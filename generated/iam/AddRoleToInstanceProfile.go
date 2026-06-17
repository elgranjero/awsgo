package iam

// AddRoleToInstanceProfile is generated as a reference stub.
// Executable command wiring lives under cmd/iam.go.
//
// Adds the specified IAM role to the specified instance profile. An instance
// profile can contain only one role, and this quota cannot be increased. You can
// remove the existing role and then add a different role to an instance profile.
// You must then wait for the change to appear across all of Amazon Web Services
// because of [eventual consistency]. To force the change, you must [disassociate the instance profile] and then [associate the instance profile], or you can stop your
// instance and then restart it.
//
// The caller of this operation must be granted the PassRole permission on the IAM
// role by a permissions policy.
//
// When using the [iam:AssociatedResourceArn] condition in a policy to restrict the [PassRole] IAM action, special
// considerations apply if the policy is intended to define access for the
// AddRoleToInstanceProfile action. In this case, you cannot specify a Region or
// instance ID in the EC2 instance ARN. The ARN value must be
// arn:aws:ec2:*:CallerAccountId:instance/* . Using any other ARN value may lead to
// unexpected evaluation results.
//
// For more information about roles, see [IAM roles] in the IAM User Guide. For more
// information about instance profiles, see [Using instance profiles]in the IAM User Guide.
//
// [disassociate the instance profile]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DisassociateIamInstanceProfile.html
// [associate the instance profile]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_AssociateIamInstanceProfile.html
// [Using instance profiles]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use_switch-role-ec2_instance-profiles.html
// [PassRole]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use_passrole.html
// [iam:AssociatedResourceArn]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_iam-condition-keys.html#available-keys-for-iam
// [eventual consistency]: https://en.wikipedia.org/wiki/Eventual_consistency
// [IAM roles]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles.html
