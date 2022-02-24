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
include scripts/make-rules/run.mk
include scripts/make-rules/image.mk
include scripts/make-rules/migrate.mk
include scripts/make-rules/ent.mk

.PHONY: build
build:
	@$(MAKE) go.build

.PHONY: clean
clean:
	@$(MAKE) go.clean

.PHONY: tools
tools:
	@$(MAKE) tools.install

.PHONY: run
run:
	@$(MAKE) run
