CURRENT_DIR=$(shell pwd)
BUILD_DIR=$(CURRENT_DIR)/build
BIN_DIR=$(BUILD_DIR)/bin/
CONTAINER_NAME_SERVER:=sbiport-server

.PHONY: build start stop test restart

build:
	CGO_ENABLED=0 go build -a -tags "netgo" -installsuffix netgo  -ldflags="-s -w -extldflags \"-static\"" -o build/bin/ ./...
	docker build -t ${CONTAINER_NAME_SERVER} -f build/Dockerfile-sbiport-server .

start:
	docker compose -f deployment/compose-local.yml up -d

stop:
	docker compose -f deployment/compose-local.yml down

test:
	go test -v ./...

restart:
	make stop
	make start
