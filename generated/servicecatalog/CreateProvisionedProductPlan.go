package servicecatalog

// CreateProvisionedProductPlan is generated as a reference stub.
// Executable command wiring lives under cmd/servicecatalog.go.
//
// Creates a plan.
//
// A plan includes the list of resources to be created (when provisioning a new
// product) or modified (when updating a provisioned product) when the plan is
// executed.
//
// You can create one plan for each provisioned product. To create a plan for an
// existing provisioned product, the product status must be AVAILABLE or TAINTED.
//
// To view the resource changes in the change set, use DescribeProvisionedProductPlan. To create or modify the
// provisioned product, use ExecuteProvisionedProductPlan.
