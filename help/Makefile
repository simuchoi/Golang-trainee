.PHONY: help fmt imports vet lint test race vuln tidy check ci

help:
	@echo "Available targets:"
	@echo "  make fmt      - run gofmt for all packages"
	@echo "  make imports  - run goimports and fix imports"
	@echo "  make vet      - run go vet"
	@echo "  make lint     - run golangci-lint (if installed), fallback to staticcheck (if installed)"
	@echo "  make test     - run tests"
	@echo "  make race     - run tests with race detector"
	@echo "  make vuln     - run govulncheck (if installed)"
	@echo "  make tidy     - run go mod tidy"
	@echo "  make check    - fmt + imports + vet + lint + test"
	@echo "  make ci       - tidy + check + race"

fmt:
	go fmt ./...

imports:
	goimports -w .

vet:
	go vet ./...

lint:
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	elif command -v staticcheck >/dev/null 2>&1; then \
		staticcheck ./...; \
	else \
		echo "No linter found. Install one:"; \
		echo "  go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
		echo "  or"; \
		echo "  go install honnef.co/go/tools/cmd/staticcheck@latest"; \
	fi

test:
	go test ./...

race:
	go test -race ./...

vuln:
	@if command -v govulncheck >/dev/null 2>&1; then \
		govulncheck ./...; \
	else \
		echo "govulncheck is not installed."; \
		echo "Install it with:"; \
		echo "  go install golang.org/x/vuln/cmd/govulncheck@latest"; \
	fi

tidy:
	go mod tidy

check: fmt imports vet lint test

ci: tidy check race
