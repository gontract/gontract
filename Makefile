
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
