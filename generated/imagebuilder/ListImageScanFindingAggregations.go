package imagebuilder

// ListImageScanFindingAggregations is generated as a reference stub.
// Executable command wiring lives under cmd/imagebuilder.go.
//
// Returns a list of image scan aggregations for your account. You can filter by
// the type of key that Image Builder uses to group results. For example, if you
// want to get a list of findings by severity level for one of your pipelines, you
// might specify your pipeline with the imagePipelineArn filter. If you don't
// specify a filter, Image Builder returns an aggregation for your account.
//
// To streamline results, you can use the following filters in your request:
//
// - accountId
//
// - imageBuildVersionArn
//
// - imagePipelineArn
//
// - vulnerabilityId
