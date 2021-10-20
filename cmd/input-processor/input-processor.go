package inputprocessor

import (
	"fmt"
	helpprinter "go-api-check/cmd/help-printer"
)

func ProcessInputArguments(args []string) {

	if len(args) == 0 {
		fmt.Println("No parameters. Use param '-help' for basic usage.")
	} else {
		switch args[0] {
		case "-help":
			helpprinter.CheckHelpInvocation(args)
		default:
			fmt.Printf("Parameter %s was not recongized.", args[0])
		}
	}

}
