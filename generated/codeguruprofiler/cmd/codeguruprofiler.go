package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codeguruprofiler"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// codeguruprofilerCmd represents the codeguruprofiler command
var _codeguruprofilerCmd = &cobra.Command{
	Use:   "codeguruprofiler",
	Short: "AWS codeguruprofiler CLI",
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
		client := codeguruprofiler.NewFromConfig(cfg)
		if _codeguruprofilerAddNotificationChannels {
			codeguruprofiler_AddNotificationChannels(cfg, client)
			return
		}
		if _codeguruprofilerBatchGetFrameMetricData {
			codeguruprofiler_BatchGetFrameMetricData(cfg, client)
			return
		}
		if _codeguruprofilerConfigureAgent {
			codeguruprofiler_ConfigureAgent(cfg, client)
			return
		}
		if _codeguruprofilerCreateProfilingGroup {
			codeguruprofiler_CreateProfilingGroup(cfg, client)
			return
		}
		if _codeguruprofilerDeleteProfilingGroup {
			codeguruprofiler_DeleteProfilingGroup(cfg, client)
			return
		}
		if _codeguruprofilerDescribeProfilingGroup {
			codeguruprofiler_DescribeProfilingGroup(cfg, client)
			return
		}
		if _codeguruprofilerGetFindingsReportAccountSummary {
			codeguruprofiler_GetFindingsReportAccountSummary(cfg, client)
			return
		}
		if _codeguruprofilerGetNotificationConfiguration {
			codeguruprofiler_GetNotificationConfiguration(cfg, client)
			return
		}
		if _codeguruprofilerGetPolicy {
			codeguruprofiler_GetPolicy(cfg, client)
			return
		}
		if _codeguruprofilerGetProfile {
			codeguruprofiler_GetProfile(cfg, client)
			return
		}
		if _codeguruprofilerGetRecommendations {
			codeguruprofiler_GetRecommendations(cfg, client)
			return
		}
		if _codeguruprofilerListFindingsReports {
			codeguruprofiler_ListFindingsReports(cfg, client)
			return
		}
		if _codeguruprofilerListProfileTimes {
			codeguruprofiler_ListProfileTimes(cfg, client)
			return
		}
		if _codeguruprofilerListProfilingGroups {
			codeguruprofiler_ListProfilingGroups(cfg, client)
			return
		}
		if _codeguruprofilerListTagsForResource {
			codeguruprofiler_ListTagsForResource(cfg, client)
			return
		}
		if _codeguruprofilerPostAgentProfile {
			codeguruprofiler_PostAgentProfile(cfg, client)
			return
		}
		if _codeguruprofilerPutPermission {
			codeguruprofiler_PutPermission(cfg, client)
			return
		}
		if _codeguruprofilerRemoveNotificationChannel {
			codeguruprofiler_RemoveNotificationChannel(cfg, client)
			return
		}
		if _codeguruprofilerRemovePermission {
			codeguruprofiler_RemovePermission(cfg, client)
			return
		}
		if _codeguruprofilerSubmitFeedback {
			codeguruprofiler_SubmitFeedback(cfg, client)
			return
		}
		if _codeguruprofilerTagResource {
			codeguruprofiler_TagResource(cfg, client)
			return
		}
		if _codeguruprofilerUntagResource {
			codeguruprofiler_UntagResource(cfg, client)
			return
		}
		if _codeguruprofilerUpdateProfilingGroup {
			codeguruprofiler_UpdateProfilingGroup(cfg, client)
			return
		}

	},
}

var (
	_codeguruprofilerAddNotificationChannels         bool
	_codeguruprofilerBatchGetFrameMetricData         bool
	_codeguruprofilerConfigureAgent                  bool
	_codeguruprofilerCreateProfilingGroup            bool
	_codeguruprofilerDeleteProfilingGroup            bool
	_codeguruprofilerDescribeProfilingGroup          bool
	_codeguruprofilerGetFindingsReportAccountSummary bool
	_codeguruprofilerGetNotificationConfiguration    bool
	_codeguruprofilerGetPolicy                       bool
	_codeguruprofilerGetProfile                      bool
	_codeguruprofilerGetRecommendations              bool
	_codeguruprofilerListFindingsReports             bool
	_codeguruprofilerListProfileTimes                bool
	_codeguruprofilerListProfilingGroups             bool
	_codeguruprofilerListTagsForResource             bool
	_codeguruprofilerPostAgentProfile                bool
	_codeguruprofilerPutPermission                   bool
	_codeguruprofilerRemoveNotificationChannel       bool
	_codeguruprofilerRemovePermission                bool
	_codeguruprofilerSubmitFeedback                  bool
	_codeguruprofilerTagResource                     bool
	_codeguruprofilerUntagResource                   bool
	_codeguruprofilerUpdateProfilingGroup            bool

	_codeguruprofilerAccept                   string
	_codeguruprofilerActionGroup              string
	_codeguruprofilerAgentOrchestrationConfig string
	_codeguruprofilerAgentProfile             string
	_codeguruprofilerAnomalyInstanceId        string
	_codeguruprofilerChannelId                string
	_codeguruprofilerChannels                 string
	_codeguruprofilerClientToken              string
	_codeguruprofilerComment                  string
	_codeguruprofilerComputePlatform          string
	_codeguruprofilerContentType              string
	_codeguruprofilerDailyReportsOnly         string
	_codeguruprofilerEndTime                  string
	_codeguruprofilerFleetInstanceId          string
	_codeguruprofilerFrameMetrics             string
	_codeguruprofilerIncludeDescription       string
	_codeguruprofilerLocale                   string
	_codeguruprofilerMaxDepth                 string
	_codeguruprofilerMaxResults               string
	_codeguruprofilerMetadata                 string
	_codeguruprofilerNextToken                string
	_codeguruprofilerOrderBy                  string
	_codeguruprofilerPeriod                   string
	_codeguruprofilerPrincipals               []string
	_codeguruprofilerProfileToken             string
	_codeguruprofilerProfilingGroupName       string
	_codeguruprofilerResourceArn              string
	_codeguruprofilerRevisionId               string
	_codeguruprofilerStartTime                string
	_codeguruprofilerTagKeys                  []string
	_codeguruprofilerTags                     string
	_codeguruprofilerTargetResolution         string
	_codeguruprofilerType                     string
)

// Add up to 2 anomaly notifications channels for a profiling group.
func codeguruprofiler_AddNotificationChannels(cfg aws.Config, client *codeguruprofiler.Client) {
	input := &codeguruprofiler.AddNotificationChannelsInput{
		// Channels: []types.Channel, // Required
		// ProfilingGroupName: *string, // Required
	}

	if len(_codeguruprofilerChannels) > 0 {
		if err := assignInputField(input, "Channels", _codeguruprofilerChannels); err != nil {
			log.Errorf("invalid --channels: %s", err.Error())
			return
		}
	}
	if len(_codeguruprofilerProfilingGroupName) > 0 {
		input.ProfilingGroupName = aws.String(_codeguruprofilerProfilingGroupName)
	}

	if resp, err := client.AddNotificationChannels(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the time series of values for a requested list of frame metrics from a
// time period.
func codeguruprofiler_BatchGetFrameMetricData(cfg aws.Config, client *codeguruprofiler.Client) {
	input := &codeguruprofiler.BatchGetFrameMetricDataInput{
		// ProfilingGroupName: *string, // Required
	}

	if len(_codeguruprofilerProfilingGroupName) > 0 {
		input.ProfilingGroupName = aws.String(_codeguruprofilerProfilingGroupName)
	}
	if len(_codeguruprofilerEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _codeguruprofilerEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_codeguruprofilerFrameMetrics) > 0 {
		if err := assignInputField(input, "FrameMetrics", _codeguruprofilerFrameMetrics); err != nil {
			log.Errorf("invalid --frame-metrics: %s", err.Error())
			return
		}
	}
	if len(_codeguruprofilerPeriod) > 0 {
		input.Period = aws.String(_codeguruprofilerPeriod)
	}
	if len(_codeguruprofilerStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _codeguruprofilerStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_codeguruprofilerTargetResolution) > 0 {
		if err := assignInputField(input, "TargetResolution", _codeguruprofilerTargetResolution); err != nil {
			log.Errorf("invalid --target-resolution: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchGetFrameMetricData(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Used by profiler agents to report their current state and to receive remote
// configuration updates. For example, ConfigureAgent can be used to tell an agent
// whether to profile or not and for how long to return profiling data.
func codeguruprofiler_ConfigureAgent(cfg aws.Config, client *codeguruprofiler.Client) {
	input := &codeguruprofiler.ConfigureAgentInput{
		// ProfilingGroupName: *string, // Required
	}

	if len(_codeguruprofilerProfilingGroupName) > 0 {
		input.ProfilingGroupName = aws.String(_codeguruprofilerProfilingGroupName)
	}
	if len(_codeguruprofilerFleetInstanceId) > 0 {
		input.FleetInstanceId = aws.String(_codeguruprofilerFleetInstanceId)
	}
	if len(_codeguruprofilerMetadata) > 0 {
		if err := assignInputField(input, "Metadata", _codeguruprofilerMetadata); err != nil {
			log.Errorf("invalid --metadata: %s", err.Error())
			return
		}
	}

	if resp, err := client.ConfigureAgent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a profiling group.
func codeguruprofiler_CreateProfilingGroup(cfg aws.Config, client *codeguruprofiler.Client) {
	input := &codeguruprofiler.CreateProfilingGroupInput{
		// ClientToken: *string, // Required
		// ProfilingGroupName: *string, // Required
	}

	if len(_codeguruprofilerClientToken) > 0 {
		input.ClientToken = aws.String(_codeguruprofilerClientToken)
	}
	if len(_codeguruprofilerProfilingGroupName) > 0 {
		input.ProfilingGroupName = aws.String(_codeguruprofilerProfilingGroupName)
	}
	if len(_codeguruprofilerAgentOrchestrationConfig) > 0 {
		if err := assignInputField(input, "AgentOrchestrationConfig", _codeguruprofilerAgentOrchestrationConfig); err != nil {
			log.Errorf("invalid --agent-orchestration-config: %s", err.Error())
			return
		}
	}
	if len(_codeguruprofilerComputePlatform) > 0 {
		if err := assignInputField(input, "ComputePlatform", _codeguruprofilerComputePlatform); err != nil {
			log.Errorf("invalid --compute-platform: %s", err.Error())
			return
		}
	}
	if len(_codeguruprofilerTags) > 0 {
		if err := assignInputField(input, "Tags", _codeguruprofilerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateProfilingGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a profiling group.
func codeguruprofiler_DeleteProfilingGroup(cfg aws.Config, client *codeguruprofiler.Client) {
	input := &codeguruprofiler.DeleteProfilingGroupInput{
		// ProfilingGroupName: *string, // Required
	}

	if len(_codeguruprofilerProfilingGroupName) > 0 {
		input.ProfilingGroupName = aws.String(_codeguruprofilerProfilingGroupName)
	}

	if resp, err := client.DeleteProfilingGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a [ProfilingGroupDescription]ProfilingGroupDescription object that contains information about the
// requested profiling group.
//
// [ProfilingGroupDescription]: https://docs.aws.amazon.com/codeguru/latest/profiler-api/API_ProfilingGroupDescription.html
func codeguruprofiler_DescribeProfilingGroup(cfg aws.Config, client *codeguruprofiler.Client) {
	input := &codeguruprofiler.DescribeProfilingGroupInput{
		// ProfilingGroupName: *string, // Required
	}

	if len(_codeguruprofilerProfilingGroupName) > 0 {
		input.ProfilingGroupName = aws.String(_codeguruprofilerProfilingGroupName)
	}

	if resp, err := client.DescribeProfilingGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of [FindingsReportSummary]FindingsReportSummary objects that contain analysis results
// for all profiling groups in your AWS account.
//
// [FindingsReportSummary]: https://docs.aws.amazon.com/codeguru/latest/profiler-api/API_FindingsReportSummary.html
func codeguruprofiler_GetFindingsReportAccountSummary(cfg aws.Config, client *codeguruprofiler.Client) {
	input := &codeguruprofiler.GetFindingsReportAccountSummaryInput{}

	if len(_codeguruprofilerDailyReportsOnly) > 0 {
		if err := assignInputField(input, "DailyReportsOnly", _codeguruprofilerDailyReportsOnly); err != nil {
			log.Errorf("invalid --daily-reports-only: %s", err.Error())
			return
		}
	}
	if len(_codeguruprofilerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codeguruprofilerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codeguruprofilerNextToken) > 0 {
		input.NextToken = aws.String(_codeguruprofilerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetFindingsReportAccountSummary(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codeguruprofiler.GetFindingsReportAccountSummaryOutput
	p := codeguruprofiler.NewGetFindingsReportAccountSummaryPaginator(client, input)
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

// Get the current configuration for anomaly notifications for a profiling group.
func codeguruprofiler_GetNotificationConfiguration(cfg aws.Config, client *codeguruprofiler.Client) {
	input := &codeguruprofiler.GetNotificationConfigurationInput{
		// ProfilingGroupName: *string, // Required
	}

	if len(_codeguruprofilerProfilingGroupName) > 0 {
		input.ProfilingGroupName = aws.String(_codeguruprofilerProfilingGroupName)
	}

	if resp, err := client.GetNotificationConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the JSON-formatted resource-based policy on a profiling group.
func codeguruprofiler_GetPolicy(cfg aws.Config, client *codeguruprofiler.Client) {
	input := &codeguruprofiler.GetPolicyInput{
		// ProfilingGroupName: *string, // Required
	}

	if len(_codeguruprofilerProfilingGroupName) > 0 {
		input.ProfilingGroupName = aws.String(_codeguruprofilerProfilingGroupName)
	}

	if resp, err := client.GetPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the aggregated profile of a profiling group for a specified time range.
// Amazon CodeGuru Profiler collects posted agent profiles for a profiling group
// into aggregated profiles.
//
// Because aggregated profiles expire over time GetProfile is not idempotent.
//
// Specify the time range for the requested aggregated profile using 1 or 2 of the
// following parameters: startTime , endTime , period . The maximum time range
// allowed is 7 days. If you specify all 3 parameters, an exception is thrown. If
// you specify only period , the latest aggregated profile is returned.
//
// Aggregated profiles are available with aggregation periods of 5 minutes, 1
// hour, and 1 day, aligned to UTC. The aggregation period of an aggregated profile
// determines how long it is retained. For more information, see [AggregatedProfileTime]
// AggregatedProfileTime . The aggregated profile's aggregation period determines
// how long
//
// it is retained by CodeGuru Profiler.
//
// - If the aggregation period is 5 minutes, the aggregated profile is retained
// for 15 days.
//
// - If the aggregation period is 1 hour, the aggregated profile is retained for
// 60 days.
//
// - If the aggregation period is 1 day, the aggregated profile is retained for
// 3 years.
//
// There are two use cases for calling GetProfile .
//
// - If you want to return an aggregated profile that already exists, use [ListProfileTimes]
// ListProfileTimes to view the time ranges of existing aggregated profiles. Use
// them in a GetProfile request to return a specific, existing aggregated
// profile.
//
// - If you want to return an aggregated profile for a time range that doesn't
// align with an existing aggregated profile, then CodeGuru Profiler makes a best
// effort to combine existing aggregated profiles from the requested time range and
// return them as one aggregated profile.
//
// # If aggregated profiles do not exist for the full time range requested, then
//
// aggregated profiles for a smaller time range are returned. For example, if the
// requested time range is from 00:00 to 00:20, and the existing aggregated
// profiles are from 00:15 and 00:25, then the aggregated profiles from 00:15 to
// 00:20 are returned.
//
// [ListProfileTimes]: https://docs.aws.amazon.com/codeguru/latest/profiler-api/API_ListProfileTimes.html
// [AggregatedProfileTime]: https://docs.aws.amazon.com/codeguru/latest/profiler-api/API_AggregatedProfileTime.html
func codeguruprofiler_GetProfile(cfg aws.Config, client *codeguruprofiler.Client) {
	input := &codeguruprofiler.GetProfileInput{
		// ProfilingGroupName: *string, // Required
	}

	if len(_codeguruprofilerProfilingGroupName) > 0 {
		input.ProfilingGroupName = aws.String(_codeguruprofilerProfilingGroupName)
	}
	if len(_codeguruprofilerAccept) > 0 {
		input.Accept = aws.String(_codeguruprofilerAccept)
	}
	if len(_codeguruprofilerEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _codeguruprofilerEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_codeguruprofilerMaxDepth) > 0 {
		if err := assignInputField(input, "MaxDepth", _codeguruprofilerMaxDepth); err != nil {
			log.Errorf("invalid --max-depth: %s", err.Error())
			return
		}
	}
	if len(_codeguruprofilerPeriod) > 0 {
		input.Period = aws.String(_codeguruprofilerPeriod)
	}
	if len(_codeguruprofilerStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _codeguruprofilerStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of [Recommendation]Recommendation objects that contain recommendations for a
// profiling group for a given time period. A list of [Anomaly]Anomaly objects that
// contains details about anomalies detected in the profiling group for the same
// time period is also returned.
//
// [Anomaly]: https://docs.aws.amazon.com/codeguru/latest/profiler-api/API_Anomaly.html
// [Recommendation]: https://docs.aws.amazon.com/codeguru/latest/profiler-api/API_Recommendation.html
func codeguruprofiler_GetRecommendations(cfg aws.Config, client *codeguruprofiler.Client) {
	input := &codeguruprofiler.GetRecommendationsInput{
		// EndTime: *time.Time, // Required
		// ProfilingGroupName: *string, // Required
		// StartTime: *time.Time, // Required
	}

	if len(_codeguruprofilerEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _codeguruprofilerEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_codeguruprofilerProfilingGroupName) > 0 {
		input.ProfilingGroupName = aws.String(_codeguruprofilerProfilingGroupName)
	}
	if len(_codeguruprofilerStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _codeguruprofilerStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_codeguruprofilerLocale) > 0 {
		input.Locale = aws.String(_codeguruprofilerLocale)
	}

	if resp, err := client.GetRecommendations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// List the available reports for a given profiling group and time range.
func codeguruprofiler_ListFindingsReports(cfg aws.Config, client *codeguruprofiler.Client) {
	input := &codeguruprofiler.ListFindingsReportsInput{
		// EndTime: *time.Time, // Required
		// ProfilingGroupName: *string, // Required
		// StartTime: *time.Time, // Required
	}

	if len(_codeguruprofilerEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _codeguruprofilerEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_codeguruprofilerProfilingGroupName) > 0 {
		input.ProfilingGroupName = aws.String(_codeguruprofilerProfilingGroupName)
	}
	if len(_codeguruprofilerStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _codeguruprofilerStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_codeguruprofilerDailyReportsOnly) > 0 {
		if err := assignInputField(input, "DailyReportsOnly", _codeguruprofilerDailyReportsOnly); err != nil {
			log.Errorf("invalid --daily-reports-only: %s", err.Error())
			return
		}
	}
	if len(_codeguruprofilerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codeguruprofilerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codeguruprofilerNextToken) > 0 {
		input.NextToken = aws.String(_codeguruprofilerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListFindingsReports(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codeguruprofiler.ListFindingsReportsOutput
	p := codeguruprofiler.NewListFindingsReportsPaginator(client, input)
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

// Lists the start times of the available aggregated profiles of a profiling group
// for an aggregation period within the specified time range.
func codeguruprofiler_ListProfileTimes(cfg aws.Config, client *codeguruprofiler.Client) {
	input := &codeguruprofiler.ListProfileTimesInput{
		// EndTime: *time.Time, // Required
		// Period: types.AggregationPeriod, // Required
		// ProfilingGroupName: *string, // Required
		// StartTime: *time.Time, // Required
	}

	if len(_codeguruprofilerEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _codeguruprofilerEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_codeguruprofilerPeriod) > 0 {
		if err := assignInputField(input, "Period", _codeguruprofilerPeriod); err != nil {
			log.Errorf("invalid --period: %s", err.Error())
			return
		}
	}
	if len(_codeguruprofilerProfilingGroupName) > 0 {
		input.ProfilingGroupName = aws.String(_codeguruprofilerProfilingGroupName)
	}
	if len(_codeguruprofilerStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _codeguruprofilerStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_codeguruprofilerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codeguruprofilerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codeguruprofilerNextToken) > 0 {
		input.NextToken = aws.String(_codeguruprofilerNextToken)
	}
	if len(_codeguruprofilerOrderBy) > 0 {
		if err := assignInputField(input, "OrderBy", _codeguruprofilerOrderBy); err != nil {
			log.Errorf("invalid --order-by: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListProfileTimes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codeguruprofiler.ListProfileTimesOutput
	p := codeguruprofiler.NewListProfileTimesPaginator(client, input)
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

// Returns a list of profiling groups. The profiling groups are returned as [ProfilingGroupDescription]
// ProfilingGroupDescription objects.
//
// [ProfilingGroupDescription]: https://docs.aws.amazon.com/codeguru/latest/profiler-api/API_ProfilingGroupDescription.html
func codeguruprofiler_ListProfilingGroups(cfg aws.Config, client *codeguruprofiler.Client) {
	input := &codeguruprofiler.ListProfilingGroupsInput{}

	if len(_codeguruprofilerIncludeDescription) > 0 {
		if err := assignInputField(input, "IncludeDescription", _codeguruprofilerIncludeDescription); err != nil {
			log.Errorf("invalid --include-description: %s", err.Error())
			return
		}
	}
	if len(_codeguruprofilerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codeguruprofilerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codeguruprofilerNextToken) > 0 {
		input.NextToken = aws.String(_codeguruprofilerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListProfilingGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codeguruprofiler.ListProfilingGroupsOutput
	p := codeguruprofiler.NewListProfilingGroupsPaginator(client, input)
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

// Returns a list of the tags that are assigned to a specified resource.
func codeguruprofiler_ListTagsForResource(cfg aws.Config, client *codeguruprofiler.Client) {
	input := &codeguruprofiler.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_codeguruprofilerResourceArn) > 0 {
		input.ResourceArn = aws.String(_codeguruprofilerResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Submits profiling data to an aggregated profile of a profiling group. To get
// an aggregated profile that is created with this profiling data, use [GetProfile]GetProfile .
//
// [GetProfile]: https://docs.aws.amazon.com/codeguru/latest/profiler-api/API_GetProfile.html
func codeguruprofiler_PostAgentProfile(cfg aws.Config, client *codeguruprofiler.Client) {
	input := &codeguruprofiler.PostAgentProfileInput{
		// AgentProfile: []byte, // Required
		// ContentType: *string, // Required
		// ProfilingGroupName: *string, // Required
	}

	if len(_codeguruprofilerAgentProfile) > 0 {
		if err := assignInputField(input, "AgentProfile", _codeguruprofilerAgentProfile); err != nil {
			log.Errorf("invalid --agent-profile: %s", err.Error())
			return
		}
	}
	if len(_codeguruprofilerContentType) > 0 {
		input.ContentType = aws.String(_codeguruprofilerContentType)
	}
	if len(_codeguruprofilerProfilingGroupName) > 0 {
		input.ProfilingGroupName = aws.String(_codeguruprofilerProfilingGroupName)
	}
	if len(_codeguruprofilerProfileToken) > 0 {
		input.ProfileToken = aws.String(_codeguruprofilerProfileToken)
	}

	if resp, err := client.PostAgentProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds permissions to a profiling group's resource-based policy that are
// provided using an action group. If a profiling group doesn't have a
// resource-based policy, one is created for it using the permissions in the action
// group and the roles and users in the principals parameter.
//
// The one supported action group that can be added is agentPermission which
// grants ConfigureAgent and PostAgent permissions. For more information, see [Resource-based policies in CodeGuru Profiler] in
// the Amazon CodeGuru Profiler User Guide, [ConfigureAgent]ConfigureAgent , and [PostAgentProfile]PostAgentProfile .
//
// The first time you call PutPermission on a profiling group, do not specify a
// revisionId because it doesn't have a resource-based policy. Subsequent calls
// must provide a revisionId to specify which revision of the resource-based
// policy to add the permissions to.
//
// The response contains the profiling group's JSON-formatted resource policy.
//
// [ConfigureAgent]: https://docs.aws.amazon.com/codeguru/latest/profiler-api/API_ConfigureAgent.html
// [Resource-based policies in CodeGuru Profiler]: https://docs.aws.amazon.com/codeguru/latest/profiler-ug/resource-based-policies.html
// [PostAgentProfile]: https://docs.aws.amazon.com/codeguru/latest/profiler-api/API_PostAgentProfile.html
func codeguruprofiler_PutPermission(cfg aws.Config, client *codeguruprofiler.Client) {
	input := &codeguruprofiler.PutPermissionInput{
		// ActionGroup: types.ActionGroup, // Required
		// Principals: []string, // Required
		// ProfilingGroupName: *string, // Required
	}

	if len(_codeguruprofilerActionGroup) > 0 {
		if err := assignInputField(input, "ActionGroup", _codeguruprofilerActionGroup); err != nil {
			log.Errorf("invalid --action-group: %s", err.Error())
			return
		}
	}
	if len(_codeguruprofilerPrincipals) > 0 {
		input.Principals = append([]string(nil), _codeguruprofilerPrincipals...)
	}
	if len(_codeguruprofilerProfilingGroupName) > 0 {
		input.ProfilingGroupName = aws.String(_codeguruprofilerProfilingGroupName)
	}
	if len(_codeguruprofilerRevisionId) > 0 {
		input.RevisionId = aws.String(_codeguruprofilerRevisionId)
	}

	if resp, err := client.PutPermission(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Remove one anomaly notifications channel for a profiling group.
func codeguruprofiler_RemoveNotificationChannel(cfg aws.Config, client *codeguruprofiler.Client) {
	input := &codeguruprofiler.RemoveNotificationChannelInput{
		// ChannelId: *string, // Required
		// ProfilingGroupName: *string, // Required
	}

	if len(_codeguruprofilerChannelId) > 0 {
		input.ChannelId = aws.String(_codeguruprofilerChannelId)
	}
	if len(_codeguruprofilerProfilingGroupName) > 0 {
		input.ProfilingGroupName = aws.String(_codeguruprofilerProfilingGroupName)
	}

	if resp, err := client.RemoveNotificationChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes permissions from a profiling group's resource-based policy that are
// provided using an action group. The one supported action group that can be
// removed is agentPermission which grants ConfigureAgent and PostAgent
// permissions. For more information, see [Resource-based policies in CodeGuru Profiler]in the Amazon CodeGuru Profiler User
// Guide, [ConfigureAgent]ConfigureAgent , and [PostAgentProfile]PostAgentProfile .
//
// [ConfigureAgent]: https://docs.aws.amazon.com/codeguru/latest/profiler-api/API_ConfigureAgent.html
// [Resource-based policies in CodeGuru Profiler]: https://docs.aws.amazon.com/codeguru/latest/profiler-ug/resource-based-policies.html
// [PostAgentProfile]: https://docs.aws.amazon.com/codeguru/latest/profiler-api/API_PostAgentProfile.html
func codeguruprofiler_RemovePermission(cfg aws.Config, client *codeguruprofiler.Client) {
	input := &codeguruprofiler.RemovePermissionInput{
		// ActionGroup: types.ActionGroup, // Required
		// ProfilingGroupName: *string, // Required
		// RevisionId: *string, // Required
	}

	if len(_codeguruprofilerActionGroup) > 0 {
		if err := assignInputField(input, "ActionGroup", _codeguruprofilerActionGroup); err != nil {
			log.Errorf("invalid --action-group: %s", err.Error())
			return
		}
	}
	if len(_codeguruprofilerProfilingGroupName) > 0 {
		input.ProfilingGroupName = aws.String(_codeguruprofilerProfilingGroupName)
	}
	if len(_codeguruprofilerRevisionId) > 0 {
		input.RevisionId = aws.String(_codeguruprofilerRevisionId)
	}

	if resp, err := client.RemovePermission(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sends feedback to CodeGuru Profiler about whether the anomaly detected by the
// analysis is useful or not.
func codeguruprofiler_SubmitFeedback(cfg aws.Config, client *codeguruprofiler.Client) {
	input := &codeguruprofiler.SubmitFeedbackInput{
		// AnomalyInstanceId: *string, // Required
		// ProfilingGroupName: *string, // Required
		// Type: types.FeedbackType, // Required
	}

	if len(_codeguruprofilerAnomalyInstanceId) > 0 {
		input.AnomalyInstanceId = aws.String(_codeguruprofilerAnomalyInstanceId)
	}
	if len(_codeguruprofilerProfilingGroupName) > 0 {
		input.ProfilingGroupName = aws.String(_codeguruprofilerProfilingGroupName)
	}
	if len(_codeguruprofilerType) > 0 {
		if err := assignInputField(input, "Type", _codeguruprofilerType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_codeguruprofilerComment) > 0 {
		input.Comment = aws.String(_codeguruprofilerComment)
	}

	if resp, err := client.SubmitFeedback(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Use to assign one or more tags to a resource.
func codeguruprofiler_TagResource(cfg aws.Config, client *codeguruprofiler.Client) {
	input := &codeguruprofiler.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_codeguruprofilerResourceArn) > 0 {
		input.ResourceArn = aws.String(_codeguruprofilerResourceArn)
	}
	if len(_codeguruprofilerTags) > 0 {
		if err := assignInputField(input, "Tags", _codeguruprofilerTags); err != nil {
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

// Use to remove one or more tags from a resource.
func codeguruprofiler_UntagResource(cfg aws.Config, client *codeguruprofiler.Client) {
	input := &codeguruprofiler.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_codeguruprofilerResourceArn) > 0 {
		input.ResourceArn = aws.String(_codeguruprofilerResourceArn)
	}
	if len(_codeguruprofilerTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _codeguruprofilerTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a profiling group.
func codeguruprofiler_UpdateProfilingGroup(cfg aws.Config, client *codeguruprofiler.Client) {
	input := &codeguruprofiler.UpdateProfilingGroupInput{
		// AgentOrchestrationConfig: *types.AgentOrchestrationConfig, // Required
		// ProfilingGroupName: *string, // Required
	}

	if len(_codeguruprofilerAgentOrchestrationConfig) > 0 {
		if err := assignInputField(input, "AgentOrchestrationConfig", _codeguruprofilerAgentOrchestrationConfig); err != nil {
			log.Errorf("invalid --agent-orchestration-config: %s", err.Error())
			return
		}
	}
	if len(_codeguruprofilerProfilingGroupName) > 0 {
		input.ProfilingGroupName = aws.String(_codeguruprofilerProfilingGroupName)
	}

	if resp, err := client.UpdateProfilingGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_codeguruprofilerCmd)
	_codeguruprofilerCmd.Flags().SortFlags = false

	_codeguruprofilerCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_codeguruprofilerCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_codeguruprofilerCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_codeguruprofilerCmd.Flags().StringVarP(&_codeguruprofilerAccept, "accept", "", "", "Accept")
	_codeguruprofilerCmd.Flags().StringVarP(&_codeguruprofilerActionGroup, "action-group", "", "", "Action Group")
	_codeguruprofilerCmd.Flags().StringVarP(&_codeguruprofilerAgentOrchestrationConfig, "agent-orchestration-config", "", "", "Agent Orchestration Config")
	_codeguruprofilerCmd.Flags().StringVarP(&_codeguruprofilerAgentProfile, "agent-profile", "", "", "Agent Profile")
	_codeguruprofilerCmd.Flags().StringVarP(&_codeguruprofilerAnomalyInstanceId, "anomaly-instance-id", "", "", "Anomaly Instance ID")
	_codeguruprofilerCmd.Flags().StringVarP(&_codeguruprofilerChannelId, "channel-id", "", "", "Channel ID")
	_codeguruprofilerCmd.Flags().StringVarP(&_codeguruprofilerChannels, "channels", "", "", "Channels")
	_codeguruprofilerCmd.Flags().StringVarP(&_codeguruprofilerClientToken, "client-token", "", "", "Client Token")
	_codeguruprofilerCmd.Flags().StringVarP(&_codeguruprofilerComment, "comment", "", "", "Comment")
	_codeguruprofilerCmd.Flags().StringVarP(&_codeguruprofilerComputePlatform, "compute-platform", "", "", "Compute Platform")
	_codeguruprofilerCmd.Flags().StringVarP(&_codeguruprofilerContentType, "content-type", "", "", "Content Type")
	_codeguruprofilerCmd.Flags().StringVarP(&_codeguruprofilerDailyReportsOnly, "daily-reports-only", "", "", "Daily Reports Only")
	_codeguruprofilerCmd.Flags().StringVarP(&_codeguruprofilerEndTime, "end-time", "", "", "End Time")
	_codeguruprofilerCmd.Flags().StringVarP(&_codeguruprofilerFleetInstanceId, "fleet-instance-id", "", "", "Fleet Instance ID")
	_codeguruprofilerCmd.Flags().StringVarP(&_codeguruprofilerFrameMetrics, "frame-metrics", "", "", "Frame Metrics")
	_codeguruprofilerCmd.Flags().StringVarP(&_codeguruprofilerIncludeDescription, "include-description", "", "", "Include Description")
	_codeguruprofilerCmd.Flags().StringVarP(&_codeguruprofilerLocale, "locale", "", "", "Locale")
	_codeguruprofilerCmd.Flags().StringVarP(&_codeguruprofilerMaxDepth, "max-depth", "", "", "Max Depth")
	_codeguruprofilerCmd.Flags().StringVarP(&_codeguruprofilerMaxResults, "max-results", "", "", "Max Results")
	_codeguruprofilerCmd.Flags().StringVarP(&_codeguruprofilerMetadata, "metadata", "", "", "Metadata")
	_codeguruprofilerCmd.Flags().StringVarP(&_codeguruprofilerNextToken, "next-token", "", "", "Next Token")
	_codeguruprofilerCmd.Flags().StringVarP(&_codeguruprofilerOrderBy, "order-by", "", "", "Order By")
	_codeguruprofilerCmd.Flags().StringVarP(&_codeguruprofilerPeriod, "period", "", "", "Period")
	_codeguruprofilerCmd.Flags().StringSliceVarP(&_codeguruprofilerPrincipals, "principals", "", nil, "Principals")
	_codeguruprofilerCmd.Flags().StringVarP(&_codeguruprofilerProfileToken, "profile-token", "", "", "Profile Token")
	_codeguruprofilerCmd.Flags().StringVarP(&_codeguruprofilerProfilingGroupName, "profiling-group-name", "", "", "Profiling Group Name")
	_codeguruprofilerCmd.Flags().StringVarP(&_codeguruprofilerResourceArn, "resource-arn", "", "", "Resource ARN")
	_codeguruprofilerCmd.Flags().StringVarP(&_codeguruprofilerRevisionId, "revision-id", "", "", "Revision ID")
	_codeguruprofilerCmd.Flags().StringVarP(&_codeguruprofilerStartTime, "start-time", "", "", "Start Time")
	_codeguruprofilerCmd.Flags().StringSliceVarP(&_codeguruprofilerTagKeys, "tag-keys", "", nil, "Tag Keys")
	_codeguruprofilerCmd.Flags().StringVarP(&_codeguruprofilerTags, "tags", "", "", "Tags")
	_codeguruprofilerCmd.Flags().StringVarP(&_codeguruprofilerTargetResolution, "target-resolution", "", "", "Target Resolution")
	_codeguruprofilerCmd.Flags().StringVarP(&_codeguruprofilerType, "type", "", "", "Type")

	_codeguruprofilerCmd.Flags().BoolVarP(&_codeguruprofilerAddNotificationChannels, "add-notification-channels", "", false, "Add Notification Channels")
	_codeguruprofilerCmd.Flags().BoolVarP(&_codeguruprofilerBatchGetFrameMetricData, "batch-get-frame-metric-data", "", false, "Batch Get Frame Metric Data")
	_codeguruprofilerCmd.Flags().BoolVarP(&_codeguruprofilerConfigureAgent, "configure-agent", "", false, "Configure Agent")
	_codeguruprofilerCmd.Flags().BoolVarP(&_codeguruprofilerCreateProfilingGroup, "create-profiling-group", "", false, "Create Profiling Group")
	_codeguruprofilerCmd.Flags().BoolVarP(&_codeguruprofilerDeleteProfilingGroup, "delete-profiling-group", "", false, "Delete Profiling Group")
	_codeguruprofilerCmd.Flags().BoolVarP(&_codeguruprofilerDescribeProfilingGroup, "describe-profiling-group", "", false, "Describe Profiling Group")
	_codeguruprofilerCmd.Flags().BoolVarP(&_codeguruprofilerGetFindingsReportAccountSummary, "get-findings-report-account-summary", "", false, "Get Findings Report Account Summary")
	_codeguruprofilerCmd.Flags().BoolVarP(&_codeguruprofilerGetNotificationConfiguration, "get-notification-configuration", "", false, "Get Notification Configuration")
	_codeguruprofilerCmd.Flags().BoolVarP(&_codeguruprofilerGetPolicy, "get-policy", "", false, "Get Policy")
	_codeguruprofilerCmd.Flags().BoolVarP(&_codeguruprofilerGetProfile, "get-profile", "", false, "Get Profile")
	_codeguruprofilerCmd.Flags().BoolVarP(&_codeguruprofilerGetRecommendations, "get-recommendations", "", false, "Get Recommendations")
	_codeguruprofilerCmd.Flags().BoolVarP(&_codeguruprofilerListFindingsReports, "list-findings-reports", "", false, "List Findings Reports")
	_codeguruprofilerCmd.Flags().BoolVarP(&_codeguruprofilerListProfileTimes, "list-profile-times", "", false, "List Profile Times")
	_codeguruprofilerCmd.Flags().BoolVarP(&_codeguruprofilerListProfilingGroups, "list-profiling-groups", "", false, "List Profiling Groups")
	_codeguruprofilerCmd.Flags().BoolVarP(&_codeguruprofilerListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_codeguruprofilerCmd.Flags().BoolVarP(&_codeguruprofilerPostAgentProfile, "post-agent-profile", "", false, "Post Agent Profile")
	_codeguruprofilerCmd.Flags().BoolVarP(&_codeguruprofilerPutPermission, "put-permission", "", false, "Put Permission")
	_codeguruprofilerCmd.Flags().BoolVarP(&_codeguruprofilerRemoveNotificationChannel, "remove-notification-channel", "", false, "Remove Notification Channel")
	_codeguruprofilerCmd.Flags().BoolVarP(&_codeguruprofilerRemovePermission, "remove-permission", "", false, "Remove Permission")
	_codeguruprofilerCmd.Flags().BoolVarP(&_codeguruprofilerSubmitFeedback, "submit-feedback", "", false, "Submit Feedback")
	_codeguruprofilerCmd.Flags().BoolVarP(&_codeguruprofilerTagResource, "tag-resource", "", false, "Tag Resource")
	_codeguruprofilerCmd.Flags().BoolVarP(&_codeguruprofilerUntagResource, "untag-resource", "", false, "Untag Resource")
	_codeguruprofilerCmd.Flags().BoolVarP(&_codeguruprofilerUpdateProfilingGroup, "update-profiling-group", "", false, "Update Profiling Group")

}
