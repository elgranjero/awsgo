package athena

// CreateDataCatalog is generated as a reference stub.
// Executable command wiring lives under cmd/athena.go.
//
// Creates (registers) a data catalog with the specified name and properties.
// Catalogs created are visible to all users of the same Amazon Web Services
// account.
//
// For a FEDERATED catalog, this API operation creates the following resources.
//
// - CFN Stack Name with a maximum length of 128 characters and prefix
// athenafederatedcatalog-CATALOG_NAME_SANITIZED with length 23 characters.
//
// - Lambda Function Name with a maximum length of 64 characters and prefix
// athenafederatedcatalog_CATALOG_NAME_SANITIZED with length 23 characters.
//
// - Glue Connection Name with a maximum length of 255 characters and a prefix
// athenafederatedcatalog_CATALOG_NAME_SANITIZED with length 23 characters.
