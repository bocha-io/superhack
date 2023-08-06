package main

import (
	"os"
	"strconv"

	"github.com/bocha-io/superhack/internal/backend"
	"github.com/bocha-io/txbuilder/x/txbuilder"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	// Defaults
	port := int64(9000)
	var err error

	portString := os.Getenv("PORT")
	if portString != "" {
		port, err = strconv.ParseInt(portString, 10, 32)
		if err != nil {
			panic("invalid port value")
		}
	}

	pk := txbuilder.PrivateKeysAnvil[0]
	privString := os.Getenv("PRIV")
	if privString != "" {
		pk, err = crypto.HexToECDSA(privString)
		if err != nil {
			panic("invalid private key")
		}
	}

	worldID := "0x5FbDB2315678afecb367f032d93F642f64180aa3"
	worldString := os.Getenv("WORLD")
	if worldString != "" {
		worldID = worldString
	}

	blockchain := "http://localhost:8545"
	endpoint := os.Getenv("endpoint")
	if endpoint != "" {
		blockchain = endpoint
	}

	mnemonicString := os.Getenv("MNEMONIC")
	if mnemonicString == "" {
		panic("MNEMONIC env var is missing")
	}

	backend.Run(int(port), pk, worldID, blockchain, mnemonicString)
}
