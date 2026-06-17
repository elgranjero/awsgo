package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/personalizeruntime/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"get-action-recommendations", "get-personalized-ranking", "get-recommendations"},
		OperationSet: map[string]bool{"get-action-recommendations": true, "get-personalized-ranking": true, "get-recommendations": true},
		OperationInputs: map[string][]string{
			"get-action-recommendations": {"CampaignArn", "FilterArn", "FilterValues", "NumResults", "UserId"},
			"get-personalized-ranking":   {"CampaignArn", "Context", "FilterArn", "FilterValues", "InputList", "MetadataColumns", "UserId"},
			"get-recommendations":        {"CampaignArn", "Context", "FilterArn", "FilterValues", "ItemId", "MetadataColumns", "NumResults", "Promotions", "RecommenderArn", "UserId"},
		},
		OperationInputTypes: map[string]map[string]string{
			"get-action-recommendations": {"CampaignArn": "*string", "FilterArn": "*string", "FilterValues": "map[string]string", "NumResults": "int32", "UserId": "*string"},
			"get-personalized-ranking":   {"CampaignArn": "*string", "Context": "map[string]string", "FilterArn": "*string", "FilterValues": "map[string]string", "InputList": "[]string", "MetadataColumns": "map[string][]string", "UserId": "*string"},
			"get-recommendations":        {"CampaignArn": "*string", "Context": "map[string]string", "FilterArn": "*string", "FilterValues": "map[string]string", "ItemId": "*string", "MetadataColumns": "map[string][]string", "NumResults": "int32", "Promotions": "[]types.Promotion", "RecommenderArn": "*string", "UserId": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"get-action-recommendations": {},
			"get-personalized-ranking":   {"CampaignArn", "InputList", "UserId"},
			"get-recommendations":        {},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("personalizeruntime", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
