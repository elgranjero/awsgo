package iam

// TagInstanceProfile is generated as a reference stub.
// Executable command wiring lives under cmd/iam.go.
//
// Adds one or more tags to an IAM instance profile. If a tag with the same key
// name already exists, then that tag is overwritten with the new value.
//
// Each tag consists of a key name and an associated value. By assigning tags to
// your resources, you can do the following:
//
// - Administrative grouping and discovery - Attach tags to resources to aid in
// organization and search. For example, you could search for all resources with
// the key name Project and the value MyImportantProject. Or search for all
// resources with the key name Cost Center and the value 41200.
//
// - Access control - Include tags in IAM user-based and resource-based
// policies. You can use tags to restrict access to only an IAM instance profile
// that has a specified tag attached. For examples of policies that show how to use
// tags to control access, see [Control access using IAM tags]in the IAM User Guide.
//
// - If any one of the tags is invalid or if you exceed the allowed maximum
// number of tags, then the entire request fails and the resource is not created.
// For more information about tagging, see [Tagging IAM resources]in the IAM User Guide.
//
// - Amazon Web Services always interprets the tag Value as a single string. If
// you need to store an array, you can store comma-separated values in the string.
// However, you must interpret the value in your code.
//
// [Control access using IAM tags]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_tags.html
// [Tagging IAM resources]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html
