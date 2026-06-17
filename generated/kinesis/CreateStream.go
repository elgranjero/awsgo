package kinesis

// CreateStream is generated as a reference stub.
// Executable command wiring lives under cmd/kinesis.go.
//
// Creates a Kinesis data stream. A stream captures and transports data records
// that are continuously emitted from different data sources or producers.
// Scale-out within a stream is explicitly supported by means of shards, which are
// uniquely identified groups of data records in a stream.
//
// You can create your data stream using either on-demand or provisioned capacity
// mode. Data streams with an on-demand mode require no capacity planning and
// automatically scale to handle gigabytes of write and read throughput per minute.
// With the on-demand mode, Kinesis Data Streams automatically manages the shards
// in order to provide the necessary throughput.
//
// If you'd still like to proactively scale your on-demand data stream’s capacity,
// you can unlock the warm throughput feature for on-demand data streams by
// enabling MinimumThroughputBillingCommitment for your account. Once your account
// has MinimumThroughputBillingCommitment enabled, you can specify the warm
// throughput in MiB per second that your stream can support in writes.
//
// For the data streams with a provisioned mode, you must specify the number of
// shards for the data stream. Each shard can support reads up to five transactions
// per second, up to a maximum data read total of 2 MiB per second. Each shard can
// support writes up to 1,000 records per second, up to a maximum data write total
// of 1 MiB per second. If the amount of data input increases or decreases, you can
// add or remove shards.
//
// The stream name identifies the stream. The name is scoped to the Amazon Web
// Services account used by the application. It is also scoped by Amazon Web
// Services Region. That is, two streams in two different accounts can have the
// same name, and two streams in the same account, but in two different Regions,
// can have the same name.
//
// CreateStream is an asynchronous operation. Upon receiving a CreateStream
// request, Kinesis Data Streams immediately returns and sets the stream status to
// CREATING . After the stream is created, Kinesis Data Streams sets the stream
// status to ACTIVE . You should perform read and write operations only on an
// ACTIVE stream.
//
// You receive a LimitExceededException when making a CreateStream request when
// you try to do one of the following:
//
// - Have more than five streams in the CREATING state at any point in time.
//
// - Create more shards than are authorized for your account.
//
// For the default shard or on-demand throughput limits for an Amazon Web Services
// account, see [Amazon Kinesis Data Streams Limits]in the Amazon Kinesis Data Streams Developer Guide. To increase
// this limit, [contact Amazon Web Services Support].
//
// You can use DescribeStreamSummary to check the stream status, which is returned in StreamStatus .
//
// CreateStreamhas a limit of five transactions per second per account.
//
// You can add tags to the stream when making a CreateStream request by setting
// the Tags parameter. If you pass the Tags parameter, in addition to having the
// kinesis:CreateStream permission, you must also have the kinesis:AddTagsToStream
// permission for the stream that will be created. The kinesis:TagResource
// permission won’t work to tag streams on creation. Tags will take effect from the
// CREATING status of the stream, but you can't make any updates to the tags until
// the stream is in ACTIVE state.
//
// [contact Amazon Web Services Support]: https://docs.aws.amazon.com/general/latest/gr/aws_service_limits.html
// [Amazon Kinesis Data Streams Limits]: https://docs.aws.amazon.com/kinesis/latest/dev/service-sizes-and-limits.html
