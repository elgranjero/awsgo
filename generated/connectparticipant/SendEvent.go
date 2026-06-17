package connectparticipant

// SendEvent is generated as a reference stub.
// Executable command wiring lives under cmd/connectparticipant.go.
//
// The application/vnd.amazonaws.connect.event.connection.acknowledged ContentType
// is no longer maintained since December 31, 2024. This event has been migrated to
// the [CreateParticipantConnection]API using the ConnectParticipant field.
//
// Sends an event. Message receipts are not supported when there are more than two
// active participants in the chat. Using the SendEvent API for message receipts
// when a supervisor is barged-in will result in a conflict exception.
//
// For security recommendations, see [Amazon Connect Chat security best practices].
//
// ConnectionToken is used for invoking this API instead of ParticipantToken .
//
// The Amazon Connect Participant Service APIs do not use [Signature Version 4 authentication].
//
// [CreateParticipantConnection]: https://docs.aws.amazon.com/connect-participant/latest/APIReference/API_CreateParticipantConnection.html
// [Signature Version 4 authentication]: https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html
// [Amazon Connect Chat security best practices]: https://docs.aws.amazon.com/connect/latest/adminguide/security-best-practices.html#bp-security-chat
