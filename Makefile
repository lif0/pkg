.PHONY: tidy test cover vet lint lint-fix sec check

tidy:
	go mod tidy

test:
	go test ./... -race

cover:
	go test ./... -race -covermode=atomic -coverprofile=coverage.out
	go tool cover -func=coverage.out

vet:
	go vet ./...

# golangci-lint v2 — install: https://golangci-lint.run/welcome/install/
lint:
	golangci-lint run ./...

# auto-fix what golangci-lint can
lint-fix:
	golangci-lint run --fix ./...

# gosec — install: go install github.com/securego/gosec/v2/cmd/gosec@latest
sec:
	gosec -quiet ./...

# everything CI runs, locally
check: vet lint sec test
