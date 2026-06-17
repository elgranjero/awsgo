package proton

// CancelServicePipelineDeployment is generated as a reference stub.
// Executable command wiring lives under cmd/proton.go.
//
// Attempts to cancel a service pipeline deployment on an UpdateServicePipeline action, if the
// deployment is IN_PROGRESS . For more information, see [Update a service pipeline] in the Proton User guide.
//
// The following list includes potential cancellation scenarios.
//
// - If the cancellation attempt succeeds, the resulting deployment state is
// CANCELLED .
//
// - If the cancellation attempt fails, the resulting deployment state is FAILED .
//
// - If the current UpdateServicePipelineaction succeeds before the cancellation attempt starts, the
// resulting deployment state is SUCCEEDED and the cancellation attempt has no
// effect.
//
// Deprecated: AWS Proton is not accepting new customers.
//
// [Update a service pipeline]: https://docs.aws.amazon.com/proton/latest/userguide/ag-svc-pipeline-update.html
