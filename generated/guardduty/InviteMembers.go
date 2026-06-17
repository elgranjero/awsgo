package guardduty

// InviteMembers is generated as a reference stub.
// Executable command wiring lives under cmd/guardduty.go.
//
// Invites Amazon Web Services accounts to become members of an organization
// administered by the Amazon Web Services account that invokes this API. If you
// are using Amazon Web Services Organizations to manage your GuardDuty
// environment, this step is not needed. For more information, see [Managing accounts with organizations].
//
// To invite Amazon Web Services accounts, the first step is to ensure that
// GuardDuty has been enabled in the potential member accounts. You can now invoke
// this API to add accounts by invitation. The invited accounts can either accept
// or decline the invitation from their GuardDuty accounts. Each invited Amazon Web
// Services account can choose to accept the invitation from only one Amazon Web
// Services account. For more information, see [Managing GuardDuty accounts by invitation].
//
// After the invite has been accepted and you choose to disassociate a member
// account (by using [DisassociateMembers]) from your account, the details of the member account
// obtained by invoking [CreateMembers], including the associated email addresses, will be
// retained. This is done so that you can invoke InviteMembers without the need to
// invoke [CreateMembers]again. To remove the details associated with a member account, you must
// also invoke [DeleteMembers].
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
// [Managing GuardDuty accounts by invitation]: https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_invitations.html
// [DeleteMembers]: https://docs.aws.amazon.com/guardduty/latest/APIReference/API_DeleteMembers.html
// [CreateMembers]: https://docs.aws.amazon.com/guardduty/latest/APIReference/API_CreateMembers.html
// [Managing accounts with organizations]: https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_organizations.html
// [DisassociateMembers]: https://docs.aws.amazon.com/guardduty/latest/APIReference/API_DisassociateMembers.html
// [InviteMembers]: https://docs.aws.amazon.com/guardduty/latest/APIReference/API_InviteMembers.html
