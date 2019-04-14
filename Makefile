start:
	@echo ">>> start app"
	@./bin/main

build:
	@echo ">>> build foreign currency"
	@go get ./...
	@mkdir bin
	@go build -o bin/main
	@echo ">>> finish build foreign currency"

test:
	@echo ">>> test all package"
	@go test ./... -v -coverprofile=cover.out; go tool cover -html=cover.out -o cover.html; rm cover.out;
	@echo ">>> finish test all package"