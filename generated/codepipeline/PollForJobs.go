package codepipeline

// PollForJobs is generated as a reference stub.
// Executable command wiring lives under cmd/codepipeline.go.
//
// Returns information about any jobs for CodePipeline to act on. PollForJobs is
// valid only for action types with "Custom" in the owner field. If the action type
// contains AWS or ThirdParty in the owner field, the PollForJobs action returns
// an error.
//
// When this API is called, CodePipeline returns temporary credentials for the S3
// bucket used to store artifacts for the pipeline, if the action requires access
// to that S3 bucket for input or output artifacts. This API also returns any
// secret values defined for the action.
