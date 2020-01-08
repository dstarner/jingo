NAME:= jingo
MAIN_LOC:=pkg/main/main.go

ifeq ($(GOOS),)
  GOOS:=$(shell uname | tr '[:upper:]' '[:lower:]')
endif

ifeq ($(GOARCH),)
  GOARCH:=arm64
endif

BINARY:=bin/jingo-$(GOOS)-$(GOARCH)

build: bin/$(BINARY)

bin/$(BINARY):
	GOOS=$(GOOS) go build -o $(BINARY) $(MAIN_LOC)

.PHONY: test
test:
	go test $(GO_TEST_ARGS) ./...

.PHONY: test-cover
COVER_PROFILE=coverage.out
test-cover: GO_TEST_ARGS = -coverprofile=$(COVER_PROFILE)
test-cover: test
	go tool cover -html=$(COVER_PROFILE) -o coverage.html

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: vendor
vendor:
	go mod vendor

.PHONY: clean
clean:
	rm -rf vendor
	rm -rf bin
	rm -rf coverage.html coverage.out

.PHONY: compile
compile:
	@echo "Compiling for every OS and Platform"
	@GOOS=linux GOARCH=arm $(MAKE) build
	@GOOS=linux GOARCH=arm64 $(MAKE) build
	@GOOS=linux GOARCH=amd64 $(MAKE) build
	@GOOS=darwin GOARCH=amd64 $(MAKE) build
	@GOOS=freebsd GOARCH=386 $(MAKE) build
	@GOOS=windows GOARCH=amd64 $(MAKE) build

all: build