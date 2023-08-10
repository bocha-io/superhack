package handlers

import (
	"fmt"
	"sync"
	"time"
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
}

func NewGameAdmin(matchID string, playerA string, playerB string) *GameAdmin {
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

		mu: &sync.Mutex{},
	}

	go g.Subrutine()
	return g
}

type GameAdmins struct {
	Admins        map[string]*GameAdmin
	MatchRequests map[string]string
	mu            *sync.Mutex
}

func (ga *GameAdmins) AddMatchRequest(playerA string, playerB string) {
	ga.mu.Lock()
	defer ga.mu.Unlock()
	if _, ok := ga.MatchRequests[playerA]; ok {
		ga.MatchRequests[playerA] = playerB
		return
	}

	ga.MatchRequests[playerA] = playerB
}

func (ga *GameAdmins) AcceptMatchRequest(playerA string) {
	ga.mu.Lock()
	defer ga.mu.Unlock()
	delete(ga.MatchRequests, playerA)
}

func (ga GameAdmins) GetMatchRequest(playerA string) (string, error) {
	ga.mu.Lock()
	defer ga.mu.Unlock()
	if v, ok := ga.MatchRequests[playerA]; ok {
		return v, nil
	}
	return "", fmt.Errorf("not found")
}

func NewGameAdmins() *GameAdmins {
	return &GameAdmins{
		Admins:        map[string]*GameAdmin{},
		MatchRequests: make(map[string]string),
		mu:            &sync.Mutex{},
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
	if _, ok := ga.Admins[matchID]; ok {
		return fmt.Errorf("match already has an admin")
	}
	ga.Admins[matchID] = NewGameAdmin(matchID, playerA, playerB)
	return nil
}

func (ga *GameAdmins) RemoveAdmin(matchID string) error {
	ga.mu.Lock()
	defer ga.mu.Unlock()
	if _, ok := ga.Admins[matchID]; !ok {
		return fmt.Errorf("already removed")
	}
	delete(ga.Admins, matchID)
	return nil
}

func (ga *GameAdmins) AddAction(matchID, player string, action uint8, pos uint8) error {
	ga.mu.Lock()
	defer ga.mu.Unlock()
	if _, ok := ga.Admins[matchID]; ok {
		return ga.Admins[matchID].AddAction(player, action, pos)
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

	if g.PlayerA.PlayerID == player {
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

	} else if g.PlayerB.PlayerID == player {
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
		g.ExecuteAction()
	}

	return fmt.Errorf("invalid player ID")
}

func (g *GameAdmin) ExecuteAction() {
	fmt.Println("execute action", g)
	g.PlayerA.Set = false
	g.PlayerB.Set = false
	g.TimeStart = time.Now()
}

func (g *GameAdmin) Subrutine() {
	for g.Active {
		if g.PlayerA.Set && g.PlayerB.Set {
			g.mu.Lock()
			g.ExecuteAction()
			g.mu.Unlock()
		}

		if time.Now().Add(-60*time.Second).Compare(g.TimeStart) == 1 {
			// The user didn't sent the action, assume surrender
			if !g.PlayerA.Set {
				fmt.Println("surrender playerA")
			} else {
				fmt.Println("surrender playerB")
			}
			g.Active = false
		}
		time.Sleep(500 * time.Millisecond)
	}

}
