package efs

// DeleteTags is generated as a reference stub.
// Executable command wiring lives under cmd/efs.go.
//
// DEPRECATED - DeleteTags is deprecated and not maintained. To remove tags from
// EFS resources, use the API action.
//
// Deletes the specified tags from a file system. If the DeleteTags request
// includes a tag key that doesn't exist, Amazon EFS ignores it and doesn't cause
// an error. For more information about tags and related restrictions, see [Tag restrictions]in the
// Billing and Cost Management User Guide.
//
// This operation requires permissions for the elasticfilesystem:DeleteTags action.
//
// Deprecated: Use UntagResource.
//
// [Tag restrictions]: https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html
