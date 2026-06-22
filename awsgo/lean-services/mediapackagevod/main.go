package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/mediapackagevod"
)

var fields_configure_logs = []leanruntime.Field{
	{Name: "EgressAccessLogs", Flag: "egress-access-logs", Type: "*types.EgressAccessLogs", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_create_asset = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "PackagingGroupId", Flag: "packaging-group-id", Type: "*string", Required: true},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: false},
	{Name: "SourceArn", Flag: "source-arn", Type: "*string", Required: true},
	{Name: "SourceRoleArn", Flag: "source-role-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_packaging_configuration = []leanruntime.Field{
	{Name: "CmafPackage", Flag: "cmaf-package", Type: "*types.CmafPackage", Required: false},
	{Name: "DashPackage", Flag: "dash-package", Type: "*types.DashPackage", Required: false},
	{Name: "HlsPackage", Flag: "hls-package", Type: "*types.HlsPackage", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "MssPackage", Flag: "mss-package", Type: "*types.MssPackage", Required: false},
	{Name: "PackagingGroupId", Flag: "packaging-group-id", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_packaging_group = []leanruntime.Field{
	{Name: "Authorization", Flag: "authorization", Type: "*types.Authorization", Required: false},
	{Name: "EgressAccessLogs", Flag: "egress-access-logs", Type: "*types.EgressAccessLogs", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_delete_asset = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_delete_packaging_configuration = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_delete_packaging_group = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_describe_asset = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_describe_packaging_configuration = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_describe_packaging_group = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_list_assets = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PackagingGroupId", Flag: "packaging-group-id", Type: "*string", Required: false},
}

var fields_list_packaging_configurations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PackagingGroupId", Flag: "packaging-group-id", Type: "*string", Required: false},
}

var fields_list_packaging_groups = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_packaging_group = []leanruntime.Field{
	{Name: "Authorization", Flag: "authorization", Type: "*types.Authorization", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"configure-logs": {
			Name:   "configure-logs",
			Fields: fields_configure_logs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ConfigureLogsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_configure_logs, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ConfigureLogs(ctx, input)
			},
		},
		"create-asset": {
			Name:   "create-asset",
			Fields: fields_create_asset,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAssetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_asset, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAsset(ctx, input)
			},
		},
		"create-packaging-configuration": {
			Name:   "create-packaging-configuration",
			Fields: fields_create_packaging_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePackagingConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_packaging_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePackagingConfiguration(ctx, input)
			},
		},
		"create-packaging-group": {
			Name:   "create-packaging-group",
			Fields: fields_create_packaging_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePackagingGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_packaging_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePackagingGroup(ctx, input)
			},
		},
		"delete-asset": {
			Name:   "delete-asset",
			Fields: fields_delete_asset,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAssetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_asset, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAsset(ctx, input)
			},
		},
		"delete-packaging-configuration": {
			Name:   "delete-packaging-configuration",
			Fields: fields_delete_packaging_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePackagingConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_packaging_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePackagingConfiguration(ctx, input)
			},
		},
		"delete-packaging-group": {
			Name:   "delete-packaging-group",
			Fields: fields_delete_packaging_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePackagingGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_packaging_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePackagingGroup(ctx, input)
			},
		},
		"describe-asset": {
			Name:   "describe-asset",
			Fields: fields_describe_asset,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAssetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_asset, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAsset(ctx, input)
			},
		},
		"describe-packaging-configuration": {
			Name:   "describe-packaging-configuration",
			Fields: fields_describe_packaging_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribePackagingConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_packaging_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribePackagingConfiguration(ctx, input)
			},
		},
		"describe-packaging-group": {
			Name:   "describe-packaging-group",
			Fields: fields_describe_packaging_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribePackagingGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_packaging_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribePackagingGroup(ctx, input)
			},
		},
		"list-assets": {
			Name:   "list-assets",
			Fields: fields_list_assets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAssetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_assets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAssets(ctx, input)
				}
				var results []*svc.ListAssetsOutput
				p := svc.NewListAssetsPaginator(client, input)
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
		"list-packaging-configurations": {
			Name:   "list-packaging-configurations",
			Fields: fields_list_packaging_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPackagingConfigurationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_packaging_configurations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPackagingConfigurations(ctx, input)
				}
				var results []*svc.ListPackagingConfigurationsOutput
				p := svc.NewListPackagingConfigurationsPaginator(client, input)
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
		"list-packaging-groups": {
			Name:   "list-packaging-groups",
			Fields: fields_list_packaging_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPackagingGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_packaging_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPackagingGroups(ctx, input)
				}
				var results []*svc.ListPackagingGroupsOutput
				p := svc.NewListPackagingGroupsPaginator(client, input)
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
		"update-packaging-group": {
			Name:   "update-packaging-group",
			Fields: fields_update_packaging_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePackagingGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_packaging_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePackagingGroup(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("mediapackagevod", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
