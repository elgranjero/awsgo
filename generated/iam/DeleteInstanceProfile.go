package iam

// DeleteInstanceProfile is generated as a reference stub.
// Executable command wiring lives under cmd/iam.go.
//
// Deletes the specified instance profile. The instance profile must not have an
// associated role.
//
// Make sure that you do not have any Amazon EC2 instances running with the
// instance profile you are about to delete. Deleting a role or instance profile
// that is associated with a running instance will break any applications running
// on the instance.
//
// For more information about instance profiles, see [Using instance profiles] in the IAM User Guide.
//
// [Using instance profiles]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use_switch-role-ec2_instance-profiles.html
