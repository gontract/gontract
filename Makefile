

GOLANGCI_LINT := go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest


.PHONY: check.go.vet
check.go.vet:
	@echo "vetting go code..."
	@go vet ./...


.PHOHY: check.go.fmt
check.go.fmt:
	@echo "Checking go formatting..."
	@if [ -n "$$(gofmt -l .)" ]; then \
			echo "Files need formatting:"; \
				gofmt -l .; \
				exit 1; \
	else \
		echo "All files formatted correctly."; \
	fi

.PHONY: fix.go.fmt
fix.go.fmt: #␣fix␣go␣formatting (if needed)
	@ go fmt ./...

.PHONY: test
test: check
	@go test ./...

.PHONY: golangci-lint
golangci-lint:
	@echo "linting go code ..."
	@$(GOLANGCI_LINT) run


.PHONY: check.go.lint
check.go.lint: check.go.vet golangci-lint

.PHONY: check
check: check.go.fmt check.go.lint
