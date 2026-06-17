package sqs

// ChangeMessageVisibilityBatch is generated as a reference stub.
// Executable command wiring lives under cmd/sqs.go.
//
// Changes the visibility timeout of multiple messages. This is a batch version of ChangeMessageVisibility
// . The result of the action on each message is reported individually in the
// response. You can send up to 10 ChangeMessageVisibilityrequests with each ChangeMessageVisibilityBatch
// action.
//
// Because the batch request can result in a combination of successful and
// unsuccessful actions, you should check for batch errors even when the call
// returns an HTTP status code of 200 .
