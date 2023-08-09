package garnethelpers

import (
	"math/big"

	"github.com/bocha-io/garnet/x/indexer/data"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func BytesEventFromString(val string) []byte {
	ret, err := hexutil.Decode(val)
	if err != nil {
		panic(err.Error())
	}
	return ret
}

func CreatePlayerEvent(ID string, value bool) data.MudEvent {
	return data.MudEvent{
		Table: "Player",
		Key:   ID,
		Fields: []data.Field{
			{Key: "value", Data: data.BoolField{Data: value}},
		},
	}
}

func DeletePlayerEvent(ID string) data.MudEvent {
	return data.MudEvent{
		Table:  "Player",
		Key:    ID,
		Fields: []data.Field{},
	}
}

func CreateStatusEvent(ID string, value int64) data.MudEvent {
	return data.MudEvent{
		Table: "Status",
		Key:   ID,
		Fields: []data.Field{
			{Key: "value", Data: data.UintField{Data: *big.NewInt(value)}},
		},
	}
}

func DeleteStatusEvent(ID string) data.MudEvent {
	return data.MudEvent{
		Table:  "Status",
		Key:    ID,
		Fields: []data.Field{},
	}
}

func CreatePositionEvent(ID string, x int64, y int64) data.MudEvent {
	return data.MudEvent{
		Table: "Position",
		Key:   ID,
		Fields: []data.Field{
			{Key: "x", Data: data.UintField{Data: *big.NewInt(x)}},
			{Key: "y", Data: data.UintField{Data: *big.NewInt(y)}},
		},
	}
}

func DeletePositionEvent(ID string) data.MudEvent {
	return data.MudEvent{
		Table:  "Position",
		Key:    ID,
		Fields: []data.Field{},
	}
}

func CreateMatchEvent(ID string, value bool) data.MudEvent {
	return data.MudEvent{
		Table: "Match",
		Key:   ID,
		Fields: []data.Field{
			{Key: "value", Data: data.BoolField{Data: value}},
		},
	}
}

func DeleteMatchEvent(ID string) data.MudEvent {
	return data.MudEvent{
		Table:  "Match",
		Key:    ID,
		Fields: []data.Field{},
	}
}

func CreateMatchResultEvent(ID string, winner string, loser string) data.MudEvent {
	return data.MudEvent{
		Table: "MatchResult",
		Key:   ID,
		Fields: []data.Field{
			{Key: "winner", Data: data.NewBytesField(BytesEventFromString(winner))},
			{Key: "loser", Data: data.NewBytesField(BytesEventFromString(loser))},
		},
	}
}

func DeleteMatchResultEvent(ID string) data.MudEvent {
	return data.MudEvent{
		Table:  "MatchResult",
		Key:    ID,
		Fields: []data.Field{},
	}
}

func CreatePlayerOneEvent(ID string, value string) data.MudEvent {
	return data.MudEvent{
		Table: "PlayerOne",
		Key:   ID,
		Fields: []data.Field{
			{Key: "value", Data: data.NewBytesField(BytesEventFromString(value))},
		},
	}
}

func DeletePlayerOneEvent(ID string) data.MudEvent {
	return data.MudEvent{
		Table:  "PlayerOne",
		Key:    ID,
		Fields: []data.Field{},
	}
}

func CreatePlayerTwoEvent(ID string, value string) data.MudEvent {
	return data.MudEvent{
		Table: "PlayerTwo",
		Key:   ID,
		Fields: []data.Field{
			{Key: "value", Data: data.NewBytesField(BytesEventFromString(value))},
		},
	}
}

func DeletePlayerTwoEvent(ID string) data.MudEvent {
	return data.MudEvent{
		Table:  "PlayerTwo",
		Key:    ID,
		Fields: []data.Field{},
	}
}

func CreatePlayerOneCurrentMonEvent(ID string, value string) data.MudEvent {
	return data.MudEvent{
		Table: "PlayerOneCurrentMon",
		Key:   ID,
		Fields: []data.Field{
			{Key: "value", Data: data.NewBytesField(BytesEventFromString(value))},
		},
	}
}

func DeletePlayerOneCurrentMonEvent(ID string) data.MudEvent {
	return data.MudEvent{
		Table:  "PlayerOneCurrentMon",
		Key:    ID,
		Fields: []data.Field{},
	}
}

func CreatePlayerTwoCurrentMonEvent(ID string, value string) data.MudEvent {
	return data.MudEvent{
		Table: "PlayerTwoCurrentMon",
		Key:   ID,
		Fields: []data.Field{
			{Key: "value", Data: data.NewBytesField(BytesEventFromString(value))},
		},
	}
}

func DeletePlayerTwoCurrentMonEvent(ID string) data.MudEvent {
	return data.MudEvent{
		Table:  "PlayerTwoCurrentMon",
		Key:    ID,
		Fields: []data.Field{},
	}
}

func CreatePlayerFirstMonEvent(ID string, value string) data.MudEvent {
	return data.MudEvent{
		Table: "PlayerFirstMon",
		Key:   ID,
		Fields: []data.Field{
			{Key: "value", Data: data.NewBytesField(BytesEventFromString(value))},
		},
	}
}

func DeletePlayerFirstMonEvent(ID string) data.MudEvent {
	return data.MudEvent{
		Table:  "PlayerFirstMon",
		Key:    ID,
		Fields: []data.Field{},
	}
}

func CreatePlayerSecondMonEvent(ID string, value string) data.MudEvent {
	return data.MudEvent{
		Table: "PlayerSecondMon",
		Key:   ID,
		Fields: []data.Field{
			{Key: "value", Data: data.NewBytesField(BytesEventFromString(value))},
		},
	}
}

func DeletePlayerSecondMonEvent(ID string) data.MudEvent {
	return data.MudEvent{
		Table:  "PlayerSecondMon",
		Key:    ID,
		Fields: []data.Field{},
	}
}

func CreatePlayerThirdMonEvent(ID string, value string) data.MudEvent {
	return data.MudEvent{
		Table: "PlayerThirdMon",
		Key:   ID,
		Fields: []data.Field{
			{Key: "value", Data: data.NewBytesField(BytesEventFromString(value))},
		},
	}
}

func DeletePlayerThirdMonEvent(ID string) data.MudEvent {
	return data.MudEvent{
		Table:  "PlayerThirdMon",
		Key:    ID,
		Fields: []data.Field{},
	}
}

func CreateMonEvent(ID string, value bool) data.MudEvent {
	return data.MudEvent{
		Table: "Mon",
		Key:   ID,
		Fields: []data.Field{
			{Key: "value", Data: data.BoolField{Data: value}},
		},
	}
}

func DeleteMonEvent(ID string) data.MudEvent {
	return data.MudEvent{
		Table:  "Mon",
		Key:    ID,
		Fields: []data.Field{},
	}
}

func CreateMonSpecieEvent(ID string, value int64) data.MudEvent {
	return data.MudEvent{
		Table: "MonSpecie",
		Key:   ID,
		Fields: []data.Field{
			{Key: "value", Data: data.UintField{Data: *big.NewInt(value)}},
		},
	}
}

func DeleteMonSpecieEvent(ID string) data.MudEvent {
	return data.MudEvent{
		Table:  "MonSpecie",
		Key:    ID,
		Fields: []data.Field{},
	}
}

func CreateMonHpEvent(ID string, value int64) data.MudEvent {
	return data.MudEvent{
		Table: "MonHp",
		Key:   ID,
		Fields: []data.Field{
			{Key: "value", Data: data.UintField{Data: *big.NewInt(value)}},
		},
	}
}

func DeleteMonHpEvent(ID string) data.MudEvent {
	return data.MudEvent{
		Table:  "MonHp",
		Key:    ID,
		Fields: []data.Field{},
	}
}

func CreateInventoryFirstMonEvent(ID string, value int64) data.MudEvent {
	return data.MudEvent{
		Table: "InventoryFirstMon",
		Key:   ID,
		Fields: []data.Field{
			{Key: "value", Data: data.UintField{Data: *big.NewInt(value)}},
		},
	}
}

func DeleteInventoryFirstMonEvent(ID string) data.MudEvent {
	return data.MudEvent{
		Table:  "InventoryFirstMon",
		Key:    ID,
		Fields: []data.Field{},
	}
}

func CreateInventorySecondMonEvent(ID string, value int64) data.MudEvent {
	return data.MudEvent{
		Table: "InventorySecondMon",
		Key:   ID,
		Fields: []data.Field{
			{Key: "value", Data: data.UintField{Data: *big.NewInt(value)}},
		},
	}
}

func DeleteInventorySecondMonEvent(ID string) data.MudEvent {
	return data.MudEvent{
		Table:  "InventorySecondMon",
		Key:    ID,
		Fields: []data.Field{},
	}
}

func CreateInventoryThirdMonEvent(ID string, value int64) data.MudEvent {
	return data.MudEvent{
		Table: "InventoryThirdMon",
		Key:   ID,
		Fields: []data.Field{
			{Key: "value", Data: data.UintField{Data: *big.NewInt(value)}},
		},
	}
}

func DeleteInventoryThirdMonEvent(ID string) data.MudEvent {
	return data.MudEvent{
		Table:  "InventoryThirdMon",
		Key:    ID,
		Fields: []data.Field{},
	}
}
