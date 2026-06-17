package connectparticipant

// GetAuthenticationUrl is generated as a reference stub.
// Executable command wiring lives under cmd/connectparticipant.go.
//
// Retrieves the AuthenticationUrl for the current authentication session for the
// AuthenticateCustomer flow block.
//
// For security recommendations, see [Amazon Connect Chat security best practices].
//
// - This API can only be called within one minute of receiving the
// authenticationInitiated event.
//
// - The current supported channel is chat. This API is not supported for Apple
// Messages for Business, WhatsApp, or SMS chats.
//
// ConnectionToken is used for invoking this API instead of ParticipantToken .
//
// The Amazon Connect Participant Service APIs do not use [Signature Version 4 authentication].
//
// [Signature Version 4 authentication]: https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html
// [Amazon Connect Chat security best practices]: https://docs.aws.amazon.com/connect/latest/adminguide/security-best-practices.html#bp-security-chat
