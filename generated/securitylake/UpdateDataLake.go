package securitylake

// UpdateDataLake is generated as a reference stub.
// Executable command wiring lives under cmd/securitylake.go.
//
// You can use UpdateDataLake to specify where to store your security data, how it
// should be encrypted at rest and for how long. You can add a [Rollup Region]to consolidate data
// from multiple Amazon Web Services Regions, replace default encryption (SSE-S3)
// with [Customer Manged Key], or specify transition and expiration actions through storage [Lifecycle management]. The
// UpdateDataLake API works as an "upsert" operation that performs an insert if the
// specified item or record does not exist, or an update if it already exists.
// Security Lake securely stores your data at rest using Amazon Web Services
// encryption solutions. For more details, see [Data protection in Amazon Security Lake].
//
// For example, omitting the key encryptionConfiguration from a Region that is
// included in an update call that currently uses KMS will leave that Region's KMS
// key in place, but specifying encryptionConfiguration: {kmsKeyId:
// 'S3_MANAGED_KEY'} for that same Region will reset the key to S3-managed .
//
// For more details about lifecycle management and how to update retention
// settings for one or more Regions after enabling Security Lake, see the [Amazon Security Lake User Guide].
//
// [Lifecycle management]: https://docs.aws.amazon.com/security-lake/latest/userguide/lifecycle-management.html
// [Rollup Region]: https://docs.aws.amazon.com/security-lake/latest/userguide/manage-regions.html#add-rollup-region
// [Data protection in Amazon Security Lake]: https://docs.aws.amazon.com/security-lake/latest/userguide/data-protection.html
// [Amazon Security Lake User Guide]: https://docs.aws.amazon.com/security-lake/latest/userguide/lifecycle-management.html
// [Customer Manged Key]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#customer-cmk
