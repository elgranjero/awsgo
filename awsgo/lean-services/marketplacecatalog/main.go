package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/marketplacecatalog"
)

var fields_batch_describe_entities = []leanruntime.Field{
	{Name: "EntityRequestList", Flag: "entity-request-list", Type: "[]types.EntityRequest", Required: true},
}

var fields_cancel_change_set = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "ChangeSetId", Flag: "change-set-id", Type: "*string", Required: true},
}

var fields_delete_resource_policy = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_describe_change_set = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "ChangeSetId", Flag: "change-set-id", Type: "*string", Required: true},
}

var fields_describe_entity = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "EntityId", Flag: "entity-id", Type: "*string", Required: true},
}

var fields_get_resource_policy = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_change_sets = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "FilterList", Flag: "filter-list", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Sort", Flag: "sort", Type: "*types.Sort", Required: false},
}

var fields_list_entities = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "EntityType", Flag: "entity-type", Type: "*string", Required: true},
	{Name: "EntityTypeFilters", Flag: "entity-type-filters", Type: "types.EntityTypeFilters", Required: false},
	{Name: "EntityTypeSort", Flag: "entity-type-sort", Type: "types.EntityTypeSort", Required: false},
	{Name: "FilterList", Flag: "filter-list", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OwnershipType", Flag: "ownership-type", Type: "types.OwnershipType", Required: false},
	{Name: "Sort", Flag: "sort", Type: "*types.Sort", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_put_resource_policy = []leanruntime.Field{
	{Name: "Policy", Flag: "policy", Type: "*string", Required: true},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_start_change_set = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "ChangeSet", Flag: "change-set", Type: "[]types.Change", Required: true},
	{Name: "ChangeSetName", Flag: "change-set-name", Type: "*string", Required: false},
	{Name: "ChangeSetTags", Flag: "change-set-tags", Type: "[]types.Tag", Required: false},
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "Intent", Flag: "intent", Type: "types.Intent", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"batch-describe-entities": {
			Name:   "batch-describe-entities",
			Fields: fields_batch_describe_entities,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchDescribeEntitiesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_describe_entities, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchDescribeEntities(ctx, input)
			},
		},
		"cancel-change-set": {
			Name:   "cancel-change-set",
			Fields: fields_cancel_change_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelChangeSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_change_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelChangeSet(ctx, input)
			},
		},
		"delete-resource-policy": {
			Name:   "delete-resource-policy",
			Fields: fields_delete_resource_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteResourcePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_resource_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteResourcePolicy(ctx, input)
			},
		},
		"describe-change-set": {
			Name:   "describe-change-set",
			Fields: fields_describe_change_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeChangeSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_change_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeChangeSet(ctx, input)
			},
		},
		"describe-entity": {
			Name:   "describe-entity",
			Fields: fields_describe_entity,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEntityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_entity, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeEntity(ctx, input)
			},
		},
		"get-resource-policy": {
			Name:   "get-resource-policy",
			Fields: fields_get_resource_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetResourcePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_resource_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetResourcePolicy(ctx, input)
			},
		},
		"list-change-sets": {
			Name:   "list-change-sets",
			Fields: fields_list_change_sets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListChangeSetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_change_sets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListChangeSets(ctx, input)
				}
				var results []*svc.ListChangeSetsOutput
				p := svc.NewListChangeSetsPaginator(client, input)
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
		"list-entities": {
			Name:   "list-entities",
			Fields: fields_list_entities,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEntitiesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_entities, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEntities(ctx, input)
				}
				var results []*svc.ListEntitiesOutput
				p := svc.NewListEntitiesPaginator(client, input)
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
		"put-resource-policy": {
			Name:   "put-resource-policy",
			Fields: fields_put_resource_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutResourcePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_resource_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutResourcePolicy(ctx, input)
			},
		},
		"start-change-set": {
			Name:   "start-change-set",
			Fields: fields_start_change_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartChangeSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_change_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartChangeSet(ctx, input)
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
	if err := leanruntime.Execute("marketplacecatalog", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
