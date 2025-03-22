build:
	go build -o bin ./...

lint:
	go tool github.com/golangci/golangci-lint/cmd/golangci-lint run

test:
	go test -race -timeout 1h -coverprofile cp.out ./...
