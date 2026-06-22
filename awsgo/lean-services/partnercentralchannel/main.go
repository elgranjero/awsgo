package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/partnercentralchannel"
)

var fields_accept_channel_handshake = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_cancel_channel_handshake = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_create_channel_handshake = []leanruntime.Field{
	{Name: "AssociatedResourceIdentifier", Flag: "associated-resource-identifier", Type: "*string", Required: true},
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "HandshakeType", Flag: "handshake-type", Type: "types.HandshakeType", Required: true},
	{Name: "Payload", Flag: "payload", Type: "types.ChannelHandshakePayload", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_program_management_account = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: true},
	{Name: "Program", Flag: "program", Type: "types.Program", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_relationship = []leanruntime.Field{
	{Name: "AssociatedAccountId", Flag: "associated-account-id", Type: "*string", Required: true},
	{Name: "AssociationType", Flag: "association-type", Type: "types.AssociationType", Required: true},
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: true},
	{Name: "ProgramManagementAccountIdentifier", Flag: "program-management-account-identifier", Type: "*string", Required: true},
	{Name: "RequestedSupportPlan", Flag: "requested-support-plan", Type: "types.SupportPlan", Required: false},
	{Name: "ResaleAccountModel", Flag: "resale-account-model", Type: "types.ResaleAccountModel", Required: false},
	{Name: "Sector", Flag: "sector", Type: "types.Sector", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_delete_program_management_account = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_delete_relationship = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "ProgramManagementAccountIdentifier", Flag: "program-management-account-identifier", Type: "*string", Required: true},
}

var fields_get_relationship = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "ProgramManagementAccountIdentifier", Flag: "program-management-account-identifier", Type: "*string", Required: true},
}

var fields_list_channel_handshakes = []leanruntime.Field{
	{Name: "AssociatedResourceIdentifiers", Flag: "associated-resource-identifiers", Type: "[]string", Required: false},
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "HandshakeType", Flag: "handshake-type", Type: "types.HandshakeType", Required: true},
	{Name: "HandshakeTypeFilters", Flag: "handshake-type-filters", Type: "types.ListChannelHandshakesTypeFilters", Required: false},
	{Name: "HandshakeTypeSort", Flag: "handshake-type-sort", Type: "types.ListChannelHandshakesTypeSort", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ParticipantType", Flag: "participant-type", Type: "types.ParticipantType", Required: true},
	{Name: "Statuses", Flag: "statuses", Type: "[]types.HandshakeStatus", Required: false},
}

var fields_list_program_management_accounts = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]string", Required: false},
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "DisplayNames", Flag: "display-names", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Programs", Flag: "programs", Type: "[]types.Program", Required: false},
	{Name: "Sort", Flag: "sort", Type: "*types.ListProgramManagementAccountsSortBase", Required: false},
	{Name: "Statuses", Flag: "statuses", Type: "[]types.ProgramManagementAccountStatus", Required: false},
}

var fields_list_relationships = []leanruntime.Field{
	{Name: "AssociatedAccountIds", Flag: "associated-account-ids", Type: "[]string", Required: false},
	{Name: "AssociationTypes", Flag: "association-types", Type: "[]types.AssociationType", Required: false},
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "DisplayNames", Flag: "display-names", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ProgramManagementAccountIdentifiers", Flag: "program-management-account-identifiers", Type: "[]string", Required: false},
	{Name: "Sort", Flag: "sort", Type: "*types.ListRelationshipsSortBase", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_reject_channel_handshake = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_program_management_account = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "Revision", Flag: "revision", Type: "*string", Required: false},
}

var fields_update_relationship = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "ProgramManagementAccountIdentifier", Flag: "program-management-account-identifier", Type: "*string", Required: true},
	{Name: "RequestedSupportPlan", Flag: "requested-support-plan", Type: "types.SupportPlan", Required: false},
	{Name: "Revision", Flag: "revision", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"accept-channel-handshake": {
			Name:   "accept-channel-handshake",
			Fields: fields_accept_channel_handshake,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AcceptChannelHandshakeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_accept_channel_handshake, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AcceptChannelHandshake(ctx, input)
			},
		},
		"cancel-channel-handshake": {
			Name:   "cancel-channel-handshake",
			Fields: fields_cancel_channel_handshake,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelChannelHandshakeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_channel_handshake, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelChannelHandshake(ctx, input)
			},
		},
		"create-channel-handshake": {
			Name:   "create-channel-handshake",
			Fields: fields_create_channel_handshake,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateChannelHandshakeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_channel_handshake, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateChannelHandshake(ctx, input)
			},
		},
		"create-program-management-account": {
			Name:   "create-program-management-account",
			Fields: fields_create_program_management_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateProgramManagementAccountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_program_management_account, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateProgramManagementAccount(ctx, input)
			},
		},
		"create-relationship": {
			Name:   "create-relationship",
			Fields: fields_create_relationship,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRelationshipInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_relationship, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRelationship(ctx, input)
			},
		},
		"delete-program-management-account": {
			Name:   "delete-program-management-account",
			Fields: fields_delete_program_management_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteProgramManagementAccountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_program_management_account, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteProgramManagementAccount(ctx, input)
			},
		},
		"delete-relationship": {
			Name:   "delete-relationship",
			Fields: fields_delete_relationship,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRelationshipInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_relationship, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRelationship(ctx, input)
			},
		},
		"get-relationship": {
			Name:   "get-relationship",
			Fields: fields_get_relationship,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRelationshipInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_relationship, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRelationship(ctx, input)
			},
		},
		"list-channel-handshakes": {
			Name:   "list-channel-handshakes",
			Fields: fields_list_channel_handshakes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListChannelHandshakesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_channel_handshakes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListChannelHandshakes(ctx, input)
				}
				var results []*svc.ListChannelHandshakesOutput
				p := svc.NewListChannelHandshakesPaginator(client, input)
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
		"list-program-management-accounts": {
			Name:   "list-program-management-accounts",
			Fields: fields_list_program_management_accounts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListProgramManagementAccountsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_program_management_accounts, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListProgramManagementAccounts(ctx, input)
				}
				var results []*svc.ListProgramManagementAccountsOutput
				p := svc.NewListProgramManagementAccountsPaginator(client, input)
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
		"list-relationships": {
			Name:   "list-relationships",
			Fields: fields_list_relationships,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRelationshipsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_relationships, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRelationships(ctx, input)
				}
				var results []*svc.ListRelationshipsOutput
				p := svc.NewListRelationshipsPaginator(client, input)
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
		"reject-channel-handshake": {
			Name:   "reject-channel-handshake",
			Fields: fields_reject_channel_handshake,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RejectChannelHandshakeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reject_channel_handshake, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RejectChannelHandshake(ctx, input)
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
		"update-program-management-account": {
			Name:   "update-program-management-account",
			Fields: fields_update_program_management_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateProgramManagementAccountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_program_management_account, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateProgramManagementAccount(ctx, input)
			},
		},
		"update-relationship": {
			Name:   "update-relationship",
			Fields: fields_update_relationship,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRelationshipInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_relationship, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRelationship(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("partnercentralchannel", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
