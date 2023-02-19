
build:
	go build minialert.go

build-linux:
	$ GOOS=linux GOARCH=amd64 go build -o bin/app-amd64-linux minialert.go

test:
	go test minialert/server minialert/client minialert/cmd

format:
	gofmt .
