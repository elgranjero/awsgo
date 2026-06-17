package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/rtbfabric/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"accept-link", "create-inbound-external-link", "create-link", "create-outbound-external-link", "create-requester-gateway", "create-responder-gateway", "delete-inbound-external-link", "delete-link", "delete-outbound-external-link", "delete-requester-gateway", "delete-responder-gateway", "get-inbound-external-link", "get-link", "get-outbound-external-link", "get-requester-gateway", "get-responder-gateway", "list-links", "list-requester-gateways", "list-responder-gateways", "list-tags-for-resource", "reject-link", "tag-resource", "untag-resource", "update-link", "update-link-module-flow", "update-requester-gateway", "update-responder-gateway"},
		OperationSet: map[string]bool{"accept-link": true, "create-inbound-external-link": true, "create-link": true, "create-outbound-external-link": true, "create-requester-gateway": true, "create-responder-gateway": true, "delete-inbound-external-link": true, "delete-link": true, "delete-outbound-external-link": true, "delete-requester-gateway": true, "delete-responder-gateway": true, "get-inbound-external-link": true, "get-link": true, "get-outbound-external-link": true, "get-requester-gateway": true, "get-responder-gateway": true, "list-links": true, "list-requester-gateways": true, "list-responder-gateways": true, "list-tags-for-resource": true, "reject-link": true, "tag-resource": true, "untag-resource": true, "update-link": true, "update-link-module-flow": true, "update-requester-gateway": true, "update-responder-gateway": true},
		OperationInputs: map[string][]string{
			"accept-link":                   {"Attributes", "GatewayId", "LinkId", "LogSettings"},
			"create-inbound-external-link":  {"Attributes", "ClientToken", "GatewayId", "LogSettings", "Tags"},
			"create-link":                   {"Attributes", "GatewayId", "HttpResponderAllowed", "LogSettings", "PeerGatewayId", "Tags"},
			"create-outbound-external-link": {"Attributes", "ClientToken", "GatewayId", "LogSettings", "PublicEndpoint", "Tags"},
			"create-requester-gateway":      {"ClientToken", "Description", "SecurityGroupIds", "SubnetIds", "Tags", "VpcId"},
			"create-responder-gateway":      {"ClientToken", "Description", "DomainName", "ManagedEndpointConfiguration", "Port", "Protocol", "SecurityGroupIds", "SubnetIds", "Tags", "TrustStoreConfiguration", "VpcId"},
			"delete-inbound-external-link":  {"GatewayId", "LinkId"},
			"delete-link":                   {"GatewayId", "LinkId"},
			"delete-outbound-external-link": {"GatewayId", "LinkId"},
			"delete-requester-gateway":      {"GatewayId"},
			"delete-responder-gateway":      {"GatewayId"},
			"get-inbound-external-link":     {"GatewayId", "LinkId"},
			"get-link":                      {"GatewayId", "LinkId"},
			"get-outbound-external-link":    {"GatewayId", "LinkId"},
			"get-requester-gateway":         {"GatewayId"},
			"get-responder-gateway":         {"GatewayId"},
			"list-links":                    {"GatewayId", "MaxResults", "NextToken"},
			"list-requester-gateways":       {"MaxResults", "NextToken"},
			"list-responder-gateways":       {"MaxResults", "NextToken"},
			"list-tags-for-resource":        {"ResourceArn"},
			"reject-link":                   {"GatewayId", "LinkId"},
			"tag-resource":                  {"ResourceArn", "Tags"},
			"untag-resource":                {"ResourceArn", "TagKeys"},
			"update-link":                   {"GatewayId", "LinkId", "LogSettings"},
			"update-link-module-flow":       {"ClientToken", "GatewayId", "LinkId", "Modules"},
			"update-requester-gateway":      {"ClientToken", "Description", "GatewayId"},
			"update-responder-gateway":      {"ClientToken", "Description", "DomainName", "GatewayId", "ManagedEndpointConfiguration", "Port", "Protocol", "TrustStoreConfiguration"},
		},
		OperationInputTypes: map[string]map[string]string{
			"accept-link":                   {"Attributes": "*types.LinkAttributes", "GatewayId": "*string", "LinkId": "*string", "LogSettings": "*types.LinkLogSettings"},
			"create-inbound-external-link":  {"Attributes": "*types.LinkAttributes", "ClientToken": "*string", "GatewayId": "*string", "LogSettings": "*types.LinkLogSettings", "Tags": "map[string]string"},
			"create-link":                   {"Attributes": "*types.LinkAttributes", "GatewayId": "*string", "HttpResponderAllowed": "*bool", "LogSettings": "*types.LinkLogSettings", "PeerGatewayId": "*string", "Tags": "map[string]string"},
			"create-outbound-external-link": {"Attributes": "*types.LinkAttributes", "ClientToken": "*string", "GatewayId": "*string", "LogSettings": "*types.LinkLogSettings", "PublicEndpoint": "*string", "Tags": "map[string]string"},
			"create-requester-gateway":      {"ClientToken": "*string", "Description": "*string", "SecurityGroupIds": "[]string", "SubnetIds": "[]string", "Tags": "map[string]string", "VpcId": "*string"},
			"create-responder-gateway":      {"ClientToken": "*string", "Description": "*string", "DomainName": "*string", "ManagedEndpointConfiguration": "types.ManagedEndpointConfiguration", "Port": "*int32", "Protocol": "types.Protocol", "SecurityGroupIds": "[]string", "SubnetIds": "[]string", "Tags": "map[string]string", "TrustStoreConfiguration": "*types.TrustStoreConfiguration", "VpcId": "*string"},
			"delete-inbound-external-link":  {"GatewayId": "*string", "LinkId": "*string"},
			"delete-link":                   {"GatewayId": "*string", "LinkId": "*string"},
			"delete-outbound-external-link": {"GatewayId": "*string", "LinkId": "*string"},
			"delete-requester-gateway":      {"GatewayId": "*string"},
			"delete-responder-gateway":      {"GatewayId": "*string"},
			"get-inbound-external-link":     {"GatewayId": "*string", "LinkId": "*string"},
			"get-link":                      {"GatewayId": "*string", "LinkId": "*string"},
			"get-outbound-external-link":    {"GatewayId": "*string", "LinkId": "*string"},
			"get-requester-gateway":         {"GatewayId": "*string"},
			"get-responder-gateway":         {"GatewayId": "*string"},
			"list-links":                    {"GatewayId": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"list-requester-gateways":       {"MaxResults": "*int32", "NextToken": "*string"},
			"list-responder-gateways":       {"MaxResults": "*int32", "NextToken": "*string"},
			"list-tags-for-resource":        {"ResourceArn": "*string"},
			"reject-link":                   {"GatewayId": "*string", "LinkId": "*string"},
			"tag-resource":                  {"ResourceArn": "*string", "Tags": "map[string]string"},
			"untag-resource":                {"ResourceArn": "*string", "TagKeys": "[]string"},
			"update-link":                   {"GatewayId": "*string", "LinkId": "*string", "LogSettings": "*types.LinkLogSettings"},
			"update-link-module-flow":       {"ClientToken": "*string", "GatewayId": "*string", "LinkId": "*string", "Modules": "[]types.ModuleConfiguration"},
			"update-requester-gateway":      {"ClientToken": "*string", "Description": "*string", "GatewayId": "*string"},
			"update-responder-gateway":      {"ClientToken": "*string", "Description": "*string", "DomainName": "*string", "GatewayId": "*string", "ManagedEndpointConfiguration": "types.ManagedEndpointConfiguration", "Port": "*int32", "Protocol": "types.Protocol", "TrustStoreConfiguration": "*types.TrustStoreConfiguration"},
		},
		OperationInputRequired: map[string][]string{
			"accept-link":                   {"GatewayId", "LinkId", "LogSettings"},
			"create-inbound-external-link":  {"ClientToken", "GatewayId", "LogSettings"},
			"create-link":                   {"GatewayId", "LogSettings", "PeerGatewayId"},
			"create-outbound-external-link": {"ClientToken", "GatewayId", "LogSettings", "PublicEndpoint"},
			"create-requester-gateway":      {"ClientToken", "SecurityGroupIds", "SubnetIds", "VpcId"},
			"create-responder-gateway":      {"ClientToken", "Port", "Protocol", "SecurityGroupIds", "SubnetIds", "VpcId"},
			"delete-inbound-external-link":  {"GatewayId", "LinkId"},
			"delete-link":                   {"GatewayId", "LinkId"},
			"delete-outbound-external-link": {"GatewayId", "LinkId"},
			"delete-requester-gateway":      {"GatewayId"},
			"delete-responder-gateway":      {"GatewayId"},
			"get-inbound-external-link":     {"GatewayId", "LinkId"},
			"get-link":                      {"GatewayId", "LinkId"},
			"get-outbound-external-link":    {"GatewayId", "LinkId"},
			"get-requester-gateway":         {"GatewayId"},
			"get-responder-gateway":         {"GatewayId"},
			"list-links":                    {"GatewayId"},
			"list-requester-gateways":       {},
			"list-responder-gateways":       {},
			"list-tags-for-resource":        {"ResourceArn"},
			"reject-link":                   {"GatewayId", "LinkId"},
			"tag-resource":                  {"ResourceArn", "Tags"},
			"untag-resource":                {"ResourceArn", "TagKeys"},
			"update-link":                   {"GatewayId", "LinkId"},
			"update-link-module-flow":       {"ClientToken", "GatewayId", "LinkId", "Modules"},
			"update-requester-gateway":      {"ClientToken", "GatewayId"},
			"update-responder-gateway":      {"ClientToken", "GatewayId", "Port", "Protocol"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("rtbfabric", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
