package garnethelpers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bocha-io/garnet/x/indexer/data"
)

func (g *GameObject) ProcessFieldsPlayer(fields []data.Field) (bool, error) {
	if len(fields) != 1 {
		return false, fmt.Errorf("invalid amount of fields")
	}

	field0 := fields[0].Data.String() == "true"
	return field0, nil
}

func (g *GameObject) GetPlayer(key string) (bool, error) {
	fields, err := data.GetRowFieldsUsingString(g.db, g.world, key, "Player")
	if err != nil {
		return false, err
	}
	return g.ProcessFieldsPlayer(fields)
}

func (g GameObject) GetAllRowsPlayer() map[string][]data.Field {
	table := g.world.GetTableByName("Player")
	return g.db.GetRows(table)
}

func (g GameObject) GetRowsPlayer(arg0 bool) []string {
	rows := g.GetAllRowsPlayer()
	for k, fields := range rows {
		if len(fields) != 1 {
			continue
		}

		field0 := fields[0].Data.String() == "true"
		if field0 != arg0 {
			continue
		}
		return []string{k}
	}
	return []string{}
}

func (g *GameObject) ProcessFieldsStatus(fields []data.Field) (int64, error) {
	if len(fields) != 1 {
		return 0, fmt.Errorf("invalid amount of fields")
	}

	field0, err := strconv.ParseInt(fields[0].Data.String(), 10, 32)
	if err != nil {
		return 0, err
	}
	return field0, nil
}

func (g *GameObject) GetStatus(key string) (int64, error) {
	fields, err := data.GetRowFieldsUsingString(g.db, g.world, key, "Status")
	if err != nil {
		return 0, err
	}
	return g.ProcessFieldsStatus(fields)
}

func (g GameObject) GetAllRowsStatus() map[string][]data.Field {
	table := g.world.GetTableByName("Status")
	return g.db.GetRows(table)
}

func (g GameObject) GetRowsStatus(arg0 int64) []string {
	rows := g.GetAllRowsStatus()
	for k, fields := range rows {
		if len(fields) != 1 {
			continue
		}

		field0, err := strconv.ParseInt(fields[0].Data.String(), 10, 32)
		if err != nil {
			continue
		}
		if field0 != arg0 {
			continue
		}
		return []string{k}
	}
	return []string{}
}

func (g *GameObject) ProcessFieldsPosition(fields []data.Field) (int64, int64, error) {
	if len(fields) != 2 {
		return 0, 0, fmt.Errorf("invalid amount of fields")
	}

	field0, err := strconv.ParseInt(fields[0].Data.String(), 10, 32)
	if err != nil {
		return 0, 0, err
	}
	field1, err := strconv.ParseInt(fields[1].Data.String(), 10, 32)
	if err != nil {
		return 0, 0, err
	}
	return field0, field1, nil
}

func (g *GameObject) GetPosition(key string) (int64, int64, error) {
	fields, err := data.GetRowFieldsUsingString(g.db, g.world, key, "Position")
	if err != nil {
		return 0, 0, err
	}
	return g.ProcessFieldsPosition(fields)
}

func (g GameObject) GetAllRowsPosition() map[string][]data.Field {
	table := g.world.GetTableByName("Position")
	return g.db.GetRows(table)
}

func (g GameObject) GetRowsPosition(arg0 int64, arg1 int64) []string {
	rows := g.GetAllRowsPosition()
	for k, fields := range rows {
		if len(fields) != 2 {
			continue
		}

		field0, err := strconv.ParseInt(fields[0].Data.String(), 10, 32)
		if err != nil {
			continue
		}
		if field0 != arg0 {
			continue
		}
		field1, err := strconv.ParseInt(fields[1].Data.String(), 10, 32)
		if err != nil {
			continue
		}
		if field1 != arg1 {
			continue
		}
		return []string{k}
	}
	return []string{}
}

func (g *GameObject) ProcessFieldsMatch(fields []data.Field) (bool, error) {
	if len(fields) != 1 {
		return false, fmt.Errorf("invalid amount of fields")
	}

	field0 := fields[0].Data.String() == "true"
	return field0, nil
}

func (g *GameObject) GetMatch(key string) (bool, error) {
	fields, err := data.GetRowFieldsUsingString(g.db, g.world, key, "Match")
	if err != nil {
		return false, err
	}
	return g.ProcessFieldsMatch(fields)
}

func (g GameObject) GetAllRowsMatch() map[string][]data.Field {
	table := g.world.GetTableByName("Match")
	return g.db.GetRows(table)
}

func (g GameObject) GetRowsMatch(arg0 bool) []string {
	rows := g.GetAllRowsMatch()
	for k, fields := range rows {
		if len(fields) != 1 {
			continue
		}

		field0 := fields[0].Data.String() == "true"
		if field0 != arg0 {
			continue
		}
		return []string{k}
	}
	return []string{}
}

func (g *GameObject) ProcessFieldsMatchResult(fields []data.Field) (string, string, error) {
	if len(fields) != 2 {
		return "", "", fmt.Errorf("invalid amount of fields")
	}

	field0 := strings.ReplaceAll(fields[0].Data.String(), "\"", "")
	field1 := strings.ReplaceAll(fields[1].Data.String(), "\"", "")
	return field0, field1, nil
}

func (g *GameObject) GetMatchResult(key string) (string, string, error) {
	fields, err := data.GetRowFieldsUsingString(g.db, g.world, key, "MatchResult")
	if err != nil {
		return "", "", err
	}
	return g.ProcessFieldsMatchResult(fields)
}

func (g GameObject) GetAllRowsMatchResult() map[string][]data.Field {
	table := g.world.GetTableByName("MatchResult")
	return g.db.GetRows(table)
}

func (g GameObject) GetRowsMatchResult(arg0 string, arg1 string) []string {
	rows := g.GetAllRowsMatchResult()
	for k, fields := range rows {
		if len(fields) != 2 {
			continue
		}

		field0 := strings.ReplaceAll(fields[0].Data.String(), "\"", "")
		if field0 != arg0 {
			continue
		}
		field1 := strings.ReplaceAll(fields[1].Data.String(), "\"", "")
		if field1 != arg1 {
			continue
		}
		return []string{k}
	}
	return []string{}
}

func (g *GameObject) ProcessFieldsPlayerOne(fields []data.Field) (string, error) {
	if len(fields) != 1 {
		return "", fmt.Errorf("invalid amount of fields")
	}

	field0 := strings.ReplaceAll(fields[0].Data.String(), "\"", "")
	return field0, nil
}

func (g *GameObject) GetPlayerOne(key string) (string, error) {
	fields, err := data.GetRowFieldsUsingString(g.db, g.world, key, "PlayerOne")
	if err != nil {
		return "", err
	}
	return g.ProcessFieldsPlayerOne(fields)
}

func (g GameObject) GetAllRowsPlayerOne() map[string][]data.Field {
	table := g.world.GetTableByName("PlayerOne")
	return g.db.GetRows(table)
}

func (g GameObject) GetRowsPlayerOne(arg0 string) []string {
	rows := g.GetAllRowsPlayerOne()
	for k, fields := range rows {
		if len(fields) != 1 {
			continue
		}

		field0 := strings.ReplaceAll(fields[0].Data.String(), "\"", "")
		if field0 != arg0 {
			continue
		}
		return []string{k}
	}
	return []string{}
}

func (g *GameObject) ProcessFieldsPlayerTwo(fields []data.Field) (string, error) {
	if len(fields) != 1 {
		return "", fmt.Errorf("invalid amount of fields")
	}

	field0 := strings.ReplaceAll(fields[0].Data.String(), "\"", "")
	return field0, nil
}

func (g *GameObject) GetPlayerTwo(key string) (string, error) {
	fields, err := data.GetRowFieldsUsingString(g.db, g.world, key, "PlayerTwo")
	if err != nil {
		return "", err
	}
	return g.ProcessFieldsPlayerTwo(fields)
}

func (g GameObject) GetAllRowsPlayerTwo() map[string][]data.Field {
	table := g.world.GetTableByName("PlayerTwo")
	return g.db.GetRows(table)
}

func (g GameObject) GetRowsPlayerTwo(arg0 string) []string {
	rows := g.GetAllRowsPlayerTwo()
	for k, fields := range rows {
		if len(fields) != 1 {
			continue
		}

		field0 := strings.ReplaceAll(fields[0].Data.String(), "\"", "")
		if field0 != arg0 {
			continue
		}
		return []string{k}
	}
	return []string{}
}

func (g *GameObject) ProcessFieldsPlayerOneCurrentMon(fields []data.Field) (string, error) {
	if len(fields) != 1 {
		return "", fmt.Errorf("invalid amount of fields")
	}

	field0 := strings.ReplaceAll(fields[0].Data.String(), "\"", "")
	return field0, nil
}

func (g *GameObject) GetPlayerOneCurrentMon(key string) (string, error) {
	fields, err := data.GetRowFieldsUsingString(g.db, g.world, key, "PlayerOneCurrentMon")
	if err != nil {
		return "", err
	}
	return g.ProcessFieldsPlayerOneCurrentMon(fields)
}

func (g GameObject) GetAllRowsPlayerOneCurrentMon() map[string][]data.Field {
	table := g.world.GetTableByName("PlayerOneCurrentMon")
	return g.db.GetRows(table)
}

func (g GameObject) GetRowsPlayerOneCurrentMon(arg0 string) []string {
	rows := g.GetAllRowsPlayerOneCurrentMon()
	for k, fields := range rows {
		if len(fields) != 1 {
			continue
		}

		field0 := strings.ReplaceAll(fields[0].Data.String(), "\"", "")
		if field0 != arg0 {
			continue
		}
		return []string{k}
	}
	return []string{}
}

func (g *GameObject) ProcessFieldsPlayerTwoCurrentMon(fields []data.Field) (string, error) {
	if len(fields) != 1 {
		return "", fmt.Errorf("invalid amount of fields")
	}

	field0 := strings.ReplaceAll(fields[0].Data.String(), "\"", "")
	return field0, nil
}

func (g *GameObject) GetPlayerTwoCurrentMon(key string) (string, error) {
	fields, err := data.GetRowFieldsUsingString(g.db, g.world, key, "PlayerTwoCurrentMon")
	if err != nil {
		return "", err
	}
	return g.ProcessFieldsPlayerTwoCurrentMon(fields)
}

func (g GameObject) GetAllRowsPlayerTwoCurrentMon() map[string][]data.Field {
	table := g.world.GetTableByName("PlayerTwoCurrentMon")
	return g.db.GetRows(table)
}

func (g GameObject) GetRowsPlayerTwoCurrentMon(arg0 string) []string {
	rows := g.GetAllRowsPlayerTwoCurrentMon()
	for k, fields := range rows {
		if len(fields) != 1 {
			continue
		}

		field0 := strings.ReplaceAll(fields[0].Data.String(), "\"", "")
		if field0 != arg0 {
			continue
		}
		return []string{k}
	}
	return []string{}
}

func (g *GameObject) ProcessFieldsPlayerFirstMon(fields []data.Field) (string, error) {
	if len(fields) != 1 {
		return "", fmt.Errorf("invalid amount of fields")
	}

	field0 := strings.ReplaceAll(fields[0].Data.String(), "\"", "")
	return field0, nil
}

func (g *GameObject) GetPlayerFirstMon(key string) (string, error) {
	fields, err := data.GetRowFieldsUsingString(g.db, g.world, key, "PlayerFirstMon")
	if err != nil {
		return "", err
	}
	return g.ProcessFieldsPlayerFirstMon(fields)
}

func (g GameObject) GetAllRowsPlayerFirstMon() map[string][]data.Field {
	table := g.world.GetTableByName("PlayerFirstMon")
	return g.db.GetRows(table)
}

func (g GameObject) GetRowsPlayerFirstMon(arg0 string) []string {
	rows := g.GetAllRowsPlayerFirstMon()
	for k, fields := range rows {
		if len(fields) != 1 {
			continue
		}

		field0 := strings.ReplaceAll(fields[0].Data.String(), "\"", "")
		if field0 != arg0 {
			continue
		}
		return []string{k}
	}
	return []string{}
}

func (g *GameObject) ProcessFieldsPlayerSecondMon(fields []data.Field) (string, error) {
	if len(fields) != 1 {
		return "", fmt.Errorf("invalid amount of fields")
	}

	field0 := strings.ReplaceAll(fields[0].Data.String(), "\"", "")
	return field0, nil
}

func (g *GameObject) GetPlayerSecondMon(key string) (string, error) {
	fields, err := data.GetRowFieldsUsingString(g.db, g.world, key, "PlayerSecondMon")
	if err != nil {
		return "", err
	}
	return g.ProcessFieldsPlayerSecondMon(fields)
}

func (g GameObject) GetAllRowsPlayerSecondMon() map[string][]data.Field {
	table := g.world.GetTableByName("PlayerSecondMon")
	return g.db.GetRows(table)
}

func (g GameObject) GetRowsPlayerSecondMon(arg0 string) []string {
	rows := g.GetAllRowsPlayerSecondMon()
	for k, fields := range rows {
		if len(fields) != 1 {
			continue
		}

		field0 := strings.ReplaceAll(fields[0].Data.String(), "\"", "")
		if field0 != arg0 {
			continue
		}
		return []string{k}
	}
	return []string{}
}

func (g *GameObject) ProcessFieldsPlayerThirdMon(fields []data.Field) (string, error) {
	if len(fields) != 1 {
		return "", fmt.Errorf("invalid amount of fields")
	}

	field0 := strings.ReplaceAll(fields[0].Data.String(), "\"", "")
	return field0, nil
}

func (g *GameObject) GetPlayerThirdMon(key string) (string, error) {
	fields, err := data.GetRowFieldsUsingString(g.db, g.world, key, "PlayerThirdMon")
	if err != nil {
		return "", err
	}
	return g.ProcessFieldsPlayerThirdMon(fields)
}

func (g GameObject) GetAllRowsPlayerThirdMon() map[string][]data.Field {
	table := g.world.GetTableByName("PlayerThirdMon")
	return g.db.GetRows(table)
}

func (g GameObject) GetRowsPlayerThirdMon(arg0 string) []string {
	rows := g.GetAllRowsPlayerThirdMon()
	for k, fields := range rows {
		if len(fields) != 1 {
			continue
		}

		field0 := strings.ReplaceAll(fields[0].Data.String(), "\"", "")
		if field0 != arg0 {
			continue
		}
		return []string{k}
	}
	return []string{}
}

func (g *GameObject) ProcessFieldsMon(fields []data.Field) (bool, error) {
	if len(fields) != 1 {
		return false, fmt.Errorf("invalid amount of fields")
	}

	field0 := fields[0].Data.String() == "true"
	return field0, nil
}

func (g *GameObject) GetMon(key string) (bool, error) {
	fields, err := data.GetRowFieldsUsingString(g.db, g.world, key, "Mon")
	if err != nil {
		return false, err
	}
	return g.ProcessFieldsMon(fields)
}

func (g GameObject) GetAllRowsMon() map[string][]data.Field {
	table := g.world.GetTableByName("Mon")
	return g.db.GetRows(table)
}

func (g GameObject) GetRowsMon(arg0 bool) []string {
	rows := g.GetAllRowsMon()
	for k, fields := range rows {
		if len(fields) != 1 {
			continue
		}

		field0 := fields[0].Data.String() == "true"
		if field0 != arg0 {
			continue
		}
		return []string{k}
	}
	return []string{}
}

func (g *GameObject) ProcessFieldsMonSpecie(fields []data.Field) (int64, error) {
	if len(fields) != 1 {
		return 0, fmt.Errorf("invalid amount of fields")
	}

	field0, err := strconv.ParseInt(fields[0].Data.String(), 10, 32)
	if err != nil {
		return 0, err
	}
	return field0, nil
}

func (g *GameObject) GetMonSpecie(key string) (int64, error) {
	fields, err := data.GetRowFieldsUsingString(g.db, g.world, key, "MonSpecie")
	if err != nil {
		return 0, err
	}
	return g.ProcessFieldsMonSpecie(fields)
}

func (g GameObject) GetAllRowsMonSpecie() map[string][]data.Field {
	table := g.world.GetTableByName("MonSpecie")
	return g.db.GetRows(table)
}

func (g GameObject) GetRowsMonSpecie(arg0 int64) []string {
	rows := g.GetAllRowsMonSpecie()
	for k, fields := range rows {
		if len(fields) != 1 {
			continue
		}

		field0, err := strconv.ParseInt(fields[0].Data.String(), 10, 32)
		if err != nil {
			continue
		}
		if field0 != arg0 {
			continue
		}
		return []string{k}
	}
	return []string{}
}

func (g *GameObject) ProcessFieldsMonHp(fields []data.Field) (int64, error) {
	if len(fields) != 1 {
		return 0, fmt.Errorf("invalid amount of fields")
	}

	field0, err := strconv.ParseInt(fields[0].Data.String(), 10, 32)
	if err != nil {
		return 0, err
	}
	return field0, nil
}

func (g *GameObject) GetMonHp(key string) (int64, error) {
	fields, err := data.GetRowFieldsUsingString(g.db, g.world, key, "MonHp")
	if err != nil {
		return 0, err
	}
	return g.ProcessFieldsMonHp(fields)
}

func (g GameObject) GetAllRowsMonHp() map[string][]data.Field {
	table := g.world.GetTableByName("MonHp")
	return g.db.GetRows(table)
}

func (g GameObject) GetRowsMonHp(arg0 int64) []string {
	rows := g.GetAllRowsMonHp()
	for k, fields := range rows {
		if len(fields) != 1 {
			continue
		}

		field0, err := strconv.ParseInt(fields[0].Data.String(), 10, 32)
		if err != nil {
			continue
		}
		if field0 != arg0 {
			continue
		}
		return []string{k}
	}
	return []string{}
}

func (g *GameObject) ProcessFieldsInventoryFirstMon(fields []data.Field) (int64, error) {
	if len(fields) != 1 {
		return 0, fmt.Errorf("invalid amount of fields")
	}

	field0, err := strconv.ParseInt(fields[0].Data.String(), 10, 32)
	if err != nil {
		return 0, err
	}
	return field0, nil
}

func (g *GameObject) GetInventoryFirstMon(key string) (int64, error) {
	fields, err := data.GetRowFieldsUsingString(g.db, g.world, key, "InventoryFirstMon")
	if err != nil {
		return 0, err
	}
	return g.ProcessFieldsInventoryFirstMon(fields)
}

func (g GameObject) GetAllRowsInventoryFirstMon() map[string][]data.Field {
	table := g.world.GetTableByName("InventoryFirstMon")
	return g.db.GetRows(table)
}

func (g GameObject) GetRowsInventoryFirstMon(arg0 int64) []string {
	rows := g.GetAllRowsInventoryFirstMon()
	for k, fields := range rows {
		if len(fields) != 1 {
			continue
		}

		field0, err := strconv.ParseInt(fields[0].Data.String(), 10, 32)
		if err != nil {
			continue
		}
		if field0 != arg0 {
			continue
		}
		return []string{k}
	}
	return []string{}
}

func (g *GameObject) ProcessFieldsInventorySecondMon(fields []data.Field) (int64, error) {
	if len(fields) != 1 {
		return 0, fmt.Errorf("invalid amount of fields")
	}

	field0, err := strconv.ParseInt(fields[0].Data.String(), 10, 32)
	if err != nil {
		return 0, err
	}
	return field0, nil
}

func (g *GameObject) GetInventorySecondMon(key string) (int64, error) {
	fields, err := data.GetRowFieldsUsingString(g.db, g.world, key, "InventorySecondMon")
	if err != nil {
		return 0, err
	}
	return g.ProcessFieldsInventorySecondMon(fields)
}

func (g GameObject) GetAllRowsInventorySecondMon() map[string][]data.Field {
	table := g.world.GetTableByName("InventorySecondMon")
	return g.db.GetRows(table)
}

func (g GameObject) GetRowsInventorySecondMon(arg0 int64) []string {
	rows := g.GetAllRowsInventorySecondMon()
	for k, fields := range rows {
		if len(fields) != 1 {
			continue
		}

		field0, err := strconv.ParseInt(fields[0].Data.String(), 10, 32)
		if err != nil {
			continue
		}
		if field0 != arg0 {
			continue
		}
		return []string{k}
	}
	return []string{}
}

func (g *GameObject) ProcessFieldsInventoryThirdMon(fields []data.Field) (int64, error) {
	if len(fields) != 1 {
		return 0, fmt.Errorf("invalid amount of fields")
	}

	field0, err := strconv.ParseInt(fields[0].Data.String(), 10, 32)
	if err != nil {
		return 0, err
	}
	return field0, nil
}

func (g *GameObject) GetInventoryThirdMon(key string) (int64, error) {
	fields, err := data.GetRowFieldsUsingString(g.db, g.world, key, "InventoryThirdMon")
	if err != nil {
		return 0, err
	}
	return g.ProcessFieldsInventoryThirdMon(fields)
}

func (g GameObject) GetAllRowsInventoryThirdMon() map[string][]data.Field {
	table := g.world.GetTableByName("InventoryThirdMon")
	return g.db.GetRows(table)
}

func (g GameObject) GetRowsInventoryThirdMon(arg0 int64) []string {
	rows := g.GetAllRowsInventoryThirdMon()
	for k, fields := range rows {
		if len(fields) != 1 {
			continue
		}

		field0, err := strconv.ParseInt(fields[0].Data.String(), 10, 32)
		if err != nil {
			continue
		}
		if field0 != arg0 {
			continue
		}
		return []string{k}
	}
	return []string{}
}
