ITERATION := $(shell date +%s)
ENV ?= "$(shell uname -n)"
COMMIT ?= $$(git rev-parse HEAD)
TAG ?= $$(git describe --tags --abbrev=0)

init:
	go mod init typidor_exporter

vet:
	go vet -mod=vendor $(shell go list ./...)

deps:
	go mod tidy
	go mod vendor

build:
	GOOS=linux GOARCH=amd64 go build -mod=vendor \
		-ldflags "-X main.version=${TAG} -X main.commit=${COMMIT} -X main.iteration=${ITERATION} -X main.env=${ENV}" \
		-o typidor-exporter
