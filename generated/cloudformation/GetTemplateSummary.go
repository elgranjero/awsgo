package cloudformation

// GetTemplateSummary is generated as a reference stub.
// Executable command wiring lives under cmd/cloudformation.go.
//
// Returns information about a new or existing template. The GetTemplateSummary
// action is useful for viewing parameter information, such as default parameter
// values and parameter types, before you create or update a stack or StackSet.
//
// You can use the GetTemplateSummary action when you submit a template, or you
// can get template information for a StackSet, or a running or deleted stack.
//
// For deleted stacks, GetTemplateSummary returns the template information for up
// to 90 days after the stack has been deleted. If the template doesn't exist, a
// ValidationError is returned.
