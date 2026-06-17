package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// cognitoidentityproviderCmd represents the cognitoidentityprovider command
var _cognitoidentityproviderCmd = &cobra.Command{
	Use:   "cognitoidentityprovider",
	Short: "AWS cognitoidentityprovider CLI",
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
		client := cognitoidentityprovider.NewFromConfig(cfg)
		if _cognitoidentityproviderAddCustomAttributes {
			cognitoidentityprovider_AddCustomAttributes(cfg, client)
			return
		}
		if _cognitoidentityproviderAddUserPoolClientSecret {
			cognitoidentityprovider_AddUserPoolClientSecret(cfg, client)
			return
		}
		if _cognitoidentityproviderAdminAddUserToGroup {
			cognitoidentityprovider_AdminAddUserToGroup(cfg, client)
			return
		}
		if _cognitoidentityproviderAdminConfirmSignUp {
			cognitoidentityprovider_AdminConfirmSignUp(cfg, client)
			return
		}
		if _cognitoidentityproviderAdminCreateUser {
			cognitoidentityprovider_AdminCreateUser(cfg, client)
			return
		}
		if _cognitoidentityproviderAdminDeleteUser {
			cognitoidentityprovider_AdminDeleteUser(cfg, client)
			return
		}
		if _cognitoidentityproviderAdminDeleteUserAttributes {
			cognitoidentityprovider_AdminDeleteUserAttributes(cfg, client)
			return
		}
		if _cognitoidentityproviderAdminDisableProviderForUser {
			cognitoidentityprovider_AdminDisableProviderForUser(cfg, client)
			return
		}
		if _cognitoidentityproviderAdminDisableUser {
			cognitoidentityprovider_AdminDisableUser(cfg, client)
			return
		}
		if _cognitoidentityproviderAdminEnableUser {
			cognitoidentityprovider_AdminEnableUser(cfg, client)
			return
		}
		if _cognitoidentityproviderAdminForgetDevice {
			cognitoidentityprovider_AdminForgetDevice(cfg, client)
			return
		}
		if _cognitoidentityproviderAdminGetDevice {
			cognitoidentityprovider_AdminGetDevice(cfg, client)
			return
		}
		if _cognitoidentityproviderAdminGetUser {
			cognitoidentityprovider_AdminGetUser(cfg, client)
			return
		}
		if _cognitoidentityproviderAdminInitiateAuth {
			cognitoidentityprovider_AdminInitiateAuth(cfg, client)
			return
		}
		if _cognitoidentityproviderAdminLinkProviderForUser {
			cognitoidentityprovider_AdminLinkProviderForUser(cfg, client)
			return
		}
		if _cognitoidentityproviderAdminListDevices {
			cognitoidentityprovider_AdminListDevices(cfg, client)
			return
		}
		if _cognitoidentityproviderAdminListGroupsForUser {
			cognitoidentityprovider_AdminListGroupsForUser(cfg, client)
			return
		}
		if _cognitoidentityproviderAdminListUserAuthEvents {
			cognitoidentityprovider_AdminListUserAuthEvents(cfg, client)
			return
		}
		if _cognitoidentityproviderAdminRemoveUserFromGroup {
			cognitoidentityprovider_AdminRemoveUserFromGroup(cfg, client)
			return
		}
		if _cognitoidentityproviderAdminResetUserPassword {
			cognitoidentityprovider_AdminResetUserPassword(cfg, client)
			return
		}
		if _cognitoidentityproviderAdminRespondToAuthChallenge {
			cognitoidentityprovider_AdminRespondToAuthChallenge(cfg, client)
			return
		}
		if _cognitoidentityproviderAdminSetUserMFAPreference {
			cognitoidentityprovider_AdminSetUserMFAPreference(cfg, client)
			return
		}
		if _cognitoidentityproviderAdminSetUserPassword {
			cognitoidentityprovider_AdminSetUserPassword(cfg, client)
			return
		}
		if _cognitoidentityproviderAdminSetUserSettings {
			cognitoidentityprovider_AdminSetUserSettings(cfg, client)
			return
		}
		if _cognitoidentityproviderAdminUpdateAuthEventFeedback {
			cognitoidentityprovider_AdminUpdateAuthEventFeedback(cfg, client)
			return
		}
		if _cognitoidentityproviderAdminUpdateDeviceStatus {
			cognitoidentityprovider_AdminUpdateDeviceStatus(cfg, client)
			return
		}
		if _cognitoidentityproviderAdminUpdateUserAttributes {
			cognitoidentityprovider_AdminUpdateUserAttributes(cfg, client)
			return
		}
		if _cognitoidentityproviderAdminUserGlobalSignOut {
			cognitoidentityprovider_AdminUserGlobalSignOut(cfg, client)
			return
		}
		if _cognitoidentityproviderAssociateSoftwareToken {
			cognitoidentityprovider_AssociateSoftwareToken(cfg, client)
			return
		}
		if _cognitoidentityproviderChangePassword {
			cognitoidentityprovider_ChangePassword(cfg, client)
			return
		}
		if _cognitoidentityproviderCompleteWebAuthnRegistration {
			cognitoidentityprovider_CompleteWebAuthnRegistration(cfg, client)
			return
		}
		if _cognitoidentityproviderConfirmDevice {
			cognitoidentityprovider_ConfirmDevice(cfg, client)
			return
		}
		if _cognitoidentityproviderConfirmForgotPassword {
			cognitoidentityprovider_ConfirmForgotPassword(cfg, client)
			return
		}
		if _cognitoidentityproviderConfirmSignUp {
			cognitoidentityprovider_ConfirmSignUp(cfg, client)
			return
		}
		if _cognitoidentityproviderCreateGroup {
			cognitoidentityprovider_CreateGroup(cfg, client)
			return
		}
		if _cognitoidentityproviderCreateIdentityProvider {
			cognitoidentityprovider_CreateIdentityProvider(cfg, client)
			return
		}
		if _cognitoidentityproviderCreateManagedLoginBranding {
			cognitoidentityprovider_CreateManagedLoginBranding(cfg, client)
			return
		}
		if _cognitoidentityproviderCreateResourceServer {
			cognitoidentityprovider_CreateResourceServer(cfg, client)
			return
		}
		if _cognitoidentityproviderCreateTerms {
			cognitoidentityprovider_CreateTerms(cfg, client)
			return
		}
		if _cognitoidentityproviderCreateUserImportJob {
			cognitoidentityprovider_CreateUserImportJob(cfg, client)
			return
		}
		if _cognitoidentityproviderCreateUserPool {
			cognitoidentityprovider_CreateUserPool(cfg, client)
			return
		}
		if _cognitoidentityproviderCreateUserPoolClient {
			cognitoidentityprovider_CreateUserPoolClient(cfg, client)
			return
		}
		if _cognitoidentityproviderCreateUserPoolDomain {
			cognitoidentityprovider_CreateUserPoolDomain(cfg, client)
			return
		}
		if _cognitoidentityproviderDeleteGroup {
			cognitoidentityprovider_DeleteGroup(cfg, client)
			return
		}
		if _cognitoidentityproviderDeleteIdentityProvider {
			cognitoidentityprovider_DeleteIdentityProvider(cfg, client)
			return
		}
		if _cognitoidentityproviderDeleteManagedLoginBranding {
			cognitoidentityprovider_DeleteManagedLoginBranding(cfg, client)
			return
		}
		if _cognitoidentityproviderDeleteResourceServer {
			cognitoidentityprovider_DeleteResourceServer(cfg, client)
			return
		}
		if _cognitoidentityproviderDeleteTerms {
			cognitoidentityprovider_DeleteTerms(cfg, client)
			return
		}
		if _cognitoidentityproviderDeleteUser {
			cognitoidentityprovider_DeleteUser(cfg, client)
			return
		}
		if _cognitoidentityproviderDeleteUserAttributes {
			cognitoidentityprovider_DeleteUserAttributes(cfg, client)
			return
		}
		if _cognitoidentityproviderDeleteUserPool {
			cognitoidentityprovider_DeleteUserPool(cfg, client)
			return
		}
		if _cognitoidentityproviderDeleteUserPoolClient {
			cognitoidentityprovider_DeleteUserPoolClient(cfg, client)
			return
		}
		if _cognitoidentityproviderDeleteUserPoolClientSecret {
			cognitoidentityprovider_DeleteUserPoolClientSecret(cfg, client)
			return
		}
		if _cognitoidentityproviderDeleteUserPoolDomain {
			cognitoidentityprovider_DeleteUserPoolDomain(cfg, client)
			return
		}
		if _cognitoidentityproviderDeleteWebAuthnCredential {
			cognitoidentityprovider_DeleteWebAuthnCredential(cfg, client)
			return
		}
		if _cognitoidentityproviderDescribeIdentityProvider {
			cognitoidentityprovider_DescribeIdentityProvider(cfg, client)
			return
		}
		if _cognitoidentityproviderDescribeManagedLoginBranding {
			cognitoidentityprovider_DescribeManagedLoginBranding(cfg, client)
			return
		}
		if _cognitoidentityproviderDescribeManagedLoginBrandingByClient {
			cognitoidentityprovider_DescribeManagedLoginBrandingByClient(cfg, client)
			return
		}
		if _cognitoidentityproviderDescribeResourceServer {
			cognitoidentityprovider_DescribeResourceServer(cfg, client)
			return
		}
		if _cognitoidentityproviderDescribeRiskConfiguration {
			cognitoidentityprovider_DescribeRiskConfiguration(cfg, client)
			return
		}
		if _cognitoidentityproviderDescribeTerms {
			cognitoidentityprovider_DescribeTerms(cfg, client)
			return
		}
		if _cognitoidentityproviderDescribeUserImportJob {
			cognitoidentityprovider_DescribeUserImportJob(cfg, client)
			return
		}
		if _cognitoidentityproviderDescribeUserPool {
			cognitoidentityprovider_DescribeUserPool(cfg, client)
			return
		}
		if _cognitoidentityproviderDescribeUserPoolClient {
			cognitoidentityprovider_DescribeUserPoolClient(cfg, client)
			return
		}
		if _cognitoidentityproviderDescribeUserPoolDomain {
			cognitoidentityprovider_DescribeUserPoolDomain(cfg, client)
			return
		}
		if _cognitoidentityproviderForgetDevice {
			cognitoidentityprovider_ForgetDevice(cfg, client)
			return
		}
		if _cognitoidentityproviderForgotPassword {
			cognitoidentityprovider_ForgotPassword(cfg, client)
			return
		}
		if _cognitoidentityproviderGetCSVHeader {
			cognitoidentityprovider_GetCSVHeader(cfg, client)
			return
		}
		if _cognitoidentityproviderGetDevice {
			cognitoidentityprovider_GetDevice(cfg, client)
			return
		}
		if _cognitoidentityproviderGetGroup {
			cognitoidentityprovider_GetGroup(cfg, client)
			return
		}
		if _cognitoidentityproviderGetIdentityProviderByIdentifier {
			cognitoidentityprovider_GetIdentityProviderByIdentifier(cfg, client)
			return
		}
		if _cognitoidentityproviderGetLogDeliveryConfiguration {
			cognitoidentityprovider_GetLogDeliveryConfiguration(cfg, client)
			return
		}
		if _cognitoidentityproviderGetSigningCertificate {
			cognitoidentityprovider_GetSigningCertificate(cfg, client)
			return
		}
		if _cognitoidentityproviderGetTokensFromRefreshToken {
			cognitoidentityprovider_GetTokensFromRefreshToken(cfg, client)
			return
		}
		if _cognitoidentityproviderGetUICustomization {
			cognitoidentityprovider_GetUICustomization(cfg, client)
			return
		}
		if _cognitoidentityproviderGetUser {
			cognitoidentityprovider_GetUser(cfg, client)
			return
		}
		if _cognitoidentityproviderGetUserAttributeVerificationCode {
			cognitoidentityprovider_GetUserAttributeVerificationCode(cfg, client)
			return
		}
		if _cognitoidentityproviderGetUserAuthFactors {
			cognitoidentityprovider_GetUserAuthFactors(cfg, client)
			return
		}
		if _cognitoidentityproviderGetUserPoolMfaConfig {
			cognitoidentityprovider_GetUserPoolMfaConfig(cfg, client)
			return
		}
		if _cognitoidentityproviderGlobalSignOut {
			cognitoidentityprovider_GlobalSignOut(cfg, client)
			return
		}
		if _cognitoidentityproviderInitiateAuth {
			cognitoidentityprovider_InitiateAuth(cfg, client)
			return
		}
		if _cognitoidentityproviderListDevices {
			cognitoidentityprovider_ListDevices(cfg, client)
			return
		}
		if _cognitoidentityproviderListGroups {
			cognitoidentityprovider_ListGroups(cfg, client)
			return
		}
		if _cognitoidentityproviderListIdentityProviders {
			cognitoidentityprovider_ListIdentityProviders(cfg, client)
			return
		}
		if _cognitoidentityproviderListResourceServers {
			cognitoidentityprovider_ListResourceServers(cfg, client)
			return
		}
		if _cognitoidentityproviderListTagsForResource {
			cognitoidentityprovider_ListTagsForResource(cfg, client)
			return
		}
		if _cognitoidentityproviderListTerms {
			cognitoidentityprovider_ListTerms(cfg, client)
			return
		}
		if _cognitoidentityproviderListUserImportJobs {
			cognitoidentityprovider_ListUserImportJobs(cfg, client)
			return
		}
		if _cognitoidentityproviderListUserPoolClientSecrets {
			cognitoidentityprovider_ListUserPoolClientSecrets(cfg, client)
			return
		}
		if _cognitoidentityproviderListUserPoolClients {
			cognitoidentityprovider_ListUserPoolClients(cfg, client)
			return
		}
		if _cognitoidentityproviderListUserPools {
			cognitoidentityprovider_ListUserPools(cfg, client)
			return
		}
		if _cognitoidentityproviderListUsers {
			cognitoidentityprovider_ListUsers(cfg, client)
			return
		}
		if _cognitoidentityproviderListUsersInGroup {
			cognitoidentityprovider_ListUsersInGroup(cfg, client)
			return
		}
		if _cognitoidentityproviderListWebAuthnCredentials {
			cognitoidentityprovider_ListWebAuthnCredentials(cfg, client)
			return
		}
		if _cognitoidentityproviderResendConfirmationCode {
			cognitoidentityprovider_ResendConfirmationCode(cfg, client)
			return
		}
		if _cognitoidentityproviderRespondToAuthChallenge {
			cognitoidentityprovider_RespondToAuthChallenge(cfg, client)
			return
		}
		if _cognitoidentityproviderRevokeToken {
			cognitoidentityprovider_RevokeToken(cfg, client)
			return
		}
		if _cognitoidentityproviderSetLogDeliveryConfiguration {
			cognitoidentityprovider_SetLogDeliveryConfiguration(cfg, client)
			return
		}
		if _cognitoidentityproviderSetRiskConfiguration {
			cognitoidentityprovider_SetRiskConfiguration(cfg, client)
			return
		}
		if _cognitoidentityproviderSetUICustomization {
			cognitoidentityprovider_SetUICustomization(cfg, client)
			return
		}
		if _cognitoidentityproviderSetUserMFAPreference {
			cognitoidentityprovider_SetUserMFAPreference(cfg, client)
			return
		}
		if _cognitoidentityproviderSetUserPoolMfaConfig {
			cognitoidentityprovider_SetUserPoolMfaConfig(cfg, client)
			return
		}
		if _cognitoidentityproviderSetUserSettings {
			cognitoidentityprovider_SetUserSettings(cfg, client)
			return
		}
		if _cognitoidentityproviderSignUp {
			cognitoidentityprovider_SignUp(cfg, client)
			return
		}
		if _cognitoidentityproviderStartUserImportJob {
			cognitoidentityprovider_StartUserImportJob(cfg, client)
			return
		}
		if _cognitoidentityproviderStartWebAuthnRegistration {
			cognitoidentityprovider_StartWebAuthnRegistration(cfg, client)
			return
		}
		if _cognitoidentityproviderStopUserImportJob {
			cognitoidentityprovider_StopUserImportJob(cfg, client)
			return
		}
		if _cognitoidentityproviderTagResource {
			cognitoidentityprovider_TagResource(cfg, client)
			return
		}
		if _cognitoidentityproviderUntagResource {
			cognitoidentityprovider_UntagResource(cfg, client)
			return
		}
		if _cognitoidentityproviderUpdateAuthEventFeedback {
			cognitoidentityprovider_UpdateAuthEventFeedback(cfg, client)
			return
		}
		if _cognitoidentityproviderUpdateDeviceStatus {
			cognitoidentityprovider_UpdateDeviceStatus(cfg, client)
			return
		}
		if _cognitoidentityproviderUpdateGroup {
			cognitoidentityprovider_UpdateGroup(cfg, client)
			return
		}
		if _cognitoidentityproviderUpdateIdentityProvider {
			cognitoidentityprovider_UpdateIdentityProvider(cfg, client)
			return
		}
		if _cognitoidentityproviderUpdateManagedLoginBranding {
			cognitoidentityprovider_UpdateManagedLoginBranding(cfg, client)
			return
		}
		if _cognitoidentityproviderUpdateResourceServer {
			cognitoidentityprovider_UpdateResourceServer(cfg, client)
			return
		}
		if _cognitoidentityproviderUpdateTerms {
			cognitoidentityprovider_UpdateTerms(cfg, client)
			return
		}
		if _cognitoidentityproviderUpdateUserAttributes {
			cognitoidentityprovider_UpdateUserAttributes(cfg, client)
			return
		}
		if _cognitoidentityproviderUpdateUserPool {
			cognitoidentityprovider_UpdateUserPool(cfg, client)
			return
		}
		if _cognitoidentityproviderUpdateUserPoolClient {
			cognitoidentityprovider_UpdateUserPoolClient(cfg, client)
			return
		}
		if _cognitoidentityproviderUpdateUserPoolDomain {
			cognitoidentityprovider_UpdateUserPoolDomain(cfg, client)
			return
		}
		if _cognitoidentityproviderVerifySoftwareToken {
			cognitoidentityprovider_VerifySoftwareToken(cfg, client)
			return
		}
		if _cognitoidentityproviderVerifyUserAttribute {
			cognitoidentityprovider_VerifyUserAttribute(cfg, client)
			return
		}

	},
}

var (
	_cognitoidentityproviderAddCustomAttributes                  bool
	_cognitoidentityproviderAddUserPoolClientSecret              bool
	_cognitoidentityproviderAdminAddUserToGroup                  bool
	_cognitoidentityproviderAdminConfirmSignUp                   bool
	_cognitoidentityproviderAdminCreateUser                      bool
	_cognitoidentityproviderAdminDeleteUser                      bool
	_cognitoidentityproviderAdminDeleteUserAttributes            bool
	_cognitoidentityproviderAdminDisableProviderForUser          bool
	_cognitoidentityproviderAdminDisableUser                     bool
	_cognitoidentityproviderAdminEnableUser                      bool
	_cognitoidentityproviderAdminForgetDevice                    bool
	_cognitoidentityproviderAdminGetDevice                       bool
	_cognitoidentityproviderAdminGetUser                         bool
	_cognitoidentityproviderAdminInitiateAuth                    bool
	_cognitoidentityproviderAdminLinkProviderForUser             bool
	_cognitoidentityproviderAdminListDevices                     bool
	_cognitoidentityproviderAdminListGroupsForUser               bool
	_cognitoidentityproviderAdminListUserAuthEvents              bool
	_cognitoidentityproviderAdminRemoveUserFromGroup             bool
	_cognitoidentityproviderAdminResetUserPassword               bool
	_cognitoidentityproviderAdminRespondToAuthChallenge          bool
	_cognitoidentityproviderAdminSetUserMFAPreference            bool
	_cognitoidentityproviderAdminSetUserPassword                 bool
	_cognitoidentityproviderAdminSetUserSettings                 bool
	_cognitoidentityproviderAdminUpdateAuthEventFeedback         bool
	_cognitoidentityproviderAdminUpdateDeviceStatus              bool
	_cognitoidentityproviderAdminUpdateUserAttributes            bool
	_cognitoidentityproviderAdminUserGlobalSignOut               bool
	_cognitoidentityproviderAssociateSoftwareToken               bool
	_cognitoidentityproviderChangePassword                       bool
	_cognitoidentityproviderCompleteWebAuthnRegistration         bool
	_cognitoidentityproviderConfirmDevice                        bool
	_cognitoidentityproviderConfirmForgotPassword                bool
	_cognitoidentityproviderConfirmSignUp                        bool
	_cognitoidentityproviderCreateGroup                          bool
	_cognitoidentityproviderCreateIdentityProvider               bool
	_cognitoidentityproviderCreateManagedLoginBranding           bool
	_cognitoidentityproviderCreateResourceServer                 bool
	_cognitoidentityproviderCreateTerms                          bool
	_cognitoidentityproviderCreateUserImportJob                  bool
	_cognitoidentityproviderCreateUserPool                       bool
	_cognitoidentityproviderCreateUserPoolClient                 bool
	_cognitoidentityproviderCreateUserPoolDomain                 bool
	_cognitoidentityproviderDeleteGroup                          bool
	_cognitoidentityproviderDeleteIdentityProvider               bool
	_cognitoidentityproviderDeleteManagedLoginBranding           bool
	_cognitoidentityproviderDeleteResourceServer                 bool
	_cognitoidentityproviderDeleteTerms                          bool
	_cognitoidentityproviderDeleteUser                           bool
	_cognitoidentityproviderDeleteUserAttributes                 bool
	_cognitoidentityproviderDeleteUserPool                       bool
	_cognitoidentityproviderDeleteUserPoolClient                 bool
	_cognitoidentityproviderDeleteUserPoolClientSecret           bool
	_cognitoidentityproviderDeleteUserPoolDomain                 bool
	_cognitoidentityproviderDeleteWebAuthnCredential             bool
	_cognitoidentityproviderDescribeIdentityProvider             bool
	_cognitoidentityproviderDescribeManagedLoginBranding         bool
	_cognitoidentityproviderDescribeManagedLoginBrandingByClient bool
	_cognitoidentityproviderDescribeResourceServer               bool
	_cognitoidentityproviderDescribeRiskConfiguration            bool
	_cognitoidentityproviderDescribeTerms                        bool
	_cognitoidentityproviderDescribeUserImportJob                bool
	_cognitoidentityproviderDescribeUserPool                     bool
	_cognitoidentityproviderDescribeUserPoolClient               bool
	_cognitoidentityproviderDescribeUserPoolDomain               bool
	_cognitoidentityproviderForgetDevice                         bool
	_cognitoidentityproviderForgotPassword                       bool
	_cognitoidentityproviderGetCSVHeader                         bool
	_cognitoidentityproviderGetDevice                            bool
	_cognitoidentityproviderGetGroup                             bool
	_cognitoidentityproviderGetIdentityProviderByIdentifier      bool
	_cognitoidentityproviderGetLogDeliveryConfiguration          bool
	_cognitoidentityproviderGetSigningCertificate                bool
	_cognitoidentityproviderGetTokensFromRefreshToken            bool
	_cognitoidentityproviderGetUICustomization                   bool
	_cognitoidentityproviderGetUser                              bool
	_cognitoidentityproviderGetUserAttributeVerificationCode     bool
	_cognitoidentityproviderGetUserAuthFactors                   bool
	_cognitoidentityproviderGetUserPoolMfaConfig                 bool
	_cognitoidentityproviderGlobalSignOut                        bool
	_cognitoidentityproviderInitiateAuth                         bool
	_cognitoidentityproviderListDevices                          bool
	_cognitoidentityproviderListGroups                           bool
	_cognitoidentityproviderListIdentityProviders                bool
	_cognitoidentityproviderListResourceServers                  bool
	_cognitoidentityproviderListTagsForResource                  bool
	_cognitoidentityproviderListTerms                            bool
	_cognitoidentityproviderListUserImportJobs                   bool
	_cognitoidentityproviderListUserPoolClientSecrets            bool
	_cognitoidentityproviderListUserPoolClients                  bool
	_cognitoidentityproviderListUserPools                        bool
	_cognitoidentityproviderListUsers                            bool
	_cognitoidentityproviderListUsersInGroup                     bool
	_cognitoidentityproviderListWebAuthnCredentials              bool
	_cognitoidentityproviderResendConfirmationCode               bool
	_cognitoidentityproviderRespondToAuthChallenge               bool
	_cognitoidentityproviderRevokeToken                          bool
	_cognitoidentityproviderSetLogDeliveryConfiguration          bool
	_cognitoidentityproviderSetRiskConfiguration                 bool
	_cognitoidentityproviderSetUICustomization                   bool
	_cognitoidentityproviderSetUserMFAPreference                 bool
	_cognitoidentityproviderSetUserPoolMfaConfig                 bool
	_cognitoidentityproviderSetUserSettings                      bool
	_cognitoidentityproviderSignUp                               bool
	_cognitoidentityproviderStartUserImportJob                   bool
	_cognitoidentityproviderStartWebAuthnRegistration            bool
	_cognitoidentityproviderStopUserImportJob                    bool
	_cognitoidentityproviderTagResource                          bool
	_cognitoidentityproviderUntagResource                        bool
	_cognitoidentityproviderUpdateAuthEventFeedback              bool
	_cognitoidentityproviderUpdateDeviceStatus                   bool
	_cognitoidentityproviderUpdateGroup                          bool
	_cognitoidentityproviderUpdateIdentityProvider               bool
	_cognitoidentityproviderUpdateManagedLoginBranding           bool
	_cognitoidentityproviderUpdateResourceServer                 bool
	_cognitoidentityproviderUpdateTerms                          bool
	_cognitoidentityproviderUpdateUserAttributes                 bool
	_cognitoidentityproviderUpdateUserPool                       bool
	_cognitoidentityproviderUpdateUserPoolClient                 bool
	_cognitoidentityproviderUpdateUserPoolDomain                 bool
	_cognitoidentityproviderVerifySoftwareToken                  bool
	_cognitoidentityproviderVerifyUserAttribute                  bool

	_cognitoidentityproviderAccessToken                              string
	_cognitoidentityproviderAccessTokenValidity                      string
	_cognitoidentityproviderAccountRecoverySetting                   string
	_cognitoidentityproviderAccountTakeoverRiskConfiguration         string
	_cognitoidentityproviderAdminCreateUserConfig                    string
	_cognitoidentityproviderAliasAttributes                          string
	_cognitoidentityproviderAllowedOAuthFlows                        string
	_cognitoidentityproviderAllowedOAuthFlowsUserPoolClient          string
	_cognitoidentityproviderAllowedOAuthScopes                       []string
	_cognitoidentityproviderAnalyticsConfiguration                   string
	_cognitoidentityproviderAnalyticsMetadata                        string
	_cognitoidentityproviderAssets                                   string
	_cognitoidentityproviderAttributeMapping                         string
	_cognitoidentityproviderAttributeName                            string
	_cognitoidentityproviderAttributesToGet                          []string
	_cognitoidentityproviderAuthFlow                                 string
	_cognitoidentityproviderAuthParameters                           string
	_cognitoidentityproviderAuthSessionValidity                      string
	_cognitoidentityproviderAutoVerifiedAttributes                   string
	_cognitoidentityproviderCallbackURLs                             []string
	_cognitoidentityproviderChallengeName                            string
	_cognitoidentityproviderChallengeResponses                       string
	_cognitoidentityproviderClientId                                 string
	_cognitoidentityproviderClientMetadata                           string
	_cognitoidentityproviderClientName                               string
	_cognitoidentityproviderClientSecret                             string
	_cognitoidentityproviderClientSecretId                           string
	_cognitoidentityproviderCloudWatchLogsRoleArn                    string
	_cognitoidentityproviderCode                                     string
	_cognitoidentityproviderCompromisedCredentialsRiskConfiguration  string
	_cognitoidentityproviderConfirmationCode                         string
	_cognitoidentityproviderContextData                              string
	_cognitoidentityproviderCredential                               string
	_cognitoidentityproviderCredentialId                             string
	_cognitoidentityproviderCSS                                      string
	_cognitoidentityproviderCustomAttributes                         string
	_cognitoidentityproviderCustomDomainConfig                       string
	_cognitoidentityproviderDefaultRedirectURI                       string
	_cognitoidentityproviderDeletionProtection                       string
	_cognitoidentityproviderDescription                              string
	_cognitoidentityproviderDesiredDeliveryMediums                   string
	_cognitoidentityproviderDestinationUser                          string
	_cognitoidentityproviderDeviceConfiguration                      string
	_cognitoidentityproviderDeviceKey                                string
	_cognitoidentityproviderDeviceName                               string
	_cognitoidentityproviderDeviceRememberedStatus                   string
	_cognitoidentityproviderDeviceSecretVerifierConfig               string
	_cognitoidentityproviderDomain                                   string
	_cognitoidentityproviderEmailConfiguration                       string
	_cognitoidentityproviderEmailMfaConfiguration                    string
	_cognitoidentityproviderEmailMfaSettings                         string
	_cognitoidentityproviderEmailVerificationMessage                 string
	_cognitoidentityproviderEmailVerificationSubject                 string
	_cognitoidentityproviderEnablePropagateAdditionalUserContextData string
	_cognitoidentityproviderEnableTokenRevocation                    string
	_cognitoidentityproviderEnforcement                              string
	_cognitoidentityproviderEventId                                  string
	_cognitoidentityproviderExplicitAuthFlows                        string
	_cognitoidentityproviderFeedbackToken                            string
	_cognitoidentityproviderFeedbackValue                            string
	_cognitoidentityproviderFilter                                   string
	_cognitoidentityproviderForceAliasCreation                       string
	_cognitoidentityproviderFriendlyDeviceName                       string
	_cognitoidentityproviderGenerateSecret                           string
	_cognitoidentityproviderGroupName                                string
	_cognitoidentityproviderIdTokenValidity                          string
	_cognitoidentityproviderIdentifier                               string
	_cognitoidentityproviderIdpIdentifier                            string
	_cognitoidentityproviderIdpIdentifiers                           []string
	_cognitoidentityproviderImageFile                                string
	_cognitoidentityproviderJobId                                    string
	_cognitoidentityproviderJobName                                  string
	_cognitoidentityproviderLambdaConfig                             string
	_cognitoidentityproviderLimit                                    string
	_cognitoidentityproviderLinks                                    string
	_cognitoidentityproviderLogConfigurations                        string
	_cognitoidentityproviderLogoutURLs                               []string
	_cognitoidentityproviderManagedLoginBrandingId                   string
	_cognitoidentityproviderManagedLoginVersion                      string
	_cognitoidentityproviderMaxResults                               string
	_cognitoidentityproviderMessageAction                            string
	_cognitoidentityproviderMfaConfiguration                         string
	_cognitoidentityproviderMFAOptions                               string
	_cognitoidentityproviderName                                     string
	_cognitoidentityproviderNextToken                                string
	_cognitoidentityproviderPaginationToken                          string
	_cognitoidentityproviderPassword                                 string
	_cognitoidentityproviderPermanent                                string
	_cognitoidentityproviderPolicies                                 string
	_cognitoidentityproviderPoolName                                 string
	_cognitoidentityproviderPrecedence                               string
	_cognitoidentityproviderPreventUserExistenceErrors               string
	_cognitoidentityproviderPreviousPassword                         string
	_cognitoidentityproviderProposedPassword                         string
	_cognitoidentityproviderProviderDetails                          string
	_cognitoidentityproviderProviderName                             string
	_cognitoidentityproviderProviderType                             string
	_cognitoidentityproviderReadAttributes                           []string
	_cognitoidentityproviderRefreshToken                             string
	_cognitoidentityproviderRefreshTokenRotation                     string
	_cognitoidentityproviderRefreshTokenValidity                     string
	_cognitoidentityproviderResourceArn                              string
	_cognitoidentityproviderReturnMergedResources                    string
	_cognitoidentityproviderRiskExceptionConfiguration               string
	_cognitoidentityproviderRoleArn                                  string
	_cognitoidentityproviderSchema                                   string
	_cognitoidentityproviderScopes                                   string
	_cognitoidentityproviderSecretHash                               string
	_cognitoidentityproviderSession                                  string
	_cognitoidentityproviderSettings                                 string
	_cognitoidentityproviderSmsAuthenticationMessage                 string
	_cognitoidentityproviderSmsConfiguration                         string
	_cognitoidentityproviderSmsMfaConfiguration                      string
	_cognitoidentityproviderSMSMfaSettings                           string
	_cognitoidentityproviderSmsVerificationMessage                   string
	_cognitoidentityproviderSoftwareTokenMfaConfiguration            string
	_cognitoidentityproviderSoftwareTokenMfaSettings                 string
	_cognitoidentityproviderSourceUser                               string
	_cognitoidentityproviderSupportedIdentityProviders               []string
	_cognitoidentityproviderTagKeys                                  []string
	_cognitoidentityproviderTags                                     string
	_cognitoidentityproviderTemporaryPassword                        string
	_cognitoidentityproviderTermsId                                  string
	_cognitoidentityproviderTermsName                                string
	_cognitoidentityproviderTermsSource                              string
	_cognitoidentityproviderToken                                    string
	_cognitoidentityproviderTokenValidityUnits                       string
	_cognitoidentityproviderUseCognitoProvidedValues                 string
	_cognitoidentityproviderUser                                     string
	_cognitoidentityproviderUserAttributeNames                       []string
	_cognitoidentityproviderUserAttributeUpdateSettings              string
	_cognitoidentityproviderUserAttributes                           string
	_cognitoidentityproviderUserCode                                 string
	_cognitoidentityproviderUserContextData                          string
	_cognitoidentityproviderUserPoolAddOns                           string
	_cognitoidentityproviderUserPoolId                               string
	_cognitoidentityproviderUserPoolTags                             string
	_cognitoidentityproviderUserPoolTier                             string
	_cognitoidentityproviderUsername                                 string
	_cognitoidentityproviderUsernameAttributes                       string
	_cognitoidentityproviderUsernameConfiguration                    string
	_cognitoidentityproviderValidationData                           string
	_cognitoidentityproviderVerificationMessageTemplate              string
	_cognitoidentityproviderWebAuthnConfiguration                    string
	_cognitoidentityproviderWriteAttributes                          []string
)

// Adds additional user attributes to the user pool schema. Custom attributes can
// be mutable or immutable and have a custom: or dev: prefix. For more
// information, see [Custom attributes].
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [Custom attributes]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-attributes.html#user-pool-settings-custom-attributes
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func cognitoidentityprovider_AddCustomAttributes(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.AddCustomAttributesInput{
		// CustomAttributes: []types.SchemaAttributeType, // Required
		// UserPoolId: *string, // Required
	}

	if len(_cognitoidentityproviderCustomAttributes) > 0 {
		if err := assignInputField(input, "CustomAttributes", _cognitoidentityproviderCustomAttributes); err != nil {
			log.Errorf("invalid --custom-attributes: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}

	if resp, err := client.AddCustomAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new client secret for an existing confidential user pool app client.
// Supports up to 2 active secrets per app client for zero-downtime credential
// rotation workflows.
func cognitoidentityprovider_AddUserPoolClientSecret(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.AddUserPoolClientSecretInput{
		// ClientId: *string, // Required
		// UserPoolId: *string, // Required
	}

	if len(_cognitoidentityproviderClientId) > 0 {
		input.ClientId = aws.String(_cognitoidentityproviderClientId)
	}
	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}
	if len(_cognitoidentityproviderClientSecret) > 0 {
		input.ClientSecret = aws.String(_cognitoidentityproviderClientSecret)
	}

	if resp, err := client.AddUserPoolClientSecret(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds a user to a group. A user who is in a group can present a preferred-role
// claim to an identity pool, and populates a cognito:groups claim to their access
// and identity tokens.
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func cognitoidentityprovider_AdminAddUserToGroup(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.AdminAddUserToGroupInput{
		// GroupName: *string, // Required
		// UserPoolId: *string, // Required
		// Username: *string, // Required
	}

	if len(_cognitoidentityproviderGroupName) > 0 {
		input.GroupName = aws.String(_cognitoidentityproviderGroupName)
	}
	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}
	if len(_cognitoidentityproviderUsername) > 0 {
		input.Username = aws.String(_cognitoidentityproviderUsername)
	}

	if resp, err := client.AdminAddUserToGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Confirms user sign-up as an administrator.
// This request sets a user account active in a user pool that [requires confirmation of new user accounts] before they can
// sign in. You can configure your user pool to not send confirmation codes to new
// users and instead confirm them with this API operation on the back end.
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// To configure your user pool to require administrative confirmation of users,
// set AllowAdminCreateUserOnly to true in a CreateUserPool or UpdateUserPool
// request.
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [requires confirmation of new user accounts]: https://docs.aws.amazon.com/cognito/latest/developerguide/signing-up-users-in-your-app.html#signing-up-users-in-your-app-and-confirming-them-as-admin
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func cognitoidentityprovider_AdminConfirmSignUp(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.AdminConfirmSignUpInput{
		// UserPoolId: *string, // Required
		// Username: *string, // Required
	}

	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}
	if len(_cognitoidentityproviderUsername) > 0 {
		input.Username = aws.String(_cognitoidentityproviderUsername)
	}
	if len(_cognitoidentityproviderClientMetadata) > 0 {
		if err := assignInputField(input, "ClientMetadata", _cognitoidentityproviderClientMetadata); err != nil {
			log.Errorf("invalid --client-metadata: %s", err.Error())
			return
		}
	}

	if resp, err := client.AdminConfirmSignUp(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new user in the specified user pool.
// If MessageAction isn't set, the default is to send a welcome message via email
// or phone (SMS).
//
// This message is based on a template that you configured in your call to create
// or update a user pool. This template includes your custom sign-up instructions
// and placeholders for user name and temporary password.
//
// Alternatively, you can call AdminCreateUser with SUPPRESS for the MessageAction
// parameter, and Amazon Cognito won't send any email.
//
// In either case, if the user has a password, they will be in the
// FORCE_CHANGE_PASSWORD state until they sign in and set their password. Your
// invitation message template must have the {####} password placeholder if your
// users have passwords. If your template doesn't have this placeholder, Amazon
// Cognito doesn't deliver the invitation message. In this case, you must update
// your message template and resend the password with a new AdminCreateUser
// request with a MessageAction value of RESEND .
//
// This action might generate an SMS text message. Starting June 1, 2021, US
// telecom carriers require you to register an origination phone number before you
// can send SMS messages to US phone numbers. If you use SMS text messages in
// Amazon Cognito, you must register a phone number with [Amazon Pinpoint]. Amazon Cognito uses the
// registered number automatically. Otherwise, Amazon Cognito users who must
// receive SMS messages might not be able to sign up, activate their accounts, or
// sign in.
//
// If you have never used SMS text messages with Amazon Cognito or any other
// Amazon Web Services service, Amazon Simple Notification Service might place your
// account in the SMS sandbox. In [sandbox mode], you can send messages only to verified phone
// numbers. After you test your app while in the sandbox environment, you can move
// out of the sandbox and into production. For more information, see [SMS message settings for Amazon Cognito user pools]in the Amazon
// Cognito Developer Guide.
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [SMS message settings for Amazon Cognito user pools]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-sms-settings.html
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [sandbox mode]: https://docs.aws.amazon.com/sns/latest/dg/sns-sms-sandbox.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
// [Amazon Pinpoint]: https://console.aws.amazon.com/pinpoint/home/
func cognitoidentityprovider_AdminCreateUser(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.AdminCreateUserInput{
		// UserPoolId: *string, // Required
		// Username: *string, // Required
	}

	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}
	if len(_cognitoidentityproviderUsername) > 0 {
		input.Username = aws.String(_cognitoidentityproviderUsername)
	}
	if len(_cognitoidentityproviderClientMetadata) > 0 {
		if err := assignInputField(input, "ClientMetadata", _cognitoidentityproviderClientMetadata); err != nil {
			log.Errorf("invalid --client-metadata: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderDesiredDeliveryMediums) > 0 {
		if err := assignInputField(input, "DesiredDeliveryMediums", _cognitoidentityproviderDesiredDeliveryMediums); err != nil {
			log.Errorf("invalid --desired-delivery-mediums: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderForceAliasCreation) > 0 {
		if err := assignInputField(input, "ForceAliasCreation", _cognitoidentityproviderForceAliasCreation); err != nil {
			log.Errorf("invalid --force-alias-creation: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderMessageAction) > 0 {
		if err := assignInputField(input, "MessageAction", _cognitoidentityproviderMessageAction); err != nil {
			log.Errorf("invalid --message-action: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderTemporaryPassword) > 0 {
		input.TemporaryPassword = aws.String(_cognitoidentityproviderTemporaryPassword)
	}
	if len(_cognitoidentityproviderUserAttributes) > 0 {
		if err := assignInputField(input, "UserAttributes", _cognitoidentityproviderUserAttributes); err != nil {
			log.Errorf("invalid --user-attributes: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderValidationData) > 0 {
		if err := assignInputField(input, "ValidationData", _cognitoidentityproviderValidationData); err != nil {
			log.Errorf("invalid --validation-data: %s", err.Error())
			return
		}
	}

	if resp, err := client.AdminCreateUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a user profile in your user pool.
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func cognitoidentityprovider_AdminDeleteUser(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.AdminDeleteUserInput{
		// UserPoolId: *string, // Required
		// Username: *string, // Required
	}

	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}
	if len(_cognitoidentityproviderUsername) > 0 {
		input.Username = aws.String(_cognitoidentityproviderUsername)
	}

	if resp, err := client.AdminDeleteUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes attribute values from a user. This operation doesn't affect tokens for
// existing user sessions. The next ID token that the user receives will no longer
// have the deleted attributes.
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func cognitoidentityprovider_AdminDeleteUserAttributes(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.AdminDeleteUserAttributesInput{
		// UserAttributeNames: []string, // Required
		// UserPoolId: *string, // Required
		// Username: *string, // Required
	}

	if len(_cognitoidentityproviderUserAttributeNames) > 0 {
		input.UserAttributeNames = append([]string(nil), _cognitoidentityproviderUserAttributeNames...)
	}
	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}
	if len(_cognitoidentityproviderUsername) > 0 {
		input.Username = aws.String(_cognitoidentityproviderUsername)
	}

	if resp, err := client.AdminDeleteUserAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Prevents the user from signing in with the specified external (SAML or social)
// identity provider (IdP). If the user that you want to deactivate is a Amazon
// Cognito user pools native username + password user, they can't use their
// password to sign in. If the user to deactivate is a linked external IdP user,
// any link between that user and an existing user is removed. When the external
// user signs in again, and the user is no longer attached to the previously linked
// DestinationUser , the user must create a new user account.
//
// The value of ProviderName must match the name of a user pool IdP.
//
// To deactivate a local user, set ProviderName to Cognito and the
// ProviderAttributeName to Cognito_Subject . The ProviderAttributeValue must be
// user's local username.
//
// The ProviderAttributeName must always be Cognito_Subject for social IdPs. The
// ProviderAttributeValue must always be the exact subject that was used when the
// user was originally linked as a source user.
//
// For de-linking a SAML identity, there are two scenarios. If the linked identity
// has not yet been used to sign in, the ProviderAttributeName and
// ProviderAttributeValue must be the same values that were used for the SourceUser
// when the identities were originally linked using AdminLinkProviderForUser call.
// This is also true if the linking was done with ProviderAttributeName set to
// Cognito_Subject . If the user has already signed in, the ProviderAttributeName
// must be Cognito_Subject and ProviderAttributeValue must be the NameID from
// their SAML assertion.
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func cognitoidentityprovider_AdminDisableProviderForUser(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.AdminDisableProviderForUserInput{
		// User: *types.ProviderUserIdentifierType, // Required
		// UserPoolId: *string, // Required
	}

	if len(_cognitoidentityproviderUser) > 0 {
		if err := assignInputField(input, "User", _cognitoidentityproviderUser); err != nil {
			log.Errorf("invalid --user: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}

	if resp, err := client.AdminDisableProviderForUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deactivates a user profile and revokes all access tokens for the user. A
// deactivated user can't sign in, but still appears in the responses to ListUsers
// API requests.
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func cognitoidentityprovider_AdminDisableUser(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.AdminDisableUserInput{
		// UserPoolId: *string, // Required
		// Username: *string, // Required
	}

	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}
	if len(_cognitoidentityproviderUsername) > 0 {
		input.Username = aws.String(_cognitoidentityproviderUsername)
	}

	if resp, err := client.AdminDisableUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Activates sign-in for a user profile that previously had sign-in access
// disabled.
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func cognitoidentityprovider_AdminEnableUser(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.AdminEnableUserInput{
		// UserPoolId: *string, // Required
		// Username: *string, // Required
	}

	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}
	if len(_cognitoidentityproviderUsername) > 0 {
		input.Username = aws.String(_cognitoidentityproviderUsername)
	}

	if resp, err := client.AdminEnableUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Forgets, or deletes, a remembered device from a user's profile. After you
// forget the device, the user can no longer complete device authentication with
// that device and when applicable, must submit MFA codes again. For more
// information, see [Working with devices].
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [Working with devices]: https://docs.aws.amazon.com/cognito/latest/developerguide/amazon-cognito-user-pools-device-tracking.html
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func cognitoidentityprovider_AdminForgetDevice(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.AdminForgetDeviceInput{
		// DeviceKey: *string, // Required
		// UserPoolId: *string, // Required
		// Username: *string, // Required
	}

	if len(_cognitoidentityproviderDeviceKey) > 0 {
		input.DeviceKey = aws.String(_cognitoidentityproviderDeviceKey)
	}
	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}
	if len(_cognitoidentityproviderUsername) > 0 {
		input.Username = aws.String(_cognitoidentityproviderUsername)
	}

	if resp, err := client.AdminForgetDevice(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Given the device key, returns details for a user's device. For more
// information, see [Working with devices].
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [Working with devices]: https://docs.aws.amazon.com/cognito/latest/developerguide/amazon-cognito-user-pools-device-tracking.html
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func cognitoidentityprovider_AdminGetDevice(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.AdminGetDeviceInput{
		// DeviceKey: *string, // Required
		// UserPoolId: *string, // Required
		// Username: *string, // Required
	}

	if len(_cognitoidentityproviderDeviceKey) > 0 {
		input.DeviceKey = aws.String(_cognitoidentityproviderDeviceKey)
	}
	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}
	if len(_cognitoidentityproviderUsername) > 0 {
		input.Username = aws.String(_cognitoidentityproviderUsername)
	}

	if resp, err := client.AdminGetDevice(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Given a username, returns details about a user profile in a user pool. You can
// specify alias attributes in the Username request parameter.
//
// This operation contributes to your monthly active user (MAU) count for the
// purpose of billing.
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func cognitoidentityprovider_AdminGetUser(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.AdminGetUserInput{
		// UserPoolId: *string, // Required
		// Username: *string, // Required
	}

	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}
	if len(_cognitoidentityproviderUsername) > 0 {
		input.Username = aws.String(_cognitoidentityproviderUsername)
	}

	if resp, err := client.AdminGetUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts sign-in for applications with a server-side component, for example a
// traditional web application. This operation specifies the authentication flow
// that you'd like to begin. The authentication flow that you specify must be
// supported in your app client configuration. For more information about
// authentication flows, see [Authentication flows].
//
// This action might generate an SMS text message. Starting June 1, 2021, US
// telecom carriers require you to register an origination phone number before you
// can send SMS messages to US phone numbers. If you use SMS text messages in
// Amazon Cognito, you must register a phone number with [Amazon Pinpoint]. Amazon Cognito uses the
// registered number automatically. Otherwise, Amazon Cognito users who must
// receive SMS messages might not be able to sign up, activate their accounts, or
// sign in.
//
// If you have never used SMS text messages with Amazon Cognito or any other
// Amazon Web Services service, Amazon Simple Notification Service might place your
// account in the SMS sandbox. In [sandbox mode], you can send messages only to verified phone
// numbers. After you test your app while in the sandbox environment, you can move
// out of the sandbox and into production. For more information, see [SMS message settings for Amazon Cognito user pools]in the Amazon
// Cognito Developer Guide.
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [SMS message settings for Amazon Cognito user pools]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-sms-settings.html
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [sandbox mode]: https://docs.aws.amazon.com/sns/latest/dg/sns-sms-sandbox.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
// [Authentication flows]: https://docs.aws.amazon.com/cognito/latest/developerguide/amazon-cognito-user-pools-authentication-flow-methods.html
// [Amazon Pinpoint]: https://console.aws.amazon.com/pinpoint/home/
func cognitoidentityprovider_AdminInitiateAuth(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.AdminInitiateAuthInput{
		// AuthFlow: types.AuthFlowType, // Required
		// ClientId: *string, // Required
		// UserPoolId: *string, // Required
	}

	if len(_cognitoidentityproviderAuthFlow) > 0 {
		if err := assignInputField(input, "AuthFlow", _cognitoidentityproviderAuthFlow); err != nil {
			log.Errorf("invalid --auth-flow: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderClientId) > 0 {
		input.ClientId = aws.String(_cognitoidentityproviderClientId)
	}
	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}
	if len(_cognitoidentityproviderAnalyticsMetadata) > 0 {
		if err := assignInputField(input, "AnalyticsMetadata", _cognitoidentityproviderAnalyticsMetadata); err != nil {
			log.Errorf("invalid --analytics-metadata: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderAuthParameters) > 0 {
		if err := assignInputField(input, "AuthParameters", _cognitoidentityproviderAuthParameters); err != nil {
			log.Errorf("invalid --auth-parameters: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderClientMetadata) > 0 {
		if err := assignInputField(input, "ClientMetadata", _cognitoidentityproviderClientMetadata); err != nil {
			log.Errorf("invalid --client-metadata: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderContextData) > 0 {
		if err := assignInputField(input, "ContextData", _cognitoidentityproviderContextData); err != nil {
			log.Errorf("invalid --context-data: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderSession) > 0 {
		input.Session = aws.String(_cognitoidentityproviderSession)
	}

	if resp, err := client.AdminInitiateAuth(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Links an existing user account in a user pool, or DestinationUser , to an
// identity from an external IdP, or SourceUser , based on a specified attribute
// name and value from the external IdP.
//
// This operation connects a local user profile with a user identity who hasn't
// yet signed in from their third-party IdP. When the user signs in with their IdP,
// they get access-control configuration from the local user profile. Linked local
// users can also sign in with SDK-based API operations like InitiateAuth after
// they sign in at least once through their IdP. For more information, see [Linking federated users].
//
// The maximum number of federated identities linked to a user is five.
//
// Because this API allows a user with an external federated identity to sign in
// as a local user, it is critical that it only be used with external IdPs and
// linked attributes that you trust.
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Linking federated users]: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-identity-federation-consolidate-users.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func cognitoidentityprovider_AdminLinkProviderForUser(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.AdminLinkProviderForUserInput{
		// DestinationUser: *types.ProviderUserIdentifierType, // Required
		// SourceUser: *types.ProviderUserIdentifierType, // Required
		// UserPoolId: *string, // Required
	}

	if len(_cognitoidentityproviderDestinationUser) > 0 {
		if err := assignInputField(input, "DestinationUser", _cognitoidentityproviderDestinationUser); err != nil {
			log.Errorf("invalid --destination-user: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderSourceUser) > 0 {
		if err := assignInputField(input, "SourceUser", _cognitoidentityproviderSourceUser); err != nil {
			log.Errorf("invalid --source-user: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}

	if resp, err := client.AdminLinkProviderForUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists a user's registered devices. Remembered devices are used in
// authentication services where you offer a "Remember me" option for users who you
// want to permit to sign in without MFA from a trusted device. Users can bypass
// MFA while your application performs device SRP authentication on the back end.
// For more information, see [Working with devices].
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [Working with devices]: https://docs.aws.amazon.com/cognito/latest/developerguide/amazon-cognito-user-pools-device-tracking.html
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func cognitoidentityprovider_AdminListDevices(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.AdminListDevicesInput{
		// UserPoolId: *string, // Required
		// Username: *string, // Required
	}

	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}
	if len(_cognitoidentityproviderUsername) > 0 {
		input.Username = aws.String(_cognitoidentityproviderUsername)
	}
	if len(_cognitoidentityproviderLimit) > 0 {
		if err := assignInputField(input, "Limit", _cognitoidentityproviderLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderPaginationToken) > 0 {
		input.PaginationToken = aws.String(_cognitoidentityproviderPaginationToken)
	}

	if resp, err := client.AdminListDevices(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the groups that a user belongs to. User pool groups are identifiers that
// you can reference from the contents of ID and access tokens, and set preferred
// IAM roles for identity-pool authentication. For more information, see [Adding groups to a user pool].
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Adding groups to a user pool]: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-user-groups.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func cognitoidentityprovider_AdminListGroupsForUser(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.AdminListGroupsForUserInput{
		// UserPoolId: *string, // Required
		// Username: *string, // Required
	}

	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}
	if len(_cognitoidentityproviderUsername) > 0 {
		input.Username = aws.String(_cognitoidentityproviderUsername)
	}
	if len(_cognitoidentityproviderLimit) > 0 {
		if err := assignInputField(input, "Limit", _cognitoidentityproviderLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderNextToken) > 0 {
		input.NextToken = aws.String(_cognitoidentityproviderNextToken)
	}

	if disablePaginator() {
		if resp, err := client.AdminListGroupsForUser(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cognitoidentityprovider.AdminListGroupsForUserOutput
	p := cognitoidentityprovider.NewAdminListGroupsForUserPaginator(client, input)
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

// Requests a history of user activity and any risks detected as part of Amazon
// Cognito threat protection. For more information, see [Viewing user event history].
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Viewing user event history]: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pool-settings-adaptive-authentication.html#user-pool-settings-adaptive-authentication-event-user-history
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func cognitoidentityprovider_AdminListUserAuthEvents(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.AdminListUserAuthEventsInput{
		// UserPoolId: *string, // Required
		// Username: *string, // Required
	}

	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}
	if len(_cognitoidentityproviderUsername) > 0 {
		input.Username = aws.String(_cognitoidentityproviderUsername)
	}
	if len(_cognitoidentityproviderMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cognitoidentityproviderMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderNextToken) > 0 {
		input.NextToken = aws.String(_cognitoidentityproviderNextToken)
	}

	if disablePaginator() {
		if resp, err := client.AdminListUserAuthEvents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cognitoidentityprovider.AdminListUserAuthEventsOutput
	p := cognitoidentityprovider.NewAdminListUserAuthEventsPaginator(client, input)
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

// Given a username and a group name, removes them from the group. User pool
// groups are identifiers that you can reference from the contents of ID and access
// tokens, and set preferred IAM roles for identity-pool authentication. For more
// information, see [Adding groups to a user pool].
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Adding groups to a user pool]: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-user-groups.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func cognitoidentityprovider_AdminRemoveUserFromGroup(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.AdminRemoveUserFromGroupInput{
		// GroupName: *string, // Required
		// UserPoolId: *string, // Required
		// Username: *string, // Required
	}

	if len(_cognitoidentityproviderGroupName) > 0 {
		input.GroupName = aws.String(_cognitoidentityproviderGroupName)
	}
	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}
	if len(_cognitoidentityproviderUsername) > 0 {
		input.Username = aws.String(_cognitoidentityproviderUsername)
	}

	if resp, err := client.AdminRemoveUserFromGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Begins the password reset process. Sets the requested user’s account into a
// RESET_REQUIRED status, and sends them a password-reset code. Your user pool also
// sends the user a notification with a reset code and the information that their
// password has been reset. At sign-in, your application or the managed login
// session receives a challenge to complete the reset by confirming the code and
// setting a new password.
//
// To use this API operation, your user pool must have self-service account
// recovery configured.
//
// This action might generate an SMS text message. Starting June 1, 2021, US
// telecom carriers require you to register an origination phone number before you
// can send SMS messages to US phone numbers. If you use SMS text messages in
// Amazon Cognito, you must register a phone number with [Amazon Pinpoint]. Amazon Cognito uses the
// registered number automatically. Otherwise, Amazon Cognito users who must
// receive SMS messages might not be able to sign up, activate their accounts, or
// sign in.
//
// If you have never used SMS text messages with Amazon Cognito or any other
// Amazon Web Services service, Amazon Simple Notification Service might place your
// account in the SMS sandbox. In [sandbox mode], you can send messages only to verified phone
// numbers. After you test your app while in the sandbox environment, you can move
// out of the sandbox and into production. For more information, see [SMS message settings for Amazon Cognito user pools]in the Amazon
// Cognito Developer Guide.
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [SMS message settings for Amazon Cognito user pools]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-sms-settings.html
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [sandbox mode]: https://docs.aws.amazon.com/sns/latest/dg/sns-sms-sandbox.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
// [Amazon Pinpoint]: https://console.aws.amazon.com/pinpoint/home/
func cognitoidentityprovider_AdminResetUserPassword(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.AdminResetUserPasswordInput{
		// UserPoolId: *string, // Required
		// Username: *string, // Required
	}

	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}
	if len(_cognitoidentityproviderUsername) > 0 {
		input.Username = aws.String(_cognitoidentityproviderUsername)
	}
	if len(_cognitoidentityproviderClientMetadata) > 0 {
		if err := assignInputField(input, "ClientMetadata", _cognitoidentityproviderClientMetadata); err != nil {
			log.Errorf("invalid --client-metadata: %s", err.Error())
			return
		}
	}

	if resp, err := client.AdminResetUserPassword(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Some API operations in a user pool generate a challenge, like a prompt for an
// MFA code, for device authentication that bypasses MFA, or for a custom
// authentication challenge. An AdminRespondToAuthChallenge API request provides
// the answer to that challenge, like a code or a secure remote password (SRP). The
// parameters of a response to an authentication challenge vary with the type of
// challenge.
//
// For more information about custom authentication challenges, see [Custom authentication challenge Lambda triggers].
//
// This action might generate an SMS text message. Starting June 1, 2021, US
// telecom carriers require you to register an origination phone number before you
// can send SMS messages to US phone numbers. If you use SMS text messages in
// Amazon Cognito, you must register a phone number with [Amazon Pinpoint]. Amazon Cognito uses the
// registered number automatically. Otherwise, Amazon Cognito users who must
// receive SMS messages might not be able to sign up, activate their accounts, or
// sign in.
//
// If you have never used SMS text messages with Amazon Cognito or any other
// Amazon Web Services service, Amazon Simple Notification Service might place your
// account in the SMS sandbox. In [sandbox mode], you can send messages only to verified phone
// numbers. After you test your app while in the sandbox environment, you can move
// out of the sandbox and into production. For more information, see [SMS message settings for Amazon Cognito user pools]in the Amazon
// Cognito Developer Guide.
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [SMS message settings for Amazon Cognito user pools]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-sms-settings.html
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [sandbox mode]: https://docs.aws.amazon.com/sns/latest/dg/sns-sms-sandbox.html
// [Custom authentication challenge Lambda triggers]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-lambda-challenge.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
// [Amazon Pinpoint]: https://console.aws.amazon.com/pinpoint/home/
func cognitoidentityprovider_AdminRespondToAuthChallenge(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.AdminRespondToAuthChallengeInput{
		// ChallengeName: types.ChallengeNameType, // Required
		// ClientId: *string, // Required
		// UserPoolId: *string, // Required
	}

	if len(_cognitoidentityproviderChallengeName) > 0 {
		if err := assignInputField(input, "ChallengeName", _cognitoidentityproviderChallengeName); err != nil {
			log.Errorf("invalid --challenge-name: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderClientId) > 0 {
		input.ClientId = aws.String(_cognitoidentityproviderClientId)
	}
	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}
	if len(_cognitoidentityproviderAnalyticsMetadata) > 0 {
		if err := assignInputField(input, "AnalyticsMetadata", _cognitoidentityproviderAnalyticsMetadata); err != nil {
			log.Errorf("invalid --analytics-metadata: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderChallengeResponses) > 0 {
		if err := assignInputField(input, "ChallengeResponses", _cognitoidentityproviderChallengeResponses); err != nil {
			log.Errorf("invalid --challenge-responses: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderClientMetadata) > 0 {
		if err := assignInputField(input, "ClientMetadata", _cognitoidentityproviderClientMetadata); err != nil {
			log.Errorf("invalid --client-metadata: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderContextData) > 0 {
		if err := assignInputField(input, "ContextData", _cognitoidentityproviderContextData); err != nil {
			log.Errorf("invalid --context-data: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderSession) > 0 {
		input.Session = aws.String(_cognitoidentityproviderSession)
	}

	if resp, err := client.AdminRespondToAuthChallenge(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets the user's multi-factor authentication (MFA) preference, including which
// MFA options are activated, and if any are preferred. Only one factor can be set
// as preferred. The preferred MFA factor will be used to authenticate a user if
// multiple factors are activated. If multiple options are activated and no
// preference is set, a challenge to choose an MFA option will be returned during
// sign-in.
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func cognitoidentityprovider_AdminSetUserMFAPreference(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.AdminSetUserMFAPreferenceInput{
		// UserPoolId: *string, // Required
		// Username: *string, // Required
	}

	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}
	if len(_cognitoidentityproviderUsername) > 0 {
		input.Username = aws.String(_cognitoidentityproviderUsername)
	}
	if len(_cognitoidentityproviderEmailMfaSettings) > 0 {
		if err := assignInputField(input, "EmailMfaSettings", _cognitoidentityproviderEmailMfaSettings); err != nil {
			log.Errorf("invalid --email-mfa-settings: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderSMSMfaSettings) > 0 {
		if err := assignInputField(input, "SMSMfaSettings", _cognitoidentityproviderSMSMfaSettings); err != nil {
			log.Errorf("invalid --sms-mfa-settings: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderSoftwareTokenMfaSettings) > 0 {
		if err := assignInputField(input, "SoftwareTokenMfaSettings", _cognitoidentityproviderSoftwareTokenMfaSettings); err != nil {
			log.Errorf("invalid --software-token-mfa-settings: %s", err.Error())
			return
		}
	}

	if resp, err := client.AdminSetUserMFAPreference(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets the specified user's password in a user pool. This operation
// administratively sets a temporary or permanent password for a user. With this
// operation, you can bypass self-service password changes and permit immediate
// sign-in with the password that you set. To do this, set Permanent to true .
//
// You can also set a new temporary password in this request, send it to a user,
// and require them to choose a new password on their next sign-in. To do this, set
// Permanent to false .
//
// If the password is temporary, the user's Status becomes FORCE_CHANGE_PASSWORD .
// When the user next tries to sign in, the InitiateAuth or AdminInitiateAuth
// response includes the NEW_PASSWORD_REQUIRED challenge. If the user doesn't sign
// in before the temporary password expires, they can no longer sign in and you
// must repeat this operation to set a temporary or permanent password for them.
//
// After the user sets a new password, or if you set a permanent password, their
// status becomes Confirmed .
//
// AdminSetUserPassword can set a password for the user profile that Amazon
// Cognito creates for third-party federated users. When you set a password, the
// federated user's status changes from EXTERNAL_PROVIDER to CONFIRMED . A user in
// this state can sign in as a federated user, and initiate authentication flows in
// the API like a linked native user. They can also modify their password and
// attributes in token-authenticated API requests like ChangePassword and
// UpdateUserAttributes . As a best security practice and to keep users in sync
// with your external IdP, don't set passwords on federated user profiles. To set
// up a federated user for native sign-in with a linked native user, refer to [Linking federated users to an existing user profile].
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [Linking federated users to an existing user profile]: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-identity-federation-consolidate-users.html
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func cognitoidentityprovider_AdminSetUserPassword(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.AdminSetUserPasswordInput{
		// Password: *string, // Required
		// UserPoolId: *string, // Required
		// Username: *string, // Required
	}

	if len(_cognitoidentityproviderPassword) > 0 {
		input.Password = aws.String(_cognitoidentityproviderPassword)
	}
	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}
	if len(_cognitoidentityproviderUsername) > 0 {
		input.Username = aws.String(_cognitoidentityproviderUsername)
	}
	if len(_cognitoidentityproviderPermanent) > 0 {
		if err := assignInputField(input, "Permanent", _cognitoidentityproviderPermanent); err != nil {
			log.Errorf("invalid --permanent: %s", err.Error())
			return
		}
	}

	if resp, err := client.AdminSetUserPassword(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This action is no longer supported. You can use it to configure only SMS MFA.
// You can't use it to configure time-based one-time password (TOTP) software token
// MFA.
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func cognitoidentityprovider_AdminSetUserSettings(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.AdminSetUserSettingsInput{
		// MFAOptions: []types.MFAOptionType, // Required
		// UserPoolId: *string, // Required
		// Username: *string, // Required
	}

	if len(_cognitoidentityproviderMFAOptions) > 0 {
		if err := assignInputField(input, "MFAOptions", _cognitoidentityproviderMFAOptions); err != nil {
			log.Errorf("invalid --mfa-options: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}
	if len(_cognitoidentityproviderUsername) > 0 {
		input.Username = aws.String(_cognitoidentityproviderUsername)
	}

	if resp, err := client.AdminSetUserSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides the feedback for an authentication event generated by threat
// protection features. Your response indicates that you think that the event
// either was from a valid user or was an unwanted authentication attempt. This
// feedback improves the risk evaluation decision for the user pool as part of
// Amazon Cognito threat protection. To activate this setting, your user pool must
// be on the [Plus tier].
//
// To train the threat-protection model to recognize trusted and untrusted sign-in
// characteristics, configure threat protection in audit-only mode and provide a
// mechanism for users or administrators to submit feedback. Your feedback can tell
// Amazon Cognito that a risk rating was assigned at a level you don't agree with.
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Plus tier]: https://docs.aws.amazon.com/cognito/latest/developerguide/feature-plans-features-plus.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func cognitoidentityprovider_AdminUpdateAuthEventFeedback(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.AdminUpdateAuthEventFeedbackInput{
		// EventId: *string, // Required
		// FeedbackValue: types.FeedbackValueType, // Required
		// UserPoolId: *string, // Required
		// Username: *string, // Required
	}

	if len(_cognitoidentityproviderEventId) > 0 {
		input.EventId = aws.String(_cognitoidentityproviderEventId)
	}
	if len(_cognitoidentityproviderFeedbackValue) > 0 {
		if err := assignInputField(input, "FeedbackValue", _cognitoidentityproviderFeedbackValue); err != nil {
			log.Errorf("invalid --feedback-value: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}
	if len(_cognitoidentityproviderUsername) > 0 {
		input.Username = aws.String(_cognitoidentityproviderUsername)
	}

	if resp, err := client.AdminUpdateAuthEventFeedback(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the status of a user's device so that it is marked as remembered or not
// remembered for the purpose of device authentication. Device authentication is a
// "remember me" mechanism that silently completes sign-in from trusted devices
// with a device key instead of a user-provided MFA code. This operation changes
// the status of a device without deleting it, so you can enable it again later.
// For more information about device authentication, see [Working with devices].
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [Working with devices]: https://docs.aws.amazon.com/cognito/latest/developerguide/amazon-cognito-user-pools-device-tracking.html
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func cognitoidentityprovider_AdminUpdateDeviceStatus(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.AdminUpdateDeviceStatusInput{
		// DeviceKey: *string, // Required
		// UserPoolId: *string, // Required
		// Username: *string, // Required
	}

	if len(_cognitoidentityproviderDeviceKey) > 0 {
		input.DeviceKey = aws.String(_cognitoidentityproviderDeviceKey)
	}
	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}
	if len(_cognitoidentityproviderUsername) > 0 {
		input.Username = aws.String(_cognitoidentityproviderUsername)
	}
	if len(_cognitoidentityproviderDeviceRememberedStatus) > 0 {
		if err := assignInputField(input, "DeviceRememberedStatus", _cognitoidentityproviderDeviceRememberedStatus); err != nil {
			log.Errorf("invalid --device-remembered-status: %s", err.Error())
			return
		}
	}

	if resp, err := client.AdminUpdateDeviceStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified user's attributes. To delete an attribute from your user,
// submit the attribute in your API request with a blank value.
//
// For custom attributes, you must add a custom: prefix to the attribute name, for
// example custom:department .
//
// This operation can set a user's email address or phone number as verified and
// permit immediate sign-in in user pools that require verification of these
// attributes. To do this, set the email_verified or phone_number_verified
// attribute to true .
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// This action might generate an SMS text message. Starting June 1, 2021, US
// telecom carriers require you to register an origination phone number before you
// can send SMS messages to US phone numbers. If you use SMS text messages in
// Amazon Cognito, you must register a phone number with [Amazon Pinpoint]. Amazon Cognito uses the
// registered number automatically. Otherwise, Amazon Cognito users who must
// receive SMS messages might not be able to sign up, activate their accounts, or
// sign in.
//
// If you have never used SMS text messages with Amazon Cognito or any other
// Amazon Web Services service, Amazon Simple Notification Service might place your
// account in the SMS sandbox. In [sandbox mode], you can send messages only to verified phone
// numbers. After you test your app while in the sandbox environment, you can move
// out of the sandbox and into production. For more information, see [SMS message settings for Amazon Cognito user pools]in the Amazon
// Cognito Developer Guide.
//
// [SMS message settings for Amazon Cognito user pools]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-sms-settings.html
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [sandbox mode]: https://docs.aws.amazon.com/sns/latest/dg/sns-sms-sandbox.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
// [Amazon Pinpoint]: https://console.aws.amazon.com/pinpoint/home/
func cognitoidentityprovider_AdminUpdateUserAttributes(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.AdminUpdateUserAttributesInput{
		// UserAttributes: []types.AttributeType, // Required
		// UserPoolId: *string, // Required
		// Username: *string, // Required
	}

	if len(_cognitoidentityproviderUserAttributes) > 0 {
		if err := assignInputField(input, "UserAttributes", _cognitoidentityproviderUserAttributes); err != nil {
			log.Errorf("invalid --user-attributes: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}
	if len(_cognitoidentityproviderUsername) > 0 {
		input.Username = aws.String(_cognitoidentityproviderUsername)
	}
	if len(_cognitoidentityproviderClientMetadata) > 0 {
		if err := assignInputField(input, "ClientMetadata", _cognitoidentityproviderClientMetadata); err != nil {
			log.Errorf("invalid --client-metadata: %s", err.Error())
			return
		}
	}

	if resp, err := client.AdminUpdateUserAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Invalidates the identity, access, and refresh tokens that Amazon Cognito issued
// to a user. Call this operation with your administrative credentials when your
// user signs out of your app. This results in the following behavior.
//
// - Amazon Cognito no longer accepts token-authorized user operations that you
// authorize with a signed-out user's access tokens. For more information, see [Using the Amazon Cognito user pools API and user pool endpoints].
//
// # Amazon Cognito returns an Access Token has been revoked error when your app
//
// attempts to authorize a user pools API request with a revoked access token that
// contains the scope aws.cognito.signin.user.admin .
//
// - Amazon Cognito no longer accepts a signed-out user's ID token in a [GetId]request
// to an identity pool with ServerSideTokenCheck enabled for its user pool IdP
// configuration in [CognitoIdentityProvider].
//
// - Amazon Cognito no longer accepts a signed-out user's refresh tokens in
// refresh requests.
//
// Other requests might be valid until your user's token expires. This operation
// doesn't clear the [managed login]session cookie. To clear the session for a user who signed in
// with managed login or the classic hosted UI, direct their browser session to the
// [logout endpoint].
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [logout endpoint]: https://docs.aws.amazon.com/cognito/latest/developerguide/logout-endpoint.html
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [managed login]: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-managed-login.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
// [CognitoIdentityProvider]: https://docs.aws.amazon.com/cognitoidentity/latest/APIReference/API_CognitoIdentityProvider.html
// [GetId]: https://docs.aws.amazon.com/cognitoidentity/latest/APIReference/API_GetId.html
func cognitoidentityprovider_AdminUserGlobalSignOut(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.AdminUserGlobalSignOutInput{
		// UserPoolId: *string, // Required
		// Username: *string, // Required
	}

	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}
	if len(_cognitoidentityproviderUsername) > 0 {
		input.Username = aws.String(_cognitoidentityproviderUsername)
	}

	if resp, err := client.AdminUserGlobalSignOut(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Begins setup of time-based one-time password (TOTP) multi-factor authentication
// (MFA) for a user, with a unique private key that Amazon Cognito generates and
// returns in the API response. You can authorize an AssociateSoftwareToken
// request with either the user's access token, or a session string from a
// challenge response that you received from Amazon Cognito.
//
// Amazon Cognito doesn't evaluate Identity and Access Management (IAM) policies
// in requests for this API operation. For this operation, you can't use IAM
// credentials to authorize requests, and you can't grant IAM permissions in
// policies. For more information about authorization models in Amazon Cognito, see
// [Using the Amazon Cognito user pools API and user pool endpoints].
//
// Authorize this action with a signed-in user's access token. It must include the
// scope aws.cognito.signin.user.admin .
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
func cognitoidentityprovider_AssociateSoftwareToken(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.AssociateSoftwareTokenInput{}

	if len(_cognitoidentityproviderAccessToken) > 0 {
		input.AccessToken = aws.String(_cognitoidentityproviderAccessToken)
	}
	if len(_cognitoidentityproviderSession) > 0 {
		input.Session = aws.String(_cognitoidentityproviderSession)
	}

	if resp, err := client.AssociateSoftwareToken(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Changes the password for the currently signed-in user.
// Authorize this action with a signed-in user's access token. It must include the
// scope aws.cognito.signin.user.admin .
//
// Amazon Cognito doesn't evaluate Identity and Access Management (IAM) policies
// in requests for this API operation. For this operation, you can't use IAM
// credentials to authorize requests, and you can't grant IAM permissions in
// policies. For more information about authorization models in Amazon Cognito, see
// [Using the Amazon Cognito user pools API and user pool endpoints].
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
func cognitoidentityprovider_ChangePassword(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.ChangePasswordInput{
		// AccessToken: *string, // Required
		// ProposedPassword: *string, // Required
	}

	if len(_cognitoidentityproviderAccessToken) > 0 {
		input.AccessToken = aws.String(_cognitoidentityproviderAccessToken)
	}
	if len(_cognitoidentityproviderProposedPassword) > 0 {
		input.ProposedPassword = aws.String(_cognitoidentityproviderProposedPassword)
	}
	if len(_cognitoidentityproviderPreviousPassword) > 0 {
		input.PreviousPassword = aws.String(_cognitoidentityproviderPreviousPassword)
	}

	if resp, err := client.ChangePassword(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Completes registration of a passkey authenticator for the currently signed-in
// user.
//
// Authorize this action with a signed-in user's access token. It must include the
// scope aws.cognito.signin.user.admin .
func cognitoidentityprovider_CompleteWebAuthnRegistration(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.CompleteWebAuthnRegistrationInput{
		// AccessToken: *string, // Required
		// Credential: document.Interface, // Required
	}

	if len(_cognitoidentityproviderAccessToken) > 0 {
		input.AccessToken = aws.String(_cognitoidentityproviderAccessToken)
	}
	if len(_cognitoidentityproviderCredential) > 0 {
		if err := assignInputField(input, "Credential", _cognitoidentityproviderCredential); err != nil {
			log.Errorf("invalid --credential: %s", err.Error())
			return
		}
	}

	if resp, err := client.CompleteWebAuthnRegistration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Confirms a device that a user wants to remember. A remembered device is a
// "Remember me on this device" option for user pools that perform authentication
// with the device key of a trusted device in the back end, instead of a
// user-provided MFA code. For more information about device authentication, see [Working with user devices in your user pool].
//
// Authorize this action with a signed-in user's access token. It must include the
// scope aws.cognito.signin.user.admin .
//
// Amazon Cognito doesn't evaluate Identity and Access Management (IAM) policies
// in requests for this API operation. For this operation, you can't use IAM
// credentials to authorize requests, and you can't grant IAM permissions in
// policies. For more information about authorization models in Amazon Cognito, see
// [Using the Amazon Cognito user pools API and user pool endpoints].
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Working with user devices in your user pool]: https://docs.aws.amazon.com/cognito/latest/developerguide/amazon-cognito-user-pools-device-tracking.html
func cognitoidentityprovider_ConfirmDevice(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.ConfirmDeviceInput{
		// AccessToken: *string, // Required
		// DeviceKey: *string, // Required
	}

	if len(_cognitoidentityproviderAccessToken) > 0 {
		input.AccessToken = aws.String(_cognitoidentityproviderAccessToken)
	}
	if len(_cognitoidentityproviderDeviceKey) > 0 {
		input.DeviceKey = aws.String(_cognitoidentityproviderDeviceKey)
	}
	if len(_cognitoidentityproviderDeviceName) > 0 {
		input.DeviceName = aws.String(_cognitoidentityproviderDeviceName)
	}
	if len(_cognitoidentityproviderDeviceSecretVerifierConfig) > 0 {
		if err := assignInputField(input, "DeviceSecretVerifierConfig", _cognitoidentityproviderDeviceSecretVerifierConfig); err != nil {
			log.Errorf("invalid --device-secret-verifier-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.ConfirmDevice(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This public API operation accepts a confirmation code that Amazon Cognito sent
// to a user and accepts a new password for that user.
//
// Amazon Cognito doesn't evaluate Identity and Access Management (IAM) policies
// in requests for this API operation. For this operation, you can't use IAM
// credentials to authorize requests, and you can't grant IAM permissions in
// policies. For more information about authorization models in Amazon Cognito, see
// [Using the Amazon Cognito user pools API and user pool endpoints].
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
func cognitoidentityprovider_ConfirmForgotPassword(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.ConfirmForgotPasswordInput{
		// ClientId: *string, // Required
		// ConfirmationCode: *string, // Required
		// Password: *string, // Required
		// Username: *string, // Required
	}

	if len(_cognitoidentityproviderClientId) > 0 {
		input.ClientId = aws.String(_cognitoidentityproviderClientId)
	}
	if len(_cognitoidentityproviderConfirmationCode) > 0 {
		input.ConfirmationCode = aws.String(_cognitoidentityproviderConfirmationCode)
	}
	if len(_cognitoidentityproviderPassword) > 0 {
		input.Password = aws.String(_cognitoidentityproviderPassword)
	}
	if len(_cognitoidentityproviderUsername) > 0 {
		input.Username = aws.String(_cognitoidentityproviderUsername)
	}
	if len(_cognitoidentityproviderAnalyticsMetadata) > 0 {
		if err := assignInputField(input, "AnalyticsMetadata", _cognitoidentityproviderAnalyticsMetadata); err != nil {
			log.Errorf("invalid --analytics-metadata: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderClientMetadata) > 0 {
		if err := assignInputField(input, "ClientMetadata", _cognitoidentityproviderClientMetadata); err != nil {
			log.Errorf("invalid --client-metadata: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderSecretHash) > 0 {
		input.SecretHash = aws.String(_cognitoidentityproviderSecretHash)
	}
	if len(_cognitoidentityproviderUserContextData) > 0 {
		if err := assignInputField(input, "UserContextData", _cognitoidentityproviderUserContextData); err != nil {
			log.Errorf("invalid --user-context-data: %s", err.Error())
			return
		}
	}

	if resp, err := client.ConfirmForgotPassword(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Confirms the account of a new user. This public API operation submits a code
// that Amazon Cognito sent to your user when they signed up in your user pool.
// After your user enters their code, they confirm ownership of the email address
// or phone number that they provided, and their user account becomes active.
// Depending on your user pool configuration, your users will receive their
// confirmation code in an email or SMS message.
//
// Local users who signed up in your user pool are the only type of user who can
// confirm sign-up with a code. Users who federate through an external identity
// provider (IdP) have already been confirmed by their IdP.
//
// Amazon Cognito doesn't evaluate Identity and Access Management (IAM) policies
// in requests for this API operation. For this operation, you can't use IAM
// credentials to authorize requests, and you can't grant IAM permissions in
// policies. For more information about authorization models in Amazon Cognito, see
// [Using the Amazon Cognito user pools API and user pool endpoints].
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
func cognitoidentityprovider_ConfirmSignUp(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.ConfirmSignUpInput{
		// ClientId: *string, // Required
		// ConfirmationCode: *string, // Required
		// Username: *string, // Required
	}

	if len(_cognitoidentityproviderClientId) > 0 {
		input.ClientId = aws.String(_cognitoidentityproviderClientId)
	}
	if len(_cognitoidentityproviderConfirmationCode) > 0 {
		input.ConfirmationCode = aws.String(_cognitoidentityproviderConfirmationCode)
	}
	if len(_cognitoidentityproviderUsername) > 0 {
		input.Username = aws.String(_cognitoidentityproviderUsername)
	}
	if len(_cognitoidentityproviderAnalyticsMetadata) > 0 {
		if err := assignInputField(input, "AnalyticsMetadata", _cognitoidentityproviderAnalyticsMetadata); err != nil {
			log.Errorf("invalid --analytics-metadata: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderClientMetadata) > 0 {
		if err := assignInputField(input, "ClientMetadata", _cognitoidentityproviderClientMetadata); err != nil {
			log.Errorf("invalid --client-metadata: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderForceAliasCreation) > 0 {
		if err := assignInputField(input, "ForceAliasCreation", _cognitoidentityproviderForceAliasCreation); err != nil {
			log.Errorf("invalid --force-alias-creation: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderSecretHash) > 0 {
		input.SecretHash = aws.String(_cognitoidentityproviderSecretHash)
	}
	if len(_cognitoidentityproviderSession) > 0 {
		input.Session = aws.String(_cognitoidentityproviderSession)
	}
	if len(_cognitoidentityproviderUserContextData) > 0 {
		if err := assignInputField(input, "UserContextData", _cognitoidentityproviderUserContextData); err != nil {
			log.Errorf("invalid --user-context-data: %s", err.Error())
			return
		}
	}

	if resp, err := client.ConfirmSignUp(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new group in the specified user pool. For more information about user
// pool groups, see [Adding groups to a user pool].
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Adding groups to a user pool]: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-user-groups.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func cognitoidentityprovider_CreateGroup(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.CreateGroupInput{
		// GroupName: *string, // Required
		// UserPoolId: *string, // Required
	}

	if len(_cognitoidentityproviderGroupName) > 0 {
		input.GroupName = aws.String(_cognitoidentityproviderGroupName)
	}
	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}
	if len(_cognitoidentityproviderDescription) > 0 {
		input.Description = aws.String(_cognitoidentityproviderDescription)
	}
	if len(_cognitoidentityproviderPrecedence) > 0 {
		if err := assignInputField(input, "Precedence", _cognitoidentityproviderPrecedence); err != nil {
			log.Errorf("invalid --precedence: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderRoleArn) > 0 {
		input.RoleArn = aws.String(_cognitoidentityproviderRoleArn)
	}

	if resp, err := client.CreateGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds a configuration and trust relationship between a third-party identity
// provider (IdP) and a user pool. Amazon Cognito accepts sign-in with third-party
// identity providers through managed login and OIDC relying-party libraries. For
// more information, see [Third-party IdP sign-in].
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Third-party IdP sign-in]: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-identity-federation.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func cognitoidentityprovider_CreateIdentityProvider(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.CreateIdentityProviderInput{
		// ProviderDetails: map[string]string, // Required
		// ProviderName: *string, // Required
		// ProviderType: types.IdentityProviderTypeType, // Required
		// UserPoolId: *string, // Required
	}

	if len(_cognitoidentityproviderProviderDetails) > 0 {
		if err := assignInputField(input, "ProviderDetails", _cognitoidentityproviderProviderDetails); err != nil {
			log.Errorf("invalid --provider-details: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderProviderName) > 0 {
		input.ProviderName = aws.String(_cognitoidentityproviderProviderName)
	}
	if len(_cognitoidentityproviderProviderType) > 0 {
		if err := assignInputField(input, "ProviderType", _cognitoidentityproviderProviderType); err != nil {
			log.Errorf("invalid --provider-type: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}
	if len(_cognitoidentityproviderAttributeMapping) > 0 {
		if err := assignInputField(input, "AttributeMapping", _cognitoidentityproviderAttributeMapping); err != nil {
			log.Errorf("invalid --attribute-mapping: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderIdpIdentifiers) > 0 {
		input.IdpIdentifiers = append([]string(nil), _cognitoidentityproviderIdpIdentifiers...)
	}

	if resp, err := client.CreateIdentityProvider(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new set of branding settings for a user pool style and associates it
// with an app client. This operation is the programmatic option for the creation
// of a new style in the branding editor.
//
// Provides values for UI customization in a Settings JSON object and image files
// in an Assets array. To send the JSON object Document type parameter in Settings
// , you might need to update to the most recent version of your Amazon Web
// Services SDK. To create a new style with default settings, set
// UseCognitoProvidedValues to true and don't provide values for any other options.
//
// This operation has a 2-megabyte request-size limit and include the CSS settings
// and image assets for your app client. Your branding settings might exceed 2MB in
// size. Amazon Cognito doesn't require that you pass all parameters in one request
// and preserves existing style settings that you don't specify. If your request is
// larger than 2MB, separate it into multiple requests, each with a size smaller
// than the limit.
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func cognitoidentityprovider_CreateManagedLoginBranding(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.CreateManagedLoginBrandingInput{
		// ClientId: *string, // Required
		// UserPoolId: *string, // Required
	}

	if len(_cognitoidentityproviderClientId) > 0 {
		input.ClientId = aws.String(_cognitoidentityproviderClientId)
	}
	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}
	if len(_cognitoidentityproviderAssets) > 0 {
		if err := assignInputField(input, "Assets", _cognitoidentityproviderAssets); err != nil {
			log.Errorf("invalid --assets: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderSettings) > 0 {
		if err := assignInputField(input, "Settings", _cognitoidentityproviderSettings); err != nil {
			log.Errorf("invalid --settings: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderUseCognitoProvidedValues) > 0 {
		if err := assignInputField(input, "UseCognitoProvidedValues", _cognitoidentityproviderUseCognitoProvidedValues); err != nil {
			log.Errorf("invalid --use-cognito-provided-values: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateManagedLoginBranding(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new OAuth2.0 resource server and defines custom scopes within it.
// Resource servers are associated with custom scopes and machine-to-machine (M2M)
// authorization. For more information, see [Access control with resource servers].
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Access control with resource servers]: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-define-resource-servers.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func cognitoidentityprovider_CreateResourceServer(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.CreateResourceServerInput{
		// Identifier: *string, // Required
		// Name: *string, // Required
		// UserPoolId: *string, // Required
	}

	if len(_cognitoidentityproviderIdentifier) > 0 {
		input.Identifier = aws.String(_cognitoidentityproviderIdentifier)
	}
	if len(_cognitoidentityproviderName) > 0 {
		input.Name = aws.String(_cognitoidentityproviderName)
	}
	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}
	if len(_cognitoidentityproviderScopes) > 0 {
		if err := assignInputField(input, "Scopes", _cognitoidentityproviderScopes); err != nil {
			log.Errorf("invalid --scopes: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateResourceServer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates terms documents for the requested app client. When Terms and conditions
// and Privacy policy documents are configured, the app client displays links to
// them in the sign-up page of managed login for the app client.
//
// You can provide URLs for terms documents in the languages that are supported by [managed login localization]
// . Amazon Cognito directs users to the terms documents for their current
// language, with fallback to default if no document exists for the language.
//
// Each request accepts one type of terms document and a map of language-to-link
// for that document type. You must provide both types of terms documents in at
// least one language before Amazon Cognito displays your terms documents. Supply
// each type in separate requests.
//
// For more information, see [Terms documents].
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Terms documents]: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-managed-login.html#managed-login-terms-documents
// [managed login localization]: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-managed-login.html#managed-login-localization
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func cognitoidentityprovider_CreateTerms(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.CreateTermsInput{
		// ClientId: *string, // Required
		// Enforcement: types.TermsEnforcementType, // Required
		// TermsName: *string, // Required
		// TermsSource: types.TermsSourceType, // Required
		// UserPoolId: *string, // Required
	}

	if len(_cognitoidentityproviderClientId) > 0 {
		input.ClientId = aws.String(_cognitoidentityproviderClientId)
	}
	if len(_cognitoidentityproviderEnforcement) > 0 {
		if err := assignInputField(input, "Enforcement", _cognitoidentityproviderEnforcement); err != nil {
			log.Errorf("invalid --enforcement: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderTermsName) > 0 {
		input.TermsName = aws.String(_cognitoidentityproviderTermsName)
	}
	if len(_cognitoidentityproviderTermsSource) > 0 {
		if err := assignInputField(input, "TermsSource", _cognitoidentityproviderTermsSource); err != nil {
			log.Errorf("invalid --terms-source: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}
	if len(_cognitoidentityproviderLinks) > 0 {
		if err := assignInputField(input, "Links", _cognitoidentityproviderLinks); err != nil {
			log.Errorf("invalid --links: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateTerms(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a user import job. You can import users into user pools from a
// comma-separated values (CSV) file without adding Amazon Cognito MAU costs to
// your Amazon Web Services bill.
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func cognitoidentityprovider_CreateUserImportJob(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.CreateUserImportJobInput{
		// CloudWatchLogsRoleArn: *string, // Required
		// JobName: *string, // Required
		// UserPoolId: *string, // Required
	}

	if len(_cognitoidentityproviderCloudWatchLogsRoleArn) > 0 {
		input.CloudWatchLogsRoleArn = aws.String(_cognitoidentityproviderCloudWatchLogsRoleArn)
	}
	if len(_cognitoidentityproviderJobName) > 0 {
		input.JobName = aws.String(_cognitoidentityproviderJobName)
	}
	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}

	if resp, err := client.CreateUserImportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new Amazon Cognito user pool. This operation sets basic and advanced
// configuration options.
//
// If you don't provide a value for an attribute, Amazon Cognito sets it to its
// default value.
//
// This action might generate an SMS text message. Starting June 1, 2021, US
// telecom carriers require you to register an origination phone number before you
// can send SMS messages to US phone numbers. If you use SMS text messages in
// Amazon Cognito, you must register a phone number with [Amazon Pinpoint]. Amazon Cognito uses the
// registered number automatically. Otherwise, Amazon Cognito users who must
// receive SMS messages might not be able to sign up, activate their accounts, or
// sign in.
//
// If you have never used SMS text messages with Amazon Cognito or any other
// Amazon Web Services service, Amazon Simple Notification Service might place your
// account in the SMS sandbox. In [sandbox mode], you can send messages only to verified phone
// numbers. After you test your app while in the sandbox environment, you can move
// out of the sandbox and into production. For more information, see [SMS message settings for Amazon Cognito user pools]in the Amazon
// Cognito Developer Guide.
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [SMS message settings for Amazon Cognito user pools]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-sms-settings.html
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [sandbox mode]: https://docs.aws.amazon.com/sns/latest/dg/sns-sms-sandbox.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
// [Amazon Pinpoint]: https://console.aws.amazon.com/pinpoint/home/
func cognitoidentityprovider_CreateUserPool(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.CreateUserPoolInput{
		// PoolName: *string, // Required
	}

	if len(_cognitoidentityproviderPoolName) > 0 {
		input.PoolName = aws.String(_cognitoidentityproviderPoolName)
	}
	if len(_cognitoidentityproviderAccountRecoverySetting) > 0 {
		if err := assignInputField(input, "AccountRecoverySetting", _cognitoidentityproviderAccountRecoverySetting); err != nil {
			log.Errorf("invalid --account-recovery-setting: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderAdminCreateUserConfig) > 0 {
		if err := assignInputField(input, "AdminCreateUserConfig", _cognitoidentityproviderAdminCreateUserConfig); err != nil {
			log.Errorf("invalid --admin-create-user-config: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderAliasAttributes) > 0 {
		if err := assignInputField(input, "AliasAttributes", _cognitoidentityproviderAliasAttributes); err != nil {
			log.Errorf("invalid --alias-attributes: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderAutoVerifiedAttributes) > 0 {
		if err := assignInputField(input, "AutoVerifiedAttributes", _cognitoidentityproviderAutoVerifiedAttributes); err != nil {
			log.Errorf("invalid --auto-verified-attributes: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderDeletionProtection) > 0 {
		if err := assignInputField(input, "DeletionProtection", _cognitoidentityproviderDeletionProtection); err != nil {
			log.Errorf("invalid --deletion-protection: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderDeviceConfiguration) > 0 {
		if err := assignInputField(input, "DeviceConfiguration", _cognitoidentityproviderDeviceConfiguration); err != nil {
			log.Errorf("invalid --device-configuration: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderEmailConfiguration) > 0 {
		if err := assignInputField(input, "EmailConfiguration", _cognitoidentityproviderEmailConfiguration); err != nil {
			log.Errorf("invalid --email-configuration: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderEmailVerificationMessage) > 0 {
		input.EmailVerificationMessage = aws.String(_cognitoidentityproviderEmailVerificationMessage)
	}
	if len(_cognitoidentityproviderEmailVerificationSubject) > 0 {
		input.EmailVerificationSubject = aws.String(_cognitoidentityproviderEmailVerificationSubject)
	}
	if len(_cognitoidentityproviderLambdaConfig) > 0 {
		if err := assignInputField(input, "LambdaConfig", _cognitoidentityproviderLambdaConfig); err != nil {
			log.Errorf("invalid --lambda-config: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderMfaConfiguration) > 0 {
		if err := assignInputField(input, "MfaConfiguration", _cognitoidentityproviderMfaConfiguration); err != nil {
			log.Errorf("invalid --mfa-configuration: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderPolicies) > 0 {
		if err := assignInputField(input, "Policies", _cognitoidentityproviderPolicies); err != nil {
			log.Errorf("invalid --policies: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderSchema) > 0 {
		if err := assignInputField(input, "Schema", _cognitoidentityproviderSchema); err != nil {
			log.Errorf("invalid --schema: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderSmsAuthenticationMessage) > 0 {
		input.SmsAuthenticationMessage = aws.String(_cognitoidentityproviderSmsAuthenticationMessage)
	}
	if len(_cognitoidentityproviderSmsConfiguration) > 0 {
		if err := assignInputField(input, "SmsConfiguration", _cognitoidentityproviderSmsConfiguration); err != nil {
			log.Errorf("invalid --sms-configuration: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderSmsVerificationMessage) > 0 {
		input.SmsVerificationMessage = aws.String(_cognitoidentityproviderSmsVerificationMessage)
	}
	if len(_cognitoidentityproviderUserAttributeUpdateSettings) > 0 {
		if err := assignInputField(input, "UserAttributeUpdateSettings", _cognitoidentityproviderUserAttributeUpdateSettings); err != nil {
			log.Errorf("invalid --user-attribute-update-settings: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderUserPoolAddOns) > 0 {
		if err := assignInputField(input, "UserPoolAddOns", _cognitoidentityproviderUserPoolAddOns); err != nil {
			log.Errorf("invalid --user-pool-add-ons: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderUserPoolTags) > 0 {
		if err := assignInputField(input, "UserPoolTags", _cognitoidentityproviderUserPoolTags); err != nil {
			log.Errorf("invalid --user-pool-tags: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderUserPoolTier) > 0 {
		if err := assignInputField(input, "UserPoolTier", _cognitoidentityproviderUserPoolTier); err != nil {
			log.Errorf("invalid --user-pool-tier: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderUsernameAttributes) > 0 {
		if err := assignInputField(input, "UsernameAttributes", _cognitoidentityproviderUsernameAttributes); err != nil {
			log.Errorf("invalid --username-attributes: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderUsernameConfiguration) > 0 {
		if err := assignInputField(input, "UsernameConfiguration", _cognitoidentityproviderUsernameConfiguration); err != nil {
			log.Errorf("invalid --username-configuration: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderVerificationMessageTemplate) > 0 {
		if err := assignInputField(input, "VerificationMessageTemplate", _cognitoidentityproviderVerificationMessageTemplate); err != nil {
			log.Errorf("invalid --verification-message-template: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateUserPool(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an app client in a user pool. This operation sets basic and advanced
// configuration options.
//
// Unlike app clients created in the console, Amazon Cognito doesn't automatically
// assign a branding style to app clients that you configure with this API
// operation. Managed login and classic hosted UI pages aren't available for your
// client until after you apply a branding style.
//
// If you don't provide a value for an attribute, Amazon Cognito sets it to its
// default value.
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func cognitoidentityprovider_CreateUserPoolClient(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.CreateUserPoolClientInput{
		// ClientName: *string, // Required
		// UserPoolId: *string, // Required
	}

	if len(_cognitoidentityproviderClientName) > 0 {
		input.ClientName = aws.String(_cognitoidentityproviderClientName)
	}
	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}
	if len(_cognitoidentityproviderAccessTokenValidity) > 0 {
		if err := assignInputField(input, "AccessTokenValidity", _cognitoidentityproviderAccessTokenValidity); err != nil {
			log.Errorf("invalid --access-token-validity: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderAllowedOAuthFlows) > 0 {
		if err := assignInputField(input, "AllowedOAuthFlows", _cognitoidentityproviderAllowedOAuthFlows); err != nil {
			log.Errorf("invalid --allowed-oauth-flows: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderAllowedOAuthFlowsUserPoolClient) > 0 {
		if err := assignInputField(input, "AllowedOAuthFlowsUserPoolClient", _cognitoidentityproviderAllowedOAuthFlowsUserPoolClient); err != nil {
			log.Errorf("invalid --allowed-oauth-flows-user-pool-client: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderAllowedOAuthScopes) > 0 {
		input.AllowedOAuthScopes = append([]string(nil), _cognitoidentityproviderAllowedOAuthScopes...)
	}
	if len(_cognitoidentityproviderAnalyticsConfiguration) > 0 {
		if err := assignInputField(input, "AnalyticsConfiguration", _cognitoidentityproviderAnalyticsConfiguration); err != nil {
			log.Errorf("invalid --analytics-configuration: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderAuthSessionValidity) > 0 {
		if err := assignInputField(input, "AuthSessionValidity", _cognitoidentityproviderAuthSessionValidity); err != nil {
			log.Errorf("invalid --auth-session-validity: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderCallbackURLs) > 0 {
		input.CallbackURLs = append([]string(nil), _cognitoidentityproviderCallbackURLs...)
	}
	if len(_cognitoidentityproviderClientSecret) > 0 {
		input.ClientSecret = aws.String(_cognitoidentityproviderClientSecret)
	}
	if len(_cognitoidentityproviderDefaultRedirectURI) > 0 {
		input.DefaultRedirectURI = aws.String(_cognitoidentityproviderDefaultRedirectURI)
	}
	if len(_cognitoidentityproviderEnablePropagateAdditionalUserContextData) > 0 {
		if err := assignInputField(input, "EnablePropagateAdditionalUserContextData", _cognitoidentityproviderEnablePropagateAdditionalUserContextData); err != nil {
			log.Errorf("invalid --enable-propagate-additional-user-context-data: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderEnableTokenRevocation) > 0 {
		if err := assignInputField(input, "EnableTokenRevocation", _cognitoidentityproviderEnableTokenRevocation); err != nil {
			log.Errorf("invalid --enable-token-revocation: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderExplicitAuthFlows) > 0 {
		if err := assignInputField(input, "ExplicitAuthFlows", _cognitoidentityproviderExplicitAuthFlows); err != nil {
			log.Errorf("invalid --explicit-auth-flows: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderGenerateSecret) > 0 {
		if err := assignInputField(input, "GenerateSecret", _cognitoidentityproviderGenerateSecret); err != nil {
			log.Errorf("invalid --generate-secret: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderIdTokenValidity) > 0 {
		if err := assignInputField(input, "IdTokenValidity", _cognitoidentityproviderIdTokenValidity); err != nil {
			log.Errorf("invalid --id-token-validity: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderLogoutURLs) > 0 {
		input.LogoutURLs = append([]string(nil), _cognitoidentityproviderLogoutURLs...)
	}
	if len(_cognitoidentityproviderPreventUserExistenceErrors) > 0 {
		if err := assignInputField(input, "PreventUserExistenceErrors", _cognitoidentityproviderPreventUserExistenceErrors); err != nil {
			log.Errorf("invalid --prevent-user-existence-errors: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderReadAttributes) > 0 {
		input.ReadAttributes = append([]string(nil), _cognitoidentityproviderReadAttributes...)
	}
	if len(_cognitoidentityproviderRefreshTokenRotation) > 0 {
		if err := assignInputField(input, "RefreshTokenRotation", _cognitoidentityproviderRefreshTokenRotation); err != nil {
			log.Errorf("invalid --refresh-token-rotation: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderRefreshTokenValidity) > 0 {
		if err := assignInputField(input, "RefreshTokenValidity", _cognitoidentityproviderRefreshTokenValidity); err != nil {
			log.Errorf("invalid --refresh-token-validity: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderSupportedIdentityProviders) > 0 {
		input.SupportedIdentityProviders = append([]string(nil), _cognitoidentityproviderSupportedIdentityProviders...)
	}
	if len(_cognitoidentityproviderTokenValidityUnits) > 0 {
		if err := assignInputField(input, "TokenValidityUnits", _cognitoidentityproviderTokenValidityUnits); err != nil {
			log.Errorf("invalid --token-validity-units: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderWriteAttributes) > 0 {
		input.WriteAttributes = append([]string(nil), _cognitoidentityproviderWriteAttributes...)
	}

	if resp, err := client.CreateUserPoolClient(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// A user pool domain hosts managed login, an authorization server and web server
// for authentication in your application. This operation creates a new user pool
// prefix domain or custom domain and sets the managed login branding version. Set
// the branding version to 1 for hosted UI (classic) or 2 for managed login. When
// you choose a custom domain, you must provide an SSL certificate in the US East
// (N. Virginia) Amazon Web Services Region in your request.
//
// Your prefix domain might take up to one minute to take effect. Your custom
// domain is online within five minutes, but it can take up to one hour to
// distribute your SSL certificate.
//
// For more information about adding a custom domain to your user pool, see [Configuring a user pool domain].
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Configuring a user pool domain]: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-add-custom-domain.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func cognitoidentityprovider_CreateUserPoolDomain(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.CreateUserPoolDomainInput{
		// Domain: *string, // Required
		// UserPoolId: *string, // Required
	}

	if len(_cognitoidentityproviderDomain) > 0 {
		input.Domain = aws.String(_cognitoidentityproviderDomain)
	}
	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}
	if len(_cognitoidentityproviderCustomDomainConfig) > 0 {
		if err := assignInputField(input, "CustomDomainConfig", _cognitoidentityproviderCustomDomainConfig); err != nil {
			log.Errorf("invalid --custom-domain-config: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderManagedLoginVersion) > 0 {
		if err := assignInputField(input, "ManagedLoginVersion", _cognitoidentityproviderManagedLoginVersion); err != nil {
			log.Errorf("invalid --managed-login-version: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateUserPoolDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a group from the specified user pool. When you delete a group, that
// group no longer contributes to users' cognito:preferred_group or cognito:groups
// claims, and no longer influence access-control decision that are based on group
// membership. For more information about user pool groups, see [Adding groups to a user pool].
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Adding groups to a user pool]: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-user-groups.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func cognitoidentityprovider_DeleteGroup(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.DeleteGroupInput{
		// GroupName: *string, // Required
		// UserPoolId: *string, // Required
	}

	if len(_cognitoidentityproviderGroupName) > 0 {
		input.GroupName = aws.String(_cognitoidentityproviderGroupName)
	}
	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}

	if resp, err := client.DeleteGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a user pool identity provider (IdP). After you delete an IdP, users can
// no longer sign in to your user pool through that IdP. For more information about
// user pool IdPs, see [Third-party IdP sign-in].
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Third-party IdP sign-in]: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-identity-federation.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func cognitoidentityprovider_DeleteIdentityProvider(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.DeleteIdentityProviderInput{
		// ProviderName: *string, // Required
		// UserPoolId: *string, // Required
	}

	if len(_cognitoidentityproviderProviderName) > 0 {
		input.ProviderName = aws.String(_cognitoidentityproviderProviderName)
	}
	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}

	if resp, err := client.DeleteIdentityProvider(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a managed login branding style. When you delete a style, you delete the
// branding association for an app client. When an app client doesn't have a style
// assigned, your managed login pages for that app client are nonfunctional until
// you create a new style or switch the domain branding version.
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func cognitoidentityprovider_DeleteManagedLoginBranding(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.DeleteManagedLoginBrandingInput{
		// ManagedLoginBrandingId: *string, // Required
		// UserPoolId: *string, // Required
	}

	if len(_cognitoidentityproviderManagedLoginBrandingId) > 0 {
		input.ManagedLoginBrandingId = aws.String(_cognitoidentityproviderManagedLoginBrandingId)
	}
	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}

	if resp, err := client.DeleteManagedLoginBranding(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a resource server. After you delete a resource server, users can no
// longer generate access tokens with scopes that are associate with that resource
// server.
//
// Resource servers are associated with custom scopes and machine-to-machine (M2M)
// authorization. For more information, see [Access control with resource servers].
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Access control with resource servers]: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-define-resource-servers.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func cognitoidentityprovider_DeleteResourceServer(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.DeleteResourceServerInput{
		// Identifier: *string, // Required
		// UserPoolId: *string, // Required
	}

	if len(_cognitoidentityproviderIdentifier) > 0 {
		input.Identifier = aws.String(_cognitoidentityproviderIdentifier)
	}
	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}

	if resp, err := client.DeleteResourceServer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the terms documents with the requested ID from your app client.
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func cognitoidentityprovider_DeleteTerms(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.DeleteTermsInput{
		// TermsId: *string, // Required
		// UserPoolId: *string, // Required
	}

	if len(_cognitoidentityproviderTermsId) > 0 {
		input.TermsId = aws.String(_cognitoidentityproviderTermsId)
	}
	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}

	if resp, err := client.DeleteTerms(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the profile of the currently signed-in user. A deleted user profile can
// no longer be used to sign in and can't be restored.
//
// Authorize this action with a signed-in user's access token. It must include the
// scope aws.cognito.signin.user.admin .
//
// Amazon Cognito doesn't evaluate Identity and Access Management (IAM) policies
// in requests for this API operation. For this operation, you can't use IAM
// credentials to authorize requests, and you can't grant IAM permissions in
// policies. For more information about authorization models in Amazon Cognito, see
// [Using the Amazon Cognito user pools API and user pool endpoints].
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
func cognitoidentityprovider_DeleteUser(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.DeleteUserInput{
		// AccessToken: *string, // Required
	}

	if len(_cognitoidentityproviderAccessToken) > 0 {
		input.AccessToken = aws.String(_cognitoidentityproviderAccessToken)
	}

	if resp, err := client.DeleteUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes attributes from the currently signed-in user. For example, your
// application can submit a request to this operation when a user wants to remove
// their birthdate attribute value.
//
// Authorize this action with a signed-in user's access token. It must include the
// scope aws.cognito.signin.user.admin .
//
// Amazon Cognito doesn't evaluate Identity and Access Management (IAM) policies
// in requests for this API operation. For this operation, you can't use IAM
// credentials to authorize requests, and you can't grant IAM permissions in
// policies. For more information about authorization models in Amazon Cognito, see
// [Using the Amazon Cognito user pools API and user pool endpoints].
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
func cognitoidentityprovider_DeleteUserAttributes(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.DeleteUserAttributesInput{
		// AccessToken: *string, // Required
		// UserAttributeNames: []string, // Required
	}

	if len(_cognitoidentityproviderAccessToken) > 0 {
		input.AccessToken = aws.String(_cognitoidentityproviderAccessToken)
	}
	if len(_cognitoidentityproviderUserAttributeNames) > 0 {
		input.UserAttributeNames = append([]string(nil), _cognitoidentityproviderUserAttributeNames...)
	}

	if resp, err := client.DeleteUserAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a user pool. After you delete a user pool, users can no longer sign in
// to any associated applications.
//
// When you delete a user pool, it's no longer visible or operational in your
// Amazon Web Services account. Amazon Cognito retains deleted user pools in an
// inactive state for 14 days, then begins a cleanup process that fully removes
// them from Amazon Web Services systems. In case of accidental deletion, contact
// Amazon Web Services Support within 14 days for restoration assistance.
//
// Amazon Cognito begins full deletion of all resources from deleted user pools
// after 14 days. In the case of large user pools, the cleanup process might take
// significant additional time before all user data is permanently deleted.
func cognitoidentityprovider_DeleteUserPool(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.DeleteUserPoolInput{
		// UserPoolId: *string, // Required
	}

	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}

	if resp, err := client.DeleteUserPool(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a user pool app client. After you delete an app client, users can no
// longer sign in to the associated application.
func cognitoidentityprovider_DeleteUserPoolClient(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.DeleteUserPoolClientInput{
		// ClientId: *string, // Required
		// UserPoolId: *string, // Required
	}

	if len(_cognitoidentityproviderClientId) > 0 {
		input.ClientId = aws.String(_cognitoidentityproviderClientId)
	}
	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}

	if resp, err := client.DeleteUserPoolClient(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a specific client secret from a user pool app client. You cannot delete
// the last remaining secret for an app client.
func cognitoidentityprovider_DeleteUserPoolClientSecret(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.DeleteUserPoolClientSecretInput{
		// ClientId: *string, // Required
		// ClientSecretId: *string, // Required
		// UserPoolId: *string, // Required
	}

	if len(_cognitoidentityproviderClientId) > 0 {
		input.ClientId = aws.String(_cognitoidentityproviderClientId)
	}
	if len(_cognitoidentityproviderClientSecretId) > 0 {
		input.ClientSecretId = aws.String(_cognitoidentityproviderClientSecretId)
	}
	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}

	if resp, err := client.DeleteUserPoolClientSecret(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Given a user pool ID and domain identifier, deletes a user pool domain. After
// you delete a user pool domain, your managed login pages and authorization server
// are no longer available.
func cognitoidentityprovider_DeleteUserPoolDomain(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.DeleteUserPoolDomainInput{
		// Domain: *string, // Required
		// UserPoolId: *string, // Required
	}

	if len(_cognitoidentityproviderDomain) > 0 {
		input.Domain = aws.String(_cognitoidentityproviderDomain)
	}
	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}

	if resp, err := client.DeleteUserPoolDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a registered passkey, or WebAuthn, authenticator for the currently
// signed-in user.
//
// Authorize this action with a signed-in user's access token. It must include the
// scope aws.cognito.signin.user.admin .
//
// Amazon Cognito doesn't evaluate Identity and Access Management (IAM) policies
// in requests for this API operation. For this operation, you can't use IAM
// credentials to authorize requests, and you can't grant IAM permissions in
// policies. For more information about authorization models in Amazon Cognito, see
// [Using the Amazon Cognito user pools API and user pool endpoints].
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
func cognitoidentityprovider_DeleteWebAuthnCredential(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.DeleteWebAuthnCredentialInput{
		// AccessToken: *string, // Required
		// CredentialId: *string, // Required
	}

	if len(_cognitoidentityproviderAccessToken) > 0 {
		input.AccessToken = aws.String(_cognitoidentityproviderAccessToken)
	}
	if len(_cognitoidentityproviderCredentialId) > 0 {
		input.CredentialId = aws.String(_cognitoidentityproviderCredentialId)
	}

	if resp, err := client.DeleteWebAuthnCredential(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Given a user pool ID and identity provider (IdP) name, returns details about
// the IdP.
func cognitoidentityprovider_DescribeIdentityProvider(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.DescribeIdentityProviderInput{
		// ProviderName: *string, // Required
		// UserPoolId: *string, // Required
	}

	if len(_cognitoidentityproviderProviderName) > 0 {
		input.ProviderName = aws.String(_cognitoidentityproviderProviderName)
	}
	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}

	if resp, err := client.DescribeIdentityProvider(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Given the ID of a managed login branding style, returns detailed information
// about the style.
func cognitoidentityprovider_DescribeManagedLoginBranding(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.DescribeManagedLoginBrandingInput{
		// ManagedLoginBrandingId: *string, // Required
		// UserPoolId: *string, // Required
	}

	if len(_cognitoidentityproviderManagedLoginBrandingId) > 0 {
		input.ManagedLoginBrandingId = aws.String(_cognitoidentityproviderManagedLoginBrandingId)
	}
	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}
	if len(_cognitoidentityproviderReturnMergedResources) > 0 {
		if err := assignInputField(input, "ReturnMergedResources", _cognitoidentityproviderReturnMergedResources); err != nil {
			log.Errorf("invalid --return-merged-resources: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeManagedLoginBranding(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Given the ID of a user pool app client, returns detailed information about the
// style assigned to the app client.
func cognitoidentityprovider_DescribeManagedLoginBrandingByClient(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.DescribeManagedLoginBrandingByClientInput{
		// ClientId: *string, // Required
		// UserPoolId: *string, // Required
	}

	if len(_cognitoidentityproviderClientId) > 0 {
		input.ClientId = aws.String(_cognitoidentityproviderClientId)
	}
	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}
	if len(_cognitoidentityproviderReturnMergedResources) > 0 {
		if err := assignInputField(input, "ReturnMergedResources", _cognitoidentityproviderReturnMergedResources); err != nil {
			log.Errorf("invalid --return-merged-resources: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeManagedLoginBrandingByClient(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a resource server. For more information about resource servers, see [Access control with resource servers].
//
// [Access control with resource servers]: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-define-resource-servers.html
func cognitoidentityprovider_DescribeResourceServer(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.DescribeResourceServerInput{
		// Identifier: *string, // Required
		// UserPoolId: *string, // Required
	}

	if len(_cognitoidentityproviderIdentifier) > 0 {
		input.Identifier = aws.String(_cognitoidentityproviderIdentifier)
	}
	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}

	if resp, err := client.DescribeResourceServer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Given an app client or user pool ID where threat protection is configured,
// describes the risk configuration. This operation returns details about adaptive
// authentication, compromised credentials, and IP-address allow- and denylists.
// For more information about threat protection, see [Threat protection].
//
// [Threat protection]: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pool-settings-threat-protection.html
func cognitoidentityprovider_DescribeRiskConfiguration(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.DescribeRiskConfigurationInput{
		// UserPoolId: *string, // Required
	}

	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}
	if len(_cognitoidentityproviderClientId) > 0 {
		input.ClientId = aws.String(_cognitoidentityproviderClientId)
	}

	if resp, err := client.DescribeRiskConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns details for the requested terms documents ID. For more information, see [Terms documents]
// .
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
// [Terms documents]: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-managed-login.html#managed-login-terms-documents
func cognitoidentityprovider_DescribeTerms(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.DescribeTermsInput{
		// TermsId: *string, // Required
		// UserPoolId: *string, // Required
	}

	if len(_cognitoidentityproviderTermsId) > 0 {
		input.TermsId = aws.String(_cognitoidentityproviderTermsId)
	}
	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}

	if resp, err := client.DescribeTerms(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a user import job. For more information about user CSV import, see [Importing users from a CSV file].
//
// [Importing users from a CSV file]: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-using-import-tool.html
func cognitoidentityprovider_DescribeUserImportJob(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.DescribeUserImportJobInput{
		// JobId: *string, // Required
		// UserPoolId: *string, // Required
	}

	if len(_cognitoidentityproviderJobId) > 0 {
		input.JobId = aws.String(_cognitoidentityproviderJobId)
	}
	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}

	if resp, err := client.DescribeUserImportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Given a user pool ID, returns configuration information. This operation is
// useful when you want to inspect an existing user pool and programmatically
// replicate the configuration to another user pool.
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func cognitoidentityprovider_DescribeUserPool(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.DescribeUserPoolInput{
		// UserPoolId: *string, // Required
	}

	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}

	if resp, err := client.DescribeUserPool(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Given an app client ID, returns configuration information. This operation is
// useful when you want to inspect an existing app client and programmatically
// replicate the configuration to another app client. For more information about
// app clients, see [App clients].
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [App clients]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-client-apps.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func cognitoidentityprovider_DescribeUserPoolClient(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.DescribeUserPoolClientInput{
		// ClientId: *string, // Required
		// UserPoolId: *string, // Required
	}

	if len(_cognitoidentityproviderClientId) > 0 {
		input.ClientId = aws.String(_cognitoidentityproviderClientId)
	}
	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}

	if resp, err := client.DescribeUserPoolClient(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Given a user pool domain name, returns information about the domain
// configuration.
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func cognitoidentityprovider_DescribeUserPoolDomain(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.DescribeUserPoolDomainInput{
		// Domain: *string, // Required
	}

	if len(_cognitoidentityproviderDomain) > 0 {
		input.Domain = aws.String(_cognitoidentityproviderDomain)
	}

	if resp, err := client.DescribeUserPoolDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Given a device key, deletes a remembered device as the currently signed-in
// user. For more information about device authentication, see [Working with user devices in your user pool].
//
// Authorize this action with a signed-in user's access token. It must include the
// scope aws.cognito.signin.user.admin .
//
// Amazon Cognito doesn't evaluate Identity and Access Management (IAM) policies
// in requests for this API operation. For this operation, you can't use IAM
// credentials to authorize requests, and you can't grant IAM permissions in
// policies. For more information about authorization models in Amazon Cognito, see
// [Using the Amazon Cognito user pools API and user pool endpoints].
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Working with user devices in your user pool]: https://docs.aws.amazon.com/cognito/latest/developerguide/amazon-cognito-user-pools-device-tracking.html
func cognitoidentityprovider_ForgetDevice(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.ForgetDeviceInput{
		// DeviceKey: *string, // Required
	}

	if len(_cognitoidentityproviderDeviceKey) > 0 {
		input.DeviceKey = aws.String(_cognitoidentityproviderDeviceKey)
	}
	if len(_cognitoidentityproviderAccessToken) > 0 {
		input.AccessToken = aws.String(_cognitoidentityproviderAccessToken)
	}

	if resp, err := client.ForgetDevice(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sends a password-reset confirmation code to the email address or phone number
// of the requested username. The message delivery method is determined by the
// user's available attributes and the AccountRecoverySetting configuration of the
// user pool.
//
// For the Username parameter, you can use the username or an email, phone, or
// preferred username alias.
//
// If neither a verified phone number nor a verified email exists, Amazon Cognito
// responds with an InvalidParameterException error . If your app client has a
// client secret and you don't provide a SECRET_HASH parameter, this API returns
// NotAuthorizedException .
//
// Amazon Cognito doesn't evaluate Identity and Access Management (IAM) policies
// in requests for this API operation. For this operation, you can't use IAM
// credentials to authorize requests, and you can't grant IAM permissions in
// policies. For more information about authorization models in Amazon Cognito, see
// [Using the Amazon Cognito user pools API and user pool endpoints].
//
// This action might generate an SMS text message. Starting June 1, 2021, US
// telecom carriers require you to register an origination phone number before you
// can send SMS messages to US phone numbers. If you use SMS text messages in
// Amazon Cognito, you must register a phone number with [Amazon Pinpoint]. Amazon Cognito uses the
// registered number automatically. Otherwise, Amazon Cognito users who must
// receive SMS messages might not be able to sign up, activate their accounts, or
// sign in.
//
// If you have never used SMS text messages with Amazon Cognito or any other
// Amazon Web Services service, Amazon Simple Notification Service might place your
// account in the SMS sandbox. In [sandbox mode], you can send messages only to verified phone
// numbers. After you test your app while in the sandbox environment, you can move
// out of the sandbox and into production. For more information, see [SMS message settings for Amazon Cognito user pools]in the Amazon
// Cognito Developer Guide.
//
// [SMS message settings for Amazon Cognito user pools]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-sms-settings.html
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [sandbox mode]: https://docs.aws.amazon.com/sns/latest/dg/sns-sms-sandbox.html
// [Amazon Pinpoint]: https://console.aws.amazon.com/pinpoint/home/
func cognitoidentityprovider_ForgotPassword(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.ForgotPasswordInput{
		// ClientId: *string, // Required
		// Username: *string, // Required
	}

	if len(_cognitoidentityproviderClientId) > 0 {
		input.ClientId = aws.String(_cognitoidentityproviderClientId)
	}
	if len(_cognitoidentityproviderUsername) > 0 {
		input.Username = aws.String(_cognitoidentityproviderUsername)
	}
	if len(_cognitoidentityproviderAnalyticsMetadata) > 0 {
		if err := assignInputField(input, "AnalyticsMetadata", _cognitoidentityproviderAnalyticsMetadata); err != nil {
			log.Errorf("invalid --analytics-metadata: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderClientMetadata) > 0 {
		if err := assignInputField(input, "ClientMetadata", _cognitoidentityproviderClientMetadata); err != nil {
			log.Errorf("invalid --client-metadata: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderSecretHash) > 0 {
		input.SecretHash = aws.String(_cognitoidentityproviderSecretHash)
	}
	if len(_cognitoidentityproviderUserContextData) > 0 {
		if err := assignInputField(input, "UserContextData", _cognitoidentityproviderUserContextData); err != nil {
			log.Errorf("invalid --user-context-data: %s", err.Error())
			return
		}
	}

	if resp, err := client.ForgotPassword(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Given a user pool ID, generates a comma-separated value (CSV) list populated
// with available user attributes in the user pool. This list is the header for the
// CSV file that determines the users in a user import job. Save the content of
// CSVHeader in the response as a .csv file and populate it with the usernames and
// attributes of users that you want to import. For more information about CSV user
// import, see [Importing users from a CSV file].
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Importing users from a CSV file]: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-using-import-tool.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func cognitoidentityprovider_GetCSVHeader(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.GetCSVHeaderInput{
		// UserPoolId: *string, // Required
	}

	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}

	if resp, err := client.GetCSVHeader(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Given a device key, returns information about a remembered device for the
// current user. For more information about device authentication, see [Working with user devices in your user pool].
//
// Authorize this action with a signed-in user's access token. It must include the
// scope aws.cognito.signin.user.admin .
//
// Amazon Cognito doesn't evaluate Identity and Access Management (IAM) policies
// in requests for this API operation. For this operation, you can't use IAM
// credentials to authorize requests, and you can't grant IAM permissions in
// policies. For more information about authorization models in Amazon Cognito, see
// [Using the Amazon Cognito user pools API and user pool endpoints].
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Working with user devices in your user pool]: https://docs.aws.amazon.com/cognito/latest/developerguide/amazon-cognito-user-pools-device-tracking.html
func cognitoidentityprovider_GetDevice(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.GetDeviceInput{
		// DeviceKey: *string, // Required
	}

	if len(_cognitoidentityproviderDeviceKey) > 0 {
		input.DeviceKey = aws.String(_cognitoidentityproviderDeviceKey)
	}
	if len(_cognitoidentityproviderAccessToken) > 0 {
		input.AccessToken = aws.String(_cognitoidentityproviderAccessToken)
	}

	if resp, err := client.GetDevice(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Given a user pool ID and a group name, returns information about the user group.
// For more information about user pool groups, see [Adding groups to a user pool].
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Adding groups to a user pool]: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-user-groups.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func cognitoidentityprovider_GetGroup(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.GetGroupInput{
		// GroupName: *string, // Required
		// UserPoolId: *string, // Required
	}

	if len(_cognitoidentityproviderGroupName) > 0 {
		input.GroupName = aws.String(_cognitoidentityproviderGroupName)
	}
	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}

	if resp, err := client.GetGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Given the identifier of an identity provider (IdP), for example examplecorp ,
// returns information about the user pool configuration for that IdP. For more
// information about IdPs, see [Third-party IdP sign-in].
//
// [Third-party IdP sign-in]: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-identity-federation.html
func cognitoidentityprovider_GetIdentityProviderByIdentifier(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.GetIdentityProviderByIdentifierInput{
		// IdpIdentifier: *string, // Required
		// UserPoolId: *string, // Required
	}

	if len(_cognitoidentityproviderIdpIdentifier) > 0 {
		input.IdpIdentifier = aws.String(_cognitoidentityproviderIdpIdentifier)
	}
	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}

	if resp, err := client.GetIdentityProviderByIdentifier(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Given a user pool ID, returns the logging configuration. User pools can export
// message-delivery error and threat-protection activity logs to external Amazon
// Web Services services. For more information, see [Exporting user pool logs].
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [Exporting user pool logs]: https://docs.aws.amazon.com/cognito/latest/developerguide/exporting-quotas-and-usage.html
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func cognitoidentityprovider_GetLogDeliveryConfiguration(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.GetLogDeliveryConfigurationInput{
		// UserPoolId: *string, // Required
	}

	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}

	if resp, err := client.GetLogDeliveryConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Given a user pool ID, returns the signing certificate for SAML 2.0 federation.
// Issued certificates are valid for 10 years from the date of issue. Amazon
// Cognito issues and assigns a new signing certificate annually. This renewal
// process returns a new value in the response to GetSigningCertificate , but
// doesn't invalidate the original certificate.
//
// For more information, see [Signing SAML requests].
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Signing SAML requests]: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-SAML-signing-encryption.html#cognito-user-pools-SAML-signing
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func cognitoidentityprovider_GetSigningCertificate(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.GetSigningCertificateInput{
		// UserPoolId: *string, // Required
	}

	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}

	if resp, err := client.GetSigningCertificate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Given a refresh token, issues new ID, access, and optionally refresh tokens for
// the user who owns the submitted token. This operation issues a new refresh token
// and invalidates the original refresh token after an optional grace period when
// refresh token rotation is enabled. If refresh token rotation is disabled, issues
// new ID and access tokens only.
func cognitoidentityprovider_GetTokensFromRefreshToken(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.GetTokensFromRefreshTokenInput{
		// ClientId: *string, // Required
		// RefreshToken: *string, // Required
	}

	if len(_cognitoidentityproviderClientId) > 0 {
		input.ClientId = aws.String(_cognitoidentityproviderClientId)
	}
	if len(_cognitoidentityproviderRefreshToken) > 0 {
		input.RefreshToken = aws.String(_cognitoidentityproviderRefreshToken)
	}
	if len(_cognitoidentityproviderClientMetadata) > 0 {
		if err := assignInputField(input, "ClientMetadata", _cognitoidentityproviderClientMetadata); err != nil {
			log.Errorf("invalid --client-metadata: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderClientSecret) > 0 {
		input.ClientSecret = aws.String(_cognitoidentityproviderClientSecret)
	}
	if len(_cognitoidentityproviderDeviceKey) > 0 {
		input.DeviceKey = aws.String(_cognitoidentityproviderDeviceKey)
	}

	if resp, err := client.GetTokensFromRefreshToken(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Given a user pool ID or app client, returns information about classic hosted UI
// branding that you applied, if any. Returns user-pool level branding information
// if no app client branding is applied, or if you don't specify an app client ID.
// Returns an empty object if you haven't applied hosted UI branding to either the
// client or the user pool. For more information, see [Hosted UI (classic) branding].
//
// [Hosted UI (classic) branding]: https://docs.aws.amazon.com/cognito/latest/developerguide/hosted-ui-classic-branding.html
func cognitoidentityprovider_GetUICustomization(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.GetUICustomizationInput{
		// UserPoolId: *string, // Required
	}

	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}
	if len(_cognitoidentityproviderClientId) > 0 {
		input.ClientId = aws.String(_cognitoidentityproviderClientId)
	}

	if resp, err := client.GetUICustomization(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets user attributes and and MFA settings for the currently signed-in user.
// Authorize this action with a signed-in user's access token. It must include the
// scope aws.cognito.signin.user.admin .
//
// Amazon Cognito doesn't evaluate Identity and Access Management (IAM) policies
// in requests for this API operation. For this operation, you can't use IAM
// credentials to authorize requests, and you can't grant IAM permissions in
// policies. For more information about authorization models in Amazon Cognito, see
// [Using the Amazon Cognito user pools API and user pool endpoints].
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
func cognitoidentityprovider_GetUser(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.GetUserInput{
		// AccessToken: *string, // Required
	}

	if len(_cognitoidentityproviderAccessToken) > 0 {
		input.AccessToken = aws.String(_cognitoidentityproviderAccessToken)
	}

	if resp, err := client.GetUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Given an attribute name, sends a user attribute verification code for the
// specified attribute name to the currently signed-in user.
//
// Authorize this action with a signed-in user's access token. It must include the
// scope aws.cognito.signin.user.admin .
//
// Amazon Cognito doesn't evaluate Identity and Access Management (IAM) policies
// in requests for this API operation. For this operation, you can't use IAM
// credentials to authorize requests, and you can't grant IAM permissions in
// policies. For more information about authorization models in Amazon Cognito, see
// [Using the Amazon Cognito user pools API and user pool endpoints].
//
// This action might generate an SMS text message. Starting June 1, 2021, US
// telecom carriers require you to register an origination phone number before you
// can send SMS messages to US phone numbers. If you use SMS text messages in
// Amazon Cognito, you must register a phone number with [Amazon Pinpoint]. Amazon Cognito uses the
// registered number automatically. Otherwise, Amazon Cognito users who must
// receive SMS messages might not be able to sign up, activate their accounts, or
// sign in.
//
// If you have never used SMS text messages with Amazon Cognito or any other
// Amazon Web Services service, Amazon Simple Notification Service might place your
// account in the SMS sandbox. In [sandbox mode], you can send messages only to verified phone
// numbers. After you test your app while in the sandbox environment, you can move
// out of the sandbox and into production. For more information, see [SMS message settings for Amazon Cognito user pools]in the Amazon
// Cognito Developer Guide.
//
// [SMS message settings for Amazon Cognito user pools]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-sms-settings.html
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [sandbox mode]: https://docs.aws.amazon.com/sns/latest/dg/sns-sms-sandbox.html
// [Amazon Pinpoint]: https://console.aws.amazon.com/pinpoint/home/
func cognitoidentityprovider_GetUserAttributeVerificationCode(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.GetUserAttributeVerificationCodeInput{
		// AccessToken: *string, // Required
		// AttributeName: *string, // Required
	}

	if len(_cognitoidentityproviderAccessToken) > 0 {
		input.AccessToken = aws.String(_cognitoidentityproviderAccessToken)
	}
	if len(_cognitoidentityproviderAttributeName) > 0 {
		input.AttributeName = aws.String(_cognitoidentityproviderAttributeName)
	}
	if len(_cognitoidentityproviderClientMetadata) > 0 {
		if err := assignInputField(input, "ClientMetadata", _cognitoidentityproviderClientMetadata); err != nil {
			log.Errorf("invalid --client-metadata: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetUserAttributeVerificationCode(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the authentication options for the currently signed-in user. Returns the
// following:
//
// - The user's multi-factor authentication (MFA) preferences.
//
// - The user's options for choice-based authentication with the USER_AUTH flow.
//
// Authorize this action with a signed-in user's access token. It must include the
// scope aws.cognito.signin.user.admin .
//
// Amazon Cognito doesn't evaluate Identity and Access Management (IAM) policies
// in requests for this API operation. For this operation, you can't use IAM
// credentials to authorize requests, and you can't grant IAM permissions in
// policies. For more information about authorization models in Amazon Cognito, see
// [Using the Amazon Cognito user pools API and user pool endpoints].
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
func cognitoidentityprovider_GetUserAuthFactors(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.GetUserAuthFactorsInput{
		// AccessToken: *string, // Required
	}

	if len(_cognitoidentityproviderAccessToken) > 0 {
		input.AccessToken = aws.String(_cognitoidentityproviderAccessToken)
	}

	if resp, err := client.GetUserAuthFactors(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Given a user pool ID, returns configuration for sign-in with WebAuthn
// authenticators and for multi-factor authentication (MFA). This operation
// describes the following:
//
// - The WebAuthn relying party (RP) ID and user-verification settings.
//
// - The required, optional, or disabled state of MFA for all user pool users.
//
// - The message templates for email and SMS MFA.
//
// - The enabled or disabled state of time-based one-time password (TOTP) MFA.
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func cognitoidentityprovider_GetUserPoolMfaConfig(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.GetUserPoolMfaConfigInput{
		// UserPoolId: *string, // Required
	}

	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}

	if resp, err := client.GetUserPoolMfaConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Invalidates the identity, access, and refresh tokens that Amazon Cognito issued
// to a user. Call this operation when your user signs out of your app. This
// results in the following behavior.
//
// - Amazon Cognito no longer accepts token-authorized user operations that you
// authorize with a signed-out user's access tokens. For more information, see [Using the Amazon Cognito user pools API and user pool endpoints].
//
// # Amazon Cognito returns an Access Token has been revoked error when your app
//
// attempts to authorize a user pools API request with a revoked access token that
// contains the scope aws.cognito.signin.user.admin .
//
// - Amazon Cognito no longer accepts a signed-out user's ID token in a [GetId]request
// to an identity pool with ServerSideTokenCheck enabled for its user pool IdP
// configuration in [CognitoIdentityProvider].
//
// - Amazon Cognito no longer accepts a signed-out user's refresh tokens in
// refresh requests.
//
// Other requests might be valid until your user's token expires. This operation
// doesn't clear the [managed login]session cookie. To clear the session for a user who signed in
// with managed login or the classic hosted UI, direct their browser session to the
// [logout endpoint].
//
// Authorize this action with a signed-in user's access token. It must include the
// scope aws.cognito.signin.user.admin .
//
// Amazon Cognito doesn't evaluate Identity and Access Management (IAM) policies
// in requests for this API operation. For this operation, you can't use IAM
// credentials to authorize requests, and you can't grant IAM permissions in
// policies. For more information about authorization models in Amazon Cognito, see
// [Using the Amazon Cognito user pools API and user pool endpoints].
//
// [logout endpoint]: https://docs.aws.amazon.com/cognito/latest/developerguide/logout-endpoint.html
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [managed login]: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-managed-login.html
// [CognitoIdentityProvider]: https://docs.aws.amazon.com/cognitoidentity/latest/APIReference/API_CognitoIdentityProvider.html
// [GetId]: https://docs.aws.amazon.com/cognitoidentity/latest/APIReference/API_GetId.html
func cognitoidentityprovider_GlobalSignOut(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.GlobalSignOutInput{
		// AccessToken: *string, // Required
	}

	if len(_cognitoidentityproviderAccessToken) > 0 {
		input.AccessToken = aws.String(_cognitoidentityproviderAccessToken)
	}

	if resp, err := client.GlobalSignOut(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Declares an authentication flow and initiates sign-in for a user in the Amazon
// Cognito user directory. Amazon Cognito might respond with an additional
// challenge or an AuthenticationResult that contains the outcome of a successful
// authentication. You can't sign in a user with a federated IdP with InitiateAuth
// . For more information, see [Authentication].
//
// Amazon Cognito doesn't evaluate Identity and Access Management (IAM) policies
// in requests for this API operation. For this operation, you can't use IAM
// credentials to authorize requests, and you can't grant IAM permissions in
// policies. For more information about authorization models in Amazon Cognito, see
// [Using the Amazon Cognito user pools API and user pool endpoints].
//
// This action might generate an SMS text message. Starting June 1, 2021, US
// telecom carriers require you to register an origination phone number before you
// can send SMS messages to US phone numbers. If you use SMS text messages in
// Amazon Cognito, you must register a phone number with [Amazon Pinpoint]. Amazon Cognito uses the
// registered number automatically. Otherwise, Amazon Cognito users who must
// receive SMS messages might not be able to sign up, activate their accounts, or
// sign in.
//
// If you have never used SMS text messages with Amazon Cognito or any other
// Amazon Web Services service, Amazon Simple Notification Service might place your
// account in the SMS sandbox. In [sandbox mode], you can send messages only to verified phone
// numbers. After you test your app while in the sandbox environment, you can move
// out of the sandbox and into production. For more information, see [SMS message settings for Amazon Cognito user pools]in the Amazon
// Cognito Developer Guide.
//
// [SMS message settings for Amazon Cognito user pools]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-sms-settings.html
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Authentication]: https://docs.aws.amazon.com/cognito/latest/developerguide/authentication.html
// [sandbox mode]: https://docs.aws.amazon.com/sns/latest/dg/sns-sms-sandbox.html
// [Amazon Pinpoint]: https://console.aws.amazon.com/pinpoint/home/
func cognitoidentityprovider_InitiateAuth(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.InitiateAuthInput{
		// AuthFlow: types.AuthFlowType, // Required
		// ClientId: *string, // Required
	}

	if len(_cognitoidentityproviderAuthFlow) > 0 {
		if err := assignInputField(input, "AuthFlow", _cognitoidentityproviderAuthFlow); err != nil {
			log.Errorf("invalid --auth-flow: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderClientId) > 0 {
		input.ClientId = aws.String(_cognitoidentityproviderClientId)
	}
	if len(_cognitoidentityproviderAnalyticsMetadata) > 0 {
		if err := assignInputField(input, "AnalyticsMetadata", _cognitoidentityproviderAnalyticsMetadata); err != nil {
			log.Errorf("invalid --analytics-metadata: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderAuthParameters) > 0 {
		if err := assignInputField(input, "AuthParameters", _cognitoidentityproviderAuthParameters); err != nil {
			log.Errorf("invalid --auth-parameters: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderClientMetadata) > 0 {
		if err := assignInputField(input, "ClientMetadata", _cognitoidentityproviderClientMetadata); err != nil {
			log.Errorf("invalid --client-metadata: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderSession) > 0 {
		input.Session = aws.String(_cognitoidentityproviderSession)
	}
	if len(_cognitoidentityproviderUserContextData) > 0 {
		if err := assignInputField(input, "UserContextData", _cognitoidentityproviderUserContextData); err != nil {
			log.Errorf("invalid --user-context-data: %s", err.Error())
			return
		}
	}

	if resp, err := client.InitiateAuth(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the devices that Amazon Cognito has registered to the currently signed-in
// user. For more information about device authentication, see [Working with user devices in your user pool].
//
// Authorize this action with a signed-in user's access token. It must include the
// scope aws.cognito.signin.user.admin .
//
// Amazon Cognito doesn't evaluate Identity and Access Management (IAM) policies
// in requests for this API operation. For this operation, you can't use IAM
// credentials to authorize requests, and you can't grant IAM permissions in
// policies. For more information about authorization models in Amazon Cognito, see
// [Using the Amazon Cognito user pools API and user pool endpoints].
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Working with user devices in your user pool]: https://docs.aws.amazon.com/cognito/latest/developerguide/amazon-cognito-user-pools-device-tracking.html
func cognitoidentityprovider_ListDevices(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.ListDevicesInput{
		// AccessToken: *string, // Required
	}

	if len(_cognitoidentityproviderAccessToken) > 0 {
		input.AccessToken = aws.String(_cognitoidentityproviderAccessToken)
	}
	if len(_cognitoidentityproviderLimit) > 0 {
		if err := assignInputField(input, "Limit", _cognitoidentityproviderLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderPaginationToken) > 0 {
		input.PaginationToken = aws.String(_cognitoidentityproviderPaginationToken)
	}

	if resp, err := client.ListDevices(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Given a user pool ID, returns user pool groups and their details.
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func cognitoidentityprovider_ListGroups(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.ListGroupsInput{
		// UserPoolId: *string, // Required
	}

	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}
	if len(_cognitoidentityproviderLimit) > 0 {
		if err := assignInputField(input, "Limit", _cognitoidentityproviderLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderNextToken) > 0 {
		input.NextToken = aws.String(_cognitoidentityproviderNextToken)
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

	var results []*cognitoidentityprovider.ListGroupsOutput
	p := cognitoidentityprovider.NewListGroupsPaginator(client, input)
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

// Given a user pool ID, returns information about configured identity providers
// (IdPs). For more information about IdPs, see [Third-party IdP sign-in].
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Third-party IdP sign-in]: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-identity-federation.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func cognitoidentityprovider_ListIdentityProviders(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.ListIdentityProvidersInput{
		// UserPoolId: *string, // Required
	}

	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}
	if len(_cognitoidentityproviderMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cognitoidentityproviderMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderNextToken) > 0 {
		input.NextToken = aws.String(_cognitoidentityproviderNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListIdentityProviders(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cognitoidentityprovider.ListIdentityProvidersOutput
	p := cognitoidentityprovider.NewListIdentityProvidersPaginator(client, input)
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

// Given a user pool ID, returns all resource servers and their details. For more
// information about resource servers, see [Access control with resource servers].
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Access control with resource servers]: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-define-resource-servers.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func cognitoidentityprovider_ListResourceServers(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.ListResourceServersInput{
		// UserPoolId: *string, // Required
	}

	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}
	if len(_cognitoidentityproviderMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cognitoidentityproviderMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderNextToken) > 0 {
		input.NextToken = aws.String(_cognitoidentityproviderNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListResourceServers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cognitoidentityprovider.ListResourceServersOutput
	p := cognitoidentityprovider.NewListResourceServersPaginator(client, input)
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

// Lists the tags that are assigned to an Amazon Cognito user pool. For more
// information, see [Tagging resources].
//
// [Tagging resources]: https://docs.aws.amazon.com/cognito/latest/developerguide/tagging.html
func cognitoidentityprovider_ListTagsForResource(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_cognitoidentityproviderResourceArn) > 0 {
		input.ResourceArn = aws.String(_cognitoidentityproviderResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns details about all terms documents for the requested user pool.
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func cognitoidentityprovider_ListTerms(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.ListTermsInput{
		// UserPoolId: *string, // Required
	}

	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}
	if len(_cognitoidentityproviderMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cognitoidentityproviderMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderNextToken) > 0 {
		input.NextToken = aws.String(_cognitoidentityproviderNextToken)
	}

	if resp, err := client.ListTerms(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Given a user pool ID, returns user import jobs and their details. Import jobs
// are retained in user pool configuration so that you can stage, stop, start,
// review, and delete them. For more information about user import, see [Importing users from a CSV file].
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Importing users from a CSV file]: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-using-import-tool.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func cognitoidentityprovider_ListUserImportJobs(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.ListUserImportJobsInput{
		// MaxResults: *int32, // Required
		// UserPoolId: *string, // Required
	}

	if len(_cognitoidentityproviderMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cognitoidentityproviderMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}
	if len(_cognitoidentityproviderPaginationToken) > 0 {
		input.PaginationToken = aws.String(_cognitoidentityproviderPaginationToken)
	}

	if resp, err := client.ListUserImportJobs(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all client secrets associated with a user pool app client. Returns
// metadata about the secrets. The response does not include pagination tokens as
// there are only 2 secrets at any given time and we return both with every
// ListUserPoolClientSecrets call. For security reasons, the response never reveals
// the actual secret value in ClientSecretValue.
func cognitoidentityprovider_ListUserPoolClientSecrets(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.ListUserPoolClientSecretsInput{
		// ClientId: *string, // Required
		// UserPoolId: *string, // Required
	}

	if len(_cognitoidentityproviderClientId) > 0 {
		input.ClientId = aws.String(_cognitoidentityproviderClientId)
	}
	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}
	if len(_cognitoidentityproviderNextToken) > 0 {
		input.NextToken = aws.String(_cognitoidentityproviderNextToken)
	}

	if resp, err := client.ListUserPoolClientSecrets(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Given a user pool ID, lists app clients. App clients are sets of rules for the
// access that you want a user pool to grant to one application. For more
// information, see [App clients].
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [App clients]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-client-apps.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func cognitoidentityprovider_ListUserPoolClients(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.ListUserPoolClientsInput{
		// UserPoolId: *string, // Required
	}

	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}
	if len(_cognitoidentityproviderMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cognitoidentityproviderMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderNextToken) > 0 {
		input.NextToken = aws.String(_cognitoidentityproviderNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListUserPoolClients(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cognitoidentityprovider.ListUserPoolClientsOutput
	p := cognitoidentityprovider.NewListUserPoolClientsPaginator(client, input)
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

// Lists user pools and their details in the current Amazon Web Services account.
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func cognitoidentityprovider_ListUserPools(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.ListUserPoolsInput{
		// MaxResults: *int32, // Required
	}

	if len(_cognitoidentityproviderMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cognitoidentityproviderMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderNextToken) > 0 {
		input.NextToken = aws.String(_cognitoidentityproviderNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListUserPools(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cognitoidentityprovider.ListUserPoolsOutput
	p := cognitoidentityprovider.NewListUserPoolsPaginator(client, input)
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

// Given a user pool ID, returns a list of users and their basic details in a user
// pool.
//
// This operation is eventually consistent. You might experience a delay before
// results are up-to-date. To validate the existence or configuration of an
// individual user, use AdminGetUser .
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func cognitoidentityprovider_ListUsers(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.ListUsersInput{
		// UserPoolId: *string, // Required
	}

	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}
	if len(_cognitoidentityproviderAttributesToGet) > 0 {
		input.AttributesToGet = append([]string(nil), _cognitoidentityproviderAttributesToGet...)
	}
	if len(_cognitoidentityproviderFilter) > 0 {
		input.Filter = aws.String(_cognitoidentityproviderFilter)
	}
	if len(_cognitoidentityproviderLimit) > 0 {
		if err := assignInputField(input, "Limit", _cognitoidentityproviderLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderPaginationToken) > 0 {
		input.PaginationToken = aws.String(_cognitoidentityproviderPaginationToken)
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

	var results []*cognitoidentityprovider.ListUsersOutput
	p := cognitoidentityprovider.NewListUsersPaginator(client, input)
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

// Given a user pool ID and a group name, returns a list of users in the group.
// For more information about user pool groups, see [Adding groups to a user pool].
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Adding groups to a user pool]: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-user-groups.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func cognitoidentityprovider_ListUsersInGroup(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.ListUsersInGroupInput{
		// GroupName: *string, // Required
		// UserPoolId: *string, // Required
	}

	if len(_cognitoidentityproviderGroupName) > 0 {
		input.GroupName = aws.String(_cognitoidentityproviderGroupName)
	}
	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}
	if len(_cognitoidentityproviderLimit) > 0 {
		if err := assignInputField(input, "Limit", _cognitoidentityproviderLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderNextToken) > 0 {
		input.NextToken = aws.String(_cognitoidentityproviderNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListUsersInGroup(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cognitoidentityprovider.ListUsersInGroupOutput
	p := cognitoidentityprovider.NewListUsersInGroupPaginator(client, input)
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

// Generates a list of the currently signed-in user's registered passkey, or
// WebAuthn, credentials.
//
// Authorize this action with a signed-in user's access token. It must include the
// scope aws.cognito.signin.user.admin .
//
// Amazon Cognito doesn't evaluate Identity and Access Management (IAM) policies
// in requests for this API operation. For this operation, you can't use IAM
// credentials to authorize requests, and you can't grant IAM permissions in
// policies. For more information about authorization models in Amazon Cognito, see
// [Using the Amazon Cognito user pools API and user pool endpoints].
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
func cognitoidentityprovider_ListWebAuthnCredentials(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.ListWebAuthnCredentialsInput{
		// AccessToken: *string, // Required
	}

	if len(_cognitoidentityproviderAccessToken) > 0 {
		input.AccessToken = aws.String(_cognitoidentityproviderAccessToken)
	}
	if len(_cognitoidentityproviderMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cognitoidentityproviderMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderNextToken) > 0 {
		input.NextToken = aws.String(_cognitoidentityproviderNextToken)
	}

	if resp, err := client.ListWebAuthnCredentials(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Resends the code that confirms a new account for a user who has signed up in
// your user pool. Amazon Cognito sends confirmation codes to the user attribute in
// the AutoVerifiedAttributes property of your user pool. When you prompt new
// users for the confirmation code, include a "Resend code" option that generates a
// call to this API operation.
//
// Amazon Cognito doesn't evaluate Identity and Access Management (IAM) policies
// in requests for this API operation. For this operation, you can't use IAM
// credentials to authorize requests, and you can't grant IAM permissions in
// policies. For more information about authorization models in Amazon Cognito, see
// [Using the Amazon Cognito user pools API and user pool endpoints].
//
// This action might generate an SMS text message. Starting June 1, 2021, US
// telecom carriers require you to register an origination phone number before you
// can send SMS messages to US phone numbers. If you use SMS text messages in
// Amazon Cognito, you must register a phone number with [Amazon Pinpoint]. Amazon Cognito uses the
// registered number automatically. Otherwise, Amazon Cognito users who must
// receive SMS messages might not be able to sign up, activate their accounts, or
// sign in.
//
// If you have never used SMS text messages with Amazon Cognito or any other
// Amazon Web Services service, Amazon Simple Notification Service might place your
// account in the SMS sandbox. In [sandbox mode], you can send messages only to verified phone
// numbers. After you test your app while in the sandbox environment, you can move
// out of the sandbox and into production. For more information, see [SMS message settings for Amazon Cognito user pools]in the Amazon
// Cognito Developer Guide.
//
// [SMS message settings for Amazon Cognito user pools]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-sms-settings.html
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [sandbox mode]: https://docs.aws.amazon.com/sns/latest/dg/sns-sms-sandbox.html
// [Amazon Pinpoint]: https://console.aws.amazon.com/pinpoint/home/
func cognitoidentityprovider_ResendConfirmationCode(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.ResendConfirmationCodeInput{
		// ClientId: *string, // Required
		// Username: *string, // Required
	}

	if len(_cognitoidentityproviderClientId) > 0 {
		input.ClientId = aws.String(_cognitoidentityproviderClientId)
	}
	if len(_cognitoidentityproviderUsername) > 0 {
		input.Username = aws.String(_cognitoidentityproviderUsername)
	}
	if len(_cognitoidentityproviderAnalyticsMetadata) > 0 {
		if err := assignInputField(input, "AnalyticsMetadata", _cognitoidentityproviderAnalyticsMetadata); err != nil {
			log.Errorf("invalid --analytics-metadata: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderClientMetadata) > 0 {
		if err := assignInputField(input, "ClientMetadata", _cognitoidentityproviderClientMetadata); err != nil {
			log.Errorf("invalid --client-metadata: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderSecretHash) > 0 {
		input.SecretHash = aws.String(_cognitoidentityproviderSecretHash)
	}
	if len(_cognitoidentityproviderUserContextData) > 0 {
		if err := assignInputField(input, "UserContextData", _cognitoidentityproviderUserContextData); err != nil {
			log.Errorf("invalid --user-context-data: %s", err.Error())
			return
		}
	}

	if resp, err := client.ResendConfirmationCode(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Some API operations in a user pool generate a challenge, like a prompt for an
// MFA code, for device authentication that bypasses MFA, or for a custom
// authentication challenge. A RespondToAuthChallenge API request provides the
// answer to that challenge, like a code or a secure remote password (SRP). The
// parameters of a response to an authentication challenge vary with the type of
// challenge.
//
// For more information about custom authentication challenges, see [Custom authentication challenge Lambda triggers].
//
// Amazon Cognito doesn't evaluate Identity and Access Management (IAM) policies
// in requests for this API operation. For this operation, you can't use IAM
// credentials to authorize requests, and you can't grant IAM permissions in
// policies. For more information about authorization models in Amazon Cognito, see
// [Using the Amazon Cognito user pools API and user pool endpoints].
//
// This action might generate an SMS text message. Starting June 1, 2021, US
// telecom carriers require you to register an origination phone number before you
// can send SMS messages to US phone numbers. If you use SMS text messages in
// Amazon Cognito, you must register a phone number with [Amazon Pinpoint]. Amazon Cognito uses the
// registered number automatically. Otherwise, Amazon Cognito users who must
// receive SMS messages might not be able to sign up, activate their accounts, or
// sign in.
//
// If you have never used SMS text messages with Amazon Cognito or any other
// Amazon Web Services service, Amazon Simple Notification Service might place your
// account in the SMS sandbox. In [sandbox mode], you can send messages only to verified phone
// numbers. After you test your app while in the sandbox environment, you can move
// out of the sandbox and into production. For more information, see [SMS message settings for Amazon Cognito user pools]in the Amazon
// Cognito Developer Guide.
//
// [SMS message settings for Amazon Cognito user pools]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-sms-settings.html
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [sandbox mode]: https://docs.aws.amazon.com/sns/latest/dg/sns-sms-sandbox.html
// [Custom authentication challenge Lambda triggers]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-lambda-challenge.html
// [Amazon Pinpoint]: https://console.aws.amazon.com/pinpoint/home/
func cognitoidentityprovider_RespondToAuthChallenge(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.RespondToAuthChallengeInput{
		// ChallengeName: types.ChallengeNameType, // Required
		// ClientId: *string, // Required
	}

	if len(_cognitoidentityproviderChallengeName) > 0 {
		if err := assignInputField(input, "ChallengeName", _cognitoidentityproviderChallengeName); err != nil {
			log.Errorf("invalid --challenge-name: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderClientId) > 0 {
		input.ClientId = aws.String(_cognitoidentityproviderClientId)
	}
	if len(_cognitoidentityproviderAnalyticsMetadata) > 0 {
		if err := assignInputField(input, "AnalyticsMetadata", _cognitoidentityproviderAnalyticsMetadata); err != nil {
			log.Errorf("invalid --analytics-metadata: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderChallengeResponses) > 0 {
		if err := assignInputField(input, "ChallengeResponses", _cognitoidentityproviderChallengeResponses); err != nil {
			log.Errorf("invalid --challenge-responses: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderClientMetadata) > 0 {
		if err := assignInputField(input, "ClientMetadata", _cognitoidentityproviderClientMetadata); err != nil {
			log.Errorf("invalid --client-metadata: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderSession) > 0 {
		input.Session = aws.String(_cognitoidentityproviderSession)
	}
	if len(_cognitoidentityproviderUserContextData) > 0 {
		if err := assignInputField(input, "UserContextData", _cognitoidentityproviderUserContextData); err != nil {
			log.Errorf("invalid --user-context-data: %s", err.Error())
			return
		}
	}

	if resp, err := client.RespondToAuthChallenge(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Revokes all of the access tokens generated by, and at the same time as, the
// specified refresh token. After a token is revoked, you can't use the revoked
// token to access Amazon Cognito user APIs, or to authorize access to your
// resource server.
//
// Amazon Cognito doesn't evaluate Identity and Access Management (IAM) policies
// in requests for this API operation. For this operation, you can't use IAM
// credentials to authorize requests, and you can't grant IAM permissions in
// policies. For more information about authorization models in Amazon Cognito, see
// [Using the Amazon Cognito user pools API and user pool endpoints].
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
func cognitoidentityprovider_RevokeToken(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.RevokeTokenInput{
		// ClientId: *string, // Required
		// Token: *string, // Required
	}

	if len(_cognitoidentityproviderClientId) > 0 {
		input.ClientId = aws.String(_cognitoidentityproviderClientId)
	}
	if len(_cognitoidentityproviderToken) > 0 {
		input.Token = aws.String(_cognitoidentityproviderToken)
	}
	if len(_cognitoidentityproviderClientSecret) > 0 {
		input.ClientSecret = aws.String(_cognitoidentityproviderClientSecret)
	}

	if resp, err := client.RevokeToken(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets up or modifies the logging configuration of a user pool. User pools can
// export user notification logs and, when threat protection is active,
// user-activity logs. For more information, see [Exporting user pool logs].
//
// [Exporting user pool logs]: https://docs.aws.amazon.com/cognito/latest/developerguide/exporting-quotas-and-usage.html
func cognitoidentityprovider_SetLogDeliveryConfiguration(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.SetLogDeliveryConfigurationInput{
		// LogConfigurations: []types.LogConfigurationType, // Required
		// UserPoolId: *string, // Required
	}

	if len(_cognitoidentityproviderLogConfigurations) > 0 {
		if err := assignInputField(input, "LogConfigurations", _cognitoidentityproviderLogConfigurations); err != nil {
			log.Errorf("invalid --log-configurations: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}

	if resp, err := client.SetLogDeliveryConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Configures threat protection for a user pool or app client. Sets configuration
// for the following.
//
// - Responses to risks with adaptive authentication
//
// - Responses to vulnerable passwords with compromised-credentials detection
//
// - Notifications to users who have had risky activity detected
//
// - IP-address denylist and allowlist
//
// To set the risk configuration for the user pool to defaults, send this request
// with only the UserPoolId parameter. To reset the threat protection settings of
// an app client to be inherited from the user pool, send UserPoolId and ClientId
// parameters only. To change threat protection to audit-only or off, update the
// value of UserPoolAddOns in an UpdateUserPool request. To activate this setting,
// your user pool must be on the [Plus tier].
//
// [Plus tier]: https://docs.aws.amazon.com/cognito/latest/developerguide/feature-plans-features-plus.html
func cognitoidentityprovider_SetRiskConfiguration(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.SetRiskConfigurationInput{
		// UserPoolId: *string, // Required
	}

	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}
	if len(_cognitoidentityproviderAccountTakeoverRiskConfiguration) > 0 {
		if err := assignInputField(input, "AccountTakeoverRiskConfiguration", _cognitoidentityproviderAccountTakeoverRiskConfiguration); err != nil {
			log.Errorf("invalid --account-takeover-risk-configuration: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderClientId) > 0 {
		input.ClientId = aws.String(_cognitoidentityproviderClientId)
	}
	if len(_cognitoidentityproviderCompromisedCredentialsRiskConfiguration) > 0 {
		if err := assignInputField(input, "CompromisedCredentialsRiskConfiguration", _cognitoidentityproviderCompromisedCredentialsRiskConfiguration); err != nil {
			log.Errorf("invalid --compromised-credentials-risk-configuration: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderRiskExceptionConfiguration) > 0 {
		if err := assignInputField(input, "RiskExceptionConfiguration", _cognitoidentityproviderRiskExceptionConfiguration); err != nil {
			log.Errorf("invalid --risk-exception-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.SetRiskConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Configures UI branding settings for domains with the hosted UI (classic)
// branding version. Your user pool must have a domain. Configure a domain with .
//
// Set the default configuration for all clients with a ClientId of ALL . When the
// ClientId value is an app client ID, the settings you pass in this request apply
// to that app client and override the default ALL configuration.
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func cognitoidentityprovider_SetUICustomization(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.SetUICustomizationInput{
		// UserPoolId: *string, // Required
	}

	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}
	if len(_cognitoidentityproviderCSS) > 0 {
		input.CSS = aws.String(_cognitoidentityproviderCSS)
	}
	if len(_cognitoidentityproviderClientId) > 0 {
		input.ClientId = aws.String(_cognitoidentityproviderClientId)
	}
	if len(_cognitoidentityproviderImageFile) > 0 {
		if err := assignInputField(input, "ImageFile", _cognitoidentityproviderImageFile); err != nil {
			log.Errorf("invalid --image-file: %s", err.Error())
			return
		}
	}

	if resp, err := client.SetUICustomization(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Set the user's multi-factor authentication (MFA) method preference, including
// which MFA factors are activated and if any are preferred. Only one factor can be
// set as preferred. The preferred MFA factor will be used to authenticate a user
// if multiple factors are activated. If multiple options are activated and no
// preference is set, a challenge to choose an MFA option will be returned during
// sign-in. If an MFA type is activated for a user, the user will be prompted for
// MFA during all sign-in attempts unless device tracking is turned on and the
// device has been trusted. If you want MFA to be applied selectively based on the
// assessed risk level of sign-in attempts, deactivate MFA for users and turn on
// Adaptive Authentication for the user pool.
//
// Authorize this action with a signed-in user's access token. It must include the
// scope aws.cognito.signin.user.admin .
//
// Amazon Cognito doesn't evaluate Identity and Access Management (IAM) policies
// in requests for this API operation. For this operation, you can't use IAM
// credentials to authorize requests, and you can't grant IAM permissions in
// policies. For more information about authorization models in Amazon Cognito, see
// [Using the Amazon Cognito user pools API and user pool endpoints].
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
func cognitoidentityprovider_SetUserMFAPreference(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.SetUserMFAPreferenceInput{
		// AccessToken: *string, // Required
	}

	if len(_cognitoidentityproviderAccessToken) > 0 {
		input.AccessToken = aws.String(_cognitoidentityproviderAccessToken)
	}
	if len(_cognitoidentityproviderEmailMfaSettings) > 0 {
		if err := assignInputField(input, "EmailMfaSettings", _cognitoidentityproviderEmailMfaSettings); err != nil {
			log.Errorf("invalid --email-mfa-settings: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderSMSMfaSettings) > 0 {
		if err := assignInputField(input, "SMSMfaSettings", _cognitoidentityproviderSMSMfaSettings); err != nil {
			log.Errorf("invalid --sms-mfa-settings: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderSoftwareTokenMfaSettings) > 0 {
		if err := assignInputField(input, "SoftwareTokenMfaSettings", _cognitoidentityproviderSoftwareTokenMfaSettings); err != nil {
			log.Errorf("invalid --software-token-mfa-settings: %s", err.Error())
			return
		}
	}

	if resp, err := client.SetUserMFAPreference(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets user pool multi-factor authentication (MFA) and passkey configuration. For
// more information about user pool MFA, see [Adding MFA]. For more information about WebAuthn
// passkeys see [Authentication flows].
//
// This action might generate an SMS text message. Starting June 1, 2021, US
// telecom carriers require you to register an origination phone number before you
// can send SMS messages to US phone numbers. If you use SMS text messages in
// Amazon Cognito, you must register a phone number with [Amazon Pinpoint]. Amazon Cognito uses the
// registered number automatically. Otherwise, Amazon Cognito users who must
// receive SMS messages might not be able to sign up, activate their accounts, or
// sign in.
//
// If you have never used SMS text messages with Amazon Cognito or any other
// Amazon Web Services service, Amazon Simple Notification Service might place your
// account in the SMS sandbox. In [sandbox mode], you can send messages only to verified phone
// numbers. After you test your app while in the sandbox environment, you can move
// out of the sandbox and into production. For more information, see [SMS message settings for Amazon Cognito user pools]in the Amazon
// Cognito Developer Guide.
//
// [SMS message settings for Amazon Cognito user pools]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-sms-settings.html
// [Adding MFA]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-mfa.html
// [sandbox mode]: https://docs.aws.amazon.com/sns/latest/dg/sns-sms-sandbox.html
// [Authentication flows]: https://docs.aws.amazon.com/cognito/latest/developerguide/amazon-cognito-user-pools-authentication-flow-methods.html#amazon-cognito-user-pools-authentication-flow-methods-passkey
// [Amazon Pinpoint]: https://console.aws.amazon.com/pinpoint/home/
func cognitoidentityprovider_SetUserPoolMfaConfig(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.SetUserPoolMfaConfigInput{
		// UserPoolId: *string, // Required
	}

	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}
	if len(_cognitoidentityproviderEmailMfaConfiguration) > 0 {
		if err := assignInputField(input, "EmailMfaConfiguration", _cognitoidentityproviderEmailMfaConfiguration); err != nil {
			log.Errorf("invalid --email-mfa-configuration: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderMfaConfiguration) > 0 {
		if err := assignInputField(input, "MfaConfiguration", _cognitoidentityproviderMfaConfiguration); err != nil {
			log.Errorf("invalid --mfa-configuration: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderSmsMfaConfiguration) > 0 {
		if err := assignInputField(input, "SmsMfaConfiguration", _cognitoidentityproviderSmsMfaConfiguration); err != nil {
			log.Errorf("invalid --sms-mfa-configuration: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderSoftwareTokenMfaConfiguration) > 0 {
		if err := assignInputField(input, "SoftwareTokenMfaConfiguration", _cognitoidentityproviderSoftwareTokenMfaConfiguration); err != nil {
			log.Errorf("invalid --software-token-mfa-configuration: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderWebAuthnConfiguration) > 0 {
		if err := assignInputField(input, "WebAuthnConfiguration", _cognitoidentityproviderWebAuthnConfiguration); err != nil {
			log.Errorf("invalid --web-authn-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.SetUserPoolMfaConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This action is no longer supported. You can use it to configure only SMS MFA.
// You can't use it to configure time-based one-time password (TOTP) software token
// or email MFA.
//
// Authorize this action with a signed-in user's access token. It must include the
// scope aws.cognito.signin.user.admin .
//
// Amazon Cognito doesn't evaluate Identity and Access Management (IAM) policies
// in requests for this API operation. For this operation, you can't use IAM
// credentials to authorize requests, and you can't grant IAM permissions in
// policies. For more information about authorization models in Amazon Cognito, see
// [Using the Amazon Cognito user pools API and user pool endpoints].
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
func cognitoidentityprovider_SetUserSettings(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.SetUserSettingsInput{
		// AccessToken: *string, // Required
		// MFAOptions: []types.MFAOptionType, // Required
	}

	if len(_cognitoidentityproviderAccessToken) > 0 {
		input.AccessToken = aws.String(_cognitoidentityproviderAccessToken)
	}
	if len(_cognitoidentityproviderMFAOptions) > 0 {
		if err := assignInputField(input, "MFAOptions", _cognitoidentityproviderMFAOptions); err != nil {
			log.Errorf("invalid --mfa-options: %s", err.Error())
			return
		}
	}

	if resp, err := client.SetUserSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Registers a user with an app client and requests a user name, password, and
// user attributes in the user pool.
//
// Amazon Cognito doesn't evaluate Identity and Access Management (IAM) policies
// in requests for this API operation. For this operation, you can't use IAM
// credentials to authorize requests, and you can't grant IAM permissions in
// policies. For more information about authorization models in Amazon Cognito, see
// [Using the Amazon Cognito user pools API and user pool endpoints].
//
// This action might generate an SMS text message. Starting June 1, 2021, US
// telecom carriers require you to register an origination phone number before you
// can send SMS messages to US phone numbers. If you use SMS text messages in
// Amazon Cognito, you must register a phone number with [Amazon Pinpoint]. Amazon Cognito uses the
// registered number automatically. Otherwise, Amazon Cognito users who must
// receive SMS messages might not be able to sign up, activate their accounts, or
// sign in.
//
// If you have never used SMS text messages with Amazon Cognito or any other
// Amazon Web Services service, Amazon Simple Notification Service might place your
// account in the SMS sandbox. In [sandbox mode], you can send messages only to verified phone
// numbers. After you test your app while in the sandbox environment, you can move
// out of the sandbox and into production. For more information, see [SMS message settings for Amazon Cognito user pools]in the Amazon
// Cognito Developer Guide.
//
// You might receive a LimitExceeded exception in response to this request if you
// have exceeded a rate quota for email or SMS messages, and if your user pool
// automatically verifies email addresses or phone numbers. When you get this
// exception in the response, the user is successfully created and is in an
// UNCONFIRMED state.
//
// [SMS message settings for Amazon Cognito user pools]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-sms-settings.html
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [sandbox mode]: https://docs.aws.amazon.com/sns/latest/dg/sns-sms-sandbox.html
// [Amazon Pinpoint]: https://console.aws.amazon.com/pinpoint/home/
func cognitoidentityprovider_SignUp(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.SignUpInput{
		// ClientId: *string, // Required
		// Username: *string, // Required
	}

	if len(_cognitoidentityproviderClientId) > 0 {
		input.ClientId = aws.String(_cognitoidentityproviderClientId)
	}
	if len(_cognitoidentityproviderUsername) > 0 {
		input.Username = aws.String(_cognitoidentityproviderUsername)
	}
	if len(_cognitoidentityproviderAnalyticsMetadata) > 0 {
		if err := assignInputField(input, "AnalyticsMetadata", _cognitoidentityproviderAnalyticsMetadata); err != nil {
			log.Errorf("invalid --analytics-metadata: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderClientMetadata) > 0 {
		if err := assignInputField(input, "ClientMetadata", _cognitoidentityproviderClientMetadata); err != nil {
			log.Errorf("invalid --client-metadata: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderPassword) > 0 {
		input.Password = aws.String(_cognitoidentityproviderPassword)
	}
	if len(_cognitoidentityproviderSecretHash) > 0 {
		input.SecretHash = aws.String(_cognitoidentityproviderSecretHash)
	}
	if len(_cognitoidentityproviderUserAttributes) > 0 {
		if err := assignInputField(input, "UserAttributes", _cognitoidentityproviderUserAttributes); err != nil {
			log.Errorf("invalid --user-attributes: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderUserContextData) > 0 {
		if err := assignInputField(input, "UserContextData", _cognitoidentityproviderUserContextData); err != nil {
			log.Errorf("invalid --user-context-data: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderValidationData) > 0 {
		if err := assignInputField(input, "ValidationData", _cognitoidentityproviderValidationData); err != nil {
			log.Errorf("invalid --validation-data: %s", err.Error())
			return
		}
	}

	if resp, err := client.SignUp(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Instructs your user pool to start importing users from a CSV file that contains
// their usernames and attributes. For more information about importing users from
// a CSV file, see [Importing users from a CSV file].
//
// [Importing users from a CSV file]: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-using-import-tool.html
func cognitoidentityprovider_StartUserImportJob(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.StartUserImportJobInput{
		// JobId: *string, // Required
		// UserPoolId: *string, // Required
	}

	if len(_cognitoidentityproviderJobId) > 0 {
		input.JobId = aws.String(_cognitoidentityproviderJobId)
	}
	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}

	if resp, err := client.StartUserImportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Requests credential creation options from your user pool for the currently
// signed-in user. Returns information about the user pool, the user profile, and
// authentication requirements. Users must provide this information in their
// request to enroll your application with their passkey provider.
//
// Authorize this action with a signed-in user's access token. It must include the
// scope aws.cognito.signin.user.admin .
func cognitoidentityprovider_StartWebAuthnRegistration(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.StartWebAuthnRegistrationInput{
		// AccessToken: *string, // Required
	}

	if len(_cognitoidentityproviderAccessToken) > 0 {
		input.AccessToken = aws.String(_cognitoidentityproviderAccessToken)
	}

	if resp, err := client.StartWebAuthnRegistration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Instructs your user pool to stop a running job that's importing users from a
// CSV file that contains their usernames and attributes. For more information
// about importing users from a CSV file, see [Importing users from a CSV file].
//
// [Importing users from a CSV file]: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-using-import-tool.html
func cognitoidentityprovider_StopUserImportJob(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.StopUserImportJobInput{
		// JobId: *string, // Required
		// UserPoolId: *string, // Required
	}

	if len(_cognitoidentityproviderJobId) > 0 {
		input.JobId = aws.String(_cognitoidentityproviderJobId)
	}
	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}

	if resp, err := client.StopUserImportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Assigns a set of tags to an Amazon Cognito user pool. A tag is a label that you
// can use to categorize and manage user pools in different ways, such as by
// purpose, owner, environment, or other criteria.
//
// Each tag consists of a key and value, both of which you define. A key is a
// general category for more specific values. For example, if you have two versions
// of a user pool, one for testing and another for production, you might assign an
// Environment tag key to both user pools. The value of this key might be Test for
// one user pool, and Production for the other.
//
// Tags are useful for cost tracking and access control. You can activate your
// tags so that they appear on the Billing and Cost Management console, where you
// can track the costs associated with your user pools. In an Identity and Access
// Management policy, you can constrain permissions for user pools based on
// specific tags or tag values.
//
// You can use this action up to 5 times per second, per account. A user pool can
// have as many as 50 tags.
func cognitoidentityprovider_TagResource(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_cognitoidentityproviderResourceArn) > 0 {
		input.ResourceArn = aws.String(_cognitoidentityproviderResourceArn)
	}
	if len(_cognitoidentityproviderTags) > 0 {
		if err := assignInputField(input, "Tags", _cognitoidentityproviderTags); err != nil {
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

// Given tag IDs that you previously assigned to a user pool, removes them.
func cognitoidentityprovider_UntagResource(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_cognitoidentityproviderResourceArn) > 0 {
		input.ResourceArn = aws.String(_cognitoidentityproviderResourceArn)
	}
	if len(_cognitoidentityproviderTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _cognitoidentityproviderTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides the feedback for an authentication event generated by threat
// protection features. The user's response indicates that you think that the event
// either was from a valid user or was an unwanted authentication attempt. This
// feedback improves the risk evaluation decision for the user pool as part of
// Amazon Cognito threat protection. To activate this setting, your user pool must
// be on the [Plus tier].
//
// This operation requires a FeedbackToken that Amazon Cognito generates and adds
// to notification emails when users have potentially suspicious authentication
// events. Users invoke this operation when they select the link that corresponds
// to {one-click-link-valid} or {one-click-link-invalid} in your notification
// template. Because FeedbackToken is a required parameter, you can't make
// requests to UpdateAuthEventFeedback without the contents of the notification
// email message.
//
// Amazon Cognito doesn't evaluate Identity and Access Management (IAM) policies
// in requests for this API operation. For this operation, you can't use IAM
// credentials to authorize requests, and you can't grant IAM permissions in
// policies. For more information about authorization models in Amazon Cognito, see
// [Using the Amazon Cognito user pools API and user pool endpoints].
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Plus tier]: https://docs.aws.amazon.com/cognito/latest/developerguide/feature-plans-features-plus.html
func cognitoidentityprovider_UpdateAuthEventFeedback(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.UpdateAuthEventFeedbackInput{
		// EventId: *string, // Required
		// FeedbackToken: *string, // Required
		// FeedbackValue: types.FeedbackValueType, // Required
		// UserPoolId: *string, // Required
		// Username: *string, // Required
	}

	if len(_cognitoidentityproviderEventId) > 0 {
		input.EventId = aws.String(_cognitoidentityproviderEventId)
	}
	if len(_cognitoidentityproviderFeedbackToken) > 0 {
		input.FeedbackToken = aws.String(_cognitoidentityproviderFeedbackToken)
	}
	if len(_cognitoidentityproviderFeedbackValue) > 0 {
		if err := assignInputField(input, "FeedbackValue", _cognitoidentityproviderFeedbackValue); err != nil {
			log.Errorf("invalid --feedback-value: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}
	if len(_cognitoidentityproviderUsername) > 0 {
		input.Username = aws.String(_cognitoidentityproviderUsername)
	}

	if resp, err := client.UpdateAuthEventFeedback(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the status of a the currently signed-in user's device so that it is
// marked as remembered or not remembered for the purpose of device authentication.
// Device authentication is a "remember me" mechanism that silently completes
// sign-in from trusted devices with a device key instead of a user-provided MFA
// code. This operation changes the status of a device without deleting it, so you
// can enable it again later. For more information about device authentication, see
// [Working with devices].
//
// Authorize this action with a signed-in user's access token. It must include the
// scope aws.cognito.signin.user.admin .
//
// Amazon Cognito doesn't evaluate Identity and Access Management (IAM) policies
// in requests for this API operation. For this operation, you can't use IAM
// credentials to authorize requests, and you can't grant IAM permissions in
// policies. For more information about authorization models in Amazon Cognito, see
// [Using the Amazon Cognito user pools API and user pool endpoints].
//
// [Working with devices]: https://docs.aws.amazon.com/cognito/latest/developerguide/amazon-cognito-user-pools-device-tracking.html
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
func cognitoidentityprovider_UpdateDeviceStatus(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.UpdateDeviceStatusInput{
		// AccessToken: *string, // Required
		// DeviceKey: *string, // Required
	}

	if len(_cognitoidentityproviderAccessToken) > 0 {
		input.AccessToken = aws.String(_cognitoidentityproviderAccessToken)
	}
	if len(_cognitoidentityproviderDeviceKey) > 0 {
		input.DeviceKey = aws.String(_cognitoidentityproviderDeviceKey)
	}
	if len(_cognitoidentityproviderDeviceRememberedStatus) > 0 {
		if err := assignInputField(input, "DeviceRememberedStatus", _cognitoidentityproviderDeviceRememberedStatus); err != nil {
			log.Errorf("invalid --device-remembered-status: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateDeviceStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Given the name of a user pool group, updates any of the properties for
// precedence, IAM role, or description. For more information about user pool
// groups, see [Adding groups to a user pool].
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Adding groups to a user pool]: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-user-groups.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func cognitoidentityprovider_UpdateGroup(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.UpdateGroupInput{
		// GroupName: *string, // Required
		// UserPoolId: *string, // Required
	}

	if len(_cognitoidentityproviderGroupName) > 0 {
		input.GroupName = aws.String(_cognitoidentityproviderGroupName)
	}
	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}
	if len(_cognitoidentityproviderDescription) > 0 {
		input.Description = aws.String(_cognitoidentityproviderDescription)
	}
	if len(_cognitoidentityproviderPrecedence) > 0 {
		if err := assignInputField(input, "Precedence", _cognitoidentityproviderPrecedence); err != nil {
			log.Errorf("invalid --precedence: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderRoleArn) > 0 {
		input.RoleArn = aws.String(_cognitoidentityproviderRoleArn)
	}

	if resp, err := client.UpdateGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the configuration and trust relationship between a third-party
// identity provider (IdP) and a user pool. Amazon Cognito accepts sign-in with
// third-party identity providers through managed login and OIDC relying-party
// libraries. For more information, see [Third-party IdP sign-in].
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Third-party IdP sign-in]: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-identity-federation.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func cognitoidentityprovider_UpdateIdentityProvider(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.UpdateIdentityProviderInput{
		// ProviderName: *string, // Required
		// UserPoolId: *string, // Required
	}

	if len(_cognitoidentityproviderProviderName) > 0 {
		input.ProviderName = aws.String(_cognitoidentityproviderProviderName)
	}
	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}
	if len(_cognitoidentityproviderAttributeMapping) > 0 {
		if err := assignInputField(input, "AttributeMapping", _cognitoidentityproviderAttributeMapping); err != nil {
			log.Errorf("invalid --attribute-mapping: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderIdpIdentifiers) > 0 {
		input.IdpIdentifiers = append([]string(nil), _cognitoidentityproviderIdpIdentifiers...)
	}
	if len(_cognitoidentityproviderProviderDetails) > 0 {
		if err := assignInputField(input, "ProviderDetails", _cognitoidentityproviderProviderDetails); err != nil {
			log.Errorf("invalid --provider-details: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateIdentityProvider(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Configures the branding settings for a user pool style. This operation is the
// programmatic option for the configuration of a style in the branding editor.
//
// Provides values for UI customization in a Settings JSON object and image files
// in an Assets array.
//
// This operation has a 2-megabyte request-size limit and include the CSS settings
// and image assets for your app client. Your branding settings might exceed 2MB in
// size. Amazon Cognito doesn't require that you pass all parameters in one request
// and preserves existing style settings that you don't specify. If your request is
// larger than 2MB, separate it into multiple requests, each with a size smaller
// than the limit.
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func cognitoidentityprovider_UpdateManagedLoginBranding(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.UpdateManagedLoginBrandingInput{}

	if len(_cognitoidentityproviderAssets) > 0 {
		if err := assignInputField(input, "Assets", _cognitoidentityproviderAssets); err != nil {
			log.Errorf("invalid --assets: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderManagedLoginBrandingId) > 0 {
		input.ManagedLoginBrandingId = aws.String(_cognitoidentityproviderManagedLoginBrandingId)
	}
	if len(_cognitoidentityproviderSettings) > 0 {
		if err := assignInputField(input, "Settings", _cognitoidentityproviderSettings); err != nil {
			log.Errorf("invalid --settings: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderUseCognitoProvidedValues) > 0 {
		if err := assignInputField(input, "UseCognitoProvidedValues", _cognitoidentityproviderUseCognitoProvidedValues); err != nil {
			log.Errorf("invalid --use-cognito-provided-values: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}

	if resp, err := client.UpdateManagedLoginBranding(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the name and scopes of a resource server. All other fields are
// read-only. For more information about resource servers, see [Access control with resource servers].
//
// If you don't provide a value for an attribute, it is set to the default value.
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Access control with resource servers]: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-define-resource-servers.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func cognitoidentityprovider_UpdateResourceServer(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.UpdateResourceServerInput{
		// Identifier: *string, // Required
		// Name: *string, // Required
		// UserPoolId: *string, // Required
	}

	if len(_cognitoidentityproviderIdentifier) > 0 {
		input.Identifier = aws.String(_cognitoidentityproviderIdentifier)
	}
	if len(_cognitoidentityproviderName) > 0 {
		input.Name = aws.String(_cognitoidentityproviderName)
	}
	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}
	if len(_cognitoidentityproviderScopes) > 0 {
		if err := assignInputField(input, "Scopes", _cognitoidentityproviderScopes); err != nil {
			log.Errorf("invalid --scopes: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateResourceServer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies existing terms documents for the requested app client. When Terms and
// conditions and Privacy policy documents are configured, the app client displays
// links to them in the sign-up page of managed login for the app client.
//
// You can provide URLs for terms documents in the languages that are supported by [managed login localization]
// . Amazon Cognito directs users to the terms documents for their current
// language, with fallback to default if no document exists for the language.
//
// Each request accepts one type of terms document and a map of language-to-link
// for that document type. You must provide both types of terms documents in at
// least one language before Amazon Cognito displays your terms documents. Supply
// each type in separate requests.
//
// For more information, see [Terms documents].
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Terms documents]: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-managed-login.html#managed-login-terms-documents
// [managed login localization]: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-managed-login.html#managed-login-localization
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func cognitoidentityprovider_UpdateTerms(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.UpdateTermsInput{
		// TermsId: *string, // Required
		// UserPoolId: *string, // Required
	}

	if len(_cognitoidentityproviderTermsId) > 0 {
		input.TermsId = aws.String(_cognitoidentityproviderTermsId)
	}
	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}
	if len(_cognitoidentityproviderEnforcement) > 0 {
		if err := assignInputField(input, "Enforcement", _cognitoidentityproviderEnforcement); err != nil {
			log.Errorf("invalid --enforcement: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderLinks) > 0 {
		if err := assignInputField(input, "Links", _cognitoidentityproviderLinks); err != nil {
			log.Errorf("invalid --links: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderTermsName) > 0 {
		input.TermsName = aws.String(_cognitoidentityproviderTermsName)
	}
	if len(_cognitoidentityproviderTermsSource) > 0 {
		if err := assignInputField(input, "TermsSource", _cognitoidentityproviderTermsSource); err != nil {
			log.Errorf("invalid --terms-source: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateTerms(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the currently signed-in user's attributes. To delete an attribute from
// the user, submit the attribute in your API request with a blank value.
//
// For custom attributes, you must add a custom: prefix to the attribute name, for
// example custom:department .
//
// Authorize this action with a signed-in user's access token. It must include the
// scope aws.cognito.signin.user.admin .
//
// Amazon Cognito doesn't evaluate Identity and Access Management (IAM) policies
// in requests for this API operation. For this operation, you can't use IAM
// credentials to authorize requests, and you can't grant IAM permissions in
// policies. For more information about authorization models in Amazon Cognito, see
// [Using the Amazon Cognito user pools API and user pool endpoints].
//
// This action might generate an SMS text message. Starting June 1, 2021, US
// telecom carriers require you to register an origination phone number before you
// can send SMS messages to US phone numbers. If you use SMS text messages in
// Amazon Cognito, you must register a phone number with [Amazon Pinpoint]. Amazon Cognito uses the
// registered number automatically. Otherwise, Amazon Cognito users who must
// receive SMS messages might not be able to sign up, activate their accounts, or
// sign in.
//
// If you have never used SMS text messages with Amazon Cognito or any other
// Amazon Web Services service, Amazon Simple Notification Service might place your
// account in the SMS sandbox. In [sandbox mode], you can send messages only to verified phone
// numbers. After you test your app while in the sandbox environment, you can move
// out of the sandbox and into production. For more information, see [SMS message settings for Amazon Cognito user pools]in the Amazon
// Cognito Developer Guide.
//
// [SMS message settings for Amazon Cognito user pools]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-sms-settings.html
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [sandbox mode]: https://docs.aws.amazon.com/sns/latest/dg/sns-sms-sandbox.html
// [Amazon Pinpoint]: https://console.aws.amazon.com/pinpoint/home/
func cognitoidentityprovider_UpdateUserAttributes(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.UpdateUserAttributesInput{
		// AccessToken: *string, // Required
		// UserAttributes: []types.AttributeType, // Required
	}

	if len(_cognitoidentityproviderAccessToken) > 0 {
		input.AccessToken = aws.String(_cognitoidentityproviderAccessToken)
	}
	if len(_cognitoidentityproviderUserAttributes) > 0 {
		if err := assignInputField(input, "UserAttributes", _cognitoidentityproviderUserAttributes); err != nil {
			log.Errorf("invalid --user-attributes: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderClientMetadata) > 0 {
		if err := assignInputField(input, "ClientMetadata", _cognitoidentityproviderClientMetadata); err != nil {
			log.Errorf("invalid --client-metadata: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateUserAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration of a user pool. To avoid setting parameters to Amazon
// Cognito defaults, construct this API request to pass the existing configuration
// of your user pool, modified to include the changes that you want to make.
//
// With the exception of UserPoolTier , if you don't provide a value for an
// attribute, Amazon Cognito sets it to its default value.
//
// This action might generate an SMS text message. Starting June 1, 2021, US
// telecom carriers require you to register an origination phone number before you
// can send SMS messages to US phone numbers. If you use SMS text messages in
// Amazon Cognito, you must register a phone number with [Amazon Pinpoint]. Amazon Cognito uses the
// registered number automatically. Otherwise, Amazon Cognito users who must
// receive SMS messages might not be able to sign up, activate their accounts, or
// sign in.
//
// If you have never used SMS text messages with Amazon Cognito or any other
// Amazon Web Services service, Amazon Simple Notification Service might place your
// account in the SMS sandbox. In [sandbox mode], you can send messages only to verified phone
// numbers. After you test your app while in the sandbox environment, you can move
// out of the sandbox and into production. For more information, see [SMS message settings for Amazon Cognito user pools]in the Amazon
// Cognito Developer Guide.
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [SMS message settings for Amazon Cognito user pools]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-sms-settings.html
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [sandbox mode]: https://docs.aws.amazon.com/sns/latest/dg/sns-sms-sandbox.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
// [Amazon Pinpoint]: https://console.aws.amazon.com/pinpoint/home/
func cognitoidentityprovider_UpdateUserPool(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.UpdateUserPoolInput{
		// UserPoolId: *string, // Required
	}

	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}
	if len(_cognitoidentityproviderAccountRecoverySetting) > 0 {
		if err := assignInputField(input, "AccountRecoverySetting", _cognitoidentityproviderAccountRecoverySetting); err != nil {
			log.Errorf("invalid --account-recovery-setting: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderAdminCreateUserConfig) > 0 {
		if err := assignInputField(input, "AdminCreateUserConfig", _cognitoidentityproviderAdminCreateUserConfig); err != nil {
			log.Errorf("invalid --admin-create-user-config: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderAutoVerifiedAttributes) > 0 {
		if err := assignInputField(input, "AutoVerifiedAttributes", _cognitoidentityproviderAutoVerifiedAttributes); err != nil {
			log.Errorf("invalid --auto-verified-attributes: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderDeletionProtection) > 0 {
		if err := assignInputField(input, "DeletionProtection", _cognitoidentityproviderDeletionProtection); err != nil {
			log.Errorf("invalid --deletion-protection: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderDeviceConfiguration) > 0 {
		if err := assignInputField(input, "DeviceConfiguration", _cognitoidentityproviderDeviceConfiguration); err != nil {
			log.Errorf("invalid --device-configuration: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderEmailConfiguration) > 0 {
		if err := assignInputField(input, "EmailConfiguration", _cognitoidentityproviderEmailConfiguration); err != nil {
			log.Errorf("invalid --email-configuration: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderEmailVerificationMessage) > 0 {
		input.EmailVerificationMessage = aws.String(_cognitoidentityproviderEmailVerificationMessage)
	}
	if len(_cognitoidentityproviderEmailVerificationSubject) > 0 {
		input.EmailVerificationSubject = aws.String(_cognitoidentityproviderEmailVerificationSubject)
	}
	if len(_cognitoidentityproviderLambdaConfig) > 0 {
		if err := assignInputField(input, "LambdaConfig", _cognitoidentityproviderLambdaConfig); err != nil {
			log.Errorf("invalid --lambda-config: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderMfaConfiguration) > 0 {
		if err := assignInputField(input, "MfaConfiguration", _cognitoidentityproviderMfaConfiguration); err != nil {
			log.Errorf("invalid --mfa-configuration: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderPolicies) > 0 {
		if err := assignInputField(input, "Policies", _cognitoidentityproviderPolicies); err != nil {
			log.Errorf("invalid --policies: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderPoolName) > 0 {
		input.PoolName = aws.String(_cognitoidentityproviderPoolName)
	}
	if len(_cognitoidentityproviderSmsAuthenticationMessage) > 0 {
		input.SmsAuthenticationMessage = aws.String(_cognitoidentityproviderSmsAuthenticationMessage)
	}
	if len(_cognitoidentityproviderSmsConfiguration) > 0 {
		if err := assignInputField(input, "SmsConfiguration", _cognitoidentityproviderSmsConfiguration); err != nil {
			log.Errorf("invalid --sms-configuration: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderSmsVerificationMessage) > 0 {
		input.SmsVerificationMessage = aws.String(_cognitoidentityproviderSmsVerificationMessage)
	}
	if len(_cognitoidentityproviderUserAttributeUpdateSettings) > 0 {
		if err := assignInputField(input, "UserAttributeUpdateSettings", _cognitoidentityproviderUserAttributeUpdateSettings); err != nil {
			log.Errorf("invalid --user-attribute-update-settings: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderUserPoolAddOns) > 0 {
		if err := assignInputField(input, "UserPoolAddOns", _cognitoidentityproviderUserPoolAddOns); err != nil {
			log.Errorf("invalid --user-pool-add-ons: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderUserPoolTags) > 0 {
		if err := assignInputField(input, "UserPoolTags", _cognitoidentityproviderUserPoolTags); err != nil {
			log.Errorf("invalid --user-pool-tags: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderUserPoolTier) > 0 {
		if err := assignInputField(input, "UserPoolTier", _cognitoidentityproviderUserPoolTier); err != nil {
			log.Errorf("invalid --user-pool-tier: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderVerificationMessageTemplate) > 0 {
		if err := assignInputField(input, "VerificationMessageTemplate", _cognitoidentityproviderVerificationMessageTemplate); err != nil {
			log.Errorf("invalid --verification-message-template: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateUserPool(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Given a user pool app client ID, updates the configuration. To avoid setting
// parameters to Amazon Cognito defaults, construct this API request to pass the
// existing configuration of your app client, modified to include the changes that
// you want to make.
//
// If you don't provide a value for an attribute, Amazon Cognito sets it to its
// default value.
//
// Unlike app clients created in the console, Amazon Cognito doesn't automatically
// assign a branding style to app clients that you configure with this API
// operation. Managed login and classic hosted UI pages aren't available for your
// client until after you apply a branding style.
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func cognitoidentityprovider_UpdateUserPoolClient(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.UpdateUserPoolClientInput{
		// ClientId: *string, // Required
		// UserPoolId: *string, // Required
	}

	if len(_cognitoidentityproviderClientId) > 0 {
		input.ClientId = aws.String(_cognitoidentityproviderClientId)
	}
	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}
	if len(_cognitoidentityproviderAccessTokenValidity) > 0 {
		if err := assignInputField(input, "AccessTokenValidity", _cognitoidentityproviderAccessTokenValidity); err != nil {
			log.Errorf("invalid --access-token-validity: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderAllowedOAuthFlows) > 0 {
		if err := assignInputField(input, "AllowedOAuthFlows", _cognitoidentityproviderAllowedOAuthFlows); err != nil {
			log.Errorf("invalid --allowed-oauth-flows: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderAllowedOAuthFlowsUserPoolClient) > 0 {
		if err := assignInputField(input, "AllowedOAuthFlowsUserPoolClient", _cognitoidentityproviderAllowedOAuthFlowsUserPoolClient); err != nil {
			log.Errorf("invalid --allowed-oauth-flows-user-pool-client: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderAllowedOAuthScopes) > 0 {
		input.AllowedOAuthScopes = append([]string(nil), _cognitoidentityproviderAllowedOAuthScopes...)
	}
	if len(_cognitoidentityproviderAnalyticsConfiguration) > 0 {
		if err := assignInputField(input, "AnalyticsConfiguration", _cognitoidentityproviderAnalyticsConfiguration); err != nil {
			log.Errorf("invalid --analytics-configuration: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderAuthSessionValidity) > 0 {
		if err := assignInputField(input, "AuthSessionValidity", _cognitoidentityproviderAuthSessionValidity); err != nil {
			log.Errorf("invalid --auth-session-validity: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderCallbackURLs) > 0 {
		input.CallbackURLs = append([]string(nil), _cognitoidentityproviderCallbackURLs...)
	}
	if len(_cognitoidentityproviderClientName) > 0 {
		input.ClientName = aws.String(_cognitoidentityproviderClientName)
	}
	if len(_cognitoidentityproviderDefaultRedirectURI) > 0 {
		input.DefaultRedirectURI = aws.String(_cognitoidentityproviderDefaultRedirectURI)
	}
	if len(_cognitoidentityproviderEnablePropagateAdditionalUserContextData) > 0 {
		if err := assignInputField(input, "EnablePropagateAdditionalUserContextData", _cognitoidentityproviderEnablePropagateAdditionalUserContextData); err != nil {
			log.Errorf("invalid --enable-propagate-additional-user-context-data: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderEnableTokenRevocation) > 0 {
		if err := assignInputField(input, "EnableTokenRevocation", _cognitoidentityproviderEnableTokenRevocation); err != nil {
			log.Errorf("invalid --enable-token-revocation: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderExplicitAuthFlows) > 0 {
		if err := assignInputField(input, "ExplicitAuthFlows", _cognitoidentityproviderExplicitAuthFlows); err != nil {
			log.Errorf("invalid --explicit-auth-flows: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderIdTokenValidity) > 0 {
		if err := assignInputField(input, "IdTokenValidity", _cognitoidentityproviderIdTokenValidity); err != nil {
			log.Errorf("invalid --id-token-validity: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderLogoutURLs) > 0 {
		input.LogoutURLs = append([]string(nil), _cognitoidentityproviderLogoutURLs...)
	}
	if len(_cognitoidentityproviderPreventUserExistenceErrors) > 0 {
		if err := assignInputField(input, "PreventUserExistenceErrors", _cognitoidentityproviderPreventUserExistenceErrors); err != nil {
			log.Errorf("invalid --prevent-user-existence-errors: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderReadAttributes) > 0 {
		input.ReadAttributes = append([]string(nil), _cognitoidentityproviderReadAttributes...)
	}
	if len(_cognitoidentityproviderRefreshTokenRotation) > 0 {
		if err := assignInputField(input, "RefreshTokenRotation", _cognitoidentityproviderRefreshTokenRotation); err != nil {
			log.Errorf("invalid --refresh-token-rotation: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderRefreshTokenValidity) > 0 {
		if err := assignInputField(input, "RefreshTokenValidity", _cognitoidentityproviderRefreshTokenValidity); err != nil {
			log.Errorf("invalid --refresh-token-validity: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderSupportedIdentityProviders) > 0 {
		input.SupportedIdentityProviders = append([]string(nil), _cognitoidentityproviderSupportedIdentityProviders...)
	}
	if len(_cognitoidentityproviderTokenValidityUnits) > 0 {
		if err := assignInputField(input, "TokenValidityUnits", _cognitoidentityproviderTokenValidityUnits); err != nil {
			log.Errorf("invalid --token-validity-units: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderWriteAttributes) > 0 {
		input.WriteAttributes = append([]string(nil), _cognitoidentityproviderWriteAttributes...)
	}

	if resp, err := client.UpdateUserPoolClient(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// A user pool domain hosts managed login, an authorization server and web server
// for authentication in your application. This operation updates the branding
// version for user pool domains between 1 for hosted UI (classic) and 2 for
// managed login. It also updates the SSL certificate for user pool custom domains.
//
// Changes to the domain branding version take up to one minute to take effect for
// a prefix domain and up to five minutes for a custom domain.
//
// This operation doesn't change the name of your user pool domain. To change your
// domain, delete it with DeleteUserPoolDomain and create a new domain with
// CreateUserPoolDomain .
//
// You can pass the ARN of a new Certificate Manager certificate in this request.
// Typically, ACM certificates automatically renew and you user pool can continue
// to use the same ARN. But if you generate a new certificate for your custom
// domain name, replace the original configuration with the new ARN in this
// request.
//
// ACM certificates for custom domains must be in the US East (N. Virginia) Amazon
// Web Services Region. After you submit your request, Amazon Cognito requires up
// to 1 hour to distribute your new certificate to your custom domain.
//
// For more information about adding a custom domain to your user pool, see [Configuring a user pool domain].
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Configuring a user pool domain]: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-add-custom-domain.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func cognitoidentityprovider_UpdateUserPoolDomain(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.UpdateUserPoolDomainInput{
		// Domain: *string, // Required
		// UserPoolId: *string, // Required
	}

	if len(_cognitoidentityproviderDomain) > 0 {
		input.Domain = aws.String(_cognitoidentityproviderDomain)
	}
	if len(_cognitoidentityproviderUserPoolId) > 0 {
		input.UserPoolId = aws.String(_cognitoidentityproviderUserPoolId)
	}
	if len(_cognitoidentityproviderCustomDomainConfig) > 0 {
		if err := assignInputField(input, "CustomDomainConfig", _cognitoidentityproviderCustomDomainConfig); err != nil {
			log.Errorf("invalid --custom-domain-config: %s", err.Error())
			return
		}
	}
	if len(_cognitoidentityproviderManagedLoginVersion) > 0 {
		if err := assignInputField(input, "ManagedLoginVersion", _cognitoidentityproviderManagedLoginVersion); err != nil {
			log.Errorf("invalid --managed-login-version: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateUserPoolDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Registers the current user's time-based one-time password (TOTP) authenticator
// with a code generated in their authenticator app from a private key that's
// supplied by your user pool. Marks the user's software token MFA status as
// "verified" if successful. The request takes an access token or a session string,
// but not both.
//
// Amazon Cognito doesn't evaluate Identity and Access Management (IAM) policies
// in requests for this API operation. For this operation, you can't use IAM
// credentials to authorize requests, and you can't grant IAM permissions in
// policies. For more information about authorization models in Amazon Cognito, see
// [Using the Amazon Cognito user pools API and user pool endpoints].
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
func cognitoidentityprovider_VerifySoftwareToken(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.VerifySoftwareTokenInput{
		// UserCode: *string, // Required
	}

	if len(_cognitoidentityproviderUserCode) > 0 {
		input.UserCode = aws.String(_cognitoidentityproviderUserCode)
	}
	if len(_cognitoidentityproviderAccessToken) > 0 {
		input.AccessToken = aws.String(_cognitoidentityproviderAccessToken)
	}
	if len(_cognitoidentityproviderFriendlyDeviceName) > 0 {
		input.FriendlyDeviceName = aws.String(_cognitoidentityproviderFriendlyDeviceName)
	}
	if len(_cognitoidentityproviderSession) > 0 {
		input.Session = aws.String(_cognitoidentityproviderSession)
	}

	if resp, err := client.VerifySoftwareToken(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Submits a verification code for a signed-in user who has added or changed a
// value of an auto-verified attribute. When successful, the user's attribute
// becomes verified and the attribute email_verified or phone_number_verified
// becomes true .
//
// If your user pool requires verification before Amazon Cognito updates the
// attribute value, this operation updates the affected attribute to its pending
// value.
//
// Authorize this action with a signed-in user's access token. It must include the
// scope aws.cognito.signin.user.admin .
//
// Amazon Cognito doesn't evaluate Identity and Access Management (IAM) policies
// in requests for this API operation. For this operation, you can't use IAM
// credentials to authorize requests, and you can't grant IAM permissions in
// policies. For more information about authorization models in Amazon Cognito, see
// [Using the Amazon Cognito user pools API and user pool endpoints].
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
func cognitoidentityprovider_VerifyUserAttribute(cfg aws.Config, client *cognitoidentityprovider.Client) {
	input := &cognitoidentityprovider.VerifyUserAttributeInput{
		// AccessToken: *string, // Required
		// AttributeName: *string, // Required
		// Code: *string, // Required
	}

	if len(_cognitoidentityproviderAccessToken) > 0 {
		input.AccessToken = aws.String(_cognitoidentityproviderAccessToken)
	}
	if len(_cognitoidentityproviderAttributeName) > 0 {
		input.AttributeName = aws.String(_cognitoidentityproviderAttributeName)
	}
	if len(_cognitoidentityproviderCode) > 0 {
		input.Code = aws.String(_cognitoidentityproviderCode)
	}

	if resp, err := client.VerifyUserAttribute(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_cognitoidentityproviderCmd)
	_cognitoidentityproviderCmd.Flags().SortFlags = false

	_cognitoidentityproviderCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_cognitoidentityproviderCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderAccessToken, "access-token", "", "", "Access Token")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderAccessTokenValidity, "access-token-validity", "", "", "Access Token Validity")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderAccountRecoverySetting, "account-recovery-setting", "", "", "Account Recovery Setting")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderAccountTakeoverRiskConfiguration, "account-takeover-risk-configuration", "", "", "Account Takeover Risk Configuration")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderAdminCreateUserConfig, "admin-create-user-config", "", "", "Admin Create User Config")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderAliasAttributes, "alias-attributes", "", "", "Alias Attributes")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderAllowedOAuthFlows, "allowed-oauth-flows", "", "", "Allowed Oauth Flows")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderAllowedOAuthFlowsUserPoolClient, "allowed-oauth-flows-user-pool-client", "", "", "Allowed Oauth Flows User Pool Client")
	_cognitoidentityproviderCmd.Flags().StringSliceVarP(&_cognitoidentityproviderAllowedOAuthScopes, "allowed-oauth-scopes", "", nil, "Allowed Oauth Scopes")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderAnalyticsConfiguration, "analytics-configuration", "", "", "Analytics Configuration")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderAnalyticsMetadata, "analytics-metadata", "", "", "Analytics Metadata")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderAssets, "assets", "", "", "Assets")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderAttributeMapping, "attribute-mapping", "", "", "Attribute Mapping")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderAttributeName, "attribute-name", "", "", "Attribute Name")
	_cognitoidentityproviderCmd.Flags().StringSliceVarP(&_cognitoidentityproviderAttributesToGet, "attributes-to-get", "", nil, "Attributes To Get")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderAuthFlow, "auth-flow", "", "", "Auth Flow")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderAuthParameters, "auth-parameters", "", "", "Auth Parameters")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderAuthSessionValidity, "auth-session-validity", "", "", "Auth Session Validity")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderAutoVerifiedAttributes, "auto-verified-attributes", "", "", "Auto Verified Attributes")
	_cognitoidentityproviderCmd.Flags().StringSliceVarP(&_cognitoidentityproviderCallbackURLs, "callback-urls", "", nil, "Callback Urls")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderChallengeName, "challenge-name", "", "", "Challenge Name")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderChallengeResponses, "challenge-responses", "", "", "Challenge Responses")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderClientId, "client-id", "", "", "Client ID")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderClientMetadata, "client-metadata", "", "", "Client Metadata")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderClientName, "client-name", "", "", "Client Name")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderClientSecret, "client-secret", "", "", "Client Secret")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderClientSecretId, "client-secret-id", "", "", "Client Secret ID")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderCloudWatchLogsRoleArn, "cloud-watch-logs-role-arn", "", "", "Cloud Watch Logs Role ARN")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderCode, "code", "", "", "Code")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderCompromisedCredentialsRiskConfiguration, "compromised-credentials-risk-configuration", "", "", "Compromised Credentials Risk Configuration")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderConfirmationCode, "confirmation-code", "", "", "Confirmation Code")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderContextData, "context-data", "", "", "Context Data")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderCredential, "credential", "", "", "Credential")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderCredentialId, "credential-id", "", "", "Credential ID")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderCSS, "css", "", "", "Css")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderCustomAttributes, "custom-attributes", "", "", "Custom Attributes")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderCustomDomainConfig, "custom-domain-config", "", "", "Custom Domain Config")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderDefaultRedirectURI, "default-redirect-uri", "", "", "Default Redirect URI")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderDeletionProtection, "deletion-protection", "", "", "Deletion Protection")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderDescription, "description", "", "", "Description")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderDesiredDeliveryMediums, "desired-delivery-mediums", "", "", "Desired Delivery Mediums")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderDestinationUser, "destination-user", "", "", "Destination User")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderDeviceConfiguration, "device-configuration", "", "", "Device Configuration")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderDeviceKey, "device-key", "", "", "Device Key")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderDeviceName, "device-name", "", "", "Device Name")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderDeviceRememberedStatus, "device-remembered-status", "", "", "Device Remembered Status")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderDeviceSecretVerifierConfig, "device-secret-verifier-config", "", "", "Device Secret Verifier Config")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderDomain, "domain", "", "", "Domain")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderEmailConfiguration, "email-configuration", "", "", "Email Configuration")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderEmailMfaConfiguration, "email-mfa-configuration", "", "", "Email MFA Configuration")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderEmailMfaSettings, "email-mfa-settings", "", "", "Email MFA Settings")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderEmailVerificationMessage, "email-verification-message", "", "", "Email Verification Message")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderEmailVerificationSubject, "email-verification-subject", "", "", "Email Verification Subject")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderEnablePropagateAdditionalUserContextData, "enable-propagate-additional-user-context-data", "", "", "Enable Propagate Additional User Context Data")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderEnableTokenRevocation, "enable-token-revocation", "", "", "Enable Token Revocation")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderEnforcement, "enforcement", "", "", "Enforcement")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderEventId, "event-id", "", "", "Event ID")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderExplicitAuthFlows, "explicit-auth-flows", "", "", "Explicit Auth Flows")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderFeedbackToken, "feedback-token", "", "", "Feedback Token")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderFeedbackValue, "feedback-value", "", "", "Feedback Value")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderFilter, "filter", "", "", "Filter")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderForceAliasCreation, "force-alias-creation", "", "", "Force Alias Creation")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderFriendlyDeviceName, "friendly-device-name", "", "", "Friendly Device Name")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderGenerateSecret, "generate-secret", "", "", "Generate Secret")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderGroupName, "group-name", "", "", "Group Name")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderIdTokenValidity, "id-token-validity", "", "", "ID Token Validity")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderIdentifier, "identifier", "", "", "Identifier")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderIdpIdentifier, "idp-identifier", "", "", "Idp Identifier")
	_cognitoidentityproviderCmd.Flags().StringSliceVarP(&_cognitoidentityproviderIdpIdentifiers, "idp-identifiers", "", nil, "Idp Identifiers")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderImageFile, "image-file", "", "", "Image File")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderJobId, "job-id", "", "", "Job ID")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderJobName, "job-name", "", "", "Job Name")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderLambdaConfig, "lambda-config", "", "", "Lambda Config")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderLimit, "limit", "", "", "Limit")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderLinks, "links", "", "", "Links")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderLogConfigurations, "log-configurations", "", "", "Log Configurations")
	_cognitoidentityproviderCmd.Flags().StringSliceVarP(&_cognitoidentityproviderLogoutURLs, "logout-urls", "", nil, "Logout Urls")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderManagedLoginBrandingId, "managed-login-branding-id", "", "", "Managed Login Branding ID")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderManagedLoginVersion, "managed-login-version", "", "", "Managed Login Version")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderMaxResults, "max-results", "", "", "Max Results")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderMessageAction, "message-action", "", "", "Message Action")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderMfaConfiguration, "mfa-configuration", "", "", "MFA Configuration")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderMFAOptions, "mfa-options", "", "", "MFA Options")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderName, "name", "", "", "Name")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderNextToken, "next-token", "", "", "Next Token")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderPaginationToken, "pagination-token", "", "", "Pagination Token")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderPassword, "password", "", "", "Password")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderPermanent, "permanent", "", "", "Permanent")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderPolicies, "policies", "", "", "Policies")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderPoolName, "pool-name", "", "", "Pool Name")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderPrecedence, "precedence", "", "", "Precedence")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderPreventUserExistenceErrors, "prevent-user-existence-errors", "", "", "Prevent User Existence Errors")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderPreviousPassword, "previous-password", "", "", "Previous Password")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderProposedPassword, "proposed-password", "", "", "Proposed Password")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderProviderDetails, "provider-details", "", "", "Provider Details")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderProviderName, "provider-name", "", "", "Provider Name")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderProviderType, "provider-type", "", "", "Provider Type")
	_cognitoidentityproviderCmd.Flags().StringSliceVarP(&_cognitoidentityproviderReadAttributes, "read-attributes", "", nil, "Read Attributes")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderRefreshToken, "refresh-token", "", "", "Refresh Token")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderRefreshTokenRotation, "refresh-token-rotation", "", "", "Refresh Token Rotation")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderRefreshTokenValidity, "refresh-token-validity", "", "", "Refresh Token Validity")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderResourceArn, "resource-arn", "", "", "Resource ARN")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderReturnMergedResources, "return-merged-resources", "", "", "Return Merged Resources")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderRiskExceptionConfiguration, "risk-exception-configuration", "", "", "Risk Exception Configuration")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderRoleArn, "role-arn", "", "", "Role ARN")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderSchema, "schema", "", "", "Schema")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderScopes, "scopes", "", "", "Scopes")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderSecretHash, "secret-hash", "", "", "Secret Hash")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderSession, "session", "", "", "Session")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderSettings, "settings", "", "", "Settings")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderSmsAuthenticationMessage, "sms-authentication-message", "", "", "Sms Authentication Message")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderSmsConfiguration, "sms-configuration", "", "", "Sms Configuration")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderSmsMfaConfiguration, "sms-mfa-configuration", "", "", "Sms MFA Configuration")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderSMSMfaSettings, "sms-mfa-settings", "", "", "Sms MFA Settings")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderSmsVerificationMessage, "sms-verification-message", "", "", "Sms Verification Message")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderSoftwareTokenMfaConfiguration, "software-token-mfa-configuration", "", "", "Software Token MFA Configuration")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderSoftwareTokenMfaSettings, "software-token-mfa-settings", "", "", "Software Token MFA Settings")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderSourceUser, "source-user", "", "", "Source User")
	_cognitoidentityproviderCmd.Flags().StringSliceVarP(&_cognitoidentityproviderSupportedIdentityProviders, "supported-identity-providers", "", nil, "Supported Identity Providers")
	_cognitoidentityproviderCmd.Flags().StringSliceVarP(&_cognitoidentityproviderTagKeys, "tag-keys", "", nil, "Tag Keys")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderTags, "tags", "", "", "Tags")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderTemporaryPassword, "temporary-password", "", "", "Temporary Password")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderTermsId, "terms-id", "", "", "Terms ID")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderTermsName, "terms-name", "", "", "Terms Name")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderTermsSource, "terms-source", "", "", "Terms Source")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderToken, "token", "", "", "Token")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderTokenValidityUnits, "token-validity-units", "", "", "Token Validity Units")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderUseCognitoProvidedValues, "use-cognito-provided-values", "", "", "Use Cognito Provided Values")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderUser, "user", "", "", "User")
	_cognitoidentityproviderCmd.Flags().StringSliceVarP(&_cognitoidentityproviderUserAttributeNames, "user-attribute-names", "", nil, "User Attribute Names")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderUserAttributeUpdateSettings, "user-attribute-update-settings", "", "", "User Attribute Update Settings")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderUserAttributes, "user-attributes", "", "", "User Attributes")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderUserCode, "user-code", "", "", "User Code")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderUserContextData, "user-context-data", "", "", "User Context Data")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderUserPoolAddOns, "user-pool-add-ons", "", "", "User Pool Add Ons")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderUserPoolId, "user-pool-id", "", "", "User Pool ID")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderUserPoolTags, "user-pool-tags", "", "", "User Pool Tags")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderUserPoolTier, "user-pool-tier", "", "", "User Pool Tier")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderUsername, "username", "", "", "Username")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderUsernameAttributes, "username-attributes", "", "", "Username Attributes")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderUsernameConfiguration, "username-configuration", "", "", "Username Configuration")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderValidationData, "validation-data", "", "", "Validation Data")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderVerificationMessageTemplate, "verification-message-template", "", "", "Verification Message Template")
	_cognitoidentityproviderCmd.Flags().StringVarP(&_cognitoidentityproviderWebAuthnConfiguration, "web-authn-configuration", "", "", "Web Authn Configuration")
	_cognitoidentityproviderCmd.Flags().StringSliceVarP(&_cognitoidentityproviderWriteAttributes, "write-attributes", "", nil, "Write Attributes")

	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderAddCustomAttributes, "add-custom-attributes", "", false, "Add Custom Attributes")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderAddUserPoolClientSecret, "add-user-pool-client-secret", "", false, "Add User Pool Client Secret")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderAdminAddUserToGroup, "admin-add-user-to-group", "", false, "Admin Add User To Group")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderAdminConfirmSignUp, "admin-confirm-sign-up", "", false, "Admin Confirm Sign Up")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderAdminCreateUser, "admin-create-user", "", false, "Admin Create User")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderAdminDeleteUser, "admin-delete-user", "", false, "Admin Delete User")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderAdminDeleteUserAttributes, "admin-delete-user-attributes", "", false, "Admin Delete User Attributes")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderAdminDisableProviderForUser, "admin-disable-provider-for-user", "", false, "Admin Disable Provider For User")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderAdminDisableUser, "admin-disable-user", "", false, "Admin Disable User")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderAdminEnableUser, "admin-enable-user", "", false, "Admin Enable User")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderAdminForgetDevice, "admin-forget-device", "", false, "Admin Forget Device")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderAdminGetDevice, "admin-get-device", "", false, "Admin Get Device")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderAdminGetUser, "admin-get-user", "", false, "Admin Get User")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderAdminInitiateAuth, "admin-initiate-auth", "", false, "Admin Initiate Auth")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderAdminLinkProviderForUser, "admin-link-provider-for-user", "", false, "Admin Link Provider For User")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderAdminListDevices, "admin-list-devices", "", false, "Admin List Devices")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderAdminListGroupsForUser, "admin-list-groups-for-user", "", false, "Admin List Groups For User")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderAdminListUserAuthEvents, "admin-list-user-auth-events", "", false, "Admin List User Auth Events")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderAdminRemoveUserFromGroup, "admin-remove-user-from-group", "", false, "Admin Remove User From Group")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderAdminResetUserPassword, "admin-reset-user-password", "", false, "Admin Reset User Password")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderAdminRespondToAuthChallenge, "admin-respond-to-auth-challenge", "", false, "Admin Respond To Auth Challenge")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderAdminSetUserMFAPreference, "admin-set-user-mfa-preference", "", false, "Admin Set User MFA Preference")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderAdminSetUserPassword, "admin-set-user-password", "", false, "Admin Set User Password")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderAdminSetUserSettings, "admin-set-user-settings", "", false, "Admin Set User Settings")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderAdminUpdateAuthEventFeedback, "admin-update-auth-event-feedback", "", false, "Admin Update Auth Event Feedback")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderAdminUpdateDeviceStatus, "admin-update-device-status", "", false, "Admin Update Device Status")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderAdminUpdateUserAttributes, "admin-update-user-attributes", "", false, "Admin Update User Attributes")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderAdminUserGlobalSignOut, "admin-user-global-sign-out", "", false, "Admin User Global Sign Out")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderAssociateSoftwareToken, "associate-software-token", "", false, "Associate Software Token")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderChangePassword, "change-password", "", false, "Change Password")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderCompleteWebAuthnRegistration, "complete-web-authn-registration", "", false, "Complete Web Authn Registration")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderConfirmDevice, "confirm-device", "", false, "Confirm Device")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderConfirmForgotPassword, "confirm-forgot-password", "", false, "Confirm Forgot Password")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderConfirmSignUp, "confirm-sign-up", "", false, "Confirm Sign Up")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderCreateGroup, "create-group", "", false, "Create Group")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderCreateIdentityProvider, "create-identity-provider", "", false, "Create Identity Provider")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderCreateManagedLoginBranding, "create-managed-login-branding", "", false, "Create Managed Login Branding")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderCreateResourceServer, "create-resource-server", "", false, "Create Resource Server")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderCreateTerms, "create-terms", "", false, "Create Terms")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderCreateUserImportJob, "create-user-import-job", "", false, "Create User Import Job")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderCreateUserPool, "create-user-pool", "", false, "Create User Pool")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderCreateUserPoolClient, "create-user-pool-client", "", false, "Create User Pool Client")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderCreateUserPoolDomain, "create-user-pool-domain", "", false, "Create User Pool Domain")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderDeleteGroup, "delete-group", "", false, "Delete Group")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderDeleteIdentityProvider, "delete-identity-provider", "", false, "Delete Identity Provider")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderDeleteManagedLoginBranding, "delete-managed-login-branding", "", false, "Delete Managed Login Branding")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderDeleteResourceServer, "delete-resource-server", "", false, "Delete Resource Server")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderDeleteTerms, "delete-terms", "", false, "Delete Terms")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderDeleteUser, "delete-user", "", false, "Delete User")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderDeleteUserAttributes, "delete-user-attributes", "", false, "Delete User Attributes")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderDeleteUserPool, "delete-user-pool", "", false, "Delete User Pool")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderDeleteUserPoolClient, "delete-user-pool-client", "", false, "Delete User Pool Client")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderDeleteUserPoolClientSecret, "delete-user-pool-client-secret", "", false, "Delete User Pool Client Secret")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderDeleteUserPoolDomain, "delete-user-pool-domain", "", false, "Delete User Pool Domain")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderDeleteWebAuthnCredential, "delete-web-authn-credential", "", false, "Delete Web Authn Credential")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderDescribeIdentityProvider, "describe-identity-provider", "", false, "Describe Identity Provider")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderDescribeManagedLoginBranding, "describe-managed-login-branding", "", false, "Describe Managed Login Branding")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderDescribeManagedLoginBrandingByClient, "describe-managed-login-branding-by-client", "", false, "Describe Managed Login Branding By Client")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderDescribeResourceServer, "describe-resource-server", "", false, "Describe Resource Server")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderDescribeRiskConfiguration, "describe-risk-configuration", "", false, "Describe Risk Configuration")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderDescribeTerms, "describe-terms", "", false, "Describe Terms")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderDescribeUserImportJob, "describe-user-import-job", "", false, "Describe User Import Job")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderDescribeUserPool, "describe-user-pool", "", false, "Describe User Pool")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderDescribeUserPoolClient, "describe-user-pool-client", "", false, "Describe User Pool Client")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderDescribeUserPoolDomain, "describe-user-pool-domain", "", false, "Describe User Pool Domain")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderForgetDevice, "forget-device", "", false, "Forget Device")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderForgotPassword, "forgot-password", "", false, "Forgot Password")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderGetCSVHeader, "get-csv-header", "", false, "Get CSV Header")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderGetDevice, "get-device", "", false, "Get Device")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderGetGroup, "get-group", "", false, "Get Group")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderGetIdentityProviderByIdentifier, "get-identity-provider-by-identifier", "", false, "Get Identity Provider By Identifier")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderGetLogDeliveryConfiguration, "get-log-delivery-configuration", "", false, "Get Log Delivery Configuration")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderGetSigningCertificate, "get-signing-certificate", "", false, "Get Signing Certificate")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderGetTokensFromRefreshToken, "get-tokens-from-refresh-token", "", false, "Get Tokens From Refresh Token")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderGetUICustomization, "get-ui-customization", "", false, "Get Ui Customization")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderGetUser, "get-user", "", false, "Get User")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderGetUserAttributeVerificationCode, "get-user-attribute-verification-code", "", false, "Get User Attribute Verification Code")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderGetUserAuthFactors, "get-user-auth-factors", "", false, "Get User Auth Factors")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderGetUserPoolMfaConfig, "get-user-pool-mfa-config", "", false, "Get User Pool MFA Config")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderGlobalSignOut, "global-sign-out", "", false, "Global Sign Out")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderInitiateAuth, "initiate-auth", "", false, "Initiate Auth")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderListDevices, "list-devices", "", false, "List Devices")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderListGroups, "list-groups", "", false, "List Groups")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderListIdentityProviders, "list-identity-providers", "", false, "List Identity Providers")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderListResourceServers, "list-resource-servers", "", false, "List Resource Servers")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderListTerms, "list-terms", "", false, "List Terms")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderListUserImportJobs, "list-user-import-jobs", "", false, "List User Import Jobs")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderListUserPoolClientSecrets, "list-user-pool-client-secrets", "", false, "List User Pool Client Secrets")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderListUserPoolClients, "list-user-pool-clients", "", false, "List User Pool Clients")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderListUserPools, "list-user-pools", "", false, "List User Pools")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderListUsers, "list-users", "", false, "List Users")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderListUsersInGroup, "list-users-in-group", "", false, "List Users In Group")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderListWebAuthnCredentials, "list-web-authn-credentials", "", false, "List Web Authn Credentials")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderResendConfirmationCode, "resend-confirmation-code", "", false, "Resend Confirmation Code")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderRespondToAuthChallenge, "respond-to-auth-challenge", "", false, "Respond To Auth Challenge")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderRevokeToken, "revoke-token", "", false, "Revoke Token")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderSetLogDeliveryConfiguration, "set-log-delivery-configuration", "", false, "Set Log Delivery Configuration")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderSetRiskConfiguration, "set-risk-configuration", "", false, "Set Risk Configuration")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderSetUICustomization, "set-ui-customization", "", false, "Set Ui Customization")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderSetUserMFAPreference, "set-user-mfa-preference", "", false, "Set User MFA Preference")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderSetUserPoolMfaConfig, "set-user-pool-mfa-config", "", false, "Set User Pool MFA Config")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderSetUserSettings, "set-user-settings", "", false, "Set User Settings")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderSignUp, "sign-up", "", false, "Sign Up")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderStartUserImportJob, "start-user-import-job", "", false, "Start User Import Job")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderStartWebAuthnRegistration, "start-web-authn-registration", "", false, "Start Web Authn Registration")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderStopUserImportJob, "stop-user-import-job", "", false, "Stop User Import Job")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderTagResource, "tag-resource", "", false, "Tag Resource")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderUntagResource, "untag-resource", "", false, "Untag Resource")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderUpdateAuthEventFeedback, "update-auth-event-feedback", "", false, "Update Auth Event Feedback")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderUpdateDeviceStatus, "update-device-status", "", false, "Update Device Status")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderUpdateGroup, "update-group", "", false, "Update Group")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderUpdateIdentityProvider, "update-identity-provider", "", false, "Update Identity Provider")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderUpdateManagedLoginBranding, "update-managed-login-branding", "", false, "Update Managed Login Branding")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderUpdateResourceServer, "update-resource-server", "", false, "Update Resource Server")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderUpdateTerms, "update-terms", "", false, "Update Terms")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderUpdateUserAttributes, "update-user-attributes", "", false, "Update User Attributes")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderUpdateUserPool, "update-user-pool", "", false, "Update User Pool")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderUpdateUserPoolClient, "update-user-pool-client", "", false, "Update User Pool Client")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderUpdateUserPoolDomain, "update-user-pool-domain", "", false, "Update User Pool Domain")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderVerifySoftwareToken, "verify-software-token", "", false, "Verify Software Token")
	_cognitoidentityproviderCmd.Flags().BoolVarP(&_cognitoidentityproviderVerifyUserAttribute, "verify-user-attribute", "", false, "Verify User Attribute")

}
