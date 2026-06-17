package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/arczonalshift/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"cancel-practice-run", "cancel-zonal-shift", "create-practice-run-configuration", "delete-practice-run-configuration", "get-autoshift-observer-notification-status", "get-managed-resource", "list-autoshifts", "list-managed-resources", "list-zonal-shifts", "start-practice-run", "start-zonal-shift", "update-autoshift-observer-notification-status", "update-practice-run-configuration", "update-zonal-autoshift-configuration", "update-zonal-shift"},
		OperationSet: map[string]bool{"cancel-practice-run": true, "cancel-zonal-shift": true, "create-practice-run-configuration": true, "delete-practice-run-configuration": true, "get-autoshift-observer-notification-status": true, "get-managed-resource": true, "list-autoshifts": true, "list-managed-resources": true, "list-zonal-shifts": true, "start-practice-run": true, "start-zonal-shift": true, "update-autoshift-observer-notification-status": true, "update-practice-run-configuration": true, "update-zonal-autoshift-configuration": true, "update-zonal-shift": true},
		OperationInputs: map[string][]string{
			"cancel-practice-run":                           {"ZonalShiftId"},
			"cancel-zonal-shift":                            {"ZonalShiftId"},
			"create-practice-run-configuration":             {"AllowedWindows", "BlockedDates", "BlockedWindows", "BlockingAlarms", "OutcomeAlarms", "ResourceIdentifier"},
			"delete-practice-run-configuration":             {"ResourceIdentifier"},
			"get-autoshift-observer-notification-status":    {},
			"get-managed-resource":                          {"ResourceIdentifier"},
			"list-autoshifts":                               {"MaxResults", "NextToken", "Status"},
			"list-managed-resources":                        {"MaxResults", "NextToken"},
			"list-zonal-shifts":                             {"MaxResults", "NextToken", "ResourceIdentifier", "Status"},
			"start-practice-run":                            {"AwayFrom", "Comment", "ResourceIdentifier"},
			"start-zonal-shift":                             {"AwayFrom", "Comment", "ExpiresIn", "ResourceIdentifier"},
			"update-autoshift-observer-notification-status": {"Status"},
			"update-practice-run-configuration":             {"AllowedWindows", "BlockedDates", "BlockedWindows", "BlockingAlarms", "OutcomeAlarms", "ResourceIdentifier"},
			"update-zonal-autoshift-configuration":          {"ResourceIdentifier", "ZonalAutoshiftStatus"},
			"update-zonal-shift":                            {"Comment", "ExpiresIn", "ZonalShiftId"},
		},
		OperationInputTypes: map[string]map[string]string{
			"cancel-practice-run":                           {"ZonalShiftId": "*string"},
			"cancel-zonal-shift":                            {"ZonalShiftId": "*string"},
			"create-practice-run-configuration":             {"AllowedWindows": "[]string", "BlockedDates": "[]string", "BlockedWindows": "[]string", "BlockingAlarms": "[]types.ControlCondition", "OutcomeAlarms": "[]types.ControlCondition", "ResourceIdentifier": "*string"},
			"delete-practice-run-configuration":             {"ResourceIdentifier": "*string"},
			"get-autoshift-observer-notification-status":    {},
			"get-managed-resource":                          {"ResourceIdentifier": "*string"},
			"list-autoshifts":                               {"MaxResults": "*int32", "NextToken": "*string", "Status": "types.AutoshiftExecutionStatus"},
			"list-managed-resources":                        {"MaxResults": "*int32", "NextToken": "*string"},
			"list-zonal-shifts":                             {"MaxResults": "*int32", "NextToken": "*string", "ResourceIdentifier": "*string", "Status": "types.ZonalShiftStatus"},
			"start-practice-run":                            {"AwayFrom": "*string", "Comment": "*string", "ResourceIdentifier": "*string"},
			"start-zonal-shift":                             {"AwayFrom": "*string", "Comment": "*string", "ExpiresIn": "*string", "ResourceIdentifier": "*string"},
			"update-autoshift-observer-notification-status": {"Status": "types.AutoshiftObserverNotificationStatus"},
			"update-practice-run-configuration":             {"AllowedWindows": "[]string", "BlockedDates": "[]string", "BlockedWindows": "[]string", "BlockingAlarms": "[]types.ControlCondition", "OutcomeAlarms": "[]types.ControlCondition", "ResourceIdentifier": "*string"},
			"update-zonal-autoshift-configuration":          {"ResourceIdentifier": "*string", "ZonalAutoshiftStatus": "types.ZonalAutoshiftStatus"},
			"update-zonal-shift":                            {"Comment": "*string", "ExpiresIn": "*string", "ZonalShiftId": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"cancel-practice-run":                           {"ZonalShiftId"},
			"cancel-zonal-shift":                            {"ZonalShiftId"},
			"create-practice-run-configuration":             {"OutcomeAlarms", "ResourceIdentifier"},
			"delete-practice-run-configuration":             {"ResourceIdentifier"},
			"get-autoshift-observer-notification-status":    {},
			"get-managed-resource":                          {"ResourceIdentifier"},
			"list-autoshifts":                               {},
			"list-managed-resources":                        {},
			"list-zonal-shifts":                             {},
			"start-practice-run":                            {"AwayFrom", "Comment", "ResourceIdentifier"},
			"start-zonal-shift":                             {"AwayFrom", "Comment", "ExpiresIn", "ResourceIdentifier"},
			"update-autoshift-observer-notification-status": {"Status"},
			"update-practice-run-configuration":             {"ResourceIdentifier"},
			"update-zonal-autoshift-configuration":          {"ResourceIdentifier", "ZonalAutoshiftStatus"},
			"update-zonal-shift":                            {"ZonalShiftId"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("arczonalshift", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
