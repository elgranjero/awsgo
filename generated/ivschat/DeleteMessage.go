package ivschat

// DeleteMessage is generated as a reference stub.
// Executable command wiring lives under cmd/ivschat.go.
//
// Sends an event to a specific room which directs clients to delete a specific
// message; that is, unrender it from view and delete it from the client’s chat
// history. This event’s EventName is aws:DELETE_MESSAGE . This replicates the [DeleteMessage]
// WebSocket operation in the Amazon IVS Chat Messaging API.
//
// [DeleteMessage]: https://docs.aws.amazon.com/ivs/latest/chatmsgapireference/actions-deletemessage-publish.html
