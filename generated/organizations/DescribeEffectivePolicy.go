package organizations

// DescribeEffectivePolicy is generated as a reference stub.
// Executable command wiring lives under cmd/organizations.go.
//
// Returns the contents of the effective policy for specified policy type and
// account. The effective policy is the aggregation of any policies of the
// specified type that the account inherits, plus any policy of that type that is
// directly attached to the account.
//
// This operation applies only to management policies. It does not apply to
// authorization policies: service control policies (SCPs) and resource control
// policies (RCPs).
//
// For more information about policy inheritance, see [Understanding management policy inheritance] in the Organizations User
// Guide.
//
// You can call this operation from any account in a organization.
//
// [Understanding management policy inheritance]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_inheritance_mgmt.html
