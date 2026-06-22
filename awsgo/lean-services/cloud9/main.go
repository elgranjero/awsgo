package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/cloud9"
)

var fields_create_environment_ec2 = []leanruntime.Field{
	{Name: "AutomaticStopTimeMinutes", Flag: "automatic-stop-time-minutes", Type: "*int32", Required: false},
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "ConnectionType", Flag: "connection-type", Type: "types.ConnectionType", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "ImageId", Flag: "image-id", Type: "*string", Required: true},
	{Name: "InstanceType", Flag: "instance-type", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "OwnerArn", Flag: "owner-arn", Type: "*string", Required: false},
	{Name: "SubnetId", Flag: "subnet-id", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_environment_membership = []leanruntime.Field{
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: true},
	{Name: "Permissions", Flag: "permissions", Type: "types.MemberPermissions", Required: true},
	{Name: "UserArn", Flag: "user-arn", Type: "*string", Required: true},
}

var fields_delete_environment = []leanruntime.Field{
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: true},
}

var fields_delete_environment_membership = []leanruntime.Field{
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: true},
	{Name: "UserArn", Flag: "user-arn", Type: "*string", Required: true},
}

var fields_describe_environment_memberships = []leanruntime.Field{
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Permissions", Flag: "permissions", Type: "[]types.Permissions", Required: false},
	{Name: "UserArn", Flag: "user-arn", Type: "*string", Required: false},
}

var fields_describe_environment_status = []leanruntime.Field{
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: true},
}

var fields_describe_environments = []leanruntime.Field{
	{Name: "EnvironmentIds", Flag: "environment-ids", Type: "[]string", Required: true},
}

var fields_list_environments = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_environment = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: true},
	{Name: "ManagedCredentialsAction", Flag: "managed-credentials-action", Type: "types.ManagedCredentialsAction", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_update_environment_membership = []leanruntime.Field{
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: true},
	{Name: "Permissions", Flag: "permissions", Type: "types.MemberPermissions", Required: true},
	{Name: "UserArn", Flag: "user-arn", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-environment-ec2": {
			Name:   "create-environment-ec2",
			Fields: fields_create_environment_ec2,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateEnvironmentEC2Input{}
				if _, err := leanruntime.ApplyInput(input, fields_create_environment_ec2, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateEnvironmentEC2(ctx, input)
			},
		},
		"create-environment-membership": {
			Name:   "create-environment-membership",
			Fields: fields_create_environment_membership,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateEnvironmentMembershipInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_environment_membership, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateEnvironmentMembership(ctx, input)
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
		"delete-environment-membership": {
			Name:   "delete-environment-membership",
			Fields: fields_delete_environment_membership,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEnvironmentMembershipInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_environment_membership, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEnvironmentMembership(ctx, input)
			},
		},
		"describe-environment-memberships": {
			Name:   "describe-environment-memberships",
			Fields: fields_describe_environment_memberships,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEnvironmentMembershipsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_environment_memberships, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeEnvironmentMemberships(ctx, input)
				}
				var results []*svc.DescribeEnvironmentMembershipsOutput
				p := svc.NewDescribeEnvironmentMembershipsPaginator(client, input)
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
		"describe-environment-status": {
			Name:   "describe-environment-status",
			Fields: fields_describe_environment_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEnvironmentStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_environment_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeEnvironmentStatus(ctx, input)
			},
		},
		"describe-environments": {
			Name:   "describe-environments",
			Fields: fields_describe_environments,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEnvironmentsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_environments, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeEnvironments(ctx, input)
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
		"update-environment-membership": {
			Name:   "update-environment-membership",
			Fields: fields_update_environment_membership,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateEnvironmentMembershipInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_environment_membership, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateEnvironmentMembership(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("cloud9", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
