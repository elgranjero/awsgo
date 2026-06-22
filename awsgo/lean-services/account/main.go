package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/account"
)

var fields_accept_primary_email_update = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "Otp", Flag: "otp", Type: "*string", Required: true},
	{Name: "PrimaryEmail", Flag: "primary-email", Type: "*string", Required: true},
}

var fields_delete_alternate_contact = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: false},
	{Name: "AlternateContactType", Flag: "alternate-contact-type", Type: "types.AlternateContactType", Required: true},
}

var fields_disable_region = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: false},
	{Name: "RegionName", Flag: "region-name", Type: "*string", Required: true},
}

var fields_enable_region = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: false},
	{Name: "RegionName", Flag: "region-name", Type: "*string", Required: true},
}

var fields_get_account_information = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: false},
}

var fields_get_alternate_contact = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: false},
	{Name: "AlternateContactType", Flag: "alternate-contact-type", Type: "types.AlternateContactType", Required: true},
}

var fields_get_contact_information = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: false},
}

var fields_get_gov_cloud_account_information = []leanruntime.Field{
	{Name: "StandardAccountId", Flag: "standard-account-id", Type: "*string", Required: false},
}

var fields_get_primary_email = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
}

var fields_get_region_opt_status = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: false},
	{Name: "RegionName", Flag: "region-name", Type: "*string", Required: true},
}

var fields_list_regions = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RegionOptStatusContains", Flag: "region-opt-status-contains", Type: "[]types.RegionOptStatus", Required: false},
}

var fields_put_account_name = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: false},
	{Name: "AccountName", Flag: "account-name", Type: "*string", Required: true},
}

var fields_put_alternate_contact = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: false},
	{Name: "AlternateContactType", Flag: "alternate-contact-type", Type: "types.AlternateContactType", Required: true},
	{Name: "EmailAddress", Flag: "email-address", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "PhoneNumber", Flag: "phone-number", Type: "*string", Required: true},
	{Name: "Title", Flag: "title", Type: "*string", Required: true},
}

var fields_put_contact_information = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: false},
	{Name: "ContactInformation", Flag: "contact-information", Type: "*types.ContactInformation", Required: true},
}

var fields_start_primary_email_update = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "PrimaryEmail", Flag: "primary-email", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"accept-primary-email-update": {
			Name:   "accept-primary-email-update",
			Fields: fields_accept_primary_email_update,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AcceptPrimaryEmailUpdateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_accept_primary_email_update, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AcceptPrimaryEmailUpdate(ctx, input)
			},
		},
		"delete-alternate-contact": {
			Name:   "delete-alternate-contact",
			Fields: fields_delete_alternate_contact,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAlternateContactInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_alternate_contact, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAlternateContact(ctx, input)
			},
		},
		"disable-region": {
			Name:   "disable-region",
			Fields: fields_disable_region,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableRegionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_region, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableRegion(ctx, input)
			},
		},
		"enable-region": {
			Name:   "enable-region",
			Fields: fields_enable_region,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableRegionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_region, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableRegion(ctx, input)
			},
		},
		"get-account-information": {
			Name:   "get-account-information",
			Fields: fields_get_account_information,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAccountInformationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_account_information, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAccountInformation(ctx, input)
			},
		},
		"get-alternate-contact": {
			Name:   "get-alternate-contact",
			Fields: fields_get_alternate_contact,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAlternateContactInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_alternate_contact, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAlternateContact(ctx, input)
			},
		},
		"get-contact-information": {
			Name:   "get-contact-information",
			Fields: fields_get_contact_information,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetContactInformationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_contact_information, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetContactInformation(ctx, input)
			},
		},
		"get-gov-cloud-account-information": {
			Name:   "get-gov-cloud-account-information",
			Fields: fields_get_gov_cloud_account_information,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetGovCloudAccountInformationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_gov_cloud_account_information, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetGovCloudAccountInformation(ctx, input)
			},
		},
		"get-primary-email": {
			Name:   "get-primary-email",
			Fields: fields_get_primary_email,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPrimaryEmailInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_primary_email, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPrimaryEmail(ctx, input)
			},
		},
		"get-region-opt-status": {
			Name:   "get-region-opt-status",
			Fields: fields_get_region_opt_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRegionOptStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_region_opt_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRegionOptStatus(ctx, input)
			},
		},
		"list-regions": {
			Name:   "list-regions",
			Fields: fields_list_regions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRegionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_regions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRegions(ctx, input)
				}
				var results []*svc.ListRegionsOutput
				p := svc.NewListRegionsPaginator(client, input)
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
		"put-account-name": {
			Name:   "put-account-name",
			Fields: fields_put_account_name,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutAccountNameInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_account_name, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutAccountName(ctx, input)
			},
		},
		"put-alternate-contact": {
			Name:   "put-alternate-contact",
			Fields: fields_put_alternate_contact,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutAlternateContactInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_alternate_contact, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutAlternateContact(ctx, input)
			},
		},
		"put-contact-information": {
			Name:   "put-contact-information",
			Fields: fields_put_contact_information,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutContactInformationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_contact_information, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutContactInformation(ctx, input)
			},
		},
		"start-primary-email-update": {
			Name:   "start-primary-email-update",
			Fields: fields_start_primary_email_update,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartPrimaryEmailUpdateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_primary_email_update, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartPrimaryEmailUpdate(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("account", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
