NAME:= jingo
MAIN_LOC:=pkg/main/main.go

.DEFAULT_GOAL:=all

ifeq ($(GOOS),)
  GOOS:=$(shell uname | tr '[:upper:]' '[:lower:]')
endif

ifeq ($(GOARCH),)
  GOARCH:=arm64
endif

BINARY:=bin/jingo-$(GOOS)-$(GOARCH)

# TARGET: Build the CLI binary easily
build: bin/$(BINARY)

# TARGET: Build the CLI binary from the $BINARY variable
bin/$(BINARY):
	GOOS=$(GOOS) go build -o $(BINARY) $(MAIN_LOC)

# TARGET: Run tests against the repository using the args passed with $GO_TEST_ARGS
.PHONY: test
test:
	go test $(GO_TEST_ARGS) ./...

# TARGET: Run tests and report the coverage details to `coverage.html`
.PHONY: test-cover
COVER_PROFILE=coverage.out
test-cover: GO_TEST_ARGS = -coverprofile=$(COVER_PROFILE)
test-cover: test
	go tool cover -html=$(COVER_PROFILE) -o coverage.html

# TARGET: clean up `go.[mod|sum] and download missing dependencies
.PHONY: tidy
tidy:
	go mod tidy

# TARGET: Move the project dependencies to a `vendor` directory
.PHONY: vendor
vendor:
	go mod vendor

# TARGET: Clean up and remove all built artifacts to restore the repository to its original state
.PHONY: clean
clean:
	rm -rf vendor
	rm -rf bin
	rm -rf coverage.html coverage.out

# TARGET: Compile the CLI for the different Operating Systems & architectures
.PHONY: cross-build
cross-build:
	@echo "Compiling for every OS and Platform"
	@GOOS=linux GOARCH=arm $(MAKE) build
	@GOOS=linux GOARCH=arm64 $(MAKE) build
	@GOOS=linux GOARCH=amd64 $(MAKE) build
	@GOOS=darwin GOARCH=amd64 $(MAKE) build
	@GOOS=freebsd GOARCH=386 $(MAKE) build
	@GOOS=windows GOARCH=amd64 $(MAKE) build

# TARGET: run tests and build the CLI for the current operating system
all: test build