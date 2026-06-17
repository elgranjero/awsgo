package connectparticipant

// GetAttachment is generated as a reference stub.
// Executable command wiring lives under cmd/connectparticipant.go.
//
// Provides a pre-signed URL for download of a completed attachment. This is an
// asynchronous API for use with active contacts.
//
// For security recommendations, see [Amazon Connect Chat security best practices].
//
// - The participant role CUSTOM_BOT is not permitted to access attachments
// customers may upload. An AccessDeniedException can indicate that the
// participant may be a CUSTOM_BOT, and it doesn't have access to attachments.
//
// - ConnectionToken is used for invoking this API instead of ParticipantToken .
//
// The Amazon Connect Participant Service APIs do not use [Signature Version 4 authentication].
//
// [Signature Version 4 authentication]: https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html
// [Amazon Connect Chat security best practices]: https://docs.aws.amazon.com/connect/latest/adminguide/security-best-practices.html#bp-security-chat
