package connect

// StopContactRecording is generated as a reference stub.
// Executable command wiring lives under cmd/connect.go.
//
// Stops recording a call when a contact is being recorded. StopContactRecording
// is a one-time action. If you use StopContactRecording to stop recording an
// ongoing call, you can't use StartContactRecording to restart it. For scenarios
// where the recording has started and you want to suspend it for sensitive
// information (for example, to collect a credit card number), and then restart it,
// use SuspendContactRecording and ResumeContactRecording.
//
// Only voice recordings are supported at this time.
