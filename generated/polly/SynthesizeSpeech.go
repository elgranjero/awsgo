package polly

// SynthesizeSpeech is generated as a reference stub.
// Executable command wiring lives under cmd/polly.go.
//
// Synthesizes UTF-8 input, plain text or SSML, to a stream of bytes. SSML input
// must be valid, well-formed SSML. Some alphabets might not be available with all
// the voices (for example, Cyrillic might not be read at all by English voices)
// unless phoneme mapping is used. For more information, see [How it Works].
//
// [How it Works]: https://docs.aws.amazon.com/polly/latest/dg/how-text-to-speech-works.html
