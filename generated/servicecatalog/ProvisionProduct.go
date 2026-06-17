package servicecatalog

// ProvisionProduct is generated as a reference stub.
// Executable command wiring lives under cmd/servicecatalog.go.
//
// Provisions the specified product.
//
// A provisioned product is a resourced instance of a product. For example,
// provisioning a product that's based on an CloudFormation template launches an
// CloudFormation stack and its underlying resources. You can check the status of
// this request using DescribeRecord.
//
// If the request contains a tag key with an empty list of values, there's a tag
// conflict for that key. Don't include conflicted keys as tags, or this will cause
// the error "Parameter validation failed: Missing required parameter in
// Tags[N]:Value".
//
// When provisioning a product that's been added to a portfolio, you must grant
// your user, group, or role access to the portfolio. For more information, see [Granting users access]in
// the Service Catalog User Guide.
//
// [Granting users access]: https://docs.aws.amazon.com/servicecatalog/latest/adminguide/catalogs_portfolios_users.html
