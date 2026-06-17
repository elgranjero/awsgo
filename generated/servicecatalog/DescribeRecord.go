package servicecatalog

// DescribeRecord is generated as a reference stub.
// Executable command wiring lives under cmd/servicecatalog.go.
//
// Gets information about the specified request operation.
//
// Use this operation after calling a request operation (for example, ProvisionProduct, TerminateProvisionedProduct, or UpdateProvisionedProduct).
//
// If a provisioned product was transferred to a new owner using UpdateProvisionedProductProperties, the new owner
// will be able to describe all past records for that product. The previous owner
// will no longer be able to describe the records, but will be able to use ListRecordHistoryto see
// the product's history from when he was the owner.
