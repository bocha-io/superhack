package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/bocha-io/game-backend/x/messages"
)

// Duel
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
) (err error) {
	if !ws.Authenticated {
		return fmt.Errorf("user is not logged in")
	}

	var duelresponseMsg DuelResponseMessage
	err = json.Unmarshal(p, &duelresponseMsg)
	if err != nil {
		return err
	}

	// We get the wallet address, so we need to add the padding to the ids
	response := newDuelResponseMessageResponse(
		Duel{
			PlayerA: duelresponseMsg.PlayerA,
			PlayerB: ws.WalletAddress,
		},
	)

	// Check if the match was created in the server
	player, err := b.gameAdmins.GetMatchRequest(response.Value.PlayerA)
	if err != nil {
		_ = messages.WriteJSON(
			ws.Conn,
			ws.ConnMutex,
			newDuelResponseMessageResponse(Duel{PlayerA: "", PlayerB: ""}),
		)
		return nil
	}

	// Invalid player
	if player != response.Value.PlayerB {
		_ = messages.WriteJSON(
			ws.Conn,
			ws.ConnMutex,
			newDuelResponseMessageResponse(Duel{PlayerA: "", PlayerB: ""}),
		)
		return nil
	}

	// Accept the match
	b.gameAdmins.AcceptMatchRequest(response.Value.PlayerA)
	// Inform the two player that the game is starting
	pA := b.GetConex(response.Value.PlayerA)
	if pA != nil {
		_ = messages.WriteJSON(pA.Conn, pA.ConnMutex, response)
	}
	pB := b.GetConex(response.Value.PlayerB)
	if pB != nil {
		_ = messages.WriteJSON(pB.Conn, pB.ConnMutex, response)
	}

	// Create a match
	matchMsg, err := createMatch(b, response.Value.PlayerA, response.Value.PlayerB)
	if err != nil {
		_ = messages.WriteJSON(
			ws.Conn,
			ws.ConnMutex,
			newDuelResponseMessageResponse(Duel{PlayerA: "", PlayerB: ""}),
		)
		return nil
	}

	// Broadcast the matchID
	if pA != nil {
		_ = messages.WriteJSON(pA.Conn, pA.ConnMutex, matchMsg)
	}
	if pB != nil {
		_ = messages.WriteJSON(pB.Conn, pB.ConnMutex, matchMsg)
	}

	// Game created
	_ = b.gameAdmins.AddAdmin(
		matchMsg.Value.MatchID,
		matchMsg.Value.PlayerOne,
		matchMsg.Value.PlayerTwo,
	)
	return nil
}
