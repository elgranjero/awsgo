package pinpointsmsvoicev2

// SendTextMessage is generated as a reference stub.
// Executable command wiring lives under cmd/pinpointsmsvoicev2.go.
//
// Creates a new text message and sends it to a recipient's phone number.
// SendTextMessage only sends an SMS message to one recipient each time it is
// invoked.
//
// SMS throughput limits are measured in Message Parts per Second (MPS). Your MPS
// limit depends on the destination country of your messages, as well as the type
// of phone number (origination number) that you use to send the message. For more
// information about MPS, see [Message Parts per Second (MPS) limits]in the End User Messaging SMS User Guide.
//
// [Message Parts per Second (MPS) limits]: https://docs.aws.amazon.com/sms-voice/latest/userguide/sms-limitations-mps.html
