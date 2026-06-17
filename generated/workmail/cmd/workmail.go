package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/workmail"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// workmailCmd represents the workmail command
var _workmailCmd = &cobra.Command{
	Use:   "workmail",
	Short: "AWS workmail CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := workmail.NewFromConfig(cfg)
		if _workmailAssociateDelegateToResource {
			workmail_AssociateDelegateToResource(cfg, client)
			return
		}
		if _workmailAssociateMemberToGroup {
			workmail_AssociateMemberToGroup(cfg, client)
			return
		}
		if _workmailAssumeImpersonationRole {
			workmail_AssumeImpersonationRole(cfg, client)
			return
		}
		if _workmailCancelMailboxExportJob {
			workmail_CancelMailboxExportJob(cfg, client)
			return
		}
		if _workmailCreateAlias {
			workmail_CreateAlias(cfg, client)
			return
		}
		if _workmailCreateAvailabilityConfiguration {
			workmail_CreateAvailabilityConfiguration(cfg, client)
			return
		}
		if _workmailCreateGroup {
			workmail_CreateGroup(cfg, client)
			return
		}
		if _workmailCreateIdentityCenterApplication {
			workmail_CreateIdentityCenterApplication(cfg, client)
			return
		}
		if _workmailCreateImpersonationRole {
			workmail_CreateImpersonationRole(cfg, client)
			return
		}
		if _workmailCreateMobileDeviceAccessRule {
			workmail_CreateMobileDeviceAccessRule(cfg, client)
			return
		}
		if _workmailCreateOrganization {
			workmail_CreateOrganization(cfg, client)
			return
		}
		if _workmailCreateResource {
			workmail_CreateResource(cfg, client)
			return
		}
		if _workmailCreateUser {
			workmail_CreateUser(cfg, client)
			return
		}
		if _workmailDeleteAccessControlRule {
			workmail_DeleteAccessControlRule(cfg, client)
			return
		}
		if _workmailDeleteAlias {
			workmail_DeleteAlias(cfg, client)
			return
		}
		if _workmailDeleteAvailabilityConfiguration {
			workmail_DeleteAvailabilityConfiguration(cfg, client)
			return
		}
		if _workmailDeleteEmailMonitoringConfiguration {
			workmail_DeleteEmailMonitoringConfiguration(cfg, client)
			return
		}
		if _workmailDeleteGroup {
			workmail_DeleteGroup(cfg, client)
			return
		}
		if _workmailDeleteIdentityCenterApplication {
			workmail_DeleteIdentityCenterApplication(cfg, client)
			return
		}
		if _workmailDeleteIdentityProviderConfiguration {
			workmail_DeleteIdentityProviderConfiguration(cfg, client)
			return
		}
		if _workmailDeleteImpersonationRole {
			workmail_DeleteImpersonationRole(cfg, client)
			return
		}
		if _workmailDeleteMailboxPermissions {
			workmail_DeleteMailboxPermissions(cfg, client)
			return
		}
		if _workmailDeleteMobileDeviceAccessOverride {
			workmail_DeleteMobileDeviceAccessOverride(cfg, client)
			return
		}
		if _workmailDeleteMobileDeviceAccessRule {
			workmail_DeleteMobileDeviceAccessRule(cfg, client)
			return
		}
		if _workmailDeleteOrganization {
			workmail_DeleteOrganization(cfg, client)
			return
		}
		if _workmailDeletePersonalAccessToken {
			workmail_DeletePersonalAccessToken(cfg, client)
			return
		}
		if _workmailDeleteResource {
			workmail_DeleteResource(cfg, client)
			return
		}
		if _workmailDeleteRetentionPolicy {
			workmail_DeleteRetentionPolicy(cfg, client)
			return
		}
		if _workmailDeleteUser {
			workmail_DeleteUser(cfg, client)
			return
		}
		if _workmailDeregisterFromWorkMail {
			workmail_DeregisterFromWorkMail(cfg, client)
			return
		}
		if _workmailDeregisterMailDomain {
			workmail_DeregisterMailDomain(cfg, client)
			return
		}
		if _workmailDescribeEmailMonitoringConfiguration {
			workmail_DescribeEmailMonitoringConfiguration(cfg, client)
			return
		}
		if _workmailDescribeEntity {
			workmail_DescribeEntity(cfg, client)
			return
		}
		if _workmailDescribeGroup {
			workmail_DescribeGroup(cfg, client)
			return
		}
		if _workmailDescribeIdentityProviderConfiguration {
			workmail_DescribeIdentityProviderConfiguration(cfg, client)
			return
		}
		if _workmailDescribeInboundDmarcSettings {
			workmail_DescribeInboundDmarcSettings(cfg, client)
			return
		}
		if _workmailDescribeMailboxExportJob {
			workmail_DescribeMailboxExportJob(cfg, client)
			return
		}
		if _workmailDescribeOrganization {
			workmail_DescribeOrganization(cfg, client)
			return
		}
		if _workmailDescribeResource {
			workmail_DescribeResource(cfg, client)
			return
		}
		if _workmailDescribeUser {
			workmail_DescribeUser(cfg, client)
			return
		}
		if _workmailDisassociateDelegateFromResource {
			workmail_DisassociateDelegateFromResource(cfg, client)
			return
		}
		if _workmailDisassociateMemberFromGroup {
			workmail_DisassociateMemberFromGroup(cfg, client)
			return
		}
		if _workmailGetAccessControlEffect {
			workmail_GetAccessControlEffect(cfg, client)
			return
		}
		if _workmailGetDefaultRetentionPolicy {
			workmail_GetDefaultRetentionPolicy(cfg, client)
			return
		}
		if _workmailGetImpersonationRole {
			workmail_GetImpersonationRole(cfg, client)
			return
		}
		if _workmailGetImpersonationRoleEffect {
			workmail_GetImpersonationRoleEffect(cfg, client)
			return
		}
		if _workmailGetMailDomain {
			workmail_GetMailDomain(cfg, client)
			return
		}
		if _workmailGetMailboxDetails {
			workmail_GetMailboxDetails(cfg, client)
			return
		}
		if _workmailGetMobileDeviceAccessEffect {
			workmail_GetMobileDeviceAccessEffect(cfg, client)
			return
		}
		if _workmailGetMobileDeviceAccessOverride {
			workmail_GetMobileDeviceAccessOverride(cfg, client)
			return
		}
		if _workmailGetPersonalAccessTokenMetadata {
			workmail_GetPersonalAccessTokenMetadata(cfg, client)
			return
		}
		if _workmailListAccessControlRules {
			workmail_ListAccessControlRules(cfg, client)
			return
		}
		if _workmailListAliases {
			workmail_ListAliases(cfg, client)
			return
		}
		if _workmailListAvailabilityConfigurations {
			workmail_ListAvailabilityConfigurations(cfg, client)
			return
		}
		if _workmailListGroupMembers {
			workmail_ListGroupMembers(cfg, client)
			return
		}
		if _workmailListGroups {
			workmail_ListGroups(cfg, client)
			return
		}
		if _workmailListGroupsForEntity {
			workmail_ListGroupsForEntity(cfg, client)
			return
		}
		if _workmailListImpersonationRoles {
			workmail_ListImpersonationRoles(cfg, client)
			return
		}
		if _workmailListMailDomains {
			workmail_ListMailDomains(cfg, client)
			return
		}
		if _workmailListMailboxExportJobs {
			workmail_ListMailboxExportJobs(cfg, client)
			return
		}
		if _workmailListMailboxPermissions {
			workmail_ListMailboxPermissions(cfg, client)
			return
		}
		if _workmailListMobileDeviceAccessOverrides {
			workmail_ListMobileDeviceAccessOverrides(cfg, client)
			return
		}
		if _workmailListMobileDeviceAccessRules {
			workmail_ListMobileDeviceAccessRules(cfg, client)
			return
		}
		if _workmailListOrganizations {
			workmail_ListOrganizations(cfg, client)
			return
		}
		if _workmailListPersonalAccessTokens {
			workmail_ListPersonalAccessTokens(cfg, client)
			return
		}
		if _workmailListResourceDelegates {
			workmail_ListResourceDelegates(cfg, client)
			return
		}
		if _workmailListResources {
			workmail_ListResources(cfg, client)
			return
		}
		if _workmailListTagsForResource {
			workmail_ListTagsForResource(cfg, client)
			return
		}
		if _workmailListUsers {
			workmail_ListUsers(cfg, client)
			return
		}
		if _workmailPutAccessControlRule {
			workmail_PutAccessControlRule(cfg, client)
			return
		}
		if _workmailPutEmailMonitoringConfiguration {
			workmail_PutEmailMonitoringConfiguration(cfg, client)
			return
		}
		if _workmailPutIdentityProviderConfiguration {
			workmail_PutIdentityProviderConfiguration(cfg, client)
			return
		}
		if _workmailPutInboundDmarcSettings {
			workmail_PutInboundDmarcSettings(cfg, client)
			return
		}
		if _workmailPutMailboxPermissions {
			workmail_PutMailboxPermissions(cfg, client)
			return
		}
		if _workmailPutMobileDeviceAccessOverride {
			workmail_PutMobileDeviceAccessOverride(cfg, client)
			return
		}
		if _workmailPutRetentionPolicy {
			workmail_PutRetentionPolicy(cfg, client)
			return
		}
		if _workmailRegisterMailDomain {
			workmail_RegisterMailDomain(cfg, client)
			return
		}
		if _workmailRegisterToWorkMail {
			workmail_RegisterToWorkMail(cfg, client)
			return
		}
		if _workmailResetPassword {
			workmail_ResetPassword(cfg, client)
			return
		}
		if _workmailStartMailboxExportJob {
			workmail_StartMailboxExportJob(cfg, client)
			return
		}
		if _workmailTagResource {
			workmail_TagResource(cfg, client)
			return
		}
		if _workmailTestAvailabilityConfiguration {
			workmail_TestAvailabilityConfiguration(cfg, client)
			return
		}
		if _workmailUntagResource {
			workmail_UntagResource(cfg, client)
			return
		}
		if _workmailUpdateAvailabilityConfiguration {
			workmail_UpdateAvailabilityConfiguration(cfg, client)
			return
		}
		if _workmailUpdateDefaultMailDomain {
			workmail_UpdateDefaultMailDomain(cfg, client)
			return
		}
		if _workmailUpdateGroup {
			workmail_UpdateGroup(cfg, client)
			return
		}
		if _workmailUpdateImpersonationRole {
			workmail_UpdateImpersonationRole(cfg, client)
			return
		}
		if _workmailUpdateMailboxQuota {
			workmail_UpdateMailboxQuota(cfg, client)
			return
		}
		if _workmailUpdateMobileDeviceAccessRule {
			workmail_UpdateMobileDeviceAccessRule(cfg, client)
			return
		}
		if _workmailUpdatePrimaryEmailAddress {
			workmail_UpdatePrimaryEmailAddress(cfg, client)
			return
		}
		if _workmailUpdateResource {
			workmail_UpdateResource(cfg, client)
			return
		}
		if _workmailUpdateUser {
			workmail_UpdateUser(cfg, client)
			return
		}

	},
}

var (
	_workmailAssociateDelegateToResource           bool
	_workmailAssociateMemberToGroup                bool
	_workmailAssumeImpersonationRole               bool
	_workmailCancelMailboxExportJob                bool
	_workmailCreateAlias                           bool
	_workmailCreateAvailabilityConfiguration       bool
	_workmailCreateGroup                           bool
	_workmailCreateIdentityCenterApplication       bool
	_workmailCreateImpersonationRole               bool
	_workmailCreateMobileDeviceAccessRule          bool
	_workmailCreateOrganization                    bool
	_workmailCreateResource                        bool
	_workmailCreateUser                            bool
	_workmailDeleteAccessControlRule               bool
	_workmailDeleteAlias                           bool
	_workmailDeleteAvailabilityConfiguration       bool
	_workmailDeleteEmailMonitoringConfiguration    bool
	_workmailDeleteGroup                           bool
	_workmailDeleteIdentityCenterApplication       bool
	_workmailDeleteIdentityProviderConfiguration   bool
	_workmailDeleteImpersonationRole               bool
	_workmailDeleteMailboxPermissions              bool
	_workmailDeleteMobileDeviceAccessOverride      bool
	_workmailDeleteMobileDeviceAccessRule          bool
	_workmailDeleteOrganization                    bool
	_workmailDeletePersonalAccessToken             bool
	_workmailDeleteResource                        bool
	_workmailDeleteRetentionPolicy                 bool
	_workmailDeleteUser                            bool
	_workmailDeregisterFromWorkMail                bool
	_workmailDeregisterMailDomain                  bool
	_workmailDescribeEmailMonitoringConfiguration  bool
	_workmailDescribeEntity                        bool
	_workmailDescribeGroup                         bool
	_workmailDescribeIdentityProviderConfiguration bool
	_workmailDescribeInboundDmarcSettings          bool
	_workmailDescribeMailboxExportJob              bool
	_workmailDescribeOrganization                  bool
	_workmailDescribeResource                      bool
	_workmailDescribeUser                          bool
	_workmailDisassociateDelegateFromResource      bool
	_workmailDisassociateMemberFromGroup           bool
	_workmailGetAccessControlEffect                bool
	_workmailGetDefaultRetentionPolicy             bool
	_workmailGetImpersonationRole                  bool
	_workmailGetImpersonationRoleEffect            bool
	_workmailGetMailDomain                         bool
	_workmailGetMailboxDetails                     bool
	_workmailGetMobileDeviceAccessEffect           bool
	_workmailGetMobileDeviceAccessOverride         bool
	_workmailGetPersonalAccessTokenMetadata        bool
	_workmailListAccessControlRules                bool
	_workmailListAliases                           bool
	_workmailListAvailabilityConfigurations        bool
	_workmailListGroupMembers                      bool
	_workmailListGroups                            bool
	_workmailListGroupsForEntity                   bool
	_workmailListImpersonationRoles                bool
	_workmailListMailDomains                       bool
	_workmailListMailboxExportJobs                 bool
	_workmailListMailboxPermissions                bool
	_workmailListMobileDeviceAccessOverrides       bool
	_workmailListMobileDeviceAccessRules           bool
	_workmailListOrganizations                     bool
	_workmailListPersonalAccessTokens              bool
	_workmailListResourceDelegates                 bool
	_workmailListResources                         bool
	_workmailListTagsForResource                   bool
	_workmailListUsers                             bool
	_workmailPutAccessControlRule                  bool
	_workmailPutEmailMonitoringConfiguration       bool
	_workmailPutIdentityProviderConfiguration      bool
	_workmailPutInboundDmarcSettings               bool
	_workmailPutMailboxPermissions                 bool
	_workmailPutMobileDeviceAccessOverride         bool
	_workmailPutRetentionPolicy                    bool
	_workmailRegisterMailDomain                    bool
	_workmailRegisterToWorkMail                    bool
	_workmailResetPassword                         bool
	_workmailStartMailboxExportJob                 bool
	_workmailTagResource                           bool
	_workmailTestAvailabilityConfiguration         bool
	_workmailUntagResource                         bool
	_workmailUpdateAvailabilityConfiguration       bool
	_workmailUpdateDefaultMailDomain               bool
	_workmailUpdateGroup                           bool
	_workmailUpdateImpersonationRole               bool
	_workmailUpdateMailboxQuota                    bool
	_workmailUpdateMobileDeviceAccessRule          bool
	_workmailUpdatePrimaryEmailAddress             bool
	_workmailUpdateResource                        bool
	_workmailUpdateUser                            bool

	_workmailAction                           string
	_workmailActions                          []string
	_workmailAlias                            string
	_workmailApplicationArn                   string
	_workmailAuthenticationMode               string
	_workmailBookingOptions                   string
	_workmailCity                             string
	_workmailClientToken                      string
	_workmailCompany                          string
	_workmailCountry                          string
	_workmailDeleteDirectory                  string
	_workmailDepartment                       string
	_workmailDescription                      string
	_workmailDeviceId                         string
	_workmailDeviceModel                      string
	_workmailDeviceModels                     []string
	_workmailDeviceOperatingSystem            string
	_workmailDeviceOperatingSystems           []string
	_workmailDeviceType                       string
	_workmailDeviceTypes                      []string
	_workmailDeviceUserAgent                  string
	_workmailDeviceUserAgents                 []string
	_workmailDirectoryId                      string
	_workmailDisplayName                      string
	_workmailDomainName                       string
	_workmailDomains                          string
	_workmailEffect                           string
	_workmailEmail                            string
	_workmailEnableInteroperability           string
	_workmailEnforced                         string
	_workmailEntityId                         string
	_workmailEwsProvider                      string
	_workmailFilters                          string
	_workmailFirstName                        string
	_workmailFolderConfigurations             string
	_workmailForceDelete                      string
	_workmailGranteeId                        string
	_workmailGroupId                          string
	_workmailHiddenFromGlobalAddressList      string
	_workmailId                               string
	_workmailIdentityCenterConfiguration      string
	_workmailIdentityProviderUserId           string
	_workmailImpersonationRoleId              string
	_workmailImpersonationRoleIds             []string
	_workmailInitials                         string
	_workmailInstanceArn                      string
	_workmailIpAddress                        string
	_workmailIpRanges                         []string
	_workmailJobId                            string
	_workmailJobTitle                         string
	_workmailKmsKeyArn                        string
	_workmailLambdaProvider                   string
	_workmailLastName                         string
	_workmailLogGroupArn                      string
	_workmailMailboxQuota                     string
	_workmailMaxResults                       string
	_workmailMemberId                         string
	_workmailMobileDeviceAccessRuleId         string
	_workmailName                             string
	_workmailNextToken                        string
	_workmailNotActions                       []string
	_workmailNotDeviceModels                  []string
	_workmailNotDeviceOperatingSystems        []string
	_workmailNotDeviceTypes                   []string
	_workmailNotDeviceUserAgents              []string
	_workmailNotImpersonationRoleIds          []string
	_workmailNotIpRanges                      []string
	_workmailNotUserIds                       []string
	_workmailOffice                           string
	_workmailOrganizationId                   string
	_workmailPassword                         string
	_workmailPermissionValues                 string
	_workmailPersonalAccessTokenConfiguration string
	_workmailPersonalAccessTokenId            string
	_workmailResourceARN                      string
	_workmailResourceId                       string
	_workmailRole                             string
	_workmailRoleArn                          string
	_workmailRules                            string
	_workmailS3BucketName                     string
	_workmailS3Prefix                         string
	_workmailStreet                           string
	_workmailTagKeys                          []string
	_workmailTags                             string
	_workmailTargetUser                       string
	_workmailTelephone                        string
	_workmailType                             string
	_workmailUserId                           string
	_workmailUserIds                          []string
	_workmailZipCode                          string
)

// Adds a member (user or group) to the resource's set of delegates.
func workmail_AssociateDelegateToResource(cfg aws.Config, client *workmail.Client) {
	input := &workmail.AssociateDelegateToResourceInput{
		// EntityId: *string, // Required
		// OrganizationId: *string, // Required
		// ResourceId: *string, // Required
	}

	if len(_workmailEntityId) > 0 {
		input.EntityId = aws.String(_workmailEntityId)
	}
	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}
	if len(_workmailResourceId) > 0 {
		input.ResourceId = aws.String(_workmailResourceId)
	}

	if resp, err := client.AssociateDelegateToResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds a member (user or group) to the group's set.
func workmail_AssociateMemberToGroup(cfg aws.Config, client *workmail.Client) {
	input := &workmail.AssociateMemberToGroupInput{
		// GroupId: *string, // Required
		// MemberId: *string, // Required
		// OrganizationId: *string, // Required
	}

	if len(_workmailGroupId) > 0 {
		input.GroupId = aws.String(_workmailGroupId)
	}
	if len(_workmailMemberId) > 0 {
		input.MemberId = aws.String(_workmailMemberId)
	}
	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}

	if resp, err := client.AssociateMemberToGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Assumes an impersonation role for the given WorkMail organization. This method
// returns an authentication token you can use to make impersonated calls.
func workmail_AssumeImpersonationRole(cfg aws.Config, client *workmail.Client) {
	input := &workmail.AssumeImpersonationRoleInput{
		// ImpersonationRoleId: *string, // Required
		// OrganizationId: *string, // Required
	}

	if len(_workmailImpersonationRoleId) > 0 {
		input.ImpersonationRoleId = aws.String(_workmailImpersonationRoleId)
	}
	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}

	if resp, err := client.AssumeImpersonationRole(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels a mailbox export job.
// If the mailbox export job is near completion, it might not be possible to
// cancel it.
func workmail_CancelMailboxExportJob(cfg aws.Config, client *workmail.Client) {
	input := &workmail.CancelMailboxExportJobInput{
		// ClientToken: *string, // Required
		// JobId: *string, // Required
		// OrganizationId: *string, // Required
	}

	if len(_workmailClientToken) > 0 {
		input.ClientToken = aws.String(_workmailClientToken)
	}
	if len(_workmailJobId) > 0 {
		input.JobId = aws.String(_workmailJobId)
	}
	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}

	if resp, err := client.CancelMailboxExportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds an alias to the set of a given member (user or group) of WorkMail.
func workmail_CreateAlias(cfg aws.Config, client *workmail.Client) {
	input := &workmail.CreateAliasInput{
		// Alias: *string, // Required
		// EntityId: *string, // Required
		// OrganizationId: *string, // Required
	}

	if len(_workmailAlias) > 0 {
		input.Alias = aws.String(_workmailAlias)
	}
	if len(_workmailEntityId) > 0 {
		input.EntityId = aws.String(_workmailEntityId)
	}
	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}

	if resp, err := client.CreateAlias(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an AvailabilityConfiguration for the given WorkMail organization and
// domain.
func workmail_CreateAvailabilityConfiguration(cfg aws.Config, client *workmail.Client) {
	input := &workmail.CreateAvailabilityConfigurationInput{
		// DomainName: *string, // Required
		// OrganizationId: *string, // Required
	}

	if len(_workmailDomainName) > 0 {
		input.DomainName = aws.String(_workmailDomainName)
	}
	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}
	if len(_workmailClientToken) > 0 {
		input.ClientToken = aws.String(_workmailClientToken)
	}
	if len(_workmailEwsProvider) > 0 {
		if err := assignInputField(input, "EwsProvider", _workmailEwsProvider); err != nil {
			log.Errorf("invalid --ews-provider: %s", err.Error())
			return
		}
	}
	if len(_workmailLambdaProvider) > 0 {
		if err := assignInputField(input, "LambdaProvider", _workmailLambdaProvider); err != nil {
			log.Errorf("invalid --lambda-provider: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAvailabilityConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a group that can be used in WorkMail by calling the RegisterToWorkMail operation.
func workmail_CreateGroup(cfg aws.Config, client *workmail.Client) {
	input := &workmail.CreateGroupInput{
		// Name: *string, // Required
		// OrganizationId: *string, // Required
	}

	if len(_workmailName) > 0 {
		input.Name = aws.String(_workmailName)
	}
	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}
	if len(_workmailHiddenFromGlobalAddressList) > 0 {
		if err := assignInputField(input, "HiddenFromGlobalAddressList", _workmailHiddenFromGlobalAddressList); err != nil {
			log.Errorf("invalid --hidden-from-global-address-list: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates the WorkMail application in IAM Identity Center that can be used later
// in the WorkMail - IdC integration. For more information, see
// PutIdentityProviderConfiguration. This action does not affect the authentication
// settings for any WorkMail organizations.
func workmail_CreateIdentityCenterApplication(cfg aws.Config, client *workmail.Client) {
	input := &workmail.CreateIdentityCenterApplicationInput{
		// InstanceArn: *string, // Required
		// Name: *string, // Required
	}

	if len(_workmailInstanceArn) > 0 {
		input.InstanceArn = aws.String(_workmailInstanceArn)
	}
	if len(_workmailName) > 0 {
		input.Name = aws.String(_workmailName)
	}
	if len(_workmailClientToken) > 0 {
		input.ClientToken = aws.String(_workmailClientToken)
	}

	if resp, err := client.CreateIdentityCenterApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an impersonation role for the given WorkMail organization.
// Idempotency ensures that an API request completes no more than one time. With
// an idempotent request, if the original request completes successfully, any
// subsequent retries also complete successfully without performing any further
// actions.
func workmail_CreateImpersonationRole(cfg aws.Config, client *workmail.Client) {
	input := &workmail.CreateImpersonationRoleInput{
		// Name: *string, // Required
		// OrganizationId: *string, // Required
		// Rules: []types.ImpersonationRule, // Required
		// Type: types.ImpersonationRoleType, // Required
	}

	if len(_workmailName) > 0 {
		input.Name = aws.String(_workmailName)
	}
	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}
	if len(_workmailRules) > 0 {
		if err := assignInputField(input, "Rules", _workmailRules); err != nil {
			log.Errorf("invalid --rules: %s", err.Error())
			return
		}
	}
	if len(_workmailType) > 0 {
		if err := assignInputField(input, "Type", _workmailType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_workmailClientToken) > 0 {
		input.ClientToken = aws.String(_workmailClientToken)
	}
	if len(_workmailDescription) > 0 {
		input.Description = aws.String(_workmailDescription)
	}

	if resp, err := client.CreateImpersonationRole(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new mobile device access rule for the specified WorkMail organization.
func workmail_CreateMobileDeviceAccessRule(cfg aws.Config, client *workmail.Client) {
	input := &workmail.CreateMobileDeviceAccessRuleInput{
		// Effect: types.MobileDeviceAccessRuleEffect, // Required
		// Name: *string, // Required
		// OrganizationId: *string, // Required
	}

	if len(_workmailEffect) > 0 {
		if err := assignInputField(input, "Effect", _workmailEffect); err != nil {
			log.Errorf("invalid --effect: %s", err.Error())
			return
		}
	}
	if len(_workmailName) > 0 {
		input.Name = aws.String(_workmailName)
	}
	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}
	if len(_workmailClientToken) > 0 {
		input.ClientToken = aws.String(_workmailClientToken)
	}
	if len(_workmailDescription) > 0 {
		input.Description = aws.String(_workmailDescription)
	}
	if len(_workmailDeviceModels) > 0 {
		input.DeviceModels = append([]string(nil), _workmailDeviceModels...)
	}
	if len(_workmailDeviceOperatingSystems) > 0 {
		input.DeviceOperatingSystems = append([]string(nil), _workmailDeviceOperatingSystems...)
	}
	if len(_workmailDeviceTypes) > 0 {
		input.DeviceTypes = append([]string(nil), _workmailDeviceTypes...)
	}
	if len(_workmailDeviceUserAgents) > 0 {
		input.DeviceUserAgents = append([]string(nil), _workmailDeviceUserAgents...)
	}
	if len(_workmailNotDeviceModels) > 0 {
		input.NotDeviceModels = append([]string(nil), _workmailNotDeviceModels...)
	}
	if len(_workmailNotDeviceOperatingSystems) > 0 {
		input.NotDeviceOperatingSystems = append([]string(nil), _workmailNotDeviceOperatingSystems...)
	}
	if len(_workmailNotDeviceTypes) > 0 {
		input.NotDeviceTypes = append([]string(nil), _workmailNotDeviceTypes...)
	}
	if len(_workmailNotDeviceUserAgents) > 0 {
		input.NotDeviceUserAgents = append([]string(nil), _workmailNotDeviceUserAgents...)
	}

	if resp, err := client.CreateMobileDeviceAccessRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new WorkMail organization. Optionally, you can choose to associate an
// existing AWS Directory Service directory with your organization. If an AWS
// Directory Service directory ID is specified, the organization alias must match
// the directory alias. If you choose not to associate an existing directory with
// your organization, then we create a new WorkMail directory for you. For more
// information, see [Adding an organization]in the WorkMail Administrator Guide.
//
// You can associate multiple email domains with an organization, then choose your
// default email domain from the WorkMail console. You can also associate a domain
// that is managed in an Amazon Route 53 public hosted zone. For more information,
// see [Adding a domain]and [Choosing the default domain] in the WorkMail Administrator Guide.
//
// Optionally, you can use a customer managed key from AWS Key Management Service
// (AWS KMS) to encrypt email for your organization. If you don't associate an AWS
// KMS key, WorkMail creates a default, AWS managed key for you.
//
// [Adding an organization]: https://docs.aws.amazon.com/workmail/latest/adminguide/add_new_organization.html
// [Adding a domain]: https://docs.aws.amazon.com/workmail/latest/adminguide/add_domain.html
// [Choosing the default domain]: https://docs.aws.amazon.com/workmail/latest/adminguide/default_domain.html
func workmail_CreateOrganization(cfg aws.Config, client *workmail.Client) {
	input := &workmail.CreateOrganizationInput{
		// Alias: *string, // Required
	}

	if len(_workmailAlias) > 0 {
		input.Alias = aws.String(_workmailAlias)
	}
	if len(_workmailClientToken) > 0 {
		input.ClientToken = aws.String(_workmailClientToken)
	}
	if len(_workmailDirectoryId) > 0 {
		input.DirectoryId = aws.String(_workmailDirectoryId)
	}
	if len(_workmailDomains) > 0 {
		if err := assignInputField(input, "Domains", _workmailDomains); err != nil {
			log.Errorf("invalid --domains: %s", err.Error())
			return
		}
	}
	if len(_workmailEnableInteroperability) > 0 {
		if err := assignInputField(input, "EnableInteroperability", _workmailEnableInteroperability); err != nil {
			log.Errorf("invalid --enable-interoperability: %s", err.Error())
			return
		}
	}
	if len(_workmailKmsKeyArn) > 0 {
		input.KmsKeyArn = aws.String(_workmailKmsKeyArn)
	}

	if resp, err := client.CreateOrganization(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new WorkMail resource.
func workmail_CreateResource(cfg aws.Config, client *workmail.Client) {
	input := &workmail.CreateResourceInput{
		// Name: *string, // Required
		// OrganizationId: *string, // Required
		// Type: types.ResourceType, // Required
	}

	if len(_workmailName) > 0 {
		input.Name = aws.String(_workmailName)
	}
	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}
	if len(_workmailType) > 0 {
		if err := assignInputField(input, "Type", _workmailType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_workmailDescription) > 0 {
		input.Description = aws.String(_workmailDescription)
	}
	if len(_workmailHiddenFromGlobalAddressList) > 0 {
		if err := assignInputField(input, "HiddenFromGlobalAddressList", _workmailHiddenFromGlobalAddressList); err != nil {
			log.Errorf("invalid --hidden-from-global-address-list: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a user who can be used in WorkMail by calling the RegisterToWorkMail operation.
func workmail_CreateUser(cfg aws.Config, client *workmail.Client) {
	input := &workmail.CreateUserInput{
		// DisplayName: *string, // Required
		// Name: *string, // Required
		// OrganizationId: *string, // Required
	}

	if len(_workmailDisplayName) > 0 {
		input.DisplayName = aws.String(_workmailDisplayName)
	}
	if len(_workmailName) > 0 {
		input.Name = aws.String(_workmailName)
	}
	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}
	if len(_workmailFirstName) > 0 {
		input.FirstName = aws.String(_workmailFirstName)
	}
	if len(_workmailHiddenFromGlobalAddressList) > 0 {
		if err := assignInputField(input, "HiddenFromGlobalAddressList", _workmailHiddenFromGlobalAddressList); err != nil {
			log.Errorf("invalid --hidden-from-global-address-list: %s", err.Error())
			return
		}
	}
	if len(_workmailIdentityProviderUserId) > 0 {
		input.IdentityProviderUserId = aws.String(_workmailIdentityProviderUserId)
	}
	if len(_workmailLastName) > 0 {
		input.LastName = aws.String(_workmailLastName)
	}
	if len(_workmailPassword) > 0 {
		input.Password = aws.String(_workmailPassword)
	}
	if len(_workmailRole) > 0 {
		if err := assignInputField(input, "Role", _workmailRole); err != nil {
			log.Errorf("invalid --role: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an access control rule for the specified WorkMail organization.
// Deleting already deleted and non-existing rules does not produce an error. In
// those cases, the service sends back an HTTP 200 response with an empty HTTP
// body.
func workmail_DeleteAccessControlRule(cfg aws.Config, client *workmail.Client) {
	input := &workmail.DeleteAccessControlRuleInput{
		// Name: *string, // Required
		// OrganizationId: *string, // Required
	}

	if len(_workmailName) > 0 {
		input.Name = aws.String(_workmailName)
	}
	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}

	if resp, err := client.DeleteAccessControlRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Remove one or more specified aliases from a set of aliases for a given user.
func workmail_DeleteAlias(cfg aws.Config, client *workmail.Client) {
	input := &workmail.DeleteAliasInput{
		// Alias: *string, // Required
		// EntityId: *string, // Required
		// OrganizationId: *string, // Required
	}

	if len(_workmailAlias) > 0 {
		input.Alias = aws.String(_workmailAlias)
	}
	if len(_workmailEntityId) > 0 {
		input.EntityId = aws.String(_workmailEntityId)
	}
	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}

	if resp, err := client.DeleteAlias(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the AvailabilityConfiguration for the given WorkMail organization and
// domain.
func workmail_DeleteAvailabilityConfiguration(cfg aws.Config, client *workmail.Client) {
	input := &workmail.DeleteAvailabilityConfigurationInput{
		// DomainName: *string, // Required
		// OrganizationId: *string, // Required
	}

	if len(_workmailDomainName) > 0 {
		input.DomainName = aws.String(_workmailDomainName)
	}
	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}

	if resp, err := client.DeleteAvailabilityConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the email monitoring configuration for a specified organization.
func workmail_DeleteEmailMonitoringConfiguration(cfg aws.Config, client *workmail.Client) {
	input := &workmail.DeleteEmailMonitoringConfigurationInput{
		// OrganizationId: *string, // Required
	}

	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}

	if resp, err := client.DeleteEmailMonitoringConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a group from WorkMail.
func workmail_DeleteGroup(cfg aws.Config, client *workmail.Client) {
	input := &workmail.DeleteGroupInput{
		// GroupId: *string, // Required
		// OrganizationId: *string, // Required
	}

	if len(_workmailGroupId) > 0 {
		input.GroupId = aws.String(_workmailGroupId)
	}
	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}

	if resp, err := client.DeleteGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the IAM Identity Center application from WorkMail. This action does
// not affect the authentication settings for any WorkMail organizations.
func workmail_DeleteIdentityCenterApplication(cfg aws.Config, client *workmail.Client) {
	input := &workmail.DeleteIdentityCenterApplicationInput{
		// ApplicationArn: *string, // Required
	}

	if len(_workmailApplicationArn) > 0 {
		input.ApplicationArn = aws.String(_workmailApplicationArn)
	}

	if resp, err := client.DeleteIdentityCenterApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disables the integration between IdC and WorkMail. Authentication will
// continue with the directory as it was before the IdC integration. You might have
// to reset your directory passwords and reconfigure your desktop and mobile email
// clients.
func workmail_DeleteIdentityProviderConfiguration(cfg aws.Config, client *workmail.Client) {
	input := &workmail.DeleteIdentityProviderConfigurationInput{
		// OrganizationId: *string, // Required
	}

	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}

	if resp, err := client.DeleteIdentityProviderConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an impersonation role for the given WorkMail organization.
func workmail_DeleteImpersonationRole(cfg aws.Config, client *workmail.Client) {
	input := &workmail.DeleteImpersonationRoleInput{
		// ImpersonationRoleId: *string, // Required
		// OrganizationId: *string, // Required
	}

	if len(_workmailImpersonationRoleId) > 0 {
		input.ImpersonationRoleId = aws.String(_workmailImpersonationRoleId)
	}
	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}

	if resp, err := client.DeleteImpersonationRole(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes permissions granted to a member (user or group).
func workmail_DeleteMailboxPermissions(cfg aws.Config, client *workmail.Client) {
	input := &workmail.DeleteMailboxPermissionsInput{
		// EntityId: *string, // Required
		// GranteeId: *string, // Required
		// OrganizationId: *string, // Required
	}

	if len(_workmailEntityId) > 0 {
		input.EntityId = aws.String(_workmailEntityId)
	}
	if len(_workmailGranteeId) > 0 {
		input.GranteeId = aws.String(_workmailGranteeId)
	}
	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}

	if resp, err := client.DeleteMailboxPermissions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the mobile device access override for the given WorkMail organization,
// user, and device.
//
// Deleting already deleted and non-existing overrides does not produce an error.
// In those cases, the service sends back an HTTP 200 response with an empty HTTP
// body.
func workmail_DeleteMobileDeviceAccessOverride(cfg aws.Config, client *workmail.Client) {
	input := &workmail.DeleteMobileDeviceAccessOverrideInput{
		// DeviceId: *string, // Required
		// OrganizationId: *string, // Required
		// UserId: *string, // Required
	}

	if len(_workmailDeviceId) > 0 {
		input.DeviceId = aws.String(_workmailDeviceId)
	}
	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}
	if len(_workmailUserId) > 0 {
		input.UserId = aws.String(_workmailUserId)
	}

	if resp, err := client.DeleteMobileDeviceAccessOverride(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a mobile device access rule for the specified WorkMail organization.
// Deleting already deleted and non-existing rules does not produce an error. In
// those cases, the service sends back an HTTP 200 response with an empty HTTP
// body.
func workmail_DeleteMobileDeviceAccessRule(cfg aws.Config, client *workmail.Client) {
	input := &workmail.DeleteMobileDeviceAccessRuleInput{
		// MobileDeviceAccessRuleId: *string, // Required
		// OrganizationId: *string, // Required
	}

	if len(_workmailMobileDeviceAccessRuleId) > 0 {
		input.MobileDeviceAccessRuleId = aws.String(_workmailMobileDeviceAccessRuleId)
	}
	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}

	if resp, err := client.DeleteMobileDeviceAccessRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an WorkMail organization and all underlying AWS resources managed by
// WorkMail as part of the organization. You can choose whether to delete the
// associated directory. For more information, see [Removing an organization]in the WorkMail Administrator
// Guide.
//
// [Removing an organization]: https://docs.aws.amazon.com/workmail/latest/adminguide/remove_organization.html
func workmail_DeleteOrganization(cfg aws.Config, client *workmail.Client) {
	input := &workmail.DeleteOrganizationInput{
		// DeleteDirectory: bool, // Required
		// OrganizationId: *string, // Required
	}

	if len(_workmailDeleteDirectory) > 0 {
		if err := assignInputField(input, "DeleteDirectory", _workmailDeleteDirectory); err != nil {
			log.Errorf("invalid --delete-directory: %s", err.Error())
			return
		}
	}
	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}
	if len(_workmailClientToken) > 0 {
		input.ClientToken = aws.String(_workmailClientToken)
	}
	if len(_workmailForceDelete) > 0 {
		if err := assignInputField(input, "ForceDelete", _workmailForceDelete); err != nil {
			log.Errorf("invalid --force-delete: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteOrganization(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the Personal Access Token from the provided WorkMail Organization.
func workmail_DeletePersonalAccessToken(cfg aws.Config, client *workmail.Client) {
	input := &workmail.DeletePersonalAccessTokenInput{
		// OrganizationId: *string, // Required
		// PersonalAccessTokenId: *string, // Required
	}

	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}
	if len(_workmailPersonalAccessTokenId) > 0 {
		input.PersonalAccessTokenId = aws.String(_workmailPersonalAccessTokenId)
	}

	if resp, err := client.DeletePersonalAccessToken(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified resource.
func workmail_DeleteResource(cfg aws.Config, client *workmail.Client) {
	input := &workmail.DeleteResourceInput{
		// OrganizationId: *string, // Required
		// ResourceId: *string, // Required
	}

	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}
	if len(_workmailResourceId) > 0 {
		input.ResourceId = aws.String(_workmailResourceId)
	}

	if resp, err := client.DeleteResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified retention policy from the specified organization.
func workmail_DeleteRetentionPolicy(cfg aws.Config, client *workmail.Client) {
	input := &workmail.DeleteRetentionPolicyInput{
		// Id: *string, // Required
		// OrganizationId: *string, // Required
	}

	if len(_workmailId) > 0 {
		input.Id = aws.String(_workmailId)
	}
	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}

	if resp, err := client.DeleteRetentionPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a user from WorkMail and all subsequent systems. Before you can delete
// a user, the user state must be DISABLED . Use the DescribeUser action to confirm the user
// state.
//
// Deleting a user is permanent and cannot be undone. WorkMail archives user
// mailboxes for 30 days before they are permanently removed.
func workmail_DeleteUser(cfg aws.Config, client *workmail.Client) {
	input := &workmail.DeleteUserInput{
		// OrganizationId: *string, // Required
		// UserId: *string, // Required
	}

	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}
	if len(_workmailUserId) > 0 {
		input.UserId = aws.String(_workmailUserId)
	}

	if resp, err := client.DeleteUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Mark a user, group, or resource as no longer used in WorkMail. This action
// disassociates the mailbox and schedules it for clean-up. WorkMail keeps
// mailboxes for 30 days before they are permanently removed. The functionality in
// the console is Disable.
func workmail_DeregisterFromWorkMail(cfg aws.Config, client *workmail.Client) {
	input := &workmail.DeregisterFromWorkMailInput{
		// EntityId: *string, // Required
		// OrganizationId: *string, // Required
	}

	if len(_workmailEntityId) > 0 {
		input.EntityId = aws.String(_workmailEntityId)
	}
	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}

	if resp, err := client.DeregisterFromWorkMail(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a domain from WorkMail, stops email routing to WorkMail, and removes
// the authorization allowing WorkMail use. SES keeps the domain because other
// applications may use it. You must first remove any email address used by
// WorkMail entities before you remove the domain.
func workmail_DeregisterMailDomain(cfg aws.Config, client *workmail.Client) {
	input := &workmail.DeregisterMailDomainInput{
		// DomainName: *string, // Required
		// OrganizationId: *string, // Required
	}

	if len(_workmailDomainName) > 0 {
		input.DomainName = aws.String(_workmailDomainName)
	}
	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}

	if resp, err := client.DeregisterMailDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the current email monitoring configuration for a specified
// organization.
func workmail_DescribeEmailMonitoringConfiguration(cfg aws.Config, client *workmail.Client) {
	input := &workmail.DescribeEmailMonitoringConfigurationInput{
		// OrganizationId: *string, // Required
	}

	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}

	if resp, err := client.DescribeEmailMonitoringConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns basic details about an entity in WorkMail.
func workmail_DescribeEntity(cfg aws.Config, client *workmail.Client) {
	input := &workmail.DescribeEntityInput{
		// Email: *string, // Required
		// OrganizationId: *string, // Required
	}

	if len(_workmailEmail) > 0 {
		input.Email = aws.String(_workmailEmail)
	}
	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}

	if resp, err := client.DescribeEntity(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the data available for the group.
func workmail_DescribeGroup(cfg aws.Config, client *workmail.Client) {
	input := &workmail.DescribeGroupInput{
		// GroupId: *string, // Required
		// OrganizationId: *string, // Required
	}

	if len(_workmailGroupId) > 0 {
		input.GroupId = aws.String(_workmailGroupId)
	}
	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}

	if resp, err := client.DescribeGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns detailed information on the current IdC setup for the WorkMail
// organization.
func workmail_DescribeIdentityProviderConfiguration(cfg aws.Config, client *workmail.Client) {
	input := &workmail.DescribeIdentityProviderConfigurationInput{
		// OrganizationId: *string, // Required
	}

	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}

	if resp, err := client.DescribeIdentityProviderConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the settings in a DMARC policy for a specified organization.
func workmail_DescribeInboundDmarcSettings(cfg aws.Config, client *workmail.Client) {
	input := &workmail.DescribeInboundDmarcSettingsInput{
		// OrganizationId: *string, // Required
	}

	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}

	if resp, err := client.DescribeInboundDmarcSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the current status of a mailbox export job.
func workmail_DescribeMailboxExportJob(cfg aws.Config, client *workmail.Client) {
	input := &workmail.DescribeMailboxExportJobInput{
		// JobId: *string, // Required
		// OrganizationId: *string, // Required
	}

	if len(_workmailJobId) > 0 {
		input.JobId = aws.String(_workmailJobId)
	}
	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}

	if resp, err := client.DescribeMailboxExportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides more information regarding a given organization based on its
// identifier.
func workmail_DescribeOrganization(cfg aws.Config, client *workmail.Client) {
	input := &workmail.DescribeOrganizationInput{
		// OrganizationId: *string, // Required
	}

	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}

	if resp, err := client.DescribeOrganization(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the data available for the resource.
func workmail_DescribeResource(cfg aws.Config, client *workmail.Client) {
	input := &workmail.DescribeResourceInput{
		// OrganizationId: *string, // Required
		// ResourceId: *string, // Required
	}

	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}
	if len(_workmailResourceId) > 0 {
		input.ResourceId = aws.String(_workmailResourceId)
	}

	if resp, err := client.DescribeResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides information regarding the user.
func workmail_DescribeUser(cfg aws.Config, client *workmail.Client) {
	input := &workmail.DescribeUserInput{
		// OrganizationId: *string, // Required
		// UserId: *string, // Required
	}

	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}
	if len(_workmailUserId) > 0 {
		input.UserId = aws.String(_workmailUserId)
	}

	if resp, err := client.DescribeUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a member from the resource's set of delegates.
func workmail_DisassociateDelegateFromResource(cfg aws.Config, client *workmail.Client) {
	input := &workmail.DisassociateDelegateFromResourceInput{
		// EntityId: *string, // Required
		// OrganizationId: *string, // Required
		// ResourceId: *string, // Required
	}

	if len(_workmailEntityId) > 0 {
		input.EntityId = aws.String(_workmailEntityId)
	}
	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}
	if len(_workmailResourceId) > 0 {
		input.ResourceId = aws.String(_workmailResourceId)
	}

	if resp, err := client.DisassociateDelegateFromResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a member from a group.
func workmail_DisassociateMemberFromGroup(cfg aws.Config, client *workmail.Client) {
	input := &workmail.DisassociateMemberFromGroupInput{
		// GroupId: *string, // Required
		// MemberId: *string, // Required
		// OrganizationId: *string, // Required
	}

	if len(_workmailGroupId) > 0 {
		input.GroupId = aws.String(_workmailGroupId)
	}
	if len(_workmailMemberId) > 0 {
		input.MemberId = aws.String(_workmailMemberId)
	}
	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}

	if resp, err := client.DisassociateMemberFromGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the effects of an organization's access control rules as they apply to a
// specified IPv4 address, access protocol action, and user ID or impersonation
// role ID. You must provide either the user ID or impersonation role ID.
// Impersonation role ID can only be used with Action EWS.
func workmail_GetAccessControlEffect(cfg aws.Config, client *workmail.Client) {
	input := &workmail.GetAccessControlEffectInput{
		// Action: *string, // Required
		// IpAddress: *string, // Required
		// OrganizationId: *string, // Required
	}

	if len(_workmailAction) > 0 {
		input.Action = aws.String(_workmailAction)
	}
	if len(_workmailIpAddress) > 0 {
		input.IpAddress = aws.String(_workmailIpAddress)
	}
	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}
	if len(_workmailImpersonationRoleId) > 0 {
		input.ImpersonationRoleId = aws.String(_workmailImpersonationRoleId)
	}
	if len(_workmailUserId) > 0 {
		input.UserId = aws.String(_workmailUserId)
	}

	if resp, err := client.GetAccessControlEffect(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the default retention policy details for the specified organization.
func workmail_GetDefaultRetentionPolicy(cfg aws.Config, client *workmail.Client) {
	input := &workmail.GetDefaultRetentionPolicyInput{
		// OrganizationId: *string, // Required
	}

	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}

	if resp, err := client.GetDefaultRetentionPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the impersonation role details for the given WorkMail organization.
func workmail_GetImpersonationRole(cfg aws.Config, client *workmail.Client) {
	input := &workmail.GetImpersonationRoleInput{
		// ImpersonationRoleId: *string, // Required
		// OrganizationId: *string, // Required
	}

	if len(_workmailImpersonationRoleId) > 0 {
		input.ImpersonationRoleId = aws.String(_workmailImpersonationRoleId)
	}
	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}

	if resp, err := client.GetImpersonationRole(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Tests whether the given impersonation role can impersonate a target user.
func workmail_GetImpersonationRoleEffect(cfg aws.Config, client *workmail.Client) {
	input := &workmail.GetImpersonationRoleEffectInput{
		// ImpersonationRoleId: *string, // Required
		// OrganizationId: *string, // Required
		// TargetUser: *string, // Required
	}

	if len(_workmailImpersonationRoleId) > 0 {
		input.ImpersonationRoleId = aws.String(_workmailImpersonationRoleId)
	}
	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}
	if len(_workmailTargetUser) > 0 {
		input.TargetUser = aws.String(_workmailTargetUser)
	}

	if resp, err := client.GetImpersonationRoleEffect(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets details for a mail domain, including domain records required to configure
// your domain with recommended security.
func workmail_GetMailDomain(cfg aws.Config, client *workmail.Client) {
	input := &workmail.GetMailDomainInput{
		// DomainName: *string, // Required
		// OrganizationId: *string, // Required
	}

	if len(_workmailDomainName) > 0 {
		input.DomainName = aws.String(_workmailDomainName)
	}
	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}

	if resp, err := client.GetMailDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Requests a user's mailbox details for a specified organization and user.
func workmail_GetMailboxDetails(cfg aws.Config, client *workmail.Client) {
	input := &workmail.GetMailboxDetailsInput{
		// OrganizationId: *string, // Required
		// UserId: *string, // Required
	}

	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}
	if len(_workmailUserId) > 0 {
		input.UserId = aws.String(_workmailUserId)
	}

	if resp, err := client.GetMailboxDetails(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Simulates the effect of the mobile device access rules for the given attributes
// of a sample access event. Use this method to test the effects of the current set
// of mobile device access rules for the WorkMail organization for a particular
// user's attributes.
func workmail_GetMobileDeviceAccessEffect(cfg aws.Config, client *workmail.Client) {
	input := &workmail.GetMobileDeviceAccessEffectInput{
		// OrganizationId: *string, // Required
	}

	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}
	if len(_workmailDeviceModel) > 0 {
		input.DeviceModel = aws.String(_workmailDeviceModel)
	}
	if len(_workmailDeviceOperatingSystem) > 0 {
		input.DeviceOperatingSystem = aws.String(_workmailDeviceOperatingSystem)
	}
	if len(_workmailDeviceType) > 0 {
		input.DeviceType = aws.String(_workmailDeviceType)
	}
	if len(_workmailDeviceUserAgent) > 0 {
		input.DeviceUserAgent = aws.String(_workmailDeviceUserAgent)
	}

	if resp, err := client.GetMobileDeviceAccessEffect(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the mobile device access override for the given WorkMail organization,
// user, and device.
func workmail_GetMobileDeviceAccessOverride(cfg aws.Config, client *workmail.Client) {
	input := &workmail.GetMobileDeviceAccessOverrideInput{
		// DeviceId: *string, // Required
		// OrganizationId: *string, // Required
		// UserId: *string, // Required
	}

	if len(_workmailDeviceId) > 0 {
		input.DeviceId = aws.String(_workmailDeviceId)
	}
	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}
	if len(_workmailUserId) > 0 {
		input.UserId = aws.String(_workmailUserId)
	}

	if resp, err := client.GetMobileDeviceAccessOverride(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Requests details of a specific Personal Access Token within the WorkMail
// organization.
func workmail_GetPersonalAccessTokenMetadata(cfg aws.Config, client *workmail.Client) {
	input := &workmail.GetPersonalAccessTokenMetadataInput{
		// OrganizationId: *string, // Required
		// PersonalAccessTokenId: *string, // Required
	}

	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}
	if len(_workmailPersonalAccessTokenId) > 0 {
		input.PersonalAccessTokenId = aws.String(_workmailPersonalAccessTokenId)
	}

	if resp, err := client.GetPersonalAccessTokenMetadata(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the access control rules for the specified organization.
func workmail_ListAccessControlRules(cfg aws.Config, client *workmail.Client) {
	input := &workmail.ListAccessControlRulesInput{
		// OrganizationId: *string, // Required
	}

	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}

	if resp, err := client.ListAccessControlRules(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a paginated call to list the aliases associated with a given entity.
func workmail_ListAliases(cfg aws.Config, client *workmail.Client) {
	input := &workmail.ListAliasesInput{
		// EntityId: *string, // Required
		// OrganizationId: *string, // Required
	}

	if len(_workmailEntityId) > 0 {
		input.EntityId = aws.String(_workmailEntityId)
	}
	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}
	if len(_workmailMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _workmailMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_workmailNextToken) > 0 {
		input.NextToken = aws.String(_workmailNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAliases(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*workmail.ListAliasesOutput
	p := workmail.NewListAliasesPaginator(client, input)
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

// List all the AvailabilityConfiguration 's for the given WorkMail organization.
func workmail_ListAvailabilityConfigurations(cfg aws.Config, client *workmail.Client) {
	input := &workmail.ListAvailabilityConfigurationsInput{
		// OrganizationId: *string, // Required
	}

	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}
	if len(_workmailMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _workmailMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_workmailNextToken) > 0 {
		input.NextToken = aws.String(_workmailNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAvailabilityConfigurations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*workmail.ListAvailabilityConfigurationsOutput
	p := workmail.NewListAvailabilityConfigurationsPaginator(client, input)
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

// Returns an overview of the members of a group. Users and groups can be members
// of a group.
func workmail_ListGroupMembers(cfg aws.Config, client *workmail.Client) {
	input := &workmail.ListGroupMembersInput{
		// GroupId: *string, // Required
		// OrganizationId: *string, // Required
	}

	if len(_workmailGroupId) > 0 {
		input.GroupId = aws.String(_workmailGroupId)
	}
	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}
	if len(_workmailMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _workmailMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_workmailNextToken) > 0 {
		input.NextToken = aws.String(_workmailNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListGroupMembers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*workmail.ListGroupMembersOutput
	p := workmail.NewListGroupMembersPaginator(client, input)
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

// Returns summaries of the organization's groups.
func workmail_ListGroups(cfg aws.Config, client *workmail.Client) {
	input := &workmail.ListGroupsInput{
		// OrganizationId: *string, // Required
	}

	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}
	if len(_workmailFilters) > 0 {
		if err := assignInputField(input, "Filters", _workmailFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_workmailMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _workmailMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_workmailNextToken) > 0 {
		input.NextToken = aws.String(_workmailNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*workmail.ListGroupsOutput
	p := workmail.NewListGroupsPaginator(client, input)
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

// Returns all the groups to which an entity belongs.
func workmail_ListGroupsForEntity(cfg aws.Config, client *workmail.Client) {
	input := &workmail.ListGroupsForEntityInput{
		// EntityId: *string, // Required
		// OrganizationId: *string, // Required
	}

	if len(_workmailEntityId) > 0 {
		input.EntityId = aws.String(_workmailEntityId)
	}
	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}
	if len(_workmailFilters) > 0 {
		if err := assignInputField(input, "Filters", _workmailFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_workmailMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _workmailMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_workmailNextToken) > 0 {
		input.NextToken = aws.String(_workmailNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListGroupsForEntity(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*workmail.ListGroupsForEntityOutput
	p := workmail.NewListGroupsForEntityPaginator(client, input)
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

// Lists all the impersonation roles for the given WorkMail organization.
func workmail_ListImpersonationRoles(cfg aws.Config, client *workmail.Client) {
	input := &workmail.ListImpersonationRolesInput{
		// OrganizationId: *string, // Required
	}

	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}
	if len(_workmailMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _workmailMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_workmailNextToken) > 0 {
		input.NextToken = aws.String(_workmailNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListImpersonationRoles(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*workmail.ListImpersonationRolesOutput
	p := workmail.NewListImpersonationRolesPaginator(client, input)
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

// Lists the mail domains in a given WorkMail organization.
func workmail_ListMailDomains(cfg aws.Config, client *workmail.Client) {
	input := &workmail.ListMailDomainsInput{
		// OrganizationId: *string, // Required
	}

	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}
	if len(_workmailMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _workmailMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_workmailNextToken) > 0 {
		input.NextToken = aws.String(_workmailNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListMailDomains(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*workmail.ListMailDomainsOutput
	p := workmail.NewListMailDomainsPaginator(client, input)
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

// Lists the mailbox export jobs started for the specified organization within the
// last seven days.
func workmail_ListMailboxExportJobs(cfg aws.Config, client *workmail.Client) {
	input := &workmail.ListMailboxExportJobsInput{
		// OrganizationId: *string, // Required
	}

	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}
	if len(_workmailMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _workmailMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_workmailNextToken) > 0 {
		input.NextToken = aws.String(_workmailNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListMailboxExportJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*workmail.ListMailboxExportJobsOutput
	p := workmail.NewListMailboxExportJobsPaginator(client, input)
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

// Lists the mailbox permissions associated with a user, group, or resource
// mailbox.
func workmail_ListMailboxPermissions(cfg aws.Config, client *workmail.Client) {
	input := &workmail.ListMailboxPermissionsInput{
		// EntityId: *string, // Required
		// OrganizationId: *string, // Required
	}

	if len(_workmailEntityId) > 0 {
		input.EntityId = aws.String(_workmailEntityId)
	}
	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}
	if len(_workmailMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _workmailMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_workmailNextToken) > 0 {
		input.NextToken = aws.String(_workmailNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListMailboxPermissions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*workmail.ListMailboxPermissionsOutput
	p := workmail.NewListMailboxPermissionsPaginator(client, input)
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

// Lists all the mobile device access overrides for any given combination of
// WorkMail organization, user, or device.
func workmail_ListMobileDeviceAccessOverrides(cfg aws.Config, client *workmail.Client) {
	input := &workmail.ListMobileDeviceAccessOverridesInput{
		// OrganizationId: *string, // Required
	}

	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}
	if len(_workmailDeviceId) > 0 {
		input.DeviceId = aws.String(_workmailDeviceId)
	}
	if len(_workmailMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _workmailMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_workmailNextToken) > 0 {
		input.NextToken = aws.String(_workmailNextToken)
	}
	if len(_workmailUserId) > 0 {
		input.UserId = aws.String(_workmailUserId)
	}

	if disablePaginator() {
		if resp, err := client.ListMobileDeviceAccessOverrides(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*workmail.ListMobileDeviceAccessOverridesOutput
	p := workmail.NewListMobileDeviceAccessOverridesPaginator(client, input)
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

// Lists the mobile device access rules for the specified WorkMail organization.
func workmail_ListMobileDeviceAccessRules(cfg aws.Config, client *workmail.Client) {
	input := &workmail.ListMobileDeviceAccessRulesInput{
		// OrganizationId: *string, // Required
	}

	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}

	if resp, err := client.ListMobileDeviceAccessRules(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns summaries of the customer's organizations.
func workmail_ListOrganizations(cfg aws.Config, client *workmail.Client) {
	input := &workmail.ListOrganizationsInput{}

	if len(_workmailMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _workmailMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_workmailNextToken) > 0 {
		input.NextToken = aws.String(_workmailNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListOrganizations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*workmail.ListOrganizationsOutput
	p := workmail.NewListOrganizationsPaginator(client, input)
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

// Returns a summary of your Personal Access Tokens.
func workmail_ListPersonalAccessTokens(cfg aws.Config, client *workmail.Client) {
	input := &workmail.ListPersonalAccessTokensInput{
		// OrganizationId: *string, // Required
	}

	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}
	if len(_workmailMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _workmailMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_workmailNextToken) > 0 {
		input.NextToken = aws.String(_workmailNextToken)
	}
	if len(_workmailUserId) > 0 {
		input.UserId = aws.String(_workmailUserId)
	}

	if disablePaginator() {
		if resp, err := client.ListPersonalAccessTokens(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*workmail.ListPersonalAccessTokensOutput
	p := workmail.NewListPersonalAccessTokensPaginator(client, input)
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

// Lists the delegates associated with a resource. Users and groups can be
// resource delegates and answer requests on behalf of the resource.
func workmail_ListResourceDelegates(cfg aws.Config, client *workmail.Client) {
	input := &workmail.ListResourceDelegatesInput{
		// OrganizationId: *string, // Required
		// ResourceId: *string, // Required
	}

	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}
	if len(_workmailResourceId) > 0 {
		input.ResourceId = aws.String(_workmailResourceId)
	}
	if len(_workmailMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _workmailMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_workmailNextToken) > 0 {
		input.NextToken = aws.String(_workmailNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListResourceDelegates(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*workmail.ListResourceDelegatesOutput
	p := workmail.NewListResourceDelegatesPaginator(client, input)
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

// Returns summaries of the organization's resources.
func workmail_ListResources(cfg aws.Config, client *workmail.Client) {
	input := &workmail.ListResourcesInput{
		// OrganizationId: *string, // Required
	}

	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}
	if len(_workmailFilters) > 0 {
		if err := assignInputField(input, "Filters", _workmailFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_workmailMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _workmailMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_workmailNextToken) > 0 {
		input.NextToken = aws.String(_workmailNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListResources(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*workmail.ListResourcesOutput
	p := workmail.NewListResourcesPaginator(client, input)
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

// Lists the tags applied to an WorkMail organization resource.
func workmail_ListTagsForResource(cfg aws.Config, client *workmail.Client) {
	input := &workmail.ListTagsForResourceInput{
		// ResourceARN: *string, // Required
	}

	if len(_workmailResourceARN) > 0 {
		input.ResourceARN = aws.String(_workmailResourceARN)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns summaries of the organization's users.
func workmail_ListUsers(cfg aws.Config, client *workmail.Client) {
	input := &workmail.ListUsersInput{
		// OrganizationId: *string, // Required
	}

	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}
	if len(_workmailFilters) > 0 {
		if err := assignInputField(input, "Filters", _workmailFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_workmailMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _workmailMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_workmailNextToken) > 0 {
		input.NextToken = aws.String(_workmailNextToken)
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

	var results []*workmail.ListUsersOutput
	p := workmail.NewListUsersPaginator(client, input)
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

// Adds a new access control rule for the specified organization. The rule allows
// or denies access to the organization for the specified IPv4 addresses, access
// protocol actions, user IDs and impersonation IDs. Adding a new rule with the
// same name as an existing rule replaces the older rule.
func workmail_PutAccessControlRule(cfg aws.Config, client *workmail.Client) {
	input := &workmail.PutAccessControlRuleInput{
		// Description: *string, // Required
		// Effect: types.AccessControlRuleEffect, // Required
		// Name: *string, // Required
		// OrganizationId: *string, // Required
	}

	if len(_workmailDescription) > 0 {
		input.Description = aws.String(_workmailDescription)
	}
	if len(_workmailEffect) > 0 {
		if err := assignInputField(input, "Effect", _workmailEffect); err != nil {
			log.Errorf("invalid --effect: %s", err.Error())
			return
		}
	}
	if len(_workmailName) > 0 {
		input.Name = aws.String(_workmailName)
	}
	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}
	if len(_workmailActions) > 0 {
		input.Actions = append([]string(nil), _workmailActions...)
	}
	if len(_workmailImpersonationRoleIds) > 0 {
		input.ImpersonationRoleIds = append([]string(nil), _workmailImpersonationRoleIds...)
	}
	if len(_workmailIpRanges) > 0 {
		input.IpRanges = append([]string(nil), _workmailIpRanges...)
	}
	if len(_workmailNotActions) > 0 {
		input.NotActions = append([]string(nil), _workmailNotActions...)
	}
	if len(_workmailNotImpersonationRoleIds) > 0 {
		input.NotImpersonationRoleIds = append([]string(nil), _workmailNotImpersonationRoleIds...)
	}
	if len(_workmailNotIpRanges) > 0 {
		input.NotIpRanges = append([]string(nil), _workmailNotIpRanges...)
	}
	if len(_workmailNotUserIds) > 0 {
		input.NotUserIds = append([]string(nil), _workmailNotUserIds...)
	}
	if len(_workmailUserIds) > 0 {
		input.UserIds = append([]string(nil), _workmailUserIds...)
	}

	if resp, err := client.PutAccessControlRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates the email monitoring configuration for a specified
// organization.
func workmail_PutEmailMonitoringConfiguration(cfg aws.Config, client *workmail.Client) {
	input := &workmail.PutEmailMonitoringConfigurationInput{
		// LogGroupArn: *string, // Required
		// OrganizationId: *string, // Required
	}

	if len(_workmailLogGroupArn) > 0 {
		input.LogGroupArn = aws.String(_workmailLogGroupArn)
	}
	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}
	if len(_workmailRoleArn) > 0 {
		input.RoleArn = aws.String(_workmailRoleArn)
	}

	if resp, err := client.PutEmailMonitoringConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables integration between IAM Identity Center (IdC) and WorkMail to proxy
// authentication requests for mailbox users. You can connect your IdC directory or
// your external directory to WorkMail through IdC and manage access to WorkMail
// mailboxes in a single place. For enhanced protection, you could enable
// Multifactor Authentication (MFA) and Personal Access Tokens.
func workmail_PutIdentityProviderConfiguration(cfg aws.Config, client *workmail.Client) {
	input := &workmail.PutIdentityProviderConfigurationInput{
		// AuthenticationMode: types.IdentityProviderAuthenticationMode, // Required
		// IdentityCenterConfiguration: *types.IdentityCenterConfiguration, // Required
		// OrganizationId: *string, // Required
		// PersonalAccessTokenConfiguration: *types.PersonalAccessTokenConfiguration, // Required
	}

	if len(_workmailAuthenticationMode) > 0 {
		if err := assignInputField(input, "AuthenticationMode", _workmailAuthenticationMode); err != nil {
			log.Errorf("invalid --authentication-mode: %s", err.Error())
			return
		}
	}
	if len(_workmailIdentityCenterConfiguration) > 0 {
		if err := assignInputField(input, "IdentityCenterConfiguration", _workmailIdentityCenterConfiguration); err != nil {
			log.Errorf("invalid --identity-center-configuration: %s", err.Error())
			return
		}
	}
	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}
	if len(_workmailPersonalAccessTokenConfiguration) > 0 {
		if err := assignInputField(input, "PersonalAccessTokenConfiguration", _workmailPersonalAccessTokenConfiguration); err != nil {
			log.Errorf("invalid --personal-access-token-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutIdentityProviderConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables or disables a DMARC policy for a given organization.
func workmail_PutInboundDmarcSettings(cfg aws.Config, client *workmail.Client) {
	input := &workmail.PutInboundDmarcSettingsInput{
		// Enforced: *bool, // Required
		// OrganizationId: *string, // Required
	}

	if len(_workmailEnforced) > 0 {
		if err := assignInputField(input, "Enforced", _workmailEnforced); err != nil {
			log.Errorf("invalid --enforced: %s", err.Error())
			return
		}
	}
	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}

	if resp, err := client.PutInboundDmarcSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets permissions for a user, group, or resource. This replaces any pre-existing
// permissions.
func workmail_PutMailboxPermissions(cfg aws.Config, client *workmail.Client) {
	input := &workmail.PutMailboxPermissionsInput{
		// EntityId: *string, // Required
		// GranteeId: *string, // Required
		// OrganizationId: *string, // Required
		// PermissionValues: []types.PermissionType, // Required
	}

	if len(_workmailEntityId) > 0 {
		input.EntityId = aws.String(_workmailEntityId)
	}
	if len(_workmailGranteeId) > 0 {
		input.GranteeId = aws.String(_workmailGranteeId)
	}
	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}
	if len(_workmailPermissionValues) > 0 {
		if err := assignInputField(input, "PermissionValues", _workmailPermissionValues); err != nil {
			log.Errorf("invalid --permission-values: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutMailboxPermissions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates a mobile device access override for the given WorkMail
// organization, user, and device.
func workmail_PutMobileDeviceAccessOverride(cfg aws.Config, client *workmail.Client) {
	input := &workmail.PutMobileDeviceAccessOverrideInput{
		// DeviceId: *string, // Required
		// Effect: types.MobileDeviceAccessRuleEffect, // Required
		// OrganizationId: *string, // Required
		// UserId: *string, // Required
	}

	if len(_workmailDeviceId) > 0 {
		input.DeviceId = aws.String(_workmailDeviceId)
	}
	if len(_workmailEffect) > 0 {
		if err := assignInputField(input, "Effect", _workmailEffect); err != nil {
			log.Errorf("invalid --effect: %s", err.Error())
			return
		}
	}
	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}
	if len(_workmailUserId) > 0 {
		input.UserId = aws.String(_workmailUserId)
	}
	if len(_workmailDescription) > 0 {
		input.Description = aws.String(_workmailDescription)
	}

	if resp, err := client.PutMobileDeviceAccessOverride(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Puts a retention policy to the specified organization.
func workmail_PutRetentionPolicy(cfg aws.Config, client *workmail.Client) {
	input := &workmail.PutRetentionPolicyInput{
		// FolderConfigurations: []types.FolderConfiguration, // Required
		// Name: *string, // Required
		// OrganizationId: *string, // Required
	}

	if len(_workmailFolderConfigurations) > 0 {
		if err := assignInputField(input, "FolderConfigurations", _workmailFolderConfigurations); err != nil {
			log.Errorf("invalid --folder-configurations: %s", err.Error())
			return
		}
	}
	if len(_workmailName) > 0 {
		input.Name = aws.String(_workmailName)
	}
	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}
	if len(_workmailDescription) > 0 {
		input.Description = aws.String(_workmailDescription)
	}
	if len(_workmailId) > 0 {
		input.Id = aws.String(_workmailId)
	}

	if resp, err := client.PutRetentionPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Registers a new domain in WorkMail and SES, and configures it for use by
// WorkMail. Emails received by SES for this domain are routed to the specified
// WorkMail organization, and WorkMail has permanent permission to use the
// specified domain for sending your users' emails.
func workmail_RegisterMailDomain(cfg aws.Config, client *workmail.Client) {
	input := &workmail.RegisterMailDomainInput{
		// DomainName: *string, // Required
		// OrganizationId: *string, // Required
	}

	if len(_workmailDomainName) > 0 {
		input.DomainName = aws.String(_workmailDomainName)
	}
	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}
	if len(_workmailClientToken) > 0 {
		input.ClientToken = aws.String(_workmailClientToken)
	}

	if resp, err := client.RegisterMailDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Registers an existing and disabled user, group, or resource for WorkMail use by
// associating a mailbox and calendaring capabilities. It performs no change if the
// user, group, or resource is enabled and fails if the user, group, or resource is
// deleted. This operation results in the accumulation of costs. For more
// information, see [Pricing]. The equivalent console functionality for this operation is
// Enable.
//
// Users can either be created by calling the CreateUser API operation or they can be
// synchronized from your directory. For more information, see DeregisterFromWorkMail.
//
// [Pricing]: https://aws.amazon.com/workmail/pricing
func workmail_RegisterToWorkMail(cfg aws.Config, client *workmail.Client) {
	input := &workmail.RegisterToWorkMailInput{
		// Email: *string, // Required
		// EntityId: *string, // Required
		// OrganizationId: *string, // Required
	}

	if len(_workmailEmail) > 0 {
		input.Email = aws.String(_workmailEmail)
	}
	if len(_workmailEntityId) > 0 {
		input.EntityId = aws.String(_workmailEntityId)
	}
	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}

	if resp, err := client.RegisterToWorkMail(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Allows the administrator to reset the password for a user.
func workmail_ResetPassword(cfg aws.Config, client *workmail.Client) {
	input := &workmail.ResetPasswordInput{
		// OrganizationId: *string, // Required
		// Password: *string, // Required
		// UserId: *string, // Required
	}

	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}
	if len(_workmailPassword) > 0 {
		input.Password = aws.String(_workmailPassword)
	}
	if len(_workmailUserId) > 0 {
		input.UserId = aws.String(_workmailUserId)
	}

	if resp, err := client.ResetPassword(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a mailbox export job to export MIME-format email messages and calendar
// items from the specified mailbox to the specified Amazon Simple Storage Service
// (Amazon S3) bucket. For more information, see [Exporting mailbox content]in the WorkMail Administrator
// Guide.
//
// [Exporting mailbox content]: https://docs.aws.amazon.com/workmail/latest/adminguide/mail-export.html
func workmail_StartMailboxExportJob(cfg aws.Config, client *workmail.Client) {
	input := &workmail.StartMailboxExportJobInput{
		// ClientToken: *string, // Required
		// EntityId: *string, // Required
		// KmsKeyArn: *string, // Required
		// OrganizationId: *string, // Required
		// RoleArn: *string, // Required
		// S3BucketName: *string, // Required
		// S3Prefix: *string, // Required
	}

	if len(_workmailClientToken) > 0 {
		input.ClientToken = aws.String(_workmailClientToken)
	}
	if len(_workmailEntityId) > 0 {
		input.EntityId = aws.String(_workmailEntityId)
	}
	if len(_workmailKmsKeyArn) > 0 {
		input.KmsKeyArn = aws.String(_workmailKmsKeyArn)
	}
	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}
	if len(_workmailRoleArn) > 0 {
		input.RoleArn = aws.String(_workmailRoleArn)
	}
	if len(_workmailS3BucketName) > 0 {
		input.S3BucketName = aws.String(_workmailS3BucketName)
	}
	if len(_workmailS3Prefix) > 0 {
		input.S3Prefix = aws.String(_workmailS3Prefix)
	}
	if len(_workmailDescription) > 0 {
		input.Description = aws.String(_workmailDescription)
	}

	if resp, err := client.StartMailboxExportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Applies the specified tags to the specified WorkMailorganization resource.
func workmail_TagResource(cfg aws.Config, client *workmail.Client) {
	input := &workmail.TagResourceInput{
		// ResourceARN: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_workmailResourceARN) > 0 {
		input.ResourceARN = aws.String(_workmailResourceARN)
	}
	if len(_workmailTags) > 0 {
		if err := assignInputField(input, "Tags", _workmailTags); err != nil {
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

// Performs a test on an availability provider to ensure that access is allowed.
// For EWS, it verifies the provided credentials can be used to successfully log
// in. For Lambda, it verifies that the Lambda function can be invoked and that the
// resource access policy was configured to deny anonymous access. An anonymous
// invocation is one done without providing either a SourceArn or SourceAccount
// header.
//
// The request must contain either one provider definition ( EwsProvider or
// LambdaProvider ) or the DomainName parameter. If the DomainName parameter is
// provided, the configuration stored under the DomainName will be tested.
func workmail_TestAvailabilityConfiguration(cfg aws.Config, client *workmail.Client) {
	input := &workmail.TestAvailabilityConfigurationInput{
		// OrganizationId: *string, // Required
	}

	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}
	if len(_workmailDomainName) > 0 {
		input.DomainName = aws.String(_workmailDomainName)
	}
	if len(_workmailEwsProvider) > 0 {
		if err := assignInputField(input, "EwsProvider", _workmailEwsProvider); err != nil {
			log.Errorf("invalid --ews-provider: %s", err.Error())
			return
		}
	}
	if len(_workmailLambdaProvider) > 0 {
		if err := assignInputField(input, "LambdaProvider", _workmailLambdaProvider); err != nil {
			log.Errorf("invalid --lambda-provider: %s", err.Error())
			return
		}
	}

	if resp, err := client.TestAvailabilityConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Untags the specified tags from the specified WorkMail organization resource.
func workmail_UntagResource(cfg aws.Config, client *workmail.Client) {
	input := &workmail.UntagResourceInput{
		// ResourceARN: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_workmailResourceARN) > 0 {
		input.ResourceARN = aws.String(_workmailResourceARN)
	}
	if len(_workmailTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _workmailTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing AvailabilityConfiguration for the given WorkMail
// organization and domain.
func workmail_UpdateAvailabilityConfiguration(cfg aws.Config, client *workmail.Client) {
	input := &workmail.UpdateAvailabilityConfigurationInput{
		// DomainName: *string, // Required
		// OrganizationId: *string, // Required
	}

	if len(_workmailDomainName) > 0 {
		input.DomainName = aws.String(_workmailDomainName)
	}
	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}
	if len(_workmailEwsProvider) > 0 {
		if err := assignInputField(input, "EwsProvider", _workmailEwsProvider); err != nil {
			log.Errorf("invalid --ews-provider: %s", err.Error())
			return
		}
	}
	if len(_workmailLambdaProvider) > 0 {
		if err := assignInputField(input, "LambdaProvider", _workmailLambdaProvider); err != nil {
			log.Errorf("invalid --lambda-provider: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateAvailabilityConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the default mail domain for an organization. The default mail domain is
// used by the WorkMail AWS Console to suggest an email address when enabling a
// mail user. You can only have one default domain.
func workmail_UpdateDefaultMailDomain(cfg aws.Config, client *workmail.Client) {
	input := &workmail.UpdateDefaultMailDomainInput{
		// DomainName: *string, // Required
		// OrganizationId: *string, // Required
	}

	if len(_workmailDomainName) > 0 {
		input.DomainName = aws.String(_workmailDomainName)
	}
	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}

	if resp, err := client.UpdateDefaultMailDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates attributes in a group.
func workmail_UpdateGroup(cfg aws.Config, client *workmail.Client) {
	input := &workmail.UpdateGroupInput{
		// GroupId: *string, // Required
		// OrganizationId: *string, // Required
	}

	if len(_workmailGroupId) > 0 {
		input.GroupId = aws.String(_workmailGroupId)
	}
	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}
	if len(_workmailHiddenFromGlobalAddressList) > 0 {
		if err := assignInputField(input, "HiddenFromGlobalAddressList", _workmailHiddenFromGlobalAddressList); err != nil {
			log.Errorf("invalid --hidden-from-global-address-list: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an impersonation role for the given WorkMail organization.
func workmail_UpdateImpersonationRole(cfg aws.Config, client *workmail.Client) {
	input := &workmail.UpdateImpersonationRoleInput{
		// ImpersonationRoleId: *string, // Required
		// Name: *string, // Required
		// OrganizationId: *string, // Required
		// Rules: []types.ImpersonationRule, // Required
		// Type: types.ImpersonationRoleType, // Required
	}

	if len(_workmailImpersonationRoleId) > 0 {
		input.ImpersonationRoleId = aws.String(_workmailImpersonationRoleId)
	}
	if len(_workmailName) > 0 {
		input.Name = aws.String(_workmailName)
	}
	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}
	if len(_workmailRules) > 0 {
		if err := assignInputField(input, "Rules", _workmailRules); err != nil {
			log.Errorf("invalid --rules: %s", err.Error())
			return
		}
	}
	if len(_workmailType) > 0 {
		if err := assignInputField(input, "Type", _workmailType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_workmailDescription) > 0 {
		input.Description = aws.String(_workmailDescription)
	}

	if resp, err := client.UpdateImpersonationRole(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a user's current mailbox quota for a specified organization and user.
func workmail_UpdateMailboxQuota(cfg aws.Config, client *workmail.Client) {
	input := &workmail.UpdateMailboxQuotaInput{
		// MailboxQuota: *int32, // Required
		// OrganizationId: *string, // Required
		// UserId: *string, // Required
	}

	if len(_workmailMailboxQuota) > 0 {
		if err := assignInputField(input, "MailboxQuota", _workmailMailboxQuota); err != nil {
			log.Errorf("invalid --mailbox-quota: %s", err.Error())
			return
		}
	}
	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}
	if len(_workmailUserId) > 0 {
		input.UserId = aws.String(_workmailUserId)
	}

	if resp, err := client.UpdateMailboxQuota(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a mobile device access rule for the specified WorkMail organization.
func workmail_UpdateMobileDeviceAccessRule(cfg aws.Config, client *workmail.Client) {
	input := &workmail.UpdateMobileDeviceAccessRuleInput{
		// Effect: types.MobileDeviceAccessRuleEffect, // Required
		// MobileDeviceAccessRuleId: *string, // Required
		// Name: *string, // Required
		// OrganizationId: *string, // Required
	}

	if len(_workmailEffect) > 0 {
		if err := assignInputField(input, "Effect", _workmailEffect); err != nil {
			log.Errorf("invalid --effect: %s", err.Error())
			return
		}
	}
	if len(_workmailMobileDeviceAccessRuleId) > 0 {
		input.MobileDeviceAccessRuleId = aws.String(_workmailMobileDeviceAccessRuleId)
	}
	if len(_workmailName) > 0 {
		input.Name = aws.String(_workmailName)
	}
	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}
	if len(_workmailDescription) > 0 {
		input.Description = aws.String(_workmailDescription)
	}
	if len(_workmailDeviceModels) > 0 {
		input.DeviceModels = append([]string(nil), _workmailDeviceModels...)
	}
	if len(_workmailDeviceOperatingSystems) > 0 {
		input.DeviceOperatingSystems = append([]string(nil), _workmailDeviceOperatingSystems...)
	}
	if len(_workmailDeviceTypes) > 0 {
		input.DeviceTypes = append([]string(nil), _workmailDeviceTypes...)
	}
	if len(_workmailDeviceUserAgents) > 0 {
		input.DeviceUserAgents = append([]string(nil), _workmailDeviceUserAgents...)
	}
	if len(_workmailNotDeviceModels) > 0 {
		input.NotDeviceModels = append([]string(nil), _workmailNotDeviceModels...)
	}
	if len(_workmailNotDeviceOperatingSystems) > 0 {
		input.NotDeviceOperatingSystems = append([]string(nil), _workmailNotDeviceOperatingSystems...)
	}
	if len(_workmailNotDeviceTypes) > 0 {
		input.NotDeviceTypes = append([]string(nil), _workmailNotDeviceTypes...)
	}
	if len(_workmailNotDeviceUserAgents) > 0 {
		input.NotDeviceUserAgents = append([]string(nil), _workmailNotDeviceUserAgents...)
	}

	if resp, err := client.UpdateMobileDeviceAccessRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the primary email for a user, group, or resource. The current email is
// moved into the list of aliases (or swapped between an existing alias and the
// current primary email), and the email provided in the input is promoted as the
// primary.
func workmail_UpdatePrimaryEmailAddress(cfg aws.Config, client *workmail.Client) {
	input := &workmail.UpdatePrimaryEmailAddressInput{
		// Email: *string, // Required
		// EntityId: *string, // Required
		// OrganizationId: *string, // Required
	}

	if len(_workmailEmail) > 0 {
		input.Email = aws.String(_workmailEmail)
	}
	if len(_workmailEntityId) > 0 {
		input.EntityId = aws.String(_workmailEntityId)
	}
	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}

	if resp, err := client.UpdatePrimaryEmailAddress(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates data for the resource. To have the latest information, it must be
// preceded by a DescribeResourcecall. The dataset in the request should be the one expected when
// performing another DescribeResource call.
func workmail_UpdateResource(cfg aws.Config, client *workmail.Client) {
	input := &workmail.UpdateResourceInput{
		// OrganizationId: *string, // Required
		// ResourceId: *string, // Required
	}

	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}
	if len(_workmailResourceId) > 0 {
		input.ResourceId = aws.String(_workmailResourceId)
	}
	if len(_workmailBookingOptions) > 0 {
		if err := assignInputField(input, "BookingOptions", _workmailBookingOptions); err != nil {
			log.Errorf("invalid --booking-options: %s", err.Error())
			return
		}
	}
	if len(_workmailDescription) > 0 {
		input.Description = aws.String(_workmailDescription)
	}
	if len(_workmailHiddenFromGlobalAddressList) > 0 {
		if err := assignInputField(input, "HiddenFromGlobalAddressList", _workmailHiddenFromGlobalAddressList); err != nil {
			log.Errorf("invalid --hidden-from-global-address-list: %s", err.Error())
			return
		}
	}
	if len(_workmailName) > 0 {
		input.Name = aws.String(_workmailName)
	}
	if len(_workmailType) > 0 {
		if err := assignInputField(input, "Type", _workmailType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates data for the user. To have the latest information, it must be preceded
// by a DescribeUsercall. The dataset in the request should be the one expected when
// performing another DescribeUser call.
func workmail_UpdateUser(cfg aws.Config, client *workmail.Client) {
	input := &workmail.UpdateUserInput{
		// OrganizationId: *string, // Required
		// UserId: *string, // Required
	}

	if len(_workmailOrganizationId) > 0 {
		input.OrganizationId = aws.String(_workmailOrganizationId)
	}
	if len(_workmailUserId) > 0 {
		input.UserId = aws.String(_workmailUserId)
	}
	if len(_workmailCity) > 0 {
		input.City = aws.String(_workmailCity)
	}
	if len(_workmailCompany) > 0 {
		input.Company = aws.String(_workmailCompany)
	}
	if len(_workmailCountry) > 0 {
		input.Country = aws.String(_workmailCountry)
	}
	if len(_workmailDepartment) > 0 {
		input.Department = aws.String(_workmailDepartment)
	}
	if len(_workmailDisplayName) > 0 {
		input.DisplayName = aws.String(_workmailDisplayName)
	}
	if len(_workmailFirstName) > 0 {
		input.FirstName = aws.String(_workmailFirstName)
	}
	if len(_workmailHiddenFromGlobalAddressList) > 0 {
		if err := assignInputField(input, "HiddenFromGlobalAddressList", _workmailHiddenFromGlobalAddressList); err != nil {
			log.Errorf("invalid --hidden-from-global-address-list: %s", err.Error())
			return
		}
	}
	if len(_workmailIdentityProviderUserId) > 0 {
		input.IdentityProviderUserId = aws.String(_workmailIdentityProviderUserId)
	}
	if len(_workmailInitials) > 0 {
		input.Initials = aws.String(_workmailInitials)
	}
	if len(_workmailJobTitle) > 0 {
		input.JobTitle = aws.String(_workmailJobTitle)
	}
	if len(_workmailLastName) > 0 {
		input.LastName = aws.String(_workmailLastName)
	}
	if len(_workmailOffice) > 0 {
		input.Office = aws.String(_workmailOffice)
	}
	if len(_workmailRole) > 0 {
		if err := assignInputField(input, "Role", _workmailRole); err != nil {
			log.Errorf("invalid --role: %s", err.Error())
			return
		}
	}
	if len(_workmailStreet) > 0 {
		input.Street = aws.String(_workmailStreet)
	}
	if len(_workmailTelephone) > 0 {
		input.Telephone = aws.String(_workmailTelephone)
	}
	if len(_workmailZipCode) > 0 {
		input.ZipCode = aws.String(_workmailZipCode)
	}

	if resp, err := client.UpdateUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_workmailCmd)
	_workmailCmd.Flags().SortFlags = false

	_workmailCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_workmailCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_workmailCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_workmailCmd.Flags().StringVarP(&_workmailAction, "action", "", "", "Action")
	_workmailCmd.Flags().StringSliceVarP(&_workmailActions, "actions", "", nil, "Actions")
	_workmailCmd.Flags().StringVarP(&_workmailAlias, "alias", "", "", "Alias")
	_workmailCmd.Flags().StringVarP(&_workmailApplicationArn, "application-arn", "", "", "Application ARN")
	_workmailCmd.Flags().StringVarP(&_workmailAuthenticationMode, "authentication-mode", "", "", "Authentication Mode")
	_workmailCmd.Flags().StringVarP(&_workmailBookingOptions, "booking-options", "", "", "Booking Options")
	_workmailCmd.Flags().StringVarP(&_workmailCity, "city", "", "", "City")
	_workmailCmd.Flags().StringVarP(&_workmailClientToken, "client-token", "", "", "Client Token")
	_workmailCmd.Flags().StringVarP(&_workmailCompany, "company", "", "", "Company")
	_workmailCmd.Flags().StringVarP(&_workmailCountry, "country", "", "", "Country")
	_workmailCmd.Flags().StringVarP(&_workmailDeleteDirectory, "delete-directory", "", "", "Delete Directory")
	_workmailCmd.Flags().StringVarP(&_workmailDepartment, "department", "", "", "Department")
	_workmailCmd.Flags().StringVarP(&_workmailDescription, "description", "", "", "Description")
	_workmailCmd.Flags().StringVarP(&_workmailDeviceId, "device-id", "", "", "Device ID")
	_workmailCmd.Flags().StringVarP(&_workmailDeviceModel, "device-model", "", "", "Device Model")
	_workmailCmd.Flags().StringSliceVarP(&_workmailDeviceModels, "device-models", "", nil, "Device Models")
	_workmailCmd.Flags().StringVarP(&_workmailDeviceOperatingSystem, "device-operating-system", "", "", "Device Operating System")
	_workmailCmd.Flags().StringSliceVarP(&_workmailDeviceOperatingSystems, "device-operating-systems", "", nil, "Device Operating Systems")
	_workmailCmd.Flags().StringVarP(&_workmailDeviceType, "device-type", "", "", "Device Type")
	_workmailCmd.Flags().StringSliceVarP(&_workmailDeviceTypes, "device-types", "", nil, "Device Types")
	_workmailCmd.Flags().StringVarP(&_workmailDeviceUserAgent, "device-user-agent", "", "", "Device User Agent")
	_workmailCmd.Flags().StringSliceVarP(&_workmailDeviceUserAgents, "device-user-agents", "", nil, "Device User Agents")
	_workmailCmd.Flags().StringVarP(&_workmailDirectoryId, "directory-id", "", "", "Directory ID")
	_workmailCmd.Flags().StringVarP(&_workmailDisplayName, "display-name", "", "", "Display Name")
	_workmailCmd.Flags().StringVarP(&_workmailDomainName, "domain-name", "", "", "Domain Name")
	_workmailCmd.Flags().StringVarP(&_workmailDomains, "domains", "", "", "Domains")
	_workmailCmd.Flags().StringVarP(&_workmailEffect, "effect", "", "", "Effect")
	_workmailCmd.Flags().StringVarP(&_workmailEmail, "email", "", "", "Email")
	_workmailCmd.Flags().StringVarP(&_workmailEnableInteroperability, "enable-interoperability", "", "", "Enable Interoperability")
	_workmailCmd.Flags().StringVarP(&_workmailEnforced, "enforced", "", "", "Enforced")
	_workmailCmd.Flags().StringVarP(&_workmailEntityId, "entity-id", "", "", "Entity ID")
	_workmailCmd.Flags().StringVarP(&_workmailEwsProvider, "ews-provider", "", "", "Ews Provider")
	_workmailCmd.Flags().StringVarP(&_workmailFilters, "filters", "", "", "Filters")
	_workmailCmd.Flags().StringVarP(&_workmailFirstName, "first-name", "", "", "First Name")
	_workmailCmd.Flags().StringVarP(&_workmailFolderConfigurations, "folder-configurations", "", "", "Folder Configurations")
	_workmailCmd.Flags().StringVarP(&_workmailForceDelete, "force-delete", "", "", "Force Delete")
	_workmailCmd.Flags().StringVarP(&_workmailGranteeId, "grantee-id", "", "", "Grantee ID")
	_workmailCmd.Flags().StringVarP(&_workmailGroupId, "group-id", "", "", "Group ID")
	_workmailCmd.Flags().StringVarP(&_workmailHiddenFromGlobalAddressList, "hidden-from-global-address-list", "", "", "Hidden From Global Address List")
	_workmailCmd.Flags().StringVarP(&_workmailId, "id", "", "", "ID")
	_workmailCmd.Flags().StringVarP(&_workmailIdentityCenterConfiguration, "identity-center-configuration", "", "", "Identity Center Configuration")
	_workmailCmd.Flags().StringVarP(&_workmailIdentityProviderUserId, "identity-provider-user-id", "", "", "Identity Provider User ID")
	_workmailCmd.Flags().StringVarP(&_workmailImpersonationRoleId, "impersonation-role-id", "", "", "Impersonation Role ID")
	_workmailCmd.Flags().StringSliceVarP(&_workmailImpersonationRoleIds, "impersonation-role-ids", "", nil, "Impersonation Role Ids")
	_workmailCmd.Flags().StringVarP(&_workmailInitials, "initials", "", "", "Initials")
	_workmailCmd.Flags().StringVarP(&_workmailInstanceArn, "instance-arn", "", "", "Instance ARN")
	_workmailCmd.Flags().StringVarP(&_workmailIpAddress, "ip-address", "", "", "IP Address")
	_workmailCmd.Flags().StringSliceVarP(&_workmailIpRanges, "ip-ranges", "", nil, "IP Ranges")
	_workmailCmd.Flags().StringVarP(&_workmailJobId, "job-id", "", "", "Job ID")
	_workmailCmd.Flags().StringVarP(&_workmailJobTitle, "job-title", "", "", "Job Title")
	_workmailCmd.Flags().StringVarP(&_workmailKmsKeyArn, "kms-key-arn", "", "", "KMS Key ARN")
	_workmailCmd.Flags().StringVarP(&_workmailLambdaProvider, "lambda-provider", "", "", "Lambda Provider")
	_workmailCmd.Flags().StringVarP(&_workmailLastName, "last-name", "", "", "Last Name")
	_workmailCmd.Flags().StringVarP(&_workmailLogGroupArn, "log-group-arn", "", "", "Log Group ARN")
	_workmailCmd.Flags().StringVarP(&_workmailMailboxQuota, "mailbox-quota", "", "", "Mailbox Quota")
	_workmailCmd.Flags().StringVarP(&_workmailMaxResults, "max-results", "", "", "Max Results")
	_workmailCmd.Flags().StringVarP(&_workmailMemberId, "member-id", "", "", "Member ID")
	_workmailCmd.Flags().StringVarP(&_workmailMobileDeviceAccessRuleId, "mobile-device-access-rule-id", "", "", "Mobile Device Access Rule ID")
	_workmailCmd.Flags().StringVarP(&_workmailName, "name", "", "", "Name")
	_workmailCmd.Flags().StringVarP(&_workmailNextToken, "next-token", "", "", "Next Token")
	_workmailCmd.Flags().StringSliceVarP(&_workmailNotActions, "not-actions", "", nil, "Not Actions")
	_workmailCmd.Flags().StringSliceVarP(&_workmailNotDeviceModels, "not-device-models", "", nil, "Not Device Models")
	_workmailCmd.Flags().StringSliceVarP(&_workmailNotDeviceOperatingSystems, "not-device-operating-systems", "", nil, "Not Device Operating Systems")
	_workmailCmd.Flags().StringSliceVarP(&_workmailNotDeviceTypes, "not-device-types", "", nil, "Not Device Types")
	_workmailCmd.Flags().StringSliceVarP(&_workmailNotDeviceUserAgents, "not-device-user-agents", "", nil, "Not Device User Agents")
	_workmailCmd.Flags().StringSliceVarP(&_workmailNotImpersonationRoleIds, "not-impersonation-role-ids", "", nil, "Not Impersonation Role Ids")
	_workmailCmd.Flags().StringSliceVarP(&_workmailNotIpRanges, "not-ip-ranges", "", nil, "Not IP Ranges")
	_workmailCmd.Flags().StringSliceVarP(&_workmailNotUserIds, "not-user-ids", "", nil, "Not User Ids")
	_workmailCmd.Flags().StringVarP(&_workmailOffice, "office", "", "", "Office")
	_workmailCmd.Flags().StringVarP(&_workmailOrganizationId, "organization-id", "", "", "Organization ID")
	_workmailCmd.Flags().StringVarP(&_workmailPassword, "password", "", "", "Password")
	_workmailCmd.Flags().StringVarP(&_workmailPermissionValues, "permission-values", "", "", "Permission Values")
	_workmailCmd.Flags().StringVarP(&_workmailPersonalAccessTokenConfiguration, "personal-access-token-configuration", "", "", "Personal Access Token Configuration")
	_workmailCmd.Flags().StringVarP(&_workmailPersonalAccessTokenId, "personal-access-token-id", "", "", "Personal Access Token ID")
	_workmailCmd.Flags().StringVarP(&_workmailResourceARN, "resource-arn", "", "", "Resource ARN")
	_workmailCmd.Flags().StringVarP(&_workmailResourceId, "resource-id", "", "", "Resource ID")
	_workmailCmd.Flags().StringVarP(&_workmailRole, "role", "", "", "Role")
	_workmailCmd.Flags().StringVarP(&_workmailRoleArn, "role-arn", "", "", "Role ARN")
	_workmailCmd.Flags().StringVarP(&_workmailRules, "rules", "", "", "Rules")
	_workmailCmd.Flags().StringVarP(&_workmailS3BucketName, "s3-bucket-name", "", "", "S3 Bucket Name")
	_workmailCmd.Flags().StringVarP(&_workmailS3Prefix, "s3-prefix", "", "", "S3 Prefix")
	_workmailCmd.Flags().StringVarP(&_workmailStreet, "street", "", "", "Street")
	_workmailCmd.Flags().StringSliceVarP(&_workmailTagKeys, "tag-keys", "", nil, "Tag Keys")
	_workmailCmd.Flags().StringVarP(&_workmailTags, "tags", "", "", "Tags")
	_workmailCmd.Flags().StringVarP(&_workmailTargetUser, "target-user", "", "", "Target User")
	_workmailCmd.Flags().StringVarP(&_workmailTelephone, "telephone", "", "", "Telephone")
	_workmailCmd.Flags().StringVarP(&_workmailType, "type", "", "", "Type")
	_workmailCmd.Flags().StringVarP(&_workmailUserId, "user-id", "", "", "User ID")
	_workmailCmd.Flags().StringSliceVarP(&_workmailUserIds, "user-ids", "", nil, "User Ids")
	_workmailCmd.Flags().StringVarP(&_workmailZipCode, "zip-code", "", "", "Zip Code")

	_workmailCmd.Flags().BoolVarP(&_workmailAssociateDelegateToResource, "associate-delegate-to-resource", "", false, "Associate Delegate To Resource")
	_workmailCmd.Flags().BoolVarP(&_workmailAssociateMemberToGroup, "associate-member-to-group", "", false, "Associate Member To Group")
	_workmailCmd.Flags().BoolVarP(&_workmailAssumeImpersonationRole, "assume-impersonation-role", "", false, "Assume Impersonation Role")
	_workmailCmd.Flags().BoolVarP(&_workmailCancelMailboxExportJob, "cancel-mailbox-export-job", "", false, "Cancel Mailbox Export Job")
	_workmailCmd.Flags().BoolVarP(&_workmailCreateAlias, "create-alias", "", false, "Create Alias")
	_workmailCmd.Flags().BoolVarP(&_workmailCreateAvailabilityConfiguration, "create-availability-configuration", "", false, "Create Availability Configuration")
	_workmailCmd.Flags().BoolVarP(&_workmailCreateGroup, "create-group", "", false, "Create Group")
	_workmailCmd.Flags().BoolVarP(&_workmailCreateIdentityCenterApplication, "create-identity-center-application", "", false, "Create Identity Center Application")
	_workmailCmd.Flags().BoolVarP(&_workmailCreateImpersonationRole, "create-impersonation-role", "", false, "Create Impersonation Role")
	_workmailCmd.Flags().BoolVarP(&_workmailCreateMobileDeviceAccessRule, "create-mobile-device-access-rule", "", false, "Create Mobile Device Access Rule")
	_workmailCmd.Flags().BoolVarP(&_workmailCreateOrganization, "create-organization", "", false, "Create Organization")
	_workmailCmd.Flags().BoolVarP(&_workmailCreateResource, "create-resource", "", false, "Create Resource")
	_workmailCmd.Flags().BoolVarP(&_workmailCreateUser, "create-user", "", false, "Create User")
	_workmailCmd.Flags().BoolVarP(&_workmailDeleteAccessControlRule, "delete-access-control-rule", "", false, "Delete Access Control Rule")
	_workmailCmd.Flags().BoolVarP(&_workmailDeleteAlias, "delete-alias", "", false, "Delete Alias")
	_workmailCmd.Flags().BoolVarP(&_workmailDeleteAvailabilityConfiguration, "delete-availability-configuration", "", false, "Delete Availability Configuration")
	_workmailCmd.Flags().BoolVarP(&_workmailDeleteEmailMonitoringConfiguration, "delete-email-monitoring-configuration", "", false, "Delete Email Monitoring Configuration")
	_workmailCmd.Flags().BoolVarP(&_workmailDeleteGroup, "delete-group", "", false, "Delete Group")
	_workmailCmd.Flags().BoolVarP(&_workmailDeleteIdentityCenterApplication, "delete-identity-center-application", "", false, "Delete Identity Center Application")
	_workmailCmd.Flags().BoolVarP(&_workmailDeleteIdentityProviderConfiguration, "delete-identity-provider-configuration", "", false, "Delete Identity Provider Configuration")
	_workmailCmd.Flags().BoolVarP(&_workmailDeleteImpersonationRole, "delete-impersonation-role", "", false, "Delete Impersonation Role")
	_workmailCmd.Flags().BoolVarP(&_workmailDeleteMailboxPermissions, "delete-mailbox-permissions", "", false, "Delete Mailbox Permissions")
	_workmailCmd.Flags().BoolVarP(&_workmailDeleteMobileDeviceAccessOverride, "delete-mobile-device-access-override", "", false, "Delete Mobile Device Access Override")
	_workmailCmd.Flags().BoolVarP(&_workmailDeleteMobileDeviceAccessRule, "delete-mobile-device-access-rule", "", false, "Delete Mobile Device Access Rule")
	_workmailCmd.Flags().BoolVarP(&_workmailDeleteOrganization, "delete-organization", "", false, "Delete Organization")
	_workmailCmd.Flags().BoolVarP(&_workmailDeletePersonalAccessToken, "delete-personal-access-token", "", false, "Delete Personal Access Token")
	_workmailCmd.Flags().BoolVarP(&_workmailDeleteResource, "delete-resource", "", false, "Delete Resource")
	_workmailCmd.Flags().BoolVarP(&_workmailDeleteRetentionPolicy, "delete-retention-policy", "", false, "Delete Retention Policy")
	_workmailCmd.Flags().BoolVarP(&_workmailDeleteUser, "delete-user", "", false, "Delete User")
	_workmailCmd.Flags().BoolVarP(&_workmailDeregisterFromWorkMail, "deregister-from-work-mail", "", false, "Deregister From Work Mail")
	_workmailCmd.Flags().BoolVarP(&_workmailDeregisterMailDomain, "deregister-mail-domain", "", false, "Deregister Mail Domain")
	_workmailCmd.Flags().BoolVarP(&_workmailDescribeEmailMonitoringConfiguration, "describe-email-monitoring-configuration", "", false, "Describe Email Monitoring Configuration")
	_workmailCmd.Flags().BoolVarP(&_workmailDescribeEntity, "describe-entity", "", false, "Describe Entity")
	_workmailCmd.Flags().BoolVarP(&_workmailDescribeGroup, "describe-group", "", false, "Describe Group")
	_workmailCmd.Flags().BoolVarP(&_workmailDescribeIdentityProviderConfiguration, "describe-identity-provider-configuration", "", false, "Describe Identity Provider Configuration")
	_workmailCmd.Flags().BoolVarP(&_workmailDescribeInboundDmarcSettings, "describe-inbound-dmarc-settings", "", false, "Describe Inbound Dmarc Settings")
	_workmailCmd.Flags().BoolVarP(&_workmailDescribeMailboxExportJob, "describe-mailbox-export-job", "", false, "Describe Mailbox Export Job")
	_workmailCmd.Flags().BoolVarP(&_workmailDescribeOrganization, "describe-organization", "", false, "Describe Organization")
	_workmailCmd.Flags().BoolVarP(&_workmailDescribeResource, "describe-resource", "", false, "Describe Resource")
	_workmailCmd.Flags().BoolVarP(&_workmailDescribeUser, "describe-user", "", false, "Describe User")
	_workmailCmd.Flags().BoolVarP(&_workmailDisassociateDelegateFromResource, "disassociate-delegate-from-resource", "", false, "Disassociate Delegate From Resource")
	_workmailCmd.Flags().BoolVarP(&_workmailDisassociateMemberFromGroup, "disassociate-member-from-group", "", false, "Disassociate Member From Group")
	_workmailCmd.Flags().BoolVarP(&_workmailGetAccessControlEffect, "get-access-control-effect", "", false, "Get Access Control Effect")
	_workmailCmd.Flags().BoolVarP(&_workmailGetDefaultRetentionPolicy, "get-default-retention-policy", "", false, "Get Default Retention Policy")
	_workmailCmd.Flags().BoolVarP(&_workmailGetImpersonationRole, "get-impersonation-role", "", false, "Get Impersonation Role")
	_workmailCmd.Flags().BoolVarP(&_workmailGetImpersonationRoleEffect, "get-impersonation-role-effect", "", false, "Get Impersonation Role Effect")
	_workmailCmd.Flags().BoolVarP(&_workmailGetMailDomain, "get-mail-domain", "", false, "Get Mail Domain")
	_workmailCmd.Flags().BoolVarP(&_workmailGetMailboxDetails, "get-mailbox-details", "", false, "Get Mailbox Details")
	_workmailCmd.Flags().BoolVarP(&_workmailGetMobileDeviceAccessEffect, "get-mobile-device-access-effect", "", false, "Get Mobile Device Access Effect")
	_workmailCmd.Flags().BoolVarP(&_workmailGetMobileDeviceAccessOverride, "get-mobile-device-access-override", "", false, "Get Mobile Device Access Override")
	_workmailCmd.Flags().BoolVarP(&_workmailGetPersonalAccessTokenMetadata, "get-personal-access-token-metadata", "", false, "Get Personal Access Token Metadata")
	_workmailCmd.Flags().BoolVarP(&_workmailListAccessControlRules, "list-access-control-rules", "", false, "List Access Control Rules")
	_workmailCmd.Flags().BoolVarP(&_workmailListAliases, "list-aliases", "", false, "List Aliases")
	_workmailCmd.Flags().BoolVarP(&_workmailListAvailabilityConfigurations, "list-availability-configurations", "", false, "List Availability Configurations")
	_workmailCmd.Flags().BoolVarP(&_workmailListGroupMembers, "list-group-members", "", false, "List Group Members")
	_workmailCmd.Flags().BoolVarP(&_workmailListGroups, "list-groups", "", false, "List Groups")
	_workmailCmd.Flags().BoolVarP(&_workmailListGroupsForEntity, "list-groups-for-entity", "", false, "List Groups For Entity")
	_workmailCmd.Flags().BoolVarP(&_workmailListImpersonationRoles, "list-impersonation-roles", "", false, "List Impersonation Roles")
	_workmailCmd.Flags().BoolVarP(&_workmailListMailDomains, "list-mail-domains", "", false, "List Mail Domains")
	_workmailCmd.Flags().BoolVarP(&_workmailListMailboxExportJobs, "list-mailbox-export-jobs", "", false, "List Mailbox Export Jobs")
	_workmailCmd.Flags().BoolVarP(&_workmailListMailboxPermissions, "list-mailbox-permissions", "", false, "List Mailbox Permissions")
	_workmailCmd.Flags().BoolVarP(&_workmailListMobileDeviceAccessOverrides, "list-mobile-device-access-overrides", "", false, "List Mobile Device Access Overrides")
	_workmailCmd.Flags().BoolVarP(&_workmailListMobileDeviceAccessRules, "list-mobile-device-access-rules", "", false, "List Mobile Device Access Rules")
	_workmailCmd.Flags().BoolVarP(&_workmailListOrganizations, "list-organizations", "", false, "List Organizations")
	_workmailCmd.Flags().BoolVarP(&_workmailListPersonalAccessTokens, "list-personal-access-tokens", "", false, "List Personal Access Tokens")
	_workmailCmd.Flags().BoolVarP(&_workmailListResourceDelegates, "list-resource-delegates", "", false, "List Resource Delegates")
	_workmailCmd.Flags().BoolVarP(&_workmailListResources, "list-resources", "", false, "List Resources")
	_workmailCmd.Flags().BoolVarP(&_workmailListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_workmailCmd.Flags().BoolVarP(&_workmailListUsers, "list-users", "", false, "List Users")
	_workmailCmd.Flags().BoolVarP(&_workmailPutAccessControlRule, "put-access-control-rule", "", false, "Put Access Control Rule")
	_workmailCmd.Flags().BoolVarP(&_workmailPutEmailMonitoringConfiguration, "put-email-monitoring-configuration", "", false, "Put Email Monitoring Configuration")
	_workmailCmd.Flags().BoolVarP(&_workmailPutIdentityProviderConfiguration, "put-identity-provider-configuration", "", false, "Put Identity Provider Configuration")
	_workmailCmd.Flags().BoolVarP(&_workmailPutInboundDmarcSettings, "put-inbound-dmarc-settings", "", false, "Put Inbound Dmarc Settings")
	_workmailCmd.Flags().BoolVarP(&_workmailPutMailboxPermissions, "put-mailbox-permissions", "", false, "Put Mailbox Permissions")
	_workmailCmd.Flags().BoolVarP(&_workmailPutMobileDeviceAccessOverride, "put-mobile-device-access-override", "", false, "Put Mobile Device Access Override")
	_workmailCmd.Flags().BoolVarP(&_workmailPutRetentionPolicy, "put-retention-policy", "", false, "Put Retention Policy")
	_workmailCmd.Flags().BoolVarP(&_workmailRegisterMailDomain, "register-mail-domain", "", false, "Register Mail Domain")
	_workmailCmd.Flags().BoolVarP(&_workmailRegisterToWorkMail, "register-to-work-mail", "", false, "Register To Work Mail")
	_workmailCmd.Flags().BoolVarP(&_workmailResetPassword, "reset-password", "", false, "Reset Password")
	_workmailCmd.Flags().BoolVarP(&_workmailStartMailboxExportJob, "start-mailbox-export-job", "", false, "Start Mailbox Export Job")
	_workmailCmd.Flags().BoolVarP(&_workmailTagResource, "tag-resource", "", false, "Tag Resource")
	_workmailCmd.Flags().BoolVarP(&_workmailTestAvailabilityConfiguration, "test-availability-configuration", "", false, "Test Availability Configuration")
	_workmailCmd.Flags().BoolVarP(&_workmailUntagResource, "untag-resource", "", false, "Untag Resource")
	_workmailCmd.Flags().BoolVarP(&_workmailUpdateAvailabilityConfiguration, "update-availability-configuration", "", false, "Update Availability Configuration")
	_workmailCmd.Flags().BoolVarP(&_workmailUpdateDefaultMailDomain, "update-default-mail-domain", "", false, "Update Default Mail Domain")
	_workmailCmd.Flags().BoolVarP(&_workmailUpdateGroup, "update-group", "", false, "Update Group")
	_workmailCmd.Flags().BoolVarP(&_workmailUpdateImpersonationRole, "update-impersonation-role", "", false, "Update Impersonation Role")
	_workmailCmd.Flags().BoolVarP(&_workmailUpdateMailboxQuota, "update-mailbox-quota", "", false, "Update Mailbox Quota")
	_workmailCmd.Flags().BoolVarP(&_workmailUpdateMobileDeviceAccessRule, "update-mobile-device-access-rule", "", false, "Update Mobile Device Access Rule")
	_workmailCmd.Flags().BoolVarP(&_workmailUpdatePrimaryEmailAddress, "update-primary-email-address", "", false, "Update Primary Email Address")
	_workmailCmd.Flags().BoolVarP(&_workmailUpdateResource, "update-resource", "", false, "Update Resource")
	_workmailCmd.Flags().BoolVarP(&_workmailUpdateUser, "update-user", "", false, "Update User")

}
