tidy:
	go mod tidy

test:
	go test ./... -race

cover:
	go test ./... -race -covermode=atomic -coverprofile=coverage.out
	go tool cover -func=coverage.out
