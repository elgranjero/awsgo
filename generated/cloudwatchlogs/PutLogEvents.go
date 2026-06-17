package cloudwatchlogs

// PutLogEvents is generated as a reference stub.
// Executable command wiring lives under cmd/cloudwatchlogs.go.
//
// Uploads a batch of log events to the specified log stream.
//
// The sequence token is now ignored in PutLogEvents actions. PutLogEvents actions
// are always accepted and never return InvalidSequenceTokenException or
// DataAlreadyAcceptedException even if the sequence token is not valid. You can
// use parallel PutLogEvents actions on the same log stream.
//
// The batch of events must satisfy the following constraints:
//
// - The maximum batch size is 1,048,576 bytes. This size is calculated as the
// sum of all event messages in UTF-8, plus 26 bytes for each log event.
//
// - Events more than 2 hours in the future are rejected while processing
// remaining valid events.
//
// - Events older than 14 days or preceding the log group's retention period are
// rejected while processing remaining valid events.
//
// - The log events in the batch must be in chronological order by their
// timestamp. The timestamp is the time that the event occurred, expressed as the
// number of milliseconds after Jan 1, 1970 00:00:00 UTC . (In Amazon Web
// Services Tools for PowerShell and the Amazon Web Services SDK for .NET, the
// timestamp is specified in .NET format: yyyy-mm-ddThh:mm:ss . For example,
// 2017-09-15T13:45:30 .)
//
// - A batch of log events in a single request must be in a chronological order.
// Otherwise, the operation fails.
//
// - Each log event can be no larger than 1 MB.
//
// - The maximum number of log events in a batch is 10,000.
//
// - For valid events (within 14 days in the past to 2 hours in future), the
// time span in a single batch cannot exceed 24 hours. Otherwise, the operation
// fails.
//
// The quota of five requests per second per log stream has been removed. Instead,
// PutLogEvents actions are throttled based on a per-second per-account quota. You
// can request an increase to the per-second throttling quota by using the Service
// Quotas service.
//
// If a call to PutLogEvents returns "UnrecognizedClientException" the most likely
// cause is a non-valid Amazon Web Services access key ID or secret key.
