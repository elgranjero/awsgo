package sagemaker

// DeleteWorkforce is generated as a reference stub.
// Executable command wiring lives under cmd/sagemaker.go.
//
// Use this operation to delete a workforce.
//
// If you want to create a new workforce in an Amazon Web Services Region where a
// workforce already exists, use this operation to delete the existing workforce
// and then use [CreateWorkforce]to create a new workforce.
//
// If a private workforce contains one or more work teams, you must use the [DeleteWorkteam]
// operation to delete all work teams before you delete the workforce. If you try
// to delete a workforce that contains one or more work teams, you will receive a
// ResourceInUse error.
//
// [CreateWorkforce]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateWorkforce.html
// [DeleteWorkteam]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DeleteWorkteam.html
