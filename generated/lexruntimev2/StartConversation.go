package lexruntimev2

// StartConversation is generated as a reference stub.
// Executable command wiring lives under cmd/lexruntimev2.go.
//
// Starts an HTTP/2 bidirectional event stream that enables you to send audio,
// text, or DTMF input in real time. After your application starts a conversation,
// users send input to Amazon Lex V2 as a stream of events. Amazon Lex V2 processes
// the incoming events and responds with streaming text or audio events.
//
// Audio input must be in the following format: audio/lpcm sample-rate=8000
// sample-size-bits=16 channel-count=1; is-big-endian=false .
//
// If the optional post-fulfillment response is specified, the messages are
// returned as follows. For more information, see [PostFulfillmentStatusSpecification].
//
// - Success message - Returned if the Lambda function completes successfully
// and the intent state is fulfilled or ready fulfillment if the message is
// present.
//
// - Failed message - The failed message is returned if the Lambda function
// throws an exception or if the Lambda function returns a failed intent state
// without a message.
//
// - Timeout message - If you don't configure a timeout message and a timeout,
// and the Lambda function doesn't return within 30 seconds, the timeout message is
// returned. If you configure a timeout, the timeout message is returned when the
// period times out.
//
// For more information, see [Completion message].
//
// If the optional update message is configured, it is played at the specified
// frequency while the Lambda function is running and the update message state is
// active. If the fulfillment update message is not active, the Lambda function
// runs with a 30 second timeout.
//
// For more information, see [Update message]
//
// The StartConversation operation is supported only in the following SDKs:
//
// [AWS SDK for C++]
//
// [AWS SDK for Java V2]
//
// [AWS SDK for Ruby V3]
//
// [AWS SDK for Ruby V3]: https://docs.aws.amazon.com/goto/SdkForRubyV3/runtime.lex.v2-2020-08-07/StartConversation
// [PostFulfillmentStatusSpecification]: https://docs.aws.amazon.com/lexv2/latest/dg/API_PostFulfillmentStatusSpecification.html
// [AWS SDK for Java V2]: https://docs.aws.amazon.com/goto/SdkForJavaV2/runtime.lex.v2-2020-08-07/StartConversation
// [AWS SDK for C++]: https://docs.aws.amazon.com/goto/SdkForCpp/runtime.lex.v2-2020-08-07/StartConversation
// [Update message]: https://docs.aws.amazon.com/lexv2/latest/dg/streaming-progress.html#progress-update.html
// [Completion message]: https://docs.aws.amazon.com/lexv2/latest/dg/streaming-progress.html#progress-complete.html
