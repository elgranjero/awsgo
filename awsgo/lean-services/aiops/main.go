package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/aiops"
)

var fields_create_investigation_group = []leanruntime.Field{
	{Name: "ChatbotNotificationChannel", Flag: "chatbot-notification-channel", Type: "map[string][]string", Required: false},
	{Name: "CrossAccountConfigurations", Flag: "cross-account-configurations", Type: "[]types.CrossAccountConfiguration", Required: false},
	{Name: "EncryptionConfiguration", Flag: "encryption-configuration", Type: "*types.EncryptionConfiguration", Required: false},
	{Name: "IsCloudTrailEventHistoryEnabled", Flag: "is-cloud-trail-event-history-enabled", Type: "*bool", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RetentionInDays", Flag: "retention-in-days", Type: "*int64", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "TagKeyBoundaries", Flag: "tag-key-boundaries", Type: "[]string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_delete_investigation_group = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_delete_investigation_group_policy = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_investigation_group = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_investigation_group_policy = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_list_investigation_groups = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_put_investigation_group_policy = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "Policy", Flag: "policy", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_investigation_group = []leanruntime.Field{
	{Name: "ChatbotNotificationChannel", Flag: "chatbot-notification-channel", Type: "map[string][]string", Required: false},
	{Name: "CrossAccountConfigurations", Flag: "cross-account-configurations", Type: "[]types.CrossAccountConfiguration", Required: false},
	{Name: "EncryptionConfiguration", Flag: "encryption-configuration", Type: "*types.EncryptionConfiguration", Required: false},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "IsCloudTrailEventHistoryEnabled", Flag: "is-cloud-trail-event-history-enabled", Type: "*bool", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "TagKeyBoundaries", Flag: "tag-key-boundaries", Type: "[]string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-investigation-group": {
			Name:   "create-investigation-group",
			Fields: fields_create_investigation_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateInvestigationGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_investigation_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateInvestigationGroup(ctx, input)
			},
		},
		"delete-investigation-group": {
			Name:   "delete-investigation-group",
			Fields: fields_delete_investigation_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteInvestigationGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_investigation_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteInvestigationGroup(ctx, input)
			},
		},
		"delete-investigation-group-policy": {
			Name:   "delete-investigation-group-policy",
			Fields: fields_delete_investigation_group_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteInvestigationGroupPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_investigation_group_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteInvestigationGroupPolicy(ctx, input)
			},
		},
		"get-investigation-group": {
			Name:   "get-investigation-group",
			Fields: fields_get_investigation_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetInvestigationGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_investigation_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetInvestigationGroup(ctx, input)
			},
		},
		"get-investigation-group-policy": {
			Name:   "get-investigation-group-policy",
			Fields: fields_get_investigation_group_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetInvestigationGroupPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_investigation_group_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetInvestigationGroupPolicy(ctx, input)
			},
		},
		"list-investigation-groups": {
			Name:   "list-investigation-groups",
			Fields: fields_list_investigation_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListInvestigationGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_investigation_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListInvestigationGroups(ctx, input)
				}
				var results []*svc.ListInvestigationGroupsOutput
				p := svc.NewListInvestigationGroupsPaginator(client, input)
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
		"put-investigation-group-policy": {
			Name:   "put-investigation-group-policy",
			Fields: fields_put_investigation_group_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutInvestigationGroupPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_investigation_group_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutInvestigationGroupPolicy(ctx, input)
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
		"update-investigation-group": {
			Name:   "update-investigation-group",
			Fields: fields_update_investigation_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateInvestigationGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_investigation_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateInvestigationGroup(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("aiops", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
