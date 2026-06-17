package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/route53recoverycluster/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"get-routing-control-state", "list-routing-controls", "update-routing-control-state", "update-routing-control-states"},
		OperationSet: map[string]bool{"get-routing-control-state": true, "list-routing-controls": true, "update-routing-control-state": true, "update-routing-control-states": true},
		OperationInputs: map[string][]string{
			"get-routing-control-state":     {"RoutingControlArn"},
			"list-routing-controls":         {"ControlPanelArn", "MaxResults", "NextToken"},
			"update-routing-control-state":  {"RoutingControlArn", "RoutingControlState", "SafetyRulesToOverride"},
			"update-routing-control-states": {"SafetyRulesToOverride", "UpdateRoutingControlStateEntries"},
		},
		OperationInputTypes: map[string]map[string]string{
			"get-routing-control-state":     {"RoutingControlArn": "*string"},
			"list-routing-controls":         {"ControlPanelArn": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"update-routing-control-state":  {"RoutingControlArn": "*string", "RoutingControlState": "types.RoutingControlState", "SafetyRulesToOverride": "[]string"},
			"update-routing-control-states": {"SafetyRulesToOverride": "[]string", "UpdateRoutingControlStateEntries": "[]types.UpdateRoutingControlStateEntry"},
		},
		OperationInputRequired: map[string][]string{
			"get-routing-control-state":     {"RoutingControlArn"},
			"list-routing-controls":         {},
			"update-routing-control-state":  {"RoutingControlArn", "RoutingControlState"},
			"update-routing-control-states": {"UpdateRoutingControlStateEntries"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("route53recoverycluster", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
