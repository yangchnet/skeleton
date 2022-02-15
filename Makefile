.DEFAULT_GOAL := all

.PHONY: all
all: 

# ==============================================================
# Build option

ROOT_PACKAGE=github.com/yangchnet/skeleton
VERSION_PACAKGE=github.com/yangchnet/component-base/pkg/version

# ==============================================================
# Include 

include scripts/make-rules/common.mk
include scripts/make-rules/tools.mk
include scripts/make-rules/golang.mk

.PHONY: build
build:
	@$(MAKE) go.build

.PHONY: clean
clean:
	@$(MAKE) go.clean

.PHONY: tools
tools:
	@$(MAKE) tools.install

generate: 
	@cd api && make
	@echo proto compile succeed!
	@cd pkg/api-gateway && make
	@echo all succeed!!

api-gateway:
	@cd cmd/api-gateway && \
	go build -o debug-api-gateway main.go && ./debug-api-gateway

sqlc:
	@cd db && make sqlc

wire:
	@cd ./internal/iam/service && wire
	@cd ./internal/hello/service && wire
	@cd ./internal/aggr/service && wire

sql-build:
	@cd db && make sql-build

sql-migrate-up:
	@cd db && make migrate-up

sql-migrate-down:
	@cd db && make migrate-down

sql-up:
	@cd db && make sql-up

sql-down:
	@cd db && make sql-down


