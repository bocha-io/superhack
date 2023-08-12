package handlers

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/bocha-io/game-backend/x/messages"
	"github.com/bocha-io/garnet/x/indexer/data"
	"github.com/bocha-io/superhack/internal/constants"
	"github.com/bocha-io/superhack/internal/garnethelpers"
)

type MoveMessage struct {
	MsgType string `json:"msgtype"`
	X       int32  `json:"x"`
	Y       int32  `json:"y"`
}

type MoveMessageResponse struct {
	MsgType string `json:"msgtype"`
	Value   bool   `json:"value"`
	Error   string `json:"error"`
}

const (
	MoveMessageType       = "move"
	MoveMessageResponseID = "moveresponse"
)

func NewMoveMessage(x int32, y int32) MoveMessage {
	return MoveMessage{
		MsgType: MoveMessageType,
		X:       x,
		Y:       y,
	}
}

func newMoveMessageError(err error) MoveMessageResponse {
	return MoveMessageResponse{
		MsgType: MoveMessageResponseID,
		Value:   false,
		Error:   err.Error(),
	}
}

func newMoveMessageResponse() MoveMessageResponse {
	return MoveMessageResponse{
		MsgType: MoveMessageResponseID,
		Value:   true,
		Error:   "",
	}
}

func (b *Backend) moveMessage(
	ws *messages.WebSocketContainer,
	p []byte,
) (resp MoveMessageResponse, err error) {
	// The prediction will panic if something fails in the database, catch it here
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("error getting info from the database")
			resp = MoveMessageResponse{}
		}
	}()

	if !ws.Authenticated {
		return MoveMessageResponse{}, fmt.Errorf("user is not logged in")
	}

	var moveMsg MoveMessage
	err = json.Unmarshal(p, &moveMsg)
	if err != nil {
		return newMoveMessageError(err), err
	}

	prediction := garnethelpers.NewPrediction(b.db)
	prediction.Move(int64(moveMsg.X), int64(moveMsg.Y), ws.WalletAddress)

	txhash, err := b.txBuilder.InteractWithContract(
		constants.WorldContractName,
		ws.WalletID,
		big.NewInt(0),
		"Move",
		moveMsg.X,
		moveMsg.Y,
	)
	if err != nil {
		value := fmt.Errorf("error sending move tx")
		return newMoveMessageError(value), value
	}

	b.db.AddTxSent(data.UnconfirmedTransaction{
		Txhash: txhash.Hex(),
		Events: prediction.Events,
	})

	return newMoveMessageResponse(), nil
}
