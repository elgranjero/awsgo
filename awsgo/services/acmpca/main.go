package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/acmpca/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"create-certificate-authority", "create-certificate-authority-audit-report", "create-permission", "delete-certificate-authority", "delete-permission", "delete-policy", "describe-certificate-authority", "describe-certificate-authority-audit-report", "get-certificate", "get-certificate-authority-certificate", "get-certificate-authority-csr", "get-policy", "import-certificate-authority-certificate", "issue-certificate", "list-certificate-authorities", "list-permissions", "list-tags", "put-policy", "restore-certificate-authority", "revoke-certificate", "tag-certificate-authority", "untag-certificate-authority", "update-certificate-authority"},
		OperationSet: map[string]bool{"create-certificate-authority": true, "create-certificate-authority-audit-report": true, "create-permission": true, "delete-certificate-authority": true, "delete-permission": true, "delete-policy": true, "describe-certificate-authority": true, "describe-certificate-authority-audit-report": true, "get-certificate": true, "get-certificate-authority-certificate": true, "get-certificate-authority-csr": true, "get-policy": true, "import-certificate-authority-certificate": true, "issue-certificate": true, "list-certificate-authorities": true, "list-permissions": true, "list-tags": true, "put-policy": true, "restore-certificate-authority": true, "revoke-certificate": true, "tag-certificate-authority": true, "untag-certificate-authority": true, "update-certificate-authority": true},
		OperationInputs: map[string][]string{
			"create-certificate-authority":                {"CertificateAuthorityConfiguration", "CertificateAuthorityType", "IdempotencyToken", "KeyStorageSecurityStandard", "RevocationConfiguration", "Tags", "UsageMode"},
			"create-certificate-authority-audit-report":   {"AuditReportResponseFormat", "CertificateAuthorityArn", "S3BucketName"},
			"create-permission":                           {"Actions", "CertificateAuthorityArn", "Principal", "SourceAccount"},
			"delete-certificate-authority":                {"CertificateAuthorityArn", "PermanentDeletionTimeInDays"},
			"delete-permission":                           {"CertificateAuthorityArn", "Principal", "SourceAccount"},
			"delete-policy":                               {"ResourceArn"},
			"describe-certificate-authority":              {"CertificateAuthorityArn"},
			"describe-certificate-authority-audit-report": {"AuditReportId", "CertificateAuthorityArn"},
			"get-certificate":                             {"CertificateArn", "CertificateAuthorityArn"},
			"get-certificate-authority-certificate":       {"CertificateAuthorityArn"},
			"get-certificate-authority-csr":               {"CertificateAuthorityArn"},
			"get-policy":                                  {"ResourceArn"},
			"import-certificate-authority-certificate":    {"Certificate", "CertificateAuthorityArn", "CertificateChain"},
			"issue-certificate":                           {"ApiPassthrough", "CertificateAuthorityArn", "Csr", "IdempotencyToken", "SigningAlgorithm", "TemplateArn", "Validity", "ValidityNotBefore"},
			"list-certificate-authorities":                {"MaxResults", "NextToken", "ResourceOwner"},
			"list-permissions":                            {"CertificateAuthorityArn", "MaxResults", "NextToken"},
			"list-tags":                                   {"CertificateAuthorityArn", "MaxResults", "NextToken"},
			"put-policy":                                  {"Policy", "ResourceArn"},
			"restore-certificate-authority":               {"CertificateAuthorityArn"},
			"revoke-certificate":                          {"CertificateAuthorityArn", "CertificateSerial", "RevocationReason"},
			"tag-certificate-authority":                   {"CertificateAuthorityArn", "Tags"},
			"untag-certificate-authority":                 {"CertificateAuthorityArn", "Tags"},
			"update-certificate-authority":                {"CertificateAuthorityArn", "RevocationConfiguration", "Status"},
		},
		OperationInputTypes: map[string]map[string]string{
			"create-certificate-authority":                {"CertificateAuthorityConfiguration": "*types.CertificateAuthorityConfiguration", "CertificateAuthorityType": "types.CertificateAuthorityType", "IdempotencyToken": "*string", "KeyStorageSecurityStandard": "types.KeyStorageSecurityStandard", "RevocationConfiguration": "*types.RevocationConfiguration", "Tags": "[]types.Tag", "UsageMode": "types.CertificateAuthorityUsageMode"},
			"create-certificate-authority-audit-report":   {"AuditReportResponseFormat": "types.AuditReportResponseFormat", "CertificateAuthorityArn": "*string", "S3BucketName": "*string"},
			"create-permission":                           {"Actions": "[]types.ActionType", "CertificateAuthorityArn": "*string", "Principal": "*string", "SourceAccount": "*string"},
			"delete-certificate-authority":                {"CertificateAuthorityArn": "*string", "PermanentDeletionTimeInDays": "*int32"},
			"delete-permission":                           {"CertificateAuthorityArn": "*string", "Principal": "*string", "SourceAccount": "*string"},
			"delete-policy":                               {"ResourceArn": "*string"},
			"describe-certificate-authority":              {"CertificateAuthorityArn": "*string"},
			"describe-certificate-authority-audit-report": {"AuditReportId": "*string", "CertificateAuthorityArn": "*string"},
			"get-certificate":                             {"CertificateArn": "*string", "CertificateAuthorityArn": "*string"},
			"get-certificate-authority-certificate":       {"CertificateAuthorityArn": "*string"},
			"get-certificate-authority-csr":               {"CertificateAuthorityArn": "*string"},
			"get-policy":                                  {"ResourceArn": "*string"},
			"import-certificate-authority-certificate":    {"Certificate": "[]byte", "CertificateAuthorityArn": "*string", "CertificateChain": "[]byte"},
			"issue-certificate":                           {"ApiPassthrough": "*types.ApiPassthrough", "CertificateAuthorityArn": "*string", "Csr": "[]byte", "IdempotencyToken": "*string", "SigningAlgorithm": "types.SigningAlgorithm", "TemplateArn": "*string", "Validity": "*types.Validity", "ValidityNotBefore": "*types.Validity"},
			"list-certificate-authorities":                {"MaxResults": "*int32", "NextToken": "*string", "ResourceOwner": "types.ResourceOwner"},
			"list-permissions":                            {"CertificateAuthorityArn": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"list-tags":                                   {"CertificateAuthorityArn": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"put-policy":                                  {"Policy": "*string", "ResourceArn": "*string"},
			"restore-certificate-authority":               {"CertificateAuthorityArn": "*string"},
			"revoke-certificate":                          {"CertificateAuthorityArn": "*string", "CertificateSerial": "*string", "RevocationReason": "types.RevocationReason"},
			"tag-certificate-authority":                   {"CertificateAuthorityArn": "*string", "Tags": "[]types.Tag"},
			"untag-certificate-authority":                 {"CertificateAuthorityArn": "*string", "Tags": "[]types.Tag"},
			"update-certificate-authority":                {"CertificateAuthorityArn": "*string", "RevocationConfiguration": "*types.RevocationConfiguration", "Status": "types.CertificateAuthorityStatus"},
		},
		OperationInputRequired: map[string][]string{
			"create-certificate-authority":                {"CertificateAuthorityConfiguration", "CertificateAuthorityType"},
			"create-certificate-authority-audit-report":   {"AuditReportResponseFormat", "CertificateAuthorityArn", "S3BucketName"},
			"create-permission":                           {"Actions", "CertificateAuthorityArn", "Principal"},
			"delete-certificate-authority":                {"CertificateAuthorityArn"},
			"delete-permission":                           {"CertificateAuthorityArn", "Principal"},
			"delete-policy":                               {"ResourceArn"},
			"describe-certificate-authority":              {"CertificateAuthorityArn"},
			"describe-certificate-authority-audit-report": {"AuditReportId", "CertificateAuthorityArn"},
			"get-certificate":                             {"CertificateArn", "CertificateAuthorityArn"},
			"get-certificate-authority-certificate":       {"CertificateAuthorityArn"},
			"get-certificate-authority-csr":               {"CertificateAuthorityArn"},
			"get-policy":                                  {"ResourceArn"},
			"import-certificate-authority-certificate":    {"Certificate", "CertificateAuthorityArn"},
			"issue-certificate":                           {"CertificateAuthorityArn", "Csr", "SigningAlgorithm", "Validity"},
			"list-certificate-authorities":                {},
			"list-permissions":                            {"CertificateAuthorityArn"},
			"list-tags":                                   {"CertificateAuthorityArn"},
			"put-policy":                                  {"Policy", "ResourceArn"},
			"restore-certificate-authority":               {"CertificateAuthorityArn"},
			"revoke-certificate":                          {"CertificateAuthorityArn", "CertificateSerial", "RevocationReason"},
			"tag-certificate-authority":                   {"CertificateAuthorityArn", "Tags"},
			"untag-certificate-authority":                 {"CertificateAuthorityArn", "Tags"},
			"update-certificate-authority":                {"CertificateAuthorityArn"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("acmpca", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
