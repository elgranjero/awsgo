package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/guardduty"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// guarddutyCmd represents the guardduty command
var _guarddutyCmd = &cobra.Command{
	Use:   "guardduty",
	Short: "AWS guardduty CLI",
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
		client := guardduty.NewFromConfig(cfg)
		if _guarddutyAcceptAdministratorInvitation {
			guardduty_AcceptAdministratorInvitation(cfg, client)
			return
		}
		if _guarddutyAcceptInvitation {
			guardduty_AcceptInvitation(cfg, client)
			return
		}
		if _guarddutyArchiveFindings {
			guardduty_ArchiveFindings(cfg, client)
			return
		}
		if _guarddutyCreateDetector {
			guardduty_CreateDetector(cfg, client)
			return
		}
		if _guarddutyCreateFilter {
			guardduty_CreateFilter(cfg, client)
			return
		}
		if _guarddutyCreateIPSet {
			guardduty_CreateIPSet(cfg, client)
			return
		}
		if _guarddutyCreateMalwareProtectionPlan {
			guardduty_CreateMalwareProtectionPlan(cfg, client)
			return
		}
		if _guarddutyCreateMembers {
			guardduty_CreateMembers(cfg, client)
			return
		}
		if _guarddutyCreatePublishingDestination {
			guardduty_CreatePublishingDestination(cfg, client)
			return
		}
		if _guarddutyCreateSampleFindings {
			guardduty_CreateSampleFindings(cfg, client)
			return
		}
		if _guarddutyCreateThreatEntitySet {
			guardduty_CreateThreatEntitySet(cfg, client)
			return
		}
		if _guarddutyCreateThreatIntelSet {
			guardduty_CreateThreatIntelSet(cfg, client)
			return
		}
		if _guarddutyCreateTrustedEntitySet {
			guardduty_CreateTrustedEntitySet(cfg, client)
			return
		}
		if _guarddutyDeclineInvitations {
			guardduty_DeclineInvitations(cfg, client)
			return
		}
		if _guarddutyDeleteDetector {
			guardduty_DeleteDetector(cfg, client)
			return
		}
		if _guarddutyDeleteFilter {
			guardduty_DeleteFilter(cfg, client)
			return
		}
		if _guarddutyDeleteInvitations {
			guardduty_DeleteInvitations(cfg, client)
			return
		}
		if _guarddutyDeleteIPSet {
			guardduty_DeleteIPSet(cfg, client)
			return
		}
		if _guarddutyDeleteMalwareProtectionPlan {
			guardduty_DeleteMalwareProtectionPlan(cfg, client)
			return
		}
		if _guarddutyDeleteMembers {
			guardduty_DeleteMembers(cfg, client)
			return
		}
		if _guarddutyDeletePublishingDestination {
			guardduty_DeletePublishingDestination(cfg, client)
			return
		}
		if _guarddutyDeleteThreatEntitySet {
			guardduty_DeleteThreatEntitySet(cfg, client)
			return
		}
		if _guarddutyDeleteThreatIntelSet {
			guardduty_DeleteThreatIntelSet(cfg, client)
			return
		}
		if _guarddutyDeleteTrustedEntitySet {
			guardduty_DeleteTrustedEntitySet(cfg, client)
			return
		}
		if _guarddutyDescribeMalwareScans {
			guardduty_DescribeMalwareScans(cfg, client)
			return
		}
		if _guarddutyDescribeOrganizationConfiguration {
			guardduty_DescribeOrganizationConfiguration(cfg, client)
			return
		}
		if _guarddutyDescribePublishingDestination {
			guardduty_DescribePublishingDestination(cfg, client)
			return
		}
		if _guarddutyDisableOrganizationAdminAccount {
			guardduty_DisableOrganizationAdminAccount(cfg, client)
			return
		}
		if _guarddutyDisassociateFromAdministratorAccount {
			guardduty_DisassociateFromAdministratorAccount(cfg, client)
			return
		}
		if _guarddutyDisassociateFromMasterAccount {
			guardduty_DisassociateFromMasterAccount(cfg, client)
			return
		}
		if _guarddutyDisassociateMembers {
			guardduty_DisassociateMembers(cfg, client)
			return
		}
		if _guarddutyEnableOrganizationAdminAccount {
			guardduty_EnableOrganizationAdminAccount(cfg, client)
			return
		}
		if _guarddutyGetAdministratorAccount {
			guardduty_GetAdministratorAccount(cfg, client)
			return
		}
		if _guarddutyGetCoverageStatistics {
			guardduty_GetCoverageStatistics(cfg, client)
			return
		}
		if _guarddutyGetDetector {
			guardduty_GetDetector(cfg, client)
			return
		}
		if _guarddutyGetFilter {
			guardduty_GetFilter(cfg, client)
			return
		}
		if _guarddutyGetFindings {
			guardduty_GetFindings(cfg, client)
			return
		}
		if _guarddutyGetFindingsStatistics {
			guardduty_GetFindingsStatistics(cfg, client)
			return
		}
		if _guarddutyGetInvitationsCount {
			guardduty_GetInvitationsCount(cfg, client)
			return
		}
		if _guarddutyGetIPSet {
			guardduty_GetIPSet(cfg, client)
			return
		}
		if _guarddutyGetMalwareProtectionPlan {
			guardduty_GetMalwareProtectionPlan(cfg, client)
			return
		}
		if _guarddutyGetMalwareScan {
			guardduty_GetMalwareScan(cfg, client)
			return
		}
		if _guarddutyGetMalwareScanSettings {
			guardduty_GetMalwareScanSettings(cfg, client)
			return
		}
		if _guarddutyGetMasterAccount {
			guardduty_GetMasterAccount(cfg, client)
			return
		}
		if _guarddutyGetMemberDetectors {
			guardduty_GetMemberDetectors(cfg, client)
			return
		}
		if _guarddutyGetMembers {
			guardduty_GetMembers(cfg, client)
			return
		}
		if _guarddutyGetOrganizationStatistics {
			guardduty_GetOrganizationStatistics(cfg, client)
			return
		}
		if _guarddutyGetRemainingFreeTrialDays {
			guardduty_GetRemainingFreeTrialDays(cfg, client)
			return
		}
		if _guarddutyGetThreatEntitySet {
			guardduty_GetThreatEntitySet(cfg, client)
			return
		}
		if _guarddutyGetThreatIntelSet {
			guardduty_GetThreatIntelSet(cfg, client)
			return
		}
		if _guarddutyGetTrustedEntitySet {
			guardduty_GetTrustedEntitySet(cfg, client)
			return
		}
		if _guarddutyGetUsageStatistics {
			guardduty_GetUsageStatistics(cfg, client)
			return
		}
		if _guarddutyInviteMembers {
			guardduty_InviteMembers(cfg, client)
			return
		}
		if _guarddutyListCoverage {
			guardduty_ListCoverage(cfg, client)
			return
		}
		if _guarddutyListDetectors {
			guardduty_ListDetectors(cfg, client)
			return
		}
		if _guarddutyListFilters {
			guardduty_ListFilters(cfg, client)
			return
		}
		if _guarddutyListFindings {
			guardduty_ListFindings(cfg, client)
			return
		}
		if _guarddutyListInvitations {
			guardduty_ListInvitations(cfg, client)
			return
		}
		if _guarddutyListIPSets {
			guardduty_ListIPSets(cfg, client)
			return
		}
		if _guarddutyListMalwareProtectionPlans {
			guardduty_ListMalwareProtectionPlans(cfg, client)
			return
		}
		if _guarddutyListMalwareScans {
			guardduty_ListMalwareScans(cfg, client)
			return
		}
		if _guarddutyListMembers {
			guardduty_ListMembers(cfg, client)
			return
		}
		if _guarddutyListOrganizationAdminAccounts {
			guardduty_ListOrganizationAdminAccounts(cfg, client)
			return
		}
		if _guarddutyListPublishingDestinations {
			guardduty_ListPublishingDestinations(cfg, client)
			return
		}
		if _guarddutyListTagsForResource {
			guardduty_ListTagsForResource(cfg, client)
			return
		}
		if _guarddutyListThreatEntitySets {
			guardduty_ListThreatEntitySets(cfg, client)
			return
		}
		if _guarddutyListThreatIntelSets {
			guardduty_ListThreatIntelSets(cfg, client)
			return
		}
		if _guarddutyListTrustedEntitySets {
			guardduty_ListTrustedEntitySets(cfg, client)
			return
		}
		if _guarddutySendObjectMalwareScan {
			guardduty_SendObjectMalwareScan(cfg, client)
			return
		}
		if _guarddutyStartMalwareScan {
			guardduty_StartMalwareScan(cfg, client)
			return
		}
		if _guarddutyStartMonitoringMembers {
			guardduty_StartMonitoringMembers(cfg, client)
			return
		}
		if _guarddutyStopMonitoringMembers {
			guardduty_StopMonitoringMembers(cfg, client)
			return
		}
		if _guarddutyTagResource {
			guardduty_TagResource(cfg, client)
			return
		}
		if _guarddutyUnarchiveFindings {
			guardduty_UnarchiveFindings(cfg, client)
			return
		}
		if _guarddutyUntagResource {
			guardduty_UntagResource(cfg, client)
			return
		}
		if _guarddutyUpdateDetector {
			guardduty_UpdateDetector(cfg, client)
			return
		}
		if _guarddutyUpdateFilter {
			guardduty_UpdateFilter(cfg, client)
			return
		}
		if _guarddutyUpdateFindingsFeedback {
			guardduty_UpdateFindingsFeedback(cfg, client)
			return
		}
		if _guarddutyUpdateIPSet {
			guardduty_UpdateIPSet(cfg, client)
			return
		}
		if _guarddutyUpdateMalwareProtectionPlan {
			guardduty_UpdateMalwareProtectionPlan(cfg, client)
			return
		}
		if _guarddutyUpdateMalwareScanSettings {
			guardduty_UpdateMalwareScanSettings(cfg, client)
			return
		}
		if _guarddutyUpdateMemberDetectors {
			guardduty_UpdateMemberDetectors(cfg, client)
			return
		}
		if _guarddutyUpdateOrganizationConfiguration {
			guardduty_UpdateOrganizationConfiguration(cfg, client)
			return
		}
		if _guarddutyUpdatePublishingDestination {
			guardduty_UpdatePublishingDestination(cfg, client)
			return
		}
		if _guarddutyUpdateThreatEntitySet {
			guardduty_UpdateThreatEntitySet(cfg, client)
			return
		}
		if _guarddutyUpdateThreatIntelSet {
			guardduty_UpdateThreatIntelSet(cfg, client)
			return
		}
		if _guarddutyUpdateTrustedEntitySet {
			guardduty_UpdateTrustedEntitySet(cfg, client)
			return
		}

	},
}

var (
	_guarddutyAcceptAdministratorInvitation        bool
	_guarddutyAcceptInvitation                     bool
	_guarddutyArchiveFindings                      bool
	_guarddutyCreateDetector                       bool
	_guarddutyCreateFilter                         bool
	_guarddutyCreateIPSet                          bool
	_guarddutyCreateMalwareProtectionPlan          bool
	_guarddutyCreateMembers                        bool
	_guarddutyCreatePublishingDestination          bool
	_guarddutyCreateSampleFindings                 bool
	_guarddutyCreateThreatEntitySet                bool
	_guarddutyCreateThreatIntelSet                 bool
	_guarddutyCreateTrustedEntitySet               bool
	_guarddutyDeclineInvitations                   bool
	_guarddutyDeleteDetector                       bool
	_guarddutyDeleteFilter                         bool
	_guarddutyDeleteInvitations                    bool
	_guarddutyDeleteIPSet                          bool
	_guarddutyDeleteMalwareProtectionPlan          bool
	_guarddutyDeleteMembers                        bool
	_guarddutyDeletePublishingDestination          bool
	_guarddutyDeleteThreatEntitySet                bool
	_guarddutyDeleteThreatIntelSet                 bool
	_guarddutyDeleteTrustedEntitySet               bool
	_guarddutyDescribeMalwareScans                 bool
	_guarddutyDescribeOrganizationConfiguration    bool
	_guarddutyDescribePublishingDestination        bool
	_guarddutyDisableOrganizationAdminAccount      bool
	_guarddutyDisassociateFromAdministratorAccount bool
	_guarddutyDisassociateFromMasterAccount        bool
	_guarddutyDisassociateMembers                  bool
	_guarddutyEnableOrganizationAdminAccount       bool
	_guarddutyGetAdministratorAccount              bool
	_guarddutyGetCoverageStatistics                bool
	_guarddutyGetDetector                          bool
	_guarddutyGetFilter                            bool
	_guarddutyGetFindings                          bool
	_guarddutyGetFindingsStatistics                bool
	_guarddutyGetInvitationsCount                  bool
	_guarddutyGetIPSet                             bool
	_guarddutyGetMalwareProtectionPlan             bool
	_guarddutyGetMalwareScan                       bool
	_guarddutyGetMalwareScanSettings               bool
	_guarddutyGetMasterAccount                     bool
	_guarddutyGetMemberDetectors                   bool
	_guarddutyGetMembers                           bool
	_guarddutyGetOrganizationStatistics            bool
	_guarddutyGetRemainingFreeTrialDays            bool
	_guarddutyGetThreatEntitySet                   bool
	_guarddutyGetThreatIntelSet                    bool
	_guarddutyGetTrustedEntitySet                  bool
	_guarddutyGetUsageStatistics                   bool
	_guarddutyInviteMembers                        bool
	_guarddutyListCoverage                         bool
	_guarddutyListDetectors                        bool
	_guarddutyListFilters                          bool
	_guarddutyListFindings                         bool
	_guarddutyListInvitations                      bool
	_guarddutyListIPSets                           bool
	_guarddutyListMalwareProtectionPlans           bool
	_guarddutyListMalwareScans                     bool
	_guarddutyListMembers                          bool
	_guarddutyListOrganizationAdminAccounts        bool
	_guarddutyListPublishingDestinations           bool
	_guarddutyListTagsForResource                  bool
	_guarddutyListThreatEntitySets                 bool
	_guarddutyListThreatIntelSets                  bool
	_guarddutyListTrustedEntitySets                bool
	_guarddutySendObjectMalwareScan                bool
	_guarddutyStartMalwareScan                     bool
	_guarddutyStartMonitoringMembers               bool
	_guarddutyStopMonitoringMembers                bool
	_guarddutyTagResource                          bool
	_guarddutyUnarchiveFindings                    bool
	_guarddutyUntagResource                        bool
	_guarddutyUpdateDetector                       bool
	_guarddutyUpdateFilter                         bool
	_guarddutyUpdateFindingsFeedback               bool
	_guarddutyUpdateIPSet                          bool
	_guarddutyUpdateMalwareProtectionPlan          bool
	_guarddutyUpdateMalwareScanSettings            bool
	_guarddutyUpdateMemberDetectors                bool
	_guarddutyUpdateOrganizationConfiguration      bool
	_guarddutyUpdatePublishingDestination          bool
	_guarddutyUpdateThreatEntitySet                bool
	_guarddutyUpdateThreatIntelSet                 bool
	_guarddutyUpdateTrustedEntitySet               bool

	_guarddutyAccountDetails                string
	_guarddutyAccountIds                    []string
	_guarddutyAction                        string
	_guarddutyActions                       string
	_guarddutyActivate                      string
	_guarddutyAdminAccountId                string
	_guarddutyAdministratorId               string
	_guarddutyAutoEnable                    string
	_guarddutyAutoEnableOrganizationMembers string
	_guarddutyClientToken                   string
	_guarddutyComments                      string
	_guarddutyDataSources                   string
	_guarddutyDescription                   string
	_guarddutyDestinationId                 string
	_guarddutyDestinationProperties         string
	_guarddutyDestinationType               string
	_guarddutyDetectorId                    string
	_guarddutyDisableEmailNotification      string
	_guarddutyEbsSnapshotPreservation       string
	_guarddutyEnable                        string
	_guarddutyExpectedBucketOwner           string
	_guarddutyFeatures                      string
	_guarddutyFeedback                      string
	_guarddutyFilterCriteria                string
	_guarddutyFilterName                    string
	_guarddutyFindingCriteria               string
	_guarddutyFindingIds                    []string
	_guarddutyFindingPublishingFrequency    string
	_guarddutyFindingStatisticTypes         string
	_guarddutyFindingTypes                  []string
	_guarddutyFormat                        string
	_guarddutyGroupBy                       string
	_guarddutyInvitationId                  string
	_guarddutyIpSetId                       string
	_guarddutyLocation                      string
	_guarddutyMalwareProtectionPlanId       string
	_guarddutyMasterId                      string
	_guarddutyMaxResults                    string
	_guarddutyMessage                       string
	_guarddutyName                          string
	_guarddutyNextToken                     string
	_guarddutyOnlyAssociated                string
	_guarddutyOrderBy                       string
	_guarddutyProtectedResource             string
	_guarddutyRank                          string
	_guarddutyResourceArn                   string
	_guarddutyRole                          string
	_guarddutyS3Object                      string
	_guarddutyScanConfiguration             string
	_guarddutyScanId                        string
	_guarddutyScanResourceCriteria          string
	_guarddutySortCriteria                  string
	_guarddutyStatisticsType                string
	_guarddutyTagKeys                       []string
	_guarddutyTags                          string
	_guarddutyThreatEntitySetId             string
	_guarddutyThreatIntelSetId              string
	_guarddutyTrustedEntitySetId            string
	_guarddutyUnit                          string
	_guarddutyUsageCriteria                 string
	_guarddutyUsageStatisticType            string
)

// Accepts the invitation to be a member account and get monitored by a GuardDuty
// administrator account that sent the invitation.
func guardduty_AcceptAdministratorInvitation(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.AcceptAdministratorInvitationInput{
		// AdministratorId: *string, // Required
		// DetectorId: *string, // Required
		// InvitationId: *string, // Required
	}

	if len(_guarddutyAdministratorId) > 0 {
		input.AdministratorId = aws.String(_guarddutyAdministratorId)
	}
	if len(_guarddutyDetectorId) > 0 {
		input.DetectorId = aws.String(_guarddutyDetectorId)
	}
	if len(_guarddutyInvitationId) > 0 {
		input.InvitationId = aws.String(_guarddutyInvitationId)
	}

	if resp, err := client.AcceptAdministratorInvitation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Accepts the invitation to be monitored by a GuardDuty administrator account.
// Deprecated: This operation is deprecated, use AcceptAdministratorInvitation
// instead
func guardduty_AcceptInvitation(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.AcceptInvitationInput{
		// DetectorId: *string, // Required
		// InvitationId: *string, // Required
		// MasterId: *string, // Required
	}

	if len(_guarddutyDetectorId) > 0 {
		input.DetectorId = aws.String(_guarddutyDetectorId)
	}
	if len(_guarddutyInvitationId) > 0 {
		input.InvitationId = aws.String(_guarddutyInvitationId)
	}
	if len(_guarddutyMasterId) > 0 {
		input.MasterId = aws.String(_guarddutyMasterId)
	}

	if resp, err := client.AcceptInvitation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Archives GuardDuty findings that are specified by the list of finding IDs.
// Only the administrator account can archive findings. Member accounts don't have
// permission to archive findings from their accounts.
func guardduty_ArchiveFindings(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.ArchiveFindingsInput{
		// DetectorId: *string, // Required
		// FindingIds: []string, // Required
	}

	if len(_guarddutyDetectorId) > 0 {
		input.DetectorId = aws.String(_guarddutyDetectorId)
	}
	if len(_guarddutyFindingIds) > 0 {
		input.FindingIds = append([]string(nil), _guarddutyFindingIds...)
	}

	if resp, err := client.ArchiveFindings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a single GuardDuty detector. A detector is a resource that represents
// the GuardDuty service. To start using GuardDuty, you must create a detector in
// each Region where you enable the service. You can have only one detector per
// account per Region. All data sources are enabled in a new detector by default.
//
// - When you don't specify any features , with an exception to
// RUNTIME_MONITORING , all the optional features are enabled by default.
//
// - When you specify some of the features , any feature that is not specified in
// the API call gets enabled by default, with an exception to RUNTIME_MONITORING
// .
//
// Specifying both EKS Runtime Monitoring ( EKS_RUNTIME_MONITORING ) and Runtime
// Monitoring ( RUNTIME_MONITORING ) will cause an error. You can add only one of
// these two features because Runtime Monitoring already includes the threat
// detection for Amazon EKS resources. For more information, see [Runtime Monitoring].
//
// There might be regional differences because some data sources might not be
// available in all the Amazon Web Services Regions where GuardDuty is presently
// supported. For more information, see [Regions and endpoints].
//
// [Regions and endpoints]: https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_regions.html
// [Runtime Monitoring]: https://docs.aws.amazon.com/guardduty/latest/ug/runtime-monitoring.html
func guardduty_CreateDetector(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.CreateDetectorInput{
		// Enable: *bool, // Required
	}

	if len(_guarddutyEnable) > 0 {
		if err := assignInputField(input, "Enable", _guarddutyEnable); err != nil {
			log.Errorf("invalid --enable: %s", err.Error())
			return
		}
	}
	if len(_guarddutyClientToken) > 0 {
		input.ClientToken = aws.String(_guarddutyClientToken)
	}
	if len(_guarddutyDataSources) > 0 {
		if err := assignInputField(input, "DataSources", _guarddutyDataSources); err != nil {
			log.Errorf("invalid --data-sources: %s", err.Error())
			return
		}
	}
	if len(_guarddutyFeatures) > 0 {
		if err := assignInputField(input, "Features", _guarddutyFeatures); err != nil {
			log.Errorf("invalid --features: %s", err.Error())
			return
		}
	}
	if len(_guarddutyFindingPublishingFrequency) > 0 {
		if err := assignInputField(input, "FindingPublishingFrequency", _guarddutyFindingPublishingFrequency); err != nil {
			log.Errorf("invalid --finding-publishing-frequency: %s", err.Error())
			return
		}
	}
	if len(_guarddutyTags) > 0 {
		if err := assignInputField(input, "Tags", _guarddutyTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDetector(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a filter using the specified finding criteria. The maximum number of
// saved filters per Amazon Web Services account per Region is 100. For more
// information, see [Quotas for GuardDuty].
//
// [Quotas for GuardDuty]: https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_limits.html
func guardduty_CreateFilter(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.CreateFilterInput{
		// DetectorId: *string, // Required
		// FindingCriteria: *types.FindingCriteria, // Required
		// Name: *string, // Required
	}

	if len(_guarddutyDetectorId) > 0 {
		input.DetectorId = aws.String(_guarddutyDetectorId)
	}
	if len(_guarddutyFindingCriteria) > 0 {
		if err := assignInputField(input, "FindingCriteria", _guarddutyFindingCriteria); err != nil {
			log.Errorf("invalid --finding-criteria: %s", err.Error())
			return
		}
	}
	if len(_guarddutyName) > 0 {
		input.Name = aws.String(_guarddutyName)
	}
	if len(_guarddutyAction) > 0 {
		if err := assignInputField(input, "Action", _guarddutyAction); err != nil {
			log.Errorf("invalid --action: %s", err.Error())
			return
		}
	}
	if len(_guarddutyClientToken) > 0 {
		input.ClientToken = aws.String(_guarddutyClientToken)
	}
	if len(_guarddutyDescription) > 0 {
		input.Description = aws.String(_guarddutyDescription)
	}
	if len(_guarddutyRank) > 0 {
		if err := assignInputField(input, "Rank", _guarddutyRank); err != nil {
			log.Errorf("invalid --rank: %s", err.Error())
			return
		}
	}
	if len(_guarddutyTags) > 0 {
		if err := assignInputField(input, "Tags", _guarddutyTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateFilter(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new IPSet, which is called a trusted IP list in the console user
// interface. An IPSet is a list of IP addresses that are trusted for secure
// communication with Amazon Web Services infrastructure and applications.
// GuardDuty doesn't generate findings for IP addresses that are included in
// IPSets. Only users from the administrator account can use this operation.
func guardduty_CreateIPSet(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.CreateIPSetInput{
		// Activate: *bool, // Required
		// DetectorId: *string, // Required
		// Format: types.IpSetFormat, // Required
		// Location: *string, // Required
		// Name: *string, // Required
	}

	if len(_guarddutyActivate) > 0 {
		if err := assignInputField(input, "Activate", _guarddutyActivate); err != nil {
			log.Errorf("invalid --activate: %s", err.Error())
			return
		}
	}
	if len(_guarddutyDetectorId) > 0 {
		input.DetectorId = aws.String(_guarddutyDetectorId)
	}
	if len(_guarddutyFormat) > 0 {
		if err := assignInputField(input, "Format", _guarddutyFormat); err != nil {
			log.Errorf("invalid --format: %s", err.Error())
			return
		}
	}
	if len(_guarddutyLocation) > 0 {
		input.Location = aws.String(_guarddutyLocation)
	}
	if len(_guarddutyName) > 0 {
		input.Name = aws.String(_guarddutyName)
	}
	if len(_guarddutyClientToken) > 0 {
		input.ClientToken = aws.String(_guarddutyClientToken)
	}
	if len(_guarddutyExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_guarddutyExpectedBucketOwner)
	}
	if len(_guarddutyTags) > 0 {
		if err := assignInputField(input, "Tags", _guarddutyTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateIPSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new Malware Protection plan for the protected resource.
// When you create a Malware Protection plan, the Amazon Web Services service
// terms for GuardDuty Malware Protection apply. For more information, see [Amazon Web Services service terms for GuardDuty Malware Protection].
//
// [Amazon Web Services service terms for GuardDuty Malware Protection]: http://aws.amazon.com/service-terms/#87._Amazon_GuardDuty
func guardduty_CreateMalwareProtectionPlan(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.CreateMalwareProtectionPlanInput{
		// ProtectedResource: *types.CreateProtectedResource, // Required
		// Role: *string, // Required
	}

	if len(_guarddutyProtectedResource) > 0 {
		if err := assignInputField(input, "ProtectedResource", _guarddutyProtectedResource); err != nil {
			log.Errorf("invalid --protected-resource: %s", err.Error())
			return
		}
	}
	if len(_guarddutyRole) > 0 {
		input.Role = aws.String(_guarddutyRole)
	}
	if len(_guarddutyActions) > 0 {
		if err := assignInputField(input, "Actions", _guarddutyActions); err != nil {
			log.Errorf("invalid --actions: %s", err.Error())
			return
		}
	}
	if len(_guarddutyClientToken) > 0 {
		input.ClientToken = aws.String(_guarddutyClientToken)
	}
	if len(_guarddutyTags) > 0 {
		if err := assignInputField(input, "Tags", _guarddutyTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateMalwareProtectionPlan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates member accounts of the current Amazon Web Services account by
// specifying a list of Amazon Web Services account IDs. This step is a
// prerequisite for managing the associated member accounts either by invitation or
// through an organization.
//
// As a delegated administrator, using CreateMembers will enable GuardDuty in the
// added member accounts, with the exception of the organization delegated
// administrator account. A delegated administrator must enable GuardDuty prior to
// being added as a member.
//
// When you use CreateMembers as an Organizations delegated administrator,
// GuardDuty applies your organization's auto-enable settings to the member
// accounts in this request, irrespective of the accounts being new or existing
// members. For more information about the existing auto-enable settings for your
// organization, see [DescribeOrganizationConfiguration].
//
// If you disassociate a member account that was added by invitation, the member
// account details obtained from this API, including the associated email
// addresses, will be retained. This is done so that the delegated administrator
// can invoke the [InviteMembers]API without the need to invoke the CreateMembers API again. To
// remove the details associated with a member account, the delegated administrator
// must invoke the [DeleteMembers]API.
//
// When the member accounts added through Organizations are later disassociated,
// you (administrator) can't invite them by calling the InviteMembers API. You can
// create an association with these member accounts again only by calling the
// CreateMembers API.
//
// [DeleteMembers]: https://docs.aws.amazon.com/guardduty/latest/APIReference/API_DeleteMembers.html
// [DescribeOrganizationConfiguration]: https://docs.aws.amazon.com/guardduty/latest/APIReference/API_DescribeOrganizationConfiguration.html
// [InviteMembers]: https://docs.aws.amazon.com/guardduty/latest/APIReference/API_InviteMembers.html
func guardduty_CreateMembers(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.CreateMembersInput{
		// AccountDetails: []types.AccountDetail, // Required
		// DetectorId: *string, // Required
	}

	if len(_guarddutyAccountDetails) > 0 {
		if err := assignInputField(input, "AccountDetails", _guarddutyAccountDetails); err != nil {
			log.Errorf("invalid --account-details: %s", err.Error())
			return
		}
	}
	if len(_guarddutyDetectorId) > 0 {
		input.DetectorId = aws.String(_guarddutyDetectorId)
	}

	if resp, err := client.CreateMembers(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a publishing destination where you can export your GuardDuty findings.
// Before you start exporting the findings, the destination resource must exist.
func guardduty_CreatePublishingDestination(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.CreatePublishingDestinationInput{
		// DestinationProperties: *types.DestinationProperties, // Required
		// DestinationType: types.DestinationType, // Required
		// DetectorId: *string, // Required
	}

	if len(_guarddutyDestinationProperties) > 0 {
		if err := assignInputField(input, "DestinationProperties", _guarddutyDestinationProperties); err != nil {
			log.Errorf("invalid --destination-properties: %s", err.Error())
			return
		}
	}
	if len(_guarddutyDestinationType) > 0 {
		if err := assignInputField(input, "DestinationType", _guarddutyDestinationType); err != nil {
			log.Errorf("invalid --destination-type: %s", err.Error())
			return
		}
	}
	if len(_guarddutyDetectorId) > 0 {
		input.DetectorId = aws.String(_guarddutyDetectorId)
	}
	if len(_guarddutyClientToken) > 0 {
		input.ClientToken = aws.String(_guarddutyClientToken)
	}
	if len(_guarddutyTags) > 0 {
		if err := assignInputField(input, "Tags", _guarddutyTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreatePublishingDestination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Generates sample findings of types specified by the list of finding types. If
// 'NULL' is specified for findingTypes , the API generates sample findings of all
// supported finding types.
func guardduty_CreateSampleFindings(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.CreateSampleFindingsInput{
		// DetectorId: *string, // Required
	}

	if len(_guarddutyDetectorId) > 0 {
		input.DetectorId = aws.String(_guarddutyDetectorId)
	}
	if len(_guarddutyFindingTypes) > 0 {
		input.FindingTypes = append([]string(nil), _guarddutyFindingTypes...)
	}

	if resp, err := client.CreateSampleFindings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new threat entity set. In a threat entity set, you can provide known
// malicious IP addresses and domains for your Amazon Web Services environment.
// GuardDuty generates findings based on the entries in the threat entity sets.
// Only users of the administrator account can manage entity sets, which
// automatically apply to member accounts.
func guardduty_CreateThreatEntitySet(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.CreateThreatEntitySetInput{
		// Activate: *bool, // Required
		// DetectorId: *string, // Required
		// Format: types.ThreatEntitySetFormat, // Required
		// Location: *string, // Required
		// Name: *string, // Required
	}

	if len(_guarddutyActivate) > 0 {
		if err := assignInputField(input, "Activate", _guarddutyActivate); err != nil {
			log.Errorf("invalid --activate: %s", err.Error())
			return
		}
	}
	if len(_guarddutyDetectorId) > 0 {
		input.DetectorId = aws.String(_guarddutyDetectorId)
	}
	if len(_guarddutyFormat) > 0 {
		if err := assignInputField(input, "Format", _guarddutyFormat); err != nil {
			log.Errorf("invalid --format: %s", err.Error())
			return
		}
	}
	if len(_guarddutyLocation) > 0 {
		input.Location = aws.String(_guarddutyLocation)
	}
	if len(_guarddutyName) > 0 {
		input.Name = aws.String(_guarddutyName)
	}
	if len(_guarddutyClientToken) > 0 {
		input.ClientToken = aws.String(_guarddutyClientToken)
	}
	if len(_guarddutyExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_guarddutyExpectedBucketOwner)
	}
	if len(_guarddutyTags) > 0 {
		if err := assignInputField(input, "Tags", _guarddutyTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateThreatEntitySet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new ThreatIntelSet. ThreatIntelSets consist of known malicious IP
// addresses. GuardDuty generates findings based on ThreatIntelSets. Only users of
// the administrator account can use this operation.
func guardduty_CreateThreatIntelSet(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.CreateThreatIntelSetInput{
		// Activate: *bool, // Required
		// DetectorId: *string, // Required
		// Format: types.ThreatIntelSetFormat, // Required
		// Location: *string, // Required
		// Name: *string, // Required
	}

	if len(_guarddutyActivate) > 0 {
		if err := assignInputField(input, "Activate", _guarddutyActivate); err != nil {
			log.Errorf("invalid --activate: %s", err.Error())
			return
		}
	}
	if len(_guarddutyDetectorId) > 0 {
		input.DetectorId = aws.String(_guarddutyDetectorId)
	}
	if len(_guarddutyFormat) > 0 {
		if err := assignInputField(input, "Format", _guarddutyFormat); err != nil {
			log.Errorf("invalid --format: %s", err.Error())
			return
		}
	}
	if len(_guarddutyLocation) > 0 {
		input.Location = aws.String(_guarddutyLocation)
	}
	if len(_guarddutyName) > 0 {
		input.Name = aws.String(_guarddutyName)
	}
	if len(_guarddutyClientToken) > 0 {
		input.ClientToken = aws.String(_guarddutyClientToken)
	}
	if len(_guarddutyExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_guarddutyExpectedBucketOwner)
	}
	if len(_guarddutyTags) > 0 {
		if err := assignInputField(input, "Tags", _guarddutyTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateThreatIntelSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new trusted entity set. In the trusted entity set, you can provide IP
// addresses and domains that you believe are secure for communication in your
// Amazon Web Services environment. GuardDuty will not generate findings for the
// entries that are specified in a trusted entity set. At any given time, you can
// have only one trusted entity set.
//
// Only users of the administrator account can manage the entity sets, which
// automatically apply to member accounts.
func guardduty_CreateTrustedEntitySet(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.CreateTrustedEntitySetInput{
		// Activate: *bool, // Required
		// DetectorId: *string, // Required
		// Format: types.TrustedEntitySetFormat, // Required
		// Location: *string, // Required
		// Name: *string, // Required
	}

	if len(_guarddutyActivate) > 0 {
		if err := assignInputField(input, "Activate", _guarddutyActivate); err != nil {
			log.Errorf("invalid --activate: %s", err.Error())
			return
		}
	}
	if len(_guarddutyDetectorId) > 0 {
		input.DetectorId = aws.String(_guarddutyDetectorId)
	}
	if len(_guarddutyFormat) > 0 {
		if err := assignInputField(input, "Format", _guarddutyFormat); err != nil {
			log.Errorf("invalid --format: %s", err.Error())
			return
		}
	}
	if len(_guarddutyLocation) > 0 {
		input.Location = aws.String(_guarddutyLocation)
	}
	if len(_guarddutyName) > 0 {
		input.Name = aws.String(_guarddutyName)
	}
	if len(_guarddutyClientToken) > 0 {
		input.ClientToken = aws.String(_guarddutyClientToken)
	}
	if len(_guarddutyExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_guarddutyExpectedBucketOwner)
	}
	if len(_guarddutyTags) > 0 {
		if err := assignInputField(input, "Tags", _guarddutyTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateTrustedEntitySet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Declines invitations sent to the current member account by Amazon Web Services
// accounts specified by their account IDs.
func guardduty_DeclineInvitations(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.DeclineInvitationsInput{
		// AccountIds: []string, // Required
	}

	if len(_guarddutyAccountIds) > 0 {
		input.AccountIds = append([]string(nil), _guarddutyAccountIds...)
	}

	if resp, err := client.DeclineInvitations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Amazon GuardDuty detector that is specified by the detector ID.
func guardduty_DeleteDetector(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.DeleteDetectorInput{
		// DetectorId: *string, // Required
	}

	if len(_guarddutyDetectorId) > 0 {
		input.DetectorId = aws.String(_guarddutyDetectorId)
	}

	if resp, err := client.DeleteDetector(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the filter specified by the filter name.
func guardduty_DeleteFilter(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.DeleteFilterInput{
		// DetectorId: *string, // Required
		// FilterName: *string, // Required
	}

	if len(_guarddutyDetectorId) > 0 {
		input.DetectorId = aws.String(_guarddutyDetectorId)
	}
	if len(_guarddutyFilterName) > 0 {
		input.FilterName = aws.String(_guarddutyFilterName)
	}

	if resp, err := client.DeleteFilter(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes invitations sent to the current member account by Amazon Web Services
// accounts specified by their account IDs.
func guardduty_DeleteInvitations(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.DeleteInvitationsInput{
		// AccountIds: []string, // Required
	}

	if len(_guarddutyAccountIds) > 0 {
		input.AccountIds = append([]string(nil), _guarddutyAccountIds...)
	}

	if resp, err := client.DeleteInvitations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the IPSet specified by the ipSetId . IPSets are called trusted IP lists
// in the console user interface.
func guardduty_DeleteIPSet(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.DeleteIPSetInput{
		// DetectorId: *string, // Required
		// IpSetId: *string, // Required
	}

	if len(_guarddutyDetectorId) > 0 {
		input.DetectorId = aws.String(_guarddutyDetectorId)
	}
	if len(_guarddutyIpSetId) > 0 {
		input.IpSetId = aws.String(_guarddutyIpSetId)
	}

	if resp, err := client.DeleteIPSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the Malware Protection plan ID associated with the Malware Protection
// plan resource. Use this API only when you no longer want to protect the resource
// associated with this Malware Protection plan ID.
func guardduty_DeleteMalwareProtectionPlan(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.DeleteMalwareProtectionPlanInput{
		// MalwareProtectionPlanId: *string, // Required
	}

	if len(_guarddutyMalwareProtectionPlanId) > 0 {
		input.MalwareProtectionPlanId = aws.String(_guarddutyMalwareProtectionPlanId)
	}

	if resp, err := client.DeleteMalwareProtectionPlan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes GuardDuty member accounts (to the current GuardDuty administrator
// account) specified by the account IDs.
//
// With autoEnableOrganizationMembers configuration for your organization set to
// ALL , you'll receive an error if you attempt to disable GuardDuty for a member
// account in your organization.
func guardduty_DeleteMembers(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.DeleteMembersInput{
		// AccountIds: []string, // Required
		// DetectorId: *string, // Required
	}

	if len(_guarddutyAccountIds) > 0 {
		input.AccountIds = append([]string(nil), _guarddutyAccountIds...)
	}
	if len(_guarddutyDetectorId) > 0 {
		input.DetectorId = aws.String(_guarddutyDetectorId)
	}

	if resp, err := client.DeleteMembers(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the publishing definition with the specified destinationId .
func guardduty_DeletePublishingDestination(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.DeletePublishingDestinationInput{
		// DestinationId: *string, // Required
		// DetectorId: *string, // Required
	}

	if len(_guarddutyDestinationId) > 0 {
		input.DestinationId = aws.String(_guarddutyDestinationId)
	}
	if len(_guarddutyDetectorId) > 0 {
		input.DetectorId = aws.String(_guarddutyDetectorId)
	}

	if resp, err := client.DeletePublishingDestination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the threat entity set that is associated with the specified
// threatEntitySetId .
func guardduty_DeleteThreatEntitySet(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.DeleteThreatEntitySetInput{
		// DetectorId: *string, // Required
		// ThreatEntitySetId: *string, // Required
	}

	if len(_guarddutyDetectorId) > 0 {
		input.DetectorId = aws.String(_guarddutyDetectorId)
	}
	if len(_guarddutyThreatEntitySetId) > 0 {
		input.ThreatEntitySetId = aws.String(_guarddutyThreatEntitySetId)
	}

	if resp, err := client.DeleteThreatEntitySet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the ThreatIntelSet specified by the ThreatIntelSet ID.
func guardduty_DeleteThreatIntelSet(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.DeleteThreatIntelSetInput{
		// DetectorId: *string, // Required
		// ThreatIntelSetId: *string, // Required
	}

	if len(_guarddutyDetectorId) > 0 {
		input.DetectorId = aws.String(_guarddutyDetectorId)
	}
	if len(_guarddutyThreatIntelSetId) > 0 {
		input.ThreatIntelSetId = aws.String(_guarddutyThreatIntelSetId)
	}

	if resp, err := client.DeleteThreatIntelSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the trusted entity set that is associated with the specified
// trustedEntitySetId .
func guardduty_DeleteTrustedEntitySet(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.DeleteTrustedEntitySetInput{
		// DetectorId: *string, // Required
		// TrustedEntitySetId: *string, // Required
	}

	if len(_guarddutyDetectorId) > 0 {
		input.DetectorId = aws.String(_guarddutyDetectorId)
	}
	if len(_guarddutyTrustedEntitySetId) > 0 {
		input.TrustedEntitySetId = aws.String(_guarddutyTrustedEntitySetId)
	}

	if resp, err := client.DeleteTrustedEntitySet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of malware scans. Each member account can view the malware scans
// for their own accounts. An administrator can view the malware scans for all the
// member accounts.
//
// There might be regional differences because some data sources might not be
// available in all the Amazon Web Services Regions where GuardDuty is presently
// supported. For more information, see [Regions and endpoints].
//
// [Regions and endpoints]: https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_regions.html
func guardduty_DescribeMalwareScans(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.DescribeMalwareScansInput{
		// DetectorId: *string, // Required
	}

	if len(_guarddutyDetectorId) > 0 {
		input.DetectorId = aws.String(_guarddutyDetectorId)
	}
	if len(_guarddutyFilterCriteria) > 0 {
		if err := assignInputField(input, "FilterCriteria", _guarddutyFilterCriteria); err != nil {
			log.Errorf("invalid --filter-criteria: %s", err.Error())
			return
		}
	}
	if len(_guarddutyMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _guarddutyMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_guarddutyNextToken) > 0 {
		input.NextToken = aws.String(_guarddutyNextToken)
	}
	if len(_guarddutySortCriteria) > 0 {
		if err := assignInputField(input, "SortCriteria", _guarddutySortCriteria); err != nil {
			log.Errorf("invalid --sort-criteria: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeMalwareScans(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*guardduty.DescribeMalwareScansOutput
	p := guardduty.NewDescribeMalwareScansPaginator(client, input)
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

// Returns information about the account selected as the delegated administrator
// for GuardDuty.
//
// There might be regional differences because some data sources might not be
// available in all the Amazon Web Services Regions where GuardDuty is presently
// supported. For more information, see [Regions and endpoints].
//
// [Regions and endpoints]: https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_regions.html
func guardduty_DescribeOrganizationConfiguration(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.DescribeOrganizationConfigurationInput{
		// DetectorId: *string, // Required
	}

	if len(_guarddutyDetectorId) > 0 {
		input.DetectorId = aws.String(_guarddutyDetectorId)
	}
	if len(_guarddutyMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _guarddutyMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_guarddutyNextToken) > 0 {
		input.NextToken = aws.String(_guarddutyNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeOrganizationConfiguration(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*guardduty.DescribeOrganizationConfigurationOutput
	p := guardduty.NewDescribeOrganizationConfigurationPaginator(client, input)
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

// Returns information about the publishing destination specified by the provided
// destinationId .
func guardduty_DescribePublishingDestination(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.DescribePublishingDestinationInput{
		// DestinationId: *string, // Required
		// DetectorId: *string, // Required
	}

	if len(_guarddutyDestinationId) > 0 {
		input.DestinationId = aws.String(_guarddutyDestinationId)
	}
	if len(_guarddutyDetectorId) > 0 {
		input.DetectorId = aws.String(_guarddutyDetectorId)
	}

	if resp, err := client.DescribePublishingDestination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the existing GuardDuty delegated administrator of the organization.
// Only the organization's management account can run this API operation.
func guardduty_DisableOrganizationAdminAccount(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.DisableOrganizationAdminAccountInput{
		// AdminAccountId: *string, // Required
	}

	if len(_guarddutyAdminAccountId) > 0 {
		input.AdminAccountId = aws.String(_guarddutyAdminAccountId)
	}

	if resp, err := client.DisableOrganizationAdminAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates the current GuardDuty member account from its administrator
// account.
//
// When you disassociate an invited member from a GuardDuty delegated
// administrator, the member account details obtained from the [CreateMembers]API, including the
// associated email addresses, are retained. This is done so that the delegated
// administrator can invoke the [InviteMembers]API without the need to invoke the CreateMembers
// API again. To remove the details associated with a member account, the delegated
// administrator must invoke the [DeleteMembers]API.
//
// With autoEnableOrganizationMembers configuration for your organization set to
// ALL , you'll receive an error if you attempt to disable GuardDuty in a member
// account.
//
// [DeleteMembers]: https://docs.aws.amazon.com/guardduty/latest/APIReference/API_DeleteMembers.html
// [CreateMembers]: https://docs.aws.amazon.com/guardduty/latest/APIReference/API_CreateMembers.html
// [InviteMembers]: https://docs.aws.amazon.com/guardduty/latest/APIReference/API_InviteMembers.html
func guardduty_DisassociateFromAdministratorAccount(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.DisassociateFromAdministratorAccountInput{
		// DetectorId: *string, // Required
	}

	if len(_guarddutyDetectorId) > 0 {
		input.DetectorId = aws.String(_guarddutyDetectorId)
	}

	if resp, err := client.DisassociateFromAdministratorAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates the current GuardDuty member account from its administrator
// account.
//
// When you disassociate an invited member from a GuardDuty delegated
// administrator, the member account details obtained from the [CreateMembers]API, including the
// associated email addresses, are retained. This is done so that the delegated
// administrator can invoke the [InviteMembers]API without the need to invoke the CreateMembers
// API again. To remove the details associated with a member account, the delegated
// administrator must invoke the [DeleteMembers]API.
//
// Deprecated: This operation is deprecated, use
// DisassociateFromAdministratorAccount instead
//
// [DeleteMembers]: https://docs.aws.amazon.com/guardduty/latest/APIReference/API_DeleteMembers.html
// [CreateMembers]: https://docs.aws.amazon.com/guardduty/latest/APIReference/API_CreateMembers.html
// [InviteMembers]: https://docs.aws.amazon.com/guardduty/latest/APIReference/API_InviteMembers.html
func guardduty_DisassociateFromMasterAccount(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.DisassociateFromMasterAccountInput{
		// DetectorId: *string, // Required
	}

	if len(_guarddutyDetectorId) > 0 {
		input.DetectorId = aws.String(_guarddutyDetectorId)
	}

	if resp, err := client.DisassociateFromMasterAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates GuardDuty member accounts (from the current administrator
// account) specified by the account IDs.
//
// When you disassociate an invited member from a GuardDuty delegated
// administrator, the member account details obtained from the [CreateMembers]API, including the
// associated email addresses, are retained. This is done so that the delegated
// administrator can invoke the [InviteMembers]API without the need to invoke the CreateMembers
// API again. To remove the details associated with a member account, the delegated
// administrator must invoke the [DeleteMembers]API.
//
// With autoEnableOrganizationMembers configuration for your organization set to
// ALL , you'll receive an error if you attempt to disassociate a member account
// before removing them from your organization.
//
// If you disassociate a member account that was added by invitation, the member
// account details obtained from this API, including the associated email
// addresses, will be retained. This is done so that the delegated administrator
// can invoke the [InviteMembers]API without the need to invoke the CreateMembers API again. To
// remove the details associated with a member account, the delegated administrator
// must invoke the [DeleteMembers]API.
//
// When the member accounts added through Organizations are later disassociated,
// you (administrator) can't invite them by calling the InviteMembers API. You can
// create an association with these member accounts again only by calling the
// CreateMembers API.
//
// [DeleteMembers]: https://docs.aws.amazon.com/guardduty/latest/APIReference/API_DeleteMembers.html
// [CreateMembers]: https://docs.aws.amazon.com/guardduty/latest/APIReference/API_CreateMembers.html
// [InviteMembers]: https://docs.aws.amazon.com/guardduty/latest/APIReference/API_InviteMembers.html
func guardduty_DisassociateMembers(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.DisassociateMembersInput{
		// AccountIds: []string, // Required
		// DetectorId: *string, // Required
	}

	if len(_guarddutyAccountIds) > 0 {
		input.AccountIds = append([]string(nil), _guarddutyAccountIds...)
	}
	if len(_guarddutyDetectorId) > 0 {
		input.DetectorId = aws.String(_guarddutyDetectorId)
	}

	if resp, err := client.DisassociateMembers(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Designates an Amazon Web Services account within the organization as your
// GuardDuty delegated administrator. Only the organization's management account
// can run this API operation.
func guardduty_EnableOrganizationAdminAccount(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.EnableOrganizationAdminAccountInput{
		// AdminAccountId: *string, // Required
	}

	if len(_guarddutyAdminAccountId) > 0 {
		input.AdminAccountId = aws.String(_guarddutyAdminAccountId)
	}

	if resp, err := client.EnableOrganizationAdminAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides the details of the GuardDuty administrator account associated with the
// current GuardDuty member account.
//
// Based on the type of account that runs this API, the following list shows how
// the API behavior varies:
//
// - When the GuardDuty administrator account runs this API, it will return
// success ( HTTP 200 ) but no content.
//
// - When a member account runs this API, it will return the details of the
// GuardDuty administrator account that is associated with this calling member
// account.
//
// - When an individual account (not associated with an organization) runs this
// API, it will return success ( HTTP 200 ) but no content.
func guardduty_GetAdministratorAccount(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.GetAdministratorAccountInput{
		// DetectorId: *string, // Required
	}

	if len(_guarddutyDetectorId) > 0 {
		input.DetectorId = aws.String(_guarddutyDetectorId)
	}

	if resp, err := client.GetAdministratorAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves aggregated statistics for your account. If you are a GuardDuty
// administrator, you can retrieve the statistics for all the resources associated
// with the active member accounts in your organization who have enabled Runtime
// Monitoring and have the GuardDuty security agent running on their resources.
func guardduty_GetCoverageStatistics(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.GetCoverageStatisticsInput{
		// DetectorId: *string, // Required
		// StatisticsType: []types.CoverageStatisticsType, // Required
	}

	if len(_guarddutyDetectorId) > 0 {
		input.DetectorId = aws.String(_guarddutyDetectorId)
	}
	if len(_guarddutyStatisticsType) > 0 {
		if err := assignInputField(input, "StatisticsType", _guarddutyStatisticsType); err != nil {
			log.Errorf("invalid --statistics-type: %s", err.Error())
			return
		}
	}
	if len(_guarddutyFilterCriteria) > 0 {
		if err := assignInputField(input, "FilterCriteria", _guarddutyFilterCriteria); err != nil {
			log.Errorf("invalid --filter-criteria: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetCoverageStatistics(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a GuardDuty detector specified by the detectorId.
// There might be regional differences because some data sources might not be
// available in all the Amazon Web Services Regions where GuardDuty is presently
// supported. For more information, see [Regions and endpoints].
//
// [Regions and endpoints]: https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_regions.html
func guardduty_GetDetector(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.GetDetectorInput{
		// DetectorId: *string, // Required
	}

	if len(_guarddutyDetectorId) > 0 {
		input.DetectorId = aws.String(_guarddutyDetectorId)
	}

	if resp, err := client.GetDetector(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the details of the filter specified by the filter name.
func guardduty_GetFilter(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.GetFilterInput{
		// DetectorId: *string, // Required
		// FilterName: *string, // Required
	}

	if len(_guarddutyDetectorId) > 0 {
		input.DetectorId = aws.String(_guarddutyDetectorId)
	}
	if len(_guarddutyFilterName) > 0 {
		input.FilterName = aws.String(_guarddutyFilterName)
	}

	if resp, err := client.GetFilter(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes Amazon GuardDuty findings specified by finding IDs.
func guardduty_GetFindings(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.GetFindingsInput{
		// DetectorId: *string, // Required
		// FindingIds: []string, // Required
	}

	if len(_guarddutyDetectorId) > 0 {
		input.DetectorId = aws.String(_guarddutyDetectorId)
	}
	if len(_guarddutyFindingIds) > 0 {
		input.FindingIds = append([]string(nil), _guarddutyFindingIds...)
	}
	if len(_guarddutySortCriteria) > 0 {
		if err := assignInputField(input, "SortCriteria", _guarddutySortCriteria); err != nil {
			log.Errorf("invalid --sort-criteria: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetFindings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists GuardDuty findings statistics for the specified detector ID.
// You must provide either findingStatisticTypes or groupBy parameter, and not
// both. You can use the maxResults and orderBy parameters only when using groupBy .
//
// There might be regional differences because some flags might not be available
// in all the Regions where GuardDuty is currently supported. For more information,
// see [Regions and endpoints].
//
// [Regions and endpoints]: https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_regions.html
func guardduty_GetFindingsStatistics(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.GetFindingsStatisticsInput{
		// DetectorId: *string, // Required
	}

	if len(_guarddutyDetectorId) > 0 {
		input.DetectorId = aws.String(_guarddutyDetectorId)
	}
	if len(_guarddutyFindingCriteria) > 0 {
		if err := assignInputField(input, "FindingCriteria", _guarddutyFindingCriteria); err != nil {
			log.Errorf("invalid --finding-criteria: %s", err.Error())
			return
		}
	}
	if len(_guarddutyFindingStatisticTypes) > 0 {
		if err := assignInputField(input, "FindingStatisticTypes", _guarddutyFindingStatisticTypes); err != nil {
			log.Errorf("invalid --finding-statistic-types: %s", err.Error())
			return
		}
	}
	if len(_guarddutyGroupBy) > 0 {
		if err := assignInputField(input, "GroupBy", _guarddutyGroupBy); err != nil {
			log.Errorf("invalid --group-by: %s", err.Error())
			return
		}
	}
	if len(_guarddutyMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _guarddutyMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_guarddutyOrderBy) > 0 {
		if err := assignInputField(input, "OrderBy", _guarddutyOrderBy); err != nil {
			log.Errorf("invalid --order-by: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetFindingsStatistics(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the count of all GuardDuty membership invitations that were sent to the
// current member account except the currently accepted invitation.
func guardduty_GetInvitationsCount(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.GetInvitationsCountInput{}

	if resp, err := client.GetInvitationsCount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the IPSet specified by the ipSetId .
func guardduty_GetIPSet(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.GetIPSetInput{
		// DetectorId: *string, // Required
		// IpSetId: *string, // Required
	}

	if len(_guarddutyDetectorId) > 0 {
		input.DetectorId = aws.String(_guarddutyDetectorId)
	}
	if len(_guarddutyIpSetId) > 0 {
		input.IpSetId = aws.String(_guarddutyIpSetId)
	}

	if resp, err := client.GetIPSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the Malware Protection plan details associated with a Malware
// Protection plan ID.
func guardduty_GetMalwareProtectionPlan(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.GetMalwareProtectionPlanInput{
		// MalwareProtectionPlanId: *string, // Required
	}

	if len(_guarddutyMalwareProtectionPlanId) > 0 {
		input.MalwareProtectionPlanId = aws.String(_guarddutyMalwareProtectionPlanId)
	}

	if resp, err := client.GetMalwareProtectionPlan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the detailed information for a specific malware scan. Each member
// account can view the malware scan details for their own account. An
// administrator can view malware scan details for all accounts in the
// organization.
//
// There might be regional differences because some data sources might not be
// available in all the Amazon Web Services Regions where GuardDuty is presently
// supported. For more information, see [Regions and endpoints].
//
// [Regions and endpoints]: https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_regions.html
func guardduty_GetMalwareScan(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.GetMalwareScanInput{
		// ScanId: *string, // Required
	}

	if len(_guarddutyScanId) > 0 {
		input.ScanId = aws.String(_guarddutyScanId)
	}

	if resp, err := client.GetMalwareScan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the details of the malware scan settings.
// There might be regional differences because some data sources might not be
// available in all the Amazon Web Services Regions where GuardDuty is presently
// supported. For more information, see [Regions and endpoints].
//
// [Regions and endpoints]: https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_regions.html
func guardduty_GetMalwareScanSettings(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.GetMalwareScanSettingsInput{
		// DetectorId: *string, // Required
	}

	if len(_guarddutyDetectorId) > 0 {
		input.DetectorId = aws.String(_guarddutyDetectorId)
	}

	if resp, err := client.GetMalwareScanSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides the details for the GuardDuty administrator account associated with
// the current GuardDuty member account.
//
// Deprecated: This operation is deprecated, use GetAdministratorAccount instead
func guardduty_GetMasterAccount(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.GetMasterAccountInput{
		// DetectorId: *string, // Required
	}

	if len(_guarddutyDetectorId) > 0 {
		input.DetectorId = aws.String(_guarddutyDetectorId)
	}

	if resp, err := client.GetMasterAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes which data sources are enabled for the member account's detector.
// There might be regional differences because some data sources might not be
// available in all the Amazon Web Services Regions where GuardDuty is presently
// supported. For more information, see [Regions and endpoints].
//
// [Regions and endpoints]: https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_regions.html
func guardduty_GetMemberDetectors(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.GetMemberDetectorsInput{
		// AccountIds: []string, // Required
		// DetectorId: *string, // Required
	}

	if len(_guarddutyAccountIds) > 0 {
		input.AccountIds = append([]string(nil), _guarddutyAccountIds...)
	}
	if len(_guarddutyDetectorId) > 0 {
		input.DetectorId = aws.String(_guarddutyDetectorId)
	}

	if resp, err := client.GetMemberDetectors(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves GuardDuty member accounts (of the current GuardDuty administrator
// account) specified by the account IDs.
func guardduty_GetMembers(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.GetMembersInput{
		// AccountIds: []string, // Required
		// DetectorId: *string, // Required
	}

	if len(_guarddutyAccountIds) > 0 {
		input.AccountIds = append([]string(nil), _guarddutyAccountIds...)
	}
	if len(_guarddutyDetectorId) > 0 {
		input.DetectorId = aws.String(_guarddutyDetectorId)
	}

	if resp, err := client.GetMembers(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves how many active member accounts have each feature enabled within
// GuardDuty. Only a delegated GuardDuty administrator of an organization can run
// this API.
//
// When you create a new organization, it might take up to 24 hours to generate
// the statistics for the entire organization.
func guardduty_GetOrganizationStatistics(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.GetOrganizationStatisticsInput{}

	if resp, err := client.GetOrganizationStatistics(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides the number of days left for each data source used in the free trial
// period.
func guardduty_GetRemainingFreeTrialDays(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.GetRemainingFreeTrialDaysInput{
		// AccountIds: []string, // Required
		// DetectorId: *string, // Required
	}

	if len(_guarddutyAccountIds) > 0 {
		input.AccountIds = append([]string(nil), _guarddutyAccountIds...)
	}
	if len(_guarddutyDetectorId) > 0 {
		input.DetectorId = aws.String(_guarddutyDetectorId)
	}

	if resp, err := client.GetRemainingFreeTrialDays(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the threat entity set associated with the specified threatEntitySetId .
func guardduty_GetThreatEntitySet(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.GetThreatEntitySetInput{
		// DetectorId: *string, // Required
		// ThreatEntitySetId: *string, // Required
	}

	if len(_guarddutyDetectorId) > 0 {
		input.DetectorId = aws.String(_guarddutyDetectorId)
	}
	if len(_guarddutyThreatEntitySetId) > 0 {
		input.ThreatEntitySetId = aws.String(_guarddutyThreatEntitySetId)
	}

	if resp, err := client.GetThreatEntitySet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the ThreatIntelSet that is specified by the ThreatIntelSet ID.
func guardduty_GetThreatIntelSet(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.GetThreatIntelSetInput{
		// DetectorId: *string, // Required
		// ThreatIntelSetId: *string, // Required
	}

	if len(_guarddutyDetectorId) > 0 {
		input.DetectorId = aws.String(_guarddutyDetectorId)
	}
	if len(_guarddutyThreatIntelSetId) > 0 {
		input.ThreatIntelSetId = aws.String(_guarddutyThreatIntelSetId)
	}

	if resp, err := client.GetThreatIntelSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the trusted entity set associated with the specified
// trustedEntitySetId .
func guardduty_GetTrustedEntitySet(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.GetTrustedEntitySetInput{
		// DetectorId: *string, // Required
		// TrustedEntitySetId: *string, // Required
	}

	if len(_guarddutyDetectorId) > 0 {
		input.DetectorId = aws.String(_guarddutyDetectorId)
	}
	if len(_guarddutyTrustedEntitySetId) > 0 {
		input.TrustedEntitySetId = aws.String(_guarddutyTrustedEntitySetId)
	}

	if resp, err := client.GetTrustedEntitySet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists Amazon GuardDuty usage statistics over the last 30 days for the specified
// detector ID. For newly enabled detectors or data sources, the cost returned will
// include only the usage so far under 30 days. This may differ from the cost
// metrics in the console, which project usage over 30 days to provide a monthly
// cost estimate. For more information, see [Understanding How Usage Costs are Calculated].
//
// [Understanding How Usage Costs are Calculated]: https://docs.aws.amazon.com/guardduty/latest/ug/monitoring_costs.html#usage-calculations
func guardduty_GetUsageStatistics(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.GetUsageStatisticsInput{
		// DetectorId: *string, // Required
		// UsageCriteria: *types.UsageCriteria, // Required
		// UsageStatisticType: types.UsageStatisticType, // Required
	}

	if len(_guarddutyDetectorId) > 0 {
		input.DetectorId = aws.String(_guarddutyDetectorId)
	}
	if len(_guarddutyUsageCriteria) > 0 {
		if err := assignInputField(input, "UsageCriteria", _guarddutyUsageCriteria); err != nil {
			log.Errorf("invalid --usage-criteria: %s", err.Error())
			return
		}
	}
	if len(_guarddutyUsageStatisticType) > 0 {
		if err := assignInputField(input, "UsageStatisticType", _guarddutyUsageStatisticType); err != nil {
			log.Errorf("invalid --usage-statistic-type: %s", err.Error())
			return
		}
	}
	if len(_guarddutyMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _guarddutyMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_guarddutyNextToken) > 0 {
		input.NextToken = aws.String(_guarddutyNextToken)
	}
	if len(_guarddutyUnit) > 0 {
		input.Unit = aws.String(_guarddutyUnit)
	}

	if disablePaginator() {
		if resp, err := client.GetUsageStatistics(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*guardduty.GetUsageStatisticsOutput
	p := guardduty.NewGetUsageStatisticsPaginator(client, input)
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

// Invites Amazon Web Services accounts to become members of an organization
// administered by the Amazon Web Services account that invokes this API. If you
// are using Amazon Web Services Organizations to manage your GuardDuty
// environment, this step is not needed. For more information, see [Managing accounts with organizations].
//
// To invite Amazon Web Services accounts, the first step is to ensure that
// GuardDuty has been enabled in the potential member accounts. You can now invoke
// this API to add accounts by invitation. The invited accounts can either accept
// or decline the invitation from their GuardDuty accounts. Each invited Amazon Web
// Services account can choose to accept the invitation from only one Amazon Web
// Services account. For more information, see [Managing GuardDuty accounts by invitation].
//
// After the invite has been accepted and you choose to disassociate a member
// account (by using [DisassociateMembers]) from your account, the details of the member account
// obtained by invoking [CreateMembers], including the associated email addresses, will be
// retained. This is done so that you can invoke InviteMembers without the need to
// invoke [CreateMembers]again. To remove the details associated with a member account, you must
// also invoke [DeleteMembers].
//
// If you disassociate a member account that was added by invitation, the member
// account details obtained from this API, including the associated email
// addresses, will be retained. This is done so that the delegated administrator
// can invoke the [InviteMembers]API without the need to invoke the CreateMembers API again. To
// remove the details associated with a member account, the delegated administrator
// must invoke the [DeleteMembers]API.
//
// When the member accounts added through Organizations are later disassociated,
// you (administrator) can't invite them by calling the InviteMembers API. You can
// create an association with these member accounts again only by calling the
// CreateMembers API.
//
// [Managing GuardDuty accounts by invitation]: https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_invitations.html
// [DeleteMembers]: https://docs.aws.amazon.com/guardduty/latest/APIReference/API_DeleteMembers.html
// [CreateMembers]: https://docs.aws.amazon.com/guardduty/latest/APIReference/API_CreateMembers.html
// [Managing accounts with organizations]: https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_organizations.html
// [DisassociateMembers]: https://docs.aws.amazon.com/guardduty/latest/APIReference/API_DisassociateMembers.html
// [InviteMembers]: https://docs.aws.amazon.com/guardduty/latest/APIReference/API_InviteMembers.html
func guardduty_InviteMembers(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.InviteMembersInput{
		// AccountIds: []string, // Required
		// DetectorId: *string, // Required
	}

	if len(_guarddutyAccountIds) > 0 {
		input.AccountIds = append([]string(nil), _guarddutyAccountIds...)
	}
	if len(_guarddutyDetectorId) > 0 {
		input.DetectorId = aws.String(_guarddutyDetectorId)
	}
	if len(_guarddutyDisableEmailNotification) > 0 {
		if err := assignInputField(input, "DisableEmailNotification", _guarddutyDisableEmailNotification); err != nil {
			log.Errorf("invalid --disable-email-notification: %s", err.Error())
			return
		}
	}
	if len(_guarddutyMessage) > 0 {
		input.Message = aws.String(_guarddutyMessage)
	}

	if resp, err := client.InviteMembers(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists coverage details for your GuardDuty account. If you're a GuardDuty
// administrator, you can retrieve all resources associated with the active member
// accounts in your organization.
//
// Make sure the accounts have Runtime Monitoring enabled and GuardDuty agent
// running on their resources.
func guardduty_ListCoverage(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.ListCoverageInput{
		// DetectorId: *string, // Required
	}

	if len(_guarddutyDetectorId) > 0 {
		input.DetectorId = aws.String(_guarddutyDetectorId)
	}
	if len(_guarddutyFilterCriteria) > 0 {
		if err := assignInputField(input, "FilterCriteria", _guarddutyFilterCriteria); err != nil {
			log.Errorf("invalid --filter-criteria: %s", err.Error())
			return
		}
	}
	if len(_guarddutyMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _guarddutyMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_guarddutyNextToken) > 0 {
		input.NextToken = aws.String(_guarddutyNextToken)
	}
	if len(_guarddutySortCriteria) > 0 {
		if err := assignInputField(input, "SortCriteria", _guarddutySortCriteria); err != nil {
			log.Errorf("invalid --sort-criteria: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListCoverage(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*guardduty.ListCoverageOutput
	p := guardduty.NewListCoveragePaginator(client, input)
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

// Lists detectorIds of all the existing Amazon GuardDuty detector resources.
func guardduty_ListDetectors(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.ListDetectorsInput{}

	if len(_guarddutyMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _guarddutyMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_guarddutyNextToken) > 0 {
		input.NextToken = aws.String(_guarddutyNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDetectors(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*guardduty.ListDetectorsOutput
	p := guardduty.NewListDetectorsPaginator(client, input)
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

// Returns a paginated list of the current filters.
func guardduty_ListFilters(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.ListFiltersInput{
		// DetectorId: *string, // Required
	}

	if len(_guarddutyDetectorId) > 0 {
		input.DetectorId = aws.String(_guarddutyDetectorId)
	}
	if len(_guarddutyMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _guarddutyMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_guarddutyNextToken) > 0 {
		input.NextToken = aws.String(_guarddutyNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListFilters(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*guardduty.ListFiltersOutput
	p := guardduty.NewListFiltersPaginator(client, input)
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

// Lists GuardDuty findings for the specified detector ID.
// There might be regional differences because some flags might not be available
// in all the Regions where GuardDuty is currently supported. For more information,
// see [Regions and endpoints].
//
// [Regions and endpoints]: https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_regions.html
func guardduty_ListFindings(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.ListFindingsInput{
		// DetectorId: *string, // Required
	}

	if len(_guarddutyDetectorId) > 0 {
		input.DetectorId = aws.String(_guarddutyDetectorId)
	}
	if len(_guarddutyFindingCriteria) > 0 {
		if err := assignInputField(input, "FindingCriteria", _guarddutyFindingCriteria); err != nil {
			log.Errorf("invalid --finding-criteria: %s", err.Error())
			return
		}
	}
	if len(_guarddutyMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _guarddutyMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_guarddutyNextToken) > 0 {
		input.NextToken = aws.String(_guarddutyNextToken)
	}
	if len(_guarddutySortCriteria) > 0 {
		if err := assignInputField(input, "SortCriteria", _guarddutySortCriteria); err != nil {
			log.Errorf("invalid --sort-criteria: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListFindings(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*guardduty.ListFindingsOutput
	p := guardduty.NewListFindingsPaginator(client, input)
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

// Lists all GuardDuty membership invitations that were sent to the current Amazon
// Web Services account.
func guardduty_ListInvitations(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.ListInvitationsInput{}

	if len(_guarddutyMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _guarddutyMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_guarddutyNextToken) > 0 {
		input.NextToken = aws.String(_guarddutyNextToken)
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

	var results []*guardduty.ListInvitationsOutput
	p := guardduty.NewListInvitationsPaginator(client, input)
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

// Lists the IPSets of the GuardDuty service specified by the detector ID. If you
// use this operation from a member account, the IPSets returned are the IPSets
// from the associated administrator account.
func guardduty_ListIPSets(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.ListIPSetsInput{
		// DetectorId: *string, // Required
	}

	if len(_guarddutyDetectorId) > 0 {
		input.DetectorId = aws.String(_guarddutyDetectorId)
	}
	if len(_guarddutyMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _guarddutyMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_guarddutyNextToken) > 0 {
		input.NextToken = aws.String(_guarddutyNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListIPSets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*guardduty.ListIPSetsOutput
	p := guardduty.NewListIPSetsPaginator(client, input)
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

// Lists the Malware Protection plan IDs associated with the protected resources
// in your Amazon Web Services account.
func guardduty_ListMalwareProtectionPlans(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.ListMalwareProtectionPlansInput{}

	if len(_guarddutyNextToken) > 0 {
		input.NextToken = aws.String(_guarddutyNextToken)
	}

	if resp, err := client.ListMalwareProtectionPlans(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of malware scans. Each member account can view the malware scans
// for their own accounts. An administrator can view the malware scans for all of
// its members' accounts.
func guardduty_ListMalwareScans(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.ListMalwareScansInput{}

	if len(_guarddutyFilterCriteria) > 0 {
		if err := assignInputField(input, "FilterCriteria", _guarddutyFilterCriteria); err != nil {
			log.Errorf("invalid --filter-criteria: %s", err.Error())
			return
		}
	}
	if len(_guarddutyMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _guarddutyMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_guarddutyNextToken) > 0 {
		input.NextToken = aws.String(_guarddutyNextToken)
	}
	if len(_guarddutySortCriteria) > 0 {
		if err := assignInputField(input, "SortCriteria", _guarddutySortCriteria); err != nil {
			log.Errorf("invalid --sort-criteria: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListMalwareScans(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*guardduty.ListMalwareScansOutput
	p := guardduty.NewListMalwareScansPaginator(client, input)
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

// Lists details about all member accounts for the current GuardDuty administrator
// account.
func guardduty_ListMembers(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.ListMembersInput{
		// DetectorId: *string, // Required
	}

	if len(_guarddutyDetectorId) > 0 {
		input.DetectorId = aws.String(_guarddutyDetectorId)
	}
	if len(_guarddutyMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _guarddutyMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_guarddutyNextToken) > 0 {
		input.NextToken = aws.String(_guarddutyNextToken)
	}
	if len(_guarddutyOnlyAssociated) > 0 {
		input.OnlyAssociated = aws.String(_guarddutyOnlyAssociated)
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

	var results []*guardduty.ListMembersOutput
	p := guardduty.NewListMembersPaginator(client, input)
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

// Lists the accounts designated as GuardDuty delegated administrators. Only the
// organization's management account can run this API operation.
func guardduty_ListOrganizationAdminAccounts(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.ListOrganizationAdminAccountsInput{}

	if len(_guarddutyMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _guarddutyMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_guarddutyNextToken) > 0 {
		input.NextToken = aws.String(_guarddutyNextToken)
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

	var results []*guardduty.ListOrganizationAdminAccountsOutput
	p := guardduty.NewListOrganizationAdminAccountsPaginator(client, input)
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

// Returns a list of publishing destinations associated with the specified
// detectorId .
func guardduty_ListPublishingDestinations(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.ListPublishingDestinationsInput{
		// DetectorId: *string, // Required
	}

	if len(_guarddutyDetectorId) > 0 {
		input.DetectorId = aws.String(_guarddutyDetectorId)
	}
	if len(_guarddutyMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _guarddutyMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_guarddutyNextToken) > 0 {
		input.NextToken = aws.String(_guarddutyNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPublishingDestinations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*guardduty.ListPublishingDestinationsOutput
	p := guardduty.NewListPublishingDestinationsPaginator(client, input)
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

// Lists tags for a resource. Tagging is currently supported for detectors,
// finding filters, IP sets, threat intel sets, and publishing destination, with a
// limit of 50 tags per resource. When invoked, this operation returns all assigned
// tags for a given resource.
func guardduty_ListTagsForResource(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_guarddutyResourceArn) > 0 {
		input.ResourceArn = aws.String(_guarddutyResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the threat entity sets associated with the specified GuardDuty detector
// ID. If you use this operation from a member account, the threat entity sets that
// are returned as a response, belong to the administrator account.
func guardduty_ListThreatEntitySets(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.ListThreatEntitySetsInput{
		// DetectorId: *string, // Required
	}

	if len(_guarddutyDetectorId) > 0 {
		input.DetectorId = aws.String(_guarddutyDetectorId)
	}
	if len(_guarddutyMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _guarddutyMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_guarddutyNextToken) > 0 {
		input.NextToken = aws.String(_guarddutyNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListThreatEntitySets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*guardduty.ListThreatEntitySetsOutput
	p := guardduty.NewListThreatEntitySetsPaginator(client, input)
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

// Lists the ThreatIntelSets of the GuardDuty service specified by the detector
// ID. If you use this operation from a member account, the ThreatIntelSets
// associated with the administrator account are returned.
func guardduty_ListThreatIntelSets(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.ListThreatIntelSetsInput{
		// DetectorId: *string, // Required
	}

	if len(_guarddutyDetectorId) > 0 {
		input.DetectorId = aws.String(_guarddutyDetectorId)
	}
	if len(_guarddutyMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _guarddutyMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_guarddutyNextToken) > 0 {
		input.NextToken = aws.String(_guarddutyNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListThreatIntelSets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*guardduty.ListThreatIntelSetsOutput
	p := guardduty.NewListThreatIntelSetsPaginator(client, input)
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

// Lists the trusted entity sets associated with the specified GuardDuty detector
// ID. If you use this operation from a member account, the trusted entity sets
// that are returned as a response, belong to the administrator account.
func guardduty_ListTrustedEntitySets(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.ListTrustedEntitySetsInput{
		// DetectorId: *string, // Required
	}

	if len(_guarddutyDetectorId) > 0 {
		input.DetectorId = aws.String(_guarddutyDetectorId)
	}
	if len(_guarddutyMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _guarddutyMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_guarddutyNextToken) > 0 {
		input.NextToken = aws.String(_guarddutyNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTrustedEntitySets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*guardduty.ListTrustedEntitySetsOutput
	p := guardduty.NewListTrustedEntitySetsPaginator(client, input)
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

// Initiates a malware scan for a specific S3 object. This API allows you to
// perform on-demand malware scanning of individual objects in S3 buckets that have
// Malware Protection for S3 enabled.
//
// When you use this API, the Amazon Web Services service terms for GuardDuty
// Malware Protection apply. For more information, see [Amazon Web Services service terms for GuardDuty Malware Protection].
//
// [Amazon Web Services service terms for GuardDuty Malware Protection]: http://aws.amazon.com/service-terms/#87._Amazon_GuardDuty
func guardduty_SendObjectMalwareScan(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.SendObjectMalwareScanInput{}

	if len(_guarddutyS3Object) > 0 {
		if err := assignInputField(input, "S3Object", _guarddutyS3Object); err != nil {
			log.Errorf("invalid --s3-object: %s", err.Error())
			return
		}
	}

	if resp, err := client.SendObjectMalwareScan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Initiates the malware scan. Invoking this API will automatically create the [Service-linked role] in
// the corresponding account if the resourceArn belongs to an EC2 instance.
//
// When the malware scan starts, you can use the associated scan ID to track the
// status of the scan. For more information, see [ListMalwareScans]and [GetMalwareScan].
//
// When you use this API, the Amazon Web Services service terms for GuardDuty
// Malware Protection apply. For more information, see [Amazon Web Services service terms for GuardDuty Malware Protection].
//
// [GetMalwareScan]: https://docs.aws.amazon.com/guardduty/latest/APIReference/API_GetMalwareScan.html
// [Amazon Web Services service terms for GuardDuty Malware Protection]: http://aws.amazon.com/service-terms/#87._Amazon_GuardDuty
// [ListMalwareScans]: https://docs.aws.amazon.com/guardduty/latest/APIReference/API_ListMalwareScans.html
// [Service-linked role]: https://docs.aws.amazon.com/guardduty/latest/ug/slr-permissions-malware-protection.html
func guardduty_StartMalwareScan(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.StartMalwareScanInput{
		// ResourceArn: *string, // Required
	}

	if len(_guarddutyResourceArn) > 0 {
		input.ResourceArn = aws.String(_guarddutyResourceArn)
	}
	if len(_guarddutyClientToken) > 0 {
		input.ClientToken = aws.String(_guarddutyClientToken)
	}
	if len(_guarddutyScanConfiguration) > 0 {
		if err := assignInputField(input, "ScanConfiguration", _guarddutyScanConfiguration); err != nil {
			log.Errorf("invalid --scan-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartMalwareScan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Turns on GuardDuty monitoring of the specified member accounts. Use this
// operation to restart monitoring of accounts that you stopped monitoring with the
// [StopMonitoringMembers]operation.
//
// [StopMonitoringMembers]: https://docs.aws.amazon.com/guardduty/latest/APIReference/API_StopMonitoringMembers.html
func guardduty_StartMonitoringMembers(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.StartMonitoringMembersInput{
		// AccountIds: []string, // Required
		// DetectorId: *string, // Required
	}

	if len(_guarddutyAccountIds) > 0 {
		input.AccountIds = append([]string(nil), _guarddutyAccountIds...)
	}
	if len(_guarddutyDetectorId) > 0 {
		input.DetectorId = aws.String(_guarddutyDetectorId)
	}

	if resp, err := client.StartMonitoringMembers(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops GuardDuty monitoring for the specified member accounts. Use the
// StartMonitoringMembers operation to restart monitoring for those accounts.
//
// With autoEnableOrganizationMembers configuration for your organization set to
// ALL , you'll receive an error if you attempt to stop monitoring the member
// accounts in your organization.
func guardduty_StopMonitoringMembers(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.StopMonitoringMembersInput{
		// AccountIds: []string, // Required
		// DetectorId: *string, // Required
	}

	if len(_guarddutyAccountIds) > 0 {
		input.AccountIds = append([]string(nil), _guarddutyAccountIds...)
	}
	if len(_guarddutyDetectorId) > 0 {
		input.DetectorId = aws.String(_guarddutyDetectorId)
	}

	if resp, err := client.StopMonitoringMembers(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds tags to a resource.
func guardduty_TagResource(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_guarddutyResourceArn) > 0 {
		input.ResourceArn = aws.String(_guarddutyResourceArn)
	}
	if len(_guarddutyTags) > 0 {
		if err := assignInputField(input, "Tags", _guarddutyTags); err != nil {
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

// Unarchives GuardDuty findings specified by the findingIds .
func guardduty_UnarchiveFindings(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.UnarchiveFindingsInput{
		// DetectorId: *string, // Required
		// FindingIds: []string, // Required
	}

	if len(_guarddutyDetectorId) > 0 {
		input.DetectorId = aws.String(_guarddutyDetectorId)
	}
	if len(_guarddutyFindingIds) > 0 {
		input.FindingIds = append([]string(nil), _guarddutyFindingIds...)
	}

	if resp, err := client.UnarchiveFindings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes tags from a resource.
func guardduty_UntagResource(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_guarddutyResourceArn) > 0 {
		input.ResourceArn = aws.String(_guarddutyResourceArn)
	}
	if len(_guarddutyTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _guarddutyTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the GuardDuty detector specified by the detector ID.
// Specifying both EKS Runtime Monitoring ( EKS_RUNTIME_MONITORING ) and Runtime
// Monitoring ( RUNTIME_MONITORING ) will cause an error. You can add only one of
// these two features because Runtime Monitoring already includes the threat
// detection for Amazon EKS resources. For more information, see [Runtime Monitoring].
//
// There might be regional differences because some data sources might not be
// available in all the Amazon Web Services Regions where GuardDuty is presently
// supported. For more information, see [Regions and endpoints].
//
// [Regions and endpoints]: https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_regions.html
// [Runtime Monitoring]: https://docs.aws.amazon.com/guardduty/latest/ug/runtime-monitoring.html
func guardduty_UpdateDetector(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.UpdateDetectorInput{
		// DetectorId: *string, // Required
	}

	if len(_guarddutyDetectorId) > 0 {
		input.DetectorId = aws.String(_guarddutyDetectorId)
	}
	if len(_guarddutyDataSources) > 0 {
		if err := assignInputField(input, "DataSources", _guarddutyDataSources); err != nil {
			log.Errorf("invalid --data-sources: %s", err.Error())
			return
		}
	}
	if len(_guarddutyEnable) > 0 {
		if err := assignInputField(input, "Enable", _guarddutyEnable); err != nil {
			log.Errorf("invalid --enable: %s", err.Error())
			return
		}
	}
	if len(_guarddutyFeatures) > 0 {
		if err := assignInputField(input, "Features", _guarddutyFeatures); err != nil {
			log.Errorf("invalid --features: %s", err.Error())
			return
		}
	}
	if len(_guarddutyFindingPublishingFrequency) > 0 {
		if err := assignInputField(input, "FindingPublishingFrequency", _guarddutyFindingPublishingFrequency); err != nil {
			log.Errorf("invalid --finding-publishing-frequency: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateDetector(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the filter specified by the filter name.
func guardduty_UpdateFilter(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.UpdateFilterInput{
		// DetectorId: *string, // Required
		// FilterName: *string, // Required
	}

	if len(_guarddutyDetectorId) > 0 {
		input.DetectorId = aws.String(_guarddutyDetectorId)
	}
	if len(_guarddutyFilterName) > 0 {
		input.FilterName = aws.String(_guarddutyFilterName)
	}
	if len(_guarddutyAction) > 0 {
		if err := assignInputField(input, "Action", _guarddutyAction); err != nil {
			log.Errorf("invalid --action: %s", err.Error())
			return
		}
	}
	if len(_guarddutyDescription) > 0 {
		input.Description = aws.String(_guarddutyDescription)
	}
	if len(_guarddutyFindingCriteria) > 0 {
		if err := assignInputField(input, "FindingCriteria", _guarddutyFindingCriteria); err != nil {
			log.Errorf("invalid --finding-criteria: %s", err.Error())
			return
		}
	}
	if len(_guarddutyRank) > 0 {
		if err := assignInputField(input, "Rank", _guarddutyRank); err != nil {
			log.Errorf("invalid --rank: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateFilter(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Marks the specified GuardDuty findings as useful or not useful.
func guardduty_UpdateFindingsFeedback(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.UpdateFindingsFeedbackInput{
		// DetectorId: *string, // Required
		// Feedback: types.Feedback, // Required
		// FindingIds: []string, // Required
	}

	if len(_guarddutyDetectorId) > 0 {
		input.DetectorId = aws.String(_guarddutyDetectorId)
	}
	if len(_guarddutyFeedback) > 0 {
		if err := assignInputField(input, "Feedback", _guarddutyFeedback); err != nil {
			log.Errorf("invalid --feedback: %s", err.Error())
			return
		}
	}
	if len(_guarddutyFindingIds) > 0 {
		input.FindingIds = append([]string(nil), _guarddutyFindingIds...)
	}
	if len(_guarddutyComments) > 0 {
		input.Comments = aws.String(_guarddutyComments)
	}

	if resp, err := client.UpdateFindingsFeedback(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the IPSet specified by the IPSet ID.
func guardduty_UpdateIPSet(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.UpdateIPSetInput{
		// DetectorId: *string, // Required
		// IpSetId: *string, // Required
	}

	if len(_guarddutyDetectorId) > 0 {
		input.DetectorId = aws.String(_guarddutyDetectorId)
	}
	if len(_guarddutyIpSetId) > 0 {
		input.IpSetId = aws.String(_guarddutyIpSetId)
	}
	if len(_guarddutyActivate) > 0 {
		if err := assignInputField(input, "Activate", _guarddutyActivate); err != nil {
			log.Errorf("invalid --activate: %s", err.Error())
			return
		}
	}
	if len(_guarddutyExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_guarddutyExpectedBucketOwner)
	}
	if len(_guarddutyLocation) > 0 {
		input.Location = aws.String(_guarddutyLocation)
	}
	if len(_guarddutyName) > 0 {
		input.Name = aws.String(_guarddutyName)
	}

	if resp, err := client.UpdateIPSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing Malware Protection plan resource.
func guardduty_UpdateMalwareProtectionPlan(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.UpdateMalwareProtectionPlanInput{
		// MalwareProtectionPlanId: *string, // Required
	}

	if len(_guarddutyMalwareProtectionPlanId) > 0 {
		input.MalwareProtectionPlanId = aws.String(_guarddutyMalwareProtectionPlanId)
	}
	if len(_guarddutyActions) > 0 {
		if err := assignInputField(input, "Actions", _guarddutyActions); err != nil {
			log.Errorf("invalid --actions: %s", err.Error())
			return
		}
	}
	if len(_guarddutyProtectedResource) > 0 {
		if err := assignInputField(input, "ProtectedResource", _guarddutyProtectedResource); err != nil {
			log.Errorf("invalid --protected-resource: %s", err.Error())
			return
		}
	}
	if len(_guarddutyRole) > 0 {
		input.Role = aws.String(_guarddutyRole)
	}

	if resp, err := client.UpdateMalwareProtectionPlan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the malware scan settings.
// There might be regional differences because some data sources might not be
// available in all the Amazon Web Services Regions where GuardDuty is presently
// supported. For more information, see [Regions and endpoints].
//
// [Regions and endpoints]: https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_regions.html
func guardduty_UpdateMalwareScanSettings(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.UpdateMalwareScanSettingsInput{
		// DetectorId: *string, // Required
	}

	if len(_guarddutyDetectorId) > 0 {
		input.DetectorId = aws.String(_guarddutyDetectorId)
	}
	if len(_guarddutyEbsSnapshotPreservation) > 0 {
		if err := assignInputField(input, "EbsSnapshotPreservation", _guarddutyEbsSnapshotPreservation); err != nil {
			log.Errorf("invalid --ebs-snapshot-preservation: %s", err.Error())
			return
		}
	}
	if len(_guarddutyScanResourceCriteria) > 0 {
		if err := assignInputField(input, "ScanResourceCriteria", _guarddutyScanResourceCriteria); err != nil {
			log.Errorf("invalid --scan-resource-criteria: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateMalwareScanSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Contains information on member accounts to be updated.
// Specifying both EKS Runtime Monitoring ( EKS_RUNTIME_MONITORING ) and Runtime
// Monitoring ( RUNTIME_MONITORING ) will cause an error. You can add only one of
// these two features because Runtime Monitoring already includes the threat
// detection for Amazon EKS resources. For more information, see [Runtime Monitoring].
//
// There might be regional differences because some data sources might not be
// available in all the Amazon Web Services Regions where GuardDuty is presently
// supported. For more information, see [Regions and endpoints].
//
// [Regions and endpoints]: https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_regions.html
// [Runtime Monitoring]: https://docs.aws.amazon.com/guardduty/latest/ug/runtime-monitoring.html
func guardduty_UpdateMemberDetectors(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.UpdateMemberDetectorsInput{
		// AccountIds: []string, // Required
		// DetectorId: *string, // Required
	}

	if len(_guarddutyAccountIds) > 0 {
		input.AccountIds = append([]string(nil), _guarddutyAccountIds...)
	}
	if len(_guarddutyDetectorId) > 0 {
		input.DetectorId = aws.String(_guarddutyDetectorId)
	}
	if len(_guarddutyDataSources) > 0 {
		if err := assignInputField(input, "DataSources", _guarddutyDataSources); err != nil {
			log.Errorf("invalid --data-sources: %s", err.Error())
			return
		}
	}
	if len(_guarddutyFeatures) > 0 {
		if err := assignInputField(input, "Features", _guarddutyFeatures); err != nil {
			log.Errorf("invalid --features: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateMemberDetectors(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Configures the delegated administrator account with the provided values. You
// must provide a value for either autoEnableOrganizationMembers or autoEnable ,
// but not both.
//
// Specifying both EKS Runtime Monitoring ( EKS_RUNTIME_MONITORING ) and Runtime
// Monitoring ( RUNTIME_MONITORING ) will cause an error. You can add only one of
// these two features because Runtime Monitoring already includes the threat
// detection for Amazon EKS resources. For more information, see [Runtime Monitoring].
//
// There might be regional differences because some data sources might not be
// available in all the Amazon Web Services Regions where GuardDuty is presently
// supported. For more information, see [Regions and endpoints].
//
// [Regions and endpoints]: https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_regions.html
// [Runtime Monitoring]: https://docs.aws.amazon.com/guardduty/latest/ug/runtime-monitoring.html
func guardduty_UpdateOrganizationConfiguration(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.UpdateOrganizationConfigurationInput{
		// DetectorId: *string, // Required
	}

	if len(_guarddutyDetectorId) > 0 {
		input.DetectorId = aws.String(_guarddutyDetectorId)
	}
	if len(_guarddutyAutoEnable) > 0 {
		if err := assignInputField(input, "AutoEnable", _guarddutyAutoEnable); err != nil {
			log.Errorf("invalid --auto-enable: %s", err.Error())
			return
		}
	}
	if len(_guarddutyAutoEnableOrganizationMembers) > 0 {
		if err := assignInputField(input, "AutoEnableOrganizationMembers", _guarddutyAutoEnableOrganizationMembers); err != nil {
			log.Errorf("invalid --auto-enable-organization-members: %s", err.Error())
			return
		}
	}
	if len(_guarddutyDataSources) > 0 {
		if err := assignInputField(input, "DataSources", _guarddutyDataSources); err != nil {
			log.Errorf("invalid --data-sources: %s", err.Error())
			return
		}
	}
	if len(_guarddutyFeatures) > 0 {
		if err := assignInputField(input, "Features", _guarddutyFeatures); err != nil {
			log.Errorf("invalid --features: %s", err.Error())
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

// Updates information about the publishing destination specified by the
// destinationId .
func guardduty_UpdatePublishingDestination(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.UpdatePublishingDestinationInput{
		// DestinationId: *string, // Required
		// DetectorId: *string, // Required
	}

	if len(_guarddutyDestinationId) > 0 {
		input.DestinationId = aws.String(_guarddutyDestinationId)
	}
	if len(_guarddutyDetectorId) > 0 {
		input.DetectorId = aws.String(_guarddutyDetectorId)
	}
	if len(_guarddutyDestinationProperties) > 0 {
		if err := assignInputField(input, "DestinationProperties", _guarddutyDestinationProperties); err != nil {
			log.Errorf("invalid --destination-properties: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdatePublishingDestination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the threat entity set associated with the specified threatEntitySetId .
func guardduty_UpdateThreatEntitySet(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.UpdateThreatEntitySetInput{
		// DetectorId: *string, // Required
		// ThreatEntitySetId: *string, // Required
	}

	if len(_guarddutyDetectorId) > 0 {
		input.DetectorId = aws.String(_guarddutyDetectorId)
	}
	if len(_guarddutyThreatEntitySetId) > 0 {
		input.ThreatEntitySetId = aws.String(_guarddutyThreatEntitySetId)
	}
	if len(_guarddutyActivate) > 0 {
		if err := assignInputField(input, "Activate", _guarddutyActivate); err != nil {
			log.Errorf("invalid --activate: %s", err.Error())
			return
		}
	}
	if len(_guarddutyExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_guarddutyExpectedBucketOwner)
	}
	if len(_guarddutyLocation) > 0 {
		input.Location = aws.String(_guarddutyLocation)
	}
	if len(_guarddutyName) > 0 {
		input.Name = aws.String(_guarddutyName)
	}

	if resp, err := client.UpdateThreatEntitySet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the ThreatIntelSet specified by the ThreatIntelSet ID.
func guardduty_UpdateThreatIntelSet(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.UpdateThreatIntelSetInput{
		// DetectorId: *string, // Required
		// ThreatIntelSetId: *string, // Required
	}

	if len(_guarddutyDetectorId) > 0 {
		input.DetectorId = aws.String(_guarddutyDetectorId)
	}
	if len(_guarddutyThreatIntelSetId) > 0 {
		input.ThreatIntelSetId = aws.String(_guarddutyThreatIntelSetId)
	}
	if len(_guarddutyActivate) > 0 {
		if err := assignInputField(input, "Activate", _guarddutyActivate); err != nil {
			log.Errorf("invalid --activate: %s", err.Error())
			return
		}
	}
	if len(_guarddutyExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_guarddutyExpectedBucketOwner)
	}
	if len(_guarddutyLocation) > 0 {
		input.Location = aws.String(_guarddutyLocation)
	}
	if len(_guarddutyName) > 0 {
		input.Name = aws.String(_guarddutyName)
	}

	if resp, err := client.UpdateThreatIntelSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the trusted entity set associated with the specified trustedEntitySetId .
func guardduty_UpdateTrustedEntitySet(cfg aws.Config, client *guardduty.Client) {
	input := &guardduty.UpdateTrustedEntitySetInput{
		// DetectorId: *string, // Required
		// TrustedEntitySetId: *string, // Required
	}

	if len(_guarddutyDetectorId) > 0 {
		input.DetectorId = aws.String(_guarddutyDetectorId)
	}
	if len(_guarddutyTrustedEntitySetId) > 0 {
		input.TrustedEntitySetId = aws.String(_guarddutyTrustedEntitySetId)
	}
	if len(_guarddutyActivate) > 0 {
		if err := assignInputField(input, "Activate", _guarddutyActivate); err != nil {
			log.Errorf("invalid --activate: %s", err.Error())
			return
		}
	}
	if len(_guarddutyExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_guarddutyExpectedBucketOwner)
	}
	if len(_guarddutyLocation) > 0 {
		input.Location = aws.String(_guarddutyLocation)
	}
	if len(_guarddutyName) > 0 {
		input.Name = aws.String(_guarddutyName)
	}

	if resp, err := client.UpdateTrustedEntitySet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_guarddutyCmd)
	_guarddutyCmd.Flags().SortFlags = false

	_guarddutyCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_guarddutyCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_guarddutyCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_guarddutyCmd.Flags().StringVarP(&_guarddutyAccountDetails, "account-details", "", "", "Account Details")
	_guarddutyCmd.Flags().StringSliceVarP(&_guarddutyAccountIds, "account-ids", "", nil, "Account Ids")
	_guarddutyCmd.Flags().StringVarP(&_guarddutyAction, "action", "", "", "Action")
	_guarddutyCmd.Flags().StringVarP(&_guarddutyActions, "actions", "", "", "Actions")
	_guarddutyCmd.Flags().StringVarP(&_guarddutyActivate, "activate", "", "", "Activate")
	_guarddutyCmd.Flags().StringVarP(&_guarddutyAdminAccountId, "admin-account-id", "", "", "Admin Account ID")
	_guarddutyCmd.Flags().StringVarP(&_guarddutyAdministratorId, "administrator-id", "", "", "Administrator ID")
	_guarddutyCmd.Flags().StringVarP(&_guarddutyAutoEnable, "auto-enable", "", "", "Auto Enable")
	_guarddutyCmd.Flags().StringVarP(&_guarddutyAutoEnableOrganizationMembers, "auto-enable-organization-members", "", "", "Auto Enable Organization Members")
	_guarddutyCmd.Flags().StringVarP(&_guarddutyClientToken, "client-token", "", "", "Client Token")
	_guarddutyCmd.Flags().StringVarP(&_guarddutyComments, "comments", "", "", "Comments")
	_guarddutyCmd.Flags().StringVarP(&_guarddutyDataSources, "data-sources", "", "", "Data Sources")
	_guarddutyCmd.Flags().StringVarP(&_guarddutyDescription, "description", "", "", "Description")
	_guarddutyCmd.Flags().StringVarP(&_guarddutyDestinationId, "destination-id", "", "", "Destination ID")
	_guarddutyCmd.Flags().StringVarP(&_guarddutyDestinationProperties, "destination-properties", "", "", "Destination Properties")
	_guarddutyCmd.Flags().StringVarP(&_guarddutyDestinationType, "destination-type", "", "", "Destination Type")
	_guarddutyCmd.Flags().StringVarP(&_guarddutyDetectorId, "detector-id", "", "", "Detector ID")
	_guarddutyCmd.Flags().StringVarP(&_guarddutyDisableEmailNotification, "disable-email-notification", "", "", "Disable Email Notification")
	_guarddutyCmd.Flags().StringVarP(&_guarddutyEbsSnapshotPreservation, "ebs-snapshot-preservation", "", "", "Ebs Snapshot Preservation")
	_guarddutyCmd.Flags().StringVarP(&_guarddutyEnable, "enable", "", "", "Enable")
	_guarddutyCmd.Flags().StringVarP(&_guarddutyExpectedBucketOwner, "expected-bucket-owner", "", "", "Expected Bucket Owner")
	_guarddutyCmd.Flags().StringVarP(&_guarddutyFeatures, "features", "", "", "Features")
	_guarddutyCmd.Flags().StringVarP(&_guarddutyFeedback, "feedback", "", "", "Feedback")
	_guarddutyCmd.Flags().StringVarP(&_guarddutyFilterCriteria, "filter-criteria", "", "", "Filter Criteria")
	_guarddutyCmd.Flags().StringVarP(&_guarddutyFilterName, "filter-name", "", "", "Filter Name")
	_guarddutyCmd.Flags().StringVarP(&_guarddutyFindingCriteria, "finding-criteria", "", "", "Finding Criteria")
	_guarddutyCmd.Flags().StringSliceVarP(&_guarddutyFindingIds, "finding-ids", "", nil, "Finding Ids")
	_guarddutyCmd.Flags().StringVarP(&_guarddutyFindingPublishingFrequency, "finding-publishing-frequency", "", "", "Finding Publishing Frequency")
	_guarddutyCmd.Flags().StringVarP(&_guarddutyFindingStatisticTypes, "finding-statistic-types", "", "", "Finding Statistic Types")
	_guarddutyCmd.Flags().StringSliceVarP(&_guarddutyFindingTypes, "finding-types", "", nil, "Finding Types")
	_guarddutyCmd.Flags().StringVarP(&_guarddutyFormat, "format", "", "", "Format")
	_guarddutyCmd.Flags().StringVarP(&_guarddutyGroupBy, "group-by", "", "", "Group By")
	_guarddutyCmd.Flags().StringVarP(&_guarddutyInvitationId, "invitation-id", "", "", "Invitation ID")
	_guarddutyCmd.Flags().StringVarP(&_guarddutyIpSetId, "ip-set-id", "", "", "IP Set ID")
	_guarddutyCmd.Flags().StringVarP(&_guarddutyLocation, "location", "", "", "Location")
	_guarddutyCmd.Flags().StringVarP(&_guarddutyMalwareProtectionPlanId, "malware-protection-plan-id", "", "", "Malware Protection Plan ID")
	_guarddutyCmd.Flags().StringVarP(&_guarddutyMasterId, "master-id", "", "", "Master ID")
	_guarddutyCmd.Flags().StringVarP(&_guarddutyMaxResults, "max-results", "", "", "Max Results")
	_guarddutyCmd.Flags().StringVarP(&_guarddutyMessage, "message", "", "", "Message")
	_guarddutyCmd.Flags().StringVarP(&_guarddutyName, "name", "", "", "Name")
	_guarddutyCmd.Flags().StringVarP(&_guarddutyNextToken, "next-token", "", "", "Next Token")
	_guarddutyCmd.Flags().StringVarP(&_guarddutyOnlyAssociated, "only-associated", "", "", "Only Associated")
	_guarddutyCmd.Flags().StringVarP(&_guarddutyOrderBy, "order-by", "", "", "Order By")
	_guarddutyCmd.Flags().StringVarP(&_guarddutyProtectedResource, "protected-resource", "", "", "Protected Resource")
	_guarddutyCmd.Flags().StringVarP(&_guarddutyRank, "rank", "", "", "Rank")
	_guarddutyCmd.Flags().StringVarP(&_guarddutyResourceArn, "resource-arn", "", "", "Resource ARN")
	_guarddutyCmd.Flags().StringVarP(&_guarddutyRole, "role", "", "", "Role")
	_guarddutyCmd.Flags().StringVarP(&_guarddutyS3Object, "s3-object", "", "", "S3 Object")
	_guarddutyCmd.Flags().StringVarP(&_guarddutyScanConfiguration, "scan-configuration", "", "", "Scan Configuration")
	_guarddutyCmd.Flags().StringVarP(&_guarddutyScanId, "scan-id", "", "", "Scan ID")
	_guarddutyCmd.Flags().StringVarP(&_guarddutyScanResourceCriteria, "scan-resource-criteria", "", "", "Scan Resource Criteria")
	_guarddutyCmd.Flags().StringVarP(&_guarddutySortCriteria, "sort-criteria", "", "", "Sort Criteria")
	_guarddutyCmd.Flags().StringVarP(&_guarddutyStatisticsType, "statistics-type", "", "", "Statistics Type")
	_guarddutyCmd.Flags().StringSliceVarP(&_guarddutyTagKeys, "tag-keys", "", nil, "Tag Keys")
	_guarddutyCmd.Flags().StringVarP(&_guarddutyTags, "tags", "", "", "Tags")
	_guarddutyCmd.Flags().StringVarP(&_guarddutyThreatEntitySetId, "threat-entity-set-id", "", "", "Threat Entity Set ID")
	_guarddutyCmd.Flags().StringVarP(&_guarddutyThreatIntelSetId, "threat-intel-set-id", "", "", "Threat Intel Set ID")
	_guarddutyCmd.Flags().StringVarP(&_guarddutyTrustedEntitySetId, "trusted-entity-set-id", "", "", "Trusted Entity Set ID")
	_guarddutyCmd.Flags().StringVarP(&_guarddutyUnit, "unit", "", "", "Unit")
	_guarddutyCmd.Flags().StringVarP(&_guarddutyUsageCriteria, "usage-criteria", "", "", "Usage Criteria")
	_guarddutyCmd.Flags().StringVarP(&_guarddutyUsageStatisticType, "usage-statistic-type", "", "", "Usage Statistic Type")

	_guarddutyCmd.Flags().BoolVarP(&_guarddutyAcceptAdministratorInvitation, "accept-administrator-invitation", "", false, "Accept Administrator Invitation")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyAcceptInvitation, "accept-invitation", "", false, "Accept Invitation")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyArchiveFindings, "archive-findings", "", false, "Archive Findings")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyCreateDetector, "create-detector", "", false, "Create Detector")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyCreateFilter, "create-filter", "", false, "Create Filter")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyCreateIPSet, "create-ip-set", "", false, "Create IP Set")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyCreateMalwareProtectionPlan, "create-malware-protection-plan", "", false, "Create Malware Protection Plan")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyCreateMembers, "create-members", "", false, "Create Members")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyCreatePublishingDestination, "create-publishing-destination", "", false, "Create Publishing Destination")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyCreateSampleFindings, "create-sample-findings", "", false, "Create Sample Findings")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyCreateThreatEntitySet, "create-threat-entity-set", "", false, "Create Threat Entity Set")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyCreateThreatIntelSet, "create-threat-intel-set", "", false, "Create Threat Intel Set")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyCreateTrustedEntitySet, "create-trusted-entity-set", "", false, "Create Trusted Entity Set")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyDeclineInvitations, "decline-invitations", "", false, "Decline Invitations")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyDeleteDetector, "delete-detector", "", false, "Delete Detector")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyDeleteFilter, "delete-filter", "", false, "Delete Filter")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyDeleteInvitations, "delete-invitations", "", false, "Delete Invitations")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyDeleteIPSet, "delete-ip-set", "", false, "Delete IP Set")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyDeleteMalwareProtectionPlan, "delete-malware-protection-plan", "", false, "Delete Malware Protection Plan")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyDeleteMembers, "delete-members", "", false, "Delete Members")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyDeletePublishingDestination, "delete-publishing-destination", "", false, "Delete Publishing Destination")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyDeleteThreatEntitySet, "delete-threat-entity-set", "", false, "Delete Threat Entity Set")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyDeleteThreatIntelSet, "delete-threat-intel-set", "", false, "Delete Threat Intel Set")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyDeleteTrustedEntitySet, "delete-trusted-entity-set", "", false, "Delete Trusted Entity Set")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyDescribeMalwareScans, "describe-malware-scans", "", false, "Describe Malware Scans")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyDescribeOrganizationConfiguration, "describe-organization-configuration", "", false, "Describe Organization Configuration")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyDescribePublishingDestination, "describe-publishing-destination", "", false, "Describe Publishing Destination")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyDisableOrganizationAdminAccount, "disable-organization-admin-account", "", false, "Disable Organization Admin Account")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyDisassociateFromAdministratorAccount, "disassociate-from-administrator-account", "", false, "Disassociate From Administrator Account")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyDisassociateFromMasterAccount, "disassociate-from-master-account", "", false, "Disassociate From Master Account")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyDisassociateMembers, "disassociate-members", "", false, "Disassociate Members")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyEnableOrganizationAdminAccount, "enable-organization-admin-account", "", false, "Enable Organization Admin Account")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyGetAdministratorAccount, "get-administrator-account", "", false, "Get Administrator Account")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyGetCoverageStatistics, "get-coverage-statistics", "", false, "Get Coverage Statistics")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyGetDetector, "get-detector", "", false, "Get Detector")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyGetFilter, "get-filter", "", false, "Get Filter")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyGetFindings, "get-findings", "", false, "Get Findings")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyGetFindingsStatistics, "get-findings-statistics", "", false, "Get Findings Statistics")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyGetInvitationsCount, "get-invitations-count", "", false, "Get Invitations Count")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyGetIPSet, "get-ip-set", "", false, "Get IP Set")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyGetMalwareProtectionPlan, "get-malware-protection-plan", "", false, "Get Malware Protection Plan")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyGetMalwareScan, "get-malware-scan", "", false, "Get Malware Scan")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyGetMalwareScanSettings, "get-malware-scan-settings", "", false, "Get Malware Scan Settings")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyGetMasterAccount, "get-master-account", "", false, "Get Master Account")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyGetMemberDetectors, "get-member-detectors", "", false, "Get Member Detectors")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyGetMembers, "get-members", "", false, "Get Members")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyGetOrganizationStatistics, "get-organization-statistics", "", false, "Get Organization Statistics")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyGetRemainingFreeTrialDays, "get-remaining-free-trial-days", "", false, "Get Remaining Free Trial Days")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyGetThreatEntitySet, "get-threat-entity-set", "", false, "Get Threat Entity Set")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyGetThreatIntelSet, "get-threat-intel-set", "", false, "Get Threat Intel Set")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyGetTrustedEntitySet, "get-trusted-entity-set", "", false, "Get Trusted Entity Set")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyGetUsageStatistics, "get-usage-statistics", "", false, "Get Usage Statistics")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyInviteMembers, "invite-members", "", false, "Invite Members")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyListCoverage, "list-coverage", "", false, "List Coverage")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyListDetectors, "list-detectors", "", false, "List Detectors")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyListFilters, "list-filters", "", false, "List Filters")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyListFindings, "list-findings", "", false, "List Findings")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyListInvitations, "list-invitations", "", false, "List Invitations")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyListIPSets, "list-ip-sets", "", false, "List IP Sets")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyListMalwareProtectionPlans, "list-malware-protection-plans", "", false, "List Malware Protection Plans")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyListMalwareScans, "list-malware-scans", "", false, "List Malware Scans")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyListMembers, "list-members", "", false, "List Members")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyListOrganizationAdminAccounts, "list-organization-admin-accounts", "", false, "List Organization Admin Accounts")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyListPublishingDestinations, "list-publishing-destinations", "", false, "List Publishing Destinations")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyListThreatEntitySets, "list-threat-entity-sets", "", false, "List Threat Entity Sets")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyListThreatIntelSets, "list-threat-intel-sets", "", false, "List Threat Intel Sets")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyListTrustedEntitySets, "list-trusted-entity-sets", "", false, "List Trusted Entity Sets")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutySendObjectMalwareScan, "send-object-malware-scan", "", false, "Send Object Malware Scan")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyStartMalwareScan, "start-malware-scan", "", false, "Start Malware Scan")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyStartMonitoringMembers, "start-monitoring-members", "", false, "Start Monitoring Members")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyStopMonitoringMembers, "stop-monitoring-members", "", false, "Stop Monitoring Members")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyTagResource, "tag-resource", "", false, "Tag Resource")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyUnarchiveFindings, "unarchive-findings", "", false, "Unarchive Findings")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyUntagResource, "untag-resource", "", false, "Untag Resource")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyUpdateDetector, "update-detector", "", false, "Update Detector")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyUpdateFilter, "update-filter", "", false, "Update Filter")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyUpdateFindingsFeedback, "update-findings-feedback", "", false, "Update Findings Feedback")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyUpdateIPSet, "update-ip-set", "", false, "Update IP Set")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyUpdateMalwareProtectionPlan, "update-malware-protection-plan", "", false, "Update Malware Protection Plan")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyUpdateMalwareScanSettings, "update-malware-scan-settings", "", false, "Update Malware Scan Settings")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyUpdateMemberDetectors, "update-member-detectors", "", false, "Update Member Detectors")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyUpdateOrganizationConfiguration, "update-organization-configuration", "", false, "Update Organization Configuration")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyUpdatePublishingDestination, "update-publishing-destination", "", false, "Update Publishing Destination")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyUpdateThreatEntitySet, "update-threat-entity-set", "", false, "Update Threat Entity Set")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyUpdateThreatIntelSet, "update-threat-intel-set", "", false, "Update Threat Intel Set")
	_guarddutyCmd.Flags().BoolVarP(&_guarddutyUpdateTrustedEntitySet, "update-trusted-entity-set", "", false, "Update Trusted Entity Set")

}
