package handlers

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/bocha-io/game-backend/x/messages"
	"github.com/bocha-io/garnet/x/indexer/data"
	"github.com/bocha-io/logger"
	"github.com/bocha-io/txbuilder/x/txbuilder"
)

type Backend struct {
	mnemonic  string
	abi       []byte
	txBuilder txbuilder.TxBuilder
	db        *data.Database
	wsList    map[string]*messages.WebSocketContainer

	inMemoryDB *InMemoryDatabase
}

func NewBackend(worldABI []byte, db *data.Database, endpoint string, worldAddress string, pk *ecdsa.PrivateKey, userMnemonics string) *Backend {
	txBuilder := txbuilder.NexTxBuilder(
		worldAddress,
		txbuilder.NewWorldABI(worldABI),
		endpoint,
		userMnemonics,
		map[string]uint64{},
		20000000,
		pk,
	)

	// TODO: move this functions to another place and only send coins if they are needed
	// Send coins to 10 accounts
	i := 0
	for i <= 10 {
		_, errFaucet := txBuilder.FoundAccount(i)
		logger.LogInfo(fmt.Sprintf("[backend] sending coins to wallet: %d", i))
		if errFaucet != nil {
			logger.LogError(fmt.Sprintf("[backend] error sending coins to wallet %d, %s", i, errFaucet.Error()))
		}
		i++
	}

	b := &Backend{
		mnemonic:  userMnemonics,
		abi:       worldABI,
		db:        db,
		txBuilder: *txBuilder,
		wsList:    map[string]*messages.WebSocketContainer{},

		inMemoryDB: NewInMemoryDatabase(txBuilder),
	}

	return b
}

func (b *Backend) HandleMessage(g *messages.Server, ws *messages.WebSocketContainer, m messages.BasicMessage, p []byte) error {
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
		}
	}

	return nil
}
