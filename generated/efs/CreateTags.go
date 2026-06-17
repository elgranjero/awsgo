package efs

// CreateTags is generated as a reference stub.
// Executable command wiring lives under cmd/efs.go.
//
// DEPRECATED - CreateTags is deprecated and not maintained. To create tags for
// EFS resources, use the API action.
//
// Creates or overwrites tags associated with a file system. Each tag is a
// key-value pair. If a tag key specified in the request already exists on the file
// system, this operation overwrites its value with the value provided in the
// request. If you add the Name tag to your file system, Amazon EFS returns it in
// the response to the DescribeFileSystemsoperation.
//
// This operation requires permission for the elasticfilesystem:CreateTags action.
//
// Deprecated: Use TagResource.
