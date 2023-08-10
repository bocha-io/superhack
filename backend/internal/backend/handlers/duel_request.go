package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/bocha-io/game-backend/x/messages"
)

type DuelRequestMessage struct {
	MsgType string `json:"msgtype"`
	Enemy   string `json:"enemy"`
}

type Duel struct {
	PlayerA string `json:"playera"`
	PlayerB string `json:"playerb"`
}

type DuelRequestMessageResponse struct {
	MsgType string `json:"msgtype"`
	Value   Duel   `json:"value"`
	Error   string `json:"error"`
}

const (
	DuelRequestMessageType       = "duelrequest"
	DuelRequestMessageResponseID = "duelrequestresponse"
)

func NewDuelRequestMessage(enemy string) DuelRequestMessage {
	return DuelRequestMessage{
		MsgType: DuelRequestMessageType,
		Enemy:   enemy,
	}
}

func newDuelRequestMessageError(err error) DuelRequestMessageResponse {
	return DuelRequestMessageResponse{
		MsgType: DuelRequestMessageResponseID,
		Value:   Duel{},
		Error:   err.Error(),
	}
}

func newDuelRequestMessageResponse(duel Duel) DuelRequestMessageResponse {
	return DuelRequestMessageResponse{
		MsgType: DuelRequestMessageResponseID,
		Value:   duel,
		Error:   "",
	}
}

func (b *Backend) duelRequestMessage(
	ws *messages.WebSocketContainer,
	p []byte,
) (resp DuelRequestMessageResponse, err error) {
	if !ws.Authenticated {
		return DuelRequestMessageResponse{}, fmt.Errorf("user is not logged in")
	}

	var duelrequestMsg DuelRequestMessage
	err = json.Unmarshal(p, &duelrequestMsg)
	if err != nil {
		return newDuelRequestMessageError(err), err
	}

	// We get the wallet address, so we need to add the padding to the ids
	return newDuelRequestMessageResponse(
		Duel{
			PlayerA: ws.WalletAddress,
			PlayerB: duelrequestMsg.Enemy,
		},
	), nil
}
