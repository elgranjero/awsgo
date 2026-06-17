package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/backupgateway/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"associate-gateway-to-server", "create-gateway", "delete-gateway", "delete-hypervisor", "disassociate-gateway-from-server", "get-bandwidth-rate-limit-schedule", "get-gateway", "get-hypervisor", "get-hypervisor-property-mappings", "get-virtual-machine", "import-hypervisor-configuration", "list-gateways", "list-hypervisors", "list-tags-for-resource", "list-virtual-machines", "put-bandwidth-rate-limit-schedule", "put-hypervisor-property-mappings", "put-maintenance-start-time", "start-virtual-machines-metadata-sync", "tag-resource", "test-hypervisor-configuration", "untag-resource", "update-gateway-information", "update-gateway-software-now", "update-hypervisor"},
		OperationSet: map[string]bool{"associate-gateway-to-server": true, "create-gateway": true, "delete-gateway": true, "delete-hypervisor": true, "disassociate-gateway-from-server": true, "get-bandwidth-rate-limit-schedule": true, "get-gateway": true, "get-hypervisor": true, "get-hypervisor-property-mappings": true, "get-virtual-machine": true, "import-hypervisor-configuration": true, "list-gateways": true, "list-hypervisors": true, "list-tags-for-resource": true, "list-virtual-machines": true, "put-bandwidth-rate-limit-schedule": true, "put-hypervisor-property-mappings": true, "put-maintenance-start-time": true, "start-virtual-machines-metadata-sync": true, "tag-resource": true, "test-hypervisor-configuration": true, "untag-resource": true, "update-gateway-information": true, "update-gateway-software-now": true, "update-hypervisor": true},
		OperationInputs: map[string][]string{
			"associate-gateway-to-server":          {"GatewayArn", "ServerArn"},
			"create-gateway":                       {"ActivationKey", "GatewayDisplayName", "GatewayType", "Tags"},
			"delete-gateway":                       {"GatewayArn"},
			"delete-hypervisor":                    {"HypervisorArn"},
			"disassociate-gateway-from-server":     {"GatewayArn"},
			"get-bandwidth-rate-limit-schedule":    {"GatewayArn"},
			"get-gateway":                          {"GatewayArn"},
			"get-hypervisor":                       {"HypervisorArn"},
			"get-hypervisor-property-mappings":     {"HypervisorArn"},
			"get-virtual-machine":                  {"ResourceArn"},
			"import-hypervisor-configuration":      {"Host", "KmsKeyArn", "Name", "Password", "Tags", "Username"},
			"list-gateways":                        {"MaxResults", "NextToken"},
			"list-hypervisors":                     {"MaxResults", "NextToken"},
			"list-tags-for-resource":               {"ResourceArn"},
			"list-virtual-machines":                {"HypervisorArn", "MaxResults", "NextToken"},
			"put-bandwidth-rate-limit-schedule":    {"BandwidthRateLimitIntervals", "GatewayArn"},
			"put-hypervisor-property-mappings":     {"HypervisorArn", "IamRoleArn", "VmwareToAwsTagMappings"},
			"put-maintenance-start-time":           {"DayOfMonth", "DayOfWeek", "GatewayArn", "HourOfDay", "MinuteOfHour"},
			"start-virtual-machines-metadata-sync": {"HypervisorArn"},
			"tag-resource":                         {"ResourceARN", "Tags"},
			"test-hypervisor-configuration":        {"GatewayArn", "Host", "Password", "Username"},
			"untag-resource":                       {"ResourceARN", "TagKeys"},
			"update-gateway-information":           {"GatewayArn", "GatewayDisplayName"},
			"update-gateway-software-now":          {"GatewayArn"},
			"update-hypervisor":                    {"Host", "HypervisorArn", "LogGroupArn", "Name", "Password", "Username"},
		},
		OperationInputTypes: map[string]map[string]string{
			"associate-gateway-to-server":          {"GatewayArn": "*string", "ServerArn": "*string"},
			"create-gateway":                       {"ActivationKey": "*string", "GatewayDisplayName": "*string", "GatewayType": "types.GatewayType", "Tags": "[]types.Tag"},
			"delete-gateway":                       {"GatewayArn": "*string"},
			"delete-hypervisor":                    {"HypervisorArn": "*string"},
			"disassociate-gateway-from-server":     {"GatewayArn": "*string"},
			"get-bandwidth-rate-limit-schedule":    {"GatewayArn": "*string"},
			"get-gateway":                          {"GatewayArn": "*string"},
			"get-hypervisor":                       {"HypervisorArn": "*string"},
			"get-hypervisor-property-mappings":     {"HypervisorArn": "*string"},
			"get-virtual-machine":                  {"ResourceArn": "*string"},
			"import-hypervisor-configuration":      {"Host": "*string", "KmsKeyArn": "*string", "Name": "*string", "Password": "*string", "Tags": "[]types.Tag", "Username": "*string"},
			"list-gateways":                        {"MaxResults": "*int32", "NextToken": "*string"},
			"list-hypervisors":                     {"MaxResults": "*int32", "NextToken": "*string"},
			"list-tags-for-resource":               {"ResourceArn": "*string"},
			"list-virtual-machines":                {"HypervisorArn": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"put-bandwidth-rate-limit-schedule":    {"BandwidthRateLimitIntervals": "[]types.BandwidthRateLimitInterval", "GatewayArn": "*string"},
			"put-hypervisor-property-mappings":     {"HypervisorArn": "*string", "IamRoleArn": "*string", "VmwareToAwsTagMappings": "[]types.VmwareToAwsTagMapping"},
			"put-maintenance-start-time":           {"DayOfMonth": "*int32", "DayOfWeek": "*int32", "GatewayArn": "*string", "HourOfDay": "*int32", "MinuteOfHour": "*int32"},
			"start-virtual-machines-metadata-sync": {"HypervisorArn": "*string"},
			"tag-resource":                         {"ResourceARN": "*string", "Tags": "[]types.Tag"},
			"test-hypervisor-configuration":        {"GatewayArn": "*string", "Host": "*string", "Password": "*string", "Username": "*string"},
			"untag-resource":                       {"ResourceARN": "*string", "TagKeys": "[]string"},
			"update-gateway-information":           {"GatewayArn": "*string", "GatewayDisplayName": "*string"},
			"update-gateway-software-now":          {"GatewayArn": "*string"},
			"update-hypervisor":                    {"Host": "*string", "HypervisorArn": "*string", "LogGroupArn": "*string", "Name": "*string", "Password": "*string", "Username": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"associate-gateway-to-server":          {"GatewayArn", "ServerArn"},
			"create-gateway":                       {"ActivationKey", "GatewayDisplayName", "GatewayType"},
			"delete-gateway":                       {"GatewayArn"},
			"delete-hypervisor":                    {"HypervisorArn"},
			"disassociate-gateway-from-server":     {"GatewayArn"},
			"get-bandwidth-rate-limit-schedule":    {"GatewayArn"},
			"get-gateway":                          {"GatewayArn"},
			"get-hypervisor":                       {"HypervisorArn"},
			"get-hypervisor-property-mappings":     {"HypervisorArn"},
			"get-virtual-machine":                  {"ResourceArn"},
			"import-hypervisor-configuration":      {"Host", "Name"},
			"list-gateways":                        {},
			"list-hypervisors":                     {},
			"list-tags-for-resource":               {"ResourceArn"},
			"list-virtual-machines":                {},
			"put-bandwidth-rate-limit-schedule":    {"BandwidthRateLimitIntervals", "GatewayArn"},
			"put-hypervisor-property-mappings":     {"HypervisorArn", "IamRoleArn", "VmwareToAwsTagMappings"},
			"put-maintenance-start-time":           {"GatewayArn", "HourOfDay", "MinuteOfHour"},
			"start-virtual-machines-metadata-sync": {"HypervisorArn"},
			"tag-resource":                         {"ResourceARN", "Tags"},
			"test-hypervisor-configuration":        {"GatewayArn", "Host"},
			"untag-resource":                       {"ResourceARN", "TagKeys"},
			"update-gateway-information":           {"GatewayArn"},
			"update-gateway-software-now":          {"GatewayArn"},
			"update-hypervisor":                    {"HypervisorArn"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("backupgateway", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
