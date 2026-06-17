package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/firehose"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// firehoseCmd represents the firehose command
var _firehoseCmd = &cobra.Command{
	Use:   "firehose",
	Short: "AWS firehose CLI",
	Run: func(cmd *cobra.Command, args []string) {
		_awsOutput = resolveAWSOutput(_awsProfile, cmd.Flags().Changed("output"))
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := firehose.NewFromConfig(cfg)
		if _firehoseCreateDeliveryStream {
			firehose_CreateDeliveryStream(cfg, client)
			return
		}
		if _firehoseDeleteDeliveryStream {
			firehose_DeleteDeliveryStream(cfg, client)
			return
		}
		if _firehoseDescribeDeliveryStream {
			firehose_DescribeDeliveryStream(cfg, client)
			return
		}
		if _firehoseListDeliveryStreams {
			firehose_ListDeliveryStreams(cfg, client)
			return
		}
		if _firehoseListTagsForDeliveryStream {
			firehose_ListTagsForDeliveryStream(cfg, client)
			return
		}
		if _firehosePutRecord {
			firehose_PutRecord(cfg, client)
			return
		}
		if _firehosePutRecordBatch {
			firehose_PutRecordBatch(cfg, client)
			return
		}
		if _firehoseStartDeliveryStreamEncryption {
			firehose_StartDeliveryStreamEncryption(cfg, client)
			return
		}
		if _firehoseStopDeliveryStreamEncryption {
			firehose_StopDeliveryStreamEncryption(cfg, client)
			return
		}
		if _firehoseTagDeliveryStream {
			firehose_TagDeliveryStream(cfg, client)
			return
		}
		if _firehoseUntagDeliveryStream {
			firehose_UntagDeliveryStream(cfg, client)
			return
		}
		if _firehoseUpdateDestination {
			firehose_UpdateDestination(cfg, client)
			return
		}

	},
}

var (
	_firehoseCreateDeliveryStream          bool
	_firehoseDeleteDeliveryStream          bool
	_firehoseDescribeDeliveryStream        bool
	_firehoseListDeliveryStreams           bool
	_firehoseListTagsForDeliveryStream     bool
	_firehosePutRecord                     bool
	_firehosePutRecordBatch                bool
	_firehoseStartDeliveryStreamEncryption bool
	_firehoseStopDeliveryStreamEncryption  bool
	_firehoseTagDeliveryStream             bool
	_firehoseUntagDeliveryStream           bool
	_firehoseUpdateDestination             bool

	_firehoseAllowForceDelete                                   string
	_firehoseAmazonOpenSearchServerlessDestinationConfiguration string
	_firehoseAmazonOpenSearchServerlessDestinationUpdate        string
	_firehoseAmazonopensearchserviceDestinationConfiguration    string
	_firehoseAmazonopensearchserviceDestinationUpdate           string
	_firehoseCurrentDeliveryStreamVersionId                     string
	_firehoseDatabaseSourceConfiguration                        string
	_firehoseDeliveryStreamEncryptionConfigurationInput         string
	_firehoseDeliveryStreamName                                 string
	_firehoseDeliveryStreamType                                 string
	_firehoseDestinationId                                      string
	_firehoseDirectPutSourceConfiguration                       string
	_firehoseElasticsearchDestinationConfiguration              string
	_firehoseElasticsearchDestinationUpdate                     string
	_firehoseExclusiveStartDeliveryStreamName                   string
	_firehoseExclusiveStartDestinationId                        string
	_firehoseExclusiveStartTagKey                               string
	_firehoseExtendedS3DestinationConfiguration                 string
	_firehoseExtendedS3DestinationUpdate                        string
	_firehoseHttpEndpointDestinationConfiguration               string
	_firehoseHttpEndpointDestinationUpdate                      string
	_firehoseIcebergDestinationConfiguration                    string
	_firehoseIcebergDestinationUpdate                           string
	_firehoseKinesisStreamSourceConfiguration                   string
	_firehoseLimit                                              string
	_firehoseMSKSourceConfiguration                             string
	_firehoseRecord                                             string
	_firehoseRecords                                            string
	_firehoseRedshiftDestinationConfiguration                   string
	_firehoseRedshiftDestinationUpdate                          string
	_firehoseS3DestinationConfiguration                         string
	_firehoseS3DestinationUpdate                                string
	_firehoseSnowflakeDestinationConfiguration                  string
	_firehoseSnowflakeDestinationUpdate                         string
	_firehoseSplunkDestinationConfiguration                     string
	_firehoseSplunkDestinationUpdate                            string
	_firehoseTagKeys                                            []string
	_firehoseTags                                               string
)

// Creates a Firehose stream.
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
func firehose_CreateDeliveryStream(cfg aws.Config, client *firehose.Client) {
	input := &firehose.CreateDeliveryStreamInput{
		// DeliveryStreamName: *string, // Required
	}

	if len(_firehoseDeliveryStreamName) > 0 {
		input.DeliveryStreamName = aws.String(_firehoseDeliveryStreamName)
	}
	if len(_firehoseAmazonOpenSearchServerlessDestinationConfiguration) > 0 {
		if err := assignInputField(input, "AmazonOpenSearchServerlessDestinationConfiguration", _firehoseAmazonOpenSearchServerlessDestinationConfiguration); err != nil {
			log.Errorf("invalid --amazon-open-search-serverless-destination-configuration: %s", err.Error())
			return
		}
	}
	if len(_firehoseAmazonopensearchserviceDestinationConfiguration) > 0 {
		if err := assignInputField(input, "AmazonopensearchserviceDestinationConfiguration", _firehoseAmazonopensearchserviceDestinationConfiguration); err != nil {
			log.Errorf("invalid --amazonopensearchservice-destination-configuration: %s", err.Error())
			return
		}
	}
	if len(_firehoseDatabaseSourceConfiguration) > 0 {
		if err := assignInputField(input, "DatabaseSourceConfiguration", _firehoseDatabaseSourceConfiguration); err != nil {
			log.Errorf("invalid --database-source-configuration: %s", err.Error())
			return
		}
	}
	if len(_firehoseDeliveryStreamEncryptionConfigurationInput) > 0 {
		if err := assignInputField(input, "DeliveryStreamEncryptionConfigurationInput", _firehoseDeliveryStreamEncryptionConfigurationInput); err != nil {
			log.Errorf("invalid --delivery-stream-encryption-configuration-input: %s", err.Error())
			return
		}
	}
	if len(_firehoseDeliveryStreamType) > 0 {
		if err := assignInputField(input, "DeliveryStreamType", _firehoseDeliveryStreamType); err != nil {
			log.Errorf("invalid --delivery-stream-type: %s", err.Error())
			return
		}
	}
	if len(_firehoseDirectPutSourceConfiguration) > 0 {
		if err := assignInputField(input, "DirectPutSourceConfiguration", _firehoseDirectPutSourceConfiguration); err != nil {
			log.Errorf("invalid --direct-put-source-configuration: %s", err.Error())
			return
		}
	}
	if len(_firehoseElasticsearchDestinationConfiguration) > 0 {
		if err := assignInputField(input, "ElasticsearchDestinationConfiguration", _firehoseElasticsearchDestinationConfiguration); err != nil {
			log.Errorf("invalid --elasticsearch-destination-configuration: %s", err.Error())
			return
		}
	}
	if len(_firehoseExtendedS3DestinationConfiguration) > 0 {
		if err := assignInputField(input, "ExtendedS3DestinationConfiguration", _firehoseExtendedS3DestinationConfiguration); err != nil {
			log.Errorf("invalid --extended-s3-destination-configuration: %s", err.Error())
			return
		}
	}
	if len(_firehoseHttpEndpointDestinationConfiguration) > 0 {
		if err := assignInputField(input, "HttpEndpointDestinationConfiguration", _firehoseHttpEndpointDestinationConfiguration); err != nil {
			log.Errorf("invalid --http-endpoint-destination-configuration: %s", err.Error())
			return
		}
	}
	if len(_firehoseIcebergDestinationConfiguration) > 0 {
		if err := assignInputField(input, "IcebergDestinationConfiguration", _firehoseIcebergDestinationConfiguration); err != nil {
			log.Errorf("invalid --iceberg-destination-configuration: %s", err.Error())
			return
		}
	}
	if len(_firehoseKinesisStreamSourceConfiguration) > 0 {
		if err := assignInputField(input, "KinesisStreamSourceConfiguration", _firehoseKinesisStreamSourceConfiguration); err != nil {
			log.Errorf("invalid --kinesis-stream-source-configuration: %s", err.Error())
			return
		}
	}
	if len(_firehoseMSKSourceConfiguration) > 0 {
		if err := assignInputField(input, "MSKSourceConfiguration", _firehoseMSKSourceConfiguration); err != nil {
			log.Errorf("invalid --msk-source-configuration: %s", err.Error())
			return
		}
	}
	if len(_firehoseRedshiftDestinationConfiguration) > 0 {
		if err := assignInputField(input, "RedshiftDestinationConfiguration", _firehoseRedshiftDestinationConfiguration); err != nil {
			log.Errorf("invalid --redshift-destination-configuration: %s", err.Error())
			return
		}
	}
	if len(_firehoseS3DestinationConfiguration) > 0 {
		if err := assignInputField(input, "S3DestinationConfiguration", _firehoseS3DestinationConfiguration); err != nil {
			log.Errorf("invalid --s3-destination-configuration: %s", err.Error())
			return
		}
	}
	if len(_firehoseSnowflakeDestinationConfiguration) > 0 {
		if err := assignInputField(input, "SnowflakeDestinationConfiguration", _firehoseSnowflakeDestinationConfiguration); err != nil {
			log.Errorf("invalid --snowflake-destination-configuration: %s", err.Error())
			return
		}
	}
	if len(_firehoseSplunkDestinationConfiguration) > 0 {
		if err := assignInputField(input, "SplunkDestinationConfiguration", _firehoseSplunkDestinationConfiguration); err != nil {
			log.Errorf("invalid --splunk-destination-configuration: %s", err.Error())
			return
		}
	}
	if len(_firehoseTags) > 0 {
		if err := assignInputField(input, "Tags", _firehoseTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDeliveryStream(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a Firehose stream and its data.
// You can delete a Firehose stream only if it is in one of the following states:
// ACTIVE , DELETING , CREATING_FAILED , or DELETING_FAILED . You can't delete a
// Firehose stream that is in the CREATING state. To check the state of a Firehose
// stream, use DescribeDeliveryStream.
//
// DeleteDeliveryStream is an asynchronous API. When an API request to
// DeleteDeliveryStream succeeds, the Firehose stream is marked for deletion, and
// it goes into the DELETING state.While the Firehose stream is in the DELETING
// state, the service might continue to accept records, but it doesn't make any
// guarantees with respect to delivering the data. Therefore, as a best practice,
// first stop any applications that are sending records before you delete a
// Firehose stream.
//
// Removal of a Firehose stream that is in the DELETING state is a low priority
// operation for the service. A stream may remain in the DELETING state for
// several minutes. Therefore, as a best practice, applications should not wait for
// streams in the DELETING state to be removed.
func firehose_DeleteDeliveryStream(cfg aws.Config, client *firehose.Client) {
	input := &firehose.DeleteDeliveryStreamInput{
		// DeliveryStreamName: *string, // Required
	}

	if len(_firehoseDeliveryStreamName) > 0 {
		input.DeliveryStreamName = aws.String(_firehoseDeliveryStreamName)
	}
	if len(_firehoseAllowForceDelete) > 0 {
		if err := assignInputField(input, "AllowForceDelete", _firehoseAllowForceDelete); err != nil {
			log.Errorf("invalid --allow-force-delete: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteDeliveryStream(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the specified Firehose stream and its status. For example, after your
// Firehose stream is created, call DescribeDeliveryStream to see whether the
// Firehose stream is ACTIVE and therefore ready for data to be sent to it.
//
// If the status of a Firehose stream is CREATING_FAILED , this status doesn't
// change, and you can't invoke CreateDeliveryStreamagain on it. However, you can invoke the DeleteDeliveryStream
// operation to delete it. If the status is DELETING_FAILED , you can force
// deletion by invoking DeleteDeliveryStreamagain but with DeleteDeliveryStreamInput$AllowForceDelete set to true.
func firehose_DescribeDeliveryStream(cfg aws.Config, client *firehose.Client) {
	input := &firehose.DescribeDeliveryStreamInput{
		// DeliveryStreamName: *string, // Required
	}

	if len(_firehoseDeliveryStreamName) > 0 {
		input.DeliveryStreamName = aws.String(_firehoseDeliveryStreamName)
	}
	if len(_firehoseExclusiveStartDestinationId) > 0 {
		input.ExclusiveStartDestinationId = aws.String(_firehoseExclusiveStartDestinationId)
	}
	if len(_firehoseLimit) > 0 {
		if err := assignInputField(input, "Limit", _firehoseLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeDeliveryStream(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists your Firehose streams in alphabetical order of their names.
// The number of Firehose streams might be too large to return using a single call
// to ListDeliveryStreams . You can limit the number of Firehose streams returned,
// using the Limit parameter. To determine whether there are more delivery streams
// to list, check the value of HasMoreDeliveryStreams in the output. If there are
// more Firehose streams to list, you can request them by calling this operation
// again and setting the ExclusiveStartDeliveryStreamName parameter to the name of
// the last Firehose stream returned in the last call.
func firehose_ListDeliveryStreams(cfg aws.Config, client *firehose.Client) {
	input := &firehose.ListDeliveryStreamsInput{}

	if len(_firehoseDeliveryStreamType) > 0 {
		if err := assignInputField(input, "DeliveryStreamType", _firehoseDeliveryStreamType); err != nil {
			log.Errorf("invalid --delivery-stream-type: %s", err.Error())
			return
		}
	}
	if len(_firehoseExclusiveStartDeliveryStreamName) > 0 {
		input.ExclusiveStartDeliveryStreamName = aws.String(_firehoseExclusiveStartDeliveryStreamName)
	}
	if len(_firehoseLimit) > 0 {
		if err := assignInputField(input, "Limit", _firehoseLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListDeliveryStreams(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the tags for the specified Firehose stream. This operation has a limit of
// five transactions per second per account.
func firehose_ListTagsForDeliveryStream(cfg aws.Config, client *firehose.Client) {
	input := &firehose.ListTagsForDeliveryStreamInput{
		// DeliveryStreamName: *string, // Required
	}

	if len(_firehoseDeliveryStreamName) > 0 {
		input.DeliveryStreamName = aws.String(_firehoseDeliveryStreamName)
	}
	if len(_firehoseExclusiveStartTagKey) > 0 {
		input.ExclusiveStartTagKey = aws.String(_firehoseExclusiveStartTagKey)
	}
	if len(_firehoseLimit) > 0 {
		if err := assignInputField(input, "Limit", _firehoseLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListTagsForDeliveryStream(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Writes a single data record into an Firehose stream. To write multiple data
// records into a Firehose stream, use PutRecordBatch. Applications using these operations are
// referred to as producers.
//
// By default, each Firehose stream can take in up to 2,000 transactions per
// second, 5,000 records per second, or 5 MB per second. If you use PutRecordand PutRecordBatch, the
// limits are an aggregate across these two operations for each Firehose stream.
// For more information about limits and how to request an increase, see [Amazon Firehose Limits].
//
// Firehose accumulates and publishes a particular metric for a customer account
// in one minute intervals. It is possible that the bursts of incoming
// bytes/records ingested to a Firehose stream last only for a few seconds. Due to
// this, the actual spikes in the traffic might not be fully visible in the
// customer's 1 minute CloudWatch metrics.
//
// You must specify the name of the Firehose stream and the data record when using PutRecord
// . The data record consists of a data blob that can be up to 1,000 KiB in size,
// and any kind of data. For example, it can be a segment from a log file,
// geographic location data, website clickstream data, and so on.
//
// For multi record de-aggregation, you can not put more than 500 records even if
// the data blob length is less than 1000 KiB. If you include more than 500
// records, the request succeeds but the record de-aggregation doesn't work as
// expected and transformation lambda is invoked with the complete base64 encoded
// data blob instead of de-aggregated base64 decoded records.
//
// Firehose buffers records before delivering them to the destination. To
// disambiguate the data blobs at the destination, a common solution is to use
// delimiters in the data, such as a newline ( \n ) or some other character unique
// within the data. This allows the consumer application to parse individual data
// items when reading the data from the destination.
//
// The PutRecord operation returns a RecordId , which is a unique string assigned
// to each record. Producer applications can use this ID for purposes such as
// auditability and investigation.
//
// If the PutRecord operation throws a ServiceUnavailableException , the API is
// automatically reinvoked (retried) 3 times. If the exception persists, it is
// possible that the throughput limits have been exceeded for the Firehose stream.
//
// Re-invoking the Put API operations (for example, PutRecord and PutRecordBatch)
// can result in data duplicates. For larger data assets, allow for a longer time
// out before retrying Put API operations.
//
// Data records sent to Firehose are stored for 24 hours from the time they are
// added to a Firehose stream as it tries to send the records to the destination.
// If the destination is unreachable for more than 24 hours, the data is no longer
// available.
//
// Don't concatenate two or more base64 strings to form the data fields of your
// records. Instead, concatenate the raw data, then perform base64 encoding.
//
// [Amazon Firehose Limits]: https://docs.aws.amazon.com/firehose/latest/dev/limits.html
func firehose_PutRecord(cfg aws.Config, client *firehose.Client) {
	input := &firehose.PutRecordInput{
		// DeliveryStreamName: *string, // Required
		// Record: *types.Record, // Required
	}

	if len(_firehoseDeliveryStreamName) > 0 {
		input.DeliveryStreamName = aws.String(_firehoseDeliveryStreamName)
	}
	if len(_firehoseRecord) > 0 {
		if err := assignInputField(input, "Record", _firehoseRecord); err != nil {
			log.Errorf("invalid --record: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutRecord(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Writes multiple data records into a Firehose stream in a single call, which can
// achieve higher throughput per producer than when writing single records. To
// write single data records into a Firehose stream, use PutRecord. Applications using
// these operations are referred to as producers.
//
// Firehose accumulates and publishes a particular metric for a customer account
// in one minute intervals. It is possible that the bursts of incoming
// bytes/records ingested to a Firehose stream last only for a few seconds. Due to
// this, the actual spikes in the traffic might not be fully visible in the
// customer's 1 minute CloudWatch metrics.
//
// For information about service quota, see [Amazon Firehose Quota].
//
// Each PutRecordBatch request supports up to 500 records. Each record in the request can be as
// large as 1,000 KB (before base64 encoding), up to a limit of 4 MB for the entire
// request. These limits cannot be changed.
//
// You must specify the name of the Firehose stream and the data record when using PutRecord
// . The data record consists of a data blob that can be up to 1,000 KB in size,
// and any kind of data. For example, it could be a segment from a log file,
// geographic location data, website clickstream data, and so on.
//
// For multi record de-aggregation, you can not put more than 500 records even if
// the data blob length is less than 1000 KiB. If you include more than 500
// records, the request succeeds but the record de-aggregation doesn't work as
// expected and transformation lambda is invoked with the complete base64 encoded
// data blob instead of de-aggregated base64 decoded records.
//
// Firehose buffers records before delivering them to the destination. To
// disambiguate the data blobs at the destination, a common solution is to use
// delimiters in the data, such as a newline ( \n ) or some other character unique
// within the data. This allows the consumer application to parse individual data
// items when reading the data from the destination.
//
// The PutRecordBatch response includes a count of failed records, FailedPutCount , and an array
// of responses, RequestResponses . Even if the PutRecordBatch call succeeds, the value of
// FailedPutCount may be greater than 0, indicating that there are records for
// which the operation didn't succeed. Each entry in the RequestResponses array
// provides additional information about the processed record. It directly
// correlates with a record in the request array using the same ordering, from the
// top to the bottom. The response array always includes the same number of records
// as the request array. RequestResponses includes both successfully and
// unsuccessfully processed records. Firehose tries to process all records in each PutRecordBatch
// request. A single record failure does not stop the processing of subsequent
// records.
//
// A successfully processed record includes a RecordId value, which is unique for
// the record. An unsuccessfully processed record includes ErrorCode and
// ErrorMessage values. ErrorCode reflects the type of error, and is one of the
// following values: ServiceUnavailableException or InternalFailure . ErrorMessage
// provides more detailed information about the error.
//
// If there is an internal server error or a timeout, the write might have
// completed or it might have failed. If FailedPutCount is greater than 0, retry
// the request, resending only those records that might have failed processing.
// This minimizes the possible duplicate records and also reduces the total bytes
// sent (and corresponding charges). We recommend that you handle any duplicates at
// the destination.
//
// If PutRecordBatch throws ServiceUnavailableException , the API is automatically reinvoked
// (retried) 3 times. If the exception persists, it is possible that the throughput
// limits have been exceeded for the Firehose stream.
//
// Re-invoking the Put API operations (for example, PutRecord and PutRecordBatch)
// can result in data duplicates. For larger data assets, allow for a longer time
// out before retrying Put API operations.
//
// Data records sent to Firehose are stored for 24 hours from the time they are
// added to a Firehose stream as it attempts to send the records to the
// destination. If the destination is unreachable for more than 24 hours, the data
// is no longer available.
//
// Don't concatenate two or more base64 strings to form the data fields of your
// records. Instead, concatenate the raw data, then perform base64 encoding.
//
// [Amazon Firehose Quota]: https://docs.aws.amazon.com/firehose/latest/dev/limits.html
func firehose_PutRecordBatch(cfg aws.Config, client *firehose.Client) {
	input := &firehose.PutRecordBatchInput{
		// DeliveryStreamName: *string, // Required
		// Records: []types.Record, // Required
	}

	if len(_firehoseDeliveryStreamName) > 0 {
		input.DeliveryStreamName = aws.String(_firehoseDeliveryStreamName)
	}
	if len(_firehoseRecords) > 0 {
		if err := assignInputField(input, "Records", _firehoseRecords); err != nil {
			log.Errorf("invalid --records: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutRecordBatch(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables server-side encryption (SSE) for the Firehose stream.
// This operation is asynchronous. It returns immediately. When you invoke it,
// Firehose first sets the encryption status of the stream to ENABLING , and then
// to ENABLED . The encryption status of a Firehose stream is the Status property
// in DeliveryStreamEncryptionConfiguration. If the operation fails, the encryption status changes to ENABLING_FAILED .
// You can continue to read and write data to your Firehose stream while the
// encryption status is ENABLING , but the data is not encrypted. It can take up to
// 5 seconds after the encryption status changes to ENABLED before all records
// written to the Firehose stream are encrypted. To find out whether a record or a
// batch of records was encrypted, check the response elements PutRecordOutput$Encryptedand PutRecordBatchOutput$Encrypted, respectively.
//
// To check the encryption status of a Firehose stream, use DescribeDeliveryStream.
//
// Even if encryption is currently enabled for a Firehose stream, you can still
// invoke this operation on it to change the ARN of the CMK or both its type and
// ARN. If you invoke this method to change the CMK, and the old CMK is of type
// CUSTOMER_MANAGED_CMK , Firehose schedules the grant it had on the old CMK for
// retirement. If the new CMK is of type CUSTOMER_MANAGED_CMK , Firehose creates a
// grant that enables it to use the new CMK to encrypt and decrypt data and to
// manage the grant.
//
// For the KMS grant creation to be successful, the Firehose API operations
// StartDeliveryStreamEncryption and CreateDeliveryStream should not be called
// with session credentials that are more than 6 hours old.
//
// If a Firehose stream already has encryption enabled and then you invoke this
// operation to change the ARN of the CMK or both its type and ARN and you get
// ENABLING_FAILED , this only means that the attempt to change the CMK failed. In
// this case, encryption remains enabled with the old CMK.
//
// If the encryption status of your Firehose stream is ENABLING_FAILED , you can
// invoke this operation again with a valid CMK. The CMK must be enabled and the
// key policy mustn't explicitly deny the permission for Firehose to invoke KMS
// encrypt and decrypt operations.
//
// You can enable SSE for a Firehose stream only if it's a Firehose stream that
// uses DirectPut as its source.
//
// The StartDeliveryStreamEncryption and StopDeliveryStreamEncryption operations
// have a combined limit of 25 calls per Firehose stream per 24 hours. For example,
// you reach the limit if you call StartDeliveryStreamEncryption 13 times and
// StopDeliveryStreamEncryption 12 times for the same Firehose stream in a 24-hour
// period.
func firehose_StartDeliveryStreamEncryption(cfg aws.Config, client *firehose.Client) {
	input := &firehose.StartDeliveryStreamEncryptionInput{
		// DeliveryStreamName: *string, // Required
	}

	if len(_firehoseDeliveryStreamName) > 0 {
		input.DeliveryStreamName = aws.String(_firehoseDeliveryStreamName)
	}
	if len(_firehoseDeliveryStreamEncryptionConfigurationInput) > 0 {
		if err := assignInputField(input, "DeliveryStreamEncryptionConfigurationInput", _firehoseDeliveryStreamEncryptionConfigurationInput); err != nil {
			log.Errorf("invalid --delivery-stream-encryption-configuration-input: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartDeliveryStreamEncryption(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disables server-side encryption (SSE) for the Firehose stream.
// This operation is asynchronous. It returns immediately. When you invoke it,
// Firehose first sets the encryption status of the stream to DISABLING , and then
// to DISABLED . You can continue to read and write data to your stream while its
// status is DISABLING . It can take up to 5 seconds after the encryption status
// changes to DISABLED before all records written to the Firehose stream are no
// longer subject to encryption. To find out whether a record or a batch of records
// was encrypted, check the response elements PutRecordOutput$Encryptedand PutRecordBatchOutput$Encrypted, respectively.
//
// To check the encryption state of a Firehose stream, use DescribeDeliveryStream.
//
// If SSE is enabled using a customer managed CMK and then you invoke
// StopDeliveryStreamEncryption , Firehose schedules the related KMS grant for
// retirement and then retires it after it ensures that it is finished delivering
// records to the destination.
//
// The StartDeliveryStreamEncryption and StopDeliveryStreamEncryption operations
// have a combined limit of 25 calls per Firehose stream per 24 hours. For example,
// you reach the limit if you call StartDeliveryStreamEncryption 13 times and
// StopDeliveryStreamEncryption 12 times for the same Firehose stream in a 24-hour
// period.
func firehose_StopDeliveryStreamEncryption(cfg aws.Config, client *firehose.Client) {
	input := &firehose.StopDeliveryStreamEncryptionInput{
		// DeliveryStreamName: *string, // Required
	}

	if len(_firehoseDeliveryStreamName) > 0 {
		input.DeliveryStreamName = aws.String(_firehoseDeliveryStreamName)
	}

	if resp, err := client.StopDeliveryStreamEncryption(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds or updates tags for the specified Firehose stream. A tag is a key-value
// pair that you can define and assign to Amazon Web Services resources. If you
// specify a tag that already exists, the tag value is replaced with the value that
// you specify in the request. Tags are metadata. For example, you can add friendly
// names and descriptions or other types of information that can help you
// distinguish the Firehose stream. For more information about tags, see [Using Cost Allocation Tags]in the
// Amazon Web Services Billing and Cost Management User Guide.
//
// Each Firehose stream can have up to 50 tags.
//
// This operation has a limit of five transactions per second per account.
//
// [Using Cost Allocation Tags]: https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html
func firehose_TagDeliveryStream(cfg aws.Config, client *firehose.Client) {
	input := &firehose.TagDeliveryStreamInput{
		// DeliveryStreamName: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_firehoseDeliveryStreamName) > 0 {
		input.DeliveryStreamName = aws.String(_firehoseDeliveryStreamName)
	}
	if len(_firehoseTags) > 0 {
		if err := assignInputField(input, "Tags", _firehoseTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.TagDeliveryStream(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes tags from the specified Firehose stream. Removed tags are deleted, and
// you can't recover them after this operation successfully completes.
//
// If you specify a tag that doesn't exist, the operation ignores it.
//
// This operation has a limit of five transactions per second per account.
func firehose_UntagDeliveryStream(cfg aws.Config, client *firehose.Client) {
	input := &firehose.UntagDeliveryStreamInput{
		// DeliveryStreamName: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_firehoseDeliveryStreamName) > 0 {
		input.DeliveryStreamName = aws.String(_firehoseDeliveryStreamName)
	}
	if len(_firehoseTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _firehoseTagKeys...)
	}

	if resp, err := client.UntagDeliveryStream(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified destination of the specified Firehose stream.
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
func firehose_UpdateDestination(cfg aws.Config, client *firehose.Client) {
	input := &firehose.UpdateDestinationInput{
		// CurrentDeliveryStreamVersionId: *string, // Required
		// DeliveryStreamName: *string, // Required
		// DestinationId: *string, // Required
	}

	if len(_firehoseCurrentDeliveryStreamVersionId) > 0 {
		input.CurrentDeliveryStreamVersionId = aws.String(_firehoseCurrentDeliveryStreamVersionId)
	}
	if len(_firehoseDeliveryStreamName) > 0 {
		input.DeliveryStreamName = aws.String(_firehoseDeliveryStreamName)
	}
	if len(_firehoseDestinationId) > 0 {
		input.DestinationId = aws.String(_firehoseDestinationId)
	}
	if len(_firehoseAmazonOpenSearchServerlessDestinationUpdate) > 0 {
		if err := assignInputField(input, "AmazonOpenSearchServerlessDestinationUpdate", _firehoseAmazonOpenSearchServerlessDestinationUpdate); err != nil {
			log.Errorf("invalid --amazon-open-search-serverless-destination-update: %s", err.Error())
			return
		}
	}
	if len(_firehoseAmazonopensearchserviceDestinationUpdate) > 0 {
		if err := assignInputField(input, "AmazonopensearchserviceDestinationUpdate", _firehoseAmazonopensearchserviceDestinationUpdate); err != nil {
			log.Errorf("invalid --amazonopensearchservice-destination-update: %s", err.Error())
			return
		}
	}
	if len(_firehoseElasticsearchDestinationUpdate) > 0 {
		if err := assignInputField(input, "ElasticsearchDestinationUpdate", _firehoseElasticsearchDestinationUpdate); err != nil {
			log.Errorf("invalid --elasticsearch-destination-update: %s", err.Error())
			return
		}
	}
	if len(_firehoseExtendedS3DestinationUpdate) > 0 {
		if err := assignInputField(input, "ExtendedS3DestinationUpdate", _firehoseExtendedS3DestinationUpdate); err != nil {
			log.Errorf("invalid --extended-s3-destination-update: %s", err.Error())
			return
		}
	}
	if len(_firehoseHttpEndpointDestinationUpdate) > 0 {
		if err := assignInputField(input, "HttpEndpointDestinationUpdate", _firehoseHttpEndpointDestinationUpdate); err != nil {
			log.Errorf("invalid --http-endpoint-destination-update: %s", err.Error())
			return
		}
	}
	if len(_firehoseIcebergDestinationUpdate) > 0 {
		if err := assignInputField(input, "IcebergDestinationUpdate", _firehoseIcebergDestinationUpdate); err != nil {
			log.Errorf("invalid --iceberg-destination-update: %s", err.Error())
			return
		}
	}
	if len(_firehoseRedshiftDestinationUpdate) > 0 {
		if err := assignInputField(input, "RedshiftDestinationUpdate", _firehoseRedshiftDestinationUpdate); err != nil {
			log.Errorf("invalid --redshift-destination-update: %s", err.Error())
			return
		}
	}
	if len(_firehoseS3DestinationUpdate) > 0 {
		if err := assignInputField(input, "S3DestinationUpdate", _firehoseS3DestinationUpdate); err != nil {
			log.Errorf("invalid --s3-destination-update: %s", err.Error())
			return
		}
	}
	if len(_firehoseSnowflakeDestinationUpdate) > 0 {
		if err := assignInputField(input, "SnowflakeDestinationUpdate", _firehoseSnowflakeDestinationUpdate); err != nil {
			log.Errorf("invalid --snowflake-destination-update: %s", err.Error())
			return
		}
	}
	if len(_firehoseSplunkDestinationUpdate) > 0 {
		if err := assignInputField(input, "SplunkDestinationUpdate", _firehoseSplunkDestinationUpdate); err != nil {
			log.Errorf("invalid --splunk-destination-update: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateDestination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_firehoseCmd)
	_firehoseCmd.Flags().SortFlags = false

	_firehoseCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_firehoseCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_firehoseCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_firehoseCmd.Flags().StringVarP(&_firehoseAllowForceDelete, "allow-force-delete", "", "", "Allow Force Delete")
	_firehoseCmd.Flags().StringVarP(&_firehoseAmazonOpenSearchServerlessDestinationConfiguration, "amazon-open-search-serverless-destination-configuration", "", "", "Amazon Open Search Serverless Destination Configuration")
	_firehoseCmd.Flags().StringVarP(&_firehoseAmazonOpenSearchServerlessDestinationUpdate, "amazon-open-search-serverless-destination-update", "", "", "Amazon Open Search Serverless Destination Update")
	_firehoseCmd.Flags().StringVarP(&_firehoseAmazonopensearchserviceDestinationConfiguration, "amazonopensearchservice-destination-configuration", "", "", "Amazonopensearchservice Destination Configuration")
	_firehoseCmd.Flags().StringVarP(&_firehoseAmazonopensearchserviceDestinationUpdate, "amazonopensearchservice-destination-update", "", "", "Amazonopensearchservice Destination Update")
	_firehoseCmd.Flags().StringVarP(&_firehoseCurrentDeliveryStreamVersionId, "current-delivery-stream-version-id", "", "", "Current Delivery Stream Version ID")
	_firehoseCmd.Flags().StringVarP(&_firehoseDatabaseSourceConfiguration, "database-source-configuration", "", "", "Database Source Configuration")
	_firehoseCmd.Flags().StringVarP(&_firehoseDeliveryStreamEncryptionConfigurationInput, "delivery-stream-encryption-configuration-input", "", "", "Delivery Stream Encryption Configuration Input")
	_firehoseCmd.Flags().StringVarP(&_firehoseDeliveryStreamName, "delivery-stream-name", "", "", "Delivery Stream Name")
	_firehoseCmd.Flags().StringVarP(&_firehoseDeliveryStreamType, "delivery-stream-type", "", "", "Delivery Stream Type")
	_firehoseCmd.Flags().StringVarP(&_firehoseDestinationId, "destination-id", "", "", "Destination ID")
	_firehoseCmd.Flags().StringVarP(&_firehoseDirectPutSourceConfiguration, "direct-put-source-configuration", "", "", "Direct Put Source Configuration")
	_firehoseCmd.Flags().StringVarP(&_firehoseElasticsearchDestinationConfiguration, "elasticsearch-destination-configuration", "", "", "Elasticsearch Destination Configuration")
	_firehoseCmd.Flags().StringVarP(&_firehoseElasticsearchDestinationUpdate, "elasticsearch-destination-update", "", "", "Elasticsearch Destination Update")
	_firehoseCmd.Flags().StringVarP(&_firehoseExclusiveStartDeliveryStreamName, "exclusive-start-delivery-stream-name", "", "", "Exclusive Start Delivery Stream Name")
	_firehoseCmd.Flags().StringVarP(&_firehoseExclusiveStartDestinationId, "exclusive-start-destination-id", "", "", "Exclusive Start Destination ID")
	_firehoseCmd.Flags().StringVarP(&_firehoseExclusiveStartTagKey, "exclusive-start-tag-key", "", "", "Exclusive Start Tag Key")
	_firehoseCmd.Flags().StringVarP(&_firehoseExtendedS3DestinationConfiguration, "extended-s3-destination-configuration", "", "", "Extended S3 Destination Configuration")
	_firehoseCmd.Flags().StringVarP(&_firehoseExtendedS3DestinationUpdate, "extended-s3-destination-update", "", "", "Extended S3 Destination Update")
	_firehoseCmd.Flags().StringVarP(&_firehoseHttpEndpointDestinationConfiguration, "http-endpoint-destination-configuration", "", "", "HTTP Endpoint Destination Configuration")
	_firehoseCmd.Flags().StringVarP(&_firehoseHttpEndpointDestinationUpdate, "http-endpoint-destination-update", "", "", "HTTP Endpoint Destination Update")
	_firehoseCmd.Flags().StringVarP(&_firehoseIcebergDestinationConfiguration, "iceberg-destination-configuration", "", "", "Iceberg Destination Configuration")
	_firehoseCmd.Flags().StringVarP(&_firehoseIcebergDestinationUpdate, "iceberg-destination-update", "", "", "Iceberg Destination Update")
	_firehoseCmd.Flags().StringVarP(&_firehoseKinesisStreamSourceConfiguration, "kinesis-stream-source-configuration", "", "", "Kinesis Stream Source Configuration")
	_firehoseCmd.Flags().StringVarP(&_firehoseLimit, "limit", "", "", "Limit")
	_firehoseCmd.Flags().StringVarP(&_firehoseMSKSourceConfiguration, "msk-source-configuration", "", "", "Msk Source Configuration")
	_firehoseCmd.Flags().StringVarP(&_firehoseRecord, "record", "", "", "Record")
	_firehoseCmd.Flags().StringVarP(&_firehoseRecords, "records", "", "", "Records")
	_firehoseCmd.Flags().StringVarP(&_firehoseRedshiftDestinationConfiguration, "redshift-destination-configuration", "", "", "Redshift Destination Configuration")
	_firehoseCmd.Flags().StringVarP(&_firehoseRedshiftDestinationUpdate, "redshift-destination-update", "", "", "Redshift Destination Update")
	_firehoseCmd.Flags().StringVarP(&_firehoseS3DestinationConfiguration, "s3-destination-configuration", "", "", "S3 Destination Configuration")
	_firehoseCmd.Flags().StringVarP(&_firehoseS3DestinationUpdate, "s3-destination-update", "", "", "S3 Destination Update")
	_firehoseCmd.Flags().StringVarP(&_firehoseSnowflakeDestinationConfiguration, "snowflake-destination-configuration", "", "", "Snowflake Destination Configuration")
	_firehoseCmd.Flags().StringVarP(&_firehoseSnowflakeDestinationUpdate, "snowflake-destination-update", "", "", "Snowflake Destination Update")
	_firehoseCmd.Flags().StringVarP(&_firehoseSplunkDestinationConfiguration, "splunk-destination-configuration", "", "", "Splunk Destination Configuration")
	_firehoseCmd.Flags().StringVarP(&_firehoseSplunkDestinationUpdate, "splunk-destination-update", "", "", "Splunk Destination Update")
	_firehoseCmd.Flags().StringSliceVarP(&_firehoseTagKeys, "tag-keys", "", nil, "Tag Keys")
	_firehoseCmd.Flags().StringVarP(&_firehoseTags, "tags", "", "", "Tags")

	_firehoseCmd.Flags().BoolVarP(&_firehoseCreateDeliveryStream, "create-delivery-stream", "", false, "Create Delivery Stream")
	_firehoseCmd.Flags().BoolVarP(&_firehoseDeleteDeliveryStream, "delete-delivery-stream", "", false, "Delete Delivery Stream")
	_firehoseCmd.Flags().BoolVarP(&_firehoseDescribeDeliveryStream, "describe-delivery-stream", "", false, "Describe Delivery Stream")
	_firehoseCmd.Flags().BoolVarP(&_firehoseListDeliveryStreams, "list-delivery-streams", "", false, "List Delivery Streams")
	_firehoseCmd.Flags().BoolVarP(&_firehoseListTagsForDeliveryStream, "list-tags-for-delivery-stream", "", false, "List Tags For Delivery Stream")
	_firehoseCmd.Flags().BoolVarP(&_firehosePutRecord, "put-record", "", false, "Put Record")
	_firehoseCmd.Flags().BoolVarP(&_firehosePutRecordBatch, "put-record-batch", "", false, "Put Record Batch")
	_firehoseCmd.Flags().BoolVarP(&_firehoseStartDeliveryStreamEncryption, "start-delivery-stream-encryption", "", false, "Start Delivery Stream Encryption")
	_firehoseCmd.Flags().BoolVarP(&_firehoseStopDeliveryStreamEncryption, "stop-delivery-stream-encryption", "", false, "Stop Delivery Stream Encryption")
	_firehoseCmd.Flags().BoolVarP(&_firehoseTagDeliveryStream, "tag-delivery-stream", "", false, "Tag Delivery Stream")
	_firehoseCmd.Flags().BoolVarP(&_firehoseUntagDeliveryStream, "untag-delivery-stream", "", false, "Untag Delivery Stream")
	_firehoseCmd.Flags().BoolVarP(&_firehoseUpdateDestination, "update-destination", "", false, "Update Destination")

}
