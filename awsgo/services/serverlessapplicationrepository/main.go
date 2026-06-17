package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/serverlessapplicationrepository/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"create-application", "create-application-version", "create-cloud-formation-change-set", "create-cloud-formation-template", "delete-application", "get-application", "get-application-policy", "get-cloud-formation-template", "list-application-dependencies", "list-application-versions", "list-applications", "put-application-policy", "unshare-application", "update-application"},
		OperationSet: map[string]bool{"create-application": true, "create-application-version": true, "create-cloud-formation-change-set": true, "create-cloud-formation-template": true, "delete-application": true, "get-application": true, "get-application-policy": true, "get-cloud-formation-template": true, "list-application-dependencies": true, "list-application-versions": true, "list-applications": true, "put-application-policy": true, "unshare-application": true, "update-application": true},
		OperationInputs: map[string][]string{
			"create-application":                {"Author", "Description", "HomePageUrl", "Labels", "LicenseBody", "LicenseUrl", "Name", "ReadmeBody", "ReadmeUrl", "SemanticVersion", "SourceCodeArchiveUrl", "SourceCodeUrl", "SpdxLicenseId", "TemplateBody", "TemplateUrl"},
			"create-application-version":        {"ApplicationId", "SemanticVersion", "SourceCodeArchiveUrl", "SourceCodeUrl", "TemplateBody", "TemplateUrl"},
			"create-cloud-formation-change-set": {"ApplicationId", "Capabilities", "ChangeSetName", "ClientToken", "Description", "NotificationArns", "ParameterOverrides", "ResourceTypes", "RollbackConfiguration", "SemanticVersion", "StackName", "Tags", "TemplateId"},
			"create-cloud-formation-template":   {"ApplicationId", "SemanticVersion"},
			"delete-application":                {"ApplicationId"},
			"get-application":                   {"ApplicationId", "SemanticVersion"},
			"get-application-policy":            {"ApplicationId"},
			"get-cloud-formation-template":      {"ApplicationId", "TemplateId"},
			"list-application-dependencies":     {"ApplicationId", "MaxItems", "NextToken", "SemanticVersion"},
			"list-application-versions":         {"ApplicationId", "MaxItems", "NextToken"},
			"list-applications":                 {"MaxItems", "NextToken"},
			"put-application-policy":            {"ApplicationId", "Statements"},
			"unshare-application":               {"ApplicationId", "OrganizationId"},
			"update-application":                {"ApplicationId", "Author", "Description", "HomePageUrl", "Labels", "ReadmeBody", "ReadmeUrl"},
		},
		OperationInputTypes: map[string]map[string]string{
			"create-application":                {"Author": "*string", "Description": "*string", "HomePageUrl": "*string", "Labels": "[]string", "LicenseBody": "*string", "LicenseUrl": "*string", "Name": "*string", "ReadmeBody": "*string", "ReadmeUrl": "*string", "SemanticVersion": "*string", "SourceCodeArchiveUrl": "*string", "SourceCodeUrl": "*string", "SpdxLicenseId": "*string", "TemplateBody": "*string", "TemplateUrl": "*string"},
			"create-application-version":        {"ApplicationId": "*string", "SemanticVersion": "*string", "SourceCodeArchiveUrl": "*string", "SourceCodeUrl": "*string", "TemplateBody": "*string", "TemplateUrl": "*string"},
			"create-cloud-formation-change-set": {"ApplicationId": "*string", "Capabilities": "[]string", "ChangeSetName": "*string", "ClientToken": "*string", "Description": "*string", "NotificationArns": "[]string", "ParameterOverrides": "[]types.ParameterValue", "ResourceTypes": "[]string", "RollbackConfiguration": "*types.RollbackConfiguration", "SemanticVersion": "*string", "StackName": "*string", "Tags": "[]types.Tag", "TemplateId": "*string"},
			"create-cloud-formation-template":   {"ApplicationId": "*string", "SemanticVersion": "*string"},
			"delete-application":                {"ApplicationId": "*string"},
			"get-application":                   {"ApplicationId": "*string", "SemanticVersion": "*string"},
			"get-application-policy":            {"ApplicationId": "*string"},
			"get-cloud-formation-template":      {"ApplicationId": "*string", "TemplateId": "*string"},
			"list-application-dependencies":     {"ApplicationId": "*string", "MaxItems": "*int32", "NextToken": "*string", "SemanticVersion": "*string"},
			"list-application-versions":         {"ApplicationId": "*string", "MaxItems": "*int32", "NextToken": "*string"},
			"list-applications":                 {"MaxItems": "*int32", "NextToken": "*string"},
			"put-application-policy":            {"ApplicationId": "*string", "Statements": "[]types.ApplicationPolicyStatement"},
			"unshare-application":               {"ApplicationId": "*string", "OrganizationId": "*string"},
			"update-application":                {"ApplicationId": "*string", "Author": "*string", "Description": "*string", "HomePageUrl": "*string", "Labels": "[]string", "ReadmeBody": "*string", "ReadmeUrl": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"create-application":                {"Author", "Description", "Name"},
			"create-application-version":        {"ApplicationId", "SemanticVersion"},
			"create-cloud-formation-change-set": {"ApplicationId", "StackName"},
			"create-cloud-formation-template":   {"ApplicationId"},
			"delete-application":                {"ApplicationId"},
			"get-application":                   {"ApplicationId"},
			"get-application-policy":            {"ApplicationId"},
			"get-cloud-formation-template":      {"ApplicationId", "TemplateId"},
			"list-application-dependencies":     {"ApplicationId"},
			"list-application-versions":         {"ApplicationId"},
			"list-applications":                 {},
			"put-application-policy":            {"ApplicationId", "Statements"},
			"unshare-application":               {"ApplicationId", "OrganizationId"},
			"update-application":                {"ApplicationId"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("serverlessapplicationrepository", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
