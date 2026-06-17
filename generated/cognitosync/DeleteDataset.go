package cognitosync

// DeleteDataset is generated as a reference stub.
// Executable command wiring lives under cmd/cognitosync.go.
//
// Deletes the specific dataset. The dataset will be deleted permanently, and the
// action can't be undone. Datasets that this dataset was merged with will no
// longer report the merge. Any subsequent operation on this dataset will result in
// a ResourceNotFoundException.
//
// This API can be called with temporary user credentials provided by Cognito
// Identity or with developer credentials.
