build:
	go build -v ./...

verify:
	go mod verify

vet:
	go vet ./...

test:
	go test ./...
