package bcmpricingcalculator

// BatchDeleteBillScenarioCommitmentModification is generated as a reference stub.
// Executable command wiring lives under cmd/bcmpricingcalculator.go.
//
// Delete commitment that you have created in a Bill Scenario. You can only
//
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
