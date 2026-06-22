package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/kinesisanalytics"
)

var fields_add_application_cloud_watch_logging_option = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: true},
	{Name: "CloudWatchLoggingOption", Flag: "cloud-watch-logging-option", Type: "*types.CloudWatchLoggingOption", Required: true},
	{Name: "CurrentApplicationVersionId", Flag: "current-application-version-id", Type: "*int64", Required: true},
}

var fields_add_application_input = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: true},
	{Name: "CurrentApplicationVersionId", Flag: "current-application-version-id", Type: "*int64", Required: true},
	{Name: "Input", Flag: "input", Type: "*types.Input", Required: true},
}

var fields_add_application_input_processing_configuration = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: true},
	{Name: "CurrentApplicationVersionId", Flag: "current-application-version-id", Type: "*int64", Required: true},
	{Name: "InputId", Flag: "input-id", Type: "*string", Required: true},
	{Name: "InputProcessingConfiguration", Flag: "input-processing-configuration", Type: "*types.InputProcessingConfiguration", Required: true},
}

var fields_add_application_output = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: true},
	{Name: "CurrentApplicationVersionId", Flag: "current-application-version-id", Type: "*int64", Required: true},
	{Name: "Output", Flag: "output", Type: "*types.Output", Required: true},
}

var fields_add_application_reference_data_source = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: true},
	{Name: "CurrentApplicationVersionId", Flag: "current-application-version-id", Type: "*int64", Required: true},
	{Name: "ReferenceDataSource", Flag: "reference-data-source", Type: "*types.ReferenceDataSource", Required: true},
}

var fields_create_application = []leanruntime.Field{
	{Name: "ApplicationCode", Flag: "application-code", Type: "*string", Required: false},
	{Name: "ApplicationDescription", Flag: "application-description", Type: "*string", Required: false},
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: true},
	{Name: "CloudWatchLoggingOptions", Flag: "cloud-watch-logging-options", Type: "[]types.CloudWatchLoggingOption", Required: false},
	{Name: "Inputs", Flag: "inputs", Type: "[]types.Input", Required: false},
	{Name: "Outputs", Flag: "outputs", Type: "[]types.Output", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_delete_application = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: true},
	{Name: "CreateTimestamp", Flag: "create-timestamp", Type: "*time.Time", Required: true},
}

var fields_delete_application_cloud_watch_logging_option = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: true},
	{Name: "CloudWatchLoggingOptionId", Flag: "cloud-watch-logging-option-id", Type: "*string", Required: true},
	{Name: "CurrentApplicationVersionId", Flag: "current-application-version-id", Type: "*int64", Required: true},
}

var fields_delete_application_input_processing_configuration = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: true},
	{Name: "CurrentApplicationVersionId", Flag: "current-application-version-id", Type: "*int64", Required: true},
	{Name: "InputId", Flag: "input-id", Type: "*string", Required: true},
}

var fields_delete_application_output = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: true},
	{Name: "CurrentApplicationVersionId", Flag: "current-application-version-id", Type: "*int64", Required: true},
	{Name: "OutputId", Flag: "output-id", Type: "*string", Required: true},
}

var fields_delete_application_reference_data_source = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: true},
	{Name: "CurrentApplicationVersionId", Flag: "current-application-version-id", Type: "*int64", Required: true},
	{Name: "ReferenceId", Flag: "reference-id", Type: "*string", Required: true},
}

var fields_describe_application = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: true},
}

var fields_discover_input_schema = []leanruntime.Field{
	{Name: "InputProcessingConfiguration", Flag: "input-processing-configuration", Type: "*types.InputProcessingConfiguration", Required: false},
	{Name: "InputStartingPositionConfiguration", Flag: "input-starting-position-configuration", Type: "*types.InputStartingPositionConfiguration", Required: false},
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: false},
	{Name: "RoleARN", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "S3Configuration", Flag: "s3-configuration", Type: "*types.S3Configuration", Required: false},
}

var fields_list_applications = []leanruntime.Field{
	{Name: "ExclusiveStartApplicationName", Flag: "exclusive-start-application-name", Type: "*string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_start_application = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: true},
	{Name: "InputConfigurations", Flag: "input-configurations", Type: "[]types.InputConfiguration", Required: true},
}

var fields_stop_application = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_application = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: true},
	{Name: "ApplicationUpdate", Flag: "application-update", Type: "*types.ApplicationUpdate", Required: true},
	{Name: "CurrentApplicationVersionId", Flag: "current-application-version-id", Type: "*int64", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"add-application-cloud-watch-logging-option": {
			Name:   "add-application-cloud-watch-logging-option",
			Fields: fields_add_application_cloud_watch_logging_option,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddApplicationCloudWatchLoggingOptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_application_cloud_watch_logging_option, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddApplicationCloudWatchLoggingOption(ctx, input)
			},
		},
		"add-application-input": {
			Name:   "add-application-input",
			Fields: fields_add_application_input,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddApplicationInputInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_application_input, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddApplicationInput(ctx, input)
			},
		},
		"add-application-input-processing-configuration": {
			Name:   "add-application-input-processing-configuration",
			Fields: fields_add_application_input_processing_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddApplicationInputProcessingConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_application_input_processing_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddApplicationInputProcessingConfiguration(ctx, input)
			},
		},
		"add-application-output": {
			Name:   "add-application-output",
			Fields: fields_add_application_output,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddApplicationOutputInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_application_output, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddApplicationOutput(ctx, input)
			},
		},
		"add-application-reference-data-source": {
			Name:   "add-application-reference-data-source",
			Fields: fields_add_application_reference_data_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddApplicationReferenceDataSourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_application_reference_data_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddApplicationReferenceDataSource(ctx, input)
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
		"delete-application-cloud-watch-logging-option": {
			Name:   "delete-application-cloud-watch-logging-option",
			Fields: fields_delete_application_cloud_watch_logging_option,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteApplicationCloudWatchLoggingOptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_application_cloud_watch_logging_option, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteApplicationCloudWatchLoggingOption(ctx, input)
			},
		},
		"delete-application-input-processing-configuration": {
			Name:   "delete-application-input-processing-configuration",
			Fields: fields_delete_application_input_processing_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteApplicationInputProcessingConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_application_input_processing_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteApplicationInputProcessingConfiguration(ctx, input)
			},
		},
		"delete-application-output": {
			Name:   "delete-application-output",
			Fields: fields_delete_application_output,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteApplicationOutputInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_application_output, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteApplicationOutput(ctx, input)
			},
		},
		"delete-application-reference-data-source": {
			Name:   "delete-application-reference-data-source",
			Fields: fields_delete_application_reference_data_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteApplicationReferenceDataSourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_application_reference_data_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteApplicationReferenceDataSource(ctx, input)
			},
		},
		"describe-application": {
			Name:   "describe-application",
			Fields: fields_describe_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeApplication(ctx, input)
			},
		},
		"discover-input-schema": {
			Name:   "discover-input-schema",
			Fields: fields_discover_input_schema,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DiscoverInputSchemaInput{}
				if _, err := leanruntime.ApplyInput(input, fields_discover_input_schema, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DiscoverInputSchema(ctx, input)
			},
		},
		"list-applications": {
			Name:   "list-applications",
			Fields: fields_list_applications,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListApplicationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_applications, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListApplications(ctx, input)
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
		"start-application": {
			Name:   "start-application",
			Fields: fields_start_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartApplication(ctx, input)
			},
		},
		"stop-application": {
			Name:   "stop-application",
			Fields: fields_stop_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopApplication(ctx, input)
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
	}
	if err := leanruntime.Execute("kinesisanalytics", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
