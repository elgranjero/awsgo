package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/appconfigdata/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"get-latest-configuration", "start-configuration-session"},
		OperationSet: map[string]bool{"get-latest-configuration": true, "start-configuration-session": true},
		OperationInputs: map[string][]string{
			"get-latest-configuration":    {"ConfigurationToken"},
			"start-configuration-session": {"ApplicationIdentifier", "ConfigurationProfileIdentifier", "EnvironmentIdentifier", "RequiredMinimumPollIntervalInSeconds"},
		},
		OperationInputTypes: map[string]map[string]string{
			"get-latest-configuration":    {"ConfigurationToken": "*string"},
			"start-configuration-session": {"ApplicationIdentifier": "*string", "ConfigurationProfileIdentifier": "*string", "EnvironmentIdentifier": "*string", "RequiredMinimumPollIntervalInSeconds": "*int32"},
		},
		OperationInputRequired: map[string][]string{
			"get-latest-configuration":    {"ConfigurationToken"},
			"start-configuration-session": {"ApplicationIdentifier", "ConfigurationProfileIdentifier", "EnvironmentIdentifier"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("appconfigdata", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
