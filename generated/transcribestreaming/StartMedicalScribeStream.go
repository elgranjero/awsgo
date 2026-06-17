package transcribestreaming

// StartMedicalScribeStream is generated as a reference stub.
// Executable command wiring lives under cmd/transcribestreaming.go.
//
// Starts a bidirectional HTTP/2 stream, where audio is streamed to Amazon Web
// Services HealthScribe and the transcription results are streamed to your
// application.
//
// When you start a stream, you first specify the stream configuration in a
// MedicalScribeConfigurationEvent . This event includes channel definitions,
// encryption settings, medical scribe context, and post-stream analytics settings,
// such as the output configuration for aggregated transcript and clinical note
// generation. These are additional streaming session configurations beyond those
// provided in your initial start request headers. Whether you are starting a new
// session or resuming an existing session, your first event must be a
// MedicalScribeConfigurationEvent .
//
// After you send a MedicalScribeConfigurationEvent , you start AudioEvents and
// Amazon Web Services HealthScribe responds with real-time transcription results.
// When you are finished, to start processing the results with the post-stream
// analytics, send a MedicalScribeSessionControlEvent with a Type of END_OF_SESSION
// and Amazon Web Services HealthScribe starts the analytics.
//
// You can pause or resume streaming. To pause streaming, complete the input
// stream without sending the MedicalScribeSessionControlEvent . To resume
// streaming, call the StartMedicalScribeStream and specify the same SessionId you
// used to start the stream.
//
// The following parameters are required:
//
// - language-code
//
// - media-encoding
//
// - media-sample-rate-hertz
//
// For more information on streaming with Amazon Web Services HealthScribe, see [Amazon Web Services HealthScribe].
//
// [Amazon Web Services HealthScribe]: https://docs.aws.amazon.com/transcribe/latest/dg/health-scribe-streaming.html
