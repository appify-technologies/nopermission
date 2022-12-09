package main

import (
	"flag"
	"github.com/appify-technologies/nopermission"
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
