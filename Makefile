MAKEFILE := $(abspath $(lastword $(MAKEFILE_LIST)))
PROJECT := $(dir $(MAKEFILE))
GOPATH := $(PROJECT)/vendor
GOVENDOR := $(GOPATH)/src/github.com/jayhanjaelee/registrybot
GOPACKAGE := $(GOPATH)/pkg
BIN := $(PROJECT)bin
OUT := $(BIN)/registrybot

export PROJECT
export GOPATH
export VENDOR
export CGO_ENABLED=0

default: build

deps:
	go get github.com/aws/aws-sdk-go/aws
	go get github.com/aws/aws-sdk-go/aws/credentials
	go get github.com/aws/aws-sdk-go/aws/session
	go get github.com/aws/aws-sdk-go/service/ecr
	@find $(GOPATH)/ | grep .git/ | xargs rm -rf
	@rm -rf $(GOPACKAGE)/

build:
	@rm -rf $(GOVENDOR)

	@mkdir -p $(BIN)
	@mkdir -p $(GOVENDOR)

	@go build -a -installsuffix cgo -ldflags '-w' -o $(OUT)

docker:
	docker run -i -v $(PROJECT):/go/src/github.com/jayhanjaelee/registrybot --rm golang:1.10.1 make -C /go/src/github.com/jayhanjaelee/registrybot

run: build
	$(OUT)
