package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// cloudtrailCmd represents the cloudtrail command
var _cloudtrailCmd = &cobra.Command{
	Use:   "cloudtrail",
	Short: "AWS cloudtrail CLI",
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
		client := cloudtrail.NewFromConfig(cfg)
		if _cloudtrailAddTags {
			cloudtrail_AddTags(cfg, client)
			return
		}
		if _cloudtrailCancelQuery {
			cloudtrail_CancelQuery(cfg, client)
			return
		}
		if _cloudtrailCreateChannel {
			cloudtrail_CreateChannel(cfg, client)
			return
		}
		if _cloudtrailCreateDashboard {
			cloudtrail_CreateDashboard(cfg, client)
			return
		}
		if _cloudtrailCreateEventDataStore {
			cloudtrail_CreateEventDataStore(cfg, client)
			return
		}
		if _cloudtrailCreateTrail {
			cloudtrail_CreateTrail(cfg, client)
			return
		}
		if _cloudtrailDeleteChannel {
			cloudtrail_DeleteChannel(cfg, client)
			return
		}
		if _cloudtrailDeleteDashboard {
			cloudtrail_DeleteDashboard(cfg, client)
			return
		}
		if _cloudtrailDeleteEventDataStore {
			cloudtrail_DeleteEventDataStore(cfg, client)
			return
		}
		if _cloudtrailDeleteResourcePolicy {
			cloudtrail_DeleteResourcePolicy(cfg, client)
			return
		}
		if _cloudtrailDeleteTrail {
			cloudtrail_DeleteTrail(cfg, client)
			return
		}
		if _cloudtrailDeregisterOrganizationDelegatedAdmin {
			cloudtrail_DeregisterOrganizationDelegatedAdmin(cfg, client)
			return
		}
		if _cloudtrailDescribeQuery {
			cloudtrail_DescribeQuery(cfg, client)
			return
		}
		if _cloudtrailDescribeTrails {
			cloudtrail_DescribeTrails(cfg, client)
			return
		}
		if _cloudtrailDisableFederation {
			cloudtrail_DisableFederation(cfg, client)
			return
		}
		if _cloudtrailEnableFederation {
			cloudtrail_EnableFederation(cfg, client)
			return
		}
		if _cloudtrailGenerateQuery {
			cloudtrail_GenerateQuery(cfg, client)
			return
		}
		if _cloudtrailGetChannel {
			cloudtrail_GetChannel(cfg, client)
			return
		}
		if _cloudtrailGetDashboard {
			cloudtrail_GetDashboard(cfg, client)
			return
		}
		if _cloudtrailGetEventConfiguration {
			cloudtrail_GetEventConfiguration(cfg, client)
			return
		}
		if _cloudtrailGetEventDataStore {
			cloudtrail_GetEventDataStore(cfg, client)
			return
		}
		if _cloudtrailGetEventSelectors {
			cloudtrail_GetEventSelectors(cfg, client)
			return
		}
		if _cloudtrailGetImport {
			cloudtrail_GetImport(cfg, client)
			return
		}
		if _cloudtrailGetInsightSelectors {
			cloudtrail_GetInsightSelectors(cfg, client)
			return
		}
		if _cloudtrailGetQueryResults {
			cloudtrail_GetQueryResults(cfg, client)
			return
		}
		if _cloudtrailGetResourcePolicy {
			cloudtrail_GetResourcePolicy(cfg, client)
			return
		}
		if _cloudtrailGetTrail {
			cloudtrail_GetTrail(cfg, client)
			return
		}
		if _cloudtrailGetTrailStatus {
			cloudtrail_GetTrailStatus(cfg, client)
			return
		}
		if _cloudtrailListChannels {
			cloudtrail_ListChannels(cfg, client)
			return
		}
		if _cloudtrailListDashboards {
			cloudtrail_ListDashboards(cfg, client)
			return
		}
		if _cloudtrailListEventDataStores {
			cloudtrail_ListEventDataStores(cfg, client)
			return
		}
		if _cloudtrailListImportFailures {
			cloudtrail_ListImportFailures(cfg, client)
			return
		}
		if _cloudtrailListImports {
			cloudtrail_ListImports(cfg, client)
			return
		}
		if _cloudtrailListInsightsData {
			cloudtrail_ListInsightsData(cfg, client)
			return
		}
		if _cloudtrailListInsightsMetricData {
			cloudtrail_ListInsightsMetricData(cfg, client)
			return
		}
		if _cloudtrailListPublicKeys {
			cloudtrail_ListPublicKeys(cfg, client)
			return
		}
		if _cloudtrailListQueries {
			cloudtrail_ListQueries(cfg, client)
			return
		}
		if _cloudtrailListTags {
			cloudtrail_ListTags(cfg, client)
			return
		}
		if _cloudtrailListTrails {
			cloudtrail_ListTrails(cfg, client)
			return
		}
		if _cloudtrailLookupEvents {
			cloudtrail_LookupEvents(cfg, client)
			return
		}
		if _cloudtrailPutEventConfiguration {
			cloudtrail_PutEventConfiguration(cfg, client)
			return
		}
		if _cloudtrailPutEventSelectors {
			cloudtrail_PutEventSelectors(cfg, client)
			return
		}
		if _cloudtrailPutInsightSelectors {
			cloudtrail_PutInsightSelectors(cfg, client)
			return
		}
		if _cloudtrailPutResourcePolicy {
			cloudtrail_PutResourcePolicy(cfg, client)
			return
		}
		if _cloudtrailRegisterOrganizationDelegatedAdmin {
			cloudtrail_RegisterOrganizationDelegatedAdmin(cfg, client)
			return
		}
		if _cloudtrailRemoveTags {
			cloudtrail_RemoveTags(cfg, client)
			return
		}
		if _cloudtrailRestoreEventDataStore {
			cloudtrail_RestoreEventDataStore(cfg, client)
			return
		}
		if _cloudtrailSearchSampleQueries {
			cloudtrail_SearchSampleQueries(cfg, client)
			return
		}
		if _cloudtrailStartDashboardRefresh {
			cloudtrail_StartDashboardRefresh(cfg, client)
			return
		}
		if _cloudtrailStartEventDataStoreIngestion {
			cloudtrail_StartEventDataStoreIngestion(cfg, client)
			return
		}
		if _cloudtrailStartImport {
			cloudtrail_StartImport(cfg, client)
			return
		}
		if _cloudtrailStartLogging {
			cloudtrail_StartLogging(cfg, client)
			return
		}
		if _cloudtrailStartQuery {
			cloudtrail_StartQuery(cfg, client)
			return
		}
		if _cloudtrailStopEventDataStoreIngestion {
			cloudtrail_StopEventDataStoreIngestion(cfg, client)
			return
		}
		if _cloudtrailStopImport {
			cloudtrail_StopImport(cfg, client)
			return
		}
		if _cloudtrailStopLogging {
			cloudtrail_StopLogging(cfg, client)
			return
		}
		if _cloudtrailUpdateChannel {
			cloudtrail_UpdateChannel(cfg, client)
			return
		}
		if _cloudtrailUpdateDashboard {
			cloudtrail_UpdateDashboard(cfg, client)
			return
		}
		if _cloudtrailUpdateEventDataStore {
			cloudtrail_UpdateEventDataStore(cfg, client)
			return
		}
		if _cloudtrailUpdateTrail {
			cloudtrail_UpdateTrail(cfg, client)
			return
		}

	},
}

var (
	_cloudtrailAddTags                              bool
	_cloudtrailCancelQuery                          bool
	_cloudtrailCreateChannel                        bool
	_cloudtrailCreateDashboard                      bool
	_cloudtrailCreateEventDataStore                 bool
	_cloudtrailCreateTrail                          bool
	_cloudtrailDeleteChannel                        bool
	_cloudtrailDeleteDashboard                      bool
	_cloudtrailDeleteEventDataStore                 bool
	_cloudtrailDeleteResourcePolicy                 bool
	_cloudtrailDeleteTrail                          bool
	_cloudtrailDeregisterOrganizationDelegatedAdmin bool
	_cloudtrailDescribeQuery                        bool
	_cloudtrailDescribeTrails                       bool
	_cloudtrailDisableFederation                    bool
	_cloudtrailEnableFederation                     bool
	_cloudtrailGenerateQuery                        bool
	_cloudtrailGetChannel                           bool
	_cloudtrailGetDashboard                         bool
	_cloudtrailGetEventConfiguration                bool
	_cloudtrailGetEventDataStore                    bool
	_cloudtrailGetEventSelectors                    bool
	_cloudtrailGetImport                            bool
	_cloudtrailGetInsightSelectors                  bool
	_cloudtrailGetQueryResults                      bool
	_cloudtrailGetResourcePolicy                    bool
	_cloudtrailGetTrail                             bool
	_cloudtrailGetTrailStatus                       bool
	_cloudtrailListChannels                         bool
	_cloudtrailListDashboards                       bool
	_cloudtrailListEventDataStores                  bool
	_cloudtrailListImportFailures                   bool
	_cloudtrailListImports                          bool
	_cloudtrailListInsightsData                     bool
	_cloudtrailListInsightsMetricData               bool
	_cloudtrailListPublicKeys                       bool
	_cloudtrailListQueries                          bool
	_cloudtrailListTags                             bool
	_cloudtrailListTrails                           bool
	_cloudtrailLookupEvents                         bool
	_cloudtrailPutEventConfiguration                bool
	_cloudtrailPutEventSelectors                    bool
	_cloudtrailPutInsightSelectors                  bool
	_cloudtrailPutResourcePolicy                    bool
	_cloudtrailRegisterOrganizationDelegatedAdmin   bool
	_cloudtrailRemoveTags                           bool
	_cloudtrailRestoreEventDataStore                bool
	_cloudtrailSearchSampleQueries                  bool
	_cloudtrailStartDashboardRefresh                bool
	_cloudtrailStartEventDataStoreIngestion         bool
	_cloudtrailStartImport                          bool
	_cloudtrailStartLogging                         bool
	_cloudtrailStartQuery                           bool
	_cloudtrailStopEventDataStoreIngestion          bool
	_cloudtrailStopImport                           bool
	_cloudtrailStopLogging                          bool
	_cloudtrailUpdateChannel                        bool
	_cloudtrailUpdateDashboard                      bool
	_cloudtrailUpdateEventDataStore                 bool
	_cloudtrailUpdateTrail                          bool

	_cloudtrailAdvancedEventSelectors       string
	_cloudtrailAggregationConfigurations    string
	_cloudtrailBillingMode                  string
	_cloudtrailChannel                      string
	_cloudtrailCloudWatchLogsLogGroupArn    string
	_cloudtrailCloudWatchLogsRoleArn        string
	_cloudtrailContextKeySelectors          string
	_cloudtrailDashboardId                  string
	_cloudtrailDataType                     string
	_cloudtrailDelegatedAdminAccountId      string
	_cloudtrailDeliveryS3Uri                string
	_cloudtrailDestination                  string
	_cloudtrailDestinations                 string
	_cloudtrailDimensions                   string
	_cloudtrailEnableLogFileValidation      string
	_cloudtrailEndEventTime                 string
	_cloudtrailEndTime                      string
	_cloudtrailErrorCode                    string
	_cloudtrailEventCategory                string
	_cloudtrailEventDataStore               string
	_cloudtrailEventDataStoreOwnerAccountId string
	_cloudtrailEventDataStores              []string
	_cloudtrailEventName                    string
	_cloudtrailEventSelectors               string
	_cloudtrailEventSource                  string
	_cloudtrailFederationRoleArn            string
	_cloudtrailImportId                     string
	_cloudtrailImportSource                 string
	_cloudtrailImportStatus                 string
	_cloudtrailIncludeGlobalServiceEvents   string
	_cloudtrailIncludeShadowTrails          string
	_cloudtrailInsightSelectors             string
	_cloudtrailInsightSource                string
	_cloudtrailInsightType                  string
	_cloudtrailInsightsDestination          string
	_cloudtrailIsMultiRegionTrail           string
	_cloudtrailIsOrganizationTrail          string
	_cloudtrailKmsKeyId                     string
	_cloudtrailLookupAttributes             string
	_cloudtrailMaxEventSize                 string
	_cloudtrailMaxQueryResults              string
	_cloudtrailMaxResults                   string
	_cloudtrailMemberAccountId              string
	_cloudtrailMultiRegionEnabled           string
	_cloudtrailName                         string
	_cloudtrailNamePrefix                   string
	_cloudtrailNextToken                    string
	_cloudtrailOrganizationEnabled          string
	_cloudtrailPeriod                       string
	_cloudtrailPrompt                       string
	_cloudtrailQueryAlias                   string
	_cloudtrailQueryId                      string
	_cloudtrailQueryParameterValues         string
	_cloudtrailQueryParameters              []string
	_cloudtrailQueryStatement               string
	_cloudtrailQueryStatus                  string
	_cloudtrailRefreshId                    string
	_cloudtrailRefreshSchedule              string
	_cloudtrailResourceArn                  string
	_cloudtrailResourceId                   string
	_cloudtrailResourceIdList               []string
	_cloudtrailResourcePolicy               string
	_cloudtrailRetentionPeriod              string
	_cloudtrailS3BucketName                 string
	_cloudtrailS3KeyPrefix                  string
	_cloudtrailSearchPhrase                 string
	_cloudtrailSnsTopicName                 string
	_cloudtrailSource                       string
	_cloudtrailStartEventTime               string
	_cloudtrailStartIngestion               string
	_cloudtrailStartTime                    string
	_cloudtrailTags                         string
	_cloudtrailTagsList                     string
	_cloudtrailTerminationProtectionEnabled string
	_cloudtrailTrailName                    string
	_cloudtrailTrailNameList                []string
	_cloudtrailType                         string
	_cloudtrailWidgets                      string
)

// Adds one or more tags to a trail, event data store, dashboard, or channel, up
// to a limit of 50. Overwrites an existing tag's value when a new value is
// specified for an existing tag key. Tag key names must be unique; you cannot have
// two keys with the same name but different values. If you specify a key without a
// value, the tag will be created with the specified key and a value of null. You
// can tag a trail or event data store that applies to all Amazon Web Services
// Regions only from the Region in which the trail or event data store was created
// (also known as its home Region).
func cloudtrail_AddTags(cfg aws.Config, client *cloudtrail.Client) {
	input := &cloudtrail.AddTagsInput{
		// ResourceId: *string, // Required
		// TagsList: []types.Tag, // Required
	}

	if len(_cloudtrailResourceId) > 0 {
		input.ResourceId = aws.String(_cloudtrailResourceId)
	}
	if len(_cloudtrailTagsList) > 0 {
		if err := assignInputField(input, "TagsList", _cloudtrailTagsList); err != nil {
			log.Errorf("invalid --tags-list: %s", err.Error())
			return
		}
	}

	if resp, err := client.AddTags(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels a query if the query is not in a terminated state, such as CANCELLED ,
// FAILED , TIMED_OUT , or FINISHED . You must specify an ARN value for
// EventDataStore . The ID of the query that you want to cancel is also required.
// When you run CancelQuery , the query status might show as CANCELLED even if the
// operation is not yet finished.
func cloudtrail_CancelQuery(cfg aws.Config, client *cloudtrail.Client) {
	input := &cloudtrail.CancelQueryInput{
		// QueryId: *string, // Required
	}

	if len(_cloudtrailQueryId) > 0 {
		input.QueryId = aws.String(_cloudtrailQueryId)
	}
	if len(_cloudtrailEventDataStore) > 0 {
		input.EventDataStore = aws.String(_cloudtrailEventDataStore)
	}
	if len(_cloudtrailEventDataStoreOwnerAccountId) > 0 {
		input.EventDataStoreOwnerAccountId = aws.String(_cloudtrailEventDataStoreOwnerAccountId)
	}

	if resp, err := client.CancelQuery(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a channel for CloudTrail to ingest events from a partner or external
// source. After you create a channel, a CloudTrail Lake event data store can log
// events from the partner or source that you specify.
func cloudtrail_CreateChannel(cfg aws.Config, client *cloudtrail.Client) {
	input := &cloudtrail.CreateChannelInput{
		// Destinations: []types.Destination, // Required
		// Name: *string, // Required
		// Source: *string, // Required
	}

	if len(_cloudtrailDestinations) > 0 {
		if err := assignInputField(input, "Destinations", _cloudtrailDestinations); err != nil {
			log.Errorf("invalid --destinations: %s", err.Error())
			return
		}
	}
	if len(_cloudtrailName) > 0 {
		input.Name = aws.String(_cloudtrailName)
	}
	if len(_cloudtrailSource) > 0 {
		input.Source = aws.String(_cloudtrailSource)
	}
	if len(_cloudtrailTags) > 0 {
		if err := assignInputField(input, "Tags", _cloudtrailTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a custom dashboard or the Highlights dashboard.
// - Custom dashboards - Custom dashboards allow you to query events in any
// event data store type. You can add up to 10 widgets to a custom dashboard. You
// can manually refresh a custom dashboard, or you can set a refresh schedule.
//
// - Highlights dashboard - You can create the Highlights dashboard to see a
// summary of key user activities and API usage across all your event data stores.
// CloudTrail Lake manages the Highlights dashboard and refreshes the dashboard
// every 6 hours. To create the Highlights dashboard, you must set and enable a
// refresh schedule.
//
// CloudTrail runs queries to populate the dashboard's widgets during a manual or
// scheduled refresh. CloudTrail must be granted permissions to run the StartQuery
// operation on your behalf. To provide permissions, run the PutResourcePolicy
// operation to attach a resource-based policy to each event data store. For more
// information, see [Example: Allow CloudTrail to run queries to populate a dashboard]in the CloudTrail User Guide.
//
// To set a refresh schedule, CloudTrail must be granted permissions to run the
// StartDashboardRefresh operation to refresh the dashboard on your behalf. To
// provide permissions, run the PutResourcePolicy operation to attach a
// resource-based policy to the dashboard. For more information, see [Resource-based policy example for a dashboard]in the
// CloudTrail User Guide.
//
// For more information about dashboards, see [CloudTrail Lake dashboards] in the CloudTrail User Guide.
//
// [CloudTrail Lake dashboards]: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/lake-dashboard.html
// [Example: Allow CloudTrail to run queries to populate a dashboard]: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/security_iam_resource-based-policy-examples.html#security_iam_resource-based-policy-examples-eds-dashboard
// [Resource-based policy example for a dashboard]: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/security_iam_resource-based-policy-examples.html#security_iam_resource-based-policy-examples-dashboards
func cloudtrail_CreateDashboard(cfg aws.Config, client *cloudtrail.Client) {
	input := &cloudtrail.CreateDashboardInput{
		// Name: *string, // Required
	}

	if len(_cloudtrailName) > 0 {
		input.Name = aws.String(_cloudtrailName)
	}
	if len(_cloudtrailRefreshSchedule) > 0 {
		if err := assignInputField(input, "RefreshSchedule", _cloudtrailRefreshSchedule); err != nil {
			log.Errorf("invalid --refresh-schedule: %s", err.Error())
			return
		}
	}
	if len(_cloudtrailTagsList) > 0 {
		if err := assignInputField(input, "TagsList", _cloudtrailTagsList); err != nil {
			log.Errorf("invalid --tags-list: %s", err.Error())
			return
		}
	}
	if len(_cloudtrailTerminationProtectionEnabled) > 0 {
		if err := assignInputField(input, "TerminationProtectionEnabled", _cloudtrailTerminationProtectionEnabled); err != nil {
			log.Errorf("invalid --termination-protection-enabled: %s", err.Error())
			return
		}
	}
	if len(_cloudtrailWidgets) > 0 {
		if err := assignInputField(input, "Widgets", _cloudtrailWidgets); err != nil {
			log.Errorf("invalid --widgets: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDashboard(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new event data store.
func cloudtrail_CreateEventDataStore(cfg aws.Config, client *cloudtrail.Client) {
	input := &cloudtrail.CreateEventDataStoreInput{
		// Name: *string, // Required
	}

	if len(_cloudtrailName) > 0 {
		input.Name = aws.String(_cloudtrailName)
	}
	if len(_cloudtrailAdvancedEventSelectors) > 0 {
		if err := assignInputField(input, "AdvancedEventSelectors", _cloudtrailAdvancedEventSelectors); err != nil {
			log.Errorf("invalid --advanced-event-selectors: %s", err.Error())
			return
		}
	}
	if len(_cloudtrailBillingMode) > 0 {
		if err := assignInputField(input, "BillingMode", _cloudtrailBillingMode); err != nil {
			log.Errorf("invalid --billing-mode: %s", err.Error())
			return
		}
	}
	if len(_cloudtrailKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_cloudtrailKmsKeyId)
	}
	if len(_cloudtrailMultiRegionEnabled) > 0 {
		if err := assignInputField(input, "MultiRegionEnabled", _cloudtrailMultiRegionEnabled); err != nil {
			log.Errorf("invalid --multi-region-enabled: %s", err.Error())
			return
		}
	}
	if len(_cloudtrailOrganizationEnabled) > 0 {
		if err := assignInputField(input, "OrganizationEnabled", _cloudtrailOrganizationEnabled); err != nil {
			log.Errorf("invalid --organization-enabled: %s", err.Error())
			return
		}
	}
	if len(_cloudtrailRetentionPeriod) > 0 {
		if err := assignInputField(input, "RetentionPeriod", _cloudtrailRetentionPeriod); err != nil {
			log.Errorf("invalid --retention-period: %s", err.Error())
			return
		}
	}
	if len(_cloudtrailStartIngestion) > 0 {
		if err := assignInputField(input, "StartIngestion", _cloudtrailStartIngestion); err != nil {
			log.Errorf("invalid --start-ingestion: %s", err.Error())
			return
		}
	}
	if len(_cloudtrailTagsList) > 0 {
		if err := assignInputField(input, "TagsList", _cloudtrailTagsList); err != nil {
			log.Errorf("invalid --tags-list: %s", err.Error())
			return
		}
	}
	if len(_cloudtrailTerminationProtectionEnabled) > 0 {
		if err := assignInputField(input, "TerminationProtectionEnabled", _cloudtrailTerminationProtectionEnabled); err != nil {
			log.Errorf("invalid --termination-protection-enabled: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateEventDataStore(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a trail that specifies the settings for delivery of log data to an
// Amazon S3 bucket.
func cloudtrail_CreateTrail(cfg aws.Config, client *cloudtrail.Client) {
	input := &cloudtrail.CreateTrailInput{
		// Name: *string, // Required
		// S3BucketName: *string, // Required
	}

	if len(_cloudtrailName) > 0 {
		input.Name = aws.String(_cloudtrailName)
	}
	if len(_cloudtrailS3BucketName) > 0 {
		input.S3BucketName = aws.String(_cloudtrailS3BucketName)
	}
	if len(_cloudtrailCloudWatchLogsLogGroupArn) > 0 {
		input.CloudWatchLogsLogGroupArn = aws.String(_cloudtrailCloudWatchLogsLogGroupArn)
	}
	if len(_cloudtrailCloudWatchLogsRoleArn) > 0 {
		input.CloudWatchLogsRoleArn = aws.String(_cloudtrailCloudWatchLogsRoleArn)
	}
	if len(_cloudtrailEnableLogFileValidation) > 0 {
		if err := assignInputField(input, "EnableLogFileValidation", _cloudtrailEnableLogFileValidation); err != nil {
			log.Errorf("invalid --enable-log-file-validation: %s", err.Error())
			return
		}
	}
	if len(_cloudtrailIncludeGlobalServiceEvents) > 0 {
		if err := assignInputField(input, "IncludeGlobalServiceEvents", _cloudtrailIncludeGlobalServiceEvents); err != nil {
			log.Errorf("invalid --include-global-service-events: %s", err.Error())
			return
		}
	}
	if len(_cloudtrailIsMultiRegionTrail) > 0 {
		if err := assignInputField(input, "IsMultiRegionTrail", _cloudtrailIsMultiRegionTrail); err != nil {
			log.Errorf("invalid --is-multi-region-trail: %s", err.Error())
			return
		}
	}
	if len(_cloudtrailIsOrganizationTrail) > 0 {
		if err := assignInputField(input, "IsOrganizationTrail", _cloudtrailIsOrganizationTrail); err != nil {
			log.Errorf("invalid --is-organization-trail: %s", err.Error())
			return
		}
	}
	if len(_cloudtrailKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_cloudtrailKmsKeyId)
	}
	if len(_cloudtrailS3KeyPrefix) > 0 {
		input.S3KeyPrefix = aws.String(_cloudtrailS3KeyPrefix)
	}
	if len(_cloudtrailSnsTopicName) > 0 {
		input.SnsTopicName = aws.String(_cloudtrailSnsTopicName)
	}
	if len(_cloudtrailTagsList) > 0 {
		if err := assignInputField(input, "TagsList", _cloudtrailTagsList); err != nil {
			log.Errorf("invalid --tags-list: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateTrail(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a channel.
func cloudtrail_DeleteChannel(cfg aws.Config, client *cloudtrail.Client) {
	input := &cloudtrail.DeleteChannelInput{
		// Channel: *string, // Required
	}

	if len(_cloudtrailChannel) > 0 {
		input.Channel = aws.String(_cloudtrailChannel)
	}

	if resp, err := client.DeleteChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified dashboard. You cannot delete a dashboard that has
// termination protection enabled.
func cloudtrail_DeleteDashboard(cfg aws.Config, client *cloudtrail.Client) {
	input := &cloudtrail.DeleteDashboardInput{
		// DashboardId: *string, // Required
	}

	if len(_cloudtrailDashboardId) > 0 {
		input.DashboardId = aws.String(_cloudtrailDashboardId)
	}

	if resp, err := client.DeleteDashboard(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disables the event data store specified by EventDataStore , which accepts an
// event data store ARN. After you run DeleteEventDataStore , the event data store
// enters a PENDING_DELETION state, and is automatically deleted after a wait
// period of seven days. TerminationProtectionEnabled must be set to False on the
// event data store and the FederationStatus must be DISABLED . You cannot delete
// an event data store if TerminationProtectionEnabled is True or the
// FederationStatus is ENABLED .
//
// After you run DeleteEventDataStore on an event data store, you cannot run
// ListQueries , DescribeQuery , or GetQueryResults on queries that are using an
// event data store in a PENDING_DELETION state. An event data store in the
// PENDING_DELETION state does not incur costs.
func cloudtrail_DeleteEventDataStore(cfg aws.Config, client *cloudtrail.Client) {
	input := &cloudtrail.DeleteEventDataStoreInput{
		// EventDataStore: *string, // Required
	}

	if len(_cloudtrailEventDataStore) > 0 {
		input.EventDataStore = aws.String(_cloudtrailEventDataStore)
	}

	if resp, err := client.DeleteEventDataStore(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the resource-based policy attached to the CloudTrail event data store,
// dashboard, or channel.
func cloudtrail_DeleteResourcePolicy(cfg aws.Config, client *cloudtrail.Client) {
	input := &cloudtrail.DeleteResourcePolicyInput{
		// ResourceArn: *string, // Required
	}

	if len(_cloudtrailResourceArn) > 0 {
		input.ResourceArn = aws.String(_cloudtrailResourceArn)
	}

	if resp, err := client.DeleteResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a trail. This operation must be called from the Region in which the
// trail was created. DeleteTrail cannot be called on the shadow trails
// (replicated trails in other Regions) of a trail that is enabled in all Regions.
//
// While deleting a CloudTrail trail is an irreversible action, CloudTrail does
// not delete log files in the Amazon S3 bucket for that trail, the Amazon S3
// bucket itself, or the CloudWatchlog group to which the trail delivers events.
// Deleting a multi-Region trail will stop logging of events in all Amazon Web
// Services Regions enabled in your Amazon Web Services account. Deleting a
// single-Region trail will stop logging of events in that Region only. It will not
// stop logging of events in other Regions even if the trails in those other
// Regions have identical names to the deleted trail.
//
// For information about account closure and deletion of CloudTrail trails, see [https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-account-closure.html].
//
// [https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-account-closure.html]: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-account-closure.html
func cloudtrail_DeleteTrail(cfg aws.Config, client *cloudtrail.Client) {
	input := &cloudtrail.DeleteTrailInput{
		// Name: *string, // Required
	}

	if len(_cloudtrailName) > 0 {
		input.Name = aws.String(_cloudtrailName)
	}

	if resp, err := client.DeleteTrail(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes CloudTrail delegated administrator permissions from a member account in
// an organization.
func cloudtrail_DeregisterOrganizationDelegatedAdmin(cfg aws.Config, client *cloudtrail.Client) {
	input := &cloudtrail.DeregisterOrganizationDelegatedAdminInput{
		// DelegatedAdminAccountId: *string, // Required
	}

	if len(_cloudtrailDelegatedAdminAccountId) > 0 {
		input.DelegatedAdminAccountId = aws.String(_cloudtrailDelegatedAdminAccountId)
	}

	if resp, err := client.DeregisterOrganizationDelegatedAdmin(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns metadata about a query, including query run time in milliseconds,
// number of events scanned and matched, and query status. If the query results
// were delivered to an S3 bucket, the response also provides the S3 URI and the
// delivery status.
//
// You must specify either QueryId or QueryAlias . Specifying the QueryAlias
// parameter returns information about the last query run for the alias. You can
// provide RefreshId along with QueryAlias to view the query results of a
// dashboard query for the specified RefreshId .
func cloudtrail_DescribeQuery(cfg aws.Config, client *cloudtrail.Client) {
	input := &cloudtrail.DescribeQueryInput{}

	if len(_cloudtrailEventDataStore) > 0 {
		input.EventDataStore = aws.String(_cloudtrailEventDataStore)
	}
	if len(_cloudtrailEventDataStoreOwnerAccountId) > 0 {
		input.EventDataStoreOwnerAccountId = aws.String(_cloudtrailEventDataStoreOwnerAccountId)
	}
	if len(_cloudtrailQueryAlias) > 0 {
		input.QueryAlias = aws.String(_cloudtrailQueryAlias)
	}
	if len(_cloudtrailQueryId) > 0 {
		input.QueryId = aws.String(_cloudtrailQueryId)
	}
	if len(_cloudtrailRefreshId) > 0 {
		input.RefreshId = aws.String(_cloudtrailRefreshId)
	}

	if resp, err := client.DescribeQuery(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves settings for one or more trails associated with the current Region
// for your account.
func cloudtrail_DescribeTrails(cfg aws.Config, client *cloudtrail.Client) {
	input := &cloudtrail.DescribeTrailsInput{}

	if len(_cloudtrailIncludeShadowTrails) > 0 {
		if err := assignInputField(input, "IncludeShadowTrails", _cloudtrailIncludeShadowTrails); err != nil {
			log.Errorf("invalid --include-shadow-trails: %s", err.Error())
			return
		}
	}
	if len(_cloudtrailTrailNameList) > 0 {
		input.TrailNameList = append([]string(nil), _cloudtrailTrailNameList...)
	}

	if resp, err := client.DescribeTrails(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disables Lake query federation on the specified event data store. When you
// disable federation, CloudTrail disables the integration with Glue, Lake
// Formation, and Amazon Athena. After disabling Lake query federation, you can no
// longer query your event data in Amazon Athena.
//
// No CloudTrail Lake data is deleted when you disable federation and you can
// continue to run queries in CloudTrail Lake.
func cloudtrail_DisableFederation(cfg aws.Config, client *cloudtrail.Client) {
	input := &cloudtrail.DisableFederationInput{
		// EventDataStore: *string, // Required
	}

	if len(_cloudtrailEventDataStore) > 0 {
		input.EventDataStore = aws.String(_cloudtrailEventDataStore)
	}

	if resp, err := client.DisableFederation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables Lake query federation on the specified event data store. Federating an
// event data store lets you view the metadata associated with the event data store
// in the Glue [Data Catalog]and run SQL queries against your event data using Amazon Athena.
// The table metadata stored in the Glue Data Catalog lets the Athena query engine
// know how to find, read, and process the data that you want to query.
//
// When you enable Lake query federation, CloudTrail creates a managed database
// named aws:cloudtrail (if the database doesn't already exist) and a managed
// federated table in the Glue Data Catalog. The event data store ID is used for
// the table name. CloudTrail registers the role ARN and event data store in [Lake Formation], the
// service responsible for allowing fine-grained access control of the federated
// resources in the Glue Data Catalog.
//
// For more information about Lake query federation, see [Federate an event data store].
//
// [Federate an event data store]: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/query-federation.html
// [Lake Formation]: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/query-federation-lake-formation.html
// [Data Catalog]: https://docs.aws.amazon.com/glue/latest/dg/components-overview.html#data-catalog-intro
func cloudtrail_EnableFederation(cfg aws.Config, client *cloudtrail.Client) {
	input := &cloudtrail.EnableFederationInput{
		// EventDataStore: *string, // Required
		// FederationRoleArn: *string, // Required
	}

	if len(_cloudtrailEventDataStore) > 0 {
		input.EventDataStore = aws.String(_cloudtrailEventDataStore)
	}
	if len(_cloudtrailFederationRoleArn) > 0 {
		input.FederationRoleArn = aws.String(_cloudtrailFederationRoleArn)
	}

	if resp, err := client.EnableFederation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Generates a query from a natural language prompt. This operation uses
// generative artificial intelligence (generative AI) to produce a ready-to-use SQL
// query from the prompt.
//
// The prompt can be a question or a statement about the event data in your event
// data store. For example, you can enter prompts like "What are my top errors in
// the past month?" and “Give me a list of users that used SNS.”
//
// The prompt must be in English. For information about limitations, permissions,
// and supported Regions, see [Create CloudTrail Lake queries from natural language prompts]in the CloudTrail user guide.
//
// Do not include any personally identifying, confidential, or sensitive
// information in your prompts.
//
// This feature uses generative AI large language models (LLMs); we recommend
// double-checking the LLM response.
//
// [Create CloudTrail Lake queries from natural language prompts]: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/lake-query-generator.html
func cloudtrail_GenerateQuery(cfg aws.Config, client *cloudtrail.Client) {
	input := &cloudtrail.GenerateQueryInput{
		// EventDataStores: []string, // Required
		// Prompt: *string, // Required
	}

	if len(_cloudtrailEventDataStores) > 0 {
		input.EventDataStores = append([]string(nil), _cloudtrailEventDataStores...)
	}
	if len(_cloudtrailPrompt) > 0 {
		input.Prompt = aws.String(_cloudtrailPrompt)
	}

	if resp, err := client.GenerateQuery(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a specific channel.
func cloudtrail_GetChannel(cfg aws.Config, client *cloudtrail.Client) {
	input := &cloudtrail.GetChannelInput{
		// Channel: *string, // Required
	}

	if len(_cloudtrailChannel) > 0 {
		input.Channel = aws.String(_cloudtrailChannel)
	}

	if resp, err := client.GetChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the specified dashboard.
func cloudtrail_GetDashboard(cfg aws.Config, client *cloudtrail.Client) {
	input := &cloudtrail.GetDashboardInput{
		// DashboardId: *string, // Required
	}

	if len(_cloudtrailDashboardId) > 0 {
		input.DashboardId = aws.String(_cloudtrailDashboardId)
	}

	if resp, err := client.GetDashboard(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the current event configuration settings for the specified event data
// store or trail. The response includes maximum event size configuration, the
// context key selectors configured for the event data store, and any aggregation
// settings configured for the trail.
func cloudtrail_GetEventConfiguration(cfg aws.Config, client *cloudtrail.Client) {
	input := &cloudtrail.GetEventConfigurationInput{}

	if len(_cloudtrailEventDataStore) > 0 {
		input.EventDataStore = aws.String(_cloudtrailEventDataStore)
	}
	if len(_cloudtrailTrailName) > 0 {
		input.TrailName = aws.String(_cloudtrailTrailName)
	}

	if resp, err := client.GetEventConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about an event data store specified as either an ARN or the
// ID portion of the ARN.
func cloudtrail_GetEventDataStore(cfg aws.Config, client *cloudtrail.Client) {
	input := &cloudtrail.GetEventDataStoreInput{
		// EventDataStore: *string, // Required
	}

	if len(_cloudtrailEventDataStore) > 0 {
		input.EventDataStore = aws.String(_cloudtrailEventDataStore)
	}

	if resp, err := client.GetEventDataStore(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the settings for the event selectors that you configured for your
// trail. The information returned for your event selectors includes the following:
//
// - If your event selector includes read-only events, write-only events, or all
// events. This applies to management events, data events, and network activity
// events.
//
// - If your event selector includes management events.
//
// - If your event selector includes network activity events, the event sources
// for which you are logging network activity events.
//
// - If your event selector includes data events, the resources on which you are
// logging data events.
//
// For more information about logging management, data, and network activity
// events, see the following topics in the CloudTrail User Guide:
//
// [Logging management events]
//
// [Logging data events]
//
// [Logging network activity events]
//
// [Logging network activity events]: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-network-events-with-cloudtrail.html
// [Logging management events]: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-management-events-with-cloudtrail.html
// [Logging data events]: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.html
func cloudtrail_GetEventSelectors(cfg aws.Config, client *cloudtrail.Client) {
	input := &cloudtrail.GetEventSelectorsInput{
		// TrailName: *string, // Required
	}

	if len(_cloudtrailTrailName) > 0 {
		input.TrailName = aws.String(_cloudtrailTrailName)
	}

	if resp, err := client.GetEventSelectors(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a specific import.
func cloudtrail_GetImport(cfg aws.Config, client *cloudtrail.Client) {
	input := &cloudtrail.GetImportInput{
		// ImportId: *string, // Required
	}

	if len(_cloudtrailImportId) > 0 {
		input.ImportId = aws.String(_cloudtrailImportId)
	}

	if resp, err := client.GetImport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the settings for the Insights event selectors that you configured for
// your trail or event data store. GetInsightSelectors shows if CloudTrail
// Insights logging is enabled and which Insights types are configured with
// corresponding event categories. If you run GetInsightSelectors on a trail or
// event data store that does not have Insights events enabled, the operation
// throws the exception InsightNotEnabledException
//
// Specify either the EventDataStore parameter to get Insights event selectors for
// an event data store, or the TrailName parameter to the get Insights event
// selectors for a trail. You cannot specify these parameters together.
//
// For more information, see [Working with CloudTrail Insights] in the CloudTrail User Guide.
//
// [Working with CloudTrail Insights]: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-insights-events-with-cloudtrail.html
func cloudtrail_GetInsightSelectors(cfg aws.Config, client *cloudtrail.Client) {
	input := &cloudtrail.GetInsightSelectorsInput{}

	if len(_cloudtrailEventDataStore) > 0 {
		input.EventDataStore = aws.String(_cloudtrailEventDataStore)
	}
	if len(_cloudtrailTrailName) > 0 {
		input.TrailName = aws.String(_cloudtrailTrailName)
	}

	if resp, err := client.GetInsightSelectors(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets event data results of a query. You must specify the QueryID value returned
// by the StartQuery operation.
func cloudtrail_GetQueryResults(cfg aws.Config, client *cloudtrail.Client) {
	input := &cloudtrail.GetQueryResultsInput{
		// QueryId: *string, // Required
	}

	if len(_cloudtrailQueryId) > 0 {
		input.QueryId = aws.String(_cloudtrailQueryId)
	}
	if len(_cloudtrailEventDataStore) > 0 {
		input.EventDataStore = aws.String(_cloudtrailEventDataStore)
	}
	if len(_cloudtrailEventDataStoreOwnerAccountId) > 0 {
		input.EventDataStoreOwnerAccountId = aws.String(_cloudtrailEventDataStoreOwnerAccountId)
	}
	if len(_cloudtrailMaxQueryResults) > 0 {
		if err := assignInputField(input, "MaxQueryResults", _cloudtrailMaxQueryResults); err != nil {
			log.Errorf("invalid --max-query-results: %s", err.Error())
			return
		}
	}
	if len(_cloudtrailNextToken) > 0 {
		input.NextToken = aws.String(_cloudtrailNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetQueryResults(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudtrail.GetQueryResultsOutput
	p := cloudtrail.NewGetQueryResultsPaginator(client, input)
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

// Retrieves the JSON text of the resource-based policy document attached to the
// CloudTrail event data store, dashboard, or channel.
func cloudtrail_GetResourcePolicy(cfg aws.Config, client *cloudtrail.Client) {
	input := &cloudtrail.GetResourcePolicyInput{
		// ResourceArn: *string, // Required
	}

	if len(_cloudtrailResourceArn) > 0 {
		input.ResourceArn = aws.String(_cloudtrailResourceArn)
	}

	if resp, err := client.GetResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns settings information for a specified trail.
func cloudtrail_GetTrail(cfg aws.Config, client *cloudtrail.Client) {
	input := &cloudtrail.GetTrailInput{
		// Name: *string, // Required
	}

	if len(_cloudtrailName) > 0 {
		input.Name = aws.String(_cloudtrailName)
	}

	if resp, err := client.GetTrail(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a JSON-formatted list of information about the specified trail. Fields
// include information on delivery errors, Amazon SNS and Amazon S3 errors, and
// start and stop logging times for each trail. This operation returns trail status
// from a single Region. To return trail status from all Regions, you must call the
// operation on each Region.
func cloudtrail_GetTrailStatus(cfg aws.Config, client *cloudtrail.Client) {
	input := &cloudtrail.GetTrailStatusInput{
		// Name: *string, // Required
	}

	if len(_cloudtrailName) > 0 {
		input.Name = aws.String(_cloudtrailName)
	}

	if resp, err := client.GetTrailStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the channels in the current account, and their source names.
func cloudtrail_ListChannels(cfg aws.Config, client *cloudtrail.Client) {
	input := &cloudtrail.ListChannelsInput{}

	if len(_cloudtrailMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cloudtrailMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cloudtrailNextToken) > 0 {
		input.NextToken = aws.String(_cloudtrailNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListChannels(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudtrail.ListChannelsOutput
	p := cloudtrail.NewListChannelsPaginator(client, input)
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

// Returns information about all dashboards in the account, in the current
// Region.
func cloudtrail_ListDashboards(cfg aws.Config, client *cloudtrail.Client) {
	input := &cloudtrail.ListDashboardsInput{}

	if len(_cloudtrailMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cloudtrailMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cloudtrailNamePrefix) > 0 {
		input.NamePrefix = aws.String(_cloudtrailNamePrefix)
	}
	if len(_cloudtrailNextToken) > 0 {
		input.NextToken = aws.String(_cloudtrailNextToken)
	}
	if len(_cloudtrailType) > 0 {
		if err := assignInputField(input, "Type", _cloudtrailType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListDashboards(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about all event data stores in the account, in the current
// Region.
func cloudtrail_ListEventDataStores(cfg aws.Config, client *cloudtrail.Client) {
	input := &cloudtrail.ListEventDataStoresInput{}

	if len(_cloudtrailMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cloudtrailMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cloudtrailNextToken) > 0 {
		input.NextToken = aws.String(_cloudtrailNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListEventDataStores(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudtrail.ListEventDataStoresOutput
	p := cloudtrail.NewListEventDataStoresPaginator(client, input)
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

// Returns a list of failures for the specified import.
func cloudtrail_ListImportFailures(cfg aws.Config, client *cloudtrail.Client) {
	input := &cloudtrail.ListImportFailuresInput{
		// ImportId: *string, // Required
	}

	if len(_cloudtrailImportId) > 0 {
		input.ImportId = aws.String(_cloudtrailImportId)
	}
	if len(_cloudtrailMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cloudtrailMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cloudtrailNextToken) > 0 {
		input.NextToken = aws.String(_cloudtrailNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListImportFailures(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudtrail.ListImportFailuresOutput
	p := cloudtrail.NewListImportFailuresPaginator(client, input)
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

// Returns information on all imports, or a select set of imports by ImportStatus
// or Destination .
func cloudtrail_ListImports(cfg aws.Config, client *cloudtrail.Client) {
	input := &cloudtrail.ListImportsInput{}

	if len(_cloudtrailDestination) > 0 {
		input.Destination = aws.String(_cloudtrailDestination)
	}
	if len(_cloudtrailImportStatus) > 0 {
		if err := assignInputField(input, "ImportStatus", _cloudtrailImportStatus); err != nil {
			log.Errorf("invalid --import-status: %s", err.Error())
			return
		}
	}
	if len(_cloudtrailMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cloudtrailMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cloudtrailNextToken) > 0 {
		input.NextToken = aws.String(_cloudtrailNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListImports(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudtrail.ListImportsOutput
	p := cloudtrail.NewListImportsPaginator(client, input)
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

// Returns Insights events generated on a trail that logs data events. You can
// list Insights events that occurred in a Region within the last 90 days.
//
// ListInsightsData supports the following Dimensions for Insights events:
//
// - Event ID
//
// - Event name
//
// - Event source
//
// All dimensions are optional. The default number of results returned is 50, with
// a maximum of 50 possible. The response includes a token that you can use to get
// the next page of results.
//
// The rate of ListInsightsData requests is limited to two per second, per
// account, per Region. If this limit is exceeded, a throttling error occurs.
func cloudtrail_ListInsightsData(cfg aws.Config, client *cloudtrail.Client) {
	input := &cloudtrail.ListInsightsDataInput{
		// DataType: types.ListInsightsDataType, // Required
		// InsightSource: *string, // Required
	}

	if len(_cloudtrailDataType) > 0 {
		if err := assignInputField(input, "DataType", _cloudtrailDataType); err != nil {
			log.Errorf("invalid --data-type: %s", err.Error())
			return
		}
	}
	if len(_cloudtrailInsightSource) > 0 {
		input.InsightSource = aws.String(_cloudtrailInsightSource)
	}
	if len(_cloudtrailDimensions) > 0 {
		if err := assignInputField(input, "Dimensions", _cloudtrailDimensions); err != nil {
			log.Errorf("invalid --dimensions: %s", err.Error())
			return
		}
	}
	if len(_cloudtrailEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _cloudtrailEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_cloudtrailMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cloudtrailMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cloudtrailNextToken) > 0 {
		input.NextToken = aws.String(_cloudtrailNextToken)
	}
	if len(_cloudtrailStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _cloudtrailStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListInsightsData(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudtrail.ListInsightsDataOutput
	p := cloudtrail.NewListInsightsDataPaginator(client, input)
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

// Returns Insights metrics data for trails that have enabled Insights. The
// request must include the EventSource , EventName , and InsightType parameters.
//
// If the InsightType is set to ApiErrorRateInsight , the request must also include
// the ErrorCode parameter.
//
// The following are the available time periods for ListInsightsMetricData . Each
// cutoff is inclusive.
//
// - Data points with a period of 60 seconds (1-minute) are available for 15
// days.
//
// - Data points with a period of 300 seconds (5-minute) are available for 63
// days.
//
// - Data points with a period of 3600 seconds (1 hour) are available for 90
// days.
//
// To use ListInsightsMetricData operation, you must have the following
// permissions:
//
// - If ListInsightsMetricData is invoked with TrailName parameter, access to the
// ListInsightsMetricData API operation is linked to the cloudtrail:LookupEvents
// action and cloudtrail:ListInsightsData . To use this operation, you must have
// permissions to perform the cloudtrail:LookupEvents and
// cloudtrail:ListInsightsData action on the specific trail.
//
// - If ListInsightsMetricData is invoked without TrailName parameter, access to
// the ListInsightsMetricData API operation is linked to the
// cloudtrail:LookupEvents action only. To use this operation, you must have
// permissions to perform the cloudtrail:LookupEvents action.
func cloudtrail_ListInsightsMetricData(cfg aws.Config, client *cloudtrail.Client) {
	input := &cloudtrail.ListInsightsMetricDataInput{
		// EventName: *string, // Required
		// EventSource: *string, // Required
		// InsightType: types.InsightType, // Required
	}

	if len(_cloudtrailEventName) > 0 {
		input.EventName = aws.String(_cloudtrailEventName)
	}
	if len(_cloudtrailEventSource) > 0 {
		input.EventSource = aws.String(_cloudtrailEventSource)
	}
	if len(_cloudtrailInsightType) > 0 {
		if err := assignInputField(input, "InsightType", _cloudtrailInsightType); err != nil {
			log.Errorf("invalid --insight-type: %s", err.Error())
			return
		}
	}
	if len(_cloudtrailDataType) > 0 {
		if err := assignInputField(input, "DataType", _cloudtrailDataType); err != nil {
			log.Errorf("invalid --data-type: %s", err.Error())
			return
		}
	}
	if len(_cloudtrailEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _cloudtrailEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_cloudtrailErrorCode) > 0 {
		input.ErrorCode = aws.String(_cloudtrailErrorCode)
	}
	if len(_cloudtrailMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cloudtrailMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cloudtrailNextToken) > 0 {
		input.NextToken = aws.String(_cloudtrailNextToken)
	}
	if len(_cloudtrailPeriod) > 0 {
		if err := assignInputField(input, "Period", _cloudtrailPeriod); err != nil {
			log.Errorf("invalid --period: %s", err.Error())
			return
		}
	}
	if len(_cloudtrailStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _cloudtrailStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_cloudtrailTrailName) > 0 {
		input.TrailName = aws.String(_cloudtrailTrailName)
	}

	if disablePaginator() {
		if resp, err := client.ListInsightsMetricData(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudtrail.ListInsightsMetricDataOutput
	p := cloudtrail.NewListInsightsMetricDataPaginator(client, input)
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

// Returns all public keys whose private keys were used to sign the digest files
// within the specified time range. The public key is needed to validate digest
// files that were signed with its corresponding private key.
//
// CloudTrail uses different private and public key pairs per Region. Each digest
// file is signed with a private key unique to its Region. When you validate a
// digest file from a specific Region, you must look in the same Region for its
// corresponding public key.
func cloudtrail_ListPublicKeys(cfg aws.Config, client *cloudtrail.Client) {
	input := &cloudtrail.ListPublicKeysInput{}

	if len(_cloudtrailEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _cloudtrailEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_cloudtrailNextToken) > 0 {
		input.NextToken = aws.String(_cloudtrailNextToken)
	}
	if len(_cloudtrailStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _cloudtrailStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListPublicKeys(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudtrail.ListPublicKeysOutput
	p := cloudtrail.NewListPublicKeysPaginator(client, input)
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

// Returns a list of queries and query statuses for the past seven days. You must
// specify an ARN value for EventDataStore . Optionally, to shorten the list of
// results, you can specify a time range, formatted as timestamps, by adding
// StartTime and EndTime parameters, and a QueryStatus value. Valid values for
// QueryStatus include QUEUED , RUNNING , FINISHED , FAILED , TIMED_OUT , or
// CANCELLED .
func cloudtrail_ListQueries(cfg aws.Config, client *cloudtrail.Client) {
	input := &cloudtrail.ListQueriesInput{
		// EventDataStore: *string, // Required
	}

	if len(_cloudtrailEventDataStore) > 0 {
		input.EventDataStore = aws.String(_cloudtrailEventDataStore)
	}
	if len(_cloudtrailEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _cloudtrailEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_cloudtrailMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cloudtrailMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cloudtrailNextToken) > 0 {
		input.NextToken = aws.String(_cloudtrailNextToken)
	}
	if len(_cloudtrailQueryStatus) > 0 {
		if err := assignInputField(input, "QueryStatus", _cloudtrailQueryStatus); err != nil {
			log.Errorf("invalid --query-status: %s", err.Error())
			return
		}
	}
	if len(_cloudtrailStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _cloudtrailStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListQueries(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudtrail.ListQueriesOutput
	p := cloudtrail.NewListQueriesPaginator(client, input)
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

// Lists the tags for the specified trails, event data stores, dashboards, or
// channels in the current Region.
func cloudtrail_ListTags(cfg aws.Config, client *cloudtrail.Client) {
	input := &cloudtrail.ListTagsInput{
		// ResourceIdList: []string, // Required
	}

	if len(_cloudtrailResourceIdList) > 0 {
		input.ResourceIdList = append([]string(nil), _cloudtrailResourceIdList...)
	}
	if len(_cloudtrailNextToken) > 0 {
		input.NextToken = aws.String(_cloudtrailNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTags(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudtrail.ListTagsOutput
	p := cloudtrail.NewListTagsPaginator(client, input)
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

// Lists trails that are in the current account.
func cloudtrail_ListTrails(cfg aws.Config, client *cloudtrail.Client) {
	input := &cloudtrail.ListTrailsInput{}

	if len(_cloudtrailNextToken) > 0 {
		input.NextToken = aws.String(_cloudtrailNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTrails(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudtrail.ListTrailsOutput
	p := cloudtrail.NewListTrailsPaginator(client, input)
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

// Looks up [management events] or [CloudTrail Insights events] that are captured by CloudTrail. You can look up events that
// occurred in a Region within the last 90 days.
//
// LookupEvents returns recent Insights events for trails that enable Insights. To
// view Insights events for an event data store, you can run queries on your
// Insights event data store, and you can also view the Lake dashboard for
// Insights.
//
// Lookup supports the following attributes for management events:
//
// - Amazon Web Services access key
//
// - Event ID
//
// - Event name
//
// - Event source
//
// - Read only
//
// - Resource name
//
// - Resource type
//
// - User name
//
// Lookup supports the following attributes for Insights events:
//
// - Event ID
//
// - Event name
//
// - Event source
//
// All attributes are optional. The default number of results returned is 50, with
// a maximum of 50 possible. The response includes a token that you can use to get
// the next page of results.
//
// The rate of lookup requests is limited to two per second, per account, per
// Region. If this limit is exceeded, a throttling error occurs.
//
// [CloudTrail Insights events]: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-concepts.html#cloudtrail-concepts-insights-events
// [management events]: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-concepts.html#cloudtrail-concepts-management-events
func cloudtrail_LookupEvents(cfg aws.Config, client *cloudtrail.Client) {
	input := &cloudtrail.LookupEventsInput{}

	if len(_cloudtrailEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _cloudtrailEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_cloudtrailEventCategory) > 0 {
		if err := assignInputField(input, "EventCategory", _cloudtrailEventCategory); err != nil {
			log.Errorf("invalid --event-category: %s", err.Error())
			return
		}
	}
	if len(_cloudtrailLookupAttributes) > 0 {
		if err := assignInputField(input, "LookupAttributes", _cloudtrailLookupAttributes); err != nil {
			log.Errorf("invalid --lookup-attributes: %s", err.Error())
			return
		}
	}
	if len(_cloudtrailMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cloudtrailMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cloudtrailNextToken) > 0 {
		input.NextToken = aws.String(_cloudtrailNextToken)
	}
	if len(_cloudtrailStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _cloudtrailStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.LookupEvents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudtrail.LookupEventsOutput
	p := cloudtrail.NewLookupEventsPaginator(client, input)
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

// Updates the event configuration settings for the specified event data store or
// trail. This operation supports updating the maximum event size, adding or
// modifying context key selectors for event data store, and configuring
// aggregation settings for the trail.
func cloudtrail_PutEventConfiguration(cfg aws.Config, client *cloudtrail.Client) {
	input := &cloudtrail.PutEventConfigurationInput{}

	if len(_cloudtrailAggregationConfigurations) > 0 {
		if err := assignInputField(input, "AggregationConfigurations", _cloudtrailAggregationConfigurations); err != nil {
			log.Errorf("invalid --aggregation-configurations: %s", err.Error())
			return
		}
	}
	if len(_cloudtrailContextKeySelectors) > 0 {
		if err := assignInputField(input, "ContextKeySelectors", _cloudtrailContextKeySelectors); err != nil {
			log.Errorf("invalid --context-key-selectors: %s", err.Error())
			return
		}
	}
	if len(_cloudtrailEventDataStore) > 0 {
		input.EventDataStore = aws.String(_cloudtrailEventDataStore)
	}
	if len(_cloudtrailMaxEventSize) > 0 {
		if err := assignInputField(input, "MaxEventSize", _cloudtrailMaxEventSize); err != nil {
			log.Errorf("invalid --max-event-size: %s", err.Error())
			return
		}
	}
	if len(_cloudtrailTrailName) > 0 {
		input.TrailName = aws.String(_cloudtrailTrailName)
	}

	if resp, err := client.PutEventConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Configures event selectors (also referred to as basic event selectors) or
// advanced event selectors for your trail. You can use either
// AdvancedEventSelectors or EventSelectors , but not both. If you apply
// AdvancedEventSelectors to a trail, any existing EventSelectors are overwritten.
//
// You can use AdvancedEventSelectors to log management events, data events for
// all resource types, and network activity events.
//
// You can use EventSelectors to log management events and data events for the
// following resource types:
//
// - AWS::DynamoDB::Table
//
// - AWS::Lambda::Function
//
// - AWS::S3::Object
//
// You can't use EventSelectors to log network activity events.
//
// If you want your trail to log Insights events, be sure the event selector or
// advanced event selector enables logging of the Insights event types you want
// configured for your trail. For more information about logging Insights events,
// see [Working with CloudTrail Insights]in the CloudTrail User Guide. By default, trails created without specific
// event selectors are configured to log all read and write management events, and
// no data events or network activity events.
//
// When an event occurs in your account, CloudTrail evaluates the event selectors
// or advanced event selectors in all trails. For each trail, if the event matches
// any event selector, the trail processes and logs the event. If the event doesn't
// match any event selector, the trail doesn't log the event.
//
// # Example
//
// - You create an event selector for a trail and specify that you want to log
// write-only events.
//
// - The EC2 GetConsoleOutput and RunInstances API operations occur in your
// account.
//
// - CloudTrail evaluates whether the events match your event selectors.
//
// - The RunInstances is a write-only event and it matches your event selector.
// The trail logs the event.
//
// - The GetConsoleOutput is a read-only event that doesn't match your event
// selector. The trail doesn't log the event.
//
// The PutEventSelectors operation must be called from the Region in which the
// trail was created; otherwise, an InvalidHomeRegionException exception is thrown.
//
// You can configure up to five event selectors for each trail.
//
// You can add advanced event selectors, and conditions for your advanced event
// selectors, up to a maximum of 500 values for all conditions and selectors on a
// trail. For more information, see [Logging management events], [Logging data events], [Logging network activity events], and [Quotas in CloudTrail] in the CloudTrail User Guide.
//
// [Logging network activity events]: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-network-events-with-cloudtrail.html
// [Logging management events]: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-management-events-with-cloudtrail.html
// [Working with CloudTrail Insights]: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-insights-events-with-cloudtrail.html
// [Quotas in CloudTrail]: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/WhatIsCloudTrail-Limits.html
// [Logging data events]: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.html
func cloudtrail_PutEventSelectors(cfg aws.Config, client *cloudtrail.Client) {
	input := &cloudtrail.PutEventSelectorsInput{
		// TrailName: *string, // Required
	}

	if len(_cloudtrailTrailName) > 0 {
		input.TrailName = aws.String(_cloudtrailTrailName)
	}
	if len(_cloudtrailAdvancedEventSelectors) > 0 {
		if err := assignInputField(input, "AdvancedEventSelectors", _cloudtrailAdvancedEventSelectors); err != nil {
			log.Errorf("invalid --advanced-event-selectors: %s", err.Error())
			return
		}
	}
	if len(_cloudtrailEventSelectors) > 0 {
		if err := assignInputField(input, "EventSelectors", _cloudtrailEventSelectors); err != nil {
			log.Errorf("invalid --event-selectors: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutEventSelectors(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lets you enable Insights event logging on specific event categories by
// specifying the Insights selectors that you want to enable on an existing trail
// or event data store. You also use PutInsightSelectors to turn off Insights
// event logging, by passing an empty list of Insights types. The valid Insights
// event types are ApiErrorRateInsight and ApiCallRateInsight , and valid
// EventCategories are Management and Data .
//
// Insights on data events are not supported on event data stores. For event data
// stores, you can only enable Insights on management events.
//
// To enable Insights on an event data store, you must specify the ARNs (or ID
// suffix of the ARNs) for the source event data store ( EventDataStore ) and the
// destination event data store ( InsightsDestination ). The source event data
// store logs management events and enables Insights. The destination event data
// store logs Insights events based upon the management event activity of the
// source event data store. The source and destination event data stores must
// belong to the same Amazon Web Services account.
//
// To log Insights events for a trail, you must specify the name ( TrailName ) of
// the CloudTrail trail for which you want to change or add Insights selectors.
//
// - For Management events Insights: To log CloudTrail Insights on the API call
// rate, the trail or event data store must log write management events. To log
// CloudTrail Insights on the API error rate, the trail or event data store must
// log read or write management events.
//
// - For Data events Insights: To log CloudTrail Insights on the API call rate
// or API error rate, the trail must log read or write data events. Data events
// Insights are not supported on event data store.
//
// To log CloudTrail Insights events on API call volume, the trail or event data
// store must log write management events. To log CloudTrail Insights events on
// API error rate, the trail or event data store must log read or write management
// events. You can call GetEventSelectors on a trail to check whether the trail
// logs management events. You can call GetEventDataStore on an event data store
// to check whether the event data store logs management events.
//
// For more information, see [Working with CloudTrail Insights] in the CloudTrail User Guide.
//
// [Working with CloudTrail Insights]: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-insights-events-with-cloudtrail.html
func cloudtrail_PutInsightSelectors(cfg aws.Config, client *cloudtrail.Client) {
	input := &cloudtrail.PutInsightSelectorsInput{
		// InsightSelectors: []types.InsightSelector, // Required
	}

	if len(_cloudtrailInsightSelectors) > 0 {
		if err := assignInputField(input, "InsightSelectors", _cloudtrailInsightSelectors); err != nil {
			log.Errorf("invalid --insight-selectors: %s", err.Error())
			return
		}
	}
	if len(_cloudtrailEventDataStore) > 0 {
		input.EventDataStore = aws.String(_cloudtrailEventDataStore)
	}
	if len(_cloudtrailInsightsDestination) > 0 {
		input.InsightsDestination = aws.String(_cloudtrailInsightsDestination)
	}
	if len(_cloudtrailTrailName) > 0 {
		input.TrailName = aws.String(_cloudtrailTrailName)
	}

	if resp, err := client.PutInsightSelectors(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Attaches a resource-based permission policy to a CloudTrail event data store,
// dashboard, or channel. For more information about resource-based policies, see [CloudTrail resource-based policy examples]
// in the CloudTrail User Guide.
//
// [CloudTrail resource-based policy examples]: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/security_iam_resource-based-policy-examples.html
func cloudtrail_PutResourcePolicy(cfg aws.Config, client *cloudtrail.Client) {
	input := &cloudtrail.PutResourcePolicyInput{
		// ResourceArn: *string, // Required
		// ResourcePolicy: *string, // Required
	}

	if len(_cloudtrailResourceArn) > 0 {
		input.ResourceArn = aws.String(_cloudtrailResourceArn)
	}
	if len(_cloudtrailResourcePolicy) > 0 {
		input.ResourcePolicy = aws.String(_cloudtrailResourcePolicy)
	}

	if resp, err := client.PutResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Registers an organization’s member account as the CloudTrail [delegated administrator].
//
// [delegated administrator]: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-delegated-administrator.html
func cloudtrail_RegisterOrganizationDelegatedAdmin(cfg aws.Config, client *cloudtrail.Client) {
	input := &cloudtrail.RegisterOrganizationDelegatedAdminInput{
		// MemberAccountId: *string, // Required
	}

	if len(_cloudtrailMemberAccountId) > 0 {
		input.MemberAccountId = aws.String(_cloudtrailMemberAccountId)
	}

	if resp, err := client.RegisterOrganizationDelegatedAdmin(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the specified tags from a trail, event data store, dashboard, or
// channel.
func cloudtrail_RemoveTags(cfg aws.Config, client *cloudtrail.Client) {
	input := &cloudtrail.RemoveTagsInput{
		// ResourceId: *string, // Required
		// TagsList: []types.Tag, // Required
	}

	if len(_cloudtrailResourceId) > 0 {
		input.ResourceId = aws.String(_cloudtrailResourceId)
	}
	if len(_cloudtrailTagsList) > 0 {
		if err := assignInputField(input, "TagsList", _cloudtrailTagsList); err != nil {
			log.Errorf("invalid --tags-list: %s", err.Error())
			return
		}
	}

	if resp, err := client.RemoveTags(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Restores a deleted event data store specified by EventDataStore , which accepts
// an event data store ARN. You can only restore a deleted event data store within
// the seven-day wait period after deletion. Restoring an event data store can take
// several minutes, depending on the size of the event data store.
func cloudtrail_RestoreEventDataStore(cfg aws.Config, client *cloudtrail.Client) {
	input := &cloudtrail.RestoreEventDataStoreInput{
		// EventDataStore: *string, // Required
	}

	if len(_cloudtrailEventDataStore) > 0 {
		input.EventDataStore = aws.String(_cloudtrailEventDataStore)
	}

	if resp, err := client.RestoreEventDataStore(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Searches sample queries and returns a list of sample queries that are sorted
// by relevance. To search for sample queries, provide a natural language
// SearchPhrase in English.
func cloudtrail_SearchSampleQueries(cfg aws.Config, client *cloudtrail.Client) {
	input := &cloudtrail.SearchSampleQueriesInput{
		// SearchPhrase: *string, // Required
	}

	if len(_cloudtrailSearchPhrase) > 0 {
		input.SearchPhrase = aws.String(_cloudtrailSearchPhrase)
	}
	if len(_cloudtrailMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cloudtrailMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cloudtrailNextToken) > 0 {
		input.NextToken = aws.String(_cloudtrailNextToken)
	}

	if resp, err := client.SearchSampleQueries(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a refresh of the specified dashboard.
// Each time a dashboard is refreshed, CloudTrail runs queries to populate the
// dashboard's widgets. CloudTrail must be granted permissions to run the
// StartQuery operation on your behalf. To provide permissions, run the
// PutResourcePolicy operation to attach a resource-based policy to each event data
// store. For more information, see [Example: Allow CloudTrail to run queries to populate a dashboard]in the CloudTrail User Guide.
//
// [Example: Allow CloudTrail to run queries to populate a dashboard]: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/security_iam_resource-based-policy-examples.html#security_iam_resource-based-policy-examples-eds-dashboard
func cloudtrail_StartDashboardRefresh(cfg aws.Config, client *cloudtrail.Client) {
	input := &cloudtrail.StartDashboardRefreshInput{
		// DashboardId: *string, // Required
	}

	if len(_cloudtrailDashboardId) > 0 {
		input.DashboardId = aws.String(_cloudtrailDashboardId)
	}
	if len(_cloudtrailQueryParameterValues) > 0 {
		if err := assignInputField(input, "QueryParameterValues", _cloudtrailQueryParameterValues); err != nil {
			log.Errorf("invalid --query-parameter-values: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartDashboardRefresh(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts the ingestion of live events on an event data store specified as either
// an ARN or the ID portion of the ARN. To start ingestion, the event data store
// Status must be STOPPED_INGESTION and the eventCategory must be Management , Data
// , NetworkActivity , or ConfigurationItem .
func cloudtrail_StartEventDataStoreIngestion(cfg aws.Config, client *cloudtrail.Client) {
	input := &cloudtrail.StartEventDataStoreIngestionInput{
		// EventDataStore: *string, // Required
	}

	if len(_cloudtrailEventDataStore) > 0 {
		input.EventDataStore = aws.String(_cloudtrailEventDataStore)
	}

	if resp, err := client.StartEventDataStoreIngestion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts an import of logged trail events from a source S3 bucket to a
// destination event data store. By default, CloudTrail only imports events
// contained in the S3 bucket's CloudTrail prefix and the prefixes inside the
// CloudTrail prefix, and does not check prefixes for other Amazon Web Services
// services. If you want to import CloudTrail events contained in another prefix,
// you must include the prefix in the S3LocationUri . For more considerations about
// importing trail events, see [Considerations for copying trail events]in the CloudTrail User Guide.
//
// When you start a new import, the Destinations and ImportSource parameters are
// required. Before starting a new import, disable any access control lists (ACLs)
// attached to the source S3 bucket. For more information about disabling ACLs, see
// [Controlling ownership of objects and disabling ACLs for your bucket].
//
// When you retry an import, the ImportID parameter is required.
//
// If the destination event data store is for an organization, you must use the
// management account to import trail events. You cannot use the delegated
// administrator account for the organization.
//
// [Considerations for copying trail events]: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-copy-trail-to-lake.html#cloudtrail-trail-copy-considerations
// [Controlling ownership of objects and disabling ACLs for your bucket]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/about-object-ownership.html
func cloudtrail_StartImport(cfg aws.Config, client *cloudtrail.Client) {
	input := &cloudtrail.StartImportInput{}

	if len(_cloudtrailDestinations) > 0 {
		input.Destinations = []string{_cloudtrailDestinations}
	}
	if len(_cloudtrailEndEventTime) > 0 {
		if err := assignInputField(input, "EndEventTime", _cloudtrailEndEventTime); err != nil {
			log.Errorf("invalid --end-event-time: %s", err.Error())
			return
		}
	}
	if len(_cloudtrailImportId) > 0 {
		input.ImportId = aws.String(_cloudtrailImportId)
	}
	if len(_cloudtrailImportSource) > 0 {
		if err := assignInputField(input, "ImportSource", _cloudtrailImportSource); err != nil {
			log.Errorf("invalid --import-source: %s", err.Error())
			return
		}
	}
	if len(_cloudtrailStartEventTime) > 0 {
		if err := assignInputField(input, "StartEventTime", _cloudtrailStartEventTime); err != nil {
			log.Errorf("invalid --start-event-time: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartImport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts the recording of Amazon Web Services API calls and log file delivery for
// a trail. For a trail that is enabled in all Regions, this operation must be
// called from the Region in which the trail was created. This operation cannot be
// called on the shadow trails (replicated trails in other Regions) of a trail that
// is enabled in all Regions.
func cloudtrail_StartLogging(cfg aws.Config, client *cloudtrail.Client) {
	input := &cloudtrail.StartLoggingInput{
		// Name: *string, // Required
	}

	if len(_cloudtrailName) > 0 {
		input.Name = aws.String(_cloudtrailName)
	}

	if resp, err := client.StartLogging(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a CloudTrail Lake query. Use the QueryStatement parameter to provide
// your SQL query, enclosed in single quotation marks. Use the optional
// DeliveryS3Uri parameter to deliver the query results to an S3 bucket.
//
// StartQuery requires you specify either the QueryStatement parameter, or a
// QueryAlias and any QueryParameters . In the current release, the QueryAlias and
// QueryParameters parameters are used only for the queries that populate the
// CloudTrail Lake dashboards.
func cloudtrail_StartQuery(cfg aws.Config, client *cloudtrail.Client) {
	input := &cloudtrail.StartQueryInput{}

	if len(_cloudtrailDeliveryS3Uri) > 0 {
		input.DeliveryS3Uri = aws.String(_cloudtrailDeliveryS3Uri)
	}
	if len(_cloudtrailEventDataStoreOwnerAccountId) > 0 {
		input.EventDataStoreOwnerAccountId = aws.String(_cloudtrailEventDataStoreOwnerAccountId)
	}
	if len(_cloudtrailQueryAlias) > 0 {
		input.QueryAlias = aws.String(_cloudtrailQueryAlias)
	}
	if len(_cloudtrailQueryParameters) > 0 {
		input.QueryParameters = append([]string(nil), _cloudtrailQueryParameters...)
	}
	if len(_cloudtrailQueryStatement) > 0 {
		input.QueryStatement = aws.String(_cloudtrailQueryStatement)
	}

	if resp, err := client.StartQuery(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops the ingestion of live events on an event data store specified as either
// an ARN or the ID portion of the ARN. To stop ingestion, the event data store
// Status must be ENABLED and the eventCategory must be Management , Data ,
// NetworkActivity , or ConfigurationItem .
func cloudtrail_StopEventDataStoreIngestion(cfg aws.Config, client *cloudtrail.Client) {
	input := &cloudtrail.StopEventDataStoreIngestionInput{
		// EventDataStore: *string, // Required
	}

	if len(_cloudtrailEventDataStore) > 0 {
		input.EventDataStore = aws.String(_cloudtrailEventDataStore)
	}

	if resp, err := client.StopEventDataStoreIngestion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a specified import.
func cloudtrail_StopImport(cfg aws.Config, client *cloudtrail.Client) {
	input := &cloudtrail.StopImportInput{
		// ImportId: *string, // Required
	}

	if len(_cloudtrailImportId) > 0 {
		input.ImportId = aws.String(_cloudtrailImportId)
	}

	if resp, err := client.StopImport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Suspends the recording of Amazon Web Services API calls and log file delivery
// for the specified trail. Under most circumstances, there is no need to use this
// action. You can update a trail without stopping it first. This action is the
// only way to stop recording. For a trail enabled in all Regions, this operation
// must be called from the Region in which the trail was created, or an
// InvalidHomeRegionException will occur. This operation cannot be called on the
// shadow trails (replicated trails in other Regions) of a trail enabled in all
// Regions.
func cloudtrail_StopLogging(cfg aws.Config, client *cloudtrail.Client) {
	input := &cloudtrail.StopLoggingInput{
		// Name: *string, // Required
	}

	if len(_cloudtrailName) > 0 {
		input.Name = aws.String(_cloudtrailName)
	}

	if resp, err := client.StopLogging(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a channel specified by a required channel ARN or UUID.
func cloudtrail_UpdateChannel(cfg aws.Config, client *cloudtrail.Client) {
	input := &cloudtrail.UpdateChannelInput{
		// Channel: *string, // Required
	}

	if len(_cloudtrailChannel) > 0 {
		input.Channel = aws.String(_cloudtrailChannel)
	}
	if len(_cloudtrailDestinations) > 0 {
		if err := assignInputField(input, "Destinations", _cloudtrailDestinations); err != nil {
			log.Errorf("invalid --destinations: %s", err.Error())
			return
		}
	}
	if len(_cloudtrailName) > 0 {
		input.Name = aws.String(_cloudtrailName)
	}

	if resp, err := client.UpdateChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified dashboard.
// To set a refresh schedule, CloudTrail must be granted permissions to run the
// StartDashboardRefresh operation to refresh the dashboard on your behalf. To
// provide permissions, run the PutResourcePolicy operation to attach a
// resource-based policy to the dashboard. For more information, see [Resource-based policy example for a dashboard]in the
// CloudTrail User Guide.
//
// CloudTrail runs queries to populate the dashboard's widgets during a manual or
// scheduled refresh. CloudTrail must be granted permissions to run the StartQuery
// operation on your behalf. To provide permissions, run the PutResourcePolicy
// operation to attach a resource-based policy to each event data store. For more
// information, see [Example: Allow CloudTrail to run queries to populate a dashboard]in the CloudTrail User Guide.
//
// [Example: Allow CloudTrail to run queries to populate a dashboard]: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/security_iam_resource-based-policy-examples.html#security_iam_resource-based-policy-examples-eds-dashboard
// [Resource-based policy example for a dashboard]: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/security_iam_resource-based-policy-examples.html#security_iam_resource-based-policy-examples-dashboards
func cloudtrail_UpdateDashboard(cfg aws.Config, client *cloudtrail.Client) {
	input := &cloudtrail.UpdateDashboardInput{
		// DashboardId: *string, // Required
	}

	if len(_cloudtrailDashboardId) > 0 {
		input.DashboardId = aws.String(_cloudtrailDashboardId)
	}
	if len(_cloudtrailRefreshSchedule) > 0 {
		if err := assignInputField(input, "RefreshSchedule", _cloudtrailRefreshSchedule); err != nil {
			log.Errorf("invalid --refresh-schedule: %s", err.Error())
			return
		}
	}
	if len(_cloudtrailTerminationProtectionEnabled) > 0 {
		if err := assignInputField(input, "TerminationProtectionEnabled", _cloudtrailTerminationProtectionEnabled); err != nil {
			log.Errorf("invalid --termination-protection-enabled: %s", err.Error())
			return
		}
	}
	if len(_cloudtrailWidgets) > 0 {
		if err := assignInputField(input, "Widgets", _cloudtrailWidgets); err != nil {
			log.Errorf("invalid --widgets: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateDashboard(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an event data store. The required EventDataStore value is an ARN or the
// ID portion of the ARN. Other parameters are optional, but at least one optional
// parameter must be specified, or CloudTrail throws an error. RetentionPeriod is
// in days, and valid values are integers between 7 and 3653 if the BillingMode is
// set to EXTENDABLE_RETENTION_PRICING , or between 7 and 2557 if BillingMode is
// set to FIXED_RETENTION_PRICING . By default, TerminationProtection is enabled.
//
// For event data stores for CloudTrail events, AdvancedEventSelectors includes or
// excludes management, data, or network activity events in your event data store.
// For more information about AdvancedEventSelectors , see [AdvancedEventSelectors].
//
// For event data stores for CloudTrail Insights events, Config configuration
// items, Audit Manager evidence, or non-Amazon Web Services events,
// AdvancedEventSelectors includes events of that type in your event data store.
//
// [AdvancedEventSelectors]: https://docs.aws.amazon.com/awscloudtrail/latest/APIReference/API_AdvancedEventSelector.html
func cloudtrail_UpdateEventDataStore(cfg aws.Config, client *cloudtrail.Client) {
	input := &cloudtrail.UpdateEventDataStoreInput{
		// EventDataStore: *string, // Required
	}

	if len(_cloudtrailEventDataStore) > 0 {
		input.EventDataStore = aws.String(_cloudtrailEventDataStore)
	}
	if len(_cloudtrailAdvancedEventSelectors) > 0 {
		if err := assignInputField(input, "AdvancedEventSelectors", _cloudtrailAdvancedEventSelectors); err != nil {
			log.Errorf("invalid --advanced-event-selectors: %s", err.Error())
			return
		}
	}
	if len(_cloudtrailBillingMode) > 0 {
		if err := assignInputField(input, "BillingMode", _cloudtrailBillingMode); err != nil {
			log.Errorf("invalid --billing-mode: %s", err.Error())
			return
		}
	}
	if len(_cloudtrailKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_cloudtrailKmsKeyId)
	}
	if len(_cloudtrailMultiRegionEnabled) > 0 {
		if err := assignInputField(input, "MultiRegionEnabled", _cloudtrailMultiRegionEnabled); err != nil {
			log.Errorf("invalid --multi-region-enabled: %s", err.Error())
			return
		}
	}
	if len(_cloudtrailName) > 0 {
		input.Name = aws.String(_cloudtrailName)
	}
	if len(_cloudtrailOrganizationEnabled) > 0 {
		if err := assignInputField(input, "OrganizationEnabled", _cloudtrailOrganizationEnabled); err != nil {
			log.Errorf("invalid --organization-enabled: %s", err.Error())
			return
		}
	}
	if len(_cloudtrailRetentionPeriod) > 0 {
		if err := assignInputField(input, "RetentionPeriod", _cloudtrailRetentionPeriod); err != nil {
			log.Errorf("invalid --retention-period: %s", err.Error())
			return
		}
	}
	if len(_cloudtrailTerminationProtectionEnabled) > 0 {
		if err := assignInputField(input, "TerminationProtectionEnabled", _cloudtrailTerminationProtectionEnabled); err != nil {
			log.Errorf("invalid --termination-protection-enabled: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateEventDataStore(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates trail settings that control what events you are logging, and how to
// handle log files. Changes to a trail do not require stopping the CloudTrail
// service. Use this action to designate an existing bucket for log delivery. If
// the existing bucket has previously been a target for CloudTrail log files, an
// IAM policy exists for the bucket. UpdateTrail must be called from the Region in
// which the trail was created; otherwise, an InvalidHomeRegionException is thrown.
func cloudtrail_UpdateTrail(cfg aws.Config, client *cloudtrail.Client) {
	input := &cloudtrail.UpdateTrailInput{
		// Name: *string, // Required
	}

	if len(_cloudtrailName) > 0 {
		input.Name = aws.String(_cloudtrailName)
	}
	if len(_cloudtrailCloudWatchLogsLogGroupArn) > 0 {
		input.CloudWatchLogsLogGroupArn = aws.String(_cloudtrailCloudWatchLogsLogGroupArn)
	}
	if len(_cloudtrailCloudWatchLogsRoleArn) > 0 {
		input.CloudWatchLogsRoleArn = aws.String(_cloudtrailCloudWatchLogsRoleArn)
	}
	if len(_cloudtrailEnableLogFileValidation) > 0 {
		if err := assignInputField(input, "EnableLogFileValidation", _cloudtrailEnableLogFileValidation); err != nil {
			log.Errorf("invalid --enable-log-file-validation: %s", err.Error())
			return
		}
	}
	if len(_cloudtrailIncludeGlobalServiceEvents) > 0 {
		if err := assignInputField(input, "IncludeGlobalServiceEvents", _cloudtrailIncludeGlobalServiceEvents); err != nil {
			log.Errorf("invalid --include-global-service-events: %s", err.Error())
			return
		}
	}
	if len(_cloudtrailIsMultiRegionTrail) > 0 {
		if err := assignInputField(input, "IsMultiRegionTrail", _cloudtrailIsMultiRegionTrail); err != nil {
			log.Errorf("invalid --is-multi-region-trail: %s", err.Error())
			return
		}
	}
	if len(_cloudtrailIsOrganizationTrail) > 0 {
		if err := assignInputField(input, "IsOrganizationTrail", _cloudtrailIsOrganizationTrail); err != nil {
			log.Errorf("invalid --is-organization-trail: %s", err.Error())
			return
		}
	}
	if len(_cloudtrailKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_cloudtrailKmsKeyId)
	}
	if len(_cloudtrailS3BucketName) > 0 {
		input.S3BucketName = aws.String(_cloudtrailS3BucketName)
	}
	if len(_cloudtrailS3KeyPrefix) > 0 {
		input.S3KeyPrefix = aws.String(_cloudtrailS3KeyPrefix)
	}
	if len(_cloudtrailSnsTopicName) > 0 {
		input.SnsTopicName = aws.String(_cloudtrailSnsTopicName)
	}

	if resp, err := client.UpdateTrail(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_cloudtrailCmd)
	_cloudtrailCmd.Flags().SortFlags = false

	_cloudtrailCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_cloudtrailCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_cloudtrailCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailAdvancedEventSelectors, "advanced-event-selectors", "", "", "Advanced Event Selectors")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailAggregationConfigurations, "aggregation-configurations", "", "", "Aggregation Configurations")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailBillingMode, "billing-mode", "", "", "Billing Mode")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailChannel, "channel", "", "", "Channel")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailCloudWatchLogsLogGroupArn, "cloud-watch-logs-log-group-arn", "", "", "Cloud Watch Logs Log Group ARN")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailCloudWatchLogsRoleArn, "cloud-watch-logs-role-arn", "", "", "Cloud Watch Logs Role ARN")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailContextKeySelectors, "context-key-selectors", "", "", "Context Key Selectors")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailDashboardId, "dashboard-id", "", "", "Dashboard ID")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailDataType, "data-type", "", "", "Data Type")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailDelegatedAdminAccountId, "delegated-admin-account-id", "", "", "Delegated Admin Account ID")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailDeliveryS3Uri, "delivery-s3-uri", "", "", "Delivery S3 URI")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailDestination, "destination", "", "", "Destination")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailDestinations, "destinations", "", "", "Destinations")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailDimensions, "dimensions", "", "", "Dimensions")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailEnableLogFileValidation, "enable-log-file-validation", "", "", "Enable Log File Validation")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailEndEventTime, "end-event-time", "", "", "End Event Time")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailEndTime, "end-time", "", "", "End Time")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailErrorCode, "error-code", "", "", "Error Code")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailEventCategory, "event-category", "", "", "Event Category")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailEventDataStore, "event-data-store", "", "", "Event Data Store")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailEventDataStoreOwnerAccountId, "event-data-store-owner-account-id", "", "", "Event Data Store Owner Account ID")
	_cloudtrailCmd.Flags().StringSliceVarP(&_cloudtrailEventDataStores, "event-data-stores", "", nil, "Event Data Stores")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailEventName, "event-name", "", "", "Event Name")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailEventSelectors, "event-selectors", "", "", "Event Selectors")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailEventSource, "event-source", "", "", "Event Source")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailFederationRoleArn, "federation-role-arn", "", "", "Federation Role ARN")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailImportId, "import-id", "", "", "Import ID")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailImportSource, "import-source", "", "", "Import Source")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailImportStatus, "import-status", "", "", "Import Status")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailIncludeGlobalServiceEvents, "include-global-service-events", "", "", "Include Global Service Events")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailIncludeShadowTrails, "include-shadow-trails", "", "", "Include Shadow Trails")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailInsightSelectors, "insight-selectors", "", "", "Insight Selectors")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailInsightSource, "insight-source", "", "", "Insight Source")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailInsightType, "insight-type", "", "", "Insight Type")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailInsightsDestination, "insights-destination", "", "", "Insights Destination")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailIsMultiRegionTrail, "is-multi-region-trail", "", "", "Is Multi Region Trail")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailIsOrganizationTrail, "is-organization-trail", "", "", "Is Organization Trail")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailKmsKeyId, "kms-key-id", "", "", "KMS Key ID")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailLookupAttributes, "lookup-attributes", "", "", "Lookup Attributes")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailMaxEventSize, "max-event-size", "", "", "Max Event Size")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailMaxQueryResults, "max-query-results", "", "", "Max Query Results")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailMaxResults, "max-results", "", "", "Max Results")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailMemberAccountId, "member-account-id", "", "", "Member Account ID")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailMultiRegionEnabled, "multi-region-enabled", "", "", "Multi Region Enabled")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailName, "name", "", "", "Name")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailNamePrefix, "name-prefix", "", "", "Name Prefix")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailNextToken, "next-token", "", "", "Next Token")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailOrganizationEnabled, "organization-enabled", "", "", "Organization Enabled")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailPeriod, "period", "", "", "Period")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailPrompt, "prompt", "", "", "Prompt")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailQueryAlias, "query-alias", "", "", "Query Alias")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailQueryId, "query-id", "", "", "Query ID")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailQueryParameterValues, "query-parameter-values", "", "", "Query Parameter Values")
	_cloudtrailCmd.Flags().StringSliceVarP(&_cloudtrailQueryParameters, "query-parameters", "", nil, "Query Parameters")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailQueryStatement, "query-statement", "", "", "Query Statement")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailQueryStatus, "query-status", "", "", "Query Status")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailRefreshId, "refresh-id", "", "", "Refresh ID")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailRefreshSchedule, "refresh-schedule", "", "", "Refresh Schedule")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailResourceArn, "resource-arn", "", "", "Resource ARN")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailResourceId, "resource-id", "", "", "Resource ID")
	_cloudtrailCmd.Flags().StringSliceVarP(&_cloudtrailResourceIdList, "resource-id-list", "", nil, "Resource ID List")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailResourcePolicy, "resource-policy", "", "", "Resource Policy")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailRetentionPeriod, "retention-period", "", "", "Retention Period")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailS3BucketName, "s3-bucket-name", "", "", "S3 Bucket Name")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailS3KeyPrefix, "s3-key-prefix", "", "", "S3 Key Prefix")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailSearchPhrase, "search-phrase", "", "", "Search Phrase")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailSnsTopicName, "sns-topic-name", "", "", "SNS Topic Name")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailSource, "source", "", "", "Source")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailStartEventTime, "start-event-time", "", "", "Start Event Time")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailStartIngestion, "start-ingestion", "", "", "Start Ingestion")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailStartTime, "start-time", "", "", "Start Time")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailTags, "tags", "", "", "Tags")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailTagsList, "tags-list", "", "", "Tags List")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailTerminationProtectionEnabled, "termination-protection-enabled", "", "", "Termination Protection Enabled")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailTrailName, "trail-name", "", "", "Trail Name")
	_cloudtrailCmd.Flags().StringSliceVarP(&_cloudtrailTrailNameList, "trail-name-list", "", nil, "Trail Name List")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailType, "type", "", "", "Type")
	_cloudtrailCmd.Flags().StringVarP(&_cloudtrailWidgets, "widgets", "", "", "Widgets")

	_cloudtrailCmd.Flags().BoolVarP(&_cloudtrailAddTags, "add-tags", "", false, "Add Tags")
	_cloudtrailCmd.Flags().BoolVarP(&_cloudtrailCancelQuery, "cancel-query", "", false, "Cancel Query")
	_cloudtrailCmd.Flags().BoolVarP(&_cloudtrailCreateChannel, "create-channel", "", false, "Create Channel")
	_cloudtrailCmd.Flags().BoolVarP(&_cloudtrailCreateDashboard, "create-dashboard", "", false, "Create Dashboard")
	_cloudtrailCmd.Flags().BoolVarP(&_cloudtrailCreateEventDataStore, "create-event-data-store", "", false, "Create Event Data Store")
	_cloudtrailCmd.Flags().BoolVarP(&_cloudtrailCreateTrail, "create-trail", "", false, "Create Trail")
	_cloudtrailCmd.Flags().BoolVarP(&_cloudtrailDeleteChannel, "delete-channel", "", false, "Delete Channel")
	_cloudtrailCmd.Flags().BoolVarP(&_cloudtrailDeleteDashboard, "delete-dashboard", "", false, "Delete Dashboard")
	_cloudtrailCmd.Flags().BoolVarP(&_cloudtrailDeleteEventDataStore, "delete-event-data-store", "", false, "Delete Event Data Store")
	_cloudtrailCmd.Flags().BoolVarP(&_cloudtrailDeleteResourcePolicy, "delete-resource-policy", "", false, "Delete Resource Policy")
	_cloudtrailCmd.Flags().BoolVarP(&_cloudtrailDeleteTrail, "delete-trail", "", false, "Delete Trail")
	_cloudtrailCmd.Flags().BoolVarP(&_cloudtrailDeregisterOrganizationDelegatedAdmin, "deregister-organization-delegated-admin", "", false, "Deregister Organization Delegated Admin")
	_cloudtrailCmd.Flags().BoolVarP(&_cloudtrailDescribeQuery, "describe-query", "", false, "Describe Query")
	_cloudtrailCmd.Flags().BoolVarP(&_cloudtrailDescribeTrails, "describe-trails", "", false, "Describe Trails")
	_cloudtrailCmd.Flags().BoolVarP(&_cloudtrailDisableFederation, "disable-federation", "", false, "Disable Federation")
	_cloudtrailCmd.Flags().BoolVarP(&_cloudtrailEnableFederation, "enable-federation", "", false, "Enable Federation")
	_cloudtrailCmd.Flags().BoolVarP(&_cloudtrailGenerateQuery, "generate-query", "", false, "Generate Query")
	_cloudtrailCmd.Flags().BoolVarP(&_cloudtrailGetChannel, "get-channel", "", false, "Get Channel")
	_cloudtrailCmd.Flags().BoolVarP(&_cloudtrailGetDashboard, "get-dashboard", "", false, "Get Dashboard")
	_cloudtrailCmd.Flags().BoolVarP(&_cloudtrailGetEventConfiguration, "get-event-configuration", "", false, "Get Event Configuration")
	_cloudtrailCmd.Flags().BoolVarP(&_cloudtrailGetEventDataStore, "get-event-data-store", "", false, "Get Event Data Store")
	_cloudtrailCmd.Flags().BoolVarP(&_cloudtrailGetEventSelectors, "get-event-selectors", "", false, "Get Event Selectors")
	_cloudtrailCmd.Flags().BoolVarP(&_cloudtrailGetImport, "get-import", "", false, "Get Import")
	_cloudtrailCmd.Flags().BoolVarP(&_cloudtrailGetInsightSelectors, "get-insight-selectors", "", false, "Get Insight Selectors")
	_cloudtrailCmd.Flags().BoolVarP(&_cloudtrailGetQueryResults, "get-query-results", "", false, "Get Query Results")
	_cloudtrailCmd.Flags().BoolVarP(&_cloudtrailGetResourcePolicy, "get-resource-policy", "", false, "Get Resource Policy")
	_cloudtrailCmd.Flags().BoolVarP(&_cloudtrailGetTrail, "get-trail", "", false, "Get Trail")
	_cloudtrailCmd.Flags().BoolVarP(&_cloudtrailGetTrailStatus, "get-trail-status", "", false, "Get Trail Status")
	_cloudtrailCmd.Flags().BoolVarP(&_cloudtrailListChannels, "list-channels", "", false, "List Channels")
	_cloudtrailCmd.Flags().BoolVarP(&_cloudtrailListDashboards, "list-dashboards", "", false, "List Dashboards")
	_cloudtrailCmd.Flags().BoolVarP(&_cloudtrailListEventDataStores, "list-event-data-stores", "", false, "List Event Data Stores")
	_cloudtrailCmd.Flags().BoolVarP(&_cloudtrailListImportFailures, "list-import-failures", "", false, "List Import Failures")
	_cloudtrailCmd.Flags().BoolVarP(&_cloudtrailListImports, "list-imports", "", false, "List Imports")
	_cloudtrailCmd.Flags().BoolVarP(&_cloudtrailListInsightsData, "list-insights-data", "", false, "List Insights Data")
	_cloudtrailCmd.Flags().BoolVarP(&_cloudtrailListInsightsMetricData, "list-insights-metric-data", "", false, "List Insights Metric Data")
	_cloudtrailCmd.Flags().BoolVarP(&_cloudtrailListPublicKeys, "list-public-keys", "", false, "List Public Keys")
	_cloudtrailCmd.Flags().BoolVarP(&_cloudtrailListQueries, "list-queries", "", false, "List Queries")
	_cloudtrailCmd.Flags().BoolVarP(&_cloudtrailListTags, "list-tags", "", false, "List Tags")
	_cloudtrailCmd.Flags().BoolVarP(&_cloudtrailListTrails, "list-trails", "", false, "List Trails")
	_cloudtrailCmd.Flags().BoolVarP(&_cloudtrailLookupEvents, "lookup-events", "", false, "Lookup Events")
	_cloudtrailCmd.Flags().BoolVarP(&_cloudtrailPutEventConfiguration, "put-event-configuration", "", false, "Put Event Configuration")
	_cloudtrailCmd.Flags().BoolVarP(&_cloudtrailPutEventSelectors, "put-event-selectors", "", false, "Put Event Selectors")
	_cloudtrailCmd.Flags().BoolVarP(&_cloudtrailPutInsightSelectors, "put-insight-selectors", "", false, "Put Insight Selectors")
	_cloudtrailCmd.Flags().BoolVarP(&_cloudtrailPutResourcePolicy, "put-resource-policy", "", false, "Put Resource Policy")
	_cloudtrailCmd.Flags().BoolVarP(&_cloudtrailRegisterOrganizationDelegatedAdmin, "register-organization-delegated-admin", "", false, "Register Organization Delegated Admin")
	_cloudtrailCmd.Flags().BoolVarP(&_cloudtrailRemoveTags, "remove-tags", "", false, "Remove Tags")
	_cloudtrailCmd.Flags().BoolVarP(&_cloudtrailRestoreEventDataStore, "restore-event-data-store", "", false, "Restore Event Data Store")
	_cloudtrailCmd.Flags().BoolVarP(&_cloudtrailSearchSampleQueries, "search-sample-queries", "", false, "Search Sample Queries")
	_cloudtrailCmd.Flags().BoolVarP(&_cloudtrailStartDashboardRefresh, "start-dashboard-refresh", "", false, "Start Dashboard Refresh")
	_cloudtrailCmd.Flags().BoolVarP(&_cloudtrailStartEventDataStoreIngestion, "start-event-data-store-ingestion", "", false, "Start Event Data Store Ingestion")
	_cloudtrailCmd.Flags().BoolVarP(&_cloudtrailStartImport, "start-import", "", false, "Start Import")
	_cloudtrailCmd.Flags().BoolVarP(&_cloudtrailStartLogging, "start-logging", "", false, "Start Logging")
	_cloudtrailCmd.Flags().BoolVarP(&_cloudtrailStartQuery, "start-query", "", false, "Start Query")
	_cloudtrailCmd.Flags().BoolVarP(&_cloudtrailStopEventDataStoreIngestion, "stop-event-data-store-ingestion", "", false, "Stop Event Data Store Ingestion")
	_cloudtrailCmd.Flags().BoolVarP(&_cloudtrailStopImport, "stop-import", "", false, "Stop Import")
	_cloudtrailCmd.Flags().BoolVarP(&_cloudtrailStopLogging, "stop-logging", "", false, "Stop Logging")
	_cloudtrailCmd.Flags().BoolVarP(&_cloudtrailUpdateChannel, "update-channel", "", false, "Update Channel")
	_cloudtrailCmd.Flags().BoolVarP(&_cloudtrailUpdateDashboard, "update-dashboard", "", false, "Update Dashboard")
	_cloudtrailCmd.Flags().BoolVarP(&_cloudtrailUpdateEventDataStore, "update-event-data-store", "", false, "Update Event Data Store")
	_cloudtrailCmd.Flags().BoolVarP(&_cloudtrailUpdateTrail, "update-trail", "", false, "Update Trail")

}
