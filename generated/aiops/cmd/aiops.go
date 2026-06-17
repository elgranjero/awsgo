package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/aiops"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// aiopsCmd represents the aiops command
var _aiopsCmd = &cobra.Command{
	Use:   "aiops",
	Short: "AWS aiops CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := aiops.NewFromConfig(cfg)
		if _aiopsCreateInvestigationGroup {
			aiops_CreateInvestigationGroup(cfg, client)
			return
		}
		if _aiopsDeleteInvestigationGroup {
			aiops_DeleteInvestigationGroup(cfg, client)
			return
		}
		if _aiopsDeleteInvestigationGroupPolicy {
			aiops_DeleteInvestigationGroupPolicy(cfg, client)
			return
		}
		if _aiopsGetInvestigationGroup {
			aiops_GetInvestigationGroup(cfg, client)
			return
		}
		if _aiopsGetInvestigationGroupPolicy {
			aiops_GetInvestigationGroupPolicy(cfg, client)
			return
		}
		if _aiopsListInvestigationGroups {
			aiops_ListInvestigationGroups(cfg, client)
			return
		}
		if _aiopsListTagsForResource {
			aiops_ListTagsForResource(cfg, client)
			return
		}
		if _aiopsPutInvestigationGroupPolicy {
			aiops_PutInvestigationGroupPolicy(cfg, client)
			return
		}
		if _aiopsTagResource {
			aiops_TagResource(cfg, client)
			return
		}
		if _aiopsUntagResource {
			aiops_UntagResource(cfg, client)
			return
		}
		if _aiopsUpdateInvestigationGroup {
			aiops_UpdateInvestigationGroup(cfg, client)
			return
		}

	},
}

var (
	_aiopsCreateInvestigationGroup       bool
	_aiopsDeleteInvestigationGroup       bool
	_aiopsDeleteInvestigationGroupPolicy bool
	_aiopsGetInvestigationGroup          bool
	_aiopsGetInvestigationGroupPolicy    bool
	_aiopsListInvestigationGroups        bool
	_aiopsListTagsForResource            bool
	_aiopsPutInvestigationGroupPolicy    bool
	_aiopsTagResource                    bool
	_aiopsUntagResource                  bool
	_aiopsUpdateInvestigationGroup       bool

	_aiopsChatbotNotificationChannel      string
	_aiopsCrossAccountConfigurations      string
	_aiopsEncryptionConfiguration         string
	_aiopsIdentifier                      string
	_aiopsIsCloudTrailEventHistoryEnabled string
	_aiopsMaxResults                      string
	_aiopsName                            string
	_aiopsNextToken                       string
	_aiopsPolicy                          string
	_aiopsResourceArn                     string
	_aiopsRetentionInDays                 string
	_aiopsRoleArn                         string
	_aiopsTagKeyBoundaries                []string
	_aiopsTagKeys                         []string
	_aiopsTags                            string
)

// Creates an investigation group in your account. Creating an investigation group
// is a one-time setup task for each Region in your account. It is a necessary task
// to be able to perform investigations.
//
// Settings in the investigation group help you centrally manage the common
// properties of your investigations, such as the following:
//
// - Who can access the investigations
//
// - Whether investigation data is encrypted with a customer managed Key
// Management Service key.
//
// - How long investigations and their data are retained by default.
//
// Currently, you can have one investigation group in each Region in your account.
// Each investigation in a Region is a part of the investigation group in that
// Region
//
// To create an investigation group and set up CloudWatch investigations, you must
// be signed in to an IAM principal that has either the AIOpsConsoleAdminPolicy or
// the AdministratorAccess IAM policy attached, or to an account that has similar
// permissions.
//
// You can configure CloudWatch alarms to start investigations and add events to
// investigations. If you create your investigation group with
// CreateInvestigationGroup and you want to enable alarms to do this, you must use
// PutInvestigationGroupPolicy to create a resource policy that grants this
// permission to CloudWatch alarms.
//
// For more information about configuring CloudWatch alarms, see [Using Amazon CloudWatch alarms]
//
// [Using Amazon CloudWatch alarms]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/AlarmThatSendsEmail.html
func aiops_CreateInvestigationGroup(cfg aws.Config, client *aiops.Client) {
	input := &aiops.CreateInvestigationGroupInput{
		// Name: *string, // Required
		// RoleArn: *string, // Required
	}

	if len(_aiopsName) > 0 {
		input.Name = aws.String(_aiopsName)
	}
	if len(_aiopsRoleArn) > 0 {
		input.RoleArn = aws.String(_aiopsRoleArn)
	}
	if len(_aiopsChatbotNotificationChannel) > 0 {
		if err := assignInputField(input, "ChatbotNotificationChannel", _aiopsChatbotNotificationChannel); err != nil {
			log.Errorf("invalid --chatbot-notification-channel: %s", err.Error())
			return
		}
	}
	if len(_aiopsCrossAccountConfigurations) > 0 {
		if err := assignInputField(input, "CrossAccountConfigurations", _aiopsCrossAccountConfigurations); err != nil {
			log.Errorf("invalid --cross-account-configurations: %s", err.Error())
			return
		}
	}
	if len(_aiopsEncryptionConfiguration) > 0 {
		if err := assignInputField(input, "EncryptionConfiguration", _aiopsEncryptionConfiguration); err != nil {
			log.Errorf("invalid --encryption-configuration: %s", err.Error())
			return
		}
	}
	if len(_aiopsIsCloudTrailEventHistoryEnabled) > 0 {
		if err := assignInputField(input, "IsCloudTrailEventHistoryEnabled", _aiopsIsCloudTrailEventHistoryEnabled); err != nil {
			log.Errorf("invalid --is-cloud-trail-event-history-enabled: %s", err.Error())
			return
		}
	}
	if len(_aiopsRetentionInDays) > 0 {
		if err := assignInputField(input, "RetentionInDays", _aiopsRetentionInDays); err != nil {
			log.Errorf("invalid --retention-in-days: %s", err.Error())
			return
		}
	}
	if len(_aiopsTagKeyBoundaries) > 0 {
		input.TagKeyBoundaries = append([]string(nil), _aiopsTagKeyBoundaries...)
	}
	if len(_aiopsTags) > 0 {
		if err := assignInputField(input, "Tags", _aiopsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateInvestigationGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified investigation group from your account. You can currently
// have one investigation group per Region in your account. After you delete an
// investigation group, you can later create a new investigation group in the same
// Region.
func aiops_DeleteInvestigationGroup(cfg aws.Config, client *aiops.Client) {
	input := &aiops.DeleteInvestigationGroupInput{
		// Identifier: *string, // Required
	}

	if len(_aiopsIdentifier) > 0 {
		input.Identifier = aws.String(_aiopsIdentifier)
	}

	if resp, err := client.DeleteInvestigationGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the IAM resource policy from being associated with the investigation
// group that you specify.
func aiops_DeleteInvestigationGroupPolicy(cfg aws.Config, client *aiops.Client) {
	input := &aiops.DeleteInvestigationGroupPolicyInput{
		// Identifier: *string, // Required
	}

	if len(_aiopsIdentifier) > 0 {
		input.Identifier = aws.String(_aiopsIdentifier)
	}

	if resp, err := client.DeleteInvestigationGroupPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the configuration information for the specified investigation group.
func aiops_GetInvestigationGroup(cfg aws.Config, client *aiops.Client) {
	input := &aiops.GetInvestigationGroupInput{
		// Identifier: *string, // Required
	}

	if len(_aiopsIdentifier) > 0 {
		input.Identifier = aws.String(_aiopsIdentifier)
	}

	if resp, err := client.GetInvestigationGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the JSON of the IAM resource policy associated with the specified
// investigation group in a string. For example,
// {\"Version\":\"2012-10-17\",\"Statement\":[{\"Effect\":\"Allow\",\"Principal\":{\"Service\":\"aiops.alarms.cloudwatch.amazonaws.com\"},\"Action\":[\"aiops:CreateInvestigation\",\"aiops:CreateInvestigationEvent\"],\"Resource\":\"*\",\"Condition\":{\"StringEquals\":{\"aws:SourceAccount\":\"111122223333\"},\"ArnLike\":{\"aws:SourceArn\":\"arn:aws:cloudwatch:us-east-1:111122223333:alarm:*\"}}}]}
// .
func aiops_GetInvestigationGroupPolicy(cfg aws.Config, client *aiops.Client) {
	input := &aiops.GetInvestigationGroupPolicyInput{
		// Identifier: *string, // Required
	}

	if len(_aiopsIdentifier) > 0 {
		input.Identifier = aws.String(_aiopsIdentifier)
	}

	if resp, err := client.GetInvestigationGroupPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the ARN and name of each investigation group in the account.
func aiops_ListInvestigationGroups(cfg aws.Config, client *aiops.Client) {
	input := &aiops.ListInvestigationGroupsInput{}

	if len(_aiopsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _aiopsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_aiopsNextToken) > 0 {
		input.NextToken = aws.String(_aiopsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListInvestigationGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*aiops.ListInvestigationGroupsOutput
	p := aiops.NewListInvestigationGroupsPaginator(client, input)
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

// Displays the tags associated with a CloudWatch investigations resource.
// Currently, investigation groups support tagging.
func aiops_ListTagsForResource(cfg aws.Config, client *aiops.Client) {
	input := &aiops.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_aiopsResourceArn) > 0 {
		input.ResourceArn = aws.String(_aiopsResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an IAM resource policy and assigns it to the specified investigation
// group.
//
// If you create your investigation group with CreateInvestigationGroup and you
// want to enable CloudWatch alarms to create investigations and add events to
// investigations, you must use this operation to create a policy similar to this
// example.
//
// { "Version": "2008-10-17", "Statement": [ { "Effect": "Allow", "Principal": {
// "Service": "aiops.alarms.cloudwatch.amazonaws.com" }, "Action": [
// "aiops:CreateInvestigation", "aiops:CreateInvestigationEvent" ], "Resource":
// "*", "Condition": { "StringEquals": { "aws:SourceAccount": "account-id" },
// "ArnLike": { "aws:SourceArn": "arn:aws:cloudwatch:region:account-id:alarm:*" } }
// } ] }
func aiops_PutInvestigationGroupPolicy(cfg aws.Config, client *aiops.Client) {
	input := &aiops.PutInvestigationGroupPolicyInput{
		// Identifier: *string, // Required
		// Policy: *string, // Required
	}

	if len(_aiopsIdentifier) > 0 {
		input.Identifier = aws.String(_aiopsIdentifier)
	}
	if len(_aiopsPolicy) > 0 {
		input.Policy = aws.String(_aiopsPolicy)
	}

	if resp, err := client.PutInvestigationGroupPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Assigns one or more tags (key-value pairs) to the specified resource.
// Tags can help you organize and categorize your resources. You can also use them
// to scope user permissions by granting a user permission to access or change only
// resources with certain tag values.
//
// Tags don't have any semantic meaning to Amazon Web Services and are interpreted
// strictly as strings of characters.
//
// You can associate as many as 50 tags with a resource.
func aiops_TagResource(cfg aws.Config, client *aiops.Client) {
	input := &aiops.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_aiopsResourceArn) > 0 {
		input.ResourceArn = aws.String(_aiopsResourceArn)
	}
	if len(_aiopsTags) > 0 {
		if err := assignInputField(input, "Tags", _aiopsTags); err != nil {
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
func aiops_UntagResource(cfg aws.Config, client *aiops.Client) {
	input := &aiops.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_aiopsResourceArn) > 0 {
		input.ResourceArn = aws.String(_aiopsResourceArn)
	}
	if len(_aiopsTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _aiopsTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration of the specified investigation group.
func aiops_UpdateInvestigationGroup(cfg aws.Config, client *aiops.Client) {
	input := &aiops.UpdateInvestigationGroupInput{
		// Identifier: *string, // Required
	}

	if len(_aiopsIdentifier) > 0 {
		input.Identifier = aws.String(_aiopsIdentifier)
	}
	if len(_aiopsChatbotNotificationChannel) > 0 {
		if err := assignInputField(input, "ChatbotNotificationChannel", _aiopsChatbotNotificationChannel); err != nil {
			log.Errorf("invalid --chatbot-notification-channel: %s", err.Error())
			return
		}
	}
	if len(_aiopsCrossAccountConfigurations) > 0 {
		if err := assignInputField(input, "CrossAccountConfigurations", _aiopsCrossAccountConfigurations); err != nil {
			log.Errorf("invalid --cross-account-configurations: %s", err.Error())
			return
		}
	}
	if len(_aiopsEncryptionConfiguration) > 0 {
		if err := assignInputField(input, "EncryptionConfiguration", _aiopsEncryptionConfiguration); err != nil {
			log.Errorf("invalid --encryption-configuration: %s", err.Error())
			return
		}
	}
	if len(_aiopsIsCloudTrailEventHistoryEnabled) > 0 {
		if err := assignInputField(input, "IsCloudTrailEventHistoryEnabled", _aiopsIsCloudTrailEventHistoryEnabled); err != nil {
			log.Errorf("invalid --is-cloud-trail-event-history-enabled: %s", err.Error())
			return
		}
	}
	if len(_aiopsRoleArn) > 0 {
		input.RoleArn = aws.String(_aiopsRoleArn)
	}
	if len(_aiopsTagKeyBoundaries) > 0 {
		input.TagKeyBoundaries = append([]string(nil), _aiopsTagKeyBoundaries...)
	}

	if resp, err := client.UpdateInvestigationGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_aiopsCmd)
	_aiopsCmd.Flags().SortFlags = false

	_aiopsCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_aiopsCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_aiopsCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_aiopsCmd.Flags().StringVarP(&_aiopsChatbotNotificationChannel, "chatbot-notification-channel", "", "", "Chatbot Notification Channel")
	_aiopsCmd.Flags().StringVarP(&_aiopsCrossAccountConfigurations, "cross-account-configurations", "", "", "Cross Account Configurations")
	_aiopsCmd.Flags().StringVarP(&_aiopsEncryptionConfiguration, "encryption-configuration", "", "", "Encryption Configuration")
	_aiopsCmd.Flags().StringVarP(&_aiopsIdentifier, "identifier", "", "", "Identifier")
	_aiopsCmd.Flags().StringVarP(&_aiopsIsCloudTrailEventHistoryEnabled, "is-cloud-trail-event-history-enabled", "", "", "Is Cloud Trail Event History Enabled")
	_aiopsCmd.Flags().StringVarP(&_aiopsMaxResults, "max-results", "", "", "Max Results")
	_aiopsCmd.Flags().StringVarP(&_aiopsName, "name", "", "", "Name")
	_aiopsCmd.Flags().StringVarP(&_aiopsNextToken, "next-token", "", "", "Next Token")
	_aiopsCmd.Flags().StringVarP(&_aiopsPolicy, "policy", "", "", "Policy")
	_aiopsCmd.Flags().StringVarP(&_aiopsResourceArn, "resource-arn", "", "", "Resource ARN")
	_aiopsCmd.Flags().StringVarP(&_aiopsRetentionInDays, "retention-in-days", "", "", "Retention In Days")
	_aiopsCmd.Flags().StringVarP(&_aiopsRoleArn, "role-arn", "", "", "Role ARN")
	_aiopsCmd.Flags().StringSliceVarP(&_aiopsTagKeyBoundaries, "tag-key-boundaries", "", nil, "Tag Key Boundaries")
	_aiopsCmd.Flags().StringSliceVarP(&_aiopsTagKeys, "tag-keys", "", nil, "Tag Keys")
	_aiopsCmd.Flags().StringVarP(&_aiopsTags, "tags", "", "", "Tags")

	_aiopsCmd.Flags().BoolVarP(&_aiopsCreateInvestigationGroup, "create-investigation-group", "", false, "Create Investigation Group")
	_aiopsCmd.Flags().BoolVarP(&_aiopsDeleteInvestigationGroup, "delete-investigation-group", "", false, "Delete Investigation Group")
	_aiopsCmd.Flags().BoolVarP(&_aiopsDeleteInvestigationGroupPolicy, "delete-investigation-group-policy", "", false, "Delete Investigation Group Policy")
	_aiopsCmd.Flags().BoolVarP(&_aiopsGetInvestigationGroup, "get-investigation-group", "", false, "Get Investigation Group")
	_aiopsCmd.Flags().BoolVarP(&_aiopsGetInvestigationGroupPolicy, "get-investigation-group-policy", "", false, "Get Investigation Group Policy")
	_aiopsCmd.Flags().BoolVarP(&_aiopsListInvestigationGroups, "list-investigation-groups", "", false, "List Investigation Groups")
	_aiopsCmd.Flags().BoolVarP(&_aiopsListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_aiopsCmd.Flags().BoolVarP(&_aiopsPutInvestigationGroupPolicy, "put-investigation-group-policy", "", false, "Put Investigation Group Policy")
	_aiopsCmd.Flags().BoolVarP(&_aiopsTagResource, "tag-resource", "", false, "Tag Resource")
	_aiopsCmd.Flags().BoolVarP(&_aiopsUntagResource, "untag-resource", "", false, "Untag Resource")
	_aiopsCmd.Flags().BoolVarP(&_aiopsUpdateInvestigationGroup, "update-investigation-group", "", false, "Update Investigation Group")

}
