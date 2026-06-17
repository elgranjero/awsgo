package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/pricing/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"describe-services", "get-attribute-values", "get-price-list-file-url", "get-products", "list-price-lists"},
		OperationSet: map[string]bool{"describe-services": true, "get-attribute-values": true, "get-price-list-file-url": true, "get-products": true, "list-price-lists": true},
		OperationInputs: map[string][]string{
			"describe-services":       {"FormatVersion", "MaxResults", "NextToken", "ServiceCode"},
			"get-attribute-values":    {"AttributeName", "MaxResults", "NextToken", "ServiceCode"},
			"get-price-list-file-url": {"FileFormat", "PriceListArn"},
			"get-products":            {"Filters", "FormatVersion", "MaxResults", "NextToken", "ServiceCode"},
			"list-price-lists":        {"CurrencyCode", "EffectiveDate", "MaxResults", "NextToken", "RegionCode", "ServiceCode"},
		},
		OperationInputTypes: map[string]map[string]string{
			"describe-services":       {"FormatVersion": "*string", "MaxResults": "*int32", "NextToken": "*string", "ServiceCode": "*string"},
			"get-attribute-values":    {"AttributeName": "*string", "MaxResults": "*int32", "NextToken": "*string", "ServiceCode": "*string"},
			"get-price-list-file-url": {"FileFormat": "*string", "PriceListArn": "*string"},
			"get-products":            {"Filters": "[]types.Filter", "FormatVersion": "*string", "MaxResults": "*int32", "NextToken": "*string", "ServiceCode": "*string"},
			"list-price-lists":        {"CurrencyCode": "*string", "EffectiveDate": "*time.Time", "MaxResults": "*int32", "NextToken": "*string", "RegionCode": "*string", "ServiceCode": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"describe-services":       {},
			"get-attribute-values":    {"AttributeName", "ServiceCode"},
			"get-price-list-file-url": {"FileFormat", "PriceListArn"},
			"get-products":            {"ServiceCode"},
			"list-price-lists":        {"CurrencyCode", "EffectiveDate", "ServiceCode"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("pricing", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
