default: generate

PROTO_FILES=$(sort $(wildcard ./*.proto))
THIRD_PARTY_PATH=../third-party/

PROTOC_FLAGS:=-I. -I$(THIRD_PARTY_PATH)

.PHONY: generate
generate:
	@rm -rf ../../pkg/pb/basic/*
	@rm -rf ${API_GATEWAY_PATH}/api.swagger.json
	@rm -rf ../../docs/basic_api.html
	@rm -rf ../../docs/basic_api.md
	@echo "remove old files ok"

	@protoc $(PROTOC_FLAGS) --go_out=plugins=grpc:../../pkg/pb/basic $(PROTO_FILES)
	@echo "generate pb.go files ok"

	@protoc $(PROTOC_FLAGS) --doc_out=../../docs/ --doc_opt=markdown,basic_api.md $(PROTO_FILES)
	@protoc $(PROTOC_FLAGS) --doc_out=../../docs/ --doc_opt=html,basic_api.html $(PROTO_FILES)
	@echo "generate api.md & api.html files ok"

	@echo "basic generate finished"