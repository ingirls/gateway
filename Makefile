GOPATH:=$(shell go env GOPATH)

.PHONY: build
build:
	go build -mod=vendor -o gateway *.go

.PHONY: run
run: build
	./gateway --metrics=prometheus --registry=consul --registry_address=consul:8500 api
