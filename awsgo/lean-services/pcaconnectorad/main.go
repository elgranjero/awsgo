package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/pcaconnectorad"
)

var fields_create_connector = []leanruntime.Field{
	{Name: "CertificateAuthorityArn", Flag: "certificate-authority-arn", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "VpcInformation", Flag: "vpc-information", Type: "*types.VpcInformation", Required: true},
}

var fields_create_directory_registration = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_service_principal_name = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ConnectorArn", Flag: "connector-arn", Type: "*string", Required: true},
	{Name: "DirectoryRegistrationArn", Flag: "directory-registration-arn", Type: "*string", Required: true},
}

var fields_create_template = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ConnectorArn", Flag: "connector-arn", Type: "*string", Required: true},
	{Name: "Definition", Flag: "definition", Type: "types.TemplateDefinition", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_template_group_access_control_entry = []leanruntime.Field{
	{Name: "AccessRights", Flag: "access-rights", Type: "*types.AccessRights", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "GroupDisplayName", Flag: "group-display-name", Type: "*string", Required: true},
	{Name: "GroupSecurityIdentifier", Flag: "group-security-identifier", Type: "*string", Required: true},
	{Name: "TemplateArn", Flag: "template-arn", Type: "*string", Required: true},
}

var fields_delete_connector = []leanruntime.Field{
	{Name: "ConnectorArn", Flag: "connector-arn", Type: "*string", Required: true},
}

var fields_delete_directory_registration = []leanruntime.Field{
	{Name: "DirectoryRegistrationArn", Flag: "directory-registration-arn", Type: "*string", Required: true},
}

var fields_delete_service_principal_name = []leanruntime.Field{
	{Name: "ConnectorArn", Flag: "connector-arn", Type: "*string", Required: true},
	{Name: "DirectoryRegistrationArn", Flag: "directory-registration-arn", Type: "*string", Required: true},
}

var fields_delete_template = []leanruntime.Field{
	{Name: "TemplateArn", Flag: "template-arn", Type: "*string", Required: true},
}

var fields_delete_template_group_access_control_entry = []leanruntime.Field{
	{Name: "GroupSecurityIdentifier", Flag: "group-security-identifier", Type: "*string", Required: true},
	{Name: "TemplateArn", Flag: "template-arn", Type: "*string", Required: true},
}

var fields_get_connector = []leanruntime.Field{
	{Name: "ConnectorArn", Flag: "connector-arn", Type: "*string", Required: true},
}

var fields_get_directory_registration = []leanruntime.Field{
	{Name: "DirectoryRegistrationArn", Flag: "directory-registration-arn", Type: "*string", Required: true},
}

var fields_get_service_principal_name = []leanruntime.Field{
	{Name: "ConnectorArn", Flag: "connector-arn", Type: "*string", Required: true},
	{Name: "DirectoryRegistrationArn", Flag: "directory-registration-arn", Type: "*string", Required: true},
}

var fields_get_template = []leanruntime.Field{
	{Name: "TemplateArn", Flag: "template-arn", Type: "*string", Required: true},
}

var fields_get_template_group_access_control_entry = []leanruntime.Field{
	{Name: "GroupSecurityIdentifier", Flag: "group-security-identifier", Type: "*string", Required: true},
	{Name: "TemplateArn", Flag: "template-arn", Type: "*string", Required: true},
}

var fields_list_connectors = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_directory_registrations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_service_principal_names = []leanruntime.Field{
	{Name: "DirectoryRegistrationArn", Flag: "directory-registration-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_template_group_access_control_entries = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TemplateArn", Flag: "template-arn", Type: "*string", Required: true},
}

var fields_list_templates = []leanruntime.Field{
	{Name: "ConnectorArn", Flag: "connector-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_template = []leanruntime.Field{
	{Name: "Definition", Flag: "definition", Type: "types.TemplateDefinition", Required: false},
	{Name: "ReenrollAllCertificateHolders", Flag: "reenroll-all-certificate-holders", Type: "*bool", Required: false},
	{Name: "TemplateArn", Flag: "template-arn", Type: "*string", Required: true},
}

var fields_update_template_group_access_control_entry = []leanruntime.Field{
	{Name: "AccessRights", Flag: "access-rights", Type: "*types.AccessRights", Required: false},
	{Name: "GroupDisplayName", Flag: "group-display-name", Type: "*string", Required: false},
	{Name: "GroupSecurityIdentifier", Flag: "group-security-identifier", Type: "*string", Required: true},
	{Name: "TemplateArn", Flag: "template-arn", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-connector": {
			Name:   "create-connector",
			Fields: fields_create_connector,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateConnectorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_connector, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateConnector(ctx, input)
			},
		},
		"create-directory-registration": {
			Name:   "create-directory-registration",
			Fields: fields_create_directory_registration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDirectoryRegistrationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_directory_registration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDirectoryRegistration(ctx, input)
			},
		},
		"create-service-principal-name": {
			Name:   "create-service-principal-name",
			Fields: fields_create_service_principal_name,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateServicePrincipalNameInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_service_principal_name, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateServicePrincipalName(ctx, input)
			},
		},
		"create-template": {
			Name:   "create-template",
			Fields: fields_create_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTemplate(ctx, input)
			},
		},
		"create-template-group-access-control-entry": {
			Name:   "create-template-group-access-control-entry",
			Fields: fields_create_template_group_access_control_entry,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTemplateGroupAccessControlEntryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_template_group_access_control_entry, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTemplateGroupAccessControlEntry(ctx, input)
			},
		},
		"delete-connector": {
			Name:   "delete-connector",
			Fields: fields_delete_connector,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteConnectorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_connector, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteConnector(ctx, input)
			},
		},
		"delete-directory-registration": {
			Name:   "delete-directory-registration",
			Fields: fields_delete_directory_registration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDirectoryRegistrationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_directory_registration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDirectoryRegistration(ctx, input)
			},
		},
		"delete-service-principal-name": {
			Name:   "delete-service-principal-name",
			Fields: fields_delete_service_principal_name,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteServicePrincipalNameInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_service_principal_name, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteServicePrincipalName(ctx, input)
			},
		},
		"delete-template": {
			Name:   "delete-template",
			Fields: fields_delete_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTemplate(ctx, input)
			},
		},
		"delete-template-group-access-control-entry": {
			Name:   "delete-template-group-access-control-entry",
			Fields: fields_delete_template_group_access_control_entry,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTemplateGroupAccessControlEntryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_template_group_access_control_entry, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTemplateGroupAccessControlEntry(ctx, input)
			},
		},
		"get-connector": {
			Name:   "get-connector",
			Fields: fields_get_connector,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetConnectorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_connector, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetConnector(ctx, input)
			},
		},
		"get-directory-registration": {
			Name:   "get-directory-registration",
			Fields: fields_get_directory_registration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDirectoryRegistrationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_directory_registration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDirectoryRegistration(ctx, input)
			},
		},
		"get-service-principal-name": {
			Name:   "get-service-principal-name",
			Fields: fields_get_service_principal_name,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetServicePrincipalNameInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_service_principal_name, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetServicePrincipalName(ctx, input)
			},
		},
		"get-template": {
			Name:   "get-template",
			Fields: fields_get_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTemplate(ctx, input)
			},
		},
		"get-template-group-access-control-entry": {
			Name:   "get-template-group-access-control-entry",
			Fields: fields_get_template_group_access_control_entry,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTemplateGroupAccessControlEntryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_template_group_access_control_entry, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTemplateGroupAccessControlEntry(ctx, input)
			},
		},
		"list-connectors": {
			Name:   "list-connectors",
			Fields: fields_list_connectors,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListConnectorsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_connectors, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListConnectors(ctx, input)
				}
				var results []*svc.ListConnectorsOutput
				p := svc.NewListConnectorsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-directory-registrations": {
			Name:   "list-directory-registrations",
			Fields: fields_list_directory_registrations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDirectoryRegistrationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_directory_registrations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDirectoryRegistrations(ctx, input)
				}
				var results []*svc.ListDirectoryRegistrationsOutput
				p := svc.NewListDirectoryRegistrationsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-service-principal-names": {
			Name:   "list-service-principal-names",
			Fields: fields_list_service_principal_names,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListServicePrincipalNamesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_service_principal_names, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListServicePrincipalNames(ctx, input)
				}
				var results []*svc.ListServicePrincipalNamesOutput
				p := svc.NewListServicePrincipalNamesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-tags-for-resource": {
			Name:   "list-tags-for-resource",
			Fields: fields_list_tags_for_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTagsForResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_tags_for_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListTagsForResource(ctx, input)
			},
		},
		"list-template-group-access-control-entries": {
			Name:   "list-template-group-access-control-entries",
			Fields: fields_list_template_group_access_control_entries,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTemplateGroupAccessControlEntriesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_template_group_access_control_entries, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTemplateGroupAccessControlEntries(ctx, input)
				}
				var results []*svc.ListTemplateGroupAccessControlEntriesOutput
				p := svc.NewListTemplateGroupAccessControlEntriesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-templates": {
			Name:   "list-templates",
			Fields: fields_list_templates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTemplatesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_templates, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTemplates(ctx, input)
				}
				var results []*svc.ListTemplatesOutput
				p := svc.NewListTemplatesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"tag-resource": {
			Name:   "tag-resource",
			Fields: fields_tag_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TagResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_tag_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TagResource(ctx, input)
			},
		},
		"untag-resource": {
			Name:   "untag-resource",
			Fields: fields_untag_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UntagResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_untag_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UntagResource(ctx, input)
			},
		},
		"update-template": {
			Name:   "update-template",
			Fields: fields_update_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTemplate(ctx, input)
			},
		},
		"update-template-group-access-control-entry": {
			Name:   "update-template-group-access-control-entry",
			Fields: fields_update_template_group_access_control_entry,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTemplateGroupAccessControlEntryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_template_group_access_control_entry, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTemplateGroupAccessControlEntry(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("pcaconnectorad", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
