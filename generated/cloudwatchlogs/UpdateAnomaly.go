package cloudwatchlogs

// UpdateAnomaly is generated as a reference stub.
// Executable command wiring lives under cmd/cloudwatchlogs.go.
//
// Use this operation to suppress anomaly detection for a specified anomaly or
// pattern. If you suppress an anomaly, CloudWatch Logs won't report new
// occurrences of that anomaly and won't update that anomaly with new data. If you
// suppress a pattern, CloudWatch Logs won't report any anomalies related to that
// pattern.
//
// You must specify either anomalyId or patternId , but you can't specify both
// parameters in the same operation.
//
// If you have previously used this operation to suppress detection of a pattern
// or anomaly, you can use it again to cause CloudWatch Logs to end the
// suppression. To do this, use this operation and specify the anomaly or pattern
// to stop suppressing, and omit the suppressionType and suppressionPeriod
// parameters.
