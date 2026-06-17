package socialmessaging

// PostWhatsAppMessageMedia is generated as a reference stub.
// Executable command wiring lives under cmd/socialmessaging.go.
//
// Upload a media file to the WhatsApp service. Only the specified
// originationPhoneNumberId has the permissions to send the media file when using [SendWhatsAppMessage]
// . You must use either sourceS3File or sourceS3PresignedUrl for the source. If
// both or neither are specified then an InvalidParameterException is returned.
//
// [SendWhatsAppMessage]: https://docs.aws.amazon.com/social-messaging/latest/APIReference/API_SendWhatsAppMessage.html
