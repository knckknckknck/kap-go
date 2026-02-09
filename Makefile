.PHONY: test lint fmt vet check

test:
	go test -race -v ./...

lint:
	golangci-lint run ./...

fmt:
	gofmt -l -w .

vet:
	go vet ./...

check: fmt vet test
