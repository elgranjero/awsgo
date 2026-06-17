package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/wickr"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// wickrCmd represents the wickr command
var _wickrCmd = &cobra.Command{
	Use:   "wickr",
	Short: "AWS wickr CLI",
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
		client := wickr.NewFromConfig(cfg)
		if _wickrBatchCreateUser {
			wickr_BatchCreateUser(cfg, client)
			return
		}
		if _wickrBatchDeleteUser {
			wickr_BatchDeleteUser(cfg, client)
			return
		}
		if _wickrBatchLookupUserUname {
			wickr_BatchLookupUserUname(cfg, client)
			return
		}
		if _wickrBatchReinviteUser {
			wickr_BatchReinviteUser(cfg, client)
			return
		}
		if _wickrBatchResetDevicesForUser {
			wickr_BatchResetDevicesForUser(cfg, client)
			return
		}
		if _wickrBatchToggleUserSuspendStatus {
			wickr_BatchToggleUserSuspendStatus(cfg, client)
			return
		}
		if _wickrCreateBot {
			wickr_CreateBot(cfg, client)
			return
		}
		if _wickrCreateDataRetentionBot {
			wickr_CreateDataRetentionBot(cfg, client)
			return
		}
		if _wickrCreateDataRetentionBotChallenge {
			wickr_CreateDataRetentionBotChallenge(cfg, client)
			return
		}
		if _wickrCreateNetwork {
			wickr_CreateNetwork(cfg, client)
			return
		}
		if _wickrCreateSecurityGroup {
			wickr_CreateSecurityGroup(cfg, client)
			return
		}
		if _wickrDeleteBot {
			wickr_DeleteBot(cfg, client)
			return
		}
		if _wickrDeleteDataRetentionBot {
			wickr_DeleteDataRetentionBot(cfg, client)
			return
		}
		if _wickrDeleteNetwork {
			wickr_DeleteNetwork(cfg, client)
			return
		}
		if _wickrDeleteSecurityGroup {
			wickr_DeleteSecurityGroup(cfg, client)
			return
		}
		if _wickrGetBot {
			wickr_GetBot(cfg, client)
			return
		}
		if _wickrGetBotsCount {
			wickr_GetBotsCount(cfg, client)
			return
		}
		if _wickrGetDataRetentionBot {
			wickr_GetDataRetentionBot(cfg, client)
			return
		}
		if _wickrGetGuestUserHistoryCount {
			wickr_GetGuestUserHistoryCount(cfg, client)
			return
		}
		if _wickrGetNetwork {
			wickr_GetNetwork(cfg, client)
			return
		}
		if _wickrGetNetworkSettings {
			wickr_GetNetworkSettings(cfg, client)
			return
		}
		if _wickrGetOidcInfo {
			wickr_GetOidcInfo(cfg, client)
			return
		}
		if _wickrGetOpentdfConfig {
			wickr_GetOpentdfConfig(cfg, client)
			return
		}
		if _wickrGetSecurityGroup {
			wickr_GetSecurityGroup(cfg, client)
			return
		}
		if _wickrGetUser {
			wickr_GetUser(cfg, client)
			return
		}
		if _wickrGetUsersCount {
			wickr_GetUsersCount(cfg, client)
			return
		}
		if _wickrListBlockedGuestUsers {
			wickr_ListBlockedGuestUsers(cfg, client)
			return
		}
		if _wickrListBots {
			wickr_ListBots(cfg, client)
			return
		}
		if _wickrListDevicesForUser {
			wickr_ListDevicesForUser(cfg, client)
			return
		}
		if _wickrListGuestUsers {
			wickr_ListGuestUsers(cfg, client)
			return
		}
		if _wickrListNetworks {
			wickr_ListNetworks(cfg, client)
			return
		}
		if _wickrListSecurityGroupUsers {
			wickr_ListSecurityGroupUsers(cfg, client)
			return
		}
		if _wickrListSecurityGroups {
			wickr_ListSecurityGroups(cfg, client)
			return
		}
		if _wickrListUsers {
			wickr_ListUsers(cfg, client)
			return
		}
		if _wickrRegisterOidcConfig {
			wickr_RegisterOidcConfig(cfg, client)
			return
		}
		if _wickrRegisterOidcConfigTest {
			wickr_RegisterOidcConfigTest(cfg, client)
			return
		}
		if _wickrRegisterOpentdfConfig {
			wickr_RegisterOpentdfConfig(cfg, client)
			return
		}
		if _wickrUpdateBot {
			wickr_UpdateBot(cfg, client)
			return
		}
		if _wickrUpdateDataRetention {
			wickr_UpdateDataRetention(cfg, client)
			return
		}
		if _wickrUpdateGuestUser {
			wickr_UpdateGuestUser(cfg, client)
			return
		}
		if _wickrUpdateNetwork {
			wickr_UpdateNetwork(cfg, client)
			return
		}
		if _wickrUpdateNetworkSettings {
			wickr_UpdateNetworkSettings(cfg, client)
			return
		}
		if _wickrUpdateSecurityGroup {
			wickr_UpdateSecurityGroup(cfg, client)
			return
		}
		if _wickrUpdateUser {
			wickr_UpdateUser(cfg, client)
			return
		}

	},
}

var (
	_wickrBatchCreateUser                 bool
	_wickrBatchDeleteUser                 bool
	_wickrBatchLookupUserUname            bool
	_wickrBatchReinviteUser               bool
	_wickrBatchResetDevicesForUser        bool
	_wickrBatchToggleUserSuspendStatus    bool
	_wickrCreateBot                       bool
	_wickrCreateDataRetentionBot          bool
	_wickrCreateDataRetentionBotChallenge bool
	_wickrCreateNetwork                   bool
	_wickrCreateSecurityGroup             bool
	_wickrDeleteBot                       bool
	_wickrDeleteDataRetentionBot          bool
	_wickrDeleteNetwork                   bool
	_wickrDeleteSecurityGroup             bool
	_wickrGetBot                          bool
	_wickrGetBotsCount                    bool
	_wickrGetDataRetentionBot             bool
	_wickrGetGuestUserHistoryCount        bool
	_wickrGetNetwork                      bool
	_wickrGetNetworkSettings              bool
	_wickrGetOidcInfo                     bool
	_wickrGetOpentdfConfig                bool
	_wickrGetSecurityGroup                bool
	_wickrGetUser                         bool
	_wickrGetUsersCount                   bool
	_wickrListBlockedGuestUsers           bool
	_wickrListBots                        bool
	_wickrListDevicesForUser              bool
	_wickrListGuestUsers                  bool
	_wickrListNetworks                    bool
	_wickrListSecurityGroupUsers          bool
	_wickrListSecurityGroups              bool
	_wickrListUsers                       bool
	_wickrRegisterOidcConfig              bool
	_wickrRegisterOidcConfigTest          bool
	_wickrRegisterOpentdfConfig           bool
	_wickrUpdateBot                       bool
	_wickrUpdateDataRetention             bool
	_wickrUpdateGuestUser                 bool
	_wickrUpdateNetwork                   bool
	_wickrUpdateNetworkSettings           bool
	_wickrUpdateSecurityGroup             bool
	_wickrUpdateUser                      bool

	_wickrAccessLevel            string
	_wickrActionType             string
	_wickrAdmin                  string
	_wickrAppIds                 []string
	_wickrBillingPeriod          string
	_wickrBlock                  string
	_wickrBotId                  string
	_wickrCertificate            string
	_wickrChallenge              string
	_wickrClientId               string
	_wickrClientSecret           string
	_wickrClientToken            string
	_wickrCode                   string
	_wickrCodeVerifier           string
	_wickrCompanyId              string
	_wickrCustomUsername         string
	_wickrDisplayName            string
	_wickrDomain                 string
	_wickrDryRun                 string
	_wickrEnablePremiumFreeTrial string
	_wickrEncryptionKeyArn       string
	_wickrEndTime                string
	_wickrExtraAuthParams        string
	_wickrFirstName              string
	_wickrGrantType              string
	_wickrGroupId                string
	_wickrIssuer                 string
	_wickrLastName               string
	_wickrMaxResults             string
	_wickrName                   string
	_wickrNetworkId              string
	_wickrNetworkName            string
	_wickrNextToken              string
	_wickrProvider               string
	_wickrRedirectUri            string
	_wickrScopes                 string
	_wickrSecret                 string
	_wickrSecurityGroupSettings  string
	_wickrSettings               string
	_wickrSortDirection          string
	_wickrSortFields             string
	_wickrSsoTokenBufferMinutes  string
	_wickrStartTime              string
	_wickrStatus                 string
	_wickrSuspend                string
	_wickrUnames                 []string
	_wickrUrl                    string
	_wickrUserDetails            string
	_wickrUserId                 string
	_wickrUserIds                []string
	_wickrUsername               string
	_wickrUsernameHash           string
	_wickrUsers                  string
)

// Creates multiple users in a specified Wickr network. This operation allows you
// to provision multiple user accounts simultaneously, optionally specifying
// security groups, and validation requirements for each user.
//
// codeValidation , inviteCode , and inviteCodeTtl are restricted to networks
// under preview only.
func wickr_BatchCreateUser(cfg aws.Config, client *wickr.Client) {
	input := &wickr.BatchCreateUserInput{
		// NetworkId: *string, // Required
		// Users: []types.BatchCreateUserRequestItem, // Required
	}

	if len(_wickrNetworkId) > 0 {
		input.NetworkId = aws.String(_wickrNetworkId)
	}
	if len(_wickrUsers) > 0 {
		if err := assignInputField(input, "Users", _wickrUsers); err != nil {
			log.Errorf("invalid --users: %s", err.Error())
			return
		}
	}
	if len(_wickrClientToken) > 0 {
		input.ClientToken = aws.String(_wickrClientToken)
	}

	if resp, err := client.BatchCreateUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes multiple users from a specified Wickr network. This operation
// permanently removes user accounts and their associated data from the network.
func wickr_BatchDeleteUser(cfg aws.Config, client *wickr.Client) {
	input := &wickr.BatchDeleteUserInput{
		// NetworkId: *string, // Required
		// UserIds: []string, // Required
	}

	if len(_wickrNetworkId) > 0 {
		input.NetworkId = aws.String(_wickrNetworkId)
	}
	if len(_wickrUserIds) > 0 {
		input.UserIds = append([]string(nil), _wickrUserIds...)
	}
	if len(_wickrClientToken) > 0 {
		input.ClientToken = aws.String(_wickrClientToken)
	}

	if resp, err := client.BatchDeleteUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Looks up multiple user usernames from their unique username hashes (unames).
// This operation allows you to retrieve the email addresses associated with a list
// of username hashes.
func wickr_BatchLookupUserUname(cfg aws.Config, client *wickr.Client) {
	input := &wickr.BatchLookupUserUnameInput{
		// NetworkId: *string, // Required
		// Unames: []string, // Required
	}

	if len(_wickrNetworkId) > 0 {
		input.NetworkId = aws.String(_wickrNetworkId)
	}
	if len(_wickrUnames) > 0 {
		input.Unames = append([]string(nil), _wickrUnames...)
	}
	if len(_wickrClientToken) > 0 {
		input.ClientToken = aws.String(_wickrClientToken)
	}

	if resp, err := client.BatchLookupUserUname(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Resends invitation codes to multiple users who have pending invitations in a
// Wickr network. This operation is useful when users haven't accepted their
// initial invitations or when invitations have expired.
func wickr_BatchReinviteUser(cfg aws.Config, client *wickr.Client) {
	input := &wickr.BatchReinviteUserInput{
		// NetworkId: *string, // Required
		// UserIds: []string, // Required
	}

	if len(_wickrNetworkId) > 0 {
		input.NetworkId = aws.String(_wickrNetworkId)
	}
	if len(_wickrUserIds) > 0 {
		input.UserIds = append([]string(nil), _wickrUserIds...)
	}
	if len(_wickrClientToken) > 0 {
		input.ClientToken = aws.String(_wickrClientToken)
	}

	if resp, err := client.BatchReinviteUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Resets multiple devices for a specific user in a Wickr network. This operation
// forces the selected devices to log out and requires users to re-authenticate,
// which is useful for security purposes or when devices need to be revoked.
func wickr_BatchResetDevicesForUser(cfg aws.Config, client *wickr.Client) {
	input := &wickr.BatchResetDevicesForUserInput{
		// AppIds: []string, // Required
		// NetworkId: *string, // Required
		// UserId: *string, // Required
	}

	if len(_wickrAppIds) > 0 {
		input.AppIds = append([]string(nil), _wickrAppIds...)
	}
	if len(_wickrNetworkId) > 0 {
		input.NetworkId = aws.String(_wickrNetworkId)
	}
	if len(_wickrUserId) > 0 {
		input.UserId = aws.String(_wickrUserId)
	}
	if len(_wickrClientToken) > 0 {
		input.ClientToken = aws.String(_wickrClientToken)
	}

	if resp, err := client.BatchResetDevicesForUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Suspends or unsuspends multiple users in a Wickr network. Suspended users
// cannot access the network until they are unsuspended. This operation is useful
// for temporarily restricting access without deleting user accounts.
func wickr_BatchToggleUserSuspendStatus(cfg aws.Config, client *wickr.Client) {
	input := &wickr.BatchToggleUserSuspendStatusInput{
		// NetworkId: *string, // Required
		// Suspend: *bool, // Required
		// UserIds: []string, // Required
	}

	if len(_wickrNetworkId) > 0 {
		input.NetworkId = aws.String(_wickrNetworkId)
	}
	if len(_wickrSuspend) > 0 {
		if err := assignInputField(input, "Suspend", _wickrSuspend); err != nil {
			log.Errorf("invalid --suspend: %s", err.Error())
			return
		}
	}
	if len(_wickrUserIds) > 0 {
		input.UserIds = append([]string(nil), _wickrUserIds...)
	}
	if len(_wickrClientToken) > 0 {
		input.ClientToken = aws.String(_wickrClientToken)
	}

	if resp, err := client.BatchToggleUserSuspendStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new bot in a specified Wickr network. Bots are automated accounts
// that can send and receive messages, enabling integration with external systems
// and automation of tasks.
func wickr_CreateBot(cfg aws.Config, client *wickr.Client) {
	input := &wickr.CreateBotInput{
		// Challenge: *string, // Required
		// GroupId: *string, // Required
		// NetworkId: *string, // Required
		// Username: *string, // Required
	}

	if len(_wickrChallenge) > 0 {
		input.Challenge = aws.String(_wickrChallenge)
	}
	if len(_wickrGroupId) > 0 {
		input.GroupId = aws.String(_wickrGroupId)
	}
	if len(_wickrNetworkId) > 0 {
		input.NetworkId = aws.String(_wickrNetworkId)
	}
	if len(_wickrUsername) > 0 {
		input.Username = aws.String(_wickrUsername)
	}
	if len(_wickrDisplayName) > 0 {
		input.DisplayName = aws.String(_wickrDisplayName)
	}

	if resp, err := client.CreateBot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a data retention bot in a Wickr network. Data retention bots are
// specialized bots that handle message archiving and compliance by capturing and
// storing messages for regulatory or organizational requirements.
func wickr_CreateDataRetentionBot(cfg aws.Config, client *wickr.Client) {
	input := &wickr.CreateDataRetentionBotInput{
		// NetworkId: *string, // Required
	}

	if len(_wickrNetworkId) > 0 {
		input.NetworkId = aws.String(_wickrNetworkId)
	}

	if resp, err := client.CreateDataRetentionBot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new challenge password for the data retention bot. This password is
// used for authentication when the bot connects to the network.
func wickr_CreateDataRetentionBotChallenge(cfg aws.Config, client *wickr.Client) {
	input := &wickr.CreateDataRetentionBotChallengeInput{
		// NetworkId: *string, // Required
	}

	if len(_wickrNetworkId) > 0 {
		input.NetworkId = aws.String(_wickrNetworkId)
	}

	if resp, err := client.CreateDataRetentionBotChallenge(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new Wickr network with specified access level and configuration. This
// operation provisions a new communication network for your organization.
func wickr_CreateNetwork(cfg aws.Config, client *wickr.Client) {
	input := &wickr.CreateNetworkInput{
		// AccessLevel: types.AccessLevel, // Required
		// NetworkName: *string, // Required
	}

	if len(_wickrAccessLevel) > 0 {
		if err := assignInputField(input, "AccessLevel", _wickrAccessLevel); err != nil {
			log.Errorf("invalid --access-level: %s", err.Error())
			return
		}
	}
	if len(_wickrNetworkName) > 0 {
		input.NetworkName = aws.String(_wickrNetworkName)
	}
	if len(_wickrEnablePremiumFreeTrial) > 0 {
		if err := assignInputField(input, "EnablePremiumFreeTrial", _wickrEnablePremiumFreeTrial); err != nil {
			log.Errorf("invalid --enable-premium-free-trial: %s", err.Error())
			return
		}
	}
	if len(_wickrEncryptionKeyArn) > 0 {
		input.EncryptionKeyArn = aws.String(_wickrEncryptionKeyArn)
	}

	if resp, err := client.CreateNetwork(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new security group in a Wickr network. Security groups allow you to
// organize users and control their permissions, features, and security settings.
func wickr_CreateSecurityGroup(cfg aws.Config, client *wickr.Client) {
	input := &wickr.CreateSecurityGroupInput{
		// Name: *string, // Required
		// NetworkId: *string, // Required
		// SecurityGroupSettings: *types.SecurityGroupSettingsRequest, // Required
	}

	if len(_wickrName) > 0 {
		input.Name = aws.String(_wickrName)
	}
	if len(_wickrNetworkId) > 0 {
		input.NetworkId = aws.String(_wickrNetworkId)
	}
	if len(_wickrSecurityGroupSettings) > 0 {
		if err := assignInputField(input, "SecurityGroupSettings", _wickrSecurityGroupSettings); err != nil {
			log.Errorf("invalid --security-group-settings: %s", err.Error())
			return
		}
	}
	if len(_wickrClientToken) > 0 {
		input.ClientToken = aws.String(_wickrClientToken)
	}

	if resp, err := client.CreateSecurityGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a bot from a specified Wickr network. This operation permanently
// removes the bot account and its associated data from the network.
func wickr_DeleteBot(cfg aws.Config, client *wickr.Client) {
	input := &wickr.DeleteBotInput{
		// BotId: *string, // Required
		// NetworkId: *string, // Required
	}

	if len(_wickrBotId) > 0 {
		input.BotId = aws.String(_wickrBotId)
	}
	if len(_wickrNetworkId) > 0 {
		input.NetworkId = aws.String(_wickrNetworkId)
	}

	if resp, err := client.DeleteBot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the data retention bot from a Wickr network. This operation permanently
// removes the bot and all its associated data from the database.
func wickr_DeleteDataRetentionBot(cfg aws.Config, client *wickr.Client) {
	input := &wickr.DeleteDataRetentionBotInput{
		// NetworkId: *string, // Required
	}

	if len(_wickrNetworkId) > 0 {
		input.NetworkId = aws.String(_wickrNetworkId)
	}

	if resp, err := client.DeleteDataRetentionBot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a Wickr network and all its associated resources, including users,
// bots, security groups, and settings. This operation is permanent and cannot be
// undone.
func wickr_DeleteNetwork(cfg aws.Config, client *wickr.Client) {
	input := &wickr.DeleteNetworkInput{
		// NetworkId: *string, // Required
	}

	if len(_wickrNetworkId) > 0 {
		input.NetworkId = aws.String(_wickrNetworkId)
	}
	if len(_wickrClientToken) > 0 {
		input.ClientToken = aws.String(_wickrClientToken)
	}

	if resp, err := client.DeleteNetwork(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a security group from a Wickr network. This operation cannot be
// performed on the default security group.
func wickr_DeleteSecurityGroup(cfg aws.Config, client *wickr.Client) {
	input := &wickr.DeleteSecurityGroupInput{
		// GroupId: *string, // Required
		// NetworkId: *string, // Required
	}

	if len(_wickrGroupId) > 0 {
		input.GroupId = aws.String(_wickrGroupId)
	}
	if len(_wickrNetworkId) > 0 {
		input.NetworkId = aws.String(_wickrNetworkId)
	}

	if resp, err := client.DeleteSecurityGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves detailed information about a specific bot in a Wickr network,
// including its status, group membership, and authentication details.
func wickr_GetBot(cfg aws.Config, client *wickr.Client) {
	input := &wickr.GetBotInput{
		// BotId: *string, // Required
		// NetworkId: *string, // Required
	}

	if len(_wickrBotId) > 0 {
		input.BotId = aws.String(_wickrBotId)
	}
	if len(_wickrNetworkId) > 0 {
		input.NetworkId = aws.String(_wickrNetworkId)
	}

	if resp, err := client.GetBot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the count of bots in a Wickr network, categorized by their status
// (pending, active, and total).
func wickr_GetBotsCount(cfg aws.Config, client *wickr.Client) {
	input := &wickr.GetBotsCountInput{
		// NetworkId: *string, // Required
	}

	if len(_wickrNetworkId) > 0 {
		input.NetworkId = aws.String(_wickrNetworkId)
	}

	if resp, err := client.GetBotsCount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the data retention bot in a Wickr network,
// including its status and whether the data retention service is enabled.
func wickr_GetDataRetentionBot(cfg aws.Config, client *wickr.Client) {
	input := &wickr.GetDataRetentionBotInput{
		// NetworkId: *string, // Required
	}

	if len(_wickrNetworkId) > 0 {
		input.NetworkId = aws.String(_wickrNetworkId)
	}

	if resp, err := client.GetDataRetentionBot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves historical guest user count data for a Wickr network, showing the
// number of guest users per billing period over the past 90 days.
func wickr_GetGuestUserHistoryCount(cfg aws.Config, client *wickr.Client) {
	input := &wickr.GetGuestUserHistoryCountInput{
		// NetworkId: *string, // Required
	}

	if len(_wickrNetworkId) > 0 {
		input.NetworkId = aws.String(_wickrNetworkId)
	}

	if resp, err := client.GetGuestUserHistoryCount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves detailed information about a specific Wickr network, including its
// configuration, access level, and status.
func wickr_GetNetwork(cfg aws.Config, client *wickr.Client) {
	input := &wickr.GetNetworkInput{
		// NetworkId: *string, // Required
	}

	if len(_wickrNetworkId) > 0 {
		input.NetworkId = aws.String(_wickrNetworkId)
	}

	if resp, err := client.GetNetwork(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves all network-level settings for a Wickr network, including client
// metrics, data retention, and other configuration options.
func wickr_GetNetworkSettings(cfg aws.Config, client *wickr.Client) {
	input := &wickr.GetNetworkSettingsInput{
		// NetworkId: *string, // Required
	}

	if len(_wickrNetworkId) > 0 {
		input.NetworkId = aws.String(_wickrNetworkId)
	}

	if resp, err := client.GetNetworkSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the OpenID Connect (OIDC) configuration for a Wickr network,
// including SSO settings and optional token information if access token parameters
// are provided.
func wickr_GetOidcInfo(cfg aws.Config, client *wickr.Client) {
	input := &wickr.GetOidcInfoInput{
		// NetworkId: *string, // Required
	}

	if len(_wickrNetworkId) > 0 {
		input.NetworkId = aws.String(_wickrNetworkId)
	}
	if len(_wickrCertificate) > 0 {
		input.Certificate = aws.String(_wickrCertificate)
	}
	if len(_wickrClientId) > 0 {
		input.ClientId = aws.String(_wickrClientId)
	}
	if len(_wickrClientSecret) > 0 {
		input.ClientSecret = aws.String(_wickrClientSecret)
	}
	if len(_wickrCode) > 0 {
		input.Code = aws.String(_wickrCode)
	}
	if len(_wickrCodeVerifier) > 0 {
		input.CodeVerifier = aws.String(_wickrCodeVerifier)
	}
	if len(_wickrGrantType) > 0 {
		input.GrantType = aws.String(_wickrGrantType)
	}
	if len(_wickrRedirectUri) > 0 {
		input.RedirectUri = aws.String(_wickrRedirectUri)
	}
	if len(_wickrUrl) > 0 {
		input.Url = aws.String(_wickrUrl)
	}

	if resp, err := client.GetOidcInfo(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the OpenTDF integration configuration for a Wickr network.
func wickr_GetOpentdfConfig(cfg aws.Config, client *wickr.Client) {
	input := &wickr.GetOpentdfConfigInput{
		// NetworkId: *string, // Required
	}

	if len(_wickrNetworkId) > 0 {
		input.NetworkId = aws.String(_wickrNetworkId)
	}

	if resp, err := client.GetOpentdfConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves detailed information about a specific security group in a Wickr
// network, including its settings, member counts, and configuration.
func wickr_GetSecurityGroup(cfg aws.Config, client *wickr.Client) {
	input := &wickr.GetSecurityGroupInput{
		// GroupId: *string, // Required
		// NetworkId: *string, // Required
	}

	if len(_wickrGroupId) > 0 {
		input.GroupId = aws.String(_wickrGroupId)
	}
	if len(_wickrNetworkId) > 0 {
		input.NetworkId = aws.String(_wickrNetworkId)
	}

	if resp, err := client.GetSecurityGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves detailed information about a specific user in a Wickr network,
// including their profile, status, and activity history.
func wickr_GetUser(cfg aws.Config, client *wickr.Client) {
	input := &wickr.GetUserInput{
		// NetworkId: *string, // Required
		// UserId: *string, // Required
	}

	if len(_wickrNetworkId) > 0 {
		input.NetworkId = aws.String(_wickrNetworkId)
	}
	if len(_wickrUserId) > 0 {
		input.UserId = aws.String(_wickrUserId)
	}
	if len(_wickrEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _wickrEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_wickrStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _wickrStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the count of users in a Wickr network, categorized by their status
// (pending, active, rejected) and showing how many users can still be added.
func wickr_GetUsersCount(cfg aws.Config, client *wickr.Client) {
	input := &wickr.GetUsersCountInput{
		// NetworkId: *string, // Required
	}

	if len(_wickrNetworkId) > 0 {
		input.NetworkId = aws.String(_wickrNetworkId)
	}

	if resp, err := client.GetUsersCount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a paginated list of guest users who have been blocked from a Wickr
// network. You can filter and sort the results.
func wickr_ListBlockedGuestUsers(cfg aws.Config, client *wickr.Client) {
	input := &wickr.ListBlockedGuestUsersInput{
		// NetworkId: *string, // Required
	}

	if len(_wickrNetworkId) > 0 {
		input.NetworkId = aws.String(_wickrNetworkId)
	}
	if len(_wickrAdmin) > 0 {
		input.Admin = aws.String(_wickrAdmin)
	}
	if len(_wickrMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _wickrMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_wickrNextToken) > 0 {
		input.NextToken = aws.String(_wickrNextToken)
	}
	if len(_wickrSortDirection) > 0 {
		if err := assignInputField(input, "SortDirection", _wickrSortDirection); err != nil {
			log.Errorf("invalid --sort-direction: %s", err.Error())
			return
		}
	}
	if len(_wickrSortFields) > 0 {
		input.SortFields = aws.String(_wickrSortFields)
	}
	if len(_wickrUsername) > 0 {
		input.Username = aws.String(_wickrUsername)
	}

	if disablePaginator() {
		if resp, err := client.ListBlockedGuestUsers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*wickr.ListBlockedGuestUsersOutput
	p := wickr.NewListBlockedGuestUsersPaginator(client, input)
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

// Retrieves a paginated list of bots in a specified Wickr network. You can filter
// and sort the results based on various criteria.
func wickr_ListBots(cfg aws.Config, client *wickr.Client) {
	input := &wickr.ListBotsInput{
		// NetworkId: *string, // Required
	}

	if len(_wickrNetworkId) > 0 {
		input.NetworkId = aws.String(_wickrNetworkId)
	}
	if len(_wickrDisplayName) > 0 {
		input.DisplayName = aws.String(_wickrDisplayName)
	}
	if len(_wickrGroupId) > 0 {
		input.GroupId = aws.String(_wickrGroupId)
	}
	if len(_wickrMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _wickrMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_wickrNextToken) > 0 {
		input.NextToken = aws.String(_wickrNextToken)
	}
	if len(_wickrSortDirection) > 0 {
		if err := assignInputField(input, "SortDirection", _wickrSortDirection); err != nil {
			log.Errorf("invalid --sort-direction: %s", err.Error())
			return
		}
	}
	if len(_wickrSortFields) > 0 {
		input.SortFields = aws.String(_wickrSortFields)
	}
	if len(_wickrStatus) > 0 {
		if err := assignInputField(input, "Status", _wickrStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_wickrUsername) > 0 {
		input.Username = aws.String(_wickrUsername)
	}

	if disablePaginator() {
		if resp, err := client.ListBots(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*wickr.ListBotsOutput
	p := wickr.NewListBotsPaginator(client, input)
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

// Retrieves a paginated list of devices associated with a specific user in a
// Wickr network. This operation returns information about all devices where the
// user has logged into Wickr.
func wickr_ListDevicesForUser(cfg aws.Config, client *wickr.Client) {
	input := &wickr.ListDevicesForUserInput{
		// NetworkId: *string, // Required
		// UserId: *string, // Required
	}

	if len(_wickrNetworkId) > 0 {
		input.NetworkId = aws.String(_wickrNetworkId)
	}
	if len(_wickrUserId) > 0 {
		input.UserId = aws.String(_wickrUserId)
	}
	if len(_wickrMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _wickrMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_wickrNextToken) > 0 {
		input.NextToken = aws.String(_wickrNextToken)
	}
	if len(_wickrSortDirection) > 0 {
		if err := assignInputField(input, "SortDirection", _wickrSortDirection); err != nil {
			log.Errorf("invalid --sort-direction: %s", err.Error())
			return
		}
	}
	if len(_wickrSortFields) > 0 {
		input.SortFields = aws.String(_wickrSortFields)
	}

	if disablePaginator() {
		if resp, err := client.ListDevicesForUser(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*wickr.ListDevicesForUserOutput
	p := wickr.NewListDevicesForUserPaginator(client, input)
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

// Retrieves a paginated list of guest users who have communicated with your Wickr
// network. Guest users are external users from federated networks who can
// communicate with network members.
func wickr_ListGuestUsers(cfg aws.Config, client *wickr.Client) {
	input := &wickr.ListGuestUsersInput{
		// NetworkId: *string, // Required
	}

	if len(_wickrNetworkId) > 0 {
		input.NetworkId = aws.String(_wickrNetworkId)
	}
	if len(_wickrBillingPeriod) > 0 {
		input.BillingPeriod = aws.String(_wickrBillingPeriod)
	}
	if len(_wickrMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _wickrMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_wickrNextToken) > 0 {
		input.NextToken = aws.String(_wickrNextToken)
	}
	if len(_wickrSortDirection) > 0 {
		if err := assignInputField(input, "SortDirection", _wickrSortDirection); err != nil {
			log.Errorf("invalid --sort-direction: %s", err.Error())
			return
		}
	}
	if len(_wickrSortFields) > 0 {
		input.SortFields = aws.String(_wickrSortFields)
	}
	if len(_wickrUsername) > 0 {
		input.Username = aws.String(_wickrUsername)
	}

	if disablePaginator() {
		if resp, err := client.ListGuestUsers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*wickr.ListGuestUsersOutput
	p := wickr.NewListGuestUsersPaginator(client, input)
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

// Retrieves a paginated list of all Wickr networks associated with your Amazon
// Web Services account. You can sort the results by network ID or name.
func wickr_ListNetworks(cfg aws.Config, client *wickr.Client) {
	input := &wickr.ListNetworksInput{}

	if len(_wickrMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _wickrMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_wickrNextToken) > 0 {
		input.NextToken = aws.String(_wickrNextToken)
	}
	if len(_wickrSortDirection) > 0 {
		if err := assignInputField(input, "SortDirection", _wickrSortDirection); err != nil {
			log.Errorf("invalid --sort-direction: %s", err.Error())
			return
		}
	}
	if len(_wickrSortFields) > 0 {
		input.SortFields = aws.String(_wickrSortFields)
	}

	if disablePaginator() {
		if resp, err := client.ListNetworks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*wickr.ListNetworksOutput
	p := wickr.NewListNetworksPaginator(client, input)
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

// Retrieves a paginated list of users who belong to a specific security group in
// a Wickr network.
func wickr_ListSecurityGroupUsers(cfg aws.Config, client *wickr.Client) {
	input := &wickr.ListSecurityGroupUsersInput{
		// GroupId: *string, // Required
		// NetworkId: *string, // Required
	}

	if len(_wickrGroupId) > 0 {
		input.GroupId = aws.String(_wickrGroupId)
	}
	if len(_wickrNetworkId) > 0 {
		input.NetworkId = aws.String(_wickrNetworkId)
	}
	if len(_wickrMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _wickrMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_wickrNextToken) > 0 {
		input.NextToken = aws.String(_wickrNextToken)
	}
	if len(_wickrSortDirection) > 0 {
		if err := assignInputField(input, "SortDirection", _wickrSortDirection); err != nil {
			log.Errorf("invalid --sort-direction: %s", err.Error())
			return
		}
	}
	if len(_wickrSortFields) > 0 {
		input.SortFields = aws.String(_wickrSortFields)
	}

	if disablePaginator() {
		if resp, err := client.ListSecurityGroupUsers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*wickr.ListSecurityGroupUsersOutput
	p := wickr.NewListSecurityGroupUsersPaginator(client, input)
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

// Retrieves a paginated list of security groups in a specified Wickr network. You
// can sort the results by various criteria.
func wickr_ListSecurityGroups(cfg aws.Config, client *wickr.Client) {
	input := &wickr.ListSecurityGroupsInput{
		// NetworkId: *string, // Required
	}

	if len(_wickrNetworkId) > 0 {
		input.NetworkId = aws.String(_wickrNetworkId)
	}
	if len(_wickrMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _wickrMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_wickrNextToken) > 0 {
		input.NextToken = aws.String(_wickrNextToken)
	}
	if len(_wickrSortDirection) > 0 {
		if err := assignInputField(input, "SortDirection", _wickrSortDirection); err != nil {
			log.Errorf("invalid --sort-direction: %s", err.Error())
			return
		}
	}
	if len(_wickrSortFields) > 0 {
		input.SortFields = aws.String(_wickrSortFields)
	}

	if disablePaginator() {
		if resp, err := client.ListSecurityGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*wickr.ListSecurityGroupsOutput
	p := wickr.NewListSecurityGroupsPaginator(client, input)
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

// Retrieves a paginated list of users in a specified Wickr network. You can
// filter and sort the results based on various criteria such as name, status, or
// security group membership.
func wickr_ListUsers(cfg aws.Config, client *wickr.Client) {
	input := &wickr.ListUsersInput{
		// NetworkId: *string, // Required
	}

	if len(_wickrNetworkId) > 0 {
		input.NetworkId = aws.String(_wickrNetworkId)
	}
	if len(_wickrFirstName) > 0 {
		input.FirstName = aws.String(_wickrFirstName)
	}
	if len(_wickrGroupId) > 0 {
		input.GroupId = aws.String(_wickrGroupId)
	}
	if len(_wickrLastName) > 0 {
		input.LastName = aws.String(_wickrLastName)
	}
	if len(_wickrMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _wickrMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_wickrNextToken) > 0 {
		input.NextToken = aws.String(_wickrNextToken)
	}
	if len(_wickrSortDirection) > 0 {
		if err := assignInputField(input, "SortDirection", _wickrSortDirection); err != nil {
			log.Errorf("invalid --sort-direction: %s", err.Error())
			return
		}
	}
	if len(_wickrSortFields) > 0 {
		input.SortFields = aws.String(_wickrSortFields)
	}
	if len(_wickrStatus) > 0 {
		if err := assignInputField(input, "Status", _wickrStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_wickrUsername) > 0 {
		input.Username = aws.String(_wickrUsername)
	}

	if disablePaginator() {
		if resp, err := client.ListUsers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*wickr.ListUsersOutput
	p := wickr.NewListUsersPaginator(client, input)
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

// Registers and saves an OpenID Connect (OIDC) configuration for a Wickr network,
// enabling Single Sign-On (SSO) authentication through an identity provider.
func wickr_RegisterOidcConfig(cfg aws.Config, client *wickr.Client) {
	input := &wickr.RegisterOidcConfigInput{
		// CompanyId: *string, // Required
		// Issuer: *string, // Required
		// NetworkId: *string, // Required
		// Scopes: *string, // Required
	}

	if len(_wickrCompanyId) > 0 {
		input.CompanyId = aws.String(_wickrCompanyId)
	}
	if len(_wickrIssuer) > 0 {
		input.Issuer = aws.String(_wickrIssuer)
	}
	if len(_wickrNetworkId) > 0 {
		input.NetworkId = aws.String(_wickrNetworkId)
	}
	if len(_wickrScopes) > 0 {
		input.Scopes = aws.String(_wickrScopes)
	}
	if len(_wickrCustomUsername) > 0 {
		input.CustomUsername = aws.String(_wickrCustomUsername)
	}
	if len(_wickrExtraAuthParams) > 0 {
		input.ExtraAuthParams = aws.String(_wickrExtraAuthParams)
	}
	if len(_wickrSecret) > 0 {
		input.Secret = aws.String(_wickrSecret)
	}
	if len(_wickrSsoTokenBufferMinutes) > 0 {
		if err := assignInputField(input, "SsoTokenBufferMinutes", _wickrSsoTokenBufferMinutes); err != nil {
			log.Errorf("invalid --sso-token-buffer-minutes: %s", err.Error())
			return
		}
	}
	if len(_wickrUserId) > 0 {
		input.UserId = aws.String(_wickrUserId)
	}

	if resp, err := client.RegisterOidcConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Tests an OpenID Connect (OIDC) configuration for a Wickr network by validating
// the connection to the identity provider and retrieving its supported
// capabilities.
func wickr_RegisterOidcConfigTest(cfg aws.Config, client *wickr.Client) {
	input := &wickr.RegisterOidcConfigTestInput{
		// Issuer: *string, // Required
		// NetworkId: *string, // Required
		// Scopes: *string, // Required
	}

	if len(_wickrIssuer) > 0 {
		input.Issuer = aws.String(_wickrIssuer)
	}
	if len(_wickrNetworkId) > 0 {
		input.NetworkId = aws.String(_wickrNetworkId)
	}
	if len(_wickrScopes) > 0 {
		input.Scopes = aws.String(_wickrScopes)
	}
	if len(_wickrCertificate) > 0 {
		input.Certificate = aws.String(_wickrCertificate)
	}
	if len(_wickrExtraAuthParams) > 0 {
		input.ExtraAuthParams = aws.String(_wickrExtraAuthParams)
	}

	if resp, err := client.RegisterOidcConfigTest(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Registers and saves OpenTDF configuration for a Wickr network, enabling
// attribute-based access control for Wickr through an OpenTDF provider.
func wickr_RegisterOpentdfConfig(cfg aws.Config, client *wickr.Client) {
	input := &wickr.RegisterOpentdfConfigInput{
		// ClientId: *string, // Required
		// ClientSecret: *string, // Required
		// Domain: *string, // Required
		// NetworkId: *string, // Required
		// Provider: *string, // Required
	}

	if len(_wickrClientId) > 0 {
		input.ClientId = aws.String(_wickrClientId)
	}
	if len(_wickrClientSecret) > 0 {
		input.ClientSecret = aws.String(_wickrClientSecret)
	}
	if len(_wickrDomain) > 0 {
		input.Domain = aws.String(_wickrDomain)
	}
	if len(_wickrNetworkId) > 0 {
		input.NetworkId = aws.String(_wickrNetworkId)
	}
	if len(_wickrProvider) > 0 {
		input.Provider = aws.String(_wickrProvider)
	}
	if len(_wickrDryRun) > 0 {
		if err := assignInputField(input, "DryRun", _wickrDryRun); err != nil {
			log.Errorf("invalid --dry-run: %s", err.Error())
			return
		}
	}

	if resp, err := client.RegisterOpentdfConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the properties of an existing bot in a Wickr network. This operation
// allows you to modify the bot's display name, security group, password, or
// suspension status.
func wickr_UpdateBot(cfg aws.Config, client *wickr.Client) {
	input := &wickr.UpdateBotInput{
		// BotId: *string, // Required
		// NetworkId: *string, // Required
	}

	if len(_wickrBotId) > 0 {
		input.BotId = aws.String(_wickrBotId)
	}
	if len(_wickrNetworkId) > 0 {
		input.NetworkId = aws.String(_wickrNetworkId)
	}
	if len(_wickrChallenge) > 0 {
		input.Challenge = aws.String(_wickrChallenge)
	}
	if len(_wickrDisplayName) > 0 {
		input.DisplayName = aws.String(_wickrDisplayName)
	}
	if len(_wickrGroupId) > 0 {
		input.GroupId = aws.String(_wickrGroupId)
	}
	if len(_wickrSuspend) > 0 {
		if err := assignInputField(input, "Suspend", _wickrSuspend); err != nil {
			log.Errorf("invalid --suspend: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateBot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the data retention bot settings, allowing you to enable or disable the
// data retention service, or acknowledge the public key message.
func wickr_UpdateDataRetention(cfg aws.Config, client *wickr.Client) {
	input := &wickr.UpdateDataRetentionInput{
		// ActionType: types.DataRetentionActionType, // Required
		// NetworkId: *string, // Required
	}

	if len(_wickrActionType) > 0 {
		if err := assignInputField(input, "ActionType", _wickrActionType); err != nil {
			log.Errorf("invalid --action-type: %s", err.Error())
			return
		}
	}
	if len(_wickrNetworkId) > 0 {
		input.NetworkId = aws.String(_wickrNetworkId)
	}

	if resp, err := client.UpdateDataRetention(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the block status of a guest user in a Wickr network. This operation
// allows you to block or unblock a guest user from accessing the network.
func wickr_UpdateGuestUser(cfg aws.Config, client *wickr.Client) {
	input := &wickr.UpdateGuestUserInput{
		// Block: *bool, // Required
		// NetworkId: *string, // Required
		// UsernameHash: *string, // Required
	}

	if len(_wickrBlock) > 0 {
		if err := assignInputField(input, "Block", _wickrBlock); err != nil {
			log.Errorf("invalid --block: %s", err.Error())
			return
		}
	}
	if len(_wickrNetworkId) > 0 {
		input.NetworkId = aws.String(_wickrNetworkId)
	}
	if len(_wickrUsernameHash) > 0 {
		input.UsernameHash = aws.String(_wickrUsernameHash)
	}

	if resp, err := client.UpdateGuestUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the properties of an existing Wickr network, such as its name or
// encryption key configuration.
func wickr_UpdateNetwork(cfg aws.Config, client *wickr.Client) {
	input := &wickr.UpdateNetworkInput{
		// NetworkId: *string, // Required
		// NetworkName: *string, // Required
	}

	if len(_wickrNetworkId) > 0 {
		input.NetworkId = aws.String(_wickrNetworkId)
	}
	if len(_wickrNetworkName) > 0 {
		input.NetworkName = aws.String(_wickrNetworkName)
	}
	if len(_wickrClientToken) > 0 {
		input.ClientToken = aws.String(_wickrClientToken)
	}
	if len(_wickrEncryptionKeyArn) > 0 {
		input.EncryptionKeyArn = aws.String(_wickrEncryptionKeyArn)
	}

	if resp, err := client.UpdateNetwork(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates network-level settings for a Wickr network. You can modify settings
// such as client metrics, data retention, and other network-wide options.
func wickr_UpdateNetworkSettings(cfg aws.Config, client *wickr.Client) {
	input := &wickr.UpdateNetworkSettingsInput{
		// NetworkId: *string, // Required
		// Settings: *types.NetworkSettings, // Required
	}

	if len(_wickrNetworkId) > 0 {
		input.NetworkId = aws.String(_wickrNetworkId)
	}
	if len(_wickrSettings) > 0 {
		if err := assignInputField(input, "Settings", _wickrSettings); err != nil {
			log.Errorf("invalid --settings: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateNetworkSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the properties of an existing security group in a Wickr network, such
// as its name or settings.
func wickr_UpdateSecurityGroup(cfg aws.Config, client *wickr.Client) {
	input := &wickr.UpdateSecurityGroupInput{
		// GroupId: *string, // Required
		// NetworkId: *string, // Required
	}

	if len(_wickrGroupId) > 0 {
		input.GroupId = aws.String(_wickrGroupId)
	}
	if len(_wickrNetworkId) > 0 {
		input.NetworkId = aws.String(_wickrNetworkId)
	}
	if len(_wickrName) > 0 {
		input.Name = aws.String(_wickrName)
	}
	if len(_wickrSecurityGroupSettings) > 0 {
		if err := assignInputField(input, "SecurityGroupSettings", _wickrSecurityGroupSettings); err != nil {
			log.Errorf("invalid --security-group-settings: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateSecurityGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the properties of an existing user in a Wickr network. This operation
// allows you to modify the user's name, password, security group membership, and
// invite code settings.
//
// codeValidation , inviteCode , and inviteCodeTtl are restricted to networks
// under preview only.
func wickr_UpdateUser(cfg aws.Config, client *wickr.Client) {
	input := &wickr.UpdateUserInput{
		// NetworkId: *string, // Required
		// UserId: *string, // Required
	}

	if len(_wickrNetworkId) > 0 {
		input.NetworkId = aws.String(_wickrNetworkId)
	}
	if len(_wickrUserId) > 0 {
		input.UserId = aws.String(_wickrUserId)
	}
	if len(_wickrUserDetails) > 0 {
		if err := assignInputField(input, "UserDetails", _wickrUserDetails); err != nil {
			log.Errorf("invalid --user-details: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_wickrCmd)
	_wickrCmd.Flags().SortFlags = false

	_wickrCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_wickrCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_wickrCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_wickrCmd.Flags().StringVarP(&_wickrAccessLevel, "access-level", "", "", "Access Level")
	_wickrCmd.Flags().StringVarP(&_wickrActionType, "action-type", "", "", "Action Type")
	_wickrCmd.Flags().StringVarP(&_wickrAdmin, "admin", "", "", "Admin")
	_wickrCmd.Flags().StringSliceVarP(&_wickrAppIds, "app-ids", "", nil, "App Ids")
	_wickrCmd.Flags().StringVarP(&_wickrBillingPeriod, "billing-period", "", "", "Billing Period")
	_wickrCmd.Flags().StringVarP(&_wickrBlock, "block", "", "", "Block")
	_wickrCmd.Flags().StringVarP(&_wickrBotId, "bot-id", "", "", "Bot ID")
	_wickrCmd.Flags().StringVarP(&_wickrCertificate, "certificate", "", "", "Certificate")
	_wickrCmd.Flags().StringVarP(&_wickrChallenge, "challenge", "", "", "Challenge")
	_wickrCmd.Flags().StringVarP(&_wickrClientId, "client-id", "", "", "Client ID")
	_wickrCmd.Flags().StringVarP(&_wickrClientSecret, "client-secret", "", "", "Client Secret")
	_wickrCmd.Flags().StringVarP(&_wickrClientToken, "client-token", "", "", "Client Token")
	_wickrCmd.Flags().StringVarP(&_wickrCode, "code", "", "", "Code")
	_wickrCmd.Flags().StringVarP(&_wickrCodeVerifier, "code-verifier", "", "", "Code Verifier")
	_wickrCmd.Flags().StringVarP(&_wickrCompanyId, "company-id", "", "", "Company ID")
	_wickrCmd.Flags().StringVarP(&_wickrCustomUsername, "custom-username", "", "", "Custom Username")
	_wickrCmd.Flags().StringVarP(&_wickrDisplayName, "display-name", "", "", "Display Name")
	_wickrCmd.Flags().StringVarP(&_wickrDomain, "domain", "", "", "Domain")
	_wickrCmd.Flags().StringVarP(&_wickrDryRun, "dry-run", "", "", "Dry Run")
	_wickrCmd.Flags().StringVarP(&_wickrEnablePremiumFreeTrial, "enable-premium-free-trial", "", "", "Enable Premium Free Trial")
	_wickrCmd.Flags().StringVarP(&_wickrEncryptionKeyArn, "encryption-key-arn", "", "", "Encryption Key ARN")
	_wickrCmd.Flags().StringVarP(&_wickrEndTime, "end-time", "", "", "End Time")
	_wickrCmd.Flags().StringVarP(&_wickrExtraAuthParams, "extra-auth-params", "", "", "Extra Auth Params")
	_wickrCmd.Flags().StringVarP(&_wickrFirstName, "first-name", "", "", "First Name")
	_wickrCmd.Flags().StringVarP(&_wickrGrantType, "grant-type", "", "", "Grant Type")
	_wickrCmd.Flags().StringVarP(&_wickrGroupId, "group-id", "", "", "Group ID")
	_wickrCmd.Flags().StringVarP(&_wickrIssuer, "issuer", "", "", "Issuer")
	_wickrCmd.Flags().StringVarP(&_wickrLastName, "last-name", "", "", "Last Name")
	_wickrCmd.Flags().StringVarP(&_wickrMaxResults, "max-results", "", "", "Max Results")
	_wickrCmd.Flags().StringVarP(&_wickrName, "name", "", "", "Name")
	_wickrCmd.Flags().StringVarP(&_wickrNetworkId, "network-id", "", "", "Network ID")
	_wickrCmd.Flags().StringVarP(&_wickrNetworkName, "network-name", "", "", "Network Name")
	_wickrCmd.Flags().StringVarP(&_wickrNextToken, "next-token", "", "", "Next Token")
	_wickrCmd.Flags().StringVarP(&_wickrProvider, "provider", "", "", "Provider")
	_wickrCmd.Flags().StringVarP(&_wickrRedirectUri, "redirect-uri", "", "", "Redirect URI")
	_wickrCmd.Flags().StringVarP(&_wickrScopes, "scopes", "", "", "Scopes")
	_wickrCmd.Flags().StringVarP(&_wickrSecret, "secret", "", "", "Secret")
	_wickrCmd.Flags().StringVarP(&_wickrSecurityGroupSettings, "security-group-settings", "", "", "Security Group Settings")
	_wickrCmd.Flags().StringVarP(&_wickrSettings, "settings", "", "", "Settings")
	_wickrCmd.Flags().StringVarP(&_wickrSortDirection, "sort-direction", "", "", "Sort Direction")
	_wickrCmd.Flags().StringVarP(&_wickrSortFields, "sort-fields", "", "", "Sort Fields")
	_wickrCmd.Flags().StringVarP(&_wickrSsoTokenBufferMinutes, "sso-token-buffer-minutes", "", "", "Sso Token Buffer Minutes")
	_wickrCmd.Flags().StringVarP(&_wickrStartTime, "start-time", "", "", "Start Time")
	_wickrCmd.Flags().StringVarP(&_wickrStatus, "status", "", "", "Status")
	_wickrCmd.Flags().StringVarP(&_wickrSuspend, "suspend", "", "", "Suspend")
	_wickrCmd.Flags().StringSliceVarP(&_wickrUnames, "unames", "", nil, "Unames")
	_wickrCmd.Flags().StringVarP(&_wickrUrl, "url", "", "", "URL")
	_wickrCmd.Flags().StringVarP(&_wickrUserDetails, "user-details", "", "", "User Details")
	_wickrCmd.Flags().StringVarP(&_wickrUserId, "user-id", "", "", "User ID")
	_wickrCmd.Flags().StringSliceVarP(&_wickrUserIds, "user-ids", "", nil, "User Ids")
	_wickrCmd.Flags().StringVarP(&_wickrUsername, "username", "", "", "Username")
	_wickrCmd.Flags().StringVarP(&_wickrUsernameHash, "username-hash", "", "", "Username Hash")
	_wickrCmd.Flags().StringVarP(&_wickrUsers, "users", "", "", "Users")

	_wickrCmd.Flags().BoolVarP(&_wickrBatchCreateUser, "batch-create-user", "", false, "Batch Create User")
	_wickrCmd.Flags().BoolVarP(&_wickrBatchDeleteUser, "batch-delete-user", "", false, "Batch Delete User")
	_wickrCmd.Flags().BoolVarP(&_wickrBatchLookupUserUname, "batch-lookup-user-uname", "", false, "Batch Lookup User Uname")
	_wickrCmd.Flags().BoolVarP(&_wickrBatchReinviteUser, "batch-reinvite-user", "", false, "Batch Reinvite User")
	_wickrCmd.Flags().BoolVarP(&_wickrBatchResetDevicesForUser, "batch-reset-devices-for-user", "", false, "Batch Reset Devices For User")
	_wickrCmd.Flags().BoolVarP(&_wickrBatchToggleUserSuspendStatus, "batch-toggle-user-suspend-status", "", false, "Batch Toggle User Suspend Status")
	_wickrCmd.Flags().BoolVarP(&_wickrCreateBot, "create-bot", "", false, "Create Bot")
	_wickrCmd.Flags().BoolVarP(&_wickrCreateDataRetentionBot, "create-data-retention-bot", "", false, "Create Data Retention Bot")
	_wickrCmd.Flags().BoolVarP(&_wickrCreateDataRetentionBotChallenge, "create-data-retention-bot-challenge", "", false, "Create Data Retention Bot Challenge")
	_wickrCmd.Flags().BoolVarP(&_wickrCreateNetwork, "create-network", "", false, "Create Network")
	_wickrCmd.Flags().BoolVarP(&_wickrCreateSecurityGroup, "create-security-group", "", false, "Create Security Group")
	_wickrCmd.Flags().BoolVarP(&_wickrDeleteBot, "delete-bot", "", false, "Delete Bot")
	_wickrCmd.Flags().BoolVarP(&_wickrDeleteDataRetentionBot, "delete-data-retention-bot", "", false, "Delete Data Retention Bot")
	_wickrCmd.Flags().BoolVarP(&_wickrDeleteNetwork, "delete-network", "", false, "Delete Network")
	_wickrCmd.Flags().BoolVarP(&_wickrDeleteSecurityGroup, "delete-security-group", "", false, "Delete Security Group")
	_wickrCmd.Flags().BoolVarP(&_wickrGetBot, "get-bot", "", false, "Get Bot")
	_wickrCmd.Flags().BoolVarP(&_wickrGetBotsCount, "get-bots-count", "", false, "Get Bots Count")
	_wickrCmd.Flags().BoolVarP(&_wickrGetDataRetentionBot, "get-data-retention-bot", "", false, "Get Data Retention Bot")
	_wickrCmd.Flags().BoolVarP(&_wickrGetGuestUserHistoryCount, "get-guest-user-history-count", "", false, "Get Guest User History Count")
	_wickrCmd.Flags().BoolVarP(&_wickrGetNetwork, "get-network", "", false, "Get Network")
	_wickrCmd.Flags().BoolVarP(&_wickrGetNetworkSettings, "get-network-settings", "", false, "Get Network Settings")
	_wickrCmd.Flags().BoolVarP(&_wickrGetOidcInfo, "get-oidc-info", "", false, "Get OIDC Info")
	_wickrCmd.Flags().BoolVarP(&_wickrGetOpentdfConfig, "get-opentdf-config", "", false, "Get Opentdf Config")
	_wickrCmd.Flags().BoolVarP(&_wickrGetSecurityGroup, "get-security-group", "", false, "Get Security Group")
	_wickrCmd.Flags().BoolVarP(&_wickrGetUser, "get-user", "", false, "Get User")
	_wickrCmd.Flags().BoolVarP(&_wickrGetUsersCount, "get-users-count", "", false, "Get Users Count")
	_wickrCmd.Flags().BoolVarP(&_wickrListBlockedGuestUsers, "list-blocked-guest-users", "", false, "List Blocked Guest Users")
	_wickrCmd.Flags().BoolVarP(&_wickrListBots, "list-bots", "", false, "List Bots")
	_wickrCmd.Flags().BoolVarP(&_wickrListDevicesForUser, "list-devices-for-user", "", false, "List Devices For User")
	_wickrCmd.Flags().BoolVarP(&_wickrListGuestUsers, "list-guest-users", "", false, "List Guest Users")
	_wickrCmd.Flags().BoolVarP(&_wickrListNetworks, "list-networks", "", false, "List Networks")
	_wickrCmd.Flags().BoolVarP(&_wickrListSecurityGroupUsers, "list-security-group-users", "", false, "List Security Group Users")
	_wickrCmd.Flags().BoolVarP(&_wickrListSecurityGroups, "list-security-groups", "", false, "List Security Groups")
	_wickrCmd.Flags().BoolVarP(&_wickrListUsers, "list-users", "", false, "List Users")
	_wickrCmd.Flags().BoolVarP(&_wickrRegisterOidcConfig, "register-oidc-config", "", false, "Register OIDC Config")
	_wickrCmd.Flags().BoolVarP(&_wickrRegisterOidcConfigTest, "register-oidc-config-test", "", false, "Register OIDC Config Test")
	_wickrCmd.Flags().BoolVarP(&_wickrRegisterOpentdfConfig, "register-opentdf-config", "", false, "Register Opentdf Config")
	_wickrCmd.Flags().BoolVarP(&_wickrUpdateBot, "update-bot", "", false, "Update Bot")
	_wickrCmd.Flags().BoolVarP(&_wickrUpdateDataRetention, "update-data-retention", "", false, "Update Data Retention")
	_wickrCmd.Flags().BoolVarP(&_wickrUpdateGuestUser, "update-guest-user", "", false, "Update Guest User")
	_wickrCmd.Flags().BoolVarP(&_wickrUpdateNetwork, "update-network", "", false, "Update Network")
	_wickrCmd.Flags().BoolVarP(&_wickrUpdateNetworkSettings, "update-network-settings", "", false, "Update Network Settings")
	_wickrCmd.Flags().BoolVarP(&_wickrUpdateSecurityGroup, "update-security-group", "", false, "Update Security Group")
	_wickrCmd.Flags().BoolVarP(&_wickrUpdateUser, "update-user", "", false, "Update User")

}
