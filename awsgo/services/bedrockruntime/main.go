package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/bedrockruntime/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"apply-guardrail", "converse", "converse-stream", "count-tokens", "get-async-invoke", "invoke-model", "invoke-model-with-bidirectional-stream", "invoke-model-with-response-stream", "list-async-invokes", "start-async-invoke"},
		OperationSet: map[string]bool{"apply-guardrail": true, "converse": true, "converse-stream": true, "count-tokens": true, "get-async-invoke": true, "invoke-model": true, "invoke-model-with-bidirectional-stream": true, "invoke-model-with-response-stream": true, "list-async-invokes": true, "start-async-invoke": true},
		OperationInputs: map[string][]string{
			"apply-guardrail":                        {"Content", "GuardrailIdentifier", "GuardrailVersion", "OutputScope", "Source"},
			"converse":                               {"AdditionalModelRequestFields", "AdditionalModelResponseFieldPaths", "GuardrailConfig", "InferenceConfig", "Messages", "ModelId", "OutputConfig", "PerformanceConfig", "PromptVariables", "RequestMetadata", "ServiceTier", "System", "ToolConfig"},
			"converse-stream":                        {"AdditionalModelRequestFields", "AdditionalModelResponseFieldPaths", "GuardrailConfig", "InferenceConfig", "Messages", "ModelId", "OutputConfig", "PerformanceConfig", "PromptVariables", "RequestMetadata", "ServiceTier", "System", "ToolConfig"},
			"count-tokens":                           {"Input", "ModelId"},
			"get-async-invoke":                       {"InvocationArn"},
			"invoke-model":                           {"Accept", "Body", "ContentType", "GuardrailIdentifier", "GuardrailVersion", "ModelId", "PerformanceConfigLatency", "ServiceTier", "Trace"},
			"invoke-model-with-bidirectional-stream": {"ModelId"},
			"invoke-model-with-response-stream":      {"Accept", "Body", "ContentType", "GuardrailIdentifier", "GuardrailVersion", "ModelId", "PerformanceConfigLatency", "ServiceTier", "Trace"},
			"list-async-invokes":                     {"MaxResults", "NextToken", "SortBy", "SortOrder", "StatusEquals", "SubmitTimeAfter", "SubmitTimeBefore"},
			"start-async-invoke":                     {"ClientRequestToken", "ModelId", "ModelInput", "OutputDataConfig", "Tags"},
		},
		OperationInputTypes: map[string]map[string]string{
			"apply-guardrail":                        {"Content": "[]types.GuardrailContentBlock", "GuardrailIdentifier": "*string", "GuardrailVersion": "*string", "OutputScope": "types.GuardrailOutputScope", "Source": "types.GuardrailContentSource"},
			"converse":                               {"AdditionalModelRequestFields": "document.Interface", "AdditionalModelResponseFieldPaths": "[]string", "GuardrailConfig": "*types.GuardrailConfiguration", "InferenceConfig": "*types.InferenceConfiguration", "Messages": "[]types.Message", "ModelId": "*string", "OutputConfig": "*types.OutputConfig", "PerformanceConfig": "*types.PerformanceConfiguration", "PromptVariables": "map[string]types.PromptVariableValues", "RequestMetadata": "map[string]string", "ServiceTier": "*types.ServiceTier", "System": "[]types.SystemContentBlock", "ToolConfig": "*types.ToolConfiguration"},
			"converse-stream":                        {"AdditionalModelRequestFields": "document.Interface", "AdditionalModelResponseFieldPaths": "[]string", "GuardrailConfig": "*types.GuardrailStreamConfiguration", "InferenceConfig": "*types.InferenceConfiguration", "Messages": "[]types.Message", "ModelId": "*string", "OutputConfig": "*types.OutputConfig", "PerformanceConfig": "*types.PerformanceConfiguration", "PromptVariables": "map[string]types.PromptVariableValues", "RequestMetadata": "map[string]string", "ServiceTier": "*types.ServiceTier", "System": "[]types.SystemContentBlock", "ToolConfig": "*types.ToolConfiguration"},
			"count-tokens":                           {"Input": "types.CountTokensInput", "ModelId": "*string"},
			"get-async-invoke":                       {"InvocationArn": "*string"},
			"invoke-model":                           {"Accept": "*string", "Body": "[]byte", "ContentType": "*string", "GuardrailIdentifier": "*string", "GuardrailVersion": "*string", "ModelId": "*string", "PerformanceConfigLatency": "types.PerformanceConfigLatency", "ServiceTier": "types.ServiceTierType", "Trace": "types.Trace"},
			"invoke-model-with-bidirectional-stream": {"ModelId": "*string"},
			"invoke-model-with-response-stream":      {"Accept": "*string", "Body": "[]byte", "ContentType": "*string", "GuardrailIdentifier": "*string", "GuardrailVersion": "*string", "ModelId": "*string", "PerformanceConfigLatency": "types.PerformanceConfigLatency", "ServiceTier": "types.ServiceTierType", "Trace": "types.Trace"},
			"list-async-invokes":                     {"MaxResults": "*int32", "NextToken": "*string", "SortBy": "types.SortAsyncInvocationBy", "SortOrder": "types.SortOrder", "StatusEquals": "types.AsyncInvokeStatus", "SubmitTimeAfter": "*time.Time", "SubmitTimeBefore": "*time.Time"},
			"start-async-invoke":                     {"ClientRequestToken": "*string", "ModelId": "*string", "ModelInput": "document.Interface", "OutputDataConfig": "types.AsyncInvokeOutputDataConfig", "Tags": "[]types.Tag"},
		},
		OperationInputRequired: map[string][]string{
			"apply-guardrail":                        {"Content", "GuardrailIdentifier", "GuardrailVersion", "Source"},
			"converse":                               {"ModelId"},
			"converse-stream":                        {"ModelId"},
			"count-tokens":                           {"Input", "ModelId"},
			"get-async-invoke":                       {"InvocationArn"},
			"invoke-model":                           {"ModelId"},
			"invoke-model-with-bidirectional-stream": {"ModelId"},
			"invoke-model-with-response-stream":      {"ModelId"},
			"list-async-invokes":                     {},
			"start-async-invoke":                     {"ModelId", "ModelInput", "OutputDataConfig"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("bedrockruntime", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
