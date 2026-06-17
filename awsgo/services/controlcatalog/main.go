package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/controlcatalog/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"get-control", "list-common-controls", "list-control-mappings", "list-controls", "list-domains", "list-objectives"},
		OperationSet: map[string]bool{"get-control": true, "list-common-controls": true, "list-control-mappings": true, "list-controls": true, "list-domains": true, "list-objectives": true},
		OperationInputs: map[string][]string{
			"get-control":           {"ControlArn"},
			"list-common-controls":  {"CommonControlFilter", "MaxResults", "NextToken"},
			"list-control-mappings": {"Filter", "MaxResults", "NextToken"},
			"list-controls":         {"Filter", "MaxResults", "NextToken"},
			"list-domains":          {"MaxResults", "NextToken"},
			"list-objectives":       {"MaxResults", "NextToken", "ObjectiveFilter"},
		},
		OperationInputTypes: map[string]map[string]string{
			"get-control":           {"ControlArn": "*string"},
			"list-common-controls":  {"CommonControlFilter": "*types.CommonControlFilter", "MaxResults": "*int32", "NextToken": "*string"},
			"list-control-mappings": {"Filter": "*types.ControlMappingFilter", "MaxResults": "*int32", "NextToken": "*string"},
			"list-controls":         {"Filter": "*types.ControlFilter", "MaxResults": "*int32", "NextToken": "*string"},
			"list-domains":          {"MaxResults": "*int32", "NextToken": "*string"},
			"list-objectives":       {"MaxResults": "*int32", "NextToken": "*string", "ObjectiveFilter": "*types.ObjectiveFilter"},
		},
		OperationInputRequired: map[string][]string{
			"get-control":           {"ControlArn"},
			"list-common-controls":  {},
			"list-control-mappings": {},
			"list-controls":         {},
			"list-domains":          {},
			"list-objectives":       {},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("controlcatalog", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
