package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/macie2"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// macie2Cmd represents the macie2 command
var _macie2Cmd = &cobra.Command{
	Use:   "macie2",
	Short: "AWS macie2 CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := macie2.NewFromConfig(cfg)
		if _macie2AcceptInvitation {
			macie2_AcceptInvitation(cfg, client)
			return
		}
		if _macie2BatchGetCustomDataIdentifiers {
			macie2_BatchGetCustomDataIdentifiers(cfg, client)
			return
		}
		if _macie2BatchUpdateAutomatedDiscoveryAccounts {
			macie2_BatchUpdateAutomatedDiscoveryAccounts(cfg, client)
			return
		}
		if _macie2CreateAllowList {
			macie2_CreateAllowList(cfg, client)
			return
		}
		if _macie2CreateClassificationJob {
			macie2_CreateClassificationJob(cfg, client)
			return
		}
		if _macie2CreateCustomDataIdentifier {
			macie2_CreateCustomDataIdentifier(cfg, client)
			return
		}
		if _macie2CreateFindingsFilter {
			macie2_CreateFindingsFilter(cfg, client)
			return
		}
		if _macie2CreateInvitations {
			macie2_CreateInvitations(cfg, client)
			return
		}
		if _macie2CreateMember {
			macie2_CreateMember(cfg, client)
			return
		}
		if _macie2CreateSampleFindings {
			macie2_CreateSampleFindings(cfg, client)
			return
		}
		if _macie2DeclineInvitations {
			macie2_DeclineInvitations(cfg, client)
			return
		}
		if _macie2DeleteAllowList {
			macie2_DeleteAllowList(cfg, client)
			return
		}
		if _macie2DeleteCustomDataIdentifier {
			macie2_DeleteCustomDataIdentifier(cfg, client)
			return
		}
		if _macie2DeleteFindingsFilter {
			macie2_DeleteFindingsFilter(cfg, client)
			return
		}
		if _macie2DeleteInvitations {
			macie2_DeleteInvitations(cfg, client)
			return
		}
		if _macie2DeleteMember {
			macie2_DeleteMember(cfg, client)
			return
		}
		if _macie2DescribeBuckets {
			macie2_DescribeBuckets(cfg, client)
			return
		}
		if _macie2DescribeClassificationJob {
			macie2_DescribeClassificationJob(cfg, client)
			return
		}
		if _macie2DescribeOrganizationConfiguration {
			macie2_DescribeOrganizationConfiguration(cfg, client)
			return
		}
		if _macie2DisableMacie {
			macie2_DisableMacie(cfg, client)
			return
		}
		if _macie2DisableOrganizationAdminAccount {
			macie2_DisableOrganizationAdminAccount(cfg, client)
			return
		}
		if _macie2DisassociateFromAdministratorAccount {
			macie2_DisassociateFromAdministratorAccount(cfg, client)
			return
		}
		if _macie2DisassociateFromMasterAccount {
			macie2_DisassociateFromMasterAccount(cfg, client)
			return
		}
		if _macie2DisassociateMember {
			macie2_DisassociateMember(cfg, client)
			return
		}
		if _macie2EnableMacie {
			macie2_EnableMacie(cfg, client)
			return
		}
		if _macie2EnableOrganizationAdminAccount {
			macie2_EnableOrganizationAdminAccount(cfg, client)
			return
		}
		if _macie2GetAdministratorAccount {
			macie2_GetAdministratorAccount(cfg, client)
			return
		}
		if _macie2GetAllowList {
			macie2_GetAllowList(cfg, client)
			return
		}
		if _macie2GetAutomatedDiscoveryConfiguration {
			macie2_GetAutomatedDiscoveryConfiguration(cfg, client)
			return
		}
		if _macie2GetBucketStatistics {
			macie2_GetBucketStatistics(cfg, client)
			return
		}
		if _macie2GetClassificationExportConfiguration {
			macie2_GetClassificationExportConfiguration(cfg, client)
			return
		}
		if _macie2GetClassificationScope {
			macie2_GetClassificationScope(cfg, client)
			return
		}
		if _macie2GetCustomDataIdentifier {
			macie2_GetCustomDataIdentifier(cfg, client)
			return
		}
		if _macie2GetFindingStatistics {
			macie2_GetFindingStatistics(cfg, client)
			return
		}
		if _macie2GetFindings {
			macie2_GetFindings(cfg, client)
			return
		}
		if _macie2GetFindingsFilter {
			macie2_GetFindingsFilter(cfg, client)
			return
		}
		if _macie2GetFindingsPublicationConfiguration {
			macie2_GetFindingsPublicationConfiguration(cfg, client)
			return
		}
		if _macie2GetInvitationsCount {
			macie2_GetInvitationsCount(cfg, client)
			return
		}
		if _macie2GetMacieSession {
			macie2_GetMacieSession(cfg, client)
			return
		}
		if _macie2GetMasterAccount {
			macie2_GetMasterAccount(cfg, client)
			return
		}
		if _macie2GetMember {
			macie2_GetMember(cfg, client)
			return
		}
		if _macie2GetResourceProfile {
			macie2_GetResourceProfile(cfg, client)
			return
		}
		if _macie2GetRevealConfiguration {
			macie2_GetRevealConfiguration(cfg, client)
			return
		}
		if _macie2GetSensitiveDataOccurrences {
			macie2_GetSensitiveDataOccurrences(cfg, client)
			return
		}
		if _macie2GetSensitiveDataOccurrencesAvailability {
			macie2_GetSensitiveDataOccurrencesAvailability(cfg, client)
			return
		}
		if _macie2GetSensitivityInspectionTemplate {
			macie2_GetSensitivityInspectionTemplate(cfg, client)
			return
		}
		if _macie2GetUsageStatistics {
			macie2_GetUsageStatistics(cfg, client)
			return
		}
		if _macie2GetUsageTotals {
			macie2_GetUsageTotals(cfg, client)
			return
		}
		if _macie2ListAllowLists {
			macie2_ListAllowLists(cfg, client)
			return
		}
		if _macie2ListAutomatedDiscoveryAccounts {
			macie2_ListAutomatedDiscoveryAccounts(cfg, client)
			return
		}
		if _macie2ListClassificationJobs {
			macie2_ListClassificationJobs(cfg, client)
			return
		}
		if _macie2ListClassificationScopes {
			macie2_ListClassificationScopes(cfg, client)
			return
		}
		if _macie2ListCustomDataIdentifiers {
			macie2_ListCustomDataIdentifiers(cfg, client)
			return
		}
		if _macie2ListFindings {
			macie2_ListFindings(cfg, client)
			return
		}
		if _macie2ListFindingsFilters {
			macie2_ListFindingsFilters(cfg, client)
			return
		}
		if _macie2ListInvitations {
			macie2_ListInvitations(cfg, client)
			return
		}
		if _macie2ListManagedDataIdentifiers {
			macie2_ListManagedDataIdentifiers(cfg, client)
			return
		}
		if _macie2ListMembers {
			macie2_ListMembers(cfg, client)
			return
		}
		if _macie2ListOrganizationAdminAccounts {
			macie2_ListOrganizationAdminAccounts(cfg, client)
			return
		}
		if _macie2ListResourceProfileArtifacts {
			macie2_ListResourceProfileArtifacts(cfg, client)
			return
		}
		if _macie2ListResourceProfileDetections {
			macie2_ListResourceProfileDetections(cfg, client)
			return
		}
		if _macie2ListSensitivityInspectionTemplates {
			macie2_ListSensitivityInspectionTemplates(cfg, client)
			return
		}
		if _macie2ListTagsForResource {
			macie2_ListTagsForResource(cfg, client)
			return
		}
		if _macie2PutClassificationExportConfiguration {
			macie2_PutClassificationExportConfiguration(cfg, client)
			return
		}
		if _macie2PutFindingsPublicationConfiguration {
			macie2_PutFindingsPublicationConfiguration(cfg, client)
			return
		}
		if _macie2SearchResources {
			macie2_SearchResources(cfg, client)
			return
		}
		if _macie2TagResource {
			macie2_TagResource(cfg, client)
			return
		}
		if _macie2TestCustomDataIdentifier {
			macie2_TestCustomDataIdentifier(cfg, client)
			return
		}
		if _macie2UntagResource {
			macie2_UntagResource(cfg, client)
			return
		}
		if _macie2UpdateAllowList {
			macie2_UpdateAllowList(cfg, client)
			return
		}
		if _macie2UpdateAutomatedDiscoveryConfiguration {
			macie2_UpdateAutomatedDiscoveryConfiguration(cfg, client)
			return
		}
		if _macie2UpdateClassificationJob {
			macie2_UpdateClassificationJob(cfg, client)
			return
		}
		if _macie2UpdateClassificationScope {
			macie2_UpdateClassificationScope(cfg, client)
			return
		}
		if _macie2UpdateFindingsFilter {
			macie2_UpdateFindingsFilter(cfg, client)
			return
		}
		if _macie2UpdateMacieSession {
			macie2_UpdateMacieSession(cfg, client)
			return
		}
		if _macie2UpdateMemberSession {
			macie2_UpdateMemberSession(cfg, client)
			return
		}
		if _macie2UpdateOrganizationConfiguration {
			macie2_UpdateOrganizationConfiguration(cfg, client)
			return
		}
		if _macie2UpdateResourceProfile {
			macie2_UpdateResourceProfile(cfg, client)
			return
		}
		if _macie2UpdateResourceProfileDetections {
			macie2_UpdateResourceProfileDetections(cfg, client)
			return
		}
		if _macie2UpdateRevealConfiguration {
			macie2_UpdateRevealConfiguration(cfg, client)
			return
		}
		if _macie2UpdateSensitivityInspectionTemplate {
			macie2_UpdateSensitivityInspectionTemplate(cfg, client)
			return
		}

	},
}

var (
	_macie2AcceptInvitation                        bool
	_macie2BatchGetCustomDataIdentifiers           bool
	_macie2BatchUpdateAutomatedDiscoveryAccounts   bool
	_macie2CreateAllowList                         bool
	_macie2CreateClassificationJob                 bool
	_macie2CreateCustomDataIdentifier              bool
	_macie2CreateFindingsFilter                    bool
	_macie2CreateInvitations                       bool
	_macie2CreateMember                            bool
	_macie2CreateSampleFindings                    bool
	_macie2DeclineInvitations                      bool
	_macie2DeleteAllowList                         bool
	_macie2DeleteCustomDataIdentifier              bool
	_macie2DeleteFindingsFilter                    bool
	_macie2DeleteInvitations                       bool
	_macie2DeleteMember                            bool
	_macie2DescribeBuckets                         bool
	_macie2DescribeClassificationJob               bool
	_macie2DescribeOrganizationConfiguration       bool
	_macie2DisableMacie                            bool
	_macie2DisableOrganizationAdminAccount         bool
	_macie2DisassociateFromAdministratorAccount    bool
	_macie2DisassociateFromMasterAccount           bool
	_macie2DisassociateMember                      bool
	_macie2EnableMacie                             bool
	_macie2EnableOrganizationAdminAccount          bool
	_macie2GetAdministratorAccount                 bool
	_macie2GetAllowList                            bool
	_macie2GetAutomatedDiscoveryConfiguration      bool
	_macie2GetBucketStatistics                     bool
	_macie2GetClassificationExportConfiguration    bool
	_macie2GetClassificationScope                  bool
	_macie2GetCustomDataIdentifier                 bool
	_macie2GetFindingStatistics                    bool
	_macie2GetFindings                             bool
	_macie2GetFindingsFilter                       bool
	_macie2GetFindingsPublicationConfiguration     bool
	_macie2GetInvitationsCount                     bool
	_macie2GetMacieSession                         bool
	_macie2GetMasterAccount                        bool
	_macie2GetMember                               bool
	_macie2GetResourceProfile                      bool
	_macie2GetRevealConfiguration                  bool
	_macie2GetSensitiveDataOccurrences             bool
	_macie2GetSensitiveDataOccurrencesAvailability bool
	_macie2GetSensitivityInspectionTemplate        bool
	_macie2GetUsageStatistics                      bool
	_macie2GetUsageTotals                          bool
	_macie2ListAllowLists                          bool
	_macie2ListAutomatedDiscoveryAccounts          bool
	_macie2ListClassificationJobs                  bool
	_macie2ListClassificationScopes                bool
	_macie2ListCustomDataIdentifiers               bool
	_macie2ListFindings                            bool
	_macie2ListFindingsFilters                     bool
	_macie2ListInvitations                         bool
	_macie2ListManagedDataIdentifiers              bool
	_macie2ListMembers                             bool
	_macie2ListOrganizationAdminAccounts           bool
	_macie2ListResourceProfileArtifacts            bool
	_macie2ListResourceProfileDetections           bool
	_macie2ListSensitivityInspectionTemplates      bool
	_macie2ListTagsForResource                     bool
	_macie2PutClassificationExportConfiguration    bool
	_macie2PutFindingsPublicationConfiguration     bool
	_macie2SearchResources                         bool
	_macie2TagResource                             bool
	_macie2TestCustomDataIdentifier                bool
	_macie2UntagResource                           bool
	_macie2UpdateAllowList                         bool
	_macie2UpdateAutomatedDiscoveryConfiguration   bool
	_macie2UpdateClassificationJob                 bool
	_macie2UpdateClassificationScope               bool
	_macie2UpdateFindingsFilter                    bool
	_macie2UpdateMacieSession                      bool
	_macie2UpdateMemberSession                     bool
	_macie2UpdateOrganizationConfiguration         bool
	_macie2UpdateResourceProfile                   bool
	_macie2UpdateResourceProfileDetections         bool
	_macie2UpdateRevealConfiguration               bool
	_macie2UpdateSensitivityInspectionTemplate     bool

	_macie2Account                       string
	_macie2AccountId                     string
	_macie2AccountIds                    []string
	_macie2Accounts                      string
	_macie2Action                        string
	_macie2AdminAccountId                string
	_macie2AdministratorAccountId        string
	_macie2AllowListIds                  []string
	_macie2AutoEnable                    string
	_macie2AutoEnableOrganizationMembers string
	_macie2BucketCriteria                string
	_macie2ClientToken                   string
	_macie2Configuration                 string
	_macie2Criteria                      string
	_macie2CustomDataIdentifierIds       []string
	_macie2Description                   string
	_macie2DisableEmailNotification      string
	_macie2Excludes                      string
	_macie2FilterBy                      string
	_macie2FilterCriteria                string
	_macie2FindingCriteria               string
	_macie2FindingId                     string
	_macie2FindingIds                    []string
	_macie2FindingPublishingFrequency    string
	_macie2FindingTypes                  string
	_macie2GroupBy                       string
	_macie2Id                            string
	_macie2Ids                           []string
	_macie2IgnoreJobChecks               string
	_macie2IgnoreWords                   []string
	_macie2Includes                      string
	_macie2InitialRun                    string
	_macie2InvitationId                  string
	_macie2JobId                         string
	_macie2JobStatus                     string
	_macie2JobType                       string
	_macie2Keywords                      []string
	_macie2ManagedDataIdentifierIds      []string
	_macie2ManagedDataIdentifierSelector string
	_macie2MasterAccount                 string
	_macie2MaxResults                    string
	_macie2MaximumMatchDistance          string
	_macie2Message                       string
	_macie2Name                          string
	_macie2NextToken                     string
	_macie2OnlyAssociated                string
	_macie2Position                      string
	_macie2Regex                         string
	_macie2ResourceArn                   string
	_macie2RetrievalConfiguration        string
	_macie2S3                            string
	_macie2S3JobDefinition               string
	_macie2SampleText                    string
	_macie2SamplingPercentage            string
	_macie2ScheduleFrequency             string
	_macie2SecurityHubConfiguration      string
	_macie2SensitivityScoreOverride      string
	_macie2SeverityLevels                string
	_macie2Size                          string
	_macie2SortBy                        string
	_macie2SortCriteria                  string
	_macie2Status                        string
	_macie2SuppressDataIdentifiers       string
	_macie2TagKeys                       []string
	_macie2Tags                          string
	_macie2TimeRange                     string
)

// Accepts an Amazon Macie membership invitation that was received from a specific
// account.
func macie2_AcceptInvitation(cfg aws.Config, client *macie2.Client) {
	input := &macie2.AcceptInvitationInput{
		// InvitationId: *string, // Required
	}

	if len(_macie2InvitationId) > 0 {
		input.InvitationId = aws.String(_macie2InvitationId)
	}
	if len(_macie2AdministratorAccountId) > 0 {
		input.AdministratorAccountId = aws.String(_macie2AdministratorAccountId)
	}
	if len(_macie2MasterAccount) > 0 {
		input.MasterAccount = aws.String(_macie2MasterAccount)
	}

	if resp, err := client.AcceptInvitation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about one or more custom data identifiers.
func macie2_BatchGetCustomDataIdentifiers(cfg aws.Config, client *macie2.Client) {
	input := &macie2.BatchGetCustomDataIdentifiersInput{}

	if len(_macie2Ids) > 0 {
		input.Ids = append([]string(nil), _macie2Ids...)
	}

	if resp, err := client.BatchGetCustomDataIdentifiers(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Changes the status of automated sensitive data discovery for one or more
// accounts.
func macie2_BatchUpdateAutomatedDiscoveryAccounts(cfg aws.Config, client *macie2.Client) {
	input := &macie2.BatchUpdateAutomatedDiscoveryAccountsInput{}

	if len(_macie2Accounts) > 0 {
		if err := assignInputField(input, "Accounts", _macie2Accounts); err != nil {
			log.Errorf("invalid --accounts: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchUpdateAutomatedDiscoveryAccounts(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates and defines the settings for an allow list.
func macie2_CreateAllowList(cfg aws.Config, client *macie2.Client) {
	input := &macie2.CreateAllowListInput{
		// ClientToken: *string, // Required
		// Criteria: *types.AllowListCriteria, // Required
		// Name: *string, // Required
	}

	if len(_macie2ClientToken) > 0 {
		input.ClientToken = aws.String(_macie2ClientToken)
	}
	if len(_macie2Criteria) > 0 {
		if err := assignInputField(input, "Criteria", _macie2Criteria); err != nil {
			log.Errorf("invalid --criteria: %s", err.Error())
			return
		}
	}
	if len(_macie2Name) > 0 {
		input.Name = aws.String(_macie2Name)
	}
	if len(_macie2Description) > 0 {
		input.Description = aws.String(_macie2Description)
	}
	if len(_macie2Tags) > 0 {
		if err := assignInputField(input, "Tags", _macie2Tags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAllowList(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates and defines the settings for a classification job.
func macie2_CreateClassificationJob(cfg aws.Config, client *macie2.Client) {
	input := &macie2.CreateClassificationJobInput{
		// ClientToken: *string, // Required
		// JobType: types.JobType, // Required
		// Name: *string, // Required
		// S3JobDefinition: *types.S3JobDefinition, // Required
	}

	if len(_macie2ClientToken) > 0 {
		input.ClientToken = aws.String(_macie2ClientToken)
	}
	if len(_macie2JobType) > 0 {
		if err := assignInputField(input, "JobType", _macie2JobType); err != nil {
			log.Errorf("invalid --job-type: %s", err.Error())
			return
		}
	}
	if len(_macie2Name) > 0 {
		input.Name = aws.String(_macie2Name)
	}
	if len(_macie2S3JobDefinition) > 0 {
		if err := assignInputField(input, "S3JobDefinition", _macie2S3JobDefinition); err != nil {
			log.Errorf("invalid --s3-job-definition: %s", err.Error())
			return
		}
	}
	if len(_macie2AllowListIds) > 0 {
		input.AllowListIds = append([]string(nil), _macie2AllowListIds...)
	}
	if len(_macie2CustomDataIdentifierIds) > 0 {
		input.CustomDataIdentifierIds = append([]string(nil), _macie2CustomDataIdentifierIds...)
	}
	if len(_macie2Description) > 0 {
		input.Description = aws.String(_macie2Description)
	}
	if len(_macie2InitialRun) > 0 {
		if err := assignInputField(input, "InitialRun", _macie2InitialRun); err != nil {
			log.Errorf("invalid --initial-run: %s", err.Error())
			return
		}
	}
	if len(_macie2ManagedDataIdentifierIds) > 0 {
		input.ManagedDataIdentifierIds = append([]string(nil), _macie2ManagedDataIdentifierIds...)
	}
	if len(_macie2ManagedDataIdentifierSelector) > 0 {
		if err := assignInputField(input, "ManagedDataIdentifierSelector", _macie2ManagedDataIdentifierSelector); err != nil {
			log.Errorf("invalid --managed-data-identifier-selector: %s", err.Error())
			return
		}
	}
	if len(_macie2SamplingPercentage) > 0 {
		if err := assignInputField(input, "SamplingPercentage", _macie2SamplingPercentage); err != nil {
			log.Errorf("invalid --sampling-percentage: %s", err.Error())
			return
		}
	}
	if len(_macie2ScheduleFrequency) > 0 {
		if err := assignInputField(input, "ScheduleFrequency", _macie2ScheduleFrequency); err != nil {
			log.Errorf("invalid --schedule-frequency: %s", err.Error())
			return
		}
	}
	if len(_macie2Tags) > 0 {
		if err := assignInputField(input, "Tags", _macie2Tags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateClassificationJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates and defines the criteria and other settings for a custom data
// identifier.
func macie2_CreateCustomDataIdentifier(cfg aws.Config, client *macie2.Client) {
	input := &macie2.CreateCustomDataIdentifierInput{
		// Name: *string, // Required
		// Regex: *string, // Required
	}

	if len(_macie2Name) > 0 {
		input.Name = aws.String(_macie2Name)
	}
	if len(_macie2Regex) > 0 {
		input.Regex = aws.String(_macie2Regex)
	}
	if len(_macie2ClientToken) > 0 {
		input.ClientToken = aws.String(_macie2ClientToken)
	}
	if len(_macie2Description) > 0 {
		input.Description = aws.String(_macie2Description)
	}
	if len(_macie2IgnoreWords) > 0 {
		input.IgnoreWords = append([]string(nil), _macie2IgnoreWords...)
	}
	if len(_macie2Keywords) > 0 {
		input.Keywords = append([]string(nil), _macie2Keywords...)
	}
	if len(_macie2MaximumMatchDistance) > 0 {
		if err := assignInputField(input, "MaximumMatchDistance", _macie2MaximumMatchDistance); err != nil {
			log.Errorf("invalid --maximum-match-distance: %s", err.Error())
			return
		}
	}
	if len(_macie2SeverityLevels) > 0 {
		if err := assignInputField(input, "SeverityLevels", _macie2SeverityLevels); err != nil {
			log.Errorf("invalid --severity-levels: %s", err.Error())
			return
		}
	}
	if len(_macie2Tags) > 0 {
		if err := assignInputField(input, "Tags", _macie2Tags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCustomDataIdentifier(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates and defines the criteria and other settings for a findings filter.
func macie2_CreateFindingsFilter(cfg aws.Config, client *macie2.Client) {
	input := &macie2.CreateFindingsFilterInput{
		// Action: types.FindingsFilterAction, // Required
		// FindingCriteria: *types.FindingCriteria, // Required
		// Name: *string, // Required
	}

	if len(_macie2Action) > 0 {
		if err := assignInputField(input, "Action", _macie2Action); err != nil {
			log.Errorf("invalid --action: %s", err.Error())
			return
		}
	}
	if len(_macie2FindingCriteria) > 0 {
		if err := assignInputField(input, "FindingCriteria", _macie2FindingCriteria); err != nil {
			log.Errorf("invalid --finding-criteria: %s", err.Error())
			return
		}
	}
	if len(_macie2Name) > 0 {
		input.Name = aws.String(_macie2Name)
	}
	if len(_macie2ClientToken) > 0 {
		input.ClientToken = aws.String(_macie2ClientToken)
	}
	if len(_macie2Description) > 0 {
		input.Description = aws.String(_macie2Description)
	}
	if len(_macie2Position) > 0 {
		if err := assignInputField(input, "Position", _macie2Position); err != nil {
			log.Errorf("invalid --position: %s", err.Error())
			return
		}
	}
	if len(_macie2Tags) > 0 {
		if err := assignInputField(input, "Tags", _macie2Tags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateFindingsFilter(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sends an Amazon Macie membership invitation to one or more accounts.
func macie2_CreateInvitations(cfg aws.Config, client *macie2.Client) {
	input := &macie2.CreateInvitationsInput{
		// AccountIds: []string, // Required
	}

	if len(_macie2AccountIds) > 0 {
		input.AccountIds = append([]string(nil), _macie2AccountIds...)
	}
	if len(_macie2DisableEmailNotification) > 0 {
		if err := assignInputField(input, "DisableEmailNotification", _macie2DisableEmailNotification); err != nil {
			log.Errorf("invalid --disable-email-notification: %s", err.Error())
			return
		}
	}
	if len(_macie2Message) > 0 {
		input.Message = aws.String(_macie2Message)
	}

	if resp, err := client.CreateInvitations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates an account with an Amazon Macie administrator account.
func macie2_CreateMember(cfg aws.Config, client *macie2.Client) {
	input := &macie2.CreateMemberInput{
		// Account: *types.AccountDetail, // Required
	}

	if len(_macie2Account) > 0 {
		if err := assignInputField(input, "Account", _macie2Account); err != nil {
			log.Errorf("invalid --account: %s", err.Error())
			return
		}
	}
	if len(_macie2Tags) > 0 {
		if err := assignInputField(input, "Tags", _macie2Tags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateMember(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates sample findings.
func macie2_CreateSampleFindings(cfg aws.Config, client *macie2.Client) {
	input := &macie2.CreateSampleFindingsInput{}

	if len(_macie2FindingTypes) > 0 {
		if err := assignInputField(input, "FindingTypes", _macie2FindingTypes); err != nil {
			log.Errorf("invalid --finding-types: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSampleFindings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Declines Amazon Macie membership invitations that were received from specific
// accounts.
func macie2_DeclineInvitations(cfg aws.Config, client *macie2.Client) {
	input := &macie2.DeclineInvitationsInput{
		// AccountIds: []string, // Required
	}

	if len(_macie2AccountIds) > 0 {
		input.AccountIds = append([]string(nil), _macie2AccountIds...)
	}

	if resp, err := client.DeclineInvitations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an allow list.
func macie2_DeleteAllowList(cfg aws.Config, client *macie2.Client) {
	input := &macie2.DeleteAllowListInput{
		// Id: *string, // Required
	}

	if len(_macie2Id) > 0 {
		input.Id = aws.String(_macie2Id)
	}
	if len(_macie2IgnoreJobChecks) > 0 {
		input.IgnoreJobChecks = aws.String(_macie2IgnoreJobChecks)
	}

	if resp, err := client.DeleteAllowList(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Soft deletes a custom data identifier.
func macie2_DeleteCustomDataIdentifier(cfg aws.Config, client *macie2.Client) {
	input := &macie2.DeleteCustomDataIdentifierInput{
		// Id: *string, // Required
	}

	if len(_macie2Id) > 0 {
		input.Id = aws.String(_macie2Id)
	}

	if resp, err := client.DeleteCustomDataIdentifier(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a findings filter.
func macie2_DeleteFindingsFilter(cfg aws.Config, client *macie2.Client) {
	input := &macie2.DeleteFindingsFilterInput{
		// Id: *string, // Required
	}

	if len(_macie2Id) > 0 {
		input.Id = aws.String(_macie2Id)
	}

	if resp, err := client.DeleteFindingsFilter(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes Amazon Macie membership invitations that were received from specific
// accounts.
func macie2_DeleteInvitations(cfg aws.Config, client *macie2.Client) {
	input := &macie2.DeleteInvitationsInput{
		// AccountIds: []string, // Required
	}

	if len(_macie2AccountIds) > 0 {
		input.AccountIds = append([]string(nil), _macie2AccountIds...)
	}

	if resp, err := client.DeleteInvitations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the association between an Amazon Macie administrator account and an
// account.
func macie2_DeleteMember(cfg aws.Config, client *macie2.Client) {
	input := &macie2.DeleteMemberInput{
		// Id: *string, // Required
	}

	if len(_macie2Id) > 0 {
		input.Id = aws.String(_macie2Id)
	}

	if resp, err := client.DeleteMember(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves (queries) statistical data and other information about one or more S3
// buckets that Amazon Macie monitors and analyzes for an account.
func macie2_DescribeBuckets(cfg aws.Config, client *macie2.Client) {
	input := &macie2.DescribeBucketsInput{}

	if len(_macie2Criteria) > 0 {
		if err := assignInputField(input, "Criteria", _macie2Criteria); err != nil {
			log.Errorf("invalid --criteria: %s", err.Error())
			return
		}
	}
	if len(_macie2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _macie2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_macie2NextToken) > 0 {
		input.NextToken = aws.String(_macie2NextToken)
	}
	if len(_macie2SortCriteria) > 0 {
		if err := assignInputField(input, "SortCriteria", _macie2SortCriteria); err != nil {
			log.Errorf("invalid --sort-criteria: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeBuckets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*macie2.DescribeBucketsOutput
	p := macie2.NewDescribeBucketsPaginator(client, input)
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

// Retrieves the status and settings for a classification job.
func macie2_DescribeClassificationJob(cfg aws.Config, client *macie2.Client) {
	input := &macie2.DescribeClassificationJobInput{
		// JobId: *string, // Required
	}

	if len(_macie2JobId) > 0 {
		input.JobId = aws.String(_macie2JobId)
	}

	if resp, err := client.DescribeClassificationJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the Amazon Macie configuration settings for an organization in
// Organizations.
func macie2_DescribeOrganizationConfiguration(cfg aws.Config, client *macie2.Client) {
	input := &macie2.DescribeOrganizationConfigurationInput{}

	if resp, err := client.DescribeOrganizationConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disables Amazon Macie and deletes all settings and resources for a Macie
// account.
func macie2_DisableMacie(cfg aws.Config, client *macie2.Client) {
	input := &macie2.DisableMacieInput{}

	if resp, err := client.DisableMacie(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disables an account as the delegated Amazon Macie administrator account for an
// organization in Organizations.
func macie2_DisableOrganizationAdminAccount(cfg aws.Config, client *macie2.Client) {
	input := &macie2.DisableOrganizationAdminAccountInput{
		// AdminAccountId: *string, // Required
	}

	if len(_macie2AdminAccountId) > 0 {
		input.AdminAccountId = aws.String(_macie2AdminAccountId)
	}

	if resp, err := client.DisableOrganizationAdminAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates a member account from its Amazon Macie administrator account.
func macie2_DisassociateFromAdministratorAccount(cfg aws.Config, client *macie2.Client) {
	input := &macie2.DisassociateFromAdministratorAccountInput{}

	if resp, err := client.DisassociateFromAdministratorAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// (Deprecated) Disassociates a member account from its Amazon Macie administrator
// account. This operation has been replaced by the
// DisassociateFromAdministratorAccount operation.
func macie2_DisassociateFromMasterAccount(cfg aws.Config, client *macie2.Client) {
	input := &macie2.DisassociateFromMasterAccountInput{}

	if resp, err := client.DisassociateFromMasterAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates an Amazon Macie administrator account from a member account.
func macie2_DisassociateMember(cfg aws.Config, client *macie2.Client) {
	input := &macie2.DisassociateMemberInput{
		// Id: *string, // Required
	}

	if len(_macie2Id) > 0 {
		input.Id = aws.String(_macie2Id)
	}

	if resp, err := client.DisassociateMember(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables Amazon Macie and specifies the configuration settings for a Macie
// account.
func macie2_EnableMacie(cfg aws.Config, client *macie2.Client) {
	input := &macie2.EnableMacieInput{}

	if len(_macie2ClientToken) > 0 {
		input.ClientToken = aws.String(_macie2ClientToken)
	}
	if len(_macie2FindingPublishingFrequency) > 0 {
		if err := assignInputField(input, "FindingPublishingFrequency", _macie2FindingPublishingFrequency); err != nil {
			log.Errorf("invalid --finding-publishing-frequency: %s", err.Error())
			return
		}
	}
	if len(_macie2Status) > 0 {
		if err := assignInputField(input, "Status", _macie2Status); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if resp, err := client.EnableMacie(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Designates an account as the delegated Amazon Macie administrator account for
// an organization in Organizations.
func macie2_EnableOrganizationAdminAccount(cfg aws.Config, client *macie2.Client) {
	input := &macie2.EnableOrganizationAdminAccountInput{
		// AdminAccountId: *string, // Required
	}

	if len(_macie2AdminAccountId) > 0 {
		input.AdminAccountId = aws.String(_macie2AdminAccountId)
	}
	if len(_macie2ClientToken) > 0 {
		input.ClientToken = aws.String(_macie2ClientToken)
	}

	if resp, err := client.EnableOrganizationAdminAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the Amazon Macie administrator account for an
// account.
func macie2_GetAdministratorAccount(cfg aws.Config, client *macie2.Client) {
	input := &macie2.GetAdministratorAccountInput{}

	if resp, err := client.GetAdministratorAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the settings and status of an allow list.
func macie2_GetAllowList(cfg aws.Config, client *macie2.Client) {
	input := &macie2.GetAllowListInput{
		// Id: *string, // Required
	}

	if len(_macie2Id) > 0 {
		input.Id = aws.String(_macie2Id)
	}

	if resp, err := client.GetAllowList(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the configuration settings and status of automated sensitive data
// discovery for an organization or standalone account.
func macie2_GetAutomatedDiscoveryConfiguration(cfg aws.Config, client *macie2.Client) {
	input := &macie2.GetAutomatedDiscoveryConfigurationInput{}

	if resp, err := client.GetAutomatedDiscoveryConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves (queries) aggregated statistical data about all the S3 buckets that
// Amazon Macie monitors and analyzes for an account.
func macie2_GetBucketStatistics(cfg aws.Config, client *macie2.Client) {
	input := &macie2.GetBucketStatisticsInput{}

	if len(_macie2AccountId) > 0 {
		input.AccountId = aws.String(_macie2AccountId)
	}

	if resp, err := client.GetBucketStatistics(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the configuration settings for storing data classification results.
func macie2_GetClassificationExportConfiguration(cfg aws.Config, client *macie2.Client) {
	input := &macie2.GetClassificationExportConfigurationInput{}

	if resp, err := client.GetClassificationExportConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the classification scope settings for an account.
func macie2_GetClassificationScope(cfg aws.Config, client *macie2.Client) {
	input := &macie2.GetClassificationScopeInput{
		// Id: *string, // Required
	}

	if len(_macie2Id) > 0 {
		input.Id = aws.String(_macie2Id)
	}

	if resp, err := client.GetClassificationScope(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the criteria and other settings for a custom data identifier.
func macie2_GetCustomDataIdentifier(cfg aws.Config, client *macie2.Client) {
	input := &macie2.GetCustomDataIdentifierInput{
		// Id: *string, // Required
	}

	if len(_macie2Id) > 0 {
		input.Id = aws.String(_macie2Id)
	}

	if resp, err := client.GetCustomDataIdentifier(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves (queries) aggregated statistical data about findings.
func macie2_GetFindingStatistics(cfg aws.Config, client *macie2.Client) {
	input := &macie2.GetFindingStatisticsInput{
		// GroupBy: types.GroupBy, // Required
	}

	if len(_macie2GroupBy) > 0 {
		if err := assignInputField(input, "GroupBy", _macie2GroupBy); err != nil {
			log.Errorf("invalid --group-by: %s", err.Error())
			return
		}
	}
	if len(_macie2FindingCriteria) > 0 {
		if err := assignInputField(input, "FindingCriteria", _macie2FindingCriteria); err != nil {
			log.Errorf("invalid --finding-criteria: %s", err.Error())
			return
		}
	}
	if len(_macie2Size) > 0 {
		if err := assignInputField(input, "Size", _macie2Size); err != nil {
			log.Errorf("invalid --size: %s", err.Error())
			return
		}
	}
	if len(_macie2SortCriteria) > 0 {
		if err := assignInputField(input, "SortCriteria", _macie2SortCriteria); err != nil {
			log.Errorf("invalid --sort-criteria: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetFindingStatistics(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the details of one or more findings.
func macie2_GetFindings(cfg aws.Config, client *macie2.Client) {
	input := &macie2.GetFindingsInput{
		// FindingIds: []string, // Required
	}

	if len(_macie2FindingIds) > 0 {
		input.FindingIds = append([]string(nil), _macie2FindingIds...)
	}
	if len(_macie2SortCriteria) > 0 {
		if err := assignInputField(input, "SortCriteria", _macie2SortCriteria); err != nil {
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

// Retrieves the criteria and other settings for a findings filter.
func macie2_GetFindingsFilter(cfg aws.Config, client *macie2.Client) {
	input := &macie2.GetFindingsFilterInput{
		// Id: *string, // Required
	}

	if len(_macie2Id) > 0 {
		input.Id = aws.String(_macie2Id)
	}

	if resp, err := client.GetFindingsFilter(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the configuration settings for publishing findings to Security Hub.
func macie2_GetFindingsPublicationConfiguration(cfg aws.Config, client *macie2.Client) {
	input := &macie2.GetFindingsPublicationConfigurationInput{}

	if resp, err := client.GetFindingsPublicationConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the count of Amazon Macie membership invitations that were received
// by an account.
func macie2_GetInvitationsCount(cfg aws.Config, client *macie2.Client) {
	input := &macie2.GetInvitationsCountInput{}

	if resp, err := client.GetInvitationsCount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the status and configuration settings for an Amazon Macie account.
func macie2_GetMacieSession(cfg aws.Config, client *macie2.Client) {
	input := &macie2.GetMacieSessionInput{}

	if resp, err := client.GetMacieSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// (Deprecated) Retrieves information about the Amazon Macie administrator account
// for an account. This operation has been replaced by the GetAdministratorAccount
// operation.
func macie2_GetMasterAccount(cfg aws.Config, client *macie2.Client) {
	input := &macie2.GetMasterAccountInput{}

	if resp, err := client.GetMasterAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about an account that's associated with an Amazon Macie
// administrator account.
func macie2_GetMember(cfg aws.Config, client *macie2.Client) {
	input := &macie2.GetMemberInput{
		// Id: *string, // Required
	}

	if len(_macie2Id) > 0 {
		input.Id = aws.String(_macie2Id)
	}

	if resp, err := client.GetMember(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves (queries) sensitive data discovery statistics and the sensitivity
// score for an S3 bucket.
func macie2_GetResourceProfile(cfg aws.Config, client *macie2.Client) {
	input := &macie2.GetResourceProfileInput{
		// ResourceArn: *string, // Required
	}

	if len(_macie2ResourceArn) > 0 {
		input.ResourceArn = aws.String(_macie2ResourceArn)
	}

	if resp, err := client.GetResourceProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the status and configuration settings for retrieving occurrences of
// sensitive data reported by findings.
func macie2_GetRevealConfiguration(cfg aws.Config, client *macie2.Client) {
	input := &macie2.GetRevealConfigurationInput{}

	if resp, err := client.GetRevealConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves occurrences of sensitive data reported by a finding.
func macie2_GetSensitiveDataOccurrences(cfg aws.Config, client *macie2.Client) {
	input := &macie2.GetSensitiveDataOccurrencesInput{
		// FindingId: *string, // Required
	}

	if len(_macie2FindingId) > 0 {
		input.FindingId = aws.String(_macie2FindingId)
	}

	if resp, err := client.GetSensitiveDataOccurrences(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Checks whether occurrences of sensitive data can be retrieved for a finding.
func macie2_GetSensitiveDataOccurrencesAvailability(cfg aws.Config, client *macie2.Client) {
	input := &macie2.GetSensitiveDataOccurrencesAvailabilityInput{
		// FindingId: *string, // Required
	}

	if len(_macie2FindingId) > 0 {
		input.FindingId = aws.String(_macie2FindingId)
	}

	if resp, err := client.GetSensitiveDataOccurrencesAvailability(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the settings for the sensitivity inspection template for an account.
func macie2_GetSensitivityInspectionTemplate(cfg aws.Config, client *macie2.Client) {
	input := &macie2.GetSensitivityInspectionTemplateInput{
		// Id: *string, // Required
	}

	if len(_macie2Id) > 0 {
		input.Id = aws.String(_macie2Id)
	}

	if resp, err := client.GetSensitivityInspectionTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves (queries) quotas and aggregated usage data for one or more accounts.
func macie2_GetUsageStatistics(cfg aws.Config, client *macie2.Client) {
	input := &macie2.GetUsageStatisticsInput{}

	if len(_macie2FilterBy) > 0 {
		if err := assignInputField(input, "FilterBy", _macie2FilterBy); err != nil {
			log.Errorf("invalid --filter-by: %s", err.Error())
			return
		}
	}
	if len(_macie2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _macie2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_macie2NextToken) > 0 {
		input.NextToken = aws.String(_macie2NextToken)
	}
	if len(_macie2SortBy) > 0 {
		if err := assignInputField(input, "SortBy", _macie2SortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_macie2TimeRange) > 0 {
		if err := assignInputField(input, "TimeRange", _macie2TimeRange); err != nil {
			log.Errorf("invalid --time-range: %s", err.Error())
			return
		}
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

	var results []*macie2.GetUsageStatisticsOutput
	p := macie2.NewGetUsageStatisticsPaginator(client, input)
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

// Retrieves (queries) aggregated usage data for an account.
func macie2_GetUsageTotals(cfg aws.Config, client *macie2.Client) {
	input := &macie2.GetUsageTotalsInput{}

	if len(_macie2TimeRange) > 0 {
		input.TimeRange = aws.String(_macie2TimeRange)
	}

	if resp, err := client.GetUsageTotals(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a subset of information about all the allow lists for an account.
func macie2_ListAllowLists(cfg aws.Config, client *macie2.Client) {
	input := &macie2.ListAllowListsInput{}

	if len(_macie2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _macie2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_macie2NextToken) > 0 {
		input.NextToken = aws.String(_macie2NextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAllowLists(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*macie2.ListAllowListsOutput
	p := macie2.NewListAllowListsPaginator(client, input)
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

// Retrieves the status of automated sensitive data discovery for one or more
// accounts.
func macie2_ListAutomatedDiscoveryAccounts(cfg aws.Config, client *macie2.Client) {
	input := &macie2.ListAutomatedDiscoveryAccountsInput{}

	if len(_macie2AccountIds) > 0 {
		input.AccountIds = append([]string(nil), _macie2AccountIds...)
	}
	if len(_macie2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _macie2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_macie2NextToken) > 0 {
		input.NextToken = aws.String(_macie2NextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAutomatedDiscoveryAccounts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*macie2.ListAutomatedDiscoveryAccountsOutput
	p := macie2.NewListAutomatedDiscoveryAccountsPaginator(client, input)
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

// Retrieves a subset of information about one or more classification jobs.
func macie2_ListClassificationJobs(cfg aws.Config, client *macie2.Client) {
	input := &macie2.ListClassificationJobsInput{}

	if len(_macie2FilterCriteria) > 0 {
		if err := assignInputField(input, "FilterCriteria", _macie2FilterCriteria); err != nil {
			log.Errorf("invalid --filter-criteria: %s", err.Error())
			return
		}
	}
	if len(_macie2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _macie2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_macie2NextToken) > 0 {
		input.NextToken = aws.String(_macie2NextToken)
	}
	if len(_macie2SortCriteria) > 0 {
		if err := assignInputField(input, "SortCriteria", _macie2SortCriteria); err != nil {
			log.Errorf("invalid --sort-criteria: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListClassificationJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*macie2.ListClassificationJobsOutput
	p := macie2.NewListClassificationJobsPaginator(client, input)
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

// Retrieves a subset of information about the classification scope for an account.
func macie2_ListClassificationScopes(cfg aws.Config, client *macie2.Client) {
	input := &macie2.ListClassificationScopesInput{}

	if len(_macie2Name) > 0 {
		input.Name = aws.String(_macie2Name)
	}
	if len(_macie2NextToken) > 0 {
		input.NextToken = aws.String(_macie2NextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListClassificationScopes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*macie2.ListClassificationScopesOutput
	p := macie2.NewListClassificationScopesPaginator(client, input)
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

// Retrieves a subset of information about the custom data identifiers for an
// account.
func macie2_ListCustomDataIdentifiers(cfg aws.Config, client *macie2.Client) {
	input := &macie2.ListCustomDataIdentifiersInput{}

	if len(_macie2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _macie2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_macie2NextToken) > 0 {
		input.NextToken = aws.String(_macie2NextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCustomDataIdentifiers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*macie2.ListCustomDataIdentifiersOutput
	p := macie2.NewListCustomDataIdentifiersPaginator(client, input)
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

// Retrieves a subset of information about one or more findings.
func macie2_ListFindings(cfg aws.Config, client *macie2.Client) {
	input := &macie2.ListFindingsInput{}

	if len(_macie2FindingCriteria) > 0 {
		if err := assignInputField(input, "FindingCriteria", _macie2FindingCriteria); err != nil {
			log.Errorf("invalid --finding-criteria: %s", err.Error())
			return
		}
	}
	if len(_macie2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _macie2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_macie2NextToken) > 0 {
		input.NextToken = aws.String(_macie2NextToken)
	}
	if len(_macie2SortCriteria) > 0 {
		if err := assignInputField(input, "SortCriteria", _macie2SortCriteria); err != nil {
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

	var results []*macie2.ListFindingsOutput
	p := macie2.NewListFindingsPaginator(client, input)
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

// Retrieves a subset of information about all the findings filters for an account.
func macie2_ListFindingsFilters(cfg aws.Config, client *macie2.Client) {
	input := &macie2.ListFindingsFiltersInput{}

	if len(_macie2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _macie2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_macie2NextToken) > 0 {
		input.NextToken = aws.String(_macie2NextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListFindingsFilters(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*macie2.ListFindingsFiltersOutput
	p := macie2.NewListFindingsFiltersPaginator(client, input)
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

// Retrieves information about Amazon Macie membership invitations that were
// received by an account.
func macie2_ListInvitations(cfg aws.Config, client *macie2.Client) {
	input := &macie2.ListInvitationsInput{}

	if len(_macie2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _macie2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_macie2NextToken) > 0 {
		input.NextToken = aws.String(_macie2NextToken)
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

	var results []*macie2.ListInvitationsOutput
	p := macie2.NewListInvitationsPaginator(client, input)
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

// Retrieves information about all the managed data identifiers that Amazon Macie
// currently provides.
func macie2_ListManagedDataIdentifiers(cfg aws.Config, client *macie2.Client) {
	input := &macie2.ListManagedDataIdentifiersInput{}

	if len(_macie2NextToken) > 0 {
		input.NextToken = aws.String(_macie2NextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListManagedDataIdentifiers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*macie2.ListManagedDataIdentifiersOutput
	p := macie2.NewListManagedDataIdentifiersPaginator(client, input)
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

// Retrieves information about the accounts that are associated with an Amazon
// Macie administrator account.
func macie2_ListMembers(cfg aws.Config, client *macie2.Client) {
	input := &macie2.ListMembersInput{}

	if len(_macie2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _macie2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_macie2NextToken) > 0 {
		input.NextToken = aws.String(_macie2NextToken)
	}
	if len(_macie2OnlyAssociated) > 0 {
		input.OnlyAssociated = aws.String(_macie2OnlyAssociated)
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

	var results []*macie2.ListMembersOutput
	p := macie2.NewListMembersPaginator(client, input)
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

// Retrieves information about the delegated Amazon Macie administrator account
// for an organization in Organizations.
func macie2_ListOrganizationAdminAccounts(cfg aws.Config, client *macie2.Client) {
	input := &macie2.ListOrganizationAdminAccountsInput{}

	if len(_macie2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _macie2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_macie2NextToken) > 0 {
		input.NextToken = aws.String(_macie2NextToken)
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

	var results []*macie2.ListOrganizationAdminAccountsOutput
	p := macie2.NewListOrganizationAdminAccountsPaginator(client, input)
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

// Retrieves information about objects that Amazon Macie selected from an S3
// bucket for automated sensitive data discovery.
func macie2_ListResourceProfileArtifacts(cfg aws.Config, client *macie2.Client) {
	input := &macie2.ListResourceProfileArtifactsInput{
		// ResourceArn: *string, // Required
	}

	if len(_macie2ResourceArn) > 0 {
		input.ResourceArn = aws.String(_macie2ResourceArn)
	}
	if len(_macie2NextToken) > 0 {
		input.NextToken = aws.String(_macie2NextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListResourceProfileArtifacts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*macie2.ListResourceProfileArtifactsOutput
	p := macie2.NewListResourceProfileArtifactsPaginator(client, input)
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

// Retrieves information about the types and amount of sensitive data that Amazon
// Macie found in an S3 bucket.
func macie2_ListResourceProfileDetections(cfg aws.Config, client *macie2.Client) {
	input := &macie2.ListResourceProfileDetectionsInput{
		// ResourceArn: *string, // Required
	}

	if len(_macie2ResourceArn) > 0 {
		input.ResourceArn = aws.String(_macie2ResourceArn)
	}
	if len(_macie2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _macie2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_macie2NextToken) > 0 {
		input.NextToken = aws.String(_macie2NextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListResourceProfileDetections(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*macie2.ListResourceProfileDetectionsOutput
	p := macie2.NewListResourceProfileDetectionsPaginator(client, input)
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

// Retrieves a subset of information about the sensitivity inspection template for
// an account.
func macie2_ListSensitivityInspectionTemplates(cfg aws.Config, client *macie2.Client) {
	input := &macie2.ListSensitivityInspectionTemplatesInput{}

	if len(_macie2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _macie2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_macie2NextToken) > 0 {
		input.NextToken = aws.String(_macie2NextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSensitivityInspectionTemplates(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*macie2.ListSensitivityInspectionTemplatesOutput
	p := macie2.NewListSensitivityInspectionTemplatesPaginator(client, input)
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

// Retrieves the tags (keys and values) that are associated with an Amazon Macie
// resource.
func macie2_ListTagsForResource(cfg aws.Config, client *macie2.Client) {
	input := &macie2.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_macie2ResourceArn) > 0 {
		input.ResourceArn = aws.String(_macie2ResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds or updates the configuration settings for storing data classification
// results.
func macie2_PutClassificationExportConfiguration(cfg aws.Config, client *macie2.Client) {
	input := &macie2.PutClassificationExportConfigurationInput{
		// Configuration: *types.ClassificationExportConfiguration, // Required
	}

	if len(_macie2Configuration) > 0 {
		if err := assignInputField(input, "Configuration", _macie2Configuration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutClassificationExportConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration settings for publishing findings to Security Hub.
func macie2_PutFindingsPublicationConfiguration(cfg aws.Config, client *macie2.Client) {
	input := &macie2.PutFindingsPublicationConfigurationInput{}

	if len(_macie2ClientToken) > 0 {
		input.ClientToken = aws.String(_macie2ClientToken)
	}
	if len(_macie2SecurityHubConfiguration) > 0 {
		if err := assignInputField(input, "SecurityHubConfiguration", _macie2SecurityHubConfiguration); err != nil {
			log.Errorf("invalid --security-hub-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutFindingsPublicationConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves (queries) statistical data and other information about Amazon Web
// Services resources that Amazon Macie monitors and analyzes for an account.
func macie2_SearchResources(cfg aws.Config, client *macie2.Client) {
	input := &macie2.SearchResourcesInput{}

	if len(_macie2BucketCriteria) > 0 {
		if err := assignInputField(input, "BucketCriteria", _macie2BucketCriteria); err != nil {
			log.Errorf("invalid --bucket-criteria: %s", err.Error())
			return
		}
	}
	if len(_macie2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _macie2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_macie2NextToken) > 0 {
		input.NextToken = aws.String(_macie2NextToken)
	}
	if len(_macie2SortCriteria) > 0 {
		if err := assignInputField(input, "SortCriteria", _macie2SortCriteria); err != nil {
			log.Errorf("invalid --sort-criteria: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.SearchResources(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*macie2.SearchResourcesOutput
	p := macie2.NewSearchResourcesPaginator(client, input)
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

// Adds or updates one or more tags (keys and values) that are associated with an
// Amazon Macie resource.
func macie2_TagResource(cfg aws.Config, client *macie2.Client) {
	input := &macie2.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_macie2ResourceArn) > 0 {
		input.ResourceArn = aws.String(_macie2ResourceArn)
	}
	if len(_macie2Tags) > 0 {
		if err := assignInputField(input, "Tags", _macie2Tags); err != nil {
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

// Tests criteria for a custom data identifier.
func macie2_TestCustomDataIdentifier(cfg aws.Config, client *macie2.Client) {
	input := &macie2.TestCustomDataIdentifierInput{
		// Regex: *string, // Required
		// SampleText: *string, // Required
	}

	if len(_macie2Regex) > 0 {
		input.Regex = aws.String(_macie2Regex)
	}
	if len(_macie2SampleText) > 0 {
		input.SampleText = aws.String(_macie2SampleText)
	}
	if len(_macie2IgnoreWords) > 0 {
		input.IgnoreWords = append([]string(nil), _macie2IgnoreWords...)
	}
	if len(_macie2Keywords) > 0 {
		input.Keywords = append([]string(nil), _macie2Keywords...)
	}
	if len(_macie2MaximumMatchDistance) > 0 {
		if err := assignInputField(input, "MaximumMatchDistance", _macie2MaximumMatchDistance); err != nil {
			log.Errorf("invalid --maximum-match-distance: %s", err.Error())
			return
		}
	}

	if resp, err := client.TestCustomDataIdentifier(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes one or more tags (keys and values) from an Amazon Macie resource.
func macie2_UntagResource(cfg aws.Config, client *macie2.Client) {
	input := &macie2.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_macie2ResourceArn) > 0 {
		input.ResourceArn = aws.String(_macie2ResourceArn)
	}
	if len(_macie2TagKeys) > 0 {
		input.TagKeys = append([]string(nil), _macie2TagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the settings for an allow list.
func macie2_UpdateAllowList(cfg aws.Config, client *macie2.Client) {
	input := &macie2.UpdateAllowListInput{
		// Criteria: *types.AllowListCriteria, // Required
		// Id: *string, // Required
		// Name: *string, // Required
	}

	if len(_macie2Criteria) > 0 {
		if err := assignInputField(input, "Criteria", _macie2Criteria); err != nil {
			log.Errorf("invalid --criteria: %s", err.Error())
			return
		}
	}
	if len(_macie2Id) > 0 {
		input.Id = aws.String(_macie2Id)
	}
	if len(_macie2Name) > 0 {
		input.Name = aws.String(_macie2Name)
	}
	if len(_macie2Description) > 0 {
		input.Description = aws.String(_macie2Description)
	}

	if resp, err := client.UpdateAllowList(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Changes the configuration settings and status of automated sensitive data
// discovery for an organization or standalone account.
func macie2_UpdateAutomatedDiscoveryConfiguration(cfg aws.Config, client *macie2.Client) {
	input := &macie2.UpdateAutomatedDiscoveryConfigurationInput{
		// Status: types.AutomatedDiscoveryStatus, // Required
	}

	if len(_macie2Status) > 0 {
		if err := assignInputField(input, "Status", _macie2Status); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_macie2AutoEnableOrganizationMembers) > 0 {
		if err := assignInputField(input, "AutoEnableOrganizationMembers", _macie2AutoEnableOrganizationMembers); err != nil {
			log.Errorf("invalid --auto-enable-organization-members: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateAutomatedDiscoveryConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Changes the status of a classification job.
func macie2_UpdateClassificationJob(cfg aws.Config, client *macie2.Client) {
	input := &macie2.UpdateClassificationJobInput{
		// JobId: *string, // Required
		// JobStatus: types.JobStatus, // Required
	}

	if len(_macie2JobId) > 0 {
		input.JobId = aws.String(_macie2JobId)
	}
	if len(_macie2JobStatus) > 0 {
		if err := assignInputField(input, "JobStatus", _macie2JobStatus); err != nil {
			log.Errorf("invalid --job-status: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateClassificationJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the classification scope settings for an account.
func macie2_UpdateClassificationScope(cfg aws.Config, client *macie2.Client) {
	input := &macie2.UpdateClassificationScopeInput{
		// Id: *string, // Required
	}

	if len(_macie2Id) > 0 {
		input.Id = aws.String(_macie2Id)
	}
	if len(_macie2S3) > 0 {
		if err := assignInputField(input, "S3", _macie2S3); err != nil {
			log.Errorf("invalid --s3: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateClassificationScope(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the criteria and other settings for a findings filter.
func macie2_UpdateFindingsFilter(cfg aws.Config, client *macie2.Client) {
	input := &macie2.UpdateFindingsFilterInput{
		// Id: *string, // Required
	}

	if len(_macie2Id) > 0 {
		input.Id = aws.String(_macie2Id)
	}
	if len(_macie2Action) > 0 {
		if err := assignInputField(input, "Action", _macie2Action); err != nil {
			log.Errorf("invalid --action: %s", err.Error())
			return
		}
	}
	if len(_macie2ClientToken) > 0 {
		input.ClientToken = aws.String(_macie2ClientToken)
	}
	if len(_macie2Description) > 0 {
		input.Description = aws.String(_macie2Description)
	}
	if len(_macie2FindingCriteria) > 0 {
		if err := assignInputField(input, "FindingCriteria", _macie2FindingCriteria); err != nil {
			log.Errorf("invalid --finding-criteria: %s", err.Error())
			return
		}
	}
	if len(_macie2Name) > 0 {
		input.Name = aws.String(_macie2Name)
	}
	if len(_macie2Position) > 0 {
		if err := assignInputField(input, "Position", _macie2Position); err != nil {
			log.Errorf("invalid --position: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateFindingsFilter(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Suspends or re-enables Amazon Macie, or updates the configuration settings for
// a Macie account.
func macie2_UpdateMacieSession(cfg aws.Config, client *macie2.Client) {
	input := &macie2.UpdateMacieSessionInput{}

	if len(_macie2FindingPublishingFrequency) > 0 {
		if err := assignInputField(input, "FindingPublishingFrequency", _macie2FindingPublishingFrequency); err != nil {
			log.Errorf("invalid --finding-publishing-frequency: %s", err.Error())
			return
		}
	}
	if len(_macie2Status) > 0 {
		if err := assignInputField(input, "Status", _macie2Status); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateMacieSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables an Amazon Macie administrator to suspend or re-enable Macie for a
// member account.
func macie2_UpdateMemberSession(cfg aws.Config, client *macie2.Client) {
	input := &macie2.UpdateMemberSessionInput{
		// Id: *string, // Required
		// Status: types.MacieStatus, // Required
	}

	if len(_macie2Id) > 0 {
		input.Id = aws.String(_macie2Id)
	}
	if len(_macie2Status) > 0 {
		if err := assignInputField(input, "Status", _macie2Status); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateMemberSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the Amazon Macie configuration settings for an organization in
// Organizations.
func macie2_UpdateOrganizationConfiguration(cfg aws.Config, client *macie2.Client) {
	input := &macie2.UpdateOrganizationConfigurationInput{
		// AutoEnable: *bool, // Required
	}

	if len(_macie2AutoEnable) > 0 {
		if err := assignInputField(input, "AutoEnable", _macie2AutoEnable); err != nil {
			log.Errorf("invalid --auto-enable: %s", err.Error())
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

// Updates the sensitivity score for an S3 bucket.
func macie2_UpdateResourceProfile(cfg aws.Config, client *macie2.Client) {
	input := &macie2.UpdateResourceProfileInput{
		// ResourceArn: *string, // Required
	}

	if len(_macie2ResourceArn) > 0 {
		input.ResourceArn = aws.String(_macie2ResourceArn)
	}
	if len(_macie2SensitivityScoreOverride) > 0 {
		if err := assignInputField(input, "SensitivityScoreOverride", _macie2SensitivityScoreOverride); err != nil {
			log.Errorf("invalid --sensitivity-score-override: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateResourceProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the sensitivity scoring settings for an S3 bucket.
func macie2_UpdateResourceProfileDetections(cfg aws.Config, client *macie2.Client) {
	input := &macie2.UpdateResourceProfileDetectionsInput{
		// ResourceArn: *string, // Required
	}

	if len(_macie2ResourceArn) > 0 {
		input.ResourceArn = aws.String(_macie2ResourceArn)
	}
	if len(_macie2SuppressDataIdentifiers) > 0 {
		if err := assignInputField(input, "SuppressDataIdentifiers", _macie2SuppressDataIdentifiers); err != nil {
			log.Errorf("invalid --suppress-data-identifiers: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateResourceProfileDetections(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the status and configuration settings for retrieving occurrences of
// sensitive data reported by findings.
func macie2_UpdateRevealConfiguration(cfg aws.Config, client *macie2.Client) {
	input := &macie2.UpdateRevealConfigurationInput{
		// Configuration: *types.RevealConfiguration, // Required
	}

	if len(_macie2Configuration) > 0 {
		if err := assignInputField(input, "Configuration", _macie2Configuration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}
	if len(_macie2RetrievalConfiguration) > 0 {
		if err := assignInputField(input, "RetrievalConfiguration", _macie2RetrievalConfiguration); err != nil {
			log.Errorf("invalid --retrieval-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateRevealConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the settings for the sensitivity inspection template for an account.
func macie2_UpdateSensitivityInspectionTemplate(cfg aws.Config, client *macie2.Client) {
	input := &macie2.UpdateSensitivityInspectionTemplateInput{
		// Id: *string, // Required
	}

	if len(_macie2Id) > 0 {
		input.Id = aws.String(_macie2Id)
	}
	if len(_macie2Description) > 0 {
		input.Description = aws.String(_macie2Description)
	}
	if len(_macie2Excludes) > 0 {
		if err := assignInputField(input, "Excludes", _macie2Excludes); err != nil {
			log.Errorf("invalid --excludes: %s", err.Error())
			return
		}
	}
	if len(_macie2Includes) > 0 {
		if err := assignInputField(input, "Includes", _macie2Includes); err != nil {
			log.Errorf("invalid --includes: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateSensitivityInspectionTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_macie2Cmd)
	_macie2Cmd.Flags().SortFlags = false

	_macie2Cmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_macie2Cmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_macie2Cmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_macie2Cmd.Flags().StringVarP(&_macie2Account, "account", "", "", "Account")
	_macie2Cmd.Flags().StringVarP(&_macie2AccountId, "account-id", "", "", "Account ID")
	_macie2Cmd.Flags().StringSliceVarP(&_macie2AccountIds, "account-ids", "", nil, "Account Ids")
	_macie2Cmd.Flags().StringVarP(&_macie2Accounts, "accounts", "", "", "Accounts")
	_macie2Cmd.Flags().StringVarP(&_macie2Action, "action", "", "", "Action")
	_macie2Cmd.Flags().StringVarP(&_macie2AdminAccountId, "admin-account-id", "", "", "Admin Account ID")
	_macie2Cmd.Flags().StringVarP(&_macie2AdministratorAccountId, "administrator-account-id", "", "", "Administrator Account ID")
	_macie2Cmd.Flags().StringSliceVarP(&_macie2AllowListIds, "allow-list-ids", "", nil, "Allow List Ids")
	_macie2Cmd.Flags().StringVarP(&_macie2AutoEnable, "auto-enable", "", "", "Auto Enable")
	_macie2Cmd.Flags().StringVarP(&_macie2AutoEnableOrganizationMembers, "auto-enable-organization-members", "", "", "Auto Enable Organization Members")
	_macie2Cmd.Flags().StringVarP(&_macie2BucketCriteria, "bucket-criteria", "", "", "Bucket Criteria")
	_macie2Cmd.Flags().StringVarP(&_macie2ClientToken, "client-token", "", "", "Client Token")
	_macie2Cmd.Flags().StringVarP(&_macie2Configuration, "configuration", "", "", "Configuration")
	_macie2Cmd.Flags().StringVarP(&_macie2Criteria, "criteria", "", "", "Criteria")
	_macie2Cmd.Flags().StringSliceVarP(&_macie2CustomDataIdentifierIds, "custom-data-identifier-ids", "", nil, "Custom Data Identifier Ids")
	_macie2Cmd.Flags().StringVarP(&_macie2Description, "description", "", "", "Description")
	_macie2Cmd.Flags().StringVarP(&_macie2DisableEmailNotification, "disable-email-notification", "", "", "Disable Email Notification")
	_macie2Cmd.Flags().StringVarP(&_macie2Excludes, "excludes", "", "", "Excludes")
	_macie2Cmd.Flags().StringVarP(&_macie2FilterBy, "filter-by", "", "", "Filter By")
	_macie2Cmd.Flags().StringVarP(&_macie2FilterCriteria, "filter-criteria", "", "", "Filter Criteria")
	_macie2Cmd.Flags().StringVarP(&_macie2FindingCriteria, "finding-criteria", "", "", "Finding Criteria")
	_macie2Cmd.Flags().StringVarP(&_macie2FindingId, "finding-id", "", "", "Finding ID")
	_macie2Cmd.Flags().StringSliceVarP(&_macie2FindingIds, "finding-ids", "", nil, "Finding Ids")
	_macie2Cmd.Flags().StringVarP(&_macie2FindingPublishingFrequency, "finding-publishing-frequency", "", "", "Finding Publishing Frequency")
	_macie2Cmd.Flags().StringVarP(&_macie2FindingTypes, "finding-types", "", "", "Finding Types")
	_macie2Cmd.Flags().StringVarP(&_macie2GroupBy, "group-by", "", "", "Group By")
	_macie2Cmd.Flags().StringVarP(&_macie2Id, "id", "", "", "ID")
	_macie2Cmd.Flags().StringSliceVarP(&_macie2Ids, "ids", "", nil, "Ids")
	_macie2Cmd.Flags().StringVarP(&_macie2IgnoreJobChecks, "ignore-job-checks", "", "", "Ignore Job Checks")
	_macie2Cmd.Flags().StringSliceVarP(&_macie2IgnoreWords, "ignore-words", "", nil, "Ignore Words")
	_macie2Cmd.Flags().StringVarP(&_macie2Includes, "includes", "", "", "Includes")
	_macie2Cmd.Flags().StringVarP(&_macie2InitialRun, "initial-run", "", "", "Initial Run")
	_macie2Cmd.Flags().StringVarP(&_macie2InvitationId, "invitation-id", "", "", "Invitation ID")
	_macie2Cmd.Flags().StringVarP(&_macie2JobId, "job-id", "", "", "Job ID")
	_macie2Cmd.Flags().StringVarP(&_macie2JobStatus, "job-status", "", "", "Job Status")
	_macie2Cmd.Flags().StringVarP(&_macie2JobType, "job-type", "", "", "Job Type")
	_macie2Cmd.Flags().StringSliceVarP(&_macie2Keywords, "keywords", "", nil, "Keywords")
	_macie2Cmd.Flags().StringSliceVarP(&_macie2ManagedDataIdentifierIds, "managed-data-identifier-ids", "", nil, "Managed Data Identifier Ids")
	_macie2Cmd.Flags().StringVarP(&_macie2ManagedDataIdentifierSelector, "managed-data-identifier-selector", "", "", "Managed Data Identifier Selector")
	_macie2Cmd.Flags().StringVarP(&_macie2MasterAccount, "master-account", "", "", "Master Account")
	_macie2Cmd.Flags().StringVarP(&_macie2MaxResults, "max-results", "", "", "Max Results")
	_macie2Cmd.Flags().StringVarP(&_macie2MaximumMatchDistance, "maximum-match-distance", "", "", "Maximum Match Distance")
	_macie2Cmd.Flags().StringVarP(&_macie2Message, "message", "", "", "Message")
	_macie2Cmd.Flags().StringVarP(&_macie2Name, "name", "", "", "Name")
	_macie2Cmd.Flags().StringVarP(&_macie2NextToken, "next-token", "", "", "Next Token")
	_macie2Cmd.Flags().StringVarP(&_macie2OnlyAssociated, "only-associated", "", "", "Only Associated")
	_macie2Cmd.Flags().StringVarP(&_macie2Position, "position", "", "", "Position")
	_macie2Cmd.Flags().StringVarP(&_macie2Regex, "regex", "", "", "Regex")
	_macie2Cmd.Flags().StringVarP(&_macie2ResourceArn, "resource-arn", "", "", "Resource ARN")
	_macie2Cmd.Flags().StringVarP(&_macie2RetrievalConfiguration, "retrieval-configuration", "", "", "Retrieval Configuration")
	_macie2Cmd.Flags().StringVarP(&_macie2S3, "s3", "", "", "S3")
	_macie2Cmd.Flags().StringVarP(&_macie2S3JobDefinition, "s3-job-definition", "", "", "S3 Job Definition")
	_macie2Cmd.Flags().StringVarP(&_macie2SampleText, "sample-text", "", "", "Sample Text")
	_macie2Cmd.Flags().StringVarP(&_macie2SamplingPercentage, "sampling-percentage", "", "", "Sampling Percentage")
	_macie2Cmd.Flags().StringVarP(&_macie2ScheduleFrequency, "schedule-frequency", "", "", "Schedule Frequency")
	_macie2Cmd.Flags().StringVarP(&_macie2SecurityHubConfiguration, "security-hub-configuration", "", "", "Security Hub Configuration")
	_macie2Cmd.Flags().StringVarP(&_macie2SensitivityScoreOverride, "sensitivity-score-override", "", "", "Sensitivity Score Override")
	_macie2Cmd.Flags().StringVarP(&_macie2SeverityLevels, "severity-levels", "", "", "Severity Levels")
	_macie2Cmd.Flags().StringVarP(&_macie2Size, "size", "", "", "Size")
	_macie2Cmd.Flags().StringVarP(&_macie2SortBy, "sort-by", "", "", "Sort By")
	_macie2Cmd.Flags().StringVarP(&_macie2SortCriteria, "sort-criteria", "", "", "Sort Criteria")
	_macie2Cmd.Flags().StringVarP(&_macie2Status, "status", "", "", "Status")
	_macie2Cmd.Flags().StringVarP(&_macie2SuppressDataIdentifiers, "suppress-data-identifiers", "", "", "Suppress Data Identifiers")
	_macie2Cmd.Flags().StringSliceVarP(&_macie2TagKeys, "tag-keys", "", nil, "Tag Keys")
	_macie2Cmd.Flags().StringVarP(&_macie2Tags, "tags", "", "", "Tags")
	_macie2Cmd.Flags().StringVarP(&_macie2TimeRange, "time-range", "", "", "Time Range")

	_macie2Cmd.Flags().BoolVarP(&_macie2AcceptInvitation, "accept-invitation", "", false, "Accept Invitation")
	_macie2Cmd.Flags().BoolVarP(&_macie2BatchGetCustomDataIdentifiers, "batch-get-custom-data-identifiers", "", false, "Batch Get Custom Data Identifiers")
	_macie2Cmd.Flags().BoolVarP(&_macie2BatchUpdateAutomatedDiscoveryAccounts, "batch-update-automated-discovery-accounts", "", false, "Batch Update Automated Discovery Accounts")
	_macie2Cmd.Flags().BoolVarP(&_macie2CreateAllowList, "create-allow-list", "", false, "Create Allow List")
	_macie2Cmd.Flags().BoolVarP(&_macie2CreateClassificationJob, "create-classification-job", "", false, "Create Classification Job")
	_macie2Cmd.Flags().BoolVarP(&_macie2CreateCustomDataIdentifier, "create-custom-data-identifier", "", false, "Create Custom Data Identifier")
	_macie2Cmd.Flags().BoolVarP(&_macie2CreateFindingsFilter, "create-findings-filter", "", false, "Create Findings Filter")
	_macie2Cmd.Flags().BoolVarP(&_macie2CreateInvitations, "create-invitations", "", false, "Create Invitations")
	_macie2Cmd.Flags().BoolVarP(&_macie2CreateMember, "create-member", "", false, "Create Member")
	_macie2Cmd.Flags().BoolVarP(&_macie2CreateSampleFindings, "create-sample-findings", "", false, "Create Sample Findings")
	_macie2Cmd.Flags().BoolVarP(&_macie2DeclineInvitations, "decline-invitations", "", false, "Decline Invitations")
	_macie2Cmd.Flags().BoolVarP(&_macie2DeleteAllowList, "delete-allow-list", "", false, "Delete Allow List")
	_macie2Cmd.Flags().BoolVarP(&_macie2DeleteCustomDataIdentifier, "delete-custom-data-identifier", "", false, "Delete Custom Data Identifier")
	_macie2Cmd.Flags().BoolVarP(&_macie2DeleteFindingsFilter, "delete-findings-filter", "", false, "Delete Findings Filter")
	_macie2Cmd.Flags().BoolVarP(&_macie2DeleteInvitations, "delete-invitations", "", false, "Delete Invitations")
	_macie2Cmd.Flags().BoolVarP(&_macie2DeleteMember, "delete-member", "", false, "Delete Member")
	_macie2Cmd.Flags().BoolVarP(&_macie2DescribeBuckets, "describe-buckets", "", false, "Describe Buckets")
	_macie2Cmd.Flags().BoolVarP(&_macie2DescribeClassificationJob, "describe-classification-job", "", false, "Describe Classification Job")
	_macie2Cmd.Flags().BoolVarP(&_macie2DescribeOrganizationConfiguration, "describe-organization-configuration", "", false, "Describe Organization Configuration")
	_macie2Cmd.Flags().BoolVarP(&_macie2DisableMacie, "disable-macie", "", false, "Disable Macie")
	_macie2Cmd.Flags().BoolVarP(&_macie2DisableOrganizationAdminAccount, "disable-organization-admin-account", "", false, "Disable Organization Admin Account")
	_macie2Cmd.Flags().BoolVarP(&_macie2DisassociateFromAdministratorAccount, "disassociate-from-administrator-account", "", false, "Disassociate From Administrator Account")
	_macie2Cmd.Flags().BoolVarP(&_macie2DisassociateFromMasterAccount, "disassociate-from-master-account", "", false, "Disassociate From Master Account")
	_macie2Cmd.Flags().BoolVarP(&_macie2DisassociateMember, "disassociate-member", "", false, "Disassociate Member")
	_macie2Cmd.Flags().BoolVarP(&_macie2EnableMacie, "enable-macie", "", false, "Enable Macie")
	_macie2Cmd.Flags().BoolVarP(&_macie2EnableOrganizationAdminAccount, "enable-organization-admin-account", "", false, "Enable Organization Admin Account")
	_macie2Cmd.Flags().BoolVarP(&_macie2GetAdministratorAccount, "get-administrator-account", "", false, "Get Administrator Account")
	_macie2Cmd.Flags().BoolVarP(&_macie2GetAllowList, "get-allow-list", "", false, "Get Allow List")
	_macie2Cmd.Flags().BoolVarP(&_macie2GetAutomatedDiscoveryConfiguration, "get-automated-discovery-configuration", "", false, "Get Automated Discovery Configuration")
	_macie2Cmd.Flags().BoolVarP(&_macie2GetBucketStatistics, "get-bucket-statistics", "", false, "Get Bucket Statistics")
	_macie2Cmd.Flags().BoolVarP(&_macie2GetClassificationExportConfiguration, "get-classification-export-configuration", "", false, "Get Classification Export Configuration")
	_macie2Cmd.Flags().BoolVarP(&_macie2GetClassificationScope, "get-classification-scope", "", false, "Get Classification Scope")
	_macie2Cmd.Flags().BoolVarP(&_macie2GetCustomDataIdentifier, "get-custom-data-identifier", "", false, "Get Custom Data Identifier")
	_macie2Cmd.Flags().BoolVarP(&_macie2GetFindingStatistics, "get-finding-statistics", "", false, "Get Finding Statistics")
	_macie2Cmd.Flags().BoolVarP(&_macie2GetFindings, "get-findings", "", false, "Get Findings")
	_macie2Cmd.Flags().BoolVarP(&_macie2GetFindingsFilter, "get-findings-filter", "", false, "Get Findings Filter")
	_macie2Cmd.Flags().BoolVarP(&_macie2GetFindingsPublicationConfiguration, "get-findings-publication-configuration", "", false, "Get Findings Publication Configuration")
	_macie2Cmd.Flags().BoolVarP(&_macie2GetInvitationsCount, "get-invitations-count", "", false, "Get Invitations Count")
	_macie2Cmd.Flags().BoolVarP(&_macie2GetMacieSession, "get-macie-session", "", false, "Get Macie Session")
	_macie2Cmd.Flags().BoolVarP(&_macie2GetMasterAccount, "get-master-account", "", false, "Get Master Account")
	_macie2Cmd.Flags().BoolVarP(&_macie2GetMember, "get-member", "", false, "Get Member")
	_macie2Cmd.Flags().BoolVarP(&_macie2GetResourceProfile, "get-resource-profile", "", false, "Get Resource Profile")
	_macie2Cmd.Flags().BoolVarP(&_macie2GetRevealConfiguration, "get-reveal-configuration", "", false, "Get Reveal Configuration")
	_macie2Cmd.Flags().BoolVarP(&_macie2GetSensitiveDataOccurrences, "get-sensitive-data-occurrences", "", false, "Get Sensitive Data Occurrences")
	_macie2Cmd.Flags().BoolVarP(&_macie2GetSensitiveDataOccurrencesAvailability, "get-sensitive-data-occurrences-availability", "", false, "Get Sensitive Data Occurrences Availability")
	_macie2Cmd.Flags().BoolVarP(&_macie2GetSensitivityInspectionTemplate, "get-sensitivity-inspection-template", "", false, "Get Sensitivity Inspection Template")
	_macie2Cmd.Flags().BoolVarP(&_macie2GetUsageStatistics, "get-usage-statistics", "", false, "Get Usage Statistics")
	_macie2Cmd.Flags().BoolVarP(&_macie2GetUsageTotals, "get-usage-totals", "", false, "Get Usage Totals")
	_macie2Cmd.Flags().BoolVarP(&_macie2ListAllowLists, "list-allow-lists", "", false, "List Allow Lists")
	_macie2Cmd.Flags().BoolVarP(&_macie2ListAutomatedDiscoveryAccounts, "list-automated-discovery-accounts", "", false, "List Automated Discovery Accounts")
	_macie2Cmd.Flags().BoolVarP(&_macie2ListClassificationJobs, "list-classification-jobs", "", false, "List Classification Jobs")
	_macie2Cmd.Flags().BoolVarP(&_macie2ListClassificationScopes, "list-classification-scopes", "", false, "List Classification Scopes")
	_macie2Cmd.Flags().BoolVarP(&_macie2ListCustomDataIdentifiers, "list-custom-data-identifiers", "", false, "List Custom Data Identifiers")
	_macie2Cmd.Flags().BoolVarP(&_macie2ListFindings, "list-findings", "", false, "List Findings")
	_macie2Cmd.Flags().BoolVarP(&_macie2ListFindingsFilters, "list-findings-filters", "", false, "List Findings Filters")
	_macie2Cmd.Flags().BoolVarP(&_macie2ListInvitations, "list-invitations", "", false, "List Invitations")
	_macie2Cmd.Flags().BoolVarP(&_macie2ListManagedDataIdentifiers, "list-managed-data-identifiers", "", false, "List Managed Data Identifiers")
	_macie2Cmd.Flags().BoolVarP(&_macie2ListMembers, "list-members", "", false, "List Members")
	_macie2Cmd.Flags().BoolVarP(&_macie2ListOrganizationAdminAccounts, "list-organization-admin-accounts", "", false, "List Organization Admin Accounts")
	_macie2Cmd.Flags().BoolVarP(&_macie2ListResourceProfileArtifacts, "list-resource-profile-artifacts", "", false, "List Resource Profile Artifacts")
	_macie2Cmd.Flags().BoolVarP(&_macie2ListResourceProfileDetections, "list-resource-profile-detections", "", false, "List Resource Profile Detections")
	_macie2Cmd.Flags().BoolVarP(&_macie2ListSensitivityInspectionTemplates, "list-sensitivity-inspection-templates", "", false, "List Sensitivity Inspection Templates")
	_macie2Cmd.Flags().BoolVarP(&_macie2ListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_macie2Cmd.Flags().BoolVarP(&_macie2PutClassificationExportConfiguration, "put-classification-export-configuration", "", false, "Put Classification Export Configuration")
	_macie2Cmd.Flags().BoolVarP(&_macie2PutFindingsPublicationConfiguration, "put-findings-publication-configuration", "", false, "Put Findings Publication Configuration")
	_macie2Cmd.Flags().BoolVarP(&_macie2SearchResources, "search-resources", "", false, "Search Resources")
	_macie2Cmd.Flags().BoolVarP(&_macie2TagResource, "tag-resource", "", false, "Tag Resource")
	_macie2Cmd.Flags().BoolVarP(&_macie2TestCustomDataIdentifier, "test-custom-data-identifier", "", false, "Test Custom Data Identifier")
	_macie2Cmd.Flags().BoolVarP(&_macie2UntagResource, "untag-resource", "", false, "Untag Resource")
	_macie2Cmd.Flags().BoolVarP(&_macie2UpdateAllowList, "update-allow-list", "", false, "Update Allow List")
	_macie2Cmd.Flags().BoolVarP(&_macie2UpdateAutomatedDiscoveryConfiguration, "update-automated-discovery-configuration", "", false, "Update Automated Discovery Configuration")
	_macie2Cmd.Flags().BoolVarP(&_macie2UpdateClassificationJob, "update-classification-job", "", false, "Update Classification Job")
	_macie2Cmd.Flags().BoolVarP(&_macie2UpdateClassificationScope, "update-classification-scope", "", false, "Update Classification Scope")
	_macie2Cmd.Flags().BoolVarP(&_macie2UpdateFindingsFilter, "update-findings-filter", "", false, "Update Findings Filter")
	_macie2Cmd.Flags().BoolVarP(&_macie2UpdateMacieSession, "update-macie-session", "", false, "Update Macie Session")
	_macie2Cmd.Flags().BoolVarP(&_macie2UpdateMemberSession, "update-member-session", "", false, "Update Member Session")
	_macie2Cmd.Flags().BoolVarP(&_macie2UpdateOrganizationConfiguration, "update-organization-configuration", "", false, "Update Organization Configuration")
	_macie2Cmd.Flags().BoolVarP(&_macie2UpdateResourceProfile, "update-resource-profile", "", false, "Update Resource Profile")
	_macie2Cmd.Flags().BoolVarP(&_macie2UpdateResourceProfileDetections, "update-resource-profile-detections", "", false, "Update Resource Profile Detections")
	_macie2Cmd.Flags().BoolVarP(&_macie2UpdateRevealConfiguration, "update-reveal-configuration", "", false, "Update Reveal Configuration")
	_macie2Cmd.Flags().BoolVarP(&_macie2UpdateSensitivityInspectionTemplate, "update-sensitivity-inspection-template", "", false, "Update Sensitivity Inspection Template")

}
