package rekognition

// DeleteDataset is generated as a reference stub.
// Executable command wiring lives under cmd/rekognition.go.
//
// This operation applies only to Amazon Rekognition Custom Labels.
//
// Deletes an existing Amazon Rekognition Custom Labels dataset. Deleting a
// dataset might take while. Use DescribeDatasetto check the current status. The dataset is still
// deleting if the value of Status is DELETE_IN_PROGRESS . If you try to access the
// dataset after it is deleted, you get a ResourceNotFoundException exception.
//
// You can't delete a dataset while it is creating ( Status = CREATE_IN_PROGRESS )
// or if the dataset is updating ( Status = UPDATE_IN_PROGRESS ).
//
// This operation requires permissions to perform the rekognition:DeleteDataset
// action.
