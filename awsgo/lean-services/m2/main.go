package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/m2"
)

var fields_cancel_batch_job_execution = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "AuthSecretsManagerArn", Flag: "auth-secrets-manager-arn", Type: "*string", Required: false},
	{Name: "ExecutionId", Flag: "execution-id", Type: "*string", Required: true},
}

var fields_create_application = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Definition", Flag: "definition", Type: "types.Definition", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EngineType", Flag: "engine-type", Type: "types.EngineType", Required: true},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_data_set_export_task = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ExportConfig", Flag: "export-config", Type: "types.DataSetExportConfig", Required: true},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
}

var fields_create_data_set_import_task = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ImportConfig", Flag: "import-config", Type: "types.DataSetImportConfig", Required: true},
}

var fields_create_deployment = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "ApplicationVersion", Flag: "application-version", Type: "*int32", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: true},
}

var fields_create_environment = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EngineType", Flag: "engine-type", Type: "types.EngineType", Required: true},
	{Name: "EngineVersion", Flag: "engine-version", Type: "*string", Required: false},
	{Name: "HighAvailabilityConfig", Flag: "high-availability-config", Type: "*types.HighAvailabilityConfig", Required: false},
	{Name: "InstanceType", Flag: "instance-type", Type: "*string", Required: true},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NetworkType", Flag: "network-type", Type: "types.NetworkType", Required: false},
	{Name: "PreferredMaintenanceWindow", Flag: "preferred-maintenance-window", Type: "*string", Required: false},
	{Name: "PubliclyAccessible", Flag: "publicly-accessible", Type: "bool", Required: false},
	{Name: "SecurityGroupIds", Flag: "security-group-ids", Type: "[]string", Required: false},
	{Name: "StorageConfigurations", Flag: "storage-configurations", Type: "[]types.StorageConfiguration", Required: false},
	{Name: "SubnetIds", Flag: "subnet-ids", Type: "[]string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_delete_application = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
}

var fields_delete_application_from_environment = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: true},
}

var fields_delete_environment = []leanruntime.Field{
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: true},
}

var fields_get_application = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
}

var fields_get_application_version = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "ApplicationVersion", Flag: "application-version", Type: "*int32", Required: true},
}

var fields_get_batch_job_execution = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "ExecutionId", Flag: "execution-id", Type: "*string", Required: true},
}

var fields_get_data_set_details = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "DataSetName", Flag: "data-set-name", Type: "*string", Required: true},
}

var fields_get_data_set_export_task = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "TaskId", Flag: "task-id", Type: "*string", Required: true},
}

var fields_get_data_set_import_task = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "TaskId", Flag: "task-id", Type: "*string", Required: true},
}

var fields_get_deployment = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "DeploymentId", Flag: "deployment-id", Type: "*string", Required: true},
}

var fields_get_environment = []leanruntime.Field{
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: true},
}

var fields_get_signed_bluinsights_url = []leanruntime.Field{}

var fields_list_application_versions = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_applications = []leanruntime.Field{
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Names", Flag: "names", Type: "[]string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_batch_job_definitions = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Prefix", Flag: "prefix", Type: "*string", Required: false},
}

var fields_list_batch_job_executions = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "ExecutionIds", Flag: "execution-ids", Type: "[]string", Required: false},
	{Name: "JobName", Flag: "job-name", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StartedAfter", Flag: "started-after", Type: "*time.Time", Required: false},
	{Name: "StartedBefore", Flag: "started-before", Type: "*time.Time", Required: false},
	{Name: "Status", Flag: "status", Type: "types.BatchJobExecutionStatus", Required: false},
}

var fields_list_batch_job_restart_points = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "AuthSecretsManagerArn", Flag: "auth-secrets-manager-arn", Type: "*string", Required: false},
	{Name: "ExecutionId", Flag: "execution-id", Type: "*string", Required: true},
}

var fields_list_data_set_export_history = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_data_set_import_history = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_data_sets = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NameFilter", Flag: "name-filter", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Prefix", Flag: "prefix", Type: "*string", Required: false},
}

var fields_list_deployments = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_engine_versions = []leanruntime.Field{
	{Name: "EngineType", Flag: "engine-type", Type: "types.EngineType", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_environments = []leanruntime.Field{
	{Name: "EngineType", Flag: "engine-type", Type: "types.EngineType", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Names", Flag: "names", Type: "[]string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_start_application = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
}

var fields_start_batch_job = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "AuthSecretsManagerArn", Flag: "auth-secrets-manager-arn", Type: "*string", Required: false},
	{Name: "BatchJobIdentifier", Flag: "batch-job-identifier", Type: "types.BatchJobIdentifier", Required: true},
	{Name: "JobParams", Flag: "job-params", Type: "map[string]string", Required: false},
}

var fields_stop_application = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "ForceStop", Flag: "force-stop", Type: "bool", Required: false},
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
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "CurrentApplicationVersion", Flag: "current-application-version", Type: "*int32", Required: true},
	{Name: "Definition", Flag: "definition", Type: "types.Definition", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
}

var fields_update_environment = []leanruntime.Field{
	{Name: "ApplyDuringMaintenanceWindow", Flag: "apply-during-maintenance-window", Type: "bool", Required: false},
	{Name: "DesiredCapacity", Flag: "desired-capacity", Type: "*int32", Required: false},
	{Name: "EngineVersion", Flag: "engine-version", Type: "*string", Required: false},
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: true},
	{Name: "ForceUpdate", Flag: "force-update", Type: "bool", Required: false},
	{Name: "InstanceType", Flag: "instance-type", Type: "*string", Required: false},
	{Name: "PreferredMaintenanceWindow", Flag: "preferred-maintenance-window", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"cancel-batch-job-execution": {
			Name:   "cancel-batch-job-execution",
			Fields: fields_cancel_batch_job_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelBatchJobExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_batch_job_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelBatchJobExecution(ctx, input)
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
		"create-data-set-export-task": {
			Name:   "create-data-set-export-task",
			Fields: fields_create_data_set_export_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDataSetExportTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_data_set_export_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDataSetExportTask(ctx, input)
			},
		},
		"create-data-set-import-task": {
			Name:   "create-data-set-import-task",
			Fields: fields_create_data_set_import_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDataSetImportTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_data_set_import_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDataSetImportTask(ctx, input)
			},
		},
		"create-deployment": {
			Name:   "create-deployment",
			Fields: fields_create_deployment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDeploymentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_deployment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDeployment(ctx, input)
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
		"delete-application-from-environment": {
			Name:   "delete-application-from-environment",
			Fields: fields_delete_application_from_environment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteApplicationFromEnvironmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_application_from_environment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteApplicationFromEnvironment(ctx, input)
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
		"get-application-version": {
			Name:   "get-application-version",
			Fields: fields_get_application_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetApplicationVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_application_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetApplicationVersion(ctx, input)
			},
		},
		"get-batch-job-execution": {
			Name:   "get-batch-job-execution",
			Fields: fields_get_batch_job_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBatchJobExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_batch_job_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBatchJobExecution(ctx, input)
			},
		},
		"get-data-set-details": {
			Name:   "get-data-set-details",
			Fields: fields_get_data_set_details,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDataSetDetailsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_data_set_details, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDataSetDetails(ctx, input)
			},
		},
		"get-data-set-export-task": {
			Name:   "get-data-set-export-task",
			Fields: fields_get_data_set_export_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDataSetExportTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_data_set_export_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDataSetExportTask(ctx, input)
			},
		},
		"get-data-set-import-task": {
			Name:   "get-data-set-import-task",
			Fields: fields_get_data_set_import_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDataSetImportTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_data_set_import_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDataSetImportTask(ctx, input)
			},
		},
		"get-deployment": {
			Name:   "get-deployment",
			Fields: fields_get_deployment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDeploymentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_deployment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDeployment(ctx, input)
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
		"get-signed-bluinsights-url": {
			Name:   "get-signed-bluinsights-url",
			Fields: fields_get_signed_bluinsights_url,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSignedBluinsightsUrlInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_signed_bluinsights_url, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSignedBluinsightsUrl(ctx, input)
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
		"list-batch-job-definitions": {
			Name:   "list-batch-job-definitions",
			Fields: fields_list_batch_job_definitions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBatchJobDefinitionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_batch_job_definitions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListBatchJobDefinitions(ctx, input)
				}
				var results []*svc.ListBatchJobDefinitionsOutput
				p := svc.NewListBatchJobDefinitionsPaginator(client, input)
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
		"list-batch-job-executions": {
			Name:   "list-batch-job-executions",
			Fields: fields_list_batch_job_executions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBatchJobExecutionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_batch_job_executions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListBatchJobExecutions(ctx, input)
				}
				var results []*svc.ListBatchJobExecutionsOutput
				p := svc.NewListBatchJobExecutionsPaginator(client, input)
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
		"list-batch-job-restart-points": {
			Name:   "list-batch-job-restart-points",
			Fields: fields_list_batch_job_restart_points,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBatchJobRestartPointsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_batch_job_restart_points, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListBatchJobRestartPoints(ctx, input)
			},
		},
		"list-data-set-export-history": {
			Name:   "list-data-set-export-history",
			Fields: fields_list_data_set_export_history,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDataSetExportHistoryInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_data_set_export_history, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDataSetExportHistory(ctx, input)
				}
				var results []*svc.ListDataSetExportHistoryOutput
				p := svc.NewListDataSetExportHistoryPaginator(client, input)
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
		"list-data-set-import-history": {
			Name:   "list-data-set-import-history",
			Fields: fields_list_data_set_import_history,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDataSetImportHistoryInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_data_set_import_history, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDataSetImportHistory(ctx, input)
				}
				var results []*svc.ListDataSetImportHistoryOutput
				p := svc.NewListDataSetImportHistoryPaginator(client, input)
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
		"list-data-sets": {
			Name:   "list-data-sets",
			Fields: fields_list_data_sets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDataSetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_data_sets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDataSets(ctx, input)
				}
				var results []*svc.ListDataSetsOutput
				p := svc.NewListDataSetsPaginator(client, input)
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
		"list-deployments": {
			Name:   "list-deployments",
			Fields: fields_list_deployments,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDeploymentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_deployments, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDeployments(ctx, input)
				}
				var results []*svc.ListDeploymentsOutput
				p := svc.NewListDeploymentsPaginator(client, input)
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
		"list-engine-versions": {
			Name:   "list-engine-versions",
			Fields: fields_list_engine_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEngineVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_engine_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEngineVersions(ctx, input)
				}
				var results []*svc.ListEngineVersionsOutput
				p := svc.NewListEngineVersionsPaginator(client, input)
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
		"start-batch-job": {
			Name:   "start-batch-job",
			Fields: fields_start_batch_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartBatchJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_batch_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartBatchJob(ctx, input)
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
	if err := leanruntime.Execute("m2", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
