package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/bocha-io/game-backend/x/messages"
)

type DuelResponseMessage struct {
	MsgType  string `json:"msgtype"`
	Response bool   `json:"response"`
	PlayerA  string `json:"playera"`
}

type DuelResponseMessageResponse struct {
	MsgType string `json:"msgtype"`
	Value   Duel   `json:"value"`
	Error   string `json:"error"`
}

const (
	DuelResponseMessageType       = "duelresponse"
	DuelResponseMessageResponseID = "duelresponseresponse"
)

func NewDuelResponseMessage(playerA string, response bool) DuelResponseMessage {
	return DuelResponseMessage{
		MsgType:  DuelResponseMessageType,
		PlayerA:  playerA,
		Response: response,
	}
}

func newDuelResponseMessageError(err error) DuelResponseMessageResponse {
	return DuelResponseMessageResponse{
		MsgType: DuelResponseMessageResponseID,
		Value:   Duel{},
		Error:   err.Error(),
	}
}

func newDuelResponseMessageResponse(duel Duel) DuelResponseMessageResponse {
	return DuelResponseMessageResponse{
		MsgType: DuelResponseMessageResponseID,
		Value:   duel,
		Error:   "",
	}
}

func (b *Backend) duelResponseMessage(
	ws *messages.WebSocketContainer,
	p []byte,
) (resp DuelResponseMessageResponse, err error) {
	if !ws.Authenticated {
		return DuelResponseMessageResponse{}, fmt.Errorf("user is not logged in")
	}

	var duelresponseMsg DuelResponseMessage
	err = json.Unmarshal(p, &duelresponseMsg)
	if err != nil {
		return newDuelResponseMessageError(err), err
	}

	// We get the wallet address, so we need to add the padding to the ids
	return newDuelResponseMessageResponse(
		Duel{
			PlayerA: duelresponseMsg.PlayerA,
			PlayerB: ws.WalletAddress,
		},
	), nil
}
