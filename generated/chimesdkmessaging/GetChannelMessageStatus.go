package chimesdkmessaging

// GetChannelMessageStatus is generated as a reference stub.
// Executable command wiring lives under cmd/chimesdkmessaging.go.
//
// Gets message status for a specified messageId . Use this API to determine the
// intermediate status of messages going through channel flow processing. The API
// provides an alternative to retrieving message status if the event was not
// received because a client wasn't connected to a websocket.
//
// Messages can have any one of these statuses.
//
// # SENT Message processed successfully
//
// # PENDING Ongoing processing
//
// # FAILED Processing failed
//
// DENIED Message denied by the processor
//
// - This API does not return statuses for denied messages, because we don't
// store them once the processor denies them.
//
// - Only the message sender can invoke this API.
//
// - The x-amz-chime-bearer request header is mandatory. Use the ARN of the
// AppInstanceUser or AppInstanceBot that makes the API call as the value in the
// header.
