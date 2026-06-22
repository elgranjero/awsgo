package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/servicecatalogappregistry"
)

var fields_associate_attribute_group = []leanruntime.Field{
	{Name: "Application", Flag: "application", Type: "*string", Required: true},
	{Name: "AttributeGroup", Flag: "attribute-group", Type: "*string", Required: true},
}

var fields_associate_resource = []leanruntime.Field{
	{Name: "Application", Flag: "application", Type: "*string", Required: true},
	{Name: "Options", Flag: "options", Type: "[]types.AssociationOption", Required: false},
	{Name: "Resource", Flag: "resource", Type: "*string", Required: true},
	{Name: "ResourceType", Flag: "resource-type", Type: "types.ResourceType", Required: true},
}

var fields_create_application = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_attribute_group = []leanruntime.Field{
	{Name: "Attributes", Flag: "attributes", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_delete_application = []leanruntime.Field{
	{Name: "Application", Flag: "application", Type: "*string", Required: true},
}

var fields_delete_attribute_group = []leanruntime.Field{
	{Name: "AttributeGroup", Flag: "attribute-group", Type: "*string", Required: true},
}

var fields_disassociate_attribute_group = []leanruntime.Field{
	{Name: "Application", Flag: "application", Type: "*string", Required: true},
	{Name: "AttributeGroup", Flag: "attribute-group", Type: "*string", Required: true},
}

var fields_disassociate_resource = []leanruntime.Field{
	{Name: "Application", Flag: "application", Type: "*string", Required: true},
	{Name: "Resource", Flag: "resource", Type: "*string", Required: true},
	{Name: "ResourceType", Flag: "resource-type", Type: "types.ResourceType", Required: true},
}

var fields_get_application = []leanruntime.Field{
	{Name: "Application", Flag: "application", Type: "*string", Required: true},
}

var fields_get_associated_resource = []leanruntime.Field{
	{Name: "Application", Flag: "application", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Resource", Flag: "resource", Type: "*string", Required: true},
	{Name: "ResourceTagStatus", Flag: "resource-tag-status", Type: "[]types.ResourceItemStatus", Required: false},
	{Name: "ResourceType", Flag: "resource-type", Type: "types.ResourceType", Required: true},
}

var fields_get_attribute_group = []leanruntime.Field{
	{Name: "AttributeGroup", Flag: "attribute-group", Type: "*string", Required: true},
}

var fields_get_configuration = []leanruntime.Field{}

var fields_list_applications = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_associated_attribute_groups = []leanruntime.Field{
	{Name: "Application", Flag: "application", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_associated_resources = []leanruntime.Field{
	{Name: "Application", Flag: "application", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_attribute_groups = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_attribute_groups_for_application = []leanruntime.Field{
	{Name: "Application", Flag: "application", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_put_configuration = []leanruntime.Field{
	{Name: "Configuration", Flag: "configuration", Type: "*types.AppRegistryConfiguration", Required: true},
}

var fields_sync_resource = []leanruntime.Field{
	{Name: "Resource", Flag: "resource", Type: "*string", Required: true},
	{Name: "ResourceType", Flag: "resource-type", Type: "types.ResourceType", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_application = []leanruntime.Field{
	{Name: "Application", Flag: "application", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_update_attribute_group = []leanruntime.Field{
	{Name: "AttributeGroup", Flag: "attribute-group", Type: "*string", Required: true},
	{Name: "Attributes", Flag: "attributes", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"associate-attribute-group": {
			Name:   "associate-attribute-group",
			Fields: fields_associate_attribute_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateAttributeGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_attribute_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateAttributeGroup(ctx, input)
			},
		},
		"associate-resource": {
			Name:   "associate-resource",
			Fields: fields_associate_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateResource(ctx, input)
			},
		},
		"create-application": {
			Name:   "create-application",
			Fields: fields_create_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateApplication(ctx, input)
			},
		},
		"create-attribute-group": {
			Name:   "create-attribute-group",
			Fields: fields_create_attribute_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAttributeGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_attribute_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAttributeGroup(ctx, input)
			},
		},
		"delete-application": {
			Name:   "delete-application",
			Fields: fields_delete_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteApplication(ctx, input)
			},
		},
		"delete-attribute-group": {
			Name:   "delete-attribute-group",
			Fields: fields_delete_attribute_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAttributeGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_attribute_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAttributeGroup(ctx, input)
			},
		},
		"disassociate-attribute-group": {
			Name:   "disassociate-attribute-group",
			Fields: fields_disassociate_attribute_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateAttributeGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_attribute_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateAttributeGroup(ctx, input)
			},
		},
		"disassociate-resource": {
			Name:   "disassociate-resource",
			Fields: fields_disassociate_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateResource(ctx, input)
			},
		},
		"get-application": {
			Name:   "get-application",
			Fields: fields_get_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetApplication(ctx, input)
			},
		},
		"get-associated-resource": {
			Name:   "get-associated-resource",
			Fields: fields_get_associated_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAssociatedResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_associated_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAssociatedResource(ctx, input)
			},
		},
		"get-attribute-group": {
			Name:   "get-attribute-group",
			Fields: fields_get_attribute_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAttributeGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_attribute_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAttributeGroup(ctx, input)
			},
		},
		"get-configuration": {
			Name:   "get-configuration",
			Fields: fields_get_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetConfiguration(ctx, input)
			},
		},
		"list-applications": {
			Name:   "list-applications",
			Fields: fields_list_applications,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListApplicationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_applications, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListApplications(ctx, input)
				}
				var results []*svc.ListApplicationsOutput
				p := svc.NewListApplicationsPaginator(client, input)
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
		"list-associated-attribute-groups": {
			Name:   "list-associated-attribute-groups",
			Fields: fields_list_associated_attribute_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAssociatedAttributeGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_associated_attribute_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAssociatedAttributeGroups(ctx, input)
				}
				var results []*svc.ListAssociatedAttributeGroupsOutput
				p := svc.NewListAssociatedAttributeGroupsPaginator(client, input)
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
		"list-associated-resources": {
			Name:   "list-associated-resources",
			Fields: fields_list_associated_resources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAssociatedResourcesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_associated_resources, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAssociatedResources(ctx, input)
				}
				var results []*svc.ListAssociatedResourcesOutput
				p := svc.NewListAssociatedResourcesPaginator(client, input)
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
		"list-attribute-groups": {
			Name:   "list-attribute-groups",
			Fields: fields_list_attribute_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAttributeGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_attribute_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAttributeGroups(ctx, input)
				}
				var results []*svc.ListAttributeGroupsOutput
				p := svc.NewListAttributeGroupsPaginator(client, input)
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
		"list-attribute-groups-for-application": {
			Name:   "list-attribute-groups-for-application",
			Fields: fields_list_attribute_groups_for_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAttributeGroupsForApplicationInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_attribute_groups_for_application, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAttributeGroupsForApplication(ctx, input)
				}
				var results []*svc.ListAttributeGroupsForApplicationOutput
				p := svc.NewListAttributeGroupsForApplicationPaginator(client, input)
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
		"put-configuration": {
			Name:   "put-configuration",
			Fields: fields_put_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutConfiguration(ctx, input)
			},
		},
		"sync-resource": {
			Name:   "sync-resource",
			Fields: fields_sync_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SyncResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_sync_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SyncResource(ctx, input)
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
		"update-application": {
			Name:   "update-application",
			Fields: fields_update_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateApplication(ctx, input)
			},
		},
		"update-attribute-group": {
			Name:   "update-attribute-group",
			Fields: fields_update_attribute_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAttributeGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_attribute_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAttributeGroup(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("servicecatalogappregistry", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
