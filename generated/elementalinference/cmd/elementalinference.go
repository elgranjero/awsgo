package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elementalinference"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// elementalinferenceCmd represents the elementalinference command
var _elementalinferenceCmd = &cobra.Command{
	Use:   "elementalinference",
	Short: "AWS elementalinference CLI",
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
		client := elementalinference.NewFromConfig(cfg)
		if _elementalinferenceAssociateFeed {
			elementalinference_AssociateFeed(cfg, client)
			return
		}
		if _elementalinferenceCreateFeed {
			elementalinference_CreateFeed(cfg, client)
			return
		}
		if _elementalinferenceDeleteFeed {
			elementalinference_DeleteFeed(cfg, client)
			return
		}
		if _elementalinferenceDisassociateFeed {
			elementalinference_DisassociateFeed(cfg, client)
			return
		}
		if _elementalinferenceGetFeed {
			elementalinference_GetFeed(cfg, client)
			return
		}
		if _elementalinferenceListFeeds {
			elementalinference_ListFeeds(cfg, client)
			return
		}
		if _elementalinferenceListTagsForResource {
			elementalinference_ListTagsForResource(cfg, client)
			return
		}
		if _elementalinferenceTagResource {
			elementalinference_TagResource(cfg, client)
			return
		}
		if _elementalinferenceUntagResource {
			elementalinference_UntagResource(cfg, client)
			return
		}
		if _elementalinferenceUpdateFeed {
			elementalinference_UpdateFeed(cfg, client)
			return
		}

	},
}

var (
	_elementalinferenceAssociateFeed       bool
	_elementalinferenceCreateFeed          bool
	_elementalinferenceDeleteFeed          bool
	_elementalinferenceDisassociateFeed    bool
	_elementalinferenceGetFeed             bool
	_elementalinferenceListFeeds           bool
	_elementalinferenceListTagsForResource bool
	_elementalinferenceTagResource         bool
	_elementalinferenceUntagResource       bool
	_elementalinferenceUpdateFeed          bool

	_elementalinferenceAssociatedResourceName string
	_elementalinferenceDryRun                 string
	_elementalinferenceId                     string
	_elementalinferenceMaxResults             string
	_elementalinferenceName                   string
	_elementalinferenceNextToken              string
	_elementalinferenceOutputs                string
	_elementalinferenceResourceArn            string
	_elementalinferenceTagKeys                []string
	_elementalinferenceTags                   string
)

// Associates a resource with the feed. The resource provides the input that
// Elemental Inference needs needs in order to perform an Elemental Inference
// feature, such as cropping video. You always provide the resource by associating
// it with a feed. You can associate only one resource with each feed.
func elementalinference_AssociateFeed(cfg aws.Config, client *elementalinference.Client) {
	input := &elementalinference.AssociateFeedInput{
		// AssociatedResourceName: *string, // Required
		// Id: *string, // Required
		// Outputs: []types.CreateOutput, // Required
	}

	if len(_elementalinferenceAssociatedResourceName) > 0 {
		input.AssociatedResourceName = aws.String(_elementalinferenceAssociatedResourceName)
	}
	if len(_elementalinferenceId) > 0 {
		input.Id = aws.String(_elementalinferenceId)
	}
	if len(_elementalinferenceOutputs) > 0 {
		if err := assignInputField(input, "Outputs", _elementalinferenceOutputs); err != nil {
			log.Errorf("invalid --outputs: %s", err.Error())
			return
		}
	}
	if len(_elementalinferenceDryRun) > 0 {
		if err := assignInputField(input, "DryRun", _elementalinferenceDryRun); err != nil {
			log.Errorf("invalid --dry-run: %s", err.Error())
			return
		}
	}

	if resp, err := client.AssociateFeed(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a feed. The feed is the target for live streams being sent by the
// calling application. An example of a calling application is AWS Elemental
// MediaLive. After you create the feed, you can associate a resource with the
// feed.
func elementalinference_CreateFeed(cfg aws.Config, client *elementalinference.Client) {
	input := &elementalinference.CreateFeedInput{
		// Name: *string, // Required
		// Outputs: []types.CreateOutput, // Required
	}

	if len(_elementalinferenceName) > 0 {
		input.Name = aws.String(_elementalinferenceName)
	}
	if len(_elementalinferenceOutputs) > 0 {
		if err := assignInputField(input, "Outputs", _elementalinferenceOutputs); err != nil {
			log.Errorf("invalid --outputs: %s", err.Error())
			return
		}
	}
	if len(_elementalinferenceTags) > 0 {
		if err := assignInputField(input, "Tags", _elementalinferenceTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateFeed(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified feed. The feed can be deleted at any time.
func elementalinference_DeleteFeed(cfg aws.Config, client *elementalinference.Client) {
	input := &elementalinference.DeleteFeedInput{
		// Id: *string, // Required
	}

	if len(_elementalinferenceId) > 0 {
		input.Id = aws.String(_elementalinferenceId)
	}

	if resp, err := client.DeleteFeed(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Releases the resource (for example, an MediaLive channel) that is associated
// with this feed. The outputs in the feed become disabled.
func elementalinference_DisassociateFeed(cfg aws.Config, client *elementalinference.Client) {
	input := &elementalinference.DisassociateFeedInput{
		// AssociatedResourceName: *string, // Required
		// Id: *string, // Required
	}

	if len(_elementalinferenceAssociatedResourceName) > 0 {
		input.AssociatedResourceName = aws.String(_elementalinferenceAssociatedResourceName)
	}
	if len(_elementalinferenceId) > 0 {
		input.Id = aws.String(_elementalinferenceId)
	}
	if len(_elementalinferenceDryRun) > 0 {
		if err := assignInputField(input, "DryRun", _elementalinferenceDryRun); err != nil {
			log.Errorf("invalid --dry-run: %s", err.Error())
			return
		}
	}

	if resp, err := client.DisassociateFeed(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the specified feed.
func elementalinference_GetFeed(cfg aws.Config, client *elementalinference.Client) {
	input := &elementalinference.GetFeedInput{
		// Id: *string, // Required
	}

	if len(_elementalinferenceId) > 0 {
		input.Id = aws.String(_elementalinferenceId)
	}

	if resp, err := client.GetFeed(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Displays a list of feeds that belong to this AWS account.
func elementalinference_ListFeeds(cfg aws.Config, client *elementalinference.Client) {
	input := &elementalinference.ListFeedsInput{}

	if len(_elementalinferenceMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _elementalinferenceMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_elementalinferenceNextToken) > 0 {
		input.NextToken = aws.String(_elementalinferenceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListFeeds(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*elementalinference.ListFeedsOutput
	p := elementalinference.NewListFeedsPaginator(client, input)
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

// List all tags that are on an Elemental Inference resource in the current region.
func elementalinference_ListTagsForResource(cfg aws.Config, client *elementalinference.Client) {
	input := &elementalinference.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_elementalinferenceResourceArn) > 0 {
		input.ResourceArn = aws.String(_elementalinferenceResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates the specified tags to the resource identified by the specified
// resourceArn in the current region. If existing tags on a resource are not
// specified in the request parameters, they are not changed. When a resource is
// deleted, the tags associated with that resource are also deleted.
func elementalinference_TagResource(cfg aws.Config, client *elementalinference.Client) {
	input := &elementalinference.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_elementalinferenceResourceArn) > 0 {
		input.ResourceArn = aws.String(_elementalinferenceResourceArn)
	}
	if len(_elementalinferenceTags) > 0 {
		if err := assignInputField(input, "Tags", _elementalinferenceTags); err != nil {
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

// Deletes specified tags from the specified resource in the current region.
func elementalinference_UntagResource(cfg aws.Config, client *elementalinference.Client) {
	input := &elementalinference.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_elementalinferenceResourceArn) > 0 {
		input.ResourceArn = aws.String(_elementalinferenceResourceArn)
	}
	if len(_elementalinferenceTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _elementalinferenceTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the name and/or outputs in a feed.
func elementalinference_UpdateFeed(cfg aws.Config, client *elementalinference.Client) {
	input := &elementalinference.UpdateFeedInput{
		// Id: *string, // Required
		// Name: *string, // Required
		// Outputs: []types.UpdateOutput, // Required
	}

	if len(_elementalinferenceId) > 0 {
		input.Id = aws.String(_elementalinferenceId)
	}
	if len(_elementalinferenceName) > 0 {
		input.Name = aws.String(_elementalinferenceName)
	}
	if len(_elementalinferenceOutputs) > 0 {
		if err := assignInputField(input, "Outputs", _elementalinferenceOutputs); err != nil {
			log.Errorf("invalid --outputs: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateFeed(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_elementalinferenceCmd)
	_elementalinferenceCmd.Flags().SortFlags = false

	_elementalinferenceCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_elementalinferenceCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_elementalinferenceCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_elementalinferenceCmd.Flags().StringVarP(&_elementalinferenceAssociatedResourceName, "associated-resource-name", "", "", "Associated Resource Name")
	_elementalinferenceCmd.Flags().StringVarP(&_elementalinferenceDryRun, "dry-run", "", "", "Dry Run")
	_elementalinferenceCmd.Flags().StringVarP(&_elementalinferenceId, "id", "", "", "ID")
	_elementalinferenceCmd.Flags().StringVarP(&_elementalinferenceMaxResults, "max-results", "", "", "Max Results")
	_elementalinferenceCmd.Flags().StringVarP(&_elementalinferenceName, "name", "", "", "Name")
	_elementalinferenceCmd.Flags().StringVarP(&_elementalinferenceNextToken, "next-token", "", "", "Next Token")
	_elementalinferenceCmd.Flags().StringVarP(&_elementalinferenceOutputs, "outputs", "", "", "Outputs")
	_elementalinferenceCmd.Flags().StringVarP(&_elementalinferenceResourceArn, "resource-arn", "", "", "Resource ARN")
	_elementalinferenceCmd.Flags().StringSliceVarP(&_elementalinferenceTagKeys, "tag-keys", "", nil, "Tag Keys")
	_elementalinferenceCmd.Flags().StringVarP(&_elementalinferenceTags, "tags", "", "", "Tags")

	_elementalinferenceCmd.Flags().BoolVarP(&_elementalinferenceAssociateFeed, "associate-feed", "", false, "Associate Feed")
	_elementalinferenceCmd.Flags().BoolVarP(&_elementalinferenceCreateFeed, "create-feed", "", false, "Create Feed")
	_elementalinferenceCmd.Flags().BoolVarP(&_elementalinferenceDeleteFeed, "delete-feed", "", false, "Delete Feed")
	_elementalinferenceCmd.Flags().BoolVarP(&_elementalinferenceDisassociateFeed, "disassociate-feed", "", false, "Disassociate Feed")
	_elementalinferenceCmd.Flags().BoolVarP(&_elementalinferenceGetFeed, "get-feed", "", false, "Get Feed")
	_elementalinferenceCmd.Flags().BoolVarP(&_elementalinferenceListFeeds, "list-feeds", "", false, "List Feeds")
	_elementalinferenceCmd.Flags().BoolVarP(&_elementalinferenceListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_elementalinferenceCmd.Flags().BoolVarP(&_elementalinferenceTagResource, "tag-resource", "", false, "Tag Resource")
	_elementalinferenceCmd.Flags().BoolVarP(&_elementalinferenceUntagResource, "untag-resource", "", false, "Untag Resource")
	_elementalinferenceCmd.Flags().BoolVarP(&_elementalinferenceUpdateFeed, "update-feed", "", false, "Update Feed")

}
