package omics

// CreateRunCache is generated as a reference stub.
// Executable command wiring lives under cmd/omics.go.
//
// Creates a run cache to store and reference task outputs from completed private
// runs. Specify an Amazon S3 location where Amazon Web Services HealthOmics saves
// the cached data. This data must be immediately accessible and not in an archived
// state. You can save intermediate task files to a run cache if they are declared
// as task outputs in the workflow definition file.
//
// For more information, see [Call caching] and [Creating a run cache] in the Amazon Web Services HealthOmics User
// Guide.
//
// [Call caching]: https://docs.aws.amazon.com/omics/latest/dev/workflows-call-caching.html
// [Creating a run cache]: https://docs.aws.amazon.com/omics/latest/dev/workflow-cache-create.html
