CURRENT_DIR=$(shell pwd)
BUILD_DIR=$(CURRENT_DIR)/build
BIN_DIR=$(BUILD_DIR)/bin/

.PHONY: build start stop test restart

build:
	CGO_ENABLED=0 go build -o build/bin/ ./...

start:
	docker compose -f deployment/compose-local.yml up -d

stop:
	docker compose -f deployment/compose-local.yml down

test:
	go test -v ./...

restart:
	make stop
	make start
