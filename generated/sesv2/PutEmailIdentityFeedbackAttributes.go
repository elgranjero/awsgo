package sesv2

// PutEmailIdentityFeedbackAttributes is generated as a reference stub.
// Executable command wiring lives under cmd/sesv2.go.
//
// Used to enable or disable feedback forwarding for an identity. This setting
// determines what happens when an identity is used to send an email that results
// in a bounce or complaint event.
//
// If the value is true , you receive email notifications when bounce or complaint
// events occur. These notifications are sent to the address that you specified in
// the Return-Path header of the original email.
//
// You're required to have a method of tracking bounces and complaints. If you
// haven't set up another mechanism for receiving bounce or complaint notifications
// (for example, by setting up an event destination), you receive an email
// notification when these events occur (even if this setting is disabled).
