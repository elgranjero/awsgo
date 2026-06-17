package guardduty

// DisassociateMembers is generated as a reference stub.
// Executable command wiring lives under cmd/guardduty.go.
//
// Disassociates GuardDuty member accounts (from the current administrator
// account) specified by the account IDs.
//
// When you disassociate an invited member from a GuardDuty delegated
// administrator, the member account details obtained from the [CreateMembers]API, including the
// associated email addresses, are retained. This is done so that the delegated
// administrator can invoke the [InviteMembers]API without the need to invoke the CreateMembers
// API again. To remove the details associated with a member account, the delegated
// administrator must invoke the [DeleteMembers]API.
//
// With autoEnableOrganizationMembers configuration for your organization set to
// ALL , you'll receive an error if you attempt to disassociate a member account
// before removing them from your organization.
//
// If you disassociate a member account that was added by invitation, the member
// account details obtained from this API, including the associated email
// addresses, will be retained. This is done so that the delegated administrator
// can invoke the [InviteMembers]API without the need to invoke the CreateMembers API again. To
// remove the details associated with a member account, the delegated administrator
// must invoke the [DeleteMembers]API.
//
// When the member accounts added through Organizations are later disassociated,
// you (administrator) can't invite them by calling the InviteMembers API. You can
// create an association with these member accounts again only by calling the
// CreateMembers API.
//
// [DeleteMembers]: https://docs.aws.amazon.com/guardduty/latest/APIReference/API_DeleteMembers.html
// [CreateMembers]: https://docs.aws.amazon.com/guardduty/latest/APIReference/API_CreateMembers.html
// [InviteMembers]: https://docs.aws.amazon.com/guardduty/latest/APIReference/API_InviteMembers.html
