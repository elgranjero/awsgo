package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/identitystore"
)

var fields_create_group = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "IdentityStoreId", Flag: "identity-store-id", Type: "*string", Required: true},
}

var fields_create_group_membership = []leanruntime.Field{
	{Name: "GroupId", Flag: "group-id", Type: "*string", Required: true},
	{Name: "IdentityStoreId", Flag: "identity-store-id", Type: "*string", Required: true},
	{Name: "MemberId", Flag: "member-id", Type: "types.MemberId", Required: true},
}

var fields_create_user = []leanruntime.Field{
	{Name: "Addresses", Flag: "addresses", Type: "[]types.Address", Required: false},
	{Name: "Birthdate", Flag: "birthdate", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "Emails", Flag: "emails", Type: "[]types.Email", Required: false},
	{Name: "Extensions", Flag: "extensions", Type: "map[string]document.Interface", Required: false},
	{Name: "IdentityStoreId", Flag: "identity-store-id", Type: "*string", Required: true},
	{Name: "Locale", Flag: "locale", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*types.Name", Required: false},
	{Name: "NickName", Flag: "nick-name", Type: "*string", Required: false},
	{Name: "PhoneNumbers", Flag: "phone-numbers", Type: "[]types.PhoneNumber", Required: false},
	{Name: "Photos", Flag: "photos", Type: "[]types.Photo", Required: false},
	{Name: "PreferredLanguage", Flag: "preferred-language", Type: "*string", Required: false},
	{Name: "ProfileUrl", Flag: "profile-url", Type: "*string", Required: false},
	{Name: "Roles", Flag: "roles", Type: "[]types.Role", Required: false},
	{Name: "Timezone", Flag: "timezone", Type: "*string", Required: false},
	{Name: "Title", Flag: "title", Type: "*string", Required: false},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: false},
	{Name: "UserType", Flag: "user-type", Type: "*string", Required: false},
	{Name: "Website", Flag: "website", Type: "*string", Required: false},
}

var fields_delete_group = []leanruntime.Field{
	{Name: "GroupId", Flag: "group-id", Type: "*string", Required: true},
	{Name: "IdentityStoreId", Flag: "identity-store-id", Type: "*string", Required: true},
}

var fields_delete_group_membership = []leanruntime.Field{
	{Name: "IdentityStoreId", Flag: "identity-store-id", Type: "*string", Required: true},
	{Name: "MembershipId", Flag: "membership-id", Type: "*string", Required: true},
}

var fields_delete_user = []leanruntime.Field{
	{Name: "IdentityStoreId", Flag: "identity-store-id", Type: "*string", Required: true},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_describe_group = []leanruntime.Field{
	{Name: "GroupId", Flag: "group-id", Type: "*string", Required: true},
	{Name: "IdentityStoreId", Flag: "identity-store-id", Type: "*string", Required: true},
}

var fields_describe_group_membership = []leanruntime.Field{
	{Name: "IdentityStoreId", Flag: "identity-store-id", Type: "*string", Required: true},
	{Name: "MembershipId", Flag: "membership-id", Type: "*string", Required: true},
}

var fields_describe_user = []leanruntime.Field{
	{Name: "Extensions", Flag: "extensions", Type: "[]string", Required: false},
	{Name: "IdentityStoreId", Flag: "identity-store-id", Type: "*string", Required: true},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_get_group_id = []leanruntime.Field{
	{Name: "AlternateIdentifier", Flag: "alternate-identifier", Type: "types.AlternateIdentifier", Required: true},
	{Name: "IdentityStoreId", Flag: "identity-store-id", Type: "*string", Required: true},
}

var fields_get_group_membership_id = []leanruntime.Field{
	{Name: "GroupId", Flag: "group-id", Type: "*string", Required: true},
	{Name: "IdentityStoreId", Flag: "identity-store-id", Type: "*string", Required: true},
	{Name: "MemberId", Flag: "member-id", Type: "types.MemberId", Required: true},
}

var fields_get_user_id = []leanruntime.Field{
	{Name: "AlternateIdentifier", Flag: "alternate-identifier", Type: "types.AlternateIdentifier", Required: true},
	{Name: "IdentityStoreId", Flag: "identity-store-id", Type: "*string", Required: true},
}

var fields_is_member_in_groups = []leanruntime.Field{
	{Name: "GroupIds", Flag: "group-ids", Type: "[]string", Required: true},
	{Name: "IdentityStoreId", Flag: "identity-store-id", Type: "*string", Required: true},
	{Name: "MemberId", Flag: "member-id", Type: "types.MemberId", Required: true},
}

var fields_list_group_memberships = []leanruntime.Field{
	{Name: "GroupId", Flag: "group-id", Type: "*string", Required: true},
	{Name: "IdentityStoreId", Flag: "identity-store-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_group_memberships_for_member = []leanruntime.Field{
	{Name: "IdentityStoreId", Flag: "identity-store-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MemberId", Flag: "member-id", Type: "types.MemberId", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_groups = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "IdentityStoreId", Flag: "identity-store-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_users = []leanruntime.Field{
	{Name: "Extensions", Flag: "extensions", Type: "[]string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "IdentityStoreId", Flag: "identity-store-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_update_group = []leanruntime.Field{
	{Name: "GroupId", Flag: "group-id", Type: "*string", Required: true},
	{Name: "IdentityStoreId", Flag: "identity-store-id", Type: "*string", Required: true},
	{Name: "Operations", Flag: "operations", Type: "[]types.AttributeOperation", Required: true},
}

var fields_update_user = []leanruntime.Field{
	{Name: "IdentityStoreId", Flag: "identity-store-id", Type: "*string", Required: true},
	{Name: "Operations", Flag: "operations", Type: "[]types.AttributeOperation", Required: true},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
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
		"create-group-membership": {
			Name:   "create-group-membership",
			Fields: fields_create_group_membership,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateGroupMembershipInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_group_membership, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateGroupMembership(ctx, input)
			},
		},
		"create-user": {
			Name:   "create-user",
			Fields: fields_create_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateUser(ctx, input)
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
		"delete-group-membership": {
			Name:   "delete-group-membership",
			Fields: fields_delete_group_membership,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteGroupMembershipInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_group_membership, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteGroupMembership(ctx, input)
			},
		},
		"delete-user": {
			Name:   "delete-user",
			Fields: fields_delete_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteUser(ctx, input)
			},
		},
		"describe-group": {
			Name:   "describe-group",
			Fields: fields_describe_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeGroup(ctx, input)
			},
		},
		"describe-group-membership": {
			Name:   "describe-group-membership",
			Fields: fields_describe_group_membership,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeGroupMembershipInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_group_membership, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeGroupMembership(ctx, input)
			},
		},
		"describe-user": {
			Name:   "describe-user",
			Fields: fields_describe_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeUser(ctx, input)
			},
		},
		"get-group-id": {
			Name:   "get-group-id",
			Fields: fields_get_group_id,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetGroupIdInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_group_id, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetGroupId(ctx, input)
			},
		},
		"get-group-membership-id": {
			Name:   "get-group-membership-id",
			Fields: fields_get_group_membership_id,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetGroupMembershipIdInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_group_membership_id, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetGroupMembershipId(ctx, input)
			},
		},
		"get-user-id": {
			Name:   "get-user-id",
			Fields: fields_get_user_id,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetUserIdInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_user_id, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetUserId(ctx, input)
			},
		},
		"is-member-in-groups": {
			Name:   "is-member-in-groups",
			Fields: fields_is_member_in_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.IsMemberInGroupsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_is_member_in_groups, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.IsMemberInGroups(ctx, input)
			},
		},
		"list-group-memberships": {
			Name:   "list-group-memberships",
			Fields: fields_list_group_memberships,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListGroupMembershipsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_group_memberships, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListGroupMemberships(ctx, input)
				}
				var results []*svc.ListGroupMembershipsOutput
				p := svc.NewListGroupMembershipsPaginator(client, input)
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
		"list-group-memberships-for-member": {
			Name:   "list-group-memberships-for-member",
			Fields: fields_list_group_memberships_for_member,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListGroupMembershipsForMemberInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_group_memberships_for_member, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListGroupMembershipsForMember(ctx, input)
				}
				var results []*svc.ListGroupMembershipsForMemberOutput
				p := svc.NewListGroupMembershipsForMemberPaginator(client, input)
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
		"list-users": {
			Name:   "list-users",
			Fields: fields_list_users,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListUsersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_users, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListUsers(ctx, input)
				}
				var results []*svc.ListUsersOutput
				p := svc.NewListUsersPaginator(client, input)
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
		"update-group": {
			Name:   "update-group",
			Fields: fields_update_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateGroup(ctx, input)
			},
		},
		"update-user": {
			Name:   "update-user",
			Fields: fields_update_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateUser(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("identitystore", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
