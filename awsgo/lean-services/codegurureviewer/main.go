package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/codegurureviewer"
)

var fields_associate_repository = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "KMSKeyDetails", Flag: "kms-key-details", Type: "*types.KMSKeyDetails", Required: false},
	{Name: "Repository", Flag: "repository", Type: "*types.Repository", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_code_review = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RepositoryAssociationArn", Flag: "repository-association-arn", Type: "*string", Required: true},
	{Name: "Type", Flag: "type", Type: "*types.CodeReviewType", Required: true},
}

var fields_describe_code_review = []leanruntime.Field{
	{Name: "CodeReviewArn", Flag: "code-review-arn", Type: "*string", Required: true},
}

var fields_describe_recommendation_feedback = []leanruntime.Field{
	{Name: "CodeReviewArn", Flag: "code-review-arn", Type: "*string", Required: true},
	{Name: "RecommendationId", Flag: "recommendation-id", Type: "*string", Required: true},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: false},
}

var fields_describe_repository_association = []leanruntime.Field{
	{Name: "AssociationArn", Flag: "association-arn", Type: "*string", Required: true},
}

var fields_disassociate_repository = []leanruntime.Field{
	{Name: "AssociationArn", Flag: "association-arn", Type: "*string", Required: true},
}

var fields_list_code_reviews = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ProviderTypes", Flag: "provider-types", Type: "[]types.ProviderType", Required: false},
	{Name: "RepositoryNames", Flag: "repository-names", Type: "[]string", Required: false},
	{Name: "States", Flag: "states", Type: "[]types.JobState", Required: false},
	{Name: "Type", Flag: "type", Type: "types.Type", Required: true},
}

var fields_list_recommendation_feedback = []leanruntime.Field{
	{Name: "CodeReviewArn", Flag: "code-review-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RecommendationIds", Flag: "recommendation-ids", Type: "[]string", Required: false},
	{Name: "UserIds", Flag: "user-ids", Type: "[]string", Required: false},
}

var fields_list_recommendations = []leanruntime.Field{
	{Name: "CodeReviewArn", Flag: "code-review-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_repository_associations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Names", Flag: "names", Type: "[]string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Owners", Flag: "owners", Type: "[]string", Required: false},
	{Name: "ProviderTypes", Flag: "provider-types", Type: "[]types.ProviderType", Required: false},
	{Name: "States", Flag: "states", Type: "[]types.RepositoryAssociationState", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_put_recommendation_feedback = []leanruntime.Field{
	{Name: "CodeReviewArn", Flag: "code-review-arn", Type: "*string", Required: true},
	{Name: "Reactions", Flag: "reactions", Type: "[]types.Reaction", Required: true},
	{Name: "RecommendationId", Flag: "recommendation-id", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"associate-repository": {
			Name:   "associate-repository",
			Fields: fields_associate_repository,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateRepositoryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_repository, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateRepository(ctx, input)
			},
		},
		"create-code-review": {
			Name:   "create-code-review",
			Fields: fields_create_code_review,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCodeReviewInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_code_review, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCodeReview(ctx, input)
			},
		},
		"describe-code-review": {
			Name:   "describe-code-review",
			Fields: fields_describe_code_review,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCodeReviewInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_code_review, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeCodeReview(ctx, input)
			},
		},
		"describe-recommendation-feedback": {
			Name:   "describe-recommendation-feedback",
			Fields: fields_describe_recommendation_feedback,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeRecommendationFeedbackInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_recommendation_feedback, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeRecommendationFeedback(ctx, input)
			},
		},
		"describe-repository-association": {
			Name:   "describe-repository-association",
			Fields: fields_describe_repository_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeRepositoryAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_repository_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeRepositoryAssociation(ctx, input)
			},
		},
		"disassociate-repository": {
			Name:   "disassociate-repository",
			Fields: fields_disassociate_repository,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateRepositoryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_repository, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateRepository(ctx, input)
			},
		},
		"list-code-reviews": {
			Name:   "list-code-reviews",
			Fields: fields_list_code_reviews,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCodeReviewsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_code_reviews, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCodeReviews(ctx, input)
				}
				var results []*svc.ListCodeReviewsOutput
				p := svc.NewListCodeReviewsPaginator(client, input)
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
		"list-recommendation-feedback": {
			Name:   "list-recommendation-feedback",
			Fields: fields_list_recommendation_feedback,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRecommendationFeedbackInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_recommendation_feedback, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRecommendationFeedback(ctx, input)
				}
				var results []*svc.ListRecommendationFeedbackOutput
				p := svc.NewListRecommendationFeedbackPaginator(client, input)
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
		"list-recommendations": {
			Name:   "list-recommendations",
			Fields: fields_list_recommendations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRecommendationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_recommendations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRecommendations(ctx, input)
				}
				var results []*svc.ListRecommendationsOutput
				p := svc.NewListRecommendationsPaginator(client, input)
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
		"list-repository-associations": {
			Name:   "list-repository-associations",
			Fields: fields_list_repository_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRepositoryAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_repository_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRepositoryAssociations(ctx, input)
				}
				var results []*svc.ListRepositoryAssociationsOutput
				p := svc.NewListRepositoryAssociationsPaginator(client, input)
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
		"list-tags-for-resource": {
			Name:   "list-tags-for-resource",
			Fields: fields_list_tags_for_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTagsForResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_tags_for_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListTagsForResource(ctx, input)
			},
		},
		"put-recommendation-feedback": {
			Name:   "put-recommendation-feedback",
			Fields: fields_put_recommendation_feedback,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutRecommendationFeedbackInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_recommendation_feedback, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutRecommendationFeedback(ctx, input)
			},
		},
		"tag-resource": {
			Name:   "tag-resource",
			Fields: fields_tag_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TagResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_tag_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TagResource(ctx, input)
			},
		},
		"untag-resource": {
			Name:   "untag-resource",
			Fields: fields_untag_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UntagResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_untag_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UntagResource(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("codegurureviewer", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
