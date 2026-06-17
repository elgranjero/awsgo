package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/computeoptimizerautomation"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// computeoptimizerautomationCmd represents the computeoptimizerautomation command
var _computeoptimizerautomationCmd = &cobra.Command{
	Use:   "computeoptimizerautomation",
	Short: "AWS computeoptimizerautomation CLI",
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
		client := computeoptimizerautomation.NewFromConfig(cfg)
		if _computeoptimizerautomationAssociateAccounts {
			computeoptimizerautomation_AssociateAccounts(cfg, client)
			return
		}
		if _computeoptimizerautomationCreateAutomationRule {
			computeoptimizerautomation_CreateAutomationRule(cfg, client)
			return
		}
		if _computeoptimizerautomationDeleteAutomationRule {
			computeoptimizerautomation_DeleteAutomationRule(cfg, client)
			return
		}
		if _computeoptimizerautomationDisassociateAccounts {
			computeoptimizerautomation_DisassociateAccounts(cfg, client)
			return
		}
		if _computeoptimizerautomationGetAutomationEvent {
			computeoptimizerautomation_GetAutomationEvent(cfg, client)
			return
		}
		if _computeoptimizerautomationGetAutomationRule {
			computeoptimizerautomation_GetAutomationRule(cfg, client)
			return
		}
		if _computeoptimizerautomationGetEnrollmentConfiguration {
			computeoptimizerautomation_GetEnrollmentConfiguration(cfg, client)
			return
		}
		if _computeoptimizerautomationListAccounts {
			computeoptimizerautomation_ListAccounts(cfg, client)
			return
		}
		if _computeoptimizerautomationListAutomationEventSteps {
			computeoptimizerautomation_ListAutomationEventSteps(cfg, client)
			return
		}
		if _computeoptimizerautomationListAutomationEventSummaries {
			computeoptimizerautomation_ListAutomationEventSummaries(cfg, client)
			return
		}
		if _computeoptimizerautomationListAutomationEvents {
			computeoptimizerautomation_ListAutomationEvents(cfg, client)
			return
		}
		if _computeoptimizerautomationListAutomationRulePreview {
			computeoptimizerautomation_ListAutomationRulePreview(cfg, client)
			return
		}
		if _computeoptimizerautomationListAutomationRulePreviewSummaries {
			computeoptimizerautomation_ListAutomationRulePreviewSummaries(cfg, client)
			return
		}
		if _computeoptimizerautomationListAutomationRules {
			computeoptimizerautomation_ListAutomationRules(cfg, client)
			return
		}
		if _computeoptimizerautomationListRecommendedActionSummaries {
			computeoptimizerautomation_ListRecommendedActionSummaries(cfg, client)
			return
		}
		if _computeoptimizerautomationListRecommendedActions {
			computeoptimizerautomation_ListRecommendedActions(cfg, client)
			return
		}
		if _computeoptimizerautomationListTagsForResource {
			computeoptimizerautomation_ListTagsForResource(cfg, client)
			return
		}
		if _computeoptimizerautomationRollbackAutomationEvent {
			computeoptimizerautomation_RollbackAutomationEvent(cfg, client)
			return
		}
		if _computeoptimizerautomationStartAutomationEvent {
			computeoptimizerautomation_StartAutomationEvent(cfg, client)
			return
		}
		if _computeoptimizerautomationTagResource {
			computeoptimizerautomation_TagResource(cfg, client)
			return
		}
		if _computeoptimizerautomationUntagResource {
			computeoptimizerautomation_UntagResource(cfg, client)
			return
		}
		if _computeoptimizerautomationUpdateAutomationRule {
			computeoptimizerautomation_UpdateAutomationRule(cfg, client)
			return
		}
		if _computeoptimizerautomationUpdateEnrollmentConfiguration {
			computeoptimizerautomation_UpdateEnrollmentConfiguration(cfg, client)
			return
		}

	},
}

var (
	_computeoptimizerautomationAssociateAccounts                  bool
	_computeoptimizerautomationCreateAutomationRule               bool
	_computeoptimizerautomationDeleteAutomationRule               bool
	_computeoptimizerautomationDisassociateAccounts               bool
	_computeoptimizerautomationGetAutomationEvent                 bool
	_computeoptimizerautomationGetAutomationRule                  bool
	_computeoptimizerautomationGetEnrollmentConfiguration         bool
	_computeoptimizerautomationListAccounts                       bool
	_computeoptimizerautomationListAutomationEventSteps           bool
	_computeoptimizerautomationListAutomationEventSummaries       bool
	_computeoptimizerautomationListAutomationEvents               bool
	_computeoptimizerautomationListAutomationRulePreview          bool
	_computeoptimizerautomationListAutomationRulePreviewSummaries bool
	_computeoptimizerautomationListAutomationRules                bool
	_computeoptimizerautomationListRecommendedActionSummaries     bool
	_computeoptimizerautomationListRecommendedActions             bool
	_computeoptimizerautomationListTagsForResource                bool
	_computeoptimizerautomationRollbackAutomationEvent            bool
	_computeoptimizerautomationStartAutomationEvent               bool
	_computeoptimizerautomationTagResource                        bool
	_computeoptimizerautomationUntagResource                      bool
	_computeoptimizerautomationUpdateAutomationRule               bool
	_computeoptimizerautomationUpdateEnrollmentConfiguration      bool

	_computeoptimizerautomationAccountIds                []string
	_computeoptimizerautomationClientToken               string
	_computeoptimizerautomationCriteria                  string
	_computeoptimizerautomationDescription               string
	_computeoptimizerautomationEndDateExclusive          string
	_computeoptimizerautomationEndTimeExclusive          string
	_computeoptimizerautomationEventId                   string
	_computeoptimizerautomationFilters                   string
	_computeoptimizerautomationMaxResults                string
	_computeoptimizerautomationName                      string
	_computeoptimizerautomationNextToken                 string
	_computeoptimizerautomationOrganizationConfiguration string
	_computeoptimizerautomationOrganizationScope         string
	_computeoptimizerautomationPriority                  string
	_computeoptimizerautomationRecommendedActionId       string
	_computeoptimizerautomationRecommendedActionTypes    string
	_computeoptimizerautomationResourceArn               string
	_computeoptimizerautomationRuleArn                   string
	_computeoptimizerautomationRuleRevision              string
	_computeoptimizerautomationRuleType                  string
	_computeoptimizerautomationSchedule                  string
	_computeoptimizerautomationStartDateInclusive        string
	_computeoptimizerautomationStartTimeInclusive        string
	_computeoptimizerautomationStatus                    string
	_computeoptimizerautomationTagKeys                   []string
	_computeoptimizerautomationTags                      string
)

// Associates one or more member accounts with your organization's management
// account, enabling centralized implementation of optimization actions across
// those accounts. Once associated, the management account (or a delegated
// administrator) can apply recommended actions to the member account. When you
// associate a member account, its organization rule mode is automatically set to
// "Any allowed," which permits the management account to create Automation rules
// that automatically apply actions to that account. If the member account has not
// previously enabled the Automation feature, the association process automatically
// enables it.
//
// Only the management account or a delegated administrator can perform this
// action.
func computeoptimizerautomation_AssociateAccounts(cfg aws.Config, client *computeoptimizerautomation.Client) {
	input := &computeoptimizerautomation.AssociateAccountsInput{
		// AccountIds: []string, // Required
	}

	if len(_computeoptimizerautomationAccountIds) > 0 {
		input.AccountIds = append([]string(nil), _computeoptimizerautomationAccountIds...)
	}
	if len(_computeoptimizerautomationClientToken) > 0 {
		input.ClientToken = aws.String(_computeoptimizerautomationClientToken)
	}

	if resp, err := client.AssociateAccounts(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new automation rule to apply recommended actions to resources based
// on specified criteria.
func computeoptimizerautomation_CreateAutomationRule(cfg aws.Config, client *computeoptimizerautomation.Client) {
	input := &computeoptimizerautomation.CreateAutomationRuleInput{
		// Name: *string, // Required
		// RecommendedActionTypes: []types.RecommendedActionType, // Required
		// RuleType: types.RuleType, // Required
		// Schedule: *types.Schedule, // Required
		// Status: types.RuleStatus, // Required
	}

	if len(_computeoptimizerautomationName) > 0 {
		input.Name = aws.String(_computeoptimizerautomationName)
	}
	if len(_computeoptimizerautomationRecommendedActionTypes) > 0 {
		if err := assignInputField(input, "RecommendedActionTypes", _computeoptimizerautomationRecommendedActionTypes); err != nil {
			log.Errorf("invalid --recommended-action-types: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerautomationRuleType) > 0 {
		if err := assignInputField(input, "RuleType", _computeoptimizerautomationRuleType); err != nil {
			log.Errorf("invalid --rule-type: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerautomationSchedule) > 0 {
		if err := assignInputField(input, "Schedule", _computeoptimizerautomationSchedule); err != nil {
			log.Errorf("invalid --schedule: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerautomationStatus) > 0 {
		if err := assignInputField(input, "Status", _computeoptimizerautomationStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerautomationClientToken) > 0 {
		input.ClientToken = aws.String(_computeoptimizerautomationClientToken)
	}
	if len(_computeoptimizerautomationCriteria) > 0 {
		if err := assignInputField(input, "Criteria", _computeoptimizerautomationCriteria); err != nil {
			log.Errorf("invalid --criteria: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerautomationDescription) > 0 {
		input.Description = aws.String(_computeoptimizerautomationDescription)
	}
	if len(_computeoptimizerautomationOrganizationConfiguration) > 0 {
		if err := assignInputField(input, "OrganizationConfiguration", _computeoptimizerautomationOrganizationConfiguration); err != nil {
			log.Errorf("invalid --organization-configuration: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerautomationPriority) > 0 {
		input.Priority = aws.String(_computeoptimizerautomationPriority)
	}
	if len(_computeoptimizerautomationTags) > 0 {
		if err := assignInputField(input, "Tags", _computeoptimizerautomationTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAutomationRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing automation rule.
func computeoptimizerautomation_DeleteAutomationRule(cfg aws.Config, client *computeoptimizerautomation.Client) {
	input := &computeoptimizerautomation.DeleteAutomationRuleInput{
		// RuleArn: *string, // Required
		// RuleRevision: *int64, // Required
	}

	if len(_computeoptimizerautomationRuleArn) > 0 {
		input.RuleArn = aws.String(_computeoptimizerautomationRuleArn)
	}
	if len(_computeoptimizerautomationRuleRevision) > 0 {
		if err := assignInputField(input, "RuleRevision", _computeoptimizerautomationRuleRevision); err != nil {
			log.Errorf("invalid --rule-revision: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerautomationClientToken) > 0 {
		input.ClientToken = aws.String(_computeoptimizerautomationClientToken)
	}

	if resp, err := client.DeleteAutomationRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates member accounts from your organization's management account,
// removing centralized automation capabilities. Once disassociated, organization
// rules no longer apply to the member account, and the management account (or
// delegated administrator) cannot create Automation rules for that account.
//
// Only the management account or a delegated administrator can perform this
// action.
func computeoptimizerautomation_DisassociateAccounts(cfg aws.Config, client *computeoptimizerautomation.Client) {
	input := &computeoptimizerautomation.DisassociateAccountsInput{
		// AccountIds: []string, // Required
	}

	if len(_computeoptimizerautomationAccountIds) > 0 {
		input.AccountIds = append([]string(nil), _computeoptimizerautomationAccountIds...)
	}
	if len(_computeoptimizerautomationClientToken) > 0 {
		input.ClientToken = aws.String(_computeoptimizerautomationClientToken)
	}

	if resp, err := client.DisassociateAccounts(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves details about a specific automation event.
func computeoptimizerautomation_GetAutomationEvent(cfg aws.Config, client *computeoptimizerautomation.Client) {
	input := &computeoptimizerautomation.GetAutomationEventInput{
		// EventId: *string, // Required
	}

	if len(_computeoptimizerautomationEventId) > 0 {
		input.EventId = aws.String(_computeoptimizerautomationEventId)
	}

	if resp, err := client.GetAutomationEvent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves details about a specific automation rule.
func computeoptimizerautomation_GetAutomationRule(cfg aws.Config, client *computeoptimizerautomation.Client) {
	input := &computeoptimizerautomation.GetAutomationRuleInput{
		// RuleArn: *string, // Required
	}

	if len(_computeoptimizerautomationRuleArn) > 0 {
		input.RuleArn = aws.String(_computeoptimizerautomationRuleArn)
	}

	if resp, err := client.GetAutomationRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the current enrollment configuration for Compute Optimizer
// Automation.
func computeoptimizerautomation_GetEnrollmentConfiguration(cfg aws.Config, client *computeoptimizerautomation.Client) {
	input := &computeoptimizerautomation.GetEnrollmentConfigurationInput{}

	if resp, err := client.GetEnrollmentConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the accounts in your organization that are enrolled in Compute Optimizer
// and whether they have enabled Automation.
//
// Only the management account or a delegated administrator can perform this
// action.
func computeoptimizerautomation_ListAccounts(cfg aws.Config, client *computeoptimizerautomation.Client) {
	input := &computeoptimizerautomation.ListAccountsInput{}

	if len(_computeoptimizerautomationMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _computeoptimizerautomationMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerautomationNextToken) > 0 {
		input.NextToken = aws.String(_computeoptimizerautomationNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAccounts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*computeoptimizerautomation.ListAccountsOutput
	p := computeoptimizerautomation.NewListAccountsPaginator(client, input)
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

// Lists the steps for a specific automation event. You can only list steps for
// events created within the past year.
func computeoptimizerautomation_ListAutomationEventSteps(cfg aws.Config, client *computeoptimizerautomation.Client) {
	input := &computeoptimizerautomation.ListAutomationEventStepsInput{
		// EventId: *string, // Required
	}

	if len(_computeoptimizerautomationEventId) > 0 {
		input.EventId = aws.String(_computeoptimizerautomationEventId)
	}
	if len(_computeoptimizerautomationMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _computeoptimizerautomationMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerautomationNextToken) > 0 {
		input.NextToken = aws.String(_computeoptimizerautomationNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAutomationEventSteps(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*computeoptimizerautomation.ListAutomationEventStepsOutput
	p := computeoptimizerautomation.NewListAutomationEventStepsPaginator(client, input)
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

// Provides a summary of automation events based on specified filters. Only events
// created within the past year will be included in the summary.
func computeoptimizerautomation_ListAutomationEventSummaries(cfg aws.Config, client *computeoptimizerautomation.Client) {
	input := &computeoptimizerautomation.ListAutomationEventSummariesInput{}

	if len(_computeoptimizerautomationEndDateExclusive) > 0 {
		input.EndDateExclusive = aws.String(_computeoptimizerautomationEndDateExclusive)
	}
	if len(_computeoptimizerautomationFilters) > 0 {
		if err := assignInputField(input, "Filters", _computeoptimizerautomationFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerautomationMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _computeoptimizerautomationMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerautomationNextToken) > 0 {
		input.NextToken = aws.String(_computeoptimizerautomationNextToken)
	}
	if len(_computeoptimizerautomationStartDateInclusive) > 0 {
		input.StartDateInclusive = aws.String(_computeoptimizerautomationStartDateInclusive)
	}

	if disablePaginator() {
		if resp, err := client.ListAutomationEventSummaries(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*computeoptimizerautomation.ListAutomationEventSummariesOutput
	p := computeoptimizerautomation.NewListAutomationEventSummariesPaginator(client, input)
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

// Lists automation events based on specified filters. You can retrieve events
// that were created within the past year.
func computeoptimizerautomation_ListAutomationEvents(cfg aws.Config, client *computeoptimizerautomation.Client) {
	input := &computeoptimizerautomation.ListAutomationEventsInput{}

	if len(_computeoptimizerautomationEndTimeExclusive) > 0 {
		if err := assignInputField(input, "EndTimeExclusive", _computeoptimizerautomationEndTimeExclusive); err != nil {
			log.Errorf("invalid --end-time-exclusive: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerautomationFilters) > 0 {
		if err := assignInputField(input, "Filters", _computeoptimizerautomationFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerautomationMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _computeoptimizerautomationMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerautomationNextToken) > 0 {
		input.NextToken = aws.String(_computeoptimizerautomationNextToken)
	}
	if len(_computeoptimizerautomationStartTimeInclusive) > 0 {
		if err := assignInputField(input, "StartTimeInclusive", _computeoptimizerautomationStartTimeInclusive); err != nil {
			log.Errorf("invalid --start-time-inclusive: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListAutomationEvents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*computeoptimizerautomation.ListAutomationEventsOutput
	p := computeoptimizerautomation.NewListAutomationEventsPaginator(client, input)
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

// Returns a preview of the recommended actions that match your Automation rule's
// configuration and criteria.
func computeoptimizerautomation_ListAutomationRulePreview(cfg aws.Config, client *computeoptimizerautomation.Client) {
	input := &computeoptimizerautomation.ListAutomationRulePreviewInput{
		// RecommendedActionTypes: []types.RecommendedActionType, // Required
		// RuleType: types.RuleType, // Required
	}

	if len(_computeoptimizerautomationRecommendedActionTypes) > 0 {
		if err := assignInputField(input, "RecommendedActionTypes", _computeoptimizerautomationRecommendedActionTypes); err != nil {
			log.Errorf("invalid --recommended-action-types: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerautomationRuleType) > 0 {
		if err := assignInputField(input, "RuleType", _computeoptimizerautomationRuleType); err != nil {
			log.Errorf("invalid --rule-type: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerautomationCriteria) > 0 {
		if err := assignInputField(input, "Criteria", _computeoptimizerautomationCriteria); err != nil {
			log.Errorf("invalid --criteria: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerautomationMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _computeoptimizerautomationMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerautomationNextToken) > 0 {
		input.NextToken = aws.String(_computeoptimizerautomationNextToken)
	}
	if len(_computeoptimizerautomationOrganizationScope) > 0 {
		if err := assignInputField(input, "OrganizationScope", _computeoptimizerautomationOrganizationScope); err != nil {
			log.Errorf("invalid --organization-scope: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListAutomationRulePreview(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*computeoptimizerautomation.ListAutomationRulePreviewOutput
	p := computeoptimizerautomation.NewListAutomationRulePreviewPaginator(client, input)
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

// Returns a summary of the recommended actions that match your rule preview
// configuration and criteria.
func computeoptimizerautomation_ListAutomationRulePreviewSummaries(cfg aws.Config, client *computeoptimizerautomation.Client) {
	input := &computeoptimizerautomation.ListAutomationRulePreviewSummariesInput{
		// RecommendedActionTypes: []types.RecommendedActionType, // Required
		// RuleType: types.RuleType, // Required
	}

	if len(_computeoptimizerautomationRecommendedActionTypes) > 0 {
		if err := assignInputField(input, "RecommendedActionTypes", _computeoptimizerautomationRecommendedActionTypes); err != nil {
			log.Errorf("invalid --recommended-action-types: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerautomationRuleType) > 0 {
		if err := assignInputField(input, "RuleType", _computeoptimizerautomationRuleType); err != nil {
			log.Errorf("invalid --rule-type: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerautomationCriteria) > 0 {
		if err := assignInputField(input, "Criteria", _computeoptimizerautomationCriteria); err != nil {
			log.Errorf("invalid --criteria: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerautomationMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _computeoptimizerautomationMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerautomationNextToken) > 0 {
		input.NextToken = aws.String(_computeoptimizerautomationNextToken)
	}
	if len(_computeoptimizerautomationOrganizationScope) > 0 {
		if err := assignInputField(input, "OrganizationScope", _computeoptimizerautomationOrganizationScope); err != nil {
			log.Errorf("invalid --organization-scope: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListAutomationRulePreviewSummaries(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*computeoptimizerautomation.ListAutomationRulePreviewSummariesOutput
	p := computeoptimizerautomation.NewListAutomationRulePreviewSummariesPaginator(client, input)
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

// Lists the automation rules that match specified filters.
func computeoptimizerautomation_ListAutomationRules(cfg aws.Config, client *computeoptimizerautomation.Client) {
	input := &computeoptimizerautomation.ListAutomationRulesInput{}

	if len(_computeoptimizerautomationFilters) > 0 {
		if err := assignInputField(input, "Filters", _computeoptimizerautomationFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerautomationMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _computeoptimizerautomationMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerautomationNextToken) > 0 {
		input.NextToken = aws.String(_computeoptimizerautomationNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAutomationRules(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*computeoptimizerautomation.ListAutomationRulesOutput
	p := computeoptimizerautomation.NewListAutomationRulesPaginator(client, input)
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

// Provides a summary of recommended actions based on specified filters.
// Management accounts and delegated administrators can retrieve recommended
// actions that include associated member accounts. You can associate a member
// account using AssociateAccounts .
func computeoptimizerautomation_ListRecommendedActionSummaries(cfg aws.Config, client *computeoptimizerautomation.Client) {
	input := &computeoptimizerautomation.ListRecommendedActionSummariesInput{}

	if len(_computeoptimizerautomationFilters) > 0 {
		if err := assignInputField(input, "Filters", _computeoptimizerautomationFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerautomationMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _computeoptimizerautomationMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerautomationNextToken) > 0 {
		input.NextToken = aws.String(_computeoptimizerautomationNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRecommendedActionSummaries(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*computeoptimizerautomation.ListRecommendedActionSummariesOutput
	p := computeoptimizerautomation.NewListRecommendedActionSummariesPaginator(client, input)
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

// Lists the recommended actions based that match specified filters.
// Management accounts and delegated administrators can retrieve recommended
// actions that include associated member accounts. You can associate a member
// account using AssociateAccounts .
func computeoptimizerautomation_ListRecommendedActions(cfg aws.Config, client *computeoptimizerautomation.Client) {
	input := &computeoptimizerautomation.ListRecommendedActionsInput{}

	if len(_computeoptimizerautomationFilters) > 0 {
		if err := assignInputField(input, "Filters", _computeoptimizerautomationFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerautomationMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _computeoptimizerautomationMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerautomationNextToken) > 0 {
		input.NextToken = aws.String(_computeoptimizerautomationNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRecommendedActions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*computeoptimizerautomation.ListRecommendedActionsOutput
	p := computeoptimizerautomation.NewListRecommendedActionsPaginator(client, input)
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

// Lists the tags for a specified resource.
func computeoptimizerautomation_ListTagsForResource(cfg aws.Config, client *computeoptimizerautomation.Client) {
	input := &computeoptimizerautomation.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_computeoptimizerautomationResourceArn) > 0 {
		input.ResourceArn = aws.String(_computeoptimizerautomationResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Initiates a rollback for a completed automation event.
// Management accounts and delegated administrators can only initiate a rollback
// for events belonging to associated member accounts. You can associate a member
// account using AssociateAccounts .
func computeoptimizerautomation_RollbackAutomationEvent(cfg aws.Config, client *computeoptimizerautomation.Client) {
	input := &computeoptimizerautomation.RollbackAutomationEventInput{
		// EventId: *string, // Required
	}

	if len(_computeoptimizerautomationEventId) > 0 {
		input.EventId = aws.String(_computeoptimizerautomationEventId)
	}
	if len(_computeoptimizerautomationClientToken) > 0 {
		input.ClientToken = aws.String(_computeoptimizerautomationClientToken)
	}

	if resp, err := client.RollbackAutomationEvent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Initiates a one-time, on-demand automation for the specified recommended
// action.
//
// Management accounts and delegated administrators can only initiate recommended
// actions for associated member accounts. You can associate a member account using
// AssociateAccounts .
func computeoptimizerautomation_StartAutomationEvent(cfg aws.Config, client *computeoptimizerautomation.Client) {
	input := &computeoptimizerautomation.StartAutomationEventInput{
		// RecommendedActionId: *string, // Required
	}

	if len(_computeoptimizerautomationRecommendedActionId) > 0 {
		input.RecommendedActionId = aws.String(_computeoptimizerautomationRecommendedActionId)
	}
	if len(_computeoptimizerautomationClientToken) > 0 {
		input.ClientToken = aws.String(_computeoptimizerautomationClientToken)
	}

	if resp, err := client.StartAutomationEvent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds tags to the specified resource.
func computeoptimizerautomation_TagResource(cfg aws.Config, client *computeoptimizerautomation.Client) {
	input := &computeoptimizerautomation.TagResourceInput{
		// ResourceArn: *string, // Required
		// RuleRevision: *int64, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_computeoptimizerautomationResourceArn) > 0 {
		input.ResourceArn = aws.String(_computeoptimizerautomationResourceArn)
	}
	if len(_computeoptimizerautomationRuleRevision) > 0 {
		if err := assignInputField(input, "RuleRevision", _computeoptimizerautomationRuleRevision); err != nil {
			log.Errorf("invalid --rule-revision: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerautomationTags) > 0 {
		if err := assignInputField(input, "Tags", _computeoptimizerautomationTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerautomationClientToken) > 0 {
		input.ClientToken = aws.String(_computeoptimizerautomationClientToken)
	}

	if resp, err := client.TagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes tags from the specified resource.
func computeoptimizerautomation_UntagResource(cfg aws.Config, client *computeoptimizerautomation.Client) {
	input := &computeoptimizerautomation.UntagResourceInput{
		// ResourceArn: *string, // Required
		// RuleRevision: *int64, // Required
		// TagKeys: []string, // Required
	}

	if len(_computeoptimizerautomationResourceArn) > 0 {
		input.ResourceArn = aws.String(_computeoptimizerautomationResourceArn)
	}
	if len(_computeoptimizerautomationRuleRevision) > 0 {
		if err := assignInputField(input, "RuleRevision", _computeoptimizerautomationRuleRevision); err != nil {
			log.Errorf("invalid --rule-revision: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerautomationTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _computeoptimizerautomationTagKeys...)
	}
	if len(_computeoptimizerautomationClientToken) > 0 {
		input.ClientToken = aws.String(_computeoptimizerautomationClientToken)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing automation rule.
func computeoptimizerautomation_UpdateAutomationRule(cfg aws.Config, client *computeoptimizerautomation.Client) {
	input := &computeoptimizerautomation.UpdateAutomationRuleInput{
		// RuleArn: *string, // Required
		// RuleRevision: *int64, // Required
	}

	if len(_computeoptimizerautomationRuleArn) > 0 {
		input.RuleArn = aws.String(_computeoptimizerautomationRuleArn)
	}
	if len(_computeoptimizerautomationRuleRevision) > 0 {
		if err := assignInputField(input, "RuleRevision", _computeoptimizerautomationRuleRevision); err != nil {
			log.Errorf("invalid --rule-revision: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerautomationClientToken) > 0 {
		input.ClientToken = aws.String(_computeoptimizerautomationClientToken)
	}
	if len(_computeoptimizerautomationCriteria) > 0 {
		if err := assignInputField(input, "Criteria", _computeoptimizerautomationCriteria); err != nil {
			log.Errorf("invalid --criteria: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerautomationDescription) > 0 {
		input.Description = aws.String(_computeoptimizerautomationDescription)
	}
	if len(_computeoptimizerautomationName) > 0 {
		input.Name = aws.String(_computeoptimizerautomationName)
	}
	if len(_computeoptimizerautomationOrganizationConfiguration) > 0 {
		if err := assignInputField(input, "OrganizationConfiguration", _computeoptimizerautomationOrganizationConfiguration); err != nil {
			log.Errorf("invalid --organization-configuration: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerautomationPriority) > 0 {
		input.Priority = aws.String(_computeoptimizerautomationPriority)
	}
	if len(_computeoptimizerautomationRecommendedActionTypes) > 0 {
		if err := assignInputField(input, "RecommendedActionTypes", _computeoptimizerautomationRecommendedActionTypes); err != nil {
			log.Errorf("invalid --recommended-action-types: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerautomationRuleType) > 0 {
		if err := assignInputField(input, "RuleType", _computeoptimizerautomationRuleType); err != nil {
			log.Errorf("invalid --rule-type: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerautomationSchedule) > 0 {
		if err := assignInputField(input, "Schedule", _computeoptimizerautomationSchedule); err != nil {
			log.Errorf("invalid --schedule: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerautomationStatus) > 0 {
		if err := assignInputField(input, "Status", _computeoptimizerautomationStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateAutomationRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates your account’s Compute Optimizer Automation enrollment configuration.
func computeoptimizerautomation_UpdateEnrollmentConfiguration(cfg aws.Config, client *computeoptimizerautomation.Client) {
	input := &computeoptimizerautomation.UpdateEnrollmentConfigurationInput{
		// Status: types.EnrollmentStatus, // Required
	}

	if len(_computeoptimizerautomationStatus) > 0 {
		if err := assignInputField(input, "Status", _computeoptimizerautomationStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_computeoptimizerautomationClientToken) > 0 {
		input.ClientToken = aws.String(_computeoptimizerautomationClientToken)
	}

	if resp, err := client.UpdateEnrollmentConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_computeoptimizerautomationCmd)
	_computeoptimizerautomationCmd.Flags().SortFlags = false

	_computeoptimizerautomationCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_computeoptimizerautomationCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_computeoptimizerautomationCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_computeoptimizerautomationCmd.Flags().StringSliceVarP(&_computeoptimizerautomationAccountIds, "account-ids", "", nil, "Account Ids")
	_computeoptimizerautomationCmd.Flags().StringVarP(&_computeoptimizerautomationClientToken, "client-token", "", "", "Client Token")
	_computeoptimizerautomationCmd.Flags().StringVarP(&_computeoptimizerautomationCriteria, "criteria", "", "", "Criteria")
	_computeoptimizerautomationCmd.Flags().StringVarP(&_computeoptimizerautomationDescription, "description", "", "", "Description")
	_computeoptimizerautomationCmd.Flags().StringVarP(&_computeoptimizerautomationEndDateExclusive, "end-date-exclusive", "", "", "End Date Exclusive")
	_computeoptimizerautomationCmd.Flags().StringVarP(&_computeoptimizerautomationEndTimeExclusive, "end-time-exclusive", "", "", "End Time Exclusive")
	_computeoptimizerautomationCmd.Flags().StringVarP(&_computeoptimizerautomationEventId, "event-id", "", "", "Event ID")
	_computeoptimizerautomationCmd.Flags().StringVarP(&_computeoptimizerautomationFilters, "filters", "", "", "Filters")
	_computeoptimizerautomationCmd.Flags().StringVarP(&_computeoptimizerautomationMaxResults, "max-results", "", "", "Max Results")
	_computeoptimizerautomationCmd.Flags().StringVarP(&_computeoptimizerautomationName, "name", "", "", "Name")
	_computeoptimizerautomationCmd.Flags().StringVarP(&_computeoptimizerautomationNextToken, "next-token", "", "", "Next Token")
	_computeoptimizerautomationCmd.Flags().StringVarP(&_computeoptimizerautomationOrganizationConfiguration, "organization-configuration", "", "", "Organization Configuration")
	_computeoptimizerautomationCmd.Flags().StringVarP(&_computeoptimizerautomationOrganizationScope, "organization-scope", "", "", "Organization Scope")
	_computeoptimizerautomationCmd.Flags().StringVarP(&_computeoptimizerautomationPriority, "priority", "", "", "Priority")
	_computeoptimizerautomationCmd.Flags().StringVarP(&_computeoptimizerautomationRecommendedActionId, "recommended-action-id", "", "", "Recommended Action ID")
	_computeoptimizerautomationCmd.Flags().StringVarP(&_computeoptimizerautomationRecommendedActionTypes, "recommended-action-types", "", "", "Recommended Action Types")
	_computeoptimizerautomationCmd.Flags().StringVarP(&_computeoptimizerautomationResourceArn, "resource-arn", "", "", "Resource ARN")
	_computeoptimizerautomationCmd.Flags().StringVarP(&_computeoptimizerautomationRuleArn, "rule-arn", "", "", "Rule ARN")
	_computeoptimizerautomationCmd.Flags().StringVarP(&_computeoptimizerautomationRuleRevision, "rule-revision", "", "", "Rule Revision")
	_computeoptimizerautomationCmd.Flags().StringVarP(&_computeoptimizerautomationRuleType, "rule-type", "", "", "Rule Type")
	_computeoptimizerautomationCmd.Flags().StringVarP(&_computeoptimizerautomationSchedule, "schedule", "", "", "Schedule")
	_computeoptimizerautomationCmd.Flags().StringVarP(&_computeoptimizerautomationStartDateInclusive, "start-date-inclusive", "", "", "Start Date Inclusive")
	_computeoptimizerautomationCmd.Flags().StringVarP(&_computeoptimizerautomationStartTimeInclusive, "start-time-inclusive", "", "", "Start Time Inclusive")
	_computeoptimizerautomationCmd.Flags().StringVarP(&_computeoptimizerautomationStatus, "status", "", "", "Status")
	_computeoptimizerautomationCmd.Flags().StringSliceVarP(&_computeoptimizerautomationTagKeys, "tag-keys", "", nil, "Tag Keys")
	_computeoptimizerautomationCmd.Flags().StringVarP(&_computeoptimizerautomationTags, "tags", "", "", "Tags")

	_computeoptimizerautomationCmd.Flags().BoolVarP(&_computeoptimizerautomationAssociateAccounts, "associate-accounts", "", false, "Associate Accounts")
	_computeoptimizerautomationCmd.Flags().BoolVarP(&_computeoptimizerautomationCreateAutomationRule, "create-automation-rule", "", false, "Create Automation Rule")
	_computeoptimizerautomationCmd.Flags().BoolVarP(&_computeoptimizerautomationDeleteAutomationRule, "delete-automation-rule", "", false, "Delete Automation Rule")
	_computeoptimizerautomationCmd.Flags().BoolVarP(&_computeoptimizerautomationDisassociateAccounts, "disassociate-accounts", "", false, "Disassociate Accounts")
	_computeoptimizerautomationCmd.Flags().BoolVarP(&_computeoptimizerautomationGetAutomationEvent, "get-automation-event", "", false, "Get Automation Event")
	_computeoptimizerautomationCmd.Flags().BoolVarP(&_computeoptimizerautomationGetAutomationRule, "get-automation-rule", "", false, "Get Automation Rule")
	_computeoptimizerautomationCmd.Flags().BoolVarP(&_computeoptimizerautomationGetEnrollmentConfiguration, "get-enrollment-configuration", "", false, "Get Enrollment Configuration")
	_computeoptimizerautomationCmd.Flags().BoolVarP(&_computeoptimizerautomationListAccounts, "list-accounts", "", false, "List Accounts")
	_computeoptimizerautomationCmd.Flags().BoolVarP(&_computeoptimizerautomationListAutomationEventSteps, "list-automation-event-steps", "", false, "List Automation Event Steps")
	_computeoptimizerautomationCmd.Flags().BoolVarP(&_computeoptimizerautomationListAutomationEventSummaries, "list-automation-event-summaries", "", false, "List Automation Event Summaries")
	_computeoptimizerautomationCmd.Flags().BoolVarP(&_computeoptimizerautomationListAutomationEvents, "list-automation-events", "", false, "List Automation Events")
	_computeoptimizerautomationCmd.Flags().BoolVarP(&_computeoptimizerautomationListAutomationRulePreview, "list-automation-rule-preview", "", false, "List Automation Rule Preview")
	_computeoptimizerautomationCmd.Flags().BoolVarP(&_computeoptimizerautomationListAutomationRulePreviewSummaries, "list-automation-rule-preview-summaries", "", false, "List Automation Rule Preview Summaries")
	_computeoptimizerautomationCmd.Flags().BoolVarP(&_computeoptimizerautomationListAutomationRules, "list-automation-rules", "", false, "List Automation Rules")
	_computeoptimizerautomationCmd.Flags().BoolVarP(&_computeoptimizerautomationListRecommendedActionSummaries, "list-recommended-action-summaries", "", false, "List Recommended Action Summaries")
	_computeoptimizerautomationCmd.Flags().BoolVarP(&_computeoptimizerautomationListRecommendedActions, "list-recommended-actions", "", false, "List Recommended Actions")
	_computeoptimizerautomationCmd.Flags().BoolVarP(&_computeoptimizerautomationListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_computeoptimizerautomationCmd.Flags().BoolVarP(&_computeoptimizerautomationRollbackAutomationEvent, "rollback-automation-event", "", false, "Rollback Automation Event")
	_computeoptimizerautomationCmd.Flags().BoolVarP(&_computeoptimizerautomationStartAutomationEvent, "start-automation-event", "", false, "Start Automation Event")
	_computeoptimizerautomationCmd.Flags().BoolVarP(&_computeoptimizerautomationTagResource, "tag-resource", "", false, "Tag Resource")
	_computeoptimizerautomationCmd.Flags().BoolVarP(&_computeoptimizerautomationUntagResource, "untag-resource", "", false, "Untag Resource")
	_computeoptimizerautomationCmd.Flags().BoolVarP(&_computeoptimizerautomationUpdateAutomationRule, "update-automation-rule", "", false, "Update Automation Rule")
	_computeoptimizerautomationCmd.Flags().BoolVarP(&_computeoptimizerautomationUpdateEnrollmentConfiguration, "update-enrollment-configuration", "", false, "Update Enrollment Configuration")

}
