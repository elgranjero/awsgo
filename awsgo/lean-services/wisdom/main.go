package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/wisdom"
)

var fields_create_assistant = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ServerSideEncryptionConfiguration", Flag: "server-side-encryption-configuration", Type: "*types.ServerSideEncryptionConfiguration", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.AssistantType", Required: true},
}

var fields_create_assistant_association = []leanruntime.Field{
	{Name: "AssistantId", Flag: "assistant-id", Type: "*string", Required: true},
	{Name: "Association", Flag: "association", Type: "types.AssistantAssociationInputData", Required: true},
	{Name: "AssociationType", Flag: "association-type", Type: "types.AssociationType", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_content = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "KnowledgeBaseId", Flag: "knowledge-base-id", Type: "*string", Required: true},
	{Name: "Metadata", Flag: "metadata", Type: "map[string]string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "OverrideLinkOutUri", Flag: "override-link-out-uri", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Title", Flag: "title", Type: "*string", Required: false},
	{Name: "UploadId", Flag: "upload-id", Type: "*string", Required: true},
}

var fields_create_knowledge_base = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "KnowledgeBaseType", Flag: "knowledge-base-type", Type: "types.KnowledgeBaseType", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RenderingConfiguration", Flag: "rendering-configuration", Type: "*types.RenderingConfiguration", Required: false},
	{Name: "ServerSideEncryptionConfiguration", Flag: "server-side-encryption-configuration", Type: "*types.ServerSideEncryptionConfiguration", Required: false},
	{Name: "SourceConfiguration", Flag: "source-configuration", Type: "types.SourceConfiguration", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_quick_response = []leanruntime.Field{
	{Name: "Channels", Flag: "channels", Type: "[]string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Content", Flag: "content", Type: "types.QuickResponseDataProvider", Required: true},
	{Name: "ContentType", Flag: "content-type", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "GroupingConfiguration", Flag: "grouping-configuration", Type: "*types.GroupingConfiguration", Required: false},
	{Name: "IsActive", Flag: "is-active", Type: "*bool", Required: false},
	{Name: "KnowledgeBaseId", Flag: "knowledge-base-id", Type: "*string", Required: true},
	{Name: "Language", Flag: "language", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ShortcutKey", Flag: "shortcut-key", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_session = []leanruntime.Field{
	{Name: "AssistantId", Flag: "assistant-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_delete_assistant = []leanruntime.Field{
	{Name: "AssistantId", Flag: "assistant-id", Type: "*string", Required: true},
}

var fields_delete_assistant_association = []leanruntime.Field{
	{Name: "AssistantAssociationId", Flag: "assistant-association-id", Type: "*string", Required: true},
	{Name: "AssistantId", Flag: "assistant-id", Type: "*string", Required: true},
}

var fields_delete_content = []leanruntime.Field{
	{Name: "ContentId", Flag: "content-id", Type: "*string", Required: true},
	{Name: "KnowledgeBaseId", Flag: "knowledge-base-id", Type: "*string", Required: true},
}

var fields_delete_import_job = []leanruntime.Field{
	{Name: "ImportJobId", Flag: "import-job-id", Type: "*string", Required: true},
	{Name: "KnowledgeBaseId", Flag: "knowledge-base-id", Type: "*string", Required: true},
}

var fields_delete_knowledge_base = []leanruntime.Field{
	{Name: "KnowledgeBaseId", Flag: "knowledge-base-id", Type: "*string", Required: true},
}

var fields_delete_quick_response = []leanruntime.Field{
	{Name: "KnowledgeBaseId", Flag: "knowledge-base-id", Type: "*string", Required: true},
	{Name: "QuickResponseId", Flag: "quick-response-id", Type: "*string", Required: true},
}

var fields_get_assistant = []leanruntime.Field{
	{Name: "AssistantId", Flag: "assistant-id", Type: "*string", Required: true},
}

var fields_get_assistant_association = []leanruntime.Field{
	{Name: "AssistantAssociationId", Flag: "assistant-association-id", Type: "*string", Required: true},
	{Name: "AssistantId", Flag: "assistant-id", Type: "*string", Required: true},
}

var fields_get_content = []leanruntime.Field{
	{Name: "ContentId", Flag: "content-id", Type: "*string", Required: true},
	{Name: "KnowledgeBaseId", Flag: "knowledge-base-id", Type: "*string", Required: true},
}

var fields_get_content_summary = []leanruntime.Field{
	{Name: "ContentId", Flag: "content-id", Type: "*string", Required: true},
	{Name: "KnowledgeBaseId", Flag: "knowledge-base-id", Type: "*string", Required: true},
}

var fields_get_import_job = []leanruntime.Field{
	{Name: "ImportJobId", Flag: "import-job-id", Type: "*string", Required: true},
	{Name: "KnowledgeBaseId", Flag: "knowledge-base-id", Type: "*string", Required: true},
}

var fields_get_knowledge_base = []leanruntime.Field{
	{Name: "KnowledgeBaseId", Flag: "knowledge-base-id", Type: "*string", Required: true},
}

var fields_get_quick_response = []leanruntime.Field{
	{Name: "KnowledgeBaseId", Flag: "knowledge-base-id", Type: "*string", Required: true},
	{Name: "QuickResponseId", Flag: "quick-response-id", Type: "*string", Required: true},
}

var fields_get_recommendations = []leanruntime.Field{
	{Name: "AssistantId", Flag: "assistant-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: true},
	{Name: "WaitTimeSeconds", Flag: "wait-time-seconds", Type: "int32", Required: false},
}

var fields_get_session = []leanruntime.Field{
	{Name: "AssistantId", Flag: "assistant-id", Type: "*string", Required: true},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: true},
}

var fields_list_assistant_associations = []leanruntime.Field{
	{Name: "AssistantId", Flag: "assistant-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_assistants = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_contents = []leanruntime.Field{
	{Name: "KnowledgeBaseId", Flag: "knowledge-base-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_import_jobs = []leanruntime.Field{
	{Name: "KnowledgeBaseId", Flag: "knowledge-base-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_knowledge_bases = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_quick_responses = []leanruntime.Field{
	{Name: "KnowledgeBaseId", Flag: "knowledge-base-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_notify_recommendations_received = []leanruntime.Field{
	{Name: "AssistantId", Flag: "assistant-id", Type: "*string", Required: true},
	{Name: "RecommendationIds", Flag: "recommendation-ids", Type: "[]string", Required: true},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: true},
}

var fields_query_assistant = []leanruntime.Field{
	{Name: "AssistantId", Flag: "assistant-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "QueryText", Flag: "query-text", Type: "*string", Required: true},
}

var fields_remove_knowledge_base_template_uri = []leanruntime.Field{
	{Name: "KnowledgeBaseId", Flag: "knowledge-base-id", Type: "*string", Required: true},
}

var fields_search_content = []leanruntime.Field{
	{Name: "KnowledgeBaseId", Flag: "knowledge-base-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SearchExpression", Flag: "search-expression", Type: "*types.SearchExpression", Required: true},
}

var fields_search_quick_responses = []leanruntime.Field{
	{Name: "Attributes", Flag: "attributes", Type: "map[string]string", Required: false},
	{Name: "KnowledgeBaseId", Flag: "knowledge-base-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SearchExpression", Flag: "search-expression", Type: "*types.QuickResponseSearchExpression", Required: true},
}

var fields_search_sessions = []leanruntime.Field{
	{Name: "AssistantId", Flag: "assistant-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SearchExpression", Flag: "search-expression", Type: "*types.SearchExpression", Required: true},
}

var fields_start_content_upload = []leanruntime.Field{
	{Name: "ContentType", Flag: "content-type", Type: "*string", Required: true},
	{Name: "KnowledgeBaseId", Flag: "knowledge-base-id", Type: "*string", Required: true},
	{Name: "PresignedUrlTimeToLive", Flag: "presigned-url-time-to-live", Type: "*int32", Required: false},
}

var fields_start_import_job = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ExternalSourceConfiguration", Flag: "external-source-configuration", Type: "*types.ExternalSourceConfiguration", Required: false},
	{Name: "ImportJobType", Flag: "import-job-type", Type: "types.ImportJobType", Required: true},
	{Name: "KnowledgeBaseId", Flag: "knowledge-base-id", Type: "*string", Required: true},
	{Name: "Metadata", Flag: "metadata", Type: "map[string]string", Required: false},
	{Name: "UploadId", Flag: "upload-id", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_content = []leanruntime.Field{
	{Name: "ContentId", Flag: "content-id", Type: "*string", Required: true},
	{Name: "KnowledgeBaseId", Flag: "knowledge-base-id", Type: "*string", Required: true},
	{Name: "Metadata", Flag: "metadata", Type: "map[string]string", Required: false},
	{Name: "OverrideLinkOutUri", Flag: "override-link-out-uri", Type: "*string", Required: false},
	{Name: "RemoveOverrideLinkOutUri", Flag: "remove-override-link-out-uri", Type: "*bool", Required: false},
	{Name: "RevisionId", Flag: "revision-id", Type: "*string", Required: false},
	{Name: "Title", Flag: "title", Type: "*string", Required: false},
	{Name: "UploadId", Flag: "upload-id", Type: "*string", Required: false},
}

var fields_update_knowledge_base_template_uri = []leanruntime.Field{
	{Name: "KnowledgeBaseId", Flag: "knowledge-base-id", Type: "*string", Required: true},
	{Name: "TemplateUri", Flag: "template-uri", Type: "*string", Required: true},
}

var fields_update_quick_response = []leanruntime.Field{
	{Name: "Channels", Flag: "channels", Type: "[]string", Required: false},
	{Name: "Content", Flag: "content", Type: "types.QuickResponseDataProvider", Required: false},
	{Name: "ContentType", Flag: "content-type", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "GroupingConfiguration", Flag: "grouping-configuration", Type: "*types.GroupingConfiguration", Required: false},
	{Name: "IsActive", Flag: "is-active", Type: "*bool", Required: false},
	{Name: "KnowledgeBaseId", Flag: "knowledge-base-id", Type: "*string", Required: true},
	{Name: "Language", Flag: "language", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "QuickResponseId", Flag: "quick-response-id", Type: "*string", Required: true},
	{Name: "RemoveDescription", Flag: "remove-description", Type: "*bool", Required: false},
	{Name: "RemoveGroupingConfiguration", Flag: "remove-grouping-configuration", Type: "*bool", Required: false},
	{Name: "RemoveShortcutKey", Flag: "remove-shortcut-key", Type: "*bool", Required: false},
	{Name: "ShortcutKey", Flag: "shortcut-key", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-assistant": {
			Name:   "create-assistant",
			Fields: fields_create_assistant,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAssistantInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_assistant, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAssistant(ctx, input)
			},
		},
		"create-assistant-association": {
			Name:   "create-assistant-association",
			Fields: fields_create_assistant_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAssistantAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_assistant_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAssistantAssociation(ctx, input)
			},
		},
		"create-content": {
			Name:   "create-content",
			Fields: fields_create_content,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateContentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_content, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateContent(ctx, input)
			},
		},
		"create-knowledge-base": {
			Name:   "create-knowledge-base",
			Fields: fields_create_knowledge_base,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateKnowledgeBaseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_knowledge_base, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateKnowledgeBase(ctx, input)
			},
		},
		"create-quick-response": {
			Name:   "create-quick-response",
			Fields: fields_create_quick_response,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateQuickResponseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_quick_response, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateQuickResponse(ctx, input)
			},
		},
		"create-session": {
			Name:   "create-session",
			Fields: fields_create_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSession(ctx, input)
			},
		},
		"delete-assistant": {
			Name:   "delete-assistant",
			Fields: fields_delete_assistant,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAssistantInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_assistant, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAssistant(ctx, input)
			},
		},
		"delete-assistant-association": {
			Name:   "delete-assistant-association",
			Fields: fields_delete_assistant_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAssistantAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_assistant_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAssistantAssociation(ctx, input)
			},
		},
		"delete-content": {
			Name:   "delete-content",
			Fields: fields_delete_content,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteContentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_content, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteContent(ctx, input)
			},
		},
		"delete-import-job": {
			Name:   "delete-import-job",
			Fields: fields_delete_import_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteImportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_import_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteImportJob(ctx, input)
			},
		},
		"delete-knowledge-base": {
			Name:   "delete-knowledge-base",
			Fields: fields_delete_knowledge_base,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteKnowledgeBaseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_knowledge_base, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteKnowledgeBase(ctx, input)
			},
		},
		"delete-quick-response": {
			Name:   "delete-quick-response",
			Fields: fields_delete_quick_response,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteQuickResponseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_quick_response, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteQuickResponse(ctx, input)
			},
		},
		"get-assistant": {
			Name:   "get-assistant",
			Fields: fields_get_assistant,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAssistantInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_assistant, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAssistant(ctx, input)
			},
		},
		"get-assistant-association": {
			Name:   "get-assistant-association",
			Fields: fields_get_assistant_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAssistantAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_assistant_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAssistantAssociation(ctx, input)
			},
		},
		"get-content": {
			Name:   "get-content",
			Fields: fields_get_content,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetContentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_content, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetContent(ctx, input)
			},
		},
		"get-content-summary": {
			Name:   "get-content-summary",
			Fields: fields_get_content_summary,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetContentSummaryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_content_summary, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetContentSummary(ctx, input)
			},
		},
		"get-import-job": {
			Name:   "get-import-job",
			Fields: fields_get_import_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetImportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_import_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetImportJob(ctx, input)
			},
		},
		"get-knowledge-base": {
			Name:   "get-knowledge-base",
			Fields: fields_get_knowledge_base,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetKnowledgeBaseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_knowledge_base, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetKnowledgeBase(ctx, input)
			},
		},
		"get-quick-response": {
			Name:   "get-quick-response",
			Fields: fields_get_quick_response,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetQuickResponseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_quick_response, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetQuickResponse(ctx, input)
			},
		},
		"get-recommendations": {
			Name:   "get-recommendations",
			Fields: fields_get_recommendations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRecommendationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_recommendations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRecommendations(ctx, input)
			},
		},
		"get-session": {
			Name:   "get-session",
			Fields: fields_get_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSession(ctx, input)
			},
		},
		"list-assistant-associations": {
			Name:   "list-assistant-associations",
			Fields: fields_list_assistant_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAssistantAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_assistant_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAssistantAssociations(ctx, input)
				}
				var results []*svc.ListAssistantAssociationsOutput
				p := svc.NewListAssistantAssociationsPaginator(client, input)
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
		"list-assistants": {
			Name:   "list-assistants",
			Fields: fields_list_assistants,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAssistantsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_assistants, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAssistants(ctx, input)
				}
				var results []*svc.ListAssistantsOutput
				p := svc.NewListAssistantsPaginator(client, input)
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
		"list-contents": {
			Name:   "list-contents",
			Fields: fields_list_contents,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListContentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_contents, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListContents(ctx, input)
				}
				var results []*svc.ListContentsOutput
				p := svc.NewListContentsPaginator(client, input)
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
		"list-import-jobs": {
			Name:   "list-import-jobs",
			Fields: fields_list_import_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListImportJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_import_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListImportJobs(ctx, input)
				}
				var results []*svc.ListImportJobsOutput
				p := svc.NewListImportJobsPaginator(client, input)
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
		"list-knowledge-bases": {
			Name:   "list-knowledge-bases",
			Fields: fields_list_knowledge_bases,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListKnowledgeBasesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_knowledge_bases, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListKnowledgeBases(ctx, input)
				}
				var results []*svc.ListKnowledgeBasesOutput
				p := svc.NewListKnowledgeBasesPaginator(client, input)
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
		"list-quick-responses": {
			Name:   "list-quick-responses",
			Fields: fields_list_quick_responses,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListQuickResponsesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_quick_responses, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListQuickResponses(ctx, input)
				}
				var results []*svc.ListQuickResponsesOutput
				p := svc.NewListQuickResponsesPaginator(client, input)
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
		"notify-recommendations-received": {
			Name:   "notify-recommendations-received",
			Fields: fields_notify_recommendations_received,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.NotifyRecommendationsReceivedInput{}
				if _, err := leanruntime.ApplyInput(input, fields_notify_recommendations_received, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.NotifyRecommendationsReceived(ctx, input)
			},
		},
		"query-assistant": {
			Name:   "query-assistant",
			Fields: fields_query_assistant,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.QueryAssistantInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_query_assistant, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.QueryAssistant(ctx, input)
				}
				var results []*svc.QueryAssistantOutput
				p := svc.NewQueryAssistantPaginator(client, input)
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
		"remove-knowledge-base-template-uri": {
			Name:   "remove-knowledge-base-template-uri",
			Fields: fields_remove_knowledge_base_template_uri,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveKnowledgeBaseTemplateUriInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_knowledge_base_template_uri, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveKnowledgeBaseTemplateUri(ctx, input)
			},
		},
		"search-content": {
			Name:   "search-content",
			Fields: fields_search_content,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchContentInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_content, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchContent(ctx, input)
				}
				var results []*svc.SearchContentOutput
				p := svc.NewSearchContentPaginator(client, input)
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
		"search-quick-responses": {
			Name:   "search-quick-responses",
			Fields: fields_search_quick_responses,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchQuickResponsesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_quick_responses, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchQuickResponses(ctx, input)
				}
				var results []*svc.SearchQuickResponsesOutput
				p := svc.NewSearchQuickResponsesPaginator(client, input)
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
		"search-sessions": {
			Name:   "search-sessions",
			Fields: fields_search_sessions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchSessionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_sessions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchSessions(ctx, input)
				}
				var results []*svc.SearchSessionsOutput
				p := svc.NewSearchSessionsPaginator(client, input)
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
		"start-content-upload": {
			Name:   "start-content-upload",
			Fields: fields_start_content_upload,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartContentUploadInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_content_upload, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartContentUpload(ctx, input)
			},
		},
		"start-import-job": {
			Name:   "start-import-job",
			Fields: fields_start_import_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartImportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_import_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartImportJob(ctx, input)
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
		"update-content": {
			Name:   "update-content",
			Fields: fields_update_content,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateContentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_content, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateContent(ctx, input)
			},
		},
		"update-knowledge-base-template-uri": {
			Name:   "update-knowledge-base-template-uri",
			Fields: fields_update_knowledge_base_template_uri,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateKnowledgeBaseTemplateUriInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_knowledge_base_template_uri, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateKnowledgeBaseTemplateUri(ctx, input)
			},
		},
		"update-quick-response": {
			Name:   "update-quick-response",
			Fields: fields_update_quick_response,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateQuickResponseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_quick_response, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateQuickResponse(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("wisdom", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
