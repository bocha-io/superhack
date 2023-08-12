package handlers

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/bocha-io/game-backend/x/messages"
	"github.com/bocha-io/logger"
	"github.com/bocha-io/superhack/internal/constants"
)

type ConnectMessage struct {
	MsgType  string `json:"msgtype"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type ConnectMessageResponse struct {
	MsgType string `json:"msgtype"`
	Value   string `json:"value"`
	Error   string `json:"error"`
}

const (
	ConnectMessageType       = "connect"
	ConnectMessageResponseID = "connectresponse"
)

func NewConnectMessage(user string, password string) ConnectMessage {
	return ConnectMessage{
		MsgType:  ConnectMessageType,
		User:     user,
		Password: password,
	}
}

func newConnectMessageError(err error) ConnectMessageResponse {
	return ConnectMessageResponse{
		MsgType: ConnectMessageResponseID,
		Value:   "",
		Error:   err.Error(),
	}
}

func newConnectMessageResponse(wallet string) ConnectMessageResponse {
	return ConnectMessageResponse{
		MsgType: ConnectMessageResponseID,
		Value:   wallet,
		Error:   "",
	}
}

func (b *Backend) connectMessage(
	ws *messages.WebSocketContainer,
	p []byte,
) (ConnectMessageResponse, error) {
	var connectMsg ConnectMessage
	err := json.Unmarshal(p, &connectMsg)
	if err != nil {
		return newConnectMessageError(err), err
	}

	user, err := b.inMemoryDB.Login(connectMsg.User, connectMsg.Password)
	if err != nil {
		// Register the user
		if connectMsg.User == "" {
			value := fmt.Errorf("username can not be empty")
			return newConnectMessageError(value), value
		}

		logger.LogDebug(fmt.Sprintf("[backend] registering user %s", connectMsg.User))

		index, address, err := b.inMemoryDB.RegisterUser(
			connectMsg.User,
			connectMsg.Password,
			b.mnemonic,
		)
		if err != nil {
			return newConnectMessageError(err), err
		}

		logger.LogDebug(
			fmt.Sprintf("[backend] registered user %s with id %d", connectMsg.User, index),
		)

		logger.LogDebug(
			fmt.Sprintf(
				"[backend] registering the wallet %s in the chain %s",
				address,
				connectMsg.User,
			),
		)

		_, err = b.txBuilder.InteractWithContract(
			constants.WorldContractName,
			index,
			big.NewInt(0),
			"register",
		)
		if err != nil {
			logger.LogError(
				fmt.Sprintf("[backend] error registering wallet %s, %s", address, err.Error()),
			)
			value := fmt.Errorf("error registering the wallet")
			return newConnectMessageError(value), value
		}
		logger.LogInfo(fmt.Sprintf("[backend] wallet registered correctly %s", address))
	}

	user, err = b.inMemoryDB.Login(connectMsg.User, connectMsg.Password)
	if err != nil {
		value := fmt.Errorf("incorrect credentials")
		return newConnectMessageError(value), value
	}

	ws.User = connectMsg.User
	ws.Authenticated = true
	ws.WalletID = user.Index
	ws.WalletAddress = strings.ToLower(user.Address)

	logger.LogInfo(fmt.Sprintf("[backend] user connected: %s (%s)", ws.User, ws.WalletAddress))
	return newConnectMessageResponse(ws.WalletAddress), nil
}
