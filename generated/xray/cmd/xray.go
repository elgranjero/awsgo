package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/xray"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// xrayCmd represents the xray command
var _xrayCmd = &cobra.Command{
	Use:   "xray",
	Short: "AWS xray CLI",
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
		client := xray.NewFromConfig(cfg)
		if _xrayBatchGetTraces {
			xray_BatchGetTraces(cfg, client)
			return
		}
		if _xrayCancelTraceRetrieval {
			xray_CancelTraceRetrieval(cfg, client)
			return
		}
		if _xrayCreateGroup {
			xray_CreateGroup(cfg, client)
			return
		}
		if _xrayCreateSamplingRule {
			xray_CreateSamplingRule(cfg, client)
			return
		}
		if _xrayDeleteGroup {
			xray_DeleteGroup(cfg, client)
			return
		}
		if _xrayDeleteResourcePolicy {
			xray_DeleteResourcePolicy(cfg, client)
			return
		}
		if _xrayDeleteSamplingRule {
			xray_DeleteSamplingRule(cfg, client)
			return
		}
		if _xrayGetEncryptionConfig {
			xray_GetEncryptionConfig(cfg, client)
			return
		}
		if _xrayGetGroup {
			xray_GetGroup(cfg, client)
			return
		}
		if _xrayGetGroups {
			xray_GetGroups(cfg, client)
			return
		}
		if _xrayGetIndexingRules {
			xray_GetIndexingRules(cfg, client)
			return
		}
		if _xrayGetInsight {
			xray_GetInsight(cfg, client)
			return
		}
		if _xrayGetInsightEvents {
			xray_GetInsightEvents(cfg, client)
			return
		}
		if _xrayGetInsightImpactGraph {
			xray_GetInsightImpactGraph(cfg, client)
			return
		}
		if _xrayGetInsightSummaries {
			xray_GetInsightSummaries(cfg, client)
			return
		}
		if _xrayGetRetrievedTracesGraph {
			xray_GetRetrievedTracesGraph(cfg, client)
			return
		}
		if _xrayGetSamplingRules {
			xray_GetSamplingRules(cfg, client)
			return
		}
		if _xrayGetSamplingStatisticSummaries {
			xray_GetSamplingStatisticSummaries(cfg, client)
			return
		}
		if _xrayGetSamplingTargets {
			xray_GetSamplingTargets(cfg, client)
			return
		}
		if _xrayGetServiceGraph {
			xray_GetServiceGraph(cfg, client)
			return
		}
		if _xrayGetTimeSeriesServiceStatistics {
			xray_GetTimeSeriesServiceStatistics(cfg, client)
			return
		}
		if _xrayGetTraceGraph {
			xray_GetTraceGraph(cfg, client)
			return
		}
		if _xrayGetTraceSegmentDestination {
			xray_GetTraceSegmentDestination(cfg, client)
			return
		}
		if _xrayGetTraceSummaries {
			xray_GetTraceSummaries(cfg, client)
			return
		}
		if _xrayListResourcePolicies {
			xray_ListResourcePolicies(cfg, client)
			return
		}
		if _xrayListRetrievedTraces {
			xray_ListRetrievedTraces(cfg, client)
			return
		}
		if _xrayListTagsForResource {
			xray_ListTagsForResource(cfg, client)
			return
		}
		if _xrayPutEncryptionConfig {
			xray_PutEncryptionConfig(cfg, client)
			return
		}
		if _xrayPutResourcePolicy {
			xray_PutResourcePolicy(cfg, client)
			return
		}
		if _xrayPutTelemetryRecords {
			xray_PutTelemetryRecords(cfg, client)
			return
		}
		if _xrayPutTraceSegments {
			xray_PutTraceSegments(cfg, client)
			return
		}
		if _xrayStartTraceRetrieval {
			xray_StartTraceRetrieval(cfg, client)
			return
		}
		if _xrayTagResource {
			xray_TagResource(cfg, client)
			return
		}
		if _xrayUntagResource {
			xray_UntagResource(cfg, client)
			return
		}
		if _xrayUpdateGroup {
			xray_UpdateGroup(cfg, client)
			return
		}
		if _xrayUpdateIndexingRule {
			xray_UpdateIndexingRule(cfg, client)
			return
		}
		if _xrayUpdateSamplingRule {
			xray_UpdateSamplingRule(cfg, client)
			return
		}
		if _xrayUpdateTraceSegmentDestination {
			xray_UpdateTraceSegmentDestination(cfg, client)
			return
		}

	},
}

var (
	_xrayBatchGetTraces                 bool
	_xrayCancelTraceRetrieval           bool
	_xrayCreateGroup                    bool
	_xrayCreateSamplingRule             bool
	_xrayDeleteGroup                    bool
	_xrayDeleteResourcePolicy           bool
	_xrayDeleteSamplingRule             bool
	_xrayGetEncryptionConfig            bool
	_xrayGetGroup                       bool
	_xrayGetGroups                      bool
	_xrayGetIndexingRules               bool
	_xrayGetInsight                     bool
	_xrayGetInsightEvents               bool
	_xrayGetInsightImpactGraph          bool
	_xrayGetInsightSummaries            bool
	_xrayGetRetrievedTracesGraph        bool
	_xrayGetSamplingRules               bool
	_xrayGetSamplingStatisticSummaries  bool
	_xrayGetSamplingTargets             bool
	_xrayGetServiceGraph                bool
	_xrayGetTimeSeriesServiceStatistics bool
	_xrayGetTraceGraph                  bool
	_xrayGetTraceSegmentDestination     bool
	_xrayGetTraceSummaries              bool
	_xrayListResourcePolicies           bool
	_xrayListRetrievedTraces            bool
	_xrayListTagsForResource            bool
	_xrayPutEncryptionConfig            bool
	_xrayPutResourcePolicy              bool
	_xrayPutTelemetryRecords            bool
	_xrayPutTraceSegments               bool
	_xrayStartTraceRetrieval            bool
	_xrayTagResource                    bool
	_xrayUntagResource                  bool
	_xrayUpdateGroup                    bool
	_xrayUpdateIndexingRule             bool
	_xrayUpdateSamplingRule             bool
	_xrayUpdateTraceSegmentDestination  bool

	_xrayBypassPolicyLockoutCheck         string
	_xrayDestination                      string
	_xrayEC2InstanceId                    string
	_xrayEndTime                          string
	_xrayEntitySelectorExpression         string
	_xrayFilterExpression                 string
	_xrayForecastStatistics               string
	_xrayGroupARN                         string
	_xrayGroupName                        string
	_xrayHostname                         string
	_xrayInsightId                        string
	_xrayInsightsConfiguration            string
	_xrayKeyId                            string
	_xrayMaxResults                       string
	_xrayName                             string
	_xrayNextToken                        string
	_xrayPeriod                           string
	_xrayPolicyDocument                   string
	_xrayPolicyName                       string
	_xrayPolicyRevisionId                 string
	_xrayResourceARN                      string
	_xrayRetrievalToken                   string
	_xrayRule                             string
	_xrayRuleARN                          string
	_xrayRuleName                         string
	_xraySampling                         string
	_xraySamplingBoostStatisticsDocuments string
	_xraySamplingRule                     string
	_xraySamplingRuleUpdate               string
	_xraySamplingStatisticsDocuments      string
	_xraySamplingStrategy                 string
	_xrayStartTime                        string
	_xrayStates                           string
	_xrayTagKeys                          []string
	_xrayTags                             string
	_xrayTelemetryRecords                 string
	_xrayTimeRangeType                    string
	_xrayTraceFormat                      string
	_xrayTraceIds                         []string
	_xrayTraceSegmentDocuments            []string
	_xrayType                             string
)

// You cannot find traces through this API if Transaction Search is enabled since
// trace is not indexed in X-Ray.
//
// Retrieves a list of traces specified by ID. Each trace is a collection of
// segment documents that originates from a single request. Use GetTraceSummaries
// to get a list of trace IDs.
func xray_BatchGetTraces(cfg aws.Config, client *xray.Client) {
	input := &xray.BatchGetTracesInput{
		// TraceIds: []string, // Required
	}

	if len(_xrayTraceIds) > 0 {
		input.TraceIds = append([]string(nil), _xrayTraceIds...)
	}
	if len(_xrayNextToken) > 0 {
		input.NextToken = aws.String(_xrayNextToken)
	}

	if disablePaginator() {
		if resp, err := client.BatchGetTraces(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*xray.BatchGetTracesOutput
	p := xray.NewBatchGetTracesPaginator(client, input)
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

// Cancels an ongoing trace retrieval job initiated by StartTraceRetrieval using
// the provided RetrievalToken . A successful cancellation will return an HTTP 200
// response.
func xray_CancelTraceRetrieval(cfg aws.Config, client *xray.Client) {
	input := &xray.CancelTraceRetrievalInput{
		// RetrievalToken: *string, // Required
	}

	if len(_xrayRetrievalToken) > 0 {
		input.RetrievalToken = aws.String(_xrayRetrievalToken)
	}

	if resp, err := client.CancelTraceRetrieval(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a group resource with a name and a filter expression.
func xray_CreateGroup(cfg aws.Config, client *xray.Client) {
	input := &xray.CreateGroupInput{
		// GroupName: *string, // Required
	}

	if len(_xrayGroupName) > 0 {
		input.GroupName = aws.String(_xrayGroupName)
	}
	if len(_xrayFilterExpression) > 0 {
		input.FilterExpression = aws.String(_xrayFilterExpression)
	}
	if len(_xrayInsightsConfiguration) > 0 {
		if err := assignInputField(input, "InsightsConfiguration", _xrayInsightsConfiguration); err != nil {
			log.Errorf("invalid --insights-configuration: %s", err.Error())
			return
		}
	}
	if len(_xrayTags) > 0 {
		if err := assignInputField(input, "Tags", _xrayTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a rule to control sampling behavior for instrumented applications.
// Services retrieve rules with [GetSamplingRules], and evaluate each rule in ascending order of
// priority for each request. If a rule matches, the service records a trace,
// borrowing it from the reservoir size. After 10 seconds, the service reports back
// to X-Ray with [GetSamplingTargets]to get updated versions of each in-use rule. The updated rule
// contains a trace quota that the service can use instead of borrowing from the
// reservoir.
//
// [GetSamplingTargets]: https://docs.aws.amazon.com/xray/latest/api/API_GetSamplingTargets.html
// [GetSamplingRules]: https://docs.aws.amazon.com/xray/latest/api/API_GetSamplingRules.html
func xray_CreateSamplingRule(cfg aws.Config, client *xray.Client) {
	input := &xray.CreateSamplingRuleInput{
		// SamplingRule: *types.SamplingRule, // Required
	}

	if len(_xraySamplingRule) > 0 {
		if err := assignInputField(input, "SamplingRule", _xraySamplingRule); err != nil {
			log.Errorf("invalid --sampling-rule: %s", err.Error())
			return
		}
	}
	if len(_xrayTags) > 0 {
		if err := assignInputField(input, "Tags", _xrayTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSamplingRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a group resource.
func xray_DeleteGroup(cfg aws.Config, client *xray.Client) {
	input := &xray.DeleteGroupInput{}

	if len(_xrayGroupARN) > 0 {
		input.GroupARN = aws.String(_xrayGroupARN)
	}
	if len(_xrayGroupName) > 0 {
		input.GroupName = aws.String(_xrayGroupName)
	}

	if resp, err := client.DeleteGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a resource policy from the target Amazon Web Services account.
func xray_DeleteResourcePolicy(cfg aws.Config, client *xray.Client) {
	input := &xray.DeleteResourcePolicyInput{
		// PolicyName: *string, // Required
	}

	if len(_xrayPolicyName) > 0 {
		input.PolicyName = aws.String(_xrayPolicyName)
	}
	if len(_xrayPolicyRevisionId) > 0 {
		input.PolicyRevisionId = aws.String(_xrayPolicyRevisionId)
	}

	if resp, err := client.DeleteResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a sampling rule.
func xray_DeleteSamplingRule(cfg aws.Config, client *xray.Client) {
	input := &xray.DeleteSamplingRuleInput{}

	if len(_xrayRuleARN) > 0 {
		input.RuleARN = aws.String(_xrayRuleARN)
	}
	if len(_xrayRuleName) > 0 {
		input.RuleName = aws.String(_xrayRuleName)
	}

	if resp, err := client.DeleteSamplingRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the current encryption configuration for X-Ray data.
func xray_GetEncryptionConfig(cfg aws.Config, client *xray.Client) {
	input := &xray.GetEncryptionConfigInput{}

	if resp, err := client.GetEncryptionConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves group resource details.
func xray_GetGroup(cfg aws.Config, client *xray.Client) {
	input := &xray.GetGroupInput{}

	if len(_xrayGroupARN) > 0 {
		input.GroupARN = aws.String(_xrayGroupARN)
	}
	if len(_xrayGroupName) > 0 {
		input.GroupName = aws.String(_xrayGroupName)
	}

	if resp, err := client.GetGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves all active group details.
func xray_GetGroups(cfg aws.Config, client *xray.Client) {
	input := &xray.GetGroupsInput{}

	if len(_xrayNextToken) > 0 {
		input.NextToken = aws.String(_xrayNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*xray.GetGroupsOutput
	p := xray.NewGetGroupsPaginator(client, input)
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

// Retrieves all indexing rules.
// Indexing rules are used to determine the server-side sampling rate for spans
// ingested through the CloudWatchLogs destination and indexed by X-Ray. For more
// information, see [Transaction Search].
//
// [Transaction Search]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Transaction-Search.html
func xray_GetIndexingRules(cfg aws.Config, client *xray.Client) {
	input := &xray.GetIndexingRulesInput{}

	if len(_xrayNextToken) > 0 {
		input.NextToken = aws.String(_xrayNextToken)
	}

	if resp, err := client.GetIndexingRules(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the summary information of an insight. This includes impact to
// clients and root cause services, the top anomalous services, the category, the
// state of the insight, and the start and end time of the insight.
func xray_GetInsight(cfg aws.Config, client *xray.Client) {
	input := &xray.GetInsightInput{
		// InsightId: *string, // Required
	}

	if len(_xrayInsightId) > 0 {
		input.InsightId = aws.String(_xrayInsightId)
	}

	if resp, err := client.GetInsight(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// X-Ray reevaluates insights periodically until they're resolved, and records
// each intermediate state as an event. You can review an insight's events in the
// Impact Timeline on the Inspect page in the X-Ray console.
func xray_GetInsightEvents(cfg aws.Config, client *xray.Client) {
	input := &xray.GetInsightEventsInput{
		// InsightId: *string, // Required
	}

	if len(_xrayInsightId) > 0 {
		input.InsightId = aws.String(_xrayInsightId)
	}
	if len(_xrayMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _xrayMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_xrayNextToken) > 0 {
		input.NextToken = aws.String(_xrayNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetInsightEvents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*xray.GetInsightEventsOutput
	p := xray.NewGetInsightEventsPaginator(client, input)
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

// Retrieves a service graph structure filtered by the specified insight. The
// service graph is limited to only structural information. For a complete service
// graph, use this API with the GetServiceGraph API.
func xray_GetInsightImpactGraph(cfg aws.Config, client *xray.Client) {
	input := &xray.GetInsightImpactGraphInput{
		// EndTime: *time.Time, // Required
		// InsightId: *string, // Required
		// StartTime: *time.Time, // Required
	}

	if len(_xrayEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _xrayEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_xrayInsightId) > 0 {
		input.InsightId = aws.String(_xrayInsightId)
	}
	if len(_xrayStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _xrayStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_xrayNextToken) > 0 {
		input.NextToken = aws.String(_xrayNextToken)
	}

	if resp, err := client.GetInsightImpactGraph(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the summaries of all insights in the specified group matching the
// provided filter values.
func xray_GetInsightSummaries(cfg aws.Config, client *xray.Client) {
	input := &xray.GetInsightSummariesInput{
		// EndTime: *time.Time, // Required
		// StartTime: *time.Time, // Required
	}

	if len(_xrayEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _xrayEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_xrayStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _xrayStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_xrayGroupARN) > 0 {
		input.GroupARN = aws.String(_xrayGroupARN)
	}
	if len(_xrayGroupName) > 0 {
		input.GroupName = aws.String(_xrayGroupName)
	}
	if len(_xrayMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _xrayMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_xrayNextToken) > 0 {
		input.NextToken = aws.String(_xrayNextToken)
	}
	if len(_xrayStates) > 0 {
		if err := assignInputField(input, "States", _xrayStates); err != nil {
			log.Errorf("invalid --states: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.GetInsightSummaries(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*xray.GetInsightSummariesOutput
	p := xray.NewGetInsightSummariesPaginator(client, input)
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

// Retrieves a service graph for traces based on the specified RetrievalToken
// from the CloudWatch log group generated by Transaction Search. This API does not
// initiate a retrieval job. You must first execute StartTraceRetrieval to obtain
// the required RetrievalToken .
//
// The trace graph describes services that process incoming requests and any
// downstream services they call, which may include Amazon Web Services resources,
// external APIs, or databases.
//
// The response is empty until the RetrievalStatus is COMPLETE. Retry the request
// after the status changes from RUNNING or SCHEDULED to COMPLETE to access the
// full service graph.
//
// When CloudWatch log is the destination, this API can support cross-account
// observability and service graph retrieval across linked accounts.
//
// For retrieving graphs from X-Ray directly as opposed to the Transaction-Search
// Log group, see [GetTraceGraph].
//
// [GetTraceGraph]: https://docs.aws.amazon.com/xray/latest/api/API_GetTraceGraph.html
func xray_GetRetrievedTracesGraph(cfg aws.Config, client *xray.Client) {
	input := &xray.GetRetrievedTracesGraphInput{
		// RetrievalToken: *string, // Required
	}

	if len(_xrayRetrievalToken) > 0 {
		input.RetrievalToken = aws.String(_xrayRetrievalToken)
	}
	if len(_xrayNextToken) > 0 {
		input.NextToken = aws.String(_xrayNextToken)
	}

	if resp, err := client.GetRetrievedTracesGraph(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves all sampling rules.
func xray_GetSamplingRules(cfg aws.Config, client *xray.Client) {
	input := &xray.GetSamplingRulesInput{}

	if len(_xrayNextToken) > 0 {
		input.NextToken = aws.String(_xrayNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetSamplingRules(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*xray.GetSamplingRulesOutput
	p := xray.NewGetSamplingRulesPaginator(client, input)
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

// Retrieves information about recent sampling results for all sampling rules.
func xray_GetSamplingStatisticSummaries(cfg aws.Config, client *xray.Client) {
	input := &xray.GetSamplingStatisticSummariesInput{}

	if len(_xrayNextToken) > 0 {
		input.NextToken = aws.String(_xrayNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetSamplingStatisticSummaries(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*xray.GetSamplingStatisticSummariesOutput
	p := xray.NewGetSamplingStatisticSummariesPaginator(client, input)
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

// Requests a sampling quota for rules that the service is using to sample
// requests.
func xray_GetSamplingTargets(cfg aws.Config, client *xray.Client) {
	input := &xray.GetSamplingTargetsInput{
		// SamplingStatisticsDocuments: []types.SamplingStatisticsDocument, // Required
	}

	if len(_xraySamplingStatisticsDocuments) > 0 {
		if err := assignInputField(input, "SamplingStatisticsDocuments", _xraySamplingStatisticsDocuments); err != nil {
			log.Errorf("invalid --sampling-statistics-documents: %s", err.Error())
			return
		}
	}
	if len(_xraySamplingBoostStatisticsDocuments) > 0 {
		if err := assignInputField(input, "SamplingBoostStatisticsDocuments", _xraySamplingBoostStatisticsDocuments); err != nil {
			log.Errorf("invalid --sampling-boost-statistics-documents: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetSamplingTargets(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a document that describes services that process incoming requests,
// and downstream services that they call as a result. Root services process
// incoming requests and make calls to downstream services. Root services are
// applications that use the [Amazon Web Services X-Ray SDK]. Downstream services can be other applications,
// Amazon Web Services resources, HTTP web APIs, or SQL databases.
//
// [Amazon Web Services X-Ray SDK]: https://docs.aws.amazon.com/xray/index.html
func xray_GetServiceGraph(cfg aws.Config, client *xray.Client) {
	input := &xray.GetServiceGraphInput{
		// EndTime: *time.Time, // Required
		// StartTime: *time.Time, // Required
	}

	if len(_xrayEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _xrayEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_xrayStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _xrayStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_xrayGroupARN) > 0 {
		input.GroupARN = aws.String(_xrayGroupARN)
	}
	if len(_xrayGroupName) > 0 {
		input.GroupName = aws.String(_xrayGroupName)
	}
	if len(_xrayNextToken) > 0 {
		input.NextToken = aws.String(_xrayNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetServiceGraph(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*xray.GetServiceGraphOutput
	p := xray.NewGetServiceGraphPaginator(client, input)
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

// Get an aggregation of service statistics defined by a specific time range.
func xray_GetTimeSeriesServiceStatistics(cfg aws.Config, client *xray.Client) {
	input := &xray.GetTimeSeriesServiceStatisticsInput{
		// EndTime: *time.Time, // Required
		// StartTime: *time.Time, // Required
	}

	if len(_xrayEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _xrayEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_xrayStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _xrayStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_xrayEntitySelectorExpression) > 0 {
		input.EntitySelectorExpression = aws.String(_xrayEntitySelectorExpression)
	}
	if len(_xrayForecastStatistics) > 0 {
		if err := assignInputField(input, "ForecastStatistics", _xrayForecastStatistics); err != nil {
			log.Errorf("invalid --forecast-statistics: %s", err.Error())
			return
		}
	}
	if len(_xrayGroupARN) > 0 {
		input.GroupARN = aws.String(_xrayGroupARN)
	}
	if len(_xrayGroupName) > 0 {
		input.GroupName = aws.String(_xrayGroupName)
	}
	if len(_xrayNextToken) > 0 {
		input.NextToken = aws.String(_xrayNextToken)
	}
	if len(_xrayPeriod) > 0 {
		if err := assignInputField(input, "Period", _xrayPeriod); err != nil {
			log.Errorf("invalid --period: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.GetTimeSeriesServiceStatistics(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*xray.GetTimeSeriesServiceStatisticsOutput
	p := xray.NewGetTimeSeriesServiceStatisticsPaginator(client, input)
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

// Retrieves a service graph for one or more specific trace IDs.
func xray_GetTraceGraph(cfg aws.Config, client *xray.Client) {
	input := &xray.GetTraceGraphInput{
		// TraceIds: []string, // Required
	}

	if len(_xrayTraceIds) > 0 {
		input.TraceIds = append([]string(nil), _xrayTraceIds...)
	}
	if len(_xrayNextToken) > 0 {
		input.NextToken = aws.String(_xrayNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetTraceGraph(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*xray.GetTraceGraphOutput
	p := xray.NewGetTraceGraphPaginator(client, input)
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

// Retrieves the current destination of data sent to PutTraceSegments and
// OpenTelemetry protocol (OTLP) endpoint. The Transaction Search feature requires
// a CloudWatchLogs destination. For more information, see [Transaction Search]and [OpenTelemetry].
//
// [Transaction Search]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Transaction-Search.html
// [OpenTelemetry]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-OpenTelemetry-Sections.html
func xray_GetTraceSegmentDestination(cfg aws.Config, client *xray.Client) {
	input := &xray.GetTraceSegmentDestinationInput{}

	if resp, err := client.GetTraceSegmentDestination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves IDs and annotations for traces available for a specified time frame
// using an optional filter. To get the full traces, pass the trace IDs to
// BatchGetTraces .
//
// A filter expression can target traced requests that hit specific service nodes
// or edges, have errors, or come from a known user. For example, the following
// filter expression targets traces that pass through api.example.com :
//
// service("api.example.com")
//
// This filter expression finds traces that have an annotation named account with
// the value 12345 :
//
// annotation.account = "12345"
//
// For a full list of indexed fields and keywords that you can use in filter
// expressions, see [Use filter expressions]in the Amazon Web Services X-Ray Developer Guide.
//
// [Use filter expressions]: https://docs.aws.amazon.com/xray/latest/devguide/aws-xray-interface-console.html#xray-console-filters
func xray_GetTraceSummaries(cfg aws.Config, client *xray.Client) {
	input := &xray.GetTraceSummariesInput{
		// EndTime: *time.Time, // Required
		// StartTime: *time.Time, // Required
	}

	if len(_xrayEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _xrayEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_xrayStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _xrayStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_xrayFilterExpression) > 0 {
		input.FilterExpression = aws.String(_xrayFilterExpression)
	}
	if len(_xrayNextToken) > 0 {
		input.NextToken = aws.String(_xrayNextToken)
	}
	if len(_xraySampling) > 0 {
		if err := assignInputField(input, "Sampling", _xraySampling); err != nil {
			log.Errorf("invalid --sampling: %s", err.Error())
			return
		}
	}
	if len(_xraySamplingStrategy) > 0 {
		if err := assignInputField(input, "SamplingStrategy", _xraySamplingStrategy); err != nil {
			log.Errorf("invalid --sampling-strategy: %s", err.Error())
			return
		}
	}
	if len(_xrayTimeRangeType) > 0 {
		if err := assignInputField(input, "TimeRangeType", _xrayTimeRangeType); err != nil {
			log.Errorf("invalid --time-range-type: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.GetTraceSummaries(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*xray.GetTraceSummariesOutput
	p := xray.NewGetTraceSummariesPaginator(client, input)
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

// Returns the list of resource policies in the target Amazon Web Services account.
func xray_ListResourcePolicies(cfg aws.Config, client *xray.Client) {
	input := &xray.ListResourcePoliciesInput{}

	if len(_xrayNextToken) > 0 {
		input.NextToken = aws.String(_xrayNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListResourcePolicies(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*xray.ListResourcePoliciesOutput
	p := xray.NewListResourcePoliciesPaginator(client, input)
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

// Retrieves a list of traces for a given RetrievalToken from the CloudWatch log
// group generated by Transaction Search. For information on what each trace
// returns, see [BatchGetTraces].
//
// This API does not initiate a retrieval process. To start a trace retrieval, use
// StartTraceRetrieval , which generates the required RetrievalToken .
//
// When the RetrievalStatus is not COMPLETE, the API will return an empty
// response. Retry the request once the retrieval has completed to access the full
// list of traces.
//
// For cross-account observability, this API can retrieve traces from linked
// accounts when CloudWatch log is set as the destination across relevant accounts.
// For more details, see [CloudWatch cross-account observability].
//
// For retrieving data from X-Ray directly as opposed to the Transaction Search
// generated log group, see [BatchGetTraces].
//
// [BatchGetTraces]: https://docs.aws.amazon.com/xray/latest/api/API_BatchGetTraces.html
// [CloudWatch cross-account observability]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Unified-Cross-Account.html
func xray_ListRetrievedTraces(cfg aws.Config, client *xray.Client) {
	input := &xray.ListRetrievedTracesInput{
		// RetrievalToken: *string, // Required
	}

	if len(_xrayRetrievalToken) > 0 {
		input.RetrievalToken = aws.String(_xrayRetrievalToken)
	}
	if len(_xrayNextToken) > 0 {
		input.NextToken = aws.String(_xrayNextToken)
	}
	if len(_xrayTraceFormat) > 0 {
		if err := assignInputField(input, "TraceFormat", _xrayTraceFormat); err != nil {
			log.Errorf("invalid --trace-format: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListRetrievedTraces(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of tags that are applied to the specified Amazon Web Services
// X-Ray group or sampling rule.
func xray_ListTagsForResource(cfg aws.Config, client *xray.Client) {
	input := &xray.ListTagsForResourceInput{
		// ResourceARN: *string, // Required
	}

	if len(_xrayResourceARN) > 0 {
		input.ResourceARN = aws.String(_xrayResourceARN)
	}
	if len(_xrayNextToken) > 0 {
		input.NextToken = aws.String(_xrayNextToken)
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

	var results []*xray.ListTagsForResourceOutput
	p := xray.NewListTagsForResourcePaginator(client, input)
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

// Updates the encryption configuration for X-Ray data.
func xray_PutEncryptionConfig(cfg aws.Config, client *xray.Client) {
	input := &xray.PutEncryptionConfigInput{
		// Type: types.EncryptionType, // Required
	}

	if len(_xrayType) > 0 {
		if err := assignInputField(input, "Type", _xrayType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_xrayKeyId) > 0 {
		input.KeyId = aws.String(_xrayKeyId)
	}

	if resp, err := client.PutEncryptionConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets the resource policy to grant one or more Amazon Web Services services and
// accounts permissions to access X-Ray. Each resource policy will be associated
// with a specific Amazon Web Services account. Each Amazon Web Services account
// can have a maximum of 5 resource policies, and each policy name must be unique
// within that account. The maximum size of each resource policy is 5KB.
func xray_PutResourcePolicy(cfg aws.Config, client *xray.Client) {
	input := &xray.PutResourcePolicyInput{
		// PolicyDocument: *string, // Required
		// PolicyName: *string, // Required
	}

	if len(_xrayPolicyDocument) > 0 {
		input.PolicyDocument = aws.String(_xrayPolicyDocument)
	}
	if len(_xrayPolicyName) > 0 {
		input.PolicyName = aws.String(_xrayPolicyName)
	}
	if len(_xrayBypassPolicyLockoutCheck) > 0 {
		if err := assignInputField(input, "BypassPolicyLockoutCheck", _xrayBypassPolicyLockoutCheck); err != nil {
			log.Errorf("invalid --bypass-policy-lockout-check: %s", err.Error())
			return
		}
	}
	if len(_xrayPolicyRevisionId) > 0 {
		input.PolicyRevisionId = aws.String(_xrayPolicyRevisionId)
	}

	if resp, err := client.PutResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Used by the Amazon Web Services X-Ray daemon to upload telemetry.
func xray_PutTelemetryRecords(cfg aws.Config, client *xray.Client) {
	input := &xray.PutTelemetryRecordsInput{
		// TelemetryRecords: []types.TelemetryRecord, // Required
	}

	if len(_xrayTelemetryRecords) > 0 {
		if err := assignInputField(input, "TelemetryRecords", _xrayTelemetryRecords); err != nil {
			log.Errorf("invalid --telemetry-records: %s", err.Error())
			return
		}
	}
	if len(_xrayEC2InstanceId) > 0 {
		input.EC2InstanceId = aws.String(_xrayEC2InstanceId)
	}
	if len(_xrayHostname) > 0 {
		input.Hostname = aws.String(_xrayHostname)
	}
	if len(_xrayResourceARN) > 0 {
		input.ResourceARN = aws.String(_xrayResourceARN)
	}

	if resp, err := client.PutTelemetryRecords(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Uploads segment documents to Amazon Web Services X-Ray. A segment document can
// be a completed segment, an in-progress segment, or an array of subsegments.
//
// Segments must include the following fields. For the full segment document
// schema, see [Amazon Web Services X-Ray Segment Documents]in the Amazon Web Services X-Ray Developer Guide.
//
// # Required segment document fields
//
// - name - The name of the service that handled the request.
//
// - id - A 64-bit identifier for the segment, unique among segments in the same
// trace, in 16 hexadecimal digits.
//
// - trace_id - A unique identifier that connects all segments and subsegments
// originating from a single client request.
//
// - start_time - Time the segment or subsegment was created, in floating point
// seconds in epoch time, accurate to milliseconds. For example, 1480615200.010
// or 1.480615200010E9 .
//
// - end_time - Time the segment or subsegment was closed. For example,
// 1480615200.090 or 1.480615200090E9 . Specify either an end_time or in_progress
// .
//
// - in_progress - Set to true instead of specifying an end_time to record that a
// segment has been started, but is not complete. Send an in-progress segment when
// your application receives a request that will take a long time to serve, to
// trace that the request was received. When the response is sent, send the
// complete segment to overwrite the in-progress segment.
//
// A trace_id consists of three numbers separated by hyphens. For example,
// 1-58406520-a006649127e371903a2de979. For trace IDs created by an X-Ray SDK, or
// by Amazon Web Services services integrated with X-Ray, a trace ID includes:
//
// # Trace ID Format
//
// - The version number, for instance, 1 .
//
// - The time of the original request, in Unix epoch time, in 8 hexadecimal
// digits. For example, 10:00AM December 2nd, 2016 PST in epoch time is
// 1480615200 seconds, or 58406520 in hexadecimal.
//
// - A 96-bit identifier for the trace, globally unique, in 24 hexadecimal
// digits.
//
// Trace IDs created via OpenTelemetry have a different format based on the [W3C Trace Context specification]. A
// W3C trace ID must be formatted in the X-Ray trace ID format when sending to
// X-Ray. For example, a W3C trace ID 4efaaf4d1e8720b39541901950019ee5 should be
// formatted as 1-4efaaf4d-1e8720b39541901950019ee5 when sending to X-Ray. While
// X-Ray trace IDs include the original request timestamp in Unix epoch time, this
// is not required or validated.
//
// [W3C Trace Context specification]: https://www.w3.org/TR/trace-context/
// [Amazon Web Services X-Ray Segment Documents]: https://docs.aws.amazon.com/xray/latest/devguide/aws-xray-interface-api.html#xray-api-segmentdocuments.html
func xray_PutTraceSegments(cfg aws.Config, client *xray.Client) {
	input := &xray.PutTraceSegmentsInput{
		// TraceSegmentDocuments: []string, // Required
	}

	if len(_xrayTraceSegmentDocuments) > 0 {
		input.TraceSegmentDocuments = append([]string(nil), _xrayTraceSegmentDocuments...)
	}

	if resp, err := client.PutTraceSegments(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Initiates a trace retrieval process using the specified time range and for the
// given trace IDs in the Transaction Search generated CloudWatch log group. For
// more information, see [Transaction Search].
//
// API returns a RetrievalToken , which can be used with ListRetrievedTraces or
// GetRetrievedTracesGraph to fetch results. Retrievals will time out after 60
// minutes. To execute long time ranges, consider segmenting into multiple
// retrievals.
//
// If you are using [CloudWatch cross-account observability], you can use this operation in a monitoring account to
// retrieve data from a linked source account, as long as both accounts have
// transaction search enabled.
//
// For retrieving data from X-Ray directly as opposed to the Transaction-Search
// Log group, see [BatchGetTraces].
//
// [BatchGetTraces]: https://docs.aws.amazon.com/xray/latest/api/API_BatchGetTraces.html
// [Transaction Search]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Transaction-Search.html
// [CloudWatch cross-account observability]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Unified-Cross-Account.html
func xray_StartTraceRetrieval(cfg aws.Config, client *xray.Client) {
	input := &xray.StartTraceRetrievalInput{
		// EndTime: *time.Time, // Required
		// StartTime: *time.Time, // Required
		// TraceIds: []string, // Required
	}

	if len(_xrayEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _xrayEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_xrayStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _xrayStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_xrayTraceIds) > 0 {
		input.TraceIds = append([]string(nil), _xrayTraceIds...)
	}

	if resp, err := client.StartTraceRetrieval(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Applies tags to an existing Amazon Web Services X-Ray group or sampling rule.
func xray_TagResource(cfg aws.Config, client *xray.Client) {
	input := &xray.TagResourceInput{
		// ResourceARN: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_xrayResourceARN) > 0 {
		input.ResourceARN = aws.String(_xrayResourceARN)
	}
	if len(_xrayTags) > 0 {
		if err := assignInputField(input, "Tags", _xrayTags); err != nil {
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

// Removes tags from an Amazon Web Services X-Ray group or sampling rule. You
// cannot edit or delete system tags (those with an aws: prefix).
func xray_UntagResource(cfg aws.Config, client *xray.Client) {
	input := &xray.UntagResourceInput{
		// ResourceARN: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_xrayResourceARN) > 0 {
		input.ResourceARN = aws.String(_xrayResourceARN)
	}
	if len(_xrayTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _xrayTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a group resource.
func xray_UpdateGroup(cfg aws.Config, client *xray.Client) {
	input := &xray.UpdateGroupInput{}

	if len(_xrayFilterExpression) > 0 {
		input.FilterExpression = aws.String(_xrayFilterExpression)
	}
	if len(_xrayGroupARN) > 0 {
		input.GroupARN = aws.String(_xrayGroupARN)
	}
	if len(_xrayGroupName) > 0 {
		input.GroupName = aws.String(_xrayGroupName)
	}
	if len(_xrayInsightsConfiguration) > 0 {
		if err := assignInputField(input, "InsightsConfiguration", _xrayInsightsConfiguration); err != nil {
			log.Errorf("invalid --insights-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies an indexing rule’s configuration.
// Indexing rules are used for determining the sampling rate for spans indexed
// from CloudWatch Logs. For more information, see [Transaction Search].
//
// [Transaction Search]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Transaction-Search.html
func xray_UpdateIndexingRule(cfg aws.Config, client *xray.Client) {
	input := &xray.UpdateIndexingRuleInput{
		// Name: *string, // Required
		// Rule: types.IndexingRuleValueUpdate, // Required
	}

	if len(_xrayName) > 0 {
		input.Name = aws.String(_xrayName)
	}
	if len(_xrayRule) > 0 {
		if err := assignInputField(input, "Rule", _xrayRule); err != nil {
			log.Errorf("invalid --rule: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateIndexingRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies a sampling rule's configuration.
func xray_UpdateSamplingRule(cfg aws.Config, client *xray.Client) {
	input := &xray.UpdateSamplingRuleInput{
		// SamplingRuleUpdate: *types.SamplingRuleUpdate, // Required
	}

	if len(_xraySamplingRuleUpdate) > 0 {
		if err := assignInputField(input, "SamplingRuleUpdate", _xraySamplingRuleUpdate); err != nil {
			log.Errorf("invalid --sampling-rule-update: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateSamplingRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the destination of data sent to PutTraceSegments . The Transaction
// Search feature requires the CloudWatchLogs destination. For more information,
// see [Transaction Search].
//
// [Transaction Search]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Transaction-Search.html
func xray_UpdateTraceSegmentDestination(cfg aws.Config, client *xray.Client) {
	input := &xray.UpdateTraceSegmentDestinationInput{}

	if len(_xrayDestination) > 0 {
		if err := assignInputField(input, "Destination", _xrayDestination); err != nil {
			log.Errorf("invalid --destination: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateTraceSegmentDestination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_xrayCmd)
	_xrayCmd.Flags().SortFlags = false

	_xrayCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_xrayCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_xrayCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_xrayCmd.Flags().StringVarP(&_xrayBypassPolicyLockoutCheck, "bypass-policy-lockout-check", "", "", "Bypass Policy Lockout Check")
	_xrayCmd.Flags().StringVarP(&_xrayDestination, "destination", "", "", "Destination")
	_xrayCmd.Flags().StringVarP(&_xrayEC2InstanceId, "ec2-instance-id", "", "", "EC2 Instance ID")
	_xrayCmd.Flags().StringVarP(&_xrayEndTime, "end-time", "", "", "End Time")
	_xrayCmd.Flags().StringVarP(&_xrayEntitySelectorExpression, "entity-selector-expression", "", "", "Entity Selector Expression")
	_xrayCmd.Flags().StringVarP(&_xrayFilterExpression, "filter-expression", "", "", "Filter Expression")
	_xrayCmd.Flags().StringVarP(&_xrayForecastStatistics, "forecast-statistics", "", "", "Forecast Statistics")
	_xrayCmd.Flags().StringVarP(&_xrayGroupARN, "group-arn", "", "", "Group ARN")
	_xrayCmd.Flags().StringVarP(&_xrayGroupName, "group-name", "", "", "Group Name")
	_xrayCmd.Flags().StringVarP(&_xrayHostname, "hostname", "", "", "Hostname")
	_xrayCmd.Flags().StringVarP(&_xrayInsightId, "insight-id", "", "", "Insight ID")
	_xrayCmd.Flags().StringVarP(&_xrayInsightsConfiguration, "insights-configuration", "", "", "Insights Configuration")
	_xrayCmd.Flags().StringVarP(&_xrayKeyId, "key-id", "", "", "Key ID")
	_xrayCmd.Flags().StringVarP(&_xrayMaxResults, "max-results", "", "", "Max Results")
	_xrayCmd.Flags().StringVarP(&_xrayName, "name", "", "", "Name")
	_xrayCmd.Flags().StringVarP(&_xrayNextToken, "next-token", "", "", "Next Token")
	_xrayCmd.Flags().StringVarP(&_xrayPeriod, "period", "", "", "Period")
	_xrayCmd.Flags().StringVarP(&_xrayPolicyDocument, "policy-document", "", "", "Policy Document")
	_xrayCmd.Flags().StringVarP(&_xrayPolicyName, "policy-name", "", "", "Policy Name")
	_xrayCmd.Flags().StringVarP(&_xrayPolicyRevisionId, "policy-revision-id", "", "", "Policy Revision ID")
	_xrayCmd.Flags().StringVarP(&_xrayResourceARN, "resource-arn", "", "", "Resource ARN")
	_xrayCmd.Flags().StringVarP(&_xrayRetrievalToken, "retrieval-token", "", "", "Retrieval Token")
	_xrayCmd.Flags().StringVarP(&_xrayRule, "rule", "", "", "Rule")
	_xrayCmd.Flags().StringVarP(&_xrayRuleARN, "rule-arn", "", "", "Rule ARN")
	_xrayCmd.Flags().StringVarP(&_xrayRuleName, "rule-name", "", "", "Rule Name")
	_xrayCmd.Flags().StringVarP(&_xraySampling, "sampling", "", "", "Sampling")
	_xrayCmd.Flags().StringVarP(&_xraySamplingBoostStatisticsDocuments, "sampling-boost-statistics-documents", "", "", "Sampling Boost Statistics Documents")
	_xrayCmd.Flags().StringVarP(&_xraySamplingRule, "sampling-rule", "", "", "Sampling Rule")
	_xrayCmd.Flags().StringVarP(&_xraySamplingRuleUpdate, "sampling-rule-update", "", "", "Sampling Rule Update")
	_xrayCmd.Flags().StringVarP(&_xraySamplingStatisticsDocuments, "sampling-statistics-documents", "", "", "Sampling Statistics Documents")
	_xrayCmd.Flags().StringVarP(&_xraySamplingStrategy, "sampling-strategy", "", "", "Sampling Strategy")
	_xrayCmd.Flags().StringVarP(&_xrayStartTime, "start-time", "", "", "Start Time")
	_xrayCmd.Flags().StringVarP(&_xrayStates, "states", "", "", "States")
	_xrayCmd.Flags().StringSliceVarP(&_xrayTagKeys, "tag-keys", "", nil, "Tag Keys")
	_xrayCmd.Flags().StringVarP(&_xrayTags, "tags", "", "", "Tags")
	_xrayCmd.Flags().StringVarP(&_xrayTelemetryRecords, "telemetry-records", "", "", "Telemetry Records")
	_xrayCmd.Flags().StringVarP(&_xrayTimeRangeType, "time-range-type", "", "", "Time Range Type")
	_xrayCmd.Flags().StringVarP(&_xrayTraceFormat, "trace-format", "", "", "Trace Format")
	_xrayCmd.Flags().StringSliceVarP(&_xrayTraceIds, "trace-ids", "", nil, "Trace Ids")
	_xrayCmd.Flags().StringSliceVarP(&_xrayTraceSegmentDocuments, "trace-segment-documents", "", nil, "Trace Segment Documents")
	_xrayCmd.Flags().StringVarP(&_xrayType, "type", "", "", "Type")

	_xrayCmd.Flags().BoolVarP(&_xrayBatchGetTraces, "batch-get-traces", "", false, "Batch Get Traces")
	_xrayCmd.Flags().BoolVarP(&_xrayCancelTraceRetrieval, "cancel-trace-retrieval", "", false, "Cancel Trace Retrieval")
	_xrayCmd.Flags().BoolVarP(&_xrayCreateGroup, "create-group", "", false, "Create Group")
	_xrayCmd.Flags().BoolVarP(&_xrayCreateSamplingRule, "create-sampling-rule", "", false, "Create Sampling Rule")
	_xrayCmd.Flags().BoolVarP(&_xrayDeleteGroup, "delete-group", "", false, "Delete Group")
	_xrayCmd.Flags().BoolVarP(&_xrayDeleteResourcePolicy, "delete-resource-policy", "", false, "Delete Resource Policy")
	_xrayCmd.Flags().BoolVarP(&_xrayDeleteSamplingRule, "delete-sampling-rule", "", false, "Delete Sampling Rule")
	_xrayCmd.Flags().BoolVarP(&_xrayGetEncryptionConfig, "get-encryption-config", "", false, "Get Encryption Config")
	_xrayCmd.Flags().BoolVarP(&_xrayGetGroup, "get-group", "", false, "Get Group")
	_xrayCmd.Flags().BoolVarP(&_xrayGetGroups, "get-groups", "", false, "Get Groups")
	_xrayCmd.Flags().BoolVarP(&_xrayGetIndexingRules, "get-indexing-rules", "", false, "Get Indexing Rules")
	_xrayCmd.Flags().BoolVarP(&_xrayGetInsight, "get-insight", "", false, "Get Insight")
	_xrayCmd.Flags().BoolVarP(&_xrayGetInsightEvents, "get-insight-events", "", false, "Get Insight Events")
	_xrayCmd.Flags().BoolVarP(&_xrayGetInsightImpactGraph, "get-insight-impact-graph", "", false, "Get Insight Impact Graph")
	_xrayCmd.Flags().BoolVarP(&_xrayGetInsightSummaries, "get-insight-summaries", "", false, "Get Insight Summaries")
	_xrayCmd.Flags().BoolVarP(&_xrayGetRetrievedTracesGraph, "get-retrieved-traces-graph", "", false, "Get Retrieved Traces Graph")
	_xrayCmd.Flags().BoolVarP(&_xrayGetSamplingRules, "get-sampling-rules", "", false, "Get Sampling Rules")
	_xrayCmd.Flags().BoolVarP(&_xrayGetSamplingStatisticSummaries, "get-sampling-statistic-summaries", "", false, "Get Sampling Statistic Summaries")
	_xrayCmd.Flags().BoolVarP(&_xrayGetSamplingTargets, "get-sampling-targets", "", false, "Get Sampling Targets")
	_xrayCmd.Flags().BoolVarP(&_xrayGetServiceGraph, "get-service-graph", "", false, "Get Service Graph")
	_xrayCmd.Flags().BoolVarP(&_xrayGetTimeSeriesServiceStatistics, "get-time-series-service-statistics", "", false, "Get Time Series Service Statistics")
	_xrayCmd.Flags().BoolVarP(&_xrayGetTraceGraph, "get-trace-graph", "", false, "Get Trace Graph")
	_xrayCmd.Flags().BoolVarP(&_xrayGetTraceSegmentDestination, "get-trace-segment-destination", "", false, "Get Trace Segment Destination")
	_xrayCmd.Flags().BoolVarP(&_xrayGetTraceSummaries, "get-trace-summaries", "", false, "Get Trace Summaries")
	_xrayCmd.Flags().BoolVarP(&_xrayListResourcePolicies, "list-resource-policies", "", false, "List Resource Policies")
	_xrayCmd.Flags().BoolVarP(&_xrayListRetrievedTraces, "list-retrieved-traces", "", false, "List Retrieved Traces")
	_xrayCmd.Flags().BoolVarP(&_xrayListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_xrayCmd.Flags().BoolVarP(&_xrayPutEncryptionConfig, "put-encryption-config", "", false, "Put Encryption Config")
	_xrayCmd.Flags().BoolVarP(&_xrayPutResourcePolicy, "put-resource-policy", "", false, "Put Resource Policy")
	_xrayCmd.Flags().BoolVarP(&_xrayPutTelemetryRecords, "put-telemetry-records", "", false, "Put Telemetry Records")
	_xrayCmd.Flags().BoolVarP(&_xrayPutTraceSegments, "put-trace-segments", "", false, "Put Trace Segments")
	_xrayCmd.Flags().BoolVarP(&_xrayStartTraceRetrieval, "start-trace-retrieval", "", false, "Start Trace Retrieval")
	_xrayCmd.Flags().BoolVarP(&_xrayTagResource, "tag-resource", "", false, "Tag Resource")
	_xrayCmd.Flags().BoolVarP(&_xrayUntagResource, "untag-resource", "", false, "Untag Resource")
	_xrayCmd.Flags().BoolVarP(&_xrayUpdateGroup, "update-group", "", false, "Update Group")
	_xrayCmd.Flags().BoolVarP(&_xrayUpdateIndexingRule, "update-indexing-rule", "", false, "Update Indexing Rule")
	_xrayCmd.Flags().BoolVarP(&_xrayUpdateSamplingRule, "update-sampling-rule", "", false, "Update Sampling Rule")
	_xrayCmd.Flags().BoolVarP(&_xrayUpdateTraceSegmentDestination, "update-trace-segment-destination", "", false, "Update Trace Segment Destination")

}
