package lookoutequipment

// DeleteDataset is generated as a reference stub.
// Executable command wiring lives under cmd/lookoutequipment.go.
//
// Deletes a dataset and associated artifacts. The operation will check to see if
//
// any inference scheduler or data ingestion job is currently using the dataset,
// and if there isn't, the dataset, its metadata, and any associated data stored in
// S3 will be deleted. This does not affect any models that used this dataset for
// training and evaluation, but does prevent it from being used in the future.
