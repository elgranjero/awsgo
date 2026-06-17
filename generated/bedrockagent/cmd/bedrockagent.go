package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/bedrockagent"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// bedrockagentCmd represents the bedrockagent command
var _bedrockagentCmd = &cobra.Command{
	Use:   "bedrockagent",
	Short: "AWS bedrockagent CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := bedrockagent.NewFromConfig(cfg)
		if _bedrockagentAssociateAgentCollaborator {
			bedrockagent_AssociateAgentCollaborator(cfg, client)
			return
		}
		if _bedrockagentAssociateAgentKnowledgeBase {
			bedrockagent_AssociateAgentKnowledgeBase(cfg, client)
			return
		}
		if _bedrockagentCreateAgent {
			bedrockagent_CreateAgent(cfg, client)
			return
		}
		if _bedrockagentCreateAgentActionGroup {
			bedrockagent_CreateAgentActionGroup(cfg, client)
			return
		}
		if _bedrockagentCreateAgentAlias {
			bedrockagent_CreateAgentAlias(cfg, client)
			return
		}
		if _bedrockagentCreateDataSource {
			bedrockagent_CreateDataSource(cfg, client)
			return
		}
		if _bedrockagentCreateFlow {
			bedrockagent_CreateFlow(cfg, client)
			return
		}
		if _bedrockagentCreateFlowAlias {
			bedrockagent_CreateFlowAlias(cfg, client)
			return
		}
		if _bedrockagentCreateFlowVersion {
			bedrockagent_CreateFlowVersion(cfg, client)
			return
		}
		if _bedrockagentCreateKnowledgeBase {
			bedrockagent_CreateKnowledgeBase(cfg, client)
			return
		}
		if _bedrockagentCreatePrompt {
			bedrockagent_CreatePrompt(cfg, client)
			return
		}
		if _bedrockagentCreatePromptVersion {
			bedrockagent_CreatePromptVersion(cfg, client)
			return
		}
		if _bedrockagentDeleteAgent {
			bedrockagent_DeleteAgent(cfg, client)
			return
		}
		if _bedrockagentDeleteAgentActionGroup {
			bedrockagent_DeleteAgentActionGroup(cfg, client)
			return
		}
		if _bedrockagentDeleteAgentAlias {
			bedrockagent_DeleteAgentAlias(cfg, client)
			return
		}
		if _bedrockagentDeleteAgentVersion {
			bedrockagent_DeleteAgentVersion(cfg, client)
			return
		}
		if _bedrockagentDeleteDataSource {
			bedrockagent_DeleteDataSource(cfg, client)
			return
		}
		if _bedrockagentDeleteFlow {
			bedrockagent_DeleteFlow(cfg, client)
			return
		}
		if _bedrockagentDeleteFlowAlias {
			bedrockagent_DeleteFlowAlias(cfg, client)
			return
		}
		if _bedrockagentDeleteFlowVersion {
			bedrockagent_DeleteFlowVersion(cfg, client)
			return
		}
		if _bedrockagentDeleteKnowledgeBase {
			bedrockagent_DeleteKnowledgeBase(cfg, client)
			return
		}
		if _bedrockagentDeleteKnowledgeBaseDocuments {
			bedrockagent_DeleteKnowledgeBaseDocuments(cfg, client)
			return
		}
		if _bedrockagentDeletePrompt {
			bedrockagent_DeletePrompt(cfg, client)
			return
		}
		if _bedrockagentDisassociateAgentCollaborator {
			bedrockagent_DisassociateAgentCollaborator(cfg, client)
			return
		}
		if _bedrockagentDisassociateAgentKnowledgeBase {
			bedrockagent_DisassociateAgentKnowledgeBase(cfg, client)
			return
		}
		if _bedrockagentGetAgent {
			bedrockagent_GetAgent(cfg, client)
			return
		}
		if _bedrockagentGetAgentActionGroup {
			bedrockagent_GetAgentActionGroup(cfg, client)
			return
		}
		if _bedrockagentGetAgentAlias {
			bedrockagent_GetAgentAlias(cfg, client)
			return
		}
		if _bedrockagentGetAgentCollaborator {
			bedrockagent_GetAgentCollaborator(cfg, client)
			return
		}
		if _bedrockagentGetAgentKnowledgeBase {
			bedrockagent_GetAgentKnowledgeBase(cfg, client)
			return
		}
		if _bedrockagentGetAgentVersion {
			bedrockagent_GetAgentVersion(cfg, client)
			return
		}
		if _bedrockagentGetDataSource {
			bedrockagent_GetDataSource(cfg, client)
			return
		}
		if _bedrockagentGetFlow {
			bedrockagent_GetFlow(cfg, client)
			return
		}
		if _bedrockagentGetFlowAlias {
			bedrockagent_GetFlowAlias(cfg, client)
			return
		}
		if _bedrockagentGetFlowVersion {
			bedrockagent_GetFlowVersion(cfg, client)
			return
		}
		if _bedrockagentGetIngestionJob {
			bedrockagent_GetIngestionJob(cfg, client)
			return
		}
		if _bedrockagentGetKnowledgeBase {
			bedrockagent_GetKnowledgeBase(cfg, client)
			return
		}
		if _bedrockagentGetKnowledgeBaseDocuments {
			bedrockagent_GetKnowledgeBaseDocuments(cfg, client)
			return
		}
		if _bedrockagentGetPrompt {
			bedrockagent_GetPrompt(cfg, client)
			return
		}
		if _bedrockagentIngestKnowledgeBaseDocuments {
			bedrockagent_IngestKnowledgeBaseDocuments(cfg, client)
			return
		}
		if _bedrockagentListAgentActionGroups {
			bedrockagent_ListAgentActionGroups(cfg, client)
			return
		}
		if _bedrockagentListAgentAliases {
			bedrockagent_ListAgentAliases(cfg, client)
			return
		}
		if _bedrockagentListAgentCollaborators {
			bedrockagent_ListAgentCollaborators(cfg, client)
			return
		}
		if _bedrockagentListAgentKnowledgeBases {
			bedrockagent_ListAgentKnowledgeBases(cfg, client)
			return
		}
		if _bedrockagentListAgentVersions {
			bedrockagent_ListAgentVersions(cfg, client)
			return
		}
		if _bedrockagentListAgents {
			bedrockagent_ListAgents(cfg, client)
			return
		}
		if _bedrockagentListDataSources {
			bedrockagent_ListDataSources(cfg, client)
			return
		}
		if _bedrockagentListFlowAliases {
			bedrockagent_ListFlowAliases(cfg, client)
			return
		}
		if _bedrockagentListFlowVersions {
			bedrockagent_ListFlowVersions(cfg, client)
			return
		}
		if _bedrockagentListFlows {
			bedrockagent_ListFlows(cfg, client)
			return
		}
		if _bedrockagentListIngestionJobs {
			bedrockagent_ListIngestionJobs(cfg, client)
			return
		}
		if _bedrockagentListKnowledgeBaseDocuments {
			bedrockagent_ListKnowledgeBaseDocuments(cfg, client)
			return
		}
		if _bedrockagentListKnowledgeBases {
			bedrockagent_ListKnowledgeBases(cfg, client)
			return
		}
		if _bedrockagentListPrompts {
			bedrockagent_ListPrompts(cfg, client)
			return
		}
		if _bedrockagentListTagsForResource {
			bedrockagent_ListTagsForResource(cfg, client)
			return
		}
		if _bedrockagentPrepareAgent {
			bedrockagent_PrepareAgent(cfg, client)
			return
		}
		if _bedrockagentPrepareFlow {
			bedrockagent_PrepareFlow(cfg, client)
			return
		}
		if _bedrockagentStartIngestionJob {
			bedrockagent_StartIngestionJob(cfg, client)
			return
		}
		if _bedrockagentStopIngestionJob {
			bedrockagent_StopIngestionJob(cfg, client)
			return
		}
		if _bedrockagentTagResource {
			bedrockagent_TagResource(cfg, client)
			return
		}
		if _bedrockagentUntagResource {
			bedrockagent_UntagResource(cfg, client)
			return
		}
		if _bedrockagentUpdateAgent {
			bedrockagent_UpdateAgent(cfg, client)
			return
		}
		if _bedrockagentUpdateAgentActionGroup {
			bedrockagent_UpdateAgentActionGroup(cfg, client)
			return
		}
		if _bedrockagentUpdateAgentAlias {
			bedrockagent_UpdateAgentAlias(cfg, client)
			return
		}
		if _bedrockagentUpdateAgentCollaborator {
			bedrockagent_UpdateAgentCollaborator(cfg, client)
			return
		}
		if _bedrockagentUpdateAgentKnowledgeBase {
			bedrockagent_UpdateAgentKnowledgeBase(cfg, client)
			return
		}
		if _bedrockagentUpdateDataSource {
			bedrockagent_UpdateDataSource(cfg, client)
			return
		}
		if _bedrockagentUpdateFlow {
			bedrockagent_UpdateFlow(cfg, client)
			return
		}
		if _bedrockagentUpdateFlowAlias {
			bedrockagent_UpdateFlowAlias(cfg, client)
			return
		}
		if _bedrockagentUpdateKnowledgeBase {
			bedrockagent_UpdateKnowledgeBase(cfg, client)
			return
		}
		if _bedrockagentUpdatePrompt {
			bedrockagent_UpdatePrompt(cfg, client)
			return
		}
		if _bedrockagentValidateFlowDefinition {
			bedrockagent_ValidateFlowDefinition(cfg, client)
			return
		}

	},
}

var (
	_bedrockagentAssociateAgentCollaborator     bool
	_bedrockagentAssociateAgentKnowledgeBase    bool
	_bedrockagentCreateAgent                    bool
	_bedrockagentCreateAgentActionGroup         bool
	_bedrockagentCreateAgentAlias               bool
	_bedrockagentCreateDataSource               bool
	_bedrockagentCreateFlow                     bool
	_bedrockagentCreateFlowAlias                bool
	_bedrockagentCreateFlowVersion              bool
	_bedrockagentCreateKnowledgeBase            bool
	_bedrockagentCreatePrompt                   bool
	_bedrockagentCreatePromptVersion            bool
	_bedrockagentDeleteAgent                    bool
	_bedrockagentDeleteAgentActionGroup         bool
	_bedrockagentDeleteAgentAlias               bool
	_bedrockagentDeleteAgentVersion             bool
	_bedrockagentDeleteDataSource               bool
	_bedrockagentDeleteFlow                     bool
	_bedrockagentDeleteFlowAlias                bool
	_bedrockagentDeleteFlowVersion              bool
	_bedrockagentDeleteKnowledgeBase            bool
	_bedrockagentDeleteKnowledgeBaseDocuments   bool
	_bedrockagentDeletePrompt                   bool
	_bedrockagentDisassociateAgentCollaborator  bool
	_bedrockagentDisassociateAgentKnowledgeBase bool
	_bedrockagentGetAgent                       bool
	_bedrockagentGetAgentActionGroup            bool
	_bedrockagentGetAgentAlias                  bool
	_bedrockagentGetAgentCollaborator           bool
	_bedrockagentGetAgentKnowledgeBase          bool
	_bedrockagentGetAgentVersion                bool
	_bedrockagentGetDataSource                  bool
	_bedrockagentGetFlow                        bool
	_bedrockagentGetFlowAlias                   bool
	_bedrockagentGetFlowVersion                 bool
	_bedrockagentGetIngestionJob                bool
	_bedrockagentGetKnowledgeBase               bool
	_bedrockagentGetKnowledgeBaseDocuments      bool
	_bedrockagentGetPrompt                      bool
	_bedrockagentIngestKnowledgeBaseDocuments   bool
	_bedrockagentListAgentActionGroups          bool
	_bedrockagentListAgentAliases               bool
	_bedrockagentListAgentCollaborators         bool
	_bedrockagentListAgentKnowledgeBases        bool
	_bedrockagentListAgentVersions              bool
	_bedrockagentListAgents                     bool
	_bedrockagentListDataSources                bool
	_bedrockagentListFlowAliases                bool
	_bedrockagentListFlowVersions               bool
	_bedrockagentListFlows                      bool
	_bedrockagentListIngestionJobs              bool
	_bedrockagentListKnowledgeBaseDocuments     bool
	_bedrockagentListKnowledgeBases             bool
	_bedrockagentListPrompts                    bool
	_bedrockagentListTagsForResource            bool
	_bedrockagentPrepareAgent                   bool
	_bedrockagentPrepareFlow                    bool
	_bedrockagentStartIngestionJob              bool
	_bedrockagentStopIngestionJob               bool
	_bedrockagentTagResource                    bool
	_bedrockagentUntagResource                  bool
	_bedrockagentUpdateAgent                    bool
	_bedrockagentUpdateAgentActionGroup         bool
	_bedrockagentUpdateAgentAlias               bool
	_bedrockagentUpdateAgentCollaborator        bool
	_bedrockagentUpdateAgentKnowledgeBase       bool
	_bedrockagentUpdateDataSource               bool
	_bedrockagentUpdateFlow                     bool
	_bedrockagentUpdateFlowAlias                bool
	_bedrockagentUpdateKnowledgeBase            bool
	_bedrockagentUpdatePrompt                   bool
	_bedrockagentValidateFlowDefinition         bool

	_bedrockagentActionGroupExecutor               string
	_bedrockagentActionGroupId                     string
	_bedrockagentActionGroupName                   string
	_bedrockagentActionGroupState                  string
	_bedrockagentAgentAliasId                      string
	_bedrockagentAgentAliasName                    string
	_bedrockagentAgentCollaboration                string
	_bedrockagentAgentDescriptor                   string
	_bedrockagentAgentId                           string
	_bedrockagentAgentName                         string
	_bedrockagentAgentResourceRoleArn              string
	_bedrockagentAgentVersion                      string
	_bedrockagentAliasIdentifier                   string
	_bedrockagentAliasInvocationState              string
	_bedrockagentApiSchema                         string
	_bedrockagentClientToken                       string
	_bedrockagentCollaborationInstruction          string
	_bedrockagentCollaboratorId                    string
	_bedrockagentCollaboratorName                  string
	_bedrockagentConcurrencyConfiguration          string
	_bedrockagentCustomOrchestration               string
	_bedrockagentCustomerEncryptionKeyArn          string
	_bedrockagentDataDeletionPolicy                string
	_bedrockagentDataSourceConfiguration           string
	_bedrockagentDataSourceId                      string
	_bedrockagentDefaultVariant                    string
	_bedrockagentDefinition                        string
	_bedrockagentDescription                       string
	_bedrockagentDocumentIdentifiers               string
	_bedrockagentDocuments                         string
	_bedrockagentExecutionRoleArn                  string
	_bedrockagentFilters                           string
	_bedrockagentFlowIdentifier                    string
	_bedrockagentFlowVersion                       string
	_bedrockagentFoundationModel                   string
	_bedrockagentFunctionSchema                    string
	_bedrockagentGuardrailConfiguration            string
	_bedrockagentIdleSessionTTLInSeconds           string
	_bedrockagentIngestionJobId                    string
	_bedrockagentInstruction                       string
	_bedrockagentKnowledgeBaseConfiguration        string
	_bedrockagentKnowledgeBaseId                   string
	_bedrockagentKnowledgeBaseState                string
	_bedrockagentMaxResults                        string
	_bedrockagentMemoryConfiguration               string
	_bedrockagentName                              string
	_bedrockagentNextToken                         string
	_bedrockagentOrchestrationType                 string
	_bedrockagentParentActionGroupSignature        string
	_bedrockagentParentActionGroupSignatureParams  string
	_bedrockagentPromptIdentifier                  string
	_bedrockagentPromptOverrideConfiguration       string
	_bedrockagentPromptVersion                     string
	_bedrockagentRelayConversationHistory          string
	_bedrockagentResourceArn                       string
	_bedrockagentRoleArn                           string
	_bedrockagentRoutingConfiguration              string
	_bedrockagentServerSideEncryptionConfiguration string
	_bedrockagentSkipResourceInUseCheck            string
	_bedrockagentSortBy                            string
	_bedrockagentStorageConfiguration              string
	_bedrockagentTagKeys                           []string
	_bedrockagentTags                              string
	_bedrockagentVariants                          string
	_bedrockagentVectorIngestionConfiguration      string
)

// Makes an agent a collaborator for another agent.
func bedrockagent_AssociateAgentCollaborator(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.AssociateAgentCollaboratorInput{
		// AgentDescriptor: *types.AgentDescriptor, // Required
		// AgentId: *string, // Required
		// AgentVersion: *string, // Required
		// CollaborationInstruction: *string, // Required
		// CollaboratorName: *string, // Required
	}

	if len(_bedrockagentAgentDescriptor) > 0 {
		if err := assignInputField(input, "AgentDescriptor", _bedrockagentAgentDescriptor); err != nil {
			log.Errorf("invalid --agent-descriptor: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentAgentId) > 0 {
		input.AgentId = aws.String(_bedrockagentAgentId)
	}
	if len(_bedrockagentAgentVersion) > 0 {
		input.AgentVersion = aws.String(_bedrockagentAgentVersion)
	}
	if len(_bedrockagentCollaborationInstruction) > 0 {
		input.CollaborationInstruction = aws.String(_bedrockagentCollaborationInstruction)
	}
	if len(_bedrockagentCollaboratorName) > 0 {
		input.CollaboratorName = aws.String(_bedrockagentCollaboratorName)
	}
	if len(_bedrockagentClientToken) > 0 {
		input.ClientToken = aws.String(_bedrockagentClientToken)
	}
	if len(_bedrockagentRelayConversationHistory) > 0 {
		if err := assignInputField(input, "RelayConversationHistory", _bedrockagentRelayConversationHistory); err != nil {
			log.Errorf("invalid --relay-conversation-history: %s", err.Error())
			return
		}
	}

	if resp, err := client.AssociateAgentCollaborator(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates a knowledge base with an agent. If a knowledge base is associated
// and its indexState is set to Enabled , the agent queries the knowledge base for
// information to augment its response to the user.
func bedrockagent_AssociateAgentKnowledgeBase(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.AssociateAgentKnowledgeBaseInput{
		// AgentId: *string, // Required
		// AgentVersion: *string, // Required
		// Description: *string, // Required
		// KnowledgeBaseId: *string, // Required
	}

	if len(_bedrockagentAgentId) > 0 {
		input.AgentId = aws.String(_bedrockagentAgentId)
	}
	if len(_bedrockagentAgentVersion) > 0 {
		input.AgentVersion = aws.String(_bedrockagentAgentVersion)
	}
	if len(_bedrockagentDescription) > 0 {
		input.Description = aws.String(_bedrockagentDescription)
	}
	if len(_bedrockagentKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_bedrockagentKnowledgeBaseId)
	}
	if len(_bedrockagentKnowledgeBaseState) > 0 {
		if err := assignInputField(input, "KnowledgeBaseState", _bedrockagentKnowledgeBaseState); err != nil {
			log.Errorf("invalid --knowledge-base-state: %s", err.Error())
			return
		}
	}

	if resp, err := client.AssociateAgentKnowledgeBase(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an agent that orchestrates interactions between foundation models, data
// sources, software applications, user conversations, and APIs to carry out tasks
// to help customers.
//
// - Specify the following fields for security purposes.
//
// - agentResourceRoleArn – The Amazon Resource Name (ARN) of the role with
// permissions to invoke API operations on an agent.
//
// - (Optional) customerEncryptionKeyArn – The Amazon Resource Name (ARN) of a
// KMS key to encrypt the creation of the agent.
//
// - (Optional) idleSessionTTLinSeconds – Specify the number of seconds for which
// the agent should maintain session information. After this time expires, the
// subsequent InvokeAgent request begins a new session.
//
// - To enable your agent to retain conversational context across multiple
// sessions, include a memoryConfiguration object. For more information, see [Configure memory].
//
// - To override the default prompt behavior for agent orchestration and to use
// advanced prompts, include a promptOverrideConfiguration object. For more
// information, see [Advanced prompts].
//
// - If your agent fails to be created, the response returns a list of
// failureReasons alongside a list of recommendedActions for you to troubleshoot.
//
// - The agent instructions will not be honored if your agent has only one
// knowledge base, uses default prompts, has no action group, and user input is
// disabled.
//
// [Advanced prompts]: https://docs.aws.amazon.com/bedrock/latest/userguide/advanced-prompts.html
// [Configure memory]: https://docs.aws.amazon.com/bedrock/latest/userguide/agents-configure-memory.html
func bedrockagent_CreateAgent(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.CreateAgentInput{
		// AgentName: *string, // Required
	}

	if len(_bedrockagentAgentName) > 0 {
		input.AgentName = aws.String(_bedrockagentAgentName)
	}
	if len(_bedrockagentAgentCollaboration) > 0 {
		if err := assignInputField(input, "AgentCollaboration", _bedrockagentAgentCollaboration); err != nil {
			log.Errorf("invalid --agent-collaboration: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentAgentResourceRoleArn) > 0 {
		input.AgentResourceRoleArn = aws.String(_bedrockagentAgentResourceRoleArn)
	}
	if len(_bedrockagentClientToken) > 0 {
		input.ClientToken = aws.String(_bedrockagentClientToken)
	}
	if len(_bedrockagentCustomOrchestration) > 0 {
		if err := assignInputField(input, "CustomOrchestration", _bedrockagentCustomOrchestration); err != nil {
			log.Errorf("invalid --custom-orchestration: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentCustomerEncryptionKeyArn) > 0 {
		input.CustomerEncryptionKeyArn = aws.String(_bedrockagentCustomerEncryptionKeyArn)
	}
	if len(_bedrockagentDescription) > 0 {
		input.Description = aws.String(_bedrockagentDescription)
	}
	if len(_bedrockagentFoundationModel) > 0 {
		input.FoundationModel = aws.String(_bedrockagentFoundationModel)
	}
	if len(_bedrockagentGuardrailConfiguration) > 0 {
		if err := assignInputField(input, "GuardrailConfiguration", _bedrockagentGuardrailConfiguration); err != nil {
			log.Errorf("invalid --guardrail-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentIdleSessionTTLInSeconds) > 0 {
		if err := assignInputField(input, "IdleSessionTTLInSeconds", _bedrockagentIdleSessionTTLInSeconds); err != nil {
			log.Errorf("invalid --idle-session-ttlin-seconds: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentInstruction) > 0 {
		input.Instruction = aws.String(_bedrockagentInstruction)
	}
	if len(_bedrockagentMemoryConfiguration) > 0 {
		if err := assignInputField(input, "MemoryConfiguration", _bedrockagentMemoryConfiguration); err != nil {
			log.Errorf("invalid --memory-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentOrchestrationType) > 0 {
		if err := assignInputField(input, "OrchestrationType", _bedrockagentOrchestrationType); err != nil {
			log.Errorf("invalid --orchestration-type: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentPromptOverrideConfiguration) > 0 {
		if err := assignInputField(input, "PromptOverrideConfiguration", _bedrockagentPromptOverrideConfiguration); err != nil {
			log.Errorf("invalid --prompt-override-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentTags) > 0 {
		if err := assignInputField(input, "Tags", _bedrockagentTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAgent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an action group for an agent. An action group represents the actions
// that an agent can carry out for the customer by defining the APIs that an agent
// can call and the logic for calling them.
//
// To allow your agent to request the user for additional information when trying
// to complete a task, add an action group with the parentActionGroupSignature
// field set to AMAZON.UserInput .
//
// To allow your agent to generate, run, and troubleshoot code when trying to
// complete a task, add an action group with the parentActionGroupSignature field
// set to AMAZON.CodeInterpreter .
//
// You must leave the description , apiSchema , and actionGroupExecutor fields
// blank for this action group. During orchestration, if your agent determines that
// it needs to invoke an API in an action group, but doesn't have enough
// information to complete the API request, it will invoke this action group
// instead and return an [Observation]reprompting the user for more information.
//
// [Observation]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent-runtime_Observation.html
func bedrockagent_CreateAgentActionGroup(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.CreateAgentActionGroupInput{
		// ActionGroupName: *string, // Required
		// AgentId: *string, // Required
		// AgentVersion: *string, // Required
	}

	if len(_bedrockagentActionGroupName) > 0 {
		input.ActionGroupName = aws.String(_bedrockagentActionGroupName)
	}
	if len(_bedrockagentAgentId) > 0 {
		input.AgentId = aws.String(_bedrockagentAgentId)
	}
	if len(_bedrockagentAgentVersion) > 0 {
		input.AgentVersion = aws.String(_bedrockagentAgentVersion)
	}
	if len(_bedrockagentActionGroupExecutor) > 0 {
		if err := assignInputField(input, "ActionGroupExecutor", _bedrockagentActionGroupExecutor); err != nil {
			log.Errorf("invalid --action-group-executor: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentActionGroupState) > 0 {
		if err := assignInputField(input, "ActionGroupState", _bedrockagentActionGroupState); err != nil {
			log.Errorf("invalid --action-group-state: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentApiSchema) > 0 {
		if err := assignInputField(input, "ApiSchema", _bedrockagentApiSchema); err != nil {
			log.Errorf("invalid --api-schema: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentClientToken) > 0 {
		input.ClientToken = aws.String(_bedrockagentClientToken)
	}
	if len(_bedrockagentDescription) > 0 {
		input.Description = aws.String(_bedrockagentDescription)
	}
	if len(_bedrockagentFunctionSchema) > 0 {
		if err := assignInputField(input, "FunctionSchema", _bedrockagentFunctionSchema); err != nil {
			log.Errorf("invalid --function-schema: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentParentActionGroupSignature) > 0 {
		if err := assignInputField(input, "ParentActionGroupSignature", _bedrockagentParentActionGroupSignature); err != nil {
			log.Errorf("invalid --parent-action-group-signature: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentParentActionGroupSignatureParams) > 0 {
		if err := assignInputField(input, "ParentActionGroupSignatureParams", _bedrockagentParentActionGroupSignatureParams); err != nil {
			log.Errorf("invalid --parent-action-group-signature-params: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAgentActionGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an alias of an agent that can be used to deploy the agent.
func bedrockagent_CreateAgentAlias(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.CreateAgentAliasInput{
		// AgentAliasName: *string, // Required
		// AgentId: *string, // Required
	}

	if len(_bedrockagentAgentAliasName) > 0 {
		input.AgentAliasName = aws.String(_bedrockagentAgentAliasName)
	}
	if len(_bedrockagentAgentId) > 0 {
		input.AgentId = aws.String(_bedrockagentAgentId)
	}
	if len(_bedrockagentClientToken) > 0 {
		input.ClientToken = aws.String(_bedrockagentClientToken)
	}
	if len(_bedrockagentDescription) > 0 {
		input.Description = aws.String(_bedrockagentDescription)
	}
	if len(_bedrockagentRoutingConfiguration) > 0 {
		if err := assignInputField(input, "RoutingConfiguration", _bedrockagentRoutingConfiguration); err != nil {
			log.Errorf("invalid --routing-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentTags) > 0 {
		if err := assignInputField(input, "Tags", _bedrockagentTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAgentAlias(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Connects a knowledge base to a data source. You specify the configuration for
// the specific data source service in the dataSourceConfiguration field.
//
// You can't change the chunkingConfiguration after you create the data source
// connector.
func bedrockagent_CreateDataSource(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.CreateDataSourceInput{
		// DataSourceConfiguration: *types.DataSourceConfiguration, // Required
		// KnowledgeBaseId: *string, // Required
		// Name: *string, // Required
	}

	if len(_bedrockagentDataSourceConfiguration) > 0 {
		if err := assignInputField(input, "DataSourceConfiguration", _bedrockagentDataSourceConfiguration); err != nil {
			log.Errorf("invalid --data-source-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_bedrockagentKnowledgeBaseId)
	}
	if len(_bedrockagentName) > 0 {
		input.Name = aws.String(_bedrockagentName)
	}
	if len(_bedrockagentClientToken) > 0 {
		input.ClientToken = aws.String(_bedrockagentClientToken)
	}
	if len(_bedrockagentDataDeletionPolicy) > 0 {
		if err := assignInputField(input, "DataDeletionPolicy", _bedrockagentDataDeletionPolicy); err != nil {
			log.Errorf("invalid --data-deletion-policy: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentDescription) > 0 {
		input.Description = aws.String(_bedrockagentDescription)
	}
	if len(_bedrockagentServerSideEncryptionConfiguration) > 0 {
		if err := assignInputField(input, "ServerSideEncryptionConfiguration", _bedrockagentServerSideEncryptionConfiguration); err != nil {
			log.Errorf("invalid --server-side-encryption-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentVectorIngestionConfiguration) > 0 {
		if err := assignInputField(input, "VectorIngestionConfiguration", _bedrockagentVectorIngestionConfiguration); err != nil {
			log.Errorf("invalid --vector-ingestion-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDataSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a prompt flow that you can use to send an input through various steps
// to yield an output. Configure nodes, each of which corresponds to a step of the
// flow, and create connections between the nodes to create paths to different
// outputs. For more information, see [How it works]and [Create a flow in Amazon Bedrock] in the Amazon Bedrock User Guide.
//
// [How it works]: https://docs.aws.amazon.com/bedrock/latest/userguide/flows-how-it-works.html
// [Create a flow in Amazon Bedrock]: https://docs.aws.amazon.com/bedrock/latest/userguide/flows-create.html
func bedrockagent_CreateFlow(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.CreateFlowInput{
		// ExecutionRoleArn: *string, // Required
		// Name: *string, // Required
	}

	if len(_bedrockagentExecutionRoleArn) > 0 {
		input.ExecutionRoleArn = aws.String(_bedrockagentExecutionRoleArn)
	}
	if len(_bedrockagentName) > 0 {
		input.Name = aws.String(_bedrockagentName)
	}
	if len(_bedrockagentClientToken) > 0 {
		input.ClientToken = aws.String(_bedrockagentClientToken)
	}
	if len(_bedrockagentCustomerEncryptionKeyArn) > 0 {
		input.CustomerEncryptionKeyArn = aws.String(_bedrockagentCustomerEncryptionKeyArn)
	}
	if len(_bedrockagentDefinition) > 0 {
		if err := assignInputField(input, "Definition", _bedrockagentDefinition); err != nil {
			log.Errorf("invalid --definition: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentDescription) > 0 {
		input.Description = aws.String(_bedrockagentDescription)
	}
	if len(_bedrockagentTags) > 0 {
		if err := assignInputField(input, "Tags", _bedrockagentTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateFlow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an alias of a flow for deployment. For more information, see [Deploy a flow in Amazon Bedrock] in the
// Amazon Bedrock User Guide.
//
// [Deploy a flow in Amazon Bedrock]: https://docs.aws.amazon.com/bedrock/latest/userguide/flows-deploy.html
func bedrockagent_CreateFlowAlias(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.CreateFlowAliasInput{
		// FlowIdentifier: *string, // Required
		// Name: *string, // Required
		// RoutingConfiguration: []types.FlowAliasRoutingConfigurationListItem, // Required
	}

	if len(_bedrockagentFlowIdentifier) > 0 {
		input.FlowIdentifier = aws.String(_bedrockagentFlowIdentifier)
	}
	if len(_bedrockagentName) > 0 {
		input.Name = aws.String(_bedrockagentName)
	}
	if len(_bedrockagentRoutingConfiguration) > 0 {
		if err := assignInputField(input, "RoutingConfiguration", _bedrockagentRoutingConfiguration); err != nil {
			log.Errorf("invalid --routing-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentClientToken) > 0 {
		input.ClientToken = aws.String(_bedrockagentClientToken)
	}
	if len(_bedrockagentConcurrencyConfiguration) > 0 {
		if err := assignInputField(input, "ConcurrencyConfiguration", _bedrockagentConcurrencyConfiguration); err != nil {
			log.Errorf("invalid --concurrency-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentDescription) > 0 {
		input.Description = aws.String(_bedrockagentDescription)
	}
	if len(_bedrockagentTags) > 0 {
		if err := assignInputField(input, "Tags", _bedrockagentTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateFlowAlias(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a version of the flow that you can deploy. For more information, see [Deploy a flow in Amazon Bedrock]
// in the Amazon Bedrock User Guide.
//
// [Deploy a flow in Amazon Bedrock]: https://docs.aws.amazon.com/bedrock/latest/userguide/flows-deploy.html
func bedrockagent_CreateFlowVersion(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.CreateFlowVersionInput{
		// FlowIdentifier: *string, // Required
	}

	if len(_bedrockagentFlowIdentifier) > 0 {
		input.FlowIdentifier = aws.String(_bedrockagentFlowIdentifier)
	}
	if len(_bedrockagentClientToken) > 0 {
		input.ClientToken = aws.String(_bedrockagentClientToken)
	}
	if len(_bedrockagentDescription) > 0 {
		input.Description = aws.String(_bedrockagentDescription)
	}

	if resp, err := client.CreateFlowVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a knowledge base. A knowledge base contains your data sources so that
// Large Language Models (LLMs) can use your data. To create a knowledge base, you
// must first set up your data sources and configure a supported vector store. For
// more information, see [Set up a knowledge base].
//
// If you prefer to let Amazon Bedrock create and manage a vector store for you in
// Amazon OpenSearch Service, use the console. For more information, see [Create a knowledge base].
//
// - Provide the name and an optional description .
//
// - Provide the Amazon Resource Name (ARN) with permissions to create a
// knowledge base in the roleArn field.
//
// - Provide the embedding model to use in the embeddingModelArn field in the
// knowledgeBaseConfiguration object.
//
// - Provide the configuration for your vector store in the storageConfiguration
// object.
//
// - For an Amazon OpenSearch Service database, use the
// opensearchServerlessConfiguration object. For more information, see [Create a vector store in Amazon OpenSearch Service].
//
// - For an Amazon Aurora database, use the RdsConfiguration object. For more
// information, see [Create a vector store in Amazon Aurora].
//
// - For a Pinecone database, use the pineconeConfiguration object. For more
// information, see [Create a vector store in Pinecone].
//
// - For a Redis Enterprise Cloud database, use the
// redisEnterpriseCloudConfiguration object. For more information, see [Create a vector store in Redis Enterprise Cloud].
//
// [Create a knowledge base]: https://docs.aws.amazon.com/bedrock/latest/userguide/knowledge-base-create
// [Create a vector store in Amazon OpenSearch Service]: https://docs.aws.amazon.com/bedrock/latest/userguide/knowledge-base-setup-oss.html
// [Create a vector store in Redis Enterprise Cloud]: https://docs.aws.amazon.com/bedrock/latest/userguide/knowledge-base-setup-redis.html
// [Set up a knowledge base]: https://docs.aws.amazon.com/bedrock/latest/userguide/knowlege-base-prereq.html
// [Create a vector store in Amazon Aurora]: https://docs.aws.amazon.com/bedrock/latest/userguide/knowledge-base-setup-rds.html
// [Create a vector store in Pinecone]: https://docs.aws.amazon.com/bedrock/latest/userguide/knowledge-base-setup-pinecone.html
func bedrockagent_CreateKnowledgeBase(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.CreateKnowledgeBaseInput{
		// KnowledgeBaseConfiguration: *types.KnowledgeBaseConfiguration, // Required
		// Name: *string, // Required
		// RoleArn: *string, // Required
	}

	if len(_bedrockagentKnowledgeBaseConfiguration) > 0 {
		if err := assignInputField(input, "KnowledgeBaseConfiguration", _bedrockagentKnowledgeBaseConfiguration); err != nil {
			log.Errorf("invalid --knowledge-base-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentName) > 0 {
		input.Name = aws.String(_bedrockagentName)
	}
	if len(_bedrockagentRoleArn) > 0 {
		input.RoleArn = aws.String(_bedrockagentRoleArn)
	}
	if len(_bedrockagentClientToken) > 0 {
		input.ClientToken = aws.String(_bedrockagentClientToken)
	}
	if len(_bedrockagentDescription) > 0 {
		input.Description = aws.String(_bedrockagentDescription)
	}
	if len(_bedrockagentStorageConfiguration) > 0 {
		if err := assignInputField(input, "StorageConfiguration", _bedrockagentStorageConfiguration); err != nil {
			log.Errorf("invalid --storage-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentTags) > 0 {
		if err := assignInputField(input, "Tags", _bedrockagentTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateKnowledgeBase(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a prompt in your prompt library that you can add to a flow. For more
// information, see [Prompt management in Amazon Bedrock], [Create a prompt using Prompt management] and [Prompt flows in Amazon Bedrock] in the Amazon Bedrock User Guide.
//
// [Prompt flows in Amazon Bedrock]: https://docs.aws.amazon.com/bedrock/latest/userguide/flows.html
// [Prompt management in Amazon Bedrock]: https://docs.aws.amazon.com/bedrock/latest/userguide/prompt-management.html
// [Create a prompt using Prompt management]: https://docs.aws.amazon.com/bedrock/latest/userguide/prompt-management-create.html
func bedrockagent_CreatePrompt(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.CreatePromptInput{
		// Name: *string, // Required
	}

	if len(_bedrockagentName) > 0 {
		input.Name = aws.String(_bedrockagentName)
	}
	if len(_bedrockagentClientToken) > 0 {
		input.ClientToken = aws.String(_bedrockagentClientToken)
	}
	if len(_bedrockagentCustomerEncryptionKeyArn) > 0 {
		input.CustomerEncryptionKeyArn = aws.String(_bedrockagentCustomerEncryptionKeyArn)
	}
	if len(_bedrockagentDefaultVariant) > 0 {
		input.DefaultVariant = aws.String(_bedrockagentDefaultVariant)
	}
	if len(_bedrockagentDescription) > 0 {
		input.Description = aws.String(_bedrockagentDescription)
	}
	if len(_bedrockagentTags) > 0 {
		if err := assignInputField(input, "Tags", _bedrockagentTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentVariants) > 0 {
		if err := assignInputField(input, "Variants", _bedrockagentVariants); err != nil {
			log.Errorf("invalid --variants: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreatePrompt(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a static snapshot of your prompt that can be deployed to production.
// For more information, see [Deploy prompts using Prompt management by creating versions]in the Amazon Bedrock User Guide.
//
// [Deploy prompts using Prompt management by creating versions]: https://docs.aws.amazon.com/bedrock/latest/userguide/prompt-management-deploy.html
func bedrockagent_CreatePromptVersion(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.CreatePromptVersionInput{
		// PromptIdentifier: *string, // Required
	}

	if len(_bedrockagentPromptIdentifier) > 0 {
		input.PromptIdentifier = aws.String(_bedrockagentPromptIdentifier)
	}
	if len(_bedrockagentClientToken) > 0 {
		input.ClientToken = aws.String(_bedrockagentClientToken)
	}
	if len(_bedrockagentDescription) > 0 {
		input.Description = aws.String(_bedrockagentDescription)
	}
	if len(_bedrockagentTags) > 0 {
		if err := assignInputField(input, "Tags", _bedrockagentTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreatePromptVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an agent.
func bedrockagent_DeleteAgent(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.DeleteAgentInput{
		// AgentId: *string, // Required
	}

	if len(_bedrockagentAgentId) > 0 {
		input.AgentId = aws.String(_bedrockagentAgentId)
	}
	if len(_bedrockagentSkipResourceInUseCheck) > 0 {
		if err := assignInputField(input, "SkipResourceInUseCheck", _bedrockagentSkipResourceInUseCheck); err != nil {
			log.Errorf("invalid --skip-resource-in-use-check: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteAgent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an action group in an agent.
func bedrockagent_DeleteAgentActionGroup(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.DeleteAgentActionGroupInput{
		// ActionGroupId: *string, // Required
		// AgentId: *string, // Required
		// AgentVersion: *string, // Required
	}

	if len(_bedrockagentActionGroupId) > 0 {
		input.ActionGroupId = aws.String(_bedrockagentActionGroupId)
	}
	if len(_bedrockagentAgentId) > 0 {
		input.AgentId = aws.String(_bedrockagentAgentId)
	}
	if len(_bedrockagentAgentVersion) > 0 {
		input.AgentVersion = aws.String(_bedrockagentAgentVersion)
	}
	if len(_bedrockagentSkipResourceInUseCheck) > 0 {
		if err := assignInputField(input, "SkipResourceInUseCheck", _bedrockagentSkipResourceInUseCheck); err != nil {
			log.Errorf("invalid --skip-resource-in-use-check: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteAgentActionGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an alias of an agent.
func bedrockagent_DeleteAgentAlias(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.DeleteAgentAliasInput{
		// AgentAliasId: *string, // Required
		// AgentId: *string, // Required
	}

	if len(_bedrockagentAgentAliasId) > 0 {
		input.AgentAliasId = aws.String(_bedrockagentAgentAliasId)
	}
	if len(_bedrockagentAgentId) > 0 {
		input.AgentId = aws.String(_bedrockagentAgentId)
	}

	if resp, err := client.DeleteAgentAlias(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a version of an agent.
func bedrockagent_DeleteAgentVersion(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.DeleteAgentVersionInput{
		// AgentId: *string, // Required
		// AgentVersion: *string, // Required
	}

	if len(_bedrockagentAgentId) > 0 {
		input.AgentId = aws.String(_bedrockagentAgentId)
	}
	if len(_bedrockagentAgentVersion) > 0 {
		input.AgentVersion = aws.String(_bedrockagentAgentVersion)
	}
	if len(_bedrockagentSkipResourceInUseCheck) > 0 {
		if err := assignInputField(input, "SkipResourceInUseCheck", _bedrockagentSkipResourceInUseCheck); err != nil {
			log.Errorf("invalid --skip-resource-in-use-check: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteAgentVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a data source from a knowledge base.
func bedrockagent_DeleteDataSource(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.DeleteDataSourceInput{
		// DataSourceId: *string, // Required
		// KnowledgeBaseId: *string, // Required
	}

	if len(_bedrockagentDataSourceId) > 0 {
		input.DataSourceId = aws.String(_bedrockagentDataSourceId)
	}
	if len(_bedrockagentKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_bedrockagentKnowledgeBaseId)
	}

	if resp, err := client.DeleteDataSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a flow.
func bedrockagent_DeleteFlow(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.DeleteFlowInput{
		// FlowIdentifier: *string, // Required
	}

	if len(_bedrockagentFlowIdentifier) > 0 {
		input.FlowIdentifier = aws.String(_bedrockagentFlowIdentifier)
	}
	if len(_bedrockagentSkipResourceInUseCheck) > 0 {
		if err := assignInputField(input, "SkipResourceInUseCheck", _bedrockagentSkipResourceInUseCheck); err != nil {
			log.Errorf("invalid --skip-resource-in-use-check: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteFlow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an alias of a flow.
func bedrockagent_DeleteFlowAlias(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.DeleteFlowAliasInput{
		// AliasIdentifier: *string, // Required
		// FlowIdentifier: *string, // Required
	}

	if len(_bedrockagentAliasIdentifier) > 0 {
		input.AliasIdentifier = aws.String(_bedrockagentAliasIdentifier)
	}
	if len(_bedrockagentFlowIdentifier) > 0 {
		input.FlowIdentifier = aws.String(_bedrockagentFlowIdentifier)
	}

	if resp, err := client.DeleteFlowAlias(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a version of a flow.
func bedrockagent_DeleteFlowVersion(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.DeleteFlowVersionInput{
		// FlowIdentifier: *string, // Required
		// FlowVersion: *string, // Required
	}

	if len(_bedrockagentFlowIdentifier) > 0 {
		input.FlowIdentifier = aws.String(_bedrockagentFlowIdentifier)
	}
	if len(_bedrockagentFlowVersion) > 0 {
		input.FlowVersion = aws.String(_bedrockagentFlowVersion)
	}
	if len(_bedrockagentSkipResourceInUseCheck) > 0 {
		if err := assignInputField(input, "SkipResourceInUseCheck", _bedrockagentSkipResourceInUseCheck); err != nil {
			log.Errorf("invalid --skip-resource-in-use-check: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteFlowVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a knowledge base. Before deleting a knowledge base, you should
// disassociate the knowledge base from any agents that it is associated with by
// making a [DisassociateAgentKnowledgeBase]request.
//
// [DisassociateAgentKnowledgeBase]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent_DisassociateAgentKnowledgeBase.html
func bedrockagent_DeleteKnowledgeBase(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.DeleteKnowledgeBaseInput{
		// KnowledgeBaseId: *string, // Required
	}

	if len(_bedrockagentKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_bedrockagentKnowledgeBaseId)
	}

	if resp, err := client.DeleteKnowledgeBase(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes documents from a data source and syncs the changes to the knowledge
// base that is connected to it. For more information, see [Ingest changes directly into a knowledge base]in the Amazon Bedrock
// User Guide.
//
// [Ingest changes directly into a knowledge base]: https://docs.aws.amazon.com/bedrock/latest/userguide/kb-direct-ingestion.html
func bedrockagent_DeleteKnowledgeBaseDocuments(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.DeleteKnowledgeBaseDocumentsInput{
		// DataSourceId: *string, // Required
		// DocumentIdentifiers: []types.DocumentIdentifier, // Required
		// KnowledgeBaseId: *string, // Required
	}

	if len(_bedrockagentDataSourceId) > 0 {
		input.DataSourceId = aws.String(_bedrockagentDataSourceId)
	}
	if len(_bedrockagentDocumentIdentifiers) > 0 {
		if err := assignInputField(input, "DocumentIdentifiers", _bedrockagentDocumentIdentifiers); err != nil {
			log.Errorf("invalid --document-identifiers: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_bedrockagentKnowledgeBaseId)
	}
	if len(_bedrockagentClientToken) > 0 {
		input.ClientToken = aws.String(_bedrockagentClientToken)
	}

	if resp, err := client.DeleteKnowledgeBaseDocuments(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a prompt or a version of it, depending on whether you include the
// promptVersion field or not. For more information, see [Delete prompts from the Prompt management tool] and [Delete a version of a prompt from the Prompt management tool] in the Amazon
// Bedrock User Guide.
//
// [Delete a version of a prompt from the Prompt management tool]: https://docs.aws.amazon.com/bedrock/latest/userguide/prompt-management-deploy.html#prompt-management-versions-delete.html
// [Delete prompts from the Prompt management tool]: https://docs.aws.amazon.com/bedrock/latest/userguide/prompt-management-manage.html#prompt-management-delete.html
func bedrockagent_DeletePrompt(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.DeletePromptInput{
		// PromptIdentifier: *string, // Required
	}

	if len(_bedrockagentPromptIdentifier) > 0 {
		input.PromptIdentifier = aws.String(_bedrockagentPromptIdentifier)
	}
	if len(_bedrockagentPromptVersion) > 0 {
		input.PromptVersion = aws.String(_bedrockagentPromptVersion)
	}

	if resp, err := client.DeletePrompt(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates an agent collaborator.
func bedrockagent_DisassociateAgentCollaborator(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.DisassociateAgentCollaboratorInput{
		// AgentId: *string, // Required
		// AgentVersion: *string, // Required
		// CollaboratorId: *string, // Required
	}

	if len(_bedrockagentAgentId) > 0 {
		input.AgentId = aws.String(_bedrockagentAgentId)
	}
	if len(_bedrockagentAgentVersion) > 0 {
		input.AgentVersion = aws.String(_bedrockagentAgentVersion)
	}
	if len(_bedrockagentCollaboratorId) > 0 {
		input.CollaboratorId = aws.String(_bedrockagentCollaboratorId)
	}

	if resp, err := client.DisassociateAgentCollaborator(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates a knowledge base from an agent.
func bedrockagent_DisassociateAgentKnowledgeBase(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.DisassociateAgentKnowledgeBaseInput{
		// AgentId: *string, // Required
		// AgentVersion: *string, // Required
		// KnowledgeBaseId: *string, // Required
	}

	if len(_bedrockagentAgentId) > 0 {
		input.AgentId = aws.String(_bedrockagentAgentId)
	}
	if len(_bedrockagentAgentVersion) > 0 {
		input.AgentVersion = aws.String(_bedrockagentAgentVersion)
	}
	if len(_bedrockagentKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_bedrockagentKnowledgeBaseId)
	}

	if resp, err := client.DisassociateAgentKnowledgeBase(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about an agent.
func bedrockagent_GetAgent(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.GetAgentInput{
		// AgentId: *string, // Required
	}

	if len(_bedrockagentAgentId) > 0 {
		input.AgentId = aws.String(_bedrockagentAgentId)
	}

	if resp, err := client.GetAgent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about an action group for an agent.
func bedrockagent_GetAgentActionGroup(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.GetAgentActionGroupInput{
		// ActionGroupId: *string, // Required
		// AgentId: *string, // Required
		// AgentVersion: *string, // Required
	}

	if len(_bedrockagentActionGroupId) > 0 {
		input.ActionGroupId = aws.String(_bedrockagentActionGroupId)
	}
	if len(_bedrockagentAgentId) > 0 {
		input.AgentId = aws.String(_bedrockagentAgentId)
	}
	if len(_bedrockagentAgentVersion) > 0 {
		input.AgentVersion = aws.String(_bedrockagentAgentVersion)
	}

	if resp, err := client.GetAgentActionGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about an alias of an agent.
func bedrockagent_GetAgentAlias(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.GetAgentAliasInput{
		// AgentAliasId: *string, // Required
		// AgentId: *string, // Required
	}

	if len(_bedrockagentAgentAliasId) > 0 {
		input.AgentAliasId = aws.String(_bedrockagentAgentAliasId)
	}
	if len(_bedrockagentAgentId) > 0 {
		input.AgentId = aws.String(_bedrockagentAgentId)
	}

	if resp, err := client.GetAgentAlias(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about an agent's collaborator.
func bedrockagent_GetAgentCollaborator(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.GetAgentCollaboratorInput{
		// AgentId: *string, // Required
		// AgentVersion: *string, // Required
		// CollaboratorId: *string, // Required
	}

	if len(_bedrockagentAgentId) > 0 {
		input.AgentId = aws.String(_bedrockagentAgentId)
	}
	if len(_bedrockagentAgentVersion) > 0 {
		input.AgentVersion = aws.String(_bedrockagentAgentVersion)
	}
	if len(_bedrockagentCollaboratorId) > 0 {
		input.CollaboratorId = aws.String(_bedrockagentCollaboratorId)
	}

	if resp, err := client.GetAgentCollaborator(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a knowledge base associated with an agent.
func bedrockagent_GetAgentKnowledgeBase(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.GetAgentKnowledgeBaseInput{
		// AgentId: *string, // Required
		// AgentVersion: *string, // Required
		// KnowledgeBaseId: *string, // Required
	}

	if len(_bedrockagentAgentId) > 0 {
		input.AgentId = aws.String(_bedrockagentAgentId)
	}
	if len(_bedrockagentAgentVersion) > 0 {
		input.AgentVersion = aws.String(_bedrockagentAgentVersion)
	}
	if len(_bedrockagentKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_bedrockagentKnowledgeBaseId)
	}

	if resp, err := client.GetAgentKnowledgeBase(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets details about a version of an agent.
func bedrockagent_GetAgentVersion(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.GetAgentVersionInput{
		// AgentId: *string, // Required
		// AgentVersion: *string, // Required
	}

	if len(_bedrockagentAgentId) > 0 {
		input.AgentId = aws.String(_bedrockagentAgentId)
	}
	if len(_bedrockagentAgentVersion) > 0 {
		input.AgentVersion = aws.String(_bedrockagentAgentVersion)
	}

	if resp, err := client.GetAgentVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a data source.
func bedrockagent_GetDataSource(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.GetDataSourceInput{
		// DataSourceId: *string, // Required
		// KnowledgeBaseId: *string, // Required
	}

	if len(_bedrockagentDataSourceId) > 0 {
		input.DataSourceId = aws.String(_bedrockagentDataSourceId)
	}
	if len(_bedrockagentKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_bedrockagentKnowledgeBaseId)
	}

	if resp, err := client.GetDataSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a flow. For more information, see [Manage a flow in Amazon Bedrock] in the Amazon
// Bedrock User Guide.
//
// [Manage a flow in Amazon Bedrock]: https://docs.aws.amazon.com/bedrock/latest/userguide/flows-manage.html
func bedrockagent_GetFlow(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.GetFlowInput{
		// FlowIdentifier: *string, // Required
	}

	if len(_bedrockagentFlowIdentifier) > 0 {
		input.FlowIdentifier = aws.String(_bedrockagentFlowIdentifier)
	}

	if resp, err := client.GetFlow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a flow. For more information, see [Deploy a flow in Amazon Bedrock] in the Amazon
// Bedrock User Guide.
//
// [Deploy a flow in Amazon Bedrock]: https://docs.aws.amazon.com/bedrock/latest/userguide/flows-deploy.html
func bedrockagent_GetFlowAlias(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.GetFlowAliasInput{
		// AliasIdentifier: *string, // Required
		// FlowIdentifier: *string, // Required
	}

	if len(_bedrockagentAliasIdentifier) > 0 {
		input.AliasIdentifier = aws.String(_bedrockagentAliasIdentifier)
	}
	if len(_bedrockagentFlowIdentifier) > 0 {
		input.FlowIdentifier = aws.String(_bedrockagentFlowIdentifier)
	}

	if resp, err := client.GetFlowAlias(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a version of a flow. For more information, see [Deploy a flow in Amazon Bedrock] in
// the Amazon Bedrock User Guide.
//
// [Deploy a flow in Amazon Bedrock]: https://docs.aws.amazon.com/bedrock/latest/userguide/flows-deploy.html
func bedrockagent_GetFlowVersion(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.GetFlowVersionInput{
		// FlowIdentifier: *string, // Required
		// FlowVersion: *string, // Required
	}

	if len(_bedrockagentFlowIdentifier) > 0 {
		input.FlowIdentifier = aws.String(_bedrockagentFlowIdentifier)
	}
	if len(_bedrockagentFlowVersion) > 0 {
		input.FlowVersion = aws.String(_bedrockagentFlowVersion)
	}

	if resp, err := client.GetFlowVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a data ingestion job. Data sources are ingested into
// your knowledge base so that Large Language Models (LLMs) can use your data.
func bedrockagent_GetIngestionJob(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.GetIngestionJobInput{
		// DataSourceId: *string, // Required
		// IngestionJobId: *string, // Required
		// KnowledgeBaseId: *string, // Required
	}

	if len(_bedrockagentDataSourceId) > 0 {
		input.DataSourceId = aws.String(_bedrockagentDataSourceId)
	}
	if len(_bedrockagentIngestionJobId) > 0 {
		input.IngestionJobId = aws.String(_bedrockagentIngestionJobId)
	}
	if len(_bedrockagentKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_bedrockagentKnowledgeBaseId)
	}

	if resp, err := client.GetIngestionJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a knowledge base.
func bedrockagent_GetKnowledgeBase(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.GetKnowledgeBaseInput{
		// KnowledgeBaseId: *string, // Required
	}

	if len(_bedrockagentKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_bedrockagentKnowledgeBaseId)
	}

	if resp, err := client.GetKnowledgeBase(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves specific documents from a data source that is connected to a
// knowledge base. For more information, see [Ingest changes directly into a knowledge base]in the Amazon Bedrock User Guide.
//
// [Ingest changes directly into a knowledge base]: https://docs.aws.amazon.com/bedrock/latest/userguide/kb-direct-ingestion.html
func bedrockagent_GetKnowledgeBaseDocuments(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.GetKnowledgeBaseDocumentsInput{
		// DataSourceId: *string, // Required
		// DocumentIdentifiers: []types.DocumentIdentifier, // Required
		// KnowledgeBaseId: *string, // Required
	}

	if len(_bedrockagentDataSourceId) > 0 {
		input.DataSourceId = aws.String(_bedrockagentDataSourceId)
	}
	if len(_bedrockagentDocumentIdentifiers) > 0 {
		if err := assignInputField(input, "DocumentIdentifiers", _bedrockagentDocumentIdentifiers); err != nil {
			log.Errorf("invalid --document-identifiers: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_bedrockagentKnowledgeBaseId)
	}

	if resp, err := client.GetKnowledgeBaseDocuments(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the working draft ( DRAFT version) of a prompt or a
// version of it, depending on whether you include the promptVersion field or not.
// For more information, see [View information about prompts using Prompt management]and [View information about a version of your prompt] in the Amazon Bedrock User Guide.
//
// [View information about a version of your prompt]: https://docs.aws.amazon.com/bedrock/latest/userguide/prompt-management-deploy.html#prompt-management-versions-view.html
// [View information about prompts using Prompt management]: https://docs.aws.amazon.com/bedrock/latest/userguide/prompt-management-manage.html#prompt-management-view.html
func bedrockagent_GetPrompt(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.GetPromptInput{
		// PromptIdentifier: *string, // Required
	}

	if len(_bedrockagentPromptIdentifier) > 0 {
		input.PromptIdentifier = aws.String(_bedrockagentPromptIdentifier)
	}
	if len(_bedrockagentPromptVersion) > 0 {
		input.PromptVersion = aws.String(_bedrockagentPromptVersion)
	}

	if resp, err := client.GetPrompt(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Ingests documents directly into the knowledge base that is connected to the
// data source. The dataSourceType specified in the content for each document must
// match the type of the data source that you specify in the header. For more
// information, see [Ingest changes directly into a knowledge base]in the Amazon Bedrock User Guide.
//
// [Ingest changes directly into a knowledge base]: https://docs.aws.amazon.com/bedrock/latest/userguide/kb-direct-ingestion.html
func bedrockagent_IngestKnowledgeBaseDocuments(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.IngestKnowledgeBaseDocumentsInput{
		// DataSourceId: *string, // Required
		// Documents: []types.KnowledgeBaseDocument, // Required
		// KnowledgeBaseId: *string, // Required
	}

	if len(_bedrockagentDataSourceId) > 0 {
		input.DataSourceId = aws.String(_bedrockagentDataSourceId)
	}
	if len(_bedrockagentDocuments) > 0 {
		if err := assignInputField(input, "Documents", _bedrockagentDocuments); err != nil {
			log.Errorf("invalid --documents: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_bedrockagentKnowledgeBaseId)
	}
	if len(_bedrockagentClientToken) > 0 {
		input.ClientToken = aws.String(_bedrockagentClientToken)
	}

	if resp, err := client.IngestKnowledgeBaseDocuments(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the action groups for an agent and information about each one.
func bedrockagent_ListAgentActionGroups(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.ListAgentActionGroupsInput{
		// AgentId: *string, // Required
		// AgentVersion: *string, // Required
	}

	if len(_bedrockagentAgentId) > 0 {
		input.AgentId = aws.String(_bedrockagentAgentId)
	}
	if len(_bedrockagentAgentVersion) > 0 {
		input.AgentVersion = aws.String(_bedrockagentAgentVersion)
	}
	if len(_bedrockagentMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bedrockagentMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentNextToken) > 0 {
		input.NextToken = aws.String(_bedrockagentNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAgentActionGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrockagent.ListAgentActionGroupsOutput
	p := bedrockagent.NewListAgentActionGroupsPaginator(client, input)
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

// Lists the aliases of an agent and information about each one.
func bedrockagent_ListAgentAliases(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.ListAgentAliasesInput{
		// AgentId: *string, // Required
	}

	if len(_bedrockagentAgentId) > 0 {
		input.AgentId = aws.String(_bedrockagentAgentId)
	}
	if len(_bedrockagentMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bedrockagentMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentNextToken) > 0 {
		input.NextToken = aws.String(_bedrockagentNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAgentAliases(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrockagent.ListAgentAliasesOutput
	p := bedrockagent.NewListAgentAliasesPaginator(client, input)
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

// Retrieve a list of an agent's collaborators.
func bedrockagent_ListAgentCollaborators(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.ListAgentCollaboratorsInput{
		// AgentId: *string, // Required
		// AgentVersion: *string, // Required
	}

	if len(_bedrockagentAgentId) > 0 {
		input.AgentId = aws.String(_bedrockagentAgentId)
	}
	if len(_bedrockagentAgentVersion) > 0 {
		input.AgentVersion = aws.String(_bedrockagentAgentVersion)
	}
	if len(_bedrockagentMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bedrockagentMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentNextToken) > 0 {
		input.NextToken = aws.String(_bedrockagentNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAgentCollaborators(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrockagent.ListAgentCollaboratorsOutput
	p := bedrockagent.NewListAgentCollaboratorsPaginator(client, input)
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

// Lists knowledge bases associated with an agent and information about each one.
func bedrockagent_ListAgentKnowledgeBases(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.ListAgentKnowledgeBasesInput{
		// AgentId: *string, // Required
		// AgentVersion: *string, // Required
	}

	if len(_bedrockagentAgentId) > 0 {
		input.AgentId = aws.String(_bedrockagentAgentId)
	}
	if len(_bedrockagentAgentVersion) > 0 {
		input.AgentVersion = aws.String(_bedrockagentAgentVersion)
	}
	if len(_bedrockagentMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bedrockagentMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentNextToken) > 0 {
		input.NextToken = aws.String(_bedrockagentNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAgentKnowledgeBases(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrockagent.ListAgentKnowledgeBasesOutput
	p := bedrockagent.NewListAgentKnowledgeBasesPaginator(client, input)
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

// Lists the versions of an agent and information about each version.
func bedrockagent_ListAgentVersions(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.ListAgentVersionsInput{
		// AgentId: *string, // Required
	}

	if len(_bedrockagentAgentId) > 0 {
		input.AgentId = aws.String(_bedrockagentAgentId)
	}
	if len(_bedrockagentMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bedrockagentMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentNextToken) > 0 {
		input.NextToken = aws.String(_bedrockagentNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAgentVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrockagent.ListAgentVersionsOutput
	p := bedrockagent.NewListAgentVersionsPaginator(client, input)
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

// Lists the agents belonging to an account and information about each agent.
func bedrockagent_ListAgents(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.ListAgentsInput{}

	if len(_bedrockagentMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bedrockagentMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentNextToken) > 0 {
		input.NextToken = aws.String(_bedrockagentNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAgents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrockagent.ListAgentsOutput
	p := bedrockagent.NewListAgentsPaginator(client, input)
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

// Lists the data sources in a knowledge base and information about each one.
func bedrockagent_ListDataSources(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.ListDataSourcesInput{
		// KnowledgeBaseId: *string, // Required
	}

	if len(_bedrockagentKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_bedrockagentKnowledgeBaseId)
	}
	if len(_bedrockagentMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bedrockagentMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentNextToken) > 0 {
		input.NextToken = aws.String(_bedrockagentNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDataSources(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrockagent.ListDataSourcesOutput
	p := bedrockagent.NewListDataSourcesPaginator(client, input)
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

// Returns a list of aliases for a flow.
func bedrockagent_ListFlowAliases(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.ListFlowAliasesInput{
		// FlowIdentifier: *string, // Required
	}

	if len(_bedrockagentFlowIdentifier) > 0 {
		input.FlowIdentifier = aws.String(_bedrockagentFlowIdentifier)
	}
	if len(_bedrockagentMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bedrockagentMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentNextToken) > 0 {
		input.NextToken = aws.String(_bedrockagentNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListFlowAliases(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrockagent.ListFlowAliasesOutput
	p := bedrockagent.NewListFlowAliasesPaginator(client, input)
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

// Returns a list of information about each flow. For more information, see [Deploy a flow in Amazon Bedrock] in
// the Amazon Bedrock User Guide.
//
// [Deploy a flow in Amazon Bedrock]: https://docs.aws.amazon.com/bedrock/latest/userguide/flows-deploy.html
func bedrockagent_ListFlowVersions(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.ListFlowVersionsInput{
		// FlowIdentifier: *string, // Required
	}

	if len(_bedrockagentFlowIdentifier) > 0 {
		input.FlowIdentifier = aws.String(_bedrockagentFlowIdentifier)
	}
	if len(_bedrockagentMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bedrockagentMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentNextToken) > 0 {
		input.NextToken = aws.String(_bedrockagentNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListFlowVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrockagent.ListFlowVersionsOutput
	p := bedrockagent.NewListFlowVersionsPaginator(client, input)
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

// Returns a list of flows and information about each flow. For more information,
// see [Manage a flow in Amazon Bedrock]in the Amazon Bedrock User Guide.
//
// [Manage a flow in Amazon Bedrock]: https://docs.aws.amazon.com/bedrock/latest/userguide/flows-manage.html
func bedrockagent_ListFlows(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.ListFlowsInput{}

	if len(_bedrockagentMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bedrockagentMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentNextToken) > 0 {
		input.NextToken = aws.String(_bedrockagentNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListFlows(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrockagent.ListFlowsOutput
	p := bedrockagent.NewListFlowsPaginator(client, input)
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

// Lists the data ingestion jobs for a data source. The list also includes
// information about each job.
func bedrockagent_ListIngestionJobs(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.ListIngestionJobsInput{
		// DataSourceId: *string, // Required
		// KnowledgeBaseId: *string, // Required
	}

	if len(_bedrockagentDataSourceId) > 0 {
		input.DataSourceId = aws.String(_bedrockagentDataSourceId)
	}
	if len(_bedrockagentKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_bedrockagentKnowledgeBaseId)
	}
	if len(_bedrockagentFilters) > 0 {
		if err := assignInputField(input, "Filters", _bedrockagentFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bedrockagentMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentNextToken) > 0 {
		input.NextToken = aws.String(_bedrockagentNextToken)
	}
	if len(_bedrockagentSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _bedrockagentSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListIngestionJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrockagent.ListIngestionJobsOutput
	p := bedrockagent.NewListIngestionJobsPaginator(client, input)
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

// Retrieves all the documents contained in a data source that is connected to a
// knowledge base. For more information, see [Ingest changes directly into a knowledge base]in the Amazon Bedrock User Guide.
//
// [Ingest changes directly into a knowledge base]: https://docs.aws.amazon.com/bedrock/latest/userguide/kb-direct-ingestion.html
func bedrockagent_ListKnowledgeBaseDocuments(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.ListKnowledgeBaseDocumentsInput{
		// DataSourceId: *string, // Required
		// KnowledgeBaseId: *string, // Required
	}

	if len(_bedrockagentDataSourceId) > 0 {
		input.DataSourceId = aws.String(_bedrockagentDataSourceId)
	}
	if len(_bedrockagentKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_bedrockagentKnowledgeBaseId)
	}
	if len(_bedrockagentMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bedrockagentMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentNextToken) > 0 {
		input.NextToken = aws.String(_bedrockagentNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListKnowledgeBaseDocuments(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrockagent.ListKnowledgeBaseDocumentsOutput
	p := bedrockagent.NewListKnowledgeBaseDocumentsPaginator(client, input)
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

// Lists the knowledge bases in an account. The list also includesinformation
// about each knowledge base.
func bedrockagent_ListKnowledgeBases(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.ListKnowledgeBasesInput{}

	if len(_bedrockagentMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bedrockagentMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentNextToken) > 0 {
		input.NextToken = aws.String(_bedrockagentNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListKnowledgeBases(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrockagent.ListKnowledgeBasesOutput
	p := bedrockagent.NewListKnowledgeBasesPaginator(client, input)
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

// Returns either information about the working draft ( DRAFT version) of each
// prompt in an account, or information about of all versions of a prompt,
// depending on whether you include the promptIdentifier field or not. For more
// information, see [View information about prompts using Prompt management]in the Amazon Bedrock User Guide.
//
// [View information about prompts using Prompt management]: https://docs.aws.amazon.com/bedrock/latest/userguide/prompt-management-manage.html#prompt-management-view.html
func bedrockagent_ListPrompts(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.ListPromptsInput{}

	if len(_bedrockagentMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bedrockagentMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentNextToken) > 0 {
		input.NextToken = aws.String(_bedrockagentNextToken)
	}
	if len(_bedrockagentPromptIdentifier) > 0 {
		input.PromptIdentifier = aws.String(_bedrockagentPromptIdentifier)
	}

	if disablePaginator() {
		if resp, err := client.ListPrompts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrockagent.ListPromptsOutput
	p := bedrockagent.NewListPromptsPaginator(client, input)
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
func bedrockagent_ListTagsForResource(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_bedrockagentResourceArn) > 0 {
		input.ResourceArn = aws.String(_bedrockagentResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a DRAFT version of the agent that can be used for internal testing.
func bedrockagent_PrepareAgent(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.PrepareAgentInput{
		// AgentId: *string, // Required
	}

	if len(_bedrockagentAgentId) > 0 {
		input.AgentId = aws.String(_bedrockagentAgentId)
	}

	if resp, err := client.PrepareAgent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Prepares the DRAFT version of a flow so that it can be invoked. For more
// information, see [Test a flow in Amazon Bedrock]in the Amazon Bedrock User Guide.
//
// [Test a flow in Amazon Bedrock]: https://docs.aws.amazon.com/bedrock/latest/userguide/flows-test.html
func bedrockagent_PrepareFlow(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.PrepareFlowInput{
		// FlowIdentifier: *string, // Required
	}

	if len(_bedrockagentFlowIdentifier) > 0 {
		input.FlowIdentifier = aws.String(_bedrockagentFlowIdentifier)
	}

	if resp, err := client.PrepareFlow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Begins a data ingestion job. Data sources are ingested into your knowledge base
// so that Large Language Models (LLMs) can use your data.
func bedrockagent_StartIngestionJob(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.StartIngestionJobInput{
		// DataSourceId: *string, // Required
		// KnowledgeBaseId: *string, // Required
	}

	if len(_bedrockagentDataSourceId) > 0 {
		input.DataSourceId = aws.String(_bedrockagentDataSourceId)
	}
	if len(_bedrockagentKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_bedrockagentKnowledgeBaseId)
	}
	if len(_bedrockagentClientToken) > 0 {
		input.ClientToken = aws.String(_bedrockagentClientToken)
	}
	if len(_bedrockagentDescription) > 0 {
		input.Description = aws.String(_bedrockagentDescription)
	}

	if resp, err := client.StartIngestionJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a currently running data ingestion job. You can send a StartIngestionJob
// request again to ingest the rest of your data when you are ready.
func bedrockagent_StopIngestionJob(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.StopIngestionJobInput{
		// DataSourceId: *string, // Required
		// IngestionJobId: *string, // Required
		// KnowledgeBaseId: *string, // Required
	}

	if len(_bedrockagentDataSourceId) > 0 {
		input.DataSourceId = aws.String(_bedrockagentDataSourceId)
	}
	if len(_bedrockagentIngestionJobId) > 0 {
		input.IngestionJobId = aws.String(_bedrockagentIngestionJobId)
	}
	if len(_bedrockagentKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_bedrockagentKnowledgeBaseId)
	}

	if resp, err := client.StopIngestionJob(context.TODO(), input); err != nil {
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
func bedrockagent_TagResource(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_bedrockagentResourceArn) > 0 {
		input.ResourceArn = aws.String(_bedrockagentResourceArn)
	}
	if len(_bedrockagentTags) > 0 {
		if err := assignInputField(input, "Tags", _bedrockagentTags); err != nil {
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
func bedrockagent_UntagResource(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_bedrockagentResourceArn) > 0 {
		input.ResourceArn = aws.String(_bedrockagentResourceArn)
	}
	if len(_bedrockagentTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _bedrockagentTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration of an agent.
func bedrockagent_UpdateAgent(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.UpdateAgentInput{
		// AgentId: *string, // Required
		// AgentName: *string, // Required
		// AgentResourceRoleArn: *string, // Required
		// FoundationModel: *string, // Required
	}

	if len(_bedrockagentAgentId) > 0 {
		input.AgentId = aws.String(_bedrockagentAgentId)
	}
	if len(_bedrockagentAgentName) > 0 {
		input.AgentName = aws.String(_bedrockagentAgentName)
	}
	if len(_bedrockagentAgentResourceRoleArn) > 0 {
		input.AgentResourceRoleArn = aws.String(_bedrockagentAgentResourceRoleArn)
	}
	if len(_bedrockagentFoundationModel) > 0 {
		input.FoundationModel = aws.String(_bedrockagentFoundationModel)
	}
	if len(_bedrockagentAgentCollaboration) > 0 {
		if err := assignInputField(input, "AgentCollaboration", _bedrockagentAgentCollaboration); err != nil {
			log.Errorf("invalid --agent-collaboration: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentCustomOrchestration) > 0 {
		if err := assignInputField(input, "CustomOrchestration", _bedrockagentCustomOrchestration); err != nil {
			log.Errorf("invalid --custom-orchestration: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentCustomerEncryptionKeyArn) > 0 {
		input.CustomerEncryptionKeyArn = aws.String(_bedrockagentCustomerEncryptionKeyArn)
	}
	if len(_bedrockagentDescription) > 0 {
		input.Description = aws.String(_bedrockagentDescription)
	}
	if len(_bedrockagentGuardrailConfiguration) > 0 {
		if err := assignInputField(input, "GuardrailConfiguration", _bedrockagentGuardrailConfiguration); err != nil {
			log.Errorf("invalid --guardrail-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentIdleSessionTTLInSeconds) > 0 {
		if err := assignInputField(input, "IdleSessionTTLInSeconds", _bedrockagentIdleSessionTTLInSeconds); err != nil {
			log.Errorf("invalid --idle-session-ttlin-seconds: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentInstruction) > 0 {
		input.Instruction = aws.String(_bedrockagentInstruction)
	}
	if len(_bedrockagentMemoryConfiguration) > 0 {
		if err := assignInputField(input, "MemoryConfiguration", _bedrockagentMemoryConfiguration); err != nil {
			log.Errorf("invalid --memory-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentOrchestrationType) > 0 {
		if err := assignInputField(input, "OrchestrationType", _bedrockagentOrchestrationType); err != nil {
			log.Errorf("invalid --orchestration-type: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentPromptOverrideConfiguration) > 0 {
		if err := assignInputField(input, "PromptOverrideConfiguration", _bedrockagentPromptOverrideConfiguration); err != nil {
			log.Errorf("invalid --prompt-override-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateAgent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration for an action group for an agent.
func bedrockagent_UpdateAgentActionGroup(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.UpdateAgentActionGroupInput{
		// ActionGroupId: *string, // Required
		// ActionGroupName: *string, // Required
		// AgentId: *string, // Required
		// AgentVersion: *string, // Required
	}

	if len(_bedrockagentActionGroupId) > 0 {
		input.ActionGroupId = aws.String(_bedrockagentActionGroupId)
	}
	if len(_bedrockagentActionGroupName) > 0 {
		input.ActionGroupName = aws.String(_bedrockagentActionGroupName)
	}
	if len(_bedrockagentAgentId) > 0 {
		input.AgentId = aws.String(_bedrockagentAgentId)
	}
	if len(_bedrockagentAgentVersion) > 0 {
		input.AgentVersion = aws.String(_bedrockagentAgentVersion)
	}
	if len(_bedrockagentActionGroupExecutor) > 0 {
		if err := assignInputField(input, "ActionGroupExecutor", _bedrockagentActionGroupExecutor); err != nil {
			log.Errorf("invalid --action-group-executor: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentActionGroupState) > 0 {
		if err := assignInputField(input, "ActionGroupState", _bedrockagentActionGroupState); err != nil {
			log.Errorf("invalid --action-group-state: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentApiSchema) > 0 {
		if err := assignInputField(input, "ApiSchema", _bedrockagentApiSchema); err != nil {
			log.Errorf("invalid --api-schema: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentDescription) > 0 {
		input.Description = aws.String(_bedrockagentDescription)
	}
	if len(_bedrockagentFunctionSchema) > 0 {
		if err := assignInputField(input, "FunctionSchema", _bedrockagentFunctionSchema); err != nil {
			log.Errorf("invalid --function-schema: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentParentActionGroupSignature) > 0 {
		if err := assignInputField(input, "ParentActionGroupSignature", _bedrockagentParentActionGroupSignature); err != nil {
			log.Errorf("invalid --parent-action-group-signature: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentParentActionGroupSignatureParams) > 0 {
		if err := assignInputField(input, "ParentActionGroupSignatureParams", _bedrockagentParentActionGroupSignatureParams); err != nil {
			log.Errorf("invalid --parent-action-group-signature-params: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateAgentActionGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates configurations for an alias of an agent.
func bedrockagent_UpdateAgentAlias(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.UpdateAgentAliasInput{
		// AgentAliasId: *string, // Required
		// AgentAliasName: *string, // Required
		// AgentId: *string, // Required
	}

	if len(_bedrockagentAgentAliasId) > 0 {
		input.AgentAliasId = aws.String(_bedrockagentAgentAliasId)
	}
	if len(_bedrockagentAgentAliasName) > 0 {
		input.AgentAliasName = aws.String(_bedrockagentAgentAliasName)
	}
	if len(_bedrockagentAgentId) > 0 {
		input.AgentId = aws.String(_bedrockagentAgentId)
	}
	if len(_bedrockagentAliasInvocationState) > 0 {
		if err := assignInputField(input, "AliasInvocationState", _bedrockagentAliasInvocationState); err != nil {
			log.Errorf("invalid --alias-invocation-state: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentDescription) > 0 {
		input.Description = aws.String(_bedrockagentDescription)
	}
	if len(_bedrockagentRoutingConfiguration) > 0 {
		if err := assignInputField(input, "RoutingConfiguration", _bedrockagentRoutingConfiguration); err != nil {
			log.Errorf("invalid --routing-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateAgentAlias(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an agent's collaborator.
func bedrockagent_UpdateAgentCollaborator(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.UpdateAgentCollaboratorInput{
		// AgentDescriptor: *types.AgentDescriptor, // Required
		// AgentId: *string, // Required
		// AgentVersion: *string, // Required
		// CollaborationInstruction: *string, // Required
		// CollaboratorId: *string, // Required
		// CollaboratorName: *string, // Required
	}

	if len(_bedrockagentAgentDescriptor) > 0 {
		if err := assignInputField(input, "AgentDescriptor", _bedrockagentAgentDescriptor); err != nil {
			log.Errorf("invalid --agent-descriptor: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentAgentId) > 0 {
		input.AgentId = aws.String(_bedrockagentAgentId)
	}
	if len(_bedrockagentAgentVersion) > 0 {
		input.AgentVersion = aws.String(_bedrockagentAgentVersion)
	}
	if len(_bedrockagentCollaborationInstruction) > 0 {
		input.CollaborationInstruction = aws.String(_bedrockagentCollaborationInstruction)
	}
	if len(_bedrockagentCollaboratorId) > 0 {
		input.CollaboratorId = aws.String(_bedrockagentCollaboratorId)
	}
	if len(_bedrockagentCollaboratorName) > 0 {
		input.CollaboratorName = aws.String(_bedrockagentCollaboratorName)
	}
	if len(_bedrockagentRelayConversationHistory) > 0 {
		if err := assignInputField(input, "RelayConversationHistory", _bedrockagentRelayConversationHistory); err != nil {
			log.Errorf("invalid --relay-conversation-history: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateAgentCollaborator(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration for a knowledge base that has been associated with an
// agent.
func bedrockagent_UpdateAgentKnowledgeBase(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.UpdateAgentKnowledgeBaseInput{
		// AgentId: *string, // Required
		// AgentVersion: *string, // Required
		// KnowledgeBaseId: *string, // Required
	}

	if len(_bedrockagentAgentId) > 0 {
		input.AgentId = aws.String(_bedrockagentAgentId)
	}
	if len(_bedrockagentAgentVersion) > 0 {
		input.AgentVersion = aws.String(_bedrockagentAgentVersion)
	}
	if len(_bedrockagentKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_bedrockagentKnowledgeBaseId)
	}
	if len(_bedrockagentDescription) > 0 {
		input.Description = aws.String(_bedrockagentDescription)
	}
	if len(_bedrockagentKnowledgeBaseState) > 0 {
		if err := assignInputField(input, "KnowledgeBaseState", _bedrockagentKnowledgeBaseState); err != nil {
			log.Errorf("invalid --knowledge-base-state: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateAgentKnowledgeBase(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configurations for a data source connector.
// You can't change the chunkingConfiguration after you create the data source
// connector. Specify the existing chunkingConfiguration .
func bedrockagent_UpdateDataSource(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.UpdateDataSourceInput{
		// DataSourceConfiguration: *types.DataSourceConfiguration, // Required
		// DataSourceId: *string, // Required
		// KnowledgeBaseId: *string, // Required
		// Name: *string, // Required
	}

	if len(_bedrockagentDataSourceConfiguration) > 0 {
		if err := assignInputField(input, "DataSourceConfiguration", _bedrockagentDataSourceConfiguration); err != nil {
			log.Errorf("invalid --data-source-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentDataSourceId) > 0 {
		input.DataSourceId = aws.String(_bedrockagentDataSourceId)
	}
	if len(_bedrockagentKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_bedrockagentKnowledgeBaseId)
	}
	if len(_bedrockagentName) > 0 {
		input.Name = aws.String(_bedrockagentName)
	}
	if len(_bedrockagentDataDeletionPolicy) > 0 {
		if err := assignInputField(input, "DataDeletionPolicy", _bedrockagentDataDeletionPolicy); err != nil {
			log.Errorf("invalid --data-deletion-policy: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentDescription) > 0 {
		input.Description = aws.String(_bedrockagentDescription)
	}
	if len(_bedrockagentServerSideEncryptionConfiguration) > 0 {
		if err := assignInputField(input, "ServerSideEncryptionConfiguration", _bedrockagentServerSideEncryptionConfiguration); err != nil {
			log.Errorf("invalid --server-side-encryption-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentVectorIngestionConfiguration) > 0 {
		if err := assignInputField(input, "VectorIngestionConfiguration", _bedrockagentVectorIngestionConfiguration); err != nil {
			log.Errorf("invalid --vector-ingestion-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateDataSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies a flow. Include both fields that you want to keep and fields that you
// want to change. For more information, see [How it works]and [Create a flow in Amazon Bedrock] in the Amazon Bedrock User Guide.
//
// [How it works]: https://docs.aws.amazon.com/bedrock/latest/userguide/flows-how-it-works.html
// [Create a flow in Amazon Bedrock]: https://docs.aws.amazon.com/bedrock/latest/userguide/flows-create.html
func bedrockagent_UpdateFlow(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.UpdateFlowInput{
		// ExecutionRoleArn: *string, // Required
		// FlowIdentifier: *string, // Required
		// Name: *string, // Required
	}

	if len(_bedrockagentExecutionRoleArn) > 0 {
		input.ExecutionRoleArn = aws.String(_bedrockagentExecutionRoleArn)
	}
	if len(_bedrockagentFlowIdentifier) > 0 {
		input.FlowIdentifier = aws.String(_bedrockagentFlowIdentifier)
	}
	if len(_bedrockagentName) > 0 {
		input.Name = aws.String(_bedrockagentName)
	}
	if len(_bedrockagentCustomerEncryptionKeyArn) > 0 {
		input.CustomerEncryptionKeyArn = aws.String(_bedrockagentCustomerEncryptionKeyArn)
	}
	if len(_bedrockagentDefinition) > 0 {
		if err := assignInputField(input, "Definition", _bedrockagentDefinition); err != nil {
			log.Errorf("invalid --definition: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentDescription) > 0 {
		input.Description = aws.String(_bedrockagentDescription)
	}

	if resp, err := client.UpdateFlow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the alias of a flow. Include both fields that you want to keep and
// ones that you want to change. For more information, see [Deploy a flow in Amazon Bedrock]in the Amazon Bedrock
// User Guide.
//
// [Deploy a flow in Amazon Bedrock]: https://docs.aws.amazon.com/bedrock/latest/userguide/flows-deploy.html
func bedrockagent_UpdateFlowAlias(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.UpdateFlowAliasInput{
		// AliasIdentifier: *string, // Required
		// FlowIdentifier: *string, // Required
		// Name: *string, // Required
		// RoutingConfiguration: []types.FlowAliasRoutingConfigurationListItem, // Required
	}

	if len(_bedrockagentAliasIdentifier) > 0 {
		input.AliasIdentifier = aws.String(_bedrockagentAliasIdentifier)
	}
	if len(_bedrockagentFlowIdentifier) > 0 {
		input.FlowIdentifier = aws.String(_bedrockagentFlowIdentifier)
	}
	if len(_bedrockagentName) > 0 {
		input.Name = aws.String(_bedrockagentName)
	}
	if len(_bedrockagentRoutingConfiguration) > 0 {
		if err := assignInputField(input, "RoutingConfiguration", _bedrockagentRoutingConfiguration); err != nil {
			log.Errorf("invalid --routing-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentConcurrencyConfiguration) > 0 {
		if err := assignInputField(input, "ConcurrencyConfiguration", _bedrockagentConcurrencyConfiguration); err != nil {
			log.Errorf("invalid --concurrency-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentDescription) > 0 {
		input.Description = aws.String(_bedrockagentDescription)
	}

	if resp, err := client.UpdateFlowAlias(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration of a knowledge base with the fields that you specify.
// Because all fields will be overwritten, you must include the same values for
// fields that you want to keep the same.
//
// You can change the following fields:
//
// - name
//
// - description
//
// - roleArn
//
// You can't change the knowledgeBaseConfiguration or storageConfiguration fields,
// so you must specify the same configurations as when you created the knowledge
// base. You can send a [GetKnowledgeBase]request and copy the same configurations.
//
// [GetKnowledgeBase]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent_GetKnowledgeBase.html
func bedrockagent_UpdateKnowledgeBase(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.UpdateKnowledgeBaseInput{
		// KnowledgeBaseConfiguration: *types.KnowledgeBaseConfiguration, // Required
		// KnowledgeBaseId: *string, // Required
		// Name: *string, // Required
		// RoleArn: *string, // Required
	}

	if len(_bedrockagentKnowledgeBaseConfiguration) > 0 {
		if err := assignInputField(input, "KnowledgeBaseConfiguration", _bedrockagentKnowledgeBaseConfiguration); err != nil {
			log.Errorf("invalid --knowledge-base-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentKnowledgeBaseId) > 0 {
		input.KnowledgeBaseId = aws.String(_bedrockagentKnowledgeBaseId)
	}
	if len(_bedrockagentName) > 0 {
		input.Name = aws.String(_bedrockagentName)
	}
	if len(_bedrockagentRoleArn) > 0 {
		input.RoleArn = aws.String(_bedrockagentRoleArn)
	}
	if len(_bedrockagentDescription) > 0 {
		input.Description = aws.String(_bedrockagentDescription)
	}
	if len(_bedrockagentStorageConfiguration) > 0 {
		if err := assignInputField(input, "StorageConfiguration", _bedrockagentStorageConfiguration); err != nil {
			log.Errorf("invalid --storage-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateKnowledgeBase(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies a prompt in your prompt library. Include both fields that you want to
// keep and fields that you want to replace. For more information, see [Prompt management in Amazon Bedrock]and [Edit prompts in your prompt library] in the
// Amazon Bedrock User Guide.
//
// [Edit prompts in your prompt library]: https://docs.aws.amazon.com/bedrock/latest/userguide/prompt-management-manage.html#prompt-management-edit
// [Prompt management in Amazon Bedrock]: https://docs.aws.amazon.com/bedrock/latest/userguide/prompt-management.html
func bedrockagent_UpdatePrompt(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.UpdatePromptInput{
		// Name: *string, // Required
		// PromptIdentifier: *string, // Required
	}

	if len(_bedrockagentName) > 0 {
		input.Name = aws.String(_bedrockagentName)
	}
	if len(_bedrockagentPromptIdentifier) > 0 {
		input.PromptIdentifier = aws.String(_bedrockagentPromptIdentifier)
	}
	if len(_bedrockagentCustomerEncryptionKeyArn) > 0 {
		input.CustomerEncryptionKeyArn = aws.String(_bedrockagentCustomerEncryptionKeyArn)
	}
	if len(_bedrockagentDefaultVariant) > 0 {
		input.DefaultVariant = aws.String(_bedrockagentDefaultVariant)
	}
	if len(_bedrockagentDescription) > 0 {
		input.Description = aws.String(_bedrockagentDescription)
	}
	if len(_bedrockagentVariants) > 0 {
		if err := assignInputField(input, "Variants", _bedrockagentVariants); err != nil {
			log.Errorf("invalid --variants: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdatePrompt(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Validates the definition of a flow.
func bedrockagent_ValidateFlowDefinition(cfg aws.Config, client *bedrockagent.Client) {
	input := &bedrockagent.ValidateFlowDefinitionInput{
		// Definition: *types.FlowDefinition, // Required
	}

	if len(_bedrockagentDefinition) > 0 {
		if err := assignInputField(input, "Definition", _bedrockagentDefinition); err != nil {
			log.Errorf("invalid --definition: %s", err.Error())
			return
		}
	}

	if resp, err := client.ValidateFlowDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_bedrockagentCmd)
	_bedrockagentCmd.Flags().SortFlags = false

	_bedrockagentCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_bedrockagentCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_bedrockagentCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_bedrockagentCmd.Flags().StringVarP(&_bedrockagentActionGroupExecutor, "action-group-executor", "", "", "Action Group Executor")
	_bedrockagentCmd.Flags().StringVarP(&_bedrockagentActionGroupId, "action-group-id", "", "", "Action Group ID")
	_bedrockagentCmd.Flags().StringVarP(&_bedrockagentActionGroupName, "action-group-name", "", "", "Action Group Name")
	_bedrockagentCmd.Flags().StringVarP(&_bedrockagentActionGroupState, "action-group-state", "", "", "Action Group State")
	_bedrockagentCmd.Flags().StringVarP(&_bedrockagentAgentAliasId, "agent-alias-id", "", "", "Agent Alias ID")
	_bedrockagentCmd.Flags().StringVarP(&_bedrockagentAgentAliasName, "agent-alias-name", "", "", "Agent Alias Name")
	_bedrockagentCmd.Flags().StringVarP(&_bedrockagentAgentCollaboration, "agent-collaboration", "", "", "Agent Collaboration")
	_bedrockagentCmd.Flags().StringVarP(&_bedrockagentAgentDescriptor, "agent-descriptor", "", "", "Agent Descriptor")
	_bedrockagentCmd.Flags().StringVarP(&_bedrockagentAgentId, "agent-id", "", "", "Agent ID")
	_bedrockagentCmd.Flags().StringVarP(&_bedrockagentAgentName, "agent-name", "", "", "Agent Name")
	_bedrockagentCmd.Flags().StringVarP(&_bedrockagentAgentResourceRoleArn, "agent-resource-role-arn", "", "", "Agent Resource Role ARN")
	_bedrockagentCmd.Flags().StringVarP(&_bedrockagentAgentVersion, "agent-version", "", "", "Agent Version")
	_bedrockagentCmd.Flags().StringVarP(&_bedrockagentAliasIdentifier, "alias-identifier", "", "", "Alias Identifier")
	_bedrockagentCmd.Flags().StringVarP(&_bedrockagentAliasInvocationState, "alias-invocation-state", "", "", "Alias Invocation State")
	_bedrockagentCmd.Flags().StringVarP(&_bedrockagentApiSchema, "api-schema", "", "", "API Schema")
	_bedrockagentCmd.Flags().StringVarP(&_bedrockagentClientToken, "client-token", "", "", "Client Token")
	_bedrockagentCmd.Flags().StringVarP(&_bedrockagentCollaborationInstruction, "collaboration-instruction", "", "", "Collaboration Instruction")
	_bedrockagentCmd.Flags().StringVarP(&_bedrockagentCollaboratorId, "collaborator-id", "", "", "Collaborator ID")
	_bedrockagentCmd.Flags().StringVarP(&_bedrockagentCollaboratorName, "collaborator-name", "", "", "Collaborator Name")
	_bedrockagentCmd.Flags().StringVarP(&_bedrockagentConcurrencyConfiguration, "concurrency-configuration", "", "", "Concurrency Configuration")
	_bedrockagentCmd.Flags().StringVarP(&_bedrockagentCustomOrchestration, "custom-orchestration", "", "", "Custom Orchestration")
	_bedrockagentCmd.Flags().StringVarP(&_bedrockagentCustomerEncryptionKeyArn, "customer-encryption-key-arn", "", "", "Customer Encryption Key ARN")
	_bedrockagentCmd.Flags().StringVarP(&_bedrockagentDataDeletionPolicy, "data-deletion-policy", "", "", "Data Deletion Policy")
	_bedrockagentCmd.Flags().StringVarP(&_bedrockagentDataSourceConfiguration, "data-source-configuration", "", "", "Data Source Configuration")
	_bedrockagentCmd.Flags().StringVarP(&_bedrockagentDataSourceId, "data-source-id", "", "", "Data Source ID")
	_bedrockagentCmd.Flags().StringVarP(&_bedrockagentDefaultVariant, "default-variant", "", "", "Default Variant")
	_bedrockagentCmd.Flags().StringVarP(&_bedrockagentDefinition, "definition", "", "", "Definition")
	_bedrockagentCmd.Flags().StringVarP(&_bedrockagentDescription, "description", "", "", "Description")
	_bedrockagentCmd.Flags().StringVarP(&_bedrockagentDocumentIdentifiers, "document-identifiers", "", "", "Document Identifiers")
	_bedrockagentCmd.Flags().StringVarP(&_bedrockagentDocuments, "documents", "", "", "Documents")
	_bedrockagentCmd.Flags().StringVarP(&_bedrockagentExecutionRoleArn, "execution-role-arn", "", "", "Execution Role ARN")
	_bedrockagentCmd.Flags().StringVarP(&_bedrockagentFilters, "filters", "", "", "Filters")
	_bedrockagentCmd.Flags().StringVarP(&_bedrockagentFlowIdentifier, "flow-identifier", "", "", "Flow Identifier")
	_bedrockagentCmd.Flags().StringVarP(&_bedrockagentFlowVersion, "flow-version", "", "", "Flow Version")
	_bedrockagentCmd.Flags().StringVarP(&_bedrockagentFoundationModel, "foundation-model", "", "", "Foundation Model")
	_bedrockagentCmd.Flags().StringVarP(&_bedrockagentFunctionSchema, "function-schema", "", "", "Function Schema")
	_bedrockagentCmd.Flags().StringVarP(&_bedrockagentGuardrailConfiguration, "guardrail-configuration", "", "", "Guardrail Configuration")
	_bedrockagentCmd.Flags().StringVarP(&_bedrockagentIdleSessionTTLInSeconds, "idle-session-ttlin-seconds", "", "", "Idle Session Ttlin Seconds")
	_bedrockagentCmd.Flags().StringVarP(&_bedrockagentIngestionJobId, "ingestion-job-id", "", "", "Ingestion Job ID")
	_bedrockagentCmd.Flags().StringVarP(&_bedrockagentInstruction, "instruction", "", "", "Instruction")
	_bedrockagentCmd.Flags().StringVarP(&_bedrockagentKnowledgeBaseConfiguration, "knowledge-base-configuration", "", "", "Knowledge Base Configuration")
	_bedrockagentCmd.Flags().StringVarP(&_bedrockagentKnowledgeBaseId, "knowledge-base-id", "", "", "Knowledge Base ID")
	_bedrockagentCmd.Flags().StringVarP(&_bedrockagentKnowledgeBaseState, "knowledge-base-state", "", "", "Knowledge Base State")
	_bedrockagentCmd.Flags().StringVarP(&_bedrockagentMaxResults, "max-results", "", "", "Max Results")
	_bedrockagentCmd.Flags().StringVarP(&_bedrockagentMemoryConfiguration, "memory-configuration", "", "", "Memory Configuration")
	_bedrockagentCmd.Flags().StringVarP(&_bedrockagentName, "name", "", "", "Name")
	_bedrockagentCmd.Flags().StringVarP(&_bedrockagentNextToken, "next-token", "", "", "Next Token")
	_bedrockagentCmd.Flags().StringVarP(&_bedrockagentOrchestrationType, "orchestration-type", "", "", "Orchestration Type")
	_bedrockagentCmd.Flags().StringVarP(&_bedrockagentParentActionGroupSignature, "parent-action-group-signature", "", "", "Parent Action Group Signature")
	_bedrockagentCmd.Flags().StringVarP(&_bedrockagentParentActionGroupSignatureParams, "parent-action-group-signature-params", "", "", "Parent Action Group Signature Params")
	_bedrockagentCmd.Flags().StringVarP(&_bedrockagentPromptIdentifier, "prompt-identifier", "", "", "Prompt Identifier")
	_bedrockagentCmd.Flags().StringVarP(&_bedrockagentPromptOverrideConfiguration, "prompt-override-configuration", "", "", "Prompt Override Configuration")
	_bedrockagentCmd.Flags().StringVarP(&_bedrockagentPromptVersion, "prompt-version", "", "", "Prompt Version")
	_bedrockagentCmd.Flags().StringVarP(&_bedrockagentRelayConversationHistory, "relay-conversation-history", "", "", "Relay Conversation History")
	_bedrockagentCmd.Flags().StringVarP(&_bedrockagentResourceArn, "resource-arn", "", "", "Resource ARN")
	_bedrockagentCmd.Flags().StringVarP(&_bedrockagentRoleArn, "role-arn", "", "", "Role ARN")
	_bedrockagentCmd.Flags().StringVarP(&_bedrockagentRoutingConfiguration, "routing-configuration", "", "", "Routing Configuration")
	_bedrockagentCmd.Flags().StringVarP(&_bedrockagentServerSideEncryptionConfiguration, "server-side-encryption-configuration", "", "", "Server Side Encryption Configuration")
	_bedrockagentCmd.Flags().StringVarP(&_bedrockagentSkipResourceInUseCheck, "skip-resource-in-use-check", "", "", "Skip Resource In Use Check")
	_bedrockagentCmd.Flags().StringVarP(&_bedrockagentSortBy, "sort-by", "", "", "Sort By")
	_bedrockagentCmd.Flags().StringVarP(&_bedrockagentStorageConfiguration, "storage-configuration", "", "", "Storage Configuration")
	_bedrockagentCmd.Flags().StringSliceVarP(&_bedrockagentTagKeys, "tag-keys", "", nil, "Tag Keys")
	_bedrockagentCmd.Flags().StringVarP(&_bedrockagentTags, "tags", "", "", "Tags")
	_bedrockagentCmd.Flags().StringVarP(&_bedrockagentVariants, "variants", "", "", "Variants")
	_bedrockagentCmd.Flags().StringVarP(&_bedrockagentVectorIngestionConfiguration, "vector-ingestion-configuration", "", "", "Vector Ingestion Configuration")

	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentAssociateAgentCollaborator, "associate-agent-collaborator", "", false, "Associate Agent Collaborator")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentAssociateAgentKnowledgeBase, "associate-agent-knowledge-base", "", false, "Associate Agent Knowledge Base")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentCreateAgent, "create-agent", "", false, "Create Agent")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentCreateAgentActionGroup, "create-agent-action-group", "", false, "Create Agent Action Group")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentCreateAgentAlias, "create-agent-alias", "", false, "Create Agent Alias")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentCreateDataSource, "create-data-source", "", false, "Create Data Source")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentCreateFlow, "create-flow", "", false, "Create Flow")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentCreateFlowAlias, "create-flow-alias", "", false, "Create Flow Alias")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentCreateFlowVersion, "create-flow-version", "", false, "Create Flow Version")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentCreateKnowledgeBase, "create-knowledge-base", "", false, "Create Knowledge Base")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentCreatePrompt, "create-prompt", "", false, "Create Prompt")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentCreatePromptVersion, "create-prompt-version", "", false, "Create Prompt Version")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentDeleteAgent, "delete-agent", "", false, "Delete Agent")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentDeleteAgentActionGroup, "delete-agent-action-group", "", false, "Delete Agent Action Group")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentDeleteAgentAlias, "delete-agent-alias", "", false, "Delete Agent Alias")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentDeleteAgentVersion, "delete-agent-version", "", false, "Delete Agent Version")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentDeleteDataSource, "delete-data-source", "", false, "Delete Data Source")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentDeleteFlow, "delete-flow", "", false, "Delete Flow")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentDeleteFlowAlias, "delete-flow-alias", "", false, "Delete Flow Alias")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentDeleteFlowVersion, "delete-flow-version", "", false, "Delete Flow Version")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentDeleteKnowledgeBase, "delete-knowledge-base", "", false, "Delete Knowledge Base")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentDeleteKnowledgeBaseDocuments, "delete-knowledge-base-documents", "", false, "Delete Knowledge Base Documents")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentDeletePrompt, "delete-prompt", "", false, "Delete Prompt")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentDisassociateAgentCollaborator, "disassociate-agent-collaborator", "", false, "Disassociate Agent Collaborator")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentDisassociateAgentKnowledgeBase, "disassociate-agent-knowledge-base", "", false, "Disassociate Agent Knowledge Base")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentGetAgent, "get-agent", "", false, "Get Agent")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentGetAgentActionGroup, "get-agent-action-group", "", false, "Get Agent Action Group")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentGetAgentAlias, "get-agent-alias", "", false, "Get Agent Alias")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentGetAgentCollaborator, "get-agent-collaborator", "", false, "Get Agent Collaborator")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentGetAgentKnowledgeBase, "get-agent-knowledge-base", "", false, "Get Agent Knowledge Base")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentGetAgentVersion, "get-agent-version", "", false, "Get Agent Version")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentGetDataSource, "get-data-source", "", false, "Get Data Source")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentGetFlow, "get-flow", "", false, "Get Flow")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentGetFlowAlias, "get-flow-alias", "", false, "Get Flow Alias")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentGetFlowVersion, "get-flow-version", "", false, "Get Flow Version")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentGetIngestionJob, "get-ingestion-job", "", false, "Get Ingestion Job")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentGetKnowledgeBase, "get-knowledge-base", "", false, "Get Knowledge Base")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentGetKnowledgeBaseDocuments, "get-knowledge-base-documents", "", false, "Get Knowledge Base Documents")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentGetPrompt, "get-prompt", "", false, "Get Prompt")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentIngestKnowledgeBaseDocuments, "ingest-knowledge-base-documents", "", false, "Ingest Knowledge Base Documents")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentListAgentActionGroups, "list-agent-action-groups", "", false, "List Agent Action Groups")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentListAgentAliases, "list-agent-aliases", "", false, "List Agent Aliases")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentListAgentCollaborators, "list-agent-collaborators", "", false, "List Agent Collaborators")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentListAgentKnowledgeBases, "list-agent-knowledge-bases", "", false, "List Agent Knowledge Bases")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentListAgentVersions, "list-agent-versions", "", false, "List Agent Versions")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentListAgents, "list-agents", "", false, "List Agents")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentListDataSources, "list-data-sources", "", false, "List Data Sources")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentListFlowAliases, "list-flow-aliases", "", false, "List Flow Aliases")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentListFlowVersions, "list-flow-versions", "", false, "List Flow Versions")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentListFlows, "list-flows", "", false, "List Flows")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentListIngestionJobs, "list-ingestion-jobs", "", false, "List Ingestion Jobs")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentListKnowledgeBaseDocuments, "list-knowledge-base-documents", "", false, "List Knowledge Base Documents")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentListKnowledgeBases, "list-knowledge-bases", "", false, "List Knowledge Bases")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentListPrompts, "list-prompts", "", false, "List Prompts")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentPrepareAgent, "prepare-agent", "", false, "Prepare Agent")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentPrepareFlow, "prepare-flow", "", false, "Prepare Flow")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentStartIngestionJob, "start-ingestion-job", "", false, "Start Ingestion Job")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentStopIngestionJob, "stop-ingestion-job", "", false, "Stop Ingestion Job")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentTagResource, "tag-resource", "", false, "Tag Resource")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentUntagResource, "untag-resource", "", false, "Untag Resource")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentUpdateAgent, "update-agent", "", false, "Update Agent")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentUpdateAgentActionGroup, "update-agent-action-group", "", false, "Update Agent Action Group")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentUpdateAgentAlias, "update-agent-alias", "", false, "Update Agent Alias")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentUpdateAgentCollaborator, "update-agent-collaborator", "", false, "Update Agent Collaborator")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentUpdateAgentKnowledgeBase, "update-agent-knowledge-base", "", false, "Update Agent Knowledge Base")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentUpdateDataSource, "update-data-source", "", false, "Update Data Source")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentUpdateFlow, "update-flow", "", false, "Update Flow")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentUpdateFlowAlias, "update-flow-alias", "", false, "Update Flow Alias")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentUpdateKnowledgeBase, "update-knowledge-base", "", false, "Update Knowledge Base")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentUpdatePrompt, "update-prompt", "", false, "Update Prompt")
	_bedrockagentCmd.Flags().BoolVarP(&_bedrockagentValidateFlowDefinition, "validate-flow-definition", "", false, "Validate Flow Definition")

}
