all: run

run:
	go run cmd/main/main.go

push:
	git push git@github.com:RB-PRO/HFLabs.git

pull:
	git pull git@github.com:RB-PRO/HFLabs.git

build-config:
	go env GOOS GOARCH

build-linux-osx:
	export GOARCH=arm export GOOS=linux go build ./cmd/main/main.go  

build-linux-linux:
	export GOARCH=amd64 export GOOS=linux go build ./cmd/main/main.go

build-linux-windows:
	export GOARCH=amd64 export GOOS=windows go build ./cmd/main/main.go  

build:
	go build ./cmd/main/main.go  