package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// cloudwatchCmd represents the cloudwatch command
var _cloudwatchCmd = &cobra.Command{
	Use:   "cloudwatch",
	Short: "AWS cloudwatch CLI",
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
		client := cloudwatch.NewFromConfig(cfg)
		if _cloudwatchDeleteAlarmMuteRule {
			cloudwatch_DeleteAlarmMuteRule(cfg, client)
			return
		}
		if _cloudwatchDeleteAlarms {
			cloudwatch_DeleteAlarms(cfg, client)
			return
		}
		if _cloudwatchDeleteAnomalyDetector {
			cloudwatch_DeleteAnomalyDetector(cfg, client)
			return
		}
		if _cloudwatchDeleteDashboards {
			cloudwatch_DeleteDashboards(cfg, client)
			return
		}
		if _cloudwatchDeleteInsightRules {
			cloudwatch_DeleteInsightRules(cfg, client)
			return
		}
		if _cloudwatchDeleteMetricStream {
			cloudwatch_DeleteMetricStream(cfg, client)
			return
		}
		if _cloudwatchDescribeAlarmContributors {
			cloudwatch_DescribeAlarmContributors(cfg, client)
			return
		}
		if _cloudwatchDescribeAlarmHistory {
			cloudwatch_DescribeAlarmHistory(cfg, client)
			return
		}
		if _cloudwatchDescribeAlarms {
			cloudwatch_DescribeAlarms(cfg, client)
			return
		}
		if _cloudwatchDescribeAlarmsForMetric {
			cloudwatch_DescribeAlarmsForMetric(cfg, client)
			return
		}
		if _cloudwatchDescribeAnomalyDetectors {
			cloudwatch_DescribeAnomalyDetectors(cfg, client)
			return
		}
		if _cloudwatchDescribeInsightRules {
			cloudwatch_DescribeInsightRules(cfg, client)
			return
		}
		if _cloudwatchDisableAlarmActions {
			cloudwatch_DisableAlarmActions(cfg, client)
			return
		}
		if _cloudwatchDisableInsightRules {
			cloudwatch_DisableInsightRules(cfg, client)
			return
		}
		if _cloudwatchEnableAlarmActions {
			cloudwatch_EnableAlarmActions(cfg, client)
			return
		}
		if _cloudwatchEnableInsightRules {
			cloudwatch_EnableInsightRules(cfg, client)
			return
		}
		if _cloudwatchGetAlarmMuteRule {
			cloudwatch_GetAlarmMuteRule(cfg, client)
			return
		}
		if _cloudwatchGetDashboard {
			cloudwatch_GetDashboard(cfg, client)
			return
		}
		if _cloudwatchGetInsightRuleReport {
			cloudwatch_GetInsightRuleReport(cfg, client)
			return
		}
		if _cloudwatchGetMetricData {
			cloudwatch_GetMetricData(cfg, client)
			return
		}
		if _cloudwatchGetMetricStatistics {
			cloudwatch_GetMetricStatistics(cfg, client)
			return
		}
		if _cloudwatchGetMetricStream {
			cloudwatch_GetMetricStream(cfg, client)
			return
		}
		if _cloudwatchGetMetricWidgetImage {
			cloudwatch_GetMetricWidgetImage(cfg, client)
			return
		}
		if _cloudwatchListAlarmMuteRules {
			cloudwatch_ListAlarmMuteRules(cfg, client)
			return
		}
		if _cloudwatchListDashboards {
			cloudwatch_ListDashboards(cfg, client)
			return
		}
		if _cloudwatchListManagedInsightRules {
			cloudwatch_ListManagedInsightRules(cfg, client)
			return
		}
		if _cloudwatchListMetricStreams {
			cloudwatch_ListMetricStreams(cfg, client)
			return
		}
		if _cloudwatchListMetrics {
			cloudwatch_ListMetrics(cfg, client)
			return
		}
		if _cloudwatchListTagsForResource {
			cloudwatch_ListTagsForResource(cfg, client)
			return
		}
		if _cloudwatchPutAlarmMuteRule {
			cloudwatch_PutAlarmMuteRule(cfg, client)
			return
		}
		if _cloudwatchPutAnomalyDetector {
			cloudwatch_PutAnomalyDetector(cfg, client)
			return
		}
		if _cloudwatchPutCompositeAlarm {
			cloudwatch_PutCompositeAlarm(cfg, client)
			return
		}
		if _cloudwatchPutDashboard {
			cloudwatch_PutDashboard(cfg, client)
			return
		}
		if _cloudwatchPutInsightRule {
			cloudwatch_PutInsightRule(cfg, client)
			return
		}
		if _cloudwatchPutManagedInsightRules {
			cloudwatch_PutManagedInsightRules(cfg, client)
			return
		}
		if _cloudwatchPutMetricAlarm {
			cloudwatch_PutMetricAlarm(cfg, client)
			return
		}
		if _cloudwatchPutMetricData {
			cloudwatch_PutMetricData(cfg, client)
			return
		}
		if _cloudwatchPutMetricStream {
			cloudwatch_PutMetricStream(cfg, client)
			return
		}
		if _cloudwatchSetAlarmState {
			cloudwatch_SetAlarmState(cfg, client)
			return
		}
		if _cloudwatchStartMetricStreams {
			cloudwatch_StartMetricStreams(cfg, client)
			return
		}
		if _cloudwatchStopMetricStreams {
			cloudwatch_StopMetricStreams(cfg, client)
			return
		}
		if _cloudwatchTagResource {
			cloudwatch_TagResource(cfg, client)
			return
		}
		if _cloudwatchUntagResource {
			cloudwatch_UntagResource(cfg, client)
			return
		}

	},
}

var (
	_cloudwatchDeleteAlarmMuteRule       bool
	_cloudwatchDeleteAlarms              bool
	_cloudwatchDeleteAnomalyDetector     bool
	_cloudwatchDeleteDashboards          bool
	_cloudwatchDeleteInsightRules        bool
	_cloudwatchDeleteMetricStream        bool
	_cloudwatchDescribeAlarmContributors bool
	_cloudwatchDescribeAlarmHistory      bool
	_cloudwatchDescribeAlarms            bool
	_cloudwatchDescribeAlarmsForMetric   bool
	_cloudwatchDescribeAnomalyDetectors  bool
	_cloudwatchDescribeInsightRules      bool
	_cloudwatchDisableAlarmActions       bool
	_cloudwatchDisableInsightRules       bool
	_cloudwatchEnableAlarmActions        bool
	_cloudwatchEnableInsightRules        bool
	_cloudwatchGetAlarmMuteRule          bool
	_cloudwatchGetDashboard              bool
	_cloudwatchGetInsightRuleReport      bool
	_cloudwatchGetMetricData             bool
	_cloudwatchGetMetricStatistics       bool
	_cloudwatchGetMetricStream           bool
	_cloudwatchGetMetricWidgetImage      bool
	_cloudwatchListAlarmMuteRules        bool
	_cloudwatchListDashboards            bool
	_cloudwatchListManagedInsightRules   bool
	_cloudwatchListMetricStreams         bool
	_cloudwatchListMetrics               bool
	_cloudwatchListTagsForResource       bool
	_cloudwatchPutAlarmMuteRule          bool
	_cloudwatchPutAnomalyDetector        bool
	_cloudwatchPutCompositeAlarm         bool
	_cloudwatchPutDashboard              bool
	_cloudwatchPutInsightRule            bool
	_cloudwatchPutManagedInsightRules    bool
	_cloudwatchPutMetricAlarm            bool
	_cloudwatchPutMetricData             bool
	_cloudwatchPutMetricStream           bool
	_cloudwatchSetAlarmState             bool
	_cloudwatchStartMetricStreams        bool
	_cloudwatchStopMetricStreams         bool
	_cloudwatchTagResource               bool
	_cloudwatchUntagResource             bool

	_cloudwatchActionPrefix                     string
	_cloudwatchActionsEnabled                   string
	_cloudwatchActionsSuppressor                string
	_cloudwatchActionsSuppressorExtensionPeriod string
	_cloudwatchActionsSuppressorWaitPeriod      string
	_cloudwatchAlarmActions                     []string
	_cloudwatchAlarmContributorId               string
	_cloudwatchAlarmDescription                 string
	_cloudwatchAlarmMuteRuleName                string
	_cloudwatchAlarmName                        string
	_cloudwatchAlarmNamePrefix                  string
	_cloudwatchAlarmNames                       []string
	_cloudwatchAlarmRule                        string
	_cloudwatchAlarmTypes                       string
	_cloudwatchAnomalyDetectorTypes             string
	_cloudwatchApplyOnTransformedLogs           string
	_cloudwatchChildrenOfAlarmName              string
	_cloudwatchComparisonOperator               string
	_cloudwatchConfiguration                    string
	_cloudwatchDashboardBody                    string
	_cloudwatchDashboardName                    string
	_cloudwatchDashboardNamePrefix              string
	_cloudwatchDashboardNames                   []string
	_cloudwatchDatapointsToAlarm                string
	_cloudwatchDescription                      string
	_cloudwatchDimensions                       string
	_cloudwatchEndDate                          string
	_cloudwatchEndTime                          string
	_cloudwatchEntityMetricData                 string
	_cloudwatchEvaluateLowSampleCountPercentile string
	_cloudwatchEvaluationPeriods                string
	_cloudwatchExcludeFilters                   string
	_cloudwatchExpireDate                       string
	_cloudwatchExtendedStatistic                string
	_cloudwatchExtendedStatistics               []string
	_cloudwatchFirehoseArn                      string
	_cloudwatchHistoryItemType                  string
	_cloudwatchIncludeFilters                   string
	_cloudwatchIncludeLinkedAccounts            string
	_cloudwatchIncludeLinkedAccountsMetrics     string
	_cloudwatchInsufficientDataActions          []string
	_cloudwatchLabelOptions                     string
	_cloudwatchManagedRules                     string
	_cloudwatchMaxContributorCount              string
	_cloudwatchMaxDatapoints                    string
	_cloudwatchMaxRecords                       string
	_cloudwatchMaxResults                       string
	_cloudwatchMetricCharacteristics            string
	_cloudwatchMetricData                       string
	_cloudwatchMetricDataQueries                string
	_cloudwatchMetricMathAnomalyDetector        string
	_cloudwatchMetricName                       string
	_cloudwatchMetricWidget                     string
	_cloudwatchMetrics                          string
	_cloudwatchMuteTargets                      string
	_cloudwatchName                             string
	_cloudwatchNames                            []string
	_cloudwatchNamespace                        string
	_cloudwatchNextToken                        string
	_cloudwatchOKActions                        []string
	_cloudwatchOrderBy                          string
	_cloudwatchOutputFormat                     string
	_cloudwatchOwningAccount                    string
	_cloudwatchParentsOfAlarmName               string
	_cloudwatchPeriod                           string
	_cloudwatchRecentlyActive                   string
	_cloudwatchResourceARN                      string
	_cloudwatchRoleArn                          string
	_cloudwatchRule                             string
	_cloudwatchRuleDefinition                   string
	_cloudwatchRuleName                         string
	_cloudwatchRuleNames                        []string
	_cloudwatchRuleState                        string
	_cloudwatchScanBy                           string
	_cloudwatchSingleMetricAnomalyDetector      string
	_cloudwatchStartDate                        string
	_cloudwatchStartTime                        string
	_cloudwatchStat                             string
	_cloudwatchStateReason                      string
	_cloudwatchStateReasonData                  string
	_cloudwatchStateValue                       string
	_cloudwatchStatistic                        string
	_cloudwatchStatistics                       string
	_cloudwatchStatisticsConfigurations         string
	_cloudwatchStatuses                         string
	_cloudwatchStrictEntityValidation           string
	_cloudwatchTagKeys                          []string
	_cloudwatchTags                             string
	_cloudwatchThreshold                        string
	_cloudwatchThresholdMetricId                string
	_cloudwatchTreatMissingData                 string
	_cloudwatchUnit                             string
)

// Deletes a specific alarm mute rule.
// When you delete a mute rule, any alarms that are currently being muted by that
// rule are immediately unmuted. If those alarms are in an ALARM state, their
// configured actions will trigger.
//
// This operation is idempotent. If you delete a mute rule that does not exist,
// the operation succeeds without returning an error.
//
// # Permissions
//
// To delete a mute rule, you need the cloudwatch:DeleteAlarmMuteRule permission
// on the alarm mute rule resource.
func cloudwatch_DeleteAlarmMuteRule(cfg aws.Config, client *cloudwatch.Client) {
	input := &cloudwatch.DeleteAlarmMuteRuleInput{
		// AlarmMuteRuleName: *string, // Required
	}

	if len(_cloudwatchAlarmMuteRuleName) > 0 {
		input.AlarmMuteRuleName = aws.String(_cloudwatchAlarmMuteRuleName)
	}

	if resp, err := client.DeleteAlarmMuteRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified alarms. You can delete up to 100 alarms in one operation.
// However, this total can include no more than one composite alarm. For example,
// you could delete 99 metric alarms and one composite alarms with one operation,
// but you can't delete two composite alarms with one operation.
//
// If you specify any incorrect alarm names, the alarms you specify with correct
// names are still deleted. Other syntax errors might result in no alarms being
// deleted. To confirm that alarms were deleted successfully, you can use the [DescribeAlarms]
// operation after using DeleteAlarms .
//
// It is possible to create a loop or cycle of composite alarms, where composite
// alarm A depends on composite alarm B, and composite alarm B also depends on
// composite alarm A. In this scenario, you can't delete any composite alarm that
// is part of the cycle because there is always still a composite alarm that
// depends on that alarm that you want to delete.
//
// To get out of such a situation, you must break the cycle by changing the rule
// of one of the composite alarms in the cycle to remove a dependency that creates
// the cycle. The simplest change to make to break a cycle is to change the
// AlarmRule of one of the alarms to false .
//
// Additionally, the evaluation of composite alarms stops if CloudWatch detects a
// cycle in the evaluation path.
//
// [DescribeAlarms]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_DescribeAlarms.html
func cloudwatch_DeleteAlarms(cfg aws.Config, client *cloudwatch.Client) {
	input := &cloudwatch.DeleteAlarmsInput{
		// AlarmNames: []string, // Required
	}

	if len(_cloudwatchAlarmNames) > 0 {
		input.AlarmNames = append([]string(nil), _cloudwatchAlarmNames...)
	}

	if resp, err := client.DeleteAlarms(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified anomaly detection model from your account. For more
// information about how to delete an anomaly detection model, see [Deleting an anomaly detection model]in the
// CloudWatch User Guide.
//
// [Deleting an anomaly detection model]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Create_Anomaly_Detection_Alarm.html#Delete_Anomaly_Detection_Model
func cloudwatch_DeleteAnomalyDetector(cfg aws.Config, client *cloudwatch.Client) {
	input := &cloudwatch.DeleteAnomalyDetectorInput{}

	if len(_cloudwatchDimensions) > 0 {
		if err := assignInputField(input, "Dimensions", _cloudwatchDimensions); err != nil {
			log.Errorf("invalid --dimensions: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchMetricMathAnomalyDetector) > 0 {
		if err := assignInputField(input, "MetricMathAnomalyDetector", _cloudwatchMetricMathAnomalyDetector); err != nil {
			log.Errorf("invalid --metric-math-anomaly-detector: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchMetricName) > 0 {
		input.MetricName = aws.String(_cloudwatchMetricName)
	}
	if len(_cloudwatchNamespace) > 0 {
		input.Namespace = aws.String(_cloudwatchNamespace)
	}
	if len(_cloudwatchSingleMetricAnomalyDetector) > 0 {
		if err := assignInputField(input, "SingleMetricAnomalyDetector", _cloudwatchSingleMetricAnomalyDetector); err != nil {
			log.Errorf("invalid --single-metric-anomaly-detector: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchStat) > 0 {
		input.Stat = aws.String(_cloudwatchStat)
	}

	if resp, err := client.DeleteAnomalyDetector(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes all dashboards that you specify. You can specify up to 100 dashboards
// to delete. If there is an error during this call, no dashboards are deleted.
func cloudwatch_DeleteDashboards(cfg aws.Config, client *cloudwatch.Client) {
	input := &cloudwatch.DeleteDashboardsInput{
		// DashboardNames: []string, // Required
	}

	if len(_cloudwatchDashboardNames) > 0 {
		input.DashboardNames = append([]string(nil), _cloudwatchDashboardNames...)
	}

	if resp, err := client.DeleteDashboards(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Permanently deletes the specified Contributor Insights rules.
// If you create a rule, delete it, and then re-create it with the same name,
// historical data from the first time the rule was created might not be available.
func cloudwatch_DeleteInsightRules(cfg aws.Config, client *cloudwatch.Client) {
	input := &cloudwatch.DeleteInsightRulesInput{
		// RuleNames: []string, // Required
	}

	if len(_cloudwatchRuleNames) > 0 {
		input.RuleNames = append([]string(nil), _cloudwatchRuleNames...)
	}

	if resp, err := client.DeleteInsightRules(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Permanently deletes the metric stream that you specify.
func cloudwatch_DeleteMetricStream(cfg aws.Config, client *cloudwatch.Client) {
	input := &cloudwatch.DeleteMetricStreamInput{
		// Name: *string, // Required
	}

	if len(_cloudwatchName) > 0 {
		input.Name = aws.String(_cloudwatchName)
	}

	if resp, err := client.DeleteMetricStream(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the information of the current alarm contributors that are in ALARM
// state. This operation returns details about the individual time series that
// contribute to the alarm's state.
func cloudwatch_DescribeAlarmContributors(cfg aws.Config, client *cloudwatch.Client) {
	input := &cloudwatch.DescribeAlarmContributorsInput{
		// AlarmName: *string, // Required
	}

	if len(_cloudwatchAlarmName) > 0 {
		input.AlarmName = aws.String(_cloudwatchAlarmName)
	}
	if len(_cloudwatchNextToken) > 0 {
		input.NextToken = aws.String(_cloudwatchNextToken)
	}

	if resp, err := client.DescribeAlarmContributors(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the history for the specified alarm. You can filter the results by
// date range or item type. If an alarm name is not specified, the histories for
// either all metric alarms or all composite alarms are returned.
//
// CloudWatch retains the history of an alarm even if you delete the alarm.
//
// To use this operation and return information about a composite alarm, you must
// be signed on with the cloudwatch:DescribeAlarmHistory permission that is scoped
// to * . You can't return information about composite alarms if your
// cloudwatch:DescribeAlarmHistory permission has a narrower scope.
func cloudwatch_DescribeAlarmHistory(cfg aws.Config, client *cloudwatch.Client) {
	input := &cloudwatch.DescribeAlarmHistoryInput{}

	if len(_cloudwatchAlarmContributorId) > 0 {
		input.AlarmContributorId = aws.String(_cloudwatchAlarmContributorId)
	}
	if len(_cloudwatchAlarmName) > 0 {
		input.AlarmName = aws.String(_cloudwatchAlarmName)
	}
	if len(_cloudwatchAlarmTypes) > 0 {
		if err := assignInputField(input, "AlarmTypes", _cloudwatchAlarmTypes); err != nil {
			log.Errorf("invalid --alarm-types: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchEndDate) > 0 {
		if err := assignInputField(input, "EndDate", _cloudwatchEndDate); err != nil {
			log.Errorf("invalid --end-date: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchHistoryItemType) > 0 {
		if err := assignInputField(input, "HistoryItemType", _cloudwatchHistoryItemType); err != nil {
			log.Errorf("invalid --history-item-type: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _cloudwatchMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchNextToken) > 0 {
		input.NextToken = aws.String(_cloudwatchNextToken)
	}
	if len(_cloudwatchScanBy) > 0 {
		if err := assignInputField(input, "ScanBy", _cloudwatchScanBy); err != nil {
			log.Errorf("invalid --scan-by: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchStartDate) > 0 {
		if err := assignInputField(input, "StartDate", _cloudwatchStartDate); err != nil {
			log.Errorf("invalid --start-date: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeAlarmHistory(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudwatch.DescribeAlarmHistoryOutput
	p := cloudwatch.NewDescribeAlarmHistoryPaginator(client, input)
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

// Retrieves the specified alarms. You can filter the results by specifying a
// prefix for the alarm name, the alarm state, or a prefix for any action.
//
// To use this operation and return information about composite alarms, you must
// be signed on with the cloudwatch:DescribeAlarms permission that is scoped to * .
// You can't return information about composite alarms if your
// cloudwatch:DescribeAlarms permission has a narrower scope.
func cloudwatch_DescribeAlarms(cfg aws.Config, client *cloudwatch.Client) {
	input := &cloudwatch.DescribeAlarmsInput{}

	if len(_cloudwatchActionPrefix) > 0 {
		input.ActionPrefix = aws.String(_cloudwatchActionPrefix)
	}
	if len(_cloudwatchAlarmNamePrefix) > 0 {
		input.AlarmNamePrefix = aws.String(_cloudwatchAlarmNamePrefix)
	}
	if len(_cloudwatchAlarmNames) > 0 {
		input.AlarmNames = append([]string(nil), _cloudwatchAlarmNames...)
	}
	if len(_cloudwatchAlarmTypes) > 0 {
		if err := assignInputField(input, "AlarmTypes", _cloudwatchAlarmTypes); err != nil {
			log.Errorf("invalid --alarm-types: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchChildrenOfAlarmName) > 0 {
		input.ChildrenOfAlarmName = aws.String(_cloudwatchChildrenOfAlarmName)
	}
	if len(_cloudwatchMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _cloudwatchMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchNextToken) > 0 {
		input.NextToken = aws.String(_cloudwatchNextToken)
	}
	if len(_cloudwatchParentsOfAlarmName) > 0 {
		input.ParentsOfAlarmName = aws.String(_cloudwatchParentsOfAlarmName)
	}
	if len(_cloudwatchStateValue) > 0 {
		if err := assignInputField(input, "StateValue", _cloudwatchStateValue); err != nil {
			log.Errorf("invalid --state-value: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeAlarms(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudwatch.DescribeAlarmsOutput
	p := cloudwatch.NewDescribeAlarmsPaginator(client, input)
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

// Retrieves the alarms for the specified metric. To filter the results, specify a
// statistic, period, or unit.
//
// This operation retrieves only standard alarms that are based on the specified
// metric. It does not return alarms based on math expressions that use the
// specified metric, or composite alarms that use the specified metric.
func cloudwatch_DescribeAlarmsForMetric(cfg aws.Config, client *cloudwatch.Client) {
	input := &cloudwatch.DescribeAlarmsForMetricInput{
		// MetricName: *string, // Required
		// Namespace: *string, // Required
	}

	if len(_cloudwatchMetricName) > 0 {
		input.MetricName = aws.String(_cloudwatchMetricName)
	}
	if len(_cloudwatchNamespace) > 0 {
		input.Namespace = aws.String(_cloudwatchNamespace)
	}
	if len(_cloudwatchDimensions) > 0 {
		if err := assignInputField(input, "Dimensions", _cloudwatchDimensions); err != nil {
			log.Errorf("invalid --dimensions: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchExtendedStatistic) > 0 {
		input.ExtendedStatistic = aws.String(_cloudwatchExtendedStatistic)
	}
	if len(_cloudwatchPeriod) > 0 {
		if err := assignInputField(input, "Period", _cloudwatchPeriod); err != nil {
			log.Errorf("invalid --period: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchStatistic) > 0 {
		if err := assignInputField(input, "Statistic", _cloudwatchStatistic); err != nil {
			log.Errorf("invalid --statistic: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchUnit) > 0 {
		if err := assignInputField(input, "Unit", _cloudwatchUnit); err != nil {
			log.Errorf("invalid --unit: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeAlarmsForMetric(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the anomaly detection models that you have created in your account. For
// single metric anomaly detectors, you can list all of the models in your account
// or filter the results to only the models that are related to a certain
// namespace, metric name, or metric dimension. For metric math anomaly detectors,
// you can list them by adding METRIC_MATH to the AnomalyDetectorTypes array. This
// will return all metric math anomaly detectors in your account.
func cloudwatch_DescribeAnomalyDetectors(cfg aws.Config, client *cloudwatch.Client) {
	input := &cloudwatch.DescribeAnomalyDetectorsInput{}

	if len(_cloudwatchAnomalyDetectorTypes) > 0 {
		if err := assignInputField(input, "AnomalyDetectorTypes", _cloudwatchAnomalyDetectorTypes); err != nil {
			log.Errorf("invalid --anomaly-detector-types: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchDimensions) > 0 {
		if err := assignInputField(input, "Dimensions", _cloudwatchDimensions); err != nil {
			log.Errorf("invalid --dimensions: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cloudwatchMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchMetricName) > 0 {
		input.MetricName = aws.String(_cloudwatchMetricName)
	}
	if len(_cloudwatchNamespace) > 0 {
		input.Namespace = aws.String(_cloudwatchNamespace)
	}
	if len(_cloudwatchNextToken) > 0 {
		input.NextToken = aws.String(_cloudwatchNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeAnomalyDetectors(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudwatch.DescribeAnomalyDetectorsOutput
	p := cloudwatch.NewDescribeAnomalyDetectorsPaginator(client, input)
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

// Returns a list of all the Contributor Insights rules in your account.
// For more information about Contributor Insights, see [Using Contributor Insights to Analyze High-Cardinality Data].
//
// [Using Contributor Insights to Analyze High-Cardinality Data]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/ContributorInsights.html
func cloudwatch_DescribeInsightRules(cfg aws.Config, client *cloudwatch.Client) {
	input := &cloudwatch.DescribeInsightRulesInput{}

	if len(_cloudwatchMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cloudwatchMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchNextToken) > 0 {
		input.NextToken = aws.String(_cloudwatchNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeInsightRules(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudwatch.DescribeInsightRulesOutput
	p := cloudwatch.NewDescribeInsightRulesPaginator(client, input)
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

// Disables the actions for the specified alarms. When an alarm's actions are
// disabled, the alarm actions do not execute when the alarm state changes.
func cloudwatch_DisableAlarmActions(cfg aws.Config, client *cloudwatch.Client) {
	input := &cloudwatch.DisableAlarmActionsInput{
		// AlarmNames: []string, // Required
	}

	if len(_cloudwatchAlarmNames) > 0 {
		input.AlarmNames = append([]string(nil), _cloudwatchAlarmNames...)
	}

	if resp, err := client.DisableAlarmActions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disables the specified Contributor Insights rules. When rules are disabled,
// they do not analyze log groups and do not incur costs.
func cloudwatch_DisableInsightRules(cfg aws.Config, client *cloudwatch.Client) {
	input := &cloudwatch.DisableInsightRulesInput{
		// RuleNames: []string, // Required
	}

	if len(_cloudwatchRuleNames) > 0 {
		input.RuleNames = append([]string(nil), _cloudwatchRuleNames...)
	}

	if resp, err := client.DisableInsightRules(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables the actions for the specified alarms.
func cloudwatch_EnableAlarmActions(cfg aws.Config, client *cloudwatch.Client) {
	input := &cloudwatch.EnableAlarmActionsInput{
		// AlarmNames: []string, // Required
	}

	if len(_cloudwatchAlarmNames) > 0 {
		input.AlarmNames = append([]string(nil), _cloudwatchAlarmNames...)
	}

	if resp, err := client.EnableAlarmActions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables the specified Contributor Insights rules. When rules are enabled, they
// immediately begin analyzing log data.
func cloudwatch_EnableInsightRules(cfg aws.Config, client *cloudwatch.Client) {
	input := &cloudwatch.EnableInsightRulesInput{
		// RuleNames: []string, // Required
	}

	if len(_cloudwatchRuleNames) > 0 {
		input.RuleNames = append([]string(nil), _cloudwatchRuleNames...)
	}

	if resp, err := client.EnableInsightRules(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves details for a specific alarm mute rule.
// This operation returns complete information about the mute rule, including its
// configuration, status, targeted alarms, and metadata.
//
// The returned status indicates the current state of the mute rule:
//
// - SCHEDULED: The mute rule is configured and will become active in the future
//
// - ACTIVE: The mute rule is currently muting alarm actions
//
// - EXPIRED: The mute rule has passed its expiration date and will no longer
// become active
//
// # Permissions
//
// To retrieve details for a mute rule, you need the cloudwatch:GetAlarmMuteRule
// permission on the alarm mute rule resource.
func cloudwatch_GetAlarmMuteRule(cfg aws.Config, client *cloudwatch.Client) {
	input := &cloudwatch.GetAlarmMuteRuleInput{
		// AlarmMuteRuleName: *string, // Required
	}

	if len(_cloudwatchAlarmMuteRuleName) > 0 {
		input.AlarmMuteRuleName = aws.String(_cloudwatchAlarmMuteRuleName)
	}

	if resp, err := client.GetAlarmMuteRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Displays the details of the dashboard that you specify.
// To copy an existing dashboard, use GetDashboard , and then use the data returned
// within DashboardBody as the template for the new dashboard when you call
// PutDashboard to create the copy.
func cloudwatch_GetDashboard(cfg aws.Config, client *cloudwatch.Client) {
	input := &cloudwatch.GetDashboardInput{
		// DashboardName: *string, // Required
	}

	if len(_cloudwatchDashboardName) > 0 {
		input.DashboardName = aws.String(_cloudwatchDashboardName)
	}

	if resp, err := client.GetDashboard(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation returns the time series data collected by a Contributor Insights
// rule. The data includes the identity and number of contributors to the log
// group.
//
// You can also optionally return one or more statistics about each data point in
// the time series. These statistics can include the following:
//
// - UniqueContributors -- the number of unique contributors for each data point.
//
// - MaxContributorValue -- the value of the top contributor for each data point.
// The identity of the contributor might change for each data point in the graph.
//
// # If this rule aggregates by COUNT, the top contributor for each data point is
//
// the contributor with the most occurrences in that period. If the rule aggregates
// by SUM, the top contributor is the contributor with the highest sum in the log
// field specified by the rule's Value , during that period.
//
// - SampleCount -- the number of data points matched by the rule.
//
// - Sum -- the sum of the values from all contributors during the time period
// represented by that data point.
//
// - Minimum -- the minimum value from a single observation during the time
// period represented by that data point.
//
// - Maximum -- the maximum value from a single observation during the time
// period represented by that data point.
//
// - Average -- the average value from all contributors during the time period
// represented by that data point.
func cloudwatch_GetInsightRuleReport(cfg aws.Config, client *cloudwatch.Client) {
	input := &cloudwatch.GetInsightRuleReportInput{
		// EndTime: *time.Time, // Required
		// Period: *int32, // Required
		// RuleName: *string, // Required
		// StartTime: *time.Time, // Required
	}

	if len(_cloudwatchEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _cloudwatchEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchPeriod) > 0 {
		if err := assignInputField(input, "Period", _cloudwatchPeriod); err != nil {
			log.Errorf("invalid --period: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchRuleName) > 0 {
		input.RuleName = aws.String(_cloudwatchRuleName)
	}
	if len(_cloudwatchStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _cloudwatchStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchMaxContributorCount) > 0 {
		if err := assignInputField(input, "MaxContributorCount", _cloudwatchMaxContributorCount); err != nil {
			log.Errorf("invalid --max-contributor-count: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchMetrics) > 0 {
		input.Metrics = []string{_cloudwatchMetrics}
	}
	if len(_cloudwatchOrderBy) > 0 {
		input.OrderBy = aws.String(_cloudwatchOrderBy)
	}

	if resp, err := client.GetInsightRuleReport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// You can use the GetMetricData API to retrieve CloudWatch metric values. The
// operation can also include a CloudWatch Metrics Insights query, and one or more
// metric math functions.
//
// A GetMetricData operation that does not include a query can retrieve as many as
// 500 different metrics in a single request, with a total of as many as 100,800
// data points. You can also optionally perform metric math expressions on the
// values of the returned statistics, to create new time series that represent new
// insights into your data. For example, using Lambda metrics, you could divide the
// Errors metric by the Invocations metric to get an error rate time series. For
// more information about metric math expressions, see [Metric Math Syntax and Functions]in the Amazon CloudWatch
// User Guide.
//
// If you include a Metrics Insights query, each GetMetricData operation can
// include only one query. But the same GetMetricData operation can also retrieve
// other metrics. Metrics Insights queries can query only the most recent three
// hours of metric data. For more information about Metrics Insights, see [Query your metrics with CloudWatch Metrics Insights].
//
// Calls to the GetMetricData API have a different pricing structure than calls to
// GetMetricStatistics . For more information about pricing, see [Amazon CloudWatch Pricing].
//
// Amazon CloudWatch retains metric data as follows:
//
// - Data points with a period of less than 60 seconds are available for 3
// hours. These data points are high-resolution metrics and are available only for
// custom metrics that have been defined with a StorageResolution of 1.
//
// - Data points with a period of 60 seconds (1-minute) are available for 15
// days.
//
// - Data points with a period of 300 seconds (5-minute) are available for 63
// days.
//
// - Data points with a period of 3600 seconds (1 hour) are available for 455
// days (15 months).
//
// Data points that are initially published with a shorter period are aggregated
// together for long-term storage. For example, if you collect data using a period
// of 1 minute, the data remains available for 15 days with 1-minute resolution.
// After 15 days, this data is still available, but is aggregated and retrievable
// only with a resolution of 5 minutes. After 63 days, the data is further
// aggregated and is available with a resolution of 1 hour.
//
// If you omit Unit in your request, all data that was collected with any unit is
// returned, along with the corresponding units that were specified when the data
// was reported to CloudWatch. If you specify a unit, the operation returns only
// data that was collected with that unit specified. If you specify a unit that
// does not match the data collected, the results of the operation are null.
// CloudWatch does not perform unit conversions.
//
// # Using Metrics Insights queries with metric math
//
// You can't mix a Metric Insights query and metric math syntax in the same
// expression, but you can reference results from a Metrics Insights query within
// other Metric math expressions. A Metrics Insights query without a GROUP BY
// clause returns a single time-series (TS), and can be used as input for a metric
// math expression that expects a single time series. A Metrics Insights query with
// a GROUP BY clause returns an array of time-series (TS[]), and can be used as
// input for a metric math expression that expects an array of time series.
//
// [Metric Math Syntax and Functions]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/using-metric-math.html#metric-math-syntax
// [Query your metrics with CloudWatch Metrics Insights]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/query_with_cloudwatch-metrics-insights.html
// [Amazon CloudWatch Pricing]: https://aws.amazon.com/cloudwatch/pricing/
func cloudwatch_GetMetricData(cfg aws.Config, client *cloudwatch.Client) {
	input := &cloudwatch.GetMetricDataInput{
		// EndTime: *time.Time, // Required
		// MetricDataQueries: []types.MetricDataQuery, // Required
		// StartTime: *time.Time, // Required
	}

	if len(_cloudwatchEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _cloudwatchEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchMetricDataQueries) > 0 {
		if err := assignInputField(input, "MetricDataQueries", _cloudwatchMetricDataQueries); err != nil {
			log.Errorf("invalid --metric-data-queries: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _cloudwatchStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchLabelOptions) > 0 {
		if err := assignInputField(input, "LabelOptions", _cloudwatchLabelOptions); err != nil {
			log.Errorf("invalid --label-options: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchMaxDatapoints) > 0 {
		if err := assignInputField(input, "MaxDatapoints", _cloudwatchMaxDatapoints); err != nil {
			log.Errorf("invalid --max-datapoints: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchNextToken) > 0 {
		input.NextToken = aws.String(_cloudwatchNextToken)
	}
	if len(_cloudwatchScanBy) > 0 {
		if err := assignInputField(input, "ScanBy", _cloudwatchScanBy); err != nil {
			log.Errorf("invalid --scan-by: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.GetMetricData(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudwatch.GetMetricDataOutput
	p := cloudwatch.NewGetMetricDataPaginator(client, input)
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

// Gets statistics for the specified metric.
// The maximum number of data points returned from a single call is 1,440. If you
// request more than 1,440 data points, CloudWatch returns an error. To reduce the
// number of data points, you can narrow the specified time range and make multiple
// requests across adjacent time ranges, or you can increase the specified period.
// Data points are not returned in chronological order.
//
// CloudWatch aggregates data points based on the length of the period that you
// specify. For example, if you request statistics with a one-hour period,
// CloudWatch aggregates all data points with time stamps that fall within each
// one-hour period. Therefore, the number of values aggregated by CloudWatch is
// larger than the number of data points returned.
//
// CloudWatch needs raw data points to calculate percentile statistics. If you
// publish data using a statistic set instead, you can only retrieve percentile
// statistics for this data if one of the following conditions is true:
//
// - The SampleCount value of the statistic set is 1.
//
// - The Min and the Max values of the statistic set are equal.
//
// Percentile statistics are not available for metrics when any of the metric
// values are negative numbers.
//
// Amazon CloudWatch retains metric data as follows:
//
// - Data points with a period of less than 60 seconds are available for 3
// hours. These data points are high-resolution metrics and are available only for
// custom metrics that have been defined with a StorageResolution of 1.
//
// - Data points with a period of 60 seconds (1-minute) are available for 15
// days.
//
// - Data points with a period of 300 seconds (5-minute) are available for 63
// days.
//
// - Data points with a period of 3600 seconds (1 hour) are available for 455
// days (15 months).
//
// Data points that are initially published with a shorter period are aggregated
// together for long-term storage. For example, if you collect data using a period
// of 1 minute, the data remains available for 15 days with 1-minute resolution.
// After 15 days, this data is still available, but is aggregated and retrievable
// only with a resolution of 5 minutes. After 63 days, the data is further
// aggregated and is available with a resolution of 1 hour.
//
// CloudWatch started retaining 5-minute and 1-hour metric data as of July 9, 2016.
//
// For information about metrics and dimensions supported by Amazon Web Services
// services, see the [Amazon CloudWatch Metrics and Dimensions Reference]in the Amazon CloudWatch User Guide.
//
// [Amazon CloudWatch Metrics and Dimensions Reference]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CW_Support_For_AWS.html
func cloudwatch_GetMetricStatistics(cfg aws.Config, client *cloudwatch.Client) {
	input := &cloudwatch.GetMetricStatisticsInput{
		// EndTime: *time.Time, // Required
		// MetricName: *string, // Required
		// Namespace: *string, // Required
		// Period: *int32, // Required
		// StartTime: *time.Time, // Required
	}

	if len(_cloudwatchEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _cloudwatchEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchMetricName) > 0 {
		input.MetricName = aws.String(_cloudwatchMetricName)
	}
	if len(_cloudwatchNamespace) > 0 {
		input.Namespace = aws.String(_cloudwatchNamespace)
	}
	if len(_cloudwatchPeriod) > 0 {
		if err := assignInputField(input, "Period", _cloudwatchPeriod); err != nil {
			log.Errorf("invalid --period: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _cloudwatchStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchDimensions) > 0 {
		if err := assignInputField(input, "Dimensions", _cloudwatchDimensions); err != nil {
			log.Errorf("invalid --dimensions: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchExtendedStatistics) > 0 {
		input.ExtendedStatistics = append([]string(nil), _cloudwatchExtendedStatistics...)
	}
	if len(_cloudwatchStatistics) > 0 {
		if err := assignInputField(input, "Statistics", _cloudwatchStatistics); err != nil {
			log.Errorf("invalid --statistics: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchUnit) > 0 {
		if err := assignInputField(input, "Unit", _cloudwatchUnit); err != nil {
			log.Errorf("invalid --unit: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetMetricStatistics(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the metric stream that you specify.
func cloudwatch_GetMetricStream(cfg aws.Config, client *cloudwatch.Client) {
	input := &cloudwatch.GetMetricStreamInput{
		// Name: *string, // Required
	}

	if len(_cloudwatchName) > 0 {
		input.Name = aws.String(_cloudwatchName)
	}

	if resp, err := client.GetMetricStream(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// You can use the GetMetricWidgetImage API to retrieve a snapshot graph of one or
// more Amazon CloudWatch metrics as a bitmap image. You can then embed this image
// into your services and products, such as wiki pages, reports, and documents. You
// could also retrieve images regularly, such as every minute, and create your own
// custom live dashboard.
//
// The graph you retrieve can include all CloudWatch metric graph features,
// including metric math and horizontal and vertical annotations.
//
// There is a limit of 20 transactions per second for this API. Each
// GetMetricWidgetImage action has the following limits:
//
// - As many as 100 metrics in the graph.
//
// - Up to 100 KB uncompressed payload.
func cloudwatch_GetMetricWidgetImage(cfg aws.Config, client *cloudwatch.Client) {
	input := &cloudwatch.GetMetricWidgetImageInput{
		// MetricWidget: *string, // Required
	}

	if len(_cloudwatchMetricWidget) > 0 {
		input.MetricWidget = aws.String(_cloudwatchMetricWidget)
	}
	if len(_cloudwatchOutputFormat) > 0 {
		input.OutputFormat = aws.String(_cloudwatchOutputFormat)
	}

	if resp, err := client.GetMetricWidgetImage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists alarm mute rules in your Amazon Web Services account and region.
// You can filter the results by alarm name to find all mute rules targeting a
// specific alarm, or by status to find rules that are scheduled, active, or
// expired.
//
// This operation supports pagination for accounts with many mute rules. Use the
// MaxRecords and NextToken parameters to retrieve results in multiple calls.
//
// # Permissions
//
// To list mute rules, you need the cloudwatch:ListAlarmMuteRules permission.
func cloudwatch_ListAlarmMuteRules(cfg aws.Config, client *cloudwatch.Client) {
	input := &cloudwatch.ListAlarmMuteRulesInput{}

	if len(_cloudwatchAlarmName) > 0 {
		input.AlarmName = aws.String(_cloudwatchAlarmName)
	}
	if len(_cloudwatchMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _cloudwatchMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchNextToken) > 0 {
		input.NextToken = aws.String(_cloudwatchNextToken)
	}
	if len(_cloudwatchStatuses) > 0 {
		if err := assignInputField(input, "Statuses", _cloudwatchStatuses); err != nil {
			log.Errorf("invalid --statuses: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListAlarmMuteRules(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudwatch.ListAlarmMuteRulesOutput
	p := cloudwatch.NewListAlarmMuteRulesPaginator(client, input)
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

// Returns a list of the dashboards for your account. If you include
// DashboardNamePrefix , only those dashboards with names starting with the prefix
// are listed. Otherwise, all dashboards in your account are listed.
//
// ListDashboards returns up to 1000 results on one page. If there are more than
// 1000 dashboards, you can call ListDashboards again and include the value you
// received for NextToken in the first call, to receive the next 1000 results.
func cloudwatch_ListDashboards(cfg aws.Config, client *cloudwatch.Client) {
	input := &cloudwatch.ListDashboardsInput{}

	if len(_cloudwatchDashboardNamePrefix) > 0 {
		input.DashboardNamePrefix = aws.String(_cloudwatchDashboardNamePrefix)
	}
	if len(_cloudwatchNextToken) > 0 {
		input.NextToken = aws.String(_cloudwatchNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDashboards(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudwatch.ListDashboardsOutput
	p := cloudwatch.NewListDashboardsPaginator(client, input)
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

// Returns a list that contains the number of managed Contributor Insights rules
// in your account.
func cloudwatch_ListManagedInsightRules(cfg aws.Config, client *cloudwatch.Client) {
	input := &cloudwatch.ListManagedInsightRulesInput{
		// ResourceARN: *string, // Required
	}

	if len(_cloudwatchResourceARN) > 0 {
		input.ResourceARN = aws.String(_cloudwatchResourceARN)
	}
	if len(_cloudwatchMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cloudwatchMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchNextToken) > 0 {
		input.NextToken = aws.String(_cloudwatchNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListManagedInsightRules(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudwatch.ListManagedInsightRulesOutput
	p := cloudwatch.NewListManagedInsightRulesPaginator(client, input)
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

// Returns a list of metric streams in this account.
func cloudwatch_ListMetricStreams(cfg aws.Config, client *cloudwatch.Client) {
	input := &cloudwatch.ListMetricStreamsInput{}

	if len(_cloudwatchMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cloudwatchMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchNextToken) > 0 {
		input.NextToken = aws.String(_cloudwatchNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListMetricStreams(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudwatch.ListMetricStreamsOutput
	p := cloudwatch.NewListMetricStreamsPaginator(client, input)
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

// List the specified metrics. You can use the returned metrics with [GetMetricData] or [GetMetricStatistics] to get
// statistical data.
//
// Up to 500 results are returned for any one call. To retrieve additional
// results, use the returned token with subsequent calls.
//
// After you create a metric, allow up to 15 minutes for the metric to appear. To
// see metric statistics sooner, use [GetMetricData]or [GetMetricStatistics].
//
// If you are using CloudWatch cross-account observability, you can use this
// operation in a monitoring account and view metrics from the linked source
// accounts. For more information, see [CloudWatch cross-account observability].
//
// ListMetrics doesn't return information about metrics if those metrics haven't
// reported data in the past two weeks. To retrieve those metrics, use [GetMetricData]or [GetMetricStatistics].
//
// [GetMetricData]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_GetMetricData.html
// [GetMetricStatistics]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_GetMetricStatistics.html
// [CloudWatch cross-account observability]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Unified-Cross-Account.html
func cloudwatch_ListMetrics(cfg aws.Config, client *cloudwatch.Client) {
	input := &cloudwatch.ListMetricsInput{}

	if len(_cloudwatchDimensions) > 0 {
		if err := assignInputField(input, "Dimensions", _cloudwatchDimensions); err != nil {
			log.Errorf("invalid --dimensions: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchIncludeLinkedAccounts) > 0 {
		if err := assignInputField(input, "IncludeLinkedAccounts", _cloudwatchIncludeLinkedAccounts); err != nil {
			log.Errorf("invalid --include-linked-accounts: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchMetricName) > 0 {
		input.MetricName = aws.String(_cloudwatchMetricName)
	}
	if len(_cloudwatchNamespace) > 0 {
		input.Namespace = aws.String(_cloudwatchNamespace)
	}
	if len(_cloudwatchNextToken) > 0 {
		input.NextToken = aws.String(_cloudwatchNextToken)
	}
	if len(_cloudwatchOwningAccount) > 0 {
		input.OwningAccount = aws.String(_cloudwatchOwningAccount)
	}
	if len(_cloudwatchRecentlyActive) > 0 {
		if err := assignInputField(input, "RecentlyActive", _cloudwatchRecentlyActive); err != nil {
			log.Errorf("invalid --recently-active: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListMetrics(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudwatch.ListMetricsOutput
	p := cloudwatch.NewListMetricsPaginator(client, input)
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

// Displays the tags associated with a CloudWatch resource. Currently, alarms and
// Contributor Insights rules support tagging.
func cloudwatch_ListTagsForResource(cfg aws.Config, client *cloudwatch.Client) {
	input := &cloudwatch.ListTagsForResourceInput{
		// ResourceARN: *string, // Required
	}

	if len(_cloudwatchResourceARN) > 0 {
		input.ResourceARN = aws.String(_cloudwatchResourceARN)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates an alarm mute rule.
// Alarm mute rules automatically mute alarm actions during predefined time
// windows. When a mute rule is active, targeted alarms continue to evaluate
// metrics and transition between states, but their configured actions (such as
// Amazon SNS notifications or Auto Scaling actions) are muted.
//
// You can create mute rules with recurring schedules using cron expressions or
// one-time mute windows using at expressions. Each mute rule can target up to 100
// specific alarms by name.
//
// If you specify a rule name that already exists, this operation updates the
// existing rule with the new configuration.
//
// # Permissions
//
// To create or update a mute rule, you must have the cloudwatch:PutAlarmMuteRule
// permission on two types of resources: the alarm mute rule resource itself, and
// each alarm that the rule targets.
//
// For example, If you want to allow a user to create mute rules that target only
// specific alarms named "WebServerCPUAlarm" and "DatabaseConnectionAlarm", you
// would create an IAM policy with one statement granting
// cloudwatch:PutAlarmMuteRule on the alarm mute rule resource (
// arn:aws:cloudwatch:[REGION]:123456789012:alarm-mute:* ), and another statement
// granting cloudwatch:PutAlarmMuteRule on the targeted alarm resources (
// arn:aws:cloudwatch:[REGION]:123456789012:alarm:WebServerCPUAlarm and
// arn:aws:cloudwatch:[REGION]:123456789012:alarm:DatabaseConnectionAlarm ).
//
// You can also use IAM policy conditions to allow targeting alarms based on
// resource tags. For example, you can restrict users to create/update mute rules
// to only target alarms that have a specific tag key-value pair, such as
// Team=TeamA .
func cloudwatch_PutAlarmMuteRule(cfg aws.Config, client *cloudwatch.Client) {
	input := &cloudwatch.PutAlarmMuteRuleInput{
		// Name: *string, // Required
		// Rule: *types.Rule, // Required
	}

	if len(_cloudwatchName) > 0 {
		input.Name = aws.String(_cloudwatchName)
	}
	if len(_cloudwatchRule) > 0 {
		if err := assignInputField(input, "Rule", _cloudwatchRule); err != nil {
			log.Errorf("invalid --rule: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchDescription) > 0 {
		input.Description = aws.String(_cloudwatchDescription)
	}
	if len(_cloudwatchExpireDate) > 0 {
		if err := assignInputField(input, "ExpireDate", _cloudwatchExpireDate); err != nil {
			log.Errorf("invalid --expire-date: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchMuteTargets) > 0 {
		if err := assignInputField(input, "MuteTargets", _cloudwatchMuteTargets); err != nil {
			log.Errorf("invalid --mute-targets: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchStartDate) > 0 {
		if err := assignInputField(input, "StartDate", _cloudwatchStartDate); err != nil {
			log.Errorf("invalid --start-date: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchTags) > 0 {
		if err := assignInputField(input, "Tags", _cloudwatchTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutAlarmMuteRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an anomaly detection model for a CloudWatch metric. You can use the
// model to display a band of expected normal values when the metric is graphed.
//
// If you have enabled unified cross-account observability, and this account is a
// monitoring account, the metric can be in the same account or a source account.
// You can specify the account ID in the object you specify in the
// SingleMetricAnomalyDetector parameter.
//
// For more information, see [CloudWatch Anomaly Detection].
//
// [CloudWatch Anomaly Detection]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Anomaly_Detection.html
func cloudwatch_PutAnomalyDetector(cfg aws.Config, client *cloudwatch.Client) {
	input := &cloudwatch.PutAnomalyDetectorInput{}

	if len(_cloudwatchConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _cloudwatchConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchDimensions) > 0 {
		if err := assignInputField(input, "Dimensions", _cloudwatchDimensions); err != nil {
			log.Errorf("invalid --dimensions: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchMetricCharacteristics) > 0 {
		if err := assignInputField(input, "MetricCharacteristics", _cloudwatchMetricCharacteristics); err != nil {
			log.Errorf("invalid --metric-characteristics: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchMetricMathAnomalyDetector) > 0 {
		if err := assignInputField(input, "MetricMathAnomalyDetector", _cloudwatchMetricMathAnomalyDetector); err != nil {
			log.Errorf("invalid --metric-math-anomaly-detector: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchMetricName) > 0 {
		input.MetricName = aws.String(_cloudwatchMetricName)
	}
	if len(_cloudwatchNamespace) > 0 {
		input.Namespace = aws.String(_cloudwatchNamespace)
	}
	if len(_cloudwatchSingleMetricAnomalyDetector) > 0 {
		if err := assignInputField(input, "SingleMetricAnomalyDetector", _cloudwatchSingleMetricAnomalyDetector); err != nil {
			log.Errorf("invalid --single-metric-anomaly-detector: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchStat) > 0 {
		input.Stat = aws.String(_cloudwatchStat)
	}

	if resp, err := client.PutAnomalyDetector(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates a composite alarm. When you create a composite alarm, you
// specify a rule expression for the alarm that takes into account the alarm states
// of other alarms that you have created. The composite alarm goes into ALARM state
// only if all conditions of the rule are met.
//
// The alarms specified in a composite alarm's rule expression can include metric
// alarms and other composite alarms. The rule expression of a composite alarm can
// include as many as 100 underlying alarms. Any single alarm can be included in
// the rule expressions of as many as 150 composite alarms.
//
// Using composite alarms can reduce alarm noise. You can create multiple metric
// alarms, and also create a composite alarm and set up alerts only for the
// composite alarm. For example, you could create a composite alarm that goes into
// ALARM state only when more than one of the underlying metric alarms are in ALARM
// state.
//
// Composite alarms can take the following actions:
//
// - Notify Amazon SNS topics.
//
// - Invoke Lambda functions.
//
// - Create OpsItems in Systems Manager Ops Center.
//
// - Create incidents in Systems Manager Incident Manager.
//
// It is possible to create a loop or cycle of composite alarms, where composite
// alarm A depends on composite alarm B, and composite alarm B also depends on
// composite alarm A. In this scenario, you can't delete any composite alarm that
// is part of the cycle because there is always still a composite alarm that
// depends on that alarm that you want to delete.
//
// To get out of such a situation, you must break the cycle by changing the rule
// of one of the composite alarms in the cycle to remove a dependency that creates
// the cycle. The simplest change to make to break a cycle is to change the
// AlarmRule of one of the alarms to false .
//
// Additionally, the evaluation of composite alarms stops if CloudWatch detects a
// cycle in the evaluation path.
//
// When this operation creates an alarm, the alarm state is immediately set to
// INSUFFICIENT_DATA . The alarm is then evaluated and its state is set
// appropriately. Any actions associated with the new state are then executed. For
// a composite alarm, this initial time after creation is the only time that the
// alarm can be in INSUFFICIENT_DATA state.
//
// When you update an existing alarm, its state is left unchanged, but the update
// completely overwrites the previous configuration of the alarm.
//
// To use this operation, you must be signed on with the
// cloudwatch:PutCompositeAlarm permission that is scoped to * . You can't create a
// composite alarms if your cloudwatch:PutCompositeAlarm permission has a narrower
// scope.
//
// If you are an IAM user, you must have iam:CreateServiceLinkedRole to create a
// composite alarm that has Systems Manager OpsItem actions.
func cloudwatch_PutCompositeAlarm(cfg aws.Config, client *cloudwatch.Client) {
	input := &cloudwatch.PutCompositeAlarmInput{
		// AlarmName: *string, // Required
		// AlarmRule: *string, // Required
	}

	if len(_cloudwatchAlarmName) > 0 {
		input.AlarmName = aws.String(_cloudwatchAlarmName)
	}
	if len(_cloudwatchAlarmRule) > 0 {
		input.AlarmRule = aws.String(_cloudwatchAlarmRule)
	}
	if len(_cloudwatchActionsEnabled) > 0 {
		if err := assignInputField(input, "ActionsEnabled", _cloudwatchActionsEnabled); err != nil {
			log.Errorf("invalid --actions-enabled: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchActionsSuppressor) > 0 {
		input.ActionsSuppressor = aws.String(_cloudwatchActionsSuppressor)
	}
	if len(_cloudwatchActionsSuppressorExtensionPeriod) > 0 {
		if err := assignInputField(input, "ActionsSuppressorExtensionPeriod", _cloudwatchActionsSuppressorExtensionPeriod); err != nil {
			log.Errorf("invalid --actions-suppressor-extension-period: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchActionsSuppressorWaitPeriod) > 0 {
		if err := assignInputField(input, "ActionsSuppressorWaitPeriod", _cloudwatchActionsSuppressorWaitPeriod); err != nil {
			log.Errorf("invalid --actions-suppressor-wait-period: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchAlarmActions) > 0 {
		input.AlarmActions = append([]string(nil), _cloudwatchAlarmActions...)
	}
	if len(_cloudwatchAlarmDescription) > 0 {
		input.AlarmDescription = aws.String(_cloudwatchAlarmDescription)
	}
	if len(_cloudwatchInsufficientDataActions) > 0 {
		input.InsufficientDataActions = append([]string(nil), _cloudwatchInsufficientDataActions...)
	}
	if len(_cloudwatchOKActions) > 0 {
		input.OKActions = append([]string(nil), _cloudwatchOKActions...)
	}
	if len(_cloudwatchTags) > 0 {
		if err := assignInputField(input, "Tags", _cloudwatchTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutCompositeAlarm(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a dashboard if it does not already exist, or updates an existing
// dashboard. If you update a dashboard, the entire contents are replaced with what
// you specify here.
//
// All dashboards in your account are global, not region-specific.
//
// A simple way to create a dashboard using PutDashboard is to copy an existing
// dashboard. To copy an existing dashboard using the console, you can load the
// dashboard and then use the View/edit source command in the Actions menu to
// display the JSON block for that dashboard. Another way to copy a dashboard is to
// use GetDashboard , and then use the data returned within DashboardBody as the
// template for the new dashboard when you call PutDashboard .
//
// When you create a dashboard with PutDashboard , a good practice is to add a text
// widget at the top of the dashboard with a message that the dashboard was created
// by script and should not be changed in the console. This message could also
// point console users to the location of the DashboardBody script or the
// CloudFormation template used to create the dashboard.
func cloudwatch_PutDashboard(cfg aws.Config, client *cloudwatch.Client) {
	input := &cloudwatch.PutDashboardInput{
		// DashboardBody: *string, // Required
		// DashboardName: *string, // Required
	}

	if len(_cloudwatchDashboardBody) > 0 {
		input.DashboardBody = aws.String(_cloudwatchDashboardBody)
	}
	if len(_cloudwatchDashboardName) > 0 {
		input.DashboardName = aws.String(_cloudwatchDashboardName)
	}

	if resp, err := client.PutDashboard(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a Contributor Insights rule. Rules evaluate log events in a CloudWatch
// Logs log group, enabling you to find contributor data for the log events in that
// log group. For more information, see [Using Contributor Insights to Analyze High-Cardinality Data].
//
// If you create a rule, delete it, and then re-create it with the same name,
// historical data from the first time the rule was created might not be available.
//
// [Using Contributor Insights to Analyze High-Cardinality Data]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/ContributorInsights.html
func cloudwatch_PutInsightRule(cfg aws.Config, client *cloudwatch.Client) {
	input := &cloudwatch.PutInsightRuleInput{
		// RuleDefinition: *string, // Required
		// RuleName: *string, // Required
	}

	if len(_cloudwatchRuleDefinition) > 0 {
		input.RuleDefinition = aws.String(_cloudwatchRuleDefinition)
	}
	if len(_cloudwatchRuleName) > 0 {
		input.RuleName = aws.String(_cloudwatchRuleName)
	}
	if len(_cloudwatchApplyOnTransformedLogs) > 0 {
		if err := assignInputField(input, "ApplyOnTransformedLogs", _cloudwatchApplyOnTransformedLogs); err != nil {
			log.Errorf("invalid --apply-on-transformed-logs: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchRuleState) > 0 {
		input.RuleState = aws.String(_cloudwatchRuleState)
	}
	if len(_cloudwatchTags) > 0 {
		if err := assignInputField(input, "Tags", _cloudwatchTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutInsightRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a managed Contributor Insights rule for a specified Amazon Web
// Services resource. When you enable a managed rule, you create a Contributor
// Insights rule that collects data from Amazon Web Services services. You cannot
// edit these rules with PutInsightRule . The rules can be enabled, disabled, and
// deleted using EnableInsightRules , DisableInsightRules , and DeleteInsightRules
// . If a previously created managed rule is currently disabled, a subsequent call
// to this API will re-enable it. Use ListManagedInsightRules to describe all
// available rules.
func cloudwatch_PutManagedInsightRules(cfg aws.Config, client *cloudwatch.Client) {
	input := &cloudwatch.PutManagedInsightRulesInput{
		// ManagedRules: []types.ManagedRule, // Required
	}

	if len(_cloudwatchManagedRules) > 0 {
		if err := assignInputField(input, "ManagedRules", _cloudwatchManagedRules); err != nil {
			log.Errorf("invalid --managed-rules: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutManagedInsightRules(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates an alarm and associates it with the specified metric, metric
// math expression, anomaly detection model, or Metrics Insights query. For more
// information about using a Metrics Insights query for an alarm, see [Create alarms on Metrics Insights queries].
//
// Alarms based on anomaly detection models cannot have Auto Scaling actions.
//
// When this operation creates an alarm, the alarm state is immediately set to
// INSUFFICIENT_DATA . The alarm is then evaluated and its state is set
// appropriately. Any actions associated with the new state are then executed.
//
// When you update an existing alarm, its state is left unchanged, but the update
// completely overwrites the previous configuration of the alarm.
//
// If you are an IAM user, you must have Amazon EC2 permissions for some alarm
// operations:
//
// - The iam:CreateServiceLinkedRole permission for all alarms with EC2 actions
//
// - The iam:CreateServiceLinkedRole permissions to create an alarm with Systems
// Manager OpsItem or response plan actions.
//
// The first time you create an alarm in the Amazon Web Services Management
// Console, the CLI, or by using the PutMetricAlarm API, CloudWatch creates the
// necessary service-linked role for you. The service-linked roles are called
// AWSServiceRoleForCloudWatchEvents and
// AWSServiceRoleForCloudWatchAlarms_ActionSSM . For more information, see [Amazon Web Services service-linked role].
//
// Each PutMetricAlarm action has a maximum uncompressed payload of 120 KB.
//
// # Cross-account alarms
//
// You can set an alarm on metrics in the current account, or in another account.
// To create a cross-account alarm that watches a metric in a different account,
// you must have completed the following pre-requisites:
//
// - The account where the metrics are located (the sharing account) must
// already have a sharing role named CloudWatch-CrossAccountSharingRole. If it does
// not already have this role, you must create it using the instructions in Set up
// a sharing account in [Cross-account cross-Region CloudWatch console]. The policy for that role must grant access to the ID
// of the account where you are creating the alarm.
//
// - The account where you are creating the alarm (the monitoring account) must
// already have a service-linked role named AWSServiceRoleForCloudWatchCrossAccount
// to allow CloudWatch to assume the sharing role in the sharing account. If it
// does not, you must create it following the directions in Set up a monitoring
// account in [Cross-account cross-Region CloudWatch console].
//
// [Amazon Web Services service-linked role]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_terms-and-concepts.html#iam-term-service-linked-role
// [Create alarms on Metrics Insights queries]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Create_Metrics_Insights_Alarm.html
// [Cross-account cross-Region CloudWatch console]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Cross-Account-Cross-Region.html#enable-cross-account-cross-Region
func cloudwatch_PutMetricAlarm(cfg aws.Config, client *cloudwatch.Client) {
	input := &cloudwatch.PutMetricAlarmInput{
		// AlarmName: *string, // Required
		// ComparisonOperator: types.ComparisonOperator, // Required
		// EvaluationPeriods: *int32, // Required
	}

	if len(_cloudwatchAlarmName) > 0 {
		input.AlarmName = aws.String(_cloudwatchAlarmName)
	}
	if len(_cloudwatchComparisonOperator) > 0 {
		if err := assignInputField(input, "ComparisonOperator", _cloudwatchComparisonOperator); err != nil {
			log.Errorf("invalid --comparison-operator: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchEvaluationPeriods) > 0 {
		if err := assignInputField(input, "EvaluationPeriods", _cloudwatchEvaluationPeriods); err != nil {
			log.Errorf("invalid --evaluation-periods: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchActionsEnabled) > 0 {
		if err := assignInputField(input, "ActionsEnabled", _cloudwatchActionsEnabled); err != nil {
			log.Errorf("invalid --actions-enabled: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchAlarmActions) > 0 {
		input.AlarmActions = append([]string(nil), _cloudwatchAlarmActions...)
	}
	if len(_cloudwatchAlarmDescription) > 0 {
		input.AlarmDescription = aws.String(_cloudwatchAlarmDescription)
	}
	if len(_cloudwatchDatapointsToAlarm) > 0 {
		if err := assignInputField(input, "DatapointsToAlarm", _cloudwatchDatapointsToAlarm); err != nil {
			log.Errorf("invalid --datapoints-to-alarm: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchDimensions) > 0 {
		if err := assignInputField(input, "Dimensions", _cloudwatchDimensions); err != nil {
			log.Errorf("invalid --dimensions: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchEvaluateLowSampleCountPercentile) > 0 {
		input.EvaluateLowSampleCountPercentile = aws.String(_cloudwatchEvaluateLowSampleCountPercentile)
	}
	if len(_cloudwatchExtendedStatistic) > 0 {
		input.ExtendedStatistic = aws.String(_cloudwatchExtendedStatistic)
	}
	if len(_cloudwatchInsufficientDataActions) > 0 {
		input.InsufficientDataActions = append([]string(nil), _cloudwatchInsufficientDataActions...)
	}
	if len(_cloudwatchMetricName) > 0 {
		input.MetricName = aws.String(_cloudwatchMetricName)
	}
	if len(_cloudwatchMetrics) > 0 {
		if err := assignInputField(input, "Metrics", _cloudwatchMetrics); err != nil {
			log.Errorf("invalid --metrics: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchNamespace) > 0 {
		input.Namespace = aws.String(_cloudwatchNamespace)
	}
	if len(_cloudwatchOKActions) > 0 {
		input.OKActions = append([]string(nil), _cloudwatchOKActions...)
	}
	if len(_cloudwatchPeriod) > 0 {
		if err := assignInputField(input, "Period", _cloudwatchPeriod); err != nil {
			log.Errorf("invalid --period: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchStatistic) > 0 {
		if err := assignInputField(input, "Statistic", _cloudwatchStatistic); err != nil {
			log.Errorf("invalid --statistic: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchTags) > 0 {
		if err := assignInputField(input, "Tags", _cloudwatchTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchThreshold) > 0 {
		if err := assignInputField(input, "Threshold", _cloudwatchThreshold); err != nil {
			log.Errorf("invalid --threshold: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchThresholdMetricId) > 0 {
		input.ThresholdMetricId = aws.String(_cloudwatchThresholdMetricId)
	}
	if len(_cloudwatchTreatMissingData) > 0 {
		input.TreatMissingData = aws.String(_cloudwatchTreatMissingData)
	}
	if len(_cloudwatchUnit) > 0 {
		if err := assignInputField(input, "Unit", _cloudwatchUnit); err != nil {
			log.Errorf("invalid --unit: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutMetricAlarm(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Publishes metric data to Amazon CloudWatch. CloudWatch associates the data with
// the specified metric. If the specified metric does not exist, CloudWatch creates
// the metric. When CloudWatch creates a metric, it can take up to fifteen minutes
// for the metric to appear in calls to [ListMetrics].
//
// You can publish metrics with associated entity data (so that related telemetry
// can be found and viewed together), or publish metric data by itself. To send
// entity data with your metrics, use the EntityMetricData parameter. To send
// metrics without entity data, use the MetricData parameter. The EntityMetricData
// structure includes MetricData structures for the metric data.
//
// You can publish either individual values in the Value field, or arrays of
// values and the number of times each value occurred during the period by using
// the Values and Counts fields in the MetricData structure. Using the Values and
// Counts method enables you to publish up to 150 values per metric with one
// PutMetricData request, and supports retrieving percentile statistics on this
// data.
//
// Each PutMetricData request is limited to 1 MB in size for HTTP POST requests.
// You can send a payload compressed by gzip. Each request is also limited to no
// more than 1000 different metrics (across both the MetricData and
// EntityMetricData properties).
//
// Although the Value parameter accepts numbers of type Double , CloudWatch rejects
// values that are either too small or too large. Values must be in the range of
// -2^360 to 2^360. In addition, special values (for example, NaN, +Infinity,
// -Infinity) are not supported.
//
// You can use up to 30 dimensions per metric to further clarify what data the
// metric collects. Each dimension consists of a Name and Value pair. For more
// information about specifying dimensions, see [Publishing Metrics]in the Amazon CloudWatch User
// Guide.
//
// You specify the time stamp to be associated with each data point. You can
// specify time stamps that are as much as two weeks before the current date, and
// as much as 2 hours after the current day and time.
//
// Data points with time stamps from 24 hours ago or longer can take at least 48
// hours to become available for [GetMetricData]or [GetMetricStatistics] from the time they are submitted. Data points
// with time stamps between 3 and 24 hours ago can take as much as 2 hours to
// become available for [GetMetricData]or [GetMetricStatistics].
//
// CloudWatch needs raw data points to calculate percentile statistics. If you
// publish data using a statistic set instead, you can only retrieve percentile
// statistics for this data if one of the following conditions is true:
//
// - The SampleCount value of the statistic set is 1 and Min , Max , and Sum are
// all equal.
//
// - The Min and Max are equal, and Sum is equal to Min multiplied by SampleCount
// .
//
// [GetMetricData]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_GetMetricData.html
// [GetMetricStatistics]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_GetMetricStatistics.html
// [ListMetrics]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_ListMetrics.html
// [Publishing Metrics]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/publishingMetrics.html
func cloudwatch_PutMetricData(cfg aws.Config, client *cloudwatch.Client) {
	input := &cloudwatch.PutMetricDataInput{
		// Namespace: *string, // Required
	}

	if len(_cloudwatchNamespace) > 0 {
		input.Namespace = aws.String(_cloudwatchNamespace)
	}
	if len(_cloudwatchEntityMetricData) > 0 {
		if err := assignInputField(input, "EntityMetricData", _cloudwatchEntityMetricData); err != nil {
			log.Errorf("invalid --entity-metric-data: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchMetricData) > 0 {
		if err := assignInputField(input, "MetricData", _cloudwatchMetricData); err != nil {
			log.Errorf("invalid --metric-data: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchStrictEntityValidation) > 0 {
		if err := assignInputField(input, "StrictEntityValidation", _cloudwatchStrictEntityValidation); err != nil {
			log.Errorf("invalid --strict-entity-validation: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutMetricData(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates a metric stream. Metric streams can automatically stream
// CloudWatch metrics to Amazon Web Services destinations, including Amazon S3, and
// to many third-party solutions.
//
// For more information, see [Using Metric Streams].
//
// To create a metric stream, you must be signed in to an account that has the
// iam:PassRole permission and either the CloudWatchFullAccess policy or the
// cloudwatch:PutMetricStream permission.
//
// When you create or update a metric stream, you choose one of the following:
//
// - Stream metrics from all metric namespaces in the account.
//
// - Stream metrics from all metric namespaces in the account, except for the
// namespaces that you list in ExcludeFilters .
//
// - Stream metrics from only the metric namespaces that you list in
// IncludeFilters .
//
// By default, a metric stream always sends the MAX , MIN , SUM , and SAMPLECOUNT
// statistics for each metric that is streamed. You can use the
// StatisticsConfigurations parameter to have the metric stream send additional
// statistics in the stream. Streaming additional statistics incurs additional
// costs. For more information, see [Amazon CloudWatch Pricing].
//
// When you use PutMetricStream to create a new metric stream, the stream is
// created in the running state. If you use it to update an existing stream, the
// state of the stream is not changed.
//
// If you are using CloudWatch cross-account observability and you create a metric
// stream in a monitoring account, you can choose whether to include metrics from
// source accounts in the stream. For more information, see [CloudWatch cross-account observability].
//
// [Using Metric Streams]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Metric-Streams.html
// [CloudWatch cross-account observability]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Unified-Cross-Account.html
// [Amazon CloudWatch Pricing]: https://aws.amazon.com/cloudwatch/pricing/
func cloudwatch_PutMetricStream(cfg aws.Config, client *cloudwatch.Client) {
	input := &cloudwatch.PutMetricStreamInput{
		// FirehoseArn: *string, // Required
		// Name: *string, // Required
		// OutputFormat: types.MetricStreamOutputFormat, // Required
		// RoleArn: *string, // Required
	}

	if len(_cloudwatchFirehoseArn) > 0 {
		input.FirehoseArn = aws.String(_cloudwatchFirehoseArn)
	}
	if len(_cloudwatchName) > 0 {
		input.Name = aws.String(_cloudwatchName)
	}
	if len(_cloudwatchOutputFormat) > 0 {
		if err := assignInputField(input, "OutputFormat", _cloudwatchOutputFormat); err != nil {
			log.Errorf("invalid --output-format: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchRoleArn) > 0 {
		input.RoleArn = aws.String(_cloudwatchRoleArn)
	}
	if len(_cloudwatchExcludeFilters) > 0 {
		if err := assignInputField(input, "ExcludeFilters", _cloudwatchExcludeFilters); err != nil {
			log.Errorf("invalid --exclude-filters: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchIncludeFilters) > 0 {
		if err := assignInputField(input, "IncludeFilters", _cloudwatchIncludeFilters); err != nil {
			log.Errorf("invalid --include-filters: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchIncludeLinkedAccountsMetrics) > 0 {
		if err := assignInputField(input, "IncludeLinkedAccountsMetrics", _cloudwatchIncludeLinkedAccountsMetrics); err != nil {
			log.Errorf("invalid --include-linked-accounts-metrics: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchStatisticsConfigurations) > 0 {
		if err := assignInputField(input, "StatisticsConfigurations", _cloudwatchStatisticsConfigurations); err != nil {
			log.Errorf("invalid --statistics-configurations: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchTags) > 0 {
		if err := assignInputField(input, "Tags", _cloudwatchTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutMetricStream(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Temporarily sets the state of an alarm for testing purposes. When the updated
// state differs from the previous value, the action configured for the appropriate
// state is invoked. For example, if your alarm is configured to send an Amazon SNS
// message when an alarm is triggered, temporarily changing the alarm state to
// ALARM sends an SNS message.
//
// Metric alarms returns to their actual state quickly, often within seconds.
// Because the metric alarm state change happens quickly, it is typically only
// visible in the alarm's History tab in the Amazon CloudWatch console or through [DescribeAlarmHistory].
//
// If you use SetAlarmState on a composite alarm, the composite alarm is not
// guaranteed to return to its actual state. It returns to its actual state only
// once any of its children alarms change state. It is also reevaluated if you
// update its configuration.
//
// If an alarm triggers EC2 Auto Scaling policies or application Auto Scaling
// policies, you must include information in the StateReasonData parameter to
// enable the policy to take the correct action.
//
// [DescribeAlarmHistory]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_DescribeAlarmHistory.html
func cloudwatch_SetAlarmState(cfg aws.Config, client *cloudwatch.Client) {
	input := &cloudwatch.SetAlarmStateInput{
		// AlarmName: *string, // Required
		// StateReason: *string, // Required
		// StateValue: types.StateValue, // Required
	}

	if len(_cloudwatchAlarmName) > 0 {
		input.AlarmName = aws.String(_cloudwatchAlarmName)
	}
	if len(_cloudwatchStateReason) > 0 {
		input.StateReason = aws.String(_cloudwatchStateReason)
	}
	if len(_cloudwatchStateValue) > 0 {
		if err := assignInputField(input, "StateValue", _cloudwatchStateValue); err != nil {
			log.Errorf("invalid --state-value: %s", err.Error())
			return
		}
	}
	if len(_cloudwatchStateReasonData) > 0 {
		input.StateReasonData = aws.String(_cloudwatchStateReasonData)
	}

	if resp, err := client.SetAlarmState(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts the streaming of metrics for one or more of your metric streams.
func cloudwatch_StartMetricStreams(cfg aws.Config, client *cloudwatch.Client) {
	input := &cloudwatch.StartMetricStreamsInput{
		// Names: []string, // Required
	}

	if len(_cloudwatchNames) > 0 {
		input.Names = append([]string(nil), _cloudwatchNames...)
	}

	if resp, err := client.StartMetricStreams(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops the streaming of metrics for one or more of your metric streams.
func cloudwatch_StopMetricStreams(cfg aws.Config, client *cloudwatch.Client) {
	input := &cloudwatch.StopMetricStreamsInput{
		// Names: []string, // Required
	}

	if len(_cloudwatchNames) > 0 {
		input.Names = append([]string(nil), _cloudwatchNames...)
	}

	if resp, err := client.StopMetricStreams(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Assigns one or more tags (key-value pairs) to the specified CloudWatch
// resource. Currently, the only CloudWatch resources that can be tagged are alarms
// and Contributor Insights rules.
//
// Tags can help you organize and categorize your resources. You can also use them
// to scope user permissions by granting a user permission to access or change only
// resources with certain tag values.
//
// Tags don't have any semantic meaning to Amazon Web Services and are interpreted
// strictly as strings of characters.
//
// You can use the TagResource action with an alarm that already has tags. If you
// specify a new tag key for the alarm, this tag is appended to the list of tags
// associated with the alarm. If you specify a tag key that is already associated
// with the alarm, the new tag value that you specify replaces the previous value
// for that tag.
//
// You can associate as many as 50 tags with a CloudWatch resource.
func cloudwatch_TagResource(cfg aws.Config, client *cloudwatch.Client) {
	input := &cloudwatch.TagResourceInput{
		// ResourceARN: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_cloudwatchResourceARN) > 0 {
		input.ResourceARN = aws.String(_cloudwatchResourceARN)
	}
	if len(_cloudwatchTags) > 0 {
		if err := assignInputField(input, "Tags", _cloudwatchTags); err != nil {
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

// Removes one or more tags from the specified resource.
func cloudwatch_UntagResource(cfg aws.Config, client *cloudwatch.Client) {
	input := &cloudwatch.UntagResourceInput{
		// ResourceARN: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_cloudwatchResourceARN) > 0 {
		input.ResourceARN = aws.String(_cloudwatchResourceARN)
	}
	if len(_cloudwatchTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _cloudwatchTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_cloudwatchCmd)
	_cloudwatchCmd.Flags().SortFlags = false

	_cloudwatchCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_cloudwatchCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_cloudwatchCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchActionPrefix, "action-prefix", "", "", "Action Prefix")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchActionsEnabled, "actions-enabled", "", "", "Actions Enabled")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchActionsSuppressor, "actions-suppressor", "", "", "Actions Suppressor")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchActionsSuppressorExtensionPeriod, "actions-suppressor-extension-period", "", "", "Actions Suppressor Extension Period")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchActionsSuppressorWaitPeriod, "actions-suppressor-wait-period", "", "", "Actions Suppressor Wait Period")
	_cloudwatchCmd.Flags().StringSliceVarP(&_cloudwatchAlarmActions, "alarm-actions", "", nil, "Alarm Actions")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchAlarmContributorId, "alarm-contributor-id", "", "", "Alarm Contributor ID")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchAlarmDescription, "alarm-description", "", "", "Alarm Description")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchAlarmMuteRuleName, "alarm-mute-rule-name", "", "", "Alarm Mute Rule Name")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchAlarmName, "alarm-name", "", "", "Alarm Name")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchAlarmNamePrefix, "alarm-name-prefix", "", "", "Alarm Name Prefix")
	_cloudwatchCmd.Flags().StringSliceVarP(&_cloudwatchAlarmNames, "alarm-names", "", nil, "Alarm Names")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchAlarmRule, "alarm-rule", "", "", "Alarm Rule")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchAlarmTypes, "alarm-types", "", "", "Alarm Types")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchAnomalyDetectorTypes, "anomaly-detector-types", "", "", "Anomaly Detector Types")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchApplyOnTransformedLogs, "apply-on-transformed-logs", "", "", "Apply On Transformed Logs")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchChildrenOfAlarmName, "children-of-alarm-name", "", "", "Children Of Alarm Name")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchComparisonOperator, "comparison-operator", "", "", "Comparison Operator")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchConfiguration, "configuration", "", "", "Configuration")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchDashboardBody, "dashboard-body", "", "", "Dashboard Body")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchDashboardName, "dashboard-name", "", "", "Dashboard Name")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchDashboardNamePrefix, "dashboard-name-prefix", "", "", "Dashboard Name Prefix")
	_cloudwatchCmd.Flags().StringSliceVarP(&_cloudwatchDashboardNames, "dashboard-names", "", nil, "Dashboard Names")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchDatapointsToAlarm, "datapoints-to-alarm", "", "", "Datapoints To Alarm")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchDescription, "description", "", "", "Description")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchDimensions, "dimensions", "", "", "Dimensions")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchEndDate, "end-date", "", "", "End Date")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchEndTime, "end-time", "", "", "End Time")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchEntityMetricData, "entity-metric-data", "", "", "Entity Metric Data")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchEvaluateLowSampleCountPercentile, "evaluate-low-sample-count-percentile", "", "", "Evaluate Low Sample Count Percentile")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchEvaluationPeriods, "evaluation-periods", "", "", "Evaluation Periods")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchExcludeFilters, "exclude-filters", "", "", "Exclude Filters")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchExpireDate, "expire-date", "", "", "Expire Date")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchExtendedStatistic, "extended-statistic", "", "", "Extended Statistic")
	_cloudwatchCmd.Flags().StringSliceVarP(&_cloudwatchExtendedStatistics, "extended-statistics", "", nil, "Extended Statistics")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchFirehoseArn, "firehose-arn", "", "", "Firehose ARN")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchHistoryItemType, "history-item-type", "", "", "History Item Type")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchIncludeFilters, "include-filters", "", "", "Include Filters")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchIncludeLinkedAccounts, "include-linked-accounts", "", "", "Include Linked Accounts")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchIncludeLinkedAccountsMetrics, "include-linked-accounts-metrics", "", "", "Include Linked Accounts Metrics")
	_cloudwatchCmd.Flags().StringSliceVarP(&_cloudwatchInsufficientDataActions, "insufficient-data-actions", "", nil, "Insufficient Data Actions")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchLabelOptions, "label-options", "", "", "Label Options")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchManagedRules, "managed-rules", "", "", "Managed Rules")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchMaxContributorCount, "max-contributor-count", "", "", "Max Contributor Count")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchMaxDatapoints, "max-datapoints", "", "", "Max Datapoints")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchMaxRecords, "max-records", "", "", "Max Records")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchMaxResults, "max-results", "", "", "Max Results")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchMetricCharacteristics, "metric-characteristics", "", "", "Metric Characteristics")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchMetricData, "metric-data", "", "", "Metric Data")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchMetricDataQueries, "metric-data-queries", "", "", "Metric Data Queries")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchMetricMathAnomalyDetector, "metric-math-anomaly-detector", "", "", "Metric Math Anomaly Detector")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchMetricName, "metric-name", "", "", "Metric Name")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchMetricWidget, "metric-widget", "", "", "Metric Widget")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchMetrics, "metrics", "", "", "Metrics")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchMuteTargets, "mute-targets", "", "", "Mute Targets")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchName, "name", "", "", "Name")
	_cloudwatchCmd.Flags().StringSliceVarP(&_cloudwatchNames, "names", "", nil, "Names")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchNamespace, "namespace", "", "", "Namespace")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchNextToken, "next-token", "", "", "Next Token")
	_cloudwatchCmd.Flags().StringSliceVarP(&_cloudwatchOKActions, "ok-actions", "", nil, "Ok Actions")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchOrderBy, "order-by", "", "", "Order By")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchOutputFormat, "output-format", "", "", "Output Format")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchOwningAccount, "owning-account", "", "", "Owning Account")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchParentsOfAlarmName, "parents-of-alarm-name", "", "", "Parents Of Alarm Name")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchPeriod, "period", "", "", "Period")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchRecentlyActive, "recently-active", "", "", "Recently Active")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchResourceARN, "resource-arn", "", "", "Resource ARN")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchRoleArn, "role-arn", "", "", "Role ARN")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchRule, "rule", "", "", "Rule")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchRuleDefinition, "rule-definition", "", "", "Rule Definition")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchRuleName, "rule-name", "", "", "Rule Name")
	_cloudwatchCmd.Flags().StringSliceVarP(&_cloudwatchRuleNames, "rule-names", "", nil, "Rule Names")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchRuleState, "rule-state", "", "", "Rule State")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchScanBy, "scan-by", "", "", "Scan By")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchSingleMetricAnomalyDetector, "single-metric-anomaly-detector", "", "", "Single Metric Anomaly Detector")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchStartDate, "start-date", "", "", "Start Date")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchStartTime, "start-time", "", "", "Start Time")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchStat, "stat", "", "", "Stat")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchStateReason, "state-reason", "", "", "State Reason")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchStateReasonData, "state-reason-data", "", "", "State Reason Data")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchStateValue, "state-value", "", "", "State Value")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchStatistic, "statistic", "", "", "Statistic")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchStatistics, "statistics", "", "", "Statistics")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchStatisticsConfigurations, "statistics-configurations", "", "", "Statistics Configurations")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchStatuses, "statuses", "", "", "Statuses")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchStrictEntityValidation, "strict-entity-validation", "", "", "Strict Entity Validation")
	_cloudwatchCmd.Flags().StringSliceVarP(&_cloudwatchTagKeys, "tag-keys", "", nil, "Tag Keys")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchTags, "tags", "", "", "Tags")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchThreshold, "threshold", "", "", "Threshold")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchThresholdMetricId, "threshold-metric-id", "", "", "Threshold Metric ID")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchTreatMissingData, "treat-missing-data", "", "", "Treat Missing Data")
	_cloudwatchCmd.Flags().StringVarP(&_cloudwatchUnit, "unit", "", "", "Unit")

	_cloudwatchCmd.Flags().BoolVarP(&_cloudwatchDeleteAlarmMuteRule, "delete-alarm-mute-rule", "", false, "Delete Alarm Mute Rule")
	_cloudwatchCmd.Flags().BoolVarP(&_cloudwatchDeleteAlarms, "delete-alarms", "", false, "Delete Alarms")
	_cloudwatchCmd.Flags().BoolVarP(&_cloudwatchDeleteAnomalyDetector, "delete-anomaly-detector", "", false, "Delete Anomaly Detector")
	_cloudwatchCmd.Flags().BoolVarP(&_cloudwatchDeleteDashboards, "delete-dashboards", "", false, "Delete Dashboards")
	_cloudwatchCmd.Flags().BoolVarP(&_cloudwatchDeleteInsightRules, "delete-insight-rules", "", false, "Delete Insight Rules")
	_cloudwatchCmd.Flags().BoolVarP(&_cloudwatchDeleteMetricStream, "delete-metric-stream", "", false, "Delete Metric Stream")
	_cloudwatchCmd.Flags().BoolVarP(&_cloudwatchDescribeAlarmContributors, "describe-alarm-contributors", "", false, "Describe Alarm Contributors")
	_cloudwatchCmd.Flags().BoolVarP(&_cloudwatchDescribeAlarmHistory, "describe-alarm-history", "", false, "Describe Alarm History")
	_cloudwatchCmd.Flags().BoolVarP(&_cloudwatchDescribeAlarms, "describe-alarms", "", false, "Describe Alarms")
	_cloudwatchCmd.Flags().BoolVarP(&_cloudwatchDescribeAlarmsForMetric, "describe-alarms-for-metric", "", false, "Describe Alarms For Metric")
	_cloudwatchCmd.Flags().BoolVarP(&_cloudwatchDescribeAnomalyDetectors, "describe-anomaly-detectors", "", false, "Describe Anomaly Detectors")
	_cloudwatchCmd.Flags().BoolVarP(&_cloudwatchDescribeInsightRules, "describe-insight-rules", "", false, "Describe Insight Rules")
	_cloudwatchCmd.Flags().BoolVarP(&_cloudwatchDisableAlarmActions, "disable-alarm-actions", "", false, "Disable Alarm Actions")
	_cloudwatchCmd.Flags().BoolVarP(&_cloudwatchDisableInsightRules, "disable-insight-rules", "", false, "Disable Insight Rules")
	_cloudwatchCmd.Flags().BoolVarP(&_cloudwatchEnableAlarmActions, "enable-alarm-actions", "", false, "Enable Alarm Actions")
	_cloudwatchCmd.Flags().BoolVarP(&_cloudwatchEnableInsightRules, "enable-insight-rules", "", false, "Enable Insight Rules")
	_cloudwatchCmd.Flags().BoolVarP(&_cloudwatchGetAlarmMuteRule, "get-alarm-mute-rule", "", false, "Get Alarm Mute Rule")
	_cloudwatchCmd.Flags().BoolVarP(&_cloudwatchGetDashboard, "get-dashboard", "", false, "Get Dashboard")
	_cloudwatchCmd.Flags().BoolVarP(&_cloudwatchGetInsightRuleReport, "get-insight-rule-report", "", false, "Get Insight Rule Report")
	_cloudwatchCmd.Flags().BoolVarP(&_cloudwatchGetMetricData, "get-metric-data", "", false, "Get Metric Data")
	_cloudwatchCmd.Flags().BoolVarP(&_cloudwatchGetMetricStatistics, "get-metric-statistics", "", false, "Get Metric Statistics")
	_cloudwatchCmd.Flags().BoolVarP(&_cloudwatchGetMetricStream, "get-metric-stream", "", false, "Get Metric Stream")
	_cloudwatchCmd.Flags().BoolVarP(&_cloudwatchGetMetricWidgetImage, "get-metric-widget-image", "", false, "Get Metric Widget Image")
	_cloudwatchCmd.Flags().BoolVarP(&_cloudwatchListAlarmMuteRules, "list-alarm-mute-rules", "", false, "List Alarm Mute Rules")
	_cloudwatchCmd.Flags().BoolVarP(&_cloudwatchListDashboards, "list-dashboards", "", false, "List Dashboards")
	_cloudwatchCmd.Flags().BoolVarP(&_cloudwatchListManagedInsightRules, "list-managed-insight-rules", "", false, "List Managed Insight Rules")
	_cloudwatchCmd.Flags().BoolVarP(&_cloudwatchListMetricStreams, "list-metric-streams", "", false, "List Metric Streams")
	_cloudwatchCmd.Flags().BoolVarP(&_cloudwatchListMetrics, "list-metrics", "", false, "List Metrics")
	_cloudwatchCmd.Flags().BoolVarP(&_cloudwatchListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_cloudwatchCmd.Flags().BoolVarP(&_cloudwatchPutAlarmMuteRule, "put-alarm-mute-rule", "", false, "Put Alarm Mute Rule")
	_cloudwatchCmd.Flags().BoolVarP(&_cloudwatchPutAnomalyDetector, "put-anomaly-detector", "", false, "Put Anomaly Detector")
	_cloudwatchCmd.Flags().BoolVarP(&_cloudwatchPutCompositeAlarm, "put-composite-alarm", "", false, "Put Composite Alarm")
	_cloudwatchCmd.Flags().BoolVarP(&_cloudwatchPutDashboard, "put-dashboard", "", false, "Put Dashboard")
	_cloudwatchCmd.Flags().BoolVarP(&_cloudwatchPutInsightRule, "put-insight-rule", "", false, "Put Insight Rule")
	_cloudwatchCmd.Flags().BoolVarP(&_cloudwatchPutManagedInsightRules, "put-managed-insight-rules", "", false, "Put Managed Insight Rules")
	_cloudwatchCmd.Flags().BoolVarP(&_cloudwatchPutMetricAlarm, "put-metric-alarm", "", false, "Put Metric Alarm")
	_cloudwatchCmd.Flags().BoolVarP(&_cloudwatchPutMetricData, "put-metric-data", "", false, "Put Metric Data")
	_cloudwatchCmd.Flags().BoolVarP(&_cloudwatchPutMetricStream, "put-metric-stream", "", false, "Put Metric Stream")
	_cloudwatchCmd.Flags().BoolVarP(&_cloudwatchSetAlarmState, "set-alarm-state", "", false, "Set Alarm State")
	_cloudwatchCmd.Flags().BoolVarP(&_cloudwatchStartMetricStreams, "start-metric-streams", "", false, "Start Metric Streams")
	_cloudwatchCmd.Flags().BoolVarP(&_cloudwatchStopMetricStreams, "stop-metric-streams", "", false, "Stop Metric Streams")
	_cloudwatchCmd.Flags().BoolVarP(&_cloudwatchTagResource, "tag-resource", "", false, "Tag Resource")
	_cloudwatchCmd.Flags().BoolVarP(&_cloudwatchUntagResource, "untag-resource", "", false, "Untag Resource")

}
