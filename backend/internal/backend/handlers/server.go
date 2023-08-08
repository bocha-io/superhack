package handlers

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/bocha-io/game-backend/x/messages"
	"github.com/bocha-io/garnet/x/indexer/data"
	"github.com/bocha-io/logger"
	"github.com/bocha-io/superhack/internal/garnethelpers"
	"github.com/bocha-io/txbuilder/x/txbuilder"
)

type Backend struct {
	mnemonic  string
	abi       []byte
	txBuilder txbuilder.TxBuilder
	db        *data.Database
	wsList    map[string]*messages.WebSocketContainer

	inMemoryDB *InMemoryDatabase

	queryClient *garnethelpers.GameObject
}

func NewBackend(
	worldABI []byte,
	db *data.Database,
	endpoint string,
	worldAddress string,
	pk *ecdsa.PrivateKey,
	userMnemonics string,
) *Backend {
	txBuilder := txbuilder.NexTxBuilder(
		worldAddress,
		txbuilder.NewWorldABI(worldABI),
		endpoint,
		userMnemonics,
		map[string]uint64{},
		20000000,
		pk,
	)

	// TODO: move this function to another place and only send coins if they are needed
	// Send coins to 10 accounts
	i := 0
	for i <= 10 {
		_, errFaucet := txBuilder.FoundAccount(i)
		logger.LogInfo(fmt.Sprintf("[backend] sending coins to wallet: %d", i))
		if errFaucet != nil {
			logger.LogError(
				fmt.Sprintf("[backend] error sending coins to wallet %d, %s", i, errFaucet.Error()),
			)
		}
		i++
	}

	b := &Backend{
		mnemonic:  userMnemonics,
		abi:       worldABI,
		db:        db,
		txBuilder: *txBuilder,
		wsList:    map[string]*messages.WebSocketContainer{},

		inMemoryDB:  NewInMemoryDatabase(txBuilder),
		queryClient: garnethelpers.NewGameObject(db),
	}

	return b
}

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

	for _, v := range b.wsList {
		logger.LogInfo(fmt.Sprintf("[TEST] broadcasting to %s", v.WalletAddress))
		if v.Conn != nil {
			_ = messages.WriteJSON(v.Conn, v.ConnMutex, status)
			logger.LogInfo(fmt.Sprintf("[TEST] msg sent %s", v.WalletAddress))
		}
	}
}

func (b *Backend) HandleMessage(
	g *messages.Server,
	ws *messages.WebSocketContainer,
	m messages.BasicMessage,
	p []byte,
) error {
	_ = g
	switch m.MsgType {
	case ConnectMessageType:
		if msg, err := b.connectMessage(ws, p); err != nil {
			return err
		} else if err = messages.WriteJSON(ws.Conn, ws.ConnMutex, msg); err != nil {
			return err
		}

	case MoveMessageType:
		if response, err := b.moveMessage(ws, p); err != nil {
			return err
		} else if err = messages.WriteJSON(ws.Conn, ws.ConnMutex, response); err != nil {
			return err
		} else {
			// Movement was ok, broadcast position of all users to the subscribers
			b.broadcastPositions()
			return nil
		}
	}

	return nil
}
