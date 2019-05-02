BASE=$(shell pwd)

build:
	protoc --proto_path=. --go_out=plugins=micro:. proto/consignment/consignment.proto 
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build
	sudo docker build -t consignment-service .

run:
	sudo docker run -p 50051:50051 \
	 -e MICRO_SERVER_ADDRESS=:50051 \
	 -e MICRO_REGISTRY=mdns \
	 consignment-service