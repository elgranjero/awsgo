package kinesis

// UpdateStreamMode is generated as a reference stub.
// Executable command wiring lives under cmd/kinesis.go.
//
// Updates the capacity mode of the data stream. Currently, in Kinesis Data
//
// Streams, you can choose between an on-demand capacity mode and a provisioned
// capacity mode for your data stream.
//
// If you'd still like to proactively scale your on-demand data stream’s capacity,
// you can unlock the warm throughput feature for on-demand data streams by
// enabling MinimumThroughputBillingCommitment for your account. Once your account
// has MinimumThroughputBillingCommitment enabled, you can specify the warm
// throughput in MiB per second that your stream can support in writes.
