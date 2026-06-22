package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/bedrockagentruntime"
)

var fields_create_invocation = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "InvocationId", Flag: "invocation-id", Type: "*string", Required: false},
	{Name: "SessionIdentifier", Flag: "session-identifier", Type: "*string", Required: true},
}

var fields_create_session = []leanruntime.Field{
	{Name: "EncryptionKeyArn", Flag: "encryption-key-arn", Type: "*string", Required: false},
	{Name: "SessionMetadata", Flag: "session-metadata", Type: "map[string]string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_delete_agent_memory = []leanruntime.Field{
	{Name: "AgentAliasId", Flag: "agent-alias-id", Type: "*string", Required: true},
	{Name: "AgentId", Flag: "agent-id", Type: "*string", Required: true},
	{Name: "MemoryId", Flag: "memory-id", Type: "*string", Required: false},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: false},
}

var fields_delete_session = []leanruntime.Field{
	{Name: "SessionIdentifier", Flag: "session-identifier", Type: "*string", Required: true},
}

var fields_end_session = []leanruntime.Field{
	{Name: "SessionIdentifier", Flag: "session-identifier", Type: "*string", Required: true},
}

var fields_generate_query = []leanruntime.Field{
	{Name: "QueryGenerationInput", Flag: "query-generation-input", Type: "*types.QueryGenerationInput", Required: true},
	{Name: "TransformationConfiguration", Flag: "transformation-configuration", Type: "*types.TransformationConfiguration", Required: true},
}

var fields_get_agent_memory = []leanruntime.Field{
	{Name: "AgentAliasId", Flag: "agent-alias-id", Type: "*string", Required: true},
	{Name: "AgentId", Flag: "agent-id", Type: "*string", Required: true},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "MemoryId", Flag: "memory-id", Type: "*string", Required: true},
	{Name: "MemoryType", Flag: "memory-type", Type: "types.MemoryType", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_execution_flow_snapshot = []leanruntime.Field{
	{Name: "ExecutionIdentifier", Flag: "execution-identifier", Type: "*string", Required: true},
	{Name: "FlowAliasIdentifier", Flag: "flow-alias-identifier", Type: "*string", Required: true},
	{Name: "FlowIdentifier", Flag: "flow-identifier", Type: "*string", Required: true},
}

var fields_get_flow_execution = []leanruntime.Field{
	{Name: "ExecutionIdentifier", Flag: "execution-identifier", Type: "*string", Required: true},
	{Name: "FlowAliasIdentifier", Flag: "flow-alias-identifier", Type: "*string", Required: true},
	{Name: "FlowIdentifier", Flag: "flow-identifier", Type: "*string", Required: true},
}

var fields_get_invocation_step = []leanruntime.Field{
	{Name: "InvocationIdentifier", Flag: "invocation-identifier", Type: "*string", Required: true},
	{Name: "InvocationStepId", Flag: "invocation-step-id", Type: "*string", Required: true},
	{Name: "SessionIdentifier", Flag: "session-identifier", Type: "*string", Required: true},
}

var fields_get_session = []leanruntime.Field{
	{Name: "SessionIdentifier", Flag: "session-identifier", Type: "*string", Required: true},
}

var fields_invoke_agent = []leanruntime.Field{
	{Name: "AgentAliasId", Flag: "agent-alias-id", Type: "*string", Required: true},
	{Name: "AgentId", Flag: "agent-id", Type: "*string", Required: true},
	{Name: "BedrockModelConfigurations", Flag: "bedrock-model-configurations", Type: "*types.BedrockModelConfigurations", Required: false},
	{Name: "EnableTrace", Flag: "enable-trace", Type: "*bool", Required: false},
	{Name: "EndSession", Flag: "end-session", Type: "*bool", Required: false},
	{Name: "InputText", Flag: "input-text", Type: "*string", Required: false},
	{Name: "MemoryId", Flag: "memory-id", Type: "*string", Required: false},
	{Name: "PromptCreationConfigurations", Flag: "prompt-creation-configurations", Type: "*types.PromptCreationConfigurations", Required: false},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: true},
	{Name: "SessionState", Flag: "session-state", Type: "*types.SessionState", Required: false},
	{Name: "SourceArn", Flag: "source-arn", Type: "*string", Required: false},
	{Name: "StreamingConfigurations", Flag: "streaming-configurations", Type: "*types.StreamingConfigurations", Required: false},
}

var fields_invoke_flow = []leanruntime.Field{
	{Name: "EnableTrace", Flag: "enable-trace", Type: "*bool", Required: false},
	{Name: "ExecutionId", Flag: "execution-id", Type: "*string", Required: false},
	{Name: "FlowAliasIdentifier", Flag: "flow-alias-identifier", Type: "*string", Required: true},
	{Name: "FlowIdentifier", Flag: "flow-identifier", Type: "*string", Required: true},
	{Name: "Inputs", Flag: "inputs", Type: "[]types.FlowInput", Required: true},
	{Name: "ModelPerformanceConfiguration", Flag: "model-performance-configuration", Type: "*types.ModelPerformanceConfiguration", Required: false},
}

var fields_invoke_inline_agent = []leanruntime.Field{
	{Name: "ActionGroups", Flag: "action-groups", Type: "[]types.AgentActionGroup", Required: false},
	{Name: "AgentCollaboration", Flag: "agent-collaboration", Type: "types.AgentCollaboration", Required: false},
	{Name: "AgentName", Flag: "agent-name", Type: "*string", Required: false},
	{Name: "BedrockModelConfigurations", Flag: "bedrock-model-configurations", Type: "*types.InlineBedrockModelConfigurations", Required: false},
	{Name: "CollaboratorConfigurations", Flag: "collaborator-configurations", Type: "[]types.CollaboratorConfiguration", Required: false},
	{Name: "Collaborators", Flag: "collaborators", Type: "[]types.Collaborator", Required: false},
	{Name: "CustomOrchestration", Flag: "custom-orchestration", Type: "*types.CustomOrchestration", Required: false},
	{Name: "CustomerEncryptionKeyArn", Flag: "customer-encryption-key-arn", Type: "*string", Required: false},
	{Name: "EnableTrace", Flag: "enable-trace", Type: "*bool", Required: false},
	{Name: "EndSession", Flag: "end-session", Type: "*bool", Required: false},
	{Name: "FoundationModel", Flag: "foundation-model", Type: "*string", Required: true},
	{Name: "GuardrailConfiguration", Flag: "guardrail-configuration", Type: "*types.GuardrailConfigurationWithArn", Required: false},
	{Name: "IdleSessionTTLInSeconds", Flag: "idle-session-ttlin-seconds", Type: "*int32", Required: false},
	{Name: "InlineSessionState", Flag: "inline-session-state", Type: "*types.InlineSessionState", Required: false},
	{Name: "InputText", Flag: "input-text", Type: "*string", Required: false},
	{Name: "Instruction", Flag: "instruction", Type: "*string", Required: true},
	{Name: "KnowledgeBases", Flag: "knowledge-bases", Type: "[]types.KnowledgeBase", Required: false},
	{Name: "OrchestrationType", Flag: "orchestration-type", Type: "types.OrchestrationType", Required: false},
	{Name: "PromptCreationConfigurations", Flag: "prompt-creation-configurations", Type: "*types.PromptCreationConfigurations", Required: false},
	{Name: "PromptOverrideConfiguration", Flag: "prompt-override-configuration", Type: "*types.PromptOverrideConfiguration", Required: false},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: true},
	{Name: "StreamingConfigurations", Flag: "streaming-configurations", Type: "*types.StreamingConfigurations", Required: false},
}

var fields_list_flow_execution_events = []leanruntime.Field{
	{Name: "EventType", Flag: "event-type", Type: "types.FlowExecutionEventType", Required: true},
	{Name: "ExecutionIdentifier", Flag: "execution-identifier", Type: "*string", Required: true},
	{Name: "FlowAliasIdentifier", Flag: "flow-alias-identifier", Type: "*string", Required: true},
	{Name: "FlowIdentifier", Flag: "flow-identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_flow_executions = []leanruntime.Field{
	{Name: "FlowAliasIdentifier", Flag: "flow-alias-identifier", Type: "*string", Required: false},
	{Name: "FlowIdentifier", Flag: "flow-identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_invocation_steps = []leanruntime.Field{
	{Name: "InvocationIdentifier", Flag: "invocation-identifier", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SessionIdentifier", Flag: "session-identifier", Type: "*string", Required: true},
}

var fields_list_invocations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SessionIdentifier", Flag: "session-identifier", Type: "*string", Required: true},
}

var fields_list_sessions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_optimize_prompt = []leanruntime.Field{
	{Name: "Input", Flag: "input", Type: "types.InputPrompt", Required: true},
	{Name: "TargetModelId", Flag: "target-model-id", Type: "*string", Required: true},
}

var fields_put_invocation_step = []leanruntime.Field{
	{Name: "InvocationIdentifier", Flag: "invocation-identifier", Type: "*string", Required: true},
	{Name: "InvocationStepId", Flag: "invocation-step-id", Type: "*string", Required: false},
	{Name: "InvocationStepTime", Flag: "invocation-step-time", Type: "*time.Time", Required: true},
	{Name: "Payload", Flag: "payload", Type: "types.InvocationStepPayload", Required: true},
	{Name: "SessionIdentifier", Flag: "session-identifier", Type: "*string", Required: true},
}

var fields_rerank = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Queries", Flag: "queries", Type: "[]types.RerankQuery", Required: true},
	{Name: "RerankingConfiguration", Flag: "reranking-configuration", Type: "*types.RerankingConfiguration", Required: true},
	{Name: "Sources", Flag: "sources", Type: "[]types.RerankSource", Required: true},
}

var fields_retrieve = []leanruntime.Field{
	{Name: "GuardrailConfiguration", Flag: "guardrail-configuration", Type: "*types.GuardrailConfiguration", Required: false},
	{Name: "KnowledgeBaseId", Flag: "knowledge-base-id", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RetrievalConfiguration", Flag: "retrieval-configuration", Type: "*types.KnowledgeBaseRetrievalConfiguration", Required: false},
	{Name: "RetrievalQuery", Flag: "retrieval-query", Type: "*types.KnowledgeBaseQuery", Required: true},
}

var fields_retrieve_and_generate = []leanruntime.Field{
	{Name: "Input", Flag: "input", Type: "*types.RetrieveAndGenerateInput", Required: true},
	{Name: "RetrieveAndGenerateConfiguration", Flag: "retrieve-and-generate-configuration", Type: "*types.RetrieveAndGenerateConfiguration", Required: false},
	{Name: "SessionConfiguration", Flag: "session-configuration", Type: "*types.RetrieveAndGenerateSessionConfiguration", Required: false},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: false},
}

var fields_retrieve_and_generate_stream = []leanruntime.Field{
	{Name: "Input", Flag: "input", Type: "*types.RetrieveAndGenerateInput", Required: true},
	{Name: "RetrieveAndGenerateConfiguration", Flag: "retrieve-and-generate-configuration", Type: "*types.RetrieveAndGenerateConfiguration", Required: false},
	{Name: "SessionConfiguration", Flag: "session-configuration", Type: "*types.RetrieveAndGenerateSessionConfiguration", Required: false},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: false},
}

var fields_start_flow_execution = []leanruntime.Field{
	{Name: "FlowAliasIdentifier", Flag: "flow-alias-identifier", Type: "*string", Required: true},
	{Name: "FlowExecutionName", Flag: "flow-execution-name", Type: "*string", Required: false},
	{Name: "FlowIdentifier", Flag: "flow-identifier", Type: "*string", Required: true},
	{Name: "Inputs", Flag: "inputs", Type: "[]types.FlowInput", Required: true},
	{Name: "ModelPerformanceConfiguration", Flag: "model-performance-configuration", Type: "*types.ModelPerformanceConfiguration", Required: false},
}

var fields_stop_flow_execution = []leanruntime.Field{
	{Name: "ExecutionIdentifier", Flag: "execution-identifier", Type: "*string", Required: true},
	{Name: "FlowAliasIdentifier", Flag: "flow-alias-identifier", Type: "*string", Required: true},
	{Name: "FlowIdentifier", Flag: "flow-identifier", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_session = []leanruntime.Field{
	{Name: "SessionIdentifier", Flag: "session-identifier", Type: "*string", Required: true},
	{Name: "SessionMetadata", Flag: "session-metadata", Type: "map[string]string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-invocation": {
			Name:   "create-invocation",
			Fields: fields_create_invocation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateInvocationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_invocation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateInvocation(ctx, input)
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
		"delete-agent-memory": {
			Name:   "delete-agent-memory",
			Fields: fields_delete_agent_memory,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAgentMemoryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_agent_memory, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAgentMemory(ctx, input)
			},
		},
		"delete-session": {
			Name:   "delete-session",
			Fields: fields_delete_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSession(ctx, input)
			},
		},
		"end-session": {
			Name:   "end-session",
			Fields: fields_end_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EndSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_end_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EndSession(ctx, input)
			},
		},
		"generate-query": {
			Name:   "generate-query",
			Fields: fields_generate_query,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GenerateQueryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_generate_query, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GenerateQuery(ctx, input)
			},
		},
		"get-agent-memory": {
			Name:   "get-agent-memory",
			Fields: fields_get_agent_memory,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAgentMemoryInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_agent_memory, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetAgentMemory(ctx, input)
				}
				var results []*svc.GetAgentMemoryOutput
				p := svc.NewGetAgentMemoryPaginator(client, input)
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
		"get-execution-flow-snapshot": {
			Name:   "get-execution-flow-snapshot",
			Fields: fields_get_execution_flow_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetExecutionFlowSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_execution_flow_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetExecutionFlowSnapshot(ctx, input)
			},
		},
		"get-flow-execution": {
			Name:   "get-flow-execution",
			Fields: fields_get_flow_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFlowExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_flow_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetFlowExecution(ctx, input)
			},
		},
		"get-invocation-step": {
			Name:   "get-invocation-step",
			Fields: fields_get_invocation_step,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetInvocationStepInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_invocation_step, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetInvocationStep(ctx, input)
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
		"invoke-agent": {
			Name:   "invoke-agent",
			Fields: fields_invoke_agent,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.InvokeAgentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_invoke_agent, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.InvokeAgent(ctx, input)
			},
		},
		"invoke-flow": {
			Name:   "invoke-flow",
			Fields: fields_invoke_flow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.InvokeFlowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_invoke_flow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.InvokeFlow(ctx, input)
			},
		},
		"invoke-inline-agent": {
			Name:   "invoke-inline-agent",
			Fields: fields_invoke_inline_agent,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.InvokeInlineAgentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_invoke_inline_agent, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.InvokeInlineAgent(ctx, input)
			},
		},
		"list-flow-execution-events": {
			Name:   "list-flow-execution-events",
			Fields: fields_list_flow_execution_events,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFlowExecutionEventsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_flow_execution_events, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFlowExecutionEvents(ctx, input)
				}
				var results []*svc.ListFlowExecutionEventsOutput
				p := svc.NewListFlowExecutionEventsPaginator(client, input)
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
		"list-flow-executions": {
			Name:   "list-flow-executions",
			Fields: fields_list_flow_executions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFlowExecutionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_flow_executions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFlowExecutions(ctx, input)
				}
				var results []*svc.ListFlowExecutionsOutput
				p := svc.NewListFlowExecutionsPaginator(client, input)
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
		"list-invocation-steps": {
			Name:   "list-invocation-steps",
			Fields: fields_list_invocation_steps,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListInvocationStepsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_invocation_steps, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListInvocationSteps(ctx, input)
				}
				var results []*svc.ListInvocationStepsOutput
				p := svc.NewListInvocationStepsPaginator(client, input)
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
		"list-invocations": {
			Name:   "list-invocations",
			Fields: fields_list_invocations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListInvocationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_invocations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListInvocations(ctx, input)
				}
				var results []*svc.ListInvocationsOutput
				p := svc.NewListInvocationsPaginator(client, input)
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
		"list-sessions": {
			Name:   "list-sessions",
			Fields: fields_list_sessions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSessionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_sessions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSessions(ctx, input)
				}
				var results []*svc.ListSessionsOutput
				p := svc.NewListSessionsPaginator(client, input)
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
		"optimize-prompt": {
			Name:   "optimize-prompt",
			Fields: fields_optimize_prompt,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.OptimizePromptInput{}
				if _, err := leanruntime.ApplyInput(input, fields_optimize_prompt, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.OptimizePrompt(ctx, input)
			},
		},
		"put-invocation-step": {
			Name:   "put-invocation-step",
			Fields: fields_put_invocation_step,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutInvocationStepInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_invocation_step, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutInvocationStep(ctx, input)
			},
		},
		"rerank": {
			Name:   "rerank",
			Fields: fields_rerank,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RerankInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_rerank, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.Rerank(ctx, input)
				}
				var results []*svc.RerankOutput
				p := svc.NewRerankPaginator(client, input)
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
		"retrieve": {
			Name:   "retrieve",
			Fields: fields_retrieve,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RetrieveInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_retrieve, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.Retrieve(ctx, input)
				}
				var results []*svc.RetrieveOutput
				p := svc.NewRetrievePaginator(client, input)
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
		"retrieve-and-generate": {
			Name:   "retrieve-and-generate",
			Fields: fields_retrieve_and_generate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RetrieveAndGenerateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_retrieve_and_generate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RetrieveAndGenerate(ctx, input)
			},
		},
		"retrieve-and-generate-stream": {
			Name:   "retrieve-and-generate-stream",
			Fields: fields_retrieve_and_generate_stream,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RetrieveAndGenerateStreamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_retrieve_and_generate_stream, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RetrieveAndGenerateStream(ctx, input)
			},
		},
		"start-flow-execution": {
			Name:   "start-flow-execution",
			Fields: fields_start_flow_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartFlowExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_flow_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartFlowExecution(ctx, input)
			},
		},
		"stop-flow-execution": {
			Name:   "stop-flow-execution",
			Fields: fields_stop_flow_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopFlowExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_flow_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopFlowExecution(ctx, input)
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
	}
	if err := leanruntime.Execute("bedrockagentruntime", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
