package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/emrserverless"
)

var fields_cancel_job_run = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "JobRunId", Flag: "job-run-id", Type: "*string", Required: true},
	{Name: "ShutdownGracePeriodInSeconds", Flag: "shutdown-grace-period-in-seconds", Type: "*int32", Required: false},
}

var fields_create_application = []leanruntime.Field{
	{Name: "Architecture", Flag: "architecture", Type: "types.Architecture", Required: false},
	{Name: "AutoStartConfiguration", Flag: "auto-start-configuration", Type: "*types.AutoStartConfig", Required: false},
	{Name: "AutoStopConfiguration", Flag: "auto-stop-configuration", Type: "*types.AutoStopConfig", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "DiskEncryptionConfiguration", Flag: "disk-encryption-configuration", Type: "*types.DiskEncryptionConfiguration", Required: false},
	{Name: "IdentityCenterConfiguration", Flag: "identity-center-configuration", Type: "*types.IdentityCenterConfigurationInput", Required: false},
	{Name: "ImageConfiguration", Flag: "image-configuration", Type: "*types.ImageConfigurationInput", Required: false},
	{Name: "InitialCapacity", Flag: "initial-capacity", Type: "map[string]types.InitialCapacityConfig", Required: false},
	{Name: "InteractiveConfiguration", Flag: "interactive-configuration", Type: "*types.InteractiveConfiguration", Required: false},
	{Name: "JobLevelCostAllocationConfiguration", Flag: "job-level-cost-allocation-configuration", Type: "*types.JobLevelCostAllocationConfiguration", Required: false},
	{Name: "MaximumCapacity", Flag: "maximum-capacity", Type: "*types.MaximumAllowedResources", Required: false},
	{Name: "MonitoringConfiguration", Flag: "monitoring-configuration", Type: "*types.MonitoringConfiguration", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "NetworkConfiguration", Flag: "network-configuration", Type: "*types.NetworkConfiguration", Required: false},
	{Name: "ReleaseLabel", Flag: "release-label", Type: "*string", Required: true},
	{Name: "RuntimeConfiguration", Flag: "runtime-configuration", Type: "[]types.Configuration", Required: false},
	{Name: "SchedulerConfiguration", Flag: "scheduler-configuration", Type: "*types.SchedulerConfiguration", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Type", Flag: "type", Type: "*string", Required: true},
	{Name: "WorkerTypeSpecifications", Flag: "worker-type-specifications", Type: "map[string]types.WorkerTypeSpecificationInput", Required: false},
}

var fields_delete_application = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
}

var fields_get_application = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
}

var fields_get_dashboard_for_job_run = []leanruntime.Field{
	{Name: "AccessSystemProfileLogs", Flag: "access-system-profile-logs", Type: "*bool", Required: false},
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "Attempt", Flag: "attempt", Type: "*int32", Required: false},
	{Name: "JobRunId", Flag: "job-run-id", Type: "*string", Required: true},
}

var fields_get_job_run = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "Attempt", Flag: "attempt", Type: "*int32", Required: false},
	{Name: "JobRunId", Flag: "job-run-id", Type: "*string", Required: true},
}

var fields_list_applications = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "States", Flag: "states", Type: "[]types.ApplicationState", Required: false},
}

var fields_list_job_run_attempts = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "JobRunId", Flag: "job-run-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_job_runs = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "CreatedAtAfter", Flag: "created-at-after", Type: "*time.Time", Required: false},
	{Name: "CreatedAtBefore", Flag: "created-at-before", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Mode", Flag: "mode", Type: "types.JobRunMode", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "States", Flag: "states", Type: "[]types.JobRunState", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_start_application = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
}

var fields_start_job_run = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "ConfigurationOverrides", Flag: "configuration-overrides", Type: "*types.ConfigurationOverrides", Required: false},
	{Name: "ExecutionIamPolicy", Flag: "execution-iam-policy", Type: "*types.JobRunExecutionIamPolicy", Required: false},
	{Name: "ExecutionRoleArn", Flag: "execution-role-arn", Type: "*string", Required: true},
	{Name: "ExecutionTimeoutMinutes", Flag: "execution-timeout-minutes", Type: "*int64", Required: false},
	{Name: "JobDriver", Flag: "job-driver", Type: "types.JobDriver", Required: false},
	{Name: "Mode", Flag: "mode", Type: "types.JobRunMode", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "RetryPolicy", Flag: "retry-policy", Type: "*types.RetryPolicy", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_stop_application = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
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
	{Name: "Architecture", Flag: "architecture", Type: "types.Architecture", Required: false},
	{Name: "AutoStartConfiguration", Flag: "auto-start-configuration", Type: "*types.AutoStartConfig", Required: false},
	{Name: "AutoStopConfiguration", Flag: "auto-stop-configuration", Type: "*types.AutoStopConfig", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "DiskEncryptionConfiguration", Flag: "disk-encryption-configuration", Type: "*types.DiskEncryptionConfiguration", Required: false},
	{Name: "IdentityCenterConfiguration", Flag: "identity-center-configuration", Type: "*types.IdentityCenterConfigurationInput", Required: false},
	{Name: "ImageConfiguration", Flag: "image-configuration", Type: "*types.ImageConfigurationInput", Required: false},
	{Name: "InitialCapacity", Flag: "initial-capacity", Type: "map[string]types.InitialCapacityConfig", Required: false},
	{Name: "InteractiveConfiguration", Flag: "interactive-configuration", Type: "*types.InteractiveConfiguration", Required: false},
	{Name: "JobLevelCostAllocationConfiguration", Flag: "job-level-cost-allocation-configuration", Type: "*types.JobLevelCostAllocationConfiguration", Required: false},
	{Name: "MaximumCapacity", Flag: "maximum-capacity", Type: "*types.MaximumAllowedResources", Required: false},
	{Name: "MonitoringConfiguration", Flag: "monitoring-configuration", Type: "*types.MonitoringConfiguration", Required: false},
	{Name: "NetworkConfiguration", Flag: "network-configuration", Type: "*types.NetworkConfiguration", Required: false},
	{Name: "ReleaseLabel", Flag: "release-label", Type: "*string", Required: false},
	{Name: "RuntimeConfiguration", Flag: "runtime-configuration", Type: "[]types.Configuration", Required: false},
	{Name: "SchedulerConfiguration", Flag: "scheduler-configuration", Type: "*types.SchedulerConfiguration", Required: false},
	{Name: "WorkerTypeSpecifications", Flag: "worker-type-specifications", Type: "map[string]types.WorkerTypeSpecificationInput", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"cancel-job-run": {
			Name:   "cancel-job-run",
			Fields: fields_cancel_job_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelJobRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_job_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelJobRun(ctx, input)
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
		"get-dashboard-for-job-run": {
			Name:   "get-dashboard-for-job-run",
			Fields: fields_get_dashboard_for_job_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDashboardForJobRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_dashboard_for_job_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDashboardForJobRun(ctx, input)
			},
		},
		"get-job-run": {
			Name:   "get-job-run",
			Fields: fields_get_job_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetJobRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_job_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetJobRun(ctx, input)
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
		"list-job-run-attempts": {
			Name:   "list-job-run-attempts",
			Fields: fields_list_job_run_attempts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListJobRunAttemptsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_job_run_attempts, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListJobRunAttempts(ctx, input)
				}
				var results []*svc.ListJobRunAttemptsOutput
				p := svc.NewListJobRunAttemptsPaginator(client, input)
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
		"list-job-runs": {
			Name:   "list-job-runs",
			Fields: fields_list_job_runs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListJobRunsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_job_runs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListJobRuns(ctx, input)
				}
				var results []*svc.ListJobRunsOutput
				p := svc.NewListJobRunsPaginator(client, input)
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
		"start-job-run": {
			Name:   "start-job-run",
			Fields: fields_start_job_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartJobRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_job_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartJobRun(ctx, input)
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
	if err := leanruntime.Execute("emrserverless", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
