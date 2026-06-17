package lexmodelsv2

// ListAggregatedUtterances is generated as a reference stub.
// Executable command wiring lives under cmd/lexmodelsv2.go.
//
// Provides a list of utterances that users have sent to the bot.
//
// Utterances are aggregated by the text of the utterance. For example, all
// instances where customers used the phrase "I want to order pizza" are aggregated
// into the same line in the response.
//
// You can see both detected utterances and missed utterances. A detected
// utterance is where the bot properly recognized the utterance and activated the
// associated intent. A missed utterance was not recognized by the bot and didn't
// activate an intent.
//
// Utterances can be aggregated for a bot alias or for a bot version, but not both
// at the same time.
//
// Utterances statistics are not generated under the following conditions:
//
// - The childDirected field was set to true when the bot was created.
//
// - You are using slot obfuscation with one or more slots.
//
// - You opted out of participating in improving Amazon Lex.
