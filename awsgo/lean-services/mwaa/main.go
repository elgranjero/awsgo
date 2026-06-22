package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/mwaa"
)

var fields_create_cli_token = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_create_environment = []leanruntime.Field{
	{Name: "AirflowConfigurationOptions", Flag: "airflow-configuration-options", Type: "map[string]string", Required: false},
	{Name: "AirflowVersion", Flag: "airflow-version", Type: "*string", Required: false},
	{Name: "DagS3Path", Flag: "dag-s3-path", Type: "*string", Required: true},
	{Name: "EndpointManagement", Flag: "endpoint-management", Type: "types.EndpointManagement", Required: false},
	{Name: "EnvironmentClass", Flag: "environment-class", Type: "*string", Required: false},
	{Name: "ExecutionRoleArn", Flag: "execution-role-arn", Type: "*string", Required: true},
	{Name: "KmsKey", Flag: "kms-key", Type: "*string", Required: false},
	{Name: "LoggingConfiguration", Flag: "logging-configuration", Type: "*types.LoggingConfigurationInput", Required: false},
	{Name: "MaxWebservers", Flag: "max-webservers", Type: "*int32", Required: false},
	{Name: "MaxWorkers", Flag: "max-workers", Type: "*int32", Required: false},
	{Name: "MinWebservers", Flag: "min-webservers", Type: "*int32", Required: false},
	{Name: "MinWorkers", Flag: "min-workers", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NetworkConfiguration", Flag: "network-configuration", Type: "*types.NetworkConfiguration", Required: true},
	{Name: "PluginsS3ObjectVersion", Flag: "plugins-s3-object-version", Type: "*string", Required: false},
	{Name: "PluginsS3Path", Flag: "plugins-s3-path", Type: "*string", Required: false},
	{Name: "RequirementsS3ObjectVersion", Flag: "requirements-s3-object-version", Type: "*string", Required: false},
	{Name: "RequirementsS3Path", Flag: "requirements-s3-path", Type: "*string", Required: false},
	{Name: "Schedulers", Flag: "schedulers", Type: "*int32", Required: false},
	{Name: "SourceBucketArn", Flag: "source-bucket-arn", Type: "*string", Required: true},
	{Name: "StartupScriptS3ObjectVersion", Flag: "startup-script-s3-object-version", Type: "*string", Required: false},
	{Name: "StartupScriptS3Path", Flag: "startup-script-s3-path", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "WebserverAccessMode", Flag: "webserver-access-mode", Type: "types.WebserverAccessMode", Required: false},
	{Name: "WeeklyMaintenanceWindowStart", Flag: "weekly-maintenance-window-start", Type: "*string", Required: false},
}

var fields_create_web_login_token = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_environment = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_environment = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_invoke_rest_api = []leanruntime.Field{
	{Name: "Body", Flag: "body", Type: "document.Interface", Required: false},
	{Name: "Method", Flag: "method", Type: "types.RestApiMethod", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Path", Flag: "path", Type: "*string", Required: true},
	{Name: "QueryParameters", Flag: "query-parameters", Type: "document.Interface", Required: false},
}

var fields_list_environments = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_publish_metrics = []leanruntime.Field{
	{Name: "EnvironmentName", Flag: "environment-name", Type: "*string", Required: true},
	{Name: "MetricData", Flag: "metric-data", Type: "[]types.MetricDatum", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_environment = []leanruntime.Field{
	{Name: "AirflowConfigurationOptions", Flag: "airflow-configuration-options", Type: "map[string]string", Required: false},
	{Name: "AirflowVersion", Flag: "airflow-version", Type: "*string", Required: false},
	{Name: "DagS3Path", Flag: "dag-s3-path", Type: "*string", Required: false},
	{Name: "EnvironmentClass", Flag: "environment-class", Type: "*string", Required: false},
	{Name: "ExecutionRoleArn", Flag: "execution-role-arn", Type: "*string", Required: false},
	{Name: "LoggingConfiguration", Flag: "logging-configuration", Type: "*types.LoggingConfigurationInput", Required: false},
	{Name: "MaxWebservers", Flag: "max-webservers", Type: "*int32", Required: false},
	{Name: "MaxWorkers", Flag: "max-workers", Type: "*int32", Required: false},
	{Name: "MinWebservers", Flag: "min-webservers", Type: "*int32", Required: false},
	{Name: "MinWorkers", Flag: "min-workers", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NetworkConfiguration", Flag: "network-configuration", Type: "*types.UpdateNetworkConfigurationInput", Required: false},
	{Name: "PluginsS3ObjectVersion", Flag: "plugins-s3-object-version", Type: "*string", Required: false},
	{Name: "PluginsS3Path", Flag: "plugins-s3-path", Type: "*string", Required: false},
	{Name: "RequirementsS3ObjectVersion", Flag: "requirements-s3-object-version", Type: "*string", Required: false},
	{Name: "RequirementsS3Path", Flag: "requirements-s3-path", Type: "*string", Required: false},
	{Name: "Schedulers", Flag: "schedulers", Type: "*int32", Required: false},
	{Name: "SourceBucketArn", Flag: "source-bucket-arn", Type: "*string", Required: false},
	{Name: "StartupScriptS3ObjectVersion", Flag: "startup-script-s3-object-version", Type: "*string", Required: false},
	{Name: "StartupScriptS3Path", Flag: "startup-script-s3-path", Type: "*string", Required: false},
	{Name: "WebserverAccessMode", Flag: "webserver-access-mode", Type: "types.WebserverAccessMode", Required: false},
	{Name: "WeeklyMaintenanceWindowStart", Flag: "weekly-maintenance-window-start", Type: "*string", Required: false},
	{Name: "WorkerReplacementStrategy", Flag: "worker-replacement-strategy", Type: "types.WorkerReplacementStrategy", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-cli-token": {
			Name:   "create-cli-token",
			Fields: fields_create_cli_token,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCliTokenInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_cli_token, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCliToken(ctx, input)
			},
		},
		"create-environment": {
			Name:   "create-environment",
			Fields: fields_create_environment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateEnvironmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_environment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateEnvironment(ctx, input)
			},
		},
		"create-web-login-token": {
			Name:   "create-web-login-token",
			Fields: fields_create_web_login_token,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateWebLoginTokenInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_web_login_token, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateWebLoginToken(ctx, input)
			},
		},
		"delete-environment": {
			Name:   "delete-environment",
			Fields: fields_delete_environment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEnvironmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_environment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEnvironment(ctx, input)
			},
		},
		"get-environment": {
			Name:   "get-environment",
			Fields: fields_get_environment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEnvironmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_environment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEnvironment(ctx, input)
			},
		},
		"invoke-rest-api": {
			Name:   "invoke-rest-api",
			Fields: fields_invoke_rest_api,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.InvokeRestApiInput{}
				if _, err := leanruntime.ApplyInput(input, fields_invoke_rest_api, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.InvokeRestApi(ctx, input)
			},
		},
		"list-environments": {
			Name:   "list-environments",
			Fields: fields_list_environments,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEnvironmentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_environments, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEnvironments(ctx, input)
				}
				var results []*svc.ListEnvironmentsOutput
				p := svc.NewListEnvironmentsPaginator(client, input)
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
		"publish-metrics": {
			Name:   "publish-metrics",
			Fields: fields_publish_metrics,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PublishMetricsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_publish_metrics, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PublishMetrics(ctx, input)
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
		"update-environment": {
			Name:   "update-environment",
			Fields: fields_update_environment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateEnvironmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_environment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateEnvironment(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("mwaa", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
