APPS ?= openapi

.PHONY: up
up: $(addprefix up.,$(APPS))

.PHONY: up.%
up.%:
	@echo "==========> Starting $* server"
	@go build -o _output/app/debug-$* cmd/$*/$*.go && _output/app/debug-$*