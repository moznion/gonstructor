check: test lint vet fmt-check

PKGS := $(shell go list ./...)
PKGS_WITHOUT_TEST := $(shell go list ./... | grep -v "gonstructor/internal/test")

clean:
	rm -rf bin/gonstructor*

build4test: clean
	go build -o bin/gonstructor cmd/gonstructor/gonstructor.go

gen4test: build4test
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

