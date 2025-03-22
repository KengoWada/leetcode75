.PHONY: test
test:
	@go test -v ./...

.PHONY: test-cov
test-cov:
	@go test -v -count=1 -coverprofile=coverage.out/cov.out ./... && go tool cover -html=coverage.out/cov.out -o=coverage.out/index.html
