package connect

// DescribeContact is generated as a reference stub.
// Executable command wiring lives under cmd/connect.go.
//
// This API is in preview release for Amazon Connect and is subject to change.
//
// Describes the specified contact.
//
// # Use cases
//
// Following are common uses cases for this API:
//
// - Retrieve contact information such as the caller's phone number and the
// specific number the caller dialed to integrate into custom monitoring or custom
// agent experience solutions.
//
// - Detect when a customer chat session disconnects due to a network issue on
// the agent's end. Use the DisconnectReason field in the [ContactTraceRecord]to detect this event
// and then re-queue the chat for followup.
//
// - Identify after contact work (ACW) duration and call recordings information
// when a COMPLETED event is received by using the [contact event stream].
//
// Important things to know
//
// - SystemEndpoint is not populated for contacts with initiation method of
// MONITOR, QUEUE_TRANSFER, or CALLBACK
//
// - Contact information remains available in Amazon Connect for 24 months from
// the InitiationTimestamp , and then it is deleted. Only contact information
// that is available in Amazon Connect is returned by this API.
//
// Endpoints: See [Amazon Connect endpoints and quotas].
//
// [ContactTraceRecord]: https://docs.aws.amazon.com/connect/latest/adminguide/ctr-data-model.html#ctr-ContactTraceRecord
// [contact event stream]: https://docs.aws.amazon.com/connect/latest/adminguide/contact-events.html
// [Amazon Connect endpoints and quotas]: https://docs.aws.amazon.com/general/latest/gr/connect_region.html
