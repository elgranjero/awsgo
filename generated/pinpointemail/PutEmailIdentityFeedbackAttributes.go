package pinpointemail

// PutEmailIdentityFeedbackAttributes is generated as a reference stub.
// Executable command wiring lives under cmd/pinpointemail.go.
//
// Used to enable or disable feedback forwarding for an identity. This setting
// determines what happens when an identity is used to send an email that results
// in a bounce or complaint event.
//
// When you enable feedback forwarding, Amazon Pinpoint sends you email
// notifications when bounce or complaint events occur. Amazon Pinpoint sends this
// notification to the address that you specified in the Return-Path header of the
// original email.
//
// When you disable feedback forwarding, Amazon Pinpoint sends notifications
// through other mechanisms, such as by notifying an Amazon SNS topic. You're
// required to have a method of tracking bounces and complaints. If you haven't set
// up another mechanism for receiving bounce or complaint notifications, Amazon
// Pinpoint sends an email notification when these events occur (even if this
// setting is disabled).
