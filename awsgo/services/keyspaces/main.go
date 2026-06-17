package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/keyspaces/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"create-keyspace", "create-table", "create-type", "delete-keyspace", "delete-table", "delete-type", "get-keyspace", "get-table", "get-table-auto-scaling-settings", "get-type", "list-keyspaces", "list-tables", "list-tags-for-resource", "list-types", "restore-table", "tag-resource", "untag-resource", "update-keyspace", "update-table"},
		OperationSet: map[string]bool{"create-keyspace": true, "create-table": true, "create-type": true, "delete-keyspace": true, "delete-table": true, "delete-type": true, "get-keyspace": true, "get-table": true, "get-table-auto-scaling-settings": true, "get-type": true, "list-keyspaces": true, "list-tables": true, "list-tags-for-resource": true, "list-types": true, "restore-table": true, "tag-resource": true, "untag-resource": true, "update-keyspace": true, "update-table": true},
		OperationInputs: map[string][]string{
			"create-keyspace":                 {"KeyspaceName", "ReplicationSpecification", "Tags"},
			"create-table":                    {"AutoScalingSpecification", "CapacitySpecification", "CdcSpecification", "ClientSideTimestamps", "Comment", "DefaultTimeToLive", "EncryptionSpecification", "KeyspaceName", "PointInTimeRecovery", "ReplicaSpecifications", "SchemaDefinition", "TableName", "Tags", "Ttl", "WarmThroughputSpecification"},
			"create-type":                     {"FieldDefinitions", "KeyspaceName", "TypeName"},
			"delete-keyspace":                 {"KeyspaceName"},
			"delete-table":                    {"KeyspaceName", "TableName"},
			"delete-type":                     {"KeyspaceName", "TypeName"},
			"get-keyspace":                    {"KeyspaceName"},
			"get-table":                       {"KeyspaceName", "TableName"},
			"get-table-auto-scaling-settings": {"KeyspaceName", "TableName"},
			"get-type":                        {"KeyspaceName", "TypeName"},
			"list-keyspaces":                  {"MaxResults", "NextToken"},
			"list-tables":                     {"KeyspaceName", "MaxResults", "NextToken"},
			"list-tags-for-resource":          {"MaxResults", "NextToken", "ResourceArn"},
			"list-types":                      {"KeyspaceName", "MaxResults", "NextToken"},
			"restore-table":                   {"AutoScalingSpecification", "CapacitySpecificationOverride", "EncryptionSpecificationOverride", "PointInTimeRecoveryOverride", "ReplicaSpecifications", "RestoreTimestamp", "SourceKeyspaceName", "SourceTableName", "TagsOverride", "TargetKeyspaceName", "TargetTableName"},
			"tag-resource":                    {"ResourceArn", "Tags"},
			"untag-resource":                  {"ResourceArn", "Tags"},
			"update-keyspace":                 {"ClientSideTimestamps", "KeyspaceName", "ReplicationSpecification"},
			"update-table":                    {"AddColumns", "AutoScalingSpecification", "CapacitySpecification", "CdcSpecification", "ClientSideTimestamps", "DefaultTimeToLive", "EncryptionSpecification", "KeyspaceName", "PointInTimeRecovery", "ReplicaSpecifications", "TableName", "Ttl", "WarmThroughputSpecification"},
		},
		OperationInputTypes: map[string]map[string]string{
			"create-keyspace":                 {"KeyspaceName": "*string", "ReplicationSpecification": "*types.ReplicationSpecification", "Tags": "[]types.Tag"},
			"create-table":                    {"AutoScalingSpecification": "*types.AutoScalingSpecification", "CapacitySpecification": "*types.CapacitySpecification", "CdcSpecification": "*types.CdcSpecification", "ClientSideTimestamps": "*types.ClientSideTimestamps", "Comment": "*types.Comment", "DefaultTimeToLive": "*int32", "EncryptionSpecification": "*types.EncryptionSpecification", "KeyspaceName": "*string", "PointInTimeRecovery": "*types.PointInTimeRecovery", "ReplicaSpecifications": "[]types.ReplicaSpecification", "SchemaDefinition": "*types.SchemaDefinition", "TableName": "*string", "Tags": "[]types.Tag", "Ttl": "*types.TimeToLive", "WarmThroughputSpecification": "*types.WarmThroughputSpecification"},
			"create-type":                     {"FieldDefinitions": "[]types.FieldDefinition", "KeyspaceName": "*string", "TypeName": "*string"},
			"delete-keyspace":                 {"KeyspaceName": "*string"},
			"delete-table":                    {"KeyspaceName": "*string", "TableName": "*string"},
			"delete-type":                     {"KeyspaceName": "*string", "TypeName": "*string"},
			"get-keyspace":                    {"KeyspaceName": "*string"},
			"get-table":                       {"KeyspaceName": "*string", "TableName": "*string"},
			"get-table-auto-scaling-settings": {"KeyspaceName": "*string", "TableName": "*string"},
			"get-type":                        {"KeyspaceName": "*string", "TypeName": "*string"},
			"list-keyspaces":                  {"MaxResults": "*int32", "NextToken": "*string"},
			"list-tables":                     {"KeyspaceName": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"list-tags-for-resource":          {"MaxResults": "*int32", "NextToken": "*string", "ResourceArn": "*string"},
			"list-types":                      {"KeyspaceName": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"restore-table":                   {"AutoScalingSpecification": "*types.AutoScalingSpecification", "CapacitySpecificationOverride": "*types.CapacitySpecification", "EncryptionSpecificationOverride": "*types.EncryptionSpecification", "PointInTimeRecoveryOverride": "*types.PointInTimeRecovery", "ReplicaSpecifications": "[]types.ReplicaSpecification", "RestoreTimestamp": "*time.Time", "SourceKeyspaceName": "*string", "SourceTableName": "*string", "TagsOverride": "[]types.Tag", "TargetKeyspaceName": "*string", "TargetTableName": "*string"},
			"tag-resource":                    {"ResourceArn": "*string", "Tags": "[]types.Tag"},
			"untag-resource":                  {"ResourceArn": "*string", "Tags": "[]types.Tag"},
			"update-keyspace":                 {"ClientSideTimestamps": "*types.ClientSideTimestamps", "KeyspaceName": "*string", "ReplicationSpecification": "*types.ReplicationSpecification"},
			"update-table":                    {"AddColumns": "[]types.ColumnDefinition", "AutoScalingSpecification": "*types.AutoScalingSpecification", "CapacitySpecification": "*types.CapacitySpecification", "CdcSpecification": "*types.CdcSpecification", "ClientSideTimestamps": "*types.ClientSideTimestamps", "DefaultTimeToLive": "*int32", "EncryptionSpecification": "*types.EncryptionSpecification", "KeyspaceName": "*string", "PointInTimeRecovery": "*types.PointInTimeRecovery", "ReplicaSpecifications": "[]types.ReplicaSpecification", "TableName": "*string", "Ttl": "*types.TimeToLive", "WarmThroughputSpecification": "*types.WarmThroughputSpecification"},
		},
		OperationInputRequired: map[string][]string{
			"create-keyspace":                 {"KeyspaceName"},
			"create-table":                    {"KeyspaceName", "SchemaDefinition", "TableName"},
			"create-type":                     {"FieldDefinitions", "KeyspaceName", "TypeName"},
			"delete-keyspace":                 {"KeyspaceName"},
			"delete-table":                    {"KeyspaceName", "TableName"},
			"delete-type":                     {"KeyspaceName", "TypeName"},
			"get-keyspace":                    {"KeyspaceName"},
			"get-table":                       {"KeyspaceName", "TableName"},
			"get-table-auto-scaling-settings": {"KeyspaceName", "TableName"},
			"get-type":                        {"KeyspaceName", "TypeName"},
			"list-keyspaces":                  {},
			"list-tables":                     {"KeyspaceName"},
			"list-tags-for-resource":          {"ResourceArn"},
			"list-types":                      {"KeyspaceName"},
			"restore-table":                   {"SourceKeyspaceName", "SourceTableName", "TargetKeyspaceName", "TargetTableName"},
			"tag-resource":                    {"ResourceArn", "Tags"},
			"untag-resource":                  {"ResourceArn", "Tags"},
			"update-keyspace":                 {"KeyspaceName", "ReplicationSpecification"},
			"update-table":                    {"KeyspaceName", "TableName"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("keyspaces", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
