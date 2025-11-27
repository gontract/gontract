
.PHONY: vet
vet:
		@go vet ./...

.PHONY: test
test:
	@go test ./...
