package firehose

// StopDeliveryStreamEncryption is generated as a reference stub.
// Executable command wiring lives under cmd/firehose.go.
//
// Disables server-side encryption (SSE) for the Firehose stream.
//
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
