VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT  := $(shell git log -1 --format='%H')

export GO111MODULE = on

###############################################################################
###                                   All                                   ###
###############################################################################

all: lint test-unit install

###############################################################################
###                                Build flags                              ###
###############################################################################

LD_FLAGS = -X github.com/spike-engine/juno/cmd.Version=$(VERSION) \
	-X github.com/spike-engine/juno/cmd.Commit=$(COMMIT)

BUILD_FLAGS := -ldflags '$(LD_FLAGS)'

###############################################################################
###                                  Build                                  ###
###############################################################################

build: go.sum
ifeq ($(OS),Windows_NT)
	@echo "building juno binary..."
	@go build -mod=readonly $(BUILD_FLAGS) -o build/juno.exe ./cmd/juno
else
	@echo "building juno binary..."
	@go build -mod=readonly $(BUILD_FLAGS) -o build/juno ./cmd/juno
endif
.PHONY: build

###############################################################################
###                                 Install                                 ###
###############################################################################

install: go.sum
	@echo "installing juno binary..."
	@go install -mod=readonly $(BUILD_FLAGS) ./cmd/juno
.PHONY: install

###############################################################################
###                           Tests & Simulation                            ###
###############################################################################

stop-docker-test:
	@echo "Stopping Docker container..."
	@docker stop bdjuno-test-db || true && docker rm bdjuno-test-db || true
.PHONY: stop-docker-test

start-docker-test: stop-docker-test
	@echo "Starting Docker container..."
	@docker run --name bdjuno-test-db -e POSTGRES_USER=bdjuno -e POSTGRES_PASSWORD=password -e POSTGRES_DB=bdjuno -d -p 6433:5432 postgres
.PHONY: start-docker-test

coverage:
	@echo "viewing test coverage..."
	@go tool cover --html=coverage.out
.PHONY: coverage

test-unit: start-docker-test
	@echo "Executing unit tests..."
	@go test -mod=readonly -v -coverprofile coverage.txt ./...
.PHONY: test-unit

lint:
	golangci-lint run --out-format=tab
.PHONY: lint

lint-fix:
	golangci-lint run --fix --out-format=tab --issues-exit-code=0
.PHONY: lint-fix

format:
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -name '*.pb.go' | xargs gofmt -w -s
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -name '*.pb.go' | xargs misspell -w
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -name '*.pb.go' | xargs goimports -w -local github.com/forbole/juno
.PHONY: format

clean:
	rm -f tools-stamp ./build/**
.PHONY: clean
