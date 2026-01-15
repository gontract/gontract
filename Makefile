

GOLANGCI_LINT := go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest

.PHONY: go.vet
go.vet: vet
.PHONY: vet
vet:
		@go vet ./...

.PHONY: go.fmt
go.fmt: fmt

.PHONY: fmt
fmt:
	@ go fmt ./...

.PHONY: test
test:
	@go test ./...

.PHONY: golangci-lint
golangci-lint:
	@$(GOLANGCI_LINT) run

.PHONY: check
check: vet fmt golangci-lint
