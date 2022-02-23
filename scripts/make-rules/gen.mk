GEN_LIST ?= wire sqlc proto

.PHONY: gen.run
gen.run: $(addprefix gen., $(GEN_LIST))

WIRE_PACKAGES ?= $(wildcard $(ROOT_DIR)/internal/*/service)
.PHONY: gen.wire
gen.wire: tools.verify.wire
	@echo "==========> Generating from wire.go"
	@wire gen $(WIRE_PACKAGES)

# deprecated
SQLC_CONFIG_FILE := $(ROOT_DIR)/db/sqlc.yaml
.PHONY: gen.sqlc
gen.sqlc: tools.verify.sqlc
	@echo "==========> Generating golang query statement from sql"
	@sqlc generate -f $(SQLC_CONFIG_FILE)

PROTO_GENS := api conf
.PHONY: gen.proto
gen.proto: $(addprefix gen.proto., $(PROTO_GENS))

.PHONY: gen.proto.api
gen.proto.api: $(addprefix tools.verify., $(PROTO_TOOLS))
	@echo "==========> Generating pb.go for api"
	@cd api && buf generate

.PHONY: gen.proto.conf
gen.proto.conf: tools.verify.protoc-gen-go
	@echo "==========> Generating pb.go for conf"
	@cd internal && buf generate

.PHONY: gen.proto.update
gen.proto.update: $(addprefix tools.verify., $(PROTO_TOOLS))
	@cd api && buf mod update
	@cd internal && buf mod update
