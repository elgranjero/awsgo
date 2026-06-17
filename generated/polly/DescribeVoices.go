package polly

// DescribeVoices is generated as a reference stub.
// Executable command wiring lives under cmd/polly.go.
//
// Returns the list of voices that are available for use when requesting speech
// synthesis. Each voice speaks a specified language, is either male or female, and
// is identified by an ID, which is the ASCII version of the voice name.
//
// When synthesizing speech ( SynthesizeSpeech ), you provide the voice ID for the
// voice you want from the list of voices returned by DescribeVoices .
//
// For example, you want your news reader application to read news in a specific
// language, but giving a user the option to choose the voice. Using the
// DescribeVoices operation you can provide the user with a list of available
// voices to select from.
//
// You can optionally specify a language code to filter the available voices. For
// example, if you specify en-US , the operation returns a list of all available US
// English voices.
//
// This operation requires permissions to perform the polly:DescribeVoices action.
