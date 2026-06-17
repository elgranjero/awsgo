package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// ssoadminCmd represents the ssoadmin command
var _ssoadminCmd = &cobra.Command{
	Use:   "ssoadmin",
	Short: "AWS ssoadmin CLI",
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
		client := ssoadmin.NewFromConfig(cfg)
		if _ssoadminAddRegion {
			ssoadmin_AddRegion(cfg, client)
			return
		}
		if _ssoadminAttachCustomerManagedPolicyReferenceToPermissionSet {
			ssoadmin_AttachCustomerManagedPolicyReferenceToPermissionSet(cfg, client)
			return
		}
		if _ssoadminAttachManagedPolicyToPermissionSet {
			ssoadmin_AttachManagedPolicyToPermissionSet(cfg, client)
			return
		}
		if _ssoadminCreateAccountAssignment {
			ssoadmin_CreateAccountAssignment(cfg, client)
			return
		}
		if _ssoadminCreateApplication {
			ssoadmin_CreateApplication(cfg, client)
			return
		}
		if _ssoadminCreateApplicationAssignment {
			ssoadmin_CreateApplicationAssignment(cfg, client)
			return
		}
		if _ssoadminCreateInstance {
			ssoadmin_CreateInstance(cfg, client)
			return
		}
		if _ssoadminCreateInstanceAccessControlAttributeConfiguration {
			ssoadmin_CreateInstanceAccessControlAttributeConfiguration(cfg, client)
			return
		}
		if _ssoadminCreatePermissionSet {
			ssoadmin_CreatePermissionSet(cfg, client)
			return
		}
		if _ssoadminCreateTrustedTokenIssuer {
			ssoadmin_CreateTrustedTokenIssuer(cfg, client)
			return
		}
		if _ssoadminDeleteAccountAssignment {
			ssoadmin_DeleteAccountAssignment(cfg, client)
			return
		}
		if _ssoadminDeleteApplication {
			ssoadmin_DeleteApplication(cfg, client)
			return
		}
		if _ssoadminDeleteApplicationAccessScope {
			ssoadmin_DeleteApplicationAccessScope(cfg, client)
			return
		}
		if _ssoadminDeleteApplicationAssignment {
			ssoadmin_DeleteApplicationAssignment(cfg, client)
			return
		}
		if _ssoadminDeleteApplicationAuthenticationMethod {
			ssoadmin_DeleteApplicationAuthenticationMethod(cfg, client)
			return
		}
		if _ssoadminDeleteApplicationGrant {
			ssoadmin_DeleteApplicationGrant(cfg, client)
			return
		}
		if _ssoadminDeleteInlinePolicyFromPermissionSet {
			ssoadmin_DeleteInlinePolicyFromPermissionSet(cfg, client)
			return
		}
		if _ssoadminDeleteInstance {
			ssoadmin_DeleteInstance(cfg, client)
			return
		}
		if _ssoadminDeleteInstanceAccessControlAttributeConfiguration {
			ssoadmin_DeleteInstanceAccessControlAttributeConfiguration(cfg, client)
			return
		}
		if _ssoadminDeletePermissionSet {
			ssoadmin_DeletePermissionSet(cfg, client)
			return
		}
		if _ssoadminDeletePermissionsBoundaryFromPermissionSet {
			ssoadmin_DeletePermissionsBoundaryFromPermissionSet(cfg, client)
			return
		}
		if _ssoadminDeleteTrustedTokenIssuer {
			ssoadmin_DeleteTrustedTokenIssuer(cfg, client)
			return
		}
		if _ssoadminDescribeAccountAssignmentCreationStatus {
			ssoadmin_DescribeAccountAssignmentCreationStatus(cfg, client)
			return
		}
		if _ssoadminDescribeAccountAssignmentDeletionStatus {
			ssoadmin_DescribeAccountAssignmentDeletionStatus(cfg, client)
			return
		}
		if _ssoadminDescribeApplication {
			ssoadmin_DescribeApplication(cfg, client)
			return
		}
		if _ssoadminDescribeApplicationAssignment {
			ssoadmin_DescribeApplicationAssignment(cfg, client)
			return
		}
		if _ssoadminDescribeApplicationProvider {
			ssoadmin_DescribeApplicationProvider(cfg, client)
			return
		}
		if _ssoadminDescribeInstance {
			ssoadmin_DescribeInstance(cfg, client)
			return
		}
		if _ssoadminDescribeInstanceAccessControlAttributeConfiguration {
			ssoadmin_DescribeInstanceAccessControlAttributeConfiguration(cfg, client)
			return
		}
		if _ssoadminDescribePermissionSet {
			ssoadmin_DescribePermissionSet(cfg, client)
			return
		}
		if _ssoadminDescribePermissionSetProvisioningStatus {
			ssoadmin_DescribePermissionSetProvisioningStatus(cfg, client)
			return
		}
		if _ssoadminDescribeRegion {
			ssoadmin_DescribeRegion(cfg, client)
			return
		}
		if _ssoadminDescribeTrustedTokenIssuer {
			ssoadmin_DescribeTrustedTokenIssuer(cfg, client)
			return
		}
		if _ssoadminDetachCustomerManagedPolicyReferenceFromPermissionSet {
			ssoadmin_DetachCustomerManagedPolicyReferenceFromPermissionSet(cfg, client)
			return
		}
		if _ssoadminDetachManagedPolicyFromPermissionSet {
			ssoadmin_DetachManagedPolicyFromPermissionSet(cfg, client)
			return
		}
		if _ssoadminGetApplicationAccessScope {
			ssoadmin_GetApplicationAccessScope(cfg, client)
			return
		}
		if _ssoadminGetApplicationAssignmentConfiguration {
			ssoadmin_GetApplicationAssignmentConfiguration(cfg, client)
			return
		}
		if _ssoadminGetApplicationAuthenticationMethod {
			ssoadmin_GetApplicationAuthenticationMethod(cfg, client)
			return
		}
		if _ssoadminGetApplicationGrant {
			ssoadmin_GetApplicationGrant(cfg, client)
			return
		}
		if _ssoadminGetApplicationSessionConfiguration {
			ssoadmin_GetApplicationSessionConfiguration(cfg, client)
			return
		}
		if _ssoadminGetInlinePolicyForPermissionSet {
			ssoadmin_GetInlinePolicyForPermissionSet(cfg, client)
			return
		}
		if _ssoadminGetPermissionsBoundaryForPermissionSet {
			ssoadmin_GetPermissionsBoundaryForPermissionSet(cfg, client)
			return
		}
		if _ssoadminListAccountAssignmentCreationStatus {
			ssoadmin_ListAccountAssignmentCreationStatus(cfg, client)
			return
		}
		if _ssoadminListAccountAssignmentDeletionStatus {
			ssoadmin_ListAccountAssignmentDeletionStatus(cfg, client)
			return
		}
		if _ssoadminListAccountAssignments {
			ssoadmin_ListAccountAssignments(cfg, client)
			return
		}
		if _ssoadminListAccountAssignmentsForPrincipal {
			ssoadmin_ListAccountAssignmentsForPrincipal(cfg, client)
			return
		}
		if _ssoadminListAccountsForProvisionedPermissionSet {
			ssoadmin_ListAccountsForProvisionedPermissionSet(cfg, client)
			return
		}
		if _ssoadminListApplicationAccessScopes {
			ssoadmin_ListApplicationAccessScopes(cfg, client)
			return
		}
		if _ssoadminListApplicationAssignments {
			ssoadmin_ListApplicationAssignments(cfg, client)
			return
		}
		if _ssoadminListApplicationAssignmentsForPrincipal {
			ssoadmin_ListApplicationAssignmentsForPrincipal(cfg, client)
			return
		}
		if _ssoadminListApplicationAuthenticationMethods {
			ssoadmin_ListApplicationAuthenticationMethods(cfg, client)
			return
		}
		if _ssoadminListApplicationGrants {
			ssoadmin_ListApplicationGrants(cfg, client)
			return
		}
		if _ssoadminListApplicationProviders {
			ssoadmin_ListApplicationProviders(cfg, client)
			return
		}
		if _ssoadminListApplications {
			ssoadmin_ListApplications(cfg, client)
			return
		}
		if _ssoadminListCustomerManagedPolicyReferencesInPermissionSet {
			ssoadmin_ListCustomerManagedPolicyReferencesInPermissionSet(cfg, client)
			return
		}
		if _ssoadminListInstances {
			ssoadmin_ListInstances(cfg, client)
			return
		}
		if _ssoadminListManagedPoliciesInPermissionSet {
			ssoadmin_ListManagedPoliciesInPermissionSet(cfg, client)
			return
		}
		if _ssoadminListPermissionSetProvisioningStatus {
			ssoadmin_ListPermissionSetProvisioningStatus(cfg, client)
			return
		}
		if _ssoadminListPermissionSets {
			ssoadmin_ListPermissionSets(cfg, client)
			return
		}
		if _ssoadminListPermissionSetsProvisionedToAccount {
			ssoadmin_ListPermissionSetsProvisionedToAccount(cfg, client)
			return
		}
		if _ssoadminListRegions {
			ssoadmin_ListRegions(cfg, client)
			return
		}
		if _ssoadminListTagsForResource {
			ssoadmin_ListTagsForResource(cfg, client)
			return
		}
		if _ssoadminListTrustedTokenIssuers {
			ssoadmin_ListTrustedTokenIssuers(cfg, client)
			return
		}
		if _ssoadminProvisionPermissionSet {
			ssoadmin_ProvisionPermissionSet(cfg, client)
			return
		}
		if _ssoadminPutApplicationAccessScope {
			ssoadmin_PutApplicationAccessScope(cfg, client)
			return
		}
		if _ssoadminPutApplicationAssignmentConfiguration {
			ssoadmin_PutApplicationAssignmentConfiguration(cfg, client)
			return
		}
		if _ssoadminPutApplicationAuthenticationMethod {
			ssoadmin_PutApplicationAuthenticationMethod(cfg, client)
			return
		}
		if _ssoadminPutApplicationGrant {
			ssoadmin_PutApplicationGrant(cfg, client)
			return
		}
		if _ssoadminPutApplicationSessionConfiguration {
			ssoadmin_PutApplicationSessionConfiguration(cfg, client)
			return
		}
		if _ssoadminPutInlinePolicyToPermissionSet {
			ssoadmin_PutInlinePolicyToPermissionSet(cfg, client)
			return
		}
		if _ssoadminPutPermissionsBoundaryToPermissionSet {
			ssoadmin_PutPermissionsBoundaryToPermissionSet(cfg, client)
			return
		}
		if _ssoadminRemoveRegion {
			ssoadmin_RemoveRegion(cfg, client)
			return
		}
		if _ssoadminTagResource {
			ssoadmin_TagResource(cfg, client)
			return
		}
		if _ssoadminUntagResource {
			ssoadmin_UntagResource(cfg, client)
			return
		}
		if _ssoadminUpdateApplication {
			ssoadmin_UpdateApplication(cfg, client)
			return
		}
		if _ssoadminUpdateInstance {
			ssoadmin_UpdateInstance(cfg, client)
			return
		}
		if _ssoadminUpdateInstanceAccessControlAttributeConfiguration {
			ssoadmin_UpdateInstanceAccessControlAttributeConfiguration(cfg, client)
			return
		}
		if _ssoadminUpdatePermissionSet {
			ssoadmin_UpdatePermissionSet(cfg, client)
			return
		}
		if _ssoadminUpdateTrustedTokenIssuer {
			ssoadmin_UpdateTrustedTokenIssuer(cfg, client)
			return
		}

	},
}

var (
	_ssoadminAddRegion                                             bool
	_ssoadminAttachCustomerManagedPolicyReferenceToPermissionSet   bool
	_ssoadminAttachManagedPolicyToPermissionSet                    bool
	_ssoadminCreateAccountAssignment                               bool
	_ssoadminCreateApplication                                     bool
	_ssoadminCreateApplicationAssignment                           bool
	_ssoadminCreateInstance                                        bool
	_ssoadminCreateInstanceAccessControlAttributeConfiguration     bool
	_ssoadminCreatePermissionSet                                   bool
	_ssoadminCreateTrustedTokenIssuer                              bool
	_ssoadminDeleteAccountAssignment                               bool
	_ssoadminDeleteApplication                                     bool
	_ssoadminDeleteApplicationAccessScope                          bool
	_ssoadminDeleteApplicationAssignment                           bool
	_ssoadminDeleteApplicationAuthenticationMethod                 bool
	_ssoadminDeleteApplicationGrant                                bool
	_ssoadminDeleteInlinePolicyFromPermissionSet                   bool
	_ssoadminDeleteInstance                                        bool
	_ssoadminDeleteInstanceAccessControlAttributeConfiguration     bool
	_ssoadminDeletePermissionSet                                   bool
	_ssoadminDeletePermissionsBoundaryFromPermissionSet            bool
	_ssoadminDeleteTrustedTokenIssuer                              bool
	_ssoadminDescribeAccountAssignmentCreationStatus               bool
	_ssoadminDescribeAccountAssignmentDeletionStatus               bool
	_ssoadminDescribeApplication                                   bool
	_ssoadminDescribeApplicationAssignment                         bool
	_ssoadminDescribeApplicationProvider                           bool
	_ssoadminDescribeInstance                                      bool
	_ssoadminDescribeInstanceAccessControlAttributeConfiguration   bool
	_ssoadminDescribePermissionSet                                 bool
	_ssoadminDescribePermissionSetProvisioningStatus               bool
	_ssoadminDescribeRegion                                        bool
	_ssoadminDescribeTrustedTokenIssuer                            bool
	_ssoadminDetachCustomerManagedPolicyReferenceFromPermissionSet bool
	_ssoadminDetachManagedPolicyFromPermissionSet                  bool
	_ssoadminGetApplicationAccessScope                             bool
	_ssoadminGetApplicationAssignmentConfiguration                 bool
	_ssoadminGetApplicationAuthenticationMethod                    bool
	_ssoadminGetApplicationGrant                                   bool
	_ssoadminGetApplicationSessionConfiguration                    bool
	_ssoadminGetInlinePolicyForPermissionSet                       bool
	_ssoadminGetPermissionsBoundaryForPermissionSet                bool
	_ssoadminListAccountAssignmentCreationStatus                   bool
	_ssoadminListAccountAssignmentDeletionStatus                   bool
	_ssoadminListAccountAssignments                                bool
	_ssoadminListAccountAssignmentsForPrincipal                    bool
	_ssoadminListAccountsForProvisionedPermissionSet               bool
	_ssoadminListApplicationAccessScopes                           bool
	_ssoadminListApplicationAssignments                            bool
	_ssoadminListApplicationAssignmentsForPrincipal                bool
	_ssoadminListApplicationAuthenticationMethods                  bool
	_ssoadminListApplicationGrants                                 bool
	_ssoadminListApplicationProviders                              bool
	_ssoadminListApplications                                      bool
	_ssoadminListCustomerManagedPolicyReferencesInPermissionSet    bool
	_ssoadminListInstances                                         bool
	_ssoadminListManagedPoliciesInPermissionSet                    bool
	_ssoadminListPermissionSetProvisioningStatus                   bool
	_ssoadminListPermissionSets                                    bool
	_ssoadminListPermissionSetsProvisionedToAccount                bool
	_ssoadminListRegions                                           bool
	_ssoadminListTagsForResource                                   bool
	_ssoadminListTrustedTokenIssuers                               bool
	_ssoadminProvisionPermissionSet                                bool
	_ssoadminPutApplicationAccessScope                             bool
	_ssoadminPutApplicationAssignmentConfiguration                 bool
	_ssoadminPutApplicationAuthenticationMethod                    bool
	_ssoadminPutApplicationGrant                                   bool
	_ssoadminPutApplicationSessionConfiguration                    bool
	_ssoadminPutInlinePolicyToPermissionSet                        bool
	_ssoadminPutPermissionsBoundaryToPermissionSet                 bool
	_ssoadminRemoveRegion                                          bool
	_ssoadminTagResource                                           bool
	_ssoadminUntagResource                                         bool
	_ssoadminUpdateApplication                                     bool
	_ssoadminUpdateInstance                                        bool
	_ssoadminUpdateInstanceAccessControlAttributeConfiguration     bool
	_ssoadminUpdatePermissionSet                                   bool
	_ssoadminUpdateTrustedTokenIssuer                              bool

	_ssoadminAccountAssignmentCreationRequestId          string
	_ssoadminAccountAssignmentDeletionRequestId          string
	_ssoadminAccountId                                   string
	_ssoadminApplicationArn                              string
	_ssoadminApplicationProviderArn                      string
	_ssoadminAssignmentRequired                          string
	_ssoadminAuthenticationMethod                        string
	_ssoadminAuthenticationMethodType                    string
	_ssoadminAuthorizedTargets                           []string
	_ssoadminClientToken                                 string
	_ssoadminCustomerManagedPolicyReference              string
	_ssoadminDescription                                 string
	_ssoadminEncryptionConfiguration                     string
	_ssoadminFilter                                      string
	_ssoadminGrant                                       string
	_ssoadminGrantType                                   string
	_ssoadminInlinePolicy                                string
	_ssoadminInstanceAccessControlAttributeConfiguration string
	_ssoadminInstanceArn                                 string
	_ssoadminManagedPolicyArn                            string
	_ssoadminMaxResults                                  string
	_ssoadminName                                        string
	_ssoadminNextToken                                   string
	_ssoadminPermissionSetArn                            string
	_ssoadminPermissionsBoundary                         string
	_ssoadminPortalOptions                               string
	_ssoadminPrincipalId                                 string
	_ssoadminPrincipalType                               string
	_ssoadminProvisionPermissionSetRequestId             string
	_ssoadminProvisioningStatus                          string
	_ssoadminRegionName                                  string
	_ssoadminRelayState                                  string
	_ssoadminResourceArn                                 string
	_ssoadminScope                                       string
	_ssoadminSessionDuration                             string
	_ssoadminStatus                                      string
	_ssoadminTagKeys                                     []string
	_ssoadminTags                                        string
	_ssoadminTargetId                                    string
	_ssoadminTargetType                                  string
	_ssoadminTrustedTokenIssuerArn                       string
	_ssoadminTrustedTokenIssuerConfiguration             string
	_ssoadminTrustedTokenIssuerType                      string
	_ssoadminUserBackgroundSessionApplicationStatus      string
)

// Adds a Region to an IAM Identity Center instance. This operation initiates an
// asynchronous workflow to replicate the IAM Identity Center instance to the
// target Region. The Region status is set to ADDING at first and changes to ACTIVE
// when the workflow completes.
//
// To use this operation, your IAM Identity Center instance and the target Region
// must meet the requirements described in the [IAM Identity Center User Guide].
//
// The following actions are related to AddRegion :
//
// [RemoveRegion]
//
// [DescribeRegion]
//
// [ListRegions]
//
// [IAM Identity Center User Guide]: https://docs.aws.amazon.com/singlesignon/latest/userguide/multi-region-iam-identity-center.html#multi-region-prerequisites
// [RemoveRegion]: https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_RemoveRegion.html
// [DescribeRegion]: https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_DescribeRegion.html
// [ListRegions]: https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_ListRegions.html
func ssoadmin_AddRegion(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.AddRegionInput{
		// InstanceArn: *string, // Required
		// RegionName: *string, // Required
	}

	if len(_ssoadminInstanceArn) > 0 {
		input.InstanceArn = aws.String(_ssoadminInstanceArn)
	}
	if len(_ssoadminRegionName) > 0 {
		input.RegionName = aws.String(_ssoadminRegionName)
	}

	if resp, err := client.AddRegion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Attaches the specified customer managed policy to the specified PermissionSet.
func ssoadmin_AttachCustomerManagedPolicyReferenceToPermissionSet(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.AttachCustomerManagedPolicyReferenceToPermissionSetInput{
		// CustomerManagedPolicyReference: *types.CustomerManagedPolicyReference, // Required
		// InstanceArn: *string, // Required
		// PermissionSetArn: *string, // Required
	}

	if len(_ssoadminCustomerManagedPolicyReference) > 0 {
		if err := assignInputField(input, "CustomerManagedPolicyReference", _ssoadminCustomerManagedPolicyReference); err != nil {
			log.Errorf("invalid --customer-managed-policy-reference: %s", err.Error())
			return
		}
	}
	if len(_ssoadminInstanceArn) > 0 {
		input.InstanceArn = aws.String(_ssoadminInstanceArn)
	}
	if len(_ssoadminPermissionSetArn) > 0 {
		input.PermissionSetArn = aws.String(_ssoadminPermissionSetArn)
	}

	if resp, err := client.AttachCustomerManagedPolicyReferenceToPermissionSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Attaches an Amazon Web Services managed policy ARN to a permission set.
// If the permission set is already referenced by one or more account assignments,
// you will need to call ProvisionPermissionSetafter this operation. Calling ProvisionPermissionSet
// applies the corresponding IAM policy updates to all assigned accounts.
func ssoadmin_AttachManagedPolicyToPermissionSet(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.AttachManagedPolicyToPermissionSetInput{
		// InstanceArn: *string, // Required
		// ManagedPolicyArn: *string, // Required
		// PermissionSetArn: *string, // Required
	}

	if len(_ssoadminInstanceArn) > 0 {
		input.InstanceArn = aws.String(_ssoadminInstanceArn)
	}
	if len(_ssoadminManagedPolicyArn) > 0 {
		input.ManagedPolicyArn = aws.String(_ssoadminManagedPolicyArn)
	}
	if len(_ssoadminPermissionSetArn) > 0 {
		input.PermissionSetArn = aws.String(_ssoadminPermissionSetArn)
	}

	if resp, err := client.AttachManagedPolicyToPermissionSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Assigns access to a principal for a specified Amazon Web Services account using
// a specified permission set.
//
// The term principal here refers to a user or group that is defined in IAM
// Identity Center.
//
// As part of a successful CreateAccountAssignment call, the specified permission
// set will automatically be provisioned to the account in the form of an IAM
// policy. That policy is attached to the IAM role created in IAM Identity Center.
// If the permission set is subsequently updated, the corresponding IAM policies
// attached to roles in your accounts will not be updated automatically. In this
// case, you must call ProvisionPermissionSetto make these updates.
//
// After a successful response, call DescribeAccountAssignmentCreationStatus to
// describe the status of an assignment creation request.
func ssoadmin_CreateAccountAssignment(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.CreateAccountAssignmentInput{
		// InstanceArn: *string, // Required
		// PermissionSetArn: *string, // Required
		// PrincipalId: *string, // Required
		// PrincipalType: types.PrincipalType, // Required
		// TargetId: *string, // Required
		// TargetType: types.TargetType, // Required
	}

	if len(_ssoadminInstanceArn) > 0 {
		input.InstanceArn = aws.String(_ssoadminInstanceArn)
	}
	if len(_ssoadminPermissionSetArn) > 0 {
		input.PermissionSetArn = aws.String(_ssoadminPermissionSetArn)
	}
	if len(_ssoadminPrincipalId) > 0 {
		input.PrincipalId = aws.String(_ssoadminPrincipalId)
	}
	if len(_ssoadminPrincipalType) > 0 {
		if err := assignInputField(input, "PrincipalType", _ssoadminPrincipalType); err != nil {
			log.Errorf("invalid --principal-type: %s", err.Error())
			return
		}
	}
	if len(_ssoadminTargetId) > 0 {
		input.TargetId = aws.String(_ssoadminTargetId)
	}
	if len(_ssoadminTargetType) > 0 {
		if err := assignInputField(input, "TargetType", _ssoadminTargetType); err != nil {
			log.Errorf("invalid --target-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAccountAssignment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an OAuth 2.0 customer managed application in IAM Identity Center for
// the given application provider.
//
// This API does not support creating SAML 2.0 customer managed applications or
// Amazon Web Services managed applications. To learn how to create an Amazon Web
// Services managed application, see the application user guide. You can create a
// SAML 2.0 customer managed application in the Amazon Web Services Management
// Console only. See [Setting up customer managed SAML 2.0 applications]. For more information on these application types, see [Amazon Web Services managed applications].
//
// [Setting up customer managed SAML 2.0 applications]: https://docs.aws.amazon.com/singlesignon/latest/userguide/customermanagedapps-saml2-setup.html
// [Amazon Web Services managed applications]: https://docs.aws.amazon.com/singlesignon/latest/userguide/awsapps.html
func ssoadmin_CreateApplication(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.CreateApplicationInput{
		// ApplicationProviderArn: *string, // Required
		// InstanceArn: *string, // Required
		// Name: *string, // Required
	}

	if len(_ssoadminApplicationProviderArn) > 0 {
		input.ApplicationProviderArn = aws.String(_ssoadminApplicationProviderArn)
	}
	if len(_ssoadminInstanceArn) > 0 {
		input.InstanceArn = aws.String(_ssoadminInstanceArn)
	}
	if len(_ssoadminName) > 0 {
		input.Name = aws.String(_ssoadminName)
	}
	if len(_ssoadminClientToken) > 0 {
		input.ClientToken = aws.String(_ssoadminClientToken)
	}
	if len(_ssoadminDescription) > 0 {
		input.Description = aws.String(_ssoadminDescription)
	}
	if len(_ssoadminPortalOptions) > 0 {
		if err := assignInputField(input, "PortalOptions", _ssoadminPortalOptions); err != nil {
			log.Errorf("invalid --portal-options: %s", err.Error())
			return
		}
	}
	if len(_ssoadminStatus) > 0 {
		if err := assignInputField(input, "Status", _ssoadminStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_ssoadminTags) > 0 {
		if err := assignInputField(input, "Tags", _ssoadminTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Grant application access to a user or group.
func ssoadmin_CreateApplicationAssignment(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.CreateApplicationAssignmentInput{
		// ApplicationArn: *string, // Required
		// PrincipalId: *string, // Required
		// PrincipalType: types.PrincipalType, // Required
	}

	if len(_ssoadminApplicationArn) > 0 {
		input.ApplicationArn = aws.String(_ssoadminApplicationArn)
	}
	if len(_ssoadminPrincipalId) > 0 {
		input.PrincipalId = aws.String(_ssoadminPrincipalId)
	}
	if len(_ssoadminPrincipalType) > 0 {
		if err := assignInputField(input, "PrincipalType", _ssoadminPrincipalType); err != nil {
			log.Errorf("invalid --principal-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateApplicationAssignment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an instance of IAM Identity Center for a standalone Amazon Web Services
// account that is not managed by Organizations or a member Amazon Web Services
// account in an organization. You can create only one instance per account and
// across all Amazon Web Services Regions.
//
// The CreateInstance request is rejected if the following apply:
//
// - The instance is created within the organization management account.
//
// - An instance already exists in the same account.
func ssoadmin_CreateInstance(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.CreateInstanceInput{}

	if len(_ssoadminClientToken) > 0 {
		input.ClientToken = aws.String(_ssoadminClientToken)
	}
	if len(_ssoadminName) > 0 {
		input.Name = aws.String(_ssoadminName)
	}
	if len(_ssoadminTags) > 0 {
		if err := assignInputField(input, "Tags", _ssoadminTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables the attributes-based access control (ABAC) feature for the specified
// IAM Identity Center instance. You can also specify new attributes to add to your
// ABAC configuration during the enabling process. For more information about ABAC,
// see Attribute-Based Access Controlin the IAM Identity Center User Guide.
//
// After a successful response, call
// DescribeInstanceAccessControlAttributeConfiguration to validate that
// InstanceAccessControlAttributeConfiguration was created.
func ssoadmin_CreateInstanceAccessControlAttributeConfiguration(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.CreateInstanceAccessControlAttributeConfigurationInput{
		// InstanceAccessControlAttributeConfiguration: *types.InstanceAccessControlAttributeConfiguration, // Required
		// InstanceArn: *string, // Required
	}

	if len(_ssoadminInstanceAccessControlAttributeConfiguration) > 0 {
		if err := assignInputField(input, "InstanceAccessControlAttributeConfiguration", _ssoadminInstanceAccessControlAttributeConfiguration); err != nil {
			log.Errorf("invalid --instance-access-control-attribute-configuration: %s", err.Error())
			return
		}
	}
	if len(_ssoadminInstanceArn) > 0 {
		input.InstanceArn = aws.String(_ssoadminInstanceArn)
	}

	if resp, err := client.CreateInstanceAccessControlAttributeConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a permission set within a specified IAM Identity Center instance.
// To grant users and groups access to Amazon Web Services account resources, use CreateAccountAssignment.
func ssoadmin_CreatePermissionSet(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.CreatePermissionSetInput{
		// InstanceArn: *string, // Required
		// Name: *string, // Required
	}

	if len(_ssoadminInstanceArn) > 0 {
		input.InstanceArn = aws.String(_ssoadminInstanceArn)
	}
	if len(_ssoadminName) > 0 {
		input.Name = aws.String(_ssoadminName)
	}
	if len(_ssoadminDescription) > 0 {
		input.Description = aws.String(_ssoadminDescription)
	}
	if len(_ssoadminRelayState) > 0 {
		input.RelayState = aws.String(_ssoadminRelayState)
	}
	if len(_ssoadminSessionDuration) > 0 {
		input.SessionDuration = aws.String(_ssoadminSessionDuration)
	}
	if len(_ssoadminTags) > 0 {
		if err := assignInputField(input, "Tags", _ssoadminTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreatePermissionSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a connection to a trusted token issuer in an instance of IAM Identity
// Center. A trusted token issuer enables trusted identity propagation to be used
// with applications that authenticate outside of Amazon Web Services.
//
// This trusted token issuer describes an external identity provider (IdP) that
// can generate claims or assertions in the form of access tokens for a user.
// Applications enabled for IAM Identity Center can use these tokens for
// authentication.
func ssoadmin_CreateTrustedTokenIssuer(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.CreateTrustedTokenIssuerInput{
		// InstanceArn: *string, // Required
		// Name: *string, // Required
		// TrustedTokenIssuerConfiguration: types.TrustedTokenIssuerConfiguration, // Required
		// TrustedTokenIssuerType: types.TrustedTokenIssuerType, // Required
	}

	if len(_ssoadminInstanceArn) > 0 {
		input.InstanceArn = aws.String(_ssoadminInstanceArn)
	}
	if len(_ssoadminName) > 0 {
		input.Name = aws.String(_ssoadminName)
	}
	if len(_ssoadminTrustedTokenIssuerConfiguration) > 0 {
		if err := assignInputField(input, "TrustedTokenIssuerConfiguration", _ssoadminTrustedTokenIssuerConfiguration); err != nil {
			log.Errorf("invalid --trusted-token-issuer-configuration: %s", err.Error())
			return
		}
	}
	if len(_ssoadminTrustedTokenIssuerType) > 0 {
		if err := assignInputField(input, "TrustedTokenIssuerType", _ssoadminTrustedTokenIssuerType); err != nil {
			log.Errorf("invalid --trusted-token-issuer-type: %s", err.Error())
			return
		}
	}
	if len(_ssoadminClientToken) > 0 {
		input.ClientToken = aws.String(_ssoadminClientToken)
	}
	if len(_ssoadminTags) > 0 {
		if err := assignInputField(input, "Tags", _ssoadminTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateTrustedTokenIssuer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a principal's access from a specified Amazon Web Services account using
// a specified permission set.
//
// After a successful response, call DescribeAccountAssignmentDeletionStatus to
// describe the status of an assignment deletion request.
func ssoadmin_DeleteAccountAssignment(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.DeleteAccountAssignmentInput{
		// InstanceArn: *string, // Required
		// PermissionSetArn: *string, // Required
		// PrincipalId: *string, // Required
		// PrincipalType: types.PrincipalType, // Required
		// TargetId: *string, // Required
		// TargetType: types.TargetType, // Required
	}

	if len(_ssoadminInstanceArn) > 0 {
		input.InstanceArn = aws.String(_ssoadminInstanceArn)
	}
	if len(_ssoadminPermissionSetArn) > 0 {
		input.PermissionSetArn = aws.String(_ssoadminPermissionSetArn)
	}
	if len(_ssoadminPrincipalId) > 0 {
		input.PrincipalId = aws.String(_ssoadminPrincipalId)
	}
	if len(_ssoadminPrincipalType) > 0 {
		if err := assignInputField(input, "PrincipalType", _ssoadminPrincipalType); err != nil {
			log.Errorf("invalid --principal-type: %s", err.Error())
			return
		}
	}
	if len(_ssoadminTargetId) > 0 {
		input.TargetId = aws.String(_ssoadminTargetId)
	}
	if len(_ssoadminTargetType) > 0 {
		if err := assignInputField(input, "TargetType", _ssoadminTargetType); err != nil {
			log.Errorf("invalid --target-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteAccountAssignment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the association with the application. The connected service resource
// still exists.
func ssoadmin_DeleteApplication(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.DeleteApplicationInput{
		// ApplicationArn: *string, // Required
	}

	if len(_ssoadminApplicationArn) > 0 {
		input.ApplicationArn = aws.String(_ssoadminApplicationArn)
	}

	if resp, err := client.DeleteApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an IAM Identity Center access scope from an application.
func ssoadmin_DeleteApplicationAccessScope(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.DeleteApplicationAccessScopeInput{
		// ApplicationArn: *string, // Required
		// Scope: *string, // Required
	}

	if len(_ssoadminApplicationArn) > 0 {
		input.ApplicationArn = aws.String(_ssoadminApplicationArn)
	}
	if len(_ssoadminScope) > 0 {
		input.Scope = aws.String(_ssoadminScope)
	}

	if resp, err := client.DeleteApplicationAccessScope(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Revoke application access to an application by deleting application assignments
// for a user or group.
func ssoadmin_DeleteApplicationAssignment(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.DeleteApplicationAssignmentInput{
		// ApplicationArn: *string, // Required
		// PrincipalId: *string, // Required
		// PrincipalType: types.PrincipalType, // Required
	}

	if len(_ssoadminApplicationArn) > 0 {
		input.ApplicationArn = aws.String(_ssoadminApplicationArn)
	}
	if len(_ssoadminPrincipalId) > 0 {
		input.PrincipalId = aws.String(_ssoadminPrincipalId)
	}
	if len(_ssoadminPrincipalType) > 0 {
		if err := assignInputField(input, "PrincipalType", _ssoadminPrincipalType); err != nil {
			log.Errorf("invalid --principal-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteApplicationAssignment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an authentication method from an application.
func ssoadmin_DeleteApplicationAuthenticationMethod(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.DeleteApplicationAuthenticationMethodInput{
		// ApplicationArn: *string, // Required
		// AuthenticationMethodType: types.AuthenticationMethodType, // Required
	}

	if len(_ssoadminApplicationArn) > 0 {
		input.ApplicationArn = aws.String(_ssoadminApplicationArn)
	}
	if len(_ssoadminAuthenticationMethodType) > 0 {
		if err := assignInputField(input, "AuthenticationMethodType", _ssoadminAuthenticationMethodType); err != nil {
			log.Errorf("invalid --authentication-method-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteApplicationAuthenticationMethod(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a grant from an application.
func ssoadmin_DeleteApplicationGrant(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.DeleteApplicationGrantInput{
		// ApplicationArn: *string, // Required
		// GrantType: types.GrantType, // Required
	}

	if len(_ssoadminApplicationArn) > 0 {
		input.ApplicationArn = aws.String(_ssoadminApplicationArn)
	}
	if len(_ssoadminGrantType) > 0 {
		if err := assignInputField(input, "GrantType", _ssoadminGrantType); err != nil {
			log.Errorf("invalid --grant-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteApplicationGrant(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the inline policy from a specified permission set.
func ssoadmin_DeleteInlinePolicyFromPermissionSet(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.DeleteInlinePolicyFromPermissionSetInput{
		// InstanceArn: *string, // Required
		// PermissionSetArn: *string, // Required
	}

	if len(_ssoadminInstanceArn) > 0 {
		input.InstanceArn = aws.String(_ssoadminInstanceArn)
	}
	if len(_ssoadminPermissionSetArn) > 0 {
		input.PermissionSetArn = aws.String(_ssoadminPermissionSetArn)
	}

	if resp, err := client.DeleteInlinePolicyFromPermissionSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the instance of IAM Identity Center. Only the account that owns the
// instance can call this API. Neither the delegated administrator nor member
// account can delete the organization instance, but those roles can delete their
// own instance.
func ssoadmin_DeleteInstance(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.DeleteInstanceInput{
		// InstanceArn: *string, // Required
	}

	if len(_ssoadminInstanceArn) > 0 {
		input.InstanceArn = aws.String(_ssoadminInstanceArn)
	}

	if resp, err := client.DeleteInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disables the attributes-based access control (ABAC) feature for the specified
// IAM Identity Center instance and deletes all of the attribute mappings that have
// been configured. Once deleted, any attributes that are received from an identity
// source and any custom attributes you have previously configured will not be
// passed. For more information about ABAC, see Attribute-Based Access Controlin the IAM Identity Center User
// Guide.
func ssoadmin_DeleteInstanceAccessControlAttributeConfiguration(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.DeleteInstanceAccessControlAttributeConfigurationInput{
		// InstanceArn: *string, // Required
	}

	if len(_ssoadminInstanceArn) > 0 {
		input.InstanceArn = aws.String(_ssoadminInstanceArn)
	}

	if resp, err := client.DeleteInstanceAccessControlAttributeConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified permission set.
func ssoadmin_DeletePermissionSet(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.DeletePermissionSetInput{
		// InstanceArn: *string, // Required
		// PermissionSetArn: *string, // Required
	}

	if len(_ssoadminInstanceArn) > 0 {
		input.InstanceArn = aws.String(_ssoadminInstanceArn)
	}
	if len(_ssoadminPermissionSetArn) > 0 {
		input.PermissionSetArn = aws.String(_ssoadminPermissionSetArn)
	}

	if resp, err := client.DeletePermissionSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the permissions boundary from a specified PermissionSet.
func ssoadmin_DeletePermissionsBoundaryFromPermissionSet(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.DeletePermissionsBoundaryFromPermissionSetInput{
		// InstanceArn: *string, // Required
		// PermissionSetArn: *string, // Required
	}

	if len(_ssoadminInstanceArn) > 0 {
		input.InstanceArn = aws.String(_ssoadminInstanceArn)
	}
	if len(_ssoadminPermissionSetArn) > 0 {
		input.PermissionSetArn = aws.String(_ssoadminPermissionSetArn)
	}

	if resp, err := client.DeletePermissionsBoundaryFromPermissionSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a trusted token issuer configuration from an instance of IAM Identity
// Center.
//
// Deleting this trusted token issuer configuration will cause users to lose
// access to any applications that are configured to use the trusted token issuer.
func ssoadmin_DeleteTrustedTokenIssuer(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.DeleteTrustedTokenIssuerInput{
		// TrustedTokenIssuerArn: *string, // Required
	}

	if len(_ssoadminTrustedTokenIssuerArn) > 0 {
		input.TrustedTokenIssuerArn = aws.String(_ssoadminTrustedTokenIssuerArn)
	}

	if resp, err := client.DeleteTrustedTokenIssuer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the status of the assignment creation request.
func ssoadmin_DescribeAccountAssignmentCreationStatus(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.DescribeAccountAssignmentCreationStatusInput{
		// AccountAssignmentCreationRequestId: *string, // Required
		// InstanceArn: *string, // Required
	}

	if len(_ssoadminAccountAssignmentCreationRequestId) > 0 {
		input.AccountAssignmentCreationRequestId = aws.String(_ssoadminAccountAssignmentCreationRequestId)
	}
	if len(_ssoadminInstanceArn) > 0 {
		input.InstanceArn = aws.String(_ssoadminInstanceArn)
	}

	if resp, err := client.DescribeAccountAssignmentCreationStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the status of the assignment deletion request.
func ssoadmin_DescribeAccountAssignmentDeletionStatus(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.DescribeAccountAssignmentDeletionStatusInput{
		// AccountAssignmentDeletionRequestId: *string, // Required
		// InstanceArn: *string, // Required
	}

	if len(_ssoadminAccountAssignmentDeletionRequestId) > 0 {
		input.AccountAssignmentDeletionRequestId = aws.String(_ssoadminAccountAssignmentDeletionRequestId)
	}
	if len(_ssoadminInstanceArn) > 0 {
		input.InstanceArn = aws.String(_ssoadminInstanceArn)
	}

	if resp, err := client.DescribeAccountAssignmentDeletionStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the details of an application associated with an instance of IAM
// Identity Center.
func ssoadmin_DescribeApplication(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.DescribeApplicationInput{
		// ApplicationArn: *string, // Required
	}

	if len(_ssoadminApplicationArn) > 0 {
		input.ApplicationArn = aws.String(_ssoadminApplicationArn)
	}

	if resp, err := client.DescribeApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a direct assignment of a user or group to an application. If the user
// doesn’t have a direct assignment to the application, the user may still have
// access to the application through a group. Therefore, don’t use this API to test
// access to an application for a user. Instead use ListApplicationAssignmentsForPrincipal.
func ssoadmin_DescribeApplicationAssignment(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.DescribeApplicationAssignmentInput{
		// ApplicationArn: *string, // Required
		// PrincipalId: *string, // Required
		// PrincipalType: types.PrincipalType, // Required
	}

	if len(_ssoadminApplicationArn) > 0 {
		input.ApplicationArn = aws.String(_ssoadminApplicationArn)
	}
	if len(_ssoadminPrincipalId) > 0 {
		input.PrincipalId = aws.String(_ssoadminPrincipalId)
	}
	if len(_ssoadminPrincipalType) > 0 {
		if err := assignInputField(input, "PrincipalType", _ssoadminPrincipalType); err != nil {
			log.Errorf("invalid --principal-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeApplicationAssignment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves details about a provider that can be used to connect an Amazon Web
// Services managed application or customer managed application to IAM Identity
// Center.
func ssoadmin_DescribeApplicationProvider(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.DescribeApplicationProviderInput{
		// ApplicationProviderArn: *string, // Required
	}

	if len(_ssoadminApplicationProviderArn) > 0 {
		input.ApplicationProviderArn = aws.String(_ssoadminApplicationProviderArn)
	}

	if resp, err := client.DescribeApplicationProvider(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the details of an instance of IAM Identity Center. The status can be
// one of the following:
//
// - CREATE_IN_PROGRESS - The instance is in the process of being created. When
// the instance is ready for use, DescribeInstance returns the status of ACTIVE .
// While the instance is in the CREATE_IN_PROGRESS state, you can call only
// DescribeInstance and DeleteInstance operations.
//
// - DELETE_IN_PROGRESS - The instance is being deleted. Returns
// AccessDeniedException after the delete operation completes.
//
// - ACTIVE - The instance is active.
func ssoadmin_DescribeInstance(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.DescribeInstanceInput{
		// InstanceArn: *string, // Required
	}

	if len(_ssoadminInstanceArn) > 0 {
		input.InstanceArn = aws.String(_ssoadminInstanceArn)
	}

	if resp, err := client.DescribeInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the list of IAM Identity Center identity store attributes that have
// been configured to work with attributes-based access control (ABAC) for the
// specified IAM Identity Center instance. This will not return attributes
// configured and sent by an external identity provider. For more information about
// ABAC, see Attribute-Based Access Controlin the IAM Identity Center User Guide.
func ssoadmin_DescribeInstanceAccessControlAttributeConfiguration(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.DescribeInstanceAccessControlAttributeConfigurationInput{
		// InstanceArn: *string, // Required
	}

	if len(_ssoadminInstanceArn) > 0 {
		input.InstanceArn = aws.String(_ssoadminInstanceArn)
	}

	if resp, err := client.DescribeInstanceAccessControlAttributeConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the details of the permission set.
func ssoadmin_DescribePermissionSet(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.DescribePermissionSetInput{
		// InstanceArn: *string, // Required
		// PermissionSetArn: *string, // Required
	}

	if len(_ssoadminInstanceArn) > 0 {
		input.InstanceArn = aws.String(_ssoadminInstanceArn)
	}
	if len(_ssoadminPermissionSetArn) > 0 {
		input.PermissionSetArn = aws.String(_ssoadminPermissionSetArn)
	}

	if resp, err := client.DescribePermissionSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the status for the given permission set provisioning request.
func ssoadmin_DescribePermissionSetProvisioningStatus(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.DescribePermissionSetProvisioningStatusInput{
		// InstanceArn: *string, // Required
		// ProvisionPermissionSetRequestId: *string, // Required
	}

	if len(_ssoadminInstanceArn) > 0 {
		input.InstanceArn = aws.String(_ssoadminInstanceArn)
	}
	if len(_ssoadminProvisionPermissionSetRequestId) > 0 {
		input.ProvisionPermissionSetRequestId = aws.String(_ssoadminProvisionPermissionSetRequestId)
	}

	if resp, err := client.DescribePermissionSetProvisioningStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves details about a specific Region enabled in an IAM Identity Center
// instance. Details include the Region name, current status (ACTIVE, ADDING, or
// REMOVING), the date when the Region was added, and whether it is the primary
// Region. The request must be made from one of the enabled Regions of the IAM
// Identity Center instance.
//
// The following actions are related to DescribeRegion :
//
// [AddRegion]
//
// [RemoveRegion]
//
// [ListRegions]
//
// [AddRegion]: https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_AddRegion.html
// [RemoveRegion]: https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_RemoveRegion.html
// [ListRegions]: https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_ListRegions.html
func ssoadmin_DescribeRegion(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.DescribeRegionInput{
		// InstanceArn: *string, // Required
		// RegionName: *string, // Required
	}

	if len(_ssoadminInstanceArn) > 0 {
		input.InstanceArn = aws.String(_ssoadminInstanceArn)
	}
	if len(_ssoadminRegionName) > 0 {
		input.RegionName = aws.String(_ssoadminRegionName)
	}

	if resp, err := client.DescribeRegion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves details about a trusted token issuer configuration stored in an
// instance of IAM Identity Center. Details include the name of the trusted token
// issuer, the issuer URL, and the path of the source attribute and the destination
// attribute for a trusted token issuer configuration.
func ssoadmin_DescribeTrustedTokenIssuer(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.DescribeTrustedTokenIssuerInput{
		// TrustedTokenIssuerArn: *string, // Required
	}

	if len(_ssoadminTrustedTokenIssuerArn) > 0 {
		input.TrustedTokenIssuerArn = aws.String(_ssoadminTrustedTokenIssuerArn)
	}

	if resp, err := client.DescribeTrustedTokenIssuer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Detaches the specified customer managed policy from the specified PermissionSet.
func ssoadmin_DetachCustomerManagedPolicyReferenceFromPermissionSet(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.DetachCustomerManagedPolicyReferenceFromPermissionSetInput{
		// CustomerManagedPolicyReference: *types.CustomerManagedPolicyReference, // Required
		// InstanceArn: *string, // Required
		// PermissionSetArn: *string, // Required
	}

	if len(_ssoadminCustomerManagedPolicyReference) > 0 {
		if err := assignInputField(input, "CustomerManagedPolicyReference", _ssoadminCustomerManagedPolicyReference); err != nil {
			log.Errorf("invalid --customer-managed-policy-reference: %s", err.Error())
			return
		}
	}
	if len(_ssoadminInstanceArn) > 0 {
		input.InstanceArn = aws.String(_ssoadminInstanceArn)
	}
	if len(_ssoadminPermissionSetArn) > 0 {
		input.PermissionSetArn = aws.String(_ssoadminPermissionSetArn)
	}

	if resp, err := client.DetachCustomerManagedPolicyReferenceFromPermissionSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Detaches the attached Amazon Web Services managed policy ARN from the specified
// permission set.
func ssoadmin_DetachManagedPolicyFromPermissionSet(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.DetachManagedPolicyFromPermissionSetInput{
		// InstanceArn: *string, // Required
		// ManagedPolicyArn: *string, // Required
		// PermissionSetArn: *string, // Required
	}

	if len(_ssoadminInstanceArn) > 0 {
		input.InstanceArn = aws.String(_ssoadminInstanceArn)
	}
	if len(_ssoadminManagedPolicyArn) > 0 {
		input.ManagedPolicyArn = aws.String(_ssoadminManagedPolicyArn)
	}
	if len(_ssoadminPermissionSetArn) > 0 {
		input.PermissionSetArn = aws.String(_ssoadminPermissionSetArn)
	}

	if resp, err := client.DetachManagedPolicyFromPermissionSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the authorized targets for an IAM Identity Center access scope for an
// application.
func ssoadmin_GetApplicationAccessScope(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.GetApplicationAccessScopeInput{
		// ApplicationArn: *string, // Required
		// Scope: *string, // Required
	}

	if len(_ssoadminApplicationArn) > 0 {
		input.ApplicationArn = aws.String(_ssoadminApplicationArn)
	}
	if len(_ssoadminScope) > 0 {
		input.Scope = aws.String(_ssoadminScope)
	}

	if resp, err := client.GetApplicationAccessScope(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the configuration of PutApplicationAssignmentConfiguration.
func ssoadmin_GetApplicationAssignmentConfiguration(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.GetApplicationAssignmentConfigurationInput{
		// ApplicationArn: *string, // Required
	}

	if len(_ssoadminApplicationArn) > 0 {
		input.ApplicationArn = aws.String(_ssoadminApplicationArn)
	}

	if resp, err := client.GetApplicationAssignmentConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves details about an authentication method used by an application.
func ssoadmin_GetApplicationAuthenticationMethod(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.GetApplicationAuthenticationMethodInput{
		// ApplicationArn: *string, // Required
		// AuthenticationMethodType: types.AuthenticationMethodType, // Required
	}

	if len(_ssoadminApplicationArn) > 0 {
		input.ApplicationArn = aws.String(_ssoadminApplicationArn)
	}
	if len(_ssoadminAuthenticationMethodType) > 0 {
		if err := assignInputField(input, "AuthenticationMethodType", _ssoadminAuthenticationMethodType); err != nil {
			log.Errorf("invalid --authentication-method-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetApplicationAuthenticationMethod(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves details about an application grant.
func ssoadmin_GetApplicationGrant(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.GetApplicationGrantInput{
		// ApplicationArn: *string, // Required
		// GrantType: types.GrantType, // Required
	}

	if len(_ssoadminApplicationArn) > 0 {
		input.ApplicationArn = aws.String(_ssoadminApplicationArn)
	}
	if len(_ssoadminGrantType) > 0 {
		if err := assignInputField(input, "GrantType", _ssoadminGrantType); err != nil {
			log.Errorf("invalid --grant-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetApplicationGrant(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the session configuration for an application in IAM Identity Center.
// The session configuration determines how users can access an application. This
// includes whether user background sessions are enabled. User background sessions
// allow users to start a job on a supported Amazon Web Services managed
// application without having to remain signed in to an active session while the
// job runs.
func ssoadmin_GetApplicationSessionConfiguration(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.GetApplicationSessionConfigurationInput{
		// ApplicationArn: *string, // Required
	}

	if len(_ssoadminApplicationArn) > 0 {
		input.ApplicationArn = aws.String(_ssoadminApplicationArn)
	}

	if resp, err := client.GetApplicationSessionConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Obtains the inline policy assigned to the permission set.
func ssoadmin_GetInlinePolicyForPermissionSet(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.GetInlinePolicyForPermissionSetInput{
		// InstanceArn: *string, // Required
		// PermissionSetArn: *string, // Required
	}

	if len(_ssoadminInstanceArn) > 0 {
		input.InstanceArn = aws.String(_ssoadminInstanceArn)
	}
	if len(_ssoadminPermissionSetArn) > 0 {
		input.PermissionSetArn = aws.String(_ssoadminPermissionSetArn)
	}

	if resp, err := client.GetInlinePolicyForPermissionSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Obtains the permissions boundary for a specified PermissionSet.
func ssoadmin_GetPermissionsBoundaryForPermissionSet(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.GetPermissionsBoundaryForPermissionSetInput{
		// InstanceArn: *string, // Required
		// PermissionSetArn: *string, // Required
	}

	if len(_ssoadminInstanceArn) > 0 {
		input.InstanceArn = aws.String(_ssoadminInstanceArn)
	}
	if len(_ssoadminPermissionSetArn) > 0 {
		input.PermissionSetArn = aws.String(_ssoadminPermissionSetArn)
	}

	if resp, err := client.GetPermissionsBoundaryForPermissionSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the status of the Amazon Web Services account assignment creation
// requests for a specified IAM Identity Center instance.
func ssoadmin_ListAccountAssignmentCreationStatus(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.ListAccountAssignmentCreationStatusInput{
		// InstanceArn: *string, // Required
	}

	if len(_ssoadminInstanceArn) > 0 {
		input.InstanceArn = aws.String(_ssoadminInstanceArn)
	}
	if len(_ssoadminFilter) > 0 {
		if err := assignInputField(input, "Filter", _ssoadminFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_ssoadminMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssoadminMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssoadminNextToken) > 0 {
		input.NextToken = aws.String(_ssoadminNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAccountAssignmentCreationStatus(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssoadmin.ListAccountAssignmentCreationStatusOutput
	p := ssoadmin.NewListAccountAssignmentCreationStatusPaginator(client, input)
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

// Lists the status of the Amazon Web Services account assignment deletion
// requests for a specified IAM Identity Center instance.
func ssoadmin_ListAccountAssignmentDeletionStatus(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.ListAccountAssignmentDeletionStatusInput{
		// InstanceArn: *string, // Required
	}

	if len(_ssoadminInstanceArn) > 0 {
		input.InstanceArn = aws.String(_ssoadminInstanceArn)
	}
	if len(_ssoadminFilter) > 0 {
		if err := assignInputField(input, "Filter", _ssoadminFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_ssoadminMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssoadminMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssoadminNextToken) > 0 {
		input.NextToken = aws.String(_ssoadminNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAccountAssignmentDeletionStatus(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssoadmin.ListAccountAssignmentDeletionStatusOutput
	p := ssoadmin.NewListAccountAssignmentDeletionStatusPaginator(client, input)
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

// Lists the assignee of the specified Amazon Web Services account with the
// specified permission set.
func ssoadmin_ListAccountAssignments(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.ListAccountAssignmentsInput{
		// AccountId: *string, // Required
		// InstanceArn: *string, // Required
		// PermissionSetArn: *string, // Required
	}

	if len(_ssoadminAccountId) > 0 {
		input.AccountId = aws.String(_ssoadminAccountId)
	}
	if len(_ssoadminInstanceArn) > 0 {
		input.InstanceArn = aws.String(_ssoadminInstanceArn)
	}
	if len(_ssoadminPermissionSetArn) > 0 {
		input.PermissionSetArn = aws.String(_ssoadminPermissionSetArn)
	}
	if len(_ssoadminMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssoadminMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssoadminNextToken) > 0 {
		input.NextToken = aws.String(_ssoadminNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAccountAssignments(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssoadmin.ListAccountAssignmentsOutput
	p := ssoadmin.NewListAccountAssignmentsPaginator(client, input)
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

// Retrieves a list of the IAM Identity Center associated Amazon Web Services
// accounts that the principal has access to. This action must be called from the
// management account containing your organization instance of IAM Identity Center.
// This action is not valid for account instances of IAM Identity Center.
func ssoadmin_ListAccountAssignmentsForPrincipal(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.ListAccountAssignmentsForPrincipalInput{
		// InstanceArn: *string, // Required
		// PrincipalId: *string, // Required
		// PrincipalType: types.PrincipalType, // Required
	}

	if len(_ssoadminInstanceArn) > 0 {
		input.InstanceArn = aws.String(_ssoadminInstanceArn)
	}
	if len(_ssoadminPrincipalId) > 0 {
		input.PrincipalId = aws.String(_ssoadminPrincipalId)
	}
	if len(_ssoadminPrincipalType) > 0 {
		if err := assignInputField(input, "PrincipalType", _ssoadminPrincipalType); err != nil {
			log.Errorf("invalid --principal-type: %s", err.Error())
			return
		}
	}
	if len(_ssoadminFilter) > 0 {
		if err := assignInputField(input, "Filter", _ssoadminFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_ssoadminMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssoadminMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssoadminNextToken) > 0 {
		input.NextToken = aws.String(_ssoadminNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAccountAssignmentsForPrincipal(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssoadmin.ListAccountAssignmentsForPrincipalOutput
	p := ssoadmin.NewListAccountAssignmentsForPrincipalPaginator(client, input)
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

// Lists all the Amazon Web Services accounts where the specified permission set
// is provisioned.
func ssoadmin_ListAccountsForProvisionedPermissionSet(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.ListAccountsForProvisionedPermissionSetInput{
		// InstanceArn: *string, // Required
		// PermissionSetArn: *string, // Required
	}

	if len(_ssoadminInstanceArn) > 0 {
		input.InstanceArn = aws.String(_ssoadminInstanceArn)
	}
	if len(_ssoadminPermissionSetArn) > 0 {
		input.PermissionSetArn = aws.String(_ssoadminPermissionSetArn)
	}
	if len(_ssoadminMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssoadminMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssoadminNextToken) > 0 {
		input.NextToken = aws.String(_ssoadminNextToken)
	}
	if len(_ssoadminProvisioningStatus) > 0 {
		if err := assignInputField(input, "ProvisioningStatus", _ssoadminProvisioningStatus); err != nil {
			log.Errorf("invalid --provisioning-status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListAccountsForProvisionedPermissionSet(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssoadmin.ListAccountsForProvisionedPermissionSetOutput
	p := ssoadmin.NewListAccountsForProvisionedPermissionSetPaginator(client, input)
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

// Lists the access scopes and authorized targets associated with an application.
func ssoadmin_ListApplicationAccessScopes(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.ListApplicationAccessScopesInput{
		// ApplicationArn: *string, // Required
	}

	if len(_ssoadminApplicationArn) > 0 {
		input.ApplicationArn = aws.String(_ssoadminApplicationArn)
	}
	if len(_ssoadminMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssoadminMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssoadminNextToken) > 0 {
		input.NextToken = aws.String(_ssoadminNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListApplicationAccessScopes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssoadmin.ListApplicationAccessScopesOutput
	p := ssoadmin.NewListApplicationAccessScopesPaginator(client, input)
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

// Lists Amazon Web Services account users that are assigned to an application.
func ssoadmin_ListApplicationAssignments(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.ListApplicationAssignmentsInput{
		// ApplicationArn: *string, // Required
	}

	if len(_ssoadminApplicationArn) > 0 {
		input.ApplicationArn = aws.String(_ssoadminApplicationArn)
	}
	if len(_ssoadminMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssoadminMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssoadminNextToken) > 0 {
		input.NextToken = aws.String(_ssoadminNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListApplicationAssignments(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssoadmin.ListApplicationAssignmentsOutput
	p := ssoadmin.NewListApplicationAssignmentsPaginator(client, input)
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

// Lists the applications to which a specified principal is assigned. You must
// provide a filter when calling this action from a member account against your
// organization instance of IAM Identity Center. A filter is not required when
// called from the management account against an organization instance of IAM
// Identity Center, or from a member account against an account instance of IAM
// Identity Center in the same account.
func ssoadmin_ListApplicationAssignmentsForPrincipal(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.ListApplicationAssignmentsForPrincipalInput{
		// InstanceArn: *string, // Required
		// PrincipalId: *string, // Required
		// PrincipalType: types.PrincipalType, // Required
	}

	if len(_ssoadminInstanceArn) > 0 {
		input.InstanceArn = aws.String(_ssoadminInstanceArn)
	}
	if len(_ssoadminPrincipalId) > 0 {
		input.PrincipalId = aws.String(_ssoadminPrincipalId)
	}
	if len(_ssoadminPrincipalType) > 0 {
		if err := assignInputField(input, "PrincipalType", _ssoadminPrincipalType); err != nil {
			log.Errorf("invalid --principal-type: %s", err.Error())
			return
		}
	}
	if len(_ssoadminFilter) > 0 {
		if err := assignInputField(input, "Filter", _ssoadminFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_ssoadminMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssoadminMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssoadminNextToken) > 0 {
		input.NextToken = aws.String(_ssoadminNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListApplicationAssignmentsForPrincipal(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssoadmin.ListApplicationAssignmentsForPrincipalOutput
	p := ssoadmin.NewListApplicationAssignmentsForPrincipalPaginator(client, input)
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

// Lists all of the authentication methods supported by the specified application.
func ssoadmin_ListApplicationAuthenticationMethods(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.ListApplicationAuthenticationMethodsInput{
		// ApplicationArn: *string, // Required
	}

	if len(_ssoadminApplicationArn) > 0 {
		input.ApplicationArn = aws.String(_ssoadminApplicationArn)
	}
	if len(_ssoadminNextToken) > 0 {
		input.NextToken = aws.String(_ssoadminNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListApplicationAuthenticationMethods(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssoadmin.ListApplicationAuthenticationMethodsOutput
	p := ssoadmin.NewListApplicationAuthenticationMethodsPaginator(client, input)
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

// List the grants associated with an application.
func ssoadmin_ListApplicationGrants(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.ListApplicationGrantsInput{
		// ApplicationArn: *string, // Required
	}

	if len(_ssoadminApplicationArn) > 0 {
		input.ApplicationArn = aws.String(_ssoadminApplicationArn)
	}
	if len(_ssoadminNextToken) > 0 {
		input.NextToken = aws.String(_ssoadminNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListApplicationGrants(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssoadmin.ListApplicationGrantsOutput
	p := ssoadmin.NewListApplicationGrantsPaginator(client, input)
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

// Lists the application providers configured in the IAM Identity Center identity
// store.
func ssoadmin_ListApplicationProviders(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.ListApplicationProvidersInput{}

	if len(_ssoadminMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssoadminMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssoadminNextToken) > 0 {
		input.NextToken = aws.String(_ssoadminNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListApplicationProviders(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssoadmin.ListApplicationProvidersOutput
	p := ssoadmin.NewListApplicationProvidersPaginator(client, input)
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

// Lists all applications associated with the instance of IAM Identity Center.
// When listing applications for an organization instance in the management
// account, member accounts must use the applicationAccount parameter to filter
// the list to only applications created from that account. When listing
// applications for an account instance in the same member account, a filter is not
// required.
func ssoadmin_ListApplications(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.ListApplicationsInput{
		// InstanceArn: *string, // Required
	}

	if len(_ssoadminInstanceArn) > 0 {
		input.InstanceArn = aws.String(_ssoadminInstanceArn)
	}
	if len(_ssoadminFilter) > 0 {
		if err := assignInputField(input, "Filter", _ssoadminFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_ssoadminMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssoadminMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssoadminNextToken) > 0 {
		input.NextToken = aws.String(_ssoadminNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListApplications(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssoadmin.ListApplicationsOutput
	p := ssoadmin.NewListApplicationsPaginator(client, input)
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

// Lists all customer managed policies attached to a specified PermissionSet.
func ssoadmin_ListCustomerManagedPolicyReferencesInPermissionSet(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.ListCustomerManagedPolicyReferencesInPermissionSetInput{
		// InstanceArn: *string, // Required
		// PermissionSetArn: *string, // Required
	}

	if len(_ssoadminInstanceArn) > 0 {
		input.InstanceArn = aws.String(_ssoadminInstanceArn)
	}
	if len(_ssoadminPermissionSetArn) > 0 {
		input.PermissionSetArn = aws.String(_ssoadminPermissionSetArn)
	}
	if len(_ssoadminMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssoadminMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssoadminNextToken) > 0 {
		input.NextToken = aws.String(_ssoadminNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCustomerManagedPolicyReferencesInPermissionSet(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssoadmin.ListCustomerManagedPolicyReferencesInPermissionSetOutput
	p := ssoadmin.NewListCustomerManagedPolicyReferencesInPermissionSetPaginator(client, input)
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

// Lists the details of the organization and account instances of IAM Identity
// Center that were created in or visible to the account calling this API.
func ssoadmin_ListInstances(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.ListInstancesInput{}

	if len(_ssoadminMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssoadminMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssoadminNextToken) > 0 {
		input.NextToken = aws.String(_ssoadminNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListInstances(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssoadmin.ListInstancesOutput
	p := ssoadmin.NewListInstancesPaginator(client, input)
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

// Lists the Amazon Web Services managed policy that is attached to a specified
// permission set.
func ssoadmin_ListManagedPoliciesInPermissionSet(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.ListManagedPoliciesInPermissionSetInput{
		// InstanceArn: *string, // Required
		// PermissionSetArn: *string, // Required
	}

	if len(_ssoadminInstanceArn) > 0 {
		input.InstanceArn = aws.String(_ssoadminInstanceArn)
	}
	if len(_ssoadminPermissionSetArn) > 0 {
		input.PermissionSetArn = aws.String(_ssoadminPermissionSetArn)
	}
	if len(_ssoadminMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssoadminMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssoadminNextToken) > 0 {
		input.NextToken = aws.String(_ssoadminNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListManagedPoliciesInPermissionSet(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssoadmin.ListManagedPoliciesInPermissionSetOutput
	p := ssoadmin.NewListManagedPoliciesInPermissionSetPaginator(client, input)
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

// Lists the status of the permission set provisioning requests for a specified
// IAM Identity Center instance.
func ssoadmin_ListPermissionSetProvisioningStatus(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.ListPermissionSetProvisioningStatusInput{
		// InstanceArn: *string, // Required
	}

	if len(_ssoadminInstanceArn) > 0 {
		input.InstanceArn = aws.String(_ssoadminInstanceArn)
	}
	if len(_ssoadminFilter) > 0 {
		if err := assignInputField(input, "Filter", _ssoadminFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_ssoadminMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssoadminMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssoadminNextToken) > 0 {
		input.NextToken = aws.String(_ssoadminNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPermissionSetProvisioningStatus(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssoadmin.ListPermissionSetProvisioningStatusOutput
	p := ssoadmin.NewListPermissionSetProvisioningStatusPaginator(client, input)
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

// Lists the PermissionSets in an IAM Identity Center instance.
func ssoadmin_ListPermissionSets(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.ListPermissionSetsInput{
		// InstanceArn: *string, // Required
	}

	if len(_ssoadminInstanceArn) > 0 {
		input.InstanceArn = aws.String(_ssoadminInstanceArn)
	}
	if len(_ssoadminMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssoadminMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssoadminNextToken) > 0 {
		input.NextToken = aws.String(_ssoadminNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPermissionSets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssoadmin.ListPermissionSetsOutput
	p := ssoadmin.NewListPermissionSetsPaginator(client, input)
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

// Lists all the permission sets that are provisioned to a specified Amazon Web
// Services account.
func ssoadmin_ListPermissionSetsProvisionedToAccount(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.ListPermissionSetsProvisionedToAccountInput{
		// AccountId: *string, // Required
		// InstanceArn: *string, // Required
	}

	if len(_ssoadminAccountId) > 0 {
		input.AccountId = aws.String(_ssoadminAccountId)
	}
	if len(_ssoadminInstanceArn) > 0 {
		input.InstanceArn = aws.String(_ssoadminInstanceArn)
	}
	if len(_ssoadminMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssoadminMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssoadminNextToken) > 0 {
		input.NextToken = aws.String(_ssoadminNextToken)
	}
	if len(_ssoadminProvisioningStatus) > 0 {
		if err := assignInputField(input, "ProvisioningStatus", _ssoadminProvisioningStatus); err != nil {
			log.Errorf("invalid --provisioning-status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListPermissionSetsProvisionedToAccount(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssoadmin.ListPermissionSetsProvisionedToAccountOutput
	p := ssoadmin.NewListPermissionSetsProvisionedToAccountPaginator(client, input)
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

// Lists all enabled Regions of an IAM Identity Center instance, including those
// that are being added or removed. This operation returns Regions with ACTIVE,
// ADDING, or REMOVING status.
//
// The following actions are related to ListRegions :
//
// [AddRegion]
//
// [RemoveRegion]
//
// [DescribeRegion]
//
// [AddRegion]: https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_AddRegion.html
// [RemoveRegion]: https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_RemoveRegion.html
// [DescribeRegion]: https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_DescribeRegion.html
func ssoadmin_ListRegions(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.ListRegionsInput{
		// InstanceArn: *string, // Required
	}

	if len(_ssoadminInstanceArn) > 0 {
		input.InstanceArn = aws.String(_ssoadminInstanceArn)
	}
	if len(_ssoadminMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssoadminMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssoadminNextToken) > 0 {
		input.NextToken = aws.String(_ssoadminNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRegions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssoadmin.ListRegionsOutput
	p := ssoadmin.NewListRegionsPaginator(client, input)
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

// Lists the tags that are attached to a specified resource.
func ssoadmin_ListTagsForResource(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_ssoadminResourceArn) > 0 {
		input.ResourceArn = aws.String(_ssoadminResourceArn)
	}
	if len(_ssoadminInstanceArn) > 0 {
		input.InstanceArn = aws.String(_ssoadminInstanceArn)
	}
	if len(_ssoadminNextToken) > 0 {
		input.NextToken = aws.String(_ssoadminNextToken)
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

	var results []*ssoadmin.ListTagsForResourceOutput
	p := ssoadmin.NewListTagsForResourcePaginator(client, input)
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

// Lists all the trusted token issuers configured in an instance of IAM Identity
// Center.
func ssoadmin_ListTrustedTokenIssuers(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.ListTrustedTokenIssuersInput{
		// InstanceArn: *string, // Required
	}

	if len(_ssoadminInstanceArn) > 0 {
		input.InstanceArn = aws.String(_ssoadminInstanceArn)
	}
	if len(_ssoadminMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssoadminMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssoadminNextToken) > 0 {
		input.NextToken = aws.String(_ssoadminNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTrustedTokenIssuers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssoadmin.ListTrustedTokenIssuersOutput
	p := ssoadmin.NewListTrustedTokenIssuersPaginator(client, input)
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

// The process by which a specified permission set is provisioned to the specified
// target.
func ssoadmin_ProvisionPermissionSet(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.ProvisionPermissionSetInput{
		// InstanceArn: *string, // Required
		// PermissionSetArn: *string, // Required
		// TargetType: types.ProvisionTargetType, // Required
	}

	if len(_ssoadminInstanceArn) > 0 {
		input.InstanceArn = aws.String(_ssoadminInstanceArn)
	}
	if len(_ssoadminPermissionSetArn) > 0 {
		input.PermissionSetArn = aws.String(_ssoadminPermissionSetArn)
	}
	if len(_ssoadminTargetType) > 0 {
		if err := assignInputField(input, "TargetType", _ssoadminTargetType); err != nil {
			log.Errorf("invalid --target-type: %s", err.Error())
			return
		}
	}
	if len(_ssoadminTargetId) > 0 {
		input.TargetId = aws.String(_ssoadminTargetId)
	}

	if resp, err := client.ProvisionPermissionSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds or updates the list of authorized targets for an IAM Identity Center
// access scope for an application.
func ssoadmin_PutApplicationAccessScope(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.PutApplicationAccessScopeInput{
		// ApplicationArn: *string, // Required
		// Scope: *string, // Required
	}

	if len(_ssoadminApplicationArn) > 0 {
		input.ApplicationArn = aws.String(_ssoadminApplicationArn)
	}
	if len(_ssoadminScope) > 0 {
		input.Scope = aws.String(_ssoadminScope)
	}
	if len(_ssoadminAuthorizedTargets) > 0 {
		input.AuthorizedTargets = append([]string(nil), _ssoadminAuthorizedTargets...)
	}

	if resp, err := client.PutApplicationAccessScope(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Configure how users gain access to an application. If AssignmentsRequired is
// true (default value), users don’t have access to the application unless an
// assignment is created using the [CreateApplicationAssignment API]. If false , all users have access to the
// application. If an assignment is created using [CreateApplicationAssignment]., the user retains access if
// AssignmentsRequired is set to true .
//
// [CreateApplicationAssignment API]: https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_CreateApplicationAssignment.html
// [CreateApplicationAssignment]: https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_CreateApplicationAssignment.html
func ssoadmin_PutApplicationAssignmentConfiguration(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.PutApplicationAssignmentConfigurationInput{
		// ApplicationArn: *string, // Required
		// AssignmentRequired: *bool, // Required
	}

	if len(_ssoadminApplicationArn) > 0 {
		input.ApplicationArn = aws.String(_ssoadminApplicationArn)
	}
	if len(_ssoadminAssignmentRequired) > 0 {
		if err := assignInputField(input, "AssignmentRequired", _ssoadminAssignmentRequired); err != nil {
			log.Errorf("invalid --assignment-required: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutApplicationAssignmentConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds or updates an authentication method for an application.
func ssoadmin_PutApplicationAuthenticationMethod(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.PutApplicationAuthenticationMethodInput{
		// ApplicationArn: *string, // Required
		// AuthenticationMethod: types.AuthenticationMethod, // Required
		// AuthenticationMethodType: types.AuthenticationMethodType, // Required
	}

	if len(_ssoadminApplicationArn) > 0 {
		input.ApplicationArn = aws.String(_ssoadminApplicationArn)
	}
	if len(_ssoadminAuthenticationMethod) > 0 {
		if err := assignInputField(input, "AuthenticationMethod", _ssoadminAuthenticationMethod); err != nil {
			log.Errorf("invalid --authentication-method: %s", err.Error())
			return
		}
	}
	if len(_ssoadminAuthenticationMethodType) > 0 {
		if err := assignInputField(input, "AuthenticationMethodType", _ssoadminAuthenticationMethodType); err != nil {
			log.Errorf("invalid --authentication-method-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutApplicationAuthenticationMethod(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a configuration for an application to use grants. Conceptually grants
// are authorization to request actions related to tokens. This configuration will
// be used when parties are requesting and receiving tokens during the trusted
// identity propagation process. For more information on the IAM Identity Center
// supported grant workflows, see [SAML 2.0 and OAuth 2.0].
//
// A grant is created between your applications and Identity Center instance which
// enables an application to use specified mechanisms to obtain tokens. These
// tokens are used by your applications to gain access to Amazon Web Services
// resources on behalf of users. The following elements are within these exchanges:
//
// - Requester - The application requesting access to Amazon Web Services
// resources.
//
// - Subject - Typically the user that is requesting access to Amazon Web
// Services resources.
//
// - Grant - Conceptually, a grant is authorization to access Amazon Web
// Services resources. These grants authorize token generation for authenticating
// access to the requester and for the request to make requests on behalf of the
// subjects. There are four types of grants:
//
// - AuthorizationCode - Allows an application to request authorization through
// a series of user-agent redirects.
//
// - JWT bearer - Authorizes an application to exchange a JSON Web Token that
// came from an external identity provider. To learn more, see [RFC 6479].
//
// - Refresh token - Enables application to request new access tokens to replace
// expiring or expired access tokens.
//
// - Exchange token - A grant that requests tokens from the authorization server
// by providing a ‘subject’ token with access scope authorizing trusted identity
// propagation to this application. To learn more, see [RFC 8693].
//
// - Authorization server - IAM Identity Center requests tokens.
//
// User credentials are never shared directly within these exchanges. Instead,
// applications use grants to request access tokens from IAM Identity Center. For
// more information, see [RFC 6479].
//
// # Use cases
//
// - Connecting to custom applications.
//
// - Configuring an Amazon Web Services service to make calls to another Amazon
// Web Services services using JWT tokens.
//
// [RFC 6479]: https://datatracker.ietf.org/doc/html/rfc6749
// [SAML 2.0 and OAuth 2.0]: https://docs.aws.amazon.com/singlesignon/latest/userguide/customermanagedapps-saml2-oauth2.html
// [RFC 8693]: https://datatracker.ietf.org/doc/html/rfc8693
func ssoadmin_PutApplicationGrant(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.PutApplicationGrantInput{
		// ApplicationArn: *string, // Required
		// Grant: types.Grant, // Required
		// GrantType: types.GrantType, // Required
	}

	if len(_ssoadminApplicationArn) > 0 {
		input.ApplicationArn = aws.String(_ssoadminApplicationArn)
	}
	if len(_ssoadminGrant) > 0 {
		if err := assignInputField(input, "Grant", _ssoadminGrant); err != nil {
			log.Errorf("invalid --grant: %s", err.Error())
			return
		}
	}
	if len(_ssoadminGrantType) > 0 {
		if err := assignInputField(input, "GrantType", _ssoadminGrantType); err != nil {
			log.Errorf("invalid --grant-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutApplicationGrant(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the session configuration for an application in IAM Identity Center.
// The session configuration determines how users can access an application. This
// includes whether user background sessions are enabled. User background sessions
// allow users to start a job on a supported Amazon Web Services managed
// application without having to remain signed in to an active session while the
// job runs.
func ssoadmin_PutApplicationSessionConfiguration(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.PutApplicationSessionConfigurationInput{
		// ApplicationArn: *string, // Required
	}

	if len(_ssoadminApplicationArn) > 0 {
		input.ApplicationArn = aws.String(_ssoadminApplicationArn)
	}
	if len(_ssoadminUserBackgroundSessionApplicationStatus) > 0 {
		if err := assignInputField(input, "UserBackgroundSessionApplicationStatus", _ssoadminUserBackgroundSessionApplicationStatus); err != nil {
			log.Errorf("invalid --user-background-session-application-status: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutApplicationSessionConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Attaches an inline policy to a permission set.
// If the permission set is already referenced by one or more account assignments,
// you will need to call ProvisionPermissionSetafter this action to apply the corresponding IAM policy
// updates to all assigned accounts.
func ssoadmin_PutInlinePolicyToPermissionSet(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.PutInlinePolicyToPermissionSetInput{
		// InlinePolicy: *string, // Required
		// InstanceArn: *string, // Required
		// PermissionSetArn: *string, // Required
	}

	if len(_ssoadminInlinePolicy) > 0 {
		input.InlinePolicy = aws.String(_ssoadminInlinePolicy)
	}
	if len(_ssoadminInstanceArn) > 0 {
		input.InstanceArn = aws.String(_ssoadminInstanceArn)
	}
	if len(_ssoadminPermissionSetArn) > 0 {
		input.PermissionSetArn = aws.String(_ssoadminPermissionSetArn)
	}

	if resp, err := client.PutInlinePolicyToPermissionSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Attaches an Amazon Web Services managed or customer managed policy to the
// specified PermissionSetas a permissions boundary.
func ssoadmin_PutPermissionsBoundaryToPermissionSet(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.PutPermissionsBoundaryToPermissionSetInput{
		// InstanceArn: *string, // Required
		// PermissionSetArn: *string, // Required
		// PermissionsBoundary: *types.PermissionsBoundary, // Required
	}

	if len(_ssoadminInstanceArn) > 0 {
		input.InstanceArn = aws.String(_ssoadminInstanceArn)
	}
	if len(_ssoadminPermissionSetArn) > 0 {
		input.PermissionSetArn = aws.String(_ssoadminPermissionSetArn)
	}
	if len(_ssoadminPermissionsBoundary) > 0 {
		if err := assignInputField(input, "PermissionsBoundary", _ssoadminPermissionsBoundary); err != nil {
			log.Errorf("invalid --permissions-boundary: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutPermissionsBoundaryToPermissionSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes an additional Region from an IAM Identity Center instance. This
// operation initiates an asynchronous workflow to clean up IAM Identity Center
// resources in the specified additional Region. The Region status is set to
// REMOVING and the Region record is deleted when the workflow completes. The
// request must be made from the primary Region. The target Region cannot be the
// primary Region, and no other add or remove Region workflows can be in progress.
//
// The following actions are related to RemoveRegion :
//
// [AddRegion]
//
// [DescribeRegion]
//
// [ListRegions]
//
// [AddRegion]: https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_AddRegion.html
// [DescribeRegion]: https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_DescribeRegion.html
// [ListRegions]: https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_ListRegions.html
func ssoadmin_RemoveRegion(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.RemoveRegionInput{
		// InstanceArn: *string, // Required
		// RegionName: *string, // Required
	}

	if len(_ssoadminInstanceArn) > 0 {
		input.InstanceArn = aws.String(_ssoadminInstanceArn)
	}
	if len(_ssoadminRegionName) > 0 {
		input.RegionName = aws.String(_ssoadminRegionName)
	}

	if resp, err := client.RemoveRegion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates a set of tags with a specified resource.
func ssoadmin_TagResource(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_ssoadminResourceArn) > 0 {
		input.ResourceArn = aws.String(_ssoadminResourceArn)
	}
	if len(_ssoadminTags) > 0 {
		if err := assignInputField(input, "Tags", _ssoadminTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_ssoadminInstanceArn) > 0 {
		input.InstanceArn = aws.String(_ssoadminInstanceArn)
	}

	if resp, err := client.TagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates a set of tags from a specified resource.
func ssoadmin_UntagResource(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_ssoadminResourceArn) > 0 {
		input.ResourceArn = aws.String(_ssoadminResourceArn)
	}
	if len(_ssoadminTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _ssoadminTagKeys...)
	}
	if len(_ssoadminInstanceArn) > 0 {
		input.InstanceArn = aws.String(_ssoadminInstanceArn)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates application properties.
func ssoadmin_UpdateApplication(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.UpdateApplicationInput{
		// ApplicationArn: *string, // Required
	}

	if len(_ssoadminApplicationArn) > 0 {
		input.ApplicationArn = aws.String(_ssoadminApplicationArn)
	}
	if len(_ssoadminDescription) > 0 {
		input.Description = aws.String(_ssoadminDescription)
	}
	if len(_ssoadminName) > 0 {
		input.Name = aws.String(_ssoadminName)
	}
	if len(_ssoadminPortalOptions) > 0 {
		if err := assignInputField(input, "PortalOptions", _ssoadminPortalOptions); err != nil {
			log.Errorf("invalid --portal-options: %s", err.Error())
			return
		}
	}
	if len(_ssoadminStatus) > 0 {
		if err := assignInputField(input, "Status", _ssoadminStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update the details for the instance of IAM Identity Center that is owned by the
// Amazon Web Services account.
func ssoadmin_UpdateInstance(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.UpdateInstanceInput{
		// InstanceArn: *string, // Required
	}

	if len(_ssoadminInstanceArn) > 0 {
		input.InstanceArn = aws.String(_ssoadminInstanceArn)
	}
	if len(_ssoadminEncryptionConfiguration) > 0 {
		if err := assignInputField(input, "EncryptionConfiguration", _ssoadminEncryptionConfiguration); err != nil {
			log.Errorf("invalid --encryption-configuration: %s", err.Error())
			return
		}
	}
	if len(_ssoadminName) > 0 {
		input.Name = aws.String(_ssoadminName)
	}

	if resp, err := client.UpdateInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the IAM Identity Center identity store attributes that you can use with
// the IAM Identity Center instance for attributes-based access control (ABAC).
// When using an external identity provider as an identity source, you can pass
// attributes through the SAML assertion as an alternative to configuring
// attributes from the IAM Identity Center identity store. If a SAML assertion
// passes any of these attributes, IAM Identity Center replaces the attribute value
// with the value from the IAM Identity Center identity store. For more information
// about ABAC, see Attribute-Based Access Controlin the IAM Identity Center User Guide.
func ssoadmin_UpdateInstanceAccessControlAttributeConfiguration(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.UpdateInstanceAccessControlAttributeConfigurationInput{
		// InstanceAccessControlAttributeConfiguration: *types.InstanceAccessControlAttributeConfiguration, // Required
		// InstanceArn: *string, // Required
	}

	if len(_ssoadminInstanceAccessControlAttributeConfiguration) > 0 {
		if err := assignInputField(input, "InstanceAccessControlAttributeConfiguration", _ssoadminInstanceAccessControlAttributeConfiguration); err != nil {
			log.Errorf("invalid --instance-access-control-attribute-configuration: %s", err.Error())
			return
		}
	}
	if len(_ssoadminInstanceArn) > 0 {
		input.InstanceArn = aws.String(_ssoadminInstanceArn)
	}

	if resp, err := client.UpdateInstanceAccessControlAttributeConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing permission set.
func ssoadmin_UpdatePermissionSet(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.UpdatePermissionSetInput{
		// InstanceArn: *string, // Required
		// PermissionSetArn: *string, // Required
	}

	if len(_ssoadminInstanceArn) > 0 {
		input.InstanceArn = aws.String(_ssoadminInstanceArn)
	}
	if len(_ssoadminPermissionSetArn) > 0 {
		input.PermissionSetArn = aws.String(_ssoadminPermissionSetArn)
	}
	if len(_ssoadminDescription) > 0 {
		input.Description = aws.String(_ssoadminDescription)
	}
	if len(_ssoadminRelayState) > 0 {
		input.RelayState = aws.String(_ssoadminRelayState)
	}
	if len(_ssoadminSessionDuration) > 0 {
		input.SessionDuration = aws.String(_ssoadminSessionDuration)
	}

	if resp, err := client.UpdatePermissionSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the name of the trusted token issuer, or the path of a source attribute
// or destination attribute for a trusted token issuer configuration.
//
// Updating this trusted token issuer configuration might cause users to lose
// access to any applications that are configured to use the trusted token issuer.
func ssoadmin_UpdateTrustedTokenIssuer(cfg aws.Config, client *ssoadmin.Client) {
	input := &ssoadmin.UpdateTrustedTokenIssuerInput{
		// TrustedTokenIssuerArn: *string, // Required
	}

	if len(_ssoadminTrustedTokenIssuerArn) > 0 {
		input.TrustedTokenIssuerArn = aws.String(_ssoadminTrustedTokenIssuerArn)
	}
	if len(_ssoadminName) > 0 {
		input.Name = aws.String(_ssoadminName)
	}
	if len(_ssoadminTrustedTokenIssuerConfiguration) > 0 {
		if err := assignInputField(input, "TrustedTokenIssuerConfiguration", _ssoadminTrustedTokenIssuerConfiguration); err != nil {
			log.Errorf("invalid --trusted-token-issuer-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateTrustedTokenIssuer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_ssoadminCmd)
	_ssoadminCmd.Flags().SortFlags = false

	_ssoadminCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_ssoadminCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_ssoadminCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_ssoadminCmd.Flags().StringVarP(&_ssoadminAccountAssignmentCreationRequestId, "account-assignment-creation-request-id", "", "", "Account Assignment Creation Request ID")
	_ssoadminCmd.Flags().StringVarP(&_ssoadminAccountAssignmentDeletionRequestId, "account-assignment-deletion-request-id", "", "", "Account Assignment Deletion Request ID")
	_ssoadminCmd.Flags().StringVarP(&_ssoadminAccountId, "account-id", "", "", "Account ID")
	_ssoadminCmd.Flags().StringVarP(&_ssoadminApplicationArn, "application-arn", "", "", "Application ARN")
	_ssoadminCmd.Flags().StringVarP(&_ssoadminApplicationProviderArn, "application-provider-arn", "", "", "Application Provider ARN")
	_ssoadminCmd.Flags().StringVarP(&_ssoadminAssignmentRequired, "assignment-required", "", "", "Assignment Required")
	_ssoadminCmd.Flags().StringVarP(&_ssoadminAuthenticationMethod, "authentication-method", "", "", "Authentication Method")
	_ssoadminCmd.Flags().StringVarP(&_ssoadminAuthenticationMethodType, "authentication-method-type", "", "", "Authentication Method Type")
	_ssoadminCmd.Flags().StringSliceVarP(&_ssoadminAuthorizedTargets, "authorized-targets", "", nil, "Authorized Targets")
	_ssoadminCmd.Flags().StringVarP(&_ssoadminClientToken, "client-token", "", "", "Client Token")
	_ssoadminCmd.Flags().StringVarP(&_ssoadminCustomerManagedPolicyReference, "customer-managed-policy-reference", "", "", "Customer Managed Policy Reference")
	_ssoadminCmd.Flags().StringVarP(&_ssoadminDescription, "description", "", "", "Description")
	_ssoadminCmd.Flags().StringVarP(&_ssoadminEncryptionConfiguration, "encryption-configuration", "", "", "Encryption Configuration")
	_ssoadminCmd.Flags().StringVarP(&_ssoadminFilter, "filter", "", "", "Filter")
	_ssoadminCmd.Flags().StringVarP(&_ssoadminGrant, "grant", "", "", "Grant")
	_ssoadminCmd.Flags().StringVarP(&_ssoadminGrantType, "grant-type", "", "", "Grant Type")
	_ssoadminCmd.Flags().StringVarP(&_ssoadminInlinePolicy, "inline-policy", "", "", "Inline Policy")
	_ssoadminCmd.Flags().StringVarP(&_ssoadminInstanceAccessControlAttributeConfiguration, "instance-access-control-attribute-configuration", "", "", "Instance Access Control Attribute Configuration")
	_ssoadminCmd.Flags().StringVarP(&_ssoadminInstanceArn, "instance-arn", "", "", "Instance ARN")
	_ssoadminCmd.Flags().StringVarP(&_ssoadminManagedPolicyArn, "managed-policy-arn", "", "", "Managed Policy ARN")
	_ssoadminCmd.Flags().StringVarP(&_ssoadminMaxResults, "max-results", "", "", "Max Results")
	_ssoadminCmd.Flags().StringVarP(&_ssoadminName, "name", "", "", "Name")
	_ssoadminCmd.Flags().StringVarP(&_ssoadminNextToken, "next-token", "", "", "Next Token")
	_ssoadminCmd.Flags().StringVarP(&_ssoadminPermissionSetArn, "permission-set-arn", "", "", "Permission Set ARN")
	_ssoadminCmd.Flags().StringVarP(&_ssoadminPermissionsBoundary, "permissions-boundary", "", "", "Permissions Boundary")
	_ssoadminCmd.Flags().StringVarP(&_ssoadminPortalOptions, "portal-options", "", "", "Portal Options")
	_ssoadminCmd.Flags().StringVarP(&_ssoadminPrincipalId, "principal-id", "", "", "Principal ID")
	_ssoadminCmd.Flags().StringVarP(&_ssoadminPrincipalType, "principal-type", "", "", "Principal Type")
	_ssoadminCmd.Flags().StringVarP(&_ssoadminProvisionPermissionSetRequestId, "provision-permission-set-request-id", "", "", "Provision Permission Set Request ID")
	_ssoadminCmd.Flags().StringVarP(&_ssoadminProvisioningStatus, "provisioning-status", "", "", "Provisioning Status")
	_ssoadminCmd.Flags().StringVarP(&_ssoadminRegionName, "region-name", "", "", "Region Name")
	_ssoadminCmd.Flags().StringVarP(&_ssoadminRelayState, "relay-state", "", "", "Relay State")
	_ssoadminCmd.Flags().StringVarP(&_ssoadminResourceArn, "resource-arn", "", "", "Resource ARN")
	_ssoadminCmd.Flags().StringVarP(&_ssoadminScope, "scope", "", "", "Scope")
	_ssoadminCmd.Flags().StringVarP(&_ssoadminSessionDuration, "session-duration", "", "", "Session Duration")
	_ssoadminCmd.Flags().StringVarP(&_ssoadminStatus, "status", "", "", "Status")
	_ssoadminCmd.Flags().StringSliceVarP(&_ssoadminTagKeys, "tag-keys", "", nil, "Tag Keys")
	_ssoadminCmd.Flags().StringVarP(&_ssoadminTags, "tags", "", "", "Tags")
	_ssoadminCmd.Flags().StringVarP(&_ssoadminTargetId, "target-id", "", "", "Target ID")
	_ssoadminCmd.Flags().StringVarP(&_ssoadminTargetType, "target-type", "", "", "Target Type")
	_ssoadminCmd.Flags().StringVarP(&_ssoadminTrustedTokenIssuerArn, "trusted-token-issuer-arn", "", "", "Trusted Token Issuer ARN")
	_ssoadminCmd.Flags().StringVarP(&_ssoadminTrustedTokenIssuerConfiguration, "trusted-token-issuer-configuration", "", "", "Trusted Token Issuer Configuration")
	_ssoadminCmd.Flags().StringVarP(&_ssoadminTrustedTokenIssuerType, "trusted-token-issuer-type", "", "", "Trusted Token Issuer Type")
	_ssoadminCmd.Flags().StringVarP(&_ssoadminUserBackgroundSessionApplicationStatus, "user-background-session-application-status", "", "", "User Background Session Application Status")

	_ssoadminCmd.Flags().BoolVarP(&_ssoadminAddRegion, "add-region", "", false, "Add Region")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminAttachCustomerManagedPolicyReferenceToPermissionSet, "attach-customer-managed-policy-reference-to-permission-set", "", false, "Attach Customer Managed Policy Reference To Permission Set")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminAttachManagedPolicyToPermissionSet, "attach-managed-policy-to-permission-set", "", false, "Attach Managed Policy To Permission Set")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminCreateAccountAssignment, "create-account-assignment", "", false, "Create Account Assignment")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminCreateApplication, "create-application", "", false, "Create Application")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminCreateApplicationAssignment, "create-application-assignment", "", false, "Create Application Assignment")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminCreateInstance, "create-instance", "", false, "Create Instance")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminCreateInstanceAccessControlAttributeConfiguration, "create-instance-access-control-attribute-configuration", "", false, "Create Instance Access Control Attribute Configuration")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminCreatePermissionSet, "create-permission-set", "", false, "Create Permission Set")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminCreateTrustedTokenIssuer, "create-trusted-token-issuer", "", false, "Create Trusted Token Issuer")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminDeleteAccountAssignment, "delete-account-assignment", "", false, "Delete Account Assignment")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminDeleteApplication, "delete-application", "", false, "Delete Application")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminDeleteApplicationAccessScope, "delete-application-access-scope", "", false, "Delete Application Access Scope")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminDeleteApplicationAssignment, "delete-application-assignment", "", false, "Delete Application Assignment")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminDeleteApplicationAuthenticationMethod, "delete-application-authentication-method", "", false, "Delete Application Authentication Method")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminDeleteApplicationGrant, "delete-application-grant", "", false, "Delete Application Grant")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminDeleteInlinePolicyFromPermissionSet, "delete-inline-policy-from-permission-set", "", false, "Delete Inline Policy From Permission Set")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminDeleteInstance, "delete-instance", "", false, "Delete Instance")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminDeleteInstanceAccessControlAttributeConfiguration, "delete-instance-access-control-attribute-configuration", "", false, "Delete Instance Access Control Attribute Configuration")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminDeletePermissionSet, "delete-permission-set", "", false, "Delete Permission Set")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminDeletePermissionsBoundaryFromPermissionSet, "delete-permissions-boundary-from-permission-set", "", false, "Delete Permissions Boundary From Permission Set")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminDeleteTrustedTokenIssuer, "delete-trusted-token-issuer", "", false, "Delete Trusted Token Issuer")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminDescribeAccountAssignmentCreationStatus, "describe-account-assignment-creation-status", "", false, "Describe Account Assignment Creation Status")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminDescribeAccountAssignmentDeletionStatus, "describe-account-assignment-deletion-status", "", false, "Describe Account Assignment Deletion Status")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminDescribeApplication, "describe-application", "", false, "Describe Application")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminDescribeApplicationAssignment, "describe-application-assignment", "", false, "Describe Application Assignment")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminDescribeApplicationProvider, "describe-application-provider", "", false, "Describe Application Provider")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminDescribeInstance, "describe-instance", "", false, "Describe Instance")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminDescribeInstanceAccessControlAttributeConfiguration, "describe-instance-access-control-attribute-configuration", "", false, "Describe Instance Access Control Attribute Configuration")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminDescribePermissionSet, "describe-permission-set", "", false, "Describe Permission Set")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminDescribePermissionSetProvisioningStatus, "describe-permission-set-provisioning-status", "", false, "Describe Permission Set Provisioning Status")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminDescribeRegion, "describe-region", "", false, "Describe Region")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminDescribeTrustedTokenIssuer, "describe-trusted-token-issuer", "", false, "Describe Trusted Token Issuer")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminDetachCustomerManagedPolicyReferenceFromPermissionSet, "detach-customer-managed-policy-reference-from-permission-set", "", false, "Detach Customer Managed Policy Reference From Permission Set")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminDetachManagedPolicyFromPermissionSet, "detach-managed-policy-from-permission-set", "", false, "Detach Managed Policy From Permission Set")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminGetApplicationAccessScope, "get-application-access-scope", "", false, "Get Application Access Scope")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminGetApplicationAssignmentConfiguration, "get-application-assignment-configuration", "", false, "Get Application Assignment Configuration")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminGetApplicationAuthenticationMethod, "get-application-authentication-method", "", false, "Get Application Authentication Method")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminGetApplicationGrant, "get-application-grant", "", false, "Get Application Grant")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminGetApplicationSessionConfiguration, "get-application-session-configuration", "", false, "Get Application Session Configuration")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminGetInlinePolicyForPermissionSet, "get-inline-policy-for-permission-set", "", false, "Get Inline Policy For Permission Set")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminGetPermissionsBoundaryForPermissionSet, "get-permissions-boundary-for-permission-set", "", false, "Get Permissions Boundary For Permission Set")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminListAccountAssignmentCreationStatus, "list-account-assignment-creation-status", "", false, "List Account Assignment Creation Status")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminListAccountAssignmentDeletionStatus, "list-account-assignment-deletion-status", "", false, "List Account Assignment Deletion Status")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminListAccountAssignments, "list-account-assignments", "", false, "List Account Assignments")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminListAccountAssignmentsForPrincipal, "list-account-assignments-for-principal", "", false, "List Account Assignments For Principal")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminListAccountsForProvisionedPermissionSet, "list-accounts-for-provisioned-permission-set", "", false, "List Accounts For Provisioned Permission Set")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminListApplicationAccessScopes, "list-application-access-scopes", "", false, "List Application Access Scopes")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminListApplicationAssignments, "list-application-assignments", "", false, "List Application Assignments")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminListApplicationAssignmentsForPrincipal, "list-application-assignments-for-principal", "", false, "List Application Assignments For Principal")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminListApplicationAuthenticationMethods, "list-application-authentication-methods", "", false, "List Application Authentication Methods")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminListApplicationGrants, "list-application-grants", "", false, "List Application Grants")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminListApplicationProviders, "list-application-providers", "", false, "List Application Providers")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminListApplications, "list-applications", "", false, "List Applications")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminListCustomerManagedPolicyReferencesInPermissionSet, "list-customer-managed-policy-references-in-permission-set", "", false, "List Customer Managed Policy References In Permission Set")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminListInstances, "list-instances", "", false, "List Instances")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminListManagedPoliciesInPermissionSet, "list-managed-policies-in-permission-set", "", false, "List Managed Policies In Permission Set")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminListPermissionSetProvisioningStatus, "list-permission-set-provisioning-status", "", false, "List Permission Set Provisioning Status")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminListPermissionSets, "list-permission-sets", "", false, "List Permission Sets")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminListPermissionSetsProvisionedToAccount, "list-permission-sets-provisioned-to-account", "", false, "List Permission Sets Provisioned To Account")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminListRegions, "list-regions", "", false, "List Regions")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminListTrustedTokenIssuers, "list-trusted-token-issuers", "", false, "List Trusted Token Issuers")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminProvisionPermissionSet, "provision-permission-set", "", false, "Provision Permission Set")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminPutApplicationAccessScope, "put-application-access-scope", "", false, "Put Application Access Scope")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminPutApplicationAssignmentConfiguration, "put-application-assignment-configuration", "", false, "Put Application Assignment Configuration")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminPutApplicationAuthenticationMethod, "put-application-authentication-method", "", false, "Put Application Authentication Method")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminPutApplicationGrant, "put-application-grant", "", false, "Put Application Grant")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminPutApplicationSessionConfiguration, "put-application-session-configuration", "", false, "Put Application Session Configuration")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminPutInlinePolicyToPermissionSet, "put-inline-policy-to-permission-set", "", false, "Put Inline Policy To Permission Set")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminPutPermissionsBoundaryToPermissionSet, "put-permissions-boundary-to-permission-set", "", false, "Put Permissions Boundary To Permission Set")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminRemoveRegion, "remove-region", "", false, "Remove Region")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminTagResource, "tag-resource", "", false, "Tag Resource")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminUntagResource, "untag-resource", "", false, "Untag Resource")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminUpdateApplication, "update-application", "", false, "Update Application")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminUpdateInstance, "update-instance", "", false, "Update Instance")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminUpdateInstanceAccessControlAttributeConfiguration, "update-instance-access-control-attribute-configuration", "", false, "Update Instance Access Control Attribute Configuration")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminUpdatePermissionSet, "update-permission-set", "", false, "Update Permission Set")
	_ssoadminCmd.Flags().BoolVarP(&_ssoadminUpdateTrustedTokenIssuer, "update-trusted-token-issuer", "", false, "Update Trusted Token Issuer")

}
