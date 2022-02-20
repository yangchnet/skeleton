# deprecated

MIGRATE := migrate

DB_NAMES := $(notdir $(wildcard $(ROOT_DIR)/db/schema/*))

.PHONY: migrate.up
migrate.up: $(addprefix migrate.up.,$(DB_NAMES))

.PHONY: migrate.up.%
migrate.up.%: tools.verify.$(MIGRATE) 
	@$(MIGRATE) -source=file://$(addprefix db/schema/,$*) -database \
	 "$(DB_TYPE)://root:$(PASSWORD)@tcp(localhost:$(DB_PORT))/$*" \
	 -verbose up

.PHONY: migrate.down
migrate.down: $(addprefix migrate.down.,$(DB_NAMES))

.PHONY: migrate.down.%
migrate.down.%: tools.verify.$(MIGRATE) 
	@$(MIGRATE) -source=file://$(addprefix db/schema/,$*) -database \
	 "$(DB_TYPE)://root:$(PASSWORD)@tcp(localhost:$(DB_PORT))/$*" \
	 -verbose down

# make migrate.goto.hello.1
.PHONY: migrate.goto.%
migrate.goto.%: tools.verify.$(MIGRATE)
	@$(MIGRATE) -source=file://$(addprefix db/schema/,$(word 1,$(subst .,$(SPACE),$*))) -database \
	 "$(DB_TYPE)://root:$(PASSWORD)@tcp(localhost:$(DB_PORT))/$(word 1,$(subst .,$(SPACE),$*))" \
	 -verbose goto $(word 2,$(subst .,$(SPACE),$*))
