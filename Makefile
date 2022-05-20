GOCMD=go
GORUN=$(GOCMD) run
PROJECT_NAME=grpc-client

init:
	mkdir -p pb | echo ""
	$(GOCMD) mod init $(PROJECT_NAME)

tidy:
	$(GOCMD) mod tidy

proto-gen:
	protoc proto/*/*.proto --go_out=pb --go_opt=paths=source_relative --go-grpc_out=require_unimplemented_servers=false:pb --go-grpc_opt=paths=source_relative -I=proto --experimental_allow_proto3_optional

run:
	go run main.go