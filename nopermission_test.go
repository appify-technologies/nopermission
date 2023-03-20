package nopermission_test

import (
	"testing"

	"github.com/appify-technologies/nopermission"
	"github.com/gqlgo/gqlanalysis/analysistest"
)

// Run analysis with no excludes rule.
func Test(t *testing.T) {
	testdata := analysistest.TestData(t)
	analysistest.Run(t, testdata, nopermission.Analyzer(""), "a")
}

// Run analysis with one excludes rule.
func TestWithSingleExclusion(t *testing.T) {
	testdata := analysistest.TestData(t)
	analysistest.Run(t, testdata, nopermission.Analyzer("ObjectWithoutPermission"), "b")
}

// Run analysis with two excludes rules.
func TestWithMultipleExclusions(t *testing.T) {
	testdata := analysistest.TestData(t)
	analysistest.Run(t, testdata, nopermission.Analyzer("ObjectWithoutPermission,InterfaceWithoutPermission"), "c")
}

// Run analysis with one excludes rule in regex.
func TestWithRegex(t *testing.T) {
	testdata := analysistest.TestData(t)
	analysistest.Run(t, testdata, nopermission.Analyzer("Object.+"), "d")
}
