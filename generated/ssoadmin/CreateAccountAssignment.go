package ssoadmin

// CreateAccountAssignment is generated as a reference stub.
// Executable command wiring lives under cmd/ssoadmin.go.
//
// Assigns access to a principal for a specified Amazon Web Services account using
// a specified permission set.
//
// The term principal here refers to a user or group that is defined in IAM
// Identity Center.
//
// As part of a successful CreateAccountAssignment call, the specified permission
// set will automatically be provisioned to the account in the form of an IAM
// policy. That policy is attached to the IAM role created in IAM Identity Center.
// If the permission set is subsequently updated, the corresponding IAM policies
// attached to roles in your accounts will not be updated automatically. In this
// case, you must call ProvisionPermissionSetto make these updates.
//
// After a successful response, call DescribeAccountAssignmentCreationStatus to
// describe the status of an assignment creation request.
