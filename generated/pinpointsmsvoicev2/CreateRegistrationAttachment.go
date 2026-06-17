package pinpointsmsvoicev2

// CreateRegistrationAttachment is generated as a reference stub.
// Executable command wiring lives under cmd/pinpointsmsvoicev2.go.
//
// Create a new registration attachment to use for uploading a file or a URL to a
// file. The maximum file size is 500KB and valid file extensions are PDF, JPEG and
// PNG. For example, many sender ID registrations require a signed “letter of
// authorization” (LOA) to be submitted.
//
// Use either AttachmentUrl or AttachmentBody to upload your attachment. If both
// are specified then an exception is returned.
