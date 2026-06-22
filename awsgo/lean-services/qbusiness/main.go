package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/qbusiness"
)

var fields_associate_permission = []leanruntime.Field{
	{Name: "Actions", Flag: "actions", Type: "[]string", Required: true},
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "Conditions", Flag: "conditions", Type: "[]types.PermissionCondition", Required: false},
	{Name: "Principal", Flag: "principal", Type: "*string", Required: true},
	{Name: "StatementId", Flag: "statement-id", Type: "*string", Required: true},
}

var fields_batch_delete_document = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "DataSourceSyncId", Flag: "data-source-sync-id", Type: "*string", Required: false},
	{Name: "Documents", Flag: "documents", Type: "[]types.DeleteDocument", Required: true},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
}

var fields_batch_put_document = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "DataSourceSyncId", Flag: "data-source-sync-id", Type: "*string", Required: false},
	{Name: "Documents", Flag: "documents", Type: "[]types.Document", Required: true},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
}

var fields_cancel_subscription = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "SubscriptionId", Flag: "subscription-id", Type: "*string", Required: true},
}

var fields_chat = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ConversationId", Flag: "conversation-id", Type: "*string", Required: false},
	{Name: "ParentMessageId", Flag: "parent-message-id", Type: "*string", Required: false},
	{Name: "UserGroups", Flag: "user-groups", Type: "[]string", Required: false},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: false},
}

var fields_chat_sync = []leanruntime.Field{
	{Name: "ActionExecution", Flag: "action-execution", Type: "*types.ActionExecution", Required: false},
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "Attachments", Flag: "attachments", Type: "[]types.AttachmentInput", Required: false},
	{Name: "AttributeFilter", Flag: "attribute-filter", Type: "*types.AttributeFilter", Required: false},
	{Name: "AuthChallengeResponse", Flag: "auth-challenge-response", Type: "*types.AuthChallengeResponse", Required: false},
	{Name: "ChatMode", Flag: "chat-mode", Type: "types.ChatMode", Required: false},
	{Name: "ChatModeConfiguration", Flag: "chat-mode-configuration", Type: "types.ChatModeConfiguration", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ConversationId", Flag: "conversation-id", Type: "*string", Required: false},
	{Name: "ParentMessageId", Flag: "parent-message-id", Type: "*string", Required: false},
	{Name: "UserGroups", Flag: "user-groups", Type: "[]string", Required: false},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: false},
	{Name: "UserMessage", Flag: "user-message", Type: "*string", Required: false},
}

var fields_check_document_access = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "DataSourceId", Flag: "data-source-id", Type: "*string", Required: false},
	{Name: "DocumentId", Flag: "document-id", Type: "*string", Required: true},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_create_anonymous_web_experience_url = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "SessionDurationInMinutes", Flag: "session-duration-in-minutes", Type: "*int32", Required: false},
	{Name: "WebExperienceId", Flag: "web-experience-id", Type: "*string", Required: true},
}

var fields_create_application = []leanruntime.Field{
	{Name: "AttachmentsConfiguration", Flag: "attachments-configuration", Type: "*types.AttachmentsConfiguration", Required: false},
	{Name: "ClientIdsForOIDC", Flag: "client-ids-for-oidc", Type: "[]string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: true},
	{Name: "EncryptionConfiguration", Flag: "encryption-configuration", Type: "*types.EncryptionConfiguration", Required: false},
	{Name: "IamIdentityProviderArn", Flag: "iam-identity-provider-arn", Type: "*string", Required: false},
	{Name: "IdentityCenterInstanceArn", Flag: "identity-center-instance-arn", Type: "*string", Required: false},
	{Name: "IdentityType", Flag: "identity-type", Type: "types.IdentityType", Required: false},
	{Name: "PersonalizationConfiguration", Flag: "personalization-configuration", Type: "*types.PersonalizationConfiguration", Required: false},
	{Name: "QAppsConfiguration", Flag: "qapps-configuration", Type: "*types.QAppsConfiguration", Required: false},
	{Name: "QuickSightConfiguration", Flag: "quicksight-configuration", Type: "*types.QuickSightConfiguration", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_chat_response_configuration = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: true},
	{Name: "ResponseConfigurations", Flag: "response-configurations", Type: "map[string]types.ResponseConfiguration", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_data_accessor = []leanruntime.Field{
	{Name: "ActionConfigurations", Flag: "action-configurations", Type: "[]types.ActionConfiguration", Required: true},
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "AuthenticationDetail", Flag: "authentication-detail", Type: "*types.DataAccessorAuthenticationDetail", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: true},
	{Name: "Principal", Flag: "principal", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_data_source = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Configuration", Flag: "configuration", Type: "document.Interface", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: true},
	{Name: "DocumentEnrichmentConfiguration", Flag: "document-enrichment-configuration", Type: "*types.DocumentEnrichmentConfiguration", Required: false},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
	{Name: "MediaExtractionConfiguration", Flag: "media-extraction-configuration", Type: "*types.MediaExtractionConfiguration", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "SyncSchedule", Flag: "sync-schedule", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "VpcConfiguration", Flag: "vpc-configuration", Type: "*types.DataSourceVpcConfiguration", Required: false},
}

var fields_create_index = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "CapacityConfiguration", Flag: "capacity-configuration", Type: "*types.IndexCapacityConfiguration", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "Type", Flag: "type", Type: "types.IndexType", Required: false},
}

var fields_create_plugin = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "AuthConfiguration", Flag: "auth-configuration", Type: "types.PluginAuthConfiguration", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "CustomPluginConfiguration", Flag: "custom-plugin-configuration", Type: "*types.CustomPluginConfiguration", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: true},
	{Name: "ServerUrl", Flag: "server-url", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "Type", Flag: "type", Type: "types.PluginType", Required: true},
}

var fields_create_retriever = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Configuration", Flag: "configuration", Type: "types.RetrieverConfiguration", Required: true},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "Type", Flag: "type", Type: "types.RetrieverType", Required: true},
}

var fields_create_subscription = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Principal", Flag: "principal", Type: "types.SubscriptionPrincipal", Required: true},
	{Name: "Type", Flag: "type", Type: "types.SubscriptionType", Required: true},
}

var fields_create_user = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "UserAliases", Flag: "user-aliases", Type: "[]types.UserAlias", Required: false},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_create_web_experience = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "BrowserExtensionConfiguration", Flag: "browser-extension-configuration", Type: "*types.BrowserExtensionConfiguration", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "CustomizationConfiguration", Flag: "customization-configuration", Type: "*types.CustomizationConfiguration", Required: false},
	{Name: "IdentityProviderConfiguration", Flag: "identity-provider-configuration", Type: "types.IdentityProviderConfiguration", Required: false},
	{Name: "Origins", Flag: "origins", Type: "[]string", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "SamplePromptsControlMode", Flag: "sample-prompts-control-mode", Type: "types.WebExperienceSamplePromptsControlMode", Required: false},
	{Name: "Subtitle", Flag: "subtitle", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "Title", Flag: "title", Type: "*string", Required: false},
	{Name: "WelcomeMessage", Flag: "welcome-message", Type: "*string", Required: false},
}

var fields_delete_application = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
}

var fields_delete_attachment = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "AttachmentId", Flag: "attachment-id", Type: "*string", Required: true},
	{Name: "ConversationId", Flag: "conversation-id", Type: "*string", Required: true},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: false},
}

var fields_delete_chat_controls_configuration = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
}

var fields_delete_chat_response_configuration = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "ChatResponseConfigurationId", Flag: "chat-response-configuration-id", Type: "*string", Required: true},
}

var fields_delete_conversation = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "ConversationId", Flag: "conversation-id", Type: "*string", Required: true},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: false},
}

var fields_delete_data_accessor = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "DataAccessorId", Flag: "data-accessor-id", Type: "*string", Required: true},
}

var fields_delete_data_source = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "DataSourceId", Flag: "data-source-id", Type: "*string", Required: true},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
}

var fields_delete_group = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "DataSourceId", Flag: "data-source-id", Type: "*string", Required: false},
	{Name: "GroupName", Flag: "group-name", Type: "*string", Required: true},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
}

var fields_delete_index = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
}

var fields_delete_plugin = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "PluginId", Flag: "plugin-id", Type: "*string", Required: true},
}

var fields_delete_retriever = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "RetrieverId", Flag: "retriever-id", Type: "*string", Required: true},
}

var fields_delete_user = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_delete_web_experience = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "WebExperienceId", Flag: "web-experience-id", Type: "*string", Required: true},
}

var fields_disassociate_permission = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "StatementId", Flag: "statement-id", Type: "*string", Required: true},
}

var fields_get_application = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
}

var fields_get_chat_controls_configuration = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_chat_response_configuration = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "ChatResponseConfigurationId", Flag: "chat-response-configuration-id", Type: "*string", Required: true},
}

var fields_get_data_accessor = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "DataAccessorId", Flag: "data-accessor-id", Type: "*string", Required: true},
}

var fields_get_data_source = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "DataSourceId", Flag: "data-source-id", Type: "*string", Required: true},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
}

var fields_get_document_content = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "DataSourceId", Flag: "data-source-id", Type: "*string", Required: false},
	{Name: "DocumentId", Flag: "document-id", Type: "*string", Required: true},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
	{Name: "OutputFormat", Flag: "output-format", Type: "types.OutputFormat", Required: false},
}

var fields_get_group = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "DataSourceId", Flag: "data-source-id", Type: "*string", Required: false},
	{Name: "GroupName", Flag: "group-name", Type: "*string", Required: true},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
}

var fields_get_index = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
}

var fields_get_media = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "ConversationId", Flag: "conversation-id", Type: "*string", Required: true},
	{Name: "MediaId", Flag: "media-id", Type: "*string", Required: true},
	{Name: "MessageId", Flag: "message-id", Type: "*string", Required: true},
}

var fields_get_plugin = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "PluginId", Flag: "plugin-id", Type: "*string", Required: true},
}

var fields_get_policy = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
}

var fields_get_retriever = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "RetrieverId", Flag: "retriever-id", Type: "*string", Required: true},
}

var fields_get_user = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_get_web_experience = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "WebExperienceId", Flag: "web-experience-id", Type: "*string", Required: true},
}

var fields_list_applications = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_attachments = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "ConversationId", Flag: "conversation-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: false},
}

var fields_list_chat_response_configurations = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_conversations = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: false},
}

var fields_list_data_accessors = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_data_source_sync_jobs = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "DataSourceId", Flag: "data-source-id", Type: "*string", Required: true},
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: false},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: false},
	{Name: "StatusFilter", Flag: "status-filter", Type: "types.DataSourceSyncJobStatus", Required: false},
}

var fields_list_data_sources = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_documents = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "DataSourceIds", Flag: "data-source-ids", Type: "[]string", Required: false},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_groups = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "DataSourceId", Flag: "data-source-id", Type: "*string", Required: false},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "UpdatedEarlierThan", Flag: "updated-earlier-than", Type: "*time.Time", Required: true},
}

var fields_list_indices = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_messages = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "ConversationId", Flag: "conversation-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: false},
}

var fields_list_plugin_actions = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PluginId", Flag: "plugin-id", Type: "*string", Required: true},
}

var fields_list_plugin_type_actions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PluginType", Flag: "plugin-type", Type: "types.PluginType", Required: true},
}

var fields_list_plugin_type_metadata = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_plugins = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_retrievers = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_subscriptions = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_web_experiences = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_put_feedback = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "ConversationId", Flag: "conversation-id", Type: "*string", Required: true},
	{Name: "MessageCopiedAt", Flag: "message-copied-at", Type: "*time.Time", Required: false},
	{Name: "MessageId", Flag: "message-id", Type: "*string", Required: true},
	{Name: "MessageUsefulness", Flag: "message-usefulness", Type: "*types.MessageUsefulnessFeedback", Required: false},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: false},
}

var fields_put_group = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "DataSourceId", Flag: "data-source-id", Type: "*string", Required: false},
	{Name: "GroupMembers", Flag: "group-members", Type: "*types.GroupMembers", Required: true},
	{Name: "GroupName", Flag: "group-name", Type: "*string", Required: true},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.MembershipType", Required: true},
}

var fields_search_relevant_content = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "AttributeFilter", Flag: "attribute-filter", Type: "*types.AttributeFilter", Required: false},
	{Name: "ContentSource", Flag: "content-source", Type: "types.ContentSource", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "QueryText", Flag: "query-text", Type: "*string", Required: true},
}

var fields_start_data_source_sync_job = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "DataSourceId", Flag: "data-source-id", Type: "*string", Required: true},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
}

var fields_stop_data_source_sync_job = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "DataSourceId", Flag: "data-source-id", Type: "*string", Required: true},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_application = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "AttachmentsConfiguration", Flag: "attachments-configuration", Type: "*types.AttachmentsConfiguration", Required: false},
	{Name: "AutoSubscriptionConfiguration", Flag: "auto-subscription-configuration", Type: "*types.AutoSubscriptionConfiguration", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "IdentityCenterInstanceArn", Flag: "identity-center-instance-arn", Type: "*string", Required: false},
	{Name: "PersonalizationConfiguration", Flag: "personalization-configuration", Type: "*types.PersonalizationConfiguration", Required: false},
	{Name: "QAppsConfiguration", Flag: "qapps-configuration", Type: "*types.QAppsConfiguration", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
}

var fields_update_chat_controls_configuration = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "BlockedPhrasesConfigurationUpdate", Flag: "blocked-phrases-configuration-update", Type: "*types.BlockedPhrasesConfigurationUpdate", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "CreatorModeConfiguration", Flag: "creator-mode-configuration", Type: "*types.CreatorModeConfiguration", Required: false},
	{Name: "HallucinationReductionConfiguration", Flag: "hallucination-reduction-configuration", Type: "*types.HallucinationReductionConfiguration", Required: false},
	{Name: "OrchestrationConfiguration", Flag: "orchestration-configuration", Type: "*types.OrchestrationConfiguration", Required: false},
	{Name: "ResponseScope", Flag: "response-scope", Type: "types.ResponseScope", Required: false},
	{Name: "TopicConfigurationsToCreateOrUpdate", Flag: "topic-configurations-to-create-or-update", Type: "[]types.TopicConfiguration", Required: false},
	{Name: "TopicConfigurationsToDelete", Flag: "topic-configurations-to-delete", Type: "[]types.TopicConfiguration", Required: false},
}

var fields_update_chat_response_configuration = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "ChatResponseConfigurationId", Flag: "chat-response-configuration-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "ResponseConfigurations", Flag: "response-configurations", Type: "map[string]types.ResponseConfiguration", Required: true},
}

var fields_update_data_accessor = []leanruntime.Field{
	{Name: "ActionConfigurations", Flag: "action-configurations", Type: "[]types.ActionConfiguration", Required: true},
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "AuthenticationDetail", Flag: "authentication-detail", Type: "*types.DataAccessorAuthenticationDetail", Required: false},
	{Name: "DataAccessorId", Flag: "data-accessor-id", Type: "*string", Required: true},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
}

var fields_update_data_source = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "Configuration", Flag: "configuration", Type: "document.Interface", Required: false},
	{Name: "DataSourceId", Flag: "data-source-id", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "DocumentEnrichmentConfiguration", Flag: "document-enrichment-configuration", Type: "*types.DocumentEnrichmentConfiguration", Required: false},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
	{Name: "MediaExtractionConfiguration", Flag: "media-extraction-configuration", Type: "*types.MediaExtractionConfiguration", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "SyncSchedule", Flag: "sync-schedule", Type: "*string", Required: false},
	{Name: "VpcConfiguration", Flag: "vpc-configuration", Type: "*types.DataSourceVpcConfiguration", Required: false},
}

var fields_update_index = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "CapacityConfiguration", Flag: "capacity-configuration", Type: "*types.IndexCapacityConfiguration", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "DocumentAttributeConfigurations", Flag: "document-attribute-configurations", Type: "[]types.DocumentAttributeConfiguration", Required: false},
	{Name: "IndexId", Flag: "index-id", Type: "*string", Required: true},
}

var fields_update_plugin = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "AuthConfiguration", Flag: "auth-configuration", Type: "types.PluginAuthConfiguration", Required: false},
	{Name: "CustomPluginConfiguration", Flag: "custom-plugin-configuration", Type: "*types.CustomPluginConfiguration", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "PluginId", Flag: "plugin-id", Type: "*string", Required: true},
	{Name: "ServerUrl", Flag: "server-url", Type: "*string", Required: false},
	{Name: "State", Flag: "state", Type: "types.PluginState", Required: false},
}

var fields_update_retriever = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "Configuration", Flag: "configuration", Type: "types.RetrieverConfiguration", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "RetrieverId", Flag: "retriever-id", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
}

var fields_update_subscription = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "SubscriptionId", Flag: "subscription-id", Type: "*string", Required: true},
	{Name: "Type", Flag: "type", Type: "types.SubscriptionType", Required: true},
}

var fields_update_user = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "UserAliasesToDelete", Flag: "user-aliases-to-delete", Type: "[]types.UserAlias", Required: false},
	{Name: "UserAliasesToUpdate", Flag: "user-aliases-to-update", Type: "[]types.UserAlias", Required: false},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_update_web_experience = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "AuthenticationConfiguration", Flag: "authentication-configuration", Type: "types.WebExperienceAuthConfiguration", Required: false},
	{Name: "BrowserExtensionConfiguration", Flag: "browser-extension-configuration", Type: "*types.BrowserExtensionConfiguration", Required: false},
	{Name: "CustomizationConfiguration", Flag: "customization-configuration", Type: "*types.CustomizationConfiguration", Required: false},
	{Name: "IdentityProviderConfiguration", Flag: "identity-provider-configuration", Type: "types.IdentityProviderConfiguration", Required: false},
	{Name: "Origins", Flag: "origins", Type: "[]string", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "SamplePromptsControlMode", Flag: "sample-prompts-control-mode", Type: "types.WebExperienceSamplePromptsControlMode", Required: false},
	{Name: "Subtitle", Flag: "subtitle", Type: "*string", Required: false},
	{Name: "Title", Flag: "title", Type: "*string", Required: false},
	{Name: "WebExperienceId", Flag: "web-experience-id", Type: "*string", Required: true},
	{Name: "WelcomeMessage", Flag: "welcome-message", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"associate-permission": {
			Name:   "associate-permission",
			Fields: fields_associate_permission,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociatePermissionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_permission, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociatePermission(ctx, input)
			},
		},
		"batch-delete-document": {
			Name:   "batch-delete-document",
			Fields: fields_batch_delete_document,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchDeleteDocumentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_delete_document, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchDeleteDocument(ctx, input)
			},
		},
		"batch-put-document": {
			Name:   "batch-put-document",
			Fields: fields_batch_put_document,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchPutDocumentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_put_document, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchPutDocument(ctx, input)
			},
		},
		"cancel-subscription": {
			Name:   "cancel-subscription",
			Fields: fields_cancel_subscription,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelSubscriptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_subscription, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelSubscription(ctx, input)
			},
		},
		"chat": {
			Name:   "chat",
			Fields: fields_chat,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ChatInput{}
				if _, err := leanruntime.ApplyInput(input, fields_chat, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.Chat(ctx, input)
			},
		},
		"chat-sync": {
			Name:   "chat-sync",
			Fields: fields_chat_sync,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ChatSyncInput{}
				if _, err := leanruntime.ApplyInput(input, fields_chat_sync, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ChatSync(ctx, input)
			},
		},
		"check-document-access": {
			Name:   "check-document-access",
			Fields: fields_check_document_access,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CheckDocumentAccessInput{}
				if _, err := leanruntime.ApplyInput(input, fields_check_document_access, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CheckDocumentAccess(ctx, input)
			},
		},
		"create-anonymous-web-experience-url": {
			Name:   "create-anonymous-web-experience-url",
			Fields: fields_create_anonymous_web_experience_url,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAnonymousWebExperienceUrlInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_anonymous_web_experience_url, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAnonymousWebExperienceUrl(ctx, input)
			},
		},
		"create-application": {
			Name:   "create-application",
			Fields: fields_create_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateApplication(ctx, input)
			},
		},
		"create-chat-response-configuration": {
			Name:   "create-chat-response-configuration",
			Fields: fields_create_chat_response_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateChatResponseConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_chat_response_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateChatResponseConfiguration(ctx, input)
			},
		},
		"create-data-accessor": {
			Name:   "create-data-accessor",
			Fields: fields_create_data_accessor,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDataAccessorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_data_accessor, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDataAccessor(ctx, input)
			},
		},
		"create-data-source": {
			Name:   "create-data-source",
			Fields: fields_create_data_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDataSourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_data_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDataSource(ctx, input)
			},
		},
		"create-index": {
			Name:   "create-index",
			Fields: fields_create_index,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateIndexInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_index, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateIndex(ctx, input)
			},
		},
		"create-plugin": {
			Name:   "create-plugin",
			Fields: fields_create_plugin,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePluginInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_plugin, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePlugin(ctx, input)
			},
		},
		"create-retriever": {
			Name:   "create-retriever",
			Fields: fields_create_retriever,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRetrieverInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_retriever, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRetriever(ctx, input)
			},
		},
		"create-subscription": {
			Name:   "create-subscription",
			Fields: fields_create_subscription,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSubscriptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_subscription, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSubscription(ctx, input)
			},
		},
		"create-user": {
			Name:   "create-user",
			Fields: fields_create_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateUser(ctx, input)
			},
		},
		"create-web-experience": {
			Name:   "create-web-experience",
			Fields: fields_create_web_experience,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateWebExperienceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_web_experience, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateWebExperience(ctx, input)
			},
		},
		"delete-application": {
			Name:   "delete-application",
			Fields: fields_delete_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteApplication(ctx, input)
			},
		},
		"delete-attachment": {
			Name:   "delete-attachment",
			Fields: fields_delete_attachment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAttachmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_attachment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAttachment(ctx, input)
			},
		},
		"delete-chat-controls-configuration": {
			Name:   "delete-chat-controls-configuration",
			Fields: fields_delete_chat_controls_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteChatControlsConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_chat_controls_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteChatControlsConfiguration(ctx, input)
			},
		},
		"delete-chat-response-configuration": {
			Name:   "delete-chat-response-configuration",
			Fields: fields_delete_chat_response_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteChatResponseConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_chat_response_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteChatResponseConfiguration(ctx, input)
			},
		},
		"delete-conversation": {
			Name:   "delete-conversation",
			Fields: fields_delete_conversation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteConversationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_conversation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteConversation(ctx, input)
			},
		},
		"delete-data-accessor": {
			Name:   "delete-data-accessor",
			Fields: fields_delete_data_accessor,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDataAccessorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_data_accessor, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDataAccessor(ctx, input)
			},
		},
		"delete-data-source": {
			Name:   "delete-data-source",
			Fields: fields_delete_data_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDataSourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_data_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDataSource(ctx, input)
			},
		},
		"delete-group": {
			Name:   "delete-group",
			Fields: fields_delete_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteGroup(ctx, input)
			},
		},
		"delete-index": {
			Name:   "delete-index",
			Fields: fields_delete_index,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteIndexInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_index, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteIndex(ctx, input)
			},
		},
		"delete-plugin": {
			Name:   "delete-plugin",
			Fields: fields_delete_plugin,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePluginInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_plugin, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePlugin(ctx, input)
			},
		},
		"delete-retriever": {
			Name:   "delete-retriever",
			Fields: fields_delete_retriever,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRetrieverInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_retriever, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRetriever(ctx, input)
			},
		},
		"delete-user": {
			Name:   "delete-user",
			Fields: fields_delete_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteUser(ctx, input)
			},
		},
		"delete-web-experience": {
			Name:   "delete-web-experience",
			Fields: fields_delete_web_experience,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteWebExperienceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_web_experience, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteWebExperience(ctx, input)
			},
		},
		"disassociate-permission": {
			Name:   "disassociate-permission",
			Fields: fields_disassociate_permission,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociatePermissionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_permission, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociatePermission(ctx, input)
			},
		},
		"get-application": {
			Name:   "get-application",
			Fields: fields_get_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetApplication(ctx, input)
			},
		},
		"get-chat-controls-configuration": {
			Name:   "get-chat-controls-configuration",
			Fields: fields_get_chat_controls_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetChatControlsConfigurationInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_chat_controls_configuration, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetChatControlsConfiguration(ctx, input)
				}
				var results []*svc.GetChatControlsConfigurationOutput
				p := svc.NewGetChatControlsConfigurationPaginator(client, input)
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
		"get-chat-response-configuration": {
			Name:   "get-chat-response-configuration",
			Fields: fields_get_chat_response_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetChatResponseConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_chat_response_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetChatResponseConfiguration(ctx, input)
			},
		},
		"get-data-accessor": {
			Name:   "get-data-accessor",
			Fields: fields_get_data_accessor,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDataAccessorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_data_accessor, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDataAccessor(ctx, input)
			},
		},
		"get-data-source": {
			Name:   "get-data-source",
			Fields: fields_get_data_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDataSourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_data_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDataSource(ctx, input)
			},
		},
		"get-document-content": {
			Name:   "get-document-content",
			Fields: fields_get_document_content,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDocumentContentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_document_content, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDocumentContent(ctx, input)
			},
		},
		"get-group": {
			Name:   "get-group",
			Fields: fields_get_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetGroup(ctx, input)
			},
		},
		"get-index": {
			Name:   "get-index",
			Fields: fields_get_index,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetIndexInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_index, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetIndex(ctx, input)
			},
		},
		"get-media": {
			Name:   "get-media",
			Fields: fields_get_media,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMediaInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_media, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMedia(ctx, input)
			},
		},
		"get-plugin": {
			Name:   "get-plugin",
			Fields: fields_get_plugin,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPluginInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_plugin, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPlugin(ctx, input)
			},
		},
		"get-policy": {
			Name:   "get-policy",
			Fields: fields_get_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPolicy(ctx, input)
			},
		},
		"get-retriever": {
			Name:   "get-retriever",
			Fields: fields_get_retriever,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRetrieverInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_retriever, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRetriever(ctx, input)
			},
		},
		"get-user": {
			Name:   "get-user",
			Fields: fields_get_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetUser(ctx, input)
			},
		},
		"get-web-experience": {
			Name:   "get-web-experience",
			Fields: fields_get_web_experience,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetWebExperienceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_web_experience, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetWebExperience(ctx, input)
			},
		},
		"list-applications": {
			Name:   "list-applications",
			Fields: fields_list_applications,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListApplicationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_applications, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListApplications(ctx, input)
				}
				var results []*svc.ListApplicationsOutput
				p := svc.NewListApplicationsPaginator(client, input)
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
		"list-attachments": {
			Name:   "list-attachments",
			Fields: fields_list_attachments,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAttachmentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_attachments, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAttachments(ctx, input)
				}
				var results []*svc.ListAttachmentsOutput
				p := svc.NewListAttachmentsPaginator(client, input)
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
		"list-chat-response-configurations": {
			Name:   "list-chat-response-configurations",
			Fields: fields_list_chat_response_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListChatResponseConfigurationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_chat_response_configurations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListChatResponseConfigurations(ctx, input)
				}
				var results []*svc.ListChatResponseConfigurationsOutput
				p := svc.NewListChatResponseConfigurationsPaginator(client, input)
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
		"list-conversations": {
			Name:   "list-conversations",
			Fields: fields_list_conversations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListConversationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_conversations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListConversations(ctx, input)
				}
				var results []*svc.ListConversationsOutput
				p := svc.NewListConversationsPaginator(client, input)
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
		"list-data-accessors": {
			Name:   "list-data-accessors",
			Fields: fields_list_data_accessors,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDataAccessorsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_data_accessors, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDataAccessors(ctx, input)
				}
				var results []*svc.ListDataAccessorsOutput
				p := svc.NewListDataAccessorsPaginator(client, input)
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
		"list-data-source-sync-jobs": {
			Name:   "list-data-source-sync-jobs",
			Fields: fields_list_data_source_sync_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDataSourceSyncJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_data_source_sync_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDataSourceSyncJobs(ctx, input)
				}
				var results []*svc.ListDataSourceSyncJobsOutput
				p := svc.NewListDataSourceSyncJobsPaginator(client, input)
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
		"list-data-sources": {
			Name:   "list-data-sources",
			Fields: fields_list_data_sources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDataSourcesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_data_sources, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDataSources(ctx, input)
				}
				var results []*svc.ListDataSourcesOutput
				p := svc.NewListDataSourcesPaginator(client, input)
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
		"list-documents": {
			Name:   "list-documents",
			Fields: fields_list_documents,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDocumentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_documents, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDocuments(ctx, input)
				}
				var results []*svc.ListDocumentsOutput
				p := svc.NewListDocumentsPaginator(client, input)
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
		"list-groups": {
			Name:   "list-groups",
			Fields: fields_list_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListGroups(ctx, input)
				}
				var results []*svc.ListGroupsOutput
				p := svc.NewListGroupsPaginator(client, input)
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
		"list-indices": {
			Name:   "list-indices",
			Fields: fields_list_indices,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListIndicesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_indices, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListIndices(ctx, input)
				}
				var results []*svc.ListIndicesOutput
				p := svc.NewListIndicesPaginator(client, input)
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
		"list-plugin-actions": {
			Name:   "list-plugin-actions",
			Fields: fields_list_plugin_actions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPluginActionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_plugin_actions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPluginActions(ctx, input)
				}
				var results []*svc.ListPluginActionsOutput
				p := svc.NewListPluginActionsPaginator(client, input)
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
		"list-plugin-type-actions": {
			Name:   "list-plugin-type-actions",
			Fields: fields_list_plugin_type_actions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPluginTypeActionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_plugin_type_actions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPluginTypeActions(ctx, input)
				}
				var results []*svc.ListPluginTypeActionsOutput
				p := svc.NewListPluginTypeActionsPaginator(client, input)
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
		"list-plugin-type-metadata": {
			Name:   "list-plugin-type-metadata",
			Fields: fields_list_plugin_type_metadata,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPluginTypeMetadataInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_plugin_type_metadata, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPluginTypeMetadata(ctx, input)
				}
				var results []*svc.ListPluginTypeMetadataOutput
				p := svc.NewListPluginTypeMetadataPaginator(client, input)
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
		"list-plugins": {
			Name:   "list-plugins",
			Fields: fields_list_plugins,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPluginsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_plugins, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPlugins(ctx, input)
				}
				var results []*svc.ListPluginsOutput
				p := svc.NewListPluginsPaginator(client, input)
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
		"list-retrievers": {
			Name:   "list-retrievers",
			Fields: fields_list_retrievers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRetrieversInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_retrievers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRetrievers(ctx, input)
				}
				var results []*svc.ListRetrieversOutput
				p := svc.NewListRetrieversPaginator(client, input)
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
		"list-subscriptions": {
			Name:   "list-subscriptions",
			Fields: fields_list_subscriptions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSubscriptionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_subscriptions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSubscriptions(ctx, input)
				}
				var results []*svc.ListSubscriptionsOutput
				p := svc.NewListSubscriptionsPaginator(client, input)
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
		"list-web-experiences": {
			Name:   "list-web-experiences",
			Fields: fields_list_web_experiences,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListWebExperiencesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_web_experiences, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListWebExperiences(ctx, input)
				}
				var results []*svc.ListWebExperiencesOutput
				p := svc.NewListWebExperiencesPaginator(client, input)
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
		"put-group": {
			Name:   "put-group",
			Fields: fields_put_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutGroup(ctx, input)
			},
		},
		"search-relevant-content": {
			Name:   "search-relevant-content",
			Fields: fields_search_relevant_content,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchRelevantContentInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_relevant_content, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchRelevantContent(ctx, input)
				}
				var results []*svc.SearchRelevantContentOutput
				p := svc.NewSearchRelevantContentPaginator(client, input)
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
		"start-data-source-sync-job": {
			Name:   "start-data-source-sync-job",
			Fields: fields_start_data_source_sync_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartDataSourceSyncJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_data_source_sync_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartDataSourceSyncJob(ctx, input)
			},
		},
		"stop-data-source-sync-job": {
			Name:   "stop-data-source-sync-job",
			Fields: fields_stop_data_source_sync_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopDataSourceSyncJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_data_source_sync_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopDataSourceSyncJob(ctx, input)
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
		"update-application": {
			Name:   "update-application",
			Fields: fields_update_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateApplication(ctx, input)
			},
		},
		"update-chat-controls-configuration": {
			Name:   "update-chat-controls-configuration",
			Fields: fields_update_chat_controls_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateChatControlsConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_chat_controls_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateChatControlsConfiguration(ctx, input)
			},
		},
		"update-chat-response-configuration": {
			Name:   "update-chat-response-configuration",
			Fields: fields_update_chat_response_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateChatResponseConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_chat_response_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateChatResponseConfiguration(ctx, input)
			},
		},
		"update-data-accessor": {
			Name:   "update-data-accessor",
			Fields: fields_update_data_accessor,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDataAccessorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_data_accessor, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDataAccessor(ctx, input)
			},
		},
		"update-data-source": {
			Name:   "update-data-source",
			Fields: fields_update_data_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDataSourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_data_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDataSource(ctx, input)
			},
		},
		"update-index": {
			Name:   "update-index",
			Fields: fields_update_index,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateIndexInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_index, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateIndex(ctx, input)
			},
		},
		"update-plugin": {
			Name:   "update-plugin",
			Fields: fields_update_plugin,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePluginInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_plugin, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePlugin(ctx, input)
			},
		},
		"update-retriever": {
			Name:   "update-retriever",
			Fields: fields_update_retriever,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRetrieverInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_retriever, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRetriever(ctx, input)
			},
		},
		"update-subscription": {
			Name:   "update-subscription",
			Fields: fields_update_subscription,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSubscriptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_subscription, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSubscription(ctx, input)
			},
		},
		"update-user": {
			Name:   "update-user",
			Fields: fields_update_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateUser(ctx, input)
			},
		},
		"update-web-experience": {
			Name:   "update-web-experience",
			Fields: fields_update_web_experience,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateWebExperienceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_web_experience, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateWebExperience(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("qbusiness", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
