package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// costexplorerCmd represents the costexplorer command
var _costexplorerCmd = &cobra.Command{
	Use:   "costexplorer",
	Short: "AWS costexplorer CLI",
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
		client := costexplorer.NewFromConfig(cfg)
		if _costexplorerCreateAnomalyMonitor {
			costexplorer_CreateAnomalyMonitor(cfg, client)
			return
		}
		if _costexplorerCreateAnomalySubscription {
			costexplorer_CreateAnomalySubscription(cfg, client)
			return
		}
		if _costexplorerCreateCostCategoryDefinition {
			costexplorer_CreateCostCategoryDefinition(cfg, client)
			return
		}
		if _costexplorerDeleteAnomalyMonitor {
			costexplorer_DeleteAnomalyMonitor(cfg, client)
			return
		}
		if _costexplorerDeleteAnomalySubscription {
			costexplorer_DeleteAnomalySubscription(cfg, client)
			return
		}
		if _costexplorerDeleteCostCategoryDefinition {
			costexplorer_DeleteCostCategoryDefinition(cfg, client)
			return
		}
		if _costexplorerDescribeCostCategoryDefinition {
			costexplorer_DescribeCostCategoryDefinition(cfg, client)
			return
		}
		if _costexplorerGetAnomalies {
			costexplorer_GetAnomalies(cfg, client)
			return
		}
		if _costexplorerGetAnomalyMonitors {
			costexplorer_GetAnomalyMonitors(cfg, client)
			return
		}
		if _costexplorerGetAnomalySubscriptions {
			costexplorer_GetAnomalySubscriptions(cfg, client)
			return
		}
		if _costexplorerGetApproximateUsageRecords {
			costexplorer_GetApproximateUsageRecords(cfg, client)
			return
		}
		if _costexplorerGetCommitmentPurchaseAnalysis {
			costexplorer_GetCommitmentPurchaseAnalysis(cfg, client)
			return
		}
		if _costexplorerGetCostAndUsage {
			costexplorer_GetCostAndUsage(cfg, client)
			return
		}
		if _costexplorerGetCostAndUsageComparisons {
			costexplorer_GetCostAndUsageComparisons(cfg, client)
			return
		}
		if _costexplorerGetCostAndUsageWithResources {
			costexplorer_GetCostAndUsageWithResources(cfg, client)
			return
		}
		if _costexplorerGetCostCategories {
			costexplorer_GetCostCategories(cfg, client)
			return
		}
		if _costexplorerGetCostComparisonDrivers {
			costexplorer_GetCostComparisonDrivers(cfg, client)
			return
		}
		if _costexplorerGetCostForecast {
			costexplorer_GetCostForecast(cfg, client)
			return
		}
		if _costexplorerGetDimensionValues {
			costexplorer_GetDimensionValues(cfg, client)
			return
		}
		if _costexplorerGetReservationCoverage {
			costexplorer_GetReservationCoverage(cfg, client)
			return
		}
		if _costexplorerGetReservationPurchaseRecommendation {
			costexplorer_GetReservationPurchaseRecommendation(cfg, client)
			return
		}
		if _costexplorerGetReservationUtilization {
			costexplorer_GetReservationUtilization(cfg, client)
			return
		}
		if _costexplorerGetRightsizingRecommendation {
			costexplorer_GetRightsizingRecommendation(cfg, client)
			return
		}
		if _costexplorerGetSavingsPlanPurchaseRecommendationDetails {
			costexplorer_GetSavingsPlanPurchaseRecommendationDetails(cfg, client)
			return
		}
		if _costexplorerGetSavingsPlansCoverage {
			costexplorer_GetSavingsPlansCoverage(cfg, client)
			return
		}
		if _costexplorerGetSavingsPlansPurchaseRecommendation {
			costexplorer_GetSavingsPlansPurchaseRecommendation(cfg, client)
			return
		}
		if _costexplorerGetSavingsPlansUtilization {
			costexplorer_GetSavingsPlansUtilization(cfg, client)
			return
		}
		if _costexplorerGetSavingsPlansUtilizationDetails {
			costexplorer_GetSavingsPlansUtilizationDetails(cfg, client)
			return
		}
		if _costexplorerGetTags {
			costexplorer_GetTags(cfg, client)
			return
		}
		if _costexplorerGetUsageForecast {
			costexplorer_GetUsageForecast(cfg, client)
			return
		}
		if _costexplorerListCommitmentPurchaseAnalyses {
			costexplorer_ListCommitmentPurchaseAnalyses(cfg, client)
			return
		}
		if _costexplorerListCostAllocationTagBackfillHistory {
			costexplorer_ListCostAllocationTagBackfillHistory(cfg, client)
			return
		}
		if _costexplorerListCostAllocationTags {
			costexplorer_ListCostAllocationTags(cfg, client)
			return
		}
		if _costexplorerListCostCategoryDefinitions {
			costexplorer_ListCostCategoryDefinitions(cfg, client)
			return
		}
		if _costexplorerListCostCategoryResourceAssociations {
			costexplorer_ListCostCategoryResourceAssociations(cfg, client)
			return
		}
		if _costexplorerListSavingsPlansPurchaseRecommendationGeneration {
			costexplorer_ListSavingsPlansPurchaseRecommendationGeneration(cfg, client)
			return
		}
		if _costexplorerListTagsForResource {
			costexplorer_ListTagsForResource(cfg, client)
			return
		}
		if _costexplorerProvideAnomalyFeedback {
			costexplorer_ProvideAnomalyFeedback(cfg, client)
			return
		}
		if _costexplorerStartCommitmentPurchaseAnalysis {
			costexplorer_StartCommitmentPurchaseAnalysis(cfg, client)
			return
		}
		if _costexplorerStartCostAllocationTagBackfill {
			costexplorer_StartCostAllocationTagBackfill(cfg, client)
			return
		}
		if _costexplorerStartSavingsPlansPurchaseRecommendationGeneration {
			costexplorer_StartSavingsPlansPurchaseRecommendationGeneration(cfg, client)
			return
		}
		if _costexplorerTagResource {
			costexplorer_TagResource(cfg, client)
			return
		}
		if _costexplorerUntagResource {
			costexplorer_UntagResource(cfg, client)
			return
		}
		if _costexplorerUpdateAnomalyMonitor {
			costexplorer_UpdateAnomalyMonitor(cfg, client)
			return
		}
		if _costexplorerUpdateAnomalySubscription {
			costexplorer_UpdateAnomalySubscription(cfg, client)
			return
		}
		if _costexplorerUpdateCostAllocationTagsStatus {
			costexplorer_UpdateCostAllocationTagsStatus(cfg, client)
			return
		}
		if _costexplorerUpdateCostCategoryDefinition {
			costexplorer_UpdateCostCategoryDefinition(cfg, client)
			return
		}

	},
}

var (
	_costexplorerCreateAnomalyMonitor                              bool
	_costexplorerCreateAnomalySubscription                         bool
	_costexplorerCreateCostCategoryDefinition                      bool
	_costexplorerDeleteAnomalyMonitor                              bool
	_costexplorerDeleteAnomalySubscription                         bool
	_costexplorerDeleteCostCategoryDefinition                      bool
	_costexplorerDescribeCostCategoryDefinition                    bool
	_costexplorerGetAnomalies                                      bool
	_costexplorerGetAnomalyMonitors                                bool
	_costexplorerGetAnomalySubscriptions                           bool
	_costexplorerGetApproximateUsageRecords                        bool
	_costexplorerGetCommitmentPurchaseAnalysis                     bool
	_costexplorerGetCostAndUsage                                   bool
	_costexplorerGetCostAndUsageComparisons                        bool
	_costexplorerGetCostAndUsageWithResources                      bool
	_costexplorerGetCostCategories                                 bool
	_costexplorerGetCostComparisonDrivers                          bool
	_costexplorerGetCostForecast                                   bool
	_costexplorerGetDimensionValues                                bool
	_costexplorerGetReservationCoverage                            bool
	_costexplorerGetReservationPurchaseRecommendation              bool
	_costexplorerGetReservationUtilization                         bool
	_costexplorerGetRightsizingRecommendation                      bool
	_costexplorerGetSavingsPlanPurchaseRecommendationDetails       bool
	_costexplorerGetSavingsPlansCoverage                           bool
	_costexplorerGetSavingsPlansPurchaseRecommendation             bool
	_costexplorerGetSavingsPlansUtilization                        bool
	_costexplorerGetSavingsPlansUtilizationDetails                 bool
	_costexplorerGetTags                                           bool
	_costexplorerGetUsageForecast                                  bool
	_costexplorerListCommitmentPurchaseAnalyses                    bool
	_costexplorerListCostAllocationTagBackfillHistory              bool
	_costexplorerListCostAllocationTags                            bool
	_costexplorerListCostCategoryDefinitions                       bool
	_costexplorerListCostCategoryResourceAssociations              bool
	_costexplorerListSavingsPlansPurchaseRecommendationGeneration  bool
	_costexplorerListTagsForResource                               bool
	_costexplorerProvideAnomalyFeedback                            bool
	_costexplorerStartCommitmentPurchaseAnalysis                   bool
	_costexplorerStartCostAllocationTagBackfill                    bool
	_costexplorerStartSavingsPlansPurchaseRecommendationGeneration bool
	_costexplorerTagResource                                       bool
	_costexplorerUntagResource                                     bool
	_costexplorerUpdateAnomalyMonitor                              bool
	_costexplorerUpdateAnomalySubscription                         bool
	_costexplorerUpdateCostAllocationTagsStatus                    bool
	_costexplorerUpdateCostCategoryDefinition                      bool

	_costexplorerAccountId                               string
	_costexplorerAccountScope                            string
	_costexplorerAnalysisId                              string
	_costexplorerAnalysisIds                             []string
	_costexplorerAnalysisStatus                          string
	_costexplorerAnomalyId                               string
	_costexplorerAnomalyMonitor                          string
	_costexplorerAnomalySubscription                     string
	_costexplorerApproximationDimension                  string
	_costexplorerBackfillFrom                            string
	_costexplorerBaselineTimePeriod                      string
	_costexplorerBillingViewArn                          string
	_costexplorerCommitmentPurchaseAnalysisConfiguration string
	_costexplorerComparisonTimePeriod                    string
	_costexplorerConfiguration                           string
	_costexplorerContext                                 string
	_costexplorerCostAllocationTagsStatus                string
	_costexplorerCostCategoryArn                         string
	_costexplorerCostCategoryName                        string
	_costexplorerDataType                                string
	_costexplorerDateInterval                            string
	_costexplorerDefaultValue                            string
	_costexplorerDimension                               string
	_costexplorerEffectiveOn                             string
	_costexplorerEffectiveStart                          string
	_costexplorerFeedback                                string
	_costexplorerFilter                                  string
	_costexplorerFrequency                               string
	_costexplorerGenerationStatus                        string
	_costexplorerGranularity                             string
	_costexplorerGroupBy                                 string
	_costexplorerLookbackPeriodInDays                    string
	_costexplorerMaxResults                              string
	_costexplorerMetric                                  string
	_costexplorerMetricForComparison                     string
	_costexplorerMetrics                                 []string
	_costexplorerMonitorArn                              string
	_costexplorerMonitorArnList                          []string
	_costexplorerMonitorName                             string
	_costexplorerName                                    string
	_costexplorerNextPageToken                           string
	_costexplorerNextToken                               string
	_costexplorerPageSize                                string
	_costexplorerPaymentOption                           string
	_costexplorerPredictionIntervalLevel                 string
	_costexplorerRecommendationDetailId                  string
	_costexplorerRecommendationIds                       []string
	_costexplorerResourceArn                             string
	_costexplorerResourceTagKeys                         []string
	_costexplorerResourceTags                            string
	_costexplorerRuleVersion                             string
	_costexplorerRules                                   string
	_costexplorerSavingsPlansType                        string
	_costexplorerSearchString                            string
	_costexplorerService                                 string
	_costexplorerServiceSpecification                    string
	_costexplorerServices                                []string
	_costexplorerSortBy                                  string
	_costexplorerSplitChargeRules                        string
	_costexplorerStatus                                  string
	_costexplorerSubscribers                             string
	_costexplorerSubscriptionArn                         string
	_costexplorerSubscriptionArnList                     []string
	_costexplorerSubscriptionName                        string
	_costexplorerSupportedResourceTypes                  []string
	_costexplorerTagKey                                  string
	_costexplorerTagKeys                                 []string
	_costexplorerTermInYears                             string
	_costexplorerThreshold                               string
	_costexplorerThresholdExpression                     string
	_costexplorerTimePeriod                              string
	_costexplorerTotalImpact                             string
	_costexplorerType                                    string
)

// Creates a new cost anomaly detection monitor with the requested type and
// monitor specification.
func costexplorer_CreateAnomalyMonitor(cfg aws.Config, client *costexplorer.Client) {
	input := &costexplorer.CreateAnomalyMonitorInput{
		// AnomalyMonitor: *types.AnomalyMonitor, // Required
	}

	if len(_costexplorerAnomalyMonitor) > 0 {
		if err := assignInputField(input, "AnomalyMonitor", _costexplorerAnomalyMonitor); err != nil {
			log.Errorf("invalid --anomaly-monitor: %s", err.Error())
			return
		}
	}
	if len(_costexplorerResourceTags) > 0 {
		if err := assignInputField(input, "ResourceTags", _costexplorerResourceTags); err != nil {
			log.Errorf("invalid --resource-tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAnomalyMonitor(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds an alert subscription to a cost anomaly detection monitor. You can use
// each subscription to define subscribers with email or SNS notifications. Email
// subscribers can set an absolute or percentage threshold and a time frequency for
// receiving notifications.
func costexplorer_CreateAnomalySubscription(cfg aws.Config, client *costexplorer.Client) {
	input := &costexplorer.CreateAnomalySubscriptionInput{
		// AnomalySubscription: *types.AnomalySubscription, // Required
	}

	if len(_costexplorerAnomalySubscription) > 0 {
		if err := assignInputField(input, "AnomalySubscription", _costexplorerAnomalySubscription); err != nil {
			log.Errorf("invalid --anomaly-subscription: %s", err.Error())
			return
		}
	}
	if len(_costexplorerResourceTags) > 0 {
		if err := assignInputField(input, "ResourceTags", _costexplorerResourceTags); err != nil {
			log.Errorf("invalid --resource-tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAnomalySubscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new cost category with the requested name and rules.
func costexplorer_CreateCostCategoryDefinition(cfg aws.Config, client *costexplorer.Client) {
	input := &costexplorer.CreateCostCategoryDefinitionInput{
		// Name: *string, // Required
		// RuleVersion: types.CostCategoryRuleVersion, // Required
		// Rules: []types.CostCategoryRule, // Required
	}

	if len(_costexplorerName) > 0 {
		input.Name = aws.String(_costexplorerName)
	}
	if len(_costexplorerRuleVersion) > 0 {
		if err := assignInputField(input, "RuleVersion", _costexplorerRuleVersion); err != nil {
			log.Errorf("invalid --rule-version: %s", err.Error())
			return
		}
	}
	if len(_costexplorerRules) > 0 {
		if err := assignInputField(input, "Rules", _costexplorerRules); err != nil {
			log.Errorf("invalid --rules: %s", err.Error())
			return
		}
	}
	if len(_costexplorerDefaultValue) > 0 {
		input.DefaultValue = aws.String(_costexplorerDefaultValue)
	}
	if len(_costexplorerEffectiveStart) > 0 {
		input.EffectiveStart = aws.String(_costexplorerEffectiveStart)
	}
	if len(_costexplorerResourceTags) > 0 {
		if err := assignInputField(input, "ResourceTags", _costexplorerResourceTags); err != nil {
			log.Errorf("invalid --resource-tags: %s", err.Error())
			return
		}
	}
	if len(_costexplorerSplitChargeRules) > 0 {
		if err := assignInputField(input, "SplitChargeRules", _costexplorerSplitChargeRules); err != nil {
			log.Errorf("invalid --split-charge-rules: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCostCategoryDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a cost anomaly monitor.
func costexplorer_DeleteAnomalyMonitor(cfg aws.Config, client *costexplorer.Client) {
	input := &costexplorer.DeleteAnomalyMonitorInput{
		// MonitorArn: *string, // Required
	}

	if len(_costexplorerMonitorArn) > 0 {
		input.MonitorArn = aws.String(_costexplorerMonitorArn)
	}

	if resp, err := client.DeleteAnomalyMonitor(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a cost anomaly subscription.
func costexplorer_DeleteAnomalySubscription(cfg aws.Config, client *costexplorer.Client) {
	input := &costexplorer.DeleteAnomalySubscriptionInput{
		// SubscriptionArn: *string, // Required
	}

	if len(_costexplorerSubscriptionArn) > 0 {
		input.SubscriptionArn = aws.String(_costexplorerSubscriptionArn)
	}

	if resp, err := client.DeleteAnomalySubscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a cost category. Expenses from this month going forward will no longer
// be categorized with this cost category.
func costexplorer_DeleteCostCategoryDefinition(cfg aws.Config, client *costexplorer.Client) {
	input := &costexplorer.DeleteCostCategoryDefinitionInput{
		// CostCategoryArn: *string, // Required
	}

	if len(_costexplorerCostCategoryArn) > 0 {
		input.CostCategoryArn = aws.String(_costexplorerCostCategoryArn)
	}

	if resp, err := client.DeleteCostCategoryDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the name, Amazon Resource Name (ARN), rules, definition, and effective
// dates of a cost category that's defined in the account.
//
// You have the option to use EffectiveOn to return a cost category that's active
// on a specific date. If there's no EffectiveOn specified, you see a Cost
// Category that's effective on the current date. If cost category is still
// effective, EffectiveEnd is omitted in the response.
func costexplorer_DescribeCostCategoryDefinition(cfg aws.Config, client *costexplorer.Client) {
	input := &costexplorer.DescribeCostCategoryDefinitionInput{
		// CostCategoryArn: *string, // Required
	}

	if len(_costexplorerCostCategoryArn) > 0 {
		input.CostCategoryArn = aws.String(_costexplorerCostCategoryArn)
	}
	if len(_costexplorerEffectiveOn) > 0 {
		input.EffectiveOn = aws.String(_costexplorerEffectiveOn)
	}

	if resp, err := client.DescribeCostCategoryDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves all of the cost anomalies detected on your account during the time
// period that's specified by the DateInterval object. Anomalies are available for
// up to 90 days.
func costexplorer_GetAnomalies(cfg aws.Config, client *costexplorer.Client) {
	input := &costexplorer.GetAnomaliesInput{
		// DateInterval: *types.AnomalyDateInterval, // Required
	}

	if len(_costexplorerDateInterval) > 0 {
		if err := assignInputField(input, "DateInterval", _costexplorerDateInterval); err != nil {
			log.Errorf("invalid --date-interval: %s", err.Error())
			return
		}
	}
	if len(_costexplorerFeedback) > 0 {
		if err := assignInputField(input, "Feedback", _costexplorerFeedback); err != nil {
			log.Errorf("invalid --feedback: %s", err.Error())
			return
		}
	}
	if len(_costexplorerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _costexplorerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_costexplorerMonitorArn) > 0 {
		input.MonitorArn = aws.String(_costexplorerMonitorArn)
	}
	if len(_costexplorerNextPageToken) > 0 {
		input.NextPageToken = aws.String(_costexplorerNextPageToken)
	}
	if len(_costexplorerTotalImpact) > 0 {
		if err := assignInputField(input, "TotalImpact", _costexplorerTotalImpact); err != nil {
			log.Errorf("invalid --total-impact: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.GetAnomalies(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*costexplorer.GetAnomaliesOutput
	p := costexplorer.NewGetAnomaliesPaginator(client, input)
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

// Retrieves the cost anomaly monitor definitions for your account. You can filter
// using a list of cost anomaly monitor Amazon Resource Names (ARNs).
func costexplorer_GetAnomalyMonitors(cfg aws.Config, client *costexplorer.Client) {
	input := &costexplorer.GetAnomalyMonitorsInput{}

	if len(_costexplorerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _costexplorerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_costexplorerMonitorArnList) > 0 {
		input.MonitorArnList = append([]string(nil), _costexplorerMonitorArnList...)
	}
	if len(_costexplorerNextPageToken) > 0 {
		input.NextPageToken = aws.String(_costexplorerNextPageToken)
	}

	if disablePaginator() {
		if resp, err := client.GetAnomalyMonitors(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*costexplorer.GetAnomalyMonitorsOutput
	p := costexplorer.NewGetAnomalyMonitorsPaginator(client, input)
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

// Retrieves the cost anomaly subscription objects for your account. You can
// filter using a list of cost anomaly monitor Amazon Resource Names (ARNs).
func costexplorer_GetAnomalySubscriptions(cfg aws.Config, client *costexplorer.Client) {
	input := &costexplorer.GetAnomalySubscriptionsInput{}

	if len(_costexplorerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _costexplorerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_costexplorerMonitorArn) > 0 {
		input.MonitorArn = aws.String(_costexplorerMonitorArn)
	}
	if len(_costexplorerNextPageToken) > 0 {
		input.NextPageToken = aws.String(_costexplorerNextPageToken)
	}
	if len(_costexplorerSubscriptionArnList) > 0 {
		input.SubscriptionArnList = append([]string(nil), _costexplorerSubscriptionArnList...)
	}

	if disablePaginator() {
		if resp, err := client.GetAnomalySubscriptions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*costexplorer.GetAnomalySubscriptionsOutput
	p := costexplorer.NewGetAnomalySubscriptionsPaginator(client, input)
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

// Retrieves estimated usage records for hourly granularity or resource-level data
// at daily granularity.
func costexplorer_GetApproximateUsageRecords(cfg aws.Config, client *costexplorer.Client) {
	input := &costexplorer.GetApproximateUsageRecordsInput{
		// ApproximationDimension: types.ApproximationDimension, // Required
		// Granularity: types.Granularity, // Required
	}

	if len(_costexplorerApproximationDimension) > 0 {
		if err := assignInputField(input, "ApproximationDimension", _costexplorerApproximationDimension); err != nil {
			log.Errorf("invalid --approximation-dimension: %s", err.Error())
			return
		}
	}
	if len(_costexplorerGranularity) > 0 {
		if err := assignInputField(input, "Granularity", _costexplorerGranularity); err != nil {
			log.Errorf("invalid --granularity: %s", err.Error())
			return
		}
	}
	if len(_costexplorerServices) > 0 {
		input.Services = append([]string(nil), _costexplorerServices...)
	}

	if resp, err := client.GetApproximateUsageRecords(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a commitment purchase analysis result based on the AnalysisId .
func costexplorer_GetCommitmentPurchaseAnalysis(cfg aws.Config, client *costexplorer.Client) {
	input := &costexplorer.GetCommitmentPurchaseAnalysisInput{
		// AnalysisId: *string, // Required
	}

	if len(_costexplorerAnalysisId) > 0 {
		input.AnalysisId = aws.String(_costexplorerAnalysisId)
	}

	if resp, err := client.GetCommitmentPurchaseAnalysis(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves cost and usage metrics for your account. You can specify which cost
// and usage-related metric that you want the request to return. For example, you
// can specify BlendedCosts or UsageQuantity . You can also filter and group your
// data by various dimensions, such as SERVICE or AZ , in a specific time range.
// For a complete list of valid dimensions, see the [GetDimensionValues]operation. Management account
// in an organization in Organizations have access to all member accounts.
//
// For information about filter limitations, see [Quotas and restrictions] in the Billing and Cost
// Management User Guide.
//
// [GetDimensionValues]: https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_GetDimensionValues.html
// [Quotas and restrictions]: https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/billing-limits.html
func costexplorer_GetCostAndUsage(cfg aws.Config, client *costexplorer.Client) {
	input := &costexplorer.GetCostAndUsageInput{
		// Granularity: types.Granularity, // Required
		// Metrics: []string, // Required
		// TimePeriod: *types.DateInterval, // Required
	}

	if len(_costexplorerGranularity) > 0 {
		if err := assignInputField(input, "Granularity", _costexplorerGranularity); err != nil {
			log.Errorf("invalid --granularity: %s", err.Error())
			return
		}
	}
	if len(_costexplorerMetrics) > 0 {
		input.Metrics = append([]string(nil), _costexplorerMetrics...)
	}
	if len(_costexplorerTimePeriod) > 0 {
		if err := assignInputField(input, "TimePeriod", _costexplorerTimePeriod); err != nil {
			log.Errorf("invalid --time-period: %s", err.Error())
			return
		}
	}
	if len(_costexplorerBillingViewArn) > 0 {
		input.BillingViewArn = aws.String(_costexplorerBillingViewArn)
	}
	if len(_costexplorerFilter) > 0 {
		if err := assignInputField(input, "Filter", _costexplorerFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_costexplorerGroupBy) > 0 {
		if err := assignInputField(input, "GroupBy", _costexplorerGroupBy); err != nil {
			log.Errorf("invalid --group-by: %s", err.Error())
			return
		}
	}
	if len(_costexplorerNextPageToken) > 0 {
		input.NextPageToken = aws.String(_costexplorerNextPageToken)
	}

	if resp, err := client.GetCostAndUsage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves cost and usage comparisons for your account between two periods
// within the last 13 months. If you have enabled multi-year data at monthly
// granularity, you can go back up to 38 months.
func costexplorer_GetCostAndUsageComparisons(cfg aws.Config, client *costexplorer.Client) {
	input := &costexplorer.GetCostAndUsageComparisonsInput{
		// BaselineTimePeriod: *types.DateInterval, // Required
		// ComparisonTimePeriod: *types.DateInterval, // Required
		// MetricForComparison: *string, // Required
	}

	if len(_costexplorerBaselineTimePeriod) > 0 {
		if err := assignInputField(input, "BaselineTimePeriod", _costexplorerBaselineTimePeriod); err != nil {
			log.Errorf("invalid --baseline-time-period: %s", err.Error())
			return
		}
	}
	if len(_costexplorerComparisonTimePeriod) > 0 {
		if err := assignInputField(input, "ComparisonTimePeriod", _costexplorerComparisonTimePeriod); err != nil {
			log.Errorf("invalid --comparison-time-period: %s", err.Error())
			return
		}
	}
	if len(_costexplorerMetricForComparison) > 0 {
		input.MetricForComparison = aws.String(_costexplorerMetricForComparison)
	}
	if len(_costexplorerBillingViewArn) > 0 {
		input.BillingViewArn = aws.String(_costexplorerBillingViewArn)
	}
	if len(_costexplorerFilter) > 0 {
		if err := assignInputField(input, "Filter", _costexplorerFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_costexplorerGroupBy) > 0 {
		if err := assignInputField(input, "GroupBy", _costexplorerGroupBy); err != nil {
			log.Errorf("invalid --group-by: %s", err.Error())
			return
		}
	}
	if len(_costexplorerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _costexplorerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_costexplorerNextPageToken) > 0 {
		input.NextPageToken = aws.String(_costexplorerNextPageToken)
	}

	if disablePaginator() {
		if resp, err := client.GetCostAndUsageComparisons(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*costexplorer.GetCostAndUsageComparisonsOutput
	p := costexplorer.NewGetCostAndUsageComparisonsPaginator(client, input)
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

// Retrieves cost and usage metrics with resources for your account. You can
// specify which cost and usage-related metric, such as BlendedCosts or
// UsageQuantity , that you want the request to return. You can also filter and
// group your data by various dimensions, such as SERVICE or AZ , in a specific
// time range. For a complete list of valid dimensions, see the [GetDimensionValues]operation.
// Management account in an organization in Organizations have access to all member
// accounts.
//
// Hourly granularity is only available for EC2-Instances (Elastic Compute Cloud)
// resource-level data. All other resource-level data is available at daily
// granularity.
//
// This is an opt-in only feature. You can enable this feature from the Cost
// Explorer Settings page. For information about how to access the Settings page,
// see [Controlling Access for Cost Explorer]in the Billing and Cost Management User Guide.
//
// [GetDimensionValues]: https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_GetDimensionValues.html
// [Controlling Access for Cost Explorer]: https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/ce-access.html
func costexplorer_GetCostAndUsageWithResources(cfg aws.Config, client *costexplorer.Client) {
	input := &costexplorer.GetCostAndUsageWithResourcesInput{
		// Filter: *types.Expression, // Required
		// Granularity: types.Granularity, // Required
		// TimePeriod: *types.DateInterval, // Required
	}

	if len(_costexplorerFilter) > 0 {
		if err := assignInputField(input, "Filter", _costexplorerFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_costexplorerGranularity) > 0 {
		if err := assignInputField(input, "Granularity", _costexplorerGranularity); err != nil {
			log.Errorf("invalid --granularity: %s", err.Error())
			return
		}
	}
	if len(_costexplorerTimePeriod) > 0 {
		if err := assignInputField(input, "TimePeriod", _costexplorerTimePeriod); err != nil {
			log.Errorf("invalid --time-period: %s", err.Error())
			return
		}
	}
	if len(_costexplorerBillingViewArn) > 0 {
		input.BillingViewArn = aws.String(_costexplorerBillingViewArn)
	}
	if len(_costexplorerGroupBy) > 0 {
		if err := assignInputField(input, "GroupBy", _costexplorerGroupBy); err != nil {
			log.Errorf("invalid --group-by: %s", err.Error())
			return
		}
	}
	if len(_costexplorerMetrics) > 0 {
		input.Metrics = append([]string(nil), _costexplorerMetrics...)
	}
	if len(_costexplorerNextPageToken) > 0 {
		input.NextPageToken = aws.String(_costexplorerNextPageToken)
	}

	if resp, err := client.GetCostAndUsageWithResources(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves an array of cost category names and values incurred cost.
// If some cost category names and values are not associated with any cost, they
// will not be returned by this API.
func costexplorer_GetCostCategories(cfg aws.Config, client *costexplorer.Client) {
	input := &costexplorer.GetCostCategoriesInput{
		// TimePeriod: *types.DateInterval, // Required
	}

	if len(_costexplorerTimePeriod) > 0 {
		if err := assignInputField(input, "TimePeriod", _costexplorerTimePeriod); err != nil {
			log.Errorf("invalid --time-period: %s", err.Error())
			return
		}
	}
	if len(_costexplorerBillingViewArn) > 0 {
		input.BillingViewArn = aws.String(_costexplorerBillingViewArn)
	}
	if len(_costexplorerCostCategoryName) > 0 {
		input.CostCategoryName = aws.String(_costexplorerCostCategoryName)
	}
	if len(_costexplorerFilter) > 0 {
		if err := assignInputField(input, "Filter", _costexplorerFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_costexplorerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _costexplorerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_costexplorerNextPageToken) > 0 {
		input.NextPageToken = aws.String(_costexplorerNextPageToken)
	}
	if len(_costexplorerSearchString) > 0 {
		input.SearchString = aws.String(_costexplorerSearchString)
	}
	if len(_costexplorerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _costexplorerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetCostCategories(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves key factors driving cost changes between two time periods within the
// last 13 months, such as usage changes, discount changes, and commitment-based
// savings. If you have enabled multi-year data at monthly granularity, you can go
// back up to 38 months.
func costexplorer_GetCostComparisonDrivers(cfg aws.Config, client *costexplorer.Client) {
	input := &costexplorer.GetCostComparisonDriversInput{
		// BaselineTimePeriod: *types.DateInterval, // Required
		// ComparisonTimePeriod: *types.DateInterval, // Required
		// MetricForComparison: *string, // Required
	}

	if len(_costexplorerBaselineTimePeriod) > 0 {
		if err := assignInputField(input, "BaselineTimePeriod", _costexplorerBaselineTimePeriod); err != nil {
			log.Errorf("invalid --baseline-time-period: %s", err.Error())
			return
		}
	}
	if len(_costexplorerComparisonTimePeriod) > 0 {
		if err := assignInputField(input, "ComparisonTimePeriod", _costexplorerComparisonTimePeriod); err != nil {
			log.Errorf("invalid --comparison-time-period: %s", err.Error())
			return
		}
	}
	if len(_costexplorerMetricForComparison) > 0 {
		input.MetricForComparison = aws.String(_costexplorerMetricForComparison)
	}
	if len(_costexplorerBillingViewArn) > 0 {
		input.BillingViewArn = aws.String(_costexplorerBillingViewArn)
	}
	if len(_costexplorerFilter) > 0 {
		if err := assignInputField(input, "Filter", _costexplorerFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_costexplorerGroupBy) > 0 {
		if err := assignInputField(input, "GroupBy", _costexplorerGroupBy); err != nil {
			log.Errorf("invalid --group-by: %s", err.Error())
			return
		}
	}
	if len(_costexplorerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _costexplorerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_costexplorerNextPageToken) > 0 {
		input.NextPageToken = aws.String(_costexplorerNextPageToken)
	}

	if disablePaginator() {
		if resp, err := client.GetCostComparisonDrivers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*costexplorer.GetCostComparisonDriversOutput
	p := costexplorer.NewGetCostComparisonDriversPaginator(client, input)
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

// Retrieves a forecast for how much Amazon Web Services predicts that you will
// spend over the forecast time period that you select, based on your past costs.
func costexplorer_GetCostForecast(cfg aws.Config, client *costexplorer.Client) {
	input := &costexplorer.GetCostForecastInput{
		// Granularity: types.Granularity, // Required
		// Metric: types.Metric, // Required
		// TimePeriod: *types.DateInterval, // Required
	}

	if len(_costexplorerGranularity) > 0 {
		if err := assignInputField(input, "Granularity", _costexplorerGranularity); err != nil {
			log.Errorf("invalid --granularity: %s", err.Error())
			return
		}
	}
	if len(_costexplorerMetric) > 0 {
		if err := assignInputField(input, "Metric", _costexplorerMetric); err != nil {
			log.Errorf("invalid --metric: %s", err.Error())
			return
		}
	}
	if len(_costexplorerTimePeriod) > 0 {
		if err := assignInputField(input, "TimePeriod", _costexplorerTimePeriod); err != nil {
			log.Errorf("invalid --time-period: %s", err.Error())
			return
		}
	}
	if len(_costexplorerBillingViewArn) > 0 {
		input.BillingViewArn = aws.String(_costexplorerBillingViewArn)
	}
	if len(_costexplorerFilter) > 0 {
		if err := assignInputField(input, "Filter", _costexplorerFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_costexplorerPredictionIntervalLevel) > 0 {
		if err := assignInputField(input, "PredictionIntervalLevel", _costexplorerPredictionIntervalLevel); err != nil {
			log.Errorf("invalid --prediction-interval-level: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetCostForecast(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves all available filter values for a specified filter over a period of
// time. You can search the dimension values for an arbitrary string.
func costexplorer_GetDimensionValues(cfg aws.Config, client *costexplorer.Client) {
	input := &costexplorer.GetDimensionValuesInput{
		// Dimension: types.Dimension, // Required
		// TimePeriod: *types.DateInterval, // Required
	}

	if len(_costexplorerDimension) > 0 {
		if err := assignInputField(input, "Dimension", _costexplorerDimension); err != nil {
			log.Errorf("invalid --dimension: %s", err.Error())
			return
		}
	}
	if len(_costexplorerTimePeriod) > 0 {
		if err := assignInputField(input, "TimePeriod", _costexplorerTimePeriod); err != nil {
			log.Errorf("invalid --time-period: %s", err.Error())
			return
		}
	}
	if len(_costexplorerBillingViewArn) > 0 {
		input.BillingViewArn = aws.String(_costexplorerBillingViewArn)
	}
	if len(_costexplorerContext) > 0 {
		if err := assignInputField(input, "Context", _costexplorerContext); err != nil {
			log.Errorf("invalid --context: %s", err.Error())
			return
		}
	}
	if len(_costexplorerFilter) > 0 {
		if err := assignInputField(input, "Filter", _costexplorerFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_costexplorerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _costexplorerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_costexplorerNextPageToken) > 0 {
		input.NextPageToken = aws.String(_costexplorerNextPageToken)
	}
	if len(_costexplorerSearchString) > 0 {
		input.SearchString = aws.String(_costexplorerSearchString)
	}
	if len(_costexplorerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _costexplorerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetDimensionValues(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the reservation coverage for your account, which you can use to see
// how much of your Amazon Elastic Compute Cloud, Amazon ElastiCache, Amazon
// Relational Database Service, or Amazon Redshift usage is covered by a
// reservation. An organization's management account can see the coverage of the
// associated member accounts. This supports dimensions, cost categories, and
// nested expressions. For any time period, you can filter data about reservation
// usage by the following dimensions:
//
// - AZ
//
// - CACHE_ENGINE
//
// - DATABASE_ENGINE
//
// - DEPLOYMENT_OPTION
//
// - INSTANCE_TYPE
//
// - LINKED_ACCOUNT
//
// - OPERATING_SYSTEM
//
// - PLATFORM
//
// - REGION
//
// - SERVICE
//
// - TAG
//
// - TENANCY
//
// To determine valid values for a dimension, use the GetDimensionValues
// operation.
func costexplorer_GetReservationCoverage(cfg aws.Config, client *costexplorer.Client) {
	input := &costexplorer.GetReservationCoverageInput{
		// TimePeriod: *types.DateInterval, // Required
	}

	if len(_costexplorerTimePeriod) > 0 {
		if err := assignInputField(input, "TimePeriod", _costexplorerTimePeriod); err != nil {
			log.Errorf("invalid --time-period: %s", err.Error())
			return
		}
	}
	if len(_costexplorerFilter) > 0 {
		if err := assignInputField(input, "Filter", _costexplorerFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_costexplorerGranularity) > 0 {
		if err := assignInputField(input, "Granularity", _costexplorerGranularity); err != nil {
			log.Errorf("invalid --granularity: %s", err.Error())
			return
		}
	}
	if len(_costexplorerGroupBy) > 0 {
		if err := assignInputField(input, "GroupBy", _costexplorerGroupBy); err != nil {
			log.Errorf("invalid --group-by: %s", err.Error())
			return
		}
	}
	if len(_costexplorerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _costexplorerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_costexplorerMetrics) > 0 {
		input.Metrics = append([]string(nil), _costexplorerMetrics...)
	}
	if len(_costexplorerNextPageToken) > 0 {
		input.NextPageToken = aws.String(_costexplorerNextPageToken)
	}
	if len(_costexplorerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _costexplorerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetReservationCoverage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets recommendations for reservation purchases. These recommendations might
// help you to reduce your costs. Reservations provide a discounted hourly rate (up
// to 75%) compared to On-Demand pricing.
//
// Amazon Web Services generates your recommendations by identifying your
// On-Demand usage during a specific time period and collecting your usage into
// categories that are eligible for a reservation. After Amazon Web Services has
// these categories, it simulates every combination of reservations in each
// category of usage to identify the best number of each type of Reserved Instance
// (RI) to purchase to maximize your estimated savings.
//
// For example, Amazon Web Services automatically aggregates your Amazon EC2
// Linux, shared tenancy, and c4 family usage in the US West (Oregon) Region and
// recommends that you buy size-flexible regional reservations to apply to the c4
// family usage. Amazon Web Services recommends the smallest size instance in an
// instance family. This makes it easier to purchase a size-flexible Reserved
// Instance (RI). Amazon Web Services also shows the equal number of normalized
// units. This way, you can purchase any instance size that you want. For this
// example, your RI recommendation is for c4.large because that is the smallest
// size instance in the c4 instance family.
func costexplorer_GetReservationPurchaseRecommendation(cfg aws.Config, client *costexplorer.Client) {
	input := &costexplorer.GetReservationPurchaseRecommendationInput{
		// Service: *string, // Required
	}

	if len(_costexplorerService) > 0 {
		input.Service = aws.String(_costexplorerService)
	}
	if len(_costexplorerAccountId) > 0 {
		input.AccountId = aws.String(_costexplorerAccountId)
	}
	if len(_costexplorerAccountScope) > 0 {
		if err := assignInputField(input, "AccountScope", _costexplorerAccountScope); err != nil {
			log.Errorf("invalid --account-scope: %s", err.Error())
			return
		}
	}
	if len(_costexplorerFilter) > 0 {
		if err := assignInputField(input, "Filter", _costexplorerFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_costexplorerLookbackPeriodInDays) > 0 {
		if err := assignInputField(input, "LookbackPeriodInDays", _costexplorerLookbackPeriodInDays); err != nil {
			log.Errorf("invalid --lookback-period-in-days: %s", err.Error())
			return
		}
	}
	if len(_costexplorerNextPageToken) > 0 {
		input.NextPageToken = aws.String(_costexplorerNextPageToken)
	}
	if len(_costexplorerPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _costexplorerPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}
	if len(_costexplorerPaymentOption) > 0 {
		if err := assignInputField(input, "PaymentOption", _costexplorerPaymentOption); err != nil {
			log.Errorf("invalid --payment-option: %s", err.Error())
			return
		}
	}
	if len(_costexplorerServiceSpecification) > 0 {
		if err := assignInputField(input, "ServiceSpecification", _costexplorerServiceSpecification); err != nil {
			log.Errorf("invalid --service-specification: %s", err.Error())
			return
		}
	}
	if len(_costexplorerTermInYears) > 0 {
		if err := assignInputField(input, "TermInYears", _costexplorerTermInYears); err != nil {
			log.Errorf("invalid --term-in-years: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.GetReservationPurchaseRecommendation(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*costexplorer.GetReservationPurchaseRecommendationOutput
	p := costexplorer.NewGetReservationPurchaseRecommendationPaginator(client, input)
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

// Retrieves the reservation utilization for your account. Management account in
// an organization have access to member accounts. You can filter data by
// dimensions in a time period. You can use GetDimensionValues to determine the
// possible dimension values. Currently, you can group only by SUBSCRIPTION_ID .
func costexplorer_GetReservationUtilization(cfg aws.Config, client *costexplorer.Client) {
	input := &costexplorer.GetReservationUtilizationInput{
		// TimePeriod: *types.DateInterval, // Required
	}

	if len(_costexplorerTimePeriod) > 0 {
		if err := assignInputField(input, "TimePeriod", _costexplorerTimePeriod); err != nil {
			log.Errorf("invalid --time-period: %s", err.Error())
			return
		}
	}
	if len(_costexplorerFilter) > 0 {
		if err := assignInputField(input, "Filter", _costexplorerFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_costexplorerGranularity) > 0 {
		if err := assignInputField(input, "Granularity", _costexplorerGranularity); err != nil {
			log.Errorf("invalid --granularity: %s", err.Error())
			return
		}
	}
	if len(_costexplorerGroupBy) > 0 {
		if err := assignInputField(input, "GroupBy", _costexplorerGroupBy); err != nil {
			log.Errorf("invalid --group-by: %s", err.Error())
			return
		}
	}
	if len(_costexplorerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _costexplorerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_costexplorerNextPageToken) > 0 {
		input.NextPageToken = aws.String(_costexplorerNextPageToken)
	}
	if len(_costexplorerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _costexplorerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetReservationUtilization(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates recommendations that help you save cost by identifying idle and
// underutilized Amazon EC2 instances.
//
// Recommendations are generated to either downsize or terminate instances, along
// with providing savings detail and metrics. For more information about
// calculation and function, see [Optimizing Your Cost with Rightsizing Recommendations]in the Billing and Cost Management User Guide.
//
// [Optimizing Your Cost with Rightsizing Recommendations]: https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/ce-rightsizing.html
func costexplorer_GetRightsizingRecommendation(cfg aws.Config, client *costexplorer.Client) {
	input := &costexplorer.GetRightsizingRecommendationInput{
		// Service: *string, // Required
	}

	if len(_costexplorerService) > 0 {
		input.Service = aws.String(_costexplorerService)
	}
	if len(_costexplorerConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _costexplorerConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}
	if len(_costexplorerFilter) > 0 {
		if err := assignInputField(input, "Filter", _costexplorerFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_costexplorerNextPageToken) > 0 {
		input.NextPageToken = aws.String(_costexplorerNextPageToken)
	}
	if len(_costexplorerPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _costexplorerPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.GetRightsizingRecommendation(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*costexplorer.GetRightsizingRecommendationOutput
	p := costexplorer.NewGetRightsizingRecommendationPaginator(client, input)
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

// Retrieves the details for a Savings Plan recommendation. These details include
// the hourly data-points that construct the cost, coverage, and utilization
// charts.
func costexplorer_GetSavingsPlanPurchaseRecommendationDetails(cfg aws.Config, client *costexplorer.Client) {
	input := &costexplorer.GetSavingsPlanPurchaseRecommendationDetailsInput{
		// RecommendationDetailId: *string, // Required
	}

	if len(_costexplorerRecommendationDetailId) > 0 {
		input.RecommendationDetailId = aws.String(_costexplorerRecommendationDetailId)
	}

	if resp, err := client.GetSavingsPlanPurchaseRecommendationDetails(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the Savings Plans covered for your account. This enables you to see
// how much of your cost is covered by a Savings Plan. An organization’s management
// account can see the coverage of the associated member accounts. This supports
// dimensions, cost categories, and nested expressions. For any time period, you
// can filter data for Savings Plans usage with the following dimensions:
//
// - LINKED_ACCOUNT
//
// - REGION
//
// - SERVICE
//
// - INSTANCE_FAMILY
//
// To determine valid values for a dimension, use the GetDimensionValues operation.
func costexplorer_GetSavingsPlansCoverage(cfg aws.Config, client *costexplorer.Client) {
	input := &costexplorer.GetSavingsPlansCoverageInput{
		// TimePeriod: *types.DateInterval, // Required
	}

	if len(_costexplorerTimePeriod) > 0 {
		if err := assignInputField(input, "TimePeriod", _costexplorerTimePeriod); err != nil {
			log.Errorf("invalid --time-period: %s", err.Error())
			return
		}
	}
	if len(_costexplorerFilter) > 0 {
		if err := assignInputField(input, "Filter", _costexplorerFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_costexplorerGranularity) > 0 {
		if err := assignInputField(input, "Granularity", _costexplorerGranularity); err != nil {
			log.Errorf("invalid --granularity: %s", err.Error())
			return
		}
	}
	if len(_costexplorerGroupBy) > 0 {
		if err := assignInputField(input, "GroupBy", _costexplorerGroupBy); err != nil {
			log.Errorf("invalid --group-by: %s", err.Error())
			return
		}
	}
	if len(_costexplorerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _costexplorerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_costexplorerMetrics) > 0 {
		input.Metrics = append([]string(nil), _costexplorerMetrics...)
	}
	if len(_costexplorerNextToken) > 0 {
		input.NextToken = aws.String(_costexplorerNextToken)
	}
	if len(_costexplorerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _costexplorerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.GetSavingsPlansCoverage(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*costexplorer.GetSavingsPlansCoverageOutput
	p := costexplorer.NewGetSavingsPlansCoveragePaginator(client, input)
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

// Retrieves the Savings Plans recommendations for your account. First use
// StartSavingsPlansPurchaseRecommendationGeneration to generate a new set of
// recommendations, and then use GetSavingsPlansPurchaseRecommendation to retrieve
// them.
func costexplorer_GetSavingsPlansPurchaseRecommendation(cfg aws.Config, client *costexplorer.Client) {
	input := &costexplorer.GetSavingsPlansPurchaseRecommendationInput{
		// LookbackPeriodInDays: types.LookbackPeriodInDays, // Required
		// PaymentOption: types.PaymentOption, // Required
		// SavingsPlansType: types.SupportedSavingsPlansType, // Required
		// TermInYears: types.TermInYears, // Required
	}

	if len(_costexplorerLookbackPeriodInDays) > 0 {
		if err := assignInputField(input, "LookbackPeriodInDays", _costexplorerLookbackPeriodInDays); err != nil {
			log.Errorf("invalid --lookback-period-in-days: %s", err.Error())
			return
		}
	}
	if len(_costexplorerPaymentOption) > 0 {
		if err := assignInputField(input, "PaymentOption", _costexplorerPaymentOption); err != nil {
			log.Errorf("invalid --payment-option: %s", err.Error())
			return
		}
	}
	if len(_costexplorerSavingsPlansType) > 0 {
		if err := assignInputField(input, "SavingsPlansType", _costexplorerSavingsPlansType); err != nil {
			log.Errorf("invalid --savings-plans-type: %s", err.Error())
			return
		}
	}
	if len(_costexplorerTermInYears) > 0 {
		if err := assignInputField(input, "TermInYears", _costexplorerTermInYears); err != nil {
			log.Errorf("invalid --term-in-years: %s", err.Error())
			return
		}
	}
	if len(_costexplorerAccountScope) > 0 {
		if err := assignInputField(input, "AccountScope", _costexplorerAccountScope); err != nil {
			log.Errorf("invalid --account-scope: %s", err.Error())
			return
		}
	}
	if len(_costexplorerFilter) > 0 {
		if err := assignInputField(input, "Filter", _costexplorerFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_costexplorerNextPageToken) > 0 {
		input.NextPageToken = aws.String(_costexplorerNextPageToken)
	}
	if len(_costexplorerPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _costexplorerPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetSavingsPlansPurchaseRecommendation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the Savings Plans utilization for your account across date ranges
// with daily or monthly granularity. Management account in an organization have
// access to member accounts. You can use GetDimensionValues in SAVINGS_PLANS to
// determine the possible dimension values.
//
// You can't group by any dimension values for GetSavingsPlansUtilization .
func costexplorer_GetSavingsPlansUtilization(cfg aws.Config, client *costexplorer.Client) {
	input := &costexplorer.GetSavingsPlansUtilizationInput{
		// TimePeriod: *types.DateInterval, // Required
	}

	if len(_costexplorerTimePeriod) > 0 {
		if err := assignInputField(input, "TimePeriod", _costexplorerTimePeriod); err != nil {
			log.Errorf("invalid --time-period: %s", err.Error())
			return
		}
	}
	if len(_costexplorerFilter) > 0 {
		if err := assignInputField(input, "Filter", _costexplorerFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_costexplorerGranularity) > 0 {
		if err := assignInputField(input, "Granularity", _costexplorerGranularity); err != nil {
			log.Errorf("invalid --granularity: %s", err.Error())
			return
		}
	}
	if len(_costexplorerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _costexplorerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetSavingsPlansUtilization(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves attribute data along with aggregate utilization and savings data for
// a given time period. This doesn't support granular or grouped data
// (daily/monthly) in response. You can't retrieve data by dates in a single
// response similar to GetSavingsPlanUtilization , but you have the option to make
// multiple calls to GetSavingsPlanUtilizationDetails by providing individual
// dates. You can use GetDimensionValues in SAVINGS_PLANS to determine the
// possible dimension values.
//
// GetSavingsPlanUtilizationDetails internally groups data by SavingsPlansArn .
func costexplorer_GetSavingsPlansUtilizationDetails(cfg aws.Config, client *costexplorer.Client) {
	input := &costexplorer.GetSavingsPlansUtilizationDetailsInput{
		// TimePeriod: *types.DateInterval, // Required
	}

	if len(_costexplorerTimePeriod) > 0 {
		if err := assignInputField(input, "TimePeriod", _costexplorerTimePeriod); err != nil {
			log.Errorf("invalid --time-period: %s", err.Error())
			return
		}
	}
	if len(_costexplorerDataType) > 0 {
		if err := assignInputField(input, "DataType", _costexplorerDataType); err != nil {
			log.Errorf("invalid --data-type: %s", err.Error())
			return
		}
	}
	if len(_costexplorerFilter) > 0 {
		if err := assignInputField(input, "Filter", _costexplorerFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_costexplorerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _costexplorerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_costexplorerNextToken) > 0 {
		input.NextToken = aws.String(_costexplorerNextToken)
	}
	if len(_costexplorerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _costexplorerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.GetSavingsPlansUtilizationDetails(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*costexplorer.GetSavingsPlansUtilizationDetailsOutput
	p := costexplorer.NewGetSavingsPlansUtilizationDetailsPaginator(client, input)
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

// Queries for available tag keys and tag values for a specified period. You can
// search the tag values for an arbitrary string.
func costexplorer_GetTags(cfg aws.Config, client *costexplorer.Client) {
	input := &costexplorer.GetTagsInput{
		// TimePeriod: *types.DateInterval, // Required
	}

	if len(_costexplorerTimePeriod) > 0 {
		if err := assignInputField(input, "TimePeriod", _costexplorerTimePeriod); err != nil {
			log.Errorf("invalid --time-period: %s", err.Error())
			return
		}
	}
	if len(_costexplorerBillingViewArn) > 0 {
		input.BillingViewArn = aws.String(_costexplorerBillingViewArn)
	}
	if len(_costexplorerFilter) > 0 {
		if err := assignInputField(input, "Filter", _costexplorerFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_costexplorerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _costexplorerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_costexplorerNextPageToken) > 0 {
		input.NextPageToken = aws.String(_costexplorerNextPageToken)
	}
	if len(_costexplorerSearchString) > 0 {
		input.SearchString = aws.String(_costexplorerSearchString)
	}
	if len(_costexplorerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _costexplorerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_costexplorerTagKey) > 0 {
		input.TagKey = aws.String(_costexplorerTagKey)
	}

	if resp, err := client.GetTags(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a forecast for how much Amazon Web Services predicts that you will
// use over the forecast time period that you select, based on your past usage.
func costexplorer_GetUsageForecast(cfg aws.Config, client *costexplorer.Client) {
	input := &costexplorer.GetUsageForecastInput{
		// Granularity: types.Granularity, // Required
		// Metric: types.Metric, // Required
		// TimePeriod: *types.DateInterval, // Required
	}

	if len(_costexplorerGranularity) > 0 {
		if err := assignInputField(input, "Granularity", _costexplorerGranularity); err != nil {
			log.Errorf("invalid --granularity: %s", err.Error())
			return
		}
	}
	if len(_costexplorerMetric) > 0 {
		if err := assignInputField(input, "Metric", _costexplorerMetric); err != nil {
			log.Errorf("invalid --metric: %s", err.Error())
			return
		}
	}
	if len(_costexplorerTimePeriod) > 0 {
		if err := assignInputField(input, "TimePeriod", _costexplorerTimePeriod); err != nil {
			log.Errorf("invalid --time-period: %s", err.Error())
			return
		}
	}
	if len(_costexplorerBillingViewArn) > 0 {
		input.BillingViewArn = aws.String(_costexplorerBillingViewArn)
	}
	if len(_costexplorerFilter) > 0 {
		if err := assignInputField(input, "Filter", _costexplorerFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_costexplorerPredictionIntervalLevel) > 0 {
		if err := assignInputField(input, "PredictionIntervalLevel", _costexplorerPredictionIntervalLevel); err != nil {
			log.Errorf("invalid --prediction-interval-level: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetUsageForecast(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the commitment purchase analyses for your account.
func costexplorer_ListCommitmentPurchaseAnalyses(cfg aws.Config, client *costexplorer.Client) {
	input := &costexplorer.ListCommitmentPurchaseAnalysesInput{}

	if len(_costexplorerAnalysisIds) > 0 {
		input.AnalysisIds = append([]string(nil), _costexplorerAnalysisIds...)
	}
	if len(_costexplorerAnalysisStatus) > 0 {
		if err := assignInputField(input, "AnalysisStatus", _costexplorerAnalysisStatus); err != nil {
			log.Errorf("invalid --analysis-status: %s", err.Error())
			return
		}
	}
	if len(_costexplorerNextPageToken) > 0 {
		input.NextPageToken = aws.String(_costexplorerNextPageToken)
	}
	if len(_costexplorerPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _costexplorerPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListCommitmentPurchaseAnalyses(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*costexplorer.ListCommitmentPurchaseAnalysesOutput
	p := costexplorer.NewListCommitmentPurchaseAnalysesPaginator(client, input)
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

// Retrieves a list of your historical cost allocation tag backfill requests.
func costexplorer_ListCostAllocationTagBackfillHistory(cfg aws.Config, client *costexplorer.Client) {
	input := &costexplorer.ListCostAllocationTagBackfillHistoryInput{}

	if len(_costexplorerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _costexplorerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_costexplorerNextToken) > 0 {
		input.NextToken = aws.String(_costexplorerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCostAllocationTagBackfillHistory(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*costexplorer.ListCostAllocationTagBackfillHistoryOutput
	p := costexplorer.NewListCostAllocationTagBackfillHistoryPaginator(client, input)
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

// Get a list of cost allocation tags. All inputs in the API are optional and
// serve as filters. By default, all cost allocation tags are returned.
func costexplorer_ListCostAllocationTags(cfg aws.Config, client *costexplorer.Client) {
	input := &costexplorer.ListCostAllocationTagsInput{}

	if len(_costexplorerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _costexplorerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_costexplorerNextToken) > 0 {
		input.NextToken = aws.String(_costexplorerNextToken)
	}
	if len(_costexplorerStatus) > 0 {
		if err := assignInputField(input, "Status", _costexplorerStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_costexplorerTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _costexplorerTagKeys...)
	}
	if len(_costexplorerType) > 0 {
		if err := assignInputField(input, "Type", _costexplorerType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListCostAllocationTags(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*costexplorer.ListCostAllocationTagsOutput
	p := costexplorer.NewListCostAllocationTagsPaginator(client, input)
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

// Returns the name, Amazon Resource Name (ARN), NumberOfRules and effective dates
// of all cost categories defined in the account. You have the option to use
// EffectiveOn and SupportedResourceTypes to return a list of cost categories that
// were active on a specific date. If there is no EffectiveOn specified, you’ll
// see cost categories that are effective on the current date. If cost category is
// still effective, EffectiveEnd is omitted in the response.
// ListCostCategoryDefinitions supports pagination. The request can have a
// MaxResults range up to 100.
func costexplorer_ListCostCategoryDefinitions(cfg aws.Config, client *costexplorer.Client) {
	input := &costexplorer.ListCostCategoryDefinitionsInput{}

	if len(_costexplorerEffectiveOn) > 0 {
		input.EffectiveOn = aws.String(_costexplorerEffectiveOn)
	}
	if len(_costexplorerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _costexplorerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_costexplorerNextToken) > 0 {
		input.NextToken = aws.String(_costexplorerNextToken)
	}
	if len(_costexplorerSupportedResourceTypes) > 0 {
		input.SupportedResourceTypes = append([]string(nil), _costexplorerSupportedResourceTypes...)
	}

	if disablePaginator() {
		if resp, err := client.ListCostCategoryDefinitions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*costexplorer.ListCostCategoryDefinitionsOutput
	p := costexplorer.NewListCostCategoryDefinitionsPaginator(client, input)
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

// Returns resource associations of all cost categories defined in the account.
// You have the option to use CostCategoryArn to get the association for a
// specific cost category. ListCostCategoryResourceAssociations supports
// pagination. The request can have a MaxResults range up to 100.
func costexplorer_ListCostCategoryResourceAssociations(cfg aws.Config, client *costexplorer.Client) {
	input := &costexplorer.ListCostCategoryResourceAssociationsInput{}

	if len(_costexplorerCostCategoryArn) > 0 {
		input.CostCategoryArn = aws.String(_costexplorerCostCategoryArn)
	}
	if len(_costexplorerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _costexplorerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_costexplorerNextToken) > 0 {
		input.NextToken = aws.String(_costexplorerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCostCategoryResourceAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*costexplorer.ListCostCategoryResourceAssociationsOutput
	p := costexplorer.NewListCostCategoryResourceAssociationsPaginator(client, input)
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

// Retrieves a list of your historical recommendation generations within the past
// 30 days.
func costexplorer_ListSavingsPlansPurchaseRecommendationGeneration(cfg aws.Config, client *costexplorer.Client) {
	input := &costexplorer.ListSavingsPlansPurchaseRecommendationGenerationInput{}

	if len(_costexplorerGenerationStatus) > 0 {
		if err := assignInputField(input, "GenerationStatus", _costexplorerGenerationStatus); err != nil {
			log.Errorf("invalid --generation-status: %s", err.Error())
			return
		}
	}
	if len(_costexplorerNextPageToken) > 0 {
		input.NextPageToken = aws.String(_costexplorerNextPageToken)
	}
	if len(_costexplorerPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _costexplorerPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}
	if len(_costexplorerRecommendationIds) > 0 {
		input.RecommendationIds = append([]string(nil), _costexplorerRecommendationIds...)
	}

	if disablePaginator() {
		if resp, err := client.ListSavingsPlansPurchaseRecommendationGeneration(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*costexplorer.ListSavingsPlansPurchaseRecommendationGenerationOutput
	p := costexplorer.NewListSavingsPlansPurchaseRecommendationGenerationPaginator(client, input)
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

// Returns a list of resource tags associated with the resource specified by the
// Amazon Resource Name (ARN).
func costexplorer_ListTagsForResource(cfg aws.Config, client *costexplorer.Client) {
	input := &costexplorer.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_costexplorerResourceArn) > 0 {
		input.ResourceArn = aws.String(_costexplorerResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the feedback property of a given cost anomaly.
func costexplorer_ProvideAnomalyFeedback(cfg aws.Config, client *costexplorer.Client) {
	input := &costexplorer.ProvideAnomalyFeedbackInput{
		// AnomalyId: *string, // Required
		// Feedback: types.AnomalyFeedbackType, // Required
	}

	if len(_costexplorerAnomalyId) > 0 {
		input.AnomalyId = aws.String(_costexplorerAnomalyId)
	}
	if len(_costexplorerFeedback) > 0 {
		if err := assignInputField(input, "Feedback", _costexplorerFeedback); err != nil {
			log.Errorf("invalid --feedback: %s", err.Error())
			return
		}
	}

	if resp, err := client.ProvideAnomalyFeedback(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Specifies the parameters of a planned commitment purchase and starts the
// generation of the analysis. This enables you to estimate the cost, coverage, and
// utilization impact of your planned commitment purchases.
func costexplorer_StartCommitmentPurchaseAnalysis(cfg aws.Config, client *costexplorer.Client) {
	input := &costexplorer.StartCommitmentPurchaseAnalysisInput{
		// CommitmentPurchaseAnalysisConfiguration: *types.CommitmentPurchaseAnalysisConfiguration, // Required
	}

	if len(_costexplorerCommitmentPurchaseAnalysisConfiguration) > 0 {
		if err := assignInputField(input, "CommitmentPurchaseAnalysisConfiguration", _costexplorerCommitmentPurchaseAnalysisConfiguration); err != nil {
			log.Errorf("invalid --commitment-purchase-analysis-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartCommitmentPurchaseAnalysis(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Request a cost allocation tag backfill. This will backfill the activation
// status (either active or inactive ) for all tag keys from para:BackfillFrom up
// to the time this request is made.
//
// You can request a backfill once every 24 hours.
func costexplorer_StartCostAllocationTagBackfill(cfg aws.Config, client *costexplorer.Client) {
	input := &costexplorer.StartCostAllocationTagBackfillInput{
		// BackfillFrom: *string, // Required
	}

	if len(_costexplorerBackfillFrom) > 0 {
		input.BackfillFrom = aws.String(_costexplorerBackfillFrom)
	}

	if resp, err := client.StartCostAllocationTagBackfill(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Requests a Savings Plans recommendation generation. This enables you to
// calculate a fresh set of Savings Plans recommendations that takes your latest
// usage data and current Savings Plans inventory into account. You can refresh
// Savings Plans recommendations up to three times daily for a consolidated billing
// family.
//
// StartSavingsPlansPurchaseRecommendationGeneration has no request syntax because
// no input parameters are needed to support this operation.
func costexplorer_StartSavingsPlansPurchaseRecommendationGeneration(cfg aws.Config, client *costexplorer.Client) {
	input := &costexplorer.StartSavingsPlansPurchaseRecommendationGenerationInput{}

	if resp, err := client.StartSavingsPlansPurchaseRecommendationGeneration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// An API operation for adding one or more tags (key-value pairs) to a resource.
// You can use the TagResource operation with a resource that already has tags. If
// you specify a new tag key for the resource, this tag is appended to the list of
// tags associated with the resource. If you specify a tag key that is already
// associated with the resource, the new tag value you specify replaces the
// previous value for that tag.
//
// Although the maximum number of array members is 200, user-tag maximum is 50.
// The remaining are reserved for Amazon Web Services use.
func costexplorer_TagResource(cfg aws.Config, client *costexplorer.Client) {
	input := &costexplorer.TagResourceInput{
		// ResourceArn: *string, // Required
		// ResourceTags: []types.ResourceTag, // Required
	}

	if len(_costexplorerResourceArn) > 0 {
		input.ResourceArn = aws.String(_costexplorerResourceArn)
	}
	if len(_costexplorerResourceTags) > 0 {
		if err := assignInputField(input, "ResourceTags", _costexplorerResourceTags); err != nil {
			log.Errorf("invalid --resource-tags: %s", err.Error())
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

// Removes one or more tags from a resource. Specify only tag keys in your
// request. Don't specify the value.
func costexplorer_UntagResource(cfg aws.Config, client *costexplorer.Client) {
	input := &costexplorer.UntagResourceInput{
		// ResourceArn: *string, // Required
		// ResourceTagKeys: []string, // Required
	}

	if len(_costexplorerResourceArn) > 0 {
		input.ResourceArn = aws.String(_costexplorerResourceArn)
	}
	if len(_costexplorerResourceTagKeys) > 0 {
		input.ResourceTagKeys = append([]string(nil), _costexplorerResourceTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing cost anomaly monitor. The changes made are applied going
// forward, and doesn't change anomalies detected in the past.
func costexplorer_UpdateAnomalyMonitor(cfg aws.Config, client *costexplorer.Client) {
	input := &costexplorer.UpdateAnomalyMonitorInput{
		// MonitorArn: *string, // Required
	}

	if len(_costexplorerMonitorArn) > 0 {
		input.MonitorArn = aws.String(_costexplorerMonitorArn)
	}
	if len(_costexplorerMonitorName) > 0 {
		input.MonitorName = aws.String(_costexplorerMonitorName)
	}

	if resp, err := client.UpdateAnomalyMonitor(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing cost anomaly subscription. Specify the fields that you want
// to update. Omitted fields are unchanged.
//
// The JSON below describes the generic construct for each type. See [Request Parameters] for possible
// values as they apply to AnomalySubscription .
//
// [Request Parameters]: https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_UpdateAnomalySubscription.html#API_UpdateAnomalySubscription_RequestParameters
func costexplorer_UpdateAnomalySubscription(cfg aws.Config, client *costexplorer.Client) {
	input := &costexplorer.UpdateAnomalySubscriptionInput{
		// SubscriptionArn: *string, // Required
	}

	if len(_costexplorerSubscriptionArn) > 0 {
		input.SubscriptionArn = aws.String(_costexplorerSubscriptionArn)
	}
	if len(_costexplorerFrequency) > 0 {
		if err := assignInputField(input, "Frequency", _costexplorerFrequency); err != nil {
			log.Errorf("invalid --frequency: %s", err.Error())
			return
		}
	}
	if len(_costexplorerMonitorArnList) > 0 {
		input.MonitorArnList = append([]string(nil), _costexplorerMonitorArnList...)
	}
	if len(_costexplorerSubscribers) > 0 {
		if err := assignInputField(input, "Subscribers", _costexplorerSubscribers); err != nil {
			log.Errorf("invalid --subscribers: %s", err.Error())
			return
		}
	}
	if len(_costexplorerSubscriptionName) > 0 {
		input.SubscriptionName = aws.String(_costexplorerSubscriptionName)
	}
	if len(_costexplorerThreshold) > 0 {
		if err := assignInputField(input, "Threshold", _costexplorerThreshold); err != nil {
			log.Errorf("invalid --threshold: %s", err.Error())
			return
		}
	}
	if len(_costexplorerThresholdExpression) > 0 {
		if err := assignInputField(input, "ThresholdExpression", _costexplorerThresholdExpression); err != nil {
			log.Errorf("invalid --threshold-expression: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateAnomalySubscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates status for cost allocation tags in bulk, with maximum batch size of 20.
// If the tag status that's updated is the same as the existing tag status, the
// request doesn't fail. Instead, it doesn't have any effect on the tag status (for
// example, activating the active tag).
func costexplorer_UpdateCostAllocationTagsStatus(cfg aws.Config, client *costexplorer.Client) {
	input := &costexplorer.UpdateCostAllocationTagsStatusInput{
		// CostAllocationTagsStatus: []types.CostAllocationTagStatusEntry, // Required
	}

	if len(_costexplorerCostAllocationTagsStatus) > 0 {
		if err := assignInputField(input, "CostAllocationTagsStatus", _costexplorerCostAllocationTagsStatus); err != nil {
			log.Errorf("invalid --cost-allocation-tags-status: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateCostAllocationTagsStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing cost category. Changes made to the cost category rules will
// be used to categorize the current month’s expenses and future expenses. This
// won’t change categorization for the previous months.
func costexplorer_UpdateCostCategoryDefinition(cfg aws.Config, client *costexplorer.Client) {
	input := &costexplorer.UpdateCostCategoryDefinitionInput{
		// CostCategoryArn: *string, // Required
		// RuleVersion: types.CostCategoryRuleVersion, // Required
		// Rules: []types.CostCategoryRule, // Required
	}

	if len(_costexplorerCostCategoryArn) > 0 {
		input.CostCategoryArn = aws.String(_costexplorerCostCategoryArn)
	}
	if len(_costexplorerRuleVersion) > 0 {
		if err := assignInputField(input, "RuleVersion", _costexplorerRuleVersion); err != nil {
			log.Errorf("invalid --rule-version: %s", err.Error())
			return
		}
	}
	if len(_costexplorerRules) > 0 {
		if err := assignInputField(input, "Rules", _costexplorerRules); err != nil {
			log.Errorf("invalid --rules: %s", err.Error())
			return
		}
	}
	if len(_costexplorerDefaultValue) > 0 {
		input.DefaultValue = aws.String(_costexplorerDefaultValue)
	}
	if len(_costexplorerEffectiveStart) > 0 {
		input.EffectiveStart = aws.String(_costexplorerEffectiveStart)
	}
	if len(_costexplorerSplitChargeRules) > 0 {
		if err := assignInputField(input, "SplitChargeRules", _costexplorerSplitChargeRules); err != nil {
			log.Errorf("invalid --split-charge-rules: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateCostCategoryDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_costexplorerCmd)
	_costexplorerCmd.Flags().SortFlags = false

	_costexplorerCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_costexplorerCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_costexplorerCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_costexplorerCmd.Flags().StringVarP(&_costexplorerAccountId, "account-id", "", "", "Account ID")
	_costexplorerCmd.Flags().StringVarP(&_costexplorerAccountScope, "account-scope", "", "", "Account Scope")
	_costexplorerCmd.Flags().StringVarP(&_costexplorerAnalysisId, "analysis-id", "", "", "Analysis ID")
	_costexplorerCmd.Flags().StringSliceVarP(&_costexplorerAnalysisIds, "analysis-ids", "", nil, "Analysis Ids")
	_costexplorerCmd.Flags().StringVarP(&_costexplorerAnalysisStatus, "analysis-status", "", "", "Analysis Status")
	_costexplorerCmd.Flags().StringVarP(&_costexplorerAnomalyId, "anomaly-id", "", "", "Anomaly ID")
	_costexplorerCmd.Flags().StringVarP(&_costexplorerAnomalyMonitor, "anomaly-monitor", "", "", "Anomaly Monitor")
	_costexplorerCmd.Flags().StringVarP(&_costexplorerAnomalySubscription, "anomaly-subscription", "", "", "Anomaly Subscription")
	_costexplorerCmd.Flags().StringVarP(&_costexplorerApproximationDimension, "approximation-dimension", "", "", "Approximation Dimension")
	_costexplorerCmd.Flags().StringVarP(&_costexplorerBackfillFrom, "backfill-from", "", "", "Backfill From")
	_costexplorerCmd.Flags().StringVarP(&_costexplorerBaselineTimePeriod, "baseline-time-period", "", "", "Baseline Time Period")
	_costexplorerCmd.Flags().StringVarP(&_costexplorerBillingViewArn, "billing-view-arn", "", "", "Billing View ARN")
	_costexplorerCmd.Flags().StringVarP(&_costexplorerCommitmentPurchaseAnalysisConfiguration, "commitment-purchase-analysis-configuration", "", "", "Commitment Purchase Analysis Configuration")
	_costexplorerCmd.Flags().StringVarP(&_costexplorerComparisonTimePeriod, "comparison-time-period", "", "", "Comparison Time Period")
	_costexplorerCmd.Flags().StringVarP(&_costexplorerConfiguration, "configuration", "", "", "Configuration")
	_costexplorerCmd.Flags().StringVarP(&_costexplorerContext, "context", "", "", "Context")
	_costexplorerCmd.Flags().StringVarP(&_costexplorerCostAllocationTagsStatus, "cost-allocation-tags-status", "", "", "Cost Allocation Tags Status")
	_costexplorerCmd.Flags().StringVarP(&_costexplorerCostCategoryArn, "cost-category-arn", "", "", "Cost Category ARN")
	_costexplorerCmd.Flags().StringVarP(&_costexplorerCostCategoryName, "cost-category-name", "", "", "Cost Category Name")
	_costexplorerCmd.Flags().StringVarP(&_costexplorerDataType, "data-type", "", "", "Data Type")
	_costexplorerCmd.Flags().StringVarP(&_costexplorerDateInterval, "date-interval", "", "", "Date Interval")
	_costexplorerCmd.Flags().StringVarP(&_costexplorerDefaultValue, "default-value", "", "", "Default Value")
	_costexplorerCmd.Flags().StringVarP(&_costexplorerDimension, "dimension", "", "", "Dimension")
	_costexplorerCmd.Flags().StringVarP(&_costexplorerEffectiveOn, "effective-on", "", "", "Effective On")
	_costexplorerCmd.Flags().StringVarP(&_costexplorerEffectiveStart, "effective-start", "", "", "Effective Start")
	_costexplorerCmd.Flags().StringVarP(&_costexplorerFeedback, "feedback", "", "", "Feedback")
	_costexplorerCmd.Flags().StringVarP(&_costexplorerFilter, "filter", "", "", "Filter")
	_costexplorerCmd.Flags().StringVarP(&_costexplorerFrequency, "frequency", "", "", "Frequency")
	_costexplorerCmd.Flags().StringVarP(&_costexplorerGenerationStatus, "generation-status", "", "", "Generation Status")
	_costexplorerCmd.Flags().StringVarP(&_costexplorerGranularity, "granularity", "", "", "Granularity")
	_costexplorerCmd.Flags().StringVarP(&_costexplorerGroupBy, "group-by", "", "", "Group By")
	_costexplorerCmd.Flags().StringVarP(&_costexplorerLookbackPeriodInDays, "lookback-period-in-days", "", "", "Lookback Period In Days")
	_costexplorerCmd.Flags().StringVarP(&_costexplorerMaxResults, "max-results", "", "", "Max Results")
	_costexplorerCmd.Flags().StringVarP(&_costexplorerMetric, "metric", "", "", "Metric")
	_costexplorerCmd.Flags().StringVarP(&_costexplorerMetricForComparison, "metric-for-comparison", "", "", "Metric For Comparison")
	_costexplorerCmd.Flags().StringSliceVarP(&_costexplorerMetrics, "metrics", "", nil, "Metrics")
	_costexplorerCmd.Flags().StringVarP(&_costexplorerMonitorArn, "monitor-arn", "", "", "Monitor ARN")
	_costexplorerCmd.Flags().StringSliceVarP(&_costexplorerMonitorArnList, "monitor-arn-list", "", nil, "Monitor ARN List")
	_costexplorerCmd.Flags().StringVarP(&_costexplorerMonitorName, "monitor-name", "", "", "Monitor Name")
	_costexplorerCmd.Flags().StringVarP(&_costexplorerName, "name", "", "", "Name")
	_costexplorerCmd.Flags().StringVarP(&_costexplorerNextPageToken, "next-page-token", "", "", "Next Page Token")
	_costexplorerCmd.Flags().StringVarP(&_costexplorerNextToken, "next-token", "", "", "Next Token")
	_costexplorerCmd.Flags().StringVarP(&_costexplorerPageSize, "page-size", "", "", "Page Size")
	_costexplorerCmd.Flags().StringVarP(&_costexplorerPaymentOption, "payment-option", "", "", "Payment Option")
	_costexplorerCmd.Flags().StringVarP(&_costexplorerPredictionIntervalLevel, "prediction-interval-level", "", "", "Prediction Interval Level")
	_costexplorerCmd.Flags().StringVarP(&_costexplorerRecommendationDetailId, "recommendation-detail-id", "", "", "Recommendation Detail ID")
	_costexplorerCmd.Flags().StringSliceVarP(&_costexplorerRecommendationIds, "recommendation-ids", "", nil, "Recommendation Ids")
	_costexplorerCmd.Flags().StringVarP(&_costexplorerResourceArn, "resource-arn", "", "", "Resource ARN")
	_costexplorerCmd.Flags().StringSliceVarP(&_costexplorerResourceTagKeys, "resource-tag-keys", "", nil, "Resource Tag Keys")
	_costexplorerCmd.Flags().StringVarP(&_costexplorerResourceTags, "resource-tags", "", "", "Resource Tags")
	_costexplorerCmd.Flags().StringVarP(&_costexplorerRuleVersion, "rule-version", "", "", "Rule Version")
	_costexplorerCmd.Flags().StringVarP(&_costexplorerRules, "rules", "", "", "Rules")
	_costexplorerCmd.Flags().StringVarP(&_costexplorerSavingsPlansType, "savings-plans-type", "", "", "Savings Plans Type")
	_costexplorerCmd.Flags().StringVarP(&_costexplorerSearchString, "search-string", "", "", "Search String")
	_costexplorerCmd.Flags().StringVarP(&_costexplorerService, "service", "", "", "Service")
	_costexplorerCmd.Flags().StringVarP(&_costexplorerServiceSpecification, "service-specification", "", "", "Service Specification")
	_costexplorerCmd.Flags().StringSliceVarP(&_costexplorerServices, "services", "", nil, "Services")
	_costexplorerCmd.Flags().StringVarP(&_costexplorerSortBy, "sort-by", "", "", "Sort By")
	_costexplorerCmd.Flags().StringVarP(&_costexplorerSplitChargeRules, "split-charge-rules", "", "", "Split Charge Rules")
	_costexplorerCmd.Flags().StringVarP(&_costexplorerStatus, "status", "", "", "Status")
	_costexplorerCmd.Flags().StringVarP(&_costexplorerSubscribers, "subscribers", "", "", "Subscribers")
	_costexplorerCmd.Flags().StringVarP(&_costexplorerSubscriptionArn, "subscription-arn", "", "", "Subscription ARN")
	_costexplorerCmd.Flags().StringSliceVarP(&_costexplorerSubscriptionArnList, "subscription-arn-list", "", nil, "Subscription ARN List")
	_costexplorerCmd.Flags().StringVarP(&_costexplorerSubscriptionName, "subscription-name", "", "", "Subscription Name")
	_costexplorerCmd.Flags().StringSliceVarP(&_costexplorerSupportedResourceTypes, "supported-resource-types", "", nil, "Supported Resource Types")
	_costexplorerCmd.Flags().StringVarP(&_costexplorerTagKey, "tag-key", "", "", "Tag Key")
	_costexplorerCmd.Flags().StringSliceVarP(&_costexplorerTagKeys, "tag-keys", "", nil, "Tag Keys")
	_costexplorerCmd.Flags().StringVarP(&_costexplorerTermInYears, "term-in-years", "", "", "Term In Years")
	_costexplorerCmd.Flags().StringVarP(&_costexplorerThreshold, "threshold", "", "", "Threshold")
	_costexplorerCmd.Flags().StringVarP(&_costexplorerThresholdExpression, "threshold-expression", "", "", "Threshold Expression")
	_costexplorerCmd.Flags().StringVarP(&_costexplorerTimePeriod, "time-period", "", "", "Time Period")
	_costexplorerCmd.Flags().StringVarP(&_costexplorerTotalImpact, "total-impact", "", "", "Total Impact")
	_costexplorerCmd.Flags().StringVarP(&_costexplorerType, "type", "", "", "Type")

	_costexplorerCmd.Flags().BoolVarP(&_costexplorerCreateAnomalyMonitor, "create-anomaly-monitor", "", false, "Create Anomaly Monitor")
	_costexplorerCmd.Flags().BoolVarP(&_costexplorerCreateAnomalySubscription, "create-anomaly-subscription", "", false, "Create Anomaly Subscription")
	_costexplorerCmd.Flags().BoolVarP(&_costexplorerCreateCostCategoryDefinition, "create-cost-category-definition", "", false, "Create Cost Category Definition")
	_costexplorerCmd.Flags().BoolVarP(&_costexplorerDeleteAnomalyMonitor, "delete-anomaly-monitor", "", false, "Delete Anomaly Monitor")
	_costexplorerCmd.Flags().BoolVarP(&_costexplorerDeleteAnomalySubscription, "delete-anomaly-subscription", "", false, "Delete Anomaly Subscription")
	_costexplorerCmd.Flags().BoolVarP(&_costexplorerDeleteCostCategoryDefinition, "delete-cost-category-definition", "", false, "Delete Cost Category Definition")
	_costexplorerCmd.Flags().BoolVarP(&_costexplorerDescribeCostCategoryDefinition, "describe-cost-category-definition", "", false, "Describe Cost Category Definition")
	_costexplorerCmd.Flags().BoolVarP(&_costexplorerGetAnomalies, "get-anomalies", "", false, "Get Anomalies")
	_costexplorerCmd.Flags().BoolVarP(&_costexplorerGetAnomalyMonitors, "get-anomaly-monitors", "", false, "Get Anomaly Monitors")
	_costexplorerCmd.Flags().BoolVarP(&_costexplorerGetAnomalySubscriptions, "get-anomaly-subscriptions", "", false, "Get Anomaly Subscriptions")
	_costexplorerCmd.Flags().BoolVarP(&_costexplorerGetApproximateUsageRecords, "get-approximate-usage-records", "", false, "Get Approximate Usage Records")
	_costexplorerCmd.Flags().BoolVarP(&_costexplorerGetCommitmentPurchaseAnalysis, "get-commitment-purchase-analysis", "", false, "Get Commitment Purchase Analysis")
	_costexplorerCmd.Flags().BoolVarP(&_costexplorerGetCostAndUsage, "get-cost-and-usage", "", false, "Get Cost And Usage")
	_costexplorerCmd.Flags().BoolVarP(&_costexplorerGetCostAndUsageComparisons, "get-cost-and-usage-comparisons", "", false, "Get Cost And Usage Comparisons")
	_costexplorerCmd.Flags().BoolVarP(&_costexplorerGetCostAndUsageWithResources, "get-cost-and-usage-with-resources", "", false, "Get Cost And Usage With Resources")
	_costexplorerCmd.Flags().BoolVarP(&_costexplorerGetCostCategories, "get-cost-categories", "", false, "Get Cost Categories")
	_costexplorerCmd.Flags().BoolVarP(&_costexplorerGetCostComparisonDrivers, "get-cost-comparison-drivers", "", false, "Get Cost Comparison Drivers")
	_costexplorerCmd.Flags().BoolVarP(&_costexplorerGetCostForecast, "get-cost-forecast", "", false, "Get Cost Forecast")
	_costexplorerCmd.Flags().BoolVarP(&_costexplorerGetDimensionValues, "get-dimension-values", "", false, "Get Dimension Values")
	_costexplorerCmd.Flags().BoolVarP(&_costexplorerGetReservationCoverage, "get-reservation-coverage", "", false, "Get Reservation Coverage")
	_costexplorerCmd.Flags().BoolVarP(&_costexplorerGetReservationPurchaseRecommendation, "get-reservation-purchase-recommendation", "", false, "Get Reservation Purchase Recommendation")
	_costexplorerCmd.Flags().BoolVarP(&_costexplorerGetReservationUtilization, "get-reservation-utilization", "", false, "Get Reservation Utilization")
	_costexplorerCmd.Flags().BoolVarP(&_costexplorerGetRightsizingRecommendation, "get-rightsizing-recommendation", "", false, "Get Rightsizing Recommendation")
	_costexplorerCmd.Flags().BoolVarP(&_costexplorerGetSavingsPlanPurchaseRecommendationDetails, "get-savings-plan-purchase-recommendation-details", "", false, "Get Savings Plan Purchase Recommendation Details")
	_costexplorerCmd.Flags().BoolVarP(&_costexplorerGetSavingsPlansCoverage, "get-savings-plans-coverage", "", false, "Get Savings Plans Coverage")
	_costexplorerCmd.Flags().BoolVarP(&_costexplorerGetSavingsPlansPurchaseRecommendation, "get-savings-plans-purchase-recommendation", "", false, "Get Savings Plans Purchase Recommendation")
	_costexplorerCmd.Flags().BoolVarP(&_costexplorerGetSavingsPlansUtilization, "get-savings-plans-utilization", "", false, "Get Savings Plans Utilization")
	_costexplorerCmd.Flags().BoolVarP(&_costexplorerGetSavingsPlansUtilizationDetails, "get-savings-plans-utilization-details", "", false, "Get Savings Plans Utilization Details")
	_costexplorerCmd.Flags().BoolVarP(&_costexplorerGetTags, "get-tags", "", false, "Get Tags")
	_costexplorerCmd.Flags().BoolVarP(&_costexplorerGetUsageForecast, "get-usage-forecast", "", false, "Get Usage Forecast")
	_costexplorerCmd.Flags().BoolVarP(&_costexplorerListCommitmentPurchaseAnalyses, "list-commitment-purchase-analyses", "", false, "List Commitment Purchase Analyses")
	_costexplorerCmd.Flags().BoolVarP(&_costexplorerListCostAllocationTagBackfillHistory, "list-cost-allocation-tag-backfill-history", "", false, "List Cost Allocation Tag Backfill History")
	_costexplorerCmd.Flags().BoolVarP(&_costexplorerListCostAllocationTags, "list-cost-allocation-tags", "", false, "List Cost Allocation Tags")
	_costexplorerCmd.Flags().BoolVarP(&_costexplorerListCostCategoryDefinitions, "list-cost-category-definitions", "", false, "List Cost Category Definitions")
	_costexplorerCmd.Flags().BoolVarP(&_costexplorerListCostCategoryResourceAssociations, "list-cost-category-resource-associations", "", false, "List Cost Category Resource Associations")
	_costexplorerCmd.Flags().BoolVarP(&_costexplorerListSavingsPlansPurchaseRecommendationGeneration, "list-savings-plans-purchase-recommendation-generation", "", false, "List Savings Plans Purchase Recommendation Generation")
	_costexplorerCmd.Flags().BoolVarP(&_costexplorerListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_costexplorerCmd.Flags().BoolVarP(&_costexplorerProvideAnomalyFeedback, "provide-anomaly-feedback", "", false, "Provide Anomaly Feedback")
	_costexplorerCmd.Flags().BoolVarP(&_costexplorerStartCommitmentPurchaseAnalysis, "start-commitment-purchase-analysis", "", false, "Start Commitment Purchase Analysis")
	_costexplorerCmd.Flags().BoolVarP(&_costexplorerStartCostAllocationTagBackfill, "start-cost-allocation-tag-backfill", "", false, "Start Cost Allocation Tag Backfill")
	_costexplorerCmd.Flags().BoolVarP(&_costexplorerStartSavingsPlansPurchaseRecommendationGeneration, "start-savings-plans-purchase-recommendation-generation", "", false, "Start Savings Plans Purchase Recommendation Generation")
	_costexplorerCmd.Flags().BoolVarP(&_costexplorerTagResource, "tag-resource", "", false, "Tag Resource")
	_costexplorerCmd.Flags().BoolVarP(&_costexplorerUntagResource, "untag-resource", "", false, "Untag Resource")
	_costexplorerCmd.Flags().BoolVarP(&_costexplorerUpdateAnomalyMonitor, "update-anomaly-monitor", "", false, "Update Anomaly Monitor")
	_costexplorerCmd.Flags().BoolVarP(&_costexplorerUpdateAnomalySubscription, "update-anomaly-subscription", "", false, "Update Anomaly Subscription")
	_costexplorerCmd.Flags().BoolVarP(&_costexplorerUpdateCostAllocationTagsStatus, "update-cost-allocation-tags-status", "", false, "Update Cost Allocation Tags Status")
	_costexplorerCmd.Flags().BoolVarP(&_costexplorerUpdateCostCategoryDefinition, "update-cost-category-definition", "", false, "Update Cost Category Definition")

}
