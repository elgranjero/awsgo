package sagemaker

// CreateFeatureGroup is generated as a reference stub.
// Executable command wiring lives under cmd/sagemaker.go.
//
// Create a new FeatureGroup . A FeatureGroup is a group of Features defined in
// the FeatureStore to describe a Record .
//
// The FeatureGroup defines the schema and features contained in the FeatureGroup .
// A FeatureGroup definition is composed of a list of Features , a
// RecordIdentifierFeatureName , an EventTimeFeatureName and configurations for
// its OnlineStore and OfflineStore . Check [Amazon Web Services service quotas] to see the FeatureGroup s quota for
// your Amazon Web Services account.
//
// Note that it can take approximately 10-15 minutes to provision an OnlineStore
// FeatureGroup with the InMemory StorageType .
//
// You must include at least one of OnlineStoreConfig and OfflineStoreConfig to
// create a FeatureGroup .
//
// [Amazon Web Services service quotas]: https://docs.aws.amazon.com/general/latest/gr/aws_service_limits.html
