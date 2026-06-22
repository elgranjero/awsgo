package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/sso"
)

var fields_get_role_credentials = []leanruntime.Field{
	{Name: "AccessToken", Flag: "access-token", Type: "*string", Required: true},
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "RoleName", Flag: "role-name", Type: "*string", Required: true},
}

var fields_list_account_roles = []leanruntime.Field{
	{Name: "AccessToken", Flag: "access-token", Type: "*string", Required: true},
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_accounts = []leanruntime.Field{
	{Name: "AccessToken", Flag: "access-token", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_logout = []leanruntime.Field{
	{Name: "AccessToken", Flag: "access-token", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"get-role-credentials": {
			Name:   "get-role-credentials",
			Fields: fields_get_role_credentials,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRoleCredentialsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_role_credentials, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRoleCredentials(ctx, input)
			},
		},
		"list-account-roles": {
			Name:   "list-account-roles",
			Fields: fields_list_account_roles,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAccountRolesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_account_roles, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAccountRoles(ctx, input)
				}
				var results []*svc.ListAccountRolesOutput
				p := svc.NewListAccountRolesPaginator(client, input)
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
		"list-accounts": {
			Name:   "list-accounts",
			Fields: fields_list_accounts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAccountsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_accounts, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAccounts(ctx, input)
				}
				var results []*svc.ListAccountsOutput
				p := svc.NewListAccountsPaginator(client, input)
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
		"logout": {
			Name:   "logout",
			Fields: fields_logout,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.LogoutInput{}
				if _, err := leanruntime.ApplyInput(input, fields_logout, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.Logout(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("sso", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
