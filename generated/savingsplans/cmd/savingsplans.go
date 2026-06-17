package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/savingsplans"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// savingsplansCmd represents the savingsplans command
var _savingsplansCmd = &cobra.Command{
	Use:   "savingsplans",
	Short: "AWS savingsplans CLI",
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
		client := savingsplans.NewFromConfig(cfg)
		if _savingsplansCreateSavingsPlan {
			savingsplans_CreateSavingsPlan(cfg, client)
			return
		}
		if _savingsplansDeleteQueuedSavingsPlan {
			savingsplans_DeleteQueuedSavingsPlan(cfg, client)
			return
		}
		if _savingsplansDescribeSavingsPlanRates {
			savingsplans_DescribeSavingsPlanRates(cfg, client)
			return
		}
		if _savingsplansDescribeSavingsPlans {
			savingsplans_DescribeSavingsPlans(cfg, client)
			return
		}
		if _savingsplansDescribeSavingsPlansOfferingRates {
			savingsplans_DescribeSavingsPlansOfferingRates(cfg, client)
			return
		}
		if _savingsplansDescribeSavingsPlansOfferings {
			savingsplans_DescribeSavingsPlansOfferings(cfg, client)
			return
		}
		if _savingsplansListTagsForResource {
			savingsplans_ListTagsForResource(cfg, client)
			return
		}
		if _savingsplansReturnSavingsPlan {
			savingsplans_ReturnSavingsPlan(cfg, client)
			return
		}
		if _savingsplansTagResource {
			savingsplans_TagResource(cfg, client)
			return
		}
		if _savingsplansUntagResource {
			savingsplans_UntagResource(cfg, client)
			return
		}

	},
}

var (
	_savingsplansCreateSavingsPlan                 bool
	_savingsplansDeleteQueuedSavingsPlan           bool
	_savingsplansDescribeSavingsPlanRates          bool
	_savingsplansDescribeSavingsPlans              bool
	_savingsplansDescribeSavingsPlansOfferingRates bool
	_savingsplansDescribeSavingsPlansOfferings     bool
	_savingsplansListTagsForResource               bool
	_savingsplansReturnSavingsPlan                 bool
	_savingsplansTagResource                       bool
	_savingsplansUntagResource                     bool

	_savingsplansClientToken               string
	_savingsplansCommitment                string
	_savingsplansCurrencies                string
	_savingsplansDescriptions              []string
	_savingsplansDurations                 string
	_savingsplansFilters                   string
	_savingsplansMaxResults                string
	_savingsplansNextToken                 string
	_savingsplansOfferingIds               []string
	_savingsplansOperations                []string
	_savingsplansPaymentOptions            string
	_savingsplansPlanTypes                 string
	_savingsplansProductType               string
	_savingsplansProducts                  string
	_savingsplansPurchaseTime              string
	_savingsplansResourceArn               string
	_savingsplansSavingsPlanArns           []string
	_savingsplansSavingsPlanId             string
	_savingsplansSavingsPlanIds            []string
	_savingsplansSavingsPlanOfferingId     string
	_savingsplansSavingsPlanOfferingIds    []string
	_savingsplansSavingsPlanPaymentOptions string
	_savingsplansSavingsPlanTypes          string
	_savingsplansServiceCodes              []string
	_savingsplansStates                    string
	_savingsplansTagKeys                   []string
	_savingsplansTags                      string
	_savingsplansUpfrontPaymentAmount      string
	_savingsplansUsageTypes                []string
)

// Creates a Savings Plan.
func savingsplans_CreateSavingsPlan(cfg aws.Config, client *savingsplans.Client) {
	input := &savingsplans.CreateSavingsPlanInput{
		// Commitment: *string, // Required
		// SavingsPlanOfferingId: *string, // Required
	}

	if len(_savingsplansCommitment) > 0 {
		input.Commitment = aws.String(_savingsplansCommitment)
	}
	if len(_savingsplansSavingsPlanOfferingId) > 0 {
		input.SavingsPlanOfferingId = aws.String(_savingsplansSavingsPlanOfferingId)
	}
	if len(_savingsplansClientToken) > 0 {
		input.ClientToken = aws.String(_savingsplansClientToken)
	}
	if len(_savingsplansPurchaseTime) > 0 {
		if err := assignInputField(input, "PurchaseTime", _savingsplansPurchaseTime); err != nil {
			log.Errorf("invalid --purchase-time: %s", err.Error())
			return
		}
	}
	if len(_savingsplansTags) > 0 {
		if err := assignInputField(input, "Tags", _savingsplansTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_savingsplansUpfrontPaymentAmount) > 0 {
		input.UpfrontPaymentAmount = aws.String(_savingsplansUpfrontPaymentAmount)
	}

	if resp, err := client.CreateSavingsPlan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the queued purchase for the specified Savings Plan.
func savingsplans_DeleteQueuedSavingsPlan(cfg aws.Config, client *savingsplans.Client) {
	input := &savingsplans.DeleteQueuedSavingsPlanInput{
		// SavingsPlanId: *string, // Required
	}

	if len(_savingsplansSavingsPlanId) > 0 {
		input.SavingsPlanId = aws.String(_savingsplansSavingsPlanId)
	}

	if resp, err := client.DeleteQueuedSavingsPlan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the rates for a specific, existing Savings Plan.
func savingsplans_DescribeSavingsPlanRates(cfg aws.Config, client *savingsplans.Client) {
	input := &savingsplans.DescribeSavingsPlanRatesInput{
		// SavingsPlanId: *string, // Required
	}

	if len(_savingsplansSavingsPlanId) > 0 {
		input.SavingsPlanId = aws.String(_savingsplansSavingsPlanId)
	}
	if len(_savingsplansFilters) > 0 {
		if err := assignInputField(input, "Filters", _savingsplansFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_savingsplansMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _savingsplansMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_savingsplansNextToken) > 0 {
		input.NextToken = aws.String(_savingsplansNextToken)
	}

	if resp, err := client.DescribeSavingsPlanRates(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the specified Savings Plans.
func savingsplans_DescribeSavingsPlans(cfg aws.Config, client *savingsplans.Client) {
	input := &savingsplans.DescribeSavingsPlansInput{}

	if len(_savingsplansFilters) > 0 {
		if err := assignInputField(input, "Filters", _savingsplansFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_savingsplansMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _savingsplansMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_savingsplansNextToken) > 0 {
		input.NextToken = aws.String(_savingsplansNextToken)
	}
	if len(_savingsplansSavingsPlanArns) > 0 {
		input.SavingsPlanArns = append([]string(nil), _savingsplansSavingsPlanArns...)
	}
	if len(_savingsplansSavingsPlanIds) > 0 {
		input.SavingsPlanIds = append([]string(nil), _savingsplansSavingsPlanIds...)
	}
	if len(_savingsplansStates) > 0 {
		if err := assignInputField(input, "States", _savingsplansStates); err != nil {
			log.Errorf("invalid --states: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeSavingsPlans(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the offering rates for Savings Plans you might want to purchase.
func savingsplans_DescribeSavingsPlansOfferingRates(cfg aws.Config, client *savingsplans.Client) {
	input := &savingsplans.DescribeSavingsPlansOfferingRatesInput{}

	if len(_savingsplansFilters) > 0 {
		if err := assignInputField(input, "Filters", _savingsplansFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_savingsplansMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _savingsplansMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_savingsplansNextToken) > 0 {
		input.NextToken = aws.String(_savingsplansNextToken)
	}
	if len(_savingsplansOperations) > 0 {
		input.Operations = append([]string(nil), _savingsplansOperations...)
	}
	if len(_savingsplansProducts) > 0 {
		if err := assignInputField(input, "Products", _savingsplansProducts); err != nil {
			log.Errorf("invalid --products: %s", err.Error())
			return
		}
	}
	if len(_savingsplansSavingsPlanOfferingIds) > 0 {
		input.SavingsPlanOfferingIds = append([]string(nil), _savingsplansSavingsPlanOfferingIds...)
	}
	if len(_savingsplansSavingsPlanPaymentOptions) > 0 {
		if err := assignInputField(input, "SavingsPlanPaymentOptions", _savingsplansSavingsPlanPaymentOptions); err != nil {
			log.Errorf("invalid --savings-plan-payment-options: %s", err.Error())
			return
		}
	}
	if len(_savingsplansSavingsPlanTypes) > 0 {
		if err := assignInputField(input, "SavingsPlanTypes", _savingsplansSavingsPlanTypes); err != nil {
			log.Errorf("invalid --savings-plan-types: %s", err.Error())
			return
		}
	}
	if len(_savingsplansServiceCodes) > 0 {
		if err := assignInputField(input, "ServiceCodes", _savingsplansServiceCodes[0]); err != nil {
			log.Errorf("invalid --service-codes: %s", err.Error())
			return
		}
	}
	if len(_savingsplansUsageTypes) > 0 {
		input.UsageTypes = append([]string(nil), _savingsplansUsageTypes...)
	}

	if resp, err := client.DescribeSavingsPlansOfferingRates(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the offerings for the specified Savings Plans.
func savingsplans_DescribeSavingsPlansOfferings(cfg aws.Config, client *savingsplans.Client) {
	input := &savingsplans.DescribeSavingsPlansOfferingsInput{}

	if len(_savingsplansCurrencies) > 0 {
		if err := assignInputField(input, "Currencies", _savingsplansCurrencies); err != nil {
			log.Errorf("invalid --currencies: %s", err.Error())
			return
		}
	}
	if len(_savingsplansDescriptions) > 0 {
		input.Descriptions = append([]string(nil), _savingsplansDescriptions...)
	}
	if len(_savingsplansDurations) > 0 {
		if err := assignInputField(input, "Durations", _savingsplansDurations); err != nil {
			log.Errorf("invalid --durations: %s", err.Error())
			return
		}
	}
	if len(_savingsplansFilters) > 0 {
		if err := assignInputField(input, "Filters", _savingsplansFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_savingsplansMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _savingsplansMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_savingsplansNextToken) > 0 {
		input.NextToken = aws.String(_savingsplansNextToken)
	}
	if len(_savingsplansOfferingIds) > 0 {
		input.OfferingIds = append([]string(nil), _savingsplansOfferingIds...)
	}
	if len(_savingsplansOperations) > 0 {
		input.Operations = append([]string(nil), _savingsplansOperations...)
	}
	if len(_savingsplansPaymentOptions) > 0 {
		if err := assignInputField(input, "PaymentOptions", _savingsplansPaymentOptions); err != nil {
			log.Errorf("invalid --payment-options: %s", err.Error())
			return
		}
	}
	if len(_savingsplansPlanTypes) > 0 {
		if err := assignInputField(input, "PlanTypes", _savingsplansPlanTypes); err != nil {
			log.Errorf("invalid --plan-types: %s", err.Error())
			return
		}
	}
	if len(_savingsplansProductType) > 0 {
		if err := assignInputField(input, "ProductType", _savingsplansProductType); err != nil {
			log.Errorf("invalid --product-type: %s", err.Error())
			return
		}
	}
	if len(_savingsplansServiceCodes) > 0 {
		input.ServiceCodes = append([]string(nil), _savingsplansServiceCodes...)
	}
	if len(_savingsplansUsageTypes) > 0 {
		input.UsageTypes = append([]string(nil), _savingsplansUsageTypes...)
	}

	if resp, err := client.DescribeSavingsPlansOfferings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the tags for the specified resource.
func savingsplans_ListTagsForResource(cfg aws.Config, client *savingsplans.Client) {
	input := &savingsplans.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_savingsplansResourceArn) > 0 {
		input.ResourceArn = aws.String(_savingsplansResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the specified Savings Plan.
func savingsplans_ReturnSavingsPlan(cfg aws.Config, client *savingsplans.Client) {
	input := &savingsplans.ReturnSavingsPlanInput{
		// SavingsPlanId: *string, // Required
	}

	if len(_savingsplansSavingsPlanId) > 0 {
		input.SavingsPlanId = aws.String(_savingsplansSavingsPlanId)
	}
	if len(_savingsplansClientToken) > 0 {
		input.ClientToken = aws.String(_savingsplansClientToken)
	}

	if resp, err := client.ReturnSavingsPlan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds the specified tags to the specified resource.
func savingsplans_TagResource(cfg aws.Config, client *savingsplans.Client) {
	input := &savingsplans.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_savingsplansResourceArn) > 0 {
		input.ResourceArn = aws.String(_savingsplansResourceArn)
	}
	if len(_savingsplansTags) > 0 {
		if err := assignInputField(input, "Tags", _savingsplansTags); err != nil {
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

// Removes the specified tags from the specified resource.
func savingsplans_UntagResource(cfg aws.Config, client *savingsplans.Client) {
	input := &savingsplans.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_savingsplansResourceArn) > 0 {
		input.ResourceArn = aws.String(_savingsplansResourceArn)
	}
	if len(_savingsplansTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _savingsplansTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_savingsplansCmd)
	_savingsplansCmd.Flags().SortFlags = false

	_savingsplansCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_savingsplansCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_savingsplansCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_savingsplansCmd.Flags().StringVarP(&_savingsplansClientToken, "client-token", "", "", "Client Token")
	_savingsplansCmd.Flags().StringVarP(&_savingsplansCommitment, "commitment", "", "", "Commitment")
	_savingsplansCmd.Flags().StringVarP(&_savingsplansCurrencies, "currencies", "", "", "Currencies")
	_savingsplansCmd.Flags().StringSliceVarP(&_savingsplansDescriptions, "descriptions", "", nil, "Descriptions")
	_savingsplansCmd.Flags().StringVarP(&_savingsplansDurations, "durations", "", "", "Durations")
	_savingsplansCmd.Flags().StringVarP(&_savingsplansFilters, "filters", "", "", "Filters")
	_savingsplansCmd.Flags().StringVarP(&_savingsplansMaxResults, "max-results", "", "", "Max Results")
	_savingsplansCmd.Flags().StringVarP(&_savingsplansNextToken, "next-token", "", "", "Next Token")
	_savingsplansCmd.Flags().StringSliceVarP(&_savingsplansOfferingIds, "offering-ids", "", nil, "Offering Ids")
	_savingsplansCmd.Flags().StringSliceVarP(&_savingsplansOperations, "operations", "", nil, "Operations")
	_savingsplansCmd.Flags().StringVarP(&_savingsplansPaymentOptions, "payment-options", "", "", "Payment Options")
	_savingsplansCmd.Flags().StringVarP(&_savingsplansPlanTypes, "plan-types", "", "", "Plan Types")
	_savingsplansCmd.Flags().StringVarP(&_savingsplansProductType, "product-type", "", "", "Product Type")
	_savingsplansCmd.Flags().StringVarP(&_savingsplansProducts, "products", "", "", "Products")
	_savingsplansCmd.Flags().StringVarP(&_savingsplansPurchaseTime, "purchase-time", "", "", "Purchase Time")
	_savingsplansCmd.Flags().StringVarP(&_savingsplansResourceArn, "resource-arn", "", "", "Resource ARN")
	_savingsplansCmd.Flags().StringSliceVarP(&_savingsplansSavingsPlanArns, "savings-plan-arns", "", nil, "Savings Plan Arns")
	_savingsplansCmd.Flags().StringVarP(&_savingsplansSavingsPlanId, "savings-plan-id", "", "", "Savings Plan ID")
	_savingsplansCmd.Flags().StringSliceVarP(&_savingsplansSavingsPlanIds, "savings-plan-ids", "", nil, "Savings Plan Ids")
	_savingsplansCmd.Flags().StringVarP(&_savingsplansSavingsPlanOfferingId, "savings-plan-offering-id", "", "", "Savings Plan Offering ID")
	_savingsplansCmd.Flags().StringSliceVarP(&_savingsplansSavingsPlanOfferingIds, "savings-plan-offering-ids", "", nil, "Savings Plan Offering Ids")
	_savingsplansCmd.Flags().StringVarP(&_savingsplansSavingsPlanPaymentOptions, "savings-plan-payment-options", "", "", "Savings Plan Payment Options")
	_savingsplansCmd.Flags().StringVarP(&_savingsplansSavingsPlanTypes, "savings-plan-types", "", "", "Savings Plan Types")
	_savingsplansCmd.Flags().StringSliceVarP(&_savingsplansServiceCodes, "service-codes", "", nil, "Service Codes")
	_savingsplansCmd.Flags().StringVarP(&_savingsplansStates, "states", "", "", "States")
	_savingsplansCmd.Flags().StringSliceVarP(&_savingsplansTagKeys, "tag-keys", "", nil, "Tag Keys")
	_savingsplansCmd.Flags().StringVarP(&_savingsplansTags, "tags", "", "", "Tags")
	_savingsplansCmd.Flags().StringVarP(&_savingsplansUpfrontPaymentAmount, "upfront-payment-amount", "", "", "Upfront Payment Amount")
	_savingsplansCmd.Flags().StringSliceVarP(&_savingsplansUsageTypes, "usage-types", "", nil, "Usage Types")

	_savingsplansCmd.Flags().BoolVarP(&_savingsplansCreateSavingsPlan, "create-savings-plan", "", false, "Create Savings Plan")
	_savingsplansCmd.Flags().BoolVarP(&_savingsplansDeleteQueuedSavingsPlan, "delete-queued-savings-plan", "", false, "Delete Queued Savings Plan")
	_savingsplansCmd.Flags().BoolVarP(&_savingsplansDescribeSavingsPlanRates, "describe-savings-plan-rates", "", false, "Describe Savings Plan Rates")
	_savingsplansCmd.Flags().BoolVarP(&_savingsplansDescribeSavingsPlans, "describe-savings-plans", "", false, "Describe Savings Plans")
	_savingsplansCmd.Flags().BoolVarP(&_savingsplansDescribeSavingsPlansOfferingRates, "describe-savings-plans-offering-rates", "", false, "Describe Savings Plans Offering Rates")
	_savingsplansCmd.Flags().BoolVarP(&_savingsplansDescribeSavingsPlansOfferings, "describe-savings-plans-offerings", "", false, "Describe Savings Plans Offerings")
	_savingsplansCmd.Flags().BoolVarP(&_savingsplansListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_savingsplansCmd.Flags().BoolVarP(&_savingsplansReturnSavingsPlan, "return-savings-plan", "", false, "Return Savings Plan")
	_savingsplansCmd.Flags().BoolVarP(&_savingsplansTagResource, "tag-resource", "", false, "Tag Resource")
	_savingsplansCmd.Flags().BoolVarP(&_savingsplansUntagResource, "untag-resource", "", false, "Untag Resource")

}
