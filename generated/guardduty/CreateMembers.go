package guardduty

// CreateMembers is generated as a reference stub.
// Executable command wiring lives under cmd/guardduty.go.
//
// Creates member accounts of the current Amazon Web Services account by
// specifying a list of Amazon Web Services account IDs. This step is a
// prerequisite for managing the associated member accounts either by invitation or
// through an organization.
//
// As a delegated administrator, using CreateMembers will enable GuardDuty in the
// added member accounts, with the exception of the organization delegated
// administrator account. A delegated administrator must enable GuardDuty prior to
// being added as a member.
//
// When you use CreateMembers as an Organizations delegated administrator,
// GuardDuty applies your organization's auto-enable settings to the member
// accounts in this request, irrespective of the accounts being new or existing
// members. For more information about the existing auto-enable settings for your
// organization, see [DescribeOrganizationConfiguration].
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
// [DescribeOrganizationConfiguration]: https://docs.aws.amazon.com/guardduty/latest/APIReference/API_DescribeOrganizationConfiguration.html
// [InviteMembers]: https://docs.aws.amazon.com/guardduty/latest/APIReference/API_InviteMembers.html
