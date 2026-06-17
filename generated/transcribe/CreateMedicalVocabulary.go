package transcribe

// CreateMedicalVocabulary is generated as a reference stub.
// Executable command wiring lives under cmd/transcribe.go.
//
// Creates a new custom medical vocabulary.
//
// Before creating a new custom medical vocabulary, you must first upload a text
// file that contains your vocabulary table into an Amazon S3 bucket. Note that
// this differs from , where you can include a list of terms within your request
// using the Phrases flag; CreateMedicalVocabulary does not support the Phrases
// flag and only accepts vocabularies in table format.
//
// Each language has a character set that contains all allowed characters for that
// specific language. If you use unsupported characters, your custom vocabulary
// request fails. Refer to [Character Sets for Custom Vocabularies]to get the character set for your language.
//
// For more information, see [Custom vocabularies].
//
// [Custom vocabularies]: https://docs.aws.amazon.com/transcribe/latest/dg/custom-vocabulary.html
// [Character Sets for Custom Vocabularies]: https://docs.aws.amazon.com/transcribe/latest/dg/charsets.html
