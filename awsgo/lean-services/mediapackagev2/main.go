package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/mediapackagev2"
)

var fields_cancel_harvest_job = []leanruntime.Field{
	{Name: "ChannelGroupName", Flag: "channel-group-name", Type: "*string", Required: true},
	{Name: "ChannelName", Flag: "channel-name", Type: "*string", Required: true},
	{Name: "ETag", Flag: "etag", Type: "*string", Required: false},
	{Name: "HarvestJobName", Flag: "harvest-job-name", Type: "*string", Required: true},
	{Name: "OriginEndpointName", Flag: "origin-endpoint-name", Type: "*string", Required: true},
}

var fields_create_channel = []leanruntime.Field{
	{Name: "ChannelGroupName", Flag: "channel-group-name", Type: "*string", Required: true},
	{Name: "ChannelName", Flag: "channel-name", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "InputSwitchConfiguration", Flag: "input-switch-configuration", Type: "*types.InputSwitchConfiguration", Required: false},
	{Name: "InputType", Flag: "input-type", Type: "types.InputType", Required: false},
	{Name: "OutputHeaderConfiguration", Flag: "output-header-configuration", Type: "*types.OutputHeaderConfiguration", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_channel_group = []leanruntime.Field{
	{Name: "ChannelGroupName", Flag: "channel-group-name", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_harvest_job = []leanruntime.Field{
	{Name: "ChannelGroupName", Flag: "channel-group-name", Type: "*string", Required: true},
	{Name: "ChannelName", Flag: "channel-name", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Destination", Flag: "destination", Type: "*types.Destination", Required: true},
	{Name: "HarvestJobName", Flag: "harvest-job-name", Type: "*string", Required: false},
	{Name: "HarvestedManifests", Flag: "harvested-manifests", Type: "*types.HarvestedManifests", Required: true},
	{Name: "OriginEndpointName", Flag: "origin-endpoint-name", Type: "*string", Required: true},
	{Name: "ScheduleConfiguration", Flag: "schedule-configuration", Type: "*types.HarvesterScheduleConfiguration", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_origin_endpoint = []leanruntime.Field{
	{Name: "ChannelGroupName", Flag: "channel-group-name", Type: "*string", Required: true},
	{Name: "ChannelName", Flag: "channel-name", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ContainerType", Flag: "container-type", Type: "types.ContainerType", Required: true},
	{Name: "DashManifests", Flag: "dash-manifests", Type: "[]types.CreateDashManifestConfiguration", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ForceEndpointErrorConfiguration", Flag: "force-endpoint-error-configuration", Type: "*types.ForceEndpointErrorConfiguration", Required: false},
	{Name: "HlsManifests", Flag: "hls-manifests", Type: "[]types.CreateHlsManifestConfiguration", Required: false},
	{Name: "LowLatencyHlsManifests", Flag: "low-latency-hls-manifests", Type: "[]types.CreateLowLatencyHlsManifestConfiguration", Required: false},
	{Name: "MssManifests", Flag: "mss-manifests", Type: "[]types.CreateMssManifestConfiguration", Required: false},
	{Name: "OriginEndpointName", Flag: "origin-endpoint-name", Type: "*string", Required: true},
	{Name: "Segment", Flag: "segment", Type: "*types.Segment", Required: false},
	{Name: "StartoverWindowSeconds", Flag: "startover-window-seconds", Type: "*int32", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_delete_channel = []leanruntime.Field{
	{Name: "ChannelGroupName", Flag: "channel-group-name", Type: "*string", Required: true},
	{Name: "ChannelName", Flag: "channel-name", Type: "*string", Required: true},
}

var fields_delete_channel_group = []leanruntime.Field{
	{Name: "ChannelGroupName", Flag: "channel-group-name", Type: "*string", Required: true},
}

var fields_delete_channel_policy = []leanruntime.Field{
	{Name: "ChannelGroupName", Flag: "channel-group-name", Type: "*string", Required: true},
	{Name: "ChannelName", Flag: "channel-name", Type: "*string", Required: true},
}

var fields_delete_origin_endpoint = []leanruntime.Field{
	{Name: "ChannelGroupName", Flag: "channel-group-name", Type: "*string", Required: true},
	{Name: "ChannelName", Flag: "channel-name", Type: "*string", Required: true},
	{Name: "OriginEndpointName", Flag: "origin-endpoint-name", Type: "*string", Required: true},
}

var fields_delete_origin_endpoint_policy = []leanruntime.Field{
	{Name: "ChannelGroupName", Flag: "channel-group-name", Type: "*string", Required: true},
	{Name: "ChannelName", Flag: "channel-name", Type: "*string", Required: true},
	{Name: "OriginEndpointName", Flag: "origin-endpoint-name", Type: "*string", Required: true},
}

var fields_get_channel = []leanruntime.Field{
	{Name: "ChannelGroupName", Flag: "channel-group-name", Type: "*string", Required: true},
	{Name: "ChannelName", Flag: "channel-name", Type: "*string", Required: true},
}

var fields_get_channel_group = []leanruntime.Field{
	{Name: "ChannelGroupName", Flag: "channel-group-name", Type: "*string", Required: true},
}

var fields_get_channel_policy = []leanruntime.Field{
	{Name: "ChannelGroupName", Flag: "channel-group-name", Type: "*string", Required: true},
	{Name: "ChannelName", Flag: "channel-name", Type: "*string", Required: true},
}

var fields_get_harvest_job = []leanruntime.Field{
	{Name: "ChannelGroupName", Flag: "channel-group-name", Type: "*string", Required: true},
	{Name: "ChannelName", Flag: "channel-name", Type: "*string", Required: true},
	{Name: "HarvestJobName", Flag: "harvest-job-name", Type: "*string", Required: true},
	{Name: "OriginEndpointName", Flag: "origin-endpoint-name", Type: "*string", Required: true},
}

var fields_get_origin_endpoint = []leanruntime.Field{
	{Name: "ChannelGroupName", Flag: "channel-group-name", Type: "*string", Required: true},
	{Name: "ChannelName", Flag: "channel-name", Type: "*string", Required: true},
	{Name: "OriginEndpointName", Flag: "origin-endpoint-name", Type: "*string", Required: true},
}

var fields_get_origin_endpoint_policy = []leanruntime.Field{
	{Name: "ChannelGroupName", Flag: "channel-group-name", Type: "*string", Required: true},
	{Name: "ChannelName", Flag: "channel-name", Type: "*string", Required: true},
	{Name: "OriginEndpointName", Flag: "origin-endpoint-name", Type: "*string", Required: true},
}

var fields_list_channel_groups = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_channels = []leanruntime.Field{
	{Name: "ChannelGroupName", Flag: "channel-group-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_harvest_jobs = []leanruntime.Field{
	{Name: "ChannelGroupName", Flag: "channel-group-name", Type: "*string", Required: true},
	{Name: "ChannelName", Flag: "channel-name", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OriginEndpointName", Flag: "origin-endpoint-name", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.HarvestJobStatus", Required: false},
}

var fields_list_origin_endpoints = []leanruntime.Field{
	{Name: "ChannelGroupName", Flag: "channel-group-name", Type: "*string", Required: true},
	{Name: "ChannelName", Flag: "channel-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_put_channel_policy = []leanruntime.Field{
	{Name: "ChannelGroupName", Flag: "channel-group-name", Type: "*string", Required: true},
	{Name: "ChannelName", Flag: "channel-name", Type: "*string", Required: true},
	{Name: "Policy", Flag: "policy", Type: "*string", Required: true},
}

var fields_put_origin_endpoint_policy = []leanruntime.Field{
	{Name: "CdnAuthConfiguration", Flag: "cdn-auth-configuration", Type: "*types.CdnAuthConfiguration", Required: false},
	{Name: "ChannelGroupName", Flag: "channel-group-name", Type: "*string", Required: true},
	{Name: "ChannelName", Flag: "channel-name", Type: "*string", Required: true},
	{Name: "OriginEndpointName", Flag: "origin-endpoint-name", Type: "*string", Required: true},
	{Name: "Policy", Flag: "policy", Type: "*string", Required: true},
}

var fields_reset_channel_state = []leanruntime.Field{
	{Name: "ChannelGroupName", Flag: "channel-group-name", Type: "*string", Required: true},
	{Name: "ChannelName", Flag: "channel-name", Type: "*string", Required: true},
}

var fields_reset_origin_endpoint_state = []leanruntime.Field{
	{Name: "ChannelGroupName", Flag: "channel-group-name", Type: "*string", Required: true},
	{Name: "ChannelName", Flag: "channel-name", Type: "*string", Required: true},
	{Name: "OriginEndpointName", Flag: "origin-endpoint-name", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_channel = []leanruntime.Field{
	{Name: "ChannelGroupName", Flag: "channel-group-name", Type: "*string", Required: true},
	{Name: "ChannelName", Flag: "channel-name", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ETag", Flag: "etag", Type: "*string", Required: false},
	{Name: "InputSwitchConfiguration", Flag: "input-switch-configuration", Type: "*types.InputSwitchConfiguration", Required: false},
	{Name: "OutputHeaderConfiguration", Flag: "output-header-configuration", Type: "*types.OutputHeaderConfiguration", Required: false},
}

var fields_update_channel_group = []leanruntime.Field{
	{Name: "ChannelGroupName", Flag: "channel-group-name", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ETag", Flag: "etag", Type: "*string", Required: false},
}

var fields_update_origin_endpoint = []leanruntime.Field{
	{Name: "ChannelGroupName", Flag: "channel-group-name", Type: "*string", Required: true},
	{Name: "ChannelName", Flag: "channel-name", Type: "*string", Required: true},
	{Name: "ContainerType", Flag: "container-type", Type: "types.ContainerType", Required: true},
	{Name: "DashManifests", Flag: "dash-manifests", Type: "[]types.CreateDashManifestConfiguration", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ETag", Flag: "etag", Type: "*string", Required: false},
	{Name: "ForceEndpointErrorConfiguration", Flag: "force-endpoint-error-configuration", Type: "*types.ForceEndpointErrorConfiguration", Required: false},
	{Name: "HlsManifests", Flag: "hls-manifests", Type: "[]types.CreateHlsManifestConfiguration", Required: false},
	{Name: "LowLatencyHlsManifests", Flag: "low-latency-hls-manifests", Type: "[]types.CreateLowLatencyHlsManifestConfiguration", Required: false},
	{Name: "MssManifests", Flag: "mss-manifests", Type: "[]types.CreateMssManifestConfiguration", Required: false},
	{Name: "OriginEndpointName", Flag: "origin-endpoint-name", Type: "*string", Required: true},
	{Name: "Segment", Flag: "segment", Type: "*types.Segment", Required: false},
	{Name: "StartoverWindowSeconds", Flag: "startover-window-seconds", Type: "*int32", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"cancel-harvest-job": {
			Name:   "cancel-harvest-job",
			Fields: fields_cancel_harvest_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelHarvestJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_harvest_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelHarvestJob(ctx, input)
			},
		},
		"create-channel": {
			Name:   "create-channel",
			Fields: fields_create_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateChannel(ctx, input)
			},
		},
		"create-channel-group": {
			Name:   "create-channel-group",
			Fields: fields_create_channel_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateChannelGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_channel_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateChannelGroup(ctx, input)
			},
		},
		"create-harvest-job": {
			Name:   "create-harvest-job",
			Fields: fields_create_harvest_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateHarvestJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_harvest_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateHarvestJob(ctx, input)
			},
		},
		"create-origin-endpoint": {
			Name:   "create-origin-endpoint",
			Fields: fields_create_origin_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateOriginEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_origin_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateOriginEndpoint(ctx, input)
			},
		},
		"delete-channel": {
			Name:   "delete-channel",
			Fields: fields_delete_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteChannel(ctx, input)
			},
		},
		"delete-channel-group": {
			Name:   "delete-channel-group",
			Fields: fields_delete_channel_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteChannelGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_channel_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteChannelGroup(ctx, input)
			},
		},
		"delete-channel-policy": {
			Name:   "delete-channel-policy",
			Fields: fields_delete_channel_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteChannelPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_channel_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteChannelPolicy(ctx, input)
			},
		},
		"delete-origin-endpoint": {
			Name:   "delete-origin-endpoint",
			Fields: fields_delete_origin_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteOriginEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_origin_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteOriginEndpoint(ctx, input)
			},
		},
		"delete-origin-endpoint-policy": {
			Name:   "delete-origin-endpoint-policy",
			Fields: fields_delete_origin_endpoint_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteOriginEndpointPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_origin_endpoint_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteOriginEndpointPolicy(ctx, input)
			},
		},
		"get-channel": {
			Name:   "get-channel",
			Fields: fields_get_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetChannel(ctx, input)
			},
		},
		"get-channel-group": {
			Name:   "get-channel-group",
			Fields: fields_get_channel_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetChannelGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_channel_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetChannelGroup(ctx, input)
			},
		},
		"get-channel-policy": {
			Name:   "get-channel-policy",
			Fields: fields_get_channel_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetChannelPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_channel_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetChannelPolicy(ctx, input)
			},
		},
		"get-harvest-job": {
			Name:   "get-harvest-job",
			Fields: fields_get_harvest_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetHarvestJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_harvest_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetHarvestJob(ctx, input)
			},
		},
		"get-origin-endpoint": {
			Name:   "get-origin-endpoint",
			Fields: fields_get_origin_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetOriginEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_origin_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetOriginEndpoint(ctx, input)
			},
		},
		"get-origin-endpoint-policy": {
			Name:   "get-origin-endpoint-policy",
			Fields: fields_get_origin_endpoint_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetOriginEndpointPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_origin_endpoint_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetOriginEndpointPolicy(ctx, input)
			},
		},
		"list-channel-groups": {
			Name:   "list-channel-groups",
			Fields: fields_list_channel_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListChannelGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_channel_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListChannelGroups(ctx, input)
				}
				var results []*svc.ListChannelGroupsOutput
				p := svc.NewListChannelGroupsPaginator(client, input)
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
		"list-channels": {
			Name:   "list-channels",
			Fields: fields_list_channels,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListChannelsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_channels, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListChannels(ctx, input)
				}
				var results []*svc.ListChannelsOutput
				p := svc.NewListChannelsPaginator(client, input)
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
		"list-harvest-jobs": {
			Name:   "list-harvest-jobs",
			Fields: fields_list_harvest_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListHarvestJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_harvest_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListHarvestJobs(ctx, input)
				}
				var results []*svc.ListHarvestJobsOutput
				p := svc.NewListHarvestJobsPaginator(client, input)
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
		"list-origin-endpoints": {
			Name:   "list-origin-endpoints",
			Fields: fields_list_origin_endpoints,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListOriginEndpointsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_origin_endpoints, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListOriginEndpoints(ctx, input)
				}
				var results []*svc.ListOriginEndpointsOutput
				p := svc.NewListOriginEndpointsPaginator(client, input)
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
		"put-channel-policy": {
			Name:   "put-channel-policy",
			Fields: fields_put_channel_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutChannelPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_channel_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutChannelPolicy(ctx, input)
			},
		},
		"put-origin-endpoint-policy": {
			Name:   "put-origin-endpoint-policy",
			Fields: fields_put_origin_endpoint_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutOriginEndpointPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_origin_endpoint_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutOriginEndpointPolicy(ctx, input)
			},
		},
		"reset-channel-state": {
			Name:   "reset-channel-state",
			Fields: fields_reset_channel_state,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ResetChannelStateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reset_channel_state, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ResetChannelState(ctx, input)
			},
		},
		"reset-origin-endpoint-state": {
			Name:   "reset-origin-endpoint-state",
			Fields: fields_reset_origin_endpoint_state,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ResetOriginEndpointStateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reset_origin_endpoint_state, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ResetOriginEndpointState(ctx, input)
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
		"update-channel": {
			Name:   "update-channel",
			Fields: fields_update_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateChannel(ctx, input)
			},
		},
		"update-channel-group": {
			Name:   "update-channel-group",
			Fields: fields_update_channel_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateChannelGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_channel_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateChannelGroup(ctx, input)
			},
		},
		"update-origin-endpoint": {
			Name:   "update-origin-endpoint",
			Fields: fields_update_origin_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateOriginEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_origin_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateOriginEndpoint(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("mediapackagev2", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
