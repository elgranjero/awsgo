package lexmodelsv2

// DeleteUtterances is generated as a reference stub.
// Executable command wiring lives under cmd/lexmodelsv2.go.
//
// Deletes stored utterances.
//
// Amazon Lex stores the utterances that users send to your bot. Utterances are
// stored for 15 days for use with the [ListAggregatedUtterances]operation, and then stored indefinitely for
// use in improving the ability of your bot to respond to user input..
//
// Use the DeleteUtterances operation to manually delete utterances for a specific
// session. When you use the DeleteUtterances operation, utterances stored for
// improving your bot's ability to respond to user input are deleted immediately.
// Utterances stored for use with the ListAggregatedUtterances operation are
// deleted after 15 days.
//
// [ListAggregatedUtterances]: https://docs.aws.amazon.com/lexv2/latest/APIReference/API_ListAggregatedUtterances.html
