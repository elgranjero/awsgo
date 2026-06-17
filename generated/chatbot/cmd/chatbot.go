package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/chatbot"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// chatbotCmd represents the chatbot command
var _chatbotCmd = &cobra.Command{
	Use:   "chatbot",
	Short: "AWS chatbot CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := chatbot.NewFromConfig(cfg)
		if _chatbotAssociateToConfiguration {
			chatbot_AssociateToConfiguration(cfg, client)
			return
		}
		if _chatbotCreateChimeWebhookConfiguration {
			chatbot_CreateChimeWebhookConfiguration(cfg, client)
			return
		}
		if _chatbotCreateCustomAction {
			chatbot_CreateCustomAction(cfg, client)
			return
		}
		if _chatbotCreateMicrosoftTeamsChannelConfiguration {
			chatbot_CreateMicrosoftTeamsChannelConfiguration(cfg, client)
			return
		}
		if _chatbotCreateSlackChannelConfiguration {
			chatbot_CreateSlackChannelConfiguration(cfg, client)
			return
		}
		if _chatbotDeleteChimeWebhookConfiguration {
			chatbot_DeleteChimeWebhookConfiguration(cfg, client)
			return
		}
		if _chatbotDeleteCustomAction {
			chatbot_DeleteCustomAction(cfg, client)
			return
		}
		if _chatbotDeleteMicrosoftTeamsChannelConfiguration {
			chatbot_DeleteMicrosoftTeamsChannelConfiguration(cfg, client)
			return
		}
		if _chatbotDeleteMicrosoftTeamsConfiguredTeam {
			chatbot_DeleteMicrosoftTeamsConfiguredTeam(cfg, client)
			return
		}
		if _chatbotDeleteMicrosoftTeamsUserIdentity {
			chatbot_DeleteMicrosoftTeamsUserIdentity(cfg, client)
			return
		}
		if _chatbotDeleteSlackChannelConfiguration {
			chatbot_DeleteSlackChannelConfiguration(cfg, client)
			return
		}
		if _chatbotDeleteSlackUserIdentity {
			chatbot_DeleteSlackUserIdentity(cfg, client)
			return
		}
		if _chatbotDeleteSlackWorkspaceAuthorization {
			chatbot_DeleteSlackWorkspaceAuthorization(cfg, client)
			return
		}
		if _chatbotDescribeChimeWebhookConfigurations {
			chatbot_DescribeChimeWebhookConfigurations(cfg, client)
			return
		}
		if _chatbotDescribeSlackChannelConfigurations {
			chatbot_DescribeSlackChannelConfigurations(cfg, client)
			return
		}
		if _chatbotDescribeSlackUserIdentities {
			chatbot_DescribeSlackUserIdentities(cfg, client)
			return
		}
		if _chatbotDescribeSlackWorkspaces {
			chatbot_DescribeSlackWorkspaces(cfg, client)
			return
		}
		if _chatbotDisassociateFromConfiguration {
			chatbot_DisassociateFromConfiguration(cfg, client)
			return
		}
		if _chatbotGetAccountPreferences {
			chatbot_GetAccountPreferences(cfg, client)
			return
		}
		if _chatbotGetCustomAction {
			chatbot_GetCustomAction(cfg, client)
			return
		}
		if _chatbotGetMicrosoftTeamsChannelConfiguration {
			chatbot_GetMicrosoftTeamsChannelConfiguration(cfg, client)
			return
		}
		if _chatbotListAssociations {
			chatbot_ListAssociations(cfg, client)
			return
		}
		if _chatbotListCustomActions {
			chatbot_ListCustomActions(cfg, client)
			return
		}
		if _chatbotListMicrosoftTeamsChannelConfigurations {
			chatbot_ListMicrosoftTeamsChannelConfigurations(cfg, client)
			return
		}
		if _chatbotListMicrosoftTeamsConfiguredTeams {
			chatbot_ListMicrosoftTeamsConfiguredTeams(cfg, client)
			return
		}
		if _chatbotListMicrosoftTeamsUserIdentities {
			chatbot_ListMicrosoftTeamsUserIdentities(cfg, client)
			return
		}
		if _chatbotListTagsForResource {
			chatbot_ListTagsForResource(cfg, client)
			return
		}
		if _chatbotTagResource {
			chatbot_TagResource(cfg, client)
			return
		}
		if _chatbotUntagResource {
			chatbot_UntagResource(cfg, client)
			return
		}
		if _chatbotUpdateAccountPreferences {
			chatbot_UpdateAccountPreferences(cfg, client)
			return
		}
		if _chatbotUpdateChimeWebhookConfiguration {
			chatbot_UpdateChimeWebhookConfiguration(cfg, client)
			return
		}
		if _chatbotUpdateCustomAction {
			chatbot_UpdateCustomAction(cfg, client)
			return
		}
		if _chatbotUpdateMicrosoftTeamsChannelConfiguration {
			chatbot_UpdateMicrosoftTeamsChannelConfiguration(cfg, client)
			return
		}
		if _chatbotUpdateSlackChannelConfiguration {
			chatbot_UpdateSlackChannelConfiguration(cfg, client)
			return
		}

	},
}

var (
	_chatbotAssociateToConfiguration                 bool
	_chatbotCreateChimeWebhookConfiguration          bool
	_chatbotCreateCustomAction                       bool
	_chatbotCreateMicrosoftTeamsChannelConfiguration bool
	_chatbotCreateSlackChannelConfiguration          bool
	_chatbotDeleteChimeWebhookConfiguration          bool
	_chatbotDeleteCustomAction                       bool
	_chatbotDeleteMicrosoftTeamsChannelConfiguration bool
	_chatbotDeleteMicrosoftTeamsConfiguredTeam       bool
	_chatbotDeleteMicrosoftTeamsUserIdentity         bool
	_chatbotDeleteSlackChannelConfiguration          bool
	_chatbotDeleteSlackUserIdentity                  bool
	_chatbotDeleteSlackWorkspaceAuthorization        bool
	_chatbotDescribeChimeWebhookConfigurations       bool
	_chatbotDescribeSlackChannelConfigurations       bool
	_chatbotDescribeSlackUserIdentities              bool
	_chatbotDescribeSlackWorkspaces                  bool
	_chatbotDisassociateFromConfiguration            bool
	_chatbotGetAccountPreferences                    bool
	_chatbotGetCustomAction                          bool
	_chatbotGetMicrosoftTeamsChannelConfiguration    bool
	_chatbotListAssociations                         bool
	_chatbotListCustomActions                        bool
	_chatbotListMicrosoftTeamsChannelConfigurations  bool
	_chatbotListMicrosoftTeamsConfiguredTeams        bool
	_chatbotListMicrosoftTeamsUserIdentities         bool
	_chatbotListTagsForResource                      bool
	_chatbotTagResource                              bool
	_chatbotUntagResource                            bool
	_chatbotUpdateAccountPreferences                 bool
	_chatbotUpdateChimeWebhookConfiguration          bool
	_chatbotUpdateCustomAction                       bool
	_chatbotUpdateMicrosoftTeamsChannelConfiguration bool
	_chatbotUpdateSlackChannelConfiguration          bool

	_chatbotActionName                    string
	_chatbotAliasName                     string
	_chatbotAttachments                   string
	_chatbotChannelId                     string
	_chatbotChannelName                   string
	_chatbotChatConfiguration             string
	_chatbotChatConfigurationArn          string
	_chatbotClientToken                   string
	_chatbotConfigurationName             string
	_chatbotCustomActionArn               string
	_chatbotDefinition                    string
	_chatbotGuardrailPolicyArns           []string
	_chatbotIamRoleArn                    string
	_chatbotLoggingLevel                  string
	_chatbotMaxResults                    string
	_chatbotNextToken                     string
	_chatbotResource                      string
	_chatbotResourceARN                   string
	_chatbotSlackChannelId                string
	_chatbotSlackChannelName              string
	_chatbotSlackTeamId                   string
	_chatbotSlackUserId                   string
	_chatbotSnsTopicArns                  []string
	_chatbotTagKeys                       []string
	_chatbotTags                          string
	_chatbotTeamId                        string
	_chatbotTeamName                      string
	_chatbotTenantId                      string
	_chatbotTrainingDataCollectionEnabled string
	_chatbotUserAuthorizationRequired     string
	_chatbotUserId                        string
	_chatbotWebhookDescription            string
	_chatbotWebhookUrl                    string
)

// Links a resource (for example, a custom action) to a channel configuration.
func chatbot_AssociateToConfiguration(cfg aws.Config, client *chatbot.Client) {
	input := &chatbot.AssociateToConfigurationInput{
		// ChatConfiguration: *string, // Required
		// Resource: *string, // Required
	}

	if len(_chatbotChatConfiguration) > 0 {
		input.ChatConfiguration = aws.String(_chatbotChatConfiguration)
	}
	if len(_chatbotResource) > 0 {
		input.Resource = aws.String(_chatbotResource)
	}

	if resp, err := client.AssociateToConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an AWS Chatbot configuration for Amazon Chime.
func chatbot_CreateChimeWebhookConfiguration(cfg aws.Config, client *chatbot.Client) {
	input := &chatbot.CreateChimeWebhookConfigurationInput{
		// ConfigurationName: *string, // Required
		// IamRoleArn: *string, // Required
		// SnsTopicArns: []string, // Required
		// WebhookDescription: *string, // Required
		// WebhookUrl: *string, // Required
	}

	if len(_chatbotConfigurationName) > 0 {
		input.ConfigurationName = aws.String(_chatbotConfigurationName)
	}
	if len(_chatbotIamRoleArn) > 0 {
		input.IamRoleArn = aws.String(_chatbotIamRoleArn)
	}
	if len(_chatbotSnsTopicArns) > 0 {
		input.SnsTopicArns = append([]string(nil), _chatbotSnsTopicArns...)
	}
	if len(_chatbotWebhookDescription) > 0 {
		input.WebhookDescription = aws.String(_chatbotWebhookDescription)
	}
	if len(_chatbotWebhookUrl) > 0 {
		input.WebhookUrl = aws.String(_chatbotWebhookUrl)
	}
	if len(_chatbotLoggingLevel) > 0 {
		input.LoggingLevel = aws.String(_chatbotLoggingLevel)
	}
	if len(_chatbotTags) > 0 {
		if err := assignInputField(input, "Tags", _chatbotTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateChimeWebhookConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a custom action that can be invoked as an alias or as a button on a
// notification.
func chatbot_CreateCustomAction(cfg aws.Config, client *chatbot.Client) {
	input := &chatbot.CreateCustomActionInput{
		// ActionName: *string, // Required
		// Definition: *types.CustomActionDefinition, // Required
	}

	if len(_chatbotActionName) > 0 {
		input.ActionName = aws.String(_chatbotActionName)
	}
	if len(_chatbotDefinition) > 0 {
		if err := assignInputField(input, "Definition", _chatbotDefinition); err != nil {
			log.Errorf("invalid --definition: %s", err.Error())
			return
		}
	}
	if len(_chatbotAliasName) > 0 {
		input.AliasName = aws.String(_chatbotAliasName)
	}
	if len(_chatbotAttachments) > 0 {
		if err := assignInputField(input, "Attachments", _chatbotAttachments); err != nil {
			log.Errorf("invalid --attachments: %s", err.Error())
			return
		}
	}
	if len(_chatbotClientToken) > 0 {
		input.ClientToken = aws.String(_chatbotClientToken)
	}
	if len(_chatbotTags) > 0 {
		if err := assignInputField(input, "Tags", _chatbotTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCustomAction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an AWS Chatbot configuration for Microsoft Teams.
func chatbot_CreateMicrosoftTeamsChannelConfiguration(cfg aws.Config, client *chatbot.Client) {
	input := &chatbot.CreateMicrosoftTeamsChannelConfigurationInput{
		// ChannelId: *string, // Required
		// ConfigurationName: *string, // Required
		// IamRoleArn: *string, // Required
		// TeamId: *string, // Required
		// TenantId: *string, // Required
	}

	if len(_chatbotChannelId) > 0 {
		input.ChannelId = aws.String(_chatbotChannelId)
	}
	if len(_chatbotConfigurationName) > 0 {
		input.ConfigurationName = aws.String(_chatbotConfigurationName)
	}
	if len(_chatbotIamRoleArn) > 0 {
		input.IamRoleArn = aws.String(_chatbotIamRoleArn)
	}
	if len(_chatbotTeamId) > 0 {
		input.TeamId = aws.String(_chatbotTeamId)
	}
	if len(_chatbotTenantId) > 0 {
		input.TenantId = aws.String(_chatbotTenantId)
	}
	if len(_chatbotChannelName) > 0 {
		input.ChannelName = aws.String(_chatbotChannelName)
	}
	if len(_chatbotGuardrailPolicyArns) > 0 {
		input.GuardrailPolicyArns = append([]string(nil), _chatbotGuardrailPolicyArns...)
	}
	if len(_chatbotLoggingLevel) > 0 {
		input.LoggingLevel = aws.String(_chatbotLoggingLevel)
	}
	if len(_chatbotSnsTopicArns) > 0 {
		input.SnsTopicArns = append([]string(nil), _chatbotSnsTopicArns...)
	}
	if len(_chatbotTags) > 0 {
		if err := assignInputField(input, "Tags", _chatbotTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_chatbotTeamName) > 0 {
		input.TeamName = aws.String(_chatbotTeamName)
	}
	if len(_chatbotUserAuthorizationRequired) > 0 {
		if err := assignInputField(input, "UserAuthorizationRequired", _chatbotUserAuthorizationRequired); err != nil {
			log.Errorf("invalid --user-authorization-required: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateMicrosoftTeamsChannelConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an AWS Chatbot confugration for Slack.
func chatbot_CreateSlackChannelConfiguration(cfg aws.Config, client *chatbot.Client) {
	input := &chatbot.CreateSlackChannelConfigurationInput{
		// ConfigurationName: *string, // Required
		// IamRoleArn: *string, // Required
		// SlackChannelId: *string, // Required
		// SlackTeamId: *string, // Required
	}

	if len(_chatbotConfigurationName) > 0 {
		input.ConfigurationName = aws.String(_chatbotConfigurationName)
	}
	if len(_chatbotIamRoleArn) > 0 {
		input.IamRoleArn = aws.String(_chatbotIamRoleArn)
	}
	if len(_chatbotSlackChannelId) > 0 {
		input.SlackChannelId = aws.String(_chatbotSlackChannelId)
	}
	if len(_chatbotSlackTeamId) > 0 {
		input.SlackTeamId = aws.String(_chatbotSlackTeamId)
	}
	if len(_chatbotGuardrailPolicyArns) > 0 {
		input.GuardrailPolicyArns = append([]string(nil), _chatbotGuardrailPolicyArns...)
	}
	if len(_chatbotLoggingLevel) > 0 {
		input.LoggingLevel = aws.String(_chatbotLoggingLevel)
	}
	if len(_chatbotSlackChannelName) > 0 {
		input.SlackChannelName = aws.String(_chatbotSlackChannelName)
	}
	if len(_chatbotSnsTopicArns) > 0 {
		input.SnsTopicArns = append([]string(nil), _chatbotSnsTopicArns...)
	}
	if len(_chatbotTags) > 0 {
		if err := assignInputField(input, "Tags", _chatbotTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_chatbotUserAuthorizationRequired) > 0 {
		if err := assignInputField(input, "UserAuthorizationRequired", _chatbotUserAuthorizationRequired); err != nil {
			log.Errorf("invalid --user-authorization-required: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSlackChannelConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a Amazon Chime webhook configuration for AWS Chatbot.
func chatbot_DeleteChimeWebhookConfiguration(cfg aws.Config, client *chatbot.Client) {
	input := &chatbot.DeleteChimeWebhookConfigurationInput{
		// ChatConfigurationArn: *string, // Required
	}

	if len(_chatbotChatConfigurationArn) > 0 {
		input.ChatConfigurationArn = aws.String(_chatbotChatConfigurationArn)
	}

	if resp, err := client.DeleteChimeWebhookConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a custom action.
func chatbot_DeleteCustomAction(cfg aws.Config, client *chatbot.Client) {
	input := &chatbot.DeleteCustomActionInput{
		// CustomActionArn: *string, // Required
	}

	if len(_chatbotCustomActionArn) > 0 {
		input.CustomActionArn = aws.String(_chatbotCustomActionArn)
	}

	if resp, err := client.DeleteCustomAction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a Microsoft Teams channel configuration for AWS Chatbot
func chatbot_DeleteMicrosoftTeamsChannelConfiguration(cfg aws.Config, client *chatbot.Client) {
	input := &chatbot.DeleteMicrosoftTeamsChannelConfigurationInput{
		// ChatConfigurationArn: *string, // Required
	}

	if len(_chatbotChatConfigurationArn) > 0 {
		input.ChatConfigurationArn = aws.String(_chatbotChatConfigurationArn)
	}

	if resp, err := client.DeleteMicrosoftTeamsChannelConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the Microsoft Teams team authorization allowing for channels to be
// configured in that Microsoft Teams team. Note that the Microsoft Teams team must
// have no channels configured to remove it.
func chatbot_DeleteMicrosoftTeamsConfiguredTeam(cfg aws.Config, client *chatbot.Client) {
	input := &chatbot.DeleteMicrosoftTeamsConfiguredTeamInput{
		// TeamId: *string, // Required
	}

	if len(_chatbotTeamId) > 0 {
		input.TeamId = aws.String(_chatbotTeamId)
	}

	if resp, err := client.DeleteMicrosoftTeamsConfiguredTeam(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Identifes a user level permission for a channel configuration.
func chatbot_DeleteMicrosoftTeamsUserIdentity(cfg aws.Config, client *chatbot.Client) {
	input := &chatbot.DeleteMicrosoftTeamsUserIdentityInput{
		// ChatConfigurationArn: *string, // Required
		// UserId: *string, // Required
	}

	if len(_chatbotChatConfigurationArn) > 0 {
		input.ChatConfigurationArn = aws.String(_chatbotChatConfigurationArn)
	}
	if len(_chatbotUserId) > 0 {
		input.UserId = aws.String(_chatbotUserId)
	}

	if resp, err := client.DeleteMicrosoftTeamsUserIdentity(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a Slack channel configuration for AWS Chatbot
func chatbot_DeleteSlackChannelConfiguration(cfg aws.Config, client *chatbot.Client) {
	input := &chatbot.DeleteSlackChannelConfigurationInput{
		// ChatConfigurationArn: *string, // Required
	}

	if len(_chatbotChatConfigurationArn) > 0 {
		input.ChatConfigurationArn = aws.String(_chatbotChatConfigurationArn)
	}

	if resp, err := client.DeleteSlackChannelConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a user level permission for a Slack channel configuration.
func chatbot_DeleteSlackUserIdentity(cfg aws.Config, client *chatbot.Client) {
	input := &chatbot.DeleteSlackUserIdentityInput{
		// ChatConfigurationArn: *string, // Required
		// SlackTeamId: *string, // Required
		// SlackUserId: *string, // Required
	}

	if len(_chatbotChatConfigurationArn) > 0 {
		input.ChatConfigurationArn = aws.String(_chatbotChatConfigurationArn)
	}
	if len(_chatbotSlackTeamId) > 0 {
		input.SlackTeamId = aws.String(_chatbotSlackTeamId)
	}
	if len(_chatbotSlackUserId) > 0 {
		input.SlackUserId = aws.String(_chatbotSlackUserId)
	}

	if resp, err := client.DeleteSlackUserIdentity(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the Slack workspace authorization that allows channels to be configured
// in that workspace. This requires all configured channels in the workspace to be
// deleted.
func chatbot_DeleteSlackWorkspaceAuthorization(cfg aws.Config, client *chatbot.Client) {
	input := &chatbot.DeleteSlackWorkspaceAuthorizationInput{
		// SlackTeamId: *string, // Required
	}

	if len(_chatbotSlackTeamId) > 0 {
		input.SlackTeamId = aws.String(_chatbotSlackTeamId)
	}

	if resp, err := client.DeleteSlackWorkspaceAuthorization(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists Amazon Chime webhook configurations optionally filtered by
// ChatConfigurationArn
func chatbot_DescribeChimeWebhookConfigurations(cfg aws.Config, client *chatbot.Client) {
	input := &chatbot.DescribeChimeWebhookConfigurationsInput{}

	if len(_chatbotChatConfigurationArn) > 0 {
		input.ChatConfigurationArn = aws.String(_chatbotChatConfigurationArn)
	}
	if len(_chatbotMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _chatbotMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_chatbotNextToken) > 0 {
		input.NextToken = aws.String(_chatbotNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeChimeWebhookConfigurations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*chatbot.DescribeChimeWebhookConfigurationsOutput
	p := chatbot.NewDescribeChimeWebhookConfigurationsPaginator(client, input)
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

// Lists Slack channel configurations optionally filtered by ChatConfigurationArn
func chatbot_DescribeSlackChannelConfigurations(cfg aws.Config, client *chatbot.Client) {
	input := &chatbot.DescribeSlackChannelConfigurationsInput{}

	if len(_chatbotChatConfigurationArn) > 0 {
		input.ChatConfigurationArn = aws.String(_chatbotChatConfigurationArn)
	}
	if len(_chatbotMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _chatbotMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_chatbotNextToken) > 0 {
		input.NextToken = aws.String(_chatbotNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeSlackChannelConfigurations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*chatbot.DescribeSlackChannelConfigurationsOutput
	p := chatbot.NewDescribeSlackChannelConfigurationsPaginator(client, input)
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

// Lists all Slack user identities with a mapped role.
func chatbot_DescribeSlackUserIdentities(cfg aws.Config, client *chatbot.Client) {
	input := &chatbot.DescribeSlackUserIdentitiesInput{}

	if len(_chatbotChatConfigurationArn) > 0 {
		input.ChatConfigurationArn = aws.String(_chatbotChatConfigurationArn)
	}
	if len(_chatbotMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _chatbotMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_chatbotNextToken) > 0 {
		input.NextToken = aws.String(_chatbotNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeSlackUserIdentities(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*chatbot.DescribeSlackUserIdentitiesOutput
	p := chatbot.NewDescribeSlackUserIdentitiesPaginator(client, input)
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

// List all authorized Slack workspaces connected to the AWS Account onboarded
// with AWS Chatbot.
func chatbot_DescribeSlackWorkspaces(cfg aws.Config, client *chatbot.Client) {
	input := &chatbot.DescribeSlackWorkspacesInput{}

	if len(_chatbotMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _chatbotMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_chatbotNextToken) > 0 {
		input.NextToken = aws.String(_chatbotNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeSlackWorkspaces(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*chatbot.DescribeSlackWorkspacesOutput
	p := chatbot.NewDescribeSlackWorkspacesPaginator(client, input)
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

// Unlink a resource, for example a custom action, from a channel configuration.
func chatbot_DisassociateFromConfiguration(cfg aws.Config, client *chatbot.Client) {
	input := &chatbot.DisassociateFromConfigurationInput{
		// ChatConfiguration: *string, // Required
		// Resource: *string, // Required
	}

	if len(_chatbotChatConfiguration) > 0 {
		input.ChatConfiguration = aws.String(_chatbotChatConfiguration)
	}
	if len(_chatbotResource) > 0 {
		input.Resource = aws.String(_chatbotResource)
	}

	if resp, err := client.DisassociateFromConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns AWS Chatbot account preferences.
func chatbot_GetAccountPreferences(cfg aws.Config, client *chatbot.Client) {
	input := &chatbot.GetAccountPreferencesInput{}

	if resp, err := client.GetAccountPreferences(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a custom action.
func chatbot_GetCustomAction(cfg aws.Config, client *chatbot.Client) {
	input := &chatbot.GetCustomActionInput{
		// CustomActionArn: *string, // Required
	}

	if len(_chatbotCustomActionArn) > 0 {
		input.CustomActionArn = aws.String(_chatbotCustomActionArn)
	}

	if resp, err := client.GetCustomAction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a Microsoft Teams channel configuration in an AWS account.
func chatbot_GetMicrosoftTeamsChannelConfiguration(cfg aws.Config, client *chatbot.Client) {
	input := &chatbot.GetMicrosoftTeamsChannelConfigurationInput{
		// ChatConfigurationArn: *string, // Required
	}

	if len(_chatbotChatConfigurationArn) > 0 {
		input.ChatConfigurationArn = aws.String(_chatbotChatConfigurationArn)
	}

	if resp, err := client.GetMicrosoftTeamsChannelConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists resources associated with a channel configuration.
func chatbot_ListAssociations(cfg aws.Config, client *chatbot.Client) {
	input := &chatbot.ListAssociationsInput{
		// ChatConfiguration: *string, // Required
	}

	if len(_chatbotChatConfiguration) > 0 {
		input.ChatConfiguration = aws.String(_chatbotChatConfiguration)
	}
	if len(_chatbotMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _chatbotMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_chatbotNextToken) > 0 {
		input.NextToken = aws.String(_chatbotNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*chatbot.ListAssociationsOutput
	p := chatbot.NewListAssociationsPaginator(client, input)
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

// Lists custom actions defined in this account.
func chatbot_ListCustomActions(cfg aws.Config, client *chatbot.Client) {
	input := &chatbot.ListCustomActionsInput{}

	if len(_chatbotMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _chatbotMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_chatbotNextToken) > 0 {
		input.NextToken = aws.String(_chatbotNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCustomActions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*chatbot.ListCustomActionsOutput
	p := chatbot.NewListCustomActionsPaginator(client, input)
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

// Lists all AWS Chatbot Microsoft Teams channel configurations in an AWS account.
func chatbot_ListMicrosoftTeamsChannelConfigurations(cfg aws.Config, client *chatbot.Client) {
	input := &chatbot.ListMicrosoftTeamsChannelConfigurationsInput{}

	if len(_chatbotMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _chatbotMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_chatbotNextToken) > 0 {
		input.NextToken = aws.String(_chatbotNextToken)
	}
	if len(_chatbotTeamId) > 0 {
		input.TeamId = aws.String(_chatbotTeamId)
	}

	if disablePaginator() {
		if resp, err := client.ListMicrosoftTeamsChannelConfigurations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*chatbot.ListMicrosoftTeamsChannelConfigurationsOutput
	p := chatbot.NewListMicrosoftTeamsChannelConfigurationsPaginator(client, input)
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

// Lists all authorized Microsoft Teams for an AWS Account
func chatbot_ListMicrosoftTeamsConfiguredTeams(cfg aws.Config, client *chatbot.Client) {
	input := &chatbot.ListMicrosoftTeamsConfiguredTeamsInput{}

	if len(_chatbotMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _chatbotMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_chatbotNextToken) > 0 {
		input.NextToken = aws.String(_chatbotNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListMicrosoftTeamsConfiguredTeams(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*chatbot.ListMicrosoftTeamsConfiguredTeamsOutput
	p := chatbot.NewListMicrosoftTeamsConfiguredTeamsPaginator(client, input)
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

// A list all Microsoft Teams user identities with a mapped role.
func chatbot_ListMicrosoftTeamsUserIdentities(cfg aws.Config, client *chatbot.Client) {
	input := &chatbot.ListMicrosoftTeamsUserIdentitiesInput{}

	if len(_chatbotChatConfigurationArn) > 0 {
		input.ChatConfigurationArn = aws.String(_chatbotChatConfigurationArn)
	}
	if len(_chatbotMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _chatbotMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_chatbotNextToken) > 0 {
		input.NextToken = aws.String(_chatbotNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListMicrosoftTeamsUserIdentities(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*chatbot.ListMicrosoftTeamsUserIdentitiesOutput
	p := chatbot.NewListMicrosoftTeamsUserIdentitiesPaginator(client, input)
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

// Lists all of the tags associated with the Amazon Resource Name (ARN) that you
// specify. The resource can be a user, server, or role.
func chatbot_ListTagsForResource(cfg aws.Config, client *chatbot.Client) {
	input := &chatbot.ListTagsForResourceInput{
		// ResourceARN: *string, // Required
	}

	if len(_chatbotResourceARN) > 0 {
		input.ResourceARN = aws.String(_chatbotResourceARN)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Attaches a key-value pair to a resource, as identified by its Amazon Resource
// Name (ARN). Resources are users, servers, roles, and other entities.
func chatbot_TagResource(cfg aws.Config, client *chatbot.Client) {
	input := &chatbot.TagResourceInput{
		// ResourceARN: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_chatbotResourceARN) > 0 {
		input.ResourceARN = aws.String(_chatbotResourceARN)
	}
	if len(_chatbotTags) > 0 {
		if err := assignInputField(input, "Tags", _chatbotTags); err != nil {
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

// Detaches a key-value pair from a resource, as identified by its Amazon Resource
// Name (ARN). Resources are users, servers, roles, and other entities.
func chatbot_UntagResource(cfg aws.Config, client *chatbot.Client) {
	input := &chatbot.UntagResourceInput{
		// ResourceARN: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_chatbotResourceARN) > 0 {
		input.ResourceARN = aws.String(_chatbotResourceARN)
	}
	if len(_chatbotTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _chatbotTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates AWS Chatbot account preferences.
func chatbot_UpdateAccountPreferences(cfg aws.Config, client *chatbot.Client) {
	input := &chatbot.UpdateAccountPreferencesInput{}

	if len(_chatbotTrainingDataCollectionEnabled) > 0 {
		if err := assignInputField(input, "TrainingDataCollectionEnabled", _chatbotTrainingDataCollectionEnabled); err != nil {
			log.Errorf("invalid --training-data-collection-enabled: %s", err.Error())
			return
		}
	}
	if len(_chatbotUserAuthorizationRequired) > 0 {
		if err := assignInputField(input, "UserAuthorizationRequired", _chatbotUserAuthorizationRequired); err != nil {
			log.Errorf("invalid --user-authorization-required: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateAccountPreferences(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a Amazon Chime webhook configuration.
func chatbot_UpdateChimeWebhookConfiguration(cfg aws.Config, client *chatbot.Client) {
	input := &chatbot.UpdateChimeWebhookConfigurationInput{
		// ChatConfigurationArn: *string, // Required
	}

	if len(_chatbotChatConfigurationArn) > 0 {
		input.ChatConfigurationArn = aws.String(_chatbotChatConfigurationArn)
	}
	if len(_chatbotIamRoleArn) > 0 {
		input.IamRoleArn = aws.String(_chatbotIamRoleArn)
	}
	if len(_chatbotLoggingLevel) > 0 {
		input.LoggingLevel = aws.String(_chatbotLoggingLevel)
	}
	if len(_chatbotSnsTopicArns) > 0 {
		input.SnsTopicArns = append([]string(nil), _chatbotSnsTopicArns...)
	}
	if len(_chatbotWebhookDescription) > 0 {
		input.WebhookDescription = aws.String(_chatbotWebhookDescription)
	}
	if len(_chatbotWebhookUrl) > 0 {
		input.WebhookUrl = aws.String(_chatbotWebhookUrl)
	}

	if resp, err := client.UpdateChimeWebhookConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a custom action.
func chatbot_UpdateCustomAction(cfg aws.Config, client *chatbot.Client) {
	input := &chatbot.UpdateCustomActionInput{
		// CustomActionArn: *string, // Required
		// Definition: *types.CustomActionDefinition, // Required
	}

	if len(_chatbotCustomActionArn) > 0 {
		input.CustomActionArn = aws.String(_chatbotCustomActionArn)
	}
	if len(_chatbotDefinition) > 0 {
		if err := assignInputField(input, "Definition", _chatbotDefinition); err != nil {
			log.Errorf("invalid --definition: %s", err.Error())
			return
		}
	}
	if len(_chatbotAliasName) > 0 {
		input.AliasName = aws.String(_chatbotAliasName)
	}
	if len(_chatbotAttachments) > 0 {
		if err := assignInputField(input, "Attachments", _chatbotAttachments); err != nil {
			log.Errorf("invalid --attachments: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateCustomAction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an Microsoft Teams channel configuration.
func chatbot_UpdateMicrosoftTeamsChannelConfiguration(cfg aws.Config, client *chatbot.Client) {
	input := &chatbot.UpdateMicrosoftTeamsChannelConfigurationInput{
		// ChannelId: *string, // Required
		// ChatConfigurationArn: *string, // Required
	}

	if len(_chatbotChannelId) > 0 {
		input.ChannelId = aws.String(_chatbotChannelId)
	}
	if len(_chatbotChatConfigurationArn) > 0 {
		input.ChatConfigurationArn = aws.String(_chatbotChatConfigurationArn)
	}
	if len(_chatbotChannelName) > 0 {
		input.ChannelName = aws.String(_chatbotChannelName)
	}
	if len(_chatbotGuardrailPolicyArns) > 0 {
		input.GuardrailPolicyArns = append([]string(nil), _chatbotGuardrailPolicyArns...)
	}
	if len(_chatbotIamRoleArn) > 0 {
		input.IamRoleArn = aws.String(_chatbotIamRoleArn)
	}
	if len(_chatbotLoggingLevel) > 0 {
		input.LoggingLevel = aws.String(_chatbotLoggingLevel)
	}
	if len(_chatbotSnsTopicArns) > 0 {
		input.SnsTopicArns = append([]string(nil), _chatbotSnsTopicArns...)
	}
	if len(_chatbotUserAuthorizationRequired) > 0 {
		if err := assignInputField(input, "UserAuthorizationRequired", _chatbotUserAuthorizationRequired); err != nil {
			log.Errorf("invalid --user-authorization-required: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateMicrosoftTeamsChannelConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a Slack channel configuration.
func chatbot_UpdateSlackChannelConfiguration(cfg aws.Config, client *chatbot.Client) {
	input := &chatbot.UpdateSlackChannelConfigurationInput{
		// ChatConfigurationArn: *string, // Required
		// SlackChannelId: *string, // Required
	}

	if len(_chatbotChatConfigurationArn) > 0 {
		input.ChatConfigurationArn = aws.String(_chatbotChatConfigurationArn)
	}
	if len(_chatbotSlackChannelId) > 0 {
		input.SlackChannelId = aws.String(_chatbotSlackChannelId)
	}
	if len(_chatbotGuardrailPolicyArns) > 0 {
		input.GuardrailPolicyArns = append([]string(nil), _chatbotGuardrailPolicyArns...)
	}
	if len(_chatbotIamRoleArn) > 0 {
		input.IamRoleArn = aws.String(_chatbotIamRoleArn)
	}
	if len(_chatbotLoggingLevel) > 0 {
		input.LoggingLevel = aws.String(_chatbotLoggingLevel)
	}
	if len(_chatbotSlackChannelName) > 0 {
		input.SlackChannelName = aws.String(_chatbotSlackChannelName)
	}
	if len(_chatbotSnsTopicArns) > 0 {
		input.SnsTopicArns = append([]string(nil), _chatbotSnsTopicArns...)
	}
	if len(_chatbotUserAuthorizationRequired) > 0 {
		if err := assignInputField(input, "UserAuthorizationRequired", _chatbotUserAuthorizationRequired); err != nil {
			log.Errorf("invalid --user-authorization-required: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateSlackChannelConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_chatbotCmd)
	_chatbotCmd.Flags().SortFlags = false

	_chatbotCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_chatbotCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_chatbotCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_chatbotCmd.Flags().StringVarP(&_chatbotActionName, "action-name", "", "", "Action Name")
	_chatbotCmd.Flags().StringVarP(&_chatbotAliasName, "alias-name", "", "", "Alias Name")
	_chatbotCmd.Flags().StringVarP(&_chatbotAttachments, "attachments", "", "", "Attachments")
	_chatbotCmd.Flags().StringVarP(&_chatbotChannelId, "channel-id", "", "", "Channel ID")
	_chatbotCmd.Flags().StringVarP(&_chatbotChannelName, "channel-name", "", "", "Channel Name")
	_chatbotCmd.Flags().StringVarP(&_chatbotChatConfiguration, "chat-configuration", "", "", "Chat Configuration")
	_chatbotCmd.Flags().StringVarP(&_chatbotChatConfigurationArn, "chat-configuration-arn", "", "", "Chat Configuration ARN")
	_chatbotCmd.Flags().StringVarP(&_chatbotClientToken, "client-token", "", "", "Client Token")
	_chatbotCmd.Flags().StringVarP(&_chatbotConfigurationName, "configuration-name", "", "", "Configuration Name")
	_chatbotCmd.Flags().StringVarP(&_chatbotCustomActionArn, "custom-action-arn", "", "", "Custom Action ARN")
	_chatbotCmd.Flags().StringVarP(&_chatbotDefinition, "definition", "", "", "Definition")
	_chatbotCmd.Flags().StringSliceVarP(&_chatbotGuardrailPolicyArns, "guardrail-policy-arns", "", nil, "Guardrail Policy Arns")
	_chatbotCmd.Flags().StringVarP(&_chatbotIamRoleArn, "iam-role-arn", "", "", "IAM Role ARN")
	_chatbotCmd.Flags().StringVarP(&_chatbotLoggingLevel, "logging-level", "", "", "Logging Level")
	_chatbotCmd.Flags().StringVarP(&_chatbotMaxResults, "max-results", "", "", "Max Results")
	_chatbotCmd.Flags().StringVarP(&_chatbotNextToken, "next-token", "", "", "Next Token")
	_chatbotCmd.Flags().StringVarP(&_chatbotResource, "resource", "", "", "Resource")
	_chatbotCmd.Flags().StringVarP(&_chatbotResourceARN, "resource-arn", "", "", "Resource ARN")
	_chatbotCmd.Flags().StringVarP(&_chatbotSlackChannelId, "slack-channel-id", "", "", "Slack Channel ID")
	_chatbotCmd.Flags().StringVarP(&_chatbotSlackChannelName, "slack-channel-name", "", "", "Slack Channel Name")
	_chatbotCmd.Flags().StringVarP(&_chatbotSlackTeamId, "slack-team-id", "", "", "Slack Team ID")
	_chatbotCmd.Flags().StringVarP(&_chatbotSlackUserId, "slack-user-id", "", "", "Slack User ID")
	_chatbotCmd.Flags().StringSliceVarP(&_chatbotSnsTopicArns, "sns-topic-arns", "", nil, "SNS Topic Arns")
	_chatbotCmd.Flags().StringSliceVarP(&_chatbotTagKeys, "tag-keys", "", nil, "Tag Keys")
	_chatbotCmd.Flags().StringVarP(&_chatbotTags, "tags", "", "", "Tags")
	_chatbotCmd.Flags().StringVarP(&_chatbotTeamId, "team-id", "", "", "Team ID")
	_chatbotCmd.Flags().StringVarP(&_chatbotTeamName, "team-name", "", "", "Team Name")
	_chatbotCmd.Flags().StringVarP(&_chatbotTenantId, "tenant-id", "", "", "Tenant ID")
	_chatbotCmd.Flags().StringVarP(&_chatbotTrainingDataCollectionEnabled, "training-data-collection-enabled", "", "", "Training Data Collection Enabled")
	_chatbotCmd.Flags().StringVarP(&_chatbotUserAuthorizationRequired, "user-authorization-required", "", "", "User Authorization Required")
	_chatbotCmd.Flags().StringVarP(&_chatbotUserId, "user-id", "", "", "User ID")
	_chatbotCmd.Flags().StringVarP(&_chatbotWebhookDescription, "webhook-description", "", "", "Webhook Description")
	_chatbotCmd.Flags().StringVarP(&_chatbotWebhookUrl, "webhook-url", "", "", "Webhook URL")

	_chatbotCmd.Flags().BoolVarP(&_chatbotAssociateToConfiguration, "associate-to-configuration", "", false, "Associate To Configuration")
	_chatbotCmd.Flags().BoolVarP(&_chatbotCreateChimeWebhookConfiguration, "create-chime-webhook-configuration", "", false, "Create Chime Webhook Configuration")
	_chatbotCmd.Flags().BoolVarP(&_chatbotCreateCustomAction, "create-custom-action", "", false, "Create Custom Action")
	_chatbotCmd.Flags().BoolVarP(&_chatbotCreateMicrosoftTeamsChannelConfiguration, "create-microsoft-teams-channel-configuration", "", false, "Create Microsoft Teams Channel Configuration")
	_chatbotCmd.Flags().BoolVarP(&_chatbotCreateSlackChannelConfiguration, "create-slack-channel-configuration", "", false, "Create Slack Channel Configuration")
	_chatbotCmd.Flags().BoolVarP(&_chatbotDeleteChimeWebhookConfiguration, "delete-chime-webhook-configuration", "", false, "Delete Chime Webhook Configuration")
	_chatbotCmd.Flags().BoolVarP(&_chatbotDeleteCustomAction, "delete-custom-action", "", false, "Delete Custom Action")
	_chatbotCmd.Flags().BoolVarP(&_chatbotDeleteMicrosoftTeamsChannelConfiguration, "delete-microsoft-teams-channel-configuration", "", false, "Delete Microsoft Teams Channel Configuration")
	_chatbotCmd.Flags().BoolVarP(&_chatbotDeleteMicrosoftTeamsConfiguredTeam, "delete-microsoft-teams-configured-team", "", false, "Delete Microsoft Teams Configured Team")
	_chatbotCmd.Flags().BoolVarP(&_chatbotDeleteMicrosoftTeamsUserIdentity, "delete-microsoft-teams-user-identity", "", false, "Delete Microsoft Teams User Identity")
	_chatbotCmd.Flags().BoolVarP(&_chatbotDeleteSlackChannelConfiguration, "delete-slack-channel-configuration", "", false, "Delete Slack Channel Configuration")
	_chatbotCmd.Flags().BoolVarP(&_chatbotDeleteSlackUserIdentity, "delete-slack-user-identity", "", false, "Delete Slack User Identity")
	_chatbotCmd.Flags().BoolVarP(&_chatbotDeleteSlackWorkspaceAuthorization, "delete-slack-workspace-authorization", "", false, "Delete Slack Workspace Authorization")
	_chatbotCmd.Flags().BoolVarP(&_chatbotDescribeChimeWebhookConfigurations, "describe-chime-webhook-configurations", "", false, "Describe Chime Webhook Configurations")
	_chatbotCmd.Flags().BoolVarP(&_chatbotDescribeSlackChannelConfigurations, "describe-slack-channel-configurations", "", false, "Describe Slack Channel Configurations")
	_chatbotCmd.Flags().BoolVarP(&_chatbotDescribeSlackUserIdentities, "describe-slack-user-identities", "", false, "Describe Slack User Identities")
	_chatbotCmd.Flags().BoolVarP(&_chatbotDescribeSlackWorkspaces, "describe-slack-workspaces", "", false, "Describe Slack Workspaces")
	_chatbotCmd.Flags().BoolVarP(&_chatbotDisassociateFromConfiguration, "disassociate-from-configuration", "", false, "Disassociate From Configuration")
	_chatbotCmd.Flags().BoolVarP(&_chatbotGetAccountPreferences, "get-account-preferences", "", false, "Get Account Preferences")
	_chatbotCmd.Flags().BoolVarP(&_chatbotGetCustomAction, "get-custom-action", "", false, "Get Custom Action")
	_chatbotCmd.Flags().BoolVarP(&_chatbotGetMicrosoftTeamsChannelConfiguration, "get-microsoft-teams-channel-configuration", "", false, "Get Microsoft Teams Channel Configuration")
	_chatbotCmd.Flags().BoolVarP(&_chatbotListAssociations, "list-associations", "", false, "List Associations")
	_chatbotCmd.Flags().BoolVarP(&_chatbotListCustomActions, "list-custom-actions", "", false, "List Custom Actions")
	_chatbotCmd.Flags().BoolVarP(&_chatbotListMicrosoftTeamsChannelConfigurations, "list-microsoft-teams-channel-configurations", "", false, "List Microsoft Teams Channel Configurations")
	_chatbotCmd.Flags().BoolVarP(&_chatbotListMicrosoftTeamsConfiguredTeams, "list-microsoft-teams-configured-teams", "", false, "List Microsoft Teams Configured Teams")
	_chatbotCmd.Flags().BoolVarP(&_chatbotListMicrosoftTeamsUserIdentities, "list-microsoft-teams-user-identities", "", false, "List Microsoft Teams User Identities")
	_chatbotCmd.Flags().BoolVarP(&_chatbotListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_chatbotCmd.Flags().BoolVarP(&_chatbotTagResource, "tag-resource", "", false, "Tag Resource")
	_chatbotCmd.Flags().BoolVarP(&_chatbotUntagResource, "untag-resource", "", false, "Untag Resource")
	_chatbotCmd.Flags().BoolVarP(&_chatbotUpdateAccountPreferences, "update-account-preferences", "", false, "Update Account Preferences")
	_chatbotCmd.Flags().BoolVarP(&_chatbotUpdateChimeWebhookConfiguration, "update-chime-webhook-configuration", "", false, "Update Chime Webhook Configuration")
	_chatbotCmd.Flags().BoolVarP(&_chatbotUpdateCustomAction, "update-custom-action", "", false, "Update Custom Action")
	_chatbotCmd.Flags().BoolVarP(&_chatbotUpdateMicrosoftTeamsChannelConfiguration, "update-microsoft-teams-channel-configuration", "", false, "Update Microsoft Teams Channel Configuration")
	_chatbotCmd.Flags().BoolVarP(&_chatbotUpdateSlackChannelConfiguration, "update-slack-channel-configuration", "", false, "Update Slack Channel Configuration")

}
