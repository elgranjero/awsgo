package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/acm/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"add-tags-to-certificate", "delete-certificate", "describe-certificate", "export-certificate", "get-account-configuration", "get-certificate", "import-certificate", "list-certificates", "list-tags-for-certificate", "put-account-configuration", "remove-tags-from-certificate", "renew-certificate", "request-certificate", "resend-validation-email", "revoke-certificate", "update-certificate-options"},
		OperationSet: map[string]bool{"add-tags-to-certificate": true, "delete-certificate": true, "describe-certificate": true, "export-certificate": true, "get-account-configuration": true, "get-certificate": true, "import-certificate": true, "list-certificates": true, "list-tags-for-certificate": true, "put-account-configuration": true, "remove-tags-from-certificate": true, "renew-certificate": true, "request-certificate": true, "resend-validation-email": true, "revoke-certificate": true, "update-certificate-options": true},
		OperationInputs: map[string][]string{
			"add-tags-to-certificate":      {"CertificateArn", "Tags"},
			"delete-certificate":           {"CertificateArn"},
			"describe-certificate":         {"CertificateArn"},
			"export-certificate":           {"CertificateArn", "Passphrase"},
			"get-account-configuration":    {},
			"get-certificate":              {"CertificateArn"},
			"import-certificate":           {"Certificate", "CertificateArn", "CertificateChain", "PrivateKey", "Tags"},
			"list-certificates":            {"CertificateStatuses", "Includes", "MaxItems", "NextToken", "SortBy", "SortOrder"},
			"list-tags-for-certificate":    {"CertificateArn"},
			"put-account-configuration":    {"ExpiryEvents", "IdempotencyToken"},
			"remove-tags-from-certificate": {"CertificateArn", "Tags"},
			"renew-certificate":            {"CertificateArn"},
			"request-certificate":          {"CertificateAuthorityArn", "DomainName", "DomainValidationOptions", "IdempotencyToken", "KeyAlgorithm", "ManagedBy", "Options", "SubjectAlternativeNames", "Tags", "ValidationMethod"},
			"resend-validation-email":      {"CertificateArn", "Domain", "ValidationDomain"},
			"revoke-certificate":           {"CertificateArn", "RevocationReason"},
			"update-certificate-options":   {"CertificateArn", "Options"},
		},
		OperationInputTypes: map[string]map[string]string{
			"add-tags-to-certificate":      {"CertificateArn": "*string", "Tags": "[]types.Tag"},
			"delete-certificate":           {"CertificateArn": "*string"},
			"describe-certificate":         {"CertificateArn": "*string"},
			"export-certificate":           {"CertificateArn": "*string", "Passphrase": "[]byte"},
			"get-account-configuration":    {},
			"get-certificate":              {"CertificateArn": "*string"},
			"import-certificate":           {"Certificate": "[]byte", "CertificateArn": "*string", "CertificateChain": "[]byte", "PrivateKey": "[]byte", "Tags": "[]types.Tag"},
			"list-certificates":            {"CertificateStatuses": "[]types.CertificateStatus", "Includes": "*types.Filters", "MaxItems": "*int32", "NextToken": "*string", "SortBy": "types.SortBy", "SortOrder": "types.SortOrder"},
			"list-tags-for-certificate":    {"CertificateArn": "*string"},
			"put-account-configuration":    {"ExpiryEvents": "*types.ExpiryEventsConfiguration", "IdempotencyToken": "*string"},
			"remove-tags-from-certificate": {"CertificateArn": "*string", "Tags": "[]types.Tag"},
			"renew-certificate":            {"CertificateArn": "*string"},
			"request-certificate":          {"CertificateAuthorityArn": "*string", "DomainName": "*string", "DomainValidationOptions": "[]types.DomainValidationOption", "IdempotencyToken": "*string", "KeyAlgorithm": "types.KeyAlgorithm", "ManagedBy": "types.CertificateManagedBy", "Options": "*types.CertificateOptions", "SubjectAlternativeNames": "[]string", "Tags": "[]types.Tag", "ValidationMethod": "types.ValidationMethod"},
			"resend-validation-email":      {"CertificateArn": "*string", "Domain": "*string", "ValidationDomain": "*string"},
			"revoke-certificate":           {"CertificateArn": "*string", "RevocationReason": "types.RevocationReason"},
			"update-certificate-options":   {"CertificateArn": "*string", "Options": "*types.CertificateOptions"},
		},
		OperationInputRequired: map[string][]string{
			"add-tags-to-certificate":      {"CertificateArn", "Tags"},
			"delete-certificate":           {"CertificateArn"},
			"describe-certificate":         {"CertificateArn"},
			"export-certificate":           {"CertificateArn", "Passphrase"},
			"get-account-configuration":    {},
			"get-certificate":              {"CertificateArn"},
			"import-certificate":           {"Certificate", "PrivateKey"},
			"list-certificates":            {},
			"list-tags-for-certificate":    {"CertificateArn"},
			"put-account-configuration":    {"IdempotencyToken"},
			"remove-tags-from-certificate": {"CertificateArn", "Tags"},
			"renew-certificate":            {"CertificateArn"},
			"request-certificate":          {"DomainName"},
			"resend-validation-email":      {"CertificateArn", "Domain", "ValidationDomain"},
			"revoke-certificate":           {"CertificateArn", "RevocationReason"},
			"update-certificate-options":   {"CertificateArn", "Options"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("acm", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
