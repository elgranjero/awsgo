package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/evs/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"associate-eip-to-vlan", "create-environment", "create-environment-host", "delete-environment", "delete-environment-host", "disassociate-eip-from-vlan", "get-environment", "get-versions", "list-environment-hosts", "list-environment-vlans", "list-environments", "list-tags-for-resource", "tag-resource", "untag-resource"},
		OperationSet: map[string]bool{"associate-eip-to-vlan": true, "create-environment": true, "create-environment-host": true, "delete-environment": true, "delete-environment-host": true, "disassociate-eip-from-vlan": true, "get-environment": true, "get-versions": true, "list-environment-hosts": true, "list-environment-vlans": true, "list-environments": true, "list-tags-for-resource": true, "tag-resource": true, "untag-resource": true},
		OperationInputs: map[string][]string{
			"associate-eip-to-vlan":      {"AllocationId", "ClientToken", "EnvironmentId", "VlanName"},
			"create-environment":         {"ClientToken", "ConnectivityInfo", "EnvironmentName", "Hosts", "InitialVlans", "KmsKeyId", "LicenseInfo", "ServiceAccessSecurityGroups", "ServiceAccessSubnetId", "SiteId", "Tags", "TermsAccepted", "VcfHostnames", "VcfVersion", "VpcId"},
			"create-environment-host":    {"ClientToken", "EnvironmentId", "EsxVersion", "Host"},
			"delete-environment":         {"ClientToken", "EnvironmentId"},
			"delete-environment-host":    {"ClientToken", "EnvironmentId", "HostName"},
			"disassociate-eip-from-vlan": {"AssociationId", "ClientToken", "EnvironmentId", "VlanName"},
			"get-environment":            {"EnvironmentId"},
			"get-versions":               {},
			"list-environment-hosts":     {"EnvironmentId", "MaxResults", "NextToken"},
			"list-environment-vlans":     {"EnvironmentId", "MaxResults", "NextToken"},
			"list-environments":          {"MaxResults", "NextToken", "State"},
			"list-tags-for-resource":     {"ResourceArn"},
			"tag-resource":               {"ResourceArn", "Tags"},
			"untag-resource":             {"ResourceArn", "TagKeys"},
		},
		OperationInputTypes: map[string]map[string]string{
			"associate-eip-to-vlan":      {"AllocationId": "*string", "ClientToken": "*string", "EnvironmentId": "*string", "VlanName": "*string"},
			"create-environment":         {"ClientToken": "*string", "ConnectivityInfo": "*types.ConnectivityInfo", "EnvironmentName": "*string", "Hosts": "[]types.HostInfoForCreate", "InitialVlans": "*types.InitialVlans", "KmsKeyId": "*string", "LicenseInfo": "[]types.LicenseInfo", "ServiceAccessSecurityGroups": "*types.ServiceAccessSecurityGroups", "ServiceAccessSubnetId": "*string", "SiteId": "*string", "Tags": "map[string]string", "TermsAccepted": "*bool", "VcfHostnames": "*types.VcfHostnames", "VcfVersion": "types.VcfVersion", "VpcId": "*string"},
			"create-environment-host":    {"ClientToken": "*string", "EnvironmentId": "*string", "EsxVersion": "*string", "Host": "*types.HostInfoForCreate"},
			"delete-environment":         {"ClientToken": "*string", "EnvironmentId": "*string"},
			"delete-environment-host":    {"ClientToken": "*string", "EnvironmentId": "*string", "HostName": "*string"},
			"disassociate-eip-from-vlan": {"AssociationId": "*string", "ClientToken": "*string", "EnvironmentId": "*string", "VlanName": "*string"},
			"get-environment":            {"EnvironmentId": "*string"},
			"get-versions":               {},
			"list-environment-hosts":     {"EnvironmentId": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"list-environment-vlans":     {"EnvironmentId": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"list-environments":          {"MaxResults": "*int32", "NextToken": "*string", "State": "[]types.EnvironmentState"},
			"list-tags-for-resource":     {"ResourceArn": "*string"},
			"tag-resource":               {"ResourceArn": "*string", "Tags": "map[string]string"},
			"untag-resource":             {"ResourceArn": "*string", "TagKeys": "[]string"},
		},
		OperationInputRequired: map[string][]string{
			"associate-eip-to-vlan":      {"AllocationId", "EnvironmentId", "VlanName"},
			"create-environment":         {"ConnectivityInfo", "Hosts", "InitialVlans", "LicenseInfo", "ServiceAccessSubnetId", "SiteId", "TermsAccepted", "VcfHostnames", "VcfVersion", "VpcId"},
			"create-environment-host":    {"EnvironmentId", "Host"},
			"delete-environment":         {"EnvironmentId"},
			"delete-environment-host":    {"EnvironmentId", "HostName"},
			"disassociate-eip-from-vlan": {"AssociationId", "EnvironmentId", "VlanName"},
			"get-environment":            {"EnvironmentId"},
			"get-versions":               {},
			"list-environment-hosts":     {"EnvironmentId"},
			"list-environment-vlans":     {"EnvironmentId"},
			"list-environments":          {},
			"list-tags-for-resource":     {"ResourceArn"},
			"tag-resource":               {"ResourceArn", "Tags"},
			"untag-resource":             {"ResourceArn", "TagKeys"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("evs", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
