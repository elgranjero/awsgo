package ses

// SetIdentityNotificationTopic is generated as a reference stub.
// Executable command wiring lives under cmd/ses.go.
//
// Sets an Amazon Simple Notification Service (Amazon SNS) topic to use when
// delivering notifications. When you use this operation, you specify a verified
// identity, such as an email address or domain. When you send an email that uses
// the chosen identity in the Source field, Amazon SES sends notifications to the
// topic you specified. You can send bounce, complaint, or delivery notifications
// (or any combination of the three) to the Amazon SNS topic that you specify.
//
// You can execute this operation no more than once per second.
//
// For more information about feedback notification, see the [Amazon SES Developer Guide].
//
// [Amazon SES Developer Guide]: https://docs.aws.amazon.com/ses/latest/dg/monitor-sending-activity-using-notifications.html
