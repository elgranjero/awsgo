package transcribe

// GetCallAnalyticsJob is generated as a reference stub.
// Executable command wiring lives under cmd/transcribe.go.
//
// Provides information about the specified Call Analytics job.
//
// To view the job's status, refer to CallAnalyticsJobStatus . If the status is
// COMPLETED , the job is finished. You can find your completed transcript at the
// URI specified in TranscriptFileUri . If the status is FAILED , FailureReason
// provides details on why your transcription job failed.
//
// If you enabled personally identifiable information (PII) redaction, the
// redacted transcript appears at the location specified in
// RedactedTranscriptFileUri .
//
// If you chose to redact the audio in your media file, you can find your redacted
// media file at the location specified in RedactedMediaFileUri .
//
// To get a list of your Call Analytics jobs, use the operation.
