package connect

// GetCurrentMetricData is generated as a reference stub.
// Executable command wiring lives under cmd/connect.go.
//
// Gets the real-time metric data from the specified Amazon Connect instance.
//
// For a description of each metric, see [Metrics definitions] in the Amazon Connect Administrator
// Guide.
//
// When you make a successful API request, you can expect the following metric
// values in the response:
//
// - Metric value is null: The calculation cannot be performed due to divide by
// zero or insufficient data
//
// - Metric value is a number (including 0) of defined type: The number provided
// is the calculation result
//
// - MetricResult list is empty: The request cannot find any data in the system
//
// The following guidelines can help you work with the API:
//
// - Each dimension in the metric response must contain a value
//
// - Each item in MetricResult must include all requested metrics
//
// - If the response is slow due to large result sets, try these approaches:
//
// - Add filters to reduce the amount of data returned
//
// [Metrics definitions]: https://docs.aws.amazon.com/connect/latest/adminguide/metrics-definitions.html
