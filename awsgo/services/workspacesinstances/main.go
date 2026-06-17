package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/workspacesinstances/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"associate-volume", "create-volume", "create-workspace-instance", "delete-volume", "delete-workspace-instance", "disassociate-volume", "get-workspace-instance", "list-instance-types", "list-regions", "list-tags-for-resource", "list-workspace-instances", "tag-resource", "untag-resource"},
		OperationSet: map[string]bool{"associate-volume": true, "create-volume": true, "create-workspace-instance": true, "delete-volume": true, "delete-workspace-instance": true, "disassociate-volume": true, "get-workspace-instance": true, "list-instance-types": true, "list-regions": true, "list-tags-for-resource": true, "list-workspace-instances": true, "tag-resource": true, "untag-resource": true},
		OperationInputs: map[string][]string{
			"associate-volume":          {"Device", "VolumeId", "WorkspaceInstanceId"},
			"create-volume":             {"AvailabilityZone", "ClientToken", "Encrypted", "Iops", "KmsKeyId", "SizeInGB", "SnapshotId", "TagSpecifications", "Throughput", "VolumeType"},
			"create-workspace-instance": {"BillingConfiguration", "ClientToken", "ManagedInstance", "Tags"},
			"delete-volume":             {"VolumeId"},
			"delete-workspace-instance": {"WorkspaceInstanceId"},
			"disassociate-volume":       {"Device", "DisassociateMode", "VolumeId", "WorkspaceInstanceId"},
			"get-workspace-instance":    {"WorkspaceInstanceId"},
			"list-instance-types":       {"InstanceConfigurationFilter", "MaxResults", "NextToken"},
			"list-regions":              {"MaxResults", "NextToken"},
			"list-tags-for-resource":    {"WorkspaceInstanceId"},
			"list-workspace-instances":  {"MaxResults", "NextToken", "ProvisionStates"},
			"tag-resource":              {"Tags", "WorkspaceInstanceId"},
			"untag-resource":            {"TagKeys", "WorkspaceInstanceId"},
		},
		OperationInputTypes: map[string]map[string]string{
			"associate-volume":          {"Device": "*string", "VolumeId": "*string", "WorkspaceInstanceId": "*string"},
			"create-volume":             {"AvailabilityZone": "*string", "ClientToken": "*string", "Encrypted": "*bool", "Iops": "*int32", "KmsKeyId": "*string", "SizeInGB": "*int32", "SnapshotId": "*string", "TagSpecifications": "[]types.TagSpecification", "Throughput": "*int32", "VolumeType": "types.VolumeTypeEnum"},
			"create-workspace-instance": {"BillingConfiguration": "*types.BillingConfiguration", "ClientToken": "*string", "ManagedInstance": "*types.ManagedInstanceRequest", "Tags": "[]types.Tag"},
			"delete-volume":             {"VolumeId": "*string"},
			"delete-workspace-instance": {"WorkspaceInstanceId": "*string"},
			"disassociate-volume":       {"Device": "*string", "DisassociateMode": "types.DisassociateModeEnum", "VolumeId": "*string", "WorkspaceInstanceId": "*string"},
			"get-workspace-instance":    {"WorkspaceInstanceId": "*string"},
			"list-instance-types":       {"InstanceConfigurationFilter": "*types.InstanceConfigurationFilter", "MaxResults": "*int32", "NextToken": "*string"},
			"list-regions":              {"MaxResults": "*int32", "NextToken": "*string"},
			"list-tags-for-resource":    {"WorkspaceInstanceId": "*string"},
			"list-workspace-instances":  {"MaxResults": "*int32", "NextToken": "*string", "ProvisionStates": "[]types.ProvisionStateEnum"},
			"tag-resource":              {"Tags": "[]types.Tag", "WorkspaceInstanceId": "*string"},
			"untag-resource":            {"TagKeys": "[]string", "WorkspaceInstanceId": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"associate-volume":          {"Device", "VolumeId", "WorkspaceInstanceId"},
			"create-volume":             {"AvailabilityZone"},
			"create-workspace-instance": {"ManagedInstance"},
			"delete-volume":             {"VolumeId"},
			"delete-workspace-instance": {"WorkspaceInstanceId"},
			"disassociate-volume":       {"VolumeId", "WorkspaceInstanceId"},
			"get-workspace-instance":    {"WorkspaceInstanceId"},
			"list-instance-types":       {},
			"list-regions":              {},
			"list-tags-for-resource":    {"WorkspaceInstanceId"},
			"list-workspace-instances":  {},
			"tag-resource":              {"Tags", "WorkspaceInstanceId"},
			"untag-resource":            {"TagKeys", "WorkspaceInstanceId"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("workspacesinstances", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
