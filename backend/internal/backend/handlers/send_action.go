package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/bocha-io/game-backend/x/messages"
)

type msgValue struct {
	MatchID string `json:"matchid"`
	Action  uint8  `json:"action"`
	Pos     uint8  `json:"pos"`
}

type SendActionMessage struct {
	MsgType string `json:"msgtype"`
	MatchID string `json:"matchid"`
	Action  uint8  `json:"action"`
	Pos     uint8  `json:"pos"`
}

type SendActionMessageResponse struct {
	MsgType string   `json:"msgtype"`
	Value   msgValue `json:"value"`
	Error   string   `json:"error"`
}

const (
	SendActionMessageType       = "sendaction"
	SendActionMessageResponseID = "sendactionresponse"
)

func NewSendActionMessage(matchID string, action uint8, pos uint8) SendActionMessage {
	return SendActionMessage{
		MsgType: SendActionMessageType,
		MatchID: matchID,
		Action:  action,
		Pos:     pos,
	}
}

func newSendActionMessageError(err error) SendActionMessageResponse {
	return SendActionMessageResponse{
		MsgType: SendActionMessageResponseID,
		Value:   msgValue{},
		Error:   err.Error(),
	}
}

func newSendActionMessageResponse(msg SendActionMessage) SendActionMessageResponse {
	return SendActionMessageResponse{
		MsgType: SendActionMessageResponseID,
		Value: msgValue{
			MatchID: msg.MatchID,
			Action:  msg.Action,
			Pos:     msg.Pos,
		},
		Error: "",
	}
}

func (b *Backend) sendActionMessage(
	ws *messages.WebSocketContainer,
	p []byte,
) (resp SendActionMessageResponse, err error) {
	// The prediction will panic if something fails in the database, catch it here
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("error getting info from the database: %v", r)
			resp = SendActionMessageResponse{}
		}
	}()

	if !ws.Authenticated {
		return SendActionMessageResponse{}, fmt.Errorf("user is not logged in")
	}

	var sendactionMsg SendActionMessage
	err = json.Unmarshal(p, &sendactionMsg)
	if err != nil {
		return newSendActionMessageError(err), err
	}

	return newSendActionMessageResponse(sendactionMsg), nil
}
