package servicecatalog

// ImportAsProvisionedProduct is generated as a reference stub.
// Executable command wiring lives under cmd/servicecatalog.go.
//
// Requests the import of a resource as an Service Catalog provisioned product
//
// that is associated to an Service Catalog product and provisioning artifact. Once
// imported, all supported governance actions are supported on the provisioned
// product.
//
// Resource import only supports CloudFormation stack ARNs. CloudFormation
// StackSets, and non-root nested stacks, are not supported.
//
// The CloudFormation stack must have one of the following statuses to be
// imported: CREATE_COMPLETE , UPDATE_COMPLETE , UPDATE_ROLLBACK_COMPLETE ,
// IMPORT_COMPLETE , and IMPORT_ROLLBACK_COMPLETE .
//
// Import of the resource requires that the CloudFormation stack template matches
// the associated Service Catalog product provisioning artifact.
//
// When you import an existing CloudFormation stack into a portfolio, Service
// Catalog does not apply the product's associated constraints during the import
// process. Service Catalog applies the constraints after you call
// UpdateProvisionedProduct for the provisioned product.
//
// The user or role that performs this operation must have the
// cloudformation:GetTemplate and cloudformation:DescribeStacks IAM policy
// permissions.
//
// You can only import one provisioned product at a time. The product's
// CloudFormation stack must have the IMPORT_COMPLETE status before you import
// another.
