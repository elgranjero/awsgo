package transcribe

// CreateVocabularyFilter is generated as a reference stub.
// Executable command wiring lives under cmd/transcribe.go.
//
// Creates a new custom vocabulary filter.
//
// You can use custom vocabulary filters to mask, delete, or flag specific words
// from your transcript. Custom vocabulary filters are commonly used to mask
// profanity in transcripts.
//
// Each language has a character set that contains all allowed characters for that
// specific language. If you use unsupported characters, your custom vocabulary
// filter request fails. Refer to [Character Sets for Custom Vocabularies]to get the character set for your language.
//
// For more information, see [Vocabulary filtering].
//
// [Character Sets for Custom Vocabularies]: https://docs.aws.amazon.com/transcribe/latest/dg/charsets.html
// [Vocabulary filtering]: https://docs.aws.amazon.com/transcribe/latest/dg/vocabulary-filtering.html
