PACKAGES=$(shell go list ./... | grep -v '/simulation')

VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=Bettery \
	-X github.com/cosmos/cosmos-sdk/version.ServerName=betteryd \
	-X github.com/cosmos/cosmos-sdk/version.ClientName=batterycli \
	-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
	-X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) 

BUILD_FLAGS := -ldflags '$(ldflags)'

all: clean install 

clean:
	rm -f betteryd 
	rm -f batterycli
	rm -f bridge

generateAbi:
	solc --abi ../contracts/contracts/Bridge.sol -o ./cmd/bridge/contract/abi --overwrite
	solc --bin ../contracts/contracts/Bridge.sol -o ./cmd/bridge/contract/abi --overwrite
	abigen --bin=./cmd/bridge/contract/abi/Bridge.bin --abi=./cmd/bridge/contract/abi/Bridge.abi --pkg=store --out=./cmd/bridge/contract/abi/Bridge.go

install: go.sum
		go install -mod=readonly $(BUILD_FLAGS) ./cmd/betteryd
		go install -mod=readonly $(BUILD_FLAGS) ./cmd/betterycli
		go install -mod=readonly $(BUILD_FLAGS) ./cmd/bridge

go.sum: go.mod
		@echo "--> Ensure dependencies have not been modified"
		GO111MODULE=on go mod verify

# Uncomment when you have some tests
# test:
# 	@go test -mod=readonly $(PACKAGES)

# look into .golangci.yml for enabling / disabling linters
lint:
	@echo "--> Running linter"
	@golangci-lint run
	@go mod verify
