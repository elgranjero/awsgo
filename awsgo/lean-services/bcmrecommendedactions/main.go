package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/bcmrecommendedactions"
)

var fields_list_recommended_actions = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.RequestFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"list-recommended-actions": {
			Name:   "list-recommended-actions",
			Fields: fields_list_recommended_actions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRecommendedActionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_recommended_actions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRecommendedActions(ctx, input)
				}
				var results []*svc.ListRecommendedActionsOutput
				p := svc.NewListRecommendedActionsPaginator(client, input)
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
	if err := leanruntime.Execute("bcmrecommendedactions", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
