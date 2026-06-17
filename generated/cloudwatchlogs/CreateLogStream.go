package cloudwatchlogs

// CreateLogStream is generated as a reference stub.
// Executable command wiring lives under cmd/cloudwatchlogs.go.
//
// Creates a log stream for the specified log group. A log stream is a sequence of
// log events that originate from a single source, such as an application instance
// or a resource that is being monitored.
//
// There is no limit on the number of log streams that you can create for a log
// group. There is a limit of 50 TPS on CreateLogStream operations, after which
// transactions are throttled.
//
// You must use the following guidelines when naming a log stream:
//
// - Log stream names must be unique within the log group.
//
// - Log stream names can be between 1 and 512 characters long.
//
// - Don't use ':' (colon) or '*' (asterisk) characters.
