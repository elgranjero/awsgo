package connect

// TransferContact is generated as a reference stub.
// Executable command wiring lives under cmd/connect.go.
//
// Transfers TASK or EMAIL contacts from one agent or queue to another agent or
// queue at any point after a contact is created. You can transfer a contact to
// another queue by providing the flow which orchestrates the contact to the
// destination queue. This gives you more control over contact handling and helps
// you adhere to the service level agreement (SLA) guaranteed to your customers.
//
// Note the following requirements:
//
// - Transfer is only supported for TASK and EMAIL contacts.
//
// - Do not use both QueueId and UserId in the same call.
//
// - The following flow types are supported: Inbound flow, Transfer to agent
// flow, and Transfer to queue flow.
//
// - The TransferContact API can be called only on active contacts.
//
// - A contact cannot be transferred more than 11 times.
