package connect

// CreateQueue is generated as a reference stub.
// Executable command wiring lives under cmd/connect.go.
//
// Creates a new queue for the specified Amazon Connect instance.
//
// - If the phone number is claimed to a traffic distribution group that was
// created in the same Region as the Amazon Connect instance where you are calling
// this API, then you can use a full phone number ARN or a UUID for
// OutboundCallerIdNumberId . However, if the phone number is claimed to a
// traffic distribution group that is in one Region, and you are calling this API
// from an instance in another Amazon Web Services Region that is associated with
// the traffic distribution group, you must provide a full phone number ARN. If a
// UUID is provided in this scenario, you will receive a
// ResourceNotFoundException .
//
// - Only use the phone number ARN format that doesn't contain instance in the
// path, for example, arn:aws:connect:us-east-1:1234567890:phone-number/uuid .
// This is the same ARN format that is returned when you call the [ListPhoneNumbersV2]API.
//
// - If you plan to use IAM policies to allow/deny access to this API for phone
// number resources claimed to a traffic distribution group, see [Allow or Deny queue API actions for phone numbers in a replica Region].
//
// [ListPhoneNumbersV2]: https://docs.aws.amazon.com/connect/latest/APIReference/API_ListPhoneNumbersV2.html
// [Allow or Deny queue API actions for phone numbers in a replica Region]: https://docs.aws.amazon.com/connect/latest/adminguide/security_iam_resource-level-policy-examples.html#allow-deny-queue-actions-replica-region
