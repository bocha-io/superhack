package garnethelpers

import (
	"strings"

	"github.com/bocha-io/garnet/x/indexer/data"
)

type Prediction struct {
	Events               []data.MudEvent
	blockchainConnection GameObject
}

func NewPrediction(db *data.Database) *Prediction {
	return &Prediction{
		Events:               []data.MudEvent{},
		blockchainConnection: *NewGameObject(db),
	}
}

func (Prediction) addressToEntityKey(address string) string {
	return strings.Replace(address, "0x", "0x000000000000000000000000", 1)
}

func (p *Prediction) PlayerGet(key string) bool {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	field0, err := p.blockchainConnection.GetPlayer(key)
	if err != nil {
		panic("value not found")
	}
	return field0
}

func (p *Prediction) PlayerSet(ID string, value bool) {
	p.Events = append(p.Events, CreatePlayerEvent(ID, value))
}

func (p *Prediction) PlayerDeleterecord(ID string) {
	p.Events = append(p.Events, DeletePlayerEvent(ID))
}

func (p *Prediction) PlayerKeys(value bool) []string {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	return p.blockchainConnection.GetRowsPlayer(value)
}

func (p *Prediction) StatusGet(key string) int64 {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	field0, err := p.blockchainConnection.GetStatus(key)
	if err != nil {
		panic("value not found")
	}
	return field0
}

func (p *Prediction) StatusSet(ID string, value int64) {
	p.Events = append(p.Events, CreateStatusEvent(ID, value))
}

func (p *Prediction) StatusDeleterecord(ID string) {
	p.Events = append(p.Events, DeleteStatusEvent(ID))
}

func (p *Prediction) StatusKeys(value int64) []string {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	return p.blockchainConnection.GetRowsStatus(value)
}

func (p *Prediction) PositionGet(key string) (int64, int64) {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	field0, field1, err := p.blockchainConnection.GetPosition(key)
	if err != nil {
		panic("value not found")
	}
	return field0, field1
}

func (p *Prediction) PositionSet(ID string, x int64, y int64) {
	p.Events = append(p.Events, CreatePositionEvent(ID, x, y))
}

func (p *Prediction) PositionDeleterecord(ID string) {
	p.Events = append(p.Events, DeletePositionEvent(ID))
}

func (p *Prediction) PositionKeys(x int64, y int64) []string {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	return p.blockchainConnection.GetRowsPosition(x, y)
}

func (p *Prediction) MatchGet(key string) bool {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	field0, err := p.blockchainConnection.GetMatch(key)
	if err != nil {
		panic("value not found")
	}
	return field0
}

func (p *Prediction) MatchSet(ID string, value bool) {
	p.Events = append(p.Events, CreateMatchEvent(ID, value))
}

func (p *Prediction) MatchDeleterecord(ID string) {
	p.Events = append(p.Events, DeleteMatchEvent(ID))
}

func (p *Prediction) MatchKeys(value bool) []string {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	return p.blockchainConnection.GetRowsMatch(value)
}

func (p *Prediction) MatchResultGet(key string) (string, string) {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	field0, field1, err := p.blockchainConnection.GetMatchResult(key)
	if err != nil {
		panic("value not found")
	}
	return field0, field1
}

func (p *Prediction) MatchResultSet(ID string, winner string, loser string) {
	p.Events = append(p.Events, CreateMatchResultEvent(ID, winner, loser))
}

func (p *Prediction) MatchResultDeleterecord(ID string) {
	p.Events = append(p.Events, DeleteMatchResultEvent(ID))
}

func (p *Prediction) MatchResultKeys(winner string, loser string) []string {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	return p.blockchainConnection.GetRowsMatchResult(winner, loser)
}

func (p *Prediction) PlayerOneGet(key string) string {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	field0, err := p.blockchainConnection.GetPlayerOne(key)
	if err != nil {
		panic("value not found")
	}
	return field0
}

func (p *Prediction) PlayerOneSet(ID string, value string) {
	p.Events = append(p.Events, CreatePlayerOneEvent(ID, value))
}

func (p *Prediction) PlayerOneDeleterecord(ID string) {
	p.Events = append(p.Events, DeletePlayerOneEvent(ID))
}

func (p *Prediction) PlayerOneKeys(value string) []string {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	return p.blockchainConnection.GetRowsPlayerOne(value)
}

func (p *Prediction) PlayerTwoGet(key string) string {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	field0, err := p.blockchainConnection.GetPlayerTwo(key)
	if err != nil {
		panic("value not found")
	}
	return field0
}

func (p *Prediction) PlayerTwoSet(ID string, value string) {
	p.Events = append(p.Events, CreatePlayerTwoEvent(ID, value))
}

func (p *Prediction) PlayerTwoDeleterecord(ID string) {
	p.Events = append(p.Events, DeletePlayerTwoEvent(ID))
}

func (p *Prediction) PlayerTwoKeys(value string) []string {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	return p.blockchainConnection.GetRowsPlayerTwo(value)
}

func (p *Prediction) PlayerOneCurrentMonGet(key string) string {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	field0, err := p.blockchainConnection.GetPlayerOneCurrentMon(key)
	if err != nil {
		panic("value not found")
	}
	return field0
}

func (p *Prediction) PlayerOneCurrentMonSet(ID string, value string) {
	p.Events = append(p.Events, CreatePlayerOneCurrentMonEvent(ID, value))
}

func (p *Prediction) PlayerOneCurrentMonDeleterecord(ID string) {
	p.Events = append(p.Events, DeletePlayerOneCurrentMonEvent(ID))
}

func (p *Prediction) PlayerOneCurrentMonKeys(value string) []string {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	return p.blockchainConnection.GetRowsPlayerOneCurrentMon(value)
}

func (p *Prediction) PlayerTwoCurrentMonGet(key string) string {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	field0, err := p.blockchainConnection.GetPlayerTwoCurrentMon(key)
	if err != nil {
		panic("value not found")
	}
	return field0
}

func (p *Prediction) PlayerTwoCurrentMonSet(ID string, value string) {
	p.Events = append(p.Events, CreatePlayerTwoCurrentMonEvent(ID, value))
}

func (p *Prediction) PlayerTwoCurrentMonDeleterecord(ID string) {
	p.Events = append(p.Events, DeletePlayerTwoCurrentMonEvent(ID))
}

func (p *Prediction) PlayerTwoCurrentMonKeys(value string) []string {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	return p.blockchainConnection.GetRowsPlayerTwoCurrentMon(value)
}

func (p *Prediction) PlayerFirstMonGet(key string) string {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	field0, err := p.blockchainConnection.GetPlayerFirstMon(key)
	if err != nil {
		panic("value not found")
	}
	return field0
}

func (p *Prediction) PlayerFirstMonSet(ID string, value string) {
	p.Events = append(p.Events, CreatePlayerFirstMonEvent(ID, value))
}

func (p *Prediction) PlayerFirstMonDeleterecord(ID string) {
	p.Events = append(p.Events, DeletePlayerFirstMonEvent(ID))
}

func (p *Prediction) PlayerFirstMonKeys(value string) []string {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	return p.blockchainConnection.GetRowsPlayerFirstMon(value)
}

func (p *Prediction) PlayerSecondMonGet(key string) string {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	field0, err := p.blockchainConnection.GetPlayerSecondMon(key)
	if err != nil {
		panic("value not found")
	}
	return field0
}

func (p *Prediction) PlayerSecondMonSet(ID string, value string) {
	p.Events = append(p.Events, CreatePlayerSecondMonEvent(ID, value))
}

func (p *Prediction) PlayerSecondMonDeleterecord(ID string) {
	p.Events = append(p.Events, DeletePlayerSecondMonEvent(ID))
}

func (p *Prediction) PlayerSecondMonKeys(value string) []string {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	return p.blockchainConnection.GetRowsPlayerSecondMon(value)
}

func (p *Prediction) PlayerThirdMonGet(key string) string {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	field0, err := p.blockchainConnection.GetPlayerThirdMon(key)
	if err != nil {
		panic("value not found")
	}
	return field0
}

func (p *Prediction) PlayerThirdMonSet(ID string, value string) {
	p.Events = append(p.Events, CreatePlayerThirdMonEvent(ID, value))
}

func (p *Prediction) PlayerThirdMonDeleterecord(ID string) {
	p.Events = append(p.Events, DeletePlayerThirdMonEvent(ID))
}

func (p *Prediction) PlayerThirdMonKeys(value string) []string {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	return p.blockchainConnection.GetRowsPlayerThirdMon(value)
}

func (p *Prediction) MonGet(key string) bool {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	field0, err := p.blockchainConnection.GetMon(key)
	if err != nil {
		panic("value not found")
	}
	return field0
}

func (p *Prediction) MonSet(ID string, value bool) {
	p.Events = append(p.Events, CreateMonEvent(ID, value))
}

func (p *Prediction) MonDeleterecord(ID string) {
	p.Events = append(p.Events, DeleteMonEvent(ID))
}

func (p *Prediction) MonKeys(value bool) []string {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	return p.blockchainConnection.GetRowsMon(value)
}

func (p *Prediction) MonSpecieGet(key string) int64 {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	field0, err := p.blockchainConnection.GetMonSpecie(key)
	if err != nil {
		panic("value not found")
	}
	return field0
}

func (p *Prediction) MonSpecieSet(ID string, value int64) {
	p.Events = append(p.Events, CreateMonSpecieEvent(ID, value))
}

func (p *Prediction) MonSpecieDeleterecord(ID string) {
	p.Events = append(p.Events, DeleteMonSpecieEvent(ID))
}

func (p *Prediction) MonSpecieKeys(value int64) []string {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	return p.blockchainConnection.GetRowsMonSpecie(value)
}

func (p *Prediction) MonHpGet(key string) int64 {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	field0, err := p.blockchainConnection.GetMonHp(key)
	if err != nil {
		panic("value not found")
	}
	return field0
}

func (p *Prediction) MonHpSet(ID string, value int64) {
	p.Events = append(p.Events, CreateMonHpEvent(ID, value))
}

func (p *Prediction) MonHpDeleterecord(ID string) {
	p.Events = append(p.Events, DeleteMonHpEvent(ID))
}

func (p *Prediction) MonHpKeys(value int64) []string {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	return p.blockchainConnection.GetRowsMonHp(value)
}

func (p *Prediction) InventoryFirstMonGet(key string) int64 {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	field0, err := p.blockchainConnection.GetInventoryFirstMon(key)
	if err != nil {
		panic("value not found")
	}
	return field0
}

func (p *Prediction) InventoryFirstMonSet(ID string, value int64) {
	p.Events = append(p.Events, CreateInventoryFirstMonEvent(ID, value))
}

func (p *Prediction) InventoryFirstMonDeleterecord(ID string) {
	p.Events = append(p.Events, DeleteInventoryFirstMonEvent(ID))
}

func (p *Prediction) InventoryFirstMonKeys(value int64) []string {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	return p.blockchainConnection.GetRowsInventoryFirstMon(value)
}

func (p *Prediction) InventorySecondMonGet(key string) int64 {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	field0, err := p.blockchainConnection.GetInventorySecondMon(key)
	if err != nil {
		panic("value not found")
	}
	return field0
}

func (p *Prediction) InventorySecondMonSet(ID string, value int64) {
	p.Events = append(p.Events, CreateInventorySecondMonEvent(ID, value))
}

func (p *Prediction) InventorySecondMonDeleterecord(ID string) {
	p.Events = append(p.Events, DeleteInventorySecondMonEvent(ID))
}

func (p *Prediction) InventorySecondMonKeys(value int64) []string {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	return p.blockchainConnection.GetRowsInventorySecondMon(value)
}

func (p *Prediction) InventoryThirdMonGet(key string) int64 {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	field0, err := p.blockchainConnection.GetInventoryThirdMon(key)
	if err != nil {
		panic("value not found")
	}
	return field0
}

func (p *Prediction) InventoryThirdMonSet(ID string, value int64) {
	p.Events = append(p.Events, CreateInventoryThirdMonEvent(ID, value))
}

func (p *Prediction) InventoryThirdMonDeleterecord(ID string) {
	p.Events = append(p.Events, DeleteInventoryThirdMonEvent(ID))
}

func (p *Prediction) InventoryThirdMonKeys(value int64) []string {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	return p.blockchainConnection.GetRowsInventoryThirdMon(value)
}
