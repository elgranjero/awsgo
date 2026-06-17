package ses

// SetIdentityFeedbackForwardingEnabled is generated as a reference stub.
// Executable command wiring lives under cmd/ses.go.
//
// Given an identity (an email address or a domain), enables or disables whether
// Amazon SES forwards bounce and complaint notifications as email. Feedback
// forwarding can only be disabled when Amazon Simple Notification Service (Amazon
// SNS) topics are specified for both bounces and complaints.
//
// Feedback forwarding does not apply to delivery notifications. Delivery
// notifications are only available through Amazon SNS.
//
// You can execute this operation no more than once per second.
//
// For more information about using notifications with Amazon SES, see the [Amazon SES Developer Guide].
//
// [Amazon SES Developer Guide]: https://docs.aws.amazon.com/ses/latest/dg/monitor-sending-activity-using-notifications.html
