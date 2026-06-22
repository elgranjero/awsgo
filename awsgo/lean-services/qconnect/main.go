package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/qconnect"
)

var fields_activate_message_template = []leanruntime.Field{
	{Name: "KnowledgeBaseId", Flag: "knowledge-base-id", Type: "*string", Required: true},
	{Name: "MessageTemplateId", Flag: "message-template-id", Type: "*string", Required: true},
	{Name: "VersionNumber", Flag: "version-number", Type: "*int64", Required: true},
}

var fields_create_ai_agent = []leanruntime.Field{
	{Name: "AssistantId", Flag: "assistant-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Configuration", Flag: "configuration", Type: "types.AIAgentConfiguration", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.AIAgentType", Required: true},
	{Name: "VisibilityStatus", Flag: "visibility-status", Type: "types.VisibilityStatus", Required: true},
}

var fields_create_ai_agent_version = []leanruntime.Field{
	{Name: "AiAgentId", Flag: "ai-agent-id", Type: "*string", Required: true},
	{Name: "AssistantId", Flag: "assistant-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ModifiedTime", Flag: "modified-time", Type: "*time.Time", Required: false},
}

var fields_create_ai_guardrail = []leanruntime.Field{
	{Name: "AssistantId", Flag: "assistant-id", Type: "*string", Required: true},
	{Name: "BlockedInputMessaging", Flag: "blocked-input-messaging", Type: "*string", Required: true},
	{Name: "BlockedOutputsMessaging", Flag: "blocked-outputs-messaging", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ContentPolicyConfig", Flag: "content-policy-config", Type: "*types.AIGuardrailContentPolicyConfig", Required: false},
	{Name: "ContextualGroundingPolicyConfig", Flag: "contextual-grounding-policy-config", Type: "*types.AIGuardrailContextualGroundingPolicyConfig", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "SensitiveInformationPolicyConfig", Flag: "sensitive-information-policy-config", Type: "*types.AIGuardrailSensitiveInformationPolicyConfig", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "TopicPolicyConfig", Flag: "topic-policy-config", Type: "*types.AIGuardrailTopicPolicyConfig", Required: false},
	{Name: "VisibilityStatus", Flag: "visibility-status", Type: "types.VisibilityStatus", Required: true},
	{Name: "WordPolicyConfig", Flag: "word-policy-config", Type: "*types.AIGuardrailWordPolicyConfig", Required: false},
}

var fields_create_ai_guardrail_version = []leanruntime.Field{
	{Name: "AiGuardrailId", Flag: "ai-guardrail-id", Type: "*string", Required: true},
	{Name: "AssistantId", Flag: "assistant-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ModifiedTime", Flag: "modified-time", Type: "*time.Time", Required: false},
}

var fields_create_ai_prompt = []leanruntime.Field{
	{Name: "ApiFormat", Flag: "api-format", Type: "types.AIPromptAPIFormat", Required: true},
	{Name: "AssistantId", Flag: "assistant-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "InferenceConfiguration", Flag: "inference-configuration", Type: "*types.AIPromptInferenceConfiguration", Required: false},
	{Name: "ModelId", Flag: "model-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "TemplateConfiguration", Flag: "template-configuration", Type: "types.AIPromptTemplateConfiguration", Required: true},
	{Name: "TemplateType", Flag: "template-type", Type: "types.AIPromptTemplateType", Required: true},
	{Name: "Type", Flag: "type", Type: "types.AIPromptType", Required: true},
	{Name: "VisibilityStatus", Flag: "visibility-status", Type: "types.VisibilityStatus", Required: true},
}

var fields_create_ai_prompt_version = []leanruntime.Field{
	{Name: "AiPromptId", Flag: "ai-prompt-id", Type: "*string", Required: true},
	{Name: "AssistantId", Flag: "assistant-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ModifiedTime", Flag: "modified-time", Type: "*time.Time", Required: false},
}

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

var fields_create_content_association = []leanruntime.Field{
	{Name: "Association", Flag: "association", Type: "types.ContentAssociationContents", Required: true},
	{Name: "AssociationType", Flag: "association-type", Type: "types.ContentAssociationType", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ContentId", Flag: "content-id", Type: "*string", Required: true},
	{Name: "KnowledgeBaseId", Flag: "knowledge-base-id", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
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
	{Name: "VectorIngestionConfiguration", Flag: "vector-ingestion-configuration", Type: "*types.VectorIngestionConfiguration", Required: false},
}

var fields_create_message_template = []leanruntime.Field{
	{Name: "ChannelSubtype", Flag: "channel-subtype", Type: "types.ChannelSubtype", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Content", Flag: "content", Type: "types.MessageTemplateContentProvider", Required: false},
	{Name: "DefaultAttributes", Flag: "default-attributes", Type: "*types.MessageTemplateAttributes", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "GroupingConfiguration", Flag: "grouping-configuration", Type: "*types.GroupingConfiguration", Required: false},
	{Name: "KnowledgeBaseId", Flag: "knowledge-base-id", Type: "*string", Required: true},
	{Name: "Language", Flag: "language", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "SourceConfiguration", Flag: "source-configuration", Type: "types.MessageTemplateSourceConfiguration", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_message_template_attachment = []leanruntime.Field{
	{Name: "Body", Flag: "body", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ContentDisposition", Flag: "content-disposition", Type: "types.ContentDisposition", Required: true},
	{Name: "KnowledgeBaseId", Flag: "knowledge-base-id", Type: "*string", Required: true},
	{Name: "MessageTemplateId", Flag: "message-template-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_create_message_template_version = []leanruntime.Field{
	{Name: "KnowledgeBaseId", Flag: "knowledge-base-id", Type: "*string", Required: true},
	{Name: "MessageTemplateContentSha256", Flag: "message-template-content-sha256", Type: "*string", Required: false},
	{Name: "MessageTemplateId", Flag: "message-template-id", Type: "*string", Required: true},
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
	{Name: "AiAgentConfiguration", Flag: "ai-agent-configuration", Type: "map[string]types.AIAgentConfigurationData", Required: false},
	{Name: "AssistantId", Flag: "assistant-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ContactArn", Flag: "contact-arn", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "OrchestratorConfigurationList", Flag: "orchestrator-configuration-list", Type: "[]types.OrchestratorConfigurationEntry", Required: false},
	{Name: "RemoveOrchestratorConfigurationList", Flag: "remove-orchestrator-configuration-list", Type: "*bool", Required: false},
	{Name: "TagFilter", Flag: "tag-filter", Type: "types.TagFilter", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_deactivate_message_template = []leanruntime.Field{
	{Name: "KnowledgeBaseId", Flag: "knowledge-base-id", Type: "*string", Required: true},
	{Name: "MessageTemplateId", Flag: "message-template-id", Type: "*string", Required: true},
	{Name: "VersionNumber", Flag: "version-number", Type: "*int64", Required: true},
}

var fields_delete_ai_agent = []leanruntime.Field{
	{Name: "AiAgentId", Flag: "ai-agent-id", Type: "*string", Required: true},
	{Name: "AssistantId", Flag: "assistant-id", Type: "*string", Required: true},
}

var fields_delete_ai_agent_version = []leanruntime.Field{
	{Name: "AiAgentId", Flag: "ai-agent-id", Type: "*string", Required: true},
	{Name: "AssistantId", Flag: "assistant-id", Type: "*string", Required: true},
	{Name: "VersionNumber", Flag: "version-number", Type: "*int64", Required: true},
}

var fields_delete_ai_guardrail = []leanruntime.Field{
	{Name: "AiGuardrailId", Flag: "ai-guardrail-id", Type: "*string", Required: true},
	{Name: "AssistantId", Flag: "assistant-id", Type: "*string", Required: true},
}

var fields_delete_ai_guardrail_version = []leanruntime.Field{
	{Name: "AiGuardrailId", Flag: "ai-guardrail-id", Type: "*string", Required: true},
	{Name: "AssistantId", Flag: "assistant-id", Type: "*string", Required: true},
	{Name: "VersionNumber", Flag: "version-number", Type: "*int64", Required: true},
}

var fields_delete_ai_prompt = []leanruntime.Field{
	{Name: "AiPromptId", Flag: "ai-prompt-id", Type: "*string", Required: true},
	{Name: "AssistantId", Flag: "assistant-id", Type: "*string", Required: true},
}

var fields_delete_ai_prompt_version = []leanruntime.Field{
	{Name: "AiPromptId", Flag: "ai-prompt-id", Type: "*string", Required: true},
	{Name: "AssistantId", Flag: "assistant-id", Type: "*string", Required: true},
	{Name: "VersionNumber", Flag: "version-number", Type: "*int64", Required: true},
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

var fields_delete_content_association = []leanruntime.Field{
	{Name: "ContentAssociationId", Flag: "content-association-id", Type: "*string", Required: true},
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

var fields_delete_message_template = []leanruntime.Field{
	{Name: "KnowledgeBaseId", Flag: "knowledge-base-id", Type: "*string", Required: true},
	{Name: "MessageTemplateId", Flag: "message-template-id", Type: "*string", Required: true},
}

var fields_delete_message_template_attachment = []leanruntime.Field{
	{Name: "AttachmentId", Flag: "attachment-id", Type: "*string", Required: true},
	{Name: "KnowledgeBaseId", Flag: "knowledge-base-id", Type: "*string", Required: true},
	{Name: "MessageTemplateId", Flag: "message-template-id", Type: "*string", Required: true},
}

var fields_delete_quick_response = []leanruntime.Field{
	{Name: "KnowledgeBaseId", Flag: "knowledge-base-id", Type: "*string", Required: true},
	{Name: "QuickResponseId", Flag: "quick-response-id", Type: "*string", Required: true},
}

var fields_get_ai_agent = []leanruntime.Field{
	{Name: "AiAgentId", Flag: "ai-agent-id", Type: "*string", Required: true},
	{Name: "AssistantId", Flag: "assistant-id", Type: "*string", Required: true},
}

var fields_get_ai_guardrail = []leanruntime.Field{
	{Name: "AiGuardrailId", Flag: "ai-guardrail-id", Type: "*string", Required: true},
	{Name: "AssistantId", Flag: "assistant-id", Type: "*string", Required: true},
}

var fields_get_ai_prompt = []leanruntime.Field{
	{Name: "AiPromptId", Flag: "ai-prompt-id", Type: "*string", Required: true},
	{Name: "AssistantId", Flag: "assistant-id", Type: "*string", Required: true},
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

var fields_get_content_association = []leanruntime.Field{
	{Name: "ContentAssociationId", Flag: "content-association-id", Type: "*string", Required: true},
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

var fields_get_message_template = []leanruntime.Field{
	{Name: "KnowledgeBaseId", Flag: "knowledge-base-id", Type: "*string", Required: true},
	{Name: "MessageTemplateId", Flag: "message-template-id", Type: "*string", Required: true},
}

var fields_get_next_message = []leanruntime.Field{
	{Name: "AssistantId", Flag: "assistant-id", Type: "*string", Required: true},
	{Name: "NextMessageToken", Flag: "next-message-token", Type: "*string", Required: true},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: true},
}

var fields_get_quick_response = []leanruntime.Field{
	{Name: "KnowledgeBaseId", Flag: "knowledge-base-id", Type: "*string", Required: true},
	{Name: "QuickResponseId", Flag: "quick-response-id", Type: "*string", Required: true},
}

var fields_get_recommendations = []leanruntime.Field{
	{Name: "AssistantId", Flag: "assistant-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextChunkToken", Flag: "next-chunk-token", Type: "*string", Required: false},
	{Name: "RecommendationType", Flag: "recommendation-type", Type: "types.RecommendationType", Required: false},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: true},
	{Name: "WaitTimeSeconds", Flag: "wait-time-seconds", Type: "int32", Required: false},
}

var fields_get_session = []leanruntime.Field{
	{Name: "AssistantId", Flag: "assistant-id", Type: "*string", Required: true},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: true},
}

var fields_list_ai_agent_versions = []leanruntime.Field{
	{Name: "AiAgentId", Flag: "ai-agent-id", Type: "*string", Required: true},
	{Name: "AssistantId", Flag: "assistant-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Origin", Flag: "origin", Type: "types.Origin", Required: false},
}

var fields_list_ai_agents = []leanruntime.Field{
	{Name: "AssistantId", Flag: "assistant-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Origin", Flag: "origin", Type: "types.Origin", Required: false},
}

var fields_list_ai_guardrail_versions = []leanruntime.Field{
	{Name: "AiGuardrailId", Flag: "ai-guardrail-id", Type: "*string", Required: true},
	{Name: "AssistantId", Flag: "assistant-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_ai_guardrails = []leanruntime.Field{
	{Name: "AssistantId", Flag: "assistant-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_ai_prompt_versions = []leanruntime.Field{
	{Name: "AiPromptId", Flag: "ai-prompt-id", Type: "*string", Required: true},
	{Name: "AssistantId", Flag: "assistant-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Origin", Flag: "origin", Type: "types.Origin", Required: false},
}

var fields_list_ai_prompts = []leanruntime.Field{
	{Name: "AssistantId", Flag: "assistant-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Origin", Flag: "origin", Type: "types.Origin", Required: false},
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

var fields_list_content_associations = []leanruntime.Field{
	{Name: "ContentId", Flag: "content-id", Type: "*string", Required: true},
	{Name: "KnowledgeBaseId", Flag: "knowledge-base-id", Type: "*string", Required: true},
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

var fields_list_message_template_versions = []leanruntime.Field{
	{Name: "KnowledgeBaseId", Flag: "knowledge-base-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MessageTemplateId", Flag: "message-template-id", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_message_templates = []leanruntime.Field{
	{Name: "KnowledgeBaseId", Flag: "knowledge-base-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_messages = []leanruntime.Field{
	{Name: "AssistantId", Flag: "assistant-id", Type: "*string", Required: true},
	{Name: "Filter", Flag: "filter", Type: "types.MessageFilterType", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: true},
}

var fields_list_quick_responses = []leanruntime.Field{
	{Name: "KnowledgeBaseId", Flag: "knowledge-base-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_spans = []leanruntime.Field{
	{Name: "AssistantId", Flag: "assistant-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: true},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_notify_recommendations_received = []leanruntime.Field{
	{Name: "AssistantId", Flag: "assistant-id", Type: "*string", Required: true},
	{Name: "RecommendationIds", Flag: "recommendation-ids", Type: "[]string", Required: true},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: true},
}

var fields_put_feedback = []leanruntime.Field{
	{Name: "AssistantId", Flag: "assistant-id", Type: "*string", Required: true},
	{Name: "ContentFeedback", Flag: "content-feedback", Type: "types.ContentFeedbackData", Required: true},
	{Name: "TargetId", Flag: "target-id", Type: "*string", Required: true},
	{Name: "TargetType", Flag: "target-type", Type: "types.TargetType", Required: true},
}

var fields_query_assistant = []leanruntime.Field{
	{Name: "AssistantId", Flag: "assistant-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OverrideKnowledgeBaseSearchType", Flag: "override-knowledge-base-search-type", Type: "types.KnowledgeBaseSearchType", Required: false},
	{Name: "QueryCondition", Flag: "query-condition", Type: "[]types.QueryCondition", Required: false},
	{Name: "QueryInputData", Flag: "query-input-data", Type: "types.QueryInputData", Required: false},
	{Name: "QueryText", Flag: "query-text", Type: "*string", Required: false},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: false},
}

var fields_remove_assistant_ai_agent = []leanruntime.Field{
	{Name: "AiAgentType", Flag: "ai-agent-type", Type: "types.AIAgentType", Required: true},
	{Name: "AssistantId", Flag: "assistant-id", Type: "*string", Required: true},
	{Name: "OrchestratorUseCase", Flag: "orchestrator-use-case", Type: "*string", Required: false},
}

var fields_remove_knowledge_base_template_uri = []leanruntime.Field{
	{Name: "KnowledgeBaseId", Flag: "knowledge-base-id", Type: "*string", Required: true},
}

var fields_render_message_template = []leanruntime.Field{
	{Name: "Attributes", Flag: "attributes", Type: "*types.MessageTemplateAttributes", Required: true},
	{Name: "KnowledgeBaseId", Flag: "knowledge-base-id", Type: "*string", Required: true},
	{Name: "MessageTemplateId", Flag: "message-template-id", Type: "*string", Required: true},
}

var fields_retrieve = []leanruntime.Field{
	{Name: "AssistantId", Flag: "assistant-id", Type: "*string", Required: true},
	{Name: "RetrievalConfiguration", Flag: "retrieval-configuration", Type: "*types.RetrievalConfiguration", Required: true},
	{Name: "RetrievalQuery", Flag: "retrieval-query", Type: "*string", Required: true},
}

var fields_search_content = []leanruntime.Field{
	{Name: "KnowledgeBaseId", Flag: "knowledge-base-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SearchExpression", Flag: "search-expression", Type: "*types.SearchExpression", Required: true},
}

var fields_search_message_templates = []leanruntime.Field{
	{Name: "KnowledgeBaseId", Flag: "knowledge-base-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SearchExpression", Flag: "search-expression", Type: "*types.MessageTemplateSearchExpression", Required: true},
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

var fields_send_message = []leanruntime.Field{
	{Name: "AiAgentId", Flag: "ai-agent-id", Type: "*string", Required: false},
	{Name: "AssistantId", Flag: "assistant-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Configuration", Flag: "configuration", Type: "*types.MessageConfiguration", Required: false},
	{Name: "ConversationContext", Flag: "conversation-context", Type: "*types.ConversationContext", Required: false},
	{Name: "Message", Flag: "message", Type: "*types.MessageInput", Required: true},
	{Name: "Metadata", Flag: "metadata", Type: "map[string]string", Required: false},
	{Name: "OrchestratorUseCase", Flag: "orchestrator-use-case", Type: "*string", Required: false},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: true},
	{Name: "Type", Flag: "type", Type: "types.MessageType", Required: true},
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

var fields_update_ai_agent = []leanruntime.Field{
	{Name: "AiAgentId", Flag: "ai-agent-id", Type: "*string", Required: true},
	{Name: "AssistantId", Flag: "assistant-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Configuration", Flag: "configuration", Type: "types.AIAgentConfiguration", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "VisibilityStatus", Flag: "visibility-status", Type: "types.VisibilityStatus", Required: true},
}

var fields_update_ai_guardrail = []leanruntime.Field{
	{Name: "AiGuardrailId", Flag: "ai-guardrail-id", Type: "*string", Required: true},
	{Name: "AssistantId", Flag: "assistant-id", Type: "*string", Required: true},
	{Name: "BlockedInputMessaging", Flag: "blocked-input-messaging", Type: "*string", Required: true},
	{Name: "BlockedOutputsMessaging", Flag: "blocked-outputs-messaging", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ContentPolicyConfig", Flag: "content-policy-config", Type: "*types.AIGuardrailContentPolicyConfig", Required: false},
	{Name: "ContextualGroundingPolicyConfig", Flag: "contextual-grounding-policy-config", Type: "*types.AIGuardrailContextualGroundingPolicyConfig", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "SensitiveInformationPolicyConfig", Flag: "sensitive-information-policy-config", Type: "*types.AIGuardrailSensitiveInformationPolicyConfig", Required: false},
	{Name: "TopicPolicyConfig", Flag: "topic-policy-config", Type: "*types.AIGuardrailTopicPolicyConfig", Required: false},
	{Name: "VisibilityStatus", Flag: "visibility-status", Type: "types.VisibilityStatus", Required: true},
	{Name: "WordPolicyConfig", Flag: "word-policy-config", Type: "*types.AIGuardrailWordPolicyConfig", Required: false},
}

var fields_update_ai_prompt = []leanruntime.Field{
	{Name: "AiPromptId", Flag: "ai-prompt-id", Type: "*string", Required: true},
	{Name: "AssistantId", Flag: "assistant-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "InferenceConfiguration", Flag: "inference-configuration", Type: "*types.AIPromptInferenceConfiguration", Required: false},
	{Name: "ModelId", Flag: "model-id", Type: "*string", Required: false},
	{Name: "TemplateConfiguration", Flag: "template-configuration", Type: "types.AIPromptTemplateConfiguration", Required: false},
	{Name: "VisibilityStatus", Flag: "visibility-status", Type: "types.VisibilityStatus", Required: true},
}

var fields_update_assistant_ai_agent = []leanruntime.Field{
	{Name: "AiAgentType", Flag: "ai-agent-type", Type: "types.AIAgentType", Required: true},
	{Name: "AssistantId", Flag: "assistant-id", Type: "*string", Required: true},
	{Name: "Configuration", Flag: "configuration", Type: "*types.AIAgentConfigurationData", Required: true},
	{Name: "OrchestratorUseCase", Flag: "orchestrator-use-case", Type: "*string", Required: false},
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

var fields_update_message_template = []leanruntime.Field{
	{Name: "Content", Flag: "content", Type: "types.MessageTemplateContentProvider", Required: false},
	{Name: "DefaultAttributes", Flag: "default-attributes", Type: "*types.MessageTemplateAttributes", Required: false},
	{Name: "KnowledgeBaseId", Flag: "knowledge-base-id", Type: "*string", Required: true},
	{Name: "Language", Flag: "language", Type: "*string", Required: false},
	{Name: "MessageTemplateId", Flag: "message-template-id", Type: "*string", Required: true},
	{Name: "SourceConfiguration", Flag: "source-configuration", Type: "types.MessageTemplateSourceConfiguration", Required: false},
}

var fields_update_message_template_metadata = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "GroupingConfiguration", Flag: "grouping-configuration", Type: "*types.GroupingConfiguration", Required: false},
	{Name: "KnowledgeBaseId", Flag: "knowledge-base-id", Type: "*string", Required: true},
	{Name: "MessageTemplateId", Flag: "message-template-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
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

var fields_update_session = []leanruntime.Field{
	{Name: "AiAgentConfiguration", Flag: "ai-agent-configuration", Type: "map[string]types.AIAgentConfigurationData", Required: false},
	{Name: "AssistantId", Flag: "assistant-id", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "OrchestratorConfigurationList", Flag: "orchestrator-configuration-list", Type: "[]types.OrchestratorConfigurationEntry", Required: false},
	{Name: "RemoveOrchestratorConfigurationList", Flag: "remove-orchestrator-configuration-list", Type: "*bool", Required: false},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: true},
	{Name: "TagFilter", Flag: "tag-filter", Type: "types.TagFilter", Required: false},
}

var fields_update_session_data = []leanruntime.Field{
	{Name: "AssistantId", Flag: "assistant-id", Type: "*string", Required: true},
	{Name: "Data", Flag: "data", Type: "[]types.RuntimeSessionData", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "types.SessionDataNamespace", Required: false},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"activate-message-template": {
			Name:   "activate-message-template",
			Fields: fields_activate_message_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ActivateMessageTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_activate_message_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ActivateMessageTemplate(ctx, input)
			},
		},
		"create-ai-agent": {
			Name:   "create-ai-agent",
			Fields: fields_create_ai_agent,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAIAgentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_ai_agent, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAIAgent(ctx, input)
			},
		},
		"create-ai-agent-version": {
			Name:   "create-ai-agent-version",
			Fields: fields_create_ai_agent_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAIAgentVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_ai_agent_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAIAgentVersion(ctx, input)
			},
		},
		"create-ai-guardrail": {
			Name:   "create-ai-guardrail",
			Fields: fields_create_ai_guardrail,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAIGuardrailInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_ai_guardrail, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAIGuardrail(ctx, input)
			},
		},
		"create-ai-guardrail-version": {
			Name:   "create-ai-guardrail-version",
			Fields: fields_create_ai_guardrail_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAIGuardrailVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_ai_guardrail_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAIGuardrailVersion(ctx, input)
			},
		},
		"create-ai-prompt": {
			Name:   "create-ai-prompt",
			Fields: fields_create_ai_prompt,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAIPromptInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_ai_prompt, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAIPrompt(ctx, input)
			},
		},
		"create-ai-prompt-version": {
			Name:   "create-ai-prompt-version",
			Fields: fields_create_ai_prompt_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAIPromptVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_ai_prompt_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAIPromptVersion(ctx, input)
			},
		},
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
		"create-content-association": {
			Name:   "create-content-association",
			Fields: fields_create_content_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateContentAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_content_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateContentAssociation(ctx, input)
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
		"create-message-template": {
			Name:   "create-message-template",
			Fields: fields_create_message_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateMessageTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_message_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateMessageTemplate(ctx, input)
			},
		},
		"create-message-template-attachment": {
			Name:   "create-message-template-attachment",
			Fields: fields_create_message_template_attachment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateMessageTemplateAttachmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_message_template_attachment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateMessageTemplateAttachment(ctx, input)
			},
		},
		"create-message-template-version": {
			Name:   "create-message-template-version",
			Fields: fields_create_message_template_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateMessageTemplateVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_message_template_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateMessageTemplateVersion(ctx, input)
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
		"deactivate-message-template": {
			Name:   "deactivate-message-template",
			Fields: fields_deactivate_message_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeactivateMessageTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deactivate_message_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeactivateMessageTemplate(ctx, input)
			},
		},
		"delete-ai-agent": {
			Name:   "delete-ai-agent",
			Fields: fields_delete_ai_agent,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAIAgentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_ai_agent, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAIAgent(ctx, input)
			},
		},
		"delete-ai-agent-version": {
			Name:   "delete-ai-agent-version",
			Fields: fields_delete_ai_agent_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAIAgentVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_ai_agent_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAIAgentVersion(ctx, input)
			},
		},
		"delete-ai-guardrail": {
			Name:   "delete-ai-guardrail",
			Fields: fields_delete_ai_guardrail,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAIGuardrailInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_ai_guardrail, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAIGuardrail(ctx, input)
			},
		},
		"delete-ai-guardrail-version": {
			Name:   "delete-ai-guardrail-version",
			Fields: fields_delete_ai_guardrail_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAIGuardrailVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_ai_guardrail_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAIGuardrailVersion(ctx, input)
			},
		},
		"delete-ai-prompt": {
			Name:   "delete-ai-prompt",
			Fields: fields_delete_ai_prompt,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAIPromptInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_ai_prompt, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAIPrompt(ctx, input)
			},
		},
		"delete-ai-prompt-version": {
			Name:   "delete-ai-prompt-version",
			Fields: fields_delete_ai_prompt_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAIPromptVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_ai_prompt_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAIPromptVersion(ctx, input)
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
		"delete-content-association": {
			Name:   "delete-content-association",
			Fields: fields_delete_content_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteContentAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_content_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteContentAssociation(ctx, input)
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
		"delete-message-template": {
			Name:   "delete-message-template",
			Fields: fields_delete_message_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMessageTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_message_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMessageTemplate(ctx, input)
			},
		},
		"delete-message-template-attachment": {
			Name:   "delete-message-template-attachment",
			Fields: fields_delete_message_template_attachment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMessageTemplateAttachmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_message_template_attachment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMessageTemplateAttachment(ctx, input)
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
		"get-ai-agent": {
			Name:   "get-ai-agent",
			Fields: fields_get_ai_agent,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAIAgentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_ai_agent, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAIAgent(ctx, input)
			},
		},
		"get-ai-guardrail": {
			Name:   "get-ai-guardrail",
			Fields: fields_get_ai_guardrail,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAIGuardrailInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_ai_guardrail, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAIGuardrail(ctx, input)
			},
		},
		"get-ai-prompt": {
			Name:   "get-ai-prompt",
			Fields: fields_get_ai_prompt,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAIPromptInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_ai_prompt, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAIPrompt(ctx, input)
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
		"get-content-association": {
			Name:   "get-content-association",
			Fields: fields_get_content_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetContentAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_content_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetContentAssociation(ctx, input)
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
		"get-message-template": {
			Name:   "get-message-template",
			Fields: fields_get_message_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMessageTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_message_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMessageTemplate(ctx, input)
			},
		},
		"get-next-message": {
			Name:   "get-next-message",
			Fields: fields_get_next_message,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetNextMessageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_next_message, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetNextMessage(ctx, input)
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
		"list-ai-agent-versions": {
			Name:   "list-ai-agent-versions",
			Fields: fields_list_ai_agent_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAIAgentVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_ai_agent_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAIAgentVersions(ctx, input)
				}
				var results []*svc.ListAIAgentVersionsOutput
				p := svc.NewListAIAgentVersionsPaginator(client, input)
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
		"list-ai-agents": {
			Name:   "list-ai-agents",
			Fields: fields_list_ai_agents,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAIAgentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_ai_agents, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAIAgents(ctx, input)
				}
				var results []*svc.ListAIAgentsOutput
				p := svc.NewListAIAgentsPaginator(client, input)
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
		"list-ai-guardrail-versions": {
			Name:   "list-ai-guardrail-versions",
			Fields: fields_list_ai_guardrail_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAIGuardrailVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_ai_guardrail_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAIGuardrailVersions(ctx, input)
				}
				var results []*svc.ListAIGuardrailVersionsOutput
				p := svc.NewListAIGuardrailVersionsPaginator(client, input)
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
		"list-ai-guardrails": {
			Name:   "list-ai-guardrails",
			Fields: fields_list_ai_guardrails,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAIGuardrailsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_ai_guardrails, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAIGuardrails(ctx, input)
				}
				var results []*svc.ListAIGuardrailsOutput
				p := svc.NewListAIGuardrailsPaginator(client, input)
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
		"list-ai-prompt-versions": {
			Name:   "list-ai-prompt-versions",
			Fields: fields_list_ai_prompt_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAIPromptVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_ai_prompt_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAIPromptVersions(ctx, input)
				}
				var results []*svc.ListAIPromptVersionsOutput
				p := svc.NewListAIPromptVersionsPaginator(client, input)
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
		"list-ai-prompts": {
			Name:   "list-ai-prompts",
			Fields: fields_list_ai_prompts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAIPromptsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_ai_prompts, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAIPrompts(ctx, input)
				}
				var results []*svc.ListAIPromptsOutput
				p := svc.NewListAIPromptsPaginator(client, input)
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
		"list-content-associations": {
			Name:   "list-content-associations",
			Fields: fields_list_content_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListContentAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_content_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListContentAssociations(ctx, input)
				}
				var results []*svc.ListContentAssociationsOutput
				p := svc.NewListContentAssociationsPaginator(client, input)
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
		"list-message-template-versions": {
			Name:   "list-message-template-versions",
			Fields: fields_list_message_template_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMessageTemplateVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_message_template_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMessageTemplateVersions(ctx, input)
				}
				var results []*svc.ListMessageTemplateVersionsOutput
				p := svc.NewListMessageTemplateVersionsPaginator(client, input)
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
		"list-message-templates": {
			Name:   "list-message-templates",
			Fields: fields_list_message_templates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMessageTemplatesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_message_templates, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMessageTemplates(ctx, input)
				}
				var results []*svc.ListMessageTemplatesOutput
				p := svc.NewListMessageTemplatesPaginator(client, input)
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
		"list-messages": {
			Name:   "list-messages",
			Fields: fields_list_messages,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMessagesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_messages, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMessages(ctx, input)
				}
				var results []*svc.ListMessagesOutput
				p := svc.NewListMessagesPaginator(client, input)
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
		"list-spans": {
			Name:   "list-spans",
			Fields: fields_list_spans,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSpansInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_spans, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSpans(ctx, input)
				}
				var results []*svc.ListSpansOutput
				p := svc.NewListSpansPaginator(client, input)
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
		"put-feedback": {
			Name:   "put-feedback",
			Fields: fields_put_feedback,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutFeedbackInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_feedback, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutFeedback(ctx, input)
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
		"remove-assistant-ai-agent": {
			Name:   "remove-assistant-ai-agent",
			Fields: fields_remove_assistant_ai_agent,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveAssistantAIAgentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_assistant_ai_agent, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveAssistantAIAgent(ctx, input)
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
		"render-message-template": {
			Name:   "render-message-template",
			Fields: fields_render_message_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RenderMessageTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_render_message_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RenderMessageTemplate(ctx, input)
			},
		},
		"retrieve": {
			Name:   "retrieve",
			Fields: fields_retrieve,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RetrieveInput{}
				if _, err := leanruntime.ApplyInput(input, fields_retrieve, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.Retrieve(ctx, input)
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
		"search-message-templates": {
			Name:   "search-message-templates",
			Fields: fields_search_message_templates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchMessageTemplatesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_message_templates, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchMessageTemplates(ctx, input)
				}
				var results []*svc.SearchMessageTemplatesOutput
				p := svc.NewSearchMessageTemplatesPaginator(client, input)
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
		"send-message": {
			Name:   "send-message",
			Fields: fields_send_message,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SendMessageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_send_message, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SendMessage(ctx, input)
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
		"update-ai-agent": {
			Name:   "update-ai-agent",
			Fields: fields_update_ai_agent,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAIAgentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_ai_agent, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAIAgent(ctx, input)
			},
		},
		"update-ai-guardrail": {
			Name:   "update-ai-guardrail",
			Fields: fields_update_ai_guardrail,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAIGuardrailInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_ai_guardrail, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAIGuardrail(ctx, input)
			},
		},
		"update-ai-prompt": {
			Name:   "update-ai-prompt",
			Fields: fields_update_ai_prompt,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAIPromptInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_ai_prompt, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAIPrompt(ctx, input)
			},
		},
		"update-assistant-ai-agent": {
			Name:   "update-assistant-ai-agent",
			Fields: fields_update_assistant_ai_agent,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAssistantAIAgentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_assistant_ai_agent, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAssistantAIAgent(ctx, input)
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
		"update-message-template": {
			Name:   "update-message-template",
			Fields: fields_update_message_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateMessageTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_message_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateMessageTemplate(ctx, input)
			},
		},
		"update-message-template-metadata": {
			Name:   "update-message-template-metadata",
			Fields: fields_update_message_template_metadata,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateMessageTemplateMetadataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_message_template_metadata, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateMessageTemplateMetadata(ctx, input)
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
		"update-session": {
			Name:   "update-session",
			Fields: fields_update_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSession(ctx, input)
			},
		},
		"update-session-data": {
			Name:   "update-session-data",
			Fields: fields_update_session_data,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSessionDataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_session_data, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSessionData(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("qconnect", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
