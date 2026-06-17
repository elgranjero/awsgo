package sqs

// DeleteMessageBatch is generated as a reference stub.
// Executable command wiring lives under cmd/sqs.go.
//
// Deletes up to ten messages from the specified queue. This is a batch version of DeleteMessage
// . The result of the action on each message is reported individually in the
// response.
//
// Because the batch request can result in a combination of successful and
// unsuccessful actions, you should check for batch errors even when the call
// returns an HTTP status code of 200 .
