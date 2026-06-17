package ec2

// DisassociateEnclaveCertificateIamRole is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Disassociates an IAM role from an Certificate Manager (ACM) certificate.
// Disassociating an IAM role from an ACM certificate removes the Amazon S3 object
// that contains the certificate, certificate chain, and encrypted private key from
// the Amazon S3 bucket. It also revokes the IAM role's permission to use the KMS
// key used to encrypt the private key. This effectively revokes the role's
// permission to use the certificate.
