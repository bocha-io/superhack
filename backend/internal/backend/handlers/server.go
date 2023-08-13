package handlers

import (
	"crypto/ecdsa"
	"strings"
	"sync"

	"github.com/bocha-io/game-backend/x/messages"
	"github.com/bocha-io/garnet/x/indexer/data"
	"github.com/bocha-io/superhack/internal/constants"
	"github.com/bocha-io/superhack/internal/garnethelpers"
	"github.com/bocha-io/txbuilder/x/txbuilder"
)

type Backend struct {
	mnemonic  string
	abi       []byte
	txBuilder txbuilder.TxBuilder
	db        *data.Database

	wsList map[string]*messages.WebSocketContainer
	muList *sync.Mutex

	inMemoryDB *InMemoryDatabase

	queryClient *garnethelpers.GameObject
	gameAdmins  *GameAdmins
}

func (b *Backend) AddWebSocket(wallet string, conex *messages.WebSocketContainer) {
	b.muList.Lock()
	defer b.muList.Unlock()
	b.wsList[strings.ToLower(wallet)] = conex
}

func (b *Backend) RemoveWebSocket(wallet string) {
	b.muList.Lock()
	defer b.muList.Unlock()
	delete(b.wsList, strings.ToLower(wallet))
}

func (b *Backend) Broadcast(callback func(conex *messages.WebSocketContainer)) {
	b.muList.Lock()
	defer b.muList.Unlock()
	for _, v := range b.wsList {
		callback(v)
	}
}

func (b *Backend) GetConex(wallet string) *messages.WebSocketContainer {
	b.muList.Lock()
	defer b.muList.Unlock()
	if v, ok := b.wsList[strings.ToLower(wallet)]; ok {
		return v
	}
	return nil
}

func NewBackend(
	worldABI []byte,
	db *data.Database,
	endpoint string,
	worldAddress string,
	pk *ecdsa.PrivateKey,
	userMnemonics string,

	erc20ABI []byte,
	erc20Address string,
	bridgeABI []byte,
	bridgeAddress string,
) *Backend {
	contracts := map[string]txbuilder.Contract{}
	contracts[constants.WorldContractName] = txbuilder.NewContract(
		worldAddress,
		txbuilder.NewWorldABI(worldABI),
	)
	contracts[constants.ERC20ContractName] = txbuilder.NewContract(
		erc20Address,
		txbuilder.NewWorldABI(erc20ABI),
	)
	contracts[constants.BridgeContractName] = txbuilder.NewContract(
		bridgeAddress,
		txbuilder.NewWorldABI(bridgeABI),
	)

	txBuilder := txbuilder.NexTxBuilder(
		contracts,
		endpoint,
		userMnemonics,
		map[string]uint64{"register": 500000, "Move": 140000, "sendFrom": 400000, "approve": 140000},
		20000000,
		pk,
	)

	b := &Backend{
		mnemonic:  userMnemonics,
		abi:       worldABI,
		db:        db,
		txBuilder: *txBuilder,

		wsList: map[string]*messages.WebSocketContainer{},
		muList: &sync.Mutex{},

		inMemoryDB: NewInMemoryDatabase(txBuilder),

		queryClient: garnethelpers.NewGameObject(db),
		gameAdmins:  NewGameAdmins(),
	}
	b.gameAdmins.SetBackend(b)

	return b
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
		} else {
			// TODO: remove websocket from list when connection is closed
			b.AddWebSocket(ws.WalletAddress, ws)
			b.broadcastPositions()
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

	case InventoryMessageType:
		response, err := b.inventoryMessage(ws, p)
		if err != nil {
			return err
		}
		return messages.WriteJSON(ws.Conn, ws.ConnMutex, response)

	case BridgeMessageType:
		response, err := b.bridgeMessage(ws, p)
		if err != nil {
			return err
		}
		return messages.WriteJSON(ws.Conn, ws.ConnMutex, response)

	case SwapMessageType:
		if response, err := b.swapMessage(ws, p); err != nil {
			return err
		} else if err = messages.WriteJSON(ws.Conn, ws.ConnMutex, response); err != nil {
			return err
		} else {
			return nil
		}

	case DuelRequestMessageType:
		if response, err := b.duelRequestMessage(ws, p); err != nil {
			return err
		} else if err = messages.WriteJSON(ws.Conn, ws.ConnMutex, response); err != nil {
			return err
		}

	case DuelResponseMessageType:
		return b.duelResponseMessage(ws, p)

	case SendActionMessageType:
		response, err := b.sendActionMessage(ws, p)
		if err != nil {
			return err
		}
		// Add action to the admin
		_ = b.gameAdmins.AddAction(
			response.Value.MatchID,
			ws.WalletAddress,
			response.Value.Action,
			response.Value.Pos,
		)
		return messages.WriteJSON(ws.Conn, ws.ConnMutex, response)
	}

	return nil
}
