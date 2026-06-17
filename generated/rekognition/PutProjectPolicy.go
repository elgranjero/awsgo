package rekognition

// PutProjectPolicy is generated as a reference stub.
// Executable command wiring lives under cmd/rekognition.go.
//
// This operation applies only to Amazon Rekognition Custom Labels.
//
// Attaches a project policy to a Amazon Rekognition Custom Labels project in a
// trusting AWS account. A project policy specifies that a trusted AWS account can
// copy a model version from a trusting AWS account to a project in the trusted AWS
// account. To copy a model version you use the CopyProjectVersionoperation. Only applies to Custom
// Labels projects.
//
// For more information about the format of a project policy document, see
// Attaching a project policy (SDK) in the Amazon Rekognition Custom Labels
// Developer Guide.
//
// The response from PutProjectPolicy is a revision ID for the project policy. You
// can attach multiple project policies to a project. You can also update an
// existing project policy by specifying the policy revision ID of the existing
// policy.
//
// To remove a project policy from a project, call DeleteProjectPolicy. To get a list of project
// policies attached to a project, call ListProjectPolicies.
//
// You copy a model version by calling CopyProjectVersion.
//
// This operation requires permissions to perform the rekognition:PutProjectPolicy
// action.
