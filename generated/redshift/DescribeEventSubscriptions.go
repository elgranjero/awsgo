package redshift

// DescribeEventSubscriptions is generated as a reference stub.
// Executable command wiring lives under cmd/redshift.go.
//
// Lists descriptions of all the Amazon Redshift event notification subscriptions
// for a customer account. If you specify a subscription name, lists the
// description for that subscription.
//
// If you specify both tag keys and tag values in the same request, Amazon
// Redshift returns all event notification subscriptions that match any combination
// of the specified keys and values. For example, if you have owner and environment
// for tag keys, and admin and test for tag values, all subscriptions that have
// any combination of those values are returned.
//
// If both tag keys and values are omitted from the request, subscriptions are
// returned regardless of whether they have tag keys or values associated with
// them.
