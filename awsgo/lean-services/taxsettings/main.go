package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/taxsettings"
)

var fields_batch_delete_tax_registration = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]string", Required: true},
}

var fields_batch_get_tax_exemptions = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]string", Required: true},
}

var fields_batch_put_tax_registration = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]string", Required: true},
	{Name: "TaxRegistrationEntry", Flag: "tax-registration-entry", Type: "*types.TaxRegistrationEntry", Required: true},
}

var fields_delete_supplemental_tax_registration = []leanruntime.Field{
	{Name: "AuthorityId", Flag: "authority-id", Type: "*string", Required: true},
}

var fields_delete_tax_registration = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: false},
}

var fields_get_tax_exemption_types = []leanruntime.Field{}

var fields_get_tax_inheritance = []leanruntime.Field{}

var fields_get_tax_registration = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: false},
}

var fields_get_tax_registration_document = []leanruntime.Field{
	{Name: "DestinationS3Location", Flag: "destination-s3-location", Type: "*types.DestinationS3Location", Required: false},
	{Name: "TaxDocumentMetadata", Flag: "tax-document-metadata", Type: "*types.TaxDocumentMetadata", Required: true},
}

var fields_list_supplemental_tax_registrations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tax_exemptions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tax_registrations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_put_supplemental_tax_registration = []leanruntime.Field{
	{Name: "TaxRegistrationEntry", Flag: "tax-registration-entry", Type: "*types.SupplementalTaxRegistrationEntry", Required: true},
}

var fields_put_tax_exemption = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]string", Required: true},
	{Name: "Authority", Flag: "authority", Type: "*types.Authority", Required: true},
	{Name: "ExemptionCertificate", Flag: "exemption-certificate", Type: "*types.ExemptionCertificate", Required: true},
	{Name: "ExemptionType", Flag: "exemption-type", Type: "*string", Required: true},
}

var fields_put_tax_inheritance = []leanruntime.Field{
	{Name: "HeritageStatus", Flag: "heritage-status", Type: "types.HeritageStatus", Required: false},
}

var fields_put_tax_registration = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: false},
	{Name: "TaxRegistrationEntry", Flag: "tax-registration-entry", Type: "*types.TaxRegistrationEntry", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"batch-delete-tax-registration": {
			Name:   "batch-delete-tax-registration",
			Fields: fields_batch_delete_tax_registration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchDeleteTaxRegistrationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_delete_tax_registration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchDeleteTaxRegistration(ctx, input)
			},
		},
		"batch-get-tax-exemptions": {
			Name:   "batch-get-tax-exemptions",
			Fields: fields_batch_get_tax_exemptions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetTaxExemptionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_tax_exemptions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetTaxExemptions(ctx, input)
			},
		},
		"batch-put-tax-registration": {
			Name:   "batch-put-tax-registration",
			Fields: fields_batch_put_tax_registration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchPutTaxRegistrationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_put_tax_registration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchPutTaxRegistration(ctx, input)
			},
		},
		"delete-supplemental-tax-registration": {
			Name:   "delete-supplemental-tax-registration",
			Fields: fields_delete_supplemental_tax_registration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSupplementalTaxRegistrationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_supplemental_tax_registration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSupplementalTaxRegistration(ctx, input)
			},
		},
		"delete-tax-registration": {
			Name:   "delete-tax-registration",
			Fields: fields_delete_tax_registration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTaxRegistrationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_tax_registration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTaxRegistration(ctx, input)
			},
		},
		"get-tax-exemption-types": {
			Name:   "get-tax-exemption-types",
			Fields: fields_get_tax_exemption_types,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTaxExemptionTypesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_tax_exemption_types, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTaxExemptionTypes(ctx, input)
			},
		},
		"get-tax-inheritance": {
			Name:   "get-tax-inheritance",
			Fields: fields_get_tax_inheritance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTaxInheritanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_tax_inheritance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTaxInheritance(ctx, input)
			},
		},
		"get-tax-registration": {
			Name:   "get-tax-registration",
			Fields: fields_get_tax_registration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTaxRegistrationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_tax_registration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTaxRegistration(ctx, input)
			},
		},
		"get-tax-registration-document": {
			Name:   "get-tax-registration-document",
			Fields: fields_get_tax_registration_document,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTaxRegistrationDocumentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_tax_registration_document, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTaxRegistrationDocument(ctx, input)
			},
		},
		"list-supplemental-tax-registrations": {
			Name:   "list-supplemental-tax-registrations",
			Fields: fields_list_supplemental_tax_registrations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSupplementalTaxRegistrationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_supplemental_tax_registrations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSupplementalTaxRegistrations(ctx, input)
				}
				var results []*svc.ListSupplementalTaxRegistrationsOutput
				p := svc.NewListSupplementalTaxRegistrationsPaginator(client, input)
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
		"list-tax-exemptions": {
			Name:   "list-tax-exemptions",
			Fields: fields_list_tax_exemptions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTaxExemptionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_tax_exemptions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTaxExemptions(ctx, input)
				}
				var results []*svc.ListTaxExemptionsOutput
				p := svc.NewListTaxExemptionsPaginator(client, input)
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
		"list-tax-registrations": {
			Name:   "list-tax-registrations",
			Fields: fields_list_tax_registrations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTaxRegistrationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_tax_registrations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTaxRegistrations(ctx, input)
				}
				var results []*svc.ListTaxRegistrationsOutput
				p := svc.NewListTaxRegistrationsPaginator(client, input)
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
		"put-supplemental-tax-registration": {
			Name:   "put-supplemental-tax-registration",
			Fields: fields_put_supplemental_tax_registration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutSupplementalTaxRegistrationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_supplemental_tax_registration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutSupplementalTaxRegistration(ctx, input)
			},
		},
		"put-tax-exemption": {
			Name:   "put-tax-exemption",
			Fields: fields_put_tax_exemption,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutTaxExemptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_tax_exemption, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutTaxExemption(ctx, input)
			},
		},
		"put-tax-inheritance": {
			Name:   "put-tax-inheritance",
			Fields: fields_put_tax_inheritance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutTaxInheritanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_tax_inheritance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutTaxInheritance(ctx, input)
			},
		},
		"put-tax-registration": {
			Name:   "put-tax-registration",
			Fields: fields_put_tax_registration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutTaxRegistrationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_tax_registration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutTaxRegistration(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("taxsettings", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
