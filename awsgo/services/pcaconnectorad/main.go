package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/pcaconnectorad/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"create-connector", "create-directory-registration", "create-service-principal-name", "create-template", "create-template-group-access-control-entry", "delete-connector", "delete-directory-registration", "delete-service-principal-name", "delete-template", "delete-template-group-access-control-entry", "get-connector", "get-directory-registration", "get-service-principal-name", "get-template", "get-template-group-access-control-entry", "list-connectors", "list-directory-registrations", "list-service-principal-names", "list-tags-for-resource", "list-template-group-access-control-entries", "list-templates", "tag-resource", "untag-resource", "update-template", "update-template-group-access-control-entry"},
		OperationSet: map[string]bool{"create-connector": true, "create-directory-registration": true, "create-service-principal-name": true, "create-template": true, "create-template-group-access-control-entry": true, "delete-connector": true, "delete-directory-registration": true, "delete-service-principal-name": true, "delete-template": true, "delete-template-group-access-control-entry": true, "get-connector": true, "get-directory-registration": true, "get-service-principal-name": true, "get-template": true, "get-template-group-access-control-entry": true, "list-connectors": true, "list-directory-registrations": true, "list-service-principal-names": true, "list-tags-for-resource": true, "list-template-group-access-control-entries": true, "list-templates": true, "tag-resource": true, "untag-resource": true, "update-template": true, "update-template-group-access-control-entry": true},
		OperationInputs: map[string][]string{
			"create-connector":                           {"CertificateAuthorityArn", "ClientToken", "DirectoryId", "Tags", "VpcInformation"},
			"create-directory-registration":              {"ClientToken", "DirectoryId", "Tags"},
			"create-service-principal-name":              {"ClientToken", "ConnectorArn", "DirectoryRegistrationArn"},
			"create-template":                            {"ClientToken", "ConnectorArn", "Definition", "Name", "Tags"},
			"create-template-group-access-control-entry": {"AccessRights", "ClientToken", "GroupDisplayName", "GroupSecurityIdentifier", "TemplateArn"},
			"delete-connector":                           {"ConnectorArn"},
			"delete-directory-registration":              {"DirectoryRegistrationArn"},
			"delete-service-principal-name":              {"ConnectorArn", "DirectoryRegistrationArn"},
			"delete-template":                            {"TemplateArn"},
			"delete-template-group-access-control-entry": {"GroupSecurityIdentifier", "TemplateArn"},
			"get-connector":                              {"ConnectorArn"},
			"get-directory-registration":                 {"DirectoryRegistrationArn"},
			"get-service-principal-name":                 {"ConnectorArn", "DirectoryRegistrationArn"},
			"get-template":                               {"TemplateArn"},
			"get-template-group-access-control-entry":    {"GroupSecurityIdentifier", "TemplateArn"},
			"list-connectors":                            {"MaxResults", "NextToken"},
			"list-directory-registrations":               {"MaxResults", "NextToken"},
			"list-service-principal-names":               {"DirectoryRegistrationArn", "MaxResults", "NextToken"},
			"list-tags-for-resource":                     {"ResourceArn"},
			"list-template-group-access-control-entries": {"MaxResults", "NextToken", "TemplateArn"},
			"list-templates":                             {"ConnectorArn", "MaxResults", "NextToken"},
			"tag-resource":                               {"ResourceArn", "Tags"},
			"untag-resource":                             {"ResourceArn", "TagKeys"},
			"update-template":                            {"Definition", "ReenrollAllCertificateHolders", "TemplateArn"},
			"update-template-group-access-control-entry": {"AccessRights", "GroupDisplayName", "GroupSecurityIdentifier", "TemplateArn"},
		},
		OperationInputTypes: map[string]map[string]string{
			"create-connector":                           {"CertificateAuthorityArn": "*string", "ClientToken": "*string", "DirectoryId": "*string", "Tags": "map[string]string", "VpcInformation": "*types.VpcInformation"},
			"create-directory-registration":              {"ClientToken": "*string", "DirectoryId": "*string", "Tags": "map[string]string"},
			"create-service-principal-name":              {"ClientToken": "*string", "ConnectorArn": "*string", "DirectoryRegistrationArn": "*string"},
			"create-template":                            {"ClientToken": "*string", "ConnectorArn": "*string", "Definition": "types.TemplateDefinition", "Name": "*string", "Tags": "map[string]string"},
			"create-template-group-access-control-entry": {"AccessRights": "*types.AccessRights", "ClientToken": "*string", "GroupDisplayName": "*string", "GroupSecurityIdentifier": "*string", "TemplateArn": "*string"},
			"delete-connector":                           {"ConnectorArn": "*string"},
			"delete-directory-registration":              {"DirectoryRegistrationArn": "*string"},
			"delete-service-principal-name":              {"ConnectorArn": "*string", "DirectoryRegistrationArn": "*string"},
			"delete-template":                            {"TemplateArn": "*string"},
			"delete-template-group-access-control-entry": {"GroupSecurityIdentifier": "*string", "TemplateArn": "*string"},
			"get-connector":                              {"ConnectorArn": "*string"},
			"get-directory-registration":                 {"DirectoryRegistrationArn": "*string"},
			"get-service-principal-name":                 {"ConnectorArn": "*string", "DirectoryRegistrationArn": "*string"},
			"get-template":                               {"TemplateArn": "*string"},
			"get-template-group-access-control-entry":    {"GroupSecurityIdentifier": "*string", "TemplateArn": "*string"},
			"list-connectors":                            {"MaxResults": "*int32", "NextToken": "*string"},
			"list-directory-registrations":               {"MaxResults": "*int32", "NextToken": "*string"},
			"list-service-principal-names":               {"DirectoryRegistrationArn": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"list-tags-for-resource":                     {"ResourceArn": "*string"},
			"list-template-group-access-control-entries": {"MaxResults": "*int32", "NextToken": "*string", "TemplateArn": "*string"},
			"list-templates":                             {"ConnectorArn": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"tag-resource":                               {"ResourceArn": "*string", "Tags": "map[string]string"},
			"untag-resource":                             {"ResourceArn": "*string", "TagKeys": "[]string"},
			"update-template":                            {"Definition": "types.TemplateDefinition", "ReenrollAllCertificateHolders": "*bool", "TemplateArn": "*string"},
			"update-template-group-access-control-entry": {"AccessRights": "*types.AccessRights", "GroupDisplayName": "*string", "GroupSecurityIdentifier": "*string", "TemplateArn": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"create-connector":                           {"CertificateAuthorityArn", "DirectoryId", "VpcInformation"},
			"create-directory-registration":              {"DirectoryId"},
			"create-service-principal-name":              {"ConnectorArn", "DirectoryRegistrationArn"},
			"create-template":                            {"ConnectorArn", "Definition", "Name"},
			"create-template-group-access-control-entry": {"AccessRights", "GroupDisplayName", "GroupSecurityIdentifier", "TemplateArn"},
			"delete-connector":                           {"ConnectorArn"},
			"delete-directory-registration":              {"DirectoryRegistrationArn"},
			"delete-service-principal-name":              {"ConnectorArn", "DirectoryRegistrationArn"},
			"delete-template":                            {"TemplateArn"},
			"delete-template-group-access-control-entry": {"GroupSecurityIdentifier", "TemplateArn"},
			"get-connector":                              {"ConnectorArn"},
			"get-directory-registration":                 {"DirectoryRegistrationArn"},
			"get-service-principal-name":                 {"ConnectorArn", "DirectoryRegistrationArn"},
			"get-template":                               {"TemplateArn"},
			"get-template-group-access-control-entry":    {"GroupSecurityIdentifier", "TemplateArn"},
			"list-connectors":                            {},
			"list-directory-registrations":               {},
			"list-service-principal-names":               {"DirectoryRegistrationArn"},
			"list-tags-for-resource":                     {"ResourceArn"},
			"list-template-group-access-control-entries": {"TemplateArn"},
			"list-templates":                             {"ConnectorArn"},
			"tag-resource":                               {"ResourceArn", "Tags"},
			"untag-resource":                             {"ResourceArn", "TagKeys"},
			"update-template":                            {"TemplateArn"},
			"update-template-group-access-control-entry": {"GroupSecurityIdentifier", "TemplateArn"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("pcaconnectorad", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
