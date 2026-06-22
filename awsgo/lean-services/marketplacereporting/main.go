package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/marketplacereporting"
)

var fields_get_buyer_dashboard = []leanruntime.Field{
	{Name: "DashboardIdentifier", Flag: "dashboard-identifier", Type: "*string", Required: true},
	{Name: "EmbeddingDomains", Flag: "embedding-domains", Type: "[]string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"get-buyer-dashboard": {
			Name:   "get-buyer-dashboard",
			Fields: fields_get_buyer_dashboard,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBuyerDashboardInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_buyer_dashboard, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBuyerDashboard(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("marketplacereporting", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
