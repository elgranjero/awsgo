package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/trustedadvisor/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"batch-update-recommendation-resource-exclusion", "get-organization-recommendation", "get-recommendation", "list-checks", "list-organization-recommendation-accounts", "list-organization-recommendation-resources", "list-organization-recommendations", "list-recommendation-resources", "list-recommendations", "update-organization-recommendation-lifecycle", "update-recommendation-lifecycle"},
		OperationSet: map[string]bool{"batch-update-recommendation-resource-exclusion": true, "get-organization-recommendation": true, "get-recommendation": true, "list-checks": true, "list-organization-recommendation-accounts": true, "list-organization-recommendation-resources": true, "list-organization-recommendations": true, "list-recommendation-resources": true, "list-recommendations": true, "update-organization-recommendation-lifecycle": true, "update-recommendation-lifecycle": true},
		OperationInputs: map[string][]string{
			"batch-update-recommendation-resource-exclusion": {"RecommendationResourceExclusions"},
			"get-organization-recommendation":                {"OrganizationRecommendationIdentifier"},
			"get-recommendation":                             {"Language", "RecommendationIdentifier"},
			"list-checks":                                    {"AwsService", "Language", "MaxResults", "NextToken", "Pillar", "Source"},
			"list-organization-recommendation-accounts":      {"AffectedAccountId", "MaxResults", "NextToken", "OrganizationRecommendationIdentifier"},
			"list-organization-recommendation-resources":     {"AffectedAccountId", "ExclusionStatus", "MaxResults", "NextToken", "OrganizationRecommendationIdentifier", "RegionCode", "Status"},
			"list-organization-recommendations":              {"AfterLastUpdatedAt", "AwsService", "BeforeLastUpdatedAt", "CheckIdentifier", "MaxResults", "NextToken", "Pillar", "Source", "Status", "Type"},
			"list-recommendation-resources":                  {"ExclusionStatus", "Language", "MaxResults", "NextToken", "RecommendationIdentifier", "RegionCode", "Status"},
			"list-recommendations":                           {"AfterLastUpdatedAt", "AwsService", "BeforeLastUpdatedAt", "CheckIdentifier", "Language", "MaxResults", "NextToken", "Pillar", "Source", "Status", "Type"},
			"update-organization-recommendation-lifecycle":   {"LifecycleStage", "OrganizationRecommendationIdentifier", "UpdateReason", "UpdateReasonCode"},
			"update-recommendation-lifecycle":                {"LifecycleStage", "RecommendationIdentifier", "UpdateReason", "UpdateReasonCode"},
		},
		OperationInputTypes: map[string]map[string]string{
			"batch-update-recommendation-resource-exclusion": {"RecommendationResourceExclusions": "[]types.RecommendationResourceExclusion"},
			"get-organization-recommendation":                {"OrganizationRecommendationIdentifier": "*string"},
			"get-recommendation":                             {"Language": "types.RecommendationLanguage", "RecommendationIdentifier": "*string"},
			"list-checks":                                    {"AwsService": "*string", "Language": "types.RecommendationLanguage", "MaxResults": "*int32", "NextToken": "*string", "Pillar": "types.RecommendationPillar", "Source": "types.RecommendationSource"},
			"list-organization-recommendation-accounts":      {"AffectedAccountId": "*string", "MaxResults": "*int32", "NextToken": "*string", "OrganizationRecommendationIdentifier": "*string"},
			"list-organization-recommendation-resources":     {"AffectedAccountId": "*string", "ExclusionStatus": "types.ExclusionStatus", "MaxResults": "*int32", "NextToken": "*string", "OrganizationRecommendationIdentifier": "*string", "RegionCode": "*string", "Status": "types.ResourceStatus"},
			"list-organization-recommendations":              {"AfterLastUpdatedAt": "*time.Time", "AwsService": "*string", "BeforeLastUpdatedAt": "*time.Time", "CheckIdentifier": "*string", "MaxResults": "*int32", "NextToken": "*string", "Pillar": "types.RecommendationPillar", "Source": "types.RecommendationSource", "Status": "types.RecommendationStatus", "Type": "types.RecommendationType"},
			"list-recommendation-resources":                  {"ExclusionStatus": "types.ExclusionStatus", "Language": "types.RecommendationLanguage", "MaxResults": "*int32", "NextToken": "*string", "RecommendationIdentifier": "*string", "RegionCode": "*string", "Status": "types.ResourceStatus"},
			"list-recommendations":                           {"AfterLastUpdatedAt": "*time.Time", "AwsService": "*string", "BeforeLastUpdatedAt": "*time.Time", "CheckIdentifier": "*string", "Language": "types.RecommendationLanguage", "MaxResults": "*int32", "NextToken": "*string", "Pillar": "types.RecommendationPillar", "Source": "types.RecommendationSource", "Status": "types.RecommendationStatus", "Type": "types.RecommendationType"},
			"update-organization-recommendation-lifecycle":   {"LifecycleStage": "types.UpdateRecommendationLifecycleStage", "OrganizationRecommendationIdentifier": "*string", "UpdateReason": "*string", "UpdateReasonCode": "types.UpdateRecommendationLifecycleStageReasonCode"},
			"update-recommendation-lifecycle":                {"LifecycleStage": "types.UpdateRecommendationLifecycleStage", "RecommendationIdentifier": "*string", "UpdateReason": "*string", "UpdateReasonCode": "types.UpdateRecommendationLifecycleStageReasonCode"},
		},
		OperationInputRequired: map[string][]string{
			"batch-update-recommendation-resource-exclusion": {"RecommendationResourceExclusions"},
			"get-organization-recommendation":                {"OrganizationRecommendationIdentifier"},
			"get-recommendation":                             {"RecommendationIdentifier"},
			"list-checks":                                    {},
			"list-organization-recommendation-accounts":      {"OrganizationRecommendationIdentifier"},
			"list-organization-recommendation-resources":     {"OrganizationRecommendationIdentifier"},
			"list-organization-recommendations":              {},
			"list-recommendation-resources":                  {"RecommendationIdentifier"},
			"list-recommendations":                           {},
			"update-organization-recommendation-lifecycle":   {"LifecycleStage", "OrganizationRecommendationIdentifier"},
			"update-recommendation-lifecycle":                {"LifecycleStage", "RecommendationIdentifier"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("trustedadvisor", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
