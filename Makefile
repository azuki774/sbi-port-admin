CURRENT_DIR=$(shell pwd)
BUILD_DIR=$(CURRENT_DIR)/build
BIN_DIR=$(BUILD_DIR)/bin/
CONTAINER_NAME_SERVER:=sbiport-server
CONTAINER_NAME_FETCHER:=sbiport-fetcher
CONTAINER_NAME_CLIENT:=sbiport-fetcher

.PHONY: build start stop test restart

build:
	go build -a -tags "netgo" -installsuffix netgo  -ldflags="-s -w -extldflags \"-static\"" -o build/bin/ ./...
	docker build -t ${CONTAINER_NAME_SERVER} -f build/Dockerfile-server .
	docker build -t ${CONTAINER_NAME_FETCHER} -f build/Dockerfile-fetcher .
	docker build -t ${CONTAINER_NAME_CLIENT} -f build/Dockerfile-client .

start:
	docker compose -f deployment/compose-local.yml up -d

stop:
	docker compose -f deployment/compose-local.yml down

test:
	go test -v ./...

restart:
	make stop
	make start

migration-test:
	make restart
	sleep 15s
	python3 test/check.py
