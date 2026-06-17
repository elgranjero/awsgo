package proton

// CancelEnvironmentDeployment is generated as a reference stub.
// Executable command wiring lives under cmd/proton.go.
//
// Attempts to cancel an environment deployment on an UpdateEnvironment action, if the deployment
// is IN_PROGRESS . For more information, see [Update an environment] in the Proton User guide.
//
// The following list includes potential cancellation scenarios.
//
// - If the cancellation attempt succeeds, the resulting deployment state is
// CANCELLED .
//
// - If the cancellation attempt fails, the resulting deployment state is FAILED .
//
// - If the current UpdateEnvironmentaction succeeds before the cancellation attempt starts, the
// resulting deployment state is SUCCEEDED and the cancellation attempt has no
// effect.
//
// Deprecated: AWS Proton is not accepting new customers.
//
// [Update an environment]: https://docs.aws.amazon.com/proton/latest/userguide/ag-env-update.html
