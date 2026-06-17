package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/configservice"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// configserviceCmd represents the configservice command
var _configserviceCmd = &cobra.Command{
	Use:   "configservice",
	Short: "AWS configservice CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := configservice.NewFromConfig(cfg)
		if _configserviceAssociateResourceTypes {
			configservice_AssociateResourceTypes(cfg, client)
			return
		}
		if _configserviceBatchGetAggregateResourceConfig {
			configservice_BatchGetAggregateResourceConfig(cfg, client)
			return
		}
		if _configserviceBatchGetResourceConfig {
			configservice_BatchGetResourceConfig(cfg, client)
			return
		}
		if _configserviceDeleteAggregationAuthorization {
			configservice_DeleteAggregationAuthorization(cfg, client)
			return
		}
		if _configserviceDeleteConfigRule {
			configservice_DeleteConfigRule(cfg, client)
			return
		}
		if _configserviceDeleteConfigurationAggregator {
			configservice_DeleteConfigurationAggregator(cfg, client)
			return
		}
		if _configserviceDeleteConfigurationRecorder {
			configservice_DeleteConfigurationRecorder(cfg, client)
			return
		}
		if _configserviceDeleteConformancePack {
			configservice_DeleteConformancePack(cfg, client)
			return
		}
		if _configserviceDeleteDeliveryChannel {
			configservice_DeleteDeliveryChannel(cfg, client)
			return
		}
		if _configserviceDeleteEvaluationResults {
			configservice_DeleteEvaluationResults(cfg, client)
			return
		}
		if _configserviceDeleteOrganizationConfigRule {
			configservice_DeleteOrganizationConfigRule(cfg, client)
			return
		}
		if _configserviceDeleteOrganizationConformancePack {
			configservice_DeleteOrganizationConformancePack(cfg, client)
			return
		}
		if _configserviceDeletePendingAggregationRequest {
			configservice_DeletePendingAggregationRequest(cfg, client)
			return
		}
		if _configserviceDeleteRemediationConfiguration {
			configservice_DeleteRemediationConfiguration(cfg, client)
			return
		}
		if _configserviceDeleteRemediationExceptions {
			configservice_DeleteRemediationExceptions(cfg, client)
			return
		}
		if _configserviceDeleteResourceConfig {
			configservice_DeleteResourceConfig(cfg, client)
			return
		}
		if _configserviceDeleteRetentionConfiguration {
			configservice_DeleteRetentionConfiguration(cfg, client)
			return
		}
		if _configserviceDeleteServiceLinkedConfigurationRecorder {
			configservice_DeleteServiceLinkedConfigurationRecorder(cfg, client)
			return
		}
		if _configserviceDeleteStoredQuery {
			configservice_DeleteStoredQuery(cfg, client)
			return
		}
		if _configserviceDeliverConfigSnapshot {
			configservice_DeliverConfigSnapshot(cfg, client)
			return
		}
		if _configserviceDescribeAggregateComplianceByConfigRules {
			configservice_DescribeAggregateComplianceByConfigRules(cfg, client)
			return
		}
		if _configserviceDescribeAggregateComplianceByConformancePacks {
			configservice_DescribeAggregateComplianceByConformancePacks(cfg, client)
			return
		}
		if _configserviceDescribeAggregationAuthorizations {
			configservice_DescribeAggregationAuthorizations(cfg, client)
			return
		}
		if _configserviceDescribeComplianceByConfigRule {
			configservice_DescribeComplianceByConfigRule(cfg, client)
			return
		}
		if _configserviceDescribeComplianceByResource {
			configservice_DescribeComplianceByResource(cfg, client)
			return
		}
		if _configserviceDescribeConfigRuleEvaluationStatus {
			configservice_DescribeConfigRuleEvaluationStatus(cfg, client)
			return
		}
		if _configserviceDescribeConfigRules {
			configservice_DescribeConfigRules(cfg, client)
			return
		}
		if _configserviceDescribeConfigurationAggregatorSourcesStatus {
			configservice_DescribeConfigurationAggregatorSourcesStatus(cfg, client)
			return
		}
		if _configserviceDescribeConfigurationAggregators {
			configservice_DescribeConfigurationAggregators(cfg, client)
			return
		}
		if _configserviceDescribeConfigurationRecorderStatus {
			configservice_DescribeConfigurationRecorderStatus(cfg, client)
			return
		}
		if _configserviceDescribeConfigurationRecorders {
			configservice_DescribeConfigurationRecorders(cfg, client)
			return
		}
		if _configserviceDescribeConformancePackCompliance {
			configservice_DescribeConformancePackCompliance(cfg, client)
			return
		}
		if _configserviceDescribeConformancePackStatus {
			configservice_DescribeConformancePackStatus(cfg, client)
			return
		}
		if _configserviceDescribeConformancePacks {
			configservice_DescribeConformancePacks(cfg, client)
			return
		}
		if _configserviceDescribeDeliveryChannelStatus {
			configservice_DescribeDeliveryChannelStatus(cfg, client)
			return
		}
		if _configserviceDescribeDeliveryChannels {
			configservice_DescribeDeliveryChannels(cfg, client)
			return
		}
		if _configserviceDescribeOrganizationConfigRuleStatuses {
			configservice_DescribeOrganizationConfigRuleStatuses(cfg, client)
			return
		}
		if _configserviceDescribeOrganizationConfigRules {
			configservice_DescribeOrganizationConfigRules(cfg, client)
			return
		}
		if _configserviceDescribeOrganizationConformancePackStatuses {
			configservice_DescribeOrganizationConformancePackStatuses(cfg, client)
			return
		}
		if _configserviceDescribeOrganizationConformancePacks {
			configservice_DescribeOrganizationConformancePacks(cfg, client)
			return
		}
		if _configserviceDescribePendingAggregationRequests {
			configservice_DescribePendingAggregationRequests(cfg, client)
			return
		}
		if _configserviceDescribeRemediationConfigurations {
			configservice_DescribeRemediationConfigurations(cfg, client)
			return
		}
		if _configserviceDescribeRemediationExceptions {
			configservice_DescribeRemediationExceptions(cfg, client)
			return
		}
		if _configserviceDescribeRemediationExecutionStatus {
			configservice_DescribeRemediationExecutionStatus(cfg, client)
			return
		}
		if _configserviceDescribeRetentionConfigurations {
			configservice_DescribeRetentionConfigurations(cfg, client)
			return
		}
		if _configserviceDisassociateResourceTypes {
			configservice_DisassociateResourceTypes(cfg, client)
			return
		}
		if _configserviceGetAggregateComplianceDetailsByConfigRule {
			configservice_GetAggregateComplianceDetailsByConfigRule(cfg, client)
			return
		}
		if _configserviceGetAggregateConfigRuleComplianceSummary {
			configservice_GetAggregateConfigRuleComplianceSummary(cfg, client)
			return
		}
		if _configserviceGetAggregateConformancePackComplianceSummary {
			configservice_GetAggregateConformancePackComplianceSummary(cfg, client)
			return
		}
		if _configserviceGetAggregateDiscoveredResourceCounts {
			configservice_GetAggregateDiscoveredResourceCounts(cfg, client)
			return
		}
		if _configserviceGetAggregateResourceConfig {
			configservice_GetAggregateResourceConfig(cfg, client)
			return
		}
		if _configserviceGetComplianceDetailsByConfigRule {
			configservice_GetComplianceDetailsByConfigRule(cfg, client)
			return
		}
		if _configserviceGetComplianceDetailsByResource {
			configservice_GetComplianceDetailsByResource(cfg, client)
			return
		}
		if _configserviceGetComplianceSummaryByConfigRule {
			configservice_GetComplianceSummaryByConfigRule(cfg, client)
			return
		}
		if _configserviceGetComplianceSummaryByResourceType {
			configservice_GetComplianceSummaryByResourceType(cfg, client)
			return
		}
		if _configserviceGetConformancePackComplianceDetails {
			configservice_GetConformancePackComplianceDetails(cfg, client)
			return
		}
		if _configserviceGetConformancePackComplianceSummary {
			configservice_GetConformancePackComplianceSummary(cfg, client)
			return
		}
		if _configserviceGetCustomRulePolicy {
			configservice_GetCustomRulePolicy(cfg, client)
			return
		}
		if _configserviceGetDiscoveredResourceCounts {
			configservice_GetDiscoveredResourceCounts(cfg, client)
			return
		}
		if _configserviceGetOrganizationConfigRuleDetailedStatus {
			configservice_GetOrganizationConfigRuleDetailedStatus(cfg, client)
			return
		}
		if _configserviceGetOrganizationConformancePackDetailedStatus {
			configservice_GetOrganizationConformancePackDetailedStatus(cfg, client)
			return
		}
		if _configserviceGetOrganizationCustomRulePolicy {
			configservice_GetOrganizationCustomRulePolicy(cfg, client)
			return
		}
		if _configserviceGetResourceConfigHistory {
			configservice_GetResourceConfigHistory(cfg, client)
			return
		}
		if _configserviceGetResourceEvaluationSummary {
			configservice_GetResourceEvaluationSummary(cfg, client)
			return
		}
		if _configserviceGetStoredQuery {
			configservice_GetStoredQuery(cfg, client)
			return
		}
		if _configserviceListAggregateDiscoveredResources {
			configservice_ListAggregateDiscoveredResources(cfg, client)
			return
		}
		if _configserviceListConfigurationRecorders {
			configservice_ListConfigurationRecorders(cfg, client)
			return
		}
		if _configserviceListConformancePackComplianceScores {
			configservice_ListConformancePackComplianceScores(cfg, client)
			return
		}
		if _configserviceListDiscoveredResources {
			configservice_ListDiscoveredResources(cfg, client)
			return
		}
		if _configserviceListResourceEvaluations {
			configservice_ListResourceEvaluations(cfg, client)
			return
		}
		if _configserviceListStoredQueries {
			configservice_ListStoredQueries(cfg, client)
			return
		}
		if _configserviceListTagsForResource {
			configservice_ListTagsForResource(cfg, client)
			return
		}
		if _configservicePutAggregationAuthorization {
			configservice_PutAggregationAuthorization(cfg, client)
			return
		}
		if _configservicePutConfigRule {
			configservice_PutConfigRule(cfg, client)
			return
		}
		if _configservicePutConfigurationAggregator {
			configservice_PutConfigurationAggregator(cfg, client)
			return
		}
		if _configservicePutConfigurationRecorder {
			configservice_PutConfigurationRecorder(cfg, client)
			return
		}
		if _configservicePutConformancePack {
			configservice_PutConformancePack(cfg, client)
			return
		}
		if _configservicePutDeliveryChannel {
			configservice_PutDeliveryChannel(cfg, client)
			return
		}
		if _configservicePutEvaluations {
			configservice_PutEvaluations(cfg, client)
			return
		}
		if _configservicePutExternalEvaluation {
			configservice_PutExternalEvaluation(cfg, client)
			return
		}
		if _configservicePutOrganizationConfigRule {
			configservice_PutOrganizationConfigRule(cfg, client)
			return
		}
		if _configservicePutOrganizationConformancePack {
			configservice_PutOrganizationConformancePack(cfg, client)
			return
		}
		if _configservicePutRemediationConfigurations {
			configservice_PutRemediationConfigurations(cfg, client)
			return
		}
		if _configservicePutRemediationExceptions {
			configservice_PutRemediationExceptions(cfg, client)
			return
		}
		if _configservicePutResourceConfig {
			configservice_PutResourceConfig(cfg, client)
			return
		}
		if _configservicePutRetentionConfiguration {
			configservice_PutRetentionConfiguration(cfg, client)
			return
		}
		if _configservicePutServiceLinkedConfigurationRecorder {
			configservice_PutServiceLinkedConfigurationRecorder(cfg, client)
			return
		}
		if _configservicePutStoredQuery {
			configservice_PutStoredQuery(cfg, client)
			return
		}
		if _configserviceSelectAggregateResourceConfig {
			configservice_SelectAggregateResourceConfig(cfg, client)
			return
		}
		if _configserviceSelectResourceConfig {
			configservice_SelectResourceConfig(cfg, client)
			return
		}
		if _configserviceStartConfigRulesEvaluation {
			configservice_StartConfigRulesEvaluation(cfg, client)
			return
		}
		if _configserviceStartConfigurationRecorder {
			configservice_StartConfigurationRecorder(cfg, client)
			return
		}
		if _configserviceStartRemediationExecution {
			configservice_StartRemediationExecution(cfg, client)
			return
		}
		if _configserviceStartResourceEvaluation {
			configservice_StartResourceEvaluation(cfg, client)
			return
		}
		if _configserviceStopConfigurationRecorder {
			configservice_StopConfigurationRecorder(cfg, client)
			return
		}
		if _configserviceTagResource {
			configservice_TagResource(cfg, client)
			return
		}
		if _configserviceUntagResource {
			configservice_UntagResource(cfg, client)
			return
		}

	},
}

var (
	_configserviceAssociateResourceTypes                        bool
	_configserviceBatchGetAggregateResourceConfig               bool
	_configserviceBatchGetResourceConfig                        bool
	_configserviceDeleteAggregationAuthorization                bool
	_configserviceDeleteConfigRule                              bool
	_configserviceDeleteConfigurationAggregator                 bool
	_configserviceDeleteConfigurationRecorder                   bool
	_configserviceDeleteConformancePack                         bool
	_configserviceDeleteDeliveryChannel                         bool
	_configserviceDeleteEvaluationResults                       bool
	_configserviceDeleteOrganizationConfigRule                  bool
	_configserviceDeleteOrganizationConformancePack             bool
	_configserviceDeletePendingAggregationRequest               bool
	_configserviceDeleteRemediationConfiguration                bool
	_configserviceDeleteRemediationExceptions                   bool
	_configserviceDeleteResourceConfig                          bool
	_configserviceDeleteRetentionConfiguration                  bool
	_configserviceDeleteServiceLinkedConfigurationRecorder      bool
	_configserviceDeleteStoredQuery                             bool
	_configserviceDeliverConfigSnapshot                         bool
	_configserviceDescribeAggregateComplianceByConfigRules      bool
	_configserviceDescribeAggregateComplianceByConformancePacks bool
	_configserviceDescribeAggregationAuthorizations             bool
	_configserviceDescribeComplianceByConfigRule                bool
	_configserviceDescribeComplianceByResource                  bool
	_configserviceDescribeConfigRuleEvaluationStatus            bool
	_configserviceDescribeConfigRules                           bool
	_configserviceDescribeConfigurationAggregatorSourcesStatus  bool
	_configserviceDescribeConfigurationAggregators              bool
	_configserviceDescribeConfigurationRecorderStatus           bool
	_configserviceDescribeConfigurationRecorders                bool
	_configserviceDescribeConformancePackCompliance             bool
	_configserviceDescribeConformancePackStatus                 bool
	_configserviceDescribeConformancePacks                      bool
	_configserviceDescribeDeliveryChannelStatus                 bool
	_configserviceDescribeDeliveryChannels                      bool
	_configserviceDescribeOrganizationConfigRuleStatuses        bool
	_configserviceDescribeOrganizationConfigRules               bool
	_configserviceDescribeOrganizationConformancePackStatuses   bool
	_configserviceDescribeOrganizationConformancePacks          bool
	_configserviceDescribePendingAggregationRequests            bool
	_configserviceDescribeRemediationConfigurations             bool
	_configserviceDescribeRemediationExceptions                 bool
	_configserviceDescribeRemediationExecutionStatus            bool
	_configserviceDescribeRetentionConfigurations               bool
	_configserviceDisassociateResourceTypes                     bool
	_configserviceGetAggregateComplianceDetailsByConfigRule     bool
	_configserviceGetAggregateConfigRuleComplianceSummary       bool
	_configserviceGetAggregateConformancePackComplianceSummary  bool
	_configserviceGetAggregateDiscoveredResourceCounts          bool
	_configserviceGetAggregateResourceConfig                    bool
	_configserviceGetComplianceDetailsByConfigRule              bool
	_configserviceGetComplianceDetailsByResource                bool
	_configserviceGetComplianceSummaryByConfigRule              bool
	_configserviceGetComplianceSummaryByResourceType            bool
	_configserviceGetConformancePackComplianceDetails           bool
	_configserviceGetConformancePackComplianceSummary           bool
	_configserviceGetCustomRulePolicy                           bool
	_configserviceGetDiscoveredResourceCounts                   bool
	_configserviceGetOrganizationConfigRuleDetailedStatus       bool
	_configserviceGetOrganizationConformancePackDetailedStatus  bool
	_configserviceGetOrganizationCustomRulePolicy               bool
	_configserviceGetResourceConfigHistory                      bool
	_configserviceGetResourceEvaluationSummary                  bool
	_configserviceGetStoredQuery                                bool
	_configserviceListAggregateDiscoveredResources              bool
	_configserviceListConfigurationRecorders                    bool
	_configserviceListConformancePackComplianceScores           bool
	_configserviceListDiscoveredResources                       bool
	_configserviceListResourceEvaluations                       bool
	_configserviceListStoredQueries                             bool
	_configserviceListTagsForResource                           bool
	_configservicePutAggregationAuthorization                   bool
	_configservicePutConfigRule                                 bool
	_configservicePutConfigurationAggregator                    bool
	_configservicePutConfigurationRecorder                      bool
	_configservicePutConformancePack                            bool
	_configservicePutDeliveryChannel                            bool
	_configservicePutEvaluations                                bool
	_configservicePutExternalEvaluation                         bool
	_configservicePutOrganizationConfigRule                     bool
	_configservicePutOrganizationConformancePack                bool
	_configservicePutRemediationConfigurations                  bool
	_configservicePutRemediationExceptions                      bool
	_configservicePutResourceConfig                             bool
	_configservicePutRetentionConfiguration                     bool
	_configservicePutServiceLinkedConfigurationRecorder         bool
	_configservicePutStoredQuery                                bool
	_configserviceSelectAggregateResourceConfig                 bool
	_configserviceSelectResourceConfig                          bool
	_configserviceStartConfigRulesEvaluation                    bool
	_configserviceStartConfigurationRecorder                    bool
	_configserviceStartRemediationExecution                     bool
	_configserviceStartResourceEvaluation                       bool
	_configserviceStopConfigurationRecorder                     bool
	_configserviceTagResource                                   bool
	_configserviceUntagResource                                 bool

	_configserviceAccountAggregationSources            string
	_configserviceAccountId                            string
	_configserviceAggregatorFilters                    string
	_configserviceArn                                  string
	_configserviceAuthorizedAccountId                  string
	_configserviceAuthorizedAwsRegion                  string
	_configserviceAwsRegion                            string
	_configserviceChronologicalOrder                   string
	_configserviceClientToken                          string
	_configserviceComplianceType                       string
	_configserviceComplianceTypes                      string
	_configserviceConfigRule                           string
	_configserviceConfigRuleName                       string
	_configserviceConfigRuleNames                      []string
	_configserviceConfiguration                        string
	_configserviceConfigurationAggregatorName          string
	_configserviceConfigurationAggregatorNames         []string
	_configserviceConfigurationRecorder                string
	_configserviceConfigurationRecorderArn             string
	_configserviceConfigurationRecorderName            string
	_configserviceConfigurationRecorderNames           []string
	_configserviceConformancePackInputParameters       string
	_configserviceConformancePackName                  string
	_configserviceConformancePackNames                 []string
	_configserviceDeliveryChannel                      string
	_configserviceDeliveryChannelName                  string
	_configserviceDeliveryChannelNames                 []string
	_configserviceDeliveryS3Bucket                     string
	_configserviceDeliveryS3KeyPrefix                  string
	_configserviceEarlierTime                          string
	_configserviceEvaluationContext                    string
	_configserviceEvaluationMode                       string
	_configserviceEvaluationTimeout                    string
	_configserviceEvaluations                          string
	_configserviceExcludedAccounts                     []string
	_configserviceExpirationTime                       string
	_configserviceExpression                           string
	_configserviceExternalEvaluation                   string
	_configserviceFilters                              string
	_configserviceGroupByKey                           string
	_configserviceIncludeDeletedResources              string
	_configserviceLaterTime                            string
	_configserviceLimit                                string
	_configserviceMaxResults                           string
	_configserviceMessage                              string
	_configserviceNextToken                            string
	_configserviceOrganizationAggregationSource        string
	_configserviceOrganizationConfigRuleName           string
	_configserviceOrganizationConfigRuleNames          []string
	_configserviceOrganizationConformancePackName      string
	_configserviceOrganizationConformancePackNames     []string
	_configserviceOrganizationCustomPolicyRuleMetadata string
	_configserviceOrganizationCustomRuleMetadata       string
	_configserviceOrganizationManagedRuleMetadata      string
	_configserviceQueryName                            string
	_configserviceRemediationConfigurations            string
	_configserviceRequesterAccountId                   string
	_configserviceRequesterAwsRegion                   string
	_configserviceResourceArn                          string
	_configserviceResourceDetails                      string
	_configserviceResourceEvaluationId                 string
	_configserviceResourceId                           string
	_configserviceResourceIdentifier                   string
	_configserviceResourceIdentifiers                  string
	_configserviceResourceIds                          []string
	_configserviceResourceKeys                         string
	_configserviceResourceName                         string
	_configserviceResourceType                         string
	_configserviceResourceTypes                        []string
	_configserviceResultToken                          string
	_configserviceRetentionConfigurationName           string
	_configserviceRetentionConfigurationNames          []string
	_configserviceRetentionPeriodInDays                string
	_configserviceSchemaVersionId                      string
	_configserviceServicePrincipal                     string
	_configserviceSortBy                               string
	_configserviceSortOrder                            string
	_configserviceStoredQuery                          string
	_configserviceTagKeys                              []string
	_configserviceTags                                 string
	_configserviceTemplateBody                         string
	_configserviceTemplateS3Uri                        string
	_configserviceTemplateSSMDocumentDetails           string
	_configserviceTestMode                             string
	_configserviceUpdateStatus                         string
)

// Adds all resource types specified in the ResourceTypes list to the [RecordingGroup] of
// specified configuration recorder and includes those resource types when
// recording.
//
// For this operation, the specified configuration recorder must use a [RecordingStrategy] that is
// either INCLUSION_BY_RESOURCE_TYPES or EXCLUSION_BY_RESOURCE_TYPES .
//
// [RecordingStrategy]: https://docs.aws.amazon.com/config/latest/APIReference/API_RecordingStrategy.html
// [RecordingGroup]: https://docs.aws.amazon.com/config/latest/APIReference/API_RecordingGroup.html
func configservice_AssociateResourceTypes(cfg aws.Config, client *configservice.Client) {
	input := &configservice.AssociateResourceTypesInput{
		// ConfigurationRecorderArn: *string, // Required
		// ResourceTypes: []types.ResourceType, // Required
	}

	if len(_configserviceConfigurationRecorderArn) > 0 {
		input.ConfigurationRecorderArn = aws.String(_configserviceConfigurationRecorderArn)
	}
	if len(_configserviceResourceTypes) > 0 {
		if err := assignInputField(input, "ResourceTypes", _configserviceResourceTypes[0]); err != nil {
			log.Errorf("invalid --resource-types: %s", err.Error())
			return
		}
	}

	if resp, err := client.AssociateResourceTypes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the current configuration items for resources that are present in your
// Config aggregator. The operation also returns a list of resources that are not
// processed in the current request. If there are no unprocessed resources, the
// operation returns an empty unprocessedResourceIdentifiers list.
//
// - The API does not return results for deleted resources.
//
// - The API does not return tags and relationships.
func configservice_BatchGetAggregateResourceConfig(cfg aws.Config, client *configservice.Client) {
	input := &configservice.BatchGetAggregateResourceConfigInput{
		// ConfigurationAggregatorName: *string, // Required
		// ResourceIdentifiers: []types.AggregateResourceIdentifier, // Required
	}

	if len(_configserviceConfigurationAggregatorName) > 0 {
		input.ConfigurationAggregatorName = aws.String(_configserviceConfigurationAggregatorName)
	}
	if len(_configserviceResourceIdentifiers) > 0 {
		if err := assignInputField(input, "ResourceIdentifiers", _configserviceResourceIdentifiers); err != nil {
			log.Errorf("invalid --resource-identifiers: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchGetAggregateResourceConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the BaseConfigurationItem for one or more requested resources. The
// operation also returns a list of resources that are not processed in the current
// request. If there are no unprocessed resources, the operation returns an empty
// unprocessedResourceKeys list.
//
// - The API does not return results for deleted resources.
//
// - The API does not return any tags for the requested resources. This
// information is filtered out of the supplementaryConfiguration section of the API
// response.
func configservice_BatchGetResourceConfig(cfg aws.Config, client *configservice.Client) {
	input := &configservice.BatchGetResourceConfigInput{
		// ResourceKeys: []types.ResourceKey, // Required
	}

	if len(_configserviceResourceKeys) > 0 {
		if err := assignInputField(input, "ResourceKeys", _configserviceResourceKeys); err != nil {
			log.Errorf("invalid --resource-keys: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchGetResourceConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the authorization granted to the specified configuration aggregator
// account in a specified region.
func configservice_DeleteAggregationAuthorization(cfg aws.Config, client *configservice.Client) {
	input := &configservice.DeleteAggregationAuthorizationInput{
		// AuthorizedAccountId: *string, // Required
		// AuthorizedAwsRegion: *string, // Required
	}

	if len(_configserviceAuthorizedAccountId) > 0 {
		input.AuthorizedAccountId = aws.String(_configserviceAuthorizedAccountId)
	}
	if len(_configserviceAuthorizedAwsRegion) > 0 {
		input.AuthorizedAwsRegion = aws.String(_configserviceAuthorizedAwsRegion)
	}

	if resp, err := client.DeleteAggregationAuthorization(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified Config rule and all of its evaluation results.
// Config sets the state of a rule to DELETING until the deletion is complete. You
// cannot update a rule while it is in this state. If you make a PutConfigRule or
// DeleteConfigRule request for the rule, you will receive a ResourceInUseException
// .
//
// You can check the state of a rule by using the DescribeConfigRules request.
//
// Recommendation: Consider excluding the AWS::Config::ResourceCompliance resource
// type from recording before deleting rules
//
// Deleting rules creates configuration items (CIs) for
// AWS::Config::ResourceCompliance that can affect your costs for the configuration
// recorder. If you are deleting rules which evaluate a large number of resource
// types, this can lead to a spike in the number of CIs recorded.
//
// To avoid the associated costs, you can opt to disable recording for the
// AWS::Config::ResourceCompliance resource type before deleting rules, and
// re-enable recording after the rules have been deleted.
//
// However, since deleting rules is an asynchronous process, it might take an hour
// or more to complete. During the time when recording is disabled for
// AWS::Config::ResourceCompliance , rule evaluations will not be recorded in the
// associated resource’s history.
func configservice_DeleteConfigRule(cfg aws.Config, client *configservice.Client) {
	input := &configservice.DeleteConfigRuleInput{
		// ConfigRuleName: *string, // Required
	}

	if len(_configserviceConfigRuleName) > 0 {
		input.ConfigRuleName = aws.String(_configserviceConfigRuleName)
	}

	if resp, err := client.DeleteConfigRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified configuration aggregator and the aggregated data
// associated with the aggregator.
func configservice_DeleteConfigurationAggregator(cfg aws.Config, client *configservice.Client) {
	input := &configservice.DeleteConfigurationAggregatorInput{
		// ConfigurationAggregatorName: *string, // Required
	}

	if len(_configserviceConfigurationAggregatorName) > 0 {
		input.ConfigurationAggregatorName = aws.String(_configserviceConfigurationAggregatorName)
	}

	if resp, err := client.DeleteConfigurationAggregator(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the customer managed configuration recorder.
// This operation does not delete the configuration information that was
// previously recorded. You will be able to access the previously recorded
// information by using the [GetResourceConfigHistory]operation, but you will not be able to access this
// information in the Config console until you have created a new customer managed
// configuration recorder.
//
// [GetResourceConfigHistory]: https://docs.aws.amazon.com/config/latest/APIReference/API_GetResourceConfigHistory.html
func configservice_DeleteConfigurationRecorder(cfg aws.Config, client *configservice.Client) {
	input := &configservice.DeleteConfigurationRecorderInput{
		// ConfigurationRecorderName: *string, // Required
	}

	if len(_configserviceConfigurationRecorderName) > 0 {
		input.ConfigurationRecorderName = aws.String(_configserviceConfigurationRecorderName)
	}

	if resp, err := client.DeleteConfigurationRecorder(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified conformance pack and all the Config rules, remediation
// actions, and all evaluation results within that conformance pack.
//
// Config sets the conformance pack to DELETE_IN_PROGRESS until the deletion is
// complete. You cannot update a conformance pack while it is in this state.
//
// Recommendation: Consider excluding the AWS::Config::ResourceCompliance resource
// type from recording before deleting rules
//
// Deleting rules creates configuration items (CIs) for
// AWS::Config::ResourceCompliance that can affect your costs for the configuration
// recorder. If you are deleting rules which evaluate a large number of resource
// types, this can lead to a spike in the number of CIs recorded.
//
// To avoid the associated costs, you can opt to disable recording for the
// AWS::Config::ResourceCompliance resource type before deleting rules, and
// re-enable recording after the rules have been deleted.
//
// However, since deleting rules is an asynchronous process, it might take an hour
// or more to complete. During the time when recording is disabled for
// AWS::Config::ResourceCompliance , rule evaluations will not be recorded in the
// associated resource’s history.
func configservice_DeleteConformancePack(cfg aws.Config, client *configservice.Client) {
	input := &configservice.DeleteConformancePackInput{
		// ConformancePackName: *string, // Required
	}

	if len(_configserviceConformancePackName) > 0 {
		input.ConformancePackName = aws.String(_configserviceConformancePackName)
	}

	if resp, err := client.DeleteConformancePack(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the delivery channel.
// Before you can delete the delivery channel, you must stop the customer managed
// configuration recorder. You can use the StopConfigurationRecorderoperation to stop the customer managed
// configuration recorder.
func configservice_DeleteDeliveryChannel(cfg aws.Config, client *configservice.Client) {
	input := &configservice.DeleteDeliveryChannelInput{
		// DeliveryChannelName: *string, // Required
	}

	if len(_configserviceDeliveryChannelName) > 0 {
		input.DeliveryChannelName = aws.String(_configserviceDeliveryChannelName)
	}

	if resp, err := client.DeleteDeliveryChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the evaluation results for the specified Config rule. You can specify
// one Config rule per request. After you delete the evaluation results, you can
// call the StartConfigRulesEvaluationAPI to start evaluating your Amazon Web Services resources against the
// rule.
func configservice_DeleteEvaluationResults(cfg aws.Config, client *configservice.Client) {
	input := &configservice.DeleteEvaluationResultsInput{
		// ConfigRuleName: *string, // Required
	}

	if len(_configserviceConfigRuleName) > 0 {
		input.ConfigRuleName = aws.String(_configserviceConfigRuleName)
	}

	if resp, err := client.DeleteEvaluationResults(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified organization Config rule and all of its evaluation
// results from all member accounts in that organization.
//
// Only a management account and a delegated administrator account can delete an
// organization Config rule. When calling this API with a delegated administrator,
// you must ensure Organizations ListDelegatedAdministrator permissions are added.
//
// Config sets the state of a rule to DELETE_IN_PROGRESS until the deletion is
// complete. You cannot update a rule while it is in this state.
//
// Recommendation: Consider excluding the AWS::Config::ResourceCompliance resource
// type from recording before deleting rules
//
// Deleting rules creates configuration items (CIs) for
// AWS::Config::ResourceCompliance that can affect your costs for the configuration
// recorder. If you are deleting rules which evaluate a large number of resource
// types, this can lead to a spike in the number of CIs recorded.
//
// To avoid the associated costs, you can opt to disable recording for the
// AWS::Config::ResourceCompliance resource type before deleting rules, and
// re-enable recording after the rules have been deleted.
//
// However, since deleting rules is an asynchronous process, it might take an hour
// or more to complete. During the time when recording is disabled for
// AWS::Config::ResourceCompliance , rule evaluations will not be recorded in the
// associated resource’s history.
func configservice_DeleteOrganizationConfigRule(cfg aws.Config, client *configservice.Client) {
	input := &configservice.DeleteOrganizationConfigRuleInput{
		// OrganizationConfigRuleName: *string, // Required
	}

	if len(_configserviceOrganizationConfigRuleName) > 0 {
		input.OrganizationConfigRuleName = aws.String(_configserviceOrganizationConfigRuleName)
	}

	if resp, err := client.DeleteOrganizationConfigRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified organization conformance pack and all of the Config rules
// and remediation actions from all member accounts in that organization.
//
// Only a management account or a delegated administrator account can delete an
// organization conformance pack. When calling this API with a delegated
// administrator, you must ensure Organizations ListDelegatedAdministrator
// permissions are added.
//
// Config sets the state of a conformance pack to DELETE_IN_PROGRESS until the
// deletion is complete. You cannot update a conformance pack while it is in this
// state.
//
// Recommendation: Consider excluding the AWS::Config::ResourceCompliance resource
// type from recording before deleting rules
//
// Deleting rules creates configuration items (CIs) for
// AWS::Config::ResourceCompliance that can affect your costs for the configuration
// recorder. If you are deleting rules which evaluate a large number of resource
// types, this can lead to a spike in the number of CIs recorded.
//
// To avoid the associated costs, you can opt to disable recording for the
// AWS::Config::ResourceCompliance resource type before deleting rules, and
// re-enable recording after the rules have been deleted.
//
// However, since deleting rules is an asynchronous process, it might take an hour
// or more to complete. During the time when recording is disabled for
// AWS::Config::ResourceCompliance , rule evaluations will not be recorded in the
// associated resource’s history.
func configservice_DeleteOrganizationConformancePack(cfg aws.Config, client *configservice.Client) {
	input := &configservice.DeleteOrganizationConformancePackInput{
		// OrganizationConformancePackName: *string, // Required
	}

	if len(_configserviceOrganizationConformancePackName) > 0 {
		input.OrganizationConformancePackName = aws.String(_configserviceOrganizationConformancePackName)
	}

	if resp, err := client.DeleteOrganizationConformancePack(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes pending authorization requests for a specified aggregator account in a
// specified region.
func configservice_DeletePendingAggregationRequest(cfg aws.Config, client *configservice.Client) {
	input := &configservice.DeletePendingAggregationRequestInput{
		// RequesterAccountId: *string, // Required
		// RequesterAwsRegion: *string, // Required
	}

	if len(_configserviceRequesterAccountId) > 0 {
		input.RequesterAccountId = aws.String(_configserviceRequesterAccountId)
	}
	if len(_configserviceRequesterAwsRegion) > 0 {
		input.RequesterAwsRegion = aws.String(_configserviceRequesterAwsRegion)
	}

	if resp, err := client.DeletePendingAggregationRequest(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the remediation configuration.
func configservice_DeleteRemediationConfiguration(cfg aws.Config, client *configservice.Client) {
	input := &configservice.DeleteRemediationConfigurationInput{
		// ConfigRuleName: *string, // Required
	}

	if len(_configserviceConfigRuleName) > 0 {
		input.ConfigRuleName = aws.String(_configserviceConfigRuleName)
	}
	if len(_configserviceResourceType) > 0 {
		input.ResourceType = aws.String(_configserviceResourceType)
	}

	if resp, err := client.DeleteRemediationConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes one or more remediation exceptions mentioned in the resource keys.
// Config generates a remediation exception when a problem occurs executing a
// remediation action to a specific resource. Remediation exceptions blocks
// auto-remediation until the exception is cleared.
func configservice_DeleteRemediationExceptions(cfg aws.Config, client *configservice.Client) {
	input := &configservice.DeleteRemediationExceptionsInput{
		// ConfigRuleName: *string, // Required
		// ResourceKeys: []types.RemediationExceptionResourceKey, // Required
	}

	if len(_configserviceConfigRuleName) > 0 {
		input.ConfigRuleName = aws.String(_configserviceConfigRuleName)
	}
	if len(_configserviceResourceKeys) > 0 {
		if err := assignInputField(input, "ResourceKeys", _configserviceResourceKeys); err != nil {
			log.Errorf("invalid --resource-keys: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteRemediationExceptions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Records the configuration state for a custom resource that has been deleted.
// This API records a new ConfigurationItem with a ResourceDeleted status. You can
// retrieve the ConfigurationItems recorded for this resource in your Config
// History.
func configservice_DeleteResourceConfig(cfg aws.Config, client *configservice.Client) {
	input := &configservice.DeleteResourceConfigInput{
		// ResourceId: *string, // Required
		// ResourceType: *string, // Required
	}

	if len(_configserviceResourceId) > 0 {
		input.ResourceId = aws.String(_configserviceResourceId)
	}
	if len(_configserviceResourceType) > 0 {
		input.ResourceType = aws.String(_configserviceResourceType)
	}

	if resp, err := client.DeleteResourceConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the retention configuration.
func configservice_DeleteRetentionConfiguration(cfg aws.Config, client *configservice.Client) {
	input := &configservice.DeleteRetentionConfigurationInput{
		// RetentionConfigurationName: *string, // Required
	}

	if len(_configserviceRetentionConfigurationName) > 0 {
		input.RetentionConfigurationName = aws.String(_configserviceRetentionConfigurationName)
	}

	if resp, err := client.DeleteRetentionConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing service-linked configuration recorder.
// This operation does not delete the configuration information that was
// previously recorded. You will be able to access the previously recorded
// information by using the [GetResourceConfigHistory]operation, but you will not be able to access this
// information in the Config console until you have created a new service-linked
// configuration recorder for the same service.
//
// # The recording scope determines if you receive configuration items
//
// The recording scope is set by the service that is linked to the configuration
// recorder and determines whether you receive configuration items (CIs) in the
// delivery channel. If the recording scope is internal, you will not receive CIs
// in the delivery channel.
//
// [GetResourceConfigHistory]: https://docs.aws.amazon.com/config/latest/APIReference/API_GetResourceConfigHistory.html
func configservice_DeleteServiceLinkedConfigurationRecorder(cfg aws.Config, client *configservice.Client) {
	input := &configservice.DeleteServiceLinkedConfigurationRecorderInput{
		// ServicePrincipal: *string, // Required
	}

	if len(_configserviceServicePrincipal) > 0 {
		input.ServicePrincipal = aws.String(_configserviceServicePrincipal)
	}

	if resp, err := client.DeleteServiceLinkedConfigurationRecorder(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the stored query for a single Amazon Web Services account and a single
// Amazon Web Services Region.
func configservice_DeleteStoredQuery(cfg aws.Config, client *configservice.Client) {
	input := &configservice.DeleteStoredQueryInput{
		// QueryName: *string, // Required
	}

	if len(_configserviceQueryName) > 0 {
		input.QueryName = aws.String(_configserviceQueryName)
	}

	if resp, err := client.DeleteStoredQuery(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Schedules delivery of a configuration snapshot to the Amazon S3 bucket in the
// specified delivery channel. After the delivery has started, Config sends the
// following notifications using an Amazon SNS topic that you have specified.
//
// - Notification of the start of the delivery.
//
// - Notification of the completion of the delivery, if the delivery was
// successfully completed.
//
// - Notification of delivery failure, if the delivery failed.
func configservice_DeliverConfigSnapshot(cfg aws.Config, client *configservice.Client) {
	input := &configservice.DeliverConfigSnapshotInput{
		// DeliveryChannelName: *string, // Required
	}

	if len(_configserviceDeliveryChannelName) > 0 {
		input.DeliveryChannelName = aws.String(_configserviceDeliveryChannelName)
	}

	if resp, err := client.DeliverConfigSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of compliant and noncompliant rules with the number of resources
// for compliant and noncompliant rules. Does not display rules that do not have
// compliance results.
//
// The results can return an empty result page, but if you have a nextToken , the
// results are displayed on the next page.
func configservice_DescribeAggregateComplianceByConfigRules(cfg aws.Config, client *configservice.Client) {
	input := &configservice.DescribeAggregateComplianceByConfigRulesInput{
		// ConfigurationAggregatorName: *string, // Required
	}

	if len(_configserviceConfigurationAggregatorName) > 0 {
		input.ConfigurationAggregatorName = aws.String(_configserviceConfigurationAggregatorName)
	}
	if len(_configserviceFilters) > 0 {
		if err := assignInputField(input, "Filters", _configserviceFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_configserviceLimit) > 0 {
		if err := assignInputField(input, "Limit", _configserviceLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_configserviceNextToken) > 0 {
		input.NextToken = aws.String(_configserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeAggregateComplianceByConfigRules(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*configservice.DescribeAggregateComplianceByConfigRulesOutput
	p := configservice.NewDescribeAggregateComplianceByConfigRulesPaginator(client, input)
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

// Returns a list of the existing and deleted conformance packs and their
// associated compliance status with the count of compliant and noncompliant Config
// rules within each conformance pack. Also returns the total rule count which
// includes compliant rules, noncompliant rules, and rules that cannot be evaluated
// due to insufficient data.
//
// The results can return an empty result page, but if you have a nextToken , the
// results are displayed on the next page.
func configservice_DescribeAggregateComplianceByConformancePacks(cfg aws.Config, client *configservice.Client) {
	input := &configservice.DescribeAggregateComplianceByConformancePacksInput{
		// ConfigurationAggregatorName: *string, // Required
	}

	if len(_configserviceConfigurationAggregatorName) > 0 {
		input.ConfigurationAggregatorName = aws.String(_configserviceConfigurationAggregatorName)
	}
	if len(_configserviceFilters) > 0 {
		if err := assignInputField(input, "Filters", _configserviceFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_configserviceLimit) > 0 {
		if err := assignInputField(input, "Limit", _configserviceLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_configserviceNextToken) > 0 {
		input.NextToken = aws.String(_configserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeAggregateComplianceByConformancePacks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*configservice.DescribeAggregateComplianceByConformancePacksOutput
	p := configservice.NewDescribeAggregateComplianceByConformancePacksPaginator(client, input)
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

// Returns a list of authorizations granted to various aggregator accounts and
// regions.
func configservice_DescribeAggregationAuthorizations(cfg aws.Config, client *configservice.Client) {
	input := &configservice.DescribeAggregationAuthorizationsInput{}

	if len(_configserviceLimit) > 0 {
		if err := assignInputField(input, "Limit", _configserviceLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_configserviceNextToken) > 0 {
		input.NextToken = aws.String(_configserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeAggregationAuthorizations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*configservice.DescribeAggregationAuthorizationsOutput
	p := configservice.NewDescribeAggregationAuthorizationsPaginator(client, input)
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

// Indicates whether the specified Config rules are compliant. If a rule is
// noncompliant, this operation returns the number of Amazon Web Services resources
// that do not comply with the rule.
//
// A rule is compliant if all of the evaluated resources comply with it. It is
// noncompliant if any of these resources do not comply.
//
// If Config has no current evaluation results for the rule, it returns
// INSUFFICIENT_DATA . This result might indicate one of the following conditions:
//
// - Config has never invoked an evaluation for the rule. To check whether it
// has, use the DescribeConfigRuleEvaluationStatus action to get the
// LastSuccessfulInvocationTime and LastFailedInvocationTime .
//
// - The rule's Lambda function is failing to send evaluation results to Config.
// Verify that the role you assigned to your configuration recorder includes the
// config:PutEvaluations permission. If the rule is a custom rule, verify that
// the Lambda execution role includes the config:PutEvaluations permission.
//
// - The rule's Lambda function has returned NOT_APPLICABLE for all evaluation
// results. This can occur if the resources were deleted or removed from the rule's
// scope.
func configservice_DescribeComplianceByConfigRule(cfg aws.Config, client *configservice.Client) {
	input := &configservice.DescribeComplianceByConfigRuleInput{}

	if len(_configserviceComplianceTypes) > 0 {
		if err := assignInputField(input, "ComplianceTypes", _configserviceComplianceTypes); err != nil {
			log.Errorf("invalid --compliance-types: %s", err.Error())
			return
		}
	}
	if len(_configserviceConfigRuleNames) > 0 {
		input.ConfigRuleNames = append([]string(nil), _configserviceConfigRuleNames...)
	}
	if len(_configserviceNextToken) > 0 {
		input.NextToken = aws.String(_configserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeComplianceByConfigRule(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*configservice.DescribeComplianceByConfigRuleOutput
	p := configservice.NewDescribeComplianceByConfigRulePaginator(client, input)
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

// Indicates whether the specified Amazon Web Services resources are compliant. If
// a resource is noncompliant, this operation returns the number of Config rules
// that the resource does not comply with.
//
// A resource is compliant if it complies with all the Config rules that evaluate
// it. It is noncompliant if it does not comply with one or more of these rules.
//
// If Config has no current evaluation results for the resource, it returns
// INSUFFICIENT_DATA . This result might indicate one of the following conditions
// about the rules that evaluate the resource:
//
// - Config has never invoked an evaluation for the rule. To check whether it
// has, use the DescribeConfigRuleEvaluationStatus action to get the
// LastSuccessfulInvocationTime and LastFailedInvocationTime .
//
// - The rule's Lambda function is failing to send evaluation results to Config.
// Verify that the role that you assigned to your configuration recorder includes
// the config:PutEvaluations permission. If the rule is a custom rule, verify
// that the Lambda execution role includes the config:PutEvaluations permission.
//
// - The rule's Lambda function has returned NOT_APPLICABLE for all evaluation
// results. This can occur if the resources were deleted or removed from the rule's
// scope.
func configservice_DescribeComplianceByResource(cfg aws.Config, client *configservice.Client) {
	input := &configservice.DescribeComplianceByResourceInput{}

	if len(_configserviceComplianceTypes) > 0 {
		if err := assignInputField(input, "ComplianceTypes", _configserviceComplianceTypes); err != nil {
			log.Errorf("invalid --compliance-types: %s", err.Error())
			return
		}
	}
	if len(_configserviceLimit) > 0 {
		if err := assignInputField(input, "Limit", _configserviceLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_configserviceNextToken) > 0 {
		input.NextToken = aws.String(_configserviceNextToken)
	}
	if len(_configserviceResourceId) > 0 {
		input.ResourceId = aws.String(_configserviceResourceId)
	}
	if len(_configserviceResourceType) > 0 {
		input.ResourceType = aws.String(_configserviceResourceType)
	}

	if disablePaginator() {
		if resp, err := client.DescribeComplianceByResource(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*configservice.DescribeComplianceByResourceOutput
	p := configservice.NewDescribeComplianceByResourcePaginator(client, input)
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

// Returns status information for each of your Config managed rules. The status
// includes information such as the last time Config invoked the rule, the last
// time Config failed to invoke the rule, and the related error for the last
// failure.
func configservice_DescribeConfigRuleEvaluationStatus(cfg aws.Config, client *configservice.Client) {
	input := &configservice.DescribeConfigRuleEvaluationStatusInput{}

	if len(_configserviceConfigRuleNames) > 0 {
		input.ConfigRuleNames = append([]string(nil), _configserviceConfigRuleNames...)
	}
	if len(_configserviceLimit) > 0 {
		if err := assignInputField(input, "Limit", _configserviceLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_configserviceNextToken) > 0 {
		input.NextToken = aws.String(_configserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeConfigRuleEvaluationStatus(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*configservice.DescribeConfigRuleEvaluationStatusOutput
	p := configservice.NewDescribeConfigRuleEvaluationStatusPaginator(client, input)
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

// Returns details about your Config rules.
func configservice_DescribeConfigRules(cfg aws.Config, client *configservice.Client) {
	input := &configservice.DescribeConfigRulesInput{}

	if len(_configserviceConfigRuleNames) > 0 {
		input.ConfigRuleNames = append([]string(nil), _configserviceConfigRuleNames...)
	}
	if len(_configserviceFilters) > 0 {
		if err := assignInputField(input, "Filters", _configserviceFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_configserviceNextToken) > 0 {
		input.NextToken = aws.String(_configserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeConfigRules(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*configservice.DescribeConfigRulesOutput
	p := configservice.NewDescribeConfigRulesPaginator(client, input)
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

// Returns status information for sources within an aggregator. The status
// includes information about the last time Config verified authorization between
// the source account and an aggregator account. In case of a failure, the status
// contains the related error code or message.
func configservice_DescribeConfigurationAggregatorSourcesStatus(cfg aws.Config, client *configservice.Client) {
	input := &configservice.DescribeConfigurationAggregatorSourcesStatusInput{
		// ConfigurationAggregatorName: *string, // Required
	}

	if len(_configserviceConfigurationAggregatorName) > 0 {
		input.ConfigurationAggregatorName = aws.String(_configserviceConfigurationAggregatorName)
	}
	if len(_configserviceLimit) > 0 {
		if err := assignInputField(input, "Limit", _configserviceLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_configserviceNextToken) > 0 {
		input.NextToken = aws.String(_configserviceNextToken)
	}
	if len(_configserviceUpdateStatus) > 0 {
		if err := assignInputField(input, "UpdateStatus", _configserviceUpdateStatus); err != nil {
			log.Errorf("invalid --update-status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeConfigurationAggregatorSourcesStatus(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*configservice.DescribeConfigurationAggregatorSourcesStatusOutput
	p := configservice.NewDescribeConfigurationAggregatorSourcesStatusPaginator(client, input)
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

// Returns the details of one or more configuration aggregators. If the
// configuration aggregator is not specified, this operation returns the details
// for all the configuration aggregators associated with the account.
func configservice_DescribeConfigurationAggregators(cfg aws.Config, client *configservice.Client) {
	input := &configservice.DescribeConfigurationAggregatorsInput{}

	if len(_configserviceConfigurationAggregatorNames) > 0 {
		input.ConfigurationAggregatorNames = append([]string(nil), _configserviceConfigurationAggregatorNames...)
	}
	if len(_configserviceLimit) > 0 {
		if err := assignInputField(input, "Limit", _configserviceLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_configserviceNextToken) > 0 {
		input.NextToken = aws.String(_configserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeConfigurationAggregators(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*configservice.DescribeConfigurationAggregatorsOutput
	p := configservice.NewDescribeConfigurationAggregatorsPaginator(client, input)
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

// Returns the current status of the configuration recorder you specify as well as
// the status of the last recording event for the configuration recorders.
//
// For a detailed status of recording events over time, add your Config events to
// Amazon CloudWatch metrics and use CloudWatch metrics.
//
// If a configuration recorder is not specified, this operation returns the status
// for the customer managed configuration recorder configured for the account, if
// applicable.
//
// When making a request to this operation, you can only specify one configuration
// recorder.
func configservice_DescribeConfigurationRecorderStatus(cfg aws.Config, client *configservice.Client) {
	input := &configservice.DescribeConfigurationRecorderStatusInput{}

	if len(_configserviceArn) > 0 {
		input.Arn = aws.String(_configserviceArn)
	}
	if len(_configserviceConfigurationRecorderNames) > 0 {
		input.ConfigurationRecorderNames = append([]string(nil), _configserviceConfigurationRecorderNames...)
	}
	if len(_configserviceServicePrincipal) > 0 {
		input.ServicePrincipal = aws.String(_configserviceServicePrincipal)
	}

	if resp, err := client.DescribeConfigurationRecorderStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns details for the configuration recorder you specify.
// If a configuration recorder is not specified, this operation returns details
// for the customer managed configuration recorder configured for the account, if
// applicable.
//
// When making a request to this operation, you can only specify one configuration
// recorder.
func configservice_DescribeConfigurationRecorders(cfg aws.Config, client *configservice.Client) {
	input := &configservice.DescribeConfigurationRecordersInput{}

	if len(_configserviceArn) > 0 {
		input.Arn = aws.String(_configserviceArn)
	}
	if len(_configserviceConfigurationRecorderNames) > 0 {
		input.ConfigurationRecorderNames = append([]string(nil), _configserviceConfigurationRecorderNames...)
	}
	if len(_configserviceServicePrincipal) > 0 {
		input.ServicePrincipal = aws.String(_configserviceServicePrincipal)
	}

	if resp, err := client.DescribeConfigurationRecorders(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns compliance details for each rule in that conformance pack.
// You must provide exact rule names.
func configservice_DescribeConformancePackCompliance(cfg aws.Config, client *configservice.Client) {
	input := &configservice.DescribeConformancePackComplianceInput{
		// ConformancePackName: *string, // Required
	}

	if len(_configserviceConformancePackName) > 0 {
		input.ConformancePackName = aws.String(_configserviceConformancePackName)
	}
	if len(_configserviceFilters) > 0 {
		if err := assignInputField(input, "Filters", _configserviceFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_configserviceLimit) > 0 {
		if err := assignInputField(input, "Limit", _configserviceLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_configserviceNextToken) > 0 {
		input.NextToken = aws.String(_configserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeConformancePackCompliance(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*configservice.DescribeConformancePackComplianceOutput
	p := configservice.NewDescribeConformancePackCompliancePaginator(client, input)
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

// Provides one or more conformance packs deployment status.
// If there are no conformance packs then you will see an empty result.
func configservice_DescribeConformancePackStatus(cfg aws.Config, client *configservice.Client) {
	input := &configservice.DescribeConformancePackStatusInput{}

	if len(_configserviceConformancePackNames) > 0 {
		input.ConformancePackNames = append([]string(nil), _configserviceConformancePackNames...)
	}
	if len(_configserviceLimit) > 0 {
		if err := assignInputField(input, "Limit", _configserviceLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_configserviceNextToken) > 0 {
		input.NextToken = aws.String(_configserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeConformancePackStatus(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*configservice.DescribeConformancePackStatusOutput
	p := configservice.NewDescribeConformancePackStatusPaginator(client, input)
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

// Returns a list of one or more conformance packs.
func configservice_DescribeConformancePacks(cfg aws.Config, client *configservice.Client) {
	input := &configservice.DescribeConformancePacksInput{}

	if len(_configserviceConformancePackNames) > 0 {
		input.ConformancePackNames = append([]string(nil), _configserviceConformancePackNames...)
	}
	if len(_configserviceLimit) > 0 {
		if err := assignInputField(input, "Limit", _configserviceLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_configserviceNextToken) > 0 {
		input.NextToken = aws.String(_configserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeConformancePacks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*configservice.DescribeConformancePacksOutput
	p := configservice.NewDescribeConformancePacksPaginator(client, input)
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

// Returns the current status of the specified delivery channel. If a delivery
// channel is not specified, this operation returns the current status of all
// delivery channels associated with the account.
//
// Currently, you can specify only one delivery channel per region in your account.
func configservice_DescribeDeliveryChannelStatus(cfg aws.Config, client *configservice.Client) {
	input := &configservice.DescribeDeliveryChannelStatusInput{}

	if len(_configserviceDeliveryChannelNames) > 0 {
		input.DeliveryChannelNames = append([]string(nil), _configserviceDeliveryChannelNames...)
	}

	if resp, err := client.DescribeDeliveryChannelStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns details about the specified delivery channel. If a delivery channel is
// not specified, this operation returns the details of all delivery channels
// associated with the account.
//
// Currently, you can specify only one delivery channel per region in your account.
func configservice_DescribeDeliveryChannels(cfg aws.Config, client *configservice.Client) {
	input := &configservice.DescribeDeliveryChannelsInput{}

	if len(_configserviceDeliveryChannelNames) > 0 {
		input.DeliveryChannelNames = append([]string(nil), _configserviceDeliveryChannelNames...)
	}

	if resp, err := client.DescribeDeliveryChannels(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides organization Config rule deployment status for an organization.
// The status is not considered successful until organization Config rule is
// successfully deployed in all the member accounts with an exception of excluded
// accounts.
//
// When you specify the limit and the next token, you receive a paginated
// response. Limit and next token are not applicable if you specify organization
// Config rule names. It is only applicable, when you request all the organization
// Config rules.
func configservice_DescribeOrganizationConfigRuleStatuses(cfg aws.Config, client *configservice.Client) {
	input := &configservice.DescribeOrganizationConfigRuleStatusesInput{}

	if len(_configserviceLimit) > 0 {
		if err := assignInputField(input, "Limit", _configserviceLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_configserviceNextToken) > 0 {
		input.NextToken = aws.String(_configserviceNextToken)
	}
	if len(_configserviceOrganizationConfigRuleNames) > 0 {
		input.OrganizationConfigRuleNames = append([]string(nil), _configserviceOrganizationConfigRuleNames...)
	}

	if disablePaginator() {
		if resp, err := client.DescribeOrganizationConfigRuleStatuses(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*configservice.DescribeOrganizationConfigRuleStatusesOutput
	p := configservice.NewDescribeOrganizationConfigRuleStatusesPaginator(client, input)
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

// Returns a list of organization Config rules.
// When you specify the limit and the next token, you receive a paginated response.
//
// Limit and next token are not applicable if you specify organization Config rule
// names. It is only applicable, when you request all the organization Config
// rules.
//
// # For accounts within an organization
//
// If you deploy an organizational rule or conformance pack in an organization
// administrator account, and then establish a delegated administrator and deploy
// an organizational rule or conformance pack in the delegated administrator
// account, you won't be able to see the organizational rule or conformance pack in
// the organization administrator account from the delegated administrator account
// or see the organizational rule or conformance pack in the delegated
// administrator account from organization administrator account. The
// DescribeOrganizationConfigRules and DescribeOrganizationConformancePacks APIs
// can only see and interact with the organization-related resource that were
// deployed from within the account calling those APIs.
func configservice_DescribeOrganizationConfigRules(cfg aws.Config, client *configservice.Client) {
	input := &configservice.DescribeOrganizationConfigRulesInput{}

	if len(_configserviceLimit) > 0 {
		if err := assignInputField(input, "Limit", _configserviceLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_configserviceNextToken) > 0 {
		input.NextToken = aws.String(_configserviceNextToken)
	}
	if len(_configserviceOrganizationConfigRuleNames) > 0 {
		input.OrganizationConfigRuleNames = append([]string(nil), _configserviceOrganizationConfigRuleNames...)
	}

	if disablePaginator() {
		if resp, err := client.DescribeOrganizationConfigRules(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*configservice.DescribeOrganizationConfigRulesOutput
	p := configservice.NewDescribeOrganizationConfigRulesPaginator(client, input)
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

// Provides organization conformance pack deployment status for an organization.
// The status is not considered successful until organization conformance pack is
// successfully deployed in all the member accounts with an exception of excluded
// accounts.
//
// When you specify the limit and the next token, you receive a paginated
// response. Limit and next token are not applicable if you specify organization
// conformance pack names. They are only applicable, when you request all the
// organization conformance packs.
func configservice_DescribeOrganizationConformancePackStatuses(cfg aws.Config, client *configservice.Client) {
	input := &configservice.DescribeOrganizationConformancePackStatusesInput{}

	if len(_configserviceLimit) > 0 {
		if err := assignInputField(input, "Limit", _configserviceLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_configserviceNextToken) > 0 {
		input.NextToken = aws.String(_configserviceNextToken)
	}
	if len(_configserviceOrganizationConformancePackNames) > 0 {
		input.OrganizationConformancePackNames = append([]string(nil), _configserviceOrganizationConformancePackNames...)
	}

	if disablePaginator() {
		if resp, err := client.DescribeOrganizationConformancePackStatuses(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*configservice.DescribeOrganizationConformancePackStatusesOutput
	p := configservice.NewDescribeOrganizationConformancePackStatusesPaginator(client, input)
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

// Returns a list of organization conformance packs.
// When you specify the limit and the next token, you receive a paginated
// response.
//
// Limit and next token are not applicable if you specify organization conformance
// packs names. They are only applicable, when you request all the organization
// conformance packs.
//
// # For accounts within an organization
//
// If you deploy an organizational rule or conformance pack in an organization
// administrator account, and then establish a delegated administrator and deploy
// an organizational rule or conformance pack in the delegated administrator
// account, you won't be able to see the organizational rule or conformance pack in
// the organization administrator account from the delegated administrator account
// or see the organizational rule or conformance pack in the delegated
// administrator account from organization administrator account. The
// DescribeOrganizationConfigRules and DescribeOrganizationConformancePacks APIs
// can only see and interact with the organization-related resource that were
// deployed from within the account calling those APIs.
func configservice_DescribeOrganizationConformancePacks(cfg aws.Config, client *configservice.Client) {
	input := &configservice.DescribeOrganizationConformancePacksInput{}

	if len(_configserviceLimit) > 0 {
		if err := assignInputField(input, "Limit", _configserviceLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_configserviceNextToken) > 0 {
		input.NextToken = aws.String(_configserviceNextToken)
	}
	if len(_configserviceOrganizationConformancePackNames) > 0 {
		input.OrganizationConformancePackNames = append([]string(nil), _configserviceOrganizationConformancePackNames...)
	}

	if disablePaginator() {
		if resp, err := client.DescribeOrganizationConformancePacks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*configservice.DescribeOrganizationConformancePacksOutput
	p := configservice.NewDescribeOrganizationConformancePacksPaginator(client, input)
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

// Returns a list of all pending aggregation requests.
func configservice_DescribePendingAggregationRequests(cfg aws.Config, client *configservice.Client) {
	input := &configservice.DescribePendingAggregationRequestsInput{}

	if len(_configserviceLimit) > 0 {
		if err := assignInputField(input, "Limit", _configserviceLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_configserviceNextToken) > 0 {
		input.NextToken = aws.String(_configserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribePendingAggregationRequests(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*configservice.DescribePendingAggregationRequestsOutput
	p := configservice.NewDescribePendingAggregationRequestsPaginator(client, input)
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

// Returns the details of one or more remediation configurations.
func configservice_DescribeRemediationConfigurations(cfg aws.Config, client *configservice.Client) {
	input := &configservice.DescribeRemediationConfigurationsInput{
		// ConfigRuleNames: []string, // Required
	}

	if len(_configserviceConfigRuleNames) > 0 {
		input.ConfigRuleNames = append([]string(nil), _configserviceConfigRuleNames...)
	}

	if resp, err := client.DescribeRemediationConfigurations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the details of one or more remediation exceptions. A detailed view of a
// remediation exception for a set of resources that includes an explanation of an
// exception and the time when the exception will be deleted. When you specify the
// limit and the next token, you receive a paginated response.
//
// Config generates a remediation exception when a problem occurs executing a
// remediation action to a specific resource. Remediation exceptions blocks
// auto-remediation until the exception is cleared.
//
// When you specify the limit and the next token, you receive a paginated
// response.
//
// Limit and next token are not applicable if you request resources in batch. It
// is only applicable, when you request all resources.
func configservice_DescribeRemediationExceptions(cfg aws.Config, client *configservice.Client) {
	input := &configservice.DescribeRemediationExceptionsInput{
		// ConfigRuleName: *string, // Required
	}

	if len(_configserviceConfigRuleName) > 0 {
		input.ConfigRuleName = aws.String(_configserviceConfigRuleName)
	}
	if len(_configserviceLimit) > 0 {
		if err := assignInputField(input, "Limit", _configserviceLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_configserviceNextToken) > 0 {
		input.NextToken = aws.String(_configserviceNextToken)
	}
	if len(_configserviceResourceKeys) > 0 {
		if err := assignInputField(input, "ResourceKeys", _configserviceResourceKeys); err != nil {
			log.Errorf("invalid --resource-keys: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeRemediationExceptions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*configservice.DescribeRemediationExceptionsOutput
	p := configservice.NewDescribeRemediationExceptionsPaginator(client, input)
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

// Provides a detailed view of a Remediation Execution for a set of resources
// including state, timestamps for when steps for the remediation execution occur,
// and any error messages for steps that have failed. When you specify the limit
// and the next token, you receive a paginated response.
func configservice_DescribeRemediationExecutionStatus(cfg aws.Config, client *configservice.Client) {
	input := &configservice.DescribeRemediationExecutionStatusInput{
		// ConfigRuleName: *string, // Required
	}

	if len(_configserviceConfigRuleName) > 0 {
		input.ConfigRuleName = aws.String(_configserviceConfigRuleName)
	}
	if len(_configserviceLimit) > 0 {
		if err := assignInputField(input, "Limit", _configserviceLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_configserviceNextToken) > 0 {
		input.NextToken = aws.String(_configserviceNextToken)
	}
	if len(_configserviceResourceKeys) > 0 {
		if err := assignInputField(input, "ResourceKeys", _configserviceResourceKeys); err != nil {
			log.Errorf("invalid --resource-keys: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeRemediationExecutionStatus(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*configservice.DescribeRemediationExecutionStatusOutput
	p := configservice.NewDescribeRemediationExecutionStatusPaginator(client, input)
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

// Returns the details of one or more retention configurations. If the retention
// configuration name is not specified, this operation returns the details for all
// the retention configurations for that account.
//
// Currently, Config supports only one retention configuration per region in your
// account.
func configservice_DescribeRetentionConfigurations(cfg aws.Config, client *configservice.Client) {
	input := &configservice.DescribeRetentionConfigurationsInput{}

	if len(_configserviceNextToken) > 0 {
		input.NextToken = aws.String(_configserviceNextToken)
	}
	if len(_configserviceRetentionConfigurationNames) > 0 {
		input.RetentionConfigurationNames = append([]string(nil), _configserviceRetentionConfigurationNames...)
	}

	if disablePaginator() {
		if resp, err := client.DescribeRetentionConfigurations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*configservice.DescribeRetentionConfigurationsOutput
	p := configservice.NewDescribeRetentionConfigurationsPaginator(client, input)
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

// Removes all resource types specified in the ResourceTypes list from the [RecordingGroup] of
// configuration recorder and excludes these resource types when recording.
//
// For this operation, the configuration recorder must use a [RecordingStrategy] that is either
// INCLUSION_BY_RESOURCE_TYPES or EXCLUSION_BY_RESOURCE_TYPES .
//
// [RecordingStrategy]: https://docs.aws.amazon.com/config/latest/APIReference/API_RecordingStrategy.html
// [RecordingGroup]: https://docs.aws.amazon.com/config/latest/APIReference/API_RecordingGroup.html
func configservice_DisassociateResourceTypes(cfg aws.Config, client *configservice.Client) {
	input := &configservice.DisassociateResourceTypesInput{
		// ConfigurationRecorderArn: *string, // Required
		// ResourceTypes: []types.ResourceType, // Required
	}

	if len(_configserviceConfigurationRecorderArn) > 0 {
		input.ConfigurationRecorderArn = aws.String(_configserviceConfigurationRecorderArn)
	}
	if len(_configserviceResourceTypes) > 0 {
		if err := assignInputField(input, "ResourceTypes", _configserviceResourceTypes[0]); err != nil {
			log.Errorf("invalid --resource-types: %s", err.Error())
			return
		}
	}

	if resp, err := client.DisassociateResourceTypes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the evaluation results for the specified Config rule for a specific
// resource in a rule. The results indicate which Amazon Web Services resources
// were evaluated by the rule, when each resource was last evaluated, and whether
// each resource complies with the rule.
//
// The results can return an empty result page. But if you have a nextToken , the
// results are displayed on the next page.
func configservice_GetAggregateComplianceDetailsByConfigRule(cfg aws.Config, client *configservice.Client) {
	input := &configservice.GetAggregateComplianceDetailsByConfigRuleInput{
		// AccountId: *string, // Required
		// AwsRegion: *string, // Required
		// ConfigRuleName: *string, // Required
		// ConfigurationAggregatorName: *string, // Required
	}

	if len(_configserviceAccountId) > 0 {
		input.AccountId = aws.String(_configserviceAccountId)
	}
	if len(_configserviceAwsRegion) > 0 {
		input.AwsRegion = aws.String(_configserviceAwsRegion)
	}
	if len(_configserviceConfigRuleName) > 0 {
		input.ConfigRuleName = aws.String(_configserviceConfigRuleName)
	}
	if len(_configserviceConfigurationAggregatorName) > 0 {
		input.ConfigurationAggregatorName = aws.String(_configserviceConfigurationAggregatorName)
	}
	if len(_configserviceComplianceType) > 0 {
		if err := assignInputField(input, "ComplianceType", _configserviceComplianceType); err != nil {
			log.Errorf("invalid --compliance-type: %s", err.Error())
			return
		}
	}
	if len(_configserviceLimit) > 0 {
		if err := assignInputField(input, "Limit", _configserviceLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_configserviceNextToken) > 0 {
		input.NextToken = aws.String(_configserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetAggregateComplianceDetailsByConfigRule(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*configservice.GetAggregateComplianceDetailsByConfigRuleOutput
	p := configservice.NewGetAggregateComplianceDetailsByConfigRulePaginator(client, input)
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

// Returns the number of compliant and noncompliant rules for one or more accounts
// and regions in an aggregator.
//
// The results can return an empty result page, but if you have a nextToken, the
// results are displayed on the next page.
func configservice_GetAggregateConfigRuleComplianceSummary(cfg aws.Config, client *configservice.Client) {
	input := &configservice.GetAggregateConfigRuleComplianceSummaryInput{
		// ConfigurationAggregatorName: *string, // Required
	}

	if len(_configserviceConfigurationAggregatorName) > 0 {
		input.ConfigurationAggregatorName = aws.String(_configserviceConfigurationAggregatorName)
	}
	if len(_configserviceFilters) > 0 {
		if err := assignInputField(input, "Filters", _configserviceFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_configserviceGroupByKey) > 0 {
		if err := assignInputField(input, "GroupByKey", _configserviceGroupByKey); err != nil {
			log.Errorf("invalid --group-by-key: %s", err.Error())
			return
		}
	}
	if len(_configserviceLimit) > 0 {
		if err := assignInputField(input, "Limit", _configserviceLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_configserviceNextToken) > 0 {
		input.NextToken = aws.String(_configserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetAggregateConfigRuleComplianceSummary(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*configservice.GetAggregateConfigRuleComplianceSummaryOutput
	p := configservice.NewGetAggregateConfigRuleComplianceSummaryPaginator(client, input)
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

// Returns the count of compliant and noncompliant conformance packs across all
// Amazon Web Services accounts and Amazon Web Services Regions in an aggregator.
// You can filter based on Amazon Web Services account ID or Amazon Web Services
// Region.
//
// The results can return an empty result page, but if you have a nextToken, the
// results are displayed on the next page.
func configservice_GetAggregateConformancePackComplianceSummary(cfg aws.Config, client *configservice.Client) {
	input := &configservice.GetAggregateConformancePackComplianceSummaryInput{
		// ConfigurationAggregatorName: *string, // Required
	}

	if len(_configserviceConfigurationAggregatorName) > 0 {
		input.ConfigurationAggregatorName = aws.String(_configserviceConfigurationAggregatorName)
	}
	if len(_configserviceFilters) > 0 {
		if err := assignInputField(input, "Filters", _configserviceFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_configserviceGroupByKey) > 0 {
		if err := assignInputField(input, "GroupByKey", _configserviceGroupByKey); err != nil {
			log.Errorf("invalid --group-by-key: %s", err.Error())
			return
		}
	}
	if len(_configserviceLimit) > 0 {
		if err := assignInputField(input, "Limit", _configserviceLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_configserviceNextToken) > 0 {
		input.NextToken = aws.String(_configserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetAggregateConformancePackComplianceSummary(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*configservice.GetAggregateConformancePackComplianceSummaryOutput
	p := configservice.NewGetAggregateConformancePackComplianceSummaryPaginator(client, input)
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

// Returns the resource counts across accounts and regions that are present in
// your Config aggregator. You can request the resource counts by providing filters
// and GroupByKey.
//
// For example, if the input contains accountID 12345678910 and region us-east-1
// in filters, the API returns the count of resources in account ID 12345678910 and
// region us-east-1. If the input contains ACCOUNT_ID as a GroupByKey, the API
// returns resource counts for all source accounts that are present in your
// aggregator.
func configservice_GetAggregateDiscoveredResourceCounts(cfg aws.Config, client *configservice.Client) {
	input := &configservice.GetAggregateDiscoveredResourceCountsInput{
		// ConfigurationAggregatorName: *string, // Required
	}

	if len(_configserviceConfigurationAggregatorName) > 0 {
		input.ConfigurationAggregatorName = aws.String(_configserviceConfigurationAggregatorName)
	}
	if len(_configserviceFilters) > 0 {
		if err := assignInputField(input, "Filters", _configserviceFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_configserviceGroupByKey) > 0 {
		if err := assignInputField(input, "GroupByKey", _configserviceGroupByKey); err != nil {
			log.Errorf("invalid --group-by-key: %s", err.Error())
			return
		}
	}
	if len(_configserviceLimit) > 0 {
		if err := assignInputField(input, "Limit", _configserviceLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_configserviceNextToken) > 0 {
		input.NextToken = aws.String(_configserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetAggregateDiscoveredResourceCounts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*configservice.GetAggregateDiscoveredResourceCountsOutput
	p := configservice.NewGetAggregateDiscoveredResourceCountsPaginator(client, input)
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

// Returns configuration item that is aggregated for your specific resource in a
// specific source account and region.
//
// The API does not return results for deleted resources.
func configservice_GetAggregateResourceConfig(cfg aws.Config, client *configservice.Client) {
	input := &configservice.GetAggregateResourceConfigInput{
		// ConfigurationAggregatorName: *string, // Required
		// ResourceIdentifier: *types.AggregateResourceIdentifier, // Required
	}

	if len(_configserviceConfigurationAggregatorName) > 0 {
		input.ConfigurationAggregatorName = aws.String(_configserviceConfigurationAggregatorName)
	}
	if len(_configserviceResourceIdentifier) > 0 {
		if err := assignInputField(input, "ResourceIdentifier", _configserviceResourceIdentifier); err != nil {
			log.Errorf("invalid --resource-identifier: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetAggregateResourceConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the evaluation results for the specified Config rule. The results
// indicate which Amazon Web Services resources were evaluated by the rule, when
// each resource was last evaluated, and whether each resource complies with the
// rule.
func configservice_GetComplianceDetailsByConfigRule(cfg aws.Config, client *configservice.Client) {
	input := &configservice.GetComplianceDetailsByConfigRuleInput{
		// ConfigRuleName: *string, // Required
	}

	if len(_configserviceConfigRuleName) > 0 {
		input.ConfigRuleName = aws.String(_configserviceConfigRuleName)
	}
	if len(_configserviceComplianceTypes) > 0 {
		if err := assignInputField(input, "ComplianceTypes", _configserviceComplianceTypes); err != nil {
			log.Errorf("invalid --compliance-types: %s", err.Error())
			return
		}
	}
	if len(_configserviceLimit) > 0 {
		if err := assignInputField(input, "Limit", _configserviceLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_configserviceNextToken) > 0 {
		input.NextToken = aws.String(_configserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetComplianceDetailsByConfigRule(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*configservice.GetComplianceDetailsByConfigRuleOutput
	p := configservice.NewGetComplianceDetailsByConfigRulePaginator(client, input)
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

// Returns the evaluation results for the specified Amazon Web Services resource.
// The results indicate which Config rules were used to evaluate the resource, when
// each rule was last invoked, and whether the resource complies with each rule.
func configservice_GetComplianceDetailsByResource(cfg aws.Config, client *configservice.Client) {
	input := &configservice.GetComplianceDetailsByResourceInput{}

	if len(_configserviceComplianceTypes) > 0 {
		if err := assignInputField(input, "ComplianceTypes", _configserviceComplianceTypes); err != nil {
			log.Errorf("invalid --compliance-types: %s", err.Error())
			return
		}
	}
	if len(_configserviceNextToken) > 0 {
		input.NextToken = aws.String(_configserviceNextToken)
	}
	if len(_configserviceResourceEvaluationId) > 0 {
		input.ResourceEvaluationId = aws.String(_configserviceResourceEvaluationId)
	}
	if len(_configserviceResourceId) > 0 {
		input.ResourceId = aws.String(_configserviceResourceId)
	}
	if len(_configserviceResourceType) > 0 {
		input.ResourceType = aws.String(_configserviceResourceType)
	}

	if disablePaginator() {
		if resp, err := client.GetComplianceDetailsByResource(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*configservice.GetComplianceDetailsByResourceOutput
	p := configservice.NewGetComplianceDetailsByResourcePaginator(client, input)
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

// Returns the number of Config rules that are compliant and noncompliant, up to a
// maximum of 25 for each.
func configservice_GetComplianceSummaryByConfigRule(cfg aws.Config, client *configservice.Client) {
	input := &configservice.GetComplianceSummaryByConfigRuleInput{}

	if resp, err := client.GetComplianceSummaryByConfigRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the number of resources that are compliant and the number that are
// noncompliant. You can specify one or more resource types to get these numbers
// for each resource type. The maximum number returned is 100.
func configservice_GetComplianceSummaryByResourceType(cfg aws.Config, client *configservice.Client) {
	input := &configservice.GetComplianceSummaryByResourceTypeInput{}

	if len(_configserviceResourceTypes) > 0 {
		input.ResourceTypes = append([]string(nil), _configserviceResourceTypes...)
	}

	if resp, err := client.GetComplianceSummaryByResourceType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns compliance details of a conformance pack for all Amazon Web Services
// resources that are monitered by conformance pack.
func configservice_GetConformancePackComplianceDetails(cfg aws.Config, client *configservice.Client) {
	input := &configservice.GetConformancePackComplianceDetailsInput{
		// ConformancePackName: *string, // Required
	}

	if len(_configserviceConformancePackName) > 0 {
		input.ConformancePackName = aws.String(_configserviceConformancePackName)
	}
	if len(_configserviceFilters) > 0 {
		if err := assignInputField(input, "Filters", _configserviceFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_configserviceLimit) > 0 {
		if err := assignInputField(input, "Limit", _configserviceLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_configserviceNextToken) > 0 {
		input.NextToken = aws.String(_configserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetConformancePackComplianceDetails(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*configservice.GetConformancePackComplianceDetailsOutput
	p := configservice.NewGetConformancePackComplianceDetailsPaginator(client, input)
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

// Returns compliance details for the conformance pack based on the cumulative
// compliance results of all the rules in that conformance pack.
func configservice_GetConformancePackComplianceSummary(cfg aws.Config, client *configservice.Client) {
	input := &configservice.GetConformancePackComplianceSummaryInput{
		// ConformancePackNames: []string, // Required
	}

	if len(_configserviceConformancePackNames) > 0 {
		input.ConformancePackNames = append([]string(nil), _configserviceConformancePackNames...)
	}
	if len(_configserviceLimit) > 0 {
		if err := assignInputField(input, "Limit", _configserviceLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_configserviceNextToken) > 0 {
		input.NextToken = aws.String(_configserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetConformancePackComplianceSummary(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*configservice.GetConformancePackComplianceSummaryOutput
	p := configservice.NewGetConformancePackComplianceSummaryPaginator(client, input)
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

// Returns the policy definition containing the logic for your Config Custom
// Policy rule.
func configservice_GetCustomRulePolicy(cfg aws.Config, client *configservice.Client) {
	input := &configservice.GetCustomRulePolicyInput{}

	if len(_configserviceConfigRuleName) > 0 {
		input.ConfigRuleName = aws.String(_configserviceConfigRuleName)
	}

	if resp, err := client.GetCustomRulePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the resource types, the number of each resource type, and the total
// number of resources that Config is recording in this region for your Amazon Web
// Services account.
//
// # Example
//
// - Config is recording three resource types in the US East (Ohio) Region for
// your account: 25 EC2 instances, 20 IAM users, and 15 S3 buckets.
//
// - You make a call to the GetDiscoveredResourceCounts action and specify that
// you want all resource types.
//
// - Config returns the following:
//
// - The resource types (EC2 instances, IAM users, and S3 buckets).
//
// - The number of each resource type (25, 20, and 15).
//
// - The total number of all resources (60).
//
// The response is paginated. By default, Config lists 100 ResourceCount objects on each page.
// You can customize this number with the limit parameter. The response includes a
// nextToken string. To get the next page of results, run the request again and
// specify the string for the nextToken parameter.
//
// If you make a call to the GetDiscoveredResourceCounts action, you might not immediately receive resource
// counts in the following situations:
//
// - You are a new Config customer.
//
// - You just enabled resource recording.
//
// It might take a few minutes for Config to record and count your resources. Wait
// a few minutes and then retry the GetDiscoveredResourceCountsaction.
func configservice_GetDiscoveredResourceCounts(cfg aws.Config, client *configservice.Client) {
	input := &configservice.GetDiscoveredResourceCountsInput{}

	if len(_configserviceLimit) > 0 {
		if err := assignInputField(input, "Limit", _configserviceLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_configserviceNextToken) > 0 {
		input.NextToken = aws.String(_configserviceNextToken)
	}
	if len(_configserviceResourceTypes) > 0 {
		input.ResourceTypes = append([]string(nil), _configserviceResourceTypes...)
	}

	if disablePaginator() {
		if resp, err := client.GetDiscoveredResourceCounts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*configservice.GetDiscoveredResourceCountsOutput
	p := configservice.NewGetDiscoveredResourceCountsPaginator(client, input)
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

// Returns detailed status for each member account within an organization for a
// given organization Config rule.
func configservice_GetOrganizationConfigRuleDetailedStatus(cfg aws.Config, client *configservice.Client) {
	input := &configservice.GetOrganizationConfigRuleDetailedStatusInput{
		// OrganizationConfigRuleName: *string, // Required
	}

	if len(_configserviceOrganizationConfigRuleName) > 0 {
		input.OrganizationConfigRuleName = aws.String(_configserviceOrganizationConfigRuleName)
	}
	if len(_configserviceFilters) > 0 {
		if err := assignInputField(input, "Filters", _configserviceFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_configserviceLimit) > 0 {
		if err := assignInputField(input, "Limit", _configserviceLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_configserviceNextToken) > 0 {
		input.NextToken = aws.String(_configserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetOrganizationConfigRuleDetailedStatus(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*configservice.GetOrganizationConfigRuleDetailedStatusOutput
	p := configservice.NewGetOrganizationConfigRuleDetailedStatusPaginator(client, input)
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

// Returns detailed status for each member account within an organization for a
// given organization conformance pack.
func configservice_GetOrganizationConformancePackDetailedStatus(cfg aws.Config, client *configservice.Client) {
	input := &configservice.GetOrganizationConformancePackDetailedStatusInput{
		// OrganizationConformancePackName: *string, // Required
	}

	if len(_configserviceOrganizationConformancePackName) > 0 {
		input.OrganizationConformancePackName = aws.String(_configserviceOrganizationConformancePackName)
	}
	if len(_configserviceFilters) > 0 {
		if err := assignInputField(input, "Filters", _configserviceFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_configserviceLimit) > 0 {
		if err := assignInputField(input, "Limit", _configserviceLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_configserviceNextToken) > 0 {
		input.NextToken = aws.String(_configserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetOrganizationConformancePackDetailedStatus(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*configservice.GetOrganizationConformancePackDetailedStatusOutput
	p := configservice.NewGetOrganizationConformancePackDetailedStatusPaginator(client, input)
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

// Returns the policy definition containing the logic for your organization Config
// Custom Policy rule.
func configservice_GetOrganizationCustomRulePolicy(cfg aws.Config, client *configservice.Client) {
	input := &configservice.GetOrganizationCustomRulePolicyInput{
		// OrganizationConfigRuleName: *string, // Required
	}

	if len(_configserviceOrganizationConfigRuleName) > 0 {
		input.OrganizationConfigRuleName = aws.String(_configserviceOrganizationConfigRuleName)
	}

	if resp, err := client.GetOrganizationCustomRulePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// For accurate reporting on the compliance status, you must record the
// AWS::Config::ResourceCompliance resource type.
//
// For more information, see [Recording Amazon Web Services Resources] in the Config Resources Developer Guide.
//
// Returns a list of configurations items (CIs) for the specified resource.
//
// # Contents
//
// The list contains details about each state of the resource during the specified
// time interval. If you specified a retention period to retain your CIs between a
// minimum of 30 days and a maximum of 7 years (2557 days), Config returns the CIs
// for the specified retention period.
//
// # Pagination
//
// The response is paginated. By default, Config returns a limit of 10
// configuration items per page. You can customize this number with the limit
// parameter. The response includes a nextToken string. To get the next page of
// results, run the request again and specify the string for the nextToken
// parameter.
//
// Each call to the API is limited to span a duration of seven days. It is likely
// that the number of records returned is smaller than the specified limit . In
// such cases, you can make another call, using the nextToken .
//
// [Recording Amazon Web Services Resources]: https://docs.aws.amazon.com/config/latest/developerguide/select-resources.html
func configservice_GetResourceConfigHistory(cfg aws.Config, client *configservice.Client) {
	input := &configservice.GetResourceConfigHistoryInput{
		// ResourceId: *string, // Required
		// ResourceType: types.ResourceType, // Required
	}

	if len(_configserviceResourceId) > 0 {
		input.ResourceId = aws.String(_configserviceResourceId)
	}
	if len(_configserviceResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _configserviceResourceType); err != nil {
			log.Errorf("invalid --resource-type: %s", err.Error())
			return
		}
	}
	if len(_configserviceChronologicalOrder) > 0 {
		if err := assignInputField(input, "ChronologicalOrder", _configserviceChronologicalOrder); err != nil {
			log.Errorf("invalid --chronological-order: %s", err.Error())
			return
		}
	}
	if len(_configserviceEarlierTime) > 0 {
		if err := assignInputField(input, "EarlierTime", _configserviceEarlierTime); err != nil {
			log.Errorf("invalid --earlier-time: %s", err.Error())
			return
		}
	}
	if len(_configserviceLaterTime) > 0 {
		if err := assignInputField(input, "LaterTime", _configserviceLaterTime); err != nil {
			log.Errorf("invalid --later-time: %s", err.Error())
			return
		}
	}
	if len(_configserviceLimit) > 0 {
		if err := assignInputField(input, "Limit", _configserviceLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_configserviceNextToken) > 0 {
		input.NextToken = aws.String(_configserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetResourceConfigHistory(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*configservice.GetResourceConfigHistoryOutput
	p := configservice.NewGetResourceConfigHistoryPaginator(client, input)
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

// Returns a summary of resource evaluation for the specified resource evaluation
// ID from the proactive rules that were run. The results indicate which evaluation
// context was used to evaluate the rules, which resource details were evaluated,
// the evaluation mode that was run, and whether the resource details comply with
// the configuration of the proactive rules.
//
// To see additional information about the evaluation result, such as which rule
// flagged a resource as NON_COMPLIANT, use the [GetComplianceDetailsByResource]API. For more information, see the [Examples]
// section.
//
// [GetComplianceDetailsByResource]: https://docs.aws.amazon.com/config/latest/APIReference/API_GetComplianceDetailsByResource.html
// [Examples]: https://docs.aws.amazon.com/config/latest/APIReference/API_GetResourceEvaluationSummary.html#API_GetResourceEvaluationSummary_Examples
func configservice_GetResourceEvaluationSummary(cfg aws.Config, client *configservice.Client) {
	input := &configservice.GetResourceEvaluationSummaryInput{
		// ResourceEvaluationId: *string, // Required
	}

	if len(_configserviceResourceEvaluationId) > 0 {
		input.ResourceEvaluationId = aws.String(_configserviceResourceEvaluationId)
	}

	if resp, err := client.GetResourceEvaluationSummary(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the details of a specific stored query.
func configservice_GetStoredQuery(cfg aws.Config, client *configservice.Client) {
	input := &configservice.GetStoredQueryInput{
		// QueryName: *string, // Required
	}

	if len(_configserviceQueryName) > 0 {
		input.QueryName = aws.String(_configserviceQueryName)
	}

	if resp, err := client.GetStoredQuery(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Accepts a resource type and returns a list of resource identifiers that are
// aggregated for a specific resource type across accounts and regions. A resource
// identifier includes the resource type, ID, (if available) the custom resource
// name, source account, and source region. You can narrow the results to include
// only resources that have specific resource IDs, or a resource name, or source
// account ID, or source region.
//
// For example, if the input consists of accountID 12345678910 and the region is
// us-east-1 for resource type AWS::EC2::Instance then the API returns all the EC2
// instance identifiers of accountID 12345678910 and region us-east-1.
func configservice_ListAggregateDiscoveredResources(cfg aws.Config, client *configservice.Client) {
	input := &configservice.ListAggregateDiscoveredResourcesInput{
		// ConfigurationAggregatorName: *string, // Required
		// ResourceType: types.ResourceType, // Required
	}

	if len(_configserviceConfigurationAggregatorName) > 0 {
		input.ConfigurationAggregatorName = aws.String(_configserviceConfigurationAggregatorName)
	}
	if len(_configserviceResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _configserviceResourceType); err != nil {
			log.Errorf("invalid --resource-type: %s", err.Error())
			return
		}
	}
	if len(_configserviceFilters) > 0 {
		if err := assignInputField(input, "Filters", _configserviceFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_configserviceLimit) > 0 {
		if err := assignInputField(input, "Limit", _configserviceLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_configserviceNextToken) > 0 {
		input.NextToken = aws.String(_configserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAggregateDiscoveredResources(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*configservice.ListAggregateDiscoveredResourcesOutput
	p := configservice.NewListAggregateDiscoveredResourcesPaginator(client, input)
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

// Returns a list of configuration recorders depending on the filters you specify.
func configservice_ListConfigurationRecorders(cfg aws.Config, client *configservice.Client) {
	input := &configservice.ListConfigurationRecordersInput{}

	if len(_configserviceFilters) > 0 {
		if err := assignInputField(input, "Filters", _configserviceFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_configserviceMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _configserviceMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_configserviceNextToken) > 0 {
		input.NextToken = aws.String(_configserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListConfigurationRecorders(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*configservice.ListConfigurationRecordersOutput
	p := configservice.NewListConfigurationRecordersPaginator(client, input)
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

// Returns a list of conformance pack compliance scores. A compliance score is the
// percentage of the number of compliant rule-resource combinations in a
// conformance pack compared to the number of total possible rule-resource
// combinations in the conformance pack. This metric provides you with a high-level
// view of the compliance state of your conformance packs. You can use it to
// identify, investigate, and understand the level of compliance in your
// conformance packs.
//
// Conformance packs with no evaluation results will have a compliance score of
// INSUFFICIENT_DATA .
func configservice_ListConformancePackComplianceScores(cfg aws.Config, client *configservice.Client) {
	input := &configservice.ListConformancePackComplianceScoresInput{}

	if len(_configserviceFilters) > 0 {
		if err := assignInputField(input, "Filters", _configserviceFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_configserviceLimit) > 0 {
		if err := assignInputField(input, "Limit", _configserviceLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_configserviceNextToken) > 0 {
		input.NextToken = aws.String(_configserviceNextToken)
	}
	if len(_configserviceSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _configserviceSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_configserviceSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _configserviceSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListConformancePackComplianceScores(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*configservice.ListConformancePackComplianceScoresOutput
	p := configservice.NewListConformancePackComplianceScoresPaginator(client, input)
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

// Returns a list of resource resource identifiers for the specified resource
// types for the resources of that type. A resource identifier includes the
// resource type, ID, and (if available) the custom resource name.
//
// The results consist of resources that Config has discovered, including those
// that Config is not currently recording. You can narrow the results to include
// only resources that have specific resource IDs or a resource name.
//
// You can specify either resource IDs or a resource name, but not both, in the
// same request.
//
// # CloudFormation stack recording behavior in Config
//
// When a CloudFormation stack fails to create (for example, it enters the
// ROLLBACK_FAILED state), Config does not record a configuration item (CI) for
// that stack. Configuration items are only recorded for stacks that reach the
// following states:
//
// - CREATE_COMPLETE
//
// - UPDATE_COMPLETE
//
// - UPDATE_ROLLBACK_COMPLETE
//
// - UPDATE_ROLLBACK_FAILED
//
// - DELETE_FAILED
//
// - DELETE_COMPLETE
//
// Because no CI is created for a failed stack creation, you won't see
// configuration history for that stack in Config, even after the stack is deleted.
// This helps make sure that Config only tracks resources that were successfully
// provisioned.
func configservice_ListDiscoveredResources(cfg aws.Config, client *configservice.Client) {
	input := &configservice.ListDiscoveredResourcesInput{
		// ResourceType: types.ResourceType, // Required
	}

	if len(_configserviceResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _configserviceResourceType); err != nil {
			log.Errorf("invalid --resource-type: %s", err.Error())
			return
		}
	}
	if len(_configserviceIncludeDeletedResources) > 0 {
		if err := assignInputField(input, "IncludeDeletedResources", _configserviceIncludeDeletedResources); err != nil {
			log.Errorf("invalid --include-deleted-resources: %s", err.Error())
			return
		}
	}
	if len(_configserviceLimit) > 0 {
		if err := assignInputField(input, "Limit", _configserviceLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_configserviceNextToken) > 0 {
		input.NextToken = aws.String(_configserviceNextToken)
	}
	if len(_configserviceResourceIds) > 0 {
		input.ResourceIds = append([]string(nil), _configserviceResourceIds...)
	}
	if len(_configserviceResourceName) > 0 {
		input.ResourceName = aws.String(_configserviceResourceName)
	}

	if disablePaginator() {
		if resp, err := client.ListDiscoveredResources(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*configservice.ListDiscoveredResourcesOutput
	p := configservice.NewListDiscoveredResourcesPaginator(client, input)
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

// Returns a list of proactive resource evaluations.
func configservice_ListResourceEvaluations(cfg aws.Config, client *configservice.Client) {
	input := &configservice.ListResourceEvaluationsInput{}

	if len(_configserviceFilters) > 0 {
		if err := assignInputField(input, "Filters", _configserviceFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_configserviceLimit) > 0 {
		if err := assignInputField(input, "Limit", _configserviceLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_configserviceNextToken) > 0 {
		input.NextToken = aws.String(_configserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListResourceEvaluations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*configservice.ListResourceEvaluationsOutput
	p := configservice.NewListResourceEvaluationsPaginator(client, input)
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

// Lists the stored queries for a single Amazon Web Services account and a single
// Amazon Web Services Region. The default is 100.
func configservice_ListStoredQueries(cfg aws.Config, client *configservice.Client) {
	input := &configservice.ListStoredQueriesInput{}

	if len(_configserviceMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _configserviceMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_configserviceNextToken) > 0 {
		input.NextToken = aws.String(_configserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListStoredQueries(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*configservice.ListStoredQueriesOutput
	p := configservice.NewListStoredQueriesPaginator(client, input)
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

// List the tags for Config resource.
func configservice_ListTagsForResource(cfg aws.Config, client *configservice.Client) {
	input := &configservice.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_configserviceResourceArn) > 0 {
		input.ResourceArn = aws.String(_configserviceResourceArn)
	}
	if len(_configserviceLimit) > 0 {
		if err := assignInputField(input, "Limit", _configserviceLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_configserviceNextToken) > 0 {
		input.NextToken = aws.String(_configserviceNextToken)
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

	var results []*configservice.ListTagsForResourceOutput
	p := configservice.NewListTagsForResourcePaginator(client, input)
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

// Authorizes the aggregator account and region to collect data from the source
// account and region.
//
// # Tags are added at creation and cannot be updated with this operation
//
// PutAggregationAuthorization is an idempotent API. Subsequent requests won’t
// create a duplicate resource if one was already created. If a following request
// has different tags values, Config will ignore these differences and treat it as
// an idempotent request of the previous. In this case, tags will not be updated,
// even if they are different.
//
// Use [TagResource] and [UntagResource] to update tags after creation.
//
// [TagResource]: https://docs.aws.amazon.com/config/latest/APIReference/API_TagResource.html
// [UntagResource]: https://docs.aws.amazon.com/config/latest/APIReference/API_UntagResource.html
func configservice_PutAggregationAuthorization(cfg aws.Config, client *configservice.Client) {
	input := &configservice.PutAggregationAuthorizationInput{
		// AuthorizedAccountId: *string, // Required
		// AuthorizedAwsRegion: *string, // Required
	}

	if len(_configserviceAuthorizedAccountId) > 0 {
		input.AuthorizedAccountId = aws.String(_configserviceAuthorizedAccountId)
	}
	if len(_configserviceAuthorizedAwsRegion) > 0 {
		input.AuthorizedAwsRegion = aws.String(_configserviceAuthorizedAwsRegion)
	}
	if len(_configserviceTags) > 0 {
		if err := assignInputField(input, "Tags", _configserviceTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutAggregationAuthorization(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds or updates an Config rule to evaluate if your Amazon Web Services
// resources comply with your desired configurations. For information on how many
// Config rules you can have per account, see [Service Limits]in the Config Developer Guide.
//
// There are two types of rules: Config Managed Rules and Config Custom Rules. You
// can use PutConfigRule to create both Config Managed Rules and Config Custom
// Rules.
//
// Config Managed Rules are predefined, customizable rules created by Config. For
// a list of managed rules, see [List of Config Managed Rules]. If you are adding an Config managed rule, you
// must specify the rule's identifier for the SourceIdentifier key.
//
// Config Custom Rules are rules that you create from scratch. There are two ways
// to create Config custom rules: with Lambda functions ([Lambda Developer Guide] ) and with Guard ([Guard GitHub Repository] ), a
// policy-as-code language.
//
// Config custom rules created with Lambda are called Config Custom Lambda Rules
// and Config custom rules created with Guard are called Config Custom Policy
// Rules.
//
// If you are adding a new Config Custom Lambda rule, you first need to create an
// Lambda function that the rule invokes to evaluate your resources. When you use
// PutConfigRule to add a Custom Lambda rule to Config, you must specify the Amazon
// Resource Name (ARN) that Lambda assigns to the function. You specify the ARN in
// the SourceIdentifier key. This key is part of the Source object, which is part
// of the ConfigRule object.
//
// For any new Config rule that you add, specify the ConfigRuleName in the
// ConfigRule object. Do not specify the ConfigRuleArn or the ConfigRuleId . These
// values are generated by Config for new rules.
//
// If you are updating a rule that you added previously, you can specify the rule
// by ConfigRuleName , ConfigRuleId , or ConfigRuleArn in the ConfigRule data type
// that you use in this request.
//
// For more information about developing and using Config rules, see [Evaluating Resources with Config Rules] in the
// Config Developer Guide.
//
// # Tags are added at creation and cannot be updated with this operation
//
// PutConfigRule is an idempotent API. Subsequent requests won’t create a
// duplicate resource if one was already created. If a following request has
// different tags values, Config will ignore these differences and treat it as an
// idempotent request of the previous. In this case, tags will not be updated,
// even if they are different.
//
// Use [TagResource] and [UntagResource] to update tags after creation.
//
// [List of Config Managed Rules]: https://docs.aws.amazon.com/config/latest/developerguide/managed-rules-by-aws-config.html
// [Lambda Developer Guide]: https://docs.aws.amazon.com/config/latest/developerguide/gettingstarted-concepts.html#gettingstarted-concepts-function
// [TagResource]: https://docs.aws.amazon.com/config/latest/APIReference/API_TagResource.html
// [Evaluating Resources with Config Rules]: https://docs.aws.amazon.com/config/latest/developerguide/evaluate-config.html
// [UntagResource]: https://docs.aws.amazon.com/config/latest/APIReference/API_UntagResource.html
// [Service Limits]: https://docs.aws.amazon.com/config/latest/developerguide/configlimits.html
// [Guard GitHub Repository]: https://github.com/aws-cloudformation/cloudformation-guard
func configservice_PutConfigRule(cfg aws.Config, client *configservice.Client) {
	input := &configservice.PutConfigRuleInput{
		// ConfigRule: *types.ConfigRule, // Required
	}

	if len(_configserviceConfigRule) > 0 {
		if err := assignInputField(input, "ConfigRule", _configserviceConfigRule); err != nil {
			log.Errorf("invalid --config-rule: %s", err.Error())
			return
		}
	}
	if len(_configserviceTags) > 0 {
		if err := assignInputField(input, "Tags", _configserviceTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutConfigRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates and updates the configuration aggregator with the selected source
// accounts and regions. The source account can be individual account(s) or an
// organization.
//
// accountIds that are passed will be replaced with existing accounts. If you want
// to add additional accounts into the aggregator, call
// DescribeConfigurationAggregators to get the previous accounts and then append
// new ones.
//
// Config should be enabled in source accounts and regions you want to aggregate.
//
// If your source type is an organization, you must be signed in to the management
// account or a registered delegated administrator and all the features must be
// enabled in your organization. If the caller is a management account, Config
// calls EnableAwsServiceAccess API to enable integration between Config and
// Organizations. If the caller is a registered delegated administrator, Config
// calls ListDelegatedAdministrators API to verify whether the caller is a valid
// delegated administrator.
//
// To register a delegated administrator, see [Register a Delegated Administrator] in the Config developer guide.
//
// # Tags are added at creation and cannot be updated with this operation
//
// PutConfigurationAggregator is an idempotent API. Subsequent requests won’t
// create a duplicate resource if one was already created. If a following request
// has different tags values, Config will ignore these differences and treat it as
// an idempotent request of the previous. In this case, tags will not be updated,
// even if they are different.
//
// Use [TagResource] and [UntagResource] to update tags after creation.
//
// [Register a Delegated Administrator]: https://docs.aws.amazon.com/config/latest/developerguide/set-up-aggregator-cli.html#register-a-delegated-administrator-cli
// [TagResource]: https://docs.aws.amazon.com/config/latest/APIReference/API_TagResource.html
// [UntagResource]: https://docs.aws.amazon.com/config/latest/APIReference/API_UntagResource.html
func configservice_PutConfigurationAggregator(cfg aws.Config, client *configservice.Client) {
	input := &configservice.PutConfigurationAggregatorInput{
		// ConfigurationAggregatorName: *string, // Required
	}

	if len(_configserviceConfigurationAggregatorName) > 0 {
		input.ConfigurationAggregatorName = aws.String(_configserviceConfigurationAggregatorName)
	}
	if len(_configserviceAccountAggregationSources) > 0 {
		if err := assignInputField(input, "AccountAggregationSources", _configserviceAccountAggregationSources); err != nil {
			log.Errorf("invalid --account-aggregation-sources: %s", err.Error())
			return
		}
	}
	if len(_configserviceAggregatorFilters) > 0 {
		if err := assignInputField(input, "AggregatorFilters", _configserviceAggregatorFilters); err != nil {
			log.Errorf("invalid --aggregator-filters: %s", err.Error())
			return
		}
	}
	if len(_configserviceOrganizationAggregationSource) > 0 {
		if err := assignInputField(input, "OrganizationAggregationSource", _configserviceOrganizationAggregationSource); err != nil {
			log.Errorf("invalid --organization-aggregation-source: %s", err.Error())
			return
		}
	}
	if len(_configserviceTags) > 0 {
		if err := assignInputField(input, "Tags", _configserviceTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutConfigurationAggregator(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates the customer managed configuration recorder.
// You can use this operation to create a new customer managed configuration
// recorder or to update the roleARN and the recordingGroup for an existing
// customer managed configuration recorder.
//
// To start the customer managed configuration recorder and begin recording
// configuration changes for the resource types you specify, use the [StartConfigurationRecorder]operation.
//
// For more information, see [Working with the Configuration Recorder] in the Config Developer Guide.
//
// # One customer managed configuration recorder per account per Region
//
// You can create only one customer managed configuration recorder for each
// account for each Amazon Web Services Region.
//
// Default is to record all supported resource types, excluding the global IAM
// resource types
//
// If you have not specified values for the recordingGroup field, the default for
// the customer managed configuration recorder is to record all supported resource
// types, excluding the global IAM resource types: AWS::IAM::Group ,
// AWS::IAM::Policy , AWS::IAM::Role , and AWS::IAM::User .
//
// # Tags are added at creation and cannot be updated
//
// PutConfigurationRecorder is an idempotent API. Subsequent requests won’t create
// a duplicate resource if one was already created. If a following request has
// different tags values, Config will ignore these differences and treat it as an
// idempotent request of the previous. In this case, tags will not be updated, even
// if they are different.
//
// Use [TagResource] and [UntagResource] to update tags after creation.
//
// [Working with the Configuration Recorder]: https://docs.aws.amazon.com/config/latest/developerguide/stop-start-recorder.html
// [TagResource]: https://docs.aws.amazon.com/config/latest/APIReference/API_TagResource.html
// [UntagResource]: https://docs.aws.amazon.com/config/latest/APIReference/API_UntagResource.html
// [StartConfigurationRecorder]: https://docs.aws.amazon.com/config/latest/APIReference/API_StartConfigurationRecorder.html
func configservice_PutConfigurationRecorder(cfg aws.Config, client *configservice.Client) {
	input := &configservice.PutConfigurationRecorderInput{
		// ConfigurationRecorder: *types.ConfigurationRecorder, // Required
	}

	if len(_configserviceConfigurationRecorder) > 0 {
		if err := assignInputField(input, "ConfigurationRecorder", _configserviceConfigurationRecorder); err != nil {
			log.Errorf("invalid --configuration-recorder: %s", err.Error())
			return
		}
	}
	if len(_configserviceTags) > 0 {
		if err := assignInputField(input, "Tags", _configserviceTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutConfigurationRecorder(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates a conformance pack. A conformance pack is a collection of
// Config rules that can be easily deployed in an account and a region and across
// an organization. For information on how many conformance packs you can have per
// account, see [Service Limits]in the Config Developer Guide.
//
// When you use PutConformancePack to deploy conformance packs in your account,
// the operation can create Config rules and remediation actions without requiring
// config:PutConfigRule or config:PutRemediationConfigurations permissions in your
// account IAM policies.
//
// This API uses the AWSServiceRoleForConfigConforms service-linked role in your
// account to create conformance pack resources. This service-linked role includes
// the permissions to create Config rules and remediation configurations, even if
// your account IAM policies explicitly deny these actions.
//
// This API creates a service-linked role AWSServiceRoleForConfigConforms in your
// account. The service-linked role is created only when the role does not exist in
// your account.
//
// You must specify only one of the follow parameters: TemplateS3Uri , TemplateBody
// or TemplateSSMDocumentDetails .
//
// # Tags are added at creation and cannot be updated with this operation
//
// PutConformancePack is an idempotent API. Subsequent requests won't create a
// duplicate resource if one was already created. If a following request has
// different tags values, Config will ignore these differences and treat it as an
// idempotent request of the previous. In this case, tags will not be updated,
// even if they are different.
//
// Use [TagResource] and [UntagResource] to update tags after creation.
//
// [TagResource]: https://docs.aws.amazon.com/config/latest/APIReference/API_TagResource.html
// [UntagResource]: https://docs.aws.amazon.com/config/latest/APIReference/API_UntagResource.html
// [Service Limits]: https://docs.aws.amazon.com/config/latest/developerguide/configlimits.html
func configservice_PutConformancePack(cfg aws.Config, client *configservice.Client) {
	input := &configservice.PutConformancePackInput{
		// ConformancePackName: *string, // Required
	}

	if len(_configserviceConformancePackName) > 0 {
		input.ConformancePackName = aws.String(_configserviceConformancePackName)
	}
	if len(_configserviceConformancePackInputParameters) > 0 {
		if err := assignInputField(input, "ConformancePackInputParameters", _configserviceConformancePackInputParameters); err != nil {
			log.Errorf("invalid --conformance-pack-input-parameters: %s", err.Error())
			return
		}
	}
	if len(_configserviceDeliveryS3Bucket) > 0 {
		input.DeliveryS3Bucket = aws.String(_configserviceDeliveryS3Bucket)
	}
	if len(_configserviceDeliveryS3KeyPrefix) > 0 {
		input.DeliveryS3KeyPrefix = aws.String(_configserviceDeliveryS3KeyPrefix)
	}
	if len(_configserviceTags) > 0 {
		if err := assignInputField(input, "Tags", _configserviceTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_configserviceTemplateBody) > 0 {
		input.TemplateBody = aws.String(_configserviceTemplateBody)
	}
	if len(_configserviceTemplateS3Uri) > 0 {
		input.TemplateS3Uri = aws.String(_configserviceTemplateS3Uri)
	}
	if len(_configserviceTemplateSSMDocumentDetails) > 0 {
		if err := assignInputField(input, "TemplateSSMDocumentDetails", _configserviceTemplateSSMDocumentDetails); err != nil {
			log.Errorf("invalid --template-ssm-document-details: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutConformancePack(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates a delivery channel to deliver configuration information and
// other compliance information.
//
// You can use this operation to create a new delivery channel or to update the
// Amazon S3 bucket and the Amazon SNS topic of an existing delivery channel.
//
// For more information, see [Working with the Delivery Channel] in the Config Developer Guide.
//
// # One delivery channel per account per Region
//
// You can have only one delivery channel for each account for each Amazon Web
// Services Region.
//
// [Working with the Delivery Channel]: https://docs.aws.amazon.com/config/latest/developerguide/manage-delivery-channel.html
func configservice_PutDeliveryChannel(cfg aws.Config, client *configservice.Client) {
	input := &configservice.PutDeliveryChannelInput{
		// DeliveryChannel: *types.DeliveryChannel, // Required
	}

	if len(_configserviceDeliveryChannel) > 0 {
		if err := assignInputField(input, "DeliveryChannel", _configserviceDeliveryChannel); err != nil {
			log.Errorf("invalid --delivery-channel: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutDeliveryChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Used by an Lambda function to deliver evaluation results to Config. This
// operation is required in every Lambda function that is invoked by an Config
// rule.
func configservice_PutEvaluations(cfg aws.Config, client *configservice.Client) {
	input := &configservice.PutEvaluationsInput{
		// ResultToken: *string, // Required
	}

	if len(_configserviceResultToken) > 0 {
		input.ResultToken = aws.String(_configserviceResultToken)
	}
	if len(_configserviceEvaluations) > 0 {
		if err := assignInputField(input, "Evaluations", _configserviceEvaluations); err != nil {
			log.Errorf("invalid --evaluations: %s", err.Error())
			return
		}
	}
	if len(_configserviceTestMode) > 0 {
		if err := assignInputField(input, "TestMode", _configserviceTestMode); err != nil {
			log.Errorf("invalid --test-mode: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutEvaluations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Add or updates the evaluations for process checks. This API checks if the rule
// is a process check when the name of the Config rule is provided.
func configservice_PutExternalEvaluation(cfg aws.Config, client *configservice.Client) {
	input := &configservice.PutExternalEvaluationInput{
		// ConfigRuleName: *string, // Required
		// ExternalEvaluation: *types.ExternalEvaluation, // Required
	}

	if len(_configserviceConfigRuleName) > 0 {
		input.ConfigRuleName = aws.String(_configserviceConfigRuleName)
	}
	if len(_configserviceExternalEvaluation) > 0 {
		if err := assignInputField(input, "ExternalEvaluation", _configserviceExternalEvaluation); err != nil {
			log.Errorf("invalid --external-evaluation: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutExternalEvaluation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds or updates an Config rule for your entire organization to evaluate if your
// Amazon Web Services resources comply with your desired configurations. For
// information on how many organization Config rules you can have per account, see [Service Limits]
// in the Config Developer Guide.
//
// Only a management account and a delegated administrator can create or update an
// organization Config rule. When calling this API with a delegated administrator,
// you must ensure Organizations ListDelegatedAdministrator permissions are added.
// An organization can have up to 3 delegated administrators.
//
// This API enables organization service access through the EnableAWSServiceAccess
// action and creates a service-linked role
// AWSServiceRoleForConfigMultiAccountSetup in the management or delegated
// administrator account of your organization. The service-linked role is created
// only when the role does not exist in the caller account. Config verifies the
// existence of role with GetRole action.
//
// To use this API with delegated administrator, register a delegated
// administrator by calling Amazon Web Services Organization
// register-delegated-administrator for config-multiaccountsetup.amazonaws.com .
//
// There are two types of rules: Config Managed Rules and Config Custom Rules. You
// can use PutOrganizationConfigRule to create both Config Managed Rules and
// Config Custom Rules.
//
// Config Managed Rules are predefined, customizable rules created by Config. For
// a list of managed rules, see [List of Config Managed Rules]. If you are adding an Config managed rule, you
// must specify the rule's identifier for the RuleIdentifier key.
//
// Config Custom Rules are rules that you create from scratch. There are two ways
// to create Config custom rules: with Lambda functions ([Lambda Developer Guide] ) and with Guard ([Guard GitHub Repository] ), a
// policy-as-code language.
//
// Config custom rules created with Lambda are called Config Custom Lambda Rules
// and Config custom rules created with Guard are called Config Custom Policy
// Rules.
//
// If you are adding a new Config Custom Lambda rule, you first need to create an
// Lambda function in the management account or a delegated administrator that the
// rule invokes to evaluate your resources. You also need to create an IAM role in
// the managed account that can be assumed by the Lambda function. When you use
// PutOrganizationConfigRule to add a Custom Lambda rule to Config, you must
// specify the Amazon Resource Name (ARN) that Lambda assigns to the function.
//
// Prerequisite: Ensure you call EnableAllFeatures API to enable all features in
// an organization.
//
// Make sure to specify one of either OrganizationCustomPolicyRuleMetadata for
// Custom Policy rules, OrganizationCustomRuleMetadata for Custom Lambda rules, or
// OrganizationManagedRuleMetadata for managed rules.
//
// [List of Config Managed Rules]: https://docs.aws.amazon.com/config/latest/developerguide/managed-rules-by-aws-config.html
// [Lambda Developer Guide]: https://docs.aws.amazon.com/config/latest/developerguide/gettingstarted-concepts.html#gettingstarted-concepts-function
// [Service Limits]: https://docs.aws.amazon.com/config/latest/developerguide/configlimits.html
// [Guard GitHub Repository]: https://github.com/aws-cloudformation/cloudformation-guard
func configservice_PutOrganizationConfigRule(cfg aws.Config, client *configservice.Client) {
	input := &configservice.PutOrganizationConfigRuleInput{
		// OrganizationConfigRuleName: *string, // Required
	}

	if len(_configserviceOrganizationConfigRuleName) > 0 {
		input.OrganizationConfigRuleName = aws.String(_configserviceOrganizationConfigRuleName)
	}
	if len(_configserviceExcludedAccounts) > 0 {
		input.ExcludedAccounts = append([]string(nil), _configserviceExcludedAccounts...)
	}
	if len(_configserviceOrganizationCustomPolicyRuleMetadata) > 0 {
		if err := assignInputField(input, "OrganizationCustomPolicyRuleMetadata", _configserviceOrganizationCustomPolicyRuleMetadata); err != nil {
			log.Errorf("invalid --organization-custom-policy-rule-metadata: %s", err.Error())
			return
		}
	}
	if len(_configserviceOrganizationCustomRuleMetadata) > 0 {
		if err := assignInputField(input, "OrganizationCustomRuleMetadata", _configserviceOrganizationCustomRuleMetadata); err != nil {
			log.Errorf("invalid --organization-custom-rule-metadata: %s", err.Error())
			return
		}
	}
	if len(_configserviceOrganizationManagedRuleMetadata) > 0 {
		if err := assignInputField(input, "OrganizationManagedRuleMetadata", _configserviceOrganizationManagedRuleMetadata); err != nil {
			log.Errorf("invalid --organization-managed-rule-metadata: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutOrganizationConfigRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deploys conformance packs across member accounts in an Amazon Web Services
// Organization. For information on how many organization conformance packs and how
// many Config rules you can have per account, see [Service Limits]in the Config Developer Guide.
//
// Only a management account and a delegated administrator can call this API. When
// calling this API with a delegated administrator, you must ensure Organizations
// ListDelegatedAdministrator permissions are added. An organization can have up to
// 3 delegated administrators.
//
// When you use PutOrganizationConformancePack to deploy conformance packs across
// member accounts, the operation can create Config rules and remediation actions
// without requiring config:PutConfigRule or config:PutRemediationConfigurations
// permissions in member account IAM policies.
//
// This API uses the AWSServiceRoleForConfigConforms service-linked role in each
// member account to create conformance pack resources. This service-linked role
// includes the permissions to create Config rules and remediation configurations,
// even if member account IAM policies explicitly deny these actions.
//
// This API enables organization service access for
// config-multiaccountsetup.amazonaws.com through the EnableAWSServiceAccess
// action and creates a service-linked role
// AWSServiceRoleForConfigMultiAccountSetup in the management or delegated
// administrator account of your organization. The service-linked role is created
// only when the role does not exist in the caller account. To use this API with
// delegated administrator, register a delegated administrator by calling Amazon
// Web Services Organization register-delegate-admin for
// config-multiaccountsetup.amazonaws.com .
//
// Prerequisite: Ensure you call EnableAllFeatures API to enable all features in
// an organization.
//
// You must specify either the TemplateS3Uri or the TemplateBody parameter, but
// not both. If you provide both Config uses the TemplateS3Uri parameter and
// ignores the TemplateBody parameter.
//
// Config sets the state of a conformance pack to CREATE_IN_PROGRESS and
// UPDATE_IN_PROGRESS until the conformance pack is created or updated. You cannot
// update a conformance pack while it is in this state.
//
// [Service Limits]: https://docs.aws.amazon.com/config/latest/developerguide/configlimits.html
func configservice_PutOrganizationConformancePack(cfg aws.Config, client *configservice.Client) {
	input := &configservice.PutOrganizationConformancePackInput{
		// OrganizationConformancePackName: *string, // Required
	}

	if len(_configserviceOrganizationConformancePackName) > 0 {
		input.OrganizationConformancePackName = aws.String(_configserviceOrganizationConformancePackName)
	}
	if len(_configserviceConformancePackInputParameters) > 0 {
		if err := assignInputField(input, "ConformancePackInputParameters", _configserviceConformancePackInputParameters); err != nil {
			log.Errorf("invalid --conformance-pack-input-parameters: %s", err.Error())
			return
		}
	}
	if len(_configserviceDeliveryS3Bucket) > 0 {
		input.DeliveryS3Bucket = aws.String(_configserviceDeliveryS3Bucket)
	}
	if len(_configserviceDeliveryS3KeyPrefix) > 0 {
		input.DeliveryS3KeyPrefix = aws.String(_configserviceDeliveryS3KeyPrefix)
	}
	if len(_configserviceExcludedAccounts) > 0 {
		input.ExcludedAccounts = append([]string(nil), _configserviceExcludedAccounts...)
	}
	if len(_configserviceTemplateBody) > 0 {
		input.TemplateBody = aws.String(_configserviceTemplateBody)
	}
	if len(_configserviceTemplateS3Uri) > 0 {
		input.TemplateS3Uri = aws.String(_configserviceTemplateS3Uri)
	}

	if resp, err := client.PutOrganizationConformancePack(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds or updates the remediation configuration with a specific Config rule with
// the selected target or action. The API creates the RemediationConfiguration
// object for the Config rule. The Config rule must already exist for you to add a
// remediation configuration. The target (SSM document) must exist and have
// permissions to use the target.
//
// # Be aware of backward incompatible changes
//
// If you make backward incompatible changes to the SSM document, you must call
// this again to ensure the remediations can run.
//
// This API does not support adding remediation configurations for service-linked
// Config Rules such as Organization Config rules, the rules deployed by
// conformance packs, and rules deployed by Amazon Web Services Security Hub.
//
// # Required fields
//
// For manual remediation configuration, you need to provide a value for
// automationAssumeRole or use a value in the assumeRole field to remediate your
// resources. The SSM automation document can use either as long as it maps to a
// valid parameter.
//
// However, for automatic remediation configuration, the only valid assumeRole
// field value is AutomationAssumeRole and you need to provide a value for
// AutomationAssumeRole to remediate your resources.
//
// # Auto remediation can be initiated even for compliant resources
//
// If you enable auto remediation for a specific Config rule using the [PutRemediationConfigurations] API or the
// Config console, it initiates the remediation process for all non-compliant
// resources for that specific rule. The auto remediation process relies on the
// compliance data snapshot which is captured on a periodic basis. Any
// non-compliant resource that is updated between the snapshot schedule will
// continue to be remediated based on the last known compliance data snapshot.
//
// This means that in some cases auto remediation can be initiated even for
// compliant resources, since the bootstrap processor uses a database that can have
// stale evaluation results based on the last known compliance data snapshot.
//
// [PutRemediationConfigurations]: https://docs.aws.amazon.com/config/latest/APIReference/emAPI_PutRemediationConfigurations.html
func configservice_PutRemediationConfigurations(cfg aws.Config, client *configservice.Client) {
	input := &configservice.PutRemediationConfigurationsInput{
		// RemediationConfigurations: []types.RemediationConfiguration, // Required
	}

	if len(_configserviceRemediationConfigurations) > 0 {
		if err := assignInputField(input, "RemediationConfigurations", _configserviceRemediationConfigurations); err != nil {
			log.Errorf("invalid --remediation-configurations: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutRemediationConfigurations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// A remediation exception is when a specified resource is no longer considered
// for auto-remediation. This API adds a new exception or updates an existing
// exception for a specified resource with a specified Config rule.
//
// # Exceptions block auto remediation
//
// Config generates a remediation exception when a problem occurs running a
// remediation action for a specified resource. Remediation exceptions blocks
// auto-remediation until the exception is cleared.
//
// # Manual remediation is recommended when placing an exception
//
// When placing an exception on an Amazon Web Services resource, it is recommended
// that remediation is set as manual remediation until the given Config rule for
// the specified resource evaluates the resource as NON_COMPLIANT . Once the
// resource has been evaluated as NON_COMPLIANT , you can add remediation
// exceptions and change the remediation type back from Manual to Auto if you want
// to use auto-remediation. Otherwise, using auto-remediation before a
// NON_COMPLIANT evaluation result can delete resources before the exception is
// applied.
//
// # Exceptions can only be performed on non-compliant resources
//
// Placing an exception can only be performed on resources that are NON_COMPLIANT .
// If you use this API for COMPLIANT resources or resources that are NOT_APPLICABLE
// , a remediation exception will not be generated. For more information on the
// conditions that initiate the possible Config evaluation results, see [Concepts | Config Rules]in the
// Config Developer Guide.
//
// # Exceptions cannot be placed on service-linked remediation actions
//
// You cannot place an exception on service-linked remediation actions, such as
// remediation actions put by an organizational conformance pack.
//
// # Auto remediation can be initiated even for compliant resources
//
// If you enable auto remediation for a specific Config rule using the [PutRemediationConfigurations] API or the
// Config console, it initiates the remediation process for all non-compliant
// resources for that specific rule. The auto remediation process relies on the
// compliance data snapshot which is captured on a periodic basis. Any
// non-compliant resource that is updated between the snapshot schedule will
// continue to be remediated based on the last known compliance data snapshot.
//
// This means that in some cases auto remediation can be initiated even for
// compliant resources, since the bootstrap processor uses a database that can have
// stale evaluation results based on the last known compliance data snapshot.
//
// [Concepts | Config Rules]: https://docs.aws.amazon.com/config/latest/developerguide/config-concepts.html#aws-config-rules
// [PutRemediationConfigurations]: https://docs.aws.amazon.com/config/latest/APIReference/emAPI_PutRemediationConfigurations.html
func configservice_PutRemediationExceptions(cfg aws.Config, client *configservice.Client) {
	input := &configservice.PutRemediationExceptionsInput{
		// ConfigRuleName: *string, // Required
		// ResourceKeys: []types.RemediationExceptionResourceKey, // Required
	}

	if len(_configserviceConfigRuleName) > 0 {
		input.ConfigRuleName = aws.String(_configserviceConfigRuleName)
	}
	if len(_configserviceResourceKeys) > 0 {
		if err := assignInputField(input, "ResourceKeys", _configserviceResourceKeys); err != nil {
			log.Errorf("invalid --resource-keys: %s", err.Error())
			return
		}
	}
	if len(_configserviceExpirationTime) > 0 {
		if err := assignInputField(input, "ExpirationTime", _configserviceExpirationTime); err != nil {
			log.Errorf("invalid --expiration-time: %s", err.Error())
			return
		}
	}
	if len(_configserviceMessage) > 0 {
		input.Message = aws.String(_configserviceMessage)
	}

	if resp, err := client.PutRemediationExceptions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Records the configuration state for the resource provided in the request.
// The configuration state of a resource is represented in Config as Configuration
// Items. Once this API records the configuration item, you can retrieve the list
// of configuration items for the custom resource type using existing Config APIs.
//
// The custom resource type must be registered with CloudFormation. This API
// accepts the configuration item registered with CloudFormation.
//
// When you call this API, Config only stores configuration state of the resource
// provided in the request. This API does not change or remediate the configuration
// of the resource.
//
// Write-only schema properites are not recorded as part of the published
// configuration item.
func configservice_PutResourceConfig(cfg aws.Config, client *configservice.Client) {
	input := &configservice.PutResourceConfigInput{
		// Configuration: *string, // Required
		// ResourceId: *string, // Required
		// ResourceType: *string, // Required
		// SchemaVersionId: *string, // Required
	}

	if len(_configserviceConfiguration) > 0 {
		input.Configuration = aws.String(_configserviceConfiguration)
	}
	if len(_configserviceResourceId) > 0 {
		input.ResourceId = aws.String(_configserviceResourceId)
	}
	if len(_configserviceResourceType) > 0 {
		input.ResourceType = aws.String(_configserviceResourceType)
	}
	if len(_configserviceSchemaVersionId) > 0 {
		input.SchemaVersionId = aws.String(_configserviceSchemaVersionId)
	}
	if len(_configserviceResourceName) > 0 {
		input.ResourceName = aws.String(_configserviceResourceName)
	}
	if len(_configserviceTags) > 0 {
		if err := assignInputField(input, "Tags", _configserviceTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutResourceConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates and updates the retention configuration with details about retention
// period (number of days) that Config stores your historical information. The API
// creates the RetentionConfiguration object and names the object as default. When
// you have a RetentionConfiguration object named default, calling the API
// modifies the default object.
//
// Currently, Config supports only one retention configuration per region in your
// account.
func configservice_PutRetentionConfiguration(cfg aws.Config, client *configservice.Client) {
	input := &configservice.PutRetentionConfigurationInput{
		// RetentionPeriodInDays: *int32, // Required
	}

	if len(_configserviceRetentionPeriodInDays) > 0 {
		if err := assignInputField(input, "RetentionPeriodInDays", _configserviceRetentionPeriodInDays); err != nil {
			log.Errorf("invalid --retention-period-in-days: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutRetentionConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a service-linked configuration recorder that is linked to a specific
// Amazon Web Services service based on the ServicePrincipal you specify.
//
// The configuration recorder's name , recordingGroup , recordingMode , and
// recordingScope is set by the service that is linked to the configuration
// recorder.
//
// For more information and a list of supported services/service principals, see [Working with the Configuration Recorder]
// in the Config Developer Guide.
//
// This API creates a service-linked role AWSServiceRoleForConfig in your account.
// The service-linked role is created only when the role does not exist in your
// account.
//
// # The recording scope determines if you receive configuration items
//
// The recording scope is set by the service that is linked to the configuration
// recorder and determines whether you receive configuration items (CIs) in the
// delivery channel. If the recording scope is internal, you will not receive CIs
// in the delivery channel.
//
// # Tags are added at creation and cannot be updated with this operation
//
// Use [TagResource] and [UntagResource] to update tags after creation.
//
// [Working with the Configuration Recorder]: https://docs.aws.amazon.com/config/latest/developerguide/stop-start-recorder.html
// [TagResource]: https://docs.aws.amazon.com/config/latest/APIReference/API_TagResource.html
// [UntagResource]: https://docs.aws.amazon.com/config/latest/APIReference/API_UntagResource.html
func configservice_PutServiceLinkedConfigurationRecorder(cfg aws.Config, client *configservice.Client) {
	input := &configservice.PutServiceLinkedConfigurationRecorderInput{
		// ServicePrincipal: *string, // Required
	}

	if len(_configserviceServicePrincipal) > 0 {
		input.ServicePrincipal = aws.String(_configserviceServicePrincipal)
	}
	if len(_configserviceTags) > 0 {
		if err := assignInputField(input, "Tags", _configserviceTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutServiceLinkedConfigurationRecorder(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Saves a new query or updates an existing saved query. The QueryName must be
// unique for a single Amazon Web Services account and a single Amazon Web Services
// Region. You can create upto 300 queries in a single Amazon Web Services account
// and a single Amazon Web Services Region.
//
// # Tags are added at creation and cannot be updated
//
// PutStoredQuery is an idempotent API. Subsequent requests won’t create a
// duplicate resource if one was already created. If a following request has
// different tags values, Config will ignore these differences and treat it as an
// idempotent request of the previous. In this case, tags will not be updated,
// even if they are different.
func configservice_PutStoredQuery(cfg aws.Config, client *configservice.Client) {
	input := &configservice.PutStoredQueryInput{
		// StoredQuery: *types.StoredQuery, // Required
	}

	if len(_configserviceStoredQuery) > 0 {
		if err := assignInputField(input, "StoredQuery", _configserviceStoredQuery); err != nil {
			log.Errorf("invalid --stored-query: %s", err.Error())
			return
		}
	}
	if len(_configserviceTags) > 0 {
		if err := assignInputField(input, "Tags", _configserviceTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutStoredQuery(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Accepts a structured query language (SQL) SELECT command and an aggregator to
// query configuration state of Amazon Web Services resources across multiple
// accounts and regions, performs the corresponding search, and returns resource
// configurations matching the properties.
//
// For more information about query components, see the [Query Components] section in the Config
// Developer Guide.
//
// If you run an aggregation query (i.e., using GROUP BY or using aggregate
// functions such as COUNT ; e.g., SELECT resourceId, COUNT(*) WHERE resourceType
// = 'AWS::IAM::Role' GROUP BY resourceId ) and do not specify the MaxResults or
// the Limit query parameters, the default page size is set to 500.
//
// If you run a non-aggregation query (i.e., not using GROUP BY or aggregate
// function; e.g., SELECT * WHERE resourceType = 'AWS::IAM::Role' ) and do not
// specify the MaxResults or the Limit query parameters, the default page size is
// set to 25.
//
// [Query Components]: https://docs.aws.amazon.com/config/latest/developerguide/query-components.html
func configservice_SelectAggregateResourceConfig(cfg aws.Config, client *configservice.Client) {
	input := &configservice.SelectAggregateResourceConfigInput{
		// ConfigurationAggregatorName: *string, // Required
		// Expression: *string, // Required
	}

	if len(_configserviceConfigurationAggregatorName) > 0 {
		input.ConfigurationAggregatorName = aws.String(_configserviceConfigurationAggregatorName)
	}
	if len(_configserviceExpression) > 0 {
		input.Expression = aws.String(_configserviceExpression)
	}
	if len(_configserviceLimit) > 0 {
		if err := assignInputField(input, "Limit", _configserviceLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_configserviceMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _configserviceMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_configserviceNextToken) > 0 {
		input.NextToken = aws.String(_configserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.SelectAggregateResourceConfig(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*configservice.SelectAggregateResourceConfigOutput
	p := configservice.NewSelectAggregateResourceConfigPaginator(client, input)
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

// Accepts a structured query language (SQL) SELECT command, performs the
// corresponding search, and returns resource configurations matching the
// properties.
//
// For more information about query components, see the [Query Components] section in the Config
// Developer Guide.
//
// [Query Components]: https://docs.aws.amazon.com/config/latest/developerguide/query-components.html
func configservice_SelectResourceConfig(cfg aws.Config, client *configservice.Client) {
	input := &configservice.SelectResourceConfigInput{
		// Expression: *string, // Required
	}

	if len(_configserviceExpression) > 0 {
		input.Expression = aws.String(_configserviceExpression)
	}
	if len(_configserviceLimit) > 0 {
		if err := assignInputField(input, "Limit", _configserviceLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_configserviceNextToken) > 0 {
		input.NextToken = aws.String(_configserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.SelectResourceConfig(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*configservice.SelectResourceConfigOutput
	p := configservice.NewSelectResourceConfigPaginator(client, input)
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

// Runs an on-demand evaluation for the specified Config rules against the last
// known configuration state of the resources. Use StartConfigRulesEvaluation when
// you want to test that a rule you updated is working as expected.
// StartConfigRulesEvaluation does not re-record the latest configuration state for
// your resources. It re-runs an evaluation against the last known state of your
// resources.
//
// You can specify up to 25 Config rules per request.
//
// An existing StartConfigRulesEvaluation call for the specified rules must
// complete before you can call the API again. If you chose to have Config stream
// to an Amazon SNS topic, you will receive a ConfigRuleEvaluationStarted
// notification when the evaluation starts.
//
// You don't need to call the StartConfigRulesEvaluation API to run an evaluation
// for a new rule. When you create a rule, Config evaluates your resources against
// the rule automatically.
//
// The StartConfigRulesEvaluation API is useful if you want to run on-demand
// evaluations, such as the following example:
//
// - You have a custom rule that evaluates your IAM resources every 24 hours.
//
// - You update your Lambda function to add additional conditions to your rule.
//
// - Instead of waiting for the next periodic evaluation, you call the
// StartConfigRulesEvaluation API.
//
// - Config invokes your Lambda function and evaluates your IAM resources.
//
// - Your custom rule will still run periodic evaluations every 24 hours.
func configservice_StartConfigRulesEvaluation(cfg aws.Config, client *configservice.Client) {
	input := &configservice.StartConfigRulesEvaluationInput{}

	if len(_configserviceConfigRuleNames) > 0 {
		input.ConfigRuleNames = append([]string(nil), _configserviceConfigRuleNames...)
	}

	if resp, err := client.StartConfigRulesEvaluation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts the customer managed configuration recorder. The customer managed
// configuration recorder will begin recording configuration changes for the
// resource types you specify.
//
// You must have created a delivery channel to successfully start the customer
// managed configuration recorder. You can use the [PutDeliveryChannel]operation to create a delivery
// channel.
//
// [PutDeliveryChannel]: https://docs.aws.amazon.com/config/latest/APIReference/API_PutDeliveryChannel.html
func configservice_StartConfigurationRecorder(cfg aws.Config, client *configservice.Client) {
	input := &configservice.StartConfigurationRecorderInput{
		// ConfigurationRecorderName: *string, // Required
	}

	if len(_configserviceConfigurationRecorderName) > 0 {
		input.ConfigurationRecorderName = aws.String(_configserviceConfigurationRecorderName)
	}

	if resp, err := client.StartConfigurationRecorder(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Runs an on-demand remediation for the specified Config rules against the last
// known remediation configuration. It runs an execution against the current state
// of your resources. Remediation execution is asynchronous.
//
// You can specify up to 100 resource keys per request. An existing
// StartRemediationExecution call for the specified resource keys must complete
// before you can call the API again.
func configservice_StartRemediationExecution(cfg aws.Config, client *configservice.Client) {
	input := &configservice.StartRemediationExecutionInput{
		// ConfigRuleName: *string, // Required
		// ResourceKeys: []types.ResourceKey, // Required
	}

	if len(_configserviceConfigRuleName) > 0 {
		input.ConfigRuleName = aws.String(_configserviceConfigRuleName)
	}
	if len(_configserviceResourceKeys) > 0 {
		if err := assignInputField(input, "ResourceKeys", _configserviceResourceKeys); err != nil {
			log.Errorf("invalid --resource-keys: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartRemediationExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Runs an on-demand evaluation for the specified resource to determine whether
// the resource details will comply with configured Config rules. You can also use
// it for evaluation purposes. Config recommends using an evaluation context. It
// runs an execution against the resource details with all of the Config rules in
// your account that match with the specified proactive mode and resource type.
//
// Ensure you have the cloudformation:DescribeType role setup to validate the
// resource type schema.
//
// You can find the [Resource type schema] in "Amazon Web Services public extensions" within the
// CloudFormation registry or with the following CLI commmand: aws cloudformation
// describe-type --type-name "AWS::S3::Bucket" --type RESOURCE .
//
// For more information, see [Managing extensions through the CloudFormation registry] and [Amazon Web Services resource and property types reference] in the CloudFormation User Guide.
//
// [Resource type schema]: https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-type-schema.html
// [Amazon Web Services resource and property types reference]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-template-resource-type-ref.html
// [Managing extensions through the CloudFormation registry]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/registry.html#registry-view
func configservice_StartResourceEvaluation(cfg aws.Config, client *configservice.Client) {
	input := &configservice.StartResourceEvaluationInput{
		// EvaluationMode: types.EvaluationMode, // Required
		// ResourceDetails: *types.ResourceDetails, // Required
	}

	if len(_configserviceEvaluationMode) > 0 {
		if err := assignInputField(input, "EvaluationMode", _configserviceEvaluationMode); err != nil {
			log.Errorf("invalid --evaluation-mode: %s", err.Error())
			return
		}
	}
	if len(_configserviceResourceDetails) > 0 {
		if err := assignInputField(input, "ResourceDetails", _configserviceResourceDetails); err != nil {
			log.Errorf("invalid --resource-details: %s", err.Error())
			return
		}
	}
	if len(_configserviceClientToken) > 0 {
		input.ClientToken = aws.String(_configserviceClientToken)
	}
	if len(_configserviceEvaluationContext) > 0 {
		if err := assignInputField(input, "EvaluationContext", _configserviceEvaluationContext); err != nil {
			log.Errorf("invalid --evaluation-context: %s", err.Error())
			return
		}
	}
	if len(_configserviceEvaluationTimeout) > 0 {
		if err := assignInputField(input, "EvaluationTimeout", _configserviceEvaluationTimeout); err != nil {
			log.Errorf("invalid --evaluation-timeout: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartResourceEvaluation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops the customer managed configuration recorder. The customer managed
// configuration recorder will stop recording configuration changes for the
// resource types you have specified.
func configservice_StopConfigurationRecorder(cfg aws.Config, client *configservice.Client) {
	input := &configservice.StopConfigurationRecorderInput{
		// ConfigurationRecorderName: *string, // Required
	}

	if len(_configserviceConfigurationRecorderName) > 0 {
		input.ConfigurationRecorderName = aws.String(_configserviceConfigurationRecorderName)
	}

	if resp, err := client.StopConfigurationRecorder(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates the specified tags to a resource with the specified ResourceArn . If
// existing tags on a resource are not specified in the request parameters, they
// are not changed. If existing tags are specified, however, then their values will
// be updated. When a resource is deleted, the tags associated with that resource
// are deleted as well.
func configservice_TagResource(cfg aws.Config, client *configservice.Client) {
	input := &configservice.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_configserviceResourceArn) > 0 {
		input.ResourceArn = aws.String(_configserviceResourceArn)
	}
	if len(_configserviceTags) > 0 {
		if err := assignInputField(input, "Tags", _configserviceTags); err != nil {
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

// Deletes specified tags from a resource.
func configservice_UntagResource(cfg aws.Config, client *configservice.Client) {
	input := &configservice.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_configserviceResourceArn) > 0 {
		input.ResourceArn = aws.String(_configserviceResourceArn)
	}
	if len(_configserviceTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _configserviceTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_configserviceCmd)
	_configserviceCmd.Flags().SortFlags = false

	_configserviceCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_configserviceCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_configserviceCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_configserviceCmd.Flags().StringVarP(&_configserviceAccountAggregationSources, "account-aggregation-sources", "", "", "Account Aggregation Sources")
	_configserviceCmd.Flags().StringVarP(&_configserviceAccountId, "account-id", "", "", "Account ID")
	_configserviceCmd.Flags().StringVarP(&_configserviceAggregatorFilters, "aggregator-filters", "", "", "Aggregator Filters")
	_configserviceCmd.Flags().StringVarP(&_configserviceArn, "arn", "", "", "ARN")
	_configserviceCmd.Flags().StringVarP(&_configserviceAuthorizedAccountId, "authorized-account-id", "", "", "Authorized Account ID")
	_configserviceCmd.Flags().StringVarP(&_configserviceAuthorizedAwsRegion, "authorized-aws-region", "", "", "Authorized AWS Region")
	_configserviceCmd.Flags().StringVarP(&_configserviceAwsRegion, "aws-region", "", "", "AWS Region")
	_configserviceCmd.Flags().StringVarP(&_configserviceChronologicalOrder, "chronological-order", "", "", "Chronological Order")
	_configserviceCmd.Flags().StringVarP(&_configserviceClientToken, "client-token", "", "", "Client Token")
	_configserviceCmd.Flags().StringVarP(&_configserviceComplianceType, "compliance-type", "", "", "Compliance Type")
	_configserviceCmd.Flags().StringVarP(&_configserviceComplianceTypes, "compliance-types", "", "", "Compliance Types")
	_configserviceCmd.Flags().StringVarP(&_configserviceConfigRule, "config-rule", "", "", "Config Rule")
	_configserviceCmd.Flags().StringVarP(&_configserviceConfigRuleName, "config-rule-name", "", "", "Config Rule Name")
	_configserviceCmd.Flags().StringSliceVarP(&_configserviceConfigRuleNames, "config-rule-names", "", nil, "Config Rule Names")
	_configserviceCmd.Flags().StringVarP(&_configserviceConfiguration, "configuration", "", "", "Configuration")
	_configserviceCmd.Flags().StringVarP(&_configserviceConfigurationAggregatorName, "configuration-aggregator-name", "", "", "Configuration Aggregator Name")
	_configserviceCmd.Flags().StringSliceVarP(&_configserviceConfigurationAggregatorNames, "configuration-aggregator-names", "", nil, "Configuration Aggregator Names")
	_configserviceCmd.Flags().StringVarP(&_configserviceConfigurationRecorder, "configuration-recorder", "", "", "Configuration Recorder")
	_configserviceCmd.Flags().StringVarP(&_configserviceConfigurationRecorderArn, "configuration-recorder-arn", "", "", "Configuration Recorder ARN")
	_configserviceCmd.Flags().StringVarP(&_configserviceConfigurationRecorderName, "configuration-recorder-name", "", "", "Configuration Recorder Name")
	_configserviceCmd.Flags().StringSliceVarP(&_configserviceConfigurationRecorderNames, "configuration-recorder-names", "", nil, "Configuration Recorder Names")
	_configserviceCmd.Flags().StringVarP(&_configserviceConformancePackInputParameters, "conformance-pack-input-parameters", "", "", "Conformance Pack Input Parameters")
	_configserviceCmd.Flags().StringVarP(&_configserviceConformancePackName, "conformance-pack-name", "", "", "Conformance Pack Name")
	_configserviceCmd.Flags().StringSliceVarP(&_configserviceConformancePackNames, "conformance-pack-names", "", nil, "Conformance Pack Names")
	_configserviceCmd.Flags().StringVarP(&_configserviceDeliveryChannel, "delivery-channel", "", "", "Delivery Channel")
	_configserviceCmd.Flags().StringVarP(&_configserviceDeliveryChannelName, "delivery-channel-name", "", "", "Delivery Channel Name")
	_configserviceCmd.Flags().StringSliceVarP(&_configserviceDeliveryChannelNames, "delivery-channel-names", "", nil, "Delivery Channel Names")
	_configserviceCmd.Flags().StringVarP(&_configserviceDeliveryS3Bucket, "delivery-s3-bucket", "", "", "Delivery S3 Bucket")
	_configserviceCmd.Flags().StringVarP(&_configserviceDeliveryS3KeyPrefix, "delivery-s3-key-prefix", "", "", "Delivery S3 Key Prefix")
	_configserviceCmd.Flags().StringVarP(&_configserviceEarlierTime, "earlier-time", "", "", "Earlier Time")
	_configserviceCmd.Flags().StringVarP(&_configserviceEvaluationContext, "evaluation-context", "", "", "Evaluation Context")
	_configserviceCmd.Flags().StringVarP(&_configserviceEvaluationMode, "evaluation-mode", "", "", "Evaluation Mode")
	_configserviceCmd.Flags().StringVarP(&_configserviceEvaluationTimeout, "evaluation-timeout", "", "", "Evaluation Timeout")
	_configserviceCmd.Flags().StringVarP(&_configserviceEvaluations, "evaluations", "", "", "Evaluations")
	_configserviceCmd.Flags().StringSliceVarP(&_configserviceExcludedAccounts, "excluded-accounts", "", nil, "Excluded Accounts")
	_configserviceCmd.Flags().StringVarP(&_configserviceExpirationTime, "expiration-time", "", "", "Expiration Time")
	_configserviceCmd.Flags().StringVarP(&_configserviceExpression, "expression", "", "", "Expression")
	_configserviceCmd.Flags().StringVarP(&_configserviceExternalEvaluation, "external-evaluation", "", "", "External Evaluation")
	_configserviceCmd.Flags().StringVarP(&_configserviceFilters, "filters", "", "", "Filters")
	_configserviceCmd.Flags().StringVarP(&_configserviceGroupByKey, "group-by-key", "", "", "Group By Key")
	_configserviceCmd.Flags().StringVarP(&_configserviceIncludeDeletedResources, "include-deleted-resources", "", "", "Include Deleted Resources")
	_configserviceCmd.Flags().StringVarP(&_configserviceLaterTime, "later-time", "", "", "Later Time")
	_configserviceCmd.Flags().StringVarP(&_configserviceLimit, "limit", "", "", "Limit")
	_configserviceCmd.Flags().StringVarP(&_configserviceMaxResults, "max-results", "", "", "Max Results")
	_configserviceCmd.Flags().StringVarP(&_configserviceMessage, "message", "", "", "Message")
	_configserviceCmd.Flags().StringVarP(&_configserviceNextToken, "next-token", "", "", "Next Token")
	_configserviceCmd.Flags().StringVarP(&_configserviceOrganizationAggregationSource, "organization-aggregation-source", "", "", "Organization Aggregation Source")
	_configserviceCmd.Flags().StringVarP(&_configserviceOrganizationConfigRuleName, "organization-config-rule-name", "", "", "Organization Config Rule Name")
	_configserviceCmd.Flags().StringSliceVarP(&_configserviceOrganizationConfigRuleNames, "organization-config-rule-names", "", nil, "Organization Config Rule Names")
	_configserviceCmd.Flags().StringVarP(&_configserviceOrganizationConformancePackName, "organization-conformance-pack-name", "", "", "Organization Conformance Pack Name")
	_configserviceCmd.Flags().StringSliceVarP(&_configserviceOrganizationConformancePackNames, "organization-conformance-pack-names", "", nil, "Organization Conformance Pack Names")
	_configserviceCmd.Flags().StringVarP(&_configserviceOrganizationCustomPolicyRuleMetadata, "organization-custom-policy-rule-metadata", "", "", "Organization Custom Policy Rule Metadata")
	_configserviceCmd.Flags().StringVarP(&_configserviceOrganizationCustomRuleMetadata, "organization-custom-rule-metadata", "", "", "Organization Custom Rule Metadata")
	_configserviceCmd.Flags().StringVarP(&_configserviceOrganizationManagedRuleMetadata, "organization-managed-rule-metadata", "", "", "Organization Managed Rule Metadata")
	_configserviceCmd.Flags().StringVarP(&_configserviceQueryName, "query-name", "", "", "Query Name")
	_configserviceCmd.Flags().StringVarP(&_configserviceRemediationConfigurations, "remediation-configurations", "", "", "Remediation Configurations")
	_configserviceCmd.Flags().StringVarP(&_configserviceRequesterAccountId, "requester-account-id", "", "", "Requester Account ID")
	_configserviceCmd.Flags().StringVarP(&_configserviceRequesterAwsRegion, "requester-aws-region", "", "", "Requester AWS Region")
	_configserviceCmd.Flags().StringVarP(&_configserviceResourceArn, "resource-arn", "", "", "Resource ARN")
	_configserviceCmd.Flags().StringVarP(&_configserviceResourceDetails, "resource-details", "", "", "Resource Details")
	_configserviceCmd.Flags().StringVarP(&_configserviceResourceEvaluationId, "resource-evaluation-id", "", "", "Resource Evaluation ID")
	_configserviceCmd.Flags().StringVarP(&_configserviceResourceId, "resource-id", "", "", "Resource ID")
	_configserviceCmd.Flags().StringVarP(&_configserviceResourceIdentifier, "resource-identifier", "", "", "Resource Identifier")
	_configserviceCmd.Flags().StringVarP(&_configserviceResourceIdentifiers, "resource-identifiers", "", "", "Resource Identifiers")
	_configserviceCmd.Flags().StringSliceVarP(&_configserviceResourceIds, "resource-ids", "", nil, "Resource Ids")
	_configserviceCmd.Flags().StringVarP(&_configserviceResourceKeys, "resource-keys", "", "", "Resource Keys")
	_configserviceCmd.Flags().StringVarP(&_configserviceResourceName, "resource-name", "", "", "Resource Name")
	_configserviceCmd.Flags().StringVarP(&_configserviceResourceType, "resource-type", "", "", "Resource Type")
	_configserviceCmd.Flags().StringSliceVarP(&_configserviceResourceTypes, "resource-types", "", nil, "Resource Types")
	_configserviceCmd.Flags().StringVarP(&_configserviceResultToken, "result-token", "", "", "Result Token")
	_configserviceCmd.Flags().StringVarP(&_configserviceRetentionConfigurationName, "retention-configuration-name", "", "", "Retention Configuration Name")
	_configserviceCmd.Flags().StringSliceVarP(&_configserviceRetentionConfigurationNames, "retention-configuration-names", "", nil, "Retention Configuration Names")
	_configserviceCmd.Flags().StringVarP(&_configserviceRetentionPeriodInDays, "retention-period-in-days", "", "", "Retention Period In Days")
	_configserviceCmd.Flags().StringVarP(&_configserviceSchemaVersionId, "schema-version-id", "", "", "Schema Version ID")
	_configserviceCmd.Flags().StringVarP(&_configserviceServicePrincipal, "service-principal", "", "", "Service Principal")
	_configserviceCmd.Flags().StringVarP(&_configserviceSortBy, "sort-by", "", "", "Sort By")
	_configserviceCmd.Flags().StringVarP(&_configserviceSortOrder, "sort-order", "", "", "Sort Order")
	_configserviceCmd.Flags().StringVarP(&_configserviceStoredQuery, "stored-query", "", "", "Stored Query")
	_configserviceCmd.Flags().StringSliceVarP(&_configserviceTagKeys, "tag-keys", "", nil, "Tag Keys")
	_configserviceCmd.Flags().StringVarP(&_configserviceTags, "tags", "", "", "Tags")
	_configserviceCmd.Flags().StringVarP(&_configserviceTemplateBody, "template-body", "", "", "Template Body")
	_configserviceCmd.Flags().StringVarP(&_configserviceTemplateS3Uri, "template-s3-uri", "", "", "Template S3 URI")
	_configserviceCmd.Flags().StringVarP(&_configserviceTemplateSSMDocumentDetails, "template-ssm-document-details", "", "", "Template Ssm Document Details")
	_configserviceCmd.Flags().StringVarP(&_configserviceTestMode, "test-mode", "", "", "Test Mode")
	_configserviceCmd.Flags().StringVarP(&_configserviceUpdateStatus, "update-status", "", "", "Update Status")

	_configserviceCmd.Flags().BoolVarP(&_configserviceAssociateResourceTypes, "associate-resource-types", "", false, "Associate Resource Types")
	_configserviceCmd.Flags().BoolVarP(&_configserviceBatchGetAggregateResourceConfig, "batch-get-aggregate-resource-config", "", false, "Batch Get Aggregate Resource Config")
	_configserviceCmd.Flags().BoolVarP(&_configserviceBatchGetResourceConfig, "batch-get-resource-config", "", false, "Batch Get Resource Config")
	_configserviceCmd.Flags().BoolVarP(&_configserviceDeleteAggregationAuthorization, "delete-aggregation-authorization", "", false, "Delete Aggregation Authorization")
	_configserviceCmd.Flags().BoolVarP(&_configserviceDeleteConfigRule, "delete-config-rule", "", false, "Delete Config Rule")
	_configserviceCmd.Flags().BoolVarP(&_configserviceDeleteConfigurationAggregator, "delete-configuration-aggregator", "", false, "Delete Configuration Aggregator")
	_configserviceCmd.Flags().BoolVarP(&_configserviceDeleteConfigurationRecorder, "delete-configuration-recorder", "", false, "Delete Configuration Recorder")
	_configserviceCmd.Flags().BoolVarP(&_configserviceDeleteConformancePack, "delete-conformance-pack", "", false, "Delete Conformance Pack")
	_configserviceCmd.Flags().BoolVarP(&_configserviceDeleteDeliveryChannel, "delete-delivery-channel", "", false, "Delete Delivery Channel")
	_configserviceCmd.Flags().BoolVarP(&_configserviceDeleteEvaluationResults, "delete-evaluation-results", "", false, "Delete Evaluation Results")
	_configserviceCmd.Flags().BoolVarP(&_configserviceDeleteOrganizationConfigRule, "delete-organization-config-rule", "", false, "Delete Organization Config Rule")
	_configserviceCmd.Flags().BoolVarP(&_configserviceDeleteOrganizationConformancePack, "delete-organization-conformance-pack", "", false, "Delete Organization Conformance Pack")
	_configserviceCmd.Flags().BoolVarP(&_configserviceDeletePendingAggregationRequest, "delete-pending-aggregation-request", "", false, "Delete Pending Aggregation Request")
	_configserviceCmd.Flags().BoolVarP(&_configserviceDeleteRemediationConfiguration, "delete-remediation-configuration", "", false, "Delete Remediation Configuration")
	_configserviceCmd.Flags().BoolVarP(&_configserviceDeleteRemediationExceptions, "delete-remediation-exceptions", "", false, "Delete Remediation Exceptions")
	_configserviceCmd.Flags().BoolVarP(&_configserviceDeleteResourceConfig, "delete-resource-config", "", false, "Delete Resource Config")
	_configserviceCmd.Flags().BoolVarP(&_configserviceDeleteRetentionConfiguration, "delete-retention-configuration", "", false, "Delete Retention Configuration")
	_configserviceCmd.Flags().BoolVarP(&_configserviceDeleteServiceLinkedConfigurationRecorder, "delete-service-linked-configuration-recorder", "", false, "Delete Service Linked Configuration Recorder")
	_configserviceCmd.Flags().BoolVarP(&_configserviceDeleteStoredQuery, "delete-stored-query", "", false, "Delete Stored Query")
	_configserviceCmd.Flags().BoolVarP(&_configserviceDeliverConfigSnapshot, "deliver-config-snapshot", "", false, "Deliver Config Snapshot")
	_configserviceCmd.Flags().BoolVarP(&_configserviceDescribeAggregateComplianceByConfigRules, "describe-aggregate-compliance-by-config-rules", "", false, "Describe Aggregate Compliance By Config Rules")
	_configserviceCmd.Flags().BoolVarP(&_configserviceDescribeAggregateComplianceByConformancePacks, "describe-aggregate-compliance-by-conformance-packs", "", false, "Describe Aggregate Compliance By Conformance Packs")
	_configserviceCmd.Flags().BoolVarP(&_configserviceDescribeAggregationAuthorizations, "describe-aggregation-authorizations", "", false, "Describe Aggregation Authorizations")
	_configserviceCmd.Flags().BoolVarP(&_configserviceDescribeComplianceByConfigRule, "describe-compliance-by-config-rule", "", false, "Describe Compliance By Config Rule")
	_configserviceCmd.Flags().BoolVarP(&_configserviceDescribeComplianceByResource, "describe-compliance-by-resource", "", false, "Describe Compliance By Resource")
	_configserviceCmd.Flags().BoolVarP(&_configserviceDescribeConfigRuleEvaluationStatus, "describe-config-rule-evaluation-status", "", false, "Describe Config Rule Evaluation Status")
	_configserviceCmd.Flags().BoolVarP(&_configserviceDescribeConfigRules, "describe-config-rules", "", false, "Describe Config Rules")
	_configserviceCmd.Flags().BoolVarP(&_configserviceDescribeConfigurationAggregatorSourcesStatus, "describe-configuration-aggregator-sources-status", "", false, "Describe Configuration Aggregator Sources Status")
	_configserviceCmd.Flags().BoolVarP(&_configserviceDescribeConfigurationAggregators, "describe-configuration-aggregators", "", false, "Describe Configuration Aggregators")
	_configserviceCmd.Flags().BoolVarP(&_configserviceDescribeConfigurationRecorderStatus, "describe-configuration-recorder-status", "", false, "Describe Configuration Recorder Status")
	_configserviceCmd.Flags().BoolVarP(&_configserviceDescribeConfigurationRecorders, "describe-configuration-recorders", "", false, "Describe Configuration Recorders")
	_configserviceCmd.Flags().BoolVarP(&_configserviceDescribeConformancePackCompliance, "describe-conformance-pack-compliance", "", false, "Describe Conformance Pack Compliance")
	_configserviceCmd.Flags().BoolVarP(&_configserviceDescribeConformancePackStatus, "describe-conformance-pack-status", "", false, "Describe Conformance Pack Status")
	_configserviceCmd.Flags().BoolVarP(&_configserviceDescribeConformancePacks, "describe-conformance-packs", "", false, "Describe Conformance Packs")
	_configserviceCmd.Flags().BoolVarP(&_configserviceDescribeDeliveryChannelStatus, "describe-delivery-channel-status", "", false, "Describe Delivery Channel Status")
	_configserviceCmd.Flags().BoolVarP(&_configserviceDescribeDeliveryChannels, "describe-delivery-channels", "", false, "Describe Delivery Channels")
	_configserviceCmd.Flags().BoolVarP(&_configserviceDescribeOrganizationConfigRuleStatuses, "describe-organization-config-rule-statuses", "", false, "Describe Organization Config Rule Statuses")
	_configserviceCmd.Flags().BoolVarP(&_configserviceDescribeOrganizationConfigRules, "describe-organization-config-rules", "", false, "Describe Organization Config Rules")
	_configserviceCmd.Flags().BoolVarP(&_configserviceDescribeOrganizationConformancePackStatuses, "describe-organization-conformance-pack-statuses", "", false, "Describe Organization Conformance Pack Statuses")
	_configserviceCmd.Flags().BoolVarP(&_configserviceDescribeOrganizationConformancePacks, "describe-organization-conformance-packs", "", false, "Describe Organization Conformance Packs")
	_configserviceCmd.Flags().BoolVarP(&_configserviceDescribePendingAggregationRequests, "describe-pending-aggregation-requests", "", false, "Describe Pending Aggregation Requests")
	_configserviceCmd.Flags().BoolVarP(&_configserviceDescribeRemediationConfigurations, "describe-remediation-configurations", "", false, "Describe Remediation Configurations")
	_configserviceCmd.Flags().BoolVarP(&_configserviceDescribeRemediationExceptions, "describe-remediation-exceptions", "", false, "Describe Remediation Exceptions")
	_configserviceCmd.Flags().BoolVarP(&_configserviceDescribeRemediationExecutionStatus, "describe-remediation-execution-status", "", false, "Describe Remediation Execution Status")
	_configserviceCmd.Flags().BoolVarP(&_configserviceDescribeRetentionConfigurations, "describe-retention-configurations", "", false, "Describe Retention Configurations")
	_configserviceCmd.Flags().BoolVarP(&_configserviceDisassociateResourceTypes, "disassociate-resource-types", "", false, "Disassociate Resource Types")
	_configserviceCmd.Flags().BoolVarP(&_configserviceGetAggregateComplianceDetailsByConfigRule, "get-aggregate-compliance-details-by-config-rule", "", false, "Get Aggregate Compliance Details By Config Rule")
	_configserviceCmd.Flags().BoolVarP(&_configserviceGetAggregateConfigRuleComplianceSummary, "get-aggregate-config-rule-compliance-summary", "", false, "Get Aggregate Config Rule Compliance Summary")
	_configserviceCmd.Flags().BoolVarP(&_configserviceGetAggregateConformancePackComplianceSummary, "get-aggregate-conformance-pack-compliance-summary", "", false, "Get Aggregate Conformance Pack Compliance Summary")
	_configserviceCmd.Flags().BoolVarP(&_configserviceGetAggregateDiscoveredResourceCounts, "get-aggregate-discovered-resource-counts", "", false, "Get Aggregate Discovered Resource Counts")
	_configserviceCmd.Flags().BoolVarP(&_configserviceGetAggregateResourceConfig, "get-aggregate-resource-config", "", false, "Get Aggregate Resource Config")
	_configserviceCmd.Flags().BoolVarP(&_configserviceGetComplianceDetailsByConfigRule, "get-compliance-details-by-config-rule", "", false, "Get Compliance Details By Config Rule")
	_configserviceCmd.Flags().BoolVarP(&_configserviceGetComplianceDetailsByResource, "get-compliance-details-by-resource", "", false, "Get Compliance Details By Resource")
	_configserviceCmd.Flags().BoolVarP(&_configserviceGetComplianceSummaryByConfigRule, "get-compliance-summary-by-config-rule", "", false, "Get Compliance Summary By Config Rule")
	_configserviceCmd.Flags().BoolVarP(&_configserviceGetComplianceSummaryByResourceType, "get-compliance-summary-by-resource-type", "", false, "Get Compliance Summary By Resource Type")
	_configserviceCmd.Flags().BoolVarP(&_configserviceGetConformancePackComplianceDetails, "get-conformance-pack-compliance-details", "", false, "Get Conformance Pack Compliance Details")
	_configserviceCmd.Flags().BoolVarP(&_configserviceGetConformancePackComplianceSummary, "get-conformance-pack-compliance-summary", "", false, "Get Conformance Pack Compliance Summary")
	_configserviceCmd.Flags().BoolVarP(&_configserviceGetCustomRulePolicy, "get-custom-rule-policy", "", false, "Get Custom Rule Policy")
	_configserviceCmd.Flags().BoolVarP(&_configserviceGetDiscoveredResourceCounts, "get-discovered-resource-counts", "", false, "Get Discovered Resource Counts")
	_configserviceCmd.Flags().BoolVarP(&_configserviceGetOrganizationConfigRuleDetailedStatus, "get-organization-config-rule-detailed-status", "", false, "Get Organization Config Rule Detailed Status")
	_configserviceCmd.Flags().BoolVarP(&_configserviceGetOrganizationConformancePackDetailedStatus, "get-organization-conformance-pack-detailed-status", "", false, "Get Organization Conformance Pack Detailed Status")
	_configserviceCmd.Flags().BoolVarP(&_configserviceGetOrganizationCustomRulePolicy, "get-organization-custom-rule-policy", "", false, "Get Organization Custom Rule Policy")
	_configserviceCmd.Flags().BoolVarP(&_configserviceGetResourceConfigHistory, "get-resource-config-history", "", false, "Get Resource Config History")
	_configserviceCmd.Flags().BoolVarP(&_configserviceGetResourceEvaluationSummary, "get-resource-evaluation-summary", "", false, "Get Resource Evaluation Summary")
	_configserviceCmd.Flags().BoolVarP(&_configserviceGetStoredQuery, "get-stored-query", "", false, "Get Stored Query")
	_configserviceCmd.Flags().BoolVarP(&_configserviceListAggregateDiscoveredResources, "list-aggregate-discovered-resources", "", false, "List Aggregate Discovered Resources")
	_configserviceCmd.Flags().BoolVarP(&_configserviceListConfigurationRecorders, "list-configuration-recorders", "", false, "List Configuration Recorders")
	_configserviceCmd.Flags().BoolVarP(&_configserviceListConformancePackComplianceScores, "list-conformance-pack-compliance-scores", "", false, "List Conformance Pack Compliance Scores")
	_configserviceCmd.Flags().BoolVarP(&_configserviceListDiscoveredResources, "list-discovered-resources", "", false, "List Discovered Resources")
	_configserviceCmd.Flags().BoolVarP(&_configserviceListResourceEvaluations, "list-resource-evaluations", "", false, "List Resource Evaluations")
	_configserviceCmd.Flags().BoolVarP(&_configserviceListStoredQueries, "list-stored-queries", "", false, "List Stored Queries")
	_configserviceCmd.Flags().BoolVarP(&_configserviceListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_configserviceCmd.Flags().BoolVarP(&_configservicePutAggregationAuthorization, "put-aggregation-authorization", "", false, "Put Aggregation Authorization")
	_configserviceCmd.Flags().BoolVarP(&_configservicePutConfigRule, "put-config-rule", "", false, "Put Config Rule")
	_configserviceCmd.Flags().BoolVarP(&_configservicePutConfigurationAggregator, "put-configuration-aggregator", "", false, "Put Configuration Aggregator")
	_configserviceCmd.Flags().BoolVarP(&_configservicePutConfigurationRecorder, "put-configuration-recorder", "", false, "Put Configuration Recorder")
	_configserviceCmd.Flags().BoolVarP(&_configservicePutConformancePack, "put-conformance-pack", "", false, "Put Conformance Pack")
	_configserviceCmd.Flags().BoolVarP(&_configservicePutDeliveryChannel, "put-delivery-channel", "", false, "Put Delivery Channel")
	_configserviceCmd.Flags().BoolVarP(&_configservicePutEvaluations, "put-evaluations", "", false, "Put Evaluations")
	_configserviceCmd.Flags().BoolVarP(&_configservicePutExternalEvaluation, "put-external-evaluation", "", false, "Put External Evaluation")
	_configserviceCmd.Flags().BoolVarP(&_configservicePutOrganizationConfigRule, "put-organization-config-rule", "", false, "Put Organization Config Rule")
	_configserviceCmd.Flags().BoolVarP(&_configservicePutOrganizationConformancePack, "put-organization-conformance-pack", "", false, "Put Organization Conformance Pack")
	_configserviceCmd.Flags().BoolVarP(&_configservicePutRemediationConfigurations, "put-remediation-configurations", "", false, "Put Remediation Configurations")
	_configserviceCmd.Flags().BoolVarP(&_configservicePutRemediationExceptions, "put-remediation-exceptions", "", false, "Put Remediation Exceptions")
	_configserviceCmd.Flags().BoolVarP(&_configservicePutResourceConfig, "put-resource-config", "", false, "Put Resource Config")
	_configserviceCmd.Flags().BoolVarP(&_configservicePutRetentionConfiguration, "put-retention-configuration", "", false, "Put Retention Configuration")
	_configserviceCmd.Flags().BoolVarP(&_configservicePutServiceLinkedConfigurationRecorder, "put-service-linked-configuration-recorder", "", false, "Put Service Linked Configuration Recorder")
	_configserviceCmd.Flags().BoolVarP(&_configservicePutStoredQuery, "put-stored-query", "", false, "Put Stored Query")
	_configserviceCmd.Flags().BoolVarP(&_configserviceSelectAggregateResourceConfig, "select-aggregate-resource-config", "", false, "Select Aggregate Resource Config")
	_configserviceCmd.Flags().BoolVarP(&_configserviceSelectResourceConfig, "select-resource-config", "", false, "Select Resource Config")
	_configserviceCmd.Flags().BoolVarP(&_configserviceStartConfigRulesEvaluation, "start-config-rules-evaluation", "", false, "Start Config Rules Evaluation")
	_configserviceCmd.Flags().BoolVarP(&_configserviceStartConfigurationRecorder, "start-configuration-recorder", "", false, "Start Configuration Recorder")
	_configserviceCmd.Flags().BoolVarP(&_configserviceStartRemediationExecution, "start-remediation-execution", "", false, "Start Remediation Execution")
	_configserviceCmd.Flags().BoolVarP(&_configserviceStartResourceEvaluation, "start-resource-evaluation", "", false, "Start Resource Evaluation")
	_configserviceCmd.Flags().BoolVarP(&_configserviceStopConfigurationRecorder, "stop-configuration-recorder", "", false, "Stop Configuration Recorder")
	_configserviceCmd.Flags().BoolVarP(&_configserviceTagResource, "tag-resource", "", false, "Tag Resource")
	_configserviceCmd.Flags().BoolVarP(&_configserviceUntagResource, "untag-resource", "", false, "Untag Resource")

}
