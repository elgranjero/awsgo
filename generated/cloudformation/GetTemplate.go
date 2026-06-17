package cloudformation

// GetTemplate is generated as a reference stub.
// Executable command wiring lives under cmd/cloudformation.go.
//
// Returns the template body for a specified stack. You can get the template for
// running or deleted stacks.
//
// For deleted stacks, GetTemplate returns the template for up to 90 days after
// the stack has been deleted.
//
// If the template doesn't exist, a ValidationError is returned.
