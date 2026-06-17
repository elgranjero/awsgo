package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/taxsettings/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"batch-delete-tax-registration", "batch-get-tax-exemptions", "batch-put-tax-registration", "delete-supplemental-tax-registration", "delete-tax-registration", "get-tax-exemption-types", "get-tax-inheritance", "get-tax-registration", "get-tax-registration-document", "list-supplemental-tax-registrations", "list-tax-exemptions", "list-tax-registrations", "put-supplemental-tax-registration", "put-tax-exemption", "put-tax-inheritance", "put-tax-registration"},
		OperationSet: map[string]bool{"batch-delete-tax-registration": true, "batch-get-tax-exemptions": true, "batch-put-tax-registration": true, "delete-supplemental-tax-registration": true, "delete-tax-registration": true, "get-tax-exemption-types": true, "get-tax-inheritance": true, "get-tax-registration": true, "get-tax-registration-document": true, "list-supplemental-tax-registrations": true, "list-tax-exemptions": true, "list-tax-registrations": true, "put-supplemental-tax-registration": true, "put-tax-exemption": true, "put-tax-inheritance": true, "put-tax-registration": true},
		OperationInputs: map[string][]string{
			"batch-delete-tax-registration":        {"AccountIds"},
			"batch-get-tax-exemptions":             {"AccountIds"},
			"batch-put-tax-registration":           {"AccountIds", "TaxRegistrationEntry"},
			"delete-supplemental-tax-registration": {"AuthorityId"},
			"delete-tax-registration":              {"AccountId"},
			"get-tax-exemption-types":              {},
			"get-tax-inheritance":                  {},
			"get-tax-registration":                 {"AccountId"},
			"get-tax-registration-document":        {"DestinationS3Location", "TaxDocumentMetadata"},
			"list-supplemental-tax-registrations":  {"MaxResults", "NextToken"},
			"list-tax-exemptions":                  {"MaxResults", "NextToken"},
			"list-tax-registrations":               {"MaxResults", "NextToken"},
			"put-supplemental-tax-registration":    {"TaxRegistrationEntry"},
			"put-tax-exemption":                    {"AccountIds", "Authority", "ExemptionCertificate", "ExemptionType"},
			"put-tax-inheritance":                  {"HeritageStatus"},
			"put-tax-registration":                 {"AccountId", "TaxRegistrationEntry"},
		},
		OperationInputTypes: map[string]map[string]string{
			"batch-delete-tax-registration":        {"AccountIds": "[]string"},
			"batch-get-tax-exemptions":             {"AccountIds": "[]string"},
			"batch-put-tax-registration":           {"AccountIds": "[]string", "TaxRegistrationEntry": "*types.TaxRegistrationEntry"},
			"delete-supplemental-tax-registration": {"AuthorityId": "*string"},
			"delete-tax-registration":              {"AccountId": "*string"},
			"get-tax-exemption-types":              {},
			"get-tax-inheritance":                  {},
			"get-tax-registration":                 {"AccountId": "*string"},
			"get-tax-registration-document":        {"DestinationS3Location": "*types.DestinationS3Location", "TaxDocumentMetadata": "*types.TaxDocumentMetadata"},
			"list-supplemental-tax-registrations":  {"MaxResults": "*int32", "NextToken": "*string"},
			"list-tax-exemptions":                  {"MaxResults": "*int32", "NextToken": "*string"},
			"list-tax-registrations":               {"MaxResults": "*int32", "NextToken": "*string"},
			"put-supplemental-tax-registration":    {"TaxRegistrationEntry": "*types.SupplementalTaxRegistrationEntry"},
			"put-tax-exemption":                    {"AccountIds": "[]string", "Authority": "*types.Authority", "ExemptionCertificate": "*types.ExemptionCertificate", "ExemptionType": "*string"},
			"put-tax-inheritance":                  {"HeritageStatus": "types.HeritageStatus"},
			"put-tax-registration":                 {"AccountId": "*string", "TaxRegistrationEntry": "*types.TaxRegistrationEntry"},
		},
		OperationInputRequired: map[string][]string{
			"batch-delete-tax-registration":        {"AccountIds"},
			"batch-get-tax-exemptions":             {"AccountIds"},
			"batch-put-tax-registration":           {"AccountIds", "TaxRegistrationEntry"},
			"delete-supplemental-tax-registration": {"AuthorityId"},
			"delete-tax-registration":              {},
			"get-tax-exemption-types":              {},
			"get-tax-inheritance":                  {},
			"get-tax-registration":                 {},
			"get-tax-registration-document":        {"TaxDocumentMetadata"},
			"list-supplemental-tax-registrations":  {},
			"list-tax-exemptions":                  {},
			"list-tax-registrations":               {},
			"put-supplemental-tax-registration":    {"TaxRegistrationEntry"},
			"put-tax-exemption":                    {"AccountIds", "Authority", "ExemptionCertificate", "ExemptionType"},
			"put-tax-inheritance":                  {},
			"put-tax-registration":                 {"TaxRegistrationEntry"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("taxsettings", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
