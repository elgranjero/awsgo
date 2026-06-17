package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/savingsplans/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"create-savings-plan", "delete-queued-savings-plan", "describe-savings-plan-rates", "describe-savings-plans", "describe-savings-plans-offering-rates", "describe-savings-plans-offerings", "list-tags-for-resource", "return-savings-plan", "tag-resource", "untag-resource"},
		OperationSet: map[string]bool{"create-savings-plan": true, "delete-queued-savings-plan": true, "describe-savings-plan-rates": true, "describe-savings-plans": true, "describe-savings-plans-offering-rates": true, "describe-savings-plans-offerings": true, "list-tags-for-resource": true, "return-savings-plan": true, "tag-resource": true, "untag-resource": true},
		OperationInputs: map[string][]string{
			"create-savings-plan":                   {"ClientToken", "Commitment", "PurchaseTime", "SavingsPlanOfferingId", "Tags", "UpfrontPaymentAmount"},
			"delete-queued-savings-plan":            {"SavingsPlanId"},
			"describe-savings-plan-rates":           {"Filters", "MaxResults", "NextToken", "SavingsPlanId"},
			"describe-savings-plans":                {"Filters", "MaxResults", "NextToken", "SavingsPlanArns", "SavingsPlanIds", "States"},
			"describe-savings-plans-offering-rates": {"Filters", "MaxResults", "NextToken", "Operations", "Products", "SavingsPlanOfferingIds", "SavingsPlanPaymentOptions", "SavingsPlanTypes", "ServiceCodes", "UsageTypes"},
			"describe-savings-plans-offerings":      {"Currencies", "Descriptions", "Durations", "Filters", "MaxResults", "NextToken", "OfferingIds", "Operations", "PaymentOptions", "PlanTypes", "ProductType", "ServiceCodes", "UsageTypes"},
			"list-tags-for-resource":                {"ResourceArn"},
			"return-savings-plan":                   {"ClientToken", "SavingsPlanId"},
			"tag-resource":                          {"ResourceArn", "Tags"},
			"untag-resource":                        {"ResourceArn", "TagKeys"},
		},
		OperationInputTypes: map[string]map[string]string{
			"create-savings-plan":                   {"ClientToken": "*string", "Commitment": "*string", "PurchaseTime": "*time.Time", "SavingsPlanOfferingId": "*string", "Tags": "map[string]string", "UpfrontPaymentAmount": "*string"},
			"delete-queued-savings-plan":            {"SavingsPlanId": "*string"},
			"describe-savings-plan-rates":           {"Filters": "[]types.SavingsPlanRateFilter", "MaxResults": "*int32", "NextToken": "*string", "SavingsPlanId": "*string"},
			"describe-savings-plans":                {"Filters": "[]types.SavingsPlanFilter", "MaxResults": "*int32", "NextToken": "*string", "SavingsPlanArns": "[]string", "SavingsPlanIds": "[]string", "States": "[]types.SavingsPlanState"},
			"describe-savings-plans-offering-rates": {"Filters": "[]types.SavingsPlanOfferingRateFilterElement", "MaxResults": "int32", "NextToken": "*string", "Operations": "[]string", "Products": "[]types.SavingsPlanProductType", "SavingsPlanOfferingIds": "[]string", "SavingsPlanPaymentOptions": "[]types.SavingsPlanPaymentOption", "SavingsPlanTypes": "[]types.SavingsPlanType", "ServiceCodes": "[]types.SavingsPlanRateServiceCode", "UsageTypes": "[]string"},
			"describe-savings-plans-offerings":      {"Currencies": "[]types.CurrencyCode", "Descriptions": "[]string", "Durations": "[]int64", "Filters": "[]types.SavingsPlanOfferingFilterElement", "MaxResults": "int32", "NextToken": "*string", "OfferingIds": "[]string", "Operations": "[]string", "PaymentOptions": "[]types.SavingsPlanPaymentOption", "PlanTypes": "[]types.SavingsPlanType", "ProductType": "types.SavingsPlanProductType", "ServiceCodes": "[]string", "UsageTypes": "[]string"},
			"list-tags-for-resource":                {"ResourceArn": "*string"},
			"return-savings-plan":                   {"ClientToken": "*string", "SavingsPlanId": "*string"},
			"tag-resource":                          {"ResourceArn": "*string", "Tags": "map[string]string"},
			"untag-resource":                        {"ResourceArn": "*string", "TagKeys": "[]string"},
		},
		OperationInputRequired: map[string][]string{
			"create-savings-plan":                   {"Commitment", "SavingsPlanOfferingId"},
			"delete-queued-savings-plan":            {"SavingsPlanId"},
			"describe-savings-plan-rates":           {"SavingsPlanId"},
			"describe-savings-plans":                {},
			"describe-savings-plans-offering-rates": {},
			"describe-savings-plans-offerings":      {},
			"list-tags-for-resource":                {"ResourceArn"},
			"return-savings-plan":                   {"SavingsPlanId"},
			"tag-resource":                          {"ResourceArn", "Tags"},
			"untag-resource":                        {"ResourceArn", "TagKeys"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("savingsplans", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
