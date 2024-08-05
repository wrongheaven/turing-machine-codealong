build:
	@go build -o bin/turing-machine

run: build
	@./bin/turing-machine

test:
	@go test -v ./...
