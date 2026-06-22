package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/marketplaceentitlementservice"
)

var fields_get_entitlements = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "map[string][]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ProductCode", Flag: "product-code", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"get-entitlements": {
			Name:   "get-entitlements",
			Fields: fields_get_entitlements,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEntitlementsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_entitlements, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetEntitlements(ctx, input)
				}
				var results []*svc.GetEntitlementsOutput
				p := svc.NewGetEntitlementsPaginator(client, input)
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
	if err := leanruntime.Execute("marketplaceentitlementservice", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
