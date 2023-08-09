.PHONY: dagger build

build:
	@go build -o ./build/backend ./cmd/backend

lint:
	golangci-lint run --fix --out-format=line-number --issues-exit-code=0 --config .golangci.yml --color always ./...

fix-lines:
	@golines -w .

clean:
	@rm users.sql indexerlogs.log

run-localnet:
	@source /opt/homebrew/opt/nvm/nvm.sh && nvm use v18.12.0 && cd contracts-builder/contracts && pnpm run devnode

contracts:
	@source /opt/homebrew/opt/nvm/nvm.sh && nvm use v18.12.0 && cd contracts-builder/contracts && pnpm run dev && cd ../.. && cp contracts-builder/contracts/out/IWorld.sol/IWorld.abi.json internal/backend/

init-contracts:
	@source /opt/homebrew/opt/nvm/nvm.sh && nvm use v18.12.0 && cd contracts-builder/contracts && pnpm install

