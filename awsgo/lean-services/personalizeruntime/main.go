package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/personalizeruntime"
)

var fields_get_action_recommendations = []leanruntime.Field{
	{Name: "CampaignArn", Flag: "campaign-arn", Type: "*string", Required: false},
	{Name: "FilterArn", Flag: "filter-arn", Type: "*string", Required: false},
	{Name: "FilterValues", Flag: "filter-values", Type: "map[string]string", Required: false},
	{Name: "NumResults", Flag: "num-results", Type: "int32", Required: false},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: false},
}

var fields_get_personalized_ranking = []leanruntime.Field{
	{Name: "CampaignArn", Flag: "campaign-arn", Type: "*string", Required: true},
	{Name: "Context", Flag: "context", Type: "map[string]string", Required: false},
	{Name: "FilterArn", Flag: "filter-arn", Type: "*string", Required: false},
	{Name: "FilterValues", Flag: "filter-values", Type: "map[string]string", Required: false},
	{Name: "InputList", Flag: "input-list", Type: "[]string", Required: true},
	{Name: "MetadataColumns", Flag: "metadata-columns", Type: "map[string][]string", Required: false},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_get_recommendations = []leanruntime.Field{
	{Name: "CampaignArn", Flag: "campaign-arn", Type: "*string", Required: false},
	{Name: "Context", Flag: "context", Type: "map[string]string", Required: false},
	{Name: "FilterArn", Flag: "filter-arn", Type: "*string", Required: false},
	{Name: "FilterValues", Flag: "filter-values", Type: "map[string]string", Required: false},
	{Name: "ItemId", Flag: "item-id", Type: "*string", Required: false},
	{Name: "MetadataColumns", Flag: "metadata-columns", Type: "map[string][]string", Required: false},
	{Name: "NumResults", Flag: "num-results", Type: "int32", Required: false},
	{Name: "Promotions", Flag: "promotions", Type: "[]types.Promotion", Required: false},
	{Name: "RecommenderArn", Flag: "recommender-arn", Type: "*string", Required: false},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"get-action-recommendations": {
			Name:   "get-action-recommendations",
			Fields: fields_get_action_recommendations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetActionRecommendationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_action_recommendations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetActionRecommendations(ctx, input)
			},
		},
		"get-personalized-ranking": {
			Name:   "get-personalized-ranking",
			Fields: fields_get_personalized_ranking,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPersonalizedRankingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_personalized_ranking, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPersonalizedRanking(ctx, input)
			},
		},
		"get-recommendations": {
			Name:   "get-recommendations",
			Fields: fields_get_recommendations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRecommendationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_recommendations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRecommendations(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("personalizeruntime", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
