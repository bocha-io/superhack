package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/bocha-io/game-backend/x/messages"
	"github.com/bocha-io/logger"
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

	validResponse := newDuelRequestMessageResponse(
		Duel{
			PlayerA: ws.WalletAddress,
			PlayerB: duelrequestMsg.Enemy,
		},
	)

	logger.LogDebug(
		fmt.Sprintf("[backend] duel request looking for enemy: %s", validResponse.Value.PlayerB),
	)
	enemy := b.GetConex(duelrequestMsg.Enemy)
	if enemy != nil {
		logger.LogDebug(
			fmt.Sprintf("[backend] duel request enemy found: %s", validResponse.Value.PlayerB),
		)
		// Send duel request to the player B
		_ = messages.WriteJSON(enemy.Conn, enemy.ConnMutex, validResponse)

		// Add this match to the pending duel list
		b.gameAdmins.AddMatchRequest(validResponse.Value.PlayerA, validResponse.Value.PlayerB)
		logger.LogInfo(
			fmt.Sprintf(
				"[backend] adding match %s vs %s",
				validResponse.Value.PlayerA,
				validResponse.Value.PlayerB,
			),
		)

		return validResponse, nil
	}

	logger.LogDebug(
		fmt.Sprintf("[backend] duel request enemy NOT found: %s!!!!", validResponse.Value.PlayerB),
	)
	// Inform that the enemy is not connected
	validResponse.Value.PlayerB = ""
	return validResponse, nil
}
