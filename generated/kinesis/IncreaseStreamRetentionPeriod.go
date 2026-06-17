package kinesis

// IncreaseStreamRetentionPeriod is generated as a reference stub.
// Executable command wiring lives under cmd/kinesis.go.
//
// Increases the Kinesis data stream's retention period, which is the length of
// time data records are accessible after they are added to the stream. The maximum
// value of a stream's retention period is 8760 hours (365 days).
//
// When invoking this API, you must use either the StreamARN or the StreamName
// parameter, or both. It is recommended that you use the StreamARN input
// parameter when you invoke this API.
//
// If you choose a longer stream retention period, this operation increases the
// time period during which records that have not yet expired are accessible.
// However, it does not make previous, expired data (older than the stream's
// previous retention period) accessible after the operation has been called. For
// example, if a stream's retention period is set to 24 hours and is increased to
// 168 hours, any data that is older than 24 hours remains inaccessible to consumer
// applications.
