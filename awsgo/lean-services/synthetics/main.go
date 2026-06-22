package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/synthetics"
)

var fields_associate_resource = []leanruntime.Field{
	{Name: "GroupIdentifier", Flag: "group-identifier", Type: "*string", Required: true},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_create_canary = []leanruntime.Field{
	{Name: "ArtifactConfig", Flag: "artifact-config", Type: "*types.ArtifactConfigInput", Required: false},
	{Name: "ArtifactS3Location", Flag: "artifact-s3-location", Type: "*string", Required: true},
	{Name: "BrowserConfigs", Flag: "browser-configs", Type: "[]types.BrowserConfig", Required: false},
	{Name: "Code", Flag: "code", Type: "*types.CanaryCodeInput", Required: true},
	{Name: "ExecutionRoleArn", Flag: "execution-role-arn", Type: "*string", Required: true},
	{Name: "FailureRetentionPeriodInDays", Flag: "failure-retention-period-in-days", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ProvisionedResourceCleanup", Flag: "provisioned-resource-cleanup", Type: "types.ProvisionedResourceCleanupSetting", Required: false},
	{Name: "ResourcesToReplicateTags", Flag: "resources-to-replicate-tags", Type: "[]types.ResourceToTag", Required: false},
	{Name: "RunConfig", Flag: "run-config", Type: "*types.CanaryRunConfigInput", Required: false},
	{Name: "RuntimeVersion", Flag: "runtime-version", Type: "*string", Required: true},
	{Name: "Schedule", Flag: "schedule", Type: "*types.CanaryScheduleInput", Required: true},
	{Name: "SuccessRetentionPeriodInDays", Flag: "success-retention-period-in-days", Type: "*int32", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "VpcConfig", Flag: "vpc-config", Type: "*types.VpcConfigInput", Required: false},
}

var fields_create_group = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_delete_canary = []leanruntime.Field{
	{Name: "DeleteLambda", Flag: "delete-lambda", Type: "bool", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_group = []leanruntime.Field{
	{Name: "GroupIdentifier", Flag: "group-identifier", Type: "*string", Required: true},
}

var fields_describe_canaries = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Names", Flag: "names", Type: "[]string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_canaries_last_run = []leanruntime.Field{
	{Name: "BrowserType", Flag: "browser-type", Type: "types.BrowserType", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Names", Flag: "names", Type: "[]string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_runtime_versions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_disassociate_resource = []leanruntime.Field{
	{Name: "GroupIdentifier", Flag: "group-identifier", Type: "*string", Required: true},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_get_canary = []leanruntime.Field{
	{Name: "DryRunId", Flag: "dry-run-id", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_canary_runs = []leanruntime.Field{
	{Name: "DryRunId", Flag: "dry-run-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RunType", Flag: "run-type", Type: "types.RunType", Required: false},
}

var fields_get_group = []leanruntime.Field{
	{Name: "GroupIdentifier", Flag: "group-identifier", Type: "*string", Required: true},
}

var fields_list_associated_groups = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_group_resources = []leanruntime.Field{
	{Name: "GroupIdentifier", Flag: "group-identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_groups = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_start_canary = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_start_canary_dry_run = []leanruntime.Field{
	{Name: "ArtifactConfig", Flag: "artifact-config", Type: "*types.ArtifactConfigInput", Required: false},
	{Name: "ArtifactS3Location", Flag: "artifact-s3-location", Type: "*string", Required: false},
	{Name: "BrowserConfigs", Flag: "browser-configs", Type: "[]types.BrowserConfig", Required: false},
	{Name: "Code", Flag: "code", Type: "*types.CanaryCodeInput", Required: false},
	{Name: "ExecutionRoleArn", Flag: "execution-role-arn", Type: "*string", Required: false},
	{Name: "FailureRetentionPeriodInDays", Flag: "failure-retention-period-in-days", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ProvisionedResourceCleanup", Flag: "provisioned-resource-cleanup", Type: "types.ProvisionedResourceCleanupSetting", Required: false},
	{Name: "RunConfig", Flag: "run-config", Type: "*types.CanaryRunConfigInput", Required: false},
	{Name: "RuntimeVersion", Flag: "runtime-version", Type: "*string", Required: false},
	{Name: "SuccessRetentionPeriodInDays", Flag: "success-retention-period-in-days", Type: "*int32", Required: false},
	{Name: "VisualReference", Flag: "visual-reference", Type: "*types.VisualReferenceInput", Required: false},
	{Name: "VisualReferences", Flag: "visual-references", Type: "[]types.VisualReferenceInput", Required: false},
	{Name: "VpcConfig", Flag: "vpc-config", Type: "*types.VpcConfigInput", Required: false},
}

var fields_stop_canary = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_canary = []leanruntime.Field{
	{Name: "ArtifactConfig", Flag: "artifact-config", Type: "*types.ArtifactConfigInput", Required: false},
	{Name: "ArtifactS3Location", Flag: "artifact-s3-location", Type: "*string", Required: false},
	{Name: "BrowserConfigs", Flag: "browser-configs", Type: "[]types.BrowserConfig", Required: false},
	{Name: "Code", Flag: "code", Type: "*types.CanaryCodeInput", Required: false},
	{Name: "DryRunId", Flag: "dry-run-id", Type: "*string", Required: false},
	{Name: "ExecutionRoleArn", Flag: "execution-role-arn", Type: "*string", Required: false},
	{Name: "FailureRetentionPeriodInDays", Flag: "failure-retention-period-in-days", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ProvisionedResourceCleanup", Flag: "provisioned-resource-cleanup", Type: "types.ProvisionedResourceCleanupSetting", Required: false},
	{Name: "RunConfig", Flag: "run-config", Type: "*types.CanaryRunConfigInput", Required: false},
	{Name: "RuntimeVersion", Flag: "runtime-version", Type: "*string", Required: false},
	{Name: "Schedule", Flag: "schedule", Type: "*types.CanaryScheduleInput", Required: false},
	{Name: "SuccessRetentionPeriodInDays", Flag: "success-retention-period-in-days", Type: "*int32", Required: false},
	{Name: "VisualReference", Flag: "visual-reference", Type: "*types.VisualReferenceInput", Required: false},
	{Name: "VisualReferences", Flag: "visual-references", Type: "[]types.VisualReferenceInput", Required: false},
	{Name: "VpcConfig", Flag: "vpc-config", Type: "*types.VpcConfigInput", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"associate-resource": {
			Name:   "associate-resource",
			Fields: fields_associate_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateResource(ctx, input)
			},
		},
		"create-canary": {
			Name:   "create-canary",
			Fields: fields_create_canary,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCanaryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_canary, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCanary(ctx, input)
			},
		},
		"create-group": {
			Name:   "create-group",
			Fields: fields_create_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateGroup(ctx, input)
			},
		},
		"delete-canary": {
			Name:   "delete-canary",
			Fields: fields_delete_canary,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCanaryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_canary, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCanary(ctx, input)
			},
		},
		"delete-group": {
			Name:   "delete-group",
			Fields: fields_delete_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteGroup(ctx, input)
			},
		},
		"describe-canaries": {
			Name:   "describe-canaries",
			Fields: fields_describe_canaries,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCanariesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_canaries, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeCanaries(ctx, input)
				}
				var results []*svc.DescribeCanariesOutput
				p := svc.NewDescribeCanariesPaginator(client, input)
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
		"describe-canaries-last-run": {
			Name:   "describe-canaries-last-run",
			Fields: fields_describe_canaries_last_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCanariesLastRunInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_canaries_last_run, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeCanariesLastRun(ctx, input)
				}
				var results []*svc.DescribeCanariesLastRunOutput
				p := svc.NewDescribeCanariesLastRunPaginator(client, input)
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
		"describe-runtime-versions": {
			Name:   "describe-runtime-versions",
			Fields: fields_describe_runtime_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeRuntimeVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_runtime_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeRuntimeVersions(ctx, input)
				}
				var results []*svc.DescribeRuntimeVersionsOutput
				p := svc.NewDescribeRuntimeVersionsPaginator(client, input)
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
		"disassociate-resource": {
			Name:   "disassociate-resource",
			Fields: fields_disassociate_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateResource(ctx, input)
			},
		},
		"get-canary": {
			Name:   "get-canary",
			Fields: fields_get_canary,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCanaryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_canary, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCanary(ctx, input)
			},
		},
		"get-canary-runs": {
			Name:   "get-canary-runs",
			Fields: fields_get_canary_runs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCanaryRunsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_canary_runs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetCanaryRuns(ctx, input)
				}
				var results []*svc.GetCanaryRunsOutput
				p := svc.NewGetCanaryRunsPaginator(client, input)
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
		"get-group": {
			Name:   "get-group",
			Fields: fields_get_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetGroup(ctx, input)
			},
		},
		"list-associated-groups": {
			Name:   "list-associated-groups",
			Fields: fields_list_associated_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAssociatedGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_associated_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAssociatedGroups(ctx, input)
				}
				var results []*svc.ListAssociatedGroupsOutput
				p := svc.NewListAssociatedGroupsPaginator(client, input)
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
		"list-group-resources": {
			Name:   "list-group-resources",
			Fields: fields_list_group_resources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListGroupResourcesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_group_resources, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListGroupResources(ctx, input)
				}
				var results []*svc.ListGroupResourcesOutput
				p := svc.NewListGroupResourcesPaginator(client, input)
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
		"list-groups": {
			Name:   "list-groups",
			Fields: fields_list_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListGroups(ctx, input)
				}
				var results []*svc.ListGroupsOutput
				p := svc.NewListGroupsPaginator(client, input)
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
		"start-canary": {
			Name:   "start-canary",
			Fields: fields_start_canary,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartCanaryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_canary, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartCanary(ctx, input)
			},
		},
		"start-canary-dry-run": {
			Name:   "start-canary-dry-run",
			Fields: fields_start_canary_dry_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartCanaryDryRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_canary_dry_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartCanaryDryRun(ctx, input)
			},
		},
		"stop-canary": {
			Name:   "stop-canary",
			Fields: fields_stop_canary,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopCanaryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_canary, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopCanary(ctx, input)
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
		"update-canary": {
			Name:   "update-canary",
			Fields: fields_update_canary,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateCanaryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_canary, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateCanary(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("synthetics", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
