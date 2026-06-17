package cloudwatch

// PutMetricStream is generated as a reference stub.
// Executable command wiring lives under cmd/cloudwatch.go.
//
// Creates or updates a metric stream. Metric streams can automatically stream
// CloudWatch metrics to Amazon Web Services destinations, including Amazon S3, and
// to many third-party solutions.
//
// For more information, see [Using Metric Streams].
//
// To create a metric stream, you must be signed in to an account that has the
// iam:PassRole permission and either the CloudWatchFullAccess policy or the
// cloudwatch:PutMetricStream permission.
//
// When you create or update a metric stream, you choose one of the following:
//
// - Stream metrics from all metric namespaces in the account.
//
// - Stream metrics from all metric namespaces in the account, except for the
// namespaces that you list in ExcludeFilters .
//
// - Stream metrics from only the metric namespaces that you list in
// IncludeFilters .
//
// By default, a metric stream always sends the MAX , MIN , SUM , and SAMPLECOUNT
// statistics for each metric that is streamed. You can use the
// StatisticsConfigurations parameter to have the metric stream send additional
// statistics in the stream. Streaming additional statistics incurs additional
// costs. For more information, see [Amazon CloudWatch Pricing].
//
// When you use PutMetricStream to create a new metric stream, the stream is
// created in the running state. If you use it to update an existing stream, the
// state of the stream is not changed.
//
// If you are using CloudWatch cross-account observability and you create a metric
// stream in a monitoring account, you can choose whether to include metrics from
// source accounts in the stream. For more information, see [CloudWatch cross-account observability].
//
// [Using Metric Streams]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Metric-Streams.html
// [CloudWatch cross-account observability]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Unified-Cross-Account.html
// [Amazon CloudWatch Pricing]: https://aws.amazon.com/cloudwatch/pricing/
