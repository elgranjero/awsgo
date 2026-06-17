package ses

// SendEmail is generated as a reference stub.
// Executable command wiring lives under cmd/ses.go.
//
// Composes an email message and immediately queues it for sending. To send email
// using this operation, your message must meet the following requirements:
//
// - The message must be sent from a verified email address or domain. If you
// attempt to send email using a non-verified address or domain, the operation
// results in an "Email address not verified" error.
//
// - If your account is still in the Amazon SES sandbox, you may only send to
// verified addresses or domains, or to email addresses associated with the Amazon
// SES Mailbox Simulator. For more information, see [Verifying Email Addresses and Domains]in the Amazon SES Developer
// Guide.
//
// - The maximum message size is 10 MB.
//
// - The message must include at least one recipient email address. The
// recipient address can be a To: address, a CC: address, or a BCC: address. If a
// recipient email address is invalid (that is, it is not in the format
// UserName(at)[SubDomain.]Domain.TopLevelDomain), the entire message is rejected,
// even if the message contains other recipients that are valid.
//
// - The message may not include more than 50 recipients, across the To:, CC:
// and BCC: fields. If you need to send an email message to a larger audience, you
// can divide your recipient list into groups of 50 or fewer, and then call the
// SendEmail operation several times to send the message to each group.
//
// For every message that you send, the total number of recipients (including each
// recipient in the To:, CC: and BCC: fields) is counted against the maximum number
// of emails you can send in a 24-hour period (your sending quota). For more
// information about sending quotas in Amazon SES, see [Managing Your Amazon SES Sending Limits]in the Amazon SES Developer
// Guide.
//
// [Verifying Email Addresses and Domains]: https://docs.aws.amazon.com/ses/latest/dg/verify-addresses-and-domains.html
// [Managing Your Amazon SES Sending Limits]: https://docs.aws.amazon.com/ses/latest/dg/manage-sending-quotas.html
