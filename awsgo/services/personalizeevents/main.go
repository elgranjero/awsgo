package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/personalizeevents/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"put-action-interactions", "put-actions", "put-events", "put-items", "put-users"},
		OperationSet: map[string]bool{"put-action-interactions": true, "put-actions": true, "put-events": true, "put-items": true, "put-users": true},
		OperationInputs: map[string][]string{
			"put-action-interactions": {"ActionInteractions", "TrackingId"},
			"put-actions":             {"Actions", "DatasetArn"},
			"put-events":              {"EventList", "SessionId", "TrackingId", "UserId"},
			"put-items":               {"DatasetArn", "Items"},
			"put-users":               {"DatasetArn", "Users"},
		},
		OperationInputTypes: map[string]map[string]string{
			"put-action-interactions": {"ActionInteractions": "[]types.ActionInteraction", "TrackingId": "*string"},
			"put-actions":             {"Actions": "[]types.Action", "DatasetArn": "*string"},
			"put-events":              {"EventList": "[]types.Event", "SessionId": "*string", "TrackingId": "*string", "UserId": "*string"},
			"put-items":               {"DatasetArn": "*string", "Items": "[]types.Item"},
			"put-users":               {"DatasetArn": "*string", "Users": "[]types.User"},
		},
		OperationInputRequired: map[string][]string{
			"put-action-interactions": {"ActionInteractions", "TrackingId"},
			"put-actions":             {"Actions", "DatasetArn"},
			"put-events":              {"EventList", "SessionId", "TrackingId"},
			"put-items":               {"DatasetArn", "Items"},
			"put-users":               {"DatasetArn", "Users"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("personalizeevents", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
