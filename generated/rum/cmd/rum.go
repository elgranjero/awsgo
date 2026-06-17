package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rum"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// rumCmd represents the rum command
var _rumCmd = &cobra.Command{
	Use:   "rum",
	Short: "AWS rum CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := rum.NewFromConfig(cfg)
		if _rumBatchCreateRumMetricDefinitions {
			rum_BatchCreateRumMetricDefinitions(cfg, client)
			return
		}
		if _rumBatchDeleteRumMetricDefinitions {
			rum_BatchDeleteRumMetricDefinitions(cfg, client)
			return
		}
		if _rumBatchGetRumMetricDefinitions {
			rum_BatchGetRumMetricDefinitions(cfg, client)
			return
		}
		if _rumCreateAppMonitor {
			rum_CreateAppMonitor(cfg, client)
			return
		}
		if _rumDeleteAppMonitor {
			rum_DeleteAppMonitor(cfg, client)
			return
		}
		if _rumDeleteResourcePolicy {
			rum_DeleteResourcePolicy(cfg, client)
			return
		}
		if _rumDeleteRumMetricsDestination {
			rum_DeleteRumMetricsDestination(cfg, client)
			return
		}
		if _rumGetAppMonitor {
			rum_GetAppMonitor(cfg, client)
			return
		}
		if _rumGetAppMonitorData {
			rum_GetAppMonitorData(cfg, client)
			return
		}
		if _rumGetResourcePolicy {
			rum_GetResourcePolicy(cfg, client)
			return
		}
		if _rumListAppMonitors {
			rum_ListAppMonitors(cfg, client)
			return
		}
		if _rumListRumMetricsDestinations {
			rum_ListRumMetricsDestinations(cfg, client)
			return
		}
		if _rumListTagsForResource {
			rum_ListTagsForResource(cfg, client)
			return
		}
		if _rumPutResourcePolicy {
			rum_PutResourcePolicy(cfg, client)
			return
		}
		if _rumPutRumEvents {
			rum_PutRumEvents(cfg, client)
			return
		}
		if _rumPutRumMetricsDestination {
			rum_PutRumMetricsDestination(cfg, client)
			return
		}
		if _rumTagResource {
			rum_TagResource(cfg, client)
			return
		}
		if _rumUntagResource {
			rum_UntagResource(cfg, client)
			return
		}
		if _rumUpdateAppMonitor {
			rum_UpdateAppMonitor(cfg, client)
			return
		}
		if _rumUpdateRumMetricDefinition {
			rum_UpdateRumMetricDefinition(cfg, client)
			return
		}

	},
}

var (
	_rumBatchCreateRumMetricDefinitions bool
	_rumBatchDeleteRumMetricDefinitions bool
	_rumBatchGetRumMetricDefinitions    bool
	_rumCreateAppMonitor                bool
	_rumDeleteAppMonitor                bool
	_rumDeleteResourcePolicy            bool
	_rumDeleteRumMetricsDestination     bool
	_rumGetAppMonitor                   bool
	_rumGetAppMonitorData               bool
	_rumGetResourcePolicy               bool
	_rumListAppMonitors                 bool
	_rumListRumMetricsDestinations      bool
	_rumListTagsForResource             bool
	_rumPutResourcePolicy               bool
	_rumPutRumEvents                    bool
	_rumPutRumMetricsDestination        bool
	_rumTagResource                     bool
	_rumUntagResource                   bool
	_rumUpdateAppMonitor                bool
	_rumUpdateRumMetricDefinition       bool

	_rumAlias                      string
	_rumAppMonitorConfiguration    string
	_rumAppMonitorDetails          string
	_rumAppMonitorName             string
	_rumBatchId                    string
	_rumCustomEvents               string
	_rumCwLogEnabled               string
	_rumDeobfuscationConfiguration string
	_rumDestination                string
	_rumDestinationArn             string
	_rumDomain                     string
	_rumDomainList                 []string
	_rumFilters                    string
	_rumIamRoleArn                 string
	_rumId                         string
	_rumMaxResults                 string
	_rumMetricDefinition           string
	_rumMetricDefinitionId         string
	_rumMetricDefinitionIds        []string
	_rumMetricDefinitions          string
	_rumName                       string
	_rumNextToken                  string
	_rumPlatform                   string
	_rumPolicyDocument             string
	_rumPolicyRevisionId           string
	_rumResourceArn                string
	_rumRumEvents                  string
	_rumTagKeys                    []string
	_rumTags                       string
	_rumTimeRange                  string
	_rumUserDetails                string
)

// Specifies the extended metrics and custom metrics that you want a CloudWatch
// RUM app monitor to send to a destination. Valid destinations include CloudWatch
// and Evidently.
//
// By default, RUM app monitors send some metrics to CloudWatch. These default
// metrics are listed in [CloudWatch metrics that you can collect with CloudWatch RUM].
//
// In addition to these default metrics, you can choose to send extended metrics,
// custom metrics, or both.
//
// - Extended metrics let you send metrics with additional dimensions that
// aren't included in the default metrics. You can also send extended metrics to
// both Evidently and CloudWatch. The valid dimension names for the additional
// dimensions for extended metrics are BrowserName , CountryCode , DeviceType ,
// FileType , OSName , and PageId . For more information, see [Extended metrics that you can send to CloudWatch and CloudWatch Evidently].
//
// - Custom metrics are metrics that you define. You can send custom metrics to
// CloudWatch. CloudWatch Evidently, or both. With custom metrics, you can use any
// metric name and namespace. To derive the metrics, you can use any custom events,
// built-in events, custom attributes, or default attributes.
//
// You can't send custom metrics to the AWS/RUM namespace. You must send custom
//
// metrics to a custom namespace that you define. The namespace that you use can't
// start with AWS/ . CloudWatch RUM prepends RUM/CustomMetrics/ to the custom
// namespace that you define, so the final namespace for your metrics in CloudWatch
// is RUM/CustomMetrics/your-custom-namespace .
//
// The maximum number of metric definitions that you can specify in one
// BatchCreateRumMetricDefinitions operation is 200.
//
// The maximum number of metric definitions that one destination can contain is
// 2000.
//
// Extended metrics sent to CloudWatch and RUM custom metrics are charged as
// CloudWatch custom metrics. Each combination of additional dimension name and
// dimension value counts as a custom metric. For more information, see [Amazon CloudWatch Pricing].
//
// You must have already created a destination for the metrics before you send
// them. For more information, see [PutRumMetricsDestination].
//
// If some metric definitions specified in a BatchCreateRumMetricDefinitions
// operations are not valid, those metric definitions fail and return errors, but
// all valid metric definitions in the same operation still succeed.
//
// [PutRumMetricsDestination]: https://docs.aws.amazon.com/cloudwatchrum/latest/APIReference/API_PutRumMetricsDestination.html
// [Extended metrics that you can send to CloudWatch and CloudWatch Evidently]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-RUM-vended-metrics.html
// [CloudWatch metrics that you can collect with CloudWatch RUM]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-RUM-metrics.html
// [Amazon CloudWatch Pricing]: https://aws.amazon.com/cloudwatch/pricing/
func rum_BatchCreateRumMetricDefinitions(cfg aws.Config, client *rum.Client) {
	input := &rum.BatchCreateRumMetricDefinitionsInput{
		// AppMonitorName: *string, // Required
		// Destination: types.MetricDestination, // Required
		// MetricDefinitions: []types.MetricDefinitionRequest, // Required
	}

	if len(_rumAppMonitorName) > 0 {
		input.AppMonitorName = aws.String(_rumAppMonitorName)
	}
	if len(_rumDestination) > 0 {
		if err := assignInputField(input, "Destination", _rumDestination); err != nil {
			log.Errorf("invalid --destination: %s", err.Error())
			return
		}
	}
	if len(_rumMetricDefinitions) > 0 {
		if err := assignInputField(input, "MetricDefinitions", _rumMetricDefinitions); err != nil {
			log.Errorf("invalid --metric-definitions: %s", err.Error())
			return
		}
	}
	if len(_rumDestinationArn) > 0 {
		input.DestinationArn = aws.String(_rumDestinationArn)
	}

	if resp, err := client.BatchCreateRumMetricDefinitions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the specified metrics from being sent to an extended metrics
// destination.
//
// If some metric definition IDs specified in a BatchDeleteRumMetricDefinitions
// operations are not valid, those metric definitions fail and return errors, but
// all valid metric definition IDs in the same operation are still deleted.
//
// The maximum number of metric definitions that you can specify in one
// BatchDeleteRumMetricDefinitions operation is 200.
func rum_BatchDeleteRumMetricDefinitions(cfg aws.Config, client *rum.Client) {
	input := &rum.BatchDeleteRumMetricDefinitionsInput{
		// AppMonitorName: *string, // Required
		// Destination: types.MetricDestination, // Required
		// MetricDefinitionIds: []string, // Required
	}

	if len(_rumAppMonitorName) > 0 {
		input.AppMonitorName = aws.String(_rumAppMonitorName)
	}
	if len(_rumDestination) > 0 {
		if err := assignInputField(input, "Destination", _rumDestination); err != nil {
			log.Errorf("invalid --destination: %s", err.Error())
			return
		}
	}
	if len(_rumMetricDefinitionIds) > 0 {
		input.MetricDefinitionIds = append([]string(nil), _rumMetricDefinitionIds...)
	}
	if len(_rumDestinationArn) > 0 {
		input.DestinationArn = aws.String(_rumDestinationArn)
	}

	if resp, err := client.BatchDeleteRumMetricDefinitions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the list of metrics and dimensions that a RUM app monitor is sending
// to a single destination.
func rum_BatchGetRumMetricDefinitions(cfg aws.Config, client *rum.Client) {
	input := &rum.BatchGetRumMetricDefinitionsInput{
		// AppMonitorName: *string, // Required
		// Destination: types.MetricDestination, // Required
	}

	if len(_rumAppMonitorName) > 0 {
		input.AppMonitorName = aws.String(_rumAppMonitorName)
	}
	if len(_rumDestination) > 0 {
		if err := assignInputField(input, "Destination", _rumDestination); err != nil {
			log.Errorf("invalid --destination: %s", err.Error())
			return
		}
	}
	if len(_rumDestinationArn) > 0 {
		input.DestinationArn = aws.String(_rumDestinationArn)
	}
	if len(_rumMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _rumMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_rumNextToken) > 0 {
		input.NextToken = aws.String(_rumNextToken)
	}

	if disablePaginator() {
		if resp, err := client.BatchGetRumMetricDefinitions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rum.BatchGetRumMetricDefinitionsOutput
	p := rum.NewBatchGetRumMetricDefinitionsPaginator(client, input)
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

// Creates a Amazon CloudWatch RUM app monitor, which collects telemetry data from
// your application and sends that data to RUM. The data includes performance and
// reliability information such as page load time, client-side errors, and user
// behavior.
//
// You use this operation only to create a new app monitor. To update an existing
// app monitor, use [UpdateAppMonitor]instead.
//
// After you create an app monitor, sign in to the CloudWatch RUM console to get
// the JavaScript code snippet to add to your web application. For more
// information, see [How do I find a code snippet that I've already generated?]
//
// [UpdateAppMonitor]: https://docs.aws.amazon.com/cloudwatchrum/latest/APIReference/API_UpdateAppMonitor.html
// [How do I find a code snippet that I've already generated?]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-RUM-find-code-snippet.html
func rum_CreateAppMonitor(cfg aws.Config, client *rum.Client) {
	input := &rum.CreateAppMonitorInput{
		// Name: *string, // Required
	}

	if len(_rumName) > 0 {
		input.Name = aws.String(_rumName)
	}
	if len(_rumAppMonitorConfiguration) > 0 {
		if err := assignInputField(input, "AppMonitorConfiguration", _rumAppMonitorConfiguration); err != nil {
			log.Errorf("invalid --app-monitor-configuration: %s", err.Error())
			return
		}
	}
	if len(_rumCustomEvents) > 0 {
		if err := assignInputField(input, "CustomEvents", _rumCustomEvents); err != nil {
			log.Errorf("invalid --custom-events: %s", err.Error())
			return
		}
	}
	if len(_rumCwLogEnabled) > 0 {
		if err := assignInputField(input, "CwLogEnabled", _rumCwLogEnabled); err != nil {
			log.Errorf("invalid --cw-log-enabled: %s", err.Error())
			return
		}
	}
	if len(_rumDeobfuscationConfiguration) > 0 {
		if err := assignInputField(input, "DeobfuscationConfiguration", _rumDeobfuscationConfiguration); err != nil {
			log.Errorf("invalid --deobfuscation-configuration: %s", err.Error())
			return
		}
	}
	if len(_rumDomain) > 0 {
		input.Domain = aws.String(_rumDomain)
	}
	if len(_rumDomainList) > 0 {
		input.DomainList = append([]string(nil), _rumDomainList...)
	}
	if len(_rumPlatform) > 0 {
		if err := assignInputField(input, "Platform", _rumPlatform); err != nil {
			log.Errorf("invalid --platform: %s", err.Error())
			return
		}
	}
	if len(_rumTags) > 0 {
		if err := assignInputField(input, "Tags", _rumTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAppMonitor(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing app monitor. This immediately stops the collection of data.
func rum_DeleteAppMonitor(cfg aws.Config, client *rum.Client) {
	input := &rum.DeleteAppMonitorInput{
		// Name: *string, // Required
	}

	if len(_rumName) > 0 {
		input.Name = aws.String(_rumName)
	}

	if resp, err := client.DeleteAppMonitor(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the association of a resource-based policy from an app monitor.
func rum_DeleteResourcePolicy(cfg aws.Config, client *rum.Client) {
	input := &rum.DeleteResourcePolicyInput{
		// Name: *string, // Required
	}

	if len(_rumName) > 0 {
		input.Name = aws.String(_rumName)
	}
	if len(_rumPolicyRevisionId) > 0 {
		input.PolicyRevisionId = aws.String(_rumPolicyRevisionId)
	}

	if resp, err := client.DeleteResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a destination for CloudWatch RUM extended metrics, so that the
// specified app monitor stops sending extended metrics to that destination.
func rum_DeleteRumMetricsDestination(cfg aws.Config, client *rum.Client) {
	input := &rum.DeleteRumMetricsDestinationInput{
		// AppMonitorName: *string, // Required
		// Destination: types.MetricDestination, // Required
	}

	if len(_rumAppMonitorName) > 0 {
		input.AppMonitorName = aws.String(_rumAppMonitorName)
	}
	if len(_rumDestination) > 0 {
		if err := assignInputField(input, "Destination", _rumDestination); err != nil {
			log.Errorf("invalid --destination: %s", err.Error())
			return
		}
	}
	if len(_rumDestinationArn) > 0 {
		input.DestinationArn = aws.String(_rumDestinationArn)
	}

	if resp, err := client.DeleteRumMetricsDestination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the complete configuration information for one app monitor.
func rum_GetAppMonitor(cfg aws.Config, client *rum.Client) {
	input := &rum.GetAppMonitorInput{
		// Name: *string, // Required
	}

	if len(_rumName) > 0 {
		input.Name = aws.String(_rumName)
	}

	if resp, err := client.GetAppMonitor(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the raw performance events that RUM has collected from your web
// application, so that you can do your own processing or analysis of this data.
func rum_GetAppMonitorData(cfg aws.Config, client *rum.Client) {
	input := &rum.GetAppMonitorDataInput{
		// Name: *string, // Required
		// TimeRange: *types.TimeRange, // Required
	}

	if len(_rumName) > 0 {
		input.Name = aws.String(_rumName)
	}
	if len(_rumTimeRange) > 0 {
		if err := assignInputField(input, "TimeRange", _rumTimeRange); err != nil {
			log.Errorf("invalid --time-range: %s", err.Error())
			return
		}
	}
	if len(_rumFilters) > 0 {
		if err := assignInputField(input, "Filters", _rumFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_rumMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _rumMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_rumNextToken) > 0 {
		input.NextToken = aws.String(_rumNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetAppMonitorData(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rum.GetAppMonitorDataOutput
	p := rum.NewGetAppMonitorDataPaginator(client, input)
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

// Use this operation to retrieve information about a resource-based policy that
// is attached to an app monitor.
func rum_GetResourcePolicy(cfg aws.Config, client *rum.Client) {
	input := &rum.GetResourcePolicyInput{
		// Name: *string, // Required
	}

	if len(_rumName) > 0 {
		input.Name = aws.String(_rumName)
	}

	if resp, err := client.GetResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of the Amazon CloudWatch RUM app monitors in the account.
func rum_ListAppMonitors(cfg aws.Config, client *rum.Client) {
	input := &rum.ListAppMonitorsInput{}

	if len(_rumMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _rumMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_rumNextToken) > 0 {
		input.NextToken = aws.String(_rumNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAppMonitors(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rum.ListAppMonitorsOutput
	p := rum.NewListAppMonitorsPaginator(client, input)
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

// Returns a list of destinations that you have created to receive RUM extended
// metrics, for the specified app monitor.
//
// For more information about extended metrics, see [AddRumMetrics].
//
// [AddRumMetrics]: https://docs.aws.amazon.com/cloudwatchrum/latest/APIReference/API_AddRumMetrcs.html
func rum_ListRumMetricsDestinations(cfg aws.Config, client *rum.Client) {
	input := &rum.ListRumMetricsDestinationsInput{
		// AppMonitorName: *string, // Required
	}

	if len(_rumAppMonitorName) > 0 {
		input.AppMonitorName = aws.String(_rumAppMonitorName)
	}
	if len(_rumMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _rumMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_rumNextToken) > 0 {
		input.NextToken = aws.String(_rumNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRumMetricsDestinations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rum.ListRumMetricsDestinationsOutput
	p := rum.NewListRumMetricsDestinationsPaginator(client, input)
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

// Displays the tags associated with a CloudWatch RUM resource.
func rum_ListTagsForResource(cfg aws.Config, client *rum.Client) {
	input := &rum.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_rumResourceArn) > 0 {
		input.ResourceArn = aws.String(_rumResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Use this operation to assign a resource-based policy to a CloudWatch RUM app
// monitor to control access to it. Each app monitor can have one resource-based
// policy. The maximum size of the policy is 4 KB. To learn more about using
// resource policies with RUM, see [Using resource-based policies with CloudWatch RUM].
//
// [Using resource-based policies with CloudWatch RUM]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-RUM-resource-policies.html
func rum_PutResourcePolicy(cfg aws.Config, client *rum.Client) {
	input := &rum.PutResourcePolicyInput{
		// Name: *string, // Required
		// PolicyDocument: *string, // Required
	}

	if len(_rumName) > 0 {
		input.Name = aws.String(_rumName)
	}
	if len(_rumPolicyDocument) > 0 {
		input.PolicyDocument = aws.String(_rumPolicyDocument)
	}
	if len(_rumPolicyRevisionId) > 0 {
		input.PolicyRevisionId = aws.String(_rumPolicyRevisionId)
	}

	if resp, err := client.PutResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sends telemetry events about your application performance and user behavior to
// CloudWatch RUM. The code snippet that RUM generates for you to add to your
// application includes PutRumEvents operations to send this data to RUM.
//
// Each PutRumEvents operation can send a batch of events from one user session.
func rum_PutRumEvents(cfg aws.Config, client *rum.Client) {
	input := &rum.PutRumEventsInput{
		// AppMonitorDetails: *types.AppMonitorDetails, // Required
		// BatchId: *string, // Required
		// Id: *string, // Required
		// RumEvents: []types.RumEvent, // Required
		// UserDetails: *types.UserDetails, // Required
	}

	if len(_rumAppMonitorDetails) > 0 {
		if err := assignInputField(input, "AppMonitorDetails", _rumAppMonitorDetails); err != nil {
			log.Errorf("invalid --app-monitor-details: %s", err.Error())
			return
		}
	}
	if len(_rumBatchId) > 0 {
		input.BatchId = aws.String(_rumBatchId)
	}
	if len(_rumId) > 0 {
		input.Id = aws.String(_rumId)
	}
	if len(_rumRumEvents) > 0 {
		if err := assignInputField(input, "RumEvents", _rumRumEvents); err != nil {
			log.Errorf("invalid --rum-events: %s", err.Error())
			return
		}
	}
	if len(_rumUserDetails) > 0 {
		if err := assignInputField(input, "UserDetails", _rumUserDetails); err != nil {
			log.Errorf("invalid --user-details: %s", err.Error())
			return
		}
	}
	if len(_rumAlias) > 0 {
		input.Alias = aws.String(_rumAlias)
	}

	if resp, err := client.PutRumEvents(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates a destination to receive extended metrics from CloudWatch
// RUM. You can send extended metrics to CloudWatch or to a CloudWatch Evidently
// experiment.
//
// For more information about extended metrics, see [BatchCreateRumMetricDefinitions].
//
// [BatchCreateRumMetricDefinitions]: https://docs.aws.amazon.com/cloudwatchrum/latest/APIReference/API_BatchCreateRumMetricDefinitions.html
func rum_PutRumMetricsDestination(cfg aws.Config, client *rum.Client) {
	input := &rum.PutRumMetricsDestinationInput{
		// AppMonitorName: *string, // Required
		// Destination: types.MetricDestination, // Required
	}

	if len(_rumAppMonitorName) > 0 {
		input.AppMonitorName = aws.String(_rumAppMonitorName)
	}
	if len(_rumDestination) > 0 {
		if err := assignInputField(input, "Destination", _rumDestination); err != nil {
			log.Errorf("invalid --destination: %s", err.Error())
			return
		}
	}
	if len(_rumDestinationArn) > 0 {
		input.DestinationArn = aws.String(_rumDestinationArn)
	}
	if len(_rumIamRoleArn) > 0 {
		input.IamRoleArn = aws.String(_rumIamRoleArn)
	}

	if resp, err := client.PutRumMetricsDestination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Assigns one or more tags (key-value pairs) to the specified CloudWatch RUM
// resource. Currently, the only resources that can be tagged app monitors.
//
// Tags can help you organize and categorize your resources. You can also use them
// to scope user permissions by granting a user permission to access or change only
// resources with certain tag values.
//
// Tags don't have any semantic meaning to Amazon Web Services and are interpreted
// strictly as strings of characters.
//
// You can use the TagResource action with a resource that already has tags. If
// you specify a new tag key for the resource, this tag is appended to the list of
// tags associated with the alarm. If you specify a tag key that is already
// associated with the resource, the new tag value that you specify replaces the
// previous value for that tag.
//
// You can associate as many as 50 tags with a resource.
//
// For more information, see [Tagging Amazon Web Services resources].
//
// [Tagging Amazon Web Services resources]: https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html
func rum_TagResource(cfg aws.Config, client *rum.Client) {
	input := &rum.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_rumResourceArn) > 0 {
		input.ResourceArn = aws.String(_rumResourceArn)
	}
	if len(_rumTags) > 0 {
		if err := assignInputField(input, "Tags", _rumTags); err != nil {
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
func rum_UntagResource(cfg aws.Config, client *rum.Client) {
	input := &rum.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_rumResourceArn) > 0 {
		input.ResourceArn = aws.String(_rumResourceArn)
	}
	if len(_rumTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _rumTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration of an existing app monitor. When you use this
// operation, only the parts of the app monitor configuration that you specify in
// this operation are changed. For any parameters that you omit, the existing
// values are kept.
//
// You can't use this operation to change the tags of an existing app monitor. To
// change the tags of an existing app monitor, use [TagResource].
//
// To create a new app monitor, use [CreateAppMonitor].
//
// After you update an app monitor, sign in to the CloudWatch RUM console to get
// the updated JavaScript code snippet to add to your web application. For more
// information, see [How do I find a code snippet that I've already generated?]
//
// [CreateAppMonitor]: https://docs.aws.amazon.com/cloudwatchrum/latest/APIReference/API_CreateAppMonitor.html
// [TagResource]: https://docs.aws.amazon.com/cloudwatchrum/latest/APIReference/API_TagResource.html
// [How do I find a code snippet that I've already generated?]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-RUM-find-code-snippet.html
func rum_UpdateAppMonitor(cfg aws.Config, client *rum.Client) {
	input := &rum.UpdateAppMonitorInput{
		// Name: *string, // Required
	}

	if len(_rumName) > 0 {
		input.Name = aws.String(_rumName)
	}
	if len(_rumAppMonitorConfiguration) > 0 {
		if err := assignInputField(input, "AppMonitorConfiguration", _rumAppMonitorConfiguration); err != nil {
			log.Errorf("invalid --app-monitor-configuration: %s", err.Error())
			return
		}
	}
	if len(_rumCustomEvents) > 0 {
		if err := assignInputField(input, "CustomEvents", _rumCustomEvents); err != nil {
			log.Errorf("invalid --custom-events: %s", err.Error())
			return
		}
	}
	if len(_rumCwLogEnabled) > 0 {
		if err := assignInputField(input, "CwLogEnabled", _rumCwLogEnabled); err != nil {
			log.Errorf("invalid --cw-log-enabled: %s", err.Error())
			return
		}
	}
	if len(_rumDeobfuscationConfiguration) > 0 {
		if err := assignInputField(input, "DeobfuscationConfiguration", _rumDeobfuscationConfiguration); err != nil {
			log.Errorf("invalid --deobfuscation-configuration: %s", err.Error())
			return
		}
	}
	if len(_rumDomain) > 0 {
		input.Domain = aws.String(_rumDomain)
	}
	if len(_rumDomainList) > 0 {
		input.DomainList = append([]string(nil), _rumDomainList...)
	}

	if resp, err := client.UpdateAppMonitor(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies one existing metric definition for CloudWatch RUM extended metrics.
// For more information about extended metrics, see [BatchCreateRumMetricsDefinitions].
//
// [BatchCreateRumMetricsDefinitions]: https://docs.aws.amazon.com/cloudwatchrum/latest/APIReference/API_BatchCreateRumMetricsDefinitions.html
func rum_UpdateRumMetricDefinition(cfg aws.Config, client *rum.Client) {
	input := &rum.UpdateRumMetricDefinitionInput{
		// AppMonitorName: *string, // Required
		// Destination: types.MetricDestination, // Required
		// MetricDefinition: *types.MetricDefinitionRequest, // Required
		// MetricDefinitionId: *string, // Required
	}

	if len(_rumAppMonitorName) > 0 {
		input.AppMonitorName = aws.String(_rumAppMonitorName)
	}
	if len(_rumDestination) > 0 {
		if err := assignInputField(input, "Destination", _rumDestination); err != nil {
			log.Errorf("invalid --destination: %s", err.Error())
			return
		}
	}
	if len(_rumMetricDefinition) > 0 {
		if err := assignInputField(input, "MetricDefinition", _rumMetricDefinition); err != nil {
			log.Errorf("invalid --metric-definition: %s", err.Error())
			return
		}
	}
	if len(_rumMetricDefinitionId) > 0 {
		input.MetricDefinitionId = aws.String(_rumMetricDefinitionId)
	}
	if len(_rumDestinationArn) > 0 {
		input.DestinationArn = aws.String(_rumDestinationArn)
	}

	if resp, err := client.UpdateRumMetricDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_rumCmd)
	_rumCmd.Flags().SortFlags = false

	_rumCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_rumCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_rumCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_rumCmd.Flags().StringVarP(&_rumAlias, "alias", "", "", "Alias")
	_rumCmd.Flags().StringVarP(&_rumAppMonitorConfiguration, "app-monitor-configuration", "", "", "App Monitor Configuration")
	_rumCmd.Flags().StringVarP(&_rumAppMonitorDetails, "app-monitor-details", "", "", "App Monitor Details")
	_rumCmd.Flags().StringVarP(&_rumAppMonitorName, "app-monitor-name", "", "", "App Monitor Name")
	_rumCmd.Flags().StringVarP(&_rumBatchId, "batch-id", "", "", "Batch ID")
	_rumCmd.Flags().StringVarP(&_rumCustomEvents, "custom-events", "", "", "Custom Events")
	_rumCmd.Flags().StringVarP(&_rumCwLogEnabled, "cw-log-enabled", "", "", "Cw Log Enabled")
	_rumCmd.Flags().StringVarP(&_rumDeobfuscationConfiguration, "deobfuscation-configuration", "", "", "Deobfuscation Configuration")
	_rumCmd.Flags().StringVarP(&_rumDestination, "destination", "", "", "Destination")
	_rumCmd.Flags().StringVarP(&_rumDestinationArn, "destination-arn", "", "", "Destination ARN")
	_rumCmd.Flags().StringVarP(&_rumDomain, "domain", "", "", "Domain")
	_rumCmd.Flags().StringSliceVarP(&_rumDomainList, "domain-list", "", nil, "Domain List")
	_rumCmd.Flags().StringVarP(&_rumFilters, "filters", "", "", "Filters")
	_rumCmd.Flags().StringVarP(&_rumIamRoleArn, "iam-role-arn", "", "", "IAM Role ARN")
	_rumCmd.Flags().StringVarP(&_rumId, "id", "", "", "ID")
	_rumCmd.Flags().StringVarP(&_rumMaxResults, "max-results", "", "", "Max Results")
	_rumCmd.Flags().StringVarP(&_rumMetricDefinition, "metric-definition", "", "", "Metric Definition")
	_rumCmd.Flags().StringVarP(&_rumMetricDefinitionId, "metric-definition-id", "", "", "Metric Definition ID")
	_rumCmd.Flags().StringSliceVarP(&_rumMetricDefinitionIds, "metric-definition-ids", "", nil, "Metric Definition Ids")
	_rumCmd.Flags().StringVarP(&_rumMetricDefinitions, "metric-definitions", "", "", "Metric Definitions")
	_rumCmd.Flags().StringVarP(&_rumName, "name", "", "", "Name")
	_rumCmd.Flags().StringVarP(&_rumNextToken, "next-token", "", "", "Next Token")
	_rumCmd.Flags().StringVarP(&_rumPlatform, "platform", "", "", "Platform")
	_rumCmd.Flags().StringVarP(&_rumPolicyDocument, "policy-document", "", "", "Policy Document")
	_rumCmd.Flags().StringVarP(&_rumPolicyRevisionId, "policy-revision-id", "", "", "Policy Revision ID")
	_rumCmd.Flags().StringVarP(&_rumResourceArn, "resource-arn", "", "", "Resource ARN")
	_rumCmd.Flags().StringVarP(&_rumRumEvents, "rum-events", "", "", "Rum Events")
	_rumCmd.Flags().StringSliceVarP(&_rumTagKeys, "tag-keys", "", nil, "Tag Keys")
	_rumCmd.Flags().StringVarP(&_rumTags, "tags", "", "", "Tags")
	_rumCmd.Flags().StringVarP(&_rumTimeRange, "time-range", "", "", "Time Range")
	_rumCmd.Flags().StringVarP(&_rumUserDetails, "user-details", "", "", "User Details")

	_rumCmd.Flags().BoolVarP(&_rumBatchCreateRumMetricDefinitions, "batch-create-rum-metric-definitions", "", false, "Batch Create Rum Metric Definitions")
	_rumCmd.Flags().BoolVarP(&_rumBatchDeleteRumMetricDefinitions, "batch-delete-rum-metric-definitions", "", false, "Batch Delete Rum Metric Definitions")
	_rumCmd.Flags().BoolVarP(&_rumBatchGetRumMetricDefinitions, "batch-get-rum-metric-definitions", "", false, "Batch Get Rum Metric Definitions")
	_rumCmd.Flags().BoolVarP(&_rumCreateAppMonitor, "create-app-monitor", "", false, "Create App Monitor")
	_rumCmd.Flags().BoolVarP(&_rumDeleteAppMonitor, "delete-app-monitor", "", false, "Delete App Monitor")
	_rumCmd.Flags().BoolVarP(&_rumDeleteResourcePolicy, "delete-resource-policy", "", false, "Delete Resource Policy")
	_rumCmd.Flags().BoolVarP(&_rumDeleteRumMetricsDestination, "delete-rum-metrics-destination", "", false, "Delete Rum Metrics Destination")
	_rumCmd.Flags().BoolVarP(&_rumGetAppMonitor, "get-app-monitor", "", false, "Get App Monitor")
	_rumCmd.Flags().BoolVarP(&_rumGetAppMonitorData, "get-app-monitor-data", "", false, "Get App Monitor Data")
	_rumCmd.Flags().BoolVarP(&_rumGetResourcePolicy, "get-resource-policy", "", false, "Get Resource Policy")
	_rumCmd.Flags().BoolVarP(&_rumListAppMonitors, "list-app-monitors", "", false, "List App Monitors")
	_rumCmd.Flags().BoolVarP(&_rumListRumMetricsDestinations, "list-rum-metrics-destinations", "", false, "List Rum Metrics Destinations")
	_rumCmd.Flags().BoolVarP(&_rumListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_rumCmd.Flags().BoolVarP(&_rumPutResourcePolicy, "put-resource-policy", "", false, "Put Resource Policy")
	_rumCmd.Flags().BoolVarP(&_rumPutRumEvents, "put-rum-events", "", false, "Put Rum Events")
	_rumCmd.Flags().BoolVarP(&_rumPutRumMetricsDestination, "put-rum-metrics-destination", "", false, "Put Rum Metrics Destination")
	_rumCmd.Flags().BoolVarP(&_rumTagResource, "tag-resource", "", false, "Tag Resource")
	_rumCmd.Flags().BoolVarP(&_rumUntagResource, "untag-resource", "", false, "Untag Resource")
	_rumCmd.Flags().BoolVarP(&_rumUpdateAppMonitor, "update-app-monitor", "", false, "Update App Monitor")
	_rumCmd.Flags().BoolVarP(&_rumUpdateRumMetricDefinition, "update-rum-metric-definition", "", false, "Update Rum Metric Definition")

}
