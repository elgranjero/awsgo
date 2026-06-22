package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/acmpca"
)

var fields_create_certificate_authority = []leanruntime.Field{
	{Name: "CertificateAuthorityConfiguration", Flag: "certificate-authority-configuration", Type: "*types.CertificateAuthorityConfiguration", Required: true},
	{Name: "CertificateAuthorityType", Flag: "certificate-authority-type", Type: "types.CertificateAuthorityType", Required: true},
	{Name: "IdempotencyToken", Flag: "idempotency-token", Type: "*string", Required: false},
	{Name: "KeyStorageSecurityStandard", Flag: "key-storage-security-standard", Type: "types.KeyStorageSecurityStandard", Required: false},
	{Name: "RevocationConfiguration", Flag: "revocation-configuration", Type: "*types.RevocationConfiguration", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "UsageMode", Flag: "usage-mode", Type: "types.CertificateAuthorityUsageMode", Required: false},
}

var fields_create_certificate_authority_audit_report = []leanruntime.Field{
	{Name: "AuditReportResponseFormat", Flag: "audit-report-response-format", Type: "types.AuditReportResponseFormat", Required: true},
	{Name: "CertificateAuthorityArn", Flag: "certificate-authority-arn", Type: "*string", Required: true},
	{Name: "S3BucketName", Flag: "s3-bucket-name", Type: "*string", Required: true},
}

var fields_create_permission = []leanruntime.Field{
	{Name: "Actions", Flag: "actions", Type: "[]types.ActionType", Required: true},
	{Name: "CertificateAuthorityArn", Flag: "certificate-authority-arn", Type: "*string", Required: true},
	{Name: "Principal", Flag: "principal", Type: "*string", Required: true},
	{Name: "SourceAccount", Flag: "source-account", Type: "*string", Required: false},
}

var fields_delete_certificate_authority = []leanruntime.Field{
	{Name: "CertificateAuthorityArn", Flag: "certificate-authority-arn", Type: "*string", Required: true},
	{Name: "PermanentDeletionTimeInDays", Flag: "permanent-deletion-time-in-days", Type: "*int32", Required: false},
}

var fields_delete_permission = []leanruntime.Field{
	{Name: "CertificateAuthorityArn", Flag: "certificate-authority-arn", Type: "*string", Required: true},
	{Name: "Principal", Flag: "principal", Type: "*string", Required: true},
	{Name: "SourceAccount", Flag: "source-account", Type: "*string", Required: false},
}

var fields_delete_policy = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_describe_certificate_authority = []leanruntime.Field{
	{Name: "CertificateAuthorityArn", Flag: "certificate-authority-arn", Type: "*string", Required: true},
}

var fields_describe_certificate_authority_audit_report = []leanruntime.Field{
	{Name: "AuditReportId", Flag: "audit-report-id", Type: "*string", Required: true},
	{Name: "CertificateAuthorityArn", Flag: "certificate-authority-arn", Type: "*string", Required: true},
}

var fields_get_certificate = []leanruntime.Field{
	{Name: "CertificateArn", Flag: "certificate-arn", Type: "*string", Required: true},
	{Name: "CertificateAuthorityArn", Flag: "certificate-authority-arn", Type: "*string", Required: true},
}

var fields_get_certificate_authority_certificate = []leanruntime.Field{
	{Name: "CertificateAuthorityArn", Flag: "certificate-authority-arn", Type: "*string", Required: true},
}

var fields_get_certificate_authority_csr = []leanruntime.Field{
	{Name: "CertificateAuthorityArn", Flag: "certificate-authority-arn", Type: "*string", Required: true},
}

var fields_get_policy = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_import_certificate_authority_certificate = []leanruntime.Field{
	{Name: "Certificate", Flag: "certificate", Type: "[]byte", Required: true},
	{Name: "CertificateAuthorityArn", Flag: "certificate-authority-arn", Type: "*string", Required: true},
	{Name: "CertificateChain", Flag: "certificate-chain", Type: "[]byte", Required: false},
}

var fields_issue_certificate = []leanruntime.Field{
	{Name: "ApiPassthrough", Flag: "api-passthrough", Type: "*types.ApiPassthrough", Required: false},
	{Name: "CertificateAuthorityArn", Flag: "certificate-authority-arn", Type: "*string", Required: true},
	{Name: "Csr", Flag: "csr", Type: "[]byte", Required: true},
	{Name: "IdempotencyToken", Flag: "idempotency-token", Type: "*string", Required: false},
	{Name: "SigningAlgorithm", Flag: "signing-algorithm", Type: "types.SigningAlgorithm", Required: true},
	{Name: "TemplateArn", Flag: "template-arn", Type: "*string", Required: false},
	{Name: "Validity", Flag: "validity", Type: "*types.Validity", Required: true},
	{Name: "ValidityNotBefore", Flag: "validity-not-before", Type: "*types.Validity", Required: false},
}

var fields_list_certificate_authorities = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceOwner", Flag: "resource-owner", Type: "types.ResourceOwner", Required: false},
}

var fields_list_permissions = []leanruntime.Field{
	{Name: "CertificateAuthorityArn", Flag: "certificate-authority-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags = []leanruntime.Field{
	{Name: "CertificateAuthorityArn", Flag: "certificate-authority-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_put_policy = []leanruntime.Field{
	{Name: "Policy", Flag: "policy", Type: "*string", Required: true},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_restore_certificate_authority = []leanruntime.Field{
	{Name: "CertificateAuthorityArn", Flag: "certificate-authority-arn", Type: "*string", Required: true},
}

var fields_revoke_certificate = []leanruntime.Field{
	{Name: "CertificateAuthorityArn", Flag: "certificate-authority-arn", Type: "*string", Required: true},
	{Name: "CertificateSerial", Flag: "certificate-serial", Type: "*string", Required: true},
	{Name: "RevocationReason", Flag: "revocation-reason", Type: "types.RevocationReason", Required: true},
}

var fields_tag_certificate_authority = []leanruntime.Field{
	{Name: "CertificateAuthorityArn", Flag: "certificate-authority-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_certificate_authority = []leanruntime.Field{
	{Name: "CertificateAuthorityArn", Flag: "certificate-authority-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_update_certificate_authority = []leanruntime.Field{
	{Name: "CertificateAuthorityArn", Flag: "certificate-authority-arn", Type: "*string", Required: true},
	{Name: "RevocationConfiguration", Flag: "revocation-configuration", Type: "*types.RevocationConfiguration", Required: false},
	{Name: "Status", Flag: "status", Type: "types.CertificateAuthorityStatus", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-certificate-authority": {
			Name:   "create-certificate-authority",
			Fields: fields_create_certificate_authority,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCertificateAuthorityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_certificate_authority, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCertificateAuthority(ctx, input)
			},
		},
		"create-certificate-authority-audit-report": {
			Name:   "create-certificate-authority-audit-report",
			Fields: fields_create_certificate_authority_audit_report,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCertificateAuthorityAuditReportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_certificate_authority_audit_report, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCertificateAuthorityAuditReport(ctx, input)
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
		"delete-certificate-authority": {
			Name:   "delete-certificate-authority",
			Fields: fields_delete_certificate_authority,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCertificateAuthorityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_certificate_authority, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCertificateAuthority(ctx, input)
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
		"delete-policy": {
			Name:   "delete-policy",
			Fields: fields_delete_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePolicy(ctx, input)
			},
		},
		"describe-certificate-authority": {
			Name:   "describe-certificate-authority",
			Fields: fields_describe_certificate_authority,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCertificateAuthorityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_certificate_authority, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeCertificateAuthority(ctx, input)
			},
		},
		"describe-certificate-authority-audit-report": {
			Name:   "describe-certificate-authority-audit-report",
			Fields: fields_describe_certificate_authority_audit_report,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCertificateAuthorityAuditReportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_certificate_authority_audit_report, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeCertificateAuthorityAuditReport(ctx, input)
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
		"get-certificate-authority-certificate": {
			Name:   "get-certificate-authority-certificate",
			Fields: fields_get_certificate_authority_certificate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCertificateAuthorityCertificateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_certificate_authority_certificate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCertificateAuthorityCertificate(ctx, input)
			},
		},
		"get-certificate-authority-csr": {
			Name:   "get-certificate-authority-csr",
			Fields: fields_get_certificate_authority_csr,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCertificateAuthorityCsrInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_certificate_authority_csr, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCertificateAuthorityCsr(ctx, input)
			},
		},
		"get-policy": {
			Name:   "get-policy",
			Fields: fields_get_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPolicy(ctx, input)
			},
		},
		"import-certificate-authority-certificate": {
			Name:   "import-certificate-authority-certificate",
			Fields: fields_import_certificate_authority_certificate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ImportCertificateAuthorityCertificateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_import_certificate_authority_certificate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ImportCertificateAuthorityCertificate(ctx, input)
			},
		},
		"issue-certificate": {
			Name:   "issue-certificate",
			Fields: fields_issue_certificate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.IssueCertificateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_issue_certificate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.IssueCertificate(ctx, input)
			},
		},
		"list-certificate-authorities": {
			Name:   "list-certificate-authorities",
			Fields: fields_list_certificate_authorities,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCertificateAuthoritiesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_certificate_authorities, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCertificateAuthorities(ctx, input)
				}
				var results []*svc.ListCertificateAuthoritiesOutput
				p := svc.NewListCertificateAuthoritiesPaginator(client, input)
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
		"list-tags": {
			Name:   "list-tags",
			Fields: fields_list_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTagsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_tags, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTags(ctx, input)
				}
				var results []*svc.ListTagsOutput
				p := svc.NewListTagsPaginator(client, input)
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
		"put-policy": {
			Name:   "put-policy",
			Fields: fields_put_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutPolicy(ctx, input)
			},
		},
		"restore-certificate-authority": {
			Name:   "restore-certificate-authority",
			Fields: fields_restore_certificate_authority,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RestoreCertificateAuthorityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_restore_certificate_authority, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RestoreCertificateAuthority(ctx, input)
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
		"tag-certificate-authority": {
			Name:   "tag-certificate-authority",
			Fields: fields_tag_certificate_authority,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TagCertificateAuthorityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_tag_certificate_authority, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TagCertificateAuthority(ctx, input)
			},
		},
		"untag-certificate-authority": {
			Name:   "untag-certificate-authority",
			Fields: fields_untag_certificate_authority,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UntagCertificateAuthorityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_untag_certificate_authority, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UntagCertificateAuthority(ctx, input)
			},
		},
		"update-certificate-authority": {
			Name:   "update-certificate-authority",
			Fields: fields_update_certificate_authority,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateCertificateAuthorityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_certificate_authority, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateCertificateAuthority(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("acmpca", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
