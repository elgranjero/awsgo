package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/mediapackage"
)

var fields_configure_logs = []leanruntime.Field{
	{Name: "EgressAccessLogs", Flag: "egress-access-logs", Type: "*types.EgressAccessLogs", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IngressAccessLogs", Flag: "ingress-access-logs", Type: "*types.IngressAccessLogs", Required: false},
}

var fields_create_channel = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_harvest_job = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*string", Required: true},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "OriginEndpointId", Flag: "origin-endpoint-id", Type: "*string", Required: true},
	{Name: "S3Destination", Flag: "s3-destination", Type: "*types.S3Destination", Required: true},
	{Name: "StartTime", Flag: "start-time", Type: "*string", Required: true},
}

var fields_create_origin_endpoint = []leanruntime.Field{
	{Name: "Authorization", Flag: "authorization", Type: "*types.Authorization", Required: false},
	{Name: "ChannelId", Flag: "channel-id", Type: "*string", Required: true},
	{Name: "CmafPackage", Flag: "cmaf-package", Type: "*types.CmafPackageCreateOrUpdateParameters", Required: false},
	{Name: "DashPackage", Flag: "dash-package", Type: "*types.DashPackage", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "HlsPackage", Flag: "hls-package", Type: "*types.HlsPackage", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "ManifestName", Flag: "manifest-name", Type: "*string", Required: false},
	{Name: "MssPackage", Flag: "mss-package", Type: "*types.MssPackage", Required: false},
	{Name: "Origination", Flag: "origination", Type: "types.Origination", Required: false},
	{Name: "StartoverWindowSeconds", Flag: "startover-window-seconds", Type: "*int32", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "TimeDelaySeconds", Flag: "time-delay-seconds", Type: "*int32", Required: false},
	{Name: "Whitelist", Flag: "whitelist", Type: "[]string", Required: false},
}

var fields_delete_channel = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_delete_origin_endpoint = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_describe_channel = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_describe_harvest_job = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_describe_origin_endpoint = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_list_channels = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_harvest_jobs = []leanruntime.Field{
	{Name: "IncludeChannelId", Flag: "include-channel-id", Type: "*string", Required: false},
	{Name: "IncludeStatus", Flag: "include-status", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_origin_endpoints = []leanruntime.Field{
	{Name: "ChannelId", Flag: "channel-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_rotate_channel_credentials = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_rotate_ingest_endpoint_credentials = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IngestEndpointId", Flag: "ingest-endpoint-id", Type: "*string", Required: true},
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
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_update_origin_endpoint = []leanruntime.Field{
	{Name: "Authorization", Flag: "authorization", Type: "*types.Authorization", Required: false},
	{Name: "CmafPackage", Flag: "cmaf-package", Type: "*types.CmafPackageCreateOrUpdateParameters", Required: false},
	{Name: "DashPackage", Flag: "dash-package", Type: "*types.DashPackage", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "HlsPackage", Flag: "hls-package", Type: "*types.HlsPackage", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "ManifestName", Flag: "manifest-name", Type: "*string", Required: false},
	{Name: "MssPackage", Flag: "mss-package", Type: "*types.MssPackage", Required: false},
	{Name: "Origination", Flag: "origination", Type: "types.Origination", Required: false},
	{Name: "StartoverWindowSeconds", Flag: "startover-window-seconds", Type: "*int32", Required: false},
	{Name: "TimeDelaySeconds", Flag: "time-delay-seconds", Type: "*int32", Required: false},
	{Name: "Whitelist", Flag: "whitelist", Type: "[]string", Required: false},
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
		"describe-channel": {
			Name:   "describe-channel",
			Fields: fields_describe_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeChannel(ctx, input)
			},
		},
		"describe-harvest-job": {
			Name:   "describe-harvest-job",
			Fields: fields_describe_harvest_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeHarvestJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_harvest_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeHarvestJob(ctx, input)
			},
		},
		"describe-origin-endpoint": {
			Name:   "describe-origin-endpoint",
			Fields: fields_describe_origin_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeOriginEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_origin_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeOriginEndpoint(ctx, input)
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
		"rotate-channel-credentials": {
			Name:   "rotate-channel-credentials",
			Fields: fields_rotate_channel_credentials,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RotateChannelCredentialsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_rotate_channel_credentials, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RotateChannelCredentials(ctx, input)
			},
		},
		"rotate-ingest-endpoint-credentials": {
			Name:   "rotate-ingest-endpoint-credentials",
			Fields: fields_rotate_ingest_endpoint_credentials,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RotateIngestEndpointCredentialsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_rotate_ingest_endpoint_credentials, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RotateIngestEndpointCredentials(ctx, input)
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
	if err := leanruntime.Execute("mediapackage", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
