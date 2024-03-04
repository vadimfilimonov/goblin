//go:generate bash -c "go build -o ../bin/gocyclo github.com/fzipp/gocyclo/cmd/gocyclo"
//go:generate bash -c "go build -o ../bin/golangci-lint github.com/golangci/golangci-lint/cmd/golangci-lint"
//go:generate bash -c "go build -o ../bin/goconst github.com/jgautheron/goconst/cmd/goconst"
package tools

import (
	_ "github.com/fzipp/gocyclo/cmd/gocyclo"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/jgautheron/goconst/cmd/goconst"
)
