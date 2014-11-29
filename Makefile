.PHONY: all clean coverage debug-install dependencies fmt install lint test tools

ROOT_DIR := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

all: clean install test

clean:
	go clean -i ./...
	go clean -i -race ./...
coverage:
	go test -coverprofile=coverage.out
	go tool cover -html=coverage.out
crosscompile:
	gox -os="linux" ./...
debug-install: clean
	go install -race -v ./...
dependencies:
	go get -t -v ./...
	go build -v ./...
fmt:
	gofmt -l -w $(ROOT_DIR)/
install: clean
	go install -v ./...
lint: install fmt
	errcheck github.com/zimmski/tavor/...
	golint $(ROOT_DIR)/...
	go tool vet -all=true -v=true $(ROOT_DIR)/ 2>&1 | grep --invert-match -P "(Checking file|\%p of wrong type|can't check non-constant format)" || true
markdown:
	orange
test:
	go test -race ./...
tools:
	go get -u code.google.com/p/go.tools/cmd/cover
	go get -u code.google.com/p/go.tools/cmd/godoc
	go get -u code.google.com/p/go.tools/cmd/vet
	go get -u github.com/golang/lint
	go install github.com/golang/lint/golint
	go get -u github.com/kisielk/errcheck
