package ec2

// AssociateEnclaveCertificateIamRole is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Associates an Identity and Access Management (IAM) role with an Certificate
// Manager (ACM) certificate. This enables the certificate to be used by the ACM
// for Nitro Enclaves application inside an enclave. For more information, see [Certificate Manager for Nitro Enclaves]in
// the Amazon Web Services Nitro Enclaves User Guide.
//
// When the IAM role is associated with the ACM certificate, the certificate,
// certificate chain, and encrypted private key are placed in an Amazon S3 location
// that only the associated IAM role can access. The private key of the certificate
// is encrypted with an Amazon Web Services managed key that has an attached
// attestation-based key policy.
//
// To enable the IAM role to access the Amazon S3 object, you must grant it
// permission to call s3:GetObject on the Amazon S3 bucket returned by the
// command. To enable the IAM role to access the KMS key, you must grant it
// permission to call kms:Decrypt on the KMS key returned by the command. For more
// information, see [Grant the role permission to access the certificate and encryption key]in the Amazon Web Services Nitro Enclaves User Guide.
//
// [Certificate Manager for Nitro Enclaves]: https://docs.aws.amazon.com/enclaves/latest/user/nitro-enclave-refapp.html
// [Grant the role permission to access the certificate and encryption key]: https://docs.aws.amazon.com/enclaves/latest/user/nitro-enclave-refapp.html#add-policy
