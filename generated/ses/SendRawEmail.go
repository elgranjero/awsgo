package ses

// SendRawEmail is generated as a reference stub.
// Executable command wiring lives under cmd/ses.go.
//
// Composes an email message and immediately queues it for sending.
//
// This operation is more flexible than the SendEmail operation. When you use the
// SendRawEmail operation, you can specify the headers of the message as well as
// its content. This flexibility is useful, for example, when you need to send a
// multipart MIME email (such a message that contains both a text and an HTML
// version). You can also use this operation to send messages that include
// attachments.
//
// The SendRawEmail operation has the following requirements:
//
// - You can only send email from [verified email addresses or domains]. If you try to send email from an address
// that isn't verified, the operation results in an "Email address not verified"
// error.
//
// - If your account is still in the [Amazon SES sandbox], you can only send email to other verified
// addresses in your account, or to addresses that are associated with the [Amazon SES mailbox simulator].
//
// - The maximum message size, including attachments, is 10 MB.
//
// - Each message has to include at least one recipient address. A recipient
// address includes any address on the To:, CC:, or BCC: lines.
//
// - If you send a single message to more than one recipient address, and one of
// the recipient addresses isn't in a valid format (that is, it's not in the format
// UserName(at)[SubDomain.]Domain.TopLevelDomain), Amazon SES rejects the entire
// message, even if the other addresses are valid.
//
// - Each message can include up to 50 recipient addresses across the To:, CC:,
// or BCC: lines. If you need to send a single message to more than 50 recipients,
// you have to split the list of recipient addresses into groups of less than 50
// recipients, and send separate messages to each group.
//
// - Amazon SES allows you to specify 8-bit Content-Transfer-Encoding for MIME
// message parts. However, if Amazon SES has to modify the contents of your message
// (for example, if you use open and click tracking), 8-bit content isn't
// preserved. For this reason, we highly recommend that you encode all content that
// isn't 7-bit ASCII. For more information, see [MIME Encoding]in the Amazon SES Developer
// Guide.
//
// Additionally, keep the following considerations in mind when using the
// SendRawEmail operation:
//
// - Although you can customize the message headers when using the SendRawEmail
// operation, Amazon SES automatically applies its own Message-ID and Date
// headers; if you passed these headers when creating the message, they are
// overwritten by the values that Amazon SES provides.
//
// - If you are using sending authorization to send on behalf of another user,
// SendRawEmail enables you to specify the cross-account identity for the email's
// Source, From, and Return-Path parameters in one of two ways: you can pass
// optional parameters SourceArn , FromArn , and/or ReturnPathArn , or you can
// include the following X-headers in the header of your raw email:
//
// - X-SES-SOURCE-ARN
//
// - X-SES-FROM-ARN
//
// - X-SES-RETURN-PATH-ARN
//
// Don't include these X-headers in the DKIM signature. Amazon SES removes these
//
// before it sends the email.
//
// If you only specify the SourceIdentityArn parameter, Amazon SES sets the From
//
// and Return-Path addresses to the same identity that you specified.
//
// For more information about sending authorization, see the [Using Sending Authorization with Amazon SES]in the Amazon SES
//
// Developer Guide.
//
// - For every message that you send, the total number of recipients (including
// each recipient in the To:, CC: and BCC: fields) is counted against the maximum
// number of emails you can send in a 24-hour period (your sending quota). For more
// information about sending quotas in Amazon SES, see [Managing Your Amazon SES Sending Limits]in the Amazon SES
// Developer Guide.
//
// [Using Sending Authorization with Amazon SES]: https://docs.aws.amazon.com/ses/latest/dg/sending-authorization.html
// [MIME Encoding]: https://docs.aws.amazon.com/ses/latest/dg/send-email-raw.html#send-email-mime-encoding
// [Amazon SES mailbox simulator]: https://docs.aws.amazon.com/ses/latest/dg/send-an-email-from-console.html
// [Amazon SES sandbox]: https://docs.aws.amazon.com/ses/latest/dg/request-production-access.html
// [verified email addresses or domains]: https://docs.aws.amazon.com/ses/latest/dg/verify-addresses-and-domains.html
//
// [Managing Your Amazon SES Sending Limits]: https://docs.aws.amazon.com/ses/latest/dg/manage-sending-quotas.html
