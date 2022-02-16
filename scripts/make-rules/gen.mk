GEN_LIST ?= wire sqlc proto
.PHONY: gen.run
gen.run: $(addprefix gen., $(GEN_LIST))

WIRE_PACKAGES ?= $(wildcard $(ROOT_DIR)/internal/*/service)
.PHONY: gen.wire
gen.wire: tools.verify.wire
	@echo "==========> Generating from wire.go"
	@wire gen $(WIRE_PACKAGES)

SQLC_CONFIG_FILE := $(ROOT_DIR)/db/sqlc.yaml
.PHONY: gen.sqlc
gen.sqlc: tools.verify.sqlc
	@echo "==========> Generating golang query statement from sql"
	@sqlc generate -f $(SQLC_CONFIG_FILE)

.PHONY: gen.proto
gen.proto: $(addprefix tools.verify., $(PROTO_TOOLS))
	@echo "==========> Generating pb.go from protobuf file"
	@cd api && buf generate

.PHONY: gen.proto.update
gen.proto.update: $(addprefix tools.verify., $(PROTO_TOOLS))
	@cd api && buf mod update