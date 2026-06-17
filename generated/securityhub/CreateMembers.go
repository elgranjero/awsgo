package securityhub

// CreateMembers is generated as a reference stub.
// Executable command wiring lives under cmd/securityhub.go.
//
// Creates a member association in Security Hub CSPM between the specified
// accounts and the account used to make the request, which is the administrator
// account. If you are integrated with Organizations, then the administrator
// account is designated by the organization management account.
//
// CreateMembers is always used to add accounts that are not organization members.
//
// For accounts that are managed using Organizations, CreateMembers is only used
// in the following cases:
//
// - Security Hub CSPM is not configured to automatically add new organization
// accounts.
//
// - The account was disassociated or deleted in Security Hub CSPM.
//
// This action can only be used by an account that has Security Hub CSPM enabled.
// To enable Security Hub CSPM, you can use the EnableSecurityHub operation.
//
// For accounts that are not organization members, you create the account
// association and then send an invitation to the member account. To send the
// invitation, you use the InviteMembers operation. If the account owner accepts
// the invitation, the account becomes a member account in Security Hub CSPM.
//
// Accounts that are managed using Organizations don't receive an invitation. They
// automatically become a member account in Security Hub CSPM.
//
// - If the organization account does not have Security Hub CSPM enabled, then
// Security Hub CSPM and the default standards are automatically enabled. Note that
// Security Hub CSPM cannot be enabled automatically for the organization
// management account. The organization management account must enable Security Hub
// CSPM before the administrator account enables it as a member account.
//
// - For organization accounts that already have Security Hub CSPM enabled,
// Security Hub CSPM does not make any other changes to those accounts. It does not
// change their enabled standards or controls.
//
// A permissions policy is added that permits the administrator account to view
// the findings generated in the member account.
//
// To remove the association between the administrator and member accounts, use
// the DisassociateFromMasterAccount or DisassociateMembers operation.
