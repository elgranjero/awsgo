package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/timestreamquery"
)

var fields_cancel_query = []leanruntime.Field{
	{Name: "QueryId", Flag: "query-id", Type: "*string", Required: true},
}

var fields_create_scheduled_query = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ErrorReportConfiguration", Flag: "error-report-configuration", Type: "*types.ErrorReportConfiguration", Required: true},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NotificationConfiguration", Flag: "notification-configuration", Type: "*types.NotificationConfiguration", Required: true},
	{Name: "QueryString", Flag: "query-string", Type: "*string", Required: true},
	{Name: "ScheduleConfiguration", Flag: "schedule-configuration", Type: "*types.ScheduleConfiguration", Required: true},
	{Name: "ScheduledQueryExecutionRoleArn", Flag: "scheduled-query-execution-role-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TargetConfiguration", Flag: "target-configuration", Type: "*types.TargetConfiguration", Required: false},
}

var fields_delete_scheduled_query = []leanruntime.Field{
	{Name: "ScheduledQueryArn", Flag: "scheduled-query-arn", Type: "*string", Required: true},
}

var fields_describe_account_settings = []leanruntime.Field{}

var fields_describe_endpoints = []leanruntime.Field{}

var fields_describe_scheduled_query = []leanruntime.Field{
	{Name: "ScheduledQueryArn", Flag: "scheduled-query-arn", Type: "*string", Required: true},
}

var fields_execute_scheduled_query = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "InvocationTime", Flag: "invocation-time", Type: "*time.Time", Required: true},
	{Name: "QueryInsights", Flag: "query-insights", Type: "*types.ScheduledQueryInsights", Required: false},
	{Name: "ScheduledQueryArn", Flag: "scheduled-query-arn", Type: "*string", Required: true},
}

var fields_list_scheduled_queries = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_prepare_query = []leanruntime.Field{
	{Name: "QueryString", Flag: "query-string", Type: "*string", Required: true},
	{Name: "ValidateOnly", Flag: "validate-only", Type: "*bool", Required: false},
}

var fields_query = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "MaxRows", Flag: "max-rows", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "QueryInsights", Flag: "query-insights", Type: "*types.QueryInsights", Required: false},
	{Name: "QueryString", Flag: "query-string", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_account_settings = []leanruntime.Field{
	{Name: "MaxQueryTCU", Flag: "max-query-tcu", Type: "*int32", Required: false},
	{Name: "QueryCompute", Flag: "query-compute", Type: "*types.QueryComputeRequest", Required: false},
	{Name: "QueryPricingModel", Flag: "query-pricing-model", Type: "types.QueryPricingModel", Required: false},
}

var fields_update_scheduled_query = []leanruntime.Field{
	{Name: "ScheduledQueryArn", Flag: "scheduled-query-arn", Type: "*string", Required: true},
	{Name: "State", Flag: "state", Type: "types.ScheduledQueryState", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"cancel-query": {
			Name:   "cancel-query",
			Fields: fields_cancel_query,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelQueryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_query, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelQuery(ctx, input)
			},
		},
		"create-scheduled-query": {
			Name:   "create-scheduled-query",
			Fields: fields_create_scheduled_query,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateScheduledQueryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_scheduled_query, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateScheduledQuery(ctx, input)
			},
		},
		"delete-scheduled-query": {
			Name:   "delete-scheduled-query",
			Fields: fields_delete_scheduled_query,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteScheduledQueryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_scheduled_query, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteScheduledQuery(ctx, input)
			},
		},
		"describe-account-settings": {
			Name:   "describe-account-settings",
			Fields: fields_describe_account_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAccountSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_account_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAccountSettings(ctx, input)
			},
		},
		"describe-endpoints": {
			Name:   "describe-endpoints",
			Fields: fields_describe_endpoints,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEndpointsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_endpoints, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeEndpoints(ctx, input)
			},
		},
		"describe-scheduled-query": {
			Name:   "describe-scheduled-query",
			Fields: fields_describe_scheduled_query,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeScheduledQueryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_scheduled_query, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeScheduledQuery(ctx, input)
			},
		},
		"execute-scheduled-query": {
			Name:   "execute-scheduled-query",
			Fields: fields_execute_scheduled_query,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ExecuteScheduledQueryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_execute_scheduled_query, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ExecuteScheduledQuery(ctx, input)
			},
		},
		"list-scheduled-queries": {
			Name:   "list-scheduled-queries",
			Fields: fields_list_scheduled_queries,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListScheduledQueriesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_scheduled_queries, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListScheduledQueries(ctx, input)
				}
				var results []*svc.ListScheduledQueriesOutput
				p := svc.NewListScheduledQueriesPaginator(client, input)
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
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_tags_for_resource, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTagsForResource(ctx, input)
				}
				var results []*svc.ListTagsForResourceOutput
				p := svc.NewListTagsForResourcePaginator(client, input)
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
		"prepare-query": {
			Name:   "prepare-query",
			Fields: fields_prepare_query,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PrepareQueryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_prepare_query, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PrepareQuery(ctx, input)
			},
		},
		"query": {
			Name:   "query",
			Fields: fields_query,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.QueryInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_query, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.Query(ctx, input)
				}
				var results []*svc.QueryOutput
				p := svc.NewQueryPaginator(client, input)
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
		"update-account-settings": {
			Name:   "update-account-settings",
			Fields: fields_update_account_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAccountSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_account_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAccountSettings(ctx, input)
			},
		},
		"update-scheduled-query": {
			Name:   "update-scheduled-query",
			Fields: fields_update_scheduled_query,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateScheduledQueryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_scheduled_query, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateScheduledQuery(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("timestreamquery", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
