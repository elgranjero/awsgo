package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/trustedadvisor"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// trustedadvisorCmd represents the trustedadvisor command
var _trustedadvisorCmd = &cobra.Command{
	Use:   "trustedadvisor",
	Short: "AWS trustedadvisor CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := trustedadvisor.NewFromConfig(cfg)
		if _trustedadvisorBatchUpdateRecommendationResourceExclusion {
			trustedadvisor_BatchUpdateRecommendationResourceExclusion(cfg, client)
			return
		}
		if _trustedadvisorGetOrganizationRecommendation {
			trustedadvisor_GetOrganizationRecommendation(cfg, client)
			return
		}
		if _trustedadvisorGetRecommendation {
			trustedadvisor_GetRecommendation(cfg, client)
			return
		}
		if _trustedadvisorListChecks {
			trustedadvisor_ListChecks(cfg, client)
			return
		}
		if _trustedadvisorListOrganizationRecommendationAccounts {
			trustedadvisor_ListOrganizationRecommendationAccounts(cfg, client)
			return
		}
		if _trustedadvisorListOrganizationRecommendationResources {
			trustedadvisor_ListOrganizationRecommendationResources(cfg, client)
			return
		}
		if _trustedadvisorListOrganizationRecommendations {
			trustedadvisor_ListOrganizationRecommendations(cfg, client)
			return
		}
		if _trustedadvisorListRecommendationResources {
			trustedadvisor_ListRecommendationResources(cfg, client)
			return
		}
		if _trustedadvisorListRecommendations {
			trustedadvisor_ListRecommendations(cfg, client)
			return
		}
		if _trustedadvisorUpdateOrganizationRecommendationLifecycle {
			trustedadvisor_UpdateOrganizationRecommendationLifecycle(cfg, client)
			return
		}
		if _trustedadvisorUpdateRecommendationLifecycle {
			trustedadvisor_UpdateRecommendationLifecycle(cfg, client)
			return
		}

	},
}

var (
	_trustedadvisorBatchUpdateRecommendationResourceExclusion bool
	_trustedadvisorGetOrganizationRecommendation              bool
	_trustedadvisorGetRecommendation                          bool
	_trustedadvisorListChecks                                 bool
	_trustedadvisorListOrganizationRecommendationAccounts     bool
	_trustedadvisorListOrganizationRecommendationResources    bool
	_trustedadvisorListOrganizationRecommendations            bool
	_trustedadvisorListRecommendationResources                bool
	_trustedadvisorListRecommendations                        bool
	_trustedadvisorUpdateOrganizationRecommendationLifecycle  bool
	_trustedadvisorUpdateRecommendationLifecycle              bool

	_trustedadvisorAffectedAccountId                    string
	_trustedadvisorAfterLastUpdatedAt                   string
	_trustedadvisorAwsService                           string
	_trustedadvisorBeforeLastUpdatedAt                  string
	_trustedadvisorCheckIdentifier                      string
	_trustedadvisorExclusionStatus                      string
	_trustedadvisorLanguage                             string
	_trustedadvisorLifecycleStage                       string
	_trustedadvisorMaxResults                           string
	_trustedadvisorNextToken                            string
	_trustedadvisorOrganizationRecommendationIdentifier string
	_trustedadvisorPillar                               string
	_trustedadvisorRecommendationIdentifier             string
	_trustedadvisorRecommendationResourceExclusions     string
	_trustedadvisorRegionCode                           string
	_trustedadvisorSource                               string
	_trustedadvisorStatus                               string
	_trustedadvisorType                                 string
	_trustedadvisorUpdateReason                         string
	_trustedadvisorUpdateReasonCode                     string
)

// Update one or more exclusion statuses for a list of recommendation resources.
// This API supports up to 25 unique recommendation resource ARNs per request. This
// API currently doesn't support prioritized recommendation resources. This API
// updates global recommendations, eliminating the need to call the API in each AWS
// Region. After submitting an exclusion update, note that it might take a few
// minutes for the changes to be reflected in the system.
func trustedadvisor_BatchUpdateRecommendationResourceExclusion(cfg aws.Config, client *trustedadvisor.Client) {
	input := &trustedadvisor.BatchUpdateRecommendationResourceExclusionInput{
		// RecommendationResourceExclusions: []types.RecommendationResourceExclusion, // Required
	}

	if len(_trustedadvisorRecommendationResourceExclusions) > 0 {
		if err := assignInputField(input, "RecommendationResourceExclusions", _trustedadvisorRecommendationResourceExclusions); err != nil {
			log.Errorf("invalid --recommendation-resource-exclusions: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchUpdateRecommendationResourceExclusion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get a specific recommendation within an AWS Organizations organization. This
// API supports only prioritized recommendations and provides global priority
// recommendations, eliminating the need to call the API in each AWS Region.
func trustedadvisor_GetOrganizationRecommendation(cfg aws.Config, client *trustedadvisor.Client) {
	input := &trustedadvisor.GetOrganizationRecommendationInput{
		// OrganizationRecommendationIdentifier: *string, // Required
	}

	if len(_trustedadvisorOrganizationRecommendationIdentifier) > 0 {
		input.OrganizationRecommendationIdentifier = aws.String(_trustedadvisorOrganizationRecommendationIdentifier)
	}

	if resp, err := client.GetOrganizationRecommendation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get a specific Recommendation. This API provides global recommendations,
// eliminating the need to call the API in each AWS Region.
func trustedadvisor_GetRecommendation(cfg aws.Config, client *trustedadvisor.Client) {
	input := &trustedadvisor.GetRecommendationInput{
		// RecommendationIdentifier: *string, // Required
	}

	if len(_trustedadvisorRecommendationIdentifier) > 0 {
		input.RecommendationIdentifier = aws.String(_trustedadvisorRecommendationIdentifier)
	}
	if len(_trustedadvisorLanguage) > 0 {
		if err := assignInputField(input, "Language", _trustedadvisorLanguage); err != nil {
			log.Errorf("invalid --language: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetRecommendation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// List a filterable set of Checks. This API provides global recommendations,
// eliminating the need to call the API in each AWS Region.
func trustedadvisor_ListChecks(cfg aws.Config, client *trustedadvisor.Client) {
	input := &trustedadvisor.ListChecksInput{}

	if len(_trustedadvisorAwsService) > 0 {
		input.AwsService = aws.String(_trustedadvisorAwsService)
	}
	if len(_trustedadvisorLanguage) > 0 {
		if err := assignInputField(input, "Language", _trustedadvisorLanguage); err != nil {
			log.Errorf("invalid --language: %s", err.Error())
			return
		}
	}
	if len(_trustedadvisorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _trustedadvisorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_trustedadvisorNextToken) > 0 {
		input.NextToken = aws.String(_trustedadvisorNextToken)
	}
	if len(_trustedadvisorPillar) > 0 {
		if err := assignInputField(input, "Pillar", _trustedadvisorPillar); err != nil {
			log.Errorf("invalid --pillar: %s", err.Error())
			return
		}
	}
	if len(_trustedadvisorSource) > 0 {
		if err := assignInputField(input, "Source", _trustedadvisorSource); err != nil {
			log.Errorf("invalid --source: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListChecks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*trustedadvisor.ListChecksOutput
	p := trustedadvisor.NewListChecksPaginator(client, input)
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

// Lists the accounts that own the resources for an organization aggregate
// recommendation. This API only supports prioritized recommendations and provides
// global priority recommendations, eliminating the need to call the API in each
// AWS Region.
func trustedadvisor_ListOrganizationRecommendationAccounts(cfg aws.Config, client *trustedadvisor.Client) {
	input := &trustedadvisor.ListOrganizationRecommendationAccountsInput{
		// OrganizationRecommendationIdentifier: *string, // Required
	}

	if len(_trustedadvisorOrganizationRecommendationIdentifier) > 0 {
		input.OrganizationRecommendationIdentifier = aws.String(_trustedadvisorOrganizationRecommendationIdentifier)
	}
	if len(_trustedadvisorAffectedAccountId) > 0 {
		input.AffectedAccountId = aws.String(_trustedadvisorAffectedAccountId)
	}
	if len(_trustedadvisorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _trustedadvisorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_trustedadvisorNextToken) > 0 {
		input.NextToken = aws.String(_trustedadvisorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListOrganizationRecommendationAccounts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*trustedadvisor.ListOrganizationRecommendationAccountsOutput
	p := trustedadvisor.NewListOrganizationRecommendationAccountsPaginator(client, input)
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

// List Resources of a Recommendation within an Organization. This API only
// supports prioritized recommendations and provides global priority
// recommendations, eliminating the need to call the API in each AWS Region.
func trustedadvisor_ListOrganizationRecommendationResources(cfg aws.Config, client *trustedadvisor.Client) {
	input := &trustedadvisor.ListOrganizationRecommendationResourcesInput{
		// OrganizationRecommendationIdentifier: *string, // Required
	}

	if len(_trustedadvisorOrganizationRecommendationIdentifier) > 0 {
		input.OrganizationRecommendationIdentifier = aws.String(_trustedadvisorOrganizationRecommendationIdentifier)
	}
	if len(_trustedadvisorAffectedAccountId) > 0 {
		input.AffectedAccountId = aws.String(_trustedadvisorAffectedAccountId)
	}
	if len(_trustedadvisorExclusionStatus) > 0 {
		if err := assignInputField(input, "ExclusionStatus", _trustedadvisorExclusionStatus); err != nil {
			log.Errorf("invalid --exclusion-status: %s", err.Error())
			return
		}
	}
	if len(_trustedadvisorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _trustedadvisorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_trustedadvisorNextToken) > 0 {
		input.NextToken = aws.String(_trustedadvisorNextToken)
	}
	if len(_trustedadvisorRegionCode) > 0 {
		input.RegionCode = aws.String(_trustedadvisorRegionCode)
	}
	if len(_trustedadvisorStatus) > 0 {
		if err := assignInputField(input, "Status", _trustedadvisorStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListOrganizationRecommendationResources(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*trustedadvisor.ListOrganizationRecommendationResourcesOutput
	p := trustedadvisor.NewListOrganizationRecommendationResourcesPaginator(client, input)
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

// List a filterable set of Recommendations within an Organization. This API only
// supports prioritized recommendations and provides global priority
// recommendations, eliminating the need to call the API in each AWS Region.
func trustedadvisor_ListOrganizationRecommendations(cfg aws.Config, client *trustedadvisor.Client) {
	input := &trustedadvisor.ListOrganizationRecommendationsInput{}

	if len(_trustedadvisorAfterLastUpdatedAt) > 0 {
		if err := assignInputField(input, "AfterLastUpdatedAt", _trustedadvisorAfterLastUpdatedAt); err != nil {
			log.Errorf("invalid --after-last-updated-at: %s", err.Error())
			return
		}
	}
	if len(_trustedadvisorAwsService) > 0 {
		input.AwsService = aws.String(_trustedadvisorAwsService)
	}
	if len(_trustedadvisorBeforeLastUpdatedAt) > 0 {
		if err := assignInputField(input, "BeforeLastUpdatedAt", _trustedadvisorBeforeLastUpdatedAt); err != nil {
			log.Errorf("invalid --before-last-updated-at: %s", err.Error())
			return
		}
	}
	if len(_trustedadvisorCheckIdentifier) > 0 {
		input.CheckIdentifier = aws.String(_trustedadvisorCheckIdentifier)
	}
	if len(_trustedadvisorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _trustedadvisorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_trustedadvisorNextToken) > 0 {
		input.NextToken = aws.String(_trustedadvisorNextToken)
	}
	if len(_trustedadvisorPillar) > 0 {
		if err := assignInputField(input, "Pillar", _trustedadvisorPillar); err != nil {
			log.Errorf("invalid --pillar: %s", err.Error())
			return
		}
	}
	if len(_trustedadvisorSource) > 0 {
		if err := assignInputField(input, "Source", _trustedadvisorSource); err != nil {
			log.Errorf("invalid --source: %s", err.Error())
			return
		}
	}
	if len(_trustedadvisorStatus) > 0 {
		if err := assignInputField(input, "Status", _trustedadvisorStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_trustedadvisorType) > 0 {
		if err := assignInputField(input, "Type", _trustedadvisorType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListOrganizationRecommendations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*trustedadvisor.ListOrganizationRecommendationsOutput
	p := trustedadvisor.NewListOrganizationRecommendationsPaginator(client, input)
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

// List Resources of a Recommendation. This API provides global recommendations,
// eliminating the need to call the API in each AWS Region.
func trustedadvisor_ListRecommendationResources(cfg aws.Config, client *trustedadvisor.Client) {
	input := &trustedadvisor.ListRecommendationResourcesInput{
		// RecommendationIdentifier: *string, // Required
	}

	if len(_trustedadvisorRecommendationIdentifier) > 0 {
		input.RecommendationIdentifier = aws.String(_trustedadvisorRecommendationIdentifier)
	}
	if len(_trustedadvisorExclusionStatus) > 0 {
		if err := assignInputField(input, "ExclusionStatus", _trustedadvisorExclusionStatus); err != nil {
			log.Errorf("invalid --exclusion-status: %s", err.Error())
			return
		}
	}
	if len(_trustedadvisorLanguage) > 0 {
		if err := assignInputField(input, "Language", _trustedadvisorLanguage); err != nil {
			log.Errorf("invalid --language: %s", err.Error())
			return
		}
	}
	if len(_trustedadvisorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _trustedadvisorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_trustedadvisorNextToken) > 0 {
		input.NextToken = aws.String(_trustedadvisorNextToken)
	}
	if len(_trustedadvisorRegionCode) > 0 {
		input.RegionCode = aws.String(_trustedadvisorRegionCode)
	}
	if len(_trustedadvisorStatus) > 0 {
		if err := assignInputField(input, "Status", _trustedadvisorStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListRecommendationResources(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*trustedadvisor.ListRecommendationResourcesOutput
	p := trustedadvisor.NewListRecommendationResourcesPaginator(client, input)
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

// List a filterable set of Recommendations. This API provides global
// recommendations, eliminating the need to call the API in each AWS Region.
func trustedadvisor_ListRecommendations(cfg aws.Config, client *trustedadvisor.Client) {
	input := &trustedadvisor.ListRecommendationsInput{}

	if len(_trustedadvisorAfterLastUpdatedAt) > 0 {
		if err := assignInputField(input, "AfterLastUpdatedAt", _trustedadvisorAfterLastUpdatedAt); err != nil {
			log.Errorf("invalid --after-last-updated-at: %s", err.Error())
			return
		}
	}
	if len(_trustedadvisorAwsService) > 0 {
		input.AwsService = aws.String(_trustedadvisorAwsService)
	}
	if len(_trustedadvisorBeforeLastUpdatedAt) > 0 {
		if err := assignInputField(input, "BeforeLastUpdatedAt", _trustedadvisorBeforeLastUpdatedAt); err != nil {
			log.Errorf("invalid --before-last-updated-at: %s", err.Error())
			return
		}
	}
	if len(_trustedadvisorCheckIdentifier) > 0 {
		input.CheckIdentifier = aws.String(_trustedadvisorCheckIdentifier)
	}
	if len(_trustedadvisorLanguage) > 0 {
		if err := assignInputField(input, "Language", _trustedadvisorLanguage); err != nil {
			log.Errorf("invalid --language: %s", err.Error())
			return
		}
	}
	if len(_trustedadvisorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _trustedadvisorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_trustedadvisorNextToken) > 0 {
		input.NextToken = aws.String(_trustedadvisorNextToken)
	}
	if len(_trustedadvisorPillar) > 0 {
		if err := assignInputField(input, "Pillar", _trustedadvisorPillar); err != nil {
			log.Errorf("invalid --pillar: %s", err.Error())
			return
		}
	}
	if len(_trustedadvisorSource) > 0 {
		if err := assignInputField(input, "Source", _trustedadvisorSource); err != nil {
			log.Errorf("invalid --source: %s", err.Error())
			return
		}
	}
	if len(_trustedadvisorStatus) > 0 {
		if err := assignInputField(input, "Status", _trustedadvisorStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_trustedadvisorType) > 0 {
		if err := assignInputField(input, "Type", _trustedadvisorType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
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

	var results []*trustedadvisor.ListRecommendationsOutput
	p := trustedadvisor.NewListRecommendationsPaginator(client, input)
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

// Update the lifecycle of a Recommendation within an Organization. This API only
// supports prioritized recommendations and updates global priority
// recommendations, eliminating the need to call the API in each AWS Region.
func trustedadvisor_UpdateOrganizationRecommendationLifecycle(cfg aws.Config, client *trustedadvisor.Client) {
	input := &trustedadvisor.UpdateOrganizationRecommendationLifecycleInput{
		// LifecycleStage: types.UpdateRecommendationLifecycleStage, // Required
		// OrganizationRecommendationIdentifier: *string, // Required
	}

	if len(_trustedadvisorLifecycleStage) > 0 {
		if err := assignInputField(input, "LifecycleStage", _trustedadvisorLifecycleStage); err != nil {
			log.Errorf("invalid --lifecycle-stage: %s", err.Error())
			return
		}
	}
	if len(_trustedadvisorOrganizationRecommendationIdentifier) > 0 {
		input.OrganizationRecommendationIdentifier = aws.String(_trustedadvisorOrganizationRecommendationIdentifier)
	}
	if len(_trustedadvisorUpdateReason) > 0 {
		input.UpdateReason = aws.String(_trustedadvisorUpdateReason)
	}
	if len(_trustedadvisorUpdateReasonCode) > 0 {
		if err := assignInputField(input, "UpdateReasonCode", _trustedadvisorUpdateReasonCode); err != nil {
			log.Errorf("invalid --update-reason-code: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateOrganizationRecommendationLifecycle(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update the lifecyle of a Recommendation. This API only supports prioritized
// recommendations and updates global priority recommendations, eliminating the
// need to call the API in each AWS Region.
func trustedadvisor_UpdateRecommendationLifecycle(cfg aws.Config, client *trustedadvisor.Client) {
	input := &trustedadvisor.UpdateRecommendationLifecycleInput{
		// LifecycleStage: types.UpdateRecommendationLifecycleStage, // Required
		// RecommendationIdentifier: *string, // Required
	}

	if len(_trustedadvisorLifecycleStage) > 0 {
		if err := assignInputField(input, "LifecycleStage", _trustedadvisorLifecycleStage); err != nil {
			log.Errorf("invalid --lifecycle-stage: %s", err.Error())
			return
		}
	}
	if len(_trustedadvisorRecommendationIdentifier) > 0 {
		input.RecommendationIdentifier = aws.String(_trustedadvisorRecommendationIdentifier)
	}
	if len(_trustedadvisorUpdateReason) > 0 {
		input.UpdateReason = aws.String(_trustedadvisorUpdateReason)
	}
	if len(_trustedadvisorUpdateReasonCode) > 0 {
		if err := assignInputField(input, "UpdateReasonCode", _trustedadvisorUpdateReasonCode); err != nil {
			log.Errorf("invalid --update-reason-code: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateRecommendationLifecycle(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_trustedadvisorCmd)
	_trustedadvisorCmd.Flags().SortFlags = false

	_trustedadvisorCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_trustedadvisorCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_trustedadvisorCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_trustedadvisorCmd.Flags().StringVarP(&_trustedadvisorAffectedAccountId, "affected-account-id", "", "", "Affected Account ID")
	_trustedadvisorCmd.Flags().StringVarP(&_trustedadvisorAfterLastUpdatedAt, "after-last-updated-at", "", "", "After Last Updated At")
	_trustedadvisorCmd.Flags().StringVarP(&_trustedadvisorAwsService, "aws-service", "", "", "AWS Service")
	_trustedadvisorCmd.Flags().StringVarP(&_trustedadvisorBeforeLastUpdatedAt, "before-last-updated-at", "", "", "Before Last Updated At")
	_trustedadvisorCmd.Flags().StringVarP(&_trustedadvisorCheckIdentifier, "check-identifier", "", "", "Check Identifier")
	_trustedadvisorCmd.Flags().StringVarP(&_trustedadvisorExclusionStatus, "exclusion-status", "", "", "Exclusion Status")
	_trustedadvisorCmd.Flags().StringVarP(&_trustedadvisorLanguage, "language", "", "", "Language")
	_trustedadvisorCmd.Flags().StringVarP(&_trustedadvisorLifecycleStage, "lifecycle-stage", "", "", "Lifecycle Stage")
	_trustedadvisorCmd.Flags().StringVarP(&_trustedadvisorMaxResults, "max-results", "", "", "Max Results")
	_trustedadvisorCmd.Flags().StringVarP(&_trustedadvisorNextToken, "next-token", "", "", "Next Token")
	_trustedadvisorCmd.Flags().StringVarP(&_trustedadvisorOrganizationRecommendationIdentifier, "organization-recommendation-identifier", "", "", "Organization Recommendation Identifier")
	_trustedadvisorCmd.Flags().StringVarP(&_trustedadvisorPillar, "pillar", "", "", "Pillar")
	_trustedadvisorCmd.Flags().StringVarP(&_trustedadvisorRecommendationIdentifier, "recommendation-identifier", "", "", "Recommendation Identifier")
	_trustedadvisorCmd.Flags().StringVarP(&_trustedadvisorRecommendationResourceExclusions, "recommendation-resource-exclusions", "", "", "Recommendation Resource Exclusions")
	_trustedadvisorCmd.Flags().StringVarP(&_trustedadvisorRegionCode, "region-code", "", "", "Region Code")
	_trustedadvisorCmd.Flags().StringVarP(&_trustedadvisorSource, "source", "", "", "Source")
	_trustedadvisorCmd.Flags().StringVarP(&_trustedadvisorStatus, "status", "", "", "Status")
	_trustedadvisorCmd.Flags().StringVarP(&_trustedadvisorType, "type", "", "", "Type")
	_trustedadvisorCmd.Flags().StringVarP(&_trustedadvisorUpdateReason, "update-reason", "", "", "Update Reason")
	_trustedadvisorCmd.Flags().StringVarP(&_trustedadvisorUpdateReasonCode, "update-reason-code", "", "", "Update Reason Code")

	_trustedadvisorCmd.Flags().BoolVarP(&_trustedadvisorBatchUpdateRecommendationResourceExclusion, "batch-update-recommendation-resource-exclusion", "", false, "Batch Update Recommendation Resource Exclusion")
	_trustedadvisorCmd.Flags().BoolVarP(&_trustedadvisorGetOrganizationRecommendation, "get-organization-recommendation", "", false, "Get Organization Recommendation")
	_trustedadvisorCmd.Flags().BoolVarP(&_trustedadvisorGetRecommendation, "get-recommendation", "", false, "Get Recommendation")
	_trustedadvisorCmd.Flags().BoolVarP(&_trustedadvisorListChecks, "list-checks", "", false, "List Checks")
	_trustedadvisorCmd.Flags().BoolVarP(&_trustedadvisorListOrganizationRecommendationAccounts, "list-organization-recommendation-accounts", "", false, "List Organization Recommendation Accounts")
	_trustedadvisorCmd.Flags().BoolVarP(&_trustedadvisorListOrganizationRecommendationResources, "list-organization-recommendation-resources", "", false, "List Organization Recommendation Resources")
	_trustedadvisorCmd.Flags().BoolVarP(&_trustedadvisorListOrganizationRecommendations, "list-organization-recommendations", "", false, "List Organization Recommendations")
	_trustedadvisorCmd.Flags().BoolVarP(&_trustedadvisorListRecommendationResources, "list-recommendation-resources", "", false, "List Recommendation Resources")
	_trustedadvisorCmd.Flags().BoolVarP(&_trustedadvisorListRecommendations, "list-recommendations", "", false, "List Recommendations")
	_trustedadvisorCmd.Flags().BoolVarP(&_trustedadvisorUpdateOrganizationRecommendationLifecycle, "update-organization-recommendation-lifecycle", "", false, "Update Organization Recommendation Lifecycle")
	_trustedadvisorCmd.Flags().BoolVarP(&_trustedadvisorUpdateRecommendationLifecycle, "update-recommendation-lifecycle", "", false, "Update Recommendation Lifecycle")

}
