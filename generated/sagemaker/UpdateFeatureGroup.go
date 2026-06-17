package sagemaker

// UpdateFeatureGroup is generated as a reference stub.
// Executable command wiring lives under cmd/sagemaker.go.
//
// Updates the feature group by either adding features or updating the online
// store configuration. Use one of the following request parameters at a time while
// using the UpdateFeatureGroup API.
//
// You can add features for your feature group using the FeatureAdditions request
// parameter. Features cannot be removed from a feature group.
//
// You can update the online store configuration by using the OnlineStoreConfig
// request parameter. If a TtlDuration is specified, the default TtlDuration
// applies for all records added to the feature group after the feature group is
// updated. If a record level TtlDuration exists from using the PutRecord API, the
// record level TtlDuration applies to that record instead of the default
// TtlDuration . To remove the default TtlDuration from an existing feature group,
// use the UpdateFeatureGroup API and set the TtlDuration Unit and Value to null .
