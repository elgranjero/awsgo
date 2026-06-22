package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/polly"
)

var fields_delete_lexicon = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_describe_voices = []leanruntime.Field{
	{Name: "Engine", Flag: "engine", Type: "types.Engine", Required: false},
	{Name: "IncludeAdditionalLanguageCodes", Flag: "include-additional-language-codes", Type: "bool", Required: false},
	{Name: "LanguageCode", Flag: "language-code", Type: "types.LanguageCode", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_lexicon = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_speech_synthesis_task = []leanruntime.Field{
	{Name: "TaskId", Flag: "task-id", Type: "*string", Required: true},
}

var fields_list_lexicons = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_speech_synthesis_tasks = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.TaskStatus", Required: false},
}

var fields_put_lexicon = []leanruntime.Field{
	{Name: "Content", Flag: "content", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_start_speech_synthesis_task = []leanruntime.Field{
	{Name: "Engine", Flag: "engine", Type: "types.Engine", Required: false},
	{Name: "LanguageCode", Flag: "language-code", Type: "types.LanguageCode", Required: false},
	{Name: "LexiconNames", Flag: "lexicon-names", Type: "[]string", Required: false},
	{Name: "OutputFormat", Flag: "output-format", Type: "types.OutputFormat", Required: true},
	{Name: "OutputS3BucketName", Flag: "output-s3-bucket-name", Type: "*string", Required: true},
	{Name: "OutputS3KeyPrefix", Flag: "output-s3-key-prefix", Type: "*string", Required: false},
	{Name: "SampleRate", Flag: "sample-rate", Type: "*string", Required: false},
	{Name: "SnsTopicArn", Flag: "sns-topic-arn", Type: "*string", Required: false},
	{Name: "SpeechMarkTypes", Flag: "speech-mark-types", Type: "[]types.SpeechMarkType", Required: false},
	{Name: "Text", Flag: "text", Type: "*string", Required: true},
	{Name: "TextType", Flag: "text-type", Type: "types.TextType", Required: false},
	{Name: "VoiceId", Flag: "voice-id", Type: "types.VoiceId", Required: true},
}

var fields_synthesize_speech = []leanruntime.Field{
	{Name: "Engine", Flag: "engine", Type: "types.Engine", Required: false},
	{Name: "LanguageCode", Flag: "language-code", Type: "types.LanguageCode", Required: false},
	{Name: "LexiconNames", Flag: "lexicon-names", Type: "[]string", Required: false},
	{Name: "OutputFormat", Flag: "output-format", Type: "types.OutputFormat", Required: true},
	{Name: "SampleRate", Flag: "sample-rate", Type: "*string", Required: false},
	{Name: "SpeechMarkTypes", Flag: "speech-mark-types", Type: "[]types.SpeechMarkType", Required: false},
	{Name: "Text", Flag: "text", Type: "*string", Required: true},
	{Name: "TextType", Flag: "text-type", Type: "types.TextType", Required: false},
	{Name: "VoiceId", Flag: "voice-id", Type: "types.VoiceId", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"delete-lexicon": {
			Name:   "delete-lexicon",
			Fields: fields_delete_lexicon,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteLexiconInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_lexicon, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteLexicon(ctx, input)
			},
		},
		"describe-voices": {
			Name:   "describe-voices",
			Fields: fields_describe_voices,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeVoicesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_voices, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeVoices(ctx, input)
			},
		},
		"get-lexicon": {
			Name:   "get-lexicon",
			Fields: fields_get_lexicon,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLexiconInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_lexicon, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetLexicon(ctx, input)
			},
		},
		"get-speech-synthesis-task": {
			Name:   "get-speech-synthesis-task",
			Fields: fields_get_speech_synthesis_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSpeechSynthesisTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_speech_synthesis_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSpeechSynthesisTask(ctx, input)
			},
		},
		"list-lexicons": {
			Name:   "list-lexicons",
			Fields: fields_list_lexicons,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLexiconsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_lexicons, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListLexicons(ctx, input)
			},
		},
		"list-speech-synthesis-tasks": {
			Name:   "list-speech-synthesis-tasks",
			Fields: fields_list_speech_synthesis_tasks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSpeechSynthesisTasksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_speech_synthesis_tasks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSpeechSynthesisTasks(ctx, input)
				}
				var results []*svc.ListSpeechSynthesisTasksOutput
				p := svc.NewListSpeechSynthesisTasksPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"put-lexicon": {
			Name:   "put-lexicon",
			Fields: fields_put_lexicon,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutLexiconInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_lexicon, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutLexicon(ctx, input)
			},
		},
		"start-speech-synthesis-task": {
			Name:   "start-speech-synthesis-task",
			Fields: fields_start_speech_synthesis_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartSpeechSynthesisTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_speech_synthesis_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartSpeechSynthesisTask(ctx, input)
			},
		},
		"synthesize-speech": {
			Name:   "synthesize-speech",
			Fields: fields_synthesize_speech,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SynthesizeSpeechInput{}
				if _, err := leanruntime.ApplyInput(input, fields_synthesize_speech, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SynthesizeSpeech(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("polly", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
