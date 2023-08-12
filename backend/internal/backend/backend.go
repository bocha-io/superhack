package backend

import (
	"crypto/ecdsa"
	_ "embed"
	"fmt"
	"time"

	server "github.com/bocha-io/game-backend/x"
	"github.com/bocha-io/garnet/x/indexer"
	"github.com/bocha-io/garnet/x/indexer/data"
	"github.com/bocha-io/logger"
	"github.com/bocha-io/superhack/internal/backend/handlers"
	"github.com/bocha-io/txbuilder/x/txbuilder"
)

//go:embed IWorld.abi.json
var iworldAbiJSON []byte

//go:embed ERC20.abi.json
var erc20AbiJSON []byte

//go:embed Bridge.abi.json
var bridgeAbiJSON []byte

func Run(
	port int,
	pk *ecdsa.PrivateKey,
	worldID string,
	endpoint string,
	usersMnemonics string,
	erc20Address string,
	bridgeAddress string,
	startingHeight uint64,
) {
	// Log to file
	file := logger.LogToFile("indexerlogs.log")
	defer file.Close()

	// Index the database
	quit := false
	database := data.NewDatabase()
	database.SetDefaultWorld(worldID)
	go indexer.Process(endpoint, database, &quit, startingHeight, 2*time.Second)
	_, b, _ := txbuilder.GetWallet(usersMnemonics, 0)
	logger.LogInfo(fmt.Sprintf("[ADMIN] admin wallet is: %s", b.Address.Hex()))

	s := handlers.NewBackend(
		iworldAbiJSON,
		database,
		endpoint,
		worldID,
		pk,
		usersMnemonics,
		erc20AbiJSON,
		erc20Address,
		bridgeAbiJSON,
		bridgeAddress,
	)
	_ = server.StartGorillaServer(port, database, s.HandleMessage)
}
