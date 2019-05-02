BASE=$(shell pwd)

build:
	protoc --proto_path=. --go_out=plugins=micro:. proto/vessel/vessel.proto
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build
	sudo docker build -t vessel-service .

run:
	sudo docker run -p 50052:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns vessel-service