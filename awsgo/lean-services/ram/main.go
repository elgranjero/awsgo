package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/ram"
)

var fields_accept_resource_share_invitation = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ResourceShareInvitationArn", Flag: "resource-share-invitation-arn", Type: "*string", Required: true},
}

var fields_associate_resource_share = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Principals", Flag: "principals", Type: "[]string", Required: false},
	{Name: "ResourceArns", Flag: "resource-arns", Type: "[]string", Required: false},
	{Name: "ResourceShareArn", Flag: "resource-share-arn", Type: "*string", Required: true},
	{Name: "Sources", Flag: "sources", Type: "[]string", Required: false},
}

var fields_associate_resource_share_permission = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "PermissionArn", Flag: "permission-arn", Type: "*string", Required: true},
	{Name: "PermissionVersion", Flag: "permission-version", Type: "*int32", Required: false},
	{Name: "Replace", Flag: "replace", Type: "*bool", Required: false},
	{Name: "ResourceShareArn", Flag: "resource-share-arn", Type: "*string", Required: true},
}

var fields_create_permission = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "PolicyTemplate", Flag: "policy-template", Type: "*string", Required: true},
	{Name: "ResourceType", Flag: "resource-type", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_permission_version = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "PermissionArn", Flag: "permission-arn", Type: "*string", Required: true},
	{Name: "PolicyTemplate", Flag: "policy-template", Type: "*string", Required: true},
}

var fields_create_resource_share = []leanruntime.Field{
	{Name: "AllowExternalPrincipals", Flag: "allow-external-principals", Type: "*bool", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "PermissionArns", Flag: "permission-arns", Type: "[]string", Required: false},
	{Name: "Principals", Flag: "principals", Type: "[]string", Required: false},
	{Name: "ResourceArns", Flag: "resource-arns", Type: "[]string", Required: false},
	{Name: "ResourceShareConfiguration", Flag: "resource-share-configuration", Type: "*types.ResourceShareConfiguration", Required: false},
	{Name: "Sources", Flag: "sources", Type: "[]string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_delete_permission = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "PermissionArn", Flag: "permission-arn", Type: "*string", Required: true},
}

var fields_delete_permission_version = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "PermissionArn", Flag: "permission-arn", Type: "*string", Required: true},
	{Name: "PermissionVersion", Flag: "permission-version", Type: "*int32", Required: true},
}

var fields_delete_resource_share = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ResourceShareArn", Flag: "resource-share-arn", Type: "*string", Required: true},
}

var fields_disassociate_resource_share = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Principals", Flag: "principals", Type: "[]string", Required: false},
	{Name: "ResourceArns", Flag: "resource-arns", Type: "[]string", Required: false},
	{Name: "ResourceShareArn", Flag: "resource-share-arn", Type: "*string", Required: true},
	{Name: "Sources", Flag: "sources", Type: "[]string", Required: false},
}

var fields_disassociate_resource_share_permission = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "PermissionArn", Flag: "permission-arn", Type: "*string", Required: true},
	{Name: "ResourceShareArn", Flag: "resource-share-arn", Type: "*string", Required: true},
}

var fields_enable_sharing_with_aws_organization = []leanruntime.Field{}

var fields_get_permission = []leanruntime.Field{
	{Name: "PermissionArn", Flag: "permission-arn", Type: "*string", Required: true},
	{Name: "PermissionVersion", Flag: "permission-version", Type: "*int32", Required: false},
}

var fields_get_resource_policies = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Principal", Flag: "principal", Type: "*string", Required: false},
	{Name: "ResourceArns", Flag: "resource-arns", Type: "[]string", Required: true},
}

var fields_get_resource_share_associations = []leanruntime.Field{
	{Name: "AssociationStatus", Flag: "association-status", Type: "types.ResourceShareAssociationStatus", Required: false},
	{Name: "AssociationType", Flag: "association-type", Type: "types.ResourceShareAssociationType", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Principal", Flag: "principal", Type: "*string", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: false},
	{Name: "ResourceShareArns", Flag: "resource-share-arns", Type: "[]string", Required: false},
}

var fields_get_resource_share_invitations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceShareArns", Flag: "resource-share-arns", Type: "[]string", Required: false},
	{Name: "ResourceShareInvitationArns", Flag: "resource-share-invitation-arns", Type: "[]string", Required: false},
}

var fields_get_resource_shares = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PermissionArn", Flag: "permission-arn", Type: "*string", Required: false},
	{Name: "PermissionVersion", Flag: "permission-version", Type: "*int32", Required: false},
	{Name: "ResourceOwner", Flag: "resource-owner", Type: "types.ResourceOwner", Required: true},
	{Name: "ResourceShareArns", Flag: "resource-share-arns", Type: "[]string", Required: false},
	{Name: "ResourceShareStatus", Flag: "resource-share-status", Type: "types.ResourceShareStatus", Required: false},
	{Name: "TagFilters", Flag: "tag-filters", Type: "[]types.TagFilter", Required: false},
}

var fields_list_pending_invitation_resources = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceRegionScope", Flag: "resource-region-scope", Type: "types.ResourceRegionScopeFilter", Required: false},
	{Name: "ResourceShareInvitationArn", Flag: "resource-share-invitation-arn", Type: "*string", Required: true},
}

var fields_list_permission_associations = []leanruntime.Field{
	{Name: "AssociationStatus", Flag: "association-status", Type: "types.ResourceShareAssociationStatus", Required: false},
	{Name: "DefaultVersion", Flag: "default-version", Type: "*bool", Required: false},
	{Name: "FeatureSet", Flag: "feature-set", Type: "types.PermissionFeatureSet", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PermissionArn", Flag: "permission-arn", Type: "*string", Required: false},
	{Name: "PermissionVersion", Flag: "permission-version", Type: "*int32", Required: false},
	{Name: "ResourceType", Flag: "resource-type", Type: "*string", Required: false},
}

var fields_list_permission_versions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PermissionArn", Flag: "permission-arn", Type: "*string", Required: true},
}

var fields_list_permissions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PermissionType", Flag: "permission-type", Type: "types.PermissionTypeFilter", Required: false},
	{Name: "ResourceType", Flag: "resource-type", Type: "*string", Required: false},
}

var fields_list_principals = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Principals", Flag: "principals", Type: "[]string", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: false},
	{Name: "ResourceOwner", Flag: "resource-owner", Type: "types.ResourceOwner", Required: true},
	{Name: "ResourceShareArns", Flag: "resource-share-arns", Type: "[]string", Required: false},
	{Name: "ResourceType", Flag: "resource-type", Type: "*string", Required: false},
}

var fields_list_replace_permission_associations_work = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.ReplacePermissionAssociationsWorkStatus", Required: false},
	{Name: "WorkIds", Flag: "work-ids", Type: "[]string", Required: false},
}

var fields_list_resource_share_permissions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceShareArn", Flag: "resource-share-arn", Type: "*string", Required: true},
}

var fields_list_resource_types = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceRegionScope", Flag: "resource-region-scope", Type: "types.ResourceRegionScopeFilter", Required: false},
}

var fields_list_resources = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Principal", Flag: "principal", Type: "*string", Required: false},
	{Name: "ResourceArns", Flag: "resource-arns", Type: "[]string", Required: false},
	{Name: "ResourceOwner", Flag: "resource-owner", Type: "types.ResourceOwner", Required: true},
	{Name: "ResourceRegionScope", Flag: "resource-region-scope", Type: "types.ResourceRegionScopeFilter", Required: false},
	{Name: "ResourceShareArns", Flag: "resource-share-arns", Type: "[]string", Required: false},
	{Name: "ResourceType", Flag: "resource-type", Type: "*string", Required: false},
}

var fields_list_source_associations = []leanruntime.Field{
	{Name: "AssociationStatus", Flag: "association-status", Type: "types.ResourceShareAssociationStatus", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceShareArns", Flag: "resource-share-arns", Type: "[]string", Required: false},
	{Name: "SourceId", Flag: "source-id", Type: "*string", Required: false},
	{Name: "SourceType", Flag: "source-type", Type: "*string", Required: false},
}

var fields_promote_permission_created_from_policy = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "PermissionArn", Flag: "permission-arn", Type: "*string", Required: true},
}

var fields_promote_resource_share_created_from_policy = []leanruntime.Field{
	{Name: "ResourceShareArn", Flag: "resource-share-arn", Type: "*string", Required: true},
}

var fields_reject_resource_share_invitation = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ResourceShareInvitationArn", Flag: "resource-share-invitation-arn", Type: "*string", Required: true},
}

var fields_replace_permission_associations = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "FromPermissionArn", Flag: "from-permission-arn", Type: "*string", Required: true},
	{Name: "FromPermissionVersion", Flag: "from-permission-version", Type: "*int32", Required: false},
	{Name: "ToPermissionArn", Flag: "to-permission-arn", Type: "*string", Required: true},
}

var fields_set_default_permission_version = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "PermissionArn", Flag: "permission-arn", Type: "*string", Required: true},
	{Name: "PermissionVersion", Flag: "permission-version", Type: "*int32", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: false},
	{Name: "ResourceShareArn", Flag: "resource-share-arn", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: false},
	{Name: "ResourceShareArn", Flag: "resource-share-arn", Type: "*string", Required: false},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_resource_share = []leanruntime.Field{
	{Name: "AllowExternalPrincipals", Flag: "allow-external-principals", Type: "*bool", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "ResourceShareArn", Flag: "resource-share-arn", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"accept-resource-share-invitation": {
			Name:   "accept-resource-share-invitation",
			Fields: fields_accept_resource_share_invitation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AcceptResourceShareInvitationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_accept_resource_share_invitation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AcceptResourceShareInvitation(ctx, input)
			},
		},
		"associate-resource-share": {
			Name:   "associate-resource-share",
			Fields: fields_associate_resource_share,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateResourceShareInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_resource_share, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateResourceShare(ctx, input)
			},
		},
		"associate-resource-share-permission": {
			Name:   "associate-resource-share-permission",
			Fields: fields_associate_resource_share_permission,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateResourceSharePermissionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_resource_share_permission, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateResourceSharePermission(ctx, input)
			},
		},
		"create-permission": {
			Name:   "create-permission",
			Fields: fields_create_permission,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePermissionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_permission, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePermission(ctx, input)
			},
		},
		"create-permission-version": {
			Name:   "create-permission-version",
			Fields: fields_create_permission_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePermissionVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_permission_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePermissionVersion(ctx, input)
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
		"delete-permission": {
			Name:   "delete-permission",
			Fields: fields_delete_permission,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePermissionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_permission, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePermission(ctx, input)
			},
		},
		"delete-permission-version": {
			Name:   "delete-permission-version",
			Fields: fields_delete_permission_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePermissionVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_permission_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePermissionVersion(ctx, input)
			},
		},
		"delete-resource-share": {
			Name:   "delete-resource-share",
			Fields: fields_delete_resource_share,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteResourceShareInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_resource_share, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteResourceShare(ctx, input)
			},
		},
		"disassociate-resource-share": {
			Name:   "disassociate-resource-share",
			Fields: fields_disassociate_resource_share,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateResourceShareInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_resource_share, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateResourceShare(ctx, input)
			},
		},
		"disassociate-resource-share-permission": {
			Name:   "disassociate-resource-share-permission",
			Fields: fields_disassociate_resource_share_permission,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateResourceSharePermissionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_resource_share_permission, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateResourceSharePermission(ctx, input)
			},
		},
		"enable-sharing-with-aws-organization": {
			Name:   "enable-sharing-with-aws-organization",
			Fields: fields_enable_sharing_with_aws_organization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableSharingWithAwsOrganizationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_sharing_with_aws_organization, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableSharingWithAwsOrganization(ctx, input)
			},
		},
		"get-permission": {
			Name:   "get-permission",
			Fields: fields_get_permission,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPermissionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_permission, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPermission(ctx, input)
			},
		},
		"get-resource-policies": {
			Name:   "get-resource-policies",
			Fields: fields_get_resource_policies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetResourcePoliciesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_resource_policies, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetResourcePolicies(ctx, input)
				}
				var results []*svc.GetResourcePoliciesOutput
				p := svc.NewGetResourcePoliciesPaginator(client, input)
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
		"get-resource-share-associations": {
			Name:   "get-resource-share-associations",
			Fields: fields_get_resource_share_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetResourceShareAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_resource_share_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetResourceShareAssociations(ctx, input)
				}
				var results []*svc.GetResourceShareAssociationsOutput
				p := svc.NewGetResourceShareAssociationsPaginator(client, input)
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
		"get-resource-share-invitations": {
			Name:   "get-resource-share-invitations",
			Fields: fields_get_resource_share_invitations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetResourceShareInvitationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_resource_share_invitations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetResourceShareInvitations(ctx, input)
				}
				var results []*svc.GetResourceShareInvitationsOutput
				p := svc.NewGetResourceShareInvitationsPaginator(client, input)
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
		"get-resource-shares": {
			Name:   "get-resource-shares",
			Fields: fields_get_resource_shares,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetResourceSharesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_resource_shares, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetResourceShares(ctx, input)
				}
				var results []*svc.GetResourceSharesOutput
				p := svc.NewGetResourceSharesPaginator(client, input)
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
		"list-pending-invitation-resources": {
			Name:   "list-pending-invitation-resources",
			Fields: fields_list_pending_invitation_resources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPendingInvitationResourcesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_pending_invitation_resources, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPendingInvitationResources(ctx, input)
				}
				var results []*svc.ListPendingInvitationResourcesOutput
				p := svc.NewListPendingInvitationResourcesPaginator(client, input)
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
		"list-permission-associations": {
			Name:   "list-permission-associations",
			Fields: fields_list_permission_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPermissionAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_permission_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPermissionAssociations(ctx, input)
				}
				var results []*svc.ListPermissionAssociationsOutput
				p := svc.NewListPermissionAssociationsPaginator(client, input)
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
		"list-permission-versions": {
			Name:   "list-permission-versions",
			Fields: fields_list_permission_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPermissionVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_permission_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPermissionVersions(ctx, input)
				}
				var results []*svc.ListPermissionVersionsOutput
				p := svc.NewListPermissionVersionsPaginator(client, input)
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
		"list-permissions": {
			Name:   "list-permissions",
			Fields: fields_list_permissions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPermissionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_permissions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPermissions(ctx, input)
				}
				var results []*svc.ListPermissionsOutput
				p := svc.NewListPermissionsPaginator(client, input)
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
		"list-principals": {
			Name:   "list-principals",
			Fields: fields_list_principals,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPrincipalsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_principals, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPrincipals(ctx, input)
				}
				var results []*svc.ListPrincipalsOutput
				p := svc.NewListPrincipalsPaginator(client, input)
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
		"list-replace-permission-associations-work": {
			Name:   "list-replace-permission-associations-work",
			Fields: fields_list_replace_permission_associations_work,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListReplacePermissionAssociationsWorkInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_replace_permission_associations_work, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListReplacePermissionAssociationsWork(ctx, input)
				}
				var results []*svc.ListReplacePermissionAssociationsWorkOutput
				p := svc.NewListReplacePermissionAssociationsWorkPaginator(client, input)
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
		"list-resource-share-permissions": {
			Name:   "list-resource-share-permissions",
			Fields: fields_list_resource_share_permissions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListResourceSharePermissionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_resource_share_permissions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListResourceSharePermissions(ctx, input)
				}
				var results []*svc.ListResourceSharePermissionsOutput
				p := svc.NewListResourceSharePermissionsPaginator(client, input)
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
		"list-resource-types": {
			Name:   "list-resource-types",
			Fields: fields_list_resource_types,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListResourceTypesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_resource_types, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListResourceTypes(ctx, input)
				}
				var results []*svc.ListResourceTypesOutput
				p := svc.NewListResourceTypesPaginator(client, input)
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
		"list-resources": {
			Name:   "list-resources",
			Fields: fields_list_resources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListResourcesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_resources, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListResources(ctx, input)
				}
				var results []*svc.ListResourcesOutput
				p := svc.NewListResourcesPaginator(client, input)
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
		"list-source-associations": {
			Name:   "list-source-associations",
			Fields: fields_list_source_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSourceAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_source_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSourceAssociations(ctx, input)
				}
				var results []*svc.ListSourceAssociationsOutput
				p := svc.NewListSourceAssociationsPaginator(client, input)
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
		"promote-permission-created-from-policy": {
			Name:   "promote-permission-created-from-policy",
			Fields: fields_promote_permission_created_from_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PromotePermissionCreatedFromPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_promote_permission_created_from_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PromotePermissionCreatedFromPolicy(ctx, input)
			},
		},
		"promote-resource-share-created-from-policy": {
			Name:   "promote-resource-share-created-from-policy",
			Fields: fields_promote_resource_share_created_from_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PromoteResourceShareCreatedFromPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_promote_resource_share_created_from_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PromoteResourceShareCreatedFromPolicy(ctx, input)
			},
		},
		"reject-resource-share-invitation": {
			Name:   "reject-resource-share-invitation",
			Fields: fields_reject_resource_share_invitation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RejectResourceShareInvitationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reject_resource_share_invitation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RejectResourceShareInvitation(ctx, input)
			},
		},
		"replace-permission-associations": {
			Name:   "replace-permission-associations",
			Fields: fields_replace_permission_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ReplacePermissionAssociationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_replace_permission_associations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ReplacePermissionAssociations(ctx, input)
			},
		},
		"set-default-permission-version": {
			Name:   "set-default-permission-version",
			Fields: fields_set_default_permission_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetDefaultPermissionVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_default_permission_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetDefaultPermissionVersion(ctx, input)
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
		"update-resource-share": {
			Name:   "update-resource-share",
			Fields: fields_update_resource_share,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateResourceShareInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_resource_share, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateResourceShare(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("ram", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
