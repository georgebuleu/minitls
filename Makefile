lint:
	golangci-lint run ./...

format:
	goimports -w .
	golines .

security:
	gosec ./...
