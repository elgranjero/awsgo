package transcribe

// StartMedicalTranscriptionJob is generated as a reference stub.
// Executable command wiring lives under cmd/transcribe.go.
//
// Transcribes the audio from a medical dictation or conversation and applies any
// additional Request Parameters you choose to include in your request.
//
// In addition to many standard transcription features, Amazon Transcribe Medical
// provides you with a robust medical vocabulary and, optionally, content
// identification, which adds flags to personal health information (PHI). To learn
// more about these features, refer to [How Amazon Transcribe Medical works].
//
// To make a StartMedicalTranscriptionJob request, you must first upload your
// media file into an Amazon S3 bucket; you can then specify the Amazon S3 location
// of the file using the Media parameter.
//
// You must include the following parameters in your StartMedicalTranscriptionJob
// request:
//
// - region : The Amazon Web Services Region where you are making your request.
// For a list of Amazon Web Services Regions supported with Amazon Transcribe,
// refer to [Amazon Transcribe endpoints and quotas].
//
// - MedicalTranscriptionJobName : A custom name you create for your
// transcription job that is unique within your Amazon Web Services account.
//
// - Media ( MediaFileUri ): The Amazon S3 location of your media file.
//
// - LanguageCode : This must be en-US .
//
// - OutputBucketName : The Amazon S3 bucket where you want your transcript
// stored. If you want your output stored in a sub-folder of this bucket, you must
// also include OutputKey .
//
// - Specialty : This must be PRIMARYCARE .
//
// - Type : Choose whether your audio is a conversation or a dictation.
//
// [How Amazon Transcribe Medical works]: https://docs.aws.amazon.com/transcribe/latest/dg/how-it-works-med.html
// [Amazon Transcribe endpoints and quotas]: https://docs.aws.amazon.com/general/latest/gr/transcribe.html
