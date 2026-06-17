package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/cloud9/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"create-environment-ec2", "create-environment-membership", "delete-environment", "delete-environment-membership", "describe-environment-memberships", "describe-environment-status", "describe-environments", "list-environments", "list-tags-for-resource", "tag-resource", "untag-resource", "update-environment", "update-environment-membership"},
		OperationSet: map[string]bool{"create-environment-ec2": true, "create-environment-membership": true, "delete-environment": true, "delete-environment-membership": true, "describe-environment-memberships": true, "describe-environment-status": true, "describe-environments": true, "list-environments": true, "list-tags-for-resource": true, "tag-resource": true, "untag-resource": true, "update-environment": true, "update-environment-membership": true},
		OperationInputs: map[string][]string{
			"create-environment-ec2":           {"AutomaticStopTimeMinutes", "ClientRequestToken", "ConnectionType", "Description", "DryRun", "ImageId", "InstanceType", "Name", "OwnerArn", "SubnetId", "Tags"},
			"create-environment-membership":    {"EnvironmentId", "Permissions", "UserArn"},
			"delete-environment":               {"EnvironmentId"},
			"delete-environment-membership":    {"EnvironmentId", "UserArn"},
			"describe-environment-memberships": {"EnvironmentId", "MaxResults", "NextToken", "Permissions", "UserArn"},
			"describe-environment-status":      {"EnvironmentId"},
			"describe-environments":            {"EnvironmentIds"},
			"list-environments":                {"MaxResults", "NextToken"},
			"list-tags-for-resource":           {"ResourceARN"},
			"tag-resource":                     {"ResourceARN", "Tags"},
			"untag-resource":                   {"ResourceARN", "TagKeys"},
			"update-environment":               {"Description", "EnvironmentId", "ManagedCredentialsAction", "Name"},
			"update-environment-membership":    {"EnvironmentId", "Permissions", "UserArn"},
		},
		OperationInputTypes: map[string]map[string]string{
			"create-environment-ec2":           {"AutomaticStopTimeMinutes": "*int32", "ClientRequestToken": "*string", "ConnectionType": "types.ConnectionType", "Description": "*string", "DryRun": "*bool", "ImageId": "*string", "InstanceType": "*string", "Name": "*string", "OwnerArn": "*string", "SubnetId": "*string", "Tags": "[]types.Tag"},
			"create-environment-membership":    {"EnvironmentId": "*string", "Permissions": "types.MemberPermissions", "UserArn": "*string"},
			"delete-environment":               {"EnvironmentId": "*string"},
			"delete-environment-membership":    {"EnvironmentId": "*string", "UserArn": "*string"},
			"describe-environment-memberships": {"EnvironmentId": "*string", "MaxResults": "*int32", "NextToken": "*string", "Permissions": "[]types.Permissions", "UserArn": "*string"},
			"describe-environment-status":      {"EnvironmentId": "*string"},
			"describe-environments":            {"EnvironmentIds": "[]string"},
			"list-environments":                {"MaxResults": "*int32", "NextToken": "*string"},
			"list-tags-for-resource":           {"ResourceARN": "*string"},
			"tag-resource":                     {"ResourceARN": "*string", "Tags": "[]types.Tag"},
			"untag-resource":                   {"ResourceARN": "*string", "TagKeys": "[]string"},
			"update-environment":               {"Description": "*string", "EnvironmentId": "*string", "ManagedCredentialsAction": "types.ManagedCredentialsAction", "Name": "*string"},
			"update-environment-membership":    {"EnvironmentId": "*string", "Permissions": "types.MemberPermissions", "UserArn": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"create-environment-ec2":           {"ImageId", "InstanceType", "Name"},
			"create-environment-membership":    {"EnvironmentId", "Permissions", "UserArn"},
			"delete-environment":               {"EnvironmentId"},
			"delete-environment-membership":    {"EnvironmentId", "UserArn"},
			"describe-environment-memberships": {},
			"describe-environment-status":      {"EnvironmentId"},
			"describe-environments":            {"EnvironmentIds"},
			"list-environments":                {},
			"list-tags-for-resource":           {"ResourceARN"},
			"tag-resource":                     {"ResourceARN", "Tags"},
			"untag-resource":                   {"ResourceARN", "TagKeys"},
			"update-environment":               {"EnvironmentId"},
			"update-environment-membership":    {"EnvironmentId", "Permissions", "UserArn"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("cloud9", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
