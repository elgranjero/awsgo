package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/supportapp"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// supportappCmd represents the supportapp command
var _supportappCmd = &cobra.Command{
	Use:   "supportapp",
	Short: "AWS supportapp CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := supportapp.NewFromConfig(cfg)
		if _supportappCreateSlackChannelConfiguration {
			supportapp_CreateSlackChannelConfiguration(cfg, client)
			return
		}
		if _supportappDeleteAccountAlias {
			supportapp_DeleteAccountAlias(cfg, client)
			return
		}
		if _supportappDeleteSlackChannelConfiguration {
			supportapp_DeleteSlackChannelConfiguration(cfg, client)
			return
		}
		if _supportappDeleteSlackWorkspaceConfiguration {
			supportapp_DeleteSlackWorkspaceConfiguration(cfg, client)
			return
		}
		if _supportappGetAccountAlias {
			supportapp_GetAccountAlias(cfg, client)
			return
		}
		if _supportappListSlackChannelConfigurations {
			supportapp_ListSlackChannelConfigurations(cfg, client)
			return
		}
		if _supportappListSlackWorkspaceConfigurations {
			supportapp_ListSlackWorkspaceConfigurations(cfg, client)
			return
		}
		if _supportappPutAccountAlias {
			supportapp_PutAccountAlias(cfg, client)
			return
		}
		if _supportappRegisterSlackWorkspaceForOrganization {
			supportapp_RegisterSlackWorkspaceForOrganization(cfg, client)
			return
		}
		if _supportappUpdateSlackChannelConfiguration {
			supportapp_UpdateSlackChannelConfiguration(cfg, client)
			return
		}

	},
}

var (
	_supportappCreateSlackChannelConfiguration       bool
	_supportappDeleteAccountAlias                    bool
	_supportappDeleteSlackChannelConfiguration       bool
	_supportappDeleteSlackWorkspaceConfiguration     bool
	_supportappGetAccountAlias                       bool
	_supportappListSlackChannelConfigurations        bool
	_supportappListSlackWorkspaceConfigurations      bool
	_supportappPutAccountAlias                       bool
	_supportappRegisterSlackWorkspaceForOrganization bool
	_supportappUpdateSlackChannelConfiguration       bool

	_supportappAccountAlias                    string
	_supportappChannelId                       string
	_supportappChannelName                     string
	_supportappChannelRoleArn                  string
	_supportappNextToken                       string
	_supportappNotifyOnAddCorrespondenceToCase string
	_supportappNotifyOnCaseSeverity            string
	_supportappNotifyOnCreateOrReopenCase      string
	_supportappNotifyOnResolveCase             string
	_supportappTeamId                          string
)

// Creates a Slack channel configuration for your Amazon Web Services account.
// - You can add up to 5 Slack workspaces for your account.
//
// - You can add up to 20 Slack channels for your account.
//
// A Slack channel can have up to 100 Amazon Web Services accounts. This means
// that only 100 accounts can add the same Slack channel to the Amazon Web Services
// Support App. We recommend that you only add the accounts that you need to manage
// support cases for your organization. This can reduce the notifications about
// case updates that you receive in the Slack channel.
//
// We recommend that you choose a private Slack channel so that only members in
// that channel have read and write access to your support cases. Anyone in your
// Slack channel can create, update, or resolve support cases for your account.
// Users require an invitation to join private channels.
func supportapp_CreateSlackChannelConfiguration(cfg aws.Config, client *supportapp.Client) {
	input := &supportapp.CreateSlackChannelConfigurationInput{
		// ChannelId: *string, // Required
		// ChannelRoleArn: *string, // Required
		// NotifyOnCaseSeverity: types.NotificationSeverityLevel, // Required
		// TeamId: *string, // Required
	}

	if len(_supportappChannelId) > 0 {
		input.ChannelId = aws.String(_supportappChannelId)
	}
	if len(_supportappChannelRoleArn) > 0 {
		input.ChannelRoleArn = aws.String(_supportappChannelRoleArn)
	}
	if len(_supportappNotifyOnCaseSeverity) > 0 {
		if err := assignInputField(input, "NotifyOnCaseSeverity", _supportappNotifyOnCaseSeverity); err != nil {
			log.Errorf("invalid --notify-on-case-severity: %s", err.Error())
			return
		}
	}
	if len(_supportappTeamId) > 0 {
		input.TeamId = aws.String(_supportappTeamId)
	}
	if len(_supportappChannelName) > 0 {
		input.ChannelName = aws.String(_supportappChannelName)
	}
	if len(_supportappNotifyOnAddCorrespondenceToCase) > 0 {
		if err := assignInputField(input, "NotifyOnAddCorrespondenceToCase", _supportappNotifyOnAddCorrespondenceToCase); err != nil {
			log.Errorf("invalid --notify-on-add-correspondence-to-case: %s", err.Error())
			return
		}
	}
	if len(_supportappNotifyOnCreateOrReopenCase) > 0 {
		if err := assignInputField(input, "NotifyOnCreateOrReopenCase", _supportappNotifyOnCreateOrReopenCase); err != nil {
			log.Errorf("invalid --notify-on-create-or-reopen-case: %s", err.Error())
			return
		}
	}
	if len(_supportappNotifyOnResolveCase) > 0 {
		if err := assignInputField(input, "NotifyOnResolveCase", _supportappNotifyOnResolveCase); err != nil {
			log.Errorf("invalid --notify-on-resolve-case: %s", err.Error())
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

// Deletes an alias for an Amazon Web Services account ID. The alias appears in
// the Amazon Web Services Support App page of the Amazon Web Services Support
// Center. The alias also appears in Slack messages from the Amazon Web Services
// Support App.
func supportapp_DeleteAccountAlias(cfg aws.Config, client *supportapp.Client) {
	input := &supportapp.DeleteAccountAliasInput{}

	if resp, err := client.DeleteAccountAlias(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a Slack channel configuration from your Amazon Web Services account.
// This operation doesn't delete your Slack channel.
func supportapp_DeleteSlackChannelConfiguration(cfg aws.Config, client *supportapp.Client) {
	input := &supportapp.DeleteSlackChannelConfigurationInput{
		// ChannelId: *string, // Required
		// TeamId: *string, // Required
	}

	if len(_supportappChannelId) > 0 {
		input.ChannelId = aws.String(_supportappChannelId)
	}
	if len(_supportappTeamId) > 0 {
		input.TeamId = aws.String(_supportappTeamId)
	}

	if resp, err := client.DeleteSlackChannelConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a Slack workspace configuration from your Amazon Web Services account.
// This operation doesn't delete your Slack workspace.
func supportapp_DeleteSlackWorkspaceConfiguration(cfg aws.Config, client *supportapp.Client) {
	input := &supportapp.DeleteSlackWorkspaceConfigurationInput{
		// TeamId: *string, // Required
	}

	if len(_supportappTeamId) > 0 {
		input.TeamId = aws.String(_supportappTeamId)
	}

	if resp, err := client.DeleteSlackWorkspaceConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the alias from an Amazon Web Services account ID. The alias appears
// in the Amazon Web Services Support App page of the Amazon Web Services Support
// Center. The alias also appears in Slack messages from the Amazon Web Services
// Support App.
func supportapp_GetAccountAlias(cfg aws.Config, client *supportapp.Client) {
	input := &supportapp.GetAccountAliasInput{}

	if resp, err := client.GetAccountAlias(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the Slack channel configurations for an Amazon Web Services account.
func supportapp_ListSlackChannelConfigurations(cfg aws.Config, client *supportapp.Client) {
	input := &supportapp.ListSlackChannelConfigurationsInput{}

	if len(_supportappNextToken) > 0 {
		input.NextToken = aws.String(_supportappNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSlackChannelConfigurations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*supportapp.ListSlackChannelConfigurationsOutput
	p := supportapp.NewListSlackChannelConfigurationsPaginator(client, input)
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

// Lists the Slack workspace configurations for an Amazon Web Services account.
func supportapp_ListSlackWorkspaceConfigurations(cfg aws.Config, client *supportapp.Client) {
	input := &supportapp.ListSlackWorkspaceConfigurationsInput{}

	if len(_supportappNextToken) > 0 {
		input.NextToken = aws.String(_supportappNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSlackWorkspaceConfigurations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*supportapp.ListSlackWorkspaceConfigurationsOutput
	p := supportapp.NewListSlackWorkspaceConfigurationsPaginator(client, input)
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

// Creates or updates an individual alias for each Amazon Web Services account ID.
// The alias appears in the Amazon Web Services Support App page of the Amazon Web
// Services Support Center. The alias also appears in Slack messages from the
// Amazon Web Services Support App.
func supportapp_PutAccountAlias(cfg aws.Config, client *supportapp.Client) {
	input := &supportapp.PutAccountAliasInput{
		// AccountAlias: *string, // Required
	}

	if len(_supportappAccountAlias) > 0 {
		input.AccountAlias = aws.String(_supportappAccountAlias)
	}

	if resp, err := client.PutAccountAlias(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Registers a Slack workspace for your Amazon Web Services account. To call this
// API, your account must be part of an organization in Organizations.
//
// If you're the management account and you want to register Slack workspaces for
// your organization, you must complete the following tasks:
//
// - Sign in to the [Amazon Web Services Support Center]and authorize the Slack workspaces where you want your
// organization to have access to. See [Authorize a Slack workspace]in the Amazon Web Services Support User
// Guide.
//
// - Call the RegisterSlackWorkspaceForOrganization API to authorize each Slack
// workspace for the organization.
//
// After the management account authorizes the Slack workspace, member accounts
// can call this API to authorize the same Slack workspace for their individual
// accounts. Member accounts don't need to authorize the Slack workspace manually
// through the [Amazon Web Services Support Center].
//
// To use the Amazon Web Services Support App, each account must then complete the
// following tasks:
//
// - Create an Identity and Access Management (IAM) role with the required
// permission. For more information, see [Managing access to the Amazon Web Services Support App].
//
// - Configure a Slack channel to use the Amazon Web Services Support App for
// support cases for that account. For more information, see [Configuring a Slack channel].
//
// [Authorize a Slack workspace]: https://docs.aws.amazon.com/awssupport/latest/user/authorize-slack-workspace.html
// [Managing access to the Amazon Web Services Support App]: https://docs.aws.amazon.com/awssupport/latest/user/support-app-permissions.html
// [Amazon Web Services Support Center]: https://console.aws.amazon.com/support/app
// [Configuring a Slack channel]: https://docs.aws.amazon.com/awssupport/latest/user/add-your-slack-channel.html
func supportapp_RegisterSlackWorkspaceForOrganization(cfg aws.Config, client *supportapp.Client) {
	input := &supportapp.RegisterSlackWorkspaceForOrganizationInput{
		// TeamId: *string, // Required
	}

	if len(_supportappTeamId) > 0 {
		input.TeamId = aws.String(_supportappTeamId)
	}

	if resp, err := client.RegisterSlackWorkspaceForOrganization(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration for a Slack channel, such as case update
// notifications.
func supportapp_UpdateSlackChannelConfiguration(cfg aws.Config, client *supportapp.Client) {
	input := &supportapp.UpdateSlackChannelConfigurationInput{
		// ChannelId: *string, // Required
		// TeamId: *string, // Required
	}

	if len(_supportappChannelId) > 0 {
		input.ChannelId = aws.String(_supportappChannelId)
	}
	if len(_supportappTeamId) > 0 {
		input.TeamId = aws.String(_supportappTeamId)
	}
	if len(_supportappChannelName) > 0 {
		input.ChannelName = aws.String(_supportappChannelName)
	}
	if len(_supportappChannelRoleArn) > 0 {
		input.ChannelRoleArn = aws.String(_supportappChannelRoleArn)
	}
	if len(_supportappNotifyOnAddCorrespondenceToCase) > 0 {
		if err := assignInputField(input, "NotifyOnAddCorrespondenceToCase", _supportappNotifyOnAddCorrespondenceToCase); err != nil {
			log.Errorf("invalid --notify-on-add-correspondence-to-case: %s", err.Error())
			return
		}
	}
	if len(_supportappNotifyOnCaseSeverity) > 0 {
		if err := assignInputField(input, "NotifyOnCaseSeverity", _supportappNotifyOnCaseSeverity); err != nil {
			log.Errorf("invalid --notify-on-case-severity: %s", err.Error())
			return
		}
	}
	if len(_supportappNotifyOnCreateOrReopenCase) > 0 {
		if err := assignInputField(input, "NotifyOnCreateOrReopenCase", _supportappNotifyOnCreateOrReopenCase); err != nil {
			log.Errorf("invalid --notify-on-create-or-reopen-case: %s", err.Error())
			return
		}
	}
	if len(_supportappNotifyOnResolveCase) > 0 {
		if err := assignInputField(input, "NotifyOnResolveCase", _supportappNotifyOnResolveCase); err != nil {
			log.Errorf("invalid --notify-on-resolve-case: %s", err.Error())
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
	_rootCmd.AddCommand(_supportappCmd)
	_supportappCmd.Flags().SortFlags = false

	_supportappCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_supportappCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_supportappCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_supportappCmd.Flags().StringVarP(&_supportappAccountAlias, "account-alias", "", "", "Account Alias")
	_supportappCmd.Flags().StringVarP(&_supportappChannelId, "channel-id", "", "", "Channel ID")
	_supportappCmd.Flags().StringVarP(&_supportappChannelName, "channel-name", "", "", "Channel Name")
	_supportappCmd.Flags().StringVarP(&_supportappChannelRoleArn, "channel-role-arn", "", "", "Channel Role ARN")
	_supportappCmd.Flags().StringVarP(&_supportappNextToken, "next-token", "", "", "Next Token")
	_supportappCmd.Flags().StringVarP(&_supportappNotifyOnAddCorrespondenceToCase, "notify-on-add-correspondence-to-case", "", "", "Notify On Add Correspondence To Case")
	_supportappCmd.Flags().StringVarP(&_supportappNotifyOnCaseSeverity, "notify-on-case-severity", "", "", "Notify On Case Severity")
	_supportappCmd.Flags().StringVarP(&_supportappNotifyOnCreateOrReopenCase, "notify-on-create-or-reopen-case", "", "", "Notify On Create Or Reopen Case")
	_supportappCmd.Flags().StringVarP(&_supportappNotifyOnResolveCase, "notify-on-resolve-case", "", "", "Notify On Resolve Case")
	_supportappCmd.Flags().StringVarP(&_supportappTeamId, "team-id", "", "", "Team ID")

	_supportappCmd.Flags().BoolVarP(&_supportappCreateSlackChannelConfiguration, "create-slack-channel-configuration", "", false, "Create Slack Channel Configuration")
	_supportappCmd.Flags().BoolVarP(&_supportappDeleteAccountAlias, "delete-account-alias", "", false, "Delete Account Alias")
	_supportappCmd.Flags().BoolVarP(&_supportappDeleteSlackChannelConfiguration, "delete-slack-channel-configuration", "", false, "Delete Slack Channel Configuration")
	_supportappCmd.Flags().BoolVarP(&_supportappDeleteSlackWorkspaceConfiguration, "delete-slack-workspace-configuration", "", false, "Delete Slack Workspace Configuration")
	_supportappCmd.Flags().BoolVarP(&_supportappGetAccountAlias, "get-account-alias", "", false, "Get Account Alias")
	_supportappCmd.Flags().BoolVarP(&_supportappListSlackChannelConfigurations, "list-slack-channel-configurations", "", false, "List Slack Channel Configurations")
	_supportappCmd.Flags().BoolVarP(&_supportappListSlackWorkspaceConfigurations, "list-slack-workspace-configurations", "", false, "List Slack Workspace Configurations")
	_supportappCmd.Flags().BoolVarP(&_supportappPutAccountAlias, "put-account-alias", "", false, "Put Account Alias")
	_supportappCmd.Flags().BoolVarP(&_supportappRegisterSlackWorkspaceForOrganization, "register-slack-workspace-for-organization", "", false, "Register Slack Workspace For Organization")
	_supportappCmd.Flags().BoolVarP(&_supportappUpdateSlackChannelConfiguration, "update-slack-channel-configuration", "", false, "Update Slack Channel Configuration")

}
