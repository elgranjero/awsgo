package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/timestreaminfluxdb/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"create-db-cluster", "create-db-instance", "create-db-parameter-group", "delete-db-cluster", "delete-db-instance", "get-db-cluster", "get-db-instance", "get-db-parameter-group", "list-db-clusters", "list-db-instances", "list-db-instances-for-cluster", "list-db-parameter-groups", "list-tags-for-resource", "reboot-db-cluster", "reboot-db-instance", "tag-resource", "untag-resource", "update-db-cluster", "update-db-instance"},
		OperationSet: map[string]bool{"create-db-cluster": true, "create-db-instance": true, "create-db-parameter-group": true, "delete-db-cluster": true, "delete-db-instance": true, "get-db-cluster": true, "get-db-instance": true, "get-db-parameter-group": true, "list-db-clusters": true, "list-db-instances": true, "list-db-instances-for-cluster": true, "list-db-parameter-groups": true, "list-tags-for-resource": true, "reboot-db-cluster": true, "reboot-db-instance": true, "tag-resource": true, "untag-resource": true, "update-db-cluster": true, "update-db-instance": true},
		OperationInputs: map[string][]string{
			"create-db-cluster":             {"AllocatedStorage", "Bucket", "DbInstanceType", "DbParameterGroupIdentifier", "DbStorageType", "DeploymentType", "FailoverMode", "LogDeliveryConfiguration", "Name", "NetworkType", "Organization", "Password", "Port", "PubliclyAccessible", "Tags", "Username", "VpcSecurityGroupIds", "VpcSubnetIds"},
			"create-db-instance":            {"AllocatedStorage", "Bucket", "DbInstanceType", "DbParameterGroupIdentifier", "DbStorageType", "DeploymentType", "LogDeliveryConfiguration", "Name", "NetworkType", "Organization", "Password", "Port", "PubliclyAccessible", "Tags", "Username", "VpcSecurityGroupIds", "VpcSubnetIds"},
			"create-db-parameter-group":     {"Description", "Name", "Parameters", "Tags"},
			"delete-db-cluster":             {"DbClusterId"},
			"delete-db-instance":            {"Identifier"},
			"get-db-cluster":                {"DbClusterId"},
			"get-db-instance":               {"Identifier"},
			"get-db-parameter-group":        {"Identifier"},
			"list-db-clusters":              {"MaxResults", "NextToken"},
			"list-db-instances":             {"MaxResults", "NextToken"},
			"list-db-instances-for-cluster": {"DbClusterId", "MaxResults", "NextToken"},
			"list-db-parameter-groups":      {"MaxResults", "NextToken"},
			"list-tags-for-resource":        {"ResourceArn"},
			"reboot-db-cluster":             {"DbClusterId", "InstanceIds"},
			"reboot-db-instance":            {"Identifier"},
			"tag-resource":                  {"ResourceArn", "Tags"},
			"untag-resource":                {"ResourceArn", "TagKeys"},
			"update-db-cluster":             {"DbClusterId", "DbInstanceType", "DbParameterGroupIdentifier", "FailoverMode", "LogDeliveryConfiguration", "Port"},
			"update-db-instance":            {"AllocatedStorage", "DbInstanceType", "DbParameterGroupIdentifier", "DbStorageType", "DeploymentType", "Identifier", "LogDeliveryConfiguration", "Port"},
		},
		OperationInputTypes: map[string]map[string]string{
			"create-db-cluster":             {"AllocatedStorage": "*int32", "Bucket": "*string", "DbInstanceType": "types.DbInstanceType", "DbParameterGroupIdentifier": "*string", "DbStorageType": "types.DbStorageType", "DeploymentType": "types.ClusterDeploymentType", "FailoverMode": "types.FailoverMode", "LogDeliveryConfiguration": "*types.LogDeliveryConfiguration", "Name": "*string", "NetworkType": "types.NetworkType", "Organization": "*string", "Password": "*string", "Port": "*int32", "PubliclyAccessible": "*bool", "Tags": "map[string]string", "Username": "*string", "VpcSecurityGroupIds": "[]string", "VpcSubnetIds": "[]string"},
			"create-db-instance":            {"AllocatedStorage": "*int32", "Bucket": "*string", "DbInstanceType": "types.DbInstanceType", "DbParameterGroupIdentifier": "*string", "DbStorageType": "types.DbStorageType", "DeploymentType": "types.DeploymentType", "LogDeliveryConfiguration": "*types.LogDeliveryConfiguration", "Name": "*string", "NetworkType": "types.NetworkType", "Organization": "*string", "Password": "*string", "Port": "*int32", "PubliclyAccessible": "*bool", "Tags": "map[string]string", "Username": "*string", "VpcSecurityGroupIds": "[]string", "VpcSubnetIds": "[]string"},
			"create-db-parameter-group":     {"Description": "*string", "Name": "*string", "Parameters": "types.Parameters", "Tags": "map[string]string"},
			"delete-db-cluster":             {"DbClusterId": "*string"},
			"delete-db-instance":            {"Identifier": "*string"},
			"get-db-cluster":                {"DbClusterId": "*string"},
			"get-db-instance":               {"Identifier": "*string"},
			"get-db-parameter-group":        {"Identifier": "*string"},
			"list-db-clusters":              {"MaxResults": "*int32", "NextToken": "*string"},
			"list-db-instances":             {"MaxResults": "*int32", "NextToken": "*string"},
			"list-db-instances-for-cluster": {"DbClusterId": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"list-db-parameter-groups":      {"MaxResults": "*int32", "NextToken": "*string"},
			"list-tags-for-resource":        {"ResourceArn": "*string"},
			"reboot-db-cluster":             {"DbClusterId": "*string", "InstanceIds": "[]string"},
			"reboot-db-instance":            {"Identifier": "*string"},
			"tag-resource":                  {"ResourceArn": "*string", "Tags": "map[string]string"},
			"untag-resource":                {"ResourceArn": "*string", "TagKeys": "[]string"},
			"update-db-cluster":             {"DbClusterId": "*string", "DbInstanceType": "types.DbInstanceType", "DbParameterGroupIdentifier": "*string", "FailoverMode": "types.FailoverMode", "LogDeliveryConfiguration": "*types.LogDeliveryConfiguration", "Port": "*int32"},
			"update-db-instance":            {"AllocatedStorage": "*int32", "DbInstanceType": "types.DbInstanceType", "DbParameterGroupIdentifier": "*string", "DbStorageType": "types.DbStorageType", "DeploymentType": "types.DeploymentType", "Identifier": "*string", "LogDeliveryConfiguration": "*types.LogDeliveryConfiguration", "Port": "*int32"},
		},
		OperationInputRequired: map[string][]string{
			"create-db-cluster":             {"DbInstanceType", "Name", "VpcSecurityGroupIds", "VpcSubnetIds"},
			"create-db-instance":            {"AllocatedStorage", "DbInstanceType", "Name", "Password", "VpcSecurityGroupIds", "VpcSubnetIds"},
			"create-db-parameter-group":     {"Name"},
			"delete-db-cluster":             {"DbClusterId"},
			"delete-db-instance":            {"Identifier"},
			"get-db-cluster":                {"DbClusterId"},
			"get-db-instance":               {"Identifier"},
			"get-db-parameter-group":        {"Identifier"},
			"list-db-clusters":              {},
			"list-db-instances":             {},
			"list-db-instances-for-cluster": {"DbClusterId"},
			"list-db-parameter-groups":      {},
			"list-tags-for-resource":        {"ResourceArn"},
			"reboot-db-cluster":             {"DbClusterId"},
			"reboot-db-instance":            {"Identifier"},
			"tag-resource":                  {"ResourceArn", "Tags"},
			"untag-resource":                {"ResourceArn", "TagKeys"},
			"update-db-cluster":             {"DbClusterId"},
			"update-db-instance":            {"Identifier"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("timestreaminfluxdb", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
