build:
	@go build -o gator

run: build
	@./gator

test: 
	@go test -v ./...
