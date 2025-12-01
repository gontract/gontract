
.PHONY: go.vet
go.vet: vet
.PHONY: vet
vet:
		@go vet ./...

.PHONY: test
test:
	@go test ./...
