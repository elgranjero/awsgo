package iotdataplane

// ListRetainedMessages is generated as a reference stub.
// Executable command wiring lives under cmd/iotdataplane.go.
//
// Lists summary information about the retained messages stored for the account.
//
// This action returns only the topic names of the retained messages. It doesn't
// return any message payloads. Although this action doesn't return a message
// payload, it can still incur messaging costs.
//
// To get the message payload of a retained message, call [GetRetainedMessage] with the topic name of
// the retained message.
//
// Requires permission to access the [ListRetainedMessages] action.
//
// For more information about messaging costs, see [Amazon Web Services IoT Core pricing - Messaging].
//
// [GetRetainedMessage]: https://docs.aws.amazon.com/iot/latest/apireference/API_iotdata_GetRetainedMessage.html
// [Amazon Web Services IoT Core pricing - Messaging]: http://aws.amazon.com/iot-core/pricing/#Messaging
// [ListRetainedMessages]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html
