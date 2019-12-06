export GOBIN ?= $(shell pwd)/bin

GOLINT = $(GOBIN)/golint
BINARY_NAME=out
COVER_OUT_FILE=cover.out
COVER_HTML_FILE=cover.html
GO_FILES := $(shell \
	find . '(' -path '*/.*' -o -path './vendor' ')' -prune \
	-o -name '*.go' -print | cut -b3-)

all: test build

.PHONY: build
build:
	go build -v -o $(BINARY_NAME) ./src

.PHONY: run
run:	
	./out

.PHONY: gen-mocks
gen-mocks:
	mockgen -source=src/models/model.go -package=mocks -destination=src/mocks/mock_model.go

.PHONY: test
test: gen-mocks
	go test -race ./...

.PHONY: cover
cover: gen-mocks
	go test -race -coverprofile=cover.out -coverpkg=./... ./... 
	go tool cover -html=cover.out -o cover.html

.PHONY: clean
clean:
	@if [ -f "$(COVER_OUT_FILE)" ]; then \
		rm $(COVER_OUT_FILE); \
	fi
	@if [ -f "$(COVER_HTML_FILE)" ]; then \
                rm $(COVER_HTML_FILE); \
        fi
	@rm -f $(BINARY_NAME)
	@rm -rf ./src/mocks
	@go clean
