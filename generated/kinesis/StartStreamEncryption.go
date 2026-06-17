package kinesis

// StartStreamEncryption is generated as a reference stub.
// Executable command wiring lives under cmd/kinesis.go.
//
// Enables or updates server-side encryption using an Amazon Web Services KMS key
// for a specified stream.
//
// When invoking this API, you must use either the StreamARN or the StreamName
// parameter, or both. It is recommended that you use the StreamARN input
// parameter when you invoke this API.
//
// Starting encryption is an asynchronous operation. Upon receiving the request,
// Kinesis Data Streams returns immediately and sets the status of the stream to
// UPDATING . After the update is complete, Kinesis Data Streams sets the status of
// the stream back to ACTIVE . Updating or applying encryption normally takes a few
// seconds to complete, but it can take minutes. You can continue to read and write
// data to your stream while its status is UPDATING . Once the status of the stream
// is ACTIVE , encryption begins for records written to the stream.
//
// API Limits: You can successfully apply a new Amazon Web Services KMS key for
// server-side encryption 25 times in a rolling 24-hour period.
//
// Note: It can take up to 5 seconds after the stream is in an ACTIVE status
// before all records written to the stream are encrypted. After you enable
// encryption, you can verify that encryption is applied by inspecting the API
// response from PutRecord or PutRecords .
