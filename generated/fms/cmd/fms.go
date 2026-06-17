package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/fms"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// fmsCmd represents the fms command
var _fmsCmd = &cobra.Command{
	Use:   "fms",
	Short: "AWS fms CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := fms.NewFromConfig(cfg)
		if _fmsAssociateAdminAccount {
			fms_AssociateAdminAccount(cfg, client)
			return
		}
		if _fmsAssociateThirdPartyFirewall {
			fms_AssociateThirdPartyFirewall(cfg, client)
			return
		}
		if _fmsBatchAssociateResource {
			fms_BatchAssociateResource(cfg, client)
			return
		}
		if _fmsBatchDisassociateResource {
			fms_BatchDisassociateResource(cfg, client)
			return
		}
		if _fmsDeleteAppsList {
			fms_DeleteAppsList(cfg, client)
			return
		}
		if _fmsDeleteNotificationChannel {
			fms_DeleteNotificationChannel(cfg, client)
			return
		}
		if _fmsDeletePolicy {
			fms_DeletePolicy(cfg, client)
			return
		}
		if _fmsDeleteProtocolsList {
			fms_DeleteProtocolsList(cfg, client)
			return
		}
		if _fmsDeleteResourceSet {
			fms_DeleteResourceSet(cfg, client)
			return
		}
		if _fmsDisassociateAdminAccount {
			fms_DisassociateAdminAccount(cfg, client)
			return
		}
		if _fmsDisassociateThirdPartyFirewall {
			fms_DisassociateThirdPartyFirewall(cfg, client)
			return
		}
		if _fmsGetAdminAccount {
			fms_GetAdminAccount(cfg, client)
			return
		}
		if _fmsGetAdminScope {
			fms_GetAdminScope(cfg, client)
			return
		}
		if _fmsGetAppsList {
			fms_GetAppsList(cfg, client)
			return
		}
		if _fmsGetComplianceDetail {
			fms_GetComplianceDetail(cfg, client)
			return
		}
		if _fmsGetNotificationChannel {
			fms_GetNotificationChannel(cfg, client)
			return
		}
		if _fmsGetPolicy {
			fms_GetPolicy(cfg, client)
			return
		}
		if _fmsGetProtectionStatus {
			fms_GetProtectionStatus(cfg, client)
			return
		}
		if _fmsGetProtocolsList {
			fms_GetProtocolsList(cfg, client)
			return
		}
		if _fmsGetResourceSet {
			fms_GetResourceSet(cfg, client)
			return
		}
		if _fmsGetThirdPartyFirewallAssociationStatus {
			fms_GetThirdPartyFirewallAssociationStatus(cfg, client)
			return
		}
		if _fmsGetViolationDetails {
			fms_GetViolationDetails(cfg, client)
			return
		}
		if _fmsListAdminAccountsForOrganization {
			fms_ListAdminAccountsForOrganization(cfg, client)
			return
		}
		if _fmsListAdminsManagingAccount {
			fms_ListAdminsManagingAccount(cfg, client)
			return
		}
		if _fmsListAppsLists {
			fms_ListAppsLists(cfg, client)
			return
		}
		if _fmsListComplianceStatus {
			fms_ListComplianceStatus(cfg, client)
			return
		}
		if _fmsListDiscoveredResources {
			fms_ListDiscoveredResources(cfg, client)
			return
		}
		if _fmsListMemberAccounts {
			fms_ListMemberAccounts(cfg, client)
			return
		}
		if _fmsListPolicies {
			fms_ListPolicies(cfg, client)
			return
		}
		if _fmsListProtocolsLists {
			fms_ListProtocolsLists(cfg, client)
			return
		}
		if _fmsListResourceSetResources {
			fms_ListResourceSetResources(cfg, client)
			return
		}
		if _fmsListResourceSets {
			fms_ListResourceSets(cfg, client)
			return
		}
		if _fmsListTagsForResource {
			fms_ListTagsForResource(cfg, client)
			return
		}
		if _fmsListThirdPartyFirewallFirewallPolicies {
			fms_ListThirdPartyFirewallFirewallPolicies(cfg, client)
			return
		}
		if _fmsPutAdminAccount {
			fms_PutAdminAccount(cfg, client)
			return
		}
		if _fmsPutAppsList {
			fms_PutAppsList(cfg, client)
			return
		}
		if _fmsPutNotificationChannel {
			fms_PutNotificationChannel(cfg, client)
			return
		}
		if _fmsPutPolicy {
			fms_PutPolicy(cfg, client)
			return
		}
		if _fmsPutProtocolsList {
			fms_PutProtocolsList(cfg, client)
			return
		}
		if _fmsPutResourceSet {
			fms_PutResourceSet(cfg, client)
			return
		}
		if _fmsTagResource {
			fms_TagResource(cfg, client)
			return
		}
		if _fmsUntagResource {
			fms_UntagResource(cfg, client)
			return
		}

	},
}

var (
	_fmsAssociateAdminAccount                  bool
	_fmsAssociateThirdPartyFirewall            bool
	_fmsBatchAssociateResource                 bool
	_fmsBatchDisassociateResource              bool
	_fmsDeleteAppsList                         bool
	_fmsDeleteNotificationChannel              bool
	_fmsDeletePolicy                           bool
	_fmsDeleteProtocolsList                    bool
	_fmsDeleteResourceSet                      bool
	_fmsDisassociateAdminAccount               bool
	_fmsDisassociateThirdPartyFirewall         bool
	_fmsGetAdminAccount                        bool
	_fmsGetAdminScope                          bool
	_fmsGetAppsList                            bool
	_fmsGetComplianceDetail                    bool
	_fmsGetNotificationChannel                 bool
	_fmsGetPolicy                              bool
	_fmsGetProtectionStatus                    bool
	_fmsGetProtocolsList                       bool
	_fmsGetResourceSet                         bool
	_fmsGetThirdPartyFirewallAssociationStatus bool
	_fmsGetViolationDetails                    bool
	_fmsListAdminAccountsForOrganization       bool
	_fmsListAdminsManagingAccount              bool
	_fmsListAppsLists                          bool
	_fmsListComplianceStatus                   bool
	_fmsListDiscoveredResources                bool
	_fmsListMemberAccounts                     bool
	_fmsListPolicies                           bool
	_fmsListProtocolsLists                     bool
	_fmsListResourceSetResources               bool
	_fmsListResourceSets                       bool
	_fmsListTagsForResource                    bool
	_fmsListThirdPartyFirewallFirewallPolicies bool
	_fmsPutAdminAccount                        bool
	_fmsPutAppsList                            bool
	_fmsPutNotificationChannel                 bool
	_fmsPutPolicy                              bool
	_fmsPutProtocolsList                       bool
	_fmsPutResourceSet                         bool
	_fmsTagResource                            bool
	_fmsUntagResource                          bool

	_fmsAdminAccount             string
	_fmsAdminScope               string
	_fmsAppsList                 string
	_fmsDefaultList              string
	_fmsDefaultLists             string
	_fmsDeleteAllPolicyResources string
	_fmsEndTime                  string
	_fmsIdentifier               string
	_fmsItems                    []string
	_fmsListId                   string
	_fmsMaxResults               string
	_fmsMemberAccount            string
	_fmsMemberAccountId          string
	_fmsMemberAccountIds         []string
	_fmsNextToken                string
	_fmsPolicy                   string
	_fmsPolicyId                 string
	_fmsProtocolsList            string
	_fmsResourceArn              string
	_fmsResourceId               string
	_fmsResourceSet              string
	_fmsResourceSetIdentifier    string
	_fmsResourceType             string
	_fmsSnsRoleName              string
	_fmsSnsTopicArn              string
	_fmsStartTime                string
	_fmsTagKeys                  []string
	_fmsTagList                  string
	_fmsThirdPartyFirewall       string
)

// Sets a Firewall Manager default administrator account. The Firewall Manager
// default administrator account can manage third-party firewalls and has full
// administrative scope that allows administration of all policy types, accounts,
// organizational units, and Regions. This account must be a member account of the
// organization in Organizations whose resources you want to protect.
//
// For information about working with Firewall Manager administrator accounts, see [Managing Firewall Manager administrators]
// in the Firewall Manager Developer Guide.
//
// [Managing Firewall Manager administrators]: https://docs.aws.amazon.com/organizations/latest/userguide/fms-administrators.html
func fms_AssociateAdminAccount(cfg aws.Config, client *fms.Client) {
	input := &fms.AssociateAdminAccountInput{
		// AdminAccount: *string, // Required
	}

	if len(_fmsAdminAccount) > 0 {
		input.AdminAccount = aws.String(_fmsAdminAccount)
	}

	if resp, err := client.AssociateAdminAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets the Firewall Manager policy administrator as a tenant administrator of a
// third-party firewall service. A tenant is an instance of the third-party
// firewall service that's associated with your Amazon Web Services customer
// account.
func fms_AssociateThirdPartyFirewall(cfg aws.Config, client *fms.Client) {
	input := &fms.AssociateThirdPartyFirewallInput{
		// ThirdPartyFirewall: types.ThirdPartyFirewall, // Required
	}

	if len(_fmsThirdPartyFirewall) > 0 {
		if err := assignInputField(input, "ThirdPartyFirewall", _fmsThirdPartyFirewall); err != nil {
			log.Errorf("invalid --third-party-firewall: %s", err.Error())
			return
		}
	}

	if resp, err := client.AssociateThirdPartyFirewall(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associate resources to a Firewall Manager resource set.
func fms_BatchAssociateResource(cfg aws.Config, client *fms.Client) {
	input := &fms.BatchAssociateResourceInput{
		// Items: []string, // Required
		// ResourceSetIdentifier: *string, // Required
	}

	if len(_fmsItems) > 0 {
		input.Items = append([]string(nil), _fmsItems...)
	}
	if len(_fmsResourceSetIdentifier) > 0 {
		input.ResourceSetIdentifier = aws.String(_fmsResourceSetIdentifier)
	}

	if resp, err := client.BatchAssociateResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates resources from a Firewall Manager resource set.
func fms_BatchDisassociateResource(cfg aws.Config, client *fms.Client) {
	input := &fms.BatchDisassociateResourceInput{
		// Items: []string, // Required
		// ResourceSetIdentifier: *string, // Required
	}

	if len(_fmsItems) > 0 {
		input.Items = append([]string(nil), _fmsItems...)
	}
	if len(_fmsResourceSetIdentifier) > 0 {
		input.ResourceSetIdentifier = aws.String(_fmsResourceSetIdentifier)
	}

	if resp, err := client.BatchDisassociateResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Permanently deletes an Firewall Manager applications list.
func fms_DeleteAppsList(cfg aws.Config, client *fms.Client) {
	input := &fms.DeleteAppsListInput{
		// ListId: *string, // Required
	}

	if len(_fmsListId) > 0 {
		input.ListId = aws.String(_fmsListId)
	}

	if resp, err := client.DeleteAppsList(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Firewall Manager association with the IAM role and the Amazon Simple
// Notification Service (SNS) topic that is used to record Firewall Manager SNS
// logs.
func fms_DeleteNotificationChannel(cfg aws.Config, client *fms.Client) {
	input := &fms.DeleteNotificationChannelInput{}

	if resp, err := client.DeleteNotificationChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Permanently deletes an Firewall Manager policy.
func fms_DeletePolicy(cfg aws.Config, client *fms.Client) {
	input := &fms.DeletePolicyInput{
		// PolicyId: *string, // Required
	}

	if len(_fmsPolicyId) > 0 {
		input.PolicyId = aws.String(_fmsPolicyId)
	}
	if len(_fmsDeleteAllPolicyResources) > 0 {
		if err := assignInputField(input, "DeleteAllPolicyResources", _fmsDeleteAllPolicyResources); err != nil {
			log.Errorf("invalid --delete-all-policy-resources: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeletePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Permanently deletes an Firewall Manager protocols list.
func fms_DeleteProtocolsList(cfg aws.Config, client *fms.Client) {
	input := &fms.DeleteProtocolsListInput{
		// ListId: *string, // Required
	}

	if len(_fmsListId) > 0 {
		input.ListId = aws.String(_fmsListId)
	}

	if resp, err := client.DeleteProtocolsList(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified ResourceSet.
func fms_DeleteResourceSet(cfg aws.Config, client *fms.Client) {
	input := &fms.DeleteResourceSetInput{
		// Identifier: *string, // Required
	}

	if len(_fmsIdentifier) > 0 {
		input.Identifier = aws.String(_fmsIdentifier)
	}

	if resp, err := client.DeleteResourceSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates an Firewall Manager administrator account. To set a different
// account as an Firewall Manager administrator, submit a PutAdminAccountrequest. To set an
// account as a default administrator account, you must submit an AssociateAdminAccountrequest.
//
// Disassociation of the default administrator account follows the first in, last
// out principle. If you are the default administrator, all Firewall Manager
// administrators within the organization must first disassociate their accounts
// before you can disassociate your account.
func fms_DisassociateAdminAccount(cfg aws.Config, client *fms.Client) {
	input := &fms.DisassociateAdminAccountInput{}

	if resp, err := client.DisassociateAdminAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates a Firewall Manager policy administrator from a third-party
// firewall tenant. When you call DisassociateThirdPartyFirewall , the third-party
// firewall vendor deletes all of the firewalls that are associated with the
// account.
func fms_DisassociateThirdPartyFirewall(cfg aws.Config, client *fms.Client) {
	input := &fms.DisassociateThirdPartyFirewallInput{
		// ThirdPartyFirewall: types.ThirdPartyFirewall, // Required
	}

	if len(_fmsThirdPartyFirewall) > 0 {
		if err := assignInputField(input, "ThirdPartyFirewall", _fmsThirdPartyFirewall); err != nil {
			log.Errorf("invalid --third-party-firewall: %s", err.Error())
			return
		}
	}

	if resp, err := client.DisassociateThirdPartyFirewall(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the Organizations account that is associated with Firewall Manager as
// the Firewall Manager default administrator.
func fms_GetAdminAccount(cfg aws.Config, client *fms.Client) {
	input := &fms.GetAdminAccountInput{}

	if resp, err := client.GetAdminAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the specified account's administrative scope. The
// administrative scope defines the resources that an Firewall Manager
// administrator can manage.
func fms_GetAdminScope(cfg aws.Config, client *fms.Client) {
	input := &fms.GetAdminScopeInput{
		// AdminAccount: *string, // Required
	}

	if len(_fmsAdminAccount) > 0 {
		input.AdminAccount = aws.String(_fmsAdminAccount)
	}

	if resp, err := client.GetAdminScope(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the specified Firewall Manager applications list.
func fms_GetAppsList(cfg aws.Config, client *fms.Client) {
	input := &fms.GetAppsListInput{
		// ListId: *string, // Required
	}

	if len(_fmsListId) > 0 {
		input.ListId = aws.String(_fmsListId)
	}
	if len(_fmsDefaultList) > 0 {
		if err := assignInputField(input, "DefaultList", _fmsDefaultList); err != nil {
			log.Errorf("invalid --default-list: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetAppsList(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns detailed compliance information about the specified member account.
// Details include resources that are in and out of compliance with the specified
// policy.
//
// The reasons for resources being considered compliant depend on the Firewall
// Manager policy type.
func fms_GetComplianceDetail(cfg aws.Config, client *fms.Client) {
	input := &fms.GetComplianceDetailInput{
		// MemberAccount: *string, // Required
		// PolicyId: *string, // Required
	}

	if len(_fmsMemberAccount) > 0 {
		input.MemberAccount = aws.String(_fmsMemberAccount)
	}
	if len(_fmsPolicyId) > 0 {
		input.PolicyId = aws.String(_fmsPolicyId)
	}

	if resp, err := client.GetComplianceDetail(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Information about the Amazon Simple Notification Service (SNS) topic that is
// used to record Firewall Manager SNS logs.
func fms_GetNotificationChannel(cfg aws.Config, client *fms.Client) {
	input := &fms.GetNotificationChannelInput{}

	if resp, err := client.GetNotificationChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the specified Firewall Manager policy.
func fms_GetPolicy(cfg aws.Config, client *fms.Client) {
	input := &fms.GetPolicyInput{
		// PolicyId: *string, // Required
	}

	if len(_fmsPolicyId) > 0 {
		input.PolicyId = aws.String(_fmsPolicyId)
	}

	if resp, err := client.GetPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// If you created a Shield Advanced policy, returns policy-level attack summary
// information in the event of a potential DDoS attack. Other policy types are
// currently unsupported.
func fms_GetProtectionStatus(cfg aws.Config, client *fms.Client) {
	input := &fms.GetProtectionStatusInput{
		// PolicyId: *string, // Required
	}

	if len(_fmsPolicyId) > 0 {
		input.PolicyId = aws.String(_fmsPolicyId)
	}
	if len(_fmsEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _fmsEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_fmsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _fmsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_fmsMemberAccountId) > 0 {
		input.MemberAccountId = aws.String(_fmsMemberAccountId)
	}
	if len(_fmsNextToken) > 0 {
		input.NextToken = aws.String(_fmsNextToken)
	}
	if len(_fmsStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _fmsStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetProtectionStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the specified Firewall Manager protocols list.
func fms_GetProtocolsList(cfg aws.Config, client *fms.Client) {
	input := &fms.GetProtocolsListInput{
		// ListId: *string, // Required
	}

	if len(_fmsListId) > 0 {
		input.ListId = aws.String(_fmsListId)
	}
	if len(_fmsDefaultList) > 0 {
		if err := assignInputField(input, "DefaultList", _fmsDefaultList); err != nil {
			log.Errorf("invalid --default-list: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetProtocolsList(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a specific resource set.
func fms_GetResourceSet(cfg aws.Config, client *fms.Client) {
	input := &fms.GetResourceSetInput{
		// Identifier: *string, // Required
	}

	if len(_fmsIdentifier) > 0 {
		input.Identifier = aws.String(_fmsIdentifier)
	}

	if resp, err := client.GetResourceSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The onboarding status of a Firewall Manager admin account to third-party
// firewall vendor tenant.
func fms_GetThirdPartyFirewallAssociationStatus(cfg aws.Config, client *fms.Client) {
	input := &fms.GetThirdPartyFirewallAssociationStatusInput{
		// ThirdPartyFirewall: types.ThirdPartyFirewall, // Required
	}

	if len(_fmsThirdPartyFirewall) > 0 {
		if err := assignInputField(input, "ThirdPartyFirewall", _fmsThirdPartyFirewall); err != nil {
			log.Errorf("invalid --third-party-firewall: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetThirdPartyFirewallAssociationStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves violations for a resource based on the specified Firewall Manager
// policy and Amazon Web Services account.
func fms_GetViolationDetails(cfg aws.Config, client *fms.Client) {
	input := &fms.GetViolationDetailsInput{
		// MemberAccount: *string, // Required
		// PolicyId: *string, // Required
		// ResourceId: *string, // Required
		// ResourceType: *string, // Required
	}

	if len(_fmsMemberAccount) > 0 {
		input.MemberAccount = aws.String(_fmsMemberAccount)
	}
	if len(_fmsPolicyId) > 0 {
		input.PolicyId = aws.String(_fmsPolicyId)
	}
	if len(_fmsResourceId) > 0 {
		input.ResourceId = aws.String(_fmsResourceId)
	}
	if len(_fmsResourceType) > 0 {
		input.ResourceType = aws.String(_fmsResourceType)
	}

	if resp, err := client.GetViolationDetails(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a AdminAccounts object that lists the Firewall Manager administrators
// within the organization that are onboarded to Firewall Manager by AssociateAdminAccount.
//
// This operation can be called only from the organization's management account.
func fms_ListAdminAccountsForOrganization(cfg aws.Config, client *fms.Client) {
	input := &fms.ListAdminAccountsForOrganizationInput{}

	if len(_fmsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _fmsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_fmsNextToken) > 0 {
		input.NextToken = aws.String(_fmsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAdminAccountsForOrganization(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*fms.ListAdminAccountsForOrganizationOutput
	p := fms.NewListAdminAccountsForOrganizationPaginator(client, input)
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

// Lists the accounts that are managing the specified Organizations member
// account. This is useful for any member account so that they can view the
// accounts who are managing their account. This operation only returns the
// managing administrators that have the requested account within their AdminScope.
func fms_ListAdminsManagingAccount(cfg aws.Config, client *fms.Client) {
	input := &fms.ListAdminsManagingAccountInput{}

	if len(_fmsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _fmsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_fmsNextToken) > 0 {
		input.NextToken = aws.String(_fmsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAdminsManagingAccount(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*fms.ListAdminsManagingAccountOutput
	p := fms.NewListAdminsManagingAccountPaginator(client, input)
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

// Returns an array of AppsListDataSummary objects.
func fms_ListAppsLists(cfg aws.Config, client *fms.Client) {
	input := &fms.ListAppsListsInput{
		// MaxResults: *int32, // Required
	}

	if len(_fmsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _fmsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_fmsDefaultLists) > 0 {
		if err := assignInputField(input, "DefaultLists", _fmsDefaultLists); err != nil {
			log.Errorf("invalid --default-lists: %s", err.Error())
			return
		}
	}
	if len(_fmsNextToken) > 0 {
		input.NextToken = aws.String(_fmsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAppsLists(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*fms.ListAppsListsOutput
	p := fms.NewListAppsListsPaginator(client, input)
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

// Returns an array of PolicyComplianceStatus objects. Use PolicyComplianceStatus
// to get a summary of which member accounts are protected by the specified policy.
func fms_ListComplianceStatus(cfg aws.Config, client *fms.Client) {
	input := &fms.ListComplianceStatusInput{
		// PolicyId: *string, // Required
	}

	if len(_fmsPolicyId) > 0 {
		input.PolicyId = aws.String(_fmsPolicyId)
	}
	if len(_fmsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _fmsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_fmsNextToken) > 0 {
		input.NextToken = aws.String(_fmsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListComplianceStatus(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*fms.ListComplianceStatusOutput
	p := fms.NewListComplianceStatusPaginator(client, input)
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

// Returns an array of resources in the organization's accounts that are available
// to be associated with a resource set.
func fms_ListDiscoveredResources(cfg aws.Config, client *fms.Client) {
	input := &fms.ListDiscoveredResourcesInput{
		// MemberAccountIds: []string, // Required
		// ResourceType: *string, // Required
	}

	if len(_fmsMemberAccountIds) > 0 {
		input.MemberAccountIds = append([]string(nil), _fmsMemberAccountIds...)
	}
	if len(_fmsResourceType) > 0 {
		input.ResourceType = aws.String(_fmsResourceType)
	}
	if len(_fmsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _fmsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_fmsNextToken) > 0 {
		input.NextToken = aws.String(_fmsNextToken)
	}

	if resp, err := client.ListDiscoveredResources(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a MemberAccounts object that lists the member accounts in the
// administrator's Amazon Web Services organization.
//
// Either an Firewall Manager administrator or the organization's management
// account can make this request.
func fms_ListMemberAccounts(cfg aws.Config, client *fms.Client) {
	input := &fms.ListMemberAccountsInput{}

	if len(_fmsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _fmsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_fmsNextToken) > 0 {
		input.NextToken = aws.String(_fmsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListMemberAccounts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*fms.ListMemberAccountsOutput
	p := fms.NewListMemberAccountsPaginator(client, input)
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

// Returns an array of PolicySummary objects.
func fms_ListPolicies(cfg aws.Config, client *fms.Client) {
	input := &fms.ListPoliciesInput{}

	if len(_fmsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _fmsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_fmsNextToken) > 0 {
		input.NextToken = aws.String(_fmsNextToken)
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

	var results []*fms.ListPoliciesOutput
	p := fms.NewListPoliciesPaginator(client, input)
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

// Returns an array of ProtocolsListDataSummary objects.
func fms_ListProtocolsLists(cfg aws.Config, client *fms.Client) {
	input := &fms.ListProtocolsListsInput{
		// MaxResults: *int32, // Required
	}

	if len(_fmsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _fmsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_fmsDefaultLists) > 0 {
		if err := assignInputField(input, "DefaultLists", _fmsDefaultLists); err != nil {
			log.Errorf("invalid --default-lists: %s", err.Error())
			return
		}
	}
	if len(_fmsNextToken) > 0 {
		input.NextToken = aws.String(_fmsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListProtocolsLists(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*fms.ListProtocolsListsOutput
	p := fms.NewListProtocolsListsPaginator(client, input)
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

// Returns an array of resources that are currently associated to a resource set.
func fms_ListResourceSetResources(cfg aws.Config, client *fms.Client) {
	input := &fms.ListResourceSetResourcesInput{
		// Identifier: *string, // Required
	}

	if len(_fmsIdentifier) > 0 {
		input.Identifier = aws.String(_fmsIdentifier)
	}
	if len(_fmsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _fmsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_fmsNextToken) > 0 {
		input.NextToken = aws.String(_fmsNextToken)
	}

	if resp, err := client.ListResourceSetResources(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns an array of ResourceSetSummary objects.
func fms_ListResourceSets(cfg aws.Config, client *fms.Client) {
	input := &fms.ListResourceSetsInput{}

	if len(_fmsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _fmsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_fmsNextToken) > 0 {
		input.NextToken = aws.String(_fmsNextToken)
	}

	if resp, err := client.ListResourceSets(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the list of tags for the specified Amazon Web Services resource.
func fms_ListTagsForResource(cfg aws.Config, client *fms.Client) {
	input := &fms.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_fmsResourceArn) > 0 {
		input.ResourceArn = aws.String(_fmsResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list of all of the third-party firewall policies that are
// associated with the third-party firewall administrator's account.
func fms_ListThirdPartyFirewallFirewallPolicies(cfg aws.Config, client *fms.Client) {
	input := &fms.ListThirdPartyFirewallFirewallPoliciesInput{
		// MaxResults: *int32, // Required
		// ThirdPartyFirewall: types.ThirdPartyFirewall, // Required
	}

	if len(_fmsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _fmsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_fmsThirdPartyFirewall) > 0 {
		if err := assignInputField(input, "ThirdPartyFirewall", _fmsThirdPartyFirewall); err != nil {
			log.Errorf("invalid --third-party-firewall: %s", err.Error())
			return
		}
	}
	if len(_fmsNextToken) > 0 {
		input.NextToken = aws.String(_fmsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListThirdPartyFirewallFirewallPolicies(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*fms.ListThirdPartyFirewallFirewallPoliciesOutput
	p := fms.NewListThirdPartyFirewallFirewallPoliciesPaginator(client, input)
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

// Creates or updates an Firewall Manager administrator account. The account must
// be a member of the organization that was onboarded to Firewall Manager by AssociateAdminAccount.
// Only the organization's management account can create an Firewall Manager
// administrator account. When you create an Firewall Manager administrator
// account, the service checks to see if the account is already a delegated
// administrator within Organizations. If the account isn't a delegated
// administrator, Firewall Manager calls Organizations to delegate the account
// within Organizations. For more information about administrator accounts within
// Organizations, see [Managing the Amazon Web Services Accounts in Your Organization].
//
// [Managing the Amazon Web Services Accounts in Your Organization]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_accounts.html
func fms_PutAdminAccount(cfg aws.Config, client *fms.Client) {
	input := &fms.PutAdminAccountInput{
		// AdminAccount: *string, // Required
	}

	if len(_fmsAdminAccount) > 0 {
		input.AdminAccount = aws.String(_fmsAdminAccount)
	}
	if len(_fmsAdminScope) > 0 {
		if err := assignInputField(input, "AdminScope", _fmsAdminScope); err != nil {
			log.Errorf("invalid --admin-scope: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutAdminAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Firewall Manager applications list.
func fms_PutAppsList(cfg aws.Config, client *fms.Client) {
	input := &fms.PutAppsListInput{
		// AppsList: *types.AppsListData, // Required
	}

	if len(_fmsAppsList) > 0 {
		if err := assignInputField(input, "AppsList", _fmsAppsList); err != nil {
			log.Errorf("invalid --apps-list: %s", err.Error())
			return
		}
	}
	if len(_fmsTagList) > 0 {
		if err := assignInputField(input, "TagList", _fmsTagList); err != nil {
			log.Errorf("invalid --tag-list: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutAppsList(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Designates the IAM role and Amazon Simple Notification Service (SNS) topic that
// Firewall Manager uses to record SNS logs.
//
// To perform this action outside of the console, you must first configure the SNS
// topic's access policy to allow the SnsRoleName to publish SNS logs. If the
// SnsRoleName provided is a role other than the AWSServiceRoleForFMS
// service-linked role, this role must have a trust relationship configured to
// allow the Firewall Manager service principal fms.amazonaws.com to assume this
// role. For information about configuring an SNS access policy, see [Service roles for Firewall Manager]in the
// Firewall Manager Developer Guide.
//
// [Service roles for Firewall Manager]: https://docs.aws.amazon.com/waf/latest/developerguide/fms-security_iam_service-with-iam.html#fms-security_iam_service-with-iam-roles-service
func fms_PutNotificationChannel(cfg aws.Config, client *fms.Client) {
	input := &fms.PutNotificationChannelInput{
		// SnsRoleName: *string, // Required
		// SnsTopicArn: *string, // Required
	}

	if len(_fmsSnsRoleName) > 0 {
		input.SnsRoleName = aws.String(_fmsSnsRoleName)
	}
	if len(_fmsSnsTopicArn) > 0 {
		input.SnsTopicArn = aws.String(_fmsSnsTopicArn)
	}

	if resp, err := client.PutNotificationChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Firewall Manager policy.
// A Firewall Manager policy is specific to the individual policy type. If you
// want to enforce multiple policy types across accounts, you can create multiple
// policies. You can create more than one policy for each type.
//
// If you add a new account to an organization that you created with
// Organizations, Firewall Manager automatically applies the policy to the
// resources in that account that are within scope of the policy.
//
// Firewall Manager provides the following types of policies:
//
// - WAF policy - This policy applies WAF web ACL protections to specified
// accounts and resources.
//
// - Shield Advanced policy - This policy applies Shield Advanced protection to
// specified accounts and resources.
//
// - Security Groups policy - This type of policy gives you control over
// security groups that are in use throughout your organization in Organizations
// and lets you enforce a baseline set of rules across your organization.
//
// - Network ACL policy - This type of policy gives you control over the network
// ACLs that are in use throughout your organization in Organizations and lets you
// enforce a baseline set of first and last network ACL rules across your
// organization.
//
// - Network Firewall policy - This policy applies Network Firewall protection
// to your organization's VPCs.
//
// - DNS Firewall policy - This policy applies Amazon Route 53 Resolver DNS
// Firewall protections to your organization's VPCs.
//
// - Third-party firewall policy - This policy applies third-party firewall
// protections. Third-party firewalls are available by subscription through the
// Amazon Web Services Marketplace console at [Amazon Web Services Marketplace].
//
// - Palo Alto Networks Cloud NGFW policy - This policy applies Palo Alto
// Networks Cloud Next Generation Firewall (NGFW) protections and Palo Alto
// Networks Cloud NGFW rulestacks to your organization's VPCs.
//
// - Fortigate CNF policy - This policy applies Fortigate Cloud Native Firewall
// (CNF) protections. Fortigate CNF is a cloud-centered solution that blocks
// Zero-Day threats and secures cloud infrastructures with industry-leading
// advanced threat prevention, smart web application firewalls (WAF), and API
// protection.
//
// [Amazon Web Services Marketplace]: http://aws.amazon.com/marketplace
func fms_PutPolicy(cfg aws.Config, client *fms.Client) {
	input := &fms.PutPolicyInput{
		// Policy: *types.Policy, // Required
	}

	if len(_fmsPolicy) > 0 {
		if err := assignInputField(input, "Policy", _fmsPolicy); err != nil {
			log.Errorf("invalid --policy: %s", err.Error())
			return
		}
	}
	if len(_fmsTagList) > 0 {
		if err := assignInputField(input, "TagList", _fmsTagList); err != nil {
			log.Errorf("invalid --tag-list: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Firewall Manager protocols list.
func fms_PutProtocolsList(cfg aws.Config, client *fms.Client) {
	input := &fms.PutProtocolsListInput{
		// ProtocolsList: *types.ProtocolsListData, // Required
	}

	if len(_fmsProtocolsList) > 0 {
		if err := assignInputField(input, "ProtocolsList", _fmsProtocolsList); err != nil {
			log.Errorf("invalid --protocols-list: %s", err.Error())
			return
		}
	}
	if len(_fmsTagList) > 0 {
		if err := assignInputField(input, "TagList", _fmsTagList); err != nil {
			log.Errorf("invalid --tag-list: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutProtocolsList(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates the resource set.
// An Firewall Manager resource set defines the resources to import into an
// Firewall Manager policy from another Amazon Web Services service.
func fms_PutResourceSet(cfg aws.Config, client *fms.Client) {
	input := &fms.PutResourceSetInput{
		// ResourceSet: *types.ResourceSet, // Required
	}

	if len(_fmsResourceSet) > 0 {
		if err := assignInputField(input, "ResourceSet", _fmsResourceSet); err != nil {
			log.Errorf("invalid --resource-set: %s", err.Error())
			return
		}
	}
	if len(_fmsTagList) > 0 {
		if err := assignInputField(input, "TagList", _fmsTagList); err != nil {
			log.Errorf("invalid --tag-list: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutResourceSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds one or more tags to an Amazon Web Services resource.
func fms_TagResource(cfg aws.Config, client *fms.Client) {
	input := &fms.TagResourceInput{
		// ResourceArn: *string, // Required
		// TagList: []types.Tag, // Required
	}

	if len(_fmsResourceArn) > 0 {
		input.ResourceArn = aws.String(_fmsResourceArn)
	}
	if len(_fmsTagList) > 0 {
		if err := assignInputField(input, "TagList", _fmsTagList); err != nil {
			log.Errorf("invalid --tag-list: %s", err.Error())
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

// Removes one or more tags from an Amazon Web Services resource.
func fms_UntagResource(cfg aws.Config, client *fms.Client) {
	input := &fms.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_fmsResourceArn) > 0 {
		input.ResourceArn = aws.String(_fmsResourceArn)
	}
	if len(_fmsTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _fmsTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_fmsCmd)
	_fmsCmd.Flags().SortFlags = false

	_fmsCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_fmsCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_fmsCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_fmsCmd.Flags().StringVarP(&_fmsAdminAccount, "admin-account", "", "", "Admin Account")
	_fmsCmd.Flags().StringVarP(&_fmsAdminScope, "admin-scope", "", "", "Admin Scope")
	_fmsCmd.Flags().StringVarP(&_fmsAppsList, "apps-list", "", "", "Apps List")
	_fmsCmd.Flags().StringVarP(&_fmsDefaultList, "default-list", "", "", "Default List")
	_fmsCmd.Flags().StringVarP(&_fmsDefaultLists, "default-lists", "", "", "Default Lists")
	_fmsCmd.Flags().StringVarP(&_fmsDeleteAllPolicyResources, "delete-all-policy-resources", "", "", "Delete All Policy Resources")
	_fmsCmd.Flags().StringVarP(&_fmsEndTime, "end-time", "", "", "End Time")
	_fmsCmd.Flags().StringVarP(&_fmsIdentifier, "identifier", "", "", "Identifier")
	_fmsCmd.Flags().StringSliceVarP(&_fmsItems, "items", "", nil, "Items")
	_fmsCmd.Flags().StringVarP(&_fmsListId, "list-id", "", "", "List ID")
	_fmsCmd.Flags().StringVarP(&_fmsMaxResults, "max-results", "", "", "Max Results")
	_fmsCmd.Flags().StringVarP(&_fmsMemberAccount, "member-account", "", "", "Member Account")
	_fmsCmd.Flags().StringVarP(&_fmsMemberAccountId, "member-account-id", "", "", "Member Account ID")
	_fmsCmd.Flags().StringSliceVarP(&_fmsMemberAccountIds, "member-account-ids", "", nil, "Member Account Ids")
	_fmsCmd.Flags().StringVarP(&_fmsNextToken, "next-token", "", "", "Next Token")
	_fmsCmd.Flags().StringVarP(&_fmsPolicy, "policy", "", "", "Policy")
	_fmsCmd.Flags().StringVarP(&_fmsPolicyId, "policy-id", "", "", "Policy ID")
	_fmsCmd.Flags().StringVarP(&_fmsProtocolsList, "protocols-list", "", "", "Protocols List")
	_fmsCmd.Flags().StringVarP(&_fmsResourceArn, "resource-arn", "", "", "Resource ARN")
	_fmsCmd.Flags().StringVarP(&_fmsResourceId, "resource-id", "", "", "Resource ID")
	_fmsCmd.Flags().StringVarP(&_fmsResourceSet, "resource-set", "", "", "Resource Set")
	_fmsCmd.Flags().StringVarP(&_fmsResourceSetIdentifier, "resource-set-identifier", "", "", "Resource Set Identifier")
	_fmsCmd.Flags().StringVarP(&_fmsResourceType, "resource-type", "", "", "Resource Type")
	_fmsCmd.Flags().StringVarP(&_fmsSnsRoleName, "sns-role-name", "", "", "SNS Role Name")
	_fmsCmd.Flags().StringVarP(&_fmsSnsTopicArn, "sns-topic-arn", "", "", "SNS Topic ARN")
	_fmsCmd.Flags().StringVarP(&_fmsStartTime, "start-time", "", "", "Start Time")
	_fmsCmd.Flags().StringSliceVarP(&_fmsTagKeys, "tag-keys", "", nil, "Tag Keys")
	_fmsCmd.Flags().StringVarP(&_fmsTagList, "tag-list", "", "", "Tag List")
	_fmsCmd.Flags().StringVarP(&_fmsThirdPartyFirewall, "third-party-firewall", "", "", "Third Party Firewall")

	_fmsCmd.Flags().BoolVarP(&_fmsAssociateAdminAccount, "associate-admin-account", "", false, "Associate Admin Account")
	_fmsCmd.Flags().BoolVarP(&_fmsAssociateThirdPartyFirewall, "associate-third-party-firewall", "", false, "Associate Third Party Firewall")
	_fmsCmd.Flags().BoolVarP(&_fmsBatchAssociateResource, "batch-associate-resource", "", false, "Batch Associate Resource")
	_fmsCmd.Flags().BoolVarP(&_fmsBatchDisassociateResource, "batch-disassociate-resource", "", false, "Batch Disassociate Resource")
	_fmsCmd.Flags().BoolVarP(&_fmsDeleteAppsList, "delete-apps-list", "", false, "Delete Apps List")
	_fmsCmd.Flags().BoolVarP(&_fmsDeleteNotificationChannel, "delete-notification-channel", "", false, "Delete Notification Channel")
	_fmsCmd.Flags().BoolVarP(&_fmsDeletePolicy, "delete-policy", "", false, "Delete Policy")
	_fmsCmd.Flags().BoolVarP(&_fmsDeleteProtocolsList, "delete-protocols-list", "", false, "Delete Protocols List")
	_fmsCmd.Flags().BoolVarP(&_fmsDeleteResourceSet, "delete-resource-set", "", false, "Delete Resource Set")
	_fmsCmd.Flags().BoolVarP(&_fmsDisassociateAdminAccount, "disassociate-admin-account", "", false, "Disassociate Admin Account")
	_fmsCmd.Flags().BoolVarP(&_fmsDisassociateThirdPartyFirewall, "disassociate-third-party-firewall", "", false, "Disassociate Third Party Firewall")
	_fmsCmd.Flags().BoolVarP(&_fmsGetAdminAccount, "get-admin-account", "", false, "Get Admin Account")
	_fmsCmd.Flags().BoolVarP(&_fmsGetAdminScope, "get-admin-scope", "", false, "Get Admin Scope")
	_fmsCmd.Flags().BoolVarP(&_fmsGetAppsList, "get-apps-list", "", false, "Get Apps List")
	_fmsCmd.Flags().BoolVarP(&_fmsGetComplianceDetail, "get-compliance-detail", "", false, "Get Compliance Detail")
	_fmsCmd.Flags().BoolVarP(&_fmsGetNotificationChannel, "get-notification-channel", "", false, "Get Notification Channel")
	_fmsCmd.Flags().BoolVarP(&_fmsGetPolicy, "get-policy", "", false, "Get Policy")
	_fmsCmd.Flags().BoolVarP(&_fmsGetProtectionStatus, "get-protection-status", "", false, "Get Protection Status")
	_fmsCmd.Flags().BoolVarP(&_fmsGetProtocolsList, "get-protocols-list", "", false, "Get Protocols List")
	_fmsCmd.Flags().BoolVarP(&_fmsGetResourceSet, "get-resource-set", "", false, "Get Resource Set")
	_fmsCmd.Flags().BoolVarP(&_fmsGetThirdPartyFirewallAssociationStatus, "get-third-party-firewall-association-status", "", false, "Get Third Party Firewall Association Status")
	_fmsCmd.Flags().BoolVarP(&_fmsGetViolationDetails, "get-violation-details", "", false, "Get Violation Details")
	_fmsCmd.Flags().BoolVarP(&_fmsListAdminAccountsForOrganization, "list-admin-accounts-for-organization", "", false, "List Admin Accounts For Organization")
	_fmsCmd.Flags().BoolVarP(&_fmsListAdminsManagingAccount, "list-admins-managing-account", "", false, "List Admins Managing Account")
	_fmsCmd.Flags().BoolVarP(&_fmsListAppsLists, "list-apps-lists", "", false, "List Apps Lists")
	_fmsCmd.Flags().BoolVarP(&_fmsListComplianceStatus, "list-compliance-status", "", false, "List Compliance Status")
	_fmsCmd.Flags().BoolVarP(&_fmsListDiscoveredResources, "list-discovered-resources", "", false, "List Discovered Resources")
	_fmsCmd.Flags().BoolVarP(&_fmsListMemberAccounts, "list-member-accounts", "", false, "List Member Accounts")
	_fmsCmd.Flags().BoolVarP(&_fmsListPolicies, "list-policies", "", false, "List Policies")
	_fmsCmd.Flags().BoolVarP(&_fmsListProtocolsLists, "list-protocols-lists", "", false, "List Protocols Lists")
	_fmsCmd.Flags().BoolVarP(&_fmsListResourceSetResources, "list-resource-set-resources", "", false, "List Resource Set Resources")
	_fmsCmd.Flags().BoolVarP(&_fmsListResourceSets, "list-resource-sets", "", false, "List Resource Sets")
	_fmsCmd.Flags().BoolVarP(&_fmsListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_fmsCmd.Flags().BoolVarP(&_fmsListThirdPartyFirewallFirewallPolicies, "list-third-party-firewall-firewall-policies", "", false, "List Third Party Firewall Firewall Policies")
	_fmsCmd.Flags().BoolVarP(&_fmsPutAdminAccount, "put-admin-account", "", false, "Put Admin Account")
	_fmsCmd.Flags().BoolVarP(&_fmsPutAppsList, "put-apps-list", "", false, "Put Apps List")
	_fmsCmd.Flags().BoolVarP(&_fmsPutNotificationChannel, "put-notification-channel", "", false, "Put Notification Channel")
	_fmsCmd.Flags().BoolVarP(&_fmsPutPolicy, "put-policy", "", false, "Put Policy")
	_fmsCmd.Flags().BoolVarP(&_fmsPutProtocolsList, "put-protocols-list", "", false, "Put Protocols List")
	_fmsCmd.Flags().BoolVarP(&_fmsPutResourceSet, "put-resource-set", "", false, "Put Resource Set")
	_fmsCmd.Flags().BoolVarP(&_fmsTagResource, "tag-resource", "", false, "Tag Resource")
	_fmsCmd.Flags().BoolVarP(&_fmsUntagResource, "untag-resource", "", false, "Untag Resource")

}
