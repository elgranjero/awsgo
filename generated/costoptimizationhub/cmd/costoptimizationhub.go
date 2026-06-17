package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/costoptimizationhub"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// costoptimizationhubCmd represents the costoptimizationhub command
var _costoptimizationhubCmd = &cobra.Command{
	Use:   "costoptimizationhub",
	Short: "AWS costoptimizationhub CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := costoptimizationhub.NewFromConfig(cfg)
		if _costoptimizationhubGetPreferences {
			costoptimizationhub_GetPreferences(cfg, client)
			return
		}
		if _costoptimizationhubGetRecommendation {
			costoptimizationhub_GetRecommendation(cfg, client)
			return
		}
		if _costoptimizationhubListEfficiencyMetrics {
			costoptimizationhub_ListEfficiencyMetrics(cfg, client)
			return
		}
		if _costoptimizationhubListEnrollmentStatuses {
			costoptimizationhub_ListEnrollmentStatuses(cfg, client)
			return
		}
		if _costoptimizationhubListRecommendationSummaries {
			costoptimizationhub_ListRecommendationSummaries(cfg, client)
			return
		}
		if _costoptimizationhubListRecommendations {
			costoptimizationhub_ListRecommendations(cfg, client)
			return
		}
		if _costoptimizationhubUpdateEnrollmentStatus {
			costoptimizationhub_UpdateEnrollmentStatus(cfg, client)
			return
		}
		if _costoptimizationhubUpdatePreferences {
			costoptimizationhub_UpdatePreferences(cfg, client)
			return
		}

	},
}

var (
	_costoptimizationhubGetPreferences              bool
	_costoptimizationhubGetRecommendation           bool
	_costoptimizationhubListEfficiencyMetrics       bool
	_costoptimizationhubListEnrollmentStatuses      bool
	_costoptimizationhubListRecommendationSummaries bool
	_costoptimizationhubListRecommendations         bool
	_costoptimizationhubUpdateEnrollmentStatus      bool
	_costoptimizationhubUpdatePreferences           bool

	_costoptimizationhubAccountId                       string
	_costoptimizationhubFilter                          string
	_costoptimizationhubGranularity                     string
	_costoptimizationhubGroupBy                         string
	_costoptimizationhubIncludeAllRecommendations       string
	_costoptimizationhubIncludeMemberAccounts           string
	_costoptimizationhubIncludeOrganizationInfo         string
	_costoptimizationhubMaxResults                      string
	_costoptimizationhubMemberAccountDiscountVisibility string
	_costoptimizationhubMetrics                         string
	_costoptimizationhubNextToken                       string
	_costoptimizationhubOrderBy                         string
	_costoptimizationhubPreferredCommitment             string
	_costoptimizationhubRecommendationId                string
	_costoptimizationhubSavingsEstimationMode           string
	_costoptimizationhubStatus                          string
	_costoptimizationhubTimePeriod                      string
)

// Returns a set of preferences for an account in order to add account-specific
// preferences into the service. These preferences impact how the savings
// associated with recommendations are presented—estimated savings after discounts
// or estimated savings before discounts, for example.
func costoptimizationhub_GetPreferences(cfg aws.Config, client *costoptimizationhub.Client) {
	input := &costoptimizationhub.GetPreferencesInput{}

	if resp, err := client.GetPreferences(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns both the current and recommended resource configuration and the
// estimated cost impact for a recommendation.
//
// The recommendationId is only valid for up to a maximum of 24 hours as
// recommendations are refreshed daily. To retrieve the recommendationId , use the
// ListRecommendations API.
func costoptimizationhub_GetRecommendation(cfg aws.Config, client *costoptimizationhub.Client) {
	input := &costoptimizationhub.GetRecommendationInput{
		// RecommendationId: *string, // Required
	}

	if len(_costoptimizationhubRecommendationId) > 0 {
		input.RecommendationId = aws.String(_costoptimizationhubRecommendationId)
	}

	if resp, err := client.GetRecommendation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns cost efficiency metrics aggregated over time and optionally grouped by
// a specified dimension. The metrics provide insights into your cost optimization
// progress by tracking estimated savings, spending, and measures how effectively
// you're optimizing your Cloud resources.
//
// The operation supports both daily and monthly time granularities and allows
// grouping results by account ID, Amazon Web Services Region. Results are returned
// as time-series data, enabling you to analyze trends in your cost optimization
// performance over the specified time period.
func costoptimizationhub_ListEfficiencyMetrics(cfg aws.Config, client *costoptimizationhub.Client) {
	input := &costoptimizationhub.ListEfficiencyMetricsInput{
		// Granularity: types.GranularityType, // Required
		// TimePeriod: *types.TimePeriod, // Required
	}

	if len(_costoptimizationhubGranularity) > 0 {
		if err := assignInputField(input, "Granularity", _costoptimizationhubGranularity); err != nil {
			log.Errorf("invalid --granularity: %s", err.Error())
			return
		}
	}
	if len(_costoptimizationhubTimePeriod) > 0 {
		if err := assignInputField(input, "TimePeriod", _costoptimizationhubTimePeriod); err != nil {
			log.Errorf("invalid --time-period: %s", err.Error())
			return
		}
	}
	if len(_costoptimizationhubGroupBy) > 0 {
		input.GroupBy = aws.String(_costoptimizationhubGroupBy)
	}
	if len(_costoptimizationhubMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _costoptimizationhubMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_costoptimizationhubNextToken) > 0 {
		input.NextToken = aws.String(_costoptimizationhubNextToken)
	}
	if len(_costoptimizationhubOrderBy) > 0 {
		if err := assignInputField(input, "OrderBy", _costoptimizationhubOrderBy); err != nil {
			log.Errorf("invalid --order-by: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListEfficiencyMetrics(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*costoptimizationhub.ListEfficiencyMetricsOutput
	p := costoptimizationhub.NewListEfficiencyMetricsPaginator(client, input)
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

// Retrieves the enrollment status for an account. It can also return the list of
// accounts that are enrolled under the organization.
func costoptimizationhub_ListEnrollmentStatuses(cfg aws.Config, client *costoptimizationhub.Client) {
	input := &costoptimizationhub.ListEnrollmentStatusesInput{}

	if len(_costoptimizationhubAccountId) > 0 {
		input.AccountId = aws.String(_costoptimizationhubAccountId)
	}
	if len(_costoptimizationhubIncludeOrganizationInfo) > 0 {
		if err := assignInputField(input, "IncludeOrganizationInfo", _costoptimizationhubIncludeOrganizationInfo); err != nil {
			log.Errorf("invalid --include-organization-info: %s", err.Error())
			return
		}
	}
	if len(_costoptimizationhubMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _costoptimizationhubMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_costoptimizationhubNextToken) > 0 {
		input.NextToken = aws.String(_costoptimizationhubNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListEnrollmentStatuses(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*costoptimizationhub.ListEnrollmentStatusesOutput
	p := costoptimizationhub.NewListEnrollmentStatusesPaginator(client, input)
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

// Returns a concise representation of savings estimates for resources. Also
// returns de-duped savings across different types of recommendations.
//
// The following filters are not supported for this API: recommendationIds ,
// resourceArns , and resourceIds .
func costoptimizationhub_ListRecommendationSummaries(cfg aws.Config, client *costoptimizationhub.Client) {
	input := &costoptimizationhub.ListRecommendationSummariesInput{
		// GroupBy: *string, // Required
	}

	if len(_costoptimizationhubGroupBy) > 0 {
		input.GroupBy = aws.String(_costoptimizationhubGroupBy)
	}
	if len(_costoptimizationhubFilter) > 0 {
		if err := assignInputField(input, "Filter", _costoptimizationhubFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_costoptimizationhubMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _costoptimizationhubMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_costoptimizationhubMetrics) > 0 {
		if err := assignInputField(input, "Metrics", _costoptimizationhubMetrics); err != nil {
			log.Errorf("invalid --metrics: %s", err.Error())
			return
		}
	}
	if len(_costoptimizationhubNextToken) > 0 {
		input.NextToken = aws.String(_costoptimizationhubNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRecommendationSummaries(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*costoptimizationhub.ListRecommendationSummariesOutput
	p := costoptimizationhub.NewListRecommendationSummariesPaginator(client, input)
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

// Returns a list of recommendations.
func costoptimizationhub_ListRecommendations(cfg aws.Config, client *costoptimizationhub.Client) {
	input := &costoptimizationhub.ListRecommendationsInput{}

	if len(_costoptimizationhubFilter) > 0 {
		if err := assignInputField(input, "Filter", _costoptimizationhubFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_costoptimizationhubIncludeAllRecommendations) > 0 {
		if err := assignInputField(input, "IncludeAllRecommendations", _costoptimizationhubIncludeAllRecommendations); err != nil {
			log.Errorf("invalid --include-all-recommendations: %s", err.Error())
			return
		}
	}
	if len(_costoptimizationhubMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _costoptimizationhubMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_costoptimizationhubNextToken) > 0 {
		input.NextToken = aws.String(_costoptimizationhubNextToken)
	}
	if len(_costoptimizationhubOrderBy) > 0 {
		if err := assignInputField(input, "OrderBy", _costoptimizationhubOrderBy); err != nil {
			log.Errorf("invalid --order-by: %s", err.Error())
			return
		}
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

	var results []*costoptimizationhub.ListRecommendationsOutput
	p := costoptimizationhub.NewListRecommendationsPaginator(client, input)
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

// Updates the enrollment (opt in and opt out) status of an account to the Cost
// Optimization Hub service.
//
// If the account is a management account of an organization, this action can also
// be used to enroll member accounts of the organization.
//
// You must have the appropriate permissions to opt in to Cost Optimization Hub
// and to view its recommendations. When you opt in, Cost Optimization Hub
// automatically creates a service-linked role in your account to access its data.
func costoptimizationhub_UpdateEnrollmentStatus(cfg aws.Config, client *costoptimizationhub.Client) {
	input := &costoptimizationhub.UpdateEnrollmentStatusInput{
		// Status: types.EnrollmentStatus, // Required
	}

	if len(_costoptimizationhubStatus) > 0 {
		if err := assignInputField(input, "Status", _costoptimizationhubStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_costoptimizationhubIncludeMemberAccounts) > 0 {
		if err := assignInputField(input, "IncludeMemberAccounts", _costoptimizationhubIncludeMemberAccounts); err != nil {
			log.Errorf("invalid --include-member-accounts: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateEnrollmentStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a set of preferences for an account in order to add account-specific
// preferences into the service. These preferences impact how the savings
// associated with recommendations are presented.
func costoptimizationhub_UpdatePreferences(cfg aws.Config, client *costoptimizationhub.Client) {
	input := &costoptimizationhub.UpdatePreferencesInput{}

	if len(_costoptimizationhubMemberAccountDiscountVisibility) > 0 {
		if err := assignInputField(input, "MemberAccountDiscountVisibility", _costoptimizationhubMemberAccountDiscountVisibility); err != nil {
			log.Errorf("invalid --member-account-discount-visibility: %s", err.Error())
			return
		}
	}
	if len(_costoptimizationhubPreferredCommitment) > 0 {
		if err := assignInputField(input, "PreferredCommitment", _costoptimizationhubPreferredCommitment); err != nil {
			log.Errorf("invalid --preferred-commitment: %s", err.Error())
			return
		}
	}
	if len(_costoptimizationhubSavingsEstimationMode) > 0 {
		if err := assignInputField(input, "SavingsEstimationMode", _costoptimizationhubSavingsEstimationMode); err != nil {
			log.Errorf("invalid --savings-estimation-mode: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdatePreferences(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_costoptimizationhubCmd)
	_costoptimizationhubCmd.Flags().SortFlags = false

	_costoptimizationhubCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_costoptimizationhubCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_costoptimizationhubCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_costoptimizationhubCmd.Flags().StringVarP(&_costoptimizationhubAccountId, "account-id", "", "", "Account ID")
	_costoptimizationhubCmd.Flags().StringVarP(&_costoptimizationhubFilter, "filter", "", "", "Filter")
	_costoptimizationhubCmd.Flags().StringVarP(&_costoptimizationhubGranularity, "granularity", "", "", "Granularity")
	_costoptimizationhubCmd.Flags().StringVarP(&_costoptimizationhubGroupBy, "group-by", "", "", "Group By")
	_costoptimizationhubCmd.Flags().StringVarP(&_costoptimizationhubIncludeAllRecommendations, "include-all-recommendations", "", "", "Include All Recommendations")
	_costoptimizationhubCmd.Flags().StringVarP(&_costoptimizationhubIncludeMemberAccounts, "include-member-accounts", "", "", "Include Member Accounts")
	_costoptimizationhubCmd.Flags().StringVarP(&_costoptimizationhubIncludeOrganizationInfo, "include-organization-info", "", "", "Include Organization Info")
	_costoptimizationhubCmd.Flags().StringVarP(&_costoptimizationhubMaxResults, "max-results", "", "", "Max Results")
	_costoptimizationhubCmd.Flags().StringVarP(&_costoptimizationhubMemberAccountDiscountVisibility, "member-account-discount-visibility", "", "", "Member Account Discount Visibility")
	_costoptimizationhubCmd.Flags().StringVarP(&_costoptimizationhubMetrics, "metrics", "", "", "Metrics")
	_costoptimizationhubCmd.Flags().StringVarP(&_costoptimizationhubNextToken, "next-token", "", "", "Next Token")
	_costoptimizationhubCmd.Flags().StringVarP(&_costoptimizationhubOrderBy, "order-by", "", "", "Order By")
	_costoptimizationhubCmd.Flags().StringVarP(&_costoptimizationhubPreferredCommitment, "preferred-commitment", "", "", "Preferred Commitment")
	_costoptimizationhubCmd.Flags().StringVarP(&_costoptimizationhubRecommendationId, "recommendation-id", "", "", "Recommendation ID")
	_costoptimizationhubCmd.Flags().StringVarP(&_costoptimizationhubSavingsEstimationMode, "savings-estimation-mode", "", "", "Savings Estimation Mode")
	_costoptimizationhubCmd.Flags().StringVarP(&_costoptimizationhubStatus, "status", "", "", "Status")
	_costoptimizationhubCmd.Flags().StringVarP(&_costoptimizationhubTimePeriod, "time-period", "", "", "Time Period")

	_costoptimizationhubCmd.Flags().BoolVarP(&_costoptimizationhubGetPreferences, "get-preferences", "", false, "Get Preferences")
	_costoptimizationhubCmd.Flags().BoolVarP(&_costoptimizationhubGetRecommendation, "get-recommendation", "", false, "Get Recommendation")
	_costoptimizationhubCmd.Flags().BoolVarP(&_costoptimizationhubListEfficiencyMetrics, "list-efficiency-metrics", "", false, "List Efficiency Metrics")
	_costoptimizationhubCmd.Flags().BoolVarP(&_costoptimizationhubListEnrollmentStatuses, "list-enrollment-statuses", "", false, "List Enrollment Statuses")
	_costoptimizationhubCmd.Flags().BoolVarP(&_costoptimizationhubListRecommendationSummaries, "list-recommendation-summaries", "", false, "List Recommendation Summaries")
	_costoptimizationhubCmd.Flags().BoolVarP(&_costoptimizationhubListRecommendations, "list-recommendations", "", false, "List Recommendations")
	_costoptimizationhubCmd.Flags().BoolVarP(&_costoptimizationhubUpdateEnrollmentStatus, "update-enrollment-status", "", false, "Update Enrollment Status")
	_costoptimizationhubCmd.Flags().BoolVarP(&_costoptimizationhubUpdatePreferences, "update-preferences", "", false, "Update Preferences")

}
