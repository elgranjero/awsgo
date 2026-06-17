package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/secretsmanager/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"batch-get-secret-value", "cancel-rotate-secret", "create-secret", "delete-resource-policy", "delete-secret", "describe-secret", "get-random-password", "get-resource-policy", "get-secret-value", "list-secret-version-ids", "list-secrets", "put-resource-policy", "put-secret-value", "remove-regions-from-replication", "replicate-secret-to-regions", "restore-secret", "rotate-secret", "stop-replication-to-replica", "tag-resource", "untag-resource", "update-secret", "update-secret-version-stage", "validate-resource-policy"},
		OperationSet: map[string]bool{"batch-get-secret-value": true, "cancel-rotate-secret": true, "create-secret": true, "delete-resource-policy": true, "delete-secret": true, "describe-secret": true, "get-random-password": true, "get-resource-policy": true, "get-secret-value": true, "list-secret-version-ids": true, "list-secrets": true, "put-resource-policy": true, "put-secret-value": true, "remove-regions-from-replication": true, "replicate-secret-to-regions": true, "restore-secret": true, "rotate-secret": true, "stop-replication-to-replica": true, "tag-resource": true, "untag-resource": true, "update-secret": true, "update-secret-version-stage": true, "validate-resource-policy": true},
		OperationInputs: map[string][]string{
			"batch-get-secret-value":          {"Filters", "MaxResults", "NextToken", "SecretIdList"},
			"cancel-rotate-secret":            {"SecretId"},
			"create-secret":                   {"AddReplicaRegions", "ClientRequestToken", "Description", "ForceOverwriteReplicaSecret", "KmsKeyId", "Name", "SecretBinary", "SecretString", "Tags", "Type"},
			"delete-resource-policy":          {"SecretId"},
			"delete-secret":                   {"ForceDeleteWithoutRecovery", "RecoveryWindowInDays", "SecretId"},
			"describe-secret":                 {"SecretId"},
			"get-random-password":             {"ExcludeCharacters", "ExcludeLowercase", "ExcludeNumbers", "ExcludePunctuation", "ExcludeUppercase", "IncludeSpace", "PasswordLength", "RequireEachIncludedType"},
			"get-resource-policy":             {"SecretId"},
			"get-secret-value":                {"SecretId", "VersionId", "VersionStage"},
			"list-secret-version-ids":         {"IncludeDeprecated", "MaxResults", "NextToken", "SecretId"},
			"list-secrets":                    {"Filters", "IncludePlannedDeletion", "MaxResults", "NextToken", "SortBy", "SortOrder"},
			"put-resource-policy":             {"BlockPublicPolicy", "ResourcePolicy", "SecretId"},
			"put-secret-value":                {"ClientRequestToken", "RotationToken", "SecretBinary", "SecretId", "SecretString", "VersionStages"},
			"remove-regions-from-replication": {"RemoveReplicaRegions", "SecretId"},
			"replicate-secret-to-regions":     {"AddReplicaRegions", "ForceOverwriteReplicaSecret", "SecretId"},
			"restore-secret":                  {"SecretId"},
			"rotate-secret":                   {"ClientRequestToken", "ExternalSecretRotationMetadata", "ExternalSecretRotationRoleArn", "RotateImmediately", "RotationLambdaARN", "RotationRules", "SecretId"},
			"stop-replication-to-replica":     {"SecretId"},
			"tag-resource":                    {"SecretId", "Tags"},
			"untag-resource":                  {"SecretId", "TagKeys"},
			"update-secret":                   {"ClientRequestToken", "Description", "KmsKeyId", "SecretBinary", "SecretId", "SecretString", "Type"},
			"update-secret-version-stage":     {"MoveToVersionId", "RemoveFromVersionId", "SecretId", "VersionStage"},
			"validate-resource-policy":        {"ResourcePolicy", "SecretId"},
		},
		OperationInputTypes: map[string]map[string]string{
			"batch-get-secret-value":          {"Filters": "[]types.Filter", "MaxResults": "*int32", "NextToken": "*string", "SecretIdList": "[]string"},
			"cancel-rotate-secret":            {"SecretId": "*string"},
			"create-secret":                   {"AddReplicaRegions": "[]types.ReplicaRegionType", "ClientRequestToken": "*string", "Description": "*string", "ForceOverwriteReplicaSecret": "bool", "KmsKeyId": "*string", "Name": "*string", "SecretBinary": "[]byte", "SecretString": "*string", "Tags": "[]types.Tag", "Type": "*string"},
			"delete-resource-policy":          {"SecretId": "*string"},
			"delete-secret":                   {"ForceDeleteWithoutRecovery": "*bool", "RecoveryWindowInDays": "*int64", "SecretId": "*string"},
			"describe-secret":                 {"SecretId": "*string"},
			"get-random-password":             {"ExcludeCharacters": "*string", "ExcludeLowercase": "*bool", "ExcludeNumbers": "*bool", "ExcludePunctuation": "*bool", "ExcludeUppercase": "*bool", "IncludeSpace": "*bool", "PasswordLength": "*int64", "RequireEachIncludedType": "*bool"},
			"get-resource-policy":             {"SecretId": "*string"},
			"get-secret-value":                {"SecretId": "*string", "VersionId": "*string", "VersionStage": "*string"},
			"list-secret-version-ids":         {"IncludeDeprecated": "*bool", "MaxResults": "*int32", "NextToken": "*string", "SecretId": "*string"},
			"list-secrets":                    {"Filters": "[]types.Filter", "IncludePlannedDeletion": "*bool", "MaxResults": "*int32", "NextToken": "*string", "SortBy": "types.SortByType", "SortOrder": "types.SortOrderType"},
			"put-resource-policy":             {"BlockPublicPolicy": "*bool", "ResourcePolicy": "*string", "SecretId": "*string"},
			"put-secret-value":                {"ClientRequestToken": "*string", "RotationToken": "*string", "SecretBinary": "[]byte", "SecretId": "*string", "SecretString": "*string", "VersionStages": "[]string"},
			"remove-regions-from-replication": {"RemoveReplicaRegions": "[]string", "SecretId": "*string"},
			"replicate-secret-to-regions":     {"AddReplicaRegions": "[]types.ReplicaRegionType", "ForceOverwriteReplicaSecret": "bool", "SecretId": "*string"},
			"restore-secret":                  {"SecretId": "*string"},
			"rotate-secret":                   {"ClientRequestToken": "*string", "ExternalSecretRotationMetadata": "[]types.ExternalSecretRotationMetadataItem", "ExternalSecretRotationRoleArn": "*string", "RotateImmediately": "*bool", "RotationLambdaARN": "*string", "RotationRules": "*types.RotationRulesType", "SecretId": "*string"},
			"stop-replication-to-replica":     {"SecretId": "*string"},
			"tag-resource":                    {"SecretId": "*string", "Tags": "[]types.Tag"},
			"untag-resource":                  {"SecretId": "*string", "TagKeys": "[]string"},
			"update-secret":                   {"ClientRequestToken": "*string", "Description": "*string", "KmsKeyId": "*string", "SecretBinary": "[]byte", "SecretId": "*string", "SecretString": "*string", "Type": "*string"},
			"update-secret-version-stage":     {"MoveToVersionId": "*string", "RemoveFromVersionId": "*string", "SecretId": "*string", "VersionStage": "*string"},
			"validate-resource-policy":        {"ResourcePolicy": "*string", "SecretId": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"batch-get-secret-value":          {},
			"cancel-rotate-secret":            {"SecretId"},
			"create-secret":                   {"Name"},
			"delete-resource-policy":          {"SecretId"},
			"delete-secret":                   {"SecretId"},
			"describe-secret":                 {"SecretId"},
			"get-random-password":             {},
			"get-resource-policy":             {"SecretId"},
			"get-secret-value":                {"SecretId"},
			"list-secret-version-ids":         {"SecretId"},
			"list-secrets":                    {},
			"put-resource-policy":             {"ResourcePolicy", "SecretId"},
			"put-secret-value":                {"SecretId"},
			"remove-regions-from-replication": {"RemoveReplicaRegions", "SecretId"},
			"replicate-secret-to-regions":     {"AddReplicaRegions", "SecretId"},
			"restore-secret":                  {"SecretId"},
			"rotate-secret":                   {"SecretId"},
			"stop-replication-to-replica":     {"SecretId"},
			"tag-resource":                    {"SecretId", "Tags"},
			"untag-resource":                  {"SecretId", "TagKeys"},
			"update-secret":                   {"SecretId"},
			"update-secret-version-stage":     {"SecretId", "VersionStage"},
			"validate-resource-policy":        {"ResourcePolicy"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("secretsmanager", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
