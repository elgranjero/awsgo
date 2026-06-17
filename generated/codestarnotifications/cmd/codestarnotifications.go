package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codestarnotifications"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// codestarnotificationsCmd represents the codestarnotifications command
var _codestarnotificationsCmd = &cobra.Command{
	Use:   "codestarnotifications",
	Short: "AWS codestarnotifications CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := codestarnotifications.NewFromConfig(cfg)
		if _codestarnotificationsCreateNotificationRule {
			codestarnotifications_CreateNotificationRule(cfg, client)
			return
		}
		if _codestarnotificationsDeleteNotificationRule {
			codestarnotifications_DeleteNotificationRule(cfg, client)
			return
		}
		if _codestarnotificationsDeleteTarget {
			codestarnotifications_DeleteTarget(cfg, client)
			return
		}
		if _codestarnotificationsDescribeNotificationRule {
			codestarnotifications_DescribeNotificationRule(cfg, client)
			return
		}
		if _codestarnotificationsListEventTypes {
			codestarnotifications_ListEventTypes(cfg, client)
			return
		}
		if _codestarnotificationsListNotificationRules {
			codestarnotifications_ListNotificationRules(cfg, client)
			return
		}
		if _codestarnotificationsListTagsForResource {
			codestarnotifications_ListTagsForResource(cfg, client)
			return
		}
		if _codestarnotificationsListTargets {
			codestarnotifications_ListTargets(cfg, client)
			return
		}
		if _codestarnotificationsSubscribe {
			codestarnotifications_Subscribe(cfg, client)
			return
		}
		if _codestarnotificationsTagResource {
			codestarnotifications_TagResource(cfg, client)
			return
		}
		if _codestarnotificationsUnsubscribe {
			codestarnotifications_Unsubscribe(cfg, client)
			return
		}
		if _codestarnotificationsUntagResource {
			codestarnotifications_UntagResource(cfg, client)
			return
		}
		if _codestarnotificationsUpdateNotificationRule {
			codestarnotifications_UpdateNotificationRule(cfg, client)
			return
		}

	},
}

var (
	_codestarnotificationsCreateNotificationRule   bool
	_codestarnotificationsDeleteNotificationRule   bool
	_codestarnotificationsDeleteTarget             bool
	_codestarnotificationsDescribeNotificationRule bool
	_codestarnotificationsListEventTypes           bool
	_codestarnotificationsListNotificationRules    bool
	_codestarnotificationsListTagsForResource      bool
	_codestarnotificationsListTargets              bool
	_codestarnotificationsSubscribe                bool
	_codestarnotificationsTagResource              bool
	_codestarnotificationsUnsubscribe              bool
	_codestarnotificationsUntagResource            bool
	_codestarnotificationsUpdateNotificationRule   bool

	_codestarnotificationsArn                 string
	_codestarnotificationsClientRequestToken  string
	_codestarnotificationsDetailType          string
	_codestarnotificationsEventTypeIds        []string
	_codestarnotificationsFilters             string
	_codestarnotificationsForceUnsubscribeAll string
	_codestarnotificationsMaxResults          string
	_codestarnotificationsName                string
	_codestarnotificationsNextToken           string
	_codestarnotificationsResource            string
	_codestarnotificationsStatus              string
	_codestarnotificationsTagKeys             []string
	_codestarnotificationsTags                string
	_codestarnotificationsTarget              string
	_codestarnotificationsTargetAddress       string
	_codestarnotificationsTargets             string
)

// Creates a notification rule for a resource. The rule specifies the events you
// want notifications about and the targets (such as Amazon Q Developer in chat
// applications topics or Amazon Q Developer in chat applications clients
// configured for Slack) where you want to receive them.
func codestarnotifications_CreateNotificationRule(cfg aws.Config, client *codestarnotifications.Client) {
	input := &codestarnotifications.CreateNotificationRuleInput{
		// DetailType: types.DetailType, // Required
		// EventTypeIds: []string, // Required
		// Name: *string, // Required
		// Resource: *string, // Required
		// Targets: []types.Target, // Required
	}

	if len(_codestarnotificationsDetailType) > 0 {
		if err := assignInputField(input, "DetailType", _codestarnotificationsDetailType); err != nil {
			log.Errorf("invalid --detail-type: %s", err.Error())
			return
		}
	}
	if len(_codestarnotificationsEventTypeIds) > 0 {
		input.EventTypeIds = append([]string(nil), _codestarnotificationsEventTypeIds...)
	}
	if len(_codestarnotificationsName) > 0 {
		input.Name = aws.String(_codestarnotificationsName)
	}
	if len(_codestarnotificationsResource) > 0 {
		input.Resource = aws.String(_codestarnotificationsResource)
	}
	if len(_codestarnotificationsTargets) > 0 {
		if err := assignInputField(input, "Targets", _codestarnotificationsTargets); err != nil {
			log.Errorf("invalid --targets: %s", err.Error())
			return
		}
	}
	if len(_codestarnotificationsClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_codestarnotificationsClientRequestToken)
	}
	if len(_codestarnotificationsStatus) > 0 {
		if err := assignInputField(input, "Status", _codestarnotificationsStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_codestarnotificationsTags) > 0 {
		if err := assignInputField(input, "Tags", _codestarnotificationsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateNotificationRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a notification rule for a resource.
func codestarnotifications_DeleteNotificationRule(cfg aws.Config, client *codestarnotifications.Client) {
	input := &codestarnotifications.DeleteNotificationRuleInput{
		// Arn: *string, // Required
	}

	if len(_codestarnotificationsArn) > 0 {
		input.Arn = aws.String(_codestarnotificationsArn)
	}

	if resp, err := client.DeleteNotificationRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a specified target for notifications.
func codestarnotifications_DeleteTarget(cfg aws.Config, client *codestarnotifications.Client) {
	input := &codestarnotifications.DeleteTargetInput{
		// TargetAddress: *string, // Required
	}

	if len(_codestarnotificationsTargetAddress) > 0 {
		input.TargetAddress = aws.String(_codestarnotificationsTargetAddress)
	}
	if len(_codestarnotificationsForceUnsubscribeAll) > 0 {
		if err := assignInputField(input, "ForceUnsubscribeAll", _codestarnotificationsForceUnsubscribeAll); err != nil {
			log.Errorf("invalid --force-unsubscribe-all: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteTarget(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a specified notification rule.
func codestarnotifications_DescribeNotificationRule(cfg aws.Config, client *codestarnotifications.Client) {
	input := &codestarnotifications.DescribeNotificationRuleInput{
		// Arn: *string, // Required
	}

	if len(_codestarnotificationsArn) > 0 {
		input.Arn = aws.String(_codestarnotificationsArn)
	}

	if resp, err := client.DescribeNotificationRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the event types available for configuring
// notifications.
func codestarnotifications_ListEventTypes(cfg aws.Config, client *codestarnotifications.Client) {
	input := &codestarnotifications.ListEventTypesInput{}

	if len(_codestarnotificationsFilters) > 0 {
		if err := assignInputField(input, "Filters", _codestarnotificationsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_codestarnotificationsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codestarnotificationsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codestarnotificationsNextToken) > 0 {
		input.NextToken = aws.String(_codestarnotificationsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListEventTypes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codestarnotifications.ListEventTypesOutput
	p := codestarnotifications.NewListEventTypesPaginator(client, input)
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

// Returns a list of the notification rules for an Amazon Web Services account.
func codestarnotifications_ListNotificationRules(cfg aws.Config, client *codestarnotifications.Client) {
	input := &codestarnotifications.ListNotificationRulesInput{}

	if len(_codestarnotificationsFilters) > 0 {
		if err := assignInputField(input, "Filters", _codestarnotificationsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_codestarnotificationsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codestarnotificationsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codestarnotificationsNextToken) > 0 {
		input.NextToken = aws.String(_codestarnotificationsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListNotificationRules(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codestarnotifications.ListNotificationRulesOutput
	p := codestarnotifications.NewListNotificationRulesPaginator(client, input)
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

// Returns a list of the tags associated with a notification rule.
func codestarnotifications_ListTagsForResource(cfg aws.Config, client *codestarnotifications.Client) {
	input := &codestarnotifications.ListTagsForResourceInput{
		// Arn: *string, // Required
	}

	if len(_codestarnotificationsArn) > 0 {
		input.Arn = aws.String(_codestarnotificationsArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of the notification rule targets for an Amazon Web Services
// account.
func codestarnotifications_ListTargets(cfg aws.Config, client *codestarnotifications.Client) {
	input := &codestarnotifications.ListTargetsInput{}

	if len(_codestarnotificationsFilters) > 0 {
		if err := assignInputField(input, "Filters", _codestarnotificationsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_codestarnotificationsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codestarnotificationsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codestarnotificationsNextToken) > 0 {
		input.NextToken = aws.String(_codestarnotificationsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTargets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codestarnotifications.ListTargetsOutput
	p := codestarnotifications.NewListTargetsPaginator(client, input)
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

// Creates an association between a notification rule and an Amazon Q Developer in
// chat applications topic or Amazon Q Developer in chat applications client so
// that the associated target can receive notifications when the events described
// in the rule are triggered.
func codestarnotifications_Subscribe(cfg aws.Config, client *codestarnotifications.Client) {
	input := &codestarnotifications.SubscribeInput{
		// Arn: *string, // Required
		// Target: *types.Target, // Required
	}

	if len(_codestarnotificationsArn) > 0 {
		input.Arn = aws.String(_codestarnotificationsArn)
	}
	if len(_codestarnotificationsTarget) > 0 {
		if err := assignInputField(input, "Target", _codestarnotificationsTarget); err != nil {
			log.Errorf("invalid --target: %s", err.Error())
			return
		}
	}
	if len(_codestarnotificationsClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_codestarnotificationsClientRequestToken)
	}

	if resp, err := client.Subscribe(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates a set of provided tags with a notification rule.
func codestarnotifications_TagResource(cfg aws.Config, client *codestarnotifications.Client) {
	input := &codestarnotifications.TagResourceInput{
		// Arn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_codestarnotificationsArn) > 0 {
		input.Arn = aws.String(_codestarnotificationsArn)
	}
	if len(_codestarnotificationsTags) > 0 {
		if err := assignInputField(input, "Tags", _codestarnotificationsTags); err != nil {
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

// Removes an association between a notification rule and an Amazon Q Developer in
// chat applications topic so that subscribers to that topic stop receiving
// notifications when the events described in the rule are triggered.
func codestarnotifications_Unsubscribe(cfg aws.Config, client *codestarnotifications.Client) {
	input := &codestarnotifications.UnsubscribeInput{
		// Arn: *string, // Required
		// TargetAddress: *string, // Required
	}

	if len(_codestarnotificationsArn) > 0 {
		input.Arn = aws.String(_codestarnotificationsArn)
	}
	if len(_codestarnotificationsTargetAddress) > 0 {
		input.TargetAddress = aws.String(_codestarnotificationsTargetAddress)
	}

	if resp, err := client.Unsubscribe(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the association between one or more provided tags and a notification
// rule.
func codestarnotifications_UntagResource(cfg aws.Config, client *codestarnotifications.Client) {
	input := &codestarnotifications.UntagResourceInput{
		// Arn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_codestarnotificationsArn) > 0 {
		input.Arn = aws.String(_codestarnotificationsArn)
	}
	if len(_codestarnotificationsTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _codestarnotificationsTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a notification rule for a resource. You can change the events that
// trigger the notification rule, the status of the rule, and the targets that
// receive the notifications.
//
// To add or remove tags for a notification rule, you must use TagResource and UntagResource.
func codestarnotifications_UpdateNotificationRule(cfg aws.Config, client *codestarnotifications.Client) {
	input := &codestarnotifications.UpdateNotificationRuleInput{
		// Arn: *string, // Required
	}

	if len(_codestarnotificationsArn) > 0 {
		input.Arn = aws.String(_codestarnotificationsArn)
	}
	if len(_codestarnotificationsDetailType) > 0 {
		if err := assignInputField(input, "DetailType", _codestarnotificationsDetailType); err != nil {
			log.Errorf("invalid --detail-type: %s", err.Error())
			return
		}
	}
	if len(_codestarnotificationsEventTypeIds) > 0 {
		input.EventTypeIds = append([]string(nil), _codestarnotificationsEventTypeIds...)
	}
	if len(_codestarnotificationsName) > 0 {
		input.Name = aws.String(_codestarnotificationsName)
	}
	if len(_codestarnotificationsStatus) > 0 {
		if err := assignInputField(input, "Status", _codestarnotificationsStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_codestarnotificationsTargets) > 0 {
		if err := assignInputField(input, "Targets", _codestarnotificationsTargets); err != nil {
			log.Errorf("invalid --targets: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateNotificationRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_codestarnotificationsCmd)
	_codestarnotificationsCmd.Flags().SortFlags = false

	_codestarnotificationsCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_codestarnotificationsCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_codestarnotificationsCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_codestarnotificationsCmd.Flags().StringVarP(&_codestarnotificationsArn, "arn", "", "", "ARN")
	_codestarnotificationsCmd.Flags().StringVarP(&_codestarnotificationsClientRequestToken, "client-request-token", "", "", "Client Request Token")
	_codestarnotificationsCmd.Flags().StringVarP(&_codestarnotificationsDetailType, "detail-type", "", "", "Detail Type")
	_codestarnotificationsCmd.Flags().StringSliceVarP(&_codestarnotificationsEventTypeIds, "event-type-ids", "", nil, "Event Type Ids")
	_codestarnotificationsCmd.Flags().StringVarP(&_codestarnotificationsFilters, "filters", "", "", "Filters")
	_codestarnotificationsCmd.Flags().StringVarP(&_codestarnotificationsForceUnsubscribeAll, "force-unsubscribe-all", "", "", "Force Unsubscribe All")
	_codestarnotificationsCmd.Flags().StringVarP(&_codestarnotificationsMaxResults, "max-results", "", "", "Max Results")
	_codestarnotificationsCmd.Flags().StringVarP(&_codestarnotificationsName, "name", "", "", "Name")
	_codestarnotificationsCmd.Flags().StringVarP(&_codestarnotificationsNextToken, "next-token", "", "", "Next Token")
	_codestarnotificationsCmd.Flags().StringVarP(&_codestarnotificationsResource, "resource", "", "", "Resource")
	_codestarnotificationsCmd.Flags().StringVarP(&_codestarnotificationsStatus, "status", "", "", "Status")
	_codestarnotificationsCmd.Flags().StringSliceVarP(&_codestarnotificationsTagKeys, "tag-keys", "", nil, "Tag Keys")
	_codestarnotificationsCmd.Flags().StringVarP(&_codestarnotificationsTags, "tags", "", "", "Tags")
	_codestarnotificationsCmd.Flags().StringVarP(&_codestarnotificationsTarget, "target", "", "", "Target")
	_codestarnotificationsCmd.Flags().StringVarP(&_codestarnotificationsTargetAddress, "target-address", "", "", "Target Address")
	_codestarnotificationsCmd.Flags().StringVarP(&_codestarnotificationsTargets, "targets", "", "", "Targets")

	_codestarnotificationsCmd.Flags().BoolVarP(&_codestarnotificationsCreateNotificationRule, "create-notification-rule", "", false, "Create Notification Rule")
	_codestarnotificationsCmd.Flags().BoolVarP(&_codestarnotificationsDeleteNotificationRule, "delete-notification-rule", "", false, "Delete Notification Rule")
	_codestarnotificationsCmd.Flags().BoolVarP(&_codestarnotificationsDeleteTarget, "delete-target", "", false, "Delete Target")
	_codestarnotificationsCmd.Flags().BoolVarP(&_codestarnotificationsDescribeNotificationRule, "describe-notification-rule", "", false, "Describe Notification Rule")
	_codestarnotificationsCmd.Flags().BoolVarP(&_codestarnotificationsListEventTypes, "list-event-types", "", false, "List Event Types")
	_codestarnotificationsCmd.Flags().BoolVarP(&_codestarnotificationsListNotificationRules, "list-notification-rules", "", false, "List Notification Rules")
	_codestarnotificationsCmd.Flags().BoolVarP(&_codestarnotificationsListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_codestarnotificationsCmd.Flags().BoolVarP(&_codestarnotificationsListTargets, "list-targets", "", false, "List Targets")
	_codestarnotificationsCmd.Flags().BoolVarP(&_codestarnotificationsSubscribe, "subscribe", "", false, "Subscribe")
	_codestarnotificationsCmd.Flags().BoolVarP(&_codestarnotificationsTagResource, "tag-resource", "", false, "Tag Resource")
	_codestarnotificationsCmd.Flags().BoolVarP(&_codestarnotificationsUnsubscribe, "unsubscribe", "", false, "Unsubscribe")
	_codestarnotificationsCmd.Flags().BoolVarP(&_codestarnotificationsUntagResource, "untag-resource", "", false, "Untag Resource")
	_codestarnotificationsCmd.Flags().BoolVarP(&_codestarnotificationsUpdateNotificationRule, "update-notification-rule", "", false, "Update Notification Rule")

}
