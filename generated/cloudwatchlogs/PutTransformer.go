package cloudwatchlogs

// PutTransformer is generated as a reference stub.
// Executable command wiring lives under cmd/cloudwatchlogs.go.
//
// Creates or updates a log transformer for a single log group. You use log
// transformers to transform log events into a different format, making them easier
// for you to process and analyze. You can also transform logs from different
// sources into standardized formats that contains relevant, source-specific
// information.
//
// After you have created a transformer, CloudWatch Logs performs the
// transformations at the time of log ingestion. You can then refer to the
// transformed versions of the logs during operations such as querying with
// CloudWatch Logs Insights or creating metric filters or subscription filers.
//
// You can also use a transformer to copy metadata from metadata keys into the log
// events themselves. This metadata can include log group name, log stream name,
// account ID and Region.
//
// A transformer for a log group is a series of processors, where each processor
// applies one type of transformation to the log events ingested into this log
// group. The processors work one after another, in the order that you list them,
// like a pipeline. For more information about the available processors to use in a
// transformer, see [Processors that you can use].
//
// Having log events in standardized format enables visibility across your
// applications for your log analysis, reporting, and alarming needs. CloudWatch
// Logs provides transformation for common log types with out-of-the-box
// transformation templates for major Amazon Web Services log sources such as VPC
// flow logs, Lambda, and Amazon RDS. You can use pre-built transformation
// templates or create custom transformation policies.
//
// You can create transformers only for the log groups in the Standard log class.
//
// You can also set up a transformer at the account level. For more information,
// see [PutAccountPolicy]. If there is both a log-group level transformer created with PutTransformer
// and an account-level transformer that could apply to the same log group, the log
// group uses only the log-group level transformer. It ignores the account-level
// transformer.
//
// [Processors that you can use]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation.html#CloudWatch-Logs-Transformation-Processors
// [PutAccountPolicy]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutAccountPolicy.html
