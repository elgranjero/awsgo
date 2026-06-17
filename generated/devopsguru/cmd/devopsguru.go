package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/devopsguru"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// devopsguruCmd represents the devopsguru command
var _devopsguruCmd = &cobra.Command{
	Use:   "devopsguru",
	Short: "AWS devopsguru CLI",
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
		client := devopsguru.NewFromConfig(cfg)
		if _devopsguruAddNotificationChannel {
			devopsguru_AddNotificationChannel(cfg, client)
			return
		}
		if _devopsguruDeleteInsight {
			devopsguru_DeleteInsight(cfg, client)
			return
		}
		if _devopsguruDescribeAccountHealth {
			devopsguru_DescribeAccountHealth(cfg, client)
			return
		}
		if _devopsguruDescribeAccountOverview {
			devopsguru_DescribeAccountOverview(cfg, client)
			return
		}
		if _devopsguruDescribeAnomaly {
			devopsguru_DescribeAnomaly(cfg, client)
			return
		}
		if _devopsguruDescribeEventSourcesConfig {
			devopsguru_DescribeEventSourcesConfig(cfg, client)
			return
		}
		if _devopsguruDescribeFeedback {
			devopsguru_DescribeFeedback(cfg, client)
			return
		}
		if _devopsguruDescribeInsight {
			devopsguru_DescribeInsight(cfg, client)
			return
		}
		if _devopsguruDescribeOrganizationHealth {
			devopsguru_DescribeOrganizationHealth(cfg, client)
			return
		}
		if _devopsguruDescribeOrganizationOverview {
			devopsguru_DescribeOrganizationOverview(cfg, client)
			return
		}
		if _devopsguruDescribeOrganizationResourceCollectionHealth {
			devopsguru_DescribeOrganizationResourceCollectionHealth(cfg, client)
			return
		}
		if _devopsguruDescribeResourceCollectionHealth {
			devopsguru_DescribeResourceCollectionHealth(cfg, client)
			return
		}
		if _devopsguruDescribeServiceIntegration {
			devopsguru_DescribeServiceIntegration(cfg, client)
			return
		}
		if _devopsguruGetCostEstimation {
			devopsguru_GetCostEstimation(cfg, client)
			return
		}
		if _devopsguruGetResourceCollection {
			devopsguru_GetResourceCollection(cfg, client)
			return
		}
		if _devopsguruListAnomaliesForInsight {
			devopsguru_ListAnomaliesForInsight(cfg, client)
			return
		}
		if _devopsguruListAnomalousLogGroups {
			devopsguru_ListAnomalousLogGroups(cfg, client)
			return
		}
		if _devopsguruListEvents {
			devopsguru_ListEvents(cfg, client)
			return
		}
		if _devopsguruListInsights {
			devopsguru_ListInsights(cfg, client)
			return
		}
		if _devopsguruListMonitoredResources {
			devopsguru_ListMonitoredResources(cfg, client)
			return
		}
		if _devopsguruListNotificationChannels {
			devopsguru_ListNotificationChannels(cfg, client)
			return
		}
		if _devopsguruListOrganizationInsights {
			devopsguru_ListOrganizationInsights(cfg, client)
			return
		}
		if _devopsguruListRecommendations {
			devopsguru_ListRecommendations(cfg, client)
			return
		}
		if _devopsguruPutFeedback {
			devopsguru_PutFeedback(cfg, client)
			return
		}
		if _devopsguruRemoveNotificationChannel {
			devopsguru_RemoveNotificationChannel(cfg, client)
			return
		}
		if _devopsguruSearchInsights {
			devopsguru_SearchInsights(cfg, client)
			return
		}
		if _devopsguruSearchOrganizationInsights {
			devopsguru_SearchOrganizationInsights(cfg, client)
			return
		}
		if _devopsguruStartCostEstimation {
			devopsguru_StartCostEstimation(cfg, client)
			return
		}
		if _devopsguruUpdateEventSourcesConfig {
			devopsguru_UpdateEventSourcesConfig(cfg, client)
			return
		}
		if _devopsguruUpdateResourceCollection {
			devopsguru_UpdateResourceCollection(cfg, client)
			return
		}
		if _devopsguruUpdateServiceIntegration {
			devopsguru_UpdateServiceIntegration(cfg, client)
			return
		}

	},
}

var (
	_devopsguruAddNotificationChannel                       bool
	_devopsguruDeleteInsight                                bool
	_devopsguruDescribeAccountHealth                        bool
	_devopsguruDescribeAccountOverview                      bool
	_devopsguruDescribeAnomaly                              bool
	_devopsguruDescribeEventSourcesConfig                   bool
	_devopsguruDescribeFeedback                             bool
	_devopsguruDescribeInsight                              bool
	_devopsguruDescribeOrganizationHealth                   bool
	_devopsguruDescribeOrganizationOverview                 bool
	_devopsguruDescribeOrganizationResourceCollectionHealth bool
	_devopsguruDescribeResourceCollectionHealth             bool
	_devopsguruDescribeServiceIntegration                   bool
	_devopsguruGetCostEstimation                            bool
	_devopsguruGetResourceCollection                        bool
	_devopsguruListAnomaliesForInsight                      bool
	_devopsguruListAnomalousLogGroups                       bool
	_devopsguruListEvents                                   bool
	_devopsguruListInsights                                 bool
	_devopsguruListMonitoredResources                       bool
	_devopsguruListNotificationChannels                     bool
	_devopsguruListOrganizationInsights                     bool
	_devopsguruListRecommendations                          bool
	_devopsguruPutFeedback                                  bool
	_devopsguruRemoveNotificationChannel                    bool
	_devopsguruSearchInsights                               bool
	_devopsguruSearchOrganizationInsights                   bool
	_devopsguruStartCostEstimation                          bool
	_devopsguruUpdateEventSourcesConfig                     bool
	_devopsguruUpdateResourceCollection                     bool
	_devopsguruUpdateServiceIntegration                     bool

	_devopsguruAccountId                          string
	_devopsguruAccountIds                         []string
	_devopsguruAction                             string
	_devopsguruClientToken                        string
	_devopsguruConfig                             string
	_devopsguruEventSources                       string
	_devopsguruFilters                            string
	_devopsguruFromTime                           string
	_devopsguruId                                 string
	_devopsguruInsightFeedback                    string
	_devopsguruInsightId                          string
	_devopsguruLocale                             string
	_devopsguruMaxResults                         string
	_devopsguruNextToken                          string
	_devopsguruOrganizationResourceCollectionType string
	_devopsguruOrganizationalUnitIds              []string
	_devopsguruResourceCollection                 string
	_devopsguruResourceCollectionType             string
	_devopsguruServiceIntegration                 string
	_devopsguruStartTimeRange                     string
	_devopsguruStatusFilter                       string
	_devopsguruToTime                             string
	_devopsguruType                               string
)

// Adds a notification channel to DevOps Guru. A notification channel is used to
// notify you about important DevOps Guru events, such as when an insight is
// generated.
//
// If you use an Amazon SNS topic in another account, you must attach a policy to
// it that grants DevOps Guru permission to send it notifications. DevOps Guru adds
// the required policy on your behalf to send notifications using Amazon SNS in
// your account. DevOps Guru only supports standard SNS topics. For more
// information, see [Permissions for Amazon SNS topics].
//
// If you use an Amazon SNS topic that is encrypted by an Amazon Web Services Key
// Management Service customer-managed key (CMK), then you must add permissions to
// the CMK. For more information, see [Permissions for Amazon Web Services KMS–encrypted Amazon SNS topics].
//
// [Permissions for Amazon SNS topics]: https://docs.aws.amazon.com/devops-guru/latest/userguide/sns-required-permissions.html
// [Permissions for Amazon Web Services KMS–encrypted Amazon SNS topics]: https://docs.aws.amazon.com/devops-guru/latest/userguide/sns-kms-permissions.html
func devopsguru_AddNotificationChannel(cfg aws.Config, client *devopsguru.Client) {
	input := &devopsguru.AddNotificationChannelInput{
		// Config: *types.NotificationChannelConfig, // Required
	}

	if len(_devopsguruConfig) > 0 {
		if err := assignInputField(input, "Config", _devopsguruConfig); err != nil {
			log.Errorf("invalid --config: %s", err.Error())
			return
		}
	}

	if resp, err := client.AddNotificationChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the insight along with the associated anomalies, events and
// recommendations.
func devopsguru_DeleteInsight(cfg aws.Config, client *devopsguru.Client) {
	input := &devopsguru.DeleteInsightInput{
		// Id: *string, // Required
	}

	if len(_devopsguruId) > 0 {
		input.Id = aws.String(_devopsguruId)
	}

	if resp, err := client.DeleteInsight(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the number of open reactive insights, the number of open proactive
// insights, and the number of metrics analyzed in your Amazon Web Services
// account. Use these numbers to gauge the health of operations in your Amazon Web
// Services account.
func devopsguru_DescribeAccountHealth(cfg aws.Config, client *devopsguru.Client) {
	input := &devopsguru.DescribeAccountHealthInput{}

	if resp, err := client.DescribeAccountHealth(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// For the time range passed in, returns the number of open reactive insight that
// were created, the number of open proactive insights that were created, and the
// Mean Time to Recover (MTTR) for all closed reactive insights.
func devopsguru_DescribeAccountOverview(cfg aws.Config, client *devopsguru.Client) {
	input := &devopsguru.DescribeAccountOverviewInput{
		// FromTime: *time.Time, // Required
	}

	if len(_devopsguruFromTime) > 0 {
		if err := assignInputField(input, "FromTime", _devopsguruFromTime); err != nil {
			log.Errorf("invalid --from-time: %s", err.Error())
			return
		}
	}
	if len(_devopsguruToTime) > 0 {
		if err := assignInputField(input, "ToTime", _devopsguruToTime); err != nil {
			log.Errorf("invalid --to-time: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeAccountOverview(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns details about an anomaly that you specify using its ID.
func devopsguru_DescribeAnomaly(cfg aws.Config, client *devopsguru.Client) {
	input := &devopsguru.DescribeAnomalyInput{
		// Id: *string, // Required
	}

	if len(_devopsguruId) > 0 {
		input.Id = aws.String(_devopsguruId)
	}
	if len(_devopsguruAccountId) > 0 {
		input.AccountId = aws.String(_devopsguruAccountId)
	}

	if resp, err := client.DescribeAnomaly(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the integration status of services that are integrated with DevOps Guru
// as Consumer via EventBridge. The one service that can be integrated with DevOps
// Guru is Amazon CodeGuru Profiler, which can produce proactive recommendations
// which can be stored and viewed in DevOps Guru.
func devopsguru_DescribeEventSourcesConfig(cfg aws.Config, client *devopsguru.Client) {
	input := &devopsguru.DescribeEventSourcesConfigInput{}

	if resp, err := client.DescribeEventSourcesConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the most recent feedback submitted in the current Amazon Web Services
// account and Region.
func devopsguru_DescribeFeedback(cfg aws.Config, client *devopsguru.Client) {
	input := &devopsguru.DescribeFeedbackInput{}

	if len(_devopsguruInsightId) > 0 {
		input.InsightId = aws.String(_devopsguruInsightId)
	}

	if resp, err := client.DescribeFeedback(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns details about an insight that you specify using its ID.
func devopsguru_DescribeInsight(cfg aws.Config, client *devopsguru.Client) {
	input := &devopsguru.DescribeInsightInput{
		// Id: *string, // Required
	}

	if len(_devopsguruId) > 0 {
		input.Id = aws.String(_devopsguruId)
	}
	if len(_devopsguruAccountId) > 0 {
		input.AccountId = aws.String(_devopsguruAccountId)
	}

	if resp, err := client.DescribeInsight(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns active insights, predictive insights, and resource hours analyzed in
// last hour.
func devopsguru_DescribeOrganizationHealth(cfg aws.Config, client *devopsguru.Client) {
	input := &devopsguru.DescribeOrganizationHealthInput{}

	if len(_devopsguruAccountIds) > 0 {
		input.AccountIds = append([]string(nil), _devopsguruAccountIds...)
	}
	if len(_devopsguruOrganizationalUnitIds) > 0 {
		input.OrganizationalUnitIds = append([]string(nil), _devopsguruOrganizationalUnitIds...)
	}

	if resp, err := client.DescribeOrganizationHealth(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns an overview of your organization's history based on the specified time
// range. The overview includes the total reactive and proactive insights.
func devopsguru_DescribeOrganizationOverview(cfg aws.Config, client *devopsguru.Client) {
	input := &devopsguru.DescribeOrganizationOverviewInput{
		// FromTime: *time.Time, // Required
	}

	if len(_devopsguruFromTime) > 0 {
		if err := assignInputField(input, "FromTime", _devopsguruFromTime); err != nil {
			log.Errorf("invalid --from-time: %s", err.Error())
			return
		}
	}
	if len(_devopsguruAccountIds) > 0 {
		input.AccountIds = append([]string(nil), _devopsguruAccountIds...)
	}
	if len(_devopsguruOrganizationalUnitIds) > 0 {
		input.OrganizationalUnitIds = append([]string(nil), _devopsguruOrganizationalUnitIds...)
	}
	if len(_devopsguruToTime) > 0 {
		if err := assignInputField(input, "ToTime", _devopsguruToTime); err != nil {
			log.Errorf("invalid --to-time: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeOrganizationOverview(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides an overview of your system's health. If additional member accounts are
// part of your organization, you can filter those accounts using the AccountIds
// field.
func devopsguru_DescribeOrganizationResourceCollectionHealth(cfg aws.Config, client *devopsguru.Client) {
	input := &devopsguru.DescribeOrganizationResourceCollectionHealthInput{
		// OrganizationResourceCollectionType: types.OrganizationResourceCollectionType, // Required
	}

	if len(_devopsguruOrganizationResourceCollectionType) > 0 {
		if err := assignInputField(input, "OrganizationResourceCollectionType", _devopsguruOrganizationResourceCollectionType); err != nil {
			log.Errorf("invalid --organization-resource-collection-type: %s", err.Error())
			return
		}
	}
	if len(_devopsguruAccountIds) > 0 {
		input.AccountIds = append([]string(nil), _devopsguruAccountIds...)
	}
	if len(_devopsguruMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _devopsguruMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_devopsguruNextToken) > 0 {
		input.NextToken = aws.String(_devopsguruNextToken)
	}
	if len(_devopsguruOrganizationalUnitIds) > 0 {
		input.OrganizationalUnitIds = append([]string(nil), _devopsguruOrganizationalUnitIds...)
	}

	if disablePaginator() {
		if resp, err := client.DescribeOrganizationResourceCollectionHealth(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*devopsguru.DescribeOrganizationResourceCollectionHealthOutput
	p := devopsguru.NewDescribeOrganizationResourceCollectionHealthPaginator(client, input)
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

// Returns the number of open proactive insights, open reactive insights, and the
// Mean Time to Recover (MTTR) for all closed insights in resource collections in
// your account. You specify the type of Amazon Web Services resources collection.
// The two types of Amazon Web Services resource collections supported are Amazon
// Web Services CloudFormation stacks and Amazon Web Services resources that
// contain the same Amazon Web Services tag. DevOps Guru can be configured to
// analyze the Amazon Web Services resources that are defined in the stacks or that
// are tagged using the same tag key. You can specify up to 500 Amazon Web Services
// CloudFormation stacks.
func devopsguru_DescribeResourceCollectionHealth(cfg aws.Config, client *devopsguru.Client) {
	input := &devopsguru.DescribeResourceCollectionHealthInput{
		// ResourceCollectionType: types.ResourceCollectionType, // Required
	}

	if len(_devopsguruResourceCollectionType) > 0 {
		if err := assignInputField(input, "ResourceCollectionType", _devopsguruResourceCollectionType); err != nil {
			log.Errorf("invalid --resource-collection-type: %s", err.Error())
			return
		}
	}
	if len(_devopsguruNextToken) > 0 {
		input.NextToken = aws.String(_devopsguruNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeResourceCollectionHealth(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*devopsguru.DescribeResourceCollectionHealthOutput
	p := devopsguru.NewDescribeResourceCollectionHealthPaginator(client, input)
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

// Returns the integration status of services that are integrated with DevOps
// Guru. The one service that can be integrated with DevOps Guru is Amazon Web
// Services Systems Manager, which can be used to create an OpsItem for each
// generated insight.
func devopsguru_DescribeServiceIntegration(cfg aws.Config, client *devopsguru.Client) {
	input := &devopsguru.DescribeServiceIntegrationInput{}

	if resp, err := client.DescribeServiceIntegration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns an estimate of the monthly cost for DevOps Guru to analyze your Amazon
// Web Services resources. For more information, see [Estimate your Amazon DevOps Guru costs]and [Amazon DevOps Guru pricing].
//
// [Amazon DevOps Guru pricing]: http://aws.amazon.com/devops-guru/pricing/
// [Estimate your Amazon DevOps Guru costs]: https://docs.aws.amazon.com/devops-guru/latest/userguide/cost-estimate.html
func devopsguru_GetCostEstimation(cfg aws.Config, client *devopsguru.Client) {
	input := &devopsguru.GetCostEstimationInput{}

	if len(_devopsguruNextToken) > 0 {
		input.NextToken = aws.String(_devopsguruNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetCostEstimation(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*devopsguru.GetCostEstimationOutput
	p := devopsguru.NewGetCostEstimationPaginator(client, input)
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

// Returns lists Amazon Web Services resources that are of the specified resource
// collection type. The two types of Amazon Web Services resource collections
// supported are Amazon Web Services CloudFormation stacks and Amazon Web Services
// resources that contain the same Amazon Web Services tag. DevOps Guru can be
// configured to analyze the Amazon Web Services resources that are defined in the
// stacks or that are tagged using the same tag key. You can specify up to 500
// Amazon Web Services CloudFormation stacks.
func devopsguru_GetResourceCollection(cfg aws.Config, client *devopsguru.Client) {
	input := &devopsguru.GetResourceCollectionInput{
		// ResourceCollectionType: types.ResourceCollectionType, // Required
	}

	if len(_devopsguruResourceCollectionType) > 0 {
		if err := assignInputField(input, "ResourceCollectionType", _devopsguruResourceCollectionType); err != nil {
			log.Errorf("invalid --resource-collection-type: %s", err.Error())
			return
		}
	}
	if len(_devopsguruNextToken) > 0 {
		input.NextToken = aws.String(_devopsguruNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetResourceCollection(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*devopsguru.GetResourceCollectionOutput
	p := devopsguru.NewGetResourceCollectionPaginator(client, input)
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

// Returns a list of the anomalies that belong to an insight that you specify
// using its ID.
func devopsguru_ListAnomaliesForInsight(cfg aws.Config, client *devopsguru.Client) {
	input := &devopsguru.ListAnomaliesForInsightInput{
		// InsightId: *string, // Required
	}

	if len(_devopsguruInsightId) > 0 {
		input.InsightId = aws.String(_devopsguruInsightId)
	}
	if len(_devopsguruAccountId) > 0 {
		input.AccountId = aws.String(_devopsguruAccountId)
	}
	if len(_devopsguruFilters) > 0 {
		if err := assignInputField(input, "Filters", _devopsguruFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_devopsguruMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _devopsguruMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_devopsguruNextToken) > 0 {
		input.NextToken = aws.String(_devopsguruNextToken)
	}
	if len(_devopsguruStartTimeRange) > 0 {
		if err := assignInputField(input, "StartTimeRange", _devopsguruStartTimeRange); err != nil {
			log.Errorf("invalid --start-time-range: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListAnomaliesForInsight(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*devopsguru.ListAnomaliesForInsightOutput
	p := devopsguru.NewListAnomaliesForInsightPaginator(client, input)
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

// Returns the list of log groups that contain log anomalies.
func devopsguru_ListAnomalousLogGroups(cfg aws.Config, client *devopsguru.Client) {
	input := &devopsguru.ListAnomalousLogGroupsInput{
		// InsightId: *string, // Required
	}

	if len(_devopsguruInsightId) > 0 {
		input.InsightId = aws.String(_devopsguruInsightId)
	}
	if len(_devopsguruMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _devopsguruMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_devopsguruNextToken) > 0 {
		input.NextToken = aws.String(_devopsguruNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAnomalousLogGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*devopsguru.ListAnomalousLogGroupsOutput
	p := devopsguru.NewListAnomalousLogGroupsPaginator(client, input)
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

// Returns a list of the events emitted by the resources that are evaluated by
// DevOps Guru. You can use filters to specify which events are returned.
func devopsguru_ListEvents(cfg aws.Config, client *devopsguru.Client) {
	input := &devopsguru.ListEventsInput{
		// Filters: *types.ListEventsFilters, // Required
	}

	if len(_devopsguruFilters) > 0 {
		if err := assignInputField(input, "Filters", _devopsguruFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_devopsguruAccountId) > 0 {
		input.AccountId = aws.String(_devopsguruAccountId)
	}
	if len(_devopsguruMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _devopsguruMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_devopsguruNextToken) > 0 {
		input.NextToken = aws.String(_devopsguruNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListEvents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*devopsguru.ListEventsOutput
	p := devopsguru.NewListEventsPaginator(client, input)
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

// Returns a list of insights in your Amazon Web Services account. You can
// specify which insights are returned by their start time and status ( ONGOING ,
// CLOSED , or ANY ).
func devopsguru_ListInsights(cfg aws.Config, client *devopsguru.Client) {
	input := &devopsguru.ListInsightsInput{
		// StatusFilter: *types.ListInsightsStatusFilter, // Required
	}

	if len(_devopsguruStatusFilter) > 0 {
		if err := assignInputField(input, "StatusFilter", _devopsguruStatusFilter); err != nil {
			log.Errorf("invalid --status-filter: %s", err.Error())
			return
		}
	}
	if len(_devopsguruMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _devopsguruMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_devopsguruNextToken) > 0 {
		input.NextToken = aws.String(_devopsguruNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListInsights(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*devopsguru.ListInsightsOutput
	p := devopsguru.NewListInsightsPaginator(client, input)
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

// Returns the list of all log groups that are being monitored and tagged by
// DevOps Guru.
func devopsguru_ListMonitoredResources(cfg aws.Config, client *devopsguru.Client) {
	input := &devopsguru.ListMonitoredResourcesInput{}

	if len(_devopsguruFilters) > 0 {
		if err := assignInputField(input, "Filters", _devopsguruFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_devopsguruMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _devopsguruMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_devopsguruNextToken) > 0 {
		input.NextToken = aws.String(_devopsguruNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListMonitoredResources(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*devopsguru.ListMonitoredResourcesOutput
	p := devopsguru.NewListMonitoredResourcesPaginator(client, input)
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

// Returns a list of notification channels configured for DevOps Guru. Each
// notification channel is used to notify you when DevOps Guru generates an insight
// that contains information about how to improve your operations. The one
// supported notification channel is Amazon Simple Notification Service (Amazon
// SNS).
func devopsguru_ListNotificationChannels(cfg aws.Config, client *devopsguru.Client) {
	input := &devopsguru.ListNotificationChannelsInput{}

	if len(_devopsguruNextToken) > 0 {
		input.NextToken = aws.String(_devopsguruNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListNotificationChannels(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*devopsguru.ListNotificationChannelsOutput
	p := devopsguru.NewListNotificationChannelsPaginator(client, input)
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

// Returns a list of insights associated with the account or OU Id.
func devopsguru_ListOrganizationInsights(cfg aws.Config, client *devopsguru.Client) {
	input := &devopsguru.ListOrganizationInsightsInput{
		// StatusFilter: *types.ListInsightsStatusFilter, // Required
	}

	if len(_devopsguruStatusFilter) > 0 {
		if err := assignInputField(input, "StatusFilter", _devopsguruStatusFilter); err != nil {
			log.Errorf("invalid --status-filter: %s", err.Error())
			return
		}
	}
	if len(_devopsguruAccountIds) > 0 {
		input.AccountIds = append([]string(nil), _devopsguruAccountIds...)
	}
	if len(_devopsguruMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _devopsguruMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_devopsguruNextToken) > 0 {
		input.NextToken = aws.String(_devopsguruNextToken)
	}
	if len(_devopsguruOrganizationalUnitIds) > 0 {
		input.OrganizationalUnitIds = append([]string(nil), _devopsguruOrganizationalUnitIds...)
	}

	if disablePaginator() {
		if resp, err := client.ListOrganizationInsights(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*devopsguru.ListOrganizationInsightsOutput
	p := devopsguru.NewListOrganizationInsightsPaginator(client, input)
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

// Returns a list of a specified insight's recommendations. Each recommendation
// includes a list of related metrics and a list of related events.
func devopsguru_ListRecommendations(cfg aws.Config, client *devopsguru.Client) {
	input := &devopsguru.ListRecommendationsInput{
		// InsightId: *string, // Required
	}

	if len(_devopsguruInsightId) > 0 {
		input.InsightId = aws.String(_devopsguruInsightId)
	}
	if len(_devopsguruAccountId) > 0 {
		input.AccountId = aws.String(_devopsguruAccountId)
	}
	if len(_devopsguruLocale) > 0 {
		if err := assignInputField(input, "Locale", _devopsguruLocale); err != nil {
			log.Errorf("invalid --locale: %s", err.Error())
			return
		}
	}
	if len(_devopsguruNextToken) > 0 {
		input.NextToken = aws.String(_devopsguruNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRecommendations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*devopsguru.ListRecommendationsOutput
	p := devopsguru.NewListRecommendationsPaginator(client, input)
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

// Collects customer feedback about the specified insight.
func devopsguru_PutFeedback(cfg aws.Config, client *devopsguru.Client) {
	input := &devopsguru.PutFeedbackInput{}

	if len(_devopsguruInsightFeedback) > 0 {
		if err := assignInputField(input, "InsightFeedback", _devopsguruInsightFeedback); err != nil {
			log.Errorf("invalid --insight-feedback: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutFeedback(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a notification channel from DevOps Guru. A notification channel is
// used to notify you when DevOps Guru generates an insight that contains
// information about how to improve your operations.
func devopsguru_RemoveNotificationChannel(cfg aws.Config, client *devopsguru.Client) {
	input := &devopsguru.RemoveNotificationChannelInput{
		// Id: *string, // Required
	}

	if len(_devopsguruId) > 0 {
		input.Id = aws.String(_devopsguruId)
	}

	if resp, err := client.RemoveNotificationChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of insights in your Amazon Web Services account. You can
// specify which insights are returned by their start time, one or more statuses (
// ONGOING or CLOSED ), one or more severities ( LOW , MEDIUM , and HIGH ), and
// type ( REACTIVE or PROACTIVE ).
//
// Use the Filters parameter to specify status and severity search parameters. Use
// the Type parameter to specify REACTIVE or PROACTIVE in your search.
func devopsguru_SearchInsights(cfg aws.Config, client *devopsguru.Client) {
	input := &devopsguru.SearchInsightsInput{
		// StartTimeRange: *types.StartTimeRange, // Required
		// Type: types.InsightType, // Required
	}

	if len(_devopsguruStartTimeRange) > 0 {
		if err := assignInputField(input, "StartTimeRange", _devopsguruStartTimeRange); err != nil {
			log.Errorf("invalid --start-time-range: %s", err.Error())
			return
		}
	}
	if len(_devopsguruType) > 0 {
		if err := assignInputField(input, "Type", _devopsguruType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_devopsguruFilters) > 0 {
		if err := assignInputField(input, "Filters", _devopsguruFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_devopsguruMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _devopsguruMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_devopsguruNextToken) > 0 {
		input.NextToken = aws.String(_devopsguruNextToken)
	}

	if disablePaginator() {
		if resp, err := client.SearchInsights(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*devopsguru.SearchInsightsOutput
	p := devopsguru.NewSearchInsightsPaginator(client, input)
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

// Returns a list of insights in your organization. You can specify which
// insights are returned by their start time, one or more statuses ( ONGOING ,
// CLOSED , and CLOSED ), one or more severities ( LOW , MEDIUM , and HIGH ), and
// type ( REACTIVE or PROACTIVE ).
//
// Use the Filters parameter to specify status and severity search parameters. Use
// the Type parameter to specify REACTIVE or PROACTIVE in your search.
func devopsguru_SearchOrganizationInsights(cfg aws.Config, client *devopsguru.Client) {
	input := &devopsguru.SearchOrganizationInsightsInput{
		// AccountIds: []string, // Required
		// StartTimeRange: *types.StartTimeRange, // Required
		// Type: types.InsightType, // Required
	}

	if len(_devopsguruAccountIds) > 0 {
		input.AccountIds = append([]string(nil), _devopsguruAccountIds...)
	}
	if len(_devopsguruStartTimeRange) > 0 {
		if err := assignInputField(input, "StartTimeRange", _devopsguruStartTimeRange); err != nil {
			log.Errorf("invalid --start-time-range: %s", err.Error())
			return
		}
	}
	if len(_devopsguruType) > 0 {
		if err := assignInputField(input, "Type", _devopsguruType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_devopsguruFilters) > 0 {
		if err := assignInputField(input, "Filters", _devopsguruFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_devopsguruMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _devopsguruMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_devopsguruNextToken) > 0 {
		input.NextToken = aws.String(_devopsguruNextToken)
	}

	if disablePaginator() {
		if resp, err := client.SearchOrganizationInsights(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*devopsguru.SearchOrganizationInsightsOutput
	p := devopsguru.NewSearchOrganizationInsightsPaginator(client, input)
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

// Starts the creation of an estimate of the monthly cost to analyze your Amazon
// Web Services resources.
func devopsguru_StartCostEstimation(cfg aws.Config, client *devopsguru.Client) {
	input := &devopsguru.StartCostEstimationInput{
		// ResourceCollection: *types.CostEstimationResourceCollectionFilter, // Required
	}

	if len(_devopsguruResourceCollection) > 0 {
		if err := assignInputField(input, "ResourceCollection", _devopsguruResourceCollection); err != nil {
			log.Errorf("invalid --resource-collection: %s", err.Error())
			return
		}
	}
	if len(_devopsguruClientToken) > 0 {
		input.ClientToken = aws.String(_devopsguruClientToken)
	}

	if resp, err := client.StartCostEstimation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables or disables integration with a service that can be integrated with
// DevOps Guru. The one service that can be integrated with DevOps Guru is Amazon
// CodeGuru Profiler, which can produce proactive recommendations which can be
// stored and viewed in DevOps Guru.
func devopsguru_UpdateEventSourcesConfig(cfg aws.Config, client *devopsguru.Client) {
	input := &devopsguru.UpdateEventSourcesConfigInput{}

	if len(_devopsguruEventSources) > 0 {
		if err := assignInputField(input, "EventSources", _devopsguruEventSources); err != nil {
			log.Errorf("invalid --event-sources: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateEventSourcesConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the collection of resources that DevOps Guru analyzes. The two types
// of Amazon Web Services resource collections supported are Amazon Web Services
// CloudFormation stacks and Amazon Web Services resources that contain the same
// Amazon Web Services tag. DevOps Guru can be configured to analyze the Amazon Web
// Services resources that are defined in the stacks or that are tagged using the
// same tag key. You can specify up to 500 Amazon Web Services CloudFormation
// stacks. This method also creates the IAM role required for you to use DevOps
// Guru.
func devopsguru_UpdateResourceCollection(cfg aws.Config, client *devopsguru.Client) {
	input := &devopsguru.UpdateResourceCollectionInput{
		// Action: types.UpdateResourceCollectionAction, // Required
		// ResourceCollection: *types.UpdateResourceCollectionFilter, // Required
	}

	if len(_devopsguruAction) > 0 {
		if err := assignInputField(input, "Action", _devopsguruAction); err != nil {
			log.Errorf("invalid --action: %s", err.Error())
			return
		}
	}
	if len(_devopsguruResourceCollection) > 0 {
		if err := assignInputField(input, "ResourceCollection", _devopsguruResourceCollection); err != nil {
			log.Errorf("invalid --resource-collection: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateResourceCollection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables or disables integration with a service that can be integrated with
// DevOps Guru. The one service that can be integrated with DevOps Guru is Amazon
// Web Services Systems Manager, which can be used to create an OpsItem for each
// generated insight.
func devopsguru_UpdateServiceIntegration(cfg aws.Config, client *devopsguru.Client) {
	input := &devopsguru.UpdateServiceIntegrationInput{
		// ServiceIntegration: *types.UpdateServiceIntegrationConfig, // Required
	}

	if len(_devopsguruServiceIntegration) > 0 {
		if err := assignInputField(input, "ServiceIntegration", _devopsguruServiceIntegration); err != nil {
			log.Errorf("invalid --service-integration: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateServiceIntegration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_devopsguruCmd)
	_devopsguruCmd.Flags().SortFlags = false

	_devopsguruCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_devopsguruCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_devopsguruCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_devopsguruCmd.Flags().StringVarP(&_devopsguruAccountId, "account-id", "", "", "Account ID")
	_devopsguruCmd.Flags().StringSliceVarP(&_devopsguruAccountIds, "account-ids", "", nil, "Account Ids")
	_devopsguruCmd.Flags().StringVarP(&_devopsguruAction, "action", "", "", "Action")
	_devopsguruCmd.Flags().StringVarP(&_devopsguruClientToken, "client-token", "", "", "Client Token")
	_devopsguruCmd.Flags().StringVarP(&_devopsguruConfig, "config", "", "", "Config")
	_devopsguruCmd.Flags().StringVarP(&_devopsguruEventSources, "event-sources", "", "", "Event Sources")
	_devopsguruCmd.Flags().StringVarP(&_devopsguruFilters, "filters", "", "", "Filters")
	_devopsguruCmd.Flags().StringVarP(&_devopsguruFromTime, "from-time", "", "", "From Time")
	_devopsguruCmd.Flags().StringVarP(&_devopsguruId, "id", "", "", "ID")
	_devopsguruCmd.Flags().StringVarP(&_devopsguruInsightFeedback, "insight-feedback", "", "", "Insight Feedback")
	_devopsguruCmd.Flags().StringVarP(&_devopsguruInsightId, "insight-id", "", "", "Insight ID")
	_devopsguruCmd.Flags().StringVarP(&_devopsguruLocale, "locale", "", "", "Locale")
	_devopsguruCmd.Flags().StringVarP(&_devopsguruMaxResults, "max-results", "", "", "Max Results")
	_devopsguruCmd.Flags().StringVarP(&_devopsguruNextToken, "next-token", "", "", "Next Token")
	_devopsguruCmd.Flags().StringVarP(&_devopsguruOrganizationResourceCollectionType, "organization-resource-collection-type", "", "", "Organization Resource Collection Type")
	_devopsguruCmd.Flags().StringSliceVarP(&_devopsguruOrganizationalUnitIds, "organizational-unit-ids", "", nil, "Organizational Unit Ids")
	_devopsguruCmd.Flags().StringVarP(&_devopsguruResourceCollection, "resource-collection", "", "", "Resource Collection")
	_devopsguruCmd.Flags().StringVarP(&_devopsguruResourceCollectionType, "resource-collection-type", "", "", "Resource Collection Type")
	_devopsguruCmd.Flags().StringVarP(&_devopsguruServiceIntegration, "service-integration", "", "", "Service Integration")
	_devopsguruCmd.Flags().StringVarP(&_devopsguruStartTimeRange, "start-time-range", "", "", "Start Time Range")
	_devopsguruCmd.Flags().StringVarP(&_devopsguruStatusFilter, "status-filter", "", "", "Status Filter")
	_devopsguruCmd.Flags().StringVarP(&_devopsguruToTime, "to-time", "", "", "To Time")
	_devopsguruCmd.Flags().StringVarP(&_devopsguruType, "type", "", "", "Type")

	_devopsguruCmd.Flags().BoolVarP(&_devopsguruAddNotificationChannel, "add-notification-channel", "", false, "Add Notification Channel")
	_devopsguruCmd.Flags().BoolVarP(&_devopsguruDeleteInsight, "delete-insight", "", false, "Delete Insight")
	_devopsguruCmd.Flags().BoolVarP(&_devopsguruDescribeAccountHealth, "describe-account-health", "", false, "Describe Account Health")
	_devopsguruCmd.Flags().BoolVarP(&_devopsguruDescribeAccountOverview, "describe-account-overview", "", false, "Describe Account Overview")
	_devopsguruCmd.Flags().BoolVarP(&_devopsguruDescribeAnomaly, "describe-anomaly", "", false, "Describe Anomaly")
	_devopsguruCmd.Flags().BoolVarP(&_devopsguruDescribeEventSourcesConfig, "describe-event-sources-config", "", false, "Describe Event Sources Config")
	_devopsguruCmd.Flags().BoolVarP(&_devopsguruDescribeFeedback, "describe-feedback", "", false, "Describe Feedback")
	_devopsguruCmd.Flags().BoolVarP(&_devopsguruDescribeInsight, "describe-insight", "", false, "Describe Insight")
	_devopsguruCmd.Flags().BoolVarP(&_devopsguruDescribeOrganizationHealth, "describe-organization-health", "", false, "Describe Organization Health")
	_devopsguruCmd.Flags().BoolVarP(&_devopsguruDescribeOrganizationOverview, "describe-organization-overview", "", false, "Describe Organization Overview")
	_devopsguruCmd.Flags().BoolVarP(&_devopsguruDescribeOrganizationResourceCollectionHealth, "describe-organization-resource-collection-health", "", false, "Describe Organization Resource Collection Health")
	_devopsguruCmd.Flags().BoolVarP(&_devopsguruDescribeResourceCollectionHealth, "describe-resource-collection-health", "", false, "Describe Resource Collection Health")
	_devopsguruCmd.Flags().BoolVarP(&_devopsguruDescribeServiceIntegration, "describe-service-integration", "", false, "Describe Service Integration")
	_devopsguruCmd.Flags().BoolVarP(&_devopsguruGetCostEstimation, "get-cost-estimation", "", false, "Get Cost Estimation")
	_devopsguruCmd.Flags().BoolVarP(&_devopsguruGetResourceCollection, "get-resource-collection", "", false, "Get Resource Collection")
	_devopsguruCmd.Flags().BoolVarP(&_devopsguruListAnomaliesForInsight, "list-anomalies-for-insight", "", false, "List Anomalies For Insight")
	_devopsguruCmd.Flags().BoolVarP(&_devopsguruListAnomalousLogGroups, "list-anomalous-log-groups", "", false, "List Anomalous Log Groups")
	_devopsguruCmd.Flags().BoolVarP(&_devopsguruListEvents, "list-events", "", false, "List Events")
	_devopsguruCmd.Flags().BoolVarP(&_devopsguruListInsights, "list-insights", "", false, "List Insights")
	_devopsguruCmd.Flags().BoolVarP(&_devopsguruListMonitoredResources, "list-monitored-resources", "", false, "List Monitored Resources")
	_devopsguruCmd.Flags().BoolVarP(&_devopsguruListNotificationChannels, "list-notification-channels", "", false, "List Notification Channels")
	_devopsguruCmd.Flags().BoolVarP(&_devopsguruListOrganizationInsights, "list-organization-insights", "", false, "List Organization Insights")
	_devopsguruCmd.Flags().BoolVarP(&_devopsguruListRecommendations, "list-recommendations", "", false, "List Recommendations")
	_devopsguruCmd.Flags().BoolVarP(&_devopsguruPutFeedback, "put-feedback", "", false, "Put Feedback")
	_devopsguruCmd.Flags().BoolVarP(&_devopsguruRemoveNotificationChannel, "remove-notification-channel", "", false, "Remove Notification Channel")
	_devopsguruCmd.Flags().BoolVarP(&_devopsguruSearchInsights, "search-insights", "", false, "Search Insights")
	_devopsguruCmd.Flags().BoolVarP(&_devopsguruSearchOrganizationInsights, "search-organization-insights", "", false, "Search Organization Insights")
	_devopsguruCmd.Flags().BoolVarP(&_devopsguruStartCostEstimation, "start-cost-estimation", "", false, "Start Cost Estimation")
	_devopsguruCmd.Flags().BoolVarP(&_devopsguruUpdateEventSourcesConfig, "update-event-sources-config", "", false, "Update Event Sources Config")
	_devopsguruCmd.Flags().BoolVarP(&_devopsguruUpdateResourceCollection, "update-resource-collection", "", false, "Update Resource Collection")
	_devopsguruCmd.Flags().BoolVarP(&_devopsguruUpdateServiceIntegration, "update-service-integration", "", false, "Update Service Integration")

}
