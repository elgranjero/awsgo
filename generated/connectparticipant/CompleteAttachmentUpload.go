package connectparticipant

// CompleteAttachmentUpload is generated as a reference stub.
// Executable command wiring lives under cmd/connectparticipant.go.
//
// Allows you to confirm that the attachment has been uploaded using the
// pre-signed URL provided in StartAttachmentUpload API. A conflict exception is
// thrown when an attachment with that identifier is already being uploaded.
//
// For security recommendations, see [Amazon Connect Chat security best practices].
//
// ConnectionToken is used for invoking this API instead of ParticipantToken .
//
// The Amazon Connect Participant Service APIs do not use [Signature Version 4 authentication].
//
// [Signature Version 4 authentication]: https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html
// [Amazon Connect Chat security best practices]: https://docs.aws.amazon.com/connect/latest/adminguide/security-best-practices.html#bp-security-chat
