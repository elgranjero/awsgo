package rekognition

// CopyProjectVersion is generated as a reference stub.
// Executable command wiring lives under cmd/rekognition.go.
//
// This operation applies only to Amazon Rekognition Custom Labels.
//
// Copies a version of an Amazon Rekognition Custom Labels model from a source
// project to a destination project. The source and destination projects can be in
// different AWS accounts but must be in the same AWS Region. You can't copy a
// model to another AWS service.
//
// To copy a model version to a different AWS account, you need to create a
// resource-based policy known as a project policy. You attach the project policy
// to the source project by calling PutProjectPolicy. The project policy gives permission to copy
// the model version from a trusting AWS account to a trusted account.
//
// For more information creating and attaching a project policy, see Attaching a
// project policy (SDK) in the Amazon Rekognition Custom Labels Developer Guide.
//
// If you are copying a model version to a project in the same AWS account, you
// don't need to create a project policy.
//
// Copying project versions is supported only for Custom Labels models.
//
// To copy a model, the destination project, source project, and source model
// version must already exist.
//
// Copying a model version takes a while to complete. To get the current status,
// call DescribeProjectVersionsand check the value of Status in the ProjectVersionDescription object. The copy operation has
// finished when the value of Status is COPYING_COMPLETED .
//
// This operation requires permissions to perform the
// rekognition:CopyProjectVersion action.
