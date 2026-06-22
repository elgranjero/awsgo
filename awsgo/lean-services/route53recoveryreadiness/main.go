package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/route53recoveryreadiness"
)

var fields_create_cell = []leanruntime.Field{
	{Name: "CellName", Flag: "cell-name", Type: "*string", Required: true},
	{Name: "Cells", Flag: "cells", Type: "[]string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_cross_account_authorization = []leanruntime.Field{
	{Name: "CrossAccountAuthorization", Flag: "cross-account-authorization", Type: "*string", Required: true},
}

var fields_create_readiness_check = []leanruntime.Field{
	{Name: "ReadinessCheckName", Flag: "readiness-check-name", Type: "*string", Required: true},
	{Name: "ResourceSetName", Flag: "resource-set-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_recovery_group = []leanruntime.Field{
	{Name: "Cells", Flag: "cells", Type: "[]string", Required: false},
	{Name: "RecoveryGroupName", Flag: "recovery-group-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_resource_set = []leanruntime.Field{
	{Name: "ResourceSetName", Flag: "resource-set-name", Type: "*string", Required: true},
	{Name: "ResourceSetType", Flag: "resource-set-type", Type: "*string", Required: true},
	{Name: "Resources", Flag: "resources", Type: "[]types.Resource", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_delete_cell = []leanruntime.Field{
	{Name: "CellName", Flag: "cell-name", Type: "*string", Required: true},
}

var fields_delete_cross_account_authorization = []leanruntime.Field{
	{Name: "CrossAccountAuthorization", Flag: "cross-account-authorization", Type: "*string", Required: true},
}

var fields_delete_readiness_check = []leanruntime.Field{
	{Name: "ReadinessCheckName", Flag: "readiness-check-name", Type: "*string", Required: true},
}

var fields_delete_recovery_group = []leanruntime.Field{
	{Name: "RecoveryGroupName", Flag: "recovery-group-name", Type: "*string", Required: true},
}

var fields_delete_resource_set = []leanruntime.Field{
	{Name: "ResourceSetName", Flag: "resource-set-name", Type: "*string", Required: true},
}

var fields_get_architecture_recommendations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RecoveryGroupName", Flag: "recovery-group-name", Type: "*string", Required: true},
}

var fields_get_cell = []leanruntime.Field{
	{Name: "CellName", Flag: "cell-name", Type: "*string", Required: true},
}

var fields_get_cell_readiness_summary = []leanruntime.Field{
	{Name: "CellName", Flag: "cell-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_readiness_check = []leanruntime.Field{
	{Name: "ReadinessCheckName", Flag: "readiness-check-name", Type: "*string", Required: true},
}

var fields_get_readiness_check_resource_status = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ReadinessCheckName", Flag: "readiness-check-name", Type: "*string", Required: true},
	{Name: "ResourceIdentifier", Flag: "resource-identifier", Type: "*string", Required: true},
}

var fields_get_readiness_check_status = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ReadinessCheckName", Flag: "readiness-check-name", Type: "*string", Required: true},
}

var fields_get_recovery_group = []leanruntime.Field{
	{Name: "RecoveryGroupName", Flag: "recovery-group-name", Type: "*string", Required: true},
}

var fields_get_recovery_group_readiness_summary = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RecoveryGroupName", Flag: "recovery-group-name", Type: "*string", Required: true},
}

var fields_get_resource_set = []leanruntime.Field{
	{Name: "ResourceSetName", Flag: "resource-set-name", Type: "*string", Required: true},
}

var fields_list_cells = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_cross_account_authorizations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_readiness_checks = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_recovery_groups = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_resource_sets = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_rules = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceType", Flag: "resource-type", Type: "*string", Required: false},
}

var fields_list_tags_for_resources = []leanruntime.Field{
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

var fields_update_cell = []leanruntime.Field{
	{Name: "CellName", Flag: "cell-name", Type: "*string", Required: true},
	{Name: "Cells", Flag: "cells", Type: "[]string", Required: true},
}

var fields_update_readiness_check = []leanruntime.Field{
	{Name: "ReadinessCheckName", Flag: "readiness-check-name", Type: "*string", Required: true},
	{Name: "ResourceSetName", Flag: "resource-set-name", Type: "*string", Required: true},
}

var fields_update_recovery_group = []leanruntime.Field{
	{Name: "Cells", Flag: "cells", Type: "[]string", Required: true},
	{Name: "RecoveryGroupName", Flag: "recovery-group-name", Type: "*string", Required: true},
}

var fields_update_resource_set = []leanruntime.Field{
	{Name: "ResourceSetName", Flag: "resource-set-name", Type: "*string", Required: true},
	{Name: "ResourceSetType", Flag: "resource-set-type", Type: "*string", Required: true},
	{Name: "Resources", Flag: "resources", Type: "[]types.Resource", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-cell": {
			Name:   "create-cell",
			Fields: fields_create_cell,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCellInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_cell, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCell(ctx, input)
			},
		},
		"create-cross-account-authorization": {
			Name:   "create-cross-account-authorization",
			Fields: fields_create_cross_account_authorization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCrossAccountAuthorizationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_cross_account_authorization, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCrossAccountAuthorization(ctx, input)
			},
		},
		"create-readiness-check": {
			Name:   "create-readiness-check",
			Fields: fields_create_readiness_check,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateReadinessCheckInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_readiness_check, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateReadinessCheck(ctx, input)
			},
		},
		"create-recovery-group": {
			Name:   "create-recovery-group",
			Fields: fields_create_recovery_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRecoveryGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_recovery_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRecoveryGroup(ctx, input)
			},
		},
		"create-resource-set": {
			Name:   "create-resource-set",
			Fields: fields_create_resource_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateResourceSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_resource_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateResourceSet(ctx, input)
			},
		},
		"delete-cell": {
			Name:   "delete-cell",
			Fields: fields_delete_cell,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCellInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_cell, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCell(ctx, input)
			},
		},
		"delete-cross-account-authorization": {
			Name:   "delete-cross-account-authorization",
			Fields: fields_delete_cross_account_authorization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCrossAccountAuthorizationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_cross_account_authorization, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCrossAccountAuthorization(ctx, input)
			},
		},
		"delete-readiness-check": {
			Name:   "delete-readiness-check",
			Fields: fields_delete_readiness_check,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteReadinessCheckInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_readiness_check, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteReadinessCheck(ctx, input)
			},
		},
		"delete-recovery-group": {
			Name:   "delete-recovery-group",
			Fields: fields_delete_recovery_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRecoveryGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_recovery_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRecoveryGroup(ctx, input)
			},
		},
		"delete-resource-set": {
			Name:   "delete-resource-set",
			Fields: fields_delete_resource_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteResourceSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_resource_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteResourceSet(ctx, input)
			},
		},
		"get-architecture-recommendations": {
			Name:   "get-architecture-recommendations",
			Fields: fields_get_architecture_recommendations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetArchitectureRecommendationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_architecture_recommendations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetArchitectureRecommendations(ctx, input)
			},
		},
		"get-cell": {
			Name:   "get-cell",
			Fields: fields_get_cell,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCellInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_cell, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCell(ctx, input)
			},
		},
		"get-cell-readiness-summary": {
			Name:   "get-cell-readiness-summary",
			Fields: fields_get_cell_readiness_summary,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCellReadinessSummaryInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_cell_readiness_summary, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetCellReadinessSummary(ctx, input)
				}
				var results []*svc.GetCellReadinessSummaryOutput
				p := svc.NewGetCellReadinessSummaryPaginator(client, input)
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
		"get-readiness-check": {
			Name:   "get-readiness-check",
			Fields: fields_get_readiness_check,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetReadinessCheckInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_readiness_check, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetReadinessCheck(ctx, input)
			},
		},
		"get-readiness-check-resource-status": {
			Name:   "get-readiness-check-resource-status",
			Fields: fields_get_readiness_check_resource_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetReadinessCheckResourceStatusInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_readiness_check_resource_status, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetReadinessCheckResourceStatus(ctx, input)
				}
				var results []*svc.GetReadinessCheckResourceStatusOutput
				p := svc.NewGetReadinessCheckResourceStatusPaginator(client, input)
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
		"get-readiness-check-status": {
			Name:   "get-readiness-check-status",
			Fields: fields_get_readiness_check_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetReadinessCheckStatusInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_readiness_check_status, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetReadinessCheckStatus(ctx, input)
				}
				var results []*svc.GetReadinessCheckStatusOutput
				p := svc.NewGetReadinessCheckStatusPaginator(client, input)
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
		"get-recovery-group": {
			Name:   "get-recovery-group",
			Fields: fields_get_recovery_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRecoveryGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_recovery_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRecoveryGroup(ctx, input)
			},
		},
		"get-recovery-group-readiness-summary": {
			Name:   "get-recovery-group-readiness-summary",
			Fields: fields_get_recovery_group_readiness_summary,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRecoveryGroupReadinessSummaryInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_recovery_group_readiness_summary, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetRecoveryGroupReadinessSummary(ctx, input)
				}
				var results []*svc.GetRecoveryGroupReadinessSummaryOutput
				p := svc.NewGetRecoveryGroupReadinessSummaryPaginator(client, input)
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
		"get-resource-set": {
			Name:   "get-resource-set",
			Fields: fields_get_resource_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetResourceSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_resource_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetResourceSet(ctx, input)
			},
		},
		"list-cells": {
			Name:   "list-cells",
			Fields: fields_list_cells,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCellsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_cells, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCells(ctx, input)
				}
				var results []*svc.ListCellsOutput
				p := svc.NewListCellsPaginator(client, input)
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
		"list-cross-account-authorizations": {
			Name:   "list-cross-account-authorizations",
			Fields: fields_list_cross_account_authorizations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCrossAccountAuthorizationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_cross_account_authorizations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCrossAccountAuthorizations(ctx, input)
				}
				var results []*svc.ListCrossAccountAuthorizationsOutput
				p := svc.NewListCrossAccountAuthorizationsPaginator(client, input)
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
		"list-readiness-checks": {
			Name:   "list-readiness-checks",
			Fields: fields_list_readiness_checks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListReadinessChecksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_readiness_checks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListReadinessChecks(ctx, input)
				}
				var results []*svc.ListReadinessChecksOutput
				p := svc.NewListReadinessChecksPaginator(client, input)
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
		"list-recovery-groups": {
			Name:   "list-recovery-groups",
			Fields: fields_list_recovery_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRecoveryGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_recovery_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRecoveryGroups(ctx, input)
				}
				var results []*svc.ListRecoveryGroupsOutput
				p := svc.NewListRecoveryGroupsPaginator(client, input)
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
		"list-resource-sets": {
			Name:   "list-resource-sets",
			Fields: fields_list_resource_sets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListResourceSetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_resource_sets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListResourceSets(ctx, input)
				}
				var results []*svc.ListResourceSetsOutput
				p := svc.NewListResourceSetsPaginator(client, input)
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
		"list-rules": {
			Name:   "list-rules",
			Fields: fields_list_rules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRulesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_rules, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRules(ctx, input)
				}
				var results []*svc.ListRulesOutput
				p := svc.NewListRulesPaginator(client, input)
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
		"list-tags-for-resources": {
			Name:   "list-tags-for-resources",
			Fields: fields_list_tags_for_resources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTagsForResourcesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_tags_for_resources, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListTagsForResources(ctx, input)
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
		"update-cell": {
			Name:   "update-cell",
			Fields: fields_update_cell,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateCellInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_cell, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateCell(ctx, input)
			},
		},
		"update-readiness-check": {
			Name:   "update-readiness-check",
			Fields: fields_update_readiness_check,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateReadinessCheckInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_readiness_check, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateReadinessCheck(ctx, input)
			},
		},
		"update-recovery-group": {
			Name:   "update-recovery-group",
			Fields: fields_update_recovery_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRecoveryGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_recovery_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRecoveryGroup(ctx, input)
			},
		},
		"update-resource-set": {
			Name:   "update-resource-set",
			Fields: fields_update_resource_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateResourceSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_resource_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateResourceSet(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("route53recoveryreadiness", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
