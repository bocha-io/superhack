package handlers

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/bocha-io/garnet/x/indexer/data"
	"github.com/bocha-io/logger"
	"github.com/bocha-io/superhack/internal/garnethelpers"
	"github.com/bocha-io/txbuilder/x/txbuilder"
)

type PlayerData struct {
	PlayerID   string
	ActionType uint8
	Pos        uint8
	Set        bool
}

type GameAdmin struct {
	MatchID string
	PlayerA PlayerData
	PlayerB PlayerData

	TimeStart time.Time
	Active    bool
	mu        *sync.Mutex

	backend *Backend
}

func NewGameAdmin(matchID string, playerA string, playerB string, backend *Backend) *GameAdmin {
	g := &GameAdmin{
		MatchID: matchID,
		PlayerA: PlayerData{
			PlayerID: playerA,
			Set:      false,
		},
		PlayerB: PlayerData{
			PlayerID: playerB,
			Set:      false,
		},
		Active:    true,
		TimeStart: time.Now(),

		mu:      &sync.Mutex{},
		backend: backend,
	}

	go g.Subrutine()
	return g
}

type GameAdmins struct {
	Admins        map[string]*GameAdmin
	MatchRequests map[string]string
	mu            *sync.Mutex
	backend       *Backend
}

func (ga *GameAdmins) SetBackend(b *Backend) {
	ga.backend = b
}

func (ga *GameAdmins) AddMatchRequest(playerA string, playerB string) {
	ga.mu.Lock()
	defer ga.mu.Unlock()
	if _, ok := ga.MatchRequests[strings.ToLower(playerA)]; ok {
		ga.MatchRequests[strings.ToLower(playerA)] = strings.ToLower(playerB)
		return
	}

	ga.MatchRequests[strings.ToLower(playerA)] = strings.ToLower(playerB)
}

func (ga *GameAdmins) AcceptMatchRequest(playerA string) {
	ga.mu.Lock()
	defer ga.mu.Unlock()
	delete(ga.MatchRequests, strings.ToLower(playerA))
}

func (ga GameAdmins) GetMatchRequest(playerA string) (string, error) {
	ga.mu.Lock()
	defer ga.mu.Unlock()
	if v, ok := ga.MatchRequests[strings.ToLower(playerA)]; ok {
		return v, nil
	}
	return "", fmt.Errorf("not found")
}

func NewGameAdmins() *GameAdmins {
	return &GameAdmins{
		Admins:        map[string]*GameAdmin{},
		MatchRequests: make(map[string]string),
		mu:            &sync.Mutex{},
		backend:       nil,
	}
}

const (
	Attack = iota
	Swap
	Surrender
)

func (ga *GameAdmins) AddAdmin(matchID string, playerA string, playerB string) error {
	ga.mu.Lock()
	defer ga.mu.Unlock()
	if _, ok := ga.Admins[strings.ToLower(matchID)]; ok {
		return fmt.Errorf("match already has an admin")
	}
	ga.Admins[strings.ToLower(matchID)] = NewGameAdmin(strings.ToLower(matchID), strings.ToLower(playerA), strings.ToLower(playerB), ga.backend)
	return nil
}

func (ga *GameAdmins) RemoveAdmin(matchID string) error {
	ga.mu.Lock()
	defer ga.mu.Unlock()
	if _, ok := ga.Admins[strings.ToLower(matchID)]; !ok {
		return fmt.Errorf("already removed")
	}
	delete(ga.Admins, strings.ToLower(matchID))
	return nil
}

func (ga *GameAdmins) AddAction(matchID, player string, action uint8, pos uint8) error {
	ga.mu.Lock()
	defer ga.mu.Unlock()
	if _, ok := ga.Admins[strings.ToLower(matchID)]; ok {
		return ga.Admins[strings.ToLower(matchID)].AddAction(strings.ToLower(player), action, pos)
	}
	return fmt.Errorf("the match is not active")
}

func validateAction(action uint8) error {
	if action > 2 {
		return fmt.Errorf("invalid action")
	}
	return nil
}

func validatePos(pos uint8) error {
	if pos > 3 {
		return fmt.Errorf("invalid pos")
	}
	return nil
}

func (g *GameAdmin) AddAction(player string, action uint8, pos uint8) error {
	g.mu.Lock()
	defer g.mu.Unlock()

	if g.PlayerA.PlayerID == strings.ToLower(player) {
		if g.PlayerA.Set {
			return fmt.Errorf("the action for this turn is already set")
		}
		if err := validateAction(action); err != nil {
			return err
		}
		if err := validatePos(pos); err != nil {
			return err
		}
		g.PlayerA.Set = true
		g.PlayerA.ActionType = action
		g.PlayerA.Pos = pos

	} else if g.PlayerB.PlayerID == strings.ToLower(player) {
		if g.PlayerB.Set {
			return fmt.Errorf("the action for this turn is already set")
		}
		if err := validateAction(action); err != nil {
			return err
		}
		if err := validatePos(pos); err != nil {
			return err
		}
		g.PlayerB.Set = true
		g.PlayerB.ActionType = action
		g.PlayerB.Pos = pos
	}

	if g.PlayerA.Set && g.PlayerB.Set {
		_ = g.ExecuteAction()
	}

	return fmt.Errorf("invalid player ID")
}

func (g *GameAdmin) ExecuteAction() error {
	fmt.Println("execute action", g)

	matchID, err := txbuilder.StringToSlice(g.MatchID)
	if err != nil {
		value := fmt.Errorf("error parsing params for battle")
		logger.LogDebug(
			fmt.Sprintf("[backend] error creating transaction to battle: %s", value),
		)
		return value
	}

	prediction := garnethelpers.NewPrediction(g.backend.db)
	prediction.Battle(
		g.MatchID,
		int64(g.PlayerA.ActionType),
		int64(g.PlayerA.Pos),
		int64(g.PlayerB.ActionType),
		int64(g.PlayerB.Pos),
	)

	txhash, err := g.backend.txBuilder.InteractWithContract(
		0,
		"Battle",
		matchID,
		g.PlayerA.ActionType,
		g.PlayerA.Pos,
		g.PlayerB.ActionType,
		g.PlayerB.Pos,
	)
	if err != nil {
		value := fmt.Errorf("error sending battle tx")
		return value
	}

	g.backend.db.AddTxSent(data.UnconfirmedTransaction{
		Txhash: txhash.Hex(),
		Events: prediction.Events,
	})

	playerOneSwapped := false
	playerTwoSwapped := false
	damaged := []string{}
	for _, v := range prediction.Events {
		if v.Table == "PlayerOneCurrentMon" {
			playerOneSwapped = true
		}
		if v.Table == "PlayerTwoCurrentMon" {
			playerTwoSwapped = true
		}
		if v.Table == "MonHp" {
			damaged = append(damaged, v.Key)
		}
	}
	playerOneAttack := int8(-1)
	playerTwoAttack := int8(-1)
	if !playerOneSwapped {
		playerOneAttack = int8(g.PlayerA.ActionType)
	}

	if !playerTwoSwapped {
		playerTwoAttack = int8(g.PlayerB.ActionType)
	}

	// TODO: if match ended, player one and two will fail
	actions := Actions{
		PlayerOneSwapped: playerOneSwapped,
		PlayerTwoSwapped: playerTwoSwapped,
		DamagedUnits:     damaged,
		PlayerOneAttack:  playerOneAttack,
		PlayerTwoAttack:  playerTwoAttack,
	}

	g.backend.broadcastMatchState(g.MatchID, g.PlayerA.PlayerID, g.PlayerB.PlayerID, actions)

	// TODO: if game ended set g.Active as false
	// Reset
	g.PlayerA.Set = false
	g.PlayerA.ActionType = 0
	g.PlayerB.Set = false
	g.PlayerB.ActionType = 0
	g.TimeStart = time.Now()
	return nil
}

func (g *GameAdmin) Subrutine() {
	for g.Active {
		if g.PlayerA.Set && g.PlayerB.Set {
			g.mu.Lock()
			_ = g.ExecuteAction()
			g.mu.Unlock()
		} else if time.Now().Add(-60*time.Second).Compare(g.TimeStart) == 1 {
			// The user didn't sent the action, assume surrender
			if !g.PlayerA.Set {
				g.PlayerA.ActionType = 2
			} else {
				g.PlayerB.ActionType = 2
			}
			g.mu.Lock()
			_ = g.ExecuteAction()
			g.mu.Unlock()
			g.Active = false
		}
		time.Sleep(500 * time.Millisecond)
	}
}
