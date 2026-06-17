package iotthingsgraph

// DeleteFlowTemplate is generated as a reference stub.
// Executable command wiring lives under cmd/iotthingsgraph.go.
//
// Deletes a workflow. Any new system or deployment that contains this workflow
// will fail to update or deploy. Existing deployments that contain the workflow
// will continue to run (since they use a snapshot of the workflow taken at the
// time of deployment).
//
// Deprecated: since: 2022-08-30
