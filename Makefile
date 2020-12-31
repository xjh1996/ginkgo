# The old school Makefile, following are required targets. The Makefile is written
# to allow building multiple binaries. You are free to add more targets or change
# existing implementations, as long as the semantics are preserved.
#
#   make              - default to 'build' target
#   make lint         - code analysis
#   make test         - run unit test (or plus integration test)
#   make dryrun       - run ginkgo dryrun test
#   make build        - alias to build-local target
#   make build-local  - build local binary targets
#   make build-linux  - build linux binary targets
#   make container    - build containers
#   $ docker login registry -u username -p xxxxx
#   make push         - push containers
#   make clean        - clean up targets
#
# Not included but recommended targets:
#   make e2e-test
#
# The makefile is also responsible to populate project version information.
#

#
# Tweak the variables based on your project.
#

# This repo's root import path (under GOPATH).
ROOT := github.com/caicloud/zeus

# Module name.
NAME := zeus

# Container image prefix and suffix added to targets.
# The final built images are:
#   $[REGISTRY]/$[IMAGE_PREFIX]$[TARGET]$[IMAGE_SUFFIX]:$[VERSION]
# $[REGISTRY] is an item from $[REGISTRIES], $[TARGET] is an item from $[TARGETS].
IMAGE_PREFIX ?= $(strip )
IMAGE_SUFFIX ?= $(strip )

# Container registries.
REGISTRY ?= cargo.dev.caicloud.xyz/release

# Container registry for base images.
BASE_REGISTRY ?= cargo.caicloud.xyz/library

#
# These variables should not need tweaking.
#

# It's necessary to set this because some environments don't link sh -> bash.
export SHELL := /bin/bash

# It's necessary to set the errexit flags for the bash shell.
export SHELLOPTS := errexit

# Project main package location.
CMD_DIR := ./cmd

# Project output directory.
OUTPUT_DIR := ./bin

# Build directory.
BUILD_DIR := ./build

IMAGE_NAME := $(IMAGE_PREFIX)$(NAME)$(IMAGE_SUFFIX)

# Current version of the project.
GOCOMMON     := $(shell if [ ! -f go.mod ]; then echo $(ROOT)/vendor/; fi)github.com/caicloud/go-common
VERSION      ?= $(shell git describe --tags --always --dirty)
BRANCH       ?= $(shell git branch | grep \* | cut -d ' ' -f2)
GITCOMMIT    ?= $(shell git rev-parse HEAD)
GITTREESTATE ?= $(if $(shell git status --porcelain),dirty,clean)
BUILDDATE    ?= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
appVersion   ?= $(VERSION)

# Available cpus for compiling, please refer to https://github.com/caicloud/engineering/issues/8186#issuecomment-518656946 for more information.
CPUS ?= $(shell /bin/bash hack/read_cpus_available.sh)

# Track code version with Docker Label.
DOCKER_LABELS ?= git-describe="$(shell date -u +v%Y%m%d)-$(shell git describe --tags --always --dirty)"

# Golang standard bin directory.
GOPATH ?= $(shell go env GOPATH)
BIN_DIR := $(GOPATH)/bin
GOLANGCI_LINT := $(BIN_DIR)/golangci-lint

# Default golang flags used in build and test
# -mod=vendor: force go to use the vendor files instead of using the `$GOPATH/pkg/mod`
# -p: the number of programs that can be run in parallel
# -count: run each test and benchmark 1 times. Set this flag to disable test cache
export GOFLAGS ?= -mod=vendor -p=$(CPUS) -count=1

#
# Define all targets. At least the following commands are required:
#

# All targets.
.PHONY: lint test build container push

build: build-local

# more info about `GOGC` env: https://github.com/golangci/golangci-lint#memory-usage-of-golangci-lint
lint: $(GOLANGCI_LINT)
	@$(GOLANGCI_LINT) run --timeout=10m0s

$(GOLANGCI_LINT):
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(BIN_DIR) v1.23.6

test:
	@go test -race -coverprofile=coverage.out ./...
	@go tool cover -func coverage.out | tail -n 1 | awk '{ print "Total coverage: " $$3 }'

build-local:
	@go build -v -o $(OUTPUT_DIR)/$(NAME)                                                                                         \
	  -ldflags "-s -w -X $(GOCOMMON)/version.module=$(NAME)                                                                       \
	    -X $(GOCOMMON)/version.version=$(VERSION)                                                                                 \
	    -X $(GOCOMMON)/version.branch=$(BRANCH)                                                                                   \
	    -X $(GOCOMMON)/version.gitCommit=$(GITCOMMIT)                                                                             \
	    -X $(GOCOMMON)/version.gitTreeState=$(GITTREESTATE)                                                                       \
	    -X $(GOCOMMON)/version.buildDate=$(BUILDDATE)"                                                                            \
	  $(CMD_DIR);

build-linux:
	@docker run --rm -it                                                                                                          \
	  -v $(PWD):/go/src/$(ROOT)                                                                                                   \
	  -w /go/src/$(ROOT)                                                                                                          \
	  -e GOOS=linux                                                                                                               \
	  -e GOARCH=amd64                                                                                                             \
	  -e GOPATH=/go                                                                                                               \
	  -e GOFLAGS="$(GOFLAGS)"                                                                                                     \
	  -e SHELLOPTS="$(SHELLOPTS)"                                                                                                 \
	  $(BASE_REGISTRY)/golang:1.13-security                                                                                       \
	    /bin/bash -c 'go run vendor/github.com/caicloud/nirvana/cmd/nirvana/main.go api --output=./apis --serve="" pkg/server &&  \
	      go build -v -o $(OUTPUT_DIR)/$(NAME)                                                                                    \
	        -ldflags "-s -w -X $(GOCOMMON)/version.module=$(NAME)                                                                 \
	          -X $(GOCOMMON)/version.version=$(VERSION)                                                                           \
	          -X $(GOCOMMON)/version.branch=$(BRANCH)                                                                             \
	          -X $(GOCOMMON)/version.gitCommit=$(GITCOMMIT)                                                                       \
	          -X $(GOCOMMON)/version.gitTreeState=$(GITTREESTATE)                                                                 \
	          -X $(GOCOMMON)/version.buildDate=$(BUILDDATE)"                                                                      \
	        $(CMD_DIR)'

container: build-linux
	@docker build -t $(REGISTRY)/$(IMAGE_NAME):$(VERSION)                                                                         \
	  --label $(DOCKER_LABELS)                                                                                                    \
	  -f $(BUILD_DIR)/Dockerfile .;

push: container
	@docker push $(REGISTRY)/$(IMAGE_NAME):$(VERSION);

.PHONY: clean
clean:
	@-rm -vrf ${OUTPUT_DIR}

.PHONY: dryrun
dryrun:
	@ginkgo -- -ginkgo.dryRun