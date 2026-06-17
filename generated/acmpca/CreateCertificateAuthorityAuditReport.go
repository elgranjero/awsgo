package acmpca

// CreateCertificateAuthorityAuditReport is generated as a reference stub.
// Executable command wiring lives under cmd/acmpca.go.
//
// Creates an audit report that lists every time that your CA private key is used
// to issue a certificate. The [IssueCertificate]and [RevokeCertificate] actions use the private key.
//
// To save the audit report to your designated Amazon S3 bucket, you must create a
// bucket policy that grants Amazon Web Services Private CA permission to access
// and write to it. For an example policy, see [Prepare an Amazon S3 bucket for audit reports].
//
// Amazon Web Services Private CA assets that are stored in Amazon S3 can be
// protected with encryption. For more information, see [Encrypting Your Audit Reports].
//
// You can generate a maximum of one report every 30 minutes.
//
// [RevokeCertificate]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_RevokeCertificate.html
// [Encrypting Your Audit Reports]: https://docs.aws.amazon.com/privateca/latest/userguide/PcaAuditReport.html#audit-report-encryption
// [IssueCertificate]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_IssueCertificate.html
// [Prepare an Amazon S3 bucket for audit reports]: https://docs.aws.amazon.com/privateca/latest/userguide/PcaAuditReport.html#s3-access
