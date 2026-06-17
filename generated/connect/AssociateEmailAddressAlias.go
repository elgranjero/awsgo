package connect

// AssociateEmailAddressAlias is generated as a reference stub.
// Executable command wiring lives under cmd/connect.go.
//
// Associates an email address alias with an existing email address in an Amazon
// Connect instance. This creates a forwarding relationship where emails sent to
// the alias email address are automatically forwarded to the primary email
// address.
//
// # Use cases
//
// Following are common uses cases for this API:
//
// - Unified customer support: Create multiple entry points (for example,
// support(at)example.com, help(at)example.com, customercare(at)example.com) that all
// forward to a single agent queue for streamlined management.
//
// - Department consolidation: Forward emails from legacy department addresses
// (for example, sales(at)example.com, info(at)example.com) to a centralized customer
// service email during organizational restructuring.
//
// - Brand management: Enable you to use familiar brand-specific email addresses
// that forward to the appropriate Amazon Connect instance email address.
//
// Important things to know
//
// - Each email address can have a maximum of one alias. You cannot create
// multiple aliases for the same email address.
//
// - If the alias email address already receives direct emails, it continues to
// receive direct emails plus forwarded emails.
//
// - You cannot chain email aliases together (that is, create an alias of an
// alias).
//
// AssociateEmailAddressAlias does not return the following information:
//
// - A confirmation of the alias relationship details (you must call [DescribeEmailAddress]to verify).
//
// - The timestamp of when the association occurred.
//
// - The status of the forwarding configuration.
//
// Endpoints: See [Amazon Connect endpoints and quotas].
//
// # Related operations
//
// [DisassociateEmailAddressAlias]
// - : Removes the alias association between two email addresses in an Amazon
// Connect instance.
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
// [DisassociateEmailAddressAlias]: https://docs.aws.amazon.com/connect/latest/APIReference/API_DisassociateEmailAddressAlias.html
// [SearchEmailAddresses]: https://docs.aws.amazon.com/connect/latest/APIReference/API_SearchEmailAddresses.html
// [UpdateEmailAddressMetadata]: https://docs.aws.amazon.com/connect/latest/APIReference/API_UpdateEmailAddressMetadata.html
// [Amazon Connect endpoints and quotas]: https://docs.aws.amazon.com/general/latest/gr/connect_region.html
// [CreateEmailAddress]: https://docs.aws.amazon.com/connect/latest/APIReference/API_CreateEmailAddress.html
