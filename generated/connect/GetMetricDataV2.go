package connect

// GetMetricDataV2 is generated as a reference stub.
// Executable command wiring lives under cmd/connect.go.
//
// Gets metric data from the specified Amazon Connect instance.
//
// GetMetricDataV2 offers more features than [GetMetricData], the previous version of this API.
// It has new metrics, offers filtering at a metric level, and offers the ability
// to filter and group data by channels, queues, routing profiles, agents, and
// agent hierarchy levels. It can retrieve historical data for the last 3 months,
// at varying intervals. It does not support agent queues.
//
// For a description of the historical metrics that are supported by
// GetMetricDataV2 and GetMetricData , see [Metrics definitions] in the Amazon Connect Administrator
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
// - Narrow the time range of your request
//
// - Add filters to reduce the amount of data returned
//
// [GetMetricData]: https://docs.aws.amazon.com/connect/latest/APIReference/API_GetMetricData.html
// [Metrics definitions]: https://docs.aws.amazon.com/connect/latest/adminguide/metrics-definitions.html
