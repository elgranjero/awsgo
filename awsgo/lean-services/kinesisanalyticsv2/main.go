package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/kinesisanalyticsv2"
)

var fields_add_application_cloud_watch_logging_option = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: true},
	{Name: "CloudWatchLoggingOption", Flag: "cloud-watch-logging-option", Type: "*types.CloudWatchLoggingOption", Required: true},
	{Name: "ConditionalToken", Flag: "conditional-token", Type: "*string", Required: false},
	{Name: "CurrentApplicationVersionId", Flag: "current-application-version-id", Type: "*int64", Required: false},
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

var fields_add_application_vpc_configuration = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: true},
	{Name: "ConditionalToken", Flag: "conditional-token", Type: "*string", Required: false},
	{Name: "CurrentApplicationVersionId", Flag: "current-application-version-id", Type: "*int64", Required: false},
	{Name: "VpcConfiguration", Flag: "vpc-configuration", Type: "*types.VpcConfiguration", Required: true},
}

var fields_create_application = []leanruntime.Field{
	{Name: "ApplicationConfiguration", Flag: "application-configuration", Type: "*types.ApplicationConfiguration", Required: false},
	{Name: "ApplicationDescription", Flag: "application-description", Type: "*string", Required: false},
	{Name: "ApplicationMode", Flag: "application-mode", Type: "types.ApplicationMode", Required: false},
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: true},
	{Name: "CloudWatchLoggingOptions", Flag: "cloud-watch-logging-options", Type: "[]types.CloudWatchLoggingOption", Required: false},
	{Name: "RuntimeEnvironment", Flag: "runtime-environment", Type: "types.RuntimeEnvironment", Required: true},
	{Name: "ServiceExecutionRole", Flag: "service-execution-role", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_application_presigned_url = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: true},
	{Name: "SessionExpirationDurationInSeconds", Flag: "session-expiration-duration-in-seconds", Type: "*int64", Required: false},
	{Name: "UrlType", Flag: "url-type", Type: "types.UrlType", Required: true},
}

var fields_create_application_snapshot = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: true},
	{Name: "SnapshotName", Flag: "snapshot-name", Type: "*string", Required: true},
}

var fields_delete_application = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: true},
	{Name: "CreateTimestamp", Flag: "create-timestamp", Type: "*time.Time", Required: true},
}

var fields_delete_application_cloud_watch_logging_option = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: true},
	{Name: "CloudWatchLoggingOptionId", Flag: "cloud-watch-logging-option-id", Type: "*string", Required: true},
	{Name: "ConditionalToken", Flag: "conditional-token", Type: "*string", Required: false},
	{Name: "CurrentApplicationVersionId", Flag: "current-application-version-id", Type: "*int64", Required: false},
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

var fields_delete_application_snapshot = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: true},
	{Name: "SnapshotCreationTimestamp", Flag: "snapshot-creation-timestamp", Type: "*time.Time", Required: true},
	{Name: "SnapshotName", Flag: "snapshot-name", Type: "*string", Required: true},
}

var fields_delete_application_vpc_configuration = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: true},
	{Name: "ConditionalToken", Flag: "conditional-token", Type: "*string", Required: false},
	{Name: "CurrentApplicationVersionId", Flag: "current-application-version-id", Type: "*int64", Required: false},
	{Name: "VpcConfigurationId", Flag: "vpc-configuration-id", Type: "*string", Required: true},
}

var fields_describe_application = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: true},
	{Name: "IncludeAdditionalDetails", Flag: "include-additional-details", Type: "*bool", Required: false},
}

var fields_describe_application_operation = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: true},
	{Name: "OperationId", Flag: "operation-id", Type: "*string", Required: true},
}

var fields_describe_application_snapshot = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: true},
	{Name: "SnapshotName", Flag: "snapshot-name", Type: "*string", Required: true},
}

var fields_describe_application_version = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: true},
	{Name: "ApplicationVersionId", Flag: "application-version-id", Type: "*int64", Required: true},
}

var fields_discover_input_schema = []leanruntime.Field{
	{Name: "InputProcessingConfiguration", Flag: "input-processing-configuration", Type: "*types.InputProcessingConfiguration", Required: false},
	{Name: "InputStartingPositionConfiguration", Flag: "input-starting-position-configuration", Type: "*types.InputStartingPositionConfiguration", Required: false},
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: false},
	{Name: "S3Configuration", Flag: "s3-configuration", Type: "*types.S3Configuration", Required: false},
	{Name: "ServiceExecutionRole", Flag: "service-execution-role", Type: "*string", Required: true},
}

var fields_list_application_operations = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: true},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Operation", Flag: "operation", Type: "*string", Required: false},
	{Name: "OperationStatus", Flag: "operation-status", Type: "types.OperationStatus", Required: false},
}

var fields_list_application_snapshots = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: true},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_application_versions = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: true},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_applications = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_rollback_application = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: true},
	{Name: "CurrentApplicationVersionId", Flag: "current-application-version-id", Type: "*int64", Required: true},
}

var fields_start_application = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: true},
	{Name: "RunConfiguration", Flag: "run-configuration", Type: "*types.RunConfiguration", Required: false},
}

var fields_stop_application = []leanruntime.Field{
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: true},
	{Name: "Force", Flag: "force", Type: "*bool", Required: false},
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
	{Name: "ApplicationConfigurationUpdate", Flag: "application-configuration-update", Type: "*types.ApplicationConfigurationUpdate", Required: false},
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: true},
	{Name: "CloudWatchLoggingOptionUpdates", Flag: "cloud-watch-logging-option-updates", Type: "[]types.CloudWatchLoggingOptionUpdate", Required: false},
	{Name: "ConditionalToken", Flag: "conditional-token", Type: "*string", Required: false},
	{Name: "CurrentApplicationVersionId", Flag: "current-application-version-id", Type: "*int64", Required: false},
	{Name: "RunConfigurationUpdate", Flag: "run-configuration-update", Type: "*types.RunConfigurationUpdate", Required: false},
	{Name: "RuntimeEnvironmentUpdate", Flag: "runtime-environment-update", Type: "types.RuntimeEnvironment", Required: false},
	{Name: "ServiceExecutionRoleUpdate", Flag: "service-execution-role-update", Type: "*string", Required: false},
}

var fields_update_application_maintenance_configuration = []leanruntime.Field{
	{Name: "ApplicationMaintenanceConfigurationUpdate", Flag: "application-maintenance-configuration-update", Type: "*types.ApplicationMaintenanceConfigurationUpdate", Required: true},
	{Name: "ApplicationName", Flag: "application-name", Type: "*string", Required: true},
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
		"add-application-vpc-configuration": {
			Name:   "add-application-vpc-configuration",
			Fields: fields_add_application_vpc_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddApplicationVpcConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_application_vpc_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddApplicationVpcConfiguration(ctx, input)
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
		"create-application-presigned-url": {
			Name:   "create-application-presigned-url",
			Fields: fields_create_application_presigned_url,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateApplicationPresignedUrlInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_application_presigned_url, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateApplicationPresignedUrl(ctx, input)
			},
		},
		"create-application-snapshot": {
			Name:   "create-application-snapshot",
			Fields: fields_create_application_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateApplicationSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_application_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateApplicationSnapshot(ctx, input)
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
		"delete-application-snapshot": {
			Name:   "delete-application-snapshot",
			Fields: fields_delete_application_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteApplicationSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_application_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteApplicationSnapshot(ctx, input)
			},
		},
		"delete-application-vpc-configuration": {
			Name:   "delete-application-vpc-configuration",
			Fields: fields_delete_application_vpc_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteApplicationVpcConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_application_vpc_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteApplicationVpcConfiguration(ctx, input)
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
		"describe-application-operation": {
			Name:   "describe-application-operation",
			Fields: fields_describe_application_operation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeApplicationOperationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_application_operation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeApplicationOperation(ctx, input)
			},
		},
		"describe-application-snapshot": {
			Name:   "describe-application-snapshot",
			Fields: fields_describe_application_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeApplicationSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_application_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeApplicationSnapshot(ctx, input)
			},
		},
		"describe-application-version": {
			Name:   "describe-application-version",
			Fields: fields_describe_application_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeApplicationVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_application_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeApplicationVersion(ctx, input)
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
		"list-application-operations": {
			Name:   "list-application-operations",
			Fields: fields_list_application_operations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListApplicationOperationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_application_operations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListApplicationOperations(ctx, input)
				}
				var results []*svc.ListApplicationOperationsOutput
				p := svc.NewListApplicationOperationsPaginator(client, input)
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
		"list-application-snapshots": {
			Name:   "list-application-snapshots",
			Fields: fields_list_application_snapshots,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListApplicationSnapshotsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_application_snapshots, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListApplicationSnapshots(ctx, input)
				}
				var results []*svc.ListApplicationSnapshotsOutput
				p := svc.NewListApplicationSnapshotsPaginator(client, input)
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
		"list-application-versions": {
			Name:   "list-application-versions",
			Fields: fields_list_application_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListApplicationVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_application_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListApplicationVersions(ctx, input)
				}
				var results []*svc.ListApplicationVersionsOutput
				p := svc.NewListApplicationVersionsPaginator(client, input)
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
		"rollback-application": {
			Name:   "rollback-application",
			Fields: fields_rollback_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RollbackApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_rollback_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RollbackApplication(ctx, input)
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
		"update-application-maintenance-configuration": {
			Name:   "update-application-maintenance-configuration",
			Fields: fields_update_application_maintenance_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateApplicationMaintenanceConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_application_maintenance_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateApplicationMaintenanceConfiguration(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("kinesisanalyticsv2", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
