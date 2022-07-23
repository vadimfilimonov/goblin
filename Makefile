build:
	go build -v ./...

verify:
	go mod verify

lint:
	go vet ./...

format:
	gofmt -w .

test:
	go test ./...

docs:
	godoc .
