package lakeformation

// ListPermissions is generated as a reference stub.
// Executable command wiring lives under cmd/lakeformation.go.
//
// Returns a list of the principal permissions on the resource, filtered by the
// permissions of the caller. For example, if you are granted an ALTER permission,
// you are able to see only the principal permissions for ALTER.
//
// This operation returns only those permissions that have been explicitly
// granted. If both Principal and Resource parameters are provided, the response
// returns effective permissions rather than the explicitly granted permissions.
//
// For information about permissions, see [Security and Access Control to Metadata and Data].
//
// [Security and Access Control to Metadata and Data]: https://docs.aws.amazon.com/lake-formation/latest/dg/security-data-access.html
