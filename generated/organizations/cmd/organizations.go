package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/organizations"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// organizationsCmd represents the organizations command
var _organizationsCmd = &cobra.Command{
	Use:   "organizations",
	Short: "AWS organizations CLI",
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
		client := organizations.NewFromConfig(cfg)
		if _organizationsAcceptHandshake {
			organizations_AcceptHandshake(cfg, client)
			return
		}
		if _organizationsAttachPolicy {
			organizations_AttachPolicy(cfg, client)
			return
		}
		if _organizationsCancelHandshake {
			organizations_CancelHandshake(cfg, client)
			return
		}
		if _organizationsCloseAccount {
			organizations_CloseAccount(cfg, client)
			return
		}
		if _organizationsCreateAccount {
			organizations_CreateAccount(cfg, client)
			return
		}
		if _organizationsCreateGovCloudAccount {
			organizations_CreateGovCloudAccount(cfg, client)
			return
		}
		if _organizationsCreateOrganization {
			organizations_CreateOrganization(cfg, client)
			return
		}
		if _organizationsCreateOrganizationalUnit {
			organizations_CreateOrganizationalUnit(cfg, client)
			return
		}
		if _organizationsCreatePolicy {
			organizations_CreatePolicy(cfg, client)
			return
		}
		if _organizationsDeclineHandshake {
			organizations_DeclineHandshake(cfg, client)
			return
		}
		if _organizationsDeleteOrganization {
			organizations_DeleteOrganization(cfg, client)
			return
		}
		if _organizationsDeleteOrganizationalUnit {
			organizations_DeleteOrganizationalUnit(cfg, client)
			return
		}
		if _organizationsDeletePolicy {
			organizations_DeletePolicy(cfg, client)
			return
		}
		if _organizationsDeleteResourcePolicy {
			organizations_DeleteResourcePolicy(cfg, client)
			return
		}
		if _organizationsDeregisterDelegatedAdministrator {
			organizations_DeregisterDelegatedAdministrator(cfg, client)
			return
		}
		if _organizationsDescribeAccount {
			organizations_DescribeAccount(cfg, client)
			return
		}
		if _organizationsDescribeCreateAccountStatus {
			organizations_DescribeCreateAccountStatus(cfg, client)
			return
		}
		if _organizationsDescribeEffectivePolicy {
			organizations_DescribeEffectivePolicy(cfg, client)
			return
		}
		if _organizationsDescribeHandshake {
			organizations_DescribeHandshake(cfg, client)
			return
		}
		if _organizationsDescribeOrganization {
			organizations_DescribeOrganization(cfg, client)
			return
		}
		if _organizationsDescribeOrganizationalUnit {
			organizations_DescribeOrganizationalUnit(cfg, client)
			return
		}
		if _organizationsDescribePolicy {
			organizations_DescribePolicy(cfg, client)
			return
		}
		if _organizationsDescribeResourcePolicy {
			organizations_DescribeResourcePolicy(cfg, client)
			return
		}
		if _organizationsDescribeResponsibilityTransfer {
			organizations_DescribeResponsibilityTransfer(cfg, client)
			return
		}
		if _organizationsDetachPolicy {
			organizations_DetachPolicy(cfg, client)
			return
		}
		if _organizationsDisableAWSServiceAccess {
			organizations_DisableAWSServiceAccess(cfg, client)
			return
		}
		if _organizationsDisablePolicyType {
			organizations_DisablePolicyType(cfg, client)
			return
		}
		if _organizationsEnableAllFeatures {
			organizations_EnableAllFeatures(cfg, client)
			return
		}
		if _organizationsEnableAWSServiceAccess {
			organizations_EnableAWSServiceAccess(cfg, client)
			return
		}
		if _organizationsEnablePolicyType {
			organizations_EnablePolicyType(cfg, client)
			return
		}
		if _organizationsInviteAccountToOrganization {
			organizations_InviteAccountToOrganization(cfg, client)
			return
		}
		if _organizationsInviteOrganizationToTransferResponsibility {
			organizations_InviteOrganizationToTransferResponsibility(cfg, client)
			return
		}
		if _organizationsLeaveOrganization {
			organizations_LeaveOrganization(cfg, client)
			return
		}
		if _organizationsListAccounts {
			organizations_ListAccounts(cfg, client)
			return
		}
		if _organizationsListAccountsForParent {
			organizations_ListAccountsForParent(cfg, client)
			return
		}
		if _organizationsListAccountsWithInvalidEffectivePolicy {
			organizations_ListAccountsWithInvalidEffectivePolicy(cfg, client)
			return
		}
		if _organizationsListAWSServiceAccessForOrganization {
			organizations_ListAWSServiceAccessForOrganization(cfg, client)
			return
		}
		if _organizationsListChildren {
			organizations_ListChildren(cfg, client)
			return
		}
		if _organizationsListCreateAccountStatus {
			organizations_ListCreateAccountStatus(cfg, client)
			return
		}
		if _organizationsListDelegatedAdministrators {
			organizations_ListDelegatedAdministrators(cfg, client)
			return
		}
		if _organizationsListDelegatedServicesForAccount {
			organizations_ListDelegatedServicesForAccount(cfg, client)
			return
		}
		if _organizationsListEffectivePolicyValidationErrors {
			organizations_ListEffectivePolicyValidationErrors(cfg, client)
			return
		}
		if _organizationsListHandshakesForAccount {
			organizations_ListHandshakesForAccount(cfg, client)
			return
		}
		if _organizationsListHandshakesForOrganization {
			organizations_ListHandshakesForOrganization(cfg, client)
			return
		}
		if _organizationsListInboundResponsibilityTransfers {
			organizations_ListInboundResponsibilityTransfers(cfg, client)
			return
		}
		if _organizationsListOrganizationalUnitsForParent {
			organizations_ListOrganizationalUnitsForParent(cfg, client)
			return
		}
		if _organizationsListOutboundResponsibilityTransfers {
			organizations_ListOutboundResponsibilityTransfers(cfg, client)
			return
		}
		if _organizationsListParents {
			organizations_ListParents(cfg, client)
			return
		}
		if _organizationsListPolicies {
			organizations_ListPolicies(cfg, client)
			return
		}
		if _organizationsListPoliciesForTarget {
			organizations_ListPoliciesForTarget(cfg, client)
			return
		}
		if _organizationsListRoots {
			organizations_ListRoots(cfg, client)
			return
		}
		if _organizationsListTagsForResource {
			organizations_ListTagsForResource(cfg, client)
			return
		}
		if _organizationsListTargetsForPolicy {
			organizations_ListTargetsForPolicy(cfg, client)
			return
		}
		if _organizationsMoveAccount {
			organizations_MoveAccount(cfg, client)
			return
		}
		if _organizationsPutResourcePolicy {
			organizations_PutResourcePolicy(cfg, client)
			return
		}
		if _organizationsRegisterDelegatedAdministrator {
			organizations_RegisterDelegatedAdministrator(cfg, client)
			return
		}
		if _organizationsRemoveAccountFromOrganization {
			organizations_RemoveAccountFromOrganization(cfg, client)
			return
		}
		if _organizationsTagResource {
			organizations_TagResource(cfg, client)
			return
		}
		if _organizationsTerminateResponsibilityTransfer {
			organizations_TerminateResponsibilityTransfer(cfg, client)
			return
		}
		if _organizationsUntagResource {
			organizations_UntagResource(cfg, client)
			return
		}
		if _organizationsUpdateOrganizationalUnit {
			organizations_UpdateOrganizationalUnit(cfg, client)
			return
		}
		if _organizationsUpdatePolicy {
			organizations_UpdatePolicy(cfg, client)
			return
		}
		if _organizationsUpdateResponsibilityTransfer {
			organizations_UpdateResponsibilityTransfer(cfg, client)
			return
		}

	},
}

var (
	_organizationsAcceptHandshake                            bool
	_organizationsAttachPolicy                               bool
	_organizationsCancelHandshake                            bool
	_organizationsCloseAccount                               bool
	_organizationsCreateAccount                              bool
	_organizationsCreateGovCloudAccount                      bool
	_organizationsCreateOrganization                         bool
	_organizationsCreateOrganizationalUnit                   bool
	_organizationsCreatePolicy                               bool
	_organizationsDeclineHandshake                           bool
	_organizationsDeleteOrganization                         bool
	_organizationsDeleteOrganizationalUnit                   bool
	_organizationsDeletePolicy                               bool
	_organizationsDeleteResourcePolicy                       bool
	_organizationsDeregisterDelegatedAdministrator           bool
	_organizationsDescribeAccount                            bool
	_organizationsDescribeCreateAccountStatus                bool
	_organizationsDescribeEffectivePolicy                    bool
	_organizationsDescribeHandshake                          bool
	_organizationsDescribeOrganization                       bool
	_organizationsDescribeOrganizationalUnit                 bool
	_organizationsDescribePolicy                             bool
	_organizationsDescribeResourcePolicy                     bool
	_organizationsDescribeResponsibilityTransfer             bool
	_organizationsDetachPolicy                               bool
	_organizationsDisableAWSServiceAccess                    bool
	_organizationsDisablePolicyType                          bool
	_organizationsEnableAllFeatures                          bool
	_organizationsEnableAWSServiceAccess                     bool
	_organizationsEnablePolicyType                           bool
	_organizationsInviteAccountToOrganization                bool
	_organizationsInviteOrganizationToTransferResponsibility bool
	_organizationsLeaveOrganization                          bool
	_organizationsListAccounts                               bool
	_organizationsListAccountsForParent                      bool
	_organizationsListAccountsWithInvalidEffectivePolicy     bool
	_organizationsListAWSServiceAccessForOrganization        bool
	_organizationsListChildren                               bool
	_organizationsListCreateAccountStatus                    bool
	_organizationsListDelegatedAdministrators                bool
	_organizationsListDelegatedServicesForAccount            bool
	_organizationsListEffectivePolicyValidationErrors        bool
	_organizationsListHandshakesForAccount                   bool
	_organizationsListHandshakesForOrganization              bool
	_organizationsListInboundResponsibilityTransfers         bool
	_organizationsListOrganizationalUnitsForParent           bool
	_organizationsListOutboundResponsibilityTransfers        bool
	_organizationsListParents                                bool
	_organizationsListPolicies                               bool
	_organizationsListPoliciesForTarget                      bool
	_organizationsListRoots                                  bool
	_organizationsListTagsForResource                        bool
	_organizationsListTargetsForPolicy                       bool
	_organizationsMoveAccount                                bool
	_organizationsPutResourcePolicy                          bool
	_organizationsRegisterDelegatedAdministrator             bool
	_organizationsRemoveAccountFromOrganization              bool
	_organizationsTagResource                                bool
	_organizationsTerminateResponsibilityTransfer            bool
	_organizationsUntagResource                              bool
	_organizationsUpdateOrganizationalUnit                   bool
	_organizationsUpdatePolicy                               bool
	_organizationsUpdateResponsibilityTransfer               bool

	_organizationsAccountId              string
	_organizationsAccountName            string
	_organizationsChildId                string
	_organizationsChildType              string
	_organizationsContent                string
	_organizationsCreateAccountRequestId string
	_organizationsDescription            string
	_organizationsDestinationParentId    string
	_organizationsEmail                  string
	_organizationsEndTimestamp           string
	_organizationsFeatureSet             string
	_organizationsFilter                 string
	_organizationsHandshakeId            string
	_organizationsIamUserAccessToBilling string
	_organizationsId                     string
	_organizationsMaxResults             string
	_organizationsName                   string
	_organizationsNextToken              string
	_organizationsNotes                  string
	_organizationsOrganizationalUnitId   string
	_organizationsParentId               string
	_organizationsPolicyId               string
	_organizationsPolicyType             string
	_organizationsResourceId             string
	_organizationsRoleName               string
	_organizationsRootId                 string
	_organizationsServicePrincipal       string
	_organizationsSourceName             string
	_organizationsSourceParentId         string
	_organizationsStartTimestamp         string
	_organizationsStates                 string
	_organizationsTagKeys                []string
	_organizationsTags                   string
	_organizationsTarget                 string
	_organizationsTargetId               string
	_organizationsType                   string
)

// Accepts a handshake by sending an ACCEPTED response to the sender. You can view
// accepted handshakes in API responses for 30 days before they are deleted.
//
// Only the management account can accept the following handshakes:
//
// - Enable all features final confirmation ( APPROVE_ALL_FEATURES )
//
// - Billing transfer ( TRANSFER_RESPONSIBILITY )
//
// For more information, see [Enabling all features] and [Responding to a billing transfer invitation] in the Organizations User Guide.
//
// Only a member account can accept the following handshakes:
//
// - Invitation to join ( INVITE )
//
// - Approve all features request ( ENABLE_ALL_FEATURES )
//
// For more information, see [Responding to invitations] and [Enabling all features] in the Organizations User Guide.
//
// [Enabling all features]: https://docs.aws.amazon.com/organizations/latest/userguide/manage-begin-all-features-standard-migration.html#manage-approve-all-features-invite
// [Responding to invitations]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_accounts_accept-decline-invite.html
// [Responding to a billing transfer invitation]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_transfer_billing-respond-invitation.html
func organizations_AcceptHandshake(cfg aws.Config, client *organizations.Client) {
	input := &organizations.AcceptHandshakeInput{
		// HandshakeId: *string, // Required
	}

	if len(_organizationsHandshakeId) > 0 {
		input.HandshakeId = aws.String(_organizationsHandshakeId)
	}

	if resp, err := client.AcceptHandshake(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Attaches a policy to a root, an organizational unit (OU), or an individual
// account. How the policy affects accounts depends on the type of policy. Refer to
// the Organizations User Guide for information about each policy type:
//
// [SERVICE_CONTROL_POLICY]
//
// [RESOURCE_CONTROL_POLICY]
//
// [DECLARATIVE_POLICY_EC2]
//
// [BACKUP_POLICY]
//
// [TAG_POLICY]
//
// [CHATBOT_POLICY]
//
// [AISERVICES_OPT_OUT_POLICY]
//
// [SECURITYHUB_POLICY]
//
// [UPGRADE_ROLLOUT_POLICY]
//
// [INSPECTOR_POLICY]
//
// [BEDROCK_POLICY]
//
// [S3_POLICY]
//
// [NETWORK_SECURITY_DIRECTOR_POLICY]
//
// You can only call this operation from the management account or a member
// account that is a delegated administrator.
//
// [BEDROCK_POLICY]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_bedrock.html
// [NETWORK_SECURITY_DIRECTOR_POLICY]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_network_security_director.html
// [UPGRADE_ROLLOUT_POLICY]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_upgrade_rollout.html
// [BACKUP_POLICY]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_backup.html
// [CHATBOT_POLICY]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_chatbot.html
// [TAG_POLICY]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_tag-policies.html
// [INSPECTOR_POLICY]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_inspector.html
// [S3_POLICY]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_s3.html
// [RESOURCE_CONTROL_POLICY]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_rcps.html
// [AISERVICES_OPT_OUT_POLICY]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_ai-opt-out.html
// [SECURITYHUB_POLICY]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_security_hub.html
// [SERVICE_CONTROL_POLICY]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_scp.html
// [DECLARATIVE_POLICY_EC2]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_declarative.html
func organizations_AttachPolicy(cfg aws.Config, client *organizations.Client) {
	input := &organizations.AttachPolicyInput{
		// PolicyId: *string, // Required
		// TargetId: *string, // Required
	}

	if len(_organizationsPolicyId) > 0 {
		input.PolicyId = aws.String(_organizationsPolicyId)
	}
	if len(_organizationsTargetId) > 0 {
		input.TargetId = aws.String(_organizationsTargetId)
	}

	if resp, err := client.AttachPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels a Handshake.
// Only the account that sent a handshake can call this operation. The recipient
// of the handshake can't cancel it, but can use DeclineHandshaketo decline. After a handshake is
// canceled, the recipient can no longer respond to the handshake.
//
// You can view canceled handshakes in API responses for 30 days before they are
// deleted.
func organizations_CancelHandshake(cfg aws.Config, client *organizations.Client) {
	input := &organizations.CancelHandshakeInput{
		// HandshakeId: *string, // Required
	}

	if len(_organizationsHandshakeId) > 0 {
		input.HandshakeId = aws.String(_organizationsHandshakeId)
	}

	if resp, err := client.CancelHandshake(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Closes an Amazon Web Services member account within an organization. You can
// close an account when [all features are enabled]. You can't close the management account with this API.
// This is an asynchronous request that Amazon Web Services performs in the
// background. Because CloseAccount operates asynchronously, it can return a
// successful completion message even though account closure might still be in
// progress. You need to wait a few minutes before the account is fully closed. To
// check the status of the request, do one of the following:
//
// - Use the AccountId that you sent in the CloseAccount request to provide as a
// parameter to the DescribeAccountoperation.
//
// # While the close account request is in progress, Account status will indicate
//
// PENDING_CLOSURE. When the close account request completes, the status will
// change to SUSPENDED.
//
// - Check the CloudTrail log for the CloseAccountResult event that gets
// published after the account closes successfully. For information on using
// CloudTrail with Organizations, see [Logging and monitoring in Organizations]in the Organizations User Guide.
//
// - Resources remaining within the account after closing will be automatically
// deleted after 90 days. During this 90-day period, the resources won't be
// available unless you contact Amazon Web Services Support to reopen the account.
// After 90 days, you can't reopen an account. You might still receive a [bill after account closure].
//
// - You can close only 10% of member accounts, between 10 and 1000, within a
// rolling 30 day period. This quota is not bound by a calendar month, but starts
// when you close an account. After you reach this limit, you can't close
// additional accounts. For more information, see [Closing a member account in your organization]and [Quotas for Organizations]in the Organizations User
// Guide.
//
// - To reinstate a closed account, contact Amazon Web Services Support within
// the 90-day grace period while the account is in SUSPENDED status.
//
// - If the Amazon Web Services account you attempt to close is linked to an
// Amazon Web Services GovCloud (US) account, the CloseAccount request will close
// both accounts. To learn important pre-closure details, see [Closing an Amazon Web Services GovCloud (US) account]in the Amazon Web
// Services GovCloud User Guide.
//
// [all features are enabled]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_org_support-all-features.html
// [Quotas for Organizations]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_reference_limits.html
// [Logging and monitoring in Organizations]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_security_incident-response.html#orgs_cloudtrail-integration
// [bill after account closure]: https://repost.aws/knowledge-center/closed-account-bill
// [Closing an Amazon Web Services GovCloud (US) account]: https://docs.aws.amazon.com/govcloud-us/latest/UserGuide/Closing-govcloud-account.html
// [Closing a member account in your organization]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_accounts_close.html
func organizations_CloseAccount(cfg aws.Config, client *organizations.Client) {
	input := &organizations.CloseAccountInput{
		// AccountId: *string, // Required
	}

	if len(_organizationsAccountId) > 0 {
		input.AccountId = aws.String(_organizationsAccountId)
	}

	if resp, err := client.CloseAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon Web Services account that is automatically a member of the
// organization whose credentials made the request. This is an asynchronous request
// that Amazon Web Services performs in the background. Because CreateAccount
// operates asynchronously, it can return a successful completion message even
// though account initialization might still be in progress. You might need to wait
// a few minutes before you can successfully access the account. To check the
// status of the request, do one of the following:
//
// - Use the Id value of the CreateAccountStatus response element from this
// operation to provide as a parameter to the DescribeCreateAccountStatusoperation.
//
// - Check the CloudTrail log for the CreateAccountResult event. For information
// on using CloudTrail with Organizations, see [Logging and monitoring in Organizations]in the Organizations User Guide.
//
// The user who calls the API to create an account must have the
// organizations:CreateAccount permission. If you enabled all features in the
// organization, Organizations creates the required service-linked role named
// AWSServiceRoleForOrganizations . For more information, see [Organizations and service-linked roles] in the
// Organizations User Guide.
//
// If the request includes tags, then the requester must have the
// organizations:TagResource permission.
//
// Organizations preconfigures the new member account with a role (named
// OrganizationAccountAccessRole by default) that grants users in the management
// account administrator permissions in the new member account. Principals in the
// management account can assume the role. Organizations clones the company name
// and address information for the new account from the organization's management
// account.
//
// You can only call this operation from the management account.
//
// For more information about creating accounts, see [Creating a member account in your organization] in the Organizations User
// Guide.
//
// - When you create an account in an organization using the Organizations
// console, API, or CLI commands, the information required for the account to
// operate as a standalone account, such as a payment method is not automatically
// collected. If you must remove an account from your organization later, you can
// do so only after you provide the missing information. For more information, see [Considerations before removing an account from an organization]
// in the Organizations User Guide.
//
// - If you get an exception that indicates that you exceeded your account
// limits for the organization, contact [Amazon Web Services Support].
//
// - If you get an exception that indicates that the operation failed because
// your organization is still initializing, wait one hour and then try again. If
// the error persists, contact [Amazon Web Services Support].
//
// - It isn't recommended to use CreateAccount to create multiple temporary
// accounts, and using the CreateAccount API to close accounts is subject to a
// 30-day usage quota. For information on the requirements and process for closing
// an account, see [Closing a member account in your organization]in the Organizations User Guide.
//
// When you create a member account with this operation, you can choose whether to
// create the account with the IAM User and Role Access to Billing Information
// switch enabled. If you enable it, IAM users and roles that have appropriate
// permissions can view billing information for the account. If you disable it,
// only the account root user can access billing information. For information about
// how to disable this switch for an account, see [Granting access to your billing information and tools].
//
// [Granting access to your billing information and tools]: https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/control-access-billing.html#grantaccess
// [Amazon Web Services Support]: https://console.aws.amazon.com/support/home#/
// [Logging and monitoring in Organizations]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_security_incident-response.html#orgs_cloudtrail-integration
// [Organizations and service-linked roles]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_integrate_services.html#orgs_integrate_services-using_slrs
// [Creating a member account in your organization]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_accounts_create.html
// [Considerations before removing an account from an organization]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_account-before-remove.html
// [Closing a member account in your organization]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_accounts_close.html
func organizations_CreateAccount(cfg aws.Config, client *organizations.Client) {
	input := &organizations.CreateAccountInput{
		// AccountName: *string, // Required
		// Email: *string, // Required
	}

	if len(_organizationsAccountName) > 0 {
		input.AccountName = aws.String(_organizationsAccountName)
	}
	if len(_organizationsEmail) > 0 {
		input.Email = aws.String(_organizationsEmail)
	}
	if len(_organizationsIamUserAccessToBilling) > 0 {
		if err := assignInputField(input, "IamUserAccessToBilling", _organizationsIamUserAccessToBilling); err != nil {
			log.Errorf("invalid --iam-user-access-to-billing: %s", err.Error())
			return
		}
	}
	if len(_organizationsRoleName) > 0 {
		input.RoleName = aws.String(_organizationsRoleName)
	}
	if len(_organizationsTags) > 0 {
		if err := assignInputField(input, "Tags", _organizationsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This action is available if all of the following are true:
// - You're authorized to create accounts in the Amazon Web Services GovCloud
// (US) Region. For more information on the Amazon Web Services GovCloud (US)
// Region, see the [Amazon Web Services GovCloud User Guide.]
//
// - You already have an account in the Amazon Web Services GovCloud (US) Region
// that is paired with a management account of an organization in the commercial
// Region.
//
// - You call this action from the management account of your organization in
// the commercial Region.
//
// - You have the organizations:CreateGovCloudAccount permission.
//
// Organizations automatically creates the required service-linked role named
// AWSServiceRoleForOrganizations . For more information, see [Organizations and service-linked roles] in the
// Organizations User Guide.
//
// Amazon Web Services automatically enables CloudTrail for Amazon Web Services
// GovCloud (US) accounts, but you should also do the following:
//
// - Verify that CloudTrail is enabled to store logs.
//
// - Create an Amazon S3 bucket for CloudTrail log storage.
//
// For more information, see [Verifying CloudTrail Is Enabled]in the Amazon Web Services GovCloud User Guide.
//
// If the request includes tags, then the requester must have the
// organizations:TagResource permission. The tags are attached to the commercial
// account associated with the GovCloud account, rather than the GovCloud account
// itself. To add tags to the GovCloud account, call the TagResourceoperation in the GovCloud
// Region after the new GovCloud account exists.
//
// You call this action from the management account of your organization in the
// commercial Region to create a standalone Amazon Web Services account in the
// Amazon Web Services GovCloud (US) Region. After the account is created, the
// management account of an organization in the Amazon Web Services GovCloud (US)
// Region can invite it to that organization. For more information on inviting
// standalone accounts in the Amazon Web Services GovCloud (US) to join an
// organization, see [Organizations]in the Amazon Web Services GovCloud User Guide.
//
// Calling CreateGovCloudAccount is an asynchronous request that Amazon Web
// Services performs in the background. Because CreateGovCloudAccount operates
// asynchronously, it can return a successful completion message even though
// account initialization might still be in progress. You might need to wait a few
// minutes before you can successfully access the account. To check the status of
// the request, do one of the following:
//
// - Use the OperationId response element from this operation to provide as a
// parameter to the DescribeCreateAccountStatusoperation.
//
// - Check the CloudTrail log for the CreateAccountResult event. For information
// on using CloudTrail with Organizations, see [Logging and monitoring in Organizations]in the Organizations User Guide.
//
// When you call the CreateGovCloudAccount action, you create two accounts: a
// standalone account in the Amazon Web Services GovCloud (US) Region and an
// associated account in the commercial Region for billing and support purposes.
// The account in the commercial Region is automatically a member of the
// organization whose credentials made the request. Both accounts are associated
// with the same email address.
//
// A role is created in the new account in the commercial Region that allows the
// management account in the organization in the commercial Region to assume it. An
// Amazon Web Services GovCloud (US) account is then created and associated with
// the commercial account that you just created. A role is also created in the new
// Amazon Web Services GovCloud (US) account that can be assumed by the Amazon Web
// Services GovCloud (US) account that is associated with the management account of
// the commercial organization. For more information and to view a diagram that
// explains how account access works, see [Organizations]in the Amazon Web Services GovCloud User
// Guide.
//
// For more information about creating accounts, see [Creating a member account in your organization] in the Organizations User
// Guide.
//
// - When you create an account in an organization using the Organizations
// console, API, or CLI commands, the information required for the account to
// operate as a standalone account is not automatically collected. This includes a
// payment method and signing the end user license agreement (EULA). If you must
// remove an account from your organization later, you can do so only after you
// provide the missing information. For more information, see [Considerations before removing an account from an organization]in the
// Organizations User Guide.
//
// - If you get an exception that indicates that you exceeded your account
// limits for the organization, contact [Amazon Web Services Support].
//
// - If you get an exception that indicates that the operation failed because
// your organization is still initializing, wait one hour and then try again. If
// the error persists, contact [Amazon Web Services Support].
//
// - Using CreateGovCloudAccount to create multiple temporary accounts isn't
// recommended. You can only close an account from the Amazon Web Services Billing
// and Cost Management console, and you must be signed in as the root user. For
// information on the requirements and process for closing an account, see [Closing a member account in your organization]in
// the Organizations User Guide.
//
// When you create a member account with this operation, you can choose whether to
// create the account with the IAM User and Role Access to Billing Information
// switch enabled. If you enable it, IAM users and roles that have appropriate
// permissions can view billing information for the account. If you disable it,
// only the account root user can access billing information. For information about
// how to disable this switch for an account, see [Granting access to your billing information and tools].
//
// [Granting access to your billing information and tools]: https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/grantaccess.html
// [Verifying CloudTrail Is Enabled]: https://docs.aws.amazon.com/govcloud-us/latest/UserGuide/verifying-cloudtrail.html
// [Amazon Web Services Support]: https://console.aws.amazon.com/support/home#/
// [Logging and monitoring in Organizations]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_security_incident-response.html
// [Organizations]: https://docs.aws.amazon.com/govcloud-us/latest/UserGuide/govcloud-organizations.html
// [Organizations and service-linked roles]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_integrate_services.html#orgs_integrate_services-using_slrs
// [Creating a member account in your organization]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_accounts_create.html
// [Considerations before removing an account from an organization]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_account-before-remove.html
// [Amazon Web Services GovCloud User Guide.]: https://docs.aws.amazon.com/govcloud-us/latest/UserGuide/welcome.html
// [Closing a member account in your organization]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_accounts_close.html
func organizations_CreateGovCloudAccount(cfg aws.Config, client *organizations.Client) {
	input := &organizations.CreateGovCloudAccountInput{
		// AccountName: *string, // Required
		// Email: *string, // Required
	}

	if len(_organizationsAccountName) > 0 {
		input.AccountName = aws.String(_organizationsAccountName)
	}
	if len(_organizationsEmail) > 0 {
		input.Email = aws.String(_organizationsEmail)
	}
	if len(_organizationsIamUserAccessToBilling) > 0 {
		if err := assignInputField(input, "IamUserAccessToBilling", _organizationsIamUserAccessToBilling); err != nil {
			log.Errorf("invalid --iam-user-access-to-billing: %s", err.Error())
			return
		}
	}
	if len(_organizationsRoleName) > 0 {
		input.RoleName = aws.String(_organizationsRoleName)
	}
	if len(_organizationsTags) > 0 {
		if err := assignInputField(input, "Tags", _organizationsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateGovCloudAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon Web Services organization. The account whose user is calling
// the CreateOrganization operation automatically becomes the [management account] of the new
// organization.
//
// This operation must be called using credentials from the account that is to
// become the new organization's management account. The principal must also have
// the relevant IAM permissions.
//
// By default (or if you set the FeatureSet parameter to ALL ), the new
// organization is created with all features enabled and service control policies
// automatically enabled in the root. If you instead choose to create the
// organization supporting only the consolidated billing features by setting the
// FeatureSet parameter to CONSOLIDATED_BILLING , no policy types are enabled by
// default and you can't use organization policies.
//
// [management account]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_getting-started_concepts.html#account
func organizations_CreateOrganization(cfg aws.Config, client *organizations.Client) {
	input := &organizations.CreateOrganizationInput{}

	if len(_organizationsFeatureSet) > 0 {
		if err := assignInputField(input, "FeatureSet", _organizationsFeatureSet); err != nil {
			log.Errorf("invalid --feature-set: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateOrganization(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an organizational unit (OU) within a root or parent OU. An OU is a
// container for accounts that enables you to organize your accounts to apply
// policies according to your business requirements. The number of levels deep that
// you can nest OUs is dependent upon the policy types enabled for that root. For
// service control policies, the limit is five.
//
// For more information about OUs, see [Managing organizational units (OUs)] in the Organizations User Guide.
//
// If the request includes tags, then the requester must have the
// organizations:TagResource permission.
//
// You can only call this operation from the management account.
//
// [Managing organizational units (OUs)]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_ous.html
func organizations_CreateOrganizationalUnit(cfg aws.Config, client *organizations.Client) {
	input := &organizations.CreateOrganizationalUnitInput{
		// Name: *string, // Required
		// ParentId: *string, // Required
	}

	if len(_organizationsName) > 0 {
		input.Name = aws.String(_organizationsName)
	}
	if len(_organizationsParentId) > 0 {
		input.ParentId = aws.String(_organizationsParentId)
	}
	if len(_organizationsTags) > 0 {
		if err := assignInputField(input, "Tags", _organizationsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateOrganizationalUnit(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a policy of a specified type that you can attach to a root, an
// organizational unit (OU), or an individual Amazon Web Services account.
//
// For more information about policies and their use, see [Managing Organizations policies].
//
// If the request includes tags, then the requester must have the
// organizations:TagResource permission.
//
// You can only call this operation from the management account or a member
// account that is a delegated administrator.
//
// [Managing Organizations policies]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies.html
func organizations_CreatePolicy(cfg aws.Config, client *organizations.Client) {
	input := &organizations.CreatePolicyInput{
		// Content: *string, // Required
		// Description: *string, // Required
		// Name: *string, // Required
		// Type: types.PolicyType, // Required
	}

	if len(_organizationsContent) > 0 {
		input.Content = aws.String(_organizationsContent)
	}
	if len(_organizationsDescription) > 0 {
		input.Description = aws.String(_organizationsDescription)
	}
	if len(_organizationsName) > 0 {
		input.Name = aws.String(_organizationsName)
	}
	if len(_organizationsType) > 0 {
		if err := assignInputField(input, "Type", _organizationsType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_organizationsTags) > 0 {
		if err := assignInputField(input, "Tags", _organizationsTags); err != nil {
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

// Declines a Handshake.
// Only the account that receives a handshake can call this operation. The sender
// of the handshake can use CancelHandshaketo cancel if the handshake hasn't yet been responded
// to.
//
// You can view canceled handshakes in API responses for 30 days before they are
// deleted.
func organizations_DeclineHandshake(cfg aws.Config, client *organizations.Client) {
	input := &organizations.DeclineHandshakeInput{
		// HandshakeId: *string, // Required
	}

	if len(_organizationsHandshakeId) > 0 {
		input.HandshakeId = aws.String(_organizationsHandshakeId)
	}

	if resp, err := client.DeclineHandshake(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the organization. You can delete an organization only by using
// credentials from the management account. The organization must be empty of
// member accounts.
func organizations_DeleteOrganization(cfg aws.Config, client *organizations.Client) {
	input := &organizations.DeleteOrganizationInput{}

	if resp, err := client.DeleteOrganization(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an organizational unit (OU) from a root or another OU. You must first
// remove all accounts and child OUs from the OU that you want to delete.
//
// You can only call this operation from the management account.
func organizations_DeleteOrganizationalUnit(cfg aws.Config, client *organizations.Client) {
	input := &organizations.DeleteOrganizationalUnitInput{
		// OrganizationalUnitId: *string, // Required
	}

	if len(_organizationsOrganizationalUnitId) > 0 {
		input.OrganizationalUnitId = aws.String(_organizationsOrganizationalUnitId)
	}

	if resp, err := client.DeleteOrganizationalUnit(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified policy from your organization. Before you perform this
// operation, you must first detach the policy from all organizational units (OUs),
// roots, and accounts.
//
// You can only call this operation from the management account or a member
// account that is a delegated administrator.
func organizations_DeletePolicy(cfg aws.Config, client *organizations.Client) {
	input := &organizations.DeletePolicyInput{
		// PolicyId: *string, // Required
	}

	if len(_organizationsPolicyId) > 0 {
		input.PolicyId = aws.String(_organizationsPolicyId)
	}

	if resp, err := client.DeletePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the resource policy from your organization.
// You can only call this operation from the management account.
func organizations_DeleteResourcePolicy(cfg aws.Config, client *organizations.Client) {
	input := &organizations.DeleteResourcePolicyInput{}

	if resp, err := client.DeleteResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the specified member Amazon Web Services account as a delegated
// administrator for the specified Amazon Web Services service.
//
// Deregistering a delegated administrator can have unintended impacts on the
// functionality of the enabled Amazon Web Services service. See the documentation
// for the enabled service before you deregister a delegated administrator so that
// you understand any potential impacts.
//
// You can run this action only for Amazon Web Services services that support this
// feature. For a current list of services that support it, see the column Supports
// Delegated Administrator in the table at [Amazon Web Services Services that you can use with Organizations]in the Organizations User Guide.
//
// You can only call this operation from the management account.
//
// [Amazon Web Services Services that you can use with Organizations]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_integrate_services_list.html
func organizations_DeregisterDelegatedAdministrator(cfg aws.Config, client *organizations.Client) {
	input := &organizations.DeregisterDelegatedAdministratorInput{
		// AccountId: *string, // Required
		// ServicePrincipal: *string, // Required
	}

	if len(_organizationsAccountId) > 0 {
		input.AccountId = aws.String(_organizationsAccountId)
	}
	if len(_organizationsServicePrincipal) > 0 {
		input.ServicePrincipal = aws.String(_organizationsServicePrincipal)
	}

	if resp, err := client.DeregisterDelegatedAdministrator(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves Organizations-related information about the specified account.
// You can only call this operation from the management account or a member
// account that is a delegated administrator.
func organizations_DescribeAccount(cfg aws.Config, client *organizations.Client) {
	input := &organizations.DescribeAccountInput{
		// AccountId: *string, // Required
	}

	if len(_organizationsAccountId) > 0 {
		input.AccountId = aws.String(_organizationsAccountId)
	}

	if resp, err := client.DescribeAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the current status of an asynchronous request to create an account.
// You can only call this operation from the management account or a member
// account that is a delegated administrator.
func organizations_DescribeCreateAccountStatus(cfg aws.Config, client *organizations.Client) {
	input := &organizations.DescribeCreateAccountStatusInput{
		// CreateAccountRequestId: *string, // Required
	}

	if len(_organizationsCreateAccountRequestId) > 0 {
		input.CreateAccountRequestId = aws.String(_organizationsCreateAccountRequestId)
	}

	if resp, err := client.DescribeCreateAccountStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the contents of the effective policy for specified policy type and
// account. The effective policy is the aggregation of any policies of the
// specified type that the account inherits, plus any policy of that type that is
// directly attached to the account.
//
// This operation applies only to management policies. It does not apply to
// authorization policies: service control policies (SCPs) and resource control
// policies (RCPs).
//
// For more information about policy inheritance, see [Understanding management policy inheritance] in the Organizations User
// Guide.
//
// You can call this operation from any account in a organization.
//
// [Understanding management policy inheritance]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_inheritance_mgmt.html
func organizations_DescribeEffectivePolicy(cfg aws.Config, client *organizations.Client) {
	input := &organizations.DescribeEffectivePolicyInput{
		// PolicyType: types.EffectivePolicyType, // Required
	}

	if len(_organizationsPolicyType) > 0 {
		if err := assignInputField(input, "PolicyType", _organizationsPolicyType); err != nil {
			log.Errorf("invalid --policy-type: %s", err.Error())
			return
		}
	}
	if len(_organizationsTargetId) > 0 {
		input.TargetId = aws.String(_organizationsTargetId)
	}

	if resp, err := client.DescribeEffectivePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns details for a handshake. A handshake is the secure exchange of
// information between two Amazon Web Services accounts: a sender and a recipient.
//
// You can view ACCEPTED , DECLINED , or CANCELED handshakes in API Responses for
// 30 days before they are deleted.
//
// You can call this operation from any account in a organization.
func organizations_DescribeHandshake(cfg aws.Config, client *organizations.Client) {
	input := &organizations.DescribeHandshakeInput{
		// HandshakeId: *string, // Required
	}

	if len(_organizationsHandshakeId) > 0 {
		input.HandshakeId = aws.String(_organizationsHandshakeId)
	}

	if resp, err := client.DescribeHandshake(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the organization that the user's account belongs to.
// You can call this operation from any account in a organization.
//
// Even if a policy type is shown as available in the organization, you can
// disable it separately at the root level with DisablePolicyType. Use ListRoots to see the status of policy
// types for a specified root.
func organizations_DescribeOrganization(cfg aws.Config, client *organizations.Client) {
	input := &organizations.DescribeOrganizationInput{}

	if resp, err := client.DescribeOrganization(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about an organizational unit (OU).
// You can only call this operation from the management account or a member
// account that is a delegated administrator.
func organizations_DescribeOrganizationalUnit(cfg aws.Config, client *organizations.Client) {
	input := &organizations.DescribeOrganizationalUnitInput{
		// OrganizationalUnitId: *string, // Required
	}

	if len(_organizationsOrganizationalUnitId) > 0 {
		input.OrganizationalUnitId = aws.String(_organizationsOrganizationalUnitId)
	}

	if resp, err := client.DescribeOrganizationalUnit(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a policy.
// You can only call this operation from the management account or a member
// account that is a delegated administrator.
func organizations_DescribePolicy(cfg aws.Config, client *organizations.Client) {
	input := &organizations.DescribePolicyInput{
		// PolicyId: *string, // Required
	}

	if len(_organizationsPolicyId) > 0 {
		input.PolicyId = aws.String(_organizationsPolicyId)
	}

	if resp, err := client.DescribePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a resource policy.
// You can only call this operation from the management account or a member
// account that is a delegated administrator.
func organizations_DescribeResourcePolicy(cfg aws.Config, client *organizations.Client) {
	input := &organizations.DescribeResourcePolicyInput{}

	if resp, err := client.DescribeResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns details for a transfer. A transfer is an arrangement between two
// management accounts where one account designates the other with specified
// responsibilities for their organization.
func organizations_DescribeResponsibilityTransfer(cfg aws.Config, client *organizations.Client) {
	input := &organizations.DescribeResponsibilityTransferInput{
		// Id: *string, // Required
	}

	if len(_organizationsId) > 0 {
		input.Id = aws.String(_organizationsId)
	}

	if resp, err := client.DescribeResponsibilityTransfer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Detaches a policy from a target root, organizational unit (OU), or account.
// If the policy being detached is a service control policy (SCP), the changes to
// permissions for Identity and Access Management (IAM) users and roles in affected
// accounts are immediate.
//
// Every root, OU, and account must have at least one SCP attached. If you want to
// replace the default FullAWSAccess policy with an SCP that limits the
// permissions that can be delegated, you must attach the replacement SCP before
// you can remove the default SCP. This is the authorization strategy of an "[allow list] ". If
// you instead attach a second SCP and leave the FullAWSAccess SCP still attached,
// and specify "Effect": "Deny" in the second SCP to override the "Effect": "Allow"
// in the FullAWSAccess policy (or any other attached SCP), you're using the
// authorization strategy of a "[deny list] ".
//
// You can only call this operation from the management account or a member
// account that is a delegated administrator.
//
// [deny list]: https://docs.aws.amazon.com/organizations/latest/userguide/SCP_strategies.html#orgs_policies_denylist
// [allow list]: https://docs.aws.amazon.com/organizations/latest/userguide/SCP_strategies.html#orgs_policies_allowlist
func organizations_DetachPolicy(cfg aws.Config, client *organizations.Client) {
	input := &organizations.DetachPolicyInput{
		// PolicyId: *string, // Required
		// TargetId: *string, // Required
	}

	if len(_organizationsPolicyId) > 0 {
		input.PolicyId = aws.String(_organizationsPolicyId)
	}
	if len(_organizationsTargetId) > 0 {
		input.TargetId = aws.String(_organizationsTargetId)
	}

	if resp, err := client.DetachPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disables the integration of an Amazon Web Services service (the service that is
// specified by ServicePrincipal ) with Organizations. When you disable
// integration, the specified service no longer can create a [service-linked role]in new accounts in
// your organization. This means the service can't perform operations on your
// behalf on any new accounts in your organization. The service can still perform
// operations in older accounts until the service completes its clean-up from
// Organizations.
//
// We strongly recommend that you don't use this command to disable integration
// between Organizations and the specified Amazon Web Services service. Instead,
// use the console or commands that are provided by the specified service. This
// lets the trusted service perform any required initialization when enabling
// trusted access, such as creating any required resources and any required clean
// up of resources when disabling trusted access.
//
// For information about how to disable trusted service access to your
// organization using the trusted service, see the Learn more link under the
// Supports Trusted Access column at [Amazon Web Services services that you can use with Organizations]. on this page.
//
// If you disable access by using this command, it causes the following actions to
// occur:
//
// - The service can no longer create a service-linked role in the accounts in
// your organization. This means that the service can't perform operations on your
// behalf on any new accounts in your organization. The service can still perform
// operations in older accounts until the service completes its clean-up from
// Organizations.
//
// - The service can no longer perform tasks in the member accounts in the
// organization, unless those operations are explicitly permitted by the IAM
// policies that are attached to your roles. This includes any data aggregation
// from the member accounts to the management account, or to a delegated
// administrator account, where relevant.
//
// - Some services detect this and clean up any remaining data or resources
// related to the integration, while other services stop accessing the organization
// but leave any historical data and configuration in place to support a possible
// re-enabling of the integration.
//
// Using the other service's console or commands to disable the integration
// ensures that the other service is aware that it can clean up any resources that
// are required only for the integration. How the service cleans up its resources
// in the organization's accounts depends on that service. For more information,
// see the documentation for the other Amazon Web Services service.
//
// After you perform the DisableAWSServiceAccess operation, the specified service
// can no longer perform operations in your organization's accounts
//
// For more information about integrating other services with Organizations,
// including the list of services that work with Organizations, see [Using Organizations with other Amazon Web Services services]in the
// Organizations User Guide.
//
// You can only call this operation from the management account.
//
// [Amazon Web Services services that you can use with Organizations]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_integrate_services_list.html
// [Using Organizations with other Amazon Web Services services]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_integrate_services.html
// [service-linked role]: https://docs.aws.amazon.com/IAM/latest/UserGuide/using-service-linked-roles.html
func organizations_DisableAWSServiceAccess(cfg aws.Config, client *organizations.Client) {
	input := &organizations.DisableAWSServiceAccessInput{
		// ServicePrincipal: *string, // Required
	}

	if len(_organizationsServicePrincipal) > 0 {
		input.ServicePrincipal = aws.String(_organizationsServicePrincipal)
	}

	if resp, err := client.DisableAWSServiceAccess(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disables an organizational policy type in a root. A policy of a certain type
// can be attached to entities in a root only if that type is enabled in the root.
// After you perform this operation, you no longer can attach policies of the
// specified type to that root or to any organizational unit (OU) or account in
// that root. You can undo this by using the EnablePolicyTypeoperation.
//
// This is an asynchronous request that Amazon Web Services performs in the
// background. If you disable a policy type for a root, it still appears enabled
// for the organization if [all features]are enabled for the organization. Amazon Web Services
// recommends that you first use ListRootsto see the status of policy types for a specified
// root, and then use this operation.
//
// You can only call this operation from the management account or a member
// account that is a delegated administrator.
//
// To view the status of available policy types in the organization, use ListRoots.
//
// [all features]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_org_support-all-features.html
func organizations_DisablePolicyType(cfg aws.Config, client *organizations.Client) {
	input := &organizations.DisablePolicyTypeInput{
		// PolicyType: types.PolicyType, // Required
		// RootId: *string, // Required
	}

	if len(_organizationsPolicyType) > 0 {
		if err := assignInputField(input, "PolicyType", _organizationsPolicyType); err != nil {
			log.Errorf("invalid --policy-type: %s", err.Error())
			return
		}
	}
	if len(_organizationsRootId) > 0 {
		input.RootId = aws.String(_organizationsRootId)
	}

	if resp, err := client.DisablePolicyType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables all features in an organization. This enables the use of organization
// policies that can restrict the services and actions that can be called in each
// account. Until you enable all features, you have access only to consolidated
// billing, and you can't use any of the advanced account administration features
// that Organizations supports. For more information, see [Enabling all features in your organization]in the Organizations
// User Guide.
//
// This operation is required only for organizations that were created explicitly
// with only the consolidated billing features enabled. Calling this operation
// sends a handshake to every invited account in the organization. The feature set
// change can be finalized and the additional features enabled only after all
// administrators in the invited accounts approve the change by accepting the
// handshake.
//
// After you enable all features, you can separately enable or disable individual
// policy types in a root using EnablePolicyTypeand DisablePolicyType. To see the status of policy types in a root,
// use ListRoots.
//
// After all invited member accounts accept the handshake, you finalize the
// feature set change by accepting the handshake that contains "Action":
// "ENABLE_ALL_FEATURES" . This completes the change.
//
// After you enable all features in your organization, the management account in
// the organization can apply policies on all member accounts. These policies can
// restrict what users and even administrators in those accounts can do. The
// management account can apply policies that prevent accounts from leaving the
// organization. Ensure that your account administrators are aware of this.
//
// You can only call this operation from the management account.
//
// [Enabling all features in your organization]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_org_support-all-features.html
func organizations_EnableAllFeatures(cfg aws.Config, client *organizations.Client) {
	input := &organizations.EnableAllFeaturesInput{}

	if resp, err := client.EnableAllFeatures(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides an Amazon Web Services service (the service that is specified by
// ServicePrincipal ) with permissions to view the structure of an organization,
// create a [service-linked role]in all the accounts in the organization, and allow the service to
// perform operations on behalf of the organization and its accounts. Establishing
// these permissions can be a first step in enabling the integration of an Amazon
// Web Services service with Organizations.
//
// We recommend that you enable integration between Organizations and the
// specified Amazon Web Services service by using the console or commands that are
// provided by the specified service. Doing so ensures that the service is aware
// that it can create the resources that are required for the integration. How the
// service creates those resources in the organization's accounts depends on that
// service. For more information, see the documentation for the other Amazon Web
// Services service.
//
// For more information about enabling services to integrate with Organizations,
// see [Using Organizations with other Amazon Web Services services]in the Organizations User Guide.
//
// You can only call this operation from the management account.
//
// [Using Organizations with other Amazon Web Services services]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_integrate_services.html
// [service-linked role]: https://docs.aws.amazon.com/IAM/latest/UserGuide/using-service-linked-roles.html
func organizations_EnableAWSServiceAccess(cfg aws.Config, client *organizations.Client) {
	input := &organizations.EnableAWSServiceAccessInput{
		// ServicePrincipal: *string, // Required
	}

	if len(_organizationsServicePrincipal) > 0 {
		input.ServicePrincipal = aws.String(_organizationsServicePrincipal)
	}

	if resp, err := client.EnableAWSServiceAccess(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables a policy type in a root. After you enable a policy type in a root, you
// can attach policies of that type to the root, any organizational unit (OU), or
// account in that root. You can undo this by using the DisablePolicyTypeoperation.
//
// This is an asynchronous request that Amazon Web Services performs in the
// background. Amazon Web Services recommends that you first use ListRootsto see the status
// of policy types for a specified root, and then use this operation.
//
// You can only call this operation from the management account or a member
// account that is a delegated administrator.
//
// You can enable a policy type in a root only if that policy type is available in
// the organization. To view the status of available policy types in the
// organization, use ListRoots.
func organizations_EnablePolicyType(cfg aws.Config, client *organizations.Client) {
	input := &organizations.EnablePolicyTypeInput{
		// PolicyType: types.PolicyType, // Required
		// RootId: *string, // Required
	}

	if len(_organizationsPolicyType) > 0 {
		if err := assignInputField(input, "PolicyType", _organizationsPolicyType); err != nil {
			log.Errorf("invalid --policy-type: %s", err.Error())
			return
		}
	}
	if len(_organizationsRootId) > 0 {
		input.RootId = aws.String(_organizationsRootId)
	}

	if resp, err := client.EnablePolicyType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sends an invitation to another account to join your organization as a member
// account. Organizations sends email on your behalf to the email address that is
// associated with the other account's owner. The invitation is implemented as a Handshake
// whose details are in the response.
//
// If you receive an exception that indicates that you exceeded your account
// limits for the organization or that the operation failed because your
// organization is still initializing, wait one hour and then try again. If the
// error persists after an hour, contact [Amazon Web Services Support].
//
// If the request includes tags, then the requester must have the
// organizations:TagResource permission.
//
// You can only call this operation from the management account.
//
// [Amazon Web Services Support]: https://console.aws.amazon.com/support/home#/
func organizations_InviteAccountToOrganization(cfg aws.Config, client *organizations.Client) {
	input := &organizations.InviteAccountToOrganizationInput{
		// Target: *types.HandshakeParty, // Required
	}

	if len(_organizationsTarget) > 0 {
		if err := assignInputField(input, "Target", _organizationsTarget); err != nil {
			log.Errorf("invalid --target: %s", err.Error())
			return
		}
	}
	if len(_organizationsNotes) > 0 {
		input.Notes = aws.String(_organizationsNotes)
	}
	if len(_organizationsTags) > 0 {
		if err := assignInputField(input, "Tags", _organizationsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.InviteAccountToOrganization(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sends an invitation to another organization's management account to designate
// your account with the specified responsibilities for their organization. The
// invitation is implemented as a Handshakewhose details are in the response.
//
// You can only call this operation from the management account.
func organizations_InviteOrganizationToTransferResponsibility(cfg aws.Config, client *organizations.Client) {
	input := &organizations.InviteOrganizationToTransferResponsibilityInput{
		// SourceName: *string, // Required
		// StartTimestamp: *time.Time, // Required
		// Target: *types.HandshakeParty, // Required
		// Type: types.ResponsibilityTransferType, // Required
	}

	if len(_organizationsSourceName) > 0 {
		input.SourceName = aws.String(_organizationsSourceName)
	}
	if len(_organizationsStartTimestamp) > 0 {
		if err := assignInputField(input, "StartTimestamp", _organizationsStartTimestamp); err != nil {
			log.Errorf("invalid --start-timestamp: %s", err.Error())
			return
		}
	}
	if len(_organizationsTarget) > 0 {
		if err := assignInputField(input, "Target", _organizationsTarget); err != nil {
			log.Errorf("invalid --target: %s", err.Error())
			return
		}
	}
	if len(_organizationsType) > 0 {
		if err := assignInputField(input, "Type", _organizationsType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_organizationsNotes) > 0 {
		input.Notes = aws.String(_organizationsNotes)
	}
	if len(_organizationsTags) > 0 {
		if err := assignInputField(input, "Tags", _organizationsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.InviteOrganizationToTransferResponsibility(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a member account from its parent organization. This version of the
// operation is performed by the account that wants to leave. To remove a member
// account as a user in the management account, use RemoveAccountFromOrganizationinstead.
//
// You can only call from operation from a member account.
//
// - The management account in an organization with all features enabled can set
// service control policies (SCPs) that can restrict what administrators of member
// accounts can do. This includes preventing them from successfully calling
// LeaveOrganization and leaving the organization.
//
// - You can leave an organization as a member account only if the account is
// configured with the information required to operate as a standalone account.
// When you create an account in an organization using the Organizations console,
// API, or CLI commands, the information required of standalone accounts is not
// automatically collected. For each account that you want to make standalone, you
// must perform the following steps. If any of the steps are already completed for
// this account, that step doesn't appear.
//
// - Choose a support plan
//
// - Provide and verify the required contact information
//
// - Provide a current payment method
//
// # Amazon Web Services uses the payment method to charge for any billable (not
//
// free tier) Amazon Web Services activity that occurs while the account isn't
// attached to an organization. For more information, see [Considerations before removing an account from an organization]in the Organizations
// User Guide.
//
// - The account that you want to leave must not be a delegated administrator
// account for any Amazon Web Services service enabled for your organization. If
// the account is a delegated administrator, you must first change the delegated
// administrator account to another account that is remaining in the organization.
//
// - After the account leaves the organization, all tags that were attached to
// the account object in the organization are deleted. Amazon Web Services accounts
// outside of an organization do not support tags.
//
// - A newly created account has a waiting period before it can be removed from
// its organization. You must wait until at least four days after the account was
// created. Invited accounts aren't subject to this waiting period.
//
// - If you are using an organization principal to call LeaveOrganization across
// multiple accounts, you can only do this up to 5 accounts per second in a single
// organization.
//
// [Considerations before removing an account from an organization]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_account-before-remove.html
func organizations_LeaveOrganization(cfg aws.Config, client *organizations.Client) {
	input := &organizations.LeaveOrganizationInput{}

	if resp, err := client.LeaveOrganization(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all the accounts in the organization. To request only the accounts in a
// specified root or organizational unit (OU), use the ListAccountsForParentoperation instead.
//
// When calling List* operations, always check the NextToken response parameter
// value, even if you receive an empty result set. These operations can
// occasionally return an empty set of results even when more results are
// available. Continue making requests until NextToken returns null. A null
// NextToken value indicates that you have retrieved all available results.
//
// You can only call this operation from the management account or a member
// account that is a delegated administrator.
func organizations_ListAccounts(cfg aws.Config, client *organizations.Client) {
	input := &organizations.ListAccountsInput{}

	if len(_organizationsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _organizationsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_organizationsNextToken) > 0 {
		input.NextToken = aws.String(_organizationsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAccounts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*organizations.ListAccountsOutput
	p := organizations.NewListAccountsPaginator(client, input)
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

// Lists the accounts in an organization that are contained by the specified
// target root or organizational unit (OU). If you specify the root, you get a list
// of all the accounts that aren't in any OU. If you specify an OU, you get a list
// of all the accounts in only that OU and not in any child OUs. To get a list of
// all accounts in the organization, use the ListAccountsoperation.
//
// When calling List* operations, always check the NextToken response parameter
// value, even if you receive an empty result set. These operations can
// occasionally return an empty set of results even when more results are
// available. Continue making requests until NextToken returns null. A null
// NextToken value indicates that you have retrieved all available results.
//
// You can only call this operation from the management account or a member
// account that is a delegated administrator.
func organizations_ListAccountsForParent(cfg aws.Config, client *organizations.Client) {
	input := &organizations.ListAccountsForParentInput{
		// ParentId: *string, // Required
	}

	if len(_organizationsParentId) > 0 {
		input.ParentId = aws.String(_organizationsParentId)
	}
	if len(_organizationsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _organizationsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_organizationsNextToken) > 0 {
		input.NextToken = aws.String(_organizationsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAccountsForParent(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*organizations.ListAccountsForParentOutput
	p := organizations.NewListAccountsForParentPaginator(client, input)
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

// Lists all the accounts in an organization that have invalid effective policies.
// An invalid effective policy is an [effective policy]that fails validation checks, resulting in
// the effective policy not being fully enforced on all the intended accounts
// within an organization.
//
// You can only call this operation from the management account or a member
// account that is a delegated administrator.
//
// [effective policy]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_effective.html
func organizations_ListAccountsWithInvalidEffectivePolicy(cfg aws.Config, client *organizations.Client) {
	input := &organizations.ListAccountsWithInvalidEffectivePolicyInput{
		// PolicyType: types.EffectivePolicyType, // Required
	}

	if len(_organizationsPolicyType) > 0 {
		if err := assignInputField(input, "PolicyType", _organizationsPolicyType); err != nil {
			log.Errorf("invalid --policy-type: %s", err.Error())
			return
		}
	}
	if len(_organizationsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _organizationsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_organizationsNextToken) > 0 {
		input.NextToken = aws.String(_organizationsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAccountsWithInvalidEffectivePolicy(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*organizations.ListAccountsWithInvalidEffectivePolicyOutput
	p := organizations.NewListAccountsWithInvalidEffectivePolicyPaginator(client, input)
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

// Returns a list of the Amazon Web Services services that you enabled to
// integrate with your organization. After a service on this list creates the
// resources that it requires for the integration, it can perform operations on
// your organization and its accounts.
//
// For more information about integrating other services with Organizations,
// including the list of services that currently work with Organizations, see [Using Organizations with other Amazon Web Services services]in
// the Organizations User Guide.
//
// You can only call this operation from the management account or a member
// account that is a delegated administrator.
//
// [Using Organizations with other Amazon Web Services services]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_integrate_services.html
func organizations_ListAWSServiceAccessForOrganization(cfg aws.Config, client *organizations.Client) {
	input := &organizations.ListAWSServiceAccessForOrganizationInput{}

	if len(_organizationsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _organizationsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_organizationsNextToken) > 0 {
		input.NextToken = aws.String(_organizationsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAWSServiceAccessForOrganization(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*organizations.ListAWSServiceAccessForOrganizationOutput
	p := organizations.NewListAWSServiceAccessForOrganizationPaginator(client, input)
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

// Lists all of the organizational units (OUs) or accounts that are contained in
// the specified parent OU or root. This operation, along with ListParentsenables you to
// traverse the tree structure that makes up this root.
//
// When calling List* operations, always check the NextToken response parameter
// value, even if you receive an empty result set. These operations can
// occasionally return an empty set of results even when more results are
// available. Continue making requests until NextToken returns null. A null
// NextToken value indicates that you have retrieved all available results.
//
// You can only call this operation from the management account or a member
// account that is a delegated administrator.
func organizations_ListChildren(cfg aws.Config, client *organizations.Client) {
	input := &organizations.ListChildrenInput{
		// ChildType: types.ChildType, // Required
		// ParentId: *string, // Required
	}

	if len(_organizationsChildType) > 0 {
		if err := assignInputField(input, "ChildType", _organizationsChildType); err != nil {
			log.Errorf("invalid --child-type: %s", err.Error())
			return
		}
	}
	if len(_organizationsParentId) > 0 {
		input.ParentId = aws.String(_organizationsParentId)
	}
	if len(_organizationsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _organizationsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_organizationsNextToken) > 0 {
		input.NextToken = aws.String(_organizationsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListChildren(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*organizations.ListChildrenOutput
	p := organizations.NewListChildrenPaginator(client, input)
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

// Lists the account creation requests that match the specified status that is
// currently being tracked for the organization.
//
// When calling List* operations, always check the NextToken response parameter
// value, even if you receive an empty result set. These operations can
// occasionally return an empty set of results even when more results are
// available. Continue making requests until NextToken returns null. A null
// NextToken value indicates that you have retrieved all available results.
//
// You can only call this operation from the management account or a member
// account that is a delegated administrator.
func organizations_ListCreateAccountStatus(cfg aws.Config, client *organizations.Client) {
	input := &organizations.ListCreateAccountStatusInput{}

	if len(_organizationsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _organizationsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_organizationsNextToken) > 0 {
		input.NextToken = aws.String(_organizationsNextToken)
	}
	if len(_organizationsStates) > 0 {
		if err := assignInputField(input, "States", _organizationsStates); err != nil {
			log.Errorf("invalid --states: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListCreateAccountStatus(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*organizations.ListCreateAccountStatusOutput
	p := organizations.NewListCreateAccountStatusPaginator(client, input)
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

// Lists the Amazon Web Services accounts that are designated as delegated
// administrators in this organization.
//
// You can only call this operation from the management account or a member
// account that is a delegated administrator.
func organizations_ListDelegatedAdministrators(cfg aws.Config, client *organizations.Client) {
	input := &organizations.ListDelegatedAdministratorsInput{}

	if len(_organizationsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _organizationsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_organizationsNextToken) > 0 {
		input.NextToken = aws.String(_organizationsNextToken)
	}
	if len(_organizationsServicePrincipal) > 0 {
		input.ServicePrincipal = aws.String(_organizationsServicePrincipal)
	}

	if disablePaginator() {
		if resp, err := client.ListDelegatedAdministrators(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*organizations.ListDelegatedAdministratorsOutput
	p := organizations.NewListDelegatedAdministratorsPaginator(client, input)
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

// List the Amazon Web Services services for which the specified account is a
// delegated administrator.
//
// You can only call this operation from the management account or a member
// account that is a delegated administrator.
func organizations_ListDelegatedServicesForAccount(cfg aws.Config, client *organizations.Client) {
	input := &organizations.ListDelegatedServicesForAccountInput{
		// AccountId: *string, // Required
	}

	if len(_organizationsAccountId) > 0 {
		input.AccountId = aws.String(_organizationsAccountId)
	}
	if len(_organizationsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _organizationsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_organizationsNextToken) > 0 {
		input.NextToken = aws.String(_organizationsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDelegatedServicesForAccount(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*organizations.ListDelegatedServicesForAccountOutput
	p := organizations.NewListDelegatedServicesForAccountPaginator(client, input)
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

// Lists all the validation errors on an [effective policy] for a specified account and policy type.
// You can only call this operation from the management account or a member
// account that is a delegated administrator.
//
// [effective policy]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_effective.html
func organizations_ListEffectivePolicyValidationErrors(cfg aws.Config, client *organizations.Client) {
	input := &organizations.ListEffectivePolicyValidationErrorsInput{
		// AccountId: *string, // Required
		// PolicyType: types.EffectivePolicyType, // Required
	}

	if len(_organizationsAccountId) > 0 {
		input.AccountId = aws.String(_organizationsAccountId)
	}
	if len(_organizationsPolicyType) > 0 {
		if err := assignInputField(input, "PolicyType", _organizationsPolicyType); err != nil {
			log.Errorf("invalid --policy-type: %s", err.Error())
			return
		}
	}
	if len(_organizationsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _organizationsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_organizationsNextToken) > 0 {
		input.NextToken = aws.String(_organizationsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListEffectivePolicyValidationErrors(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*organizations.ListEffectivePolicyValidationErrorsOutput
	p := organizations.NewListEffectivePolicyValidationErrorsPaginator(client, input)
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

// Lists the recent handshakes that you have received.
// You can view CANCELED , ACCEPTED , DECLINED , or EXPIRED handshakes in API
// responses for 30 days before they are deleted.
//
// You can call this operation from any account in a organization.
//
// When calling List* operations, always check the NextToken response parameter
// value, even if you receive an empty result set. These operations can
// occasionally return an empty set of results even when more results are
// available. Continue making requests until NextToken returns null. A null
// NextToken value indicates that you have retrieved all available results.
func organizations_ListHandshakesForAccount(cfg aws.Config, client *organizations.Client) {
	input := &organizations.ListHandshakesForAccountInput{}

	if len(_organizationsFilter) > 0 {
		if err := assignInputField(input, "Filter", _organizationsFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_organizationsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _organizationsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_organizationsNextToken) > 0 {
		input.NextToken = aws.String(_organizationsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListHandshakesForAccount(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*organizations.ListHandshakesForAccountOutput
	p := organizations.NewListHandshakesForAccountPaginator(client, input)
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

// Lists the recent handshakes that you have sent.
// You can view CANCELED , ACCEPTED , DECLINED , or EXPIRED handshakes in API
// responses for 30 days before they are deleted.
//
// You can only call this operation from the management account or a member
// account that is a delegated administrator.
//
// When calling List* operations, always check the NextToken response parameter
// value, even if you receive an empty result set. These operations can
// occasionally return an empty set of results even when more results are
// available. Continue making requests until NextToken returns null. A null
// NextToken value indicates that you have retrieved all available results.
func organizations_ListHandshakesForOrganization(cfg aws.Config, client *organizations.Client) {
	input := &organizations.ListHandshakesForOrganizationInput{}

	if len(_organizationsFilter) > 0 {
		if err := assignInputField(input, "Filter", _organizationsFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_organizationsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _organizationsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_organizationsNextToken) > 0 {
		input.NextToken = aws.String(_organizationsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListHandshakesForOrganization(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*organizations.ListHandshakesForOrganizationOutput
	p := organizations.NewListHandshakesForOrganizationPaginator(client, input)
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

// Lists transfers that allow you to manage the specified responsibilities for
// another organization. This operation returns both transfer invitations and
// transfers.
//
// When calling List* operations, always check the NextToken response parameter
// value, even if you receive an empty result set. These operations can
// occasionally return an empty set of results even when more results are
// available. Continue making requests until NextToken returns null. A null
// NextToken value indicates that you have retrieved all available results.
func organizations_ListInboundResponsibilityTransfers(cfg aws.Config, client *organizations.Client) {
	input := &organizations.ListInboundResponsibilityTransfersInput{
		// Type: types.ResponsibilityTransferType, // Required
	}

	if len(_organizationsType) > 0 {
		if err := assignInputField(input, "Type", _organizationsType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_organizationsId) > 0 {
		input.Id = aws.String(_organizationsId)
	}
	if len(_organizationsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _organizationsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_organizationsNextToken) > 0 {
		input.NextToken = aws.String(_organizationsNextToken)
	}

	if resp, err := client.ListInboundResponsibilityTransfers(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the organizational units (OUs) in a parent organizational unit or root.
// When calling List* operations, always check the NextToken response parameter
// value, even if you receive an empty result set. These operations can
// occasionally return an empty set of results even when more results are
// available. Continue making requests until NextToken returns null. A null
// NextToken value indicates that you have retrieved all available results.
//
// You can only call this operation from the management account or a member
// account that is a delegated administrator.
func organizations_ListOrganizationalUnitsForParent(cfg aws.Config, client *organizations.Client) {
	input := &organizations.ListOrganizationalUnitsForParentInput{
		// ParentId: *string, // Required
	}

	if len(_organizationsParentId) > 0 {
		input.ParentId = aws.String(_organizationsParentId)
	}
	if len(_organizationsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _organizationsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_organizationsNextToken) > 0 {
		input.NextToken = aws.String(_organizationsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListOrganizationalUnitsForParent(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*organizations.ListOrganizationalUnitsForParentOutput
	p := organizations.NewListOrganizationalUnitsForParentPaginator(client, input)
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

// Lists transfers that allow an account outside your organization to manage the
// specified responsibilities for your organization. This operation returns both
// transfer invitations and transfers.
//
// When calling List* operations, always check the NextToken response parameter
// value, even if you receive an empty result set. These operations can
// occasionally return an empty set of results even when more results are
// available. Continue making requests until NextToken returns null. A null
// NextToken value indicates that you have retrieved all available results.
func organizations_ListOutboundResponsibilityTransfers(cfg aws.Config, client *organizations.Client) {
	input := &organizations.ListOutboundResponsibilityTransfersInput{
		// Type: types.ResponsibilityTransferType, // Required
	}

	if len(_organizationsType) > 0 {
		if err := assignInputField(input, "Type", _organizationsType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_organizationsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _organizationsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_organizationsNextToken) > 0 {
		input.NextToken = aws.String(_organizationsNextToken)
	}

	if resp, err := client.ListOutboundResponsibilityTransfers(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the root or organizational units (OUs) that serve as the immediate parent
// of the specified child OU or account. This operation, along with ListChildrenenables you to
// traverse the tree structure that makes up this root.
//
// When calling List* operations, always check the NextToken response parameter
// value, even if you receive an empty result set. These operations can
// occasionally return an empty set of results even when more results are
// available. Continue making requests until NextToken returns null. A null
// NextToken value indicates that you have retrieved all available results.
//
// You can only call this operation from the management account or a member
// account that is a delegated administrator.
//
// In the current release, a child can have only a single parent.
func organizations_ListParents(cfg aws.Config, client *organizations.Client) {
	input := &organizations.ListParentsInput{
		// ChildId: *string, // Required
	}

	if len(_organizationsChildId) > 0 {
		input.ChildId = aws.String(_organizationsChildId)
	}
	if len(_organizationsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _organizationsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_organizationsNextToken) > 0 {
		input.NextToken = aws.String(_organizationsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListParents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*organizations.ListParentsOutput
	p := organizations.NewListParentsPaginator(client, input)
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

// Retrieves the list of all policies in an organization of a specified type.
// When calling List* operations, always check the NextToken response parameter
// value, even if you receive an empty result set. These operations can
// occasionally return an empty set of results even when more results are
// available. Continue making requests until NextToken returns null. A null
// NextToken value indicates that you have retrieved all available results.
//
// You can only call this operation from the management account or a member
// account that is a delegated administrator.
func organizations_ListPolicies(cfg aws.Config, client *organizations.Client) {
	input := &organizations.ListPoliciesInput{
		// Filter: types.PolicyType, // Required
	}

	if len(_organizationsFilter) > 0 {
		if err := assignInputField(input, "Filter", _organizationsFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_organizationsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _organizationsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_organizationsNextToken) > 0 {
		input.NextToken = aws.String(_organizationsNextToken)
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

	var results []*organizations.ListPoliciesOutput
	p := organizations.NewListPoliciesPaginator(client, input)
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

// Lists the policies that are directly attached to the specified target root,
// organizational unit (OU), or account. You must specify the policy type that you
// want included in the returned list.
//
// When calling List* operations, always check the NextToken response parameter
// value, even if you receive an empty result set. These operations can
// occasionally return an empty set of results even when more results are
// available. Continue making requests until NextToken returns null. A null
// NextToken value indicates that you have retrieved all available results.
//
// You can only call this operation from the management account or a member
// account that is a delegated administrator.
func organizations_ListPoliciesForTarget(cfg aws.Config, client *organizations.Client) {
	input := &organizations.ListPoliciesForTargetInput{
		// Filter: types.PolicyType, // Required
		// TargetId: *string, // Required
	}

	if len(_organizationsFilter) > 0 {
		if err := assignInputField(input, "Filter", _organizationsFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_organizationsTargetId) > 0 {
		input.TargetId = aws.String(_organizationsTargetId)
	}
	if len(_organizationsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _organizationsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_organizationsNextToken) > 0 {
		input.NextToken = aws.String(_organizationsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPoliciesForTarget(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*organizations.ListPoliciesForTargetOutput
	p := organizations.NewListPoliciesForTargetPaginator(client, input)
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

// Lists the roots that are defined in the current organization.
// When calling List* operations, always check the NextToken response parameter
// value, even if you receive an empty result set. These operations can
// occasionally return an empty set of results even when more results are
// available. Continue making requests until NextToken returns null. A null
// NextToken value indicates that you have retrieved all available results.
//
// You can only call this operation from the management account or a member
// account that is a delegated administrator.
//
// Policy types can be enabled and disabled in roots. This is distinct from
// whether they're available in the organization. When you enable all features, you
// make policy types available for use in that organization. Individual policy
// types can then be enabled and disabled in a root. To see the availability of a
// policy type in an organization, use DescribeOrganization.
func organizations_ListRoots(cfg aws.Config, client *organizations.Client) {
	input := &organizations.ListRootsInput{}

	if len(_organizationsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _organizationsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_organizationsNextToken) > 0 {
		input.NextToken = aws.String(_organizationsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRoots(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*organizations.ListRootsOutput
	p := organizations.NewListRootsPaginator(client, input)
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

// Lists tags that are attached to the specified resource.
// You can attach tags to the following resources in Organizations.
//
// - Amazon Web Services account
//
// - Organization root
//
// - Organizational unit (OU)
//
// - Policy (any type)
//
// You can only call this operation from the management account or a member
// account that is a delegated administrator.
func organizations_ListTagsForResource(cfg aws.Config, client *organizations.Client) {
	input := &organizations.ListTagsForResourceInput{
		// ResourceId: *string, // Required
	}

	if len(_organizationsResourceId) > 0 {
		input.ResourceId = aws.String(_organizationsResourceId)
	}
	if len(_organizationsNextToken) > 0 {
		input.NextToken = aws.String(_organizationsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*organizations.ListTagsForResourceOutput
	p := organizations.NewListTagsForResourcePaginator(client, input)
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

// Lists all the roots, organizational units (OUs), and accounts that the
// specified policy is attached to.
//
// When calling List* operations, always check the NextToken response parameter
// value, even if you receive an empty result set. These operations can
// occasionally return an empty set of results even when more results are
// available. Continue making requests until NextToken returns null. A null
// NextToken value indicates that you have retrieved all available results.
//
// You can only call this operation from the management account or a member
// account that is a delegated administrator.
func organizations_ListTargetsForPolicy(cfg aws.Config, client *organizations.Client) {
	input := &organizations.ListTargetsForPolicyInput{
		// PolicyId: *string, // Required
	}

	if len(_organizationsPolicyId) > 0 {
		input.PolicyId = aws.String(_organizationsPolicyId)
	}
	if len(_organizationsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _organizationsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_organizationsNextToken) > 0 {
		input.NextToken = aws.String(_organizationsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTargetsForPolicy(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*organizations.ListTargetsForPolicyOutput
	p := organizations.NewListTargetsForPolicyPaginator(client, input)
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

// Moves an account from its current source parent root or organizational unit
// (OU) to the specified destination parent root or OU.
//
// You can only call this operation from the management account.
func organizations_MoveAccount(cfg aws.Config, client *organizations.Client) {
	input := &organizations.MoveAccountInput{
		// AccountId: *string, // Required
		// DestinationParentId: *string, // Required
		// SourceParentId: *string, // Required
	}

	if len(_organizationsAccountId) > 0 {
		input.AccountId = aws.String(_organizationsAccountId)
	}
	if len(_organizationsDestinationParentId) > 0 {
		input.DestinationParentId = aws.String(_organizationsDestinationParentId)
	}
	if len(_organizationsSourceParentId) > 0 {
		input.SourceParentId = aws.String(_organizationsSourceParentId)
	}

	if resp, err := client.MoveAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates a resource policy.
// You can only call this operation from the management account..
func organizations_PutResourcePolicy(cfg aws.Config, client *organizations.Client) {
	input := &organizations.PutResourcePolicyInput{
		// Content: *string, // Required
	}

	if len(_organizationsContent) > 0 {
		input.Content = aws.String(_organizationsContent)
	}
	if len(_organizationsTags) > 0 {
		if err := assignInputField(input, "Tags", _organizationsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables the specified member account to administer the Organizations features
// of the specified Amazon Web Services service. It grants read-only access to
// Organizations service data. The account still requires IAM permissions to access
// and administer the Amazon Web Services service.
//
// You can run this action only for Amazon Web Services services that support this
// feature. For a current list of services that support it, see the column Supports
// Delegated Administrator in the table at [Amazon Web Services Services that you can use with Organizations]in the Organizations User Guide.
//
// You can only call this operation from the management account.
//
// [Amazon Web Services Services that you can use with Organizations]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_integrate_services_list.html
func organizations_RegisterDelegatedAdministrator(cfg aws.Config, client *organizations.Client) {
	input := &organizations.RegisterDelegatedAdministratorInput{
		// AccountId: *string, // Required
		// ServicePrincipal: *string, // Required
	}

	if len(_organizationsAccountId) > 0 {
		input.AccountId = aws.String(_organizationsAccountId)
	}
	if len(_organizationsServicePrincipal) > 0 {
		input.ServicePrincipal = aws.String(_organizationsServicePrincipal)
	}

	if resp, err := client.RegisterDelegatedAdministrator(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the specified account from the organization.
// The removed account becomes a standalone account that isn't a member of any
// organization. It's no longer subject to any policies and is responsible for its
// own bill payments. The organization's management account is no longer charged
// for any expenses accrued by the member account after it's removed from the
// organization.
//
// You can only call this operation from the management account. Member accounts
// can remove themselves with LeaveOrganizationinstead.
//
// - You can remove an account from your organization only if the account is
// configured with the information required to operate as a standalone account.
// When you create an account in an organization using the Organizations console,
// API, or CLI commands, the information required of standalone accounts is not
// automatically collected. For more information, see [Considerations before removing an account from an organization]in the Organizations User
// Guide.
//
// - The account that you want to leave must not be a delegated administrator
// account for any Amazon Web Services service enabled for your organization. If
// the account is a delegated administrator, you must first change the delegated
// administrator account to another account that is remaining in the organization.
//
// - After the account leaves the organization, all tags that were attached to
// the account object in the organization are deleted. Amazon Web Services accounts
// outside of an organization do not support tags.
//
// [Considerations before removing an account from an organization]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_account-before-remove.html
func organizations_RemoveAccountFromOrganization(cfg aws.Config, client *organizations.Client) {
	input := &organizations.RemoveAccountFromOrganizationInput{
		// AccountId: *string, // Required
	}

	if len(_organizationsAccountId) > 0 {
		input.AccountId = aws.String(_organizationsAccountId)
	}

	if resp, err := client.RemoveAccountFromOrganization(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds one or more tags to the specified resource.
// Currently, you can attach tags to the following resources in Organizations.
//
// - Amazon Web Services account
//
// - Organization root
//
// - Organizational unit (OU)
//
// - Policy (any type)
//
// You can only call this operation from the management account or a member
// account that is a delegated administrator.
func organizations_TagResource(cfg aws.Config, client *organizations.Client) {
	input := &organizations.TagResourceInput{
		// ResourceId: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_organizationsResourceId) > 0 {
		input.ResourceId = aws.String(_organizationsResourceId)
	}
	if len(_organizationsTags) > 0 {
		if err := assignInputField(input, "Tags", _organizationsTags); err != nil {
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

// Ends a transfer. A transfer is an arrangement between two management accounts
// where one account designates the other with specified responsibilities for their
// organization.
func organizations_TerminateResponsibilityTransfer(cfg aws.Config, client *organizations.Client) {
	input := &organizations.TerminateResponsibilityTransferInput{
		// Id: *string, // Required
	}

	if len(_organizationsId) > 0 {
		input.Id = aws.String(_organizationsId)
	}
	if len(_organizationsEndTimestamp) > 0 {
		if err := assignInputField(input, "EndTimestamp", _organizationsEndTimestamp); err != nil {
			log.Errorf("invalid --end-timestamp: %s", err.Error())
			return
		}
	}

	if resp, err := client.TerminateResponsibilityTransfer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes any tags with the specified keys from the specified resource.
// You can attach tags to the following resources in Organizations.
//
// - Amazon Web Services account
//
// - Organization root
//
// - Organizational unit (OU)
//
// - Policy (any type)
//
// You can only call this operation from the management account or a member
// account that is a delegated administrator.
func organizations_UntagResource(cfg aws.Config, client *organizations.Client) {
	input := &organizations.UntagResourceInput{
		// ResourceId: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_organizationsResourceId) > 0 {
		input.ResourceId = aws.String(_organizationsResourceId)
	}
	if len(_organizationsTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _organizationsTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Renames the specified organizational unit (OU). The ID and ARN don't change.
// The child OUs and accounts remain in place, and any attached policies of the OU
// remain attached.
//
// You can only call this operation from the management account.
func organizations_UpdateOrganizationalUnit(cfg aws.Config, client *organizations.Client) {
	input := &organizations.UpdateOrganizationalUnitInput{
		// OrganizationalUnitId: *string, // Required
	}

	if len(_organizationsOrganizationalUnitId) > 0 {
		input.OrganizationalUnitId = aws.String(_organizationsOrganizationalUnitId)
	}
	if len(_organizationsName) > 0 {
		input.Name = aws.String(_organizationsName)
	}

	if resp, err := client.UpdateOrganizationalUnit(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing policy with a new name, description, or content. If you
// don't supply any parameter, that value remains unchanged. You can't change a
// policy's type.
//
// You can only call this operation from the management account or a member
// account that is a delegated administrator.
func organizations_UpdatePolicy(cfg aws.Config, client *organizations.Client) {
	input := &organizations.UpdatePolicyInput{
		// PolicyId: *string, // Required
	}

	if len(_organizationsPolicyId) > 0 {
		input.PolicyId = aws.String(_organizationsPolicyId)
	}
	if len(_organizationsContent) > 0 {
		input.Content = aws.String(_organizationsContent)
	}
	if len(_organizationsDescription) > 0 {
		input.Description = aws.String(_organizationsDescription)
	}
	if len(_organizationsName) > 0 {
		input.Name = aws.String(_organizationsName)
	}

	if resp, err := client.UpdatePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a transfer. A transfer is the arrangement between two management
// accounts where one account designates the other with specified responsibilities
// for their organization.
//
// You can update the name assigned to a transfer.
func organizations_UpdateResponsibilityTransfer(cfg aws.Config, client *organizations.Client) {
	input := &organizations.UpdateResponsibilityTransferInput{
		// Id: *string, // Required
		// Name: *string, // Required
	}

	if len(_organizationsId) > 0 {
		input.Id = aws.String(_organizationsId)
	}
	if len(_organizationsName) > 0 {
		input.Name = aws.String(_organizationsName)
	}

	if resp, err := client.UpdateResponsibilityTransfer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_organizationsCmd)
	_organizationsCmd.Flags().SortFlags = false

	_organizationsCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_organizationsCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_organizationsCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_organizationsCmd.Flags().StringVarP(&_organizationsAccountId, "account-id", "", "", "Account ID")
	_organizationsCmd.Flags().StringVarP(&_organizationsAccountName, "account-name", "", "", "Account Name")
	_organizationsCmd.Flags().StringVarP(&_organizationsChildId, "child-id", "", "", "Child ID")
	_organizationsCmd.Flags().StringVarP(&_organizationsChildType, "child-type", "", "", "Child Type")
	_organizationsCmd.Flags().StringVarP(&_organizationsContent, "content", "", "", "Content")
	_organizationsCmd.Flags().StringVarP(&_organizationsCreateAccountRequestId, "create-account-request-id", "", "", "Create Account Request ID")
	_organizationsCmd.Flags().StringVarP(&_organizationsDescription, "description", "", "", "Description")
	_organizationsCmd.Flags().StringVarP(&_organizationsDestinationParentId, "destination-parent-id", "", "", "Destination Parent ID")
	_organizationsCmd.Flags().StringVarP(&_organizationsEmail, "email", "", "", "Email")
	_organizationsCmd.Flags().StringVarP(&_organizationsEndTimestamp, "end-timestamp", "", "", "End Timestamp")
	_organizationsCmd.Flags().StringVarP(&_organizationsFeatureSet, "feature-set", "", "", "Feature Set")
	_organizationsCmd.Flags().StringVarP(&_organizationsFilter, "filter", "", "", "Filter")
	_organizationsCmd.Flags().StringVarP(&_organizationsHandshakeId, "handshake-id", "", "", "Handshake ID")
	_organizationsCmd.Flags().StringVarP(&_organizationsIamUserAccessToBilling, "iam-user-access-to-billing", "", "", "IAM User Access To Billing")
	_organizationsCmd.Flags().StringVarP(&_organizationsId, "id", "", "", "ID")
	_organizationsCmd.Flags().StringVarP(&_organizationsMaxResults, "max-results", "", "", "Max Results")
	_organizationsCmd.Flags().StringVarP(&_organizationsName, "name", "", "", "Name")
	_organizationsCmd.Flags().StringVarP(&_organizationsNextToken, "next-token", "", "", "Next Token")
	_organizationsCmd.Flags().StringVarP(&_organizationsNotes, "notes", "", "", "Notes")
	_organizationsCmd.Flags().StringVarP(&_organizationsOrganizationalUnitId, "organizational-unit-id", "", "", "Organizational Unit ID")
	_organizationsCmd.Flags().StringVarP(&_organizationsParentId, "parent-id", "", "", "Parent ID")
	_organizationsCmd.Flags().StringVarP(&_organizationsPolicyId, "policy-id", "", "", "Policy ID")
	_organizationsCmd.Flags().StringVarP(&_organizationsPolicyType, "policy-type", "", "", "Policy Type")
	_organizationsCmd.Flags().StringVarP(&_organizationsResourceId, "resource-id", "", "", "Resource ID")
	_organizationsCmd.Flags().StringVarP(&_organizationsRoleName, "role-name", "", "", "Role Name")
	_organizationsCmd.Flags().StringVarP(&_organizationsRootId, "root-id", "", "", "Root ID")
	_organizationsCmd.Flags().StringVarP(&_organizationsServicePrincipal, "service-principal", "", "", "Service Principal")
	_organizationsCmd.Flags().StringVarP(&_organizationsSourceName, "source-name", "", "", "Source Name")
	_organizationsCmd.Flags().StringVarP(&_organizationsSourceParentId, "source-parent-id", "", "", "Source Parent ID")
	_organizationsCmd.Flags().StringVarP(&_organizationsStartTimestamp, "start-timestamp", "", "", "Start Timestamp")
	_organizationsCmd.Flags().StringVarP(&_organizationsStates, "states", "", "", "States")
	_organizationsCmd.Flags().StringSliceVarP(&_organizationsTagKeys, "tag-keys", "", nil, "Tag Keys")
	_organizationsCmd.Flags().StringVarP(&_organizationsTags, "tags", "", "", "Tags")
	_organizationsCmd.Flags().StringVarP(&_organizationsTarget, "target", "", "", "Target")
	_organizationsCmd.Flags().StringVarP(&_organizationsTargetId, "target-id", "", "", "Target ID")
	_organizationsCmd.Flags().StringVarP(&_organizationsType, "type", "", "", "Type")

	_organizationsCmd.Flags().BoolVarP(&_organizationsAcceptHandshake, "accept-handshake", "", false, "Accept Handshake")
	_organizationsCmd.Flags().BoolVarP(&_organizationsAttachPolicy, "attach-policy", "", false, "Attach Policy")
	_organizationsCmd.Flags().BoolVarP(&_organizationsCancelHandshake, "cancel-handshake", "", false, "Cancel Handshake")
	_organizationsCmd.Flags().BoolVarP(&_organizationsCloseAccount, "close-account", "", false, "Close Account")
	_organizationsCmd.Flags().BoolVarP(&_organizationsCreateAccount, "create-account", "", false, "Create Account")
	_organizationsCmd.Flags().BoolVarP(&_organizationsCreateGovCloudAccount, "create-gov-cloud-account", "", false, "Create Gov Cloud Account")
	_organizationsCmd.Flags().BoolVarP(&_organizationsCreateOrganization, "create-organization", "", false, "Create Organization")
	_organizationsCmd.Flags().BoolVarP(&_organizationsCreateOrganizationalUnit, "create-organizational-unit", "", false, "Create Organizational Unit")
	_organizationsCmd.Flags().BoolVarP(&_organizationsCreatePolicy, "create-policy", "", false, "Create Policy")
	_organizationsCmd.Flags().BoolVarP(&_organizationsDeclineHandshake, "decline-handshake", "", false, "Decline Handshake")
	_organizationsCmd.Flags().BoolVarP(&_organizationsDeleteOrganization, "delete-organization", "", false, "Delete Organization")
	_organizationsCmd.Flags().BoolVarP(&_organizationsDeleteOrganizationalUnit, "delete-organizational-unit", "", false, "Delete Organizational Unit")
	_organizationsCmd.Flags().BoolVarP(&_organizationsDeletePolicy, "delete-policy", "", false, "Delete Policy")
	_organizationsCmd.Flags().BoolVarP(&_organizationsDeleteResourcePolicy, "delete-resource-policy", "", false, "Delete Resource Policy")
	_organizationsCmd.Flags().BoolVarP(&_organizationsDeregisterDelegatedAdministrator, "deregister-delegated-administrator", "", false, "Deregister Delegated Administrator")
	_organizationsCmd.Flags().BoolVarP(&_organizationsDescribeAccount, "describe-account", "", false, "Describe Account")
	_organizationsCmd.Flags().BoolVarP(&_organizationsDescribeCreateAccountStatus, "describe-create-account-status", "", false, "Describe Create Account Status")
	_organizationsCmd.Flags().BoolVarP(&_organizationsDescribeEffectivePolicy, "describe-effective-policy", "", false, "Describe Effective Policy")
	_organizationsCmd.Flags().BoolVarP(&_organizationsDescribeHandshake, "describe-handshake", "", false, "Describe Handshake")
	_organizationsCmd.Flags().BoolVarP(&_organizationsDescribeOrganization, "describe-organization", "", false, "Describe Organization")
	_organizationsCmd.Flags().BoolVarP(&_organizationsDescribeOrganizationalUnit, "describe-organizational-unit", "", false, "Describe Organizational Unit")
	_organizationsCmd.Flags().BoolVarP(&_organizationsDescribePolicy, "describe-policy", "", false, "Describe Policy")
	_organizationsCmd.Flags().BoolVarP(&_organizationsDescribeResourcePolicy, "describe-resource-policy", "", false, "Describe Resource Policy")
	_organizationsCmd.Flags().BoolVarP(&_organizationsDescribeResponsibilityTransfer, "describe-responsibility-transfer", "", false, "Describe Responsibility Transfer")
	_organizationsCmd.Flags().BoolVarP(&_organizationsDetachPolicy, "detach-policy", "", false, "Detach Policy")
	_organizationsCmd.Flags().BoolVarP(&_organizationsDisableAWSServiceAccess, "disable-aws-service-access", "", false, "Disable AWS Service Access")
	_organizationsCmd.Flags().BoolVarP(&_organizationsDisablePolicyType, "disable-policy-type", "", false, "Disable Policy Type")
	_organizationsCmd.Flags().BoolVarP(&_organizationsEnableAllFeatures, "enable-all-features", "", false, "Enable All Features")
	_organizationsCmd.Flags().BoolVarP(&_organizationsEnableAWSServiceAccess, "enable-aws-service-access", "", false, "Enable AWS Service Access")
	_organizationsCmd.Flags().BoolVarP(&_organizationsEnablePolicyType, "enable-policy-type", "", false, "Enable Policy Type")
	_organizationsCmd.Flags().BoolVarP(&_organizationsInviteAccountToOrganization, "invite-account-to-organization", "", false, "Invite Account To Organization")
	_organizationsCmd.Flags().BoolVarP(&_organizationsInviteOrganizationToTransferResponsibility, "invite-organization-to-transfer-responsibility", "", false, "Invite Organization To Transfer Responsibility")
	_organizationsCmd.Flags().BoolVarP(&_organizationsLeaveOrganization, "leave-organization", "", false, "Leave Organization")
	_organizationsCmd.Flags().BoolVarP(&_organizationsListAccounts, "list-accounts", "", false, "List Accounts")
	_organizationsCmd.Flags().BoolVarP(&_organizationsListAccountsForParent, "list-accounts-for-parent", "", false, "List Accounts For Parent")
	_organizationsCmd.Flags().BoolVarP(&_organizationsListAccountsWithInvalidEffectivePolicy, "list-accounts-with-invalid-effective-policy", "", false, "List Accounts With Invalid Effective Policy")
	_organizationsCmd.Flags().BoolVarP(&_organizationsListAWSServiceAccessForOrganization, "list-aws-service-access-for-organization", "", false, "List AWS Service Access For Organization")
	_organizationsCmd.Flags().BoolVarP(&_organizationsListChildren, "list-children", "", false, "List Children")
	_organizationsCmd.Flags().BoolVarP(&_organizationsListCreateAccountStatus, "list-create-account-status", "", false, "List Create Account Status")
	_organizationsCmd.Flags().BoolVarP(&_organizationsListDelegatedAdministrators, "list-delegated-administrators", "", false, "List Delegated Administrators")
	_organizationsCmd.Flags().BoolVarP(&_organizationsListDelegatedServicesForAccount, "list-delegated-services-for-account", "", false, "List Delegated Services For Account")
	_organizationsCmd.Flags().BoolVarP(&_organizationsListEffectivePolicyValidationErrors, "list-effective-policy-validation-errors", "", false, "List Effective Policy Validation Errors")
	_organizationsCmd.Flags().BoolVarP(&_organizationsListHandshakesForAccount, "list-handshakes-for-account", "", false, "List Handshakes For Account")
	_organizationsCmd.Flags().BoolVarP(&_organizationsListHandshakesForOrganization, "list-handshakes-for-organization", "", false, "List Handshakes For Organization")
	_organizationsCmd.Flags().BoolVarP(&_organizationsListInboundResponsibilityTransfers, "list-inbound-responsibility-transfers", "", false, "List Inbound Responsibility Transfers")
	_organizationsCmd.Flags().BoolVarP(&_organizationsListOrganizationalUnitsForParent, "list-organizational-units-for-parent", "", false, "List Organizational Units For Parent")
	_organizationsCmd.Flags().BoolVarP(&_organizationsListOutboundResponsibilityTransfers, "list-outbound-responsibility-transfers", "", false, "List Outbound Responsibility Transfers")
	_organizationsCmd.Flags().BoolVarP(&_organizationsListParents, "list-parents", "", false, "List Parents")
	_organizationsCmd.Flags().BoolVarP(&_organizationsListPolicies, "list-policies", "", false, "List Policies")
	_organizationsCmd.Flags().BoolVarP(&_organizationsListPoliciesForTarget, "list-policies-for-target", "", false, "List Policies For Target")
	_organizationsCmd.Flags().BoolVarP(&_organizationsListRoots, "list-roots", "", false, "List Roots")
	_organizationsCmd.Flags().BoolVarP(&_organizationsListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_organizationsCmd.Flags().BoolVarP(&_organizationsListTargetsForPolicy, "list-targets-for-policy", "", false, "List Targets For Policy")
	_organizationsCmd.Flags().BoolVarP(&_organizationsMoveAccount, "move-account", "", false, "Move Account")
	_organizationsCmd.Flags().BoolVarP(&_organizationsPutResourcePolicy, "put-resource-policy", "", false, "Put Resource Policy")
	_organizationsCmd.Flags().BoolVarP(&_organizationsRegisterDelegatedAdministrator, "register-delegated-administrator", "", false, "Register Delegated Administrator")
	_organizationsCmd.Flags().BoolVarP(&_organizationsRemoveAccountFromOrganization, "remove-account-from-organization", "", false, "Remove Account From Organization")
	_organizationsCmd.Flags().BoolVarP(&_organizationsTagResource, "tag-resource", "", false, "Tag Resource")
	_organizationsCmd.Flags().BoolVarP(&_organizationsTerminateResponsibilityTransfer, "terminate-responsibility-transfer", "", false, "Terminate Responsibility Transfer")
	_organizationsCmd.Flags().BoolVarP(&_organizationsUntagResource, "untag-resource", "", false, "Untag Resource")
	_organizationsCmd.Flags().BoolVarP(&_organizationsUpdateOrganizationalUnit, "update-organizational-unit", "", false, "Update Organizational Unit")
	_organizationsCmd.Flags().BoolVarP(&_organizationsUpdatePolicy, "update-policy", "", false, "Update Policy")
	_organizationsCmd.Flags().BoolVarP(&_organizationsUpdateResponsibilityTransfer, "update-responsibility-transfer", "", false, "Update Responsibility Transfer")

}
