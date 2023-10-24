#!/usr/bin/env bash

CURRENT_DIR := $(shell pwd)

setup: setup-linters

setup-linters:
	go install github.com/yoheimuta/protolint/cmd/protolint@latest

build: compile-protos dc-build

compile-protos:
	docker build -f ./docker/Dockerfile.proto-compiler -t proto-compiler .
	docker run --rm -v "$(CURRENT_DIR)":/workspace \
		proto-compiler \
			--go_out=. \
			--go_opt=paths=source_relative \
			--go-grpc_out=. \
			--go-grpc_opt=paths=source_relative \
			pkg/server/protos/*

cli:
	go run cmd/cli/main.go

server:
	go run cmd/server/main.go

dc-build:
	docker compose build

dc-up:
	docker compose up -d

dc-up-d:
	docker compose up -d

dc-down:
	docker compose down

dc-stop:
	docker compose stop

dc-ps:
	docker compose ps

dc-logs:
	docker compose logs

dc-follow:
	docker compose logs server --follow