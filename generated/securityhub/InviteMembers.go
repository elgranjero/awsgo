package securityhub

// InviteMembers is generated as a reference stub.
// Executable command wiring lives under cmd/securityhub.go.
//
// We recommend using Organizations instead of Security Hub CSPM invitations to
// manage your member accounts. For information, see [Managing Security Hub CSPM administrator and member accounts with Organizations]in the Security Hub CSPM User
// Guide.
//
// Invites other Amazon Web Services accounts to become member accounts for the
// Security Hub CSPM administrator account that the invitation is sent from.
//
// This operation is only used to invite accounts that don't belong to an Amazon
// Web Services organization. Organization accounts don't receive invitations.
//
// Before you can use this action to invite a member, you must first use the
// CreateMembers action to create the member account in Security Hub CSPM.
//
// When the account owner enables Security Hub CSPM and accepts the invitation to
// become a member account, the administrator account can view the findings
// generated in the member account.
//
// [Managing Security Hub CSPM administrator and member accounts with Organizations]: https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-accounts-orgs.html
