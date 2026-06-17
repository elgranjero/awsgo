package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/bcmpricingcalculator"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// bcmpricingcalculatorCmd represents the bcmpricingcalculator command
var _bcmpricingcalculatorCmd = &cobra.Command{
	Use:   "bcmpricingcalculator",
	Short: "AWS bcmpricingcalculator CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := bcmpricingcalculator.NewFromConfig(cfg)
		if _bcmpricingcalculatorBatchCreateBillScenarioCommitmentModification {
			bcmpricingcalculator_BatchCreateBillScenarioCommitmentModification(cfg, client)
			return
		}
		if _bcmpricingcalculatorBatchCreateBillScenarioUsageModification {
			bcmpricingcalculator_BatchCreateBillScenarioUsageModification(cfg, client)
			return
		}
		if _bcmpricingcalculatorBatchCreateWorkloadEstimateUsage {
			bcmpricingcalculator_BatchCreateWorkloadEstimateUsage(cfg, client)
			return
		}
		if _bcmpricingcalculatorBatchDeleteBillScenarioCommitmentModification {
			bcmpricingcalculator_BatchDeleteBillScenarioCommitmentModification(cfg, client)
			return
		}
		if _bcmpricingcalculatorBatchDeleteBillScenarioUsageModification {
			bcmpricingcalculator_BatchDeleteBillScenarioUsageModification(cfg, client)
			return
		}
		if _bcmpricingcalculatorBatchDeleteWorkloadEstimateUsage {
			bcmpricingcalculator_BatchDeleteWorkloadEstimateUsage(cfg, client)
			return
		}
		if _bcmpricingcalculatorBatchUpdateBillScenarioCommitmentModification {
			bcmpricingcalculator_BatchUpdateBillScenarioCommitmentModification(cfg, client)
			return
		}
		if _bcmpricingcalculatorBatchUpdateBillScenarioUsageModification {
			bcmpricingcalculator_BatchUpdateBillScenarioUsageModification(cfg, client)
			return
		}
		if _bcmpricingcalculatorBatchUpdateWorkloadEstimateUsage {
			bcmpricingcalculator_BatchUpdateWorkloadEstimateUsage(cfg, client)
			return
		}
		if _bcmpricingcalculatorCreateBillEstimate {
			bcmpricingcalculator_CreateBillEstimate(cfg, client)
			return
		}
		if _bcmpricingcalculatorCreateBillScenario {
			bcmpricingcalculator_CreateBillScenario(cfg, client)
			return
		}
		if _bcmpricingcalculatorCreateWorkloadEstimate {
			bcmpricingcalculator_CreateWorkloadEstimate(cfg, client)
			return
		}
		if _bcmpricingcalculatorDeleteBillEstimate {
			bcmpricingcalculator_DeleteBillEstimate(cfg, client)
			return
		}
		if _bcmpricingcalculatorDeleteBillScenario {
			bcmpricingcalculator_DeleteBillScenario(cfg, client)
			return
		}
		if _bcmpricingcalculatorDeleteWorkloadEstimate {
			bcmpricingcalculator_DeleteWorkloadEstimate(cfg, client)
			return
		}
		if _bcmpricingcalculatorGetBillEstimate {
			bcmpricingcalculator_GetBillEstimate(cfg, client)
			return
		}
		if _bcmpricingcalculatorGetBillScenario {
			bcmpricingcalculator_GetBillScenario(cfg, client)
			return
		}
		if _bcmpricingcalculatorGetPreferences {
			bcmpricingcalculator_GetPreferences(cfg, client)
			return
		}
		if _bcmpricingcalculatorGetWorkloadEstimate {
			bcmpricingcalculator_GetWorkloadEstimate(cfg, client)
			return
		}
		if _bcmpricingcalculatorListBillEstimateCommitments {
			bcmpricingcalculator_ListBillEstimateCommitments(cfg, client)
			return
		}
		if _bcmpricingcalculatorListBillEstimateInputCommitmentModifications {
			bcmpricingcalculator_ListBillEstimateInputCommitmentModifications(cfg, client)
			return
		}
		if _bcmpricingcalculatorListBillEstimateInputUsageModifications {
			bcmpricingcalculator_ListBillEstimateInputUsageModifications(cfg, client)
			return
		}
		if _bcmpricingcalculatorListBillEstimateLineItems {
			bcmpricingcalculator_ListBillEstimateLineItems(cfg, client)
			return
		}
		if _bcmpricingcalculatorListBillEstimates {
			bcmpricingcalculator_ListBillEstimates(cfg, client)
			return
		}
		if _bcmpricingcalculatorListBillScenarioCommitmentModifications {
			bcmpricingcalculator_ListBillScenarioCommitmentModifications(cfg, client)
			return
		}
		if _bcmpricingcalculatorListBillScenarioUsageModifications {
			bcmpricingcalculator_ListBillScenarioUsageModifications(cfg, client)
			return
		}
		if _bcmpricingcalculatorListBillScenarios {
			bcmpricingcalculator_ListBillScenarios(cfg, client)
			return
		}
		if _bcmpricingcalculatorListTagsForResource {
			bcmpricingcalculator_ListTagsForResource(cfg, client)
			return
		}
		if _bcmpricingcalculatorListWorkloadEstimateUsage {
			bcmpricingcalculator_ListWorkloadEstimateUsage(cfg, client)
			return
		}
		if _bcmpricingcalculatorListWorkloadEstimates {
			bcmpricingcalculator_ListWorkloadEstimates(cfg, client)
			return
		}
		if _bcmpricingcalculatorTagResource {
			bcmpricingcalculator_TagResource(cfg, client)
			return
		}
		if _bcmpricingcalculatorUntagResource {
			bcmpricingcalculator_UntagResource(cfg, client)
			return
		}
		if _bcmpricingcalculatorUpdateBillEstimate {
			bcmpricingcalculator_UpdateBillEstimate(cfg, client)
			return
		}
		if _bcmpricingcalculatorUpdateBillScenario {
			bcmpricingcalculator_UpdateBillScenario(cfg, client)
			return
		}
		if _bcmpricingcalculatorUpdatePreferences {
			bcmpricingcalculator_UpdatePreferences(cfg, client)
			return
		}
		if _bcmpricingcalculatorUpdateWorkloadEstimate {
			bcmpricingcalculator_UpdateWorkloadEstimate(cfg, client)
			return
		}

	},
}

var (
	_bcmpricingcalculatorBatchCreateBillScenarioCommitmentModification bool
	_bcmpricingcalculatorBatchCreateBillScenarioUsageModification      bool
	_bcmpricingcalculatorBatchCreateWorkloadEstimateUsage              bool
	_bcmpricingcalculatorBatchDeleteBillScenarioCommitmentModification bool
	_bcmpricingcalculatorBatchDeleteBillScenarioUsageModification      bool
	_bcmpricingcalculatorBatchDeleteWorkloadEstimateUsage              bool
	_bcmpricingcalculatorBatchUpdateBillScenarioCommitmentModification bool
	_bcmpricingcalculatorBatchUpdateBillScenarioUsageModification      bool
	_bcmpricingcalculatorBatchUpdateWorkloadEstimateUsage              bool
	_bcmpricingcalculatorCreateBillEstimate                            bool
	_bcmpricingcalculatorCreateBillScenario                            bool
	_bcmpricingcalculatorCreateWorkloadEstimate                        bool
	_bcmpricingcalculatorDeleteBillEstimate                            bool
	_bcmpricingcalculatorDeleteBillScenario                            bool
	_bcmpricingcalculatorDeleteWorkloadEstimate                        bool
	_bcmpricingcalculatorGetBillEstimate                               bool
	_bcmpricingcalculatorGetBillScenario                               bool
	_bcmpricingcalculatorGetPreferences                                bool
	_bcmpricingcalculatorGetWorkloadEstimate                           bool
	_bcmpricingcalculatorListBillEstimateCommitments                   bool
	_bcmpricingcalculatorListBillEstimateInputCommitmentModifications  bool
	_bcmpricingcalculatorListBillEstimateInputUsageModifications       bool
	_bcmpricingcalculatorListBillEstimateLineItems                     bool
	_bcmpricingcalculatorListBillEstimates                             bool
	_bcmpricingcalculatorListBillScenarioCommitmentModifications       bool
	_bcmpricingcalculatorListBillScenarioUsageModifications            bool
	_bcmpricingcalculatorListBillScenarios                             bool
	_bcmpricingcalculatorListTagsForResource                           bool
	_bcmpricingcalculatorListWorkloadEstimateUsage                     bool
	_bcmpricingcalculatorListWorkloadEstimates                         bool
	_bcmpricingcalculatorTagResource                                   bool
	_bcmpricingcalculatorUntagResource                                 bool
	_bcmpricingcalculatorUpdateBillEstimate                            bool
	_bcmpricingcalculatorUpdateBillScenario                            bool
	_bcmpricingcalculatorUpdatePreferences                             bool
	_bcmpricingcalculatorUpdateWorkloadEstimate                        bool

	_bcmpricingcalculatorArn                                   string
	_bcmpricingcalculatorBillEstimateId                        string
	_bcmpricingcalculatorBillScenarioId                        string
	_bcmpricingcalculatorClientToken                           string
	_bcmpricingcalculatorCommitmentModifications               string
	_bcmpricingcalculatorCostCategoryGroupSharingPreferenceArn string
	_bcmpricingcalculatorCreatedAtFilter                       string
	_bcmpricingcalculatorExpiresAt                             string
	_bcmpricingcalculatorExpiresAtFilter                       string
	_bcmpricingcalculatorFilters                               string
	_bcmpricingcalculatorGroupSharingPreference                string
	_bcmpricingcalculatorIdentifier                            string
	_bcmpricingcalculatorIds                                   []string
	_bcmpricingcalculatorManagementAccountRateTypeSelections   string
	_bcmpricingcalculatorMaxResults                            string
	_bcmpricingcalculatorMemberAccountRateTypeSelections       string
	_bcmpricingcalculatorName                                  string
	_bcmpricingcalculatorNextToken                             string
	_bcmpricingcalculatorRateType                              string
	_bcmpricingcalculatorStandaloneAccountRateTypeSelections   string
	_bcmpricingcalculatorTagKeys                               []string
	_bcmpricingcalculatorTags                                  string
	_bcmpricingcalculatorUsage                                 string
	_bcmpricingcalculatorUsageModifications                    string
	_bcmpricingcalculatorWorkloadEstimateId                    string
)

// Create Compute Savings Plans, EC2 Instance Savings Plans, or EC2 Reserved
// Instances commitments that you want to model in a Bill Scenario.
//
// The BatchCreateBillScenarioCommitmentModification operation doesn't have its
// own IAM permission. To authorize this operation for Amazon Web Services
// principals, include the permission
// bcm-pricing-calculator:CreateBillScenarioCommitmentModification in your policies.
func bcmpricingcalculator_BatchCreateBillScenarioCommitmentModification(cfg aws.Config, client *bcmpricingcalculator.Client) {
	input := &bcmpricingcalculator.BatchCreateBillScenarioCommitmentModificationInput{
		// BillScenarioId: *string, // Required
		// CommitmentModifications: []types.BatchCreateBillScenarioCommitmentModificationEntry, // Required
	}

	if len(_bcmpricingcalculatorBillScenarioId) > 0 {
		input.BillScenarioId = aws.String(_bcmpricingcalculatorBillScenarioId)
	}
	if len(_bcmpricingcalculatorCommitmentModifications) > 0 {
		if err := assignInputField(input, "CommitmentModifications", _bcmpricingcalculatorCommitmentModifications); err != nil {
			log.Errorf("invalid --commitment-modifications: %s", err.Error())
			return
		}
	}
	if len(_bcmpricingcalculatorClientToken) > 0 {
		input.ClientToken = aws.String(_bcmpricingcalculatorClientToken)
	}

	if resp, err := client.BatchCreateBillScenarioCommitmentModification(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create Amazon Web Services service usage that you want to model in a Bill
// Scenario.
//
// The BatchCreateBillScenarioUsageModification operation doesn't have its own IAM
// permission. To authorize this operation for Amazon Web Services principals,
// include the permission
// bcm-pricing-calculator:CreateBillScenarioUsageModification in your policies.
func bcmpricingcalculator_BatchCreateBillScenarioUsageModification(cfg aws.Config, client *bcmpricingcalculator.Client) {
	input := &bcmpricingcalculator.BatchCreateBillScenarioUsageModificationInput{
		// BillScenarioId: *string, // Required
		// UsageModifications: []types.BatchCreateBillScenarioUsageModificationEntry, // Required
	}

	if len(_bcmpricingcalculatorBillScenarioId) > 0 {
		input.BillScenarioId = aws.String(_bcmpricingcalculatorBillScenarioId)
	}
	if len(_bcmpricingcalculatorUsageModifications) > 0 {
		if err := assignInputField(input, "UsageModifications", _bcmpricingcalculatorUsageModifications); err != nil {
			log.Errorf("invalid --usage-modifications: %s", err.Error())
			return
		}
	}
	if len(_bcmpricingcalculatorClientToken) > 0 {
		input.ClientToken = aws.String(_bcmpricingcalculatorClientToken)
	}

	if resp, err := client.BatchCreateBillScenarioUsageModification(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create Amazon Web Services service usage that you want to model in a Workload
// Estimate.
//
// The BatchCreateWorkloadEstimateUsage operation doesn't have its own IAM
// permission. To authorize this operation for Amazon Web Services principals,
// include the permission bcm-pricing-calculator:CreateWorkloadEstimateUsage in
// your policies.
func bcmpricingcalculator_BatchCreateWorkloadEstimateUsage(cfg aws.Config, client *bcmpricingcalculator.Client) {
	input := &bcmpricingcalculator.BatchCreateWorkloadEstimateUsageInput{
		// Usage: []types.BatchCreateWorkloadEstimateUsageEntry, // Required
		// WorkloadEstimateId: *string, // Required
	}

	if len(_bcmpricingcalculatorUsage) > 0 {
		if err := assignInputField(input, "Usage", _bcmpricingcalculatorUsage); err != nil {
			log.Errorf("invalid --usage: %s", err.Error())
			return
		}
	}
	if len(_bcmpricingcalculatorWorkloadEstimateId) > 0 {
		input.WorkloadEstimateId = aws.String(_bcmpricingcalculatorWorkloadEstimateId)
	}
	if len(_bcmpricingcalculatorClientToken) > 0 {
		input.ClientToken = aws.String(_bcmpricingcalculatorClientToken)
	}

	if resp, err := client.BatchCreateWorkloadEstimateUsage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete commitment that you have created in a Bill Scenario. You can only
// delete a commitment that you had added and cannot model deletion (or removal) of
// a existing commitment. If you want model deletion of an existing commitment, see
// the negate [BillScenarioCommitmentModificationAction]of [BatchCreateBillScenarioCommitmentModification] operation.
//
// The BatchDeleteBillScenarioCommitmentModification operation doesn't have its
// own IAM permission. To authorize this operation for Amazon Web Services
// principals, include the permission
// bcm-pricing-calculator:DeleteBillScenarioCommitmentModification in your policies.
//
// [BillScenarioCommitmentModificationAction]: https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_AWSBCMPricingCalculator_BillScenarioCommitmentModificationAction.html
// [BatchCreateBillScenarioCommitmentModification]: https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_AWSBCMPricingCalculator_BatchCreateBillScenarioUsageModification.html
func bcmpricingcalculator_BatchDeleteBillScenarioCommitmentModification(cfg aws.Config, client *bcmpricingcalculator.Client) {
	input := &bcmpricingcalculator.BatchDeleteBillScenarioCommitmentModificationInput{
		// BillScenarioId: *string, // Required
		// Ids: []string, // Required
	}

	if len(_bcmpricingcalculatorBillScenarioId) > 0 {
		input.BillScenarioId = aws.String(_bcmpricingcalculatorBillScenarioId)
	}
	if len(_bcmpricingcalculatorIds) > 0 {
		input.Ids = append([]string(nil), _bcmpricingcalculatorIds...)
	}

	if resp, err := client.BatchDeleteBillScenarioCommitmentModification(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete usage that you have created in a Bill Scenario. You can only delete
// usage that you had added and cannot model deletion (or removal) of a existing
// usage. If you want model removal of an existing usage, see [BatchUpdateBillScenarioUsageModification].
//
// The BatchDeleteBillScenarioUsageModification operation doesn't have its own IAM
// permission. To authorize this operation for Amazon Web Services principals,
// include the permission
// bcm-pricing-calculator:DeleteBillScenarioUsageModification in your policies.
//
// [BatchUpdateBillScenarioUsageModification]: https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_AWSBCMPricingCalculator_BatchUpdateBillScenarioUsageModification.html
func bcmpricingcalculator_BatchDeleteBillScenarioUsageModification(cfg aws.Config, client *bcmpricingcalculator.Client) {
	input := &bcmpricingcalculator.BatchDeleteBillScenarioUsageModificationInput{
		// BillScenarioId: *string, // Required
		// Ids: []string, // Required
	}

	if len(_bcmpricingcalculatorBillScenarioId) > 0 {
		input.BillScenarioId = aws.String(_bcmpricingcalculatorBillScenarioId)
	}
	if len(_bcmpricingcalculatorIds) > 0 {
		input.Ids = append([]string(nil), _bcmpricingcalculatorIds...)
	}

	if resp, err := client.BatchDeleteBillScenarioUsageModification(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete usage that you have created in a Workload estimate. You can only delete
// usage that you had added and cannot model deletion (or removal) of a existing
// usage. If you want model removal of an existing usage, see [BatchUpdateWorkloadEstimateUsage].
//
// The BatchDeleteWorkloadEstimateUsage operation doesn't have its own IAM
// permission. To authorize this operation for Amazon Web Services principals,
// include the permission bcm-pricing-calculator:DeleteWorkloadEstimateUsage in
// your policies.
//
// [BatchUpdateWorkloadEstimateUsage]: https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_AWSBCMPricingCalculator_BatchUpdateWorkloadEstimateUsage.html
func bcmpricingcalculator_BatchDeleteWorkloadEstimateUsage(cfg aws.Config, client *bcmpricingcalculator.Client) {
	input := &bcmpricingcalculator.BatchDeleteWorkloadEstimateUsageInput{
		// Ids: []string, // Required
		// WorkloadEstimateId: *string, // Required
	}

	if len(_bcmpricingcalculatorIds) > 0 {
		input.Ids = append([]string(nil), _bcmpricingcalculatorIds...)
	}
	if len(_bcmpricingcalculatorWorkloadEstimateId) > 0 {
		input.WorkloadEstimateId = aws.String(_bcmpricingcalculatorWorkloadEstimateId)
	}

	if resp, err := client.BatchDeleteWorkloadEstimateUsage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update a newly added or existing commitment. You can update the commitment
// group based on a commitment ID and a Bill scenario ID.
//
// The BatchUpdateBillScenarioCommitmentModification operation doesn't have its
// own IAM permission. To authorize this operation for Amazon Web Services
// principals, include the permission
// bcm-pricing-calculator:UpdateBillScenarioCommitmentModification in your policies.
func bcmpricingcalculator_BatchUpdateBillScenarioCommitmentModification(cfg aws.Config, client *bcmpricingcalculator.Client) {
	input := &bcmpricingcalculator.BatchUpdateBillScenarioCommitmentModificationInput{
		// BillScenarioId: *string, // Required
		// CommitmentModifications: []types.BatchUpdateBillScenarioCommitmentModificationEntry, // Required
	}

	if len(_bcmpricingcalculatorBillScenarioId) > 0 {
		input.BillScenarioId = aws.String(_bcmpricingcalculatorBillScenarioId)
	}
	if len(_bcmpricingcalculatorCommitmentModifications) > 0 {
		if err := assignInputField(input, "CommitmentModifications", _bcmpricingcalculatorCommitmentModifications); err != nil {
			log.Errorf("invalid --commitment-modifications: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchUpdateBillScenarioCommitmentModification(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update a newly added or existing usage lines. You can update the usage
// amounts, usage hour, and usage group based on a usage ID and a Bill scenario ID.
//
// The BatchUpdateBillScenarioUsageModification operation doesn't have its own IAM
// permission. To authorize this operation for Amazon Web Services principals,
// include the permission
// bcm-pricing-calculator:UpdateBillScenarioUsageModification in your policies.
func bcmpricingcalculator_BatchUpdateBillScenarioUsageModification(cfg aws.Config, client *bcmpricingcalculator.Client) {
	input := &bcmpricingcalculator.BatchUpdateBillScenarioUsageModificationInput{
		// BillScenarioId: *string, // Required
		// UsageModifications: []types.BatchUpdateBillScenarioUsageModificationEntry, // Required
	}

	if len(_bcmpricingcalculatorBillScenarioId) > 0 {
		input.BillScenarioId = aws.String(_bcmpricingcalculatorBillScenarioId)
	}
	if len(_bcmpricingcalculatorUsageModifications) > 0 {
		if err := assignInputField(input, "UsageModifications", _bcmpricingcalculatorUsageModifications); err != nil {
			log.Errorf("invalid --usage-modifications: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchUpdateBillScenarioUsageModification(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update a newly added or existing usage lines. You can update the usage amounts
// and usage group based on a usage ID and a Workload estimate ID.
//
// The BatchUpdateWorkloadEstimateUsage operation doesn't have its own IAM
// permission. To authorize this operation for Amazon Web Services principals,
// include the permission bcm-pricing-calculator:UpdateWorkloadEstimateUsage in
// your policies.
func bcmpricingcalculator_BatchUpdateWorkloadEstimateUsage(cfg aws.Config, client *bcmpricingcalculator.Client) {
	input := &bcmpricingcalculator.BatchUpdateWorkloadEstimateUsageInput{
		// Usage: []types.BatchUpdateWorkloadEstimateUsageEntry, // Required
		// WorkloadEstimateId: *string, // Required
	}

	if len(_bcmpricingcalculatorUsage) > 0 {
		if err := assignInputField(input, "Usage", _bcmpricingcalculatorUsage); err != nil {
			log.Errorf("invalid --usage: %s", err.Error())
			return
		}
	}
	if len(_bcmpricingcalculatorWorkloadEstimateId) > 0 {
		input.WorkloadEstimateId = aws.String(_bcmpricingcalculatorWorkloadEstimateId)
	}

	if resp, err := client.BatchUpdateWorkloadEstimateUsage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a Bill estimate from a Bill scenario. In the Bill scenario you can
// model usage addition, usage changes, and usage removal. You can also model
// commitment addition and commitment removal. After all changes in a Bill scenario
// is made satisfactorily, you can call this API with a Bill scenario ID to
// generate the Bill estimate. Bill estimate calculates the pre-tax cost for your
// consolidated billing family, incorporating all modeled usage and commitments
// alongside existing usage and commitments from your most recent completed
// anniversary bill, with any applicable discounts applied.
func bcmpricingcalculator_CreateBillEstimate(cfg aws.Config, client *bcmpricingcalculator.Client) {
	input := &bcmpricingcalculator.CreateBillEstimateInput{
		// BillScenarioId: *string, // Required
		// Name: *string, // Required
	}

	if len(_bcmpricingcalculatorBillScenarioId) > 0 {
		input.BillScenarioId = aws.String(_bcmpricingcalculatorBillScenarioId)
	}
	if len(_bcmpricingcalculatorName) > 0 {
		input.Name = aws.String(_bcmpricingcalculatorName)
	}
	if len(_bcmpricingcalculatorClientToken) > 0 {
		input.ClientToken = aws.String(_bcmpricingcalculatorClientToken)
	}
	if len(_bcmpricingcalculatorTags) > 0 {
		if err := assignInputField(input, "Tags", _bcmpricingcalculatorTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateBillEstimate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new bill scenario to model potential changes to Amazon Web Services
// usage and costs.
func bcmpricingcalculator_CreateBillScenario(cfg aws.Config, client *bcmpricingcalculator.Client) {
	input := &bcmpricingcalculator.CreateBillScenarioInput{
		// Name: *string, // Required
	}

	if len(_bcmpricingcalculatorName) > 0 {
		input.Name = aws.String(_bcmpricingcalculatorName)
	}
	if len(_bcmpricingcalculatorClientToken) > 0 {
		input.ClientToken = aws.String(_bcmpricingcalculatorClientToken)
	}
	if len(_bcmpricingcalculatorCostCategoryGroupSharingPreferenceArn) > 0 {
		input.CostCategoryGroupSharingPreferenceArn = aws.String(_bcmpricingcalculatorCostCategoryGroupSharingPreferenceArn)
	}
	if len(_bcmpricingcalculatorGroupSharingPreference) > 0 {
		if err := assignInputField(input, "GroupSharingPreference", _bcmpricingcalculatorGroupSharingPreference); err != nil {
			log.Errorf("invalid --group-sharing-preference: %s", err.Error())
			return
		}
	}
	if len(_bcmpricingcalculatorTags) > 0 {
		if err := assignInputField(input, "Tags", _bcmpricingcalculatorTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateBillScenario(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new workload estimate to model costs for a specific workload.
func bcmpricingcalculator_CreateWorkloadEstimate(cfg aws.Config, client *bcmpricingcalculator.Client) {
	input := &bcmpricingcalculator.CreateWorkloadEstimateInput{
		// Name: *string, // Required
	}

	if len(_bcmpricingcalculatorName) > 0 {
		input.Name = aws.String(_bcmpricingcalculatorName)
	}
	if len(_bcmpricingcalculatorClientToken) > 0 {
		input.ClientToken = aws.String(_bcmpricingcalculatorClientToken)
	}
	if len(_bcmpricingcalculatorRateType) > 0 {
		if err := assignInputField(input, "RateType", _bcmpricingcalculatorRateType); err != nil {
			log.Errorf("invalid --rate-type: %s", err.Error())
			return
		}
	}
	if len(_bcmpricingcalculatorTags) > 0 {
		if err := assignInputField(input, "Tags", _bcmpricingcalculatorTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateWorkloadEstimate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing bill estimate.
func bcmpricingcalculator_DeleteBillEstimate(cfg aws.Config, client *bcmpricingcalculator.Client) {
	input := &bcmpricingcalculator.DeleteBillEstimateInput{
		// Identifier: *string, // Required
	}

	if len(_bcmpricingcalculatorIdentifier) > 0 {
		input.Identifier = aws.String(_bcmpricingcalculatorIdentifier)
	}

	if resp, err := client.DeleteBillEstimate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing bill scenario.
func bcmpricingcalculator_DeleteBillScenario(cfg aws.Config, client *bcmpricingcalculator.Client) {
	input := &bcmpricingcalculator.DeleteBillScenarioInput{
		// Identifier: *string, // Required
	}

	if len(_bcmpricingcalculatorIdentifier) > 0 {
		input.Identifier = aws.String(_bcmpricingcalculatorIdentifier)
	}

	if resp, err := client.DeleteBillScenario(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing workload estimate.
func bcmpricingcalculator_DeleteWorkloadEstimate(cfg aws.Config, client *bcmpricingcalculator.Client) {
	input := &bcmpricingcalculator.DeleteWorkloadEstimateInput{
		// Identifier: *string, // Required
	}

	if len(_bcmpricingcalculatorIdentifier) > 0 {
		input.Identifier = aws.String(_bcmpricingcalculatorIdentifier)
	}

	if resp, err := client.DeleteWorkloadEstimate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves details of a specific bill estimate.
func bcmpricingcalculator_GetBillEstimate(cfg aws.Config, client *bcmpricingcalculator.Client) {
	input := &bcmpricingcalculator.GetBillEstimateInput{
		// Identifier: *string, // Required
	}

	if len(_bcmpricingcalculatorIdentifier) > 0 {
		input.Identifier = aws.String(_bcmpricingcalculatorIdentifier)
	}

	if resp, err := client.GetBillEstimate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves details of a specific bill scenario.
func bcmpricingcalculator_GetBillScenario(cfg aws.Config, client *bcmpricingcalculator.Client) {
	input := &bcmpricingcalculator.GetBillScenarioInput{
		// Identifier: *string, // Required
	}

	if len(_bcmpricingcalculatorIdentifier) > 0 {
		input.Identifier = aws.String(_bcmpricingcalculatorIdentifier)
	}

	if resp, err := client.GetBillScenario(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the current preferences for Pricing Calculator.
func bcmpricingcalculator_GetPreferences(cfg aws.Config, client *bcmpricingcalculator.Client) {
	input := &bcmpricingcalculator.GetPreferencesInput{}

	if resp, err := client.GetPreferences(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves details of a specific workload estimate.
func bcmpricingcalculator_GetWorkloadEstimate(cfg aws.Config, client *bcmpricingcalculator.Client) {
	input := &bcmpricingcalculator.GetWorkloadEstimateInput{
		// Identifier: *string, // Required
	}

	if len(_bcmpricingcalculatorIdentifier) > 0 {
		input.Identifier = aws.String(_bcmpricingcalculatorIdentifier)
	}

	if resp, err := client.GetWorkloadEstimate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the commitments associated with a bill estimate.
func bcmpricingcalculator_ListBillEstimateCommitments(cfg aws.Config, client *bcmpricingcalculator.Client) {
	input := &bcmpricingcalculator.ListBillEstimateCommitmentsInput{
		// BillEstimateId: *string, // Required
	}

	if len(_bcmpricingcalculatorBillEstimateId) > 0 {
		input.BillEstimateId = aws.String(_bcmpricingcalculatorBillEstimateId)
	}
	if len(_bcmpricingcalculatorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bcmpricingcalculatorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bcmpricingcalculatorNextToken) > 0 {
		input.NextToken = aws.String(_bcmpricingcalculatorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListBillEstimateCommitments(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bcmpricingcalculator.ListBillEstimateCommitmentsOutput
	p := bcmpricingcalculator.NewListBillEstimateCommitmentsPaginator(client, input)
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

// Lists the input commitment modifications associated with a bill estimate.
func bcmpricingcalculator_ListBillEstimateInputCommitmentModifications(cfg aws.Config, client *bcmpricingcalculator.Client) {
	input := &bcmpricingcalculator.ListBillEstimateInputCommitmentModificationsInput{
		// BillEstimateId: *string, // Required
	}

	if len(_bcmpricingcalculatorBillEstimateId) > 0 {
		input.BillEstimateId = aws.String(_bcmpricingcalculatorBillEstimateId)
	}
	if len(_bcmpricingcalculatorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bcmpricingcalculatorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bcmpricingcalculatorNextToken) > 0 {
		input.NextToken = aws.String(_bcmpricingcalculatorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListBillEstimateInputCommitmentModifications(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bcmpricingcalculator.ListBillEstimateInputCommitmentModificationsOutput
	p := bcmpricingcalculator.NewListBillEstimateInputCommitmentModificationsPaginator(client, input)
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

// Lists the input usage modifications associated with a bill estimate.
func bcmpricingcalculator_ListBillEstimateInputUsageModifications(cfg aws.Config, client *bcmpricingcalculator.Client) {
	input := &bcmpricingcalculator.ListBillEstimateInputUsageModificationsInput{
		// BillEstimateId: *string, // Required
	}

	if len(_bcmpricingcalculatorBillEstimateId) > 0 {
		input.BillEstimateId = aws.String(_bcmpricingcalculatorBillEstimateId)
	}
	if len(_bcmpricingcalculatorFilters) > 0 {
		if err := assignInputField(input, "Filters", _bcmpricingcalculatorFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_bcmpricingcalculatorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bcmpricingcalculatorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bcmpricingcalculatorNextToken) > 0 {
		input.NextToken = aws.String(_bcmpricingcalculatorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListBillEstimateInputUsageModifications(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bcmpricingcalculator.ListBillEstimateInputUsageModificationsOutput
	p := bcmpricingcalculator.NewListBillEstimateInputUsageModificationsPaginator(client, input)
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

// Lists the line items associated with a bill estimate.
func bcmpricingcalculator_ListBillEstimateLineItems(cfg aws.Config, client *bcmpricingcalculator.Client) {
	input := &bcmpricingcalculator.ListBillEstimateLineItemsInput{
		// BillEstimateId: *string, // Required
	}

	if len(_bcmpricingcalculatorBillEstimateId) > 0 {
		input.BillEstimateId = aws.String(_bcmpricingcalculatorBillEstimateId)
	}
	if len(_bcmpricingcalculatorFilters) > 0 {
		if err := assignInputField(input, "Filters", _bcmpricingcalculatorFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_bcmpricingcalculatorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bcmpricingcalculatorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bcmpricingcalculatorNextToken) > 0 {
		input.NextToken = aws.String(_bcmpricingcalculatorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListBillEstimateLineItems(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bcmpricingcalculator.ListBillEstimateLineItemsOutput
	p := bcmpricingcalculator.NewListBillEstimateLineItemsPaginator(client, input)
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

// Lists all bill estimates for the account.
func bcmpricingcalculator_ListBillEstimates(cfg aws.Config, client *bcmpricingcalculator.Client) {
	input := &bcmpricingcalculator.ListBillEstimatesInput{}

	if len(_bcmpricingcalculatorCreatedAtFilter) > 0 {
		if err := assignInputField(input, "CreatedAtFilter", _bcmpricingcalculatorCreatedAtFilter); err != nil {
			log.Errorf("invalid --created-at-filter: %s", err.Error())
			return
		}
	}
	if len(_bcmpricingcalculatorExpiresAtFilter) > 0 {
		if err := assignInputField(input, "ExpiresAtFilter", _bcmpricingcalculatorExpiresAtFilter); err != nil {
			log.Errorf("invalid --expires-at-filter: %s", err.Error())
			return
		}
	}
	if len(_bcmpricingcalculatorFilters) > 0 {
		if err := assignInputField(input, "Filters", _bcmpricingcalculatorFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_bcmpricingcalculatorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bcmpricingcalculatorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bcmpricingcalculatorNextToken) > 0 {
		input.NextToken = aws.String(_bcmpricingcalculatorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListBillEstimates(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bcmpricingcalculator.ListBillEstimatesOutput
	p := bcmpricingcalculator.NewListBillEstimatesPaginator(client, input)
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

// Lists the commitment modifications associated with a bill scenario.
func bcmpricingcalculator_ListBillScenarioCommitmentModifications(cfg aws.Config, client *bcmpricingcalculator.Client) {
	input := &bcmpricingcalculator.ListBillScenarioCommitmentModificationsInput{
		// BillScenarioId: *string, // Required
	}

	if len(_bcmpricingcalculatorBillScenarioId) > 0 {
		input.BillScenarioId = aws.String(_bcmpricingcalculatorBillScenarioId)
	}
	if len(_bcmpricingcalculatorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bcmpricingcalculatorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bcmpricingcalculatorNextToken) > 0 {
		input.NextToken = aws.String(_bcmpricingcalculatorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListBillScenarioCommitmentModifications(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bcmpricingcalculator.ListBillScenarioCommitmentModificationsOutput
	p := bcmpricingcalculator.NewListBillScenarioCommitmentModificationsPaginator(client, input)
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

// Lists the usage modifications associated with a bill scenario.
func bcmpricingcalculator_ListBillScenarioUsageModifications(cfg aws.Config, client *bcmpricingcalculator.Client) {
	input := &bcmpricingcalculator.ListBillScenarioUsageModificationsInput{
		// BillScenarioId: *string, // Required
	}

	if len(_bcmpricingcalculatorBillScenarioId) > 0 {
		input.BillScenarioId = aws.String(_bcmpricingcalculatorBillScenarioId)
	}
	if len(_bcmpricingcalculatorFilters) > 0 {
		if err := assignInputField(input, "Filters", _bcmpricingcalculatorFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_bcmpricingcalculatorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bcmpricingcalculatorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bcmpricingcalculatorNextToken) > 0 {
		input.NextToken = aws.String(_bcmpricingcalculatorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListBillScenarioUsageModifications(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bcmpricingcalculator.ListBillScenarioUsageModificationsOutput
	p := bcmpricingcalculator.NewListBillScenarioUsageModificationsPaginator(client, input)
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

// Lists all bill scenarios for the account.
func bcmpricingcalculator_ListBillScenarios(cfg aws.Config, client *bcmpricingcalculator.Client) {
	input := &bcmpricingcalculator.ListBillScenariosInput{}

	if len(_bcmpricingcalculatorCreatedAtFilter) > 0 {
		if err := assignInputField(input, "CreatedAtFilter", _bcmpricingcalculatorCreatedAtFilter); err != nil {
			log.Errorf("invalid --created-at-filter: %s", err.Error())
			return
		}
	}
	if len(_bcmpricingcalculatorExpiresAtFilter) > 0 {
		if err := assignInputField(input, "ExpiresAtFilter", _bcmpricingcalculatorExpiresAtFilter); err != nil {
			log.Errorf("invalid --expires-at-filter: %s", err.Error())
			return
		}
	}
	if len(_bcmpricingcalculatorFilters) > 0 {
		if err := assignInputField(input, "Filters", _bcmpricingcalculatorFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_bcmpricingcalculatorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bcmpricingcalculatorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bcmpricingcalculatorNextToken) > 0 {
		input.NextToken = aws.String(_bcmpricingcalculatorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListBillScenarios(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bcmpricingcalculator.ListBillScenariosOutput
	p := bcmpricingcalculator.NewListBillScenariosPaginator(client, input)
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

// Lists all tags associated with a specified resource.
func bcmpricingcalculator_ListTagsForResource(cfg aws.Config, client *bcmpricingcalculator.Client) {
	input := &bcmpricingcalculator.ListTagsForResourceInput{
		// Arn: *string, // Required
	}

	if len(_bcmpricingcalculatorArn) > 0 {
		input.Arn = aws.String(_bcmpricingcalculatorArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the usage associated with a workload estimate.
func bcmpricingcalculator_ListWorkloadEstimateUsage(cfg aws.Config, client *bcmpricingcalculator.Client) {
	input := &bcmpricingcalculator.ListWorkloadEstimateUsageInput{
		// WorkloadEstimateId: *string, // Required
	}

	if len(_bcmpricingcalculatorWorkloadEstimateId) > 0 {
		input.WorkloadEstimateId = aws.String(_bcmpricingcalculatorWorkloadEstimateId)
	}
	if len(_bcmpricingcalculatorFilters) > 0 {
		if err := assignInputField(input, "Filters", _bcmpricingcalculatorFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_bcmpricingcalculatorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bcmpricingcalculatorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bcmpricingcalculatorNextToken) > 0 {
		input.NextToken = aws.String(_bcmpricingcalculatorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListWorkloadEstimateUsage(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bcmpricingcalculator.ListWorkloadEstimateUsageOutput
	p := bcmpricingcalculator.NewListWorkloadEstimateUsagePaginator(client, input)
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

// Lists all workload estimates for the account.
func bcmpricingcalculator_ListWorkloadEstimates(cfg aws.Config, client *bcmpricingcalculator.Client) {
	input := &bcmpricingcalculator.ListWorkloadEstimatesInput{}

	if len(_bcmpricingcalculatorCreatedAtFilter) > 0 {
		if err := assignInputField(input, "CreatedAtFilter", _bcmpricingcalculatorCreatedAtFilter); err != nil {
			log.Errorf("invalid --created-at-filter: %s", err.Error())
			return
		}
	}
	if len(_bcmpricingcalculatorExpiresAtFilter) > 0 {
		if err := assignInputField(input, "ExpiresAtFilter", _bcmpricingcalculatorExpiresAtFilter); err != nil {
			log.Errorf("invalid --expires-at-filter: %s", err.Error())
			return
		}
	}
	if len(_bcmpricingcalculatorFilters) > 0 {
		if err := assignInputField(input, "Filters", _bcmpricingcalculatorFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_bcmpricingcalculatorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bcmpricingcalculatorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bcmpricingcalculatorNextToken) > 0 {
		input.NextToken = aws.String(_bcmpricingcalculatorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListWorkloadEstimates(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bcmpricingcalculator.ListWorkloadEstimatesOutput
	p := bcmpricingcalculator.NewListWorkloadEstimatesPaginator(client, input)
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

// Adds one or more tags to a specified resource.
func bcmpricingcalculator_TagResource(cfg aws.Config, client *bcmpricingcalculator.Client) {
	input := &bcmpricingcalculator.TagResourceInput{
		// Arn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_bcmpricingcalculatorArn) > 0 {
		input.Arn = aws.String(_bcmpricingcalculatorArn)
	}
	if len(_bcmpricingcalculatorTags) > 0 {
		if err := assignInputField(input, "Tags", _bcmpricingcalculatorTags); err != nil {
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

// Removes one or more tags from a specified resource.
func bcmpricingcalculator_UntagResource(cfg aws.Config, client *bcmpricingcalculator.Client) {
	input := &bcmpricingcalculator.UntagResourceInput{
		// Arn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_bcmpricingcalculatorArn) > 0 {
		input.Arn = aws.String(_bcmpricingcalculatorArn)
	}
	if len(_bcmpricingcalculatorTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _bcmpricingcalculatorTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing bill estimate.
func bcmpricingcalculator_UpdateBillEstimate(cfg aws.Config, client *bcmpricingcalculator.Client) {
	input := &bcmpricingcalculator.UpdateBillEstimateInput{
		// Identifier: *string, // Required
	}

	if len(_bcmpricingcalculatorIdentifier) > 0 {
		input.Identifier = aws.String(_bcmpricingcalculatorIdentifier)
	}
	if len(_bcmpricingcalculatorExpiresAt) > 0 {
		if err := assignInputField(input, "ExpiresAt", _bcmpricingcalculatorExpiresAt); err != nil {
			log.Errorf("invalid --expires-at: %s", err.Error())
			return
		}
	}
	if len(_bcmpricingcalculatorName) > 0 {
		input.Name = aws.String(_bcmpricingcalculatorName)
	}

	if resp, err := client.UpdateBillEstimate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing bill scenario.
func bcmpricingcalculator_UpdateBillScenario(cfg aws.Config, client *bcmpricingcalculator.Client) {
	input := &bcmpricingcalculator.UpdateBillScenarioInput{
		// Identifier: *string, // Required
	}

	if len(_bcmpricingcalculatorIdentifier) > 0 {
		input.Identifier = aws.String(_bcmpricingcalculatorIdentifier)
	}
	if len(_bcmpricingcalculatorCostCategoryGroupSharingPreferenceArn) > 0 {
		input.CostCategoryGroupSharingPreferenceArn = aws.String(_bcmpricingcalculatorCostCategoryGroupSharingPreferenceArn)
	}
	if len(_bcmpricingcalculatorExpiresAt) > 0 {
		if err := assignInputField(input, "ExpiresAt", _bcmpricingcalculatorExpiresAt); err != nil {
			log.Errorf("invalid --expires-at: %s", err.Error())
			return
		}
	}
	if len(_bcmpricingcalculatorGroupSharingPreference) > 0 {
		if err := assignInputField(input, "GroupSharingPreference", _bcmpricingcalculatorGroupSharingPreference); err != nil {
			log.Errorf("invalid --group-sharing-preference: %s", err.Error())
			return
		}
	}
	if len(_bcmpricingcalculatorName) > 0 {
		input.Name = aws.String(_bcmpricingcalculatorName)
	}

	if resp, err := client.UpdateBillScenario(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the preferences for Pricing Calculator.
func bcmpricingcalculator_UpdatePreferences(cfg aws.Config, client *bcmpricingcalculator.Client) {
	input := &bcmpricingcalculator.UpdatePreferencesInput{}

	if len(_bcmpricingcalculatorManagementAccountRateTypeSelections) > 0 {
		if err := assignInputField(input, "ManagementAccountRateTypeSelections", _bcmpricingcalculatorManagementAccountRateTypeSelections); err != nil {
			log.Errorf("invalid --management-account-rate-type-selections: %s", err.Error())
			return
		}
	}
	if len(_bcmpricingcalculatorMemberAccountRateTypeSelections) > 0 {
		if err := assignInputField(input, "MemberAccountRateTypeSelections", _bcmpricingcalculatorMemberAccountRateTypeSelections); err != nil {
			log.Errorf("invalid --member-account-rate-type-selections: %s", err.Error())
			return
		}
	}
	if len(_bcmpricingcalculatorStandaloneAccountRateTypeSelections) > 0 {
		if err := assignInputField(input, "StandaloneAccountRateTypeSelections", _bcmpricingcalculatorStandaloneAccountRateTypeSelections); err != nil {
			log.Errorf("invalid --standalone-account-rate-type-selections: %s", err.Error())
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

// Updates an existing workload estimate.
func bcmpricingcalculator_UpdateWorkloadEstimate(cfg aws.Config, client *bcmpricingcalculator.Client) {
	input := &bcmpricingcalculator.UpdateWorkloadEstimateInput{
		// Identifier: *string, // Required
	}

	if len(_bcmpricingcalculatorIdentifier) > 0 {
		input.Identifier = aws.String(_bcmpricingcalculatorIdentifier)
	}
	if len(_bcmpricingcalculatorExpiresAt) > 0 {
		if err := assignInputField(input, "ExpiresAt", _bcmpricingcalculatorExpiresAt); err != nil {
			log.Errorf("invalid --expires-at: %s", err.Error())
			return
		}
	}
	if len(_bcmpricingcalculatorName) > 0 {
		input.Name = aws.String(_bcmpricingcalculatorName)
	}

	if resp, err := client.UpdateWorkloadEstimate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_bcmpricingcalculatorCmd)
	_bcmpricingcalculatorCmd.Flags().SortFlags = false

	_bcmpricingcalculatorCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_bcmpricingcalculatorCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_bcmpricingcalculatorCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_bcmpricingcalculatorCmd.Flags().StringVarP(&_bcmpricingcalculatorArn, "arn", "", "", "ARN")
	_bcmpricingcalculatorCmd.Flags().StringVarP(&_bcmpricingcalculatorBillEstimateId, "bill-estimate-id", "", "", "Bill Estimate ID")
	_bcmpricingcalculatorCmd.Flags().StringVarP(&_bcmpricingcalculatorBillScenarioId, "bill-scenario-id", "", "", "Bill Scenario ID")
	_bcmpricingcalculatorCmd.Flags().StringVarP(&_bcmpricingcalculatorClientToken, "client-token", "", "", "Client Token")
	_bcmpricingcalculatorCmd.Flags().StringVarP(&_bcmpricingcalculatorCommitmentModifications, "commitment-modifications", "", "", "Commitment Modifications")
	_bcmpricingcalculatorCmd.Flags().StringVarP(&_bcmpricingcalculatorCostCategoryGroupSharingPreferenceArn, "cost-category-group-sharing-preference-arn", "", "", "Cost Category Group Sharing Preference ARN")
	_bcmpricingcalculatorCmd.Flags().StringVarP(&_bcmpricingcalculatorCreatedAtFilter, "created-at-filter", "", "", "Created At Filter")
	_bcmpricingcalculatorCmd.Flags().StringVarP(&_bcmpricingcalculatorExpiresAt, "expires-at", "", "", "Expires At")
	_bcmpricingcalculatorCmd.Flags().StringVarP(&_bcmpricingcalculatorExpiresAtFilter, "expires-at-filter", "", "", "Expires At Filter")
	_bcmpricingcalculatorCmd.Flags().StringVarP(&_bcmpricingcalculatorFilters, "filters", "", "", "Filters")
	_bcmpricingcalculatorCmd.Flags().StringVarP(&_bcmpricingcalculatorGroupSharingPreference, "group-sharing-preference", "", "", "Group Sharing Preference")
	_bcmpricingcalculatorCmd.Flags().StringVarP(&_bcmpricingcalculatorIdentifier, "identifier", "", "", "Identifier")
	_bcmpricingcalculatorCmd.Flags().StringSliceVarP(&_bcmpricingcalculatorIds, "ids", "", nil, "Ids")
	_bcmpricingcalculatorCmd.Flags().StringVarP(&_bcmpricingcalculatorManagementAccountRateTypeSelections, "management-account-rate-type-selections", "", "", "Management Account Rate Type Selections")
	_bcmpricingcalculatorCmd.Flags().StringVarP(&_bcmpricingcalculatorMaxResults, "max-results", "", "", "Max Results")
	_bcmpricingcalculatorCmd.Flags().StringVarP(&_bcmpricingcalculatorMemberAccountRateTypeSelections, "member-account-rate-type-selections", "", "", "Member Account Rate Type Selections")
	_bcmpricingcalculatorCmd.Flags().StringVarP(&_bcmpricingcalculatorName, "name", "", "", "Name")
	_bcmpricingcalculatorCmd.Flags().StringVarP(&_bcmpricingcalculatorNextToken, "next-token", "", "", "Next Token")
	_bcmpricingcalculatorCmd.Flags().StringVarP(&_bcmpricingcalculatorRateType, "rate-type", "", "", "Rate Type")
	_bcmpricingcalculatorCmd.Flags().StringVarP(&_bcmpricingcalculatorStandaloneAccountRateTypeSelections, "standalone-account-rate-type-selections", "", "", "Standalone Account Rate Type Selections")
	_bcmpricingcalculatorCmd.Flags().StringSliceVarP(&_bcmpricingcalculatorTagKeys, "tag-keys", "", nil, "Tag Keys")
	_bcmpricingcalculatorCmd.Flags().StringVarP(&_bcmpricingcalculatorTags, "tags", "", "", "Tags")
	_bcmpricingcalculatorCmd.Flags().StringVarP(&_bcmpricingcalculatorUsage, "usage", "", "", "Usage")
	_bcmpricingcalculatorCmd.Flags().StringVarP(&_bcmpricingcalculatorUsageModifications, "usage-modifications", "", "", "Usage Modifications")
	_bcmpricingcalculatorCmd.Flags().StringVarP(&_bcmpricingcalculatorWorkloadEstimateId, "workload-estimate-id", "", "", "Workload Estimate ID")

	_bcmpricingcalculatorCmd.Flags().BoolVarP(&_bcmpricingcalculatorBatchCreateBillScenarioCommitmentModification, "batch-create-bill-scenario-commitment-modification", "", false, "Batch Create Bill Scenario Commitment Modification")
	_bcmpricingcalculatorCmd.Flags().BoolVarP(&_bcmpricingcalculatorBatchCreateBillScenarioUsageModification, "batch-create-bill-scenario-usage-modification", "", false, "Batch Create Bill Scenario Usage Modification")
	_bcmpricingcalculatorCmd.Flags().BoolVarP(&_bcmpricingcalculatorBatchCreateWorkloadEstimateUsage, "batch-create-workload-estimate-usage", "", false, "Batch Create Workload Estimate Usage")
	_bcmpricingcalculatorCmd.Flags().BoolVarP(&_bcmpricingcalculatorBatchDeleteBillScenarioCommitmentModification, "batch-delete-bill-scenario-commitment-modification", "", false, "Batch Delete Bill Scenario Commitment Modification")
	_bcmpricingcalculatorCmd.Flags().BoolVarP(&_bcmpricingcalculatorBatchDeleteBillScenarioUsageModification, "batch-delete-bill-scenario-usage-modification", "", false, "Batch Delete Bill Scenario Usage Modification")
	_bcmpricingcalculatorCmd.Flags().BoolVarP(&_bcmpricingcalculatorBatchDeleteWorkloadEstimateUsage, "batch-delete-workload-estimate-usage", "", false, "Batch Delete Workload Estimate Usage")
	_bcmpricingcalculatorCmd.Flags().BoolVarP(&_bcmpricingcalculatorBatchUpdateBillScenarioCommitmentModification, "batch-update-bill-scenario-commitment-modification", "", false, "Batch Update Bill Scenario Commitment Modification")
	_bcmpricingcalculatorCmd.Flags().BoolVarP(&_bcmpricingcalculatorBatchUpdateBillScenarioUsageModification, "batch-update-bill-scenario-usage-modification", "", false, "Batch Update Bill Scenario Usage Modification")
	_bcmpricingcalculatorCmd.Flags().BoolVarP(&_bcmpricingcalculatorBatchUpdateWorkloadEstimateUsage, "batch-update-workload-estimate-usage", "", false, "Batch Update Workload Estimate Usage")
	_bcmpricingcalculatorCmd.Flags().BoolVarP(&_bcmpricingcalculatorCreateBillEstimate, "create-bill-estimate", "", false, "Create Bill Estimate")
	_bcmpricingcalculatorCmd.Flags().BoolVarP(&_bcmpricingcalculatorCreateBillScenario, "create-bill-scenario", "", false, "Create Bill Scenario")
	_bcmpricingcalculatorCmd.Flags().BoolVarP(&_bcmpricingcalculatorCreateWorkloadEstimate, "create-workload-estimate", "", false, "Create Workload Estimate")
	_bcmpricingcalculatorCmd.Flags().BoolVarP(&_bcmpricingcalculatorDeleteBillEstimate, "delete-bill-estimate", "", false, "Delete Bill Estimate")
	_bcmpricingcalculatorCmd.Flags().BoolVarP(&_bcmpricingcalculatorDeleteBillScenario, "delete-bill-scenario", "", false, "Delete Bill Scenario")
	_bcmpricingcalculatorCmd.Flags().BoolVarP(&_bcmpricingcalculatorDeleteWorkloadEstimate, "delete-workload-estimate", "", false, "Delete Workload Estimate")
	_bcmpricingcalculatorCmd.Flags().BoolVarP(&_bcmpricingcalculatorGetBillEstimate, "get-bill-estimate", "", false, "Get Bill Estimate")
	_bcmpricingcalculatorCmd.Flags().BoolVarP(&_bcmpricingcalculatorGetBillScenario, "get-bill-scenario", "", false, "Get Bill Scenario")
	_bcmpricingcalculatorCmd.Flags().BoolVarP(&_bcmpricingcalculatorGetPreferences, "get-preferences", "", false, "Get Preferences")
	_bcmpricingcalculatorCmd.Flags().BoolVarP(&_bcmpricingcalculatorGetWorkloadEstimate, "get-workload-estimate", "", false, "Get Workload Estimate")
	_bcmpricingcalculatorCmd.Flags().BoolVarP(&_bcmpricingcalculatorListBillEstimateCommitments, "list-bill-estimate-commitments", "", false, "List Bill Estimate Commitments")
	_bcmpricingcalculatorCmd.Flags().BoolVarP(&_bcmpricingcalculatorListBillEstimateInputCommitmentModifications, "list-bill-estimate-input-commitment-modifications", "", false, "List Bill Estimate Input Commitment Modifications")
	_bcmpricingcalculatorCmd.Flags().BoolVarP(&_bcmpricingcalculatorListBillEstimateInputUsageModifications, "list-bill-estimate-input-usage-modifications", "", false, "List Bill Estimate Input Usage Modifications")
	_bcmpricingcalculatorCmd.Flags().BoolVarP(&_bcmpricingcalculatorListBillEstimateLineItems, "list-bill-estimate-line-items", "", false, "List Bill Estimate Line Items")
	_bcmpricingcalculatorCmd.Flags().BoolVarP(&_bcmpricingcalculatorListBillEstimates, "list-bill-estimates", "", false, "List Bill Estimates")
	_bcmpricingcalculatorCmd.Flags().BoolVarP(&_bcmpricingcalculatorListBillScenarioCommitmentModifications, "list-bill-scenario-commitment-modifications", "", false, "List Bill Scenario Commitment Modifications")
	_bcmpricingcalculatorCmd.Flags().BoolVarP(&_bcmpricingcalculatorListBillScenarioUsageModifications, "list-bill-scenario-usage-modifications", "", false, "List Bill Scenario Usage Modifications")
	_bcmpricingcalculatorCmd.Flags().BoolVarP(&_bcmpricingcalculatorListBillScenarios, "list-bill-scenarios", "", false, "List Bill Scenarios")
	_bcmpricingcalculatorCmd.Flags().BoolVarP(&_bcmpricingcalculatorListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_bcmpricingcalculatorCmd.Flags().BoolVarP(&_bcmpricingcalculatorListWorkloadEstimateUsage, "list-workload-estimate-usage", "", false, "List Workload Estimate Usage")
	_bcmpricingcalculatorCmd.Flags().BoolVarP(&_bcmpricingcalculatorListWorkloadEstimates, "list-workload-estimates", "", false, "List Workload Estimates")
	_bcmpricingcalculatorCmd.Flags().BoolVarP(&_bcmpricingcalculatorTagResource, "tag-resource", "", false, "Tag Resource")
	_bcmpricingcalculatorCmd.Flags().BoolVarP(&_bcmpricingcalculatorUntagResource, "untag-resource", "", false, "Untag Resource")
	_bcmpricingcalculatorCmd.Flags().BoolVarP(&_bcmpricingcalculatorUpdateBillEstimate, "update-bill-estimate", "", false, "Update Bill Estimate")
	_bcmpricingcalculatorCmd.Flags().BoolVarP(&_bcmpricingcalculatorUpdateBillScenario, "update-bill-scenario", "", false, "Update Bill Scenario")
	_bcmpricingcalculatorCmd.Flags().BoolVarP(&_bcmpricingcalculatorUpdatePreferences, "update-preferences", "", false, "Update Preferences")
	_bcmpricingcalculatorCmd.Flags().BoolVarP(&_bcmpricingcalculatorUpdateWorkloadEstimate, "update-workload-estimate", "", false, "Update Workload Estimate")

}
