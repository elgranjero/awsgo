package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/transcribestreaming"
)

var fields_get_medical_scribe_stream = []leanruntime.Field{
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: true},
}

var fields_start_call_analytics_stream_transcription = []leanruntime.Field{
	{Name: "ContentIdentificationType", Flag: "content-identification-type", Type: "types.ContentIdentificationType", Required: false},
	{Name: "ContentRedactionType", Flag: "content-redaction-type", Type: "types.ContentRedactionType", Required: false},
	{Name: "EnablePartialResultsStabilization", Flag: "enable-partial-results-stabilization", Type: "bool", Required: false},
	{Name: "IdentifyLanguage", Flag: "identify-language", Type: "bool", Required: false},
	{Name: "LanguageCode", Flag: "language-code", Type: "types.CallAnalyticsLanguageCode", Required: false},
	{Name: "LanguageModelName", Flag: "language-model-name", Type: "*string", Required: false},
	{Name: "LanguageOptions", Flag: "language-options", Type: "*string", Required: false},
	{Name: "MediaEncoding", Flag: "media-encoding", Type: "types.MediaEncoding", Required: true},
	{Name: "MediaSampleRateHertz", Flag: "media-sample-rate-hertz", Type: "*int32", Required: true},
	{Name: "PartialResultsStability", Flag: "partial-results-stability", Type: "types.PartialResultsStability", Required: false},
	{Name: "PiiEntityTypes", Flag: "pii-entity-types", Type: "*string", Required: false},
	{Name: "PreferredLanguage", Flag: "preferred-language", Type: "types.CallAnalyticsLanguageCode", Required: false},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: false},
	{Name: "VocabularyFilterMethod", Flag: "vocabulary-filter-method", Type: "types.VocabularyFilterMethod", Required: false},
	{Name: "VocabularyFilterName", Flag: "vocabulary-filter-name", Type: "*string", Required: false},
	{Name: "VocabularyFilterNames", Flag: "vocabulary-filter-names", Type: "*string", Required: false},
	{Name: "VocabularyName", Flag: "vocabulary-name", Type: "*string", Required: false},
	{Name: "VocabularyNames", Flag: "vocabulary-names", Type: "*string", Required: false},
}

var fields_start_medical_scribe_stream = []leanruntime.Field{
	{Name: "LanguageCode", Flag: "language-code", Type: "types.MedicalScribeLanguageCode", Required: true},
	{Name: "MediaEncoding", Flag: "media-encoding", Type: "types.MedicalScribeMediaEncoding", Required: true},
	{Name: "MediaSampleRateHertz", Flag: "media-sample-rate-hertz", Type: "*int32", Required: true},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: false},
}

var fields_start_medical_stream_transcription = []leanruntime.Field{
	{Name: "ContentIdentificationType", Flag: "content-identification-type", Type: "types.MedicalContentIdentificationType", Required: false},
	{Name: "EnableChannelIdentification", Flag: "enable-channel-identification", Type: "bool", Required: false},
	{Name: "LanguageCode", Flag: "language-code", Type: "types.LanguageCode", Required: true},
	{Name: "MediaEncoding", Flag: "media-encoding", Type: "types.MediaEncoding", Required: true},
	{Name: "MediaSampleRateHertz", Flag: "media-sample-rate-hertz", Type: "*int32", Required: true},
	{Name: "NumberOfChannels", Flag: "number-of-channels", Type: "*int32", Required: false},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: false},
	{Name: "ShowSpeakerLabel", Flag: "show-speaker-label", Type: "bool", Required: false},
	{Name: "Specialty", Flag: "specialty", Type: "types.Specialty", Required: true},
	{Name: "Type", Flag: "type", Type: "types.Type", Required: true},
	{Name: "VocabularyName", Flag: "vocabulary-name", Type: "*string", Required: false},
}

var fields_start_stream_transcription = []leanruntime.Field{
	{Name: "ContentIdentificationType", Flag: "content-identification-type", Type: "types.ContentIdentificationType", Required: false},
	{Name: "ContentRedactionType", Flag: "content-redaction-type", Type: "types.ContentRedactionType", Required: false},
	{Name: "EnableChannelIdentification", Flag: "enable-channel-identification", Type: "bool", Required: false},
	{Name: "EnablePartialResultsStabilization", Flag: "enable-partial-results-stabilization", Type: "bool", Required: false},
	{Name: "IdentifyLanguage", Flag: "identify-language", Type: "bool", Required: false},
	{Name: "IdentifyMultipleLanguages", Flag: "identify-multiple-languages", Type: "bool", Required: false},
	{Name: "LanguageCode", Flag: "language-code", Type: "types.LanguageCode", Required: false},
	{Name: "LanguageModelName", Flag: "language-model-name", Type: "*string", Required: false},
	{Name: "LanguageOptions", Flag: "language-options", Type: "*string", Required: false},
	{Name: "MediaEncoding", Flag: "media-encoding", Type: "types.MediaEncoding", Required: true},
	{Name: "MediaSampleRateHertz", Flag: "media-sample-rate-hertz", Type: "*int32", Required: true},
	{Name: "NumberOfChannels", Flag: "number-of-channels", Type: "*int32", Required: false},
	{Name: "PartialResultsStability", Flag: "partial-results-stability", Type: "types.PartialResultsStability", Required: false},
	{Name: "PiiEntityTypes", Flag: "pii-entity-types", Type: "*string", Required: false},
	{Name: "PreferredLanguage", Flag: "preferred-language", Type: "types.LanguageCode", Required: false},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: false},
	{Name: "SessionResumeWindow", Flag: "session-resume-window", Type: "*int32", Required: false},
	{Name: "ShowSpeakerLabel", Flag: "show-speaker-label", Type: "bool", Required: false},
	{Name: "VocabularyFilterMethod", Flag: "vocabulary-filter-method", Type: "types.VocabularyFilterMethod", Required: false},
	{Name: "VocabularyFilterName", Flag: "vocabulary-filter-name", Type: "*string", Required: false},
	{Name: "VocabularyFilterNames", Flag: "vocabulary-filter-names", Type: "*string", Required: false},
	{Name: "VocabularyName", Flag: "vocabulary-name", Type: "*string", Required: false},
	{Name: "VocabularyNames", Flag: "vocabulary-names", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"get-medical-scribe-stream": {
			Name:   "get-medical-scribe-stream",
			Fields: fields_get_medical_scribe_stream,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMedicalScribeStreamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_medical_scribe_stream, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMedicalScribeStream(ctx, input)
			},
		},
		"start-call-analytics-stream-transcription": {
			Name:   "start-call-analytics-stream-transcription",
			Fields: fields_start_call_analytics_stream_transcription,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartCallAnalyticsStreamTranscriptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_call_analytics_stream_transcription, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartCallAnalyticsStreamTranscription(ctx, input)
			},
		},
		"start-medical-scribe-stream": {
			Name:   "start-medical-scribe-stream",
			Fields: fields_start_medical_scribe_stream,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartMedicalScribeStreamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_medical_scribe_stream, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartMedicalScribeStream(ctx, input)
			},
		},
		"start-medical-stream-transcription": {
			Name:   "start-medical-stream-transcription",
			Fields: fields_start_medical_stream_transcription,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartMedicalStreamTranscriptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_medical_stream_transcription, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartMedicalStreamTranscription(ctx, input)
			},
		},
		"start-stream-transcription": {
			Name:   "start-stream-transcription",
			Fields: fields_start_stream_transcription,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartStreamTranscriptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_stream_transcription, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartStreamTranscription(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("transcribestreaming", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
