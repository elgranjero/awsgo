package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/textract"
)

var fields_analyze_document = []leanruntime.Field{
	{Name: "AdaptersConfig", Flag: "adapters-config", Type: "*types.AdaptersConfig", Required: false},
	{Name: "Document", Flag: "document", Type: "*types.Document", Required: true},
	{Name: "FeatureTypes", Flag: "feature-types", Type: "[]types.FeatureType", Required: true},
	{Name: "HumanLoopConfig", Flag: "human-loop-config", Type: "*types.HumanLoopConfig", Required: false},
	{Name: "QueriesConfig", Flag: "queries-config", Type: "*types.QueriesConfig", Required: false},
}

var fields_analyze_expense = []leanruntime.Field{
	{Name: "Document", Flag: "document", Type: "*types.Document", Required: true},
}

var fields_analyze_id = []leanruntime.Field{
	{Name: "DocumentPages", Flag: "document-pages", Type: "[]types.Document", Required: true},
}

var fields_create_adapter = []leanruntime.Field{
	{Name: "AdapterName", Flag: "adapter-name", Type: "*string", Required: true},
	{Name: "AutoUpdate", Flag: "auto-update", Type: "types.AutoUpdate", Required: false},
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "FeatureTypes", Flag: "feature-types", Type: "[]types.FeatureType", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_adapter_version = []leanruntime.Field{
	{Name: "AdapterId", Flag: "adapter-id", Type: "*string", Required: true},
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "DatasetConfig", Flag: "dataset-config", Type: "*types.AdapterVersionDatasetConfig", Required: true},
	{Name: "KMSKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "OutputConfig", Flag: "output-config", Type: "*types.OutputConfig", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_delete_adapter = []leanruntime.Field{
	{Name: "AdapterId", Flag: "adapter-id", Type: "*string", Required: true},
}

var fields_delete_adapter_version = []leanruntime.Field{
	{Name: "AdapterId", Flag: "adapter-id", Type: "*string", Required: true},
	{Name: "AdapterVersion", Flag: "adapter-version", Type: "*string", Required: true},
}

var fields_detect_document_text = []leanruntime.Field{
	{Name: "Document", Flag: "document", Type: "*types.Document", Required: true},
}

var fields_get_adapter = []leanruntime.Field{
	{Name: "AdapterId", Flag: "adapter-id", Type: "*string", Required: true},
}

var fields_get_adapter_version = []leanruntime.Field{
	{Name: "AdapterId", Flag: "adapter-id", Type: "*string", Required: true},
	{Name: "AdapterVersion", Flag: "adapter-version", Type: "*string", Required: true},
}

var fields_get_document_analysis = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_document_text_detection = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_expense_analysis = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_lending_analysis = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_lending_analysis_summary = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_list_adapter_versions = []leanruntime.Field{
	{Name: "AdapterId", Flag: "adapter-id", Type: "*string", Required: false},
	{Name: "AfterCreationTime", Flag: "after-creation-time", Type: "*time.Time", Required: false},
	{Name: "BeforeCreationTime", Flag: "before-creation-time", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_adapters = []leanruntime.Field{
	{Name: "AfterCreationTime", Flag: "after-creation-time", Type: "*time.Time", Required: false},
	{Name: "BeforeCreationTime", Flag: "before-creation-time", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_start_document_analysis = []leanruntime.Field{
	{Name: "AdaptersConfig", Flag: "adapters-config", Type: "*types.AdaptersConfig", Required: false},
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "DocumentLocation", Flag: "document-location", Type: "*types.DocumentLocation", Required: true},
	{Name: "FeatureTypes", Flag: "feature-types", Type: "[]types.FeatureType", Required: true},
	{Name: "JobTag", Flag: "job-tag", Type: "*string", Required: false},
	{Name: "KMSKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "NotificationChannel", Flag: "notification-channel", Type: "*types.NotificationChannel", Required: false},
	{Name: "OutputConfig", Flag: "output-config", Type: "*types.OutputConfig", Required: false},
	{Name: "QueriesConfig", Flag: "queries-config", Type: "*types.QueriesConfig", Required: false},
}

var fields_start_document_text_detection = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "DocumentLocation", Flag: "document-location", Type: "*types.DocumentLocation", Required: true},
	{Name: "JobTag", Flag: "job-tag", Type: "*string", Required: false},
	{Name: "KMSKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "NotificationChannel", Flag: "notification-channel", Type: "*types.NotificationChannel", Required: false},
	{Name: "OutputConfig", Flag: "output-config", Type: "*types.OutputConfig", Required: false},
}

var fields_start_expense_analysis = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "DocumentLocation", Flag: "document-location", Type: "*types.DocumentLocation", Required: true},
	{Name: "JobTag", Flag: "job-tag", Type: "*string", Required: false},
	{Name: "KMSKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "NotificationChannel", Flag: "notification-channel", Type: "*types.NotificationChannel", Required: false},
	{Name: "OutputConfig", Flag: "output-config", Type: "*types.OutputConfig", Required: false},
}

var fields_start_lending_analysis = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "DocumentLocation", Flag: "document-location", Type: "*types.DocumentLocation", Required: true},
	{Name: "JobTag", Flag: "job-tag", Type: "*string", Required: false},
	{Name: "KMSKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "NotificationChannel", Flag: "notification-channel", Type: "*types.NotificationChannel", Required: false},
	{Name: "OutputConfig", Flag: "output-config", Type: "*types.OutputConfig", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_adapter = []leanruntime.Field{
	{Name: "AdapterId", Flag: "adapter-id", Type: "*string", Required: true},
	{Name: "AdapterName", Flag: "adapter-name", Type: "*string", Required: false},
	{Name: "AutoUpdate", Flag: "auto-update", Type: "types.AutoUpdate", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"analyze-document": {
			Name:   "analyze-document",
			Fields: fields_analyze_document,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AnalyzeDocumentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_analyze_document, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AnalyzeDocument(ctx, input)
			},
		},
		"analyze-expense": {
			Name:   "analyze-expense",
			Fields: fields_analyze_expense,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AnalyzeExpenseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_analyze_expense, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AnalyzeExpense(ctx, input)
			},
		},
		"analyze-id": {
			Name:   "analyze-id",
			Fields: fields_analyze_id,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AnalyzeIDInput{}
				if _, err := leanruntime.ApplyInput(input, fields_analyze_id, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AnalyzeID(ctx, input)
			},
		},
		"create-adapter": {
			Name:   "create-adapter",
			Fields: fields_create_adapter,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAdapterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_adapter, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAdapter(ctx, input)
			},
		},
		"create-adapter-version": {
			Name:   "create-adapter-version",
			Fields: fields_create_adapter_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAdapterVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_adapter_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAdapterVersion(ctx, input)
			},
		},
		"delete-adapter": {
			Name:   "delete-adapter",
			Fields: fields_delete_adapter,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAdapterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_adapter, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAdapter(ctx, input)
			},
		},
		"delete-adapter-version": {
			Name:   "delete-adapter-version",
			Fields: fields_delete_adapter_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAdapterVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_adapter_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAdapterVersion(ctx, input)
			},
		},
		"detect-document-text": {
			Name:   "detect-document-text",
			Fields: fields_detect_document_text,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DetectDocumentTextInput{}
				if _, err := leanruntime.ApplyInput(input, fields_detect_document_text, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DetectDocumentText(ctx, input)
			},
		},
		"get-adapter": {
			Name:   "get-adapter",
			Fields: fields_get_adapter,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAdapterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_adapter, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAdapter(ctx, input)
			},
		},
		"get-adapter-version": {
			Name:   "get-adapter-version",
			Fields: fields_get_adapter_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAdapterVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_adapter_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAdapterVersion(ctx, input)
			},
		},
		"get-document-analysis": {
			Name:   "get-document-analysis",
			Fields: fields_get_document_analysis,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDocumentAnalysisInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_document_analysis, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDocumentAnalysis(ctx, input)
			},
		},
		"get-document-text-detection": {
			Name:   "get-document-text-detection",
			Fields: fields_get_document_text_detection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDocumentTextDetectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_document_text_detection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDocumentTextDetection(ctx, input)
			},
		},
		"get-expense-analysis": {
			Name:   "get-expense-analysis",
			Fields: fields_get_expense_analysis,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetExpenseAnalysisInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_expense_analysis, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetExpenseAnalysis(ctx, input)
			},
		},
		"get-lending-analysis": {
			Name:   "get-lending-analysis",
			Fields: fields_get_lending_analysis,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLendingAnalysisInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_lending_analysis, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetLendingAnalysis(ctx, input)
			},
		},
		"get-lending-analysis-summary": {
			Name:   "get-lending-analysis-summary",
			Fields: fields_get_lending_analysis_summary,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLendingAnalysisSummaryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_lending_analysis_summary, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetLendingAnalysisSummary(ctx, input)
			},
		},
		"list-adapter-versions": {
			Name:   "list-adapter-versions",
			Fields: fields_list_adapter_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAdapterVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_adapter_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAdapterVersions(ctx, input)
				}
				var results []*svc.ListAdapterVersionsOutput
				p := svc.NewListAdapterVersionsPaginator(client, input)
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
		"list-adapters": {
			Name:   "list-adapters",
			Fields: fields_list_adapters,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAdaptersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_adapters, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAdapters(ctx, input)
				}
				var results []*svc.ListAdaptersOutput
				p := svc.NewListAdaptersPaginator(client, input)
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
		"start-document-analysis": {
			Name:   "start-document-analysis",
			Fields: fields_start_document_analysis,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartDocumentAnalysisInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_document_analysis, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartDocumentAnalysis(ctx, input)
			},
		},
		"start-document-text-detection": {
			Name:   "start-document-text-detection",
			Fields: fields_start_document_text_detection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartDocumentTextDetectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_document_text_detection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartDocumentTextDetection(ctx, input)
			},
		},
		"start-expense-analysis": {
			Name:   "start-expense-analysis",
			Fields: fields_start_expense_analysis,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartExpenseAnalysisInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_expense_analysis, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartExpenseAnalysis(ctx, input)
			},
		},
		"start-lending-analysis": {
			Name:   "start-lending-analysis",
			Fields: fields_start_lending_analysis,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartLendingAnalysisInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_lending_analysis, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartLendingAnalysis(ctx, input)
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
		"update-adapter": {
			Name:   "update-adapter",
			Fields: fields_update_adapter,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAdapterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_adapter, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAdapter(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("textract", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
