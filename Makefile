GOCMD := go

GOBUILD    := $(GOCMD) build
GOGET      := $(GOCMD) get
GOTEST     := $(GOCMD) test
GOGENERATE := $(GOCMD) generate

GOOS   := $(shell go env GOOS)
GOARCH := $(shell go env GOARCH)
GOPATH := $(shell go env GOPATH)

OUT ?= bin

.PHONY: all clean test generate

all: clean build test

build:
	rm -rf  $(OUT)/*
#	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(OUT)/autotrade-linux .
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(OUT)/autotrade .

zip: build
	zip -r $(OUT)/netbox.zip ${OUT}/*

test:
	$(GOTEST) -covermode=atomic -cover ./...

format:
	find . -name '*.go' -type f -not -path './vendor/*' -exec go fmt {} \;;

clean:
	rm -f $(OUT)/*
	packr2 clean
