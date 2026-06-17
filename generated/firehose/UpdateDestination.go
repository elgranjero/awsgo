package firehose

// UpdateDestination is generated as a reference stub.
// Executable command wiring lives under cmd/firehose.go.
//
// Updates the specified destination of the specified Firehose stream.
//
// Use this operation to change the destination type (for example, to replace the
// Amazon S3 destination with Amazon Redshift) or change the parameters associated
// with a destination (for example, to change the bucket name of the Amazon S3
// destination). The update might not occur immediately. The target Firehose stream
// remains active while the configurations are updated, so data writes to the
// Firehose stream can continue during this process. The updated configurations are
// usually effective within a few minutes.
//
// Switching between Amazon OpenSearch Service and other services is not
// supported. For an Amazon OpenSearch Service destination, you can only update to
// another Amazon OpenSearch Service destination.
//
// If the destination type is the same, Firehose merges the configuration
// parameters specified with the destination configuration that already exists on
// the delivery stream. If any of the parameters are not specified in the call, the
// existing values are retained. For example, in the Amazon S3 destination, if EncryptionConfigurationis
// not specified, then the existing EncryptionConfiguration is maintained on the
// destination.
//
// If the destination type is not the same, for example, changing the destination
// from Amazon S3 to Amazon Redshift, Firehose does not merge any parameters. In
// this case, all parameters must be specified.
//
// Firehose uses CurrentDeliveryStreamVersionId to avoid race conditions and
// conflicting merges. This is a required field, and the service updates the
// configuration only if the existing configuration has a version ID that matches.
// After the update is applied successfully, the version ID is updated, and can be
// retrieved using DescribeDeliveryStream. Use the new version ID to set CurrentDeliveryStreamVersionId
// in the next call.
