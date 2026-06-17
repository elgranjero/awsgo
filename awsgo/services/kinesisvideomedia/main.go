package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/kinesisvideomedia/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"get-media"},
		OperationSet: map[string]bool{"get-media": true},
		OperationInputs: map[string][]string{
			"get-media": {"StartSelector", "StreamARN", "StreamName"},
		},
		OperationInputTypes: map[string]map[string]string{
			"get-media": {"StartSelector": "*types.StartSelector", "StreamARN": "*string", "StreamName": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"get-media": {"StartSelector"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("kinesisvideomedia", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
