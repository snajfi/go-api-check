package helpprinter

import "fmt"

func CheckHelpInvocation(args []string) {
	if len(args) == 1 && args[0] == "-help" {
		printHelp()
	}
}

func printHelp() {
	fmt.Printf("This is help")
}
