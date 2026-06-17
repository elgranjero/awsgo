package lakeformation

// UpdateLFTag is generated as a reference stub.
// Executable command wiring lives under cmd/lakeformation.go.
//
// Updates the list of possible values for the specified LF-tag key. If the LF-tag
// does not exist, the operation throws an EntityNotFoundException. The values in
// the delete key values will be deleted from list of possible values. If any value
// in the delete key values is attached to a resource, then API errors out with a
// 400 Exception - "Update not allowed". Untag the attribute before deleting the
// LF-tag key's value.
