package main

import (
	inputprocessor "go-api-check/cmd/input-processor"
	"os"
)

func main() {

	args := os.Args[1:]

	inputprocessor.ProcessInputArguments(args)

}
