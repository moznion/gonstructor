clean:
	rm -rf bin/gonstructor*
build4test:
	go build -o bin/gonstructor cmd/gonstructor/gonstructor.go
gen4test: build4test
	go generate ./...
test: gen4test
	go test ./...
