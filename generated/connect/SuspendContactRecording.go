package connect

// SuspendContactRecording is generated as a reference stub.
// Executable command wiring lives under cmd/connect.go.
//
// When a contact is being recorded, this API suspends recording whatever is
// selected in the flow configuration: call (IVR or agent), screen, or both. If
// only call recording or only screen recording is enabled, then it would be
// suspended. For example, you might suspend the screen recording while collecting
// sensitive information, such as a credit card number. Then use [ResumeContactRecording]to restart
// recording the screen.
//
// The period of time that the recording is suspended is filled with silence in
// the final recording.
//
// Voice (IVR, agent) and screen recordings are supported.
//
// [ResumeContactRecording]: https://docs.aws.amazon.com/connect/latest/APIReference/API_ResumeContactRecording.html
