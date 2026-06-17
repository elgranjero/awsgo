package sagemakerfeaturestoreruntime

// DeleteRecord is generated as a reference stub.
// Executable command wiring lives under cmd/sagemakerfeaturestoreruntime.go.
//
// Deletes a Record from a FeatureGroup in the OnlineStore . Feature Store supports
// both SoftDelete and HardDelete . For SoftDelete (default), feature columns are
// set to null and the record is no longer retrievable by GetRecord or
// BatchGetRecord . For HardDelete , the complete Record is removed from the
// OnlineStore . In both cases, Feature Store appends the deleted record marker to
// the OfflineStore . The deleted record marker is a record with the same
// RecordIdentifer as the original, but with is_deleted value set to True ,
// EventTime set to the delete input EventTime , and other feature values set to
// null .
//
// Note that the EventTime specified in DeleteRecord should be set later than the
// EventTime of the existing record in the OnlineStore for that RecordIdentifer .
// If it is not, the deletion does not occur:
//
// - For SoftDelete , the existing (not deleted) record remains in the
// OnlineStore , though the delete record marker is still written to the
// OfflineStore .
//
// - HardDelete returns EventTime : 400 ValidationException to indicate that the
// delete operation failed. No delete record marker is written to the
// OfflineStore .
//
// When a record is deleted from the OnlineStore , the deleted record marker is
// appended to the OfflineStore . If you have the Iceberg table format enabled for
// your OfflineStore , you can remove all history of a record from the OfflineStore
// using Amazon Athena or Apache Spark. For information on how to hard delete a
// record from the OfflineStore with the Iceberg table format enabled, see [Delete records from the offline store].
//
// [Delete records from the offline store]: https://docs.aws.amazon.com/sagemaker/latest/dg/feature-store-delete-records-offline-store.html#feature-store-delete-records-offline-store
