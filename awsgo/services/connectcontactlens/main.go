package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/connectcontactlens/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"list-realtime-contact-analysis-segments"},
		OperationSet: map[string]bool{"list-realtime-contact-analysis-segments": true},
		OperationInputs: map[string][]string{
			"list-realtime-contact-analysis-segments": {"ContactId", "InstanceId", "MaxResults", "NextToken"},
		},
		OperationInputTypes: map[string]map[string]string{
			"list-realtime-contact-analysis-segments": {"ContactId": "*string", "InstanceId": "*string", "MaxResults": "*int32", "NextToken": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"list-realtime-contact-analysis-segments": {"ContactId", "InstanceId"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("connectcontactlens", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
