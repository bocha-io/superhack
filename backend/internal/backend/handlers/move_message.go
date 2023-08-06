package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/bocha-io/game-backend/x/messages"
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

func (b *Backend) moveMessage(ws *messages.WebSocketContainer, p []byte) (MoveMessageResponse, error) {
	if !ws.Authenticated {
		return MoveMessageResponse{}, fmt.Errorf("user is not logged in")
	}

	var moveMsg MoveMessage
	err := json.Unmarshal(p, &moveMsg)
	if err != nil {
		return newMoveMessageError(err), err
	}

	// TODO: autogenerate predictions and call it here!
	_, err = b.txBuilder.InteractWithContract(ws.WalletID, "move", moveMsg.X, moveMsg.Y)
	if err != nil {
		value := fmt.Errorf("error sending move tx")
		return newMoveMessageError(value), value
	}

	return newMoveMessageResponse(), nil
}
