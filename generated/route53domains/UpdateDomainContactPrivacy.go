package route53domains

// UpdateDomainContactPrivacy is generated as a reference stub.
// Executable command wiring lives under cmd/route53domains.go.
//
// This operation updates the specified domain contact's privacy setting. When
// privacy protection is enabled, your contact information is replaced with contact
// information for the registrar or with the phrase "REDACTED FOR PRIVACY", or "On
// behalf of owner."
//
// While some domains may allow different privacy settings per contact, we
// recommend specifying the same privacy setting for all contacts.
//
// This operation affects only the contact information for the specified contact
// type (administrative, registrant, or technical). If the request succeeds, Amazon
// Route 53 returns an operation ID that you can use with [GetOperationDetail]to track the progress
// and completion of the action. If the request doesn't complete successfully, the
// domain registrant will be notified by email.
//
// By disabling the privacy service via API, you consent to the publication of the
// contact information provided for this domain via the public WHOIS database. You
// certify that you are the registrant of this domain name and have the authority
// to make this decision. You may withdraw your consent at any time by enabling
// privacy protection using either UpdateDomainContactPrivacy or the Route 53
// console. Enabling privacy protection removes the contact information provided
// for this domain from the WHOIS database. For more information on our privacy
// practices, see [https://aws.amazon.com/privacy/].
//
// [https://aws.amazon.com/privacy/]: https://aws.amazon.com/privacy/
// [GetOperationDetail]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_GetOperationDetail.html
