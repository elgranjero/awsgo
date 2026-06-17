package connectparticipant

// GetTranscript is generated as a reference stub.
// Executable command wiring lives under cmd/connectparticipant.go.
//
// Retrieves a transcript of the session, including details about any attachments.
// For information about accessing past chat contact transcripts for a persistent
// chat, see [Enable persistent chat].
//
// For security recommendations, see [Amazon Connect Chat security best practices].
//
// If you have a process that consumes events in the transcript of an chat that
// has ended, note that chat transcripts contain the following event content types
// if the event has occurred during the chat session:
//
// - application/vnd.amazonaws.connect.event.participant.invited
//
// - application/vnd.amazonaws.connect.event.participant.joined
//
// - application/vnd.amazonaws.connect.event.participant.left
//
// - application/vnd.amazonaws.connect.event.chat.ended
//
// - application/vnd.amazonaws.connect.event.transfer.succeeded
//
// - application/vnd.amazonaws.connect.event.transfer.failed
//
// ConnectionToken is used for invoking this API instead of ParticipantToken .
//
// The Amazon Connect Participant Service APIs do not use [Signature Version 4 authentication].
//
// [Enable persistent chat]: https://docs.aws.amazon.com/connect/latest/adminguide/chat-persistence.html
// [Signature Version 4 authentication]: https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html
// [Amazon Connect Chat security best practices]: https://docs.aws.amazon.com/connect/latest/adminguide/security-best-practices.html#bp-security-chat
