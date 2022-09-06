APP = web-server-tools

all: start

start:
	go run apps/main.go web-server --config ./config/config.yaml

clean:
	rm ./apps.exe

build:
	#CGO_ENABLED=0SET GOOS=darwin3 SET GOARCH=amd64 go build apps/main.go -o apps
    #CGO_ENABLED=0 SET GOOS=linux SET GOARCH=amd64 go build  apps/main.go -o apps
	go build -v apps/ -o apps

proto:
	cd proto
	protoc -I . --go_out=plugins=grpc:. ./*.proto

