# make build     # builds the project
# make run       # builds and runs the binary
# make test      # runs tests
# make fmt       # formats code
# make tidy      # go mod tidy
# make lint      # runs golint
# make clean     # removes binaries


# Project settings
APP_NAME := go-rest-api
PKG := ./...
BIN_DIR := bin

# Commands
GO := go
GOTEST := $(GO) test
GOBUILD := $(GO) build
GOLINT := golint

# Flags
LDFLAGS := -s -w
BUILD_FLAGS := -o bin/go-rest-api

.PHONY: all build run clean test lint tidy fmt deps

all: build

build:
	@go build -o bin/go-rest-api ./cmd/main.go

run: build
	@./bin/go-rest-api

