package sgego8_test

import (
	"testing"

	"github.com/QualiArts/myanalyzer"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, sgego8.Analyzer, "c")
}