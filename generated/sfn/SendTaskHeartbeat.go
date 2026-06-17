package sfn

// SendTaskHeartbeat is generated as a reference stub.
// Executable command wiring lives under cmd/sfn.go.
//
// Used by activity workers and Task states using the [callback] pattern, and optionally
// Task states using the [job run]pattern to report to Step Functions that the task
// represented by the specified taskToken is still making progress. This action
// resets the Heartbeat clock. The Heartbeat threshold is specified in the state
// machine's Amazon States Language definition ( HeartbeatSeconds ). This action
// does not in itself create an event in the execution history. However, if the
// task times out, the execution history contains an ActivityTimedOut entry for
// activities, or a TaskTimedOut entry for tasks using the [job run] or [callback] pattern.
//
// The Timeout of a task, defined in the state machine's Amazon States Language
// definition, is its maximum allowed duration, regardless of the number of SendTaskHeartbeat
// requests received. Use HeartbeatSeconds to configure the timeout interval for
// heartbeats.
//
// [callback]: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token
// [job run]: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-sync
