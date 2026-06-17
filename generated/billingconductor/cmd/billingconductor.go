package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/billingconductor"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// billingconductorCmd represents the billingconductor command
var _billingconductorCmd = &cobra.Command{
	Use:   "billingconductor",
	Short: "AWS billingconductor CLI",
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
		client := billingconductor.NewFromConfig(cfg)
		if _billingconductorAssociateAccounts {
			billingconductor_AssociateAccounts(cfg, client)
			return
		}
		if _billingconductorAssociatePricingRules {
			billingconductor_AssociatePricingRules(cfg, client)
			return
		}
		if _billingconductorBatchAssociateResourcesToCustomLineItem {
			billingconductor_BatchAssociateResourcesToCustomLineItem(cfg, client)
			return
		}
		if _billingconductorBatchDisassociateResourcesFromCustomLineItem {
			billingconductor_BatchDisassociateResourcesFromCustomLineItem(cfg, client)
			return
		}
		if _billingconductorCreateBillingGroup {
			billingconductor_CreateBillingGroup(cfg, client)
			return
		}
		if _billingconductorCreateCustomLineItem {
			billingconductor_CreateCustomLineItem(cfg, client)
			return
		}
		if _billingconductorCreatePricingPlan {
			billingconductor_CreatePricingPlan(cfg, client)
			return
		}
		if _billingconductorCreatePricingRule {
			billingconductor_CreatePricingRule(cfg, client)
			return
		}
		if _billingconductorDeleteBillingGroup {
			billingconductor_DeleteBillingGroup(cfg, client)
			return
		}
		if _billingconductorDeleteCustomLineItem {
			billingconductor_DeleteCustomLineItem(cfg, client)
			return
		}
		if _billingconductorDeletePricingPlan {
			billingconductor_DeletePricingPlan(cfg, client)
			return
		}
		if _billingconductorDeletePricingRule {
			billingconductor_DeletePricingRule(cfg, client)
			return
		}
		if _billingconductorDisassociateAccounts {
			billingconductor_DisassociateAccounts(cfg, client)
			return
		}
		if _billingconductorDisassociatePricingRules {
			billingconductor_DisassociatePricingRules(cfg, client)
			return
		}
		if _billingconductorGetBillingGroupCostReport {
			billingconductor_GetBillingGroupCostReport(cfg, client)
			return
		}
		if _billingconductorListAccountAssociations {
			billingconductor_ListAccountAssociations(cfg, client)
			return
		}
		if _billingconductorListBillingGroupCostReports {
			billingconductor_ListBillingGroupCostReports(cfg, client)
			return
		}
		if _billingconductorListBillingGroups {
			billingconductor_ListBillingGroups(cfg, client)
			return
		}
		if _billingconductorListCustomLineItemVersions {
			billingconductor_ListCustomLineItemVersions(cfg, client)
			return
		}
		if _billingconductorListCustomLineItems {
			billingconductor_ListCustomLineItems(cfg, client)
			return
		}
		if _billingconductorListPricingPlans {
			billingconductor_ListPricingPlans(cfg, client)
			return
		}
		if _billingconductorListPricingPlansAssociatedWithPricingRule {
			billingconductor_ListPricingPlansAssociatedWithPricingRule(cfg, client)
			return
		}
		if _billingconductorListPricingRules {
			billingconductor_ListPricingRules(cfg, client)
			return
		}
		if _billingconductorListPricingRulesAssociatedToPricingPlan {
			billingconductor_ListPricingRulesAssociatedToPricingPlan(cfg, client)
			return
		}
		if _billingconductorListResourcesAssociatedToCustomLineItem {
			billingconductor_ListResourcesAssociatedToCustomLineItem(cfg, client)
			return
		}
		if _billingconductorListTagsForResource {
			billingconductor_ListTagsForResource(cfg, client)
			return
		}
		if _billingconductorTagResource {
			billingconductor_TagResource(cfg, client)
			return
		}
		if _billingconductorUntagResource {
			billingconductor_UntagResource(cfg, client)
			return
		}
		if _billingconductorUpdateBillingGroup {
			billingconductor_UpdateBillingGroup(cfg, client)
			return
		}
		if _billingconductorUpdateCustomLineItem {
			billingconductor_UpdateCustomLineItem(cfg, client)
			return
		}
		if _billingconductorUpdatePricingPlan {
			billingconductor_UpdatePricingPlan(cfg, client)
			return
		}
		if _billingconductorUpdatePricingRule {
			billingconductor_UpdatePricingRule(cfg, client)
			return
		}

	},
}

var (
	_billingconductorAssociateAccounts                            bool
	_billingconductorAssociatePricingRules                        bool
	_billingconductorBatchAssociateResourcesToCustomLineItem      bool
	_billingconductorBatchDisassociateResourcesFromCustomLineItem bool
	_billingconductorCreateBillingGroup                           bool
	_billingconductorCreateCustomLineItem                         bool
	_billingconductorCreatePricingPlan                            bool
	_billingconductorCreatePricingRule                            bool
	_billingconductorDeleteBillingGroup                           bool
	_billingconductorDeleteCustomLineItem                         bool
	_billingconductorDeletePricingPlan                            bool
	_billingconductorDeletePricingRule                            bool
	_billingconductorDisassociateAccounts                         bool
	_billingconductorDisassociatePricingRules                     bool
	_billingconductorGetBillingGroupCostReport                    bool
	_billingconductorListAccountAssociations                      bool
	_billingconductorListBillingGroupCostReports                  bool
	_billingconductorListBillingGroups                            bool
	_billingconductorListCustomLineItemVersions                   bool
	_billingconductorListCustomLineItems                          bool
	_billingconductorListPricingPlans                             bool
	_billingconductorListPricingPlansAssociatedWithPricingRule    bool
	_billingconductorListPricingRules                             bool
	_billingconductorListPricingRulesAssociatedToPricingPlan      bool
	_billingconductorListResourcesAssociatedToCustomLineItem      bool
	_billingconductorListTagsForResource                          bool
	_billingconductorTagResource                                  bool
	_billingconductorUntagResource                                bool
	_billingconductorUpdateBillingGroup                           bool
	_billingconductorUpdateCustomLineItem                         bool
	_billingconductorUpdatePricingPlan                            bool
	_billingconductorUpdatePricingRule                            bool

	_billingconductorAccountGrouping       string
	_billingconductorAccountId             string
	_billingconductorAccountIds            []string
	_billingconductorArn                   string
	_billingconductorBillingEntity         string
	_billingconductorBillingGroupArn       string
	_billingconductorBillingPeriod         string
	_billingconductorBillingPeriodRange    string
	_billingconductorChargeDetails         string
	_billingconductorClientToken           string
	_billingconductorComputationPreference string
	_billingconductorComputationRule       string
	_billingconductorDescription           string
	_billingconductorFilters               string
	_billingconductorGroupBy               string
	_billingconductorMaxResults            string
	_billingconductorModifierPercentage    string
	_billingconductorName                  string
	_billingconductorNextToken             string
	_billingconductorOperation             string
	_billingconductorPresentationDetails   string
	_billingconductorPricingPlanArn        string
	_billingconductorPricingRuleArn        string
	_billingconductorPricingRuleArns       []string
	_billingconductorPrimaryAccountId      string
	_billingconductorResourceArn           string
	_billingconductorResourceArns          []string
	_billingconductorScope                 string
	_billingconductorService               string
	_billingconductorStatus                string
	_billingconductorTagKeys               []string
	_billingconductorTags                  string
	_billingconductorTargetArn             string
	_billingconductorTiering               string
	_billingconductorType                  string
	_billingconductorUsageType             string
)

// Connects an array of account IDs in a consolidated billing family to a
// predefined billing group. The account IDs must be a part of the consolidated
// billing family during the current month, and not already associated with another
// billing group. The maximum number of accounts that can be associated in one call
// is 30.
func billingconductor_AssociateAccounts(cfg aws.Config, client *billingconductor.Client) {
	input := &billingconductor.AssociateAccountsInput{
		// AccountIds: []string, // Required
		// Arn: *string, // Required
	}

	if len(_billingconductorAccountIds) > 0 {
		input.AccountIds = append([]string(nil), _billingconductorAccountIds...)
	}
	if len(_billingconductorArn) > 0 {
		input.Arn = aws.String(_billingconductorArn)
	}

	if resp, err := client.AssociateAccounts(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Connects an array of PricingRuleArns to a defined PricingPlan . The maximum
// number PricingRuleArn that can be associated in one call is 30.
func billingconductor_AssociatePricingRules(cfg aws.Config, client *billingconductor.Client) {
	input := &billingconductor.AssociatePricingRulesInput{
		// Arn: *string, // Required
		// PricingRuleArns: []string, // Required
	}

	if len(_billingconductorArn) > 0 {
		input.Arn = aws.String(_billingconductorArn)
	}
	if len(_billingconductorPricingRuleArns) > 0 {
		input.PricingRuleArns = append([]string(nil), _billingconductorPricingRuleArns...)
	}

	if resp, err := client.AssociatePricingRules(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates a batch of resources to a percentage custom line item.
func billingconductor_BatchAssociateResourcesToCustomLineItem(cfg aws.Config, client *billingconductor.Client) {
	input := &billingconductor.BatchAssociateResourcesToCustomLineItemInput{
		// ResourceArns: []string, // Required
		// TargetArn: *string, // Required
	}

	if len(_billingconductorResourceArns) > 0 {
		input.ResourceArns = append([]string(nil), _billingconductorResourceArns...)
	}
	if len(_billingconductorTargetArn) > 0 {
		input.TargetArn = aws.String(_billingconductorTargetArn)
	}
	if len(_billingconductorBillingPeriodRange) > 0 {
		if err := assignInputField(input, "BillingPeriodRange", _billingconductorBillingPeriodRange); err != nil {
			log.Errorf("invalid --billing-period-range: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchAssociateResourcesToCustomLineItem(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates a batch of resources from a percentage custom line item.
func billingconductor_BatchDisassociateResourcesFromCustomLineItem(cfg aws.Config, client *billingconductor.Client) {
	input := &billingconductor.BatchDisassociateResourcesFromCustomLineItemInput{
		// ResourceArns: []string, // Required
		// TargetArn: *string, // Required
	}

	if len(_billingconductorResourceArns) > 0 {
		input.ResourceArns = append([]string(nil), _billingconductorResourceArns...)
	}
	if len(_billingconductorTargetArn) > 0 {
		input.TargetArn = aws.String(_billingconductorTargetArn)
	}
	if len(_billingconductorBillingPeriodRange) > 0 {
		if err := assignInputField(input, "BillingPeriodRange", _billingconductorBillingPeriodRange); err != nil {
			log.Errorf("invalid --billing-period-range: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchDisassociateResourcesFromCustomLineItem(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a billing group that resembles a consolidated billing family that
// Amazon Web Services charges, based off of the predefined pricing plan
// computation.
func billingconductor_CreateBillingGroup(cfg aws.Config, client *billingconductor.Client) {
	input := &billingconductor.CreateBillingGroupInput{
		// AccountGrouping: *types.AccountGrouping, // Required
		// ComputationPreference: *types.ComputationPreference, // Required
		// Name: *string, // Required
	}

	if len(_billingconductorAccountGrouping) > 0 {
		if err := assignInputField(input, "AccountGrouping", _billingconductorAccountGrouping); err != nil {
			log.Errorf("invalid --account-grouping: %s", err.Error())
			return
		}
	}
	if len(_billingconductorComputationPreference) > 0 {
		if err := assignInputField(input, "ComputationPreference", _billingconductorComputationPreference); err != nil {
			log.Errorf("invalid --computation-preference: %s", err.Error())
			return
		}
	}
	if len(_billingconductorName) > 0 {
		input.Name = aws.String(_billingconductorName)
	}
	if len(_billingconductorClientToken) > 0 {
		input.ClientToken = aws.String(_billingconductorClientToken)
	}
	if len(_billingconductorDescription) > 0 {
		input.Description = aws.String(_billingconductorDescription)
	}
	if len(_billingconductorPrimaryAccountId) > 0 {
		input.PrimaryAccountId = aws.String(_billingconductorPrimaryAccountId)
	}
	if len(_billingconductorTags) > 0 {
		if err := assignInputField(input, "Tags", _billingconductorTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateBillingGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a custom line item that can be used to create a one-time fixed charge
// that can be applied to a single billing group for the current or previous
// billing period. The one-time fixed charge is either a fee or discount.
func billingconductor_CreateCustomLineItem(cfg aws.Config, client *billingconductor.Client) {
	input := &billingconductor.CreateCustomLineItemInput{
		// BillingGroupArn: *string, // Required
		// ChargeDetails: *types.CustomLineItemChargeDetails, // Required
		// Description: *string, // Required
		// Name: *string, // Required
	}

	if len(_billingconductorBillingGroupArn) > 0 {
		input.BillingGroupArn = aws.String(_billingconductorBillingGroupArn)
	}
	if len(_billingconductorChargeDetails) > 0 {
		if err := assignInputField(input, "ChargeDetails", _billingconductorChargeDetails); err != nil {
			log.Errorf("invalid --charge-details: %s", err.Error())
			return
		}
	}
	if len(_billingconductorDescription) > 0 {
		input.Description = aws.String(_billingconductorDescription)
	}
	if len(_billingconductorName) > 0 {
		input.Name = aws.String(_billingconductorName)
	}
	if len(_billingconductorAccountId) > 0 {
		input.AccountId = aws.String(_billingconductorAccountId)
	}
	if len(_billingconductorBillingPeriodRange) > 0 {
		if err := assignInputField(input, "BillingPeriodRange", _billingconductorBillingPeriodRange); err != nil {
			log.Errorf("invalid --billing-period-range: %s", err.Error())
			return
		}
	}
	if len(_billingconductorClientToken) > 0 {
		input.ClientToken = aws.String(_billingconductorClientToken)
	}
	if len(_billingconductorComputationRule) > 0 {
		if err := assignInputField(input, "ComputationRule", _billingconductorComputationRule); err != nil {
			log.Errorf("invalid --computation-rule: %s", err.Error())
			return
		}
	}
	if len(_billingconductorPresentationDetails) > 0 {
		if err := assignInputField(input, "PresentationDetails", _billingconductorPresentationDetails); err != nil {
			log.Errorf("invalid --presentation-details: %s", err.Error())
			return
		}
	}
	if len(_billingconductorTags) > 0 {
		if err := assignInputField(input, "Tags", _billingconductorTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCustomLineItem(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a pricing plan that is used for computing Amazon Web Services charges
// for billing groups.
func billingconductor_CreatePricingPlan(cfg aws.Config, client *billingconductor.Client) {
	input := &billingconductor.CreatePricingPlanInput{
		// Name: *string, // Required
	}

	if len(_billingconductorName) > 0 {
		input.Name = aws.String(_billingconductorName)
	}
	if len(_billingconductorClientToken) > 0 {
		input.ClientToken = aws.String(_billingconductorClientToken)
	}
	if len(_billingconductorDescription) > 0 {
		input.Description = aws.String(_billingconductorDescription)
	}
	if len(_billingconductorPricingRuleArns) > 0 {
		input.PricingRuleArns = append([]string(nil), _billingconductorPricingRuleArns...)
	}
	if len(_billingconductorTags) > 0 {
		if err := assignInputField(input, "Tags", _billingconductorTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreatePricingPlan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a pricing rule can be associated to a pricing plan, or a set of
// pricing plans.
func billingconductor_CreatePricingRule(cfg aws.Config, client *billingconductor.Client) {
	input := &billingconductor.CreatePricingRuleInput{
		// Name: *string, // Required
		// Scope: types.PricingRuleScope, // Required
		// Type: types.PricingRuleType, // Required
	}

	if len(_billingconductorName) > 0 {
		input.Name = aws.String(_billingconductorName)
	}
	if len(_billingconductorScope) > 0 {
		if err := assignInputField(input, "Scope", _billingconductorScope); err != nil {
			log.Errorf("invalid --scope: %s", err.Error())
			return
		}
	}
	if len(_billingconductorType) > 0 {
		if err := assignInputField(input, "Type", _billingconductorType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_billingconductorBillingEntity) > 0 {
		input.BillingEntity = aws.String(_billingconductorBillingEntity)
	}
	if len(_billingconductorClientToken) > 0 {
		input.ClientToken = aws.String(_billingconductorClientToken)
	}
	if len(_billingconductorDescription) > 0 {
		input.Description = aws.String(_billingconductorDescription)
	}
	if len(_billingconductorModifierPercentage) > 0 {
		if err := assignInputField(input, "ModifierPercentage", _billingconductorModifierPercentage); err != nil {
			log.Errorf("invalid --modifier-percentage: %s", err.Error())
			return
		}
	}
	if len(_billingconductorOperation) > 0 {
		input.Operation = aws.String(_billingconductorOperation)
	}
	if len(_billingconductorService) > 0 {
		input.Service = aws.String(_billingconductorService)
	}
	if len(_billingconductorTags) > 0 {
		if err := assignInputField(input, "Tags", _billingconductorTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_billingconductorTiering) > 0 {
		if err := assignInputField(input, "Tiering", _billingconductorTiering); err != nil {
			log.Errorf("invalid --tiering: %s", err.Error())
			return
		}
	}
	if len(_billingconductorUsageType) > 0 {
		input.UsageType = aws.String(_billingconductorUsageType)
	}

	if resp, err := client.CreatePricingRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a billing group.
func billingconductor_DeleteBillingGroup(cfg aws.Config, client *billingconductor.Client) {
	input := &billingconductor.DeleteBillingGroupInput{
		// Arn: *string, // Required
	}

	if len(_billingconductorArn) > 0 {
		input.Arn = aws.String(_billingconductorArn)
	}

	if resp, err := client.DeleteBillingGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the custom line item identified by the given ARN in the current, or
// previous billing period.
func billingconductor_DeleteCustomLineItem(cfg aws.Config, client *billingconductor.Client) {
	input := &billingconductor.DeleteCustomLineItemInput{
		// Arn: *string, // Required
	}

	if len(_billingconductorArn) > 0 {
		input.Arn = aws.String(_billingconductorArn)
	}
	if len(_billingconductorBillingPeriodRange) > 0 {
		if err := assignInputField(input, "BillingPeriodRange", _billingconductorBillingPeriodRange); err != nil {
			log.Errorf("invalid --billing-period-range: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteCustomLineItem(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a pricing plan. The pricing plan must not be associated with any
// billing groups to delete successfully.
func billingconductor_DeletePricingPlan(cfg aws.Config, client *billingconductor.Client) {
	input := &billingconductor.DeletePricingPlanInput{
		// Arn: *string, // Required
	}

	if len(_billingconductorArn) > 0 {
		input.Arn = aws.String(_billingconductorArn)
	}

	if resp, err := client.DeletePricingPlan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the pricing rule that's identified by the input Amazon Resource Name
// (ARN).
func billingconductor_DeletePricingRule(cfg aws.Config, client *billingconductor.Client) {
	input := &billingconductor.DeletePricingRuleInput{
		// Arn: *string, // Required
	}

	if len(_billingconductorArn) > 0 {
		input.Arn = aws.String(_billingconductorArn)
	}

	if resp, err := client.DeletePricingRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the specified list of account IDs from the given billing group.
func billingconductor_DisassociateAccounts(cfg aws.Config, client *billingconductor.Client) {
	input := &billingconductor.DisassociateAccountsInput{
		// AccountIds: []string, // Required
		// Arn: *string, // Required
	}

	if len(_billingconductorAccountIds) > 0 {
		input.AccountIds = append([]string(nil), _billingconductorAccountIds...)
	}
	if len(_billingconductorArn) > 0 {
		input.Arn = aws.String(_billingconductorArn)
	}

	if resp, err := client.DisassociateAccounts(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates a list of pricing rules from a pricing plan.
func billingconductor_DisassociatePricingRules(cfg aws.Config, client *billingconductor.Client) {
	input := &billingconductor.DisassociatePricingRulesInput{
		// Arn: *string, // Required
		// PricingRuleArns: []string, // Required
	}

	if len(_billingconductorArn) > 0 {
		input.Arn = aws.String(_billingconductorArn)
	}
	if len(_billingconductorPricingRuleArns) > 0 {
		input.PricingRuleArns = append([]string(nil), _billingconductorPricingRuleArns...)
	}

	if resp, err := client.DisassociatePricingRules(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the margin summary report, which includes the Amazon Web Services
// cost and charged amount (pro forma cost) by Amazon Web Services service for a
// specific billing group.
func billingconductor_GetBillingGroupCostReport(cfg aws.Config, client *billingconductor.Client) {
	input := &billingconductor.GetBillingGroupCostReportInput{
		// Arn: *string, // Required
	}

	if len(_billingconductorArn) > 0 {
		input.Arn = aws.String(_billingconductorArn)
	}
	if len(_billingconductorBillingPeriodRange) > 0 {
		if err := assignInputField(input, "BillingPeriodRange", _billingconductorBillingPeriodRange); err != nil {
			log.Errorf("invalid --billing-period-range: %s", err.Error())
			return
		}
	}
	if len(_billingconductorGroupBy) > 0 {
		if err := assignInputField(input, "GroupBy", _billingconductorGroupBy); err != nil {
			log.Errorf("invalid --group-by: %s", err.Error())
			return
		}
	}
	if len(_billingconductorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _billingconductorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_billingconductorNextToken) > 0 {
		input.NextToken = aws.String(_billingconductorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetBillingGroupCostReport(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*billingconductor.GetBillingGroupCostReportOutput
	p := billingconductor.NewGetBillingGroupCostReportPaginator(client, input)
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

// This is a paginated call to list linked accounts that are linked to the payer
// account for the specified time period. If no information is provided, the
// current billing period is used. The response will optionally include the billing
// group that's associated with the linked account.
func billingconductor_ListAccountAssociations(cfg aws.Config, client *billingconductor.Client) {
	input := &billingconductor.ListAccountAssociationsInput{}

	if len(_billingconductorBillingPeriod) > 0 {
		input.BillingPeriod = aws.String(_billingconductorBillingPeriod)
	}
	if len(_billingconductorFilters) > 0 {
		if err := assignInputField(input, "Filters", _billingconductorFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_billingconductorNextToken) > 0 {
		input.NextToken = aws.String(_billingconductorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAccountAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*billingconductor.ListAccountAssociationsOutput
	p := billingconductor.NewListAccountAssociationsPaginator(client, input)
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

// A paginated call to retrieve a summary report of actual Amazon Web Services
// charges and the calculated Amazon Web Services charges based on the associated
// pricing plan of a billing group.
func billingconductor_ListBillingGroupCostReports(cfg aws.Config, client *billingconductor.Client) {
	input := &billingconductor.ListBillingGroupCostReportsInput{}

	if len(_billingconductorBillingPeriod) > 0 {
		input.BillingPeriod = aws.String(_billingconductorBillingPeriod)
	}
	if len(_billingconductorFilters) > 0 {
		if err := assignInputField(input, "Filters", _billingconductorFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_billingconductorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _billingconductorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_billingconductorNextToken) > 0 {
		input.NextToken = aws.String(_billingconductorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListBillingGroupCostReports(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*billingconductor.ListBillingGroupCostReportsOutput
	p := billingconductor.NewListBillingGroupCostReportsPaginator(client, input)
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

// A paginated call to retrieve a list of billing groups for the given billing
// period. If you don't provide a billing group, the current billing period is
// used.
func billingconductor_ListBillingGroups(cfg aws.Config, client *billingconductor.Client) {
	input := &billingconductor.ListBillingGroupsInput{}

	if len(_billingconductorBillingPeriod) > 0 {
		input.BillingPeriod = aws.String(_billingconductorBillingPeriod)
	}
	if len(_billingconductorFilters) > 0 {
		if err := assignInputField(input, "Filters", _billingconductorFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_billingconductorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _billingconductorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_billingconductorNextToken) > 0 {
		input.NextToken = aws.String(_billingconductorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListBillingGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*billingconductor.ListBillingGroupsOutput
	p := billingconductor.NewListBillingGroupsPaginator(client, input)
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

// A paginated call to get a list of all custom line item versions.
func billingconductor_ListCustomLineItemVersions(cfg aws.Config, client *billingconductor.Client) {
	input := &billingconductor.ListCustomLineItemVersionsInput{
		// Arn: *string, // Required
	}

	if len(_billingconductorArn) > 0 {
		input.Arn = aws.String(_billingconductorArn)
	}
	if len(_billingconductorFilters) > 0 {
		if err := assignInputField(input, "Filters", _billingconductorFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_billingconductorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _billingconductorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_billingconductorNextToken) > 0 {
		input.NextToken = aws.String(_billingconductorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCustomLineItemVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*billingconductor.ListCustomLineItemVersionsOutput
	p := billingconductor.NewListCustomLineItemVersionsPaginator(client, input)
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

// A paginated call to get a list of all custom line items (FFLIs) for the given
// billing period. If you don't provide a billing period, the current billing
// period is used.
func billingconductor_ListCustomLineItems(cfg aws.Config, client *billingconductor.Client) {
	input := &billingconductor.ListCustomLineItemsInput{}

	if len(_billingconductorBillingPeriod) > 0 {
		input.BillingPeriod = aws.String(_billingconductorBillingPeriod)
	}
	if len(_billingconductorFilters) > 0 {
		if err := assignInputField(input, "Filters", _billingconductorFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_billingconductorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _billingconductorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_billingconductorNextToken) > 0 {
		input.NextToken = aws.String(_billingconductorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCustomLineItems(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*billingconductor.ListCustomLineItemsOutput
	p := billingconductor.NewListCustomLineItemsPaginator(client, input)
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

// A paginated call to get pricing plans for the given billing period. If you
// don't provide a billing period, the current billing period is used.
func billingconductor_ListPricingPlans(cfg aws.Config, client *billingconductor.Client) {
	input := &billingconductor.ListPricingPlansInput{}

	if len(_billingconductorBillingPeriod) > 0 {
		input.BillingPeriod = aws.String(_billingconductorBillingPeriod)
	}
	if len(_billingconductorFilters) > 0 {
		if err := assignInputField(input, "Filters", _billingconductorFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_billingconductorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _billingconductorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_billingconductorNextToken) > 0 {
		input.NextToken = aws.String(_billingconductorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPricingPlans(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*billingconductor.ListPricingPlansOutput
	p := billingconductor.NewListPricingPlansPaginator(client, input)
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

// A list of the pricing plans that are associated with a pricing rule.
func billingconductor_ListPricingPlansAssociatedWithPricingRule(cfg aws.Config, client *billingconductor.Client) {
	input := &billingconductor.ListPricingPlansAssociatedWithPricingRuleInput{
		// PricingRuleArn: *string, // Required
	}

	if len(_billingconductorPricingRuleArn) > 0 {
		input.PricingRuleArn = aws.String(_billingconductorPricingRuleArn)
	}
	if len(_billingconductorBillingPeriod) > 0 {
		input.BillingPeriod = aws.String(_billingconductorBillingPeriod)
	}
	if len(_billingconductorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _billingconductorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_billingconductorNextToken) > 0 {
		input.NextToken = aws.String(_billingconductorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPricingPlansAssociatedWithPricingRule(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*billingconductor.ListPricingPlansAssociatedWithPricingRuleOutput
	p := billingconductor.NewListPricingPlansAssociatedWithPricingRulePaginator(client, input)
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

// Describes a pricing rule that can be associated to a pricing plan, or set of
// pricing plans.
func billingconductor_ListPricingRules(cfg aws.Config, client *billingconductor.Client) {
	input := &billingconductor.ListPricingRulesInput{}

	if len(_billingconductorBillingPeriod) > 0 {
		input.BillingPeriod = aws.String(_billingconductorBillingPeriod)
	}
	if len(_billingconductorFilters) > 0 {
		if err := assignInputField(input, "Filters", _billingconductorFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_billingconductorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _billingconductorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_billingconductorNextToken) > 0 {
		input.NextToken = aws.String(_billingconductorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPricingRules(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*billingconductor.ListPricingRulesOutput
	p := billingconductor.NewListPricingRulesPaginator(client, input)
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

// Lists the pricing rules that are associated with a pricing plan.
func billingconductor_ListPricingRulesAssociatedToPricingPlan(cfg aws.Config, client *billingconductor.Client) {
	input := &billingconductor.ListPricingRulesAssociatedToPricingPlanInput{
		// PricingPlanArn: *string, // Required
	}

	if len(_billingconductorPricingPlanArn) > 0 {
		input.PricingPlanArn = aws.String(_billingconductorPricingPlanArn)
	}
	if len(_billingconductorBillingPeriod) > 0 {
		input.BillingPeriod = aws.String(_billingconductorBillingPeriod)
	}
	if len(_billingconductorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _billingconductorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_billingconductorNextToken) > 0 {
		input.NextToken = aws.String(_billingconductorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPricingRulesAssociatedToPricingPlan(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*billingconductor.ListPricingRulesAssociatedToPricingPlanOutput
	p := billingconductor.NewListPricingRulesAssociatedToPricingPlanPaginator(client, input)
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

// List the resources that are associated to a custom line item.
func billingconductor_ListResourcesAssociatedToCustomLineItem(cfg aws.Config, client *billingconductor.Client) {
	input := &billingconductor.ListResourcesAssociatedToCustomLineItemInput{
		// Arn: *string, // Required
	}

	if len(_billingconductorArn) > 0 {
		input.Arn = aws.String(_billingconductorArn)
	}
	if len(_billingconductorBillingPeriod) > 0 {
		input.BillingPeriod = aws.String(_billingconductorBillingPeriod)
	}
	if len(_billingconductorFilters) > 0 {
		if err := assignInputField(input, "Filters", _billingconductorFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_billingconductorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _billingconductorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_billingconductorNextToken) > 0 {
		input.NextToken = aws.String(_billingconductorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListResourcesAssociatedToCustomLineItem(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*billingconductor.ListResourcesAssociatedToCustomLineItemOutput
	p := billingconductor.NewListResourcesAssociatedToCustomLineItemPaginator(client, input)
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

// A list the tags for a resource.
func billingconductor_ListTagsForResource(cfg aws.Config, client *billingconductor.Client) {
	input := &billingconductor.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_billingconductorResourceArn) > 0 {
		input.ResourceArn = aws.String(_billingconductorResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates the specified tags to a resource with the specified resourceArn . If
// existing tags on a resource are not specified in the request parameters, they
// are not changed.
func billingconductor_TagResource(cfg aws.Config, client *billingconductor.Client) {
	input := &billingconductor.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_billingconductorResourceArn) > 0 {
		input.ResourceArn = aws.String(_billingconductorResourceArn)
	}
	if len(_billingconductorTags) > 0 {
		if err := assignInputField(input, "Tags", _billingconductorTags); err != nil {
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

// Deletes specified tags from a resource.
func billingconductor_UntagResource(cfg aws.Config, client *billingconductor.Client) {
	input := &billingconductor.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_billingconductorResourceArn) > 0 {
		input.ResourceArn = aws.String(_billingconductorResourceArn)
	}
	if len(_billingconductorTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _billingconductorTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This updates an existing billing group.
func billingconductor_UpdateBillingGroup(cfg aws.Config, client *billingconductor.Client) {
	input := &billingconductor.UpdateBillingGroupInput{
		// Arn: *string, // Required
	}

	if len(_billingconductorArn) > 0 {
		input.Arn = aws.String(_billingconductorArn)
	}
	if len(_billingconductorAccountGrouping) > 0 {
		if err := assignInputField(input, "AccountGrouping", _billingconductorAccountGrouping); err != nil {
			log.Errorf("invalid --account-grouping: %s", err.Error())
			return
		}
	}
	if len(_billingconductorComputationPreference) > 0 {
		if err := assignInputField(input, "ComputationPreference", _billingconductorComputationPreference); err != nil {
			log.Errorf("invalid --computation-preference: %s", err.Error())
			return
		}
	}
	if len(_billingconductorDescription) > 0 {
		input.Description = aws.String(_billingconductorDescription)
	}
	if len(_billingconductorName) > 0 {
		input.Name = aws.String(_billingconductorName)
	}
	if len(_billingconductorStatus) > 0 {
		if err := assignInputField(input, "Status", _billingconductorStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateBillingGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update an existing custom line item in the current or previous billing period.
func billingconductor_UpdateCustomLineItem(cfg aws.Config, client *billingconductor.Client) {
	input := &billingconductor.UpdateCustomLineItemInput{
		// Arn: *string, // Required
	}

	if len(_billingconductorArn) > 0 {
		input.Arn = aws.String(_billingconductorArn)
	}
	if len(_billingconductorBillingPeriodRange) > 0 {
		if err := assignInputField(input, "BillingPeriodRange", _billingconductorBillingPeriodRange); err != nil {
			log.Errorf("invalid --billing-period-range: %s", err.Error())
			return
		}
	}
	if len(_billingconductorChargeDetails) > 0 {
		if err := assignInputField(input, "ChargeDetails", _billingconductorChargeDetails); err != nil {
			log.Errorf("invalid --charge-details: %s", err.Error())
			return
		}
	}
	if len(_billingconductorDescription) > 0 {
		input.Description = aws.String(_billingconductorDescription)
	}
	if len(_billingconductorName) > 0 {
		input.Name = aws.String(_billingconductorName)
	}

	if resp, err := client.UpdateCustomLineItem(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This updates an existing pricing plan.
func billingconductor_UpdatePricingPlan(cfg aws.Config, client *billingconductor.Client) {
	input := &billingconductor.UpdatePricingPlanInput{
		// Arn: *string, // Required
	}

	if len(_billingconductorArn) > 0 {
		input.Arn = aws.String(_billingconductorArn)
	}
	if len(_billingconductorDescription) > 0 {
		input.Description = aws.String(_billingconductorDescription)
	}
	if len(_billingconductorName) > 0 {
		input.Name = aws.String(_billingconductorName)
	}

	if resp, err := client.UpdatePricingPlan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing pricing rule.
func billingconductor_UpdatePricingRule(cfg aws.Config, client *billingconductor.Client) {
	input := &billingconductor.UpdatePricingRuleInput{
		// Arn: *string, // Required
	}

	if len(_billingconductorArn) > 0 {
		input.Arn = aws.String(_billingconductorArn)
	}
	if len(_billingconductorDescription) > 0 {
		input.Description = aws.String(_billingconductorDescription)
	}
	if len(_billingconductorModifierPercentage) > 0 {
		if err := assignInputField(input, "ModifierPercentage", _billingconductorModifierPercentage); err != nil {
			log.Errorf("invalid --modifier-percentage: %s", err.Error())
			return
		}
	}
	if len(_billingconductorName) > 0 {
		input.Name = aws.String(_billingconductorName)
	}
	if len(_billingconductorTiering) > 0 {
		if err := assignInputField(input, "Tiering", _billingconductorTiering); err != nil {
			log.Errorf("invalid --tiering: %s", err.Error())
			return
		}
	}
	if len(_billingconductorType) > 0 {
		if err := assignInputField(input, "Type", _billingconductorType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdatePricingRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_billingconductorCmd)
	_billingconductorCmd.Flags().SortFlags = false

	_billingconductorCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_billingconductorCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_billingconductorCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_billingconductorCmd.Flags().StringVarP(&_billingconductorAccountGrouping, "account-grouping", "", "", "Account Grouping")
	_billingconductorCmd.Flags().StringVarP(&_billingconductorAccountId, "account-id", "", "", "Account ID")
	_billingconductorCmd.Flags().StringSliceVarP(&_billingconductorAccountIds, "account-ids", "", nil, "Account Ids")
	_billingconductorCmd.Flags().StringVarP(&_billingconductorArn, "arn", "", "", "ARN")
	_billingconductorCmd.Flags().StringVarP(&_billingconductorBillingEntity, "billing-entity", "", "", "Billing Entity")
	_billingconductorCmd.Flags().StringVarP(&_billingconductorBillingGroupArn, "billing-group-arn", "", "", "Billing Group ARN")
	_billingconductorCmd.Flags().StringVarP(&_billingconductorBillingPeriod, "billing-period", "", "", "Billing Period")
	_billingconductorCmd.Flags().StringVarP(&_billingconductorBillingPeriodRange, "billing-period-range", "", "", "Billing Period Range")
	_billingconductorCmd.Flags().StringVarP(&_billingconductorChargeDetails, "charge-details", "", "", "Charge Details")
	_billingconductorCmd.Flags().StringVarP(&_billingconductorClientToken, "client-token", "", "", "Client Token")
	_billingconductorCmd.Flags().StringVarP(&_billingconductorComputationPreference, "computation-preference", "", "", "Computation Preference")
	_billingconductorCmd.Flags().StringVarP(&_billingconductorComputationRule, "computation-rule", "", "", "Computation Rule")
	_billingconductorCmd.Flags().StringVarP(&_billingconductorDescription, "description", "", "", "Description")
	_billingconductorCmd.Flags().StringVarP(&_billingconductorFilters, "filters", "", "", "Filters")
	_billingconductorCmd.Flags().StringVarP(&_billingconductorGroupBy, "group-by", "", "", "Group By")
	_billingconductorCmd.Flags().StringVarP(&_billingconductorMaxResults, "max-results", "", "", "Max Results")
	_billingconductorCmd.Flags().StringVarP(&_billingconductorModifierPercentage, "modifier-percentage", "", "", "Modifier Percentage")
	_billingconductorCmd.Flags().StringVarP(&_billingconductorName, "name", "", "", "Name")
	_billingconductorCmd.Flags().StringVarP(&_billingconductorNextToken, "next-token", "", "", "Next Token")
	_billingconductorCmd.Flags().StringVarP(&_billingconductorOperation, "operation", "", "", "Operation")
	_billingconductorCmd.Flags().StringVarP(&_billingconductorPresentationDetails, "presentation-details", "", "", "Presentation Details")
	_billingconductorCmd.Flags().StringVarP(&_billingconductorPricingPlanArn, "pricing-plan-arn", "", "", "Pricing Plan ARN")
	_billingconductorCmd.Flags().StringVarP(&_billingconductorPricingRuleArn, "pricing-rule-arn", "", "", "Pricing Rule ARN")
	_billingconductorCmd.Flags().StringSliceVarP(&_billingconductorPricingRuleArns, "pricing-rule-arns", "", nil, "Pricing Rule Arns")
	_billingconductorCmd.Flags().StringVarP(&_billingconductorPrimaryAccountId, "primary-account-id", "", "", "Primary Account ID")
	_billingconductorCmd.Flags().StringVarP(&_billingconductorResourceArn, "resource-arn", "", "", "Resource ARN")
	_billingconductorCmd.Flags().StringSliceVarP(&_billingconductorResourceArns, "resource-arns", "", nil, "Resource Arns")
	_billingconductorCmd.Flags().StringVarP(&_billingconductorScope, "scope", "", "", "Scope")
	_billingconductorCmd.Flags().StringVarP(&_billingconductorService, "service", "", "", "Service")
	_billingconductorCmd.Flags().StringVarP(&_billingconductorStatus, "status", "", "", "Status")
	_billingconductorCmd.Flags().StringSliceVarP(&_billingconductorTagKeys, "tag-keys", "", nil, "Tag Keys")
	_billingconductorCmd.Flags().StringVarP(&_billingconductorTags, "tags", "", "", "Tags")
	_billingconductorCmd.Flags().StringVarP(&_billingconductorTargetArn, "target-arn", "", "", "Target ARN")
	_billingconductorCmd.Flags().StringVarP(&_billingconductorTiering, "tiering", "", "", "Tiering")
	_billingconductorCmd.Flags().StringVarP(&_billingconductorType, "type", "", "", "Type")
	_billingconductorCmd.Flags().StringVarP(&_billingconductorUsageType, "usage-type", "", "", "Usage Type")

	_billingconductorCmd.Flags().BoolVarP(&_billingconductorAssociateAccounts, "associate-accounts", "", false, "Associate Accounts")
	_billingconductorCmd.Flags().BoolVarP(&_billingconductorAssociatePricingRules, "associate-pricing-rules", "", false, "Associate Pricing Rules")
	_billingconductorCmd.Flags().BoolVarP(&_billingconductorBatchAssociateResourcesToCustomLineItem, "batch-associate-resources-to-custom-line-item", "", false, "Batch Associate Resources To Custom Line Item")
	_billingconductorCmd.Flags().BoolVarP(&_billingconductorBatchDisassociateResourcesFromCustomLineItem, "batch-disassociate-resources-from-custom-line-item", "", false, "Batch Disassociate Resources From Custom Line Item")
	_billingconductorCmd.Flags().BoolVarP(&_billingconductorCreateBillingGroup, "create-billing-group", "", false, "Create Billing Group")
	_billingconductorCmd.Flags().BoolVarP(&_billingconductorCreateCustomLineItem, "create-custom-line-item", "", false, "Create Custom Line Item")
	_billingconductorCmd.Flags().BoolVarP(&_billingconductorCreatePricingPlan, "create-pricing-plan", "", false, "Create Pricing Plan")
	_billingconductorCmd.Flags().BoolVarP(&_billingconductorCreatePricingRule, "create-pricing-rule", "", false, "Create Pricing Rule")
	_billingconductorCmd.Flags().BoolVarP(&_billingconductorDeleteBillingGroup, "delete-billing-group", "", false, "Delete Billing Group")
	_billingconductorCmd.Flags().BoolVarP(&_billingconductorDeleteCustomLineItem, "delete-custom-line-item", "", false, "Delete Custom Line Item")
	_billingconductorCmd.Flags().BoolVarP(&_billingconductorDeletePricingPlan, "delete-pricing-plan", "", false, "Delete Pricing Plan")
	_billingconductorCmd.Flags().BoolVarP(&_billingconductorDeletePricingRule, "delete-pricing-rule", "", false, "Delete Pricing Rule")
	_billingconductorCmd.Flags().BoolVarP(&_billingconductorDisassociateAccounts, "disassociate-accounts", "", false, "Disassociate Accounts")
	_billingconductorCmd.Flags().BoolVarP(&_billingconductorDisassociatePricingRules, "disassociate-pricing-rules", "", false, "Disassociate Pricing Rules")
	_billingconductorCmd.Flags().BoolVarP(&_billingconductorGetBillingGroupCostReport, "get-billing-group-cost-report", "", false, "Get Billing Group Cost Report")
	_billingconductorCmd.Flags().BoolVarP(&_billingconductorListAccountAssociations, "list-account-associations", "", false, "List Account Associations")
	_billingconductorCmd.Flags().BoolVarP(&_billingconductorListBillingGroupCostReports, "list-billing-group-cost-reports", "", false, "List Billing Group Cost Reports")
	_billingconductorCmd.Flags().BoolVarP(&_billingconductorListBillingGroups, "list-billing-groups", "", false, "List Billing Groups")
	_billingconductorCmd.Flags().BoolVarP(&_billingconductorListCustomLineItemVersions, "list-custom-line-item-versions", "", false, "List Custom Line Item Versions")
	_billingconductorCmd.Flags().BoolVarP(&_billingconductorListCustomLineItems, "list-custom-line-items", "", false, "List Custom Line Items")
	_billingconductorCmd.Flags().BoolVarP(&_billingconductorListPricingPlans, "list-pricing-plans", "", false, "List Pricing Plans")
	_billingconductorCmd.Flags().BoolVarP(&_billingconductorListPricingPlansAssociatedWithPricingRule, "list-pricing-plans-associated-with-pricing-rule", "", false, "List Pricing Plans Associated With Pricing Rule")
	_billingconductorCmd.Flags().BoolVarP(&_billingconductorListPricingRules, "list-pricing-rules", "", false, "List Pricing Rules")
	_billingconductorCmd.Flags().BoolVarP(&_billingconductorListPricingRulesAssociatedToPricingPlan, "list-pricing-rules-associated-to-pricing-plan", "", false, "List Pricing Rules Associated To Pricing Plan")
	_billingconductorCmd.Flags().BoolVarP(&_billingconductorListResourcesAssociatedToCustomLineItem, "list-resources-associated-to-custom-line-item", "", false, "List Resources Associated To Custom Line Item")
	_billingconductorCmd.Flags().BoolVarP(&_billingconductorListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_billingconductorCmd.Flags().BoolVarP(&_billingconductorTagResource, "tag-resource", "", false, "Tag Resource")
	_billingconductorCmd.Flags().BoolVarP(&_billingconductorUntagResource, "untag-resource", "", false, "Untag Resource")
	_billingconductorCmd.Flags().BoolVarP(&_billingconductorUpdateBillingGroup, "update-billing-group", "", false, "Update Billing Group")
	_billingconductorCmd.Flags().BoolVarP(&_billingconductorUpdateCustomLineItem, "update-custom-line-item", "", false, "Update Custom Line Item")
	_billingconductorCmd.Flags().BoolVarP(&_billingconductorUpdatePricingPlan, "update-pricing-plan", "", false, "Update Pricing Plan")
	_billingconductorCmd.Flags().BoolVarP(&_billingconductorUpdatePricingRule, "update-pricing-rule", "", false, "Update Pricing Rule")

}
