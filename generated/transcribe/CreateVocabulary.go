package transcribe

// CreateVocabulary is generated as a reference stub.
// Executable command wiring lives under cmd/transcribe.go.
//
// Creates a new custom vocabulary.
//
// When creating a new custom vocabulary, you can either upload a text file that
// contains your new entries, phrases, and terms into an Amazon S3 bucket and
// include the URI in your request. Or you can include a list of terms directly in
// your request using the Phrases flag.
//
// Each language has a character set that contains all allowed characters for that
// specific language. If you use unsupported characters, your custom vocabulary
// request fails. Refer to [Character Sets for Custom Vocabularies]to get the character set for your language.
//
// For more information, see [Custom vocabularies].
//
// [Custom vocabularies]: https://docs.aws.amazon.com/transcribe/latest/dg/custom-vocabulary.html
// [Character Sets for Custom Vocabularies]: https://docs.aws.amazon.com/transcribe/latest/dg/charsets.html
