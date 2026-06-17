package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/ssmguiconnect/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"delete-connection-recording-preferences", "get-connection-recording-preferences", "update-connection-recording-preferences"},
		OperationSet: map[string]bool{"delete-connection-recording-preferences": true, "get-connection-recording-preferences": true, "update-connection-recording-preferences": true},
		OperationInputs: map[string][]string{
			"delete-connection-recording-preferences": {"ClientToken"},
			"get-connection-recording-preferences":    {},
			"update-connection-recording-preferences": {"ClientToken", "ConnectionRecordingPreferences"},
		},
		OperationInputTypes: map[string]map[string]string{
			"delete-connection-recording-preferences": {"ClientToken": "*string"},
			"get-connection-recording-preferences":    {},
			"update-connection-recording-preferences": {"ClientToken": "*string", "ConnectionRecordingPreferences": "*types.ConnectionRecordingPreferences"},
		},
		OperationInputRequired: map[string][]string{
			"delete-connection-recording-preferences": {},
			"get-connection-recording-preferences":    {},
			"update-connection-recording-preferences": {"ConnectionRecordingPreferences"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("ssmguiconnect", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
