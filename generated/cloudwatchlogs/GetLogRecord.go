package cloudwatchlogs

// GetLogRecord is generated as a reference stub.
// Executable command wiring lives under cmd/cloudwatchlogs.go.
//
// Retrieves all of the fields and values of a single log event. All fields are
// retrieved, even if the original query that produced the logRecordPointer
// retrieved only a subset of fields. Fields are returned as field name/field value
// pairs.
//
// The full unparsed log event is returned within (at)message .
