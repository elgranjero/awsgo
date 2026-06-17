package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/migrationhubrefactorspaces/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"create-application", "create-environment", "create-route", "create-service", "delete-application", "delete-environment", "delete-resource-policy", "delete-route", "delete-service", "get-application", "get-environment", "get-resource-policy", "get-route", "get-service", "list-applications", "list-environment-vpcs", "list-environments", "list-routes", "list-services", "list-tags-for-resource", "put-resource-policy", "tag-resource", "untag-resource", "update-route"},
		OperationSet: map[string]bool{"create-application": true, "create-environment": true, "create-route": true, "create-service": true, "delete-application": true, "delete-environment": true, "delete-resource-policy": true, "delete-route": true, "delete-service": true, "get-application": true, "get-environment": true, "get-resource-policy": true, "get-route": true, "get-service": true, "list-applications": true, "list-environment-vpcs": true, "list-environments": true, "list-routes": true, "list-services": true, "list-tags-for-resource": true, "put-resource-policy": true, "tag-resource": true, "untag-resource": true, "update-route": true},
		OperationInputs: map[string][]string{
			"create-application":     {"ApiGatewayProxy", "ClientToken", "EnvironmentIdentifier", "Name", "ProxyType", "Tags", "VpcId"},
			"create-environment":     {"ClientToken", "Description", "Name", "NetworkFabricType", "Tags"},
			"create-route":           {"ApplicationIdentifier", "ClientToken", "DefaultRoute", "EnvironmentIdentifier", "RouteType", "ServiceIdentifier", "Tags", "UriPathRoute"},
			"create-service":         {"ApplicationIdentifier", "ClientToken", "Description", "EndpointType", "EnvironmentIdentifier", "LambdaEndpoint", "Name", "Tags", "UrlEndpoint", "VpcId"},
			"delete-application":     {"ApplicationIdentifier", "EnvironmentIdentifier"},
			"delete-environment":     {"EnvironmentIdentifier"},
			"delete-resource-policy": {"Identifier"},
			"delete-route":           {"ApplicationIdentifier", "EnvironmentIdentifier", "RouteIdentifier"},
			"delete-service":         {"ApplicationIdentifier", "EnvironmentIdentifier", "ServiceIdentifier"},
			"get-application":        {"ApplicationIdentifier", "EnvironmentIdentifier"},
			"get-environment":        {"EnvironmentIdentifier"},
			"get-resource-policy":    {"Identifier"},
			"get-route":              {"ApplicationIdentifier", "EnvironmentIdentifier", "RouteIdentifier"},
			"get-service":            {"ApplicationIdentifier", "EnvironmentIdentifier", "ServiceIdentifier"},
			"list-applications":      {"EnvironmentIdentifier", "MaxResults", "NextToken"},
			"list-environment-vpcs":  {"EnvironmentIdentifier", "MaxResults", "NextToken"},
			"list-environments":      {"MaxResults", "NextToken"},
			"list-routes":            {"ApplicationIdentifier", "EnvironmentIdentifier", "MaxResults", "NextToken"},
			"list-services":          {"ApplicationIdentifier", "EnvironmentIdentifier", "MaxResults", "NextToken"},
			"list-tags-for-resource": {"ResourceArn"},
			"put-resource-policy":    {"Policy", "ResourceArn"},
			"tag-resource":           {"ResourceArn", "Tags"},
			"untag-resource":         {"ResourceArn", "TagKeys"},
			"update-route":           {"ActivationState", "ApplicationIdentifier", "EnvironmentIdentifier", "RouteIdentifier"},
		},
		OperationInputTypes: map[string]map[string]string{
			"create-application":     {"ApiGatewayProxy": "*types.ApiGatewayProxyInput", "ClientToken": "*string", "EnvironmentIdentifier": "*string", "Name": "*string", "ProxyType": "types.ProxyType", "Tags": "map[string]string", "VpcId": "*string"},
			"create-environment":     {"ClientToken": "*string", "Description": "*string", "Name": "*string", "NetworkFabricType": "types.NetworkFabricType", "Tags": "map[string]string"},
			"create-route":           {"ApplicationIdentifier": "*string", "ClientToken": "*string", "DefaultRoute": "*types.DefaultRouteInput", "EnvironmentIdentifier": "*string", "RouteType": "types.RouteType", "ServiceIdentifier": "*string", "Tags": "map[string]string", "UriPathRoute": "*types.UriPathRouteInput"},
			"create-service":         {"ApplicationIdentifier": "*string", "ClientToken": "*string", "Description": "*string", "EndpointType": "types.ServiceEndpointType", "EnvironmentIdentifier": "*string", "LambdaEndpoint": "*types.LambdaEndpointInput", "Name": "*string", "Tags": "map[string]string", "UrlEndpoint": "*types.UrlEndpointInput", "VpcId": "*string"},
			"delete-application":     {"ApplicationIdentifier": "*string", "EnvironmentIdentifier": "*string"},
			"delete-environment":     {"EnvironmentIdentifier": "*string"},
			"delete-resource-policy": {"Identifier": "*string"},
			"delete-route":           {"ApplicationIdentifier": "*string", "EnvironmentIdentifier": "*string", "RouteIdentifier": "*string"},
			"delete-service":         {"ApplicationIdentifier": "*string", "EnvironmentIdentifier": "*string", "ServiceIdentifier": "*string"},
			"get-application":        {"ApplicationIdentifier": "*string", "EnvironmentIdentifier": "*string"},
			"get-environment":        {"EnvironmentIdentifier": "*string"},
			"get-resource-policy":    {"Identifier": "*string"},
			"get-route":              {"ApplicationIdentifier": "*string", "EnvironmentIdentifier": "*string", "RouteIdentifier": "*string"},
			"get-service":            {"ApplicationIdentifier": "*string", "EnvironmentIdentifier": "*string", "ServiceIdentifier": "*string"},
			"list-applications":      {"EnvironmentIdentifier": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"list-environment-vpcs":  {"EnvironmentIdentifier": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"list-environments":      {"MaxResults": "*int32", "NextToken": "*string"},
			"list-routes":            {"ApplicationIdentifier": "*string", "EnvironmentIdentifier": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"list-services":          {"ApplicationIdentifier": "*string", "EnvironmentIdentifier": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"list-tags-for-resource": {"ResourceArn": "*string"},
			"put-resource-policy":    {"Policy": "*string", "ResourceArn": "*string"},
			"tag-resource":           {"ResourceArn": "*string", "Tags": "map[string]string"},
			"untag-resource":         {"ResourceArn": "*string", "TagKeys": "[]string"},
			"update-route":           {"ActivationState": "types.RouteActivationState", "ApplicationIdentifier": "*string", "EnvironmentIdentifier": "*string", "RouteIdentifier": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"create-application":     {"EnvironmentIdentifier", "Name", "ProxyType", "VpcId"},
			"create-environment":     {"Name", "NetworkFabricType"},
			"create-route":           {"ApplicationIdentifier", "EnvironmentIdentifier", "RouteType", "ServiceIdentifier"},
			"create-service":         {"ApplicationIdentifier", "EndpointType", "EnvironmentIdentifier", "Name"},
			"delete-application":     {"ApplicationIdentifier", "EnvironmentIdentifier"},
			"delete-environment":     {"EnvironmentIdentifier"},
			"delete-resource-policy": {"Identifier"},
			"delete-route":           {"ApplicationIdentifier", "EnvironmentIdentifier", "RouteIdentifier"},
			"delete-service":         {"ApplicationIdentifier", "EnvironmentIdentifier", "ServiceIdentifier"},
			"get-application":        {"ApplicationIdentifier", "EnvironmentIdentifier"},
			"get-environment":        {"EnvironmentIdentifier"},
			"get-resource-policy":    {"Identifier"},
			"get-route":              {"ApplicationIdentifier", "EnvironmentIdentifier", "RouteIdentifier"},
			"get-service":            {"ApplicationIdentifier", "EnvironmentIdentifier", "ServiceIdentifier"},
			"list-applications":      {"EnvironmentIdentifier"},
			"list-environment-vpcs":  {"EnvironmentIdentifier"},
			"list-environments":      {},
			"list-routes":            {"ApplicationIdentifier", "EnvironmentIdentifier"},
			"list-services":          {"ApplicationIdentifier", "EnvironmentIdentifier"},
			"list-tags-for-resource": {"ResourceArn"},
			"put-resource-policy":    {"Policy", "ResourceArn"},
			"tag-resource":           {"ResourceArn", "Tags"},
			"untag-resource":         {"ResourceArn", "TagKeys"},
			"update-route":           {"ActivationState", "ApplicationIdentifier", "EnvironmentIdentifier", "RouteIdentifier"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("migrationhubrefactorspaces", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
