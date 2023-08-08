package garnethelpers

import (
	"fmt"
	"strings"
)

func (Prediction) monKey(playerKey string, ID int64) string {
	// The mons index starts at 1
	ID++
	if ID < 0 || ID > 9 {
		return playerKey
	}
	return strings.Replace(
		playerKey,
		"0x000000000000000000000000",
		fmt.Sprintf("0x3%d0000000000000000000000", ID),
		1,
	)
}
