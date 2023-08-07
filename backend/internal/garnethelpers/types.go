package garnethelpers

import "github.com/bocha-io/garnet/x/indexer/data"

type GameObject struct {
	db     *data.Database
	world  *data.World
	active bool
}

func NewGameObject(db *data.Database) *GameObject {
	return &GameObject{
		db:     db,
		world:  db.GetDefaultWorld(),
		active: true,
	}
}
