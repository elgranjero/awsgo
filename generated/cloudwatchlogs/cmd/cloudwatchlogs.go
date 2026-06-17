package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// cloudwatchlogsCmd represents the cloudwatchlogs command
var _cloudwatchlogsCmd = &cobra.Command{
	Use:   "cloudwatchlogs",
	Short: "AWS cloudwatchlogs CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := cloudwatchlogs.NewFromConfig(cfg)
		if _cloudwatchlogsAssociateKmsKey {
			cloudwatchlogs_AssociateKmsKey(cfg, client)
			return
		}
		if _cloudwatchlogsAssociateSourceToS3TableIntegration {
			cloudwatchlogs_AssociateSourceToS3TableIntegration(cfg, client)
			return
		}
		if _cloudwatchlogsCancelExportTask {
			cloudwatchlogs_CancelExportTask(cfg, client)
			return
		}
		if _cloudwatchlogsCancelImportTask {
			cloudwatchlogs_CancelImportTask(cfg, client)
			return
		}
		if _cloudwatchlogsCreateDelivery {
			cloudwatchlogs_CreateDelivery(cfg, client)
			return
		}
		if _cloudwatchlogsCreateExportTask {
			cloudwatchlogs_CreateExportTask(cfg, client)
			return
		}
		if _cloudwatchlogsCreateImportTask {
			cloudwatchlogs_CreateImportTask(cfg, client)
			return
		}
		if _cloudwatchlogsCreateLogAnomalyDetector {
			cloudwatchlogs_CreateLogAnomalyDetector(cfg, client)
			return
		}
		if _cloudwatchlogsCreateLogGroup {
			cloudwatchlogs_CreateLogGroup(cfg, client)
			return
		}
		if _cloudwatchlogsCreateLogStream {
			cloudwatchlogs_CreateLogStream(cfg, client)
			return
		}
		if _cloudwatchlogsCreateScheduledQuery {
			cloudwatchlogs_CreateScheduledQuery(cfg, client)
			return
		}
		if _cloudwatchlogsDeleteAccountPolicy {
			cloudwatchlogs_DeleteAccountPolicy(cfg, client)
			return
		}
		if _cloudwatchlogsDeleteDataProtectionPolicy {
			cloudwatchlogs_DeleteDataProtectionPolicy(cfg, client)
			return
		}
		if _cloudwatchlogsDeleteDelivery {
			cloudwatchlogs_DeleteDelivery(cfg, client)
			return
		}
		if _cloudwatchlogsDeleteDeliveryDestination {
			cloudwatchlogs_DeleteDeliveryDestination(cfg, client)
			return
		}
		if _cloudwatchlogsDeleteDeliveryDestinationPolicy {
			cloudwatchlogs_DeleteDeliveryDestinationPolicy(cfg, client)
			return
		}
		if _cloudwatchlogsDeleteDeliverySource {
			cloudwatchlogs_DeleteDeliverySource(cfg, client)
			return
		}
		if _cloudwatchlogsDeleteDestination {
			cloudwatchlogs_DeleteDestination(cfg, client)
			return
		}
		if _cloudwatchlogsDeleteIndexPolicy {
			cloudwatchlogs_DeleteIndexPolicy(cfg, client)
			return
		}
		if _cloudwatchlogsDeleteIntegration {
			cloudwatchlogs_DeleteIntegration(cfg, client)
			return
		}
		if _cloudwatchlogsDeleteLogAnomalyDetector {
			cloudwatchlogs_DeleteLogAnomalyDetector(cfg, client)
			return
		}
		if _cloudwatchlogsDeleteLogGroup {
			cloudwatchlogs_DeleteLogGroup(cfg, client)
			return
		}
		if _cloudwatchlogsDeleteLogStream {
			cloudwatchlogs_DeleteLogStream(cfg, client)
			return
		}
		if _cloudwatchlogsDeleteMetricFilter {
			cloudwatchlogs_DeleteMetricFilter(cfg, client)
			return
		}
		if _cloudwatchlogsDeleteQueryDefinition {
			cloudwatchlogs_DeleteQueryDefinition(cfg, client)
			return
		}
		if _cloudwatchlogsDeleteResourcePolicy {
			cloudwatchlogs_DeleteResourcePolicy(cfg, client)
			return
		}
		if _cloudwatchlogsDeleteRetentionPolicy {
			cloudwatchlogs_DeleteRetentionPolicy(cfg, client)
			return
		}
		if _cloudwatchlogsDeleteScheduledQuery {
			cloudwatchlogs_DeleteScheduledQuery(cfg, client)
			return
		}
		if _cloudwatchlogsDeleteSubscriptionFilter {
			cloudwatchlogs_DeleteSubscriptionFilter(cfg, client)
			return
		}
		if _cloudwatchlogsDeleteTransformer {
			cloudwatchlogs_DeleteTransformer(cfg, client)
			return
		}
		if _cloudwatchlogsDescribeAccountPolicies {
			cloudwatchlogs_DescribeAccountPolicies(cfg, client)
			return
		}
		if _cloudwatchlogsDescribeConfigurationTemplates {
			cloudwatchlogs_DescribeConfigurationTemplates(cfg, client)
			return
		}
		if _cloudwatchlogsDescribeDeliveries {
			cloudwatchlogs_DescribeDeliveries(cfg, client)
			return
		}
		if _cloudwatchlogsDescribeDeliveryDestinations {
			cloudwatchlogs_DescribeDeliveryDestinations(cfg, client)
			return
		}
		if _cloudwatchlogsDescribeDeliverySources {
			cloudwatchlogs_DescribeDeliverySources(cfg, client)
			return
		}
		if _cloudwatchlogsDescribeDestinations {
			cloudwatchlogs_DescribeDestinations(cfg, client)
			return
		}
		if _cloudwatchlogsDescribeExportTasks {
			cloudwatchlogs_DescribeExportTasks(cfg, client)
			return
		}
		if _cloudwatchlogsDescribeFieldIndexes {
			cloudwatchlogs_DescribeFieldIndexes(cfg, client)
			return
		}
		if _cloudwatchlogsDescribeImportTaskBatches {
			cloudwatchlogs_DescribeImportTaskBatches(cfg, client)
			return
		}
		if _cloudwatchlogsDescribeImportTasks {
			cloudwatchlogs_DescribeImportTasks(cfg, client)
			return
		}
		if _cloudwatchlogsDescribeIndexPolicies {
			cloudwatchlogs_DescribeIndexPolicies(cfg, client)
			return
		}
		if _cloudwatchlogsDescribeLogGroups {
			cloudwatchlogs_DescribeLogGroups(cfg, client)
			return
		}
		if _cloudwatchlogsDescribeLogStreams {
			cloudwatchlogs_DescribeLogStreams(cfg, client)
			return
		}
		if _cloudwatchlogsDescribeMetricFilters {
			cloudwatchlogs_DescribeMetricFilters(cfg, client)
			return
		}
		if _cloudwatchlogsDescribeQueries {
			cloudwatchlogs_DescribeQueries(cfg, client)
			return
		}
		if _cloudwatchlogsDescribeQueryDefinitions {
			cloudwatchlogs_DescribeQueryDefinitions(cfg, client)
			return
		}
		if _cloudwatchlogsDescribeResourcePolicies {
			cloudwatchlogs_DescribeResourcePolicies(cfg, client)
			return
		}
		if _cloudwatchlogsDescribeSubscriptionFilters {
			cloudwatchlogs_DescribeSubscriptionFilters(cfg, client)
			return
		}
		if _cloudwatchlogsDisassociateKmsKey {
			cloudwatchlogs_DisassociateKmsKey(cfg, client)
			return
		}
		if _cloudwatchlogsDisassociateSourceFromS3TableIntegration {
			cloudwatchlogs_DisassociateSourceFromS3TableIntegration(cfg, client)
			return
		}
		if _cloudwatchlogsFilterLogEvents {
			cloudwatchlogs_FilterLogEvents(cfg, client)
			return
		}
		if _cloudwatchlogsGetDataProtectionPolicy {
			cloudwatchlogs_GetDataProtectionPolicy(cfg, client)
			return
		}
		if _cloudwatchlogsGetDelivery {
			cloudwatchlogs_GetDelivery(cfg, client)
			return
		}
		if _cloudwatchlogsGetDeliveryDestination {
			cloudwatchlogs_GetDeliveryDestination(cfg, client)
			return
		}
		if _cloudwatchlogsGetDeliveryDestinationPolicy {
			cloudwatchlogs_GetDeliveryDestinationPolicy(cfg, client)
			return
		}
		if _cloudwatchlogsGetDeliverySource {
			cloudwatchlogs_GetDeliverySource(cfg, client)
			return
		}
		if _cloudwatchlogsGetIntegration {
			cloudwatchlogs_GetIntegration(cfg, client)
			return
		}
		if _cloudwatchlogsGetLogAnomalyDetector {
			cloudwatchlogs_GetLogAnomalyDetector(cfg, client)
			return
		}
		if _cloudwatchlogsGetLogEvents {
			cloudwatchlogs_GetLogEvents(cfg, client)
			return
		}
		if _cloudwatchlogsGetLogFields {
			cloudwatchlogs_GetLogFields(cfg, client)
			return
		}
		if _cloudwatchlogsGetLogGroupFields {
			cloudwatchlogs_GetLogGroupFields(cfg, client)
			return
		}
		if _cloudwatchlogsGetLogObject {
			cloudwatchlogs_GetLogObject(cfg, client)
			return
		}
		if _cloudwatchlogsGetLogRecord {
			cloudwatchlogs_GetLogRecord(cfg, client)
			return
		}
		if _cloudwatchlogsGetQueryResults {
			cloudwatchlogs_GetQueryResults(cfg, client)
			return
		}
		if _cloudwatchlogsGetScheduledQuery {
			cloudwatchlogs_GetScheduledQuery(cfg, client)
			return
		}
		if _cloudwatchlogsGetScheduledQueryHistory {
			cloudwatchlogs_GetScheduledQueryHistory(cfg, client)
			return
		}
		if _cloudwatchlogsGetTransformer {
			cloudwatchlogs_GetTransformer(cfg, client)
			return
		}
		if _cloudwatchlogsListAggregateLogGroupSummaries {
			cloudwatchlogs_ListAggregateLogGroupSummaries(cfg, client)
			return
		}
		if _cloudwatchlogsListAnomalies {
			cloudwatchlogs_ListAnomalies(cfg, client)
			return
		}
		if _cloudwatchlogsListIntegrations {
			cloudwatchlogs_ListIntegrations(cfg, client)
			return
		}
		if _cloudwatchlogsListLogAnomalyDetectors {
			cloudwatchlogs_ListLogAnomalyDetectors(cfg, client)
			return
		}
		if _cloudwatchlogsListLogGroups {
			cloudwatchlogs_ListLogGroups(cfg, client)
			return
		}
		if _cloudwatchlogsListLogGroupsForQuery {
			cloudwatchlogs_ListLogGroupsForQuery(cfg, client)
			return
		}
		if _cloudwatchlogsListScheduledQueries {
			cloudwatchlogs_ListScheduledQueries(cfg, client)
			return
		}
		if _cloudwatchlogsListSourcesForS3TableIntegration {
			cloudwatchlogs_ListSourcesForS3TableIntegration(cfg, client)
			return
		}
		if _cloudwatchlogsListTagsForResource {
			cloudwatchlogs_ListTagsForResource(cfg, client)
			return
		}
		if _cloudwatchlogsListTagsLogGroup {
			cloudwatchlogs_ListTagsLogGroup(cfg, client)
			return
		}
		if _cloudwatchlogsPutAccountPolicy {
			cloudwatchlogs_PutAccountPolicy(cfg, client)
			return
		}
		if _cloudwatchlogsPutBearerTokenAuthentication {
			cloudwatchlogs_PutBearerTokenAuthentication(cfg, client)
			return
		}
		if _cloudwatchlogsPutDataProtectionPolicy {
			cloudwatchlogs_PutDataProtectionPolicy(cfg, client)
			return
		}
		if _cloudwatchlogsPutDeliveryDestination {
			cloudwatchlogs_PutDeliveryDestination(cfg, client)
			return
		}
		if _cloudwatchlogsPutDeliveryDestinationPolicy {
			cloudwatchlogs_PutDeliveryDestinationPolicy(cfg, client)
			return
		}
		if _cloudwatchlogsPutDeliverySource {
			cloudwatchlogs_PutDeliverySource(cfg, client)
			return
		}
		if _cloudwatchlogsPutDestination {
			cloudwatchlogs_PutDestination(cfg, client)
			return
		}
		if _cloudwatchlogsPutDestinationPolicy {
			cloudwatchlogs_PutDestinationPolicy(cfg, client)
			return
		}
		if _cloudwatchlogsPutIndexPolicy {
			cloudwatchlogs_PutIndexPolicy(cfg, client)
			return
		}
		if _cloudwatchlogsPutIntegration {
			cloudwatchlogs_PutIntegration(cfg, client)
			return
		}
		if _cloudwatchlogsPutLogEvents {
			cloudwatchlogs_PutLogEvents(cfg, client)
			return
		}
		if _cloudwatchlogsPutLogGroupDeletionProtection {
			cloudwatchlogs_PutLogGroupDeletionProtection(cfg, client)
			return
		}
		if _cloudwatchlogsPutMetricFilter {
			cloudwatchlogs_PutMetricFilter(cfg, client)
			return
		}
		if _cloudwatchlogsPutQueryDefinition {
			cloudwatchlogs_PutQueryDefinition(cfg, client)
			return
		}
		if _cloudwatchlogsPutResourcePolicy {
			cloudwatchlogs_PutResourcePolicy(cfg, client)
			return
		}
		if _cloudwatchlogsPutRetentionPolicy {
			cloudwatchlogs_PutRetentionPolicy(cfg, client)
			return
		}
		if _cloudwatchlogsPutSubscriptionFilter {
			cloudwatchlogs_PutSubscriptionFilter(cfg, client)
			return
		}
		if _cloudwatchlogsPutTransformer {
			cloudwatchlogs_PutTransformer(cfg, client)
			return
		}
		if _cloudwatchlogsStartLiveTail {
			cloudwatchlogs_StartLiveTail(cfg, client)
			return
		}
		if _cloudwatchlogsStartQuery {
			cloudwatchlogs_StartQuery(cfg, client)
			return
		}
		if _cloudwatchlogsStopQuery {
			cloudwatchlogs_StopQuery(cfg, client)
			return
		}
		if _cloudwatchlogsTagLogGroup {
			cloudwatchlogs_TagLogGroup(cfg, client)
			return
		}
		if _cloudwatchlogsTagResource {
			cloudwatchlogs_TagResource(cfg, client)
			return
		}
		if _cloudwatchlogsTestMetricFilter {
			cloudwatchlogs_TestMetricFilter(cfg, client)
			return
		}
		if _cloudwatchlogsTestTransformer {
			cloudwatchlogs_TestTransformer(cfg, client)
			return
		}
		if _cloudwatchlogsUntagLogGroup {
			cloudwatchlogs_UntagLogGroup(cfg, client)
			return
		}
		if _cloudwatchlogsUntagResource {
			cloudwatchlogs_UntagResource(cfg, client)
			return
		}
		if _cloudwatchlogsUpdateAnomaly {
			cloudwatchlogs_UpdateAnomaly(cfg, client)
			return
		}
		if _cloudwatchlogsUpdateDeliveryConfiguration {
			cloudwatchlogs_UpdateDeliveryConfiguration(cfg, client)
			return
		}
		if _cloudwatchlogsUpdateLogAnomalyDetector {
			cloudwatchlogs_UpdateLogAnomalyDetector(cfg, client)
			return
		}
		if _cloudwatchlogsUpdateScheduledQuery {
			cloudwatchlogs_UpdateScheduledQuery(cfg, client)
			return
		}

	},
}

var (
	_cloudwatchlogsAssociateKmsKey                          bool
	_cloudwatchlogsAssociateSourceToS3TableIntegration      bool
	_cloudwatchlogsCancelExportTask                         bool
	_cloudwatchlogsCancelImportTask                         bool
	_cloudwatchlogsCreateDelivery                           bool
	_cloudwatchlogsCreateExportTask                         bool
	_cloudwatchlogsCreateImportTask                         bool
	_cloudwatchlogsCreateLogAnomalyDetector                 bool
	_cloudwatchlogsCreateLogGroup                           bool
	_cloudwatchlogsCreateLogStream                          bool
	_cloudwatchlogsCreateScheduledQuery                     bool
	_cloudwatchlogsDeleteAccountPolicy                      bool
	_cloudwatchlogsDeleteDataProtectionPolicy               bool
	_cloudwatchlogsDeleteDelivery                           bool
	_cloudwatchlogsDeleteDeliveryDestination                bool
	_cloudwatchlogsDeleteDeliveryDestinationPolicy          bool
	_cloudwatchlogsDeleteDeliverySource                     bool
	_cloudwatchlogsDeleteDestination                        bool
	_cloudwatchlogsDeleteIndexPolicy                        bool
	_cloudwatchlogsDeleteIntegration                        bool
	_cloudwatchlogsDeleteLogAnomalyDetector                 bool
	_cloudwatchlogsDeleteLogGroup                           bool
	_cloudwatchlogsDeleteLogStream                          bool
	_cloudwatchlogsDeleteMetricFilter                       bool
	_cloudwatchlogsDeleteQueryDefinition                    bool
	_cloudwatchlogsDeleteResourcePolicy                     bool
	_cloudwatchlogsDeleteRetentionPolicy                    bool
	_cloudwatchlogsDeleteScheduledQuery                     bool
	_cloudwatchlogsDeleteSubscriptionFilter                 bool
	_cloudwatchlogsDeleteTransformer                        bool
	_cloudwatchlogsDescribeAccountPolicies                  bool
	_cloudwatchlogsDescribeConfigurationTemplates           bool
	_cloudwatchlogsDescribeDeliveries                       bool
	_cloudwatchlogsDescribeDeliveryDestinations             bool
	_cloudwatchlogsDescribeDeliverySources                  bool
	_cloudwatchlogsDescribeDestinations                     bool
	_cloudwatchlogsDescribeExportTasks                      bool
	_cloudwatchlogsDescribeFieldIndexes                     bool
	_cloudwatchlogsDescribeImportTaskBatches                bool
	_cloudwatchlogsDescribeImportTasks                      bool
	_cloudwatchlogsDescribeIndexPolicies                    bool
	_cloudwatchlogsDescribeLogGroups                        bool
	_cloudwatchlogsDescribeLogStreams                       bool
	_cloudwatchlogsDescribeMetricFilters                    bool
	_cloudwatchlogsDescribeQueries                          bool
	_cloudwatchlogsDescribeQueryDefinitions                 bool
	_cloudwatchlogsDescribeResourcePolicies                 bool
	_cloudwatchlogsDescribeSubscriptionFilters              bool
	_cloudwatchlogsDisassociateKmsKey                       bool
	_cloudwatchlogsDisassociateSourceFromS3TableIntegration bool
	_cloudwatchlogsFilterLogEvents                          bool
	_cloudwatchlogsGetDataProtectionPolicy                  bool
	_cloudwatchlogsGetDelivery                              bool
	_cloudwatchlogsGetDeliveryDestination                   bool
	_cloudwatchlogsGetDeliveryDestinationPolicy             bool
	_cloudwatchlogsGetDeliverySource                        bool
	_cloudwatchlogsGetIntegration                           bool
	_cloudwatchlogsGetLogAnomalyDetector                    bool
	_cloudwatchlogsGetLogEvents                             bool
	_cloudwatchlogsGetLogFields                             bool
	_cloudwatchlogsGetLogGroupFields                        bool
	_cloudwatchlogsGetLogObject                             bool
	_cloudwatchlogsGetLogRecord                             bool
	_cloudwatchlogsGetQueryResults                          bool
	_cloudwatchlogsGetScheduledQuery                        bool
	_cloudwatchlogsGetScheduledQueryHistory                 bool
	_cloudwatchlogsGetTransformer                           bool
	_cloudwatchlogsListAggregateLogGroupSummaries           bool
	_cloudwatchlogsListAnomalies                            bool
	_cloudwatchlogsListIntegrations                         bool
	_cloudwatchlogsListLogAnomalyDetectors                  bool
	_cloudwatchlogsListLogGroups                            bool
	_cloudwatchlogsListLogGroupsForQuery                    bool
	_cloudwatchlogsListScheduledQueries                     bool
	_cloudwatchlogsListSourcesForS3TableIntegration         bool
	_cloudwatchlogsListTagsForResource                      bool
	_cloudwatchlogsListTagsLogGroup                         bool
	_cloudwatchlogsPutAccountPolicy                         bool
	_cloudwatchlogsPutBearerTokenAuthentication             bool
	_cloudwatchlogsPutDataProtectionPolicy                  bool
	_cloudwatchlogsPutDeliveryDestination                   bool
	_cloudwatchlogsPutDeliveryDestinationPolicy             bool
	_cloudwatchlogsPutDeliverySource                        bool
	_cloudwatchlogsPutDestination                           bool
	_cloudwatchlogsPutDestinationPolicy                     bool
	_cloudwatchlogsPutIndexPolicy                           bool
	_cloudwatchlogsPutIntegration                           bool
	_cloudwatchlogsPutLogEvents                             bool
	_cloudwatchlogsPutLogGroupDeletionProtection            bool
	_cloudwatchlogsPutMetricFilter                          bool
	_cloudwatchlogsPutQueryDefinition                       bool
	_cloudwatchlogsPutResourcePolicy                        bool
	_cloudwatchlogsPutRetentionPolicy                       bool
	_cloudwatchlogsPutSubscriptionFilter                    bool
	_cloudwatchlogsPutTransformer                           bool
	_cloudwatchlogsStartLiveTail                            bool
	_cloudwatchlogsStartQuery                               bool
	_cloudwatchlogsStopQuery                                bool
	_cloudwatchlogsTagLogGroup                              bool
	_cloudwatchlogsTagResource                              bool
	_cloudwatchlogsTestMetricFilter                         bool
	_cloudwatchlogsTestTransformer                          bool
	_cloudwatchlogsUntagLogGroup                            bool
	_cloudwatchlogsUntagResource                            bool
	_cloudwatchlogsUpdateAnomaly                            bool
	_cloudwatchlogsUpdateDeliveryConfiguration              bool
	_cloudwatchlogsUpdateLogAnomalyDetector                 bool
	_cloudwatchlogsUpdateScheduledQuery                     bool

	_cloudwatchlogsAccessPolicy                     string
	_cloudwatchlogsAccountIdentifiers               []string
	_cloudwatchlogsAnomalyDetectorArn               string
	_cloudwatchlogsAnomalyId                        string
	_cloudwatchlogsAnomalyVisibilityTime            string
	_cloudwatchlogsApplyOnTransformedLogs           string
	_cloudwatchlogsBaseline                         string
	_cloudwatchlogsBatchImportStatus                string
	_cloudwatchlogsBearerTokenAuthenticationEnabled string
	_cloudwatchlogsClientToken                      string
	_cloudwatchlogsDataSource                       string
	_cloudwatchlogsDataSourceName                   string
	_cloudwatchlogsDataSourceType                   string
	_cloudwatchlogsDataSources                      string
	_cloudwatchlogsDeletionProtectionEnabled        string
	_cloudwatchlogsDeliveryDestinationArn           string
	_cloudwatchlogsDeliveryDestinationConfiguration string
	_cloudwatchlogsDeliveryDestinationName          string
	_cloudwatchlogsDeliveryDestinationPolicy        string
	_cloudwatchlogsDeliveryDestinationType          string
	_cloudwatchlogsDeliveryDestinationTypes         string
	_cloudwatchlogsDeliverySourceName               string
	_cloudwatchlogsDescending                       string
	_cloudwatchlogsDescription                      string
	_cloudwatchlogsDestination                      string
	_cloudwatchlogsDestinationArn                   string
	_cloudwatchlogsDestinationConfiguration         string
	_cloudwatchlogsDestinationName                  string
	_cloudwatchlogsDestinationNamePrefix            string
	_cloudwatchlogsDestinationPrefix                string
	_cloudwatchlogsDetectorName                     string
	_cloudwatchlogsDistribution                     string
	_cloudwatchlogsEmitSystemFieldDimensions        []string
	_cloudwatchlogsEmitSystemFields                 []string
	_cloudwatchlogsEnabled                          string
	_cloudwatchlogsEndTime                          string
	_cloudwatchlogsEntity                           string
	_cloudwatchlogsEvaluationFrequency              string
	_cloudwatchlogsExecutionRoleArn                 string
	_cloudwatchlogsExecutionStatuses                string
	_cloudwatchlogsExpectedRevisionId               string
	_cloudwatchlogsFieldDelimiter                   string
	_cloudwatchlogsFieldIndexNames                  []string
	_cloudwatchlogsFieldSelectionCriteria           string
	_cloudwatchlogsFilterLogGroupArn                string
	_cloudwatchlogsFilterName                       string
	_cloudwatchlogsFilterNamePrefix                 string
	_cloudwatchlogsFilterPattern                    string
	_cloudwatchlogsForce                            string
	_cloudwatchlogsForceUpdate                      string
	_cloudwatchlogsFrom                             string
	_cloudwatchlogsGroupBy                          string
	_cloudwatchlogsId                               string
	_cloudwatchlogsIdentifier                       string
	_cloudwatchlogsImportFilter                     string
	_cloudwatchlogsImportId                         string
	_cloudwatchlogsImportRoleArn                    string
	_cloudwatchlogsImportSourceArn                  string
	_cloudwatchlogsImportStatus                     string
	_cloudwatchlogsIncludeLinkedAccounts            string
	_cloudwatchlogsIntegrationArn                   string
	_cloudwatchlogsIntegrationName                  string
	_cloudwatchlogsIntegrationNamePrefix            string
	_cloudwatchlogsIntegrationStatus                string
	_cloudwatchlogsIntegrationType                  string
	_cloudwatchlogsInterleaved                      string
	_cloudwatchlogsKmsKeyId                         string
	_cloudwatchlogsLimit                            string
	_cloudwatchlogsLogEventFilterPattern            string
	_cloudwatchlogsLogEventMessages                 []string
	_cloudwatchlogsLogEvents                        string
	_cloudwatchlogsLogGroupArnList                  []string
	_cloudwatchlogsLogGroupClass                    string
	_cloudwatchlogsLogGroupIdentifier               string
	_cloudwatchlogsLogGroupIdentifiers              []string
	_cloudwatchlogsLogGroupName                     string
	_cloudwatchlogsLogGroupNamePattern              string
	_cloudwatchlogsLogGroupNamePrefix               string
	_cloudwatchlogsLogGroupNames                    []string
	_cloudwatchlogsLogObjectPointer                 string
	_cloudwatchlogsLogRecordPointer                 string
	_cloudwatchlogsLogStreamName                    string
	_cloudwatchlogsLogStreamNamePrefix              string
	_cloudwatchlogsLogStreamNamePrefixes            []string
	_cloudwatchlogsLogStreamNames                   []string
	_cloudwatchlogsLogType                          string
	_cloudwatchlogsLogTypes                         []string
	_cloudwatchlogsMaxResults                       string
	_cloudwatchlogsMetricName                       string
	_cloudwatchlogsMetricNamespace                  string
	_cloudwatchlogsMetricTransformations            string
	_cloudwatchlogsName                             string
	_cloudwatchlogsNextToken                        string
	_cloudwatchlogsOrderBy                          string
	_cloudwatchlogsOutputFormat                     string
	_cloudwatchlogsPatternId                        string
	_cloudwatchlogsPolicyDocument                   string
	_cloudwatchlogsPolicyName                       string
	_cloudwatchlogsPolicyScope                      string
	_cloudwatchlogsPolicyType                       string
	_cloudwatchlogsQueryDefinitionId                string
	_cloudwatchlogsQueryDefinitionNamePrefix        string
	_cloudwatchlogsQueryId                          string
	_cloudwatchlogsQueryLanguage                    string
	_cloudwatchlogsQueryString                      string
	_cloudwatchlogsRecordFields                     []string
	_cloudwatchlogsResourceArn                      string
	_cloudwatchlogsResourceConfig                   string
	_cloudwatchlogsResourceIdentifier               string
	_cloudwatchlogsResourceTypes                    []string
	_cloudwatchlogsRetentionInDays                  string
	_cloudwatchlogsRoleArn                          string
	_cloudwatchlogsS3DeliveryConfiguration          string
	_cloudwatchlogsScheduleEndTime                  string
	_cloudwatchlogsScheduleExpression               string
	_cloudwatchlogsScheduleStartTime                string
	_cloudwatchlogsScope                            string
	_cloudwatchlogsSelectionCriteria                string
	_cloudwatchlogsSequenceToken                    string
	_cloudwatchlogsService                          string
	_cloudwatchlogsStartFromHead                    string
	_cloudwatchlogsStartTime                        string
	_cloudwatchlogsStartTimeOffset                  string
	_cloudwatchlogsState                            string
	_cloudwatchlogsStatus                           string
	_cloudwatchlogsStatusCode                       string
	_cloudwatchlogsSuppressionPeriod                string
	_cloudwatchlogsSuppressionState                 string
	_cloudwatchlogsSuppressionType                  string
	_cloudwatchlogsTagKeys                          []string
	_cloudwatchlogsTags                             []string
	_cloudwatchlogsTargetArn                        string
	_cloudwatchlogsTaskId                           string
	_cloudwatchlogsTaskName                         string
	_cloudwatchlogsTime                             string
	_cloudwatchlogsTimezone                         string
	_cloudwatchlogsTo                               string
	_cloudwatchlogsTransformerConfig                string
	_cloudwatchlogsUnmask                           string
)

// Associates the specified KMS key with either one log group in the account, or
// with all stored CloudWatch Logs query insights results in the account.
//
// When you use AssociateKmsKey , you specify either the logGroupName parameter or
// the resourceIdentifier parameter. You can't specify both of those parameters in
// the same operation.
//
// - Specify the logGroupName parameter to cause log events ingested into that
// log group to be encrypted with that key. Only the log events ingested after the
// key is associated are encrypted with that key.
//
// # Associating a KMS key with a log group overrides any existing associations
//
// between the log group and a KMS key. After a KMS key is associated with a log
// group, all newly ingested data for the log group is encrypted using the KMS key.
// This association is stored as long as the data encrypted with the KMS key is
// still within CloudWatch Logs. This enables CloudWatch Logs to decrypt this data
// whenever it is requested.
//
// # Associating a key with a log group does not cause the results of queries of
//
// that log group to be encrypted with that key. To have query results encrypted
// with a KMS key, you must use an AssociateKmsKey operation with the
// resourceIdentifier parameter that specifies a query-result resource.
//
// - Specify the resourceIdentifier parameter with a query-result resource, to
// use that key to encrypt the stored results of all future [StartQuery]operations in the
// account. The response from a [GetQueryResults]operation will still return the query results in
// plain text.
//
// # Even if you have not associated a key with your query results, the query
//
// results are encrypted when stored, using the default CloudWatch Logs method.
//
// # If you run a query from a monitoring account that queries logs in a source
//
// account, the query results key from the monitoring account, if any, is used.
//
// If you delete the key that is used to encrypt log events or log group query
// results, then all the associated stored log events or query results that were
// encrypted with that key will be unencryptable and unusable.
//
// CloudWatch Logs supports only symmetric KMS keys. Do not associate an
// asymmetric KMS key with your log group or query results. For more information,
// see [Using Symmetric and Asymmetric Keys].
//
// It can take up to 5 minutes for this operation to take effect.
//
// If you attempt to associate a KMS key with a log group but the KMS key does not
// exist or the KMS key is disabled, you receive an InvalidParameterException
// error.
//
// [Using Symmetric and Asymmetric Keys]: https://docs.aws.amazon.com/kms/latest/developerguide/symmetric-asymmetric.html
// [StartQuery]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_StartQuery.html
// [GetQueryResults]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_GetQueryResults.html
func cloudwatchlogs_AssociateKmsKey(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.AssociateKmsKeyInput{
		// KmsKeyId: *string, // Required
	}

	if len(_cloudwatchlogsKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_cloudwatchlogsKmsKeyId)
	}
	if len(_cloudwatchlogsLogGroupName) > 0 {
		input.LogGroupName = aws.String(_cloudwatchlogsLogGroupName)
	}
	if len(_cloudwatchlogsResourceIdentifier) > 0 {
		input.ResourceIdentifier = aws.String(_cloudwatchlogsResourceIdentifier)
	}

	if resp, err := client.AssociateKmsKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates a data source with an S3 Table Integration for query access in the
// 'logs' namespace. This enables querying log data using analytics engines that
// support Iceberg such as Amazon Athena, Amazon Redshift, and Apache Spark.
func cloudwatchlogs_AssociateSourceToS3TableIntegration(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.AssociateSourceToS3TableIntegrationInput{
		// DataSource: *types.DataSource, // Required
		// IntegrationArn: *string, // Required
	}

	if len(_cloudwatchlogsDataSource) > 0 {
		if err := assignInputField(input, "DataSource", _cloudwatchlogsDataSource); err != nil {
			log.Errorf("invalid --data-source: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsIntegrationArn) > 0 {
		input.IntegrationArn = aws.String(_cloudwatchlogsIntegrationArn)
	}

	if resp, err := client.AssociateSourceToS3TableIntegration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels the specified export task.
// The task must be in the PENDING or RUNNING state.
func cloudwatchlogs_CancelExportTask(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.CancelExportTaskInput{
		// TaskId: *string, // Required
	}

	if len(_cloudwatchlogsTaskId) > 0 {
		input.TaskId = aws.String(_cloudwatchlogsTaskId)
	}

	if resp, err := client.CancelExportTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels an active import task and stops importing data from the CloudTrail Lake
// Event Data Store.
func cloudwatchlogs_CancelImportTask(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.CancelImportTaskInput{
		// ImportId: *string, // Required
	}

	if len(_cloudwatchlogsImportId) > 0 {
		input.ImportId = aws.String(_cloudwatchlogsImportId)
	}

	if resp, err := client.CancelImportTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a delivery. A delivery is a connection between a logical delivery
// source and a logical delivery destination that you have already created.
//
// Only some Amazon Web Services services support being configured as a delivery
// source using this operation. These services are listed as Supported [V2
// Permissions] in the table at [Enabling logging from Amazon Web Services services.]
//
// A delivery destination can represent a log group in CloudWatch Logs, an Amazon
// S3 bucket, a delivery stream in Firehose, or X-Ray.
//
// To configure logs delivery between a supported Amazon Web Services service and
// a destination, you must do the following:
//
// - Create a delivery source, which is a logical object that represents the
// resource that is actually sending the logs. For more information, see [PutDeliverySource].
//
// - Create a delivery destination, which is a logical object that represents
// the actual delivery destination. For more information, see [PutDeliveryDestination].
//
// - If you are delivering logs cross-account, you must use [PutDeliveryDestinationPolicy]in the destination
// account to assign an IAM policy to the destination. This policy allows delivery
// to that destination.
//
// - Use CreateDelivery to create a delivery by pairing exactly one delivery
// source and one delivery destination.
//
// You can configure a single delivery source to send logs to multiple
// destinations by creating multiple deliveries. You can also create multiple
// deliveries to configure multiple delivery sources to send logs to the same
// delivery destination.
//
// To update an existing delivery configuration, use [UpdateDeliveryConfiguration].
//
// [PutDeliveryDestination]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutDeliveryDestination.html
// [PutDeliverySource]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutDeliverySource.html
// [Enabling logging from Amazon Web Services services.]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-logs-and-resource-policy.html
// [PutDeliveryDestinationPolicy]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutDeliveryDestinationPolicy.html
// [UpdateDeliveryConfiguration]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_UpdateDeliveryConfiguration.html
func cloudwatchlogs_CreateDelivery(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.CreateDeliveryInput{
		// DeliveryDestinationArn: *string, // Required
		// DeliverySourceName: *string, // Required
	}

	if len(_cloudwatchlogsDeliveryDestinationArn) > 0 {
		input.DeliveryDestinationArn = aws.String(_cloudwatchlogsDeliveryDestinationArn)
	}
	if len(_cloudwatchlogsDeliverySourceName) > 0 {
		input.DeliverySourceName = aws.String(_cloudwatchlogsDeliverySourceName)
	}
	if len(_cloudwatchlogsFieldDelimiter) > 0 {
		input.FieldDelimiter = aws.String(_cloudwatchlogsFieldDelimiter)
	}
	if len(_cloudwatchlogsRecordFields) > 0 {
		input.RecordFields = append([]string(nil), _cloudwatchlogsRecordFields...)
	}
	if len(_cloudwatchlogsS3DeliveryConfiguration) > 0 {
		if err := assignInputField(input, "S3DeliveryConfiguration", _cloudwatchlogsS3DeliveryConfiguration); err != nil {
			log.Errorf("invalid --s3-delivery-configuration: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsTags) > 0 {
		if err := assignInputField(input, "Tags", _cloudwatchlogsTags[0]); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDelivery(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an export task so that you can efficiently export data from a log group
// to an Amazon S3 bucket. When you perform a CreateExportTask operation, you must
// use credentials that have permission to write to the S3 bucket that you specify
// as the destination.
//
// Exporting log data to S3 buckets that are encrypted by KMS is supported.
// Exporting log data to Amazon S3 buckets that have S3 Object Lock enabled with a
// retention period is also supported.
//
// Exporting to S3 buckets that are encrypted with AES-256 is supported.
//
// This is an asynchronous call. If all the required information is provided, this
// operation initiates an export task and responds with the ID of the task. After
// the task has started, you can use [DescribeExportTasks]to get the status of the export task. Each
// account can only have one active ( RUNNING or PENDING ) export task at a time.
// To cancel an export task, use [CancelExportTask].
//
// You can export logs from multiple log groups or multiple time ranges to the
// same S3 bucket. To separate log data for each export task, specify a prefix to
// be used as the Amazon S3 key prefix for all exported objects.
//
// We recommend that you don't regularly export to Amazon S3 as a way to
// continuously archive your logs. For that use case, we instead recommend that you
// use subscriptions. For more information about subscriptions, see [Real-time processing of log data with subscriptions].
//
// Time-based sorting on chunks of log data inside an exported file is not
// guaranteed. You can sort the exported log field data by using Linux utilities.
//
// [CancelExportTask]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_CancelExportTask.html
// [DescribeExportTasks]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_DescribeExportTasks.html
// [Real-time processing of log data with subscriptions]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/Subscriptions.html
func cloudwatchlogs_CreateExportTask(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.CreateExportTaskInput{
		// Destination: *string, // Required
		// From: *int64, // Required
		// LogGroupName: *string, // Required
		// To: *int64, // Required
	}

	if len(_cloudwatchlogsDestination) > 0 {
		input.Destination = aws.String(_cloudwatchlogsDestination)
	}
	if len(_cloudwatchlogsFrom) > 0 {
		if err := assignInputField(input, "From", _cloudwatchlogsFrom); err != nil {
			log.Errorf("invalid --from: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsLogGroupName) > 0 {
		input.LogGroupName = aws.String(_cloudwatchlogsLogGroupName)
	}
	if len(_cloudwatchlogsTo) > 0 {
		if err := assignInputField(input, "To", _cloudwatchlogsTo); err != nil {
			log.Errorf("invalid --to: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsDestinationPrefix) > 0 {
		input.DestinationPrefix = aws.String(_cloudwatchlogsDestinationPrefix)
	}
	if len(_cloudwatchlogsLogStreamNamePrefix) > 0 {
		input.LogStreamNamePrefix = aws.String(_cloudwatchlogsLogStreamNamePrefix)
	}
	if len(_cloudwatchlogsTaskName) > 0 {
		input.TaskName = aws.String(_cloudwatchlogsTaskName)
	}

	if resp, err := client.CreateExportTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts an import from a data source to CloudWatch Log and creates a managed log
// group as the destination for the imported data. Currently, [CloudTrail Event Data Store]is the only
// supported data source.
//
// The import task must satisfy the following constraints:
//
// - The specified source must be in an ACTIVE state.
//
// - The API caller must have permissions to access the data in the provided
// source and to perform iam:PassRole on the provided import role which has the
// same permissions, as described below.
//
// - The provided IAM role must trust the "cloudtrail.amazonaws.com" principal
// and have the following permissions:
//
// - cloudtrail:GetEventDataStoreData
//
// - logs:CreateLogGroup
//
// - logs:CreateLogStream
//
// - logs:PutResourcePolicy
//
// - (If source has an associated Amazon Web Services KMS Key) kms:Decrypt
//
// - (If source has an associated Amazon Web Services KMS Key)
// kms:GenerateDataKey
//
// Example IAM policy for provided import role:
//
// [ { "Effect": "Allow", "Action": "iam:PassRole", "Resource":
//
// "arn:aws:iam::123456789012:role/apiCallerCredentials", "Condition": {
// "StringLike": { "iam:AssociatedResourceARN":
// "arn:aws:logs:us-east-1:123456789012:log-group:aws/cloudtrail/f1d45bff-d0e3-4868-b5d9-2eb678aa32fb:*"
// } } }, { "Effect": "Allow", "Action": [ "cloudtrail:GetEventDataStoreData" ],
// "Resource": [
// "arn:aws:cloudtrail:us-east-1:123456789012:eventdatastore/f1d45bff-d0e3-4868-b5d9-2eb678aa32fb"
// ] }, { "Effect": "Allow", "Action": [ "logs:CreateImportTask",
// "logs:CreateLogGroup", "logs:CreateLogStream", "logs:PutResourcePolicy" ],
// "Resource": [ "arn:aws:logs:us-east-1:123456789012:log-group:/aws/cloudtrail/*"
// ] }, { "Effect": "Allow", "Action": [ "kms:Decrypt", "kms:GenerateDataKey" ],
// "Resource": [
// "arn:aws:kms:us-east-1:123456789012:key/12345678-1234-1234-1234-123456789012" ]
// } ]
//
// - If the import source has a customer managed key, the
// "cloudtrail.amazonaws.com" principal needs permissions to perform kms:Decrypt
// and kms:GenerateDataKey.
//
// - There can be no more than 3 active imports per account at a given time.
//
// - The startEventTime must be less than or equal to endEventTime.
//
// - The data being imported must be within the specified source's retention
// period.
//
// [CloudTrail Event Data Store]: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/query-event-data-store.html
func cloudwatchlogs_CreateImportTask(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.CreateImportTaskInput{
		// ImportRoleArn: *string, // Required
		// ImportSourceArn: *string, // Required
	}

	if len(_cloudwatchlogsImportRoleArn) > 0 {
		input.ImportRoleArn = aws.String(_cloudwatchlogsImportRoleArn)
	}
	if len(_cloudwatchlogsImportSourceArn) > 0 {
		input.ImportSourceArn = aws.String(_cloudwatchlogsImportSourceArn)
	}
	if len(_cloudwatchlogsImportFilter) > 0 {
		if err := assignInputField(input, "ImportFilter", _cloudwatchlogsImportFilter); err != nil {
			log.Errorf("invalid --import-filter: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateImportTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an anomaly detector that regularly scans one or more log groups and
// look for patterns and anomalies in the logs.
//
// An anomaly detector can help surface issues by automatically discovering
// anomalies in your log event traffic. An anomaly detector uses machine learning
// algorithms to scan log events and find patterns. A pattern is a shared text
// structure that recurs among your log fields. Patterns provide a useful tool for
// analyzing large sets of logs because a large number of log events can often be
// compressed into a few patterns.
//
// The anomaly detector uses pattern recognition to find anomalies , which are
// unusual log events. It uses the evaluationFrequency to compare current log
// events and patterns with trained baselines.
//
// Fields within a pattern are called tokens. Fields that vary within a pattern,
// such as a request ID or timestamp, are referred to as dynamic tokens and
// represented by <*> .
//
// The following is an example of a pattern:
//
// [INFO] Request time: <*> ms
//
// This pattern represents log events like [INFO] Request time: 327 ms and other
// similar log events that differ only by the number, in this csse 327. When the
// pattern is displayed, the different numbers are replaced by <*>
//
// Any parts of log events that are masked as sensitive data are not scanned for
// anomalies. For more information about masking sensitive data, see [Help protect sensitive log data with masking].
//
// [Help protect sensitive log data with masking]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/mask-sensitive-log-data.html
func cloudwatchlogs_CreateLogAnomalyDetector(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.CreateLogAnomalyDetectorInput{
		// LogGroupArnList: []string, // Required
	}

	if len(_cloudwatchlogsLogGroupArnList) > 0 {
		input.LogGroupArnList = append([]string(nil), _cloudwatchlogsLogGroupArnList...)
	}
	if len(_cloudwatchlogsAnomalyVisibilityTime) > 0 {
		if err := assignInputField(input, "AnomalyVisibilityTime", _cloudwatchlogsAnomalyVisibilityTime); err != nil {
			log.Errorf("invalid --anomaly-visibility-time: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsDetectorName) > 0 {
		input.DetectorName = aws.String(_cloudwatchlogsDetectorName)
	}
	if len(_cloudwatchlogsEvaluationFrequency) > 0 {
		if err := assignInputField(input, "EvaluationFrequency", _cloudwatchlogsEvaluationFrequency); err != nil {
			log.Errorf("invalid --evaluation-frequency: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsFilterPattern) > 0 {
		input.FilterPattern = aws.String(_cloudwatchlogsFilterPattern)
	}
	if len(_cloudwatchlogsKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_cloudwatchlogsKmsKeyId)
	}
	if len(_cloudwatchlogsTags) > 0 {
		if err := assignInputField(input, "Tags", _cloudwatchlogsTags[0]); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateLogAnomalyDetector(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a log group with the specified name. You can create up to 1,000,000 log
// groups per Region per account.
//
// You must use the following guidelines when naming a log group:
//
// - Log group names must be unique within a Region for an Amazon Web Services
// account.
//
// - Log group names can be between 1 and 512 characters long.
//
// - Log group names consist of the following characters: a-z, A-Z, 0-9, '_'
// (underscore), '-' (hyphen), '/' (forward slash), '.' (period), and '#' (number
// sign)
//
// - Log group names can't start with the string aws/
//
// When you create a log group, by default the log events in the log group do not
// expire. To set a retention policy so that events expire and are deleted after a
// specified time, use [PutRetentionPolicy].
//
// If you associate an KMS key with the log group, ingested data is encrypted
// using the KMS key. This association is stored as long as the data encrypted with
// the KMS key is still within CloudWatch Logs. This enables CloudWatch Logs to
// decrypt this data whenever it is requested.
//
// If you attempt to associate a KMS key with the log group but the KMS key does
// not exist or the KMS key is disabled, you receive an InvalidParameterException
// error.
//
// CloudWatch Logs supports only symmetric KMS keys. Do not associate an
// asymmetric KMS key with your log group. For more information, see [Using Symmetric and Asymmetric Keys].
//
// [Using Symmetric and Asymmetric Keys]: https://docs.aws.amazon.com/kms/latest/developerguide/symmetric-asymmetric.html
// [PutRetentionPolicy]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutRetentionPolicy.html
func cloudwatchlogs_CreateLogGroup(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.CreateLogGroupInput{
		// LogGroupName: *string, // Required
	}

	if len(_cloudwatchlogsLogGroupName) > 0 {
		input.LogGroupName = aws.String(_cloudwatchlogsLogGroupName)
	}
	if len(_cloudwatchlogsDeletionProtectionEnabled) > 0 {
		if err := assignInputField(input, "DeletionProtectionEnabled", _cloudwatchlogsDeletionProtectionEnabled); err != nil {
			log.Errorf("invalid --deletion-protection-enabled: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_cloudwatchlogsKmsKeyId)
	}
	if len(_cloudwatchlogsLogGroupClass) > 0 {
		if err := assignInputField(input, "LogGroupClass", _cloudwatchlogsLogGroupClass); err != nil {
			log.Errorf("invalid --log-group-class: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsTags) > 0 {
		if err := assignInputField(input, "Tags", _cloudwatchlogsTags[0]); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateLogGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a log stream for the specified log group. A log stream is a sequence of
// log events that originate from a single source, such as an application instance
// or a resource that is being monitored.
//
// There is no limit on the number of log streams that you can create for a log
// group. There is a limit of 50 TPS on CreateLogStream operations, after which
// transactions are throttled.
//
// You must use the following guidelines when naming a log stream:
//
// - Log stream names must be unique within the log group.
//
// - Log stream names can be between 1 and 512 characters long.
//
// - Don't use ':' (colon) or '*' (asterisk) characters.
func cloudwatchlogs_CreateLogStream(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.CreateLogStreamInput{
		// LogGroupName: *string, // Required
		// LogStreamName: *string, // Required
	}

	if len(_cloudwatchlogsLogGroupName) > 0 {
		input.LogGroupName = aws.String(_cloudwatchlogsLogGroupName)
	}
	if len(_cloudwatchlogsLogStreamName) > 0 {
		input.LogStreamName = aws.String(_cloudwatchlogsLogStreamName)
	}

	if resp, err := client.CreateLogStream(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a scheduled query that runs CloudWatch Logs Insights queries at regular
// intervals. Scheduled queries enable proactive monitoring by automatically
// executing queries to detect patterns and anomalies in your log data. Query
// results can be delivered to Amazon S3 for analysis or further processing.
func cloudwatchlogs_CreateScheduledQuery(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.CreateScheduledQueryInput{
		// ExecutionRoleArn: *string, // Required
		// Name: *string, // Required
		// QueryLanguage: types.QueryLanguage, // Required
		// QueryString: *string, // Required
		// ScheduleExpression: *string, // Required
	}

	if len(_cloudwatchlogsExecutionRoleArn) > 0 {
		input.ExecutionRoleArn = aws.String(_cloudwatchlogsExecutionRoleArn)
	}
	if len(_cloudwatchlogsName) > 0 {
		input.Name = aws.String(_cloudwatchlogsName)
	}
	if len(_cloudwatchlogsQueryLanguage) > 0 {
		if err := assignInputField(input, "QueryLanguage", _cloudwatchlogsQueryLanguage); err != nil {
			log.Errorf("invalid --query-language: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsQueryString) > 0 {
		input.QueryString = aws.String(_cloudwatchlogsQueryString)
	}
	if len(_cloudwatchlogsScheduleExpression) > 0 {
		input.ScheduleExpression = aws.String(_cloudwatchlogsScheduleExpression)
	}
	if len(_cloudwatchlogsDescription) > 0 {
		input.Description = aws.String(_cloudwatchlogsDescription)
	}
	if len(_cloudwatchlogsDestinationConfiguration) > 0 {
		if err := assignInputField(input, "DestinationConfiguration", _cloudwatchlogsDestinationConfiguration); err != nil {
			log.Errorf("invalid --destination-configuration: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsLogGroupIdentifiers) > 0 {
		input.LogGroupIdentifiers = append([]string(nil), _cloudwatchlogsLogGroupIdentifiers...)
	}
	if len(_cloudwatchlogsScheduleEndTime) > 0 {
		if err := assignInputField(input, "ScheduleEndTime", _cloudwatchlogsScheduleEndTime); err != nil {
			log.Errorf("invalid --schedule-end-time: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsScheduleStartTime) > 0 {
		if err := assignInputField(input, "ScheduleStartTime", _cloudwatchlogsScheduleStartTime); err != nil {
			log.Errorf("invalid --schedule-start-time: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsStartTimeOffset) > 0 {
		if err := assignInputField(input, "StartTimeOffset", _cloudwatchlogsStartTimeOffset); err != nil {
			log.Errorf("invalid --start-time-offset: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsState) > 0 {
		if err := assignInputField(input, "State", _cloudwatchlogsState); err != nil {
			log.Errorf("invalid --state: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsTags) > 0 {
		if err := assignInputField(input, "Tags", _cloudwatchlogsTags[0]); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsTimezone) > 0 {
		input.Timezone = aws.String(_cloudwatchlogsTimezone)
	}

	if resp, err := client.CreateScheduledQuery(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a CloudWatch Logs account policy. This stops the account-wide policy
// from applying to log groups or data sources in the account. If you delete a data
// protection policy or subscription filter policy, any log-group level policies of
// those types remain in effect. This operation supports deletion of data
// source-based field index policies, including facet configurations, in addition
// to log group-based policies.
//
// To use this operation, you must be signed on with the correct permissions
// depending on the type of policy that you are deleting.
//
// - To delete a data protection policy, you must have the
// logs:DeleteDataProtectionPolicy and logs:DeleteAccountPolicy permissions.
//
// - To delete a subscription filter policy, you must have the
// logs:DeleteSubscriptionFilter and logs:DeleteAccountPolicy permissions.
//
// - To delete a transformer policy, you must have the logs:DeleteTransformer and
// logs:DeleteAccountPolicy permissions.
//
// - To delete a field index policy, you must have the logs:DeleteIndexPolicy and
// logs:DeleteAccountPolicy permissions.
//
// # If you delete a field index policy that included facet configurations, those
//
// facets will no longer be available for interactive exploration in the CloudWatch
// Logs Insights console. However, facet data is retained for up to 30 days.
//
// If you delete a field index policy, the indexing of the log events that
// happened before you deleted the policy will still be used for up to 30 days to
// improve CloudWatch Logs Insights queries.
func cloudwatchlogs_DeleteAccountPolicy(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.DeleteAccountPolicyInput{
		// PolicyName: *string, // Required
		// PolicyType: types.PolicyType, // Required
	}

	if len(_cloudwatchlogsPolicyName) > 0 {
		input.PolicyName = aws.String(_cloudwatchlogsPolicyName)
	}
	if len(_cloudwatchlogsPolicyType) > 0 {
		if err := assignInputField(input, "PolicyType", _cloudwatchlogsPolicyType); err != nil {
			log.Errorf("invalid --policy-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteAccountPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the data protection policy from the specified log group.
// For more information about data protection policies, see [PutDataProtectionPolicy].
//
// [PutDataProtectionPolicy]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutDataProtectionPolicy.html
func cloudwatchlogs_DeleteDataProtectionPolicy(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.DeleteDataProtectionPolicyInput{
		// LogGroupIdentifier: *string, // Required
	}

	if len(_cloudwatchlogsLogGroupIdentifier) > 0 {
		input.LogGroupIdentifier = aws.String(_cloudwatchlogsLogGroupIdentifier)
	}

	if resp, err := client.DeleteDataProtectionPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a delivery. A delivery is a connection between a logical delivery
// source and a logical delivery destination. Deleting a delivery only deletes the
// connection between the delivery source and delivery destination. It does not
// delete the delivery destination or the delivery source.
func cloudwatchlogs_DeleteDelivery(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.DeleteDeliveryInput{
		// Id: *string, // Required
	}

	if len(_cloudwatchlogsId) > 0 {
		input.Id = aws.String(_cloudwatchlogsId)
	}

	if resp, err := client.DeleteDelivery(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a delivery destination. A delivery is a connection between a logical
// delivery source and a logical delivery destination.
//
// You can't delete a delivery destination if any current deliveries are
// associated with it. To find whether any deliveries are associated with this
// delivery destination, use the [DescribeDeliveries]operation and check the deliveryDestinationArn
// field in the results.
//
// [DescribeDeliveries]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_DescribeDeliveries.html
func cloudwatchlogs_DeleteDeliveryDestination(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.DeleteDeliveryDestinationInput{
		// Name: *string, // Required
	}

	if len(_cloudwatchlogsName) > 0 {
		input.Name = aws.String(_cloudwatchlogsName)
	}

	if resp, err := client.DeleteDeliveryDestination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a delivery destination policy. For more information about these
// policies, see [PutDeliveryDestinationPolicy].
//
// [PutDeliveryDestinationPolicy]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutDeliveryDestinationPolicy.html
func cloudwatchlogs_DeleteDeliveryDestinationPolicy(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.DeleteDeliveryDestinationPolicyInput{
		// DeliveryDestinationName: *string, // Required
	}

	if len(_cloudwatchlogsDeliveryDestinationName) > 0 {
		input.DeliveryDestinationName = aws.String(_cloudwatchlogsDeliveryDestinationName)
	}

	if resp, err := client.DeleteDeliveryDestinationPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a delivery source. A delivery is a connection between a logical
// delivery source and a logical delivery destination.
//
// You can't delete a delivery source if any current deliveries are associated
// with it. To find whether any deliveries are associated with this delivery
// source, use the [DescribeDeliveries]operation and check the deliverySourceName field in the results.
//
// [DescribeDeliveries]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_DescribeDeliveries.html
func cloudwatchlogs_DeleteDeliverySource(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.DeleteDeliverySourceInput{
		// Name: *string, // Required
	}

	if len(_cloudwatchlogsName) > 0 {
		input.Name = aws.String(_cloudwatchlogsName)
	}

	if resp, err := client.DeleteDeliverySource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified destination, and eventually disables all the subscription
// filters that publish to it. This operation does not delete the physical resource
// encapsulated by the destination.
func cloudwatchlogs_DeleteDestination(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.DeleteDestinationInput{
		// DestinationName: *string, // Required
	}

	if len(_cloudwatchlogsDestinationName) > 0 {
		input.DestinationName = aws.String(_cloudwatchlogsDestinationName)
	}

	if resp, err := client.DeleteDestination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a log-group level field index policy that was applied to a single log
// group. The indexing of the log events that happened before you delete the policy
// will still be used for as many as 30 days to improve CloudWatch Logs Insights
// queries.
//
// If the deleted policy included facet configurations, those facets will no
// longer be available for interactive exploration in the CloudWatch Logs Insights
// console for this log group. However, facet data is retained for up to 30 days.
//
// You can't use this operation to delete an account-level index policy. Instead,
// use [DeleteAccountPolicy].
//
// If you delete a log-group level field index policy and there is an
// account-level field index policy, in a few minutes the log group begins using
// that account-wide policy to index new incoming log events. This operation only
// affects log group-level policies, including any facet configurations, and
// preserves any data source-based account policies that may apply to the log
// group.
//
// [DeleteAccountPolicy]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_DeleteAccountPolicy.html
func cloudwatchlogs_DeleteIndexPolicy(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.DeleteIndexPolicyInput{
		// LogGroupIdentifier: *string, // Required
	}

	if len(_cloudwatchlogsLogGroupIdentifier) > 0 {
		input.LogGroupIdentifier = aws.String(_cloudwatchlogsLogGroupIdentifier)
	}

	if resp, err := client.DeleteIndexPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the integration between CloudWatch Logs and OpenSearch Service. If your
// integration has active vended logs dashboards, you must specify true for the
// force parameter, otherwise the operation will fail. If you delete the
// integration by setting force to true , all your vended logs dashboards powered
// by OpenSearch Service will be deleted and the data that was on them will no
// longer be accessible.
func cloudwatchlogs_DeleteIntegration(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.DeleteIntegrationInput{
		// IntegrationName: *string, // Required
	}

	if len(_cloudwatchlogsIntegrationName) > 0 {
		input.IntegrationName = aws.String(_cloudwatchlogsIntegrationName)
	}
	if len(_cloudwatchlogsForce) > 0 {
		if err := assignInputField(input, "Force", _cloudwatchlogsForce); err != nil {
			log.Errorf("invalid --force: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteIntegration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified CloudWatch Logs anomaly detector.
func cloudwatchlogs_DeleteLogAnomalyDetector(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.DeleteLogAnomalyDetectorInput{
		// AnomalyDetectorArn: *string, // Required
	}

	if len(_cloudwatchlogsAnomalyDetectorArn) > 0 {
		input.AnomalyDetectorArn = aws.String(_cloudwatchlogsAnomalyDetectorArn)
	}

	if resp, err := client.DeleteLogAnomalyDetector(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified log group and permanently deletes all the archived log
// events associated with the log group.
func cloudwatchlogs_DeleteLogGroup(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.DeleteLogGroupInput{
		// LogGroupName: *string, // Required
	}

	if len(_cloudwatchlogsLogGroupName) > 0 {
		input.LogGroupName = aws.String(_cloudwatchlogsLogGroupName)
	}

	if resp, err := client.DeleteLogGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified log stream and permanently deletes all the archived log
// events associated with the log stream.
func cloudwatchlogs_DeleteLogStream(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.DeleteLogStreamInput{
		// LogGroupName: *string, // Required
		// LogStreamName: *string, // Required
	}

	if len(_cloudwatchlogsLogGroupName) > 0 {
		input.LogGroupName = aws.String(_cloudwatchlogsLogGroupName)
	}
	if len(_cloudwatchlogsLogStreamName) > 0 {
		input.LogStreamName = aws.String(_cloudwatchlogsLogStreamName)
	}

	if resp, err := client.DeleteLogStream(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified metric filter.
func cloudwatchlogs_DeleteMetricFilter(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.DeleteMetricFilterInput{
		// FilterName: *string, // Required
		// LogGroupName: *string, // Required
	}

	if len(_cloudwatchlogsFilterName) > 0 {
		input.FilterName = aws.String(_cloudwatchlogsFilterName)
	}
	if len(_cloudwatchlogsLogGroupName) > 0 {
		input.LogGroupName = aws.String(_cloudwatchlogsLogGroupName)
	}

	if resp, err := client.DeleteMetricFilter(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a saved CloudWatch Logs Insights query definition. A query definition
// contains details about a saved CloudWatch Logs Insights query.
//
// Each DeleteQueryDefinition operation can delete one query definition.
//
// You must have the logs:DeleteQueryDefinition permission to be able to perform
// this operation.
func cloudwatchlogs_DeleteQueryDefinition(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.DeleteQueryDefinitionInput{
		// QueryDefinitionId: *string, // Required
	}

	if len(_cloudwatchlogsQueryDefinitionId) > 0 {
		input.QueryDefinitionId = aws.String(_cloudwatchlogsQueryDefinitionId)
	}

	if resp, err := client.DeleteQueryDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a resource policy from this account. This revokes the access of the
// identities in that policy to put log events to this account.
func cloudwatchlogs_DeleteResourcePolicy(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.DeleteResourcePolicyInput{}

	if len(_cloudwatchlogsExpectedRevisionId) > 0 {
		input.ExpectedRevisionId = aws.String(_cloudwatchlogsExpectedRevisionId)
	}
	if len(_cloudwatchlogsPolicyName) > 0 {
		input.PolicyName = aws.String(_cloudwatchlogsPolicyName)
	}
	if len(_cloudwatchlogsResourceArn) > 0 {
		input.ResourceArn = aws.String(_cloudwatchlogsResourceArn)
	}

	if resp, err := client.DeleteResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified retention policy.
// Log events do not expire if they belong to log groups without a retention
// policy.
func cloudwatchlogs_DeleteRetentionPolicy(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.DeleteRetentionPolicyInput{
		// LogGroupName: *string, // Required
	}

	if len(_cloudwatchlogsLogGroupName) > 0 {
		input.LogGroupName = aws.String(_cloudwatchlogsLogGroupName)
	}

	if resp, err := client.DeleteRetentionPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a scheduled query and stops all future executions. This operation also
// removes any configured actions and associated resources.
func cloudwatchlogs_DeleteScheduledQuery(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.DeleteScheduledQueryInput{
		// Identifier: *string, // Required
	}

	if len(_cloudwatchlogsIdentifier) > 0 {
		input.Identifier = aws.String(_cloudwatchlogsIdentifier)
	}

	if resp, err := client.DeleteScheduledQuery(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified subscription filter.
func cloudwatchlogs_DeleteSubscriptionFilter(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.DeleteSubscriptionFilterInput{
		// FilterName: *string, // Required
		// LogGroupName: *string, // Required
	}

	if len(_cloudwatchlogsFilterName) > 0 {
		input.FilterName = aws.String(_cloudwatchlogsFilterName)
	}
	if len(_cloudwatchlogsLogGroupName) > 0 {
		input.LogGroupName = aws.String(_cloudwatchlogsLogGroupName)
	}

	if resp, err := client.DeleteSubscriptionFilter(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the log transformer for the specified log group. As soon as you do
// this, the transformation of incoming log events according to that transformer
// stops. If this account has an account-level transformer that applies to this log
// group, the log group begins using that account-level transformer when this
// log-group level transformer is deleted.
//
// After you delete a transformer, be sure to edit any metric filters or
// subscription filters that relied on the transformed versions of the log events.
func cloudwatchlogs_DeleteTransformer(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.DeleteTransformerInput{
		// LogGroupIdentifier: *string, // Required
	}

	if len(_cloudwatchlogsLogGroupIdentifier) > 0 {
		input.LogGroupIdentifier = aws.String(_cloudwatchlogsLogGroupIdentifier)
	}

	if resp, err := client.DeleteTransformer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of all CloudWatch Logs account policies in the account.
// To use this operation, you must be signed on with the correct permissions
// depending on the type of policy that you are retrieving information for.
//
// - To see data protection policies, you must have the
// logs:GetDataProtectionPolicy and logs:DescribeAccountPolicies permissions.
//
// - To see subscription filter policies, you must have the
// logs:DescribeSubscriptionFilters and logs:DescribeAccountPolicies permissions.
//
// - To see transformer policies, you must have the logs:GetTransformer and
// logs:DescribeAccountPolicies permissions.
//
// - To see field index policies, you must have the logs:DescribeIndexPolicies
// and logs:DescribeAccountPolicies permissions.
func cloudwatchlogs_DescribeAccountPolicies(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.DescribeAccountPoliciesInput{
		// PolicyType: types.PolicyType, // Required
	}

	if len(_cloudwatchlogsPolicyType) > 0 {
		if err := assignInputField(input, "PolicyType", _cloudwatchlogsPolicyType); err != nil {
			log.Errorf("invalid --policy-type: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsAccountIdentifiers) > 0 {
		input.AccountIdentifiers = append([]string(nil), _cloudwatchlogsAccountIdentifiers...)
	}
	if len(_cloudwatchlogsNextToken) > 0 {
		input.NextToken = aws.String(_cloudwatchlogsNextToken)
	}
	if len(_cloudwatchlogsPolicyName) > 0 {
		input.PolicyName = aws.String(_cloudwatchlogsPolicyName)
	}

	if resp, err := client.DescribeAccountPolicies(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Use this operation to return the valid and default values that are used when
// creating delivery sources, delivery destinations, and deliveries. For more
// information about deliveries, see [CreateDelivery].
//
// [CreateDelivery]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_CreateDelivery.html
func cloudwatchlogs_DescribeConfigurationTemplates(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.DescribeConfigurationTemplatesInput{}

	if len(_cloudwatchlogsDeliveryDestinationTypes) > 0 {
		if err := assignInputField(input, "DeliveryDestinationTypes", _cloudwatchlogsDeliveryDestinationTypes); err != nil {
			log.Errorf("invalid --delivery-destination-types: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsLimit) > 0 {
		if err := assignInputField(input, "Limit", _cloudwatchlogsLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsLogTypes) > 0 {
		input.LogTypes = append([]string(nil), _cloudwatchlogsLogTypes...)
	}
	if len(_cloudwatchlogsNextToken) > 0 {
		input.NextToken = aws.String(_cloudwatchlogsNextToken)
	}
	if len(_cloudwatchlogsResourceTypes) > 0 {
		input.ResourceTypes = append([]string(nil), _cloudwatchlogsResourceTypes...)
	}
	if len(_cloudwatchlogsService) > 0 {
		input.Service = aws.String(_cloudwatchlogsService)
	}

	if disablePaginator() {
		if resp, err := client.DescribeConfigurationTemplates(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudwatchlogs.DescribeConfigurationTemplatesOutput
	p := cloudwatchlogs.NewDescribeConfigurationTemplatesPaginator(client, input)
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

// Retrieves a list of the deliveries that have been created in the account.
// A delivery is a connection between a [delivery source] and a [delivery destination].
//
// A delivery source represents an Amazon Web Services resource that sends logs to
// an logs delivery destination. The destination can be CloudWatch Logs, Amazon S3,
// Firehose or X-Ray. Only some Amazon Web Services services support being
// configured as a delivery source. These services are listed in [Enable logging from Amazon Web Services services.]
//
// [delivery destination]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutDeliveryDestination.html
// [delivery source]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutDeliverySource.html
// [Enable logging from Amazon Web Services services.]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-logs-and-resource-policy.html
func cloudwatchlogs_DescribeDeliveries(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.DescribeDeliveriesInput{}

	if len(_cloudwatchlogsLimit) > 0 {
		if err := assignInputField(input, "Limit", _cloudwatchlogsLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsNextToken) > 0 {
		input.NextToken = aws.String(_cloudwatchlogsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeDeliveries(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudwatchlogs.DescribeDeliveriesOutput
	p := cloudwatchlogs.NewDescribeDeliveriesPaginator(client, input)
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

// Retrieves a list of the delivery destinations that have been created in the
// account.
func cloudwatchlogs_DescribeDeliveryDestinations(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.DescribeDeliveryDestinationsInput{}

	if len(_cloudwatchlogsLimit) > 0 {
		if err := assignInputField(input, "Limit", _cloudwatchlogsLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsNextToken) > 0 {
		input.NextToken = aws.String(_cloudwatchlogsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeDeliveryDestinations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudwatchlogs.DescribeDeliveryDestinationsOutput
	p := cloudwatchlogs.NewDescribeDeliveryDestinationsPaginator(client, input)
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

// Retrieves a list of the delivery sources that have been created in the account.
func cloudwatchlogs_DescribeDeliverySources(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.DescribeDeliverySourcesInput{}

	if len(_cloudwatchlogsLimit) > 0 {
		if err := assignInputField(input, "Limit", _cloudwatchlogsLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsNextToken) > 0 {
		input.NextToken = aws.String(_cloudwatchlogsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeDeliverySources(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudwatchlogs.DescribeDeliverySourcesOutput
	p := cloudwatchlogs.NewDescribeDeliverySourcesPaginator(client, input)
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

// Lists all your destinations. The results are ASCII-sorted by destination name.
func cloudwatchlogs_DescribeDestinations(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.DescribeDestinationsInput{}

	if len(_cloudwatchlogsDestinationNamePrefix) > 0 {
		input.DestinationNamePrefix = aws.String(_cloudwatchlogsDestinationNamePrefix)
	}
	if len(_cloudwatchlogsLimit) > 0 {
		if err := assignInputField(input, "Limit", _cloudwatchlogsLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsNextToken) > 0 {
		input.NextToken = aws.String(_cloudwatchlogsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeDestinations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudwatchlogs.DescribeDestinationsOutput
	p := cloudwatchlogs.NewDescribeDestinationsPaginator(client, input)
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

// Lists the specified export tasks. You can list all your export tasks or filter
// the results based on task ID or task status.
func cloudwatchlogs_DescribeExportTasks(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.DescribeExportTasksInput{}

	if len(_cloudwatchlogsLimit) > 0 {
		if err := assignInputField(input, "Limit", _cloudwatchlogsLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsNextToken) > 0 {
		input.NextToken = aws.String(_cloudwatchlogsNextToken)
	}
	if len(_cloudwatchlogsStatusCode) > 0 {
		if err := assignInputField(input, "StatusCode", _cloudwatchlogsStatusCode); err != nil {
			log.Errorf("invalid --status-code: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsTaskId) > 0 {
		input.TaskId = aws.String(_cloudwatchlogsTaskId)
	}

	if resp, err := client.DescribeExportTasks(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of custom and default field indexes which are discovered in log
// data. For more information about field index policies, see [PutIndexPolicy].
//
// [PutIndexPolicy]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutIndexPolicy.html
func cloudwatchlogs_DescribeFieldIndexes(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.DescribeFieldIndexesInput{
		// LogGroupIdentifiers: []string, // Required
	}

	if len(_cloudwatchlogsLogGroupIdentifiers) > 0 {
		input.LogGroupIdentifiers = append([]string(nil), _cloudwatchlogsLogGroupIdentifiers...)
	}
	if len(_cloudwatchlogsNextToken) > 0 {
		input.NextToken = aws.String(_cloudwatchlogsNextToken)
	}

	if resp, err := client.DescribeFieldIndexes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets detailed information about the individual batches within an import task,
// including their status and any error messages. For CloudTrail Event Data Store
// sources, a batch refers to a subset of stored events grouped by their eventTime.
func cloudwatchlogs_DescribeImportTaskBatches(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.DescribeImportTaskBatchesInput{
		// ImportId: *string, // Required
	}

	if len(_cloudwatchlogsImportId) > 0 {
		input.ImportId = aws.String(_cloudwatchlogsImportId)
	}
	if len(_cloudwatchlogsBatchImportStatus) > 0 {
		if err := assignInputField(input, "BatchImportStatus", _cloudwatchlogsBatchImportStatus); err != nil {
			log.Errorf("invalid --batch-import-status: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsLimit) > 0 {
		if err := assignInputField(input, "Limit", _cloudwatchlogsLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsNextToken) > 0 {
		input.NextToken = aws.String(_cloudwatchlogsNextToken)
	}

	if resp, err := client.DescribeImportTaskBatches(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists and describes import tasks, with optional filtering by import status and
// source ARN.
func cloudwatchlogs_DescribeImportTasks(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.DescribeImportTasksInput{}

	if len(_cloudwatchlogsImportId) > 0 {
		input.ImportId = aws.String(_cloudwatchlogsImportId)
	}
	if len(_cloudwatchlogsImportSourceArn) > 0 {
		input.ImportSourceArn = aws.String(_cloudwatchlogsImportSourceArn)
	}
	if len(_cloudwatchlogsImportStatus) > 0 {
		if err := assignInputField(input, "ImportStatus", _cloudwatchlogsImportStatus); err != nil {
			log.Errorf("invalid --import-status: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsLimit) > 0 {
		if err := assignInputField(input, "Limit", _cloudwatchlogsLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsNextToken) > 0 {
		input.NextToken = aws.String(_cloudwatchlogsNextToken)
	}

	if resp, err := client.DescribeImportTasks(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the field index policies of the specified log group. For more
// information about field index policies, see [PutIndexPolicy].
//
// If a specified log group has a log-group level index policy, that policy is
// returned by this operation.
//
// If a specified log group doesn't have a log-group level index policy, but an
// account-wide index policy applies to it, that account-wide policy is returned by
// this operation.
//
// To find information about only account-level policies, use [DescribeAccountPolicies] instead.
//
// [PutIndexPolicy]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutIndexPolicy.html
// [DescribeAccountPolicies]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_DescribeAccountPolicies.html
func cloudwatchlogs_DescribeIndexPolicies(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.DescribeIndexPoliciesInput{
		// LogGroupIdentifiers: []string, // Required
	}

	if len(_cloudwatchlogsLogGroupIdentifiers) > 0 {
		input.LogGroupIdentifiers = append([]string(nil), _cloudwatchlogsLogGroupIdentifiers...)
	}
	if len(_cloudwatchlogsNextToken) > 0 {
		input.NextToken = aws.String(_cloudwatchlogsNextToken)
	}

	if resp, err := client.DescribeIndexPolicies(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about log groups, including data sources that ingest into
// each log group. You can return all your log groups or filter the results by
// prefix. The results are ASCII-sorted by log group name.
//
// CloudWatch Logs doesn't support IAM policies that control access to the
// DescribeLogGroups action by using the aws:ResourceTag/key-name  condition key.
// Other CloudWatch Logs actions do support the use of the
// aws:ResourceTag/key-name condition key to control access. For more information
// about using tags to control access, see [Controlling access to Amazon Web Services resources using tags].
//
// If you are using CloudWatch cross-account observability, you can use this
// operation in a monitoring account and view data from the linked source accounts.
// For more information, see [CloudWatch cross-account observability].
//
// [CloudWatch cross-account observability]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Unified-Cross-Account.html
// [Controlling access to Amazon Web Services resources using tags]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_tags.html
func cloudwatchlogs_DescribeLogGroups(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.DescribeLogGroupsInput{}

	if len(_cloudwatchlogsAccountIdentifiers) > 0 {
		input.AccountIdentifiers = append([]string(nil), _cloudwatchlogsAccountIdentifiers...)
	}
	if len(_cloudwatchlogsIncludeLinkedAccounts) > 0 {
		if err := assignInputField(input, "IncludeLinkedAccounts", _cloudwatchlogsIncludeLinkedAccounts); err != nil {
			log.Errorf("invalid --include-linked-accounts: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsLimit) > 0 {
		if err := assignInputField(input, "Limit", _cloudwatchlogsLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsLogGroupClass) > 0 {
		if err := assignInputField(input, "LogGroupClass", _cloudwatchlogsLogGroupClass); err != nil {
			log.Errorf("invalid --log-group-class: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsLogGroupIdentifiers) > 0 {
		input.LogGroupIdentifiers = append([]string(nil), _cloudwatchlogsLogGroupIdentifiers...)
	}
	if len(_cloudwatchlogsLogGroupNamePattern) > 0 {
		input.LogGroupNamePattern = aws.String(_cloudwatchlogsLogGroupNamePattern)
	}
	if len(_cloudwatchlogsLogGroupNamePrefix) > 0 {
		input.LogGroupNamePrefix = aws.String(_cloudwatchlogsLogGroupNamePrefix)
	}
	if len(_cloudwatchlogsNextToken) > 0 {
		input.NextToken = aws.String(_cloudwatchlogsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeLogGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudwatchlogs.DescribeLogGroupsOutput
	p := cloudwatchlogs.NewDescribeLogGroupsPaginator(client, input)
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

// Lists the log streams for the specified log group. You can list all the log
// streams or filter the results by prefix. You can also control how the results
// are ordered.
//
// You can specify the log group to search by using either logGroupIdentifier or
// logGroupName . You must include one of these two parameters, but you can't
// include both.
//
// This operation has a limit of 25 transactions per second, after which
// transactions are throttled.
//
// If you are using CloudWatch cross-account observability, you can use this
// operation in a monitoring account and view data from the linked source accounts.
// For more information, see [CloudWatch cross-account observability].
//
// [CloudWatch cross-account observability]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Unified-Cross-Account.html
func cloudwatchlogs_DescribeLogStreams(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.DescribeLogStreamsInput{}

	if len(_cloudwatchlogsDescending) > 0 {
		if err := assignInputField(input, "Descending", _cloudwatchlogsDescending); err != nil {
			log.Errorf("invalid --descending: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsLimit) > 0 {
		if err := assignInputField(input, "Limit", _cloudwatchlogsLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsLogGroupIdentifier) > 0 {
		input.LogGroupIdentifier = aws.String(_cloudwatchlogsLogGroupIdentifier)
	}
	if len(_cloudwatchlogsLogGroupName) > 0 {
		input.LogGroupName = aws.String(_cloudwatchlogsLogGroupName)
	}
	if len(_cloudwatchlogsLogStreamNamePrefix) > 0 {
		input.LogStreamNamePrefix = aws.String(_cloudwatchlogsLogStreamNamePrefix)
	}
	if len(_cloudwatchlogsNextToken) > 0 {
		input.NextToken = aws.String(_cloudwatchlogsNextToken)
	}
	if len(_cloudwatchlogsOrderBy) > 0 {
		if err := assignInputField(input, "OrderBy", _cloudwatchlogsOrderBy); err != nil {
			log.Errorf("invalid --order-by: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeLogStreams(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudwatchlogs.DescribeLogStreamsOutput
	p := cloudwatchlogs.NewDescribeLogStreamsPaginator(client, input)
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

// Lists the specified metric filters. You can list all of the metric filters or
// filter the results by log name, prefix, metric name, or metric namespace. The
// results are ASCII-sorted by filter name.
func cloudwatchlogs_DescribeMetricFilters(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.DescribeMetricFiltersInput{}

	if len(_cloudwatchlogsFilterNamePrefix) > 0 {
		input.FilterNamePrefix = aws.String(_cloudwatchlogsFilterNamePrefix)
	}
	if len(_cloudwatchlogsLimit) > 0 {
		if err := assignInputField(input, "Limit", _cloudwatchlogsLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsLogGroupName) > 0 {
		input.LogGroupName = aws.String(_cloudwatchlogsLogGroupName)
	}
	if len(_cloudwatchlogsMetricName) > 0 {
		input.MetricName = aws.String(_cloudwatchlogsMetricName)
	}
	if len(_cloudwatchlogsMetricNamespace) > 0 {
		input.MetricNamespace = aws.String(_cloudwatchlogsMetricNamespace)
	}
	if len(_cloudwatchlogsNextToken) > 0 {
		input.NextToken = aws.String(_cloudwatchlogsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeMetricFilters(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudwatchlogs.DescribeMetricFiltersOutput
	p := cloudwatchlogs.NewDescribeMetricFiltersPaginator(client, input)
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

// Returns a list of CloudWatch Logs Insights queries that are scheduled, running,
// or have been run recently in this account. You can request all queries or limit
// it to queries of a specific log group or queries with a certain status.
//
// This operation includes both interactive queries started directly by users and
// automated queries executed by scheduled query configurations. Scheduled query
// executions appear in the results alongside manually initiated queries, providing
// visibility into all query activity in your account.
func cloudwatchlogs_DescribeQueries(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.DescribeQueriesInput{}

	if len(_cloudwatchlogsLogGroupName) > 0 {
		input.LogGroupName = aws.String(_cloudwatchlogsLogGroupName)
	}
	if len(_cloudwatchlogsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cloudwatchlogsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsNextToken) > 0 {
		input.NextToken = aws.String(_cloudwatchlogsNextToken)
	}
	if len(_cloudwatchlogsQueryLanguage) > 0 {
		if err := assignInputField(input, "QueryLanguage", _cloudwatchlogsQueryLanguage); err != nil {
			log.Errorf("invalid --query-language: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsStatus) > 0 {
		if err := assignInputField(input, "Status", _cloudwatchlogsStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeQueries(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation returns a paginated list of your saved CloudWatch Logs Insights
// query definitions. You can retrieve query definitions from the current account
// or from a source account that is linked to the current account.
//
// You can use the queryDefinitionNamePrefix parameter to limit the results to
// only the query definitions that have names that start with a certain string.
func cloudwatchlogs_DescribeQueryDefinitions(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.DescribeQueryDefinitionsInput{}

	if len(_cloudwatchlogsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cloudwatchlogsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsNextToken) > 0 {
		input.NextToken = aws.String(_cloudwatchlogsNextToken)
	}
	if len(_cloudwatchlogsQueryDefinitionNamePrefix) > 0 {
		input.QueryDefinitionNamePrefix = aws.String(_cloudwatchlogsQueryDefinitionNamePrefix)
	}
	if len(_cloudwatchlogsQueryLanguage) > 0 {
		if err := assignInputField(input, "QueryLanguage", _cloudwatchlogsQueryLanguage); err != nil {
			log.Errorf("invalid --query-language: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeQueryDefinitions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the resource policies in this account.
func cloudwatchlogs_DescribeResourcePolicies(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.DescribeResourcePoliciesInput{}

	if len(_cloudwatchlogsLimit) > 0 {
		if err := assignInputField(input, "Limit", _cloudwatchlogsLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsNextToken) > 0 {
		input.NextToken = aws.String(_cloudwatchlogsNextToken)
	}
	if len(_cloudwatchlogsPolicyScope) > 0 {
		if err := assignInputField(input, "PolicyScope", _cloudwatchlogsPolicyScope); err != nil {
			log.Errorf("invalid --policy-scope: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsResourceArn) > 0 {
		input.ResourceArn = aws.String(_cloudwatchlogsResourceArn)
	}

	if resp, err := client.DescribeResourcePolicies(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the subscription filters for the specified log group. You can list all
// the subscription filters or filter the results by prefix. The results are
// ASCII-sorted by filter name.
func cloudwatchlogs_DescribeSubscriptionFilters(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.DescribeSubscriptionFiltersInput{
		// LogGroupName: *string, // Required
	}

	if len(_cloudwatchlogsLogGroupName) > 0 {
		input.LogGroupName = aws.String(_cloudwatchlogsLogGroupName)
	}
	if len(_cloudwatchlogsFilterNamePrefix) > 0 {
		input.FilterNamePrefix = aws.String(_cloudwatchlogsFilterNamePrefix)
	}
	if len(_cloudwatchlogsLimit) > 0 {
		if err := assignInputField(input, "Limit", _cloudwatchlogsLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsNextToken) > 0 {
		input.NextToken = aws.String(_cloudwatchlogsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeSubscriptionFilters(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudwatchlogs.DescribeSubscriptionFiltersOutput
	p := cloudwatchlogs.NewDescribeSubscriptionFiltersPaginator(client, input)
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

// Disassociates the specified KMS key from the specified log group or from all
// CloudWatch Logs Insights query results in the account.
//
// When you use DisassociateKmsKey , you specify either the logGroupName parameter
// or the resourceIdentifier parameter. You can't specify both of those parameters
// in the same operation.
//
// - Specify the logGroupName parameter to stop using the KMS key to encrypt
// future log events ingested and stored in the log group. Instead, they will be
// encrypted with the default CloudWatch Logs method. The log events that were
// ingested while the key was associated with the log group are still encrypted
// with that key. Therefore, CloudWatch Logs will need permissions for the key
// whenever that data is accessed.
//
// - Specify the resourceIdentifier parameter with the query-result resource to
// stop using the KMS key to encrypt the results of all future [StartQuery]operations in the
// account. They will instead be encrypted with the default CloudWatch Logs method.
// The results from queries that ran while the key was associated with the account
// are still encrypted with that key. Therefore, CloudWatch Logs will need
// permissions for the key whenever that data is accessed.
//
// It can take up to 5 minutes for this operation to take effect.
//
// [StartQuery]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_StartQuery.html
func cloudwatchlogs_DisassociateKmsKey(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.DisassociateKmsKeyInput{}

	if len(_cloudwatchlogsLogGroupName) > 0 {
		input.LogGroupName = aws.String(_cloudwatchlogsLogGroupName)
	}
	if len(_cloudwatchlogsResourceIdentifier) > 0 {
		input.ResourceIdentifier = aws.String(_cloudwatchlogsResourceIdentifier)
	}

	if resp, err := client.DisassociateKmsKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates a data source from an S3 Table Integration, removing query access
// and deleting all associated data from the integration.
func cloudwatchlogs_DisassociateSourceFromS3TableIntegration(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.DisassociateSourceFromS3TableIntegrationInput{
		// Identifier: *string, // Required
	}

	if len(_cloudwatchlogsIdentifier) > 0 {
		input.Identifier = aws.String(_cloudwatchlogsIdentifier)
	}

	if resp, err := client.DisassociateSourceFromS3TableIntegration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists log events from the specified log group. You can list all the log events
// or filter the results using one or more of the following:
//
// - A filter pattern
//
// - A time range
//
// - The log stream name, or a log stream name prefix that matches multiple log
// streams
//
// You must have the logs:FilterLogEvents permission to perform this operation.
//
// You can specify the log group to search by using either logGroupIdentifier or
// logGroupName . You must include one of these two parameters, but you can't
// include both.
//
// FilterLogEvents is a paginated operation. Each page returned can contain up to
// 1 MB of log events or up to 10,000 log events. A returned page might only be
// partially full, or even empty. For example, if the result of a query would
// return 15,000 log events, the first page isn't guaranteed to have 10,000 log
// events even if they all fit into 1 MB.
//
// Partially full or empty pages don't necessarily mean that pagination is
// finished. If the results include a nextToken , there might be more log events
// available. You can return these additional log events by providing the nextToken
// in a subsequent FilterLogEvents operation. If the results don't include a
// nextToken , then pagination is finished.
//
// Specifying the limit parameter only guarantees that a single page doesn't
// return more log events than the specified limit, but it might return fewer
// events than the limit. This is the expected API behavior.
//
// The returned log events are sorted by event timestamp, the timestamp when the
// event was ingested by CloudWatch Logs, and the ID of the PutLogEvents request.
//
// If you are using CloudWatch cross-account observability, you can use this
// operation in a monitoring account and view data from the linked source accounts.
// For more information, see [CloudWatch cross-account observability].
//
// If you are using [log transformation], the FilterLogEvents operation returns only the original
// versions of log events, before they were transformed. To view the transformed
// versions, you must use a [CloudWatch Logs query.]
//
// [log transformation]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation.html
// [CloudWatch cross-account observability]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Unified-Cross-Account.html
// [CloudWatch Logs query.]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AnalyzingLogData.html
func cloudwatchlogs_FilterLogEvents(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.FilterLogEventsInput{}

	if len(_cloudwatchlogsEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _cloudwatchlogsEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsFilterPattern) > 0 {
		input.FilterPattern = aws.String(_cloudwatchlogsFilterPattern)
	}
	if len(_cloudwatchlogsInterleaved) > 0 {
		if err := assignInputField(input, "Interleaved", _cloudwatchlogsInterleaved); err != nil {
			log.Errorf("invalid --interleaved: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsLimit) > 0 {
		if err := assignInputField(input, "Limit", _cloudwatchlogsLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsLogGroupIdentifier) > 0 {
		input.LogGroupIdentifier = aws.String(_cloudwatchlogsLogGroupIdentifier)
	}
	if len(_cloudwatchlogsLogGroupName) > 0 {
		input.LogGroupName = aws.String(_cloudwatchlogsLogGroupName)
	}
	if len(_cloudwatchlogsLogStreamNamePrefix) > 0 {
		input.LogStreamNamePrefix = aws.String(_cloudwatchlogsLogStreamNamePrefix)
	}
	if len(_cloudwatchlogsLogStreamNames) > 0 {
		input.LogStreamNames = append([]string(nil), _cloudwatchlogsLogStreamNames...)
	}
	if len(_cloudwatchlogsNextToken) > 0 {
		input.NextToken = aws.String(_cloudwatchlogsNextToken)
	}
	if len(_cloudwatchlogsStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _cloudwatchlogsStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsUnmask) > 0 {
		if err := assignInputField(input, "Unmask", _cloudwatchlogsUnmask); err != nil {
			log.Errorf("invalid --unmask: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.FilterLogEvents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudwatchlogs.FilterLogEventsOutput
	p := cloudwatchlogs.NewFilterLogEventsPaginator(client, input)
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

// Returns information about a log group data protection policy.
func cloudwatchlogs_GetDataProtectionPolicy(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.GetDataProtectionPolicyInput{
		// LogGroupIdentifier: *string, // Required
	}

	if len(_cloudwatchlogsLogGroupIdentifier) > 0 {
		input.LogGroupIdentifier = aws.String(_cloudwatchlogsLogGroupIdentifier)
	}

	if resp, err := client.GetDataProtectionPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns complete information about one logical delivery. A delivery is a
// connection between a [delivery source]and a [delivery destination].
//
// A delivery source represents an Amazon Web Services resource that sends logs to
// an logs delivery destination. The destination can be CloudWatch Logs, Amazon S3,
// or Firehose. Only some Amazon Web Services services support being configured as
// a delivery source. These services are listed in [Enable logging from Amazon Web Services services.]
//
// You need to specify the delivery id in this operation. You can find the IDs of
// the deliveries in your account with the [DescribeDeliveries]operation.
//
// [delivery destination]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutDeliveryDestination.html
// [delivery source]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutDeliverySource.html
// [Enable logging from Amazon Web Services services.]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-logs-and-resource-policy.html
// [DescribeDeliveries]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_DescribeDeliveries.html
func cloudwatchlogs_GetDelivery(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.GetDeliveryInput{
		// Id: *string, // Required
	}

	if len(_cloudwatchlogsId) > 0 {
		input.Id = aws.String(_cloudwatchlogsId)
	}

	if resp, err := client.GetDelivery(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves complete information about one delivery destination.
func cloudwatchlogs_GetDeliveryDestination(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.GetDeliveryDestinationInput{
		// Name: *string, // Required
	}

	if len(_cloudwatchlogsName) > 0 {
		input.Name = aws.String(_cloudwatchlogsName)
	}

	if resp, err := client.GetDeliveryDestination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the delivery destination policy assigned to the delivery destination
// that you specify. For more information about delivery destinations and their
// policies, see [PutDeliveryDestinationPolicy].
//
// [PutDeliveryDestinationPolicy]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutDeliveryDestinationPolicy.html
func cloudwatchlogs_GetDeliveryDestinationPolicy(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.GetDeliveryDestinationPolicyInput{
		// DeliveryDestinationName: *string, // Required
	}

	if len(_cloudwatchlogsDeliveryDestinationName) > 0 {
		input.DeliveryDestinationName = aws.String(_cloudwatchlogsDeliveryDestinationName)
	}

	if resp, err := client.GetDeliveryDestinationPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves complete information about one delivery source.
func cloudwatchlogs_GetDeliverySource(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.GetDeliverySourceInput{
		// Name: *string, // Required
	}

	if len(_cloudwatchlogsName) > 0 {
		input.Name = aws.String(_cloudwatchlogsName)
	}

	if resp, err := client.GetDeliverySource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about one integration between CloudWatch Logs and
// OpenSearch Service.
func cloudwatchlogs_GetIntegration(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.GetIntegrationInput{
		// IntegrationName: *string, // Required
	}

	if len(_cloudwatchlogsIntegrationName) > 0 {
		input.IntegrationName = aws.String(_cloudwatchlogsIntegrationName)
	}

	if resp, err := client.GetIntegration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the log anomaly detector that you specify. The KMS
// key ARN detected is valid.
func cloudwatchlogs_GetLogAnomalyDetector(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.GetLogAnomalyDetectorInput{
		// AnomalyDetectorArn: *string, // Required
	}

	if len(_cloudwatchlogsAnomalyDetectorArn) > 0 {
		input.AnomalyDetectorArn = aws.String(_cloudwatchlogsAnomalyDetectorArn)
	}

	if resp, err := client.GetLogAnomalyDetector(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists log events from the specified log stream. You can list all of the log
// events or filter using a time range.
//
// GetLogEvents is a paginated operation. Each page returned can contain up to 1
// MB of log events or up to 10,000 log events. A returned page might only be
// partially full, or even empty. For example, if the result of a query would
// return 15,000 log events, the first page isn't guaranteed to have 10,000 log
// events even if they all fit into 1 MB.
//
// Partially full or empty pages don't necessarily mean that pagination is
// finished. As long as the nextBackwardToken or nextForwardToken returned is NOT
// equal to the nextToken that you passed into the API call, there might be more
// log events available. The token that you use depends on the direction you want
// to move in along the log stream. The returned tokens are never null.
//
// If you set startFromHead to true and you don’t include endTime in your request,
// you can end up in a situation where the pagination doesn't terminate. This can
// happen when the new log events are being added to the target log streams faster
// than they are being read. This situation is a good use case for the CloudWatch
// Logs [Live Tail]feature.
//
// If you are using CloudWatch cross-account observability, you can use this
// operation in a monitoring account and view data from the linked source accounts.
// For more information, see [CloudWatch cross-account observability].
//
// You can specify the log group to search by using either logGroupIdentifier or
// logGroupName . You must include one of these two parameters, but you can't
// include both.
//
// If you are using [log transformation], the GetLogEvents operation returns only the original
// versions of log events, before they were transformed. To view the transformed
// versions, you must use a [CloudWatch Logs query.]
//
// [log transformation]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation.html
// [CloudWatch cross-account observability]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Unified-Cross-Account.html
// [CloudWatch Logs query.]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AnalyzingLogData.html
// [Live Tail]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatchLogs_LiveTail.html
func cloudwatchlogs_GetLogEvents(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.GetLogEventsInput{
		// LogStreamName: *string, // Required
	}

	if len(_cloudwatchlogsLogStreamName) > 0 {
		input.LogStreamName = aws.String(_cloudwatchlogsLogStreamName)
	}
	if len(_cloudwatchlogsEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _cloudwatchlogsEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsLimit) > 0 {
		if err := assignInputField(input, "Limit", _cloudwatchlogsLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsLogGroupIdentifier) > 0 {
		input.LogGroupIdentifier = aws.String(_cloudwatchlogsLogGroupIdentifier)
	}
	if len(_cloudwatchlogsLogGroupName) > 0 {
		input.LogGroupName = aws.String(_cloudwatchlogsLogGroupName)
	}
	if len(_cloudwatchlogsNextToken) > 0 {
		input.NextToken = aws.String(_cloudwatchlogsNextToken)
	}
	if len(_cloudwatchlogsStartFromHead) > 0 {
		if err := assignInputField(input, "StartFromHead", _cloudwatchlogsStartFromHead); err != nil {
			log.Errorf("invalid --start-from-head: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _cloudwatchlogsStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsUnmask) > 0 {
		if err := assignInputField(input, "Unmask", _cloudwatchlogsUnmask); err != nil {
			log.Errorf("invalid --unmask: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.GetLogEvents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudwatchlogs.GetLogEventsOutput
	p := cloudwatchlogs.NewGetLogEventsPaginator(client, input)
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

// Discovers available fields for a specific data source and type. The response
// includes any field modifications introduced through pipelines, such as new
// fields or changed field types.
func cloudwatchlogs_GetLogFields(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.GetLogFieldsInput{
		// DataSourceName: *string, // Required
		// DataSourceType: *string, // Required
	}

	if len(_cloudwatchlogsDataSourceName) > 0 {
		input.DataSourceName = aws.String(_cloudwatchlogsDataSourceName)
	}
	if len(_cloudwatchlogsDataSourceType) > 0 {
		input.DataSourceType = aws.String(_cloudwatchlogsDataSourceType)
	}

	if resp, err := client.GetLogFields(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of the fields that are included in log events in the specified
// log group. Includes the percentage of log events that contain each field. The
// search is limited to a time period that you specify.
//
// This operation is used for discovering fields within log group events. For
// discovering fields across data sources, use the GetLogFields operation.
//
// You can specify the log group to search by using either logGroupIdentifier or
// logGroupName . You must specify one of these parameters, but you can't specify
// both.
//
// In the results, fields that start with (at) are fields generated by CloudWatch
// Logs. For example, (at)timestamp is the timestamp of each log event. For more
// information about the fields that are generated by CloudWatch logs, see [Supported Logs and Discovered Fields].
//
// The response results are sorted by the frequency percentage, starting with the
// highest percentage.
//
// If you are using CloudWatch cross-account observability, you can use this
// operation in a monitoring account and view data from the linked source accounts.
// For more information, see [CloudWatch cross-account observability].
//
// [CloudWatch cross-account observability]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Unified-Cross-Account.html
// [Supported Logs and Discovered Fields]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_AnalyzeLogData-discoverable-fields.html
func cloudwatchlogs_GetLogGroupFields(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.GetLogGroupFieldsInput{}

	if len(_cloudwatchlogsLogGroupIdentifier) > 0 {
		input.LogGroupIdentifier = aws.String(_cloudwatchlogsLogGroupIdentifier)
	}
	if len(_cloudwatchlogsLogGroupName) > 0 {
		input.LogGroupName = aws.String(_cloudwatchlogsLogGroupName)
	}
	if len(_cloudwatchlogsTime) > 0 {
		if err := assignInputField(input, "Time", _cloudwatchlogsTime); err != nil {
			log.Errorf("invalid --time: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetLogGroupFields(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a large logging object (LLO) and streams it back. This API is used to
// fetch the content of large portions of log events that have been ingested
// through the PutOpenTelemetryLogs API. When log events contain fields that would
// cause the total event size to exceed 1MB, CloudWatch Logs automatically
// processes up to 10 fields, starting with the largest fields. Each field is
// truncated as needed to keep the total event size as close to 1MB as possible.
// The excess portions are stored as Large Log Objects (LLOs) and these fields are
// processed separately and LLO reference system fields (in the format
// (at)ptr.$[path.to.field] ) are added. The path in the reference field reflects the
// original JSON structure where the large field was located. For example, this
// could be (at)ptr.$['input']['message'] , (at)ptr.$['AAA']['BBB']['CCC']['DDD'] ,
// (at)ptr.$['AAA'] , or any other path matching your log structure.
//
// The GetLogObject API routes requests using SDK host prefix injection. SDK
// versions released before April 1, 2026 route to
// streaming-logs.Region.amazonaws.com , which does not support VPC endpoints. SDK
// versions released on or after April 1, 2026 route to
// stream-logs.Region.amazonaws.com , which supports VPC endpoints. To set up a VPC
// endpoint for this API, see [Creating a VPC endpoint for CloudWatch Logs].
//
// [Creating a VPC endpoint for CloudWatch Logs]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/cloudwatch-logs-and-interface-VPC.html#create-VPC-endpoint-for-CloudWatchLogs
func cloudwatchlogs_GetLogObject(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.GetLogObjectInput{
		// LogObjectPointer: *string, // Required
	}

	if len(_cloudwatchlogsLogObjectPointer) > 0 {
		input.LogObjectPointer = aws.String(_cloudwatchlogsLogObjectPointer)
	}
	if len(_cloudwatchlogsUnmask) > 0 {
		if err := assignInputField(input, "Unmask", _cloudwatchlogsUnmask); err != nil {
			log.Errorf("invalid --unmask: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetLogObject(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves all of the fields and values of a single log event. All fields are
// retrieved, even if the original query that produced the logRecordPointer
// retrieved only a subset of fields. Fields are returned as field name/field value
// pairs.
//
// The full unparsed log event is returned within (at)message .
func cloudwatchlogs_GetLogRecord(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.GetLogRecordInput{
		// LogRecordPointer: *string, // Required
	}

	if len(_cloudwatchlogsLogRecordPointer) > 0 {
		input.LogRecordPointer = aws.String(_cloudwatchlogsLogRecordPointer)
	}
	if len(_cloudwatchlogsUnmask) > 0 {
		if err := assignInputField(input, "Unmask", _cloudwatchlogsUnmask); err != nil {
			log.Errorf("invalid --unmask: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetLogRecord(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the results from the specified query.
// Only the fields requested in the query are returned, along with a (at)ptr field,
// which is the identifier for the log record. You can use the value of (at)ptr in a [GetLogRecord]
// operation to get the full log record.
//
// GetQueryResults does not start running a query. To run a query, use [StartQuery]. For more
// information about how long results of previous queries are available, see [CloudWatch Logs quotas].
//
// If the value of the Status field in the output is Running , this operation
// returns only partial results. If you see a value of Scheduled or Running for
// the status, you can retry the operation later to see the final results.
//
// This operation is used both for retrieving results from interactive queries and
// from automated scheduled query executions. Scheduled queries use GetQueryResults
// internally to retrieve query results for processing and delivery to configured
// destinations.
//
// If you are using CloudWatch cross-account observability, you can use this
// operation in a monitoring account to start queries in linked source accounts.
// For more information, see [CloudWatch cross-account observability].
//
// [CloudWatch cross-account observability]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Unified-Cross-Account.html
// [GetLogRecord]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_GetLogRecord.html
// [CloudWatch Logs quotas]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/cloudwatch_limits_cwl.html
// [StartQuery]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_StartQuery.html
func cloudwatchlogs_GetQueryResults(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.GetQueryResultsInput{
		// QueryId: *string, // Required
	}

	if len(_cloudwatchlogsQueryId) > 0 {
		input.QueryId = aws.String(_cloudwatchlogsQueryId)
	}

	if resp, err := client.GetQueryResults(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves details about a specific scheduled query, including its
// configuration, execution status, and metadata.
func cloudwatchlogs_GetScheduledQuery(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.GetScheduledQueryInput{
		// Identifier: *string, // Required
	}

	if len(_cloudwatchlogsIdentifier) > 0 {
		input.Identifier = aws.String(_cloudwatchlogsIdentifier)
	}

	if resp, err := client.GetScheduledQuery(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the execution history of a scheduled query within a specified time
// range, including query results and destination processing status.
func cloudwatchlogs_GetScheduledQueryHistory(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.GetScheduledQueryHistoryInput{
		// EndTime: *int64, // Required
		// Identifier: *string, // Required
		// StartTime: *int64, // Required
	}

	if len(_cloudwatchlogsEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _cloudwatchlogsEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsIdentifier) > 0 {
		input.Identifier = aws.String(_cloudwatchlogsIdentifier)
	}
	if len(_cloudwatchlogsStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _cloudwatchlogsStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsExecutionStatuses) > 0 {
		if err := assignInputField(input, "ExecutionStatuses", _cloudwatchlogsExecutionStatuses); err != nil {
			log.Errorf("invalid --execution-statuses: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cloudwatchlogsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsNextToken) > 0 {
		input.NextToken = aws.String(_cloudwatchlogsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetScheduledQueryHistory(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudwatchlogs.GetScheduledQueryHistoryOutput
	p := cloudwatchlogs.NewGetScheduledQueryHistoryPaginator(client, input)
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

// Returns the information about the log transformer associated with this log
// group.
//
// This operation returns data only for transformers created at the log group
// level. To get information for an account-level transformer, use [DescribeAccountPolicies].
//
// [DescribeAccountPolicies]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_DescribeAccountPolicies.html
func cloudwatchlogs_GetTransformer(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.GetTransformerInput{
		// LogGroupIdentifier: *string, // Required
	}

	if len(_cloudwatchlogsLogGroupIdentifier) > 0 {
		input.LogGroupIdentifier = aws.String(_cloudwatchlogsLogGroupIdentifier)
	}

	if resp, err := client.GetTransformer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns an aggregate summary of all log groups in the Region grouped by
// specified data source characteristics. Supports optional filtering by log group
// class, name patterns, and data sources. If you perform this action in a
// monitoring account, you can also return aggregated summaries of log groups from
// source accounts that are linked to the monitoring account. For more information
// about using cross-account observability to set up monitoring accounts and source
// accounts, see [CloudWatch cross-account observability].
//
// The operation aggregates log groups by data source name and type and optionally
// format, providing counts of log groups that share these characteristics. The
// operation paginates results. By default, it returns up to 50 results and
// includes a token to retrieve more results.
//
// [CloudWatch cross-account observability]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Unified-Cross-Account.html
func cloudwatchlogs_ListAggregateLogGroupSummaries(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.ListAggregateLogGroupSummariesInput{
		// GroupBy: types.ListAggregateLogGroupSummariesGroupBy, // Required
	}

	if len(_cloudwatchlogsGroupBy) > 0 {
		if err := assignInputField(input, "GroupBy", _cloudwatchlogsGroupBy); err != nil {
			log.Errorf("invalid --group-by: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsAccountIdentifiers) > 0 {
		input.AccountIdentifiers = append([]string(nil), _cloudwatchlogsAccountIdentifiers...)
	}
	if len(_cloudwatchlogsDataSources) > 0 {
		if err := assignInputField(input, "DataSources", _cloudwatchlogsDataSources); err != nil {
			log.Errorf("invalid --data-sources: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsIncludeLinkedAccounts) > 0 {
		if err := assignInputField(input, "IncludeLinkedAccounts", _cloudwatchlogsIncludeLinkedAccounts); err != nil {
			log.Errorf("invalid --include-linked-accounts: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsLimit) > 0 {
		if err := assignInputField(input, "Limit", _cloudwatchlogsLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsLogGroupClass) > 0 {
		if err := assignInputField(input, "LogGroupClass", _cloudwatchlogsLogGroupClass); err != nil {
			log.Errorf("invalid --log-group-class: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsLogGroupNamePattern) > 0 {
		input.LogGroupNamePattern = aws.String(_cloudwatchlogsLogGroupNamePattern)
	}
	if len(_cloudwatchlogsNextToken) > 0 {
		input.NextToken = aws.String(_cloudwatchlogsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAggregateLogGroupSummaries(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudwatchlogs.ListAggregateLogGroupSummariesOutput
	p := cloudwatchlogs.NewListAggregateLogGroupSummariesPaginator(client, input)
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

// Returns a list of anomalies that log anomaly detectors have found. For details
// about the structure format of each anomaly object that is returned, see the
// example in this section.
func cloudwatchlogs_ListAnomalies(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.ListAnomaliesInput{}

	if len(_cloudwatchlogsAnomalyDetectorArn) > 0 {
		input.AnomalyDetectorArn = aws.String(_cloudwatchlogsAnomalyDetectorArn)
	}
	if len(_cloudwatchlogsLimit) > 0 {
		if err := assignInputField(input, "Limit", _cloudwatchlogsLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsNextToken) > 0 {
		input.NextToken = aws.String(_cloudwatchlogsNextToken)
	}
	if len(_cloudwatchlogsSuppressionState) > 0 {
		if err := assignInputField(input, "SuppressionState", _cloudwatchlogsSuppressionState); err != nil {
			log.Errorf("invalid --suppression-state: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListAnomalies(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudwatchlogs.ListAnomaliesOutput
	p := cloudwatchlogs.NewListAnomaliesPaginator(client, input)
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

// Returns a list of integrations between CloudWatch Logs and other services in
// this account. Currently, only one integration can be created in an account, and
// this integration must be with OpenSearch Service.
func cloudwatchlogs_ListIntegrations(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.ListIntegrationsInput{}

	if len(_cloudwatchlogsIntegrationNamePrefix) > 0 {
		input.IntegrationNamePrefix = aws.String(_cloudwatchlogsIntegrationNamePrefix)
	}
	if len(_cloudwatchlogsIntegrationStatus) > 0 {
		if err := assignInputField(input, "IntegrationStatus", _cloudwatchlogsIntegrationStatus); err != nil {
			log.Errorf("invalid --integration-status: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsIntegrationType) > 0 {
		if err := assignInputField(input, "IntegrationType", _cloudwatchlogsIntegrationType); err != nil {
			log.Errorf("invalid --integration-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListIntegrations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list of the log anomaly detectors in the account.
func cloudwatchlogs_ListLogAnomalyDetectors(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.ListLogAnomalyDetectorsInput{}

	if len(_cloudwatchlogsFilterLogGroupArn) > 0 {
		input.FilterLogGroupArn = aws.String(_cloudwatchlogsFilterLogGroupArn)
	}
	if len(_cloudwatchlogsLimit) > 0 {
		if err := assignInputField(input, "Limit", _cloudwatchlogsLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsNextToken) > 0 {
		input.NextToken = aws.String(_cloudwatchlogsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListLogAnomalyDetectors(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudwatchlogs.ListLogAnomalyDetectorsOutput
	p := cloudwatchlogs.NewListLogAnomalyDetectorsPaginator(client, input)
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

// Returns a list of log groups in the Region in your account. If you are
// performing this action in a monitoring account, you can choose to also return
// log groups from source accounts that are linked to the monitoring account. For
// more information about using cross-account observability to set up monitoring
// accounts and source accounts, see [CloudWatch cross-account observability].
//
// You can optionally filter the list by log group class, by using regular
// expressions in your request to match strings in the log group names, by using
// the fieldIndexes parameter to filter log groups based on which field indexes are
// configured, by using the dataSources parameter to filter log groups by data
// source types, and by using the fieldIndexNames parameter to filter by specific
// field index names.
//
// This operation is paginated. By default, your first use of this operation
// returns 50 results, and includes a token to use in a subsequent operation to
// return more results.
//
// [CloudWatch cross-account observability]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Unified-Cross-Account.html
func cloudwatchlogs_ListLogGroups(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.ListLogGroupsInput{}

	if len(_cloudwatchlogsAccountIdentifiers) > 0 {
		input.AccountIdentifiers = append([]string(nil), _cloudwatchlogsAccountIdentifiers...)
	}
	if len(_cloudwatchlogsDataSources) > 0 {
		if err := assignInputField(input, "DataSources", _cloudwatchlogsDataSources); err != nil {
			log.Errorf("invalid --data-sources: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsFieldIndexNames) > 0 {
		input.FieldIndexNames = append([]string(nil), _cloudwatchlogsFieldIndexNames...)
	}
	if len(_cloudwatchlogsIncludeLinkedAccounts) > 0 {
		if err := assignInputField(input, "IncludeLinkedAccounts", _cloudwatchlogsIncludeLinkedAccounts); err != nil {
			log.Errorf("invalid --include-linked-accounts: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsLimit) > 0 {
		if err := assignInputField(input, "Limit", _cloudwatchlogsLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsLogGroupClass) > 0 {
		if err := assignInputField(input, "LogGroupClass", _cloudwatchlogsLogGroupClass); err != nil {
			log.Errorf("invalid --log-group-class: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsLogGroupNamePattern) > 0 {
		input.LogGroupNamePattern = aws.String(_cloudwatchlogsLogGroupNamePattern)
	}
	if len(_cloudwatchlogsNextToken) > 0 {
		input.NextToken = aws.String(_cloudwatchlogsNextToken)
	}

	if resp, err := client.ListLogGroups(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of the log groups that were analyzed during a single CloudWatch
// Logs Insights query. This can be useful for queries that use log group name
// prefixes or the filterIndex command, because the log groups are dynamically
// selected in these cases.
//
// For more information about field indexes, see [Create field indexes to improve query performance and reduce costs].
//
// [Create field indexes to improve query performance and reduce costs]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatchLogs-Field-Indexing.html
func cloudwatchlogs_ListLogGroupsForQuery(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.ListLogGroupsForQueryInput{
		// QueryId: *string, // Required
	}

	if len(_cloudwatchlogsQueryId) > 0 {
		input.QueryId = aws.String(_cloudwatchlogsQueryId)
	}
	if len(_cloudwatchlogsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cloudwatchlogsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsNextToken) > 0 {
		input.NextToken = aws.String(_cloudwatchlogsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListLogGroupsForQuery(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudwatchlogs.ListLogGroupsForQueryOutput
	p := cloudwatchlogs.NewListLogGroupsForQueryPaginator(client, input)
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

// Lists all scheduled queries in your account and region. You can filter results
// by state to show only enabled or disabled queries.
func cloudwatchlogs_ListScheduledQueries(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.ListScheduledQueriesInput{}

	if len(_cloudwatchlogsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cloudwatchlogsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsNextToken) > 0 {
		input.NextToken = aws.String(_cloudwatchlogsNextToken)
	}
	if len(_cloudwatchlogsState) > 0 {
		if err := assignInputField(input, "State", _cloudwatchlogsState); err != nil {
			log.Errorf("invalid --state: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListScheduledQueries(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudwatchlogs.ListScheduledQueriesOutput
	p := cloudwatchlogs.NewListScheduledQueriesPaginator(client, input)
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

// Returns a list of data source associations for a specified S3 Table
// Integration, showing which data sources are currently associated for query
// access.
func cloudwatchlogs_ListSourcesForS3TableIntegration(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.ListSourcesForS3TableIntegrationInput{
		// IntegrationArn: *string, // Required
	}

	if len(_cloudwatchlogsIntegrationArn) > 0 {
		input.IntegrationArn = aws.String(_cloudwatchlogsIntegrationArn)
	}
	if len(_cloudwatchlogsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cloudwatchlogsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsNextToken) > 0 {
		input.NextToken = aws.String(_cloudwatchlogsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSourcesForS3TableIntegration(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudwatchlogs.ListSourcesForS3TableIntegrationOutput
	p := cloudwatchlogs.NewListSourcesForS3TableIntegrationPaginator(client, input)
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

// Displays the tags associated with a CloudWatch Logs resource. Currently, log
// groups and destinations support tagging.
func cloudwatchlogs_ListTagsForResource(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_cloudwatchlogsResourceArn) > 0 {
		input.ResourceArn = aws.String(_cloudwatchlogsResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The ListTagsLogGroup operation is on the path to deprecation. We recommend that
// you use [ListTagsForResource]instead.
//
// Lists the tags for the specified log group.
//
// Deprecated: Please use the generic tagging API ListTagsForResource
//
// [ListTagsForResource]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_ListTagsForResource.html
func cloudwatchlogs_ListTagsLogGroup(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.ListTagsLogGroupInput{
		// LogGroupName: *string, // Required
	}

	if len(_cloudwatchlogsLogGroupName) > 0 {
		input.LogGroupName = aws.String(_cloudwatchlogsLogGroupName)
	}

	if resp, err := client.ListTagsLogGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an account-level data protection policy, subscription filter policy,
// field index policy, transformer policy, or metric extraction policy that applies
// to all log groups, a subset of log groups, or a data source name and type
// combination in the account.
//
// For field index policies, you can configure indexed fields as facets to enable
// interactive exploration of your logs. Facets provide value distributions and
// counts for indexed fields in the CloudWatch Logs Insights console without
// requiring query execution. For more information, see [Use facets to group and explore logs].
//
// To use this operation, you must be signed on with the correct permissions
// depending on the type of policy that you are creating.
//
// - To create a data protection policy, you must have the
// logs:PutDataProtectionPolicy and logs:PutAccountPolicy permissions.
//
// - To create a subscription filter policy, you must have the
// logs:PutSubscriptionFilter and logs:PutAccountPolicy permissions.
//
// - To create a transformer policy, you must have the logs:PutTransformer and
// logs:PutAccountPolicy permissions.
//
// - To create a field index policy, you must have the logs:PutIndexPolicy and
// logs:PutAccountPolicy permissions.
//
// - To configure facets for field index policies, you must have the
// logs:PutIndexPolicy and logs:PutAccountPolicy permissions.
//
// - To create a metric extraction policy, you must have the
// logs:PutMetricExtractionPolicy and logs:PutAccountPolicy permissions.
//
// # Data protection policy
//
// A data protection policy can help safeguard sensitive data that's ingested by
// your log groups by auditing and masking the sensitive log data. Each account can
// have only one account-level data protection policy.
//
// Sensitive data is detected and masked when it is ingested into a log group.
// When you set a data protection policy, log events ingested into the log groups
// before that time are not masked.
//
// If you use PutAccountPolicy to create a data protection policy for your whole
// account, it applies to both existing log groups and all log groups that are
// created later in this account. The account-level policy is applied to existing
// log groups with eventual consistency. It might take up to 5 minutes before
// sensitive data in existing log groups begins to be masked.
//
// By default, when a user views a log event that includes masked data, the
// sensitive data is replaced by asterisks. A user who has the logs:Unmask
// permission can use a [GetLogEvents]or [FilterLogEvents] operation with the unmask parameter set to true to
// view the unmasked log events. Users with the logs:Unmask can also view unmasked
// data in the CloudWatch Logs console by running a CloudWatch Logs Insights query
// with the unmask query command.
//
// For more information, including a list of types of data that can be audited and
// masked, see [Protect sensitive log data with masking].
//
// To use the PutAccountPolicy operation for a data protection policy, you must be
// signed on with the logs:PutDataProtectionPolicy and logs:PutAccountPolicy
// permissions.
//
// The PutAccountPolicy operation applies to all log groups in the account. You
// can use [PutDataProtectionPolicy]to create a data protection policy that applies to just one log group.
// If a log group has its own data protection policy and the account also has an
// account-level data protection policy, then the two policies are cumulative. Any
// sensitive term specified in either policy is masked.
//
// # Subscription filter policy
//
// A subscription filter policy sets up a real-time feed of log events from
// CloudWatch Logs to other Amazon Web Services services. Account-level
// subscription filter policies apply to both existing log groups and log groups
// that are created later in this account. Supported destinations are Kinesis Data
// Streams, Firehose, and Lambda. When log events are sent to the receiving
// service, they are Base64 encoded and compressed with the GZIP format.
//
// The following destinations are supported for subscription filters:
//
// - An Kinesis Data Streams data stream in the same account as the subscription
// policy, for same-account delivery.
//
// - An Firehose data stream in the same account as the subscription policy, for
// same-account delivery.
//
// - A Lambda function in the same account as the subscription policy, for
// same-account delivery.
//
// - A logical destination in a different account created with [PutDestination], for
// cross-account delivery. Kinesis Data Streams and Firehose are supported as
// logical destinations.
//
// Each account can have one account-level subscription filter policy per Region.
// If you are updating an existing filter, you must specify the correct name in
// PolicyName . To perform a PutAccountPolicy subscription filter operation for
// any destination except a Lambda function, you must also have the iam:PassRole
// permission.
//
// # Transformer policy
//
// Creates or updates a log transformer policy for your account. You use log
// transformers to transform log events into a different format, making them easier
// for you to process and analyze. You can also transform logs from different
// sources into standardized formats that contain relevant, source-specific
// information. After you have created a transformer, CloudWatch Logs performs this
// transformation at the time of log ingestion. You can then refer to the
// transformed versions of the logs during operations such as querying with
// CloudWatch Logs Insights or creating metric filters or subscription filters.
//
// You can also use a transformer to copy metadata from metadata keys into the log
// events themselves. This metadata can include log group name, log stream name,
// account ID and Region.
//
// A transformer for a log group is a series of processors, where each processor
// applies one type of transformation to the log events ingested into this log
// group. For more information about the available processors to use in a
// transformer, see [Processors that you can use].
//
// Having log events in standardized format enables visibility across your
// applications for your log analysis, reporting, and alarming needs. CloudWatch
// Logs provides transformation for common log types with out-of-the-box
// transformation templates for major Amazon Web Services log sources such as VPC
// flow logs, Lambda, and Amazon RDS. You can use pre-built transformation
// templates or create custom transformation policies.
//
// You can create transformers only for the log groups in the Standard log class.
//
// You can have one account-level transformer policy that applies to all log
// groups in the account. Or you can create as many as 20 account-level transformer
// policies that are each scoped to a subset of log groups with the
// selectionCriteria parameter. If you have multiple account-level transformer
// policies with selection criteria, no two of them can use the same or overlapping
// log group name prefixes. For example, if you have one policy filtered to log
// groups that start with my-log , you can't have another transformer policy
// filtered to my-logpprod or my-logging .
//
// You can also set up a transformer at the log-group level. For more information,
// see [PutTransformer]. If there is both a log-group level transformer created with PutTransformer
// and an account-level transformer that could apply to the same log group, the log
// group uses only the log-group level transformer. It ignores the account-level
// transformer.
//
// # Field index policy
//
// You can use field index policies to create indexes on fields found in log
// events for a log group or data source name and type combination. Creating field
// indexes can help lower the scan volume for CloudWatch Logs Insights queries that
// reference those fields, because these queries attempt to skip the processing of
// log events that are known to not match the indexed field. Good fields to index
// are fields that you often need to query for and fields or values that match only
// a small fraction of the total log events. Common examples of indexes include
// request ID, session ID, user IDs, or instance IDs. For more information, see [Create field indexes to improve query performance and reduce costs]
//
// To find the fields that are in your log group events, use the [GetLogGroupFields] operation. To
// find the fields for a data source use the [GetLogFields]operation.
//
// For example, suppose you have created a field index for requestId . Then, any
// CloudWatch Logs Insights query on that log group that includes requestId =
// value or requestId in [value, value, ...] will attempt to process only the log
// events where the indexed field matches the specified value.
//
// Matches of log events to the names of indexed fields are case-sensitive. For
// example, an indexed field of RequestId won't match a log event containing
// requestId .
//
// You can have one account-level field index policy that applies to all log
// groups in the account. Or you can create as many as 20 account-level field index
// policies that are each scoped to a subset of log groups using LogGroupNamePrefix
// with the selectionCriteria parameter. You can have another 20 account-level
// field index policies using DataSourceName and DataSourceType for the
// selectionCriteria parameter. If you have multiple account-level index policies
// with LogGroupNamePrefix selection criteria, no two of them can use the same or
// overlapping log group name prefixes. For example, if you have one policy
// filtered to log groups that start with my-log, you can't have another field
// index policy filtered to my-logpprod or my-logging. Similarly, if you have
// multiple account-level index policies with DataSourceName and DataSourceType
// selection criteria, no two of them can use the same data source name and type
// combination. For example, if you have one policy filtered to the data source
// name amazon_vpc and data source type flow you cannot create another policy with
// this combination.
//
// If you create an account-level field index policy in a monitoring account in
// cross-account observability, the policy is applied only to the monitoring
// account and not to any source accounts.
//
// CloudWatch Logs provides default field indexes for all log groups in the
// Standard log class. Default field indexes are automatically available for the
// following fields:
//
// - (at)logStream
//
// - (at)aws.region
//
// - (at)aws.account
//
// - (at)source.log
//
// - (at)data_source_name
//
// - (at)data_source_type
//
// - (at)data_format
//
// - traceId
//
// - severityText
//
// - attributes.session.id
//
// CloudWatch Logs provides default field indexes for certain data source name and
// type combinations as well. Default field indexes are automatically available for
// the following data source name and type combinations as identified in the
// following list:
//
// amazon_vpc.flow
//
// - action
//
// - logStatus
//
// - region
//
// - flowDirection
//
// - type
//
// amazon_route53.resolver_query
//
// - transport
//
// - rcode
//
// aws_waf.access
//
// - action
//
// - httpRequest.country
//
// aws_cloudtrail.data , aws_cloudtrail.management
//
// - eventSource
//
// - eventName
//
// - awsRegion
//
// - userAgent
//
// - errorCode
//
// - eventType
//
// - managementEvent
//
// - readOnly
//
// - eventCategory
//
// - requestId
//
// Default field indexes are in addition to any custom field indexes you define
// within your policy. Default field indexes are not counted towards your [field index quota].
//
// If you want to create a field index policy for a single log group, you can use [PutIndexPolicy]
// instead of PutAccountPolicy . If you do so, that log group will use that
// log-group level policy and any account-level policies that match at the data
// source level; any account-level policy that matches at the log group level (for
// example, no selection criteria or log group name prefix selection criteria) will
// be ignored.
//
// # Metric extraction policy
//
// A metric extraction policy controls whether CloudWatch Metrics can be created
// through the Embedded Metrics Format (EMF) for log groups in your account. By
// default, EMF metric creation is enabled for all log groups. You can use metric
// extraction policies to disable EMF metric creation for your entire account or
// specific log groups.
//
// When a policy disables EMF metric creation for a log group, log events in the
// EMF format are still ingested, but no CloudWatch Metrics are created from them.
//
// Creating a policy disables metrics for Amazon Web Services features that use
// EMF to create metrics, such as CloudWatch Container Insights and CloudWatch
// Application Signals. To prevent turning off those features by accident, we
// recommend that you exclude the underlying log-groups through a
// selection-criteria such as LogGroupNamePrefix NOT IN ["/aws/containerinsights",
// "/aws/ecs/containerinsights", "/aws/application-signals/data"] .
//
// Each account can have either one account-level metric extraction policy that
// applies to all log groups, or up to 5 policies that are each scoped to a subset
// of log groups with the selectionCriteria parameter. The selection criteria
// supports filtering by LogGroupName and LogGroupNamePrefix using the operators IN
// and NOT IN . You can specify up to 50 values in each IN or NOT IN list.
//
// The selection criteria can be specified in these formats:
//
// LogGroupName IN ["log-group-1", "log-group-2"]
//
// LogGroupNamePrefix NOT IN ["/aws/prefix1", "/aws/prefix2"]
//
// If you have multiple account-level metric extraction policies with selection
// criteria, no two of them can have overlapping criteria. For example, if you have
// one policy with selection criteria LogGroupNamePrefix IN ["my-log"] , you can't
// have another metric extraction policy with selection criteria
// LogGroupNamePrefix IN ["/my-log-prod"] or LogGroupNamePrefix IN ["/my-logging"]
// , as the set of log groups matching these prefixes would be a subset of the log
// groups matching the first policy's prefix, creating an overlap.
//
// When using NOT IN , only one policy with this operator is allowed per account.
//
// When combining policies with IN and NOT IN operators, the overlap check ensures
// that policies don't have conflicting effects. Two policies with IN and NOT IN
// operators do not overlap if and only if every value in the IN policy is
// completely contained within some value in the NOT IN policy. For example:
//
// - If you have a NOT IN policy for prefix "/aws/lambda" , you can create an IN
// policy for the exact log group name "/aws/lambda/function1" because the set of
// log groups matching "/aws/lambda/function1" is a subset of the log groups
// matching "/aws/lambda" .
//
// - If you have a NOT IN policy for prefix "/aws/lambda" , you cannot create an
// IN policy for prefix "/aws" because the set of log groups matching "/aws" is
// not a subset of the log groups matching "/aws/lambda" .
//
// [PutDestination]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutDestination.html
// [PutTransformer]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutTransformer.html
// [GetLogFields]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_GetLogFields.html
// [PutDataProtectionPolicy]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutDataProtectionPolicy.html
// [FilterLogEvents]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_FilterLogEvents.html
// [Processors that you can use]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation.html#CloudWatch-Logs-Transformation-Processors
// [GetLogEvents]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_GetLogEvents.html
// [field index quota]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatchLogs-Field-Indexing-Syntax
// [PutIndexPolicy]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutIndexPolicy.html
// [Protect sensitive log data with masking]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/mask-sensitive-log-data.html
// [GetLogGroupFields]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_GetLogGroupFields.html
// [Use facets to group and explore logs]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatchLogs-Facets.html
// [Create field indexes to improve query performance and reduce costs]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatchLogs-Field-Indexing.html
func cloudwatchlogs_PutAccountPolicy(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.PutAccountPolicyInput{
		// PolicyDocument: *string, // Required
		// PolicyName: *string, // Required
		// PolicyType: types.PolicyType, // Required
	}

	if len(_cloudwatchlogsPolicyDocument) > 0 {
		input.PolicyDocument = aws.String(_cloudwatchlogsPolicyDocument)
	}
	if len(_cloudwatchlogsPolicyName) > 0 {
		input.PolicyName = aws.String(_cloudwatchlogsPolicyName)
	}
	if len(_cloudwatchlogsPolicyType) > 0 {
		if err := assignInputField(input, "PolicyType", _cloudwatchlogsPolicyType); err != nil {
			log.Errorf("invalid --policy-type: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsScope) > 0 {
		if err := assignInputField(input, "Scope", _cloudwatchlogsScope); err != nil {
			log.Errorf("invalid --scope: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsSelectionCriteria) > 0 {
		input.SelectionCriteria = aws.String(_cloudwatchlogsSelectionCriteria)
	}

	if resp, err := client.PutAccountPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables or disables bearer token authentication for the specified log group.
// When enabled on a log group, bearer token authentication is enabled on
// operations until it is explicitly disabled.
//
// For information about the parameters that are common to all actions, see [Common Parameters].
//
// [Common Parameters]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/CommonParameters.html
func cloudwatchlogs_PutBearerTokenAuthentication(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.PutBearerTokenAuthenticationInput{
		// BearerTokenAuthenticationEnabled: *bool, // Required
		// LogGroupIdentifier: *string, // Required
	}

	if len(_cloudwatchlogsBearerTokenAuthenticationEnabled) > 0 {
		if err := assignInputField(input, "BearerTokenAuthenticationEnabled", _cloudwatchlogsBearerTokenAuthenticationEnabled); err != nil {
			log.Errorf("invalid --bearer-token-authentication-enabled: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsLogGroupIdentifier) > 0 {
		input.LogGroupIdentifier = aws.String(_cloudwatchlogsLogGroupIdentifier)
	}

	if resp, err := client.PutBearerTokenAuthentication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a data protection policy for the specified log group. A data protection
// policy can help safeguard sensitive data that's ingested by the log group by
// auditing and masking the sensitive log data.
//
// Sensitive data is detected and masked when it is ingested into the log group.
// When you set a data protection policy, log events ingested into the log group
// before that time are not masked.
//
// By default, when a user views a log event that includes masked data, the
// sensitive data is replaced by asterisks. A user who has the logs:Unmask
// permission can use a [GetLogEvents]or [FilterLogEvents] operation with the unmask parameter set to true to
// view the unmasked log events. Users with the logs:Unmask can also view unmasked
// data in the CloudWatch Logs console by running a CloudWatch Logs Insights query
// with the unmask query command.
//
// For more information, including a list of types of data that can be audited and
// masked, see [Protect sensitive log data with masking].
//
// The PutDataProtectionPolicy operation applies to only the specified log group.
// You can also use [PutAccountPolicy]to create an account-level data protection policy that applies
// to all log groups in the account, including both existing log groups and log
// groups that are created level. If a log group has its own data protection policy
// and the account also has an account-level data protection policy, then the two
// policies are cumulative. Any sensitive term specified in either policy is
// masked.
//
// [Protect sensitive log data with masking]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/mask-sensitive-log-data.html
// [FilterLogEvents]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_FilterLogEvents.html
// [PutAccountPolicy]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutAccountPolicy.html
// [GetLogEvents]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_GetLogEvents.html
func cloudwatchlogs_PutDataProtectionPolicy(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.PutDataProtectionPolicyInput{
		// LogGroupIdentifier: *string, // Required
		// PolicyDocument: *string, // Required
	}

	if len(_cloudwatchlogsLogGroupIdentifier) > 0 {
		input.LogGroupIdentifier = aws.String(_cloudwatchlogsLogGroupIdentifier)
	}
	if len(_cloudwatchlogsPolicyDocument) > 0 {
		input.PolicyDocument = aws.String(_cloudwatchlogsPolicyDocument)
	}

	if resp, err := client.PutDataProtectionPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates a logical delivery destination. A delivery destination is an
// Amazon Web Services resource that represents an Amazon Web Services service that
// logs can be sent to. CloudWatch Logs, Amazon S3, and Firehose are supported as
// logs delivery destinations and X-Ray as the trace delivery destination.
//
// To configure logs delivery between a supported Amazon Web Services service and
// a destination, you must do the following:
//
// - Create a delivery source, which is a logical object that represents the
// resource that is actually sending the logs. For more information, see [PutDeliverySource].
//
// - Use PutDeliveryDestination to create a delivery destination in the same
// account of the actual delivery destination. The delivery destination that you
// create is a logical object that represents the actual delivery destination.
//
// - If you are delivering logs cross-account, you must use [PutDeliveryDestinationPolicy]in the destination
// account to assign an IAM policy to the destination. This policy allows delivery
// to that destination.
//
// - Use CreateDelivery to create a delivery by pairing exactly one delivery
// source and one delivery destination. For more information, see [CreateDelivery].
//
// You can configure a single delivery source to send logs to multiple
// destinations by creating multiple deliveries. You can also create multiple
// deliveries to configure multiple delivery sources to send logs to the same
// delivery destination.
//
// Only some Amazon Web Services services support being configured as a delivery
// source. These services are listed as Supported [V2 Permissions] in the table at [Enabling logging from Amazon Web Services services.]
//
// If you use this operation to update an existing delivery destination, all the
// current delivery destination parameters are overwritten with the new parameter
// values that you specify.
//
// [PutDeliverySource]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutDeliverySource.html
// [Enabling logging from Amazon Web Services services.]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-logs-and-resource-policy.html
// [CreateDelivery]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_CreateDelivery.html
// [PutDeliveryDestinationPolicy]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutDeliveryDestinationPolicy.html
func cloudwatchlogs_PutDeliveryDestination(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.PutDeliveryDestinationInput{
		// Name: *string, // Required
	}

	if len(_cloudwatchlogsName) > 0 {
		input.Name = aws.String(_cloudwatchlogsName)
	}
	if len(_cloudwatchlogsDeliveryDestinationConfiguration) > 0 {
		if err := assignInputField(input, "DeliveryDestinationConfiguration", _cloudwatchlogsDeliveryDestinationConfiguration); err != nil {
			log.Errorf("invalid --delivery-destination-configuration: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsDeliveryDestinationType) > 0 {
		if err := assignInputField(input, "DeliveryDestinationType", _cloudwatchlogsDeliveryDestinationType); err != nil {
			log.Errorf("invalid --delivery-destination-type: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsOutputFormat) > 0 {
		if err := assignInputField(input, "OutputFormat", _cloudwatchlogsOutputFormat); err != nil {
			log.Errorf("invalid --output-format: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsTags) > 0 {
		if err := assignInputField(input, "Tags", _cloudwatchlogsTags[0]); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutDeliveryDestination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates and assigns an IAM policy that grants permissions to CloudWatch Logs to
// deliver logs cross-account to a specified destination in this account. To
// configure the delivery of logs from an Amazon Web Services service in another
// account to a logs delivery destination in the current account, you must do the
// following:
//
// - Create a delivery source, which is a logical object that represents the
// resource that is actually sending the logs. For more information, see [PutDeliverySource].
//
// - Create a delivery destination, which is a logical object that represents
// the actual delivery destination. For more information, see [PutDeliveryDestination].
//
// - Use this operation in the destination account to assign an IAM policy to
// the destination. This policy allows delivery to that destination.
//
// - Create a delivery by pairing exactly one delivery source and one delivery
// destination. For more information, see [CreateDelivery].
//
// Only some Amazon Web Services services support being configured as a delivery
// source. These services are listed as Supported [V2 Permissions] in the table at [Enabling logging from Amazon Web Services services.]
//
// The contents of the policy must include two statements. One statement enables
// general logs delivery, and the other allows delivery to the chosen destination.
// See the examples for the needed policies.
//
// [PutDeliveryDestination]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutDeliveryDestination.html
// [PutDeliverySource]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutDeliverySource.html
// [Enabling logging from Amazon Web Services services.]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-logs-and-resource-policy.html
// [CreateDelivery]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_CreateDelivery.html
func cloudwatchlogs_PutDeliveryDestinationPolicy(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.PutDeliveryDestinationPolicyInput{
		// DeliveryDestinationName: *string, // Required
		// DeliveryDestinationPolicy: *string, // Required
	}

	if len(_cloudwatchlogsDeliveryDestinationName) > 0 {
		input.DeliveryDestinationName = aws.String(_cloudwatchlogsDeliveryDestinationName)
	}
	if len(_cloudwatchlogsDeliveryDestinationPolicy) > 0 {
		input.DeliveryDestinationPolicy = aws.String(_cloudwatchlogsDeliveryDestinationPolicy)
	}

	if resp, err := client.PutDeliveryDestinationPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates a logical delivery source. A delivery source represents an
// Amazon Web Services resource that sends logs to an logs delivery destination.
// The destination can be CloudWatch Logs, Amazon S3, Firehose or X-Ray for sending
// traces.
//
// To configure logs delivery between a delivery destination and an Amazon Web
// Services service that is supported as a delivery source, you must do the
// following:
//
// - Use PutDeliverySource to create a delivery source, which is a logical object
// that represents the resource that is actually sending the logs.
//
// - Use PutDeliveryDestination to create a delivery destination, which is a
// logical object that represents the actual delivery destination. For more
// information, see [PutDeliveryDestination].
//
// - If you are delivering logs cross-account, you must use [PutDeliveryDestinationPolicy]in the destination
// account to assign an IAM policy to the destination. This policy allows delivery
// to that destination.
//
// - Use CreateDelivery to create a delivery by pairing exactly one delivery
// source and one delivery destination. For more information, see [CreateDelivery].
//
// You can configure a single delivery source to send logs to multiple
// destinations by creating multiple deliveries. You can also create multiple
// deliveries to configure multiple delivery sources to send logs to the same
// delivery destination.
//
// Only some Amazon Web Services services support being configured as a delivery
// source. These services are listed as Supported [V2 Permissions] in the table at [Enabling logging from Amazon Web Services services.]
//
// If you use this operation to update an existing delivery source, all the
// current delivery source parameters are overwritten with the new parameter values
// that you specify.
//
// [PutDeliveryDestination]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutDeliveryDestination.html
// [Enabling logging from Amazon Web Services services.]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-logs-and-resource-policy.html
// [CreateDelivery]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_CreateDelivery.html
// [PutDeliveryDestinationPolicy]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutDeliveryDestinationPolicy.html
func cloudwatchlogs_PutDeliverySource(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.PutDeliverySourceInput{
		// LogType: *string, // Required
		// Name: *string, // Required
		// ResourceArn: *string, // Required
	}

	if len(_cloudwatchlogsLogType) > 0 {
		input.LogType = aws.String(_cloudwatchlogsLogType)
	}
	if len(_cloudwatchlogsName) > 0 {
		input.Name = aws.String(_cloudwatchlogsName)
	}
	if len(_cloudwatchlogsResourceArn) > 0 {
		input.ResourceArn = aws.String(_cloudwatchlogsResourceArn)
	}
	if len(_cloudwatchlogsTags) > 0 {
		if err := assignInputField(input, "Tags", _cloudwatchlogsTags[0]); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutDeliverySource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates a destination. This operation is used only to create
// destinations for cross-account subscriptions.
//
// A destination encapsulates a physical resource (such as an Amazon Kinesis
// stream). With a destination, you can subscribe to a real-time stream of log
// events for a different account, ingested using [PutLogEvents].
//
// Through an access policy, a destination controls what is written to it. By
// default, PutDestination does not set any access policy with the destination,
// which means a cross-account user cannot call [PutSubscriptionFilter]against this destination. To
// enable this, the destination owner must call [PutDestinationPolicy]after PutDestination .
//
// To perform a PutDestination operation, you must also have the iam:PassRole
// permission.
//
// [PutSubscriptionFilter]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutSubscriptionFilter.html
// [PutLogEvents]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutLogEvents.html
// [PutDestinationPolicy]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutDestinationPolicy.html
func cloudwatchlogs_PutDestination(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.PutDestinationInput{
		// DestinationName: *string, // Required
		// RoleArn: *string, // Required
		// TargetArn: *string, // Required
	}

	if len(_cloudwatchlogsDestinationName) > 0 {
		input.DestinationName = aws.String(_cloudwatchlogsDestinationName)
	}
	if len(_cloudwatchlogsRoleArn) > 0 {
		input.RoleArn = aws.String(_cloudwatchlogsRoleArn)
	}
	if len(_cloudwatchlogsTargetArn) > 0 {
		input.TargetArn = aws.String(_cloudwatchlogsTargetArn)
	}
	if len(_cloudwatchlogsTags) > 0 {
		if err := assignInputField(input, "Tags", _cloudwatchlogsTags[0]); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutDestination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates an access policy associated with an existing destination. An
// access policy is an [IAM policy document]that is used to authorize claims to register a subscription
// filter against a given destination.
//
// [IAM policy document]: https://docs.aws.amazon.com/IAM/latest/UserGuide/policies_overview.html
func cloudwatchlogs_PutDestinationPolicy(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.PutDestinationPolicyInput{
		// AccessPolicy: *string, // Required
		// DestinationName: *string, // Required
	}

	if len(_cloudwatchlogsAccessPolicy) > 0 {
		input.AccessPolicy = aws.String(_cloudwatchlogsAccessPolicy)
	}
	if len(_cloudwatchlogsDestinationName) > 0 {
		input.DestinationName = aws.String(_cloudwatchlogsDestinationName)
	}
	if len(_cloudwatchlogsForceUpdate) > 0 {
		if err := assignInputField(input, "ForceUpdate", _cloudwatchlogsForceUpdate); err != nil {
			log.Errorf("invalid --force-update: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutDestinationPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates a field index policy for the specified log group. Only log
// groups in the Standard log class support field index policies. For more
// information about log classes, see [Log classes].
//
// You can use field index policies to create field indexes on fields found in log
// events in the log group. Creating field indexes speeds up and lowers the costs
// for CloudWatch Logs Insights queries that reference those field indexes, because
// these queries attempt to skip the processing of log events that are known to not
// match the indexed field. Good fields to index are fields that you often need to
// query for and fields or values that match only a small fraction of the total log
// events. Common examples of indexes include request ID, session ID, userID, and
// instance IDs. For more information, see [Create field indexes to improve query performance and reduce costs].
//
// You can configure indexed fields as facets to enable interactive exploration
// and filtering of your logs in the CloudWatch Logs Insights console. Facets allow
// you to view value distributions and counts for indexed fields without running
// queries. When you create a field index, you can optionally set it as a facet to
// enable this interactive analysis capability. For more information, see [Use facets to group and explore logs].
//
// To find the fields that are in your log group events, use the [GetLogGroupFields] operation.
//
// For example, suppose you have created a field index for requestId . Then, any
// CloudWatch Logs Insights query on that log group that includes requestId =
// value or requestId IN [value, value, ...] will process fewer log events to
// reduce costs, and have improved performance.
//
// CloudWatch Logs provides default field indexes for all log groups in the
// Standard log class. Default field indexes are automatically available for the
// following fields:
//
// - (at)logStream
//
// - (at)aws.region
//
// - (at)aws.account
//
// - (at)source.log
//
// - traceId
//
// Default field indexes are in addition to any custom field indexes you define
// within your policy. Default field indexes are not counted towards your field
// index quota.
//
// Each index policy has the following quotas and restrictions:
//
// - As many as 20 fields can be included in the policy.
//
// - Each field name can include as many as 100 characters.
//
// Matches of log events to the names of indexed fields are case-sensitive. For
// example, a field index of RequestId won't match a log event containing requestId
// .
//
// Log group-level field index policies created with PutIndexPolicy override
// account-level field index policies created with [PutAccountPolicy]that apply to log groups. If
// you use PutIndexPolicy to create a field index policy for a log group, that log
// group uses only that policy for log group-level indexing, including any facet
// configurations. The log group ignores any account-wide field index policy that
// applies to log groups, but data source-based account policies may still apply.
//
// [Log classes]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch_Logs_Log_Classes.html
// [GetLogGroupFields]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_GetLogGroupFields.html
// [PutAccountPolicy]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutAccountPolicy.html
// [Create field indexes to improve query performance and reduce costs]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatchLogs-Field-Indexing.html
// [Use facets to group and explore logs]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatchLogs-Facets.html
func cloudwatchlogs_PutIndexPolicy(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.PutIndexPolicyInput{
		// LogGroupIdentifier: *string, // Required
		// PolicyDocument: *string, // Required
	}

	if len(_cloudwatchlogsLogGroupIdentifier) > 0 {
		input.LogGroupIdentifier = aws.String(_cloudwatchlogsLogGroupIdentifier)
	}
	if len(_cloudwatchlogsPolicyDocument) > 0 {
		input.PolicyDocument = aws.String(_cloudwatchlogsPolicyDocument)
	}

	if resp, err := client.PutIndexPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an integration between CloudWatch Logs and another service in this
// account. Currently, only integrations with OpenSearch Service are supported, and
// currently you can have only one integration in your account.
//
// Integrating with OpenSearch Service makes it possible for you to create curated
// vended logs dashboards, powered by OpenSearch Service analytics. For more
// information, see [Vended log dashboards powered by Amazon OpenSearch Service].
//
// You can use this operation only to create a new integration. You can't modify
// an existing integration.
//
// [Vended log dashboards powered by Amazon OpenSearch Service]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatchLogs-OpenSearch-Dashboards.html
func cloudwatchlogs_PutIntegration(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.PutIntegrationInput{
		// IntegrationName: *string, // Required
		// IntegrationType: types.IntegrationType, // Required
		// ResourceConfig: types.ResourceConfig, // Required
	}

	if len(_cloudwatchlogsIntegrationName) > 0 {
		input.IntegrationName = aws.String(_cloudwatchlogsIntegrationName)
	}
	if len(_cloudwatchlogsIntegrationType) > 0 {
		if err := assignInputField(input, "IntegrationType", _cloudwatchlogsIntegrationType); err != nil {
			log.Errorf("invalid --integration-type: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsResourceConfig) > 0 {
		if err := assignInputField(input, "ResourceConfig", _cloudwatchlogsResourceConfig); err != nil {
			log.Errorf("invalid --resource-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutIntegration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Uploads a batch of log events to the specified log stream.
// The sequence token is now ignored in PutLogEvents actions. PutLogEvents actions
// are always accepted and never return InvalidSequenceTokenException or
// DataAlreadyAcceptedException even if the sequence token is not valid. You can
// use parallel PutLogEvents actions on the same log stream.
//
// The batch of events must satisfy the following constraints:
//
// - The maximum batch size is 1,048,576 bytes. This size is calculated as the
// sum of all event messages in UTF-8, plus 26 bytes for each log event.
//
// - Events more than 2 hours in the future are rejected while processing
// remaining valid events.
//
// - Events older than 14 days or preceding the log group's retention period are
// rejected while processing remaining valid events.
//
// - The log events in the batch must be in chronological order by their
// timestamp. The timestamp is the time that the event occurred, expressed as the
// number of milliseconds after Jan 1, 1970 00:00:00 UTC . (In Amazon Web
// Services Tools for PowerShell and the Amazon Web Services SDK for .NET, the
// timestamp is specified in .NET format: yyyy-mm-ddThh:mm:ss . For example,
// 2017-09-15T13:45:30 .)
//
// - A batch of log events in a single request must be in a chronological order.
// Otherwise, the operation fails.
//
// - Each log event can be no larger than 1 MB.
//
// - The maximum number of log events in a batch is 10,000.
//
// - For valid events (within 14 days in the past to 2 hours in future), the
// time span in a single batch cannot exceed 24 hours. Otherwise, the operation
// fails.
//
// The quota of five requests per second per log stream has been removed. Instead,
// PutLogEvents actions are throttled based on a per-second per-account quota. You
// can request an increase to the per-second throttling quota by using the Service
// Quotas service.
//
// If a call to PutLogEvents returns "UnrecognizedClientException" the most likely
// cause is a non-valid Amazon Web Services access key ID or secret key.
func cloudwatchlogs_PutLogEvents(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.PutLogEventsInput{
		// LogEvents: []types.InputLogEvent, // Required
		// LogGroupName: *string, // Required
		// LogStreamName: *string, // Required
	}

	if len(_cloudwatchlogsLogEvents) > 0 {
		if err := assignInputField(input, "LogEvents", _cloudwatchlogsLogEvents); err != nil {
			log.Errorf("invalid --log-events: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsLogGroupName) > 0 {
		input.LogGroupName = aws.String(_cloudwatchlogsLogGroupName)
	}
	if len(_cloudwatchlogsLogStreamName) > 0 {
		input.LogStreamName = aws.String(_cloudwatchlogsLogStreamName)
	}
	if len(_cloudwatchlogsEntity) > 0 {
		if err := assignInputField(input, "Entity", _cloudwatchlogsEntity); err != nil {
			log.Errorf("invalid --entity: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsSequenceToken) > 0 {
		input.SequenceToken = aws.String(_cloudwatchlogsSequenceToken)
	}

	if resp, err := client.PutLogEvents(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables or disables deletion protection for the specified log group. When
// enabled on a log group, deletion protection blocks all deletion operations until
// it is explicitly disabled.
//
// For information about the parameters that are common to all actions, see [Common Parameters].
//
// [Common Parameters]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/CommonParameters.html
func cloudwatchlogs_PutLogGroupDeletionProtection(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.PutLogGroupDeletionProtectionInput{
		// DeletionProtectionEnabled: *bool, // Required
		// LogGroupIdentifier: *string, // Required
	}

	if len(_cloudwatchlogsDeletionProtectionEnabled) > 0 {
		if err := assignInputField(input, "DeletionProtectionEnabled", _cloudwatchlogsDeletionProtectionEnabled); err != nil {
			log.Errorf("invalid --deletion-protection-enabled: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsLogGroupIdentifier) > 0 {
		input.LogGroupIdentifier = aws.String(_cloudwatchlogsLogGroupIdentifier)
	}

	if resp, err := client.PutLogGroupDeletionProtection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates a metric filter and associates it with the specified log
// group. With metric filters, you can configure rules to extract metric data from
// log events ingested through [PutLogEvents].
//
// The maximum number of metric filters that can be associated with a log group is
// 100.
//
// Using regular expressions in filter patterns is supported. For these filters,
// there is a quota of two regular expression patterns within a single filter
// pattern. There is also a quota of five regular expression patterns per log
// group. For more information about using regular expressions in filter patterns,
// see [Filter pattern syntax for metric filters, subscription filters, filter log events, and Live Tail].
//
// When you create a metric filter, you can also optionally assign a unit and
// dimensions to the metric that is created.
//
// Metrics extracted from log events are charged as custom metrics. To prevent
// unexpected high charges, do not specify high-cardinality fields such as
// IPAddress or requestID as dimensions. Each different value found for a
// dimension is treated as a separate metric and accrues charges as a separate
// custom metric.
//
// CloudWatch Logs might disable a metric filter if it generates 1,000 different
// name/value pairs for your specified dimensions within one hour.
//
// You can also set up a billing alarm to alert you if your charges are higher
// than expected. For more information, see [Creating a Billing Alarm to Monitor Your Estimated Amazon Web Services Charges].
//
// [Creating a Billing Alarm to Monitor Your Estimated Amazon Web Services Charges]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/monitor_estimated_charges_with_cloudwatch.html
// [PutLogEvents]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutLogEvents.html
// [Filter pattern syntax for metric filters, subscription filters, filter log events, and Live Tail]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/FilterAndPatternSyntax.html
func cloudwatchlogs_PutMetricFilter(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.PutMetricFilterInput{
		// FilterName: *string, // Required
		// FilterPattern: *string, // Required
		// LogGroupName: *string, // Required
		// MetricTransformations: []types.MetricTransformation, // Required
	}

	if len(_cloudwatchlogsFilterName) > 0 {
		input.FilterName = aws.String(_cloudwatchlogsFilterName)
	}
	if len(_cloudwatchlogsFilterPattern) > 0 {
		input.FilterPattern = aws.String(_cloudwatchlogsFilterPattern)
	}
	if len(_cloudwatchlogsLogGroupName) > 0 {
		input.LogGroupName = aws.String(_cloudwatchlogsLogGroupName)
	}
	if len(_cloudwatchlogsMetricTransformations) > 0 {
		if err := assignInputField(input, "MetricTransformations", _cloudwatchlogsMetricTransformations); err != nil {
			log.Errorf("invalid --metric-transformations: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsApplyOnTransformedLogs) > 0 {
		if err := assignInputField(input, "ApplyOnTransformedLogs", _cloudwatchlogsApplyOnTransformedLogs); err != nil {
			log.Errorf("invalid --apply-on-transformed-logs: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsEmitSystemFieldDimensions) > 0 {
		input.EmitSystemFieldDimensions = append([]string(nil), _cloudwatchlogsEmitSystemFieldDimensions...)
	}
	if len(_cloudwatchlogsFieldSelectionCriteria) > 0 {
		input.FieldSelectionCriteria = aws.String(_cloudwatchlogsFieldSelectionCriteria)
	}

	if resp, err := client.PutMetricFilter(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates a query definition for CloudWatch Logs Insights. For more
// information, see [Analyzing Log Data with CloudWatch Logs Insights].
//
// To update a query definition, specify its queryDefinitionId in your request.
// The values of name , queryString , and logGroupNames are changed to the values
// that you specify in your update operation. No current values are retained from
// the current query definition. For example, imagine updating a current query
// definition that includes log groups. If you don't specify the logGroupNames
// parameter in your update operation, the query definition changes to contain no
// log groups.
//
// You must have the logs:PutQueryDefinition permission to be able to perform this
// operation.
//
// [Analyzing Log Data with CloudWatch Logs Insights]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AnalyzingLogData.html
func cloudwatchlogs_PutQueryDefinition(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.PutQueryDefinitionInput{
		// Name: *string, // Required
		// QueryString: *string, // Required
	}

	if len(_cloudwatchlogsName) > 0 {
		input.Name = aws.String(_cloudwatchlogsName)
	}
	if len(_cloudwatchlogsQueryString) > 0 {
		input.QueryString = aws.String(_cloudwatchlogsQueryString)
	}
	if len(_cloudwatchlogsClientToken) > 0 {
		input.ClientToken = aws.String(_cloudwatchlogsClientToken)
	}
	if len(_cloudwatchlogsLogGroupNames) > 0 {
		input.LogGroupNames = append([]string(nil), _cloudwatchlogsLogGroupNames...)
	}
	if len(_cloudwatchlogsQueryDefinitionId) > 0 {
		input.QueryDefinitionId = aws.String(_cloudwatchlogsQueryDefinitionId)
	}
	if len(_cloudwatchlogsQueryLanguage) > 0 {
		if err := assignInputField(input, "QueryLanguage", _cloudwatchlogsQueryLanguage); err != nil {
			log.Errorf("invalid --query-language: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutQueryDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates a resource policy allowing other Amazon Web Services
// services to put log events to this account, such as Amazon Route 53. This API
// has the following restrictions:
//
// - Supported actions - Policy only supports logs:PutLogEvents and
// logs:CreateLogStream actions
//
// - Supported principals - Policy only applies when operations are invoked by
// Amazon Web Services service principals (not IAM users, roles, or cross-account
// principals
//
// - Policy limits - An account can have a maximum of 10 policies without
// resourceARN and one per LogGroup resourceARN
//
// Resource policies with actions invoked by non-Amazon Web Services service
// principals (such as IAM users, roles, or other Amazon Web Services accounts)
// will not be enforced. For access control involving these principals, use the IAM
// policies.
func cloudwatchlogs_PutResourcePolicy(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.PutResourcePolicyInput{}

	if len(_cloudwatchlogsExpectedRevisionId) > 0 {
		input.ExpectedRevisionId = aws.String(_cloudwatchlogsExpectedRevisionId)
	}
	if len(_cloudwatchlogsPolicyDocument) > 0 {
		input.PolicyDocument = aws.String(_cloudwatchlogsPolicyDocument)
	}
	if len(_cloudwatchlogsPolicyName) > 0 {
		input.PolicyName = aws.String(_cloudwatchlogsPolicyName)
	}
	if len(_cloudwatchlogsResourceArn) > 0 {
		input.ResourceArn = aws.String(_cloudwatchlogsResourceArn)
	}

	if resp, err := client.PutResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets the retention of the specified log group. With a retention policy, you can
// configure the number of days for which to retain log events in the specified log
// group.
//
// CloudWatch Logs doesn't immediately delete log events when they reach their
// retention setting. It typically takes up to 72 hours after that before log
// events are deleted, but in rare situations might take longer.
//
// To illustrate, imagine that you change a log group to have a longer retention
// setting when it contains log events that are past the expiration date, but
// haven't been deleted. Those log events will take up to 72 hours to be deleted
// after the new retention date is reached. To make sure that log data is deleted
// permanently, keep a log group at its lower retention setting until 72 hours
// after the previous retention period ends. Alternatively, wait to change the
// retention setting until you confirm that the earlier log events are deleted.
//
// When log events reach their retention setting they are marked for deletion.
// After they are marked for deletion, they do not add to your archival storage
// costs anymore, even if they are not actually deleted until later. These log
// events marked for deletion are also not included when you use an API to retrieve
// the storedBytes value to see how many bytes a log group is storing.
func cloudwatchlogs_PutRetentionPolicy(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.PutRetentionPolicyInput{
		// LogGroupName: *string, // Required
		// RetentionInDays: *int32, // Required
	}

	if len(_cloudwatchlogsLogGroupName) > 0 {
		input.LogGroupName = aws.String(_cloudwatchlogsLogGroupName)
	}
	if len(_cloudwatchlogsRetentionInDays) > 0 {
		if err := assignInputField(input, "RetentionInDays", _cloudwatchlogsRetentionInDays); err != nil {
			log.Errorf("invalid --retention-in-days: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutRetentionPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates a subscription filter and associates it with the specified
// log group. With subscription filters, you can subscribe to a real-time stream of
// log events ingested through [PutLogEvents]and have them delivered to a specific destination.
// When log events are sent to the receiving service, they are Base64 encoded and
// compressed with the GZIP format.
//
// The following destinations are supported for subscription filters:
//
// - An Amazon Kinesis data stream belonging to the same account as the
// subscription filter, for same-account delivery.
//
// - A logical destination created with [PutDestination]that belongs to a different account, for
// cross-account delivery. We currently support Kinesis Data Streams and Firehose
// as logical destinations.
//
// - An Amazon Kinesis Data Firehose delivery stream that belongs to the same
// account as the subscription filter, for same-account delivery.
//
// - An Lambda function that belongs to the same account as the subscription
// filter, for same-account delivery.
//
// Each log group can have up to two subscription filters associated with it. If
// you are updating an existing filter, you must specify the correct name in
// filterName .
//
// Using regular expressions in filter patterns is supported. For these filters,
// there is a quotas of quota of two regular expression patterns within a single
// filter pattern. There is also a quota of five regular expression patterns per
// log group. For more information about using regular expressions in filter
// patterns, see [Filter pattern syntax for metric filters, subscription filters, filter log events, and Live Tail].
//
// To perform a PutSubscriptionFilter operation for any destination except a
// Lambda function, you must also have the iam:PassRole permission.
//
// [PutDestination]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutDestination.html
// [PutLogEvents]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutLogEvents.html
// [Filter pattern syntax for metric filters, subscription filters, filter log events, and Live Tail]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/FilterAndPatternSyntax.html
func cloudwatchlogs_PutSubscriptionFilter(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.PutSubscriptionFilterInput{
		// DestinationArn: *string, // Required
		// FilterName: *string, // Required
		// FilterPattern: *string, // Required
		// LogGroupName: *string, // Required
	}

	if len(_cloudwatchlogsDestinationArn) > 0 {
		input.DestinationArn = aws.String(_cloudwatchlogsDestinationArn)
	}
	if len(_cloudwatchlogsFilterName) > 0 {
		input.FilterName = aws.String(_cloudwatchlogsFilterName)
	}
	if len(_cloudwatchlogsFilterPattern) > 0 {
		input.FilterPattern = aws.String(_cloudwatchlogsFilterPattern)
	}
	if len(_cloudwatchlogsLogGroupName) > 0 {
		input.LogGroupName = aws.String(_cloudwatchlogsLogGroupName)
	}
	if len(_cloudwatchlogsApplyOnTransformedLogs) > 0 {
		if err := assignInputField(input, "ApplyOnTransformedLogs", _cloudwatchlogsApplyOnTransformedLogs); err != nil {
			log.Errorf("invalid --apply-on-transformed-logs: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsDistribution) > 0 {
		if err := assignInputField(input, "Distribution", _cloudwatchlogsDistribution); err != nil {
			log.Errorf("invalid --distribution: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsEmitSystemFields) > 0 {
		input.EmitSystemFields = append([]string(nil), _cloudwatchlogsEmitSystemFields...)
	}
	if len(_cloudwatchlogsFieldSelectionCriteria) > 0 {
		input.FieldSelectionCriteria = aws.String(_cloudwatchlogsFieldSelectionCriteria)
	}
	if len(_cloudwatchlogsRoleArn) > 0 {
		input.RoleArn = aws.String(_cloudwatchlogsRoleArn)
	}

	if resp, err := client.PutSubscriptionFilter(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates a log transformer for a single log group. You use log
// transformers to transform log events into a different format, making them easier
// for you to process and analyze. You can also transform logs from different
// sources into standardized formats that contains relevant, source-specific
// information.
//
// After you have created a transformer, CloudWatch Logs performs the
// transformations at the time of log ingestion. You can then refer to the
// transformed versions of the logs during operations such as querying with
// CloudWatch Logs Insights or creating metric filters or subscription filers.
//
// You can also use a transformer to copy metadata from metadata keys into the log
// events themselves. This metadata can include log group name, log stream name,
// account ID and Region.
//
// A transformer for a log group is a series of processors, where each processor
// applies one type of transformation to the log events ingested into this log
// group. The processors work one after another, in the order that you list them,
// like a pipeline. For more information about the available processors to use in a
// transformer, see [Processors that you can use].
//
// Having log events in standardized format enables visibility across your
// applications for your log analysis, reporting, and alarming needs. CloudWatch
// Logs provides transformation for common log types with out-of-the-box
// transformation templates for major Amazon Web Services log sources such as VPC
// flow logs, Lambda, and Amazon RDS. You can use pre-built transformation
// templates or create custom transformation policies.
//
// You can create transformers only for the log groups in the Standard log class.
//
// You can also set up a transformer at the account level. For more information,
// see [PutAccountPolicy]. If there is both a log-group level transformer created with PutTransformer
// and an account-level transformer that could apply to the same log group, the log
// group uses only the log-group level transformer. It ignores the account-level
// transformer.
//
// [Processors that you can use]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation.html#CloudWatch-Logs-Transformation-Processors
// [PutAccountPolicy]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutAccountPolicy.html
func cloudwatchlogs_PutTransformer(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.PutTransformerInput{
		// LogGroupIdentifier: *string, // Required
		// TransformerConfig: []types.Processor, // Required
	}

	if len(_cloudwatchlogsLogGroupIdentifier) > 0 {
		input.LogGroupIdentifier = aws.String(_cloudwatchlogsLogGroupIdentifier)
	}
	if len(_cloudwatchlogsTransformerConfig) > 0 {
		if err := assignInputField(input, "TransformerConfig", _cloudwatchlogsTransformerConfig); err != nil {
			log.Errorf("invalid --transformer-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutTransformer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a Live Tail streaming session for one or more log groups. A Live Tail
// session returns a stream of log events that have been recently ingested in the
// log groups. For more information, see [Use Live Tail to view logs in near real time].
//
// The response to this operation is a response stream, over which the server
// sends live log events and the client receives them.
//
// The following objects are sent over the stream:
//
// - A single [LiveTailSessionStart]object is sent at the start of the session.
//
// - Every second, a [LiveTailSessionUpdate]object is sent. Each of these objects contains an array of
// the actual log events.
//
// # If no new log events were ingested in the past second, the LiveTailSessionUpdate
//
// object will contain an empty array.
//
// # The array of log events contained in a LiveTailSessionUpdate can include as many
//
// as 500 log events. If the number of log events matching the request exceeds 500
// per second, the log events are sampled down to 500 log events to be included in
// each LiveTailSessionUpdate object.
//
// If your client consumes the log events slower than the server produces them,
//
// CloudWatch Logs buffers up to 10 LiveTailSessionUpdate events or 5000 log
// events, after which it starts dropping the oldest events.
//
// - A [SessionStreamingException]object is returned if an unknown error occurs on the server side.
//
// - A [SessionTimeoutException]object is returned when the session times out, after it has been kept
// open for three hours.
//
// The StartLiveTail API routes requests using SDK host prefix injection. SDK
// versions released before April 1, 2026 route to
// streaming-logs.Region.amazonaws.com , which does not support VPC endpoints. SDK
// versions released on or after April 1, 2026 route to
// stream-logs.Region.amazonaws.com , which supports VPC endpoints. To set up a VPC
// endpoint for this API, see [Creating a VPC endpoint for CloudWatch Logs].
//
// You can end a session before it times out by closing the session stream or by
// closing the client that is receiving the stream. The session also ends if the
// established connection between the client and the server breaks.
//
// For examples of using an SDK to start a Live Tail session, see [Start a Live Tail session using an Amazon Web Services SDK].
//
// [LiveTailSessionStart]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_LiveTailSessionStart.html
// [LiveTailSessionUpdate]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_LiveTailSessionUpdate.html
// [Use Live Tail to view logs in near real time]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatchLogs_LiveTail.html
// [Creating a VPC endpoint for CloudWatch Logs]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/cloudwatch-logs-and-interface-VPC.html#create-VPC-endpoint-for-CloudWatchLogs
// [Start a Live Tail session using an Amazon Web Services SDK]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/example_cloudwatch-logs_StartLiveTail_section.html
// [SessionTimeoutException]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_StartLiveTailResponseStream.html#CWL-Type-StartLiveTailResponseStream-SessionTimeoutException
// [SessionStreamingException]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_StartLiveTailResponseStream.html#CWL-Type-StartLiveTailResponseStream-SessionStreamingException
func cloudwatchlogs_StartLiveTail(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.StartLiveTailInput{
		// LogGroupIdentifiers: []string, // Required
	}

	if len(_cloudwatchlogsLogGroupIdentifiers) > 0 {
		input.LogGroupIdentifiers = append([]string(nil), _cloudwatchlogsLogGroupIdentifiers...)
	}
	if len(_cloudwatchlogsLogEventFilterPattern) > 0 {
		input.LogEventFilterPattern = aws.String(_cloudwatchlogsLogEventFilterPattern)
	}
	if len(_cloudwatchlogsLogStreamNamePrefixes) > 0 {
		input.LogStreamNamePrefixes = append([]string(nil), _cloudwatchlogsLogStreamNamePrefixes...)
	}
	if len(_cloudwatchlogsLogStreamNames) > 0 {
		input.LogStreamNames = append([]string(nil), _cloudwatchlogsLogStreamNames...)
	}

	if resp, err := client.StartLiveTail(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a query of one or more log groups or data sources using CloudWatch Logs
// Insights. You specify the log groups or data sources and time range to query and
// the query string to use. You can query up to 10 data sources in a single query.
//
// For more information, see [CloudWatch Logs Insights Query Syntax].
//
// After you run a query using StartQuery , the query results are stored by
// CloudWatch Logs. You can use [GetQueryResults]to retrieve the results of a query, using the
// queryId that StartQuery returns.
//
// Interactive queries started with StartQuery share concurrency limits with
// automated scheduled query executions. Both types of queries count toward the
// same regional concurrent query quota, so high scheduled query activity may
// affect the availability of concurrent slots for interactive queries.
//
// To specify the log groups to query, a StartQuery operation must include one of
// the following:
//
// - Either exactly one of the following parameters: logGroupName , logGroupNames
// , or logGroupIdentifiers
//
// - Or the queryString must include a SOURCE command to select log groups for
// the query. The SOURCE command can select log groups based on log group name
// prefix, account ID, and log class, or select data sources using dataSource
// syntax in LogsQL, PPL, and SQL.
//
// For more information about the SOURCE command, see [SOURCE].
//
// If you have associated a KMS key with the query results in this account, then [StartQuery]
// uses that key to encrypt the results when it stores them. If no key is
// associated with query results, the query results are encrypted with the default
// CloudWatch Logs encryption method.
//
// Queries time out after 60 minutes of runtime. If your queries are timing out,
// reduce the time range being searched or partition your query into a number of
// queries.
//
// If you are using CloudWatch cross-account observability, you can use this
// operation in a monitoring account to start a query in a linked source account.
// For more information, see [CloudWatch cross-account observability]. For a cross-account StartQuery operation, the query
// definition must be defined in the monitoring account.
//
// You can have up to 30 concurrent CloudWatch Logs insights queries, including
// queries that have been added to dashboards.
//
// [CloudWatch Logs Insights Query Syntax]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_QuerySyntax.html
// [CloudWatch cross-account observability]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Unified-Cross-Account.html
// [SOURCE]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_QuerySyntax-Source.html
// [GetQueryResults]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_GetQueryResults.html
// [StartQuery]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_StartQuery.html
func cloudwatchlogs_StartQuery(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.StartQueryInput{
		// EndTime: *int64, // Required
		// QueryString: *string, // Required
		// StartTime: *int64, // Required
	}

	if len(_cloudwatchlogsEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _cloudwatchlogsEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsQueryString) > 0 {
		input.QueryString = aws.String(_cloudwatchlogsQueryString)
	}
	if len(_cloudwatchlogsStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _cloudwatchlogsStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsLimit) > 0 {
		if err := assignInputField(input, "Limit", _cloudwatchlogsLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsLogGroupIdentifiers) > 0 {
		input.LogGroupIdentifiers = append([]string(nil), _cloudwatchlogsLogGroupIdentifiers...)
	}
	if len(_cloudwatchlogsLogGroupName) > 0 {
		input.LogGroupName = aws.String(_cloudwatchlogsLogGroupName)
	}
	if len(_cloudwatchlogsLogGroupNames) > 0 {
		input.LogGroupNames = append([]string(nil), _cloudwatchlogsLogGroupNames...)
	}
	if len(_cloudwatchlogsQueryLanguage) > 0 {
		if err := assignInputField(input, "QueryLanguage", _cloudwatchlogsQueryLanguage); err != nil {
			log.Errorf("invalid --query-language: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartQuery(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a CloudWatch Logs Insights query that is in progress. If the query has
// already ended, the operation returns an error indicating that the specified
// query is not running.
//
// This operation can be used to cancel both interactive queries and individual
// scheduled query executions. When used with scheduled queries, StopQuery cancels
// only the specific execution identified by the query ID, not the scheduled query
// configuration itself.
func cloudwatchlogs_StopQuery(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.StopQueryInput{
		// QueryId: *string, // Required
	}

	if len(_cloudwatchlogsQueryId) > 0 {
		input.QueryId = aws.String(_cloudwatchlogsQueryId)
	}

	if resp, err := client.StopQuery(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The TagLogGroup operation is on the path to deprecation. We recommend that you
// use [TagResource]instead.
//
// Adds or updates the specified tags for the specified log group.
//
// To list the tags for a log group, use [ListTagsForResource]. To remove tags, use [UntagResource].
//
// For more information about tags, see [Tag Log Groups in Amazon CloudWatch Logs] in the Amazon CloudWatch Logs User Guide.
//
// CloudWatch Logs doesn't support IAM policies that prevent users from assigning
// specified tags to log groups using the aws:Resource/key-name  or aws:TagKeys
// condition keys. For more information about using tags to control access, see [Controlling access to Amazon Web Services resources using tags].
//
// Deprecated: Please use the generic tagging API TagResource
//
// [TagResource]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_TagResource.html
// [UntagResource]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_UntagResource.html
// [Tag Log Groups in Amazon CloudWatch Logs]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/Working-with-log-groups-and-streams.html#log-group-tagging
// [Controlling access to Amazon Web Services resources using tags]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_tags.html
// [ListTagsForResource]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_ListTagsForResource.html
func cloudwatchlogs_TagLogGroup(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.TagLogGroupInput{
		// LogGroupName: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_cloudwatchlogsLogGroupName) > 0 {
		input.LogGroupName = aws.String(_cloudwatchlogsLogGroupName)
	}
	if len(_cloudwatchlogsTags) > 0 {
		if err := assignInputField(input, "Tags", _cloudwatchlogsTags[0]); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.TagLogGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Assigns one or more tags (key-value pairs) to the specified CloudWatch Logs
// resource. Currently, the only CloudWatch Logs resources that can be tagged are
// log groups and destinations.
//
// Tags can help you organize and categorize your resources. You can also use them
// to scope user permissions by granting a user permission to access or change only
// resources with certain tag values.
//
// Tags don't have any semantic meaning to Amazon Web Services and are interpreted
// strictly as strings of characters.
//
// You can use the TagResource action with a resource that already has tags. If
// you specify a new tag key for the alarm, this tag is appended to the list of
// tags associated with the alarm. If you specify a tag key that is already
// associated with the alarm, the new tag value that you specify replaces the
// previous value for that tag.
//
// You can associate as many as 50 tags with a CloudWatch Logs resource.
func cloudwatchlogs_TagResource(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_cloudwatchlogsResourceArn) > 0 {
		input.ResourceArn = aws.String(_cloudwatchlogsResourceArn)
	}
	if len(_cloudwatchlogsTags) > 0 {
		if err := assignInputField(input, "Tags", _cloudwatchlogsTags[0]); err != nil {
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

// Tests the filter pattern of a metric filter against a sample of log event
// messages. You can use this operation to validate the correctness of a metric
// filter pattern.
func cloudwatchlogs_TestMetricFilter(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.TestMetricFilterInput{
		// FilterPattern: *string, // Required
		// LogEventMessages: []string, // Required
	}

	if len(_cloudwatchlogsFilterPattern) > 0 {
		input.FilterPattern = aws.String(_cloudwatchlogsFilterPattern)
	}
	if len(_cloudwatchlogsLogEventMessages) > 0 {
		input.LogEventMessages = append([]string(nil), _cloudwatchlogsLogEventMessages...)
	}

	if resp, err := client.TestMetricFilter(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Use this operation to test a log transformer. You enter the transformer
// configuration and a set of log events to test with. The operation responds with
// an array that includes the original log events and the transformed versions.
func cloudwatchlogs_TestTransformer(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.TestTransformerInput{
		// LogEventMessages: []string, // Required
		// TransformerConfig: []types.Processor, // Required
	}

	if len(_cloudwatchlogsLogEventMessages) > 0 {
		input.LogEventMessages = append([]string(nil), _cloudwatchlogsLogEventMessages...)
	}
	if len(_cloudwatchlogsTransformerConfig) > 0 {
		if err := assignInputField(input, "TransformerConfig", _cloudwatchlogsTransformerConfig); err != nil {
			log.Errorf("invalid --transformer-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.TestTransformer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The UntagLogGroup operation is on the path to deprecation. We recommend that
// you use [UntagResource]instead.
//
// Removes the specified tags from the specified log group.
//
// To list the tags for a log group, use [ListTagsForResource]. To add tags, use [TagResource].
//
// When using IAM policies to control tag management for CloudWatch Logs log
// groups, the condition keys aws:Resource/key-name and aws:TagKeys cannot be used
// to restrict which tags users can assign.
//
// Deprecated: Please use the generic tagging API UntagResource
//
// [TagResource]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_TagResource.html
// [UntagResource]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_UntagResource.html
// [ListTagsForResource]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_ListTagsForResource.html
func cloudwatchlogs_UntagLogGroup(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.UntagLogGroupInput{
		// LogGroupName: *string, // Required
		// Tags: []string, // Required
	}

	if len(_cloudwatchlogsLogGroupName) > 0 {
		input.LogGroupName = aws.String(_cloudwatchlogsLogGroupName)
	}
	if len(_cloudwatchlogsTags) > 0 {
		input.Tags = append([]string(nil), _cloudwatchlogsTags...)
	}

	if resp, err := client.UntagLogGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes one or more tags from the specified resource.
func cloudwatchlogs_UntagResource(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_cloudwatchlogsResourceArn) > 0 {
		input.ResourceArn = aws.String(_cloudwatchlogsResourceArn)
	}
	if len(_cloudwatchlogsTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _cloudwatchlogsTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Use this operation to suppress anomaly detection for a specified anomaly or
// pattern. If you suppress an anomaly, CloudWatch Logs won't report new
// occurrences of that anomaly and won't update that anomaly with new data. If you
// suppress a pattern, CloudWatch Logs won't report any anomalies related to that
// pattern.
//
// You must specify either anomalyId or patternId , but you can't specify both
// parameters in the same operation.
//
// If you have previously used this operation to suppress detection of a pattern
// or anomaly, you can use it again to cause CloudWatch Logs to end the
// suppression. To do this, use this operation and specify the anomaly or pattern
// to stop suppressing, and omit the suppressionType and suppressionPeriod
// parameters.
func cloudwatchlogs_UpdateAnomaly(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.UpdateAnomalyInput{
		// AnomalyDetectorArn: *string, // Required
	}

	if len(_cloudwatchlogsAnomalyDetectorArn) > 0 {
		input.AnomalyDetectorArn = aws.String(_cloudwatchlogsAnomalyDetectorArn)
	}
	if len(_cloudwatchlogsAnomalyId) > 0 {
		input.AnomalyId = aws.String(_cloudwatchlogsAnomalyId)
	}
	if len(_cloudwatchlogsBaseline) > 0 {
		if err := assignInputField(input, "Baseline", _cloudwatchlogsBaseline); err != nil {
			log.Errorf("invalid --baseline: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsPatternId) > 0 {
		input.PatternId = aws.String(_cloudwatchlogsPatternId)
	}
	if len(_cloudwatchlogsSuppressionPeriod) > 0 {
		if err := assignInputField(input, "SuppressionPeriod", _cloudwatchlogsSuppressionPeriod); err != nil {
			log.Errorf("invalid --suppression-period: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsSuppressionType) > 0 {
		if err := assignInputField(input, "SuppressionType", _cloudwatchlogsSuppressionType); err != nil {
			log.Errorf("invalid --suppression-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateAnomaly(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Use this operation to update the configuration of a [delivery] to change either the S3
// path pattern or the format of the delivered logs. You can't use this operation
// to change the source or destination of the delivery.
//
// [delivery]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_Delivery.html
func cloudwatchlogs_UpdateDeliveryConfiguration(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.UpdateDeliveryConfigurationInput{
		// Id: *string, // Required
	}

	if len(_cloudwatchlogsId) > 0 {
		input.Id = aws.String(_cloudwatchlogsId)
	}
	if len(_cloudwatchlogsFieldDelimiter) > 0 {
		input.FieldDelimiter = aws.String(_cloudwatchlogsFieldDelimiter)
	}
	if len(_cloudwatchlogsRecordFields) > 0 {
		input.RecordFields = append([]string(nil), _cloudwatchlogsRecordFields...)
	}
	if len(_cloudwatchlogsS3DeliveryConfiguration) > 0 {
		if err := assignInputField(input, "S3DeliveryConfiguration", _cloudwatchlogsS3DeliveryConfiguration); err != nil {
			log.Errorf("invalid --s3-delivery-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateDeliveryConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing log anomaly detector.
func cloudwatchlogs_UpdateLogAnomalyDetector(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.UpdateLogAnomalyDetectorInput{
		// AnomalyDetectorArn: *string, // Required
		// Enabled: *bool, // Required
	}

	if len(_cloudwatchlogsAnomalyDetectorArn) > 0 {
		input.AnomalyDetectorArn = aws.String(_cloudwatchlogsAnomalyDetectorArn)
	}
	if len(_cloudwatchlogsEnabled) > 0 {
		if err := assignInputField(input, "Enabled", _cloudwatchlogsEnabled); err != nil {
			log.Errorf("invalid --enabled: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsAnomalyVisibilityTime) > 0 {
		if err := assignInputField(input, "AnomalyVisibilityTime", _cloudwatchlogsAnomalyVisibilityTime); err != nil {
			log.Errorf("invalid --anomaly-visibility-time: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsEvaluationFrequency) > 0 {
		if err := assignInputField(input, "EvaluationFrequency", _cloudwatchlogsEvaluationFrequency); err != nil {
			log.Errorf("invalid --evaluation-frequency: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsFilterPattern) > 0 {
		input.FilterPattern = aws.String(_cloudwatchlogsFilterPattern)
	}

	if resp, err := client.UpdateLogAnomalyDetector(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing scheduled query with new configuration. This operation uses
// PUT semantics, allowing modification of query parameters, schedule, and
// destinations.
func cloudwatchlogs_UpdateScheduledQuery(cfg aws.Config, client *cloudwatchlogs.Client) {
	input := &cloudwatchlogs.UpdateScheduledQueryInput{
		// ExecutionRoleArn: *string, // Required
		// Identifier: *string, // Required
		// QueryLanguage: types.QueryLanguage, // Required
		// QueryString: *string, // Required
		// ScheduleExpression: *string, // Required
	}

	if len(_cloudwatchlogsExecutionRoleArn) > 0 {
		input.ExecutionRoleArn = aws.String(_cloudwatchlogsExecutionRoleArn)
	}
	if len(_cloudwatchlogsIdentifier) > 0 {
		input.Identifier = aws.String(_cloudwatchlogsIdentifier)
	}
	if len(_cloudwatchlogsQueryLanguage) > 0 {
		if err := assignInputField(input, "QueryLanguage", _cloudwatchlogsQueryLanguage); err != nil {
			log.Errorf("invalid --query-language: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsQueryString) > 0 {
		input.QueryString = aws.String(_cloudwatchlogsQueryString)
	}
	if len(_cloudwatchlogsScheduleExpression) > 0 {
		input.ScheduleExpression = aws.String(_cloudwatchlogsScheduleExpression)
	}
	if len(_cloudwatchlogsDescription) > 0 {
		input.Description = aws.String(_cloudwatchlogsDescription)
	}
	if len(_cloudwatchlogsDestinationConfiguration) > 0 {
		if err := assignInputField(input, "DestinationConfiguration", _cloudwatchlogsDestinationConfiguration); err != nil {
			log.Errorf("invalid --destination-configuration: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsLogGroupIdentifiers) > 0 {
		input.LogGroupIdentifiers = append([]string(nil), _cloudwatchlogsLogGroupIdentifiers...)
	}
	if len(_cloudwatchlogsScheduleEndTime) > 0 {
		if err := assignInputField(input, "ScheduleEndTime", _cloudwatchlogsScheduleEndTime); err != nil {
			log.Errorf("invalid --schedule-end-time: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsScheduleStartTime) > 0 {
		if err := assignInputField(input, "ScheduleStartTime", _cloudwatchlogsScheduleStartTime); err != nil {
			log.Errorf("invalid --schedule-start-time: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsStartTimeOffset) > 0 {
		if err := assignInputField(input, "StartTimeOffset", _cloudwatchlogsStartTimeOffset); err != nil {
			log.Errorf("invalid --start-time-offset: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsState) > 0 {
		if err := assignInputField(input, "State", _cloudwatchlogsState); err != nil {
			log.Errorf("invalid --state: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchlogsTimezone) > 0 {
		input.Timezone = aws.String(_cloudwatchlogsTimezone)
	}

	if resp, err := client.UpdateScheduledQuery(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_cloudwatchlogsCmd)
	_cloudwatchlogsCmd.Flags().SortFlags = false

	_cloudwatchlogsCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_cloudwatchlogsCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_cloudwatchlogsCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsAccessPolicy, "access-policy", "", "", "Access Policy")
	_cloudwatchlogsCmd.Flags().StringSliceVarP(&_cloudwatchlogsAccountIdentifiers, "account-identifiers", "", nil, "Account Identifiers")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsAnomalyDetectorArn, "anomaly-detector-arn", "", "", "Anomaly Detector ARN")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsAnomalyId, "anomaly-id", "", "", "Anomaly ID")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsAnomalyVisibilityTime, "anomaly-visibility-time", "", "", "Anomaly Visibility Time")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsApplyOnTransformedLogs, "apply-on-transformed-logs", "", "", "Apply On Transformed Logs")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsBaseline, "baseline", "", "", "Baseline")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsBatchImportStatus, "batch-import-status", "", "", "Batch Import Status")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsBearerTokenAuthenticationEnabled, "bearer-token-authentication-enabled", "", "", "Bearer Token Authentication Enabled")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsClientToken, "client-token", "", "", "Client Token")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsDataSource, "data-source", "", "", "Data Source")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsDataSourceName, "data-source-name", "", "", "Data Source Name")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsDataSourceType, "data-source-type", "", "", "Data Source Type")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsDataSources, "data-sources", "", "", "Data Sources")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsDeletionProtectionEnabled, "deletion-protection-enabled", "", "", "Deletion Protection Enabled")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsDeliveryDestinationArn, "delivery-destination-arn", "", "", "Delivery Destination ARN")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsDeliveryDestinationConfiguration, "delivery-destination-configuration", "", "", "Delivery Destination Configuration")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsDeliveryDestinationName, "delivery-destination-name", "", "", "Delivery Destination Name")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsDeliveryDestinationPolicy, "delivery-destination-policy", "", "", "Delivery Destination Policy")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsDeliveryDestinationType, "delivery-destination-type", "", "", "Delivery Destination Type")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsDeliveryDestinationTypes, "delivery-destination-types", "", "", "Delivery Destination Types")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsDeliverySourceName, "delivery-source-name", "", "", "Delivery Source Name")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsDescending, "descending", "", "", "Descending")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsDescription, "description", "", "", "Description")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsDestination, "destination", "", "", "Destination")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsDestinationArn, "destination-arn", "", "", "Destination ARN")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsDestinationConfiguration, "destination-configuration", "", "", "Destination Configuration")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsDestinationName, "destination-name", "", "", "Destination Name")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsDestinationNamePrefix, "destination-name-prefix", "", "", "Destination Name Prefix")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsDestinationPrefix, "destination-prefix", "", "", "Destination Prefix")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsDetectorName, "detector-name", "", "", "Detector Name")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsDistribution, "distribution", "", "", "Distribution")
	_cloudwatchlogsCmd.Flags().StringSliceVarP(&_cloudwatchlogsEmitSystemFieldDimensions, "emit-system-field-dimensions", "", nil, "Emit System Field Dimensions")
	_cloudwatchlogsCmd.Flags().StringSliceVarP(&_cloudwatchlogsEmitSystemFields, "emit-system-fields", "", nil, "Emit System Fields")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsEnabled, "enabled", "", "", "Enabled")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsEndTime, "end-time", "", "", "End Time")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsEntity, "entity", "", "", "Entity")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsEvaluationFrequency, "evaluation-frequency", "", "", "Evaluation Frequency")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsExecutionRoleArn, "execution-role-arn", "", "", "Execution Role ARN")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsExecutionStatuses, "execution-statuses", "", "", "Execution Statuses")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsExpectedRevisionId, "expected-revision-id", "", "", "Expected Revision ID")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsFieldDelimiter, "field-delimiter", "", "", "Field Delimiter")
	_cloudwatchlogsCmd.Flags().StringSliceVarP(&_cloudwatchlogsFieldIndexNames, "field-index-names", "", nil, "Field Index Names")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsFieldSelectionCriteria, "field-selection-criteria", "", "", "Field Selection Criteria")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsFilterLogGroupArn, "filter-log-group-arn", "", "", "Filter Log Group ARN")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsFilterName, "filter-name", "", "", "Filter Name")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsFilterNamePrefix, "filter-name-prefix", "", "", "Filter Name Prefix")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsFilterPattern, "filter-pattern", "", "", "Filter Pattern")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsForce, "force", "", "", "Force")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsForceUpdate, "force-update", "", "", "Force Update")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsFrom, "from", "", "", "From")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsGroupBy, "group-by", "", "", "Group By")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsId, "id", "", "", "ID")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsIdentifier, "identifier", "", "", "Identifier")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsImportFilter, "import-filter", "", "", "Import Filter")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsImportId, "import-id", "", "", "Import ID")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsImportRoleArn, "import-role-arn", "", "", "Import Role ARN")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsImportSourceArn, "import-source-arn", "", "", "Import Source ARN")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsImportStatus, "import-status", "", "", "Import Status")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsIncludeLinkedAccounts, "include-linked-accounts", "", "", "Include Linked Accounts")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsIntegrationArn, "integration-arn", "", "", "Integration ARN")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsIntegrationName, "integration-name", "", "", "Integration Name")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsIntegrationNamePrefix, "integration-name-prefix", "", "", "Integration Name Prefix")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsIntegrationStatus, "integration-status", "", "", "Integration Status")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsIntegrationType, "integration-type", "", "", "Integration Type")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsInterleaved, "interleaved", "", "", "Interleaved")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsKmsKeyId, "kms-key-id", "", "", "KMS Key ID")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsLimit, "limit", "", "", "Limit")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsLogEventFilterPattern, "log-event-filter-pattern", "", "", "Log Event Filter Pattern")
	_cloudwatchlogsCmd.Flags().StringSliceVarP(&_cloudwatchlogsLogEventMessages, "log-event-messages", "", nil, "Log Event Messages")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsLogEvents, "log-events", "", "", "Log Events")
	_cloudwatchlogsCmd.Flags().StringSliceVarP(&_cloudwatchlogsLogGroupArnList, "log-group-arn-list", "", nil, "Log Group ARN List")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsLogGroupClass, "log-group-class", "", "", "Log Group Class")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsLogGroupIdentifier, "log-group-identifier", "", "", "Log Group Identifier")
	_cloudwatchlogsCmd.Flags().StringSliceVarP(&_cloudwatchlogsLogGroupIdentifiers, "log-group-identifiers", "", nil, "Log Group Identifiers")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsLogGroupName, "log-group-name", "", "", "Log Group Name")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsLogGroupNamePattern, "log-group-name-pattern", "", "", "Log Group Name Pattern")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsLogGroupNamePrefix, "log-group-name-prefix", "", "", "Log Group Name Prefix")
	_cloudwatchlogsCmd.Flags().StringSliceVarP(&_cloudwatchlogsLogGroupNames, "log-group-names", "", nil, "Log Group Names")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsLogObjectPointer, "log-object-pointer", "", "", "Log Object Pointer")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsLogRecordPointer, "log-record-pointer", "", "", "Log Record Pointer")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsLogStreamName, "log-stream-name", "", "", "Log Stream Name")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsLogStreamNamePrefix, "log-stream-name-prefix", "", "", "Log Stream Name Prefix")
	_cloudwatchlogsCmd.Flags().StringSliceVarP(&_cloudwatchlogsLogStreamNamePrefixes, "log-stream-name-prefixes", "", nil, "Log Stream Name Prefixes")
	_cloudwatchlogsCmd.Flags().StringSliceVarP(&_cloudwatchlogsLogStreamNames, "log-stream-names", "", nil, "Log Stream Names")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsLogType, "log-type", "", "", "Log Type")
	_cloudwatchlogsCmd.Flags().StringSliceVarP(&_cloudwatchlogsLogTypes, "log-types", "", nil, "Log Types")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsMaxResults, "max-results", "", "", "Max Results")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsMetricName, "metric-name", "", "", "Metric Name")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsMetricNamespace, "metric-namespace", "", "", "Metric Namespace")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsMetricTransformations, "metric-transformations", "", "", "Metric Transformations")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsName, "name", "", "", "Name")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsNextToken, "next-token", "", "", "Next Token")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsOrderBy, "order-by", "", "", "Order By")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsOutputFormat, "output-format", "", "", "Output Format")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsPatternId, "pattern-id", "", "", "Pattern ID")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsPolicyDocument, "policy-document", "", "", "Policy Document")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsPolicyName, "policy-name", "", "", "Policy Name")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsPolicyScope, "policy-scope", "", "", "Policy Scope")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsPolicyType, "policy-type", "", "", "Policy Type")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsQueryDefinitionId, "query-definition-id", "", "", "Query Definition ID")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsQueryDefinitionNamePrefix, "query-definition-name-prefix", "", "", "Query Definition Name Prefix")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsQueryId, "query-id", "", "", "Query ID")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsQueryLanguage, "query-language", "", "", "Query Language")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsQueryString, "query-string", "", "", "Query String")
	_cloudwatchlogsCmd.Flags().StringSliceVarP(&_cloudwatchlogsRecordFields, "record-fields", "", nil, "Record Fields")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsResourceArn, "resource-arn", "", "", "Resource ARN")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsResourceConfig, "resource-config", "", "", "Resource Config")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsResourceIdentifier, "resource-identifier", "", "", "Resource Identifier")
	_cloudwatchlogsCmd.Flags().StringSliceVarP(&_cloudwatchlogsResourceTypes, "resource-types", "", nil, "Resource Types")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsRetentionInDays, "retention-in-days", "", "", "Retention In Days")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsRoleArn, "role-arn", "", "", "Role ARN")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsS3DeliveryConfiguration, "s3-delivery-configuration", "", "", "S3 Delivery Configuration")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsScheduleEndTime, "schedule-end-time", "", "", "Schedule End Time")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsScheduleExpression, "schedule-expression", "", "", "Schedule Expression")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsScheduleStartTime, "schedule-start-time", "", "", "Schedule Start Time")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsScope, "scope", "", "", "Scope")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsSelectionCriteria, "selection-criteria", "", "", "Selection Criteria")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsSequenceToken, "sequence-token", "", "", "Sequence Token")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsService, "service", "", "", "Service")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsStartFromHead, "start-from-head", "", "", "Start From Head")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsStartTime, "start-time", "", "", "Start Time")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsStartTimeOffset, "start-time-offset", "", "", "Start Time Offset")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsState, "state", "", "", "State")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsStatus, "status", "", "", "Status")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsStatusCode, "status-code", "", "", "Status Code")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsSuppressionPeriod, "suppression-period", "", "", "Suppression Period")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsSuppressionState, "suppression-state", "", "", "Suppression State")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsSuppressionType, "suppression-type", "", "", "Suppression Type")
	_cloudwatchlogsCmd.Flags().StringSliceVarP(&_cloudwatchlogsTagKeys, "tag-keys", "", nil, "Tag Keys")
	_cloudwatchlogsCmd.Flags().StringSliceVarP(&_cloudwatchlogsTags, "tags", "", nil, "Tags")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsTargetArn, "target-arn", "", "", "Target ARN")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsTaskId, "task-id", "", "", "Task ID")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsTaskName, "task-name", "", "", "Task Name")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsTime, "time", "", "", "Time")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsTimezone, "timezone", "", "", "Timezone")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsTo, "to", "", "", "To")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsTransformerConfig, "transformer-config", "", "", "Transformer Config")
	_cloudwatchlogsCmd.Flags().StringVarP(&_cloudwatchlogsUnmask, "unmask", "", "", "Unmask")

	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsAssociateKmsKey, "associate-kms-key", "", false, "Associate KMS Key")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsAssociateSourceToS3TableIntegration, "associate-source-to-s3-table-integration", "", false, "Associate Source To S3 Table Integration")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsCancelExportTask, "cancel-export-task", "", false, "Cancel Export Task")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsCancelImportTask, "cancel-import-task", "", false, "Cancel Import Task")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsCreateDelivery, "create-delivery", "", false, "Create Delivery")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsCreateExportTask, "create-export-task", "", false, "Create Export Task")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsCreateImportTask, "create-import-task", "", false, "Create Import Task")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsCreateLogAnomalyDetector, "create-log-anomaly-detector", "", false, "Create Log Anomaly Detector")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsCreateLogGroup, "create-log-group", "", false, "Create Log Group")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsCreateLogStream, "create-log-stream", "", false, "Create Log Stream")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsCreateScheduledQuery, "create-scheduled-query", "", false, "Create Scheduled Query")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsDeleteAccountPolicy, "delete-account-policy", "", false, "Delete Account Policy")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsDeleteDataProtectionPolicy, "delete-data-protection-policy", "", false, "Delete Data Protection Policy")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsDeleteDelivery, "delete-delivery", "", false, "Delete Delivery")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsDeleteDeliveryDestination, "delete-delivery-destination", "", false, "Delete Delivery Destination")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsDeleteDeliveryDestinationPolicy, "delete-delivery-destination-policy", "", false, "Delete Delivery Destination Policy")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsDeleteDeliverySource, "delete-delivery-source", "", false, "Delete Delivery Source")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsDeleteDestination, "delete-destination", "", false, "Delete Destination")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsDeleteIndexPolicy, "delete-index-policy", "", false, "Delete Index Policy")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsDeleteIntegration, "delete-integration", "", false, "Delete Integration")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsDeleteLogAnomalyDetector, "delete-log-anomaly-detector", "", false, "Delete Log Anomaly Detector")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsDeleteLogGroup, "delete-log-group", "", false, "Delete Log Group")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsDeleteLogStream, "delete-log-stream", "", false, "Delete Log Stream")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsDeleteMetricFilter, "delete-metric-filter", "", false, "Delete Metric Filter")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsDeleteQueryDefinition, "delete-query-definition", "", false, "Delete Query Definition")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsDeleteResourcePolicy, "delete-resource-policy", "", false, "Delete Resource Policy")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsDeleteRetentionPolicy, "delete-retention-policy", "", false, "Delete Retention Policy")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsDeleteScheduledQuery, "delete-scheduled-query", "", false, "Delete Scheduled Query")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsDeleteSubscriptionFilter, "delete-subscription-filter", "", false, "Delete Subscription Filter")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsDeleteTransformer, "delete-transformer", "", false, "Delete Transformer")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsDescribeAccountPolicies, "describe-account-policies", "", false, "Describe Account Policies")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsDescribeConfigurationTemplates, "describe-configuration-templates", "", false, "Describe Configuration Templates")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsDescribeDeliveries, "describe-deliveries", "", false, "Describe Deliveries")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsDescribeDeliveryDestinations, "describe-delivery-destinations", "", false, "Describe Delivery Destinations")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsDescribeDeliverySources, "describe-delivery-sources", "", false, "Describe Delivery Sources")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsDescribeDestinations, "describe-destinations", "", false, "Describe Destinations")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsDescribeExportTasks, "describe-export-tasks", "", false, "Describe Export Tasks")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsDescribeFieldIndexes, "describe-field-indexes", "", false, "Describe Field Indexes")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsDescribeImportTaskBatches, "describe-import-task-batches", "", false, "Describe Import Task Batches")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsDescribeImportTasks, "describe-import-tasks", "", false, "Describe Import Tasks")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsDescribeIndexPolicies, "describe-index-policies", "", false, "Describe Index Policies")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsDescribeLogGroups, "describe-log-groups", "", false, "Describe Log Groups")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsDescribeLogStreams, "describe-log-streams", "", false, "Describe Log Streams")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsDescribeMetricFilters, "describe-metric-filters", "", false, "Describe Metric Filters")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsDescribeQueries, "describe-queries", "", false, "Describe Queries")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsDescribeQueryDefinitions, "describe-query-definitions", "", false, "Describe Query Definitions")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsDescribeResourcePolicies, "describe-resource-policies", "", false, "Describe Resource Policies")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsDescribeSubscriptionFilters, "describe-subscription-filters", "", false, "Describe Subscription Filters")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsDisassociateKmsKey, "disassociate-kms-key", "", false, "Disassociate KMS Key")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsDisassociateSourceFromS3TableIntegration, "disassociate-source-from-s3-table-integration", "", false, "Disassociate Source From S3 Table Integration")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsFilterLogEvents, "filter-log-events", "", false, "Filter Log Events")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsGetDataProtectionPolicy, "get-data-protection-policy", "", false, "Get Data Protection Policy")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsGetDelivery, "get-delivery", "", false, "Get Delivery")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsGetDeliveryDestination, "get-delivery-destination", "", false, "Get Delivery Destination")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsGetDeliveryDestinationPolicy, "get-delivery-destination-policy", "", false, "Get Delivery Destination Policy")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsGetDeliverySource, "get-delivery-source", "", false, "Get Delivery Source")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsGetIntegration, "get-integration", "", false, "Get Integration")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsGetLogAnomalyDetector, "get-log-anomaly-detector", "", false, "Get Log Anomaly Detector")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsGetLogEvents, "get-log-events", "", false, "Get Log Events")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsGetLogFields, "get-log-fields", "", false, "Get Log Fields")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsGetLogGroupFields, "get-log-group-fields", "", false, "Get Log Group Fields")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsGetLogObject, "get-log-object", "", false, "Get Log Object")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsGetLogRecord, "get-log-record", "", false, "Get Log Record")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsGetQueryResults, "get-query-results", "", false, "Get Query Results")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsGetScheduledQuery, "get-scheduled-query", "", false, "Get Scheduled Query")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsGetScheduledQueryHistory, "get-scheduled-query-history", "", false, "Get Scheduled Query History")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsGetTransformer, "get-transformer", "", false, "Get Transformer")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsListAggregateLogGroupSummaries, "list-aggregate-log-group-summaries", "", false, "List Aggregate Log Group Summaries")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsListAnomalies, "list-anomalies", "", false, "List Anomalies")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsListIntegrations, "list-integrations", "", false, "List Integrations")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsListLogAnomalyDetectors, "list-log-anomaly-detectors", "", false, "List Log Anomaly Detectors")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsListLogGroups, "list-log-groups", "", false, "List Log Groups")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsListLogGroupsForQuery, "list-log-groups-for-query", "", false, "List Log Groups For Query")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsListScheduledQueries, "list-scheduled-queries", "", false, "List Scheduled Queries")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsListSourcesForS3TableIntegration, "list-sources-for-s3-table-integration", "", false, "List Sources For S3 Table Integration")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsListTagsLogGroup, "list-tags-log-group", "", false, "List Tags Log Group")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsPutAccountPolicy, "put-account-policy", "", false, "Put Account Policy")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsPutBearerTokenAuthentication, "put-bearer-token-authentication", "", false, "Put Bearer Token Authentication")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsPutDataProtectionPolicy, "put-data-protection-policy", "", false, "Put Data Protection Policy")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsPutDeliveryDestination, "put-delivery-destination", "", false, "Put Delivery Destination")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsPutDeliveryDestinationPolicy, "put-delivery-destination-policy", "", false, "Put Delivery Destination Policy")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsPutDeliverySource, "put-delivery-source", "", false, "Put Delivery Source")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsPutDestination, "put-destination", "", false, "Put Destination")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsPutDestinationPolicy, "put-destination-policy", "", false, "Put Destination Policy")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsPutIndexPolicy, "put-index-policy", "", false, "Put Index Policy")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsPutIntegration, "put-integration", "", false, "Put Integration")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsPutLogEvents, "put-log-events", "", false, "Put Log Events")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsPutLogGroupDeletionProtection, "put-log-group-deletion-protection", "", false, "Put Log Group Deletion Protection")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsPutMetricFilter, "put-metric-filter", "", false, "Put Metric Filter")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsPutQueryDefinition, "put-query-definition", "", false, "Put Query Definition")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsPutResourcePolicy, "put-resource-policy", "", false, "Put Resource Policy")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsPutRetentionPolicy, "put-retention-policy", "", false, "Put Retention Policy")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsPutSubscriptionFilter, "put-subscription-filter", "", false, "Put Subscription Filter")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsPutTransformer, "put-transformer", "", false, "Put Transformer")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsStartLiveTail, "start-live-tail", "", false, "Start Live Tail")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsStartQuery, "start-query", "", false, "Start Query")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsStopQuery, "stop-query", "", false, "Stop Query")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsTagLogGroup, "tag-log-group", "", false, "Tag Log Group")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsTagResource, "tag-resource", "", false, "Tag Resource")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsTestMetricFilter, "test-metric-filter", "", false, "Test Metric Filter")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsTestTransformer, "test-transformer", "", false, "Test Transformer")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsUntagLogGroup, "untag-log-group", "", false, "Untag Log Group")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsUntagResource, "untag-resource", "", false, "Untag Resource")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsUpdateAnomaly, "update-anomaly", "", false, "Update Anomaly")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsUpdateDeliveryConfiguration, "update-delivery-configuration", "", false, "Update Delivery Configuration")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsUpdateLogAnomalyDetector, "update-log-anomaly-detector", "", false, "Update Log Anomaly Detector")
	_cloudwatchlogsCmd.Flags().BoolVarP(&_cloudwatchlogsUpdateScheduledQuery, "update-scheduled-query", "", false, "Update Scheduled Query")

}
