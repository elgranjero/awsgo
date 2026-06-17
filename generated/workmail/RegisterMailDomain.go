package workmail

// RegisterMailDomain is generated as a reference stub.
// Executable command wiring lives under cmd/workmail.go.
//
// Registers a new domain in WorkMail and SES, and configures it for use by
// WorkMail. Emails received by SES for this domain are routed to the specified
// WorkMail organization, and WorkMail has permanent permission to use the
// specified domain for sending your users' emails.
