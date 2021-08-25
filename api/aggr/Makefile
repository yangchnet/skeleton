default: generate

PROTO_FILES=$(sort $(wildcard ./*.proto))
THIRD_PARTY_PATH=../third-party/
API_GATEWAY_PATH=../../pkg/api-gateway/spec

PROTOC_FLAGS:=-I. -I$(THIRD_PARTY_PATH)

.PHONY: generate
generate:
	@rm -rf ../../pkg/pb/aggr/*
	@rm -rf ${API_GATEWAY_PATH}/api.swagger.json
	@rm -rf ../../docs/aggr_api.html
	@rm -rf ../../docs/aggr_api.md
	@echo "remove old files ok"

	@protoc $(PROTOC_FLAGS) --go_out=plugins=grpc:../../pkg/pb/aggr $(PROTO_FILES)
	@echo "generate pb.go files ok"

	@protoc $(PROTOC_FLAGS) --grpc-gateway_out=logtostderr=true,allow_delete_body=true:../../pkg/pb/aggr $(PROTO_FILES)
	@echo "generate pb.gw.go files ok"

	@protoc $(PROTOC_FLAGS) --swagger_out=logtostderr=true,allow_delete_body=true:${API_GATEWAY_PATH} $(PROTO_FILES)
	@cd ${API_GATEWAY_PATH}/ && cat ./*.swagger.json 0.swagger.json | jq --slurp 'reduce .[] as $$item ({}; . * $$item)' | jq -f ./preprocess.jq > ./api.swagger.json
	@cd ${API_GATEWAY_PATH}/ && find . | grep .swagger.json | grep -v "api" | xargs rm
	@echo "generate api.swagger.json files ok"

	@protoc $(PROTOC_FLAGS) --doc_out=../../docs/ --doc_opt=markdown,aggr_api.md $(PROTO_FILES)
	@protoc $(PROTOC_FLAGS) --doc_out=../../docs/ --doc_opt=html,aggr_api.html $(PROTO_FILES)
	@echo "generate api.md & api.html files ok"

	@echo "aggr generate finished"