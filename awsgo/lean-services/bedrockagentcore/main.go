package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/bedrockagentcore"
)

var fields_batch_create_memory_records = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "MemoryId", Flag: "memory-id", Type: "*string", Required: true},
	{Name: "Records", Flag: "records", Type: "[]types.MemoryRecordCreateInput", Required: true},
}

var fields_batch_delete_memory_records = []leanruntime.Field{
	{Name: "MemoryId", Flag: "memory-id", Type: "*string", Required: true},
	{Name: "Records", Flag: "records", Type: "[]types.MemoryRecordDeleteInput", Required: true},
}

var fields_batch_update_memory_records = []leanruntime.Field{
	{Name: "MemoryId", Flag: "memory-id", Type: "*string", Required: true},
	{Name: "Records", Flag: "records", Type: "[]types.MemoryRecordUpdateInput", Required: true},
}

var fields_complete_resource_token_auth = []leanruntime.Field{
	{Name: "SessionUri", Flag: "session-uri", Type: "*string", Required: true},
	{Name: "UserIdentifier", Flag: "user-identifier", Type: "types.UserIdentifier", Required: true},
}

var fields_create_event = []leanruntime.Field{
	{Name: "ActorId", Flag: "actor-id", Type: "*string", Required: true},
	{Name: "Branch", Flag: "branch", Type: "*types.Branch", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "EventTimestamp", Flag: "event-timestamp", Type: "*time.Time", Required: true},
	{Name: "MemoryId", Flag: "memory-id", Type: "*string", Required: true},
	{Name: "Metadata", Flag: "metadata", Type: "map[string]types.MetadataValue", Required: false},
	{Name: "Payload", Flag: "payload", Type: "[]types.PayloadType", Required: true},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: false},
}

var fields_delete_event = []leanruntime.Field{
	{Name: "ActorId", Flag: "actor-id", Type: "*string", Required: true},
	{Name: "EventId", Flag: "event-id", Type: "*string", Required: true},
	{Name: "MemoryId", Flag: "memory-id", Type: "*string", Required: true},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: true},
}

var fields_delete_memory_record = []leanruntime.Field{
	{Name: "MemoryId", Flag: "memory-id", Type: "*string", Required: true},
	{Name: "MemoryRecordId", Flag: "memory-record-id", Type: "*string", Required: true},
}

var fields_evaluate = []leanruntime.Field{
	{Name: "EvaluationInput", Flag: "evaluation-input", Type: "types.EvaluationInput", Required: true},
	{Name: "EvaluationTarget", Flag: "evaluation-target", Type: "types.EvaluationTarget", Required: false},
	{Name: "EvaluatorId", Flag: "evaluator-id", Type: "*string", Required: true},
}

var fields_get_agent_card = []leanruntime.Field{
	{Name: "AgentRuntimeArn", Flag: "agent-runtime-arn", Type: "*string", Required: true},
	{Name: "Qualifier", Flag: "qualifier", Type: "*string", Required: false},
	{Name: "RuntimeSessionId", Flag: "runtime-session-id", Type: "*string", Required: false},
}

var fields_get_browser_session = []leanruntime.Field{
	{Name: "BrowserIdentifier", Flag: "browser-identifier", Type: "*string", Required: true},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: true},
}

var fields_get_code_interpreter_session = []leanruntime.Field{
	{Name: "CodeInterpreterIdentifier", Flag: "code-interpreter-identifier", Type: "*string", Required: true},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: true},
}

var fields_get_event = []leanruntime.Field{
	{Name: "ActorId", Flag: "actor-id", Type: "*string", Required: true},
	{Name: "EventId", Flag: "event-id", Type: "*string", Required: true},
	{Name: "MemoryId", Flag: "memory-id", Type: "*string", Required: true},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: true},
}

var fields_get_memory_record = []leanruntime.Field{
	{Name: "MemoryId", Flag: "memory-id", Type: "*string", Required: true},
	{Name: "MemoryRecordId", Flag: "memory-record-id", Type: "*string", Required: true},
}

var fields_get_resource_api_key = []leanruntime.Field{
	{Name: "ResourceCredentialProviderName", Flag: "resource-credential-provider-name", Type: "*string", Required: true},
	{Name: "WorkloadIdentityToken", Flag: "workload-identity-token", Type: "*string", Required: true},
}

var fields_get_resource_oauth2_token = []leanruntime.Field{
	{Name: "CustomParameters", Flag: "custom-parameters", Type: "map[string]string", Required: false},
	{Name: "CustomState", Flag: "custom-state", Type: "*string", Required: false},
	{Name: "ForceAuthentication", Flag: "force-authentication", Type: "*bool", Required: false},
	{Name: "Oauth2Flow", Flag: "oauth2-flow", Type: "types.Oauth2FlowType", Required: true},
	{Name: "ResourceCredentialProviderName", Flag: "resource-credential-provider-name", Type: "*string", Required: true},
	{Name: "ResourceOauth2ReturnUrl", Flag: "resource-oauth2-return-url", Type: "*string", Required: false},
	{Name: "Scopes", Flag: "scopes", Type: "[]string", Required: true},
	{Name: "SessionUri", Flag: "session-uri", Type: "*string", Required: false},
	{Name: "WorkloadIdentityToken", Flag: "workload-identity-token", Type: "*string", Required: true},
}

var fields_get_workload_access_token = []leanruntime.Field{
	{Name: "WorkloadName", Flag: "workload-name", Type: "*string", Required: true},
}

var fields_get_workload_access_token_for_jwt = []leanruntime.Field{
	{Name: "UserToken", Flag: "user-token", Type: "*string", Required: true},
	{Name: "WorkloadName", Flag: "workload-name", Type: "*string", Required: true},
}

var fields_get_workload_access_token_for_user_id = []leanruntime.Field{
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
	{Name: "WorkloadName", Flag: "workload-name", Type: "*string", Required: true},
}

var fields_invoke_agent_runtime = []leanruntime.Field{
	{Name: "Accept", Flag: "accept", Type: "*string", Required: false},
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: false},
	{Name: "AgentRuntimeArn", Flag: "agent-runtime-arn", Type: "*string", Required: true},
	{Name: "Baggage", Flag: "baggage", Type: "*string", Required: false},
	{Name: "ContentType", Flag: "content-type", Type: "*string", Required: false},
	{Name: "McpProtocolVersion", Flag: "mcp-protocol-version", Type: "*string", Required: false},
	{Name: "McpSessionId", Flag: "mcp-session-id", Type: "*string", Required: false},
	{Name: "Payload", Flag: "payload", Type: "[]byte", Required: true},
	{Name: "Qualifier", Flag: "qualifier", Type: "*string", Required: false},
	{Name: "RuntimeSessionId", Flag: "runtime-session-id", Type: "*string", Required: false},
	{Name: "RuntimeUserId", Flag: "runtime-user-id", Type: "*string", Required: false},
	{Name: "TraceId", Flag: "trace-id", Type: "*string", Required: false},
	{Name: "TraceParent", Flag: "trace-parent", Type: "*string", Required: false},
	{Name: "TraceState", Flag: "trace-state", Type: "*string", Required: false},
}

var fields_invoke_code_interpreter = []leanruntime.Field{
	{Name: "Arguments", Flag: "arguments", Type: "*types.ToolArguments", Required: false},
	{Name: "CodeInterpreterIdentifier", Flag: "code-interpreter-identifier", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "types.ToolName", Required: true},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: false},
	{Name: "TraceId", Flag: "trace-id", Type: "*string", Required: false},
	{Name: "TraceParent", Flag: "trace-parent", Type: "*string", Required: false},
}

var fields_list_actors = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MemoryId", Flag: "memory-id", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_browser_sessions = []leanruntime.Field{
	{Name: "BrowserIdentifier", Flag: "browser-identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.BrowserSessionStatus", Required: false},
}

var fields_list_code_interpreter_sessions = []leanruntime.Field{
	{Name: "CodeInterpreterIdentifier", Flag: "code-interpreter-identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.CodeInterpreterSessionStatus", Required: false},
}

var fields_list_events = []leanruntime.Field{
	{Name: "ActorId", Flag: "actor-id", Type: "*string", Required: true},
	{Name: "Filter", Flag: "filter", Type: "*types.FilterInput", Required: false},
	{Name: "IncludePayloads", Flag: "include-payloads", Type: "*bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MemoryId", Flag: "memory-id", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: true},
}

var fields_list_memory_extraction_jobs = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.ExtractionJobFilterInput", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MemoryId", Flag: "memory-id", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_memory_records = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MemoryId", Flag: "memory-id", Type: "*string", Required: true},
	{Name: "MemoryStrategyId", Flag: "memory-strategy-id", Type: "*string", Required: false},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_sessions = []leanruntime.Field{
	{Name: "ActorId", Flag: "actor-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MemoryId", Flag: "memory-id", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_retrieve_memory_records = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MemoryId", Flag: "memory-id", Type: "*string", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SearchCriteria", Flag: "search-criteria", Type: "*types.SearchCriteria", Required: true},
}

var fields_save_browser_session_profile = []leanruntime.Field{
	{Name: "BrowserIdentifier", Flag: "browser-identifier", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ProfileIdentifier", Flag: "profile-identifier", Type: "*string", Required: true},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: true},
	{Name: "TraceId", Flag: "trace-id", Type: "*string", Required: false},
	{Name: "TraceParent", Flag: "trace-parent", Type: "*string", Required: false},
}

var fields_start_browser_session = []leanruntime.Field{
	{Name: "BrowserIdentifier", Flag: "browser-identifier", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Extensions", Flag: "extensions", Type: "[]types.BrowserExtension", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "ProfileConfiguration", Flag: "profile-configuration", Type: "*types.BrowserProfileConfiguration", Required: false},
	{Name: "ProxyConfiguration", Flag: "proxy-configuration", Type: "*types.ProxyConfiguration", Required: false},
	{Name: "SessionTimeoutSeconds", Flag: "session-timeout-seconds", Type: "*int32", Required: false},
	{Name: "TraceId", Flag: "trace-id", Type: "*string", Required: false},
	{Name: "TraceParent", Flag: "trace-parent", Type: "*string", Required: false},
	{Name: "ViewPort", Flag: "view-port", Type: "*types.ViewPort", Required: false},
}

var fields_start_code_interpreter_session = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "CodeInterpreterIdentifier", Flag: "code-interpreter-identifier", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "SessionTimeoutSeconds", Flag: "session-timeout-seconds", Type: "*int32", Required: false},
	{Name: "TraceId", Flag: "trace-id", Type: "*string", Required: false},
	{Name: "TraceParent", Flag: "trace-parent", Type: "*string", Required: false},
}

var fields_start_memory_extraction_job = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ExtractionJob", Flag: "extraction-job", Type: "*types.ExtractionJob", Required: true},
	{Name: "MemoryId", Flag: "memory-id", Type: "*string", Required: true},
}

var fields_stop_browser_session = []leanruntime.Field{
	{Name: "BrowserIdentifier", Flag: "browser-identifier", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: true},
	{Name: "TraceId", Flag: "trace-id", Type: "*string", Required: false},
	{Name: "TraceParent", Flag: "trace-parent", Type: "*string", Required: false},
}

var fields_stop_code_interpreter_session = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "CodeInterpreterIdentifier", Flag: "code-interpreter-identifier", Type: "*string", Required: true},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: true},
	{Name: "TraceId", Flag: "trace-id", Type: "*string", Required: false},
	{Name: "TraceParent", Flag: "trace-parent", Type: "*string", Required: false},
}

var fields_stop_runtime_session = []leanruntime.Field{
	{Name: "AgentRuntimeArn", Flag: "agent-runtime-arn", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Qualifier", Flag: "qualifier", Type: "*string", Required: false},
	{Name: "RuntimeSessionId", Flag: "runtime-session-id", Type: "*string", Required: true},
}

var fields_update_browser_stream = []leanruntime.Field{
	{Name: "BrowserIdentifier", Flag: "browser-identifier", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: true},
	{Name: "StreamUpdate", Flag: "stream-update", Type: "types.StreamUpdate", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"batch-create-memory-records": {
			Name:   "batch-create-memory-records",
			Fields: fields_batch_create_memory_records,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchCreateMemoryRecordsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_create_memory_records, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchCreateMemoryRecords(ctx, input)
			},
		},
		"batch-delete-memory-records": {
			Name:   "batch-delete-memory-records",
			Fields: fields_batch_delete_memory_records,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchDeleteMemoryRecordsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_delete_memory_records, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchDeleteMemoryRecords(ctx, input)
			},
		},
		"batch-update-memory-records": {
			Name:   "batch-update-memory-records",
			Fields: fields_batch_update_memory_records,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchUpdateMemoryRecordsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_update_memory_records, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchUpdateMemoryRecords(ctx, input)
			},
		},
		"complete-resource-token-auth": {
			Name:   "complete-resource-token-auth",
			Fields: fields_complete_resource_token_auth,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CompleteResourceTokenAuthInput{}
				if _, err := leanruntime.ApplyInput(input, fields_complete_resource_token_auth, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CompleteResourceTokenAuth(ctx, input)
			},
		},
		"create-event": {
			Name:   "create-event",
			Fields: fields_create_event,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateEventInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_event, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateEvent(ctx, input)
			},
		},
		"delete-event": {
			Name:   "delete-event",
			Fields: fields_delete_event,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEventInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_event, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEvent(ctx, input)
			},
		},
		"delete-memory-record": {
			Name:   "delete-memory-record",
			Fields: fields_delete_memory_record,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMemoryRecordInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_memory_record, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMemoryRecord(ctx, input)
			},
		},
		"evaluate": {
			Name:   "evaluate",
			Fields: fields_evaluate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EvaluateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_evaluate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.Evaluate(ctx, input)
			},
		},
		"get-agent-card": {
			Name:   "get-agent-card",
			Fields: fields_get_agent_card,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAgentCardInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_agent_card, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAgentCard(ctx, input)
			},
		},
		"get-browser-session": {
			Name:   "get-browser-session",
			Fields: fields_get_browser_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBrowserSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_browser_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBrowserSession(ctx, input)
			},
		},
		"get-code-interpreter-session": {
			Name:   "get-code-interpreter-session",
			Fields: fields_get_code_interpreter_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCodeInterpreterSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_code_interpreter_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCodeInterpreterSession(ctx, input)
			},
		},
		"get-event": {
			Name:   "get-event",
			Fields: fields_get_event,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEventInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_event, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEvent(ctx, input)
			},
		},
		"get-memory-record": {
			Name:   "get-memory-record",
			Fields: fields_get_memory_record,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMemoryRecordInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_memory_record, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMemoryRecord(ctx, input)
			},
		},
		"get-resource-api-key": {
			Name:   "get-resource-api-key",
			Fields: fields_get_resource_api_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetResourceApiKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_resource_api_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetResourceApiKey(ctx, input)
			},
		},
		"get-resource-oauth2-token": {
			Name:   "get-resource-oauth2-token",
			Fields: fields_get_resource_oauth2_token,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetResourceOauth2TokenInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_resource_oauth2_token, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetResourceOauth2Token(ctx, input)
			},
		},
		"get-workload-access-token": {
			Name:   "get-workload-access-token",
			Fields: fields_get_workload_access_token,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetWorkloadAccessTokenInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_workload_access_token, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetWorkloadAccessToken(ctx, input)
			},
		},
		"get-workload-access-token-for-jwt": {
			Name:   "get-workload-access-token-for-jwt",
			Fields: fields_get_workload_access_token_for_jwt,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetWorkloadAccessTokenForJWTInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_workload_access_token_for_jwt, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetWorkloadAccessTokenForJWT(ctx, input)
			},
		},
		"get-workload-access-token-for-user-id": {
			Name:   "get-workload-access-token-for-user-id",
			Fields: fields_get_workload_access_token_for_user_id,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetWorkloadAccessTokenForUserIdInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_workload_access_token_for_user_id, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetWorkloadAccessTokenForUserId(ctx, input)
			},
		},
		"invoke-agent-runtime": {
			Name:   "invoke-agent-runtime",
			Fields: fields_invoke_agent_runtime,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.InvokeAgentRuntimeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_invoke_agent_runtime, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.InvokeAgentRuntime(ctx, input)
			},
		},
		"invoke-code-interpreter": {
			Name:   "invoke-code-interpreter",
			Fields: fields_invoke_code_interpreter,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.InvokeCodeInterpreterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_invoke_code_interpreter, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.InvokeCodeInterpreter(ctx, input)
			},
		},
		"list-actors": {
			Name:   "list-actors",
			Fields: fields_list_actors,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListActorsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_actors, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListActors(ctx, input)
				}
				var results []*svc.ListActorsOutput
				p := svc.NewListActorsPaginator(client, input)
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
		"list-browser-sessions": {
			Name:   "list-browser-sessions",
			Fields: fields_list_browser_sessions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBrowserSessionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_browser_sessions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListBrowserSessions(ctx, input)
			},
		},
		"list-code-interpreter-sessions": {
			Name:   "list-code-interpreter-sessions",
			Fields: fields_list_code_interpreter_sessions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCodeInterpreterSessionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_code_interpreter_sessions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListCodeInterpreterSessions(ctx, input)
			},
		},
		"list-events": {
			Name:   "list-events",
			Fields: fields_list_events,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEventsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_events, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEvents(ctx, input)
				}
				var results []*svc.ListEventsOutput
				p := svc.NewListEventsPaginator(client, input)
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
		"list-memory-extraction-jobs": {
			Name:   "list-memory-extraction-jobs",
			Fields: fields_list_memory_extraction_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMemoryExtractionJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_memory_extraction_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMemoryExtractionJobs(ctx, input)
				}
				var results []*svc.ListMemoryExtractionJobsOutput
				p := svc.NewListMemoryExtractionJobsPaginator(client, input)
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
		"list-memory-records": {
			Name:   "list-memory-records",
			Fields: fields_list_memory_records,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMemoryRecordsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_memory_records, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMemoryRecords(ctx, input)
				}
				var results []*svc.ListMemoryRecordsOutput
				p := svc.NewListMemoryRecordsPaginator(client, input)
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
		"retrieve-memory-records": {
			Name:   "retrieve-memory-records",
			Fields: fields_retrieve_memory_records,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RetrieveMemoryRecordsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_retrieve_memory_records, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.RetrieveMemoryRecords(ctx, input)
				}
				var results []*svc.RetrieveMemoryRecordsOutput
				p := svc.NewRetrieveMemoryRecordsPaginator(client, input)
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
		"save-browser-session-profile": {
			Name:   "save-browser-session-profile",
			Fields: fields_save_browser_session_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SaveBrowserSessionProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_save_browser_session_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SaveBrowserSessionProfile(ctx, input)
			},
		},
		"start-browser-session": {
			Name:   "start-browser-session",
			Fields: fields_start_browser_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartBrowserSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_browser_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartBrowserSession(ctx, input)
			},
		},
		"start-code-interpreter-session": {
			Name:   "start-code-interpreter-session",
			Fields: fields_start_code_interpreter_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartCodeInterpreterSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_code_interpreter_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartCodeInterpreterSession(ctx, input)
			},
		},
		"start-memory-extraction-job": {
			Name:   "start-memory-extraction-job",
			Fields: fields_start_memory_extraction_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartMemoryExtractionJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_memory_extraction_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartMemoryExtractionJob(ctx, input)
			},
		},
		"stop-browser-session": {
			Name:   "stop-browser-session",
			Fields: fields_stop_browser_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopBrowserSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_browser_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopBrowserSession(ctx, input)
			},
		},
		"stop-code-interpreter-session": {
			Name:   "stop-code-interpreter-session",
			Fields: fields_stop_code_interpreter_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopCodeInterpreterSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_code_interpreter_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopCodeInterpreterSession(ctx, input)
			},
		},
		"stop-runtime-session": {
			Name:   "stop-runtime-session",
			Fields: fields_stop_runtime_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopRuntimeSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_runtime_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopRuntimeSession(ctx, input)
			},
		},
		"update-browser-stream": {
			Name:   "update-browser-stream",
			Fields: fields_update_browser_stream,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateBrowserStreamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_browser_stream, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateBrowserStream(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("bedrockagentcore", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
