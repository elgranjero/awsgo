package ses

// SendTemplatedEmail is generated as a reference stub.
// Executable command wiring lives under cmd/ses.go.
//
// Composes an email message using an email template and immediately queues it for
// sending.
//
// To send email using this operation, your call must meet the following
// requirements:
//
// - The call must refer to an existing email template. You can create email
// templates using the CreateTemplateoperation.
//
// - The message must be sent from a verified email address or domain.
//
// - If your account is still in the Amazon SES sandbox, you may only send to
// verified addresses or domains, or to email addresses associated with the Amazon
// SES Mailbox Simulator. For more information, see [Verifying Email Addresses and Domains]in the Amazon SES Developer
// Guide.
//
// - The maximum message size is 10 MB.
//
// - Calls to the SendTemplatedEmail operation may only include one Destination
// parameter. A destination is a set of recipients that receives the same version
// of the email. The Destination parameter can include up to 50 recipients,
// across the To:, CC: and BCC: fields.
//
// - The Destination parameter must include at least one recipient email address.
// The recipient address can be a To: address, a CC: address, or a BCC: address. If
// a recipient email address is invalid (that is, it is not in the format
// UserName(at)[SubDomain.]Domain.TopLevelDomain), the entire message is rejected,
// even if the message contains other recipients that are valid.
//
// If your call to the SendTemplatedEmail operation includes all of the required
// parameters, Amazon SES accepts it and returns a Message ID. However, if Amazon
// SES can't render the email because the template contains errors, it doesn't send
// the email. Additionally, because it already accepted the message, Amazon SES
// doesn't return a message stating that it was unable to send the email.
//
// For these reasons, we highly recommend that you set up Amazon SES to send you
// notifications when Rendering Failure events occur. For more information, see [Sending Personalized Email Using the Amazon SES API]in
// the Amazon Simple Email Service Developer Guide.
//
// [Sending Personalized Email Using the Amazon SES API]: https://docs.aws.amazon.com/ses/latest/dg/send-personalized-email-api.html
// [Verifying Email Addresses and Domains]: https://docs.aws.amazon.com/ses/latest/dg/verify-addresses-and-domains.html
