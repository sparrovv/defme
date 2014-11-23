all: build

deps:
	go get -d

build: deps
	go build

test: deps
	go test ./...

itest:
	ruby test/integration_test.rb

install: deps
	go install

