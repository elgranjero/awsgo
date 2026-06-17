package sns

// Publish is generated as a reference stub.
// Executable command wiring lives under cmd/sns.go.
//
// Sends a message to an Amazon SNS topic, a text message (SMS message) directly
// to a phone number, or a message to a mobile platform endpoint (when you specify
// the TargetArn ).
//
// If you send a message to a topic, Amazon SNS delivers the message to each
// endpoint that is subscribed to the topic. The format of the message depends on
// the notification protocol for each subscribed endpoint.
//
// When a messageId is returned, the message is saved and Amazon SNS immediately
// delivers it to subscribers.
//
// To use the Publish action for publishing a message to a mobile endpoint, such
// as an app on a Kindle device or mobile phone, you must specify the EndpointArn
// for the TargetArn parameter. The EndpointArn is returned when making a call with
// the CreatePlatformEndpoint action.
//
// For more information about formatting messages, see [Send Custom Platform-Specific Payloads in Messages to Mobile Devices].
//
// You can publish messages only to topics and endpoints in the same Amazon Web
// Services Region.
//
// [Send Custom Platform-Specific Payloads in Messages to Mobile Devices]: https://docs.aws.amazon.com/sns/latest/dg/mobile-push-send-custommessage.html
