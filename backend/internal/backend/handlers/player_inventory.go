package handlers

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/bocha-io/game-backend/x/messages"
)

type InventoryMessage struct {
	MsgType string `json:"msgtype"`
}

type InventoryMessageResponse struct {
	MsgType string  `json:"msgtype"`
	Value   []int64 `json:"value"`
	Error   string  `json:"error"`
}

const (
	InventoryMessageType       = "inventory"
	InventoryMessageResponseID = "inventoryresponse"
)

func NewInventoryMessage() InventoryMessage {
	return InventoryMessage{
		MsgType: InventoryMessageType,
	}
}

func newInventoryMessageError(err error) InventoryMessageResponse {
	return InventoryMessageResponse{
		MsgType: InventoryMessageResponseID,
		Value:   []int64{},
		Error:   err.Error(),
	}
}

func newInventoryMessageResponse(mons []int64) InventoryMessageResponse {
	return InventoryMessageResponse{
		MsgType: InventoryMessageResponseID,
		Value:   mons,
		Error:   "",
	}
}

func (b *Backend) inventoryMessage(
	ws *messages.WebSocketContainer,
	p []byte,
) (resp InventoryMessageResponse, err error) {
	// The prediction will panic if something fails in the database, catch it here
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("error getting info from the database")
			resp = InventoryMessageResponse{}
		}
	}()

	if !ws.Authenticated {
		return InventoryMessageResponse{}, fmt.Errorf("user is not logged in")
	}

	var inventoryMsg InventoryMessage
	err = json.Unmarshal(p, &inventoryMsg)
	if err != nil {
		return newInventoryMessageError(err), err
	}

	res := make([]int64, 0, 3)
	temp := strings.Replace(ws.WalletAddress, "0x", "0x000000000000000000000000", 1)
	mon, err := b.queryClient.GetInventoryFirstMon(temp)
	if err != nil {
		return newInventoryMessageResponse([]int64{}), nil
	}
	res = append(res, mon)

	mon, err = b.queryClient.GetInventorySecondMon(temp)
	if err != nil {
		return newInventoryMessageResponse([]int64{}), nil
	}
	res = append(res, mon)

	mon, err = b.queryClient.GetInventoryThirdMon(temp)
	if err != nil {
		return newInventoryMessageResponse([]int64{}), nil
	}
	res = append(res, mon)

	return newInventoryMessageResponse(res), nil
}
