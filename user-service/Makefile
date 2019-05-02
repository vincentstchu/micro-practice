build:
	protoc --proto_path=. --go_out=plugins=micro:. proto/user/user.proto
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build
	sudo docker build -t user-service .

run:
	sudo docker run -e MICRO_REGISTRY=mdns user-service