package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/transcribestreaming"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// transcribestreamingCmd represents the transcribestreaming command
var _transcribestreamingCmd = &cobra.Command{
	Use:   "transcribestreaming",
	Short: "AWS transcribestreaming CLI",
	Run: func(cmd *cobra.Command, args []string) {
		_awsOutput = resolveAWSOutput(_awsProfile, cmd.Flags().Changed("output"))
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := transcribestreaming.NewFromConfig(cfg)
		if _transcribestreamingGetMedicalScribeStream {
			transcribestreaming_GetMedicalScribeStream(cfg, client)
			return
		}
		if _transcribestreamingStartCallAnalyticsStreamTranscription {
			transcribestreaming_StartCallAnalyticsStreamTranscription(cfg, client)
			return
		}
		if _transcribestreamingStartMedicalScribeStream {
			transcribestreaming_StartMedicalScribeStream(cfg, client)
			return
		}
		if _transcribestreamingStartMedicalStreamTranscription {
			transcribestreaming_StartMedicalStreamTranscription(cfg, client)
			return
		}
		if _transcribestreamingStartStreamTranscription {
			transcribestreaming_StartStreamTranscription(cfg, client)
			return
		}

	},
}

var (
	_transcribestreamingGetMedicalScribeStream                bool
	_transcribestreamingStartCallAnalyticsStreamTranscription bool
	_transcribestreamingStartMedicalScribeStream              bool
	_transcribestreamingStartMedicalStreamTranscription       bool
	_transcribestreamingStartStreamTranscription              bool

	_transcribestreamingContentIdentificationType         string
	_transcribestreamingContentRedactionType              string
	_transcribestreamingEnableChannelIdentification       string
	_transcribestreamingEnablePartialResultsStabilization string
	_transcribestreamingIdentifyLanguage                  string
	_transcribestreamingIdentifyMultipleLanguages         string
	_transcribestreamingLanguageCode                      string
	_transcribestreamingLanguageModelName                 string
	_transcribestreamingLanguageOptions                   string
	_transcribestreamingMediaEncoding                     string
	_transcribestreamingMediaSampleRateHertz              string
	_transcribestreamingNumberOfChannels                  string
	_transcribestreamingPartialResultsStability           string
	_transcribestreamingPiiEntityTypes                    string
	_transcribestreamingPreferredLanguage                 string
	_transcribestreamingSessionId                         string
	_transcribestreamingSessionResumeWindow               string
	_transcribestreamingShowSpeakerLabel                  string
	_transcribestreamingSpecialty                         string
	_transcribestreamingType                              string
	_transcribestreamingVocabularyFilterMethod            string
	_transcribestreamingVocabularyFilterName              string
	_transcribestreamingVocabularyFilterNames             string
	_transcribestreamingVocabularyName                    string
	_transcribestreamingVocabularyNames                   string
)

// Provides details about the specified Amazon Web Services HealthScribe streaming
// session. To view the status of the streaming session, check the StreamStatus
// field in the response. To get the details of post-stream analytics, including
// its status, check the PostStreamAnalyticsResult field in the response.
func transcribestreaming_GetMedicalScribeStream(cfg aws.Config, client *transcribestreaming.Client) {
	input := &transcribestreaming.GetMedicalScribeStreamInput{
		// SessionId: *string, // Required
	}

	if len(_transcribestreamingSessionId) > 0 {
		input.SessionId = aws.String(_transcribestreamingSessionId)
	}

	if resp, err := client.GetMedicalScribeStream(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a bidirectional HTTP/2 or WebSocket stream where audio is streamed to
// Amazon Transcribe and the transcription results are streamed to your
// application. Use this operation for [Call Analytics]transcriptions.
//
// The following parameters are required:
//
// - language-code or identify-language
//
// - media-encoding
//
// - sample-rate
//
// For more information on streaming with Amazon Transcribe, see [Transcribing streaming audio].
//
// [Call Analytics]: https://docs.aws.amazon.com/transcribe/latest/dg/call-analytics.html
// [Transcribing streaming audio]: https://docs.aws.amazon.com/transcribe/latest/dg/streaming.html
func transcribestreaming_StartCallAnalyticsStreamTranscription(cfg aws.Config, client *transcribestreaming.Client) {
	input := &transcribestreaming.StartCallAnalyticsStreamTranscriptionInput{
		// MediaEncoding: types.MediaEncoding, // Required
		// MediaSampleRateHertz: *int32, // Required
	}

	if len(_transcribestreamingMediaEncoding) > 0 {
		if err := assignInputField(input, "MediaEncoding", _transcribestreamingMediaEncoding); err != nil {
			log.Errorf("invalid --media-encoding: %s", err.Error())
			return
		}
	}
	if len(_transcribestreamingMediaSampleRateHertz) > 0 {
		if err := assignInputField(input, "MediaSampleRateHertz", _transcribestreamingMediaSampleRateHertz); err != nil {
			log.Errorf("invalid --media-sample-rate-hertz: %s", err.Error())
			return
		}
	}
	if len(_transcribestreamingContentIdentificationType) > 0 {
		if err := assignInputField(input, "ContentIdentificationType", _transcribestreamingContentIdentificationType); err != nil {
			log.Errorf("invalid --content-identification-type: %s", err.Error())
			return
		}
	}
	if len(_transcribestreamingContentRedactionType) > 0 {
		if err := assignInputField(input, "ContentRedactionType", _transcribestreamingContentRedactionType); err != nil {
			log.Errorf("invalid --content-redaction-type: %s", err.Error())
			return
		}
	}
	if len(_transcribestreamingEnablePartialResultsStabilization) > 0 {
		if err := assignInputField(input, "EnablePartialResultsStabilization", _transcribestreamingEnablePartialResultsStabilization); err != nil {
			log.Errorf("invalid --enable-partial-results-stabilization: %s", err.Error())
			return
		}
	}
	if len(_transcribestreamingIdentifyLanguage) > 0 {
		if err := assignInputField(input, "IdentifyLanguage", _transcribestreamingIdentifyLanguage); err != nil {
			log.Errorf("invalid --identify-language: %s", err.Error())
			return
		}
	}
	if len(_transcribestreamingLanguageCode) > 0 {
		if err := assignInputField(input, "LanguageCode", _transcribestreamingLanguageCode); err != nil {
			log.Errorf("invalid --language-code: %s", err.Error())
			return
		}
	}
	if len(_transcribestreamingLanguageModelName) > 0 {
		input.LanguageModelName = aws.String(_transcribestreamingLanguageModelName)
	}
	if len(_transcribestreamingLanguageOptions) > 0 {
		input.LanguageOptions = aws.String(_transcribestreamingLanguageOptions)
	}
	if len(_transcribestreamingPartialResultsStability) > 0 {
		if err := assignInputField(input, "PartialResultsStability", _transcribestreamingPartialResultsStability); err != nil {
			log.Errorf("invalid --partial-results-stability: %s", err.Error())
			return
		}
	}
	if len(_transcribestreamingPiiEntityTypes) > 0 {
		input.PiiEntityTypes = aws.String(_transcribestreamingPiiEntityTypes)
	}
	if len(_transcribestreamingPreferredLanguage) > 0 {
		if err := assignInputField(input, "PreferredLanguage", _transcribestreamingPreferredLanguage); err != nil {
			log.Errorf("invalid --preferred-language: %s", err.Error())
			return
		}
	}
	if len(_transcribestreamingSessionId) > 0 {
		input.SessionId = aws.String(_transcribestreamingSessionId)
	}
	if len(_transcribestreamingVocabularyFilterMethod) > 0 {
		if err := assignInputField(input, "VocabularyFilterMethod", _transcribestreamingVocabularyFilterMethod); err != nil {
			log.Errorf("invalid --vocabulary-filter-method: %s", err.Error())
			return
		}
	}
	if len(_transcribestreamingVocabularyFilterName) > 0 {
		input.VocabularyFilterName = aws.String(_transcribestreamingVocabularyFilterName)
	}
	if len(_transcribestreamingVocabularyFilterNames) > 0 {
		input.VocabularyFilterNames = aws.String(_transcribestreamingVocabularyFilterNames)
	}
	if len(_transcribestreamingVocabularyName) > 0 {
		input.VocabularyName = aws.String(_transcribestreamingVocabularyName)
	}
	if len(_transcribestreamingVocabularyNames) > 0 {
		input.VocabularyNames = aws.String(_transcribestreamingVocabularyNames)
	}

	if resp, err := client.StartCallAnalyticsStreamTranscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

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
func transcribestreaming_StartMedicalScribeStream(cfg aws.Config, client *transcribestreaming.Client) {
	input := &transcribestreaming.StartMedicalScribeStreamInput{
		// LanguageCode: types.MedicalScribeLanguageCode, // Required
		// MediaEncoding: types.MedicalScribeMediaEncoding, // Required
		// MediaSampleRateHertz: *int32, // Required
	}

	if len(_transcribestreamingLanguageCode) > 0 {
		if err := assignInputField(input, "LanguageCode", _transcribestreamingLanguageCode); err != nil {
			log.Errorf("invalid --language-code: %s", err.Error())
			return
		}
	}
	if len(_transcribestreamingMediaEncoding) > 0 {
		if err := assignInputField(input, "MediaEncoding", _transcribestreamingMediaEncoding); err != nil {
			log.Errorf("invalid --media-encoding: %s", err.Error())
			return
		}
	}
	if len(_transcribestreamingMediaSampleRateHertz) > 0 {
		if err := assignInputField(input, "MediaSampleRateHertz", _transcribestreamingMediaSampleRateHertz); err != nil {
			log.Errorf("invalid --media-sample-rate-hertz: %s", err.Error())
			return
		}
	}
	if len(_transcribestreamingSessionId) > 0 {
		input.SessionId = aws.String(_transcribestreamingSessionId)
	}

	if resp, err := client.StartMedicalScribeStream(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a bidirectional HTTP/2 or WebSocket stream where audio is streamed to
// Amazon Transcribe Medical and the transcription results are streamed to your
// application.
//
// The following parameters are required:
//
// - language-code
//
// - media-encoding
//
// - sample-rate
//
// For more information on streaming with Amazon Transcribe Medical, see [Transcribing streaming audio].
//
// [Transcribing streaming audio]: https://docs.aws.amazon.com/transcribe/latest/dg/streaming.html
func transcribestreaming_StartMedicalStreamTranscription(cfg aws.Config, client *transcribestreaming.Client) {
	input := &transcribestreaming.StartMedicalStreamTranscriptionInput{
		// LanguageCode: types.LanguageCode, // Required
		// MediaEncoding: types.MediaEncoding, // Required
		// MediaSampleRateHertz: *int32, // Required
		// Specialty: types.Specialty, // Required
		// Type: types.Type, // Required
	}

	if len(_transcribestreamingLanguageCode) > 0 {
		if err := assignInputField(input, "LanguageCode", _transcribestreamingLanguageCode); err != nil {
			log.Errorf("invalid --language-code: %s", err.Error())
			return
		}
	}
	if len(_transcribestreamingMediaEncoding) > 0 {
		if err := assignInputField(input, "MediaEncoding", _transcribestreamingMediaEncoding); err != nil {
			log.Errorf("invalid --media-encoding: %s", err.Error())
			return
		}
	}
	if len(_transcribestreamingMediaSampleRateHertz) > 0 {
		if err := assignInputField(input, "MediaSampleRateHertz", _transcribestreamingMediaSampleRateHertz); err != nil {
			log.Errorf("invalid --media-sample-rate-hertz: %s", err.Error())
			return
		}
	}
	if len(_transcribestreamingSpecialty) > 0 {
		if err := assignInputField(input, "Specialty", _transcribestreamingSpecialty); err != nil {
			log.Errorf("invalid --specialty: %s", err.Error())
			return
		}
	}
	if len(_transcribestreamingType) > 0 {
		if err := assignInputField(input, "Type", _transcribestreamingType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_transcribestreamingContentIdentificationType) > 0 {
		if err := assignInputField(input, "ContentIdentificationType", _transcribestreamingContentIdentificationType); err != nil {
			log.Errorf("invalid --content-identification-type: %s", err.Error())
			return
		}
	}
	if len(_transcribestreamingEnableChannelIdentification) > 0 {
		if err := assignInputField(input, "EnableChannelIdentification", _transcribestreamingEnableChannelIdentification); err != nil {
			log.Errorf("invalid --enable-channel-identification: %s", err.Error())
			return
		}
	}
	if len(_transcribestreamingNumberOfChannels) > 0 {
		if err := assignInputField(input, "NumberOfChannels", _transcribestreamingNumberOfChannels); err != nil {
			log.Errorf("invalid --number-of-channels: %s", err.Error())
			return
		}
	}
	if len(_transcribestreamingSessionId) > 0 {
		input.SessionId = aws.String(_transcribestreamingSessionId)
	}
	if len(_transcribestreamingShowSpeakerLabel) > 0 {
		if err := assignInputField(input, "ShowSpeakerLabel", _transcribestreamingShowSpeakerLabel); err != nil {
			log.Errorf("invalid --show-speaker-label: %s", err.Error())
			return
		}
	}
	if len(_transcribestreamingVocabularyName) > 0 {
		input.VocabularyName = aws.String(_transcribestreamingVocabularyName)
	}

	if resp, err := client.StartMedicalStreamTranscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a bidirectional HTTP/2 or WebSocket stream where audio is streamed to
// Amazon Transcribe and the transcription results are streamed to your
// application.
//
// The following parameters are required:
//
// - language-code or identify-language or identify-multiple-language
//
// - media-encoding
//
// - sample-rate
//
// For more information on streaming with Amazon Transcribe, see [Transcribing streaming audio].
//
// [Transcribing streaming audio]: https://docs.aws.amazon.com/transcribe/latest/dg/streaming.html
func transcribestreaming_StartStreamTranscription(cfg aws.Config, client *transcribestreaming.Client) {
	input := &transcribestreaming.StartStreamTranscriptionInput{
		// MediaEncoding: types.MediaEncoding, // Required
		// MediaSampleRateHertz: *int32, // Required
	}

	if len(_transcribestreamingMediaEncoding) > 0 {
		if err := assignInputField(input, "MediaEncoding", _transcribestreamingMediaEncoding); err != nil {
			log.Errorf("invalid --media-encoding: %s", err.Error())
			return
		}
	}
	if len(_transcribestreamingMediaSampleRateHertz) > 0 {
		if err := assignInputField(input, "MediaSampleRateHertz", _transcribestreamingMediaSampleRateHertz); err != nil {
			log.Errorf("invalid --media-sample-rate-hertz: %s", err.Error())
			return
		}
	}
	if len(_transcribestreamingContentIdentificationType) > 0 {
		if err := assignInputField(input, "ContentIdentificationType", _transcribestreamingContentIdentificationType); err != nil {
			log.Errorf("invalid --content-identification-type: %s", err.Error())
			return
		}
	}
	if len(_transcribestreamingContentRedactionType) > 0 {
		if err := assignInputField(input, "ContentRedactionType", _transcribestreamingContentRedactionType); err != nil {
			log.Errorf("invalid --content-redaction-type: %s", err.Error())
			return
		}
	}
	if len(_transcribestreamingEnableChannelIdentification) > 0 {
		if err := assignInputField(input, "EnableChannelIdentification", _transcribestreamingEnableChannelIdentification); err != nil {
			log.Errorf("invalid --enable-channel-identification: %s", err.Error())
			return
		}
	}
	if len(_transcribestreamingEnablePartialResultsStabilization) > 0 {
		if err := assignInputField(input, "EnablePartialResultsStabilization", _transcribestreamingEnablePartialResultsStabilization); err != nil {
			log.Errorf("invalid --enable-partial-results-stabilization: %s", err.Error())
			return
		}
	}
	if len(_transcribestreamingIdentifyLanguage) > 0 {
		if err := assignInputField(input, "IdentifyLanguage", _transcribestreamingIdentifyLanguage); err != nil {
			log.Errorf("invalid --identify-language: %s", err.Error())
			return
		}
	}
	if len(_transcribestreamingIdentifyMultipleLanguages) > 0 {
		if err := assignInputField(input, "IdentifyMultipleLanguages", _transcribestreamingIdentifyMultipleLanguages); err != nil {
			log.Errorf("invalid --identify-multiple-languages: %s", err.Error())
			return
		}
	}
	if len(_transcribestreamingLanguageCode) > 0 {
		if err := assignInputField(input, "LanguageCode", _transcribestreamingLanguageCode); err != nil {
			log.Errorf("invalid --language-code: %s", err.Error())
			return
		}
	}
	if len(_transcribestreamingLanguageModelName) > 0 {
		input.LanguageModelName = aws.String(_transcribestreamingLanguageModelName)
	}
	if len(_transcribestreamingLanguageOptions) > 0 {
		input.LanguageOptions = aws.String(_transcribestreamingLanguageOptions)
	}
	if len(_transcribestreamingNumberOfChannels) > 0 {
		if err := assignInputField(input, "NumberOfChannels", _transcribestreamingNumberOfChannels); err != nil {
			log.Errorf("invalid --number-of-channels: %s", err.Error())
			return
		}
	}
	if len(_transcribestreamingPartialResultsStability) > 0 {
		if err := assignInputField(input, "PartialResultsStability", _transcribestreamingPartialResultsStability); err != nil {
			log.Errorf("invalid --partial-results-stability: %s", err.Error())
			return
		}
	}
	if len(_transcribestreamingPiiEntityTypes) > 0 {
		input.PiiEntityTypes = aws.String(_transcribestreamingPiiEntityTypes)
	}
	if len(_transcribestreamingPreferredLanguage) > 0 {
		if err := assignInputField(input, "PreferredLanguage", _transcribestreamingPreferredLanguage); err != nil {
			log.Errorf("invalid --preferred-language: %s", err.Error())
			return
		}
	}
	if len(_transcribestreamingSessionId) > 0 {
		input.SessionId = aws.String(_transcribestreamingSessionId)
	}
	if len(_transcribestreamingSessionResumeWindow) > 0 {
		if err := assignInputField(input, "SessionResumeWindow", _transcribestreamingSessionResumeWindow); err != nil {
			log.Errorf("invalid --session-resume-window: %s", err.Error())
			return
		}
	}
	if len(_transcribestreamingShowSpeakerLabel) > 0 {
		if err := assignInputField(input, "ShowSpeakerLabel", _transcribestreamingShowSpeakerLabel); err != nil {
			log.Errorf("invalid --show-speaker-label: %s", err.Error())
			return
		}
	}
	if len(_transcribestreamingVocabularyFilterMethod) > 0 {
		if err := assignInputField(input, "VocabularyFilterMethod", _transcribestreamingVocabularyFilterMethod); err != nil {
			log.Errorf("invalid --vocabulary-filter-method: %s", err.Error())
			return
		}
	}
	if len(_transcribestreamingVocabularyFilterName) > 0 {
		input.VocabularyFilterName = aws.String(_transcribestreamingVocabularyFilterName)
	}
	if len(_transcribestreamingVocabularyFilterNames) > 0 {
		input.VocabularyFilterNames = aws.String(_transcribestreamingVocabularyFilterNames)
	}
	if len(_transcribestreamingVocabularyName) > 0 {
		input.VocabularyName = aws.String(_transcribestreamingVocabularyName)
	}
	if len(_transcribestreamingVocabularyNames) > 0 {
		input.VocabularyNames = aws.String(_transcribestreamingVocabularyNames)
	}

	if resp, err := client.StartStreamTranscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_transcribestreamingCmd)
	_transcribestreamingCmd.Flags().SortFlags = false

	_transcribestreamingCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_transcribestreamingCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_transcribestreamingCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_transcribestreamingCmd.Flags().StringVarP(&_transcribestreamingContentIdentificationType, "content-identification-type", "", "", "Content Identification Type")
	_transcribestreamingCmd.Flags().StringVarP(&_transcribestreamingContentRedactionType, "content-redaction-type", "", "", "Content Redaction Type")
	_transcribestreamingCmd.Flags().StringVarP(&_transcribestreamingEnableChannelIdentification, "enable-channel-identification", "", "", "Enable Channel Identification")
	_transcribestreamingCmd.Flags().StringVarP(&_transcribestreamingEnablePartialResultsStabilization, "enable-partial-results-stabilization", "", "", "Enable Partial Results Stabilization")
	_transcribestreamingCmd.Flags().StringVarP(&_transcribestreamingIdentifyLanguage, "identify-language", "", "", "Identify Language")
	_transcribestreamingCmd.Flags().StringVarP(&_transcribestreamingIdentifyMultipleLanguages, "identify-multiple-languages", "", "", "Identify Multiple Languages")
	_transcribestreamingCmd.Flags().StringVarP(&_transcribestreamingLanguageCode, "language-code", "", "", "Language Code")
	_transcribestreamingCmd.Flags().StringVarP(&_transcribestreamingLanguageModelName, "language-model-name", "", "", "Language Model Name")
	_transcribestreamingCmd.Flags().StringVarP(&_transcribestreamingLanguageOptions, "language-options", "", "", "Language Options")
	_transcribestreamingCmd.Flags().StringVarP(&_transcribestreamingMediaEncoding, "media-encoding", "", "", "Media Encoding")
	_transcribestreamingCmd.Flags().StringVarP(&_transcribestreamingMediaSampleRateHertz, "media-sample-rate-hertz", "", "", "Media Sample Rate Hertz")
	_transcribestreamingCmd.Flags().StringVarP(&_transcribestreamingNumberOfChannels, "number-of-channels", "", "", "Number Of Channels")
	_transcribestreamingCmd.Flags().StringVarP(&_transcribestreamingPartialResultsStability, "partial-results-stability", "", "", "Partial Results Stability")
	_transcribestreamingCmd.Flags().StringVarP(&_transcribestreamingPiiEntityTypes, "pii-entity-types", "", "", "Pii Entity Types")
	_transcribestreamingCmd.Flags().StringVarP(&_transcribestreamingPreferredLanguage, "preferred-language", "", "", "Preferred Language")
	_transcribestreamingCmd.Flags().StringVarP(&_transcribestreamingSessionId, "session-id", "", "", "Session ID")
	_transcribestreamingCmd.Flags().StringVarP(&_transcribestreamingSessionResumeWindow, "session-resume-window", "", "", "Session Resume Window")
	_transcribestreamingCmd.Flags().StringVarP(&_transcribestreamingShowSpeakerLabel, "show-speaker-label", "", "", "Show Speaker Label")
	_transcribestreamingCmd.Flags().StringVarP(&_transcribestreamingSpecialty, "specialty", "", "", "Specialty")
	_transcribestreamingCmd.Flags().StringVarP(&_transcribestreamingType, "type", "", "", "Type")
	_transcribestreamingCmd.Flags().StringVarP(&_transcribestreamingVocabularyFilterMethod, "vocabulary-filter-method", "", "", "Vocabulary Filter Method")
	_transcribestreamingCmd.Flags().StringVarP(&_transcribestreamingVocabularyFilterName, "vocabulary-filter-name", "", "", "Vocabulary Filter Name")
	_transcribestreamingCmd.Flags().StringVarP(&_transcribestreamingVocabularyFilterNames, "vocabulary-filter-names", "", "", "Vocabulary Filter Names")
	_transcribestreamingCmd.Flags().StringVarP(&_transcribestreamingVocabularyName, "vocabulary-name", "", "", "Vocabulary Name")
	_transcribestreamingCmd.Flags().StringVarP(&_transcribestreamingVocabularyNames, "vocabulary-names", "", "", "Vocabulary Names")

	_transcribestreamingCmd.Flags().BoolVarP(&_transcribestreamingGetMedicalScribeStream, "get-medical-scribe-stream", "", false, "Get Medical Scribe Stream")
	_transcribestreamingCmd.Flags().BoolVarP(&_transcribestreamingStartCallAnalyticsStreamTranscription, "start-call-analytics-stream-transcription", "", false, "Start Call Analytics Stream Transcription")
	_transcribestreamingCmd.Flags().BoolVarP(&_transcribestreamingStartMedicalScribeStream, "start-medical-scribe-stream", "", false, "Start Medical Scribe Stream")
	_transcribestreamingCmd.Flags().BoolVarP(&_transcribestreamingStartMedicalStreamTranscription, "start-medical-stream-transcription", "", false, "Start Medical Stream Transcription")
	_transcribestreamingCmd.Flags().BoolVarP(&_transcribestreamingStartStreamTranscription, "start-stream-transcription", "", false, "Start Stream Transcription")

}
