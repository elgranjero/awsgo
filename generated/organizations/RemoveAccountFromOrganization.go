package organizations

// RemoveAccountFromOrganization is generated as a reference stub.
// Executable command wiring lives under cmd/organizations.go.
//
// Removes the specified account from the organization.
//
// The removed account becomes a standalone account that isn't a member of any
// organization. It's no longer subject to any policies and is responsible for its
// own bill payments. The organization's management account is no longer charged
// for any expenses accrued by the member account after it's removed from the
// organization.
//
// You can only call this operation from the management account. Member accounts
// can remove themselves with LeaveOrganizationinstead.
//
// - You can remove an account from your organization only if the account is
// configured with the information required to operate as a standalone account.
// When you create an account in an organization using the Organizations console,
// API, or CLI commands, the information required of standalone accounts is not
// automatically collected. For more information, see [Considerations before removing an account from an organization]in the Organizations User
// Guide.
//
// - The account that you want to leave must not be a delegated administrator
// account for any Amazon Web Services service enabled for your organization. If
// the account is a delegated administrator, you must first change the delegated
// administrator account to another account that is remaining in the organization.
//
// - After the account leaves the organization, all tags that were attached to
// the account object in the organization are deleted. Amazon Web Services accounts
// outside of an organization do not support tags.
//
// [Considerations before removing an account from an organization]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_account-before-remove.html
