package connect

// StartContactRecording is generated as a reference stub.
// Executable command wiring lives under cmd/connect.go.
//
// Starts recording the contact:
//
// - If the API is called before the agent joins the call, recording starts when
// the agent joins the call.
//
// - If the API is called after the agent joins the call, recording starts at
// the time of the API call.
//
// StartContactRecording is a one-time action. For example, if you use
// StopContactRecording to stop recording an ongoing call, you can't use
// StartContactRecording to restart it. For scenarios where the recording has
// started and you want to suspend and resume it, such as when collecting sensitive
// information (for example, a credit card number), use SuspendContactRecording and
// ResumeContactRecording.
//
// You can use this API to override the recording behavior configured in the [Set recording behavior]
// block.
//
// Only voice recordings are supported at this time.
//
// [Set recording behavior]: https://docs.aws.amazon.com/connect/latest/adminguide/set-recording-behavior.html
