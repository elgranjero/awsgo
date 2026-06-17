package ssm

// DescribePatchProperties is generated as a reference stub.
// Executable command wiring lives under cmd/ssm.go.
//
// Lists the properties of available patches organized by product, product family,
// classification, severity, and other properties of available patches. You can use
// the reported properties in the filters you specify in requests for operations
// such as CreatePatchBaseline, UpdatePatchBaseline, DescribeAvailablePatches, and DescribePatchBaselines.
//
// The following section lists the properties that can be used in filters for each
// major operating system type:
//
// AMAZON_LINUX Valid properties: PRODUCT | CLASSIFICATION | SEVERITY
//
// AMAZON_LINUX_2 Valid properties: PRODUCT | CLASSIFICATION | SEVERITY
//
// AMAZON_LINUX_2023 Valid properties: PRODUCT | CLASSIFICATION | SEVERITY
//
// CENTOS Valid properties: PRODUCT | CLASSIFICATION | SEVERITY
//
// DEBIAN Valid properties: PRODUCT | PRIORITY
//
// MACOS Valid properties: PRODUCT | CLASSIFICATION
//
// ORACLE_LINUX Valid properties: PRODUCT | CLASSIFICATION | SEVERITY
//
// REDHAT_ENTERPRISE_LINUX Valid properties: PRODUCT | CLASSIFICATION | SEVERITY
//
// SUSE Valid properties: PRODUCT | CLASSIFICATION | SEVERITY
//
// UBUNTU Valid properties: PRODUCT | PRIORITY
//
// WINDOWS Valid properties: PRODUCT | PRODUCT_FAMILY | CLASSIFICATION |
// MSRC_SEVERITY
