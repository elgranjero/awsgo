package codepipeline

// GetThirdPartyJobDetails is generated as a reference stub.
// Executable command wiring lives under cmd/codepipeline.go.
//
// Requests the details of a job for a third party action. Used for partner
// actions only.
//
// When this API is called, CodePipeline returns temporary credentials for the S3
// bucket used to store artifacts for the pipeline, if the action requires access
// to that S3 bucket for input or output artifacts. This API also returns any
// secret values defined for the action.
