package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/dlm"
)

var fields_create_lifecycle_policy = []leanruntime.Field{
	{Name: "CopyTags", Flag: "copy-tags", Type: "*bool", Required: false},
	{Name: "CreateInterval", Flag: "create-interval", Type: "*int32", Required: false},
	{Name: "CrossRegionCopyTargets", Flag: "cross-region-copy-targets", Type: "[]types.CrossRegionCopyTarget", Required: false},
	{Name: "DefaultPolicy", Flag: "default-policy", Type: "types.DefaultPolicyTypeValues", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: true},
	{Name: "Exclusions", Flag: "exclusions", Type: "*types.Exclusions", Required: false},
	{Name: "ExecutionRoleArn", Flag: "execution-role-arn", Type: "*string", Required: true},
	{Name: "ExtendDeletion", Flag: "extend-deletion", Type: "*bool", Required: false},
	{Name: "PolicyDetails", Flag: "policy-details", Type: "*types.PolicyDetails", Required: false},
	{Name: "RetainInterval", Flag: "retain-interval", Type: "*int32", Required: false},
	{Name: "State", Flag: "state", Type: "types.SettablePolicyStateValues", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_delete_lifecycle_policy = []leanruntime.Field{
	{Name: "PolicyId", Flag: "policy-id", Type: "*string", Required: true},
}

var fields_get_lifecycle_policies = []leanruntime.Field{
	{Name: "DefaultPolicyType", Flag: "default-policy-type", Type: "types.DefaultPoliciesTypeValues", Required: false},
	{Name: "PolicyIds", Flag: "policy-ids", Type: "[]string", Required: false},
	{Name: "ResourceTypes", Flag: "resource-types", Type: "[]types.ResourceTypeValues", Required: false},
	{Name: "State", Flag: "state", Type: "types.GettablePolicyStateValues", Required: false},
	{Name: "TagsToAdd", Flag: "tags-to-add", Type: "[]string", Required: false},
	{Name: "TargetTags", Flag: "target-tags", Type: "[]string", Required: false},
}

var fields_get_lifecycle_policy = []leanruntime.Field{
	{Name: "PolicyId", Flag: "policy-id", Type: "*string", Required: true},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_lifecycle_policy = []leanruntime.Field{
	{Name: "CopyTags", Flag: "copy-tags", Type: "*bool", Required: false},
	{Name: "CreateInterval", Flag: "create-interval", Type: "*int32", Required: false},
	{Name: "CrossRegionCopyTargets", Flag: "cross-region-copy-targets", Type: "[]types.CrossRegionCopyTarget", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Exclusions", Flag: "exclusions", Type: "*types.Exclusions", Required: false},
	{Name: "ExecutionRoleArn", Flag: "execution-role-arn", Type: "*string", Required: false},
	{Name: "ExtendDeletion", Flag: "extend-deletion", Type: "*bool", Required: false},
	{Name: "PolicyDetails", Flag: "policy-details", Type: "*types.PolicyDetails", Required: false},
	{Name: "PolicyId", Flag: "policy-id", Type: "*string", Required: true},
	{Name: "RetainInterval", Flag: "retain-interval", Type: "*int32", Required: false},
	{Name: "State", Flag: "state", Type: "types.SettablePolicyStateValues", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-lifecycle-policy": {
			Name:   "create-lifecycle-policy",
			Fields: fields_create_lifecycle_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLifecyclePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_lifecycle_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLifecyclePolicy(ctx, input)
			},
		},
		"delete-lifecycle-policy": {
			Name:   "delete-lifecycle-policy",
			Fields: fields_delete_lifecycle_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteLifecyclePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_lifecycle_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteLifecyclePolicy(ctx, input)
			},
		},
		"get-lifecycle-policies": {
			Name:   "get-lifecycle-policies",
			Fields: fields_get_lifecycle_policies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLifecyclePoliciesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_lifecycle_policies, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetLifecyclePolicies(ctx, input)
			},
		},
		"get-lifecycle-policy": {
			Name:   "get-lifecycle-policy",
			Fields: fields_get_lifecycle_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLifecyclePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_lifecycle_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetLifecyclePolicy(ctx, input)
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
		"update-lifecycle-policy": {
			Name:   "update-lifecycle-policy",
			Fields: fields_update_lifecycle_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateLifecyclePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_lifecycle_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateLifecyclePolicy(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("dlm", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
