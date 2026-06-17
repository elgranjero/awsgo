package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/gameliftstreams/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"add-stream-group-locations", "associate-applications", "create-application", "create-stream-group", "create-stream-session-connection", "delete-application", "delete-stream-group", "disassociate-applications", "export-stream-session-files", "get-application", "get-stream-group", "get-stream-session", "list-applications", "list-stream-groups", "list-stream-sessions", "list-stream-sessions-by-account", "list-tags-for-resource", "remove-stream-group-locations", "start-stream-session", "tag-resource", "terminate-stream-session", "untag-resource", "update-application", "update-stream-group"},
		OperationSet: map[string]bool{"add-stream-group-locations": true, "associate-applications": true, "create-application": true, "create-stream-group": true, "create-stream-session-connection": true, "delete-application": true, "delete-stream-group": true, "disassociate-applications": true, "export-stream-session-files": true, "get-application": true, "get-stream-group": true, "get-stream-session": true, "list-applications": true, "list-stream-groups": true, "list-stream-sessions": true, "list-stream-sessions-by-account": true, "list-tags-for-resource": true, "remove-stream-group-locations": true, "start-stream-session": true, "tag-resource": true, "terminate-stream-session": true, "untag-resource": true, "update-application": true, "update-stream-group": true},
		OperationInputs: map[string][]string{
			"add-stream-group-locations":       {"Identifier", "LocationConfigurations"},
			"associate-applications":           {"ApplicationIdentifiers", "Identifier"},
			"create-application":               {"ApplicationLogOutputUri", "ApplicationLogPaths", "ApplicationSourceUri", "ClientToken", "Description", "ExecutablePath", "RuntimeEnvironment", "Tags"},
			"create-stream-group":              {"ClientToken", "DefaultApplicationIdentifier", "Description", "LocationConfigurations", "StreamClass", "Tags"},
			"create-stream-session-connection": {"ClientToken", "Identifier", "SignalRequest", "StreamSessionIdentifier"},
			"delete-application":               {"Identifier"},
			"delete-stream-group":              {"Identifier"},
			"disassociate-applications":        {"ApplicationIdentifiers", "Identifier"},
			"export-stream-session-files":      {"Identifier", "OutputUri", "StreamSessionIdentifier"},
			"get-application":                  {"Identifier"},
			"get-stream-group":                 {"Identifier"},
			"get-stream-session":               {"Identifier", "StreamSessionIdentifier"},
			"list-applications":                {"MaxResults", "NextToken"},
			"list-stream-groups":               {"MaxResults", "NextToken"},
			"list-stream-sessions":             {"ExportFilesStatus", "Identifier", "MaxResults", "NextToken", "Status"},
			"list-stream-sessions-by-account":  {"ExportFilesStatus", "MaxResults", "NextToken", "Status"},
			"list-tags-for-resource":           {"ResourceArn"},
			"remove-stream-group-locations":    {"Identifier", "Locations"},
			"start-stream-session":             {"AdditionalEnvironmentVariables", "AdditionalLaunchArgs", "ApplicationIdentifier", "ClientToken", "ConnectionTimeoutSeconds", "Description", "Identifier", "Locations", "PerformanceStatsConfiguration", "Protocol", "SessionLengthSeconds", "SignalRequest", "UserId"},
			"tag-resource":                     {"ResourceArn", "Tags"},
			"terminate-stream-session":         {"Identifier", "StreamSessionIdentifier"},
			"untag-resource":                   {"ResourceArn", "TagKeys"},
			"update-application":               {"ApplicationLogOutputUri", "ApplicationLogPaths", "Description", "Identifier"},
			"update-stream-group":              {"DefaultApplicationIdentifier", "Description", "Identifier", "LocationConfigurations"},
		},
		OperationInputTypes: map[string]map[string]string{
			"add-stream-group-locations":       {"Identifier": "*string", "LocationConfigurations": "[]types.LocationConfiguration"},
			"associate-applications":           {"ApplicationIdentifiers": "[]string", "Identifier": "*string"},
			"create-application":               {"ApplicationLogOutputUri": "*string", "ApplicationLogPaths": "[]string", "ApplicationSourceUri": "*string", "ClientToken": "*string", "Description": "*string", "ExecutablePath": "*string", "RuntimeEnvironment": "*types.RuntimeEnvironment", "Tags": "map[string]string"},
			"create-stream-group":              {"ClientToken": "*string", "DefaultApplicationIdentifier": "*string", "Description": "*string", "LocationConfigurations": "[]types.LocationConfiguration", "StreamClass": "types.StreamClass", "Tags": "map[string]string"},
			"create-stream-session-connection": {"ClientToken": "*string", "Identifier": "*string", "SignalRequest": "*string", "StreamSessionIdentifier": "*string"},
			"delete-application":               {"Identifier": "*string"},
			"delete-stream-group":              {"Identifier": "*string"},
			"disassociate-applications":        {"ApplicationIdentifiers": "[]string", "Identifier": "*string"},
			"export-stream-session-files":      {"Identifier": "*string", "OutputUri": "*string", "StreamSessionIdentifier": "*string"},
			"get-application":                  {"Identifier": "*string"},
			"get-stream-group":                 {"Identifier": "*string"},
			"get-stream-session":               {"Identifier": "*string", "StreamSessionIdentifier": "*string"},
			"list-applications":                {"MaxResults": "*int32", "NextToken": "*string"},
			"list-stream-groups":               {"MaxResults": "*int32", "NextToken": "*string"},
			"list-stream-sessions":             {"ExportFilesStatus": "types.ExportFilesStatus", "Identifier": "*string", "MaxResults": "*int32", "NextToken": "*string", "Status": "types.StreamSessionStatus"},
			"list-stream-sessions-by-account":  {"ExportFilesStatus": "types.ExportFilesStatus", "MaxResults": "*int32", "NextToken": "*string", "Status": "types.StreamSessionStatus"},
			"list-tags-for-resource":           {"ResourceArn": "*string"},
			"remove-stream-group-locations":    {"Identifier": "*string", "Locations": "[]string"},
			"start-stream-session":             {"AdditionalEnvironmentVariables": "map[string]string", "AdditionalLaunchArgs": "[]string", "ApplicationIdentifier": "*string", "ClientToken": "*string", "ConnectionTimeoutSeconds": "*int32", "Description": "*string", "Identifier": "*string", "Locations": "[]string", "PerformanceStatsConfiguration": "*types.PerformanceStatsConfiguration", "Protocol": "types.Protocol", "SessionLengthSeconds": "*int32", "SignalRequest": "*string", "UserId": "*string"},
			"tag-resource":                     {"ResourceArn": "*string", "Tags": "map[string]string"},
			"terminate-stream-session":         {"Identifier": "*string", "StreamSessionIdentifier": "*string"},
			"untag-resource":                   {"ResourceArn": "*string", "TagKeys": "[]string"},
			"update-application":               {"ApplicationLogOutputUri": "*string", "ApplicationLogPaths": "[]string", "Description": "*string", "Identifier": "*string"},
			"update-stream-group":              {"DefaultApplicationIdentifier": "*string", "Description": "*string", "Identifier": "*string", "LocationConfigurations": "[]types.LocationConfiguration"},
		},
		OperationInputRequired: map[string][]string{
			"add-stream-group-locations":       {"Identifier", "LocationConfigurations"},
			"associate-applications":           {"ApplicationIdentifiers", "Identifier"},
			"create-application":               {"ApplicationSourceUri", "Description", "ExecutablePath", "RuntimeEnvironment"},
			"create-stream-group":              {"Description", "StreamClass"},
			"create-stream-session-connection": {"Identifier", "SignalRequest", "StreamSessionIdentifier"},
			"delete-application":               {"Identifier"},
			"delete-stream-group":              {"Identifier"},
			"disassociate-applications":        {"ApplicationIdentifiers", "Identifier"},
			"export-stream-session-files":      {"Identifier", "OutputUri", "StreamSessionIdentifier"},
			"get-application":                  {"Identifier"},
			"get-stream-group":                 {"Identifier"},
			"get-stream-session":               {"Identifier", "StreamSessionIdentifier"},
			"list-applications":                {},
			"list-stream-groups":               {},
			"list-stream-sessions":             {"Identifier"},
			"list-stream-sessions-by-account":  {},
			"list-tags-for-resource":           {"ResourceArn"},
			"remove-stream-group-locations":    {"Identifier", "Locations"},
			"start-stream-session":             {"ApplicationIdentifier", "Identifier", "Protocol", "SignalRequest"},
			"tag-resource":                     {"ResourceArn", "Tags"},
			"terminate-stream-session":         {"Identifier", "StreamSessionIdentifier"},
			"untag-resource":                   {"ResourceArn", "TagKeys"},
			"update-application":               {"Identifier"},
			"update-stream-group":              {"Identifier"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("gameliftstreams", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
