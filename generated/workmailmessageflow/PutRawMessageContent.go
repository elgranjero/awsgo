package workmailmessageflow

// PutRawMessageContent is generated as a reference stub.
// Executable command wiring lives under cmd/workmailmessageflow.go.
//
// Updates the raw content of an in-transit email message, in MIME format.
//
// This example describes how to update in-transit email message. For more
// information and examples for using this API, see [Updating message content with AWS Lambda].
//
// Updates to an in-transit message only appear when you call PutRawMessageContent
// from an AWS Lambda function configured with a synchronous [Run Lambda]rule. If you call
// PutRawMessageContent on a delivered or sent message, the message remains
// unchanged, even though [GetRawMessageContent]returns an updated message.
//
// [GetRawMessageContent]: https://docs.aws.amazon.com/workmail/latest/APIReference/API_messageflow_GetRawMessageContent.html
// [Run Lambda]: https://docs.aws.amazon.com/workmail/latest/adminguide/lambda.html#synchronous-rules
// [Updating message content with AWS Lambda]: https://docs.aws.amazon.com/workmail/latest/adminguide/update-with-lambda.html
