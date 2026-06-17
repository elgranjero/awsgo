package transcribe

// CreateLanguageModel is generated as a reference stub.
// Executable command wiring lives under cmd/transcribe.go.
//
// Creates a new custom language model.
//
// When creating a new custom language model, you must specify:
//
// - If you want a Wideband (audio sample rates over 16,000 Hz) or Narrowband
// (audio sample rates under 16,000 Hz) base model
//
// - The location of your training and tuning files (this must be an Amazon S3
// URI)
//
// - The language of your model
//
// - A unique name for your model
