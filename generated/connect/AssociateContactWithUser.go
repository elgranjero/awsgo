package connect

// AssociateContactWithUser is generated as a reference stub.
// Executable command wiring lives under cmd/connect.go.
//
// Associates a queued contact with an agent.
//
// # Use cases
//
// Following are common uses cases for this API:
//
// - Programmatically assign queued contacts to available users.
//
// - Leverage the IAM context key connect:PreferredUserArn to restrict contact
// association to specific preferred user.
//
// Important things to know
//
// - Use this API with chat, email, and task contacts. It does not support voice
// contacts.
//
// - Use it to associate contacts with users regardless of their current state,
// including custom states. Ensure your application logic accounts for user
// availability before making associations.
//
// - It honors the IAM context key connect:PreferredUserArn to prevent
// unauthorized contact associations.
//
// - It respects the IAM context key connect:PreferredUserArn to enforce
// authorization controls and prevent unauthorized contact associations. Verify
// that your IAM policies are properly configured to support your intended use
// cases.
//
// - The service quota Queues per routing profile per instance applies to
// manually assigned queues, too. For more information about this quota, see [Amazon Connect quotas]in
// the Amazon Connect Administrator Guide.
//
// Endpoints: See [Amazon Connect endpoints and quotas].
//
// [Amazon Connect endpoints and quotas]: https://docs.aws.amazon.com/general/latest/gr/connect_region.html
// [Amazon Connect quotas]: https://docs.aws.amazon.com/connect/latest/adminguide/amazon-connect-service-limits.html#connect-quotas
