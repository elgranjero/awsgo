package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/mediapackage/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"configure-logs", "create-channel", "create-harvest-job", "create-origin-endpoint", "delete-channel", "delete-origin-endpoint", "describe-channel", "describe-harvest-job", "describe-origin-endpoint", "list-channels", "list-harvest-jobs", "list-origin-endpoints", "list-tags-for-resource", "rotate-channel-credentials", "rotate-ingest-endpoint-credentials", "tag-resource", "untag-resource", "update-channel", "update-origin-endpoint"},
		OperationSet: map[string]bool{"configure-logs": true, "create-channel": true, "create-harvest-job": true, "create-origin-endpoint": true, "delete-channel": true, "delete-origin-endpoint": true, "describe-channel": true, "describe-harvest-job": true, "describe-origin-endpoint": true, "list-channels": true, "list-harvest-jobs": true, "list-origin-endpoints": true, "list-tags-for-resource": true, "rotate-channel-credentials": true, "rotate-ingest-endpoint-credentials": true, "tag-resource": true, "untag-resource": true, "update-channel": true, "update-origin-endpoint": true},
		OperationInputs: map[string][]string{
			"configure-logs":                     {"EgressAccessLogs", "Id", "IngressAccessLogs"},
			"create-channel":                     {"Description", "Id", "Tags"},
			"create-harvest-job":                 {"EndTime", "Id", "OriginEndpointId", "S3Destination", "StartTime"},
			"create-origin-endpoint":             {"Authorization", "ChannelId", "CmafPackage", "DashPackage", "Description", "HlsPackage", "Id", "ManifestName", "MssPackage", "Origination", "StartoverWindowSeconds", "Tags", "TimeDelaySeconds", "Whitelist"},
			"delete-channel":                     {"Id"},
			"delete-origin-endpoint":             {"Id"},
			"describe-channel":                   {"Id"},
			"describe-harvest-job":               {"Id"},
			"describe-origin-endpoint":           {"Id"},
			"list-channels":                      {"MaxResults", "NextToken"},
			"list-harvest-jobs":                  {"IncludeChannelId", "IncludeStatus", "MaxResults", "NextToken"},
			"list-origin-endpoints":              {"ChannelId", "MaxResults", "NextToken"},
			"list-tags-for-resource":             {"ResourceArn"},
			"rotate-channel-credentials":         {"Id"},
			"rotate-ingest-endpoint-credentials": {"Id", "IngestEndpointId"},
			"tag-resource":                       {"ResourceArn", "Tags"},
			"untag-resource":                     {"ResourceArn", "TagKeys"},
			"update-channel":                     {"Description", "Id"},
			"update-origin-endpoint":             {"Authorization", "CmafPackage", "DashPackage", "Description", "HlsPackage", "Id", "ManifestName", "MssPackage", "Origination", "StartoverWindowSeconds", "TimeDelaySeconds", "Whitelist"},
		},
		OperationInputTypes: map[string]map[string]string{
			"configure-logs":                     {"EgressAccessLogs": "*types.EgressAccessLogs", "Id": "*string", "IngressAccessLogs": "*types.IngressAccessLogs"},
			"create-channel":                     {"Description": "*string", "Id": "*string", "Tags": "map[string]string"},
			"create-harvest-job":                 {"EndTime": "*string", "Id": "*string", "OriginEndpointId": "*string", "S3Destination": "*types.S3Destination", "StartTime": "*string"},
			"create-origin-endpoint":             {"Authorization": "*types.Authorization", "ChannelId": "*string", "CmafPackage": "*types.CmafPackageCreateOrUpdateParameters", "DashPackage": "*types.DashPackage", "Description": "*string", "HlsPackage": "*types.HlsPackage", "Id": "*string", "ManifestName": "*string", "MssPackage": "*types.MssPackage", "Origination": "types.Origination", "StartoverWindowSeconds": "*int32", "Tags": "map[string]string", "TimeDelaySeconds": "*int32", "Whitelist": "[]string"},
			"delete-channel":                     {"Id": "*string"},
			"delete-origin-endpoint":             {"Id": "*string"},
			"describe-channel":                   {"Id": "*string"},
			"describe-harvest-job":               {"Id": "*string"},
			"describe-origin-endpoint":           {"Id": "*string"},
			"list-channels":                      {"MaxResults": "*int32", "NextToken": "*string"},
			"list-harvest-jobs":                  {"IncludeChannelId": "*string", "IncludeStatus": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"list-origin-endpoints":              {"ChannelId": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"list-tags-for-resource":             {"ResourceArn": "*string"},
			"rotate-channel-credentials":         {"Id": "*string"},
			"rotate-ingest-endpoint-credentials": {"Id": "*string", "IngestEndpointId": "*string"},
			"tag-resource":                       {"ResourceArn": "*string", "Tags": "map[string]string"},
			"untag-resource":                     {"ResourceArn": "*string", "TagKeys": "[]string"},
			"update-channel":                     {"Description": "*string", "Id": "*string"},
			"update-origin-endpoint":             {"Authorization": "*types.Authorization", "CmafPackage": "*types.CmafPackageCreateOrUpdateParameters", "DashPackage": "*types.DashPackage", "Description": "*string", "HlsPackage": "*types.HlsPackage", "Id": "*string", "ManifestName": "*string", "MssPackage": "*types.MssPackage", "Origination": "types.Origination", "StartoverWindowSeconds": "*int32", "TimeDelaySeconds": "*int32", "Whitelist": "[]string"},
		},
		OperationInputRequired: map[string][]string{
			"configure-logs":                     {"Id"},
			"create-channel":                     {"Id"},
			"create-harvest-job":                 {"EndTime", "Id", "OriginEndpointId", "S3Destination", "StartTime"},
			"create-origin-endpoint":             {"ChannelId", "Id"},
			"delete-channel":                     {"Id"},
			"delete-origin-endpoint":             {"Id"},
			"describe-channel":                   {"Id"},
			"describe-harvest-job":               {"Id"},
			"describe-origin-endpoint":           {"Id"},
			"list-channels":                      {},
			"list-harvest-jobs":                  {},
			"list-origin-endpoints":              {},
			"list-tags-for-resource":             {"ResourceArn"},
			"rotate-channel-credentials":         {"Id"},
			"rotate-ingest-endpoint-credentials": {"Id", "IngestEndpointId"},
			"tag-resource":                       {"ResourceArn", "Tags"},
			"untag-resource":                     {"ResourceArn", "TagKeys"},
			"update-channel":                     {"Id"},
			"update-origin-endpoint":             {"Id"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("mediapackage", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
