package rum

// BatchCreateRumMetricDefinitions is generated as a reference stub.
// Executable command wiring lives under cmd/rum.go.
//
// Specifies the extended metrics and custom metrics that you want a CloudWatch
// RUM app monitor to send to a destination. Valid destinations include CloudWatch
// and Evidently.
//
// By default, RUM app monitors send some metrics to CloudWatch. These default
// metrics are listed in [CloudWatch metrics that you can collect with CloudWatch RUM].
//
// In addition to these default metrics, you can choose to send extended metrics,
// custom metrics, or both.
//
// - Extended metrics let you send metrics with additional dimensions that
// aren't included in the default metrics. You can also send extended metrics to
// both Evidently and CloudWatch. The valid dimension names for the additional
// dimensions for extended metrics are BrowserName , CountryCode , DeviceType ,
// FileType , OSName , and PageId . For more information, see [Extended metrics that you can send to CloudWatch and CloudWatch Evidently].
//
// - Custom metrics are metrics that you define. You can send custom metrics to
// CloudWatch. CloudWatch Evidently, or both. With custom metrics, you can use any
// metric name and namespace. To derive the metrics, you can use any custom events,
// built-in events, custom attributes, or default attributes.
//
// You can't send custom metrics to the AWS/RUM namespace. You must send custom
//
// metrics to a custom namespace that you define. The namespace that you use can't
// start with AWS/ . CloudWatch RUM prepends RUM/CustomMetrics/ to the custom
// namespace that you define, so the final namespace for your metrics in CloudWatch
// is RUM/CustomMetrics/your-custom-namespace .
//
// The maximum number of metric definitions that you can specify in one
// BatchCreateRumMetricDefinitions operation is 200.
//
// The maximum number of metric definitions that one destination can contain is
// 2000.
//
// Extended metrics sent to CloudWatch and RUM custom metrics are charged as
// CloudWatch custom metrics. Each combination of additional dimension name and
// dimension value counts as a custom metric. For more information, see [Amazon CloudWatch Pricing].
//
// You must have already created a destination for the metrics before you send
// them. For more information, see [PutRumMetricsDestination].
//
// If some metric definitions specified in a BatchCreateRumMetricDefinitions
// operations are not valid, those metric definitions fail and return errors, but
// all valid metric definitions in the same operation still succeed.
//
// [PutRumMetricsDestination]: https://docs.aws.amazon.com/cloudwatchrum/latest/APIReference/API_PutRumMetricsDestination.html
// [Extended metrics that you can send to CloudWatch and CloudWatch Evidently]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-RUM-vended-metrics.html
// [CloudWatch metrics that you can collect with CloudWatch RUM]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-RUM-metrics.html
// [Amazon CloudWatch Pricing]: https://aws.amazon.com/cloudwatch/pricing/
