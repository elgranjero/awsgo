package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/securityhub"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// securityhubCmd represents the securityhub command
var _securityhubCmd = &cobra.Command{
	Use:   "securityhub",
	Short: "AWS securityhub CLI",
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
		client := securityhub.NewFromConfig(cfg)
		if _securityhubAcceptAdministratorInvitation {
			securityhub_AcceptAdministratorInvitation(cfg, client)
			return
		}
		if _securityhubAcceptInvitation {
			securityhub_AcceptInvitation(cfg, client)
			return
		}
		if _securityhubBatchDeleteAutomationRules {
			securityhub_BatchDeleteAutomationRules(cfg, client)
			return
		}
		if _securityhubBatchDisableStandards {
			securityhub_BatchDisableStandards(cfg, client)
			return
		}
		if _securityhubBatchEnableStandards {
			securityhub_BatchEnableStandards(cfg, client)
			return
		}
		if _securityhubBatchGetAutomationRules {
			securityhub_BatchGetAutomationRules(cfg, client)
			return
		}
		if _securityhubBatchGetConfigurationPolicyAssociations {
			securityhub_BatchGetConfigurationPolicyAssociations(cfg, client)
			return
		}
		if _securityhubBatchGetSecurityControls {
			securityhub_BatchGetSecurityControls(cfg, client)
			return
		}
		if _securityhubBatchGetStandardsControlAssociations {
			securityhub_BatchGetStandardsControlAssociations(cfg, client)
			return
		}
		if _securityhubBatchImportFindings {
			securityhub_BatchImportFindings(cfg, client)
			return
		}
		if _securityhubBatchUpdateAutomationRules {
			securityhub_BatchUpdateAutomationRules(cfg, client)
			return
		}
		if _securityhubBatchUpdateFindings {
			securityhub_BatchUpdateFindings(cfg, client)
			return
		}
		if _securityhubBatchUpdateFindingsV2 {
			securityhub_BatchUpdateFindingsV2(cfg, client)
			return
		}
		if _securityhubBatchUpdateStandardsControlAssociations {
			securityhub_BatchUpdateStandardsControlAssociations(cfg, client)
			return
		}
		if _securityhubCreateActionTarget {
			securityhub_CreateActionTarget(cfg, client)
			return
		}
		if _securityhubCreateAggregatorV2 {
			securityhub_CreateAggregatorV2(cfg, client)
			return
		}
		if _securityhubCreateAutomationRule {
			securityhub_CreateAutomationRule(cfg, client)
			return
		}
		if _securityhubCreateAutomationRuleV2 {
			securityhub_CreateAutomationRuleV2(cfg, client)
			return
		}
		if _securityhubCreateConfigurationPolicy {
			securityhub_CreateConfigurationPolicy(cfg, client)
			return
		}
		if _securityhubCreateConnectorV2 {
			securityhub_CreateConnectorV2(cfg, client)
			return
		}
		if _securityhubCreateFindingAggregator {
			securityhub_CreateFindingAggregator(cfg, client)
			return
		}
		if _securityhubCreateInsight {
			securityhub_CreateInsight(cfg, client)
			return
		}
		if _securityhubCreateMembers {
			securityhub_CreateMembers(cfg, client)
			return
		}
		if _securityhubCreateTicketV2 {
			securityhub_CreateTicketV2(cfg, client)
			return
		}
		if _securityhubDeclineInvitations {
			securityhub_DeclineInvitations(cfg, client)
			return
		}
		if _securityhubDeleteActionTarget {
			securityhub_DeleteActionTarget(cfg, client)
			return
		}
		if _securityhubDeleteAggregatorV2 {
			securityhub_DeleteAggregatorV2(cfg, client)
			return
		}
		if _securityhubDeleteAutomationRuleV2 {
			securityhub_DeleteAutomationRuleV2(cfg, client)
			return
		}
		if _securityhubDeleteConfigurationPolicy {
			securityhub_DeleteConfigurationPolicy(cfg, client)
			return
		}
		if _securityhubDeleteConnectorV2 {
			securityhub_DeleteConnectorV2(cfg, client)
			return
		}
		if _securityhubDeleteFindingAggregator {
			securityhub_DeleteFindingAggregator(cfg, client)
			return
		}
		if _securityhubDeleteInsight {
			securityhub_DeleteInsight(cfg, client)
			return
		}
		if _securityhubDeleteInvitations {
			securityhub_DeleteInvitations(cfg, client)
			return
		}
		if _securityhubDeleteMembers {
			securityhub_DeleteMembers(cfg, client)
			return
		}
		if _securityhubDescribeActionTargets {
			securityhub_DescribeActionTargets(cfg, client)
			return
		}
		if _securityhubDescribeHub {
			securityhub_DescribeHub(cfg, client)
			return
		}
		if _securityhubDescribeOrganizationConfiguration {
			securityhub_DescribeOrganizationConfiguration(cfg, client)
			return
		}
		if _securityhubDescribeProducts {
			securityhub_DescribeProducts(cfg, client)
			return
		}
		if _securityhubDescribeProductsV2 {
			securityhub_DescribeProductsV2(cfg, client)
			return
		}
		if _securityhubDescribeSecurityHubV2 {
			securityhub_DescribeSecurityHubV2(cfg, client)
			return
		}
		if _securityhubDescribeStandards {
			securityhub_DescribeStandards(cfg, client)
			return
		}
		if _securityhubDescribeStandardsControls {
			securityhub_DescribeStandardsControls(cfg, client)
			return
		}
		if _securityhubDisableImportFindingsForProduct {
			securityhub_DisableImportFindingsForProduct(cfg, client)
			return
		}
		if _securityhubDisableOrganizationAdminAccount {
			securityhub_DisableOrganizationAdminAccount(cfg, client)
			return
		}
		if _securityhubDisableSecurityHub {
			securityhub_DisableSecurityHub(cfg, client)
			return
		}
		if _securityhubDisableSecurityHubV2 {
			securityhub_DisableSecurityHubV2(cfg, client)
			return
		}
		if _securityhubDisassociateFromAdministratorAccount {
			securityhub_DisassociateFromAdministratorAccount(cfg, client)
			return
		}
		if _securityhubDisassociateFromMasterAccount {
			securityhub_DisassociateFromMasterAccount(cfg, client)
			return
		}
		if _securityhubDisassociateMembers {
			securityhub_DisassociateMembers(cfg, client)
			return
		}
		if _securityhubEnableImportFindingsForProduct {
			securityhub_EnableImportFindingsForProduct(cfg, client)
			return
		}
		if _securityhubEnableOrganizationAdminAccount {
			securityhub_EnableOrganizationAdminAccount(cfg, client)
			return
		}
		if _securityhubEnableSecurityHub {
			securityhub_EnableSecurityHub(cfg, client)
			return
		}
		if _securityhubEnableSecurityHubV2 {
			securityhub_EnableSecurityHubV2(cfg, client)
			return
		}
		if _securityhubGetAdministratorAccount {
			securityhub_GetAdministratorAccount(cfg, client)
			return
		}
		if _securityhubGetAggregatorV2 {
			securityhub_GetAggregatorV2(cfg, client)
			return
		}
		if _securityhubGetAutomationRuleV2 {
			securityhub_GetAutomationRuleV2(cfg, client)
			return
		}
		if _securityhubGetConfigurationPolicy {
			securityhub_GetConfigurationPolicy(cfg, client)
			return
		}
		if _securityhubGetConfigurationPolicyAssociation {
			securityhub_GetConfigurationPolicyAssociation(cfg, client)
			return
		}
		if _securityhubGetConnectorV2 {
			securityhub_GetConnectorV2(cfg, client)
			return
		}
		if _securityhubGetEnabledStandards {
			securityhub_GetEnabledStandards(cfg, client)
			return
		}
		if _securityhubGetFindingAggregator {
			securityhub_GetFindingAggregator(cfg, client)
			return
		}
		if _securityhubGetFindingHistory {
			securityhub_GetFindingHistory(cfg, client)
			return
		}
		if _securityhubGetFindingStatisticsV2 {
			securityhub_GetFindingStatisticsV2(cfg, client)
			return
		}
		if _securityhubGetFindings {
			securityhub_GetFindings(cfg, client)
			return
		}
		if _securityhubGetFindingsTrendsV2 {
			securityhub_GetFindingsTrendsV2(cfg, client)
			return
		}
		if _securityhubGetFindingsV2 {
			securityhub_GetFindingsV2(cfg, client)
			return
		}
		if _securityhubGetInsightResults {
			securityhub_GetInsightResults(cfg, client)
			return
		}
		if _securityhubGetInsights {
			securityhub_GetInsights(cfg, client)
			return
		}
		if _securityhubGetInvitationsCount {
			securityhub_GetInvitationsCount(cfg, client)
			return
		}
		if _securityhubGetMasterAccount {
			securityhub_GetMasterAccount(cfg, client)
			return
		}
		if _securityhubGetMembers {
			securityhub_GetMembers(cfg, client)
			return
		}
		if _securityhubGetResourcesStatisticsV2 {
			securityhub_GetResourcesStatisticsV2(cfg, client)
			return
		}
		if _securityhubGetResourcesTrendsV2 {
			securityhub_GetResourcesTrendsV2(cfg, client)
			return
		}
		if _securityhubGetResourcesV2 {
			securityhub_GetResourcesV2(cfg, client)
			return
		}
		if _securityhubGetSecurityControlDefinition {
			securityhub_GetSecurityControlDefinition(cfg, client)
			return
		}
		if _securityhubInviteMembers {
			securityhub_InviteMembers(cfg, client)
			return
		}
		if _securityhubListAggregatorsV2 {
			securityhub_ListAggregatorsV2(cfg, client)
			return
		}
		if _securityhubListAutomationRules {
			securityhub_ListAutomationRules(cfg, client)
			return
		}
		if _securityhubListAutomationRulesV2 {
			securityhub_ListAutomationRulesV2(cfg, client)
			return
		}
		if _securityhubListConfigurationPolicies {
			securityhub_ListConfigurationPolicies(cfg, client)
			return
		}
		if _securityhubListConfigurationPolicyAssociations {
			securityhub_ListConfigurationPolicyAssociations(cfg, client)
			return
		}
		if _securityhubListConnectorsV2 {
			securityhub_ListConnectorsV2(cfg, client)
			return
		}
		if _securityhubListEnabledProductsForImport {
			securityhub_ListEnabledProductsForImport(cfg, client)
			return
		}
		if _securityhubListFindingAggregators {
			securityhub_ListFindingAggregators(cfg, client)
			return
		}
		if _securityhubListInvitations {
			securityhub_ListInvitations(cfg, client)
			return
		}
		if _securityhubListMembers {
			securityhub_ListMembers(cfg, client)
			return
		}
		if _securityhubListOrganizationAdminAccounts {
			securityhub_ListOrganizationAdminAccounts(cfg, client)
			return
		}
		if _securityhubListSecurityControlDefinitions {
			securityhub_ListSecurityControlDefinitions(cfg, client)
			return
		}
		if _securityhubListStandardsControlAssociations {
			securityhub_ListStandardsControlAssociations(cfg, client)
			return
		}
		if _securityhubListTagsForResource {
			securityhub_ListTagsForResource(cfg, client)
			return
		}
		if _securityhubRegisterConnectorV2 {
			securityhub_RegisterConnectorV2(cfg, client)
			return
		}
		if _securityhubStartConfigurationPolicyAssociation {
			securityhub_StartConfigurationPolicyAssociation(cfg, client)
			return
		}
		if _securityhubStartConfigurationPolicyDisassociation {
			securityhub_StartConfigurationPolicyDisassociation(cfg, client)
			return
		}
		if _securityhubTagResource {
			securityhub_TagResource(cfg, client)
			return
		}
		if _securityhubUntagResource {
			securityhub_UntagResource(cfg, client)
			return
		}
		if _securityhubUpdateActionTarget {
			securityhub_UpdateActionTarget(cfg, client)
			return
		}
		if _securityhubUpdateAggregatorV2 {
			securityhub_UpdateAggregatorV2(cfg, client)
			return
		}
		if _securityhubUpdateAutomationRuleV2 {
			securityhub_UpdateAutomationRuleV2(cfg, client)
			return
		}
		if _securityhubUpdateConfigurationPolicy {
			securityhub_UpdateConfigurationPolicy(cfg, client)
			return
		}
		if _securityhubUpdateConnectorV2 {
			securityhub_UpdateConnectorV2(cfg, client)
			return
		}
		if _securityhubUpdateFindingAggregator {
			securityhub_UpdateFindingAggregator(cfg, client)
			return
		}
		if _securityhubUpdateFindings {
			securityhub_UpdateFindings(cfg, client)
			return
		}
		if _securityhubUpdateInsight {
			securityhub_UpdateInsight(cfg, client)
			return
		}
		if _securityhubUpdateOrganizationConfiguration {
			securityhub_UpdateOrganizationConfiguration(cfg, client)
			return
		}
		if _securityhubUpdateSecurityControl {
			securityhub_UpdateSecurityControl(cfg, client)
			return
		}
		if _securityhubUpdateSecurityHubConfiguration {
			securityhub_UpdateSecurityHubConfiguration(cfg, client)
			return
		}
		if _securityhubUpdateStandardsControl {
			securityhub_UpdateStandardsControl(cfg, client)
			return
		}

	},
}

var (
	_securityhubAcceptAdministratorInvitation           bool
	_securityhubAcceptInvitation                        bool
	_securityhubBatchDeleteAutomationRules              bool
	_securityhubBatchDisableStandards                   bool
	_securityhubBatchEnableStandards                    bool
	_securityhubBatchGetAutomationRules                 bool
	_securityhubBatchGetConfigurationPolicyAssociations bool
	_securityhubBatchGetSecurityControls                bool
	_securityhubBatchGetStandardsControlAssociations    bool
	_securityhubBatchImportFindings                     bool
	_securityhubBatchUpdateAutomationRules              bool
	_securityhubBatchUpdateFindings                     bool
	_securityhubBatchUpdateFindingsV2                   bool
	_securityhubBatchUpdateStandardsControlAssociations bool
	_securityhubCreateActionTarget                      bool
	_securityhubCreateAggregatorV2                      bool
	_securityhubCreateAutomationRule                    bool
	_securityhubCreateAutomationRuleV2                  bool
	_securityhubCreateConfigurationPolicy               bool
	_securityhubCreateConnectorV2                       bool
	_securityhubCreateFindingAggregator                 bool
	_securityhubCreateInsight                           bool
	_securityhubCreateMembers                           bool
	_securityhubCreateTicketV2                          bool
	_securityhubDeclineInvitations                      bool
	_securityhubDeleteActionTarget                      bool
	_securityhubDeleteAggregatorV2                      bool
	_securityhubDeleteAutomationRuleV2                  bool
	_securityhubDeleteConfigurationPolicy               bool
	_securityhubDeleteConnectorV2                       bool
	_securityhubDeleteFindingAggregator                 bool
	_securityhubDeleteInsight                           bool
	_securityhubDeleteInvitations                       bool
	_securityhubDeleteMembers                           bool
	_securityhubDescribeActionTargets                   bool
	_securityhubDescribeHub                             bool
	_securityhubDescribeOrganizationConfiguration       bool
	_securityhubDescribeProducts                        bool
	_securityhubDescribeProductsV2                      bool
	_securityhubDescribeSecurityHubV2                   bool
	_securityhubDescribeStandards                       bool
	_securityhubDescribeStandardsControls               bool
	_securityhubDisableImportFindingsForProduct         bool
	_securityhubDisableOrganizationAdminAccount         bool
	_securityhubDisableSecurityHub                      bool
	_securityhubDisableSecurityHubV2                    bool
	_securityhubDisassociateFromAdministratorAccount    bool
	_securityhubDisassociateFromMasterAccount           bool
	_securityhubDisassociateMembers                     bool
	_securityhubEnableImportFindingsForProduct          bool
	_securityhubEnableOrganizationAdminAccount          bool
	_securityhubEnableSecurityHub                       bool
	_securityhubEnableSecurityHubV2                     bool
	_securityhubGetAdministratorAccount                 bool
	_securityhubGetAggregatorV2                         bool
	_securityhubGetAutomationRuleV2                     bool
	_securityhubGetConfigurationPolicy                  bool
	_securityhubGetConfigurationPolicyAssociation       bool
	_securityhubGetConnectorV2                          bool
	_securityhubGetEnabledStandards                     bool
	_securityhubGetFindingAggregator                    bool
	_securityhubGetFindingHistory                       bool
	_securityhubGetFindingStatisticsV2                  bool
	_securityhubGetFindings                             bool
	_securityhubGetFindingsTrendsV2                     bool
	_securityhubGetFindingsV2                           bool
	_securityhubGetInsightResults                       bool
	_securityhubGetInsights                             bool
	_securityhubGetInvitationsCount                     bool
	_securityhubGetMasterAccount                        bool
	_securityhubGetMembers                              bool
	_securityhubGetResourcesStatisticsV2                bool
	_securityhubGetResourcesTrendsV2                    bool
	_securityhubGetResourcesV2                          bool
	_securityhubGetSecurityControlDefinition            bool
	_securityhubInviteMembers                           bool
	_securityhubListAggregatorsV2                       bool
	_securityhubListAutomationRules                     bool
	_securityhubListAutomationRulesV2                   bool
	_securityhubListConfigurationPolicies               bool
	_securityhubListConfigurationPolicyAssociations     bool
	_securityhubListConnectorsV2                        bool
	_securityhubListEnabledProductsForImport            bool
	_securityhubListFindingAggregators                  bool
	_securityhubListInvitations                         bool
	_securityhubListMembers                             bool
	_securityhubListOrganizationAdminAccounts           bool
	_securityhubListSecurityControlDefinitions          bool
	_securityhubListStandardsControlAssociations        bool
	_securityhubListTagsForResource                     bool
	_securityhubRegisterConnectorV2                     bool
	_securityhubStartConfigurationPolicyAssociation     bool
	_securityhubStartConfigurationPolicyDisassociation  bool
	_securityhubTagResource                             bool
	_securityhubUntagResource                           bool
	_securityhubUpdateActionTarget                      bool
	_securityhubUpdateAggregatorV2                      bool
	_securityhubUpdateAutomationRuleV2                  bool
	_securityhubUpdateConfigurationPolicy               bool
	_securityhubUpdateConnectorV2                       bool
	_securityhubUpdateFindingAggregator                 bool
	_securityhubUpdateFindings                          bool
	_securityhubUpdateInsight                           bool
	_securityhubUpdateOrganizationConfiguration         bool
	_securityhubUpdateSecurityControl                   bool
	_securityhubUpdateSecurityHubConfiguration          bool
	_securityhubUpdateStandardsControl                  bool

	_securityhubAccountDetails                            string
	_securityhubAccountIds                                []string
	_securityhubActionTargetArn                           string
	_securityhubActionTargetArns                          []string
	_securityhubActions                                   string
	_securityhubAdminAccountId                            string
	_securityhubAdministratorId                           string
	_securityhubAggregatorV2Arn                           string
	_securityhubAuthCode                                  string
	_securityhubAuthState                                 string
	_securityhubAutoEnable                                string
	_securityhubAutoEnableControls                        string
	_securityhubAutoEnableStandards                       string
	_securityhubAutomationRulesArns                       []string
	_securityhubClientToken                               string
	_securityhubComment                                   string
	_securityhubConfidence                                string
	_securityhubConfigurationPolicy                       string
	_securityhubConfigurationPolicyAssociationIdentifiers string
	_securityhubConfigurationPolicyIdentifier             string
	_securityhubConnectorId                               string
	_securityhubConnectorStatus                           string
	_securityhubControlFindingGenerator                   string
	_securityhubControlStatus                             string
	_securityhubCriteria                                  string
	_securityhubCriticality                               string
	_securityhubDescription                               string
	_securityhubDisabledReason                            string
	_securityhubEnableDefaultStandards                    string
	_securityhubEndTime                                   string
	_securityhubFeature                                   string
	_securityhubFilters                                   string
	_securityhubFindingAggregatorArn                      string
	_securityhubFindingIdentifier                         string
	_securityhubFindingIdentifiers                        string
	_securityhubFindingMetadataUid                        string
	_securityhubFindings                                  string
	_securityhubGroupByAttribute                          string
	_securityhubGroupByRules                              string
	_securityhubHubArn                                    string
	_securityhubId                                        string
	_securityhubIdentifier                                string
	_securityhubInsightArn                                string
	_securityhubInsightArns                               []string
	_securityhubInvitationId                              string
	_securityhubIsTerminal                                string
	_securityhubKmsKeyArn                                 string
	_securityhubLastUpdateReason                          string
	_securityhubLinkedRegions                             []string
	_securityhubMasterId                                  string
	_securityhubMaxResults                                string
	_securityhubMaxStatisticResults                       string
	_securityhubMetadataUids                              []string
	_securityhubMode                                      string
	_securityhubName                                      string
	_securityhubNextToken                                 string
	_securityhubNote                                      string
	_securityhubOnlyAssociated                            string
	_securityhubOrganizationConfiguration                 string
	_securityhubParameters                                string
	_securityhubProductArn                                string
	_securityhubProductSubscriptionArn                    string
	_securityhubProvider                                  string
	_securityhubProviderName                              string
	_securityhubRecordState                               string
	_securityhubRegionLinkingMode                         string
	_securityhubRegions                                   []string
	_securityhubRelatedFindings                           string
	_securityhubResourceArn                               string
	_securityhubRuleName                                  string
	_securityhubRuleOrder                                 string
	_securityhubRuleStatus                                string
	_securityhubSecurityControlId                         string
	_securityhubSecurityControlIds                        []string
	_securityhubSeverity                                  string
	_securityhubSeverityId                                string
	_securityhubSortCriteria                              string
	_securityhubSortOrder                                 string
	_securityhubStandardsArn                              string
	_securityhubStandardsControlArn                       string
	_securityhubStandardsControlAssociationIds            string
	_securityhubStandardsControlAssociationUpdates        string
	_securityhubStandardsSubscriptionArn                  string
	_securityhubStandardsSubscriptionArns                 []string
	_securityhubStandardsSubscriptionRequests             string
	_securityhubStartTime                                 string
	_securityhubStatusId                                  string
	_securityhubTagKeys                                   []string
	_securityhubTags                                      string
	_securityhubTarget                                    string
	_securityhubTypes                                     []string
	_securityhubUpdateAutomationRulesRequestItems         string
	_securityhubUpdatedReason                             string
	_securityhubUserDefinedFields                         string
	_securityhubVerificationState                         string
	_securityhubWorkflow                                  string
)

// We recommend using Organizations instead of Security Hub CSPM invitations to
// manage your member accounts. For information, see [Managing Security Hub CSPM administrator and member accounts with Organizations]in the Security Hub CSPM User
// Guide.
//
// Accepts the invitation to be a member account and be monitored by the Security
// Hub CSPM administrator account that the invitation was sent from.
//
// This operation is only used by member accounts that are not added through
// Organizations.
//
// When the member account accepts the invitation, permission is granted to the
// administrator account to view findings generated in the member account.
//
// [Managing Security Hub CSPM administrator and member accounts with Organizations]: https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-accounts-orgs.html
func securityhub_AcceptAdministratorInvitation(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.AcceptAdministratorInvitationInput{
		// AdministratorId: *string, // Required
		// InvitationId: *string, // Required
	}

	if len(_securityhubAdministratorId) > 0 {
		input.AdministratorId = aws.String(_securityhubAdministratorId)
	}
	if len(_securityhubInvitationId) > 0 {
		input.InvitationId = aws.String(_securityhubInvitationId)
	}

	if resp, err := client.AcceptAdministratorInvitation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This method is deprecated. Instead, use AcceptAdministratorInvitation .
// The Security Hub CSPM console continues to use AcceptInvitation . It will
// eventually change to use AcceptAdministratorInvitation . Any IAM policies that
// specifically control access to this function must continue to use
// AcceptInvitation . You should also add AcceptAdministratorInvitation to your
// policies to ensure that the correct permissions are in place after the console
// begins to use AcceptAdministratorInvitation .
//
// Accepts the invitation to be a member account and be monitored by the Security
// Hub CSPM administrator account that the invitation was sent from.
//
// This operation is only used by member accounts that are not added through
// Organizations.
//
// When the member account accepts the invitation, permission is granted to the
// administrator account to view findings generated in the member account.
//
// Deprecated: This API has been deprecated, use AcceptAdministratorInvitation API
// instead.
func securityhub_AcceptInvitation(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.AcceptInvitationInput{
		// InvitationId: *string, // Required
		// MasterId: *string, // Required
	}

	if len(_securityhubInvitationId) > 0 {
		input.InvitationId = aws.String(_securityhubInvitationId)
	}
	if len(_securityhubMasterId) > 0 {
		input.MasterId = aws.String(_securityhubMasterId)
	}

	if resp, err := client.AcceptInvitation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes one or more automation rules.
func securityhub_BatchDeleteAutomationRules(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.BatchDeleteAutomationRulesInput{
		// AutomationRulesArns: []string, // Required
	}

	if len(_securityhubAutomationRulesArns) > 0 {
		input.AutomationRulesArns = append([]string(nil), _securityhubAutomationRulesArns...)
	}

	if resp, err := client.BatchDeleteAutomationRules(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disables the standards specified by the provided StandardsSubscriptionArns .
// For more information, see [Security Standards] section of the Security Hub CSPM User Guide.
//
// [Security Standards]: https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-standards.html
func securityhub_BatchDisableStandards(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.BatchDisableStandardsInput{
		// StandardsSubscriptionArns: []string, // Required
	}

	if len(_securityhubStandardsSubscriptionArns) > 0 {
		input.StandardsSubscriptionArns = append([]string(nil), _securityhubStandardsSubscriptionArns...)
	}

	if resp, err := client.BatchDisableStandards(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables the standards specified by the provided StandardsArn . To obtain the ARN
// for a standard, use the DescribeStandards operation.
//
// For more information, see the [Security Standards] section of the Security Hub CSPM User Guide.
//
// [Security Standards]: https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-standards.html
func securityhub_BatchEnableStandards(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.BatchEnableStandardsInput{
		// StandardsSubscriptionRequests: []types.StandardsSubscriptionRequest, // Required
	}

	if len(_securityhubStandardsSubscriptionRequests) > 0 {
		if err := assignInputField(input, "StandardsSubscriptionRequests", _securityhubStandardsSubscriptionRequests); err != nil {
			log.Errorf("invalid --standards-subscription-requests: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchEnableStandards(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list of details for automation rules based on rule Amazon Resource
// Names (ARNs).
func securityhub_BatchGetAutomationRules(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.BatchGetAutomationRulesInput{
		// AutomationRulesArns: []string, // Required
	}

	if len(_securityhubAutomationRulesArns) > 0 {
		input.AutomationRulesArns = append([]string(nil), _securityhubAutomationRulesArns...)
	}

	if resp, err := client.BatchGetAutomationRules(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns associations between an Security Hub CSPM configuration and a batch of
// target accounts, organizational units, or the root. Only the Security Hub CSPM
// delegated administrator can invoke this operation from the home Region. A
// configuration can refer to a configuration policy or to a self-managed
// configuration.
func securityhub_BatchGetConfigurationPolicyAssociations(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.BatchGetConfigurationPolicyAssociationsInput{
		// ConfigurationPolicyAssociationIdentifiers: []types.ConfigurationPolicyAssociation, // Required
	}

	if len(_securityhubConfigurationPolicyAssociationIdentifiers) > 0 {
		if err := assignInputField(input, "ConfigurationPolicyAssociationIdentifiers", _securityhubConfigurationPolicyAssociationIdentifiers); err != nil {
			log.Errorf("invalid --configuration-policy-association-identifiers: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchGetConfigurationPolicyAssociations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides details about a batch of security controls for the current Amazon Web
// Services account and Amazon Web Services Region.
func securityhub_BatchGetSecurityControls(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.BatchGetSecurityControlsInput{
		// SecurityControlIds: []string, // Required
	}

	if len(_securityhubSecurityControlIds) > 0 {
		input.SecurityControlIds = append([]string(nil), _securityhubSecurityControlIds...)
	}

	if resp, err := client.BatchGetSecurityControls(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// For a batch of security controls and standards, identifies whether each
// control is currently enabled or disabled in a standard.
//
// Calls to this operation return a RESOURCE_NOT_FOUND_EXCEPTION error when the
// standard subscription for the association has a NOT_READY_FOR_UPDATES value for
// StandardsControlsUpdatable .
func securityhub_BatchGetStandardsControlAssociations(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.BatchGetStandardsControlAssociationsInput{
		// StandardsControlAssociationIds: []types.StandardsControlAssociationId, // Required
	}

	if len(_securityhubStandardsControlAssociationIds) > 0 {
		if err := assignInputField(input, "StandardsControlAssociationIds", _securityhubStandardsControlAssociationIds); err != nil {
			log.Errorf("invalid --standards-control-association-ids: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchGetStandardsControlAssociations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Imports security findings generated by a finding provider into Security Hub
// CSPM. This action is requested by the finding provider to import its findings
// into Security Hub CSPM.
//
// BatchImportFindings must be called by one of the following:
//
// - The Amazon Web Services account that is associated with a finding if you
// are using the [default product ARN]or are a partner sending findings from within a customer's
// Amazon Web Services account. In these cases, the identifier of the account that
// you are calling BatchImportFindings from needs to be the same as the
// AwsAccountId attribute for the finding.
//
// - An Amazon Web Services account that Security Hub CSPM has allow-listed for
// an official partner integration. In this case, you can call
// BatchImportFindings from the allow-listed account and send findings from
// different customer accounts in the same batch.
//
// The maximum allowed size for a finding is 240 Kb. An error is returned for any
// finding larger than 240 Kb.
//
// After a finding is created, BatchImportFindings cannot be used to update the
// following finding fields and objects, which Security Hub CSPM customers use to
// manage their investigation workflow.
//
// - Note
//
// - UserDefinedFields
//
// - VerificationState
//
// - Workflow
//
// Finding providers also should not use BatchImportFindings to update the
// following attributes.
//
// - Confidence
//
// - Criticality
//
// - RelatedFindings
//
// - Severity
//
// - Types
//
// Instead, finding providers use FindingProviderFields to provide values for
// these attributes.
//
// [default product ARN]: https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-custom-providers.html#securityhub-custom-providers-bfi-reqs
func securityhub_BatchImportFindings(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.BatchImportFindingsInput{
		// Findings: []types.AwsSecurityFinding, // Required
	}

	if len(_securityhubFindings) > 0 {
		if err := assignInputField(input, "Findings", _securityhubFindings); err != nil {
			log.Errorf("invalid --findings: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchImportFindings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates one or more automation rules based on rule Amazon Resource Names
// (ARNs) and input parameters.
func securityhub_BatchUpdateAutomationRules(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.BatchUpdateAutomationRulesInput{
		// UpdateAutomationRulesRequestItems: []types.UpdateAutomationRulesRequestItem, // Required
	}

	if len(_securityhubUpdateAutomationRulesRequestItems) > 0 {
		if err := assignInputField(input, "UpdateAutomationRulesRequestItems", _securityhubUpdateAutomationRulesRequestItems); err != nil {
			log.Errorf("invalid --update-automation-rules-request-items: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchUpdateAutomationRules(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Used by Security Hub CSPM customers to update information about their
// investigation into one or more findings. Requested by administrator accounts or
// member accounts. Administrator accounts can update findings for their account
// and their member accounts. A member account can update findings only for their
// own account. Administrator and member accounts can use this operation to update
// the following fields and objects for one or more findings:
//
// - Confidence
//
// - Criticality
//
// - Note
//
// - RelatedFindings
//
// - Severity
//
// - Types
//
// - UserDefinedFields
//
// - VerificationState
//
// - Workflow
//
// If you use this operation to update a finding, your updates don’t affect the
// value for the UpdatedAt field of the finding. Also note that it can take
// several minutes for Security Hub CSPM to process your request and update each
// finding specified in the request.
//
// You can configure IAM policies to restrict access to fields and field values.
// For example, you might not want member accounts to be able to suppress findings
// or change the finding severity. For more information see [Configuring access to BatchUpdateFindings]in the Security Hub
// CSPM User Guide.
//
// [Configuring access to BatchUpdateFindings]: https://docs.aws.amazon.com/securityhub/latest/userguide/finding-update-batchupdatefindings.html#batchupdatefindings-configure-access
func securityhub_BatchUpdateFindings(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.BatchUpdateFindingsInput{
		// FindingIdentifiers: []types.AwsSecurityFindingIdentifier, // Required
	}

	if len(_securityhubFindingIdentifiers) > 0 {
		if err := assignInputField(input, "FindingIdentifiers", _securityhubFindingIdentifiers); err != nil {
			log.Errorf("invalid --finding-identifiers: %s", err.Error())
			return
		}
	}
	if len(_securityhubConfidence) > 0 {
		if err := assignInputField(input, "Confidence", _securityhubConfidence); err != nil {
			log.Errorf("invalid --confidence: %s", err.Error())
			return
		}
	}
	if len(_securityhubCriticality) > 0 {
		if err := assignInputField(input, "Criticality", _securityhubCriticality); err != nil {
			log.Errorf("invalid --criticality: %s", err.Error())
			return
		}
	}
	if len(_securityhubNote) > 0 {
		if err := assignInputField(input, "Note", _securityhubNote); err != nil {
			log.Errorf("invalid --note: %s", err.Error())
			return
		}
	}
	if len(_securityhubRelatedFindings) > 0 {
		if err := assignInputField(input, "RelatedFindings", _securityhubRelatedFindings); err != nil {
			log.Errorf("invalid --related-findings: %s", err.Error())
			return
		}
	}
	if len(_securityhubSeverity) > 0 {
		if err := assignInputField(input, "Severity", _securityhubSeverity); err != nil {
			log.Errorf("invalid --severity: %s", err.Error())
			return
		}
	}
	if len(_securityhubTypes) > 0 {
		input.Types = append([]string(nil), _securityhubTypes...)
	}
	if len(_securityhubUserDefinedFields) > 0 {
		if err := assignInputField(input, "UserDefinedFields", _securityhubUserDefinedFields); err != nil {
			log.Errorf("invalid --user-defined-fields: %s", err.Error())
			return
		}
	}
	if len(_securityhubVerificationState) > 0 {
		if err := assignInputField(input, "VerificationState", _securityhubVerificationState); err != nil {
			log.Errorf("invalid --verification-state: %s", err.Error())
			return
		}
	}
	if len(_securityhubWorkflow) > 0 {
		if err := assignInputField(input, "Workflow", _securityhubWorkflow); err != nil {
			log.Errorf("invalid --workflow: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchUpdateFindings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Used by customers to update information about their investigation into a
// finding. Requested by delegated administrator accounts or member accounts.
// Delegated administrator accounts can update findings for their account and their
// member accounts. Member accounts can update findings for their account.
// BatchUpdateFindings and BatchUpdateFindingV2 both use
// securityhub:BatchUpdateFindings in the Action element of an IAM policy
// statement. You must have permission to perform the
// securityhub:BatchUpdateFindings action. Updates from BatchUpdateFindingsV2
// don't affect the value of f inding_info.modified_time ,
// finding_info.modified_time_dt , time , time_dt for a finding .
func securityhub_BatchUpdateFindingsV2(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.BatchUpdateFindingsV2Input{}

	if len(_securityhubComment) > 0 {
		input.Comment = aws.String(_securityhubComment)
	}
	if len(_securityhubFindingIdentifiers) > 0 {
		if err := assignInputField(input, "FindingIdentifiers", _securityhubFindingIdentifiers); err != nil {
			log.Errorf("invalid --finding-identifiers: %s", err.Error())
			return
		}
	}
	if len(_securityhubMetadataUids) > 0 {
		input.MetadataUids = append([]string(nil), _securityhubMetadataUids...)
	}
	if len(_securityhubSeverityId) > 0 {
		if err := assignInputField(input, "SeverityId", _securityhubSeverityId); err != nil {
			log.Errorf("invalid --severity-id: %s", err.Error())
			return
		}
	}
	if len(_securityhubStatusId) > 0 {
		if err := assignInputField(input, "StatusId", _securityhubStatusId); err != nil {
			log.Errorf("invalid --status-id: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchUpdateFindingsV2(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// For a batch of security controls and standards, this operation updates the
// enablement status of a control in a standard.
func securityhub_BatchUpdateStandardsControlAssociations(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.BatchUpdateStandardsControlAssociationsInput{
		// StandardsControlAssociationUpdates: []types.StandardsControlAssociationUpdate, // Required
	}

	if len(_securityhubStandardsControlAssociationUpdates) > 0 {
		if err := assignInputField(input, "StandardsControlAssociationUpdates", _securityhubStandardsControlAssociationUpdates); err != nil {
			log.Errorf("invalid --standards-control-association-updates: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchUpdateStandardsControlAssociations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a custom action target in Security Hub CSPM.
// You can use custom actions on findings and insights in Security Hub CSPM to
// trigger target actions in Amazon CloudWatch Events.
func securityhub_CreateActionTarget(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.CreateActionTargetInput{
		// Description: *string, // Required
		// Id: *string, // Required
		// Name: *string, // Required
	}

	if len(_securityhubDescription) > 0 {
		input.Description = aws.String(_securityhubDescription)
	}
	if len(_securityhubId) > 0 {
		input.Id = aws.String(_securityhubId)
	}
	if len(_securityhubName) > 0 {
		input.Name = aws.String(_securityhubName)
	}

	if resp, err := client.CreateActionTarget(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables aggregation across Amazon Web Services Regions.
func securityhub_CreateAggregatorV2(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.CreateAggregatorV2Input{
		// RegionLinkingMode: *string, // Required
	}

	if len(_securityhubRegionLinkingMode) > 0 {
		input.RegionLinkingMode = aws.String(_securityhubRegionLinkingMode)
	}
	if len(_securityhubClientToken) > 0 {
		input.ClientToken = aws.String(_securityhubClientToken)
	}
	if len(_securityhubLinkedRegions) > 0 {
		input.LinkedRegions = append([]string(nil), _securityhubLinkedRegions...)
	}
	if len(_securityhubTags) > 0 {
		if err := assignInputField(input, "Tags", _securityhubTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAggregatorV2(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an automation rule based on input parameters.
func securityhub_CreateAutomationRule(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.CreateAutomationRuleInput{
		// Actions: []types.AutomationRulesAction, // Required
		// Criteria: *types.AutomationRulesFindingFilters, // Required
		// Description: *string, // Required
		// RuleName: *string, // Required
		// RuleOrder: *int32, // Required
	}

	if len(_securityhubActions) > 0 {
		if err := assignInputField(input, "Actions", _securityhubActions); err != nil {
			log.Errorf("invalid --actions: %s", err.Error())
			return
		}
	}
	if len(_securityhubCriteria) > 0 {
		if err := assignInputField(input, "Criteria", _securityhubCriteria); err != nil {
			log.Errorf("invalid --criteria: %s", err.Error())
			return
		}
	}
	if len(_securityhubDescription) > 0 {
		input.Description = aws.String(_securityhubDescription)
	}
	if len(_securityhubRuleName) > 0 {
		input.RuleName = aws.String(_securityhubRuleName)
	}
	if len(_securityhubRuleOrder) > 0 {
		if err := assignInputField(input, "RuleOrder", _securityhubRuleOrder); err != nil {
			log.Errorf("invalid --rule-order: %s", err.Error())
			return
		}
	}
	if len(_securityhubIsTerminal) > 0 {
		if err := assignInputField(input, "IsTerminal", _securityhubIsTerminal); err != nil {
			log.Errorf("invalid --is-terminal: %s", err.Error())
			return
		}
	}
	if len(_securityhubRuleStatus) > 0 {
		if err := assignInputField(input, "RuleStatus", _securityhubRuleStatus); err != nil {
			log.Errorf("invalid --rule-status: %s", err.Error())
			return
		}
	}
	if len(_securityhubTags) > 0 {
		if err := assignInputField(input, "Tags", _securityhubTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAutomationRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a V2 automation rule.
func securityhub_CreateAutomationRuleV2(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.CreateAutomationRuleV2Input{
		// Actions: []types.AutomationRulesActionV2, // Required
		// Criteria: types.Criteria, // Required
		// Description: *string, // Required
		// RuleName: *string, // Required
		// RuleOrder: *float32, // Required
	}

	if len(_securityhubActions) > 0 {
		if err := assignInputField(input, "Actions", _securityhubActions); err != nil {
			log.Errorf("invalid --actions: %s", err.Error())
			return
		}
	}
	if len(_securityhubCriteria) > 0 {
		if err := assignInputField(input, "Criteria", _securityhubCriteria); err != nil {
			log.Errorf("invalid --criteria: %s", err.Error())
			return
		}
	}
	if len(_securityhubDescription) > 0 {
		input.Description = aws.String(_securityhubDescription)
	}
	if len(_securityhubRuleName) > 0 {
		input.RuleName = aws.String(_securityhubRuleName)
	}
	if len(_securityhubRuleOrder) > 0 {
		if err := assignInputField(input, "RuleOrder", _securityhubRuleOrder); err != nil {
			log.Errorf("invalid --rule-order: %s", err.Error())
			return
		}
	}
	if len(_securityhubClientToken) > 0 {
		input.ClientToken = aws.String(_securityhubClientToken)
	}
	if len(_securityhubRuleStatus) > 0 {
		if err := assignInputField(input, "RuleStatus", _securityhubRuleStatus); err != nil {
			log.Errorf("invalid --rule-status: %s", err.Error())
			return
		}
	}
	if len(_securityhubTags) > 0 {
		if err := assignInputField(input, "Tags", _securityhubTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAutomationRuleV2(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a configuration policy with the defined configuration. Only the
// Security Hub CSPM delegated administrator can invoke this operation from the
// home Region.
func securityhub_CreateConfigurationPolicy(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.CreateConfigurationPolicyInput{
		// ConfigurationPolicy: types.Policy, // Required
		// Name: *string, // Required
	}

	if len(_securityhubConfigurationPolicy) > 0 {
		if err := assignInputField(input, "ConfigurationPolicy", _securityhubConfigurationPolicy); err != nil {
			log.Errorf("invalid --configuration-policy: %s", err.Error())
			return
		}
	}
	if len(_securityhubName) > 0 {
		input.Name = aws.String(_securityhubName)
	}
	if len(_securityhubDescription) > 0 {
		input.Description = aws.String(_securityhubDescription)
	}
	if len(_securityhubTags) > 0 {
		if err := assignInputField(input, "Tags", _securityhubTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateConfigurationPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Grants permission to create a connectorV2 based on input parameters.
func securityhub_CreateConnectorV2(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.CreateConnectorV2Input{
		// Name: *string, // Required
		// Provider: types.ProviderConfiguration, // Required
	}

	if len(_securityhubName) > 0 {
		input.Name = aws.String(_securityhubName)
	}
	if len(_securityhubProvider) > 0 {
		if err := assignInputField(input, "Provider", _securityhubProvider); err != nil {
			log.Errorf("invalid --provider: %s", err.Error())
			return
		}
	}
	if len(_securityhubClientToken) > 0 {
		input.ClientToken = aws.String(_securityhubClientToken)
	}
	if len(_securityhubDescription) > 0 {
		input.Description = aws.String(_securityhubDescription)
	}
	if len(_securityhubKmsKeyArn) > 0 {
		input.KmsKeyArn = aws.String(_securityhubKmsKeyArn)
	}
	if len(_securityhubTags) > 0 {
		if err := assignInputField(input, "Tags", _securityhubTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateConnectorV2(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The aggregation Region is now called the home Region.
// Used to enable cross-Region aggregation. This operation can be invoked from the
// home Region only.
//
// For information about how cross-Region aggregation works, see [Understanding cross-Region aggregation in Security Hub CSPM] in the Security
// Hub CSPM User Guide.
//
// [Understanding cross-Region aggregation in Security Hub CSPM]: https://docs.aws.amazon.com/securityhub/latest/userguide/finding-aggregation.html
func securityhub_CreateFindingAggregator(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.CreateFindingAggregatorInput{
		// RegionLinkingMode: *string, // Required
	}

	if len(_securityhubRegionLinkingMode) > 0 {
		input.RegionLinkingMode = aws.String(_securityhubRegionLinkingMode)
	}
	if len(_securityhubRegions) > 0 {
		input.Regions = append([]string(nil), _securityhubRegions...)
	}

	if resp, err := client.CreateFindingAggregator(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a custom insight in Security Hub CSPM. An insight is a consolidation of
// findings that relate to a security issue that requires attention or remediation.
//
// To group the related findings in the insight, use the GroupByAttribute .
func securityhub_CreateInsight(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.CreateInsightInput{
		// Filters: *types.AwsSecurityFindingFilters, // Required
		// GroupByAttribute: *string, // Required
		// Name: *string, // Required
	}

	if len(_securityhubFilters) > 0 {
		if err := assignInputField(input, "Filters", _securityhubFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_securityhubGroupByAttribute) > 0 {
		input.GroupByAttribute = aws.String(_securityhubGroupByAttribute)
	}
	if len(_securityhubName) > 0 {
		input.Name = aws.String(_securityhubName)
	}

	if resp, err := client.CreateInsight(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a member association in Security Hub CSPM between the specified
// accounts and the account used to make the request, which is the administrator
// account. If you are integrated with Organizations, then the administrator
// account is designated by the organization management account.
//
// CreateMembers is always used to add accounts that are not organization members.
//
// For accounts that are managed using Organizations, CreateMembers is only used
// in the following cases:
//
// - Security Hub CSPM is not configured to automatically add new organization
// accounts.
//
// - The account was disassociated or deleted in Security Hub CSPM.
//
// This action can only be used by an account that has Security Hub CSPM enabled.
// To enable Security Hub CSPM, you can use the EnableSecurityHub operation.
//
// For accounts that are not organization members, you create the account
// association and then send an invitation to the member account. To send the
// invitation, you use the InviteMembers operation. If the account owner accepts
// the invitation, the account becomes a member account in Security Hub CSPM.
//
// Accounts that are managed using Organizations don't receive an invitation. They
// automatically become a member account in Security Hub CSPM.
//
// - If the organization account does not have Security Hub CSPM enabled, then
// Security Hub CSPM and the default standards are automatically enabled. Note that
// Security Hub CSPM cannot be enabled automatically for the organization
// management account. The organization management account must enable Security Hub
// CSPM before the administrator account enables it as a member account.
//
// - For organization accounts that already have Security Hub CSPM enabled,
// Security Hub CSPM does not make any other changes to those accounts. It does not
// change their enabled standards or controls.
//
// A permissions policy is added that permits the administrator account to view
// the findings generated in the member account.
//
// To remove the association between the administrator and member accounts, use
// the DisassociateFromMasterAccount or DisassociateMembers operation.
func securityhub_CreateMembers(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.CreateMembersInput{
		// AccountDetails: []types.AccountDetails, // Required
	}

	if len(_securityhubAccountDetails) > 0 {
		if err := assignInputField(input, "AccountDetails", _securityhubAccountDetails); err != nil {
			log.Errorf("invalid --account-details: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateMembers(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Grants permission to create a ticket in the chosen ITSM based on finding
// information for the provided finding metadata UID.
func securityhub_CreateTicketV2(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.CreateTicketV2Input{
		// ConnectorId: *string, // Required
		// FindingMetadataUid: *string, // Required
	}

	if len(_securityhubConnectorId) > 0 {
		input.ConnectorId = aws.String(_securityhubConnectorId)
	}
	if len(_securityhubFindingMetadataUid) > 0 {
		input.FindingMetadataUid = aws.String(_securityhubFindingMetadataUid)
	}
	if len(_securityhubClientToken) > 0 {
		input.ClientToken = aws.String(_securityhubClientToken)
	}
	if len(_securityhubMode) > 0 {
		if err := assignInputField(input, "Mode", _securityhubMode); err != nil {
			log.Errorf("invalid --mode: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateTicketV2(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// We recommend using Organizations instead of Security Hub CSPM invitations to
// manage your member accounts. For information, see [Managing Security Hub CSPM administrator and member accounts with Organizations]in the Security Hub CSPM User
// Guide.
//
// Declines invitations to become a Security Hub CSPM member account.
//
// A prospective member account uses this operation to decline an invitation to
// become a member.
//
// Only member accounts that aren't part of an Amazon Web Services organization
// should use this operation. Organization accounts don't receive invitations.
//
// [Managing Security Hub CSPM administrator and member accounts with Organizations]: https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-accounts-orgs.html
func securityhub_DeclineInvitations(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.DeclineInvitationsInput{
		// AccountIds: []string, // Required
	}

	if len(_securityhubAccountIds) > 0 {
		input.AccountIds = append([]string(nil), _securityhubAccountIds...)
	}

	if resp, err := client.DeclineInvitations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a custom action target from Security Hub CSPM.
// Deleting a custom action target does not affect any findings or insights that
// were already sent to Amazon CloudWatch Events using the custom action.
func securityhub_DeleteActionTarget(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.DeleteActionTargetInput{
		// ActionTargetArn: *string, // Required
	}

	if len(_securityhubActionTargetArn) > 0 {
		input.ActionTargetArn = aws.String(_securityhubActionTargetArn)
	}

	if resp, err := client.DeleteActionTarget(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the Aggregator V2.
func securityhub_DeleteAggregatorV2(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.DeleteAggregatorV2Input{
		// AggregatorV2Arn: *string, // Required
	}

	if len(_securityhubAggregatorV2Arn) > 0 {
		input.AggregatorV2Arn = aws.String(_securityhubAggregatorV2Arn)
	}

	if resp, err := client.DeleteAggregatorV2(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a V2 automation rule.
func securityhub_DeleteAutomationRuleV2(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.DeleteAutomationRuleV2Input{
		// Identifier: *string, // Required
	}

	if len(_securityhubIdentifier) > 0 {
		input.Identifier = aws.String(_securityhubIdentifier)
	}

	if resp, err := client.DeleteAutomationRuleV2(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a configuration policy. Only the Security Hub CSPM delegated
// administrator can invoke this operation from the home Region. For the deletion
// to succeed, you must first disassociate a configuration policy from target
// accounts, organizational units, or the root by invoking the
// StartConfigurationPolicyDisassociation operation.
func securityhub_DeleteConfigurationPolicy(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.DeleteConfigurationPolicyInput{
		// Identifier: *string, // Required
	}

	if len(_securityhubIdentifier) > 0 {
		input.Identifier = aws.String(_securityhubIdentifier)
	}

	if resp, err := client.DeleteConfigurationPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Grants permission to delete a connectorV2.
func securityhub_DeleteConnectorV2(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.DeleteConnectorV2Input{
		// ConnectorId: *string, // Required
	}

	if len(_securityhubConnectorId) > 0 {
		input.ConnectorId = aws.String(_securityhubConnectorId)
	}

	if resp, err := client.DeleteConnectorV2(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The aggregation Region is now called the home Region.
// Deletes a finding aggregator. When you delete the finding aggregator, you stop
// cross-Region aggregation. Finding replication stops occurring from the linked
// Regions to the home Region.
//
// When you stop cross-Region aggregation, findings that were already replicated
// and sent to the home Region are still visible from the home Region. However, new
// findings and finding updates are no longer replicated and sent to the home
// Region.
func securityhub_DeleteFindingAggregator(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.DeleteFindingAggregatorInput{
		// FindingAggregatorArn: *string, // Required
	}

	if len(_securityhubFindingAggregatorArn) > 0 {
		input.FindingAggregatorArn = aws.String(_securityhubFindingAggregatorArn)
	}

	if resp, err := client.DeleteFindingAggregator(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the insight specified by the InsightArn .
func securityhub_DeleteInsight(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.DeleteInsightInput{
		// InsightArn: *string, // Required
	}

	if len(_securityhubInsightArn) > 0 {
		input.InsightArn = aws.String(_securityhubInsightArn)
	}

	if resp, err := client.DeleteInsight(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// We recommend using Organizations instead of Security Hub CSPM invitations to
// manage your member accounts. For information, see [Managing Security Hub CSPM administrator and member accounts with Organizations]in the Security Hub CSPM User
// Guide.
//
// Deletes invitations to become a Security Hub CSPM member account.
//
// A Security Hub CSPM administrator account can use this operation to delete
// invitations sent to one or more prospective member accounts.
//
// This operation is only used to delete invitations that are sent to prospective
// member accounts that aren't part of an Amazon Web Services organization.
// Organization accounts don't receive invitations.
//
// [Managing Security Hub CSPM administrator and member accounts with Organizations]: https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-accounts-orgs.html
func securityhub_DeleteInvitations(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.DeleteInvitationsInput{
		// AccountIds: []string, // Required
	}

	if len(_securityhubAccountIds) > 0 {
		input.AccountIds = append([]string(nil), _securityhubAccountIds...)
	}

	if resp, err := client.DeleteInvitations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified member accounts from Security Hub CSPM.
// You can invoke this API only to delete accounts that became members through
// invitation. You can't invoke this API to delete accounts that belong to an
// Organizations organization.
func securityhub_DeleteMembers(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.DeleteMembersInput{
		// AccountIds: []string, // Required
	}

	if len(_securityhubAccountIds) > 0 {
		input.AccountIds = append([]string(nil), _securityhubAccountIds...)
	}

	if resp, err := client.DeleteMembers(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of the custom action targets in Security Hub CSPM in your
// account.
func securityhub_DescribeActionTargets(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.DescribeActionTargetsInput{}

	if len(_securityhubActionTargetArns) > 0 {
		input.ActionTargetArns = append([]string(nil), _securityhubActionTargetArns...)
	}
	if len(_securityhubMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _securityhubMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_securityhubNextToken) > 0 {
		input.NextToken = aws.String(_securityhubNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeActionTargets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*securityhub.DescribeActionTargetsOutput
	p := securityhub.NewDescribeActionTargetsPaginator(client, input)
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

// Returns details about the Hub resource in your account, including the HubArn
// and the time when you enabled Security Hub CSPM.
func securityhub_DescribeHub(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.DescribeHubInput{}

	if len(_securityhubHubArn) > 0 {
		input.HubArn = aws.String(_securityhubHubArn)
	}

	if resp, err := client.DescribeHub(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the way your organization is configured in Security
// Hub CSPM. Only the Security Hub CSPM administrator account can invoke this
// operation.
func securityhub_DescribeOrganizationConfiguration(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.DescribeOrganizationConfigurationInput{}

	if resp, err := client.DescribeOrganizationConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about product integrations in Security Hub CSPM.
// You can optionally provide an integration ARN. If you provide an integration
// ARN, then the results only include that integration.
//
// If you don't provide an integration ARN, then the results include all of the
// available product integrations.
func securityhub_DescribeProducts(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.DescribeProductsInput{}

	if len(_securityhubMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _securityhubMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_securityhubNextToken) > 0 {
		input.NextToken = aws.String(_securityhubNextToken)
	}
	if len(_securityhubProductArn) > 0 {
		input.ProductArn = aws.String(_securityhubProductArn)
	}

	if disablePaginator() {
		if resp, err := client.DescribeProducts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*securityhub.DescribeProductsOutput
	p := securityhub.NewDescribeProductsPaginator(client, input)
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

// Gets information about the product integration.
func securityhub_DescribeProductsV2(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.DescribeProductsV2Input{}

	if len(_securityhubMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _securityhubMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_securityhubNextToken) > 0 {
		input.NextToken = aws.String(_securityhubNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeProductsV2(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*securityhub.DescribeProductsV2Output
	p := securityhub.NewDescribeProductsV2Paginator(client, input)
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

// Returns details about the service resource in your account.
func securityhub_DescribeSecurityHubV2(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.DescribeSecurityHubV2Input{}

	if resp, err := client.DescribeSecurityHubV2(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of the available standards in Security Hub CSPM.
// For each standard, the results include the standard ARN, the name, and a
// description.
func securityhub_DescribeStandards(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.DescribeStandardsInput{}

	if len(_securityhubMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _securityhubMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_securityhubNextToken) > 0 {
		input.NextToken = aws.String(_securityhubNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeStandards(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*securityhub.DescribeStandardsOutput
	p := securityhub.NewDescribeStandardsPaginator(client, input)
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

// Returns a list of security standards controls.
// For each control, the results include information about whether it is currently
// enabled, the severity, and a link to remediation information.
//
// This operation returns an empty list for standard subscriptions where
// StandardsControlsUpdatable has value NOT_READY_FOR_UPDATES .
func securityhub_DescribeStandardsControls(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.DescribeStandardsControlsInput{
		// StandardsSubscriptionArn: *string, // Required
	}

	if len(_securityhubStandardsSubscriptionArn) > 0 {
		input.StandardsSubscriptionArn = aws.String(_securityhubStandardsSubscriptionArn)
	}
	if len(_securityhubMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _securityhubMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_securityhubNextToken) > 0 {
		input.NextToken = aws.String(_securityhubNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeStandardsControls(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*securityhub.DescribeStandardsControlsOutput
	p := securityhub.NewDescribeStandardsControlsPaginator(client, input)
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

// Disables the integration of the specified product with Security Hub CSPM. After
// the integration is disabled, findings from that product are no longer sent to
// Security Hub CSPM.
func securityhub_DisableImportFindingsForProduct(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.DisableImportFindingsForProductInput{
		// ProductSubscriptionArn: *string, // Required
	}

	if len(_securityhubProductSubscriptionArn) > 0 {
		input.ProductSubscriptionArn = aws.String(_securityhubProductSubscriptionArn)
	}

	if resp, err := client.DisableImportFindingsForProduct(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disables a Security Hub CSPM administrator account. Can only be called by the
// organization management account.
func securityhub_DisableOrganizationAdminAccount(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.DisableOrganizationAdminAccountInput{
		// AdminAccountId: *string, // Required
	}

	if len(_securityhubAdminAccountId) > 0 {
		input.AdminAccountId = aws.String(_securityhubAdminAccountId)
	}
	if len(_securityhubFeature) > 0 {
		if err := assignInputField(input, "Feature", _securityhubFeature); err != nil {
			log.Errorf("invalid --feature: %s", err.Error())
			return
		}
	}

	if resp, err := client.DisableOrganizationAdminAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disables Security Hub CSPM in your account only in the current Amazon Web
// Services Region. To disable Security Hub CSPM in all Regions, you must submit
// one request per Region where you have enabled Security Hub CSPM.
//
// You can't disable Security Hub CSPM in an account that is currently the
// Security Hub CSPM administrator.
//
// When you disable Security Hub CSPM, your existing findings and insights and any
// Security Hub CSPM configuration settings are deleted after 90 days and cannot be
// recovered. Any standards that were enabled are disabled, and your administrator
// and member account associations are removed.
//
// If you want to save your existing findings, you must export them before you
// disable Security Hub CSPM.
func securityhub_DisableSecurityHub(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.DisableSecurityHubInput{}

	if resp, err := client.DisableSecurityHub(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disable the service for the current Amazon Web Services Region or specified
// Amazon Web Services Region.
func securityhub_DisableSecurityHubV2(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.DisableSecurityHubV2Input{}

	if resp, err := client.DisableSecurityHubV2(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates the current Security Hub CSPM member account from the associated
// administrator account.
//
// This operation is only used by accounts that are not part of an organization.
// For organization accounts, only the administrator account can disassociate a
// member account.
func securityhub_DisassociateFromAdministratorAccount(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.DisassociateFromAdministratorAccountInput{}

	if resp, err := client.DisassociateFromAdministratorAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This method is deprecated. Instead, use DisassociateFromAdministratorAccount .
// The Security Hub CSPM console continues to use DisassociateFromMasterAccount .
// It will eventually change to use DisassociateFromAdministratorAccount . Any IAM
// policies that specifically control access to this function must continue to use
// DisassociateFromMasterAccount . You should also add
// DisassociateFromAdministratorAccount to your policies to ensure that the correct
// permissions are in place after the console begins to use
// DisassociateFromAdministratorAccount .
//
// Disassociates the current Security Hub CSPM member account from the associated
// administrator account.
//
// This operation is only used by accounts that are not part of an organization.
// For organization accounts, only the administrator account can disassociate a
// member account.
//
// Deprecated: This API has been deprecated, use
// DisassociateFromAdministratorAccount API instead.
func securityhub_DisassociateFromMasterAccount(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.DisassociateFromMasterAccountInput{}

	if resp, err := client.DisassociateFromMasterAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates the specified member accounts from the associated administrator
// account.
//
// Can be used to disassociate both accounts that are managed using Organizations
// and accounts that were invited manually.
func securityhub_DisassociateMembers(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.DisassociateMembersInput{
		// AccountIds: []string, // Required
	}

	if len(_securityhubAccountIds) > 0 {
		input.AccountIds = append([]string(nil), _securityhubAccountIds...)
	}

	if resp, err := client.DisassociateMembers(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables the integration of a partner product with Security Hub CSPM. Integrated
// products send findings to Security Hub CSPM.
//
// When you enable a product integration, a permissions policy that grants
// permission for the product to send findings to Security Hub CSPM is applied.
func securityhub_EnableImportFindingsForProduct(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.EnableImportFindingsForProductInput{
		// ProductArn: *string, // Required
	}

	if len(_securityhubProductArn) > 0 {
		input.ProductArn = aws.String(_securityhubProductArn)
	}

	if resp, err := client.EnableImportFindingsForProduct(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Designates the Security Hub CSPM administrator account for an organization. Can
// only be called by the organization management account.
func securityhub_EnableOrganizationAdminAccount(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.EnableOrganizationAdminAccountInput{
		// AdminAccountId: *string, // Required
	}

	if len(_securityhubAdminAccountId) > 0 {
		input.AdminAccountId = aws.String(_securityhubAdminAccountId)
	}
	if len(_securityhubFeature) > 0 {
		if err := assignInputField(input, "Feature", _securityhubFeature); err != nil {
			log.Errorf("invalid --feature: %s", err.Error())
			return
		}
	}

	if resp, err := client.EnableOrganizationAdminAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables Security Hub CSPM for your account in the current Region or the Region
// you specify in the request.
//
// When you enable Security Hub CSPM, you grant to Security Hub CSPM the
// permissions necessary to gather findings from other services that are integrated
// with Security Hub CSPM.
//
// When you use the EnableSecurityHub operation to enable Security Hub CSPM, you
// also automatically enable the following standards:
//
// - Center for Internet Security (CIS) Amazon Web Services Foundations
// Benchmark v1.2.0
//
// - Amazon Web Services Foundational Security Best Practices
//
// Other standards are not automatically enabled.
//
// To opt out of automatically enabled standards, set EnableDefaultStandards to
// false .
//
// After you enable Security Hub CSPM, to enable a standard, use the
// BatchEnableStandards operation. To disable a standard, use the
// BatchDisableStandards operation.
//
// To learn more, see the [setup information] in the Security Hub CSPM User Guide.
//
// [setup information]: https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-settingup.html
func securityhub_EnableSecurityHub(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.EnableSecurityHubInput{}

	if len(_securityhubControlFindingGenerator) > 0 {
		if err := assignInputField(input, "ControlFindingGenerator", _securityhubControlFindingGenerator); err != nil {
			log.Errorf("invalid --control-finding-generator: %s", err.Error())
			return
		}
	}
	if len(_securityhubEnableDefaultStandards) > 0 {
		if err := assignInputField(input, "EnableDefaultStandards", _securityhubEnableDefaultStandards); err != nil {
			log.Errorf("invalid --enable-default-standards: %s", err.Error())
			return
		}
	}
	if len(_securityhubTags) > 0 {
		if err := assignInputField(input, "Tags", _securityhubTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.EnableSecurityHub(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables the service in account for the current Amazon Web Services Region or
// specified Amazon Web Services Region.
func securityhub_EnableSecurityHubV2(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.EnableSecurityHubV2Input{}

	if len(_securityhubTags) > 0 {
		if err := assignInputField(input, "Tags", _securityhubTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.EnableSecurityHubV2(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides the details for the Security Hub CSPM administrator account for the
// current member account.
//
// Can be used by both member accounts that are managed using Organizations and
// accounts that were invited manually.
func securityhub_GetAdministratorAccount(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.GetAdministratorAccountInput{}

	if resp, err := client.GetAdministratorAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the configuration of the specified Aggregator V2.
func securityhub_GetAggregatorV2(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.GetAggregatorV2Input{
		// AggregatorV2Arn: *string, // Required
	}

	if len(_securityhubAggregatorV2Arn) > 0 {
		input.AggregatorV2Arn = aws.String(_securityhubAggregatorV2Arn)
	}

	if resp, err := client.GetAggregatorV2(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns an automation rule for the V2 service.
func securityhub_GetAutomationRuleV2(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.GetAutomationRuleV2Input{
		// Identifier: *string, // Required
	}

	if len(_securityhubIdentifier) > 0 {
		input.Identifier = aws.String(_securityhubIdentifier)
	}

	if resp, err := client.GetAutomationRuleV2(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides information about a configuration policy. Only the Security Hub CSPM
// delegated administrator can invoke this operation from the home Region.
func securityhub_GetConfigurationPolicy(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.GetConfigurationPolicyInput{
		// Identifier: *string, // Required
	}

	if len(_securityhubIdentifier) > 0 {
		input.Identifier = aws.String(_securityhubIdentifier)
	}

	if resp, err := client.GetConfigurationPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the association between a configuration and a target account,
// organizational unit, or the root. The configuration can be a configuration
// policy or self-managed behavior. Only the Security Hub CSPM delegated
// administrator can invoke this operation from the home Region.
func securityhub_GetConfigurationPolicyAssociation(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.GetConfigurationPolicyAssociationInput{
		// Target: types.Target, // Required
	}

	if len(_securityhubTarget) > 0 {
		if err := assignInputField(input, "Target", _securityhubTarget); err != nil {
			log.Errorf("invalid --target: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetConfigurationPolicyAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Grants permission to retrieve details for a connectorV2 based on connector id.
func securityhub_GetConnectorV2(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.GetConnectorV2Input{
		// ConnectorId: *string, // Required
	}

	if len(_securityhubConnectorId) > 0 {
		input.ConnectorId = aws.String(_securityhubConnectorId)
	}

	if resp, err := client.GetConnectorV2(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of the standards that are currently enabled.
func securityhub_GetEnabledStandards(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.GetEnabledStandardsInput{}

	if len(_securityhubMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _securityhubMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_securityhubNextToken) > 0 {
		input.NextToken = aws.String(_securityhubNextToken)
	}
	if len(_securityhubStandardsSubscriptionArns) > 0 {
		input.StandardsSubscriptionArns = append([]string(nil), _securityhubStandardsSubscriptionArns...)
	}

	if disablePaginator() {
		if resp, err := client.GetEnabledStandards(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*securityhub.GetEnabledStandardsOutput
	p := securityhub.NewGetEnabledStandardsPaginator(client, input)
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

// The aggregation Region is now called the home Region.
// Returns the current configuration in the calling account for cross-Region
// aggregation. A finding aggregator is a resource that establishes the home Region
// and any linked Regions.
func securityhub_GetFindingAggregator(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.GetFindingAggregatorInput{
		// FindingAggregatorArn: *string, // Required
	}

	if len(_securityhubFindingAggregatorArn) > 0 {
		input.FindingAggregatorArn = aws.String(_securityhubFindingAggregatorArn)
	}

	if resp, err := client.GetFindingAggregator(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the history of a Security Hub CSPM finding. The history includes
// changes made to any fields in the Amazon Web Services Security Finding Format
// (ASFF) except top-level timestamp fields, such as the CreatedAt and UpdatedAt
// fields.
//
// This operation might return fewer results than the maximum number of results (
// MaxResults ) specified in a request, even when more results are available. If
// this occurs, the response includes a NextToken value, which you should use to
// retrieve the next set of results in the response. The presence of a NextToken
// value in a response doesn't necessarily indicate that the results are
// incomplete. However, you should continue to specify a NextToken value until you
// receive a response that doesn't include this value.
func securityhub_GetFindingHistory(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.GetFindingHistoryInput{
		// FindingIdentifier: *types.AwsSecurityFindingIdentifier, // Required
	}

	if len(_securityhubFindingIdentifier) > 0 {
		if err := assignInputField(input, "FindingIdentifier", _securityhubFindingIdentifier); err != nil {
			log.Errorf("invalid --finding-identifier: %s", err.Error())
			return
		}
	}
	if len(_securityhubEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _securityhubEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_securityhubMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _securityhubMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_securityhubNextToken) > 0 {
		input.NextToken = aws.String(_securityhubNextToken)
	}
	if len(_securityhubStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _securityhubStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.GetFindingHistory(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*securityhub.GetFindingHistoryOutput
	p := securityhub.NewGetFindingHistoryPaginator(client, input)
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

// Returns aggregated statistical data about findings. GetFindingStatisticsV2 use
// securityhub:GetAdhocInsightResults in the Action element of an IAM policy
// statement. You must have permission to perform the s action.
func securityhub_GetFindingStatisticsV2(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.GetFindingStatisticsV2Input{
		// GroupByRules: []types.GroupByRule, // Required
	}

	if len(_securityhubGroupByRules) > 0 {
		if err := assignInputField(input, "GroupByRules", _securityhubGroupByRules); err != nil {
			log.Errorf("invalid --group-by-rules: %s", err.Error())
			return
		}
	}
	if len(_securityhubMaxStatisticResults) > 0 {
		if err := assignInputField(input, "MaxStatisticResults", _securityhubMaxStatisticResults); err != nil {
			log.Errorf("invalid --max-statistic-results: %s", err.Error())
			return
		}
	}
	if len(_securityhubSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _securityhubSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetFindingStatisticsV2(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of findings that match the specified criteria.
// If cross-Region aggregation is enabled, then when you call GetFindings from the
// home Region, the results include all of the matching findings from both the home
// Region and linked Regions.
func securityhub_GetFindings(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.GetFindingsInput{}

	if len(_securityhubFilters) > 0 {
		if err := assignInputField(input, "Filters", _securityhubFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_securityhubMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _securityhubMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_securityhubNextToken) > 0 {
		input.NextToken = aws.String(_securityhubNextToken)
	}
	if len(_securityhubSortCriteria) > 0 {
		if err := assignInputField(input, "SortCriteria", _securityhubSortCriteria); err != nil {
			log.Errorf("invalid --sort-criteria: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.GetFindings(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*securityhub.GetFindingsOutput
	p := securityhub.NewGetFindingsPaginator(client, input)
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

// Returns findings trend data based on the specified criteria. This operation
// helps you analyze patterns and changes in findings over time.
func securityhub_GetFindingsTrendsV2(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.GetFindingsTrendsV2Input{
		// EndTime: *time.Time, // Required
		// StartTime: *time.Time, // Required
	}

	if len(_securityhubEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _securityhubEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_securityhubStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _securityhubStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_securityhubFilters) > 0 {
		if err := assignInputField(input, "Filters", _securityhubFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_securityhubMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _securityhubMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_securityhubNextToken) > 0 {
		input.NextToken = aws.String(_securityhubNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetFindingsTrendsV2(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*securityhub.GetFindingsTrendsV2Output
	p := securityhub.NewGetFindingsTrendsV2Paginator(client, input)
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

// Return a list of findings that match the specified criteria. GetFindings and
// GetFindingsV2 both use securityhub:GetFindings in the Action element of an IAM
// policy statement. You must have permission to perform the
// securityhub:GetFindings action.
func securityhub_GetFindingsV2(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.GetFindingsV2Input{}

	if len(_securityhubFilters) > 0 {
		if err := assignInputField(input, "Filters", _securityhubFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_securityhubMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _securityhubMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_securityhubNextToken) > 0 {
		input.NextToken = aws.String(_securityhubNextToken)
	}
	if len(_securityhubSortCriteria) > 0 {
		if err := assignInputField(input, "SortCriteria", _securityhubSortCriteria); err != nil {
			log.Errorf("invalid --sort-criteria: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.GetFindingsV2(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*securityhub.GetFindingsV2Output
	p := securityhub.NewGetFindingsV2Paginator(client, input)
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

// Lists the results of the Security Hub CSPM insight specified by the insight ARN.
func securityhub_GetInsightResults(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.GetInsightResultsInput{
		// InsightArn: *string, // Required
	}

	if len(_securityhubInsightArn) > 0 {
		input.InsightArn = aws.String(_securityhubInsightArn)
	}

	if resp, err := client.GetInsightResults(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists and describes insights for the specified insight ARNs.
func securityhub_GetInsights(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.GetInsightsInput{}

	if len(_securityhubInsightArns) > 0 {
		input.InsightArns = append([]string(nil), _securityhubInsightArns...)
	}
	if len(_securityhubMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _securityhubMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_securityhubNextToken) > 0 {
		input.NextToken = aws.String(_securityhubNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetInsights(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*securityhub.GetInsightsOutput
	p := securityhub.NewGetInsightsPaginator(client, input)
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

// We recommend using Organizations instead of Security Hub CSPM invitations to
// manage your member accounts. For information, see [Managing Security Hub CSPM administrator and member accounts with Organizations]in the Security Hub CSPM User
// Guide.
//
// Returns the count of all Security Hub CSPM membership invitations that were
// sent to the calling member account, not including the currently accepted
// invitation.
//
// [Managing Security Hub CSPM administrator and member accounts with Organizations]: https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-accounts-orgs.html
func securityhub_GetInvitationsCount(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.GetInvitationsCountInput{}

	if resp, err := client.GetInvitationsCount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This method is deprecated. Instead, use GetAdministratorAccount .
// The Security Hub CSPM console continues to use GetMasterAccount . It will
// eventually change to use GetAdministratorAccount . Any IAM policies that
// specifically control access to this function must continue to use
// GetMasterAccount . You should also add GetAdministratorAccount to your policies
// to ensure that the correct permissions are in place after the console begins to
// use GetAdministratorAccount .
//
// Provides the details for the Security Hub CSPM administrator account for the
// current member account.
//
// Can be used by both member accounts that are managed using Organizations and
// accounts that were invited manually.
//
// Deprecated: This API has been deprecated, use GetAdministratorAccount API
// instead.
func securityhub_GetMasterAccount(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.GetMasterAccountInput{}

	if resp, err := client.GetMasterAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the details for the Security Hub CSPM member accounts for the specified
// account IDs.
//
// An administrator account can be either the delegated Security Hub CSPM
// administrator account for an organization or an administrator account that
// enabled Security Hub CSPM manually.
//
// The results include both member accounts that are managed using Organizations
// and accounts that were invited manually.
func securityhub_GetMembers(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.GetMembersInput{
		// AccountIds: []string, // Required
	}

	if len(_securityhubAccountIds) > 0 {
		input.AccountIds = append([]string(nil), _securityhubAccountIds...)
	}

	if resp, err := client.GetMembers(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves statistical information about Amazon Web Services resources and their
// associated security findings.
func securityhub_GetResourcesStatisticsV2(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.GetResourcesStatisticsV2Input{
		// GroupByRules: []types.ResourceGroupByRule, // Required
	}

	if len(_securityhubGroupByRules) > 0 {
		if err := assignInputField(input, "GroupByRules", _securityhubGroupByRules); err != nil {
			log.Errorf("invalid --group-by-rules: %s", err.Error())
			return
		}
	}
	if len(_securityhubMaxStatisticResults) > 0 {
		if err := assignInputField(input, "MaxStatisticResults", _securityhubMaxStatisticResults); err != nil {
			log.Errorf("invalid --max-statistic-results: %s", err.Error())
			return
		}
	}
	if len(_securityhubSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _securityhubSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetResourcesStatisticsV2(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns resource trend data based on the specified criteria. This operation
// helps you analyze patterns and changes in resource compliance over time.
func securityhub_GetResourcesTrendsV2(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.GetResourcesTrendsV2Input{
		// EndTime: *time.Time, // Required
		// StartTime: *time.Time, // Required
	}

	if len(_securityhubEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _securityhubEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_securityhubStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _securityhubStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_securityhubFilters) > 0 {
		if err := assignInputField(input, "Filters", _securityhubFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_securityhubMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _securityhubMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_securityhubNextToken) > 0 {
		input.NextToken = aws.String(_securityhubNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetResourcesTrendsV2(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*securityhub.GetResourcesTrendsV2Output
	p := securityhub.NewGetResourcesTrendsV2Paginator(client, input)
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

// Returns a list of resources.
func securityhub_GetResourcesV2(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.GetResourcesV2Input{}

	if len(_securityhubFilters) > 0 {
		if err := assignInputField(input, "Filters", _securityhubFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_securityhubMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _securityhubMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_securityhubNextToken) > 0 {
		input.NextToken = aws.String(_securityhubNextToken)
	}
	if len(_securityhubSortCriteria) > 0 {
		if err := assignInputField(input, "SortCriteria", _securityhubSortCriteria); err != nil {
			log.Errorf("invalid --sort-criteria: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.GetResourcesV2(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*securityhub.GetResourcesV2Output
	p := securityhub.NewGetResourcesV2Paginator(client, input)
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

// Retrieves the definition of a security control. The definition includes the
// control title, description, Region availability, parameter definitions, and
// other details.
func securityhub_GetSecurityControlDefinition(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.GetSecurityControlDefinitionInput{
		// SecurityControlId: *string, // Required
	}

	if len(_securityhubSecurityControlId) > 0 {
		input.SecurityControlId = aws.String(_securityhubSecurityControlId)
	}

	if resp, err := client.GetSecurityControlDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// We recommend using Organizations instead of Security Hub CSPM invitations to
// manage your member accounts. For information, see [Managing Security Hub CSPM administrator and member accounts with Organizations]in the Security Hub CSPM User
// Guide.
//
// Invites other Amazon Web Services accounts to become member accounts for the
// Security Hub CSPM administrator account that the invitation is sent from.
//
// This operation is only used to invite accounts that don't belong to an Amazon
// Web Services organization. Organization accounts don't receive invitations.
//
// Before you can use this action to invite a member, you must first use the
// CreateMembers action to create the member account in Security Hub CSPM.
//
// When the account owner enables Security Hub CSPM and accepts the invitation to
// become a member account, the administrator account can view the findings
// generated in the member account.
//
// [Managing Security Hub CSPM administrator and member accounts with Organizations]: https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-accounts-orgs.html
func securityhub_InviteMembers(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.InviteMembersInput{
		// AccountIds: []string, // Required
	}

	if len(_securityhubAccountIds) > 0 {
		input.AccountIds = append([]string(nil), _securityhubAccountIds...)
	}

	if resp, err := client.InviteMembers(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list of V2 aggregators.
func securityhub_ListAggregatorsV2(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.ListAggregatorsV2Input{}

	if len(_securityhubMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _securityhubMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_securityhubNextToken) > 0 {
		input.NextToken = aws.String(_securityhubNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAggregatorsV2(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*securityhub.ListAggregatorsV2Output
	p := securityhub.NewListAggregatorsV2Paginator(client, input)
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

// A list of automation rules and their metadata for the calling account.
func securityhub_ListAutomationRules(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.ListAutomationRulesInput{}

	if len(_securityhubMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _securityhubMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_securityhubNextToken) > 0 {
		input.NextToken = aws.String(_securityhubNextToken)
	}

	if resp, err := client.ListAutomationRules(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of automation rules and metadata for the calling account.
func securityhub_ListAutomationRulesV2(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.ListAutomationRulesV2Input{}

	if len(_securityhubMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _securityhubMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_securityhubNextToken) > 0 {
		input.NextToken = aws.String(_securityhubNextToken)
	}

	if resp, err := client.ListAutomationRulesV2(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the configuration policies that the Security Hub CSPM delegated
// administrator has created for your organization. Only the delegated
// administrator can invoke this operation from the home Region.
func securityhub_ListConfigurationPolicies(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.ListConfigurationPoliciesInput{}

	if len(_securityhubMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _securityhubMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_securityhubNextToken) > 0 {
		input.NextToken = aws.String(_securityhubNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListConfigurationPolicies(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*securityhub.ListConfigurationPoliciesOutput
	p := securityhub.NewListConfigurationPoliciesPaginator(client, input)
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

// Provides information about the associations for your configuration policies
// and self-managed behavior. Only the Security Hub CSPM delegated administrator
// can invoke this operation from the home Region.
func securityhub_ListConfigurationPolicyAssociations(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.ListConfigurationPolicyAssociationsInput{}

	if len(_securityhubFilters) > 0 {
		if err := assignInputField(input, "Filters", _securityhubFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_securityhubMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _securityhubMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_securityhubNextToken) > 0 {
		input.NextToken = aws.String(_securityhubNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListConfigurationPolicyAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*securityhub.ListConfigurationPolicyAssociationsOutput
	p := securityhub.NewListConfigurationPolicyAssociationsPaginator(client, input)
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

// Grants permission to retrieve a list of connectorsV2 and their metadata for the
// calling account.
func securityhub_ListConnectorsV2(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.ListConnectorsV2Input{}

	if len(_securityhubConnectorStatus) > 0 {
		if err := assignInputField(input, "ConnectorStatus", _securityhubConnectorStatus); err != nil {
			log.Errorf("invalid --connector-status: %s", err.Error())
			return
		}
	}
	if len(_securityhubMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _securityhubMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_securityhubNextToken) > 0 {
		input.NextToken = aws.String(_securityhubNextToken)
	}
	if len(_securityhubProviderName) > 0 {
		if err := assignInputField(input, "ProviderName", _securityhubProviderName); err != nil {
			log.Errorf("invalid --provider-name: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListConnectorsV2(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all findings-generating solutions (products) that you are subscribed to
// receive findings from in Security Hub CSPM.
func securityhub_ListEnabledProductsForImport(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.ListEnabledProductsForImportInput{}

	if len(_securityhubMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _securityhubMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_securityhubNextToken) > 0 {
		input.NextToken = aws.String(_securityhubNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListEnabledProductsForImport(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*securityhub.ListEnabledProductsForImportOutput
	p := securityhub.NewListEnabledProductsForImportPaginator(client, input)
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

// If cross-Region aggregation is enabled, then ListFindingAggregators returns the
// Amazon Resource Name (ARN) of the finding aggregator. You can run this operation
// from any Amazon Web Services Region.
func securityhub_ListFindingAggregators(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.ListFindingAggregatorsInput{}

	if len(_securityhubMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _securityhubMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_securityhubNextToken) > 0 {
		input.NextToken = aws.String(_securityhubNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListFindingAggregators(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*securityhub.ListFindingAggregatorsOutput
	p := securityhub.NewListFindingAggregatorsPaginator(client, input)
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

// We recommend using Organizations instead of Security Hub CSPM invitations to
// manage your member accounts. For information, see [Managing Security Hub CSPM administrator and member accounts with Organizations]in the Security Hub CSPM User
// Guide.
//
// Lists all Security Hub CSPM membership invitations that were sent to the
// calling account.
//
// Only accounts that are managed by invitation can use this operation. Accounts
// that are managed using the integration with Organizations don't receive
// invitations.
//
// [Managing Security Hub CSPM administrator and member accounts with Organizations]: https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-accounts-orgs.html
func securityhub_ListInvitations(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.ListInvitationsInput{}

	if len(_securityhubMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _securityhubMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_securityhubNextToken) > 0 {
		input.NextToken = aws.String(_securityhubNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListInvitations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*securityhub.ListInvitationsOutput
	p := securityhub.NewListInvitationsPaginator(client, input)
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

// Lists details about all member accounts for the current Security Hub CSPM
// administrator account.
//
// The results include both member accounts that belong to an organization and
// member accounts that were invited manually.
func securityhub_ListMembers(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.ListMembersInput{}

	if len(_securityhubMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _securityhubMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_securityhubNextToken) > 0 {
		input.NextToken = aws.String(_securityhubNextToken)
	}
	if len(_securityhubOnlyAssociated) > 0 {
		if err := assignInputField(input, "OnlyAssociated", _securityhubOnlyAssociated); err != nil {
			log.Errorf("invalid --only-associated: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListMembers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*securityhub.ListMembersOutput
	p := securityhub.NewListMembersPaginator(client, input)
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

// Lists the Security Hub CSPM administrator accounts. Can only be called by the
// organization management account.
func securityhub_ListOrganizationAdminAccounts(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.ListOrganizationAdminAccountsInput{}

	if len(_securityhubFeature) > 0 {
		if err := assignInputField(input, "Feature", _securityhubFeature); err != nil {
			log.Errorf("invalid --feature: %s", err.Error())
			return
		}
	}
	if len(_securityhubMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _securityhubMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_securityhubNextToken) > 0 {
		input.NextToken = aws.String(_securityhubNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListOrganizationAdminAccounts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*securityhub.ListOrganizationAdminAccountsOutput
	p := securityhub.NewListOrganizationAdminAccountsPaginator(client, input)
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

// Lists all of the security controls that apply to a specified standard.
func securityhub_ListSecurityControlDefinitions(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.ListSecurityControlDefinitionsInput{}

	if len(_securityhubMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _securityhubMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_securityhubNextToken) > 0 {
		input.NextToken = aws.String(_securityhubNextToken)
	}
	if len(_securityhubStandardsArn) > 0 {
		input.StandardsArn = aws.String(_securityhubStandardsArn)
	}

	if disablePaginator() {
		if resp, err := client.ListSecurityControlDefinitions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*securityhub.ListSecurityControlDefinitionsOutput
	p := securityhub.NewListSecurityControlDefinitionsPaginator(client, input)
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

// Specifies whether a control is currently enabled or disabled in each enabled
// standard in the calling account.
//
// This operation omits standards control associations for standard subscriptions
// where StandardsControlsUpdatable has value NOT_READY_FOR_UPDATES .
func securityhub_ListStandardsControlAssociations(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.ListStandardsControlAssociationsInput{
		// SecurityControlId: *string, // Required
	}

	if len(_securityhubSecurityControlId) > 0 {
		input.SecurityControlId = aws.String(_securityhubSecurityControlId)
	}
	if len(_securityhubMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _securityhubMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_securityhubNextToken) > 0 {
		input.NextToken = aws.String(_securityhubNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListStandardsControlAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*securityhub.ListStandardsControlAssociationsOutput
	p := securityhub.NewListStandardsControlAssociationsPaginator(client, input)
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

// Returns a list of tags associated with a resource.
func securityhub_ListTagsForResource(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_securityhubResourceArn) > 0 {
		input.ResourceArn = aws.String(_securityhubResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Grants permission to complete the authorization based on input parameters.
func securityhub_RegisterConnectorV2(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.RegisterConnectorV2Input{
		// AuthCode: *string, // Required
		// AuthState: *string, // Required
	}

	if len(_securityhubAuthCode) > 0 {
		input.AuthCode = aws.String(_securityhubAuthCode)
	}
	if len(_securityhubAuthState) > 0 {
		input.AuthState = aws.String(_securityhubAuthState)
	}

	if resp, err := client.RegisterConnectorV2(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates a target account, organizational unit, or the root with a specified
// configuration. The target can be associated with a configuration policy or
// self-managed behavior. Only the Security Hub CSPM delegated administrator can
// invoke this operation from the home Region.
func securityhub_StartConfigurationPolicyAssociation(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.StartConfigurationPolicyAssociationInput{
		// ConfigurationPolicyIdentifier: *string, // Required
		// Target: types.Target, // Required
	}

	if len(_securityhubConfigurationPolicyIdentifier) > 0 {
		input.ConfigurationPolicyIdentifier = aws.String(_securityhubConfigurationPolicyIdentifier)
	}
	if len(_securityhubTarget) > 0 {
		if err := assignInputField(input, "Target", _securityhubTarget); err != nil {
			log.Errorf("invalid --target: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartConfigurationPolicyAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates a target account, organizational unit, or the root from a
// specified configuration. When you disassociate a configuration from its target,
// the target inherits the configuration of the closest parent. If there’s no
// configuration to inherit, the target retains its settings but becomes a
// self-managed account. A target can be disassociated from a configuration policy
// or self-managed behavior. Only the Security Hub CSPM delegated administrator can
// invoke this operation from the home Region.
func securityhub_StartConfigurationPolicyDisassociation(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.StartConfigurationPolicyDisassociationInput{
		// ConfigurationPolicyIdentifier: *string, // Required
	}

	if len(_securityhubConfigurationPolicyIdentifier) > 0 {
		input.ConfigurationPolicyIdentifier = aws.String(_securityhubConfigurationPolicyIdentifier)
	}
	if len(_securityhubTarget) > 0 {
		if err := assignInputField(input, "Target", _securityhubTarget); err != nil {
			log.Errorf("invalid --target: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartConfigurationPolicyDisassociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds one or more tags to a resource.
func securityhub_TagResource(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_securityhubResourceArn) > 0 {
		input.ResourceArn = aws.String(_securityhubResourceArn)
	}
	if len(_securityhubTags) > 0 {
		if err := assignInputField(input, "Tags", _securityhubTags); err != nil {
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

// Removes one or more tags from a resource.
func securityhub_UntagResource(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_securityhubResourceArn) > 0 {
		input.ResourceArn = aws.String(_securityhubResourceArn)
	}
	if len(_securityhubTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _securityhubTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the name and description of a custom action target in Security Hub CSPM.
func securityhub_UpdateActionTarget(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.UpdateActionTargetInput{
		// ActionTargetArn: *string, // Required
	}

	if len(_securityhubActionTargetArn) > 0 {
		input.ActionTargetArn = aws.String(_securityhubActionTargetArn)
	}
	if len(_securityhubDescription) > 0 {
		input.Description = aws.String(_securityhubDescription)
	}
	if len(_securityhubName) > 0 {
		input.Name = aws.String(_securityhubName)
	}

	if resp, err := client.UpdateActionTarget(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Udpates the configuration for the Aggregator V2.
func securityhub_UpdateAggregatorV2(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.UpdateAggregatorV2Input{
		// AggregatorV2Arn: *string, // Required
		// RegionLinkingMode: *string, // Required
	}

	if len(_securityhubAggregatorV2Arn) > 0 {
		input.AggregatorV2Arn = aws.String(_securityhubAggregatorV2Arn)
	}
	if len(_securityhubRegionLinkingMode) > 0 {
		input.RegionLinkingMode = aws.String(_securityhubRegionLinkingMode)
	}
	if len(_securityhubLinkedRegions) > 0 {
		input.LinkedRegions = append([]string(nil), _securityhubLinkedRegions...)
	}

	if resp, err := client.UpdateAggregatorV2(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a V2 automation rule.
func securityhub_UpdateAutomationRuleV2(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.UpdateAutomationRuleV2Input{
		// Identifier: *string, // Required
	}

	if len(_securityhubIdentifier) > 0 {
		input.Identifier = aws.String(_securityhubIdentifier)
	}
	if len(_securityhubActions) > 0 {
		if err := assignInputField(input, "Actions", _securityhubActions); err != nil {
			log.Errorf("invalid --actions: %s", err.Error())
			return
		}
	}
	if len(_securityhubCriteria) > 0 {
		if err := assignInputField(input, "Criteria", _securityhubCriteria); err != nil {
			log.Errorf("invalid --criteria: %s", err.Error())
			return
		}
	}
	if len(_securityhubDescription) > 0 {
		input.Description = aws.String(_securityhubDescription)
	}
	if len(_securityhubRuleName) > 0 {
		input.RuleName = aws.String(_securityhubRuleName)
	}
	if len(_securityhubRuleOrder) > 0 {
		if err := assignInputField(input, "RuleOrder", _securityhubRuleOrder); err != nil {
			log.Errorf("invalid --rule-order: %s", err.Error())
			return
		}
	}
	if len(_securityhubRuleStatus) > 0 {
		if err := assignInputField(input, "RuleStatus", _securityhubRuleStatus); err != nil {
			log.Errorf("invalid --rule-status: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateAutomationRuleV2(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a configuration policy. Only the Security Hub CSPM delegated
// administrator can invoke this operation from the home Region.
func securityhub_UpdateConfigurationPolicy(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.UpdateConfigurationPolicyInput{
		// Identifier: *string, // Required
	}

	if len(_securityhubIdentifier) > 0 {
		input.Identifier = aws.String(_securityhubIdentifier)
	}
	if len(_securityhubConfigurationPolicy) > 0 {
		if err := assignInputField(input, "ConfigurationPolicy", _securityhubConfigurationPolicy); err != nil {
			log.Errorf("invalid --configuration-policy: %s", err.Error())
			return
		}
	}
	if len(_securityhubDescription) > 0 {
		input.Description = aws.String(_securityhubDescription)
	}
	if len(_securityhubName) > 0 {
		input.Name = aws.String(_securityhubName)
	}
	if len(_securityhubUpdatedReason) > 0 {
		input.UpdatedReason = aws.String(_securityhubUpdatedReason)
	}

	if resp, err := client.UpdateConfigurationPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Grants permission to update a connectorV2 based on its id and input parameters.
func securityhub_UpdateConnectorV2(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.UpdateConnectorV2Input{
		// ConnectorId: *string, // Required
	}

	if len(_securityhubConnectorId) > 0 {
		input.ConnectorId = aws.String(_securityhubConnectorId)
	}
	if len(_securityhubDescription) > 0 {
		input.Description = aws.String(_securityhubDescription)
	}
	if len(_securityhubProvider) > 0 {
		if err := assignInputField(input, "Provider", _securityhubProvider); err != nil {
			log.Errorf("invalid --provider: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateConnectorV2(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The aggregation Region is now called the home Region.
// Updates cross-Region aggregation settings. You can use this operation to update
// the Region linking mode and the list of included or excluded Amazon Web Services
// Regions. However, you can't use this operation to change the home Region.
//
// You can invoke this operation from the current home Region only.
func securityhub_UpdateFindingAggregator(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.UpdateFindingAggregatorInput{
		// FindingAggregatorArn: *string, // Required
		// RegionLinkingMode: *string, // Required
	}

	if len(_securityhubFindingAggregatorArn) > 0 {
		input.FindingAggregatorArn = aws.String(_securityhubFindingAggregatorArn)
	}
	if len(_securityhubRegionLinkingMode) > 0 {
		input.RegionLinkingMode = aws.String(_securityhubRegionLinkingMode)
	}
	if len(_securityhubRegions) > 0 {
		input.Regions = append([]string(nil), _securityhubRegions...)
	}

	if resp, err := client.UpdateFindingAggregator(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// UpdateFindings is a deprecated operation. Instead of UpdateFindings , use the
// BatchUpdateFindings operation.
//
// The UpdateFindings operation updates the Note and RecordState of the Security
// Hub CSPM aggregated findings that the filter attributes specify. Any member
// account that can view the finding can also see the update to the finding.
//
// Finding updates made with UpdateFindings aren't persisted if the same finding
// is later updated by the finding provider through the BatchImportFindings
// operation. In addition, Security Hub CSPM doesn't record updates made with
// UpdateFindings in the finding history.
func securityhub_UpdateFindings(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.UpdateFindingsInput{
		// Filters: *types.AwsSecurityFindingFilters, // Required
	}

	if len(_securityhubFilters) > 0 {
		if err := assignInputField(input, "Filters", _securityhubFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_securityhubNote) > 0 {
		if err := assignInputField(input, "Note", _securityhubNote); err != nil {
			log.Errorf("invalid --note: %s", err.Error())
			return
		}
	}
	if len(_securityhubRecordState) > 0 {
		if err := assignInputField(input, "RecordState", _securityhubRecordState); err != nil {
			log.Errorf("invalid --record-state: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateFindings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the Security Hub CSPM insight identified by the specified insight ARN.
func securityhub_UpdateInsight(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.UpdateInsightInput{
		// InsightArn: *string, // Required
	}

	if len(_securityhubInsightArn) > 0 {
		input.InsightArn = aws.String(_securityhubInsightArn)
	}
	if len(_securityhubFilters) > 0 {
		if err := assignInputField(input, "Filters", _securityhubFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_securityhubGroupByAttribute) > 0 {
		input.GroupByAttribute = aws.String(_securityhubGroupByAttribute)
	}
	if len(_securityhubName) > 0 {
		input.Name = aws.String(_securityhubName)
	}

	if resp, err := client.UpdateInsight(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration of your organization in Security Hub CSPM. Only the
// Security Hub CSPM administrator account can invoke this operation.
func securityhub_UpdateOrganizationConfiguration(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.UpdateOrganizationConfigurationInput{
		// AutoEnable: *bool, // Required
	}

	if len(_securityhubAutoEnable) > 0 {
		if err := assignInputField(input, "AutoEnable", _securityhubAutoEnable); err != nil {
			log.Errorf("invalid --auto-enable: %s", err.Error())
			return
		}
	}
	if len(_securityhubAutoEnableStandards) > 0 {
		if err := assignInputField(input, "AutoEnableStandards", _securityhubAutoEnableStandards); err != nil {
			log.Errorf("invalid --auto-enable-standards: %s", err.Error())
			return
		}
	}
	if len(_securityhubOrganizationConfiguration) > 0 {
		if err := assignInputField(input, "OrganizationConfiguration", _securityhubOrganizationConfiguration); err != nil {
			log.Errorf("invalid --organization-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateOrganizationConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the properties of a security control.
func securityhub_UpdateSecurityControl(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.UpdateSecurityControlInput{
		// Parameters: map[string]types.ParameterConfiguration, // Required
		// SecurityControlId: *string, // Required
	}

	if len(_securityhubParameters) > 0 {
		if err := assignInputField(input, "Parameters", _securityhubParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}
	if len(_securityhubSecurityControlId) > 0 {
		input.SecurityControlId = aws.String(_securityhubSecurityControlId)
	}
	if len(_securityhubLastUpdateReason) > 0 {
		input.LastUpdateReason = aws.String(_securityhubLastUpdateReason)
	}

	if resp, err := client.UpdateSecurityControl(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates configuration options for Security Hub CSPM.
func securityhub_UpdateSecurityHubConfiguration(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.UpdateSecurityHubConfigurationInput{}

	if len(_securityhubAutoEnableControls) > 0 {
		if err := assignInputField(input, "AutoEnableControls", _securityhubAutoEnableControls); err != nil {
			log.Errorf("invalid --auto-enable-controls: %s", err.Error())
			return
		}
	}
	if len(_securityhubControlFindingGenerator) > 0 {
		if err := assignInputField(input, "ControlFindingGenerator", _securityhubControlFindingGenerator); err != nil {
			log.Errorf("invalid --control-finding-generator: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateSecurityHubConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Used to control whether an individual security standard control is enabled or
// disabled.
//
// Calls to this operation return a RESOURCE_NOT_FOUND_EXCEPTION error when the
// standard subscription for the control has StandardsControlsUpdatable value
// NOT_READY_FOR_UPDATES .
func securityhub_UpdateStandardsControl(cfg aws.Config, client *securityhub.Client) {
	input := &securityhub.UpdateStandardsControlInput{
		// StandardsControlArn: *string, // Required
	}

	if len(_securityhubStandardsControlArn) > 0 {
		input.StandardsControlArn = aws.String(_securityhubStandardsControlArn)
	}
	if len(_securityhubControlStatus) > 0 {
		if err := assignInputField(input, "ControlStatus", _securityhubControlStatus); err != nil {
			log.Errorf("invalid --control-status: %s", err.Error())
			return
		}
	}
	if len(_securityhubDisabledReason) > 0 {
		input.DisabledReason = aws.String(_securityhubDisabledReason)
	}

	if resp, err := client.UpdateStandardsControl(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_securityhubCmd)
	_securityhubCmd.Flags().SortFlags = false

	_securityhubCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_securityhubCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_securityhubCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_securityhubCmd.Flags().StringVarP(&_securityhubAccountDetails, "account-details", "", "", "Account Details")
	_securityhubCmd.Flags().StringSliceVarP(&_securityhubAccountIds, "account-ids", "", nil, "Account Ids")
	_securityhubCmd.Flags().StringVarP(&_securityhubActionTargetArn, "action-target-arn", "", "", "Action Target ARN")
	_securityhubCmd.Flags().StringSliceVarP(&_securityhubActionTargetArns, "action-target-arns", "", nil, "Action Target Arns")
	_securityhubCmd.Flags().StringVarP(&_securityhubActions, "actions", "", "", "Actions")
	_securityhubCmd.Flags().StringVarP(&_securityhubAdminAccountId, "admin-account-id", "", "", "Admin Account ID")
	_securityhubCmd.Flags().StringVarP(&_securityhubAdministratorId, "administrator-id", "", "", "Administrator ID")
	_securityhubCmd.Flags().StringVarP(&_securityhubAggregatorV2Arn, "aggregator-v2-arn", "", "", "Aggregator V2 ARN")
	_securityhubCmd.Flags().StringVarP(&_securityhubAuthCode, "auth-code", "", "", "Auth Code")
	_securityhubCmd.Flags().StringVarP(&_securityhubAuthState, "auth-state", "", "", "Auth State")
	_securityhubCmd.Flags().StringVarP(&_securityhubAutoEnable, "auto-enable", "", "", "Auto Enable")
	_securityhubCmd.Flags().StringVarP(&_securityhubAutoEnableControls, "auto-enable-controls", "", "", "Auto Enable Controls")
	_securityhubCmd.Flags().StringVarP(&_securityhubAutoEnableStandards, "auto-enable-standards", "", "", "Auto Enable Standards")
	_securityhubCmd.Flags().StringSliceVarP(&_securityhubAutomationRulesArns, "automation-rules-arns", "", nil, "Automation Rules Arns")
	_securityhubCmd.Flags().StringVarP(&_securityhubClientToken, "client-token", "", "", "Client Token")
	_securityhubCmd.Flags().StringVarP(&_securityhubComment, "comment", "", "", "Comment")
	_securityhubCmd.Flags().StringVarP(&_securityhubConfidence, "confidence", "", "", "Confidence")
	_securityhubCmd.Flags().StringVarP(&_securityhubConfigurationPolicy, "configuration-policy", "", "", "Configuration Policy")
	_securityhubCmd.Flags().StringVarP(&_securityhubConfigurationPolicyAssociationIdentifiers, "configuration-policy-association-identifiers", "", "", "Configuration Policy Association Identifiers")
	_securityhubCmd.Flags().StringVarP(&_securityhubConfigurationPolicyIdentifier, "configuration-policy-identifier", "", "", "Configuration Policy Identifier")
	_securityhubCmd.Flags().StringVarP(&_securityhubConnectorId, "connector-id", "", "", "Connector ID")
	_securityhubCmd.Flags().StringVarP(&_securityhubConnectorStatus, "connector-status", "", "", "Connector Status")
	_securityhubCmd.Flags().StringVarP(&_securityhubControlFindingGenerator, "control-finding-generator", "", "", "Control Finding Generator")
	_securityhubCmd.Flags().StringVarP(&_securityhubControlStatus, "control-status", "", "", "Control Status")
	_securityhubCmd.Flags().StringVarP(&_securityhubCriteria, "criteria", "", "", "Criteria")
	_securityhubCmd.Flags().StringVarP(&_securityhubCriticality, "criticality", "", "", "Criticality")
	_securityhubCmd.Flags().StringVarP(&_securityhubDescription, "description", "", "", "Description")
	_securityhubCmd.Flags().StringVarP(&_securityhubDisabledReason, "disabled-reason", "", "", "Disabled Reason")
	_securityhubCmd.Flags().StringVarP(&_securityhubEnableDefaultStandards, "enable-default-standards", "", "", "Enable Default Standards")
	_securityhubCmd.Flags().StringVarP(&_securityhubEndTime, "end-time", "", "", "End Time")
	_securityhubCmd.Flags().StringVarP(&_securityhubFeature, "feature", "", "", "Feature")
	_securityhubCmd.Flags().StringVarP(&_securityhubFilters, "filters", "", "", "Filters")
	_securityhubCmd.Flags().StringVarP(&_securityhubFindingAggregatorArn, "finding-aggregator-arn", "", "", "Finding Aggregator ARN")
	_securityhubCmd.Flags().StringVarP(&_securityhubFindingIdentifier, "finding-identifier", "", "", "Finding Identifier")
	_securityhubCmd.Flags().StringVarP(&_securityhubFindingIdentifiers, "finding-identifiers", "", "", "Finding Identifiers")
	_securityhubCmd.Flags().StringVarP(&_securityhubFindingMetadataUid, "finding-metadata-uid", "", "", "Finding Metadata UID")
	_securityhubCmd.Flags().StringVarP(&_securityhubFindings, "findings", "", "", "Findings")
	_securityhubCmd.Flags().StringVarP(&_securityhubGroupByAttribute, "group-by-attribute", "", "", "Group By Attribute")
	_securityhubCmd.Flags().StringVarP(&_securityhubGroupByRules, "group-by-rules", "", "", "Group By Rules")
	_securityhubCmd.Flags().StringVarP(&_securityhubHubArn, "hub-arn", "", "", "Hub ARN")
	_securityhubCmd.Flags().StringVarP(&_securityhubId, "id", "", "", "ID")
	_securityhubCmd.Flags().StringVarP(&_securityhubIdentifier, "identifier", "", "", "Identifier")
	_securityhubCmd.Flags().StringVarP(&_securityhubInsightArn, "insight-arn", "", "", "Insight ARN")
	_securityhubCmd.Flags().StringSliceVarP(&_securityhubInsightArns, "insight-arns", "", nil, "Insight Arns")
	_securityhubCmd.Flags().StringVarP(&_securityhubInvitationId, "invitation-id", "", "", "Invitation ID")
	_securityhubCmd.Flags().StringVarP(&_securityhubIsTerminal, "is-terminal", "", "", "Is Terminal")
	_securityhubCmd.Flags().StringVarP(&_securityhubKmsKeyArn, "kms-key-arn", "", "", "KMS Key ARN")
	_securityhubCmd.Flags().StringVarP(&_securityhubLastUpdateReason, "last-update-reason", "", "", "Last Update Reason")
	_securityhubCmd.Flags().StringSliceVarP(&_securityhubLinkedRegions, "linked-regions", "", nil, "Linked Regions")
	_securityhubCmd.Flags().StringVarP(&_securityhubMasterId, "master-id", "", "", "Master ID")
	_securityhubCmd.Flags().StringVarP(&_securityhubMaxResults, "max-results", "", "", "Max Results")
	_securityhubCmd.Flags().StringVarP(&_securityhubMaxStatisticResults, "max-statistic-results", "", "", "Max Statistic Results")
	_securityhubCmd.Flags().StringSliceVarP(&_securityhubMetadataUids, "metadata-uids", "", nil, "Metadata Uids")
	_securityhubCmd.Flags().StringVarP(&_securityhubMode, "mode", "", "", "Mode")
	_securityhubCmd.Flags().StringVarP(&_securityhubName, "name", "", "", "Name")
	_securityhubCmd.Flags().StringVarP(&_securityhubNextToken, "next-token", "", "", "Next Token")
	_securityhubCmd.Flags().StringVarP(&_securityhubNote, "note", "", "", "Note")
	_securityhubCmd.Flags().StringVarP(&_securityhubOnlyAssociated, "only-associated", "", "", "Only Associated")
	_securityhubCmd.Flags().StringVarP(&_securityhubOrganizationConfiguration, "organization-configuration", "", "", "Organization Configuration")
	_securityhubCmd.Flags().StringVarP(&_securityhubParameters, "parameters", "", "", "Parameters")
	_securityhubCmd.Flags().StringVarP(&_securityhubProductArn, "product-arn", "", "", "Product ARN")
	_securityhubCmd.Flags().StringVarP(&_securityhubProductSubscriptionArn, "product-subscription-arn", "", "", "Product Subscription ARN")
	_securityhubCmd.Flags().StringVarP(&_securityhubProvider, "provider", "", "", "Provider")
	_securityhubCmd.Flags().StringVarP(&_securityhubProviderName, "provider-name", "", "", "Provider Name")
	_securityhubCmd.Flags().StringVarP(&_securityhubRecordState, "record-state", "", "", "Record State")
	_securityhubCmd.Flags().StringVarP(&_securityhubRegionLinkingMode, "region-linking-mode", "", "", "Region Linking Mode")
	_securityhubCmd.Flags().StringSliceVarP(&_securityhubRegions, "regions", "", nil, "Regions")
	_securityhubCmd.Flags().StringVarP(&_securityhubRelatedFindings, "related-findings", "", "", "Related Findings")
	_securityhubCmd.Flags().StringVarP(&_securityhubResourceArn, "resource-arn", "", "", "Resource ARN")
	_securityhubCmd.Flags().StringVarP(&_securityhubRuleName, "rule-name", "", "", "Rule Name")
	_securityhubCmd.Flags().StringVarP(&_securityhubRuleOrder, "rule-order", "", "", "Rule Order")
	_securityhubCmd.Flags().StringVarP(&_securityhubRuleStatus, "rule-status", "", "", "Rule Status")
	_securityhubCmd.Flags().StringVarP(&_securityhubSecurityControlId, "security-control-id", "", "", "Security Control ID")
	_securityhubCmd.Flags().StringSliceVarP(&_securityhubSecurityControlIds, "security-control-ids", "", nil, "Security Control Ids")
	_securityhubCmd.Flags().StringVarP(&_securityhubSeverity, "severity", "", "", "Severity")
	_securityhubCmd.Flags().StringVarP(&_securityhubSeverityId, "severity-id", "", "", "Severity ID")
	_securityhubCmd.Flags().StringVarP(&_securityhubSortCriteria, "sort-criteria", "", "", "Sort Criteria")
	_securityhubCmd.Flags().StringVarP(&_securityhubSortOrder, "sort-order", "", "", "Sort Order")
	_securityhubCmd.Flags().StringVarP(&_securityhubStandardsArn, "standards-arn", "", "", "Standards ARN")
	_securityhubCmd.Flags().StringVarP(&_securityhubStandardsControlArn, "standards-control-arn", "", "", "Standards Control ARN")
	_securityhubCmd.Flags().StringVarP(&_securityhubStandardsControlAssociationIds, "standards-control-association-ids", "", "", "Standards Control Association Ids")
	_securityhubCmd.Flags().StringVarP(&_securityhubStandardsControlAssociationUpdates, "standards-control-association-updates", "", "", "Standards Control Association Updates")
	_securityhubCmd.Flags().StringVarP(&_securityhubStandardsSubscriptionArn, "standards-subscription-arn", "", "", "Standards Subscription ARN")
	_securityhubCmd.Flags().StringSliceVarP(&_securityhubStandardsSubscriptionArns, "standards-subscription-arns", "", nil, "Standards Subscription Arns")
	_securityhubCmd.Flags().StringVarP(&_securityhubStandardsSubscriptionRequests, "standards-subscription-requests", "", "", "Standards Subscription Requests")
	_securityhubCmd.Flags().StringVarP(&_securityhubStartTime, "start-time", "", "", "Start Time")
	_securityhubCmd.Flags().StringVarP(&_securityhubStatusId, "status-id", "", "", "Status ID")
	_securityhubCmd.Flags().StringSliceVarP(&_securityhubTagKeys, "tag-keys", "", nil, "Tag Keys")
	_securityhubCmd.Flags().StringVarP(&_securityhubTags, "tags", "", "", "Tags")
	_securityhubCmd.Flags().StringVarP(&_securityhubTarget, "target", "", "", "Target")
	_securityhubCmd.Flags().StringSliceVarP(&_securityhubTypes, "types", "", nil, "Types")
	_securityhubCmd.Flags().StringVarP(&_securityhubUpdateAutomationRulesRequestItems, "update-automation-rules-request-items", "", "", "Update Automation Rules Request Items")
	_securityhubCmd.Flags().StringVarP(&_securityhubUpdatedReason, "updated-reason", "", "", "Updated Reason")
	_securityhubCmd.Flags().StringVarP(&_securityhubUserDefinedFields, "user-defined-fields", "", "", "User Defined Fields")
	_securityhubCmd.Flags().StringVarP(&_securityhubVerificationState, "verification-state", "", "", "Verification State")
	_securityhubCmd.Flags().StringVarP(&_securityhubWorkflow, "workflow", "", "", "Workflow")

	_securityhubCmd.Flags().BoolVarP(&_securityhubAcceptAdministratorInvitation, "accept-administrator-invitation", "", false, "Accept Administrator Invitation")
	_securityhubCmd.Flags().BoolVarP(&_securityhubAcceptInvitation, "accept-invitation", "", false, "Accept Invitation")
	_securityhubCmd.Flags().BoolVarP(&_securityhubBatchDeleteAutomationRules, "batch-delete-automation-rules", "", false, "Batch Delete Automation Rules")
	_securityhubCmd.Flags().BoolVarP(&_securityhubBatchDisableStandards, "batch-disable-standards", "", false, "Batch Disable Standards")
	_securityhubCmd.Flags().BoolVarP(&_securityhubBatchEnableStandards, "batch-enable-standards", "", false, "Batch Enable Standards")
	_securityhubCmd.Flags().BoolVarP(&_securityhubBatchGetAutomationRules, "batch-get-automation-rules", "", false, "Batch Get Automation Rules")
	_securityhubCmd.Flags().BoolVarP(&_securityhubBatchGetConfigurationPolicyAssociations, "batch-get-configuration-policy-associations", "", false, "Batch Get Configuration Policy Associations")
	_securityhubCmd.Flags().BoolVarP(&_securityhubBatchGetSecurityControls, "batch-get-security-controls", "", false, "Batch Get Security Controls")
	_securityhubCmd.Flags().BoolVarP(&_securityhubBatchGetStandardsControlAssociations, "batch-get-standards-control-associations", "", false, "Batch Get Standards Control Associations")
	_securityhubCmd.Flags().BoolVarP(&_securityhubBatchImportFindings, "batch-import-findings", "", false, "Batch Import Findings")
	_securityhubCmd.Flags().BoolVarP(&_securityhubBatchUpdateAutomationRules, "batch-update-automation-rules", "", false, "Batch Update Automation Rules")
	_securityhubCmd.Flags().BoolVarP(&_securityhubBatchUpdateFindings, "batch-update-findings", "", false, "Batch Update Findings")
	_securityhubCmd.Flags().BoolVarP(&_securityhubBatchUpdateFindingsV2, "batch-update-findings-v2", "", false, "Batch Update Findings V2")
	_securityhubCmd.Flags().BoolVarP(&_securityhubBatchUpdateStandardsControlAssociations, "batch-update-standards-control-associations", "", false, "Batch Update Standards Control Associations")
	_securityhubCmd.Flags().BoolVarP(&_securityhubCreateActionTarget, "create-action-target", "", false, "Create Action Target")
	_securityhubCmd.Flags().BoolVarP(&_securityhubCreateAggregatorV2, "create-aggregator-v2", "", false, "Create Aggregator V2")
	_securityhubCmd.Flags().BoolVarP(&_securityhubCreateAutomationRule, "create-automation-rule", "", false, "Create Automation Rule")
	_securityhubCmd.Flags().BoolVarP(&_securityhubCreateAutomationRuleV2, "create-automation-rule-v2", "", false, "Create Automation Rule V2")
	_securityhubCmd.Flags().BoolVarP(&_securityhubCreateConfigurationPolicy, "create-configuration-policy", "", false, "Create Configuration Policy")
	_securityhubCmd.Flags().BoolVarP(&_securityhubCreateConnectorV2, "create-connector-v2", "", false, "Create Connector V2")
	_securityhubCmd.Flags().BoolVarP(&_securityhubCreateFindingAggregator, "create-finding-aggregator", "", false, "Create Finding Aggregator")
	_securityhubCmd.Flags().BoolVarP(&_securityhubCreateInsight, "create-insight", "", false, "Create Insight")
	_securityhubCmd.Flags().BoolVarP(&_securityhubCreateMembers, "create-members", "", false, "Create Members")
	_securityhubCmd.Flags().BoolVarP(&_securityhubCreateTicketV2, "create-ticket-v2", "", false, "Create Ticket V2")
	_securityhubCmd.Flags().BoolVarP(&_securityhubDeclineInvitations, "decline-invitations", "", false, "Decline Invitations")
	_securityhubCmd.Flags().BoolVarP(&_securityhubDeleteActionTarget, "delete-action-target", "", false, "Delete Action Target")
	_securityhubCmd.Flags().BoolVarP(&_securityhubDeleteAggregatorV2, "delete-aggregator-v2", "", false, "Delete Aggregator V2")
	_securityhubCmd.Flags().BoolVarP(&_securityhubDeleteAutomationRuleV2, "delete-automation-rule-v2", "", false, "Delete Automation Rule V2")
	_securityhubCmd.Flags().BoolVarP(&_securityhubDeleteConfigurationPolicy, "delete-configuration-policy", "", false, "Delete Configuration Policy")
	_securityhubCmd.Flags().BoolVarP(&_securityhubDeleteConnectorV2, "delete-connector-v2", "", false, "Delete Connector V2")
	_securityhubCmd.Flags().BoolVarP(&_securityhubDeleteFindingAggregator, "delete-finding-aggregator", "", false, "Delete Finding Aggregator")
	_securityhubCmd.Flags().BoolVarP(&_securityhubDeleteInsight, "delete-insight", "", false, "Delete Insight")
	_securityhubCmd.Flags().BoolVarP(&_securityhubDeleteInvitations, "delete-invitations", "", false, "Delete Invitations")
	_securityhubCmd.Flags().BoolVarP(&_securityhubDeleteMembers, "delete-members", "", false, "Delete Members")
	_securityhubCmd.Flags().BoolVarP(&_securityhubDescribeActionTargets, "describe-action-targets", "", false, "Describe Action Targets")
	_securityhubCmd.Flags().BoolVarP(&_securityhubDescribeHub, "describe-hub", "", false, "Describe Hub")
	_securityhubCmd.Flags().BoolVarP(&_securityhubDescribeOrganizationConfiguration, "describe-organization-configuration", "", false, "Describe Organization Configuration")
	_securityhubCmd.Flags().BoolVarP(&_securityhubDescribeProducts, "describe-products", "", false, "Describe Products")
	_securityhubCmd.Flags().BoolVarP(&_securityhubDescribeProductsV2, "describe-products-v2", "", false, "Describe Products V2")
	_securityhubCmd.Flags().BoolVarP(&_securityhubDescribeSecurityHubV2, "describe-security-hub-v2", "", false, "Describe Security Hub V2")
	_securityhubCmd.Flags().BoolVarP(&_securityhubDescribeStandards, "describe-standards", "", false, "Describe Standards")
	_securityhubCmd.Flags().BoolVarP(&_securityhubDescribeStandardsControls, "describe-standards-controls", "", false, "Describe Standards Controls")
	_securityhubCmd.Flags().BoolVarP(&_securityhubDisableImportFindingsForProduct, "disable-import-findings-for-product", "", false, "Disable Import Findings For Product")
	_securityhubCmd.Flags().BoolVarP(&_securityhubDisableOrganizationAdminAccount, "disable-organization-admin-account", "", false, "Disable Organization Admin Account")
	_securityhubCmd.Flags().BoolVarP(&_securityhubDisableSecurityHub, "disable-security-hub", "", false, "Disable Security Hub")
	_securityhubCmd.Flags().BoolVarP(&_securityhubDisableSecurityHubV2, "disable-security-hub-v2", "", false, "Disable Security Hub V2")
	_securityhubCmd.Flags().BoolVarP(&_securityhubDisassociateFromAdministratorAccount, "disassociate-from-administrator-account", "", false, "Disassociate From Administrator Account")
	_securityhubCmd.Flags().BoolVarP(&_securityhubDisassociateFromMasterAccount, "disassociate-from-master-account", "", false, "Disassociate From Master Account")
	_securityhubCmd.Flags().BoolVarP(&_securityhubDisassociateMembers, "disassociate-members", "", false, "Disassociate Members")
	_securityhubCmd.Flags().BoolVarP(&_securityhubEnableImportFindingsForProduct, "enable-import-findings-for-product", "", false, "Enable Import Findings For Product")
	_securityhubCmd.Flags().BoolVarP(&_securityhubEnableOrganizationAdminAccount, "enable-organization-admin-account", "", false, "Enable Organization Admin Account")
	_securityhubCmd.Flags().BoolVarP(&_securityhubEnableSecurityHub, "enable-security-hub", "", false, "Enable Security Hub")
	_securityhubCmd.Flags().BoolVarP(&_securityhubEnableSecurityHubV2, "enable-security-hub-v2", "", false, "Enable Security Hub V2")
	_securityhubCmd.Flags().BoolVarP(&_securityhubGetAdministratorAccount, "get-administrator-account", "", false, "Get Administrator Account")
	_securityhubCmd.Flags().BoolVarP(&_securityhubGetAggregatorV2, "get-aggregator-v2", "", false, "Get Aggregator V2")
	_securityhubCmd.Flags().BoolVarP(&_securityhubGetAutomationRuleV2, "get-automation-rule-v2", "", false, "Get Automation Rule V2")
	_securityhubCmd.Flags().BoolVarP(&_securityhubGetConfigurationPolicy, "get-configuration-policy", "", false, "Get Configuration Policy")
	_securityhubCmd.Flags().BoolVarP(&_securityhubGetConfigurationPolicyAssociation, "get-configuration-policy-association", "", false, "Get Configuration Policy Association")
	_securityhubCmd.Flags().BoolVarP(&_securityhubGetConnectorV2, "get-connector-v2", "", false, "Get Connector V2")
	_securityhubCmd.Flags().BoolVarP(&_securityhubGetEnabledStandards, "get-enabled-standards", "", false, "Get Enabled Standards")
	_securityhubCmd.Flags().BoolVarP(&_securityhubGetFindingAggregator, "get-finding-aggregator", "", false, "Get Finding Aggregator")
	_securityhubCmd.Flags().BoolVarP(&_securityhubGetFindingHistory, "get-finding-history", "", false, "Get Finding History")
	_securityhubCmd.Flags().BoolVarP(&_securityhubGetFindingStatisticsV2, "get-finding-statistics-v2", "", false, "Get Finding Statistics V2")
	_securityhubCmd.Flags().BoolVarP(&_securityhubGetFindings, "get-findings", "", false, "Get Findings")
	_securityhubCmd.Flags().BoolVarP(&_securityhubGetFindingsTrendsV2, "get-findings-trends-v2", "", false, "Get Findings Trends V2")
	_securityhubCmd.Flags().BoolVarP(&_securityhubGetFindingsV2, "get-findings-v2", "", false, "Get Findings V2")
	_securityhubCmd.Flags().BoolVarP(&_securityhubGetInsightResults, "get-insight-results", "", false, "Get Insight Results")
	_securityhubCmd.Flags().BoolVarP(&_securityhubGetInsights, "get-insights", "", false, "Get Insights")
	_securityhubCmd.Flags().BoolVarP(&_securityhubGetInvitationsCount, "get-invitations-count", "", false, "Get Invitations Count")
	_securityhubCmd.Flags().BoolVarP(&_securityhubGetMasterAccount, "get-master-account", "", false, "Get Master Account")
	_securityhubCmd.Flags().BoolVarP(&_securityhubGetMembers, "get-members", "", false, "Get Members")
	_securityhubCmd.Flags().BoolVarP(&_securityhubGetResourcesStatisticsV2, "get-resources-statistics-v2", "", false, "Get Resources Statistics V2")
	_securityhubCmd.Flags().BoolVarP(&_securityhubGetResourcesTrendsV2, "get-resources-trends-v2", "", false, "Get Resources Trends V2")
	_securityhubCmd.Flags().BoolVarP(&_securityhubGetResourcesV2, "get-resources-v2", "", false, "Get Resources V2")
	_securityhubCmd.Flags().BoolVarP(&_securityhubGetSecurityControlDefinition, "get-security-control-definition", "", false, "Get Security Control Definition")
	_securityhubCmd.Flags().BoolVarP(&_securityhubInviteMembers, "invite-members", "", false, "Invite Members")
	_securityhubCmd.Flags().BoolVarP(&_securityhubListAggregatorsV2, "list-aggregators-v2", "", false, "List Aggregators V2")
	_securityhubCmd.Flags().BoolVarP(&_securityhubListAutomationRules, "list-automation-rules", "", false, "List Automation Rules")
	_securityhubCmd.Flags().BoolVarP(&_securityhubListAutomationRulesV2, "list-automation-rules-v2", "", false, "List Automation Rules V2")
	_securityhubCmd.Flags().BoolVarP(&_securityhubListConfigurationPolicies, "list-configuration-policies", "", false, "List Configuration Policies")
	_securityhubCmd.Flags().BoolVarP(&_securityhubListConfigurationPolicyAssociations, "list-configuration-policy-associations", "", false, "List Configuration Policy Associations")
	_securityhubCmd.Flags().BoolVarP(&_securityhubListConnectorsV2, "list-connectors-v2", "", false, "List Connectors V2")
	_securityhubCmd.Flags().BoolVarP(&_securityhubListEnabledProductsForImport, "list-enabled-products-for-import", "", false, "List Enabled Products For Import")
	_securityhubCmd.Flags().BoolVarP(&_securityhubListFindingAggregators, "list-finding-aggregators", "", false, "List Finding Aggregators")
	_securityhubCmd.Flags().BoolVarP(&_securityhubListInvitations, "list-invitations", "", false, "List Invitations")
	_securityhubCmd.Flags().BoolVarP(&_securityhubListMembers, "list-members", "", false, "List Members")
	_securityhubCmd.Flags().BoolVarP(&_securityhubListOrganizationAdminAccounts, "list-organization-admin-accounts", "", false, "List Organization Admin Accounts")
	_securityhubCmd.Flags().BoolVarP(&_securityhubListSecurityControlDefinitions, "list-security-control-definitions", "", false, "List Security Control Definitions")
	_securityhubCmd.Flags().BoolVarP(&_securityhubListStandardsControlAssociations, "list-standards-control-associations", "", false, "List Standards Control Associations")
	_securityhubCmd.Flags().BoolVarP(&_securityhubListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_securityhubCmd.Flags().BoolVarP(&_securityhubRegisterConnectorV2, "register-connector-v2", "", false, "Register Connector V2")
	_securityhubCmd.Flags().BoolVarP(&_securityhubStartConfigurationPolicyAssociation, "start-configuration-policy-association", "", false, "Start Configuration Policy Association")
	_securityhubCmd.Flags().BoolVarP(&_securityhubStartConfigurationPolicyDisassociation, "start-configuration-policy-disassociation", "", false, "Start Configuration Policy Disassociation")
	_securityhubCmd.Flags().BoolVarP(&_securityhubTagResource, "tag-resource", "", false, "Tag Resource")
	_securityhubCmd.Flags().BoolVarP(&_securityhubUntagResource, "untag-resource", "", false, "Untag Resource")
	_securityhubCmd.Flags().BoolVarP(&_securityhubUpdateActionTarget, "update-action-target", "", false, "Update Action Target")
	_securityhubCmd.Flags().BoolVarP(&_securityhubUpdateAggregatorV2, "update-aggregator-v2", "", false, "Update Aggregator V2")
	_securityhubCmd.Flags().BoolVarP(&_securityhubUpdateAutomationRuleV2, "update-automation-rule-v2", "", false, "Update Automation Rule V2")
	_securityhubCmd.Flags().BoolVarP(&_securityhubUpdateConfigurationPolicy, "update-configuration-policy", "", false, "Update Configuration Policy")
	_securityhubCmd.Flags().BoolVarP(&_securityhubUpdateConnectorV2, "update-connector-v2", "", false, "Update Connector V2")
	_securityhubCmd.Flags().BoolVarP(&_securityhubUpdateFindingAggregator, "update-finding-aggregator", "", false, "Update Finding Aggregator")
	_securityhubCmd.Flags().BoolVarP(&_securityhubUpdateFindings, "update-findings", "", false, "Update Findings")
	_securityhubCmd.Flags().BoolVarP(&_securityhubUpdateInsight, "update-insight", "", false, "Update Insight")
	_securityhubCmd.Flags().BoolVarP(&_securityhubUpdateOrganizationConfiguration, "update-organization-configuration", "", false, "Update Organization Configuration")
	_securityhubCmd.Flags().BoolVarP(&_securityhubUpdateSecurityControl, "update-security-control", "", false, "Update Security Control")
	_securityhubCmd.Flags().BoolVarP(&_securityhubUpdateSecurityHubConfiguration, "update-security-hub-configuration", "", false, "Update Security Hub Configuration")
	_securityhubCmd.Flags().BoolVarP(&_securityhubUpdateStandardsControl, "update-standards-control", "", false, "Update Standards Control")

}
