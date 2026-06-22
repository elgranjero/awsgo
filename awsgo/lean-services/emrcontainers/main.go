package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/emrcontainers"
)

var fields_cancel_job_run = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "VirtualClusterId", Flag: "virtual-cluster-id", Type: "*string", Required: true},
}

var fields_create_job_template = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "JobTemplateData", Flag: "job-template-data", Type: "*types.JobTemplateData", Required: true},
	{Name: "KmsKeyArn", Flag: "kms-key-arn", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_managed_endpoint = []leanruntime.Field{
	{Name: "CertificateArn", Flag: "certificate-arn", Type: "*string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "ConfigurationOverrides", Flag: "configuration-overrides", Type: "*types.ConfigurationOverrides", Required: false},
	{Name: "ExecutionRoleArn", Flag: "execution-role-arn", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ReleaseLabel", Flag: "release-label", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Type", Flag: "type", Type: "*string", Required: true},
	{Name: "VirtualClusterId", Flag: "virtual-cluster-id", Type: "*string", Required: true},
}

var fields_create_security_configuration = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "ContainerProvider", Flag: "container-provider", Type: "*types.ContainerProvider", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "SecurityConfigurationData", Flag: "security-configuration-data", Type: "*types.SecurityConfigurationData", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_virtual_cluster = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "ContainerProvider", Flag: "container-provider", Type: "*types.ContainerProvider", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "SecurityConfigurationId", Flag: "security-configuration-id", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_delete_job_template = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_delete_managed_endpoint = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "VirtualClusterId", Flag: "virtual-cluster-id", Type: "*string", Required: true},
}

var fields_delete_virtual_cluster = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_describe_job_run = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "VirtualClusterId", Flag: "virtual-cluster-id", Type: "*string", Required: true},
}

var fields_describe_job_template = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_describe_managed_endpoint = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "VirtualClusterId", Flag: "virtual-cluster-id", Type: "*string", Required: true},
}

var fields_describe_security_configuration = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_describe_virtual_cluster = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_managed_endpoint_session_credentials = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "CredentialType", Flag: "credential-type", Type: "*string", Required: true},
	{Name: "DurationInSeconds", Flag: "duration-in-seconds", Type: "*int32", Required: false},
	{Name: "EndpointIdentifier", Flag: "endpoint-identifier", Type: "*string", Required: true},
	{Name: "ExecutionRoleArn", Flag: "execution-role-arn", Type: "*string", Required: true},
	{Name: "LogContext", Flag: "log-context", Type: "*string", Required: false},
	{Name: "VirtualClusterIdentifier", Flag: "virtual-cluster-identifier", Type: "*string", Required: true},
}

var fields_list_job_runs = []leanruntime.Field{
	{Name: "CreatedAfter", Flag: "created-after", Type: "*time.Time", Required: false},
	{Name: "CreatedBefore", Flag: "created-before", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "States", Flag: "states", Type: "[]types.JobRunState", Required: false},
	{Name: "VirtualClusterId", Flag: "virtual-cluster-id", Type: "*string", Required: true},
}

var fields_list_job_templates = []leanruntime.Field{
	{Name: "CreatedAfter", Flag: "created-after", Type: "*time.Time", Required: false},
	{Name: "CreatedBefore", Flag: "created-before", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_managed_endpoints = []leanruntime.Field{
	{Name: "CreatedAfter", Flag: "created-after", Type: "*time.Time", Required: false},
	{Name: "CreatedBefore", Flag: "created-before", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "States", Flag: "states", Type: "[]types.EndpointState", Required: false},
	{Name: "Types", Flag: "types", Type: "[]string", Required: false},
	{Name: "VirtualClusterId", Flag: "virtual-cluster-id", Type: "*string", Required: true},
}

var fields_list_security_configurations = []leanruntime.Field{
	{Name: "CreatedAfter", Flag: "created-after", Type: "*time.Time", Required: false},
	{Name: "CreatedBefore", Flag: "created-before", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_virtual_clusters = []leanruntime.Field{
	{Name: "ContainerProviderId", Flag: "container-provider-id", Type: "*string", Required: false},
	{Name: "ContainerProviderType", Flag: "container-provider-type", Type: "types.ContainerProviderType", Required: false},
	{Name: "CreatedAfter", Flag: "created-after", Type: "*time.Time", Required: false},
	{Name: "CreatedBefore", Flag: "created-before", Type: "*time.Time", Required: false},
	{Name: "EksAccessEntryIntegrated", Flag: "eks-access-entry-integrated", Type: "*bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "States", Flag: "states", Type: "[]types.VirtualClusterState", Required: false},
}

var fields_start_job_run = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "ConfigurationOverrides", Flag: "configuration-overrides", Type: "*types.ConfigurationOverrides", Required: false},
	{Name: "ExecutionRoleArn", Flag: "execution-role-arn", Type: "*string", Required: false},
	{Name: "JobDriver", Flag: "job-driver", Type: "*types.JobDriver", Required: false},
	{Name: "JobTemplateId", Flag: "job-template-id", Type: "*string", Required: false},
	{Name: "JobTemplateParameters", Flag: "job-template-parameters", Type: "map[string]string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "ReleaseLabel", Flag: "release-label", Type: "*string", Required: false},
	{Name: "RetryPolicyConfiguration", Flag: "retry-policy-configuration", Type: "*types.RetryPolicyConfiguration", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "VirtualClusterId", Flag: "virtual-cluster-id", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
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
		"create-managed-endpoint": {
			Name:   "create-managed-endpoint",
			Fields: fields_create_managed_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateManagedEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_managed_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateManagedEndpoint(ctx, input)
			},
		},
		"create-security-configuration": {
			Name:   "create-security-configuration",
			Fields: fields_create_security_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSecurityConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_security_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSecurityConfiguration(ctx, input)
			},
		},
		"create-virtual-cluster": {
			Name:   "create-virtual-cluster",
			Fields: fields_create_virtual_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateVirtualClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_virtual_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateVirtualCluster(ctx, input)
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
		"delete-managed-endpoint": {
			Name:   "delete-managed-endpoint",
			Fields: fields_delete_managed_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteManagedEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_managed_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteManagedEndpoint(ctx, input)
			},
		},
		"delete-virtual-cluster": {
			Name:   "delete-virtual-cluster",
			Fields: fields_delete_virtual_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVirtualClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_virtual_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVirtualCluster(ctx, input)
			},
		},
		"describe-job-run": {
			Name:   "describe-job-run",
			Fields: fields_describe_job_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeJobRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_job_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeJobRun(ctx, input)
			},
		},
		"describe-job-template": {
			Name:   "describe-job-template",
			Fields: fields_describe_job_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeJobTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_job_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeJobTemplate(ctx, input)
			},
		},
		"describe-managed-endpoint": {
			Name:   "describe-managed-endpoint",
			Fields: fields_describe_managed_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeManagedEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_managed_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeManagedEndpoint(ctx, input)
			},
		},
		"describe-security-configuration": {
			Name:   "describe-security-configuration",
			Fields: fields_describe_security_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSecurityConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_security_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeSecurityConfiguration(ctx, input)
			},
		},
		"describe-virtual-cluster": {
			Name:   "describe-virtual-cluster",
			Fields: fields_describe_virtual_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeVirtualClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_virtual_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeVirtualCluster(ctx, input)
			},
		},
		"get-managed-endpoint-session-credentials": {
			Name:   "get-managed-endpoint-session-credentials",
			Fields: fields_get_managed_endpoint_session_credentials,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetManagedEndpointSessionCredentialsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_managed_endpoint_session_credentials, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetManagedEndpointSessionCredentials(ctx, input)
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
		"list-managed-endpoints": {
			Name:   "list-managed-endpoints",
			Fields: fields_list_managed_endpoints,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListManagedEndpointsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_managed_endpoints, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListManagedEndpoints(ctx, input)
				}
				var results []*svc.ListManagedEndpointsOutput
				p := svc.NewListManagedEndpointsPaginator(client, input)
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
		"list-security-configurations": {
			Name:   "list-security-configurations",
			Fields: fields_list_security_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSecurityConfigurationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_security_configurations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSecurityConfigurations(ctx, input)
				}
				var results []*svc.ListSecurityConfigurationsOutput
				p := svc.NewListSecurityConfigurationsPaginator(client, input)
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
		"list-virtual-clusters": {
			Name:   "list-virtual-clusters",
			Fields: fields_list_virtual_clusters,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListVirtualClustersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_virtual_clusters, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListVirtualClusters(ctx, input)
				}
				var results []*svc.ListVirtualClustersOutput
				p := svc.NewListVirtualClustersPaginator(client, input)
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
	}
	if err := leanruntime.Execute("emrcontainers", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
