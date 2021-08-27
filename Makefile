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


