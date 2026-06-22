package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/transcribe"
)

var fields_create_call_analytics_category = []leanruntime.Field{
	{Name: "CategoryName", Flag: "category-name", Type: "*string", Required: true},
	{Name: "InputType", Flag: "input-type", Type: "types.InputType", Required: false},
	{Name: "Rules", Flag: "rules", Type: "[]types.Rule", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_language_model = []leanruntime.Field{
	{Name: "BaseModelName", Flag: "base-model-name", Type: "types.BaseModelName", Required: true},
	{Name: "InputDataConfig", Flag: "input-data-config", Type: "*types.InputDataConfig", Required: true},
	{Name: "LanguageCode", Flag: "language-code", Type: "types.CLMLanguageCode", Required: true},
	{Name: "ModelName", Flag: "model-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_medical_vocabulary = []leanruntime.Field{
	{Name: "LanguageCode", Flag: "language-code", Type: "types.LanguageCode", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "VocabularyFileUri", Flag: "vocabulary-file-uri", Type: "*string", Required: true},
	{Name: "VocabularyName", Flag: "vocabulary-name", Type: "*string", Required: true},
}

var fields_create_vocabulary = []leanruntime.Field{
	{Name: "DataAccessRoleArn", Flag: "data-access-role-arn", Type: "*string", Required: false},
	{Name: "LanguageCode", Flag: "language-code", Type: "types.LanguageCode", Required: true},
	{Name: "Phrases", Flag: "phrases", Type: "[]string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "VocabularyFileUri", Flag: "vocabulary-file-uri", Type: "*string", Required: false},
	{Name: "VocabularyName", Flag: "vocabulary-name", Type: "*string", Required: true},
}

var fields_create_vocabulary_filter = []leanruntime.Field{
	{Name: "DataAccessRoleArn", Flag: "data-access-role-arn", Type: "*string", Required: false},
	{Name: "LanguageCode", Flag: "language-code", Type: "types.LanguageCode", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "VocabularyFilterFileUri", Flag: "vocabulary-filter-file-uri", Type: "*string", Required: false},
	{Name: "VocabularyFilterName", Flag: "vocabulary-filter-name", Type: "*string", Required: true},
	{Name: "Words", Flag: "words", Type: "[]string", Required: false},
}

var fields_delete_call_analytics_category = []leanruntime.Field{
	{Name: "CategoryName", Flag: "category-name", Type: "*string", Required: true},
}

var fields_delete_call_analytics_job = []leanruntime.Field{
	{Name: "CallAnalyticsJobName", Flag: "call-analytics-job-name", Type: "*string", Required: true},
}

var fields_delete_language_model = []leanruntime.Field{
	{Name: "ModelName", Flag: "model-name", Type: "*string", Required: true},
}

var fields_delete_medical_scribe_job = []leanruntime.Field{
	{Name: "MedicalScribeJobName", Flag: "medical-scribe-job-name", Type: "*string", Required: true},
}

var fields_delete_medical_transcription_job = []leanruntime.Field{
	{Name: "MedicalTranscriptionJobName", Flag: "medical-transcription-job-name", Type: "*string", Required: true},
}

var fields_delete_medical_vocabulary = []leanruntime.Field{
	{Name: "VocabularyName", Flag: "vocabulary-name", Type: "*string", Required: true},
}

var fields_delete_transcription_job = []leanruntime.Field{
	{Name: "TranscriptionJobName", Flag: "transcription-job-name", Type: "*string", Required: true},
}

var fields_delete_vocabulary = []leanruntime.Field{
	{Name: "VocabularyName", Flag: "vocabulary-name", Type: "*string", Required: true},
}

var fields_delete_vocabulary_filter = []leanruntime.Field{
	{Name: "VocabularyFilterName", Flag: "vocabulary-filter-name", Type: "*string", Required: true},
}

var fields_describe_language_model = []leanruntime.Field{
	{Name: "ModelName", Flag: "model-name", Type: "*string", Required: true},
}

var fields_get_call_analytics_category = []leanruntime.Field{
	{Name: "CategoryName", Flag: "category-name", Type: "*string", Required: true},
}

var fields_get_call_analytics_job = []leanruntime.Field{
	{Name: "CallAnalyticsJobName", Flag: "call-analytics-job-name", Type: "*string", Required: true},
}

var fields_get_medical_scribe_job = []leanruntime.Field{
	{Name: "MedicalScribeJobName", Flag: "medical-scribe-job-name", Type: "*string", Required: true},
}

var fields_get_medical_transcription_job = []leanruntime.Field{
	{Name: "MedicalTranscriptionJobName", Flag: "medical-transcription-job-name", Type: "*string", Required: true},
}

var fields_get_medical_vocabulary = []leanruntime.Field{
	{Name: "VocabularyName", Flag: "vocabulary-name", Type: "*string", Required: true},
}

var fields_get_transcription_job = []leanruntime.Field{
	{Name: "TranscriptionJobName", Flag: "transcription-job-name", Type: "*string", Required: true},
}

var fields_get_vocabulary = []leanruntime.Field{
	{Name: "VocabularyName", Flag: "vocabulary-name", Type: "*string", Required: true},
}

var fields_get_vocabulary_filter = []leanruntime.Field{
	{Name: "VocabularyFilterName", Flag: "vocabulary-filter-name", Type: "*string", Required: true},
}

var fields_list_call_analytics_categories = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_call_analytics_jobs = []leanruntime.Field{
	{Name: "JobNameContains", Flag: "job-name-contains", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.CallAnalyticsJobStatus", Required: false},
}

var fields_list_language_models = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NameContains", Flag: "name-contains", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StatusEquals", Flag: "status-equals", Type: "types.ModelStatus", Required: false},
}

var fields_list_medical_scribe_jobs = []leanruntime.Field{
	{Name: "JobNameContains", Flag: "job-name-contains", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.MedicalScribeJobStatus", Required: false},
}

var fields_list_medical_transcription_jobs = []leanruntime.Field{
	{Name: "JobNameContains", Flag: "job-name-contains", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.TranscriptionJobStatus", Required: false},
}

var fields_list_medical_vocabularies = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NameContains", Flag: "name-contains", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StateEquals", Flag: "state-equals", Type: "types.VocabularyState", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_transcription_jobs = []leanruntime.Field{
	{Name: "JobNameContains", Flag: "job-name-contains", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.TranscriptionJobStatus", Required: false},
}

var fields_list_vocabularies = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NameContains", Flag: "name-contains", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StateEquals", Flag: "state-equals", Type: "types.VocabularyState", Required: false},
}

var fields_list_vocabulary_filters = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NameContains", Flag: "name-contains", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_start_call_analytics_job = []leanruntime.Field{
	{Name: "CallAnalyticsJobName", Flag: "call-analytics-job-name", Type: "*string", Required: true},
	{Name: "ChannelDefinitions", Flag: "channel-definitions", Type: "[]types.ChannelDefinition", Required: false},
	{Name: "DataAccessRoleArn", Flag: "data-access-role-arn", Type: "*string", Required: false},
	{Name: "Media", Flag: "media", Type: "*types.Media", Required: true},
	{Name: "OutputEncryptionKMSKeyId", Flag: "output-encryption-kms-key-id", Type: "*string", Required: false},
	{Name: "OutputLocation", Flag: "output-location", Type: "*string", Required: false},
	{Name: "Settings", Flag: "settings", Type: "*types.CallAnalyticsJobSettings", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_start_medical_scribe_job = []leanruntime.Field{
	{Name: "ChannelDefinitions", Flag: "channel-definitions", Type: "[]types.MedicalScribeChannelDefinition", Required: false},
	{Name: "DataAccessRoleArn", Flag: "data-access-role-arn", Type: "*string", Required: true},
	{Name: "KMSEncryptionContext", Flag: "kms-encryption-context", Type: "map[string]string", Required: false},
	{Name: "Media", Flag: "media", Type: "*types.Media", Required: true},
	{Name: "MedicalScribeContext", Flag: "medical-scribe-context", Type: "*types.MedicalScribeContext", Required: false},
	{Name: "MedicalScribeJobName", Flag: "medical-scribe-job-name", Type: "*string", Required: true},
	{Name: "OutputBucketName", Flag: "output-bucket-name", Type: "*string", Required: true},
	{Name: "OutputEncryptionKMSKeyId", Flag: "output-encryption-kms-key-id", Type: "*string", Required: false},
	{Name: "Settings", Flag: "settings", Type: "*types.MedicalScribeSettings", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_start_medical_transcription_job = []leanruntime.Field{
	{Name: "ContentIdentificationType", Flag: "content-identification-type", Type: "types.MedicalContentIdentificationType", Required: false},
	{Name: "KMSEncryptionContext", Flag: "kms-encryption-context", Type: "map[string]string", Required: false},
	{Name: "LanguageCode", Flag: "language-code", Type: "types.LanguageCode", Required: true},
	{Name: "Media", Flag: "media", Type: "*types.Media", Required: true},
	{Name: "MediaFormat", Flag: "media-format", Type: "types.MediaFormat", Required: false},
	{Name: "MediaSampleRateHertz", Flag: "media-sample-rate-hertz", Type: "*int32", Required: false},
	{Name: "MedicalTranscriptionJobName", Flag: "medical-transcription-job-name", Type: "*string", Required: true},
	{Name: "OutputBucketName", Flag: "output-bucket-name", Type: "*string", Required: true},
	{Name: "OutputEncryptionKMSKeyId", Flag: "output-encryption-kms-key-id", Type: "*string", Required: false},
	{Name: "OutputKey", Flag: "output-key", Type: "*string", Required: false},
	{Name: "Settings", Flag: "settings", Type: "*types.MedicalTranscriptionSetting", Required: false},
	{Name: "Specialty", Flag: "specialty", Type: "types.Specialty", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "Type", Flag: "type", Type: "types.Type", Required: true},
}

var fields_start_transcription_job = []leanruntime.Field{
	{Name: "ContentRedaction", Flag: "content-redaction", Type: "*types.ContentRedaction", Required: false},
	{Name: "IdentifyLanguage", Flag: "identify-language", Type: "*bool", Required: false},
	{Name: "IdentifyMultipleLanguages", Flag: "identify-multiple-languages", Type: "*bool", Required: false},
	{Name: "JobExecutionSettings", Flag: "job-execution-settings", Type: "*types.JobExecutionSettings", Required: false},
	{Name: "KMSEncryptionContext", Flag: "kms-encryption-context", Type: "map[string]string", Required: false},
	{Name: "LanguageCode", Flag: "language-code", Type: "types.LanguageCode", Required: false},
	{Name: "LanguageIdSettings", Flag: "language-id-settings", Type: "map[string]types.LanguageIdSettings", Required: false},
	{Name: "LanguageOptions", Flag: "language-options", Type: "[]types.LanguageCode", Required: false},
	{Name: "Media", Flag: "media", Type: "*types.Media", Required: true},
	{Name: "MediaFormat", Flag: "media-format", Type: "types.MediaFormat", Required: false},
	{Name: "MediaSampleRateHertz", Flag: "media-sample-rate-hertz", Type: "*int32", Required: false},
	{Name: "ModelSettings", Flag: "model-settings", Type: "*types.ModelSettings", Required: false},
	{Name: "OutputBucketName", Flag: "output-bucket-name", Type: "*string", Required: false},
	{Name: "OutputEncryptionKMSKeyId", Flag: "output-encryption-kms-key-id", Type: "*string", Required: false},
	{Name: "OutputKey", Flag: "output-key", Type: "*string", Required: false},
	{Name: "Settings", Flag: "settings", Type: "*types.Settings", Required: false},
	{Name: "Subtitles", Flag: "subtitles", Type: "*types.Subtitles", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "ToxicityDetection", Flag: "toxicity-detection", Type: "[]types.ToxicityDetectionSettings", Required: false},
	{Name: "TranscriptionJobName", Flag: "transcription-job-name", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_call_analytics_category = []leanruntime.Field{
	{Name: "CategoryName", Flag: "category-name", Type: "*string", Required: true},
	{Name: "InputType", Flag: "input-type", Type: "types.InputType", Required: false},
	{Name: "Rules", Flag: "rules", Type: "[]types.Rule", Required: true},
}

var fields_update_medical_vocabulary = []leanruntime.Field{
	{Name: "LanguageCode", Flag: "language-code", Type: "types.LanguageCode", Required: true},
	{Name: "VocabularyFileUri", Flag: "vocabulary-file-uri", Type: "*string", Required: true},
	{Name: "VocabularyName", Flag: "vocabulary-name", Type: "*string", Required: true},
}

var fields_update_vocabulary = []leanruntime.Field{
	{Name: "DataAccessRoleArn", Flag: "data-access-role-arn", Type: "*string", Required: false},
	{Name: "LanguageCode", Flag: "language-code", Type: "types.LanguageCode", Required: true},
	{Name: "Phrases", Flag: "phrases", Type: "[]string", Required: false},
	{Name: "VocabularyFileUri", Flag: "vocabulary-file-uri", Type: "*string", Required: false},
	{Name: "VocabularyName", Flag: "vocabulary-name", Type: "*string", Required: true},
}

var fields_update_vocabulary_filter = []leanruntime.Field{
	{Name: "DataAccessRoleArn", Flag: "data-access-role-arn", Type: "*string", Required: false},
	{Name: "VocabularyFilterFileUri", Flag: "vocabulary-filter-file-uri", Type: "*string", Required: false},
	{Name: "VocabularyFilterName", Flag: "vocabulary-filter-name", Type: "*string", Required: true},
	{Name: "Words", Flag: "words", Type: "[]string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-call-analytics-category": {
			Name:   "create-call-analytics-category",
			Fields: fields_create_call_analytics_category,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCallAnalyticsCategoryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_call_analytics_category, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCallAnalyticsCategory(ctx, input)
			},
		},
		"create-language-model": {
			Name:   "create-language-model",
			Fields: fields_create_language_model,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLanguageModelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_language_model, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLanguageModel(ctx, input)
			},
		},
		"create-medical-vocabulary": {
			Name:   "create-medical-vocabulary",
			Fields: fields_create_medical_vocabulary,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateMedicalVocabularyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_medical_vocabulary, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateMedicalVocabulary(ctx, input)
			},
		},
		"create-vocabulary": {
			Name:   "create-vocabulary",
			Fields: fields_create_vocabulary,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateVocabularyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_vocabulary, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateVocabulary(ctx, input)
			},
		},
		"create-vocabulary-filter": {
			Name:   "create-vocabulary-filter",
			Fields: fields_create_vocabulary_filter,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateVocabularyFilterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_vocabulary_filter, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateVocabularyFilter(ctx, input)
			},
		},
		"delete-call-analytics-category": {
			Name:   "delete-call-analytics-category",
			Fields: fields_delete_call_analytics_category,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCallAnalyticsCategoryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_call_analytics_category, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCallAnalyticsCategory(ctx, input)
			},
		},
		"delete-call-analytics-job": {
			Name:   "delete-call-analytics-job",
			Fields: fields_delete_call_analytics_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCallAnalyticsJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_call_analytics_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCallAnalyticsJob(ctx, input)
			},
		},
		"delete-language-model": {
			Name:   "delete-language-model",
			Fields: fields_delete_language_model,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteLanguageModelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_language_model, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteLanguageModel(ctx, input)
			},
		},
		"delete-medical-scribe-job": {
			Name:   "delete-medical-scribe-job",
			Fields: fields_delete_medical_scribe_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMedicalScribeJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_medical_scribe_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMedicalScribeJob(ctx, input)
			},
		},
		"delete-medical-transcription-job": {
			Name:   "delete-medical-transcription-job",
			Fields: fields_delete_medical_transcription_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMedicalTranscriptionJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_medical_transcription_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMedicalTranscriptionJob(ctx, input)
			},
		},
		"delete-medical-vocabulary": {
			Name:   "delete-medical-vocabulary",
			Fields: fields_delete_medical_vocabulary,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMedicalVocabularyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_medical_vocabulary, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMedicalVocabulary(ctx, input)
			},
		},
		"delete-transcription-job": {
			Name:   "delete-transcription-job",
			Fields: fields_delete_transcription_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTranscriptionJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_transcription_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTranscriptionJob(ctx, input)
			},
		},
		"delete-vocabulary": {
			Name:   "delete-vocabulary",
			Fields: fields_delete_vocabulary,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVocabularyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_vocabulary, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVocabulary(ctx, input)
			},
		},
		"delete-vocabulary-filter": {
			Name:   "delete-vocabulary-filter",
			Fields: fields_delete_vocabulary_filter,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVocabularyFilterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_vocabulary_filter, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVocabularyFilter(ctx, input)
			},
		},
		"describe-language-model": {
			Name:   "describe-language-model",
			Fields: fields_describe_language_model,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeLanguageModelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_language_model, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeLanguageModel(ctx, input)
			},
		},
		"get-call-analytics-category": {
			Name:   "get-call-analytics-category",
			Fields: fields_get_call_analytics_category,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCallAnalyticsCategoryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_call_analytics_category, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCallAnalyticsCategory(ctx, input)
			},
		},
		"get-call-analytics-job": {
			Name:   "get-call-analytics-job",
			Fields: fields_get_call_analytics_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCallAnalyticsJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_call_analytics_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCallAnalyticsJob(ctx, input)
			},
		},
		"get-medical-scribe-job": {
			Name:   "get-medical-scribe-job",
			Fields: fields_get_medical_scribe_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMedicalScribeJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_medical_scribe_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMedicalScribeJob(ctx, input)
			},
		},
		"get-medical-transcription-job": {
			Name:   "get-medical-transcription-job",
			Fields: fields_get_medical_transcription_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMedicalTranscriptionJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_medical_transcription_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMedicalTranscriptionJob(ctx, input)
			},
		},
		"get-medical-vocabulary": {
			Name:   "get-medical-vocabulary",
			Fields: fields_get_medical_vocabulary,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMedicalVocabularyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_medical_vocabulary, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMedicalVocabulary(ctx, input)
			},
		},
		"get-transcription-job": {
			Name:   "get-transcription-job",
			Fields: fields_get_transcription_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTranscriptionJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_transcription_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTranscriptionJob(ctx, input)
			},
		},
		"get-vocabulary": {
			Name:   "get-vocabulary",
			Fields: fields_get_vocabulary,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetVocabularyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_vocabulary, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetVocabulary(ctx, input)
			},
		},
		"get-vocabulary-filter": {
			Name:   "get-vocabulary-filter",
			Fields: fields_get_vocabulary_filter,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetVocabularyFilterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_vocabulary_filter, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetVocabularyFilter(ctx, input)
			},
		},
		"list-call-analytics-categories": {
			Name:   "list-call-analytics-categories",
			Fields: fields_list_call_analytics_categories,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCallAnalyticsCategoriesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_call_analytics_categories, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCallAnalyticsCategories(ctx, input)
				}
				var results []*svc.ListCallAnalyticsCategoriesOutput
				p := svc.NewListCallAnalyticsCategoriesPaginator(client, input)
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
		"list-call-analytics-jobs": {
			Name:   "list-call-analytics-jobs",
			Fields: fields_list_call_analytics_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCallAnalyticsJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_call_analytics_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCallAnalyticsJobs(ctx, input)
				}
				var results []*svc.ListCallAnalyticsJobsOutput
				p := svc.NewListCallAnalyticsJobsPaginator(client, input)
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
		"list-language-models": {
			Name:   "list-language-models",
			Fields: fields_list_language_models,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLanguageModelsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_language_models, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListLanguageModels(ctx, input)
				}
				var results []*svc.ListLanguageModelsOutput
				p := svc.NewListLanguageModelsPaginator(client, input)
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
		"list-medical-scribe-jobs": {
			Name:   "list-medical-scribe-jobs",
			Fields: fields_list_medical_scribe_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMedicalScribeJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_medical_scribe_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMedicalScribeJobs(ctx, input)
				}
				var results []*svc.ListMedicalScribeJobsOutput
				p := svc.NewListMedicalScribeJobsPaginator(client, input)
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
		"list-medical-transcription-jobs": {
			Name:   "list-medical-transcription-jobs",
			Fields: fields_list_medical_transcription_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMedicalTranscriptionJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_medical_transcription_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMedicalTranscriptionJobs(ctx, input)
				}
				var results []*svc.ListMedicalTranscriptionJobsOutput
				p := svc.NewListMedicalTranscriptionJobsPaginator(client, input)
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
		"list-medical-vocabularies": {
			Name:   "list-medical-vocabularies",
			Fields: fields_list_medical_vocabularies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMedicalVocabulariesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_medical_vocabularies, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMedicalVocabularies(ctx, input)
				}
				var results []*svc.ListMedicalVocabulariesOutput
				p := svc.NewListMedicalVocabulariesPaginator(client, input)
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
		"list-tags-for-resource": {
			Name:   "list-tags-for-resource",
			Fields: fields_list_tags_for_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTagsForResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_tags_for_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListTagsForResource(ctx, input)
			},
		},
		"list-transcription-jobs": {
			Name:   "list-transcription-jobs",
			Fields: fields_list_transcription_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTranscriptionJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_transcription_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTranscriptionJobs(ctx, input)
				}
				var results []*svc.ListTranscriptionJobsOutput
				p := svc.NewListTranscriptionJobsPaginator(client, input)
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
		"list-vocabularies": {
			Name:   "list-vocabularies",
			Fields: fields_list_vocabularies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListVocabulariesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_vocabularies, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListVocabularies(ctx, input)
				}
				var results []*svc.ListVocabulariesOutput
				p := svc.NewListVocabulariesPaginator(client, input)
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
		"list-vocabulary-filters": {
			Name:   "list-vocabulary-filters",
			Fields: fields_list_vocabulary_filters,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListVocabularyFiltersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_vocabulary_filters, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListVocabularyFilters(ctx, input)
				}
				var results []*svc.ListVocabularyFiltersOutput
				p := svc.NewListVocabularyFiltersPaginator(client, input)
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
		"start-call-analytics-job": {
			Name:   "start-call-analytics-job",
			Fields: fields_start_call_analytics_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartCallAnalyticsJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_call_analytics_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartCallAnalyticsJob(ctx, input)
			},
		},
		"start-medical-scribe-job": {
			Name:   "start-medical-scribe-job",
			Fields: fields_start_medical_scribe_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartMedicalScribeJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_medical_scribe_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartMedicalScribeJob(ctx, input)
			},
		},
		"start-medical-transcription-job": {
			Name:   "start-medical-transcription-job",
			Fields: fields_start_medical_transcription_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartMedicalTranscriptionJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_medical_transcription_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartMedicalTranscriptionJob(ctx, input)
			},
		},
		"start-transcription-job": {
			Name:   "start-transcription-job",
			Fields: fields_start_transcription_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartTranscriptionJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_transcription_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartTranscriptionJob(ctx, input)
			},
		},
		"tag-resource": {
			Name:   "tag-resource",
			Fields: fields_tag_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TagResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_tag_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TagResource(ctx, input)
			},
		},
		"untag-resource": {
			Name:   "untag-resource",
			Fields: fields_untag_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UntagResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_untag_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UntagResource(ctx, input)
			},
		},
		"update-call-analytics-category": {
			Name:   "update-call-analytics-category",
			Fields: fields_update_call_analytics_category,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateCallAnalyticsCategoryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_call_analytics_category, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateCallAnalyticsCategory(ctx, input)
			},
		},
		"update-medical-vocabulary": {
			Name:   "update-medical-vocabulary",
			Fields: fields_update_medical_vocabulary,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateMedicalVocabularyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_medical_vocabulary, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateMedicalVocabulary(ctx, input)
			},
		},
		"update-vocabulary": {
			Name:   "update-vocabulary",
			Fields: fields_update_vocabulary,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateVocabularyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_vocabulary, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateVocabulary(ctx, input)
			},
		},
		"update-vocabulary-filter": {
			Name:   "update-vocabulary-filter",
			Fields: fields_update_vocabulary_filter,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateVocabularyFilterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_vocabulary_filter, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateVocabularyFilter(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("transcribe", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
