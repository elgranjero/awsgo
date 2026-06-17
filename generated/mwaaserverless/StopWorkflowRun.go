package mwaaserverless

// StopWorkflowRun is generated as a reference stub.
// Executable command wiring lives under cmd/mwaaserverless.go.
//
// Stops a running workflow execution. This operation terminates all running tasks
// and prevents new tasks from starting. Amazon Managed Workflows for Apache
// Airflow Serverless gracefully shuts down the workflow execution by stopping task
// scheduling and terminating active ECS worker containers. The operation
// transitions the workflow run to a STOPPING state and then to STOPPED once all
// cleanup is complete. In-flight tasks may complete or be terminated depending on
// their current execution state.
