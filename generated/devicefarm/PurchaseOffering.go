package devicefarm

// PurchaseOffering is generated as a reference stub.
// Executable command wiring lives under cmd/devicefarm.go.
//
// Immediately purchases offerings for an AWS account. Offerings renew with the
// latest total purchased quantity for an offering, unless the renewal was
// overridden. The API returns a NotEligible error if the user is not permitted to
// invoke the operation. If you must be able to invoke this operation, contact aws-devicefarm-support(at)amazon.com.
