package lexruntimev2

// RecognizeText is generated as a reference stub.
// Executable command wiring lives under cmd/lexruntimev2.go.
//
// Sends user input to Amazon Lex V2. Client applications use this API to send
// requests to Amazon Lex V2 at runtime. Amazon Lex V2 then interprets the user
// input using the machine learning model that it build for the bot.
//
// In response, Amazon Lex V2 returns the next message to convey to the user and
// an optional response card to display.
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
