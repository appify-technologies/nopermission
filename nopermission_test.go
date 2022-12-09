package nopermission_test

import (
	"testing"

	"github.com/appify-technologies/nopermission"
	"github.com/gqlgo/gqlanalysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData(t)
	analysistest.Run(t, testdata, nopermission.Analyzer(""), "a")
}

func TestWithSingleExclusion(t *testing.T) {
	testdata := analysistest.TestData(t)
	analysistest.Run(t, testdata, nopermission.Analyzer("ObjectWithoutPermission"), "b")
}

func TestWithMultipleExclusions(t *testing.T) {
	testdata := analysistest.TestData(t)
	analysistest.Run(t, testdata, nopermission.Analyzer("ObjectWithoutPermission,InterfaceWithoutPermission"), "c")
}

func TestWithRegex(t *testing.T) {
	testdata := analysistest.TestData(t)
	analysistest.Run(t, testdata, nopermission.Analyzer("Object.+"), "d")
}
