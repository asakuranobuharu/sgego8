package main

import (
	"github.com/asakuranobuharu/sgego8"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(sgego8.Analyzer) }
