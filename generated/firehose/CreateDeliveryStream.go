package firehose

// CreateDeliveryStream is generated as a reference stub.
// Executable command wiring lives under cmd/firehose.go.
//
// Creates a Firehose stream.
//
// By default, you can create up to 5,000 Firehose streams per Amazon Web Services
// Region.
//
// This is an asynchronous operation that immediately returns. The initial status
// of the Firehose stream is CREATING . After the Firehose stream is created, its
// status is ACTIVE and it now accepts data. If the Firehose stream creation
// fails, the status transitions to CREATING_FAILED . Attempts to send data to a
// delivery stream that is not in the ACTIVE state cause an exception. To check
// the state of a Firehose stream, use DescribeDeliveryStream.
//
// If the status of a Firehose stream is CREATING_FAILED , this status doesn't
// change, and you can't invoke CreateDeliveryStream again on it. However, you can
// invoke the DeleteDeliveryStreamoperation to delete it.
//
// A Firehose stream can be configured to receive records directly from providers
// using PutRecordor PutRecordBatch, or it can be configured to use an existing Kinesis stream as its
// source. To specify a Kinesis data stream as input, set the DeliveryStreamType
// parameter to KinesisStreamAsSource , and provide the Kinesis stream Amazon
// Resource Name (ARN) and role ARN in the KinesisStreamSourceConfiguration
// parameter.
//
// To create a Firehose stream with server-side encryption (SSE) enabled, include DeliveryStreamEncryptionConfigurationInput
// in your request. This is optional. You can also invoke StartDeliveryStreamEncryptionto turn on SSE for an
// existing Firehose stream that doesn't have SSE enabled.
//
// A Firehose stream is configured with a single destination, such as Amazon
// Simple Storage Service (Amazon S3), Amazon Redshift, Amazon OpenSearch Service,
// Amazon OpenSearch Serverless, Splunk, and any custom HTTP endpoint or HTTP
// endpoints owned by or supported by third-party service providers, including
// Datadog, Dynatrace, LogicMonitor, MongoDB, New Relic, and Sumo Logic. You must
// specify only one of the following destination configuration parameters:
// ExtendedS3DestinationConfiguration , S3DestinationConfiguration ,
// ElasticsearchDestinationConfiguration , RedshiftDestinationConfiguration , or
// SplunkDestinationConfiguration .
//
// When you specify S3DestinationConfiguration , you can also provide the following
// optional values: BufferingHints, EncryptionConfiguration , and CompressionFormat
// . By default, if no BufferingHints value is provided, Firehose buffers data up
// to 5 MB or for 5 minutes, whichever condition is satisfied first. BufferingHints
// is a hint, so there are some cases where the service cannot adhere to these
// conditions strictly. For example, record boundaries might be such that the size
// is a little over or under the configured buffering size. By default, no
// encryption is performed. We strongly recommend that you enable encryption to
// ensure secure data storage in Amazon S3.
//
// A few notes about Amazon Redshift as a destination:
//
// - An Amazon Redshift destination requires an S3 bucket as intermediate
// location. Firehose first delivers data to Amazon S3 and then uses COPY syntax
// to load data into an Amazon Redshift table. This is specified in the
// RedshiftDestinationConfiguration.S3Configuration parameter.
//
// - The compression formats SNAPPY or ZIP cannot be specified in
// RedshiftDestinationConfiguration.S3Configuration because the Amazon Redshift
// COPY operation that reads from the S3 bucket doesn't support these compression
// formats.
//
// - We strongly recommend that you use the user name and password you provide
// exclusively with Firehose, and that the permissions for the account are
// restricted for Amazon Redshift INSERT permissions.
//
// Firehose assumes the IAM role that is configured as part of the destination.
// The role should allow the Firehose principal to assume the role, and the role
// should have permissions that allow the service to deliver the data. For more
// information, see [Grant Firehose Access to an Amazon S3 Destination]in the Amazon Firehose Developer Guide.
//
// [Grant Firehose Access to an Amazon S3 Destination]: https://docs.aws.amazon.com/firehose/latest/dev/controlling-access.html#using-iam-s3
