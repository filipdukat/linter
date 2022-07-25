package main

import (
	"github.com/filipdukat/linter/modules"
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {

	// Set slice of banned libraries
	libs := []string{
		"fmt",
	}

	multichecker.Main(modules.NewAnalyzer(libs))
}
