package connect

// StartChatContact is generated as a reference stub.
// Executable command wiring lives under cmd/connect.go.
//
// Initiates a flow to start a new chat for the customer. Response of this API
// provides a token required to obtain credentials from the [CreateParticipantConnection]API in the Amazon
// Connect Participant Service.
//
// When a new chat contact is successfully created, clients must subscribe to the
// participant’s connection for the created chat within 5 minutes. This is achieved
// by invoking [CreateParticipantConnection]with WEBSOCKET and CONNECTION_CREDENTIALS.
//
// A 429 error occurs in the following situations:
//
// - API rate limit is exceeded. API TPS throttling returns a TooManyRequests
// exception.
//
// - The [quota for concurrent active chats]is exceeded. Active chat throttling returns a LimitExceededException .
//
// If you use the ChatDurationInMinutes parameter and receive a 400 error, your
// account may not support the ability to configure custom chat durations. For more
// information, contact Amazon Web Services Support.
//
// For more information about chat, see the following topics in the Amazon Connect
// Administrator Guide:
//
// [Concepts: Web and mobile messaging capabilities in Amazon Connect]
//
// [Amazon Connect Chat security best practices]
//
// [CreateParticipantConnection]: https://docs.aws.amazon.com/connect-participant/latest/APIReference/API_CreateParticipantConnection.html
// [Concepts: Web and mobile messaging capabilities in Amazon Connect]: https://docs.aws.amazon.com/connect/latest/adminguide/web-and-mobile-chat.html
// [quota for concurrent active chats]: https://docs.aws.amazon.com/connect/latest/adminguide/amazon-connect-service-limits.html
// [Amazon Connect Chat security best practices]: https://docs.aws.amazon.com/connect/latest/adminguide/security-best-practices.html#bp-security-chat
