package main

import (
	"fmt"
	"os"

	servicecmd "aws/generated/budgets/cmd"
)

func main() {
	if err := servicecmd.Execute(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
