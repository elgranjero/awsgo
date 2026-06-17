package datapipeline

// DeactivatePipeline is generated as a reference stub.
// Executable command wiring lives under cmd/datapipeline.go.
//
// Deactivates the specified running pipeline. The pipeline is set to the
// DEACTIVATING state until the deactivation process completes.
//
// To resume a deactivated pipeline, use ActivatePipeline. By default, the pipeline resumes from
// the last completed execution. Optionally, you can specify the date and time to
// resume the pipeline.
