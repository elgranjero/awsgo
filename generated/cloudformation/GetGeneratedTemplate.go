package cloudformation

// GetGeneratedTemplate is generated as a reference stub.
// Executable command wiring lives under cmd/cloudformation.go.
//
// Retrieves a generated template. If the template is in an InProgress or Pending
// status then the template returned will be the template when the template was
// last in a Complete status. If the template has not yet been in a Complete
// status then an empty template will be returned.
