package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/pi"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// piCmd represents the pi command
var _piCmd = &cobra.Command{
	Use:   "pi",
	Short: "AWS pi CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := pi.NewFromConfig(cfg)
		if _piCreatePerformanceAnalysisReport {
			pi_CreatePerformanceAnalysisReport(cfg, client)
			return
		}
		if _piDeletePerformanceAnalysisReport {
			pi_DeletePerformanceAnalysisReport(cfg, client)
			return
		}
		if _piDescribeDimensionKeys {
			pi_DescribeDimensionKeys(cfg, client)
			return
		}
		if _piGetDimensionKeyDetails {
			pi_GetDimensionKeyDetails(cfg, client)
			return
		}
		if _piGetPerformanceAnalysisReport {
			pi_GetPerformanceAnalysisReport(cfg, client)
			return
		}
		if _piGetResourceMetadata {
			pi_GetResourceMetadata(cfg, client)
			return
		}
		if _piGetResourceMetrics {
			pi_GetResourceMetrics(cfg, client)
			return
		}
		if _piListAvailableResourceDimensions {
			pi_ListAvailableResourceDimensions(cfg, client)
			return
		}
		if _piListAvailableResourceMetrics {
			pi_ListAvailableResourceMetrics(cfg, client)
			return
		}
		if _piListPerformanceAnalysisReports {
			pi_ListPerformanceAnalysisReports(cfg, client)
			return
		}
		if _piListTagsForResource {
			pi_ListTagsForResource(cfg, client)
			return
		}
		if _piTagResource {
			pi_TagResource(cfg, client)
			return
		}
		if _piUntagResource {
			pi_UntagResource(cfg, client)
			return
		}

	},
}

var (
	_piCreatePerformanceAnalysisReport bool
	_piDeletePerformanceAnalysisReport bool
	_piDescribeDimensionKeys           bool
	_piGetDimensionKeyDetails          bool
	_piGetPerformanceAnalysisReport    bool
	_piGetResourceMetadata             bool
	_piGetResourceMetrics              bool
	_piListAvailableResourceDimensions bool
	_piListAvailableResourceMetrics    bool
	_piListPerformanceAnalysisReports  bool
	_piListTagsForResource             bool
	_piTagResource                     bool
	_piUntagResource                   bool

	_piAcceptLanguage      string
	_piAdditionalMetrics   []string
	_piAnalysisReportId    string
	_piAuthorizedActions   string
	_piEndTime             string
	_piFilter              string
	_piGroup               string
	_piGroupBy             string
	_piGroupIdentifier     string
	_piIdentifier          string
	_piListTags            string
	_piMaxResults          string
	_piMetric              string
	_piMetricQueries       string
	_piMetricTypes         []string
	_piMetrics             []string
	_piNextToken           string
	_piPartitionBy         string
	_piPeriodAlignment     string
	_piPeriodInSeconds     string
	_piRequestedDimensions []string
	_piResourceARN         string
	_piServiceType         string
	_piStartTime           string
	_piTagKeys             []string
	_piTags                string
	_piTextFormat          string
)

// Creates a new performance analysis report for a specific time period for the DB
// instance.
func pi_CreatePerformanceAnalysisReport(cfg aws.Config, client *pi.Client) {
	input := &pi.CreatePerformanceAnalysisReportInput{
		// EndTime: *time.Time, // Required
		// Identifier: *string, // Required
		// ServiceType: types.ServiceType, // Required
		// StartTime: *time.Time, // Required
	}

	if len(_piEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _piEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_piIdentifier) > 0 {
		input.Identifier = aws.String(_piIdentifier)
	}
	if len(_piServiceType) > 0 {
		if err := assignInputField(input, "ServiceType", _piServiceType); err != nil {
			log.Errorf("invalid --service-type: %s", err.Error())
			return
		}
	}
	if len(_piStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _piStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_piTags) > 0 {
		if err := assignInputField(input, "Tags", _piTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreatePerformanceAnalysisReport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a performance analysis report.
func pi_DeletePerformanceAnalysisReport(cfg aws.Config, client *pi.Client) {
	input := &pi.DeletePerformanceAnalysisReportInput{
		// AnalysisReportId: *string, // Required
		// Identifier: *string, // Required
		// ServiceType: types.ServiceType, // Required
	}

	if len(_piAnalysisReportId) > 0 {
		input.AnalysisReportId = aws.String(_piAnalysisReportId)
	}
	if len(_piIdentifier) > 0 {
		input.Identifier = aws.String(_piIdentifier)
	}
	if len(_piServiceType) > 0 {
		if err := assignInputField(input, "ServiceType", _piServiceType); err != nil {
			log.Errorf("invalid --service-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeletePerformanceAnalysisReport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// For a specific time period, retrieve the top N dimension keys for a metric.
// Each response element returns a maximum of 500 bytes. For larger elements, such
// as SQL statements, only the first 500 bytes are returned.
func pi_DescribeDimensionKeys(cfg aws.Config, client *pi.Client) {
	input := &pi.DescribeDimensionKeysInput{
		// EndTime: *time.Time, // Required
		// GroupBy: *types.DimensionGroup, // Required
		// Identifier: *string, // Required
		// Metric: *string, // Required
		// ServiceType: types.ServiceType, // Required
		// StartTime: *time.Time, // Required
	}

	if len(_piEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _piEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_piGroupBy) > 0 {
		if err := assignInputField(input, "GroupBy", _piGroupBy); err != nil {
			log.Errorf("invalid --group-by: %s", err.Error())
			return
		}
	}
	if len(_piIdentifier) > 0 {
		input.Identifier = aws.String(_piIdentifier)
	}
	if len(_piMetric) > 0 {
		input.Metric = aws.String(_piMetric)
	}
	if len(_piServiceType) > 0 {
		if err := assignInputField(input, "ServiceType", _piServiceType); err != nil {
			log.Errorf("invalid --service-type: %s", err.Error())
			return
		}
	}
	if len(_piStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _piStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_piAdditionalMetrics) > 0 {
		input.AdditionalMetrics = append([]string(nil), _piAdditionalMetrics...)
	}
	if len(_piFilter) > 0 {
		if err := assignInputField(input, "Filter", _piFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_piMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _piMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_piNextToken) > 0 {
		input.NextToken = aws.String(_piNextToken)
	}
	if len(_piPartitionBy) > 0 {
		if err := assignInputField(input, "PartitionBy", _piPartitionBy); err != nil {
			log.Errorf("invalid --partition-by: %s", err.Error())
			return
		}
	}
	if len(_piPeriodInSeconds) > 0 {
		if err := assignInputField(input, "PeriodInSeconds", _piPeriodInSeconds); err != nil {
			log.Errorf("invalid --period-in-seconds: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeDimensionKeys(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*pi.DescribeDimensionKeysOutput
	p := pi.NewDescribeDimensionKeysPaginator(client, input)
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

// Get the attributes of the specified dimension group for a DB instance or data
// source. For example, if you specify a SQL ID, GetDimensionKeyDetails retrieves
// the full text of the dimension db.sql.statement associated with this ID. This
// operation is useful because GetResourceMetrics and DescribeDimensionKeys don't
// support retrieval of large SQL statement text, lock snapshots, and execution
// plans.
func pi_GetDimensionKeyDetails(cfg aws.Config, client *pi.Client) {
	input := &pi.GetDimensionKeyDetailsInput{
		// Group: *string, // Required
		// GroupIdentifier: *string, // Required
		// Identifier: *string, // Required
		// ServiceType: types.ServiceType, // Required
	}

	if len(_piGroup) > 0 {
		input.Group = aws.String(_piGroup)
	}
	if len(_piGroupIdentifier) > 0 {
		input.GroupIdentifier = aws.String(_piGroupIdentifier)
	}
	if len(_piIdentifier) > 0 {
		input.Identifier = aws.String(_piIdentifier)
	}
	if len(_piServiceType) > 0 {
		if err := assignInputField(input, "ServiceType", _piServiceType); err != nil {
			log.Errorf("invalid --service-type: %s", err.Error())
			return
		}
	}
	if len(_piRequestedDimensions) > 0 {
		input.RequestedDimensions = append([]string(nil), _piRequestedDimensions...)
	}

	if resp, err := client.GetDimensionKeyDetails(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the report including the report ID, status, time details, and the
// insights with recommendations. The report status can be RUNNING , SUCCEEDED , or
// FAILED . The insights include the description and recommendation fields.
func pi_GetPerformanceAnalysisReport(cfg aws.Config, client *pi.Client) {
	input := &pi.GetPerformanceAnalysisReportInput{
		// AnalysisReportId: *string, // Required
		// Identifier: *string, // Required
		// ServiceType: types.ServiceType, // Required
	}

	if len(_piAnalysisReportId) > 0 {
		input.AnalysisReportId = aws.String(_piAnalysisReportId)
	}
	if len(_piIdentifier) > 0 {
		input.Identifier = aws.String(_piIdentifier)
	}
	if len(_piServiceType) > 0 {
		if err := assignInputField(input, "ServiceType", _piServiceType); err != nil {
			log.Errorf("invalid --service-type: %s", err.Error())
			return
		}
	}
	if len(_piAcceptLanguage) > 0 {
		if err := assignInputField(input, "AcceptLanguage", _piAcceptLanguage); err != nil {
			log.Errorf("invalid --accept-language: %s", err.Error())
			return
		}
	}
	if len(_piTextFormat) > 0 {
		if err := assignInputField(input, "TextFormat", _piTextFormat); err != nil {
			log.Errorf("invalid --text-format: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetPerformanceAnalysisReport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieve the metadata for different features. For example, the metadata might
// indicate that a feature is turned on or off on a specific DB instance.
func pi_GetResourceMetadata(cfg aws.Config, client *pi.Client) {
	input := &pi.GetResourceMetadataInput{
		// Identifier: *string, // Required
		// ServiceType: types.ServiceType, // Required
	}

	if len(_piIdentifier) > 0 {
		input.Identifier = aws.String(_piIdentifier)
	}
	if len(_piServiceType) > 0 {
		if err := assignInputField(input, "ServiceType", _piServiceType); err != nil {
			log.Errorf("invalid --service-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetResourceMetadata(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieve Performance Insights metrics for a set of data sources over a time
// period. You can provide specific dimension groups and dimensions, and provide
// filtering criteria for each group. You must specify an aggregate function for
// each metric.
//
// Each response element returns a maximum of 500 bytes. For larger elements, such
// as SQL statements, only the first 500 bytes are returned.
func pi_GetResourceMetrics(cfg aws.Config, client *pi.Client) {
	input := &pi.GetResourceMetricsInput{
		// EndTime: *time.Time, // Required
		// Identifier: *string, // Required
		// MetricQueries: []types.MetricQuery, // Required
		// ServiceType: types.ServiceType, // Required
		// StartTime: *time.Time, // Required
	}

	if len(_piEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _piEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_piIdentifier) > 0 {
		input.Identifier = aws.String(_piIdentifier)
	}
	if len(_piMetricQueries) > 0 {
		if err := assignInputField(input, "MetricQueries", _piMetricQueries); err != nil {
			log.Errorf("invalid --metric-queries: %s", err.Error())
			return
		}
	}
	if len(_piServiceType) > 0 {
		if err := assignInputField(input, "ServiceType", _piServiceType); err != nil {
			log.Errorf("invalid --service-type: %s", err.Error())
			return
		}
	}
	if len(_piStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _piStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_piMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _piMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_piNextToken) > 0 {
		input.NextToken = aws.String(_piNextToken)
	}
	if len(_piPeriodAlignment) > 0 {
		if err := assignInputField(input, "PeriodAlignment", _piPeriodAlignment); err != nil {
			log.Errorf("invalid --period-alignment: %s", err.Error())
			return
		}
	}
	if len(_piPeriodInSeconds) > 0 {
		if err := assignInputField(input, "PeriodInSeconds", _piPeriodInSeconds); err != nil {
			log.Errorf("invalid --period-in-seconds: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.GetResourceMetrics(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*pi.GetResourceMetricsOutput
	p := pi.NewGetResourceMetricsPaginator(client, input)
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

// Retrieve the dimensions that can be queried for each specified metric type on a
// specified DB instance.
func pi_ListAvailableResourceDimensions(cfg aws.Config, client *pi.Client) {
	input := &pi.ListAvailableResourceDimensionsInput{
		// Identifier: *string, // Required
		// Metrics: []string, // Required
		// ServiceType: types.ServiceType, // Required
	}

	if len(_piIdentifier) > 0 {
		input.Identifier = aws.String(_piIdentifier)
	}
	if len(_piMetrics) > 0 {
		input.Metrics = append([]string(nil), _piMetrics...)
	}
	if len(_piServiceType) > 0 {
		if err := assignInputField(input, "ServiceType", _piServiceType); err != nil {
			log.Errorf("invalid --service-type: %s", err.Error())
			return
		}
	}
	if len(_piAuthorizedActions) > 0 {
		if err := assignInputField(input, "AuthorizedActions", _piAuthorizedActions); err != nil {
			log.Errorf("invalid --authorized-actions: %s", err.Error())
			return
		}
	}
	if len(_piMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _piMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_piNextToken) > 0 {
		input.NextToken = aws.String(_piNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAvailableResourceDimensions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*pi.ListAvailableResourceDimensionsOutput
	p := pi.NewListAvailableResourceDimensionsPaginator(client, input)
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

// Retrieve metrics of the specified types that can be queried for a specified DB
// instance.
func pi_ListAvailableResourceMetrics(cfg aws.Config, client *pi.Client) {
	input := &pi.ListAvailableResourceMetricsInput{
		// Identifier: *string, // Required
		// MetricTypes: []string, // Required
		// ServiceType: types.ServiceType, // Required
	}

	if len(_piIdentifier) > 0 {
		input.Identifier = aws.String(_piIdentifier)
	}
	if len(_piMetricTypes) > 0 {
		input.MetricTypes = append([]string(nil), _piMetricTypes...)
	}
	if len(_piServiceType) > 0 {
		if err := assignInputField(input, "ServiceType", _piServiceType); err != nil {
			log.Errorf("invalid --service-type: %s", err.Error())
			return
		}
	}
	if len(_piMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _piMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_piNextToken) > 0 {
		input.NextToken = aws.String(_piNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAvailableResourceMetrics(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*pi.ListAvailableResourceMetricsOutput
	p := pi.NewListAvailableResourceMetricsPaginator(client, input)
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

// Lists all the analysis reports created for the DB instance. The reports are
// sorted based on the start time of each report.
func pi_ListPerformanceAnalysisReports(cfg aws.Config, client *pi.Client) {
	input := &pi.ListPerformanceAnalysisReportsInput{
		// Identifier: *string, // Required
		// ServiceType: types.ServiceType, // Required
	}

	if len(_piIdentifier) > 0 {
		input.Identifier = aws.String(_piIdentifier)
	}
	if len(_piServiceType) > 0 {
		if err := assignInputField(input, "ServiceType", _piServiceType); err != nil {
			log.Errorf("invalid --service-type: %s", err.Error())
			return
		}
	}
	if len(_piListTags) > 0 {
		if err := assignInputField(input, "ListTags", _piListTags); err != nil {
			log.Errorf("invalid --list-tags: %s", err.Error())
			return
		}
	}
	if len(_piMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _piMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_piNextToken) > 0 {
		input.NextToken = aws.String(_piNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPerformanceAnalysisReports(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*pi.ListPerformanceAnalysisReportsOutput
	p := pi.NewListPerformanceAnalysisReportsPaginator(client, input)
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

// Retrieves all the metadata tags associated with Amazon RDS Performance Insights
// resource.
func pi_ListTagsForResource(cfg aws.Config, client *pi.Client) {
	input := &pi.ListTagsForResourceInput{
		// ResourceARN: *string, // Required
		// ServiceType: types.ServiceType, // Required
	}

	if len(_piResourceARN) > 0 {
		input.ResourceARN = aws.String(_piResourceARN)
	}
	if len(_piServiceType) > 0 {
		if err := assignInputField(input, "ServiceType", _piServiceType); err != nil {
			log.Errorf("invalid --service-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds metadata tags to the Amazon RDS Performance Insights resource.
func pi_TagResource(cfg aws.Config, client *pi.Client) {
	input := &pi.TagResourceInput{
		// ResourceARN: *string, // Required
		// ServiceType: types.ServiceType, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_piResourceARN) > 0 {
		input.ResourceARN = aws.String(_piResourceARN)
	}
	if len(_piServiceType) > 0 {
		if err := assignInputField(input, "ServiceType", _piServiceType); err != nil {
			log.Errorf("invalid --service-type: %s", err.Error())
			return
		}
	}
	if len(_piTags) > 0 {
		if err := assignInputField(input, "Tags", _piTags); err != nil {
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

// Deletes the metadata tags from the Amazon RDS Performance Insights resource.
func pi_UntagResource(cfg aws.Config, client *pi.Client) {
	input := &pi.UntagResourceInput{
		// ResourceARN: *string, // Required
		// ServiceType: types.ServiceType, // Required
		// TagKeys: []string, // Required
	}

	if len(_piResourceARN) > 0 {
		input.ResourceARN = aws.String(_piResourceARN)
	}
	if len(_piServiceType) > 0 {
		if err := assignInputField(input, "ServiceType", _piServiceType); err != nil {
			log.Errorf("invalid --service-type: %s", err.Error())
			return
		}
	}
	if len(_piTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _piTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_piCmd)
	_piCmd.Flags().SortFlags = false

	_piCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_piCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_piCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_piCmd.Flags().StringVarP(&_piAcceptLanguage, "accept-language", "", "", "Accept Language")
	_piCmd.Flags().StringSliceVarP(&_piAdditionalMetrics, "additional-metrics", "", nil, "Additional Metrics")
	_piCmd.Flags().StringVarP(&_piAnalysisReportId, "analysis-report-id", "", "", "Analysis Report ID")
	_piCmd.Flags().StringVarP(&_piAuthorizedActions, "authorized-actions", "", "", "Authorized Actions")
	_piCmd.Flags().StringVarP(&_piEndTime, "end-time", "", "", "End Time")
	_piCmd.Flags().StringVarP(&_piFilter, "filter", "", "", "Filter")
	_piCmd.Flags().StringVarP(&_piGroup, "group", "", "", "Group")
	_piCmd.Flags().StringVarP(&_piGroupBy, "group-by", "", "", "Group By")
	_piCmd.Flags().StringVarP(&_piGroupIdentifier, "group-identifier", "", "", "Group Identifier")
	_piCmd.Flags().StringVarP(&_piIdentifier, "identifier", "", "", "Identifier")
	_piCmd.Flags().StringVarP(&_piListTags, "list-tags", "", "", "List Tags")
	_piCmd.Flags().StringVarP(&_piMaxResults, "max-results", "", "", "Max Results")
	_piCmd.Flags().StringVarP(&_piMetric, "metric", "", "", "Metric")
	_piCmd.Flags().StringVarP(&_piMetricQueries, "metric-queries", "", "", "Metric Queries")
	_piCmd.Flags().StringSliceVarP(&_piMetricTypes, "metric-types", "", nil, "Metric Types")
	_piCmd.Flags().StringSliceVarP(&_piMetrics, "metrics", "", nil, "Metrics")
	_piCmd.Flags().StringVarP(&_piNextToken, "next-token", "", "", "Next Token")
	_piCmd.Flags().StringVarP(&_piPartitionBy, "partition-by", "", "", "Partition By")
	_piCmd.Flags().StringVarP(&_piPeriodAlignment, "period-alignment", "", "", "Period Alignment")
	_piCmd.Flags().StringVarP(&_piPeriodInSeconds, "period-in-seconds", "", "", "Period In Seconds")
	_piCmd.Flags().StringSliceVarP(&_piRequestedDimensions, "requested-dimensions", "", nil, "Requested Dimensions")
	_piCmd.Flags().StringVarP(&_piResourceARN, "resource-arn", "", "", "Resource ARN")
	_piCmd.Flags().StringVarP(&_piServiceType, "service-type", "", "", "Service Type")
	_piCmd.Flags().StringVarP(&_piStartTime, "start-time", "", "", "Start Time")
	_piCmd.Flags().StringSliceVarP(&_piTagKeys, "tag-keys", "", nil, "Tag Keys")
	_piCmd.Flags().StringVarP(&_piTags, "tags", "", "", "Tags")
	_piCmd.Flags().StringVarP(&_piTextFormat, "text-format", "", "", "Text Format")

	_piCmd.Flags().BoolVarP(&_piCreatePerformanceAnalysisReport, "create-performance-analysis-report", "", false, "Create Performance Analysis Report")
	_piCmd.Flags().BoolVarP(&_piDeletePerformanceAnalysisReport, "delete-performance-analysis-report", "", false, "Delete Performance Analysis Report")
	_piCmd.Flags().BoolVarP(&_piDescribeDimensionKeys, "describe-dimension-keys", "", false, "Describe Dimension Keys")
	_piCmd.Flags().BoolVarP(&_piGetDimensionKeyDetails, "get-dimension-key-details", "", false, "Get Dimension Key Details")
	_piCmd.Flags().BoolVarP(&_piGetPerformanceAnalysisReport, "get-performance-analysis-report", "", false, "Get Performance Analysis Report")
	_piCmd.Flags().BoolVarP(&_piGetResourceMetadata, "get-resource-metadata", "", false, "Get Resource Metadata")
	_piCmd.Flags().BoolVarP(&_piGetResourceMetrics, "get-resource-metrics", "", false, "Get Resource Metrics")
	_piCmd.Flags().BoolVarP(&_piListAvailableResourceDimensions, "list-available-resource-dimensions", "", false, "List Available Resource Dimensions")
	_piCmd.Flags().BoolVarP(&_piListAvailableResourceMetrics, "list-available-resource-metrics", "", false, "List Available Resource Metrics")
	_piCmd.Flags().BoolVarP(&_piListPerformanceAnalysisReports, "list-performance-analysis-reports", "", false, "List Performance Analysis Reports")
	_piCmd.Flags().BoolVarP(&_piListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_piCmd.Flags().BoolVarP(&_piTagResource, "tag-resource", "", false, "Tag Resource")
	_piCmd.Flags().BoolVarP(&_piUntagResource, "untag-resource", "", false, "Untag Resource")

}
