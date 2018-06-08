GOCMD=go
GOGEN=$(GOCMD) generate
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOCOVER=$(GOCMD) tool cover
ABIGEN=abigen

## Check to see if `go` is installed
GO=$(shell command -v go 2> /dev/null)

ABI_DIR=abi

## List of expected dirs for generated code
GENERATED_DIR=pkg/generated
GENERATED_CONTRACT_DIR=pkg/generated/contract
GENERATED_WATCHER_DIR=pkg/generated/watcher
GENERATED_FILTERER_DIR=pkg/generated/filterer

EVENTHANDLER_GEN_MAIN=cmd/eventhandlergen/main.go

## Reliant on go and $GOPATH being set.
check-env:
ifndef GO
	$(error go command is not installed or in PATH)
endif
ifndef GOPATH
	$(error GOPATH is not set)
endif

.PHONY: install-dep
install-dep: check-env ## Installs dep
	mkdir -p $(GOPATH)/bin
	curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

.PHONY: install-linter
install-linter: check-env ## Installs linter
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install

.PHONY: install-cover
install-cover: check-env ## Installs code coverage tool
	go get -u golang.org/x/tools/cmd/cover

.PHONY: install-abigen
install-abigen: check-env ## Installs the Ethereum abigen tool
	go get -u github.com/ethereum/go-ethereum/cmd/abigen

.PHONY: setup
setup: check-env install-dep install-linter install-cover install-abigen ## Sets up the tooling.

.PHONY: lint
lint: generate-contracts generate ## Runs linting.
	# gometalinter config in .gometalinter.json
	gometalinter ./...

.PHONY: generate
generate: generate-contracts generate-watchers generate-filterers ## Runs all the code generation

.PHONY: generate-watchers
generate-watchers: ## Runs watchergen to generate contract Watch* wrapper code.
	mkdir -p $(GENERATED_WATCHER_DIR)
	$(GORUN) $(EVENTHANDLER_GEN_MAIN) civiltcr watcher watcher > ./$(GENERATED_WATCHER_DIR)/civiltcr.go
	$(GORUN) $(EVENTHANDLER_GEN_MAIN) newsroom watcher watcher > ./$(GENERATED_WATCHER_DIR)/newsroom.go

.PHONY: generate-filterers
generate-filterers: ## Runs filterergen to generate contract Filter* wrapper code.
	mkdir -p $(GENERATED_FILTERER_DIR)
	$(GORUN) $(EVENTHANDLER_GEN_MAIN) civiltcr filterer filterer > ./$(GENERATED_FILTERER_DIR)/civiltcr.go
	$(GORUN) $(EVENTHANDLER_GEN_MAIN) newsroom filterer filterer > ./$(GENERATED_FILTERER_DIR)/newsroom.go

.PHONY: generate-contracts
generate-contracts: ## Builds the contract wrapper code from the ABIs in /abi.
ifneq ("$(wildcard $(ABI_DIR)/*.abi)", "")
	mkdir -p $(GENERATED_CONTRACT_DIR)
	$(ABIGEN) -abi ./$(ABI_DIR)/CivilTCR.abi -bin ./$(ABI_DIR)/CivilTCR.bin -type CivilTCRContract -out ./$(GENERATED_CONTRACT_DIR)/CivilTCRContract.go -pkg contract
	$(ABIGEN) -abi ./$(ABI_DIR)/Newsroom.abi -bin ./$(ABI_DIR)/Newsroom.bin -type NewsroomContract -out ./$(GENERATED_CONTRACT_DIR)/NewsroomContract.go -pkg contract
	$(ABIGEN) -abi ./$(ABI_DIR)/PLCRVoting.abi -bin ./$(ABI_DIR)/PLCRVoting.bin -type PLCRVotingContract -out ./$(GENERATED_CONTRACT_DIR)/PLCRVotingContract.go -pkg contract
	$(ABIGEN) -abi ./$(ABI_DIR)/Parameterizer.abi -bin ./$(ABI_DIR)/Parameterizer.bin -type ParameterizerContract -out ./$(GENERATED_CONTRACT_DIR)/ParameterizerContract.go -pkg contract
	$(ABIGEN) -abi ./$(ABI_DIR)/Government.abi -bin ./$(ABI_DIR)/Government.bin -type GovernmentContract -out ./$(GENERATED_CONTRACT_DIR)/GovernmentContract.go -pkg contract
	$(ABIGEN) -abi ./$(ABI_DIR)/EIP20.abi -bin ./$(ABI_DIR)/EIP20.bin -type EIP20Contract -out ./$(GENERATED_CONTRACT_DIR)/EIP20.go -pkg contract
else
	$(error No abi files found; copy them to /abi after generation)
endif

.PHONY: build
build: ## Builds the code.
	$(GOBUILD) ./...

.PHONY: test
test: ## Runs unit tests and tests code coverage.
	echo 'mode: atomic' > coverage.txt && $(GOTEST) -covermode=atomic -coverprofile=coverage.txt -v -race -timeout=30s ./...

.PHONY: cover
cover: test ## Runs unit tests, code coverage, and runs HTML coverage tool.
	$(GOCOVER) -html=coverage.txt

.PHONY: clean
clean: ## go clean and clean up of artifacts.
	$(GOCLEAN) ./... || true
	rm coverage.txt || true

## Some magic from http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'