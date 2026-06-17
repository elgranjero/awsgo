package workmail

// DeregisterMailDomain is generated as a reference stub.
// Executable command wiring lives under cmd/workmail.go.
//
// Removes a domain from WorkMail, stops email routing to WorkMail, and removes
// the authorization allowing WorkMail use. SES keeps the domain because other
// applications may use it. You must first remove any email address used by
// WorkMail entities before you remove the domain.
