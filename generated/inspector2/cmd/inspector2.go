package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/inspector2"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// inspector2Cmd represents the inspector2 command
var _inspector2Cmd = &cobra.Command{
	Use:   "inspector2",
	Short: "AWS inspector2 CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := inspector2.NewFromConfig(cfg)
		if _inspector2AssociateMember {
			inspector2_AssociateMember(cfg, client)
			return
		}
		if _inspector2BatchAssociateCodeSecurityScanConfiguration {
			inspector2_BatchAssociateCodeSecurityScanConfiguration(cfg, client)
			return
		}
		if _inspector2BatchDisassociateCodeSecurityScanConfiguration {
			inspector2_BatchDisassociateCodeSecurityScanConfiguration(cfg, client)
			return
		}
		if _inspector2BatchGetAccountStatus {
			inspector2_BatchGetAccountStatus(cfg, client)
			return
		}
		if _inspector2BatchGetCodeSnippet {
			inspector2_BatchGetCodeSnippet(cfg, client)
			return
		}
		if _inspector2BatchGetFindingDetails {
			inspector2_BatchGetFindingDetails(cfg, client)
			return
		}
		if _inspector2BatchGetFreeTrialInfo {
			inspector2_BatchGetFreeTrialInfo(cfg, client)
			return
		}
		if _inspector2BatchGetMemberEc2DeepInspectionStatus {
			inspector2_BatchGetMemberEc2DeepInspectionStatus(cfg, client)
			return
		}
		if _inspector2BatchUpdateMemberEc2DeepInspectionStatus {
			inspector2_BatchUpdateMemberEc2DeepInspectionStatus(cfg, client)
			return
		}
		if _inspector2CancelFindingsReport {
			inspector2_CancelFindingsReport(cfg, client)
			return
		}
		if _inspector2CancelSbomExport {
			inspector2_CancelSbomExport(cfg, client)
			return
		}
		if _inspector2CreateCisScanConfiguration {
			inspector2_CreateCisScanConfiguration(cfg, client)
			return
		}
		if _inspector2CreateCodeSecurityIntegration {
			inspector2_CreateCodeSecurityIntegration(cfg, client)
			return
		}
		if _inspector2CreateCodeSecurityScanConfiguration {
			inspector2_CreateCodeSecurityScanConfiguration(cfg, client)
			return
		}
		if _inspector2CreateFilter {
			inspector2_CreateFilter(cfg, client)
			return
		}
		if _inspector2CreateFindingsReport {
			inspector2_CreateFindingsReport(cfg, client)
			return
		}
		if _inspector2CreateSbomExport {
			inspector2_CreateSbomExport(cfg, client)
			return
		}
		if _inspector2DeleteCisScanConfiguration {
			inspector2_DeleteCisScanConfiguration(cfg, client)
			return
		}
		if _inspector2DeleteCodeSecurityIntegration {
			inspector2_DeleteCodeSecurityIntegration(cfg, client)
			return
		}
		if _inspector2DeleteCodeSecurityScanConfiguration {
			inspector2_DeleteCodeSecurityScanConfiguration(cfg, client)
			return
		}
		if _inspector2DeleteFilter {
			inspector2_DeleteFilter(cfg, client)
			return
		}
		if _inspector2DescribeOrganizationConfiguration {
			inspector2_DescribeOrganizationConfiguration(cfg, client)
			return
		}
		if _inspector2Disable {
			inspector2_Disable(cfg, client)
			return
		}
		if _inspector2DisableDelegatedAdminAccount {
			inspector2_DisableDelegatedAdminAccount(cfg, client)
			return
		}
		if _inspector2DisassociateMember {
			inspector2_DisassociateMember(cfg, client)
			return
		}
		if _inspector2Enable {
			inspector2_Enable(cfg, client)
			return
		}
		if _inspector2EnableDelegatedAdminAccount {
			inspector2_EnableDelegatedAdminAccount(cfg, client)
			return
		}
		if _inspector2GetCisScanReport {
			inspector2_GetCisScanReport(cfg, client)
			return
		}
		if _inspector2GetCisScanResultDetails {
			inspector2_GetCisScanResultDetails(cfg, client)
			return
		}
		if _inspector2GetClustersForImage {
			inspector2_GetClustersForImage(cfg, client)
			return
		}
		if _inspector2GetCodeSecurityIntegration {
			inspector2_GetCodeSecurityIntegration(cfg, client)
			return
		}
		if _inspector2GetCodeSecurityScan {
			inspector2_GetCodeSecurityScan(cfg, client)
			return
		}
		if _inspector2GetCodeSecurityScanConfiguration {
			inspector2_GetCodeSecurityScanConfiguration(cfg, client)
			return
		}
		if _inspector2GetConfiguration {
			inspector2_GetConfiguration(cfg, client)
			return
		}
		if _inspector2GetDelegatedAdminAccount {
			inspector2_GetDelegatedAdminAccount(cfg, client)
			return
		}
		if _inspector2GetEc2DeepInspectionConfiguration {
			inspector2_GetEc2DeepInspectionConfiguration(cfg, client)
			return
		}
		if _inspector2GetEncryptionKey {
			inspector2_GetEncryptionKey(cfg, client)
			return
		}
		if _inspector2GetFindingsReportStatus {
			inspector2_GetFindingsReportStatus(cfg, client)
			return
		}
		if _inspector2GetMember {
			inspector2_GetMember(cfg, client)
			return
		}
		if _inspector2GetSbomExport {
			inspector2_GetSbomExport(cfg, client)
			return
		}
		if _inspector2ListAccountPermissions {
			inspector2_ListAccountPermissions(cfg, client)
			return
		}
		if _inspector2ListCisScanConfigurations {
			inspector2_ListCisScanConfigurations(cfg, client)
			return
		}
		if _inspector2ListCisScanResultsAggregatedByChecks {
			inspector2_ListCisScanResultsAggregatedByChecks(cfg, client)
			return
		}
		if _inspector2ListCisScanResultsAggregatedByTargetResource {
			inspector2_ListCisScanResultsAggregatedByTargetResource(cfg, client)
			return
		}
		if _inspector2ListCisScans {
			inspector2_ListCisScans(cfg, client)
			return
		}
		if _inspector2ListCodeSecurityIntegrations {
			inspector2_ListCodeSecurityIntegrations(cfg, client)
			return
		}
		if _inspector2ListCodeSecurityScanConfigurationAssociations {
			inspector2_ListCodeSecurityScanConfigurationAssociations(cfg, client)
			return
		}
		if _inspector2ListCodeSecurityScanConfigurations {
			inspector2_ListCodeSecurityScanConfigurations(cfg, client)
			return
		}
		if _inspector2ListCoverage {
			inspector2_ListCoverage(cfg, client)
			return
		}
		if _inspector2ListCoverageStatistics {
			inspector2_ListCoverageStatistics(cfg, client)
			return
		}
		if _inspector2ListDelegatedAdminAccounts {
			inspector2_ListDelegatedAdminAccounts(cfg, client)
			return
		}
		if _inspector2ListFilters {
			inspector2_ListFilters(cfg, client)
			return
		}
		if _inspector2ListFindingAggregations {
			inspector2_ListFindingAggregations(cfg, client)
			return
		}
		if _inspector2ListFindings {
			inspector2_ListFindings(cfg, client)
			return
		}
		if _inspector2ListMembers {
			inspector2_ListMembers(cfg, client)
			return
		}
		if _inspector2ListTagsForResource {
			inspector2_ListTagsForResource(cfg, client)
			return
		}
		if _inspector2ListUsageTotals {
			inspector2_ListUsageTotals(cfg, client)
			return
		}
		if _inspector2ResetEncryptionKey {
			inspector2_ResetEncryptionKey(cfg, client)
			return
		}
		if _inspector2SearchVulnerabilities {
			inspector2_SearchVulnerabilities(cfg, client)
			return
		}
		if _inspector2SendCisSessionHealth {
			inspector2_SendCisSessionHealth(cfg, client)
			return
		}
		if _inspector2SendCisSessionTelemetry {
			inspector2_SendCisSessionTelemetry(cfg, client)
			return
		}
		if _inspector2StartCisSession {
			inspector2_StartCisSession(cfg, client)
			return
		}
		if _inspector2StartCodeSecurityScan {
			inspector2_StartCodeSecurityScan(cfg, client)
			return
		}
		if _inspector2StopCisSession {
			inspector2_StopCisSession(cfg, client)
			return
		}
		if _inspector2TagResource {
			inspector2_TagResource(cfg, client)
			return
		}
		if _inspector2UntagResource {
			inspector2_UntagResource(cfg, client)
			return
		}
		if _inspector2UpdateCisScanConfiguration {
			inspector2_UpdateCisScanConfiguration(cfg, client)
			return
		}
		if _inspector2UpdateCodeSecurityIntegration {
			inspector2_UpdateCodeSecurityIntegration(cfg, client)
			return
		}
		if _inspector2UpdateCodeSecurityScanConfiguration {
			inspector2_UpdateCodeSecurityScanConfiguration(cfg, client)
			return
		}
		if _inspector2UpdateConfiguration {
			inspector2_UpdateConfiguration(cfg, client)
			return
		}
		if _inspector2UpdateEc2DeepInspectionConfiguration {
			inspector2_UpdateEc2DeepInspectionConfiguration(cfg, client)
			return
		}
		if _inspector2UpdateEncryptionKey {
			inspector2_UpdateEncryptionKey(cfg, client)
			return
		}
		if _inspector2UpdateFilter {
			inspector2_UpdateFilter(cfg, client)
			return
		}
		if _inspector2UpdateOrgEc2DeepInspectionConfiguration {
			inspector2_UpdateOrgEc2DeepInspectionConfiguration(cfg, client)
			return
		}
		if _inspector2UpdateOrganizationConfiguration {
			inspector2_UpdateOrganizationConfiguration(cfg, client)
			return
		}

	},
}

var (
	_inspector2AssociateMember                                bool
	_inspector2BatchAssociateCodeSecurityScanConfiguration    bool
	_inspector2BatchDisassociateCodeSecurityScanConfiguration bool
	_inspector2BatchGetAccountStatus                          bool
	_inspector2BatchGetCodeSnippet                            bool
	_inspector2BatchGetFindingDetails                         bool
	_inspector2BatchGetFreeTrialInfo                          bool
	_inspector2BatchGetMemberEc2DeepInspectionStatus          bool
	_inspector2BatchUpdateMemberEc2DeepInspectionStatus       bool
	_inspector2CancelFindingsReport                           bool
	_inspector2CancelSbomExport                               bool
	_inspector2CreateCisScanConfiguration                     bool
	_inspector2CreateCodeSecurityIntegration                  bool
	_inspector2CreateCodeSecurityScanConfiguration            bool
	_inspector2CreateFilter                                   bool
	_inspector2CreateFindingsReport                           bool
	_inspector2CreateSbomExport                               bool
	_inspector2DeleteCisScanConfiguration                     bool
	_inspector2DeleteCodeSecurityIntegration                  bool
	_inspector2DeleteCodeSecurityScanConfiguration            bool
	_inspector2DeleteFilter                                   bool
	_inspector2DescribeOrganizationConfiguration              bool
	_inspector2Disable                                        bool
	_inspector2DisableDelegatedAdminAccount                   bool
	_inspector2DisassociateMember                             bool
	_inspector2Enable                                         bool
	_inspector2EnableDelegatedAdminAccount                    bool
	_inspector2GetCisScanReport                               bool
	_inspector2GetCisScanResultDetails                        bool
	_inspector2GetClustersForImage                            bool
	_inspector2GetCodeSecurityIntegration                     bool
	_inspector2GetCodeSecurityScan                            bool
	_inspector2GetCodeSecurityScanConfiguration               bool
	_inspector2GetConfiguration                               bool
	_inspector2GetDelegatedAdminAccount                       bool
	_inspector2GetEc2DeepInspectionConfiguration              bool
	_inspector2GetEncryptionKey                               bool
	_inspector2GetFindingsReportStatus                        bool
	_inspector2GetMember                                      bool
	_inspector2GetSbomExport                                  bool
	_inspector2ListAccountPermissions                         bool
	_inspector2ListCisScanConfigurations                      bool
	_inspector2ListCisScanResultsAggregatedByChecks           bool
	_inspector2ListCisScanResultsAggregatedByTargetResource   bool
	_inspector2ListCisScans                                   bool
	_inspector2ListCodeSecurityIntegrations                   bool
	_inspector2ListCodeSecurityScanConfigurationAssociations  bool
	_inspector2ListCodeSecurityScanConfigurations             bool
	_inspector2ListCoverage                                   bool
	_inspector2ListCoverageStatistics                         bool
	_inspector2ListDelegatedAdminAccounts                     bool
	_inspector2ListFilters                                    bool
	_inspector2ListFindingAggregations                        bool
	_inspector2ListFindings                                   bool
	_inspector2ListMembers                                    bool
	_inspector2ListTagsForResource                            bool
	_inspector2ListUsageTotals                                bool
	_inspector2ResetEncryptionKey                             bool
	_inspector2SearchVulnerabilities                          bool
	_inspector2SendCisSessionHealth                           bool
	_inspector2SendCisSessionTelemetry                        bool
	_inspector2StartCisSession                                bool
	_inspector2StartCodeSecurityScan                          bool
	_inspector2StopCisSession                                 bool
	_inspector2TagResource                                    bool
	_inspector2UntagResource                                  bool
	_inspector2UpdateCisScanConfiguration                     bool
	_inspector2UpdateCodeSecurityIntegration                  bool
	_inspector2UpdateCodeSecurityScanConfiguration            bool
	_inspector2UpdateConfiguration                            bool
	_inspector2UpdateEc2DeepInspectionConfiguration           bool
	_inspector2UpdateEncryptionKey                            bool
	_inspector2UpdateFilter                                   bool
	_inspector2UpdateOrgEc2DeepInspectionConfiguration        bool
	_inspector2UpdateOrganizationConfiguration                bool

	_inspector2AccountId                         string
	_inspector2AccountIds                        []string
	_inspector2Action                            string
	_inspector2ActivateDeepInspection            string
	_inspector2AggregationRequest                string
	_inspector2AggregationType                   string
	_inspector2Arn                               string
	_inspector2Arns                              []string
	_inspector2AssociateConfigurationRequests    string
	_inspector2AutoEnable                        string
	_inspector2ClientToken                       string
	_inspector2Configuration                     string
	_inspector2DelegatedAdminAccountId           string
	_inspector2Description                       string
	_inspector2DetailLevel                       string
	_inspector2Details                           string
	_inspector2DisassociateConfigurationRequests string
	_inspector2Ec2Configuration                  string
	_inspector2EcrConfiguration                  string
	_inspector2Filter                            string
	_inspector2FilterArn                         string
	_inspector2FilterCriteria                    string
	_inspector2FindingArns                       []string
	_inspector2GroupBy                           string
	_inspector2IntegrationArn                    string
	_inspector2KmsKeyId                          string
	_inspector2Level                             string
	_inspector2MaxResults                        string
	_inspector2Message                           string
	_inspector2Messages                          string
	_inspector2Name                              string
	_inspector2NextToken                         string
	_inspector2OnlyAssociated                    string
	_inspector2OrgPackagePaths                   []string
	_inspector2PackagePaths                      []string
	_inspector2Reason                            string
	_inspector2ReportFormat                      string
	_inspector2ReportId                          string
	_inspector2Resource                          string
	_inspector2ResourceArn                       string
	_inspector2ResourceFilterCriteria            string
	_inspector2ResourceType                      string
	_inspector2ResourceTypes                     string
	_inspector2S3Destination                     string
	_inspector2ScanArn                           string
	_inspector2ScanConfigurationArn              string
	_inspector2ScanId                            string
	_inspector2ScanJobId                         string
	_inspector2ScanName                          string
	_inspector2ScanType                          string
	_inspector2Schedule                          string
	_inspector2ScopeSettings                     string
	_inspector2SecurityLevel                     string
	_inspector2Service                           string
	_inspector2SessionToken                      string
	_inspector2SortBy                            string
	_inspector2SortCriteria                      string
	_inspector2SortOrder                         string
	_inspector2TagKeys                           []string
	_inspector2Tags                              string
	_inspector2TargetAccounts                    []string
	_inspector2TargetResourceId                  string
	_inspector2Targets                           string
	_inspector2Type                              string
)

// Associates an Amazon Web Services account with an Amazon Inspector delegated
// administrator. An HTTP 200 response indicates the association was successfully
// started, but doesn’t indicate whether it was completed. You can check if the
// association completed by using [ListMembers]for multiple accounts or [GetMembers] for a single account.
//
// [ListMembers]: https://docs.aws.amazon.com/inspector/v2/APIReference/API_ListMembers.html
// [GetMembers]: https://docs.aws.amazon.com/inspector/v2/APIReference/API_GetMember.html
func inspector2_AssociateMember(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.AssociateMemberInput{
		// AccountId: *string, // Required
	}

	if len(_inspector2AccountId) > 0 {
		input.AccountId = aws.String(_inspector2AccountId)
	}

	if resp, err := client.AssociateMember(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates multiple code repositories with an Amazon Inspector code security
// scan configuration.
func inspector2_BatchAssociateCodeSecurityScanConfiguration(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.BatchAssociateCodeSecurityScanConfigurationInput{
		// AssociateConfigurationRequests: []types.AssociateConfigurationRequest, // Required
	}

	if len(_inspector2AssociateConfigurationRequests) > 0 {
		if err := assignInputField(input, "AssociateConfigurationRequests", _inspector2AssociateConfigurationRequests); err != nil {
			log.Errorf("invalid --associate-configuration-requests: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchAssociateCodeSecurityScanConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates multiple code repositories from an Amazon Inspector code security
// scan configuration.
func inspector2_BatchDisassociateCodeSecurityScanConfiguration(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.BatchDisassociateCodeSecurityScanConfigurationInput{
		// DisassociateConfigurationRequests: []types.DisassociateConfigurationRequest, // Required
	}

	if len(_inspector2DisassociateConfigurationRequests) > 0 {
		if err := assignInputField(input, "DisassociateConfigurationRequests", _inspector2DisassociateConfigurationRequests); err != nil {
			log.Errorf("invalid --disassociate-configuration-requests: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchDisassociateCodeSecurityScanConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the Amazon Inspector status of multiple Amazon Web Services accounts
// within your environment.
func inspector2_BatchGetAccountStatus(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.BatchGetAccountStatusInput{}

	if len(_inspector2AccountIds) > 0 {
		input.AccountIds = append([]string(nil), _inspector2AccountIds...)
	}

	if resp, err := client.BatchGetAccountStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves code snippets from findings that Amazon Inspector detected code
// vulnerabilities in.
func inspector2_BatchGetCodeSnippet(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.BatchGetCodeSnippetInput{
		// FindingArns: []string, // Required
	}

	if len(_inspector2FindingArns) > 0 {
		input.FindingArns = append([]string(nil), _inspector2FindingArns...)
	}

	if resp, err := client.BatchGetCodeSnippet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets vulnerability details for findings.
func inspector2_BatchGetFindingDetails(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.BatchGetFindingDetailsInput{
		// FindingArns: []string, // Required
	}

	if len(_inspector2FindingArns) > 0 {
		input.FindingArns = append([]string(nil), _inspector2FindingArns...)
	}

	if resp, err := client.BatchGetFindingDetails(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets free trial status for multiple Amazon Web Services accounts.
func inspector2_BatchGetFreeTrialInfo(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.BatchGetFreeTrialInfoInput{
		// AccountIds: []string, // Required
	}

	if len(_inspector2AccountIds) > 0 {
		input.AccountIds = append([]string(nil), _inspector2AccountIds...)
	}

	if resp, err := client.BatchGetFreeTrialInfo(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves Amazon Inspector deep inspection activation status of multiple member
// accounts within your organization. You must be the delegated administrator of an
// organization in Amazon Inspector to use this API.
func inspector2_BatchGetMemberEc2DeepInspectionStatus(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.BatchGetMemberEc2DeepInspectionStatusInput{}

	if len(_inspector2AccountIds) > 0 {
		input.AccountIds = append([]string(nil), _inspector2AccountIds...)
	}

	if resp, err := client.BatchGetMemberEc2DeepInspectionStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Activates or deactivates Amazon Inspector deep inspection for the provided
// member accounts in your organization. You must be the delegated administrator of
// an organization in Amazon Inspector to use this API.
func inspector2_BatchUpdateMemberEc2DeepInspectionStatus(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.BatchUpdateMemberEc2DeepInspectionStatusInput{
		// AccountIds: []types.MemberAccountEc2DeepInspectionStatus, // Required
	}

	if len(_inspector2AccountIds) > 0 {
		if err := assignInputField(input, "AccountIds", _inspector2AccountIds[0]); err != nil {
			log.Errorf("invalid --account-ids: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchUpdateMemberEc2DeepInspectionStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels the given findings report.
func inspector2_CancelFindingsReport(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.CancelFindingsReportInput{
		// ReportId: *string, // Required
	}

	if len(_inspector2ReportId) > 0 {
		input.ReportId = aws.String(_inspector2ReportId)
	}

	if resp, err := client.CancelFindingsReport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels a software bill of materials (SBOM) report.
func inspector2_CancelSbomExport(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.CancelSbomExportInput{
		// ReportId: *string, // Required
	}

	if len(_inspector2ReportId) > 0 {
		input.ReportId = aws.String(_inspector2ReportId)
	}

	if resp, err := client.CancelSbomExport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a CIS scan configuration.
func inspector2_CreateCisScanConfiguration(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.CreateCisScanConfigurationInput{
		// ScanName: *string, // Required
		// Schedule: types.Schedule, // Required
		// SecurityLevel: types.CisSecurityLevel, // Required
		// Targets: *types.CreateCisTargets, // Required
	}

	if len(_inspector2ScanName) > 0 {
		input.ScanName = aws.String(_inspector2ScanName)
	}
	if len(_inspector2Schedule) > 0 {
		if err := assignInputField(input, "Schedule", _inspector2Schedule); err != nil {
			log.Errorf("invalid --schedule: %s", err.Error())
			return
		}
	}
	if len(_inspector2SecurityLevel) > 0 {
		if err := assignInputField(input, "SecurityLevel", _inspector2SecurityLevel); err != nil {
			log.Errorf("invalid --security-level: %s", err.Error())
			return
		}
	}
	if len(_inspector2Targets) > 0 {
		if err := assignInputField(input, "Targets", _inspector2Targets); err != nil {
			log.Errorf("invalid --targets: %s", err.Error())
			return
		}
	}
	if len(_inspector2Tags) > 0 {
		if err := assignInputField(input, "Tags", _inspector2Tags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCisScanConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a code security integration with a source code repository provider.
// After calling the CreateCodeSecurityIntegration operation, you complete
// authentication and authorization with your provider. Next you call the
// UpdateCodeSecurityIntegration operation to provide the details to complete the
// integration setup
func inspector2_CreateCodeSecurityIntegration(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.CreateCodeSecurityIntegrationInput{
		// Name: *string, // Required
		// Type: types.IntegrationType, // Required
	}

	if len(_inspector2Name) > 0 {
		input.Name = aws.String(_inspector2Name)
	}
	if len(_inspector2Type) > 0 {
		if err := assignInputField(input, "Type", _inspector2Type); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_inspector2Details) > 0 {
		if err := assignInputField(input, "Details", _inspector2Details); err != nil {
			log.Errorf("invalid --details: %s", err.Error())
			return
		}
	}
	if len(_inspector2Tags) > 0 {
		if err := assignInputField(input, "Tags", _inspector2Tags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCodeSecurityIntegration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a scan configuration for code security scanning.
func inspector2_CreateCodeSecurityScanConfiguration(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.CreateCodeSecurityScanConfigurationInput{
		// Configuration: *types.CodeSecurityScanConfiguration, // Required
		// Level: types.ConfigurationLevel, // Required
		// Name: *string, // Required
	}

	if len(_inspector2Configuration) > 0 {
		if err := assignInputField(input, "Configuration", _inspector2Configuration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}
	if len(_inspector2Level) > 0 {
		if err := assignInputField(input, "Level", _inspector2Level); err != nil {
			log.Errorf("invalid --level: %s", err.Error())
			return
		}
	}
	if len(_inspector2Name) > 0 {
		input.Name = aws.String(_inspector2Name)
	}
	if len(_inspector2ScopeSettings) > 0 {
		if err := assignInputField(input, "ScopeSettings", _inspector2ScopeSettings); err != nil {
			log.Errorf("invalid --scope-settings: %s", err.Error())
			return
		}
	}
	if len(_inspector2Tags) > 0 {
		if err := assignInputField(input, "Tags", _inspector2Tags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCodeSecurityScanConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a filter resource using specified filter criteria. When the filter
// action is set to SUPPRESS this action creates a suppression rule.
func inspector2_CreateFilter(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.CreateFilterInput{
		// Action: types.FilterAction, // Required
		// FilterCriteria: *types.FilterCriteria, // Required
		// Name: *string, // Required
	}

	if len(_inspector2Action) > 0 {
		if err := assignInputField(input, "Action", _inspector2Action); err != nil {
			log.Errorf("invalid --action: %s", err.Error())
			return
		}
	}
	if len(_inspector2FilterCriteria) > 0 {
		if err := assignInputField(input, "FilterCriteria", _inspector2FilterCriteria); err != nil {
			log.Errorf("invalid --filter-criteria: %s", err.Error())
			return
		}
	}
	if len(_inspector2Name) > 0 {
		input.Name = aws.String(_inspector2Name)
	}
	if len(_inspector2Description) > 0 {
		input.Description = aws.String(_inspector2Description)
	}
	if len(_inspector2Reason) > 0 {
		input.Reason = aws.String(_inspector2Reason)
	}
	if len(_inspector2Tags) > 0 {
		if err := assignInputField(input, "Tags", _inspector2Tags); err != nil {
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

// Creates a finding report. By default only ACTIVE findings are returned in the
// report. To see SUPRESSED or CLOSED findings you must specify a value for the
// findingStatus filter criteria.
func inspector2_CreateFindingsReport(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.CreateFindingsReportInput{
		// ReportFormat: types.ReportFormat, // Required
		// S3Destination: *types.Destination, // Required
	}

	if len(_inspector2ReportFormat) > 0 {
		if err := assignInputField(input, "ReportFormat", _inspector2ReportFormat); err != nil {
			log.Errorf("invalid --report-format: %s", err.Error())
			return
		}
	}
	if len(_inspector2S3Destination) > 0 {
		if err := assignInputField(input, "S3Destination", _inspector2S3Destination); err != nil {
			log.Errorf("invalid --s3-destination: %s", err.Error())
			return
		}
	}
	if len(_inspector2FilterCriteria) > 0 {
		if err := assignInputField(input, "FilterCriteria", _inspector2FilterCriteria); err != nil {
			log.Errorf("invalid --filter-criteria: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateFindingsReport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a software bill of materials (SBOM) report.
func inspector2_CreateSbomExport(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.CreateSbomExportInput{
		// ReportFormat: types.SbomReportFormat, // Required
		// S3Destination: *types.Destination, // Required
	}

	if len(_inspector2ReportFormat) > 0 {
		if err := assignInputField(input, "ReportFormat", _inspector2ReportFormat); err != nil {
			log.Errorf("invalid --report-format: %s", err.Error())
			return
		}
	}
	if len(_inspector2S3Destination) > 0 {
		if err := assignInputField(input, "S3Destination", _inspector2S3Destination); err != nil {
			log.Errorf("invalid --s3-destination: %s", err.Error())
			return
		}
	}
	if len(_inspector2ResourceFilterCriteria) > 0 {
		if err := assignInputField(input, "ResourceFilterCriteria", _inspector2ResourceFilterCriteria); err != nil {
			log.Errorf("invalid --resource-filter-criteria: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSbomExport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a CIS scan configuration.
func inspector2_DeleteCisScanConfiguration(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.DeleteCisScanConfigurationInput{
		// ScanConfigurationArn: *string, // Required
	}

	if len(_inspector2ScanConfigurationArn) > 0 {
		input.ScanConfigurationArn = aws.String(_inspector2ScanConfigurationArn)
	}

	if resp, err := client.DeleteCisScanConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a code security integration.
func inspector2_DeleteCodeSecurityIntegration(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.DeleteCodeSecurityIntegrationInput{
		// IntegrationArn: *string, // Required
	}

	if len(_inspector2IntegrationArn) > 0 {
		input.IntegrationArn = aws.String(_inspector2IntegrationArn)
	}

	if resp, err := client.DeleteCodeSecurityIntegration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a code security scan configuration.
func inspector2_DeleteCodeSecurityScanConfiguration(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.DeleteCodeSecurityScanConfigurationInput{
		// ScanConfigurationArn: *string, // Required
	}

	if len(_inspector2ScanConfigurationArn) > 0 {
		input.ScanConfigurationArn = aws.String(_inspector2ScanConfigurationArn)
	}

	if resp, err := client.DeleteCodeSecurityScanConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a filter resource.
func inspector2_DeleteFilter(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.DeleteFilterInput{
		// Arn: *string, // Required
	}

	if len(_inspector2Arn) > 0 {
		input.Arn = aws.String(_inspector2Arn)
	}

	if resp, err := client.DeleteFilter(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describe Amazon Inspector configuration settings for an Amazon Web Services
// organization.
func inspector2_DescribeOrganizationConfiguration(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.DescribeOrganizationConfigurationInput{}

	if resp, err := client.DescribeOrganizationConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disables Amazon Inspector scans for one or more Amazon Web Services accounts.
// Disabling all scan types in an account disables the Amazon Inspector service.
func inspector2_Disable(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.DisableInput{}

	if len(_inspector2AccountIds) > 0 {
		input.AccountIds = append([]string(nil), _inspector2AccountIds...)
	}
	if len(_inspector2ResourceTypes) > 0 {
		if err := assignInputField(input, "ResourceTypes", _inspector2ResourceTypes); err != nil {
			log.Errorf("invalid --resource-types: %s", err.Error())
			return
		}
	}

	if resp, err := client.Disable(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disables the Amazon Inspector delegated administrator for your organization.
func inspector2_DisableDelegatedAdminAccount(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.DisableDelegatedAdminAccountInput{
		// DelegatedAdminAccountId: *string, // Required
	}

	if len(_inspector2DelegatedAdminAccountId) > 0 {
		input.DelegatedAdminAccountId = aws.String(_inspector2DelegatedAdminAccountId)
	}

	if resp, err := client.DisableDelegatedAdminAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates a member account from an Amazon Inspector delegated administrator.
func inspector2_DisassociateMember(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.DisassociateMemberInput{
		// AccountId: *string, // Required
	}

	if len(_inspector2AccountId) > 0 {
		input.AccountId = aws.String(_inspector2AccountId)
	}

	if resp, err := client.DisassociateMember(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables Amazon Inspector scans for one or more Amazon Web Services accounts.
func inspector2_Enable(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.EnableInput{
		// ResourceTypes: []types.ResourceScanType, // Required
	}

	if len(_inspector2ResourceTypes) > 0 {
		if err := assignInputField(input, "ResourceTypes", _inspector2ResourceTypes); err != nil {
			log.Errorf("invalid --resource-types: %s", err.Error())
			return
		}
	}
	if len(_inspector2AccountIds) > 0 {
		input.AccountIds = append([]string(nil), _inspector2AccountIds...)
	}
	if len(_inspector2ClientToken) > 0 {
		input.ClientToken = aws.String(_inspector2ClientToken)
	}

	if resp, err := client.Enable(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables the Amazon Inspector delegated administrator for your Organizations
// organization.
func inspector2_EnableDelegatedAdminAccount(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.EnableDelegatedAdminAccountInput{
		// DelegatedAdminAccountId: *string, // Required
	}

	if len(_inspector2DelegatedAdminAccountId) > 0 {
		input.DelegatedAdminAccountId = aws.String(_inspector2DelegatedAdminAccountId)
	}
	if len(_inspector2ClientToken) > 0 {
		input.ClientToken = aws.String(_inspector2ClientToken)
	}

	if resp, err := client.EnableDelegatedAdminAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a CIS scan report.
func inspector2_GetCisScanReport(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.GetCisScanReportInput{
		// ScanArn: *string, // Required
	}

	if len(_inspector2ScanArn) > 0 {
		input.ScanArn = aws.String(_inspector2ScanArn)
	}
	if len(_inspector2ReportFormat) > 0 {
		if err := assignInputField(input, "ReportFormat", _inspector2ReportFormat); err != nil {
			log.Errorf("invalid --report-format: %s", err.Error())
			return
		}
	}
	if len(_inspector2TargetAccounts) > 0 {
		input.TargetAccounts = append([]string(nil), _inspector2TargetAccounts...)
	}

	if resp, err := client.GetCisScanReport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves CIS scan result details.
func inspector2_GetCisScanResultDetails(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.GetCisScanResultDetailsInput{
		// AccountId: *string, // Required
		// ScanArn: *string, // Required
		// TargetResourceId: *string, // Required
	}

	if len(_inspector2AccountId) > 0 {
		input.AccountId = aws.String(_inspector2AccountId)
	}
	if len(_inspector2ScanArn) > 0 {
		input.ScanArn = aws.String(_inspector2ScanArn)
	}
	if len(_inspector2TargetResourceId) > 0 {
		input.TargetResourceId = aws.String(_inspector2TargetResourceId)
	}
	if len(_inspector2FilterCriteria) > 0 {
		if err := assignInputField(input, "FilterCriteria", _inspector2FilterCriteria); err != nil {
			log.Errorf("invalid --filter-criteria: %s", err.Error())
			return
		}
	}
	if len(_inspector2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _inspector2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_inspector2NextToken) > 0 {
		input.NextToken = aws.String(_inspector2NextToken)
	}
	if len(_inspector2SortBy) > 0 {
		if err := assignInputField(input, "SortBy", _inspector2SortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_inspector2SortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _inspector2SortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.GetCisScanResultDetails(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*inspector2.GetCisScanResultDetailsOutput
	p := inspector2.NewGetCisScanResultDetailsPaginator(client, input)
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

// Returns a list of clusters and metadata associated with an image.
func inspector2_GetClustersForImage(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.GetClustersForImageInput{
		// Filter: *types.ClusterForImageFilterCriteria, // Required
	}

	if len(_inspector2Filter) > 0 {
		if err := assignInputField(input, "Filter", _inspector2Filter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_inspector2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _inspector2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_inspector2NextToken) > 0 {
		input.NextToken = aws.String(_inspector2NextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetClustersForImage(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*inspector2.GetClustersForImageOutput
	p := inspector2.NewGetClustersForImagePaginator(client, input)
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

// Retrieves information about a code security integration.
func inspector2_GetCodeSecurityIntegration(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.GetCodeSecurityIntegrationInput{
		// IntegrationArn: *string, // Required
	}

	if len(_inspector2IntegrationArn) > 0 {
		input.IntegrationArn = aws.String(_inspector2IntegrationArn)
	}
	if len(_inspector2Tags) > 0 {
		if err := assignInputField(input, "Tags", _inspector2Tags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetCodeSecurityIntegration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a specific code security scan.
func inspector2_GetCodeSecurityScan(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.GetCodeSecurityScanInput{
		// Resource: types.CodeSecurityResource, // Required
		// ScanId: *string, // Required
	}

	if len(_inspector2Resource) > 0 {
		if err := assignInputField(input, "Resource", _inspector2Resource); err != nil {
			log.Errorf("invalid --resource: %s", err.Error())
			return
		}
	}
	if len(_inspector2ScanId) > 0 {
		input.ScanId = aws.String(_inspector2ScanId)
	}

	if resp, err := client.GetCodeSecurityScan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a code security scan configuration.
func inspector2_GetCodeSecurityScanConfiguration(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.GetCodeSecurityScanConfigurationInput{
		// ScanConfigurationArn: *string, // Required
	}

	if len(_inspector2ScanConfigurationArn) > 0 {
		input.ScanConfigurationArn = aws.String(_inspector2ScanConfigurationArn)
	}

	if resp, err := client.GetCodeSecurityScanConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves setting configurations for Inspector scans.
func inspector2_GetConfiguration(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.GetConfigurationInput{}

	if resp, err := client.GetConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the Amazon Inspector delegated administrator for
// your organization.
func inspector2_GetDelegatedAdminAccount(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.GetDelegatedAdminAccountInput{}

	if resp, err := client.GetDelegatedAdminAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the activation status of Amazon Inspector deep inspection and custom
// paths associated with your account.
func inspector2_GetEc2DeepInspectionConfiguration(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.GetEc2DeepInspectionConfigurationInput{}

	if resp, err := client.GetEc2DeepInspectionConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets an encryption key.
func inspector2_GetEncryptionKey(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.GetEncryptionKeyInput{
		// ResourceType: types.ResourceType, // Required
		// ScanType: types.ScanType, // Required
	}

	if len(_inspector2ResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _inspector2ResourceType); err != nil {
			log.Errorf("invalid --resource-type: %s", err.Error())
			return
		}
	}
	if len(_inspector2ScanType) > 0 {
		if err := assignInputField(input, "ScanType", _inspector2ScanType); err != nil {
			log.Errorf("invalid --scan-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetEncryptionKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the status of a findings report.
func inspector2_GetFindingsReportStatus(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.GetFindingsReportStatusInput{}

	if len(_inspector2ReportId) > 0 {
		input.ReportId = aws.String(_inspector2ReportId)
	}

	if resp, err := client.GetFindingsReportStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets member information for your organization.
func inspector2_GetMember(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.GetMemberInput{
		// AccountId: *string, // Required
	}

	if len(_inspector2AccountId) > 0 {
		input.AccountId = aws.String(_inspector2AccountId)
	}

	if resp, err := client.GetMember(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets details of a software bill of materials (SBOM) report.
func inspector2_GetSbomExport(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.GetSbomExportInput{
		// ReportId: *string, // Required
	}

	if len(_inspector2ReportId) > 0 {
		input.ReportId = aws.String(_inspector2ReportId)
	}

	if resp, err := client.GetSbomExport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the permissions an account has to configure Amazon Inspector. If the
// account is a member account or standalone account with resources managed by an
// Organizations policy, the operation returns fewer permissions.
func inspector2_ListAccountPermissions(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.ListAccountPermissionsInput{}

	if len(_inspector2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _inspector2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_inspector2NextToken) > 0 {
		input.NextToken = aws.String(_inspector2NextToken)
	}
	if len(_inspector2Service) > 0 {
		if err := assignInputField(input, "Service", _inspector2Service); err != nil {
			log.Errorf("invalid --service: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListAccountPermissions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*inspector2.ListAccountPermissionsOutput
	p := inspector2.NewListAccountPermissionsPaginator(client, input)
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

// Lists CIS scan configurations.
func inspector2_ListCisScanConfigurations(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.ListCisScanConfigurationsInput{}

	if len(_inspector2FilterCriteria) > 0 {
		if err := assignInputField(input, "FilterCriteria", _inspector2FilterCriteria); err != nil {
			log.Errorf("invalid --filter-criteria: %s", err.Error())
			return
		}
	}
	if len(_inspector2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _inspector2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_inspector2NextToken) > 0 {
		input.NextToken = aws.String(_inspector2NextToken)
	}
	if len(_inspector2SortBy) > 0 {
		if err := assignInputField(input, "SortBy", _inspector2SortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_inspector2SortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _inspector2SortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListCisScanConfigurations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*inspector2.ListCisScanConfigurationsOutput
	p := inspector2.NewListCisScanConfigurationsPaginator(client, input)
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

// Lists scan results aggregated by checks.
func inspector2_ListCisScanResultsAggregatedByChecks(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.ListCisScanResultsAggregatedByChecksInput{
		// ScanArn: *string, // Required
	}

	if len(_inspector2ScanArn) > 0 {
		input.ScanArn = aws.String(_inspector2ScanArn)
	}
	if len(_inspector2FilterCriteria) > 0 {
		if err := assignInputField(input, "FilterCriteria", _inspector2FilterCriteria); err != nil {
			log.Errorf("invalid --filter-criteria: %s", err.Error())
			return
		}
	}
	if len(_inspector2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _inspector2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_inspector2NextToken) > 0 {
		input.NextToken = aws.String(_inspector2NextToken)
	}
	if len(_inspector2SortBy) > 0 {
		if err := assignInputField(input, "SortBy", _inspector2SortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_inspector2SortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _inspector2SortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListCisScanResultsAggregatedByChecks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*inspector2.ListCisScanResultsAggregatedByChecksOutput
	p := inspector2.NewListCisScanResultsAggregatedByChecksPaginator(client, input)
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

// Lists scan results aggregated by a target resource.
func inspector2_ListCisScanResultsAggregatedByTargetResource(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.ListCisScanResultsAggregatedByTargetResourceInput{
		// ScanArn: *string, // Required
	}

	if len(_inspector2ScanArn) > 0 {
		input.ScanArn = aws.String(_inspector2ScanArn)
	}
	if len(_inspector2FilterCriteria) > 0 {
		if err := assignInputField(input, "FilterCriteria", _inspector2FilterCriteria); err != nil {
			log.Errorf("invalid --filter-criteria: %s", err.Error())
			return
		}
	}
	if len(_inspector2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _inspector2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_inspector2NextToken) > 0 {
		input.NextToken = aws.String(_inspector2NextToken)
	}
	if len(_inspector2SortBy) > 0 {
		if err := assignInputField(input, "SortBy", _inspector2SortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_inspector2SortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _inspector2SortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListCisScanResultsAggregatedByTargetResource(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*inspector2.ListCisScanResultsAggregatedByTargetResourceOutput
	p := inspector2.NewListCisScanResultsAggregatedByTargetResourcePaginator(client, input)
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

// Returns a CIS scan list.
func inspector2_ListCisScans(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.ListCisScansInput{}

	if len(_inspector2DetailLevel) > 0 {
		if err := assignInputField(input, "DetailLevel", _inspector2DetailLevel); err != nil {
			log.Errorf("invalid --detail-level: %s", err.Error())
			return
		}
	}
	if len(_inspector2FilterCriteria) > 0 {
		if err := assignInputField(input, "FilterCriteria", _inspector2FilterCriteria); err != nil {
			log.Errorf("invalid --filter-criteria: %s", err.Error())
			return
		}
	}
	if len(_inspector2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _inspector2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_inspector2NextToken) > 0 {
		input.NextToken = aws.String(_inspector2NextToken)
	}
	if len(_inspector2SortBy) > 0 {
		if err := assignInputField(input, "SortBy", _inspector2SortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_inspector2SortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _inspector2SortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListCisScans(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*inspector2.ListCisScansOutput
	p := inspector2.NewListCisScansPaginator(client, input)
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

// Lists all code security integrations in your account.
func inspector2_ListCodeSecurityIntegrations(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.ListCodeSecurityIntegrationsInput{}

	if len(_inspector2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _inspector2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_inspector2NextToken) > 0 {
		input.NextToken = aws.String(_inspector2NextToken)
	}

	if resp, err := client.ListCodeSecurityIntegrations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the associations between code repositories and Amazon Inspector code
// security scan configurations.
func inspector2_ListCodeSecurityScanConfigurationAssociations(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.ListCodeSecurityScanConfigurationAssociationsInput{
		// ScanConfigurationArn: *string, // Required
	}

	if len(_inspector2ScanConfigurationArn) > 0 {
		input.ScanConfigurationArn = aws.String(_inspector2ScanConfigurationArn)
	}
	if len(_inspector2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _inspector2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_inspector2NextToken) > 0 {
		input.NextToken = aws.String(_inspector2NextToken)
	}

	if resp, err := client.ListCodeSecurityScanConfigurationAssociations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all code security scan configurations in your account.
func inspector2_ListCodeSecurityScanConfigurations(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.ListCodeSecurityScanConfigurationsInput{}

	if len(_inspector2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _inspector2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_inspector2NextToken) > 0 {
		input.NextToken = aws.String(_inspector2NextToken)
	}

	if resp, err := client.ListCodeSecurityScanConfigurations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists coverage details for your environment.
func inspector2_ListCoverage(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.ListCoverageInput{}

	if len(_inspector2FilterCriteria) > 0 {
		if err := assignInputField(input, "FilterCriteria", _inspector2FilterCriteria); err != nil {
			log.Errorf("invalid --filter-criteria: %s", err.Error())
			return
		}
	}
	if len(_inspector2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _inspector2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_inspector2NextToken) > 0 {
		input.NextToken = aws.String(_inspector2NextToken)
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

	var results []*inspector2.ListCoverageOutput
	p := inspector2.NewListCoveragePaginator(client, input)
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

// Lists Amazon Inspector coverage statistics for your environment.
func inspector2_ListCoverageStatistics(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.ListCoverageStatisticsInput{}

	if len(_inspector2FilterCriteria) > 0 {
		if err := assignInputField(input, "FilterCriteria", _inspector2FilterCriteria); err != nil {
			log.Errorf("invalid --filter-criteria: %s", err.Error())
			return
		}
	}
	if len(_inspector2GroupBy) > 0 {
		if err := assignInputField(input, "GroupBy", _inspector2GroupBy); err != nil {
			log.Errorf("invalid --group-by: %s", err.Error())
			return
		}
	}
	if len(_inspector2NextToken) > 0 {
		input.NextToken = aws.String(_inspector2NextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCoverageStatistics(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*inspector2.ListCoverageStatisticsOutput
	p := inspector2.NewListCoverageStatisticsPaginator(client, input)
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

// Lists information about the Amazon Inspector delegated administrator of your
// organization.
func inspector2_ListDelegatedAdminAccounts(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.ListDelegatedAdminAccountsInput{}

	if len(_inspector2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _inspector2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_inspector2NextToken) > 0 {
		input.NextToken = aws.String(_inspector2NextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDelegatedAdminAccounts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*inspector2.ListDelegatedAdminAccountsOutput
	p := inspector2.NewListDelegatedAdminAccountsPaginator(client, input)
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

// Lists the filters associated with your account.
func inspector2_ListFilters(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.ListFiltersInput{}

	if len(_inspector2Action) > 0 {
		if err := assignInputField(input, "Action", _inspector2Action); err != nil {
			log.Errorf("invalid --action: %s", err.Error())
			return
		}
	}
	if len(_inspector2Arns) > 0 {
		input.Arns = append([]string(nil), _inspector2Arns...)
	}
	if len(_inspector2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _inspector2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_inspector2NextToken) > 0 {
		input.NextToken = aws.String(_inspector2NextToken)
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

	var results []*inspector2.ListFiltersOutput
	p := inspector2.NewListFiltersPaginator(client, input)
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

// Lists aggregated finding data for your environment based on specific criteria.
func inspector2_ListFindingAggregations(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.ListFindingAggregationsInput{
		// AggregationType: types.AggregationType, // Required
	}

	if len(_inspector2AggregationType) > 0 {
		if err := assignInputField(input, "AggregationType", _inspector2AggregationType); err != nil {
			log.Errorf("invalid --aggregation-type: %s", err.Error())
			return
		}
	}
	if len(_inspector2AccountIds) > 0 {
		if err := assignInputField(input, "AccountIds", _inspector2AccountIds[0]); err != nil {
			log.Errorf("invalid --account-ids: %s", err.Error())
			return
		}
	}
	if len(_inspector2AggregationRequest) > 0 {
		if err := assignInputField(input, "AggregationRequest", _inspector2AggregationRequest); err != nil {
			log.Errorf("invalid --aggregation-request: %s", err.Error())
			return
		}
	}
	if len(_inspector2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _inspector2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_inspector2NextToken) > 0 {
		input.NextToken = aws.String(_inspector2NextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListFindingAggregations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*inspector2.ListFindingAggregationsOutput
	p := inspector2.NewListFindingAggregationsPaginator(client, input)
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

// Lists findings for your environment.
func inspector2_ListFindings(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.ListFindingsInput{}

	if len(_inspector2FilterCriteria) > 0 {
		if err := assignInputField(input, "FilterCriteria", _inspector2FilterCriteria); err != nil {
			log.Errorf("invalid --filter-criteria: %s", err.Error())
			return
		}
	}
	if len(_inspector2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _inspector2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_inspector2NextToken) > 0 {
		input.NextToken = aws.String(_inspector2NextToken)
	}
	if len(_inspector2SortCriteria) > 0 {
		if err := assignInputField(input, "SortCriteria", _inspector2SortCriteria); err != nil {
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

	var results []*inspector2.ListFindingsOutput
	p := inspector2.NewListFindingsPaginator(client, input)
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

// List members associated with the Amazon Inspector delegated administrator for
// your organization.
func inspector2_ListMembers(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.ListMembersInput{}

	if len(_inspector2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _inspector2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_inspector2NextToken) > 0 {
		input.NextToken = aws.String(_inspector2NextToken)
	}
	if len(_inspector2OnlyAssociated) > 0 {
		if err := assignInputField(input, "OnlyAssociated", _inspector2OnlyAssociated); err != nil {
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

	var results []*inspector2.ListMembersOutput
	p := inspector2.NewListMembersPaginator(client, input)
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

// Lists all tags attached to a given resource.
func inspector2_ListTagsForResource(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_inspector2ResourceArn) > 0 {
		input.ResourceArn = aws.String(_inspector2ResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the Amazon Inspector usage totals over the last 30 days.
func inspector2_ListUsageTotals(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.ListUsageTotalsInput{}

	if len(_inspector2AccountIds) > 0 {
		input.AccountIds = append([]string(nil), _inspector2AccountIds...)
	}
	if len(_inspector2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _inspector2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_inspector2NextToken) > 0 {
		input.NextToken = aws.String(_inspector2NextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListUsageTotals(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*inspector2.ListUsageTotalsOutput
	p := inspector2.NewListUsageTotalsPaginator(client, input)
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

// Resets an encryption key. After the key is reset your resources will be
// encrypted by an Amazon Web Services owned key.
func inspector2_ResetEncryptionKey(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.ResetEncryptionKeyInput{
		// ResourceType: types.ResourceType, // Required
		// ScanType: types.ScanType, // Required
	}

	if len(_inspector2ResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _inspector2ResourceType); err != nil {
			log.Errorf("invalid --resource-type: %s", err.Error())
			return
		}
	}
	if len(_inspector2ScanType) > 0 {
		if err := assignInputField(input, "ScanType", _inspector2ScanType); err != nil {
			log.Errorf("invalid --scan-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.ResetEncryptionKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists Amazon Inspector coverage details for a specific vulnerability.
func inspector2_SearchVulnerabilities(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.SearchVulnerabilitiesInput{
		// FilterCriteria: *types.SearchVulnerabilitiesFilterCriteria, // Required
	}

	if len(_inspector2FilterCriteria) > 0 {
		if err := assignInputField(input, "FilterCriteria", _inspector2FilterCriteria); err != nil {
			log.Errorf("invalid --filter-criteria: %s", err.Error())
			return
		}
	}
	if len(_inspector2NextToken) > 0 {
		input.NextToken = aws.String(_inspector2NextToken)
	}

	if disablePaginator() {
		if resp, err := client.SearchVulnerabilities(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*inspector2.SearchVulnerabilitiesOutput
	p := inspector2.NewSearchVulnerabilitiesPaginator(client, input)
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

// Sends a CIS session health. This API is used by the Amazon Inspector SSM
// plugin to communicate with the Amazon Inspector service. The Amazon Inspector
// SSM plugin calls this API to start a CIS scan session for the scan ID supplied
// by the service.
func inspector2_SendCisSessionHealth(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.SendCisSessionHealthInput{
		// ScanJobId: *string, // Required
		// SessionToken: *string, // Required
	}

	if len(_inspector2ScanJobId) > 0 {
		input.ScanJobId = aws.String(_inspector2ScanJobId)
	}
	if len(_inspector2SessionToken) > 0 {
		input.SessionToken = aws.String(_inspector2SessionToken)
	}

	if resp, err := client.SendCisSessionHealth(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sends a CIS session telemetry. This API is used by the Amazon Inspector SSM
// plugin to communicate with the Amazon Inspector service. The Amazon Inspector
// SSM plugin calls this API to start a CIS scan session for the scan ID supplied
// by the service.
func inspector2_SendCisSessionTelemetry(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.SendCisSessionTelemetryInput{
		// Messages: []types.CisSessionMessage, // Required
		// ScanJobId: *string, // Required
		// SessionToken: *string, // Required
	}

	if len(_inspector2Messages) > 0 {
		if err := assignInputField(input, "Messages", _inspector2Messages); err != nil {
			log.Errorf("invalid --messages: %s", err.Error())
			return
		}
	}
	if len(_inspector2ScanJobId) > 0 {
		input.ScanJobId = aws.String(_inspector2ScanJobId)
	}
	if len(_inspector2SessionToken) > 0 {
		input.SessionToken = aws.String(_inspector2SessionToken)
	}

	if resp, err := client.SendCisSessionTelemetry(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a CIS session. This API is used by the Amazon Inspector SSM plugin to
// communicate with the Amazon Inspector service. The Amazon Inspector SSM plugin
// calls this API to start a CIS scan session for the scan ID supplied by the
// service.
func inspector2_StartCisSession(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.StartCisSessionInput{
		// Message: *types.StartCisSessionMessage, // Required
		// ScanJobId: *string, // Required
	}

	if len(_inspector2Message) > 0 {
		if err := assignInputField(input, "Message", _inspector2Message); err != nil {
			log.Errorf("invalid --message: %s", err.Error())
			return
		}
	}
	if len(_inspector2ScanJobId) > 0 {
		input.ScanJobId = aws.String(_inspector2ScanJobId)
	}

	if resp, err := client.StartCisSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Initiates a code security scan on a specified repository.
func inspector2_StartCodeSecurityScan(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.StartCodeSecurityScanInput{
		// Resource: types.CodeSecurityResource, // Required
	}

	if len(_inspector2Resource) > 0 {
		if err := assignInputField(input, "Resource", _inspector2Resource); err != nil {
			log.Errorf("invalid --resource: %s", err.Error())
			return
		}
	}
	if len(_inspector2ClientToken) > 0 {
		input.ClientToken = aws.String(_inspector2ClientToken)
	}

	if resp, err := client.StartCodeSecurityScan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a CIS session. This API is used by the Amazon Inspector SSM plugin to
// communicate with the Amazon Inspector service. The Amazon Inspector SSM plugin
// calls this API to stop a CIS scan session for the scan ID supplied by the
// service.
func inspector2_StopCisSession(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.StopCisSessionInput{
		// Message: *types.StopCisSessionMessage, // Required
		// ScanJobId: *string, // Required
		// SessionToken: *string, // Required
	}

	if len(_inspector2Message) > 0 {
		if err := assignInputField(input, "Message", _inspector2Message); err != nil {
			log.Errorf("invalid --message: %s", err.Error())
			return
		}
	}
	if len(_inspector2ScanJobId) > 0 {
		input.ScanJobId = aws.String(_inspector2ScanJobId)
	}
	if len(_inspector2SessionToken) > 0 {
		input.SessionToken = aws.String(_inspector2SessionToken)
	}

	if resp, err := client.StopCisSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds tags to a resource.
func inspector2_TagResource(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_inspector2ResourceArn) > 0 {
		input.ResourceArn = aws.String(_inspector2ResourceArn)
	}
	if len(_inspector2Tags) > 0 {
		if err := assignInputField(input, "Tags", _inspector2Tags); err != nil {
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

// Removes tags from a resource.
func inspector2_UntagResource(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_inspector2ResourceArn) > 0 {
		input.ResourceArn = aws.String(_inspector2ResourceArn)
	}
	if len(_inspector2TagKeys) > 0 {
		input.TagKeys = append([]string(nil), _inspector2TagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a CIS scan configuration.
func inspector2_UpdateCisScanConfiguration(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.UpdateCisScanConfigurationInput{
		// ScanConfigurationArn: *string, // Required
	}

	if len(_inspector2ScanConfigurationArn) > 0 {
		input.ScanConfigurationArn = aws.String(_inspector2ScanConfigurationArn)
	}
	if len(_inspector2ScanName) > 0 {
		input.ScanName = aws.String(_inspector2ScanName)
	}
	if len(_inspector2Schedule) > 0 {
		if err := assignInputField(input, "Schedule", _inspector2Schedule); err != nil {
			log.Errorf("invalid --schedule: %s", err.Error())
			return
		}
	}
	if len(_inspector2SecurityLevel) > 0 {
		if err := assignInputField(input, "SecurityLevel", _inspector2SecurityLevel); err != nil {
			log.Errorf("invalid --security-level: %s", err.Error())
			return
		}
	}
	if len(_inspector2Targets) > 0 {
		if err := assignInputField(input, "Targets", _inspector2Targets); err != nil {
			log.Errorf("invalid --targets: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateCisScanConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing code security integration.
// After calling the CreateCodeSecurityIntegration operation, you complete
// authentication and authorization with your provider. Next you call the
// UpdateCodeSecurityIntegration operation to provide the details to complete the
// integration setup
func inspector2_UpdateCodeSecurityIntegration(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.UpdateCodeSecurityIntegrationInput{
		// Details: types.UpdateIntegrationDetails, // Required
		// IntegrationArn: *string, // Required
	}

	if len(_inspector2Details) > 0 {
		if err := assignInputField(input, "Details", _inspector2Details); err != nil {
			log.Errorf("invalid --details: %s", err.Error())
			return
		}
	}
	if len(_inspector2IntegrationArn) > 0 {
		input.IntegrationArn = aws.String(_inspector2IntegrationArn)
	}

	if resp, err := client.UpdateCodeSecurityIntegration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing code security scan configuration.
func inspector2_UpdateCodeSecurityScanConfiguration(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.UpdateCodeSecurityScanConfigurationInput{
		// Configuration: *types.CodeSecurityScanConfiguration, // Required
		// ScanConfigurationArn: *string, // Required
	}

	if len(_inspector2Configuration) > 0 {
		if err := assignInputField(input, "Configuration", _inspector2Configuration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}
	if len(_inspector2ScanConfigurationArn) > 0 {
		input.ScanConfigurationArn = aws.String(_inspector2ScanConfigurationArn)
	}

	if resp, err := client.UpdateCodeSecurityScanConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates setting configurations for your Amazon Inspector account. When you use
// this API as an Amazon Inspector delegated administrator this updates the setting
// for all accounts you manage. Member accounts in an organization cannot update
// this setting.
func inspector2_UpdateConfiguration(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.UpdateConfigurationInput{}

	if len(_inspector2Ec2Configuration) > 0 {
		if err := assignInputField(input, "Ec2Configuration", _inspector2Ec2Configuration); err != nil {
			log.Errorf("invalid --ec2-configuration: %s", err.Error())
			return
		}
	}
	if len(_inspector2EcrConfiguration) > 0 {
		if err := assignInputField(input, "EcrConfiguration", _inspector2EcrConfiguration); err != nil {
			log.Errorf("invalid --ecr-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Activates, deactivates Amazon Inspector deep inspection, or updates custom
// paths for your account.
func inspector2_UpdateEc2DeepInspectionConfiguration(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.UpdateEc2DeepInspectionConfigurationInput{}

	if len(_inspector2ActivateDeepInspection) > 0 {
		if err := assignInputField(input, "ActivateDeepInspection", _inspector2ActivateDeepInspection); err != nil {
			log.Errorf("invalid --activate-deep-inspection: %s", err.Error())
			return
		}
	}
	if len(_inspector2PackagePaths) > 0 {
		input.PackagePaths = append([]string(nil), _inspector2PackagePaths...)
	}

	if resp, err := client.UpdateEc2DeepInspectionConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an encryption key. A ResourceNotFoundException means that an Amazon Web
// Services owned key is being used for encryption.
func inspector2_UpdateEncryptionKey(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.UpdateEncryptionKeyInput{
		// KmsKeyId: *string, // Required
		// ResourceType: types.ResourceType, // Required
		// ScanType: types.ScanType, // Required
	}

	if len(_inspector2KmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_inspector2KmsKeyId)
	}
	if len(_inspector2ResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _inspector2ResourceType); err != nil {
			log.Errorf("invalid --resource-type: %s", err.Error())
			return
		}
	}
	if len(_inspector2ScanType) > 0 {
		if err := assignInputField(input, "ScanType", _inspector2ScanType); err != nil {
			log.Errorf("invalid --scan-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateEncryptionKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Specifies the action that is to be applied to the findings that match the
// filter.
func inspector2_UpdateFilter(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.UpdateFilterInput{
		// FilterArn: *string, // Required
	}

	if len(_inspector2FilterArn) > 0 {
		input.FilterArn = aws.String(_inspector2FilterArn)
	}
	if len(_inspector2Action) > 0 {
		if err := assignInputField(input, "Action", _inspector2Action); err != nil {
			log.Errorf("invalid --action: %s", err.Error())
			return
		}
	}
	if len(_inspector2Description) > 0 {
		input.Description = aws.String(_inspector2Description)
	}
	if len(_inspector2FilterCriteria) > 0 {
		if err := assignInputField(input, "FilterCriteria", _inspector2FilterCriteria); err != nil {
			log.Errorf("invalid --filter-criteria: %s", err.Error())
			return
		}
	}
	if len(_inspector2Name) > 0 {
		input.Name = aws.String(_inspector2Name)
	}
	if len(_inspector2Reason) > 0 {
		input.Reason = aws.String(_inspector2Reason)
	}

	if resp, err := client.UpdateFilter(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the Amazon Inspector deep inspection custom paths for your
// organization. You must be an Amazon Inspector delegated administrator to use
// this API.
func inspector2_UpdateOrgEc2DeepInspectionConfiguration(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.UpdateOrgEc2DeepInspectionConfigurationInput{
		// OrgPackagePaths: []string, // Required
	}

	if len(_inspector2OrgPackagePaths) > 0 {
		input.OrgPackagePaths = append([]string(nil), _inspector2OrgPackagePaths...)
	}

	if resp, err := client.UpdateOrgEc2DeepInspectionConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configurations for your Amazon Inspector organization.
func inspector2_UpdateOrganizationConfiguration(cfg aws.Config, client *inspector2.Client) {
	input := &inspector2.UpdateOrganizationConfigurationInput{
		// AutoEnable: *types.AutoEnable, // Required
	}

	if len(_inspector2AutoEnable) > 0 {
		if err := assignInputField(input, "AutoEnable", _inspector2AutoEnable); err != nil {
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

func init() {
	_rootCmd.AddCommand(_inspector2Cmd)
	_inspector2Cmd.Flags().SortFlags = false

	_inspector2Cmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_inspector2Cmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_inspector2Cmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_inspector2Cmd.Flags().StringVarP(&_inspector2AccountId, "account-id", "", "", "Account ID")
	_inspector2Cmd.Flags().StringSliceVarP(&_inspector2AccountIds, "account-ids", "", nil, "Account Ids")
	_inspector2Cmd.Flags().StringVarP(&_inspector2Action, "action", "", "", "Action")
	_inspector2Cmd.Flags().StringVarP(&_inspector2ActivateDeepInspection, "activate-deep-inspection", "", "", "Activate Deep Inspection")
	_inspector2Cmd.Flags().StringVarP(&_inspector2AggregationRequest, "aggregation-request", "", "", "Aggregation Request")
	_inspector2Cmd.Flags().StringVarP(&_inspector2AggregationType, "aggregation-type", "", "", "Aggregation Type")
	_inspector2Cmd.Flags().StringVarP(&_inspector2Arn, "arn", "", "", "ARN")
	_inspector2Cmd.Flags().StringSliceVarP(&_inspector2Arns, "arns", "", nil, "Arns")
	_inspector2Cmd.Flags().StringVarP(&_inspector2AssociateConfigurationRequests, "associate-configuration-requests", "", "", "Associate Configuration Requests")
	_inspector2Cmd.Flags().StringVarP(&_inspector2AutoEnable, "auto-enable", "", "", "Auto Enable")
	_inspector2Cmd.Flags().StringVarP(&_inspector2ClientToken, "client-token", "", "", "Client Token")
	_inspector2Cmd.Flags().StringVarP(&_inspector2Configuration, "configuration", "", "", "Configuration")
	_inspector2Cmd.Flags().StringVarP(&_inspector2DelegatedAdminAccountId, "delegated-admin-account-id", "", "", "Delegated Admin Account ID")
	_inspector2Cmd.Flags().StringVarP(&_inspector2Description, "description", "", "", "Description")
	_inspector2Cmd.Flags().StringVarP(&_inspector2DetailLevel, "detail-level", "", "", "Detail Level")
	_inspector2Cmd.Flags().StringVarP(&_inspector2Details, "details", "", "", "Details")
	_inspector2Cmd.Flags().StringVarP(&_inspector2DisassociateConfigurationRequests, "disassociate-configuration-requests", "", "", "Disassociate Configuration Requests")
	_inspector2Cmd.Flags().StringVarP(&_inspector2Ec2Configuration, "ec2-configuration", "", "", "EC2 Configuration")
	_inspector2Cmd.Flags().StringVarP(&_inspector2EcrConfiguration, "ecr-configuration", "", "", "ECR Configuration")
	_inspector2Cmd.Flags().StringVarP(&_inspector2Filter, "filter", "", "", "Filter")
	_inspector2Cmd.Flags().StringVarP(&_inspector2FilterArn, "filter-arn", "", "", "Filter ARN")
	_inspector2Cmd.Flags().StringVarP(&_inspector2FilterCriteria, "filter-criteria", "", "", "Filter Criteria")
	_inspector2Cmd.Flags().StringSliceVarP(&_inspector2FindingArns, "finding-arns", "", nil, "Finding Arns")
	_inspector2Cmd.Flags().StringVarP(&_inspector2GroupBy, "group-by", "", "", "Group By")
	_inspector2Cmd.Flags().StringVarP(&_inspector2IntegrationArn, "integration-arn", "", "", "Integration ARN")
	_inspector2Cmd.Flags().StringVarP(&_inspector2KmsKeyId, "kms-key-id", "", "", "KMS Key ID")
	_inspector2Cmd.Flags().StringVarP(&_inspector2Level, "level", "", "", "Level")
	_inspector2Cmd.Flags().StringVarP(&_inspector2MaxResults, "max-results", "", "", "Max Results")
	_inspector2Cmd.Flags().StringVarP(&_inspector2Message, "message", "", "", "Message")
	_inspector2Cmd.Flags().StringVarP(&_inspector2Messages, "messages", "", "", "Messages")
	_inspector2Cmd.Flags().StringVarP(&_inspector2Name, "name", "", "", "Name")
	_inspector2Cmd.Flags().StringVarP(&_inspector2NextToken, "next-token", "", "", "Next Token")
	_inspector2Cmd.Flags().StringVarP(&_inspector2OnlyAssociated, "only-associated", "", "", "Only Associated")
	_inspector2Cmd.Flags().StringSliceVarP(&_inspector2OrgPackagePaths, "org-package-paths", "", nil, "Org Package Paths")
	_inspector2Cmd.Flags().StringSliceVarP(&_inspector2PackagePaths, "package-paths", "", nil, "Package Paths")
	_inspector2Cmd.Flags().StringVarP(&_inspector2Reason, "reason", "", "", "Reason")
	_inspector2Cmd.Flags().StringVarP(&_inspector2ReportFormat, "report-format", "", "", "Report Format")
	_inspector2Cmd.Flags().StringVarP(&_inspector2ReportId, "report-id", "", "", "Report ID")
	_inspector2Cmd.Flags().StringVarP(&_inspector2Resource, "resource", "", "", "Resource")
	_inspector2Cmd.Flags().StringVarP(&_inspector2ResourceArn, "resource-arn", "", "", "Resource ARN")
	_inspector2Cmd.Flags().StringVarP(&_inspector2ResourceFilterCriteria, "resource-filter-criteria", "", "", "Resource Filter Criteria")
	_inspector2Cmd.Flags().StringVarP(&_inspector2ResourceType, "resource-type", "", "", "Resource Type")
	_inspector2Cmd.Flags().StringVarP(&_inspector2ResourceTypes, "resource-types", "", "", "Resource Types")
	_inspector2Cmd.Flags().StringVarP(&_inspector2S3Destination, "s3-destination", "", "", "S3 Destination")
	_inspector2Cmd.Flags().StringVarP(&_inspector2ScanArn, "scan-arn", "", "", "Scan ARN")
	_inspector2Cmd.Flags().StringVarP(&_inspector2ScanConfigurationArn, "scan-configuration-arn", "", "", "Scan Configuration ARN")
	_inspector2Cmd.Flags().StringVarP(&_inspector2ScanId, "scan-id", "", "", "Scan ID")
	_inspector2Cmd.Flags().StringVarP(&_inspector2ScanJobId, "scan-job-id", "", "", "Scan Job ID")
	_inspector2Cmd.Flags().StringVarP(&_inspector2ScanName, "scan-name", "", "", "Scan Name")
	_inspector2Cmd.Flags().StringVarP(&_inspector2ScanType, "scan-type", "", "", "Scan Type")
	_inspector2Cmd.Flags().StringVarP(&_inspector2Schedule, "schedule", "", "", "Schedule")
	_inspector2Cmd.Flags().StringVarP(&_inspector2ScopeSettings, "scope-settings", "", "", "Scope Settings")
	_inspector2Cmd.Flags().StringVarP(&_inspector2SecurityLevel, "security-level", "", "", "Security Level")
	_inspector2Cmd.Flags().StringVarP(&_inspector2Service, "service", "", "", "Service")
	_inspector2Cmd.Flags().StringVarP(&_inspector2SessionToken, "session-token", "", "", "Session Token")
	_inspector2Cmd.Flags().StringVarP(&_inspector2SortBy, "sort-by", "", "", "Sort By")
	_inspector2Cmd.Flags().StringVarP(&_inspector2SortCriteria, "sort-criteria", "", "", "Sort Criteria")
	_inspector2Cmd.Flags().StringVarP(&_inspector2SortOrder, "sort-order", "", "", "Sort Order")
	_inspector2Cmd.Flags().StringSliceVarP(&_inspector2TagKeys, "tag-keys", "", nil, "Tag Keys")
	_inspector2Cmd.Flags().StringVarP(&_inspector2Tags, "tags", "", "", "Tags")
	_inspector2Cmd.Flags().StringSliceVarP(&_inspector2TargetAccounts, "target-accounts", "", nil, "Target Accounts")
	_inspector2Cmd.Flags().StringVarP(&_inspector2TargetResourceId, "target-resource-id", "", "", "Target Resource ID")
	_inspector2Cmd.Flags().StringVarP(&_inspector2Targets, "targets", "", "", "Targets")
	_inspector2Cmd.Flags().StringVarP(&_inspector2Type, "type", "", "", "Type")

	_inspector2Cmd.Flags().BoolVarP(&_inspector2AssociateMember, "associate-member", "", false, "Associate Member")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2BatchAssociateCodeSecurityScanConfiguration, "batch-associate-code-security-scan-configuration", "", false, "Batch Associate Code Security Scan Configuration")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2BatchDisassociateCodeSecurityScanConfiguration, "batch-disassociate-code-security-scan-configuration", "", false, "Batch Disassociate Code Security Scan Configuration")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2BatchGetAccountStatus, "batch-get-account-status", "", false, "Batch Get Account Status")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2BatchGetCodeSnippet, "batch-get-code-snippet", "", false, "Batch Get Code Snippet")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2BatchGetFindingDetails, "batch-get-finding-details", "", false, "Batch Get Finding Details")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2BatchGetFreeTrialInfo, "batch-get-free-trial-info", "", false, "Batch Get Free Trial Info")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2BatchGetMemberEc2DeepInspectionStatus, "batch-get-member-ec2-deep-inspection-status", "", false, "Batch Get Member EC2 Deep Inspection Status")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2BatchUpdateMemberEc2DeepInspectionStatus, "batch-update-member-ec2-deep-inspection-status", "", false, "Batch Update Member EC2 Deep Inspection Status")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2CancelFindingsReport, "cancel-findings-report", "", false, "Cancel Findings Report")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2CancelSbomExport, "cancel-sbom-export", "", false, "Cancel Sbom Export")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2CreateCisScanConfiguration, "create-cis-scan-configuration", "", false, "Create Cis Scan Configuration")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2CreateCodeSecurityIntegration, "create-code-security-integration", "", false, "Create Code Security Integration")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2CreateCodeSecurityScanConfiguration, "create-code-security-scan-configuration", "", false, "Create Code Security Scan Configuration")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2CreateFilter, "create-filter", "", false, "Create Filter")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2CreateFindingsReport, "create-findings-report", "", false, "Create Findings Report")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2CreateSbomExport, "create-sbom-export", "", false, "Create Sbom Export")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2DeleteCisScanConfiguration, "delete-cis-scan-configuration", "", false, "Delete Cis Scan Configuration")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2DeleteCodeSecurityIntegration, "delete-code-security-integration", "", false, "Delete Code Security Integration")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2DeleteCodeSecurityScanConfiguration, "delete-code-security-scan-configuration", "", false, "Delete Code Security Scan Configuration")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2DeleteFilter, "delete-filter", "", false, "Delete Filter")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2DescribeOrganizationConfiguration, "describe-organization-configuration", "", false, "Describe Organization Configuration")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2Disable, "disable", "", false, "Disable")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2DisableDelegatedAdminAccount, "disable-delegated-admin-account", "", false, "Disable Delegated Admin Account")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2DisassociateMember, "disassociate-member", "", false, "Disassociate Member")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2Enable, "enable", "", false, "Enable")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2EnableDelegatedAdminAccount, "enable-delegated-admin-account", "", false, "Enable Delegated Admin Account")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2GetCisScanReport, "get-cis-scan-report", "", false, "Get Cis Scan Report")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2GetCisScanResultDetails, "get-cis-scan-result-details", "", false, "Get Cis Scan Result Details")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2GetClustersForImage, "get-clusters-for-image", "", false, "Get Clusters For Image")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2GetCodeSecurityIntegration, "get-code-security-integration", "", false, "Get Code Security Integration")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2GetCodeSecurityScan, "get-code-security-scan", "", false, "Get Code Security Scan")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2GetCodeSecurityScanConfiguration, "get-code-security-scan-configuration", "", false, "Get Code Security Scan Configuration")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2GetConfiguration, "get-configuration", "", false, "Get Configuration")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2GetDelegatedAdminAccount, "get-delegated-admin-account", "", false, "Get Delegated Admin Account")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2GetEc2DeepInspectionConfiguration, "get-ec2-deep-inspection-configuration", "", false, "Get EC2 Deep Inspection Configuration")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2GetEncryptionKey, "get-encryption-key", "", false, "Get Encryption Key")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2GetFindingsReportStatus, "get-findings-report-status", "", false, "Get Findings Report Status")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2GetMember, "get-member", "", false, "Get Member")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2GetSbomExport, "get-sbom-export", "", false, "Get Sbom Export")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2ListAccountPermissions, "list-account-permissions", "", false, "List Account Permissions")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2ListCisScanConfigurations, "list-cis-scan-configurations", "", false, "List Cis Scan Configurations")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2ListCisScanResultsAggregatedByChecks, "list-cis-scan-results-aggregated-by-checks", "", false, "List Cis Scan Results Aggregated By Checks")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2ListCisScanResultsAggregatedByTargetResource, "list-cis-scan-results-aggregated-by-target-resource", "", false, "List Cis Scan Results Aggregated By Target Resource")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2ListCisScans, "list-cis-scans", "", false, "List Cis Scans")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2ListCodeSecurityIntegrations, "list-code-security-integrations", "", false, "List Code Security Integrations")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2ListCodeSecurityScanConfigurationAssociations, "list-code-security-scan-configuration-associations", "", false, "List Code Security Scan Configuration Associations")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2ListCodeSecurityScanConfigurations, "list-code-security-scan-configurations", "", false, "List Code Security Scan Configurations")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2ListCoverage, "list-coverage", "", false, "List Coverage")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2ListCoverageStatistics, "list-coverage-statistics", "", false, "List Coverage Statistics")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2ListDelegatedAdminAccounts, "list-delegated-admin-accounts", "", false, "List Delegated Admin Accounts")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2ListFilters, "list-filters", "", false, "List Filters")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2ListFindingAggregations, "list-finding-aggregations", "", false, "List Finding Aggregations")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2ListFindings, "list-findings", "", false, "List Findings")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2ListMembers, "list-members", "", false, "List Members")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2ListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2ListUsageTotals, "list-usage-totals", "", false, "List Usage Totals")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2ResetEncryptionKey, "reset-encryption-key", "", false, "Reset Encryption Key")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2SearchVulnerabilities, "search-vulnerabilities", "", false, "Search Vulnerabilities")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2SendCisSessionHealth, "send-cis-session-health", "", false, "Send Cis Session Health")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2SendCisSessionTelemetry, "send-cis-session-telemetry", "", false, "Send Cis Session Telemetry")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2StartCisSession, "start-cis-session", "", false, "Start Cis Session")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2StartCodeSecurityScan, "start-code-security-scan", "", false, "Start Code Security Scan")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2StopCisSession, "stop-cis-session", "", false, "Stop Cis Session")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2TagResource, "tag-resource", "", false, "Tag Resource")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2UntagResource, "untag-resource", "", false, "Untag Resource")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2UpdateCisScanConfiguration, "update-cis-scan-configuration", "", false, "Update Cis Scan Configuration")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2UpdateCodeSecurityIntegration, "update-code-security-integration", "", false, "Update Code Security Integration")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2UpdateCodeSecurityScanConfiguration, "update-code-security-scan-configuration", "", false, "Update Code Security Scan Configuration")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2UpdateConfiguration, "update-configuration", "", false, "Update Configuration")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2UpdateEc2DeepInspectionConfiguration, "update-ec2-deep-inspection-configuration", "", false, "Update EC2 Deep Inspection Configuration")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2UpdateEncryptionKey, "update-encryption-key", "", false, "Update Encryption Key")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2UpdateFilter, "update-filter", "", false, "Update Filter")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2UpdateOrgEc2DeepInspectionConfiguration, "update-org-ec2-deep-inspection-configuration", "", false, "Update Org EC2 Deep Inspection Configuration")
	_inspector2Cmd.Flags().BoolVarP(&_inspector2UpdateOrganizationConfiguration, "update-organization-configuration", "", false, "Update Organization Configuration")

}
