package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/cloudhsmv2/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"copy-backup-to-region", "create-cluster", "create-hsm", "delete-backup", "delete-cluster", "delete-hsm", "delete-resource-policy", "describe-backups", "describe-clusters", "get-resource-policy", "initialize-cluster", "list-tags", "modify-backup-attributes", "modify-cluster", "put-resource-policy", "restore-backup", "tag-resource", "untag-resource"},
		OperationSet: map[string]bool{"copy-backup-to-region": true, "create-cluster": true, "create-hsm": true, "delete-backup": true, "delete-cluster": true, "delete-hsm": true, "delete-resource-policy": true, "describe-backups": true, "describe-clusters": true, "get-resource-policy": true, "initialize-cluster": true, "list-tags": true, "modify-backup-attributes": true, "modify-cluster": true, "put-resource-policy": true, "restore-backup": true, "tag-resource": true, "untag-resource": true},
		OperationInputs: map[string][]string{
			"copy-backup-to-region":    {"BackupId", "DestinationRegion", "TagList"},
			"create-cluster":           {"BackupRetentionPolicy", "HsmType", "Mode", "NetworkType", "SourceBackupId", "SubnetIds", "TagList"},
			"create-hsm":               {"AvailabilityZone", "ClusterId", "IpAddress"},
			"delete-backup":            {"BackupId"},
			"delete-cluster":           {"ClusterId"},
			"delete-hsm":               {"ClusterId", "EniId", "EniIp", "HsmId"},
			"delete-resource-policy":   {"ResourceArn"},
			"describe-backups":         {"Filters", "MaxResults", "NextToken", "Shared", "SortAscending"},
			"describe-clusters":        {"Filters", "MaxResults", "NextToken"},
			"get-resource-policy":      {"ResourceArn"},
			"initialize-cluster":       {"ClusterId", "SignedCert", "TrustAnchor"},
			"list-tags":                {"MaxResults", "NextToken", "ResourceId"},
			"modify-backup-attributes": {"BackupId", "NeverExpires"},
			"modify-cluster":           {"BackupRetentionPolicy", "ClusterId", "HsmType"},
			"put-resource-policy":      {"Policy", "ResourceArn"},
			"restore-backup":           {"BackupId"},
			"tag-resource":             {"ResourceId", "TagList"},
			"untag-resource":           {"ResourceId", "TagKeyList"},
		},
		OperationInputTypes: map[string]map[string]string{
			"copy-backup-to-region":    {"BackupId": "*string", "DestinationRegion": "*string", "TagList": "[]types.Tag"},
			"create-cluster":           {"BackupRetentionPolicy": "*types.BackupRetentionPolicy", "HsmType": "*string", "Mode": "types.ClusterMode", "NetworkType": "types.NetworkType", "SourceBackupId": "*string", "SubnetIds": "[]string", "TagList": "[]types.Tag"},
			"create-hsm":               {"AvailabilityZone": "*string", "ClusterId": "*string", "IpAddress": "*string"},
			"delete-backup":            {"BackupId": "*string"},
			"delete-cluster":           {"ClusterId": "*string"},
			"delete-hsm":               {"ClusterId": "*string", "EniId": "*string", "EniIp": "*string", "HsmId": "*string"},
			"delete-resource-policy":   {"ResourceArn": "*string"},
			"describe-backups":         {"Filters": "map[string][]string", "MaxResults": "*int32", "NextToken": "*string", "Shared": "*bool", "SortAscending": "*bool"},
			"describe-clusters":        {"Filters": "map[string][]string", "MaxResults": "*int32", "NextToken": "*string"},
			"get-resource-policy":      {"ResourceArn": "*string"},
			"initialize-cluster":       {"ClusterId": "*string", "SignedCert": "*string", "TrustAnchor": "*string"},
			"list-tags":                {"MaxResults": "*int32", "NextToken": "*string", "ResourceId": "*string"},
			"modify-backup-attributes": {"BackupId": "*string", "NeverExpires": "*bool"},
			"modify-cluster":           {"BackupRetentionPolicy": "*types.BackupRetentionPolicy", "ClusterId": "*string", "HsmType": "*string"},
			"put-resource-policy":      {"Policy": "*string", "ResourceArn": "*string"},
			"restore-backup":           {"BackupId": "*string"},
			"tag-resource":             {"ResourceId": "*string", "TagList": "[]types.Tag"},
			"untag-resource":           {"ResourceId": "*string", "TagKeyList": "[]string"},
		},
		OperationInputRequired: map[string][]string{
			"copy-backup-to-region":    {"BackupId", "DestinationRegion"},
			"create-cluster":           {"HsmType", "SubnetIds"},
			"create-hsm":               {"AvailabilityZone", "ClusterId"},
			"delete-backup":            {"BackupId"},
			"delete-cluster":           {"ClusterId"},
			"delete-hsm":               {"ClusterId"},
			"delete-resource-policy":   {},
			"describe-backups":         {},
			"describe-clusters":        {},
			"get-resource-policy":      {},
			"initialize-cluster":       {"ClusterId", "SignedCert", "TrustAnchor"},
			"list-tags":                {"ResourceId"},
			"modify-backup-attributes": {"BackupId", "NeverExpires"},
			"modify-cluster":           {"ClusterId"},
			"put-resource-policy":      {},
			"restore-backup":           {"BackupId"},
			"tag-resource":             {"ResourceId", "TagList"},
			"untag-resource":           {"ResourceId", "TagKeyList"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("cloudhsmv2", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
