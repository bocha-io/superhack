#!/bin/bash
make build
export PORT=9000
# Anvil key
export PRIV=ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
# Random mnemonic for players accounts
export MNEMONIC="cash enlist manage casino tuition creek name material toe doctor ridge region slush second trust manual quit tissue purchase rebel curve element loyal duck"
# Current world address
export WORLD=0x5FbDB2315678afecb367f032d93F642f64180aa3
./build/backend
