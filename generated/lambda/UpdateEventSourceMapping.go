package lambda

// UpdateEventSourceMapping is generated as a reference stub.
// Executable command wiring lives under cmd/lambda.go.
//
// Updates an event source mapping. You can change the function that Lambda
// invokes, or pause invocation and resume later from the same location.
//
// For details about how to configure different event sources, see the following
// topics.
//
// [Amazon DynamoDB Streams]
//
// [Amazon Kinesis]
//
// [Amazon SQS]
//
// [Amazon MQ and RabbitMQ]
//
// [Amazon MSK]
//
// [Apache Kafka]
//
// [Amazon DocumentDB]
//
// The following error handling options are available for stream sources
// (DynamoDB, Kinesis, Amazon MSK, and self-managed Apache Kafka):
//
// - BisectBatchOnFunctionError – If the function returns an error, split the
// batch in two and retry.
//
// - MaximumRecordAgeInSeconds – Discard records older than the specified age.
// The default value is infinite (-1). When set to infinite (-1), failed records
// are retried until the record expires
//
// - MaximumRetryAttempts – Discard records after the specified number of
// retries. The default value is infinite (-1). When set to infinite (-1), failed
// records are retried until the record expires.
//
// - OnFailure – Send discarded records to an Amazon SQS queue, Amazon SNS topic,
// Kafka topic, or Amazon S3 bucket. For more information, see [Adding a destination].
//
// The following option is available only for DynamoDB and Kinesis event sources:
//
// - ParallelizationFactor – Process multiple batches from each shard
// concurrently.
//
// For information about which configuration parameters apply to each event
// source, see the following topics.
//
// [Amazon DynamoDB Streams]
//
// [Amazon Kinesis]
//
// [Amazon SQS]
//
// [Amazon MQ and RabbitMQ]
//
// [Amazon MSK]
//
// [Apache Kafka]
//
// [Amazon DocumentDB]
//
// [Amazon DynamoDB Streams]: https://docs.aws.amazon.com/lambda/latest/dg/with-ddb.html#services-ddb-params
// [Amazon SQS]: https://docs.aws.amazon.com/lambda/latest/dg/with-sqs.html#services-sqs-params
// [Amazon MSK]: https://docs.aws.amazon.com/lambda/latest/dg/with-msk.html#services-msk-parms
// [Amazon Kinesis]: https://docs.aws.amazon.com/lambda/latest/dg/with-kinesis.html#services-kinesis-params
// [Amazon MQ and RabbitMQ]: https://docs.aws.amazon.com/lambda/latest/dg/with-mq.html#services-mq-params
// [Apache Kafka]: https://docs.aws.amazon.com/lambda/latest/dg/with-kafka.html#services-kafka-parms
// [Amazon DocumentDB]: https://docs.aws.amazon.com/lambda/latest/dg/with-documentdb.html#docdb-configuration
// [Adding a destination]: https://docs.aws.amazon.com/lambda/latest/dg/invocation-async-retain-records.html#invocation-async-destinations
