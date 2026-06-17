package acmpca

// GetCertificate is generated as a reference stub.
// Executable command wiring lives under cmd/acmpca.go.
//
// Retrieves a certificate from your private CA or one that has been shared with
// you. The ARN of the certificate is returned when you call the [IssueCertificate]action. You must
// specify both the ARN of your private CA and the ARN of the issued certificate
// when calling the GetCertificate action. You can retrieve the certificate if it
// is in the ISSUED, EXPIRED, or REVOKED state. You can call the [CreateCertificateAuthorityAuditReport]action to create
// a report that contains information about all of the certificates issued and
// revoked by your private CA.
//
// [IssueCertificate]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_IssueCertificate.html
// [CreateCertificateAuthorityAuditReport]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_CreateCertificateAuthorityAuditReport.html
