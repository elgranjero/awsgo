package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/bedrockagentcore"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// bedrockagentcoreCmd represents the bedrockagentcore command
var _bedrockagentcoreCmd = &cobra.Command{
	Use:   "bedrockagentcore",
	Short: "AWS bedrockagentcore CLI",
	Run: func(cmd *cobra.Command, args []string) {
		_awsOutput = resolveAWSOutput(_awsProfile, cmd.Flags().Changed("output"))
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := bedrockagentcore.NewFromConfig(cfg)
		if _bedrockagentcoreBatchCreateMemoryRecords {
			bedrockagentcore_BatchCreateMemoryRecords(cfg, client)
			return
		}
		if _bedrockagentcoreBatchDeleteMemoryRecords {
			bedrockagentcore_BatchDeleteMemoryRecords(cfg, client)
			return
		}
		if _bedrockagentcoreBatchUpdateMemoryRecords {
			bedrockagentcore_BatchUpdateMemoryRecords(cfg, client)
			return
		}
		if _bedrockagentcoreCompleteResourceTokenAuth {
			bedrockagentcore_CompleteResourceTokenAuth(cfg, client)
			return
		}
		if _bedrockagentcoreCreateEvent {
			bedrockagentcore_CreateEvent(cfg, client)
			return
		}
		if _bedrockagentcoreDeleteEvent {
			bedrockagentcore_DeleteEvent(cfg, client)
			return
		}
		if _bedrockagentcoreDeleteMemoryRecord {
			bedrockagentcore_DeleteMemoryRecord(cfg, client)
			return
		}
		if _bedrockagentcoreEvaluate {
			bedrockagentcore_Evaluate(cfg, client)
			return
		}
		if _bedrockagentcoreGetAgentCard {
			bedrockagentcore_GetAgentCard(cfg, client)
			return
		}
		if _bedrockagentcoreGetBrowserSession {
			bedrockagentcore_GetBrowserSession(cfg, client)
			return
		}
		if _bedrockagentcoreGetCodeInterpreterSession {
			bedrockagentcore_GetCodeInterpreterSession(cfg, client)
			return
		}
		if _bedrockagentcoreGetEvent {
			bedrockagentcore_GetEvent(cfg, client)
			return
		}
		if _bedrockagentcoreGetMemoryRecord {
			bedrockagentcore_GetMemoryRecord(cfg, client)
			return
		}
		if _bedrockagentcoreGetResourceApiKey {
			bedrockagentcore_GetResourceApiKey(cfg, client)
			return
		}
		if _bedrockagentcoreGetResourceOauth2Token {
			bedrockagentcore_GetResourceOauth2Token(cfg, client)
			return
		}
		if _bedrockagentcoreGetWorkloadAccessToken {
			bedrockagentcore_GetWorkloadAccessToken(cfg, client)
			return
		}
		if _bedrockagentcoreGetWorkloadAccessTokenForJWT {
			bedrockagentcore_GetWorkloadAccessTokenForJWT(cfg, client)
			return
		}
		if _bedrockagentcoreGetWorkloadAccessTokenForUserId {
			bedrockagentcore_GetWorkloadAccessTokenForUserId(cfg, client)
			return
		}
		if _bedrockagentcoreInvokeAgentRuntime {
			bedrockagentcore_InvokeAgentRuntime(cfg, client)
			return
		}
		if _bedrockagentcoreInvokeCodeInterpreter {
			bedrockagentcore_InvokeCodeInterpreter(cfg, client)
			return
		}
		if _bedrockagentcoreListActors {
			bedrockagentcore_ListActors(cfg, client)
			return
		}
		if _bedrockagentcoreListBrowserSessions {
			bedrockagentcore_ListBrowserSessions(cfg, client)
			return
		}
		if _bedrockagentcoreListCodeInterpreterSessions {
			bedrockagentcore_ListCodeInterpreterSessions(cfg, client)
			return
		}
		if _bedrockagentcoreListEvents {
			bedrockagentcore_ListEvents(cfg, client)
			return
		}
		if _bedrockagentcoreListMemoryExtractionJobs {
			bedrockagentcore_ListMemoryExtractionJobs(cfg, client)
			return
		}
		if _bedrockagentcoreListMemoryRecords {
			bedrockagentcore_ListMemoryRecords(cfg, client)
			return
		}
		if _bedrockagentcoreListSessions {
			bedrockagentcore_ListSessions(cfg, client)
			return
		}
		if _bedrockagentcoreRetrieveMemoryRecords {
			bedrockagentcore_RetrieveMemoryRecords(cfg, client)
			return
		}
		if _bedrockagentcoreSaveBrowserSessionProfile {
			bedrockagentcore_SaveBrowserSessionProfile(cfg, client)
			return
		}
		if _bedrockagentcoreStartBrowserSession {
			bedrockagentcore_StartBrowserSession(cfg, client)
			return
		}
		if _bedrockagentcoreStartCodeInterpreterSession {
			bedrockagentcore_StartCodeInterpreterSession(cfg, client)
			return
		}
		if _bedrockagentcoreStartMemoryExtractionJob {
			bedrockagentcore_StartMemoryExtractionJob(cfg, client)
			return
		}
		if _bedrockagentcoreStopBrowserSession {
			bedrockagentcore_StopBrowserSession(cfg, client)
			return
		}
		if _bedrockagentcoreStopCodeInterpreterSession {
			bedrockagentcore_StopCodeInterpreterSession(cfg, client)
			return
		}
		if _bedrockagentcoreStopRuntimeSession {
			bedrockagentcore_StopRuntimeSession(cfg, client)
			return
		}
		if _bedrockagentcoreUpdateBrowserStream {
			bedrockagentcore_UpdateBrowserStream(cfg, client)
			return
		}

	},
}

var (
	_bedrockagentcoreBatchCreateMemoryRecords        bool
	_bedrockagentcoreBatchDeleteMemoryRecords        bool
	_bedrockagentcoreBatchUpdateMemoryRecords        bool
	_bedrockagentcoreCompleteResourceTokenAuth       bool
	_bedrockagentcoreCreateEvent                     bool
	_bedrockagentcoreDeleteEvent                     bool
	_bedrockagentcoreDeleteMemoryRecord              bool
	_bedrockagentcoreEvaluate                        bool
	_bedrockagentcoreGetAgentCard                    bool
	_bedrockagentcoreGetBrowserSession               bool
	_bedrockagentcoreGetCodeInterpreterSession       bool
	_bedrockagentcoreGetEvent                        bool
	_bedrockagentcoreGetMemoryRecord                 bool
	_bedrockagentcoreGetResourceApiKey               bool
	_bedrockagentcoreGetResourceOauth2Token          bool
	_bedrockagentcoreGetWorkloadAccessToken          bool
	_bedrockagentcoreGetWorkloadAccessTokenForJWT    bool
	_bedrockagentcoreGetWorkloadAccessTokenForUserId bool
	_bedrockagentcoreInvokeAgentRuntime              bool
	_bedrockagentcoreInvokeCodeInterpreter           bool
	_bedrockagentcoreListActors                      bool
	_bedrockagentcoreListBrowserSessions             bool
	_bedrockagentcoreListCodeInterpreterSessions     bool
	_bedrockagentcoreListEvents                      bool
	_bedrockagentcoreListMemoryExtractionJobs        bool
	_bedrockagentcoreListMemoryRecords               bool
	_bedrockagentcoreListSessions                    bool
	_bedrockagentcoreRetrieveMemoryRecords           bool
	_bedrockagentcoreSaveBrowserSessionProfile       bool
	_bedrockagentcoreStartBrowserSession             bool
	_bedrockagentcoreStartCodeInterpreterSession     bool
	_bedrockagentcoreStartMemoryExtractionJob        bool
	_bedrockagentcoreStopBrowserSession              bool
	_bedrockagentcoreStopCodeInterpreterSession      bool
	_bedrockagentcoreStopRuntimeSession              bool
	_bedrockagentcoreUpdateBrowserStream             bool

	_bedrockagentcoreAccept                         string
	_bedrockagentcoreAccountId                      string
	_bedrockagentcoreActorId                        string
	_bedrockagentcoreAgentRuntimeArn                string
	_bedrockagentcoreArguments                      string
	_bedrockagentcoreBaggage                        string
	_bedrockagentcoreBranch                         string
	_bedrockagentcoreBrowserIdentifier              string
	_bedrockagentcoreClientToken                    string
	_bedrockagentcoreCodeInterpreterIdentifier      string
	_bedrockagentcoreContentType                    string
	_bedrockagentcoreCustomParameters               string
	_bedrockagentcoreCustomState                    string
	_bedrockagentcoreEvaluationInput                string
	_bedrockagentcoreEvaluationTarget               string
	_bedrockagentcoreEvaluatorId                    string
	_bedrockagentcoreEventId                        string
	_bedrockagentcoreEventTimestamp                 string
	_bedrockagentcoreExtensions                     string
	_bedrockagentcoreExtractionJob                  string
	_bedrockagentcoreFilter                         string
	_bedrockagentcoreForceAuthentication            string
	_bedrockagentcoreIncludePayloads                string
	_bedrockagentcoreMaxResults                     string
	_bedrockagentcoreMcpProtocolVersion             string
	_bedrockagentcoreMcpSessionId                   string
	_bedrockagentcoreMemoryId                       string
	_bedrockagentcoreMemoryRecordId                 string
	_bedrockagentcoreMemoryStrategyId               string
	_bedrockagentcoreMetadata                       string
	_bedrockagentcoreName                           string
	_bedrockagentcoreNamespace                      string
	_bedrockagentcoreNextToken                      string
	_bedrockagentcoreOauth2Flow                     string
	_bedrockagentcorePayload                        string
	_bedrockagentcoreProfileConfiguration           string
	_bedrockagentcoreProfileIdentifier              string
	_bedrockagentcoreProxyConfiguration             string
	_bedrockagentcoreQualifier                      string
	_bedrockagentcoreRecords                        string
	_bedrockagentcoreResourceCredentialProviderName string
	_bedrockagentcoreResourceOauth2ReturnUrl        string
	_bedrockagentcoreRuntimeSessionId               string
	_bedrockagentcoreRuntimeUserId                  string
	_bedrockagentcoreScopes                         []string
	_bedrockagentcoreSearchCriteria                 string
	_bedrockagentcoreSessionId                      string
	_bedrockagentcoreSessionTimeoutSeconds          string
	_bedrockagentcoreSessionUri                     string
	_bedrockagentcoreStatus                         string
	_bedrockagentcoreStreamUpdate                   string
	_bedrockagentcoreTraceId                        string
	_bedrockagentcoreTraceParent                    string
	_bedrockagentcoreTraceState                     string
	_bedrockagentcoreUserId                         string
	_bedrockagentcoreUserIdentifier                 string
	_bedrockagentcoreUserToken                      string
	_bedrockagentcoreViewPort                       string
	_bedrockagentcoreWorkloadIdentityToken          string
	_bedrockagentcoreWorkloadName                   string
)

// Creates multiple memory records in a single batch operation for the specified
// memory with custom content.
func bedrockagentcore_BatchCreateMemoryRecords(cfg aws.Config, client *bedrockagentcore.Client) {
	input := &bedrockagentcore.BatchCreateMemoryRecordsInput{
		// MemoryId: *string, // Required
		// Records: []types.MemoryRecordCreateInput, // Required
	}

	if len(_bedrockagentcoreMemoryId) > 0 {
		input.MemoryId = aws.String(_bedrockagentcoreMemoryId)
	}
	if len(_bedrockagentcoreRecords) > 0 {
		if err := assignInputField(input, "Records", _bedrockagentcoreRecords); err != nil {
			log.Errorf("invalid --records: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcoreClientToken) > 0 {
		input.ClientToken = aws.String(_bedrockagentcoreClientToken)
	}

	if resp, err := client.BatchCreateMemoryRecords(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes multiple memory records in a single batch operation from the specified
// memory.
func bedrockagentcore_BatchDeleteMemoryRecords(cfg aws.Config, client *bedrockagentcore.Client) {
	input := &bedrockagentcore.BatchDeleteMemoryRecordsInput{
		// MemoryId: *string, // Required
		// Records: []types.MemoryRecordDeleteInput, // Required
	}

	if len(_bedrockagentcoreMemoryId) > 0 {
		input.MemoryId = aws.String(_bedrockagentcoreMemoryId)
	}
	if len(_bedrockagentcoreRecords) > 0 {
		if err := assignInputField(input, "Records", _bedrockagentcoreRecords); err != nil {
			log.Errorf("invalid --records: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchDeleteMemoryRecords(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates multiple memory records with custom content in a single batch operation
// within the specified memory.
func bedrockagentcore_BatchUpdateMemoryRecords(cfg aws.Config, client *bedrockagentcore.Client) {
	input := &bedrockagentcore.BatchUpdateMemoryRecordsInput{
		// MemoryId: *string, // Required
		// Records: []types.MemoryRecordUpdateInput, // Required
	}

	if len(_bedrockagentcoreMemoryId) > 0 {
		input.MemoryId = aws.String(_bedrockagentcoreMemoryId)
	}
	if len(_bedrockagentcoreRecords) > 0 {
		if err := assignInputField(input, "Records", _bedrockagentcoreRecords); err != nil {
			log.Errorf("invalid --records: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchUpdateMemoryRecords(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Confirms the user authentication session for obtaining OAuth2.0 tokens for a
// resource.
func bedrockagentcore_CompleteResourceTokenAuth(cfg aws.Config, client *bedrockagentcore.Client) {
	input := &bedrockagentcore.CompleteResourceTokenAuthInput{
		// SessionUri: *string, // Required
		// UserIdentifier: types.UserIdentifier, // Required
	}

	if len(_bedrockagentcoreSessionUri) > 0 {
		input.SessionUri = aws.String(_bedrockagentcoreSessionUri)
	}
	if len(_bedrockagentcoreUserIdentifier) > 0 {
		if err := assignInputField(input, "UserIdentifier", _bedrockagentcoreUserIdentifier); err != nil {
			log.Errorf("invalid --user-identifier: %s", err.Error())
			return
		}
	}

	if resp, err := client.CompleteResourceTokenAuth(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an event in an AgentCore Memory resource. Events represent interactions
// or activities that occur within a session and are associated with specific
// actors.
//
// To use this operation, you must have the bedrock-agentcore:CreateEvent
// permission.
//
// This operation is subject to request rate limiting.
func bedrockagentcore_CreateEvent(cfg aws.Config, client *bedrockagentcore.Client) {
	input := &bedrockagentcore.CreateEventInput{
		// ActorId: *string, // Required
		// EventTimestamp: *time.Time, // Required
		// MemoryId: *string, // Required
		// Payload: []types.PayloadType, // Required
	}

	if len(_bedrockagentcoreActorId) > 0 {
		input.ActorId = aws.String(_bedrockagentcoreActorId)
	}
	if len(_bedrockagentcoreEventTimestamp) > 0 {
		if err := assignInputField(input, "EventTimestamp", _bedrockagentcoreEventTimestamp); err != nil {
			log.Errorf("invalid --event-timestamp: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcoreMemoryId) > 0 {
		input.MemoryId = aws.String(_bedrockagentcoreMemoryId)
	}
	if len(_bedrockagentcorePayload) > 0 {
		if err := assignInputField(input, "Payload", _bedrockagentcorePayload); err != nil {
			log.Errorf("invalid --payload: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcoreBranch) > 0 {
		if err := assignInputField(input, "Branch", _bedrockagentcoreBranch); err != nil {
			log.Errorf("invalid --branch: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcoreClientToken) > 0 {
		input.ClientToken = aws.String(_bedrockagentcoreClientToken)
	}
	if len(_bedrockagentcoreMetadata) > 0 {
		if err := assignInputField(input, "Metadata", _bedrockagentcoreMetadata); err != nil {
			log.Errorf("invalid --metadata: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcoreSessionId) > 0 {
		input.SessionId = aws.String(_bedrockagentcoreSessionId)
	}

	if resp, err := client.CreateEvent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an event from an AgentCore Memory resource. When you delete an event,
// it is permanently removed.
//
// To use this operation, you must have the bedrock-agentcore:DeleteEvent
// permission.
func bedrockagentcore_DeleteEvent(cfg aws.Config, client *bedrockagentcore.Client) {
	input := &bedrockagentcore.DeleteEventInput{
		// ActorId: *string, // Required
		// EventId: *string, // Required
		// MemoryId: *string, // Required
		// SessionId: *string, // Required
	}

	if len(_bedrockagentcoreActorId) > 0 {
		input.ActorId = aws.String(_bedrockagentcoreActorId)
	}
	if len(_bedrockagentcoreEventId) > 0 {
		input.EventId = aws.String(_bedrockagentcoreEventId)
	}
	if len(_bedrockagentcoreMemoryId) > 0 {
		input.MemoryId = aws.String(_bedrockagentcoreMemoryId)
	}
	if len(_bedrockagentcoreSessionId) > 0 {
		input.SessionId = aws.String(_bedrockagentcoreSessionId)
	}

	if resp, err := client.DeleteEvent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a memory record from an AgentCore Memory resource. When you delete a
// memory record, it is permanently removed.
//
// To use this operation, you must have the bedrock-agentcore:DeleteMemoryRecord
// permission.
func bedrockagentcore_DeleteMemoryRecord(cfg aws.Config, client *bedrockagentcore.Client) {
	input := &bedrockagentcore.DeleteMemoryRecordInput{
		// MemoryId: *string, // Required
		// MemoryRecordId: *string, // Required
	}

	if len(_bedrockagentcoreMemoryId) > 0 {
		input.MemoryId = aws.String(_bedrockagentcoreMemoryId)
	}
	if len(_bedrockagentcoreMemoryRecordId) > 0 {
		input.MemoryRecordId = aws.String(_bedrockagentcoreMemoryRecordId)
	}

	if resp, err := client.DeleteMemoryRecord(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Performs on-demand evaluation of agent traces using a specified evaluator.
// This synchronous API accepts traces in OpenTelemetry format and returns
// immediate scoring results with detailed explanations.
func bedrockagentcore_Evaluate(cfg aws.Config, client *bedrockagentcore.Client) {
	input := &bedrockagentcore.EvaluateInput{
		// EvaluationInput: types.EvaluationInput, // Required
		// EvaluatorId: *string, // Required
	}

	if len(_bedrockagentcoreEvaluationInput) > 0 {
		if err := assignInputField(input, "EvaluationInput", _bedrockagentcoreEvaluationInput); err != nil {
			log.Errorf("invalid --evaluation-input: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcoreEvaluatorId) > 0 {
		input.EvaluatorId = aws.String(_bedrockagentcoreEvaluatorId)
	}
	if len(_bedrockagentcoreEvaluationTarget) > 0 {
		if err := assignInputField(input, "EvaluationTarget", _bedrockagentcoreEvaluationTarget); err != nil {
			log.Errorf("invalid --evaluation-target: %s", err.Error())
			return
		}
	}

	if resp, err := client.Evaluate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the A2A agent card associated with an AgentCore Runtime agent.
func bedrockagentcore_GetAgentCard(cfg aws.Config, client *bedrockagentcore.Client) {
	input := &bedrockagentcore.GetAgentCardInput{
		// AgentRuntimeArn: *string, // Required
	}

	if len(_bedrockagentcoreAgentRuntimeArn) > 0 {
		input.AgentRuntimeArn = aws.String(_bedrockagentcoreAgentRuntimeArn)
	}
	if len(_bedrockagentcoreQualifier) > 0 {
		input.Qualifier = aws.String(_bedrockagentcoreQualifier)
	}
	if len(_bedrockagentcoreRuntimeSessionId) > 0 {
		input.RuntimeSessionId = aws.String(_bedrockagentcoreRuntimeSessionId)
	}

	if resp, err := client.GetAgentCard(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves detailed information about a specific browser session in Amazon
// Bedrock AgentCore. This operation returns the session's configuration, current
// status, associated streams, and metadata.
//
// To get a browser session, you must specify both the browser identifier and the
// session ID. The response includes information about the session's viewport
// configuration, timeout settings, and stream endpoints.
//
// The following operations are related to GetBrowserSession :
//
// [StartBrowserSession]
//
// [ListBrowserSessions]
//
// [StopBrowserSession]
//
// [ListBrowserSessions]: https://docs.aws.amazon.com/bedrock-agentcore/latest/APIReference/API_ListBrowserSessions.html
// [StopBrowserSession]: https://docs.aws.amazon.com/bedrock-agentcore/latest/APIReference/API_StopBrowserSession.html
// [StartBrowserSession]: https://docs.aws.amazon.com/bedrock-agentcore/latest/APIReference/API_StartBrowserSession.html
func bedrockagentcore_GetBrowserSession(cfg aws.Config, client *bedrockagentcore.Client) {
	input := &bedrockagentcore.GetBrowserSessionInput{
		// BrowserIdentifier: *string, // Required
		// SessionId: *string, // Required
	}

	if len(_bedrockagentcoreBrowserIdentifier) > 0 {
		input.BrowserIdentifier = aws.String(_bedrockagentcoreBrowserIdentifier)
	}
	if len(_bedrockagentcoreSessionId) > 0 {
		input.SessionId = aws.String(_bedrockagentcoreSessionId)
	}

	if resp, err := client.GetBrowserSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves detailed information about a specific code interpreter session in
// Amazon Bedrock AgentCore. This operation returns the session's configuration,
// current status, and metadata.
//
// To get a code interpreter session, you must specify both the code interpreter
// identifier and the session ID. The response includes information about the
// session's timeout settings and current status.
//
// The following operations are related to GetCodeInterpreterSession :
//
// [StartCodeInterpreterSession]
//
// [ListCodeInterpreterSessions]
//
// [StopCodeInterpreterSession]
//
// [StopCodeInterpreterSession]: https://docs.aws.amazon.com/bedrock-agentcore/latest/APIReference/API_StopCodeInterpreterSession.html
// [StartCodeInterpreterSession]: https://docs.aws.amazon.com/bedrock-agentcore/latest/APIReference/API_StartCodeInterpreterSession.html
// [ListCodeInterpreterSessions]: https://docs.aws.amazon.com/bedrock-agentcore/latest/APIReference/API_ListCodeInterpreterSessions.html
func bedrockagentcore_GetCodeInterpreterSession(cfg aws.Config, client *bedrockagentcore.Client) {
	input := &bedrockagentcore.GetCodeInterpreterSessionInput{
		// CodeInterpreterIdentifier: *string, // Required
		// SessionId: *string, // Required
	}

	if len(_bedrockagentcoreCodeInterpreterIdentifier) > 0 {
		input.CodeInterpreterIdentifier = aws.String(_bedrockagentcoreCodeInterpreterIdentifier)
	}
	if len(_bedrockagentcoreSessionId) > 0 {
		input.SessionId = aws.String(_bedrockagentcoreSessionId)
	}

	if resp, err := client.GetCodeInterpreterSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a specific event in an AgentCore Memory resource.
// To use this operation, you must have the bedrock-agentcore:GetEvent permission.
func bedrockagentcore_GetEvent(cfg aws.Config, client *bedrockagentcore.Client) {
	input := &bedrockagentcore.GetEventInput{
		// ActorId: *string, // Required
		// EventId: *string, // Required
		// MemoryId: *string, // Required
		// SessionId: *string, // Required
	}

	if len(_bedrockagentcoreActorId) > 0 {
		input.ActorId = aws.String(_bedrockagentcoreActorId)
	}
	if len(_bedrockagentcoreEventId) > 0 {
		input.EventId = aws.String(_bedrockagentcoreEventId)
	}
	if len(_bedrockagentcoreMemoryId) > 0 {
		input.MemoryId = aws.String(_bedrockagentcoreMemoryId)
	}
	if len(_bedrockagentcoreSessionId) > 0 {
		input.SessionId = aws.String(_bedrockagentcoreSessionId)
	}

	if resp, err := client.GetEvent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a specific memory record from an AgentCore Memory resource.
// To use this operation, you must have the bedrock-agentcore:GetMemoryRecord
// permission.
func bedrockagentcore_GetMemoryRecord(cfg aws.Config, client *bedrockagentcore.Client) {
	input := &bedrockagentcore.GetMemoryRecordInput{
		// MemoryId: *string, // Required
		// MemoryRecordId: *string, // Required
	}

	if len(_bedrockagentcoreMemoryId) > 0 {
		input.MemoryId = aws.String(_bedrockagentcoreMemoryId)
	}
	if len(_bedrockagentcoreMemoryRecordId) > 0 {
		input.MemoryRecordId = aws.String(_bedrockagentcoreMemoryRecordId)
	}

	if resp, err := client.GetMemoryRecord(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the API key associated with an API key credential provider.
func bedrockagentcore_GetResourceApiKey(cfg aws.Config, client *bedrockagentcore.Client) {
	input := &bedrockagentcore.GetResourceApiKeyInput{
		// ResourceCredentialProviderName: *string, // Required
		// WorkloadIdentityToken: *string, // Required
	}

	if len(_bedrockagentcoreResourceCredentialProviderName) > 0 {
		input.ResourceCredentialProviderName = aws.String(_bedrockagentcoreResourceCredentialProviderName)
	}
	if len(_bedrockagentcoreWorkloadIdentityToken) > 0 {
		input.WorkloadIdentityToken = aws.String(_bedrockagentcoreWorkloadIdentityToken)
	}

	if resp, err := client.GetResourceApiKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the OAuth 2.0 token of the provided resource.
func bedrockagentcore_GetResourceOauth2Token(cfg aws.Config, client *bedrockagentcore.Client) {
	input := &bedrockagentcore.GetResourceOauth2TokenInput{
		// Oauth2Flow: types.Oauth2FlowType, // Required
		// ResourceCredentialProviderName: *string, // Required
		// Scopes: []string, // Required
		// WorkloadIdentityToken: *string, // Required
	}

	if len(_bedrockagentcoreOauth2Flow) > 0 {
		if err := assignInputField(input, "Oauth2Flow", _bedrockagentcoreOauth2Flow); err != nil {
			log.Errorf("invalid --oauth2-flow: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcoreResourceCredentialProviderName) > 0 {
		input.ResourceCredentialProviderName = aws.String(_bedrockagentcoreResourceCredentialProviderName)
	}
	if len(_bedrockagentcoreScopes) > 0 {
		input.Scopes = append([]string(nil), _bedrockagentcoreScopes...)
	}
	if len(_bedrockagentcoreWorkloadIdentityToken) > 0 {
		input.WorkloadIdentityToken = aws.String(_bedrockagentcoreWorkloadIdentityToken)
	}
	if len(_bedrockagentcoreCustomParameters) > 0 {
		if err := assignInputField(input, "CustomParameters", _bedrockagentcoreCustomParameters); err != nil {
			log.Errorf("invalid --custom-parameters: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcoreCustomState) > 0 {
		input.CustomState = aws.String(_bedrockagentcoreCustomState)
	}
	if len(_bedrockagentcoreForceAuthentication) > 0 {
		if err := assignInputField(input, "ForceAuthentication", _bedrockagentcoreForceAuthentication); err != nil {
			log.Errorf("invalid --force-authentication: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcoreResourceOauth2ReturnUrl) > 0 {
		input.ResourceOauth2ReturnUrl = aws.String(_bedrockagentcoreResourceOauth2ReturnUrl)
	}
	if len(_bedrockagentcoreSessionUri) > 0 {
		input.SessionUri = aws.String(_bedrockagentcoreSessionUri)
	}

	if resp, err := client.GetResourceOauth2Token(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Obtains a workload access token for agentic workloads not acting on behalf of a
// user.
func bedrockagentcore_GetWorkloadAccessToken(cfg aws.Config, client *bedrockagentcore.Client) {
	input := &bedrockagentcore.GetWorkloadAccessTokenInput{
		// WorkloadName: *string, // Required
	}

	if len(_bedrockagentcoreWorkloadName) > 0 {
		input.WorkloadName = aws.String(_bedrockagentcoreWorkloadName)
	}

	if resp, err := client.GetWorkloadAccessToken(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Obtains a workload access token for agentic workloads acting on behalf of a
// user, using a JWT token.
func bedrockagentcore_GetWorkloadAccessTokenForJWT(cfg aws.Config, client *bedrockagentcore.Client) {
	input := &bedrockagentcore.GetWorkloadAccessTokenForJWTInput{
		// UserToken: *string, // Required
		// WorkloadName: *string, // Required
	}

	if len(_bedrockagentcoreUserToken) > 0 {
		input.UserToken = aws.String(_bedrockagentcoreUserToken)
	}
	if len(_bedrockagentcoreWorkloadName) > 0 {
		input.WorkloadName = aws.String(_bedrockagentcoreWorkloadName)
	}

	if resp, err := client.GetWorkloadAccessTokenForJWT(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Obtains a workload access token for agentic workloads acting on behalf of a
// user, using the user's ID.
func bedrockagentcore_GetWorkloadAccessTokenForUserId(cfg aws.Config, client *bedrockagentcore.Client) {
	input := &bedrockagentcore.GetWorkloadAccessTokenForUserIdInput{
		// UserId: *string, // Required
		// WorkloadName: *string, // Required
	}

	if len(_bedrockagentcoreUserId) > 0 {
		input.UserId = aws.String(_bedrockagentcoreUserId)
	}
	if len(_bedrockagentcoreWorkloadName) > 0 {
		input.WorkloadName = aws.String(_bedrockagentcoreWorkloadName)
	}

	if resp, err := client.GetWorkloadAccessTokenForUserId(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sends a request to an agent or tool hosted in an Amazon Bedrock AgentCore
// Runtime and receives responses in real-time.
//
// To invoke an agent you must specify the AgentCore Runtime ARN and provide a
// payload containing your request. You can optionally specify a qualifier to
// target a specific version or endpoint of the agent.
//
// This operation supports streaming responses, allowing you to receive partial
// responses as they become available. We recommend using pagination to ensure that
// the operation returns quickly and successfully when processing large responses.
//
// For example code, see [Invoke an AgentCore Runtime agent].
//
// If you're integrating your agent with OAuth, you can't use the Amazon Web
// Services SDK to call InvokeAgentRuntime . Instead, make a HTTPS request to
// InvokeAgentRuntime . For an example, see [Authenticate and authorize with Inbound Auth and Outbound Auth].
//
// To use this operation, you must have the bedrock-agentcore:InvokeAgentRuntime
// permission. If you are making a call to InvokeAgentRuntime on behalf of a user
// ID with the X-Amzn-Bedrock-AgentCore-Runtime-User-Id header, You require
// permissions to both actions ( bedrock-agentcore:InvokeAgentRuntime and
// bedrock-agentcore:InvokeAgentRuntimeForUser ).
//
// [Invoke an AgentCore Runtime agent]: https://docs.aws.amazon.com/bedrock-agentcore/latest/devguide/runtime-invoke-agent.html
// [Authenticate and authorize with Inbound Auth and Outbound Auth]: https://docs.aws.amazon.com/bedrock-agentcore/latest/devguide/runtime-oauth.html
func bedrockagentcore_InvokeAgentRuntime(cfg aws.Config, client *bedrockagentcore.Client) {
	input := &bedrockagentcore.InvokeAgentRuntimeInput{
		// AgentRuntimeArn: *string, // Required
		// Payload: []byte, // Required
	}

	if len(_bedrockagentcoreAgentRuntimeArn) > 0 {
		input.AgentRuntimeArn = aws.String(_bedrockagentcoreAgentRuntimeArn)
	}
	if len(_bedrockagentcorePayload) > 0 {
		if err := assignInputField(input, "Payload", _bedrockagentcorePayload); err != nil {
			log.Errorf("invalid --payload: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcoreAccept) > 0 {
		input.Accept = aws.String(_bedrockagentcoreAccept)
	}
	if len(_bedrockagentcoreAccountId) > 0 {
		input.AccountId = aws.String(_bedrockagentcoreAccountId)
	}
	if len(_bedrockagentcoreBaggage) > 0 {
		input.Baggage = aws.String(_bedrockagentcoreBaggage)
	}
	if len(_bedrockagentcoreContentType) > 0 {
		input.ContentType = aws.String(_bedrockagentcoreContentType)
	}
	if len(_bedrockagentcoreMcpProtocolVersion) > 0 {
		input.McpProtocolVersion = aws.String(_bedrockagentcoreMcpProtocolVersion)
	}
	if len(_bedrockagentcoreMcpSessionId) > 0 {
		input.McpSessionId = aws.String(_bedrockagentcoreMcpSessionId)
	}
	if len(_bedrockagentcoreQualifier) > 0 {
		input.Qualifier = aws.String(_bedrockagentcoreQualifier)
	}
	if len(_bedrockagentcoreRuntimeSessionId) > 0 {
		input.RuntimeSessionId = aws.String(_bedrockagentcoreRuntimeSessionId)
	}
	if len(_bedrockagentcoreRuntimeUserId) > 0 {
		input.RuntimeUserId = aws.String(_bedrockagentcoreRuntimeUserId)
	}
	if len(_bedrockagentcoreTraceId) > 0 {
		input.TraceId = aws.String(_bedrockagentcoreTraceId)
	}
	if len(_bedrockagentcoreTraceParent) > 0 {
		input.TraceParent = aws.String(_bedrockagentcoreTraceParent)
	}
	if len(_bedrockagentcoreTraceState) > 0 {
		input.TraceState = aws.String(_bedrockagentcoreTraceState)
	}

	if resp, err := client.InvokeAgentRuntime(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Executes code within an active code interpreter session in Amazon Bedrock
// AgentCore. This operation processes the provided code, runs it in a secure
// environment, and returns the execution results including output, errors, and
// generated visualizations.
//
// To execute code, you must specify the code interpreter identifier, session ID,
// and the code to run in the arguments parameter. The operation returns a stream
// containing the execution results, which can include text output, error messages,
// and data visualizations.
//
// This operation is subject to request rate limiting based on your account's
// service quotas.
//
// The following operations are related to InvokeCodeInterpreter :
//
// [StartCodeInterpreterSession]
//
// [GetCodeInterpreterSession]
//
// [StartCodeInterpreterSession]: https://docs.aws.amazon.com/bedrock-agentcore/latest/APIReference/API_StartCodeInterpreterSession.html
// [GetCodeInterpreterSession]: https://docs.aws.amazon.com/bedrock-agentcore/latest/APIReference/API_GetCodeInterpreterSession.html
func bedrockagentcore_InvokeCodeInterpreter(cfg aws.Config, client *bedrockagentcore.Client) {
	input := &bedrockagentcore.InvokeCodeInterpreterInput{
		// CodeInterpreterIdentifier: *string, // Required
		// Name: types.ToolName, // Required
	}

	if len(_bedrockagentcoreCodeInterpreterIdentifier) > 0 {
		input.CodeInterpreterIdentifier = aws.String(_bedrockagentcoreCodeInterpreterIdentifier)
	}
	if len(_bedrockagentcoreName) > 0 {
		if err := assignInputField(input, "Name", _bedrockagentcoreName); err != nil {
			log.Errorf("invalid --name: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcoreArguments) > 0 {
		if err := assignInputField(input, "Arguments", _bedrockagentcoreArguments); err != nil {
			log.Errorf("invalid --arguments: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcoreSessionId) > 0 {
		input.SessionId = aws.String(_bedrockagentcoreSessionId)
	}
	if len(_bedrockagentcoreTraceId) > 0 {
		input.TraceId = aws.String(_bedrockagentcoreTraceId)
	}
	if len(_bedrockagentcoreTraceParent) > 0 {
		input.TraceParent = aws.String(_bedrockagentcoreTraceParent)
	}

	if resp, err := client.InvokeCodeInterpreter(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all actors in an AgentCore Memory resource. We recommend using pagination
// to ensure that the operation returns quickly and successfully.
//
// To use this operation, you must have the bedrock-agentcore:ListActors
// permission.
func bedrockagentcore_ListActors(cfg aws.Config, client *bedrockagentcore.Client) {
	input := &bedrockagentcore.ListActorsInput{
		// MemoryId: *string, // Required
	}

	if len(_bedrockagentcoreMemoryId) > 0 {
		input.MemoryId = aws.String(_bedrockagentcoreMemoryId)
	}
	if len(_bedrockagentcoreMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bedrockagentcoreMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcoreNextToken) > 0 {
		input.NextToken = aws.String(_bedrockagentcoreNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListActors(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrockagentcore.ListActorsOutput
	p := bedrockagentcore.NewListActorsPaginator(client, input)
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

// Retrieves a list of browser sessions in Amazon Bedrock AgentCore that match the
// specified criteria. This operation returns summary information about each
// session, including identifiers, status, and timestamps.
//
// You can filter the results by browser identifier and session status. The
// operation supports pagination to handle large result sets efficiently.
//
// We recommend using pagination to ensure that the operation returns quickly and
// successfully when retrieving large numbers of sessions.
//
// The following operations are related to ListBrowserSessions :
//
// [StartBrowserSession]
//
// [GetBrowserSession]
//
// [GetBrowserSession]: https://docs.aws.amazon.com/bedrock-agentcore/latest/APIReference/API_GetBrowserSession.html
// [StartBrowserSession]: https://docs.aws.amazon.com/bedrock-agentcore/latest/APIReference/API_StartBrowserSession.html
func bedrockagentcore_ListBrowserSessions(cfg aws.Config, client *bedrockagentcore.Client) {
	input := &bedrockagentcore.ListBrowserSessionsInput{
		// BrowserIdentifier: *string, // Required
	}

	if len(_bedrockagentcoreBrowserIdentifier) > 0 {
		input.BrowserIdentifier = aws.String(_bedrockagentcoreBrowserIdentifier)
	}
	if len(_bedrockagentcoreMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bedrockagentcoreMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcoreNextToken) > 0 {
		input.NextToken = aws.String(_bedrockagentcoreNextToken)
	}
	if len(_bedrockagentcoreStatus) > 0 {
		if err := assignInputField(input, "Status", _bedrockagentcoreStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListBrowserSessions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list of code interpreter sessions in Amazon Bedrock AgentCore that
// match the specified criteria. This operation returns summary information about
// each session, including identifiers, status, and timestamps.
//
// You can filter the results by code interpreter identifier and session status.
// The operation supports pagination to handle large result sets efficiently.
//
// We recommend using pagination to ensure that the operation returns quickly and
// successfully when retrieving large numbers of sessions.
//
// The following operations are related to ListCodeInterpreterSessions :
//
// [StartCodeInterpreterSession]
//
// [GetCodeInterpreterSession]
//
// [StartCodeInterpreterSession]: https://docs.aws.amazon.com/bedrock-agentcore/latest/APIReference/API_StartCodeInterpreterSession.html
// [GetCodeInterpreterSession]: https://docs.aws.amazon.com/bedrock-agentcore/latest/APIReference/API_GetCodeInterpreterSession.html
func bedrockagentcore_ListCodeInterpreterSessions(cfg aws.Config, client *bedrockagentcore.Client) {
	input := &bedrockagentcore.ListCodeInterpreterSessionsInput{
		// CodeInterpreterIdentifier: *string, // Required
	}

	if len(_bedrockagentcoreCodeInterpreterIdentifier) > 0 {
		input.CodeInterpreterIdentifier = aws.String(_bedrockagentcoreCodeInterpreterIdentifier)
	}
	if len(_bedrockagentcoreMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bedrockagentcoreMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcoreNextToken) > 0 {
		input.NextToken = aws.String(_bedrockagentcoreNextToken)
	}
	if len(_bedrockagentcoreStatus) > 0 {
		if err := assignInputField(input, "Status", _bedrockagentcoreStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListCodeInterpreterSessions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists events in an AgentCore Memory resource based on specified criteria. We
// recommend using pagination to ensure that the operation returns quickly and
// successfully.
//
// To use this operation, you must have the bedrock-agentcore:ListEvents
// permission.
func bedrockagentcore_ListEvents(cfg aws.Config, client *bedrockagentcore.Client) {
	input := &bedrockagentcore.ListEventsInput{
		// ActorId: *string, // Required
		// MemoryId: *string, // Required
		// SessionId: *string, // Required
	}

	if len(_bedrockagentcoreActorId) > 0 {
		input.ActorId = aws.String(_bedrockagentcoreActorId)
	}
	if len(_bedrockagentcoreMemoryId) > 0 {
		input.MemoryId = aws.String(_bedrockagentcoreMemoryId)
	}
	if len(_bedrockagentcoreSessionId) > 0 {
		input.SessionId = aws.String(_bedrockagentcoreSessionId)
	}
	if len(_bedrockagentcoreFilter) > 0 {
		if err := assignInputField(input, "Filter", _bedrockagentcoreFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcoreIncludePayloads) > 0 {
		if err := assignInputField(input, "IncludePayloads", _bedrockagentcoreIncludePayloads); err != nil {
			log.Errorf("invalid --include-payloads: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcoreMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bedrockagentcoreMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcoreNextToken) > 0 {
		input.NextToken = aws.String(_bedrockagentcoreNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListEvents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrockagentcore.ListEventsOutput
	p := bedrockagentcore.NewListEventsPaginator(client, input)
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

// Lists all long-term memory extraction jobs that are eligible to be started with
// optional filtering.
//
// To use this operation, you must have the
// bedrock-agentcore:ListMemoryExtractionJobs permission.
func bedrockagentcore_ListMemoryExtractionJobs(cfg aws.Config, client *bedrockagentcore.Client) {
	input := &bedrockagentcore.ListMemoryExtractionJobsInput{
		// MemoryId: *string, // Required
	}

	if len(_bedrockagentcoreMemoryId) > 0 {
		input.MemoryId = aws.String(_bedrockagentcoreMemoryId)
	}
	if len(_bedrockagentcoreFilter) > 0 {
		if err := assignInputField(input, "Filter", _bedrockagentcoreFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcoreMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bedrockagentcoreMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcoreNextToken) > 0 {
		input.NextToken = aws.String(_bedrockagentcoreNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListMemoryExtractionJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrockagentcore.ListMemoryExtractionJobsOutput
	p := bedrockagentcore.NewListMemoryExtractionJobsPaginator(client, input)
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

// Lists memory records in an AgentCore Memory resource based on specified
// criteria. We recommend using pagination to ensure that the operation returns
// quickly and successfully.
//
// To use this operation, you must have the bedrock-agentcore:ListMemoryRecords
// permission.
func bedrockagentcore_ListMemoryRecords(cfg aws.Config, client *bedrockagentcore.Client) {
	input := &bedrockagentcore.ListMemoryRecordsInput{
		// MemoryId: *string, // Required
		// Namespace: *string, // Required
	}

	if len(_bedrockagentcoreMemoryId) > 0 {
		input.MemoryId = aws.String(_bedrockagentcoreMemoryId)
	}
	if len(_bedrockagentcoreNamespace) > 0 {
		input.Namespace = aws.String(_bedrockagentcoreNamespace)
	}
	if len(_bedrockagentcoreMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bedrockagentcoreMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcoreMemoryStrategyId) > 0 {
		input.MemoryStrategyId = aws.String(_bedrockagentcoreMemoryStrategyId)
	}
	if len(_bedrockagentcoreNextToken) > 0 {
		input.NextToken = aws.String(_bedrockagentcoreNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListMemoryRecords(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrockagentcore.ListMemoryRecordsOutput
	p := bedrockagentcore.NewListMemoryRecordsPaginator(client, input)
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

// Lists sessions in an AgentCore Memory resource based on specified criteria. We
// recommend using pagination to ensure that the operation returns quickly and
// successfully.
//
// To use this operation, you must have the bedrock-agentcore:ListSessions
// permission.
func bedrockagentcore_ListSessions(cfg aws.Config, client *bedrockagentcore.Client) {
	input := &bedrockagentcore.ListSessionsInput{
		// ActorId: *string, // Required
		// MemoryId: *string, // Required
	}

	if len(_bedrockagentcoreActorId) > 0 {
		input.ActorId = aws.String(_bedrockagentcoreActorId)
	}
	if len(_bedrockagentcoreMemoryId) > 0 {
		input.MemoryId = aws.String(_bedrockagentcoreMemoryId)
	}
	if len(_bedrockagentcoreMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bedrockagentcoreMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcoreNextToken) > 0 {
		input.NextToken = aws.String(_bedrockagentcoreNextToken)
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

	var results []*bedrockagentcore.ListSessionsOutput
	p := bedrockagentcore.NewListSessionsPaginator(client, input)
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

// Searches for and retrieves memory records from an AgentCore Memory resource
// based on specified search criteria. We recommend using pagination to ensure that
// the operation returns quickly and successfully.
//
// To use this operation, you must have the bedrock-agentcore:RetrieveMemoryRecords
// permission.
func bedrockagentcore_RetrieveMemoryRecords(cfg aws.Config, client *bedrockagentcore.Client) {
	input := &bedrockagentcore.RetrieveMemoryRecordsInput{
		// MemoryId: *string, // Required
		// Namespace: *string, // Required
		// SearchCriteria: *types.SearchCriteria, // Required
	}

	if len(_bedrockagentcoreMemoryId) > 0 {
		input.MemoryId = aws.String(_bedrockagentcoreMemoryId)
	}
	if len(_bedrockagentcoreNamespace) > 0 {
		input.Namespace = aws.String(_bedrockagentcoreNamespace)
	}
	if len(_bedrockagentcoreSearchCriteria) > 0 {
		if err := assignInputField(input, "SearchCriteria", _bedrockagentcoreSearchCriteria); err != nil {
			log.Errorf("invalid --search-criteria: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcoreMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bedrockagentcoreMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcoreNextToken) > 0 {
		input.NextToken = aws.String(_bedrockagentcoreNextToken)
	}

	if disablePaginator() {
		if resp, err := client.RetrieveMemoryRecords(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrockagentcore.RetrieveMemoryRecordsOutput
	p := bedrockagentcore.NewRetrieveMemoryRecordsPaginator(client, input)
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

// Saves the current state of a browser session as a reusable profile in Amazon
// Bedrock AgentCore. A browser profile captures persistent browser data such as
// cookies and local storage from an active session, enabling you to reuse this
// data in future browser sessions.
//
// To save a browser session profile, you must specify the profile identifier,
// browser identifier, and session ID. The session must be active when saving the
// profile. Once saved, the profile can be used with the StartBrowserSession
// operation to initialize new sessions with the stored browser state.
//
// Browser profiles are useful for scenarios that require persistent
// authentication, maintaining user preferences across sessions, or continuing
// tasks that depend on previously stored browser data.
//
// The following operations are related to SaveBrowserSessionProfile :
//
// [StartBrowserSession]
//
// [GetBrowserSession]
//
// [GetBrowserSession]: https://docs.aws.amazon.com/bedrock-agentcore/latest/APIReference/API_GetBrowserSession.html
// [StartBrowserSession]: https://docs.aws.amazon.com/bedrock-agentcore/latest/APIReference/API_StartBrowserSession.html
func bedrockagentcore_SaveBrowserSessionProfile(cfg aws.Config, client *bedrockagentcore.Client) {
	input := &bedrockagentcore.SaveBrowserSessionProfileInput{
		// BrowserIdentifier: *string, // Required
		// ProfileIdentifier: *string, // Required
		// SessionId: *string, // Required
	}

	if len(_bedrockagentcoreBrowserIdentifier) > 0 {
		input.BrowserIdentifier = aws.String(_bedrockagentcoreBrowserIdentifier)
	}
	if len(_bedrockagentcoreProfileIdentifier) > 0 {
		input.ProfileIdentifier = aws.String(_bedrockagentcoreProfileIdentifier)
	}
	if len(_bedrockagentcoreSessionId) > 0 {
		input.SessionId = aws.String(_bedrockagentcoreSessionId)
	}
	if len(_bedrockagentcoreClientToken) > 0 {
		input.ClientToken = aws.String(_bedrockagentcoreClientToken)
	}
	if len(_bedrockagentcoreTraceId) > 0 {
		input.TraceId = aws.String(_bedrockagentcoreTraceId)
	}
	if len(_bedrockagentcoreTraceParent) > 0 {
		input.TraceParent = aws.String(_bedrockagentcoreTraceParent)
	}

	if resp, err := client.SaveBrowserSessionProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates and initializes a browser session in Amazon Bedrock AgentCore. The
// session enables agents to navigate and interact with web content, extract
// information from websites, and perform web-based tasks as part of their response
// generation.
//
// To create a session, you must specify a browser identifier and a name. You can
// also configure the viewport dimensions to control the visible area of web
// content. The session remains active until it times out or you explicitly stop it
// using the StopBrowserSession operation.
//
// The following operations are related to StartBrowserSession :
//
// [GetBrowserSession]
//
// [UpdateBrowserStream]
//
// [SaveBrowserSessionProfile]
//
// [StopBrowserSession]
//
// [SaveBrowserSessionProfile]: https://docs.aws.amazon.com/bedrock-agentcore/latest/APIReference/API_SaveBrowserSessionProfile.html
// [UpdateBrowserStream]: https://docs.aws.amazon.com/bedrock-agentcore/latest/APIReference/API_UpdateBrowserStream.html
// [GetBrowserSession]: https://docs.aws.amazon.com/bedrock-agentcore/latest/APIReference/API_GetBrowserSession.html
// [StopBrowserSession]: https://docs.aws.amazon.com/bedrock-agentcore/latest/APIReference/API_StopBrowserSession.html
func bedrockagentcore_StartBrowserSession(cfg aws.Config, client *bedrockagentcore.Client) {
	input := &bedrockagentcore.StartBrowserSessionInput{
		// BrowserIdentifier: *string, // Required
	}

	if len(_bedrockagentcoreBrowserIdentifier) > 0 {
		input.BrowserIdentifier = aws.String(_bedrockagentcoreBrowserIdentifier)
	}
	if len(_bedrockagentcoreClientToken) > 0 {
		input.ClientToken = aws.String(_bedrockagentcoreClientToken)
	}
	if len(_bedrockagentcoreExtensions) > 0 {
		if err := assignInputField(input, "Extensions", _bedrockagentcoreExtensions); err != nil {
			log.Errorf("invalid --extensions: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcoreName) > 0 {
		input.Name = aws.String(_bedrockagentcoreName)
	}
	if len(_bedrockagentcoreProfileConfiguration) > 0 {
		if err := assignInputField(input, "ProfileConfiguration", _bedrockagentcoreProfileConfiguration); err != nil {
			log.Errorf("invalid --profile-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcoreProxyConfiguration) > 0 {
		if err := assignInputField(input, "ProxyConfiguration", _bedrockagentcoreProxyConfiguration); err != nil {
			log.Errorf("invalid --proxy-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcoreSessionTimeoutSeconds) > 0 {
		if err := assignInputField(input, "SessionTimeoutSeconds", _bedrockagentcoreSessionTimeoutSeconds); err != nil {
			log.Errorf("invalid --session-timeout-seconds: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcoreTraceId) > 0 {
		input.TraceId = aws.String(_bedrockagentcoreTraceId)
	}
	if len(_bedrockagentcoreTraceParent) > 0 {
		input.TraceParent = aws.String(_bedrockagentcoreTraceParent)
	}
	if len(_bedrockagentcoreViewPort) > 0 {
		if err := assignInputField(input, "ViewPort", _bedrockagentcoreViewPort); err != nil {
			log.Errorf("invalid --view-port: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartBrowserSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates and initializes a code interpreter session in Amazon Bedrock AgentCore.
// The session enables agents to execute code as part of their response generation,
// supporting programming languages such as Python for data analysis,
// visualization, and computation tasks.
//
// To create a session, you must specify a code interpreter identifier and a name.
// The session remains active until it times out or you explicitly stop it using
// the StopCodeInterpreterSession operation.
//
// The following operations are related to StartCodeInterpreterSession :
//
// [InvokeCodeInterpreter]
//
// [GetCodeInterpreterSession]
//
// [StopCodeInterpreterSession]
//
// [InvokeCodeInterpreter]: https://docs.aws.amazon.com/bedrock-agentcore/latest/APIReference/API_InvokeCodeInterpreter.html
// [StopCodeInterpreterSession]: https://docs.aws.amazon.com/bedrock-agentcore/latest/APIReference/API_StopCodeInterpreterSession.html
// [GetCodeInterpreterSession]: https://docs.aws.amazon.com/bedrock-agentcore/latest/APIReference/API_GetCodeInterpreterSession.html
func bedrockagentcore_StartCodeInterpreterSession(cfg aws.Config, client *bedrockagentcore.Client) {
	input := &bedrockagentcore.StartCodeInterpreterSessionInput{
		// CodeInterpreterIdentifier: *string, // Required
	}

	if len(_bedrockagentcoreCodeInterpreterIdentifier) > 0 {
		input.CodeInterpreterIdentifier = aws.String(_bedrockagentcoreCodeInterpreterIdentifier)
	}
	if len(_bedrockagentcoreClientToken) > 0 {
		input.ClientToken = aws.String(_bedrockagentcoreClientToken)
	}
	if len(_bedrockagentcoreName) > 0 {
		input.Name = aws.String(_bedrockagentcoreName)
	}
	if len(_bedrockagentcoreSessionTimeoutSeconds) > 0 {
		if err := assignInputField(input, "SessionTimeoutSeconds", _bedrockagentcoreSessionTimeoutSeconds); err != nil {
			log.Errorf("invalid --session-timeout-seconds: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcoreTraceId) > 0 {
		input.TraceId = aws.String(_bedrockagentcoreTraceId)
	}
	if len(_bedrockagentcoreTraceParent) > 0 {
		input.TraceParent = aws.String(_bedrockagentcoreTraceParent)
	}

	if resp, err := client.StartCodeInterpreterSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a memory extraction job that processes events that failed extraction
// previously in an AgentCore Memory resource and produces structured memory
// records. When earlier extraction attempts have left events unprocessed, this job
// will pick up and extract those as well.
//
// To use this operation, you must have the
// bedrock-agentcore:StartMemoryExtractionJob permission.
func bedrockagentcore_StartMemoryExtractionJob(cfg aws.Config, client *bedrockagentcore.Client) {
	input := &bedrockagentcore.StartMemoryExtractionJobInput{
		// ExtractionJob: *types.ExtractionJob, // Required
		// MemoryId: *string, // Required
	}

	if len(_bedrockagentcoreExtractionJob) > 0 {
		if err := assignInputField(input, "ExtractionJob", _bedrockagentcoreExtractionJob); err != nil {
			log.Errorf("invalid --extraction-job: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcoreMemoryId) > 0 {
		input.MemoryId = aws.String(_bedrockagentcoreMemoryId)
	}
	if len(_bedrockagentcoreClientToken) > 0 {
		input.ClientToken = aws.String(_bedrockagentcoreClientToken)
	}

	if resp, err := client.StartMemoryExtractionJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Terminates an active browser session in Amazon Bedrock AgentCore. This
// operation stops the session, releases associated resources, and makes the
// session unavailable for further use.
//
// To stop a browser session, you must specify both the browser identifier and the
// session ID. Once stopped, a session cannot be restarted; you must create a new
// session using StartBrowserSession .
//
// The following operations are related to StopBrowserSession :
//
// [StartBrowserSession]
//
// [GetBrowserSession]
//
// [GetBrowserSession]: https://docs.aws.amazon.com/bedrock-agentcore/latest/APIReference/API_GetBrowserSession.html
// [StartBrowserSession]: https://docs.aws.amazon.com/bedrock-agentcore/latest/APIReference/API_StartBrowserSession.html
func bedrockagentcore_StopBrowserSession(cfg aws.Config, client *bedrockagentcore.Client) {
	input := &bedrockagentcore.StopBrowserSessionInput{
		// BrowserIdentifier: *string, // Required
		// SessionId: *string, // Required
	}

	if len(_bedrockagentcoreBrowserIdentifier) > 0 {
		input.BrowserIdentifier = aws.String(_bedrockagentcoreBrowserIdentifier)
	}
	if len(_bedrockagentcoreSessionId) > 0 {
		input.SessionId = aws.String(_bedrockagentcoreSessionId)
	}
	if len(_bedrockagentcoreClientToken) > 0 {
		input.ClientToken = aws.String(_bedrockagentcoreClientToken)
	}
	if len(_bedrockagentcoreTraceId) > 0 {
		input.TraceId = aws.String(_bedrockagentcoreTraceId)
	}
	if len(_bedrockagentcoreTraceParent) > 0 {
		input.TraceParent = aws.String(_bedrockagentcoreTraceParent)
	}

	if resp, err := client.StopBrowserSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Terminates an active code interpreter session in Amazon Bedrock AgentCore. This
// operation stops the session, releases associated resources, and makes the
// session unavailable for further use.
//
// To stop a code interpreter session, you must specify both the code interpreter
// identifier and the session ID. Once stopped, a session cannot be restarted; you
// must create a new session using StartCodeInterpreterSession .
//
// The following operations are related to StopCodeInterpreterSession :
//
// [StartCodeInterpreterSession]
//
// [GetCodeInterpreterSession]
//
// [StartCodeInterpreterSession]: https://docs.aws.amazon.com/bedrock-agentcore/latest/APIReference/API_StartCodeInterpreterSession.html
// [GetCodeInterpreterSession]: https://docs.aws.amazon.com/bedrock-agentcore/latest/APIReference/API_GetCodeInterpreterSession.html
func bedrockagentcore_StopCodeInterpreterSession(cfg aws.Config, client *bedrockagentcore.Client) {
	input := &bedrockagentcore.StopCodeInterpreterSessionInput{
		// CodeInterpreterIdentifier: *string, // Required
		// SessionId: *string, // Required
	}

	if len(_bedrockagentcoreCodeInterpreterIdentifier) > 0 {
		input.CodeInterpreterIdentifier = aws.String(_bedrockagentcoreCodeInterpreterIdentifier)
	}
	if len(_bedrockagentcoreSessionId) > 0 {
		input.SessionId = aws.String(_bedrockagentcoreSessionId)
	}
	if len(_bedrockagentcoreClientToken) > 0 {
		input.ClientToken = aws.String(_bedrockagentcoreClientToken)
	}
	if len(_bedrockagentcoreTraceId) > 0 {
		input.TraceId = aws.String(_bedrockagentcoreTraceId)
	}
	if len(_bedrockagentcoreTraceParent) > 0 {
		input.TraceParent = aws.String(_bedrockagentcoreTraceParent)
	}

	if resp, err := client.StopCodeInterpreterSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a session that is running in an running AgentCore Runtime agent.
func bedrockagentcore_StopRuntimeSession(cfg aws.Config, client *bedrockagentcore.Client) {
	input := &bedrockagentcore.StopRuntimeSessionInput{
		// AgentRuntimeArn: *string, // Required
		// RuntimeSessionId: *string, // Required
	}

	if len(_bedrockagentcoreAgentRuntimeArn) > 0 {
		input.AgentRuntimeArn = aws.String(_bedrockagentcoreAgentRuntimeArn)
	}
	if len(_bedrockagentcoreRuntimeSessionId) > 0 {
		input.RuntimeSessionId = aws.String(_bedrockagentcoreRuntimeSessionId)
	}
	if len(_bedrockagentcoreClientToken) > 0 {
		input.ClientToken = aws.String(_bedrockagentcoreClientToken)
	}
	if len(_bedrockagentcoreQualifier) > 0 {
		input.Qualifier = aws.String(_bedrockagentcoreQualifier)
	}

	if resp, err := client.StopRuntimeSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a browser stream. To use this operation, you must have permissions to
// perform the bedrock:UpdateBrowserStream action.
func bedrockagentcore_UpdateBrowserStream(cfg aws.Config, client *bedrockagentcore.Client) {
	input := &bedrockagentcore.UpdateBrowserStreamInput{
		// BrowserIdentifier: *string, // Required
		// SessionId: *string, // Required
		// StreamUpdate: types.StreamUpdate, // Required
	}

	if len(_bedrockagentcoreBrowserIdentifier) > 0 {
		input.BrowserIdentifier = aws.String(_bedrockagentcoreBrowserIdentifier)
	}
	if len(_bedrockagentcoreSessionId) > 0 {
		input.SessionId = aws.String(_bedrockagentcoreSessionId)
	}
	if len(_bedrockagentcoreStreamUpdate) > 0 {
		if err := assignInputField(input, "StreamUpdate", _bedrockagentcoreStreamUpdate); err != nil {
			log.Errorf("invalid --stream-update: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcoreClientToken) > 0 {
		input.ClientToken = aws.String(_bedrockagentcoreClientToken)
	}

	if resp, err := client.UpdateBrowserStream(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_bedrockagentcoreCmd)
	_bedrockagentcoreCmd.Flags().SortFlags = false

	_bedrockagentcoreCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_bedrockagentcoreCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_bedrockagentcoreCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_bedrockagentcoreCmd.Flags().StringVarP(&_bedrockagentcoreAccept, "accept", "", "", "Accept")
	_bedrockagentcoreCmd.Flags().StringVarP(&_bedrockagentcoreAccountId, "account-id", "", "", "Account ID")
	_bedrockagentcoreCmd.Flags().StringVarP(&_bedrockagentcoreActorId, "actor-id", "", "", "Actor ID")
	_bedrockagentcoreCmd.Flags().StringVarP(&_bedrockagentcoreAgentRuntimeArn, "agent-runtime-arn", "", "", "Agent Runtime ARN")
	_bedrockagentcoreCmd.Flags().StringVarP(&_bedrockagentcoreArguments, "arguments", "", "", "Arguments")
	_bedrockagentcoreCmd.Flags().StringVarP(&_bedrockagentcoreBaggage, "baggage", "", "", "Baggage")
	_bedrockagentcoreCmd.Flags().StringVarP(&_bedrockagentcoreBranch, "branch", "", "", "Branch")
	_bedrockagentcoreCmd.Flags().StringVarP(&_bedrockagentcoreBrowserIdentifier, "browser-identifier", "", "", "Browser Identifier")
	_bedrockagentcoreCmd.Flags().StringVarP(&_bedrockagentcoreClientToken, "client-token", "", "", "Client Token")
	_bedrockagentcoreCmd.Flags().StringVarP(&_bedrockagentcoreCodeInterpreterIdentifier, "code-interpreter-identifier", "", "", "Code Interpreter Identifier")
	_bedrockagentcoreCmd.Flags().StringVarP(&_bedrockagentcoreContentType, "content-type", "", "", "Content Type")
	_bedrockagentcoreCmd.Flags().StringVarP(&_bedrockagentcoreCustomParameters, "custom-parameters", "", "", "Custom Parameters")
	_bedrockagentcoreCmd.Flags().StringVarP(&_bedrockagentcoreCustomState, "custom-state", "", "", "Custom State")
	_bedrockagentcoreCmd.Flags().StringVarP(&_bedrockagentcoreEvaluationInput, "evaluation-input", "", "", "Evaluation Input")
	_bedrockagentcoreCmd.Flags().StringVarP(&_bedrockagentcoreEvaluationTarget, "evaluation-target", "", "", "Evaluation Target")
	_bedrockagentcoreCmd.Flags().StringVarP(&_bedrockagentcoreEvaluatorId, "evaluator-id", "", "", "Evaluator ID")
	_bedrockagentcoreCmd.Flags().StringVarP(&_bedrockagentcoreEventId, "event-id", "", "", "Event ID")
	_bedrockagentcoreCmd.Flags().StringVarP(&_bedrockagentcoreEventTimestamp, "event-timestamp", "", "", "Event Timestamp")
	_bedrockagentcoreCmd.Flags().StringVarP(&_bedrockagentcoreExtensions, "extensions", "", "", "Extensions")
	_bedrockagentcoreCmd.Flags().StringVarP(&_bedrockagentcoreExtractionJob, "extraction-job", "", "", "Extraction Job")
	_bedrockagentcoreCmd.Flags().StringVarP(&_bedrockagentcoreFilter, "filter", "", "", "Filter")
	_bedrockagentcoreCmd.Flags().StringVarP(&_bedrockagentcoreForceAuthentication, "force-authentication", "", "", "Force Authentication")
	_bedrockagentcoreCmd.Flags().StringVarP(&_bedrockagentcoreIncludePayloads, "include-payloads", "", "", "Include Payloads")
	_bedrockagentcoreCmd.Flags().StringVarP(&_bedrockagentcoreMaxResults, "max-results", "", "", "Max Results")
	_bedrockagentcoreCmd.Flags().StringVarP(&_bedrockagentcoreMcpProtocolVersion, "mcp-protocol-version", "", "", "Mcp Protocol Version")
	_bedrockagentcoreCmd.Flags().StringVarP(&_bedrockagentcoreMcpSessionId, "mcp-session-id", "", "", "Mcp Session ID")
	_bedrockagentcoreCmd.Flags().StringVarP(&_bedrockagentcoreMemoryId, "memory-id", "", "", "Memory ID")
	_bedrockagentcoreCmd.Flags().StringVarP(&_bedrockagentcoreMemoryRecordId, "memory-record-id", "", "", "Memory Record ID")
	_bedrockagentcoreCmd.Flags().StringVarP(&_bedrockagentcoreMemoryStrategyId, "memory-strategy-id", "", "", "Memory Strategy ID")
	_bedrockagentcoreCmd.Flags().StringVarP(&_bedrockagentcoreMetadata, "metadata", "", "", "Metadata")
	_bedrockagentcoreCmd.Flags().StringVarP(&_bedrockagentcoreName, "name", "", "", "Name")
	_bedrockagentcoreCmd.Flags().StringVarP(&_bedrockagentcoreNamespace, "namespace", "", "", "Namespace")
	_bedrockagentcoreCmd.Flags().StringVarP(&_bedrockagentcoreNextToken, "next-token", "", "", "Next Token")
	_bedrockagentcoreCmd.Flags().StringVarP(&_bedrockagentcoreOauth2Flow, "oauth2-flow", "", "", "Oauth2 Flow")
	_bedrockagentcoreCmd.Flags().StringVarP(&_bedrockagentcorePayload, "payload", "", "", "Payload")
	_bedrockagentcoreCmd.Flags().StringVarP(&_bedrockagentcoreProfileConfiguration, "profile-configuration", "", "", "Profile Configuration")
	_bedrockagentcoreCmd.Flags().StringVarP(&_bedrockagentcoreProfileIdentifier, "profile-identifier", "", "", "Profile Identifier")
	_bedrockagentcoreCmd.Flags().StringVarP(&_bedrockagentcoreProxyConfiguration, "proxy-configuration", "", "", "Proxy Configuration")
	_bedrockagentcoreCmd.Flags().StringVarP(&_bedrockagentcoreQualifier, "qualifier", "", "", "Qualifier")
	_bedrockagentcoreCmd.Flags().StringVarP(&_bedrockagentcoreRecords, "records", "", "", "Records")
	_bedrockagentcoreCmd.Flags().StringVarP(&_bedrockagentcoreResourceCredentialProviderName, "resource-credential-provider-name", "", "", "Resource Credential Provider Name")
	_bedrockagentcoreCmd.Flags().StringVarP(&_bedrockagentcoreResourceOauth2ReturnUrl, "resource-oauth2-return-url", "", "", "Resource Oauth2 Return URL")
	_bedrockagentcoreCmd.Flags().StringVarP(&_bedrockagentcoreRuntimeSessionId, "runtime-session-id", "", "", "Runtime Session ID")
	_bedrockagentcoreCmd.Flags().StringVarP(&_bedrockagentcoreRuntimeUserId, "runtime-user-id", "", "", "Runtime User ID")
	_bedrockagentcoreCmd.Flags().StringSliceVarP(&_bedrockagentcoreScopes, "scopes", "", nil, "Scopes")
	_bedrockagentcoreCmd.Flags().StringVarP(&_bedrockagentcoreSearchCriteria, "search-criteria", "", "", "Search Criteria")
	_bedrockagentcoreCmd.Flags().StringVarP(&_bedrockagentcoreSessionId, "session-id", "", "", "Session ID")
	_bedrockagentcoreCmd.Flags().StringVarP(&_bedrockagentcoreSessionTimeoutSeconds, "session-timeout-seconds", "", "", "Session Timeout Seconds")
	_bedrockagentcoreCmd.Flags().StringVarP(&_bedrockagentcoreSessionUri, "session-uri", "", "", "Session URI")
	_bedrockagentcoreCmd.Flags().StringVarP(&_bedrockagentcoreStatus, "status", "", "", "Status")
	_bedrockagentcoreCmd.Flags().StringVarP(&_bedrockagentcoreStreamUpdate, "stream-update", "", "", "Stream Update")
	_bedrockagentcoreCmd.Flags().StringVarP(&_bedrockagentcoreTraceId, "trace-id", "", "", "Trace ID")
	_bedrockagentcoreCmd.Flags().StringVarP(&_bedrockagentcoreTraceParent, "trace-parent", "", "", "Trace Parent")
	_bedrockagentcoreCmd.Flags().StringVarP(&_bedrockagentcoreTraceState, "trace-state", "", "", "Trace State")
	_bedrockagentcoreCmd.Flags().StringVarP(&_bedrockagentcoreUserId, "user-id", "", "", "User ID")
	_bedrockagentcoreCmd.Flags().StringVarP(&_bedrockagentcoreUserIdentifier, "user-identifier", "", "", "User Identifier")
	_bedrockagentcoreCmd.Flags().StringVarP(&_bedrockagentcoreUserToken, "user-token", "", "", "User Token")
	_bedrockagentcoreCmd.Flags().StringVarP(&_bedrockagentcoreViewPort, "view-port", "", "", "View Port")
	_bedrockagentcoreCmd.Flags().StringVarP(&_bedrockagentcoreWorkloadIdentityToken, "workload-identity-token", "", "", "Workload Identity Token")
	_bedrockagentcoreCmd.Flags().StringVarP(&_bedrockagentcoreWorkloadName, "workload-name", "", "", "Workload Name")

	_bedrockagentcoreCmd.Flags().BoolVarP(&_bedrockagentcoreBatchCreateMemoryRecords, "batch-create-memory-records", "", false, "Batch Create Memory Records")
	_bedrockagentcoreCmd.Flags().BoolVarP(&_bedrockagentcoreBatchDeleteMemoryRecords, "batch-delete-memory-records", "", false, "Batch Delete Memory Records")
	_bedrockagentcoreCmd.Flags().BoolVarP(&_bedrockagentcoreBatchUpdateMemoryRecords, "batch-update-memory-records", "", false, "Batch Update Memory Records")
	_bedrockagentcoreCmd.Flags().BoolVarP(&_bedrockagentcoreCompleteResourceTokenAuth, "complete-resource-token-auth", "", false, "Complete Resource Token Auth")
	_bedrockagentcoreCmd.Flags().BoolVarP(&_bedrockagentcoreCreateEvent, "create-event", "", false, "Create Event")
	_bedrockagentcoreCmd.Flags().BoolVarP(&_bedrockagentcoreDeleteEvent, "delete-event", "", false, "Delete Event")
	_bedrockagentcoreCmd.Flags().BoolVarP(&_bedrockagentcoreDeleteMemoryRecord, "delete-memory-record", "", false, "Delete Memory Record")
	_bedrockagentcoreCmd.Flags().BoolVarP(&_bedrockagentcoreEvaluate, "evaluate", "", false, "Evaluate")
	_bedrockagentcoreCmd.Flags().BoolVarP(&_bedrockagentcoreGetAgentCard, "get-agent-card", "", false, "Get Agent Card")
	_bedrockagentcoreCmd.Flags().BoolVarP(&_bedrockagentcoreGetBrowserSession, "get-browser-session", "", false, "Get Browser Session")
	_bedrockagentcoreCmd.Flags().BoolVarP(&_bedrockagentcoreGetCodeInterpreterSession, "get-code-interpreter-session", "", false, "Get Code Interpreter Session")
	_bedrockagentcoreCmd.Flags().BoolVarP(&_bedrockagentcoreGetEvent, "get-event", "", false, "Get Event")
	_bedrockagentcoreCmd.Flags().BoolVarP(&_bedrockagentcoreGetMemoryRecord, "get-memory-record", "", false, "Get Memory Record")
	_bedrockagentcoreCmd.Flags().BoolVarP(&_bedrockagentcoreGetResourceApiKey, "get-resource-api-key", "", false, "Get Resource API Key")
	_bedrockagentcoreCmd.Flags().BoolVarP(&_bedrockagentcoreGetResourceOauth2Token, "get-resource-oauth2-token", "", false, "Get Resource Oauth2 Token")
	_bedrockagentcoreCmd.Flags().BoolVarP(&_bedrockagentcoreGetWorkloadAccessToken, "get-workload-access-token", "", false, "Get Workload Access Token")
	_bedrockagentcoreCmd.Flags().BoolVarP(&_bedrockagentcoreGetWorkloadAccessTokenForJWT, "get-workload-access-token-for-jwt", "", false, "Get Workload Access Token For Jwt")
	_bedrockagentcoreCmd.Flags().BoolVarP(&_bedrockagentcoreGetWorkloadAccessTokenForUserId, "get-workload-access-token-for-user-id", "", false, "Get Workload Access Token For User ID")
	_bedrockagentcoreCmd.Flags().BoolVarP(&_bedrockagentcoreInvokeAgentRuntime, "invoke-agent-runtime", "", false, "Invoke Agent Runtime")
	_bedrockagentcoreCmd.Flags().BoolVarP(&_bedrockagentcoreInvokeCodeInterpreter, "invoke-code-interpreter", "", false, "Invoke Code Interpreter")
	_bedrockagentcoreCmd.Flags().BoolVarP(&_bedrockagentcoreListActors, "list-actors", "", false, "List Actors")
	_bedrockagentcoreCmd.Flags().BoolVarP(&_bedrockagentcoreListBrowserSessions, "list-browser-sessions", "", false, "List Browser Sessions")
	_bedrockagentcoreCmd.Flags().BoolVarP(&_bedrockagentcoreListCodeInterpreterSessions, "list-code-interpreter-sessions", "", false, "List Code Interpreter Sessions")
	_bedrockagentcoreCmd.Flags().BoolVarP(&_bedrockagentcoreListEvents, "list-events", "", false, "List Events")
	_bedrockagentcoreCmd.Flags().BoolVarP(&_bedrockagentcoreListMemoryExtractionJobs, "list-memory-extraction-jobs", "", false, "List Memory Extraction Jobs")
	_bedrockagentcoreCmd.Flags().BoolVarP(&_bedrockagentcoreListMemoryRecords, "list-memory-records", "", false, "List Memory Records")
	_bedrockagentcoreCmd.Flags().BoolVarP(&_bedrockagentcoreListSessions, "list-sessions", "", false, "List Sessions")
	_bedrockagentcoreCmd.Flags().BoolVarP(&_bedrockagentcoreRetrieveMemoryRecords, "retrieve-memory-records", "", false, "Retrieve Memory Records")
	_bedrockagentcoreCmd.Flags().BoolVarP(&_bedrockagentcoreSaveBrowserSessionProfile, "save-browser-session-profile", "", false, "Save Browser Session Profile")
	_bedrockagentcoreCmd.Flags().BoolVarP(&_bedrockagentcoreStartBrowserSession, "start-browser-session", "", false, "Start Browser Session")
	_bedrockagentcoreCmd.Flags().BoolVarP(&_bedrockagentcoreStartCodeInterpreterSession, "start-code-interpreter-session", "", false, "Start Code Interpreter Session")
	_bedrockagentcoreCmd.Flags().BoolVarP(&_bedrockagentcoreStartMemoryExtractionJob, "start-memory-extraction-job", "", false, "Start Memory Extraction Job")
	_bedrockagentcoreCmd.Flags().BoolVarP(&_bedrockagentcoreStopBrowserSession, "stop-browser-session", "", false, "Stop Browser Session")
	_bedrockagentcoreCmd.Flags().BoolVarP(&_bedrockagentcoreStopCodeInterpreterSession, "stop-code-interpreter-session", "", false, "Stop Code Interpreter Session")
	_bedrockagentcoreCmd.Flags().BoolVarP(&_bedrockagentcoreStopRuntimeSession, "stop-runtime-session", "", false, "Stop Runtime Session")
	_bedrockagentcoreCmd.Flags().BoolVarP(&_bedrockagentcoreUpdateBrowserStream, "update-browser-stream", "", false, "Update Browser Stream")

}
