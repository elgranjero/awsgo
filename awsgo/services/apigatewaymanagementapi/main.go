package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/apigatewaymanagementapi/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"delete-connection", "get-connection", "post-to-connection"},
		OperationSet: map[string]bool{"delete-connection": true, "get-connection": true, "post-to-connection": true},
		OperationInputs: map[string][]string{
			"delete-connection":  {"ConnectionId"},
			"get-connection":     {"ConnectionId"},
			"post-to-connection": {"ConnectionId", "Data"},
		},
		OperationInputTypes: map[string]map[string]string{
			"delete-connection":  {"ConnectionId": "*string"},
			"get-connection":     {"ConnectionId": "*string"},
			"post-to-connection": {"ConnectionId": "*string", "Data": "[]byte"},
		},
		OperationInputRequired: map[string][]string{
			"delete-connection":  {"ConnectionId"},
			"get-connection":     {"ConnectionId"},
			"post-to-connection": {"ConnectionId", "Data"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("apigatewaymanagementapi", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
