generate: 
	@cd api && make
	@echo proto compile succeed!
	@cd pkg/api-gateway && make
	@echo all succeed!!

api-gateway:
	@cd cmd/api-gateway && \
	go build -o debug-api-gateway main.go && ./debug-api-gateway

sqlc:
	@cd db && make sqlc

wire:
	@cd ./internal/iam/service && wire
	@cd ./internal/hello/service && wire
	@cd ./internal/aggr/service && wire

sql-build:
	@cd db && make sql-build

sql-migrate-up:
	@cd db && make migrate-up

sql-migrate-down:
	@cd db && make migrate-down

sql-up:
	@cd db && make sql-up

sql-down:
	@cd db && make sql-down


