package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/workmailmessageflow/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"get-raw-message-content", "put-raw-message-content"},
		OperationSet: map[string]bool{"get-raw-message-content": true, "put-raw-message-content": true},
		OperationInputs: map[string][]string{
			"get-raw-message-content": {"MessageId"},
			"put-raw-message-content": {"Content", "MessageId"},
		},
		OperationInputTypes: map[string]map[string]string{
			"get-raw-message-content": {"MessageId": "*string"},
			"put-raw-message-content": {"Content": "*types.RawMessageContent", "MessageId": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"get-raw-message-content": {"MessageId"},
			"put-raw-message-content": {"Content", "MessageId"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("workmailmessageflow", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
