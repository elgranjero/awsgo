package acm

// ResendValidationEmail is generated as a reference stub.
// Executable command wiring lives under cmd/acm.go.
//
// Resends the email that requests domain ownership validation. The domain owner
// or an authorized representative must approve the ACM certificate before it can
// be issued. The certificate can be approved by clicking a link in the mail to
// navigate to the Amazon certificate approval website and then clicking I Approve.
// However, the validation email can be blocked by spam filters. Therefore, if you
// do not receive the original mail, you can request that the mail be resent within
// 72 hours of requesting the ACM certificate. If more than 72 hours have elapsed
// since your original request or since your last attempt to resend validation
// mail, you must request a new certificate. For more information about setting up
// your contact email addresses, see [Configure Email for your Domain].
//
// [Configure Email for your Domain]: https://docs.aws.amazon.com/acm/latest/userguide/setup-email.html
