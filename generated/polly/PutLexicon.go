package polly

// PutLexicon is generated as a reference stub.
// Executable command wiring lives under cmd/polly.go.
//
// Stores a pronunciation lexicon in an Amazon Web Services Region. If a lexicon
// with the same name already exists in the region, it is overwritten by the new
// lexicon. Lexicon operations have eventual consistency, therefore, it might take
// some time before the lexicon is available to the SynthesizeSpeech operation.
//
// For more information, see [Managing Lexicons].
//
// [Managing Lexicons]: https://docs.aws.amazon.com/polly/latest/dg/managing-lexicons.html
