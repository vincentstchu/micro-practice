build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build
	sudo docker build -t user-cli .
run:
	sudo docker run -e MICRO_REGISTRY=mdns user-cli


	