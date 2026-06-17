package acmpca

// RevokeCertificate is generated as a reference stub.
// Executable command wiring lives under cmd/acmpca.go.
//
// Revokes a certificate that was issued inside Amazon Web Services Private CA. If
// you enable a certificate revocation list (CRL) when you create or update your
// private CA, information about the revoked certificates will be included in the
// CRL. Amazon Web Services Private CA writes the CRL to an S3 bucket that you
// specify. A CRL is typically updated approximately 30 minutes after a certificate
// is revoked. If for any reason the CRL update fails, Amazon Web Services Private
// CA attempts makes further attempts every 15 minutes. With Amazon CloudWatch, you
// can create alarms for the metrics CRLGenerated and MisconfiguredCRLBucket . For
// more information, see [Supported CloudWatch Metrics].
//
// Both Amazon Web Services Private CA and the IAM principal must have permission
// to write to the S3 bucket that you specify. If the IAM principal making the call
// does not have permission to write to the bucket, then an exception is thrown.
// For more information, see [Access policies for CRLs in Amazon S3].
//
// Amazon Web Services Private CA also writes revocation information to the audit
// report. For more information, see [CreateCertificateAuthorityAuditReport].
//
// You cannot revoke a root CA self-signed certificate.
//
// [Access policies for CRLs in Amazon S3]: https://docs.aws.amazon.com/privateca/latest/userguide/crl-planning.html#s3-policies
// [CreateCertificateAuthorityAuditReport]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_CreateCertificateAuthorityAuditReport.html
// [Supported CloudWatch Metrics]: https://docs.aws.amazon.com/privateca/latest/userguide/PcaCloudWatch.html
