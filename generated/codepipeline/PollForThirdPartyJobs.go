package codepipeline

// PollForThirdPartyJobs is generated as a reference stub.
// Executable command wiring lives under cmd/codepipeline.go.
//
// Determines whether there are any third party jobs for a job worker to act on.
// Used for partner actions only.
//
// When this API is called, CodePipeline returns temporary credentials for the S3
// bucket used to store artifacts for the pipeline, if the action requires access
// to that S3 bucket for input or output artifacts.
