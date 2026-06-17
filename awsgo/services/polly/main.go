package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/polly/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"delete-lexicon", "describe-voices", "get-lexicon", "get-speech-synthesis-task", "list-lexicons", "list-speech-synthesis-tasks", "put-lexicon", "start-speech-synthesis-task", "synthesize-speech"},
		OperationSet: map[string]bool{"delete-lexicon": true, "describe-voices": true, "get-lexicon": true, "get-speech-synthesis-task": true, "list-lexicons": true, "list-speech-synthesis-tasks": true, "put-lexicon": true, "start-speech-synthesis-task": true, "synthesize-speech": true},
		OperationInputs: map[string][]string{
			"delete-lexicon":              {"Name"},
			"describe-voices":             {"Engine", "IncludeAdditionalLanguageCodes", "LanguageCode", "NextToken"},
			"get-lexicon":                 {"Name"},
			"get-speech-synthesis-task":   {"TaskId"},
			"list-lexicons":               {"NextToken"},
			"list-speech-synthesis-tasks": {"MaxResults", "NextToken", "Status"},
			"put-lexicon":                 {"Content", "Name"},
			"start-speech-synthesis-task": {"Engine", "LanguageCode", "LexiconNames", "OutputFormat", "OutputS3BucketName", "OutputS3KeyPrefix", "SampleRate", "SnsTopicArn", "SpeechMarkTypes", "Text", "TextType", "VoiceId"},
			"synthesize-speech":           {"Engine", "LanguageCode", "LexiconNames", "OutputFormat", "SampleRate", "SpeechMarkTypes", "Text", "TextType", "VoiceId"},
		},
		OperationInputTypes: map[string]map[string]string{
			"delete-lexicon":              {"Name": "*string"},
			"describe-voices":             {"Engine": "types.Engine", "IncludeAdditionalLanguageCodes": "bool", "LanguageCode": "types.LanguageCode", "NextToken": "*string"},
			"get-lexicon":                 {"Name": "*string"},
			"get-speech-synthesis-task":   {"TaskId": "*string"},
			"list-lexicons":               {"NextToken": "*string"},
			"list-speech-synthesis-tasks": {"MaxResults": "*int32", "NextToken": "*string", "Status": "types.TaskStatus"},
			"put-lexicon":                 {"Content": "*string", "Name": "*string"},
			"start-speech-synthesis-task": {"Engine": "types.Engine", "LanguageCode": "types.LanguageCode", "LexiconNames": "[]string", "OutputFormat": "types.OutputFormat", "OutputS3BucketName": "*string", "OutputS3KeyPrefix": "*string", "SampleRate": "*string", "SnsTopicArn": "*string", "SpeechMarkTypes": "[]types.SpeechMarkType", "Text": "*string", "TextType": "types.TextType", "VoiceId": "types.VoiceId"},
			"synthesize-speech":           {"Engine": "types.Engine", "LanguageCode": "types.LanguageCode", "LexiconNames": "[]string", "OutputFormat": "types.OutputFormat", "SampleRate": "*string", "SpeechMarkTypes": "[]types.SpeechMarkType", "Text": "*string", "TextType": "types.TextType", "VoiceId": "types.VoiceId"},
		},
		OperationInputRequired: map[string][]string{
			"delete-lexicon":              {"Name"},
			"describe-voices":             {},
			"get-lexicon":                 {"Name"},
			"get-speech-synthesis-task":   {"TaskId"},
			"list-lexicons":               {},
			"list-speech-synthesis-tasks": {},
			"put-lexicon":                 {"Content", "Name"},
			"start-speech-synthesis-task": {"OutputFormat", "OutputS3BucketName", "Text", "VoiceId"},
			"synthesize-speech":           {"OutputFormat", "Text", "VoiceId"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("polly", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
