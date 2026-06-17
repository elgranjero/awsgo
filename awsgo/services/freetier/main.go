package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/freetier/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"get-account-activity", "get-account-plan-state", "get-free-tier-usage", "list-account-activities", "upgrade-account-plan"},
		OperationSet: map[string]bool{"get-account-activity": true, "get-account-plan-state": true, "get-free-tier-usage": true, "list-account-activities": true, "upgrade-account-plan": true},
		OperationInputs: map[string][]string{
			"get-account-activity":    {"ActivityId", "LanguageCode"},
			"get-account-plan-state":  {},
			"get-free-tier-usage":     {"Filter", "MaxResults", "NextToken"},
			"list-account-activities": {"FilterActivityStatuses", "LanguageCode", "MaxResults", "NextToken"},
			"upgrade-account-plan":    {"AccountPlanType"},
		},
		OperationInputTypes: map[string]map[string]string{
			"get-account-activity":    {"ActivityId": "*string", "LanguageCode": "types.LanguageCode"},
			"get-account-plan-state":  {},
			"get-free-tier-usage":     {"Filter": "*types.Expression", "MaxResults": "*int32", "NextToken": "*string"},
			"list-account-activities": {"FilterActivityStatuses": "[]types.ActivityStatus", "LanguageCode": "types.LanguageCode", "MaxResults": "*int32", "NextToken": "*string"},
			"upgrade-account-plan":    {"AccountPlanType": "types.AccountPlanType"},
		},
		OperationInputRequired: map[string][]string{
			"get-account-activity":    {"ActivityId"},
			"get-account-plan-state":  {},
			"get-free-tier-usage":     {},
			"list-account-activities": {},
			"upgrade-account-plan":    {"AccountPlanType"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("freetier", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
