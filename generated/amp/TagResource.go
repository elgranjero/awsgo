package amp

// TagResource is generated as a reference stub.
// Executable command wiring lives under cmd/amp.go.
//
// The TagResource operation associates tags with an Amazon Managed Service for
// Prometheus resource. The only resources that can be tagged are rule groups
// namespaces, scrapers, and workspaces.
//
// If you specify a new tag key for the resource, this tag is appended to the list
// of tags associated with the resource. If you specify a tag key that is already
// associated with the resource, the new tag value that you specify replaces the
// previous value for that tag. To remove a tag, use UntagResource .
