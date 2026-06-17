package cloudwatch

// ListMetrics is generated as a reference stub.
// Executable command wiring lives under cmd/cloudwatch.go.
//
// List the specified metrics. You can use the returned metrics with [GetMetricData] or [GetMetricStatistics] to get
// statistical data.
//
// Up to 500 results are returned for any one call. To retrieve additional
// results, use the returned token with subsequent calls.
//
// After you create a metric, allow up to 15 minutes for the metric to appear. To
// see metric statistics sooner, use [GetMetricData]or [GetMetricStatistics].
//
// If you are using CloudWatch cross-account observability, you can use this
// operation in a monitoring account and view metrics from the linked source
// accounts. For more information, see [CloudWatch cross-account observability].
//
// ListMetrics doesn't return information about metrics if those metrics haven't
// reported data in the past two weeks. To retrieve those metrics, use [GetMetricData]or [GetMetricStatistics].
//
// [GetMetricData]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_GetMetricData.html
// [GetMetricStatistics]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_GetMetricStatistics.html
// [CloudWatch cross-account observability]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Unified-Cross-Account.html
