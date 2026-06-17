package sns

// CreatePlatformEndpoint is generated as a reference stub.
// Executable command wiring lives under cmd/sns.go.
//
// Creates an endpoint for a device and mobile app on one of the supported push
// notification services, such as GCM (Firebase Cloud Messaging) and APNS.
// CreatePlatformEndpoint requires the PlatformApplicationArn that is returned
// from CreatePlatformApplication . You can use the returned EndpointArn to send a
// message to a mobile app or by the Subscribe action for subscription to a topic.
// The CreatePlatformEndpoint action is idempotent, so if the requester already
// owns an endpoint with the same device token and attributes, that endpoint's ARN
// is returned without creating a new endpoint. For more information, see [Using Amazon SNS Mobile Push Notifications].
//
// When using CreatePlatformEndpoint with Baidu, two attributes must be provided:
// ChannelId and UserId. The token field must also contain the ChannelId. For more
// information, see [Creating an Amazon SNS Endpoint for Baidu].
//
// [Creating an Amazon SNS Endpoint for Baidu]: https://docs.aws.amazon.com/sns/latest/dg/SNSMobilePushBaiduEndpoint.html
// [Using Amazon SNS Mobile Push Notifications]: https://docs.aws.amazon.com/sns/latest/dg/SNSMobilePush.html
