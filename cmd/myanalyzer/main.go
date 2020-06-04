package main

import (
	"github.com/QualiArts/myanalyzer"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(sgego8.Analyzer) }
