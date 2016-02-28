package main

import (
	"flag"
	"os"

	"github.com/cweill/gotests/gotests/process"
)

var (
	onlyFuncs     = flag.String("only", "", `regexp. generate tests for functions and methods that match only. Takes precedence over -all`)
	exclFuncs     = flag.String("excl", "", `regexp. generate tests for functions and methods that don't match. Takes precedence over -only, -exported, and -all`)
	exportedFuncs = flag.Bool("exported", false, `generate tests for exported functions and methods. Takes precedence over -only and -all`)
	allFuncs      = flag.Bool("all", false, "generate tests for all functions and methods")
	printInputs   = flag.Bool("i", false, "print test inputs in error messages")
	writeOutput   = flag.Bool("w", false, "write output to (test) files instead of stdout")
)

func main() {
	flag.Parse()
	args := flag.Args()
	process.Run(os.Stdout, args, &process.Options{
		OnlyFuncs:     *onlyFuncs,
		ExclFuncs:     *exclFuncs,
		ExportedFuncs: *exportedFuncs,
		AllFuncs:      *allFuncs,
		PrintInputs:   *printInputs,
		WriteOutput:   *writeOutput,
	})
}
