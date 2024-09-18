GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)

ifeq ($(GOHOSTOS), windows)
	#the `find.exe` is different from `find` in bash/shell.
	#to see https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/find.
	#changed to use git-bash.exe to run find cli or other cli friendly, caused of every developer has a Git.
	Git_Bash= $(subst cmd\,bin\bash.exe,$(dir $(shell where git)))
	INTERNAL_PROTO_FILES=$(shell $(Git_Bash) -c "find internal -name *.proto")
	API_PROTO_FILES=$(shell $(Git_Bash) -c "find api -name *.proto")
else
	INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
	API_PROTO_FILES=$(shell find api -maxdepth 2 -name *.proto)
endif

.PHONY: init
# init env
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.32.0
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@v0.7.0
	go install github.com/envoyproxy/protoc-gen-validate@v1.0.4
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOPATH)/bin v1.55.2
	go install github.com/go-cinch/cinch/cmd/cinch@latest

.PHONY: config
# generate internal proto
config:
	protoc --proto_path=./internal \
	       --proto_path=./third_party \
	       --go_out=paths=source_relative:./internal \
	       $(INTERNAL_PROTO_FILES)

.PHONY: sub
# update submodule
sub:
	git submodule update --force --recursive --init --remote

.PHONY: api
# generate api proto
api:
	@mkdir -p docs
	for NAME in $(API_PROTO_FILES); do \
		ROOT=$(shell pwd); \
		DIR=`echo $$NAME | awk -F '-proto/[^/]*$$' '{print $$1}'`; \
		echo $$NAME; \
		protoc --proto_path=./api \
        	--proto_path=./third_party \
        	--go_out=. \
        	--go-errors_out=. \
        	--go-http_out=. \
        	--go-grpc_out=. \
        	--validate_out=lang=go:. \
        	--openapi_out=fq_schema_naming=true,default_response=false,output_mode=source_relative:docs \
        	$$NAME; \
	done
	@echo 'You can import *.json into https://editor.swagger.io/'

.PHONY: build
# build
build:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...

.PHONY: gen
# generate wire
gen:
	go mod tidy
	go get github.com/google/wire/cmd/wire@latest
	go generate ./...
	go mod tidy

.PHONY: lint
# golangci-lint
lint:
	golangci-lint run --fix

.PHONY: all
# generate all
all:
	make api;
	make config;
	make gen;
	make lint;

local:
	cinch run

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
