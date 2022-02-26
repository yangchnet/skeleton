DOCKER := docker

IMAGES := mysql

ifeq (${IMAGES},)
  $(error Could not determine IMAGES, set ROOT_DIR or run in source dir)
endif

.PHONY: image.build
image.build: $(addprefix image.build.,$(IMAGES))

.PHONY: image.build.%
image.build.%: #FIXME
	@$(DOCKER) build --file $(ROOT_DIR)/build/docker/$*/Dockerfile -t $(PROJECT_NAME)-$* $(ROOT_DIR)

MYSQL_PORT := 13306
.PHONY: image.up.mysql
image.up.mysql:
	@$(DOCKER) run --name $(PROJECT_NAME)-mysql-runtime -e MYSQL_ROOT_PASSWORD=$(PASSWORD) --restart=always -p $(MYSQL_PORT):3306 -v ~/.$(PROJECT_NAME)-data/mysql:/var/lib/mysql -d $(PROJECT_NAME)-mysql

.PHONY: image.up.redis
image.up.redis:
	@$(DOCKER) run --name $(PROJECT_NAME)-redis-runtime --restart=always -p 16379:6379 -d $(PROJECT_NAME)-redis

