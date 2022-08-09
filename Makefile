build:
	go build -v ./...

verify:
	go mod verify

lint:
	go vet ./...
	golangci-lint run

format:
	gofmt -w .

test:
	go test ./...

test-coverage:
	go test ./... -cover

build-test-coverage-report:
	go test . -coverprofile=coverage.out
	go tool cover -html=coverage.out
docs:
	godoc .
