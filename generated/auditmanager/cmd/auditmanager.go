package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/auditmanager"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// auditmanagerCmd represents the auditmanager command
var _auditmanagerCmd = &cobra.Command{
	Use:   "auditmanager",
	Short: "AWS auditmanager CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := auditmanager.NewFromConfig(cfg)
		if _auditmanagerAssociateAssessmentReportEvidenceFolder {
			auditmanager_AssociateAssessmentReportEvidenceFolder(cfg, client)
			return
		}
		if _auditmanagerBatchAssociateAssessmentReportEvidence {
			auditmanager_BatchAssociateAssessmentReportEvidence(cfg, client)
			return
		}
		if _auditmanagerBatchCreateDelegationByAssessment {
			auditmanager_BatchCreateDelegationByAssessment(cfg, client)
			return
		}
		if _auditmanagerBatchDeleteDelegationByAssessment {
			auditmanager_BatchDeleteDelegationByAssessment(cfg, client)
			return
		}
		if _auditmanagerBatchDisassociateAssessmentReportEvidence {
			auditmanager_BatchDisassociateAssessmentReportEvidence(cfg, client)
			return
		}
		if _auditmanagerBatchImportEvidenceToAssessmentControl {
			auditmanager_BatchImportEvidenceToAssessmentControl(cfg, client)
			return
		}
		if _auditmanagerCreateAssessment {
			auditmanager_CreateAssessment(cfg, client)
			return
		}
		if _auditmanagerCreateAssessmentFramework {
			auditmanager_CreateAssessmentFramework(cfg, client)
			return
		}
		if _auditmanagerCreateAssessmentReport {
			auditmanager_CreateAssessmentReport(cfg, client)
			return
		}
		if _auditmanagerCreateControl {
			auditmanager_CreateControl(cfg, client)
			return
		}
		if _auditmanagerDeleteAssessment {
			auditmanager_DeleteAssessment(cfg, client)
			return
		}
		if _auditmanagerDeleteAssessmentFramework {
			auditmanager_DeleteAssessmentFramework(cfg, client)
			return
		}
		if _auditmanagerDeleteAssessmentFrameworkShare {
			auditmanager_DeleteAssessmentFrameworkShare(cfg, client)
			return
		}
		if _auditmanagerDeleteAssessmentReport {
			auditmanager_DeleteAssessmentReport(cfg, client)
			return
		}
		if _auditmanagerDeleteControl {
			auditmanager_DeleteControl(cfg, client)
			return
		}
		if _auditmanagerDeregisterAccount {
			auditmanager_DeregisterAccount(cfg, client)
			return
		}
		if _auditmanagerDeregisterOrganizationAdminAccount {
			auditmanager_DeregisterOrganizationAdminAccount(cfg, client)
			return
		}
		if _auditmanagerDisassociateAssessmentReportEvidenceFolder {
			auditmanager_DisassociateAssessmentReportEvidenceFolder(cfg, client)
			return
		}
		if _auditmanagerGetAccountStatus {
			auditmanager_GetAccountStatus(cfg, client)
			return
		}
		if _auditmanagerGetAssessment {
			auditmanager_GetAssessment(cfg, client)
			return
		}
		if _auditmanagerGetAssessmentFramework {
			auditmanager_GetAssessmentFramework(cfg, client)
			return
		}
		if _auditmanagerGetAssessmentReportUrl {
			auditmanager_GetAssessmentReportUrl(cfg, client)
			return
		}
		if _auditmanagerGetChangeLogs {
			auditmanager_GetChangeLogs(cfg, client)
			return
		}
		if _auditmanagerGetControl {
			auditmanager_GetControl(cfg, client)
			return
		}
		if _auditmanagerGetDelegations {
			auditmanager_GetDelegations(cfg, client)
			return
		}
		if _auditmanagerGetEvidence {
			auditmanager_GetEvidence(cfg, client)
			return
		}
		if _auditmanagerGetEvidenceByEvidenceFolder {
			auditmanager_GetEvidenceByEvidenceFolder(cfg, client)
			return
		}
		if _auditmanagerGetEvidenceFileUploadUrl {
			auditmanager_GetEvidenceFileUploadUrl(cfg, client)
			return
		}
		if _auditmanagerGetEvidenceFolder {
			auditmanager_GetEvidenceFolder(cfg, client)
			return
		}
		if _auditmanagerGetEvidenceFoldersByAssessment {
			auditmanager_GetEvidenceFoldersByAssessment(cfg, client)
			return
		}
		if _auditmanagerGetEvidenceFoldersByAssessmentControl {
			auditmanager_GetEvidenceFoldersByAssessmentControl(cfg, client)
			return
		}
		if _auditmanagerGetInsights {
			auditmanager_GetInsights(cfg, client)
			return
		}
		if _auditmanagerGetInsightsByAssessment {
			auditmanager_GetInsightsByAssessment(cfg, client)
			return
		}
		if _auditmanagerGetOrganizationAdminAccount {
			auditmanager_GetOrganizationAdminAccount(cfg, client)
			return
		}
		if _auditmanagerGetServicesInScope {
			auditmanager_GetServicesInScope(cfg, client)
			return
		}
		if _auditmanagerGetSettings {
			auditmanager_GetSettings(cfg, client)
			return
		}
		if _auditmanagerListAssessmentControlInsightsByControlDomain {
			auditmanager_ListAssessmentControlInsightsByControlDomain(cfg, client)
			return
		}
		if _auditmanagerListAssessmentFrameworkShareRequests {
			auditmanager_ListAssessmentFrameworkShareRequests(cfg, client)
			return
		}
		if _auditmanagerListAssessmentFrameworks {
			auditmanager_ListAssessmentFrameworks(cfg, client)
			return
		}
		if _auditmanagerListAssessmentReports {
			auditmanager_ListAssessmentReports(cfg, client)
			return
		}
		if _auditmanagerListAssessments {
			auditmanager_ListAssessments(cfg, client)
			return
		}
		if _auditmanagerListControlDomainInsights {
			auditmanager_ListControlDomainInsights(cfg, client)
			return
		}
		if _auditmanagerListControlDomainInsightsByAssessment {
			auditmanager_ListControlDomainInsightsByAssessment(cfg, client)
			return
		}
		if _auditmanagerListControlInsightsByControlDomain {
			auditmanager_ListControlInsightsByControlDomain(cfg, client)
			return
		}
		if _auditmanagerListControls {
			auditmanager_ListControls(cfg, client)
			return
		}
		if _auditmanagerListKeywordsForDataSource {
			auditmanager_ListKeywordsForDataSource(cfg, client)
			return
		}
		if _auditmanagerListNotifications {
			auditmanager_ListNotifications(cfg, client)
			return
		}
		if _auditmanagerListTagsForResource {
			auditmanager_ListTagsForResource(cfg, client)
			return
		}
		if _auditmanagerRegisterAccount {
			auditmanager_RegisterAccount(cfg, client)
			return
		}
		if _auditmanagerRegisterOrganizationAdminAccount {
			auditmanager_RegisterOrganizationAdminAccount(cfg, client)
			return
		}
		if _auditmanagerStartAssessmentFrameworkShare {
			auditmanager_StartAssessmentFrameworkShare(cfg, client)
			return
		}
		if _auditmanagerTagResource {
			auditmanager_TagResource(cfg, client)
			return
		}
		if _auditmanagerUntagResource {
			auditmanager_UntagResource(cfg, client)
			return
		}
		if _auditmanagerUpdateAssessment {
			auditmanager_UpdateAssessment(cfg, client)
			return
		}
		if _auditmanagerUpdateAssessmentControl {
			auditmanager_UpdateAssessmentControl(cfg, client)
			return
		}
		if _auditmanagerUpdateAssessmentControlSetStatus {
			auditmanager_UpdateAssessmentControlSetStatus(cfg, client)
			return
		}
		if _auditmanagerUpdateAssessmentFramework {
			auditmanager_UpdateAssessmentFramework(cfg, client)
			return
		}
		if _auditmanagerUpdateAssessmentFrameworkShare {
			auditmanager_UpdateAssessmentFrameworkShare(cfg, client)
			return
		}
		if _auditmanagerUpdateAssessmentStatus {
			auditmanager_UpdateAssessmentStatus(cfg, client)
			return
		}
		if _auditmanagerUpdateControl {
			auditmanager_UpdateControl(cfg, client)
			return
		}
		if _auditmanagerUpdateSettings {
			auditmanager_UpdateSettings(cfg, client)
			return
		}
		if _auditmanagerValidateAssessmentReportIntegrity {
			auditmanager_ValidateAssessmentReportIntegrity(cfg, client)
			return
		}

	},
}

var (
	_auditmanagerAssociateAssessmentReportEvidenceFolder      bool
	_auditmanagerBatchAssociateAssessmentReportEvidence       bool
	_auditmanagerBatchCreateDelegationByAssessment            bool
	_auditmanagerBatchDeleteDelegationByAssessment            bool
	_auditmanagerBatchDisassociateAssessmentReportEvidence    bool
	_auditmanagerBatchImportEvidenceToAssessmentControl       bool
	_auditmanagerCreateAssessment                             bool
	_auditmanagerCreateAssessmentFramework                    bool
	_auditmanagerCreateAssessmentReport                       bool
	_auditmanagerCreateControl                                bool
	_auditmanagerDeleteAssessment                             bool
	_auditmanagerDeleteAssessmentFramework                    bool
	_auditmanagerDeleteAssessmentFrameworkShare               bool
	_auditmanagerDeleteAssessmentReport                       bool
	_auditmanagerDeleteControl                                bool
	_auditmanagerDeregisterAccount                            bool
	_auditmanagerDeregisterOrganizationAdminAccount           bool
	_auditmanagerDisassociateAssessmentReportEvidenceFolder   bool
	_auditmanagerGetAccountStatus                             bool
	_auditmanagerGetAssessment                                bool
	_auditmanagerGetAssessmentFramework                       bool
	_auditmanagerGetAssessmentReportUrl                       bool
	_auditmanagerGetChangeLogs                                bool
	_auditmanagerGetControl                                   bool
	_auditmanagerGetDelegations                               bool
	_auditmanagerGetEvidence                                  bool
	_auditmanagerGetEvidenceByEvidenceFolder                  bool
	_auditmanagerGetEvidenceFileUploadUrl                     bool
	_auditmanagerGetEvidenceFolder                            bool
	_auditmanagerGetEvidenceFoldersByAssessment               bool
	_auditmanagerGetEvidenceFoldersByAssessmentControl        bool
	_auditmanagerGetInsights                                  bool
	_auditmanagerGetInsightsByAssessment                      bool
	_auditmanagerGetOrganizationAdminAccount                  bool
	_auditmanagerGetServicesInScope                           bool
	_auditmanagerGetSettings                                  bool
	_auditmanagerListAssessmentControlInsightsByControlDomain bool
	_auditmanagerListAssessmentFrameworkShareRequests         bool
	_auditmanagerListAssessmentFrameworks                     bool
	_auditmanagerListAssessmentReports                        bool
	_auditmanagerListAssessments                              bool
	_auditmanagerListControlDomainInsights                    bool
	_auditmanagerListControlDomainInsightsByAssessment        bool
	_auditmanagerListControlInsightsByControlDomain           bool
	_auditmanagerListControls                                 bool
	_auditmanagerListKeywordsForDataSource                    bool
	_auditmanagerListNotifications                            bool
	_auditmanagerListTagsForResource                          bool
	_auditmanagerRegisterAccount                              bool
	_auditmanagerRegisterOrganizationAdminAccount             bool
	_auditmanagerStartAssessmentFrameworkShare                bool
	_auditmanagerTagResource                                  bool
	_auditmanagerUntagResource                                bool
	_auditmanagerUpdateAssessment                             bool
	_auditmanagerUpdateAssessmentControl                      bool
	_auditmanagerUpdateAssessmentControlSetStatus             bool
	_auditmanagerUpdateAssessmentFramework                    bool
	_auditmanagerUpdateAssessmentFrameworkShare               bool
	_auditmanagerUpdateAssessmentStatus                       bool
	_auditmanagerUpdateControl                                bool
	_auditmanagerUpdateSettings                               bool
	_auditmanagerValidateAssessmentReportIntegrity            bool

	_auditmanagerAction                              string
	_auditmanagerActionPlanInstructions              string
	_auditmanagerActionPlanTitle                     string
	_auditmanagerAdminAccountId                      string
	_auditmanagerAssessmentDescription               string
	_auditmanagerAssessmentId                        string
	_auditmanagerAssessmentName                      string
	_auditmanagerAssessmentReportId                  string
	_auditmanagerAssessmentReportsDestination        string
	_auditmanagerAttribute                           string
	_auditmanagerComment                             string
	_auditmanagerCommentBody                         string
	_auditmanagerComplianceType                      string
	_auditmanagerControlCatalogId                    string
	_auditmanagerControlDomainId                     string
	_auditmanagerControlId                           string
	_auditmanagerControlMappingSources               string
	_auditmanagerControlSetId                        string
	_auditmanagerControlSets                         string
	_auditmanagerControlStatus                       string
	_auditmanagerControlType                         string
	_auditmanagerCreateDelegationRequests            string
	_auditmanagerDefaultAssessmentReportsDestination string
	_auditmanagerDefaultExportDestination            string
	_auditmanagerDefaultProcessOwners                string
	_auditmanagerDelegatedAdminAccount               string
	_auditmanagerDelegationIds                       []string
	_auditmanagerDeregistrationPolicy                string
	_auditmanagerDescription                         string
	_auditmanagerDestinationAccount                  string
	_auditmanagerDestinationRegion                   string
	_auditmanagerEvidenceFinderEnabled               string
	_auditmanagerEvidenceFolderId                    string
	_auditmanagerEvidenceId                          string
	_auditmanagerEvidenceIds                         []string
	_auditmanagerFileName                            string
	_auditmanagerFrameworkId                         string
	_auditmanagerFrameworkType                       string
	_auditmanagerKmsKey                              string
	_auditmanagerManualEvidence                      string
	_auditmanagerMaxResults                          string
	_auditmanagerName                                string
	_auditmanagerNextToken                           string
	_auditmanagerQueryStatement                      string
	_auditmanagerRequestId                           string
	_auditmanagerRequestType                         string
	_auditmanagerResourceArn                         string
	_auditmanagerRoles                               string
	_auditmanagerS3RelativePath                      string
	_auditmanagerScope                               string
	_auditmanagerSnsTopic                            string
	_auditmanagerSource                              string
	_auditmanagerStatus                              string
	_auditmanagerTagKeys                             []string
	_auditmanagerTags                                string
	_auditmanagerTestingInformation                  string
)

// Associates an evidence folder to an assessment report in an Audit Manager
// assessment.
func auditmanager_AssociateAssessmentReportEvidenceFolder(cfg aws.Config, client *auditmanager.Client) {
	input := &auditmanager.AssociateAssessmentReportEvidenceFolderInput{
		// AssessmentId: *string, // Required
		// EvidenceFolderId: *string, // Required
	}

	if len(_auditmanagerAssessmentId) > 0 {
		input.AssessmentId = aws.String(_auditmanagerAssessmentId)
	}
	if len(_auditmanagerEvidenceFolderId) > 0 {
		input.EvidenceFolderId = aws.String(_auditmanagerEvidenceFolderId)
	}

	if resp, err := client.AssociateAssessmentReportEvidenceFolder(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates a list of evidence to an assessment report in an Audit Manager
// assessment.
func auditmanager_BatchAssociateAssessmentReportEvidence(cfg aws.Config, client *auditmanager.Client) {
	input := &auditmanager.BatchAssociateAssessmentReportEvidenceInput{
		// AssessmentId: *string, // Required
		// EvidenceFolderId: *string, // Required
		// EvidenceIds: []string, // Required
	}

	if len(_auditmanagerAssessmentId) > 0 {
		input.AssessmentId = aws.String(_auditmanagerAssessmentId)
	}
	if len(_auditmanagerEvidenceFolderId) > 0 {
		input.EvidenceFolderId = aws.String(_auditmanagerEvidenceFolderId)
	}
	if len(_auditmanagerEvidenceIds) > 0 {
		input.EvidenceIds = append([]string(nil), _auditmanagerEvidenceIds...)
	}

	if resp, err := client.BatchAssociateAssessmentReportEvidence(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a batch of delegations for an assessment in Audit Manager.
func auditmanager_BatchCreateDelegationByAssessment(cfg aws.Config, client *auditmanager.Client) {
	input := &auditmanager.BatchCreateDelegationByAssessmentInput{
		// AssessmentId: *string, // Required
		// CreateDelegationRequests: []types.CreateDelegationRequest, // Required
	}

	if len(_auditmanagerAssessmentId) > 0 {
		input.AssessmentId = aws.String(_auditmanagerAssessmentId)
	}
	if len(_auditmanagerCreateDelegationRequests) > 0 {
		if err := assignInputField(input, "CreateDelegationRequests", _auditmanagerCreateDelegationRequests); err != nil {
			log.Errorf("invalid --create-delegation-requests: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchCreateDelegationByAssessment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a batch of delegations for an assessment in Audit Manager.
func auditmanager_BatchDeleteDelegationByAssessment(cfg aws.Config, client *auditmanager.Client) {
	input := &auditmanager.BatchDeleteDelegationByAssessmentInput{
		// AssessmentId: *string, // Required
		// DelegationIds: []string, // Required
	}

	if len(_auditmanagerAssessmentId) > 0 {
		input.AssessmentId = aws.String(_auditmanagerAssessmentId)
	}
	if len(_auditmanagerDelegationIds) > 0 {
		input.DelegationIds = append([]string(nil), _auditmanagerDelegationIds...)
	}

	if resp, err := client.BatchDeleteDelegationByAssessment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates a list of evidence from an assessment report in Audit Manager.
func auditmanager_BatchDisassociateAssessmentReportEvidence(cfg aws.Config, client *auditmanager.Client) {
	input := &auditmanager.BatchDisassociateAssessmentReportEvidenceInput{
		// AssessmentId: *string, // Required
		// EvidenceFolderId: *string, // Required
		// EvidenceIds: []string, // Required
	}

	if len(_auditmanagerAssessmentId) > 0 {
		input.AssessmentId = aws.String(_auditmanagerAssessmentId)
	}
	if len(_auditmanagerEvidenceFolderId) > 0 {
		input.EvidenceFolderId = aws.String(_auditmanagerEvidenceFolderId)
	}
	if len(_auditmanagerEvidenceIds) > 0 {
		input.EvidenceIds = append([]string(nil), _auditmanagerEvidenceIds...)
	}

	if resp, err := client.BatchDisassociateAssessmentReportEvidence(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds one or more pieces of evidence to a control in an Audit Manager
// assessment.
//
// You can import manual evidence from any S3 bucket by specifying the S3 URI of
// the object. You can also upload a file from your browser, or enter plain text in
// response to a risk assessment question.
//
// The following restrictions apply to this action:
//
// - manualEvidence can be only one of the following: evidenceFileName ,
// s3ResourcePath , or textResponse
//
// - Maximum size of an individual evidence file: 100 MB
//
// - Number of daily manual evidence uploads per control: 100
//
// - Supported file formats: See [Supported file types for manual evidence]in the Audit Manager User Guide
//
// For more information about Audit Manager service restrictions, see [Quotas and restrictions for Audit Manager].
//
// [Supported file types for manual evidence]: https://docs.aws.amazon.com/audit-manager/latest/userguide/upload-evidence.html#supported-manual-evidence-files
// [Quotas and restrictions for Audit Manager]: https://docs.aws.amazon.com/audit-manager/latest/userguide/service-quotas.html
func auditmanager_BatchImportEvidenceToAssessmentControl(cfg aws.Config, client *auditmanager.Client) {
	input := &auditmanager.BatchImportEvidenceToAssessmentControlInput{
		// AssessmentId: *string, // Required
		// ControlId: *string, // Required
		// ControlSetId: *string, // Required
		// ManualEvidence: []types.ManualEvidence, // Required
	}

	if len(_auditmanagerAssessmentId) > 0 {
		input.AssessmentId = aws.String(_auditmanagerAssessmentId)
	}
	if len(_auditmanagerControlId) > 0 {
		input.ControlId = aws.String(_auditmanagerControlId)
	}
	if len(_auditmanagerControlSetId) > 0 {
		input.ControlSetId = aws.String(_auditmanagerControlSetId)
	}
	if len(_auditmanagerManualEvidence) > 0 {
		if err := assignInputField(input, "ManualEvidence", _auditmanagerManualEvidence); err != nil {
			log.Errorf("invalid --manual-evidence: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchImportEvidenceToAssessmentControl(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an assessment in Audit Manager.
func auditmanager_CreateAssessment(cfg aws.Config, client *auditmanager.Client) {
	input := &auditmanager.CreateAssessmentInput{
		// AssessmentReportsDestination: *types.AssessmentReportsDestination, // Required
		// FrameworkId: *string, // Required
		// Name: *string, // Required
		// Roles: []types.Role, // Required
		// Scope: *types.Scope, // Required
	}

	if len(_auditmanagerAssessmentReportsDestination) > 0 {
		if err := assignInputField(input, "AssessmentReportsDestination", _auditmanagerAssessmentReportsDestination); err != nil {
			log.Errorf("invalid --assessment-reports-destination: %s", err.Error())
			return
		}
	}
	if len(_auditmanagerFrameworkId) > 0 {
		input.FrameworkId = aws.String(_auditmanagerFrameworkId)
	}
	if len(_auditmanagerName) > 0 {
		input.Name = aws.String(_auditmanagerName)
	}
	if len(_auditmanagerRoles) > 0 {
		if err := assignInputField(input, "Roles", _auditmanagerRoles); err != nil {
			log.Errorf("invalid --roles: %s", err.Error())
			return
		}
	}
	if len(_auditmanagerScope) > 0 {
		if err := assignInputField(input, "Scope", _auditmanagerScope); err != nil {
			log.Errorf("invalid --scope: %s", err.Error())
			return
		}
	}
	if len(_auditmanagerDescription) > 0 {
		input.Description = aws.String(_auditmanagerDescription)
	}
	if len(_auditmanagerTags) > 0 {
		if err := assignInputField(input, "Tags", _auditmanagerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAssessment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a custom framework in Audit Manager.
func auditmanager_CreateAssessmentFramework(cfg aws.Config, client *auditmanager.Client) {
	input := &auditmanager.CreateAssessmentFrameworkInput{
		// ControlSets: []types.CreateAssessmentFrameworkControlSet, // Required
		// Name: *string, // Required
	}

	if len(_auditmanagerControlSets) > 0 {
		if err := assignInputField(input, "ControlSets", _auditmanagerControlSets); err != nil {
			log.Errorf("invalid --control-sets: %s", err.Error())
			return
		}
	}
	if len(_auditmanagerName) > 0 {
		input.Name = aws.String(_auditmanagerName)
	}
	if len(_auditmanagerComplianceType) > 0 {
		input.ComplianceType = aws.String(_auditmanagerComplianceType)
	}
	if len(_auditmanagerDescription) > 0 {
		input.Description = aws.String(_auditmanagerDescription)
	}
	if len(_auditmanagerTags) > 0 {
		if err := assignInputField(input, "Tags", _auditmanagerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAssessmentFramework(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an assessment report for the specified assessment.
func auditmanager_CreateAssessmentReport(cfg aws.Config, client *auditmanager.Client) {
	input := &auditmanager.CreateAssessmentReportInput{
		// AssessmentId: *string, // Required
		// Name: *string, // Required
	}

	if len(_auditmanagerAssessmentId) > 0 {
		input.AssessmentId = aws.String(_auditmanagerAssessmentId)
	}
	if len(_auditmanagerName) > 0 {
		input.Name = aws.String(_auditmanagerName)
	}
	if len(_auditmanagerDescription) > 0 {
		input.Description = aws.String(_auditmanagerDescription)
	}
	if len(_auditmanagerQueryStatement) > 0 {
		input.QueryStatement = aws.String(_auditmanagerQueryStatement)
	}

	if resp, err := client.CreateAssessmentReport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new custom control in Audit Manager.
func auditmanager_CreateControl(cfg aws.Config, client *auditmanager.Client) {
	input := &auditmanager.CreateControlInput{
		// ControlMappingSources: []types.CreateControlMappingSource, // Required
		// Name: *string, // Required
	}

	if len(_auditmanagerControlMappingSources) > 0 {
		if err := assignInputField(input, "ControlMappingSources", _auditmanagerControlMappingSources); err != nil {
			log.Errorf("invalid --control-mapping-sources: %s", err.Error())
			return
		}
	}
	if len(_auditmanagerName) > 0 {
		input.Name = aws.String(_auditmanagerName)
	}
	if len(_auditmanagerActionPlanInstructions) > 0 {
		input.ActionPlanInstructions = aws.String(_auditmanagerActionPlanInstructions)
	}
	if len(_auditmanagerActionPlanTitle) > 0 {
		input.ActionPlanTitle = aws.String(_auditmanagerActionPlanTitle)
	}
	if len(_auditmanagerDescription) > 0 {
		input.Description = aws.String(_auditmanagerDescription)
	}
	if len(_auditmanagerTags) > 0 {
		if err := assignInputField(input, "Tags", _auditmanagerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_auditmanagerTestingInformation) > 0 {
		input.TestingInformation = aws.String(_auditmanagerTestingInformation)
	}

	if resp, err := client.CreateControl(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an assessment in Audit Manager.
func auditmanager_DeleteAssessment(cfg aws.Config, client *auditmanager.Client) {
	input := &auditmanager.DeleteAssessmentInput{
		// AssessmentId: *string, // Required
	}

	if len(_auditmanagerAssessmentId) > 0 {
		input.AssessmentId = aws.String(_auditmanagerAssessmentId)
	}

	if resp, err := client.DeleteAssessment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a custom framework in Audit Manager.
func auditmanager_DeleteAssessmentFramework(cfg aws.Config, client *auditmanager.Client) {
	input := &auditmanager.DeleteAssessmentFrameworkInput{
		// FrameworkId: *string, // Required
	}

	if len(_auditmanagerFrameworkId) > 0 {
		input.FrameworkId = aws.String(_auditmanagerFrameworkId)
	}

	if resp, err := client.DeleteAssessmentFramework(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a share request for a custom framework in Audit Manager.
func auditmanager_DeleteAssessmentFrameworkShare(cfg aws.Config, client *auditmanager.Client) {
	input := &auditmanager.DeleteAssessmentFrameworkShareInput{
		// RequestId: *string, // Required
		// RequestType: types.ShareRequestType, // Required
	}

	if len(_auditmanagerRequestId) > 0 {
		input.RequestId = aws.String(_auditmanagerRequestId)
	}
	if len(_auditmanagerRequestType) > 0 {
		if err := assignInputField(input, "RequestType", _auditmanagerRequestType); err != nil {
			log.Errorf("invalid --request-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteAssessmentFrameworkShare(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an assessment report in Audit Manager.
// When you run the DeleteAssessmentReport operation, Audit Manager attempts to
// delete the following data:
//
// - The specified assessment report that’s stored in your S3 bucket
//
// - The associated metadata that’s stored in Audit Manager
//
// If Audit Manager can’t access the assessment report in your S3 bucket, the
// report isn’t deleted. In this event, the DeleteAssessmentReport operation
// doesn’t fail. Instead, it proceeds to delete the associated metadata only. You
// must then delete the assessment report from the S3 bucket yourself.
//
// This scenario happens when Audit Manager receives a 403 (Forbidden) or 404 (Not
// Found) error from Amazon S3. To avoid this, make sure that your S3 bucket is
// available, and that you configured the correct permissions for Audit Manager to
// delete resources in your S3 bucket. For an example permissions policy that you
// can use, see [Assessment report destination permissions]in the Audit Manager User Guide. For information about the issues
// that could cause a 403 (Forbidden) or 404 (Not Found ) error from Amazon S3, see [List of Error Codes]
// in the Amazon Simple Storage Service API Reference.
//
// [List of Error Codes]: https://docs.aws.amazon.com/AmazonS3/latest/API/ErrorResponses.html#ErrorCodeList
// [Assessment report destination permissions]: https://docs.aws.amazon.com/audit-manager/latest/userguide/security_iam_id-based-policy-examples.html#full-administrator-access-assessment-report-destination
func auditmanager_DeleteAssessmentReport(cfg aws.Config, client *auditmanager.Client) {
	input := &auditmanager.DeleteAssessmentReportInput{
		// AssessmentId: *string, // Required
		// AssessmentReportId: *string, // Required
	}

	if len(_auditmanagerAssessmentId) > 0 {
		input.AssessmentId = aws.String(_auditmanagerAssessmentId)
	}
	if len(_auditmanagerAssessmentReportId) > 0 {
		input.AssessmentReportId = aws.String(_auditmanagerAssessmentReportId)
	}

	if resp, err := client.DeleteAssessmentReport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a custom control in Audit Manager.
// When you invoke this operation, the custom control is deleted from any
// frameworks or assessments that it’s currently part of. As a result, Audit
// Manager will stop collecting evidence for that custom control in all of your
// assessments. This includes assessments that you previously created before you
// deleted the custom control.
func auditmanager_DeleteControl(cfg aws.Config, client *auditmanager.Client) {
	input := &auditmanager.DeleteControlInput{
		// ControlId: *string, // Required
	}

	if len(_auditmanagerControlId) > 0 {
		input.ControlId = aws.String(_auditmanagerControlId)
	}

	if resp, err := client.DeleteControl(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deregisters an account in Audit Manager.
// Before you deregister, you can use the [UpdateSettings] API operation to set your preferred
// data retention policy. By default, Audit Manager retains your data. If you want
// to delete your data, you can use the DeregistrationPolicy attribute to request
// the deletion of your data.
//
// For more information about data retention, see [Data Protection] in the Audit Manager User
// Guide.
//
// [Data Protection]: https://docs.aws.amazon.com/audit-manager/latest/userguide/data-protection.html
// [UpdateSettings]: https://docs.aws.amazon.com/audit-manager/latest/APIReference/API_UpdateSettings.html
func auditmanager_DeregisterAccount(cfg aws.Config, client *auditmanager.Client) {
	input := &auditmanager.DeregisterAccountInput{}

	if resp, err := client.DeregisterAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the specified Amazon Web Services account as a delegated administrator
// for Audit Manager.
//
// When you remove a delegated administrator from your Audit Manager settings, you
// continue to have access to the evidence that you previously collected under that
// account. This is also the case when you deregister a delegated administrator
// from Organizations. However, Audit Manager stops collecting and attaching
// evidence to that delegated administrator account moving forward.
//
// Keep in mind the following cleanup task if you use evidence finder:
//
// Before you use your management account to remove a delegated administrator,
// make sure that the current delegated administrator account signs in to Audit
// Manager and disables evidence finder first. Disabling evidence finder
// automatically deletes the event data store that was created in their account
// when they enabled evidence finder. If this task isn’t completed, the event data
// store remains in their account. In this case, we recommend that the original
// delegated administrator goes to CloudTrail Lake and manually [deletes the event data store].
//
// This cleanup task is necessary to ensure that you don't end up with multiple
// event data stores. Audit Manager ignores an unused event data store after you
// remove or change a delegated administrator account. However, the unused event
// data store continues to incur storage costs from CloudTrail Lake if you don't
// delete it.
//
// When you deregister a delegated administrator account for Audit Manager, the
// data for that account isn’t deleted. If you want to delete resource data for a
// delegated administrator account, you must perform that task separately before
// you deregister the account. Either, you can do this in the Audit Manager
// console. Or, you can use one of the delete API operations that are provided by
// Audit Manager.
//
// To delete your Audit Manager resource data, see the following instructions:
//
// [DeleteAssessment]
// - (see also: [Deleting an assessment]in the Audit Manager User Guide)
//
// [DeleteAssessmentFramework]
// - (see also: [Deleting a custom framework]in the Audit Manager User Guide)
//
// [DeleteAssessmentFrameworkShare]
// - (see also: [Deleting a share request]in the Audit Manager User Guide)
//
// [DeleteAssessmentReport]
// - (see also: [Deleting an assessment report]in the Audit Manager User Guide)
//
// [DeleteControl]
// - (see also: [Deleting a custom control]in the Audit Manager User Guide)
//
// At this time, Audit Manager doesn't provide an option to delete evidence for a
// specific delegated administrator. Instead, when your management account
// deregisters Audit Manager, we perform a cleanup for the current delegated
// administrator account at the time of deregistration.
//
// [DeleteControl]: https://docs.aws.amazon.com/audit-manager/latest/APIReference/API_DeleteControl.html
// [deletes the event data store]: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/query-eds-disable-termination.html
// [DeleteAssessmentFrameworkShare]: https://docs.aws.amazon.com/audit-manager/latest/APIReference/API_DeleteAssessmentFrameworkShare.html
// [Deleting a custom control]: https://docs.aws.amazon.com/audit-manager/latest/userguide/delete-controls.html
// [Deleting an assessment report]: https://docs.aws.amazon.com/audit-manager/latest/userguide/generate-assessment-report.html#delete-assessment-report-steps
// [Deleting a custom framework]: https://docs.aws.amazon.com/audit-manager/latest/userguide/delete-custom-framework.html
// [Deleting a share request]: https://docs.aws.amazon.com/audit-manager/latest/userguide/deleting-shared-framework-requests.html
// [Deleting an assessment]: https://docs.aws.amazon.com/audit-manager/latest/userguide/delete-assessment.html
// [DeleteAssessment]: https://docs.aws.amazon.com/audit-manager/latest/APIReference/API_DeleteAssessment.html
// [DeleteAssessmentReport]: https://docs.aws.amazon.com/audit-manager/latest/APIReference/API_DeleteAssessmentReport.html
// [DeleteAssessmentFramework]: https://docs.aws.amazon.com/audit-manager/latest/APIReference/API_DeleteAssessmentFramework.html
func auditmanager_DeregisterOrganizationAdminAccount(cfg aws.Config, client *auditmanager.Client) {
	input := &auditmanager.DeregisterOrganizationAdminAccountInput{}

	if len(_auditmanagerAdminAccountId) > 0 {
		input.AdminAccountId = aws.String(_auditmanagerAdminAccountId)
	}

	if resp, err := client.DeregisterOrganizationAdminAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates an evidence folder from the specified assessment report in Audit
// Manager.
func auditmanager_DisassociateAssessmentReportEvidenceFolder(cfg aws.Config, client *auditmanager.Client) {
	input := &auditmanager.DisassociateAssessmentReportEvidenceFolderInput{
		// AssessmentId: *string, // Required
		// EvidenceFolderId: *string, // Required
	}

	if len(_auditmanagerAssessmentId) > 0 {
		input.AssessmentId = aws.String(_auditmanagerAssessmentId)
	}
	if len(_auditmanagerEvidenceFolderId) > 0 {
		input.EvidenceFolderId = aws.String(_auditmanagerEvidenceFolderId)
	}

	if resp, err := client.DisassociateAssessmentReportEvidenceFolder(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the registration status of an account in Audit Manager.
func auditmanager_GetAccountStatus(cfg aws.Config, client *auditmanager.Client) {
	input := &auditmanager.GetAccountStatusInput{}

	if resp, err := client.GetAccountStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a specified assessment.
func auditmanager_GetAssessment(cfg aws.Config, client *auditmanager.Client) {
	input := &auditmanager.GetAssessmentInput{
		// AssessmentId: *string, // Required
	}

	if len(_auditmanagerAssessmentId) > 0 {
		input.AssessmentId = aws.String(_auditmanagerAssessmentId)
	}

	if resp, err := client.GetAssessment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a specified framework.
func auditmanager_GetAssessmentFramework(cfg aws.Config, client *auditmanager.Client) {
	input := &auditmanager.GetAssessmentFrameworkInput{
		// FrameworkId: *string, // Required
	}

	if len(_auditmanagerFrameworkId) > 0 {
		input.FrameworkId = aws.String(_auditmanagerFrameworkId)
	}

	if resp, err := client.GetAssessmentFramework(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the URL of an assessment report in Audit Manager.
func auditmanager_GetAssessmentReportUrl(cfg aws.Config, client *auditmanager.Client) {
	input := &auditmanager.GetAssessmentReportUrlInput{
		// AssessmentId: *string, // Required
		// AssessmentReportId: *string, // Required
	}

	if len(_auditmanagerAssessmentId) > 0 {
		input.AssessmentId = aws.String(_auditmanagerAssessmentId)
	}
	if len(_auditmanagerAssessmentReportId) > 0 {
		input.AssessmentReportId = aws.String(_auditmanagerAssessmentReportId)
	}

	if resp, err := client.GetAssessmentReportUrl(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a list of changelogs from Audit Manager.
func auditmanager_GetChangeLogs(cfg aws.Config, client *auditmanager.Client) {
	input := &auditmanager.GetChangeLogsInput{
		// AssessmentId: *string, // Required
	}

	if len(_auditmanagerAssessmentId) > 0 {
		input.AssessmentId = aws.String(_auditmanagerAssessmentId)
	}
	if len(_auditmanagerControlId) > 0 {
		input.ControlId = aws.String(_auditmanagerControlId)
	}
	if len(_auditmanagerControlSetId) > 0 {
		input.ControlSetId = aws.String(_auditmanagerControlSetId)
	}
	if len(_auditmanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _auditmanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_auditmanagerNextToken) > 0 {
		input.NextToken = aws.String(_auditmanagerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetChangeLogs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*auditmanager.GetChangeLogsOutput
	p := auditmanager.NewGetChangeLogsPaginator(client, input)
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

// Gets information about a specified control.
func auditmanager_GetControl(cfg aws.Config, client *auditmanager.Client) {
	input := &auditmanager.GetControlInput{
		// ControlId: *string, // Required
	}

	if len(_auditmanagerControlId) > 0 {
		input.ControlId = aws.String(_auditmanagerControlId)
	}

	if resp, err := client.GetControl(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a list of delegations from an audit owner to a delegate.
func auditmanager_GetDelegations(cfg aws.Config, client *auditmanager.Client) {
	input := &auditmanager.GetDelegationsInput{}

	if len(_auditmanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _auditmanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_auditmanagerNextToken) > 0 {
		input.NextToken = aws.String(_auditmanagerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetDelegations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*auditmanager.GetDelegationsOutput
	p := auditmanager.NewGetDelegationsPaginator(client, input)
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

// Gets information about a specified evidence item.
func auditmanager_GetEvidence(cfg aws.Config, client *auditmanager.Client) {
	input := &auditmanager.GetEvidenceInput{
		// AssessmentId: *string, // Required
		// ControlSetId: *string, // Required
		// EvidenceFolderId: *string, // Required
		// EvidenceId: *string, // Required
	}

	if len(_auditmanagerAssessmentId) > 0 {
		input.AssessmentId = aws.String(_auditmanagerAssessmentId)
	}
	if len(_auditmanagerControlSetId) > 0 {
		input.ControlSetId = aws.String(_auditmanagerControlSetId)
	}
	if len(_auditmanagerEvidenceFolderId) > 0 {
		input.EvidenceFolderId = aws.String(_auditmanagerEvidenceFolderId)
	}
	if len(_auditmanagerEvidenceId) > 0 {
		input.EvidenceId = aws.String(_auditmanagerEvidenceId)
	}

	if resp, err := client.GetEvidence(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets all evidence from a specified evidence folder in Audit Manager.
func auditmanager_GetEvidenceByEvidenceFolder(cfg aws.Config, client *auditmanager.Client) {
	input := &auditmanager.GetEvidenceByEvidenceFolderInput{
		// AssessmentId: *string, // Required
		// ControlSetId: *string, // Required
		// EvidenceFolderId: *string, // Required
	}

	if len(_auditmanagerAssessmentId) > 0 {
		input.AssessmentId = aws.String(_auditmanagerAssessmentId)
	}
	if len(_auditmanagerControlSetId) > 0 {
		input.ControlSetId = aws.String(_auditmanagerControlSetId)
	}
	if len(_auditmanagerEvidenceFolderId) > 0 {
		input.EvidenceFolderId = aws.String(_auditmanagerEvidenceFolderId)
	}
	if len(_auditmanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _auditmanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_auditmanagerNextToken) > 0 {
		input.NextToken = aws.String(_auditmanagerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetEvidenceByEvidenceFolder(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*auditmanager.GetEvidenceByEvidenceFolderOutput
	p := auditmanager.NewGetEvidenceByEvidenceFolderPaginator(client, input)
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

// Creates a presigned Amazon S3 URL that can be used to upload a file as manual
// evidence. For instructions on how to use this operation, see [Upload a file from your browser]in the Audit
// Manager User Guide.
//
// The following restrictions apply to this operation:
//
// - Maximum size of an individual evidence file: 100 MB
//
// - Number of daily manual evidence uploads per control: 100
//
// - Supported file formats: See [Supported file types for manual evidence]in the Audit Manager User Guide
//
// For more information about Audit Manager service restrictions, see [Quotas and restrictions for Audit Manager].
//
// [Supported file types for manual evidence]: https://docs.aws.amazon.com/audit-manager/latest/userguide/upload-evidence.html#supported-manual-evidence-files
// [Upload a file from your browser]: https://docs.aws.amazon.com/audit-manager/latest/userguide/upload-evidence.html#how-to-upload-manual-evidence-files
// [Quotas and restrictions for Audit Manager]: https://docs.aws.amazon.com/audit-manager/latest/userguide/service-quotas.html
func auditmanager_GetEvidenceFileUploadUrl(cfg aws.Config, client *auditmanager.Client) {
	input := &auditmanager.GetEvidenceFileUploadUrlInput{
		// FileName: *string, // Required
	}

	if len(_auditmanagerFileName) > 0 {
		input.FileName = aws.String(_auditmanagerFileName)
	}

	if resp, err := client.GetEvidenceFileUploadUrl(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets an evidence folder from a specified assessment in Audit Manager.
func auditmanager_GetEvidenceFolder(cfg aws.Config, client *auditmanager.Client) {
	input := &auditmanager.GetEvidenceFolderInput{
		// AssessmentId: *string, // Required
		// ControlSetId: *string, // Required
		// EvidenceFolderId: *string, // Required
	}

	if len(_auditmanagerAssessmentId) > 0 {
		input.AssessmentId = aws.String(_auditmanagerAssessmentId)
	}
	if len(_auditmanagerControlSetId) > 0 {
		input.ControlSetId = aws.String(_auditmanagerControlSetId)
	}
	if len(_auditmanagerEvidenceFolderId) > 0 {
		input.EvidenceFolderId = aws.String(_auditmanagerEvidenceFolderId)
	}

	if resp, err := client.GetEvidenceFolder(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the evidence folders from a specified assessment in Audit Manager.
func auditmanager_GetEvidenceFoldersByAssessment(cfg aws.Config, client *auditmanager.Client) {
	input := &auditmanager.GetEvidenceFoldersByAssessmentInput{
		// AssessmentId: *string, // Required
	}

	if len(_auditmanagerAssessmentId) > 0 {
		input.AssessmentId = aws.String(_auditmanagerAssessmentId)
	}
	if len(_auditmanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _auditmanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_auditmanagerNextToken) > 0 {
		input.NextToken = aws.String(_auditmanagerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetEvidenceFoldersByAssessment(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*auditmanager.GetEvidenceFoldersByAssessmentOutput
	p := auditmanager.NewGetEvidenceFoldersByAssessmentPaginator(client, input)
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

// Gets a list of evidence folders that are associated with a specified control
// in an Audit Manager assessment.
func auditmanager_GetEvidenceFoldersByAssessmentControl(cfg aws.Config, client *auditmanager.Client) {
	input := &auditmanager.GetEvidenceFoldersByAssessmentControlInput{
		// AssessmentId: *string, // Required
		// ControlId: *string, // Required
		// ControlSetId: *string, // Required
	}

	if len(_auditmanagerAssessmentId) > 0 {
		input.AssessmentId = aws.String(_auditmanagerAssessmentId)
	}
	if len(_auditmanagerControlId) > 0 {
		input.ControlId = aws.String(_auditmanagerControlId)
	}
	if len(_auditmanagerControlSetId) > 0 {
		input.ControlSetId = aws.String(_auditmanagerControlSetId)
	}
	if len(_auditmanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _auditmanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_auditmanagerNextToken) > 0 {
		input.NextToken = aws.String(_auditmanagerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetEvidenceFoldersByAssessmentControl(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*auditmanager.GetEvidenceFoldersByAssessmentControlOutput
	p := auditmanager.NewGetEvidenceFoldersByAssessmentControlPaginator(client, input)
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

// Gets the latest analytics data for all your current active assessments.
func auditmanager_GetInsights(cfg aws.Config, client *auditmanager.Client) {
	input := &auditmanager.GetInsightsInput{}

	if resp, err := client.GetInsights(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the latest analytics data for a specific active assessment.
func auditmanager_GetInsightsByAssessment(cfg aws.Config, client *auditmanager.Client) {
	input := &auditmanager.GetInsightsByAssessmentInput{
		// AssessmentId: *string, // Required
	}

	if len(_auditmanagerAssessmentId) > 0 {
		input.AssessmentId = aws.String(_auditmanagerAssessmentId)
	}

	if resp, err := client.GetInsightsByAssessment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the name of the delegated Amazon Web Services administrator account for a
// specified organization.
func auditmanager_GetOrganizationAdminAccount(cfg aws.Config, client *auditmanager.Client) {
	input := &auditmanager.GetOrganizationAdminAccountInput{}

	if resp, err := client.GetOrganizationAdminAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a list of the Amazon Web Services services from which Audit Manager can
// collect evidence.
//
// Audit Manager defines which Amazon Web Services services are in scope for an
// assessment. Audit Manager infers this scope by examining the assessment’s
// controls and their data sources, and then mapping this information to one or
// more of the corresponding Amazon Web Services services that are in this list.
//
// For information about why it's no longer possible to specify services in scope
// manually, see [I can't edit the services in scope for my assessment]in the Troubleshooting section of the Audit Manager user guide.
//
// [I can't edit the services in scope for my assessment]: https://docs.aws.amazon.com/audit-manager/latest/userguide/evidence-collection-issues.html#unable-to-edit-services
func auditmanager_GetServicesInScope(cfg aws.Config, client *auditmanager.Client) {
	input := &auditmanager.GetServicesInScopeInput{}

	if resp, err := client.GetServicesInScope(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the settings for a specified Amazon Web Services account.
func auditmanager_GetSettings(cfg aws.Config, client *auditmanager.Client) {
	input := &auditmanager.GetSettingsInput{
		// Attribute: types.SettingAttribute, // Required
	}

	if len(_auditmanagerAttribute) > 0 {
		if err := assignInputField(input, "Attribute", _auditmanagerAttribute); err != nil {
			log.Errorf("invalid --attribute: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the latest analytics data for controls within a specific control domain
// and a specific active assessment.
//
// Control insights are listed only if the control belongs to the control domain
// and assessment that was specified. Moreover, the control must have collected
// evidence on the lastUpdated date of controlInsightsByAssessment . If neither of
// these conditions are met, no data is listed for that control.
func auditmanager_ListAssessmentControlInsightsByControlDomain(cfg aws.Config, client *auditmanager.Client) {
	input := &auditmanager.ListAssessmentControlInsightsByControlDomainInput{
		// AssessmentId: *string, // Required
		// ControlDomainId: *string, // Required
	}

	if len(_auditmanagerAssessmentId) > 0 {
		input.AssessmentId = aws.String(_auditmanagerAssessmentId)
	}
	if len(_auditmanagerControlDomainId) > 0 {
		input.ControlDomainId = aws.String(_auditmanagerControlDomainId)
	}
	if len(_auditmanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _auditmanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_auditmanagerNextToken) > 0 {
		input.NextToken = aws.String(_auditmanagerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAssessmentControlInsightsByControlDomain(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*auditmanager.ListAssessmentControlInsightsByControlDomainOutput
	p := auditmanager.NewListAssessmentControlInsightsByControlDomainPaginator(client, input)
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

// Returns a list of sent or received share requests for custom frameworks in
// Audit Manager.
func auditmanager_ListAssessmentFrameworkShareRequests(cfg aws.Config, client *auditmanager.Client) {
	input := &auditmanager.ListAssessmentFrameworkShareRequestsInput{
		// RequestType: types.ShareRequestType, // Required
	}

	if len(_auditmanagerRequestType) > 0 {
		if err := assignInputField(input, "RequestType", _auditmanagerRequestType); err != nil {
			log.Errorf("invalid --request-type: %s", err.Error())
			return
		}
	}
	if len(_auditmanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _auditmanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_auditmanagerNextToken) > 0 {
		input.NextToken = aws.String(_auditmanagerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAssessmentFrameworkShareRequests(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*auditmanager.ListAssessmentFrameworkShareRequestsOutput
	p := auditmanager.NewListAssessmentFrameworkShareRequestsPaginator(client, input)
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

// Returns a list of the frameworks that are available in the Audit Manager
// framework library.
func auditmanager_ListAssessmentFrameworks(cfg aws.Config, client *auditmanager.Client) {
	input := &auditmanager.ListAssessmentFrameworksInput{
		// FrameworkType: types.FrameworkType, // Required
	}

	if len(_auditmanagerFrameworkType) > 0 {
		if err := assignInputField(input, "FrameworkType", _auditmanagerFrameworkType); err != nil {
			log.Errorf("invalid --framework-type: %s", err.Error())
			return
		}
	}
	if len(_auditmanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _auditmanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_auditmanagerNextToken) > 0 {
		input.NextToken = aws.String(_auditmanagerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAssessmentFrameworks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*auditmanager.ListAssessmentFrameworksOutput
	p := auditmanager.NewListAssessmentFrameworksPaginator(client, input)
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

// Returns a list of assessment reports created in Audit Manager.
func auditmanager_ListAssessmentReports(cfg aws.Config, client *auditmanager.Client) {
	input := &auditmanager.ListAssessmentReportsInput{}

	if len(_auditmanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _auditmanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_auditmanagerNextToken) > 0 {
		input.NextToken = aws.String(_auditmanagerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAssessmentReports(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*auditmanager.ListAssessmentReportsOutput
	p := auditmanager.NewListAssessmentReportsPaginator(client, input)
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

// Returns a list of current and past assessments from Audit Manager.
func auditmanager_ListAssessments(cfg aws.Config, client *auditmanager.Client) {
	input := &auditmanager.ListAssessmentsInput{}

	if len(_auditmanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _auditmanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_auditmanagerNextToken) > 0 {
		input.NextToken = aws.String(_auditmanagerNextToken)
	}
	if len(_auditmanagerStatus) > 0 {
		if err := assignInputField(input, "Status", _auditmanagerStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListAssessments(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*auditmanager.ListAssessmentsOutput
	p := auditmanager.NewListAssessmentsPaginator(client, input)
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

// Lists the latest analytics data for control domains across all of your active
// assessments.
//
// Audit Manager supports the control domains that are provided by Amazon Web
// Services Control Catalog. For information about how to find a list of available
// control domains, see [ListDomains]ListDomains in the Amazon Web Services Control Catalog API
// Reference.
//
// A control domain is listed only if at least one of the controls within that
// domain collected evidence on the lastUpdated date of controlDomainInsights . If
// this condition isn’t met, no data is listed for that control domain.
//
// [ListDomains]: https://docs.aws.amazon.com/controlcatalog/latest/APIReference/API_ListDomains.html
func auditmanager_ListControlDomainInsights(cfg aws.Config, client *auditmanager.Client) {
	input := &auditmanager.ListControlDomainInsightsInput{}

	if len(_auditmanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _auditmanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_auditmanagerNextToken) > 0 {
		input.NextToken = aws.String(_auditmanagerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListControlDomainInsights(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*auditmanager.ListControlDomainInsightsOutput
	p := auditmanager.NewListControlDomainInsightsPaginator(client, input)
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

// Lists analytics data for control domains within a specified active assessment.
// Audit Manager supports the control domains that are provided by Amazon Web
// Services Control Catalog. For information about how to find a list of available
// control domains, see [ListDomains]ListDomains in the Amazon Web Services Control Catalog API
// Reference.
//
// A control domain is listed only if at least one of the controls within that
// domain collected evidence on the lastUpdated date of controlDomainInsights . If
// this condition isn’t met, no data is listed for that domain.
//
// [ListDomains]: https://docs.aws.amazon.com/controlcatalog/latest/APIReference/API_ListDomains.html
func auditmanager_ListControlDomainInsightsByAssessment(cfg aws.Config, client *auditmanager.Client) {
	input := &auditmanager.ListControlDomainInsightsByAssessmentInput{
		// AssessmentId: *string, // Required
	}

	if len(_auditmanagerAssessmentId) > 0 {
		input.AssessmentId = aws.String(_auditmanagerAssessmentId)
	}
	if len(_auditmanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _auditmanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_auditmanagerNextToken) > 0 {
		input.NextToken = aws.String(_auditmanagerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListControlDomainInsightsByAssessment(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*auditmanager.ListControlDomainInsightsByAssessmentOutput
	p := auditmanager.NewListControlDomainInsightsByAssessmentPaginator(client, input)
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

// Lists the latest analytics data for controls within a specific control domain
// across all active assessments.
//
// Control insights are listed only if the control belongs to the control domain
// that was specified and the control collected evidence on the lastUpdated date
// of controlInsightsMetadata . If neither of these conditions are met, no data is
// listed for that control.
func auditmanager_ListControlInsightsByControlDomain(cfg aws.Config, client *auditmanager.Client) {
	input := &auditmanager.ListControlInsightsByControlDomainInput{
		// ControlDomainId: *string, // Required
	}

	if len(_auditmanagerControlDomainId) > 0 {
		input.ControlDomainId = aws.String(_auditmanagerControlDomainId)
	}
	if len(_auditmanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _auditmanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_auditmanagerNextToken) > 0 {
		input.NextToken = aws.String(_auditmanagerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListControlInsightsByControlDomain(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*auditmanager.ListControlInsightsByControlDomainOutput
	p := auditmanager.NewListControlInsightsByControlDomainPaginator(client, input)
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

// Returns a list of controls from Audit Manager.
func auditmanager_ListControls(cfg aws.Config, client *auditmanager.Client) {
	input := &auditmanager.ListControlsInput{
		// ControlType: types.ControlType, // Required
	}

	if len(_auditmanagerControlType) > 0 {
		if err := assignInputField(input, "ControlType", _auditmanagerControlType); err != nil {
			log.Errorf("invalid --control-type: %s", err.Error())
			return
		}
	}
	if len(_auditmanagerControlCatalogId) > 0 {
		input.ControlCatalogId = aws.String(_auditmanagerControlCatalogId)
	}
	if len(_auditmanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _auditmanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_auditmanagerNextToken) > 0 {
		input.NextToken = aws.String(_auditmanagerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListControls(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*auditmanager.ListControlsOutput
	p := auditmanager.NewListControlsPaginator(client, input)
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

// Returns a list of keywords that are pre-mapped to the specified control data
// source.
func auditmanager_ListKeywordsForDataSource(cfg aws.Config, client *auditmanager.Client) {
	input := &auditmanager.ListKeywordsForDataSourceInput{
		// Source: types.DataSourceType, // Required
	}

	if len(_auditmanagerSource) > 0 {
		if err := assignInputField(input, "Source", _auditmanagerSource); err != nil {
			log.Errorf("invalid --source: %s", err.Error())
			return
		}
	}
	if len(_auditmanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _auditmanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_auditmanagerNextToken) > 0 {
		input.NextToken = aws.String(_auditmanagerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListKeywordsForDataSource(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*auditmanager.ListKeywordsForDataSourceOutput
	p := auditmanager.NewListKeywordsForDataSourcePaginator(client, input)
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

// Returns a list of all Audit Manager notifications.
func auditmanager_ListNotifications(cfg aws.Config, client *auditmanager.Client) {
	input := &auditmanager.ListNotificationsInput{}

	if len(_auditmanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _auditmanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_auditmanagerNextToken) > 0 {
		input.NextToken = aws.String(_auditmanagerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListNotifications(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*auditmanager.ListNotificationsOutput
	p := auditmanager.NewListNotificationsPaginator(client, input)
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

// Returns a list of tags for the specified resource in Audit Manager.
func auditmanager_ListTagsForResource(cfg aws.Config, client *auditmanager.Client) {
	input := &auditmanager.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_auditmanagerResourceArn) > 0 {
		input.ResourceArn = aws.String(_auditmanagerResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables Audit Manager for the specified Amazon Web Services account.
func auditmanager_RegisterAccount(cfg aws.Config, client *auditmanager.Client) {
	input := &auditmanager.RegisterAccountInput{}

	if len(_auditmanagerDelegatedAdminAccount) > 0 {
		input.DelegatedAdminAccount = aws.String(_auditmanagerDelegatedAdminAccount)
	}
	if len(_auditmanagerKmsKey) > 0 {
		input.KmsKey = aws.String(_auditmanagerKmsKey)
	}

	if resp, err := client.RegisterAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables an Amazon Web Services account within the organization as the
// delegated administrator for Audit Manager.
func auditmanager_RegisterOrganizationAdminAccount(cfg aws.Config, client *auditmanager.Client) {
	input := &auditmanager.RegisterOrganizationAdminAccountInput{
		// AdminAccountId: *string, // Required
	}

	if len(_auditmanagerAdminAccountId) > 0 {
		input.AdminAccountId = aws.String(_auditmanagerAdminAccountId)
	}

	if resp, err := client.RegisterOrganizationAdminAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a share request for a custom framework in Audit Manager.
// The share request specifies a recipient and notifies them that a custom
// framework is available. Recipients have 120 days to accept or decline the
// request. If no action is taken, the share request expires.
//
// When you create a share request, Audit Manager stores a snapshot of your custom
// framework in the US East (N. Virginia) Amazon Web Services Region. Audit Manager
// also stores a backup of the same snapshot in the US West (Oregon) Amazon Web
// Services Region.
//
// Audit Manager deletes the snapshot and the backup snapshot when one of the
// following events occurs:
//
// - The sender revokes the share request.
//
// - The recipient declines the share request.
//
// - The recipient encounters an error and doesn't successfully accept the share
// request.
//
// - The share request expires before the recipient responds to the request.
//
// When a sender [resends a share request], the snapshot is replaced with an updated version that
// corresponds with the latest version of the custom framework.
//
// When a recipient accepts a share request, the snapshot is replicated into their
// Amazon Web Services account under the Amazon Web Services Region that was
// specified in the share request.
//
// When you invoke the StartAssessmentFrameworkShare API, you are about to share a
// custom framework with another Amazon Web Services account. You may not share a
// custom framework that is derived from a standard framework if the standard
// framework is designated as not eligible for sharing by Amazon Web Services,
// unless you have obtained permission to do so from the owner of the standard
// framework. To learn more about which standard frameworks are eligible for
// sharing, see [Framework sharing eligibility]in the Audit Manager User Guide.
//
// [resends a share request]: https://docs.aws.amazon.com/audit-manager/latest/userguide/framework-sharing.html#framework-sharing-resend
// [Framework sharing eligibility]: https://docs.aws.amazon.com/audit-manager/latest/userguide/share-custom-framework-concepts-and-terminology.html#eligibility
func auditmanager_StartAssessmentFrameworkShare(cfg aws.Config, client *auditmanager.Client) {
	input := &auditmanager.StartAssessmentFrameworkShareInput{
		// DestinationAccount: *string, // Required
		// DestinationRegion: *string, // Required
		// FrameworkId: *string, // Required
	}

	if len(_auditmanagerDestinationAccount) > 0 {
		input.DestinationAccount = aws.String(_auditmanagerDestinationAccount)
	}
	if len(_auditmanagerDestinationRegion) > 0 {
		input.DestinationRegion = aws.String(_auditmanagerDestinationRegion)
	}
	if len(_auditmanagerFrameworkId) > 0 {
		input.FrameworkId = aws.String(_auditmanagerFrameworkId)
	}
	if len(_auditmanagerComment) > 0 {
		input.Comment = aws.String(_auditmanagerComment)
	}

	if resp, err := client.StartAssessmentFrameworkShare(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Tags the specified resource in Audit Manager.
func auditmanager_TagResource(cfg aws.Config, client *auditmanager.Client) {
	input := &auditmanager.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_auditmanagerResourceArn) > 0 {
		input.ResourceArn = aws.String(_auditmanagerResourceArn)
	}
	if len(_auditmanagerTags) > 0 {
		if err := assignInputField(input, "Tags", _auditmanagerTags); err != nil {
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

// Removes a tag from a resource in Audit Manager.
func auditmanager_UntagResource(cfg aws.Config, client *auditmanager.Client) {
	input := &auditmanager.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_auditmanagerResourceArn) > 0 {
		input.ResourceArn = aws.String(_auditmanagerResourceArn)
	}
	if len(_auditmanagerTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _auditmanagerTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Edits an Audit Manager assessment.
func auditmanager_UpdateAssessment(cfg aws.Config, client *auditmanager.Client) {
	input := &auditmanager.UpdateAssessmentInput{
		// AssessmentId: *string, // Required
		// Scope: *types.Scope, // Required
	}

	if len(_auditmanagerAssessmentId) > 0 {
		input.AssessmentId = aws.String(_auditmanagerAssessmentId)
	}
	if len(_auditmanagerScope) > 0 {
		if err := assignInputField(input, "Scope", _auditmanagerScope); err != nil {
			log.Errorf("invalid --scope: %s", err.Error())
			return
		}
	}
	if len(_auditmanagerAssessmentDescription) > 0 {
		input.AssessmentDescription = aws.String(_auditmanagerAssessmentDescription)
	}
	if len(_auditmanagerAssessmentName) > 0 {
		input.AssessmentName = aws.String(_auditmanagerAssessmentName)
	}
	if len(_auditmanagerAssessmentReportsDestination) > 0 {
		if err := assignInputField(input, "AssessmentReportsDestination", _auditmanagerAssessmentReportsDestination); err != nil {
			log.Errorf("invalid --assessment-reports-destination: %s", err.Error())
			return
		}
	}
	if len(_auditmanagerRoles) > 0 {
		if err := assignInputField(input, "Roles", _auditmanagerRoles); err != nil {
			log.Errorf("invalid --roles: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateAssessment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a control within an assessment in Audit Manager.
func auditmanager_UpdateAssessmentControl(cfg aws.Config, client *auditmanager.Client) {
	input := &auditmanager.UpdateAssessmentControlInput{
		// AssessmentId: *string, // Required
		// ControlId: *string, // Required
		// ControlSetId: *string, // Required
	}

	if len(_auditmanagerAssessmentId) > 0 {
		input.AssessmentId = aws.String(_auditmanagerAssessmentId)
	}
	if len(_auditmanagerControlId) > 0 {
		input.ControlId = aws.String(_auditmanagerControlId)
	}
	if len(_auditmanagerControlSetId) > 0 {
		input.ControlSetId = aws.String(_auditmanagerControlSetId)
	}
	if len(_auditmanagerCommentBody) > 0 {
		input.CommentBody = aws.String(_auditmanagerCommentBody)
	}
	if len(_auditmanagerControlStatus) > 0 {
		if err := assignInputField(input, "ControlStatus", _auditmanagerControlStatus); err != nil {
			log.Errorf("invalid --control-status: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateAssessmentControl(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the status of a control set in an Audit Manager assessment.
func auditmanager_UpdateAssessmentControlSetStatus(cfg aws.Config, client *auditmanager.Client) {
	input := &auditmanager.UpdateAssessmentControlSetStatusInput{
		// AssessmentId: *string, // Required
		// Comment: *string, // Required
		// ControlSetId: *string, // Required
		// Status: types.ControlSetStatus, // Required
	}

	if len(_auditmanagerAssessmentId) > 0 {
		input.AssessmentId = aws.String(_auditmanagerAssessmentId)
	}
	if len(_auditmanagerComment) > 0 {
		input.Comment = aws.String(_auditmanagerComment)
	}
	if len(_auditmanagerControlSetId) > 0 {
		input.ControlSetId = aws.String(_auditmanagerControlSetId)
	}
	if len(_auditmanagerStatus) > 0 {
		if err := assignInputField(input, "Status", _auditmanagerStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateAssessmentControlSetStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a custom framework in Audit Manager.
func auditmanager_UpdateAssessmentFramework(cfg aws.Config, client *auditmanager.Client) {
	input := &auditmanager.UpdateAssessmentFrameworkInput{
		// ControlSets: []types.UpdateAssessmentFrameworkControlSet, // Required
		// FrameworkId: *string, // Required
		// Name: *string, // Required
	}

	if len(_auditmanagerControlSets) > 0 {
		if err := assignInputField(input, "ControlSets", _auditmanagerControlSets); err != nil {
			log.Errorf("invalid --control-sets: %s", err.Error())
			return
		}
	}
	if len(_auditmanagerFrameworkId) > 0 {
		input.FrameworkId = aws.String(_auditmanagerFrameworkId)
	}
	if len(_auditmanagerName) > 0 {
		input.Name = aws.String(_auditmanagerName)
	}
	if len(_auditmanagerComplianceType) > 0 {
		input.ComplianceType = aws.String(_auditmanagerComplianceType)
	}
	if len(_auditmanagerDescription) > 0 {
		input.Description = aws.String(_auditmanagerDescription)
	}

	if resp, err := client.UpdateAssessmentFramework(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a share request for a custom framework in Audit Manager.
func auditmanager_UpdateAssessmentFrameworkShare(cfg aws.Config, client *auditmanager.Client) {
	input := &auditmanager.UpdateAssessmentFrameworkShareInput{
		// Action: types.ShareRequestAction, // Required
		// RequestId: *string, // Required
		// RequestType: types.ShareRequestType, // Required
	}

	if len(_auditmanagerAction) > 0 {
		if err := assignInputField(input, "Action", _auditmanagerAction); err != nil {
			log.Errorf("invalid --action: %s", err.Error())
			return
		}
	}
	if len(_auditmanagerRequestId) > 0 {
		input.RequestId = aws.String(_auditmanagerRequestId)
	}
	if len(_auditmanagerRequestType) > 0 {
		if err := assignInputField(input, "RequestType", _auditmanagerRequestType); err != nil {
			log.Errorf("invalid --request-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateAssessmentFrameworkShare(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the status of an assessment in Audit Manager.
func auditmanager_UpdateAssessmentStatus(cfg aws.Config, client *auditmanager.Client) {
	input := &auditmanager.UpdateAssessmentStatusInput{
		// AssessmentId: *string, // Required
		// Status: types.AssessmentStatus, // Required
	}

	if len(_auditmanagerAssessmentId) > 0 {
		input.AssessmentId = aws.String(_auditmanagerAssessmentId)
	}
	if len(_auditmanagerStatus) > 0 {
		if err := assignInputField(input, "Status", _auditmanagerStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateAssessmentStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a custom control in Audit Manager.
func auditmanager_UpdateControl(cfg aws.Config, client *auditmanager.Client) {
	input := &auditmanager.UpdateControlInput{
		// ControlId: *string, // Required
		// ControlMappingSources: []types.ControlMappingSource, // Required
		// Name: *string, // Required
	}

	if len(_auditmanagerControlId) > 0 {
		input.ControlId = aws.String(_auditmanagerControlId)
	}
	if len(_auditmanagerControlMappingSources) > 0 {
		if err := assignInputField(input, "ControlMappingSources", _auditmanagerControlMappingSources); err != nil {
			log.Errorf("invalid --control-mapping-sources: %s", err.Error())
			return
		}
	}
	if len(_auditmanagerName) > 0 {
		input.Name = aws.String(_auditmanagerName)
	}
	if len(_auditmanagerActionPlanInstructions) > 0 {
		input.ActionPlanInstructions = aws.String(_auditmanagerActionPlanInstructions)
	}
	if len(_auditmanagerActionPlanTitle) > 0 {
		input.ActionPlanTitle = aws.String(_auditmanagerActionPlanTitle)
	}
	if len(_auditmanagerDescription) > 0 {
		input.Description = aws.String(_auditmanagerDescription)
	}
	if len(_auditmanagerTestingInformation) > 0 {
		input.TestingInformation = aws.String(_auditmanagerTestingInformation)
	}

	if resp, err := client.UpdateControl(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates Audit Manager settings for the current account.
func auditmanager_UpdateSettings(cfg aws.Config, client *auditmanager.Client) {
	input := &auditmanager.UpdateSettingsInput{}

	if len(_auditmanagerDefaultAssessmentReportsDestination) > 0 {
		if err := assignInputField(input, "DefaultAssessmentReportsDestination", _auditmanagerDefaultAssessmentReportsDestination); err != nil {
			log.Errorf("invalid --default-assessment-reports-destination: %s", err.Error())
			return
		}
	}
	if len(_auditmanagerDefaultExportDestination) > 0 {
		if err := assignInputField(input, "DefaultExportDestination", _auditmanagerDefaultExportDestination); err != nil {
			log.Errorf("invalid --default-export-destination: %s", err.Error())
			return
		}
	}
	if len(_auditmanagerDefaultProcessOwners) > 0 {
		if err := assignInputField(input, "DefaultProcessOwners", _auditmanagerDefaultProcessOwners); err != nil {
			log.Errorf("invalid --default-process-owners: %s", err.Error())
			return
		}
	}
	if len(_auditmanagerDeregistrationPolicy) > 0 {
		if err := assignInputField(input, "DeregistrationPolicy", _auditmanagerDeregistrationPolicy); err != nil {
			log.Errorf("invalid --deregistration-policy: %s", err.Error())
			return
		}
	}
	if len(_auditmanagerEvidenceFinderEnabled) > 0 {
		if err := assignInputField(input, "EvidenceFinderEnabled", _auditmanagerEvidenceFinderEnabled); err != nil {
			log.Errorf("invalid --evidence-finder-enabled: %s", err.Error())
			return
		}
	}
	if len(_auditmanagerKmsKey) > 0 {
		input.KmsKey = aws.String(_auditmanagerKmsKey)
	}
	if len(_auditmanagerSnsTopic) > 0 {
		input.SnsTopic = aws.String(_auditmanagerSnsTopic)
	}

	if resp, err := client.UpdateSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Validates the integrity of an assessment report in Audit Manager.
func auditmanager_ValidateAssessmentReportIntegrity(cfg aws.Config, client *auditmanager.Client) {
	input := &auditmanager.ValidateAssessmentReportIntegrityInput{
		// S3RelativePath: *string, // Required
	}

	if len(_auditmanagerS3RelativePath) > 0 {
		input.S3RelativePath = aws.String(_auditmanagerS3RelativePath)
	}

	if resp, err := client.ValidateAssessmentReportIntegrity(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_auditmanagerCmd)
	_auditmanagerCmd.Flags().SortFlags = false

	_auditmanagerCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_auditmanagerCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_auditmanagerCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_auditmanagerCmd.Flags().StringVarP(&_auditmanagerAction, "action", "", "", "Action")
	_auditmanagerCmd.Flags().StringVarP(&_auditmanagerActionPlanInstructions, "action-plan-instructions", "", "", "Action Plan Instructions")
	_auditmanagerCmd.Flags().StringVarP(&_auditmanagerActionPlanTitle, "action-plan-title", "", "", "Action Plan Title")
	_auditmanagerCmd.Flags().StringVarP(&_auditmanagerAdminAccountId, "admin-account-id", "", "", "Admin Account ID")
	_auditmanagerCmd.Flags().StringVarP(&_auditmanagerAssessmentDescription, "assessment-description", "", "", "Assessment Description")
	_auditmanagerCmd.Flags().StringVarP(&_auditmanagerAssessmentId, "assessment-id", "", "", "Assessment ID")
	_auditmanagerCmd.Flags().StringVarP(&_auditmanagerAssessmentName, "assessment-name", "", "", "Assessment Name")
	_auditmanagerCmd.Flags().StringVarP(&_auditmanagerAssessmentReportId, "assessment-report-id", "", "", "Assessment Report ID")
	_auditmanagerCmd.Flags().StringVarP(&_auditmanagerAssessmentReportsDestination, "assessment-reports-destination", "", "", "Assessment Reports Destination")
	_auditmanagerCmd.Flags().StringVarP(&_auditmanagerAttribute, "attribute", "", "", "Attribute")
	_auditmanagerCmd.Flags().StringVarP(&_auditmanagerComment, "comment", "", "", "Comment")
	_auditmanagerCmd.Flags().StringVarP(&_auditmanagerCommentBody, "comment-body", "", "", "Comment Body")
	_auditmanagerCmd.Flags().StringVarP(&_auditmanagerComplianceType, "compliance-type", "", "", "Compliance Type")
	_auditmanagerCmd.Flags().StringVarP(&_auditmanagerControlCatalogId, "control-catalog-id", "", "", "Control Catalog ID")
	_auditmanagerCmd.Flags().StringVarP(&_auditmanagerControlDomainId, "control-domain-id", "", "", "Control Domain ID")
	_auditmanagerCmd.Flags().StringVarP(&_auditmanagerControlId, "control-id", "", "", "Control ID")
	_auditmanagerCmd.Flags().StringVarP(&_auditmanagerControlMappingSources, "control-mapping-sources", "", "", "Control Mapping Sources")
	_auditmanagerCmd.Flags().StringVarP(&_auditmanagerControlSetId, "control-set-id", "", "", "Control Set ID")
	_auditmanagerCmd.Flags().StringVarP(&_auditmanagerControlSets, "control-sets", "", "", "Control Sets")
	_auditmanagerCmd.Flags().StringVarP(&_auditmanagerControlStatus, "control-status", "", "", "Control Status")
	_auditmanagerCmd.Flags().StringVarP(&_auditmanagerControlType, "control-type", "", "", "Control Type")
	_auditmanagerCmd.Flags().StringVarP(&_auditmanagerCreateDelegationRequests, "create-delegation-requests", "", "", "Create Delegation Requests")
	_auditmanagerCmd.Flags().StringVarP(&_auditmanagerDefaultAssessmentReportsDestination, "default-assessment-reports-destination", "", "", "Default Assessment Reports Destination")
	_auditmanagerCmd.Flags().StringVarP(&_auditmanagerDefaultExportDestination, "default-export-destination", "", "", "Default Export Destination")
	_auditmanagerCmd.Flags().StringVarP(&_auditmanagerDefaultProcessOwners, "default-process-owners", "", "", "Default Process Owners")
	_auditmanagerCmd.Flags().StringVarP(&_auditmanagerDelegatedAdminAccount, "delegated-admin-account", "", "", "Delegated Admin Account")
	_auditmanagerCmd.Flags().StringSliceVarP(&_auditmanagerDelegationIds, "delegation-ids", "", nil, "Delegation Ids")
	_auditmanagerCmd.Flags().StringVarP(&_auditmanagerDeregistrationPolicy, "deregistration-policy", "", "", "Deregistration Policy")
	_auditmanagerCmd.Flags().StringVarP(&_auditmanagerDescription, "description", "", "", "Description")
	_auditmanagerCmd.Flags().StringVarP(&_auditmanagerDestinationAccount, "destination-account", "", "", "Destination Account")
	_auditmanagerCmd.Flags().StringVarP(&_auditmanagerDestinationRegion, "destination-region", "", "", "Destination Region")
	_auditmanagerCmd.Flags().StringVarP(&_auditmanagerEvidenceFinderEnabled, "evidence-finder-enabled", "", "", "Evidence Finder Enabled")
	_auditmanagerCmd.Flags().StringVarP(&_auditmanagerEvidenceFolderId, "evidence-folder-id", "", "", "Evidence Folder ID")
	_auditmanagerCmd.Flags().StringVarP(&_auditmanagerEvidenceId, "evidence-id", "", "", "Evidence ID")
	_auditmanagerCmd.Flags().StringSliceVarP(&_auditmanagerEvidenceIds, "evidence-ids", "", nil, "Evidence Ids")
	_auditmanagerCmd.Flags().StringVarP(&_auditmanagerFileName, "file-name", "", "", "File Name")
	_auditmanagerCmd.Flags().StringVarP(&_auditmanagerFrameworkId, "framework-id", "", "", "Framework ID")
	_auditmanagerCmd.Flags().StringVarP(&_auditmanagerFrameworkType, "framework-type", "", "", "Framework Type")
	_auditmanagerCmd.Flags().StringVarP(&_auditmanagerKmsKey, "kms-key", "", "", "KMS Key")
	_auditmanagerCmd.Flags().StringVarP(&_auditmanagerManualEvidence, "manual-evidence", "", "", "Manual Evidence")
	_auditmanagerCmd.Flags().StringVarP(&_auditmanagerMaxResults, "max-results", "", "", "Max Results")
	_auditmanagerCmd.Flags().StringVarP(&_auditmanagerName, "name", "", "", "Name")
	_auditmanagerCmd.Flags().StringVarP(&_auditmanagerNextToken, "next-token", "", "", "Next Token")
	_auditmanagerCmd.Flags().StringVarP(&_auditmanagerQueryStatement, "query-statement", "", "", "Query Statement")
	_auditmanagerCmd.Flags().StringVarP(&_auditmanagerRequestId, "request-id", "", "", "Request ID")
	_auditmanagerCmd.Flags().StringVarP(&_auditmanagerRequestType, "request-type", "", "", "Request Type")
	_auditmanagerCmd.Flags().StringVarP(&_auditmanagerResourceArn, "resource-arn", "", "", "Resource ARN")
	_auditmanagerCmd.Flags().StringVarP(&_auditmanagerRoles, "roles", "", "", "Roles")
	_auditmanagerCmd.Flags().StringVarP(&_auditmanagerS3RelativePath, "s3-relative-path", "", "", "S3 Relative Path")
	_auditmanagerCmd.Flags().StringVarP(&_auditmanagerScope, "scope", "", "", "Scope")
	_auditmanagerCmd.Flags().StringVarP(&_auditmanagerSnsTopic, "sns-topic", "", "", "SNS Topic")
	_auditmanagerCmd.Flags().StringVarP(&_auditmanagerSource, "source", "", "", "Source")
	_auditmanagerCmd.Flags().StringVarP(&_auditmanagerStatus, "status", "", "", "Status")
	_auditmanagerCmd.Flags().StringSliceVarP(&_auditmanagerTagKeys, "tag-keys", "", nil, "Tag Keys")
	_auditmanagerCmd.Flags().StringVarP(&_auditmanagerTags, "tags", "", "", "Tags")
	_auditmanagerCmd.Flags().StringVarP(&_auditmanagerTestingInformation, "testing-information", "", "", "Testing Information")

	_auditmanagerCmd.Flags().BoolVarP(&_auditmanagerAssociateAssessmentReportEvidenceFolder, "associate-assessment-report-evidence-folder", "", false, "Associate Assessment Report Evidence Folder")
	_auditmanagerCmd.Flags().BoolVarP(&_auditmanagerBatchAssociateAssessmentReportEvidence, "batch-associate-assessment-report-evidence", "", false, "Batch Associate Assessment Report Evidence")
	_auditmanagerCmd.Flags().BoolVarP(&_auditmanagerBatchCreateDelegationByAssessment, "batch-create-delegation-by-assessment", "", false, "Batch Create Delegation By Assessment")
	_auditmanagerCmd.Flags().BoolVarP(&_auditmanagerBatchDeleteDelegationByAssessment, "batch-delete-delegation-by-assessment", "", false, "Batch Delete Delegation By Assessment")
	_auditmanagerCmd.Flags().BoolVarP(&_auditmanagerBatchDisassociateAssessmentReportEvidence, "batch-disassociate-assessment-report-evidence", "", false, "Batch Disassociate Assessment Report Evidence")
	_auditmanagerCmd.Flags().BoolVarP(&_auditmanagerBatchImportEvidenceToAssessmentControl, "batch-import-evidence-to-assessment-control", "", false, "Batch Import Evidence To Assessment Control")
	_auditmanagerCmd.Flags().BoolVarP(&_auditmanagerCreateAssessment, "create-assessment", "", false, "Create Assessment")
	_auditmanagerCmd.Flags().BoolVarP(&_auditmanagerCreateAssessmentFramework, "create-assessment-framework", "", false, "Create Assessment Framework")
	_auditmanagerCmd.Flags().BoolVarP(&_auditmanagerCreateAssessmentReport, "create-assessment-report", "", false, "Create Assessment Report")
	_auditmanagerCmd.Flags().BoolVarP(&_auditmanagerCreateControl, "create-control", "", false, "Create Control")
	_auditmanagerCmd.Flags().BoolVarP(&_auditmanagerDeleteAssessment, "delete-assessment", "", false, "Delete Assessment")
	_auditmanagerCmd.Flags().BoolVarP(&_auditmanagerDeleteAssessmentFramework, "delete-assessment-framework", "", false, "Delete Assessment Framework")
	_auditmanagerCmd.Flags().BoolVarP(&_auditmanagerDeleteAssessmentFrameworkShare, "delete-assessment-framework-share", "", false, "Delete Assessment Framework Share")
	_auditmanagerCmd.Flags().BoolVarP(&_auditmanagerDeleteAssessmentReport, "delete-assessment-report", "", false, "Delete Assessment Report")
	_auditmanagerCmd.Flags().BoolVarP(&_auditmanagerDeleteControl, "delete-control", "", false, "Delete Control")
	_auditmanagerCmd.Flags().BoolVarP(&_auditmanagerDeregisterAccount, "deregister-account", "", false, "Deregister Account")
	_auditmanagerCmd.Flags().BoolVarP(&_auditmanagerDeregisterOrganizationAdminAccount, "deregister-organization-admin-account", "", false, "Deregister Organization Admin Account")
	_auditmanagerCmd.Flags().BoolVarP(&_auditmanagerDisassociateAssessmentReportEvidenceFolder, "disassociate-assessment-report-evidence-folder", "", false, "Disassociate Assessment Report Evidence Folder")
	_auditmanagerCmd.Flags().BoolVarP(&_auditmanagerGetAccountStatus, "get-account-status", "", false, "Get Account Status")
	_auditmanagerCmd.Flags().BoolVarP(&_auditmanagerGetAssessment, "get-assessment", "", false, "Get Assessment")
	_auditmanagerCmd.Flags().BoolVarP(&_auditmanagerGetAssessmentFramework, "get-assessment-framework", "", false, "Get Assessment Framework")
	_auditmanagerCmd.Flags().BoolVarP(&_auditmanagerGetAssessmentReportUrl, "get-assessment-report-url", "", false, "Get Assessment Report URL")
	_auditmanagerCmd.Flags().BoolVarP(&_auditmanagerGetChangeLogs, "get-change-logs", "", false, "Get Change Logs")
	_auditmanagerCmd.Flags().BoolVarP(&_auditmanagerGetControl, "get-control", "", false, "Get Control")
	_auditmanagerCmd.Flags().BoolVarP(&_auditmanagerGetDelegations, "get-delegations", "", false, "Get Delegations")
	_auditmanagerCmd.Flags().BoolVarP(&_auditmanagerGetEvidence, "get-evidence", "", false, "Get Evidence")
	_auditmanagerCmd.Flags().BoolVarP(&_auditmanagerGetEvidenceByEvidenceFolder, "get-evidence-by-evidence-folder", "", false, "Get Evidence By Evidence Folder")
	_auditmanagerCmd.Flags().BoolVarP(&_auditmanagerGetEvidenceFileUploadUrl, "get-evidence-file-upload-url", "", false, "Get Evidence File Upload URL")
	_auditmanagerCmd.Flags().BoolVarP(&_auditmanagerGetEvidenceFolder, "get-evidence-folder", "", false, "Get Evidence Folder")
	_auditmanagerCmd.Flags().BoolVarP(&_auditmanagerGetEvidenceFoldersByAssessment, "get-evidence-folders-by-assessment", "", false, "Get Evidence Folders By Assessment")
	_auditmanagerCmd.Flags().BoolVarP(&_auditmanagerGetEvidenceFoldersByAssessmentControl, "get-evidence-folders-by-assessment-control", "", false, "Get Evidence Folders By Assessment Control")
	_auditmanagerCmd.Flags().BoolVarP(&_auditmanagerGetInsights, "get-insights", "", false, "Get Insights")
	_auditmanagerCmd.Flags().BoolVarP(&_auditmanagerGetInsightsByAssessment, "get-insights-by-assessment", "", false, "Get Insights By Assessment")
	_auditmanagerCmd.Flags().BoolVarP(&_auditmanagerGetOrganizationAdminAccount, "get-organization-admin-account", "", false, "Get Organization Admin Account")
	_auditmanagerCmd.Flags().BoolVarP(&_auditmanagerGetServicesInScope, "get-services-in-scope", "", false, "Get Services In Scope")
	_auditmanagerCmd.Flags().BoolVarP(&_auditmanagerGetSettings, "get-settings", "", false, "Get Settings")
	_auditmanagerCmd.Flags().BoolVarP(&_auditmanagerListAssessmentControlInsightsByControlDomain, "list-assessment-control-insights-by-control-domain", "", false, "List Assessment Control Insights By Control Domain")
	_auditmanagerCmd.Flags().BoolVarP(&_auditmanagerListAssessmentFrameworkShareRequests, "list-assessment-framework-share-requests", "", false, "List Assessment Framework Share Requests")
	_auditmanagerCmd.Flags().BoolVarP(&_auditmanagerListAssessmentFrameworks, "list-assessment-frameworks", "", false, "List Assessment Frameworks")
	_auditmanagerCmd.Flags().BoolVarP(&_auditmanagerListAssessmentReports, "list-assessment-reports", "", false, "List Assessment Reports")
	_auditmanagerCmd.Flags().BoolVarP(&_auditmanagerListAssessments, "list-assessments", "", false, "List Assessments")
	_auditmanagerCmd.Flags().BoolVarP(&_auditmanagerListControlDomainInsights, "list-control-domain-insights", "", false, "List Control Domain Insights")
	_auditmanagerCmd.Flags().BoolVarP(&_auditmanagerListControlDomainInsightsByAssessment, "list-control-domain-insights-by-assessment", "", false, "List Control Domain Insights By Assessment")
	_auditmanagerCmd.Flags().BoolVarP(&_auditmanagerListControlInsightsByControlDomain, "list-control-insights-by-control-domain", "", false, "List Control Insights By Control Domain")
	_auditmanagerCmd.Flags().BoolVarP(&_auditmanagerListControls, "list-controls", "", false, "List Controls")
	_auditmanagerCmd.Flags().BoolVarP(&_auditmanagerListKeywordsForDataSource, "list-keywords-for-data-source", "", false, "List Keywords For Data Source")
	_auditmanagerCmd.Flags().BoolVarP(&_auditmanagerListNotifications, "list-notifications", "", false, "List Notifications")
	_auditmanagerCmd.Flags().BoolVarP(&_auditmanagerListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_auditmanagerCmd.Flags().BoolVarP(&_auditmanagerRegisterAccount, "register-account", "", false, "Register Account")
	_auditmanagerCmd.Flags().BoolVarP(&_auditmanagerRegisterOrganizationAdminAccount, "register-organization-admin-account", "", false, "Register Organization Admin Account")
	_auditmanagerCmd.Flags().BoolVarP(&_auditmanagerStartAssessmentFrameworkShare, "start-assessment-framework-share", "", false, "Start Assessment Framework Share")
	_auditmanagerCmd.Flags().BoolVarP(&_auditmanagerTagResource, "tag-resource", "", false, "Tag Resource")
	_auditmanagerCmd.Flags().BoolVarP(&_auditmanagerUntagResource, "untag-resource", "", false, "Untag Resource")
	_auditmanagerCmd.Flags().BoolVarP(&_auditmanagerUpdateAssessment, "update-assessment", "", false, "Update Assessment")
	_auditmanagerCmd.Flags().BoolVarP(&_auditmanagerUpdateAssessmentControl, "update-assessment-control", "", false, "Update Assessment Control")
	_auditmanagerCmd.Flags().BoolVarP(&_auditmanagerUpdateAssessmentControlSetStatus, "update-assessment-control-set-status", "", false, "Update Assessment Control Set Status")
	_auditmanagerCmd.Flags().BoolVarP(&_auditmanagerUpdateAssessmentFramework, "update-assessment-framework", "", false, "Update Assessment Framework")
	_auditmanagerCmd.Flags().BoolVarP(&_auditmanagerUpdateAssessmentFrameworkShare, "update-assessment-framework-share", "", false, "Update Assessment Framework Share")
	_auditmanagerCmd.Flags().BoolVarP(&_auditmanagerUpdateAssessmentStatus, "update-assessment-status", "", false, "Update Assessment Status")
	_auditmanagerCmd.Flags().BoolVarP(&_auditmanagerUpdateControl, "update-control", "", false, "Update Control")
	_auditmanagerCmd.Flags().BoolVarP(&_auditmanagerUpdateSettings, "update-settings", "", false, "Update Settings")
	_auditmanagerCmd.Flags().BoolVarP(&_auditmanagerValidateAssessmentReportIntegrity, "validate-assessment-report-integrity", "", false, "Validate Assessment Report Integrity")

}
