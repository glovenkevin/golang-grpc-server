# Application Setup
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

# MongoDB Cluster Setup
docker-prepare:
	mkdir data mongodb mongo-c1 mongo-c2
	docker network create mongo-cluster

mongo-gen-key:
	openssl rand -base64 700 > file.key
	chmod 400 file.key
	mv file.key ./key/file.key

mongo-cluster-init:
	docker-compose exec mongo mongo -u root -p password --eval "rs.initiate({_id: 'rs-mongo', members: [{_id: 0, host: 'mongo:27017'}, {_id: 1, host: 'mongo-c1:27018'}, {_id: 1, host: 'mongo-c2:27019'}]});"
