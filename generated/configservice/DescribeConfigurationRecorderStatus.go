package configservice

// DescribeConfigurationRecorderStatus is generated as a reference stub.
// Executable command wiring lives under cmd/configservice.go.
//
// Returns the current status of the configuration recorder you specify as well as
// the status of the last recording event for the configuration recorders.
//
// For a detailed status of recording events over time, add your Config events to
// Amazon CloudWatch metrics and use CloudWatch metrics.
//
// If a configuration recorder is not specified, this operation returns the status
// for the customer managed configuration recorder configured for the account, if
// applicable.
//
// When making a request to this operation, you can only specify one configuration
// recorder.
