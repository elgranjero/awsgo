package ec2

// EnableEbsEncryptionByDefault is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Enables EBS encryption by default for your account in the current Region.
//
// After you enable encryption by default, the EBS volumes that you create are
// always encrypted, either using the default KMS key or the KMS key that you
// specified when you created each volume. For more information, see [Amazon EBS encryption]in the Amazon
// EBS User Guide.
//
// Enabling encryption by default has no effect on the encryption status of your
// existing volumes.
//
// After you enable encryption by default, you can no longer launch instances
// using instance types that do not support encryption. For more information, see [Supported instance types].
//
// [Amazon EBS encryption]: https://docs.aws.amazon.com/ebs/latest/userguide/ebs-encryption.html
// [Supported instance types]: https://docs.aws.amazon.com/ebs/latest/userguide/ebs-encryption-requirements.html#ebs-encryption_supported_instances
