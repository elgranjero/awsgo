package sagemaker

// DeleteFeatureGroup is generated as a reference stub.
// Executable command wiring lives under cmd/sagemaker.go.
//
// Delete the FeatureGroup and any data that was written to the OnlineStore of the
// FeatureGroup . Data cannot be accessed from the OnlineStore immediately after
// DeleteFeatureGroup is called.
//
// Data written into the OfflineStore will not be deleted. The Amazon Web Services
// Glue database and tables that are automatically created for your OfflineStore
// are not deleted.
//
// Note that it can take approximately 10-15 minutes to delete an OnlineStore
// FeatureGroup with the InMemory StorageType .
