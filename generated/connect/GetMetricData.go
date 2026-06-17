package connect

// GetMetricData is generated as a reference stub.
// Executable command wiring lives under cmd/connect.go.
//
// Gets historical metric data from the specified Amazon Connect instance.
//
// For a description of each historical metric, see [Metrics definitions] in the Amazon Connect
// Administrator Guide.
//
// We recommend using the [GetMetricDataV2] API. It provides more flexibility, features, and the
// ability to query longer time ranges than GetMetricData . Use it to retrieve
// historical agent and contact metrics for the last 3 months, at varying
// intervals. You can also use it to build custom dashboards to measure historical
// queue and agent performance. For example, you can track the number of incoming
// contacts for the last 7 days, with data split by day, to see how contact volume
// changed per day of the week.
//
// [GetMetricDataV2]: https://docs.aws.amazon.com/connect/latest/APIReference/API_GetMetricDataV2.html
// [Metrics definitions]: https://docs.aws.amazon.com/connect/latest/adminguide/metrics-definitions.html
