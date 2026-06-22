package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/translate"
)

var fields_create_parallel_data = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EncryptionKey", Flag: "encryption-key", Type: "*types.EncryptionKey", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ParallelDataConfig", Flag: "parallel-data-config", Type: "*types.ParallelDataConfig", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_delete_parallel_data = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_terminology = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_describe_text_translation_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_get_parallel_data = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_terminology = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "TerminologyDataFormat", Flag: "terminology-data-format", Type: "types.TerminologyDataFormat", Required: false},
}

var fields_import_terminology = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EncryptionKey", Flag: "encryption-key", Type: "*types.EncryptionKey", Required: false},
	{Name: "MergeStrategy", Flag: "merge-strategy", Type: "types.MergeStrategy", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TerminologyData", Flag: "terminology-data", Type: "*types.TerminologyData", Required: true},
}

var fields_list_languages = []leanruntime.Field{
	{Name: "DisplayLanguageCode", Flag: "display-language-code", Type: "types.DisplayLanguageCode", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_parallel_data = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_terminologies = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_text_translation_jobs = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.TextTranslationJobFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_start_text_translation_job = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "DataAccessRoleArn", Flag: "data-access-role-arn", Type: "*string", Required: true},
	{Name: "InputDataConfig", Flag: "input-data-config", Type: "*types.InputDataConfig", Required: true},
	{Name: "JobName", Flag: "job-name", Type: "*string", Required: false},
	{Name: "OutputDataConfig", Flag: "output-data-config", Type: "*types.OutputDataConfig", Required: true},
	{Name: "ParallelDataNames", Flag: "parallel-data-names", Type: "[]string", Required: false},
	{Name: "Settings", Flag: "settings", Type: "*types.TranslationSettings", Required: false},
	{Name: "SourceLanguageCode", Flag: "source-language-code", Type: "*string", Required: true},
	{Name: "TargetLanguageCodes", Flag: "target-language-codes", Type: "[]string", Required: true},
	{Name: "TerminologyNames", Flag: "terminology-names", Type: "[]string", Required: false},
}

var fields_stop_text_translation_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_translate_document = []leanruntime.Field{
	{Name: "Document", Flag: "document", Type: "*types.Document", Required: true},
	{Name: "Settings", Flag: "settings", Type: "*types.TranslationSettings", Required: false},
	{Name: "SourceLanguageCode", Flag: "source-language-code", Type: "*string", Required: true},
	{Name: "TargetLanguageCode", Flag: "target-language-code", Type: "*string", Required: true},
	{Name: "TerminologyNames", Flag: "terminology-names", Type: "[]string", Required: false},
}

var fields_translate_text = []leanruntime.Field{
	{Name: "Settings", Flag: "settings", Type: "*types.TranslationSettings", Required: false},
	{Name: "SourceLanguageCode", Flag: "source-language-code", Type: "*string", Required: true},
	{Name: "TargetLanguageCode", Flag: "target-language-code", Type: "*string", Required: true},
	{Name: "TerminologyNames", Flag: "terminology-names", Type: "[]string", Required: false},
	{Name: "Text", Flag: "text", Type: "*string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_parallel_data = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ParallelDataConfig", Flag: "parallel-data-config", Type: "*types.ParallelDataConfig", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-parallel-data": {
			Name:   "create-parallel-data",
			Fields: fields_create_parallel_data,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateParallelDataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_parallel_data, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateParallelData(ctx, input)
			},
		},
		"delete-parallel-data": {
			Name:   "delete-parallel-data",
			Fields: fields_delete_parallel_data,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteParallelDataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_parallel_data, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteParallelData(ctx, input)
			},
		},
		"delete-terminology": {
			Name:   "delete-terminology",
			Fields: fields_delete_terminology,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTerminologyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_terminology, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTerminology(ctx, input)
			},
		},
		"describe-text-translation-job": {
			Name:   "describe-text-translation-job",
			Fields: fields_describe_text_translation_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTextTranslationJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_text_translation_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeTextTranslationJob(ctx, input)
			},
		},
		"get-parallel-data": {
			Name:   "get-parallel-data",
			Fields: fields_get_parallel_data,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetParallelDataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_parallel_data, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetParallelData(ctx, input)
			},
		},
		"get-terminology": {
			Name:   "get-terminology",
			Fields: fields_get_terminology,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTerminologyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_terminology, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTerminology(ctx, input)
			},
		},
		"import-terminology": {
			Name:   "import-terminology",
			Fields: fields_import_terminology,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ImportTerminologyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_import_terminology, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ImportTerminology(ctx, input)
			},
		},
		"list-languages": {
			Name:   "list-languages",
			Fields: fields_list_languages,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLanguagesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_languages, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListLanguages(ctx, input)
				}
				var results []*svc.ListLanguagesOutput
				p := svc.NewListLanguagesPaginator(client, input)
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
		"list-parallel-data": {
			Name:   "list-parallel-data",
			Fields: fields_list_parallel_data,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListParallelDataInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_parallel_data, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListParallelData(ctx, input)
				}
				var results []*svc.ListParallelDataOutput
				p := svc.NewListParallelDataPaginator(client, input)
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
		"list-terminologies": {
			Name:   "list-terminologies",
			Fields: fields_list_terminologies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTerminologiesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_terminologies, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTerminologies(ctx, input)
				}
				var results []*svc.ListTerminologiesOutput
				p := svc.NewListTerminologiesPaginator(client, input)
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
		"list-text-translation-jobs": {
			Name:   "list-text-translation-jobs",
			Fields: fields_list_text_translation_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTextTranslationJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_text_translation_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTextTranslationJobs(ctx, input)
				}
				var results []*svc.ListTextTranslationJobsOutput
				p := svc.NewListTextTranslationJobsPaginator(client, input)
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
		"start-text-translation-job": {
			Name:   "start-text-translation-job",
			Fields: fields_start_text_translation_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartTextTranslationJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_text_translation_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartTextTranslationJob(ctx, input)
			},
		},
		"stop-text-translation-job": {
			Name:   "stop-text-translation-job",
			Fields: fields_stop_text_translation_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopTextTranslationJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_text_translation_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopTextTranslationJob(ctx, input)
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
		"translate-document": {
			Name:   "translate-document",
			Fields: fields_translate_document,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TranslateDocumentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_translate_document, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TranslateDocument(ctx, input)
			},
		},
		"translate-text": {
			Name:   "translate-text",
			Fields: fields_translate_text,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TranslateTextInput{}
				if _, err := leanruntime.ApplyInput(input, fields_translate_text, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TranslateText(ctx, input)
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
		"update-parallel-data": {
			Name:   "update-parallel-data",
			Fields: fields_update_parallel_data,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateParallelDataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_parallel_data, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateParallelData(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("translate", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
