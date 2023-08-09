package handlers

import (
	"crypto/ecdsa"
	"fmt"
	"strings"

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
		logger.LogInfo(fmt.Sprintf("[backend] broadcasting position to %s", v.WalletAddress))
		if v.Conn != nil {
			_ = messages.WriteJSON(v.Conn, v.ConnMutex, status)
		}
	}
}

type Mon struct {
	MonID   string `json:"id"`
	MonHP   int64  `json:"hp"`
	MonType int64  `json:"montype"`
}

type Mons struct {
	First  Mon `json:"first"`
	Second Mon `json:"second"`
	Third  Mon `json:"third"`
}

type Battle struct {
	MsgType   string `json:"msgtype"`
	PlayerOne string `json:"playerone"`
	PlayerTwo string `json:"playertwo"`
	MatchID   string `json:"matchid"`

	Actions Actions `json:"actions"`

	PlayerOneMons       Mons   `json:"playeronemons"`
	PlayerTwoMons       Mons   `json:"playertwomons"`
	PlayerOneCurrentMon string `json:"playeronecurrentmon"`
	PlayerTwoCurrentMon string `json:"playertwocurrentmon"`
}

func (b *Backend) broadcastMatchState(
	matchID string,
	playerA string,
	playerB string,
	actions Actions,
) {
	res := Battle{
		MsgType:   "battlestatus",
		PlayerOne: playerA,
		PlayerTwo: playerB,
		MatchID:   matchID,
		Actions:   actions,
	}

	playerAKey := strings.ToLower(
		strings.Replace(playerA, "0x", "0x000000000000000000000000", 1),
	)
	res.PlayerOneMons.First.MonID, _ = b.queryClient.GetPlayerFirstMon(playerAKey)
	res.PlayerOneMons.First.MonHP, _ = b.queryClient.GetMonHp(res.PlayerOneMons.First.MonID)
	res.PlayerOneMons.First.MonType, _ = b.queryClient.GetMonSpecie(res.PlayerOneMons.First.MonID)

	res.PlayerOneMons.Second.MonID, _ = b.queryClient.GetPlayerSecondMon(playerAKey)
	res.PlayerOneMons.Second.MonHP, _ = b.queryClient.GetMonHp(res.PlayerOneMons.Second.MonID)
	res.PlayerOneMons.Second.MonType, _ = b.queryClient.GetMonSpecie(res.PlayerOneMons.Second.MonID)

	res.PlayerOneMons.Third.MonID, _ = b.queryClient.GetPlayerThirdMon(playerAKey)
	res.PlayerOneMons.Third.MonHP, _ = b.queryClient.GetMonHp(res.PlayerOneMons.Third.MonID)
	res.PlayerOneMons.Third.MonType, _ = b.queryClient.GetMonSpecie(res.PlayerOneMons.Third.MonID)

	res.PlayerOneCurrentMon, _ = b.queryClient.GetPlayerOneCurrentMon(playerAKey)

	playerBKey := strings.ToLower(
		strings.Replace(playerB, "0x", "0x000000000000000000000000", 1),
	)
	res.PlayerTwoMons.First.MonID, _ = b.queryClient.GetPlayerFirstMon(playerBKey)
	res.PlayerTwoMons.First.MonHP, _ = b.queryClient.GetMonHp(res.PlayerTwoMons.First.MonID)
	res.PlayerTwoMons.First.MonType, _ = b.queryClient.GetMonSpecie(res.PlayerTwoMons.First.MonID)

	res.PlayerTwoMons.Second.MonID, _ = b.queryClient.GetPlayerSecondMon(playerBKey)
	res.PlayerTwoMons.Second.MonHP, _ = b.queryClient.GetMonHp(res.PlayerTwoMons.Second.MonID)
	res.PlayerTwoMons.Second.MonType, _ = b.queryClient.GetMonSpecie(res.PlayerTwoMons.Second.MonID)

	res.PlayerTwoMons.Third.MonID, _ = b.queryClient.GetPlayerThirdMon(playerBKey)
	res.PlayerTwoMons.Third.MonHP, _ = b.queryClient.GetMonHp(res.PlayerTwoMons.Third.MonID)
	res.PlayerTwoMons.Third.MonType, _ = b.queryClient.GetMonSpecie(res.PlayerTwoMons.Third.MonID)

	res.PlayerTwoCurrentMon, _ = b.queryClient.GetPlayerTwoCurrentMon(playerBKey)

	for k, v := range b.wsList {
		if k == playerA || k == playerB {
			logger.LogInfo(fmt.Sprintf("[backend] broadcasting position to %s", v.WalletAddress))
			if v.Conn != nil {
				_ = messages.WriteJSON(v.Conn, v.ConnMutex, res)
			}
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
		} else {
			// TODO: remove websocket from list when connection is closed
			b.wsList[ws.WalletAddress] = ws
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

	case CreateMatchMessageType:
		if response, err := b.createMatchMessage(ws, p); err != nil {
			return err
		} else if err = messages.WriteJSON(ws.Conn, ws.ConnMutex, response); err != nil {
			return err
		} else {
			return nil
		}

	case BattleMessageType:
		if response, err := b.battleMessage(ws, p); err != nil {
			return err
		} else if err = messages.WriteJSON(ws.Conn, ws.ConnMutex, response); err != nil {
			return err
		} else {
			b.broadcastMatchState(response.Value.Match.MatchID, response.Value.Match.PlayerOne, response.Value.Match.PlayerTwo, response.Value.Actions)
			return nil
		}

	}

	return nil
}
