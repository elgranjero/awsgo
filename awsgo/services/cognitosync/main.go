package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/cognitosync/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"bulk-publish", "delete-dataset", "describe-dataset", "describe-identity-pool-usage", "describe-identity-usage", "get-bulk-publish-details", "get-cognito-events", "get-identity-pool-configuration", "list-datasets", "list-identity-pool-usage", "list-records", "register-device", "set-cognito-events", "set-identity-pool-configuration", "subscribe-to-dataset", "unsubscribe-from-dataset", "update-records"},
		OperationSet: map[string]bool{"bulk-publish": true, "delete-dataset": true, "describe-dataset": true, "describe-identity-pool-usage": true, "describe-identity-usage": true, "get-bulk-publish-details": true, "get-cognito-events": true, "get-identity-pool-configuration": true, "list-datasets": true, "list-identity-pool-usage": true, "list-records": true, "register-device": true, "set-cognito-events": true, "set-identity-pool-configuration": true, "subscribe-to-dataset": true, "unsubscribe-from-dataset": true, "update-records": true},
		OperationInputs: map[string][]string{
			"bulk-publish":                    {"IdentityPoolId"},
			"delete-dataset":                  {"DatasetName", "IdentityId", "IdentityPoolId"},
			"describe-dataset":                {"DatasetName", "IdentityId", "IdentityPoolId"},
			"describe-identity-pool-usage":    {"IdentityPoolId"},
			"describe-identity-usage":         {"IdentityId", "IdentityPoolId"},
			"get-bulk-publish-details":        {"IdentityPoolId"},
			"get-cognito-events":              {"IdentityPoolId"},
			"get-identity-pool-configuration": {"IdentityPoolId"},
			"list-datasets":                   {"IdentityId", "IdentityPoolId", "MaxResults", "NextToken"},
			"list-identity-pool-usage":        {"MaxResults", "NextToken"},
			"list-records":                    {"DatasetName", "IdentityId", "IdentityPoolId", "LastSyncCount", "MaxResults", "NextToken", "SyncSessionToken"},
			"register-device":                 {"IdentityId", "IdentityPoolId", "Platform", "Token"},
			"set-cognito-events":              {"Events", "IdentityPoolId"},
			"set-identity-pool-configuration": {"CognitoStreams", "IdentityPoolId", "PushSync"},
			"subscribe-to-dataset":            {"DatasetName", "DeviceId", "IdentityId", "IdentityPoolId"},
			"unsubscribe-from-dataset":        {"DatasetName", "DeviceId", "IdentityId", "IdentityPoolId"},
			"update-records":                  {"ClientContext", "DatasetName", "DeviceId", "IdentityId", "IdentityPoolId", "RecordPatches", "SyncSessionToken"},
		},
		OperationInputTypes: map[string]map[string]string{
			"bulk-publish":                    {"IdentityPoolId": "*string"},
			"delete-dataset":                  {"DatasetName": "*string", "IdentityId": "*string", "IdentityPoolId": "*string"},
			"describe-dataset":                {"DatasetName": "*string", "IdentityId": "*string", "IdentityPoolId": "*string"},
			"describe-identity-pool-usage":    {"IdentityPoolId": "*string"},
			"describe-identity-usage":         {"IdentityId": "*string", "IdentityPoolId": "*string"},
			"get-bulk-publish-details":        {"IdentityPoolId": "*string"},
			"get-cognito-events":              {"IdentityPoolId": "*string"},
			"get-identity-pool-configuration": {"IdentityPoolId": "*string"},
			"list-datasets":                   {"IdentityId": "*string", "IdentityPoolId": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"list-identity-pool-usage":        {"MaxResults": "*int32", "NextToken": "*string"},
			"list-records":                    {"DatasetName": "*string", "IdentityId": "*string", "IdentityPoolId": "*string", "LastSyncCount": "*int64", "MaxResults": "*int32", "NextToken": "*string", "SyncSessionToken": "*string"},
			"register-device":                 {"IdentityId": "*string", "IdentityPoolId": "*string", "Platform": "types.Platform", "Token": "*string"},
			"set-cognito-events":              {"Events": "map[string]string", "IdentityPoolId": "*string"},
			"set-identity-pool-configuration": {"CognitoStreams": "*types.CognitoStreams", "IdentityPoolId": "*string", "PushSync": "*types.PushSync"},
			"subscribe-to-dataset":            {"DatasetName": "*string", "DeviceId": "*string", "IdentityId": "*string", "IdentityPoolId": "*string"},
			"unsubscribe-from-dataset":        {"DatasetName": "*string", "DeviceId": "*string", "IdentityId": "*string", "IdentityPoolId": "*string"},
			"update-records":                  {"ClientContext": "*string", "DatasetName": "*string", "DeviceId": "*string", "IdentityId": "*string", "IdentityPoolId": "*string", "RecordPatches": "[]types.RecordPatch", "SyncSessionToken": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"bulk-publish":                    {"IdentityPoolId"},
			"delete-dataset":                  {"DatasetName", "IdentityId", "IdentityPoolId"},
			"describe-dataset":                {"DatasetName", "IdentityId", "IdentityPoolId"},
			"describe-identity-pool-usage":    {"IdentityPoolId"},
			"describe-identity-usage":         {"IdentityId", "IdentityPoolId"},
			"get-bulk-publish-details":        {"IdentityPoolId"},
			"get-cognito-events":              {"IdentityPoolId"},
			"get-identity-pool-configuration": {"IdentityPoolId"},
			"list-datasets":                   {"IdentityId", "IdentityPoolId"},
			"list-identity-pool-usage":        {},
			"list-records":                    {"DatasetName", "IdentityId", "IdentityPoolId"},
			"register-device":                 {"IdentityId", "IdentityPoolId", "Platform", "Token"},
			"set-cognito-events":              {"Events", "IdentityPoolId"},
			"set-identity-pool-configuration": {"IdentityPoolId"},
			"subscribe-to-dataset":            {"DatasetName", "DeviceId", "IdentityId", "IdentityPoolId"},
			"unsubscribe-from-dataset":        {"DatasetName", "DeviceId", "IdentityId", "IdentityPoolId"},
			"update-records":                  {"DatasetName", "IdentityId", "IdentityPoolId", "SyncSessionToken"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("cognitosync", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
