package main

import (
	helpprinter "go-api-check/cmd/help-printer"
	"os"
)

func main() {

	args := os.Args[1:]

	helpprinter.CheckHelpInvocation(args)

}
