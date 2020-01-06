PROJECT?=github.com/dicaormu/go-xke-jan-2020

test:
	go test -race ./...

build:
	go build -o bin/go-xke ${PROJECT}/cmd/go-xke

docker-build:
	docker build -t go-xke .