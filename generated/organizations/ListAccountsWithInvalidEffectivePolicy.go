package organizations

// ListAccountsWithInvalidEffectivePolicy is generated as a reference stub.
// Executable command wiring lives under cmd/organizations.go.
//
// Lists all the accounts in an organization that have invalid effective policies.
// An invalid effective policy is an [effective policy]that fails validation checks, resulting in
// the effective policy not being fully enforced on all the intended accounts
// within an organization.
//
// You can only call this operation from the management account or a member
// account that is a delegated administrator.
//
// [effective policy]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_effective.html
