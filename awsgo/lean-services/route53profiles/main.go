package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/route53profiles"
)

var fields_associate_profile = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ProfileId", Flag: "profile-id", Type: "*string", Required: true},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_associate_resource_to_profile = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ProfileId", Flag: "profile-id", Type: "*string", Required: true},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "ResourceProperties", Flag: "resource-properties", Type: "*string", Required: false},
}

var fields_create_profile = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_delete_profile = []leanruntime.Field{
	{Name: "ProfileId", Flag: "profile-id", Type: "*string", Required: true},
}

var fields_disassociate_profile = []leanruntime.Field{
	{Name: "ProfileId", Flag: "profile-id", Type: "*string", Required: true},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
}

var fields_disassociate_resource_from_profile = []leanruntime.Field{
	{Name: "ProfileId", Flag: "profile-id", Type: "*string", Required: true},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_get_profile = []leanruntime.Field{
	{Name: "ProfileId", Flag: "profile-id", Type: "*string", Required: true},
}

var fields_get_profile_association = []leanruntime.Field{
	{Name: "ProfileAssociationId", Flag: "profile-association-id", Type: "*string", Required: true},
}

var fields_get_profile_resource_association = []leanruntime.Field{
	{Name: "ProfileResourceAssociationId", Flag: "profile-resource-association-id", Type: "*string", Required: true},
}

var fields_list_profile_associations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ProfileId", Flag: "profile-id", Type: "*string", Required: false},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: false},
}

var fields_list_profile_resource_associations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ProfileId", Flag: "profile-id", Type: "*string", Required: true},
	{Name: "ResourceType", Flag: "resource-type", Type: "*string", Required: false},
}

var fields_list_profiles = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
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

var fields_update_profile_resource_association = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "ProfileResourceAssociationId", Flag: "profile-resource-association-id", Type: "*string", Required: true},
	{Name: "ResourceProperties", Flag: "resource-properties", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"associate-profile": {
			Name:   "associate-profile",
			Fields: fields_associate_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateProfile(ctx, input)
			},
		},
		"associate-resource-to-profile": {
			Name:   "associate-resource-to-profile",
			Fields: fields_associate_resource_to_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateResourceToProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_resource_to_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateResourceToProfile(ctx, input)
			},
		},
		"create-profile": {
			Name:   "create-profile",
			Fields: fields_create_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateProfile(ctx, input)
			},
		},
		"delete-profile": {
			Name:   "delete-profile",
			Fields: fields_delete_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteProfile(ctx, input)
			},
		},
		"disassociate-profile": {
			Name:   "disassociate-profile",
			Fields: fields_disassociate_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateProfile(ctx, input)
			},
		},
		"disassociate-resource-from-profile": {
			Name:   "disassociate-resource-from-profile",
			Fields: fields_disassociate_resource_from_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateResourceFromProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_resource_from_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateResourceFromProfile(ctx, input)
			},
		},
		"get-profile": {
			Name:   "get-profile",
			Fields: fields_get_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetProfile(ctx, input)
			},
		},
		"get-profile-association": {
			Name:   "get-profile-association",
			Fields: fields_get_profile_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetProfileAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_profile_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetProfileAssociation(ctx, input)
			},
		},
		"get-profile-resource-association": {
			Name:   "get-profile-resource-association",
			Fields: fields_get_profile_resource_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetProfileResourceAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_profile_resource_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetProfileResourceAssociation(ctx, input)
			},
		},
		"list-profile-associations": {
			Name:   "list-profile-associations",
			Fields: fields_list_profile_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListProfileAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_profile_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListProfileAssociations(ctx, input)
				}
				var results []*svc.ListProfileAssociationsOutput
				p := svc.NewListProfileAssociationsPaginator(client, input)
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
		"list-profile-resource-associations": {
			Name:   "list-profile-resource-associations",
			Fields: fields_list_profile_resource_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListProfileResourceAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_profile_resource_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListProfileResourceAssociations(ctx, input)
				}
				var results []*svc.ListProfileResourceAssociationsOutput
				p := svc.NewListProfileResourceAssociationsPaginator(client, input)
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
		"list-profiles": {
			Name:   "list-profiles",
			Fields: fields_list_profiles,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListProfilesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_profiles, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListProfiles(ctx, input)
				}
				var results []*svc.ListProfilesOutput
				p := svc.NewListProfilesPaginator(client, input)
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
		"update-profile-resource-association": {
			Name:   "update-profile-resource-association",
			Fields: fields_update_profile_resource_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateProfileResourceAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_profile_resource_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateProfileResourceAssociation(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("route53profiles", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
