package nopermission

import (
	"regexp"
	"strings"

	"github.com/gqlgo/gqlanalysis"
	"github.com/vektah/gqlparser/v2/ast"
)

func Analyzer(excludes string) *gqlanalysis.Analyzer {
	return &gqlanalysis.Analyzer{
		Name: "nopermission",
		Doc:  "nopermission finds id fields with no id directive.",
		Run:  run(excludes),
	}
}

func run(excludes string) func(pass *gqlanalysis.Pass) (interface{}, error) {
	return func(pass *gqlanalysis.Pass) (interface{}, error) {
		for _, t := range pass.Schema.Types {
			if t.BuiltIn {
				continue
			}
			exclude := false
			for _, rule := range strings.Split(excludes, ",") {
				if len(rule) > 0 {
					regex := regexp.MustCompile(rule)

					if regex.MatchString(t.Name) {
						exclude = true
					}
				}
			}
			if exclude {
				continue
			}
			if t.Kind == ast.Object {
				if t.Directives.ForName("hasPermission") == nil {
					pass.Reportf(t.Position, "%s has no hasPermission directive", t.Name)
				}
			}
			if t.Kind == ast.Interface {
				if t.Directives.ForName("hasPermission") == nil {
					pass.Reportf(t.Position, "%s has no hasPermission directive", t.Name)
				}
			}
		}

		return nil, nil
	}
}
