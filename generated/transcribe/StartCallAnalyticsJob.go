package transcribe

// StartCallAnalyticsJob is generated as a reference stub.
// Executable command wiring lives under cmd/transcribe.go.
//
// Transcribes the audio from a customer service call and applies any additional
// Request Parameters you choose to include in your request.
//
// In addition to many standard transcription features, Call Analytics provides
// you with call characteristics, call summarization, speaker sentiment, and
// optional redaction of your text transcript and your audio file. You can also
// apply custom categories to flag specified conditions. To learn more about these
// features and insights, refer to [Analyzing call center audio with Call Analytics].
//
// If you want to apply categories to your Call Analytics job, you must create
// them before submitting your job request. Categories cannot be retroactively
// applied to a job. To create a new category, use the operation. To learn more
// about Call Analytics categories, see [Creating categories for post-call transcriptions]and [Creating categories for real-time transcriptions].
//
// To make a StartCallAnalyticsJob request, you must first upload your media file
// into an Amazon S3 bucket; you can then specify the Amazon S3 location of the
// file using the Media parameter.
//
// Job queuing is available for Call Analytics jobs. If you pass a
// DataAccessRoleArn in your request and you exceed your Concurrent Job Limit, your
// job will automatically be added to a queue to be processed once your concurrent
// job count is below the limit.
//
// You must include the following parameters in your StartCallAnalyticsJob request:
//
// - region : The Amazon Web Services Region where you are making your request.
// For a list of Amazon Web Services Regions supported with Amazon Transcribe,
// refer to [Amazon Transcribe endpoints and quotas].
//
// - CallAnalyticsJobName : A custom name that you create for your transcription
// job that's unique within your Amazon Web Services account.
//
// - Media ( MediaFileUri or RedactedMediaFileUri ): The Amazon S3 location of
// your media file.
//
// With Call Analytics, you can redact the audio contained in your media file by
// including RedactedMediaFileUri , instead of MediaFileUri , to specify the
// location of your input audio. If you choose to redact your audio, you can find
// your redacted media at the location specified in the RedactedMediaFileUri field
// of your response.
//
// [Amazon Transcribe endpoints and quotas]: https://docs.aws.amazon.com/general/latest/gr/transcribe.html
// [Analyzing call center audio with Call Analytics]: https://docs.aws.amazon.com/transcribe/latest/dg/call-analytics.html
// [Creating categories for post-call transcriptions]: https://docs.aws.amazon.com/transcribe/latest/dg/tca-categories-batch.html
// [Creating categories for real-time transcriptions]: https://docs.aws.amazon.com/transcribe/latest/dg/tca-categories-stream.html
