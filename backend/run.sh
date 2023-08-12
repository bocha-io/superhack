#!/bin/bash
make build
export PORT=9000
# Anvil key
export PRIV=ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
# Random mnemonic for players accounts
export MNEMONIC="cash enlist manage casino tuition creek name material toe doctor ridge region slush second trust manual quit tissue purchase rebel curve element loyal duck"
# Current world address
export WORLD=0x5FbDB2315678afecb367f032d93F642f64180aa3
# Contract addresses
export ERC20ADDRESS=0x8464135c8F25Da09e49BC8782676a84730C318bC
export BRIDGEADDRESS=0x948B3c65b89DF0B4894ABE91E6D02FE579834F8F

./build/backend
