package sns

// ListEndpointsByPlatformApplication is generated as a reference stub.
// Executable command wiring lives under cmd/sns.go.
//
// Lists the endpoints and endpoint attributes for devices in a supported push
// notification service, such as GCM (Firebase Cloud Messaging) and APNS. The
// results for ListEndpointsByPlatformApplication are paginated and return a
// limited list of endpoints, up to 100. If additional records are available after
// the first page results, then a NextToken string will be returned. To receive the
// next page, you call ListEndpointsByPlatformApplication again using the
// NextToken string received from the previous call. When there are no more records
// to return, NextToken will be null. For more information, see [Using Amazon SNS Mobile Push Notifications].
//
// This action is throttled at 30 transactions per second (TPS).
//
// [Using Amazon SNS Mobile Push Notifications]: https://docs.aws.amazon.com/sns/latest/dg/SNSMobilePush.html
