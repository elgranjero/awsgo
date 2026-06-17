package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codegurureviewer"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// codegurureviewerCmd represents the codegurureviewer command
var _codegurureviewerCmd = &cobra.Command{
	Use:   "codegurureviewer",
	Short: "AWS codegurureviewer CLI",
	Run: func(cmd *cobra.Command, args []string) {
		_awsOutput = resolveAWSOutput(_awsProfile, cmd.Flags().Changed("output"))
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := codegurureviewer.NewFromConfig(cfg)
		if _codegurureviewerAssociateRepository {
			codegurureviewer_AssociateRepository(cfg, client)
			return
		}
		if _codegurureviewerCreateCodeReview {
			codegurureviewer_CreateCodeReview(cfg, client)
			return
		}
		if _codegurureviewerDescribeCodeReview {
			codegurureviewer_DescribeCodeReview(cfg, client)
			return
		}
		if _codegurureviewerDescribeRecommendationFeedback {
			codegurureviewer_DescribeRecommendationFeedback(cfg, client)
			return
		}
		if _codegurureviewerDescribeRepositoryAssociation {
			codegurureviewer_DescribeRepositoryAssociation(cfg, client)
			return
		}
		if _codegurureviewerDisassociateRepository {
			codegurureviewer_DisassociateRepository(cfg, client)
			return
		}
		if _codegurureviewerListCodeReviews {
			codegurureviewer_ListCodeReviews(cfg, client)
			return
		}
		if _codegurureviewerListRecommendationFeedback {
			codegurureviewer_ListRecommendationFeedback(cfg, client)
			return
		}
		if _codegurureviewerListRecommendations {
			codegurureviewer_ListRecommendations(cfg, client)
			return
		}
		if _codegurureviewerListRepositoryAssociations {
			codegurureviewer_ListRepositoryAssociations(cfg, client)
			return
		}
		if _codegurureviewerListTagsForResource {
			codegurureviewer_ListTagsForResource(cfg, client)
			return
		}
		if _codegurureviewerPutRecommendationFeedback {
			codegurureviewer_PutRecommendationFeedback(cfg, client)
			return
		}
		if _codegurureviewerTagResource {
			codegurureviewer_TagResource(cfg, client)
			return
		}
		if _codegurureviewerUntagResource {
			codegurureviewer_UntagResource(cfg, client)
			return
		}

	},
}

var (
	_codegurureviewerAssociateRepository            bool
	_codegurureviewerCreateCodeReview               bool
	_codegurureviewerDescribeCodeReview             bool
	_codegurureviewerDescribeRecommendationFeedback bool
	_codegurureviewerDescribeRepositoryAssociation  bool
	_codegurureviewerDisassociateRepository         bool
	_codegurureviewerListCodeReviews                bool
	_codegurureviewerListRecommendationFeedback     bool
	_codegurureviewerListRecommendations            bool
	_codegurureviewerListRepositoryAssociations     bool
	_codegurureviewerListTagsForResource            bool
	_codegurureviewerPutRecommendationFeedback      bool
	_codegurureviewerTagResource                    bool
	_codegurureviewerUntagResource                  bool

	_codegurureviewerAssociationArn           string
	_codegurureviewerClientRequestToken       string
	_codegurureviewerCodeReviewArn            string
	_codegurureviewerKMSKeyDetails            string
	_codegurureviewerMaxResults               string
	_codegurureviewerName                     string
	_codegurureviewerNames                    []string
	_codegurureviewerNextToken                string
	_codegurureviewerOwners                   []string
	_codegurureviewerProviderTypes            string
	_codegurureviewerReactions                string
	_codegurureviewerRecommendationId         string
	_codegurureviewerRecommendationIds        []string
	_codegurureviewerRepository               string
	_codegurureviewerRepositoryAssociationArn string
	_codegurureviewerRepositoryNames          []string
	_codegurureviewerResourceArn              string
	_codegurureviewerStates                   string
	_codegurureviewerTagKeys                  []string
	_codegurureviewerTags                     string
	_codegurureviewerType                     string
	_codegurureviewerUserId                   string
	_codegurureviewerUserIds                  []string
)

// Use to associate an Amazon Web Services CodeCommit repository or a repository
// managed by Amazon Web Services CodeStar Connections with Amazon CodeGuru
// Reviewer. When you associate a repository, CodeGuru Reviewer reviews source code
// changes in the repository's pull requests and provides automatic
// recommendations. You can view recommendations using the CodeGuru Reviewer
// console. For more information, see [Recommendations in Amazon CodeGuru Reviewer]in the Amazon CodeGuru Reviewer User Guide.
//
// If you associate a CodeCommit or S3 repository, it must be in the same Amazon
// Web Services Region and Amazon Web Services account where its CodeGuru Reviewer
// code reviews are configured.
//
// Bitbucket and GitHub Enterprise Server repositories are managed by Amazon Web
// Services CodeStar Connections to connect to CodeGuru Reviewer. For more
// information, see [Associate a repository]in the Amazon CodeGuru Reviewer User Guide.
//
// You cannot use the CodeGuru Reviewer SDK or the Amazon Web Services CLI to
// associate a GitHub repository with Amazon CodeGuru Reviewer. To associate a
// GitHub repository, use the console. For more information, see [Getting started with CodeGuru Reviewer]in the CodeGuru
// Reviewer User Guide.
//
// [Recommendations in Amazon CodeGuru Reviewer]: https://docs.aws.amazon.com/codeguru/latest/reviewer-ug/recommendations.html
// [Getting started with CodeGuru Reviewer]: https://docs.aws.amazon.com/codeguru/latest/reviewer-ug/getting-started-with-guru.html
// [Associate a repository]: https://docs.aws.amazon.com/codeguru/latest/reviewer-ug/getting-started-associate-repository.html
func codegurureviewer_AssociateRepository(cfg aws.Config, client *codegurureviewer.Client) {
	input := &codegurureviewer.AssociateRepositoryInput{
		// Repository: *types.Repository, // Required
	}

	if len(_codegurureviewerRepository) > 0 {
		if err := assignInputField(input, "Repository", _codegurureviewerRepository); err != nil {
			log.Errorf("invalid --repository: %s", err.Error())
			return
		}
	}
	if len(_codegurureviewerClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_codegurureviewerClientRequestToken)
	}
	if len(_codegurureviewerKMSKeyDetails) > 0 {
		if err := assignInputField(input, "KMSKeyDetails", _codegurureviewerKMSKeyDetails); err != nil {
			log.Errorf("invalid --kms-key-details: %s", err.Error())
			return
		}
	}
	if len(_codegurureviewerTags) > 0 {
		if err := assignInputField(input, "Tags", _codegurureviewerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.AssociateRepository(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Use to create a code review with a [CodeReviewType] of RepositoryAnalysis . This type of code
// review analyzes all code under a specified branch in an associated repository.
// PullRequest code reviews are automatically triggered by a pull request.
//
// [CodeReviewType]: https://docs.aws.amazon.com/codeguru/latest/reviewer-api/API_CodeReviewType.html
func codegurureviewer_CreateCodeReview(cfg aws.Config, client *codegurureviewer.Client) {
	input := &codegurureviewer.CreateCodeReviewInput{
		// Name: *string, // Required
		// RepositoryAssociationArn: *string, // Required
		// Type: *types.CodeReviewType, // Required
	}

	if len(_codegurureviewerName) > 0 {
		input.Name = aws.String(_codegurureviewerName)
	}
	if len(_codegurureviewerRepositoryAssociationArn) > 0 {
		input.RepositoryAssociationArn = aws.String(_codegurureviewerRepositoryAssociationArn)
	}
	if len(_codegurureviewerType) > 0 {
		if err := assignInputField(input, "Type", _codegurureviewerType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_codegurureviewerClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_codegurureviewerClientRequestToken)
	}

	if resp, err := client.CreateCodeReview(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the metadata associated with the code review along with its status.
func codegurureviewer_DescribeCodeReview(cfg aws.Config, client *codegurureviewer.Client) {
	input := &codegurureviewer.DescribeCodeReviewInput{
		// CodeReviewArn: *string, // Required
	}

	if len(_codegurureviewerCodeReviewArn) > 0 {
		input.CodeReviewArn = aws.String(_codegurureviewerCodeReviewArn)
	}

	if resp, err := client.DescribeCodeReview(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the customer feedback for a CodeGuru Reviewer recommendation.
func codegurureviewer_DescribeRecommendationFeedback(cfg aws.Config, client *codegurureviewer.Client) {
	input := &codegurureviewer.DescribeRecommendationFeedbackInput{
		// CodeReviewArn: *string, // Required
		// RecommendationId: *string, // Required
	}

	if len(_codegurureviewerCodeReviewArn) > 0 {
		input.CodeReviewArn = aws.String(_codegurureviewerCodeReviewArn)
	}
	if len(_codegurureviewerRecommendationId) > 0 {
		input.RecommendationId = aws.String(_codegurureviewerRecommendationId)
	}
	if len(_codegurureviewerUserId) > 0 {
		input.UserId = aws.String(_codegurureviewerUserId)
	}

	if resp, err := client.DescribeRecommendationFeedback(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a [RepositoryAssociation] object that contains information about the requested repository
// association.
//
// [RepositoryAssociation]: https://docs.aws.amazon.com/codeguru/latest/reviewer-api/API_RepositoryAssociation.html
func codegurureviewer_DescribeRepositoryAssociation(cfg aws.Config, client *codegurureviewer.Client) {
	input := &codegurureviewer.DescribeRepositoryAssociationInput{
		// AssociationArn: *string, // Required
	}

	if len(_codegurureviewerAssociationArn) > 0 {
		input.AssociationArn = aws.String(_codegurureviewerAssociationArn)
	}

	if resp, err := client.DescribeRepositoryAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the association between Amazon CodeGuru Reviewer and a repository.
func codegurureviewer_DisassociateRepository(cfg aws.Config, client *codegurureviewer.Client) {
	input := &codegurureviewer.DisassociateRepositoryInput{
		// AssociationArn: *string, // Required
	}

	if len(_codegurureviewerAssociationArn) > 0 {
		input.AssociationArn = aws.String(_codegurureviewerAssociationArn)
	}

	if resp, err := client.DisassociateRepository(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all the code reviews that the customer has created in the past 90 days.
func codegurureviewer_ListCodeReviews(cfg aws.Config, client *codegurureviewer.Client) {
	input := &codegurureviewer.ListCodeReviewsInput{
		// Type: types.Type, // Required
	}

	if len(_codegurureviewerType) > 0 {
		if err := assignInputField(input, "Type", _codegurureviewerType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_codegurureviewerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codegurureviewerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codegurureviewerNextToken) > 0 {
		input.NextToken = aws.String(_codegurureviewerNextToken)
	}
	if len(_codegurureviewerProviderTypes) > 0 {
		if err := assignInputField(input, "ProviderTypes", _codegurureviewerProviderTypes); err != nil {
			log.Errorf("invalid --provider-types: %s", err.Error())
			return
		}
	}
	if len(_codegurureviewerRepositoryNames) > 0 {
		input.RepositoryNames = append([]string(nil), _codegurureviewerRepositoryNames...)
	}
	if len(_codegurureviewerStates) > 0 {
		if err := assignInputField(input, "States", _codegurureviewerStates); err != nil {
			log.Errorf("invalid --states: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListCodeReviews(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codegurureviewer.ListCodeReviewsOutput
	p := codegurureviewer.NewListCodeReviewsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns a list of [RecommendationFeedbackSummary] objects that contain customer recommendation feedback for
// all CodeGuru Reviewer users.
//
// [RecommendationFeedbackSummary]: https://docs.aws.amazon.com/codeguru/latest/reviewer-api/API_RecommendationFeedbackSummary.html
func codegurureviewer_ListRecommendationFeedback(cfg aws.Config, client *codegurureviewer.Client) {
	input := &codegurureviewer.ListRecommendationFeedbackInput{
		// CodeReviewArn: *string, // Required
	}

	if len(_codegurureviewerCodeReviewArn) > 0 {
		input.CodeReviewArn = aws.String(_codegurureviewerCodeReviewArn)
	}
	if len(_codegurureviewerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codegurureviewerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codegurureviewerNextToken) > 0 {
		input.NextToken = aws.String(_codegurureviewerNextToken)
	}
	if len(_codegurureviewerRecommendationIds) > 0 {
		input.RecommendationIds = append([]string(nil), _codegurureviewerRecommendationIds...)
	}
	if len(_codegurureviewerUserIds) > 0 {
		input.UserIds = append([]string(nil), _codegurureviewerUserIds...)
	}

	if disablePaginator() {
		if resp, err := client.ListRecommendationFeedback(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codegurureviewer.ListRecommendationFeedbackOutput
	p := codegurureviewer.NewListRecommendationFeedbackPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns the list of all recommendations for a completed code review.
func codegurureviewer_ListRecommendations(cfg aws.Config, client *codegurureviewer.Client) {
	input := &codegurureviewer.ListRecommendationsInput{
		// CodeReviewArn: *string, // Required
	}

	if len(_codegurureviewerCodeReviewArn) > 0 {
		input.CodeReviewArn = aws.String(_codegurureviewerCodeReviewArn)
	}
	if len(_codegurureviewerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codegurureviewerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codegurureviewerNextToken) > 0 {
		input.NextToken = aws.String(_codegurureviewerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRecommendations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codegurureviewer.ListRecommendationsOutput
	p := codegurureviewer.NewListRecommendationsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns a list of [RepositoryAssociationSummary] objects that contain summary information about a repository
// association. You can filter the returned list by [ProviderType], [Name], [State], and [Owner].
//
// [Owner]: https://docs.aws.amazon.com/codeguru/latest/reviewer-api/API_RepositoryAssociationSummary.html#reviewer-Type-RepositoryAssociationSummary-Owner
// [State]: https://docs.aws.amazon.com/codeguru/latest/reviewer-api/API_RepositoryAssociationSummary.html#reviewer-Type-RepositoryAssociationSummary-State
// [ProviderType]: https://docs.aws.amazon.com/codeguru/latest/reviewer-api/API_RepositoryAssociationSummary.html#reviewer-Type-RepositoryAssociationSummary-ProviderType
// [Name]: https://docs.aws.amazon.com/codeguru/latest/reviewer-api/API_RepositoryAssociationSummary.html#reviewer-Type-RepositoryAssociationSummary-Name
// [RepositoryAssociationSummary]: https://docs.aws.amazon.com/codeguru/latest/reviewer-api/API_RepositoryAssociationSummary.html
func codegurureviewer_ListRepositoryAssociations(cfg aws.Config, client *codegurureviewer.Client) {
	input := &codegurureviewer.ListRepositoryAssociationsInput{}

	if len(_codegurureviewerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codegurureviewerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codegurureviewerNames) > 0 {
		input.Names = append([]string(nil), _codegurureviewerNames...)
	}
	if len(_codegurureviewerNextToken) > 0 {
		input.NextToken = aws.String(_codegurureviewerNextToken)
	}
	if len(_codegurureviewerOwners) > 0 {
		input.Owners = append([]string(nil), _codegurureviewerOwners...)
	}
	if len(_codegurureviewerProviderTypes) > 0 {
		if err := assignInputField(input, "ProviderTypes", _codegurureviewerProviderTypes); err != nil {
			log.Errorf("invalid --provider-types: %s", err.Error())
			return
		}
	}
	if len(_codegurureviewerStates) > 0 {
		if err := assignInputField(input, "States", _codegurureviewerStates); err != nil {
			log.Errorf("invalid --states: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListRepositoryAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codegurureviewer.ListRepositoryAssociationsOutput
	p := codegurureviewer.NewListRepositoryAssociationsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns the list of tags associated with an associated repository resource.
func codegurureviewer_ListTagsForResource(cfg aws.Config, client *codegurureviewer.Client) {
	input := &codegurureviewer.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_codegurureviewerResourceArn) > 0 {
		input.ResourceArn = aws.String(_codegurureviewerResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stores customer feedback for a CodeGuru Reviewer recommendation. When this API
// is called again with different reactions the previous feedback is overwritten.
func codegurureviewer_PutRecommendationFeedback(cfg aws.Config, client *codegurureviewer.Client) {
	input := &codegurureviewer.PutRecommendationFeedbackInput{
		// CodeReviewArn: *string, // Required
		// Reactions: []types.Reaction, // Required
		// RecommendationId: *string, // Required
	}

	if len(_codegurureviewerCodeReviewArn) > 0 {
		input.CodeReviewArn = aws.String(_codegurureviewerCodeReviewArn)
	}
	if len(_codegurureviewerReactions) > 0 {
		if err := assignInputField(input, "Reactions", _codegurureviewerReactions); err != nil {
			log.Errorf("invalid --reactions: %s", err.Error())
			return
		}
	}
	if len(_codegurureviewerRecommendationId) > 0 {
		input.RecommendationId = aws.String(_codegurureviewerRecommendationId)
	}

	if resp, err := client.PutRecommendationFeedback(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds one or more tags to an associated repository.
func codegurureviewer_TagResource(cfg aws.Config, client *codegurureviewer.Client) {
	input := &codegurureviewer.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_codegurureviewerResourceArn) > 0 {
		input.ResourceArn = aws.String(_codegurureviewerResourceArn)
	}
	if len(_codegurureviewerTags) > 0 {
		if err := assignInputField(input, "Tags", _codegurureviewerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.TagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a tag from an associated repository.
func codegurureviewer_UntagResource(cfg aws.Config, client *codegurureviewer.Client) {
	input := &codegurureviewer.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_codegurureviewerResourceArn) > 0 {
		input.ResourceArn = aws.String(_codegurureviewerResourceArn)
	}
	if len(_codegurureviewerTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _codegurureviewerTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_codegurureviewerCmd)
	_codegurureviewerCmd.Flags().SortFlags = false

	_codegurureviewerCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_codegurureviewerCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_codegurureviewerCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_codegurureviewerCmd.Flags().StringVarP(&_codegurureviewerAssociationArn, "association-arn", "", "", "Association ARN")
	_codegurureviewerCmd.Flags().StringVarP(&_codegurureviewerClientRequestToken, "client-request-token", "", "", "Client Request Token")
	_codegurureviewerCmd.Flags().StringVarP(&_codegurureviewerCodeReviewArn, "code-review-arn", "", "", "Code Review ARN")
	_codegurureviewerCmd.Flags().StringVarP(&_codegurureviewerKMSKeyDetails, "kms-key-details", "", "", "KMS Key Details")
	_codegurureviewerCmd.Flags().StringVarP(&_codegurureviewerMaxResults, "max-results", "", "", "Max Results")
	_codegurureviewerCmd.Flags().StringVarP(&_codegurureviewerName, "name", "", "", "Name")
	_codegurureviewerCmd.Flags().StringSliceVarP(&_codegurureviewerNames, "names", "", nil, "Names")
	_codegurureviewerCmd.Flags().StringVarP(&_codegurureviewerNextToken, "next-token", "", "", "Next Token")
	_codegurureviewerCmd.Flags().StringSliceVarP(&_codegurureviewerOwners, "owners", "", nil, "Owners")
	_codegurureviewerCmd.Flags().StringVarP(&_codegurureviewerProviderTypes, "provider-types", "", "", "Provider Types")
	_codegurureviewerCmd.Flags().StringVarP(&_codegurureviewerReactions, "reactions", "", "", "Reactions")
	_codegurureviewerCmd.Flags().StringVarP(&_codegurureviewerRecommendationId, "recommendation-id", "", "", "Recommendation ID")
	_codegurureviewerCmd.Flags().StringSliceVarP(&_codegurureviewerRecommendationIds, "recommendation-ids", "", nil, "Recommendation Ids")
	_codegurureviewerCmd.Flags().StringVarP(&_codegurureviewerRepository, "repository", "", "", "Repository")
	_codegurureviewerCmd.Flags().StringVarP(&_codegurureviewerRepositoryAssociationArn, "repository-association-arn", "", "", "Repository Association ARN")
	_codegurureviewerCmd.Flags().StringSliceVarP(&_codegurureviewerRepositoryNames, "repository-names", "", nil, "Repository Names")
	_codegurureviewerCmd.Flags().StringVarP(&_codegurureviewerResourceArn, "resource-arn", "", "", "Resource ARN")
	_codegurureviewerCmd.Flags().StringVarP(&_codegurureviewerStates, "states", "", "", "States")
	_codegurureviewerCmd.Flags().StringSliceVarP(&_codegurureviewerTagKeys, "tag-keys", "", nil, "Tag Keys")
	_codegurureviewerCmd.Flags().StringVarP(&_codegurureviewerTags, "tags", "", "", "Tags")
	_codegurureviewerCmd.Flags().StringVarP(&_codegurureviewerType, "type", "", "", "Type")
	_codegurureviewerCmd.Flags().StringVarP(&_codegurureviewerUserId, "user-id", "", "", "User ID")
	_codegurureviewerCmd.Flags().StringSliceVarP(&_codegurureviewerUserIds, "user-ids", "", nil, "User Ids")

	_codegurureviewerCmd.Flags().BoolVarP(&_codegurureviewerAssociateRepository, "associate-repository", "", false, "Associate Repository")
	_codegurureviewerCmd.Flags().BoolVarP(&_codegurureviewerCreateCodeReview, "create-code-review", "", false, "Create Code Review")
	_codegurureviewerCmd.Flags().BoolVarP(&_codegurureviewerDescribeCodeReview, "describe-code-review", "", false, "Describe Code Review")
	_codegurureviewerCmd.Flags().BoolVarP(&_codegurureviewerDescribeRecommendationFeedback, "describe-recommendation-feedback", "", false, "Describe Recommendation Feedback")
	_codegurureviewerCmd.Flags().BoolVarP(&_codegurureviewerDescribeRepositoryAssociation, "describe-repository-association", "", false, "Describe Repository Association")
	_codegurureviewerCmd.Flags().BoolVarP(&_codegurureviewerDisassociateRepository, "disassociate-repository", "", false, "Disassociate Repository")
	_codegurureviewerCmd.Flags().BoolVarP(&_codegurureviewerListCodeReviews, "list-code-reviews", "", false, "List Code Reviews")
	_codegurureviewerCmd.Flags().BoolVarP(&_codegurureviewerListRecommendationFeedback, "list-recommendation-feedback", "", false, "List Recommendation Feedback")
	_codegurureviewerCmd.Flags().BoolVarP(&_codegurureviewerListRecommendations, "list-recommendations", "", false, "List Recommendations")
	_codegurureviewerCmd.Flags().BoolVarP(&_codegurureviewerListRepositoryAssociations, "list-repository-associations", "", false, "List Repository Associations")
	_codegurureviewerCmd.Flags().BoolVarP(&_codegurureviewerListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_codegurureviewerCmd.Flags().BoolVarP(&_codegurureviewerPutRecommendationFeedback, "put-recommendation-feedback", "", false, "Put Recommendation Feedback")
	_codegurureviewerCmd.Flags().BoolVarP(&_codegurureviewerTagResource, "tag-resource", "", false, "Tag Resource")
	_codegurureviewerCmd.Flags().BoolVarP(&_codegurureviewerUntagResource, "untag-resource", "", false, "Untag Resource")

}
