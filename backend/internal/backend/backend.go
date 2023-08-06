package backend

import (
	"crypto/ecdsa"
	_ "embed"

	server "github.com/bocha-io/game-backend/x"
	"github.com/bocha-io/garnet/x/indexer"
	"github.com/bocha-io/garnet/x/indexer/data"
	"github.com/bocha-io/logger"
	"github.com/bocha-io/superhack/internal/backend/handlers"
)

//go:embed IWorld.abi.json
var iworldAbiJSON []byte

func Run(port int, pk *ecdsa.PrivateKey, worldID string, endpoint string, usersMnemonics string) {
	// Log to file
	file := logger.LogToFile("indexerlogs.log")
	defer file.Close()

	// Index the database
	quit := false
	database := data.NewDatabase()
	database.SetDefaultWorld(worldID)
	go indexer.Process(endpoint, database, &quit)

	s := handlers.NewBackend(iworldAbiJSON, database, endpoint, worldID, pk, usersMnemonics)
	_ = server.StartGorillaServer(port, database, s.HandleMessage)
}
