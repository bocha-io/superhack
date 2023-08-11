package handlers

import (
	"fmt"

	"github.com/bocha-io/game-backend/x/messages"
	"github.com/bocha-io/logger"
)

type PlayerPos struct {
	X  int64  `json:"X"`
	Y  int64  `json:"Y"`
	ID string `json:"playerid"`
}

type MapStatus struct {
	Players []PlayerPos `json:"playerspos"`
	MsgType string      `json:"msgtype"`
}

func (b *Backend) broadcastPositions() {
	rows := b.queryClient.GetAllRowsPosition()
	ret := make([]PlayerPos, 0, len(rows))
	for k, v := range rows {
		x, y, err := b.queryClient.ProcessFieldsPosition(v)
		if err != nil {
			continue
		}
		ret = append(ret, PlayerPos{X: x, Y: y, ID: k})
	}

	// TODO: instead of broadcasting all the positions in the database, filter and return only players with a WS active connection
	// It would be simpler to get the current position on WS connection, cache it and update it on message move. Remove the item when the ws disconnects
	status := MapStatus{Players: ret, MsgType: "mapstatus"}

	b.Broadcast(func(conex *messages.WebSocketContainer) {
		logger.LogInfo(fmt.Sprintf("[backend] broadcasting position to %s", conex.WalletAddress))
		if conex.Conn != nil {
			_ = messages.WriteJSON(conex.Conn, conex.ConnMutex, status)
		}
	})
}
