package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/repostspace"
)

var fields_batch_add_channel_role_to_accessors = []leanruntime.Field{
	{Name: "AccessorIds", Flag: "accessor-ids", Type: "[]string", Required: true},
	{Name: "ChannelId", Flag: "channel-id", Type: "*string", Required: true},
	{Name: "ChannelRole", Flag: "channel-role", Type: "types.ChannelRole", Required: true},
	{Name: "SpaceId", Flag: "space-id", Type: "*string", Required: true},
}

var fields_batch_add_role = []leanruntime.Field{
	{Name: "AccessorIds", Flag: "accessor-ids", Type: "[]string", Required: true},
	{Name: "Role", Flag: "role", Type: "types.Role", Required: true},
	{Name: "SpaceId", Flag: "space-id", Type: "*string", Required: true},
}

var fields_batch_remove_channel_role_from_accessors = []leanruntime.Field{
	{Name: "AccessorIds", Flag: "accessor-ids", Type: "[]string", Required: true},
	{Name: "ChannelId", Flag: "channel-id", Type: "*string", Required: true},
	{Name: "ChannelRole", Flag: "channel-role", Type: "types.ChannelRole", Required: true},
	{Name: "SpaceId", Flag: "space-id", Type: "*string", Required: true},
}

var fields_batch_remove_role = []leanruntime.Field{
	{Name: "AccessorIds", Flag: "accessor-ids", Type: "[]string", Required: true},
	{Name: "Role", Flag: "role", Type: "types.Role", Required: true},
	{Name: "SpaceId", Flag: "space-id", Type: "*string", Required: true},
}

var fields_create_channel = []leanruntime.Field{
	{Name: "ChannelDescription", Flag: "channel-description", Type: "*string", Required: false},
	{Name: "ChannelName", Flag: "channel-name", Type: "*string", Required: true},
	{Name: "SpaceId", Flag: "space-id", Type: "*string", Required: true},
}

var fields_create_space = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "Subdomain", Flag: "subdomain", Type: "*string", Required: true},
	{Name: "SupportedEmailDomains", Flag: "supported-email-domains", Type: "*types.SupportedEmailDomainsParameters", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Tier", Flag: "tier", Type: "types.TierLevel", Required: true},
	{Name: "UserKMSKey", Flag: "user-kms-key", Type: "*string", Required: false},
}

var fields_delete_space = []leanruntime.Field{
	{Name: "SpaceId", Flag: "space-id", Type: "*string", Required: true},
}

var fields_deregister_admin = []leanruntime.Field{
	{Name: "AdminId", Flag: "admin-id", Type: "*string", Required: true},
	{Name: "SpaceId", Flag: "space-id", Type: "*string", Required: true},
}

var fields_get_channel = []leanruntime.Field{
	{Name: "ChannelId", Flag: "channel-id", Type: "*string", Required: true},
	{Name: "SpaceId", Flag: "space-id", Type: "*string", Required: true},
}

var fields_get_space = []leanruntime.Field{
	{Name: "SpaceId", Flag: "space-id", Type: "*string", Required: true},
}

var fields_list_channels = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SpaceId", Flag: "space-id", Type: "*string", Required: true},
}

var fields_list_spaces = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_register_admin = []leanruntime.Field{
	{Name: "AdminId", Flag: "admin-id", Type: "*string", Required: true},
	{Name: "SpaceId", Flag: "space-id", Type: "*string", Required: true},
}

var fields_send_invites = []leanruntime.Field{
	{Name: "AccessorIds", Flag: "accessor-ids", Type: "[]string", Required: true},
	{Name: "Body", Flag: "body", Type: "*string", Required: true},
	{Name: "SpaceId", Flag: "space-id", Type: "*string", Required: true},
	{Name: "Title", Flag: "title", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_channel = []leanruntime.Field{
	{Name: "ChannelDescription", Flag: "channel-description", Type: "*string", Required: false},
	{Name: "ChannelId", Flag: "channel-id", Type: "*string", Required: true},
	{Name: "ChannelName", Flag: "channel-name", Type: "*string", Required: true},
	{Name: "SpaceId", Flag: "space-id", Type: "*string", Required: true},
}

var fields_update_space = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "SpaceId", Flag: "space-id", Type: "*string", Required: true},
	{Name: "SupportedEmailDomains", Flag: "supported-email-domains", Type: "*types.SupportedEmailDomainsParameters", Required: false},
	{Name: "Tier", Flag: "tier", Type: "types.TierLevel", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"batch-add-channel-role-to-accessors": {
			Name:   "batch-add-channel-role-to-accessors",
			Fields: fields_batch_add_channel_role_to_accessors,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchAddChannelRoleToAccessorsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_add_channel_role_to_accessors, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchAddChannelRoleToAccessors(ctx, input)
			},
		},
		"batch-add-role": {
			Name:   "batch-add-role",
			Fields: fields_batch_add_role,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchAddRoleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_add_role, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchAddRole(ctx, input)
			},
		},
		"batch-remove-channel-role-from-accessors": {
			Name:   "batch-remove-channel-role-from-accessors",
			Fields: fields_batch_remove_channel_role_from_accessors,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchRemoveChannelRoleFromAccessorsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_remove_channel_role_from_accessors, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchRemoveChannelRoleFromAccessors(ctx, input)
			},
		},
		"batch-remove-role": {
			Name:   "batch-remove-role",
			Fields: fields_batch_remove_role,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchRemoveRoleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_remove_role, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchRemoveRole(ctx, input)
			},
		},
		"create-channel": {
			Name:   "create-channel",
			Fields: fields_create_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateChannel(ctx, input)
			},
		},
		"create-space": {
			Name:   "create-space",
			Fields: fields_create_space,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSpaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_space, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSpace(ctx, input)
			},
		},
		"delete-space": {
			Name:   "delete-space",
			Fields: fields_delete_space,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSpaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_space, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSpace(ctx, input)
			},
		},
		"deregister-admin": {
			Name:   "deregister-admin",
			Fields: fields_deregister_admin,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeregisterAdminInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deregister_admin, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeregisterAdmin(ctx, input)
			},
		},
		"get-channel": {
			Name:   "get-channel",
			Fields: fields_get_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetChannel(ctx, input)
			},
		},
		"get-space": {
			Name:   "get-space",
			Fields: fields_get_space,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSpaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_space, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSpace(ctx, input)
			},
		},
		"list-channels": {
			Name:   "list-channels",
			Fields: fields_list_channels,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListChannelsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_channels, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListChannels(ctx, input)
				}
				var results []*svc.ListChannelsOutput
				p := svc.NewListChannelsPaginator(client, input)
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
		"list-spaces": {
			Name:   "list-spaces",
			Fields: fields_list_spaces,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSpacesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_spaces, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSpaces(ctx, input)
				}
				var results []*svc.ListSpacesOutput
				p := svc.NewListSpacesPaginator(client, input)
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
		"register-admin": {
			Name:   "register-admin",
			Fields: fields_register_admin,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterAdminInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_admin, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterAdmin(ctx, input)
			},
		},
		"send-invites": {
			Name:   "send-invites",
			Fields: fields_send_invites,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SendInvitesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_send_invites, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SendInvites(ctx, input)
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
		"update-channel": {
			Name:   "update-channel",
			Fields: fields_update_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateChannel(ctx, input)
			},
		},
		"update-space": {
			Name:   "update-space",
			Fields: fields_update_space,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSpaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_space, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSpace(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("repostspace", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
