package fms

// PutAdminAccount is generated as a reference stub.
// Executable command wiring lives under cmd/fms.go.
//
// Creates or updates an Firewall Manager administrator account. The account must
// be a member of the organization that was onboarded to Firewall Manager by AssociateAdminAccount.
// Only the organization's management account can create an Firewall Manager
// administrator account. When you create an Firewall Manager administrator
// account, the service checks to see if the account is already a delegated
// administrator within Organizations. If the account isn't a delegated
// administrator, Firewall Manager calls Organizations to delegate the account
// within Organizations. For more information about administrator accounts within
// Organizations, see [Managing the Amazon Web Services Accounts in Your Organization].
//
// [Managing the Amazon Web Services Accounts in Your Organization]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_accounts.html
