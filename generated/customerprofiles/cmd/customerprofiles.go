package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/customerprofiles"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// customerprofilesCmd represents the customerprofiles command
var _customerprofilesCmd = &cobra.Command{
	Use:   "customerprofiles",
	Short: "AWS customerprofiles CLI",
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
		client := customerprofiles.NewFromConfig(cfg)
		if _customerprofilesAddProfileKey {
			customerprofiles_AddProfileKey(cfg, client)
			return
		}
		if _customerprofilesBatchGetCalculatedAttributeForProfile {
			customerprofiles_BatchGetCalculatedAttributeForProfile(cfg, client)
			return
		}
		if _customerprofilesBatchGetProfile {
			customerprofiles_BatchGetProfile(cfg, client)
			return
		}
		if _customerprofilesCreateCalculatedAttributeDefinition {
			customerprofiles_CreateCalculatedAttributeDefinition(cfg, client)
			return
		}
		if _customerprofilesCreateDomain {
			customerprofiles_CreateDomain(cfg, client)
			return
		}
		if _customerprofilesCreateDomainLayout {
			customerprofiles_CreateDomainLayout(cfg, client)
			return
		}
		if _customerprofilesCreateEventStream {
			customerprofiles_CreateEventStream(cfg, client)
			return
		}
		if _customerprofilesCreateEventTrigger {
			customerprofiles_CreateEventTrigger(cfg, client)
			return
		}
		if _customerprofilesCreateIntegrationWorkflow {
			customerprofiles_CreateIntegrationWorkflow(cfg, client)
			return
		}
		if _customerprofilesCreateProfile {
			customerprofiles_CreateProfile(cfg, client)
			return
		}
		if _customerprofilesCreateRecommender {
			customerprofiles_CreateRecommender(cfg, client)
			return
		}
		if _customerprofilesCreateSegmentDefinition {
			customerprofiles_CreateSegmentDefinition(cfg, client)
			return
		}
		if _customerprofilesCreateSegmentEstimate {
			customerprofiles_CreateSegmentEstimate(cfg, client)
			return
		}
		if _customerprofilesCreateSegmentSnapshot {
			customerprofiles_CreateSegmentSnapshot(cfg, client)
			return
		}
		if _customerprofilesCreateUploadJob {
			customerprofiles_CreateUploadJob(cfg, client)
			return
		}
		if _customerprofilesDeleteCalculatedAttributeDefinition {
			customerprofiles_DeleteCalculatedAttributeDefinition(cfg, client)
			return
		}
		if _customerprofilesDeleteDomain {
			customerprofiles_DeleteDomain(cfg, client)
			return
		}
		if _customerprofilesDeleteDomainLayout {
			customerprofiles_DeleteDomainLayout(cfg, client)
			return
		}
		if _customerprofilesDeleteDomainObjectType {
			customerprofiles_DeleteDomainObjectType(cfg, client)
			return
		}
		if _customerprofilesDeleteEventStream {
			customerprofiles_DeleteEventStream(cfg, client)
			return
		}
		if _customerprofilesDeleteEventTrigger {
			customerprofiles_DeleteEventTrigger(cfg, client)
			return
		}
		if _customerprofilesDeleteIntegration {
			customerprofiles_DeleteIntegration(cfg, client)
			return
		}
		if _customerprofilesDeleteProfile {
			customerprofiles_DeleteProfile(cfg, client)
			return
		}
		if _customerprofilesDeleteProfileKey {
			customerprofiles_DeleteProfileKey(cfg, client)
			return
		}
		if _customerprofilesDeleteProfileObject {
			customerprofiles_DeleteProfileObject(cfg, client)
			return
		}
		if _customerprofilesDeleteProfileObjectType {
			customerprofiles_DeleteProfileObjectType(cfg, client)
			return
		}
		if _customerprofilesDeleteRecommender {
			customerprofiles_DeleteRecommender(cfg, client)
			return
		}
		if _customerprofilesDeleteSegmentDefinition {
			customerprofiles_DeleteSegmentDefinition(cfg, client)
			return
		}
		if _customerprofilesDeleteWorkflow {
			customerprofiles_DeleteWorkflow(cfg, client)
			return
		}
		if _customerprofilesDetectProfileObjectType {
			customerprofiles_DetectProfileObjectType(cfg, client)
			return
		}
		if _customerprofilesGetAutoMergingPreview {
			customerprofiles_GetAutoMergingPreview(cfg, client)
			return
		}
		if _customerprofilesGetCalculatedAttributeDefinition {
			customerprofiles_GetCalculatedAttributeDefinition(cfg, client)
			return
		}
		if _customerprofilesGetCalculatedAttributeForProfile {
			customerprofiles_GetCalculatedAttributeForProfile(cfg, client)
			return
		}
		if _customerprofilesGetDomain {
			customerprofiles_GetDomain(cfg, client)
			return
		}
		if _customerprofilesGetDomainLayout {
			customerprofiles_GetDomainLayout(cfg, client)
			return
		}
		if _customerprofilesGetDomainObjectType {
			customerprofiles_GetDomainObjectType(cfg, client)
			return
		}
		if _customerprofilesGetEventStream {
			customerprofiles_GetEventStream(cfg, client)
			return
		}
		if _customerprofilesGetEventTrigger {
			customerprofiles_GetEventTrigger(cfg, client)
			return
		}
		if _customerprofilesGetIdentityResolutionJob {
			customerprofiles_GetIdentityResolutionJob(cfg, client)
			return
		}
		if _customerprofilesGetIntegration {
			customerprofiles_GetIntegration(cfg, client)
			return
		}
		if _customerprofilesGetMatches {
			customerprofiles_GetMatches(cfg, client)
			return
		}
		if _customerprofilesGetObjectTypeAttributeStatistics {
			customerprofiles_GetObjectTypeAttributeStatistics(cfg, client)
			return
		}
		if _customerprofilesGetProfileHistoryRecord {
			customerprofiles_GetProfileHistoryRecord(cfg, client)
			return
		}
		if _customerprofilesGetProfileObjectType {
			customerprofiles_GetProfileObjectType(cfg, client)
			return
		}
		if _customerprofilesGetProfileObjectTypeTemplate {
			customerprofiles_GetProfileObjectTypeTemplate(cfg, client)
			return
		}
		if _customerprofilesGetProfileRecommendations {
			customerprofiles_GetProfileRecommendations(cfg, client)
			return
		}
		if _customerprofilesGetRecommender {
			customerprofiles_GetRecommender(cfg, client)
			return
		}
		if _customerprofilesGetSegmentDefinition {
			customerprofiles_GetSegmentDefinition(cfg, client)
			return
		}
		if _customerprofilesGetSegmentEstimate {
			customerprofiles_GetSegmentEstimate(cfg, client)
			return
		}
		if _customerprofilesGetSegmentMembership {
			customerprofiles_GetSegmentMembership(cfg, client)
			return
		}
		if _customerprofilesGetSegmentSnapshot {
			customerprofiles_GetSegmentSnapshot(cfg, client)
			return
		}
		if _customerprofilesGetSimilarProfiles {
			customerprofiles_GetSimilarProfiles(cfg, client)
			return
		}
		if _customerprofilesGetUploadJob {
			customerprofiles_GetUploadJob(cfg, client)
			return
		}
		if _customerprofilesGetUploadJobPath {
			customerprofiles_GetUploadJobPath(cfg, client)
			return
		}
		if _customerprofilesGetWorkflow {
			customerprofiles_GetWorkflow(cfg, client)
			return
		}
		if _customerprofilesGetWorkflowSteps {
			customerprofiles_GetWorkflowSteps(cfg, client)
			return
		}
		if _customerprofilesListAccountIntegrations {
			customerprofiles_ListAccountIntegrations(cfg, client)
			return
		}
		if _customerprofilesListCalculatedAttributeDefinitions {
			customerprofiles_ListCalculatedAttributeDefinitions(cfg, client)
			return
		}
		if _customerprofilesListCalculatedAttributesForProfile {
			customerprofiles_ListCalculatedAttributesForProfile(cfg, client)
			return
		}
		if _customerprofilesListDomainLayouts {
			customerprofiles_ListDomainLayouts(cfg, client)
			return
		}
		if _customerprofilesListDomainObjectTypes {
			customerprofiles_ListDomainObjectTypes(cfg, client)
			return
		}
		if _customerprofilesListDomains {
			customerprofiles_ListDomains(cfg, client)
			return
		}
		if _customerprofilesListEventStreams {
			customerprofiles_ListEventStreams(cfg, client)
			return
		}
		if _customerprofilesListEventTriggers {
			customerprofiles_ListEventTriggers(cfg, client)
			return
		}
		if _customerprofilesListIdentityResolutionJobs {
			customerprofiles_ListIdentityResolutionJobs(cfg, client)
			return
		}
		if _customerprofilesListIntegrations {
			customerprofiles_ListIntegrations(cfg, client)
			return
		}
		if _customerprofilesListObjectTypeAttributeValues {
			customerprofiles_ListObjectTypeAttributeValues(cfg, client)
			return
		}
		if _customerprofilesListObjectTypeAttributes {
			customerprofiles_ListObjectTypeAttributes(cfg, client)
			return
		}
		if _customerprofilesListProfileAttributeValues {
			customerprofiles_ListProfileAttributeValues(cfg, client)
			return
		}
		if _customerprofilesListProfileHistoryRecords {
			customerprofiles_ListProfileHistoryRecords(cfg, client)
			return
		}
		if _customerprofilesListProfileObjectTypeTemplates {
			customerprofiles_ListProfileObjectTypeTemplates(cfg, client)
			return
		}
		if _customerprofilesListProfileObjectTypes {
			customerprofiles_ListProfileObjectTypes(cfg, client)
			return
		}
		if _customerprofilesListProfileObjects {
			customerprofiles_ListProfileObjects(cfg, client)
			return
		}
		if _customerprofilesListRecommenderRecipes {
			customerprofiles_ListRecommenderRecipes(cfg, client)
			return
		}
		if _customerprofilesListRecommenders {
			customerprofiles_ListRecommenders(cfg, client)
			return
		}
		if _customerprofilesListRuleBasedMatches {
			customerprofiles_ListRuleBasedMatches(cfg, client)
			return
		}
		if _customerprofilesListSegmentDefinitions {
			customerprofiles_ListSegmentDefinitions(cfg, client)
			return
		}
		if _customerprofilesListTagsForResource {
			customerprofiles_ListTagsForResource(cfg, client)
			return
		}
		if _customerprofilesListUploadJobs {
			customerprofiles_ListUploadJobs(cfg, client)
			return
		}
		if _customerprofilesListWorkflows {
			customerprofiles_ListWorkflows(cfg, client)
			return
		}
		if _customerprofilesMergeProfiles {
			customerprofiles_MergeProfiles(cfg, client)
			return
		}
		if _customerprofilesPutDomainObjectType {
			customerprofiles_PutDomainObjectType(cfg, client)
			return
		}
		if _customerprofilesPutIntegration {
			customerprofiles_PutIntegration(cfg, client)
			return
		}
		if _customerprofilesPutProfileObject {
			customerprofiles_PutProfileObject(cfg, client)
			return
		}
		if _customerprofilesPutProfileObjectType {
			customerprofiles_PutProfileObjectType(cfg, client)
			return
		}
		if _customerprofilesSearchProfiles {
			customerprofiles_SearchProfiles(cfg, client)
			return
		}
		if _customerprofilesStartRecommender {
			customerprofiles_StartRecommender(cfg, client)
			return
		}
		if _customerprofilesStartUploadJob {
			customerprofiles_StartUploadJob(cfg, client)
			return
		}
		if _customerprofilesStopRecommender {
			customerprofiles_StopRecommender(cfg, client)
			return
		}
		if _customerprofilesStopUploadJob {
			customerprofiles_StopUploadJob(cfg, client)
			return
		}
		if _customerprofilesTagResource {
			customerprofiles_TagResource(cfg, client)
			return
		}
		if _customerprofilesUntagResource {
			customerprofiles_UntagResource(cfg, client)
			return
		}
		if _customerprofilesUpdateCalculatedAttributeDefinition {
			customerprofiles_UpdateCalculatedAttributeDefinition(cfg, client)
			return
		}
		if _customerprofilesUpdateDomain {
			customerprofiles_UpdateDomain(cfg, client)
			return
		}
		if _customerprofilesUpdateDomainLayout {
			customerprofiles_UpdateDomainLayout(cfg, client)
			return
		}
		if _customerprofilesUpdateEventTrigger {
			customerprofiles_UpdateEventTrigger(cfg, client)
			return
		}
		if _customerprofilesUpdateProfile {
			customerprofiles_UpdateProfile(cfg, client)
			return
		}
		if _customerprofilesUpdateRecommender {
			customerprofiles_UpdateRecommender(cfg, client)
			return
		}

	},
}

var (
	_customerprofilesAddProfileKey                         bool
	_customerprofilesBatchGetCalculatedAttributeForProfile bool
	_customerprofilesBatchGetProfile                       bool
	_customerprofilesCreateCalculatedAttributeDefinition   bool
	_customerprofilesCreateDomain                          bool
	_customerprofilesCreateDomainLayout                    bool
	_customerprofilesCreateEventStream                     bool
	_customerprofilesCreateEventTrigger                    bool
	_customerprofilesCreateIntegrationWorkflow             bool
	_customerprofilesCreateProfile                         bool
	_customerprofilesCreateRecommender                     bool
	_customerprofilesCreateSegmentDefinition               bool
	_customerprofilesCreateSegmentEstimate                 bool
	_customerprofilesCreateSegmentSnapshot                 bool
	_customerprofilesCreateUploadJob                       bool
	_customerprofilesDeleteCalculatedAttributeDefinition   bool
	_customerprofilesDeleteDomain                          bool
	_customerprofilesDeleteDomainLayout                    bool
	_customerprofilesDeleteDomainObjectType                bool
	_customerprofilesDeleteEventStream                     bool
	_customerprofilesDeleteEventTrigger                    bool
	_customerprofilesDeleteIntegration                     bool
	_customerprofilesDeleteProfile                         bool
	_customerprofilesDeleteProfileKey                      bool
	_customerprofilesDeleteProfileObject                   bool
	_customerprofilesDeleteProfileObjectType               bool
	_customerprofilesDeleteRecommender                     bool
	_customerprofilesDeleteSegmentDefinition               bool
	_customerprofilesDeleteWorkflow                        bool
	_customerprofilesDetectProfileObjectType               bool
	_customerprofilesGetAutoMergingPreview                 bool
	_customerprofilesGetCalculatedAttributeDefinition      bool
	_customerprofilesGetCalculatedAttributeForProfile      bool
	_customerprofilesGetDomain                             bool
	_customerprofilesGetDomainLayout                       bool
	_customerprofilesGetDomainObjectType                   bool
	_customerprofilesGetEventStream                        bool
	_customerprofilesGetEventTrigger                       bool
	_customerprofilesGetIdentityResolutionJob              bool
	_customerprofilesGetIntegration                        bool
	_customerprofilesGetMatches                            bool
	_customerprofilesGetObjectTypeAttributeStatistics      bool
	_customerprofilesGetProfileHistoryRecord               bool
	_customerprofilesGetProfileObjectType                  bool
	_customerprofilesGetProfileObjectTypeTemplate          bool
	_customerprofilesGetProfileRecommendations             bool
	_customerprofilesGetRecommender                        bool
	_customerprofilesGetSegmentDefinition                  bool
	_customerprofilesGetSegmentEstimate                    bool
	_customerprofilesGetSegmentMembership                  bool
	_customerprofilesGetSegmentSnapshot                    bool
	_customerprofilesGetSimilarProfiles                    bool
	_customerprofilesGetUploadJob                          bool
	_customerprofilesGetUploadJobPath                      bool
	_customerprofilesGetWorkflow                           bool
	_customerprofilesGetWorkflowSteps                      bool
	_customerprofilesListAccountIntegrations               bool
	_customerprofilesListCalculatedAttributeDefinitions    bool
	_customerprofilesListCalculatedAttributesForProfile    bool
	_customerprofilesListDomainLayouts                     bool
	_customerprofilesListDomainObjectTypes                 bool
	_customerprofilesListDomains                           bool
	_customerprofilesListEventStreams                      bool
	_customerprofilesListEventTriggers                     bool
	_customerprofilesListIdentityResolutionJobs            bool
	_customerprofilesListIntegrations                      bool
	_customerprofilesListObjectTypeAttributeValues         bool
	_customerprofilesListObjectTypeAttributes              bool
	_customerprofilesListProfileAttributeValues            bool
	_customerprofilesListProfileHistoryRecords             bool
	_customerprofilesListProfileObjectTypeTemplates        bool
	_customerprofilesListProfileObjectTypes                bool
	_customerprofilesListProfileObjects                    bool
	_customerprofilesListRecommenderRecipes                bool
	_customerprofilesListRecommenders                      bool
	_customerprofilesListRuleBasedMatches                  bool
	_customerprofilesListSegmentDefinitions                bool
	_customerprofilesListTagsForResource                   bool
	_customerprofilesListUploadJobs                        bool
	_customerprofilesListWorkflows                         bool
	_customerprofilesMergeProfiles                         bool
	_customerprofilesPutDomainObjectType                   bool
	_customerprofilesPutIntegration                        bool
	_customerprofilesPutProfileObject                      bool
	_customerprofilesPutProfileObjectType                  bool
	_customerprofilesSearchProfiles                        bool
	_customerprofilesStartRecommender                      bool
	_customerprofilesStartUploadJob                        bool
	_customerprofilesStopRecommender                       bool
	_customerprofilesStopUploadJob                         bool
	_customerprofilesTagResource                           bool
	_customerprofilesUntagResource                         bool
	_customerprofilesUpdateCalculatedAttributeDefinition   bool
	_customerprofilesUpdateDomain                          bool
	_customerprofilesUpdateDomainLayout                    bool
	_customerprofilesUpdateEventTrigger                    bool
	_customerprofilesUpdateProfile                         bool
	_customerprofilesUpdateRecommender                     bool

	_customerprofilesAccountNumber                       string
	_customerprofilesActionType                          string
	_customerprofilesAdditionalInformation               string
	_customerprofilesAdditionalSearchKeys                string
	_customerprofilesAddress                             string
	_customerprofilesAllowProfileCreation                string
	_customerprofilesAttributeDetails                    string
	_customerprofilesAttributeName                       string
	_customerprofilesAttributes                          string
	_customerprofilesBillingAddress                      string
	_customerprofilesBirthDate                           string
	_customerprofilesBusinessEmailAddress                string
	_customerprofilesBusinessName                        string
	_customerprofilesBusinessPhoneNumber                 string
	_customerprofilesCalculatedAttributeName             string
	_customerprofilesConditionOverrides                  string
	_customerprofilesConditions                          string
	_customerprofilesConflictResolution                  string
	_customerprofilesConsolidation                       string
	_customerprofilesContext                             string
	_customerprofilesDataExpiry                          string
	_customerprofilesDataFormat                          string
	_customerprofilesDataStore                           string
	_customerprofilesDeadLetterQueueUrl                  string
	_customerprofilesDefaultEncryptionKey                string
	_customerprofilesDefaultExpirationDays               string
	_customerprofilesDescription                         string
	_customerprofilesDestinationUri                      string
	_customerprofilesDisplayName                         string
	_customerprofilesDomainName                          string
	_customerprofilesEmailAddress                        string
	_customerprofilesEncryptionKey                       string
	_customerprofilesEngagementPreferences               string
	_customerprofilesEstimateId                          string
	_customerprofilesEventStreamName                     string
	_customerprofilesEventTriggerConditions              string
	_customerprofilesEventTriggerLimits                  string
	_customerprofilesEventTriggerName                    string
	_customerprofilesEventTriggerNames                   []string
	_customerprofilesExpirationDays                      string
	_customerprofilesFieldSourceProfileIds               string
	_customerprofilesFields                              string
	_customerprofilesFilter                              string
	_customerprofilesFirstName                           string
	_customerprofilesFlowDefinition                      string
	_customerprofilesGender                              string
	_customerprofilesGenderString                        string
	_customerprofilesHomePhoneNumber                     string
	_customerprofilesId                                  string
	_customerprofilesIncludeHidden                       string
	_customerprofilesIntegrationConfig                   string
	_customerprofilesIsDefault                           string
	_customerprofilesJobId                               string
	_customerprofilesKeyName                             string
	_customerprofilesKeys                                string
	_customerprofilesLastName                            string
	_customerprofilesLayout                              string
	_customerprofilesLayoutDefinitionName                string
	_customerprofilesLayoutType                          string
	_customerprofilesLogicalOperator                     string
	_customerprofilesMailingAddress                      string
	_customerprofilesMainProfileId                       string
	_customerprofilesMatchType                           string
	_customerprofilesMatching                            string
	_customerprofilesMaxProfileObjectCount               string
	_customerprofilesMaxResults                          string
	_customerprofilesMiddleName                          string
	_customerprofilesMinAllowedConfidenceScoreForMerging string
	_customerprofilesMobilePhoneNumber                   string
	_customerprofilesNextToken                           string
	_customerprofilesObject                              string
	_customerprofilesObjectFilter                        string
	_customerprofilesObjectTypeName                      string
	_customerprofilesObjectTypeNames                     string
	_customerprofilesObjects                             []string
	_customerprofilesPartyType                           string
	_customerprofilesPartyTypeString                     string
	_customerprofilesPerformedBy                         string
	_customerprofilesPersonalEmailAddress                string
	_customerprofilesPhoneNumber                         string
	_customerprofilesProfileId                           string
	_customerprofilesProfileIds                          []string
	_customerprofilesProfileIdsToBeMerged                []string
	_customerprofilesProfileObjectUniqueKey              string
	_customerprofilesProfileType                         string
	_customerprofilesQueryEndDate                        string
	_customerprofilesQueryStartDate                      string
	_customerprofilesRecommenderConfig                   string
	_customerprofilesRecommenderName                     string
	_customerprofilesRecommenderRecipeName               string
	_customerprofilesResourceArn                         string
	_customerprofilesRoleArn                             string
	_customerprofilesRuleBasedMatching                   string
	_customerprofilesScope                               string
	_customerprofilesSearchKey                           string
	_customerprofilesSearchValue                         string
	_customerprofilesSegmentDefinitionName               string
	_customerprofilesSegmentFilter                       string
	_customerprofilesSegmentGroups                       string
	_customerprofilesSegmentQuery                        string
	_customerprofilesSegmentSqlQuery                     string
	_customerprofilesShippingAddress                     string
	_customerprofilesSnapshotId                          string
	_customerprofilesSourceLastUpdatedTimestampFormat    string
	_customerprofilesSourcePriority                      string
	_customerprofilesStatistic                           string
	_customerprofilesStatus                              string
	_customerprofilesTagKeys                             []string
	_customerprofilesTags                                string
	_customerprofilesTemplateId                          string
	_customerprofilesTrainingMetricsCount                string
	_customerprofilesUniqueKey                           string
	_customerprofilesUri                                 string
	_customerprofilesUseHistoricalData                   string
	_customerprofilesValues                              []string
	_customerprofilesWorkflowId                          string
	_customerprofilesWorkflowType                        string
)

// Associates a new key value with a specific profile, such as a Contact Record
// ContactId.
//
// A profile object can have a single unique key and any number of additional keys
// that can be used to identify the profile that it belongs to.
func customerprofiles_AddProfileKey(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.AddProfileKeyInput{
		// DomainName: *string, // Required
		// KeyName: *string, // Required
		// ProfileId: *string, // Required
		// Values: []string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesKeyName) > 0 {
		input.KeyName = aws.String(_customerprofilesKeyName)
	}
	if len(_customerprofilesProfileId) > 0 {
		input.ProfileId = aws.String(_customerprofilesProfileId)
	}
	if len(_customerprofilesValues) > 0 {
		input.Values = append([]string(nil), _customerprofilesValues...)
	}

	if resp, err := client.AddProfileKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Fetch the possible attribute values given the attribute name.
func customerprofiles_BatchGetCalculatedAttributeForProfile(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.BatchGetCalculatedAttributeForProfileInput{
		// CalculatedAttributeName: *string, // Required
		// DomainName: *string, // Required
		// ProfileIds: []string, // Required
	}

	if len(_customerprofilesCalculatedAttributeName) > 0 {
		input.CalculatedAttributeName = aws.String(_customerprofilesCalculatedAttributeName)
	}
	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesProfileIds) > 0 {
		input.ProfileIds = append([]string(nil), _customerprofilesProfileIds...)
	}
	if len(_customerprofilesConditionOverrides) > 0 {
		if err := assignInputField(input, "ConditionOverrides", _customerprofilesConditionOverrides); err != nil {
			log.Errorf("invalid --condition-overrides: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchGetCalculatedAttributeForProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get a batch of profiles.
func customerprofiles_BatchGetProfile(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.BatchGetProfileInput{
		// DomainName: *string, // Required
		// ProfileIds: []string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesProfileIds) > 0 {
		input.ProfileIds = append([]string(nil), _customerprofilesProfileIds...)
	}

	if resp, err := client.BatchGetProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new calculated attribute definition. After creation, new object data
// ingested into Customer Profiles will be included in the calculated attribute,
// which can be retrieved for a profile using the [GetCalculatedAttributeForProfile]API. Defining a calculated
// attribute makes it available for all profiles within a domain. Each calculated
// attribute can only reference one ObjectType and at most, two fields from that
// ObjectType .
//
// [GetCalculatedAttributeForProfile]: https://docs.aws.amazon.com/customerprofiles/latest/APIReference/API_GetCalculatedAttributeForProfile.html
func customerprofiles_CreateCalculatedAttributeDefinition(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.CreateCalculatedAttributeDefinitionInput{
		// AttributeDetails: *types.AttributeDetails, // Required
		// CalculatedAttributeName: *string, // Required
		// DomainName: *string, // Required
		// Statistic: types.Statistic, // Required
	}

	if len(_customerprofilesAttributeDetails) > 0 {
		if err := assignInputField(input, "AttributeDetails", _customerprofilesAttributeDetails); err != nil {
			log.Errorf("invalid --attribute-details: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesCalculatedAttributeName) > 0 {
		input.CalculatedAttributeName = aws.String(_customerprofilesCalculatedAttributeName)
	}
	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesStatistic) > 0 {
		if err := assignInputField(input, "Statistic", _customerprofilesStatistic); err != nil {
			log.Errorf("invalid --statistic: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesConditions) > 0 {
		if err := assignInputField(input, "Conditions", _customerprofilesConditions); err != nil {
			log.Errorf("invalid --conditions: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesDescription) > 0 {
		input.Description = aws.String(_customerprofilesDescription)
	}
	if len(_customerprofilesDisplayName) > 0 {
		input.DisplayName = aws.String(_customerprofilesDisplayName)
	}
	if len(_customerprofilesFilter) > 0 {
		if err := assignInputField(input, "Filter", _customerprofilesFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesTags) > 0 {
		if err := assignInputField(input, "Tags", _customerprofilesTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesUseHistoricalData) > 0 {
		if err := assignInputField(input, "UseHistoricalData", _customerprofilesUseHistoricalData); err != nil {
			log.Errorf("invalid --use-historical-data: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCalculatedAttributeDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a domain, which is a container for all customer data, such as customer
// profile attributes, object types, profile keys, and encryption keys. You can
// create multiple domains, and each domain can have multiple third-party
// integrations.
//
// Each Amazon Connect instance can be associated with only one domain. Multiple
// Amazon Connect instances can be associated with one domain.
//
// Use this API or [UpdateDomain] to enable [identity resolution]: set Matching to true.
//
// To prevent cross-service impersonation when you call this API, see [Cross-service confused deputy prevention] for sample
// policies that you should apply.
//
// It is not possible to associate a Customer Profiles domain with an Amazon
// Connect Instance directly from the API. If you would like to create a domain and
// associate a Customer Profiles domain, use the Amazon Connect admin website. For
// more information, see [Enable Customer Profiles].
//
// Each Amazon Connect instance can be associated with only one domain. Multiple
// Amazon Connect instances can be associated with one domain.
//
// [UpdateDomain]: https://docs.aws.amazon.com/customerprofiles/latest/APIReference/API_UpdateDomain.html
// [Enable Customer Profiles]: https://docs.aws.amazon.com/connect/latest/adminguide/enable-customer-profiles.html#enable-customer-profiles-step1
// [Cross-service confused deputy prevention]: https://docs.aws.amazon.com/connect/latest/adminguide/cross-service-confused-deputy-prevention.html
// [identity resolution]: https://docs.aws.amazon.com/customerprofiles/latest/APIReference/API_GetMatches.html
func customerprofiles_CreateDomain(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.CreateDomainInput{
		// DefaultExpirationDays: *int32, // Required
		// DomainName: *string, // Required
	}

	if len(_customerprofilesDefaultExpirationDays) > 0 {
		if err := assignInputField(input, "DefaultExpirationDays", _customerprofilesDefaultExpirationDays); err != nil {
			log.Errorf("invalid --default-expiration-days: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesDataStore) > 0 {
		if err := assignInputField(input, "DataStore", _customerprofilesDataStore); err != nil {
			log.Errorf("invalid --data-store: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesDeadLetterQueueUrl) > 0 {
		input.DeadLetterQueueUrl = aws.String(_customerprofilesDeadLetterQueueUrl)
	}
	if len(_customerprofilesDefaultEncryptionKey) > 0 {
		input.DefaultEncryptionKey = aws.String(_customerprofilesDefaultEncryptionKey)
	}
	if len(_customerprofilesMatching) > 0 {
		if err := assignInputField(input, "Matching", _customerprofilesMatching); err != nil {
			log.Errorf("invalid --matching: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesRuleBasedMatching) > 0 {
		if err := assignInputField(input, "RuleBasedMatching", _customerprofilesRuleBasedMatching); err != nil {
			log.Errorf("invalid --rule-based-matching: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesTags) > 0 {
		if err := assignInputField(input, "Tags", _customerprofilesTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates the layout to view data for a specific domain. This API can only be
// invoked from the Amazon Connect admin website.
func customerprofiles_CreateDomainLayout(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.CreateDomainLayoutInput{
		// Description: *string, // Required
		// DisplayName: *string, // Required
		// DomainName: *string, // Required
		// Layout: *string, // Required
		// LayoutDefinitionName: *string, // Required
		// LayoutType: types.LayoutType, // Required
	}

	if len(_customerprofilesDescription) > 0 {
		input.Description = aws.String(_customerprofilesDescription)
	}
	if len(_customerprofilesDisplayName) > 0 {
		input.DisplayName = aws.String(_customerprofilesDisplayName)
	}
	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesLayout) > 0 {
		input.Layout = aws.String(_customerprofilesLayout)
	}
	if len(_customerprofilesLayoutDefinitionName) > 0 {
		input.LayoutDefinitionName = aws.String(_customerprofilesLayoutDefinitionName)
	}
	if len(_customerprofilesLayoutType) > 0 {
		if err := assignInputField(input, "LayoutType", _customerprofilesLayoutType); err != nil {
			log.Errorf("invalid --layout-type: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesIsDefault) > 0 {
		if err := assignInputField(input, "IsDefault", _customerprofilesIsDefault); err != nil {
			log.Errorf("invalid --is-default: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesTags) > 0 {
		if err := assignInputField(input, "Tags", _customerprofilesTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDomainLayout(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an event stream, which is a subscription to real-time events, such as
// when profiles are created and updated through Amazon Connect Customer Profiles.
//
// Each event stream can be associated with only one Kinesis Data Stream
// destination in the same region and Amazon Web Services account as the customer
// profiles domain
func customerprofiles_CreateEventStream(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.CreateEventStreamInput{
		// DomainName: *string, // Required
		// EventStreamName: *string, // Required
		// Uri: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesEventStreamName) > 0 {
		input.EventStreamName = aws.String(_customerprofilesEventStreamName)
	}
	if len(_customerprofilesUri) > 0 {
		input.Uri = aws.String(_customerprofilesUri)
	}
	if len(_customerprofilesTags) > 0 {
		if err := assignInputField(input, "Tags", _customerprofilesTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateEventStream(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an event trigger, which specifies the rules when to perform action
// based on customer's ingested data.
//
// Each event stream can be associated with only one integration in the same
// region and AWS account as the event stream.
func customerprofiles_CreateEventTrigger(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.CreateEventTriggerInput{
		// DomainName: *string, // Required
		// EventTriggerConditions: []types.EventTriggerCondition, // Required
		// EventTriggerName: *string, // Required
		// ObjectTypeName: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesEventTriggerConditions) > 0 {
		if err := assignInputField(input, "EventTriggerConditions", _customerprofilesEventTriggerConditions); err != nil {
			log.Errorf("invalid --event-trigger-conditions: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesEventTriggerName) > 0 {
		input.EventTriggerName = aws.String(_customerprofilesEventTriggerName)
	}
	if len(_customerprofilesObjectTypeName) > 0 {
		input.ObjectTypeName = aws.String(_customerprofilesObjectTypeName)
	}
	if len(_customerprofilesDescription) > 0 {
		input.Description = aws.String(_customerprofilesDescription)
	}
	if len(_customerprofilesEventTriggerLimits) > 0 {
		if err := assignInputField(input, "EventTriggerLimits", _customerprofilesEventTriggerLimits); err != nil {
			log.Errorf("invalid --event-trigger-limits: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesSegmentFilter) > 0 {
		input.SegmentFilter = aws.String(_customerprofilesSegmentFilter)
	}
	if len(_customerprofilesTags) > 0 {
		if err := assignInputField(input, "Tags", _customerprofilesTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateEventTrigger(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an integration workflow. An integration workflow is an async process
// which ingests historic data and sets up an integration for ongoing updates. The
// supported Amazon AppFlow sources are Salesforce, ServiceNow, and Marketo.
func customerprofiles_CreateIntegrationWorkflow(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.CreateIntegrationWorkflowInput{
		// DomainName: *string, // Required
		// IntegrationConfig: *types.IntegrationConfig, // Required
		// ObjectTypeName: *string, // Required
		// RoleArn: *string, // Required
		// WorkflowType: types.WorkflowType, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesIntegrationConfig) > 0 {
		if err := assignInputField(input, "IntegrationConfig", _customerprofilesIntegrationConfig); err != nil {
			log.Errorf("invalid --integration-config: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesObjectTypeName) > 0 {
		input.ObjectTypeName = aws.String(_customerprofilesObjectTypeName)
	}
	if len(_customerprofilesRoleArn) > 0 {
		input.RoleArn = aws.String(_customerprofilesRoleArn)
	}
	if len(_customerprofilesWorkflowType) > 0 {
		if err := assignInputField(input, "WorkflowType", _customerprofilesWorkflowType); err != nil {
			log.Errorf("invalid --workflow-type: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesTags) > 0 {
		if err := assignInputField(input, "Tags", _customerprofilesTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateIntegrationWorkflow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a standard profile.
// A standard profile represents the following attributes for a customer profile
// in a domain.
func customerprofiles_CreateProfile(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.CreateProfileInput{
		// DomainName: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesAccountNumber) > 0 {
		input.AccountNumber = aws.String(_customerprofilesAccountNumber)
	}
	if len(_customerprofilesAdditionalInformation) > 0 {
		input.AdditionalInformation = aws.String(_customerprofilesAdditionalInformation)
	}
	if len(_customerprofilesAddress) > 0 {
		if err := assignInputField(input, "Address", _customerprofilesAddress); err != nil {
			log.Errorf("invalid --address: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesAttributes) > 0 {
		if err := assignInputField(input, "Attributes", _customerprofilesAttributes); err != nil {
			log.Errorf("invalid --attributes: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesBillingAddress) > 0 {
		if err := assignInputField(input, "BillingAddress", _customerprofilesBillingAddress); err != nil {
			log.Errorf("invalid --billing-address: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesBirthDate) > 0 {
		input.BirthDate = aws.String(_customerprofilesBirthDate)
	}
	if len(_customerprofilesBusinessEmailAddress) > 0 {
		input.BusinessEmailAddress = aws.String(_customerprofilesBusinessEmailAddress)
	}
	if len(_customerprofilesBusinessName) > 0 {
		input.BusinessName = aws.String(_customerprofilesBusinessName)
	}
	if len(_customerprofilesBusinessPhoneNumber) > 0 {
		input.BusinessPhoneNumber = aws.String(_customerprofilesBusinessPhoneNumber)
	}
	if len(_customerprofilesEmailAddress) > 0 {
		input.EmailAddress = aws.String(_customerprofilesEmailAddress)
	}
	if len(_customerprofilesEngagementPreferences) > 0 {
		if err := assignInputField(input, "EngagementPreferences", _customerprofilesEngagementPreferences); err != nil {
			log.Errorf("invalid --engagement-preferences: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesFirstName) > 0 {
		input.FirstName = aws.String(_customerprofilesFirstName)
	}
	if len(_customerprofilesGender) > 0 {
		if err := assignInputField(input, "Gender", _customerprofilesGender); err != nil {
			log.Errorf("invalid --gender: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesGenderString) > 0 {
		input.GenderString = aws.String(_customerprofilesGenderString)
	}
	if len(_customerprofilesHomePhoneNumber) > 0 {
		input.HomePhoneNumber = aws.String(_customerprofilesHomePhoneNumber)
	}
	if len(_customerprofilesLastName) > 0 {
		input.LastName = aws.String(_customerprofilesLastName)
	}
	if len(_customerprofilesMailingAddress) > 0 {
		if err := assignInputField(input, "MailingAddress", _customerprofilesMailingAddress); err != nil {
			log.Errorf("invalid --mailing-address: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesMiddleName) > 0 {
		input.MiddleName = aws.String(_customerprofilesMiddleName)
	}
	if len(_customerprofilesMobilePhoneNumber) > 0 {
		input.MobilePhoneNumber = aws.String(_customerprofilesMobilePhoneNumber)
	}
	if len(_customerprofilesPartyType) > 0 {
		if err := assignInputField(input, "PartyType", _customerprofilesPartyType); err != nil {
			log.Errorf("invalid --party-type: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesPartyTypeString) > 0 {
		input.PartyTypeString = aws.String(_customerprofilesPartyTypeString)
	}
	if len(_customerprofilesPersonalEmailAddress) > 0 {
		input.PersonalEmailAddress = aws.String(_customerprofilesPersonalEmailAddress)
	}
	if len(_customerprofilesPhoneNumber) > 0 {
		input.PhoneNumber = aws.String(_customerprofilesPhoneNumber)
	}
	if len(_customerprofilesProfileType) > 0 {
		if err := assignInputField(input, "ProfileType", _customerprofilesProfileType); err != nil {
			log.Errorf("invalid --profile-type: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesShippingAddress) > 0 {
		if err := assignInputField(input, "ShippingAddress", _customerprofilesShippingAddress); err != nil {
			log.Errorf("invalid --shipping-address: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a recommender
func customerprofiles_CreateRecommender(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.CreateRecommenderInput{
		// DomainName: *string, // Required
		// RecommenderName: *string, // Required
		// RecommenderRecipeName: types.RecommenderRecipeName, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesRecommenderName) > 0 {
		input.RecommenderName = aws.String(_customerprofilesRecommenderName)
	}
	if len(_customerprofilesRecommenderRecipeName) > 0 {
		if err := assignInputField(input, "RecommenderRecipeName", _customerprofilesRecommenderRecipeName); err != nil {
			log.Errorf("invalid --recommender-recipe-name: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesDescription) > 0 {
		input.Description = aws.String(_customerprofilesDescription)
	}
	if len(_customerprofilesRecommenderConfig) > 0 {
		if err := assignInputField(input, "RecommenderConfig", _customerprofilesRecommenderConfig); err != nil {
			log.Errorf("invalid --recommender-config: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesTags) > 0 {
		if err := assignInputField(input, "Tags", _customerprofilesTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateRecommender(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a segment definition associated to the given domain.
func customerprofiles_CreateSegmentDefinition(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.CreateSegmentDefinitionInput{
		// DisplayName: *string, // Required
		// DomainName: *string, // Required
		// SegmentDefinitionName: *string, // Required
	}

	if len(_customerprofilesDisplayName) > 0 {
		input.DisplayName = aws.String(_customerprofilesDisplayName)
	}
	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesSegmentDefinitionName) > 0 {
		input.SegmentDefinitionName = aws.String(_customerprofilesSegmentDefinitionName)
	}
	if len(_customerprofilesDescription) > 0 {
		input.Description = aws.String(_customerprofilesDescription)
	}
	if len(_customerprofilesSegmentGroups) > 0 {
		if err := assignInputField(input, "SegmentGroups", _customerprofilesSegmentGroups); err != nil {
			log.Errorf("invalid --segment-groups: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesSegmentSqlQuery) > 0 {
		input.SegmentSqlQuery = aws.String(_customerprofilesSegmentSqlQuery)
	}
	if len(_customerprofilesTags) > 0 {
		if err := assignInputField(input, "Tags", _customerprofilesTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSegmentDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a segment estimate query.
func customerprofiles_CreateSegmentEstimate(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.CreateSegmentEstimateInput{
		// DomainName: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesSegmentQuery) > 0 {
		if err := assignInputField(input, "SegmentQuery", _customerprofilesSegmentQuery); err != nil {
			log.Errorf("invalid --segment-query: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesSegmentSqlQuery) > 0 {
		input.SegmentSqlQuery = aws.String(_customerprofilesSegmentSqlQuery)
	}

	if resp, err := client.CreateSegmentEstimate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Triggers a job to export a segment to a specified destination.
func customerprofiles_CreateSegmentSnapshot(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.CreateSegmentSnapshotInput{
		// DataFormat: types.DataFormat, // Required
		// DomainName: *string, // Required
		// SegmentDefinitionName: *string, // Required
	}

	if len(_customerprofilesDataFormat) > 0 {
		if err := assignInputField(input, "DataFormat", _customerprofilesDataFormat); err != nil {
			log.Errorf("invalid --data-format: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesSegmentDefinitionName) > 0 {
		input.SegmentDefinitionName = aws.String(_customerprofilesSegmentDefinitionName)
	}
	if len(_customerprofilesDestinationUri) > 0 {
		input.DestinationUri = aws.String(_customerprofilesDestinationUri)
	}
	if len(_customerprofilesEncryptionKey) > 0 {
		input.EncryptionKey = aws.String(_customerprofilesEncryptionKey)
	}
	if len(_customerprofilesRoleArn) > 0 {
		input.RoleArn = aws.String(_customerprofilesRoleArn)
	}

	if resp, err := client.CreateSegmentSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Upload job to ingest data for segment imports. The metadata is
// created for the job with the provided field mapping and unique key.
func customerprofiles_CreateUploadJob(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.CreateUploadJobInput{
		// DisplayName: *string, // Required
		// DomainName: *string, // Required
		// Fields: map[string]types.ObjectTypeField, // Required
		// UniqueKey: *string, // Required
	}

	if len(_customerprofilesDisplayName) > 0 {
		input.DisplayName = aws.String(_customerprofilesDisplayName)
	}
	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesFields) > 0 {
		if err := assignInputField(input, "Fields", _customerprofilesFields); err != nil {
			log.Errorf("invalid --fields: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesUniqueKey) > 0 {
		input.UniqueKey = aws.String(_customerprofilesUniqueKey)
	}
	if len(_customerprofilesDataExpiry) > 0 {
		if err := assignInputField(input, "DataExpiry", _customerprofilesDataExpiry); err != nil {
			log.Errorf("invalid --data-expiry: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateUploadJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing calculated attribute definition. Note that deleting a
// default calculated attribute is possible, however once deleted, you will be
// unable to undo that action and will need to recreate it on your own using the
// CreateCalculatedAttributeDefinition API if you want it back.
func customerprofiles_DeleteCalculatedAttributeDefinition(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.DeleteCalculatedAttributeDefinitionInput{
		// CalculatedAttributeName: *string, // Required
		// DomainName: *string, // Required
	}

	if len(_customerprofilesCalculatedAttributeName) > 0 {
		input.CalculatedAttributeName = aws.String(_customerprofilesCalculatedAttributeName)
	}
	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}

	if resp, err := client.DeleteCalculatedAttributeDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a specific domain and all of its customer data, such as customer
// profile attributes and their related objects.
func customerprofiles_DeleteDomain(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.DeleteDomainInput{
		// DomainName: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}

	if resp, err := client.DeleteDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the layout used to view data for a specific domain. This API can only
// be invoked from the Amazon Connect admin website.
func customerprofiles_DeleteDomainLayout(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.DeleteDomainLayoutInput{
		// DomainName: *string, // Required
		// LayoutDefinitionName: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesLayoutDefinitionName) > 0 {
		input.LayoutDefinitionName = aws.String(_customerprofilesLayoutDefinitionName)
	}

	if resp, err := client.DeleteDomainLayout(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete a DomainObjectType for the given Domain and ObjectType name.
func customerprofiles_DeleteDomainObjectType(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.DeleteDomainObjectTypeInput{
		// DomainName: *string, // Required
		// ObjectTypeName: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesObjectTypeName) > 0 {
		input.ObjectTypeName = aws.String(_customerprofilesObjectTypeName)
	}

	if resp, err := client.DeleteDomainObjectType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disables and deletes the specified event stream.
func customerprofiles_DeleteEventStream(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.DeleteEventStreamInput{
		// DomainName: *string, // Required
		// EventStreamName: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesEventStreamName) > 0 {
		input.EventStreamName = aws.String(_customerprofilesEventStreamName)
	}

	if resp, err := client.DeleteEventStream(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disable and deletes the Event Trigger.
// You cannot delete an Event Trigger with an active Integration associated.
func customerprofiles_DeleteEventTrigger(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.DeleteEventTriggerInput{
		// DomainName: *string, // Required
		// EventTriggerName: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesEventTriggerName) > 0 {
		input.EventTriggerName = aws.String(_customerprofilesEventTriggerName)
	}

	if resp, err := client.DeleteEventTrigger(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes an integration from a specific domain.
func customerprofiles_DeleteIntegration(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.DeleteIntegrationInput{
		// DomainName: *string, // Required
		// Uri: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesUri) > 0 {
		input.Uri = aws.String(_customerprofilesUri)
	}

	if resp, err := client.DeleteIntegration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the standard customer profile and all data pertaining to the profile.
func customerprofiles_DeleteProfile(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.DeleteProfileInput{
		// DomainName: *string, // Required
		// ProfileId: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesProfileId) > 0 {
		input.ProfileId = aws.String(_customerprofilesProfileId)
	}

	if resp, err := client.DeleteProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a searchable key from a customer profile.
func customerprofiles_DeleteProfileKey(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.DeleteProfileKeyInput{
		// DomainName: *string, // Required
		// KeyName: *string, // Required
		// ProfileId: *string, // Required
		// Values: []string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesKeyName) > 0 {
		input.KeyName = aws.String(_customerprofilesKeyName)
	}
	if len(_customerprofilesProfileId) > 0 {
		input.ProfileId = aws.String(_customerprofilesProfileId)
	}
	if len(_customerprofilesValues) > 0 {
		input.Values = append([]string(nil), _customerprofilesValues...)
	}

	if resp, err := client.DeleteProfileKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes an object associated with a profile of a given ProfileObjectType.
func customerprofiles_DeleteProfileObject(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.DeleteProfileObjectInput{
		// DomainName: *string, // Required
		// ObjectTypeName: *string, // Required
		// ProfileId: *string, // Required
		// ProfileObjectUniqueKey: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesObjectTypeName) > 0 {
		input.ObjectTypeName = aws.String(_customerprofilesObjectTypeName)
	}
	if len(_customerprofilesProfileId) > 0 {
		input.ProfileId = aws.String(_customerprofilesProfileId)
	}
	if len(_customerprofilesProfileObjectUniqueKey) > 0 {
		input.ProfileObjectUniqueKey = aws.String(_customerprofilesProfileObjectUniqueKey)
	}

	if resp, err := client.DeleteProfileObject(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a ProfileObjectType from a specific domain as well as removes all the
// ProfileObjects of that type. It also disables integrations from this specific
// ProfileObjectType. In addition, it scrubs all of the fields of the standard
// profile that were populated from this ProfileObjectType.
func customerprofiles_DeleteProfileObjectType(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.DeleteProfileObjectTypeInput{
		// DomainName: *string, // Required
		// ObjectTypeName: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesObjectTypeName) > 0 {
		input.ObjectTypeName = aws.String(_customerprofilesObjectTypeName)
	}

	if resp, err := client.DeleteProfileObjectType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a recommender.
func customerprofiles_DeleteRecommender(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.DeleteRecommenderInput{
		// DomainName: *string, // Required
		// RecommenderName: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesRecommenderName) > 0 {
		input.RecommenderName = aws.String(_customerprofilesRecommenderName)
	}

	if resp, err := client.DeleteRecommender(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a segment definition from the domain.
func customerprofiles_DeleteSegmentDefinition(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.DeleteSegmentDefinitionInput{
		// DomainName: *string, // Required
		// SegmentDefinitionName: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesSegmentDefinitionName) > 0 {
		input.SegmentDefinitionName = aws.String(_customerprofilesSegmentDefinitionName)
	}

	if resp, err := client.DeleteSegmentDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified workflow and all its corresponding resources. This is an
// async process.
func customerprofiles_DeleteWorkflow(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.DeleteWorkflowInput{
		// DomainName: *string, // Required
		// WorkflowId: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesWorkflowId) > 0 {
		input.WorkflowId = aws.String(_customerprofilesWorkflowId)
	}

	if resp, err := client.DeleteWorkflow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The process of detecting profile object type mapping by using given objects.
func customerprofiles_DetectProfileObjectType(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.DetectProfileObjectTypeInput{
		// DomainName: *string, // Required
		// Objects: []string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesObjects) > 0 {
		input.Objects = append([]string(nil), _customerprofilesObjects...)
	}

	if resp, err := client.DetectProfileObjectType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Tests the auto-merging settings of your Identity Resolution Job without merging
// your data. It randomly selects a sample of matching groups from the existing
// matching results, and applies the automerging settings that you provided. You
// can then view the number of profiles in the sample, the number of matches, and
// the number of profiles identified to be merged. This enables you to evaluate the
// accuracy of the attributes in your matching list.
//
// You can't view which profiles are matched and would be merged.
//
// We strongly recommend you use this API to do a dry run of the automerging
// process before running the Identity Resolution Job. Include at least two
// matching attributes. If your matching list includes too few attributes (such as
// only FirstName or only LastName ), there may be a large number of matches. This
// increases the chances of erroneous merges.
func customerprofiles_GetAutoMergingPreview(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.GetAutoMergingPreviewInput{
		// ConflictResolution: *types.ConflictResolution, // Required
		// Consolidation: *types.Consolidation, // Required
		// DomainName: *string, // Required
	}

	if len(_customerprofilesConflictResolution) > 0 {
		if err := assignInputField(input, "ConflictResolution", _customerprofilesConflictResolution); err != nil {
			log.Errorf("invalid --conflict-resolution: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesConsolidation) > 0 {
		if err := assignInputField(input, "Consolidation", _customerprofilesConsolidation); err != nil {
			log.Errorf("invalid --consolidation: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesMinAllowedConfidenceScoreForMerging) > 0 {
		if err := assignInputField(input, "MinAllowedConfidenceScoreForMerging", _customerprofilesMinAllowedConfidenceScoreForMerging); err != nil {
			log.Errorf("invalid --min-allowed-confidence-score-for-merging: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetAutoMergingPreview(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides more information on a calculated attribute definition for Customer
// Profiles.
func customerprofiles_GetCalculatedAttributeDefinition(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.GetCalculatedAttributeDefinitionInput{
		// CalculatedAttributeName: *string, // Required
		// DomainName: *string, // Required
	}

	if len(_customerprofilesCalculatedAttributeName) > 0 {
		input.CalculatedAttributeName = aws.String(_customerprofilesCalculatedAttributeName)
	}
	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}

	if resp, err := client.GetCalculatedAttributeDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieve a calculated attribute for a customer profile.
func customerprofiles_GetCalculatedAttributeForProfile(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.GetCalculatedAttributeForProfileInput{
		// CalculatedAttributeName: *string, // Required
		// DomainName: *string, // Required
		// ProfileId: *string, // Required
	}

	if len(_customerprofilesCalculatedAttributeName) > 0 {
		input.CalculatedAttributeName = aws.String(_customerprofilesCalculatedAttributeName)
	}
	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesProfileId) > 0 {
		input.ProfileId = aws.String(_customerprofilesProfileId)
	}

	if resp, err := client.GetCalculatedAttributeForProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a specific domain.
func customerprofiles_GetDomain(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.GetDomainInput{
		// DomainName: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}

	if resp, err := client.GetDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the layout to view data for a specific domain. This API can only be
// invoked from the Amazon Connect admin website.
func customerprofiles_GetDomainLayout(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.GetDomainLayoutInput{
		// DomainName: *string, // Required
		// LayoutDefinitionName: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesLayoutDefinitionName) > 0 {
		input.LayoutDefinitionName = aws.String(_customerprofilesLayoutDefinitionName)
	}

	if resp, err := client.GetDomainLayout(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Return a DomainObjectType for the input Domain and ObjectType names.
func customerprofiles_GetDomainObjectType(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.GetDomainObjectTypeInput{
		// DomainName: *string, // Required
		// ObjectTypeName: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesObjectTypeName) > 0 {
		input.ObjectTypeName = aws.String(_customerprofilesObjectTypeName)
	}

	if resp, err := client.GetDomainObjectType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the specified event stream in a specific domain.
func customerprofiles_GetEventStream(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.GetEventStreamInput{
		// DomainName: *string, // Required
		// EventStreamName: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesEventStreamName) > 0 {
		input.EventStreamName = aws.String(_customerprofilesEventStreamName)
	}

	if resp, err := client.GetEventStream(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get a specific Event Trigger from the domain.
func customerprofiles_GetEventTrigger(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.GetEventTriggerInput{
		// DomainName: *string, // Required
		// EventTriggerName: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesEventTriggerName) > 0 {
		input.EventTriggerName = aws.String(_customerprofilesEventTriggerName)
	}

	if resp, err := client.GetEventTrigger(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about an Identity Resolution Job in a specific domain.
// Identity Resolution Jobs are set up using the Amazon Connect admin console. For
// more information, see [Use Identity Resolution to consolidate similar profiles].
//
// [Use Identity Resolution to consolidate similar profiles]: https://docs.aws.amazon.com/connect/latest/adminguide/use-identity-resolution.html
func customerprofiles_GetIdentityResolutionJob(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.GetIdentityResolutionJobInput{
		// DomainName: *string, // Required
		// JobId: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesJobId) > 0 {
		input.JobId = aws.String(_customerprofilesJobId)
	}

	if resp, err := client.GetIdentityResolutionJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns an integration for a domain.
func customerprofiles_GetIntegration(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.GetIntegrationInput{
		// DomainName: *string, // Required
		// Uri: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesUri) > 0 {
		input.Uri = aws.String(_customerprofilesUri)
	}

	if resp, err := client.GetIntegration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Before calling this API, use [CreateDomain] or [UpdateDomain] to enable identity resolution: set Matching
// to true.
//
// GetMatches returns potentially matching profiles, based on the results of the
// latest run of a machine learning process.
//
// The process of matching duplicate profiles. If Matching = true , Amazon Connect
// Customer Profiles starts a weekly batch process called Identity Resolution Job.
// If you do not specify a date and time for Identity Resolution Job to run, by
// default it runs every Saturday at 12AM UTC to detect duplicate profiles in your
// domains.
//
// After the Identity Resolution Job completes, use the [GetMatches] API to return and review
// the results. Or, if you have configured ExportingConfig in the MatchingRequest ,
// you can download the results from S3.
//
// Amazon Connect uses the following profile attributes to identify matches:
//
// - PhoneNumber
//
// - HomePhoneNumber
//
// - BusinessPhoneNumber
//
// - MobilePhoneNumber
//
// - EmailAddress
//
// - PersonalEmailAddress
//
// - BusinessEmailAddress
//
// - FullName
//
// For example, two or more profiles—with spelling mistakes such as John Doe and
// Jhn Doe, or different casing email addresses such as JOHN_DOE(at)ANYCOMPANY.COM and
// johndoe(at)anycompany.com, or different phone number formats such as 555-010-0000
// and +1-555-010-0000—can be detected as belonging to the same customer John Doe
// and merged into a unified profile.
//
// [GetMatches]: https://docs.aws.amazon.com/customerprofiles/latest/APIReference/API_GetMatches.html
// [UpdateDomain]: https://docs.aws.amazon.com/customerprofiles/latest/APIReference/API_UpdateDomain.html
// [CreateDomain]: https://docs.aws.amazon.com/customerprofiles/latest/APIReference/API_CreateDomain.html
func customerprofiles_GetMatches(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.GetMatchesInput{
		// DomainName: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _customerprofilesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesNextToken) > 0 {
		input.NextToken = aws.String(_customerprofilesNextToken)
	}

	if resp, err := client.GetMatches(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The GetObjectTypeAttributeValues API delivers statistical insights about
// attributes within a specific object type, but is exclusively available for
// domains with data store enabled. This API performs daily calculations to provide
// statistical information about your attribute values, helping you understand
// patterns and trends in your data. The statistical calculations are performed
// once per day, providing a consistent snapshot of your attribute data
// characteristics.
//
// You'll receive null values in two scenarios:
//
// During the first period after enabling data vault (unless a calculation cycle
// occurs, which happens once daily).
//
// For attributes that don't contain numeric values.
func customerprofiles_GetObjectTypeAttributeStatistics(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.GetObjectTypeAttributeStatisticsInput{
		// AttributeName: *string, // Required
		// DomainName: *string, // Required
		// ObjectTypeName: *string, // Required
	}

	if len(_customerprofilesAttributeName) > 0 {
		input.AttributeName = aws.String(_customerprofilesAttributeName)
	}
	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesObjectTypeName) > 0 {
		input.ObjectTypeName = aws.String(_customerprofilesObjectTypeName)
	}

	if resp, err := client.GetObjectTypeAttributeStatistics(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a history record for a specific profile, for a specific domain.
func customerprofiles_GetProfileHistoryRecord(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.GetProfileHistoryRecordInput{
		// DomainName: *string, // Required
		// Id: *string, // Required
		// ProfileId: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesId) > 0 {
		input.Id = aws.String(_customerprofilesId)
	}
	if len(_customerprofilesProfileId) > 0 {
		input.ProfileId = aws.String(_customerprofilesProfileId)
	}

	if resp, err := client.GetProfileHistoryRecord(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the object types for a specific domain.
func customerprofiles_GetProfileObjectType(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.GetProfileObjectTypeInput{
		// DomainName: *string, // Required
		// ObjectTypeName: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesObjectTypeName) > 0 {
		input.ObjectTypeName = aws.String(_customerprofilesObjectTypeName)
	}

	if resp, err := client.GetProfileObjectType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the template information for a specific object type.
// A template is a predefined ProfileObjectType, such as “Salesforce-Account” or
// “Salesforce-Contact.” When a user sends a ProfileObject, using the
// PutProfileObject API, with an ObjectTypeName that matches one of the
// TemplateIds, it uses the mappings from the template.
func customerprofiles_GetProfileObjectTypeTemplate(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.GetProfileObjectTypeTemplateInput{
		// TemplateId: *string, // Required
	}

	if len(_customerprofilesTemplateId) > 0 {
		input.TemplateId = aws.String(_customerprofilesTemplateId)
	}

	if resp, err := client.GetProfileObjectTypeTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Fetches the recommendations for a profile in the input Customer Profiles
// domain. Fetches all the profile recommendations
func customerprofiles_GetProfileRecommendations(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.GetProfileRecommendationsInput{
		// DomainName: *string, // Required
		// ProfileId: *string, // Required
		// RecommenderName: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesProfileId) > 0 {
		input.ProfileId = aws.String(_customerprofilesProfileId)
	}
	if len(_customerprofilesRecommenderName) > 0 {
		input.RecommenderName = aws.String(_customerprofilesRecommenderName)
	}
	if len(_customerprofilesContext) > 0 {
		if err := assignInputField(input, "Context", _customerprofilesContext); err != nil {
			log.Errorf("invalid --context: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _customerprofilesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetProfileRecommendations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a recommender.
func customerprofiles_GetRecommender(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.GetRecommenderInput{
		// DomainName: *string, // Required
		// RecommenderName: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesRecommenderName) > 0 {
		input.RecommenderName = aws.String(_customerprofilesRecommenderName)
	}
	if len(_customerprofilesTrainingMetricsCount) > 0 {
		if err := assignInputField(input, "TrainingMetricsCount", _customerprofilesTrainingMetricsCount); err != nil {
			log.Errorf("invalid --training-metrics-count: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetRecommender(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a segment definition from the domain.
func customerprofiles_GetSegmentDefinition(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.GetSegmentDefinitionInput{
		// DomainName: *string, // Required
		// SegmentDefinitionName: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesSegmentDefinitionName) > 0 {
		input.SegmentDefinitionName = aws.String(_customerprofilesSegmentDefinitionName)
	}

	if resp, err := client.GetSegmentDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the result of a segment estimate query.
func customerprofiles_GetSegmentEstimate(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.GetSegmentEstimateInput{
		// DomainName: *string, // Required
		// EstimateId: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesEstimateId) > 0 {
		input.EstimateId = aws.String(_customerprofilesEstimateId)
	}

	if resp, err := client.GetSegmentEstimate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Determines if the given profiles are within a segment.
func customerprofiles_GetSegmentMembership(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.GetSegmentMembershipInput{
		// DomainName: *string, // Required
		// ProfileIds: []string, // Required
		// SegmentDefinitionName: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesProfileIds) > 0 {
		input.ProfileIds = append([]string(nil), _customerprofilesProfileIds...)
	}
	if len(_customerprofilesSegmentDefinitionName) > 0 {
		input.SegmentDefinitionName = aws.String(_customerprofilesSegmentDefinitionName)
	}

	if resp, err := client.GetSegmentMembership(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieve the latest status of a segment snapshot.
func customerprofiles_GetSegmentSnapshot(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.GetSegmentSnapshotInput{
		// DomainName: *string, // Required
		// SegmentDefinitionName: *string, // Required
		// SnapshotId: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesSegmentDefinitionName) > 0 {
		input.SegmentDefinitionName = aws.String(_customerprofilesSegmentDefinitionName)
	}
	if len(_customerprofilesSnapshotId) > 0 {
		input.SnapshotId = aws.String(_customerprofilesSnapshotId)
	}

	if resp, err := client.GetSegmentSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a set of profiles that belong to the same matching group using the
// matchId or profileId . You can also specify the type of matching that you want
// for finding similar profiles using either RULE_BASED_MATCHING or
// ML_BASED_MATCHING .
func customerprofiles_GetSimilarProfiles(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.GetSimilarProfilesInput{
		// DomainName: *string, // Required
		// MatchType: types.MatchType, // Required
		// SearchKey: *string, // Required
		// SearchValue: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesMatchType) > 0 {
		if err := assignInputField(input, "MatchType", _customerprofilesMatchType); err != nil {
			log.Errorf("invalid --match-type: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesSearchKey) > 0 {
		input.SearchKey = aws.String(_customerprofilesSearchKey)
	}
	if len(_customerprofilesSearchValue) > 0 {
		input.SearchValue = aws.String(_customerprofilesSearchValue)
	}
	if len(_customerprofilesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _customerprofilesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesNextToken) > 0 {
		input.NextToken = aws.String(_customerprofilesNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetSimilarProfiles(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*customerprofiles.GetSimilarProfilesOutput
	p := customerprofiles.NewGetSimilarProfilesPaginator(client, input)
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

// This API retrieves the details of a specific upload job.
func customerprofiles_GetUploadJob(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.GetUploadJobInput{
		// DomainName: *string, // Required
		// JobId: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesJobId) > 0 {
		input.JobId = aws.String(_customerprofilesJobId)
	}

	if resp, err := client.GetUploadJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This API retrieves the pre-signed URL and client token for uploading the file
// associated with the upload job.
func customerprofiles_GetUploadJobPath(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.GetUploadJobPathInput{
		// DomainName: *string, // Required
		// JobId: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesJobId) > 0 {
		input.JobId = aws.String(_customerprofilesJobId)
	}

	if resp, err := client.GetUploadJobPath(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get details of specified workflow.
func customerprofiles_GetWorkflow(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.GetWorkflowInput{
		// DomainName: *string, // Required
		// WorkflowId: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesWorkflowId) > 0 {
		input.WorkflowId = aws.String(_customerprofilesWorkflowId)
	}

	if resp, err := client.GetWorkflow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get granular list of steps in workflow.
func customerprofiles_GetWorkflowSteps(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.GetWorkflowStepsInput{
		// DomainName: *string, // Required
		// WorkflowId: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesWorkflowId) > 0 {
		input.WorkflowId = aws.String(_customerprofilesWorkflowId)
	}
	if len(_customerprofilesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _customerprofilesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesNextToken) > 0 {
		input.NextToken = aws.String(_customerprofilesNextToken)
	}

	if resp, err := client.GetWorkflowSteps(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all of the integrations associated to a specific URI in the AWS account.
func customerprofiles_ListAccountIntegrations(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.ListAccountIntegrationsInput{
		// Uri: *string, // Required
	}

	if len(_customerprofilesUri) > 0 {
		input.Uri = aws.String(_customerprofilesUri)
	}
	if len(_customerprofilesIncludeHidden) > 0 {
		if err := assignInputField(input, "IncludeHidden", _customerprofilesIncludeHidden); err != nil {
			log.Errorf("invalid --include-hidden: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _customerprofilesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesNextToken) > 0 {
		input.NextToken = aws.String(_customerprofilesNextToken)
	}

	if resp, err := client.ListAccountIntegrations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists calculated attribute definitions for Customer Profiles
func customerprofiles_ListCalculatedAttributeDefinitions(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.ListCalculatedAttributeDefinitionsInput{
		// DomainName: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _customerprofilesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesNextToken) > 0 {
		input.NextToken = aws.String(_customerprofilesNextToken)
	}

	if resp, err := client.ListCalculatedAttributeDefinitions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieve a list of calculated attributes for a customer profile.
func customerprofiles_ListCalculatedAttributesForProfile(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.ListCalculatedAttributesForProfileInput{
		// DomainName: *string, // Required
		// ProfileId: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesProfileId) > 0 {
		input.ProfileId = aws.String(_customerprofilesProfileId)
	}
	if len(_customerprofilesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _customerprofilesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesNextToken) > 0 {
		input.NextToken = aws.String(_customerprofilesNextToken)
	}

	if resp, err := client.ListCalculatedAttributesForProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the existing layouts that can be used to view data for a specific domain.
// This API can only be invoked from the Amazon Connect admin website.
func customerprofiles_ListDomainLayouts(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.ListDomainLayoutsInput{
		// DomainName: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _customerprofilesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesNextToken) > 0 {
		input.NextToken = aws.String(_customerprofilesNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDomainLayouts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*customerprofiles.ListDomainLayoutsOutput
	p := customerprofiles.NewListDomainLayoutsPaginator(client, input)
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

// List all DomainObjectType(s) in a Customer Profiles domain.
func customerprofiles_ListDomainObjectTypes(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.ListDomainObjectTypesInput{
		// DomainName: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _customerprofilesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesNextToken) > 0 {
		input.NextToken = aws.String(_customerprofilesNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDomainObjectTypes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*customerprofiles.ListDomainObjectTypesOutput
	p := customerprofiles.NewListDomainObjectTypesPaginator(client, input)
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

// Returns a list of all the domains for an AWS account that have been created.
func customerprofiles_ListDomains(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.ListDomainsInput{}

	if len(_customerprofilesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _customerprofilesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesNextToken) > 0 {
		input.NextToken = aws.String(_customerprofilesNextToken)
	}

	if resp, err := client.ListDomains(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of all the event streams in a specific domain.
func customerprofiles_ListEventStreams(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.ListEventStreamsInput{
		// DomainName: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _customerprofilesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesNextToken) > 0 {
		input.NextToken = aws.String(_customerprofilesNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListEventStreams(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*customerprofiles.ListEventStreamsOutput
	p := customerprofiles.NewListEventStreamsPaginator(client, input)
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

// List all Event Triggers under a domain.
func customerprofiles_ListEventTriggers(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.ListEventTriggersInput{
		// DomainName: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _customerprofilesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesNextToken) > 0 {
		input.NextToken = aws.String(_customerprofilesNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListEventTriggers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*customerprofiles.ListEventTriggersOutput
	p := customerprofiles.NewListEventTriggersPaginator(client, input)
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

// Lists all of the Identity Resolution Jobs in your domain. The response sorts
// the list by JobStartTime .
func customerprofiles_ListIdentityResolutionJobs(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.ListIdentityResolutionJobsInput{
		// DomainName: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _customerprofilesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesNextToken) > 0 {
		input.NextToken = aws.String(_customerprofilesNextToken)
	}

	if resp, err := client.ListIdentityResolutionJobs(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all of the integrations in your domain.
func customerprofiles_ListIntegrations(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.ListIntegrationsInput{
		// DomainName: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesIncludeHidden) > 0 {
		if err := assignInputField(input, "IncludeHidden", _customerprofilesIncludeHidden); err != nil {
			log.Errorf("invalid --include-hidden: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _customerprofilesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesNextToken) > 0 {
		input.NextToken = aws.String(_customerprofilesNextToken)
	}

	if resp, err := client.ListIntegrations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The ListObjectTypeAttributeValues API provides access to the most recent
// distinct values for any specified attribute, making it valuable for real-time
// data validation and consistency checks within your object types. This API works
// across domain, supporting both custom and standard object types. The API accepts
// the object type name, attribute name, and domain name as input parameters and
// returns values up to the storage limit of approximately 350KB.
func customerprofiles_ListObjectTypeAttributeValues(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.ListObjectTypeAttributeValuesInput{
		// AttributeName: *string, // Required
		// DomainName: *string, // Required
		// ObjectTypeName: *string, // Required
	}

	if len(_customerprofilesAttributeName) > 0 {
		input.AttributeName = aws.String(_customerprofilesAttributeName)
	}
	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesObjectTypeName) > 0 {
		input.ObjectTypeName = aws.String(_customerprofilesObjectTypeName)
	}
	if len(_customerprofilesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _customerprofilesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesNextToken) > 0 {
		input.NextToken = aws.String(_customerprofilesNextToken)
	}

	if resp, err := client.ListObjectTypeAttributeValues(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Fetch the possible attribute values given the attribute name.
func customerprofiles_ListObjectTypeAttributes(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.ListObjectTypeAttributesInput{
		// DomainName: *string, // Required
		// ObjectTypeName: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesObjectTypeName) > 0 {
		input.ObjectTypeName = aws.String(_customerprofilesObjectTypeName)
	}
	if len(_customerprofilesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _customerprofilesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesNextToken) > 0 {
		input.NextToken = aws.String(_customerprofilesNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListObjectTypeAttributes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*customerprofiles.ListObjectTypeAttributesOutput
	p := customerprofiles.NewListObjectTypeAttributesPaginator(client, input)
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

// Fetch the possible attribute values given the attribute name.
func customerprofiles_ListProfileAttributeValues(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.ListProfileAttributeValuesInput{
		// AttributeName: *string, // Required
		// DomainName: *string, // Required
	}

	if len(_customerprofilesAttributeName) > 0 {
		input.AttributeName = aws.String(_customerprofilesAttributeName)
	}
	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}

	if resp, err := client.ListProfileAttributeValues(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of history records for a specific profile, for a specific domain.
func customerprofiles_ListProfileHistoryRecords(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.ListProfileHistoryRecordsInput{
		// DomainName: *string, // Required
		// ProfileId: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesProfileId) > 0 {
		input.ProfileId = aws.String(_customerprofilesProfileId)
	}
	if len(_customerprofilesActionType) > 0 {
		if err := assignInputField(input, "ActionType", _customerprofilesActionType); err != nil {
			log.Errorf("invalid --action-type: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _customerprofilesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesNextToken) > 0 {
		input.NextToken = aws.String(_customerprofilesNextToken)
	}
	if len(_customerprofilesObjectTypeName) > 0 {
		input.ObjectTypeName = aws.String(_customerprofilesObjectTypeName)
	}
	if len(_customerprofilesPerformedBy) > 0 {
		input.PerformedBy = aws.String(_customerprofilesPerformedBy)
	}

	if resp, err := client.ListProfileHistoryRecords(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all of the template information for object types.
func customerprofiles_ListProfileObjectTypeTemplates(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.ListProfileObjectTypeTemplatesInput{}

	if len(_customerprofilesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _customerprofilesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesNextToken) > 0 {
		input.NextToken = aws.String(_customerprofilesNextToken)
	}

	if resp, err := client.ListProfileObjectTypeTemplates(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all of the templates available within the service.
func customerprofiles_ListProfileObjectTypes(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.ListProfileObjectTypesInput{
		// DomainName: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _customerprofilesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesNextToken) > 0 {
		input.NextToken = aws.String(_customerprofilesNextToken)
	}

	if resp, err := client.ListProfileObjectTypes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of objects associated with a profile of a given
// ProfileObjectType.
func customerprofiles_ListProfileObjects(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.ListProfileObjectsInput{
		// DomainName: *string, // Required
		// ObjectTypeName: *string, // Required
		// ProfileId: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesObjectTypeName) > 0 {
		input.ObjectTypeName = aws.String(_customerprofilesObjectTypeName)
	}
	if len(_customerprofilesProfileId) > 0 {
		input.ProfileId = aws.String(_customerprofilesProfileId)
	}
	if len(_customerprofilesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _customerprofilesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesNextToken) > 0 {
		input.NextToken = aws.String(_customerprofilesNextToken)
	}
	if len(_customerprofilesObjectFilter) > 0 {
		if err := assignInputField(input, "ObjectFilter", _customerprofilesObjectFilter); err != nil {
			log.Errorf("invalid --object-filter: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListProfileObjects(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of available recommender recipes that can be used to create
// recommenders.
func customerprofiles_ListRecommenderRecipes(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.ListRecommenderRecipesInput{}

	if len(_customerprofilesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _customerprofilesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesNextToken) > 0 {
		input.NextToken = aws.String(_customerprofilesNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRecommenderRecipes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*customerprofiles.ListRecommenderRecipesOutput
	p := customerprofiles.NewListRecommenderRecipesPaginator(client, input)
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

// Returns a list of recommenders in the specified domain.
func customerprofiles_ListRecommenders(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.ListRecommendersInput{
		// DomainName: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _customerprofilesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesNextToken) > 0 {
		input.NextToken = aws.String(_customerprofilesNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRecommenders(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*customerprofiles.ListRecommendersOutput
	p := customerprofiles.NewListRecommendersPaginator(client, input)
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

// Returns a set of MatchIds that belong to the given domain.
func customerprofiles_ListRuleBasedMatches(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.ListRuleBasedMatchesInput{
		// DomainName: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _customerprofilesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesNextToken) > 0 {
		input.NextToken = aws.String(_customerprofilesNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRuleBasedMatches(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*customerprofiles.ListRuleBasedMatchesOutput
	p := customerprofiles.NewListRuleBasedMatchesPaginator(client, input)
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

// Lists all segment definitions under a domain.
func customerprofiles_ListSegmentDefinitions(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.ListSegmentDefinitionsInput{
		// DomainName: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _customerprofilesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesNextToken) > 0 {
		input.NextToken = aws.String(_customerprofilesNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSegmentDefinitions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*customerprofiles.ListSegmentDefinitionsOutput
	p := customerprofiles.NewListSegmentDefinitionsPaginator(client, input)
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

// Displays the tags associated with an Amazon Connect Customer Profiles resource.
// In Connect Customer Profiles, domains, profile object types, and integrations
// can be tagged.
func customerprofiles_ListTagsForResource(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_customerprofilesResourceArn) > 0 {
		input.ResourceArn = aws.String(_customerprofilesResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This API retrieves a list of upload jobs for the specified domain.
func customerprofiles_ListUploadJobs(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.ListUploadJobsInput{
		// DomainName: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _customerprofilesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesNextToken) > 0 {
		input.NextToken = aws.String(_customerprofilesNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListUploadJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*customerprofiles.ListUploadJobsOutput
	p := customerprofiles.NewListUploadJobsPaginator(client, input)
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

// Query to list all workflows.
func customerprofiles_ListWorkflows(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.ListWorkflowsInput{
		// DomainName: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _customerprofilesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesNextToken) > 0 {
		input.NextToken = aws.String(_customerprofilesNextToken)
	}
	if len(_customerprofilesQueryEndDate) > 0 {
		if err := assignInputField(input, "QueryEndDate", _customerprofilesQueryEndDate); err != nil {
			log.Errorf("invalid --query-end-date: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesQueryStartDate) > 0 {
		if err := assignInputField(input, "QueryStartDate", _customerprofilesQueryStartDate); err != nil {
			log.Errorf("invalid --query-start-date: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesStatus) > 0 {
		if err := assignInputField(input, "Status", _customerprofilesStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesWorkflowType) > 0 {
		if err := assignInputField(input, "WorkflowType", _customerprofilesWorkflowType); err != nil {
			log.Errorf("invalid --workflow-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListWorkflows(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Runs an AWS Lambda job that does the following:
// - All the profileKeys in the ProfileToBeMerged will be moved to the main
// profile.
//
// - All the objects in the ProfileToBeMerged will be moved to the main profile.
//
// - All the ProfileToBeMerged will be deleted at the end.
//
// - All the profileKeys in the ProfileIdsToBeMerged will be moved to the main
// profile.
//
// - Standard fields are merged as follows:
//
// - Fields are always "union"-ed if there are no conflicts in standard fields
// or attributeKeys.
//
// - When there are conflicting fields:
//
// - If no SourceProfileIds entry is specified, the main Profile value is always
// taken.
//
// - If a SourceProfileIds entry is specified, the specified profileId is always
// taken, even if it is a NULL value.
//
// You can use MergeProfiles together with [GetMatches], which returns potentially matching
// profiles, or use it with the results of another matching system. After profiles
// have been merged, they cannot be separated (unmerged).
//
// [GetMatches]: https://docs.aws.amazon.com/customerprofiles/latest/APIReference/API_GetMatches.html
func customerprofiles_MergeProfiles(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.MergeProfilesInput{
		// DomainName: *string, // Required
		// MainProfileId: *string, // Required
		// ProfileIdsToBeMerged: []string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesMainProfileId) > 0 {
		input.MainProfileId = aws.String(_customerprofilesMainProfileId)
	}
	if len(_customerprofilesProfileIdsToBeMerged) > 0 {
		input.ProfileIdsToBeMerged = append([]string(nil), _customerprofilesProfileIdsToBeMerged...)
	}
	if len(_customerprofilesFieldSourceProfileIds) > 0 {
		if err := assignInputField(input, "FieldSourceProfileIds", _customerprofilesFieldSourceProfileIds); err != nil {
			log.Errorf("invalid --field-source-profile-ids: %s", err.Error())
			return
		}
	}

	if resp, err := client.MergeProfiles(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create/Update a DomainObjectType in a Customer Profiles domain. To create a new
// DomainObjectType, Data Store needs to be enabled on the Domain.
func customerprofiles_PutDomainObjectType(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.PutDomainObjectTypeInput{
		// DomainName: *string, // Required
		// Fields: map[string]types.DomainObjectTypeField, // Required
		// ObjectTypeName: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesFields) > 0 {
		if err := assignInputField(input, "Fields", _customerprofilesFields); err != nil {
			log.Errorf("invalid --fields: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesObjectTypeName) > 0 {
		input.ObjectTypeName = aws.String(_customerprofilesObjectTypeName)
	}
	if len(_customerprofilesDescription) > 0 {
		input.Description = aws.String(_customerprofilesDescription)
	}
	if len(_customerprofilesEncryptionKey) > 0 {
		input.EncryptionKey = aws.String(_customerprofilesEncryptionKey)
	}
	if len(_customerprofilesTags) > 0 {
		if err := assignInputField(input, "Tags", _customerprofilesTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutDomainObjectType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds an integration between the service and a third-party service, which
// includes Amazon AppFlow and Amazon Connect.
//
// An integration can belong to only one domain.
//
// To add or remove tags on an existing Integration, see [TagResource]/[UntagResource] .
//
// [TagResource]: https://docs.aws.amazon.com/customerprofiles/latest/APIReference/API_TagResource.html
// [UntagResource]: https://docs.aws.amazon.com/customerprofiles/latest/APIReference/API_UntagResource.html
func customerprofiles_PutIntegration(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.PutIntegrationInput{
		// DomainName: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesEventTriggerNames) > 0 {
		input.EventTriggerNames = append([]string(nil), _customerprofilesEventTriggerNames...)
	}
	if len(_customerprofilesFlowDefinition) > 0 {
		if err := assignInputField(input, "FlowDefinition", _customerprofilesFlowDefinition); err != nil {
			log.Errorf("invalid --flow-definition: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesObjectTypeName) > 0 {
		input.ObjectTypeName = aws.String(_customerprofilesObjectTypeName)
	}
	if len(_customerprofilesObjectTypeNames) > 0 {
		if err := assignInputField(input, "ObjectTypeNames", _customerprofilesObjectTypeNames); err != nil {
			log.Errorf("invalid --object-type-names: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesRoleArn) > 0 {
		input.RoleArn = aws.String(_customerprofilesRoleArn)
	}
	if len(_customerprofilesScope) > 0 {
		if err := assignInputField(input, "Scope", _customerprofilesScope); err != nil {
			log.Errorf("invalid --scope: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesTags) > 0 {
		if err := assignInputField(input, "Tags", _customerprofilesTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesUri) > 0 {
		input.Uri = aws.String(_customerprofilesUri)
	}

	if resp, err := client.PutIntegration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds additional objects to customer profiles of a given ObjectType.
// When adding a specific profile object, like a Contact Record, an inferred
// profile can get created if it is not mapped to an existing profile. The
// resulting profile will only have a phone number populated in the standard
// ProfileObject. Any additional Contact Records with the same phone number will be
// mapped to the same inferred profile.
//
// When a ProfileObject is created and if a ProfileObjectType already exists for
// the ProfileObject, it will provide data to a standard profile depending on the
// ProfileObjectType definition.
//
// PutProfileObject needs an ObjectType, which can be created using
// PutProfileObjectType.
func customerprofiles_PutProfileObject(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.PutProfileObjectInput{
		// DomainName: *string, // Required
		// Object: *string, // Required
		// ObjectTypeName: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesObject) > 0 {
		input.Object = aws.String(_customerprofilesObject)
	}
	if len(_customerprofilesObjectTypeName) > 0 {
		input.ObjectTypeName = aws.String(_customerprofilesObjectTypeName)
	}

	if resp, err := client.PutProfileObject(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Defines a ProfileObjectType.
// To add or remove tags on an existing ObjectType, see [TagResource]/[UntagResource] .
//
// [TagResource]: https://docs.aws.amazon.com/customerprofiles/latest/APIReference/API_TagResource.html
// [UntagResource]: https://docs.aws.amazon.com/customerprofiles/latest/APIReference/API_UntagResource.html
func customerprofiles_PutProfileObjectType(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.PutProfileObjectTypeInput{
		// Description: *string, // Required
		// DomainName: *string, // Required
		// ObjectTypeName: *string, // Required
	}

	if len(_customerprofilesDescription) > 0 {
		input.Description = aws.String(_customerprofilesDescription)
	}
	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesObjectTypeName) > 0 {
		input.ObjectTypeName = aws.String(_customerprofilesObjectTypeName)
	}
	if len(_customerprofilesAllowProfileCreation) > 0 {
		if err := assignInputField(input, "AllowProfileCreation", _customerprofilesAllowProfileCreation); err != nil {
			log.Errorf("invalid --allow-profile-creation: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesEncryptionKey) > 0 {
		input.EncryptionKey = aws.String(_customerprofilesEncryptionKey)
	}
	if len(_customerprofilesExpirationDays) > 0 {
		if err := assignInputField(input, "ExpirationDays", _customerprofilesExpirationDays); err != nil {
			log.Errorf("invalid --expiration-days: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesFields) > 0 {
		if err := assignInputField(input, "Fields", _customerprofilesFields); err != nil {
			log.Errorf("invalid --fields: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesKeys) > 0 {
		if err := assignInputField(input, "Keys", _customerprofilesKeys); err != nil {
			log.Errorf("invalid --keys: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesMaxProfileObjectCount) > 0 {
		if err := assignInputField(input, "MaxProfileObjectCount", _customerprofilesMaxProfileObjectCount); err != nil {
			log.Errorf("invalid --max-profile-object-count: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesSourceLastUpdatedTimestampFormat) > 0 {
		input.SourceLastUpdatedTimestampFormat = aws.String(_customerprofilesSourceLastUpdatedTimestampFormat)
	}
	if len(_customerprofilesSourcePriority) > 0 {
		if err := assignInputField(input, "SourcePriority", _customerprofilesSourcePriority); err != nil {
			log.Errorf("invalid --source-priority: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesTags) > 0 {
		if err := assignInputField(input, "Tags", _customerprofilesTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesTemplateId) > 0 {
		input.TemplateId = aws.String(_customerprofilesTemplateId)
	}

	if resp, err := client.PutProfileObjectType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Searches for profiles within a specific domain using one or more predefined
// search keys (e.g., _fullName, _phone, _email, _account, etc.) and/or
// custom-defined search keys. A search key is a data type pair that consists of a
// KeyName and Values list.
//
// This operation supports searching for profiles with a minimum of 1 key-value(s)
// pair and up to 5 key-value(s) pairs using either AND or OR logic.
func customerprofiles_SearchProfiles(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.SearchProfilesInput{
		// DomainName: *string, // Required
		// KeyName: *string, // Required
		// Values: []string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesKeyName) > 0 {
		input.KeyName = aws.String(_customerprofilesKeyName)
	}
	if len(_customerprofilesValues) > 0 {
		input.Values = append([]string(nil), _customerprofilesValues...)
	}
	if len(_customerprofilesAdditionalSearchKeys) > 0 {
		if err := assignInputField(input, "AdditionalSearchKeys", _customerprofilesAdditionalSearchKeys); err != nil {
			log.Errorf("invalid --additional-search-keys: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesLogicalOperator) > 0 {
		if err := assignInputField(input, "LogicalOperator", _customerprofilesLogicalOperator); err != nil {
			log.Errorf("invalid --logical-operator: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _customerprofilesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesNextToken) > 0 {
		input.NextToken = aws.String(_customerprofilesNextToken)
	}

	if resp, err := client.SearchProfiles(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a recommender that was previously stopped. Starting a recommender
// resumes its ability to generate recommendations.
func customerprofiles_StartRecommender(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.StartRecommenderInput{
		// DomainName: *string, // Required
		// RecommenderName: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesRecommenderName) > 0 {
		input.RecommenderName = aws.String(_customerprofilesRecommenderName)
	}

	if resp, err := client.StartRecommender(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This API starts the processing of an upload job to ingest profile data.
func customerprofiles_StartUploadJob(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.StartUploadJobInput{
		// DomainName: *string, // Required
		// JobId: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesJobId) > 0 {
		input.JobId = aws.String(_customerprofilesJobId)
	}

	if resp, err := client.StartUploadJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a recommender, suspending its ability to generate recommendations. The
// recommender can be restarted later using StartRecommender.
func customerprofiles_StopRecommender(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.StopRecommenderInput{
		// DomainName: *string, // Required
		// RecommenderName: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesRecommenderName) > 0 {
		input.RecommenderName = aws.String(_customerprofilesRecommenderName)
	}

	if resp, err := client.StopRecommender(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This API stops the processing of an upload job.
func customerprofiles_StopUploadJob(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.StopUploadJobInput{
		// DomainName: *string, // Required
		// JobId: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesJobId) > 0 {
		input.JobId = aws.String(_customerprofilesJobId)
	}

	if resp, err := client.StopUploadJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Assigns one or more tags (key-value pairs) to the specified Amazon Connect
// Customer Profiles resource. Tags can help you organize and categorize your
// resources. You can also use them to scope user permissions by granting a user
// permission to access or change only resources with certain tag values. In
// Connect Customer Profiles, domains, profile object types, and integrations can
// be tagged.
//
// Tags don't have any semantic meaning to AWS and are interpreted strictly as
// strings of characters.
//
// You can use the TagResource action with a resource that already has tags. If
// you specify a new tag key, this tag is appended to the list of tags associated
// with the resource. If you specify a tag key that is already associated with the
// resource, the new tag value that you specify replaces the previous value for
// that tag.
//
// You can associate as many as 50 tags with a resource.
func customerprofiles_TagResource(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_customerprofilesResourceArn) > 0 {
		input.ResourceArn = aws.String(_customerprofilesResourceArn)
	}
	if len(_customerprofilesTags) > 0 {
		if err := assignInputField(input, "Tags", _customerprofilesTags); err != nil {
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

// Removes one or more tags from the specified Amazon Connect Customer Profiles
// resource. In Connect Customer Profiles, domains, profile object types, and
// integrations can be tagged.
func customerprofiles_UntagResource(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_customerprofilesResourceArn) > 0 {
		input.ResourceArn = aws.String(_customerprofilesResourceArn)
	}
	if len(_customerprofilesTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _customerprofilesTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing calculated attribute definition. When updating the
// Conditions, note that increasing the date range of a calculated attribute will
// not trigger inclusion of historical data greater than the current date range.
func customerprofiles_UpdateCalculatedAttributeDefinition(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.UpdateCalculatedAttributeDefinitionInput{
		// CalculatedAttributeName: *string, // Required
		// DomainName: *string, // Required
	}

	if len(_customerprofilesCalculatedAttributeName) > 0 {
		input.CalculatedAttributeName = aws.String(_customerprofilesCalculatedAttributeName)
	}
	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesConditions) > 0 {
		if err := assignInputField(input, "Conditions", _customerprofilesConditions); err != nil {
			log.Errorf("invalid --conditions: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesDescription) > 0 {
		input.Description = aws.String(_customerprofilesDescription)
	}
	if len(_customerprofilesDisplayName) > 0 {
		input.DisplayName = aws.String(_customerprofilesDisplayName)
	}

	if resp, err := client.UpdateCalculatedAttributeDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the properties of a domain, including creating or selecting a dead
// letter queue or an encryption key.
//
// After a domain is created, the name can’t be changed.
//
// Use this API or [CreateDomain] to enable [identity resolution]: set Matching to true.
//
// To prevent cross-service impersonation when you call this API, see [Cross-service confused deputy prevention] for sample
// policies that you should apply.
//
// To add or remove tags on an existing Domain, see [TagResource]/[UntagResource] .
//
// [CreateDomain]: https://docs.aws.amazon.com/customerprofiles/latest/APIReference/API_CreateDomain.html
// [TagResource]: https://docs.aws.amazon.com/customerprofiles/latest/APIReference/API_TagResource.html
// [Cross-service confused deputy prevention]: https://docs.aws.amazon.com/connect/latest/adminguide/cross-service-confused-deputy-prevention.html
// [UntagResource]: https://docs.aws.amazon.com/customerprofiles/latest/APIReference/API_UntagResource.html
// [identity resolution]: https://docs.aws.amazon.com/customerprofiles/latest/APIReference/API_GetMatches.html
func customerprofiles_UpdateDomain(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.UpdateDomainInput{
		// DomainName: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesDataStore) > 0 {
		if err := assignInputField(input, "DataStore", _customerprofilesDataStore); err != nil {
			log.Errorf("invalid --data-store: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesDeadLetterQueueUrl) > 0 {
		input.DeadLetterQueueUrl = aws.String(_customerprofilesDeadLetterQueueUrl)
	}
	if len(_customerprofilesDefaultEncryptionKey) > 0 {
		input.DefaultEncryptionKey = aws.String(_customerprofilesDefaultEncryptionKey)
	}
	if len(_customerprofilesDefaultExpirationDays) > 0 {
		if err := assignInputField(input, "DefaultExpirationDays", _customerprofilesDefaultExpirationDays); err != nil {
			log.Errorf("invalid --default-expiration-days: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesMatching) > 0 {
		if err := assignInputField(input, "Matching", _customerprofilesMatching); err != nil {
			log.Errorf("invalid --matching: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesRuleBasedMatching) > 0 {
		if err := assignInputField(input, "RuleBasedMatching", _customerprofilesRuleBasedMatching); err != nil {
			log.Errorf("invalid --rule-based-matching: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesTags) > 0 {
		if err := assignInputField(input, "Tags", _customerprofilesTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the layout used to view data for a specific domain. This API can only
// be invoked from the Amazon Connect admin website.
func customerprofiles_UpdateDomainLayout(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.UpdateDomainLayoutInput{
		// DomainName: *string, // Required
		// LayoutDefinitionName: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesLayoutDefinitionName) > 0 {
		input.LayoutDefinitionName = aws.String(_customerprofilesLayoutDefinitionName)
	}
	if len(_customerprofilesDescription) > 0 {
		input.Description = aws.String(_customerprofilesDescription)
	}
	if len(_customerprofilesDisplayName) > 0 {
		input.DisplayName = aws.String(_customerprofilesDisplayName)
	}
	if len(_customerprofilesIsDefault) > 0 {
		if err := assignInputField(input, "IsDefault", _customerprofilesIsDefault); err != nil {
			log.Errorf("invalid --is-default: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesLayout) > 0 {
		input.Layout = aws.String(_customerprofilesLayout)
	}
	if len(_customerprofilesLayoutType) > 0 {
		if err := assignInputField(input, "LayoutType", _customerprofilesLayoutType); err != nil {
			log.Errorf("invalid --layout-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateDomainLayout(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update the properties of an Event Trigger.
func customerprofiles_UpdateEventTrigger(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.UpdateEventTriggerInput{
		// DomainName: *string, // Required
		// EventTriggerName: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesEventTriggerName) > 0 {
		input.EventTriggerName = aws.String(_customerprofilesEventTriggerName)
	}
	if len(_customerprofilesDescription) > 0 {
		input.Description = aws.String(_customerprofilesDescription)
	}
	if len(_customerprofilesEventTriggerConditions) > 0 {
		if err := assignInputField(input, "EventTriggerConditions", _customerprofilesEventTriggerConditions); err != nil {
			log.Errorf("invalid --event-trigger-conditions: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesEventTriggerLimits) > 0 {
		if err := assignInputField(input, "EventTriggerLimits", _customerprofilesEventTriggerLimits); err != nil {
			log.Errorf("invalid --event-trigger-limits: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesObjectTypeName) > 0 {
		input.ObjectTypeName = aws.String(_customerprofilesObjectTypeName)
	}
	if len(_customerprofilesSegmentFilter) > 0 {
		input.SegmentFilter = aws.String(_customerprofilesSegmentFilter)
	}

	if resp, err := client.UpdateEventTrigger(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the properties of a profile. The ProfileId is required for updating a
// customer profile.
//
// When calling the UpdateProfile API, specifying an empty string value means that
// any existing value will be removed. Not specifying a string value means that any
// value already there will be kept.
func customerprofiles_UpdateProfile(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.UpdateProfileInput{
		// DomainName: *string, // Required
		// ProfileId: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesProfileId) > 0 {
		input.ProfileId = aws.String(_customerprofilesProfileId)
	}
	if len(_customerprofilesAccountNumber) > 0 {
		input.AccountNumber = aws.String(_customerprofilesAccountNumber)
	}
	if len(_customerprofilesAdditionalInformation) > 0 {
		input.AdditionalInformation = aws.String(_customerprofilesAdditionalInformation)
	}
	if len(_customerprofilesAddress) > 0 {
		if err := assignInputField(input, "Address", _customerprofilesAddress); err != nil {
			log.Errorf("invalid --address: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesAttributes) > 0 {
		if err := assignInputField(input, "Attributes", _customerprofilesAttributes); err != nil {
			log.Errorf("invalid --attributes: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesBillingAddress) > 0 {
		if err := assignInputField(input, "BillingAddress", _customerprofilesBillingAddress); err != nil {
			log.Errorf("invalid --billing-address: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesBirthDate) > 0 {
		input.BirthDate = aws.String(_customerprofilesBirthDate)
	}
	if len(_customerprofilesBusinessEmailAddress) > 0 {
		input.BusinessEmailAddress = aws.String(_customerprofilesBusinessEmailAddress)
	}
	if len(_customerprofilesBusinessName) > 0 {
		input.BusinessName = aws.String(_customerprofilesBusinessName)
	}
	if len(_customerprofilesBusinessPhoneNumber) > 0 {
		input.BusinessPhoneNumber = aws.String(_customerprofilesBusinessPhoneNumber)
	}
	if len(_customerprofilesEmailAddress) > 0 {
		input.EmailAddress = aws.String(_customerprofilesEmailAddress)
	}
	if len(_customerprofilesEngagementPreferences) > 0 {
		if err := assignInputField(input, "EngagementPreferences", _customerprofilesEngagementPreferences); err != nil {
			log.Errorf("invalid --engagement-preferences: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesFirstName) > 0 {
		input.FirstName = aws.String(_customerprofilesFirstName)
	}
	if len(_customerprofilesGender) > 0 {
		if err := assignInputField(input, "Gender", _customerprofilesGender); err != nil {
			log.Errorf("invalid --gender: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesGenderString) > 0 {
		input.GenderString = aws.String(_customerprofilesGenderString)
	}
	if len(_customerprofilesHomePhoneNumber) > 0 {
		input.HomePhoneNumber = aws.String(_customerprofilesHomePhoneNumber)
	}
	if len(_customerprofilesLastName) > 0 {
		input.LastName = aws.String(_customerprofilesLastName)
	}
	if len(_customerprofilesMailingAddress) > 0 {
		if err := assignInputField(input, "MailingAddress", _customerprofilesMailingAddress); err != nil {
			log.Errorf("invalid --mailing-address: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesMiddleName) > 0 {
		input.MiddleName = aws.String(_customerprofilesMiddleName)
	}
	if len(_customerprofilesMobilePhoneNumber) > 0 {
		input.MobilePhoneNumber = aws.String(_customerprofilesMobilePhoneNumber)
	}
	if len(_customerprofilesPartyType) > 0 {
		if err := assignInputField(input, "PartyType", _customerprofilesPartyType); err != nil {
			log.Errorf("invalid --party-type: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesPartyTypeString) > 0 {
		input.PartyTypeString = aws.String(_customerprofilesPartyTypeString)
	}
	if len(_customerprofilesPersonalEmailAddress) > 0 {
		input.PersonalEmailAddress = aws.String(_customerprofilesPersonalEmailAddress)
	}
	if len(_customerprofilesPhoneNumber) > 0 {
		input.PhoneNumber = aws.String(_customerprofilesPhoneNumber)
	}
	if len(_customerprofilesProfileType) > 0 {
		if err := assignInputField(input, "ProfileType", _customerprofilesProfileType); err != nil {
			log.Errorf("invalid --profile-type: %s", err.Error())
			return
		}
	}
	if len(_customerprofilesShippingAddress) > 0 {
		if err := assignInputField(input, "ShippingAddress", _customerprofilesShippingAddress); err != nil {
			log.Errorf("invalid --shipping-address: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the properties of an existing recommender, allowing you to modify its
// configuration and description.
func customerprofiles_UpdateRecommender(cfg aws.Config, client *customerprofiles.Client) {
	input := &customerprofiles.UpdateRecommenderInput{
		// DomainName: *string, // Required
		// RecommenderName: *string, // Required
	}

	if len(_customerprofilesDomainName) > 0 {
		input.DomainName = aws.String(_customerprofilesDomainName)
	}
	if len(_customerprofilesRecommenderName) > 0 {
		input.RecommenderName = aws.String(_customerprofilesRecommenderName)
	}
	if len(_customerprofilesDescription) > 0 {
		input.Description = aws.String(_customerprofilesDescription)
	}
	if len(_customerprofilesRecommenderConfig) > 0 {
		if err := assignInputField(input, "RecommenderConfig", _customerprofilesRecommenderConfig); err != nil {
			log.Errorf("invalid --recommender-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateRecommender(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_customerprofilesCmd)
	_customerprofilesCmd.Flags().SortFlags = false

	_customerprofilesCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_customerprofilesCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_customerprofilesCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesAccountNumber, "account-number", "", "", "Account Number")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesActionType, "action-type", "", "", "Action Type")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesAdditionalInformation, "additional-information", "", "", "Additional Information")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesAdditionalSearchKeys, "additional-search-keys", "", "", "Additional Search Keys")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesAddress, "address", "", "", "Address")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesAllowProfileCreation, "allow-profile-creation", "", "", "Allow Profile Creation")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesAttributeDetails, "attribute-details", "", "", "Attribute Details")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesAttributeName, "attribute-name", "", "", "Attribute Name")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesAttributes, "attributes", "", "", "Attributes")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesBillingAddress, "billing-address", "", "", "Billing Address")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesBirthDate, "birth-date", "", "", "Birth Date")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesBusinessEmailAddress, "business-email-address", "", "", "Business Email Address")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesBusinessName, "business-name", "", "", "Business Name")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesBusinessPhoneNumber, "business-phone-number", "", "", "Business Phone Number")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesCalculatedAttributeName, "calculated-attribute-name", "", "", "Calculated Attribute Name")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesConditionOverrides, "condition-overrides", "", "", "Condition Overrides")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesConditions, "conditions", "", "", "Conditions")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesConflictResolution, "conflict-resolution", "", "", "Conflict Resolution")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesConsolidation, "consolidation", "", "", "Consolidation")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesContext, "context", "", "", "Context")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesDataExpiry, "data-expiry", "", "", "Data Expiry")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesDataFormat, "data-format", "", "", "Data Format")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesDataStore, "data-store", "", "", "Data Store")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesDeadLetterQueueUrl, "dead-letter-queue-url", "", "", "Dead Letter Queue URL")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesDefaultEncryptionKey, "default-encryption-key", "", "", "Default Encryption Key")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesDefaultExpirationDays, "default-expiration-days", "", "", "Default Expiration Days")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesDescription, "description", "", "", "Description")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesDestinationUri, "destination-uri", "", "", "Destination URI")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesDisplayName, "display-name", "", "", "Display Name")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesDomainName, "domain-name", "", "", "Domain Name")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesEmailAddress, "email-address", "", "", "Email Address")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesEncryptionKey, "encryption-key", "", "", "Encryption Key")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesEngagementPreferences, "engagement-preferences", "", "", "Engagement Preferences")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesEstimateId, "estimate-id", "", "", "Estimate ID")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesEventStreamName, "event-stream-name", "", "", "Event Stream Name")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesEventTriggerConditions, "event-trigger-conditions", "", "", "Event Trigger Conditions")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesEventTriggerLimits, "event-trigger-limits", "", "", "Event Trigger Limits")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesEventTriggerName, "event-trigger-name", "", "", "Event Trigger Name")
	_customerprofilesCmd.Flags().StringSliceVarP(&_customerprofilesEventTriggerNames, "event-trigger-names", "", nil, "Event Trigger Names")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesExpirationDays, "expiration-days", "", "", "Expiration Days")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesFieldSourceProfileIds, "field-source-profile-ids", "", "", "Field Source Profile Ids")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesFields, "fields", "", "", "Fields")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesFilter, "filter", "", "", "Filter")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesFirstName, "first-name", "", "", "First Name")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesFlowDefinition, "flow-definition", "", "", "Flow Definition")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesGender, "gender", "", "", "Gender")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesGenderString, "gender-string", "", "", "Gender String")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesHomePhoneNumber, "home-phone-number", "", "", "Home Phone Number")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesId, "id", "", "", "ID")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesIncludeHidden, "include-hidden", "", "", "Include Hidden")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesIntegrationConfig, "integration-config", "", "", "Integration Config")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesIsDefault, "is-default", "", "", "Is Default")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesJobId, "job-id", "", "", "Job ID")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesKeyName, "key-name", "", "", "Key Name")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesKeys, "keys", "", "", "Keys")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesLastName, "last-name", "", "", "Last Name")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesLayout, "layout", "", "", "Layout")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesLayoutDefinitionName, "layout-definition-name", "", "", "Layout Definition Name")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesLayoutType, "layout-type", "", "", "Layout Type")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesLogicalOperator, "logical-operator", "", "", "Logical Operator")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesMailingAddress, "mailing-address", "", "", "Mailing Address")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesMainProfileId, "main-profile-id", "", "", "Main Profile ID")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesMatchType, "match-type", "", "", "Match Type")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesMatching, "matching", "", "", "Matching")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesMaxProfileObjectCount, "max-profile-object-count", "", "", "Max Profile Object Count")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesMaxResults, "max-results", "", "", "Max Results")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesMiddleName, "middle-name", "", "", "Middle Name")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesMinAllowedConfidenceScoreForMerging, "min-allowed-confidence-score-for-merging", "", "", "Min Allowed Confidence Score For Merging")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesMobilePhoneNumber, "mobile-phone-number", "", "", "Mobile Phone Number")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesNextToken, "next-token", "", "", "Next Token")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesObject, "object", "", "", "Object")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesObjectFilter, "object-filter", "", "", "Object Filter")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesObjectTypeName, "object-type-name", "", "", "Object Type Name")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesObjectTypeNames, "object-type-names", "", "", "Object Type Names")
	_customerprofilesCmd.Flags().StringSliceVarP(&_customerprofilesObjects, "objects", "", nil, "Objects")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesPartyType, "party-type", "", "", "Party Type")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesPartyTypeString, "party-type-string", "", "", "Party Type String")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesPerformedBy, "performed-by", "", "", "Performed By")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesPersonalEmailAddress, "personal-email-address", "", "", "Personal Email Address")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesPhoneNumber, "phone-number", "", "", "Phone Number")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesProfileId, "profile-id", "", "", "Profile ID")
	_customerprofilesCmd.Flags().StringSliceVarP(&_customerprofilesProfileIds, "profile-ids", "", nil, "Profile Ids")
	_customerprofilesCmd.Flags().StringSliceVarP(&_customerprofilesProfileIdsToBeMerged, "profile-ids-to-be-merged", "", nil, "Profile Ids To Be Merged")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesProfileObjectUniqueKey, "profile-object-unique-key", "", "", "Profile Object Unique Key")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesProfileType, "profile-type", "", "", "Profile Type")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesQueryEndDate, "query-end-date", "", "", "Query End Date")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesQueryStartDate, "query-start-date", "", "", "Query Start Date")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesRecommenderConfig, "recommender-config", "", "", "Recommender Config")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesRecommenderName, "recommender-name", "", "", "Recommender Name")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesRecommenderRecipeName, "recommender-recipe-name", "", "", "Recommender Recipe Name")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesResourceArn, "resource-arn", "", "", "Resource ARN")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesRoleArn, "role-arn", "", "", "Role ARN")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesRuleBasedMatching, "rule-based-matching", "", "", "Rule Based Matching")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesScope, "scope", "", "", "Scope")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesSearchKey, "search-key", "", "", "Search Key")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesSearchValue, "search-value", "", "", "Search Value")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesSegmentDefinitionName, "segment-definition-name", "", "", "Segment Definition Name")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesSegmentFilter, "segment-filter", "", "", "Segment Filter")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesSegmentGroups, "segment-groups", "", "", "Segment Groups")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesSegmentQuery, "segment-query", "", "", "Segment Query")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesSegmentSqlQuery, "segment-sql-query", "", "", "Segment Sql Query")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesShippingAddress, "shipping-address", "", "", "Shipping Address")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesSnapshotId, "snapshot-id", "", "", "Snapshot ID")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesSourceLastUpdatedTimestampFormat, "source-last-updated-timestamp-format", "", "", "Source Last Updated Timestamp Format")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesSourcePriority, "source-priority", "", "", "Source Priority")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesStatistic, "statistic", "", "", "Statistic")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesStatus, "status", "", "", "Status")
	_customerprofilesCmd.Flags().StringSliceVarP(&_customerprofilesTagKeys, "tag-keys", "", nil, "Tag Keys")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesTags, "tags", "", "", "Tags")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesTemplateId, "template-id", "", "", "Template ID")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesTrainingMetricsCount, "training-metrics-count", "", "", "Training Metrics Count")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesUniqueKey, "unique-key", "", "", "Unique Key")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesUri, "uri", "", "", "URI")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesUseHistoricalData, "use-historical-data", "", "", "Use Historical Data")
	_customerprofilesCmd.Flags().StringSliceVarP(&_customerprofilesValues, "values", "", nil, "Values")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesWorkflowId, "workflow-id", "", "", "Workflow ID")
	_customerprofilesCmd.Flags().StringVarP(&_customerprofilesWorkflowType, "workflow-type", "", "", "Workflow Type")

	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesAddProfileKey, "add-profile-key", "", false, "Add Profile Key")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesBatchGetCalculatedAttributeForProfile, "batch-get-calculated-attribute-for-profile", "", false, "Batch Get Calculated Attribute For Profile")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesBatchGetProfile, "batch-get-profile", "", false, "Batch Get Profile")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesCreateCalculatedAttributeDefinition, "create-calculated-attribute-definition", "", false, "Create Calculated Attribute Definition")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesCreateDomain, "create-domain", "", false, "Create Domain")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesCreateDomainLayout, "create-domain-layout", "", false, "Create Domain Layout")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesCreateEventStream, "create-event-stream", "", false, "Create Event Stream")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesCreateEventTrigger, "create-event-trigger", "", false, "Create Event Trigger")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesCreateIntegrationWorkflow, "create-integration-workflow", "", false, "Create Integration Workflow")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesCreateProfile, "create-profile", "", false, "Create Profile")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesCreateRecommender, "create-recommender", "", false, "Create Recommender")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesCreateSegmentDefinition, "create-segment-definition", "", false, "Create Segment Definition")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesCreateSegmentEstimate, "create-segment-estimate", "", false, "Create Segment Estimate")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesCreateSegmentSnapshot, "create-segment-snapshot", "", false, "Create Segment Snapshot")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesCreateUploadJob, "create-upload-job", "", false, "Create Upload Job")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesDeleteCalculatedAttributeDefinition, "delete-calculated-attribute-definition", "", false, "Delete Calculated Attribute Definition")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesDeleteDomain, "delete-domain", "", false, "Delete Domain")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesDeleteDomainLayout, "delete-domain-layout", "", false, "Delete Domain Layout")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesDeleteDomainObjectType, "delete-domain-object-type", "", false, "Delete Domain Object Type")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesDeleteEventStream, "delete-event-stream", "", false, "Delete Event Stream")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesDeleteEventTrigger, "delete-event-trigger", "", false, "Delete Event Trigger")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesDeleteIntegration, "delete-integration", "", false, "Delete Integration")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesDeleteProfile, "delete-profile", "", false, "Delete Profile")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesDeleteProfileKey, "delete-profile-key", "", false, "Delete Profile Key")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesDeleteProfileObject, "delete-profile-object", "", false, "Delete Profile Object")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesDeleteProfileObjectType, "delete-profile-object-type", "", false, "Delete Profile Object Type")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesDeleteRecommender, "delete-recommender", "", false, "Delete Recommender")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesDeleteSegmentDefinition, "delete-segment-definition", "", false, "Delete Segment Definition")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesDeleteWorkflow, "delete-workflow", "", false, "Delete Workflow")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesDetectProfileObjectType, "detect-profile-object-type", "", false, "Detect Profile Object Type")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesGetAutoMergingPreview, "get-auto-merging-preview", "", false, "Get Auto Merging Preview")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesGetCalculatedAttributeDefinition, "get-calculated-attribute-definition", "", false, "Get Calculated Attribute Definition")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesGetCalculatedAttributeForProfile, "get-calculated-attribute-for-profile", "", false, "Get Calculated Attribute For Profile")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesGetDomain, "get-domain", "", false, "Get Domain")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesGetDomainLayout, "get-domain-layout", "", false, "Get Domain Layout")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesGetDomainObjectType, "get-domain-object-type", "", false, "Get Domain Object Type")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesGetEventStream, "get-event-stream", "", false, "Get Event Stream")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesGetEventTrigger, "get-event-trigger", "", false, "Get Event Trigger")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesGetIdentityResolutionJob, "get-identity-resolution-job", "", false, "Get Identity Resolution Job")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesGetIntegration, "get-integration", "", false, "Get Integration")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesGetMatches, "get-matches", "", false, "Get Matches")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesGetObjectTypeAttributeStatistics, "get-object-type-attribute-statistics", "", false, "Get Object Type Attribute Statistics")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesGetProfileHistoryRecord, "get-profile-history-record", "", false, "Get Profile History Record")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesGetProfileObjectType, "get-profile-object-type", "", false, "Get Profile Object Type")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesGetProfileObjectTypeTemplate, "get-profile-object-type-template", "", false, "Get Profile Object Type Template")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesGetProfileRecommendations, "get-profile-recommendations", "", false, "Get Profile Recommendations")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesGetRecommender, "get-recommender", "", false, "Get Recommender")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesGetSegmentDefinition, "get-segment-definition", "", false, "Get Segment Definition")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesGetSegmentEstimate, "get-segment-estimate", "", false, "Get Segment Estimate")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesGetSegmentMembership, "get-segment-membership", "", false, "Get Segment Membership")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesGetSegmentSnapshot, "get-segment-snapshot", "", false, "Get Segment Snapshot")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesGetSimilarProfiles, "get-similar-profiles", "", false, "Get Similar Profiles")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesGetUploadJob, "get-upload-job", "", false, "Get Upload Job")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesGetUploadJobPath, "get-upload-job-path", "", false, "Get Upload Job Path")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesGetWorkflow, "get-workflow", "", false, "Get Workflow")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesGetWorkflowSteps, "get-workflow-steps", "", false, "Get Workflow Steps")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesListAccountIntegrations, "list-account-integrations", "", false, "List Account Integrations")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesListCalculatedAttributeDefinitions, "list-calculated-attribute-definitions", "", false, "List Calculated Attribute Definitions")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesListCalculatedAttributesForProfile, "list-calculated-attributes-for-profile", "", false, "List Calculated Attributes For Profile")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesListDomainLayouts, "list-domain-layouts", "", false, "List Domain Layouts")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesListDomainObjectTypes, "list-domain-object-types", "", false, "List Domain Object Types")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesListDomains, "list-domains", "", false, "List Domains")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesListEventStreams, "list-event-streams", "", false, "List Event Streams")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesListEventTriggers, "list-event-triggers", "", false, "List Event Triggers")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesListIdentityResolutionJobs, "list-identity-resolution-jobs", "", false, "List Identity Resolution Jobs")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesListIntegrations, "list-integrations", "", false, "List Integrations")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesListObjectTypeAttributeValues, "list-object-type-attribute-values", "", false, "List Object Type Attribute Values")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesListObjectTypeAttributes, "list-object-type-attributes", "", false, "List Object Type Attributes")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesListProfileAttributeValues, "list-profile-attribute-values", "", false, "List Profile Attribute Values")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesListProfileHistoryRecords, "list-profile-history-records", "", false, "List Profile History Records")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesListProfileObjectTypeTemplates, "list-profile-object-type-templates", "", false, "List Profile Object Type Templates")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesListProfileObjectTypes, "list-profile-object-types", "", false, "List Profile Object Types")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesListProfileObjects, "list-profile-objects", "", false, "List Profile Objects")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesListRecommenderRecipes, "list-recommender-recipes", "", false, "List Recommender Recipes")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesListRecommenders, "list-recommenders", "", false, "List Recommenders")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesListRuleBasedMatches, "list-rule-based-matches", "", false, "List Rule Based Matches")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesListSegmentDefinitions, "list-segment-definitions", "", false, "List Segment Definitions")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesListUploadJobs, "list-upload-jobs", "", false, "List Upload Jobs")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesListWorkflows, "list-workflows", "", false, "List Workflows")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesMergeProfiles, "merge-profiles", "", false, "Merge Profiles")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesPutDomainObjectType, "put-domain-object-type", "", false, "Put Domain Object Type")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesPutIntegration, "put-integration", "", false, "Put Integration")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesPutProfileObject, "put-profile-object", "", false, "Put Profile Object")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesPutProfileObjectType, "put-profile-object-type", "", false, "Put Profile Object Type")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesSearchProfiles, "search-profiles", "", false, "Search Profiles")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesStartRecommender, "start-recommender", "", false, "Start Recommender")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesStartUploadJob, "start-upload-job", "", false, "Start Upload Job")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesStopRecommender, "stop-recommender", "", false, "Stop Recommender")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesStopUploadJob, "stop-upload-job", "", false, "Stop Upload Job")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesTagResource, "tag-resource", "", false, "Tag Resource")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesUntagResource, "untag-resource", "", false, "Untag Resource")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesUpdateCalculatedAttributeDefinition, "update-calculated-attribute-definition", "", false, "Update Calculated Attribute Definition")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesUpdateDomain, "update-domain", "", false, "Update Domain")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesUpdateDomainLayout, "update-domain-layout", "", false, "Update Domain Layout")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesUpdateEventTrigger, "update-event-trigger", "", false, "Update Event Trigger")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesUpdateProfile, "update-profile", "", false, "Update Profile")
	_customerprofilesCmd.Flags().BoolVarP(&_customerprofilesUpdateRecommender, "update-recommender", "", false, "Update Recommender")

}
