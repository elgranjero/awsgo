package cloudtrail

// EnableFederation is generated as a reference stub.
// Executable command wiring lives under cmd/cloudtrail.go.
//
// Enables Lake query federation on the specified event data store. Federating an
//
// event data store lets you view the metadata associated with the event data store
// in the Glue [Data Catalog]and run SQL queries against your event data using Amazon Athena.
// The table metadata stored in the Glue Data Catalog lets the Athena query engine
// know how to find, read, and process the data that you want to query.
//
// When you enable Lake query federation, CloudTrail creates a managed database
// named aws:cloudtrail (if the database doesn't already exist) and a managed
// federated table in the Glue Data Catalog. The event data store ID is used for
// the table name. CloudTrail registers the role ARN and event data store in [Lake Formation], the
// service responsible for allowing fine-grained access control of the federated
// resources in the Glue Data Catalog.
//
// For more information about Lake query federation, see [Federate an event data store].
//
// [Federate an event data store]: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/query-federation.html
// [Lake Formation]: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/query-federation-lake-formation.html
// [Data Catalog]: https://docs.aws.amazon.com/glue/latest/dg/components-overview.html#data-catalog-intro
