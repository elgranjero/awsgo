package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/bedrockagentruntime"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// bedrockagentruntimeCmd represents the bedrockagentruntime command
var _bedrockagentruntimeCmd = &cobra.Command{
	Use:   "bedrockagentruntime",
	Short: "AWS bedrockagentruntime CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := bedrockagentruntime.NewFromConfig(cfg)
		if _bedrockagentruntimeCreateInvocation {
			bedrockagentruntime_CreateInvocation(cfg, client)
			return
		}
		if _bedrockagentruntimeCreateSession {
			bedrockagentruntime_CreateSession(cfg, client)
			return
		}
		if _bedrockagentruntimeDeleteAgentMemory {
			bedrockagentruntime_DeleteAgentMemory(cfg, client)
			return
		}
		if _bedrockagentruntimeDeleteSession {
			bedrockagentruntime_DeleteSession(cfg, client)
			return
		}
		if _bedrockagentruntimeEndSession {
			bedrockagentruntime_EndSession(cfg, client)
			return
		}
		if _bedrockagentruntimeGenerateQuery {
			bedrockagentruntime_GenerateQuery(cfg, client)
			return
		}
		if _bedrockagentruntimeGetAgentMemory {
			bedrockagentruntime_GetAgentMemory(cfg, client)
			return
		}
		if _bedrockagentruntimeGetExecutionFlowSnapshot {
			bedrockagentruntime_GetExecutionFlowSnapshot(cfg, client)
			return
		}
		if _bedrockagentruntimeGetFlowExecution {
			bedrockagentruntime_GetFlowExecution(cfg, client)
			return
		}
		if _bedrockagentruntimeGetInvocationStep {
			bedrockagentruntime_GetInvocationStep(cfg, client)
			return
		}
		if _bedrockagentruntimeGetSession {
			bedrockagentruntime_GetSession(cfg, client)
			return
		}
		if _bedrockagentruntimeInvokeAgent {
			bedrockagentruntime_InvokeAgent(cfg, client)
			return
		}
		if _bedrockagentruntimeInvokeFlow {
			bedrockagentruntime_InvokeFlow(cfg, client)
			return
		}
		if _bedrockagentruntimeInvokeInlineAgent {
			bedrockagentruntime_InvokeInlineAgent(cfg, client)
			return
		}
		if _bedrockagentruntimeListFlowExecutionEvents {
			bedrockagentruntime_ListFlowExecutionEvents(cfg, client)
			return
		}
		if _bedrockagentruntimeListFlowExecutions {
			bedrockagentruntime_ListFlowExecutions(cfg, client)
			return
		}
		if _bedrockagentruntimeListInvocationSteps {
			bedrockagentruntime_ListInvocationSteps(cfg, client)
			return
		}
		if _bedrockagentruntimeListInvocations {
			bedrockagentruntime_ListInvocations(cfg, client)
			return
		}
		if _bedrockagentruntimeListSessions {
			bedrockagentruntime_ListSessions(cfg, client)
			return
		}
		if _bedrockagentruntimeListTagsForResource {
			bedrockagentruntime_ListTagsForResource(cfg, client)
			return
		}
		if _bedrockagentruntimeOptimizePrompt {
			bedrockagentruntime_OptimizePrompt(cfg, client)
			return
		}
		if _bedrockagentruntimePutInvocationStep {
			bedrockagentruntime_PutInvocationStep(cfg, client)
			return
		}
		if _bedrockagentruntimeRerank {
			bedrockagentruntime_Rerank(cfg, client)
			return
		}
		if _bedrockagentruntimeRetrieve {
			bedrockagentruntime_Retrieve(cfg, client)
			return
		}
		if _bedrockagentruntimeRetrieveAndGenerate {
			bedrockagentruntime_RetrieveAndGenerate(cfg, client)
			return
		}
		if _bedrockagentruntimeRetrieveAndGenerateStream {
			bedrockagentruntime_RetrieveAndGenerateStream(cfg, client)
			return
		}
		if _bedrockagentruntimeStartFlowExecution {
			bedrockagentruntime_StartFlowExecution(cfg, client)
			return
		}
		if _bedrockagentruntimeStopFlowExecution {
			bedrockagentruntime_StopFlowExecution(cfg, client)
			return
		}
		if _bedrockagentruntimeTagResource {
			bedrockagentruntime_TagResource(cfg, client)
			return
		}
		if _bedrockagentruntimeUntagResource {
			bedrockagentruntime_UntagResource(cfg, client)
			return
		}
		if _bedrockagentruntimeUpdateSession {
			bedrockagentruntime_UpdateSession(cfg, client)
			return
		}

	},
}

var (
	_bedrockagentruntimeCreateInvocation          bool
	_bedrockagentruntimeCreateSession             bool
	_bedrockagentruntimeDeleteAgentMemory         bool
	_bedrockagentruntimeDeleteSession             bool
	_bedrockagentruntimeEndSession                bool
	_bedrockagentruntimeGenerateQuery             bool
	_bedrockagentruntimeGetAgentMemory            bool
	_bedrockagentruntimeGetExecutionFlowSnapshot  bool
	_bedrockagentruntimeGetFlowExecution          bool
	_bedrockagentruntimeGetInvocationStep         bool
	_bedrockagentruntimeGetSession                bool
	_bedrockagentruntimeInvokeAgent               bool
	_bedrockagentruntimeInvokeFlow                bool
	_bedrockagentruntimeInvokeInlineAgent         bool
	_bedrockagentruntimeListFlowExecutionEvents   bool
	_bedrockagentruntimeListFlowExecutions        bool
	_bedrockagentruntimeListInvocationSteps       bool
	_bedrockagentruntimeListInvocations           bool
	_bedrockagentruntimeListSessions              bool
	_bedrockagentruntimeListTagsForResource       bool
	_bedrockagentruntimeOptimizePrompt            bool
	_bedrockagentruntimePutInvocationStep         bool
	_bedrockagentruntimeRerank                    bool
	_bedrockagentruntimeRetrieve                  bool
	_bedrockagentruntimeRetrieveAndGenerate       bool
	_bedrockagentruntimeRetrieveAndGenerateStream bool
	_bedrockagentruntimeStartFlowExecution        bool
	_bedrockagentruntimeStopFlowExecution         bool
	_bedrockagentruntimeTagResource               bool
	_bedrockagentruntimeUntagResource             bool
	_bedrockagentruntimeUpdateSession             bool

	_bedrockagentruntimeActionGroups                     string
	_bedrockagentruntimeAgentAliasId                     string
	_bedrockagentruntimeAgentCollaboration               string
	_bedrockagentruntimeAgentId                          string
	_bedrockagentruntimeAgentName                        string
	_bedrockagentruntimeBedrockModelConfigurations       string
	_bedrockagentruntimeCollaboratorConfigurations       string
	_bedrockagentruntimeCollaborators                    string
	_bedrockagentruntimeCustomOrchestration              string
	_bedrockagentruntimeCustomerEncryptionKeyArn         string
	_bedrockagentruntimeDescription                      string
	_bedrockagentruntimeEnableTrace                      string
	_bedrockagentruntimeEncryptionKeyArn                 string
	_bedrockagentruntimeEventType                        string
	_bedrockagentruntimeExecutionId                      string
	_bedrockagentruntimeExecutionIdentifier              string
	_bedrockagentruntimeFlowAliasIdentifier              string
	_bedrockagentruntimeFlowExecutionName                string
	_bedrockagentruntimeFlowIdentifier                   string
	_bedrockagentruntimeFoundationModel                  string
	_bedrockagentruntimeGuardrailConfiguration           string
	_bedrockagentruntimeIdleSessionTTLInSeconds          string
	_bedrockagentruntimeInlineSessionState               string
	_bedrockagentruntimeInput                            string
	_bedrockagentruntimeInputText                        string
	_bedrockagentruntimeInputs                           string
	_bedrockagentruntimeInstruction                      string
	_bedrockagentruntimeInvocationId                     string
	_bedrockagentruntimeInvocationIdentifier             string
	_bedrockagentruntimeInvocationStepId                 string
	_bedrockagentruntimeInvocationStepTime               string
	_bedrockagentruntimeKnowledgeBaseId                  string
	_bedrockagentruntimeKnowledgeBases                   string
	_bedrockagentruntimeMaxItems                         string
	_bedrockagentruntimeMaxResults                       string
	_bedrockagentruntimeMemoryId                         string
	_bedrockagentruntimeMemoryType                       string
	_bedrockagentruntimeModelPerformanceConfiguration    string
	_bedrockagentruntimeNextToken                        string
	_bedrockagentruntimeOrchestrationType                string
	_bedrockagentruntimePayload                          string
	_bedrockagentruntimePromptCreationConfigurations     string
	_bedrockagentruntimePromptOverrideConfiguration      string
	_bedrockagentruntimeQueries                          string
	_bedrockagentruntimeQueryGenerationInput             string
	_bedrockagentruntimeRerankingConfiguration           string
	_bedrockagentruntimeResourceArn                      string
	_bedrockagentruntimeRetrievalConfiguration           string
	_bedrockagentruntimeRetrievalQuery                   string
	_bedrockagentruntimeRetrieveAndGenerateConfiguration string
	_bedrockagentruntimeSessionConfiguration             string
	_bedrockagentruntimeSessionId                        string
	_bedrockagentruntimeSessionIdentifier                string
	_bedrockagentruntimeSessionMetadata                  string
	_bedrockagentruntimeSessionState                     string
	_bedrockagentruntimeSourceArn                        string
	_bedrockagentruntimeSources                          string
	_bedrockagentruntimeStreamingConfigurations          string
	_bedrockagentruntimeTagKeys                          []string
	_bedrockagentruntimeTags                             string
	_bedrockagentruntimeTargetModelId                    string
	_bedrockagentruntimeTransformationConfiguration      string
)

// Creates a new invocation within a session. An invocation groups the related
// invocation steps that store the content from a conversation. For more
// information about sessions, see [Store and retrieve conversation history and context with Amazon Bedrock sessions].
//
// # Related APIs
//
// [ListInvocations]
//
// [ListSessions]
//
// [GetSession]
//
// [Store and retrieve conversation history and context with Amazon Bedrock sessions]: https://docs.aws.amazon.com/bedrock/latest/userguide/sessions.html
// [ListInvocations]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent-runtime_ListInvocations.html
// [ListSessions]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent-runtime_ListSessions.html
// [GetSession]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent-runtime_GetSession.html
func bedrockagentruntime_CreateInvocation(cfg aws.Config, client *bedrockagentruntime.Client) {
	input := &bedrockagentruntime.CreateInvocationInput{
		// SessionIdentifier: *string, // Required
	}

	if len(_bedrockagentruntimeSessionIdentifier) > 0 {
		input.SessionIdentifier = aws.String(_bedrockagentruntimeSessionIdentifier)
	}
	if len(_bedrockagentruntimeDescription) > 0 {
		input.Description = aws.String(_bedrockagentruntimeDescription)
	}
	if len(_bedrockagentruntimeInvocationId) > 0 {
		input.InvocationId = aws.String(_bedrockagentruntimeInvocationId)
	}

	if resp, err := client.CreateInvocation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a session to temporarily store conversations for generative AI (GenAI)
// applications built with open-source frameworks such as LangGraph and LlamaIndex.
// Sessions enable you to save the state of conversations at checkpoints, with the
// added security and infrastructure of Amazon Web Services. For more information,
// see [Store and retrieve conversation history and context with Amazon Bedrock sessions].
//
// By default, Amazon Bedrock uses Amazon Web Services-managed keys for session
// encryption, including session metadata, or you can use your own KMS key. For
// more information, see [Amazon Bedrock session encryption].
//
// You use a session to store state and conversation history for generative AI
// applications built with open-source frameworks. For Amazon Bedrock Agents, the
// service automatically manages conversation context and associates them with the
// agent-specific sessionId you specify in the [InvokeAgent]API operation.
//
// Related APIs:
//
// [ListSessions]
//
// [GetSession]
//
// [EndSession]
//
// [DeleteSession]
//
// [Store and retrieve conversation history and context with Amazon Bedrock sessions]: https://docs.aws.amazon.com/bedrock/latest/userguide/sessions.html
// [DeleteSession]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent-runtime_DeleteSession.html
// [Amazon Bedrock session encryption]: https://docs.aws.amazon.com/bedrock/latest/userguide/session-encryption.html
// [InvokeAgent]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent-runtime_InvokeAgent.html
// [EndSession]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent-runtime_EndSession.html
// [ListSessions]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent-runtime_ListSessions.html
// [GetSession]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent-runtime_GetSession.html
func bedrockagentruntime_CreateSession(cfg aws.Config, client *bedrockagentruntime.Client) {
	input := &bedrockagentruntime.CreateSessionInput{}

	if len(_bedrockagentruntimeEncryptionKeyArn) > 0 {
		input.EncryptionKeyArn = aws.String(_bedrockagentruntimeEncryptionKeyArn)
	}
	if len(_bedrockagentruntimeSessionMetadata) > 0 {
		if err := assignInputField(input, "SessionMetadata", _bedrockagentruntimeSessionMetadata); err != nil {
			log.Errorf("invalid --session-metadata: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentruntimeTags) > 0 {
		if err := assignInputField(input, "Tags", _bedrockagentruntimeTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes memory from the specified memory identifier.
func bedrockagentruntime_DeleteAgentMemory(cfg aws.Config, client *bedrockagentruntime.Client) {
	input := &bedrockagentruntime.DeleteAgentMemoryInput{
		// AgentAliasId: *string, // Required
		// AgentId: *string, // Required
	}

	if len(_bedrockagentruntimeAgentAliasId) > 0 {
		input.AgentAliasId = aws.String(_bedrockagentruntimeAgentAliasId)
	}
	if len(_bedrockagentruntimeAgentId) > 0 {
		input.AgentId = aws.String(_bedrockagentruntimeAgentId)
	}
	if len(_bedrockagentruntimeMemoryId) > 0 {
		input.MemoryId = aws.String(_bedrockagentruntimeMemoryId)
	}
	if len(_bedrockagentruntimeSessionId) > 0 {
		input.SessionId = aws.String(_bedrockagentruntimeSessionId)
	}

	if resp, err := client.DeleteAgentMemory(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a session that you ended. You can't delete a session with an ACTIVE
// status. To delete an active session, you must first end it with the [EndSession]API
// operation. For more information about sessions, see [Store and retrieve conversation history and context with Amazon Bedrock sessions].
//
// [Store and retrieve conversation history and context with Amazon Bedrock sessions]: https://docs.aws.amazon.com/bedrock/latest/userguide/sessions.html
// [EndSession]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent-runtime_EndSession.html
func bedrockagentruntime_DeleteSession(cfg aws.Config, client *bedrockagentruntime.Client) {
	input := &bedrockagentruntime.DeleteSessionInput{
		// SessionIdentifier: *string, // Required
	}

	if len(_bedrockagentruntimeSessionIdentifier) > 0 {
		input.SessionIdentifier = aws.String(_bedrockagentruntimeSessionIdentifier)
	}

	if resp, err := client.DeleteSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Ends the session. After you end a session, you can still access its content but
// you can’t add to it. To delete the session and it's content, you use the
// DeleteSession API operation. For more information about sessions, see [Store and retrieve conversation history and context with Amazon Bedrock sessions].
//
// [Store and retrieve conversation history and context with Amazon Bedrock sessions]: https://docs.aws.amazon.com/bedrock/latest/userguide/sessions.html
func bedrockagentruntime_EndSession(cfg aws.Config, client *bedrockagentruntime.Client) {
	input := &bedrockagentruntime.EndSessionInput{
		// SessionIdentifier: *string, // Required
	}

	if len(_bedrockagentruntimeSessionIdentifier) > 0 {
		input.SessionIdentifier = aws.String(_bedrockagentruntimeSessionIdentifier)
	}

	if resp, err := client.EndSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Generates an SQL query from a natural language query. For more information, see [Generate a query for structured data]
// in the Amazon Bedrock User Guide.
//
// [Generate a query for structured data]: https://docs.aws.amazon.com/bedrock/latest/userguide/knowledge-base-generate-query.html
func bedrockagentruntime_GenerateQuery(cfg aws.Config, client *bedrockagentruntime.Client) {
	input := &bedrockagentruntime.GenerateQueryInput{
		// QueryGenerationInput: *types.QueryGenerationInput, // Required
		// TransformationConfiguration: *types.TransformationConfiguration, // Required
	}

	if len(_bedrockagentruntimeQueryGenerationInput) > 0 {
		if err := assignInputField(input, "QueryGenerationInput", _bedrockagentruntimeQueryGenerationInput); err != nil {
			log.Errorf("invalid --query-generation-input: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentruntimeTransformationConfiguration) > 0 {
		if err := assignInputField(input, "TransformationConfiguration", _bedrockagentruntimeTransformationConfiguration); err != nil {
			log.Errorf("invalid --transformation-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.GenerateQuery(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the sessions stored in the memory of the agent.
func bedrockagentruntime_GetAgentMemory(cfg aws.Config, client *bedrockagentruntime.Client) {
	input := &bedrockagentruntime.GetAgentMemoryInput{
		// AgentAliasId: *string, // Required
		// AgentId: *string, // Required
		// MemoryId: *string, // Required
		// MemoryType: types.MemoryType, // Required
	}

	if len(_bedrockagentruntimeAgentAliasId) > 0 {
		input.AgentAliasId = aws.String(_bedrockagentruntimeAgentAliasId)
	}
	if len(_bedrockagentruntimeAgentId) > 0 {
		input.AgentId = aws.String(_bedrockagentruntimeAgentId)
	}
	if len(_bedrockagentruntimeMemoryId) > 0 {
		input.MemoryId = aws.String(_bedrockagentruntimeMemoryId)
	}
	if len(_bedrockagentruntimeMemoryType) > 0 {
		if err := assignInputField(input, "MemoryType", _bedrockagentruntimeMemoryType); err != nil {
			log.Errorf("invalid --memory-type: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentruntimeMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _bedrockagentruntimeMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentruntimeNextToken) > 0 {
		input.NextToken = aws.String(_bedrockagentruntimeNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetAgentMemory(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrockagentruntime.GetAgentMemoryOutput
	p := bedrockagentruntime.NewGetAgentMemoryPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Retrieves the flow definition snapshot used for a flow execution. The snapshot
// represents the flow metadata and definition as it existed at the time the
// execution was started. Note that even if the flow is edited after an execution
// starts, the snapshot connected to the execution remains unchanged.
//
// Flow executions is in preview release for Amazon Bedrock and is subject to
// change.
func bedrockagentruntime_GetExecutionFlowSnapshot(cfg aws.Config, client *bedrockagentruntime.Client) {
	input := &bedrockagentruntime.GetExecutionFlowSnapshotInput{
		// ExecutionIdentifier: *string, // Required
		// FlowAliasIdentifier: *string, // Required
		// FlowIdentifier: *string, // Required
	}

	if len(_bedrockagentruntimeExecutionIdentifier) > 0 {
		input.ExecutionIdentifier = aws.String(_bedrockagentruntimeExecutionIdentifier)
	}
	if len(_bedrockagentruntimeFlowAliasIdentifier) > 0 {
		input.FlowAliasIdentifier = aws.String(_bedrockagentruntimeFlowAliasIdentifier)
	}
	if len(_bedrockagentruntimeFlowIdentifier) > 0 {
		input.FlowIdentifier = aws.String(_bedrockagentruntimeFlowIdentifier)
	}

	if resp, err := client.GetExecutionFlowSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves details about a specific flow execution, including its status, start
// and end times, and any errors that occurred during execution.
func bedrockagentruntime_GetFlowExecution(cfg aws.Config, client *bedrockagentruntime.Client) {
	input := &bedrockagentruntime.GetFlowExecutionInput{
		// ExecutionIdentifier: *string, // Required
		// FlowAliasIdentifier: *string, // Required
		// FlowIdentifier: *string, // Required
	}

	if len(_bedrockagentruntimeExecutionIdentifier) > 0 {
		input.ExecutionIdentifier = aws.String(_bedrockagentruntimeExecutionIdentifier)
	}
	if len(_bedrockagentruntimeFlowAliasIdentifier) > 0 {
		input.FlowAliasIdentifier = aws.String(_bedrockagentruntimeFlowAliasIdentifier)
	}
	if len(_bedrockagentruntimeFlowIdentifier) > 0 {
		input.FlowIdentifier = aws.String(_bedrockagentruntimeFlowIdentifier)
	}

	if resp, err := client.GetFlowExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the details of a specific invocation step within an invocation in a
// session. For more information about sessions, see [Store and retrieve conversation history and context with Amazon Bedrock sessions].
//
// [Store and retrieve conversation history and context with Amazon Bedrock sessions]: https://docs.aws.amazon.com/bedrock/latest/userguide/sessions.html
func bedrockagentruntime_GetInvocationStep(cfg aws.Config, client *bedrockagentruntime.Client) {
	input := &bedrockagentruntime.GetInvocationStepInput{
		// InvocationIdentifier: *string, // Required
		// InvocationStepId: *string, // Required
		// SessionIdentifier: *string, // Required
	}

	if len(_bedrockagentruntimeInvocationIdentifier) > 0 {
		input.InvocationIdentifier = aws.String(_bedrockagentruntimeInvocationIdentifier)
	}
	if len(_bedrockagentruntimeInvocationStepId) > 0 {
		input.InvocationStepId = aws.String(_bedrockagentruntimeInvocationStepId)
	}
	if len(_bedrockagentruntimeSessionIdentifier) > 0 {
		input.SessionIdentifier = aws.String(_bedrockagentruntimeSessionIdentifier)
	}

	if resp, err := client.GetInvocationStep(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves details about a specific session. For more information about
// sessions, see [Store and retrieve conversation history and context with Amazon Bedrock sessions].
//
// [Store and retrieve conversation history and context with Amazon Bedrock sessions]: https://docs.aws.amazon.com/bedrock/latest/userguide/sessions.html
func bedrockagentruntime_GetSession(cfg aws.Config, client *bedrockagentruntime.Client) {
	input := &bedrockagentruntime.GetSessionInput{
		// SessionIdentifier: *string, // Required
	}

	if len(_bedrockagentruntimeSessionIdentifier) > 0 {
		input.SessionIdentifier = aws.String(_bedrockagentruntimeSessionIdentifier)
	}

	if resp, err := client.GetSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sends a prompt for the agent to process and respond to. Note the following
// fields for the request:
//
// - To continue the same conversation with an agent, use the same sessionId
// value in the request.
//
// - To activate trace enablement, turn enableTrace to true . Trace enablement
// helps you follow the agent's reasoning process that led it to the information it
// processed, the actions it took, and the final result it yielded. For more
// information, see [Trace enablement].
//
// - End a conversation by setting endSession to true .
//
// - In the sessionState object, you can include attributes for the session or
// prompt or, if you configured an action group to return control, results from
// invocation of the action group.
//
// The response contains both chunk and trace attributes.
//
// The final response is returned in the bytes field of the chunk object. The
// InvokeAgent returns one chunk for the entire interaction.
//
// - The attribution object contains citations for parts of the response.
//
// - If you set enableTrace to true in the request, you can trace the agent's
// steps and reasoning process that led it to the response.
//
// - If the action predicted was configured to return control, the response
// returns parameters for the action, elicited from the user, in the
// returnControl field.
//
// - Errors are also surfaced in the response.
//
// [Trace enablement]: https://docs.aws.amazon.com/bedrock/latest/userguide/agents-test.html#trace-events
func bedrockagentruntime_InvokeAgent(cfg aws.Config, client *bedrockagentruntime.Client) {
	input := &bedrockagentruntime.InvokeAgentInput{
		// AgentAliasId: *string, // Required
		// AgentId: *string, // Required
		// SessionId: *string, // Required
	}

	if len(_bedrockagentruntimeAgentAliasId) > 0 {
		input.AgentAliasId = aws.String(_bedrockagentruntimeAgentAliasId)
	}
	if len(_bedrockagentruntimeAgentId) > 0 {
		input.AgentId = aws.String(_bedrockagentruntimeAgentId)
	}
	if len(_bedrockagentruntimeSessionId) > 0 {
		input.SessionId = aws.String(_bedrockagentruntimeSessionId)
	}
	if len(_bedrockagentruntimeBedrockModelConfigurations) > 0 {
		if err := assignInputField(input, "BedrockModelConfigurations", _bedrockagentruntimeBedrockModelConfigurations); err != nil {
			log.Errorf("invalid --bedrock-model-configurations: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentruntimeEnableTrace) > 0 {
		if err := assignInputField(input, "EnableTrace", _bedrockagentruntimeEnableTrace); err != nil {
			log.Errorf("invalid --enable-trace: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentruntimeInputText) > 0 {
		input.InputText = aws.String(_bedrockagentruntimeInputText)
	}
	if len(_bedrockagentruntimeMemoryId) > 0 {
		input.MemoryId = aws.String(_bedrockagentruntimeMemoryId)
	}
	if len(_bedrockagentruntimePromptCreationConfigurations) > 0 {
		if err := assignInputField(input, "PromptCreationConfigurations", _bedrockagentruntimePromptCreationConfigurations); err != nil {
			log.Errorf("invalid --prompt-creation-configurations: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentruntimeSessionState) > 0 {
		if err := assignInputField(input, "SessionState", _bedrockagentruntimeSessionState); err != nil {
			log.Errorf("invalid --session-state: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentruntimeSourceArn) > 0 {
		input.SourceArn = aws.String(_bedrockagentruntimeSourceArn)
	}
	if len(_bedrockagentruntimeStreamingConfigurations) > 0 {
		if err := assignInputField(input, "StreamingConfigurations", _bedrockagentruntimeStreamingConfigurations); err != nil {
			log.Errorf("invalid --streaming-configurations: %s", err.Error())
			return
		}
	}

	if resp, err := client.InvokeAgent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Invokes an alias of a flow to run the inputs that you specify and return the
// output of each node as a stream. If there's an error, the error is returned. For
// more information, see [Test a flow in Amazon Bedrock]in the [Amazon Bedrock User Guide].
//
// The CLI doesn't support streaming operations in Amazon Bedrock, including
// InvokeFlow .
//
// [Test a flow in Amazon Bedrock]: https://docs.aws.amazon.com/bedrock/latest/userguide/flows-test.html
// [Amazon Bedrock User Guide]: https://docs.aws.amazon.com/bedrock/latest/userguide/what-is-service.html
func bedrockagentruntime_InvokeFlow(cfg aws.Config, client *bedrockagentruntime.Client) {
	input := &bedrockagentruntime.InvokeFlowInput{
		// FlowAliasIdentifier: *string, // Required
		// FlowIdentifier: *string, // Required
		// Inputs: []types.FlowInput, // Required
	}

	if len(_bedrockagentruntimeFlowAliasIdentifier) > 0 {
		input.FlowAliasIdentifier = aws.String(_bedrockagentruntimeFlowAliasIdentifier)
	}
	if len(_bedrockagentruntimeFlowIdentifier) > 0 {
		input.FlowIdentifier = aws.String(_bedrockagentruntimeFlowIdentifier)
	}
	if len(_bedrockagentruntimeInputs) > 0 {
		if err := assignInputField(input, "Inputs", _bedrockagentruntimeInputs); err != nil {
			log.Errorf("invalid --inputs: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentruntimeEnableTrace) > 0 {
		if err := assignInputField(input, "EnableTrace", _bedrockagentruntimeEnableTrace); err != nil {
			log.Errorf("invalid --enable-trace: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentruntimeExecutionId) > 0 {
		input.ExecutionId = aws.String(_bedrockagentruntimeExecutionId)
	}
	if len(_bedrockagentruntimeModelPerformanceConfiguration) > 0 {
		if err := assignInputField(input, "ModelPerformanceConfiguration", _bedrockagentruntimeModelPerformanceConfiguration); err != nil {
			log.Errorf("invalid --model-performance-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.InvokeFlow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Invokes an inline Amazon Bedrock agent using the configurations you provide
// with the request.
//
// - Specify the following fields for security purposes.
//
// - (Optional) customerEncryptionKeyArn – The Amazon Resource Name (ARN) of a
// KMS key to encrypt the creation of the agent.
//
// - (Optional) idleSessionTTLinSeconds – Specify the number of seconds for which
// the agent should maintain session information. After this time expires, the
// subsequent InvokeInlineAgent request begins a new session.
//
// - To override the default prompt behavior for agent orchestration and to use
// advanced prompts, include a promptOverrideConfiguration object. For more
// information, see [Advanced prompts].
//
// - The agent instructions will not be honored if your agent has only one
// knowledge base, uses default prompts, has no action group, and user input is
// disabled.
//
// [Advanced prompts]: https://docs.aws.amazon.com/bedrock/latest/userguide/advanced-prompts.html
func bedrockagentruntime_InvokeInlineAgent(cfg aws.Config, client *bedrockagentruntime.Client) {
	input := &bedrockagentruntime.InvokeInlineAgentInput{
		// FoundationModel: *string, // Required
		// Instruction: *string, // Required
		// SessionId: *string, // Required
	}

	if len(_bedrockagentruntimeFoundationModel) > 0 {
		input.FoundationModel = aws.String(_bedrockagentruntimeFoundationModel)
	}
	if len(_bedrockagentruntimeInstruction) > 0 {
		input.Instruction = aws.String(_bedrockagentruntimeInstruction)
	}
	if len(_bedrockagentruntimeSessionId) > 0 {
		input.SessionId = aws.String(_bedrockagentruntimeSessionId)
	}
	if len(_bedrockagentruntimeActionGroups) > 0 {
		if err := assignInputField(input, "ActionGroups", _bedrockagentruntimeActionGroups); err != nil {
			log.Errorf("invalid --action-groups: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentruntimeAgentCollaboration) > 0 {
		if err := assignInputField(input, "AgentCollaboration", _bedrockagentruntimeAgentCollaboration); err != nil {
			log.Errorf("invalid --agent-collaboration: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentruntimeAgentName) > 0 {
		input.AgentName = aws.String(_bedrockagentruntimeAgentName)
	}
	if len(_bedrockagentruntimeBedrockModelConfigurations) > 0 {
		if err := assignInputField(input, "BedrockModelConfigurations", _bedrockagentruntimeBedrockModelConfigurations); err != nil {
			log.Errorf("invalid --bedrock-model-configurations: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentruntimeCollaboratorConfigurations) > 0 {
		if err := assignInputField(input, "CollaboratorConfigurations", _bedrockagentruntimeCollaboratorConfigurations); err != nil {
			log.Errorf("invalid --collaborator-configurations: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentruntimeCollaborators) > 0 {
		if err := assignInputField(input, "Collaborators", _bedrockagentruntimeCollaborators); err != nil {
			log.Errorf("invalid --collaborators: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentruntimeCustomOrchestration) > 0 {
		if err := assignInputField(input, "CustomOrchestration", _bedrockagentruntimeCustomOrchestration); err != nil {
			log.Errorf("invalid --custom-orchestration: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentruntimeCustomerEncryptionKeyArn) > 0 {
		input.CustomerEncryptionKeyArn = aws.String(_bedrockagentruntimeCustomerEncryptionKeyArn)
	}
	if len(_bedrockagentruntimeEnableTrace) > 0 {
		if err := assignInputField(input, "EnableTrace", _bedrockagentruntimeEnableTrace); err != nil {
			log.Errorf("invalid --enable-trace: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentruntimeGuardrailConfiguration) > 0 {
		if err := assignInputField(input, "GuardrailConfiguration", _bedrockagentruntimeGuardrailConfiguration); err != nil {
			log.Errorf("invalid --guardrail-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentruntimeIdleSessionTTLInSeconds) > 0 {
		if err := assignInputField(input, "IdleSessionTTLInSeconds", _bedrockagentruntimeIdleSessionTTLInSeconds); err != nil {
			log.Errorf("invalid --idle-session-ttlin-seconds: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentruntimeInlineSessionState) > 0 {
		if err := assignInputField(input, "InlineSessionState", _bedrockagentruntimeInlineSessionState); err != nil {
			log.Errorf("invalid --inline-session-state: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentruntimeInputText) > 0 {
		input.InputText = aws.String(_bedrockagentruntimeInputText)
	}
	if len(_bedrockagentruntimeKnowledgeBases) > 0 {
		if err := assignInputField(input, "KnowledgeBases", _bedrockagentruntimeKnowledgeBases); err != nil {
			log.Errorf("invalid --knowledge-bases: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentruntimeOrchestrationType) > 0 {
		if err := assignInputField(input, "OrchestrationType", _bedrockagentruntimeOrchestrationType); err != nil {
			log.Errorf("invalid --orchestration-type: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentruntimePromptCreationConfigurations) > 0 {
		if err := assignInputField(input, "PromptCreationConfigurations", _bedrockagentruntimePromptCreationConfigurations); err != nil {
			log.Errorf("invalid --prompt-creation-configurations: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentruntimePromptOverrideConfiguration) > 0 {
		if err := assignInputField(input, "PromptOverrideConfiguration", _bedrockagentruntimePromptOverrideConfiguration); err != nil {
			log.Errorf("invalid --prompt-override-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentruntimeStreamingConfigurations) > 0 {
		if err := assignInputField(input, "StreamingConfigurations", _bedrockagentruntimeStreamingConfigurations); err != nil {
			log.Errorf("invalid --streaming-configurations: %s", err.Error())
			return
		}
	}

	if resp, err := client.InvokeInlineAgent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists events that occurred during a flow execution. Events provide detailed
// information about the execution progress, including node inputs and outputs,
// flow inputs and outputs, condition results, and failure events.
//
// Flow executions is in preview release for Amazon Bedrock and is subject to
// change.
func bedrockagentruntime_ListFlowExecutionEvents(cfg aws.Config, client *bedrockagentruntime.Client) {
	input := &bedrockagentruntime.ListFlowExecutionEventsInput{
		// EventType: types.FlowExecutionEventType, // Required
		// ExecutionIdentifier: *string, // Required
		// FlowAliasIdentifier: *string, // Required
		// FlowIdentifier: *string, // Required
	}

	if len(_bedrockagentruntimeEventType) > 0 {
		if err := assignInputField(input, "EventType", _bedrockagentruntimeEventType); err != nil {
			log.Errorf("invalid --event-type: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentruntimeExecutionIdentifier) > 0 {
		input.ExecutionIdentifier = aws.String(_bedrockagentruntimeExecutionIdentifier)
	}
	if len(_bedrockagentruntimeFlowAliasIdentifier) > 0 {
		input.FlowAliasIdentifier = aws.String(_bedrockagentruntimeFlowAliasIdentifier)
	}
	if len(_bedrockagentruntimeFlowIdentifier) > 0 {
		input.FlowIdentifier = aws.String(_bedrockagentruntimeFlowIdentifier)
	}
	if len(_bedrockagentruntimeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bedrockagentruntimeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentruntimeNextToken) > 0 {
		input.NextToken = aws.String(_bedrockagentruntimeNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListFlowExecutionEvents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrockagentruntime.ListFlowExecutionEventsOutput
	p := bedrockagentruntime.NewListFlowExecutionEventsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Lists all executions of a flow. Results can be paginated and include summary
// information about each execution, such as status, start and end times, and the
// execution's Amazon Resource Name (ARN).
//
// Flow executions is in preview release for Amazon Bedrock and is subject to
// change.
func bedrockagentruntime_ListFlowExecutions(cfg aws.Config, client *bedrockagentruntime.Client) {
	input := &bedrockagentruntime.ListFlowExecutionsInput{
		// FlowIdentifier: *string, // Required
	}

	if len(_bedrockagentruntimeFlowIdentifier) > 0 {
		input.FlowIdentifier = aws.String(_bedrockagentruntimeFlowIdentifier)
	}
	if len(_bedrockagentruntimeFlowAliasIdentifier) > 0 {
		input.FlowAliasIdentifier = aws.String(_bedrockagentruntimeFlowAliasIdentifier)
	}
	if len(_bedrockagentruntimeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bedrockagentruntimeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentruntimeNextToken) > 0 {
		input.NextToken = aws.String(_bedrockagentruntimeNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListFlowExecutions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrockagentruntime.ListFlowExecutionsOutput
	p := bedrockagentruntime.NewListFlowExecutionsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Lists all invocation steps associated with a session and optionally, an
// invocation within the session. For more information about sessions, see [Store and retrieve conversation history and context with Amazon Bedrock sessions].
//
// [Store and retrieve conversation history and context with Amazon Bedrock sessions]: https://docs.aws.amazon.com/bedrock/latest/userguide/sessions.html
func bedrockagentruntime_ListInvocationSteps(cfg aws.Config, client *bedrockagentruntime.Client) {
	input := &bedrockagentruntime.ListInvocationStepsInput{
		// SessionIdentifier: *string, // Required
	}

	if len(_bedrockagentruntimeSessionIdentifier) > 0 {
		input.SessionIdentifier = aws.String(_bedrockagentruntimeSessionIdentifier)
	}
	if len(_bedrockagentruntimeInvocationIdentifier) > 0 {
		input.InvocationIdentifier = aws.String(_bedrockagentruntimeInvocationIdentifier)
	}
	if len(_bedrockagentruntimeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bedrockagentruntimeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentruntimeNextToken) > 0 {
		input.NextToken = aws.String(_bedrockagentruntimeNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListInvocationSteps(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrockagentruntime.ListInvocationStepsOutput
	p := bedrockagentruntime.NewListInvocationStepsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Lists all invocations associated with a specific session. For more information
// about sessions, see [Store and retrieve conversation history and context with Amazon Bedrock sessions].
//
// [Store and retrieve conversation history and context with Amazon Bedrock sessions]: https://docs.aws.amazon.com/bedrock/latest/userguide/sessions.html
func bedrockagentruntime_ListInvocations(cfg aws.Config, client *bedrockagentruntime.Client) {
	input := &bedrockagentruntime.ListInvocationsInput{
		// SessionIdentifier: *string, // Required
	}

	if len(_bedrockagentruntimeSessionIdentifier) > 0 {
		input.SessionIdentifier = aws.String(_bedrockagentruntimeSessionIdentifier)
	}
	if len(_bedrockagentruntimeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bedrockagentruntimeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentruntimeNextToken) > 0 {
		input.NextToken = aws.String(_bedrockagentruntimeNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListInvocations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrockagentruntime.ListInvocationsOutput
	p := bedrockagentruntime.NewListInvocationsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Lists all sessions in your Amazon Web Services account. For more information
// about sessions, see [Store and retrieve conversation history and context with Amazon Bedrock sessions].
//
// [Store and retrieve conversation history and context with Amazon Bedrock sessions]: https://docs.aws.amazon.com/bedrock/latest/userguide/sessions.html
func bedrockagentruntime_ListSessions(cfg aws.Config, client *bedrockagentruntime.Client) {
	input := &bedrockagentruntime.ListSessionsInput{}

	if len(_bedrockagentruntimeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bedrockagentruntimeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentruntimeNextToken) > 0 {
		input.NextToken = aws.String(_bedrockagentruntimeNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSessions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrockagentruntime.ListSessionsOutput
	p := bedrockagentruntime.NewListSessionsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// List all the tags for the resource you specify.
func bedrockagentruntime_ListTagsForResource(cfg aws.Config, client *bedrockagentruntime.Client) {
	input := &bedrockagentruntime.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_bedrockagentruntimeResourceArn) > 0 {
		input.ResourceArn = aws.String(_bedrockagentruntimeResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Optimizes a prompt for the task that you specify. For more information, see [Optimize a prompt] in
// the [Amazon Bedrock User Guide].
//
// [Amazon Bedrock User Guide]: https://docs.aws.amazon.com/bedrock/latest/userguide/what-is-service.html
// [Optimize a prompt]: https://docs.aws.amazon.com/bedrock/latest/userguide/prompt-management-optimize.html
func bedrockagentruntime_OptimizePrompt(cfg aws.Config, client *bedrockagentruntime.Client) {
	input := &bedrockagentruntime.OptimizePromptInput{
		// Input: types.InputPrompt, // Required
		// TargetModelId: *string, // Required
	}

	if len(_bedrockagentruntimeInput) > 0 {
		if err := assignInputField(input, "Input", _bedrockagentruntimeInput); err != nil {
			log.Errorf("invalid --input: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentruntimeTargetModelId) > 0 {
		input.TargetModelId = aws.String(_bedrockagentruntimeTargetModelId)
	}

	if resp, err := client.OptimizePrompt(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Add an invocation step to an invocation in a session. An invocation step stores
// fine-grained state checkpoints, including text and images, for each interaction.
// For more information about sessions, see [Store and retrieve conversation history and context with Amazon Bedrock sessions].
//
// Related APIs:
//
// [GetInvocationStep]
//
// [ListInvocationSteps]
//
// [ListInvocations]
//
// [ListSessions]
//
// [Store and retrieve conversation history and context with Amazon Bedrock sessions]: https://docs.aws.amazon.com/bedrock/latest/userguide/sessions.html
// [GetInvocationStep]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent-runtime_GetInvocationStep.html
// [ListInvocationSteps]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent-runtime_ListInvocationSteps.html
// [ListInvocations]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent-runtime_ListInvocations.html
// [ListSessions]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent-runtime_ListInvocations.html
func bedrockagentruntime_PutInvocationStep(cfg aws.Config, client *bedrockagentruntime.Client) {
	input := &bedrockagentruntime.PutInvocationStepInput{
		// InvocationIdentifier: *string, // Required
		// InvocationStepTime: *time.Time, // Required
		// Payload: types.InvocationStepPayload, // Required
		// SessionIdentifier: *string, // Required
	}

	if len(_bedrockagentruntimeInvocationIdentifier) > 0 {
		input.InvocationIdentifier = aws.String(_bedrockagentruntimeInvocationIdentifier)
	}
	if len(_bedrockagentruntimeInvocationStepTime) > 0 {
		if err := assignInputField(input, "InvocationStepTime", _bedrockagentruntimeInvocationStepTime); err != nil {
			log.Errorf("invalid --invocation-step-time: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentruntimePayload) > 0 {
		if err := assignInputField(input, "Payload", _bedrockagentruntimePayload); err != nil {
			log.Errorf("invalid --payload: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentruntimeSessionIdentifier) > 0 {
		input.SessionIdentifier = aws.String(_bedrockagentruntimeSessionIdentifier)
	}
	if len(_bedrockagentruntimeInvocationStepId) > 0 {
		input.InvocationStepId = aws.String(_bedrockagentruntimeInvocationStepId)
	}

	if resp, err := client.PutInvocationStep(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Reranks the relevance of sources based on queries. For more information, see [Improve the relevance of query responses with a reranker model].
//
// [Improve the relevance of query responses with a reranker model]: https://docs.aws.amazon.com/bedrock/latest/userguide/rerank.html
func bedrockagentruntime_Rerank(cfg aws.Config, client *bedrockagentruntime.Client) {
	input := &bedrockagentruntime.RerankInput{
		// Queries: []types.RerankQuery, // Required
		// RerankingConfiguration: *types.RerankingConfiguration, // Required
		// Sources: []types.RerankSource, // Required
	}

	if len(_bedrockagentruntimeQueries) > 0 {
		if err := assignInputField(input, "Queries", _bedrockagentruntimeQueries); err != nil {
			log.Errorf("invalid --queries: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentruntimeRerankingConfiguration) > 0 {
		if err := assignInputField(input, "RerankingConfiguration", _bedrockagentruntimeRerankingConfiguration); err != nil {
			log.Errorf("invalid --reranking-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentruntimeSources) > 0 {
		if err := assignInputField(input, "Sources", _bedrockagentruntimeSources); err != nil {
			log.Errorf("invalid --sources: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentruntimeNextToken) > 0 {
		input.NextToken = aws.String(_bedrockagentruntimeNextToken)
	}

	if disablePaginator() {
		if resp, err := client.Rerank(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrockagentruntime.RerankOutput
	p := bedrockagentruntime.NewRerankPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Queries a knowledge base and retrieves information from it.
func bedrockagentruntime_Retrieve(cfg aws.Config, client *bedrockagentruntime.Client) {
	input := &bedrockagentruntime.RetrieveInput{
		// KnowledgeBaseId: *string, // Required
		// RetrievalQuery: *types.KnowledgeBaseQuery, // Required
	}

	if len(_bedrockagentruntimeKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_bedrockagentruntimeKnowledgeBaseId)
	}
	if len(_bedrockagentruntimeRetrievalQuery) > 0 {
		if err := assignInputField(input, "RetrievalQuery", _bedrockagentruntimeRetrievalQuery); err != nil {
			log.Errorf("invalid --retrieval-query: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentruntimeGuardrailConfiguration) > 0 {
		if err := assignInputField(input, "GuardrailConfiguration", _bedrockagentruntimeGuardrailConfiguration); err != nil {
			log.Errorf("invalid --guardrail-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentruntimeNextToken) > 0 {
		input.NextToken = aws.String(_bedrockagentruntimeNextToken)
	}
	if len(_bedrockagentruntimeRetrievalConfiguration) > 0 {
		if err := assignInputField(input, "RetrievalConfiguration", _bedrockagentruntimeRetrievalConfiguration); err != nil {
			log.Errorf("invalid --retrieval-configuration: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.Retrieve(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrockagentruntime.RetrieveOutput
	p := bedrockagentruntime.NewRetrievePaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Queries a knowledge base and generates responses based on the retrieved results
// and using the specified foundation model or [inference profile]. The response only cites sources
// that are relevant to the query.
//
// [inference profile]: https://docs.aws.amazon.com/bedrock/latest/userguide/cross-region-inference.html
func bedrockagentruntime_RetrieveAndGenerate(cfg aws.Config, client *bedrockagentruntime.Client) {
	input := &bedrockagentruntime.RetrieveAndGenerateInput{
		// Input: *types.RetrieveAndGenerateInput, // Required
	}

	if len(_bedrockagentruntimeInput) > 0 {
		if err := assignInputField(input, "Input", _bedrockagentruntimeInput); err != nil {
			log.Errorf("invalid --input: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentruntimeRetrieveAndGenerateConfiguration) > 0 {
		if err := assignInputField(input, "RetrieveAndGenerateConfiguration", _bedrockagentruntimeRetrieveAndGenerateConfiguration); err != nil {
			log.Errorf("invalid --retrieve-and-generate-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentruntimeSessionConfiguration) > 0 {
		if err := assignInputField(input, "SessionConfiguration", _bedrockagentruntimeSessionConfiguration); err != nil {
			log.Errorf("invalid --session-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentruntimeSessionId) > 0 {
		input.SessionId = aws.String(_bedrockagentruntimeSessionId)
	}

	if resp, err := client.RetrieveAndGenerate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Queries a knowledge base and generates responses based on the retrieved
// results, with output in streaming format.
//
// The CLI doesn't support streaming operations in Amazon Bedrock, including
// InvokeModelWithResponseStream .
//
// This operation requires permission for the  bedrock:RetrieveAndGenerate action.
func bedrockagentruntime_RetrieveAndGenerateStream(cfg aws.Config, client *bedrockagentruntime.Client) {
	input := &bedrockagentruntime.RetrieveAndGenerateStreamInput{
		// Input: *types.RetrieveAndGenerateInput, // Required
	}

	if len(_bedrockagentruntimeInput) > 0 {
		if err := assignInputField(input, "Input", _bedrockagentruntimeInput); err != nil {
			log.Errorf("invalid --input: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentruntimeRetrieveAndGenerateConfiguration) > 0 {
		if err := assignInputField(input, "RetrieveAndGenerateConfiguration", _bedrockagentruntimeRetrieveAndGenerateConfiguration); err != nil {
			log.Errorf("invalid --retrieve-and-generate-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentruntimeSessionConfiguration) > 0 {
		if err := assignInputField(input, "SessionConfiguration", _bedrockagentruntimeSessionConfiguration); err != nil {
			log.Errorf("invalid --session-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentruntimeSessionId) > 0 {
		input.SessionId = aws.String(_bedrockagentruntimeSessionId)
	}

	if resp, err := client.RetrieveAndGenerateStream(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts an execution of an Amazon Bedrock flow. Unlike flows that run until
// completion or time out after five minutes, flow executions let you run flows
// asynchronously for longer durations. Flow executions also yield control so that
// your application can perform other tasks.
//
// This operation returns an Amazon Resource Name (ARN) that you can use to track
// and manage your flow execution.
//
// Flow executions is in preview release for Amazon Bedrock and is subject to
// change.
func bedrockagentruntime_StartFlowExecution(cfg aws.Config, client *bedrockagentruntime.Client) {
	input := &bedrockagentruntime.StartFlowExecutionInput{
		// FlowAliasIdentifier: *string, // Required
		// FlowIdentifier: *string, // Required
		// Inputs: []types.FlowInput, // Required
	}

	if len(_bedrockagentruntimeFlowAliasIdentifier) > 0 {
		input.FlowAliasIdentifier = aws.String(_bedrockagentruntimeFlowAliasIdentifier)
	}
	if len(_bedrockagentruntimeFlowIdentifier) > 0 {
		input.FlowIdentifier = aws.String(_bedrockagentruntimeFlowIdentifier)
	}
	if len(_bedrockagentruntimeInputs) > 0 {
		if err := assignInputField(input, "Inputs", _bedrockagentruntimeInputs); err != nil {
			log.Errorf("invalid --inputs: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentruntimeFlowExecutionName) > 0 {
		input.FlowExecutionName = aws.String(_bedrockagentruntimeFlowExecutionName)
	}
	if len(_bedrockagentruntimeModelPerformanceConfiguration) > 0 {
		if err := assignInputField(input, "ModelPerformanceConfiguration", _bedrockagentruntimeModelPerformanceConfiguration); err != nil {
			log.Errorf("invalid --model-performance-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartFlowExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops an Amazon Bedrock flow's execution. This operation prevents further
// processing of the flow and changes the execution status to Aborted .
func bedrockagentruntime_StopFlowExecution(cfg aws.Config, client *bedrockagentruntime.Client) {
	input := &bedrockagentruntime.StopFlowExecutionInput{
		// ExecutionIdentifier: *string, // Required
		// FlowAliasIdentifier: *string, // Required
		// FlowIdentifier: *string, // Required
	}

	if len(_bedrockagentruntimeExecutionIdentifier) > 0 {
		input.ExecutionIdentifier = aws.String(_bedrockagentruntimeExecutionIdentifier)
	}
	if len(_bedrockagentruntimeFlowAliasIdentifier) > 0 {
		input.FlowAliasIdentifier = aws.String(_bedrockagentruntimeFlowAliasIdentifier)
	}
	if len(_bedrockagentruntimeFlowIdentifier) > 0 {
		input.FlowIdentifier = aws.String(_bedrockagentruntimeFlowIdentifier)
	}

	if resp, err := client.StopFlowExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associate tags with a resource. For more information, see [Tagging resources] in the Amazon
// Bedrock User Guide.
//
// [Tagging resources]: https://docs.aws.amazon.com/bedrock/latest/userguide/what-is-service.html
func bedrockagentruntime_TagResource(cfg aws.Config, client *bedrockagentruntime.Client) {
	input := &bedrockagentruntime.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_bedrockagentruntimeResourceArn) > 0 {
		input.ResourceArn = aws.String(_bedrockagentruntimeResourceArn)
	}
	if len(_bedrockagentruntimeTags) > 0 {
		if err := assignInputField(input, "Tags", _bedrockagentruntimeTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.TagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Remove tags from a resource.
func bedrockagentruntime_UntagResource(cfg aws.Config, client *bedrockagentruntime.Client) {
	input := &bedrockagentruntime.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_bedrockagentruntimeResourceArn) > 0 {
		input.ResourceArn = aws.String(_bedrockagentruntimeResourceArn)
	}
	if len(_bedrockagentruntimeTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _bedrockagentruntimeTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the metadata or encryption settings of a session. For more information
// about sessions, see [Store and retrieve conversation history and context with Amazon Bedrock sessions].
//
// [Store and retrieve conversation history and context with Amazon Bedrock sessions]: https://docs.aws.amazon.com/bedrock/latest/userguide/sessions.html
func bedrockagentruntime_UpdateSession(cfg aws.Config, client *bedrockagentruntime.Client) {
	input := &bedrockagentruntime.UpdateSessionInput{
		// SessionIdentifier: *string, // Required
	}

	if len(_bedrockagentruntimeSessionIdentifier) > 0 {
		input.SessionIdentifier = aws.String(_bedrockagentruntimeSessionIdentifier)
	}
	if len(_bedrockagentruntimeSessionMetadata) > 0 {
		if err := assignInputField(input, "SessionMetadata", _bedrockagentruntimeSessionMetadata); err != nil {
			log.Errorf("invalid --session-metadata: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_bedrockagentruntimeCmd)
	_bedrockagentruntimeCmd.Flags().SortFlags = false

	_bedrockagentruntimeCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_bedrockagentruntimeCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_bedrockagentruntimeCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_bedrockagentruntimeCmd.Flags().StringVarP(&_bedrockagentruntimeActionGroups, "action-groups", "", "", "Action Groups")
	_bedrockagentruntimeCmd.Flags().StringVarP(&_bedrockagentruntimeAgentAliasId, "agent-alias-id", "", "", "Agent Alias ID")
	_bedrockagentruntimeCmd.Flags().StringVarP(&_bedrockagentruntimeAgentCollaboration, "agent-collaboration", "", "", "Agent Collaboration")
	_bedrockagentruntimeCmd.Flags().StringVarP(&_bedrockagentruntimeAgentId, "agent-id", "", "", "Agent ID")
	_bedrockagentruntimeCmd.Flags().StringVarP(&_bedrockagentruntimeAgentName, "agent-name", "", "", "Agent Name")
	_bedrockagentruntimeCmd.Flags().StringVarP(&_bedrockagentruntimeBedrockModelConfigurations, "bedrock-model-configurations", "", "", "Bedrock Model Configurations")
	_bedrockagentruntimeCmd.Flags().StringVarP(&_bedrockagentruntimeCollaboratorConfigurations, "collaborator-configurations", "", "", "Collaborator Configurations")
	_bedrockagentruntimeCmd.Flags().StringVarP(&_bedrockagentruntimeCollaborators, "collaborators", "", "", "Collaborators")
	_bedrockagentruntimeCmd.Flags().StringVarP(&_bedrockagentruntimeCustomOrchestration, "custom-orchestration", "", "", "Custom Orchestration")
	_bedrockagentruntimeCmd.Flags().StringVarP(&_bedrockagentruntimeCustomerEncryptionKeyArn, "customer-encryption-key-arn", "", "", "Customer Encryption Key ARN")
	_bedrockagentruntimeCmd.Flags().StringVarP(&_bedrockagentruntimeDescription, "description", "", "", "Description")
	_bedrockagentruntimeCmd.Flags().StringVarP(&_bedrockagentruntimeEnableTrace, "enable-trace", "", "", "Enable Trace")
	_bedrockagentruntimeCmd.Flags().StringVarP(&_bedrockagentruntimeEncryptionKeyArn, "encryption-key-arn", "", "", "Encryption Key ARN")
	_bedrockagentruntimeCmd.Flags().StringVarP(&_bedrockagentruntimeEventType, "event-type", "", "", "Event Type")
	_bedrockagentruntimeCmd.Flags().StringVarP(&_bedrockagentruntimeExecutionId, "execution-id", "", "", "Execution ID")
	_bedrockagentruntimeCmd.Flags().StringVarP(&_bedrockagentruntimeExecutionIdentifier, "execution-identifier", "", "", "Execution Identifier")
	_bedrockagentruntimeCmd.Flags().StringVarP(&_bedrockagentruntimeFlowAliasIdentifier, "flow-alias-identifier", "", "", "Flow Alias Identifier")
	_bedrockagentruntimeCmd.Flags().StringVarP(&_bedrockagentruntimeFlowExecutionName, "flow-execution-name", "", "", "Flow Execution Name")
	_bedrockagentruntimeCmd.Flags().StringVarP(&_bedrockagentruntimeFlowIdentifier, "flow-identifier", "", "", "Flow Identifier")
	_bedrockagentruntimeCmd.Flags().StringVarP(&_bedrockagentruntimeFoundationModel, "foundation-model", "", "", "Foundation Model")
	_bedrockagentruntimeCmd.Flags().StringVarP(&_bedrockagentruntimeGuardrailConfiguration, "guardrail-configuration", "", "", "Guardrail Configuration")
	_bedrockagentruntimeCmd.Flags().StringVarP(&_bedrockagentruntimeIdleSessionTTLInSeconds, "idle-session-ttlin-seconds", "", "", "Idle Session Ttlin Seconds")
	_bedrockagentruntimeCmd.Flags().StringVarP(&_bedrockagentruntimeInlineSessionState, "inline-session-state", "", "", "Inline Session State")
	_bedrockagentruntimeCmd.Flags().StringVarP(&_bedrockagentruntimeInput, "input", "", "", "Input")
	_bedrockagentruntimeCmd.Flags().StringVarP(&_bedrockagentruntimeInputText, "input-text", "", "", "Input Text")
	_bedrockagentruntimeCmd.Flags().StringVarP(&_bedrockagentruntimeInputs, "inputs", "", "", "Inputs")
	_bedrockagentruntimeCmd.Flags().StringVarP(&_bedrockagentruntimeInstruction, "instruction", "", "", "Instruction")
	_bedrockagentruntimeCmd.Flags().StringVarP(&_bedrockagentruntimeInvocationId, "invocation-id", "", "", "Invocation ID")
	_bedrockagentruntimeCmd.Flags().StringVarP(&_bedrockagentruntimeInvocationIdentifier, "invocation-identifier", "", "", "Invocation Identifier")
	_bedrockagentruntimeCmd.Flags().StringVarP(&_bedrockagentruntimeInvocationStepId, "invocation-step-id", "", "", "Invocation Step ID")
	_bedrockagentruntimeCmd.Flags().StringVarP(&_bedrockagentruntimeInvocationStepTime, "invocation-step-time", "", "", "Invocation Step Time")
	_bedrockagentruntimeCmd.Flags().StringVarP(&_bedrockagentruntimeKnowledgeBaseId, "knowledge-base-id", "", "", "Knowledge Base ID")
	_bedrockagentruntimeCmd.Flags().StringVarP(&_bedrockagentruntimeKnowledgeBases, "knowledge-bases", "", "", "Knowledge Bases")
	_bedrockagentruntimeCmd.Flags().StringVarP(&_bedrockagentruntimeMaxItems, "max-items", "", "", "Max Items")
	_bedrockagentruntimeCmd.Flags().StringVarP(&_bedrockagentruntimeMaxResults, "max-results", "", "", "Max Results")
	_bedrockagentruntimeCmd.Flags().StringVarP(&_bedrockagentruntimeMemoryId, "memory-id", "", "", "Memory ID")
	_bedrockagentruntimeCmd.Flags().StringVarP(&_bedrockagentruntimeMemoryType, "memory-type", "", "", "Memory Type")
	_bedrockagentruntimeCmd.Flags().StringVarP(&_bedrockagentruntimeModelPerformanceConfiguration, "model-performance-configuration", "", "", "Model Performance Configuration")
	_bedrockagentruntimeCmd.Flags().StringVarP(&_bedrockagentruntimeNextToken, "next-token", "", "", "Next Token")
	_bedrockagentruntimeCmd.Flags().StringVarP(&_bedrockagentruntimeOrchestrationType, "orchestration-type", "", "", "Orchestration Type")
	_bedrockagentruntimeCmd.Flags().StringVarP(&_bedrockagentruntimePayload, "payload", "", "", "Payload")
	_bedrockagentruntimeCmd.Flags().StringVarP(&_bedrockagentruntimePromptCreationConfigurations, "prompt-creation-configurations", "", "", "Prompt Creation Configurations")
	_bedrockagentruntimeCmd.Flags().StringVarP(&_bedrockagentruntimePromptOverrideConfiguration, "prompt-override-configuration", "", "", "Prompt Override Configuration")
	_bedrockagentruntimeCmd.Flags().StringVarP(&_bedrockagentruntimeQueries, "queries", "", "", "Queries")
	_bedrockagentruntimeCmd.Flags().StringVarP(&_bedrockagentruntimeQueryGenerationInput, "query-generation-input", "", "", "Query Generation Input")
	_bedrockagentruntimeCmd.Flags().StringVarP(&_bedrockagentruntimeRerankingConfiguration, "reranking-configuration", "", "", "Reranking Configuration")
	_bedrockagentruntimeCmd.Flags().StringVarP(&_bedrockagentruntimeResourceArn, "resource-arn", "", "", "Resource ARN")
	_bedrockagentruntimeCmd.Flags().StringVarP(&_bedrockagentruntimeRetrievalConfiguration, "retrieval-configuration", "", "", "Retrieval Configuration")
	_bedrockagentruntimeCmd.Flags().StringVarP(&_bedrockagentruntimeRetrievalQuery, "retrieval-query", "", "", "Retrieval Query")
	_bedrockagentruntimeCmd.Flags().StringVarP(&_bedrockagentruntimeRetrieveAndGenerateConfiguration, "retrieve-and-generate-configuration", "", "", "Retrieve And Generate Configuration")
	_bedrockagentruntimeCmd.Flags().StringVarP(&_bedrockagentruntimeSessionConfiguration, "session-configuration", "", "", "Session Configuration")
	_bedrockagentruntimeCmd.Flags().StringVarP(&_bedrockagentruntimeSessionId, "session-id", "", "", "Session ID")
	_bedrockagentruntimeCmd.Flags().StringVarP(&_bedrockagentruntimeSessionIdentifier, "session-identifier", "", "", "Session Identifier")
	_bedrockagentruntimeCmd.Flags().StringVarP(&_bedrockagentruntimeSessionMetadata, "session-metadata", "", "", "Session Metadata")
	_bedrockagentruntimeCmd.Flags().StringVarP(&_bedrockagentruntimeSessionState, "session-state", "", "", "Session State")
	_bedrockagentruntimeCmd.Flags().StringVarP(&_bedrockagentruntimeSourceArn, "source-arn", "", "", "Source ARN")
	_bedrockagentruntimeCmd.Flags().StringVarP(&_bedrockagentruntimeSources, "sources", "", "", "Sources")
	_bedrockagentruntimeCmd.Flags().StringVarP(&_bedrockagentruntimeStreamingConfigurations, "streaming-configurations", "", "", "Streaming Configurations")
	_bedrockagentruntimeCmd.Flags().StringSliceVarP(&_bedrockagentruntimeTagKeys, "tag-keys", "", nil, "Tag Keys")
	_bedrockagentruntimeCmd.Flags().StringVarP(&_bedrockagentruntimeTags, "tags", "", "", "Tags")
	_bedrockagentruntimeCmd.Flags().StringVarP(&_bedrockagentruntimeTargetModelId, "target-model-id", "", "", "Target Model ID")
	_bedrockagentruntimeCmd.Flags().StringVarP(&_bedrockagentruntimeTransformationConfiguration, "transformation-configuration", "", "", "Transformation Configuration")

	_bedrockagentruntimeCmd.Flags().BoolVarP(&_bedrockagentruntimeCreateInvocation, "create-invocation", "", false, "Create Invocation")
	_bedrockagentruntimeCmd.Flags().BoolVarP(&_bedrockagentruntimeCreateSession, "create-session", "", false, "Create Session")
	_bedrockagentruntimeCmd.Flags().BoolVarP(&_bedrockagentruntimeDeleteAgentMemory, "delete-agent-memory", "", false, "Delete Agent Memory")
	_bedrockagentruntimeCmd.Flags().BoolVarP(&_bedrockagentruntimeDeleteSession, "delete-session", "", false, "Delete Session")
	_bedrockagentruntimeCmd.Flags().BoolVarP(&_bedrockagentruntimeEndSession, "end-session", "", false, "End Session")
	_bedrockagentruntimeCmd.Flags().BoolVarP(&_bedrockagentruntimeGenerateQuery, "generate-query", "", false, "Generate Query")
	_bedrockagentruntimeCmd.Flags().BoolVarP(&_bedrockagentruntimeGetAgentMemory, "get-agent-memory", "", false, "Get Agent Memory")
	_bedrockagentruntimeCmd.Flags().BoolVarP(&_bedrockagentruntimeGetExecutionFlowSnapshot, "get-execution-flow-snapshot", "", false, "Get Execution Flow Snapshot")
	_bedrockagentruntimeCmd.Flags().BoolVarP(&_bedrockagentruntimeGetFlowExecution, "get-flow-execution", "", false, "Get Flow Execution")
	_bedrockagentruntimeCmd.Flags().BoolVarP(&_bedrockagentruntimeGetInvocationStep, "get-invocation-step", "", false, "Get Invocation Step")
	_bedrockagentruntimeCmd.Flags().BoolVarP(&_bedrockagentruntimeGetSession, "get-session", "", false, "Get Session")
	_bedrockagentruntimeCmd.Flags().BoolVarP(&_bedrockagentruntimeInvokeAgent, "invoke-agent", "", false, "Invoke Agent")
	_bedrockagentruntimeCmd.Flags().BoolVarP(&_bedrockagentruntimeInvokeFlow, "invoke-flow", "", false, "Invoke Flow")
	_bedrockagentruntimeCmd.Flags().BoolVarP(&_bedrockagentruntimeInvokeInlineAgent, "invoke-inline-agent", "", false, "Invoke Inline Agent")
	_bedrockagentruntimeCmd.Flags().BoolVarP(&_bedrockagentruntimeListFlowExecutionEvents, "list-flow-execution-events", "", false, "List Flow Execution Events")
	_bedrockagentruntimeCmd.Flags().BoolVarP(&_bedrockagentruntimeListFlowExecutions, "list-flow-executions", "", false, "List Flow Executions")
	_bedrockagentruntimeCmd.Flags().BoolVarP(&_bedrockagentruntimeListInvocationSteps, "list-invocation-steps", "", false, "List Invocation Steps")
	_bedrockagentruntimeCmd.Flags().BoolVarP(&_bedrockagentruntimeListInvocations, "list-invocations", "", false, "List Invocations")
	_bedrockagentruntimeCmd.Flags().BoolVarP(&_bedrockagentruntimeListSessions, "list-sessions", "", false, "List Sessions")
	_bedrockagentruntimeCmd.Flags().BoolVarP(&_bedrockagentruntimeListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_bedrockagentruntimeCmd.Flags().BoolVarP(&_bedrockagentruntimeOptimizePrompt, "optimize-prompt", "", false, "Optimize Prompt")
	_bedrockagentruntimeCmd.Flags().BoolVarP(&_bedrockagentruntimePutInvocationStep, "put-invocation-step", "", false, "Put Invocation Step")
	_bedrockagentruntimeCmd.Flags().BoolVarP(&_bedrockagentruntimeRerank, "rerank", "", false, "Rerank")
	_bedrockagentruntimeCmd.Flags().BoolVarP(&_bedrockagentruntimeRetrieve, "retrieve", "", false, "Retrieve")
	_bedrockagentruntimeCmd.Flags().BoolVarP(&_bedrockagentruntimeRetrieveAndGenerate, "retrieve-and-generate", "", false, "Retrieve And Generate")
	_bedrockagentruntimeCmd.Flags().BoolVarP(&_bedrockagentruntimeRetrieveAndGenerateStream, "retrieve-and-generate-stream", "", false, "Retrieve And Generate Stream")
	_bedrockagentruntimeCmd.Flags().BoolVarP(&_bedrockagentruntimeStartFlowExecution, "start-flow-execution", "", false, "Start Flow Execution")
	_bedrockagentruntimeCmd.Flags().BoolVarP(&_bedrockagentruntimeStopFlowExecution, "stop-flow-execution", "", false, "Stop Flow Execution")
	_bedrockagentruntimeCmd.Flags().BoolVarP(&_bedrockagentruntimeTagResource, "tag-resource", "", false, "Tag Resource")
	_bedrockagentruntimeCmd.Flags().BoolVarP(&_bedrockagentruntimeUntagResource, "untag-resource", "", false, "Untag Resource")
	_bedrockagentruntimeCmd.Flags().BoolVarP(&_bedrockagentruntimeUpdateSession, "update-session", "", false, "Update Session")

}
