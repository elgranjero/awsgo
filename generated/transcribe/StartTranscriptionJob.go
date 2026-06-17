package transcribe

// StartTranscriptionJob is generated as a reference stub.
// Executable command wiring lives under cmd/transcribe.go.
//
// Transcribes the audio from a media file and applies any additional Request
// Parameters you choose to include in your request.
//
// To make a StartTranscriptionJob request, you must first upload your media file
// into an Amazon S3 bucket; you can then specify the Amazon S3 location of the
// file using the Media parameter.
//
// You must include the following parameters in your StartTranscriptionJob request:
//
// - region : The Amazon Web Services Region where you are making your request.
// For a list of Amazon Web Services Regions supported with Amazon Transcribe,
// refer to [Amazon Transcribe endpoints and quotas].
//
// - TranscriptionJobName : A custom name you create for your transcription job
// that is unique within your Amazon Web Services account.
//
// - Media ( MediaFileUri ): The Amazon S3 location of your media file.
//
// - One of LanguageCode , IdentifyLanguage , or IdentifyMultipleLanguages : If
// you know the language of your media file, specify it using the LanguageCode
// parameter; you can find all valid language codes in the [Supported languages]table. If you do not
// know the languages spoken in your media, use either IdentifyLanguage or
// IdentifyMultipleLanguages and let Amazon Transcribe identify the languages for
// you.
//
// [Amazon Transcribe endpoints and quotas]: https://docs.aws.amazon.com/general/latest/gr/transcribe.html
// [Supported languages]: https://docs.aws.amazon.com/transcribe/latest/dg/supported-languages.html
