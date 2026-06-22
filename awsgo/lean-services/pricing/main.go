package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/pricing"
)

var fields_describe_services = []leanruntime.Field{
	{Name: "FormatVersion", Flag: "format-version", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ServiceCode", Flag: "service-code", Type: "*string", Required: false},
}

var fields_get_attribute_values = []leanruntime.Field{
	{Name: "AttributeName", Flag: "attribute-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ServiceCode", Flag: "service-code", Type: "*string", Required: true},
}

var fields_get_price_list_file_url = []leanruntime.Field{
	{Name: "FileFormat", Flag: "file-format", Type: "*string", Required: true},
	{Name: "PriceListArn", Flag: "price-list-arn", Type: "*string", Required: true},
}

var fields_get_products = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "FormatVersion", Flag: "format-version", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ServiceCode", Flag: "service-code", Type: "*string", Required: true},
}

var fields_list_price_lists = []leanruntime.Field{
	{Name: "CurrencyCode", Flag: "currency-code", Type: "*string", Required: true},
	{Name: "EffectiveDate", Flag: "effective-date", Type: "*time.Time", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RegionCode", Flag: "region-code", Type: "*string", Required: false},
	{Name: "ServiceCode", Flag: "service-code", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"describe-services": {
			Name:   "describe-services",
			Fields: fields_describe_services,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeServicesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_services, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeServices(ctx, input)
				}
				var results []*svc.DescribeServicesOutput
				p := svc.NewDescribeServicesPaginator(client, input)
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
		"get-attribute-values": {
			Name:   "get-attribute-values",
			Fields: fields_get_attribute_values,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAttributeValuesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_attribute_values, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetAttributeValues(ctx, input)
				}
				var results []*svc.GetAttributeValuesOutput
				p := svc.NewGetAttributeValuesPaginator(client, input)
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
		"get-price-list-file-url": {
			Name:   "get-price-list-file-url",
			Fields: fields_get_price_list_file_url,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPriceListFileUrlInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_price_list_file_url, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPriceListFileUrl(ctx, input)
			},
		},
		"get-products": {
			Name:   "get-products",
			Fields: fields_get_products,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetProductsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_products, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetProducts(ctx, input)
				}
				var results []*svc.GetProductsOutput
				p := svc.NewGetProductsPaginator(client, input)
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
		"list-price-lists": {
			Name:   "list-price-lists",
			Fields: fields_list_price_lists,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPriceListsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_price_lists, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPriceLists(ctx, input)
				}
				var results []*svc.ListPriceListsOutput
				p := svc.NewListPriceListsPaginator(client, input)
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
	}
	if err := leanruntime.Execute("pricing", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
