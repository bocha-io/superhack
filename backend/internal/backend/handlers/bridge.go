package handlers

import (
	"encoding/json"
	"fmt"
	"math/big"
	"time"

	"github.com/bocha-io/game-backend/x/messages"
	"github.com/bocha-io/logger"
	"github.com/bocha-io/superhack/internal/constants"
	"github.com/ethereum/go-ethereum/common"
)

type BridgeMessage struct {
	MsgType string `json:"msgtype"`
	Amount  int64  `json:"amount"`
}

type BridgeMessageResponse struct {
	MsgType string `json:"msgtype"`
	Value   string `json:"value"`
	Error   string `json:"error"`
}

const (
	BridgeMessageType       = "bridge"
	BridgeMessageResponseID = "bridgeresponse"
)

func NewBridgeMessage() BridgeMessage {
	return BridgeMessage{
		MsgType: BridgeMessageType,
	}
}

func newBridgeMessageError(err error) BridgeMessageResponse {
	return BridgeMessageResponse{
		MsgType: BridgeMessageResponseID,
		Value:   "",
		Error:   err.Error(),
	}
}

func newBridgeMessageResponse(txhash string) BridgeMessageResponse {
	return BridgeMessageResponse{
		MsgType: BridgeMessageResponseID,
		Value:   txhash,
		Error:   "",
	}
}

func (b *Backend) bridgeMessage(
	ws *messages.WebSocketContainer,
	p []byte,
) (resp BridgeMessageResponse, err error) {
	// The prediction will panic if something fails in the database, catch it here
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("error getting info from the database %v", r)
			resp = BridgeMessageResponse{}
		}
	}()

	if !ws.Authenticated {
		return BridgeMessageResponse{}, fmt.Errorf("user is not logged in")
	}

	var bridgeMsg BridgeMessage
	err = json.Unmarshal(p, &bridgeMsg)
	if err != nil {
		return newBridgeMessageError(err), err
	}

	bridgeAddress := common.HexToAddress("0xE77710Ae15c5F9F1b8E31135ca4f5FBe5bEc2097")
	paddedBridgeAddress := common.LeftPadBytes(bridgeAddress.Bytes(), 32)
	var sliceBridgeAddress [32]byte
	copy(sliceBridgeAddress[:], paddedBridgeAddress)

	toAddress := common.HexToAddress(ws.WalletAddress)
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)

	var sliceAddress [32]byte
	copy(sliceAddress[:], paddedAddress)

	emptyAddress := common.HexToAddress("0x0000000000000000000000000000000000000000")
	paddedEmptyAddress := common.LeftPadBytes(emptyAddress.Bytes(), 32)
	var sliceEmptyAddress [32]byte
	copy(sliceEmptyAddress[:], paddedEmptyAddress)

	amount := new(big.Int)
	amount.SetString(fmt.Sprintf("%d000000000000000000", bridgeMsg.Amount), 10)

	// var emptyByteArray [0]byte
	// Create allow transaction
	txhash, err := b.txBuilder.InteractWithContract(
		constants.ERC20ContractName,
		ws.WalletID,
		big.NewInt(0),
		"approve",
		sliceBridgeAddress,
		amount,
	)
	if err != nil {
		logger.LogError(fmt.Sprintf("[APPROVE] error sending approve: %s", err.Error()))
		return newBridgeMessageError(err), err
	}
	logger.LogError(fmt.Sprintf("[APPROVE] sent with hash: %s", txhash))

	// Wait 5 seconds
	time.Sleep(5 * time.Second)

	// Create bridge transaction
	txhash, err = b.txBuilder.InteractWithContract(
		constants.BridgeContractName,
		ws.WalletID,
		big.NewInt(94313772364324),
		"sendFrom",
		sliceAddress,
		uint16(184),
		toAddress.Bytes(),
		amount,
		sliceAddress,
		sliceEmptyAddress,
		[]byte{},
	)
	if err != nil {
		logger.LogError(fmt.Sprintf("[BRIDGE] error sending bridge: %s", err.Error()))
		return newBridgeMessageError(err), err
	}

	return newBridgeMessageResponse(txhash.Hex()), nil
}
