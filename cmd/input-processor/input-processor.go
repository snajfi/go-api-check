package inputprocessor

import (
	"fmt"
	"go-api-check/cmd/caller"
	helpprinter "go-api-check/cmd/help-printer"
	utils "go-api-check/internal/utils"
	"os"
)

func ProcessInputArguments(args []string) {

	if len(args) == 0 {
		fmt.Println("No parameters. Use param '-help' for basic usage.")
		os.Exit(0)
	}

	verbose := utils.Contains(args, "-v")
	if verbose {
		fmt.Println("Verbose mode was activated.")
	}

	if utils.Contains(args, "-help") {
		helpprinter.CheckHelpInvocation(args)
	}

	switch args[0] {
	case "-call":
		caller.Call(args)
	default:
		fmt.Printf("Parameter %s was not recongized.", args[0])
	}

}
