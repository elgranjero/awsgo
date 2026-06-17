package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/partnercentralbenefits/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"amend-benefit-application", "associate-benefit-application-resource", "cancel-benefit-application", "create-benefit-application", "disassociate-benefit-application-resource", "get-benefit", "get-benefit-allocation", "get-benefit-application", "list-benefit-allocations", "list-benefit-applications", "list-benefits", "list-tags-for-resource", "recall-benefit-application", "submit-benefit-application", "tag-resource", "untag-resource", "update-benefit-application"},
		OperationSet: map[string]bool{"amend-benefit-application": true, "associate-benefit-application-resource": true, "cancel-benefit-application": true, "create-benefit-application": true, "disassociate-benefit-application-resource": true, "get-benefit": true, "get-benefit-allocation": true, "get-benefit-application": true, "list-benefit-allocations": true, "list-benefit-applications": true, "list-benefits": true, "list-tags-for-resource": true, "recall-benefit-application": true, "submit-benefit-application": true, "tag-resource": true, "untag-resource": true, "update-benefit-application": true},
		OperationInputs: map[string][]string{
			"amend-benefit-application":                 {"AmendmentReason", "Amendments", "Catalog", "ClientToken", "Identifier", "Revision"},
			"associate-benefit-application-resource":    {"BenefitApplicationIdentifier", "Catalog", "ResourceArn"},
			"cancel-benefit-application":                {"Catalog", "ClientToken", "Identifier", "Reason"},
			"create-benefit-application":                {"AssociatedResources", "BenefitApplicationDetails", "BenefitIdentifier", "Catalog", "ClientToken", "Description", "FileDetails", "FulfillmentTypes", "Name", "PartnerContacts", "Tags"},
			"disassociate-benefit-application-resource": {"BenefitApplicationIdentifier", "Catalog", "ResourceArn"},
			"get-benefit":                {"Catalog", "Identifier"},
			"get-benefit-allocation":     {"Catalog", "Identifier"},
			"get-benefit-application":    {"Catalog", "Identifier"},
			"list-benefit-allocations":   {"BenefitApplicationIdentifiers", "BenefitIdentifiers", "Catalog", "FulfillmentTypes", "MaxResults", "NextToken", "Status"},
			"list-benefit-applications":  {"AssociatedResourceArns", "AssociatedResources", "BenefitIdentifiers", "Catalog", "FulfillmentTypes", "MaxResults", "NextToken", "Programs", "Stages", "Status"},
			"list-benefits":              {"Catalog", "FulfillmentTypes", "MaxResults", "NextToken", "Programs", "Status"},
			"list-tags-for-resource":     {"ResourceArn"},
			"recall-benefit-application": {"Catalog", "ClientToken", "Identifier", "Reason"},
			"submit-benefit-application": {"Catalog", "Identifier"},
			"tag-resource":               {"ResourceArn", "Tags"},
			"untag-resource":             {"ResourceArn", "TagKeys"},
			"update-benefit-application": {"BenefitApplicationDetails", "Catalog", "ClientToken", "Description", "FileDetails", "Identifier", "Name", "PartnerContacts", "Revision"},
		},
		OperationInputTypes: map[string]map[string]string{
			"amend-benefit-application":                 {"AmendmentReason": "*string", "Amendments": "[]types.Amendment", "Catalog": "*string", "ClientToken": "*string", "Identifier": "*string", "Revision": "*string"},
			"associate-benefit-application-resource":    {"BenefitApplicationIdentifier": "*string", "Catalog": "*string", "ResourceArn": "*string"},
			"cancel-benefit-application":                {"Catalog": "*string", "ClientToken": "*string", "Identifier": "*string", "Reason": "*string"},
			"create-benefit-application":                {"AssociatedResources": "[]string", "BenefitApplicationDetails": "document.Interface", "BenefitIdentifier": "*string", "Catalog": "*string", "ClientToken": "*string", "Description": "*string", "FileDetails": "[]types.FileInput", "FulfillmentTypes": "[]types.FulfillmentType", "Name": "*string", "PartnerContacts": "[]types.Contact", "Tags": "[]types.Tag"},
			"disassociate-benefit-application-resource": {"BenefitApplicationIdentifier": "*string", "Catalog": "*string", "ResourceArn": "*string"},
			"get-benefit":                {"Catalog": "*string", "Identifier": "*string"},
			"get-benefit-allocation":     {"Catalog": "*string", "Identifier": "*string"},
			"get-benefit-application":    {"Catalog": "*string", "Identifier": "*string"},
			"list-benefit-allocations":   {"BenefitApplicationIdentifiers": "[]string", "BenefitIdentifiers": "[]string", "Catalog": "*string", "FulfillmentTypes": "[]types.FulfillmentType", "MaxResults": "*int32", "NextToken": "*string", "Status": "[]types.BenefitAllocationStatus"},
			"list-benefit-applications":  {"AssociatedResourceArns": "[]string", "AssociatedResources": "[]types.AssociatedResource", "BenefitIdentifiers": "[]string", "Catalog": "*string", "FulfillmentTypes": "[]types.FulfillmentType", "MaxResults": "*int32", "NextToken": "*string", "Programs": "[]string", "Stages": "[]string", "Status": "[]types.BenefitApplicationStatus"},
			"list-benefits":              {"Catalog": "*string", "FulfillmentTypes": "[]types.FulfillmentType", "MaxResults": "*int32", "NextToken": "*string", "Programs": "[]string", "Status": "[]types.BenefitStatus"},
			"list-tags-for-resource":     {"ResourceArn": "*string"},
			"recall-benefit-application": {"Catalog": "*string", "ClientToken": "*string", "Identifier": "*string", "Reason": "*string"},
			"submit-benefit-application": {"Catalog": "*string", "Identifier": "*string"},
			"tag-resource":               {"ResourceArn": "*string", "Tags": "[]types.Tag"},
			"untag-resource":             {"ResourceArn": "*string", "TagKeys": "[]string"},
			"update-benefit-application": {"BenefitApplicationDetails": "document.Interface", "Catalog": "*string", "ClientToken": "*string", "Description": "*string", "FileDetails": "[]types.FileInput", "Identifier": "*string", "Name": "*string", "PartnerContacts": "[]types.Contact", "Revision": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"amend-benefit-application":                 {"AmendmentReason", "Amendments", "Catalog", "ClientToken", "Identifier", "Revision"},
			"associate-benefit-application-resource":    {"BenefitApplicationIdentifier", "Catalog", "ResourceArn"},
			"cancel-benefit-application":                {"Catalog", "ClientToken", "Identifier"},
			"create-benefit-application":                {"BenefitIdentifier", "Catalog", "ClientToken"},
			"disassociate-benefit-application-resource": {"BenefitApplicationIdentifier", "Catalog", "ResourceArn"},
			"get-benefit":                {"Catalog", "Identifier"},
			"get-benefit-allocation":     {"Catalog", "Identifier"},
			"get-benefit-application":    {"Catalog", "Identifier"},
			"list-benefit-allocations":   {"Catalog"},
			"list-benefit-applications":  {"Catalog"},
			"list-benefits":              {"Catalog"},
			"list-tags-for-resource":     {"ResourceArn"},
			"recall-benefit-application": {"Catalog", "Identifier", "Reason"},
			"submit-benefit-application": {"Catalog", "Identifier"},
			"tag-resource":               {"ResourceArn", "Tags"},
			"untag-resource":             {"ResourceArn", "TagKeys"},
			"update-benefit-application": {"Catalog", "ClientToken", "Identifier", "Revision"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("partnercentralbenefits", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
