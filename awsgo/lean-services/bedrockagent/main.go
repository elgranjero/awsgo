package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/bedrockagent"
)

var fields_associate_agent_collaborator = []leanruntime.Field{
	{Name: "AgentDescriptor", Flag: "agent-descriptor", Type: "*types.AgentDescriptor", Required: true},
	{Name: "AgentId", Flag: "agent-id", Type: "*string", Required: true},
	{Name: "AgentVersion", Flag: "agent-version", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "CollaborationInstruction", Flag: "collaboration-instruction", Type: "*string", Required: true},
	{Name: "CollaboratorName", Flag: "collaborator-name", Type: "*string", Required: true},
	{Name: "RelayConversationHistory", Flag: "relay-conversation-history", Type: "types.RelayConversationHistory", Required: false},
}

var fields_associate_agent_knowledge_base = []leanruntime.Field{
	{Name: "AgentId", Flag: "agent-id", Type: "*string", Required: true},
	{Name: "AgentVersion", Flag: "agent-version", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: true},
	{Name: "KnowledgeBaseId", Flag: "knowledge-base-id", Type: "*string", Required: true},
	{Name: "KnowledgeBaseState", Flag: "knowledge-base-state", Type: "types.KnowledgeBaseState", Required: false},
}

var fields_create_agent = []leanruntime.Field{
	{Name: "AgentCollaboration", Flag: "agent-collaboration", Type: "types.AgentCollaboration", Required: false},
	{Name: "AgentName", Flag: "agent-name", Type: "*string", Required: true},
	{Name: "AgentResourceRoleArn", Flag: "agent-resource-role-arn", Type: "*string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "CustomOrchestration", Flag: "custom-orchestration", Type: "*types.CustomOrchestration", Required: false},
	{Name: "CustomerEncryptionKeyArn", Flag: "customer-encryption-key-arn", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "FoundationModel", Flag: "foundation-model", Type: "*string", Required: false},
	{Name: "GuardrailConfiguration", Flag: "guardrail-configuration", Type: "*types.GuardrailConfiguration", Required: false},
	{Name: "IdleSessionTTLInSeconds", Flag: "idle-session-ttlin-seconds", Type: "*int32", Required: false},
	{Name: "Instruction", Flag: "instruction", Type: "*string", Required: false},
	{Name: "MemoryConfiguration", Flag: "memory-configuration", Type: "*types.MemoryConfiguration", Required: false},
	{Name: "OrchestrationType", Flag: "orchestration-type", Type: "types.OrchestrationType", Required: false},
	{Name: "PromptOverrideConfiguration", Flag: "prompt-override-configuration", Type: "*types.PromptOverrideConfiguration", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_agent_action_group = []leanruntime.Field{
	{Name: "ActionGroupExecutor", Flag: "action-group-executor", Type: "types.ActionGroupExecutor", Required: false},
	{Name: "ActionGroupName", Flag: "action-group-name", Type: "*string", Required: true},
	{Name: "ActionGroupState", Flag: "action-group-state", Type: "types.ActionGroupState", Required: false},
	{Name: "AgentId", Flag: "agent-id", Type: "*string", Required: true},
	{Name: "AgentVersion", Flag: "agent-version", Type: "*string", Required: true},
	{Name: "ApiSchema", Flag: "api-schema", Type: "types.APISchema", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "FunctionSchema", Flag: "function-schema", Type: "types.FunctionSchema", Required: false},
	{Name: "ParentActionGroupSignature", Flag: "parent-action-group-signature", Type: "types.ActionGroupSignature", Required: false},
	{Name: "ParentActionGroupSignatureParams", Flag: "parent-action-group-signature-params", Type: "map[string]string", Required: false},
}

var fields_create_agent_alias = []leanruntime.Field{
	{Name: "AgentAliasName", Flag: "agent-alias-name", Type: "*string", Required: true},
	{Name: "AgentId", Flag: "agent-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "RoutingConfiguration", Flag: "routing-configuration", Type: "[]types.AgentAliasRoutingConfigurationListItem", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_data_source = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DataDeletionPolicy", Flag: "data-deletion-policy", Type: "types.DataDeletionPolicy", Required: false},
	{Name: "DataSourceConfiguration", Flag: "data-source-configuration", Type: "*types.DataSourceConfiguration", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "KnowledgeBaseId", Flag: "knowledge-base-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ServerSideEncryptionConfiguration", Flag: "server-side-encryption-configuration", Type: "*types.ServerSideEncryptionConfiguration", Required: false},
	{Name: "VectorIngestionConfiguration", Flag: "vector-ingestion-configuration", Type: "*types.VectorIngestionConfiguration", Required: false},
}

var fields_create_flow = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "CustomerEncryptionKeyArn", Flag: "customer-encryption-key-arn", Type: "*string", Required: false},
	{Name: "Definition", Flag: "definition", Type: "*types.FlowDefinition", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ExecutionRoleArn", Flag: "execution-role-arn", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_flow_alias = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ConcurrencyConfiguration", Flag: "concurrency-configuration", Type: "*types.FlowAliasConcurrencyConfiguration", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "FlowIdentifier", Flag: "flow-identifier", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RoutingConfiguration", Flag: "routing-configuration", Type: "[]types.FlowAliasRoutingConfigurationListItem", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_flow_version = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "FlowIdentifier", Flag: "flow-identifier", Type: "*string", Required: true},
}

var fields_create_knowledge_base = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "KnowledgeBaseConfiguration", Flag: "knowledge-base-configuration", Type: "*types.KnowledgeBaseConfiguration", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "StorageConfiguration", Flag: "storage-configuration", Type: "*types.StorageConfiguration", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_prompt = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "CustomerEncryptionKeyArn", Flag: "customer-encryption-key-arn", Type: "*string", Required: false},
	{Name: "DefaultVariant", Flag: "default-variant", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Variants", Flag: "variants", Type: "[]types.PromptVariant", Required: false},
}

var fields_create_prompt_version = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "PromptIdentifier", Flag: "prompt-identifier", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_delete_agent = []leanruntime.Field{
	{Name: "AgentId", Flag: "agent-id", Type: "*string", Required: true},
	{Name: "SkipResourceInUseCheck", Flag: "skip-resource-in-use-check", Type: "bool", Required: false},
}

var fields_delete_agent_action_group = []leanruntime.Field{
	{Name: "ActionGroupId", Flag: "action-group-id", Type: "*string", Required: true},
	{Name: "AgentId", Flag: "agent-id", Type: "*string", Required: true},
	{Name: "AgentVersion", Flag: "agent-version", Type: "*string", Required: true},
	{Name: "SkipResourceInUseCheck", Flag: "skip-resource-in-use-check", Type: "bool", Required: false},
}

var fields_delete_agent_alias = []leanruntime.Field{
	{Name: "AgentAliasId", Flag: "agent-alias-id", Type: "*string", Required: true},
	{Name: "AgentId", Flag: "agent-id", Type: "*string", Required: true},
}

var fields_delete_agent_version = []leanruntime.Field{
	{Name: "AgentId", Flag: "agent-id", Type: "*string", Required: true},
	{Name: "AgentVersion", Flag: "agent-version", Type: "*string", Required: true},
	{Name: "SkipResourceInUseCheck", Flag: "skip-resource-in-use-check", Type: "bool", Required: false},
}

var fields_delete_data_source = []leanruntime.Field{
	{Name: "DataSourceId", Flag: "data-source-id", Type: "*string", Required: true},
	{Name: "KnowledgeBaseId", Flag: "knowledge-base-id", Type: "*string", Required: true},
}

var fields_delete_flow = []leanruntime.Field{
	{Name: "FlowIdentifier", Flag: "flow-identifier", Type: "*string", Required: true},
	{Name: "SkipResourceInUseCheck", Flag: "skip-resource-in-use-check", Type: "bool", Required: false},
}

var fields_delete_flow_alias = []leanruntime.Field{
	{Name: "AliasIdentifier", Flag: "alias-identifier", Type: "*string", Required: true},
	{Name: "FlowIdentifier", Flag: "flow-identifier", Type: "*string", Required: true},
}

var fields_delete_flow_version = []leanruntime.Field{
	{Name: "FlowIdentifier", Flag: "flow-identifier", Type: "*string", Required: true},
	{Name: "FlowVersion", Flag: "flow-version", Type: "*string", Required: true},
	{Name: "SkipResourceInUseCheck", Flag: "skip-resource-in-use-check", Type: "bool", Required: false},
}

var fields_delete_knowledge_base = []leanruntime.Field{
	{Name: "KnowledgeBaseId", Flag: "knowledge-base-id", Type: "*string", Required: true},
}

var fields_delete_knowledge_base_documents = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DataSourceId", Flag: "data-source-id", Type: "*string", Required: true},
	{Name: "DocumentIdentifiers", Flag: "document-identifiers", Type: "[]types.DocumentIdentifier", Required: true},
	{Name: "KnowledgeBaseId", Flag: "knowledge-base-id", Type: "*string", Required: true},
}

var fields_delete_prompt = []leanruntime.Field{
	{Name: "PromptIdentifier", Flag: "prompt-identifier", Type: "*string", Required: true},
	{Name: "PromptVersion", Flag: "prompt-version", Type: "*string", Required: false},
}

var fields_disassociate_agent_collaborator = []leanruntime.Field{
	{Name: "AgentId", Flag: "agent-id", Type: "*string", Required: true},
	{Name: "AgentVersion", Flag: "agent-version", Type: "*string", Required: true},
	{Name: "CollaboratorId", Flag: "collaborator-id", Type: "*string", Required: true},
}

var fields_disassociate_agent_knowledge_base = []leanruntime.Field{
	{Name: "AgentId", Flag: "agent-id", Type: "*string", Required: true},
	{Name: "AgentVersion", Flag: "agent-version", Type: "*string", Required: true},
	{Name: "KnowledgeBaseId", Flag: "knowledge-base-id", Type: "*string", Required: true},
}

var fields_get_agent = []leanruntime.Field{
	{Name: "AgentId", Flag: "agent-id", Type: "*string", Required: true},
}

var fields_get_agent_action_group = []leanruntime.Field{
	{Name: "ActionGroupId", Flag: "action-group-id", Type: "*string", Required: true},
	{Name: "AgentId", Flag: "agent-id", Type: "*string", Required: true},
	{Name: "AgentVersion", Flag: "agent-version", Type: "*string", Required: true},
}

var fields_get_agent_alias = []leanruntime.Field{
	{Name: "AgentAliasId", Flag: "agent-alias-id", Type: "*string", Required: true},
	{Name: "AgentId", Flag: "agent-id", Type: "*string", Required: true},
}

var fields_get_agent_collaborator = []leanruntime.Field{
	{Name: "AgentId", Flag: "agent-id", Type: "*string", Required: true},
	{Name: "AgentVersion", Flag: "agent-version", Type: "*string", Required: true},
	{Name: "CollaboratorId", Flag: "collaborator-id", Type: "*string", Required: true},
}

var fields_get_agent_knowledge_base = []leanruntime.Field{
	{Name: "AgentId", Flag: "agent-id", Type: "*string", Required: true},
	{Name: "AgentVersion", Flag: "agent-version", Type: "*string", Required: true},
	{Name: "KnowledgeBaseId", Flag: "knowledge-base-id", Type: "*string", Required: true},
}

var fields_get_agent_version = []leanruntime.Field{
	{Name: "AgentId", Flag: "agent-id", Type: "*string", Required: true},
	{Name: "AgentVersion", Flag: "agent-version", Type: "*string", Required: true},
}

var fields_get_data_source = []leanruntime.Field{
	{Name: "DataSourceId", Flag: "data-source-id", Type: "*string", Required: true},
	{Name: "KnowledgeBaseId", Flag: "knowledge-base-id", Type: "*string", Required: true},
}

var fields_get_flow = []leanruntime.Field{
	{Name: "FlowIdentifier", Flag: "flow-identifier", Type: "*string", Required: true},
}

var fields_get_flow_alias = []leanruntime.Field{
	{Name: "AliasIdentifier", Flag: "alias-identifier", Type: "*string", Required: true},
	{Name: "FlowIdentifier", Flag: "flow-identifier", Type: "*string", Required: true},
}

var fields_get_flow_version = []leanruntime.Field{
	{Name: "FlowIdentifier", Flag: "flow-identifier", Type: "*string", Required: true},
	{Name: "FlowVersion", Flag: "flow-version", Type: "*string", Required: true},
}

var fields_get_ingestion_job = []leanruntime.Field{
	{Name: "DataSourceId", Flag: "data-source-id", Type: "*string", Required: true},
	{Name: "IngestionJobId", Flag: "ingestion-job-id", Type: "*string", Required: true},
	{Name: "KnowledgeBaseId", Flag: "knowledge-base-id", Type: "*string", Required: true},
}

var fields_get_knowledge_base = []leanruntime.Field{
	{Name: "KnowledgeBaseId", Flag: "knowledge-base-id", Type: "*string", Required: true},
}

var fields_get_knowledge_base_documents = []leanruntime.Field{
	{Name: "DataSourceId", Flag: "data-source-id", Type: "*string", Required: true},
	{Name: "DocumentIdentifiers", Flag: "document-identifiers", Type: "[]types.DocumentIdentifier", Required: true},
	{Name: "KnowledgeBaseId", Flag: "knowledge-base-id", Type: "*string", Required: true},
}

var fields_get_prompt = []leanruntime.Field{
	{Name: "PromptIdentifier", Flag: "prompt-identifier", Type: "*string", Required: true},
	{Name: "PromptVersion", Flag: "prompt-version", Type: "*string", Required: false},
}

var fields_ingest_knowledge_base_documents = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DataSourceId", Flag: "data-source-id", Type: "*string", Required: true},
	{Name: "Documents", Flag: "documents", Type: "[]types.KnowledgeBaseDocument", Required: true},
	{Name: "KnowledgeBaseId", Flag: "knowledge-base-id", Type: "*string", Required: true},
}

var fields_list_agent_action_groups = []leanruntime.Field{
	{Name: "AgentId", Flag: "agent-id", Type: "*string", Required: true},
	{Name: "AgentVersion", Flag: "agent-version", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_agent_aliases = []leanruntime.Field{
	{Name: "AgentId", Flag: "agent-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_agent_collaborators = []leanruntime.Field{
	{Name: "AgentId", Flag: "agent-id", Type: "*string", Required: true},
	{Name: "AgentVersion", Flag: "agent-version", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_agent_knowledge_bases = []leanruntime.Field{
	{Name: "AgentId", Flag: "agent-id", Type: "*string", Required: true},
	{Name: "AgentVersion", Flag: "agent-version", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_agent_versions = []leanruntime.Field{
	{Name: "AgentId", Flag: "agent-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_agents = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_data_sources = []leanruntime.Field{
	{Name: "KnowledgeBaseId", Flag: "knowledge-base-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_flow_aliases = []leanruntime.Field{
	{Name: "FlowIdentifier", Flag: "flow-identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_flow_versions = []leanruntime.Field{
	{Name: "FlowIdentifier", Flag: "flow-identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_flows = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_ingestion_jobs = []leanruntime.Field{
	{Name: "DataSourceId", Flag: "data-source-id", Type: "*string", Required: true},
	{Name: "Filters", Flag: "filters", Type: "[]types.IngestionJobFilter", Required: false},
	{Name: "KnowledgeBaseId", Flag: "knowledge-base-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "*types.IngestionJobSortBy", Required: false},
}

var fields_list_knowledge_base_documents = []leanruntime.Field{
	{Name: "DataSourceId", Flag: "data-source-id", Type: "*string", Required: true},
	{Name: "KnowledgeBaseId", Flag: "knowledge-base-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_knowledge_bases = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_prompts = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PromptIdentifier", Flag: "prompt-identifier", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_prepare_agent = []leanruntime.Field{
	{Name: "AgentId", Flag: "agent-id", Type: "*string", Required: true},
}

var fields_prepare_flow = []leanruntime.Field{
	{Name: "FlowIdentifier", Flag: "flow-identifier", Type: "*string", Required: true},
}

var fields_start_ingestion_job = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DataSourceId", Flag: "data-source-id", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "KnowledgeBaseId", Flag: "knowledge-base-id", Type: "*string", Required: true},
}

var fields_stop_ingestion_job = []leanruntime.Field{
	{Name: "DataSourceId", Flag: "data-source-id", Type: "*string", Required: true},
	{Name: "IngestionJobId", Flag: "ingestion-job-id", Type: "*string", Required: true},
	{Name: "KnowledgeBaseId", Flag: "knowledge-base-id", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_agent = []leanruntime.Field{
	{Name: "AgentCollaboration", Flag: "agent-collaboration", Type: "types.AgentCollaboration", Required: false},
	{Name: "AgentId", Flag: "agent-id", Type: "*string", Required: true},
	{Name: "AgentName", Flag: "agent-name", Type: "*string", Required: true},
	{Name: "AgentResourceRoleArn", Flag: "agent-resource-role-arn", Type: "*string", Required: true},
	{Name: "CustomOrchestration", Flag: "custom-orchestration", Type: "*types.CustomOrchestration", Required: false},
	{Name: "CustomerEncryptionKeyArn", Flag: "customer-encryption-key-arn", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "FoundationModel", Flag: "foundation-model", Type: "*string", Required: true},
	{Name: "GuardrailConfiguration", Flag: "guardrail-configuration", Type: "*types.GuardrailConfiguration", Required: false},
	{Name: "IdleSessionTTLInSeconds", Flag: "idle-session-ttlin-seconds", Type: "*int32", Required: false},
	{Name: "Instruction", Flag: "instruction", Type: "*string", Required: false},
	{Name: "MemoryConfiguration", Flag: "memory-configuration", Type: "*types.MemoryConfiguration", Required: false},
	{Name: "OrchestrationType", Flag: "orchestration-type", Type: "types.OrchestrationType", Required: false},
	{Name: "PromptOverrideConfiguration", Flag: "prompt-override-configuration", Type: "*types.PromptOverrideConfiguration", Required: false},
}

var fields_update_agent_action_group = []leanruntime.Field{
	{Name: "ActionGroupExecutor", Flag: "action-group-executor", Type: "types.ActionGroupExecutor", Required: false},
	{Name: "ActionGroupId", Flag: "action-group-id", Type: "*string", Required: true},
	{Name: "ActionGroupName", Flag: "action-group-name", Type: "*string", Required: true},
	{Name: "ActionGroupState", Flag: "action-group-state", Type: "types.ActionGroupState", Required: false},
	{Name: "AgentId", Flag: "agent-id", Type: "*string", Required: true},
	{Name: "AgentVersion", Flag: "agent-version", Type: "*string", Required: true},
	{Name: "ApiSchema", Flag: "api-schema", Type: "types.APISchema", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "FunctionSchema", Flag: "function-schema", Type: "types.FunctionSchema", Required: false},
	{Name: "ParentActionGroupSignature", Flag: "parent-action-group-signature", Type: "types.ActionGroupSignature", Required: false},
	{Name: "ParentActionGroupSignatureParams", Flag: "parent-action-group-signature-params", Type: "map[string]string", Required: false},
}

var fields_update_agent_alias = []leanruntime.Field{
	{Name: "AgentAliasId", Flag: "agent-alias-id", Type: "*string", Required: true},
	{Name: "AgentAliasName", Flag: "agent-alias-name", Type: "*string", Required: true},
	{Name: "AgentId", Flag: "agent-id", Type: "*string", Required: true},
	{Name: "AliasInvocationState", Flag: "alias-invocation-state", Type: "types.AliasInvocationState", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "RoutingConfiguration", Flag: "routing-configuration", Type: "[]types.AgentAliasRoutingConfigurationListItem", Required: false},
}

var fields_update_agent_collaborator = []leanruntime.Field{
	{Name: "AgentDescriptor", Flag: "agent-descriptor", Type: "*types.AgentDescriptor", Required: true},
	{Name: "AgentId", Flag: "agent-id", Type: "*string", Required: true},
	{Name: "AgentVersion", Flag: "agent-version", Type: "*string", Required: true},
	{Name: "CollaborationInstruction", Flag: "collaboration-instruction", Type: "*string", Required: true},
	{Name: "CollaboratorId", Flag: "collaborator-id", Type: "*string", Required: true},
	{Name: "CollaboratorName", Flag: "collaborator-name", Type: "*string", Required: true},
	{Name: "RelayConversationHistory", Flag: "relay-conversation-history", Type: "types.RelayConversationHistory", Required: false},
}

var fields_update_agent_knowledge_base = []leanruntime.Field{
	{Name: "AgentId", Flag: "agent-id", Type: "*string", Required: true},
	{Name: "AgentVersion", Flag: "agent-version", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "KnowledgeBaseId", Flag: "knowledge-base-id", Type: "*string", Required: true},
	{Name: "KnowledgeBaseState", Flag: "knowledge-base-state", Type: "types.KnowledgeBaseState", Required: false},
}

var fields_update_data_source = []leanruntime.Field{
	{Name: "DataDeletionPolicy", Flag: "data-deletion-policy", Type: "types.DataDeletionPolicy", Required: false},
	{Name: "DataSourceConfiguration", Flag: "data-source-configuration", Type: "*types.DataSourceConfiguration", Required: true},
	{Name: "DataSourceId", Flag: "data-source-id", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "KnowledgeBaseId", Flag: "knowledge-base-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ServerSideEncryptionConfiguration", Flag: "server-side-encryption-configuration", Type: "*types.ServerSideEncryptionConfiguration", Required: false},
	{Name: "VectorIngestionConfiguration", Flag: "vector-ingestion-configuration", Type: "*types.VectorIngestionConfiguration", Required: false},
}

var fields_update_flow = []leanruntime.Field{
	{Name: "CustomerEncryptionKeyArn", Flag: "customer-encryption-key-arn", Type: "*string", Required: false},
	{Name: "Definition", Flag: "definition", Type: "*types.FlowDefinition", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ExecutionRoleArn", Flag: "execution-role-arn", Type: "*string", Required: true},
	{Name: "FlowIdentifier", Flag: "flow-identifier", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_update_flow_alias = []leanruntime.Field{
	{Name: "AliasIdentifier", Flag: "alias-identifier", Type: "*string", Required: true},
	{Name: "ConcurrencyConfiguration", Flag: "concurrency-configuration", Type: "*types.FlowAliasConcurrencyConfiguration", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "FlowIdentifier", Flag: "flow-identifier", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RoutingConfiguration", Flag: "routing-configuration", Type: "[]types.FlowAliasRoutingConfigurationListItem", Required: true},
}

var fields_update_knowledge_base = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "KnowledgeBaseConfiguration", Flag: "knowledge-base-configuration", Type: "*types.KnowledgeBaseConfiguration", Required: true},
	{Name: "KnowledgeBaseId", Flag: "knowledge-base-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "StorageConfiguration", Flag: "storage-configuration", Type: "*types.StorageConfiguration", Required: false},
}

var fields_update_prompt = []leanruntime.Field{
	{Name: "CustomerEncryptionKeyArn", Flag: "customer-encryption-key-arn", Type: "*string", Required: false},
	{Name: "DefaultVariant", Flag: "default-variant", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "PromptIdentifier", Flag: "prompt-identifier", Type: "*string", Required: true},
	{Name: "Variants", Flag: "variants", Type: "[]types.PromptVariant", Required: false},
}

var fields_validate_flow_definition = []leanruntime.Field{
	{Name: "Definition", Flag: "definition", Type: "*types.FlowDefinition", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"associate-agent-collaborator": {
			Name:   "associate-agent-collaborator",
			Fields: fields_associate_agent_collaborator,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateAgentCollaboratorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_agent_collaborator, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateAgentCollaborator(ctx, input)
			},
		},
		"associate-agent-knowledge-base": {
			Name:   "associate-agent-knowledge-base",
			Fields: fields_associate_agent_knowledge_base,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateAgentKnowledgeBaseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_agent_knowledge_base, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateAgentKnowledgeBase(ctx, input)
			},
		},
		"create-agent": {
			Name:   "create-agent",
			Fields: fields_create_agent,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAgentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_agent, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAgent(ctx, input)
			},
		},
		"create-agent-action-group": {
			Name:   "create-agent-action-group",
			Fields: fields_create_agent_action_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAgentActionGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_agent_action_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAgentActionGroup(ctx, input)
			},
		},
		"create-agent-alias": {
			Name:   "create-agent-alias",
			Fields: fields_create_agent_alias,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAgentAliasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_agent_alias, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAgentAlias(ctx, input)
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
		"create-flow": {
			Name:   "create-flow",
			Fields: fields_create_flow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateFlowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_flow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateFlow(ctx, input)
			},
		},
		"create-flow-alias": {
			Name:   "create-flow-alias",
			Fields: fields_create_flow_alias,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateFlowAliasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_flow_alias, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateFlowAlias(ctx, input)
			},
		},
		"create-flow-version": {
			Name:   "create-flow-version",
			Fields: fields_create_flow_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateFlowVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_flow_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateFlowVersion(ctx, input)
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
		"create-prompt": {
			Name:   "create-prompt",
			Fields: fields_create_prompt,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePromptInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_prompt, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePrompt(ctx, input)
			},
		},
		"create-prompt-version": {
			Name:   "create-prompt-version",
			Fields: fields_create_prompt_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePromptVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_prompt_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePromptVersion(ctx, input)
			},
		},
		"delete-agent": {
			Name:   "delete-agent",
			Fields: fields_delete_agent,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAgentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_agent, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAgent(ctx, input)
			},
		},
		"delete-agent-action-group": {
			Name:   "delete-agent-action-group",
			Fields: fields_delete_agent_action_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAgentActionGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_agent_action_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAgentActionGroup(ctx, input)
			},
		},
		"delete-agent-alias": {
			Name:   "delete-agent-alias",
			Fields: fields_delete_agent_alias,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAgentAliasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_agent_alias, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAgentAlias(ctx, input)
			},
		},
		"delete-agent-version": {
			Name:   "delete-agent-version",
			Fields: fields_delete_agent_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAgentVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_agent_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAgentVersion(ctx, input)
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
		"delete-flow": {
			Name:   "delete-flow",
			Fields: fields_delete_flow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteFlowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_flow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteFlow(ctx, input)
			},
		},
		"delete-flow-alias": {
			Name:   "delete-flow-alias",
			Fields: fields_delete_flow_alias,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteFlowAliasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_flow_alias, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteFlowAlias(ctx, input)
			},
		},
		"delete-flow-version": {
			Name:   "delete-flow-version",
			Fields: fields_delete_flow_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteFlowVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_flow_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteFlowVersion(ctx, input)
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
		"delete-knowledge-base-documents": {
			Name:   "delete-knowledge-base-documents",
			Fields: fields_delete_knowledge_base_documents,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteKnowledgeBaseDocumentsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_knowledge_base_documents, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteKnowledgeBaseDocuments(ctx, input)
			},
		},
		"delete-prompt": {
			Name:   "delete-prompt",
			Fields: fields_delete_prompt,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePromptInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_prompt, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePrompt(ctx, input)
			},
		},
		"disassociate-agent-collaborator": {
			Name:   "disassociate-agent-collaborator",
			Fields: fields_disassociate_agent_collaborator,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateAgentCollaboratorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_agent_collaborator, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateAgentCollaborator(ctx, input)
			},
		},
		"disassociate-agent-knowledge-base": {
			Name:   "disassociate-agent-knowledge-base",
			Fields: fields_disassociate_agent_knowledge_base,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateAgentKnowledgeBaseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_agent_knowledge_base, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateAgentKnowledgeBase(ctx, input)
			},
		},
		"get-agent": {
			Name:   "get-agent",
			Fields: fields_get_agent,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAgentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_agent, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAgent(ctx, input)
			},
		},
		"get-agent-action-group": {
			Name:   "get-agent-action-group",
			Fields: fields_get_agent_action_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAgentActionGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_agent_action_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAgentActionGroup(ctx, input)
			},
		},
		"get-agent-alias": {
			Name:   "get-agent-alias",
			Fields: fields_get_agent_alias,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAgentAliasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_agent_alias, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAgentAlias(ctx, input)
			},
		},
		"get-agent-collaborator": {
			Name:   "get-agent-collaborator",
			Fields: fields_get_agent_collaborator,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAgentCollaboratorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_agent_collaborator, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAgentCollaborator(ctx, input)
			},
		},
		"get-agent-knowledge-base": {
			Name:   "get-agent-knowledge-base",
			Fields: fields_get_agent_knowledge_base,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAgentKnowledgeBaseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_agent_knowledge_base, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAgentKnowledgeBase(ctx, input)
			},
		},
		"get-agent-version": {
			Name:   "get-agent-version",
			Fields: fields_get_agent_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAgentVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_agent_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAgentVersion(ctx, input)
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
		"get-flow": {
			Name:   "get-flow",
			Fields: fields_get_flow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFlowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_flow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetFlow(ctx, input)
			},
		},
		"get-flow-alias": {
			Name:   "get-flow-alias",
			Fields: fields_get_flow_alias,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFlowAliasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_flow_alias, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetFlowAlias(ctx, input)
			},
		},
		"get-flow-version": {
			Name:   "get-flow-version",
			Fields: fields_get_flow_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFlowVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_flow_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetFlowVersion(ctx, input)
			},
		},
		"get-ingestion-job": {
			Name:   "get-ingestion-job",
			Fields: fields_get_ingestion_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetIngestionJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_ingestion_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetIngestionJob(ctx, input)
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
		"get-knowledge-base-documents": {
			Name:   "get-knowledge-base-documents",
			Fields: fields_get_knowledge_base_documents,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetKnowledgeBaseDocumentsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_knowledge_base_documents, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetKnowledgeBaseDocuments(ctx, input)
			},
		},
		"get-prompt": {
			Name:   "get-prompt",
			Fields: fields_get_prompt,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPromptInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_prompt, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPrompt(ctx, input)
			},
		},
		"ingest-knowledge-base-documents": {
			Name:   "ingest-knowledge-base-documents",
			Fields: fields_ingest_knowledge_base_documents,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.IngestKnowledgeBaseDocumentsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_ingest_knowledge_base_documents, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.IngestKnowledgeBaseDocuments(ctx, input)
			},
		},
		"list-agent-action-groups": {
			Name:   "list-agent-action-groups",
			Fields: fields_list_agent_action_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAgentActionGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_agent_action_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAgentActionGroups(ctx, input)
				}
				var results []*svc.ListAgentActionGroupsOutput
				p := svc.NewListAgentActionGroupsPaginator(client, input)
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
		"list-agent-aliases": {
			Name:   "list-agent-aliases",
			Fields: fields_list_agent_aliases,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAgentAliasesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_agent_aliases, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAgentAliases(ctx, input)
				}
				var results []*svc.ListAgentAliasesOutput
				p := svc.NewListAgentAliasesPaginator(client, input)
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
		"list-agent-collaborators": {
			Name:   "list-agent-collaborators",
			Fields: fields_list_agent_collaborators,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAgentCollaboratorsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_agent_collaborators, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAgentCollaborators(ctx, input)
				}
				var results []*svc.ListAgentCollaboratorsOutput
				p := svc.NewListAgentCollaboratorsPaginator(client, input)
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
		"list-agent-knowledge-bases": {
			Name:   "list-agent-knowledge-bases",
			Fields: fields_list_agent_knowledge_bases,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAgentKnowledgeBasesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_agent_knowledge_bases, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAgentKnowledgeBases(ctx, input)
				}
				var results []*svc.ListAgentKnowledgeBasesOutput
				p := svc.NewListAgentKnowledgeBasesPaginator(client, input)
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
		"list-agent-versions": {
			Name:   "list-agent-versions",
			Fields: fields_list_agent_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAgentVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_agent_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAgentVersions(ctx, input)
				}
				var results []*svc.ListAgentVersionsOutput
				p := svc.NewListAgentVersionsPaginator(client, input)
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
		"list-agents": {
			Name:   "list-agents",
			Fields: fields_list_agents,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAgentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_agents, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAgents(ctx, input)
				}
				var results []*svc.ListAgentsOutput
				p := svc.NewListAgentsPaginator(client, input)
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
		"list-flow-aliases": {
			Name:   "list-flow-aliases",
			Fields: fields_list_flow_aliases,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFlowAliasesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_flow_aliases, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFlowAliases(ctx, input)
				}
				var results []*svc.ListFlowAliasesOutput
				p := svc.NewListFlowAliasesPaginator(client, input)
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
		"list-flow-versions": {
			Name:   "list-flow-versions",
			Fields: fields_list_flow_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFlowVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_flow_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFlowVersions(ctx, input)
				}
				var results []*svc.ListFlowVersionsOutput
				p := svc.NewListFlowVersionsPaginator(client, input)
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
		"list-flows": {
			Name:   "list-flows",
			Fields: fields_list_flows,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFlowsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_flows, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFlows(ctx, input)
				}
				var results []*svc.ListFlowsOutput
				p := svc.NewListFlowsPaginator(client, input)
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
		"list-ingestion-jobs": {
			Name:   "list-ingestion-jobs",
			Fields: fields_list_ingestion_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListIngestionJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_ingestion_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListIngestionJobs(ctx, input)
				}
				var results []*svc.ListIngestionJobsOutput
				p := svc.NewListIngestionJobsPaginator(client, input)
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
		"list-knowledge-base-documents": {
			Name:   "list-knowledge-base-documents",
			Fields: fields_list_knowledge_base_documents,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListKnowledgeBaseDocumentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_knowledge_base_documents, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListKnowledgeBaseDocuments(ctx, input)
				}
				var results []*svc.ListKnowledgeBaseDocumentsOutput
				p := svc.NewListKnowledgeBaseDocumentsPaginator(client, input)
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
		"list-prompts": {
			Name:   "list-prompts",
			Fields: fields_list_prompts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPromptsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_prompts, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPrompts(ctx, input)
				}
				var results []*svc.ListPromptsOutput
				p := svc.NewListPromptsPaginator(client, input)
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
		"prepare-agent": {
			Name:   "prepare-agent",
			Fields: fields_prepare_agent,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PrepareAgentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_prepare_agent, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PrepareAgent(ctx, input)
			},
		},
		"prepare-flow": {
			Name:   "prepare-flow",
			Fields: fields_prepare_flow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PrepareFlowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_prepare_flow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PrepareFlow(ctx, input)
			},
		},
		"start-ingestion-job": {
			Name:   "start-ingestion-job",
			Fields: fields_start_ingestion_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartIngestionJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_ingestion_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartIngestionJob(ctx, input)
			},
		},
		"stop-ingestion-job": {
			Name:   "stop-ingestion-job",
			Fields: fields_stop_ingestion_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopIngestionJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_ingestion_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopIngestionJob(ctx, input)
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
		"update-agent": {
			Name:   "update-agent",
			Fields: fields_update_agent,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAgentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_agent, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAgent(ctx, input)
			},
		},
		"update-agent-action-group": {
			Name:   "update-agent-action-group",
			Fields: fields_update_agent_action_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAgentActionGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_agent_action_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAgentActionGroup(ctx, input)
			},
		},
		"update-agent-alias": {
			Name:   "update-agent-alias",
			Fields: fields_update_agent_alias,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAgentAliasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_agent_alias, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAgentAlias(ctx, input)
			},
		},
		"update-agent-collaborator": {
			Name:   "update-agent-collaborator",
			Fields: fields_update_agent_collaborator,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAgentCollaboratorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_agent_collaborator, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAgentCollaborator(ctx, input)
			},
		},
		"update-agent-knowledge-base": {
			Name:   "update-agent-knowledge-base",
			Fields: fields_update_agent_knowledge_base,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAgentKnowledgeBaseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_agent_knowledge_base, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAgentKnowledgeBase(ctx, input)
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
		"update-flow": {
			Name:   "update-flow",
			Fields: fields_update_flow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateFlowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_flow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateFlow(ctx, input)
			},
		},
		"update-flow-alias": {
			Name:   "update-flow-alias",
			Fields: fields_update_flow_alias,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateFlowAliasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_flow_alias, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateFlowAlias(ctx, input)
			},
		},
		"update-knowledge-base": {
			Name:   "update-knowledge-base",
			Fields: fields_update_knowledge_base,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateKnowledgeBaseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_knowledge_base, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateKnowledgeBase(ctx, input)
			},
		},
		"update-prompt": {
			Name:   "update-prompt",
			Fields: fields_update_prompt,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePromptInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_prompt, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePrompt(ctx, input)
			},
		},
		"validate-flow-definition": {
			Name:   "validate-flow-definition",
			Fields: fields_validate_flow_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ValidateFlowDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_validate_flow_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ValidateFlowDefinition(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("bedrockagent", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
