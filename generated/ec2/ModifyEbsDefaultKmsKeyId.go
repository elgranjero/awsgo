package ec2

// ModifyEbsDefaultKmsKeyId is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Changes the default KMS key for EBS encryption by default for your account in
// this Region.
//
// Amazon Web Services creates a unique Amazon Web Services managed KMS key in
// each Region for use with encryption by default. If you change the default KMS
// key to a symmetric customer managed KMS key, it is used instead of the Amazon
// Web Services managed KMS key. Amazon EBS does not support asymmetric KMS keys.
//
// If you delete or disable the customer managed KMS key that you specified for
// use with encryption by default, your instances will fail to launch.
//
// For more information, see [Amazon EBS encryption] in the Amazon EBS User Guide.
//
// [Amazon EBS encryption]: https://docs.aws.amazon.com/ebs/latest/userguide/ebs-encryption.html
