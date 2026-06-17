package iam

// ListInstanceProfiles is generated as a reference stub.
// Executable command wiring lives under cmd/iam.go.
//
// Lists the instance profiles that have the specified path prefix. If there are
// none, the operation returns an empty list. For more information about instance
// profiles, see [Using instance profiles]in the IAM User Guide.
//
// IAM resource-listing operations return a subset of the available attributes for
// the resource. For example, this operation does not return tags, even though they
// are an attribute of the returned object. To view all of the information for an
// instance profile, see [GetInstanceProfile].
//
// You can paginate the results using the MaxItems and Marker parameters.
//
// [Using instance profiles]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use_switch-role-ec2_instance-profiles.html
// [GetInstanceProfile]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetInstanceProfile.html
