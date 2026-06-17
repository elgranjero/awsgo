package acmpca

// CreateCertificateAuthority is generated as a reference stub.
// Executable command wiring lives under cmd/acmpca.go.
//
// Creates a root or subordinate private certificate authority (CA). You must
// specify the CA configuration, an optional configuration for Online Certificate
// Status Protocol (OCSP) and/or a certificate revocation list (CRL), the CA type,
// and an optional idempotency token to avoid accidental creation of multiple CAs.
// The CA configuration specifies the name of the algorithm and key size to be used
// to create the CA private key, the type of signing algorithm that the CA uses,
// and X.500 subject information. The OCSP configuration can optionally specify a
// custom URL for the OCSP responder. The CRL configuration specifies the CRL
// expiration period in days (the validity period of the CRL), the Amazon S3 bucket
// that will contain the CRL, and a CNAME alias for the S3 bucket that is included
// in certificates issued by the CA. If successful, this action returns the Amazon
// Resource Name (ARN) of the CA.
//
// Both Amazon Web Services Private CA and the IAM principal must have permission
// to write to the S3 bucket that you specify. If the IAM principal making the call
// does not have permission to write to the bucket, then an exception is thrown.
// For more information, see [Access policies for CRLs in Amazon S3].
//
// Amazon Web Services Private CA assets that are stored in Amazon S3 can be
// protected with encryption. For more information, see [Encrypting Your CRLs].
//
// [Access policies for CRLs in Amazon S3]: https://docs.aws.amazon.com/privateca/latest/userguide/crl-planning.html#s3-policies
// [Encrypting Your CRLs]: https://docs.aws.amazon.com/privateca/latest/userguide/crl-planning.html#crl-encryption
