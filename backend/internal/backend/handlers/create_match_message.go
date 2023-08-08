package handlers

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/bocha-io/game-backend/x/messages"
	"github.com/bocha-io/logger"
	"github.com/bocha-io/txbuilder/x/txbuilder"
)

type CreateMatchMessage struct {
	MsgType string `json:"msgtype"`
	PlayerA string `json:"playera"`
	PlayerB string `json:"playerb"`
}

type CreateMatchMessageResponse struct {
	MsgType string `json:"msgtype"`
	Value   bool   `json:"value"`
	Error   string `json:"error"`
}

const (
	CreateMatchMessageType       = "creatematch"
	CreateMatchMessageResponseID = "creatematchresponse"
)

func NewCreateMatchMessage(playerA string, playerB string) CreateMatchMessage {
	return CreateMatchMessage{
		MsgType: CreateMatchMessageType,
		PlayerA: playerA,
		PlayerB: playerB,
	}
}

func newCreateMatchMessageError(err error) CreateMatchMessageResponse {
	return CreateMatchMessageResponse{
		MsgType: CreateMatchMessageResponseID,
		Value:   false,
		Error:   err.Error(),
	}
}

func newCreateMatchMessageResponse() CreateMatchMessageResponse {
	return CreateMatchMessageResponse{
		MsgType: CreateMatchMessageResponseID,
		Value:   true,
		Error:   "",
	}
}

func (b *Backend) createMatchMessage(
	ws *messages.WebSocketContainer,
	p []byte,
) (resp CreateMatchMessageResponse, err error) {
	// The prediction will panic if something fails in the database, catch it here
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("error getting info from the database")
			resp = CreateMatchMessageResponse{}
		}
	}()

	if !ws.Authenticated {
		return CreateMatchMessageResponse{}, fmt.Errorf("user is not logged in")
	}

	var creatematchMsg CreateMatchMessage
	err = json.Unmarshal(p, &creatematchMsg)
	if err != nil {
		return newCreateMatchMessageError(err), err
	}

	playerA, err := txbuilder.StringToSlice(strings.Replace(creatematchMsg.PlayerA, "0x", "0x000000000000000000000000", 1))
	if err != nil {
		value := fmt.Errorf("error parsing params for create match")
		logger.LogDebug(fmt.Sprintf("[backend] error creating transaction to create match: %s", value))
		return newCreateMatchMessageError(value), value
	}

	playerB, err := txbuilder.StringToSlice(strings.Replace(creatematchMsg.PlayerB, "0x", "0x000000000000000000000000", 1))
	if err != nil {
		value := fmt.Errorf("error parsing params for create match")
		logger.LogDebug(fmt.Sprintf("[backend] error creating transaction to create match: %s", value))
		return newCreateMatchMessageError(value), value
	}

	_, err = b.txBuilder.InteractWithContract(ws.WalletID, "CreateMatch", playerA, playerB)
	if err != nil {
		value := fmt.Errorf("error sending creatematch tx")
		return newCreateMatchMessageError(value), value
	}

	return newCreateMatchMessageResponse(), nil
}