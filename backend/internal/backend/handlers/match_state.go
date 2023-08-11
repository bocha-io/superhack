package handlers

import (
	"fmt"
	"strings"

	"github.com/bocha-io/game-backend/x/messages"
	"github.com/bocha-io/logger"
)

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

type Actions struct {
	PlayerOneSwapped bool     `json:"playeroneswapped"`
	PlayerTwoSwapped bool     `json:"playertwoswapped"`
	DamagedUnits     []string `json:"damagedunits"`
	PlayerOneAttack  int8     `json:"playeroneattack"`
	PlayerTwoAttack  int8     `json:"playertwoattack"`
	Winner           string   `json:"winner"`
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

	res.PlayerOneCurrentMon, _ = b.queryClient.GetPlayerOneCurrentMon(matchID)

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

	res.PlayerTwoCurrentMon, _ = b.queryClient.GetPlayerTwoCurrentMon(matchID)

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
