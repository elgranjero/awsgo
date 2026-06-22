package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/workspacesinstances"
)

var fields_associate_volume = []leanruntime.Field{
	{Name: "Device", Flag: "device", Type: "*string", Required: true},
	{Name: "VolumeId", Flag: "volume-id", Type: "*string", Required: true},
	{Name: "WorkspaceInstanceId", Flag: "workspace-instance-id", Type: "*string", Required: true},
}

var fields_create_volume = []leanruntime.Field{
	{Name: "AvailabilityZone", Flag: "availability-zone", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Encrypted", Flag: "encrypted", Type: "*bool", Required: false},
	{Name: "Iops", Flag: "iops", Type: "*int32", Required: false},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "SizeInGB", Flag: "size-in-gb", Type: "*int32", Required: false},
	{Name: "SnapshotId", Flag: "snapshot-id", Type: "*string", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
	{Name: "Throughput", Flag: "throughput", Type: "*int32", Required: false},
	{Name: "VolumeType", Flag: "volume-type", Type: "types.VolumeTypeEnum", Required: false},
}

var fields_create_workspace_instance = []leanruntime.Field{
	{Name: "BillingConfiguration", Flag: "billing-configuration", Type: "*types.BillingConfiguration", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ManagedInstance", Flag: "managed-instance", Type: "*types.ManagedInstanceRequest", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_delete_volume = []leanruntime.Field{
	{Name: "VolumeId", Flag: "volume-id", Type: "*string", Required: true},
}

var fields_delete_workspace_instance = []leanruntime.Field{
	{Name: "WorkspaceInstanceId", Flag: "workspace-instance-id", Type: "*string", Required: true},
}

var fields_disassociate_volume = []leanruntime.Field{
	{Name: "Device", Flag: "device", Type: "*string", Required: false},
	{Name: "DisassociateMode", Flag: "disassociate-mode", Type: "types.DisassociateModeEnum", Required: false},
	{Name: "VolumeId", Flag: "volume-id", Type: "*string", Required: true},
	{Name: "WorkspaceInstanceId", Flag: "workspace-instance-id", Type: "*string", Required: true},
}

var fields_get_workspace_instance = []leanruntime.Field{
	{Name: "WorkspaceInstanceId", Flag: "workspace-instance-id", Type: "*string", Required: true},
}

var fields_list_instance_types = []leanruntime.Field{
	{Name: "InstanceConfigurationFilter", Flag: "instance-configuration-filter", Type: "*types.InstanceConfigurationFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_regions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "WorkspaceInstanceId", Flag: "workspace-instance-id", Type: "*string", Required: true},
}

var fields_list_workspace_instances = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ProvisionStates", Flag: "provision-states", Type: "[]types.ProvisionStateEnum", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
	{Name: "WorkspaceInstanceId", Flag: "workspace-instance-id", Type: "*string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
	{Name: "WorkspaceInstanceId", Flag: "workspace-instance-id", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"associate-volume": {
			Name:   "associate-volume",
			Fields: fields_associate_volume,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateVolumeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_volume, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateVolume(ctx, input)
			},
		},
		"create-volume": {
			Name:   "create-volume",
			Fields: fields_create_volume,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateVolumeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_volume, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateVolume(ctx, input)
			},
		},
		"create-workspace-instance": {
			Name:   "create-workspace-instance",
			Fields: fields_create_workspace_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateWorkspaceInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_workspace_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateWorkspaceInstance(ctx, input)
			},
		},
		"delete-volume": {
			Name:   "delete-volume",
			Fields: fields_delete_volume,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVolumeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_volume, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVolume(ctx, input)
			},
		},
		"delete-workspace-instance": {
			Name:   "delete-workspace-instance",
			Fields: fields_delete_workspace_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteWorkspaceInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_workspace_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteWorkspaceInstance(ctx, input)
			},
		},
		"disassociate-volume": {
			Name:   "disassociate-volume",
			Fields: fields_disassociate_volume,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateVolumeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_volume, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateVolume(ctx, input)
			},
		},
		"get-workspace-instance": {
			Name:   "get-workspace-instance",
			Fields: fields_get_workspace_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetWorkspaceInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_workspace_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetWorkspaceInstance(ctx, input)
			},
		},
		"list-instance-types": {
			Name:   "list-instance-types",
			Fields: fields_list_instance_types,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListInstanceTypesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_instance_types, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListInstanceTypes(ctx, input)
				}
				var results []*svc.ListInstanceTypesOutput
				p := svc.NewListInstanceTypesPaginator(client, input)
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
		"list-regions": {
			Name:   "list-regions",
			Fields: fields_list_regions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRegionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_regions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRegions(ctx, input)
				}
				var results []*svc.ListRegionsOutput
				p := svc.NewListRegionsPaginator(client, input)
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
		"list-workspace-instances": {
			Name:   "list-workspace-instances",
			Fields: fields_list_workspace_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListWorkspaceInstancesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_workspace_instances, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListWorkspaceInstances(ctx, input)
				}
				var results []*svc.ListWorkspaceInstancesOutput
				p := svc.NewListWorkspaceInstancesPaginator(client, input)
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
	}
	if err := leanruntime.Execute("workspacesinstances", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
