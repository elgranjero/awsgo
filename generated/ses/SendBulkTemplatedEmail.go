package ses

// SendBulkTemplatedEmail is generated as a reference stub.
// Executable command wiring lives under cmd/ses.go.
//
// Composes an email message to multiple destinations. The message body is created
// using an email template.
//
// To send email using this operation, your call must meet the following
// requirements:
//
// - The call must refer to an existing email template. You can create email
// templates using CreateTemplate.
//
// - The message must be sent from a verified email address or domain.
//
// - If your account is still in the Amazon SES sandbox, you may send only to
// verified addresses or domains, or to email addresses associated with the Amazon
// SES Mailbox Simulator. For more information, see [Verifying Email Addresses and Domains]in the Amazon SES Developer
// Guide.
//
// - The maximum message size is 10 MB.
//
// - Each Destination parameter must include at least one recipient email
// address. The recipient address can be a To: address, a CC: address, or a BCC:
// address. If a recipient email address is invalid (that is, it is not in the
// format UserName(at)[SubDomain.]Domain.TopLevelDomain), the entire message is
// rejected, even if the message contains other recipients that are valid.
//
// - The message may not include more than 50 recipients, across the To:, CC:
// and BCC: fields. If you need to send an email message to a larger audience, you
// can divide your recipient list into groups of 50 or fewer, and then call the
// SendBulkTemplatedEmail operation several times to send the message to each
// group.
//
// - The number of destinations you can contact in a single call can be limited
// by your account's maximum sending rate.
//
// [Verifying Email Addresses and Domains]: https://docs.aws.amazon.com/ses/latest/dg/verify-addresses-and-domains.html
