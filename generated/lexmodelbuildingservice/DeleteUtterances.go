package lexmodelbuildingservice

// DeleteUtterances is generated as a reference stub.
// Executable command wiring lives under cmd/lexmodelbuildingservice.go.
//
// Deletes stored utterances.
//
// Amazon Lex stores the utterances that users send to your bot. Utterances are
// stored for 15 days for use with the GetUtterancesViewoperation, and then stored indefinitely for
// use in improving the ability of your bot to respond to user input.
//
// Use the DeleteUtterances operation to manually delete stored utterances for a
// specific user. When you use the DeleteUtterances operation, utterances stored
// for improving your bot's ability to respond to user input are deleted
// immediately. Utterances stored for use with the GetUtterancesView operation are
// deleted after 15 days.
//
// This operation requires permissions for the lex:DeleteUtterances action.
