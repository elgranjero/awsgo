package iam

// RemoveRoleFromInstanceProfile is generated as a reference stub.
// Executable command wiring lives under cmd/iam.go.
//
// Removes the specified IAM role from the specified Amazon EC2 instance profile.
//
// Make sure that you do not have any Amazon EC2 instances running with the role
// you are about to remove from the instance profile. Removing a role from an
// instance profile that is associated with a running instance might break any
// applications running on the instance.
//
// For more information about roles, see [IAM roles] in the IAM User Guide. For more
// information about instance profiles, see [Using instance profiles]in the IAM User Guide.
//
// [Using instance profiles]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use_switch-role-ec2_instance-profiles.html
// [IAM roles]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles.html
