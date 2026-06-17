package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/transcribestreaming/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"get-medical-scribe-stream", "start-call-analytics-stream-transcription", "start-medical-scribe-stream", "start-medical-stream-transcription", "start-stream-transcription"},
		OperationSet: map[string]bool{"get-medical-scribe-stream": true, "start-call-analytics-stream-transcription": true, "start-medical-scribe-stream": true, "start-medical-stream-transcription": true, "start-stream-transcription": true},
		OperationInputs: map[string][]string{
			"get-medical-scribe-stream":                 {"SessionId"},
			"start-call-analytics-stream-transcription": {"ContentIdentificationType", "ContentRedactionType", "EnablePartialResultsStabilization", "IdentifyLanguage", "LanguageCode", "LanguageModelName", "LanguageOptions", "MediaEncoding", "MediaSampleRateHertz", "PartialResultsStability", "PiiEntityTypes", "PreferredLanguage", "SessionId", "VocabularyFilterMethod", "VocabularyFilterName", "VocabularyFilterNames", "VocabularyName", "VocabularyNames"},
			"start-medical-scribe-stream":               {"LanguageCode", "MediaEncoding", "MediaSampleRateHertz", "SessionId"},
			"start-medical-stream-transcription":        {"ContentIdentificationType", "EnableChannelIdentification", "LanguageCode", "MediaEncoding", "MediaSampleRateHertz", "NumberOfChannels", "SessionId", "ShowSpeakerLabel", "Specialty", "Type", "VocabularyName"},
			"start-stream-transcription":                {"ContentIdentificationType", "ContentRedactionType", "EnableChannelIdentification", "EnablePartialResultsStabilization", "IdentifyLanguage", "IdentifyMultipleLanguages", "LanguageCode", "LanguageModelName", "LanguageOptions", "MediaEncoding", "MediaSampleRateHertz", "NumberOfChannels", "PartialResultsStability", "PiiEntityTypes", "PreferredLanguage", "SessionId", "SessionResumeWindow", "ShowSpeakerLabel", "VocabularyFilterMethod", "VocabularyFilterName", "VocabularyFilterNames", "VocabularyName", "VocabularyNames"},
		},
		OperationInputTypes: map[string]map[string]string{
			"get-medical-scribe-stream":                 {"SessionId": "*string"},
			"start-call-analytics-stream-transcription": {"ContentIdentificationType": "types.ContentIdentificationType", "ContentRedactionType": "types.ContentRedactionType", "EnablePartialResultsStabilization": "bool", "IdentifyLanguage": "bool", "LanguageCode": "types.CallAnalyticsLanguageCode", "LanguageModelName": "*string", "LanguageOptions": "*string", "MediaEncoding": "types.MediaEncoding", "MediaSampleRateHertz": "*int32", "PartialResultsStability": "types.PartialResultsStability", "PiiEntityTypes": "*string", "PreferredLanguage": "types.CallAnalyticsLanguageCode", "SessionId": "*string", "VocabularyFilterMethod": "types.VocabularyFilterMethod", "VocabularyFilterName": "*string", "VocabularyFilterNames": "*string", "VocabularyName": "*string", "VocabularyNames": "*string"},
			"start-medical-scribe-stream":               {"LanguageCode": "types.MedicalScribeLanguageCode", "MediaEncoding": "types.MedicalScribeMediaEncoding", "MediaSampleRateHertz": "*int32", "SessionId": "*string"},
			"start-medical-stream-transcription":        {"ContentIdentificationType": "types.MedicalContentIdentificationType", "EnableChannelIdentification": "bool", "LanguageCode": "types.LanguageCode", "MediaEncoding": "types.MediaEncoding", "MediaSampleRateHertz": "*int32", "NumberOfChannels": "*int32", "SessionId": "*string", "ShowSpeakerLabel": "bool", "Specialty": "types.Specialty", "Type": "types.Type", "VocabularyName": "*string"},
			"start-stream-transcription":                {"ContentIdentificationType": "types.ContentIdentificationType", "ContentRedactionType": "types.ContentRedactionType", "EnableChannelIdentification": "bool", "EnablePartialResultsStabilization": "bool", "IdentifyLanguage": "bool", "IdentifyMultipleLanguages": "bool", "LanguageCode": "types.LanguageCode", "LanguageModelName": "*string", "LanguageOptions": "*string", "MediaEncoding": "types.MediaEncoding", "MediaSampleRateHertz": "*int32", "NumberOfChannels": "*int32", "PartialResultsStability": "types.PartialResultsStability", "PiiEntityTypes": "*string", "PreferredLanguage": "types.LanguageCode", "SessionId": "*string", "SessionResumeWindow": "*int32", "ShowSpeakerLabel": "bool", "VocabularyFilterMethod": "types.VocabularyFilterMethod", "VocabularyFilterName": "*string", "VocabularyFilterNames": "*string", "VocabularyName": "*string", "VocabularyNames": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"get-medical-scribe-stream":                 {"SessionId"},
			"start-call-analytics-stream-transcription": {"MediaEncoding", "MediaSampleRateHertz"},
			"start-medical-scribe-stream":               {"LanguageCode", "MediaEncoding", "MediaSampleRateHertz"},
			"start-medical-stream-transcription":        {"LanguageCode", "MediaEncoding", "MediaSampleRateHertz", "Specialty", "Type"},
			"start-stream-transcription":                {"MediaEncoding", "MediaSampleRateHertz"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("transcribestreaming", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
