package kinesis

// DecreaseStreamRetentionPeriod is generated as a reference stub.
// Executable command wiring lives under cmd/kinesis.go.
//
// Decreases the Kinesis data stream's retention period, which is the length of
// time data records are accessible after they are added to the stream. The minimum
// value of a stream's retention period is 24 hours.
//
// When invoking this API, you must use either the StreamARN or the StreamName
// parameter, or both. It is recommended that you use the StreamARN input
// parameter when you invoke this API.
//
// This operation may result in lost data. For example, if the stream's retention
// period is 48 hours and is decreased to 24 hours, any data already in the stream
// that is older than 24 hours is inaccessible.
