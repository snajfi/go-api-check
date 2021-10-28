package helpprinter

import (
	"fmt"
	"os"
)

func CheckHelpInvocation(args []string) {
	if args[0] == "-help" {
		printHelp()
		os.Exit(0)
	}
}

func printHelp() {
	fmt.Printf("This is help")
}
