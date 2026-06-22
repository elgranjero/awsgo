package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
)

var fields_apply_guardrail = []leanruntime.Field{
	{Name: "Content", Flag: "content", Type: "[]types.GuardrailContentBlock", Required: true},
	{Name: "GuardrailIdentifier", Flag: "guardrail-identifier", Type: "*string", Required: true},
	{Name: "GuardrailVersion", Flag: "guardrail-version", Type: "*string", Required: true},
	{Name: "OutputScope", Flag: "output-scope", Type: "types.GuardrailOutputScope", Required: false},
	{Name: "Source", Flag: "source", Type: "types.GuardrailContentSource", Required: true},
}

var fields_converse = []leanruntime.Field{
	{Name: "AdditionalModelRequestFields", Flag: "additional-model-request-fields", Type: "document.Interface", Required: false},
	{Name: "AdditionalModelResponseFieldPaths", Flag: "additional-model-response-field-paths", Type: "[]string", Required: false},
	{Name: "GuardrailConfig", Flag: "guardrail-config", Type: "*types.GuardrailConfiguration", Required: false},
	{Name: "InferenceConfig", Flag: "inference-config", Type: "*types.InferenceConfiguration", Required: false},
	{Name: "Messages", Flag: "messages", Type: "[]types.Message", Required: false},
	{Name: "ModelId", Flag: "model-id", Type: "*string", Required: true},
	{Name: "OutputConfig", Flag: "output-config", Type: "*types.OutputConfig", Required: false},
	{Name: "PerformanceConfig", Flag: "performance-config", Type: "*types.PerformanceConfiguration", Required: false},
	{Name: "PromptVariables", Flag: "prompt-variables", Type: "map[string]types.PromptVariableValues", Required: false},
	{Name: "RequestMetadata", Flag: "request-metadata", Type: "map[string]string", Required: false},
	{Name: "ServiceTier", Flag: "service-tier", Type: "*types.ServiceTier", Required: false},
	{Name: "System", Flag: "system", Type: "[]types.SystemContentBlock", Required: false},
	{Name: "ToolConfig", Flag: "tool-config", Type: "*types.ToolConfiguration", Required: false},
}

var fields_converse_stream = []leanruntime.Field{
	{Name: "AdditionalModelRequestFields", Flag: "additional-model-request-fields", Type: "document.Interface", Required: false},
	{Name: "AdditionalModelResponseFieldPaths", Flag: "additional-model-response-field-paths", Type: "[]string", Required: false},
	{Name: "GuardrailConfig", Flag: "guardrail-config", Type: "*types.GuardrailStreamConfiguration", Required: false},
	{Name: "InferenceConfig", Flag: "inference-config", Type: "*types.InferenceConfiguration", Required: false},
	{Name: "Messages", Flag: "messages", Type: "[]types.Message", Required: false},
	{Name: "ModelId", Flag: "model-id", Type: "*string", Required: true},
	{Name: "OutputConfig", Flag: "output-config", Type: "*types.OutputConfig", Required: false},
	{Name: "PerformanceConfig", Flag: "performance-config", Type: "*types.PerformanceConfiguration", Required: false},
	{Name: "PromptVariables", Flag: "prompt-variables", Type: "map[string]types.PromptVariableValues", Required: false},
	{Name: "RequestMetadata", Flag: "request-metadata", Type: "map[string]string", Required: false},
	{Name: "ServiceTier", Flag: "service-tier", Type: "*types.ServiceTier", Required: false},
	{Name: "System", Flag: "system", Type: "[]types.SystemContentBlock", Required: false},
	{Name: "ToolConfig", Flag: "tool-config", Type: "*types.ToolConfiguration", Required: false},
}

var fields_count_tokens = []leanruntime.Field{
	{Name: "Input", Flag: "input", Type: "types.CountTokensInput", Required: true},
	{Name: "ModelId", Flag: "model-id", Type: "*string", Required: true},
}

var fields_get_async_invoke = []leanruntime.Field{
	{Name: "InvocationArn", Flag: "invocation-arn", Type: "*string", Required: true},
}

var fields_invoke_model = []leanruntime.Field{
	{Name: "Accept", Flag: "accept", Type: "*string", Required: false},
	{Name: "Body", Flag: "body", Type: "[]byte", Required: false},
	{Name: "ContentType", Flag: "content-type", Type: "*string", Required: false},
	{Name: "GuardrailIdentifier", Flag: "guardrail-identifier", Type: "*string", Required: false},
	{Name: "GuardrailVersion", Flag: "guardrail-version", Type: "*string", Required: false},
	{Name: "ModelId", Flag: "model-id", Type: "*string", Required: true},
	{Name: "PerformanceConfigLatency", Flag: "performance-config-latency", Type: "types.PerformanceConfigLatency", Required: false},
	{Name: "ServiceTier", Flag: "service-tier", Type: "types.ServiceTierType", Required: false},
	{Name: "Trace", Flag: "trace", Type: "types.Trace", Required: false},
}

var fields_invoke_model_with_bidirectional_stream = []leanruntime.Field{
	{Name: "ModelId", Flag: "model-id", Type: "*string", Required: true},
}

var fields_invoke_model_with_response_stream = []leanruntime.Field{
	{Name: "Accept", Flag: "accept", Type: "*string", Required: false},
	{Name: "Body", Flag: "body", Type: "[]byte", Required: false},
	{Name: "ContentType", Flag: "content-type", Type: "*string", Required: false},
	{Name: "GuardrailIdentifier", Flag: "guardrail-identifier", Type: "*string", Required: false},
	{Name: "GuardrailVersion", Flag: "guardrail-version", Type: "*string", Required: false},
	{Name: "ModelId", Flag: "model-id", Type: "*string", Required: true},
	{Name: "PerformanceConfigLatency", Flag: "performance-config-latency", Type: "types.PerformanceConfigLatency", Required: false},
	{Name: "ServiceTier", Flag: "service-tier", Type: "types.ServiceTierType", Required: false},
	{Name: "Trace", Flag: "trace", Type: "types.Trace", Required: false},
}

var fields_list_async_invokes = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.SortAsyncInvocationBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
	{Name: "StatusEquals", Flag: "status-equals", Type: "types.AsyncInvokeStatus", Required: false},
	{Name: "SubmitTimeAfter", Flag: "submit-time-after", Type: "*time.Time", Required: false},
	{Name: "SubmitTimeBefore", Flag: "submit-time-before", Type: "*time.Time", Required: false},
}

var fields_start_async_invoke = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "ModelId", Flag: "model-id", Type: "*string", Required: true},
	{Name: "ModelInput", Flag: "model-input", Type: "document.Interface", Required: true},
	{Name: "OutputDataConfig", Flag: "output-data-config", Type: "types.AsyncInvokeOutputDataConfig", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"apply-guardrail": {
			Name:   "apply-guardrail",
			Fields: fields_apply_guardrail,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ApplyGuardrailInput{}
				if _, err := leanruntime.ApplyInput(input, fields_apply_guardrail, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ApplyGuardrail(ctx, input)
			},
		},
		"converse": {
			Name:   "converse",
			Fields: fields_converse,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ConverseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_converse, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.Converse(ctx, input)
			},
		},
		"converse-stream": {
			Name:   "converse-stream",
			Fields: fields_converse_stream,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ConverseStreamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_converse_stream, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ConverseStream(ctx, input)
			},
		},
		"count-tokens": {
			Name:   "count-tokens",
			Fields: fields_count_tokens,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CountTokensInput{}
				if _, err := leanruntime.ApplyInput(input, fields_count_tokens, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CountTokens(ctx, input)
			},
		},
		"get-async-invoke": {
			Name:   "get-async-invoke",
			Fields: fields_get_async_invoke,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAsyncInvokeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_async_invoke, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAsyncInvoke(ctx, input)
			},
		},
		"invoke-model": {
			Name:   "invoke-model",
			Fields: fields_invoke_model,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.InvokeModelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_invoke_model, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.InvokeModel(ctx, input)
			},
		},
		"invoke-model-with-bidirectional-stream": {
			Name:   "invoke-model-with-bidirectional-stream",
			Fields: fields_invoke_model_with_bidirectional_stream,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.InvokeModelWithBidirectionalStreamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_invoke_model_with_bidirectional_stream, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.InvokeModelWithBidirectionalStream(ctx, input)
			},
		},
		"invoke-model-with-response-stream": {
			Name:   "invoke-model-with-response-stream",
			Fields: fields_invoke_model_with_response_stream,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.InvokeModelWithResponseStreamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_invoke_model_with_response_stream, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.InvokeModelWithResponseStream(ctx, input)
			},
		},
		"list-async-invokes": {
			Name:   "list-async-invokes",
			Fields: fields_list_async_invokes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAsyncInvokesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_async_invokes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAsyncInvokes(ctx, input)
				}
				var results []*svc.ListAsyncInvokesOutput
				p := svc.NewListAsyncInvokesPaginator(client, input)
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
		"start-async-invoke": {
			Name:   "start-async-invoke",
			Fields: fields_start_async_invoke,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartAsyncInvokeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_async_invoke, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartAsyncInvoke(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("bedrockruntime", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
