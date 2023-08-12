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

type SwapMessage struct {
	MsgType string `json:"msgtype"`
	MonType uint8  `json:"montype"`
	Pos     uint8  `json:"pos"`
}

type SwapMessageResponse struct {
	MsgType string `json:"msgtype"`
	Value   bool   `json:"value"`
	Error   string `json:"error"`
}

const (
	SwapMessageType       = "swap"
	SwapMessageResponseID = "swapresponse"
)

func NewSwapMessage(mon uint8, pos uint8) SwapMessage {
	return SwapMessage{
		MsgType: SwapMessageType,
		MonType: mon,
		Pos:     pos,
	}
}

func newSwapMessageError(err error) SwapMessageResponse {
	return SwapMessageResponse{
		MsgType: SwapMessageResponseID,
		Value:   false,
		Error:   err.Error(),
	}
}

func newSwapMessageResponse() SwapMessageResponse {
	return SwapMessageResponse{
		MsgType: SwapMessageResponseID,
		Value:   true,
		Error:   "",
	}
}

func (b *Backend) swapMessage(
	ws *messages.WebSocketContainer,
	p []byte,
) (resp SwapMessageResponse, err error) {
	// The prediction will panic if something fails in the database, catch it here
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("error getting info from the database")
			resp = SwapMessageResponse{}
		}
	}()

	if !ws.Authenticated {
		return SwapMessageResponse{}, fmt.Errorf("user is not logged in")
	}

	var swapMsg SwapMessage
	err = json.Unmarshal(p, &swapMsg)
	if err != nil {
		return newSwapMessageError(err), err
	}

	prediction := garnethelpers.NewPrediction(b.db)
	prediction.SwapMon(int64(swapMsg.MonType), int64(swapMsg.Pos), ws.WalletAddress)

	fmt.Println(swapMsg.MonType)
	fmt.Println(swapMsg.Pos)

	txhash, err := b.txBuilder.InteractWithContract(
		constants.WorldContractName,
		ws.WalletID,
		big.NewInt(0),
		"SwapMon",
		swapMsg.MonType,
		swapMsg.Pos,
	)
	if err != nil {
		fmt.Println(err)
		value := fmt.Errorf("error sending swap tx")
		return newSwapMessageError(value), value
	}

	b.db.AddTxSent(data.UnconfirmedTransaction{
		Txhash: txhash.Hex(),
		Events: prediction.Events,
	})

	return newSwapMessageResponse(), nil
}
