APPS ?= openapi iam-service

.PHONY: run
run: $(addprefix run.,$(APPS))

.PHONY: run.%
run.%:
	@echo "==========> Starting $* server"
	@go build -o _output/app/debug-$* cmd/$*/$*.go && _output/app/debug-$*