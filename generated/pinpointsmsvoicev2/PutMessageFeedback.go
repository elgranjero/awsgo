package pinpointsmsvoicev2

// PutMessageFeedback is generated as a reference stub.
// Executable command wiring lives under cmd/pinpointsmsvoicev2.go.
//
// Set the MessageFeedbackStatus as RECEIVED or FAILED for the passed in
// MessageId.
//
// If you use message feedback then you must update message feedback record. When
// you receive a signal that a user has received the message you must use
// PutMessageFeedback to set the message feedback record as RECEIVED ; Otherwise,
// an hour after the message feedback record is set to FAILED .
