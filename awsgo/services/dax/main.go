package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/dax/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"create-cluster", "create-parameter-group", "create-subnet-group", "decrease-replication-factor", "delete-cluster", "delete-parameter-group", "delete-subnet-group", "describe-clusters", "describe-default-parameters", "describe-events", "describe-parameter-groups", "describe-parameters", "describe-subnet-groups", "increase-replication-factor", "list-tags", "reboot-node", "tag-resource", "untag-resource", "update-cluster", "update-parameter-group", "update-subnet-group"},
		OperationSet: map[string]bool{"create-cluster": true, "create-parameter-group": true, "create-subnet-group": true, "decrease-replication-factor": true, "delete-cluster": true, "delete-parameter-group": true, "delete-subnet-group": true, "describe-clusters": true, "describe-default-parameters": true, "describe-events": true, "describe-parameter-groups": true, "describe-parameters": true, "describe-subnet-groups": true, "increase-replication-factor": true, "list-tags": true, "reboot-node": true, "tag-resource": true, "untag-resource": true, "update-cluster": true, "update-parameter-group": true, "update-subnet-group": true},
		OperationInputs: map[string][]string{
			"create-cluster":              {"AvailabilityZones", "ClusterEndpointEncryptionType", "ClusterName", "Description", "IamRoleArn", "NetworkType", "NodeType", "NotificationTopicArn", "ParameterGroupName", "PreferredMaintenanceWindow", "ReplicationFactor", "SSESpecification", "SecurityGroupIds", "SubnetGroupName", "Tags"},
			"create-parameter-group":      {"Description", "ParameterGroupName"},
			"create-subnet-group":         {"Description", "SubnetGroupName", "SubnetIds"},
			"decrease-replication-factor": {"AvailabilityZones", "ClusterName", "NewReplicationFactor", "NodeIdsToRemove"},
			"delete-cluster":              {"ClusterName"},
			"delete-parameter-group":      {"ParameterGroupName"},
			"delete-subnet-group":         {"SubnetGroupName"},
			"describe-clusters":           {"ClusterNames", "MaxResults", "NextToken"},
			"describe-default-parameters": {"MaxResults", "NextToken"},
			"describe-events":             {"Duration", "EndTime", "MaxResults", "NextToken", "SourceName", "SourceType", "StartTime"},
			"describe-parameter-groups":   {"MaxResults", "NextToken", "ParameterGroupNames"},
			"describe-parameters":         {"MaxResults", "NextToken", "ParameterGroupName", "Source"},
			"describe-subnet-groups":      {"MaxResults", "NextToken", "SubnetGroupNames"},
			"increase-replication-factor": {"AvailabilityZones", "ClusterName", "NewReplicationFactor"},
			"list-tags":                   {"NextToken", "ResourceName"},
			"reboot-node":                 {"ClusterName", "NodeId"},
			"tag-resource":                {"ResourceName", "Tags"},
			"untag-resource":              {"ResourceName", "TagKeys"},
			"update-cluster":              {"ClusterName", "Description", "NotificationTopicArn", "NotificationTopicStatus", "ParameterGroupName", "PreferredMaintenanceWindow", "SecurityGroupIds"},
			"update-parameter-group":      {"ParameterGroupName", "ParameterNameValues"},
			"update-subnet-group":         {"Description", "SubnetGroupName", "SubnetIds"},
		},
		OperationInputTypes: map[string]map[string]string{
			"create-cluster":              {"AvailabilityZones": "[]string", "ClusterEndpointEncryptionType": "types.ClusterEndpointEncryptionType", "ClusterName": "*string", "Description": "*string", "IamRoleArn": "*string", "NetworkType": "types.NetworkType", "NodeType": "*string", "NotificationTopicArn": "*string", "ParameterGroupName": "*string", "PreferredMaintenanceWindow": "*string", "ReplicationFactor": "int32", "SSESpecification": "*types.SSESpecification", "SecurityGroupIds": "[]string", "SubnetGroupName": "*string", "Tags": "[]types.Tag"},
			"create-parameter-group":      {"Description": "*string", "ParameterGroupName": "*string"},
			"create-subnet-group":         {"Description": "*string", "SubnetGroupName": "*string", "SubnetIds": "[]string"},
			"decrease-replication-factor": {"AvailabilityZones": "[]string", "ClusterName": "*string", "NewReplicationFactor": "int32", "NodeIdsToRemove": "[]string"},
			"delete-cluster":              {"ClusterName": "*string"},
			"delete-parameter-group":      {"ParameterGroupName": "*string"},
			"delete-subnet-group":         {"SubnetGroupName": "*string"},
			"describe-clusters":           {"ClusterNames": "[]string", "MaxResults": "*int32", "NextToken": "*string"},
			"describe-default-parameters": {"MaxResults": "*int32", "NextToken": "*string"},
			"describe-events":             {"Duration": "*int32", "EndTime": "*time.Time", "MaxResults": "*int32", "NextToken": "*string", "SourceName": "*string", "SourceType": "types.SourceType", "StartTime": "*time.Time"},
			"describe-parameter-groups":   {"MaxResults": "*int32", "NextToken": "*string", "ParameterGroupNames": "[]string"},
			"describe-parameters":         {"MaxResults": "*int32", "NextToken": "*string", "ParameterGroupName": "*string", "Source": "*string"},
			"describe-subnet-groups":      {"MaxResults": "*int32", "NextToken": "*string", "SubnetGroupNames": "[]string"},
			"increase-replication-factor": {"AvailabilityZones": "[]string", "ClusterName": "*string", "NewReplicationFactor": "int32"},
			"list-tags":                   {"NextToken": "*string", "ResourceName": "*string"},
			"reboot-node":                 {"ClusterName": "*string", "NodeId": "*string"},
			"tag-resource":                {"ResourceName": "*string", "Tags": "[]types.Tag"},
			"untag-resource":              {"ResourceName": "*string", "TagKeys": "[]string"},
			"update-cluster":              {"ClusterName": "*string", "Description": "*string", "NotificationTopicArn": "*string", "NotificationTopicStatus": "*string", "ParameterGroupName": "*string", "PreferredMaintenanceWindow": "*string", "SecurityGroupIds": "[]string"},
			"update-parameter-group":      {"ParameterGroupName": "*string", "ParameterNameValues": "[]types.ParameterNameValue"},
			"update-subnet-group":         {"Description": "*string", "SubnetGroupName": "*string", "SubnetIds": "[]string"},
		},
		OperationInputRequired: map[string][]string{
			"create-cluster":              {"ClusterName", "IamRoleArn", "NodeType", "ReplicationFactor"},
			"create-parameter-group":      {"ParameterGroupName"},
			"create-subnet-group":         {"SubnetGroupName", "SubnetIds"},
			"decrease-replication-factor": {"ClusterName", "NewReplicationFactor"},
			"delete-cluster":              {"ClusterName"},
			"delete-parameter-group":      {"ParameterGroupName"},
			"delete-subnet-group":         {"SubnetGroupName"},
			"describe-clusters":           {},
			"describe-default-parameters": {},
			"describe-events":             {},
			"describe-parameter-groups":   {},
			"describe-parameters":         {"ParameterGroupName"},
			"describe-subnet-groups":      {},
			"increase-replication-factor": {"ClusterName", "NewReplicationFactor"},
			"list-tags":                   {"ResourceName"},
			"reboot-node":                 {"ClusterName", "NodeId"},
			"tag-resource":                {"ResourceName", "Tags"},
			"untag-resource":              {"ResourceName", "TagKeys"},
			"update-cluster":              {"ClusterName"},
			"update-parameter-group":      {"ParameterGroupName", "ParameterNameValues"},
			"update-subnet-group":         {"SubnetGroupName"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("dax", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
