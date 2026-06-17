package transcribe

// StartMedicalScribeJob is generated as a reference stub.
// Executable command wiring lives under cmd/transcribe.go.
//
// Transcribes patient-clinician conversations and generates clinical notes.
//
// Amazon Web Services HealthScribe automatically provides rich conversation
// transcripts, identifies speaker roles, classifies dialogues, extracts medical
// terms, and generates preliminary clinical notes. To learn more about these
// features, refer to [Amazon Web Services HealthScribe].
//
// To make a StartMedicalScribeJob request, you must first upload your media file
// into an Amazon S3 bucket; you can then specify the Amazon S3 location of the
// file using the Media parameter.
//
// You must include the following parameters in your StartMedicalTranscriptionJob
// request:
//
// - DataAccessRoleArn : The ARN of an IAM role with the these minimum
// permissions: read permission on input file Amazon S3 bucket specified in Media
// , write permission on the Amazon S3 bucket specified in OutputBucketName , and
// full permissions on the KMS key specified in OutputEncryptionKMSKeyId (if
// set). The role should also allow transcribe.amazonaws.com to assume it.
//
// - Media ( MediaFileUri ): The Amazon S3 location of your media file.
//
// - MedicalScribeJobName : A custom name you create for your MedicalScribe job
// that is unique within your Amazon Web Services account.
//
// - OutputBucketName : The Amazon S3 bucket where you want your output files
// stored.
//
// - Settings : A MedicalScribeSettings object that must set exactly one of
// ShowSpeakerLabels or ChannelIdentification to true. If ShowSpeakerLabels is
// true, MaxSpeakerLabels must also be set.
//
// - ChannelDefinitions : A MedicalScribeChannelDefinitions array should be set
// if and only if the ChannelIdentification value of Settings is set to true.
//
// [Amazon Web Services HealthScribe]: https://docs.aws.amazon.com/transcribe/latest/dg/health-scribe.html
