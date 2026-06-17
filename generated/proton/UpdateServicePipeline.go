package proton

// UpdateServicePipeline is generated as a reference stub.
// Executable command wiring lives under cmd/proton.go.
//
// Update the service pipeline.
//
// There are four modes for updating a service pipeline. The deploymentType field
// defines the mode.
//
// NONE
//
// In this mode, a deployment doesn't occur. Only the requested metadata
// parameters are updated.
//
// CURRENT_VERSION
//
// In this mode, the service pipeline is deployed and updated with the new spec
// that you provide. Only requested parameters are updated. Don’t include major or
// minor version parameters when you use this deployment-type .
//
// MINOR_VERSION
//
// In this mode, the service pipeline is deployed and updated with the published,
// recommended (latest) minor version of the current major version in use, by
// default. You can specify a different minor version of the current major version
// in use.
//
// MAJOR_VERSION
//
// In this mode, the service pipeline is deployed and updated with the published,
// recommended (latest) major and minor version of the current template by default.
// You can specify a different major version that's higher than the major version
// in use and a minor version.
//
// Deprecated: AWS Proton is not accepting new customers.
