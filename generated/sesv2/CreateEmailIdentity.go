package sesv2

// CreateEmailIdentity is generated as a reference stub.
// Executable command wiring lives under cmd/sesv2.go.
//
// Starts the process of verifying an email identity. An identity is an email
// address or domain that you use when you send email. Before you can use an
// identity to send email, you first have to verify it. By verifying an identity,
// you demonstrate that you're the owner of the identity, and that you've given
// Amazon SES API v2 permission to send email from the identity.
//
// When you verify an email address, Amazon SES sends an email to the address.
// Your email address is verified as soon as you follow the link in the
// verification email.
//
// When you verify a domain without specifying the DkimSigningAttributes object,
// this operation provides a set of DKIM tokens. You can convert these tokens into
// CNAME records, which you then add to the DNS configuration for your domain. Your
// domain is verified when Amazon SES detects these records in the DNS
// configuration for your domain. This verification method is known as [Easy DKIM].
//
// Alternatively, you can perform the verification process by providing your own
// public-private key pair. This verification method is known as Bring Your Own
// DKIM (BYODKIM). To use BYODKIM, your call to the CreateEmailIdentity operation
// has to include the DkimSigningAttributes object. When you specify this object,
// you provide a selector (a component of the DNS record name that identifies the
// public key to use for DKIM authentication) and a private key.
//
// When you verify a domain, this operation provides a set of DKIM tokens, which
// you can convert into CNAME tokens. You add these CNAME tokens to the DNS
// configuration for your domain. Your domain is verified when Amazon SES detects
// these records in the DNS configuration for your domain. For some DNS providers,
// it can take 72 hours or more to complete the domain verification process.
//
// Additionally, you can associate an existing configuration set with the email
// identity that you're verifying.
//
// [Easy DKIM]: https://docs.aws.amazon.com/ses/latest/DeveloperGuide/easy-dkim.html
