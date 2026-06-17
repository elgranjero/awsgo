package omics

// CreateWorkflowVersion is generated as a reference stub.
// Executable command wiring lives under cmd/omics.go.
//
// Creates a new workflow version for the workflow that you specify with the
// workflowId parameter.
//
// When you create a new version of a workflow, you need to specify the
// configuration for the new version. It doesn't inherit any configuration values
// from the workflow.
//
// Provide a version name that is unique for this workflow. You cannot change the
// name after HealthOmics creates the version.
//
// Don't include any personally identifiable information (PII) in the version
// name. Version names appear in the workflow version ARN.
//
// For more information, see [Workflow versioning in Amazon Web Services HealthOmics] in the Amazon Web Services HealthOmics User Guide.
//
// [Workflow versioning in Amazon Web Services HealthOmics]: https://docs.aws.amazon.com/omics/latest/dev/workflow-versions.html
