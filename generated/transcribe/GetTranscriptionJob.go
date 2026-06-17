package transcribe

// GetTranscriptionJob is generated as a reference stub.
// Executable command wiring lives under cmd/transcribe.go.
//
// Provides information about the specified transcription job.
//
// To view the status of the specified transcription job, check the
// TranscriptionJobStatus field. If the status is COMPLETED , the job is finished.
// You can find the results at the location specified in TranscriptFileUri . If the
// status is FAILED , FailureReason provides details on why your transcription job
// failed.
//
// If you enabled content redaction, the redacted transcript can be found at the
// location specified in RedactedTranscriptFileUri .
//
// To get a list of your transcription jobs, use the operation.
