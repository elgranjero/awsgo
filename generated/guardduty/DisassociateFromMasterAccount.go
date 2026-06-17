package guardduty

// DisassociateFromMasterAccount is generated as a reference stub.
// Executable command wiring lives under cmd/guardduty.go.
//
// Disassociates the current GuardDuty member account from its administrator
// account.
//
// When you disassociate an invited member from a GuardDuty delegated
// administrator, the member account details obtained from the [CreateMembers]API, including the
// associated email addresses, are retained. This is done so that the delegated
// administrator can invoke the [InviteMembers]API without the need to invoke the CreateMembers
// API again. To remove the details associated with a member account, the delegated
// administrator must invoke the [DeleteMembers]API.
//
// Deprecated: This operation is deprecated, use
// DisassociateFromAdministratorAccount instead
//
// [DeleteMembers]: https://docs.aws.amazon.com/guardduty/latest/APIReference/API_DeleteMembers.html
// [CreateMembers]: https://docs.aws.amazon.com/guardduty/latest/APIReference/API_CreateMembers.html
// [InviteMembers]: https://docs.aws.amazon.com/guardduty/latest/APIReference/API_InviteMembers.html
