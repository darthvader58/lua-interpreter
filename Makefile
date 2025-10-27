.PHONY: build run test clean repl help

# Build the interpreter
build:
	go build -o lua-interpreter .

# Run the REPL
run: build
	./lua-interpreter

# Run the REPL without building binary
repl:
	go run main.go

# Run all tests
test:
	go test ./... -v

# Run tests with coverage
test-coverage:
	go test ./... -cover

# Run specific package tests
test-lexer:
	go test ./lexer -v

# Clean build artifacts
clean:
	rm -f lua-interpreter
	go clean

# Format code
fmt:
	go fmt ./...

# Run linter (requires golangci-lint)
lint:
	golangci-lint run

# Display help
help:
	@echo "Available targets:"
	@echo "  build          - Build the interpreter binary"
	@echo "  run            - Build and run the REPL"
	@echo "  repl           - Run the REPL without building binary"
	@echo "  test           - Run all tests"
	@echo "  test-coverage  - Run tests with coverage report"
	@echo "  test-lexer     - Run lexer tests only"
	@echo "  clean          - Remove build artifacts"
	@echo "  fmt            - Format source code"
	@echo "  lint           - Run linter (requires golangci-lint)"
	@echo "  help           - Show this help message"