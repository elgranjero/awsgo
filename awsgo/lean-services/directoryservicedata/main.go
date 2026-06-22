package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/directoryservicedata"
)

var fields_add_group_member = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
	{Name: "GroupName", Flag: "group-name", Type: "*string", Required: true},
	{Name: "MemberName", Flag: "member-name", Type: "*string", Required: true},
	{Name: "MemberRealm", Flag: "member-realm", Type: "*string", Required: false},
}

var fields_create_group = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
	{Name: "GroupScope", Flag: "group-scope", Type: "types.GroupScope", Required: false},
	{Name: "GroupType", Flag: "group-type", Type: "types.GroupType", Required: false},
	{Name: "OtherAttributes", Flag: "other-attributes", Type: "map[string]types.AttributeValue", Required: false},
	{Name: "SAMAccountName", Flag: "sam-account-name", Type: "*string", Required: true},
}

var fields_create_user = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
	{Name: "EmailAddress", Flag: "email-address", Type: "*string", Required: false},
	{Name: "GivenName", Flag: "given-name", Type: "*string", Required: false},
	{Name: "OtherAttributes", Flag: "other-attributes", Type: "map[string]types.AttributeValue", Required: false},
	{Name: "SAMAccountName", Flag: "sam-account-name", Type: "*string", Required: true},
	{Name: "Surname", Flag: "surname", Type: "*string", Required: false},
}

var fields_delete_group = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
	{Name: "SAMAccountName", Flag: "sam-account-name", Type: "*string", Required: true},
}

var fields_delete_user = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
	{Name: "SAMAccountName", Flag: "sam-account-name", Type: "*string", Required: true},
}

var fields_describe_group = []leanruntime.Field{
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
	{Name: "OtherAttributes", Flag: "other-attributes", Type: "[]string", Required: false},
	{Name: "Realm", Flag: "realm", Type: "*string", Required: false},
	{Name: "SAMAccountName", Flag: "sam-account-name", Type: "*string", Required: true},
}

var fields_describe_user = []leanruntime.Field{
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
	{Name: "OtherAttributes", Flag: "other-attributes", Type: "[]string", Required: false},
	{Name: "Realm", Flag: "realm", Type: "*string", Required: false},
	{Name: "SAMAccountName", Flag: "sam-account-name", Type: "*string", Required: true},
}

var fields_disable_user = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
	{Name: "SAMAccountName", Flag: "sam-account-name", Type: "*string", Required: true},
}

var fields_list_group_members = []leanruntime.Field{
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MemberRealm", Flag: "member-realm", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Realm", Flag: "realm", Type: "*string", Required: false},
	{Name: "SAMAccountName", Flag: "sam-account-name", Type: "*string", Required: true},
}

var fields_list_groups = []leanruntime.Field{
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Realm", Flag: "realm", Type: "*string", Required: false},
}

var fields_list_groups_for_member = []leanruntime.Field{
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MemberRealm", Flag: "member-realm", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Realm", Flag: "realm", Type: "*string", Required: false},
	{Name: "SAMAccountName", Flag: "sam-account-name", Type: "*string", Required: true},
}

var fields_list_users = []leanruntime.Field{
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Realm", Flag: "realm", Type: "*string", Required: false},
}

var fields_remove_group_member = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
	{Name: "GroupName", Flag: "group-name", Type: "*string", Required: true},
	{Name: "MemberName", Flag: "member-name", Type: "*string", Required: true},
	{Name: "MemberRealm", Flag: "member-realm", Type: "*string", Required: false},
}

var fields_search_groups = []leanruntime.Field{
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Realm", Flag: "realm", Type: "*string", Required: false},
	{Name: "SearchAttributes", Flag: "search-attributes", Type: "[]string", Required: true},
	{Name: "SearchString", Flag: "search-string", Type: "*string", Required: true},
}

var fields_search_users = []leanruntime.Field{
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Realm", Flag: "realm", Type: "*string", Required: false},
	{Name: "SearchAttributes", Flag: "search-attributes", Type: "[]string", Required: true},
	{Name: "SearchString", Flag: "search-string", Type: "*string", Required: true},
}

var fields_update_group = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
	{Name: "GroupScope", Flag: "group-scope", Type: "types.GroupScope", Required: false},
	{Name: "GroupType", Flag: "group-type", Type: "types.GroupType", Required: false},
	{Name: "OtherAttributes", Flag: "other-attributes", Type: "map[string]types.AttributeValue", Required: false},
	{Name: "SAMAccountName", Flag: "sam-account-name", Type: "*string", Required: true},
	{Name: "UpdateType", Flag: "update-type", Type: "types.UpdateType", Required: false},
}

var fields_update_user = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
	{Name: "EmailAddress", Flag: "email-address", Type: "*string", Required: false},
	{Name: "GivenName", Flag: "given-name", Type: "*string", Required: false},
	{Name: "OtherAttributes", Flag: "other-attributes", Type: "map[string]types.AttributeValue", Required: false},
	{Name: "SAMAccountName", Flag: "sam-account-name", Type: "*string", Required: true},
	{Name: "Surname", Flag: "surname", Type: "*string", Required: false},
	{Name: "UpdateType", Flag: "update-type", Type: "types.UpdateType", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"add-group-member": {
			Name:   "add-group-member",
			Fields: fields_add_group_member,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddGroupMemberInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_group_member, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddGroupMember(ctx, input)
			},
		},
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
		"disable-user": {
			Name:   "disable-user",
			Fields: fields_disable_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableUser(ctx, input)
			},
		},
		"list-group-members": {
			Name:   "list-group-members",
			Fields: fields_list_group_members,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListGroupMembersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_group_members, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListGroupMembers(ctx, input)
				}
				var results []*svc.ListGroupMembersOutput
				p := svc.NewListGroupMembersPaginator(client, input)
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
		"list-groups-for-member": {
			Name:   "list-groups-for-member",
			Fields: fields_list_groups_for_member,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListGroupsForMemberInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_groups_for_member, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListGroupsForMember(ctx, input)
				}
				var results []*svc.ListGroupsForMemberOutput
				p := svc.NewListGroupsForMemberPaginator(client, input)
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
		"remove-group-member": {
			Name:   "remove-group-member",
			Fields: fields_remove_group_member,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveGroupMemberInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_group_member, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveGroupMember(ctx, input)
			},
		},
		"search-groups": {
			Name:   "search-groups",
			Fields: fields_search_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchGroups(ctx, input)
				}
				var results []*svc.SearchGroupsOutput
				p := svc.NewSearchGroupsPaginator(client, input)
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
		"search-users": {
			Name:   "search-users",
			Fields: fields_search_users,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchUsersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_users, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchUsers(ctx, input)
				}
				var results []*svc.SearchUsersOutput
				p := svc.NewSearchUsersPaginator(client, input)
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
	if err := leanruntime.Execute("directoryservicedata", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
