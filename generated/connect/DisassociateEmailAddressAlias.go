package connect

// DisassociateEmailAddressAlias is generated as a reference stub.
// Executable command wiring lives under cmd/connect.go.
//
// Removes the alias association between two email addresses in an Amazon Connect
// instance. After disassociation, emails sent to the former alias email address
// are no longer forwarded to the primary email address. Both email addresses
// continue to exist independently and can receive emails directly.
//
// # Use cases
//
// Following are common uses cases for this API:
//
// - Department separation: Remove alias relationships when splitting a
// consolidated support queue back into separate department-specific queues.
//
// - Email address retirement: Cleanly remove forwarding relationships before
// decommissioning old email addresses.
//
// - Organizational restructuring: Reconfigure email routing when business
// processes change and aliases are no longer needed.
//
// Important things to know
//
// - Concurrent operations: This API uses distributed locking, so concurrent
// operations on the same email addresses may be temporarily blocked.
//
// - Emails sent to the former alias address are still delivered directly to
// that address if it exists.
//
// - You do not need to delete the email addresses after disassociation. Both
// addresses remain active independently.
//
// - After a successful disassociation, you can immediately create a new alias
// relationship with the same addresses.
//
// - 200 status means alias was successfully disassociated.
//
// DisassociateEmailAddressAlias does not return the following information:
//
// - Details in the response about the email that was disassociated. The
// response returns an empty body.
//
// - The timestamp of when the disassociation occurred.
//
// Endpoints: See [Amazon Connect endpoints and quotas].
//
// # Related operations
//
// [AssociateEmailAddressAlias]
// - : Associates an email address alias with an existing email address in an
// Amazon Connect instance.
//
// [DescribeEmailAddress]
// - : View current alias configurations for an email address.
//
// [SearchEmailAddresses]
// - : Find email addresses and their alias relationships across an instance.
//
// [CreateEmailAddress]
// - : Create new email addresses that can participate in alias relationships.
//
// [DeleteEmailAddress]
// - : Remove email addresses (automatically removes any alias relationships).
//
// [UpdateEmailAddressMetadata]
// - : Modify email address properties (does not affect alias relationships).
//
// [DescribeEmailAddress]: https://docs.aws.amazon.com/connect/latest/APIReference/API_DescribeEmailAddress.html
// [DeleteEmailAddress]: https://docs.aws.amazon.com/connect/latest/APIReference/API_DeleteEmailAddress.html
// [SearchEmailAddresses]: https://docs.aws.amazon.com/connect/latest/APIReference/API_SearchEmailAddresses.html
// [UpdateEmailAddressMetadata]: https://docs.aws.amazon.com/connect/latest/APIReference/API_UpdateEmailAddressMetadata.html
// [Amazon Connect endpoints and quotas]: https://docs.aws.amazon.com/general/latest/gr/connect_region.html
// [CreateEmailAddress]: https://docs.aws.amazon.com/connect/latest/APIReference/API_CreateEmailAddress.html
// [AssociateEmailAddressAlias]: https://docs.aws.amazon.com/connect/latest/APIReference/API_AssociateEmailAddressAlias.html
