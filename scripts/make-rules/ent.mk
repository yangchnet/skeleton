ENT = ent

## usage example: make ent.<service_name>.<table_name>
.PHONY: ent.%
ent.%: tools.verify.ent
	@$(eval list := $(subst .,$(SPACE),$*))
	@$(eval service := $(word 1,$(list)))
	@$(eval table := $(word 2,$(list)))
	@cd $(ROOT_DIR)/internal/$(service)/data && $(ENT) init $(table)

## usage example: make ent.gen.<service_name>
.PHONY: ent.gen.%
ent.gen.%: tools.verify.ent
	@cd $(ROOT_DIR)/internal/$*/data && $(GO) generate ./ent