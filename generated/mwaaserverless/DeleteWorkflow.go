package mwaaserverless

// DeleteWorkflow is generated as a reference stub.
// Executable command wiring lives under cmd/mwaaserverless.go.
//
// Deletes a workflow and all its versions. This operation permanently removes the
// workflow and cannot be undone. Amazon Managed Workflows for Apache Airflow
// Serverless ensures that all associated resources are properly cleaned up,
// including stopping any running executions, removing scheduled triggers, and
// cleaning up execution history. The deletion process respects the multi-tenant
// isolation boundaries and ensures that no residual data or configurations remain
// that could affect other customers or workflows.
