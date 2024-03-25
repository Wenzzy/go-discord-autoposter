include .env

LOCAL_BIN := $(CURDIR)/LDE/bin


# Dependencies ===============

install:
	mkdir -p $(LOCAL_BIN)
	make install-golangci-lint
	make install-go-air-livereload

install-golangci-lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.3

install-go-air-livereload:
	GOBIN=$(LOCAL_BIN) go install github.com/cosmtrek/air@v1.51.0

# ============================



# Go Dependencies ============

t: 
	go mod tidy

# ============================




# Main =======================


run:
# make lint

	go run cmd/app/main.go

run-watch:
	$(LOCAL_BIN)/air -c .air.toml

lint:
	$(LOCAL_BIN)/golangci-lint run ./... --config .golangci.pipeline.yaml

build:
	go build -o bin/app cmd/app/main.go

build-win:
	GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CXX=x86_64-w64-mingw32-g++ CC=x86_64-w64-mingw32-gcc go build -o bin/app.exe cmd/app/main.go

local-build:
	go build -o bin/macos/autoposter cmd/app/main.go
	cp config.yml.example bin/macos/config.yml
	cp README.md bin/macos/README.md

local-build-win:
	GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CXX=x86_64-w64-mingw32-g++ CC=x86_64-w64-mingw32-gcc go build -o bin/win64/autoposter.exe cmd/app/main.go
	cp config.yml.example bin/win64/config.yml
	cp README.md bin/win64/README.md

# ============================




# Datavase migrations ========

local-migration-create:
	make _required_param-name

	$(LOCAL_BIN)/goose -dir ${MIGRATIONS_DIR} create $(name) sql

local-migration-status:
	$(LOCAL_BIN)/goose -dir ${MIGRATIONS_DIR} sqlite3 ${SQLITE_FILE} status -v

local-migration-up:
	$(LOCAL_BIN)/goose -dir ${MIGRATIONS_DIR} sqlite3 ${SQLITE_FILE} up -v

local-migration-down:
	$(LOCAL_BIN)/goose -dir ${MIGRATIONS_DIR} sqlite3 ${SQLITE_FILE} down -v

# ============================




# Make sys ===================

_required_param-%:
	@if [ "${${*}}" == "" ]; then echo "\n\033[0;91mPlease provide arg: \"$*\"\033[0m\n"; exit 1; fi

# ============================