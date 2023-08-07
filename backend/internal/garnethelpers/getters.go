package garnethelpers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bocha-io/garnet/x/indexer/data"
)

func (g *GameObject) getPlayer(key string) (bool, error) {
	fields, err := data.GetRowFieldsUsingString(g.db, g.world, key, "Player")
	if err != nil {
		return false, err
	}

	if len(fields) != 1 {
		return false, fmt.Errorf("invalid amount of fields")
	}

	field0 := fields[0].Data.String() == "true"
	return field0, nil
}

func (g GameObject) getRowsPlayer(arg0 bool) []string {
	table := g.world.GetTableByName("Player")
	rows := g.db.GetRows(table)
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

func (g *GameObject) getStatus(key string) (int64, error) {
	fields, err := data.GetRowFieldsUsingString(g.db, g.world, key, "Status")
	if err != nil {
		return 0, err
	}

	if len(fields) != 1 {
		return 0, fmt.Errorf("invalid amount of fields")
	}

	field0, err := strconv.ParseInt(fields[0].Data.String(), 10, 32)
	if err != nil {
		return 0, err
	}
	return field0, nil
}

func (g GameObject) getRowsStatus(arg0 int64) []string {
	table := g.world.GetTableByName("Status")
	rows := g.db.GetRows(table)
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

func (g *GameObject) getPosition(key string) (int64, int64, error) {
	fields, err := data.GetRowFieldsUsingString(g.db, g.world, key, "Position")
	if err != nil {
		return 0, 0, err
	}

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

func (g GameObject) getRowsPosition(arg0 int64, arg1 int64) []string {
	table := g.world.GetTableByName("Position")
	rows := g.db.GetRows(table)
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

func (g *GameObject) getMatch(key string) (bool, error) {
	fields, err := data.GetRowFieldsUsingString(g.db, g.world, key, "Match")
	if err != nil {
		return false, err
	}

	if len(fields) != 1 {
		return false, fmt.Errorf("invalid amount of fields")
	}

	field0 := fields[0].Data.String() == "true"
	return field0, nil
}

func (g GameObject) getRowsMatch(arg0 bool) []string {
	table := g.world.GetTableByName("Match")
	rows := g.db.GetRows(table)
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

func (g *GameObject) getPlayerOne(key string) (string, error) {
	fields, err := data.GetRowFieldsUsingString(g.db, g.world, key, "PlayerOne")
	if err != nil {
		return "", err
	}

	if len(fields) != 1 {
		return "", fmt.Errorf("invalid amount of fields")
	}

	field0 := strings.ReplaceAll(fields[0].Data.String(), "\"", "")
	return field0, nil
}

func (g GameObject) getRowsPlayerOne(arg0 string) []string {
	table := g.world.GetTableByName("PlayerOne")
	rows := g.db.GetRows(table)
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

func (g *GameObject) getPlayerTwo(key string) (string, error) {
	fields, err := data.GetRowFieldsUsingString(g.db, g.world, key, "PlayerTwo")
	if err != nil {
		return "", err
	}

	if len(fields) != 1 {
		return "", fmt.Errorf("invalid amount of fields")
	}

	field0 := strings.ReplaceAll(fields[0].Data.String(), "\"", "")
	return field0, nil
}

func (g GameObject) getRowsPlayerTwo(arg0 string) []string {
	table := g.world.GetTableByName("PlayerTwo")
	rows := g.db.GetRows(table)
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

func (g *GameObject) getPlayerOneCurrentMon(key string) (string, error) {
	fields, err := data.GetRowFieldsUsingString(g.db, g.world, key, "PlayerOneCurrentMon")
	if err != nil {
		return "", err
	}

	if len(fields) != 1 {
		return "", fmt.Errorf("invalid amount of fields")
	}

	field0 := strings.ReplaceAll(fields[0].Data.String(), "\"", "")
	return field0, nil
}

func (g GameObject) getRowsPlayerOneCurrentMon(arg0 string) []string {
	table := g.world.GetTableByName("PlayerOneCurrentMon")
	rows := g.db.GetRows(table)
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

func (g *GameObject) getPlayerTwoCurrentMon(key string) (string, error) {
	fields, err := data.GetRowFieldsUsingString(g.db, g.world, key, "PlayerTwoCurrentMon")
	if err != nil {
		return "", err
	}

	if len(fields) != 1 {
		return "", fmt.Errorf("invalid amount of fields")
	}

	field0 := strings.ReplaceAll(fields[0].Data.String(), "\"", "")
	return field0, nil
}

func (g GameObject) getRowsPlayerTwoCurrentMon(arg0 string) []string {
	table := g.world.GetTableByName("PlayerTwoCurrentMon")
	rows := g.db.GetRows(table)
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

func (g *GameObject) getPlayerFirstMon(key string) (string, error) {
	fields, err := data.GetRowFieldsUsingString(g.db, g.world, key, "PlayerFirstMon")
	if err != nil {
		return "", err
	}

	if len(fields) != 1 {
		return "", fmt.Errorf("invalid amount of fields")
	}

	field0 := strings.ReplaceAll(fields[0].Data.String(), "\"", "")
	return field0, nil
}

func (g GameObject) getRowsPlayerFirstMon(arg0 string) []string {
	table := g.world.GetTableByName("PlayerFirstMon")
	rows := g.db.GetRows(table)
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

func (g *GameObject) getPlayerSecondMon(key string) (string, error) {
	fields, err := data.GetRowFieldsUsingString(g.db, g.world, key, "PlayerSecondMon")
	if err != nil {
		return "", err
	}

	if len(fields) != 1 {
		return "", fmt.Errorf("invalid amount of fields")
	}

	field0 := strings.ReplaceAll(fields[0].Data.String(), "\"", "")
	return field0, nil
}

func (g GameObject) getRowsPlayerSecondMon(arg0 string) []string {
	table := g.world.GetTableByName("PlayerSecondMon")
	rows := g.db.GetRows(table)
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

func (g *GameObject) getPlayerThirdMon(key string) (string, error) {
	fields, err := data.GetRowFieldsUsingString(g.db, g.world, key, "PlayerThirdMon")
	if err != nil {
		return "", err
	}

	if len(fields) != 1 {
		return "", fmt.Errorf("invalid amount of fields")
	}

	field0 := strings.ReplaceAll(fields[0].Data.String(), "\"", "")
	return field0, nil
}

func (g GameObject) getRowsPlayerThirdMon(arg0 string) []string {
	table := g.world.GetTableByName("PlayerThirdMon")
	rows := g.db.GetRows(table)
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

func (g *GameObject) getMon(key string) (bool, error) {
	fields, err := data.GetRowFieldsUsingString(g.db, g.world, key, "Mon")
	if err != nil {
		return false, err
	}

	if len(fields) != 1 {
		return false, fmt.Errorf("invalid amount of fields")
	}

	field0 := fields[0].Data.String() == "true"
	return field0, nil
}

func (g GameObject) getRowsMon(arg0 bool) []string {
	table := g.world.GetTableByName("Mon")
	rows := g.db.GetRows(table)
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

func (g *GameObject) getMonSpecie(key string) (int64, error) {
	fields, err := data.GetRowFieldsUsingString(g.db, g.world, key, "MonSpecie")
	if err != nil {
		return 0, err
	}

	if len(fields) != 1 {
		return 0, fmt.Errorf("invalid amount of fields")
	}

	field0, err := strconv.ParseInt(fields[0].Data.String(), 10, 32)
	if err != nil {
		return 0, err
	}
	return field0, nil
}

func (g GameObject) getRowsMonSpecie(arg0 int64) []string {
	table := g.world.GetTableByName("MonSpecie")
	rows := g.db.GetRows(table)
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

func (g *GameObject) getMonHp(key string) (int64, error) {
	fields, err := data.GetRowFieldsUsingString(g.db, g.world, key, "MonHp")
	if err != nil {
		return 0, err
	}

	if len(fields) != 1 {
		return 0, fmt.Errorf("invalid amount of fields")
	}

	field0, err := strconv.ParseInt(fields[0].Data.String(), 10, 32)
	if err != nil {
		return 0, err
	}
	return field0, nil
}

func (g GameObject) getRowsMonHp(arg0 int64) []string {
	table := g.world.GetTableByName("MonHp")
	rows := g.db.GetRows(table)
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

func (g *GameObject) getInventoryFirstMon(key string) (int64, error) {
	fields, err := data.GetRowFieldsUsingString(g.db, g.world, key, "InventoryFirstMon")
	if err != nil {
		return 0, err
	}

	if len(fields) != 1 {
		return 0, fmt.Errorf("invalid amount of fields")
	}

	field0, err := strconv.ParseInt(fields[0].Data.String(), 10, 32)
	if err != nil {
		return 0, err
	}
	return field0, nil
}

func (g GameObject) getRowsInventoryFirstMon(arg0 int64) []string {
	table := g.world.GetTableByName("InventoryFirstMon")
	rows := g.db.GetRows(table)
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

func (g *GameObject) getInventorySecondMon(key string) (int64, error) {
	fields, err := data.GetRowFieldsUsingString(g.db, g.world, key, "InventorySecondMon")
	if err != nil {
		return 0, err
	}

	if len(fields) != 1 {
		return 0, fmt.Errorf("invalid amount of fields")
	}

	field0, err := strconv.ParseInt(fields[0].Data.String(), 10, 32)
	if err != nil {
		return 0, err
	}
	return field0, nil
}

func (g GameObject) getRowsInventorySecondMon(arg0 int64) []string {
	table := g.world.GetTableByName("InventorySecondMon")
	rows := g.db.GetRows(table)
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

func (g *GameObject) getInventoryThirdMon(key string) (int64, error) {
	fields, err := data.GetRowFieldsUsingString(g.db, g.world, key, "InventoryThirdMon")
	if err != nil {
		return 0, err
	}

	if len(fields) != 1 {
		return 0, fmt.Errorf("invalid amount of fields")
	}

	field0, err := strconv.ParseInt(fields[0].Data.String(), 10, 32)
	if err != nil {
		return 0, err
	}
	return field0, nil
}

func (g GameObject) getRowsInventoryThirdMon(arg0 int64) []string {
	table := g.world.GetTableByName("InventoryThirdMon")
	rows := g.db.GetRows(table)
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
