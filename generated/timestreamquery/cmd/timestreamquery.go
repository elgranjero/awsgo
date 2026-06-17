package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/timestreamquery"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// timestreamqueryCmd represents the timestreamquery command
var _timestreamqueryCmd = &cobra.Command{
	Use:   "timestreamquery",
	Short: "AWS timestreamquery CLI",
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
		client := timestreamquery.NewFromConfig(cfg)
		if _timestreamqueryCancelQuery {
			timestreamquery_CancelQuery(cfg, client)
			return
		}
		if _timestreamqueryCreateScheduledQuery {
			timestreamquery_CreateScheduledQuery(cfg, client)
			return
		}
		if _timestreamqueryDeleteScheduledQuery {
			timestreamquery_DeleteScheduledQuery(cfg, client)
			return
		}
		if _timestreamqueryDescribeAccountSettings {
			timestreamquery_DescribeAccountSettings(cfg, client)
			return
		}
		if _timestreamqueryDescribeEndpoints {
			timestreamquery_DescribeEndpoints(cfg, client)
			return
		}
		if _timestreamqueryDescribeScheduledQuery {
			timestreamquery_DescribeScheduledQuery(cfg, client)
			return
		}
		if _timestreamqueryExecuteScheduledQuery {
			timestreamquery_ExecuteScheduledQuery(cfg, client)
			return
		}
		if _timestreamqueryListScheduledQueries {
			timestreamquery_ListScheduledQueries(cfg, client)
			return
		}
		if _timestreamqueryListTagsForResource {
			timestreamquery_ListTagsForResource(cfg, client)
			return
		}
		if _timestreamqueryPrepareQuery {
			timestreamquery_PrepareQuery(cfg, client)
			return
		}
		if _timestreamqueryQuery {
			timestreamquery_Query(cfg, client)
			return
		}
		if _timestreamqueryTagResource {
			timestreamquery_TagResource(cfg, client)
			return
		}
		if _timestreamqueryUntagResource {
			timestreamquery_UntagResource(cfg, client)
			return
		}
		if _timestreamqueryUpdateAccountSettings {
			timestreamquery_UpdateAccountSettings(cfg, client)
			return
		}
		if _timestreamqueryUpdateScheduledQuery {
			timestreamquery_UpdateScheduledQuery(cfg, client)
			return
		}

	},
}

var (
	_timestreamqueryCancelQuery             bool
	_timestreamqueryCreateScheduledQuery    bool
	_timestreamqueryDeleteScheduledQuery    bool
	_timestreamqueryDescribeAccountSettings bool
	_timestreamqueryDescribeEndpoints       bool
	_timestreamqueryDescribeScheduledQuery  bool
	_timestreamqueryExecuteScheduledQuery   bool
	_timestreamqueryListScheduledQueries    bool
	_timestreamqueryListTagsForResource     bool
	_timestreamqueryPrepareQuery            bool
	_timestreamqueryQuery                   bool
	_timestreamqueryTagResource             bool
	_timestreamqueryUntagResource           bool
	_timestreamqueryUpdateAccountSettings   bool
	_timestreamqueryUpdateScheduledQuery    bool

	_timestreamqueryClientToken                    string
	_timestreamqueryErrorReportConfiguration       string
	_timestreamqueryInvocationTime                 string
	_timestreamqueryKmsKeyId                       string
	_timestreamqueryMaxQueryTCU                    string
	_timestreamqueryMaxResults                     string
	_timestreamqueryMaxRows                        string
	_timestreamqueryName                           string
	_timestreamqueryNextToken                      string
	_timestreamqueryNotificationConfiguration      string
	_timestreamqueryQueryCompute                   string
	_timestreamqueryQueryId                        string
	_timestreamqueryQueryInsights                  string
	_timestreamqueryQueryPricingModel              string
	_timestreamqueryQueryString                    string
	_timestreamqueryResourceARN                    string
	_timestreamqueryScheduleConfiguration          string
	_timestreamqueryScheduledQueryArn              string
	_timestreamqueryScheduledQueryExecutionRoleArn string
	_timestreamqueryState                          string
	_timestreamqueryTagKeys                        []string
	_timestreamqueryTags                           string
	_timestreamqueryTargetConfiguration            string
	_timestreamqueryValidateOnly                   string
)

// Cancels a query that has been issued. Cancellation is provided only if the
// query has not completed running before the cancellation request was issued.
// Because cancellation is an idempotent operation, subsequent cancellation
// requests will return a CancellationMessage , indicating that the query has
// already been canceled. See [code sample]for details.
//
// [code sample]: https://docs.aws.amazon.com/timestream/latest/developerguide/code-samples.cancel-query.html
func timestreamquery_CancelQuery(cfg aws.Config, client *timestreamquery.Client) {
	input := &timestreamquery.CancelQueryInput{
		// QueryId: *string, // Required
	}

	if len(_timestreamqueryQueryId) > 0 {
		input.QueryId = aws.String(_timestreamqueryQueryId)
	}

	if resp, err := client.CancelQuery(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a scheduled query that will be run on your behalf at the configured
// schedule. Timestream assumes the execution role provided as part of the
// ScheduledQueryExecutionRoleArn parameter to run the query. You can use the
// NotificationConfiguration parameter to configure notification for your scheduled
// query operations.
func timestreamquery_CreateScheduledQuery(cfg aws.Config, client *timestreamquery.Client) {
	input := &timestreamquery.CreateScheduledQueryInput{
		// ErrorReportConfiguration: *types.ErrorReportConfiguration, // Required
		// Name: *string, // Required
		// NotificationConfiguration: *types.NotificationConfiguration, // Required
		// QueryString: *string, // Required
		// ScheduleConfiguration: *types.ScheduleConfiguration, // Required
		// ScheduledQueryExecutionRoleArn: *string, // Required
	}

	if len(_timestreamqueryErrorReportConfiguration) > 0 {
		if err := assignInputField(input, "ErrorReportConfiguration", _timestreamqueryErrorReportConfiguration); err != nil {
			log.Errorf("invalid --error-report-configuration: %s", err.Error())
			return
		}
	}
	if len(_timestreamqueryName) > 0 {
		input.Name = aws.String(_timestreamqueryName)
	}
	if len(_timestreamqueryNotificationConfiguration) > 0 {
		if err := assignInputField(input, "NotificationConfiguration", _timestreamqueryNotificationConfiguration); err != nil {
			log.Errorf("invalid --notification-configuration: %s", err.Error())
			return
		}
	}
	if len(_timestreamqueryQueryString) > 0 {
		input.QueryString = aws.String(_timestreamqueryQueryString)
	}
	if len(_timestreamqueryScheduleConfiguration) > 0 {
		if err := assignInputField(input, "ScheduleConfiguration", _timestreamqueryScheduleConfiguration); err != nil {
			log.Errorf("invalid --schedule-configuration: %s", err.Error())
			return
		}
	}
	if len(_timestreamqueryScheduledQueryExecutionRoleArn) > 0 {
		input.ScheduledQueryExecutionRoleArn = aws.String(_timestreamqueryScheduledQueryExecutionRoleArn)
	}
	if len(_timestreamqueryClientToken) > 0 {
		input.ClientToken = aws.String(_timestreamqueryClientToken)
	}
	if len(_timestreamqueryKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_timestreamqueryKmsKeyId)
	}
	if len(_timestreamqueryTags) > 0 {
		if err := assignInputField(input, "Tags", _timestreamqueryTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_timestreamqueryTargetConfiguration) > 0 {
		if err := assignInputField(input, "TargetConfiguration", _timestreamqueryTargetConfiguration); err != nil {
			log.Errorf("invalid --target-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateScheduledQuery(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a given scheduled query. This is an irreversible operation.
func timestreamquery_DeleteScheduledQuery(cfg aws.Config, client *timestreamquery.Client) {
	input := &timestreamquery.DeleteScheduledQueryInput{
		// ScheduledQueryArn: *string, // Required
	}

	if len(_timestreamqueryScheduledQueryArn) > 0 {
		input.ScheduledQueryArn = aws.String(_timestreamqueryScheduledQueryArn)
	}

	if resp, err := client.DeleteScheduledQuery(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the settings for your account that include the query pricing model
// and the configured maximum TCUs the service can use for your query workload.
//
// You're charged only for the duration of compute units used for your workloads.
func timestreamquery_DescribeAccountSettings(cfg aws.Config, client *timestreamquery.Client) {
	input := &timestreamquery.DescribeAccountSettingsInput{}

	if resp, err := client.DescribeAccountSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// DescribeEndpoints returns a list of available endpoints to make Timestream API
// calls against. This API is available through both Write and Query.
//
// Because the Timestream SDKs are designed to transparently work with the
// service’s architecture, including the management and mapping of the service
// endpoints, it is not recommended that you use this API unless:
//
// - You are using [VPC endpoints (Amazon Web Services PrivateLink) with Timestream]
//
// - Your application uses a programming language that does not yet have SDK
// support
//
// - You require better control over the client-side implementation
//
// For detailed information on how and when to use and implement
// DescribeEndpoints, see [The Endpoint Discovery Pattern].
//
// [The Endpoint Discovery Pattern]: https://docs.aws.amazon.com/timestream/latest/developerguide/Using.API.html#Using-API.endpoint-discovery
// [VPC endpoints (Amazon Web Services PrivateLink) with Timestream]: https://docs.aws.amazon.com/timestream/latest/developerguide/VPCEndpoints
func timestreamquery_DescribeEndpoints(cfg aws.Config, client *timestreamquery.Client) {
	input := &timestreamquery.DescribeEndpointsInput{}

	if resp, err := client.DescribeEndpoints(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides detailed information about a scheduled query.
func timestreamquery_DescribeScheduledQuery(cfg aws.Config, client *timestreamquery.Client) {
	input := &timestreamquery.DescribeScheduledQueryInput{
		// ScheduledQueryArn: *string, // Required
	}

	if len(_timestreamqueryScheduledQueryArn) > 0 {
		input.ScheduledQueryArn = aws.String(_timestreamqueryScheduledQueryArn)
	}

	if resp, err := client.DescribeScheduledQuery(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// You can use this API to run a scheduled query manually.
// If you enabled QueryInsights , this API also returns insights and metrics
// related to the query that you executed as part of an Amazon SNS notification.
// QueryInsights helps with performance tuning of your query. For more information
// about QueryInsights , see [Using query insights to optimize queries in Amazon Timestream].
//
// [Using query insights to optimize queries in Amazon Timestream]: https://docs.aws.amazon.com/timestream/latest/developerguide/using-query-insights.html
func timestreamquery_ExecuteScheduledQuery(cfg aws.Config, client *timestreamquery.Client) {
	input := &timestreamquery.ExecuteScheduledQueryInput{
		// InvocationTime: *time.Time, // Required
		// ScheduledQueryArn: *string, // Required
	}

	if len(_timestreamqueryInvocationTime) > 0 {
		if err := assignInputField(input, "InvocationTime", _timestreamqueryInvocationTime); err != nil {
			log.Errorf("invalid --invocation-time: %s", err.Error())
			return
		}
	}
	if len(_timestreamqueryScheduledQueryArn) > 0 {
		input.ScheduledQueryArn = aws.String(_timestreamqueryScheduledQueryArn)
	}
	if len(_timestreamqueryClientToken) > 0 {
		input.ClientToken = aws.String(_timestreamqueryClientToken)
	}
	if len(_timestreamqueryQueryInsights) > 0 {
		if err := assignInputField(input, "QueryInsights", _timestreamqueryQueryInsights); err != nil {
			log.Errorf("invalid --query-insights: %s", err.Error())
			return
		}
	}

	if resp, err := client.ExecuteScheduledQuery(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a list of all scheduled queries in the caller's Amazon account and Region.
// ListScheduledQueries is eventually consistent.
func timestreamquery_ListScheduledQueries(cfg aws.Config, client *timestreamquery.Client) {
	input := &timestreamquery.ListScheduledQueriesInput{}

	if len(_timestreamqueryMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _timestreamqueryMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_timestreamqueryNextToken) > 0 {
		input.NextToken = aws.String(_timestreamqueryNextToken)
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

	var results []*timestreamquery.ListScheduledQueriesOutput
	p := timestreamquery.NewListScheduledQueriesPaginator(client, input)
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

// List all tags on a Timestream query resource.
func timestreamquery_ListTagsForResource(cfg aws.Config, client *timestreamquery.Client) {
	input := &timestreamquery.ListTagsForResourceInput{
		// ResourceARN: *string, // Required
	}

	if len(_timestreamqueryResourceARN) > 0 {
		input.ResourceARN = aws.String(_timestreamqueryResourceARN)
	}
	if len(_timestreamqueryMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _timestreamqueryMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_timestreamqueryNextToken) > 0 {
		input.NextToken = aws.String(_timestreamqueryNextToken)
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

	var results []*timestreamquery.ListTagsForResourceOutput
	p := timestreamquery.NewListTagsForResourcePaginator(client, input)
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

// A synchronous operation that allows you to submit a query with parameters to be
// stored by Timestream for later running. Timestream only supports using this
// operation with ValidateOnly set to true .
func timestreamquery_PrepareQuery(cfg aws.Config, client *timestreamquery.Client) {
	input := &timestreamquery.PrepareQueryInput{
		// QueryString: *string, // Required
	}

	if len(_timestreamqueryQueryString) > 0 {
		input.QueryString = aws.String(_timestreamqueryQueryString)
	}
	if len(_timestreamqueryValidateOnly) > 0 {
		if err := assignInputField(input, "ValidateOnly", _timestreamqueryValidateOnly); err != nil {
			log.Errorf("invalid --validate-only: %s", err.Error())
			return
		}
	}

	if resp, err := client.PrepareQuery(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Query is a synchronous operation that enables you to run a query against your
// Amazon Timestream data.
//
// If you enabled QueryInsights , this API also returns insights and metrics
// related to the query that you executed. QueryInsights helps with performance
// tuning of your query. For more information about QueryInsights , see [Using query insights to optimize queries in Amazon Timestream].
//
// The maximum number of Query API requests you're allowed to make with
// QueryInsights enabled is 1 query per second (QPS). If you exceed this query
// rate, it might result in throttling.
//
// Query will time out after 60 seconds. You must update the default timeout in
// the SDK to support a timeout of 60 seconds. See the [code sample]for details.
//
// Your query request will fail in the following cases:
//
// - If you submit a Query request with the same client token outside of the
// 5-minute idempotency window.
//
// - If you submit a Query request with the same client token, but change other
// parameters, within the 5-minute idempotency window.
//
// - If the size of the row (including the query metadata) exceeds 1 MB, then
// the query will fail with the following error message:
//
// # Query aborted as max page response size has been exceeded by the output result
//
// row
//
// - If the IAM principal of the query initiator and the result reader are not
// the same and/or the query initiator and the result reader do not have the same
// query string in the query requests, the query will fail with an Invalid
// pagination token error.
//
// [code sample]: https://docs.aws.amazon.com/timestream/latest/developerguide/code-samples.run-query.html
// [Using query insights to optimize queries in Amazon Timestream]: https://docs.aws.amazon.com/timestream/latest/developerguide/using-query-insights.html
func timestreamquery_Query(cfg aws.Config, client *timestreamquery.Client) {
	input := &timestreamquery.QueryInput{
		// QueryString: *string, // Required
	}

	if len(_timestreamqueryQueryString) > 0 {
		input.QueryString = aws.String(_timestreamqueryQueryString)
	}
	if len(_timestreamqueryClientToken) > 0 {
		input.ClientToken = aws.String(_timestreamqueryClientToken)
	}
	if len(_timestreamqueryMaxRows) > 0 {
		if err := assignInputField(input, "MaxRows", _timestreamqueryMaxRows); err != nil {
			log.Errorf("invalid --max-rows: %s", err.Error())
			return
		}
	}
	if len(_timestreamqueryNextToken) > 0 {
		input.NextToken = aws.String(_timestreamqueryNextToken)
	}
	if len(_timestreamqueryQueryInsights) > 0 {
		if err := assignInputField(input, "QueryInsights", _timestreamqueryQueryInsights); err != nil {
			log.Errorf("invalid --query-insights: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.Query(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*timestreamquery.QueryOutput
	p := timestreamquery.NewQueryPaginator(client, input)
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

// Associate a set of tags with a Timestream resource. You can then activate these
// user-defined tags so that they appear on the Billing and Cost Management console
// for cost allocation tracking.
func timestreamquery_TagResource(cfg aws.Config, client *timestreamquery.Client) {
	input := &timestreamquery.TagResourceInput{
		// ResourceARN: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_timestreamqueryResourceARN) > 0 {
		input.ResourceARN = aws.String(_timestreamqueryResourceARN)
	}
	if len(_timestreamqueryTags) > 0 {
		if err := assignInputField(input, "Tags", _timestreamqueryTags); err != nil {
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

// Removes the association of tags from a Timestream query resource.
func timestreamquery_UntagResource(cfg aws.Config, client *timestreamquery.Client) {
	input := &timestreamquery.UntagResourceInput{
		// ResourceARN: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_timestreamqueryResourceARN) > 0 {
		input.ResourceARN = aws.String(_timestreamqueryResourceARN)
	}
	if len(_timestreamqueryTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _timestreamqueryTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Transitions your account to use TCUs for query pricing and modifies the maximum
// query compute units that you've configured. If you reduce the value of
// MaxQueryTCU to a desired configuration, the new value can take up to 24 hours to
// be effective.
//
// After you've transitioned your account to use TCUs for query pricing, you can't
// transition to using bytes scanned for query pricing.
func timestreamquery_UpdateAccountSettings(cfg aws.Config, client *timestreamquery.Client) {
	input := &timestreamquery.UpdateAccountSettingsInput{}

	if len(_timestreamqueryMaxQueryTCU) > 0 {
		if err := assignInputField(input, "MaxQueryTCU", _timestreamqueryMaxQueryTCU); err != nil {
			log.Errorf("invalid --max-query-tcu: %s", err.Error())
			return
		}
	}
	if len(_timestreamqueryQueryCompute) > 0 {
		if err := assignInputField(input, "QueryCompute", _timestreamqueryQueryCompute); err != nil {
			log.Errorf("invalid --query-compute: %s", err.Error())
			return
		}
	}
	if len(_timestreamqueryQueryPricingModel) > 0 {
		if err := assignInputField(input, "QueryPricingModel", _timestreamqueryQueryPricingModel); err != nil {
			log.Errorf("invalid --query-pricing-model: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateAccountSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update a scheduled query.
func timestreamquery_UpdateScheduledQuery(cfg aws.Config, client *timestreamquery.Client) {
	input := &timestreamquery.UpdateScheduledQueryInput{
		// ScheduledQueryArn: *string, // Required
		// State: types.ScheduledQueryState, // Required
	}

	if len(_timestreamqueryScheduledQueryArn) > 0 {
		input.ScheduledQueryArn = aws.String(_timestreamqueryScheduledQueryArn)
	}
	if len(_timestreamqueryState) > 0 {
		if err := assignInputField(input, "State", _timestreamqueryState); err != nil {
			log.Errorf("invalid --state: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateScheduledQuery(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_timestreamqueryCmd)
	_timestreamqueryCmd.Flags().SortFlags = false

	_timestreamqueryCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_timestreamqueryCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_timestreamqueryCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_timestreamqueryCmd.Flags().StringVarP(&_timestreamqueryClientToken, "client-token", "", "", "Client Token")
	_timestreamqueryCmd.Flags().StringVarP(&_timestreamqueryErrorReportConfiguration, "error-report-configuration", "", "", "Error Report Configuration")
	_timestreamqueryCmd.Flags().StringVarP(&_timestreamqueryInvocationTime, "invocation-time", "", "", "Invocation Time")
	_timestreamqueryCmd.Flags().StringVarP(&_timestreamqueryKmsKeyId, "kms-key-id", "", "", "KMS Key ID")
	_timestreamqueryCmd.Flags().StringVarP(&_timestreamqueryMaxQueryTCU, "max-query-tcu", "", "", "Max Query Tcu")
	_timestreamqueryCmd.Flags().StringVarP(&_timestreamqueryMaxResults, "max-results", "", "", "Max Results")
	_timestreamqueryCmd.Flags().StringVarP(&_timestreamqueryMaxRows, "max-rows", "", "", "Max Rows")
	_timestreamqueryCmd.Flags().StringVarP(&_timestreamqueryName, "name", "", "", "Name")
	_timestreamqueryCmd.Flags().StringVarP(&_timestreamqueryNextToken, "next-token", "", "", "Next Token")
	_timestreamqueryCmd.Flags().StringVarP(&_timestreamqueryNotificationConfiguration, "notification-configuration", "", "", "Notification Configuration")
	_timestreamqueryCmd.Flags().StringVarP(&_timestreamqueryQueryCompute, "query-compute", "", "", "Query Compute")
	_timestreamqueryCmd.Flags().StringVarP(&_timestreamqueryQueryId, "query-id", "", "", "Query ID")
	_timestreamqueryCmd.Flags().StringVarP(&_timestreamqueryQueryInsights, "query-insights", "", "", "Query Insights")
	_timestreamqueryCmd.Flags().StringVarP(&_timestreamqueryQueryPricingModel, "query-pricing-model", "", "", "Query Pricing Model")
	_timestreamqueryCmd.Flags().StringVarP(&_timestreamqueryQueryString, "query-string", "", "", "Query String")
	_timestreamqueryCmd.Flags().StringVarP(&_timestreamqueryResourceARN, "resource-arn", "", "", "Resource ARN")
	_timestreamqueryCmd.Flags().StringVarP(&_timestreamqueryScheduleConfiguration, "schedule-configuration", "", "", "Schedule Configuration")
	_timestreamqueryCmd.Flags().StringVarP(&_timestreamqueryScheduledQueryArn, "scheduled-query-arn", "", "", "Scheduled Query ARN")
	_timestreamqueryCmd.Flags().StringVarP(&_timestreamqueryScheduledQueryExecutionRoleArn, "scheduled-query-execution-role-arn", "", "", "Scheduled Query Execution Role ARN")
	_timestreamqueryCmd.Flags().StringVarP(&_timestreamqueryState, "state", "", "", "State")
	_timestreamqueryCmd.Flags().StringSliceVarP(&_timestreamqueryTagKeys, "tag-keys", "", nil, "Tag Keys")
	_timestreamqueryCmd.Flags().StringVarP(&_timestreamqueryTags, "tags", "", "", "Tags")
	_timestreamqueryCmd.Flags().StringVarP(&_timestreamqueryTargetConfiguration, "target-configuration", "", "", "Target Configuration")
	_timestreamqueryCmd.Flags().StringVarP(&_timestreamqueryValidateOnly, "validate-only", "", "", "Validate Only")

	_timestreamqueryCmd.Flags().BoolVarP(&_timestreamqueryCancelQuery, "cancel-query", "", false, "Cancel Query")
	_timestreamqueryCmd.Flags().BoolVarP(&_timestreamqueryCreateScheduledQuery, "create-scheduled-query", "", false, "Create Scheduled Query")
	_timestreamqueryCmd.Flags().BoolVarP(&_timestreamqueryDeleteScheduledQuery, "delete-scheduled-query", "", false, "Delete Scheduled Query")
	_timestreamqueryCmd.Flags().BoolVarP(&_timestreamqueryDescribeAccountSettings, "describe-account-settings", "", false, "Describe Account Settings")
	_timestreamqueryCmd.Flags().BoolVarP(&_timestreamqueryDescribeEndpoints, "describe-endpoints", "", false, "Describe Endpoints")
	_timestreamqueryCmd.Flags().BoolVarP(&_timestreamqueryDescribeScheduledQuery, "describe-scheduled-query", "", false, "Describe Scheduled Query")
	_timestreamqueryCmd.Flags().BoolVarP(&_timestreamqueryExecuteScheduledQuery, "execute-scheduled-query", "", false, "Execute Scheduled Query")
	_timestreamqueryCmd.Flags().BoolVarP(&_timestreamqueryListScheduledQueries, "list-scheduled-queries", "", false, "List Scheduled Queries")
	_timestreamqueryCmd.Flags().BoolVarP(&_timestreamqueryListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_timestreamqueryCmd.Flags().BoolVarP(&_timestreamqueryPrepareQuery, "prepare-query", "", false, "Prepare Query")
	_timestreamqueryCmd.Flags().BoolVarP(&_timestreamqueryQuery, "query", "", false, "Query")
	_timestreamqueryCmd.Flags().BoolVarP(&_timestreamqueryTagResource, "tag-resource", "", false, "Tag Resource")
	_timestreamqueryCmd.Flags().BoolVarP(&_timestreamqueryUntagResource, "untag-resource", "", false, "Untag Resource")
	_timestreamqueryCmd.Flags().BoolVarP(&_timestreamqueryUpdateAccountSettings, "update-account-settings", "", false, "Update Account Settings")
	_timestreamqueryCmd.Flags().BoolVarP(&_timestreamqueryUpdateScheduledQuery, "update-scheduled-query", "", false, "Update Scheduled Query")

}
