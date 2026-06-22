package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/iotdeviceadvisor"
)

var fields_create_suite_definition = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "SuiteDefinitionConfiguration", Flag: "suite-definition-configuration", Type: "*types.SuiteDefinitionConfiguration", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_delete_suite_definition = []leanruntime.Field{
	{Name: "SuiteDefinitionId", Flag: "suite-definition-id", Type: "*string", Required: true},
}

var fields_get_endpoint = []leanruntime.Field{
	{Name: "AuthenticationMethod", Flag: "authentication-method", Type: "types.AuthenticationMethod", Required: false},
	{Name: "CertificateArn", Flag: "certificate-arn", Type: "*string", Required: false},
	{Name: "DeviceRoleArn", Flag: "device-role-arn", Type: "*string", Required: false},
	{Name: "ThingArn", Flag: "thing-arn", Type: "*string", Required: false},
}

var fields_get_suite_definition = []leanruntime.Field{
	{Name: "SuiteDefinitionId", Flag: "suite-definition-id", Type: "*string", Required: true},
	{Name: "SuiteDefinitionVersion", Flag: "suite-definition-version", Type: "*string", Required: false},
}

var fields_get_suite_run = []leanruntime.Field{
	{Name: "SuiteDefinitionId", Flag: "suite-definition-id", Type: "*string", Required: true},
	{Name: "SuiteRunId", Flag: "suite-run-id", Type: "*string", Required: true},
}

var fields_get_suite_run_report = []leanruntime.Field{
	{Name: "SuiteDefinitionId", Flag: "suite-definition-id", Type: "*string", Required: true},
	{Name: "SuiteRunId", Flag: "suite-run-id", Type: "*string", Required: true},
}

var fields_list_suite_definitions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_suite_runs = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SuiteDefinitionId", Flag: "suite-definition-id", Type: "*string", Required: false},
	{Name: "SuiteDefinitionVersion", Flag: "suite-definition-version", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_start_suite_run = []leanruntime.Field{
	{Name: "SuiteDefinitionId", Flag: "suite-definition-id", Type: "*string", Required: true},
	{Name: "SuiteDefinitionVersion", Flag: "suite-definition-version", Type: "*string", Required: false},
	{Name: "SuiteRunConfiguration", Flag: "suite-run-configuration", Type: "*types.SuiteRunConfiguration", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_stop_suite_run = []leanruntime.Field{
	{Name: "SuiteDefinitionId", Flag: "suite-definition-id", Type: "*string", Required: true},
	{Name: "SuiteRunId", Flag: "suite-run-id", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_suite_definition = []leanruntime.Field{
	{Name: "SuiteDefinitionConfiguration", Flag: "suite-definition-configuration", Type: "*types.SuiteDefinitionConfiguration", Required: true},
	{Name: "SuiteDefinitionId", Flag: "suite-definition-id", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-suite-definition": {
			Name:   "create-suite-definition",
			Fields: fields_create_suite_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSuiteDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_suite_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSuiteDefinition(ctx, input)
			},
		},
		"delete-suite-definition": {
			Name:   "delete-suite-definition",
			Fields: fields_delete_suite_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSuiteDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_suite_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSuiteDefinition(ctx, input)
			},
		},
		"get-endpoint": {
			Name:   "get-endpoint",
			Fields: fields_get_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEndpoint(ctx, input)
			},
		},
		"get-suite-definition": {
			Name:   "get-suite-definition",
			Fields: fields_get_suite_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSuiteDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_suite_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSuiteDefinition(ctx, input)
			},
		},
		"get-suite-run": {
			Name:   "get-suite-run",
			Fields: fields_get_suite_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSuiteRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_suite_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSuiteRun(ctx, input)
			},
		},
		"get-suite-run-report": {
			Name:   "get-suite-run-report",
			Fields: fields_get_suite_run_report,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSuiteRunReportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_suite_run_report, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSuiteRunReport(ctx, input)
			},
		},
		"list-suite-definitions": {
			Name:   "list-suite-definitions",
			Fields: fields_list_suite_definitions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSuiteDefinitionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_suite_definitions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSuiteDefinitions(ctx, input)
				}
				var results []*svc.ListSuiteDefinitionsOutput
				p := svc.NewListSuiteDefinitionsPaginator(client, input)
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
		"list-suite-runs": {
			Name:   "list-suite-runs",
			Fields: fields_list_suite_runs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSuiteRunsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_suite_runs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSuiteRuns(ctx, input)
				}
				var results []*svc.ListSuiteRunsOutput
				p := svc.NewListSuiteRunsPaginator(client, input)
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
		"start-suite-run": {
			Name:   "start-suite-run",
			Fields: fields_start_suite_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartSuiteRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_suite_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartSuiteRun(ctx, input)
			},
		},
		"stop-suite-run": {
			Name:   "stop-suite-run",
			Fields: fields_stop_suite_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopSuiteRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_suite_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopSuiteRun(ctx, input)
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
		"update-suite-definition": {
			Name:   "update-suite-definition",
			Fields: fields_update_suite_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSuiteDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_suite_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSuiteDefinition(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("iotdeviceadvisor", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
