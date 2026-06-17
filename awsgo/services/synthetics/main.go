package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/synthetics/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"associate-resource", "create-canary", "create-group", "delete-canary", "delete-group", "describe-canaries", "describe-canaries-last-run", "describe-runtime-versions", "disassociate-resource", "get-canary", "get-canary-runs", "get-group", "list-associated-groups", "list-group-resources", "list-groups", "list-tags-for-resource", "start-canary", "start-canary-dry-run", "stop-canary", "tag-resource", "untag-resource", "update-canary"},
		OperationSet: map[string]bool{"associate-resource": true, "create-canary": true, "create-group": true, "delete-canary": true, "delete-group": true, "describe-canaries": true, "describe-canaries-last-run": true, "describe-runtime-versions": true, "disassociate-resource": true, "get-canary": true, "get-canary-runs": true, "get-group": true, "list-associated-groups": true, "list-group-resources": true, "list-groups": true, "list-tags-for-resource": true, "start-canary": true, "start-canary-dry-run": true, "stop-canary": true, "tag-resource": true, "untag-resource": true, "update-canary": true},
		OperationInputs: map[string][]string{
			"associate-resource":         {"GroupIdentifier", "ResourceArn"},
			"create-canary":              {"ArtifactConfig", "ArtifactS3Location", "BrowserConfigs", "Code", "ExecutionRoleArn", "FailureRetentionPeriodInDays", "Name", "ProvisionedResourceCleanup", "ResourcesToReplicateTags", "RunConfig", "RuntimeVersion", "Schedule", "SuccessRetentionPeriodInDays", "Tags", "VpcConfig"},
			"create-group":               {"Name", "Tags"},
			"delete-canary":              {"DeleteLambda", "Name"},
			"delete-group":               {"GroupIdentifier"},
			"describe-canaries":          {"MaxResults", "Names", "NextToken"},
			"describe-canaries-last-run": {"BrowserType", "MaxResults", "Names", "NextToken"},
			"describe-runtime-versions":  {"MaxResults", "NextToken"},
			"disassociate-resource":      {"GroupIdentifier", "ResourceArn"},
			"get-canary":                 {"DryRunId", "Name"},
			"get-canary-runs":            {"DryRunId", "MaxResults", "Name", "NextToken", "RunType"},
			"get-group":                  {"GroupIdentifier"},
			"list-associated-groups":     {"MaxResults", "NextToken", "ResourceArn"},
			"list-group-resources":       {"GroupIdentifier", "MaxResults", "NextToken"},
			"list-groups":                {"MaxResults", "NextToken"},
			"list-tags-for-resource":     {"ResourceArn"},
			"start-canary":               {"Name"},
			"start-canary-dry-run":       {"ArtifactConfig", "ArtifactS3Location", "BrowserConfigs", "Code", "ExecutionRoleArn", "FailureRetentionPeriodInDays", "Name", "ProvisionedResourceCleanup", "RunConfig", "RuntimeVersion", "SuccessRetentionPeriodInDays", "VisualReference", "VisualReferences", "VpcConfig"},
			"stop-canary":                {"Name"},
			"tag-resource":               {"ResourceArn", "Tags"},
			"untag-resource":             {"ResourceArn", "TagKeys"},
			"update-canary":              {"ArtifactConfig", "ArtifactS3Location", "BrowserConfigs", "Code", "DryRunId", "ExecutionRoleArn", "FailureRetentionPeriodInDays", "Name", "ProvisionedResourceCleanup", "RunConfig", "RuntimeVersion", "Schedule", "SuccessRetentionPeriodInDays", "VisualReference", "VisualReferences", "VpcConfig"},
		},
		OperationInputTypes: map[string]map[string]string{
			"associate-resource":         {"GroupIdentifier": "*string", "ResourceArn": "*string"},
			"create-canary":              {"ArtifactConfig": "*types.ArtifactConfigInput", "ArtifactS3Location": "*string", "BrowserConfigs": "[]types.BrowserConfig", "Code": "*types.CanaryCodeInput", "ExecutionRoleArn": "*string", "FailureRetentionPeriodInDays": "*int32", "Name": "*string", "ProvisionedResourceCleanup": "types.ProvisionedResourceCleanupSetting", "ResourcesToReplicateTags": "[]types.ResourceToTag", "RunConfig": "*types.CanaryRunConfigInput", "RuntimeVersion": "*string", "Schedule": "*types.CanaryScheduleInput", "SuccessRetentionPeriodInDays": "*int32", "Tags": "map[string]string", "VpcConfig": "*types.VpcConfigInput"},
			"create-group":               {"Name": "*string", "Tags": "map[string]string"},
			"delete-canary":              {"DeleteLambda": "bool", "Name": "*string"},
			"delete-group":               {"GroupIdentifier": "*string"},
			"describe-canaries":          {"MaxResults": "*int32", "Names": "[]string", "NextToken": "*string"},
			"describe-canaries-last-run": {"BrowserType": "types.BrowserType", "MaxResults": "*int32", "Names": "[]string", "NextToken": "*string"},
			"describe-runtime-versions":  {"MaxResults": "*int32", "NextToken": "*string"},
			"disassociate-resource":      {"GroupIdentifier": "*string", "ResourceArn": "*string"},
			"get-canary":                 {"DryRunId": "*string", "Name": "*string"},
			"get-canary-runs":            {"DryRunId": "*string", "MaxResults": "*int32", "Name": "*string", "NextToken": "*string", "RunType": "types.RunType"},
			"get-group":                  {"GroupIdentifier": "*string"},
			"list-associated-groups":     {"MaxResults": "*int32", "NextToken": "*string", "ResourceArn": "*string"},
			"list-group-resources":       {"GroupIdentifier": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"list-groups":                {"MaxResults": "*int32", "NextToken": "*string"},
			"list-tags-for-resource":     {"ResourceArn": "*string"},
			"start-canary":               {"Name": "*string"},
			"start-canary-dry-run":       {"ArtifactConfig": "*types.ArtifactConfigInput", "ArtifactS3Location": "*string", "BrowserConfigs": "[]types.BrowserConfig", "Code": "*types.CanaryCodeInput", "ExecutionRoleArn": "*string", "FailureRetentionPeriodInDays": "*int32", "Name": "*string", "ProvisionedResourceCleanup": "types.ProvisionedResourceCleanupSetting", "RunConfig": "*types.CanaryRunConfigInput", "RuntimeVersion": "*string", "SuccessRetentionPeriodInDays": "*int32", "VisualReference": "*types.VisualReferenceInput", "VisualReferences": "[]types.VisualReferenceInput", "VpcConfig": "*types.VpcConfigInput"},
			"stop-canary":                {"Name": "*string"},
			"tag-resource":               {"ResourceArn": "*string", "Tags": "map[string]string"},
			"untag-resource":             {"ResourceArn": "*string", "TagKeys": "[]string"},
			"update-canary":              {"ArtifactConfig": "*types.ArtifactConfigInput", "ArtifactS3Location": "*string", "BrowserConfigs": "[]types.BrowserConfig", "Code": "*types.CanaryCodeInput", "DryRunId": "*string", "ExecutionRoleArn": "*string", "FailureRetentionPeriodInDays": "*int32", "Name": "*string", "ProvisionedResourceCleanup": "types.ProvisionedResourceCleanupSetting", "RunConfig": "*types.CanaryRunConfigInput", "RuntimeVersion": "*string", "Schedule": "*types.CanaryScheduleInput", "SuccessRetentionPeriodInDays": "*int32", "VisualReference": "*types.VisualReferenceInput", "VisualReferences": "[]types.VisualReferenceInput", "VpcConfig": "*types.VpcConfigInput"},
		},
		OperationInputRequired: map[string][]string{
			"associate-resource":         {"GroupIdentifier", "ResourceArn"},
			"create-canary":              {"ArtifactS3Location", "Code", "ExecutionRoleArn", "Name", "RuntimeVersion", "Schedule"},
			"create-group":               {"Name"},
			"delete-canary":              {"Name"},
			"delete-group":               {"GroupIdentifier"},
			"describe-canaries":          {},
			"describe-canaries-last-run": {},
			"describe-runtime-versions":  {},
			"disassociate-resource":      {"GroupIdentifier", "ResourceArn"},
			"get-canary":                 {"Name"},
			"get-canary-runs":            {"Name"},
			"get-group":                  {"GroupIdentifier"},
			"list-associated-groups":     {"ResourceArn"},
			"list-group-resources":       {"GroupIdentifier"},
			"list-groups":                {},
			"list-tags-for-resource":     {"ResourceArn"},
			"start-canary":               {"Name"},
			"start-canary-dry-run":       {"Name"},
			"stop-canary":                {"Name"},
			"tag-resource":               {"ResourceArn", "Tags"},
			"untag-resource":             {"ResourceArn", "TagKeys"},
			"update-canary":              {"Name"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("synthetics", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
