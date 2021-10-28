package main

import (
	inputprocessor "go-api-check/cmd/input-processor"
	"os"
)

var verbose bool = false
var version = "undefined"

func main() {

	args := os.Args[1:]

	inputprocessor.ProcessInputArguments(args)

}
