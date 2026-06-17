package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/iotdataplane/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"delete-connection", "delete-thing-shadow", "get-retained-message", "get-thing-shadow", "list-named-shadows-for-thing", "list-retained-messages", "publish", "update-thing-shadow"},
		OperationSet: map[string]bool{"delete-connection": true, "delete-thing-shadow": true, "get-retained-message": true, "get-thing-shadow": true, "list-named-shadows-for-thing": true, "list-retained-messages": true, "publish": true, "update-thing-shadow": true},
		OperationInputs: map[string][]string{
			"delete-connection":            {"CleanSession", "ClientId", "PreventWillMessage"},
			"delete-thing-shadow":          {"ShadowName", "ThingName"},
			"get-retained-message":         {"Topic"},
			"get-thing-shadow":             {"ShadowName", "ThingName"},
			"list-named-shadows-for-thing": {"NextToken", "PageSize", "ThingName"},
			"list-retained-messages":       {"MaxResults", "NextToken"},
			"publish":                      {"ContentType", "CorrelationData", "MessageExpiry", "Payload", "PayloadFormatIndicator", "Qos", "ResponseTopic", "Retain", "Topic", "UserProperties"},
			"update-thing-shadow":          {"Payload", "ShadowName", "ThingName"},
		},
		OperationInputTypes: map[string]map[string]string{
			"delete-connection":            {"CleanSession": "bool", "ClientId": "*string", "PreventWillMessage": "bool"},
			"delete-thing-shadow":          {"ShadowName": "*string", "ThingName": "*string"},
			"get-retained-message":         {"Topic": "*string"},
			"get-thing-shadow":             {"ShadowName": "*string", "ThingName": "*string"},
			"list-named-shadows-for-thing": {"NextToken": "*string", "PageSize": "*int32", "ThingName": "*string"},
			"list-retained-messages":       {"MaxResults": "*int32", "NextToken": "*string"},
			"publish":                      {"ContentType": "*string", "CorrelationData": "*string", "MessageExpiry": "int64", "Payload": "[]byte", "PayloadFormatIndicator": "types.PayloadFormatIndicator", "Qos": "int32", "ResponseTopic": "*string", "Retain": "bool", "Topic": "*string", "UserProperties": "*string"},
			"update-thing-shadow":          {"Payload": "[]byte", "ShadowName": "*string", "ThingName": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"delete-connection":            {"ClientId"},
			"delete-thing-shadow":          {"ThingName"},
			"get-retained-message":         {"Topic"},
			"get-thing-shadow":             {"ThingName"},
			"list-named-shadows-for-thing": {"ThingName"},
			"list-retained-messages":       {},
			"publish":                      {"Topic"},
			"update-thing-shadow":          {"Payload", "ThingName"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("iotdataplane", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
