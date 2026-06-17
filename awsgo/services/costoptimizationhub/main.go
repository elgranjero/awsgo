package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/costoptimizationhub/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"get-preferences", "get-recommendation", "list-efficiency-metrics", "list-enrollment-statuses", "list-recommendation-summaries", "list-recommendations", "update-enrollment-status", "update-preferences"},
		OperationSet: map[string]bool{"get-preferences": true, "get-recommendation": true, "list-efficiency-metrics": true, "list-enrollment-statuses": true, "list-recommendation-summaries": true, "list-recommendations": true, "update-enrollment-status": true, "update-preferences": true},
		OperationInputs: map[string][]string{
			"get-preferences":               {},
			"get-recommendation":            {"RecommendationId"},
			"list-efficiency-metrics":       {"Granularity", "GroupBy", "MaxResults", "NextToken", "OrderBy", "TimePeriod"},
			"list-enrollment-statuses":      {"AccountId", "IncludeOrganizationInfo", "MaxResults", "NextToken"},
			"list-recommendation-summaries": {"Filter", "GroupBy", "MaxResults", "Metrics", "NextToken"},
			"list-recommendations":          {"Filter", "IncludeAllRecommendations", "MaxResults", "NextToken", "OrderBy"},
			"update-enrollment-status":      {"IncludeMemberAccounts", "Status"},
			"update-preferences":            {"MemberAccountDiscountVisibility", "PreferredCommitment", "SavingsEstimationMode"},
		},
		OperationInputTypes: map[string]map[string]string{
			"get-preferences":               {},
			"get-recommendation":            {"RecommendationId": "*string"},
			"list-efficiency-metrics":       {"Granularity": "types.GranularityType", "GroupBy": "*string", "MaxResults": "*int32", "NextToken": "*string", "OrderBy": "*types.OrderBy", "TimePeriod": "*types.TimePeriod"},
			"list-enrollment-statuses":      {"AccountId": "*string", "IncludeOrganizationInfo": "bool", "MaxResults": "*int32", "NextToken": "*string"},
			"list-recommendation-summaries": {"Filter": "*types.Filter", "GroupBy": "*string", "MaxResults": "*int32", "Metrics": "[]types.SummaryMetrics", "NextToken": "*string"},
			"list-recommendations":          {"Filter": "*types.Filter", "IncludeAllRecommendations": "bool", "MaxResults": "*int32", "NextToken": "*string", "OrderBy": "*types.OrderBy"},
			"update-enrollment-status":      {"IncludeMemberAccounts": "*bool", "Status": "types.EnrollmentStatus"},
			"update-preferences":            {"MemberAccountDiscountVisibility": "types.MemberAccountDiscountVisibility", "PreferredCommitment": "*types.PreferredCommitment", "SavingsEstimationMode": "types.SavingsEstimationMode"},
		},
		OperationInputRequired: map[string][]string{
			"get-preferences":               {},
			"get-recommendation":            {"RecommendationId"},
			"list-efficiency-metrics":       {"Granularity", "TimePeriod"},
			"list-enrollment-statuses":      {},
			"list-recommendation-summaries": {"GroupBy"},
			"list-recommendations":          {},
			"update-enrollment-status":      {"Status"},
			"update-preferences":            {},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("costoptimizationhub", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
