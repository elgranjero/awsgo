package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rbin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// rbinCmd represents the rbin command
var _rbinCmd = &cobra.Command{
	Use:   "rbin",
	Short: "AWS rbin CLI",
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
		client := rbin.NewFromConfig(cfg)
		if _rbinCreateRule {
			rbin_CreateRule(cfg, client)
			return
		}
		if _rbinDeleteRule {
			rbin_DeleteRule(cfg, client)
			return
		}
		if _rbinGetRule {
			rbin_GetRule(cfg, client)
			return
		}
		if _rbinListRules {
			rbin_ListRules(cfg, client)
			return
		}
		if _rbinListTagsForResource {
			rbin_ListTagsForResource(cfg, client)
			return
		}
		if _rbinLockRule {
			rbin_LockRule(cfg, client)
			return
		}
		if _rbinTagResource {
			rbin_TagResource(cfg, client)
			return
		}
		if _rbinUnlockRule {
			rbin_UnlockRule(cfg, client)
			return
		}
		if _rbinUntagResource {
			rbin_UntagResource(cfg, client)
			return
		}
		if _rbinUpdateRule {
			rbin_UpdateRule(cfg, client)
			return
		}

	},
}

var (
	_rbinCreateRule          bool
	_rbinDeleteRule          bool
	_rbinGetRule             bool
	_rbinListRules           bool
	_rbinListTagsForResource bool
	_rbinLockRule            bool
	_rbinTagResource         bool
	_rbinUnlockRule          bool
	_rbinUntagResource       bool
	_rbinUpdateRule          bool

	_rbinDescription         string
	_rbinExcludeResourceTags string
	_rbinIdentifier          string
	_rbinLockConfiguration   string
	_rbinLockState           string
	_rbinMaxResults          string
	_rbinNextToken           string
	_rbinResourceArn         string
	_rbinResourceTags        string
	_rbinResourceType        string
	_rbinRetentionPeriod     string
	_rbinTagKeys             []string
	_rbinTags                string
)

// Creates a Recycle Bin retention rule. You can create two types of retention
// rules:
//
// - Tag-level retention rules - These retention rules use resource tags to
// identify the resources to protect. For each retention rule, you specify one or
// more tag key and value pairs. Resources (of the specified type) that have at
// least one of these tag key and value pairs are automatically retained in the
// Recycle Bin upon deletion. Use this type of retention rule to protect specific
// resources in your account based on their tags.
//
// - Region-level retention rules - These retention rules, by default, apply to
// all of the resources (of the specified type) in the Region, even if the
// resources are not tagged. However, you can specify exclusion tags to exclude
// resources that have specific tags. Use this type of retention rule to protect
// all resources of a specific type in a Region.
//
// For more information, see [Create Recycle Bin retention rules] in the Amazon EBS User Guide.
//
// [Create Recycle Bin retention rules]: https://docs.aws.amazon.com/ebs/latest/userguide/recycle-bin.html
func rbin_CreateRule(cfg aws.Config, client *rbin.Client) {
	input := &rbin.CreateRuleInput{
		// ResourceType: types.ResourceType, // Required
		// RetentionPeriod: *types.RetentionPeriod, // Required
	}

	if len(_rbinResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _rbinResourceType); err != nil {
			log.Errorf("invalid --resource-type: %s", err.Error())
			return
		}
	}
	if len(_rbinRetentionPeriod) > 0 {
		if err := assignInputField(input, "RetentionPeriod", _rbinRetentionPeriod); err != nil {
			log.Errorf("invalid --retention-period: %s", err.Error())
			return
		}
	}
	if len(_rbinDescription) > 0 {
		input.Description = aws.String(_rbinDescription)
	}
	if len(_rbinExcludeResourceTags) > 0 {
		if err := assignInputField(input, "ExcludeResourceTags", _rbinExcludeResourceTags); err != nil {
			log.Errorf("invalid --exclude-resource-tags: %s", err.Error())
			return
		}
	}
	if len(_rbinLockConfiguration) > 0 {
		if err := assignInputField(input, "LockConfiguration", _rbinLockConfiguration); err != nil {
			log.Errorf("invalid --lock-configuration: %s", err.Error())
			return
		}
	}
	if len(_rbinResourceTags) > 0 {
		if err := assignInputField(input, "ResourceTags", _rbinResourceTags); err != nil {
			log.Errorf("invalid --resource-tags: %s", err.Error())
			return
		}
	}
	if len(_rbinTags) > 0 {
		if err := assignInputField(input, "Tags", _rbinTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a Recycle Bin retention rule. For more information, see [Delete Recycle Bin retention rules] in the Amazon
// Elastic Compute Cloud User Guide.
//
// [Delete Recycle Bin retention rules]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/recycle-bin-working-with-rules.html#recycle-bin-delete-rule
func rbin_DeleteRule(cfg aws.Config, client *rbin.Client) {
	input := &rbin.DeleteRuleInput{
		// Identifier: *string, // Required
	}

	if len(_rbinIdentifier) > 0 {
		input.Identifier = aws.String(_rbinIdentifier)
	}

	if resp, err := client.DeleteRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a Recycle Bin retention rule.
func rbin_GetRule(cfg aws.Config, client *rbin.Client) {
	input := &rbin.GetRuleInput{
		// Identifier: *string, // Required
	}

	if len(_rbinIdentifier) > 0 {
		input.Identifier = aws.String(_rbinIdentifier)
	}

	if resp, err := client.GetRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the Recycle Bin retention rules in the Region.
func rbin_ListRules(cfg aws.Config, client *rbin.Client) {
	input := &rbin.ListRulesInput{
		// ResourceType: types.ResourceType, // Required
	}

	if len(_rbinResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _rbinResourceType); err != nil {
			log.Errorf("invalid --resource-type: %s", err.Error())
			return
		}
	}
	if len(_rbinExcludeResourceTags) > 0 {
		if err := assignInputField(input, "ExcludeResourceTags", _rbinExcludeResourceTags); err != nil {
			log.Errorf("invalid --exclude-resource-tags: %s", err.Error())
			return
		}
	}
	if len(_rbinLockState) > 0 {
		if err := assignInputField(input, "LockState", _rbinLockState); err != nil {
			log.Errorf("invalid --lock-state: %s", err.Error())
			return
		}
	}
	if len(_rbinMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _rbinMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_rbinNextToken) > 0 {
		input.NextToken = aws.String(_rbinNextToken)
	}
	if len(_rbinResourceTags) > 0 {
		if err := assignInputField(input, "ResourceTags", _rbinResourceTags); err != nil {
			log.Errorf("invalid --resource-tags: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListRules(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rbin.ListRulesOutput
	p := rbin.NewListRulesPaginator(client, input)
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

// Lists the tags assigned to a retention rule.
func rbin_ListTagsForResource(cfg aws.Config, client *rbin.Client) {
	input := &rbin.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_rbinResourceArn) > 0 {
		input.ResourceArn = aws.String(_rbinResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Locks a Region-level retention rule. A locked retention rule can't be modified
// or deleted.
//
// You can't lock tag-level retention rules, or Region-level retention rules that
// have exclusion tags.
func rbin_LockRule(cfg aws.Config, client *rbin.Client) {
	input := &rbin.LockRuleInput{
		// Identifier: *string, // Required
		// LockConfiguration: *types.LockConfiguration, // Required
	}

	if len(_rbinIdentifier) > 0 {
		input.Identifier = aws.String(_rbinIdentifier)
	}
	if len(_rbinLockConfiguration) > 0 {
		if err := assignInputField(input, "LockConfiguration", _rbinLockConfiguration); err != nil {
			log.Errorf("invalid --lock-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.LockRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Assigns tags to the specified retention rule.
func rbin_TagResource(cfg aws.Config, client *rbin.Client) {
	input := &rbin.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_rbinResourceArn) > 0 {
		input.ResourceArn = aws.String(_rbinResourceArn)
	}
	if len(_rbinTags) > 0 {
		if err := assignInputField(input, "Tags", _rbinTags); err != nil {
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

// Unlocks a retention rule. After a retention rule is unlocked, it can be
// modified or deleted only after the unlock delay period expires.
func rbin_UnlockRule(cfg aws.Config, client *rbin.Client) {
	input := &rbin.UnlockRuleInput{
		// Identifier: *string, // Required
	}

	if len(_rbinIdentifier) > 0 {
		input.Identifier = aws.String(_rbinIdentifier)
	}

	if resp, err := client.UnlockRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Unassigns a tag from a retention rule.
func rbin_UntagResource(cfg aws.Config, client *rbin.Client) {
	input := &rbin.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_rbinResourceArn) > 0 {
		input.ResourceArn = aws.String(_rbinResourceArn)
	}
	if len(_rbinTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _rbinTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing Recycle Bin retention rule. You can update a retention
// rule's description, resource tags, and retention period at any time after
// creation. You can't update a retention rule's resource type after creation. For
// more information, see [Update Recycle Bin retention rules]in the Amazon Elastic Compute Cloud User Guide.
//
// [Update Recycle Bin retention rules]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/recycle-bin-working-with-rules.html#recycle-bin-update-rule
func rbin_UpdateRule(cfg aws.Config, client *rbin.Client) {
	input := &rbin.UpdateRuleInput{
		// Identifier: *string, // Required
	}

	if len(_rbinIdentifier) > 0 {
		input.Identifier = aws.String(_rbinIdentifier)
	}
	if len(_rbinDescription) > 0 {
		input.Description = aws.String(_rbinDescription)
	}
	if len(_rbinExcludeResourceTags) > 0 {
		if err := assignInputField(input, "ExcludeResourceTags", _rbinExcludeResourceTags); err != nil {
			log.Errorf("invalid --exclude-resource-tags: %s", err.Error())
			return
		}
	}
	if len(_rbinResourceTags) > 0 {
		if err := assignInputField(input, "ResourceTags", _rbinResourceTags); err != nil {
			log.Errorf("invalid --resource-tags: %s", err.Error())
			return
		}
	}
	if len(_rbinResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _rbinResourceType); err != nil {
			log.Errorf("invalid --resource-type: %s", err.Error())
			return
		}
	}
	if len(_rbinRetentionPeriod) > 0 {
		if err := assignInputField(input, "RetentionPeriod", _rbinRetentionPeriod); err != nil {
			log.Errorf("invalid --retention-period: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_rbinCmd)
	_rbinCmd.Flags().SortFlags = false

	_rbinCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_rbinCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_rbinCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_rbinCmd.Flags().StringVarP(&_rbinDescription, "description", "", "", "Description")
	_rbinCmd.Flags().StringVarP(&_rbinExcludeResourceTags, "exclude-resource-tags", "", "", "Exclude Resource Tags")
	_rbinCmd.Flags().StringVarP(&_rbinIdentifier, "identifier", "", "", "Identifier")
	_rbinCmd.Flags().StringVarP(&_rbinLockConfiguration, "lock-configuration", "", "", "Lock Configuration")
	_rbinCmd.Flags().StringVarP(&_rbinLockState, "lock-state", "", "", "Lock State")
	_rbinCmd.Flags().StringVarP(&_rbinMaxResults, "max-results", "", "", "Max Results")
	_rbinCmd.Flags().StringVarP(&_rbinNextToken, "next-token", "", "", "Next Token")
	_rbinCmd.Flags().StringVarP(&_rbinResourceArn, "resource-arn", "", "", "Resource ARN")
	_rbinCmd.Flags().StringVarP(&_rbinResourceTags, "resource-tags", "", "", "Resource Tags")
	_rbinCmd.Flags().StringVarP(&_rbinResourceType, "resource-type", "", "", "Resource Type")
	_rbinCmd.Flags().StringVarP(&_rbinRetentionPeriod, "retention-period", "", "", "Retention Period")
	_rbinCmd.Flags().StringSliceVarP(&_rbinTagKeys, "tag-keys", "", nil, "Tag Keys")
	_rbinCmd.Flags().StringVarP(&_rbinTags, "tags", "", "", "Tags")

	_rbinCmd.Flags().BoolVarP(&_rbinCreateRule, "create-rule", "", false, "Create Rule")
	_rbinCmd.Flags().BoolVarP(&_rbinDeleteRule, "delete-rule", "", false, "Delete Rule")
	_rbinCmd.Flags().BoolVarP(&_rbinGetRule, "get-rule", "", false, "Get Rule")
	_rbinCmd.Flags().BoolVarP(&_rbinListRules, "list-rules", "", false, "List Rules")
	_rbinCmd.Flags().BoolVarP(&_rbinListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_rbinCmd.Flags().BoolVarP(&_rbinLockRule, "lock-rule", "", false, "Lock Rule")
	_rbinCmd.Flags().BoolVarP(&_rbinTagResource, "tag-resource", "", false, "Tag Resource")
	_rbinCmd.Flags().BoolVarP(&_rbinUnlockRule, "unlock-rule", "", false, "Unlock Rule")
	_rbinCmd.Flags().BoolVarP(&_rbinUntagResource, "untag-resource", "", false, "Untag Resource")
	_rbinCmd.Flags().BoolVarP(&_rbinUpdateRule, "update-rule", "", false, "Update Rule")

}
