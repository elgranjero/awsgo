package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cleanrooms"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// cleanroomsCmd represents the cleanrooms command
var _cleanroomsCmd = &cobra.Command{
	Use:   "cleanrooms",
	Short: "AWS cleanrooms CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := cleanrooms.NewFromConfig(cfg)
		if _cleanroomsBatchGetCollaborationAnalysisTemplate {
			cleanrooms_BatchGetCollaborationAnalysisTemplate(cfg, client)
			return
		}
		if _cleanroomsBatchGetSchema {
			cleanrooms_BatchGetSchema(cfg, client)
			return
		}
		if _cleanroomsBatchGetSchemaAnalysisRule {
			cleanrooms_BatchGetSchemaAnalysisRule(cfg, client)
			return
		}
		if _cleanroomsCreateAnalysisTemplate {
			cleanrooms_CreateAnalysisTemplate(cfg, client)
			return
		}
		if _cleanroomsCreateCollaboration {
			cleanrooms_CreateCollaboration(cfg, client)
			return
		}
		if _cleanroomsCreateCollaborationChangeRequest {
			cleanrooms_CreateCollaborationChangeRequest(cfg, client)
			return
		}
		if _cleanroomsCreateConfiguredAudienceModelAssociation {
			cleanrooms_CreateConfiguredAudienceModelAssociation(cfg, client)
			return
		}
		if _cleanroomsCreateConfiguredTable {
			cleanrooms_CreateConfiguredTable(cfg, client)
			return
		}
		if _cleanroomsCreateConfiguredTableAnalysisRule {
			cleanrooms_CreateConfiguredTableAnalysisRule(cfg, client)
			return
		}
		if _cleanroomsCreateConfiguredTableAssociation {
			cleanrooms_CreateConfiguredTableAssociation(cfg, client)
			return
		}
		if _cleanroomsCreateConfiguredTableAssociationAnalysisRule {
			cleanrooms_CreateConfiguredTableAssociationAnalysisRule(cfg, client)
			return
		}
		if _cleanroomsCreateIdMappingTable {
			cleanrooms_CreateIdMappingTable(cfg, client)
			return
		}
		if _cleanroomsCreateIdNamespaceAssociation {
			cleanrooms_CreateIdNamespaceAssociation(cfg, client)
			return
		}
		if _cleanroomsCreateMembership {
			cleanrooms_CreateMembership(cfg, client)
			return
		}
		if _cleanroomsCreatePrivacyBudgetTemplate {
			cleanrooms_CreatePrivacyBudgetTemplate(cfg, client)
			return
		}
		if _cleanroomsDeleteAnalysisTemplate {
			cleanrooms_DeleteAnalysisTemplate(cfg, client)
			return
		}
		if _cleanroomsDeleteCollaboration {
			cleanrooms_DeleteCollaboration(cfg, client)
			return
		}
		if _cleanroomsDeleteConfiguredAudienceModelAssociation {
			cleanrooms_DeleteConfiguredAudienceModelAssociation(cfg, client)
			return
		}
		if _cleanroomsDeleteConfiguredTable {
			cleanrooms_DeleteConfiguredTable(cfg, client)
			return
		}
		if _cleanroomsDeleteConfiguredTableAnalysisRule {
			cleanrooms_DeleteConfiguredTableAnalysisRule(cfg, client)
			return
		}
		if _cleanroomsDeleteConfiguredTableAssociation {
			cleanrooms_DeleteConfiguredTableAssociation(cfg, client)
			return
		}
		if _cleanroomsDeleteConfiguredTableAssociationAnalysisRule {
			cleanrooms_DeleteConfiguredTableAssociationAnalysisRule(cfg, client)
			return
		}
		if _cleanroomsDeleteIdMappingTable {
			cleanrooms_DeleteIdMappingTable(cfg, client)
			return
		}
		if _cleanroomsDeleteIdNamespaceAssociation {
			cleanrooms_DeleteIdNamespaceAssociation(cfg, client)
			return
		}
		if _cleanroomsDeleteMember {
			cleanrooms_DeleteMember(cfg, client)
			return
		}
		if _cleanroomsDeleteMembership {
			cleanrooms_DeleteMembership(cfg, client)
			return
		}
		if _cleanroomsDeletePrivacyBudgetTemplate {
			cleanrooms_DeletePrivacyBudgetTemplate(cfg, client)
			return
		}
		if _cleanroomsGetAnalysisTemplate {
			cleanrooms_GetAnalysisTemplate(cfg, client)
			return
		}
		if _cleanroomsGetCollaboration {
			cleanrooms_GetCollaboration(cfg, client)
			return
		}
		if _cleanroomsGetCollaborationAnalysisTemplate {
			cleanrooms_GetCollaborationAnalysisTemplate(cfg, client)
			return
		}
		if _cleanroomsGetCollaborationChangeRequest {
			cleanrooms_GetCollaborationChangeRequest(cfg, client)
			return
		}
		if _cleanroomsGetCollaborationConfiguredAudienceModelAssociation {
			cleanrooms_GetCollaborationConfiguredAudienceModelAssociation(cfg, client)
			return
		}
		if _cleanroomsGetCollaborationIdNamespaceAssociation {
			cleanrooms_GetCollaborationIdNamespaceAssociation(cfg, client)
			return
		}
		if _cleanroomsGetCollaborationPrivacyBudgetTemplate {
			cleanrooms_GetCollaborationPrivacyBudgetTemplate(cfg, client)
			return
		}
		if _cleanroomsGetConfiguredAudienceModelAssociation {
			cleanrooms_GetConfiguredAudienceModelAssociation(cfg, client)
			return
		}
		if _cleanroomsGetConfiguredTable {
			cleanrooms_GetConfiguredTable(cfg, client)
			return
		}
		if _cleanroomsGetConfiguredTableAnalysisRule {
			cleanrooms_GetConfiguredTableAnalysisRule(cfg, client)
			return
		}
		if _cleanroomsGetConfiguredTableAssociation {
			cleanrooms_GetConfiguredTableAssociation(cfg, client)
			return
		}
		if _cleanroomsGetConfiguredTableAssociationAnalysisRule {
			cleanrooms_GetConfiguredTableAssociationAnalysisRule(cfg, client)
			return
		}
		if _cleanroomsGetIdMappingTable {
			cleanrooms_GetIdMappingTable(cfg, client)
			return
		}
		if _cleanroomsGetIdNamespaceAssociation {
			cleanrooms_GetIdNamespaceAssociation(cfg, client)
			return
		}
		if _cleanroomsGetMembership {
			cleanrooms_GetMembership(cfg, client)
			return
		}
		if _cleanroomsGetPrivacyBudgetTemplate {
			cleanrooms_GetPrivacyBudgetTemplate(cfg, client)
			return
		}
		if _cleanroomsGetProtectedJob {
			cleanrooms_GetProtectedJob(cfg, client)
			return
		}
		if _cleanroomsGetProtectedQuery {
			cleanrooms_GetProtectedQuery(cfg, client)
			return
		}
		if _cleanroomsGetSchema {
			cleanrooms_GetSchema(cfg, client)
			return
		}
		if _cleanroomsGetSchemaAnalysisRule {
			cleanrooms_GetSchemaAnalysisRule(cfg, client)
			return
		}
		if _cleanroomsListAnalysisTemplates {
			cleanrooms_ListAnalysisTemplates(cfg, client)
			return
		}
		if _cleanroomsListCollaborationAnalysisTemplates {
			cleanrooms_ListCollaborationAnalysisTemplates(cfg, client)
			return
		}
		if _cleanroomsListCollaborationChangeRequests {
			cleanrooms_ListCollaborationChangeRequests(cfg, client)
			return
		}
		if _cleanroomsListCollaborationConfiguredAudienceModelAssociations {
			cleanrooms_ListCollaborationConfiguredAudienceModelAssociations(cfg, client)
			return
		}
		if _cleanroomsListCollaborationIdNamespaceAssociations {
			cleanrooms_ListCollaborationIdNamespaceAssociations(cfg, client)
			return
		}
		if _cleanroomsListCollaborationPrivacyBudgetTemplates {
			cleanrooms_ListCollaborationPrivacyBudgetTemplates(cfg, client)
			return
		}
		if _cleanroomsListCollaborationPrivacyBudgets {
			cleanrooms_ListCollaborationPrivacyBudgets(cfg, client)
			return
		}
		if _cleanroomsListCollaborations {
			cleanrooms_ListCollaborations(cfg, client)
			return
		}
		if _cleanroomsListConfiguredAudienceModelAssociations {
			cleanrooms_ListConfiguredAudienceModelAssociations(cfg, client)
			return
		}
		if _cleanroomsListConfiguredTableAssociations {
			cleanrooms_ListConfiguredTableAssociations(cfg, client)
			return
		}
		if _cleanroomsListConfiguredTables {
			cleanrooms_ListConfiguredTables(cfg, client)
			return
		}
		if _cleanroomsListIdMappingTables {
			cleanrooms_ListIdMappingTables(cfg, client)
			return
		}
		if _cleanroomsListIdNamespaceAssociations {
			cleanrooms_ListIdNamespaceAssociations(cfg, client)
			return
		}
		if _cleanroomsListMembers {
			cleanrooms_ListMembers(cfg, client)
			return
		}
		if _cleanroomsListMemberships {
			cleanrooms_ListMemberships(cfg, client)
			return
		}
		if _cleanroomsListPrivacyBudgetTemplates {
			cleanrooms_ListPrivacyBudgetTemplates(cfg, client)
			return
		}
		if _cleanroomsListPrivacyBudgets {
			cleanrooms_ListPrivacyBudgets(cfg, client)
			return
		}
		if _cleanroomsListProtectedJobs {
			cleanrooms_ListProtectedJobs(cfg, client)
			return
		}
		if _cleanroomsListProtectedQueries {
			cleanrooms_ListProtectedQueries(cfg, client)
			return
		}
		if _cleanroomsListSchemas {
			cleanrooms_ListSchemas(cfg, client)
			return
		}
		if _cleanroomsListTagsForResource {
			cleanrooms_ListTagsForResource(cfg, client)
			return
		}
		if _cleanroomsPopulateIdMappingTable {
			cleanrooms_PopulateIdMappingTable(cfg, client)
			return
		}
		if _cleanroomsPreviewPrivacyImpact {
			cleanrooms_PreviewPrivacyImpact(cfg, client)
			return
		}
		if _cleanroomsStartProtectedJob {
			cleanrooms_StartProtectedJob(cfg, client)
			return
		}
		if _cleanroomsStartProtectedQuery {
			cleanrooms_StartProtectedQuery(cfg, client)
			return
		}
		if _cleanroomsTagResource {
			cleanrooms_TagResource(cfg, client)
			return
		}
		if _cleanroomsUntagResource {
			cleanrooms_UntagResource(cfg, client)
			return
		}
		if _cleanroomsUpdateAnalysisTemplate {
			cleanrooms_UpdateAnalysisTemplate(cfg, client)
			return
		}
		if _cleanroomsUpdateCollaboration {
			cleanrooms_UpdateCollaboration(cfg, client)
			return
		}
		if _cleanroomsUpdateCollaborationChangeRequest {
			cleanrooms_UpdateCollaborationChangeRequest(cfg, client)
			return
		}
		if _cleanroomsUpdateConfiguredAudienceModelAssociation {
			cleanrooms_UpdateConfiguredAudienceModelAssociation(cfg, client)
			return
		}
		if _cleanroomsUpdateConfiguredTable {
			cleanrooms_UpdateConfiguredTable(cfg, client)
			return
		}
		if _cleanroomsUpdateConfiguredTableAnalysisRule {
			cleanrooms_UpdateConfiguredTableAnalysisRule(cfg, client)
			return
		}
		if _cleanroomsUpdateConfiguredTableAssociation {
			cleanrooms_UpdateConfiguredTableAssociation(cfg, client)
			return
		}
		if _cleanroomsUpdateConfiguredTableAssociationAnalysisRule {
			cleanrooms_UpdateConfiguredTableAssociationAnalysisRule(cfg, client)
			return
		}
		if _cleanroomsUpdateIdMappingTable {
			cleanrooms_UpdateIdMappingTable(cfg, client)
			return
		}
		if _cleanroomsUpdateIdNamespaceAssociation {
			cleanrooms_UpdateIdNamespaceAssociation(cfg, client)
			return
		}
		if _cleanroomsUpdateMembership {
			cleanrooms_UpdateMembership(cfg, client)
			return
		}
		if _cleanroomsUpdatePrivacyBudgetTemplate {
			cleanrooms_UpdatePrivacyBudgetTemplate(cfg, client)
			return
		}
		if _cleanroomsUpdateProtectedJob {
			cleanrooms_UpdateProtectedJob(cfg, client)
			return
		}
		if _cleanroomsUpdateProtectedQuery {
			cleanrooms_UpdateProtectedQuery(cfg, client)
			return
		}

	},
}

var (
	_cleanroomsBatchGetCollaborationAnalysisTemplate                bool
	_cleanroomsBatchGetSchema                                       bool
	_cleanroomsBatchGetSchemaAnalysisRule                           bool
	_cleanroomsCreateAnalysisTemplate                               bool
	_cleanroomsCreateCollaboration                                  bool
	_cleanroomsCreateCollaborationChangeRequest                     bool
	_cleanroomsCreateConfiguredAudienceModelAssociation             bool
	_cleanroomsCreateConfiguredTable                                bool
	_cleanroomsCreateConfiguredTableAnalysisRule                    bool
	_cleanroomsCreateConfiguredTableAssociation                     bool
	_cleanroomsCreateConfiguredTableAssociationAnalysisRule         bool
	_cleanroomsCreateIdMappingTable                                 bool
	_cleanroomsCreateIdNamespaceAssociation                         bool
	_cleanroomsCreateMembership                                     bool
	_cleanroomsCreatePrivacyBudgetTemplate                          bool
	_cleanroomsDeleteAnalysisTemplate                               bool
	_cleanroomsDeleteCollaboration                                  bool
	_cleanroomsDeleteConfiguredAudienceModelAssociation             bool
	_cleanroomsDeleteConfiguredTable                                bool
	_cleanroomsDeleteConfiguredTableAnalysisRule                    bool
	_cleanroomsDeleteConfiguredTableAssociation                     bool
	_cleanroomsDeleteConfiguredTableAssociationAnalysisRule         bool
	_cleanroomsDeleteIdMappingTable                                 bool
	_cleanroomsDeleteIdNamespaceAssociation                         bool
	_cleanroomsDeleteMember                                         bool
	_cleanroomsDeleteMembership                                     bool
	_cleanroomsDeletePrivacyBudgetTemplate                          bool
	_cleanroomsGetAnalysisTemplate                                  bool
	_cleanroomsGetCollaboration                                     bool
	_cleanroomsGetCollaborationAnalysisTemplate                     bool
	_cleanroomsGetCollaborationChangeRequest                        bool
	_cleanroomsGetCollaborationConfiguredAudienceModelAssociation   bool
	_cleanroomsGetCollaborationIdNamespaceAssociation               bool
	_cleanroomsGetCollaborationPrivacyBudgetTemplate                bool
	_cleanroomsGetConfiguredAudienceModelAssociation                bool
	_cleanroomsGetConfiguredTable                                   bool
	_cleanroomsGetConfiguredTableAnalysisRule                       bool
	_cleanroomsGetConfiguredTableAssociation                        bool
	_cleanroomsGetConfiguredTableAssociationAnalysisRule            bool
	_cleanroomsGetIdMappingTable                                    bool
	_cleanroomsGetIdNamespaceAssociation                            bool
	_cleanroomsGetMembership                                        bool
	_cleanroomsGetPrivacyBudgetTemplate                             bool
	_cleanroomsGetProtectedJob                                      bool
	_cleanroomsGetProtectedQuery                                    bool
	_cleanroomsGetSchema                                            bool
	_cleanroomsGetSchemaAnalysisRule                                bool
	_cleanroomsListAnalysisTemplates                                bool
	_cleanroomsListCollaborationAnalysisTemplates                   bool
	_cleanroomsListCollaborationChangeRequests                      bool
	_cleanroomsListCollaborationConfiguredAudienceModelAssociations bool
	_cleanroomsListCollaborationIdNamespaceAssociations             bool
	_cleanroomsListCollaborationPrivacyBudgetTemplates              bool
	_cleanroomsListCollaborationPrivacyBudgets                      bool
	_cleanroomsListCollaborations                                   bool
	_cleanroomsListConfiguredAudienceModelAssociations              bool
	_cleanroomsListConfiguredTableAssociations                      bool
	_cleanroomsListConfiguredTables                                 bool
	_cleanroomsListIdMappingTables                                  bool
	_cleanroomsListIdNamespaceAssociations                          bool
	_cleanroomsListMembers                                          bool
	_cleanroomsListMemberships                                      bool
	_cleanroomsListPrivacyBudgetTemplates                           bool
	_cleanroomsListPrivacyBudgets                                   bool
	_cleanroomsListProtectedJobs                                    bool
	_cleanroomsListProtectedQueries                                 bool
	_cleanroomsListSchemas                                          bool
	_cleanroomsListTagsForResource                                  bool
	_cleanroomsPopulateIdMappingTable                               bool
	_cleanroomsPreviewPrivacyImpact                                 bool
	_cleanroomsStartProtectedJob                                    bool
	_cleanroomsStartProtectedQuery                                  bool
	_cleanroomsTagResource                                          bool
	_cleanroomsUntagResource                                        bool
	_cleanroomsUpdateAnalysisTemplate                               bool
	_cleanroomsUpdateCollaboration                                  bool
	_cleanroomsUpdateCollaborationChangeRequest                     bool
	_cleanroomsUpdateConfiguredAudienceModelAssociation             bool
	_cleanroomsUpdateConfiguredTable                                bool
	_cleanroomsUpdateConfiguredTableAnalysisRule                    bool
	_cleanroomsUpdateConfiguredTableAssociation                     bool
	_cleanroomsUpdateConfiguredTableAssociationAnalysisRule         bool
	_cleanroomsUpdateIdMappingTable                                 bool
	_cleanroomsUpdateIdNamespaceAssociation                         bool
	_cleanroomsUpdateMembership                                     bool
	_cleanroomsUpdatePrivacyBudgetTemplate                          bool
	_cleanroomsUpdateProtectedJob                                   bool
	_cleanroomsUpdateProtectedQuery                                 bool

	_cleanroomsAccessBudgetResourceArn                      string
	_cleanroomsAccountId                                    string
	_cleanroomsAction                                       string
	_cleanroomsAllowedColumns                               []string
	_cleanroomsAllowedResultRegions                         string
	_cleanroomsAnalysisMethod                               string
	_cleanroomsAnalysisParameters                           string
	_cleanroomsAnalysisRulePolicy                           string
	_cleanroomsAnalysisRuleType                             string
	_cleanroomsAnalysisTemplateArn                          string
	_cleanroomsAnalysisTemplateArns                         []string
	_cleanroomsAnalysisTemplateIdentifier                   string
	_cleanroomsAnalyticsEngine                              string
	_cleanroomsAutoApprovedChangeRequestTypes               string
	_cleanroomsAutoRefresh                                  string
	_cleanroomsChangeRequestIdentifier                      string
	_cleanroomsChanges                                      string
	_cleanroomsCollaborationIdentifier                      string
	_cleanroomsComputeConfiguration                         string
	_cleanroomsConfiguredAudienceModelArn                   string
	_cleanroomsConfiguredAudienceModelAssociationIdentifier string
	_cleanroomsConfiguredAudienceModelAssociationName       string
	_cleanroomsConfiguredTableAssociationIdentifier         string
	_cleanroomsConfiguredTableIdentifier                    string
	_cleanroomsCreatorDisplayName                           string
	_cleanroomsCreatorMemberAbilities                       string
	_cleanroomsCreatorMLMemberAbilities                     string
	_cleanroomsCreatorPaymentConfiguration                  string
	_cleanroomsDataEncryptionMetadata                       string
	_cleanroomsDefaultJobResultConfiguration                string
	_cleanroomsDefaultResultConfiguration                   string
	_cleanroomsDescription                                  string
	_cleanroomsErrorMessageConfiguration                    string
	_cleanroomsFormat                                       string
	_cleanroomsIdMappingConfig                              string
	_cleanroomsIdMappingTableIdentifier                     string
	_cleanroomsIdNamespaceAssociationIdentifier             string
	_cleanroomsInputReferenceConfig                         string
	_cleanroomsIsMetricsEnabled                             string
	_cleanroomsJobLogStatus                                 string
	_cleanroomsJobParameters                                string
	_cleanroomsJobType                                      string
	_cleanroomsKmsKeyArn                                    string
	_cleanroomsManageResourcePolicies                       string
	_cleanroomsMaxResults                                   string
	_cleanroomsMemberStatus                                 string
	_cleanroomsMembers                                      string
	_cleanroomsMembershipIdentifier                         string
	_cleanroomsName                                         string
	_cleanroomsNames                                        []string
	_cleanroomsNextToken                                    string
	_cleanroomsParameters                                   string
	_cleanroomsPaymentConfiguration                         string
	_cleanroomsPrivacyBudgetTemplateIdentifier              string
	_cleanroomsPrivacyBudgetType                            string
	_cleanroomsProtectedJobIdentifier                       string
	_cleanroomsProtectedQueryIdentifier                     string
	_cleanroomsQueryLogStatus                               string
	_cleanroomsResourceArn                                  string
	_cleanroomsResultConfiguration                          string
	_cleanroomsRoleArn                                      string
	_cleanroomsSchema                                       string
	_cleanroomsSchemaAnalysisRuleRequests                   string
	_cleanroomsSchemaType                                   string
	_cleanroomsSelectedAnalysisMethods                      string
	_cleanroomsSource                                       string
	_cleanroomsSqlParameters                                string
	_cleanroomsStatus                                       string
	_cleanroomsSyntheticDataParameters                      string
	_cleanroomsTableReference                               string
	_cleanroomsTagKeys                                      []string
	_cleanroomsTags                                         string
	_cleanroomsTargetStatus                                 string
	_cleanroomsType                                         string
)

// Retrieves multiple analysis templates within a collaboration by their Amazon
// Resource Names (ARNs).
func cleanrooms_BatchGetCollaborationAnalysisTemplate(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.BatchGetCollaborationAnalysisTemplateInput{
		// AnalysisTemplateArns: []string, // Required
		// CollaborationIdentifier: *string, // Required
	}

	if len(_cleanroomsAnalysisTemplateArns) > 0 {
		input.AnalysisTemplateArns = append([]string(nil), _cleanroomsAnalysisTemplateArns...)
	}
	if len(_cleanroomsCollaborationIdentifier) > 0 {
		input.CollaborationIdentifier = aws.String(_cleanroomsCollaborationIdentifier)
	}

	if resp, err := client.BatchGetCollaborationAnalysisTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves multiple schemas by their identifiers.
func cleanrooms_BatchGetSchema(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.BatchGetSchemaInput{
		// CollaborationIdentifier: *string, // Required
		// Names: []string, // Required
	}

	if len(_cleanroomsCollaborationIdentifier) > 0 {
		input.CollaborationIdentifier = aws.String(_cleanroomsCollaborationIdentifier)
	}
	if len(_cleanroomsNames) > 0 {
		input.Names = append([]string(nil), _cleanroomsNames...)
	}

	if resp, err := client.BatchGetSchema(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves multiple analysis rule schemas.
func cleanrooms_BatchGetSchemaAnalysisRule(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.BatchGetSchemaAnalysisRuleInput{
		// CollaborationIdentifier: *string, // Required
		// SchemaAnalysisRuleRequests: []types.SchemaAnalysisRuleRequest, // Required
	}

	if len(_cleanroomsCollaborationIdentifier) > 0 {
		input.CollaborationIdentifier = aws.String(_cleanroomsCollaborationIdentifier)
	}
	if len(_cleanroomsSchemaAnalysisRuleRequests) > 0 {
		if err := assignInputField(input, "SchemaAnalysisRuleRequests", _cleanroomsSchemaAnalysisRuleRequests); err != nil {
			log.Errorf("invalid --schema-analysis-rule-requests: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchGetSchemaAnalysisRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new analysis template.
func cleanrooms_CreateAnalysisTemplate(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.CreateAnalysisTemplateInput{
		// Format: types.AnalysisFormat, // Required
		// MembershipIdentifier: *string, // Required
		// Name: *string, // Required
		// Source: types.AnalysisSource, // Required
	}

	if len(_cleanroomsFormat) > 0 {
		if err := assignInputField(input, "Format", _cleanroomsFormat); err != nil {
			log.Errorf("invalid --format: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsMembershipIdentifier)
	}
	if len(_cleanroomsName) > 0 {
		input.Name = aws.String(_cleanroomsName)
	}
	if len(_cleanroomsSource) > 0 {
		if err := assignInputField(input, "Source", _cleanroomsSource); err != nil {
			log.Errorf("invalid --source: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsAnalysisParameters) > 0 {
		if err := assignInputField(input, "AnalysisParameters", _cleanroomsAnalysisParameters); err != nil {
			log.Errorf("invalid --analysis-parameters: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsDescription) > 0 {
		input.Description = aws.String(_cleanroomsDescription)
	}
	if len(_cleanroomsErrorMessageConfiguration) > 0 {
		if err := assignInputField(input, "ErrorMessageConfiguration", _cleanroomsErrorMessageConfiguration); err != nil {
			log.Errorf("invalid --error-message-configuration: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsSchema) > 0 {
		if err := assignInputField(input, "Schema", _cleanroomsSchema); err != nil {
			log.Errorf("invalid --schema: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsSyntheticDataParameters) > 0 {
		if err := assignInputField(input, "SyntheticDataParameters", _cleanroomsSyntheticDataParameters); err != nil {
			log.Errorf("invalid --synthetic-data-parameters: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsTags) > 0 {
		if err := assignInputField(input, "Tags", _cleanroomsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAnalysisTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new collaboration.
func cleanrooms_CreateCollaboration(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.CreateCollaborationInput{
		// CreatorDisplayName: *string, // Required
		// CreatorMemberAbilities: []types.MemberAbility, // Required
		// Description: *string, // Required
		// Members: []types.MemberSpecification, // Required
		// Name: *string, // Required
		// QueryLogStatus: types.CollaborationQueryLogStatus, // Required
	}

	if len(_cleanroomsCreatorDisplayName) > 0 {
		input.CreatorDisplayName = aws.String(_cleanroomsCreatorDisplayName)
	}
	if len(_cleanroomsCreatorMemberAbilities) > 0 {
		if err := assignInputField(input, "CreatorMemberAbilities", _cleanroomsCreatorMemberAbilities); err != nil {
			log.Errorf("invalid --creator-member-abilities: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsDescription) > 0 {
		input.Description = aws.String(_cleanroomsDescription)
	}
	if len(_cleanroomsMembers) > 0 {
		if err := assignInputField(input, "Members", _cleanroomsMembers); err != nil {
			log.Errorf("invalid --members: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsName) > 0 {
		input.Name = aws.String(_cleanroomsName)
	}
	if len(_cleanroomsQueryLogStatus) > 0 {
		if err := assignInputField(input, "QueryLogStatus", _cleanroomsQueryLogStatus); err != nil {
			log.Errorf("invalid --query-log-status: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsAllowedResultRegions) > 0 {
		if err := assignInputField(input, "AllowedResultRegions", _cleanroomsAllowedResultRegions); err != nil {
			log.Errorf("invalid --allowed-result-regions: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsAnalyticsEngine) > 0 {
		if err := assignInputField(input, "AnalyticsEngine", _cleanroomsAnalyticsEngine); err != nil {
			log.Errorf("invalid --analytics-engine: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsAutoApprovedChangeRequestTypes) > 0 {
		if err := assignInputField(input, "AutoApprovedChangeRequestTypes", _cleanroomsAutoApprovedChangeRequestTypes); err != nil {
			log.Errorf("invalid --auto-approved-change-request-types: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsCreatorMLMemberAbilities) > 0 {
		if err := assignInputField(input, "CreatorMLMemberAbilities", _cleanroomsCreatorMLMemberAbilities); err != nil {
			log.Errorf("invalid --creator-ml-member-abilities: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsCreatorPaymentConfiguration) > 0 {
		if err := assignInputField(input, "CreatorPaymentConfiguration", _cleanroomsCreatorPaymentConfiguration); err != nil {
			log.Errorf("invalid --creator-payment-configuration: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsDataEncryptionMetadata) > 0 {
		if err := assignInputField(input, "DataEncryptionMetadata", _cleanroomsDataEncryptionMetadata); err != nil {
			log.Errorf("invalid --data-encryption-metadata: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsIsMetricsEnabled) > 0 {
		if err := assignInputField(input, "IsMetricsEnabled", _cleanroomsIsMetricsEnabled); err != nil {
			log.Errorf("invalid --is-metrics-enabled: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsJobLogStatus) > 0 {
		if err := assignInputField(input, "JobLogStatus", _cleanroomsJobLogStatus); err != nil {
			log.Errorf("invalid --job-log-status: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsTags) > 0 {
		if err := assignInputField(input, "Tags", _cleanroomsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCollaboration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new change request to modify an existing collaboration. This enables
// post-creation modifications to collaborations through a structured API-driven
// approach.
func cleanrooms_CreateCollaborationChangeRequest(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.CreateCollaborationChangeRequestInput{
		// Changes: []types.ChangeInput, // Required
		// CollaborationIdentifier: *string, // Required
	}

	if len(_cleanroomsChanges) > 0 {
		if err := assignInputField(input, "Changes", _cleanroomsChanges); err != nil {
			log.Errorf("invalid --changes: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsCollaborationIdentifier) > 0 {
		input.CollaborationIdentifier = aws.String(_cleanroomsCollaborationIdentifier)
	}

	if resp, err := client.CreateCollaborationChangeRequest(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides the details necessary to create a configured audience model
// association.
func cleanrooms_CreateConfiguredAudienceModelAssociation(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.CreateConfiguredAudienceModelAssociationInput{
		// ConfiguredAudienceModelArn: *string, // Required
		// ConfiguredAudienceModelAssociationName: *string, // Required
		// ManageResourcePolicies: *bool, // Required
		// MembershipIdentifier: *string, // Required
	}

	if len(_cleanroomsConfiguredAudienceModelArn) > 0 {
		input.ConfiguredAudienceModelArn = aws.String(_cleanroomsConfiguredAudienceModelArn)
	}
	if len(_cleanroomsConfiguredAudienceModelAssociationName) > 0 {
		input.ConfiguredAudienceModelAssociationName = aws.String(_cleanroomsConfiguredAudienceModelAssociationName)
	}
	if len(_cleanroomsManageResourcePolicies) > 0 {
		if err := assignInputField(input, "ManageResourcePolicies", _cleanroomsManageResourcePolicies); err != nil {
			log.Errorf("invalid --manage-resource-policies: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsMembershipIdentifier)
	}
	if len(_cleanroomsDescription) > 0 {
		input.Description = aws.String(_cleanroomsDescription)
	}
	if len(_cleanroomsTags) > 0 {
		if err := assignInputField(input, "Tags", _cleanroomsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateConfiguredAudienceModelAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new configured table resource.
func cleanrooms_CreateConfiguredTable(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.CreateConfiguredTableInput{
		// AllowedColumns: []string, // Required
		// AnalysisMethod: types.AnalysisMethod, // Required
		// Name: *string, // Required
		// TableReference: types.TableReference, // Required
	}

	if len(_cleanroomsAllowedColumns) > 0 {
		input.AllowedColumns = append([]string(nil), _cleanroomsAllowedColumns...)
	}
	if len(_cleanroomsAnalysisMethod) > 0 {
		if err := assignInputField(input, "AnalysisMethod", _cleanroomsAnalysisMethod); err != nil {
			log.Errorf("invalid --analysis-method: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsName) > 0 {
		input.Name = aws.String(_cleanroomsName)
	}
	if len(_cleanroomsTableReference) > 0 {
		if err := assignInputField(input, "TableReference", _cleanroomsTableReference); err != nil {
			log.Errorf("invalid --table-reference: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsDescription) > 0 {
		input.Description = aws.String(_cleanroomsDescription)
	}
	if len(_cleanroomsSelectedAnalysisMethods) > 0 {
		if err := assignInputField(input, "SelectedAnalysisMethods", _cleanroomsSelectedAnalysisMethods); err != nil {
			log.Errorf("invalid --selected-analysis-methods: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsTags) > 0 {
		if err := assignInputField(input, "Tags", _cleanroomsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateConfiguredTable(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new analysis rule for a configured table. Currently, only one
// analysis rule can be created for a given configured table.
func cleanrooms_CreateConfiguredTableAnalysisRule(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.CreateConfiguredTableAnalysisRuleInput{
		// AnalysisRulePolicy: types.ConfiguredTableAnalysisRulePolicy, // Required
		// AnalysisRuleType: types.ConfiguredTableAnalysisRuleType, // Required
		// ConfiguredTableIdentifier: *string, // Required
	}

	if len(_cleanroomsAnalysisRulePolicy) > 0 {
		if err := assignInputField(input, "AnalysisRulePolicy", _cleanroomsAnalysisRulePolicy); err != nil {
			log.Errorf("invalid --analysis-rule-policy: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsAnalysisRuleType) > 0 {
		if err := assignInputField(input, "AnalysisRuleType", _cleanroomsAnalysisRuleType); err != nil {
			log.Errorf("invalid --analysis-rule-type: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsConfiguredTableIdentifier) > 0 {
		input.ConfiguredTableIdentifier = aws.String(_cleanroomsConfiguredTableIdentifier)
	}

	if resp, err := client.CreateConfiguredTableAnalysisRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a configured table association. A configured table association links a
// configured table with a collaboration.
func cleanrooms_CreateConfiguredTableAssociation(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.CreateConfiguredTableAssociationInput{
		// ConfiguredTableIdentifier: *string, // Required
		// MembershipIdentifier: *string, // Required
		// Name: *string, // Required
		// RoleArn: *string, // Required
	}

	if len(_cleanroomsConfiguredTableIdentifier) > 0 {
		input.ConfiguredTableIdentifier = aws.String(_cleanroomsConfiguredTableIdentifier)
	}
	if len(_cleanroomsMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsMembershipIdentifier)
	}
	if len(_cleanroomsName) > 0 {
		input.Name = aws.String(_cleanroomsName)
	}
	if len(_cleanroomsRoleArn) > 0 {
		input.RoleArn = aws.String(_cleanroomsRoleArn)
	}
	if len(_cleanroomsDescription) > 0 {
		input.Description = aws.String(_cleanroomsDescription)
	}
	if len(_cleanroomsTags) > 0 {
		if err := assignInputField(input, "Tags", _cleanroomsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateConfiguredTableAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new analysis rule for an associated configured table.
func cleanrooms_CreateConfiguredTableAssociationAnalysisRule(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.CreateConfiguredTableAssociationAnalysisRuleInput{
		// AnalysisRulePolicy: types.ConfiguredTableAssociationAnalysisRulePolicy, // Required
		// AnalysisRuleType: types.ConfiguredTableAssociationAnalysisRuleType, // Required
		// ConfiguredTableAssociationIdentifier: *string, // Required
		// MembershipIdentifier: *string, // Required
	}

	if len(_cleanroomsAnalysisRulePolicy) > 0 {
		if err := assignInputField(input, "AnalysisRulePolicy", _cleanroomsAnalysisRulePolicy); err != nil {
			log.Errorf("invalid --analysis-rule-policy: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsAnalysisRuleType) > 0 {
		if err := assignInputField(input, "AnalysisRuleType", _cleanroomsAnalysisRuleType); err != nil {
			log.Errorf("invalid --analysis-rule-type: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsConfiguredTableAssociationIdentifier) > 0 {
		input.ConfiguredTableAssociationIdentifier = aws.String(_cleanroomsConfiguredTableAssociationIdentifier)
	}
	if len(_cleanroomsMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsMembershipIdentifier)
	}

	if resp, err := client.CreateConfiguredTableAssociationAnalysisRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an ID mapping table.
func cleanrooms_CreateIdMappingTable(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.CreateIdMappingTableInput{
		// InputReferenceConfig: *types.IdMappingTableInputReferenceConfig, // Required
		// MembershipIdentifier: *string, // Required
		// Name: *string, // Required
	}

	if len(_cleanroomsInputReferenceConfig) > 0 {
		if err := assignInputField(input, "InputReferenceConfig", _cleanroomsInputReferenceConfig); err != nil {
			log.Errorf("invalid --input-reference-config: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsMembershipIdentifier)
	}
	if len(_cleanroomsName) > 0 {
		input.Name = aws.String(_cleanroomsName)
	}
	if len(_cleanroomsDescription) > 0 {
		input.Description = aws.String(_cleanroomsDescription)
	}
	if len(_cleanroomsKmsKeyArn) > 0 {
		input.KmsKeyArn = aws.String(_cleanroomsKmsKeyArn)
	}
	if len(_cleanroomsTags) > 0 {
		if err := assignInputField(input, "Tags", _cleanroomsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateIdMappingTable(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an ID namespace association.
func cleanrooms_CreateIdNamespaceAssociation(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.CreateIdNamespaceAssociationInput{
		// InputReferenceConfig: *types.IdNamespaceAssociationInputReferenceConfig, // Required
		// MembershipIdentifier: *string, // Required
		// Name: *string, // Required
	}

	if len(_cleanroomsInputReferenceConfig) > 0 {
		if err := assignInputField(input, "InputReferenceConfig", _cleanroomsInputReferenceConfig); err != nil {
			log.Errorf("invalid --input-reference-config: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsMembershipIdentifier)
	}
	if len(_cleanroomsName) > 0 {
		input.Name = aws.String(_cleanroomsName)
	}
	if len(_cleanroomsDescription) > 0 {
		input.Description = aws.String(_cleanroomsDescription)
	}
	if len(_cleanroomsIdMappingConfig) > 0 {
		if err := assignInputField(input, "IdMappingConfig", _cleanroomsIdMappingConfig); err != nil {
			log.Errorf("invalid --id-mapping-config: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsTags) > 0 {
		if err := assignInputField(input, "Tags", _cleanroomsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateIdNamespaceAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a membership for a specific collaboration identifier and joins the
// collaboration.
func cleanrooms_CreateMembership(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.CreateMembershipInput{
		// CollaborationIdentifier: *string, // Required
		// QueryLogStatus: types.MembershipQueryLogStatus, // Required
	}

	if len(_cleanroomsCollaborationIdentifier) > 0 {
		input.CollaborationIdentifier = aws.String(_cleanroomsCollaborationIdentifier)
	}
	if len(_cleanroomsQueryLogStatus) > 0 {
		if err := assignInputField(input, "QueryLogStatus", _cleanroomsQueryLogStatus); err != nil {
			log.Errorf("invalid --query-log-status: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsDefaultJobResultConfiguration) > 0 {
		if err := assignInputField(input, "DefaultJobResultConfiguration", _cleanroomsDefaultJobResultConfiguration); err != nil {
			log.Errorf("invalid --default-job-result-configuration: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsDefaultResultConfiguration) > 0 {
		if err := assignInputField(input, "DefaultResultConfiguration", _cleanroomsDefaultResultConfiguration); err != nil {
			log.Errorf("invalid --default-result-configuration: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsIsMetricsEnabled) > 0 {
		if err := assignInputField(input, "IsMetricsEnabled", _cleanroomsIsMetricsEnabled); err != nil {
			log.Errorf("invalid --is-metrics-enabled: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsJobLogStatus) > 0 {
		if err := assignInputField(input, "JobLogStatus", _cleanroomsJobLogStatus); err != nil {
			log.Errorf("invalid --job-log-status: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsPaymentConfiguration) > 0 {
		if err := assignInputField(input, "PaymentConfiguration", _cleanroomsPaymentConfiguration); err != nil {
			log.Errorf("invalid --payment-configuration: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsTags) > 0 {
		if err := assignInputField(input, "Tags", _cleanroomsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateMembership(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a privacy budget template for a specified collaboration. Each
// collaboration can have only one privacy budget template. If you need to change
// the privacy budget template, use the UpdatePrivacyBudgetTemplateoperation.
func cleanrooms_CreatePrivacyBudgetTemplate(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.CreatePrivacyBudgetTemplateInput{
		// MembershipIdentifier: *string, // Required
		// Parameters: types.PrivacyBudgetTemplateParametersInput, // Required
		// PrivacyBudgetType: types.PrivacyBudgetType, // Required
	}

	if len(_cleanroomsMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsMembershipIdentifier)
	}
	if len(_cleanroomsParameters) > 0 {
		if err := assignInputField(input, "Parameters", _cleanroomsParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsPrivacyBudgetType) > 0 {
		if err := assignInputField(input, "PrivacyBudgetType", _cleanroomsPrivacyBudgetType); err != nil {
			log.Errorf("invalid --privacy-budget-type: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsAutoRefresh) > 0 {
		if err := assignInputField(input, "AutoRefresh", _cleanroomsAutoRefresh); err != nil {
			log.Errorf("invalid --auto-refresh: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsTags) > 0 {
		if err := assignInputField(input, "Tags", _cleanroomsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreatePrivacyBudgetTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an analysis template.
func cleanrooms_DeleteAnalysisTemplate(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.DeleteAnalysisTemplateInput{
		// AnalysisTemplateIdentifier: *string, // Required
		// MembershipIdentifier: *string, // Required
	}

	if len(_cleanroomsAnalysisTemplateIdentifier) > 0 {
		input.AnalysisTemplateIdentifier = aws.String(_cleanroomsAnalysisTemplateIdentifier)
	}
	if len(_cleanroomsMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsMembershipIdentifier)
	}

	if resp, err := client.DeleteAnalysisTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a collaboration. It can only be called by the collaboration owner.
func cleanrooms_DeleteCollaboration(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.DeleteCollaborationInput{
		// CollaborationIdentifier: *string, // Required
	}

	if len(_cleanroomsCollaborationIdentifier) > 0 {
		input.CollaborationIdentifier = aws.String(_cleanroomsCollaborationIdentifier)
	}

	if resp, err := client.DeleteCollaboration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides the information necessary to delete a configured audience model
// association.
func cleanrooms_DeleteConfiguredAudienceModelAssociation(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.DeleteConfiguredAudienceModelAssociationInput{
		// ConfiguredAudienceModelAssociationIdentifier: *string, // Required
		// MembershipIdentifier: *string, // Required
	}

	if len(_cleanroomsConfiguredAudienceModelAssociationIdentifier) > 0 {
		input.ConfiguredAudienceModelAssociationIdentifier = aws.String(_cleanroomsConfiguredAudienceModelAssociationIdentifier)
	}
	if len(_cleanroomsMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsMembershipIdentifier)
	}

	if resp, err := client.DeleteConfiguredAudienceModelAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a configured table.
func cleanrooms_DeleteConfiguredTable(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.DeleteConfiguredTableInput{
		// ConfiguredTableIdentifier: *string, // Required
	}

	if len(_cleanroomsConfiguredTableIdentifier) > 0 {
		input.ConfiguredTableIdentifier = aws.String(_cleanroomsConfiguredTableIdentifier)
	}

	if resp, err := client.DeleteConfiguredTable(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a configured table analysis rule.
func cleanrooms_DeleteConfiguredTableAnalysisRule(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.DeleteConfiguredTableAnalysisRuleInput{
		// AnalysisRuleType: types.ConfiguredTableAnalysisRuleType, // Required
		// ConfiguredTableIdentifier: *string, // Required
	}

	if len(_cleanroomsAnalysisRuleType) > 0 {
		if err := assignInputField(input, "AnalysisRuleType", _cleanroomsAnalysisRuleType); err != nil {
			log.Errorf("invalid --analysis-rule-type: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsConfiguredTableIdentifier) > 0 {
		input.ConfiguredTableIdentifier = aws.String(_cleanroomsConfiguredTableIdentifier)
	}

	if resp, err := client.DeleteConfiguredTableAnalysisRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a configured table association.
func cleanrooms_DeleteConfiguredTableAssociation(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.DeleteConfiguredTableAssociationInput{
		// ConfiguredTableAssociationIdentifier: *string, // Required
		// MembershipIdentifier: *string, // Required
	}

	if len(_cleanroomsConfiguredTableAssociationIdentifier) > 0 {
		input.ConfiguredTableAssociationIdentifier = aws.String(_cleanroomsConfiguredTableAssociationIdentifier)
	}
	if len(_cleanroomsMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsMembershipIdentifier)
	}

	if resp, err := client.DeleteConfiguredTableAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an analysis rule for a configured table association.
func cleanrooms_DeleteConfiguredTableAssociationAnalysisRule(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.DeleteConfiguredTableAssociationAnalysisRuleInput{
		// AnalysisRuleType: types.ConfiguredTableAssociationAnalysisRuleType, // Required
		// ConfiguredTableAssociationIdentifier: *string, // Required
		// MembershipIdentifier: *string, // Required
	}

	if len(_cleanroomsAnalysisRuleType) > 0 {
		if err := assignInputField(input, "AnalysisRuleType", _cleanroomsAnalysisRuleType); err != nil {
			log.Errorf("invalid --analysis-rule-type: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsConfiguredTableAssociationIdentifier) > 0 {
		input.ConfiguredTableAssociationIdentifier = aws.String(_cleanroomsConfiguredTableAssociationIdentifier)
	}
	if len(_cleanroomsMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsMembershipIdentifier)
	}

	if resp, err := client.DeleteConfiguredTableAssociationAnalysisRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an ID mapping table.
func cleanrooms_DeleteIdMappingTable(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.DeleteIdMappingTableInput{
		// IdMappingTableIdentifier: *string, // Required
		// MembershipIdentifier: *string, // Required
	}

	if len(_cleanroomsIdMappingTableIdentifier) > 0 {
		input.IdMappingTableIdentifier = aws.String(_cleanroomsIdMappingTableIdentifier)
	}
	if len(_cleanroomsMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsMembershipIdentifier)
	}

	if resp, err := client.DeleteIdMappingTable(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an ID namespace association.
func cleanrooms_DeleteIdNamespaceAssociation(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.DeleteIdNamespaceAssociationInput{
		// IdNamespaceAssociationIdentifier: *string, // Required
		// MembershipIdentifier: *string, // Required
	}

	if len(_cleanroomsIdNamespaceAssociationIdentifier) > 0 {
		input.IdNamespaceAssociationIdentifier = aws.String(_cleanroomsIdNamespaceAssociationIdentifier)
	}
	if len(_cleanroomsMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsMembershipIdentifier)
	}

	if resp, err := client.DeleteIdNamespaceAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the specified member from a collaboration. The removed member is placed
// in the Removed status and can't interact with the collaboration. The removed
// member's data is inaccessible to active members of the collaboration.
func cleanrooms_DeleteMember(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.DeleteMemberInput{
		// AccountId: *string, // Required
		// CollaborationIdentifier: *string, // Required
	}

	if len(_cleanroomsAccountId) > 0 {
		input.AccountId = aws.String(_cleanroomsAccountId)
	}
	if len(_cleanroomsCollaborationIdentifier) > 0 {
		input.CollaborationIdentifier = aws.String(_cleanroomsCollaborationIdentifier)
	}

	if resp, err := client.DeleteMember(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a specified membership. All resources under a membership must be
// deleted.
func cleanrooms_DeleteMembership(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.DeleteMembershipInput{
		// MembershipIdentifier: *string, // Required
	}

	if len(_cleanroomsMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsMembershipIdentifier)
	}

	if resp, err := client.DeleteMembership(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a privacy budget template for a specified collaboration.
func cleanrooms_DeletePrivacyBudgetTemplate(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.DeletePrivacyBudgetTemplateInput{
		// MembershipIdentifier: *string, // Required
		// PrivacyBudgetTemplateIdentifier: *string, // Required
	}

	if len(_cleanroomsMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsMembershipIdentifier)
	}
	if len(_cleanroomsPrivacyBudgetTemplateIdentifier) > 0 {
		input.PrivacyBudgetTemplateIdentifier = aws.String(_cleanroomsPrivacyBudgetTemplateIdentifier)
	}

	if resp, err := client.DeletePrivacyBudgetTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves an analysis template.
func cleanrooms_GetAnalysisTemplate(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.GetAnalysisTemplateInput{
		// AnalysisTemplateIdentifier: *string, // Required
		// MembershipIdentifier: *string, // Required
	}

	if len(_cleanroomsAnalysisTemplateIdentifier) > 0 {
		input.AnalysisTemplateIdentifier = aws.String(_cleanroomsAnalysisTemplateIdentifier)
	}
	if len(_cleanroomsMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsMembershipIdentifier)
	}

	if resp, err := client.GetAnalysisTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns metadata about a collaboration.
func cleanrooms_GetCollaboration(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.GetCollaborationInput{
		// CollaborationIdentifier: *string, // Required
	}

	if len(_cleanroomsCollaborationIdentifier) > 0 {
		input.CollaborationIdentifier = aws.String(_cleanroomsCollaborationIdentifier)
	}

	if resp, err := client.GetCollaboration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves an analysis template within a collaboration.
func cleanrooms_GetCollaborationAnalysisTemplate(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.GetCollaborationAnalysisTemplateInput{
		// AnalysisTemplateArn: *string, // Required
		// CollaborationIdentifier: *string, // Required
	}

	if len(_cleanroomsAnalysisTemplateArn) > 0 {
		input.AnalysisTemplateArn = aws.String(_cleanroomsAnalysisTemplateArn)
	}
	if len(_cleanroomsCollaborationIdentifier) > 0 {
		input.CollaborationIdentifier = aws.String(_cleanroomsCollaborationIdentifier)
	}

	if resp, err := client.GetCollaborationAnalysisTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves detailed information about a specific collaboration change request.
func cleanrooms_GetCollaborationChangeRequest(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.GetCollaborationChangeRequestInput{
		// ChangeRequestIdentifier: *string, // Required
		// CollaborationIdentifier: *string, // Required
	}

	if len(_cleanroomsChangeRequestIdentifier) > 0 {
		input.ChangeRequestIdentifier = aws.String(_cleanroomsChangeRequestIdentifier)
	}
	if len(_cleanroomsCollaborationIdentifier) > 0 {
		input.CollaborationIdentifier = aws.String(_cleanroomsCollaborationIdentifier)
	}

	if resp, err := client.GetCollaborationChangeRequest(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a configured audience model association within a collaboration.
func cleanrooms_GetCollaborationConfiguredAudienceModelAssociation(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.GetCollaborationConfiguredAudienceModelAssociationInput{
		// CollaborationIdentifier: *string, // Required
		// ConfiguredAudienceModelAssociationIdentifier: *string, // Required
	}

	if len(_cleanroomsCollaborationIdentifier) > 0 {
		input.CollaborationIdentifier = aws.String(_cleanroomsCollaborationIdentifier)
	}
	if len(_cleanroomsConfiguredAudienceModelAssociationIdentifier) > 0 {
		input.ConfiguredAudienceModelAssociationIdentifier = aws.String(_cleanroomsConfiguredAudienceModelAssociationIdentifier)
	}

	if resp, err := client.GetCollaborationConfiguredAudienceModelAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves an ID namespace association from a specific collaboration.
func cleanrooms_GetCollaborationIdNamespaceAssociation(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.GetCollaborationIdNamespaceAssociationInput{
		// CollaborationIdentifier: *string, // Required
		// IdNamespaceAssociationIdentifier: *string, // Required
	}

	if len(_cleanroomsCollaborationIdentifier) > 0 {
		input.CollaborationIdentifier = aws.String(_cleanroomsCollaborationIdentifier)
	}
	if len(_cleanroomsIdNamespaceAssociationIdentifier) > 0 {
		input.IdNamespaceAssociationIdentifier = aws.String(_cleanroomsIdNamespaceAssociationIdentifier)
	}

	if resp, err := client.GetCollaborationIdNamespaceAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns details about a specified privacy budget template.
func cleanrooms_GetCollaborationPrivacyBudgetTemplate(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.GetCollaborationPrivacyBudgetTemplateInput{
		// CollaborationIdentifier: *string, // Required
		// PrivacyBudgetTemplateIdentifier: *string, // Required
	}

	if len(_cleanroomsCollaborationIdentifier) > 0 {
		input.CollaborationIdentifier = aws.String(_cleanroomsCollaborationIdentifier)
	}
	if len(_cleanroomsPrivacyBudgetTemplateIdentifier) > 0 {
		input.PrivacyBudgetTemplateIdentifier = aws.String(_cleanroomsPrivacyBudgetTemplateIdentifier)
	}

	if resp, err := client.GetCollaborationPrivacyBudgetTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a configured audience model association.
func cleanrooms_GetConfiguredAudienceModelAssociation(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.GetConfiguredAudienceModelAssociationInput{
		// ConfiguredAudienceModelAssociationIdentifier: *string, // Required
		// MembershipIdentifier: *string, // Required
	}

	if len(_cleanroomsConfiguredAudienceModelAssociationIdentifier) > 0 {
		input.ConfiguredAudienceModelAssociationIdentifier = aws.String(_cleanroomsConfiguredAudienceModelAssociationIdentifier)
	}
	if len(_cleanroomsMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsMembershipIdentifier)
	}

	if resp, err := client.GetConfiguredAudienceModelAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a configured table.
func cleanrooms_GetConfiguredTable(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.GetConfiguredTableInput{
		// ConfiguredTableIdentifier: *string, // Required
	}

	if len(_cleanroomsConfiguredTableIdentifier) > 0 {
		input.ConfiguredTableIdentifier = aws.String(_cleanroomsConfiguredTableIdentifier)
	}

	if resp, err := client.GetConfiguredTable(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a configured table analysis rule.
func cleanrooms_GetConfiguredTableAnalysisRule(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.GetConfiguredTableAnalysisRuleInput{
		// AnalysisRuleType: types.ConfiguredTableAnalysisRuleType, // Required
		// ConfiguredTableIdentifier: *string, // Required
	}

	if len(_cleanroomsAnalysisRuleType) > 0 {
		if err := assignInputField(input, "AnalysisRuleType", _cleanroomsAnalysisRuleType); err != nil {
			log.Errorf("invalid --analysis-rule-type: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsConfiguredTableIdentifier) > 0 {
		input.ConfiguredTableIdentifier = aws.String(_cleanroomsConfiguredTableIdentifier)
	}

	if resp, err := client.GetConfiguredTableAnalysisRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a configured table association.
func cleanrooms_GetConfiguredTableAssociation(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.GetConfiguredTableAssociationInput{
		// ConfiguredTableAssociationIdentifier: *string, // Required
		// MembershipIdentifier: *string, // Required
	}

	if len(_cleanroomsConfiguredTableAssociationIdentifier) > 0 {
		input.ConfiguredTableAssociationIdentifier = aws.String(_cleanroomsConfiguredTableAssociationIdentifier)
	}
	if len(_cleanroomsMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsMembershipIdentifier)
	}

	if resp, err := client.GetConfiguredTableAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the analysis rule for a configured table association.
func cleanrooms_GetConfiguredTableAssociationAnalysisRule(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.GetConfiguredTableAssociationAnalysisRuleInput{
		// AnalysisRuleType: types.ConfiguredTableAssociationAnalysisRuleType, // Required
		// ConfiguredTableAssociationIdentifier: *string, // Required
		// MembershipIdentifier: *string, // Required
	}

	if len(_cleanroomsAnalysisRuleType) > 0 {
		if err := assignInputField(input, "AnalysisRuleType", _cleanroomsAnalysisRuleType); err != nil {
			log.Errorf("invalid --analysis-rule-type: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsConfiguredTableAssociationIdentifier) > 0 {
		input.ConfiguredTableAssociationIdentifier = aws.String(_cleanroomsConfiguredTableAssociationIdentifier)
	}
	if len(_cleanroomsMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsMembershipIdentifier)
	}

	if resp, err := client.GetConfiguredTableAssociationAnalysisRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves an ID mapping table.
func cleanrooms_GetIdMappingTable(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.GetIdMappingTableInput{
		// IdMappingTableIdentifier: *string, // Required
		// MembershipIdentifier: *string, // Required
	}

	if len(_cleanroomsIdMappingTableIdentifier) > 0 {
		input.IdMappingTableIdentifier = aws.String(_cleanroomsIdMappingTableIdentifier)
	}
	if len(_cleanroomsMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsMembershipIdentifier)
	}

	if resp, err := client.GetIdMappingTable(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves an ID namespace association.
func cleanrooms_GetIdNamespaceAssociation(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.GetIdNamespaceAssociationInput{
		// IdNamespaceAssociationIdentifier: *string, // Required
		// MembershipIdentifier: *string, // Required
	}

	if len(_cleanroomsIdNamespaceAssociationIdentifier) > 0 {
		input.IdNamespaceAssociationIdentifier = aws.String(_cleanroomsIdNamespaceAssociationIdentifier)
	}
	if len(_cleanroomsMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsMembershipIdentifier)
	}

	if resp, err := client.GetIdNamespaceAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a specified membership for an identifier.
func cleanrooms_GetMembership(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.GetMembershipInput{
		// MembershipIdentifier: *string, // Required
	}

	if len(_cleanroomsMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsMembershipIdentifier)
	}

	if resp, err := client.GetMembership(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns details for a specified privacy budget template.
func cleanrooms_GetPrivacyBudgetTemplate(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.GetPrivacyBudgetTemplateInput{
		// MembershipIdentifier: *string, // Required
		// PrivacyBudgetTemplateIdentifier: *string, // Required
	}

	if len(_cleanroomsMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsMembershipIdentifier)
	}
	if len(_cleanroomsPrivacyBudgetTemplateIdentifier) > 0 {
		input.PrivacyBudgetTemplateIdentifier = aws.String(_cleanroomsPrivacyBudgetTemplateIdentifier)
	}

	if resp, err := client.GetPrivacyBudgetTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns job processing metadata.
func cleanrooms_GetProtectedJob(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.GetProtectedJobInput{
		// MembershipIdentifier: *string, // Required
		// ProtectedJobIdentifier: *string, // Required
	}

	if len(_cleanroomsMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsMembershipIdentifier)
	}
	if len(_cleanroomsProtectedJobIdentifier) > 0 {
		input.ProtectedJobIdentifier = aws.String(_cleanroomsProtectedJobIdentifier)
	}

	if resp, err := client.GetProtectedJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns query processing metadata.
func cleanrooms_GetProtectedQuery(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.GetProtectedQueryInput{
		// MembershipIdentifier: *string, // Required
		// ProtectedQueryIdentifier: *string, // Required
	}

	if len(_cleanroomsMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsMembershipIdentifier)
	}
	if len(_cleanroomsProtectedQueryIdentifier) > 0 {
		input.ProtectedQueryIdentifier = aws.String(_cleanroomsProtectedQueryIdentifier)
	}

	if resp, err := client.GetProtectedQuery(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the schema for a relation within a collaboration.
func cleanrooms_GetSchema(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.GetSchemaInput{
		// CollaborationIdentifier: *string, // Required
		// Name: *string, // Required
	}

	if len(_cleanroomsCollaborationIdentifier) > 0 {
		input.CollaborationIdentifier = aws.String(_cleanroomsCollaborationIdentifier)
	}
	if len(_cleanroomsName) > 0 {
		input.Name = aws.String(_cleanroomsName)
	}

	if resp, err := client.GetSchema(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a schema analysis rule.
func cleanrooms_GetSchemaAnalysisRule(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.GetSchemaAnalysisRuleInput{
		// CollaborationIdentifier: *string, // Required
		// Name: *string, // Required
		// Type: types.AnalysisRuleType, // Required
	}

	if len(_cleanroomsCollaborationIdentifier) > 0 {
		input.CollaborationIdentifier = aws.String(_cleanroomsCollaborationIdentifier)
	}
	if len(_cleanroomsName) > 0 {
		input.Name = aws.String(_cleanroomsName)
	}
	if len(_cleanroomsType) > 0 {
		if err := assignInputField(input, "Type", _cleanroomsType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetSchemaAnalysisRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists analysis templates that the caller owns.
func cleanrooms_ListAnalysisTemplates(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.ListAnalysisTemplatesInput{
		// MembershipIdentifier: *string, // Required
	}

	if len(_cleanroomsMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsMembershipIdentifier)
	}
	if len(_cleanroomsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cleanroomsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsNextToken) > 0 {
		input.NextToken = aws.String(_cleanroomsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAnalysisTemplates(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cleanrooms.ListAnalysisTemplatesOutput
	p := cleanrooms.NewListAnalysisTemplatesPaginator(client, input)
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

// Lists analysis templates within a collaboration.
func cleanrooms_ListCollaborationAnalysisTemplates(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.ListCollaborationAnalysisTemplatesInput{
		// CollaborationIdentifier: *string, // Required
	}

	if len(_cleanroomsCollaborationIdentifier) > 0 {
		input.CollaborationIdentifier = aws.String(_cleanroomsCollaborationIdentifier)
	}
	if len(_cleanroomsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cleanroomsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsNextToken) > 0 {
		input.NextToken = aws.String(_cleanroomsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCollaborationAnalysisTemplates(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cleanrooms.ListCollaborationAnalysisTemplatesOutput
	p := cleanrooms.NewListCollaborationAnalysisTemplatesPaginator(client, input)
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

// Lists all change requests for a collaboration with pagination support. Returns
// change requests sorted by creation time.
func cleanrooms_ListCollaborationChangeRequests(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.ListCollaborationChangeRequestsInput{
		// CollaborationIdentifier: *string, // Required
	}

	if len(_cleanroomsCollaborationIdentifier) > 0 {
		input.CollaborationIdentifier = aws.String(_cleanroomsCollaborationIdentifier)
	}
	if len(_cleanroomsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cleanroomsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsNextToken) > 0 {
		input.NextToken = aws.String(_cleanroomsNextToken)
	}
	if len(_cleanroomsStatus) > 0 {
		if err := assignInputField(input, "Status", _cleanroomsStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListCollaborationChangeRequests(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cleanrooms.ListCollaborationChangeRequestsOutput
	p := cleanrooms.NewListCollaborationChangeRequestsPaginator(client, input)
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

// Lists configured audience model associations within a collaboration.
func cleanrooms_ListCollaborationConfiguredAudienceModelAssociations(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.ListCollaborationConfiguredAudienceModelAssociationsInput{
		// CollaborationIdentifier: *string, // Required
	}

	if len(_cleanroomsCollaborationIdentifier) > 0 {
		input.CollaborationIdentifier = aws.String(_cleanroomsCollaborationIdentifier)
	}
	if len(_cleanroomsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cleanroomsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsNextToken) > 0 {
		input.NextToken = aws.String(_cleanroomsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCollaborationConfiguredAudienceModelAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cleanrooms.ListCollaborationConfiguredAudienceModelAssociationsOutput
	p := cleanrooms.NewListCollaborationConfiguredAudienceModelAssociationsPaginator(client, input)
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

// Returns a list of the ID namespace associations in a collaboration.
func cleanrooms_ListCollaborationIdNamespaceAssociations(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.ListCollaborationIdNamespaceAssociationsInput{
		// CollaborationIdentifier: *string, // Required
	}

	if len(_cleanroomsCollaborationIdentifier) > 0 {
		input.CollaborationIdentifier = aws.String(_cleanroomsCollaborationIdentifier)
	}
	if len(_cleanroomsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cleanroomsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsNextToken) > 0 {
		input.NextToken = aws.String(_cleanroomsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCollaborationIdNamespaceAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cleanrooms.ListCollaborationIdNamespaceAssociationsOutput
	p := cleanrooms.NewListCollaborationIdNamespaceAssociationsPaginator(client, input)
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

// Returns an array that summarizes each privacy budget template in a specified
// collaboration.
func cleanrooms_ListCollaborationPrivacyBudgetTemplates(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.ListCollaborationPrivacyBudgetTemplatesInput{
		// CollaborationIdentifier: *string, // Required
	}

	if len(_cleanroomsCollaborationIdentifier) > 0 {
		input.CollaborationIdentifier = aws.String(_cleanroomsCollaborationIdentifier)
	}
	if len(_cleanroomsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cleanroomsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsNextToken) > 0 {
		input.NextToken = aws.String(_cleanroomsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCollaborationPrivacyBudgetTemplates(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cleanrooms.ListCollaborationPrivacyBudgetTemplatesOutput
	p := cleanrooms.NewListCollaborationPrivacyBudgetTemplatesPaginator(client, input)
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

// Returns an array that summarizes each privacy budget in a specified
// collaboration. The summary includes the collaboration ARN, creation time,
// creating account, and privacy budget details.
func cleanrooms_ListCollaborationPrivacyBudgets(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.ListCollaborationPrivacyBudgetsInput{
		// CollaborationIdentifier: *string, // Required
		// PrivacyBudgetType: types.PrivacyBudgetType, // Required
	}

	if len(_cleanroomsCollaborationIdentifier) > 0 {
		input.CollaborationIdentifier = aws.String(_cleanroomsCollaborationIdentifier)
	}
	if len(_cleanroomsPrivacyBudgetType) > 0 {
		if err := assignInputField(input, "PrivacyBudgetType", _cleanroomsPrivacyBudgetType); err != nil {
			log.Errorf("invalid --privacy-budget-type: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsAccessBudgetResourceArn) > 0 {
		input.AccessBudgetResourceArn = aws.String(_cleanroomsAccessBudgetResourceArn)
	}
	if len(_cleanroomsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cleanroomsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsNextToken) > 0 {
		input.NextToken = aws.String(_cleanroomsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCollaborationPrivacyBudgets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cleanrooms.ListCollaborationPrivacyBudgetsOutput
	p := cleanrooms.NewListCollaborationPrivacyBudgetsPaginator(client, input)
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

// Lists collaborations the caller owns, is active in, or has been invited to.
func cleanrooms_ListCollaborations(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.ListCollaborationsInput{}

	if len(_cleanroomsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cleanroomsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsMemberStatus) > 0 {
		if err := assignInputField(input, "MemberStatus", _cleanroomsMemberStatus); err != nil {
			log.Errorf("invalid --member-status: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsNextToken) > 0 {
		input.NextToken = aws.String(_cleanroomsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCollaborations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cleanrooms.ListCollaborationsOutput
	p := cleanrooms.NewListCollaborationsPaginator(client, input)
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

// Lists information about requested configured audience model associations.
func cleanrooms_ListConfiguredAudienceModelAssociations(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.ListConfiguredAudienceModelAssociationsInput{
		// MembershipIdentifier: *string, // Required
	}

	if len(_cleanroomsMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsMembershipIdentifier)
	}
	if len(_cleanroomsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cleanroomsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsNextToken) > 0 {
		input.NextToken = aws.String(_cleanroomsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListConfiguredAudienceModelAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cleanrooms.ListConfiguredAudienceModelAssociationsOutput
	p := cleanrooms.NewListConfiguredAudienceModelAssociationsPaginator(client, input)
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

// Lists configured table associations for a membership.
func cleanrooms_ListConfiguredTableAssociations(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.ListConfiguredTableAssociationsInput{
		// MembershipIdentifier: *string, // Required
	}

	if len(_cleanroomsMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsMembershipIdentifier)
	}
	if len(_cleanroomsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cleanroomsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsNextToken) > 0 {
		input.NextToken = aws.String(_cleanroomsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListConfiguredTableAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cleanrooms.ListConfiguredTableAssociationsOutput
	p := cleanrooms.NewListConfiguredTableAssociationsPaginator(client, input)
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

// Lists configured tables.
func cleanrooms_ListConfiguredTables(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.ListConfiguredTablesInput{}

	if len(_cleanroomsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cleanroomsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsNextToken) > 0 {
		input.NextToken = aws.String(_cleanroomsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListConfiguredTables(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cleanrooms.ListConfiguredTablesOutput
	p := cleanrooms.NewListConfiguredTablesPaginator(client, input)
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

// Returns a list of ID mapping tables.
func cleanrooms_ListIdMappingTables(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.ListIdMappingTablesInput{
		// MembershipIdentifier: *string, // Required
	}

	if len(_cleanroomsMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsMembershipIdentifier)
	}
	if len(_cleanroomsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cleanroomsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsNextToken) > 0 {
		input.NextToken = aws.String(_cleanroomsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListIdMappingTables(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cleanrooms.ListIdMappingTablesOutput
	p := cleanrooms.NewListIdMappingTablesPaginator(client, input)
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

// Returns a list of ID namespace associations.
func cleanrooms_ListIdNamespaceAssociations(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.ListIdNamespaceAssociationsInput{
		// MembershipIdentifier: *string, // Required
	}

	if len(_cleanroomsMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsMembershipIdentifier)
	}
	if len(_cleanroomsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cleanroomsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsNextToken) > 0 {
		input.NextToken = aws.String(_cleanroomsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListIdNamespaceAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cleanrooms.ListIdNamespaceAssociationsOutput
	p := cleanrooms.NewListIdNamespaceAssociationsPaginator(client, input)
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

// Lists all members within a collaboration.
func cleanrooms_ListMembers(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.ListMembersInput{
		// CollaborationIdentifier: *string, // Required
	}

	if len(_cleanroomsCollaborationIdentifier) > 0 {
		input.CollaborationIdentifier = aws.String(_cleanroomsCollaborationIdentifier)
	}
	if len(_cleanroomsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cleanroomsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsNextToken) > 0 {
		input.NextToken = aws.String(_cleanroomsNextToken)
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

	var results []*cleanrooms.ListMembersOutput
	p := cleanrooms.NewListMembersPaginator(client, input)
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

// Lists all memberships resources within the caller's account.
func cleanrooms_ListMemberships(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.ListMembershipsInput{}

	if len(_cleanroomsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cleanroomsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsNextToken) > 0 {
		input.NextToken = aws.String(_cleanroomsNextToken)
	}
	if len(_cleanroomsStatus) > 0 {
		if err := assignInputField(input, "Status", _cleanroomsStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListMemberships(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cleanrooms.ListMembershipsOutput
	p := cleanrooms.NewListMembershipsPaginator(client, input)
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

// Returns detailed information about the privacy budget templates in a specified
// membership.
func cleanrooms_ListPrivacyBudgetTemplates(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.ListPrivacyBudgetTemplatesInput{
		// MembershipIdentifier: *string, // Required
	}

	if len(_cleanroomsMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsMembershipIdentifier)
	}
	if len(_cleanroomsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cleanroomsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsNextToken) > 0 {
		input.NextToken = aws.String(_cleanroomsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPrivacyBudgetTemplates(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cleanrooms.ListPrivacyBudgetTemplatesOutput
	p := cleanrooms.NewListPrivacyBudgetTemplatesPaginator(client, input)
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

// Returns detailed information about the privacy budgets in a specified
// membership.
func cleanrooms_ListPrivacyBudgets(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.ListPrivacyBudgetsInput{
		// MembershipIdentifier: *string, // Required
		// PrivacyBudgetType: types.PrivacyBudgetType, // Required
	}

	if len(_cleanroomsMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsMembershipIdentifier)
	}
	if len(_cleanroomsPrivacyBudgetType) > 0 {
		if err := assignInputField(input, "PrivacyBudgetType", _cleanroomsPrivacyBudgetType); err != nil {
			log.Errorf("invalid --privacy-budget-type: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsAccessBudgetResourceArn) > 0 {
		input.AccessBudgetResourceArn = aws.String(_cleanroomsAccessBudgetResourceArn)
	}
	if len(_cleanroomsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cleanroomsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsNextToken) > 0 {
		input.NextToken = aws.String(_cleanroomsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPrivacyBudgets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cleanrooms.ListPrivacyBudgetsOutput
	p := cleanrooms.NewListPrivacyBudgetsPaginator(client, input)
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

// Lists protected jobs, sorted by most recent job.
func cleanrooms_ListProtectedJobs(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.ListProtectedJobsInput{
		// MembershipIdentifier: *string, // Required
	}

	if len(_cleanroomsMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsMembershipIdentifier)
	}
	if len(_cleanroomsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cleanroomsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsNextToken) > 0 {
		input.NextToken = aws.String(_cleanroomsNextToken)
	}
	if len(_cleanroomsStatus) > 0 {
		if err := assignInputField(input, "Status", _cleanroomsStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListProtectedJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cleanrooms.ListProtectedJobsOutput
	p := cleanrooms.NewListProtectedJobsPaginator(client, input)
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

// Lists protected queries, sorted by the most recent query.
func cleanrooms_ListProtectedQueries(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.ListProtectedQueriesInput{
		// MembershipIdentifier: *string, // Required
	}

	if len(_cleanroomsMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsMembershipIdentifier)
	}
	if len(_cleanroomsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cleanroomsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsNextToken) > 0 {
		input.NextToken = aws.String(_cleanroomsNextToken)
	}
	if len(_cleanroomsStatus) > 0 {
		if err := assignInputField(input, "Status", _cleanroomsStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListProtectedQueries(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cleanrooms.ListProtectedQueriesOutput
	p := cleanrooms.NewListProtectedQueriesPaginator(client, input)
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

// Lists the schemas for relations within a collaboration.
func cleanrooms_ListSchemas(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.ListSchemasInput{
		// CollaborationIdentifier: *string, // Required
	}

	if len(_cleanroomsCollaborationIdentifier) > 0 {
		input.CollaborationIdentifier = aws.String(_cleanroomsCollaborationIdentifier)
	}
	if len(_cleanroomsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cleanroomsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsNextToken) > 0 {
		input.NextToken = aws.String(_cleanroomsNextToken)
	}
	if len(_cleanroomsSchemaType) > 0 {
		if err := assignInputField(input, "SchemaType", _cleanroomsSchemaType); err != nil {
			log.Errorf("invalid --schema-type: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListSchemas(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cleanrooms.ListSchemasOutput
	p := cleanrooms.NewListSchemasPaginator(client, input)
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

// Lists all of the tags that have been added to a resource.
func cleanrooms_ListTagsForResource(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_cleanroomsResourceArn) > 0 {
		input.ResourceArn = aws.String(_cleanroomsResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Defines the information that's necessary to populate an ID mapping table.
func cleanrooms_PopulateIdMappingTable(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.PopulateIdMappingTableInput{
		// IdMappingTableIdentifier: *string, // Required
		// MembershipIdentifier: *string, // Required
	}

	if len(_cleanroomsIdMappingTableIdentifier) > 0 {
		input.IdMappingTableIdentifier = aws.String(_cleanroomsIdMappingTableIdentifier)
	}
	if len(_cleanroomsMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsMembershipIdentifier)
	}
	if len(_cleanroomsJobType) > 0 {
		if err := assignInputField(input, "JobType", _cleanroomsJobType); err != nil {
			log.Errorf("invalid --job-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.PopulateIdMappingTable(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// An estimate of the number of aggregation functions that the member who can
// query can run given epsilon and noise parameters.
func cleanrooms_PreviewPrivacyImpact(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.PreviewPrivacyImpactInput{
		// MembershipIdentifier: *string, // Required
		// Parameters: types.PreviewPrivacyImpactParametersInput, // Required
	}

	if len(_cleanroomsMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsMembershipIdentifier)
	}
	if len(_cleanroomsParameters) > 0 {
		if err := assignInputField(input, "Parameters", _cleanroomsParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}

	if resp, err := client.PreviewPrivacyImpact(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a protected job that is started by Clean Rooms.
func cleanrooms_StartProtectedJob(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.StartProtectedJobInput{
		// JobParameters: *types.ProtectedJobParameters, // Required
		// MembershipIdentifier: *string, // Required
		// Type: types.ProtectedJobType, // Required
	}

	if len(_cleanroomsJobParameters) > 0 {
		if err := assignInputField(input, "JobParameters", _cleanroomsJobParameters); err != nil {
			log.Errorf("invalid --job-parameters: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsMembershipIdentifier)
	}
	if len(_cleanroomsType) > 0 {
		if err := assignInputField(input, "Type", _cleanroomsType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsComputeConfiguration) > 0 {
		if err := assignInputField(input, "ComputeConfiguration", _cleanroomsComputeConfiguration); err != nil {
			log.Errorf("invalid --compute-configuration: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsResultConfiguration) > 0 {
		if err := assignInputField(input, "ResultConfiguration", _cleanroomsResultConfiguration); err != nil {
			log.Errorf("invalid --result-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartProtectedJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a protected query that is started by Clean Rooms.
func cleanrooms_StartProtectedQuery(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.StartProtectedQueryInput{
		// MembershipIdentifier: *string, // Required
		// SqlParameters: *types.ProtectedQuerySQLParameters, // Required
		// Type: types.ProtectedQueryType, // Required
	}

	if len(_cleanroomsMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsMembershipIdentifier)
	}
	if len(_cleanroomsSqlParameters) > 0 {
		if err := assignInputField(input, "SqlParameters", _cleanroomsSqlParameters); err != nil {
			log.Errorf("invalid --sql-parameters: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsType) > 0 {
		if err := assignInputField(input, "Type", _cleanroomsType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsComputeConfiguration) > 0 {
		if err := assignInputField(input, "ComputeConfiguration", _cleanroomsComputeConfiguration); err != nil {
			log.Errorf("invalid --compute-configuration: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsResultConfiguration) > 0 {
		if err := assignInputField(input, "ResultConfiguration", _cleanroomsResultConfiguration); err != nil {
			log.Errorf("invalid --result-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartProtectedQuery(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Tags a resource.
func cleanrooms_TagResource(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_cleanroomsResourceArn) > 0 {
		input.ResourceArn = aws.String(_cleanroomsResourceArn)
	}
	if len(_cleanroomsTags) > 0 {
		if err := assignInputField(input, "Tags", _cleanroomsTags); err != nil {
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

// Removes a tag or list of tags from a resource.
func cleanrooms_UntagResource(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_cleanroomsResourceArn) > 0 {
		input.ResourceArn = aws.String(_cleanroomsResourceArn)
	}
	if len(_cleanroomsTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _cleanroomsTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the analysis template metadata.
func cleanrooms_UpdateAnalysisTemplate(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.UpdateAnalysisTemplateInput{
		// AnalysisTemplateIdentifier: *string, // Required
		// MembershipIdentifier: *string, // Required
	}

	if len(_cleanroomsAnalysisTemplateIdentifier) > 0 {
		input.AnalysisTemplateIdentifier = aws.String(_cleanroomsAnalysisTemplateIdentifier)
	}
	if len(_cleanroomsMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsMembershipIdentifier)
	}
	if len(_cleanroomsDescription) > 0 {
		input.Description = aws.String(_cleanroomsDescription)
	}

	if resp, err := client.UpdateAnalysisTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates collaboration metadata and can only be called by the collaboration
// owner.
func cleanrooms_UpdateCollaboration(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.UpdateCollaborationInput{
		// CollaborationIdentifier: *string, // Required
	}

	if len(_cleanroomsCollaborationIdentifier) > 0 {
		input.CollaborationIdentifier = aws.String(_cleanroomsCollaborationIdentifier)
	}
	if len(_cleanroomsAnalyticsEngine) > 0 {
		if err := assignInputField(input, "AnalyticsEngine", _cleanroomsAnalyticsEngine); err != nil {
			log.Errorf("invalid --analytics-engine: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsDescription) > 0 {
		input.Description = aws.String(_cleanroomsDescription)
	}
	if len(_cleanroomsName) > 0 {
		input.Name = aws.String(_cleanroomsName)
	}

	if resp, err := client.UpdateCollaboration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing collaboration change request. This operation allows
// approval actions for pending change requests in collaborations (APPROVE, DENY,
// CANCEL, COMMIT).
//
// For change requests without automatic approval, a member in the collaboration
// can manually APPROVE or DENY a change request. The collaboration owner can
// manually CANCEL or COMMIT a change request.
func cleanrooms_UpdateCollaborationChangeRequest(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.UpdateCollaborationChangeRequestInput{
		// Action: types.ChangeRequestAction, // Required
		// ChangeRequestIdentifier: *string, // Required
		// CollaborationIdentifier: *string, // Required
	}

	if len(_cleanroomsAction) > 0 {
		if err := assignInputField(input, "Action", _cleanroomsAction); err != nil {
			log.Errorf("invalid --action: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsChangeRequestIdentifier) > 0 {
		input.ChangeRequestIdentifier = aws.String(_cleanroomsChangeRequestIdentifier)
	}
	if len(_cleanroomsCollaborationIdentifier) > 0 {
		input.CollaborationIdentifier = aws.String(_cleanroomsCollaborationIdentifier)
	}

	if resp, err := client.UpdateCollaborationChangeRequest(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides the details necessary to update a configured audience model
// association.
func cleanrooms_UpdateConfiguredAudienceModelAssociation(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.UpdateConfiguredAudienceModelAssociationInput{
		// ConfiguredAudienceModelAssociationIdentifier: *string, // Required
		// MembershipIdentifier: *string, // Required
	}

	if len(_cleanroomsConfiguredAudienceModelAssociationIdentifier) > 0 {
		input.ConfiguredAudienceModelAssociationIdentifier = aws.String(_cleanroomsConfiguredAudienceModelAssociationIdentifier)
	}
	if len(_cleanroomsMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsMembershipIdentifier)
	}
	if len(_cleanroomsDescription) > 0 {
		input.Description = aws.String(_cleanroomsDescription)
	}
	if len(_cleanroomsName) > 0 {
		input.Name = aws.String(_cleanroomsName)
	}

	if resp, err := client.UpdateConfiguredAudienceModelAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a configured table.
func cleanrooms_UpdateConfiguredTable(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.UpdateConfiguredTableInput{
		// ConfiguredTableIdentifier: *string, // Required
	}

	if len(_cleanroomsConfiguredTableIdentifier) > 0 {
		input.ConfiguredTableIdentifier = aws.String(_cleanroomsConfiguredTableIdentifier)
	}
	if len(_cleanroomsAllowedColumns) > 0 {
		input.AllowedColumns = append([]string(nil), _cleanroomsAllowedColumns...)
	}
	if len(_cleanroomsAnalysisMethod) > 0 {
		if err := assignInputField(input, "AnalysisMethod", _cleanroomsAnalysisMethod); err != nil {
			log.Errorf("invalid --analysis-method: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsDescription) > 0 {
		input.Description = aws.String(_cleanroomsDescription)
	}
	if len(_cleanroomsName) > 0 {
		input.Name = aws.String(_cleanroomsName)
	}
	if len(_cleanroomsSelectedAnalysisMethods) > 0 {
		if err := assignInputField(input, "SelectedAnalysisMethods", _cleanroomsSelectedAnalysisMethods); err != nil {
			log.Errorf("invalid --selected-analysis-methods: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsTableReference) > 0 {
		if err := assignInputField(input, "TableReference", _cleanroomsTableReference); err != nil {
			log.Errorf("invalid --table-reference: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateConfiguredTable(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a configured table analysis rule.
func cleanrooms_UpdateConfiguredTableAnalysisRule(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.UpdateConfiguredTableAnalysisRuleInput{
		// AnalysisRulePolicy: types.ConfiguredTableAnalysisRulePolicy, // Required
		// AnalysisRuleType: types.ConfiguredTableAnalysisRuleType, // Required
		// ConfiguredTableIdentifier: *string, // Required
	}

	if len(_cleanroomsAnalysisRulePolicy) > 0 {
		if err := assignInputField(input, "AnalysisRulePolicy", _cleanroomsAnalysisRulePolicy); err != nil {
			log.Errorf("invalid --analysis-rule-policy: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsAnalysisRuleType) > 0 {
		if err := assignInputField(input, "AnalysisRuleType", _cleanroomsAnalysisRuleType); err != nil {
			log.Errorf("invalid --analysis-rule-type: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsConfiguredTableIdentifier) > 0 {
		input.ConfiguredTableIdentifier = aws.String(_cleanroomsConfiguredTableIdentifier)
	}

	if resp, err := client.UpdateConfiguredTableAnalysisRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a configured table association.
func cleanrooms_UpdateConfiguredTableAssociation(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.UpdateConfiguredTableAssociationInput{
		// ConfiguredTableAssociationIdentifier: *string, // Required
		// MembershipIdentifier: *string, // Required
	}

	if len(_cleanroomsConfiguredTableAssociationIdentifier) > 0 {
		input.ConfiguredTableAssociationIdentifier = aws.String(_cleanroomsConfiguredTableAssociationIdentifier)
	}
	if len(_cleanroomsMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsMembershipIdentifier)
	}
	if len(_cleanroomsDescription) > 0 {
		input.Description = aws.String(_cleanroomsDescription)
	}
	if len(_cleanroomsRoleArn) > 0 {
		input.RoleArn = aws.String(_cleanroomsRoleArn)
	}

	if resp, err := client.UpdateConfiguredTableAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the analysis rule for a configured table association.
func cleanrooms_UpdateConfiguredTableAssociationAnalysisRule(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.UpdateConfiguredTableAssociationAnalysisRuleInput{
		// AnalysisRulePolicy: types.ConfiguredTableAssociationAnalysisRulePolicy, // Required
		// AnalysisRuleType: types.ConfiguredTableAssociationAnalysisRuleType, // Required
		// ConfiguredTableAssociationIdentifier: *string, // Required
		// MembershipIdentifier: *string, // Required
	}

	if len(_cleanroomsAnalysisRulePolicy) > 0 {
		if err := assignInputField(input, "AnalysisRulePolicy", _cleanroomsAnalysisRulePolicy); err != nil {
			log.Errorf("invalid --analysis-rule-policy: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsAnalysisRuleType) > 0 {
		if err := assignInputField(input, "AnalysisRuleType", _cleanroomsAnalysisRuleType); err != nil {
			log.Errorf("invalid --analysis-rule-type: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsConfiguredTableAssociationIdentifier) > 0 {
		input.ConfiguredTableAssociationIdentifier = aws.String(_cleanroomsConfiguredTableAssociationIdentifier)
	}
	if len(_cleanroomsMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsMembershipIdentifier)
	}

	if resp, err := client.UpdateConfiguredTableAssociationAnalysisRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides the details that are necessary to update an ID mapping table.
func cleanrooms_UpdateIdMappingTable(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.UpdateIdMappingTableInput{
		// IdMappingTableIdentifier: *string, // Required
		// MembershipIdentifier: *string, // Required
	}

	if len(_cleanroomsIdMappingTableIdentifier) > 0 {
		input.IdMappingTableIdentifier = aws.String(_cleanroomsIdMappingTableIdentifier)
	}
	if len(_cleanroomsMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsMembershipIdentifier)
	}
	if len(_cleanroomsDescription) > 0 {
		input.Description = aws.String(_cleanroomsDescription)
	}
	if len(_cleanroomsKmsKeyArn) > 0 {
		input.KmsKeyArn = aws.String(_cleanroomsKmsKeyArn)
	}

	if resp, err := client.UpdateIdMappingTable(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides the details that are necessary to update an ID namespace association.
func cleanrooms_UpdateIdNamespaceAssociation(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.UpdateIdNamespaceAssociationInput{
		// IdNamespaceAssociationIdentifier: *string, // Required
		// MembershipIdentifier: *string, // Required
	}

	if len(_cleanroomsIdNamespaceAssociationIdentifier) > 0 {
		input.IdNamespaceAssociationIdentifier = aws.String(_cleanroomsIdNamespaceAssociationIdentifier)
	}
	if len(_cleanroomsMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsMembershipIdentifier)
	}
	if len(_cleanroomsDescription) > 0 {
		input.Description = aws.String(_cleanroomsDescription)
	}
	if len(_cleanroomsIdMappingConfig) > 0 {
		if err := assignInputField(input, "IdMappingConfig", _cleanroomsIdMappingConfig); err != nil {
			log.Errorf("invalid --id-mapping-config: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsName) > 0 {
		input.Name = aws.String(_cleanroomsName)
	}

	if resp, err := client.UpdateIdNamespaceAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a membership.
func cleanrooms_UpdateMembership(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.UpdateMembershipInput{
		// MembershipIdentifier: *string, // Required
	}

	if len(_cleanroomsMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsMembershipIdentifier)
	}
	if len(_cleanroomsDefaultJobResultConfiguration) > 0 {
		if err := assignInputField(input, "DefaultJobResultConfiguration", _cleanroomsDefaultJobResultConfiguration); err != nil {
			log.Errorf("invalid --default-job-result-configuration: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsDefaultResultConfiguration) > 0 {
		if err := assignInputField(input, "DefaultResultConfiguration", _cleanroomsDefaultResultConfiguration); err != nil {
			log.Errorf("invalid --default-result-configuration: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsJobLogStatus) > 0 {
		if err := assignInputField(input, "JobLogStatus", _cleanroomsJobLogStatus); err != nil {
			log.Errorf("invalid --job-log-status: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsQueryLogStatus) > 0 {
		if err := assignInputField(input, "QueryLogStatus", _cleanroomsQueryLogStatus); err != nil {
			log.Errorf("invalid --query-log-status: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateMembership(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the privacy budget template for the specified collaboration.
func cleanrooms_UpdatePrivacyBudgetTemplate(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.UpdatePrivacyBudgetTemplateInput{
		// MembershipIdentifier: *string, // Required
		// PrivacyBudgetTemplateIdentifier: *string, // Required
		// PrivacyBudgetType: types.PrivacyBudgetType, // Required
	}

	if len(_cleanroomsMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsMembershipIdentifier)
	}
	if len(_cleanroomsPrivacyBudgetTemplateIdentifier) > 0 {
		input.PrivacyBudgetTemplateIdentifier = aws.String(_cleanroomsPrivacyBudgetTemplateIdentifier)
	}
	if len(_cleanroomsPrivacyBudgetType) > 0 {
		if err := assignInputField(input, "PrivacyBudgetType", _cleanroomsPrivacyBudgetType); err != nil {
			log.Errorf("invalid --privacy-budget-type: %s", err.Error())
			return
		}
	}
	if len(_cleanroomsParameters) > 0 {
		if err := assignInputField(input, "Parameters", _cleanroomsParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdatePrivacyBudgetTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the processing of a currently running job.
func cleanrooms_UpdateProtectedJob(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.UpdateProtectedJobInput{
		// MembershipIdentifier: *string, // Required
		// ProtectedJobIdentifier: *string, // Required
		// TargetStatus: types.TargetProtectedJobStatus, // Required
	}

	if len(_cleanroomsMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsMembershipIdentifier)
	}
	if len(_cleanroomsProtectedJobIdentifier) > 0 {
		input.ProtectedJobIdentifier = aws.String(_cleanroomsProtectedJobIdentifier)
	}
	if len(_cleanroomsTargetStatus) > 0 {
		if err := assignInputField(input, "TargetStatus", _cleanroomsTargetStatus); err != nil {
			log.Errorf("invalid --target-status: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateProtectedJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the processing of a currently running query.
func cleanrooms_UpdateProtectedQuery(cfg aws.Config, client *cleanrooms.Client) {
	input := &cleanrooms.UpdateProtectedQueryInput{
		// MembershipIdentifier: *string, // Required
		// ProtectedQueryIdentifier: *string, // Required
		// TargetStatus: types.TargetProtectedQueryStatus, // Required
	}

	if len(_cleanroomsMembershipIdentifier) > 0 {
		input.MembershipIdentifier = aws.String(_cleanroomsMembershipIdentifier)
	}
	if len(_cleanroomsProtectedQueryIdentifier) > 0 {
		input.ProtectedQueryIdentifier = aws.String(_cleanroomsProtectedQueryIdentifier)
	}
	if len(_cleanroomsTargetStatus) > 0 {
		if err := assignInputField(input, "TargetStatus", _cleanroomsTargetStatus); err != nil {
			log.Errorf("invalid --target-status: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateProtectedQuery(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_cleanroomsCmd)
	_cleanroomsCmd.Flags().SortFlags = false

	_cleanroomsCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_cleanroomsCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_cleanroomsCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsAccessBudgetResourceArn, "access-budget-resource-arn", "", "", "Access Budget Resource ARN")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsAccountId, "account-id", "", "", "Account ID")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsAction, "action", "", "", "Action")
	_cleanroomsCmd.Flags().StringSliceVarP(&_cleanroomsAllowedColumns, "allowed-columns", "", nil, "Allowed Columns")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsAllowedResultRegions, "allowed-result-regions", "", "", "Allowed Result Regions")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsAnalysisMethod, "analysis-method", "", "", "Analysis Method")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsAnalysisParameters, "analysis-parameters", "", "", "Analysis Parameters")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsAnalysisRulePolicy, "analysis-rule-policy", "", "", "Analysis Rule Policy")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsAnalysisRuleType, "analysis-rule-type", "", "", "Analysis Rule Type")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsAnalysisTemplateArn, "analysis-template-arn", "", "", "Analysis Template ARN")
	_cleanroomsCmd.Flags().StringSliceVarP(&_cleanroomsAnalysisTemplateArns, "analysis-template-arns", "", nil, "Analysis Template Arns")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsAnalysisTemplateIdentifier, "analysis-template-identifier", "", "", "Analysis Template Identifier")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsAnalyticsEngine, "analytics-engine", "", "", "Analytics Engine")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsAutoApprovedChangeRequestTypes, "auto-approved-change-request-types", "", "", "Auto Approved Change Request Types")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsAutoRefresh, "auto-refresh", "", "", "Auto Refresh")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsChangeRequestIdentifier, "change-request-identifier", "", "", "Change Request Identifier")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsChanges, "changes", "", "", "Changes")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsCollaborationIdentifier, "collaboration-identifier", "", "", "Collaboration Identifier")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsComputeConfiguration, "compute-configuration", "", "", "Compute Configuration")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsConfiguredAudienceModelArn, "configured-audience-model-arn", "", "", "Configured Audience Model ARN")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsConfiguredAudienceModelAssociationIdentifier, "configured-audience-model-association-identifier", "", "", "Configured Audience Model Association Identifier")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsConfiguredAudienceModelAssociationName, "configured-audience-model-association-name", "", "", "Configured Audience Model Association Name")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsConfiguredTableAssociationIdentifier, "configured-table-association-identifier", "", "", "Configured Table Association Identifier")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsConfiguredTableIdentifier, "configured-table-identifier", "", "", "Configured Table Identifier")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsCreatorDisplayName, "creator-display-name", "", "", "Creator Display Name")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsCreatorMemberAbilities, "creator-member-abilities", "", "", "Creator Member Abilities")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsCreatorMLMemberAbilities, "creator-ml-member-abilities", "", "", "Creator Ml Member Abilities")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsCreatorPaymentConfiguration, "creator-payment-configuration", "", "", "Creator Payment Configuration")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsDataEncryptionMetadata, "data-encryption-metadata", "", "", "Data Encryption Metadata")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsDefaultJobResultConfiguration, "default-job-result-configuration", "", "", "Default Job Result Configuration")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsDefaultResultConfiguration, "default-result-configuration", "", "", "Default Result Configuration")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsDescription, "description", "", "", "Description")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsErrorMessageConfiguration, "error-message-configuration", "", "", "Error Message Configuration")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsFormat, "format", "", "", "Format")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsIdMappingConfig, "id-mapping-config", "", "", "ID Mapping Config")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsIdMappingTableIdentifier, "id-mapping-table-identifier", "", "", "ID Mapping Table Identifier")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsIdNamespaceAssociationIdentifier, "id-namespace-association-identifier", "", "", "ID Namespace Association Identifier")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsInputReferenceConfig, "input-reference-config", "", "", "Input Reference Config")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsIsMetricsEnabled, "is-metrics-enabled", "", "", "Is Metrics Enabled")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsJobLogStatus, "job-log-status", "", "", "Job Log Status")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsJobParameters, "job-parameters", "", "", "Job Parameters")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsJobType, "job-type", "", "", "Job Type")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsKmsKeyArn, "kms-key-arn", "", "", "KMS Key ARN")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsManageResourcePolicies, "manage-resource-policies", "", "", "Manage Resource Policies")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsMaxResults, "max-results", "", "", "Max Results")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsMemberStatus, "member-status", "", "", "Member Status")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsMembers, "members", "", "", "Members")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsMembershipIdentifier, "membership-identifier", "", "", "Membership Identifier")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsName, "name", "", "", "Name")
	_cleanroomsCmd.Flags().StringSliceVarP(&_cleanroomsNames, "names", "", nil, "Names")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsNextToken, "next-token", "", "", "Next Token")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsParameters, "parameters", "", "", "Parameters")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsPaymentConfiguration, "payment-configuration", "", "", "Payment Configuration")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsPrivacyBudgetTemplateIdentifier, "privacy-budget-template-identifier", "", "", "Privacy Budget Template Identifier")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsPrivacyBudgetType, "privacy-budget-type", "", "", "Privacy Budget Type")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsProtectedJobIdentifier, "protected-job-identifier", "", "", "Protected Job Identifier")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsProtectedQueryIdentifier, "protected-query-identifier", "", "", "Protected Query Identifier")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsQueryLogStatus, "query-log-status", "", "", "Query Log Status")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsResourceArn, "resource-arn", "", "", "Resource ARN")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsResultConfiguration, "result-configuration", "", "", "Result Configuration")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsRoleArn, "role-arn", "", "", "Role ARN")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsSchema, "schema", "", "", "Schema")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsSchemaAnalysisRuleRequests, "schema-analysis-rule-requests", "", "", "Schema Analysis Rule Requests")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsSchemaType, "schema-type", "", "", "Schema Type")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsSelectedAnalysisMethods, "selected-analysis-methods", "", "", "Selected Analysis Methods")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsSource, "source", "", "", "Source")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsSqlParameters, "sql-parameters", "", "", "Sql Parameters")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsStatus, "status", "", "", "Status")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsSyntheticDataParameters, "synthetic-data-parameters", "", "", "Synthetic Data Parameters")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsTableReference, "table-reference", "", "", "Table Reference")
	_cleanroomsCmd.Flags().StringSliceVarP(&_cleanroomsTagKeys, "tag-keys", "", nil, "Tag Keys")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsTags, "tags", "", "", "Tags")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsTargetStatus, "target-status", "", "", "Target Status")
	_cleanroomsCmd.Flags().StringVarP(&_cleanroomsType, "type", "", "", "Type")

	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsBatchGetCollaborationAnalysisTemplate, "batch-get-collaboration-analysis-template", "", false, "Batch Get Collaboration Analysis Template")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsBatchGetSchema, "batch-get-schema", "", false, "Batch Get Schema")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsBatchGetSchemaAnalysisRule, "batch-get-schema-analysis-rule", "", false, "Batch Get Schema Analysis Rule")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsCreateAnalysisTemplate, "create-analysis-template", "", false, "Create Analysis Template")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsCreateCollaboration, "create-collaboration", "", false, "Create Collaboration")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsCreateCollaborationChangeRequest, "create-collaboration-change-request", "", false, "Create Collaboration Change Request")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsCreateConfiguredAudienceModelAssociation, "create-configured-audience-model-association", "", false, "Create Configured Audience Model Association")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsCreateConfiguredTable, "create-configured-table", "", false, "Create Configured Table")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsCreateConfiguredTableAnalysisRule, "create-configured-table-analysis-rule", "", false, "Create Configured Table Analysis Rule")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsCreateConfiguredTableAssociation, "create-configured-table-association", "", false, "Create Configured Table Association")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsCreateConfiguredTableAssociationAnalysisRule, "create-configured-table-association-analysis-rule", "", false, "Create Configured Table Association Analysis Rule")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsCreateIdMappingTable, "create-id-mapping-table", "", false, "Create ID Mapping Table")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsCreateIdNamespaceAssociation, "create-id-namespace-association", "", false, "Create ID Namespace Association")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsCreateMembership, "create-membership", "", false, "Create Membership")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsCreatePrivacyBudgetTemplate, "create-privacy-budget-template", "", false, "Create Privacy Budget Template")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsDeleteAnalysisTemplate, "delete-analysis-template", "", false, "Delete Analysis Template")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsDeleteCollaboration, "delete-collaboration", "", false, "Delete Collaboration")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsDeleteConfiguredAudienceModelAssociation, "delete-configured-audience-model-association", "", false, "Delete Configured Audience Model Association")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsDeleteConfiguredTable, "delete-configured-table", "", false, "Delete Configured Table")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsDeleteConfiguredTableAnalysisRule, "delete-configured-table-analysis-rule", "", false, "Delete Configured Table Analysis Rule")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsDeleteConfiguredTableAssociation, "delete-configured-table-association", "", false, "Delete Configured Table Association")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsDeleteConfiguredTableAssociationAnalysisRule, "delete-configured-table-association-analysis-rule", "", false, "Delete Configured Table Association Analysis Rule")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsDeleteIdMappingTable, "delete-id-mapping-table", "", false, "Delete ID Mapping Table")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsDeleteIdNamespaceAssociation, "delete-id-namespace-association", "", false, "Delete ID Namespace Association")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsDeleteMember, "delete-member", "", false, "Delete Member")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsDeleteMembership, "delete-membership", "", false, "Delete Membership")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsDeletePrivacyBudgetTemplate, "delete-privacy-budget-template", "", false, "Delete Privacy Budget Template")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsGetAnalysisTemplate, "get-analysis-template", "", false, "Get Analysis Template")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsGetCollaboration, "get-collaboration", "", false, "Get Collaboration")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsGetCollaborationAnalysisTemplate, "get-collaboration-analysis-template", "", false, "Get Collaboration Analysis Template")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsGetCollaborationChangeRequest, "get-collaboration-change-request", "", false, "Get Collaboration Change Request")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsGetCollaborationConfiguredAudienceModelAssociation, "get-collaboration-configured-audience-model-association", "", false, "Get Collaboration Configured Audience Model Association")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsGetCollaborationIdNamespaceAssociation, "get-collaboration-id-namespace-association", "", false, "Get Collaboration ID Namespace Association")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsGetCollaborationPrivacyBudgetTemplate, "get-collaboration-privacy-budget-template", "", false, "Get Collaboration Privacy Budget Template")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsGetConfiguredAudienceModelAssociation, "get-configured-audience-model-association", "", false, "Get Configured Audience Model Association")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsGetConfiguredTable, "get-configured-table", "", false, "Get Configured Table")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsGetConfiguredTableAnalysisRule, "get-configured-table-analysis-rule", "", false, "Get Configured Table Analysis Rule")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsGetConfiguredTableAssociation, "get-configured-table-association", "", false, "Get Configured Table Association")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsGetConfiguredTableAssociationAnalysisRule, "get-configured-table-association-analysis-rule", "", false, "Get Configured Table Association Analysis Rule")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsGetIdMappingTable, "get-id-mapping-table", "", false, "Get ID Mapping Table")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsGetIdNamespaceAssociation, "get-id-namespace-association", "", false, "Get ID Namespace Association")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsGetMembership, "get-membership", "", false, "Get Membership")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsGetPrivacyBudgetTemplate, "get-privacy-budget-template", "", false, "Get Privacy Budget Template")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsGetProtectedJob, "get-protected-job", "", false, "Get Protected Job")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsGetProtectedQuery, "get-protected-query", "", false, "Get Protected Query")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsGetSchema, "get-schema", "", false, "Get Schema")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsGetSchemaAnalysisRule, "get-schema-analysis-rule", "", false, "Get Schema Analysis Rule")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsListAnalysisTemplates, "list-analysis-templates", "", false, "List Analysis Templates")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsListCollaborationAnalysisTemplates, "list-collaboration-analysis-templates", "", false, "List Collaboration Analysis Templates")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsListCollaborationChangeRequests, "list-collaboration-change-requests", "", false, "List Collaboration Change Requests")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsListCollaborationConfiguredAudienceModelAssociations, "list-collaboration-configured-audience-model-associations", "", false, "List Collaboration Configured Audience Model Associations")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsListCollaborationIdNamespaceAssociations, "list-collaboration-id-namespace-associations", "", false, "List Collaboration ID Namespace Associations")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsListCollaborationPrivacyBudgetTemplates, "list-collaboration-privacy-budget-templates", "", false, "List Collaboration Privacy Budget Templates")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsListCollaborationPrivacyBudgets, "list-collaboration-privacy-budgets", "", false, "List Collaboration Privacy Budgets")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsListCollaborations, "list-collaborations", "", false, "List Collaborations")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsListConfiguredAudienceModelAssociations, "list-configured-audience-model-associations", "", false, "List Configured Audience Model Associations")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsListConfiguredTableAssociations, "list-configured-table-associations", "", false, "List Configured Table Associations")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsListConfiguredTables, "list-configured-tables", "", false, "List Configured Tables")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsListIdMappingTables, "list-id-mapping-tables", "", false, "List ID Mapping Tables")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsListIdNamespaceAssociations, "list-id-namespace-associations", "", false, "List ID Namespace Associations")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsListMembers, "list-members", "", false, "List Members")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsListMemberships, "list-memberships", "", false, "List Memberships")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsListPrivacyBudgetTemplates, "list-privacy-budget-templates", "", false, "List Privacy Budget Templates")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsListPrivacyBudgets, "list-privacy-budgets", "", false, "List Privacy Budgets")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsListProtectedJobs, "list-protected-jobs", "", false, "List Protected Jobs")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsListProtectedQueries, "list-protected-queries", "", false, "List Protected Queries")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsListSchemas, "list-schemas", "", false, "List Schemas")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsPopulateIdMappingTable, "populate-id-mapping-table", "", false, "Populate ID Mapping Table")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsPreviewPrivacyImpact, "preview-privacy-impact", "", false, "Preview Privacy Impact")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsStartProtectedJob, "start-protected-job", "", false, "Start Protected Job")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsStartProtectedQuery, "start-protected-query", "", false, "Start Protected Query")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsTagResource, "tag-resource", "", false, "Tag Resource")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsUntagResource, "untag-resource", "", false, "Untag Resource")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsUpdateAnalysisTemplate, "update-analysis-template", "", false, "Update Analysis Template")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsUpdateCollaboration, "update-collaboration", "", false, "Update Collaboration")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsUpdateCollaborationChangeRequest, "update-collaboration-change-request", "", false, "Update Collaboration Change Request")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsUpdateConfiguredAudienceModelAssociation, "update-configured-audience-model-association", "", false, "Update Configured Audience Model Association")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsUpdateConfiguredTable, "update-configured-table", "", false, "Update Configured Table")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsUpdateConfiguredTableAnalysisRule, "update-configured-table-analysis-rule", "", false, "Update Configured Table Analysis Rule")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsUpdateConfiguredTableAssociation, "update-configured-table-association", "", false, "Update Configured Table Association")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsUpdateConfiguredTableAssociationAnalysisRule, "update-configured-table-association-analysis-rule", "", false, "Update Configured Table Association Analysis Rule")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsUpdateIdMappingTable, "update-id-mapping-table", "", false, "Update ID Mapping Table")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsUpdateIdNamespaceAssociation, "update-id-namespace-association", "", false, "Update ID Namespace Association")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsUpdateMembership, "update-membership", "", false, "Update Membership")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsUpdatePrivacyBudgetTemplate, "update-privacy-budget-template", "", false, "Update Privacy Budget Template")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsUpdateProtectedJob, "update-protected-job", "", false, "Update Protected Job")
	_cleanroomsCmd.Flags().BoolVarP(&_cleanroomsUpdateProtectedQuery, "update-protected-query", "", false, "Update Protected Query")

}
