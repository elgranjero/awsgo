package kinesis

// StopStreamEncryption is generated as a reference stub.
// Executable command wiring lives under cmd/kinesis.go.
//
// Disables server-side encryption for a specified stream.
//
// When invoking this API, you must use either the StreamARN or the StreamName
// parameter, or both. It is recommended that you use the StreamARN input
// parameter when you invoke this API.
//
// Stopping encryption is an asynchronous operation. Upon receiving the request,
// Kinesis Data Streams returns immediately and sets the status of the stream to
// UPDATING . After the update is complete, Kinesis Data Streams sets the status of
// the stream back to ACTIVE . Stopping encryption normally takes a few seconds to
// complete, but it can take minutes. You can continue to read and write data to
// your stream while its status is UPDATING . Once the status of the stream is
// ACTIVE , records written to the stream are no longer encrypted by Kinesis Data
// Streams.
//
// API Limits: You can successfully disable server-side encryption 25 times in a
// rolling 24-hour period.
//
// Note: It can take up to 5 seconds after the stream is in an ACTIVE status
// before all records written to the stream are no longer subject to encryption.
// After you disabled encryption, you can verify that encryption is not applied by
// inspecting the API response from PutRecord or PutRecords .
