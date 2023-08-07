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
	field0, err := p.blockchainConnection.getPlayer(key)
	if err != nil {
		panic("value not found")
	}
	return field0
}

func (p *Prediction) PlayerSet(ID string, value bool) {
	p.Events = append(p.Events, CreatePlayerEvent(ID, value))
}

func (p *Prediction) PlayerDeleteRecord(ID string) {
	p.Events = append(p.Events, DeletePlayerEvent(ID))
}

func (p *Prediction) PlayerKeys(value bool) []string {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	return p.blockchainConnection.getRowsPlayer(value)
}

func (p *Prediction) StatusGet(key string) int64 {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	field0, err := p.blockchainConnection.getStatus(key)
	if err != nil {
		panic("value not found")
	}
	return field0
}

func (p *Prediction) StatusSet(ID string, value int64) {
	p.Events = append(p.Events, CreateStatusEvent(ID, value))
}

func (p *Prediction) StatusDeleteRecord(ID string) {
	p.Events = append(p.Events, DeleteStatusEvent(ID))
}

func (p *Prediction) StatusKeys(value int64) []string {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	return p.blockchainConnection.getRowsStatus(value)
}

func (p *Prediction) PositionGet(key string) (int64, int64) {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	field0, field1, err := p.blockchainConnection.getPosition(key)
	if err != nil {
		panic("value not found")
	}
	return field0, field1
}

func (p *Prediction) PositionSet(ID string, x int64, y int64) {
	p.Events = append(p.Events, CreatePositionEvent(ID, x, y))
}

func (p *Prediction) PositionDeleteRecord(ID string) {
	p.Events = append(p.Events, DeletePositionEvent(ID))
}

func (p *Prediction) PositionKeys(x int64, y int64) []string {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	return p.blockchainConnection.getRowsPosition(x, y)
}

func (p *Prediction) MatchGet(key string) bool {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	field0, err := p.blockchainConnection.getMatch(key)
	if err != nil {
		panic("value not found")
	}
	return field0
}

func (p *Prediction) MatchSet(ID string, value bool) {
	p.Events = append(p.Events, CreateMatchEvent(ID, value))
}

func (p *Prediction) MatchDeleteRecord(ID string) {
	p.Events = append(p.Events, DeleteMatchEvent(ID))
}

func (p *Prediction) MatchKeys(value bool) []string {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	return p.blockchainConnection.getRowsMatch(value)
}

func (p *Prediction) PlayerOneGet(key string) string {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	field0, err := p.blockchainConnection.getPlayerOne(key)
	if err != nil {
		panic("value not found")
	}
	return field0
}

func (p *Prediction) PlayerOneSet(ID string, value string) {
	p.Events = append(p.Events, CreatePlayerOneEvent(ID, value))
}

func (p *Prediction) PlayerOneDeleteRecord(ID string) {
	p.Events = append(p.Events, DeletePlayerOneEvent(ID))
}

func (p *Prediction) PlayerOneKeys(value string) []string {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	return p.blockchainConnection.getRowsPlayerOne(value)
}

func (p *Prediction) PlayerTwoGet(key string) string {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	field0, err := p.blockchainConnection.getPlayerTwo(key)
	if err != nil {
		panic("value not found")
	}
	return field0
}

func (p *Prediction) PlayerTwoSet(ID string, value string) {
	p.Events = append(p.Events, CreatePlayerTwoEvent(ID, value))
}

func (p *Prediction) PlayerTwoDeleteRecord(ID string) {
	p.Events = append(p.Events, DeletePlayerTwoEvent(ID))
}

func (p *Prediction) PlayerTwoKeys(value string) []string {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	return p.blockchainConnection.getRowsPlayerTwo(value)
}

func (p *Prediction) PlayerOneCurrentMonGet(key string) string {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	field0, err := p.blockchainConnection.getPlayerOneCurrentMon(key)
	if err != nil {
		panic("value not found")
	}
	return field0
}

func (p *Prediction) PlayerOneCurrentMonSet(ID string, value string) {
	p.Events = append(p.Events, CreatePlayerOneCurrentMonEvent(ID, value))
}

func (p *Prediction) PlayerOneCurrentMonDeleteRecord(ID string) {
	p.Events = append(p.Events, DeletePlayerOneCurrentMonEvent(ID))
}

func (p *Prediction) PlayerOneCurrentMonKeys(value string) []string {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	return p.blockchainConnection.getRowsPlayerOneCurrentMon(value)
}

func (p *Prediction) PlayerTwoCurrentMonGet(key string) string {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	field0, err := p.blockchainConnection.getPlayerTwoCurrentMon(key)
	if err != nil {
		panic("value not found")
	}
	return field0
}

func (p *Prediction) PlayerTwoCurrentMonSet(ID string, value string) {
	p.Events = append(p.Events, CreatePlayerTwoCurrentMonEvent(ID, value))
}

func (p *Prediction) PlayerTwoCurrentMonDeleteRecord(ID string) {
	p.Events = append(p.Events, DeletePlayerTwoCurrentMonEvent(ID))
}

func (p *Prediction) PlayerTwoCurrentMonKeys(value string) []string {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	return p.blockchainConnection.getRowsPlayerTwoCurrentMon(value)
}

func (p *Prediction) PlayerFirstMonGet(key string) string {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	field0, err := p.blockchainConnection.getPlayerFirstMon(key)
	if err != nil {
		panic("value not found")
	}
	return field0
}

func (p *Prediction) PlayerFirstMonSet(ID string, value string) {
	p.Events = append(p.Events, CreatePlayerFirstMonEvent(ID, value))
}

func (p *Prediction) PlayerFirstMonDeleteRecord(ID string) {
	p.Events = append(p.Events, DeletePlayerFirstMonEvent(ID))
}

func (p *Prediction) PlayerFirstMonKeys(value string) []string {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	return p.blockchainConnection.getRowsPlayerFirstMon(value)
}

func (p *Prediction) PlayerSecondMonGet(key string) string {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	field0, err := p.blockchainConnection.getPlayerSecondMon(key)
	if err != nil {
		panic("value not found")
	}
	return field0
}

func (p *Prediction) PlayerSecondMonSet(ID string, value string) {
	p.Events = append(p.Events, CreatePlayerSecondMonEvent(ID, value))
}

func (p *Prediction) PlayerSecondMonDeleteRecord(ID string) {
	p.Events = append(p.Events, DeletePlayerSecondMonEvent(ID))
}

func (p *Prediction) PlayerSecondMonKeys(value string) []string {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	return p.blockchainConnection.getRowsPlayerSecondMon(value)
}

func (p *Prediction) PlayerThirdMonGet(key string) string {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	field0, err := p.blockchainConnection.getPlayerThirdMon(key)
	if err != nil {
		panic("value not found")
	}
	return field0
}

func (p *Prediction) PlayerThirdMonSet(ID string, value string) {
	p.Events = append(p.Events, CreatePlayerThirdMonEvent(ID, value))
}

func (p *Prediction) PlayerThirdMonDeleteRecord(ID string) {
	p.Events = append(p.Events, DeletePlayerThirdMonEvent(ID))
}

func (p *Prediction) PlayerThirdMonKeys(value string) []string {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	return p.blockchainConnection.getRowsPlayerThirdMon(value)
}

func (p *Prediction) MonGet(key string) bool {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	field0, err := p.blockchainConnection.getMon(key)
	if err != nil {
		panic("value not found")
	}
	return field0
}

func (p *Prediction) MonSet(ID string, value bool) {
	p.Events = append(p.Events, CreateMonEvent(ID, value))
}

func (p *Prediction) MonDeleteRecord(ID string) {
	p.Events = append(p.Events, DeleteMonEvent(ID))
}

func (p *Prediction) MonKeys(value bool) []string {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	return p.blockchainConnection.getRowsMon(value)
}

func (p *Prediction) MonSpecieGet(key string) int64 {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	field0, err := p.blockchainConnection.getMonSpecie(key)
	if err != nil {
		panic("value not found")
	}
	return field0
}

func (p *Prediction) MonSpecieSet(ID string, value int64) {
	p.Events = append(p.Events, CreateMonSpecieEvent(ID, value))
}

func (p *Prediction) MonSpecieDeleteRecord(ID string) {
	p.Events = append(p.Events, DeleteMonSpecieEvent(ID))
}

func (p *Prediction) MonSpecieKeys(value int64) []string {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	return p.blockchainConnection.getRowsMonSpecie(value)
}

func (p *Prediction) MonHpGet(key string) int64 {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	field0, err := p.blockchainConnection.getMonHp(key)
	if err != nil {
		panic("value not found")
	}
	return field0
}

func (p *Prediction) MonHpSet(ID string, value int64) {
	p.Events = append(p.Events, CreateMonHpEvent(ID, value))
}

func (p *Prediction) MonHpDeleteRecord(ID string) {
	p.Events = append(p.Events, DeleteMonHpEvent(ID))
}

func (p *Prediction) MonHpKeys(value int64) []string {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	return p.blockchainConnection.getRowsMonHp(value)
}

func (p *Prediction) InventoryFirstMonGet(key string) int64 {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	field0, err := p.blockchainConnection.getInventoryFirstMon(key)
	if err != nil {
		panic("value not found")
	}
	return field0
}

func (p *Prediction) InventoryFirstMonSet(ID string, value int64) {
	p.Events = append(p.Events, CreateInventoryFirstMonEvent(ID, value))
}

func (p *Prediction) InventoryFirstMonDeleteRecord(ID string) {
	p.Events = append(p.Events, DeleteInventoryFirstMonEvent(ID))
}

func (p *Prediction) InventoryFirstMonKeys(value int64) []string {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	return p.blockchainConnection.getRowsInventoryFirstMon(value)
}

func (p *Prediction) InventorySecondMonGet(key string) int64 {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	field0, err := p.blockchainConnection.getInventorySecondMon(key)
	if err != nil {
		panic("value not found")
	}
	return field0
}

func (p *Prediction) InventorySecondMonSet(ID string, value int64) {
	p.Events = append(p.Events, CreateInventorySecondMonEvent(ID, value))
}

func (p *Prediction) InventorySecondMonDeleteRecord(ID string) {
	p.Events = append(p.Events, DeleteInventorySecondMonEvent(ID))
}

func (p *Prediction) InventorySecondMonKeys(value int64) []string {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	return p.blockchainConnection.getRowsInventorySecondMon(value)
}

func (p *Prediction) InventoryThirdMonGet(key string) int64 {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	field0, err := p.blockchainConnection.getInventoryThirdMon(key)
	if err != nil {
		panic("value not found")
	}
	return field0
}

func (p *Prediction) InventoryThirdMonSet(ID string, value int64) {
	p.Events = append(p.Events, CreateInventoryThirdMonEvent(ID, value))
}

func (p *Prediction) InventoryThirdMonDeleteRecord(ID string) {
	p.Events = append(p.Events, DeleteInventoryThirdMonEvent(ID))
}

func (p *Prediction) InventoryThirdMonKeys(value int64) []string {
	if !p.blockchainConnection.active {
		panic("game object is not active")
	}
	return p.blockchainConnection.getRowsInventoryThirdMon(value)
}
