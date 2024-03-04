//go:generate bash -c "go build -o ../bin/gocyclo github.com/fzipp/gocyclo/cmd/gocyclo"
//go:generate bash -c "go build -o ../bin/golangci-lint github.com/golangci/golangci-lint/cmd/golangci-lint"
//go:generate bash -c "go build -o ../bin/goconst github.com/jgautheron/goconst/cmd/goconst"
//go:generate bash -c "go build -o ../bin/errcheck github.com/kisielk/errcheck"
package tools

import (
	_ "github.com/fzipp/gocyclo/cmd/gocyclo"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/jgautheron/goconst/cmd/goconst"
	_ "github.com/kisielk/errcheck/errcheck"
)
