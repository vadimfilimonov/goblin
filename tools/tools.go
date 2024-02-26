//go:generate bash -c "go build -o ../bin/golangci-lint github.com/golangci/golangci-lint/cmd/golangci-lint"
package tools

import _ "github.com/golangci/golangci-lint/cmd/golangci-lint"
