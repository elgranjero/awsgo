package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/marketplaceagreement"
)

var fields_describe_agreement = []leanruntime.Field{
	{Name: "AgreementId", Flag: "agreement-id", Type: "*string", Required: true},
}

var fields_get_agreement_terms = []leanruntime.Field{
	{Name: "AgreementId", Flag: "agreement-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_search_agreements = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Sort", Flag: "sort", Type: "*types.Sort", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"describe-agreement": {
			Name:   "describe-agreement",
			Fields: fields_describe_agreement,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAgreementInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_agreement, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAgreement(ctx, input)
			},
		},
		"get-agreement-terms": {
			Name:   "get-agreement-terms",
			Fields: fields_get_agreement_terms,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAgreementTermsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_agreement_terms, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetAgreementTerms(ctx, input)
				}
				var results []*svc.GetAgreementTermsOutput
				p := svc.NewGetAgreementTermsPaginator(client, input)
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
		"search-agreements": {
			Name:   "search-agreements",
			Fields: fields_search_agreements,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchAgreementsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_agreements, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchAgreements(ctx, input)
				}
				var results []*svc.SearchAgreementsOutput
				p := svc.NewSearchAgreementsPaginator(client, input)
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
	if err := leanruntime.Execute("marketplaceagreement", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
