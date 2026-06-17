package redshiftdata

// ListSchemas is generated as a reference stub.
// Executable command wiring lives under cmd/redshiftdata.go.
//
// Lists the schemas in a database. A token is returned to page through the schema
// list. Depending on the authorization method, use one of the following
// combinations of request parameters:
//
// - Secrets Manager - when connecting to a cluster, provide the secret-arn of a
// secret stored in Secrets Manager which has username and password . The
// specified secret contains credentials to connect to the database you specify.
// When you are connecting to a cluster, you also supply the database name, If you
// provide a cluster identifier ( dbClusterIdentifier ), it must match the
// cluster identifier stored in the secret. When you are connecting to a serverless
// workgroup, you also supply the database name.
//
// - Temporary credentials - when connecting to your data warehouse, choose one
// of the following options:
//
// - When connecting to a serverless workgroup, specify the workgroup name and
// database name. The database user name is derived from the IAM identity. For
// example, arn:iam::123456789012:user:foo has the database user name IAM:foo .
// Also, permission to call the redshift-serverless:GetCredentials operation is
// required.
//
// - When connecting to a cluster as an IAM identity, specify the cluster
// identifier and the database name. The database user name is derived from the IAM
// identity. For example, arn:iam::123456789012:user:foo has the database user
// name IAM:foo . Also, permission to call the
// redshift:GetClusterCredentialsWithIAM operation is required.
//
// - When connecting to a cluster as a database user, specify the cluster
// identifier, the database name, and the database user name. Also, permission to
// call the redshift:GetClusterCredentials operation is required.
//
// For more information about the Amazon Redshift Data API and CLI usage examples,
// see [Using the Amazon Redshift Data API]in the Amazon Redshift Management Guide.
//
// [Using the Amazon Redshift Data API]: https://docs.aws.amazon.com/redshift/latest/mgmt/data-api.html
