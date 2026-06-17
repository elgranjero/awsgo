package lexruntimev2

// RecognizeUtterance is generated as a reference stub.
// Executable command wiring lives under cmd/lexruntimev2.go.
//
// Sends user input to Amazon Lex V2. You can send text or speech. Clients use
// this API to send text and audio requests to Amazon Lex V2 at runtime. Amazon Lex
// V2 interprets the user input using the machine learning model built for the bot.
//
// The following request fields must be compressed with gzip and then base64
// encoded before you send them to Amazon Lex V2.
//
// - requestAttributes
//
// - sessionState
//
// The following response fields are compressed using gzip and then base64 encoded
// by Amazon Lex V2. Before you can use these fields, you must decode and
// decompress them.
//
// - inputTranscript
//
// - interpretations
//
// - messages
//
// - requestAttributes
//
// - sessionState
//
// The example contains a Java application that compresses and encodes a Java
// object to send to Amazon Lex V2, and a second that decodes and decompresses a
// response from Amazon Lex V2.
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
// [PostFulfillmentStatusSpecification]: https://docs.aws.amazon.com/lexv2/latest/dg/API_PostFulfillmentStatusSpecification.html
// [Completion message]: https://docs.aws.amazon.com/lexv2/latest/dg/streaming-progress.html#progress-complete.html
