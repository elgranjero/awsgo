package kendra

// UpdateAccessControlConfiguration is generated as a reference stub.
// Executable command wiring lives under cmd/kendra.go.
//
// Updates an access control configuration for your documents in an index. This
// includes user and group access information for your documents. This is useful
// for user context filtering, where search results are filtered based on the user
// or their group access to documents.
//
// You can update an access control configuration you created without indexing all
// of your documents again. For example, your index contains top-secret company
// documents that only certain employees or users should access. You created an
// 'allow' access control configuration for one user who recently joined the
// 'top-secret' team, switching from a team with 'deny' access to top-secret
// documents. However, the user suddenly returns to their previous team and should
// no longer have access to top secret documents. You can update the access control
// configuration to re-configure access control for your documents as circumstances
// change.
//
// You call the [BatchPutDocument] API to apply the updated access control configuration, with the
// AccessControlConfigurationId included in the [Document] object. If you use an S3 bucket
// as a data source, you synchronize your data source to apply the
// AccessControlConfigurationId in the .metadata.json file. Amazon Kendra
// currently only supports access control configuration for S3 data sources and
// documents indexed using the BatchPutDocument API.
//
// You can't configure access control using CreateAccessControlConfiguration for
// an Amazon Kendra Gen AI Enterprise Edition index. Amazon Kendra will return a
// ValidationException error for a Gen_AI_ENTERPRISE_EDITION index.
//
// [BatchPutDocument]: https://docs.aws.amazon.com/kendra/latest/dg/API_BatchPutDocument.html
// [Document]: https://docs.aws.amazon.com/kendra/latest/dg/API_Document.html
