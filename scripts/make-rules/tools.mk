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
	@$(GO) install google.golang.org/protobuf/cmd/protoc-gen-go@latest

.PHONY: install.protoc-gen-grpc-gateway
install.protoc-gen-grpc-gateway:
	@$(GO) install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest

.PHONY: install.protoc-gen-openapiv2
install.protoc-gen-openapiv2:
	@$(GO) install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest

.PHONY: install.protoc-gen-go-grpc
install.protoc-gen-go-grpc:
	@$(GO) install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

.PHONY: install.goimports
install.goimports:
	@$(GO) install golang.org/x/tools/cmd/goimports@latest

.PHONY: install.wire
install.wire:
	@$(GO) install github.com/google/wire/cmd/wire@latest

# deprecated
.PHONY: install.sqlc
install.sqlc:
	@$(GO) install github.com/kyleconroy/sqlc/cmd/sqlc@latest

BUF_BIN := $(GOBIN)
BUF_VERSION := 1.0.0-rc12
.PHONY: install.buf
install.buf: 
	@$(ROOT_DIR)/scripts/buf-install.sh

# deprecated
.PHONY: install.migrate
install.migrate: 
	@$(GO) install -tags '$(DB_TYPE)' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

.PHONY: install.ent
install.ent:
	@$(GO) install entgo.io/ent/cmd/ent

.PHONY: install.pm2
install.pm2:
	@npm install -g pm2