PKGS := $(shell go list ./...)
PKGS_WITHOUT_TEST := $(shell go list ./... | grep -v "gonstructor/internal/test")
GOVERSION=$(shell go version)
GOOS=$(word 1,$(subst /, ,$(lastword $(GOVERSION))))
GOARCH=$(word 2,$(subst /, ,$(lastword $(GOVERSION))))
RELEASE_DIR=bin
REVISION=$(shell git rev-parse --verify HEAD)
INTERNAL_PACKAGE=github.com/moznion/gonstructor/internal

check: test lint vet fmt-check

all: check clean build-linux-amd64 build-linux-386 build-darwin-amd64 build-darwin-386 build-windows-amd64 build-windows-386

build: $(RELEASE_DIR)/gonstructor_$(GOOS)_$(GOARCH)

build-linux-amd64:
	@$(MAKE) build GOOS=linux GOARCH=amd64

build-linux-386:
	@$(MAKE) build GOOS=linux GOARCH=386

build-darwin-amd64:
	@$(MAKE) build GOOS=darwin GOARCH=amd64

build-darwin-386:
	@$(MAKE) build GOOS=darwin GOARCH=386

build-windows-amd64:
	@$(MAKE) build GOOS=windows GOARCH=amd64

build-windows-386:
	@$(MAKE) build GOOS=windows GOARCH=386

$(RELEASE_DIR)/gonstructor_$(GOOS)_$(GOARCH):
ifndef VERSION
	@echo '[ERROR] $$VERSION must be specified'
	exit 255
endif
	go build -ldflags "-X $(INTERNAL_PACKAGE).rev=$(REVISION) -X $(INTERNAL_PACKAGE).ver=$(VERSION)" \
		-o $(RELEASE_DIR)/gonstructor_$(GOOS)_$(GOARCH)_$(VERSION) cmd/gonstructor/gonstructor.go

build4test: clean
	go build -ldflags "-X $(INTERNAL_PACKAGE).rev=$(REVISION) -X $(INTERNAL_PACKAGE).ver=TESTING" \
		-o $(RELEASE_DIR)/gonstructor_test cmd/gonstructor/gonstructor.go

gen4test: build4test
	rm -f internal/test/*_gen.go
	go generate $(PKGS)

test: gen4test
	go test -v $(PKGS)

lint:
	golint -set_exit_status $(PKGS_WITHOUT_TEST)

vet:
	go vet $(PKGS)

fmt-check:
	gofmt -l -s **/*.go | grep [^*][.]go$$; \
	EXIT_CODE=$$?; \
	if [ $$EXIT_CODE -eq 0 ]; then exit 1; fi; \
	goimports -l **/*.go | grep [^*][.]go$$; \
	EXIT_CODE=$$?; \
	if [ $$EXIT_CODE -eq 0 ]; then exit 1; fi \

fmt:
	gofmt -w -s **/*.go
	goimports -w **/*.go

clean:
	rm -rf bin/gonstructor*

