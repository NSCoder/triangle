.PHONY: clean build

all: test

clean: 
	rm -rf ./bin

build_cmd = go build

build:
	$(build_cmd) -o ./bin/triangle ./cmd/cli.go

test: build
	go test -coverprofile=cp.out $(shell go list ./... | grep -v /vendor/)

run: build
	./bin/triangle solve -alpha 1.0471975512 -b 2 -gamma 0.78

start: run