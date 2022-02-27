APPS ?= openapi iam-service

.PHONY: run
run: tools.verify.pm2
	@pm2 start skeleton.config.js

.PHONY: stop
stop: tools.verify.pm2
	@pm2 del all

.PHONY: run.%
run.%:
	@echo "==========> Starting $* server"
	@go build -o _output/app/debug-$* cmd/$*/$*.go && _output/app/debug-$*