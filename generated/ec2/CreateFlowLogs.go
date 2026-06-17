package ec2

// CreateFlowLogs is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Creates one or more flow logs to capture information about IP traffic for a
// specific network interface, subnet, or VPC.
//
// Flow log data for a monitored network interface is recorded as flow log
// records, which are log events consisting of fields that describe the traffic
// flow. For more information, see [Flow log records]in the Amazon VPC User Guide.
//
// When publishing to CloudWatch Logs, flow log records are published to a log
// group, and each network interface has a unique log stream in the log group. When
// publishing to Amazon S3, flow log records for all of the monitored network
// interfaces are published to a single log file object that is stored in the
// specified bucket.
//
// For more information, see [VPC Flow Logs] in the Amazon VPC User Guide.
//
// [Flow log records]: https://docs.aws.amazon.com/vpc/latest/userguide/flow-log-records.html
// [VPC Flow Logs]: https://docs.aws.amazon.com/vpc/latest/userguide/flow-logs.html
