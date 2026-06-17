package sfn

// GetActivityTask is generated as a reference stub.
// Executable command wiring lives under cmd/sfn.go.
//
// Used by workers to retrieve a task (with the specified activity ARN) which has
// been scheduled for execution by a running state machine. This initiates a long
// poll, where the service holds the HTTP connection open and responds as soon as a
// task becomes available (i.e. an execution of a task of this type is needed.) The
// maximum time the service holds on to the request before responding is 60
// seconds. If no task is available within 60 seconds, the poll returns a taskToken
// with a null string.
//
// This API action isn't logged in CloudTrail.
//
// Workers should set their client side socket timeout to at least 65 seconds (5
// seconds higher than the maximum time the service may hold the poll request).
//
// Polling with GetActivityTask can cause latency in some implementations. See [Avoid Latency When Polling for Activity Tasks] in
// the Step Functions Developer Guide.
//
// [Avoid Latency When Polling for Activity Tasks]: https://docs.aws.amazon.com/step-functions/latest/dg/bp-activity-pollers.html
