package organizations

// InviteAccountToOrganization is generated as a reference stub.
// Executable command wiring lives under cmd/organizations.go.
//
// Sends an invitation to another account to join your organization as a member
// account. Organizations sends email on your behalf to the email address that is
// associated with the other account's owner. The invitation is implemented as a Handshake
// whose details are in the response.
//
// If you receive an exception that indicates that you exceeded your account
// limits for the organization or that the operation failed because your
// organization is still initializing, wait one hour and then try again. If the
// error persists after an hour, contact [Amazon Web Services Support].
//
// If the request includes tags, then the requester must have the
// organizations:TagResource permission.
//
// You can only call this operation from the management account.
//
// [Amazon Web Services Support]: https://console.aws.amazon.com/support/home#/
