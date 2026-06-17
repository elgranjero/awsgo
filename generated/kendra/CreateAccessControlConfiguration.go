package kendra

// CreateAccessControlConfiguration is generated as a reference stub.
// Executable command wiring lives under cmd/kendra.go.
//
// Creates an access configuration for your documents. This includes user and
// group access information for your documents. This is useful for user context
// filtering, where search results are filtered based on the user or their group
// access to documents.
//
// You can use this to re-configure your existing document level access control
// without indexing all of your documents again. For example, your index contains
// top-secret company documents that only certain employees or users should access.
// One of these users leaves the company or switches to a team that should be
// blocked from accessing top-secret documents. The user still has access to
// top-secret documents because the user had access when your documents were
// previously indexed. You can create a specific access control configuration for
// the user with deny access. You can later update the access control configuration
// to allow access if the user returns to the company and re-joins the 'top-secret'
// team. You can re-configure access control for your documents as circumstances
// change.
//
// To apply your access control configuration to certain documents, you call the [BatchPutDocument]
// API with the AccessControlConfigurationId included in the [Document] object. If you use
// an S3 bucket as a data source, you update the .metadata.json with the
// AccessControlConfigurationId and synchronize your data source. Amazon Kendra
// currently only supports access control configuration for S3 data sources and
// documents indexed using the BatchPutDocument API.
//
// You can't configure access control using CreateAccessControlConfiguration for
// an Amazon Kendra Gen AI Enterprise Edition index. Amazon Kendra will return a
// ValidationException error for a Gen_AI_ENTERPRISE_EDITION index.
//
// [BatchPutDocument]: https://docs.aws.amazon.com/kendra/latest/dg/API_BatchPutDocument.html
// [Document]: https://docs.aws.amazon.com/kendra/latest/dg/API_Document.html
