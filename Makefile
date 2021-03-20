VERSION      = v0.1.0
TARGETS     ?= builder fetcher preupgradechecks cli bundle
PKG          = github.com/fnlize/fnlize
BUILD_DIR    = ./build
IMAGE_PREFIX = $(strip )
IMAGE_SUFFIX = $(strip )
OUTPUT_DIR   = ./bin
CMD_DIR      = ./cmd

COMMIT := $(strip $(shell git rev-parse --short HEAD 2>/dev/null))
COMMIT := $(COMMIT)$(shell git diff-files --quiet || echo '-dirty')
COMMIT := $(if $(COMMIT),$(COMMIT),"Unknown")

.PHONY: check
check: test-run build

.PHONY: test-run
test-run:
	hack/verify-gofmt.sh
	hack/verify-govet.sh
	hack/verify-staticcheck.sh
	hack/runtests.sh
	@rm -f coverage.txt

.PHONY: build
build:
	@for target in $(TARGETS); do                                   \
	  go build -v -o $(OUTPUT_DIR)/$${target}                       \
	    -ldflags "-s -w -X $(PKG)/pkg/version.Version=$(VERSION)    \
	    -X $(PKG)/pkg/version.Commit=$(COMMIT)                      \
	    -X $(PKG)/pkg/version.Package=$(PKG)"                       \
	    $(CMD_DIR)/$${target};                                      \
	done

.PHONY: install
install: build
	$(INSTALL) $(OUTPUT_DIR)/cli $(GOPATH)/bin/fission

.PHONY: image
image:
	@for target in $(TARGETS); do                                     \
	  image=$(IMAGE_PREFIX)$${target}$(IMAGE_SUFFIX);                 \
	  echo Building $${image}:$(VERSION);                             \
	  docker build -t $${image}:$(VERSION)                            \
	    --build-arg PKG=${PKG}                                        \
	    --build-arg GITCOMMIT=${COMMIT}                               \
	    --build-arg BUILDVERSION=${COMMIT}                            \
	    --build-arg BUILDDATE=${COMMIT}                               \
	    --build-arg TARGET=$${target}                                 \
	    -f $(BUILD_DIR)/$${target}/Dockerfile .;                      \
	done
