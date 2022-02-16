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
include scripts/make-rules/gen.mk
include scripts/make-rules/up.mk

.PHONY: build
build:
	@$(MAKE) go.build

.PHONY: clean
clean:
	@$(MAKE) go.clean

.PHONY: tools
tools:
	@$(MAKE) tools.install

.PHONY: up
up:
	@$(MAKE) up

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


