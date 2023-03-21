# nopermission

[![pkg.go.dev][gopkg-badge]][gopkg]

`nopermission` finds types and interfaces with no @hasPermission directive in your GraphQL schema files.

```graphql
type ObjectWithoutPermission { # want "ObjectWithoutPermission has no hasPermission directive"
    id: ID!
    field: String!
}
```

## How to use

A runnable linter can be created with multichecker package.
You can create own linter with your favorite Analyzers.

```go
package main

import (
	"flag"
	"github.com/builtbystack/nopermission"
	"github.com/gqlgo/gqlanalysis/multichecker"
)

func main() {
	var excludes string
	flag.StringVar(&excludes, "excludes", "", "exclude GraphQL types for node check. it can specify multiple values separated by `,` and it can use regex(e.g .+Connection")
	flag.Parse()

	analyzer := nopermission.Analyzer(excludes)

	multichecker.Main(
		analyzer,
	)
}
```

`nopermission` provides a typical main function and you can install with `go install` command.

```sh
$ go install github.com/builtbystack/nopermission/cmd/nopermission@latest
```

The `nopermission` command has a flag, `schema` which will be parsed and analyzed by nopermission's Analyzer.

```sh
$ nopermission -schema="server/graphql/schema/**/*.graphql"
```

The default value of `schema` is "schema/*/**.graphql".

## Author

[![Appify Technologies, Inc.](appify-logo.png)](http://github.com/builtbystack)

<!-- links -->
[gopkg]: https://pkg.go.dev/github.com/builtbystack/nopermission
[gopkg-badge]: https://pkg.go.dev/badge/github.com/builtbystack/nopermission?status.svg
