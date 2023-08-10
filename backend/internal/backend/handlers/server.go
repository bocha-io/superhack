package handlers

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"strings"
	"sync"

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

		wsList: map[string]*messages.WebSocketContainer{},
		muList: &sync.Mutex{},

		inMemoryDB: NewInMemoryDatabase(txBuilder),

		queryClient: garnethelpers.NewGameObject(db),
		gameAdmins:  NewGameAdmins(),
	}
	b.gameAdmins.SetBackend(b)

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

	b.Broadcast(func(conex *messages.WebSocketContainer) {
		logger.LogInfo(fmt.Sprintf("[backend] broadcasting position to %s", conex.WalletAddress))
		if conex.Conn != nil {
			_ = messages.WriteJSON(conex.Conn, conex.ConnMutex, status)
		}
	})
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

	conex := b.GetConex(playerA)
	if conex != nil {
		logger.LogInfo(fmt.Sprintf("[backend] broadcasting position to %s", conex.WalletAddress))
		if conex.Conn != nil {
			_ = messages.WriteJSON(conex.Conn, conex.ConnMutex, res)
		}
	}

	conex = b.GetConex(playerB)
	if conex != nil {
		logger.LogInfo(fmt.Sprintf("[backend] broadcasting position to %s", conex.WalletAddress))
		if conex.Conn != nil {
			_ = messages.WriteJSON(conex.Conn, conex.ConnMutex, res)
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

	case DuelRequestMessageType:

		if response, err := b.duelRequestMessage(ws, p); err != nil {
			return err
		} else {
			logger.LogDebug(fmt.Sprintf("[backend] duel request looking for enemy: %s", response.Value.PlayerB))
			enemy := b.GetConex(response.Value.PlayerB)
			if enemy != nil {
				logger.LogDebug(fmt.Sprintf("[backend] duel request enemy found: %s", response.Value.PlayerB))
				// Send duel request to the player B
				_ = messages.WriteJSON(enemy.Conn, enemy.ConnMutex, response)

				// Add this match to the pending duel list
				b.gameAdmins.AddMatchRequest(response.Value.PlayerA, response.Value.PlayerB)
				logger.LogInfo(fmt.Sprintf("[backend] adding match %s vs %s", response.Value.PlayerA, response.Value.PlayerB))

				// Inform that the request was sent
				if err = messages.WriteJSON(ws.Conn, ws.ConnMutex, response); err != nil {
					return err
				}
			} else {
				logger.LogDebug(fmt.Sprintf("[backend] duel request enemy NOT found: %s!!!!", response.Value.PlayerB))
				// Inform that the enemy is not connected
				if err = messages.WriteJSON(ws.Conn, ws.ConnMutex, newDuelRequestMessageError(fmt.Errorf("player b is not connected"))); err != nil {
					return nil
				}
			}
		}

	case DuelResponseMessageType:
		if response, err := b.duelResponseMessage(ws, p); err != nil {
			return err
		} else {
			if player, err := b.gameAdmins.GetMatchRequest(response.Value.PlayerA); err == nil {
				if player == response.Value.PlayerB {
					b.gameAdmins.AcceptMatchRequest(response.Value.PlayerA)
					pA := b.GetConex(response.Value.PlayerA)
					if pA != nil {
						_ = messages.WriteJSON(pA.Conn, pA.ConnMutex, response)
					}
					pB := b.GetConex(response.Value.PlayerB)
					if pB != nil {
						_ = messages.WriteJSON(pB.Conn, pB.ConnMutex, response)
					}

					// Create a match
					msg := NewCreateMatchMessage(response.Value.PlayerA, response.Value.PlayerB)
					bMsg, _ := json.Marshal(msg)

					if response, err := b.createMatchMessage(ws, bMsg); err != nil {
						// TODO, maybe return something different
						if pA != nil {
							_ = messages.WriteJSON(pA.Conn, pA.ConnMutex, response)
						}
						if pB != nil {
							_ = messages.WriteJSON(pB.Conn, pB.ConnMutex, response)
						}
						return nil
					} else {
						// Game created
						_ = b.gameAdmins.AddAdmin(response.Value.MatchID, response.Value.PlayerOne, response.Value.PlayerTwo)
						return nil
					}
				}
				return fmt.Errorf("invalid player")
			} else {
				return err
			}
		}

	case SendActionMessageType:
		if response, err := b.sendActionMessage(ws, p); err != nil {
			return err
		} else {
			// Add action to the admin
			_ = b.gameAdmins.AddAction(response.Value.MatchID, ws.WalletAddress, response.Value.Action, response.Value.Pos)
			if err = messages.WriteJSON(ws.Conn, ws.ConnMutex, response); err != nil {
				return err
			}
			return nil
		}

		// case CreateMatchMessageType:
		// 	if response, err := b.createMatchMessage(ws, p); err != nil {
		// 		return err
		// 	} else if err = messages.WriteJSON(ws.Conn, ws.ConnMutex, response); err != nil {
		// 		return err
		// 	} else {
		// 		return nil
		// 	}

		// case BattleMessageType:
		// 	if response, err := b.battleMessage(ws, p); err != nil {
		// 		return err
		// 	} else if err = messages.WriteJSON(ws.Conn, ws.ConnMutex, response); err != nil {
		// 		return err
		// 	} else {
		// 		b.broadcastMatchState(response.Value.Match.MatchID, response.Value.Match.PlayerOne, response.Value.Match.PlayerTwo, response.Value.Actions)
		// 		return nil
		// 	}

	}

	return nil
}
