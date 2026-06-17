package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/codegurureviewer/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"associate-repository", "create-code-review", "describe-code-review", "describe-recommendation-feedback", "describe-repository-association", "disassociate-repository", "list-code-reviews", "list-recommendation-feedback", "list-recommendations", "list-repository-associations", "list-tags-for-resource", "put-recommendation-feedback", "tag-resource", "untag-resource"},
		OperationSet: map[string]bool{"associate-repository": true, "create-code-review": true, "describe-code-review": true, "describe-recommendation-feedback": true, "describe-repository-association": true, "disassociate-repository": true, "list-code-reviews": true, "list-recommendation-feedback": true, "list-recommendations": true, "list-repository-associations": true, "list-tags-for-resource": true, "put-recommendation-feedback": true, "tag-resource": true, "untag-resource": true},
		OperationInputs: map[string][]string{
			"associate-repository":             {"ClientRequestToken", "KMSKeyDetails", "Repository", "Tags"},
			"create-code-review":               {"ClientRequestToken", "Name", "RepositoryAssociationArn", "Type"},
			"describe-code-review":             {"CodeReviewArn"},
			"describe-recommendation-feedback": {"CodeReviewArn", "RecommendationId", "UserId"},
			"describe-repository-association":  {"AssociationArn"},
			"disassociate-repository":          {"AssociationArn"},
			"list-code-reviews":                {"MaxResults", "NextToken", "ProviderTypes", "RepositoryNames", "States", "Type"},
			"list-recommendation-feedback":     {"CodeReviewArn", "MaxResults", "NextToken", "RecommendationIds", "UserIds"},
			"list-recommendations":             {"CodeReviewArn", "MaxResults", "NextToken"},
			"list-repository-associations":     {"MaxResults", "Names", "NextToken", "Owners", "ProviderTypes", "States"},
			"list-tags-for-resource":           {"ResourceArn"},
			"put-recommendation-feedback":      {"CodeReviewArn", "Reactions", "RecommendationId"},
			"tag-resource":                     {"ResourceArn", "Tags"},
			"untag-resource":                   {"ResourceArn", "TagKeys"},
		},
		OperationInputTypes: map[string]map[string]string{
			"associate-repository":             {"ClientRequestToken": "*string", "KMSKeyDetails": "*types.KMSKeyDetails", "Repository": "*types.Repository", "Tags": "map[string]string"},
			"create-code-review":               {"ClientRequestToken": "*string", "Name": "*string", "RepositoryAssociationArn": "*string", "Type": "*types.CodeReviewType"},
			"describe-code-review":             {"CodeReviewArn": "*string"},
			"describe-recommendation-feedback": {"CodeReviewArn": "*string", "RecommendationId": "*string", "UserId": "*string"},
			"describe-repository-association":  {"AssociationArn": "*string"},
			"disassociate-repository":          {"AssociationArn": "*string"},
			"list-code-reviews":                {"MaxResults": "*int32", "NextToken": "*string", "ProviderTypes": "[]types.ProviderType", "RepositoryNames": "[]string", "States": "[]types.JobState", "Type": "types.Type"},
			"list-recommendation-feedback":     {"CodeReviewArn": "*string", "MaxResults": "*int32", "NextToken": "*string", "RecommendationIds": "[]string", "UserIds": "[]string"},
			"list-recommendations":             {"CodeReviewArn": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"list-repository-associations":     {"MaxResults": "*int32", "Names": "[]string", "NextToken": "*string", "Owners": "[]string", "ProviderTypes": "[]types.ProviderType", "States": "[]types.RepositoryAssociationState"},
			"list-tags-for-resource":           {"ResourceArn": "*string"},
			"put-recommendation-feedback":      {"CodeReviewArn": "*string", "Reactions": "[]types.Reaction", "RecommendationId": "*string"},
			"tag-resource":                     {"ResourceArn": "*string", "Tags": "map[string]string"},
			"untag-resource":                   {"ResourceArn": "*string", "TagKeys": "[]string"},
		},
		OperationInputRequired: map[string][]string{
			"associate-repository":             {"Repository"},
			"create-code-review":               {"Name", "RepositoryAssociationArn", "Type"},
			"describe-code-review":             {"CodeReviewArn"},
			"describe-recommendation-feedback": {"CodeReviewArn", "RecommendationId"},
			"describe-repository-association":  {"AssociationArn"},
			"disassociate-repository":          {"AssociationArn"},
			"list-code-reviews":                {"Type"},
			"list-recommendation-feedback":     {"CodeReviewArn"},
			"list-recommendations":             {"CodeReviewArn"},
			"list-repository-associations":     {},
			"list-tags-for-resource":           {"ResourceArn"},
			"put-recommendation-feedback":      {"CodeReviewArn", "Reactions", "RecommendationId"},
			"tag-resource":                     {"ResourceArn", "Tags"},
			"untag-resource":                   {"ResourceArn", "TagKeys"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("codegurureviewer", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
