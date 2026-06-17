package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/ivschat/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"create-chat-token", "create-logging-configuration", "create-room", "delete-logging-configuration", "delete-message", "delete-room", "disconnect-user", "get-logging-configuration", "get-room", "list-logging-configurations", "list-rooms", "list-tags-for-resource", "send-event", "tag-resource", "untag-resource", "update-logging-configuration", "update-room"},
		OperationSet: map[string]bool{"create-chat-token": true, "create-logging-configuration": true, "create-room": true, "delete-logging-configuration": true, "delete-message": true, "delete-room": true, "disconnect-user": true, "get-logging-configuration": true, "get-room": true, "list-logging-configurations": true, "list-rooms": true, "list-tags-for-resource": true, "send-event": true, "tag-resource": true, "untag-resource": true, "update-logging-configuration": true, "update-room": true},
		OperationInputs: map[string][]string{
			"create-chat-token":            {"Attributes", "Capabilities", "RoomIdentifier", "SessionDurationInMinutes", "UserId"},
			"create-logging-configuration": {"DestinationConfiguration", "Name", "Tags"},
			"create-room":                  {"LoggingConfigurationIdentifiers", "MaximumMessageLength", "MaximumMessageRatePerSecond", "MessageReviewHandler", "Name", "Tags"},
			"delete-logging-configuration": {"Identifier"},
			"delete-message":               {"Id", "Reason", "RoomIdentifier"},
			"delete-room":                  {"Identifier"},
			"disconnect-user":              {"Reason", "RoomIdentifier", "UserId"},
			"get-logging-configuration":    {"Identifier"},
			"get-room":                     {"Identifier"},
			"list-logging-configurations":  {"MaxResults", "NextToken"},
			"list-rooms":                   {"LoggingConfigurationIdentifier", "MaxResults", "MessageReviewHandlerUri", "Name", "NextToken"},
			"list-tags-for-resource":       {"ResourceArn"},
			"send-event":                   {"Attributes", "EventName", "RoomIdentifier"},
			"tag-resource":                 {"ResourceArn", "Tags"},
			"untag-resource":               {"ResourceArn", "TagKeys"},
			"update-logging-configuration": {"DestinationConfiguration", "Identifier", "Name"},
			"update-room":                  {"Identifier", "LoggingConfigurationIdentifiers", "MaximumMessageLength", "MaximumMessageRatePerSecond", "MessageReviewHandler", "Name"},
		},
		OperationInputTypes: map[string]map[string]string{
			"create-chat-token":            {"Attributes": "map[string]string", "Capabilities": "[]types.ChatTokenCapability", "RoomIdentifier": "*string", "SessionDurationInMinutes": "*int32", "UserId": "*string"},
			"create-logging-configuration": {"DestinationConfiguration": "types.DestinationConfiguration", "Name": "*string", "Tags": "map[string]string"},
			"create-room":                  {"LoggingConfigurationIdentifiers": "[]string", "MaximumMessageLength": "*int32", "MaximumMessageRatePerSecond": "*int32", "MessageReviewHandler": "*types.MessageReviewHandler", "Name": "*string", "Tags": "map[string]string"},
			"delete-logging-configuration": {"Identifier": "*string"},
			"delete-message":               {"Id": "*string", "Reason": "*string", "RoomIdentifier": "*string"},
			"delete-room":                  {"Identifier": "*string"},
			"disconnect-user":              {"Reason": "*string", "RoomIdentifier": "*string", "UserId": "*string"},
			"get-logging-configuration":    {"Identifier": "*string"},
			"get-room":                     {"Identifier": "*string"},
			"list-logging-configurations":  {"MaxResults": "*int32", "NextToken": "*string"},
			"list-rooms":                   {"LoggingConfigurationIdentifier": "*string", "MaxResults": "*int32", "MessageReviewHandlerUri": "*string", "Name": "*string", "NextToken": "*string"},
			"list-tags-for-resource":       {"ResourceArn": "*string"},
			"send-event":                   {"Attributes": "map[string]string", "EventName": "*string", "RoomIdentifier": "*string"},
			"tag-resource":                 {"ResourceArn": "*string", "Tags": "map[string]string"},
			"untag-resource":               {"ResourceArn": "*string", "TagKeys": "[]string"},
			"update-logging-configuration": {"DestinationConfiguration": "types.DestinationConfiguration", "Identifier": "*string", "Name": "*string"},
			"update-room":                  {"Identifier": "*string", "LoggingConfigurationIdentifiers": "[]string", "MaximumMessageLength": "*int32", "MaximumMessageRatePerSecond": "*int32", "MessageReviewHandler": "*types.MessageReviewHandler", "Name": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"create-chat-token":            {"RoomIdentifier", "UserId"},
			"create-logging-configuration": {"DestinationConfiguration"},
			"create-room":                  {},
			"delete-logging-configuration": {"Identifier"},
			"delete-message":               {"Id", "RoomIdentifier"},
			"delete-room":                  {"Identifier"},
			"disconnect-user":              {"RoomIdentifier", "UserId"},
			"get-logging-configuration":    {"Identifier"},
			"get-room":                     {"Identifier"},
			"list-logging-configurations":  {},
			"list-rooms":                   {},
			"list-tags-for-resource":       {"ResourceArn"},
			"send-event":                   {"EventName", "RoomIdentifier"},
			"tag-resource":                 {"ResourceArn", "Tags"},
			"untag-resource":               {"ResourceArn", "TagKeys"},
			"update-logging-configuration": {"Identifier"},
			"update-room":                  {"Identifier"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("ivschat", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
