package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/sagemakera2iruntime/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"delete-human-loop", "describe-human-loop", "list-human-loops", "start-human-loop", "stop-human-loop"},
		OperationSet: map[string]bool{"delete-human-loop": true, "describe-human-loop": true, "list-human-loops": true, "start-human-loop": true, "stop-human-loop": true},
		OperationInputs: map[string][]string{
			"delete-human-loop":   {"HumanLoopName"},
			"describe-human-loop": {"HumanLoopName"},
			"list-human-loops":    {"CreationTimeAfter", "CreationTimeBefore", "FlowDefinitionArn", "MaxResults", "NextToken", "SortOrder"},
			"start-human-loop":    {"DataAttributes", "FlowDefinitionArn", "HumanLoopInput", "HumanLoopName"},
			"stop-human-loop":     {"HumanLoopName"},
		},
		OperationInputTypes: map[string]map[string]string{
			"delete-human-loop":   {"HumanLoopName": "*string"},
			"describe-human-loop": {"HumanLoopName": "*string"},
			"list-human-loops":    {"CreationTimeAfter": "*time.Time", "CreationTimeBefore": "*time.Time", "FlowDefinitionArn": "*string", "MaxResults": "*int32", "NextToken": "*string", "SortOrder": "types.SortOrder"},
			"start-human-loop":    {"DataAttributes": "*types.HumanLoopDataAttributes", "FlowDefinitionArn": "*string", "HumanLoopInput": "*types.HumanLoopInput", "HumanLoopName": "*string"},
			"stop-human-loop":     {"HumanLoopName": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"delete-human-loop":   {"HumanLoopName"},
			"describe-human-loop": {"HumanLoopName"},
			"list-human-loops":    {"FlowDefinitionArn"},
			"start-human-loop":    {"FlowDefinitionArn", "HumanLoopInput", "HumanLoopName"},
			"stop-human-loop":     {"HumanLoopName"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("sagemakera2iruntime", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
