package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/workspacesthinclient/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"create-environment", "delete-device", "delete-environment", "deregister-device", "get-device", "get-environment", "get-software-set", "list-devices", "list-environments", "list-software-sets", "list-tags-for-resource", "tag-resource", "untag-resource", "update-device", "update-environment", "update-software-set"},
		OperationSet: map[string]bool{"create-environment": true, "delete-device": true, "delete-environment": true, "deregister-device": true, "get-device": true, "get-environment": true, "get-software-set": true, "list-devices": true, "list-environments": true, "list-software-sets": true, "list-tags-for-resource": true, "tag-resource": true, "untag-resource": true, "update-device": true, "update-environment": true, "update-software-set": true},
		OperationInputs: map[string][]string{
			"create-environment":     {"ClientToken", "DesiredSoftwareSetId", "DesktopArn", "DesktopEndpoint", "DeviceCreationTags", "KmsKeyArn", "MaintenanceWindow", "Name", "SoftwareSetUpdateMode", "SoftwareSetUpdateSchedule", "Tags"},
			"delete-device":          {"ClientToken", "Id"},
			"delete-environment":     {"ClientToken", "Id"},
			"deregister-device":      {"ClientToken", "Id", "TargetDeviceStatus"},
			"get-device":             {"Id"},
			"get-environment":        {"Id"},
			"get-software-set":       {"Id"},
			"list-devices":           {"MaxResults", "NextToken"},
			"list-environments":      {"MaxResults", "NextToken"},
			"list-software-sets":     {"MaxResults", "NextToken"},
			"list-tags-for-resource": {"ResourceArn"},
			"tag-resource":           {"ResourceArn", "Tags"},
			"untag-resource":         {"ResourceArn", "TagKeys"},
			"update-device":          {"DesiredSoftwareSetId", "Id", "Name", "SoftwareSetUpdateSchedule"},
			"update-environment":     {"DesiredSoftwareSetId", "DesktopArn", "DesktopEndpoint", "DeviceCreationTags", "Id", "MaintenanceWindow", "Name", "SoftwareSetUpdateMode", "SoftwareSetUpdateSchedule"},
			"update-software-set":    {"Id", "ValidationStatus"},
		},
		OperationInputTypes: map[string]map[string]string{
			"create-environment":     {"ClientToken": "*string", "DesiredSoftwareSetId": "*string", "DesktopArn": "*string", "DesktopEndpoint": "*string", "DeviceCreationTags": "map[string]string", "KmsKeyArn": "*string", "MaintenanceWindow": "*types.MaintenanceWindow", "Name": "*string", "SoftwareSetUpdateMode": "types.SoftwareSetUpdateMode", "SoftwareSetUpdateSchedule": "types.SoftwareSetUpdateSchedule", "Tags": "map[string]string"},
			"delete-device":          {"ClientToken": "*string", "Id": "*string"},
			"delete-environment":     {"ClientToken": "*string", "Id": "*string"},
			"deregister-device":      {"ClientToken": "*string", "Id": "*string", "TargetDeviceStatus": "types.TargetDeviceStatus"},
			"get-device":             {"Id": "*string"},
			"get-environment":        {"Id": "*string"},
			"get-software-set":       {"Id": "*string"},
			"list-devices":           {"MaxResults": "*int32", "NextToken": "*string"},
			"list-environments":      {"MaxResults": "*int32", "NextToken": "*string"},
			"list-software-sets":     {"MaxResults": "*int32", "NextToken": "*string"},
			"list-tags-for-resource": {"ResourceArn": "*string"},
			"tag-resource":           {"ResourceArn": "*string", "Tags": "map[string]string"},
			"untag-resource":         {"ResourceArn": "*string", "TagKeys": "[]string"},
			"update-device":          {"DesiredSoftwareSetId": "*string", "Id": "*string", "Name": "*string", "SoftwareSetUpdateSchedule": "types.SoftwareSetUpdateSchedule"},
			"update-environment":     {"DesiredSoftwareSetId": "*string", "DesktopArn": "*string", "DesktopEndpoint": "*string", "DeviceCreationTags": "map[string]string", "Id": "*string", "MaintenanceWindow": "*types.MaintenanceWindow", "Name": "*string", "SoftwareSetUpdateMode": "types.SoftwareSetUpdateMode", "SoftwareSetUpdateSchedule": "types.SoftwareSetUpdateSchedule"},
			"update-software-set":    {"Id": "*string", "ValidationStatus": "types.SoftwareSetValidationStatus"},
		},
		OperationInputRequired: map[string][]string{
			"create-environment":     {"DesktopArn"},
			"delete-device":          {"Id"},
			"delete-environment":     {"Id"},
			"deregister-device":      {"Id"},
			"get-device":             {"Id"},
			"get-environment":        {"Id"},
			"get-software-set":       {"Id"},
			"list-devices":           {},
			"list-environments":      {},
			"list-software-sets":     {},
			"list-tags-for-resource": {"ResourceArn"},
			"tag-resource":           {"ResourceArn", "Tags"},
			"untag-resource":         {"ResourceArn", "TagKeys"},
			"update-device":          {"Id"},
			"update-environment":     {"Id"},
			"update-software-set":    {"Id", "ValidationStatus"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("workspacesthinclient", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
