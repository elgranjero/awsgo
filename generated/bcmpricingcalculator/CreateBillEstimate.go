package bcmpricingcalculator

// CreateBillEstimate is generated as a reference stub.
// Executable command wiring lives under cmd/bcmpricingcalculator.go.
//
// Create a Bill estimate from a Bill scenario. In the Bill scenario you can
//
// model usage addition, usage changes, and usage removal. You can also model
// commitment addition and commitment removal. After all changes in a Bill scenario
// is made satisfactorily, you can call this API with a Bill scenario ID to
// generate the Bill estimate. Bill estimate calculates the pre-tax cost for your
// consolidated billing family, incorporating all modeled usage and commitments
// alongside existing usage and commitments from your most recent completed
// anniversary bill, with any applicable discounts applied.
