package configservice

// DescribeRemediationExceptions is generated as a reference stub.
// Executable command wiring lives under cmd/configservice.go.
//
// Returns the details of one or more remediation exceptions. A detailed view of a
// remediation exception for a set of resources that includes an explanation of an
// exception and the time when the exception will be deleted. When you specify the
// limit and the next token, you receive a paginated response.
//
// Config generates a remediation exception when a problem occurs executing a
// remediation action to a specific resource. Remediation exceptions blocks
// auto-remediation until the exception is cleared.
//
// When you specify the limit and the next token, you receive a paginated
// response.
//
// Limit and next token are not applicable if you request resources in batch. It
// is only applicable, when you request all resources.
