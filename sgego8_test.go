package sgego8_test

import (
	"github.com/asakuranobuharu/sgego8"
	"golang.org/x/tools/go/analysis/analysistest"
	"testing"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, sgego8.Analyzer, "c")
}
