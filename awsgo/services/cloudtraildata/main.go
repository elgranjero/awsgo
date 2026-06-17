package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/cloudtraildata/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"put-audit-events"},
		OperationSet: map[string]bool{"put-audit-events": true},
		OperationInputs: map[string][]string{
			"put-audit-events": {"AuditEvents", "ChannelArn", "ExternalId"},
		},
		OperationInputTypes: map[string]map[string]string{
			"put-audit-events": {"AuditEvents": "[]types.AuditEvent", "ChannelArn": "*string", "ExternalId": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"put-audit-events": {"AuditEvents", "ChannelArn"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("cloudtraildata", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
