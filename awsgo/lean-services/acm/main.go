package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/acm"
)

var fields_add_tags_to_certificate = []leanruntime.Field{
	{Name: "CertificateArn", Flag: "certificate-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_delete_certificate = []leanruntime.Field{
	{Name: "CertificateArn", Flag: "certificate-arn", Type: "*string", Required: true},
}

var fields_describe_certificate = []leanruntime.Field{
	{Name: "CertificateArn", Flag: "certificate-arn", Type: "*string", Required: true},
}

var fields_export_certificate = []leanruntime.Field{
	{Name: "CertificateArn", Flag: "certificate-arn", Type: "*string", Required: true},
	{Name: "Passphrase", Flag: "passphrase", Type: "[]byte", Required: true},
}

var fields_get_account_configuration = []leanruntime.Field{}

var fields_get_certificate = []leanruntime.Field{
	{Name: "CertificateArn", Flag: "certificate-arn", Type: "*string", Required: true},
}

var fields_import_certificate = []leanruntime.Field{
	{Name: "Certificate", Flag: "certificate", Type: "[]byte", Required: true},
	{Name: "CertificateArn", Flag: "certificate-arn", Type: "*string", Required: false},
	{Name: "CertificateChain", Flag: "certificate-chain", Type: "[]byte", Required: false},
	{Name: "PrivateKey", Flag: "private-key", Type: "[]byte", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_list_certificates = []leanruntime.Field{
	{Name: "CertificateStatuses", Flag: "certificate-statuses", Type: "[]types.CertificateStatus", Required: false},
	{Name: "Includes", Flag: "includes", Type: "*types.Filters", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.SortBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
}

var fields_list_tags_for_certificate = []leanruntime.Field{
	{Name: "CertificateArn", Flag: "certificate-arn", Type: "*string", Required: true},
}

var fields_put_account_configuration = []leanruntime.Field{
	{Name: "ExpiryEvents", Flag: "expiry-events", Type: "*types.ExpiryEventsConfiguration", Required: false},
	{Name: "IdempotencyToken", Flag: "idempotency-token", Type: "*string", Required: true},
}

var fields_remove_tags_from_certificate = []leanruntime.Field{
	{Name: "CertificateArn", Flag: "certificate-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_renew_certificate = []leanruntime.Field{
	{Name: "CertificateArn", Flag: "certificate-arn", Type: "*string", Required: true},
}

var fields_request_certificate = []leanruntime.Field{
	{Name: "CertificateAuthorityArn", Flag: "certificate-authority-arn", Type: "*string", Required: false},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "DomainValidationOptions", Flag: "domain-validation-options", Type: "[]types.DomainValidationOption", Required: false},
	{Name: "IdempotencyToken", Flag: "idempotency-token", Type: "*string", Required: false},
	{Name: "KeyAlgorithm", Flag: "key-algorithm", Type: "types.KeyAlgorithm", Required: false},
	{Name: "ManagedBy", Flag: "managed-by", Type: "types.CertificateManagedBy", Required: false},
	{Name: "Options", Flag: "options", Type: "*types.CertificateOptions", Required: false},
	{Name: "SubjectAlternativeNames", Flag: "subject-alternative-names", Type: "[]string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "ValidationMethod", Flag: "validation-method", Type: "types.ValidationMethod", Required: false},
}

var fields_resend_validation_email = []leanruntime.Field{
	{Name: "CertificateArn", Flag: "certificate-arn", Type: "*string", Required: true},
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "ValidationDomain", Flag: "validation-domain", Type: "*string", Required: true},
}

var fields_revoke_certificate = []leanruntime.Field{
	{Name: "CertificateArn", Flag: "certificate-arn", Type: "*string", Required: true},
	{Name: "RevocationReason", Flag: "revocation-reason", Type: "types.RevocationReason", Required: true},
}

var fields_update_certificate_options = []leanruntime.Field{
	{Name: "CertificateArn", Flag: "certificate-arn", Type: "*string", Required: true},
	{Name: "Options", Flag: "options", Type: "*types.CertificateOptions", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"add-tags-to-certificate": {
			Name:   "add-tags-to-certificate",
			Fields: fields_add_tags_to_certificate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddTagsToCertificateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_tags_to_certificate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddTagsToCertificate(ctx, input)
			},
		},
		"delete-certificate": {
			Name:   "delete-certificate",
			Fields: fields_delete_certificate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCertificateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_certificate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCertificate(ctx, input)
			},
		},
		"describe-certificate": {
			Name:   "describe-certificate",
			Fields: fields_describe_certificate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCertificateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_certificate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeCertificate(ctx, input)
			},
		},
		"export-certificate": {
			Name:   "export-certificate",
			Fields: fields_export_certificate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ExportCertificateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_export_certificate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ExportCertificate(ctx, input)
			},
		},
		"get-account-configuration": {
			Name:   "get-account-configuration",
			Fields: fields_get_account_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAccountConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_account_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAccountConfiguration(ctx, input)
			},
		},
		"get-certificate": {
			Name:   "get-certificate",
			Fields: fields_get_certificate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCertificateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_certificate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCertificate(ctx, input)
			},
		},
		"import-certificate": {
			Name:   "import-certificate",
			Fields: fields_import_certificate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ImportCertificateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_import_certificate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ImportCertificate(ctx, input)
			},
		},
		"list-certificates": {
			Name:   "list-certificates",
			Fields: fields_list_certificates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCertificatesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_certificates, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCertificates(ctx, input)
				}
				var results []*svc.ListCertificatesOutput
				p := svc.NewListCertificatesPaginator(client, input)
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
		"list-tags-for-certificate": {
			Name:   "list-tags-for-certificate",
			Fields: fields_list_tags_for_certificate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTagsForCertificateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_tags_for_certificate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListTagsForCertificate(ctx, input)
			},
		},
		"put-account-configuration": {
			Name:   "put-account-configuration",
			Fields: fields_put_account_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutAccountConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_account_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutAccountConfiguration(ctx, input)
			},
		},
		"remove-tags-from-certificate": {
			Name:   "remove-tags-from-certificate",
			Fields: fields_remove_tags_from_certificate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveTagsFromCertificateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_tags_from_certificate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveTagsFromCertificate(ctx, input)
			},
		},
		"renew-certificate": {
			Name:   "renew-certificate",
			Fields: fields_renew_certificate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RenewCertificateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_renew_certificate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RenewCertificate(ctx, input)
			},
		},
		"request-certificate": {
			Name:   "request-certificate",
			Fields: fields_request_certificate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RequestCertificateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_request_certificate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RequestCertificate(ctx, input)
			},
		},
		"resend-validation-email": {
			Name:   "resend-validation-email",
			Fields: fields_resend_validation_email,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ResendValidationEmailInput{}
				if _, err := leanruntime.ApplyInput(input, fields_resend_validation_email, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ResendValidationEmail(ctx, input)
			},
		},
		"revoke-certificate": {
			Name:   "revoke-certificate",
			Fields: fields_revoke_certificate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RevokeCertificateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_revoke_certificate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RevokeCertificate(ctx, input)
			},
		},
		"update-certificate-options": {
			Name:   "update-certificate-options",
			Fields: fields_update_certificate_options,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateCertificateOptionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_certificate_options, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateCertificateOptions(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("acm", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
