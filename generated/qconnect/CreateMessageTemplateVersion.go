package qconnect

// CreateMessageTemplateVersion is generated as a reference stub.
// Executable command wiring lives under cmd/qconnect.go.
//
// Creates a new Amazon Q in Connect message template version from the current
// content and configuration of a message template. Versions are immutable and
// monotonically increasing. Once a version is created, you can reference a
// specific version of the message template by passing in
// <message-template-id>:<versionNumber> as the message template identifier. An
// error is displayed if the supplied messageTemplateContentSha256 is different
// from the messageTemplateContentSha256 of the message template with $LATEST
// qualifier. If multiple CreateMessageTemplateVersion requests are made while the
// message template remains the same, only the first invocation creates a new
// version and the succeeding requests will return the same response as the first
// invocation.
