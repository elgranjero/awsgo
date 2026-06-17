package detective

// CreateMembers is generated as a reference stub.
// Executable command wiring lives under cmd/detective.go.
//
// CreateMembers is used to send invitations to accounts. For the organization
// behavior graph, the Detective administrator account uses CreateMembers to
// enable organization accounts as member accounts.
//
// For invited accounts, CreateMembers sends a request to invite the specified
// Amazon Web Services accounts to be member accounts in the behavior graph. This
// operation can only be called by the administrator account for a behavior graph.
//
// CreateMembers verifies the accounts and then invites the verified accounts. The
// administrator can optionally specify to not send invitation emails to the member
// accounts. This would be used when the administrator manages their member
// accounts centrally.
//
// For organization accounts in the organization behavior graph, CreateMembers
// attempts to enable the accounts. The organization accounts do not receive
// invitations.
//
// The request provides the behavior graph ARN and the list of accounts to invite
// or to enable.
//
// The response separates the requested accounts into two lists:
//
// - The accounts that CreateMembers was able to process. For invited accounts,
// includes member accounts that are being verified, that have passed verification
// and are to be invited, and that have failed verification. For organization
// accounts in the organization behavior graph, includes accounts that can be
// enabled and that cannot be enabled.
//
// - The accounts that CreateMembers was unable to process. This list includes
// accounts that were already invited to be member accounts in the behavior graph.
