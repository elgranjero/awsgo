package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/mediaconvert"
)

var fields_associate_certificate = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_cancel_job = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_create_job = []leanruntime.Field{
	{Name: "AccelerationSettings", Flag: "acceleration-settings", Type: "*types.AccelerationSettings", Required: false},
	{Name: "BillingTagsSource", Flag: "billing-tags-source", Type: "types.BillingTagsSource", Required: false},
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "HopDestinations", Flag: "hop-destinations", Type: "[]types.HopDestination", Required: false},
	{Name: "JobEngineVersion", Flag: "job-engine-version", Type: "*string", Required: false},
	{Name: "JobTemplate", Flag: "job-template", Type: "*string", Required: false},
	{Name: "Priority", Flag: "priority", Type: "*int32", Required: false},
	{Name: "Queue", Flag: "queue", Type: "*string", Required: false},
	{Name: "Role", Flag: "role", Type: "*string", Required: true},
	{Name: "Settings", Flag: "settings", Type: "*types.JobSettings", Required: true},
	{Name: "SimulateReservedQueue", Flag: "simulate-reserved-queue", Type: "types.SimulateReservedQueue", Required: false},
	{Name: "StatusUpdateInterval", Flag: "status-update-interval", Type: "types.StatusUpdateInterval", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "UserMetadata", Flag: "user-metadata", Type: "map[string]string", Required: false},
}

var fields_create_job_template = []leanruntime.Field{
	{Name: "AccelerationSettings", Flag: "acceleration-settings", Type: "*types.AccelerationSettings", Required: false},
	{Name: "Category", Flag: "category", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "HopDestinations", Flag: "hop-destinations", Type: "[]types.HopDestination", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Priority", Flag: "priority", Type: "*int32", Required: false},
	{Name: "Queue", Flag: "queue", Type: "*string", Required: false},
	{Name: "Settings", Flag: "settings", Type: "*types.JobTemplateSettings", Required: true},
	{Name: "StatusUpdateInterval", Flag: "status-update-interval", Type: "types.StatusUpdateInterval", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_preset = []leanruntime.Field{
	{Name: "Category", Flag: "category", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Settings", Flag: "settings", Type: "*types.PresetSettings", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_queue = []leanruntime.Field{
	{Name: "ConcurrentJobs", Flag: "concurrent-jobs", Type: "*int32", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "PricingPlan", Flag: "pricing-plan", Type: "types.PricingPlan", Required: false},
	{Name: "ReservationPlanSettings", Flag: "reservation-plan-settings", Type: "*types.ReservationPlanSettings", Required: false},
	{Name: "Status", Flag: "status", Type: "types.QueueStatus", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_resource_share = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "SupportCaseId", Flag: "support-case-id", Type: "*string", Required: true},
}

var fields_delete_job_template = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_policy = []leanruntime.Field{}

var fields_delete_preset = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_queue = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_describe_endpoints = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Mode", Flag: "mode", Type: "types.DescribeEndpointsMode", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_disassociate_certificate = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_get_job = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_job_template = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_jobs_query_results = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_policy = []leanruntime.Field{}

var fields_get_preset = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_queue = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_list_job_templates = []leanruntime.Field{
	{Name: "Category", Flag: "category", Type: "*string", Required: false},
	{Name: "ListBy", Flag: "list-by", Type: "types.JobTemplateListBy", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Order", Flag: "order", Type: "types.Order", Required: false},
}

var fields_list_jobs = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Order", Flag: "order", Type: "types.Order", Required: false},
	{Name: "Queue", Flag: "queue", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.JobStatus", Required: false},
}

var fields_list_presets = []leanruntime.Field{
	{Name: "Category", Flag: "category", Type: "*string", Required: false},
	{Name: "ListBy", Flag: "list-by", Type: "types.PresetListBy", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Order", Flag: "order", Type: "types.Order", Required: false},
}

var fields_list_queues = []leanruntime.Field{
	{Name: "ListBy", Flag: "list-by", Type: "types.QueueListBy", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Order", Flag: "order", Type: "types.Order", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_list_versions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_probe = []leanruntime.Field{
	{Name: "InputFiles", Flag: "input-files", Type: "[]types.ProbeInputFile", Required: false},
}

var fields_put_policy = []leanruntime.Field{
	{Name: "Policy", Flag: "policy", Type: "*types.Policy", Required: true},
}

var fields_search_jobs = []leanruntime.Field{
	{Name: "InputFile", Flag: "input-file", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Order", Flag: "order", Type: "types.Order", Required: false},
	{Name: "Queue", Flag: "queue", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.JobStatus", Required: false},
}

var fields_start_jobs_query = []leanruntime.Field{
	{Name: "FilterList", Flag: "filter-list", Type: "[]types.JobsQueryFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Order", Flag: "order", Type: "types.Order", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: false},
}

var fields_update_job_template = []leanruntime.Field{
	{Name: "AccelerationSettings", Flag: "acceleration-settings", Type: "*types.AccelerationSettings", Required: false},
	{Name: "Category", Flag: "category", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "HopDestinations", Flag: "hop-destinations", Type: "[]types.HopDestination", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Priority", Flag: "priority", Type: "*int32", Required: false},
	{Name: "Queue", Flag: "queue", Type: "*string", Required: false},
	{Name: "Settings", Flag: "settings", Type: "*types.JobTemplateSettings", Required: false},
	{Name: "StatusUpdateInterval", Flag: "status-update-interval", Type: "types.StatusUpdateInterval", Required: false},
}

var fields_update_preset = []leanruntime.Field{
	{Name: "Category", Flag: "category", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Settings", Flag: "settings", Type: "*types.PresetSettings", Required: false},
}

var fields_update_queue = []leanruntime.Field{
	{Name: "ConcurrentJobs", Flag: "concurrent-jobs", Type: "*int32", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ReservationPlanSettings", Flag: "reservation-plan-settings", Type: "*types.ReservationPlanSettings", Required: false},
	{Name: "Status", Flag: "status", Type: "types.QueueStatus", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"associate-certificate": {
			Name:   "associate-certificate",
			Fields: fields_associate_certificate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateCertificateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_certificate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateCertificate(ctx, input)
			},
		},
		"cancel-job": {
			Name:   "cancel-job",
			Fields: fields_cancel_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelJob(ctx, input)
			},
		},
		"create-job": {
			Name:   "create-job",
			Fields: fields_create_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateJob(ctx, input)
			},
		},
		"create-job-template": {
			Name:   "create-job-template",
			Fields: fields_create_job_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateJobTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_job_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateJobTemplate(ctx, input)
			},
		},
		"create-preset": {
			Name:   "create-preset",
			Fields: fields_create_preset,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePresetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_preset, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePreset(ctx, input)
			},
		},
		"create-queue": {
			Name:   "create-queue",
			Fields: fields_create_queue,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateQueueInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_queue, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateQueue(ctx, input)
			},
		},
		"create-resource-share": {
			Name:   "create-resource-share",
			Fields: fields_create_resource_share,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateResourceShareInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_resource_share, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateResourceShare(ctx, input)
			},
		},
		"delete-job-template": {
			Name:   "delete-job-template",
			Fields: fields_delete_job_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteJobTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_job_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteJobTemplate(ctx, input)
			},
		},
		"delete-policy": {
			Name:   "delete-policy",
			Fields: fields_delete_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePolicy(ctx, input)
			},
		},
		"delete-preset": {
			Name:   "delete-preset",
			Fields: fields_delete_preset,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePresetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_preset, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePreset(ctx, input)
			},
		},
		"delete-queue": {
			Name:   "delete-queue",
			Fields: fields_delete_queue,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteQueueInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_queue, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteQueue(ctx, input)
			},
		},
		"describe-endpoints": {
			Name:   "describe-endpoints",
			Fields: fields_describe_endpoints,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEndpointsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_endpoints, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeEndpoints(ctx, input)
				}
				var results []*svc.DescribeEndpointsOutput
				p := svc.NewDescribeEndpointsPaginator(client, input)
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
		"disassociate-certificate": {
			Name:   "disassociate-certificate",
			Fields: fields_disassociate_certificate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateCertificateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_certificate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateCertificate(ctx, input)
			},
		},
		"get-job": {
			Name:   "get-job",
			Fields: fields_get_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetJob(ctx, input)
			},
		},
		"get-job-template": {
			Name:   "get-job-template",
			Fields: fields_get_job_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetJobTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_job_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetJobTemplate(ctx, input)
			},
		},
		"get-jobs-query-results": {
			Name:   "get-jobs-query-results",
			Fields: fields_get_jobs_query_results,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetJobsQueryResultsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_jobs_query_results, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetJobsQueryResults(ctx, input)
			},
		},
		"get-policy": {
			Name:   "get-policy",
			Fields: fields_get_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPolicy(ctx, input)
			},
		},
		"get-preset": {
			Name:   "get-preset",
			Fields: fields_get_preset,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPresetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_preset, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPreset(ctx, input)
			},
		},
		"get-queue": {
			Name:   "get-queue",
			Fields: fields_get_queue,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetQueueInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_queue, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetQueue(ctx, input)
			},
		},
		"list-job-templates": {
			Name:   "list-job-templates",
			Fields: fields_list_job_templates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListJobTemplatesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_job_templates, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListJobTemplates(ctx, input)
				}
				var results []*svc.ListJobTemplatesOutput
				p := svc.NewListJobTemplatesPaginator(client, input)
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
		"list-jobs": {
			Name:   "list-jobs",
			Fields: fields_list_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListJobs(ctx, input)
				}
				var results []*svc.ListJobsOutput
				p := svc.NewListJobsPaginator(client, input)
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
		"list-presets": {
			Name:   "list-presets",
			Fields: fields_list_presets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPresetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_presets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPresets(ctx, input)
				}
				var results []*svc.ListPresetsOutput
				p := svc.NewListPresetsPaginator(client, input)
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
		"list-queues": {
			Name:   "list-queues",
			Fields: fields_list_queues,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListQueuesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_queues, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListQueues(ctx, input)
				}
				var results []*svc.ListQueuesOutput
				p := svc.NewListQueuesPaginator(client, input)
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
		"list-versions": {
			Name:   "list-versions",
			Fields: fields_list_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListVersions(ctx, input)
				}
				var results []*svc.ListVersionsOutput
				p := svc.NewListVersionsPaginator(client, input)
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
		"probe": {
			Name:   "probe",
			Fields: fields_probe,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ProbeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_probe, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.Probe(ctx, input)
			},
		},
		"put-policy": {
			Name:   "put-policy",
			Fields: fields_put_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutPolicy(ctx, input)
			},
		},
		"search-jobs": {
			Name:   "search-jobs",
			Fields: fields_search_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchJobs(ctx, input)
				}
				var results []*svc.SearchJobsOutput
				p := svc.NewSearchJobsPaginator(client, input)
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
		"start-jobs-query": {
			Name:   "start-jobs-query",
			Fields: fields_start_jobs_query,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartJobsQueryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_jobs_query, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartJobsQuery(ctx, input)
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
		"update-job-template": {
			Name:   "update-job-template",
			Fields: fields_update_job_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateJobTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_job_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateJobTemplate(ctx, input)
			},
		},
		"update-preset": {
			Name:   "update-preset",
			Fields: fields_update_preset,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePresetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_preset, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePreset(ctx, input)
			},
		},
		"update-queue": {
			Name:   "update-queue",
			Fields: fields_update_queue,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateQueueInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_queue, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateQueue(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("mediaconvert", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
