package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/oam"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// oamCmd represents the oam command
var _oamCmd = &cobra.Command{
	Use:   "oam",
	Short: "AWS oam CLI",
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
		client := oam.NewFromConfig(cfg)
		if _oamCreateLink {
			oam_CreateLink(cfg, client)
			return
		}
		if _oamCreateSink {
			oam_CreateSink(cfg, client)
			return
		}
		if _oamDeleteLink {
			oam_DeleteLink(cfg, client)
			return
		}
		if _oamDeleteSink {
			oam_DeleteSink(cfg, client)
			return
		}
		if _oamGetLink {
			oam_GetLink(cfg, client)
			return
		}
		if _oamGetSink {
			oam_GetSink(cfg, client)
			return
		}
		if _oamGetSinkPolicy {
			oam_GetSinkPolicy(cfg, client)
			return
		}
		if _oamListAttachedLinks {
			oam_ListAttachedLinks(cfg, client)
			return
		}
		if _oamListLinks {
			oam_ListLinks(cfg, client)
			return
		}
		if _oamListSinks {
			oam_ListSinks(cfg, client)
			return
		}
		if _oamListTagsForResource {
			oam_ListTagsForResource(cfg, client)
			return
		}
		if _oamPutSinkPolicy {
			oam_PutSinkPolicy(cfg, client)
			return
		}
		if _oamTagResource {
			oam_TagResource(cfg, client)
			return
		}
		if _oamUntagResource {
			oam_UntagResource(cfg, client)
			return
		}
		if _oamUpdateLink {
			oam_UpdateLink(cfg, client)
			return
		}

	},
}

var (
	_oamCreateLink          bool
	_oamCreateSink          bool
	_oamDeleteLink          bool
	_oamDeleteSink          bool
	_oamGetLink             bool
	_oamGetSink             bool
	_oamGetSinkPolicy       bool
	_oamListAttachedLinks   bool
	_oamListLinks           bool
	_oamListSinks           bool
	_oamListTagsForResource bool
	_oamPutSinkPolicy       bool
	_oamTagResource         bool
	_oamUntagResource       bool
	_oamUpdateLink          bool

	_oamIdentifier        string
	_oamIncludeTags       string
	_oamLabelTemplate     string
	_oamLinkConfiguration string
	_oamMaxResults        string
	_oamName              string
	_oamNextToken         string
	_oamPolicy            string
	_oamResourceArn       string
	_oamResourceTypes     string
	_oamSinkIdentifier    string
	_oamTagKeys           []string
	_oamTags              string
)

// Creates a link between a source account and a sink that you have created in a
// monitoring account. After the link is created, data is sent from the source
// account to the monitoring account. When you create a link, you can optionally
// specify filters that specify which metric namespaces and which log groups are
// shared from the source account to the monitoring account.
//
// Before you create a link, you must create a sink in the monitoring account and
// create a sink policy in that account. The sink policy must permit the source
// account to link to it. You can grant permission to source accounts by granting
// permission to an entire organization or to individual accounts.
//
// For more information, see [CreateSink] and [PutSinkPolicy].
//
// Each monitoring account can be linked to as many as 100,000 source accounts.
//
// Each source account can be linked to as many as five monitoring accounts.
//
// [CreateSink]: https://docs.aws.amazon.com/OAM/latest/APIReference/API_CreateSink.html
// [PutSinkPolicy]: https://docs.aws.amazon.com/OAM/latest/APIReference/API_PutSinkPolicy.html
func oam_CreateLink(cfg aws.Config, client *oam.Client) {
	input := &oam.CreateLinkInput{
		// LabelTemplate: *string, // Required
		// ResourceTypes: []types.ResourceType, // Required
		// SinkIdentifier: *string, // Required
	}

	if len(_oamLabelTemplate) > 0 {
		input.LabelTemplate = aws.String(_oamLabelTemplate)
	}
	if len(_oamResourceTypes) > 0 {
		if err := assignInputField(input, "ResourceTypes", _oamResourceTypes); err != nil {
			log.Errorf("invalid --resource-types: %s", err.Error())
			return
		}
	}
	if len(_oamSinkIdentifier) > 0 {
		input.SinkIdentifier = aws.String(_oamSinkIdentifier)
	}
	if len(_oamLinkConfiguration) > 0 {
		if err := assignInputField(input, "LinkConfiguration", _oamLinkConfiguration); err != nil {
			log.Errorf("invalid --link-configuration: %s", err.Error())
			return
		}
	}
	if len(_oamTags) > 0 {
		if err := assignInputField(input, "Tags", _oamTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateLink(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Use this to create a sink in the current account, so that it can be used as a
// monitoring account in CloudWatch cross-account observability. A sink is a
// resource that represents an attachment point in a monitoring account. Source
// accounts can link to the sink to send observability data.
//
// After you create a sink, you must create a sink policy that allows source
// accounts to attach to it. For more information, see [PutSinkPolicy].
//
// Each account can contain one sink per Region. If you delete a sink, you can
// then create a new one in that Region.
//
// [PutSinkPolicy]: https://docs.aws.amazon.com/OAM/latest/APIReference/API_PutSinkPolicy.html
func oam_CreateSink(cfg aws.Config, client *oam.Client) {
	input := &oam.CreateSinkInput{
		// Name: *string, // Required
	}

	if len(_oamName) > 0 {
		input.Name = aws.String(_oamName)
	}
	if len(_oamTags) > 0 {
		if err := assignInputField(input, "Tags", _oamTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSink(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a link between a monitoring account sink and a source account. You must
// run this operation in the source account.
func oam_DeleteLink(cfg aws.Config, client *oam.Client) {
	input := &oam.DeleteLinkInput{
		// Identifier: *string, // Required
	}

	if len(_oamIdentifier) > 0 {
		input.Identifier = aws.String(_oamIdentifier)
	}

	if resp, err := client.DeleteLink(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a sink. You must delete all links to a sink before you can delete that
// sink.
func oam_DeleteSink(cfg aws.Config, client *oam.Client) {
	input := &oam.DeleteSinkInput{
		// Identifier: *string, // Required
	}

	if len(_oamIdentifier) > 0 {
		input.Identifier = aws.String(_oamIdentifier)
	}

	if resp, err := client.DeleteSink(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns complete information about one link.
// To use this operation, provide the link ARN. To retrieve a list of link ARNs,
// use [ListLinks].
//
// [ListLinks]: https://docs.aws.amazon.com/OAM/latest/APIReference/API_ListLinks.html
func oam_GetLink(cfg aws.Config, client *oam.Client) {
	input := &oam.GetLinkInput{
		// Identifier: *string, // Required
	}

	if len(_oamIdentifier) > 0 {
		input.Identifier = aws.String(_oamIdentifier)
	}
	if len(_oamIncludeTags) > 0 {
		if err := assignInputField(input, "IncludeTags", _oamIncludeTags); err != nil {
			log.Errorf("invalid --include-tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetLink(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns complete information about one monitoring account sink.
// To use this operation, provide the sink ARN. To retrieve a list of sink ARNs,
// use [ListSinks].
//
// [ListSinks]: https://docs.aws.amazon.com/OAM/latest/APIReference/API_ListSinks.html
func oam_GetSink(cfg aws.Config, client *oam.Client) {
	input := &oam.GetSinkInput{
		// Identifier: *string, // Required
	}

	if len(_oamIdentifier) > 0 {
		input.Identifier = aws.String(_oamIdentifier)
	}
	if len(_oamIncludeTags) > 0 {
		if err := assignInputField(input, "IncludeTags", _oamIncludeTags); err != nil {
			log.Errorf("invalid --include-tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetSink(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the current sink policy attached to this sink. The sink policy
// specifies what accounts can attach to this sink as source accounts, and what
// types of data they can share.
func oam_GetSinkPolicy(cfg aws.Config, client *oam.Client) {
	input := &oam.GetSinkPolicyInput{
		// SinkIdentifier: *string, // Required
	}

	if len(_oamSinkIdentifier) > 0 {
		input.SinkIdentifier = aws.String(_oamSinkIdentifier)
	}

	if resp, err := client.GetSinkPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of source account links that are linked to this monitoring
// account sink.
//
// To use this operation, provide the sink ARN. To retrieve a list of sink ARNs,
// use [ListSinks].
//
// To find a list of links for one source account, use [ListLinks].
//
// [ListLinks]: https://docs.aws.amazon.com/OAM/latest/APIReference/API_ListLinks.html
// [ListSinks]: https://docs.aws.amazon.com/OAM/latest/APIReference/API_ListSinks.html
func oam_ListAttachedLinks(cfg aws.Config, client *oam.Client) {
	input := &oam.ListAttachedLinksInput{
		// SinkIdentifier: *string, // Required
	}

	if len(_oamSinkIdentifier) > 0 {
		input.SinkIdentifier = aws.String(_oamSinkIdentifier)
	}
	if len(_oamMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _oamMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_oamNextToken) > 0 {
		input.NextToken = aws.String(_oamNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAttachedLinks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*oam.ListAttachedLinksOutput
	p := oam.NewListAttachedLinksPaginator(client, input)
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

// Use this operation in a source account to return a list of links to monitoring
// account sinks that this source account has.
//
// To find a list of links for one monitoring account sink, use [ListAttachedLinks] from within the
// monitoring account.
//
// [ListAttachedLinks]: https://docs.aws.amazon.com/OAM/latest/APIReference/API_ListAttachedLinks.html
func oam_ListLinks(cfg aws.Config, client *oam.Client) {
	input := &oam.ListLinksInput{}

	if len(_oamMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _oamMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_oamNextToken) > 0 {
		input.NextToken = aws.String(_oamNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListLinks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*oam.ListLinksOutput
	p := oam.NewListLinksPaginator(client, input)
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

// Use this operation in a monitoring account to return the list of sinks created
// in that account.
func oam_ListSinks(cfg aws.Config, client *oam.Client) {
	input := &oam.ListSinksInput{}

	if len(_oamMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _oamMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_oamNextToken) > 0 {
		input.NextToken = aws.String(_oamNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSinks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*oam.ListSinksOutput
	p := oam.NewListSinksPaginator(client, input)
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

// Displays the tags associated with a resource. Both sinks and links support
// tagging.
func oam_ListTagsForResource(cfg aws.Config, client *oam.Client) {
	input := &oam.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_oamResourceArn) > 0 {
		input.ResourceArn = aws.String(_oamResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates the resource policy that grants permissions to source
// accounts to link to the monitoring account sink. When you create a sink policy,
// you can grant permissions to all accounts in an organization or to individual
// accounts.
//
// You can also use a sink policy to limit the types of data that is shared. The
// six types of services with their respective resource types that you can allow or
// deny are:
//
// - Metrics - Specify with AWS::CloudWatch::Metric
//
// - Log groups - Specify with AWS::Logs::LogGroup
//
// - Traces - Specify with AWS::XRay::Trace
//
// - Application Insights - Applications - Specify with
// AWS::ApplicationInsights::Application
//
// - Internet Monitor - Specify with AWS::InternetMonitor::Monitor
//
// - Application Signals - Specify with AWS::ApplicationSignals::Service and
// AWS::ApplicationSignals::ServiceLevelObjective
//
// See the examples in this section to see how to specify permitted source
// accounts and data types.
func oam_PutSinkPolicy(cfg aws.Config, client *oam.Client) {
	input := &oam.PutSinkPolicyInput{
		// Policy: *string, // Required
		// SinkIdentifier: *string, // Required
	}

	if len(_oamPolicy) > 0 {
		input.Policy = aws.String(_oamPolicy)
	}
	if len(_oamSinkIdentifier) > 0 {
		input.SinkIdentifier = aws.String(_oamSinkIdentifier)
	}

	if resp, err := client.PutSinkPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Assigns one or more tags (key-value pairs) to the specified resource. Both
// sinks and links can be tagged.
//
// Tags can help you organize and categorize your resources. You can also use them
// to scope user permissions by granting a user permission to access or change only
// resources with certain tag values.
//
// Tags don't have any semantic meaning to Amazon Web Services and are interpreted
// strictly as strings of characters.
//
// You can use the TagResource action with a resource that already has tags. If
// you specify a new tag key for the alarm, this tag is appended to the list of
// tags associated with the alarm. If you specify a tag key that is already
// associated with the alarm, the new tag value that you specify replaces the
// previous value for that tag.
//
// You can associate as many as 50 tags with a resource.
//
// Unlike tagging permissions in other Amazon Web Services services, to tag or
// untag links and sinks you must have the oam:ResourceTag permission. The
// iam:ResourceTag permission does not allow you to tag and untag links and sinks.
func oam_TagResource(cfg aws.Config, client *oam.Client) {
	input := &oam.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_oamResourceArn) > 0 {
		input.ResourceArn = aws.String(_oamResourceArn)
	}
	if len(_oamTags) > 0 {
		if err := assignInputField(input, "Tags", _oamTags); err != nil {
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

// Removes one or more tags from the specified resource.
// Unlike tagging permissions in other Amazon Web Services services, to tag or
// untag links and sinks you must have the oam:ResourceTag permission. The
// iam:TagResource permission does not allow you to tag and untag links and sinks.
func oam_UntagResource(cfg aws.Config, client *oam.Client) {
	input := &oam.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_oamResourceArn) > 0 {
		input.ResourceArn = aws.String(_oamResourceArn)
	}
	if len(_oamTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _oamTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Use this operation to change what types of data are shared from a source
// account to its linked monitoring account sink. You can't change the sink or
// change the monitoring account with this operation.
//
// When you update a link, you can optionally specify filters that specify which
// metric namespaces and which log groups are shared from the source account to the
// monitoring account.
//
// To update the list of tags associated with the sink, use [TagResource].
//
// [TagResource]: https://docs.aws.amazon.com/OAM/latest/APIReference/API_TagResource.html
func oam_UpdateLink(cfg aws.Config, client *oam.Client) {
	input := &oam.UpdateLinkInput{
		// Identifier: *string, // Required
		// ResourceTypes: []types.ResourceType, // Required
	}

	if len(_oamIdentifier) > 0 {
		input.Identifier = aws.String(_oamIdentifier)
	}
	if len(_oamResourceTypes) > 0 {
		if err := assignInputField(input, "ResourceTypes", _oamResourceTypes); err != nil {
			log.Errorf("invalid --resource-types: %s", err.Error())
			return
		}
	}
	if len(_oamIncludeTags) > 0 {
		if err := assignInputField(input, "IncludeTags", _oamIncludeTags); err != nil {
			log.Errorf("invalid --include-tags: %s", err.Error())
			return
		}
	}
	if len(_oamLinkConfiguration) > 0 {
		if err := assignInputField(input, "LinkConfiguration", _oamLinkConfiguration); err != nil {
			log.Errorf("invalid --link-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateLink(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_oamCmd)
	_oamCmd.Flags().SortFlags = false

	_oamCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_oamCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_oamCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_oamCmd.Flags().StringVarP(&_oamIdentifier, "identifier", "", "", "Identifier")
	_oamCmd.Flags().StringVarP(&_oamIncludeTags, "include-tags", "", "", "Include Tags")
	_oamCmd.Flags().StringVarP(&_oamLabelTemplate, "label-template", "", "", "Label Template")
	_oamCmd.Flags().StringVarP(&_oamLinkConfiguration, "link-configuration", "", "", "Link Configuration")
	_oamCmd.Flags().StringVarP(&_oamMaxResults, "max-results", "", "", "Max Results")
	_oamCmd.Flags().StringVarP(&_oamName, "name", "", "", "Name")
	_oamCmd.Flags().StringVarP(&_oamNextToken, "next-token", "", "", "Next Token")
	_oamCmd.Flags().StringVarP(&_oamPolicy, "policy", "", "", "Policy")
	_oamCmd.Flags().StringVarP(&_oamResourceArn, "resource-arn", "", "", "Resource ARN")
	_oamCmd.Flags().StringVarP(&_oamResourceTypes, "resource-types", "", "", "Resource Types")
	_oamCmd.Flags().StringVarP(&_oamSinkIdentifier, "sink-identifier", "", "", "Sink Identifier")
	_oamCmd.Flags().StringSliceVarP(&_oamTagKeys, "tag-keys", "", nil, "Tag Keys")
	_oamCmd.Flags().StringVarP(&_oamTags, "tags", "", "", "Tags")

	_oamCmd.Flags().BoolVarP(&_oamCreateLink, "create-link", "", false, "Create Link")
	_oamCmd.Flags().BoolVarP(&_oamCreateSink, "create-sink", "", false, "Create Sink")
	_oamCmd.Flags().BoolVarP(&_oamDeleteLink, "delete-link", "", false, "Delete Link")
	_oamCmd.Flags().BoolVarP(&_oamDeleteSink, "delete-sink", "", false, "Delete Sink")
	_oamCmd.Flags().BoolVarP(&_oamGetLink, "get-link", "", false, "Get Link")
	_oamCmd.Flags().BoolVarP(&_oamGetSink, "get-sink", "", false, "Get Sink")
	_oamCmd.Flags().BoolVarP(&_oamGetSinkPolicy, "get-sink-policy", "", false, "Get Sink Policy")
	_oamCmd.Flags().BoolVarP(&_oamListAttachedLinks, "list-attached-links", "", false, "List Attached Links")
	_oamCmd.Flags().BoolVarP(&_oamListLinks, "list-links", "", false, "List Links")
	_oamCmd.Flags().BoolVarP(&_oamListSinks, "list-sinks", "", false, "List Sinks")
	_oamCmd.Flags().BoolVarP(&_oamListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_oamCmd.Flags().BoolVarP(&_oamPutSinkPolicy, "put-sink-policy", "", false, "Put Sink Policy")
	_oamCmd.Flags().BoolVarP(&_oamTagResource, "tag-resource", "", false, "Tag Resource")
	_oamCmd.Flags().BoolVarP(&_oamUntagResource, "untag-resource", "", false, "Untag Resource")
	_oamCmd.Flags().BoolVarP(&_oamUpdateLink, "update-link", "", false, "Update Link")

}
