package sns

// ListPlatformApplications is generated as a reference stub.
// Executable command wiring lives under cmd/sns.go.
//
// Lists the platform application objects for the supported push notification
// services, such as APNS and GCM (Firebase Cloud Messaging). The results for
// ListPlatformApplications are paginated and return a limited list of
// applications, up to 100. If additional records are available after the first
// page results, then a NextToken string will be returned. To receive the next
// page, you call ListPlatformApplications using the NextToken string received
// from the previous call. When there are no more records to return, NextToken
// will be null. For more information, see [Using Amazon SNS Mobile Push Notifications].
//
// This action is throttled at 15 transactions per second (TPS).
//
// [Using Amazon SNS Mobile Push Notifications]: https://docs.aws.amazon.com/sns/latest/dg/SNSMobilePush.html
