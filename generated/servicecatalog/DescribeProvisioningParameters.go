package servicecatalog

// DescribeProvisioningParameters is generated as a reference stub.
// Executable command wiring lives under cmd/servicecatalog.go.
//
// Gets information about the configuration required to provision the specified
// product using the specified provisioning artifact.
//
// If the output contains a TagOption key with an empty list of values, there is a
// TagOption conflict for that key. The end user cannot take action to fix the
// conflict, and launch is not blocked. In subsequent calls to ProvisionProduct, do not include
// conflicted TagOption keys as tags, or this causes the error "Parameter
// validation failed: Missing required parameter in Tags[N]:Value". Tag the
// provisioned product with the value sc-tagoption-conflict-portfolioId-productId .
