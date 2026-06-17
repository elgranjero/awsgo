package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// iamCmd represents the iam command
var _iamCmd = &cobra.Command{
	Use:   "iam",
	Short: "AWS iam CLI",
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
		client := iam.NewFromConfig(cfg)
		if _iamAcceptDelegationRequest {
			iam_AcceptDelegationRequest(cfg, client)
			return
		}
		if _iamAddClientIDToOpenIDConnectProvider {
			iam_AddClientIDToOpenIDConnectProvider(cfg, client)
			return
		}
		if _iamAddRoleToInstanceProfile {
			iam_AddRoleToInstanceProfile(cfg, client)
			return
		}
		if _iamAddUserToGroup {
			iam_AddUserToGroup(cfg, client)
			return
		}
		if _iamAssociateDelegationRequest {
			iam_AssociateDelegationRequest(cfg, client)
			return
		}
		if _iamAttachGroupPolicy {
			iam_AttachGroupPolicy(cfg, client)
			return
		}
		if _iamAttachRolePolicy {
			iam_AttachRolePolicy(cfg, client)
			return
		}
		if _iamAttachUserPolicy {
			iam_AttachUserPolicy(cfg, client)
			return
		}
		if _iamChangePassword {
			iam_ChangePassword(cfg, client)
			return
		}
		if _iamCreateAccessKey {
			iam_CreateAccessKey(cfg, client)
			return
		}
		if _iamCreateAccountAlias {
			iam_CreateAccountAlias(cfg, client)
			return
		}
		if _iamCreateDelegationRequest {
			iam_CreateDelegationRequest(cfg, client)
			return
		}
		if _iamCreateGroup {
			iam_CreateGroup(cfg, client)
			return
		}
		if _iamCreateInstanceProfile {
			iam_CreateInstanceProfile(cfg, client)
			return
		}
		if _iamCreateLoginProfile {
			iam_CreateLoginProfile(cfg, client)
			return
		}
		if _iamCreateOpenIDConnectProvider {
			iam_CreateOpenIDConnectProvider(cfg, client)
			return
		}
		if _iamCreatePolicy {
			iam_CreatePolicy(cfg, client)
			return
		}
		if _iamCreatePolicyVersion {
			iam_CreatePolicyVersion(cfg, client)
			return
		}
		if _iamCreateRole {
			iam_CreateRole(cfg, client)
			return
		}
		if _iamCreateSAMLProvider {
			iam_CreateSAMLProvider(cfg, client)
			return
		}
		if _iamCreateServiceLinkedRole {
			iam_CreateServiceLinkedRole(cfg, client)
			return
		}
		if _iamCreateServiceSpecificCredential {
			iam_CreateServiceSpecificCredential(cfg, client)
			return
		}
		if _iamCreateUser {
			iam_CreateUser(cfg, client)
			return
		}
		if _iamCreateVirtualMFADevice {
			iam_CreateVirtualMFADevice(cfg, client)
			return
		}
		if _iamDeactivateMFADevice {
			iam_DeactivateMFADevice(cfg, client)
			return
		}
		if _iamDeleteAccessKey {
			iam_DeleteAccessKey(cfg, client)
			return
		}
		if _iamDeleteAccountAlias {
			iam_DeleteAccountAlias(cfg, client)
			return
		}
		if _iamDeleteAccountPasswordPolicy {
			iam_DeleteAccountPasswordPolicy(cfg, client)
			return
		}
		if _iamDeleteGroup {
			iam_DeleteGroup(cfg, client)
			return
		}
		if _iamDeleteGroupPolicy {
			iam_DeleteGroupPolicy(cfg, client)
			return
		}
		if _iamDeleteInstanceProfile {
			iam_DeleteInstanceProfile(cfg, client)
			return
		}
		if _iamDeleteLoginProfile {
			iam_DeleteLoginProfile(cfg, client)
			return
		}
		if _iamDeleteOpenIDConnectProvider {
			iam_DeleteOpenIDConnectProvider(cfg, client)
			return
		}
		if _iamDeletePolicy {
			iam_DeletePolicy(cfg, client)
			return
		}
		if _iamDeletePolicyVersion {
			iam_DeletePolicyVersion(cfg, client)
			return
		}
		if _iamDeleteRole {
			iam_DeleteRole(cfg, client)
			return
		}
		if _iamDeleteRolePermissionsBoundary {
			iam_DeleteRolePermissionsBoundary(cfg, client)
			return
		}
		if _iamDeleteRolePolicy {
			iam_DeleteRolePolicy(cfg, client)
			return
		}
		if _iamDeleteSAMLProvider {
			iam_DeleteSAMLProvider(cfg, client)
			return
		}
		if _iamDeleteServerCertificate {
			iam_DeleteServerCertificate(cfg, client)
			return
		}
		if _iamDeleteServiceLinkedRole {
			iam_DeleteServiceLinkedRole(cfg, client)
			return
		}
		if _iamDeleteServiceSpecificCredential {
			iam_DeleteServiceSpecificCredential(cfg, client)
			return
		}
		if _iamDeleteSigningCertificate {
			iam_DeleteSigningCertificate(cfg, client)
			return
		}
		if _iamDeleteSSHPublicKey {
			iam_DeleteSSHPublicKey(cfg, client)
			return
		}
		if _iamDeleteUser {
			iam_DeleteUser(cfg, client)
			return
		}
		if _iamDeleteUserPermissionsBoundary {
			iam_DeleteUserPermissionsBoundary(cfg, client)
			return
		}
		if _iamDeleteUserPolicy {
			iam_DeleteUserPolicy(cfg, client)
			return
		}
		if _iamDeleteVirtualMFADevice {
			iam_DeleteVirtualMFADevice(cfg, client)
			return
		}
		if _iamDetachGroupPolicy {
			iam_DetachGroupPolicy(cfg, client)
			return
		}
		if _iamDetachRolePolicy {
			iam_DetachRolePolicy(cfg, client)
			return
		}
		if _iamDetachUserPolicy {
			iam_DetachUserPolicy(cfg, client)
			return
		}
		if _iamDisableOrganizationsRootCredentialsManagement {
			iam_DisableOrganizationsRootCredentialsManagement(cfg, client)
			return
		}
		if _iamDisableOrganizationsRootSessions {
			iam_DisableOrganizationsRootSessions(cfg, client)
			return
		}
		if _iamDisableOutboundWebIdentityFederation {
			iam_DisableOutboundWebIdentityFederation(cfg, client)
			return
		}
		if _iamEnableMFADevice {
			iam_EnableMFADevice(cfg, client)
			return
		}
		if _iamEnableOrganizationsRootCredentialsManagement {
			iam_EnableOrganizationsRootCredentialsManagement(cfg, client)
			return
		}
		if _iamEnableOrganizationsRootSessions {
			iam_EnableOrganizationsRootSessions(cfg, client)
			return
		}
		if _iamEnableOutboundWebIdentityFederation {
			iam_EnableOutboundWebIdentityFederation(cfg, client)
			return
		}
		if _iamGenerateCredentialReport {
			iam_GenerateCredentialReport(cfg, client)
			return
		}
		if _iamGenerateOrganizationsAccessReport {
			iam_GenerateOrganizationsAccessReport(cfg, client)
			return
		}
		if _iamGenerateServiceLastAccessedDetails {
			iam_GenerateServiceLastAccessedDetails(cfg, client)
			return
		}
		if _iamGetAccessKeyLastUsed {
			iam_GetAccessKeyLastUsed(cfg, client)
			return
		}
		if _iamGetAccountAuthorizationDetails {
			iam_GetAccountAuthorizationDetails(cfg, client)
			return
		}
		if _iamGetAccountPasswordPolicy {
			iam_GetAccountPasswordPolicy(cfg, client)
			return
		}
		if _iamGetAccountSummary {
			iam_GetAccountSummary(cfg, client)
			return
		}
		if _iamGetContextKeysForCustomPolicy {
			iam_GetContextKeysForCustomPolicy(cfg, client)
			return
		}
		if _iamGetContextKeysForPrincipalPolicy {
			iam_GetContextKeysForPrincipalPolicy(cfg, client)
			return
		}
		if _iamGetCredentialReport {
			iam_GetCredentialReport(cfg, client)
			return
		}
		if _iamGetDelegationRequest {
			iam_GetDelegationRequest(cfg, client)
			return
		}
		if _iamGetGroup {
			iam_GetGroup(cfg, client)
			return
		}
		if _iamGetGroupPolicy {
			iam_GetGroupPolicy(cfg, client)
			return
		}
		if _iamGetHumanReadableSummary {
			iam_GetHumanReadableSummary(cfg, client)
			return
		}
		if _iamGetInstanceProfile {
			iam_GetInstanceProfile(cfg, client)
			return
		}
		if _iamGetLoginProfile {
			iam_GetLoginProfile(cfg, client)
			return
		}
		if _iamGetMFADevice {
			iam_GetMFADevice(cfg, client)
			return
		}
		if _iamGetOpenIDConnectProvider {
			iam_GetOpenIDConnectProvider(cfg, client)
			return
		}
		if _iamGetOrganizationsAccessReport {
			iam_GetOrganizationsAccessReport(cfg, client)
			return
		}
		if _iamGetOutboundWebIdentityFederationInfo {
			iam_GetOutboundWebIdentityFederationInfo(cfg, client)
			return
		}
		if _iamGetPolicy {
			iam_GetPolicy(cfg, client)
			return
		}
		if _iamGetPolicyVersion {
			iam_GetPolicyVersion(cfg, client)
			return
		}
		if _iamGetRole {
			iam_GetRole(cfg, client)
			return
		}
		if _iamGetRolePolicy {
			iam_GetRolePolicy(cfg, client)
			return
		}
		if _iamGetSAMLProvider {
			iam_GetSAMLProvider(cfg, client)
			return
		}
		if _iamGetServerCertificate {
			iam_GetServerCertificate(cfg, client)
			return
		}
		if _iamGetServiceLastAccessedDetails {
			iam_GetServiceLastAccessedDetails(cfg, client)
			return
		}
		if _iamGetServiceLastAccessedDetailsWithEntities {
			iam_GetServiceLastAccessedDetailsWithEntities(cfg, client)
			return
		}
		if _iamGetServiceLinkedRoleDeletionStatus {
			iam_GetServiceLinkedRoleDeletionStatus(cfg, client)
			return
		}
		if _iamGetSSHPublicKey {
			iam_GetSSHPublicKey(cfg, client)
			return
		}
		if _iamGetUser {
			iam_GetUser(cfg, client)
			return
		}
		if _iamGetUserPolicy {
			iam_GetUserPolicy(cfg, client)
			return
		}
		if _iamListAccessKeys {
			iam_ListAccessKeys(cfg, client)
			return
		}
		if _iamListAccountAliases {
			iam_ListAccountAliases(cfg, client)
			return
		}
		if _iamListAttachedGroupPolicies {
			iam_ListAttachedGroupPolicies(cfg, client)
			return
		}
		if _iamListAttachedRolePolicies {
			iam_ListAttachedRolePolicies(cfg, client)
			return
		}
		if _iamListAttachedUserPolicies {
			iam_ListAttachedUserPolicies(cfg, client)
			return
		}
		if _iamListDelegationRequests {
			iam_ListDelegationRequests(cfg, client)
			return
		}
		if _iamListEntitiesForPolicy {
			iam_ListEntitiesForPolicy(cfg, client)
			return
		}
		if _iamListGroupPolicies {
			iam_ListGroupPolicies(cfg, client)
			return
		}
		if _iamListGroups {
			iam_ListGroups(cfg, client)
			return
		}
		if _iamListGroupsForUser {
			iam_ListGroupsForUser(cfg, client)
			return
		}
		if _iamListInstanceProfileTags {
			iam_ListInstanceProfileTags(cfg, client)
			return
		}
		if _iamListInstanceProfiles {
			iam_ListInstanceProfiles(cfg, client)
			return
		}
		if _iamListInstanceProfilesForRole {
			iam_ListInstanceProfilesForRole(cfg, client)
			return
		}
		if _iamListMFADeviceTags {
			iam_ListMFADeviceTags(cfg, client)
			return
		}
		if _iamListMFADevices {
			iam_ListMFADevices(cfg, client)
			return
		}
		if _iamListOpenIDConnectProviderTags {
			iam_ListOpenIDConnectProviderTags(cfg, client)
			return
		}
		if _iamListOpenIDConnectProviders {
			iam_ListOpenIDConnectProviders(cfg, client)
			return
		}
		if _iamListOrganizationsFeatures {
			iam_ListOrganizationsFeatures(cfg, client)
			return
		}
		if _iamListPolicies {
			iam_ListPolicies(cfg, client)
			return
		}
		if _iamListPoliciesGrantingServiceAccess {
			iam_ListPoliciesGrantingServiceAccess(cfg, client)
			return
		}
		if _iamListPolicyTags {
			iam_ListPolicyTags(cfg, client)
			return
		}
		if _iamListPolicyVersions {
			iam_ListPolicyVersions(cfg, client)
			return
		}
		if _iamListRolePolicies {
			iam_ListRolePolicies(cfg, client)
			return
		}
		if _iamListRoleTags {
			iam_ListRoleTags(cfg, client)
			return
		}
		if _iamListRoles {
			iam_ListRoles(cfg, client)
			return
		}
		if _iamListSAMLProviderTags {
			iam_ListSAMLProviderTags(cfg, client)
			return
		}
		if _iamListSAMLProviders {
			iam_ListSAMLProviders(cfg, client)
			return
		}
		if _iamListServerCertificateTags {
			iam_ListServerCertificateTags(cfg, client)
			return
		}
		if _iamListServerCertificates {
			iam_ListServerCertificates(cfg, client)
			return
		}
		if _iamListServiceSpecificCredentials {
			iam_ListServiceSpecificCredentials(cfg, client)
			return
		}
		if _iamListSigningCertificates {
			iam_ListSigningCertificates(cfg, client)
			return
		}
		if _iamListSSHPublicKeys {
			iam_ListSSHPublicKeys(cfg, client)
			return
		}
		if _iamListUserPolicies {
			iam_ListUserPolicies(cfg, client)
			return
		}
		if _iamListUserTags {
			iam_ListUserTags(cfg, client)
			return
		}
		if _iamListUsers {
			iam_ListUsers(cfg, client)
			return
		}
		if _iamListVirtualMFADevices {
			iam_ListVirtualMFADevices(cfg, client)
			return
		}
		if _iamPutGroupPolicy {
			iam_PutGroupPolicy(cfg, client)
			return
		}
		if _iamPutRolePermissionsBoundary {
			iam_PutRolePermissionsBoundary(cfg, client)
			return
		}
		if _iamPutRolePolicy {
			iam_PutRolePolicy(cfg, client)
			return
		}
		if _iamPutUserPermissionsBoundary {
			iam_PutUserPermissionsBoundary(cfg, client)
			return
		}
		if _iamPutUserPolicy {
			iam_PutUserPolicy(cfg, client)
			return
		}
		if _iamRejectDelegationRequest {
			iam_RejectDelegationRequest(cfg, client)
			return
		}
		if _iamRemoveClientIDFromOpenIDConnectProvider {
			iam_RemoveClientIDFromOpenIDConnectProvider(cfg, client)
			return
		}
		if _iamRemoveRoleFromInstanceProfile {
			iam_RemoveRoleFromInstanceProfile(cfg, client)
			return
		}
		if _iamRemoveUserFromGroup {
			iam_RemoveUserFromGroup(cfg, client)
			return
		}
		if _iamResetServiceSpecificCredential {
			iam_ResetServiceSpecificCredential(cfg, client)
			return
		}
		if _iamResyncMFADevice {
			iam_ResyncMFADevice(cfg, client)
			return
		}
		if _iamSendDelegationToken {
			iam_SendDelegationToken(cfg, client)
			return
		}
		if _iamSetDefaultPolicyVersion {
			iam_SetDefaultPolicyVersion(cfg, client)
			return
		}
		if _iamSetSecurityTokenServicePreferences {
			iam_SetSecurityTokenServicePreferences(cfg, client)
			return
		}
		if _iamSimulateCustomPolicy {
			iam_SimulateCustomPolicy(cfg, client)
			return
		}
		if _iamSimulatePrincipalPolicy {
			iam_SimulatePrincipalPolicy(cfg, client)
			return
		}
		if _iamTagInstanceProfile {
			iam_TagInstanceProfile(cfg, client)
			return
		}
		if _iamTagMFADevice {
			iam_TagMFADevice(cfg, client)
			return
		}
		if _iamTagOpenIDConnectProvider {
			iam_TagOpenIDConnectProvider(cfg, client)
			return
		}
		if _iamTagPolicy {
			iam_TagPolicy(cfg, client)
			return
		}
		if _iamTagRole {
			iam_TagRole(cfg, client)
			return
		}
		if _iamTagSAMLProvider {
			iam_TagSAMLProvider(cfg, client)
			return
		}
		if _iamTagServerCertificate {
			iam_TagServerCertificate(cfg, client)
			return
		}
		if _iamTagUser {
			iam_TagUser(cfg, client)
			return
		}
		if _iamUntagInstanceProfile {
			iam_UntagInstanceProfile(cfg, client)
			return
		}
		if _iamUntagMFADevice {
			iam_UntagMFADevice(cfg, client)
			return
		}
		if _iamUntagOpenIDConnectProvider {
			iam_UntagOpenIDConnectProvider(cfg, client)
			return
		}
		if _iamUntagPolicy {
			iam_UntagPolicy(cfg, client)
			return
		}
		if _iamUntagRole {
			iam_UntagRole(cfg, client)
			return
		}
		if _iamUntagSAMLProvider {
			iam_UntagSAMLProvider(cfg, client)
			return
		}
		if _iamUntagServerCertificate {
			iam_UntagServerCertificate(cfg, client)
			return
		}
		if _iamUntagUser {
			iam_UntagUser(cfg, client)
			return
		}
		if _iamUpdateAccessKey {
			iam_UpdateAccessKey(cfg, client)
			return
		}
		if _iamUpdateAccountPasswordPolicy {
			iam_UpdateAccountPasswordPolicy(cfg, client)
			return
		}
		if _iamUpdateAssumeRolePolicy {
			iam_UpdateAssumeRolePolicy(cfg, client)
			return
		}
		if _iamUpdateDelegationRequest {
			iam_UpdateDelegationRequest(cfg, client)
			return
		}
		if _iamUpdateGroup {
			iam_UpdateGroup(cfg, client)
			return
		}
		if _iamUpdateLoginProfile {
			iam_UpdateLoginProfile(cfg, client)
			return
		}
		if _iamUpdateOpenIDConnectProviderThumbprint {
			iam_UpdateOpenIDConnectProviderThumbprint(cfg, client)
			return
		}
		if _iamUpdateRole {
			iam_UpdateRole(cfg, client)
			return
		}
		if _iamUpdateRoleDescription {
			iam_UpdateRoleDescription(cfg, client)
			return
		}
		if _iamUpdateSAMLProvider {
			iam_UpdateSAMLProvider(cfg, client)
			return
		}
		if _iamUpdateServerCertificate {
			iam_UpdateServerCertificate(cfg, client)
			return
		}
		if _iamUpdateServiceSpecificCredential {
			iam_UpdateServiceSpecificCredential(cfg, client)
			return
		}
		if _iamUpdateSigningCertificate {
			iam_UpdateSigningCertificate(cfg, client)
			return
		}
		if _iamUpdateSSHPublicKey {
			iam_UpdateSSHPublicKey(cfg, client)
			return
		}
		if _iamUpdateUser {
			iam_UpdateUser(cfg, client)
			return
		}
		if _iamUploadServerCertificate {
			iam_UploadServerCertificate(cfg, client)
			return
		}
		if _iamUploadSigningCertificate {
			iam_UploadSigningCertificate(cfg, client)
			return
		}
		if _iamUploadSSHPublicKey {
			iam_UploadSSHPublicKey(cfg, client)
			return
		}

	},
}

var (
	_iamAcceptDelegationRequest                       bool
	_iamAddClientIDToOpenIDConnectProvider            bool
	_iamAddRoleToInstanceProfile                      bool
	_iamAddUserToGroup                                bool
	_iamAssociateDelegationRequest                    bool
	_iamAttachGroupPolicy                             bool
	_iamAttachRolePolicy                              bool
	_iamAttachUserPolicy                              bool
	_iamChangePassword                                bool
	_iamCreateAccessKey                               bool
	_iamCreateAccountAlias                            bool
	_iamCreateDelegationRequest                       bool
	_iamCreateGroup                                   bool
	_iamCreateInstanceProfile                         bool
	_iamCreateLoginProfile                            bool
	_iamCreateOpenIDConnectProvider                   bool
	_iamCreatePolicy                                  bool
	_iamCreatePolicyVersion                           bool
	_iamCreateRole                                    bool
	_iamCreateSAMLProvider                            bool
	_iamCreateServiceLinkedRole                       bool
	_iamCreateServiceSpecificCredential               bool
	_iamCreateUser                                    bool
	_iamCreateVirtualMFADevice                        bool
	_iamDeactivateMFADevice                           bool
	_iamDeleteAccessKey                               bool
	_iamDeleteAccountAlias                            bool
	_iamDeleteAccountPasswordPolicy                   bool
	_iamDeleteGroup                                   bool
	_iamDeleteGroupPolicy                             bool
	_iamDeleteInstanceProfile                         bool
	_iamDeleteLoginProfile                            bool
	_iamDeleteOpenIDConnectProvider                   bool
	_iamDeletePolicy                                  bool
	_iamDeletePolicyVersion                           bool
	_iamDeleteRole                                    bool
	_iamDeleteRolePermissionsBoundary                 bool
	_iamDeleteRolePolicy                              bool
	_iamDeleteSAMLProvider                            bool
	_iamDeleteServerCertificate                       bool
	_iamDeleteServiceLinkedRole                       bool
	_iamDeleteServiceSpecificCredential               bool
	_iamDeleteSigningCertificate                      bool
	_iamDeleteSSHPublicKey                            bool
	_iamDeleteUser                                    bool
	_iamDeleteUserPermissionsBoundary                 bool
	_iamDeleteUserPolicy                              bool
	_iamDeleteVirtualMFADevice                        bool
	_iamDetachGroupPolicy                             bool
	_iamDetachRolePolicy                              bool
	_iamDetachUserPolicy                              bool
	_iamDisableOrganizationsRootCredentialsManagement bool
	_iamDisableOrganizationsRootSessions              bool
	_iamDisableOutboundWebIdentityFederation          bool
	_iamEnableMFADevice                               bool
	_iamEnableOrganizationsRootCredentialsManagement  bool
	_iamEnableOrganizationsRootSessions               bool
	_iamEnableOutboundWebIdentityFederation           bool
	_iamGenerateCredentialReport                      bool
	_iamGenerateOrganizationsAccessReport             bool
	_iamGenerateServiceLastAccessedDetails            bool
	_iamGetAccessKeyLastUsed                          bool
	_iamGetAccountAuthorizationDetails                bool
	_iamGetAccountPasswordPolicy                      bool
	_iamGetAccountSummary                             bool
	_iamGetContextKeysForCustomPolicy                 bool
	_iamGetContextKeysForPrincipalPolicy              bool
	_iamGetCredentialReport                           bool
	_iamGetDelegationRequest                          bool
	_iamGetGroup                                      bool
	_iamGetGroupPolicy                                bool
	_iamGetHumanReadableSummary                       bool
	_iamGetInstanceProfile                            bool
	_iamGetLoginProfile                               bool
	_iamGetMFADevice                                  bool
	_iamGetOpenIDConnectProvider                      bool
	_iamGetOrganizationsAccessReport                  bool
	_iamGetOutboundWebIdentityFederationInfo          bool
	_iamGetPolicy                                     bool
	_iamGetPolicyVersion                              bool
	_iamGetRole                                       bool
	_iamGetRolePolicy                                 bool
	_iamGetSAMLProvider                               bool
	_iamGetServerCertificate                          bool
	_iamGetServiceLastAccessedDetails                 bool
	_iamGetServiceLastAccessedDetailsWithEntities     bool
	_iamGetServiceLinkedRoleDeletionStatus            bool
	_iamGetSSHPublicKey                               bool
	_iamGetUser                                       bool
	_iamGetUserPolicy                                 bool
	_iamListAccessKeys                                bool
	_iamListAccountAliases                            bool
	_iamListAttachedGroupPolicies                     bool
	_iamListAttachedRolePolicies                      bool
	_iamListAttachedUserPolicies                      bool
	_iamListDelegationRequests                        bool
	_iamListEntitiesForPolicy                         bool
	_iamListGroupPolicies                             bool
	_iamListGroups                                    bool
	_iamListGroupsForUser                             bool
	_iamListInstanceProfileTags                       bool
	_iamListInstanceProfiles                          bool
	_iamListInstanceProfilesForRole                   bool
	_iamListMFADeviceTags                             bool
	_iamListMFADevices                                bool
	_iamListOpenIDConnectProviderTags                 bool
	_iamListOpenIDConnectProviders                    bool
	_iamListOrganizationsFeatures                     bool
	_iamListPolicies                                  bool
	_iamListPoliciesGrantingServiceAccess             bool
	_iamListPolicyTags                                bool
	_iamListPolicyVersions                            bool
	_iamListRolePolicies                              bool
	_iamListRoleTags                                  bool
	_iamListRoles                                     bool
	_iamListSAMLProviderTags                          bool
	_iamListSAMLProviders                             bool
	_iamListServerCertificateTags                     bool
	_iamListServerCertificates                        bool
	_iamListServiceSpecificCredentials                bool
	_iamListSigningCertificates                       bool
	_iamListSSHPublicKeys                             bool
	_iamListUserPolicies                              bool
	_iamListUserTags                                  bool
	_iamListUsers                                     bool
	_iamListVirtualMFADevices                         bool
	_iamPutGroupPolicy                                bool
	_iamPutRolePermissionsBoundary                    bool
	_iamPutRolePolicy                                 bool
	_iamPutUserPermissionsBoundary                    bool
	_iamPutUserPolicy                                 bool
	_iamRejectDelegationRequest                       bool
	_iamRemoveClientIDFromOpenIDConnectProvider       bool
	_iamRemoveRoleFromInstanceProfile                 bool
	_iamRemoveUserFromGroup                           bool
	_iamResetServiceSpecificCredential                bool
	_iamResyncMFADevice                               bool
	_iamSendDelegationToken                           bool
	_iamSetDefaultPolicyVersion                       bool
	_iamSetSecurityTokenServicePreferences            bool
	_iamSimulateCustomPolicy                          bool
	_iamSimulatePrincipalPolicy                       bool
	_iamTagInstanceProfile                            bool
	_iamTagMFADevice                                  bool
	_iamTagOpenIDConnectProvider                      bool
	_iamTagPolicy                                     bool
	_iamTagRole                                       bool
	_iamTagSAMLProvider                               bool
	_iamTagServerCertificate                          bool
	_iamTagUser                                       bool
	_iamUntagInstanceProfile                          bool
	_iamUntagMFADevice                                bool
	_iamUntagOpenIDConnectProvider                    bool
	_iamUntagPolicy                                   bool
	_iamUntagRole                                     bool
	_iamUntagSAMLProvider                             bool
	_iamUntagServerCertificate                        bool
	_iamUntagUser                                     bool
	_iamUpdateAccessKey                               bool
	_iamUpdateAccountPasswordPolicy                   bool
	_iamUpdateAssumeRolePolicy                        bool
	_iamUpdateDelegationRequest                       bool
	_iamUpdateGroup                                   bool
	_iamUpdateLoginProfile                            bool
	_iamUpdateOpenIDConnectProviderThumbprint         bool
	_iamUpdateRole                                    bool
	_iamUpdateRoleDescription                         bool
	_iamUpdateSAMLProvider                            bool
	_iamUpdateServerCertificate                       bool
	_iamUpdateServiceSpecificCredential               bool
	_iamUpdateSigningCertificate                      bool
	_iamUpdateSSHPublicKey                            bool
	_iamUpdateUser                                    bool
	_iamUploadServerCertificate                       bool
	_iamUploadSigningCertificate                      bool
	_iamUploadSSHPublicKey                            bool

	_iamAccessKeyId                        string
	_iamAccountAlias                       string
	_iamActionNames                        []string
	_iamAddPrivateKey                      string
	_iamAllUsers                           string
	_iamAllowUsersToChangePassword         string
	_iamArn                                string
	_iamAssertionEncryptionMode            string
	_iamAssignmentStatus                   string
	_iamAssumeRolePolicyDocument           string
	_iamAuthenticationCode1                string
	_iamAuthenticationCode2                string
	_iamAWSServiceName                     string
	_iamCallerArn                          string
	_iamCertificateBody                    string
	_iamCertificateChain                   string
	_iamCertificateId                      string
	_iamClientID                           string
	_iamClientIDList                       []string
	_iamContextEntries                     string
	_iamCredentialAgeDays                  string
	_iamCustomSuffix                       string
	_iamDelegationPermissionCheck          string
	_iamDelegationRequestId                string
	_iamDeletionTaskId                     string
	_iamDescription                        string
	_iamEncoding                           string
	_iamEntityArn                          string
	_iamEntityFilter                       string
	_iamEntityPath                         string
	_iamFilter                             string
	_iamGlobalEndpointTokenVersion         string
	_iamGranularity                        string
	_iamGroupName                          string
	_iamHardExpiry                         string
	_iamInstanceProfileName                string
	_iamJobId                              string
	_iamLocale                             string
	_iamMarker                             string
	_iamMaxItems                           string
	_iamMaxPasswordAge                     string
	_iamMaxSessionDuration                 string
	_iamMinimumPasswordLength              string
	_iamName                               string
	_iamNewGroupName                       string
	_iamNewPassword                        string
	_iamNewPath                            string
	_iamNewServerCertificateName           string
	_iamNewUserName                        string
	_iamNotes                              string
	_iamNotificationChannel                string
	_iamOldPassword                        string
	_iamOnlyAttached                       string
	_iamOnlySendByOwner                    string
	_iamOpenIDConnectProviderArn           string
	_iamOrganizationsPolicyId              string
	_iamOwnerAccountId                     string
	_iamOwnerId                            string
	_iamPassword                           string
	_iamPasswordResetRequired              string
	_iamPasswordReusePrevention            string
	_iamPath                               string
	_iamPathPrefix                         string
	_iamPermissions                        string
	_iamPermissionsBoundary                string
	_iamPermissionsBoundaryPolicyInputList []string
	_iamPolicyArn                          string
	_iamPolicyDocument                     string
	_iamPolicyInputList                    []string
	_iamPolicyName                         string
	_iamPolicySourceArn                    string
	_iamPolicyUsageFilter                  string
	_iamPrivateKey                         string
	_iamRedirectUrl                        string
	_iamRemovePrivateKey                   string
	_iamRequestMessage                     string
	_iamRequestorWorkflowId                string
	_iamRequireLowercaseCharacters         string
	_iamRequireNumbers                     string
	_iamRequireSymbols                     string
	_iamRequireUppercaseCharacters         string
	_iamResourceArns                       []string
	_iamResourceHandlingOption             string
	_iamResourceOwner                      string
	_iamResourcePolicy                     string
	_iamRoleName                           string
	_iamSAMLMetadataDocument               string
	_iamSAMLProviderArn                    string
	_iamScope                              string
	_iamSerialNumber                       string
	_iamServerCertificateName              string
	_iamServiceName                        string
	_iamServiceNamespace                   string
	_iamServiceNamespaces                  []string
	_iamServiceSpecificCredentialId        string
	_iamSessionDuration                    string
	_iamSetAsDefault                       string
	_iamSortKey                            string
	_iamSSHPublicKeyBody                   string
	_iamSSHPublicKeyId                     string
	_iamStatus                             string
	_iamTagKeys                            []string
	_iamTags                               string
	_iamThumbprintList                     []string
	_iamUrl                                string
	_iamUserName                           string
	_iamVersionId                          string
	_iamVirtualMFADeviceName               string
)

// Accepts a delegation request, granting the requested temporary access.
// Once the delegation request is accepted, it is eligible to send the exchange
// token to the partner. The [SendDelegationToken]API has to be explicitly called to send the
// delegation token.
//
// At the time of acceptance, IAM records the details and the state of the
// identity that called this API. This is the identity that gets mapped to the
// delegated credential.
//
// An accepted request may be rejected before the exchange token is sent to the
// partner.
//
// [SendDelegationToken]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_SendDelegationToken.html
func iam_AcceptDelegationRequest(cfg aws.Config, client *iam.Client) {
	input := &iam.AcceptDelegationRequestInput{
		// DelegationRequestId: *string, // Required
	}

	if len(_iamDelegationRequestId) > 0 {
		input.DelegationRequestId = aws.String(_iamDelegationRequestId)
	}

	if resp, err := client.AcceptDelegationRequest(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds a new client ID (also known as audience) to the list of client IDs already
// registered for the specified IAM OpenID Connect (OIDC) provider resource.
//
// This operation is idempotent; it does not fail or return an error if you add an
// existing client ID to the provider.
func iam_AddClientIDToOpenIDConnectProvider(cfg aws.Config, client *iam.Client) {
	input := &iam.AddClientIDToOpenIDConnectProviderInput{
		// ClientID: *string, // Required
		// OpenIDConnectProviderArn: *string, // Required
	}

	if len(_iamClientID) > 0 {
		input.ClientID = aws.String(_iamClientID)
	}
	if len(_iamOpenIDConnectProviderArn) > 0 {
		input.OpenIDConnectProviderArn = aws.String(_iamOpenIDConnectProviderArn)
	}

	if resp, err := client.AddClientIDToOpenIDConnectProvider(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds the specified IAM role to the specified instance profile. An instance
// profile can contain only one role, and this quota cannot be increased. You can
// remove the existing role and then add a different role to an instance profile.
// You must then wait for the change to appear across all of Amazon Web Services
// because of [eventual consistency]. To force the change, you must [disassociate the instance profile] and then [associate the instance profile], or you can stop your
// instance and then restart it.
//
// The caller of this operation must be granted the PassRole permission on the IAM
// role by a permissions policy.
//
// When using the [iam:AssociatedResourceArn] condition in a policy to restrict the [PassRole] IAM action, special
// considerations apply if the policy is intended to define access for the
// AddRoleToInstanceProfile action. In this case, you cannot specify a Region or
// instance ID in the EC2 instance ARN. The ARN value must be
// arn:aws:ec2:*:CallerAccountId:instance/* . Using any other ARN value may lead to
// unexpected evaluation results.
//
// For more information about roles, see [IAM roles] in the IAM User Guide. For more
// information about instance profiles, see [Using instance profiles]in the IAM User Guide.
//
// [disassociate the instance profile]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DisassociateIamInstanceProfile.html
// [associate the instance profile]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_AssociateIamInstanceProfile.html
// [Using instance profiles]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use_switch-role-ec2_instance-profiles.html
// [PassRole]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use_passrole.html
// [iam:AssociatedResourceArn]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_iam-condition-keys.html#available-keys-for-iam
// [eventual consistency]: https://en.wikipedia.org/wiki/Eventual_consistency
// [IAM roles]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles.html
func iam_AddRoleToInstanceProfile(cfg aws.Config, client *iam.Client) {
	input := &iam.AddRoleToInstanceProfileInput{
		// InstanceProfileName: *string, // Required
		// RoleName: *string, // Required
	}

	if len(_iamInstanceProfileName) > 0 {
		input.InstanceProfileName = aws.String(_iamInstanceProfileName)
	}
	if len(_iamRoleName) > 0 {
		input.RoleName = aws.String(_iamRoleName)
	}

	if resp, err := client.AddRoleToInstanceProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds the specified user to the specified group.
func iam_AddUserToGroup(cfg aws.Config, client *iam.Client) {
	input := &iam.AddUserToGroupInput{
		// GroupName: *string, // Required
		// UserName: *string, // Required
	}

	if len(_iamGroupName) > 0 {
		input.GroupName = aws.String(_iamGroupName)
	}
	if len(_iamUserName) > 0 {
		input.UserName = aws.String(_iamUserName)
	}

	if resp, err := client.AddUserToGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates a delegation request with the current identity.
// If the partner that created the delegation request has specified the owner
// account during creation, only an identity from that owner account can call the
// AssociateDelegationRequest API for the specified delegation request. Once the
// AssociateDelegationRequest API call is successful, the ARN of the current
// calling identity will be stored as the ownerId of the request.
//
// If the partner that created the delegation request has not specified the owner
// account during creation, any caller from any account can call the
// AssociateDelegationRequest API for the delegation request. Once this API call is
// successful, the ARN of the current calling identity will be stored as the
// ownerId and the Amazon Web Services account ID of the current calling identity
// will be stored as the ownerAccount of the request.
//
// For more details, see [Managing Permissions for Delegation Requests].
//
// [Managing Permissions for Delegation Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies-temporary-delegation.html#temporary-delegation-managing-permissions
func iam_AssociateDelegationRequest(cfg aws.Config, client *iam.Client) {
	input := &iam.AssociateDelegationRequestInput{
		// DelegationRequestId: *string, // Required
	}

	if len(_iamDelegationRequestId) > 0 {
		input.DelegationRequestId = aws.String(_iamDelegationRequestId)
	}

	if resp, err := client.AssociateDelegationRequest(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Attaches the specified managed policy to the specified IAM group.
// You use this operation to attach a managed policy to a group. To embed an
// inline policy in a group, use [PutGroupPolicy]PutGroupPolicy .
//
// As a best practice, you can validate your IAM policies. To learn more, see [Validating IAM policies] in
// the IAM User Guide.
//
// For more information about policies, see [Managed policies and inline policies] in the IAM User Guide.
//
// [PutGroupPolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_PutGroupPolicy.html
// [Validating IAM policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_policy-validator.html
// [Managed policies and inline policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html
func iam_AttachGroupPolicy(cfg aws.Config, client *iam.Client) {
	input := &iam.AttachGroupPolicyInput{
		// GroupName: *string, // Required
		// PolicyArn: *string, // Required
	}

	if len(_iamGroupName) > 0 {
		input.GroupName = aws.String(_iamGroupName)
	}
	if len(_iamPolicyArn) > 0 {
		input.PolicyArn = aws.String(_iamPolicyArn)
	}

	if resp, err := client.AttachGroupPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Attaches the specified managed policy to the specified IAM role. When you
// attach a managed policy to a role, the managed policy becomes part of the role's
// permission (access) policy.
//
// You cannot use a managed policy as the role's trust policy. The role's trust
// policy is created at the same time as the role, using [CreateRole]CreateRole . You can
// update a role's trust policy using [UpdateAssumerolePolicy]UpdateAssumerolePolicy .
//
// Use this operation to attach a managed policy to a role. To embed an inline
// policy in a role, use [PutRolePolicy]PutRolePolicy . For more information about policies, see [Managed policies and inline policies]
// in the IAM User Guide.
//
// As a best practice, you can validate your IAM policies. To learn more, see [Validating IAM policies] in
// the IAM User Guide.
//
// [Validating IAM policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_policy-validator.html
// [UpdateAssumerolePolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_UpdateAssumeRolePolicy.html
// [PutRolePolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_PutRolePolicy.html
// [CreateRole]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_CreateRole.html
// [Managed policies and inline policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html
func iam_AttachRolePolicy(cfg aws.Config, client *iam.Client) {
	input := &iam.AttachRolePolicyInput{
		// PolicyArn: *string, // Required
		// RoleName: *string, // Required
	}

	if len(_iamPolicyArn) > 0 {
		input.PolicyArn = aws.String(_iamPolicyArn)
	}
	if len(_iamRoleName) > 0 {
		input.RoleName = aws.String(_iamRoleName)
	}

	if resp, err := client.AttachRolePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Attaches the specified managed policy to the specified user.
// You use this operation to attach a managed policy to a user. To embed an inline
// policy in a user, use [PutUserPolicy]PutUserPolicy .
//
// As a best practice, you can validate your IAM policies. To learn more, see [Validating IAM policies] in
// the IAM User Guide.
//
// For more information about policies, see [Managed policies and inline policies] in the IAM User Guide.
//
// [Validating IAM policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_policy-validator.html
// [PutUserPolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_PutUserPolicy.html
// [Managed policies and inline policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html
func iam_AttachUserPolicy(cfg aws.Config, client *iam.Client) {
	input := &iam.AttachUserPolicyInput{
		// PolicyArn: *string, // Required
		// UserName: *string, // Required
	}

	if len(_iamPolicyArn) > 0 {
		input.PolicyArn = aws.String(_iamPolicyArn)
	}
	if len(_iamUserName) > 0 {
		input.UserName = aws.String(_iamUserName)
	}

	if resp, err := client.AttachUserPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Changes the password of the IAM user who is calling this operation. This
// operation can be performed using the CLI, the Amazon Web Services API, or the My
// Security Credentials page in the Amazon Web Services Management Console. The
// Amazon Web Services account root user password is not affected by this
// operation.
//
// Use [UpdateLoginProfile] to use the CLI, the Amazon Web Services API, or the Users page in the IAM
// console to change the password for any IAM user. For more information about
// modifying passwords, see [Managing passwords]in the IAM User Guide.
//
// [UpdateLoginProfile]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_UpdateLoginProfile.html
// [Managing passwords]: https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_ManagingLogins.html
func iam_ChangePassword(cfg aws.Config, client *iam.Client) {
	input := &iam.ChangePasswordInput{
		// NewPassword: *string, // Required
		// OldPassword: *string, // Required
	}

	if len(_iamNewPassword) > 0 {
		input.NewPassword = aws.String(_iamNewPassword)
	}
	if len(_iamOldPassword) > 0 {
		input.OldPassword = aws.String(_iamOldPassword)
	}

	if resp, err := client.ChangePassword(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new Amazon Web Services secret access key and corresponding Amazon
// Web Services access key ID for the specified user. The default status for new
// keys is Active .
//
// If you do not specify a user name, IAM determines the user name implicitly
// based on the Amazon Web Services access key ID signing the request. This
// operation works for access keys under the Amazon Web Services account.
// Consequently, you can use this operation to manage Amazon Web Services account
// root user credentials. This is true even if the Amazon Web Services account has
// no associated users.
//
// For information about quotas on the number of keys you can create, see [IAM and STS quotas] in the
// IAM User Guide.
//
// To ensure the security of your Amazon Web Services account, the secret access
// key is accessible only during key and user creation. You must save the key (for
// example, in a text file) if you want to be able to access it again. If a secret
// key is lost, you can delete the access keys for the associated user and then
// create new keys.
//
// [IAM and STS quotas]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-quotas.html
func iam_CreateAccessKey(cfg aws.Config, client *iam.Client) {
	input := &iam.CreateAccessKeyInput{}

	if len(_iamUserName) > 0 {
		input.UserName = aws.String(_iamUserName)
	}

	if resp, err := client.CreateAccessKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an alias for your Amazon Web Services account. For information about
// using an Amazon Web Services account alias, see [Creating, deleting, and listing an Amazon Web Services account alias]in the Amazon Web Services
// Sign-In User Guide.
//
// [Creating, deleting, and listing an Amazon Web Services account alias]: https://docs.aws.amazon.com/signin/latest/userguide/CreateAccountAlias.html
func iam_CreateAccountAlias(cfg aws.Config, client *iam.Client) {
	input := &iam.CreateAccountAliasInput{
		// AccountAlias: *string, // Required
	}

	if len(_iamAccountAlias) > 0 {
		input.AccountAlias = aws.String(_iamAccountAlias)
	}

	if resp, err := client.CreateAccountAlias(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an IAM delegation request for temporary access delegation.
// This API is not available for general use. In order to use this API, a caller
// first need to go through an onboarding process described in the [partner onboarding documentation].
//
// [partner onboarding documentation]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies-temporary-delegation-partner-guide.html
func iam_CreateDelegationRequest(cfg aws.Config, client *iam.Client) {
	input := &iam.CreateDelegationRequestInput{
		// Description: *string, // Required
		// NotificationChannel: *string, // Required
		// Permissions: *types.DelegationPermission, // Required
		// RequestorWorkflowId: *string, // Required
		// SessionDuration: *int32, // Required
	}

	if len(_iamDescription) > 0 {
		input.Description = aws.String(_iamDescription)
	}
	if len(_iamNotificationChannel) > 0 {
		input.NotificationChannel = aws.String(_iamNotificationChannel)
	}
	if len(_iamPermissions) > 0 {
		if err := assignInputField(input, "Permissions", _iamPermissions); err != nil {
			log.Errorf("invalid --permissions: %s", err.Error())
			return
		}
	}
	if len(_iamRequestorWorkflowId) > 0 {
		input.RequestorWorkflowId = aws.String(_iamRequestorWorkflowId)
	}
	if len(_iamSessionDuration) > 0 {
		if err := assignInputField(input, "SessionDuration", _iamSessionDuration); err != nil {
			log.Errorf("invalid --session-duration: %s", err.Error())
			return
		}
	}
	if len(_iamOnlySendByOwner) > 0 {
		if err := assignInputField(input, "OnlySendByOwner", _iamOnlySendByOwner); err != nil {
			log.Errorf("invalid --only-send-by-owner: %s", err.Error())
			return
		}
	}
	if len(_iamOwnerAccountId) > 0 {
		input.OwnerAccountId = aws.String(_iamOwnerAccountId)
	}
	if len(_iamRedirectUrl) > 0 {
		input.RedirectUrl = aws.String(_iamRedirectUrl)
	}
	if len(_iamRequestMessage) > 0 {
		input.RequestMessage = aws.String(_iamRequestMessage)
	}

	if resp, err := client.CreateDelegationRequest(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new group.
// For information about the number of groups you can create, see [IAM and STS quotas] in the IAM User
// Guide.
//
// [IAM and STS quotas]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-quotas.html
func iam_CreateGroup(cfg aws.Config, client *iam.Client) {
	input := &iam.CreateGroupInput{
		// GroupName: *string, // Required
	}

	if len(_iamGroupName) > 0 {
		input.GroupName = aws.String(_iamGroupName)
	}
	if len(_iamPath) > 0 {
		input.Path = aws.String(_iamPath)
	}

	if resp, err := client.CreateGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new instance profile. For information about instance profiles, see [Using roles for applications on Amazon EC2]
// in the IAM User Guide, and [Instance profiles]in the Amazon EC2 User Guide.
//
// For information about the number of instance profiles you can create, see [IAM object quotas] in
// the IAM User Guide.
//
// [Instance profiles]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/iam-roles-for-amazon-ec2.html#ec2-instance-profile
// [IAM object quotas]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-quotas.html
// [Using roles for applications on Amazon EC2]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use_switch-role-ec2.html
func iam_CreateInstanceProfile(cfg aws.Config, client *iam.Client) {
	input := &iam.CreateInstanceProfileInput{
		// InstanceProfileName: *string, // Required
	}

	if len(_iamInstanceProfileName) > 0 {
		input.InstanceProfileName = aws.String(_iamInstanceProfileName)
	}
	if len(_iamPath) > 0 {
		input.Path = aws.String(_iamPath)
	}
	if len(_iamTags) > 0 {
		if err := assignInputField(input, "Tags", _iamTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateInstanceProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a password for the specified IAM user. A password allows an IAM user to
// access Amazon Web Services services through the Amazon Web Services Management
// Console.
//
// You can use the CLI, the Amazon Web Services API, or the Users page in the IAM
// console to create a password for any IAM user. Use [ChangePassword]to update your own existing
// password in the My Security Credentials page in the Amazon Web Services
// Management Console.
//
// For more information about managing passwords, see [Managing passwords] in the IAM User Guide.
//
// [ChangePassword]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_ChangePassword.html
// [Managing passwords]: https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_ManagingLogins.html
func iam_CreateLoginProfile(cfg aws.Config, client *iam.Client) {
	input := &iam.CreateLoginProfileInput{}

	if len(_iamPassword) > 0 {
		input.Password = aws.String(_iamPassword)
	}
	if len(_iamPasswordResetRequired) > 0 {
		if err := assignInputField(input, "PasswordResetRequired", _iamPasswordResetRequired); err != nil {
			log.Errorf("invalid --password-reset-required: %s", err.Error())
			return
		}
	}
	if len(_iamUserName) > 0 {
		input.UserName = aws.String(_iamUserName)
	}

	if resp, err := client.CreateLoginProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an IAM entity to describe an identity provider (IdP) that supports [OpenID Connect (OIDC)].
// The OIDC provider that you create with this operation can be used as a
// principal in a role's trust policy. Such a policy establishes a trust
// relationship between Amazon Web Services and the OIDC provider.
//
// If you are using an OIDC identity provider from Google, Facebook, or Amazon
// Cognito, you don't need to create a separate IAM identity provider. These OIDC
// identity providers are already built-in to Amazon Web Services and are available
// for your use. Instead, you can move directly to creating new roles using your
// identity provider. To learn more, see [Creating a role for web identity or OpenID connect federation]in the IAM User Guide.
//
// When you create the IAM OIDC provider, you specify the following:
//
// - The URL of the OIDC identity provider (IdP) to trust
//
// - A list of client IDs (also known as audiences) that identify the
// application or applications allowed to authenticate using the OIDC provider
//
// - A list of tags that are attached to the specified IAM OIDC provider
//
// - A list of thumbprints of one or more server certificates that the IdP uses
//
// You get all of this information from the OIDC IdP you want to use to access
// Amazon Web Services.
//
// Amazon Web Services secures communication with OIDC identity providers (IdPs)
// using our library of trusted root certificate authorities (CAs) to verify the
// JSON Web Key Set (JWKS) endpoint's TLS certificate. If your OIDC IdP relies on a
// certificate that is not signed by one of these trusted CAs, only then we secure
// communication using the thumbprints set in the IdP's configuration.
//
// The trust for the OIDC provider is derived from the IAM provider that this
// operation creates. Therefore, it is best to limit access to the [CreateOpenIDConnectProvider]operation to
// highly privileged users.
//
// [Creating a role for web identity or OpenID connect federation]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_create_for-idp_oidc.html
// [CreateOpenIDConnectProvider]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_CreateOpenIDConnectProvider.html
// [OpenID Connect (OIDC)]: http://openid.net/connect/
func iam_CreateOpenIDConnectProvider(cfg aws.Config, client *iam.Client) {
	input := &iam.CreateOpenIDConnectProviderInput{
		// Url: *string, // Required
	}

	if len(_iamUrl) > 0 {
		input.Url = aws.String(_iamUrl)
	}
	if len(_iamClientIDList) > 0 {
		input.ClientIDList = append([]string(nil), _iamClientIDList...)
	}
	if len(_iamTags) > 0 {
		if err := assignInputField(input, "Tags", _iamTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_iamThumbprintList) > 0 {
		input.ThumbprintList = append([]string(nil), _iamThumbprintList...)
	}

	if resp, err := client.CreateOpenIDConnectProvider(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new managed policy for your Amazon Web Services account.
// This operation creates a policy version with a version identifier of v1 and
// sets v1 as the policy's default version. For more information about policy
// versions, see [Versioning for managed policies]in the IAM User Guide.
//
// As a best practice, you can validate your IAM policies. To learn more, see [Validating IAM policies] in
// the IAM User Guide.
//
// For more information about managed policies in general, see [Managed policies and inline policies] in the IAM User
// Guide.
//
// [Validating IAM policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_policy-validator.html
// [Versioning for managed policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-versions.html
// [Managed policies and inline policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html
func iam_CreatePolicy(cfg aws.Config, client *iam.Client) {
	input := &iam.CreatePolicyInput{
		// PolicyDocument: *string, // Required
		// PolicyName: *string, // Required
	}

	if len(_iamPolicyDocument) > 0 {
		input.PolicyDocument = aws.String(_iamPolicyDocument)
	}
	if len(_iamPolicyName) > 0 {
		input.PolicyName = aws.String(_iamPolicyName)
	}
	if len(_iamDescription) > 0 {
		input.Description = aws.String(_iamDescription)
	}
	if len(_iamPath) > 0 {
		input.Path = aws.String(_iamPath)
	}
	if len(_iamTags) > 0 {
		if err := assignInputField(input, "Tags", _iamTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreatePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new version of the specified managed policy. To update a managed
// policy, you create a new policy version. A managed policy can have up to five
// versions. If the policy has five versions, you must delete an existing version
// using [DeletePolicyVersion]before you create a new version.
//
// Optionally, you can set the new version as the policy's default version. The
// default version is the version that is in effect for the IAM users, groups, and
// roles to which the policy is attached.
//
// For more information about managed policy versions, see [Versioning for managed policies] in the IAM User Guide.
//
// [DeletePolicyVersion]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_DeletePolicyVersion.html
// [Versioning for managed policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-versions.html
func iam_CreatePolicyVersion(cfg aws.Config, client *iam.Client) {
	input := &iam.CreatePolicyVersionInput{
		// PolicyArn: *string, // Required
		// PolicyDocument: *string, // Required
	}

	if len(_iamPolicyArn) > 0 {
		input.PolicyArn = aws.String(_iamPolicyArn)
	}
	if len(_iamPolicyDocument) > 0 {
		input.PolicyDocument = aws.String(_iamPolicyDocument)
	}
	if len(_iamSetAsDefault) > 0 {
		if err := assignInputField(input, "SetAsDefault", _iamSetAsDefault); err != nil {
			log.Errorf("invalid --set-as-default: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreatePolicyVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new role for your Amazon Web Services account.
// For more information about roles, see [IAM roles] in the IAM User Guide. For information
// about quotas for role names and the number of roles you can create, see [IAM and STS quotas]in the
// IAM User Guide.
//
// [IAM and STS quotas]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-quotas.html
// [IAM roles]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles.html
func iam_CreateRole(cfg aws.Config, client *iam.Client) {
	input := &iam.CreateRoleInput{
		// AssumeRolePolicyDocument: *string, // Required
		// RoleName: *string, // Required
	}

	if len(_iamAssumeRolePolicyDocument) > 0 {
		input.AssumeRolePolicyDocument = aws.String(_iamAssumeRolePolicyDocument)
	}
	if len(_iamRoleName) > 0 {
		input.RoleName = aws.String(_iamRoleName)
	}
	if len(_iamDescription) > 0 {
		input.Description = aws.String(_iamDescription)
	}
	if len(_iamMaxSessionDuration) > 0 {
		if err := assignInputField(input, "MaxSessionDuration", _iamMaxSessionDuration); err != nil {
			log.Errorf("invalid --max-session-duration: %s", err.Error())
			return
		}
	}
	if len(_iamPath) > 0 {
		input.Path = aws.String(_iamPath)
	}
	if len(_iamPermissionsBoundary) > 0 {
		input.PermissionsBoundary = aws.String(_iamPermissionsBoundary)
	}
	if len(_iamTags) > 0 {
		if err := assignInputField(input, "Tags", _iamTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateRole(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an IAM resource that describes an identity provider (IdP) that supports
// SAML 2.0.
//
// The SAML provider resource that you create with this operation can be used as a
// principal in an IAM role's trust policy. Such a policy can enable federated
// users who sign in using the SAML IdP to assume the role. You can create an IAM
// role that supports Web-based single sign-on (SSO) to the Amazon Web Services
// Management Console or one that supports API access to Amazon Web Services.
//
// When you create the SAML provider resource, you upload a SAML metadata document
// that you get from your IdP. That document includes the issuer's name, expiration
// information, and keys that can be used to validate the SAML authentication
// response (assertions) that the IdP sends. You must generate the metadata
// document using the identity management software that is used as your
// organization's IdP.
//
// This operation requires [Signature Version 4].
//
// For more information, see [Enabling SAML 2.0 federated users to access the Amazon Web Services Management Console] and [About SAML 2.0-based federation] in the IAM User Guide.
//
// [Signature Version 4]: https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html
// [About SAML 2.0-based federation]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_providers_saml.html
// [Enabling SAML 2.0 federated users to access the Amazon Web Services Management Console]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_providers_enable-console-saml.html
func iam_CreateSAMLProvider(cfg aws.Config, client *iam.Client) {
	input := &iam.CreateSAMLProviderInput{
		// Name: *string, // Required
		// SAMLMetadataDocument: *string, // Required
	}

	if len(_iamName) > 0 {
		input.Name = aws.String(_iamName)
	}
	if len(_iamSAMLMetadataDocument) > 0 {
		input.SAMLMetadataDocument = aws.String(_iamSAMLMetadataDocument)
	}
	if len(_iamAddPrivateKey) > 0 {
		input.AddPrivateKey = aws.String(_iamAddPrivateKey)
	}
	if len(_iamAssertionEncryptionMode) > 0 {
		if err := assignInputField(input, "AssertionEncryptionMode", _iamAssertionEncryptionMode); err != nil {
			log.Errorf("invalid --assertion-encryption-mode: %s", err.Error())
			return
		}
	}
	if len(_iamTags) > 0 {
		if err := assignInputField(input, "Tags", _iamTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSAMLProvider(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an IAM role that is linked to a specific Amazon Web Services service.
// The service controls the attached policies and when the role can be deleted.
// This helps ensure that the service is not broken by an unexpectedly changed or
// deleted role, which could put your Amazon Web Services resources into an unknown
// state. Allowing the service to control the role helps improve service stability
// and proper cleanup when a service and its role are no longer needed. For more
// information, see [Using service-linked roles]in the IAM User Guide.
//
// To attach a policy to this service-linked role, you must make the request using
// the Amazon Web Services service that depends on this role.
//
// [Using service-linked roles]: https://docs.aws.amazon.com/IAM/latest/UserGuide/using-service-linked-roles.html
func iam_CreateServiceLinkedRole(cfg aws.Config, client *iam.Client) {
	input := &iam.CreateServiceLinkedRoleInput{
		// AWSServiceName: *string, // Required
	}

	if len(_iamAWSServiceName) > 0 {
		input.AWSServiceName = aws.String(_iamAWSServiceName)
	}
	if len(_iamCustomSuffix) > 0 {
		input.CustomSuffix = aws.String(_iamCustomSuffix)
	}
	if len(_iamDescription) > 0 {
		input.Description = aws.String(_iamDescription)
	}

	if resp, err := client.CreateServiceLinkedRole(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Generates a set of credentials consisting of a user name and password that can
// be used to access the service specified in the request. These credentials are
// generated by IAM, and can be used only for the specified service.
//
// You can have a maximum of two sets of service-specific credentials for each
// supported service per user.
//
// You can create service-specific credentials for Amazon Bedrock, CodeCommit and
// Amazon Keyspaces (for Apache Cassandra).
//
// You can reset the password to a new service-generated value by calling [ResetServiceSpecificCredential].
//
// For more information about service-specific credentials, see [Service-specific credentials for IAM users] in the IAM User
// Guide.
//
// [ResetServiceSpecificCredential]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_ResetServiceSpecificCredential.html
// [Service-specific credentials for IAM users]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_bedrock.html
func iam_CreateServiceSpecificCredential(cfg aws.Config, client *iam.Client) {
	input := &iam.CreateServiceSpecificCredentialInput{
		// ServiceName: *string, // Required
		// UserName: *string, // Required
	}

	if len(_iamServiceName) > 0 {
		input.ServiceName = aws.String(_iamServiceName)
	}
	if len(_iamUserName) > 0 {
		input.UserName = aws.String(_iamUserName)
	}
	if len(_iamCredentialAgeDays) > 0 {
		if err := assignInputField(input, "CredentialAgeDays", _iamCredentialAgeDays); err != nil {
			log.Errorf("invalid --credential-age-days: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateServiceSpecificCredential(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new IAM user for your Amazon Web Services account.
// For information about quotas for the number of IAM users you can create, see [IAM and STS quotas]
// in the IAM User Guide.
//
// [IAM and STS quotas]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-quotas.html
func iam_CreateUser(cfg aws.Config, client *iam.Client) {
	input := &iam.CreateUserInput{
		// UserName: *string, // Required
	}

	if len(_iamUserName) > 0 {
		input.UserName = aws.String(_iamUserName)
	}
	if len(_iamPath) > 0 {
		input.Path = aws.String(_iamPath)
	}
	if len(_iamPermissionsBoundary) > 0 {
		input.PermissionsBoundary = aws.String(_iamPermissionsBoundary)
	}
	if len(_iamTags) > 0 {
		if err := assignInputField(input, "Tags", _iamTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
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

// Creates a new virtual MFA device for the Amazon Web Services account. After
// creating the virtual MFA, use [EnableMFADevice]to attach the MFA device to an IAM user. For more
// information about creating and working with virtual MFA devices, see [Using a virtual MFA device]in the IAM
// User Guide.
//
// For information about the maximum number of MFA devices you can create, see [IAM and STS quotas] in
// the IAM User Guide.
//
// The seed information contained in the QR code and the Base32 string should be
// treated like any other secret access information. In other words, protect the
// seed information as you would your Amazon Web Services access keys or your
// passwords. After you provision your virtual device, you should ensure that the
// information is destroyed following secure procedures.
//
// [Using a virtual MFA device]: https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_VirtualMFA.html
// [EnableMFADevice]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_EnableMFADevice.html
// [IAM and STS quotas]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-quotas.html
func iam_CreateVirtualMFADevice(cfg aws.Config, client *iam.Client) {
	input := &iam.CreateVirtualMFADeviceInput{
		// VirtualMFADeviceName: *string, // Required
	}

	if len(_iamVirtualMFADeviceName) > 0 {
		input.VirtualMFADeviceName = aws.String(_iamVirtualMFADeviceName)
	}
	if len(_iamPath) > 0 {
		input.Path = aws.String(_iamPath)
	}
	if len(_iamTags) > 0 {
		if err := assignInputField(input, "Tags", _iamTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateVirtualMFADevice(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deactivates the specified MFA device and removes it from association with the
// user name for which it was originally enabled.
//
// For more information about creating and working with virtual MFA devices, see [Enabling a virtual multi-factor authentication (MFA) device]
// in the IAM User Guide.
//
// [Enabling a virtual multi-factor authentication (MFA) device]: https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_VirtualMFA.html
func iam_DeactivateMFADevice(cfg aws.Config, client *iam.Client) {
	input := &iam.DeactivateMFADeviceInput{
		// SerialNumber: *string, // Required
	}

	if len(_iamSerialNumber) > 0 {
		input.SerialNumber = aws.String(_iamSerialNumber)
	}
	if len(_iamUserName) > 0 {
		input.UserName = aws.String(_iamUserName)
	}

	if resp, err := client.DeactivateMFADevice(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the access key pair associated with the specified IAM user.
// If you do not specify a user name, IAM determines the user name implicitly
// based on the Amazon Web Services access key ID signing the request. This
// operation works for access keys under the Amazon Web Services account.
// Consequently, you can use this operation to manage Amazon Web Services account
// root user credentials even if the Amazon Web Services account has no associated
// users.
func iam_DeleteAccessKey(cfg aws.Config, client *iam.Client) {
	input := &iam.DeleteAccessKeyInput{
		// AccessKeyId: *string, // Required
	}

	if len(_iamAccessKeyId) > 0 {
		input.AccessKeyId = aws.String(_iamAccessKeyId)
	}
	if len(_iamUserName) > 0 {
		input.UserName = aws.String(_iamUserName)
	}

	if resp, err := client.DeleteAccessKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified Amazon Web Services account alias. For information about
// using an Amazon Web Services account alias, see [Creating, deleting, and listing an Amazon Web Services account alias]in the Amazon Web Services
// Sign-In User Guide.
//
// [Creating, deleting, and listing an Amazon Web Services account alias]: https://docs.aws.amazon.com/signin/latest/userguide/CreateAccountAlias.html
func iam_DeleteAccountAlias(cfg aws.Config, client *iam.Client) {
	input := &iam.DeleteAccountAliasInput{
		// AccountAlias: *string, // Required
	}

	if len(_iamAccountAlias) > 0 {
		input.AccountAlias = aws.String(_iamAccountAlias)
	}

	if resp, err := client.DeleteAccountAlias(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the password policy for the Amazon Web Services account. There are no
// parameters.
func iam_DeleteAccountPasswordPolicy(cfg aws.Config, client *iam.Client) {
	input := &iam.DeleteAccountPasswordPolicyInput{}

	if resp, err := client.DeleteAccountPasswordPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified IAM group. The group must not contain any users or have
// any attached policies.
func iam_DeleteGroup(cfg aws.Config, client *iam.Client) {
	input := &iam.DeleteGroupInput{
		// GroupName: *string, // Required
	}

	if len(_iamGroupName) > 0 {
		input.GroupName = aws.String(_iamGroupName)
	}

	if resp, err := client.DeleteGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified inline policy that is embedded in the specified IAM group.
// A group can also have managed policies attached to it. To detach a managed
// policy from a group, use [DetachGroupPolicy]. For more information about policies, refer to [Managed policies and inline policies] in
// the IAM User Guide.
//
// [DetachGroupPolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_DetachGroupPolicy.html
// [Managed policies and inline policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html
func iam_DeleteGroupPolicy(cfg aws.Config, client *iam.Client) {
	input := &iam.DeleteGroupPolicyInput{
		// GroupName: *string, // Required
		// PolicyName: *string, // Required
	}

	if len(_iamGroupName) > 0 {
		input.GroupName = aws.String(_iamGroupName)
	}
	if len(_iamPolicyName) > 0 {
		input.PolicyName = aws.String(_iamPolicyName)
	}

	if resp, err := client.DeleteGroupPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified instance profile. The instance profile must not have an
// associated role.
//
// Make sure that you do not have any Amazon EC2 instances running with the
// instance profile you are about to delete. Deleting a role or instance profile
// that is associated with a running instance will break any applications running
// on the instance.
//
// For more information about instance profiles, see [Using instance profiles] in the IAM User Guide.
//
// [Using instance profiles]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use_switch-role-ec2_instance-profiles.html
func iam_DeleteInstanceProfile(cfg aws.Config, client *iam.Client) {
	input := &iam.DeleteInstanceProfileInput{
		// InstanceProfileName: *string, // Required
	}

	if len(_iamInstanceProfileName) > 0 {
		input.InstanceProfileName = aws.String(_iamInstanceProfileName)
	}

	if resp, err := client.DeleteInstanceProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the password for the specified IAM user or root user, For more
// information, see [Managing passwords for IAM users].
//
// You can use the CLI, the Amazon Web Services API, or the Users page in the IAM
// console to delete a password for any IAM user. You can use [ChangePassword]to update, but not
// delete, your own password in the My Security Credentials page in the Amazon Web
// Services Management Console.
//
// Deleting a user's password does not prevent a user from accessing Amazon Web
// Services through the command line interface or the API. To prevent all user
// access, you must also either make any access keys inactive or delete them. For
// more information about making keys inactive or deleting them, see [UpdateAccessKey]and [DeleteAccessKey].
//
// [ChangePassword]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_ChangePassword.html
// [DeleteAccessKey]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_DeleteAccessKey.html
// [Managing passwords for IAM users]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_passwords_admin-change-user.html
// [UpdateAccessKey]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_UpdateAccessKey.html
func iam_DeleteLoginProfile(cfg aws.Config, client *iam.Client) {
	input := &iam.DeleteLoginProfileInput{}

	if len(_iamUserName) > 0 {
		input.UserName = aws.String(_iamUserName)
	}

	if resp, err := client.DeleteLoginProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an OpenID Connect identity provider (IdP) resource object in IAM.
// Deleting an IAM OIDC provider resource does not update any roles that reference
// the provider as a principal in their trust policies. Any attempt to assume a
// role that references a deleted provider fails.
//
// This operation is idempotent; it does not fail or return an error if you call
// the operation for a provider that does not exist.
func iam_DeleteOpenIDConnectProvider(cfg aws.Config, client *iam.Client) {
	input := &iam.DeleteOpenIDConnectProviderInput{
		// OpenIDConnectProviderArn: *string, // Required
	}

	if len(_iamOpenIDConnectProviderArn) > 0 {
		input.OpenIDConnectProviderArn = aws.String(_iamOpenIDConnectProviderArn)
	}

	if resp, err := client.DeleteOpenIDConnectProvider(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified managed policy.
// Before you can delete a managed policy, you must first detach the policy from
// all users, groups, and roles that it is attached to. In addition, you must
// delete all the policy's versions. The following steps describe the process for
// deleting a managed policy:
//
// - Detach the policy from all users, groups, and roles that the policy is
// attached to, using [DetachUserPolicy], [DetachGroupPolicy], or [DetachRolePolicy]. To list all the users, groups, and roles that a
// policy is attached to, use [ListEntitiesForPolicy].
//
// - Delete all versions of the policy using [DeletePolicyVersion]. To list the policy's versions,
// use [ListPolicyVersions]. You cannot use [DeletePolicyVersion]to delete the version that is marked as the default
// version. You delete the policy's default version in the next step of the
// process.
//
// - Delete the policy (this automatically deletes the policy's default version)
// using this operation.
//
// For information about managed policies, see [Managed policies and inline policies] in the IAM User Guide.
//
// [DetachUserPolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_DetachUserPolicy.html
// [DetachRolePolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_DetachRolePolicy.html
// [ListEntitiesForPolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_ListEntitiesForPolicy.html
// [DeletePolicyVersion]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_DeletePolicyVersion.html
// [DetachGroupPolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_DetachGroupPolicy.html
// [ListPolicyVersions]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_ListPolicyVersions.html
// [Managed policies and inline policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html
func iam_DeletePolicy(cfg aws.Config, client *iam.Client) {
	input := &iam.DeletePolicyInput{
		// PolicyArn: *string, // Required
	}

	if len(_iamPolicyArn) > 0 {
		input.PolicyArn = aws.String(_iamPolicyArn)
	}

	if resp, err := client.DeletePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified version from the specified managed policy.
// You cannot delete the default version from a policy using this operation. To
// delete the default version from a policy, use [DeletePolicy]. To find out which version of a
// policy is marked as the default version, use [ListPolicyVersions].
//
// For information about versions for managed policies, see [Versioning for managed policies] in the IAM User Guide.
//
// [DeletePolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_DeletePolicy.html
// [Versioning for managed policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-versions.html
// [ListPolicyVersions]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_ListPolicyVersions.html
func iam_DeletePolicyVersion(cfg aws.Config, client *iam.Client) {
	input := &iam.DeletePolicyVersionInput{
		// PolicyArn: *string, // Required
		// VersionId: *string, // Required
	}

	if len(_iamPolicyArn) > 0 {
		input.PolicyArn = aws.String(_iamPolicyArn)
	}
	if len(_iamVersionId) > 0 {
		input.VersionId = aws.String(_iamVersionId)
	}

	if resp, err := client.DeletePolicyVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified role. Unlike the Amazon Web Services Management Console,
// when you delete a role programmatically, you must delete the items attached to
// the role manually, or the deletion fails. For more information, see [Deleting an IAM role]. Before
// attempting to delete a role, remove the following attached items:
//
// - Inline policies ([DeleteRolePolicy] )
//
// - Attached managed policies ([DetachRolePolicy] )
//
// - Instance profile ([RemoveRoleFromInstanceProfile] )
//
// - Optional – Delete instance profile after detaching from role for resource
// clean up ([DeleteInstanceProfile] )
//
// Make sure that you do not have any Amazon EC2 instances running with the role
// you are about to delete. Deleting a role or instance profile that is associated
// with a running instance will break any applications running on the instance.
//
// [DetachRolePolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_DetachRolePolicy.html
// [RemoveRoleFromInstanceProfile]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_RemoveRoleFromInstanceProfile.html
// [DeleteRolePolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_DeleteRolePolicy.html
// [DeleteInstanceProfile]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_DeleteInstanceProfile.html
// [Deleting an IAM role]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_manage_delete.html#roles-managingrole-deleting-cli
func iam_DeleteRole(cfg aws.Config, client *iam.Client) {
	input := &iam.DeleteRoleInput{
		// RoleName: *string, // Required
	}

	if len(_iamRoleName) > 0 {
		input.RoleName = aws.String(_iamRoleName)
	}

	if resp, err := client.DeleteRole(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the permissions boundary for the specified IAM role.
// You cannot set the boundary for a service-linked role.
//
// Deleting the permissions boundary for a role might increase its permissions.
// For example, it might allow anyone who assumes the role to perform all the
// actions granted in its permissions policies.
func iam_DeleteRolePermissionsBoundary(cfg aws.Config, client *iam.Client) {
	input := &iam.DeleteRolePermissionsBoundaryInput{
		// RoleName: *string, // Required
	}

	if len(_iamRoleName) > 0 {
		input.RoleName = aws.String(_iamRoleName)
	}

	if resp, err := client.DeleteRolePermissionsBoundary(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified inline policy that is embedded in the specified IAM role.
// A role can also have managed policies attached to it. To detach a managed
// policy from a role, use [DetachRolePolicy]. For more information about policies, refer to [Managed policies and inline policies] in the
// IAM User Guide.
//
// [DetachRolePolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_DetachRolePolicy.html
// [Managed policies and inline policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html
func iam_DeleteRolePolicy(cfg aws.Config, client *iam.Client) {
	input := &iam.DeleteRolePolicyInput{
		// PolicyName: *string, // Required
		// RoleName: *string, // Required
	}

	if len(_iamPolicyName) > 0 {
		input.PolicyName = aws.String(_iamPolicyName)
	}
	if len(_iamRoleName) > 0 {
		input.RoleName = aws.String(_iamRoleName)
	}

	if resp, err := client.DeleteRolePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a SAML provider resource in IAM.
// Deleting the provider resource from IAM does not update any roles that
// reference the SAML provider resource's ARN as a principal in their trust
// policies. Any attempt to assume a role that references a non-existent provider
// resource ARN fails.
//
// This operation requires [Signature Version 4].
//
// [Signature Version 4]: https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html
func iam_DeleteSAMLProvider(cfg aws.Config, client *iam.Client) {
	input := &iam.DeleteSAMLProviderInput{
		// SAMLProviderArn: *string, // Required
	}

	if len(_iamSAMLProviderArn) > 0 {
		input.SAMLProviderArn = aws.String(_iamSAMLProviderArn)
	}

	if resp, err := client.DeleteSAMLProvider(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified server certificate.
// For more information about working with server certificates, see [Working with server certificates] in the IAM
// User Guide. This topic also includes a list of Amazon Web Services services that
// can use the server certificates that you manage with IAM.
//
// If you are using a server certificate with Elastic Load Balancing, deleting the
// certificate could have implications for your application. If Elastic Load
// Balancing doesn't detect the deletion of bound certificates, it may continue to
// use the certificates. This could cause Elastic Load Balancing to stop accepting
// traffic. We recommend that you remove the reference to the certificate from
// Elastic Load Balancing before using this command to delete the certificate. For
// more information, see [DeleteLoadBalancerListeners]in the Elastic Load Balancing API Reference.
//
// [Working with server certificates]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_server-certs.html
// [DeleteLoadBalancerListeners]: https://docs.aws.amazon.com/ElasticLoadBalancing/latest/APIReference/API_DeleteLoadBalancerListeners.html
func iam_DeleteServerCertificate(cfg aws.Config, client *iam.Client) {
	input := &iam.DeleteServerCertificateInput{
		// ServerCertificateName: *string, // Required
	}

	if len(_iamServerCertificateName) > 0 {
		input.ServerCertificateName = aws.String(_iamServerCertificateName)
	}

	if resp, err := client.DeleteServerCertificate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Submits a service-linked role deletion request and returns a DeletionTaskId ,
// which you can use to check the status of the deletion. Before you call this
// operation, confirm that the role has no active sessions and that any resources
// used by the role in the linked service are deleted. If you call this operation
// more than once for the same service-linked role and an earlier deletion task is
// not complete, then the DeletionTaskId of the earlier request is returned.
//
// If you submit a deletion request for a service-linked role whose linked service
// is still accessing a resource, then the deletion task fails. If it fails, the [GetServiceLinkedRoleDeletionStatus]
// operation returns the reason for the failure, usually including the resources
// that must be deleted. To delete the service-linked role, you must first remove
// those resources from the linked service and then submit the deletion request
// again. Resources are specific to the service that is linked to the role. For
// more information about removing resources from a service, see the [Amazon Web Services documentation]for your
// service.
//
// For more information about service-linked roles, see [Roles terms and concepts: Amazon Web Services service-linked role] in the IAM User Guide.
//
// [Roles terms and concepts: Amazon Web Services service-linked role]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_terms-and-concepts.html#iam-term-service-linked-role
// [Amazon Web Services documentation]: http://docs.aws.amazon.com/
// [GetServiceLinkedRoleDeletionStatus]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetServiceLinkedRoleDeletionStatus.html
func iam_DeleteServiceLinkedRole(cfg aws.Config, client *iam.Client) {
	input := &iam.DeleteServiceLinkedRoleInput{
		// RoleName: *string, // Required
	}

	if len(_iamRoleName) > 0 {
		input.RoleName = aws.String(_iamRoleName)
	}

	if resp, err := client.DeleteServiceLinkedRole(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified service-specific credential.
func iam_DeleteServiceSpecificCredential(cfg aws.Config, client *iam.Client) {
	input := &iam.DeleteServiceSpecificCredentialInput{
		// ServiceSpecificCredentialId: *string, // Required
	}

	if len(_iamServiceSpecificCredentialId) > 0 {
		input.ServiceSpecificCredentialId = aws.String(_iamServiceSpecificCredentialId)
	}
	if len(_iamUserName) > 0 {
		input.UserName = aws.String(_iamUserName)
	}

	if resp, err := client.DeleteServiceSpecificCredential(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a signing certificate associated with the specified IAM user.
// If you do not specify a user name, IAM determines the user name implicitly
// based on the Amazon Web Services access key ID signing the request. This
// operation works for access keys under the Amazon Web Services account.
// Consequently, you can use this operation to manage Amazon Web Services account
// root user credentials even if the Amazon Web Services account has no associated
// IAM users.
func iam_DeleteSigningCertificate(cfg aws.Config, client *iam.Client) {
	input := &iam.DeleteSigningCertificateInput{
		// CertificateId: *string, // Required
	}

	if len(_iamCertificateId) > 0 {
		input.CertificateId = aws.String(_iamCertificateId)
	}
	if len(_iamUserName) > 0 {
		input.UserName = aws.String(_iamUserName)
	}

	if resp, err := client.DeleteSigningCertificate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified SSH public key.
// The SSH public key deleted by this operation is used only for authenticating
// the associated IAM user to an CodeCommit repository. For more information about
// using SSH keys to authenticate to an CodeCommit repository, see [Set up CodeCommit for SSH connections]in the
// CodeCommit User Guide.
//
// [Set up CodeCommit for SSH connections]: https://docs.aws.amazon.com/codecommit/latest/userguide/setting-up-credentials-ssh.html
func iam_DeleteSSHPublicKey(cfg aws.Config, client *iam.Client) {
	input := &iam.DeleteSSHPublicKeyInput{
		// SSHPublicKeyId: *string, // Required
		// UserName: *string, // Required
	}

	if len(_iamSSHPublicKeyId) > 0 {
		input.SSHPublicKeyId = aws.String(_iamSSHPublicKeyId)
	}
	if len(_iamUserName) > 0 {
		input.UserName = aws.String(_iamUserName)
	}

	if resp, err := client.DeleteSSHPublicKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified IAM user. Unlike the Amazon Web Services Management
// Console, when you delete a user programmatically, you must delete the items
// attached to the user manually, or the deletion fails. For more information, see [Deleting an IAM user]
// . Before attempting to delete a user, remove the following items:
//
// - Password ([DeleteLoginProfile] )
//
// - Access keys ([DeleteAccessKey] )
//
// - Signing certificate ([DeleteSigningCertificate] )
//
// - SSH public key ([DeleteSSHPublicKey] )
//
// - Git credentials ([DeleteServiceSpecificCredential] )
//
// - Multi-factor authentication (MFA) device ([DeactivateMFADevice] , [DeleteVirtualMFADevice])
//
// - Inline policies ([DeleteUserPolicy] )
//
// - Attached managed policies ([DetachUserPolicy] )
//
// - Group memberships ([RemoveUserFromGroup] )
//
// [DetachUserPolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_DetachUserPolicy.html
// [DeleteAccessKey]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_DeleteAccessKey.html
// [DeleteVirtualMFADevice]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_DeleteVirtualMFADevice.html
// [Deleting an IAM user]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_users_manage.html#id_users_deleting_cli
// [DeleteUserPolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_DeleteUserPolicy.html
// [RemoveUserFromGroup]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_RemoveUserFromGroup.html
// [DeleteLoginProfile]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_DeleteLoginProfile.html
// [DeleteServiceSpecificCredential]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_DeleteServiceSpecificCredential.html
// [DeleteSigningCertificate]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_DeleteSigningCertificate.html
// [DeleteSSHPublicKey]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_DeleteSSHPublicKey.html
// [DeactivateMFADevice]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_DeactivateMFADevice.html
func iam_DeleteUser(cfg aws.Config, client *iam.Client) {
	input := &iam.DeleteUserInput{
		// UserName: *string, // Required
	}

	if len(_iamUserName) > 0 {
		input.UserName = aws.String(_iamUserName)
	}

	if resp, err := client.DeleteUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the permissions boundary for the specified IAM user.
// Deleting the permissions boundary for a user might increase its permissions by
// allowing the user to perform all the actions granted in its permissions
// policies.
func iam_DeleteUserPermissionsBoundary(cfg aws.Config, client *iam.Client) {
	input := &iam.DeleteUserPermissionsBoundaryInput{
		// UserName: *string, // Required
	}

	if len(_iamUserName) > 0 {
		input.UserName = aws.String(_iamUserName)
	}

	if resp, err := client.DeleteUserPermissionsBoundary(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified inline policy that is embedded in the specified IAM user.
// A user can also have managed policies attached to it. To detach a managed
// policy from a user, use [DetachUserPolicy]. For more information about policies, refer to [Managed policies and inline policies] in the
// IAM User Guide.
//
// [DetachUserPolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_DetachUserPolicy.html
// [Managed policies and inline policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html
func iam_DeleteUserPolicy(cfg aws.Config, client *iam.Client) {
	input := &iam.DeleteUserPolicyInput{
		// PolicyName: *string, // Required
		// UserName: *string, // Required
	}

	if len(_iamPolicyName) > 0 {
		input.PolicyName = aws.String(_iamPolicyName)
	}
	if len(_iamUserName) > 0 {
		input.UserName = aws.String(_iamUserName)
	}

	if resp, err := client.DeleteUserPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a virtual MFA device.
// You must deactivate a user's virtual MFA device before you can delete it. For
// information about deactivating MFA devices, see [DeactivateMFADevice].
//
// [DeactivateMFADevice]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_DeactivateMFADevice.html
func iam_DeleteVirtualMFADevice(cfg aws.Config, client *iam.Client) {
	input := &iam.DeleteVirtualMFADeviceInput{
		// SerialNumber: *string, // Required
	}

	if len(_iamSerialNumber) > 0 {
		input.SerialNumber = aws.String(_iamSerialNumber)
	}

	if resp, err := client.DeleteVirtualMFADevice(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the specified managed policy from the specified IAM group.
// A group can also have inline policies embedded with it. To delete an inline
// policy, use [DeleteGroupPolicy]. For information about policies, see [Managed policies and inline policies] in the IAM User Guide.
//
// [DeleteGroupPolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_DeleteGroupPolicy.html
// [Managed policies and inline policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html
func iam_DetachGroupPolicy(cfg aws.Config, client *iam.Client) {
	input := &iam.DetachGroupPolicyInput{
		// GroupName: *string, // Required
		// PolicyArn: *string, // Required
	}

	if len(_iamGroupName) > 0 {
		input.GroupName = aws.String(_iamGroupName)
	}
	if len(_iamPolicyArn) > 0 {
		input.PolicyArn = aws.String(_iamPolicyArn)
	}

	if resp, err := client.DetachGroupPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the specified managed policy from the specified role.
// A role can also have inline policies embedded with it. To delete an inline
// policy, use [DeleteRolePolicy]. For information about policies, see [Managed policies and inline policies] in the IAM User Guide.
//
// [DeleteRolePolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_DeleteRolePolicy.html
// [Managed policies and inline policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html
func iam_DetachRolePolicy(cfg aws.Config, client *iam.Client) {
	input := &iam.DetachRolePolicyInput{
		// PolicyArn: *string, // Required
		// RoleName: *string, // Required
	}

	if len(_iamPolicyArn) > 0 {
		input.PolicyArn = aws.String(_iamPolicyArn)
	}
	if len(_iamRoleName) > 0 {
		input.RoleName = aws.String(_iamRoleName)
	}

	if resp, err := client.DetachRolePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the specified managed policy from the specified user.
// A user can also have inline policies embedded with it. To delete an inline
// policy, use [DeleteUserPolicy]. For information about policies, see [Managed policies and inline policies] in the IAM User Guide.
//
// [DeleteUserPolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_DeleteUserPolicy.html
// [Managed policies and inline policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html
func iam_DetachUserPolicy(cfg aws.Config, client *iam.Client) {
	input := &iam.DetachUserPolicyInput{
		// PolicyArn: *string, // Required
		// UserName: *string, // Required
	}

	if len(_iamPolicyArn) > 0 {
		input.PolicyArn = aws.String(_iamPolicyArn)
	}
	if len(_iamUserName) > 0 {
		input.UserName = aws.String(_iamUserName)
	}

	if resp, err := client.DetachUserPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disables the management of privileged root user credentials across member
// accounts in your organization. When you disable this feature, the management
// account and the delegated administrator for IAM can no longer manage root user
// credentials for member accounts in your organization.
func iam_DisableOrganizationsRootCredentialsManagement(cfg aws.Config, client *iam.Client) {
	input := &iam.DisableOrganizationsRootCredentialsManagementInput{}

	if resp, err := client.DisableOrganizationsRootCredentialsManagement(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disables root user sessions for privileged tasks across member accounts in your
// organization. When you disable this feature, the management account and the
// delegated administrator for IAM can no longer perform privileged tasks on member
// accounts in your organization.
func iam_DisableOrganizationsRootSessions(cfg aws.Config, client *iam.Client) {
	input := &iam.DisableOrganizationsRootSessionsInput{}

	if resp, err := client.DisableOrganizationsRootSessions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disables the outbound identity federation feature for your Amazon Web Services
// account. When disabled, IAM principals in the account cannot use the
// GetWebIdentityToken API to obtain JSON Web Tokens (JWTs) for authentication with
// external services. This operation does not affect tokens that were issued before
// the feature was disabled.
func iam_DisableOutboundWebIdentityFederation(cfg aws.Config, client *iam.Client) {
	input := &iam.DisableOutboundWebIdentityFederationInput{}

	if resp, err := client.DisableOutboundWebIdentityFederation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables the specified MFA device and associates it with the specified IAM user.
// When enabled, the MFA device is required for every subsequent login by the IAM
// user associated with the device.
func iam_EnableMFADevice(cfg aws.Config, client *iam.Client) {
	input := &iam.EnableMFADeviceInput{
		// AuthenticationCode1: *string, // Required
		// AuthenticationCode2: *string, // Required
		// SerialNumber: *string, // Required
		// UserName: *string, // Required
	}

	if len(_iamAuthenticationCode1) > 0 {
		input.AuthenticationCode1 = aws.String(_iamAuthenticationCode1)
	}
	if len(_iamAuthenticationCode2) > 0 {
		input.AuthenticationCode2 = aws.String(_iamAuthenticationCode2)
	}
	if len(_iamSerialNumber) > 0 {
		input.SerialNumber = aws.String(_iamSerialNumber)
	}
	if len(_iamUserName) > 0 {
		input.UserName = aws.String(_iamUserName)
	}

	if resp, err := client.EnableMFADevice(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables the management of privileged root user credentials across member
// accounts in your organization. When you enable root credentials management for [centralized root access]
// , the management account and the delegated administrator for IAM can manage root
// user credentials for member accounts in your organization.
//
// Before you enable centralized root access, you must have an account configured
// with the following settings:
//
// - You must manage your Amazon Web Services accounts in [Organizations].
//
// - Enable trusted access for Identity and Access Management in Organizations.
// For details, see [IAM and Organizations]in the Organizations User Guide.
//
// [Organizations]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_introduction.html
// [centralized root access]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_root-user.html#id_root-user-access-management
// [IAM and Organizations]: https://docs.aws.amazon.com/organizations/latest/userguide/services-that-can-integrate-iam.html
func iam_EnableOrganizationsRootCredentialsManagement(cfg aws.Config, client *iam.Client) {
	input := &iam.EnableOrganizationsRootCredentialsManagementInput{}

	if resp, err := client.EnableOrganizationsRootCredentialsManagement(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Allows the management account or delegated administrator to perform privileged
// tasks on member accounts in your organization. For more information, see [Centrally manage root access for member accounts]in the
// Identity and Access Management User Guide.
//
// Before you enable this feature, you must have an account configured with the
// following settings:
//
// - You must manage your Amazon Web Services accounts in [Organizations].
//
// - Enable trusted access for Identity and Access Management in Organizations.
// For details, see [IAM and Organizations]in the Organizations User Guide.
//
// [Centrally manage root access for member accounts]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_root-user.html#id_root-user-access-management
// [Organizations]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_introduction.html
// [IAM and Organizations]: https://docs.aws.amazon.com/organizations/latest/userguide/services-that-can-integrate-ra.html
func iam_EnableOrganizationsRootSessions(cfg aws.Config, client *iam.Client) {
	input := &iam.EnableOrganizationsRootSessionsInput{}

	if resp, err := client.EnableOrganizationsRootSessions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables the outbound identity federation feature for your Amazon Web Services
// account. When enabled, IAM principals in your account can use the
// GetWebIdentityToken API to obtain JSON Web Tokens (JWTs) for secure
// authentication with external services. This operation also generates a unique
// issuer URL for your Amazon Web Services account.
func iam_EnableOutboundWebIdentityFederation(cfg aws.Config, client *iam.Client) {
	input := &iam.EnableOutboundWebIdentityFederationInput{}

	if resp, err := client.EnableOutboundWebIdentityFederation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Generates a credential report for the Amazon Web Services account. For more
// information about the credential report, see [Getting credential reports]in the IAM User Guide.
//
// [Getting credential reports]: https://docs.aws.amazon.com/IAM/latest/UserGuide/credential-reports.html
func iam_GenerateCredentialReport(cfg aws.Config, client *iam.Client) {
	input := &iam.GenerateCredentialReportInput{}

	if resp, err := client.GenerateCredentialReport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Generates a report for service last accessed data for Organizations. You can
// generate a report for any entities (organization root, organizational unit, or
// account) or policies in your organization.
//
// To call this operation, you must be signed in using your Organizations
// management account credentials. You can use your long-term IAM user or root user
// credentials, or temporary credentials from assuming an IAM role. SCPs must be
// enabled for your organization root. You must have the required IAM and
// Organizations permissions. For more information, see [Refining permissions using service last accessed data]in the IAM User Guide.
//
// You can generate a service last accessed data report for entities by specifying
// only the entity's path. This data includes a list of services that are allowed
// by any service control policies (SCPs) that apply to the entity.
//
// You can generate a service last accessed data report for a policy by specifying
// an entity's path and an optional Organizations policy ID. This data includes a
// list of services that are allowed by the specified SCP.
//
// For each service in both report types, the data includes the most recent
// account activity that the policy allows to account principals in the entity or
// the entity's children. For important information about the data, reporting
// period, permissions required, troubleshooting, and supported Regions see [Reducing permissions using service last accessed data]in the
// IAM User Guide.
//
// The data includes all attempts to access Amazon Web Services, not just the
// successful ones. This includes all attempts that were made using the Amazon Web
// Services Management Console, the Amazon Web Services API through any of the
// SDKs, or any of the command line tools. An unexpected entry in the service last
// accessed data does not mean that an account has been compromised, because the
// request might have been denied. Refer to your CloudTrail logs as the
// authoritative source for information about all API calls and whether they were
// successful or denied access. For more information, see [Logging IAM events with CloudTrail]in the IAM User Guide.
//
// This operation returns a JobId . Use this parameter in the [GetOrganizationsAccessReport] operation to check
// the status of the report generation. To check the status of this request, use
// the JobId parameter in the [GetOrganizationsAccessReport] operation and test the JobStatus response
// parameter. When the job is complete, you can retrieve the report.
//
// To generate a service last accessed data report for entities, specify an entity
// path without specifying the optional Organizations policy ID. The type of entity
// that you specify determines the data returned in the report.
//
// - Root – When you specify the organizations root as the entity, the resulting
// report lists all of the services allowed by SCPs that are attached to your root.
// For each service, the report includes data for all accounts in your organization
// except the management account, because the management account is not limited by
// SCPs.
//
// - OU – When you specify an organizational unit (OU) as the entity, the
// resulting report lists all of the services allowed by SCPs that are attached to
// the OU and its parents. For each service, the report includes data for all
// accounts in the OU or its children. This data excludes the management account,
// because the management account is not limited by SCPs.
//
// - management account – When you specify the management account, the resulting
// report lists all Amazon Web Services services, because the management account is
// not limited by SCPs. For each service, the report includes data for only the
// management account.
//
// - Account – When you specify another account as the entity, the resulting
// report lists all of the services allowed by SCPs that are attached to the
// account and its parents. For each service, the report includes data for only the
// specified account.
//
// To generate a service last accessed data report for policies, specify an entity
// path and the optional Organizations policy ID. The type of entity that you
// specify determines the data returned for each service.
//
// - Root – When you specify the root entity and a policy ID, the resulting
// report lists all of the services that are allowed by the specified SCP. For each
// service, the report includes data for all accounts in your organization to which
// the SCP applies. This data excludes the management account, because the
// management account is not limited by SCPs. If the SCP is not attached to any
// entities in the organization, then the report will return a list of services
// with no data.
//
// - OU – When you specify an OU entity and a policy ID, the resulting report
// lists all of the services that are allowed by the specified SCP. For each
// service, the report includes data for all accounts in the OU or its children to
// which the SCP applies. This means that other accounts outside the OU that are
// affected by the SCP might not be included in the data. This data excludes the
// management account, because the management account is not limited by SCPs. If
// the SCP is not attached to the OU or one of its children, the report will return
// a list of services with no data.
//
// - management account – When you specify the management account, the resulting
// report lists all Amazon Web Services services, because the management account is
// not limited by SCPs. If you specify a policy ID in the CLI or API, the policy is
// ignored. For each service, the report includes data for only the management
// account.
//
// - Account – When you specify another account entity and a policy ID, the
// resulting report lists all of the services that are allowed by the specified
// SCP. For each service, the report includes data for only the specified account.
// This means that other accounts in the organization that are affected by the SCP
// might not be included in the data. If the SCP is not attached to the account,
// the report will return a list of services with no data.
//
// Service last accessed data does not use other policy types when determining
// whether a principal could access a service. These other policy types include
// identity-based policies, resource-based policies, access control lists, IAM
// permissions boundaries, and STS assume role policies. It only applies SCP logic.
// For more about the evaluation of policy types, see [Evaluating policies]in the IAM User Guide.
//
// For more information about service last accessed data, see [Reducing policy scope by viewing user activity] in the IAM User
// Guide.
//
// [Logging IAM events with CloudTrail]: https://docs.aws.amazon.com/IAM/latest/UserGuide/cloudtrail-integration.html
// [Refining permissions using service last accessed data]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_access-advisor.html
// [Reducing permissions using service last accessed data]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_access-advisor.html
// [GetOrganizationsAccessReport]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetOrganizationsAccessReport.html
// [Evaluating policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_evaluation-logic.html#policy-eval-basics
// [Reducing policy scope by viewing user activity]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_access-advisor.html
func iam_GenerateOrganizationsAccessReport(cfg aws.Config, client *iam.Client) {
	input := &iam.GenerateOrganizationsAccessReportInput{
		// EntityPath: *string, // Required
	}

	if len(_iamEntityPath) > 0 {
		input.EntityPath = aws.String(_iamEntityPath)
	}
	if len(_iamOrganizationsPolicyId) > 0 {
		input.OrganizationsPolicyId = aws.String(_iamOrganizationsPolicyId)
	}

	if resp, err := client.GenerateOrganizationsAccessReport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Generates a report that includes details about when an IAM resource (user,
// group, role, or policy) was last used in an attempt to access Amazon Web
// Services services. Recent activity usually appears within four hours. IAM
// reports activity for at least the last 400 days, or less if your Region began
// supporting this feature within the last year. For more information, see [Regions where data is tracked]. For
// more information about services and actions for which action last accessed
// information is displayed, see [IAM action last accessed information services and actions].
//
// The service last accessed data includes all attempts to access an Amazon Web
// Services API, not just the successful ones. This includes all attempts that were
// made using the Amazon Web Services Management Console, the Amazon Web Services
// API through any of the SDKs, or any of the command line tools. An unexpected
// entry in the service last accessed data does not mean that your account has been
// compromised, because the request might have been denied. Refer to your
// CloudTrail logs as the authoritative source for information about all API calls
// and whether they were successful or denied access. For more information, see [Logging IAM events with CloudTrail]in
// the IAM User Guide.
//
// The GenerateServiceLastAccessedDetails operation returns a JobId . Use this
// parameter in the following operations to retrieve the following details from
// your report:
//
// [GetServiceLastAccessedDetails]
// - – Use this operation for users, groups, roles, or policies to list every
// Amazon Web Services service that the resource could access using permissions
// policies. For each service, the response includes information about the most
// recent access attempt.
//
// # The JobId returned by GenerateServiceLastAccessedDetail must be used by the same
//
// role within a session, or by the same user when used to call
// GetServiceLastAccessedDetail .
//
// [GetServiceLastAccessedDetailsWithEntities]
// - – Use this operation for groups and policies to list information about the
// associated entities (users or roles) that attempted to access a specific Amazon
// Web Services service.
//
// To check the status of the GenerateServiceLastAccessedDetails request, use the
// JobId parameter in the same operations and test the JobStatus response
// parameter.
//
// For additional information about the permissions policies that allow an
// identity (user, group, or role) to access specific services, use the [ListPoliciesGrantingServiceAccess]operation.
//
// Service last accessed data does not use other policy types when determining
// whether a resource could access a service. These other policy types include
// resource-based policies, access control lists, Organizations policies, IAM
// permissions boundaries, and STS assume role policies. It only applies
// permissions policy logic. For more about the evaluation of policy types, see [Evaluating policies]in
// the IAM User Guide.
//
// For more information about service and action last accessed data, see [Reducing permissions using service last accessed data] in the
// IAM User Guide.
//
// [Logging IAM events with CloudTrail]: https://docs.aws.amazon.com/IAM/latest/UserGuide/cloudtrail-integration.html
// [GetServiceLastAccessedDetails]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetServiceLastAccessedDetails.html
// [ListPoliciesGrantingServiceAccess]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_ListPoliciesGrantingServiceAccess.html
// [Reducing permissions using service last accessed data]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_access-advisor.html
// [Regions where data is tracked]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_access-advisor.html#access-advisor_tracking-period
// [Evaluating policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_evaluation-logic.html#policy-eval-basics
// [GetServiceLastAccessedDetailsWithEntities]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetServiceLastAccessedDetailsWithEntities.html
// [IAM action last accessed information services and actions]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_access-advisor-action-last-accessed.html
func iam_GenerateServiceLastAccessedDetails(cfg aws.Config, client *iam.Client) {
	input := &iam.GenerateServiceLastAccessedDetailsInput{
		// Arn: *string, // Required
	}

	if len(_iamArn) > 0 {
		input.Arn = aws.String(_iamArn)
	}
	if len(_iamGranularity) > 0 {
		if err := assignInputField(input, "Granularity", _iamGranularity); err != nil {
			log.Errorf("invalid --granularity: %s", err.Error())
			return
		}
	}

	if resp, err := client.GenerateServiceLastAccessedDetails(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about when the specified access key was last used. The
// information includes the date and time of last use, along with the Amazon Web
// Services service and Region that were specified in the last request made with
// that key.
func iam_GetAccessKeyLastUsed(cfg aws.Config, client *iam.Client) {
	input := &iam.GetAccessKeyLastUsedInput{
		// AccessKeyId: *string, // Required
	}

	if len(_iamAccessKeyId) > 0 {
		input.AccessKeyId = aws.String(_iamAccessKeyId)
	}

	if resp, err := client.GetAccessKeyLastUsed(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about all IAM users, groups, roles, and policies in your
// Amazon Web Services account, including their relationships to one another. Use
// this operation to obtain a snapshot of the configuration of IAM permissions
// (users, groups, roles, and policies) in your account.
//
// Policies returned by this operation are URL-encoded compliant with [RFC 3986]. You can
// use a URL decoding method to convert the policy back to plain JSON text. For
// example, if you use Java, you can use the decode method of the
// java.net.URLDecoder utility class in the Java SDK. Other languages and SDKs
// provide similar functionality, and some SDKs do this decoding automatically.
//
// You can optionally filter the results using the Filter parameter. You can
// paginate the results using the MaxItems and Marker parameters.
//
// [RFC 3986]: https://tools.ietf.org/html/rfc3986
func iam_GetAccountAuthorizationDetails(cfg aws.Config, client *iam.Client) {
	input := &iam.GetAccountAuthorizationDetailsInput{}

	if len(_iamFilter) > 0 {
		if err := assignInputField(input, "Filter", _iamFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_iamMarker) > 0 {
		input.Marker = aws.String(_iamMarker)
	}
	if len(_iamMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _iamMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.GetAccountAuthorizationDetails(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iam.GetAccountAuthorizationDetailsOutput
	p := iam.NewGetAccountAuthorizationDetailsPaginator(client, input)
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

// Retrieves the password policy for the Amazon Web Services account. This tells
// you the complexity requirements and mandatory rotation periods for the IAM user
// passwords in your account. For more information about using a password policy,
// see [Managing an IAM password policy].
//
// [Managing an IAM password policy]: https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_ManagingPasswordPolicies.html
func iam_GetAccountPasswordPolicy(cfg aws.Config, client *iam.Client) {
	input := &iam.GetAccountPasswordPolicyInput{}

	if resp, err := client.GetAccountPasswordPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about IAM entity usage and IAM quotas in the Amazon Web
// Services account.
//
// For information about IAM quotas, see [IAM and STS quotas] in the IAM User Guide.
//
// [IAM and STS quotas]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-quotas.html
func iam_GetAccountSummary(cfg aws.Config, client *iam.Client) {
	input := &iam.GetAccountSummaryInput{}

	if resp, err := client.GetAccountSummary(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a list of all of the context keys referenced in the input policies. The
// policies are supplied as a list of one or more strings. To get the context keys
// from policies associated with an IAM user, group, or role, use [GetContextKeysForPrincipalPolicy].
//
// Context keys are variables maintained by Amazon Web Services and its services
// that provide details about the context of an API query request. Context keys can
// be evaluated by testing against a value specified in an IAM policy. Use
// GetContextKeysForCustomPolicy to understand what key names and values you must
// supply when you call [SimulateCustomPolicy]. Note that all parameters are shown in unencoded form
// here for clarity but must be URL encoded to be included as a part of a real HTML
// request.
//
// [GetContextKeysForPrincipalPolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetContextKeysForPrincipalPolicy.html
// [SimulateCustomPolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_SimulateCustomPolicy.html
func iam_GetContextKeysForCustomPolicy(cfg aws.Config, client *iam.Client) {
	input := &iam.GetContextKeysForCustomPolicyInput{
		// PolicyInputList: []string, // Required
	}

	if len(_iamPolicyInputList) > 0 {
		input.PolicyInputList = append([]string(nil), _iamPolicyInputList...)
	}

	if resp, err := client.GetContextKeysForCustomPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a list of all of the context keys referenced in all the IAM policies that
// are attached to the specified IAM entity. The entity can be an IAM user, group,
// or role. If you specify a user, then the request also includes all of the
// policies attached to groups that the user is a member of.
//
// You can optionally include a list of one or more additional policies, specified
// as strings. If you want to include only a list of policies by string, use [GetContextKeysForCustomPolicy]
// instead.
//
// Note: This operation discloses information about the permissions granted to
// other users. If you do not want users to see other user's permissions, then
// consider allowing them to use [GetContextKeysForCustomPolicy]instead.
//
// Context keys are variables maintained by Amazon Web Services and its services
// that provide details about the context of an API query request. Context keys can
// be evaluated by testing against a value in an IAM policy. Use [GetContextKeysForPrincipalPolicy]to understand
// what key names and values you must supply when you call [SimulatePrincipalPolicy].
//
// [GetContextKeysForPrincipalPolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetContextKeysForPrincipalPolicy.html
// [GetContextKeysForCustomPolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetContextKeysForCustomPolicy.html
// [SimulatePrincipalPolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_SimulatePrincipalPolicy.html
func iam_GetContextKeysForPrincipalPolicy(cfg aws.Config, client *iam.Client) {
	input := &iam.GetContextKeysForPrincipalPolicyInput{
		// PolicySourceArn: *string, // Required
	}

	if len(_iamPolicySourceArn) > 0 {
		input.PolicySourceArn = aws.String(_iamPolicySourceArn)
	}
	if len(_iamPolicyInputList) > 0 {
		input.PolicyInputList = append([]string(nil), _iamPolicyInputList...)
	}

	if resp, err := client.GetContextKeysForPrincipalPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a credential report for the Amazon Web Services account. For more
// information about the credential report, see [Getting credential reports]in the IAM User Guide.
//
// [Getting credential reports]: https://docs.aws.amazon.com/IAM/latest/UserGuide/credential-reports.html
func iam_GetCredentialReport(cfg aws.Config, client *iam.Client) {
	input := &iam.GetCredentialReportInput{}

	if resp, err := client.GetCredentialReport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a specific delegation request.
// If a delegation request has no owner or owner account, GetDelegationRequest for
// that delegation request can be called by any account. If the owner account is
// assigned but there is no owner id, only identities within that owner account can
// call GetDelegationRequest for the delegation request. Once the delegation
// request is fully owned, the owner of the request gets a default permission to
// get that delegation request. For more details, see [Managing Permissions for Delegation Requests].
//
// [Managing Permissions for Delegation Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies-temporary-delegation.html#temporary-delegation-managing-permissions
func iam_GetDelegationRequest(cfg aws.Config, client *iam.Client) {
	input := &iam.GetDelegationRequestInput{
		// DelegationRequestId: *string, // Required
	}

	if len(_iamDelegationRequestId) > 0 {
		input.DelegationRequestId = aws.String(_iamDelegationRequestId)
	}
	if len(_iamDelegationPermissionCheck) > 0 {
		if err := assignInputField(input, "DelegationPermissionCheck", _iamDelegationPermissionCheck); err != nil {
			log.Errorf("invalid --delegation-permission-check: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetDelegationRequest(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of IAM users that are in the specified IAM group. You can
// paginate the results using the MaxItems and Marker parameters.
func iam_GetGroup(cfg aws.Config, client *iam.Client) {
	input := &iam.GetGroupInput{
		// GroupName: *string, // Required
	}

	if len(_iamGroupName) > 0 {
		input.GroupName = aws.String(_iamGroupName)
	}
	if len(_iamMarker) > 0 {
		input.Marker = aws.String(_iamMarker)
	}
	if len(_iamMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _iamMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.GetGroup(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iam.GetGroupOutput
	p := iam.NewGetGroupPaginator(client, input)
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

// Retrieves the specified inline policy document that is embedded in the
// specified IAM group.
//
// Policies returned by this operation are URL-encoded compliant with [RFC 3986]. You can
// use a URL decoding method to convert the policy back to plain JSON text. For
// example, if you use Java, you can use the decode method of the
// java.net.URLDecoder utility class in the Java SDK. Other languages and SDKs
// provide similar functionality, and some SDKs do this decoding automatically.
//
// An IAM group can also have managed policies attached to it. To retrieve a
// managed policy document that is attached to a group, use [GetPolicy]to determine the
// policy's default version, then use [GetPolicyVersion]to retrieve the policy document.
//
// For more information about policies, see [Managed policies and inline policies] in the IAM User Guide.
//
// [RFC 3986]: https://tools.ietf.org/html/rfc3986
// [GetPolicyVersion]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetPolicyVersion.html
// [GetPolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetPolicy.html
// [Managed policies and inline policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html
func iam_GetGroupPolicy(cfg aws.Config, client *iam.Client) {
	input := &iam.GetGroupPolicyInput{
		// GroupName: *string, // Required
		// PolicyName: *string, // Required
	}

	if len(_iamGroupName) > 0 {
		input.GroupName = aws.String(_iamGroupName)
	}
	if len(_iamPolicyName) > 0 {
		input.PolicyName = aws.String(_iamPolicyName)
	}

	if resp, err := client.GetGroupPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a human readable summary for a given entity. At this time, the only
// supported entity type is delegation-request
//
// This method uses a Large Language Model (LLM) to generate the summary.
//
// If a delegation request has no owner or owner account, GetHumanReadableSummary
// for that delegation request can be called by any account. If the owner account
// is assigned but there is no owner id, only identities within that owner account
// can call GetHumanReadableSummary for the delegation request to retrieve a
// summary of that request. Once the delegation request is fully owned, the owner
// of the request gets a default permission to get that delegation request. For
// more details, read default permissions granted to delegation requests. These rules are identical to [GetDelegationRequest] API behavior, such that a
// party who has permissions to call [GetDelegationRequest]for a given delegation request will always be
// able to retrieve the human readable summary for that request.
//
// [GetDelegationRequest]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetDelegationRequest.html
func iam_GetHumanReadableSummary(cfg aws.Config, client *iam.Client) {
	input := &iam.GetHumanReadableSummaryInput{
		// EntityArn: *string, // Required
	}

	if len(_iamEntityArn) > 0 {
		input.EntityArn = aws.String(_iamEntityArn)
	}
	if len(_iamLocale) > 0 {
		input.Locale = aws.String(_iamLocale)
	}

	if resp, err := client.GetHumanReadableSummary(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the specified instance profile, including the
// instance profile's path, GUID, ARN, and role. For more information about
// instance profiles, see [Using instance profiles]in the IAM User Guide.
//
// [Using instance profiles]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use_switch-role-ec2_instance-profiles.html
func iam_GetInstanceProfile(cfg aws.Config, client *iam.Client) {
	input := &iam.GetInstanceProfileInput{
		// InstanceProfileName: *string, // Required
	}

	if len(_iamInstanceProfileName) > 0 {
		input.InstanceProfileName = aws.String(_iamInstanceProfileName)
	}

	if resp, err := client.GetInstanceProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the user name for the specified IAM user. A login profile is created
// when you create a password for the user to access the Amazon Web Services
// Management Console. If the user does not exist or does not have a password, the
// operation returns a 404 ( NoSuchEntity ) error.
//
// If you create an IAM user with access to the console, the CreateDate reflects
// the date you created the initial password for the user.
//
// If you create an IAM user with programmatic access, and then later add a
// password for the user to access the Amazon Web Services Management Console, the
// CreateDate reflects the initial password creation date. A user with programmatic
// access does not have a login profile unless you create a password for the user
// to access the Amazon Web Services Management Console.
func iam_GetLoginProfile(cfg aws.Config, client *iam.Client) {
	input := &iam.GetLoginProfileInput{}

	if len(_iamUserName) > 0 {
		input.UserName = aws.String(_iamUserName)
	}

	if resp, err := client.GetLoginProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about an MFA device for a specified user.
func iam_GetMFADevice(cfg aws.Config, client *iam.Client) {
	input := &iam.GetMFADeviceInput{
		// SerialNumber: *string, // Required
	}

	if len(_iamSerialNumber) > 0 {
		input.SerialNumber = aws.String(_iamSerialNumber)
	}
	if len(_iamUserName) > 0 {
		input.UserName = aws.String(_iamUserName)
	}

	if resp, err := client.GetMFADevice(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the specified OpenID Connect (OIDC) provider resource
// object in IAM.
func iam_GetOpenIDConnectProvider(cfg aws.Config, client *iam.Client) {
	input := &iam.GetOpenIDConnectProviderInput{
		// OpenIDConnectProviderArn: *string, // Required
	}

	if len(_iamOpenIDConnectProviderArn) > 0 {
		input.OpenIDConnectProviderArn = aws.String(_iamOpenIDConnectProviderArn)
	}

	if resp, err := client.GetOpenIDConnectProvider(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the service last accessed data report for Organizations that was
// previously generated using the [GenerateOrganizationsAccessReport]operation. This operation retrieves the status
// of your report job and the report contents.
//
// Depending on the parameters that you passed when you generated the report, the
// data returned could include different information. For details, see [GenerateOrganizationsAccessReport].
//
// To call this operation, you must be signed in to the management account in your
// organization. SCPs must be enabled for your organization root. You must have
// permissions to perform this operation. For more information, see [Refining permissions using service last accessed data]in the IAM
// User Guide.
//
// For each service that principals in an account (root user, IAM users, or IAM
// roles) could access using SCPs, the operation returns details about the most
// recent access attempt. If there was no attempt, the service is listed without
// details about the most recent attempt to access the service. If the operation
// fails, it returns the reason that it failed.
//
// By default, the list is sorted by service namespace.
//
// [GenerateOrganizationsAccessReport]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GenerateOrganizationsAccessReport.html
// [Refining permissions using service last accessed data]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_access-advisor.html
func iam_GetOrganizationsAccessReport(cfg aws.Config, client *iam.Client) {
	input := &iam.GetOrganizationsAccessReportInput{
		// JobId: *string, // Required
	}

	if len(_iamJobId) > 0 {
		input.JobId = aws.String(_iamJobId)
	}
	if len(_iamMarker) > 0 {
		input.Marker = aws.String(_iamMarker)
	}
	if len(_iamMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _iamMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}
	if len(_iamSortKey) > 0 {
		if err := assignInputField(input, "SortKey", _iamSortKey); err != nil {
			log.Errorf("invalid --sort-key: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetOrganizationsAccessReport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the configuration information for the outbound identity federation
// feature in your Amazon Web Services account. The response includes the unique
// issuer URL for your Amazon Web Services account and the current enabled/disabled
// status of the feature. Use this operation to obtain the issuer URL that you need
// to configure trust relationships with external services.
func iam_GetOutboundWebIdentityFederationInfo(cfg aws.Config, client *iam.Client) {
	input := &iam.GetOutboundWebIdentityFederationInfoInput{}

	if resp, err := client.GetOutboundWebIdentityFederationInfo(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the specified managed policy, including the
// policy's default version and the total number of IAM users, groups, and roles to
// which the policy is attached. To retrieve the list of the specific users,
// groups, and roles that the policy is attached to, use [ListEntitiesForPolicy]. This operation returns
// metadata about the policy. To retrieve the actual policy document for a specific
// version of the policy, use [GetPolicyVersion].
//
// This operation retrieves information about managed policies. To retrieve
// information about an inline policy that is embedded with an IAM user, group, or
// role, use [GetUserPolicy], [GetGroupPolicy], or [GetRolePolicy].
//
// For more information about policies, see [Managed policies and inline policies] in the IAM User Guide.
//
// [ListEntitiesForPolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_ListEntitiesForPolicy.html
// [GetRolePolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetRolePolicy.html
// [GetPolicyVersion]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetPolicyVersion.html
// [GetGroupPolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetGroupPolicy.html
// [GetUserPolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetUserPolicy.html
// [Managed policies and inline policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html
func iam_GetPolicy(cfg aws.Config, client *iam.Client) {
	input := &iam.GetPolicyInput{
		// PolicyArn: *string, // Required
	}

	if len(_iamPolicyArn) > 0 {
		input.PolicyArn = aws.String(_iamPolicyArn)
	}

	if resp, err := client.GetPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the specified version of the specified managed
// policy, including the policy document.
//
// Policies returned by this operation are URL-encoded compliant with [RFC 3986]. You can
// use a URL decoding method to convert the policy back to plain JSON text. For
// example, if you use Java, you can use the decode method of the
// java.net.URLDecoder utility class in the Java SDK. Other languages and SDKs
// provide similar functionality, and some SDKs do this decoding automatically.
//
// To list the available versions for a policy, use [ListPolicyVersions].
//
// This operation retrieves information about managed policies. To retrieve
// information about an inline policy that is embedded in a user, group, or role,
// use [GetUserPolicy], [GetGroupPolicy], or [GetRolePolicy].
//
// For more information about the types of policies, see [Managed policies and inline policies] in the IAM User Guide.
//
// For more information about managed policy versions, see [Versioning for managed policies] in the IAM User Guide.
//
// [RFC 3986]: https://tools.ietf.org/html/rfc3986
// [GetRolePolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetRolePolicy.html
// [GetGroupPolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetGroupPolicy.html
// [GetUserPolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetUserPolicy.html
// [Versioning for managed policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-versions.html
// [ListPolicyVersions]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_ListPolicyVersions.html
// [Managed policies and inline policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html
func iam_GetPolicyVersion(cfg aws.Config, client *iam.Client) {
	input := &iam.GetPolicyVersionInput{
		// PolicyArn: *string, // Required
		// VersionId: *string, // Required
	}

	if len(_iamPolicyArn) > 0 {
		input.PolicyArn = aws.String(_iamPolicyArn)
	}
	if len(_iamVersionId) > 0 {
		input.VersionId = aws.String(_iamVersionId)
	}

	if resp, err := client.GetPolicyVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the specified role, including the role's path,
// GUID, ARN, and the role's trust policy that grants permission to assume the
// role. For more information about roles, see [IAM roles]in the IAM User Guide.
//
// Policies returned by this operation are URL-encoded compliant with [RFC 3986]. You can
// use a URL decoding method to convert the policy back to plain JSON text. For
// example, if you use Java, you can use the decode method of the
// java.net.URLDecoder utility class in the Java SDK. Other languages and SDKs
// provide similar functionality, and some SDKs do this decoding automatically.
//
// [RFC 3986]: https://tools.ietf.org/html/rfc3986
// [IAM roles]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles.html
func iam_GetRole(cfg aws.Config, client *iam.Client) {
	input := &iam.GetRoleInput{
		// RoleName: *string, // Required
	}

	if len(_iamRoleName) > 0 {
		input.RoleName = aws.String(_iamRoleName)
	}

	if resp, err := client.GetRole(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the specified inline policy document that is embedded with the
// specified IAM role.
//
// Policies returned by this operation are URL-encoded compliant with [RFC 3986]. You can
// use a URL decoding method to convert the policy back to plain JSON text. For
// example, if you use Java, you can use the decode method of the
// java.net.URLDecoder utility class in the Java SDK. Other languages and SDKs
// provide similar functionality, and some SDKs do this decoding automatically.
//
// An IAM role can also have managed policies attached to it. To retrieve a
// managed policy document that is attached to a role, use [GetPolicy]to determine the
// policy's default version, then use [GetPolicyVersion]to retrieve the policy document.
//
// For more information about policies, see [Managed policies and inline policies] in the IAM User Guide.
//
// For more information about roles, see [IAM roles] in the IAM User Guide.
//
// [RFC 3986]: https://tools.ietf.org/html/rfc3986
// [IAM roles]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles.html
// [GetPolicyVersion]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetPolicyVersion.html
// [GetPolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetPolicy.html
// [Managed policies and inline policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html
func iam_GetRolePolicy(cfg aws.Config, client *iam.Client) {
	input := &iam.GetRolePolicyInput{
		// PolicyName: *string, // Required
		// RoleName: *string, // Required
	}

	if len(_iamPolicyName) > 0 {
		input.PolicyName = aws.String(_iamPolicyName)
	}
	if len(_iamRoleName) > 0 {
		input.RoleName = aws.String(_iamRoleName)
	}

	if resp, err := client.GetRolePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the SAML provider metadocument that was uploaded when the IAM SAML
// provider resource object was created or updated.
//
// This operation requires [Signature Version 4].
//
// [Signature Version 4]: https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html
func iam_GetSAMLProvider(cfg aws.Config, client *iam.Client) {
	input := &iam.GetSAMLProviderInput{
		// SAMLProviderArn: *string, // Required
	}

	if len(_iamSAMLProviderArn) > 0 {
		input.SAMLProviderArn = aws.String(_iamSAMLProviderArn)
	}

	if resp, err := client.GetSAMLProvider(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the specified server certificate stored in IAM.
// For more information about working with server certificates, see [Working with server certificates] in the IAM
// User Guide. This topic includes a list of Amazon Web Services services that can
// use the server certificates that you manage with IAM.
//
// [Working with server certificates]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_server-certs.html
func iam_GetServerCertificate(cfg aws.Config, client *iam.Client) {
	input := &iam.GetServerCertificateInput{
		// ServerCertificateName: *string, // Required
	}

	if len(_iamServerCertificateName) > 0 {
		input.ServerCertificateName = aws.String(_iamServerCertificateName)
	}

	if resp, err := client.GetServerCertificate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a service last accessed report that was created using the
// GenerateServiceLastAccessedDetails operation. You can use the JobId parameter
// in GetServiceLastAccessedDetails to retrieve the status of your report job.
// When the report is complete, you can retrieve the generated report. The report
// includes a list of Amazon Web Services services that the resource (user, group,
// role, or managed policy) can access.
//
// Service last accessed data does not use other policy types when determining
// whether a resource could access a service. These other policy types include
// resource-based policies, access control lists, Organizations policies, IAM
// permissions boundaries, and STS assume role policies. It only applies
// permissions policy logic. For more about the evaluation of policy types, see [Evaluating policies]in
// the IAM User Guide.
//
// For each service that the resource could access using permissions policies, the
// operation returns details about the most recent access attempt. If there was no
// attempt, the service is listed without details about the most recent attempt to
// access the service. If the operation fails, the GetServiceLastAccessedDetails
// operation returns the reason that it failed.
//
// The GetServiceLastAccessedDetails operation returns a list of services. This
// list includes the number of entities that have attempted to access the service
// and the date and time of the last attempt. It also returns the ARN of the
// following entity, depending on the resource ARN that you used to generate the
// report:
//
// - User – Returns the user ARN that you used to generate the report
//
// - Group – Returns the ARN of the group member (user) that last attempted to
// access the service
//
// - Role – Returns the role ARN that you used to generate the report
//
// - Policy – Returns the ARN of the user or role that last used the policy to
// attempt to access the service
//
// By default, the list is sorted by service namespace.
//
// If you specified ACTION_LEVEL granularity when you generated the report, this
// operation returns service and action last accessed data. This includes the most
// recent access attempt for each tracked action within a service. Otherwise, this
// operation returns only service data.
//
// For more information about service and action last accessed data, see [Reducing permissions using service last accessed data] in the
// IAM User Guide.
//
// [Reducing permissions using service last accessed data]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_access-advisor.html
// [Evaluating policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_evaluation-logic.html#policy-eval-basics
func iam_GetServiceLastAccessedDetails(cfg aws.Config, client *iam.Client) {
	input := &iam.GetServiceLastAccessedDetailsInput{
		// JobId: *string, // Required
	}

	if len(_iamJobId) > 0 {
		input.JobId = aws.String(_iamJobId)
	}
	if len(_iamMarker) > 0 {
		input.Marker = aws.String(_iamMarker)
	}
	if len(_iamMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _iamMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetServiceLastAccessedDetails(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// After you generate a group or policy report using the
// GenerateServiceLastAccessedDetails operation, you can use the JobId parameter
// in GetServiceLastAccessedDetailsWithEntities . This operation retrieves the
// status of your report job and a list of entities that could have used group or
// policy permissions to access the specified service.
//
// - Group – For a group report, this operation returns a list of users in the
// group that could have used the group’s policies in an attempt to access the
// service.
//
// - Policy – For a policy report, this operation returns a list of entities
// (users or roles) that could have used the policy in an attempt to access the
// service.
//
// You can also use this operation for user or role reports to retrieve details
// about those entities.
//
// If the operation fails, the GetServiceLastAccessedDetailsWithEntities operation
// returns the reason that it failed.
//
// By default, the list of associated entities is sorted by date, with the most
// recent access listed first.
func iam_GetServiceLastAccessedDetailsWithEntities(cfg aws.Config, client *iam.Client) {
	input := &iam.GetServiceLastAccessedDetailsWithEntitiesInput{
		// JobId: *string, // Required
		// ServiceNamespace: *string, // Required
	}

	if len(_iamJobId) > 0 {
		input.JobId = aws.String(_iamJobId)
	}
	if len(_iamServiceNamespace) > 0 {
		input.ServiceNamespace = aws.String(_iamServiceNamespace)
	}
	if len(_iamMarker) > 0 {
		input.Marker = aws.String(_iamMarker)
	}
	if len(_iamMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _iamMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetServiceLastAccessedDetailsWithEntities(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the status of your service-linked role deletion. After you use [DeleteServiceLinkedRole] to
// submit a service-linked role for deletion, you can use the DeletionTaskId
// parameter in GetServiceLinkedRoleDeletionStatus to check the status of the
// deletion. If the deletion fails, this operation returns the reason that it
// failed, if that information is returned by the service.
//
// [DeleteServiceLinkedRole]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_DeleteServiceLinkedRole.html
func iam_GetServiceLinkedRoleDeletionStatus(cfg aws.Config, client *iam.Client) {
	input := &iam.GetServiceLinkedRoleDeletionStatusInput{
		// DeletionTaskId: *string, // Required
	}

	if len(_iamDeletionTaskId) > 0 {
		input.DeletionTaskId = aws.String(_iamDeletionTaskId)
	}

	if resp, err := client.GetServiceLinkedRoleDeletionStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the specified SSH public key, including metadata about the key.
// The SSH public key retrieved by this operation is used only for authenticating
// the associated IAM user to an CodeCommit repository. For more information about
// using SSH keys to authenticate to an CodeCommit repository, see [Set up CodeCommit for SSH connections]in the
// CodeCommit User Guide.
//
// [Set up CodeCommit for SSH connections]: https://docs.aws.amazon.com/codecommit/latest/userguide/setting-up-credentials-ssh.html
func iam_GetSSHPublicKey(cfg aws.Config, client *iam.Client) {
	input := &iam.GetSSHPublicKeyInput{
		// Encoding: types.EncodingType, // Required
		// SSHPublicKeyId: *string, // Required
		// UserName: *string, // Required
	}

	if len(_iamEncoding) > 0 {
		if err := assignInputField(input, "Encoding", _iamEncoding); err != nil {
			log.Errorf("invalid --encoding: %s", err.Error())
			return
		}
	}
	if len(_iamSSHPublicKeyId) > 0 {
		input.SSHPublicKeyId = aws.String(_iamSSHPublicKeyId)
	}
	if len(_iamUserName) > 0 {
		input.UserName = aws.String(_iamUserName)
	}

	if resp, err := client.GetSSHPublicKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the specified IAM user, including the user's
// creation date, path, unique ID, and ARN.
//
// If you do not specify a user name, IAM determines the user name implicitly
// based on the Amazon Web Services access key ID used to sign the request to this
// operation.
func iam_GetUser(cfg aws.Config, client *iam.Client) {
	input := &iam.GetUserInput{}

	if len(_iamUserName) > 0 {
		input.UserName = aws.String(_iamUserName)
	}

	if resp, err := client.GetUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the specified inline policy document that is embedded in the
// specified IAM user.
//
// Policies returned by this operation are URL-encoded compliant with [RFC 3986]. You can
// use a URL decoding method to convert the policy back to plain JSON text. For
// example, if you use Java, you can use the decode method of the
// java.net.URLDecoder utility class in the Java SDK. Other languages and SDKs
// provide similar functionality, and some SDKs do this decoding automatically.
//
// An IAM user can also have managed policies attached to it. To retrieve a
// managed policy document that is attached to a user, use [GetPolicy]to determine the
// policy's default version. Then use [GetPolicyVersion]to retrieve the policy document.
//
// For more information about policies, see [Managed policies and inline policies] in the IAM User Guide.
//
// [RFC 3986]: https://tools.ietf.org/html/rfc3986
// [GetPolicyVersion]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetPolicyVersion.html
// [GetPolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetPolicy.html
// [Managed policies and inline policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html
func iam_GetUserPolicy(cfg aws.Config, client *iam.Client) {
	input := &iam.GetUserPolicyInput{
		// PolicyName: *string, // Required
		// UserName: *string, // Required
	}

	if len(_iamPolicyName) > 0 {
		input.PolicyName = aws.String(_iamPolicyName)
	}
	if len(_iamUserName) > 0 {
		input.UserName = aws.String(_iamUserName)
	}

	if resp, err := client.GetUserPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the access key IDs associated with the specified IAM
// user. If there is none, the operation returns an empty list.
//
// Although each user is limited to a small number of keys, you can still paginate
// the results using the MaxItems and Marker parameters.
//
// If the UserName is not specified, the user name is determined implicitly based
// on the Amazon Web Services access key ID used to sign the request. If a
// temporary access key is used, then UserName is required. If a long-term key is
// assigned to the user, then UserName is not required.
//
// This operation works for access keys under the Amazon Web Services account. If
// the Amazon Web Services account has no associated users, the root user returns
// it's own access key IDs by running this command.
//
// To ensure the security of your Amazon Web Services account, the secret access
// key is accessible only during key and user creation.
func iam_ListAccessKeys(cfg aws.Config, client *iam.Client) {
	input := &iam.ListAccessKeysInput{}

	if len(_iamMarker) > 0 {
		input.Marker = aws.String(_iamMarker)
	}
	if len(_iamMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _iamMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}
	if len(_iamUserName) > 0 {
		input.UserName = aws.String(_iamUserName)
	}

	if disablePaginator() {
		if resp, err := client.ListAccessKeys(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iam.ListAccessKeysOutput
	p := iam.NewListAccessKeysPaginator(client, input)
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

// Lists the account alias associated with the Amazon Web Services account (Note:
// you can have only one). For information about using an Amazon Web Services
// account alias, see [Creating, deleting, and listing an Amazon Web Services account alias]in the IAM User Guide.
//
// [Creating, deleting, and listing an Amazon Web Services account alias]: https://docs.aws.amazon.com/IAM/latest/UserGuide/console_account-alias.html#CreateAccountAlias
func iam_ListAccountAliases(cfg aws.Config, client *iam.Client) {
	input := &iam.ListAccountAliasesInput{}

	if len(_iamMarker) > 0 {
		input.Marker = aws.String(_iamMarker)
	}
	if len(_iamMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _iamMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListAccountAliases(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iam.ListAccountAliasesOutput
	p := iam.NewListAccountAliasesPaginator(client, input)
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

// Lists all managed policies that are attached to the specified IAM group.
// An IAM group can also have inline policies embedded with it. To list the inline
// policies for a group, use [ListGroupPolicies]. For information about policies, see [Managed policies and inline policies] in the IAM
// User Guide.
//
// You can paginate the results using the MaxItems and Marker parameters. You can
// use the PathPrefix parameter to limit the list of policies to only those
// matching the specified path prefix. If there are no policies attached to the
// specified group (or none that match the specified path prefix), the operation
// returns an empty list.
//
// [ListGroupPolicies]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_ListGroupPolicies.html
// [Managed policies and inline policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html
func iam_ListAttachedGroupPolicies(cfg aws.Config, client *iam.Client) {
	input := &iam.ListAttachedGroupPoliciesInput{
		// GroupName: *string, // Required
	}

	if len(_iamGroupName) > 0 {
		input.GroupName = aws.String(_iamGroupName)
	}
	if len(_iamMarker) > 0 {
		input.Marker = aws.String(_iamMarker)
	}
	if len(_iamMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _iamMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}
	if len(_iamPathPrefix) > 0 {
		input.PathPrefix = aws.String(_iamPathPrefix)
	}

	if disablePaginator() {
		if resp, err := client.ListAttachedGroupPolicies(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iam.ListAttachedGroupPoliciesOutput
	p := iam.NewListAttachedGroupPoliciesPaginator(client, input)
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

// Lists all managed policies that are attached to the specified IAM role.
// An IAM role can also have inline policies embedded with it. To list the inline
// policies for a role, use [ListRolePolicies]. For information about policies, see [Managed policies and inline policies] in the IAM User
// Guide.
//
// You can paginate the results using the MaxItems and Marker parameters. You can
// use the PathPrefix parameter to limit the list of policies to only those
// matching the specified path prefix. If there are no policies attached to the
// specified role (or none that match the specified path prefix), the operation
// returns an empty list.
//
// [ListRolePolicies]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_ListRolePolicies.html
// [Managed policies and inline policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html
func iam_ListAttachedRolePolicies(cfg aws.Config, client *iam.Client) {
	input := &iam.ListAttachedRolePoliciesInput{
		// RoleName: *string, // Required
	}

	if len(_iamRoleName) > 0 {
		input.RoleName = aws.String(_iamRoleName)
	}
	if len(_iamMarker) > 0 {
		input.Marker = aws.String(_iamMarker)
	}
	if len(_iamMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _iamMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}
	if len(_iamPathPrefix) > 0 {
		input.PathPrefix = aws.String(_iamPathPrefix)
	}

	if disablePaginator() {
		if resp, err := client.ListAttachedRolePolicies(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iam.ListAttachedRolePoliciesOutput
	p := iam.NewListAttachedRolePoliciesPaginator(client, input)
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

// Lists all managed policies that are attached to the specified IAM user.
// An IAM user can also have inline policies embedded with it. To list the inline
// policies for a user, use [ListUserPolicies]. For information about policies, see [Managed policies and inline policies] in the IAM User
// Guide.
//
// You can paginate the results using the MaxItems and Marker parameters. You can
// use the PathPrefix parameter to limit the list of policies to only those
// matching the specified path prefix. If there are no policies attached to the
// specified group (or none that match the specified path prefix), the operation
// returns an empty list.
//
// [ListUserPolicies]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_ListUserPolicies.html
// [Managed policies and inline policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html
func iam_ListAttachedUserPolicies(cfg aws.Config, client *iam.Client) {
	input := &iam.ListAttachedUserPoliciesInput{
		// UserName: *string, // Required
	}

	if len(_iamUserName) > 0 {
		input.UserName = aws.String(_iamUserName)
	}
	if len(_iamMarker) > 0 {
		input.Marker = aws.String(_iamMarker)
	}
	if len(_iamMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _iamMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}
	if len(_iamPathPrefix) > 0 {
		input.PathPrefix = aws.String(_iamPathPrefix)
	}

	if disablePaginator() {
		if resp, err := client.ListAttachedUserPolicies(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iam.ListAttachedUserPoliciesOutput
	p := iam.NewListAttachedUserPoliciesPaginator(client, input)
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

// Lists delegation requests based on the specified criteria.
// If a delegation request has no owner, even if it is assigned to a specific
// account, it will not be part of the ListDelegationRequests output for that
// account.
//
// For more details, see [Managing Permissions for Delegation Requests].
//
// [Managing Permissions for Delegation Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies-temporary-delegation.html#temporary-delegation-managing-permissions
func iam_ListDelegationRequests(cfg aws.Config, client *iam.Client) {
	input := &iam.ListDelegationRequestsInput{}

	if len(_iamMarker) > 0 {
		input.Marker = aws.String(_iamMarker)
	}
	if len(_iamMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _iamMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}
	if len(_iamOwnerId) > 0 {
		input.OwnerId = aws.String(_iamOwnerId)
	}

	if resp, err := client.ListDelegationRequests(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all IAM users, groups, and roles that the specified managed policy is
// attached to.
//
// You can use the optional EntityFilter parameter to limit the results to a
// particular type of entity (users, groups, or roles). For example, to list only
// the roles that are attached to the specified policy, set EntityFilter to Role .
//
// You can paginate the results using the MaxItems and Marker parameters.
func iam_ListEntitiesForPolicy(cfg aws.Config, client *iam.Client) {
	input := &iam.ListEntitiesForPolicyInput{
		// PolicyArn: *string, // Required
	}

	if len(_iamPolicyArn) > 0 {
		input.PolicyArn = aws.String(_iamPolicyArn)
	}
	if len(_iamEntityFilter) > 0 {
		if err := assignInputField(input, "EntityFilter", _iamEntityFilter); err != nil {
			log.Errorf("invalid --entity-filter: %s", err.Error())
			return
		}
	}
	if len(_iamMarker) > 0 {
		input.Marker = aws.String(_iamMarker)
	}
	if len(_iamMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _iamMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}
	if len(_iamPathPrefix) > 0 {
		input.PathPrefix = aws.String(_iamPathPrefix)
	}
	if len(_iamPolicyUsageFilter) > 0 {
		if err := assignInputField(input, "PolicyUsageFilter", _iamPolicyUsageFilter); err != nil {
			log.Errorf("invalid --policy-usage-filter: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListEntitiesForPolicy(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iam.ListEntitiesForPolicyOutput
	p := iam.NewListEntitiesForPolicyPaginator(client, input)
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

// Lists the names of the inline policies that are embedded in the specified IAM
// group.
//
// An IAM group can also have managed policies attached to it. To list the managed
// policies that are attached to a group, use [ListAttachedGroupPolicies]. For more information about
// policies, see [Managed policies and inline policies]in the IAM User Guide.
//
// You can paginate the results using the MaxItems and Marker parameters. If there
// are no inline policies embedded with the specified group, the operation returns
// an empty list.
//
// [ListAttachedGroupPolicies]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_ListAttachedGroupPolicies.html
// [Managed policies and inline policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html
func iam_ListGroupPolicies(cfg aws.Config, client *iam.Client) {
	input := &iam.ListGroupPoliciesInput{
		// GroupName: *string, // Required
	}

	if len(_iamGroupName) > 0 {
		input.GroupName = aws.String(_iamGroupName)
	}
	if len(_iamMarker) > 0 {
		input.Marker = aws.String(_iamMarker)
	}
	if len(_iamMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _iamMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListGroupPolicies(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iam.ListGroupPoliciesOutput
	p := iam.NewListGroupPoliciesPaginator(client, input)
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

// Lists the IAM groups that have the specified path prefix.
// You can paginate the results using the MaxItems and Marker parameters.
func iam_ListGroups(cfg aws.Config, client *iam.Client) {
	input := &iam.ListGroupsInput{}

	if len(_iamMarker) > 0 {
		input.Marker = aws.String(_iamMarker)
	}
	if len(_iamMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _iamMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}
	if len(_iamPathPrefix) > 0 {
		input.PathPrefix = aws.String(_iamPathPrefix)
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

	var results []*iam.ListGroupsOutput
	p := iam.NewListGroupsPaginator(client, input)
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

// Lists the IAM groups that the specified IAM user belongs to.
// You can paginate the results using the MaxItems and Marker parameters.
func iam_ListGroupsForUser(cfg aws.Config, client *iam.Client) {
	input := &iam.ListGroupsForUserInput{
		// UserName: *string, // Required
	}

	if len(_iamUserName) > 0 {
		input.UserName = aws.String(_iamUserName)
	}
	if len(_iamMarker) > 0 {
		input.Marker = aws.String(_iamMarker)
	}
	if len(_iamMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _iamMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListGroupsForUser(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iam.ListGroupsForUserOutput
	p := iam.NewListGroupsForUserPaginator(client, input)
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

// Lists the tags that are attached to the specified IAM instance profile. The
// returned list of tags is sorted by tag key. For more information about tagging,
// see [Tagging IAM resources]in the IAM User Guide.
//
// [Tagging IAM resources]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html
func iam_ListInstanceProfileTags(cfg aws.Config, client *iam.Client) {
	input := &iam.ListInstanceProfileTagsInput{
		// InstanceProfileName: *string, // Required
	}

	if len(_iamInstanceProfileName) > 0 {
		input.InstanceProfileName = aws.String(_iamInstanceProfileName)
	}
	if len(_iamMarker) > 0 {
		input.Marker = aws.String(_iamMarker)
	}
	if len(_iamMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _iamMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListInstanceProfileTags(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iam.ListInstanceProfileTagsOutput
	p := iam.NewListInstanceProfileTagsPaginator(client, input)
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

// Lists the instance profiles that have the specified path prefix. If there are
// none, the operation returns an empty list. For more information about instance
// profiles, see [Using instance profiles]in the IAM User Guide.
//
// IAM resource-listing operations return a subset of the available attributes for
// the resource. For example, this operation does not return tags, even though they
// are an attribute of the returned object. To view all of the information for an
// instance profile, see [GetInstanceProfile].
//
// You can paginate the results using the MaxItems and Marker parameters.
//
// [Using instance profiles]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use_switch-role-ec2_instance-profiles.html
// [GetInstanceProfile]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetInstanceProfile.html
func iam_ListInstanceProfiles(cfg aws.Config, client *iam.Client) {
	input := &iam.ListInstanceProfilesInput{}

	if len(_iamMarker) > 0 {
		input.Marker = aws.String(_iamMarker)
	}
	if len(_iamMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _iamMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}
	if len(_iamPathPrefix) > 0 {
		input.PathPrefix = aws.String(_iamPathPrefix)
	}

	if disablePaginator() {
		if resp, err := client.ListInstanceProfiles(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iam.ListInstanceProfilesOutput
	p := iam.NewListInstanceProfilesPaginator(client, input)
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

// Lists the instance profiles that have the specified associated IAM role. If
// there are none, the operation returns an empty list. For more information about
// instance profiles, go to [Using instance profiles]in the IAM User Guide.
//
// You can paginate the results using the MaxItems and Marker parameters.
//
// [Using instance profiles]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use_switch-role-ec2_instance-profiles.html
func iam_ListInstanceProfilesForRole(cfg aws.Config, client *iam.Client) {
	input := &iam.ListInstanceProfilesForRoleInput{
		// RoleName: *string, // Required
	}

	if len(_iamRoleName) > 0 {
		input.RoleName = aws.String(_iamRoleName)
	}
	if len(_iamMarker) > 0 {
		input.Marker = aws.String(_iamMarker)
	}
	if len(_iamMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _iamMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListInstanceProfilesForRole(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iam.ListInstanceProfilesForRoleOutput
	p := iam.NewListInstanceProfilesForRolePaginator(client, input)
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

// Lists the tags that are attached to the specified IAM virtual multi-factor
// authentication (MFA) device. The returned list of tags is sorted by tag key. For
// more information about tagging, see [Tagging IAM resources]in the IAM User Guide.
//
// [Tagging IAM resources]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html
func iam_ListMFADeviceTags(cfg aws.Config, client *iam.Client) {
	input := &iam.ListMFADeviceTagsInput{
		// SerialNumber: *string, // Required
	}

	if len(_iamSerialNumber) > 0 {
		input.SerialNumber = aws.String(_iamSerialNumber)
	}
	if len(_iamMarker) > 0 {
		input.Marker = aws.String(_iamMarker)
	}
	if len(_iamMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _iamMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListMFADeviceTags(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iam.ListMFADeviceTagsOutput
	p := iam.NewListMFADeviceTagsPaginator(client, input)
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

// Lists the MFA devices for an IAM user. If the request includes a IAM user name,
// then this operation lists all the MFA devices associated with the specified
// user. If you do not specify a user name, IAM determines the user name implicitly
// based on the Amazon Web Services access key ID signing the request for this
// operation.
//
// You can paginate the results using the MaxItems and Marker parameters.
func iam_ListMFADevices(cfg aws.Config, client *iam.Client) {
	input := &iam.ListMFADevicesInput{}

	if len(_iamMarker) > 0 {
		input.Marker = aws.String(_iamMarker)
	}
	if len(_iamMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _iamMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}
	if len(_iamUserName) > 0 {
		input.UserName = aws.String(_iamUserName)
	}

	if disablePaginator() {
		if resp, err := client.ListMFADevices(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iam.ListMFADevicesOutput
	p := iam.NewListMFADevicesPaginator(client, input)
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

// Lists the tags that are attached to the specified OpenID Connect
// (OIDC)-compatible identity provider. The returned list of tags is sorted by tag
// key. For more information, see [About web identity federation].
//
// For more information about tagging, see [Tagging IAM resources] in the IAM User Guide.
//
// [Tagging IAM resources]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html
// [About web identity federation]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_providers_oidc.html
func iam_ListOpenIDConnectProviderTags(cfg aws.Config, client *iam.Client) {
	input := &iam.ListOpenIDConnectProviderTagsInput{
		// OpenIDConnectProviderArn: *string, // Required
	}

	if len(_iamOpenIDConnectProviderArn) > 0 {
		input.OpenIDConnectProviderArn = aws.String(_iamOpenIDConnectProviderArn)
	}
	if len(_iamMarker) > 0 {
		input.Marker = aws.String(_iamMarker)
	}
	if len(_iamMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _iamMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListOpenIDConnectProviderTags(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iam.ListOpenIDConnectProviderTagsOutput
	p := iam.NewListOpenIDConnectProviderTagsPaginator(client, input)
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

// Lists information about the IAM OpenID Connect (OIDC) provider resource objects
// defined in the Amazon Web Services account.
//
// IAM resource-listing operations return a subset of the available attributes for
// the resource. For example, this operation does not return tags, even though they
// are an attribute of the returned object. To view all of the information for an
// OIDC provider, see [GetOpenIDConnectProvider].
//
// [GetOpenIDConnectProvider]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetOpenIDConnectProvider.html
func iam_ListOpenIDConnectProviders(cfg aws.Config, client *iam.Client) {
	input := &iam.ListOpenIDConnectProvidersInput{}

	if resp, err := client.ListOpenIDConnectProviders(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the centralized root access features enabled for your organization. For
// more information, see [Centrally manage root access for member accounts].
//
// [Centrally manage root access for member accounts]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_root-user.html#id_root-user-access-management
func iam_ListOrganizationsFeatures(cfg aws.Config, client *iam.Client) {
	input := &iam.ListOrganizationsFeaturesInput{}

	if resp, err := client.ListOrganizationsFeatures(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all the managed policies that are available in your Amazon Web Services
// account, including your own customer-defined managed policies and all Amazon Web
// Services managed policies.
//
// You can filter the list of policies that is returned using the optional
// OnlyAttached , Scope , and PathPrefix parameters. For example, to list only the
// customer managed policies in your Amazon Web Services account, set Scope to
// Local . To list only Amazon Web Services managed policies, set Scope to AWS .
//
// You can paginate the results using the MaxItems and Marker parameters.
//
// For more information about managed policies, see [Managed policies and inline policies] in the IAM User Guide.
//
// IAM resource-listing operations return a subset of the available attributes for
// the resource. For example, this operation does not return tags, even though they
// are an attribute of the returned object. To view all of the information for a
// customer manged policy, see [GetPolicy].
//
// [GetPolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetPolicy.html
// [Managed policies and inline policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html
func iam_ListPolicies(cfg aws.Config, client *iam.Client) {
	input := &iam.ListPoliciesInput{}

	if len(_iamMarker) > 0 {
		input.Marker = aws.String(_iamMarker)
	}
	if len(_iamMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _iamMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}
	if len(_iamOnlyAttached) > 0 {
		if err := assignInputField(input, "OnlyAttached", _iamOnlyAttached); err != nil {
			log.Errorf("invalid --only-attached: %s", err.Error())
			return
		}
	}
	if len(_iamPathPrefix) > 0 {
		input.PathPrefix = aws.String(_iamPathPrefix)
	}
	if len(_iamPolicyUsageFilter) > 0 {
		if err := assignInputField(input, "PolicyUsageFilter", _iamPolicyUsageFilter); err != nil {
			log.Errorf("invalid --policy-usage-filter: %s", err.Error())
			return
		}
	}
	if len(_iamScope) > 0 {
		if err := assignInputField(input, "Scope", _iamScope); err != nil {
			log.Errorf("invalid --scope: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListPolicies(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iam.ListPoliciesOutput
	p := iam.NewListPoliciesPaginator(client, input)
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

// Retrieves a list of policies that the IAM identity (user, group, or role) can
// use to access each specified service.
//
// This operation does not use other policy types when determining whether a
// resource could access a service. These other policy types include resource-based
// policies, access control lists, Organizations policies, IAM permissions
// boundaries, and STS assume role policies. It only applies permissions policy
// logic. For more about the evaluation of policy types, see [Evaluating policies]in the IAM User Guide.
//
// The list of policies returned by the operation depends on the ARN of the
// identity that you provide.
//
// - User – The list of policies includes the managed and inline policies that
// are attached to the user directly. The list also includes any additional managed
// and inline policies that are attached to the group to which the user belongs.
//
// - Group – The list of policies includes only the managed and inline policies
// that are attached to the group directly. Policies that are attached to the
// group’s user are not included.
//
// - Role – The list of policies includes only the managed and inline policies
// that are attached to the role.
//
// For each managed policy, this operation returns the ARN and policy name. For
// each inline policy, it returns the policy name and the entity to which it is
// attached. Inline policies do not have an ARN. For more information about these
// policy types, see [Managed policies and inline policies]in the IAM User Guide.
//
// Policies that are attached to users and roles as permissions boundaries are not
// returned. To view which managed policy is currently used to set the permissions
// boundary for a user or role, use the [GetUser]or [GetRole] operations.
//
// [GetRole]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetRole.html
// [Evaluating policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_evaluation-logic.html#policy-eval-basics
// [Managed policies and inline policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_managed-vs-inline.html
// [GetUser]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetUser.html
func iam_ListPoliciesGrantingServiceAccess(cfg aws.Config, client *iam.Client) {
	input := &iam.ListPoliciesGrantingServiceAccessInput{
		// Arn: *string, // Required
		// ServiceNamespaces: []string, // Required
	}

	if len(_iamArn) > 0 {
		input.Arn = aws.String(_iamArn)
	}
	if len(_iamServiceNamespaces) > 0 {
		input.ServiceNamespaces = append([]string(nil), _iamServiceNamespaces...)
	}
	if len(_iamMarker) > 0 {
		input.Marker = aws.String(_iamMarker)
	}

	if resp, err := client.ListPoliciesGrantingServiceAccess(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the tags that are attached to the specified IAM customer managed policy.
// The returned list of tags is sorted by tag key. For more information about
// tagging, see [Tagging IAM resources]in the IAM User Guide.
//
// [Tagging IAM resources]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html
func iam_ListPolicyTags(cfg aws.Config, client *iam.Client) {
	input := &iam.ListPolicyTagsInput{
		// PolicyArn: *string, // Required
	}

	if len(_iamPolicyArn) > 0 {
		input.PolicyArn = aws.String(_iamPolicyArn)
	}
	if len(_iamMarker) > 0 {
		input.Marker = aws.String(_iamMarker)
	}
	if len(_iamMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _iamMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListPolicyTags(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iam.ListPolicyTagsOutput
	p := iam.NewListPolicyTagsPaginator(client, input)
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

// Lists information about the versions of the specified managed policy, including
// the version that is currently set as the policy's default version.
//
// For more information about managed policies, see [Managed policies and inline policies] in the IAM User Guide.
//
// [Managed policies and inline policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html
func iam_ListPolicyVersions(cfg aws.Config, client *iam.Client) {
	input := &iam.ListPolicyVersionsInput{
		// PolicyArn: *string, // Required
	}

	if len(_iamPolicyArn) > 0 {
		input.PolicyArn = aws.String(_iamPolicyArn)
	}
	if len(_iamMarker) > 0 {
		input.Marker = aws.String(_iamMarker)
	}
	if len(_iamMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _iamMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListPolicyVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iam.ListPolicyVersionsOutput
	p := iam.NewListPolicyVersionsPaginator(client, input)
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

// Lists the names of the inline policies that are embedded in the specified IAM
// role.
//
// An IAM role can also have managed policies attached to it. To list the managed
// policies that are attached to a role, use [ListAttachedRolePolicies]. For more information about
// policies, see [Managed policies and inline policies]in the IAM User Guide.
//
// You can paginate the results using the MaxItems and Marker parameters. If there
// are no inline policies embedded with the specified role, the operation returns
// an empty list.
//
// [ListAttachedRolePolicies]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_ListAttachedRolePolicies.html
// [Managed policies and inline policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html
func iam_ListRolePolicies(cfg aws.Config, client *iam.Client) {
	input := &iam.ListRolePoliciesInput{
		// RoleName: *string, // Required
	}

	if len(_iamRoleName) > 0 {
		input.RoleName = aws.String(_iamRoleName)
	}
	if len(_iamMarker) > 0 {
		input.Marker = aws.String(_iamMarker)
	}
	if len(_iamMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _iamMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListRolePolicies(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iam.ListRolePoliciesOutput
	p := iam.NewListRolePoliciesPaginator(client, input)
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

// Lists the tags that are attached to the specified role. The returned list of
// tags is sorted by tag key. For more information about tagging, see [Tagging IAM resources]in the IAM
// User Guide.
//
// [Tagging IAM resources]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html
func iam_ListRoleTags(cfg aws.Config, client *iam.Client) {
	input := &iam.ListRoleTagsInput{
		// RoleName: *string, // Required
	}

	if len(_iamRoleName) > 0 {
		input.RoleName = aws.String(_iamRoleName)
	}
	if len(_iamMarker) > 0 {
		input.Marker = aws.String(_iamMarker)
	}
	if len(_iamMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _iamMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListRoleTags(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iam.ListRoleTagsOutput
	p := iam.NewListRoleTagsPaginator(client, input)
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

// Lists the IAM roles that have the specified path prefix. If there are none, the
// operation returns an empty list. For more information about roles, see [IAM roles]in the
// IAM User Guide.
//
// IAM resource-listing operations return a subset of the available attributes for
// the resource. This operation does not return the following attributes, even
// though they are an attribute of the returned object:
//
// - PermissionsBoundary
//
// - RoleLastUsed
//
// - Tags
//
// To view all of the information for a role, see [GetRole].
//
// You can paginate the results using the MaxItems and Marker parameters.
//
// [GetRole]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetRole.html
// [IAM roles]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles.html
func iam_ListRoles(cfg aws.Config, client *iam.Client) {
	input := &iam.ListRolesInput{}

	if len(_iamMarker) > 0 {
		input.Marker = aws.String(_iamMarker)
	}
	if len(_iamMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _iamMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}
	if len(_iamPathPrefix) > 0 {
		input.PathPrefix = aws.String(_iamPathPrefix)
	}

	if disablePaginator() {
		if resp, err := client.ListRoles(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iam.ListRolesOutput
	p := iam.NewListRolesPaginator(client, input)
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

// Lists the tags that are attached to the specified Security Assertion Markup
// Language (SAML) identity provider. The returned list of tags is sorted by tag
// key. For more information, see [About SAML 2.0-based federation].
//
// For more information about tagging, see [Tagging IAM resources] in the IAM User Guide.
//
// [About SAML 2.0-based federation]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_providers_saml.html
// [Tagging IAM resources]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html
func iam_ListSAMLProviderTags(cfg aws.Config, client *iam.Client) {
	input := &iam.ListSAMLProviderTagsInput{
		// SAMLProviderArn: *string, // Required
	}

	if len(_iamSAMLProviderArn) > 0 {
		input.SAMLProviderArn = aws.String(_iamSAMLProviderArn)
	}
	if len(_iamMarker) > 0 {
		input.Marker = aws.String(_iamMarker)
	}
	if len(_iamMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _iamMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListSAMLProviderTags(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iam.ListSAMLProviderTagsOutput
	p := iam.NewListSAMLProviderTagsPaginator(client, input)
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

// Lists the SAML provider resource objects defined in IAM in the account. IAM
// resource-listing operations return a subset of the available attributes for the
// resource. For example, this operation does not return tags, even though they are
// an attribute of the returned object. To view all of the information for a SAML
// provider, see [GetSAMLProvider].
//
// This operation requires [Signature Version 4].
//
// [Signature Version 4]: https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html
// [GetSAMLProvider]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetSAMLProvider.html
func iam_ListSAMLProviders(cfg aws.Config, client *iam.Client) {
	input := &iam.ListSAMLProvidersInput{}

	if resp, err := client.ListSAMLProviders(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the tags that are attached to the specified IAM server certificate. The
// returned list of tags is sorted by tag key. For more information about tagging,
// see [Tagging IAM resources]in the IAM User Guide.
//
// For certificates in a Region supported by Certificate Manager (ACM), we
// recommend that you don't use IAM server certificates. Instead, use ACM to
// provision, manage, and deploy your server certificates. For more information
// about IAM server certificates, [Working with server certificates]in the IAM User Guide.
//
// [Working with server certificates]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_server-certs.html
// [Tagging IAM resources]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html
func iam_ListServerCertificateTags(cfg aws.Config, client *iam.Client) {
	input := &iam.ListServerCertificateTagsInput{
		// ServerCertificateName: *string, // Required
	}

	if len(_iamServerCertificateName) > 0 {
		input.ServerCertificateName = aws.String(_iamServerCertificateName)
	}
	if len(_iamMarker) > 0 {
		input.Marker = aws.String(_iamMarker)
	}
	if len(_iamMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _iamMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListServerCertificateTags(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iam.ListServerCertificateTagsOutput
	p := iam.NewListServerCertificateTagsPaginator(client, input)
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

// Lists the server certificates stored in IAM that have the specified path
// prefix. If none exist, the operation returns an empty list.
//
// You can paginate the results using the MaxItems and Marker parameters.
//
// For more information about working with server certificates, see [Working with server certificates] in the IAM
// User Guide. This topic also includes a list of Amazon Web Services services that
// can use the server certificates that you manage with IAM.
//
// IAM resource-listing operations return a subset of the available attributes for
// the resource. For example, this operation does not return tags, even though they
// are an attribute of the returned object. To view all of the information for a
// servercertificate, see [GetServerCertificate].
//
// [GetServerCertificate]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetServerCertificate.html
// [Working with server certificates]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_server-certs.html
func iam_ListServerCertificates(cfg aws.Config, client *iam.Client) {
	input := &iam.ListServerCertificatesInput{}

	if len(_iamMarker) > 0 {
		input.Marker = aws.String(_iamMarker)
	}
	if len(_iamMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _iamMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}
	if len(_iamPathPrefix) > 0 {
		input.PathPrefix = aws.String(_iamPathPrefix)
	}

	if disablePaginator() {
		if resp, err := client.ListServerCertificates(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iam.ListServerCertificatesOutput
	p := iam.NewListServerCertificatesPaginator(client, input)
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

// Returns information about the service-specific credentials associated with the
// specified IAM user. If none exists, the operation returns an empty list. The
// service-specific credentials returned by this operation are used only for
// authenticating the IAM user to a specific service. For more information about
// using service-specific credentials to authenticate to an Amazon Web Services
// service, see [Set up service-specific credentials]in the CodeCommit User Guide.
//
// [Set up service-specific credentials]: https://docs.aws.amazon.com/codecommit/latest/userguide/setting-up-gc.html
func iam_ListServiceSpecificCredentials(cfg aws.Config, client *iam.Client) {
	input := &iam.ListServiceSpecificCredentialsInput{}

	if len(_iamAllUsers) > 0 {
		if err := assignInputField(input, "AllUsers", _iamAllUsers); err != nil {
			log.Errorf("invalid --all-users: %s", err.Error())
			return
		}
	}
	if len(_iamMarker) > 0 {
		input.Marker = aws.String(_iamMarker)
	}
	if len(_iamMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _iamMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}
	if len(_iamServiceName) > 0 {
		input.ServiceName = aws.String(_iamServiceName)
	}
	if len(_iamUserName) > 0 {
		input.UserName = aws.String(_iamUserName)
	}

	if resp, err := client.ListServiceSpecificCredentials(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the signing certificates associated with the
// specified IAM user. If none exists, the operation returns an empty list.
//
// Although each user is limited to a small number of signing certificates, you
// can still paginate the results using the MaxItems and Marker parameters.
//
// If the UserName field is not specified, the user name is determined implicitly
// based on the Amazon Web Services access key ID used to sign the request for this
// operation. This operation works for access keys under the Amazon Web Services
// account. Consequently, you can use this operation to manage Amazon Web Services
// account root user credentials even if the Amazon Web Services account has no
// associated users.
func iam_ListSigningCertificates(cfg aws.Config, client *iam.Client) {
	input := &iam.ListSigningCertificatesInput{}

	if len(_iamMarker) > 0 {
		input.Marker = aws.String(_iamMarker)
	}
	if len(_iamMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _iamMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}
	if len(_iamUserName) > 0 {
		input.UserName = aws.String(_iamUserName)
	}

	if disablePaginator() {
		if resp, err := client.ListSigningCertificates(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iam.ListSigningCertificatesOutput
	p := iam.NewListSigningCertificatesPaginator(client, input)
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

// Returns information about the SSH public keys associated with the specified IAM
// user. If none exists, the operation returns an empty list.
//
// The SSH public keys returned by this operation are used only for authenticating
// the IAM user to an CodeCommit repository. For more information about using SSH
// keys to authenticate to an CodeCommit repository, see [Set up CodeCommit for SSH connections]in the CodeCommit User
// Guide.
//
// Although each user is limited to a small number of keys, you can still paginate
// the results using the MaxItems and Marker parameters.
//
// [Set up CodeCommit for SSH connections]: https://docs.aws.amazon.com/codecommit/latest/userguide/setting-up-credentials-ssh.html
func iam_ListSSHPublicKeys(cfg aws.Config, client *iam.Client) {
	input := &iam.ListSSHPublicKeysInput{}

	if len(_iamMarker) > 0 {
		input.Marker = aws.String(_iamMarker)
	}
	if len(_iamMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _iamMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}
	if len(_iamUserName) > 0 {
		input.UserName = aws.String(_iamUserName)
	}

	if disablePaginator() {
		if resp, err := client.ListSSHPublicKeys(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iam.ListSSHPublicKeysOutput
	p := iam.NewListSSHPublicKeysPaginator(client, input)
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

// Lists the names of the inline policies embedded in the specified IAM user.
// An IAM user can also have managed policies attached to it. To list the managed
// policies that are attached to a user, use [ListAttachedUserPolicies]. For more information about
// policies, see [Managed policies and inline policies]in the IAM User Guide.
//
// You can paginate the results using the MaxItems and Marker parameters. If there
// are no inline policies embedded with the specified user, the operation returns
// an empty list.
//
// [ListAttachedUserPolicies]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_ListAttachedUserPolicies.html
// [Managed policies and inline policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html
func iam_ListUserPolicies(cfg aws.Config, client *iam.Client) {
	input := &iam.ListUserPoliciesInput{
		// UserName: *string, // Required
	}

	if len(_iamUserName) > 0 {
		input.UserName = aws.String(_iamUserName)
	}
	if len(_iamMarker) > 0 {
		input.Marker = aws.String(_iamMarker)
	}
	if len(_iamMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _iamMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListUserPolicies(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iam.ListUserPoliciesOutput
	p := iam.NewListUserPoliciesPaginator(client, input)
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

// Lists the tags that are attached to the specified IAM user. The returned list
// of tags is sorted by tag key. For more information about tagging, see [Tagging IAM resources]in the
// IAM User Guide.
//
// [Tagging IAM resources]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html
func iam_ListUserTags(cfg aws.Config, client *iam.Client) {
	input := &iam.ListUserTagsInput{
		// UserName: *string, // Required
	}

	if len(_iamUserName) > 0 {
		input.UserName = aws.String(_iamUserName)
	}
	if len(_iamMarker) > 0 {
		input.Marker = aws.String(_iamMarker)
	}
	if len(_iamMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _iamMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListUserTags(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iam.ListUserTagsOutput
	p := iam.NewListUserTagsPaginator(client, input)
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

// Lists the IAM users that have the specified path prefix. If no path prefix is
// specified, the operation returns all users in the Amazon Web Services account.
// If there are none, the operation returns an empty list.
//
// IAM resource-listing operations return a subset of the available attributes for
// the resource. This operation does not return the following attributes, even
// though they are an attribute of the returned object:
//
// - PermissionsBoundary
//
// - Tags
//
// To view all of the information for a user, see [GetUser].
//
// You can paginate the results using the MaxItems and Marker parameters.
//
// [GetUser]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetUser.html
func iam_ListUsers(cfg aws.Config, client *iam.Client) {
	input := &iam.ListUsersInput{}

	if len(_iamMarker) > 0 {
		input.Marker = aws.String(_iamMarker)
	}
	if len(_iamMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _iamMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}
	if len(_iamPathPrefix) > 0 {
		input.PathPrefix = aws.String(_iamPathPrefix)
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

	var results []*iam.ListUsersOutput
	p := iam.NewListUsersPaginator(client, input)
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

// Lists the virtual MFA devices defined in the Amazon Web Services account by
// assignment status. If you do not specify an assignment status, the operation
// returns a list of all virtual MFA devices. Assignment status can be Assigned ,
// Unassigned , or Any .
//
// IAM resource-listing operations return a subset of the available attributes for
// the resource. For example, this operation does not return tags, even though they
// are an attribute of the returned object. To view tag information for a virtual
// MFA device, see [ListMFADeviceTags].
//
// You can paginate the results using the MaxItems and Marker parameters.
//
// [ListMFADeviceTags]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_ListMFADeviceTags.html
func iam_ListVirtualMFADevices(cfg aws.Config, client *iam.Client) {
	input := &iam.ListVirtualMFADevicesInput{}

	if len(_iamAssignmentStatus) > 0 {
		if err := assignInputField(input, "AssignmentStatus", _iamAssignmentStatus); err != nil {
			log.Errorf("invalid --assignment-status: %s", err.Error())
			return
		}
	}
	if len(_iamMarker) > 0 {
		input.Marker = aws.String(_iamMarker)
	}
	if len(_iamMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _iamMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListVirtualMFADevices(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iam.ListVirtualMFADevicesOutput
	p := iam.NewListVirtualMFADevicesPaginator(client, input)
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

// Adds or updates an inline policy document that is embedded in the specified IAM
// group.
//
// A user can also have managed policies attached to it. To attach a managed
// policy to a group, use [AttachGroupPolicy]AttachGroupPolicy . To create a new managed policy, use [CreatePolicy]
// CreatePolicy . For information about policies, see [Managed policies and inline policies] in the IAM User Guide.
//
// For information about the maximum number of inline policies that you can embed
// in a group, see [IAM and STS quotas]in the IAM User Guide.
//
// Because policy documents can be large, you should use POST rather than GET when
// calling PutGroupPolicy . For general information about using the Query API with
// IAM, see [Making query requests]in the IAM User Guide.
//
// [CreatePolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_CreatePolicy.html
// [IAM and STS quotas]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-quotas.html
// [Making query requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/IAM_UsingQueryAPI.html
// [AttachGroupPolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_AttachGroupPolicy.html
// [Managed policies and inline policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html
func iam_PutGroupPolicy(cfg aws.Config, client *iam.Client) {
	input := &iam.PutGroupPolicyInput{
		// GroupName: *string, // Required
		// PolicyDocument: *string, // Required
		// PolicyName: *string, // Required
	}

	if len(_iamGroupName) > 0 {
		input.GroupName = aws.String(_iamGroupName)
	}
	if len(_iamPolicyDocument) > 0 {
		input.PolicyDocument = aws.String(_iamPolicyDocument)
	}
	if len(_iamPolicyName) > 0 {
		input.PolicyName = aws.String(_iamPolicyName)
	}

	if resp, err := client.PutGroupPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds or updates the policy that is specified as the IAM role's permissions
// boundary. You can use an Amazon Web Services managed policy or a customer
// managed policy to set the boundary for a role. Use the boundary to control the
// maximum permissions that the role can have. Setting a permissions boundary is an
// advanced feature that can affect the permissions for the role.
//
// You cannot set the boundary for a service-linked role.
//
// Policies used as permissions boundaries do not provide permissions. You must
// also attach a permissions policy to the role. To learn how the effective
// permissions for a role are evaluated, see [IAM JSON policy evaluation logic]in the IAM User Guide.
//
// [IAM JSON policy evaluation logic]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_evaluation-logic.html
func iam_PutRolePermissionsBoundary(cfg aws.Config, client *iam.Client) {
	input := &iam.PutRolePermissionsBoundaryInput{
		// PermissionsBoundary: *string, // Required
		// RoleName: *string, // Required
	}

	if len(_iamPermissionsBoundary) > 0 {
		input.PermissionsBoundary = aws.String(_iamPermissionsBoundary)
	}
	if len(_iamRoleName) > 0 {
		input.RoleName = aws.String(_iamRoleName)
	}

	if resp, err := client.PutRolePermissionsBoundary(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds or updates an inline policy document that is embedded in the specified IAM
// role.
//
// When you embed an inline policy in a role, the inline policy is used as part of
// the role's access (permissions) policy. The role's trust policy is created at
// the same time as the role, using [CreateRole]CreateRole . You can update a role's trust
// policy using [UpdateAssumeRolePolicy]UpdateAssumeRolePolicy . For more information about roles, see [IAM roles] in
// the IAM User Guide.
//
// A role can also have a managed policy attached to it. To attach a managed
// policy to a role, use [AttachRolePolicy]AttachRolePolicy . To create a new managed policy, use [CreatePolicy]
// CreatePolicy . For information about policies, see [Managed policies and inline policies] in the IAM User Guide.
//
// For information about the maximum number of inline policies that you can embed
// with a role, see [IAM and STS quotas]in the IAM User Guide.
//
// Because policy documents can be large, you should use POST rather than GET when
// calling PutRolePolicy . For general information about using the Query API with
// IAM, see [Making query requests]in the IAM User Guide.
//
// [UpdateAssumeRolePolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_UpdateAssumeRolePolicy.html
// [AttachRolePolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_AttachRolePolicy.html
// [CreatePolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_CreatePolicy.html
// [IAM and STS quotas]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-quotas.html
// [Making query requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/IAM_UsingQueryAPI.html
// [IAM roles]: https://docs.aws.amazon.com/IAM/latest/UserGuide/roles-toplevel.html
// [CreateRole]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_CreateRole.html
// [Managed policies and inline policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html
func iam_PutRolePolicy(cfg aws.Config, client *iam.Client) {
	input := &iam.PutRolePolicyInput{
		// PolicyDocument: *string, // Required
		// PolicyName: *string, // Required
		// RoleName: *string, // Required
	}

	if len(_iamPolicyDocument) > 0 {
		input.PolicyDocument = aws.String(_iamPolicyDocument)
	}
	if len(_iamPolicyName) > 0 {
		input.PolicyName = aws.String(_iamPolicyName)
	}
	if len(_iamRoleName) > 0 {
		input.RoleName = aws.String(_iamRoleName)
	}

	if resp, err := client.PutRolePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds or updates the policy that is specified as the IAM user's permissions
// boundary. You can use an Amazon Web Services managed policy or a customer
// managed policy to set the boundary for a user. Use the boundary to control the
// maximum permissions that the user can have. Setting a permissions boundary is an
// advanced feature that can affect the permissions for the user.
//
// Policies that are used as permissions boundaries do not provide permissions.
// You must also attach a permissions policy to the user. To learn how the
// effective permissions for a user are evaluated, see [IAM JSON policy evaluation logic]in the IAM User Guide.
//
// [IAM JSON policy evaluation logic]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_evaluation-logic.html
func iam_PutUserPermissionsBoundary(cfg aws.Config, client *iam.Client) {
	input := &iam.PutUserPermissionsBoundaryInput{
		// PermissionsBoundary: *string, // Required
		// UserName: *string, // Required
	}

	if len(_iamPermissionsBoundary) > 0 {
		input.PermissionsBoundary = aws.String(_iamPermissionsBoundary)
	}
	if len(_iamUserName) > 0 {
		input.UserName = aws.String(_iamUserName)
	}

	if resp, err := client.PutUserPermissionsBoundary(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds or updates an inline policy document that is embedded in the specified IAM
// user.
//
// An IAM user can also have a managed policy attached to it. To attach a managed
// policy to a user, use [AttachUserPolicy]AttachUserPolicy . To create a new managed policy, use [CreatePolicy]
// CreatePolicy . For information about policies, see [Managed policies and inline policies] in the IAM User Guide.
//
// For information about the maximum number of inline policies that you can embed
// in a user, see [IAM and STS quotas]in the IAM User Guide.
//
// Because policy documents can be large, you should use POST rather than GET when
// calling PutUserPolicy . For general information about using the Query API with
// IAM, see [Making query requests]in the IAM User Guide.
//
// [CreatePolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_CreatePolicy.html
// [IAM and STS quotas]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-quotas.html
// [Making query requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/IAM_UsingQueryAPI.html
// [AttachUserPolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_AttachUserPolicy.html
// [Managed policies and inline policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html
func iam_PutUserPolicy(cfg aws.Config, client *iam.Client) {
	input := &iam.PutUserPolicyInput{
		// PolicyDocument: *string, // Required
		// PolicyName: *string, // Required
		// UserName: *string, // Required
	}

	if len(_iamPolicyDocument) > 0 {
		input.PolicyDocument = aws.String(_iamPolicyDocument)
	}
	if len(_iamPolicyName) > 0 {
		input.PolicyName = aws.String(_iamPolicyName)
	}
	if len(_iamUserName) > 0 {
		input.UserName = aws.String(_iamUserName)
	}

	if resp, err := client.PutUserPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Rejects a delegation request, denying the requested temporary access.
// Once a request is rejected, it cannot be accepted or updated later. Rejected
// requests expire after 7 days.
//
// When rejecting a request, an optional explanation can be added using the Notes
// request parameter.
//
// For more details, see [Managing Permissions for Delegation Requests].
//
// [Managing Permissions for Delegation Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies-temporary-delegation.html#temporary-delegation-managing-permissions
func iam_RejectDelegationRequest(cfg aws.Config, client *iam.Client) {
	input := &iam.RejectDelegationRequestInput{
		// DelegationRequestId: *string, // Required
	}

	if len(_iamDelegationRequestId) > 0 {
		input.DelegationRequestId = aws.String(_iamDelegationRequestId)
	}
	if len(_iamNotes) > 0 {
		input.Notes = aws.String(_iamNotes)
	}

	if resp, err := client.RejectDelegationRequest(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the specified client ID (also known as audience) from the list of
// client IDs registered for the specified IAM OpenID Connect (OIDC) provider
// resource object.
//
// This operation is idempotent; it does not fail or return an error if you try to
// remove a client ID that does not exist.
func iam_RemoveClientIDFromOpenIDConnectProvider(cfg aws.Config, client *iam.Client) {
	input := &iam.RemoveClientIDFromOpenIDConnectProviderInput{
		// ClientID: *string, // Required
		// OpenIDConnectProviderArn: *string, // Required
	}

	if len(_iamClientID) > 0 {
		input.ClientID = aws.String(_iamClientID)
	}
	if len(_iamOpenIDConnectProviderArn) > 0 {
		input.OpenIDConnectProviderArn = aws.String(_iamOpenIDConnectProviderArn)
	}

	if resp, err := client.RemoveClientIDFromOpenIDConnectProvider(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the specified IAM role from the specified Amazon EC2 instance profile.
// Make sure that you do not have any Amazon EC2 instances running with the role
// you are about to remove from the instance profile. Removing a role from an
// instance profile that is associated with a running instance might break any
// applications running on the instance.
//
// For more information about roles, see [IAM roles] in the IAM User Guide. For more
// information about instance profiles, see [Using instance profiles]in the IAM User Guide.
//
// [Using instance profiles]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use_switch-role-ec2_instance-profiles.html
// [IAM roles]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles.html
func iam_RemoveRoleFromInstanceProfile(cfg aws.Config, client *iam.Client) {
	input := &iam.RemoveRoleFromInstanceProfileInput{
		// InstanceProfileName: *string, // Required
		// RoleName: *string, // Required
	}

	if len(_iamInstanceProfileName) > 0 {
		input.InstanceProfileName = aws.String(_iamInstanceProfileName)
	}
	if len(_iamRoleName) > 0 {
		input.RoleName = aws.String(_iamRoleName)
	}

	if resp, err := client.RemoveRoleFromInstanceProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the specified user from the specified group.
func iam_RemoveUserFromGroup(cfg aws.Config, client *iam.Client) {
	input := &iam.RemoveUserFromGroupInput{
		// GroupName: *string, // Required
		// UserName: *string, // Required
	}

	if len(_iamGroupName) > 0 {
		input.GroupName = aws.String(_iamGroupName)
	}
	if len(_iamUserName) > 0 {
		input.UserName = aws.String(_iamUserName)
	}

	if resp, err := client.RemoveUserFromGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Resets the password for a service-specific credential. The new password is
// Amazon Web Services generated and cryptographically strong. It cannot be
// configured by the user. Resetting the password immediately invalidates the
// previous password associated with this user.
func iam_ResetServiceSpecificCredential(cfg aws.Config, client *iam.Client) {
	input := &iam.ResetServiceSpecificCredentialInput{
		// ServiceSpecificCredentialId: *string, // Required
	}

	if len(_iamServiceSpecificCredentialId) > 0 {
		input.ServiceSpecificCredentialId = aws.String(_iamServiceSpecificCredentialId)
	}
	if len(_iamUserName) > 0 {
		input.UserName = aws.String(_iamUserName)
	}

	if resp, err := client.ResetServiceSpecificCredential(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Synchronizes the specified MFA device with its IAM resource object on the
// Amazon Web Services servers.
//
// For more information about creating and working with virtual MFA devices, see [Using a virtual MFA device]
// in the IAM User Guide.
//
// [Using a virtual MFA device]: https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_VirtualMFA.html
func iam_ResyncMFADevice(cfg aws.Config, client *iam.Client) {
	input := &iam.ResyncMFADeviceInput{
		// AuthenticationCode1: *string, // Required
		// AuthenticationCode2: *string, // Required
		// SerialNumber: *string, // Required
		// UserName: *string, // Required
	}

	if len(_iamAuthenticationCode1) > 0 {
		input.AuthenticationCode1 = aws.String(_iamAuthenticationCode1)
	}
	if len(_iamAuthenticationCode2) > 0 {
		input.AuthenticationCode2 = aws.String(_iamAuthenticationCode2)
	}
	if len(_iamSerialNumber) > 0 {
		input.SerialNumber = aws.String(_iamSerialNumber)
	}
	if len(_iamUserName) > 0 {
		input.UserName = aws.String(_iamUserName)
	}

	if resp, err := client.ResyncMFADevice(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sends the exchange token for an accepted delegation request.
// The exchange token is sent to the partner via an asynchronous notification
// channel, established by the partner.
//
// The delegation request must be in the ACCEPTED state when calling this API.
// After the SendDelegationToken API call is successful, the request transitions
// to a FINALIZED state and cannot be rolled back. However, a user may reject an
// accepted request before the SendDelegationToken API is called.
//
// For more details, see [Managing Permissions for Delegation Requests].
//
// [Managing Permissions for Delegation Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies-temporary-delegation.html#temporary-delegation-managing-permissions
func iam_SendDelegationToken(cfg aws.Config, client *iam.Client) {
	input := &iam.SendDelegationTokenInput{
		// DelegationRequestId: *string, // Required
	}

	if len(_iamDelegationRequestId) > 0 {
		input.DelegationRequestId = aws.String(_iamDelegationRequestId)
	}

	if resp, err := client.SendDelegationToken(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets the specified version of the specified policy as the policy's default
// (operative) version.
//
// This operation affects all users, groups, and roles that the policy is attached
// to. To list the users, groups, and roles that the policy is attached to, use [ListEntitiesForPolicy].
//
// For information about managed policies, see [Managed policies and inline policies] in the IAM User Guide.
//
// [ListEntitiesForPolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_ListEntitiesForPolicy.html
// [Managed policies and inline policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html
func iam_SetDefaultPolicyVersion(cfg aws.Config, client *iam.Client) {
	input := &iam.SetDefaultPolicyVersionInput{
		// PolicyArn: *string, // Required
		// VersionId: *string, // Required
	}

	if len(_iamPolicyArn) > 0 {
		input.PolicyArn = aws.String(_iamPolicyArn)
	}
	if len(_iamVersionId) > 0 {
		input.VersionId = aws.String(_iamVersionId)
	}

	if resp, err := client.SetDefaultPolicyVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets the specified version of the global endpoint token as the token version
// used for the Amazon Web Services account.
//
// By default, Security Token Service (STS) is available as a global service, and
// all STS requests go to a single endpoint at https://sts.amazonaws.com . Amazon
// Web Services recommends using Regional STS endpoints to reduce latency, build in
// redundancy, and increase session token availability. For information about
// Regional endpoints for STS, see [Security Token Service endpoints and quotas]in the Amazon Web Services General Reference.
//
// If you make an STS call to the global endpoint, the resulting session tokens
// might be valid in some Regions but not others. It depends on the version that is
// set in this operation. Version 1 tokens are valid only in Amazon Web Services
// Regions that are available by default. These tokens do not work in manually
// enabled Regions, such as Asia Pacific (Hong Kong). Version 2 tokens are valid in
// all Regions. However, version 2 tokens are longer and might affect systems where
// you temporarily store tokens. For information, see [Activating and deactivating STS in an Amazon Web Services Region]in the IAM User Guide.
//
// To view the current session token version, see the GlobalEndpointTokenVersion
// entry in the response of the [GetAccountSummary]operation.
//
// [Activating and deactivating STS in an Amazon Web Services Region]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_enable-regions.html
// [GetAccountSummary]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetAccountSummary.html
// [Security Token Service endpoints and quotas]: https://docs.aws.amazon.com/general/latest/gr/sts.html
func iam_SetSecurityTokenServicePreferences(cfg aws.Config, client *iam.Client) {
	input := &iam.SetSecurityTokenServicePreferencesInput{
		// GlobalEndpointTokenVersion: types.GlobalEndpointTokenVersion, // Required
	}

	if len(_iamGlobalEndpointTokenVersion) > 0 {
		if err := assignInputField(input, "GlobalEndpointTokenVersion", _iamGlobalEndpointTokenVersion); err != nil {
			log.Errorf("invalid --global-endpoint-token-version: %s", err.Error())
			return
		}
	}

	if resp, err := client.SetSecurityTokenServicePreferences(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Simulate how a set of IAM policies and optionally a resource-based policy works
// with a list of API operations and Amazon Web Services resources to determine the
// policies' effective permissions. The policies are provided as strings.
//
// The simulation does not perform the API operations; it only checks the
// authorization to determine if the simulated policies allow or deny the
// operations. You can simulate resources that don't exist in your account.
//
// If you want to simulate existing policies that are attached to an IAM user,
// group, or role, use [SimulatePrincipalPolicy]instead.
//
// Context keys are variables that are maintained by Amazon Web Services and its
// services and which provide details about the context of an API query request.
// You can use the Condition element of an IAM policy to evaluate context keys. To
// get the list of context keys that the policies require for correct simulation,
// use [GetContextKeysForCustomPolicy].
//
// If the output is long, you can use MaxItems and Marker parameters to paginate
// the results.
//
// The IAM policy simulator evaluates statements in the identity-based policy and
// the inputs that you provide during simulation. The policy simulator results can
// differ from your live Amazon Web Services environment. We recommend that you
// check your policies against your live Amazon Web Services environment after
// testing using the policy simulator to confirm that you have the desired results.
// For more information about using the policy simulator, see [Testing IAM policies with the IAM policy simulator]in the IAM User
// Guide.
//
// [GetContextKeysForCustomPolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetContextKeysForCustomPolicy.html
// [Testing IAM policies with the IAM policy simulator]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_testing-policies.html
// [SimulatePrincipalPolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_SimulatePrincipalPolicy.html
func iam_SimulateCustomPolicy(cfg aws.Config, client *iam.Client) {
	input := &iam.SimulateCustomPolicyInput{
		// ActionNames: []string, // Required
		// PolicyInputList: []string, // Required
	}

	if len(_iamActionNames) > 0 {
		input.ActionNames = append([]string(nil), _iamActionNames...)
	}
	if len(_iamPolicyInputList) > 0 {
		input.PolicyInputList = append([]string(nil), _iamPolicyInputList...)
	}
	if len(_iamCallerArn) > 0 {
		input.CallerArn = aws.String(_iamCallerArn)
	}
	if len(_iamContextEntries) > 0 {
		if err := assignInputField(input, "ContextEntries", _iamContextEntries); err != nil {
			log.Errorf("invalid --context-entries: %s", err.Error())
			return
		}
	}
	if len(_iamMarker) > 0 {
		input.Marker = aws.String(_iamMarker)
	}
	if len(_iamMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _iamMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}
	if len(_iamPermissionsBoundaryPolicyInputList) > 0 {
		input.PermissionsBoundaryPolicyInputList = append([]string(nil), _iamPermissionsBoundaryPolicyInputList...)
	}
	if len(_iamResourceArns) > 0 {
		input.ResourceArns = append([]string(nil), _iamResourceArns...)
	}
	if len(_iamResourceHandlingOption) > 0 {
		input.ResourceHandlingOption = aws.String(_iamResourceHandlingOption)
	}
	if len(_iamResourceOwner) > 0 {
		input.ResourceOwner = aws.String(_iamResourceOwner)
	}
	if len(_iamResourcePolicy) > 0 {
		input.ResourcePolicy = aws.String(_iamResourcePolicy)
	}

	if disablePaginator() {
		if resp, err := client.SimulateCustomPolicy(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iam.SimulateCustomPolicyOutput
	p := iam.NewSimulateCustomPolicyPaginator(client, input)
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

// Simulate how a set of IAM policies attached to an IAM entity works with a list
// of API operations and Amazon Web Services resources to determine the policies'
// effective permissions. The entity can be an IAM user, group, or role. If you
// specify a user, then the simulation also includes all of the policies that are
// attached to groups that the user belongs to. You can simulate resources that
// don't exist in your account.
//
// You can optionally include a list of one or more additional policies specified
// as strings to include in the simulation. If you want to simulate only policies
// specified as strings, use [SimulateCustomPolicy]instead.
//
// You can also optionally include one resource-based policy to be evaluated with
// each of the resources included in the simulation for IAM users only.
//
// The simulation does not perform the API operations; it only checks the
// authorization to determine if the simulated policies allow or deny the
// operations.
//
// Note: This operation discloses information about the permissions granted to
// other users. If you do not want users to see other user's permissions, then
// consider allowing them to use [SimulateCustomPolicy]instead.
//
// Context keys are variables maintained by Amazon Web Services and its services
// that provide details about the context of an API query request. You can use the
// Condition element of an IAM policy to evaluate context keys. To get the list of
// context keys that the policies require for correct simulation, use [GetContextKeysForPrincipalPolicy].
//
// If the output is long, you can use the MaxItems and Marker parameters to
// paginate the results.
//
// The IAM policy simulator evaluates statements in the identity-based policy and
// the inputs that you provide during simulation. The policy simulator results can
// differ from your live Amazon Web Services environment. We recommend that you
// check your policies against your live Amazon Web Services environment after
// testing using the policy simulator to confirm that you have the desired results.
// For more information about using the policy simulator, see [Testing IAM policies with the IAM policy simulator]in the IAM User
// Guide.
//
// [GetContextKeysForPrincipalPolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetContextKeysForPrincipalPolicy.html
// [Testing IAM policies with the IAM policy simulator]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_testing-policies.html
// [SimulateCustomPolicy]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_SimulateCustomPolicy.html
func iam_SimulatePrincipalPolicy(cfg aws.Config, client *iam.Client) {
	input := &iam.SimulatePrincipalPolicyInput{
		// ActionNames: []string, // Required
		// PolicySourceArn: *string, // Required
	}

	if len(_iamActionNames) > 0 {
		input.ActionNames = append([]string(nil), _iamActionNames...)
	}
	if len(_iamPolicySourceArn) > 0 {
		input.PolicySourceArn = aws.String(_iamPolicySourceArn)
	}
	if len(_iamCallerArn) > 0 {
		input.CallerArn = aws.String(_iamCallerArn)
	}
	if len(_iamContextEntries) > 0 {
		if err := assignInputField(input, "ContextEntries", _iamContextEntries); err != nil {
			log.Errorf("invalid --context-entries: %s", err.Error())
			return
		}
	}
	if len(_iamMarker) > 0 {
		input.Marker = aws.String(_iamMarker)
	}
	if len(_iamMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _iamMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}
	if len(_iamPermissionsBoundaryPolicyInputList) > 0 {
		input.PermissionsBoundaryPolicyInputList = append([]string(nil), _iamPermissionsBoundaryPolicyInputList...)
	}
	if len(_iamPolicyInputList) > 0 {
		input.PolicyInputList = append([]string(nil), _iamPolicyInputList...)
	}
	if len(_iamResourceArns) > 0 {
		input.ResourceArns = append([]string(nil), _iamResourceArns...)
	}
	if len(_iamResourceHandlingOption) > 0 {
		input.ResourceHandlingOption = aws.String(_iamResourceHandlingOption)
	}
	if len(_iamResourceOwner) > 0 {
		input.ResourceOwner = aws.String(_iamResourceOwner)
	}
	if len(_iamResourcePolicy) > 0 {
		input.ResourcePolicy = aws.String(_iamResourcePolicy)
	}

	if disablePaginator() {
		if resp, err := client.SimulatePrincipalPolicy(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iam.SimulatePrincipalPolicyOutput
	p := iam.NewSimulatePrincipalPolicyPaginator(client, input)
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

// Adds one or more tags to an IAM instance profile. If a tag with the same key
// name already exists, then that tag is overwritten with the new value.
//
// Each tag consists of a key name and an associated value. By assigning tags to
// your resources, you can do the following:
//
// - Administrative grouping and discovery - Attach tags to resources to aid in
// organization and search. For example, you could search for all resources with
// the key name Project and the value MyImportantProject. Or search for all
// resources with the key name Cost Center and the value 41200.
//
// - Access control - Include tags in IAM user-based and resource-based
// policies. You can use tags to restrict access to only an IAM instance profile
// that has a specified tag attached. For examples of policies that show how to use
// tags to control access, see [Control access using IAM tags]in the IAM User Guide.
//
// - If any one of the tags is invalid or if you exceed the allowed maximum
// number of tags, then the entire request fails and the resource is not created.
// For more information about tagging, see [Tagging IAM resources]in the IAM User Guide.
//
// - Amazon Web Services always interprets the tag Value as a single string. If
// you need to store an array, you can store comma-separated values in the string.
// However, you must interpret the value in your code.
//
// [Control access using IAM tags]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_tags.html
// [Tagging IAM resources]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html
func iam_TagInstanceProfile(cfg aws.Config, client *iam.Client) {
	input := &iam.TagInstanceProfileInput{
		// InstanceProfileName: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_iamInstanceProfileName) > 0 {
		input.InstanceProfileName = aws.String(_iamInstanceProfileName)
	}
	if len(_iamTags) > 0 {
		if err := assignInputField(input, "Tags", _iamTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.TagInstanceProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds one or more tags to an IAM virtual multi-factor authentication (MFA)
// device. If a tag with the same key name already exists, then that tag is
// overwritten with the new value.
//
// A tag consists of a key name and an associated value. By assigning tags to your
// resources, you can do the following:
//
// - Administrative grouping and discovery - Attach tags to resources to aid in
// organization and search. For example, you could search for all resources with
// the key name Project and the value MyImportantProject. Or search for all
// resources with the key name Cost Center and the value 41200.
//
// - Access control - Include tags in IAM user-based and resource-based
// policies. You can use tags to restrict access to only an IAM virtual MFA device
// that has a specified tag attached. For examples of policies that show how to use
// tags to control access, see [Control access using IAM tags]in the IAM User Guide.
//
// - If any one of the tags is invalid or if you exceed the allowed maximum
// number of tags, then the entire request fails and the resource is not created.
// For more information about tagging, see [Tagging IAM resources]in the IAM User Guide.
//
// - Amazon Web Services always interprets the tag Value as a single string. If
// you need to store an array, you can store comma-separated values in the string.
// However, you must interpret the value in your code.
//
// [Control access using IAM tags]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_tags.html
// [Tagging IAM resources]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html
func iam_TagMFADevice(cfg aws.Config, client *iam.Client) {
	input := &iam.TagMFADeviceInput{
		// SerialNumber: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_iamSerialNumber) > 0 {
		input.SerialNumber = aws.String(_iamSerialNumber)
	}
	if len(_iamTags) > 0 {
		if err := assignInputField(input, "Tags", _iamTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.TagMFADevice(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds one or more tags to an OpenID Connect (OIDC)-compatible identity provider.
// For more information about these providers, see [About web identity federation]. If a tag with the same key
// name already exists, then that tag is overwritten with the new value.
//
// A tag consists of a key name and an associated value. By assigning tags to your
// resources, you can do the following:
//
// - Administrative grouping and discovery - Attach tags to resources to aid in
// organization and search. For example, you could search for all resources with
// the key name Project and the value MyImportantProject. Or search for all
// resources with the key name Cost Center and the value 41200.
//
// - Access control - Include tags in IAM identity-based and resource-based
// policies. You can use tags to restrict access to only an OIDC provider that has
// a specified tag attached. For examples of policies that show how to use tags to
// control access, see [Control access using IAM tags]in the IAM User Guide.
//
// - If any one of the tags is invalid or if you exceed the allowed maximum
// number of tags, then the entire request fails and the resource is not created.
// For more information about tagging, see [Tagging IAM resources]in the IAM User Guide.
//
// - Amazon Web Services always interprets the tag Value as a single string. If
// you need to store an array, you can store comma-separated values in the string.
// However, you must interpret the value in your code.
//
// [Control access using IAM tags]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_tags.html
// [Tagging IAM resources]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html
// [About web identity federation]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_providers_oidc.html
func iam_TagOpenIDConnectProvider(cfg aws.Config, client *iam.Client) {
	input := &iam.TagOpenIDConnectProviderInput{
		// OpenIDConnectProviderArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_iamOpenIDConnectProviderArn) > 0 {
		input.OpenIDConnectProviderArn = aws.String(_iamOpenIDConnectProviderArn)
	}
	if len(_iamTags) > 0 {
		if err := assignInputField(input, "Tags", _iamTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.TagOpenIDConnectProvider(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds one or more tags to an IAM customer managed policy. If a tag with the same
// key name already exists, then that tag is overwritten with the new value.
//
// A tag consists of a key name and an associated value. By assigning tags to your
// resources, you can do the following:
//
// - Administrative grouping and discovery - Attach tags to resources to aid in
// organization and search. For example, you could search for all resources with
// the key name Project and the value MyImportantProject. Or search for all
// resources with the key name Cost Center and the value 41200.
//
// - Access control - Include tags in IAM user-based and resource-based
// policies. You can use tags to restrict access to only an IAM customer managed
// policy that has a specified tag attached. For examples of policies that show how
// to use tags to control access, see [Control access using IAM tags]in the IAM User Guide.
//
// - If any one of the tags is invalid or if you exceed the allowed maximum
// number of tags, then the entire request fails and the resource is not created.
// For more information about tagging, see [Tagging IAM resources]in the IAM User Guide.
//
// - Amazon Web Services always interprets the tag Value as a single string. If
// you need to store an array, you can store comma-separated values in the string.
// However, you must interpret the value in your code.
//
// [Control access using IAM tags]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_tags.html
// [Tagging IAM resources]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html
func iam_TagPolicy(cfg aws.Config, client *iam.Client) {
	input := &iam.TagPolicyInput{
		// PolicyArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_iamPolicyArn) > 0 {
		input.PolicyArn = aws.String(_iamPolicyArn)
	}
	if len(_iamTags) > 0 {
		if err := assignInputField(input, "Tags", _iamTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.TagPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds one or more tags to an IAM role. The role can be a regular role or a
// service-linked role. If a tag with the same key name already exists, then that
// tag is overwritten with the new value.
//
// A tag consists of a key name and an associated value. By assigning tags to your
// resources, you can do the following:
//
// - Administrative grouping and discovery - Attach tags to resources to aid in
// organization and search. For example, you could search for all resources with
// the key name Project and the value MyImportantProject. Or search for all
// resources with the key name Cost Center and the value 41200.
//
// - Access control - Include tags in IAM user-based and resource-based
// policies. You can use tags to restrict access to only an IAM role that has a
// specified tag attached. You can also restrict access to only those resources
// that have a certain tag attached. For examples of policies that show how to use
// tags to control access, see [Control access using IAM tags]in the IAM User Guide.
//
// - Cost allocation - Use tags to help track which individuals and teams are
// using which Amazon Web Services resources.
//
// - If any one of the tags is invalid or if you exceed the allowed maximum
// number of tags, then the entire request fails and the resource is not created.
// For more information about tagging, see [Tagging IAM resources]in the IAM User Guide.
//
// - Amazon Web Services always interprets the tag Value as a single string. If
// you need to store an array, you can store comma-separated values in the string.
// However, you must interpret the value in your code.
//
// For more information about tagging, see [Tagging IAM identities] in the IAM User Guide.
//
// [Control access using IAM tags]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_tags.html
// [Tagging IAM resources]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html
// [Tagging IAM identities]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html
func iam_TagRole(cfg aws.Config, client *iam.Client) {
	input := &iam.TagRoleInput{
		// RoleName: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_iamRoleName) > 0 {
		input.RoleName = aws.String(_iamRoleName)
	}
	if len(_iamTags) > 0 {
		if err := assignInputField(input, "Tags", _iamTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.TagRole(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds one or more tags to a Security Assertion Markup Language (SAML) identity
// provider. For more information about these providers, see [About SAML 2.0-based federation]. If a tag with the
// same key name already exists, then that tag is overwritten with the new value.
//
// A tag consists of a key name and an associated value. By assigning tags to your
// resources, you can do the following:
//
// - Administrative grouping and discovery - Attach tags to resources to aid in
// organization and search. For example, you could search for all resources with
// the key name Project and the value MyImportantProject. Or search for all
// resources with the key name Cost Center and the value 41200.
//
// - Access control - Include tags in IAM user-based and resource-based
// policies. You can use tags to restrict access to only a SAML identity provider
// that has a specified tag attached. For examples of policies that show how to use
// tags to control access, see [Control access using IAM tags]in the IAM User Guide.
//
// - If any one of the tags is invalid or if you exceed the allowed maximum
// number of tags, then the entire request fails and the resource is not created.
// For more information about tagging, see [Tagging IAM resources]in the IAM User Guide.
//
// - Amazon Web Services always interprets the tag Value as a single string. If
// you need to store an array, you can store comma-separated values in the string.
// However, you must interpret the value in your code.
//
// [Control access using IAM tags]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_tags.html
// [About SAML 2.0-based federation]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_providers_saml.html
// [Tagging IAM resources]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html
func iam_TagSAMLProvider(cfg aws.Config, client *iam.Client) {
	input := &iam.TagSAMLProviderInput{
		// SAMLProviderArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_iamSAMLProviderArn) > 0 {
		input.SAMLProviderArn = aws.String(_iamSAMLProviderArn)
	}
	if len(_iamTags) > 0 {
		if err := assignInputField(input, "Tags", _iamTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.TagSAMLProvider(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds one or more tags to an IAM server certificate. If a tag with the same key
// name already exists, then that tag is overwritten with the new value.
//
// For certificates in a Region supported by Certificate Manager (ACM), we
// recommend that you don't use IAM server certificates. Instead, use ACM to
// provision, manage, and deploy your server certificates. For more information
// about IAM server certificates, [Working with server certificates]in the IAM User Guide.
//
// A tag consists of a key name and an associated value. By assigning tags to your
// resources, you can do the following:
//
// - Administrative grouping and discovery - Attach tags to resources to aid in
// organization and search. For example, you could search for all resources with
// the key name Project and the value MyImportantProject. Or search for all
// resources with the key name Cost Center and the value 41200.
//
// - Access control - Include tags in IAM user-based and resource-based
// policies. You can use tags to restrict access to only a server certificate that
// has a specified tag attached. For examples of policies that show how to use tags
// to control access, see [Control access using IAM tags]in the IAM User Guide.
//
// - Cost allocation - Use tags to help track which individuals and teams are
// using which Amazon Web Services resources.
//
// - If any one of the tags is invalid or if you exceed the allowed maximum
// number of tags, then the entire request fails and the resource is not created.
// For more information about tagging, see [Tagging IAM resources]in the IAM User Guide.
//
// - Amazon Web Services always interprets the tag Value as a single string. If
// you need to store an array, you can store comma-separated values in the string.
// However, you must interpret the value in your code.
//
// [Control access using IAM tags]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_tags.html
// [Working with server certificates]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_server-certs.html
// [Tagging IAM resources]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html
func iam_TagServerCertificate(cfg aws.Config, client *iam.Client) {
	input := &iam.TagServerCertificateInput{
		// ServerCertificateName: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_iamServerCertificateName) > 0 {
		input.ServerCertificateName = aws.String(_iamServerCertificateName)
	}
	if len(_iamTags) > 0 {
		if err := assignInputField(input, "Tags", _iamTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.TagServerCertificate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds one or more tags to an IAM user. If a tag with the same key name already
// exists, then that tag is overwritten with the new value.
//
// A tag consists of a key name and an associated value. By assigning tags to your
// resources, you can do the following:
//
// - Administrative grouping and discovery - Attach tags to resources to aid in
// organization and search. For example, you could search for all resources with
// the key name Project and the value MyImportantProject. Or search for all
// resources with the key name Cost Center and the value 41200.
//
// - Access control - Include tags in IAM identity-based and resource-based
// policies. You can use tags to restrict access to only an IAM requesting user
// that has a specified tag attached. You can also restrict access to only those
// resources that have a certain tag attached. For examples of policies that show
// how to use tags to control access, see [Control access using IAM tags]in the IAM User Guide.
//
// - Cost allocation - Use tags to help track which individuals and teams are
// using which Amazon Web Services resources.
//
// - If any one of the tags is invalid or if you exceed the allowed maximum
// number of tags, then the entire request fails and the resource is not created.
// For more information about tagging, see [Tagging IAM resources]in the IAM User Guide.
//
// - Amazon Web Services always interprets the tag Value as a single string. If
// you need to store an array, you can store comma-separated values in the string.
// However, you must interpret the value in your code.
//
// For more information about tagging, see [Tagging IAM identities] in the IAM User Guide.
//
// [Control access using IAM tags]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_tags.html
// [Tagging IAM resources]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html
// [Tagging IAM identities]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html
func iam_TagUser(cfg aws.Config, client *iam.Client) {
	input := &iam.TagUserInput{
		// Tags: []types.Tag, // Required
		// UserName: *string, // Required
	}

	if len(_iamTags) > 0 {
		if err := assignInputField(input, "Tags", _iamTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_iamUserName) > 0 {
		input.UserName = aws.String(_iamUserName)
	}

	if resp, err := client.TagUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the specified tags from the IAM instance profile. For more information
// about tagging, see [Tagging IAM resources]in the IAM User Guide.
//
// [Tagging IAM resources]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html
func iam_UntagInstanceProfile(cfg aws.Config, client *iam.Client) {
	input := &iam.UntagInstanceProfileInput{
		// InstanceProfileName: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_iamInstanceProfileName) > 0 {
		input.InstanceProfileName = aws.String(_iamInstanceProfileName)
	}
	if len(_iamTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _iamTagKeys...)
	}

	if resp, err := client.UntagInstanceProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the specified tags from the IAM virtual multi-factor authentication
// (MFA) device. For more information about tagging, see [Tagging IAM resources]in the IAM User Guide.
//
// [Tagging IAM resources]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html
func iam_UntagMFADevice(cfg aws.Config, client *iam.Client) {
	input := &iam.UntagMFADeviceInput{
		// SerialNumber: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_iamSerialNumber) > 0 {
		input.SerialNumber = aws.String(_iamSerialNumber)
	}
	if len(_iamTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _iamTagKeys...)
	}

	if resp, err := client.UntagMFADevice(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the specified tags from the specified OpenID Connect (OIDC)-compatible
// identity provider in IAM. For more information about OIDC providers, see [About web identity federation]. For
// more information about tagging, see [Tagging IAM resources]in the IAM User Guide.
//
// [Tagging IAM resources]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html
// [About web identity federation]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_providers_oidc.html
func iam_UntagOpenIDConnectProvider(cfg aws.Config, client *iam.Client) {
	input := &iam.UntagOpenIDConnectProviderInput{
		// OpenIDConnectProviderArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_iamOpenIDConnectProviderArn) > 0 {
		input.OpenIDConnectProviderArn = aws.String(_iamOpenIDConnectProviderArn)
	}
	if len(_iamTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _iamTagKeys...)
	}

	if resp, err := client.UntagOpenIDConnectProvider(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the specified tags from the customer managed policy. For more
// information about tagging, see [Tagging IAM resources]in the IAM User Guide.
//
// [Tagging IAM resources]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html
func iam_UntagPolicy(cfg aws.Config, client *iam.Client) {
	input := &iam.UntagPolicyInput{
		// PolicyArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_iamPolicyArn) > 0 {
		input.PolicyArn = aws.String(_iamPolicyArn)
	}
	if len(_iamTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _iamTagKeys...)
	}

	if resp, err := client.UntagPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the specified tags from the role. For more information about tagging,
// see [Tagging IAM resources]in the IAM User Guide.
//
// [Tagging IAM resources]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html
func iam_UntagRole(cfg aws.Config, client *iam.Client) {
	input := &iam.UntagRoleInput{
		// RoleName: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_iamRoleName) > 0 {
		input.RoleName = aws.String(_iamRoleName)
	}
	if len(_iamTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _iamTagKeys...)
	}

	if resp, err := client.UntagRole(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the specified tags from the specified Security Assertion Markup
// Language (SAML) identity provider in IAM. For more information about these
// providers, see [About web identity federation]. For more information about tagging, see [Tagging IAM resources] in the IAM User Guide.
//
// [Tagging IAM resources]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html
// [About web identity federation]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_providers_oidc.html
func iam_UntagSAMLProvider(cfg aws.Config, client *iam.Client) {
	input := &iam.UntagSAMLProviderInput{
		// SAMLProviderArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_iamSAMLProviderArn) > 0 {
		input.SAMLProviderArn = aws.String(_iamSAMLProviderArn)
	}
	if len(_iamTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _iamTagKeys...)
	}

	if resp, err := client.UntagSAMLProvider(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the specified tags from the IAM server certificate. For more
// information about tagging, see [Tagging IAM resources]in the IAM User Guide.
//
// For certificates in a Region supported by Certificate Manager (ACM), we
// recommend that you don't use IAM server certificates. Instead, use ACM to
// provision, manage, and deploy your server certificates. For more information
// about IAM server certificates, [Working with server certificates]in the IAM User Guide.
//
// [Working with server certificates]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_server-certs.html
// [Tagging IAM resources]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html
func iam_UntagServerCertificate(cfg aws.Config, client *iam.Client) {
	input := &iam.UntagServerCertificateInput{
		// ServerCertificateName: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_iamServerCertificateName) > 0 {
		input.ServerCertificateName = aws.String(_iamServerCertificateName)
	}
	if len(_iamTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _iamTagKeys...)
	}

	if resp, err := client.UntagServerCertificate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the specified tags from the user. For more information about tagging,
// see [Tagging IAM resources]in the IAM User Guide.
//
// [Tagging IAM resources]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html
func iam_UntagUser(cfg aws.Config, client *iam.Client) {
	input := &iam.UntagUserInput{
		// TagKeys: []string, // Required
		// UserName: *string, // Required
	}

	if len(_iamTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _iamTagKeys...)
	}
	if len(_iamUserName) > 0 {
		input.UserName = aws.String(_iamUserName)
	}

	if resp, err := client.UntagUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Changes the status of the specified access key from Active to Inactive, or vice
// versa. This operation can be used to disable a user's key as part of a key
// rotation workflow.
//
// If the UserName is not specified, the user name is determined implicitly based
// on the Amazon Web Services access key ID used to sign the request. If a
// temporary access key is used, then UserName is required. If a long-term key is
// assigned to the user, then UserName is not required. This operation works for
// access keys under the Amazon Web Services account. Consequently, you can use
// this operation to manage Amazon Web Services account root user credentials even
// if the Amazon Web Services account has no associated users.
//
// For information about rotating keys, see [Managing keys and certificates] in the IAM User Guide.
//
// [Managing keys and certificates]: https://docs.aws.amazon.com/IAM/latest/UserGuide/ManagingCredentials.html
func iam_UpdateAccessKey(cfg aws.Config, client *iam.Client) {
	input := &iam.UpdateAccessKeyInput{
		// AccessKeyId: *string, // Required
		// Status: types.StatusType, // Required
	}

	if len(_iamAccessKeyId) > 0 {
		input.AccessKeyId = aws.String(_iamAccessKeyId)
	}
	if len(_iamStatus) > 0 {
		if err := assignInputField(input, "Status", _iamStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_iamUserName) > 0 {
		input.UserName = aws.String(_iamUserName)
	}

	if resp, err := client.UpdateAccessKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the password policy settings for the Amazon Web Services account.
// This operation does not support partial updates. No parameters are required,
// but if you do not specify a parameter, that parameter's value reverts to its
// default value. See the Request Parameters section for each parameter's default
// value. Also note that some parameters do not allow the default parameter to be
// explicitly set. Instead, to invoke the default value, do not include that
// parameter when you invoke the operation.
//
// For more information about using a password policy, see [Managing an IAM password policy] in the IAM User Guide.
//
// [Managing an IAM password policy]: https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_ManagingPasswordPolicies.html
func iam_UpdateAccountPasswordPolicy(cfg aws.Config, client *iam.Client) {
	input := &iam.UpdateAccountPasswordPolicyInput{}

	if len(_iamAllowUsersToChangePassword) > 0 {
		if err := assignInputField(input, "AllowUsersToChangePassword", _iamAllowUsersToChangePassword); err != nil {
			log.Errorf("invalid --allow-users-to-change-password: %s", err.Error())
			return
		}
	}
	if len(_iamHardExpiry) > 0 {
		if err := assignInputField(input, "HardExpiry", _iamHardExpiry); err != nil {
			log.Errorf("invalid --hard-expiry: %s", err.Error())
			return
		}
	}
	if len(_iamMaxPasswordAge) > 0 {
		if err := assignInputField(input, "MaxPasswordAge", _iamMaxPasswordAge); err != nil {
			log.Errorf("invalid --max-password-age: %s", err.Error())
			return
		}
	}
	if len(_iamMinimumPasswordLength) > 0 {
		if err := assignInputField(input, "MinimumPasswordLength", _iamMinimumPasswordLength); err != nil {
			log.Errorf("invalid --minimum-password-length: %s", err.Error())
			return
		}
	}
	if len(_iamPasswordReusePrevention) > 0 {
		if err := assignInputField(input, "PasswordReusePrevention", _iamPasswordReusePrevention); err != nil {
			log.Errorf("invalid --password-reuse-prevention: %s", err.Error())
			return
		}
	}
	if len(_iamRequireLowercaseCharacters) > 0 {
		if err := assignInputField(input, "RequireLowercaseCharacters", _iamRequireLowercaseCharacters); err != nil {
			log.Errorf("invalid --require-lowercase-characters: %s", err.Error())
			return
		}
	}
	if len(_iamRequireNumbers) > 0 {
		if err := assignInputField(input, "RequireNumbers", _iamRequireNumbers); err != nil {
			log.Errorf("invalid --require-numbers: %s", err.Error())
			return
		}
	}
	if len(_iamRequireSymbols) > 0 {
		if err := assignInputField(input, "RequireSymbols", _iamRequireSymbols); err != nil {
			log.Errorf("invalid --require-symbols: %s", err.Error())
			return
		}
	}
	if len(_iamRequireUppercaseCharacters) > 0 {
		if err := assignInputField(input, "RequireUppercaseCharacters", _iamRequireUppercaseCharacters); err != nil {
			log.Errorf("invalid --require-uppercase-characters: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateAccountPasswordPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the policy that grants an IAM entity permission to assume a role. This
// is typically referred to as the "role trust policy". For more information about
// roles, see [Using roles to delegate permissions and federate identities].
//
// [Using roles to delegate permissions and federate identities]: https://docs.aws.amazon.com/IAM/latest/UserGuide/roles-toplevel.html
func iam_UpdateAssumeRolePolicy(cfg aws.Config, client *iam.Client) {
	input := &iam.UpdateAssumeRolePolicyInput{
		// PolicyDocument: *string, // Required
		// RoleName: *string, // Required
	}

	if len(_iamPolicyDocument) > 0 {
		input.PolicyDocument = aws.String(_iamPolicyDocument)
	}
	if len(_iamRoleName) > 0 {
		input.RoleName = aws.String(_iamRoleName)
	}

	if resp, err := client.UpdateAssumeRolePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing delegation request with additional information. When the
// delegation request is updated, it reaches the PENDING_APPROVAL state.
//
// Once a delegation request has an owner, that owner gets a default permission to
// update the delegation request. For more details, see [Managing Permissions for Delegation Requests].
//
// [Managing Permissions for Delegation Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies-temporary-delegation.html#temporary-delegation-managing-permissions
func iam_UpdateDelegationRequest(cfg aws.Config, client *iam.Client) {
	input := &iam.UpdateDelegationRequestInput{
		// DelegationRequestId: *string, // Required
	}

	if len(_iamDelegationRequestId) > 0 {
		input.DelegationRequestId = aws.String(_iamDelegationRequestId)
	}
	if len(_iamNotes) > 0 {
		input.Notes = aws.String(_iamNotes)
	}

	if resp, err := client.UpdateDelegationRequest(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the name and/or the path of the specified IAM group.
// You should understand the implications of changing a group's path or name. For
// more information, see [Renaming users and groups]in the IAM User Guide.
//
// The person making the request (the principal), must have permission to change
// the role group with the old name and the new name. For example, to change the
// group named Managers to MGRs , the principal must have a policy that allows them
// to update both groups. If the principal has permission to update the Managers
// group, but not the MGRs group, then the update fails. For more information
// about permissions, see [Access management].
//
// [Access management]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access.html
// [Renaming users and groups]: https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_WorkingWithGroupsAndUsers.html
func iam_UpdateGroup(cfg aws.Config, client *iam.Client) {
	input := &iam.UpdateGroupInput{
		// GroupName: *string, // Required
	}

	if len(_iamGroupName) > 0 {
		input.GroupName = aws.String(_iamGroupName)
	}
	if len(_iamNewGroupName) > 0 {
		input.NewGroupName = aws.String(_iamNewGroupName)
	}
	if len(_iamNewPath) > 0 {
		input.NewPath = aws.String(_iamNewPath)
	}

	if resp, err := client.UpdateGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Changes the password for the specified IAM user. You can use the CLI, the
// Amazon Web Services API, or the Users page in the IAM console to change the
// password for any IAM user. Use [ChangePassword]to change your own password in the My Security
// Credentials page in the Amazon Web Services Management Console.
//
// For more information about modifying passwords, see [Managing passwords] in the IAM User Guide.
//
// [ChangePassword]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_ChangePassword.html
// [Managing passwords]: https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_ManagingLogins.html
func iam_UpdateLoginProfile(cfg aws.Config, client *iam.Client) {
	input := &iam.UpdateLoginProfileInput{
		// UserName: *string, // Required
	}

	if len(_iamUserName) > 0 {
		input.UserName = aws.String(_iamUserName)
	}
	if len(_iamPassword) > 0 {
		input.Password = aws.String(_iamPassword)
	}
	if len(_iamPasswordResetRequired) > 0 {
		if err := assignInputField(input, "PasswordResetRequired", _iamPasswordResetRequired); err != nil {
			log.Errorf("invalid --password-reset-required: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateLoginProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Replaces the existing list of server certificate thumbprints associated with an
// OpenID Connect (OIDC) provider resource object with a new list of thumbprints.
//
// The list that you pass with this operation completely replaces the existing
// list of thumbprints. (The lists are not merged.)
//
// Typically, you need to update a thumbprint only when the identity provider
// certificate changes, which occurs rarely. However, if the provider's certificate
// does change, any attempt to assume an IAM role that specifies the OIDC provider
// as a principal fails until the certificate thumbprint is updated.
//
// Amazon Web Services secures communication with OIDC identity providers (IdPs)
// using our library of trusted root certificate authorities (CAs) to verify the
// JSON Web Key Set (JWKS) endpoint's TLS certificate. If your OIDC IdP relies on a
// certificate that is not signed by one of these trusted CAs, only then we secure
// communication using the thumbprints set in the IdP's configuration.
//
// Trust for the OIDC provider is derived from the provider certificate and is
// validated by the thumbprint. Therefore, it is best to limit access to the
// UpdateOpenIDConnectProviderThumbprint operation to highly privileged users.
func iam_UpdateOpenIDConnectProviderThumbprint(cfg aws.Config, client *iam.Client) {
	input := &iam.UpdateOpenIDConnectProviderThumbprintInput{
		// OpenIDConnectProviderArn: *string, // Required
		// ThumbprintList: []string, // Required
	}

	if len(_iamOpenIDConnectProviderArn) > 0 {
		input.OpenIDConnectProviderArn = aws.String(_iamOpenIDConnectProviderArn)
	}
	if len(_iamThumbprintList) > 0 {
		input.ThumbprintList = append([]string(nil), _iamThumbprintList...)
	}

	if resp, err := client.UpdateOpenIDConnectProviderThumbprint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the description or maximum session duration setting of a role.
func iam_UpdateRole(cfg aws.Config, client *iam.Client) {
	input := &iam.UpdateRoleInput{
		// RoleName: *string, // Required
	}

	if len(_iamRoleName) > 0 {
		input.RoleName = aws.String(_iamRoleName)
	}
	if len(_iamDescription) > 0 {
		input.Description = aws.String(_iamDescription)
	}
	if len(_iamMaxSessionDuration) > 0 {
		if err := assignInputField(input, "MaxSessionDuration", _iamMaxSessionDuration); err != nil {
			log.Errorf("invalid --max-session-duration: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateRole(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Use [UpdateRole] instead.
// Modifies only the description of a role. This operation performs the same
// function as the Description parameter in the UpdateRole operation.
//
// [UpdateRole]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_UpdateRole.html
func iam_UpdateRoleDescription(cfg aws.Config, client *iam.Client) {
	input := &iam.UpdateRoleDescriptionInput{
		// Description: *string, // Required
		// RoleName: *string, // Required
	}

	if len(_iamDescription) > 0 {
		input.Description = aws.String(_iamDescription)
	}
	if len(_iamRoleName) > 0 {
		input.RoleName = aws.String(_iamRoleName)
	}

	if resp, err := client.UpdateRoleDescription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the metadata document, SAML encryption settings, and private keys for
// an existing SAML provider. To rotate private keys, add your new private key and
// then remove the old key in a separate request.
func iam_UpdateSAMLProvider(cfg aws.Config, client *iam.Client) {
	input := &iam.UpdateSAMLProviderInput{
		// SAMLProviderArn: *string, // Required
	}

	if len(_iamSAMLProviderArn) > 0 {
		input.SAMLProviderArn = aws.String(_iamSAMLProviderArn)
	}
	if len(_iamAddPrivateKey) > 0 {
		input.AddPrivateKey = aws.String(_iamAddPrivateKey)
	}
	if len(_iamAssertionEncryptionMode) > 0 {
		if err := assignInputField(input, "AssertionEncryptionMode", _iamAssertionEncryptionMode); err != nil {
			log.Errorf("invalid --assertion-encryption-mode: %s", err.Error())
			return
		}
	}
	if len(_iamRemovePrivateKey) > 0 {
		input.RemovePrivateKey = aws.String(_iamRemovePrivateKey)
	}
	if len(_iamSAMLMetadataDocument) > 0 {
		input.SAMLMetadataDocument = aws.String(_iamSAMLMetadataDocument)
	}

	if resp, err := client.UpdateSAMLProvider(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the name and/or the path of the specified server certificate stored in
// IAM.
//
// For more information about working with server certificates, see [Working with server certificates] in the IAM
// User Guide. This topic also includes a list of Amazon Web Services services that
// can use the server certificates that you manage with IAM.
//
// You should understand the implications of changing a server certificate's path
// or name. For more information, see [Renaming a server certificate]in the IAM User Guide.
//
// The person making the request (the principal), must have permission to change
// the server certificate with the old name and the new name. For example, to
// change the certificate named ProductionCert to ProdCert , the principal must
// have a policy that allows them to update both certificates. If the principal has
// permission to update the ProductionCert group, but not the ProdCert
// certificate, then the update fails. For more information about permissions, see [Access management]
// in the IAM User Guide.
//
// [Renaming a server certificate]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_server-certs_manage.html#RenamingServerCerts
// [Access management]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access.html
// [Working with server certificates]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_server-certs.html
func iam_UpdateServerCertificate(cfg aws.Config, client *iam.Client) {
	input := &iam.UpdateServerCertificateInput{
		// ServerCertificateName: *string, // Required
	}

	if len(_iamServerCertificateName) > 0 {
		input.ServerCertificateName = aws.String(_iamServerCertificateName)
	}
	if len(_iamNewPath) > 0 {
		input.NewPath = aws.String(_iamNewPath)
	}
	if len(_iamNewServerCertificateName) > 0 {
		input.NewServerCertificateName = aws.String(_iamNewServerCertificateName)
	}

	if resp, err := client.UpdateServerCertificate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets the status of a service-specific credential to Active or Inactive .
// Service-specific credentials that are inactive cannot be used for authentication
// to the service. This operation can be used to disable a user's service-specific
// credential as part of a credential rotation work flow.
func iam_UpdateServiceSpecificCredential(cfg aws.Config, client *iam.Client) {
	input := &iam.UpdateServiceSpecificCredentialInput{
		// ServiceSpecificCredentialId: *string, // Required
		// Status: types.StatusType, // Required
	}

	if len(_iamServiceSpecificCredentialId) > 0 {
		input.ServiceSpecificCredentialId = aws.String(_iamServiceSpecificCredentialId)
	}
	if len(_iamStatus) > 0 {
		if err := assignInputField(input, "Status", _iamStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_iamUserName) > 0 {
		input.UserName = aws.String(_iamUserName)
	}

	if resp, err := client.UpdateServiceSpecificCredential(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Changes the status of the specified user signing certificate from active to
// disabled, or vice versa. This operation can be used to disable an IAM user's
// signing certificate as part of a certificate rotation work flow.
//
// If the UserName field is not specified, the user name is determined implicitly
// based on the Amazon Web Services access key ID used to sign the request. This
// operation works for access keys under the Amazon Web Services account.
// Consequently, you can use this operation to manage Amazon Web Services account
// root user credentials even if the Amazon Web Services account has no associated
// users.
func iam_UpdateSigningCertificate(cfg aws.Config, client *iam.Client) {
	input := &iam.UpdateSigningCertificateInput{
		// CertificateId: *string, // Required
		// Status: types.StatusType, // Required
	}

	if len(_iamCertificateId) > 0 {
		input.CertificateId = aws.String(_iamCertificateId)
	}
	if len(_iamStatus) > 0 {
		if err := assignInputField(input, "Status", _iamStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_iamUserName) > 0 {
		input.UserName = aws.String(_iamUserName)
	}

	if resp, err := client.UpdateSigningCertificate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets the status of an IAM user's SSH public key to active or inactive. SSH
// public keys that are inactive cannot be used for authentication. This operation
// can be used to disable a user's SSH public key as part of a key rotation work
// flow.
//
// The SSH public key affected by this operation is used only for authenticating
// the associated IAM user to an CodeCommit repository. For more information about
// using SSH keys to authenticate to an CodeCommit repository, see [Set up CodeCommit for SSH connections]in the
// CodeCommit User Guide.
//
// [Set up CodeCommit for SSH connections]: https://docs.aws.amazon.com/codecommit/latest/userguide/setting-up-credentials-ssh.html
func iam_UpdateSSHPublicKey(cfg aws.Config, client *iam.Client) {
	input := &iam.UpdateSSHPublicKeyInput{
		// SSHPublicKeyId: *string, // Required
		// Status: types.StatusType, // Required
		// UserName: *string, // Required
	}

	if len(_iamSSHPublicKeyId) > 0 {
		input.SSHPublicKeyId = aws.String(_iamSSHPublicKeyId)
	}
	if len(_iamStatus) > 0 {
		if err := assignInputField(input, "Status", _iamStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_iamUserName) > 0 {
		input.UserName = aws.String(_iamUserName)
	}

	if resp, err := client.UpdateSSHPublicKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the name and/or the path of the specified IAM user.
// You should understand the implications of changing an IAM user's path or name.
// For more information, see [Renaming an IAM user]and [Renaming an IAM group] in the IAM User Guide.
//
// To change a user name, the requester must have appropriate permissions on both
// the source object and the target object. For example, to change Bob to Robert,
// the entity making the request must have permission on Bob and Robert, or must
// have permission on all (*). For more information about permissions, see [Permissions and policies].
//
// [Renaming an IAM user]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_users_manage.html#id_users_renaming
// [Renaming an IAM group]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_groups_manage_rename.html
// [Permissions and policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/PermissionsAndPolicies.html
func iam_UpdateUser(cfg aws.Config, client *iam.Client) {
	input := &iam.UpdateUserInput{
		// UserName: *string, // Required
	}

	if len(_iamUserName) > 0 {
		input.UserName = aws.String(_iamUserName)
	}
	if len(_iamNewPath) > 0 {
		input.NewPath = aws.String(_iamNewPath)
	}
	if len(_iamNewUserName) > 0 {
		input.NewUserName = aws.String(_iamNewUserName)
	}

	if resp, err := client.UpdateUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Uploads a server certificate entity for the Amazon Web Services account. The
// server certificate entity includes a public key certificate, a private key, and
// an optional certificate chain, which should all be PEM-encoded.
//
// We recommend that you use [Certificate Manager] to provision, manage, and deploy your server
// certificates. With ACM you can request a certificate, deploy it to Amazon Web
// Services resources, and let ACM handle certificate renewals for you.
// Certificates provided by ACM are free. For more information about using ACM, see
// the [Certificate Manager User Guide].
//
// For more information about working with server certificates, see [Working with server certificates] in the IAM
// User Guide. This topic includes a list of Amazon Web Services services that can
// use the server certificates that you manage with IAM.
//
// For information about the number of server certificates you can upload, see [IAM and STS quotas] in
// the IAM User Guide.
//
// Because the body of the public key certificate, private key, and the
// certificate chain can be large, you should use POST rather than GET when calling
// UploadServerCertificate . For information about setting up signatures and
// authorization through the API, see [Signing Amazon Web Services API requests]in the Amazon Web Services General
// Reference. For general information about using the Query API with IAM, see [Calling the API by making HTTP query requests]in
// the IAM User Guide.
//
// [Certificate Manager]: https://docs.aws.amazon.com/acm/
// [Certificate Manager User Guide]: https://docs.aws.amazon.com/acm/latest/userguide/
// [IAM and STS quotas]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-quotas.html
// [Working with server certificates]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_server-certs.html
// [Signing Amazon Web Services API requests]: https://docs.aws.amazon.com/general/latest/gr/signing_aws_api_requests.html
// [Calling the API by making HTTP query requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/programming.html
func iam_UploadServerCertificate(cfg aws.Config, client *iam.Client) {
	input := &iam.UploadServerCertificateInput{
		// CertificateBody: *string, // Required
		// PrivateKey: *string, // Required
		// ServerCertificateName: *string, // Required
	}

	if len(_iamCertificateBody) > 0 {
		input.CertificateBody = aws.String(_iamCertificateBody)
	}
	if len(_iamPrivateKey) > 0 {
		input.PrivateKey = aws.String(_iamPrivateKey)
	}
	if len(_iamServerCertificateName) > 0 {
		input.ServerCertificateName = aws.String(_iamServerCertificateName)
	}
	if len(_iamCertificateChain) > 0 {
		input.CertificateChain = aws.String(_iamCertificateChain)
	}
	if len(_iamPath) > 0 {
		input.Path = aws.String(_iamPath)
	}
	if len(_iamTags) > 0 {
		if err := assignInputField(input, "Tags", _iamTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.UploadServerCertificate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Uploads an X.509 signing certificate and associates it with the specified IAM
// user. Some Amazon Web Services services require you to use certificates to
// validate requests that are signed with a corresponding private key. When you
// upload the certificate, its default status is Active .
//
// For information about when you would use an X.509 signing certificate, see [Managing server certificates in IAM] in
// the IAM User Guide.
//
// If the UserName is not specified, the IAM user name is determined implicitly
// based on the Amazon Web Services access key ID used to sign the request. This
// operation works for access keys under the Amazon Web Services account.
// Consequently, you can use this operation to manage Amazon Web Services account
// root user credentials even if the Amazon Web Services account has no associated
// users.
//
// Because the body of an X.509 certificate can be large, you should use POST
// rather than GET when calling UploadSigningCertificate . For information about
// setting up signatures and authorization through the API, see [Signing Amazon Web Services API requests]in the Amazon Web
// Services General Reference. For general information about using the Query API
// with IAM, see [Making query requests]in the IAM User Guide.
//
// [Managing server certificates in IAM]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_server-certs.html
// [Making query requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/IAM_UsingQueryAPI.html
// [Signing Amazon Web Services API requests]: https://docs.aws.amazon.com/general/latest/gr/signing_aws_api_requests.html
func iam_UploadSigningCertificate(cfg aws.Config, client *iam.Client) {
	input := &iam.UploadSigningCertificateInput{
		// CertificateBody: *string, // Required
	}

	if len(_iamCertificateBody) > 0 {
		input.CertificateBody = aws.String(_iamCertificateBody)
	}
	if len(_iamUserName) > 0 {
		input.UserName = aws.String(_iamUserName)
	}

	if resp, err := client.UploadSigningCertificate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Uploads an SSH public key and associates it with the specified IAM user.
// The SSH public key uploaded by this operation can be used only for
// authenticating the associated IAM user to an CodeCommit repository. For more
// information about using SSH keys to authenticate to an CodeCommit repository,
// see [Set up CodeCommit for SSH connections]in the CodeCommit User Guide.
//
// [Set up CodeCommit for SSH connections]: https://docs.aws.amazon.com/codecommit/latest/userguide/setting-up-credentials-ssh.html
func iam_UploadSSHPublicKey(cfg aws.Config, client *iam.Client) {
	input := &iam.UploadSSHPublicKeyInput{
		// SSHPublicKeyBody: *string, // Required
		// UserName: *string, // Required
	}

	if len(_iamSSHPublicKeyBody) > 0 {
		input.SSHPublicKeyBody = aws.String(_iamSSHPublicKeyBody)
	}
	if len(_iamUserName) > 0 {
		input.UserName = aws.String(_iamUserName)
	}

	if resp, err := client.UploadSSHPublicKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_iamCmd)
	_iamCmd.Flags().SortFlags = false

	_iamCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_iamCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_iamCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_iamCmd.Flags().StringVarP(&_iamAccessKeyId, "access-key-id", "", "", "Access Key ID")
	_iamCmd.Flags().StringVarP(&_iamAccountAlias, "account-alias", "", "", "Account Alias")
	_iamCmd.Flags().StringSliceVarP(&_iamActionNames, "action-names", "", nil, "Action Names")
	_iamCmd.Flags().StringVarP(&_iamAddPrivateKey, "add-private-key", "", "", "Add Private Key")
	_iamCmd.Flags().StringVarP(&_iamAllUsers, "all-users", "", "", "All Users")
	_iamCmd.Flags().StringVarP(&_iamAllowUsersToChangePassword, "allow-users-to-change-password", "", "", "Allow Users To Change Password")
	_iamCmd.Flags().StringVarP(&_iamArn, "arn", "", "", "ARN")
	_iamCmd.Flags().StringVarP(&_iamAssertionEncryptionMode, "assertion-encryption-mode", "", "", "Assertion Encryption Mode")
	_iamCmd.Flags().StringVarP(&_iamAssignmentStatus, "assignment-status", "", "", "Assignment Status")
	_iamCmd.Flags().StringVarP(&_iamAssumeRolePolicyDocument, "assume-role-policy-document", "", "", "Assume Role Policy Document")
	_iamCmd.Flags().StringVarP(&_iamAuthenticationCode1, "authentication-code1", "", "", "Authentication Code1")
	_iamCmd.Flags().StringVarP(&_iamAuthenticationCode2, "authentication-code2", "", "", "Authentication Code2")
	_iamCmd.Flags().StringVarP(&_iamAWSServiceName, "aws-service-name", "", "", "AWS Service Name")
	_iamCmd.Flags().StringVarP(&_iamCallerArn, "caller-arn", "", "", "Caller ARN")
	_iamCmd.Flags().StringVarP(&_iamCertificateBody, "certificate-body", "", "", "Certificate Body")
	_iamCmd.Flags().StringVarP(&_iamCertificateChain, "certificate-chain", "", "", "Certificate Chain")
	_iamCmd.Flags().StringVarP(&_iamCertificateId, "certificate-id", "", "", "Certificate ID")
	_iamCmd.Flags().StringVarP(&_iamClientID, "client-id", "", "", "Client ID")
	_iamCmd.Flags().StringSliceVarP(&_iamClientIDList, "client-id-list", "", nil, "Client ID List")
	_iamCmd.Flags().StringVarP(&_iamContextEntries, "context-entries", "", "", "Context Entries")
	_iamCmd.Flags().StringVarP(&_iamCredentialAgeDays, "credential-age-days", "", "", "Credential Age Days")
	_iamCmd.Flags().StringVarP(&_iamCustomSuffix, "custom-suffix", "", "", "Custom Suffix")
	_iamCmd.Flags().StringVarP(&_iamDelegationPermissionCheck, "delegation-permission-check", "", "", "Delegation Permission Check")
	_iamCmd.Flags().StringVarP(&_iamDelegationRequestId, "delegation-request-id", "", "", "Delegation Request ID")
	_iamCmd.Flags().StringVarP(&_iamDeletionTaskId, "deletion-task-id", "", "", "Deletion Task ID")
	_iamCmd.Flags().StringVarP(&_iamDescription, "description", "", "", "Description")
	_iamCmd.Flags().StringVarP(&_iamEncoding, "encoding", "", "", "Encoding")
	_iamCmd.Flags().StringVarP(&_iamEntityArn, "entity-arn", "", "", "Entity ARN")
	_iamCmd.Flags().StringVarP(&_iamEntityFilter, "entity-filter", "", "", "Entity Filter")
	_iamCmd.Flags().StringVarP(&_iamEntityPath, "entity-path", "", "", "Entity Path")
	_iamCmd.Flags().StringVarP(&_iamFilter, "filter", "", "", "Filter")
	_iamCmd.Flags().StringVarP(&_iamGlobalEndpointTokenVersion, "global-endpoint-token-version", "", "", "Global Endpoint Token Version")
	_iamCmd.Flags().StringVarP(&_iamGranularity, "granularity", "", "", "Granularity")
	_iamCmd.Flags().StringVarP(&_iamGroupName, "group-name", "", "", "Group Name")
	_iamCmd.Flags().StringVarP(&_iamHardExpiry, "hard-expiry", "", "", "Hard Expiry")
	_iamCmd.Flags().StringVarP(&_iamInstanceProfileName, "instance-profile-name", "", "", "Instance Profile Name")
	_iamCmd.Flags().StringVarP(&_iamJobId, "job-id", "", "", "Job ID")
	_iamCmd.Flags().StringVarP(&_iamLocale, "locale", "", "", "Locale")
	_iamCmd.Flags().StringVarP(&_iamMarker, "marker", "", "", "Marker")
	_iamCmd.Flags().StringVarP(&_iamMaxItems, "max-items", "", "", "Max Items")
	_iamCmd.Flags().StringVarP(&_iamMaxPasswordAge, "max-password-age", "", "", "Max Password Age")
	_iamCmd.Flags().StringVarP(&_iamMaxSessionDuration, "max-session-duration", "", "", "Max Session Duration")
	_iamCmd.Flags().StringVarP(&_iamMinimumPasswordLength, "minimum-password-length", "", "", "Minimum Password Length")
	_iamCmd.Flags().StringVarP(&_iamName, "name", "", "", "Name")
	_iamCmd.Flags().StringVarP(&_iamNewGroupName, "new-group-name", "", "", "New Group Name")
	_iamCmd.Flags().StringVarP(&_iamNewPassword, "new-password", "", "", "New Password")
	_iamCmd.Flags().StringVarP(&_iamNewPath, "new-path", "", "", "New Path")
	_iamCmd.Flags().StringVarP(&_iamNewServerCertificateName, "new-server-certificate-name", "", "", "New Server Certificate Name")
	_iamCmd.Flags().StringVarP(&_iamNewUserName, "new-user-name", "", "", "New User Name")
	_iamCmd.Flags().StringVarP(&_iamNotes, "notes", "", "", "Notes")
	_iamCmd.Flags().StringVarP(&_iamNotificationChannel, "notification-channel", "", "", "Notification Channel")
	_iamCmd.Flags().StringVarP(&_iamOldPassword, "old-password", "", "", "Old Password")
	_iamCmd.Flags().StringVarP(&_iamOnlyAttached, "only-attached", "", "", "Only Attached")
	_iamCmd.Flags().StringVarP(&_iamOnlySendByOwner, "only-send-by-owner", "", "", "Only Send By Owner")
	_iamCmd.Flags().StringVarP(&_iamOpenIDConnectProviderArn, "open-id-connect-provider-arn", "", "", "Open ID Connect Provider ARN")
	_iamCmd.Flags().StringVarP(&_iamOrganizationsPolicyId, "organizations-policy-id", "", "", "Organizations Policy ID")
	_iamCmd.Flags().StringVarP(&_iamOwnerAccountId, "owner-account-id", "", "", "Owner Account ID")
	_iamCmd.Flags().StringVarP(&_iamOwnerId, "owner-id", "", "", "Owner ID")
	_iamCmd.Flags().StringVarP(&_iamPassword, "password", "", "", "Password")
	_iamCmd.Flags().StringVarP(&_iamPasswordResetRequired, "password-reset-required", "", "", "Password Reset Required")
	_iamCmd.Flags().StringVarP(&_iamPasswordReusePrevention, "password-reuse-prevention", "", "", "Password Reuse Prevention")
	_iamCmd.Flags().StringVarP(&_iamPath, "path", "", "", "Path")
	_iamCmd.Flags().StringVarP(&_iamPathPrefix, "path-prefix", "", "", "Path Prefix")
	_iamCmd.Flags().StringVarP(&_iamPermissions, "permissions", "", "", "Permissions")
	_iamCmd.Flags().StringVarP(&_iamPermissionsBoundary, "permissions-boundary", "", "", "Permissions Boundary")
	_iamCmd.Flags().StringSliceVarP(&_iamPermissionsBoundaryPolicyInputList, "permissions-boundary-policy-input-list", "", nil, "Permissions Boundary Policy Input List")
	_iamCmd.Flags().StringVarP(&_iamPolicyArn, "policy-arn", "", "", "Policy ARN")
	_iamCmd.Flags().StringVarP(&_iamPolicyDocument, "policy-document", "", "", "Policy Document")
	_iamCmd.Flags().StringSliceVarP(&_iamPolicyInputList, "policy-input-list", "", nil, "Policy Input List")
	_iamCmd.Flags().StringVarP(&_iamPolicyName, "policy-name", "", "", "Policy Name")
	_iamCmd.Flags().StringVarP(&_iamPolicySourceArn, "policy-source-arn", "", "", "Policy Source ARN")
	_iamCmd.Flags().StringVarP(&_iamPolicyUsageFilter, "policy-usage-filter", "", "", "Policy Usage Filter")
	_iamCmd.Flags().StringVarP(&_iamPrivateKey, "private-key", "", "", "Private Key")
	_iamCmd.Flags().StringVarP(&_iamRedirectUrl, "redirect-url", "", "", "Redirect URL")
	_iamCmd.Flags().StringVarP(&_iamRemovePrivateKey, "remove-private-key", "", "", "Remove Private Key")
	_iamCmd.Flags().StringVarP(&_iamRequestMessage, "request-message", "", "", "Request Message")
	_iamCmd.Flags().StringVarP(&_iamRequestorWorkflowId, "requestor-workflow-id", "", "", "Requestor Workflow ID")
	_iamCmd.Flags().StringVarP(&_iamRequireLowercaseCharacters, "require-lowercase-characters", "", "", "Require Lowercase Characters")
	_iamCmd.Flags().StringVarP(&_iamRequireNumbers, "require-numbers", "", "", "Require Numbers")
	_iamCmd.Flags().StringVarP(&_iamRequireSymbols, "require-symbols", "", "", "Require Symbols")
	_iamCmd.Flags().StringVarP(&_iamRequireUppercaseCharacters, "require-uppercase-characters", "", "", "Require Uppercase Characters")
	_iamCmd.Flags().StringSliceVarP(&_iamResourceArns, "resource-arns", "", nil, "Resource Arns")
	_iamCmd.Flags().StringVarP(&_iamResourceHandlingOption, "resource-handling-option", "", "", "Resource Handling Option")
	_iamCmd.Flags().StringVarP(&_iamResourceOwner, "resource-owner", "", "", "Resource Owner")
	_iamCmd.Flags().StringVarP(&_iamResourcePolicy, "resource-policy", "", "", "Resource Policy")
	_iamCmd.Flags().StringVarP(&_iamRoleName, "role-name", "", "", "Role Name")
	_iamCmd.Flags().StringVarP(&_iamSAMLMetadataDocument, "saml-metadata-document", "", "", "Saml Metadata Document")
	_iamCmd.Flags().StringVarP(&_iamSAMLProviderArn, "saml-provider-arn", "", "", "Saml Provider ARN")
	_iamCmd.Flags().StringVarP(&_iamScope, "scope", "", "", "Scope")
	_iamCmd.Flags().StringVarP(&_iamSerialNumber, "serial-number", "", "", "Serial Number")
	_iamCmd.Flags().StringVarP(&_iamServerCertificateName, "server-certificate-name", "", "", "Server Certificate Name")
	_iamCmd.Flags().StringVarP(&_iamServiceName, "service-name", "", "", "Service Name")
	_iamCmd.Flags().StringVarP(&_iamServiceNamespace, "service-namespace", "", "", "Service Namespace")
	_iamCmd.Flags().StringSliceVarP(&_iamServiceNamespaces, "service-namespaces", "", nil, "Service Namespaces")
	_iamCmd.Flags().StringVarP(&_iamServiceSpecificCredentialId, "service-specific-credential-id", "", "", "Service Specific Credential ID")
	_iamCmd.Flags().StringVarP(&_iamSessionDuration, "session-duration", "", "", "Session Duration")
	_iamCmd.Flags().StringVarP(&_iamSetAsDefault, "set-as-default", "", "", "Set As Default")
	_iamCmd.Flags().StringVarP(&_iamSortKey, "sort-key", "", "", "Sort Key")
	_iamCmd.Flags().StringVarP(&_iamSSHPublicKeyBody, "ssh-public-key-body", "", "", "SSH Public Key Body")
	_iamCmd.Flags().StringVarP(&_iamSSHPublicKeyId, "ssh-public-key-id", "", "", "SSH Public Key ID")
	_iamCmd.Flags().StringVarP(&_iamStatus, "status", "", "", "Status")
	_iamCmd.Flags().StringSliceVarP(&_iamTagKeys, "tag-keys", "", nil, "Tag Keys")
	_iamCmd.Flags().StringVarP(&_iamTags, "tags", "", "", "Tags")
	_iamCmd.Flags().StringSliceVarP(&_iamThumbprintList, "thumbprint-list", "", nil, "Thumbprint List")
	_iamCmd.Flags().StringVarP(&_iamUrl, "url", "", "", "URL")
	_iamCmd.Flags().StringVarP(&_iamUserName, "user-name", "", "", "User Name")
	_iamCmd.Flags().StringVarP(&_iamVersionId, "version-id", "", "", "Version ID")
	_iamCmd.Flags().StringVarP(&_iamVirtualMFADeviceName, "virtual-mfa-device-name", "", "", "Virtual MFA Device Name")

	_iamCmd.Flags().BoolVarP(&_iamAcceptDelegationRequest, "accept-delegation-request", "", false, "Accept Delegation Request")
	_iamCmd.Flags().BoolVarP(&_iamAddClientIDToOpenIDConnectProvider, "add-client-idto-open-id-connect-provider", "", false, "Add Client Idto Open ID Connect Provider")
	_iamCmd.Flags().BoolVarP(&_iamAddRoleToInstanceProfile, "add-role-to-instance-profile", "", false, "Add Role To Instance Profile")
	_iamCmd.Flags().BoolVarP(&_iamAddUserToGroup, "add-user-to-group", "", false, "Add User To Group")
	_iamCmd.Flags().BoolVarP(&_iamAssociateDelegationRequest, "associate-delegation-request", "", false, "Associate Delegation Request")
	_iamCmd.Flags().BoolVarP(&_iamAttachGroupPolicy, "attach-group-policy", "", false, "Attach Group Policy")
	_iamCmd.Flags().BoolVarP(&_iamAttachRolePolicy, "attach-role-policy", "", false, "Attach Role Policy")
	_iamCmd.Flags().BoolVarP(&_iamAttachUserPolicy, "attach-user-policy", "", false, "Attach User Policy")
	_iamCmd.Flags().BoolVarP(&_iamChangePassword, "change-password", "", false, "Change Password")
	_iamCmd.Flags().BoolVarP(&_iamCreateAccessKey, "create-access-key", "", false, "Create Access Key")
	_iamCmd.Flags().BoolVarP(&_iamCreateAccountAlias, "create-account-alias", "", false, "Create Account Alias")
	_iamCmd.Flags().BoolVarP(&_iamCreateDelegationRequest, "create-delegation-request", "", false, "Create Delegation Request")
	_iamCmd.Flags().BoolVarP(&_iamCreateGroup, "create-group", "", false, "Create Group")
	_iamCmd.Flags().BoolVarP(&_iamCreateInstanceProfile, "create-instance-profile", "", false, "Create Instance Profile")
	_iamCmd.Flags().BoolVarP(&_iamCreateLoginProfile, "create-login-profile", "", false, "Create Login Profile")
	_iamCmd.Flags().BoolVarP(&_iamCreateOpenIDConnectProvider, "create-open-id-connect-provider", "", false, "Create Open ID Connect Provider")
	_iamCmd.Flags().BoolVarP(&_iamCreatePolicy, "create-policy", "", false, "Create Policy")
	_iamCmd.Flags().BoolVarP(&_iamCreatePolicyVersion, "create-policy-version", "", false, "Create Policy Version")
	_iamCmd.Flags().BoolVarP(&_iamCreateRole, "create-role", "", false, "Create Role")
	_iamCmd.Flags().BoolVarP(&_iamCreateSAMLProvider, "create-saml-provider", "", false, "Create Saml Provider")
	_iamCmd.Flags().BoolVarP(&_iamCreateServiceLinkedRole, "create-service-linked-role", "", false, "Create Service Linked Role")
	_iamCmd.Flags().BoolVarP(&_iamCreateServiceSpecificCredential, "create-service-specific-credential", "", false, "Create Service Specific Credential")
	_iamCmd.Flags().BoolVarP(&_iamCreateUser, "create-user", "", false, "Create User")
	_iamCmd.Flags().BoolVarP(&_iamCreateVirtualMFADevice, "create-virtual-mfa-device", "", false, "Create Virtual MFA Device")
	_iamCmd.Flags().BoolVarP(&_iamDeactivateMFADevice, "deactivate-mfa-device", "", false, "Deactivate MFA Device")
	_iamCmd.Flags().BoolVarP(&_iamDeleteAccessKey, "delete-access-key", "", false, "Delete Access Key")
	_iamCmd.Flags().BoolVarP(&_iamDeleteAccountAlias, "delete-account-alias", "", false, "Delete Account Alias")
	_iamCmd.Flags().BoolVarP(&_iamDeleteAccountPasswordPolicy, "delete-account-password-policy", "", false, "Delete Account Password Policy")
	_iamCmd.Flags().BoolVarP(&_iamDeleteGroup, "delete-group", "", false, "Delete Group")
	_iamCmd.Flags().BoolVarP(&_iamDeleteGroupPolicy, "delete-group-policy", "", false, "Delete Group Policy")
	_iamCmd.Flags().BoolVarP(&_iamDeleteInstanceProfile, "delete-instance-profile", "", false, "Delete Instance Profile")
	_iamCmd.Flags().BoolVarP(&_iamDeleteLoginProfile, "delete-login-profile", "", false, "Delete Login Profile")
	_iamCmd.Flags().BoolVarP(&_iamDeleteOpenIDConnectProvider, "delete-open-id-connect-provider", "", false, "Delete Open ID Connect Provider")
	_iamCmd.Flags().BoolVarP(&_iamDeletePolicy, "delete-policy", "", false, "Delete Policy")
	_iamCmd.Flags().BoolVarP(&_iamDeletePolicyVersion, "delete-policy-version", "", false, "Delete Policy Version")
	_iamCmd.Flags().BoolVarP(&_iamDeleteRole, "delete-role", "", false, "Delete Role")
	_iamCmd.Flags().BoolVarP(&_iamDeleteRolePermissionsBoundary, "delete-role-permissions-boundary", "", false, "Delete Role Permissions Boundary")
	_iamCmd.Flags().BoolVarP(&_iamDeleteRolePolicy, "delete-role-policy", "", false, "Delete Role Policy")
	_iamCmd.Flags().BoolVarP(&_iamDeleteSAMLProvider, "delete-saml-provider", "", false, "Delete Saml Provider")
	_iamCmd.Flags().BoolVarP(&_iamDeleteServerCertificate, "delete-server-certificate", "", false, "Delete Server Certificate")
	_iamCmd.Flags().BoolVarP(&_iamDeleteServiceLinkedRole, "delete-service-linked-role", "", false, "Delete Service Linked Role")
	_iamCmd.Flags().BoolVarP(&_iamDeleteServiceSpecificCredential, "delete-service-specific-credential", "", false, "Delete Service Specific Credential")
	_iamCmd.Flags().BoolVarP(&_iamDeleteSigningCertificate, "delete-signing-certificate", "", false, "Delete Signing Certificate")
	_iamCmd.Flags().BoolVarP(&_iamDeleteSSHPublicKey, "delete-ssh-public-key", "", false, "Delete SSH Public Key")
	_iamCmd.Flags().BoolVarP(&_iamDeleteUser, "delete-user", "", false, "Delete User")
	_iamCmd.Flags().BoolVarP(&_iamDeleteUserPermissionsBoundary, "delete-user-permissions-boundary", "", false, "Delete User Permissions Boundary")
	_iamCmd.Flags().BoolVarP(&_iamDeleteUserPolicy, "delete-user-policy", "", false, "Delete User Policy")
	_iamCmd.Flags().BoolVarP(&_iamDeleteVirtualMFADevice, "delete-virtual-mfa-device", "", false, "Delete Virtual MFA Device")
	_iamCmd.Flags().BoolVarP(&_iamDetachGroupPolicy, "detach-group-policy", "", false, "Detach Group Policy")
	_iamCmd.Flags().BoolVarP(&_iamDetachRolePolicy, "detach-role-policy", "", false, "Detach Role Policy")
	_iamCmd.Flags().BoolVarP(&_iamDetachUserPolicy, "detach-user-policy", "", false, "Detach User Policy")
	_iamCmd.Flags().BoolVarP(&_iamDisableOrganizationsRootCredentialsManagement, "disable-organizations-root-credentials-management", "", false, "Disable Organizations Root Credentials Management")
	_iamCmd.Flags().BoolVarP(&_iamDisableOrganizationsRootSessions, "disable-organizations-root-sessions", "", false, "Disable Organizations Root Sessions")
	_iamCmd.Flags().BoolVarP(&_iamDisableOutboundWebIdentityFederation, "disable-outbound-web-identity-federation", "", false, "Disable Outbound Web Identity Federation")
	_iamCmd.Flags().BoolVarP(&_iamEnableMFADevice, "enable-mfa-device", "", false, "Enable MFA Device")
	_iamCmd.Flags().BoolVarP(&_iamEnableOrganizationsRootCredentialsManagement, "enable-organizations-root-credentials-management", "", false, "Enable Organizations Root Credentials Management")
	_iamCmd.Flags().BoolVarP(&_iamEnableOrganizationsRootSessions, "enable-organizations-root-sessions", "", false, "Enable Organizations Root Sessions")
	_iamCmd.Flags().BoolVarP(&_iamEnableOutboundWebIdentityFederation, "enable-outbound-web-identity-federation", "", false, "Enable Outbound Web Identity Federation")
	_iamCmd.Flags().BoolVarP(&_iamGenerateCredentialReport, "generate-credential-report", "", false, "Generate Credential Report")
	_iamCmd.Flags().BoolVarP(&_iamGenerateOrganizationsAccessReport, "generate-organizations-access-report", "", false, "Generate Organizations Access Report")
	_iamCmd.Flags().BoolVarP(&_iamGenerateServiceLastAccessedDetails, "generate-service-last-accessed-details", "", false, "Generate Service Last Accessed Details")
	_iamCmd.Flags().BoolVarP(&_iamGetAccessKeyLastUsed, "get-access-key-last-used", "", false, "Get Access Key Last Used")
	_iamCmd.Flags().BoolVarP(&_iamGetAccountAuthorizationDetails, "get-account-authorization-details", "", false, "Get Account Authorization Details")
	_iamCmd.Flags().BoolVarP(&_iamGetAccountPasswordPolicy, "get-account-password-policy", "", false, "Get Account Password Policy")
	_iamCmd.Flags().BoolVarP(&_iamGetAccountSummary, "get-account-summary", "", false, "Get Account Summary")
	_iamCmd.Flags().BoolVarP(&_iamGetContextKeysForCustomPolicy, "get-context-keys-for-custom-policy", "", false, "Get Context Keys For Custom Policy")
	_iamCmd.Flags().BoolVarP(&_iamGetContextKeysForPrincipalPolicy, "get-context-keys-for-principal-policy", "", false, "Get Context Keys For Principal Policy")
	_iamCmd.Flags().BoolVarP(&_iamGetCredentialReport, "get-credential-report", "", false, "Get Credential Report")
	_iamCmd.Flags().BoolVarP(&_iamGetDelegationRequest, "get-delegation-request", "", false, "Get Delegation Request")
	_iamCmd.Flags().BoolVarP(&_iamGetGroup, "get-group", "", false, "Get Group")
	_iamCmd.Flags().BoolVarP(&_iamGetGroupPolicy, "get-group-policy", "", false, "Get Group Policy")
	_iamCmd.Flags().BoolVarP(&_iamGetHumanReadableSummary, "get-human-readable-summary", "", false, "Get Human Readable Summary")
	_iamCmd.Flags().BoolVarP(&_iamGetInstanceProfile, "get-instance-profile", "", false, "Get Instance Profile")
	_iamCmd.Flags().BoolVarP(&_iamGetLoginProfile, "get-login-profile", "", false, "Get Login Profile")
	_iamCmd.Flags().BoolVarP(&_iamGetMFADevice, "get-mfa-device", "", false, "Get MFA Device")
	_iamCmd.Flags().BoolVarP(&_iamGetOpenIDConnectProvider, "get-open-id-connect-provider", "", false, "Get Open ID Connect Provider")
	_iamCmd.Flags().BoolVarP(&_iamGetOrganizationsAccessReport, "get-organizations-access-report", "", false, "Get Organizations Access Report")
	_iamCmd.Flags().BoolVarP(&_iamGetOutboundWebIdentityFederationInfo, "get-outbound-web-identity-federation-info", "", false, "Get Outbound Web Identity Federation Info")
	_iamCmd.Flags().BoolVarP(&_iamGetPolicy, "get-policy", "", false, "Get Policy")
	_iamCmd.Flags().BoolVarP(&_iamGetPolicyVersion, "get-policy-version", "", false, "Get Policy Version")
	_iamCmd.Flags().BoolVarP(&_iamGetRole, "get-role", "", false, "Get Role")
	_iamCmd.Flags().BoolVarP(&_iamGetRolePolicy, "get-role-policy", "", false, "Get Role Policy")
	_iamCmd.Flags().BoolVarP(&_iamGetSAMLProvider, "get-saml-provider", "", false, "Get Saml Provider")
	_iamCmd.Flags().BoolVarP(&_iamGetServerCertificate, "get-server-certificate", "", false, "Get Server Certificate")
	_iamCmd.Flags().BoolVarP(&_iamGetServiceLastAccessedDetails, "get-service-last-accessed-details", "", false, "Get Service Last Accessed Details")
	_iamCmd.Flags().BoolVarP(&_iamGetServiceLastAccessedDetailsWithEntities, "get-service-last-accessed-details-with-entities", "", false, "Get Service Last Accessed Details With Entities")
	_iamCmd.Flags().BoolVarP(&_iamGetServiceLinkedRoleDeletionStatus, "get-service-linked-role-deletion-status", "", false, "Get Service Linked Role Deletion Status")
	_iamCmd.Flags().BoolVarP(&_iamGetSSHPublicKey, "get-ssh-public-key", "", false, "Get SSH Public Key")
	_iamCmd.Flags().BoolVarP(&_iamGetUser, "get-user", "", false, "Get User")
	_iamCmd.Flags().BoolVarP(&_iamGetUserPolicy, "get-user-policy", "", false, "Get User Policy")
	_iamCmd.Flags().BoolVarP(&_iamListAccessKeys, "list-access-keys", "", false, "List Access Keys")
	_iamCmd.Flags().BoolVarP(&_iamListAccountAliases, "list-account-aliases", "", false, "List Account Aliases")
	_iamCmd.Flags().BoolVarP(&_iamListAttachedGroupPolicies, "list-attached-group-policies", "", false, "List Attached Group Policies")
	_iamCmd.Flags().BoolVarP(&_iamListAttachedRolePolicies, "list-attached-role-policies", "", false, "List Attached Role Policies")
	_iamCmd.Flags().BoolVarP(&_iamListAttachedUserPolicies, "list-attached-user-policies", "", false, "List Attached User Policies")
	_iamCmd.Flags().BoolVarP(&_iamListDelegationRequests, "list-delegation-requests", "", false, "List Delegation Requests")
	_iamCmd.Flags().BoolVarP(&_iamListEntitiesForPolicy, "list-entities-for-policy", "", false, "List Entities For Policy")
	_iamCmd.Flags().BoolVarP(&_iamListGroupPolicies, "list-group-policies", "", false, "List Group Policies")
	_iamCmd.Flags().BoolVarP(&_iamListGroups, "list-groups", "", false, "List Groups")
	_iamCmd.Flags().BoolVarP(&_iamListGroupsForUser, "list-groups-for-user", "", false, "List Groups For User")
	_iamCmd.Flags().BoolVarP(&_iamListInstanceProfileTags, "list-instance-profile-tags", "", false, "List Instance Profile Tags")
	_iamCmd.Flags().BoolVarP(&_iamListInstanceProfiles, "list-instance-profiles", "", false, "List Instance Profiles")
	_iamCmd.Flags().BoolVarP(&_iamListInstanceProfilesForRole, "list-instance-profiles-for-role", "", false, "List Instance Profiles For Role")
	_iamCmd.Flags().BoolVarP(&_iamListMFADeviceTags, "list-mfa-device-tags", "", false, "List MFA Device Tags")
	_iamCmd.Flags().BoolVarP(&_iamListMFADevices, "list-mfa-devices", "", false, "List MFA Devices")
	_iamCmd.Flags().BoolVarP(&_iamListOpenIDConnectProviderTags, "list-open-id-connect-provider-tags", "", false, "List Open ID Connect Provider Tags")
	_iamCmd.Flags().BoolVarP(&_iamListOpenIDConnectProviders, "list-open-id-connect-providers", "", false, "List Open ID Connect Providers")
	_iamCmd.Flags().BoolVarP(&_iamListOrganizationsFeatures, "list-organizations-features", "", false, "List Organizations Features")
	_iamCmd.Flags().BoolVarP(&_iamListPolicies, "list-policies", "", false, "List Policies")
	_iamCmd.Flags().BoolVarP(&_iamListPoliciesGrantingServiceAccess, "list-policies-granting-service-access", "", false, "List Policies Granting Service Access")
	_iamCmd.Flags().BoolVarP(&_iamListPolicyTags, "list-policy-tags", "", false, "List Policy Tags")
	_iamCmd.Flags().BoolVarP(&_iamListPolicyVersions, "list-policy-versions", "", false, "List Policy Versions")
	_iamCmd.Flags().BoolVarP(&_iamListRolePolicies, "list-role-policies", "", false, "List Role Policies")
	_iamCmd.Flags().BoolVarP(&_iamListRoleTags, "list-role-tags", "", false, "List Role Tags")
	_iamCmd.Flags().BoolVarP(&_iamListRoles, "list-roles", "", false, "List Roles")
	_iamCmd.Flags().BoolVarP(&_iamListSAMLProviderTags, "list-saml-provider-tags", "", false, "List Saml Provider Tags")
	_iamCmd.Flags().BoolVarP(&_iamListSAMLProviders, "list-saml-providers", "", false, "List Saml Providers")
	_iamCmd.Flags().BoolVarP(&_iamListServerCertificateTags, "list-server-certificate-tags", "", false, "List Server Certificate Tags")
	_iamCmd.Flags().BoolVarP(&_iamListServerCertificates, "list-server-certificates", "", false, "List Server Certificates")
	_iamCmd.Flags().BoolVarP(&_iamListServiceSpecificCredentials, "list-service-specific-credentials", "", false, "List Service Specific Credentials")
	_iamCmd.Flags().BoolVarP(&_iamListSigningCertificates, "list-signing-certificates", "", false, "List Signing Certificates")
	_iamCmd.Flags().BoolVarP(&_iamListSSHPublicKeys, "list-ssh-public-keys", "", false, "List SSH Public Keys")
	_iamCmd.Flags().BoolVarP(&_iamListUserPolicies, "list-user-policies", "", false, "List User Policies")
	_iamCmd.Flags().BoolVarP(&_iamListUserTags, "list-user-tags", "", false, "List User Tags")
	_iamCmd.Flags().BoolVarP(&_iamListUsers, "list-users", "", false, "List Users")
	_iamCmd.Flags().BoolVarP(&_iamListVirtualMFADevices, "list-virtual-mfa-devices", "", false, "List Virtual MFA Devices")
	_iamCmd.Flags().BoolVarP(&_iamPutGroupPolicy, "put-group-policy", "", false, "Put Group Policy")
	_iamCmd.Flags().BoolVarP(&_iamPutRolePermissionsBoundary, "put-role-permissions-boundary", "", false, "Put Role Permissions Boundary")
	_iamCmd.Flags().BoolVarP(&_iamPutRolePolicy, "put-role-policy", "", false, "Put Role Policy")
	_iamCmd.Flags().BoolVarP(&_iamPutUserPermissionsBoundary, "put-user-permissions-boundary", "", false, "Put User Permissions Boundary")
	_iamCmd.Flags().BoolVarP(&_iamPutUserPolicy, "put-user-policy", "", false, "Put User Policy")
	_iamCmd.Flags().BoolVarP(&_iamRejectDelegationRequest, "reject-delegation-request", "", false, "Reject Delegation Request")
	_iamCmd.Flags().BoolVarP(&_iamRemoveClientIDFromOpenIDConnectProvider, "remove-client-id-from-open-id-connect-provider", "", false, "Remove Client ID From Open ID Connect Provider")
	_iamCmd.Flags().BoolVarP(&_iamRemoveRoleFromInstanceProfile, "remove-role-from-instance-profile", "", false, "Remove Role From Instance Profile")
	_iamCmd.Flags().BoolVarP(&_iamRemoveUserFromGroup, "remove-user-from-group", "", false, "Remove User From Group")
	_iamCmd.Flags().BoolVarP(&_iamResetServiceSpecificCredential, "reset-service-specific-credential", "", false, "Reset Service Specific Credential")
	_iamCmd.Flags().BoolVarP(&_iamResyncMFADevice, "resync-mfa-device", "", false, "Resync MFA Device")
	_iamCmd.Flags().BoolVarP(&_iamSendDelegationToken, "send-delegation-token", "", false, "Send Delegation Token")
	_iamCmd.Flags().BoolVarP(&_iamSetDefaultPolicyVersion, "set-default-policy-version", "", false, "Set Default Policy Version")
	_iamCmd.Flags().BoolVarP(&_iamSetSecurityTokenServicePreferences, "set-security-token-service-preferences", "", false, "Set Security Token Service Preferences")
	_iamCmd.Flags().BoolVarP(&_iamSimulateCustomPolicy, "simulate-custom-policy", "", false, "Simulate Custom Policy")
	_iamCmd.Flags().BoolVarP(&_iamSimulatePrincipalPolicy, "simulate-principal-policy", "", false, "Simulate Principal Policy")
	_iamCmd.Flags().BoolVarP(&_iamTagInstanceProfile, "tag-instance-profile", "", false, "Tag Instance Profile")
	_iamCmd.Flags().BoolVarP(&_iamTagMFADevice, "tag-mfa-device", "", false, "Tag MFA Device")
	_iamCmd.Flags().BoolVarP(&_iamTagOpenIDConnectProvider, "tag-open-id-connect-provider", "", false, "Tag Open ID Connect Provider")
	_iamCmd.Flags().BoolVarP(&_iamTagPolicy, "tag-policy", "", false, "Tag Policy")
	_iamCmd.Flags().BoolVarP(&_iamTagRole, "tag-role", "", false, "Tag Role")
	_iamCmd.Flags().BoolVarP(&_iamTagSAMLProvider, "tag-saml-provider", "", false, "Tag Saml Provider")
	_iamCmd.Flags().BoolVarP(&_iamTagServerCertificate, "tag-server-certificate", "", false, "Tag Server Certificate")
	_iamCmd.Flags().BoolVarP(&_iamTagUser, "tag-user", "", false, "Tag User")
	_iamCmd.Flags().BoolVarP(&_iamUntagInstanceProfile, "untag-instance-profile", "", false, "Untag Instance Profile")
	_iamCmd.Flags().BoolVarP(&_iamUntagMFADevice, "untag-mfa-device", "", false, "Untag MFA Device")
	_iamCmd.Flags().BoolVarP(&_iamUntagOpenIDConnectProvider, "untag-open-id-connect-provider", "", false, "Untag Open ID Connect Provider")
	_iamCmd.Flags().BoolVarP(&_iamUntagPolicy, "untag-policy", "", false, "Untag Policy")
	_iamCmd.Flags().BoolVarP(&_iamUntagRole, "untag-role", "", false, "Untag Role")
	_iamCmd.Flags().BoolVarP(&_iamUntagSAMLProvider, "untag-saml-provider", "", false, "Untag Saml Provider")
	_iamCmd.Flags().BoolVarP(&_iamUntagServerCertificate, "untag-server-certificate", "", false, "Untag Server Certificate")
	_iamCmd.Flags().BoolVarP(&_iamUntagUser, "untag-user", "", false, "Untag User")
	_iamCmd.Flags().BoolVarP(&_iamUpdateAccessKey, "update-access-key", "", false, "Update Access Key")
	_iamCmd.Flags().BoolVarP(&_iamUpdateAccountPasswordPolicy, "update-account-password-policy", "", false, "Update Account Password Policy")
	_iamCmd.Flags().BoolVarP(&_iamUpdateAssumeRolePolicy, "update-assume-role-policy", "", false, "Update Assume Role Policy")
	_iamCmd.Flags().BoolVarP(&_iamUpdateDelegationRequest, "update-delegation-request", "", false, "Update Delegation Request")
	_iamCmd.Flags().BoolVarP(&_iamUpdateGroup, "update-group", "", false, "Update Group")
	_iamCmd.Flags().BoolVarP(&_iamUpdateLoginProfile, "update-login-profile", "", false, "Update Login Profile")
	_iamCmd.Flags().BoolVarP(&_iamUpdateOpenIDConnectProviderThumbprint, "update-open-id-connect-provider-thumbprint", "", false, "Update Open ID Connect Provider Thumbprint")
	_iamCmd.Flags().BoolVarP(&_iamUpdateRole, "update-role", "", false, "Update Role")
	_iamCmd.Flags().BoolVarP(&_iamUpdateRoleDescription, "update-role-description", "", false, "Update Role Description")
	_iamCmd.Flags().BoolVarP(&_iamUpdateSAMLProvider, "update-saml-provider", "", false, "Update Saml Provider")
	_iamCmd.Flags().BoolVarP(&_iamUpdateServerCertificate, "update-server-certificate", "", false, "Update Server Certificate")
	_iamCmd.Flags().BoolVarP(&_iamUpdateServiceSpecificCredential, "update-service-specific-credential", "", false, "Update Service Specific Credential")
	_iamCmd.Flags().BoolVarP(&_iamUpdateSigningCertificate, "update-signing-certificate", "", false, "Update Signing Certificate")
	_iamCmd.Flags().BoolVarP(&_iamUpdateSSHPublicKey, "update-ssh-public-key", "", false, "Update SSH Public Key")
	_iamCmd.Flags().BoolVarP(&_iamUpdateUser, "update-user", "", false, "Update User")
	_iamCmd.Flags().BoolVarP(&_iamUploadServerCertificate, "upload-server-certificate", "", false, "Upload Server Certificate")
	_iamCmd.Flags().BoolVarP(&_iamUploadSigningCertificate, "upload-signing-certificate", "", false, "Upload Signing Certificate")
	_iamCmd.Flags().BoolVarP(&_iamUploadSSHPublicKey, "upload-ssh-public-key", "", false, "Upload SSH Public Key")

}
