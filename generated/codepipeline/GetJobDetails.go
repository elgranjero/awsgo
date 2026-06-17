package codepipeline

// GetJobDetails is generated as a reference stub.
// Executable command wiring lives under cmd/codepipeline.go.
//
// Returns information about a job. Used for custom actions only.
//
// When this API is called, CodePipeline returns temporary credentials for the S3
// bucket used to store artifacts for the pipeline, if the action requires access
// to that S3 bucket for input or output artifacts. This API also returns any
// secret values defined for the action.
