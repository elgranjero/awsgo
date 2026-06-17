package connect

// SendChatIntegrationEvent is generated as a reference stub.
// Executable command wiring lives under cmd/connect.go.
//
// Processes chat integration events from Amazon Web Services or external
// integrations to Amazon Connect. A chat integration event includes:
//
// - SourceId, DestinationId, and Subtype: a set of identifiers, uniquely
// representing a chat
//
// - ChatEvent: details of the chat action to perform such as sending a message,
// event, or disconnecting from a chat
//
// When a chat integration event is sent with chat identifiers that do not map to
// an active chat contact, a new chat contact is also created before handling chat
// action.
//
// Access to this API is currently restricted to Amazon Web Services End User
// Messaging for supporting SMS integration.
