package securityhub

// AcceptInvitation is generated as a reference stub.
// Executable command wiring lives under cmd/securityhub.go.
//
// This method is deprecated. Instead, use AcceptAdministratorInvitation .
//
// The Security Hub CSPM console continues to use AcceptInvitation . It will
// eventually change to use AcceptAdministratorInvitation . Any IAM policies that
// specifically control access to this function must continue to use
// AcceptInvitation . You should also add AcceptAdministratorInvitation to your
// policies to ensure that the correct permissions are in place after the console
// begins to use AcceptAdministratorInvitation .
//
// Accepts the invitation to be a member account and be monitored by the Security
// Hub CSPM administrator account that the invitation was sent from.
//
// This operation is only used by member accounts that are not added through
// Organizations.
//
// When the member account accepts the invitation, permission is granted to the
// administrator account to view findings generated in the member account.
//
// Deprecated: This API has been deprecated, use AcceptAdministratorInvitation API
// instead.
