package connect

// UpdateContactRoutingData is generated as a reference stub.
// Executable command wiring lives under cmd/connect.go.
//
// Updates routing priority and age on the contact (QueuePriority and
// QueueTimeAdjustmentInSeconds). These properties can be used to change a
// customer's position in the queue. For example, you can move a contact to the
// back of the queue by setting a lower routing priority relative to other contacts
// in queue; or you can move a contact to the front of the queue by increasing the
// routing age which will make the contact look artificially older and therefore
// higher up in the first-in-first-out routing order. Note that adjusting the
// routing age of a contact affects only its position in queue, and not its actual
// queue wait time as reported through metrics. These properties can also be
// updated by using [the Set routing priority / age flow block].
//
// Either QueuePriority or QueueTimeAdjustmentInSeconds should be provided within
// the request body, but not both.
//
// [the Set routing priority / age flow block]: https://docs.aws.amazon.com/connect/latest/adminguide/change-routing-priority.html
