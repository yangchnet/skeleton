TOOLS ?= $(BLOCKER_TOOLS) $(CRITICAL_TOOLS) $(TRIVIAL_TOOLS)

.PHONY: tools.install
tools.install: $(addprefix tools.install., $(TOOLS))

.PHONY: tools.install.%
tools.install.%:
	@echo "==========> Install $*"
	@$(MAKE) install.$*

.PHONY: tools.verify.%
tools.verify.%:
	@if ! which $* &>/dev/null; then $(MAKE) tools.install.$*; fi

.PHONY: install.golangci-lint
install.golangci-lint:
	@$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@echo 'source <(golangci-lint completion bash)' >>~/.bashrc
	@$(shell source ~/.bashrc)

.PHONY: install.mockgen
install.mockgen:
	@$(GO) install github.com/golang/mock/mockgen@latest

.PHONY: install.protoc-gen-go
install.protoc-gen-go:
	@$(GO) install github.com/golang/protobuf/protoc-gen-go@latest

.PHONY: install.goimports
install.goimports:
	@$(GO) install golang.org/x/tools/cmd/goimports@latest

.PHONY: install.wire
install.wire:
	@$(GO) install github.com/google/wire/cmd/wire@latest

.PHONY: install.sqlc
install.sqlc:
	@$(GO) install github.com/kyleconroy/sqlc/cmd/sqlc@latest
