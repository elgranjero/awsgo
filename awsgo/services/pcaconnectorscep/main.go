package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/pcaconnectorscep/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"create-challenge", "create-connector", "delete-challenge", "delete-connector", "get-challenge-metadata", "get-challenge-password", "get-connector", "list-challenge-metadata", "list-connectors", "list-tags-for-resource", "tag-resource", "untag-resource"},
		OperationSet: map[string]bool{"create-challenge": true, "create-connector": true, "delete-challenge": true, "delete-connector": true, "get-challenge-metadata": true, "get-challenge-password": true, "get-connector": true, "list-challenge-metadata": true, "list-connectors": true, "list-tags-for-resource": true, "tag-resource": true, "untag-resource": true},
		OperationInputs: map[string][]string{
			"create-challenge":        {"ClientToken", "ConnectorArn", "Tags"},
			"create-connector":        {"CertificateAuthorityArn", "ClientToken", "MobileDeviceManagement", "Tags", "VpcEndpointId"},
			"delete-challenge":        {"ChallengeArn"},
			"delete-connector":        {"ConnectorArn"},
			"get-challenge-metadata":  {"ChallengeArn"},
			"get-challenge-password":  {"ChallengeArn"},
			"get-connector":           {"ConnectorArn"},
			"list-challenge-metadata": {"ConnectorArn", "MaxResults", "NextToken"},
			"list-connectors":         {"MaxResults", "NextToken"},
			"list-tags-for-resource":  {"ResourceArn"},
			"tag-resource":            {"ResourceArn", "Tags"},
			"untag-resource":          {"ResourceArn", "TagKeys"},
		},
		OperationInputTypes: map[string]map[string]string{
			"create-challenge":        {"ClientToken": "*string", "ConnectorArn": "*string", "Tags": "map[string]string"},
			"create-connector":        {"CertificateAuthorityArn": "*string", "ClientToken": "*string", "MobileDeviceManagement": "types.MobileDeviceManagement", "Tags": "map[string]string", "VpcEndpointId": "*string"},
			"delete-challenge":        {"ChallengeArn": "*string"},
			"delete-connector":        {"ConnectorArn": "*string"},
			"get-challenge-metadata":  {"ChallengeArn": "*string"},
			"get-challenge-password":  {"ChallengeArn": "*string"},
			"get-connector":           {"ConnectorArn": "*string"},
			"list-challenge-metadata": {"ConnectorArn": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"list-connectors":         {"MaxResults": "*int32", "NextToken": "*string"},
			"list-tags-for-resource":  {"ResourceArn": "*string"},
			"tag-resource":            {"ResourceArn": "*string", "Tags": "map[string]string"},
			"untag-resource":          {"ResourceArn": "*string", "TagKeys": "[]string"},
		},
		OperationInputRequired: map[string][]string{
			"create-challenge":        {"ConnectorArn"},
			"create-connector":        {"CertificateAuthorityArn"},
			"delete-challenge":        {"ChallengeArn"},
			"delete-connector":        {"ConnectorArn"},
			"get-challenge-metadata":  {"ChallengeArn"},
			"get-challenge-password":  {"ChallengeArn"},
			"get-connector":           {"ConnectorArn"},
			"list-challenge-metadata": {"ConnectorArn"},
			"list-connectors":         {},
			"list-tags-for-resource":  {"ResourceArn"},
			"tag-resource":            {"ResourceArn", "Tags"},
			"untag-resource":          {"ResourceArn", "TagKeys"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("pcaconnectorscep", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
