package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/pinpointsmsvoice/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"create-configuration-set", "create-configuration-set-event-destination", "delete-configuration-set", "delete-configuration-set-event-destination", "get-configuration-set-event-destinations", "list-configuration-sets", "send-voice-message", "update-configuration-set-event-destination"},
		OperationSet: map[string]bool{"create-configuration-set": true, "create-configuration-set-event-destination": true, "delete-configuration-set": true, "delete-configuration-set-event-destination": true, "get-configuration-set-event-destinations": true, "list-configuration-sets": true, "send-voice-message": true, "update-configuration-set-event-destination": true},
		OperationInputs: map[string][]string{
			"create-configuration-set":                   {"ConfigurationSetName"},
			"create-configuration-set-event-destination": {"ConfigurationSetName", "EventDestination", "EventDestinationName"},
			"delete-configuration-set":                   {"ConfigurationSetName"},
			"delete-configuration-set-event-destination": {"ConfigurationSetName", "EventDestinationName"},
			"get-configuration-set-event-destinations":   {"ConfigurationSetName"},
			"list-configuration-sets":                    {"NextToken", "PageSize"},
			"send-voice-message":                         {"CallerId", "ConfigurationSetName", "Content", "DestinationPhoneNumber", "OriginationPhoneNumber"},
			"update-configuration-set-event-destination": {"ConfigurationSetName", "EventDestination", "EventDestinationName"},
		},
		OperationInputTypes: map[string]map[string]string{
			"create-configuration-set":                   {"ConfigurationSetName": "*string"},
			"create-configuration-set-event-destination": {"ConfigurationSetName": "*string", "EventDestination": "*types.EventDestinationDefinition", "EventDestinationName": "*string"},
			"delete-configuration-set":                   {"ConfigurationSetName": "*string"},
			"delete-configuration-set-event-destination": {"ConfigurationSetName": "*string", "EventDestinationName": "*string"},
			"get-configuration-set-event-destinations":   {"ConfigurationSetName": "*string"},
			"list-configuration-sets":                    {"NextToken": "*string", "PageSize": "*string"},
			"send-voice-message":                         {"CallerId": "*string", "ConfigurationSetName": "*string", "Content": "*types.VoiceMessageContent", "DestinationPhoneNumber": "*string", "OriginationPhoneNumber": "*string"},
			"update-configuration-set-event-destination": {"ConfigurationSetName": "*string", "EventDestination": "*types.EventDestinationDefinition", "EventDestinationName": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"create-configuration-set":                   {},
			"create-configuration-set-event-destination": {"ConfigurationSetName"},
			"delete-configuration-set":                   {"ConfigurationSetName"},
			"delete-configuration-set-event-destination": {"ConfigurationSetName", "EventDestinationName"},
			"get-configuration-set-event-destinations":   {"ConfigurationSetName"},
			"list-configuration-sets":                    {},
			"send-voice-message":                         {},
			"update-configuration-set-event-destination": {"ConfigurationSetName", "EventDestinationName"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("pinpointsmsvoice", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
