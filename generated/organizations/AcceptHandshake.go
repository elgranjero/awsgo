package organizations

// AcceptHandshake is generated as a reference stub.
// Executable command wiring lives under cmd/organizations.go.
//
// Accepts a handshake by sending an ACCEPTED response to the sender. You can view
// accepted handshakes in API responses for 30 days before they are deleted.
//
// Only the management account can accept the following handshakes:
//
// - Enable all features final confirmation ( APPROVE_ALL_FEATURES )
//
// - Billing transfer ( TRANSFER_RESPONSIBILITY )
//
// For more information, see [Enabling all features] and [Responding to a billing transfer invitation] in the Organizations User Guide.
//
// Only a member account can accept the following handshakes:
//
// - Invitation to join ( INVITE )
//
// - Approve all features request ( ENABLE_ALL_FEATURES )
//
// For more information, see [Responding to invitations] and [Enabling all features] in the Organizations User Guide.
//
// [Enabling all features]: https://docs.aws.amazon.com/organizations/latest/userguide/manage-begin-all-features-standard-migration.html#manage-approve-all-features-invite
// [Responding to invitations]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_accounts_accept-decline-invite.html
// [Responding to a billing transfer invitation]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_transfer_billing-respond-invitation.html
