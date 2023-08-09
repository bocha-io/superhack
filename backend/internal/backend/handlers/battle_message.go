package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/bocha-io/game-backend/x/messages"
	"github.com/bocha-io/garnet/x/indexer/data"
	"github.com/bocha-io/logger"
	"github.com/bocha-io/superhack/internal/garnethelpers"
	"github.com/bocha-io/txbuilder/x/txbuilder"
)

type BattleMessage struct {
	MsgType            string `json:"msgtype"`
	MatchID            string `json:"matchid"`
	PlayerOneAction    uint8  `json:"playeroneaction"`
	PlayerTwoAction    uint8  `json:"playertwoaction"`
	PlayerOneActionPos uint8  `json:"playeroneactionpos"`
	PlayerTwoActionPos uint8  `json:"playertwoactionpos"`
}

type Actions struct {
	PlayerOneSwapped bool     `json:"playeroneswapped"`
	PlayerTwoSwapped bool     `json:"playertwoswapped"`
	DamagedUnits     []string `json:"damagedunits"`
	PlayerOneAttack  int8     `json:"playeroneattack"`
	PlayerTwoAttack  int8     `json:"playertwoattack"`
}

type Values struct {
	Match   Match   `json:"match"`
	Actions Actions `json:"actions"`
}

type BattleMessageResponse struct {
	MsgType string `json:"msgtype"`
	Value   Values `json:"value"`
	Error   string `json:"error"`
}

const (
	BattleMessageType       = "battle"
	BattleMessageResponseID = "battleresponse"
)

func NewBattleMessage(
	matchID string,
	playerOneAction uint8,
	playerTwoAction uint8,
	playerOneActionPos uint8,
	playerTwoActionPos uint8,
) BattleMessage {
	return BattleMessage{
		MsgType:            BattleMessageType,
		MatchID:            matchID,
		PlayerOneAction:    playerOneAction,
		PlayerTwoAction:    playerTwoAction,
		PlayerOneActionPos: playerOneActionPos,
		PlayerTwoActionPos: playerTwoActionPos,
	}
}

func newBattleMessageError(err error) BattleMessageResponse {
	return BattleMessageResponse{
		MsgType: BattleMessageResponseID,
		Value:   Values{},
		Error:   err.Error(),
	}
}

func newBattleMessageResponse(values Values) BattleMessageResponse {
	return BattleMessageResponse{
		MsgType: BattleMessageResponseID,
		Value:   values,
		Error:   "",
	}
}

func (b *Backend) battleMessage(
	ws *messages.WebSocketContainer,
	p []byte,
) (resp BattleMessageResponse, err error) {
	// The prediction will panic if something fails in the database, catch it here
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("error getting info from the database: %v", r)
			resp = BattleMessageResponse{}
		}
	}()

	if !ws.Authenticated {
		return BattleMessageResponse{}, fmt.Errorf("user is not logged in")
	}

	var battleMsg BattleMessage
	err = json.Unmarshal(p, &battleMsg)
	if err != nil {
		return newBattleMessageError(err), err
	}

	matchID, err := txbuilder.StringToSlice(battleMsg.MatchID)
	if err != nil {
		value := fmt.Errorf("error parsing params for battle")
		logger.LogDebug(
			fmt.Sprintf("[backend] error creating transaction to battle: %s", value),
		)
		return newBattleMessageError(value), value
	}

	prediction := garnethelpers.NewPrediction(b.db)
	prediction.Battle(
		battleMsg.MatchID,
		int64(battleMsg.PlayerOneAction),
		int64(battleMsg.PlayerOneActionPos),
		int64(battleMsg.PlayerTwoAction),
		int64(battleMsg.PlayerTwoActionPos),
	)

	txhash, err := b.txBuilder.InteractWithContract(
		ws.WalletID,
		"Battle",
		matchID,
		battleMsg.PlayerOneAction,
		battleMsg.PlayerOneActionPos,
		battleMsg.PlayerTwoAction,
		battleMsg.PlayerTwoActionPos,
	)
	if err != nil {
		value := fmt.Errorf("error sending battle tx")
		return newBattleMessageError(value), value
	}

	b.db.AddTxSent(data.UnconfirmedTransaction{
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
		playerOneAttack = int8(battleMsg.PlayerOneAction)
	}

	if !playerTwoSwapped {
		playerTwoAttack = int8(battleMsg.PlayerTwoAction)
	}

	return newBattleMessageResponse(
		Values{
			Match: Match{
				MatchID:   battleMsg.MatchID,
				PlayerOne: prediction.PlayerOneGet(battleMsg.MatchID),
				PlayerTwo: prediction.PlayerTwoGet(battleMsg.MatchID),
			},
			Actions: Actions{
				PlayerOneSwapped: playerOneSwapped,
				PlayerTwoSwapped: playerTwoSwapped,
				DamagedUnits:     damaged,
				PlayerOneAttack:  playerOneAttack,
				PlayerTwoAttack:  playerTwoAttack,
			},
		},
	), nil
}
