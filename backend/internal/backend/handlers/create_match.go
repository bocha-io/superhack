package handlers

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/bocha-io/garnet/x/indexer/data"
	"github.com/bocha-io/logger"
	"github.com/bocha-io/superhack/internal/constants"
	"github.com/bocha-io/superhack/internal/garnethelpers"
	"github.com/bocha-io/txbuilder/x/txbuilder"
)

// Match
type Match struct {
	MatchID   string `json:"id"`
	PlayerOne string `json:"playerone"`
	PlayerTwo string `json:"playertwo"`
}

type CreateMatchMessageResponse struct {
	MsgType string `json:"msgtype"`
	Value   Match  `json:"value"`
	Error   string `json:"error"`
}

const (
	CreateMatchMessageResponseID = "creatematchresponse"
)

func newCreateMatchMessageError(err error) CreateMatchMessageResponse {
	return CreateMatchMessageResponse{
		MsgType: CreateMatchMessageResponseID,
		Value:   Match{},
		Error:   err.Error(),
	}
}

func newCreateMatchMessageResponse(match Match) CreateMatchMessageResponse {
	return CreateMatchMessageResponse{
		MsgType: CreateMatchMessageResponseID,
		Value:   match,
		Error:   "",
	}
}

func createMatch(
	b *Backend,
	playerA string,
	playerB string,
) (resp CreateMatchMessageResponse, err error) {
	// The prediction will panic if something fails in the database, catch it here
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("error getting info from the database: %v", r)
			resp = CreateMatchMessageResponse{}
		}
	}()

	paddedPlayerA := strings.ToLower(
		strings.Replace(playerA, "0x", "0x000000000000000000000000", 1),
	)
	playerABytes, err := txbuilder.StringToSlice(paddedPlayerA)
	if err != nil {
		value := fmt.Errorf("error parsing params for create match")
		logger.LogDebug(
			fmt.Sprintf("[backend] error creating transaction to create match: %s", value),
		)
		return newCreateMatchMessageError(value), value
	}

	paddedPlayerB := strings.ToLower(
		strings.Replace(playerB, "0x", "0x000000000000000000000000", 1),
	)
	playerBBytes, err := txbuilder.StringToSlice(paddedPlayerB)
	if err != nil {
		value := fmt.Errorf("error parsing params for create match")
		logger.LogDebug(
			fmt.Sprintf("[backend] error creating transaction to create match: %s", value),
		)
		return newCreateMatchMessageError(value), value
	}

	prediction := garnethelpers.NewPrediction(b.db)
	prediction.CreateMatch(paddedPlayerA, paddedPlayerB)

	txhash, err := b.txBuilder.InteractWithContract(
		constants.WorldContractName,
		0,
		big.NewInt(0),
		"CreateMatch",
		playerABytes,
		playerBBytes,
	)
	if err != nil {
		value := fmt.Errorf("error sending creatematch tx")
		return newCreateMatchMessageError(value), value
	}

	b.db.AddTxSent(data.UnconfirmedTransaction{
		Txhash: txhash.Hex(),
		Events: prediction.Events,
	})

	matchID := ""
	for _, v := range prediction.Events {
		if v.Table == "Match" {
			matchID = v.Key
		}
	}
	logger.LogInfo(fmt.Sprintf("[backend] match created with id %s", matchID))

	return newCreateMatchMessageResponse(
		Match{
			MatchID:   matchID,
			PlayerOne: playerA,
			PlayerTwo: playerB,
		},
	), nil
}
