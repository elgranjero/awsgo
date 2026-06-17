package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/bedrockagentcorecontrol"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// bedrockagentcorecontrolCmd represents the bedrockagentcorecontrol command
var _bedrockagentcorecontrolCmd = &cobra.Command{
	Use:   "bedrockagentcorecontrol",
	Short: "AWS bedrockagentcorecontrol CLI",
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
		client := bedrockagentcorecontrol.NewFromConfig(cfg)
		if _bedrockagentcorecontrolCreateAgentRuntime {
			bedrockagentcorecontrol_CreateAgentRuntime(cfg, client)
			return
		}
		if _bedrockagentcorecontrolCreateAgentRuntimeEndpoint {
			bedrockagentcorecontrol_CreateAgentRuntimeEndpoint(cfg, client)
			return
		}
		if _bedrockagentcorecontrolCreateApiKeyCredentialProvider {
			bedrockagentcorecontrol_CreateApiKeyCredentialProvider(cfg, client)
			return
		}
		if _bedrockagentcorecontrolCreateBrowser {
			bedrockagentcorecontrol_CreateBrowser(cfg, client)
			return
		}
		if _bedrockagentcorecontrolCreateBrowserProfile {
			bedrockagentcorecontrol_CreateBrowserProfile(cfg, client)
			return
		}
		if _bedrockagentcorecontrolCreateCodeInterpreter {
			bedrockagentcorecontrol_CreateCodeInterpreter(cfg, client)
			return
		}
		if _bedrockagentcorecontrolCreateEvaluator {
			bedrockagentcorecontrol_CreateEvaluator(cfg, client)
			return
		}
		if _bedrockagentcorecontrolCreateGateway {
			bedrockagentcorecontrol_CreateGateway(cfg, client)
			return
		}
		if _bedrockagentcorecontrolCreateGatewayTarget {
			bedrockagentcorecontrol_CreateGatewayTarget(cfg, client)
			return
		}
		if _bedrockagentcorecontrolCreateMemory {
			bedrockagentcorecontrol_CreateMemory(cfg, client)
			return
		}
		if _bedrockagentcorecontrolCreateOauth2CredentialProvider {
			bedrockagentcorecontrol_CreateOauth2CredentialProvider(cfg, client)
			return
		}
		if _bedrockagentcorecontrolCreateOnlineEvaluationConfig {
			bedrockagentcorecontrol_CreateOnlineEvaluationConfig(cfg, client)
			return
		}
		if _bedrockagentcorecontrolCreatePolicy {
			bedrockagentcorecontrol_CreatePolicy(cfg, client)
			return
		}
		if _bedrockagentcorecontrolCreatePolicyEngine {
			bedrockagentcorecontrol_CreatePolicyEngine(cfg, client)
			return
		}
		if _bedrockagentcorecontrolCreateWorkloadIdentity {
			bedrockagentcorecontrol_CreateWorkloadIdentity(cfg, client)
			return
		}
		if _bedrockagentcorecontrolDeleteAgentRuntime {
			bedrockagentcorecontrol_DeleteAgentRuntime(cfg, client)
			return
		}
		if _bedrockagentcorecontrolDeleteAgentRuntimeEndpoint {
			bedrockagentcorecontrol_DeleteAgentRuntimeEndpoint(cfg, client)
			return
		}
		if _bedrockagentcorecontrolDeleteApiKeyCredentialProvider {
			bedrockagentcorecontrol_DeleteApiKeyCredentialProvider(cfg, client)
			return
		}
		if _bedrockagentcorecontrolDeleteBrowser {
			bedrockagentcorecontrol_DeleteBrowser(cfg, client)
			return
		}
		if _bedrockagentcorecontrolDeleteBrowserProfile {
			bedrockagentcorecontrol_DeleteBrowserProfile(cfg, client)
			return
		}
		if _bedrockagentcorecontrolDeleteCodeInterpreter {
			bedrockagentcorecontrol_DeleteCodeInterpreter(cfg, client)
			return
		}
		if _bedrockagentcorecontrolDeleteEvaluator {
			bedrockagentcorecontrol_DeleteEvaluator(cfg, client)
			return
		}
		if _bedrockagentcorecontrolDeleteGateway {
			bedrockagentcorecontrol_DeleteGateway(cfg, client)
			return
		}
		if _bedrockagentcorecontrolDeleteGatewayTarget {
			bedrockagentcorecontrol_DeleteGatewayTarget(cfg, client)
			return
		}
		if _bedrockagentcorecontrolDeleteMemory {
			bedrockagentcorecontrol_DeleteMemory(cfg, client)
			return
		}
		if _bedrockagentcorecontrolDeleteOauth2CredentialProvider {
			bedrockagentcorecontrol_DeleteOauth2CredentialProvider(cfg, client)
			return
		}
		if _bedrockagentcorecontrolDeleteOnlineEvaluationConfig {
			bedrockagentcorecontrol_DeleteOnlineEvaluationConfig(cfg, client)
			return
		}
		if _bedrockagentcorecontrolDeletePolicy {
			bedrockagentcorecontrol_DeletePolicy(cfg, client)
			return
		}
		if _bedrockagentcorecontrolDeletePolicyEngine {
			bedrockagentcorecontrol_DeletePolicyEngine(cfg, client)
			return
		}
		if _bedrockagentcorecontrolDeleteResourcePolicy {
			bedrockagentcorecontrol_DeleteResourcePolicy(cfg, client)
			return
		}
		if _bedrockagentcorecontrolDeleteWorkloadIdentity {
			bedrockagentcorecontrol_DeleteWorkloadIdentity(cfg, client)
			return
		}
		if _bedrockagentcorecontrolGetAgentRuntime {
			bedrockagentcorecontrol_GetAgentRuntime(cfg, client)
			return
		}
		if _bedrockagentcorecontrolGetAgentRuntimeEndpoint {
			bedrockagentcorecontrol_GetAgentRuntimeEndpoint(cfg, client)
			return
		}
		if _bedrockagentcorecontrolGetApiKeyCredentialProvider {
			bedrockagentcorecontrol_GetApiKeyCredentialProvider(cfg, client)
			return
		}
		if _bedrockagentcorecontrolGetBrowser {
			bedrockagentcorecontrol_GetBrowser(cfg, client)
			return
		}
		if _bedrockagentcorecontrolGetBrowserProfile {
			bedrockagentcorecontrol_GetBrowserProfile(cfg, client)
			return
		}
		if _bedrockagentcorecontrolGetCodeInterpreter {
			bedrockagentcorecontrol_GetCodeInterpreter(cfg, client)
			return
		}
		if _bedrockagentcorecontrolGetEvaluator {
			bedrockagentcorecontrol_GetEvaluator(cfg, client)
			return
		}
		if _bedrockagentcorecontrolGetGateway {
			bedrockagentcorecontrol_GetGateway(cfg, client)
			return
		}
		if _bedrockagentcorecontrolGetGatewayTarget {
			bedrockagentcorecontrol_GetGatewayTarget(cfg, client)
			return
		}
		if _bedrockagentcorecontrolGetMemory {
			bedrockagentcorecontrol_GetMemory(cfg, client)
			return
		}
		if _bedrockagentcorecontrolGetOauth2CredentialProvider {
			bedrockagentcorecontrol_GetOauth2CredentialProvider(cfg, client)
			return
		}
		if _bedrockagentcorecontrolGetOnlineEvaluationConfig {
			bedrockagentcorecontrol_GetOnlineEvaluationConfig(cfg, client)
			return
		}
		if _bedrockagentcorecontrolGetPolicy {
			bedrockagentcorecontrol_GetPolicy(cfg, client)
			return
		}
		if _bedrockagentcorecontrolGetPolicyEngine {
			bedrockagentcorecontrol_GetPolicyEngine(cfg, client)
			return
		}
		if _bedrockagentcorecontrolGetPolicyGeneration {
			bedrockagentcorecontrol_GetPolicyGeneration(cfg, client)
			return
		}
		if _bedrockagentcorecontrolGetResourcePolicy {
			bedrockagentcorecontrol_GetResourcePolicy(cfg, client)
			return
		}
		if _bedrockagentcorecontrolGetTokenVault {
			bedrockagentcorecontrol_GetTokenVault(cfg, client)
			return
		}
		if _bedrockagentcorecontrolGetWorkloadIdentity {
			bedrockagentcorecontrol_GetWorkloadIdentity(cfg, client)
			return
		}
		if _bedrockagentcorecontrolListAgentRuntimeEndpoints {
			bedrockagentcorecontrol_ListAgentRuntimeEndpoints(cfg, client)
			return
		}
		if _bedrockagentcorecontrolListAgentRuntimeVersions {
			bedrockagentcorecontrol_ListAgentRuntimeVersions(cfg, client)
			return
		}
		if _bedrockagentcorecontrolListAgentRuntimes {
			bedrockagentcorecontrol_ListAgentRuntimes(cfg, client)
			return
		}
		if _bedrockagentcorecontrolListApiKeyCredentialProviders {
			bedrockagentcorecontrol_ListApiKeyCredentialProviders(cfg, client)
			return
		}
		if _bedrockagentcorecontrolListBrowserProfiles {
			bedrockagentcorecontrol_ListBrowserProfiles(cfg, client)
			return
		}
		if _bedrockagentcorecontrolListBrowsers {
			bedrockagentcorecontrol_ListBrowsers(cfg, client)
			return
		}
		if _bedrockagentcorecontrolListCodeInterpreters {
			bedrockagentcorecontrol_ListCodeInterpreters(cfg, client)
			return
		}
		if _bedrockagentcorecontrolListEvaluators {
			bedrockagentcorecontrol_ListEvaluators(cfg, client)
			return
		}
		if _bedrockagentcorecontrolListGatewayTargets {
			bedrockagentcorecontrol_ListGatewayTargets(cfg, client)
			return
		}
		if _bedrockagentcorecontrolListGateways {
			bedrockagentcorecontrol_ListGateways(cfg, client)
			return
		}
		if _bedrockagentcorecontrolListMemories {
			bedrockagentcorecontrol_ListMemories(cfg, client)
			return
		}
		if _bedrockagentcorecontrolListOauth2CredentialProviders {
			bedrockagentcorecontrol_ListOauth2CredentialProviders(cfg, client)
			return
		}
		if _bedrockagentcorecontrolListOnlineEvaluationConfigs {
			bedrockagentcorecontrol_ListOnlineEvaluationConfigs(cfg, client)
			return
		}
		if _bedrockagentcorecontrolListPolicies {
			bedrockagentcorecontrol_ListPolicies(cfg, client)
			return
		}
		if _bedrockagentcorecontrolListPolicyEngines {
			bedrockagentcorecontrol_ListPolicyEngines(cfg, client)
			return
		}
		if _bedrockagentcorecontrolListPolicyGenerationAssets {
			bedrockagentcorecontrol_ListPolicyGenerationAssets(cfg, client)
			return
		}
		if _bedrockagentcorecontrolListPolicyGenerations {
			bedrockagentcorecontrol_ListPolicyGenerations(cfg, client)
			return
		}
		if _bedrockagentcorecontrolListTagsForResource {
			bedrockagentcorecontrol_ListTagsForResource(cfg, client)
			return
		}
		if _bedrockagentcorecontrolListWorkloadIdentities {
			bedrockagentcorecontrol_ListWorkloadIdentities(cfg, client)
			return
		}
		if _bedrockagentcorecontrolPutResourcePolicy {
			bedrockagentcorecontrol_PutResourcePolicy(cfg, client)
			return
		}
		if _bedrockagentcorecontrolSetTokenVaultCMK {
			bedrockagentcorecontrol_SetTokenVaultCMK(cfg, client)
			return
		}
		if _bedrockagentcorecontrolStartPolicyGeneration {
			bedrockagentcorecontrol_StartPolicyGeneration(cfg, client)
			return
		}
		if _bedrockagentcorecontrolSynchronizeGatewayTargets {
			bedrockagentcorecontrol_SynchronizeGatewayTargets(cfg, client)
			return
		}
		if _bedrockagentcorecontrolTagResource {
			bedrockagentcorecontrol_TagResource(cfg, client)
			return
		}
		if _bedrockagentcorecontrolUntagResource {
			bedrockagentcorecontrol_UntagResource(cfg, client)
			return
		}
		if _bedrockagentcorecontrolUpdateAgentRuntime {
			bedrockagentcorecontrol_UpdateAgentRuntime(cfg, client)
			return
		}
		if _bedrockagentcorecontrolUpdateAgentRuntimeEndpoint {
			bedrockagentcorecontrol_UpdateAgentRuntimeEndpoint(cfg, client)
			return
		}
		if _bedrockagentcorecontrolUpdateApiKeyCredentialProvider {
			bedrockagentcorecontrol_UpdateApiKeyCredentialProvider(cfg, client)
			return
		}
		if _bedrockagentcorecontrolUpdateEvaluator {
			bedrockagentcorecontrol_UpdateEvaluator(cfg, client)
			return
		}
		if _bedrockagentcorecontrolUpdateGateway {
			bedrockagentcorecontrol_UpdateGateway(cfg, client)
			return
		}
		if _bedrockagentcorecontrolUpdateGatewayTarget {
			bedrockagentcorecontrol_UpdateGatewayTarget(cfg, client)
			return
		}
		if _bedrockagentcorecontrolUpdateMemory {
			bedrockagentcorecontrol_UpdateMemory(cfg, client)
			return
		}
		if _bedrockagentcorecontrolUpdateOauth2CredentialProvider {
			bedrockagentcorecontrol_UpdateOauth2CredentialProvider(cfg, client)
			return
		}
		if _bedrockagentcorecontrolUpdateOnlineEvaluationConfig {
			bedrockagentcorecontrol_UpdateOnlineEvaluationConfig(cfg, client)
			return
		}
		if _bedrockagentcorecontrolUpdatePolicy {
			bedrockagentcorecontrol_UpdatePolicy(cfg, client)
			return
		}
		if _bedrockagentcorecontrolUpdatePolicyEngine {
			bedrockagentcorecontrol_UpdatePolicyEngine(cfg, client)
			return
		}
		if _bedrockagentcorecontrolUpdateWorkloadIdentity {
			bedrockagentcorecontrol_UpdateWorkloadIdentity(cfg, client)
			return
		}

	},
}

var (
	_bedrockagentcorecontrolCreateAgentRuntime             bool
	_bedrockagentcorecontrolCreateAgentRuntimeEndpoint     bool
	_bedrockagentcorecontrolCreateApiKeyCredentialProvider bool
	_bedrockagentcorecontrolCreateBrowser                  bool
	_bedrockagentcorecontrolCreateBrowserProfile           bool
	_bedrockagentcorecontrolCreateCodeInterpreter          bool
	_bedrockagentcorecontrolCreateEvaluator                bool
	_bedrockagentcorecontrolCreateGateway                  bool
	_bedrockagentcorecontrolCreateGatewayTarget            bool
	_bedrockagentcorecontrolCreateMemory                   bool
	_bedrockagentcorecontrolCreateOauth2CredentialProvider bool
	_bedrockagentcorecontrolCreateOnlineEvaluationConfig   bool
	_bedrockagentcorecontrolCreatePolicy                   bool
	_bedrockagentcorecontrolCreatePolicyEngine             bool
	_bedrockagentcorecontrolCreateWorkloadIdentity         bool
	_bedrockagentcorecontrolDeleteAgentRuntime             bool
	_bedrockagentcorecontrolDeleteAgentRuntimeEndpoint     bool
	_bedrockagentcorecontrolDeleteApiKeyCredentialProvider bool
	_bedrockagentcorecontrolDeleteBrowser                  bool
	_bedrockagentcorecontrolDeleteBrowserProfile           bool
	_bedrockagentcorecontrolDeleteCodeInterpreter          bool
	_bedrockagentcorecontrolDeleteEvaluator                bool
	_bedrockagentcorecontrolDeleteGateway                  bool
	_bedrockagentcorecontrolDeleteGatewayTarget            bool
	_bedrockagentcorecontrolDeleteMemory                   bool
	_bedrockagentcorecontrolDeleteOauth2CredentialProvider bool
	_bedrockagentcorecontrolDeleteOnlineEvaluationConfig   bool
	_bedrockagentcorecontrolDeletePolicy                   bool
	_bedrockagentcorecontrolDeletePolicyEngine             bool
	_bedrockagentcorecontrolDeleteResourcePolicy           bool
	_bedrockagentcorecontrolDeleteWorkloadIdentity         bool
	_bedrockagentcorecontrolGetAgentRuntime                bool
	_bedrockagentcorecontrolGetAgentRuntimeEndpoint        bool
	_bedrockagentcorecontrolGetApiKeyCredentialProvider    bool
	_bedrockagentcorecontrolGetBrowser                     bool
	_bedrockagentcorecontrolGetBrowserProfile              bool
	_bedrockagentcorecontrolGetCodeInterpreter             bool
	_bedrockagentcorecontrolGetEvaluator                   bool
	_bedrockagentcorecontrolGetGateway                     bool
	_bedrockagentcorecontrolGetGatewayTarget               bool
	_bedrockagentcorecontrolGetMemory                      bool
	_bedrockagentcorecontrolGetOauth2CredentialProvider    bool
	_bedrockagentcorecontrolGetOnlineEvaluationConfig      bool
	_bedrockagentcorecontrolGetPolicy                      bool
	_bedrockagentcorecontrolGetPolicyEngine                bool
	_bedrockagentcorecontrolGetPolicyGeneration            bool
	_bedrockagentcorecontrolGetResourcePolicy              bool
	_bedrockagentcorecontrolGetTokenVault                  bool
	_bedrockagentcorecontrolGetWorkloadIdentity            bool
	_bedrockagentcorecontrolListAgentRuntimeEndpoints      bool
	_bedrockagentcorecontrolListAgentRuntimeVersions       bool
	_bedrockagentcorecontrolListAgentRuntimes              bool
	_bedrockagentcorecontrolListApiKeyCredentialProviders  bool
	_bedrockagentcorecontrolListBrowserProfiles            bool
	_bedrockagentcorecontrolListBrowsers                   bool
	_bedrockagentcorecontrolListCodeInterpreters           bool
	_bedrockagentcorecontrolListEvaluators                 bool
	_bedrockagentcorecontrolListGatewayTargets             bool
	_bedrockagentcorecontrolListGateways                   bool
	_bedrockagentcorecontrolListMemories                   bool
	_bedrockagentcorecontrolListOauth2CredentialProviders  bool
	_bedrockagentcorecontrolListOnlineEvaluationConfigs    bool
	_bedrockagentcorecontrolListPolicies                   bool
	_bedrockagentcorecontrolListPolicyEngines              bool
	_bedrockagentcorecontrolListPolicyGenerationAssets     bool
	_bedrockagentcorecontrolListPolicyGenerations          bool
	_bedrockagentcorecontrolListTagsForResource            bool
	_bedrockagentcorecontrolListWorkloadIdentities         bool
	_bedrockagentcorecontrolPutResourcePolicy              bool
	_bedrockagentcorecontrolSetTokenVaultCMK               bool
	_bedrockagentcorecontrolStartPolicyGeneration          bool
	_bedrockagentcorecontrolSynchronizeGatewayTargets      bool
	_bedrockagentcorecontrolTagResource                    bool
	_bedrockagentcorecontrolUntagResource                  bool
	_bedrockagentcorecontrolUpdateAgentRuntime             bool
	_bedrockagentcorecontrolUpdateAgentRuntimeEndpoint     bool
	_bedrockagentcorecontrolUpdateApiKeyCredentialProvider bool
	_bedrockagentcorecontrolUpdateEvaluator                bool
	_bedrockagentcorecontrolUpdateGateway                  bool
	_bedrockagentcorecontrolUpdateGatewayTarget            bool
	_bedrockagentcorecontrolUpdateMemory                   bool
	_bedrockagentcorecontrolUpdateOauth2CredentialProvider bool
	_bedrockagentcorecontrolUpdateOnlineEvaluationConfig   bool
	_bedrockagentcorecontrolUpdatePolicy                   bool
	_bedrockagentcorecontrolUpdatePolicyEngine             bool
	_bedrockagentcorecontrolUpdateWorkloadIdentity         bool

	_bedrockagentcorecontrolAgentRuntimeArtifact             string
	_bedrockagentcorecontrolAgentRuntimeId                   string
	_bedrockagentcorecontrolAgentRuntimeName                 string
	_bedrockagentcorecontrolAgentRuntimeVersion              string
	_bedrockagentcorecontrolAllowedResourceOauth2ReturnUrls  []string
	_bedrockagentcorecontrolApiKey                           string
	_bedrockagentcorecontrolAuthorizerConfiguration          string
	_bedrockagentcorecontrolAuthorizerType                   string
	_bedrockagentcorecontrolBrowserId                        string
	_bedrockagentcorecontrolBrowserSigning                   string
	_bedrockagentcorecontrolClientToken                      string
	_bedrockagentcorecontrolCodeInterpreterId                string
	_bedrockagentcorecontrolContent                          string
	_bedrockagentcorecontrolCredentialProviderConfigurations string
	_bedrockagentcorecontrolCredentialProviderVendor         string
	_bedrockagentcorecontrolDataSourceConfig                 string
	_bedrockagentcorecontrolDefinition                       string
	_bedrockagentcorecontrolDescription                      string
	_bedrockagentcorecontrolEnableOnCreate                   string
	_bedrockagentcorecontrolEncryptionKeyArn                 string
	_bedrockagentcorecontrolEndpointName                     string
	_bedrockagentcorecontrolEnvironmentVariables             string
	_bedrockagentcorecontrolEvaluationExecutionRoleArn       string
	_bedrockagentcorecontrolEvaluatorConfig                  string
	_bedrockagentcorecontrolEvaluatorId                      string
	_bedrockagentcorecontrolEvaluatorName                    string
	_bedrockagentcorecontrolEvaluators                       string
	_bedrockagentcorecontrolEventExpiryDuration              string
	_bedrockagentcorecontrolExceptionLevel                   string
	_bedrockagentcorecontrolExecutionRoleArn                 string
	_bedrockagentcorecontrolExecutionStatus                  string
	_bedrockagentcorecontrolGatewayIdentifier                string
	_bedrockagentcorecontrolInterceptorConfigurations        string
	_bedrockagentcorecontrolKmsConfiguration                 string
	_bedrockagentcorecontrolKmsKeyArn                        string
	_bedrockagentcorecontrolLevel                            string
	_bedrockagentcorecontrolLifecycleConfiguration           string
	_bedrockagentcorecontrolMaxResults                       string
	_bedrockagentcorecontrolMemoryExecutionRoleArn           string
	_bedrockagentcorecontrolMemoryId                         string
	_bedrockagentcorecontrolMemoryStrategies                 string
	_bedrockagentcorecontrolMetadataConfiguration            string
	_bedrockagentcorecontrolName                             string
	_bedrockagentcorecontrolNetworkConfiguration             string
	_bedrockagentcorecontrolNextToken                        string
	_bedrockagentcorecontrolOauth2ProviderConfigInput        string
	_bedrockagentcorecontrolOnlineEvaluationConfigId         string
	_bedrockagentcorecontrolOnlineEvaluationConfigName       string
	_bedrockagentcorecontrolPolicy                           string
	_bedrockagentcorecontrolPolicyEngineConfiguration        string
	_bedrockagentcorecontrolPolicyEngineId                   string
	_bedrockagentcorecontrolPolicyGenerationId               string
	_bedrockagentcorecontrolPolicyId                         string
	_bedrockagentcorecontrolProfileId                        string
	_bedrockagentcorecontrolProtocolConfiguration            string
	_bedrockagentcorecontrolProtocolType                     string
	_bedrockagentcorecontrolRecording                        string
	_bedrockagentcorecontrolRequestHeaderConfiguration       string
	_bedrockagentcorecontrolResource                         string
	_bedrockagentcorecontrolResourceArn                      string
	_bedrockagentcorecontrolRoleArn                          string
	_bedrockagentcorecontrolRule                             string
	_bedrockagentcorecontrolTagKeys                          []string
	_bedrockagentcorecontrolTags                             string
	_bedrockagentcorecontrolTargetConfiguration              string
	_bedrockagentcorecontrolTargetId                         string
	_bedrockagentcorecontrolTargetIdList                     []string
	_bedrockagentcorecontrolTargetResourceScope              string
	_bedrockagentcorecontrolTokenVaultId                     string
	_bedrockagentcorecontrolType                             string
	_bedrockagentcorecontrolValidationMode                   string
	_bedrockagentcorecontrolView                             string
)

// Creates an Amazon Bedrock AgentCore Runtime.
func bedrockagentcorecontrol_CreateAgentRuntime(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.CreateAgentRuntimeInput{
		// AgentRuntimeArtifact: types.AgentRuntimeArtifact, // Required
		// AgentRuntimeName: *string, // Required
		// NetworkConfiguration: *types.NetworkConfiguration, // Required
		// RoleArn: *string, // Required
	}

	if len(_bedrockagentcorecontrolAgentRuntimeArtifact) > 0 {
		if err := assignInputField(input, "AgentRuntimeArtifact", _bedrockagentcorecontrolAgentRuntimeArtifact); err != nil {
			log.Errorf("invalid --agent-runtime-artifact: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolAgentRuntimeName) > 0 {
		input.AgentRuntimeName = aws.String(_bedrockagentcorecontrolAgentRuntimeName)
	}
	if len(_bedrockagentcorecontrolNetworkConfiguration) > 0 {
		if err := assignInputField(input, "NetworkConfiguration", _bedrockagentcorecontrolNetworkConfiguration); err != nil {
			log.Errorf("invalid --network-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolRoleArn) > 0 {
		input.RoleArn = aws.String(_bedrockagentcorecontrolRoleArn)
	}
	if len(_bedrockagentcorecontrolAuthorizerConfiguration) > 0 {
		if err := assignInputField(input, "AuthorizerConfiguration", _bedrockagentcorecontrolAuthorizerConfiguration); err != nil {
			log.Errorf("invalid --authorizer-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolClientToken) > 0 {
		input.ClientToken = aws.String(_bedrockagentcorecontrolClientToken)
	}
	if len(_bedrockagentcorecontrolDescription) > 0 {
		input.Description = aws.String(_bedrockagentcorecontrolDescription)
	}
	if len(_bedrockagentcorecontrolEnvironmentVariables) > 0 {
		if err := assignInputField(input, "EnvironmentVariables", _bedrockagentcorecontrolEnvironmentVariables); err != nil {
			log.Errorf("invalid --environment-variables: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolLifecycleConfiguration) > 0 {
		if err := assignInputField(input, "LifecycleConfiguration", _bedrockagentcorecontrolLifecycleConfiguration); err != nil {
			log.Errorf("invalid --lifecycle-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolProtocolConfiguration) > 0 {
		if err := assignInputField(input, "ProtocolConfiguration", _bedrockagentcorecontrolProtocolConfiguration); err != nil {
			log.Errorf("invalid --protocol-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolRequestHeaderConfiguration) > 0 {
		if err := assignInputField(input, "RequestHeaderConfiguration", _bedrockagentcorecontrolRequestHeaderConfiguration); err != nil {
			log.Errorf("invalid --request-header-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolTags) > 0 {
		if err := assignInputField(input, "Tags", _bedrockagentcorecontrolTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAgentRuntime(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an AgentCore Runtime endpoint.
func bedrockagentcorecontrol_CreateAgentRuntimeEndpoint(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.CreateAgentRuntimeEndpointInput{
		// AgentRuntimeId: *string, // Required
		// Name: *string, // Required
	}

	if len(_bedrockagentcorecontrolAgentRuntimeId) > 0 {
		input.AgentRuntimeId = aws.String(_bedrockagentcorecontrolAgentRuntimeId)
	}
	if len(_bedrockagentcorecontrolName) > 0 {
		input.Name = aws.String(_bedrockagentcorecontrolName)
	}
	if len(_bedrockagentcorecontrolAgentRuntimeVersion) > 0 {
		input.AgentRuntimeVersion = aws.String(_bedrockagentcorecontrolAgentRuntimeVersion)
	}
	if len(_bedrockagentcorecontrolClientToken) > 0 {
		input.ClientToken = aws.String(_bedrockagentcorecontrolClientToken)
	}
	if len(_bedrockagentcorecontrolDescription) > 0 {
		input.Description = aws.String(_bedrockagentcorecontrolDescription)
	}
	if len(_bedrockagentcorecontrolTags) > 0 {
		if err := assignInputField(input, "Tags", _bedrockagentcorecontrolTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAgentRuntimeEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new API key credential provider.
func bedrockagentcorecontrol_CreateApiKeyCredentialProvider(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.CreateApiKeyCredentialProviderInput{
		// ApiKey: *string, // Required
		// Name: *string, // Required
	}

	if len(_bedrockagentcorecontrolApiKey) > 0 {
		input.ApiKey = aws.String(_bedrockagentcorecontrolApiKey)
	}
	if len(_bedrockagentcorecontrolName) > 0 {
		input.Name = aws.String(_bedrockagentcorecontrolName)
	}
	if len(_bedrockagentcorecontrolTags) > 0 {
		if err := assignInputField(input, "Tags", _bedrockagentcorecontrolTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateApiKeyCredentialProvider(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a custom browser.
func bedrockagentcorecontrol_CreateBrowser(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.CreateBrowserInput{
		// Name: *string, // Required
		// NetworkConfiguration: *types.BrowserNetworkConfiguration, // Required
	}

	if len(_bedrockagentcorecontrolName) > 0 {
		input.Name = aws.String(_bedrockagentcorecontrolName)
	}
	if len(_bedrockagentcorecontrolNetworkConfiguration) > 0 {
		if err := assignInputField(input, "NetworkConfiguration", _bedrockagentcorecontrolNetworkConfiguration); err != nil {
			log.Errorf("invalid --network-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolBrowserSigning) > 0 {
		if err := assignInputField(input, "BrowserSigning", _bedrockagentcorecontrolBrowserSigning); err != nil {
			log.Errorf("invalid --browser-signing: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolClientToken) > 0 {
		input.ClientToken = aws.String(_bedrockagentcorecontrolClientToken)
	}
	if len(_bedrockagentcorecontrolDescription) > 0 {
		input.Description = aws.String(_bedrockagentcorecontrolDescription)
	}
	if len(_bedrockagentcorecontrolExecutionRoleArn) > 0 {
		input.ExecutionRoleArn = aws.String(_bedrockagentcorecontrolExecutionRoleArn)
	}
	if len(_bedrockagentcorecontrolRecording) > 0 {
		if err := assignInputField(input, "Recording", _bedrockagentcorecontrolRecording); err != nil {
			log.Errorf("invalid --recording: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolTags) > 0 {
		if err := assignInputField(input, "Tags", _bedrockagentcorecontrolTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateBrowser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a browser profile in Amazon Bedrock AgentCore. A browser profile stores
// persistent browser data such as cookies, local storage, session storage, and
// browsing history that can be saved from browser sessions and reused in
// subsequent sessions.
func bedrockagentcorecontrol_CreateBrowserProfile(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.CreateBrowserProfileInput{
		// Name: *string, // Required
	}

	if len(_bedrockagentcorecontrolName) > 0 {
		input.Name = aws.String(_bedrockagentcorecontrolName)
	}
	if len(_bedrockagentcorecontrolClientToken) > 0 {
		input.ClientToken = aws.String(_bedrockagentcorecontrolClientToken)
	}
	if len(_bedrockagentcorecontrolDescription) > 0 {
		input.Description = aws.String(_bedrockagentcorecontrolDescription)
	}
	if len(_bedrockagentcorecontrolTags) > 0 {
		if err := assignInputField(input, "Tags", _bedrockagentcorecontrolTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateBrowserProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a custom code interpreter.
func bedrockagentcorecontrol_CreateCodeInterpreter(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.CreateCodeInterpreterInput{
		// Name: *string, // Required
		// NetworkConfiguration: *types.CodeInterpreterNetworkConfiguration, // Required
	}

	if len(_bedrockagentcorecontrolName) > 0 {
		input.Name = aws.String(_bedrockagentcorecontrolName)
	}
	if len(_bedrockagentcorecontrolNetworkConfiguration) > 0 {
		if err := assignInputField(input, "NetworkConfiguration", _bedrockagentcorecontrolNetworkConfiguration); err != nil {
			log.Errorf("invalid --network-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolClientToken) > 0 {
		input.ClientToken = aws.String(_bedrockagentcorecontrolClientToken)
	}
	if len(_bedrockagentcorecontrolDescription) > 0 {
		input.Description = aws.String(_bedrockagentcorecontrolDescription)
	}
	if len(_bedrockagentcorecontrolExecutionRoleArn) > 0 {
		input.ExecutionRoleArn = aws.String(_bedrockagentcorecontrolExecutionRoleArn)
	}
	if len(_bedrockagentcorecontrolTags) > 0 {
		if err := assignInputField(input, "Tags", _bedrockagentcorecontrolTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCodeInterpreter(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a custom evaluator for agent quality assessment. Custom evaluators use
// LLM-as-a-Judge configurations with user-defined prompts, rating scales, and
// model settings to evaluate agent performance at tool call, trace, or session
// levels.
func bedrockagentcorecontrol_CreateEvaluator(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.CreateEvaluatorInput{
		// EvaluatorConfig: types.EvaluatorConfig, // Required
		// EvaluatorName: *string, // Required
		// Level: types.EvaluatorLevel, // Required
	}

	if len(_bedrockagentcorecontrolEvaluatorConfig) > 0 {
		if err := assignInputField(input, "EvaluatorConfig", _bedrockagentcorecontrolEvaluatorConfig); err != nil {
			log.Errorf("invalid --evaluator-config: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolEvaluatorName) > 0 {
		input.EvaluatorName = aws.String(_bedrockagentcorecontrolEvaluatorName)
	}
	if len(_bedrockagentcorecontrolLevel) > 0 {
		if err := assignInputField(input, "Level", _bedrockagentcorecontrolLevel); err != nil {
			log.Errorf("invalid --level: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolClientToken) > 0 {
		input.ClientToken = aws.String(_bedrockagentcorecontrolClientToken)
	}
	if len(_bedrockagentcorecontrolDescription) > 0 {
		input.Description = aws.String(_bedrockagentcorecontrolDescription)
	}
	if len(_bedrockagentcorecontrolTags) > 0 {
		if err := assignInputField(input, "Tags", _bedrockagentcorecontrolTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateEvaluator(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a gateway for Amazon Bedrock Agent. A gateway serves as an integration
// point between your agent and external services.
//
// If you specify CUSTOM_JWT as the authorizerType , you must provide an
// authorizerConfiguration .
func bedrockagentcorecontrol_CreateGateway(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.CreateGatewayInput{
		// AuthorizerType: types.AuthorizerType, // Required
		// Name: *string, // Required
		// ProtocolType: types.GatewayProtocolType, // Required
		// RoleArn: *string, // Required
	}

	if len(_bedrockagentcorecontrolAuthorizerType) > 0 {
		if err := assignInputField(input, "AuthorizerType", _bedrockagentcorecontrolAuthorizerType); err != nil {
			log.Errorf("invalid --authorizer-type: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolName) > 0 {
		input.Name = aws.String(_bedrockagentcorecontrolName)
	}
	if len(_bedrockagentcorecontrolProtocolType) > 0 {
		if err := assignInputField(input, "ProtocolType", _bedrockagentcorecontrolProtocolType); err != nil {
			log.Errorf("invalid --protocol-type: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolRoleArn) > 0 {
		input.RoleArn = aws.String(_bedrockagentcorecontrolRoleArn)
	}
	if len(_bedrockagentcorecontrolAuthorizerConfiguration) > 0 {
		if err := assignInputField(input, "AuthorizerConfiguration", _bedrockagentcorecontrolAuthorizerConfiguration); err != nil {
			log.Errorf("invalid --authorizer-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolClientToken) > 0 {
		input.ClientToken = aws.String(_bedrockagentcorecontrolClientToken)
	}
	if len(_bedrockagentcorecontrolDescription) > 0 {
		input.Description = aws.String(_bedrockagentcorecontrolDescription)
	}
	if len(_bedrockagentcorecontrolExceptionLevel) > 0 {
		if err := assignInputField(input, "ExceptionLevel", _bedrockagentcorecontrolExceptionLevel); err != nil {
			log.Errorf("invalid --exception-level: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolInterceptorConfigurations) > 0 {
		if err := assignInputField(input, "InterceptorConfigurations", _bedrockagentcorecontrolInterceptorConfigurations); err != nil {
			log.Errorf("invalid --interceptor-configurations: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolKmsKeyArn) > 0 {
		input.KmsKeyArn = aws.String(_bedrockagentcorecontrolKmsKeyArn)
	}
	if len(_bedrockagentcorecontrolPolicyEngineConfiguration) > 0 {
		if err := assignInputField(input, "PolicyEngineConfiguration", _bedrockagentcorecontrolPolicyEngineConfiguration); err != nil {
			log.Errorf("invalid --policy-engine-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolProtocolConfiguration) > 0 {
		if err := assignInputField(input, "ProtocolConfiguration", _bedrockagentcorecontrolProtocolConfiguration); err != nil {
			log.Errorf("invalid --protocol-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolTags) > 0 {
		if err := assignInputField(input, "Tags", _bedrockagentcorecontrolTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateGateway(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a target for a gateway. A target defines an endpoint that the gateway
// can connect to.
func bedrockagentcorecontrol_CreateGatewayTarget(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.CreateGatewayTargetInput{
		// GatewayIdentifier: *string, // Required
		// Name: *string, // Required
		// TargetConfiguration: types.TargetConfiguration, // Required
	}

	if len(_bedrockagentcorecontrolGatewayIdentifier) > 0 {
		input.GatewayIdentifier = aws.String(_bedrockagentcorecontrolGatewayIdentifier)
	}
	if len(_bedrockagentcorecontrolName) > 0 {
		input.Name = aws.String(_bedrockagentcorecontrolName)
	}
	if len(_bedrockagentcorecontrolTargetConfiguration) > 0 {
		if err := assignInputField(input, "TargetConfiguration", _bedrockagentcorecontrolTargetConfiguration); err != nil {
			log.Errorf("invalid --target-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolClientToken) > 0 {
		input.ClientToken = aws.String(_bedrockagentcorecontrolClientToken)
	}
	if len(_bedrockagentcorecontrolCredentialProviderConfigurations) > 0 {
		if err := assignInputField(input, "CredentialProviderConfigurations", _bedrockagentcorecontrolCredentialProviderConfigurations); err != nil {
			log.Errorf("invalid --credential-provider-configurations: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolDescription) > 0 {
		input.Description = aws.String(_bedrockagentcorecontrolDescription)
	}
	if len(_bedrockagentcorecontrolMetadataConfiguration) > 0 {
		if err := assignInputField(input, "MetadataConfiguration", _bedrockagentcorecontrolMetadataConfiguration); err != nil {
			log.Errorf("invalid --metadata-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateGatewayTarget(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new Amazon Bedrock AgentCore Memory resource.
func bedrockagentcorecontrol_CreateMemory(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.CreateMemoryInput{
		// EventExpiryDuration: *int32, // Required
		// Name: *string, // Required
	}

	if len(_bedrockagentcorecontrolEventExpiryDuration) > 0 {
		if err := assignInputField(input, "EventExpiryDuration", _bedrockagentcorecontrolEventExpiryDuration); err != nil {
			log.Errorf("invalid --event-expiry-duration: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolName) > 0 {
		input.Name = aws.String(_bedrockagentcorecontrolName)
	}
	if len(_bedrockagentcorecontrolClientToken) > 0 {
		input.ClientToken = aws.String(_bedrockagentcorecontrolClientToken)
	}
	if len(_bedrockagentcorecontrolDescription) > 0 {
		input.Description = aws.String(_bedrockagentcorecontrolDescription)
	}
	if len(_bedrockagentcorecontrolEncryptionKeyArn) > 0 {
		input.EncryptionKeyArn = aws.String(_bedrockagentcorecontrolEncryptionKeyArn)
	}
	if len(_bedrockagentcorecontrolMemoryExecutionRoleArn) > 0 {
		input.MemoryExecutionRoleArn = aws.String(_bedrockagentcorecontrolMemoryExecutionRoleArn)
	}
	if len(_bedrockagentcorecontrolMemoryStrategies) > 0 {
		if err := assignInputField(input, "MemoryStrategies", _bedrockagentcorecontrolMemoryStrategies); err != nil {
			log.Errorf("invalid --memory-strategies: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolTags) > 0 {
		if err := assignInputField(input, "Tags", _bedrockagentcorecontrolTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateMemory(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new OAuth2 credential provider.
func bedrockagentcorecontrol_CreateOauth2CredentialProvider(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.CreateOauth2CredentialProviderInput{
		// CredentialProviderVendor: types.CredentialProviderVendorType, // Required
		// Name: *string, // Required
		// Oauth2ProviderConfigInput: types.Oauth2ProviderConfigInput, // Required
	}

	if len(_bedrockagentcorecontrolCredentialProviderVendor) > 0 {
		if err := assignInputField(input, "CredentialProviderVendor", _bedrockagentcorecontrolCredentialProviderVendor); err != nil {
			log.Errorf("invalid --credential-provider-vendor: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolName) > 0 {
		input.Name = aws.String(_bedrockagentcorecontrolName)
	}
	if len(_bedrockagentcorecontrolOauth2ProviderConfigInput) > 0 {
		if err := assignInputField(input, "Oauth2ProviderConfigInput", _bedrockagentcorecontrolOauth2ProviderConfigInput); err != nil {
			log.Errorf("invalid --oauth2-provider-config-input: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolTags) > 0 {
		if err := assignInputField(input, "Tags", _bedrockagentcorecontrolTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateOauth2CredentialProvider(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an online evaluation configuration for continuous monitoring of agent
// performance. Online evaluation automatically samples live traffic from
// CloudWatch logs at specified rates and applies evaluators to assess agent
// quality in production.
func bedrockagentcorecontrol_CreateOnlineEvaluationConfig(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.CreateOnlineEvaluationConfigInput{
		// DataSourceConfig: types.DataSourceConfig, // Required
		// EnableOnCreate: *bool, // Required
		// EvaluationExecutionRoleArn: *string, // Required
		// Evaluators: []types.EvaluatorReference, // Required
		// OnlineEvaluationConfigName: *string, // Required
		// Rule: *types.Rule, // Required
	}

	if len(_bedrockagentcorecontrolDataSourceConfig) > 0 {
		if err := assignInputField(input, "DataSourceConfig", _bedrockagentcorecontrolDataSourceConfig); err != nil {
			log.Errorf("invalid --data-source-config: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolEnableOnCreate) > 0 {
		if err := assignInputField(input, "EnableOnCreate", _bedrockagentcorecontrolEnableOnCreate); err != nil {
			log.Errorf("invalid --enable-on-create: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolEvaluationExecutionRoleArn) > 0 {
		input.EvaluationExecutionRoleArn = aws.String(_bedrockagentcorecontrolEvaluationExecutionRoleArn)
	}
	if len(_bedrockagentcorecontrolEvaluators) > 0 {
		if err := assignInputField(input, "Evaluators", _bedrockagentcorecontrolEvaluators); err != nil {
			log.Errorf("invalid --evaluators: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolOnlineEvaluationConfigName) > 0 {
		input.OnlineEvaluationConfigName = aws.String(_bedrockagentcorecontrolOnlineEvaluationConfigName)
	}
	if len(_bedrockagentcorecontrolRule) > 0 {
		if err := assignInputField(input, "Rule", _bedrockagentcorecontrolRule); err != nil {
			log.Errorf("invalid --rule: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolClientToken) > 0 {
		input.ClientToken = aws.String(_bedrockagentcorecontrolClientToken)
	}
	if len(_bedrockagentcorecontrolDescription) > 0 {
		input.Description = aws.String(_bedrockagentcorecontrolDescription)
	}
	if len(_bedrockagentcorecontrolTags) > 0 {
		if err := assignInputField(input, "Tags", _bedrockagentcorecontrolTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateOnlineEvaluationConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a policy within the AgentCore Policy system. Policies provide
// real-time, deterministic control over agentic interactions with AgentCore
// Gateway. Using the Cedar policy language, you can define fine-grained policies
// that specify which interactions with Gateway tools are permitted based on input
// parameters and OAuth claims, ensuring agents operate within defined boundaries
// and business rules. The policy is validated during creation against the Cedar
// schema generated from the Gateway's tools' input schemas, which defines the
// available tools, their parameters, and expected data types. This is an
// asynchronous operation. Use the [GetPolicy]operation to poll the status field to track
// completion.
//
// [GetPolicy]: https://docs.aws.amazon.com/bedrock-agentcore-control/latest/APIReference/API_GetPolicy.html
func bedrockagentcorecontrol_CreatePolicy(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.CreatePolicyInput{
		// Definition: types.PolicyDefinition, // Required
		// Name: *string, // Required
		// PolicyEngineId: *string, // Required
	}

	if len(_bedrockagentcorecontrolDefinition) > 0 {
		if err := assignInputField(input, "Definition", _bedrockagentcorecontrolDefinition); err != nil {
			log.Errorf("invalid --definition: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolName) > 0 {
		input.Name = aws.String(_bedrockagentcorecontrolName)
	}
	if len(_bedrockagentcorecontrolPolicyEngineId) > 0 {
		input.PolicyEngineId = aws.String(_bedrockagentcorecontrolPolicyEngineId)
	}
	if len(_bedrockagentcorecontrolClientToken) > 0 {
		input.ClientToken = aws.String(_bedrockagentcorecontrolClientToken)
	}
	if len(_bedrockagentcorecontrolDescription) > 0 {
		input.Description = aws.String(_bedrockagentcorecontrolDescription)
	}
	if len(_bedrockagentcorecontrolValidationMode) > 0 {
		if err := assignInputField(input, "ValidationMode", _bedrockagentcorecontrolValidationMode); err != nil {
			log.Errorf("invalid --validation-mode: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreatePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new policy engine within the AgentCore Policy system. A policy engine
// is a collection of policies that evaluates and authorizes agent tool calls. When
// associated with Gateways (each Gateway can be associated with at most one policy
// engine, but multiple Gateways can be associated with the same engine), the
// policy engine intercepts all agent requests and determines whether to allow or
// deny each action based on the defined policies. This is an asynchronous
// operation. Use the [GetPolicyEngine]operation to poll the status field to track completion.
//
// [GetPolicyEngine]: https://docs.aws.amazon.com/bedrock-agentcore-control/latest/APIReference/API_GetPolicyEngine.html
func bedrockagentcorecontrol_CreatePolicyEngine(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.CreatePolicyEngineInput{
		// Name: *string, // Required
	}

	if len(_bedrockagentcorecontrolName) > 0 {
		input.Name = aws.String(_bedrockagentcorecontrolName)
	}
	if len(_bedrockagentcorecontrolClientToken) > 0 {
		input.ClientToken = aws.String(_bedrockagentcorecontrolClientToken)
	}
	if len(_bedrockagentcorecontrolDescription) > 0 {
		input.Description = aws.String(_bedrockagentcorecontrolDescription)
	}
	if len(_bedrockagentcorecontrolEncryptionKeyArn) > 0 {
		input.EncryptionKeyArn = aws.String(_bedrockagentcorecontrolEncryptionKeyArn)
	}
	if len(_bedrockagentcorecontrolTags) > 0 {
		if err := assignInputField(input, "Tags", _bedrockagentcorecontrolTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreatePolicyEngine(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new workload identity.
func bedrockagentcorecontrol_CreateWorkloadIdentity(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.CreateWorkloadIdentityInput{
		// Name: *string, // Required
	}

	if len(_bedrockagentcorecontrolName) > 0 {
		input.Name = aws.String(_bedrockagentcorecontrolName)
	}
	if len(_bedrockagentcorecontrolAllowedResourceOauth2ReturnUrls) > 0 {
		input.AllowedResourceOauth2ReturnUrls = append([]string(nil), _bedrockagentcorecontrolAllowedResourceOauth2ReturnUrls...)
	}
	if len(_bedrockagentcorecontrolTags) > 0 {
		if err := assignInputField(input, "Tags", _bedrockagentcorecontrolTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateWorkloadIdentity(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Amazon Bedrock AgentCore Runtime.
func bedrockagentcorecontrol_DeleteAgentRuntime(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.DeleteAgentRuntimeInput{
		// AgentRuntimeId: *string, // Required
	}

	if len(_bedrockagentcorecontrolAgentRuntimeId) > 0 {
		input.AgentRuntimeId = aws.String(_bedrockagentcorecontrolAgentRuntimeId)
	}
	if len(_bedrockagentcorecontrolClientToken) > 0 {
		input.ClientToken = aws.String(_bedrockagentcorecontrolClientToken)
	}

	if resp, err := client.DeleteAgentRuntime(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an AAgentCore Runtime endpoint.
func bedrockagentcorecontrol_DeleteAgentRuntimeEndpoint(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.DeleteAgentRuntimeEndpointInput{
		// AgentRuntimeId: *string, // Required
		// EndpointName: *string, // Required
	}

	if len(_bedrockagentcorecontrolAgentRuntimeId) > 0 {
		input.AgentRuntimeId = aws.String(_bedrockagentcorecontrolAgentRuntimeId)
	}
	if len(_bedrockagentcorecontrolEndpointName) > 0 {
		input.EndpointName = aws.String(_bedrockagentcorecontrolEndpointName)
	}
	if len(_bedrockagentcorecontrolClientToken) > 0 {
		input.ClientToken = aws.String(_bedrockagentcorecontrolClientToken)
	}

	if resp, err := client.DeleteAgentRuntimeEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an API key credential provider.
func bedrockagentcorecontrol_DeleteApiKeyCredentialProvider(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.DeleteApiKeyCredentialProviderInput{
		// Name: *string, // Required
	}

	if len(_bedrockagentcorecontrolName) > 0 {
		input.Name = aws.String(_bedrockagentcorecontrolName)
	}

	if resp, err := client.DeleteApiKeyCredentialProvider(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a custom browser.
func bedrockagentcorecontrol_DeleteBrowser(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.DeleteBrowserInput{
		// BrowserId: *string, // Required
	}

	if len(_bedrockagentcorecontrolBrowserId) > 0 {
		input.BrowserId = aws.String(_bedrockagentcorecontrolBrowserId)
	}
	if len(_bedrockagentcorecontrolClientToken) > 0 {
		input.ClientToken = aws.String(_bedrockagentcorecontrolClientToken)
	}

	if resp, err := client.DeleteBrowser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a browser profile.
func bedrockagentcorecontrol_DeleteBrowserProfile(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.DeleteBrowserProfileInput{
		// ProfileId: *string, // Required
	}

	if len(_bedrockagentcorecontrolProfileId) > 0 {
		input.ProfileId = aws.String(_bedrockagentcorecontrolProfileId)
	}
	if len(_bedrockagentcorecontrolClientToken) > 0 {
		input.ClientToken = aws.String(_bedrockagentcorecontrolClientToken)
	}

	if resp, err := client.DeleteBrowserProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a custom code interpreter.
func bedrockagentcorecontrol_DeleteCodeInterpreter(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.DeleteCodeInterpreterInput{
		// CodeInterpreterId: *string, // Required
	}

	if len(_bedrockagentcorecontrolCodeInterpreterId) > 0 {
		input.CodeInterpreterId = aws.String(_bedrockagentcorecontrolCodeInterpreterId)
	}
	if len(_bedrockagentcorecontrolClientToken) > 0 {
		input.ClientToken = aws.String(_bedrockagentcorecontrolClientToken)
	}

	if resp, err := client.DeleteCodeInterpreter(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a custom evaluator. Builtin evaluators cannot be deleted. The
// evaluator must not be referenced by any active online evaluation configurations.
func bedrockagentcorecontrol_DeleteEvaluator(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.DeleteEvaluatorInput{
		// EvaluatorId: *string, // Required
	}

	if len(_bedrockagentcorecontrolEvaluatorId) > 0 {
		input.EvaluatorId = aws.String(_bedrockagentcorecontrolEvaluatorId)
	}

	if resp, err := client.DeleteEvaluator(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a gateway.
func bedrockagentcorecontrol_DeleteGateway(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.DeleteGatewayInput{
		// GatewayIdentifier: *string, // Required
	}

	if len(_bedrockagentcorecontrolGatewayIdentifier) > 0 {
		input.GatewayIdentifier = aws.String(_bedrockagentcorecontrolGatewayIdentifier)
	}

	if resp, err := client.DeleteGateway(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a gateway target.
func bedrockagentcorecontrol_DeleteGatewayTarget(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.DeleteGatewayTargetInput{
		// GatewayIdentifier: *string, // Required
		// TargetId: *string, // Required
	}

	if len(_bedrockagentcorecontrolGatewayIdentifier) > 0 {
		input.GatewayIdentifier = aws.String(_bedrockagentcorecontrolGatewayIdentifier)
	}
	if len(_bedrockagentcorecontrolTargetId) > 0 {
		input.TargetId = aws.String(_bedrockagentcorecontrolTargetId)
	}

	if resp, err := client.DeleteGatewayTarget(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Amazon Bedrock AgentCore Memory resource.
func bedrockagentcorecontrol_DeleteMemory(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.DeleteMemoryInput{
		// MemoryId: *string, // Required
	}

	if len(_bedrockagentcorecontrolMemoryId) > 0 {
		input.MemoryId = aws.String(_bedrockagentcorecontrolMemoryId)
	}
	if len(_bedrockagentcorecontrolClientToken) > 0 {
		input.ClientToken = aws.String(_bedrockagentcorecontrolClientToken)
	}

	if resp, err := client.DeleteMemory(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an OAuth2 credential provider.
func bedrockagentcorecontrol_DeleteOauth2CredentialProvider(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.DeleteOauth2CredentialProviderInput{
		// Name: *string, // Required
	}

	if len(_bedrockagentcorecontrolName) > 0 {
		input.Name = aws.String(_bedrockagentcorecontrolName)
	}

	if resp, err := client.DeleteOauth2CredentialProvider(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an online evaluation configuration and stops any ongoing evaluation
// processes associated with it.
func bedrockagentcorecontrol_DeleteOnlineEvaluationConfig(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.DeleteOnlineEvaluationConfigInput{
		// OnlineEvaluationConfigId: *string, // Required
	}

	if len(_bedrockagentcorecontrolOnlineEvaluationConfigId) > 0 {
		input.OnlineEvaluationConfigId = aws.String(_bedrockagentcorecontrolOnlineEvaluationConfigId)
	}

	if resp, err := client.DeleteOnlineEvaluationConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing policy from the AgentCore Policy system. Once deleted, the
// policy can no longer be used for agent behavior control and all references to it
// become invalid. This is an asynchronous operation. Use the GetPolicy operation
// to poll the status field to track completion.
func bedrockagentcorecontrol_DeletePolicy(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.DeletePolicyInput{
		// PolicyEngineId: *string, // Required
		// PolicyId: *string, // Required
	}

	if len(_bedrockagentcorecontrolPolicyEngineId) > 0 {
		input.PolicyEngineId = aws.String(_bedrockagentcorecontrolPolicyEngineId)
	}
	if len(_bedrockagentcorecontrolPolicyId) > 0 {
		input.PolicyId = aws.String(_bedrockagentcorecontrolPolicyId)
	}

	if resp, err := client.DeletePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing policy engine from the AgentCore Policy system. The policy
// engine must not have any associated policies before deletion. Once deleted, the
// policy engine and all its configurations become unavailable for policy
// management and evaluation. This is an asynchronous operation. Use the
// GetPolicyEngine operation to poll the status field to track completion.
func bedrockagentcorecontrol_DeletePolicyEngine(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.DeletePolicyEngineInput{
		// PolicyEngineId: *string, // Required
	}

	if len(_bedrockagentcorecontrolPolicyEngineId) > 0 {
		input.PolicyEngineId = aws.String(_bedrockagentcorecontrolPolicyEngineId)
	}

	if resp, err := client.DeletePolicyEngine(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the resource-based policy for a specified resource.
// This feature is currently available only for AgentCore Runtime and Gateway.
func bedrockagentcorecontrol_DeleteResourcePolicy(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.DeleteResourcePolicyInput{
		// ResourceArn: *string, // Required
	}

	if len(_bedrockagentcorecontrolResourceArn) > 0 {
		input.ResourceArn = aws.String(_bedrockagentcorecontrolResourceArn)
	}

	if resp, err := client.DeleteResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a workload identity.
func bedrockagentcorecontrol_DeleteWorkloadIdentity(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.DeleteWorkloadIdentityInput{
		// Name: *string, // Required
	}

	if len(_bedrockagentcorecontrolName) > 0 {
		input.Name = aws.String(_bedrockagentcorecontrolName)
	}

	if resp, err := client.DeleteWorkloadIdentity(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets an Amazon Bedrock AgentCore Runtime.
func bedrockagentcorecontrol_GetAgentRuntime(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.GetAgentRuntimeInput{
		// AgentRuntimeId: *string, // Required
	}

	if len(_bedrockagentcorecontrolAgentRuntimeId) > 0 {
		input.AgentRuntimeId = aws.String(_bedrockagentcorecontrolAgentRuntimeId)
	}
	if len(_bedrockagentcorecontrolAgentRuntimeVersion) > 0 {
		input.AgentRuntimeVersion = aws.String(_bedrockagentcorecontrolAgentRuntimeVersion)
	}

	if resp, err := client.GetAgentRuntime(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about an Amazon Secure AgentEndpoint.
func bedrockagentcorecontrol_GetAgentRuntimeEndpoint(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.GetAgentRuntimeEndpointInput{
		// AgentRuntimeId: *string, // Required
		// EndpointName: *string, // Required
	}

	if len(_bedrockagentcorecontrolAgentRuntimeId) > 0 {
		input.AgentRuntimeId = aws.String(_bedrockagentcorecontrolAgentRuntimeId)
	}
	if len(_bedrockagentcorecontrolEndpointName) > 0 {
		input.EndpointName = aws.String(_bedrockagentcorecontrolEndpointName)
	}

	if resp, err := client.GetAgentRuntimeEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about an API key credential provider.
func bedrockagentcorecontrol_GetApiKeyCredentialProvider(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.GetApiKeyCredentialProviderInput{
		// Name: *string, // Required
	}

	if len(_bedrockagentcorecontrolName) > 0 {
		input.Name = aws.String(_bedrockagentcorecontrolName)
	}

	if resp, err := client.GetApiKeyCredentialProvider(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a custom browser.
func bedrockagentcorecontrol_GetBrowser(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.GetBrowserInput{
		// BrowserId: *string, // Required
	}

	if len(_bedrockagentcorecontrolBrowserId) > 0 {
		input.BrowserId = aws.String(_bedrockagentcorecontrolBrowserId)
	}

	if resp, err := client.GetBrowser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a browser profile.
func bedrockagentcorecontrol_GetBrowserProfile(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.GetBrowserProfileInput{
		// ProfileId: *string, // Required
	}

	if len(_bedrockagentcorecontrolProfileId) > 0 {
		input.ProfileId = aws.String(_bedrockagentcorecontrolProfileId)
	}

	if resp, err := client.GetBrowserProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a custom code interpreter.
func bedrockagentcorecontrol_GetCodeInterpreter(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.GetCodeInterpreterInput{
		// CodeInterpreterId: *string, // Required
	}

	if len(_bedrockagentcorecontrolCodeInterpreterId) > 0 {
		input.CodeInterpreterId = aws.String(_bedrockagentcorecontrolCodeInterpreterId)
	}

	if resp, err := client.GetCodeInterpreter(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves detailed information about an evaluator, including its
// configuration, status, and metadata. Works with both built-in and custom
// evaluators.
func bedrockagentcorecontrol_GetEvaluator(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.GetEvaluatorInput{
		// EvaluatorId: *string, // Required
	}

	if len(_bedrockagentcorecontrolEvaluatorId) > 0 {
		input.EvaluatorId = aws.String(_bedrockagentcorecontrolEvaluatorId)
	}

	if resp, err := client.GetEvaluator(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a specific Gateway.
func bedrockagentcorecontrol_GetGateway(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.GetGatewayInput{
		// GatewayIdentifier: *string, // Required
	}

	if len(_bedrockagentcorecontrolGatewayIdentifier) > 0 {
		input.GatewayIdentifier = aws.String(_bedrockagentcorecontrolGatewayIdentifier)
	}

	if resp, err := client.GetGateway(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a specific gateway target.
func bedrockagentcorecontrol_GetGatewayTarget(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.GetGatewayTargetInput{
		// GatewayIdentifier: *string, // Required
		// TargetId: *string, // Required
	}

	if len(_bedrockagentcorecontrolGatewayIdentifier) > 0 {
		input.GatewayIdentifier = aws.String(_bedrockagentcorecontrolGatewayIdentifier)
	}
	if len(_bedrockagentcorecontrolTargetId) > 0 {
		input.TargetId = aws.String(_bedrockagentcorecontrolTargetId)
	}

	if resp, err := client.GetGatewayTarget(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieve an existing Amazon Bedrock AgentCore Memory resource.
func bedrockagentcorecontrol_GetMemory(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.GetMemoryInput{
		// MemoryId: *string, // Required
	}

	if len(_bedrockagentcorecontrolMemoryId) > 0 {
		input.MemoryId = aws.String(_bedrockagentcorecontrolMemoryId)
	}
	if len(_bedrockagentcorecontrolView) > 0 {
		if err := assignInputField(input, "View", _bedrockagentcorecontrolView); err != nil {
			log.Errorf("invalid --view: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetMemory(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about an OAuth2 credential provider.
func bedrockagentcorecontrol_GetOauth2CredentialProvider(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.GetOauth2CredentialProviderInput{
		// Name: *string, // Required
	}

	if len(_bedrockagentcorecontrolName) > 0 {
		input.Name = aws.String(_bedrockagentcorecontrolName)
	}

	if resp, err := client.GetOauth2CredentialProvider(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves detailed information about an online evaluation configuration,
// including its rules, data sources, evaluators, and execution status.
func bedrockagentcorecontrol_GetOnlineEvaluationConfig(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.GetOnlineEvaluationConfigInput{
		// OnlineEvaluationConfigId: *string, // Required
	}

	if len(_bedrockagentcorecontrolOnlineEvaluationConfigId) > 0 {
		input.OnlineEvaluationConfigId = aws.String(_bedrockagentcorecontrolOnlineEvaluationConfigId)
	}

	if resp, err := client.GetOnlineEvaluationConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves detailed information about a specific policy within the AgentCore
// Policy system. This operation returns the complete policy definition, metadata,
// and current status, allowing administrators to review and manage policy
// configurations.
func bedrockagentcorecontrol_GetPolicy(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.GetPolicyInput{
		// PolicyEngineId: *string, // Required
		// PolicyId: *string, // Required
	}

	if len(_bedrockagentcorecontrolPolicyEngineId) > 0 {
		input.PolicyEngineId = aws.String(_bedrockagentcorecontrolPolicyEngineId)
	}
	if len(_bedrockagentcorecontrolPolicyId) > 0 {
		input.PolicyId = aws.String(_bedrockagentcorecontrolPolicyId)
	}

	if resp, err := client.GetPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves detailed information about a specific policy engine within the
// AgentCore Policy system. This operation returns the complete policy engine
// configuration, metadata, and current status, allowing administrators to review
// and manage policy engine settings.
func bedrockagentcorecontrol_GetPolicyEngine(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.GetPolicyEngineInput{
		// PolicyEngineId: *string, // Required
	}

	if len(_bedrockagentcorecontrolPolicyEngineId) > 0 {
		input.PolicyEngineId = aws.String(_bedrockagentcorecontrolPolicyEngineId)
	}

	if resp, err := client.GetPolicyEngine(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a policy generation request within the AgentCore
// Policy system. Policy generation converts natural language descriptions into
// Cedar policy statements using AI-powered translation, enabling non-technical
// users to create policies.
func bedrockagentcorecontrol_GetPolicyGeneration(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.GetPolicyGenerationInput{
		// PolicyEngineId: *string, // Required
		// PolicyGenerationId: *string, // Required
	}

	if len(_bedrockagentcorecontrolPolicyEngineId) > 0 {
		input.PolicyEngineId = aws.String(_bedrockagentcorecontrolPolicyEngineId)
	}
	if len(_bedrockagentcorecontrolPolicyGenerationId) > 0 {
		input.PolicyGenerationId = aws.String(_bedrockagentcorecontrolPolicyGenerationId)
	}

	if resp, err := client.GetPolicyGeneration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the resource-based policy for a specified resource.
// This feature is currently available only for AgentCore Runtime and Gateway.
func bedrockagentcorecontrol_GetResourcePolicy(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.GetResourcePolicyInput{
		// ResourceArn: *string, // Required
	}

	if len(_bedrockagentcorecontrolResourceArn) > 0 {
		input.ResourceArn = aws.String(_bedrockagentcorecontrolResourceArn)
	}

	if resp, err := client.GetResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a token vault.
func bedrockagentcorecontrol_GetTokenVault(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.GetTokenVaultInput{}

	if len(_bedrockagentcorecontrolTokenVaultId) > 0 {
		input.TokenVaultId = aws.String(_bedrockagentcorecontrolTokenVaultId)
	}

	if resp, err := client.GetTokenVault(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a workload identity.
func bedrockagentcorecontrol_GetWorkloadIdentity(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.GetWorkloadIdentityInput{
		// Name: *string, // Required
	}

	if len(_bedrockagentcorecontrolName) > 0 {
		input.Name = aws.String(_bedrockagentcorecontrolName)
	}

	if resp, err := client.GetWorkloadIdentity(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all endpoints for a specific Amazon Secure Agent.
func bedrockagentcorecontrol_ListAgentRuntimeEndpoints(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.ListAgentRuntimeEndpointsInput{
		// AgentRuntimeId: *string, // Required
	}

	if len(_bedrockagentcorecontrolAgentRuntimeId) > 0 {
		input.AgentRuntimeId = aws.String(_bedrockagentcorecontrolAgentRuntimeId)
	}
	if len(_bedrockagentcorecontrolMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bedrockagentcorecontrolMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolNextToken) > 0 {
		input.NextToken = aws.String(_bedrockagentcorecontrolNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAgentRuntimeEndpoints(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrockagentcorecontrol.ListAgentRuntimeEndpointsOutput
	p := bedrockagentcorecontrol.NewListAgentRuntimeEndpointsPaginator(client, input)
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

// Lists all versions of a specific Amazon Secure Agent.
func bedrockagentcorecontrol_ListAgentRuntimeVersions(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.ListAgentRuntimeVersionsInput{
		// AgentRuntimeId: *string, // Required
	}

	if len(_bedrockagentcorecontrolAgentRuntimeId) > 0 {
		input.AgentRuntimeId = aws.String(_bedrockagentcorecontrolAgentRuntimeId)
	}
	if len(_bedrockagentcorecontrolMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bedrockagentcorecontrolMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolNextToken) > 0 {
		input.NextToken = aws.String(_bedrockagentcorecontrolNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAgentRuntimeVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrockagentcorecontrol.ListAgentRuntimeVersionsOutput
	p := bedrockagentcorecontrol.NewListAgentRuntimeVersionsPaginator(client, input)
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

// Lists all Amazon Secure Agents in your account.
func bedrockagentcorecontrol_ListAgentRuntimes(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.ListAgentRuntimesInput{}

	if len(_bedrockagentcorecontrolMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bedrockagentcorecontrolMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolNextToken) > 0 {
		input.NextToken = aws.String(_bedrockagentcorecontrolNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAgentRuntimes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrockagentcorecontrol.ListAgentRuntimesOutput
	p := bedrockagentcorecontrol.NewListAgentRuntimesPaginator(client, input)
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

// Lists all API key credential providers in your account.
func bedrockagentcorecontrol_ListApiKeyCredentialProviders(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.ListApiKeyCredentialProvidersInput{}

	if len(_bedrockagentcorecontrolMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bedrockagentcorecontrolMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolNextToken) > 0 {
		input.NextToken = aws.String(_bedrockagentcorecontrolNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListApiKeyCredentialProviders(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrockagentcorecontrol.ListApiKeyCredentialProvidersOutput
	p := bedrockagentcorecontrol.NewListApiKeyCredentialProvidersPaginator(client, input)
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

// Lists all browser profiles in your account.
func bedrockagentcorecontrol_ListBrowserProfiles(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.ListBrowserProfilesInput{}

	if len(_bedrockagentcorecontrolMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bedrockagentcorecontrolMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolNextToken) > 0 {
		input.NextToken = aws.String(_bedrockagentcorecontrolNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListBrowserProfiles(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrockagentcorecontrol.ListBrowserProfilesOutput
	p := bedrockagentcorecontrol.NewListBrowserProfilesPaginator(client, input)
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

// Lists all custom browsers in your account.
func bedrockagentcorecontrol_ListBrowsers(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.ListBrowsersInput{}

	if len(_bedrockagentcorecontrolMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bedrockagentcorecontrolMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolNextToken) > 0 {
		input.NextToken = aws.String(_bedrockagentcorecontrolNextToken)
	}
	if len(_bedrockagentcorecontrolType) > 0 {
		if err := assignInputField(input, "Type", _bedrockagentcorecontrolType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListBrowsers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrockagentcorecontrol.ListBrowsersOutput
	p := bedrockagentcorecontrol.NewListBrowsersPaginator(client, input)
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

// Lists all custom code interpreters in your account.
func bedrockagentcorecontrol_ListCodeInterpreters(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.ListCodeInterpretersInput{}

	if len(_bedrockagentcorecontrolMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bedrockagentcorecontrolMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolNextToken) > 0 {
		input.NextToken = aws.String(_bedrockagentcorecontrolNextToken)
	}
	if len(_bedrockagentcorecontrolType) > 0 {
		if err := assignInputField(input, "Type", _bedrockagentcorecontrolType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListCodeInterpreters(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrockagentcorecontrol.ListCodeInterpretersOutput
	p := bedrockagentcorecontrol.NewListCodeInterpretersPaginator(client, input)
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

// Lists all available evaluators, including both builtin evaluators provided by
// the service and custom evaluators created by the user.
func bedrockagentcorecontrol_ListEvaluators(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.ListEvaluatorsInput{}

	if len(_bedrockagentcorecontrolMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bedrockagentcorecontrolMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolNextToken) > 0 {
		input.NextToken = aws.String(_bedrockagentcorecontrolNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListEvaluators(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrockagentcorecontrol.ListEvaluatorsOutput
	p := bedrockagentcorecontrol.NewListEvaluatorsPaginator(client, input)
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

// Lists all targets for a specific gateway.
func bedrockagentcorecontrol_ListGatewayTargets(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.ListGatewayTargetsInput{
		// GatewayIdentifier: *string, // Required
	}

	if len(_bedrockagentcorecontrolGatewayIdentifier) > 0 {
		input.GatewayIdentifier = aws.String(_bedrockagentcorecontrolGatewayIdentifier)
	}
	if len(_bedrockagentcorecontrolMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bedrockagentcorecontrolMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolNextToken) > 0 {
		input.NextToken = aws.String(_bedrockagentcorecontrolNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListGatewayTargets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrockagentcorecontrol.ListGatewayTargetsOutput
	p := bedrockagentcorecontrol.NewListGatewayTargetsPaginator(client, input)
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

// Lists all gateways in the account.
func bedrockagentcorecontrol_ListGateways(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.ListGatewaysInput{}

	if len(_bedrockagentcorecontrolMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bedrockagentcorecontrolMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolNextToken) > 0 {
		input.NextToken = aws.String(_bedrockagentcorecontrolNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListGateways(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrockagentcorecontrol.ListGatewaysOutput
	p := bedrockagentcorecontrol.NewListGatewaysPaginator(client, input)
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

// Lists the available Amazon Bedrock AgentCore Memory resources in the current
// Amazon Web Services Region.
func bedrockagentcorecontrol_ListMemories(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.ListMemoriesInput{}

	if len(_bedrockagentcorecontrolMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bedrockagentcorecontrolMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolNextToken) > 0 {
		input.NextToken = aws.String(_bedrockagentcorecontrolNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListMemories(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrockagentcorecontrol.ListMemoriesOutput
	p := bedrockagentcorecontrol.NewListMemoriesPaginator(client, input)
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

// Lists all OAuth2 credential providers in your account.
func bedrockagentcorecontrol_ListOauth2CredentialProviders(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.ListOauth2CredentialProvidersInput{}

	if len(_bedrockagentcorecontrolMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bedrockagentcorecontrolMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolNextToken) > 0 {
		input.NextToken = aws.String(_bedrockagentcorecontrolNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListOauth2CredentialProviders(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrockagentcorecontrol.ListOauth2CredentialProvidersOutput
	p := bedrockagentcorecontrol.NewListOauth2CredentialProvidersPaginator(client, input)
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

// Lists all online evaluation configurations in the account, providing summary
// information about each configuration's status and settings.
func bedrockagentcorecontrol_ListOnlineEvaluationConfigs(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.ListOnlineEvaluationConfigsInput{}

	if len(_bedrockagentcorecontrolMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bedrockagentcorecontrolMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolNextToken) > 0 {
		input.NextToken = aws.String(_bedrockagentcorecontrolNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListOnlineEvaluationConfigs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrockagentcorecontrol.ListOnlineEvaluationConfigsOutput
	p := bedrockagentcorecontrol.NewListOnlineEvaluationConfigsPaginator(client, input)
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

// Retrieves a list of policies within the AgentCore Policy engine. This operation
// supports pagination and filtering to help administrators manage and discover
// policies across policy engines. Results can be filtered by policy engine or
// resource associations.
func bedrockagentcorecontrol_ListPolicies(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.ListPoliciesInput{
		// PolicyEngineId: *string, // Required
	}

	if len(_bedrockagentcorecontrolPolicyEngineId) > 0 {
		input.PolicyEngineId = aws.String(_bedrockagentcorecontrolPolicyEngineId)
	}
	if len(_bedrockagentcorecontrolMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bedrockagentcorecontrolMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolNextToken) > 0 {
		input.NextToken = aws.String(_bedrockagentcorecontrolNextToken)
	}
	if len(_bedrockagentcorecontrolTargetResourceScope) > 0 {
		input.TargetResourceScope = aws.String(_bedrockagentcorecontrolTargetResourceScope)
	}

	if disablePaginator() {
		if resp, err := client.ListPolicies(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrockagentcorecontrol.ListPoliciesOutput
	p := bedrockagentcorecontrol.NewListPoliciesPaginator(client, input)
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

// Retrieves a list of policy engines within the AgentCore Policy system. This
// operation supports pagination to help administrators discover and manage policy
// engines across their account. Each policy engine serves as a container for
// related policies.
func bedrockagentcorecontrol_ListPolicyEngines(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.ListPolicyEnginesInput{}

	if len(_bedrockagentcorecontrolMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bedrockagentcorecontrolMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolNextToken) > 0 {
		input.NextToken = aws.String(_bedrockagentcorecontrolNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPolicyEngines(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrockagentcorecontrol.ListPolicyEnginesOutput
	p := bedrockagentcorecontrol.NewListPolicyEnginesPaginator(client, input)
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

// Retrieves a list of generated policy assets from a policy generation request
// within the AgentCore Policy system. This operation returns the actual Cedar
// policies and related artifacts produced by the AI-powered policy generation
// process, allowing users to review and select from multiple generated policy
// options.
func bedrockagentcorecontrol_ListPolicyGenerationAssets(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.ListPolicyGenerationAssetsInput{
		// PolicyEngineId: *string, // Required
		// PolicyGenerationId: *string, // Required
	}

	if len(_bedrockagentcorecontrolPolicyEngineId) > 0 {
		input.PolicyEngineId = aws.String(_bedrockagentcorecontrolPolicyEngineId)
	}
	if len(_bedrockagentcorecontrolPolicyGenerationId) > 0 {
		input.PolicyGenerationId = aws.String(_bedrockagentcorecontrolPolicyGenerationId)
	}
	if len(_bedrockagentcorecontrolMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bedrockagentcorecontrolMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolNextToken) > 0 {
		input.NextToken = aws.String(_bedrockagentcorecontrolNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPolicyGenerationAssets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrockagentcorecontrol.ListPolicyGenerationAssetsOutput
	p := bedrockagentcorecontrol.NewListPolicyGenerationAssetsPaginator(client, input)
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

// Retrieves a list of policy generation requests within the AgentCore Policy
// system. This operation supports pagination and filtering to help track and
// manage AI-powered policy generation operations.
func bedrockagentcorecontrol_ListPolicyGenerations(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.ListPolicyGenerationsInput{
		// PolicyEngineId: *string, // Required
	}

	if len(_bedrockagentcorecontrolPolicyEngineId) > 0 {
		input.PolicyEngineId = aws.String(_bedrockagentcorecontrolPolicyEngineId)
	}
	if len(_bedrockagentcorecontrolMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bedrockagentcorecontrolMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolNextToken) > 0 {
		input.NextToken = aws.String(_bedrockagentcorecontrolNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPolicyGenerations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrockagentcorecontrol.ListPolicyGenerationsOutput
	p := bedrockagentcorecontrol.NewListPolicyGenerationsPaginator(client, input)
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

// Lists the tags associated with the specified resource.
// This feature is currently available only for AgentCore Runtime, Browser,
// Browser Profile, Code Interpreter tool, and Gateway.
func bedrockagentcorecontrol_ListTagsForResource(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_bedrockagentcorecontrolResourceArn) > 0 {
		input.ResourceArn = aws.String(_bedrockagentcorecontrolResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all workload identities in your account.
func bedrockagentcorecontrol_ListWorkloadIdentities(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.ListWorkloadIdentitiesInput{}

	if len(_bedrockagentcorecontrolMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bedrockagentcorecontrolMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolNextToken) > 0 {
		input.NextToken = aws.String(_bedrockagentcorecontrolNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListWorkloadIdentities(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrockagentcorecontrol.ListWorkloadIdentitiesOutput
	p := bedrockagentcorecontrol.NewListWorkloadIdentitiesPaginator(client, input)
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

// Creates or updates a resource-based policy for a resource with the specified
// resourceArn.
//
// This feature is currently available only for AgentCore Runtime and Gateway.
func bedrockagentcorecontrol_PutResourcePolicy(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.PutResourcePolicyInput{
		// Policy: *string, // Required
		// ResourceArn: *string, // Required
	}

	if len(_bedrockagentcorecontrolPolicy) > 0 {
		input.Policy = aws.String(_bedrockagentcorecontrolPolicy)
	}
	if len(_bedrockagentcorecontrolResourceArn) > 0 {
		input.ResourceArn = aws.String(_bedrockagentcorecontrolResourceArn)
	}

	if resp, err := client.PutResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets the customer master key (CMK) for a token vault.
func bedrockagentcorecontrol_SetTokenVaultCMK(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.SetTokenVaultCMKInput{
		// KmsConfiguration: *types.KmsConfiguration, // Required
	}

	if len(_bedrockagentcorecontrolKmsConfiguration) > 0 {
		if err := assignInputField(input, "KmsConfiguration", _bedrockagentcorecontrolKmsConfiguration); err != nil {
			log.Errorf("invalid --kms-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolTokenVaultId) > 0 {
		input.TokenVaultId = aws.String(_bedrockagentcorecontrolTokenVaultId)
	}

	if resp, err := client.SetTokenVaultCMK(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Initiates the AI-powered generation of Cedar policies from natural language
// descriptions within the AgentCore Policy system. This feature enables both
// technical and non-technical users to create policies by describing their
// authorization requirements in plain English, which is then automatically
// translated into formal Cedar policy statements. The generation process analyzes
// the natural language input along with the Gateway's tool context to produce
// validated policy options. Generated policy assets are automatically deleted
// after 7 days, so you should review and create policies from the generated assets
// within this timeframe. Once created, policies are permanent and not subject to
// this expiration. Generated policies should be reviewed and tested in log-only
// mode before deploying to production. Use this when you want to describe policy
// intent naturally rather than learning Cedar syntax, though generated policies
// may require refinement for complex scenarios.
func bedrockagentcorecontrol_StartPolicyGeneration(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.StartPolicyGenerationInput{
		// Content: types.Content, // Required
		// Name: *string, // Required
		// PolicyEngineId: *string, // Required
		// Resource: types.Resource, // Required
	}

	if len(_bedrockagentcorecontrolContent) > 0 {
		if err := assignInputField(input, "Content", _bedrockagentcorecontrolContent); err != nil {
			log.Errorf("invalid --content: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolName) > 0 {
		input.Name = aws.String(_bedrockagentcorecontrolName)
	}
	if len(_bedrockagentcorecontrolPolicyEngineId) > 0 {
		input.PolicyEngineId = aws.String(_bedrockagentcorecontrolPolicyEngineId)
	}
	if len(_bedrockagentcorecontrolResource) > 0 {
		if err := assignInputField(input, "Resource", _bedrockagentcorecontrolResource); err != nil {
			log.Errorf("invalid --resource: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolClientToken) > 0 {
		input.ClientToken = aws.String(_bedrockagentcorecontrolClientToken)
	}

	if resp, err := client.StartPolicyGeneration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The gateway targets.
func bedrockagentcorecontrol_SynchronizeGatewayTargets(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.SynchronizeGatewayTargetsInput{
		// GatewayIdentifier: *string, // Required
		// TargetIdList: []string, // Required
	}

	if len(_bedrockagentcorecontrolGatewayIdentifier) > 0 {
		input.GatewayIdentifier = aws.String(_bedrockagentcorecontrolGatewayIdentifier)
	}
	if len(_bedrockagentcorecontrolTargetIdList) > 0 {
		input.TargetIdList = append([]string(nil), _bedrockagentcorecontrolTargetIdList...)
	}

	if resp, err := client.SynchronizeGatewayTargets(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates the specified tags to a resource with the specified resourceArn. If
// existing tags on a resource are not specified in the request parameters, they
// are not changed. When a resource is deleted, the tags associated with that
// resource are also deleted.
//
// This feature is currently available only for AgentCore Runtime, Browser,
// Browser Profile, Code Interpreter tool, and Gateway.
func bedrockagentcorecontrol_TagResource(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_bedrockagentcorecontrolResourceArn) > 0 {
		input.ResourceArn = aws.String(_bedrockagentcorecontrolResourceArn)
	}
	if len(_bedrockagentcorecontrolTags) > 0 {
		if err := assignInputField(input, "Tags", _bedrockagentcorecontrolTags); err != nil {
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

// Removes the specified tags from the specified resource.
// This feature is currently available only for AgentCore Runtime, Browser,
// Browser Profile, Code Interpreter tool, and Gateway.
func bedrockagentcorecontrol_UntagResource(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_bedrockagentcorecontrolResourceArn) > 0 {
		input.ResourceArn = aws.String(_bedrockagentcorecontrolResourceArn)
	}
	if len(_bedrockagentcorecontrolTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _bedrockagentcorecontrolTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing Amazon Secure Agent.
func bedrockagentcorecontrol_UpdateAgentRuntime(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.UpdateAgentRuntimeInput{
		// AgentRuntimeArtifact: types.AgentRuntimeArtifact, // Required
		// AgentRuntimeId: *string, // Required
		// NetworkConfiguration: *types.NetworkConfiguration, // Required
		// RoleArn: *string, // Required
	}

	if len(_bedrockagentcorecontrolAgentRuntimeArtifact) > 0 {
		if err := assignInputField(input, "AgentRuntimeArtifact", _bedrockagentcorecontrolAgentRuntimeArtifact); err != nil {
			log.Errorf("invalid --agent-runtime-artifact: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolAgentRuntimeId) > 0 {
		input.AgentRuntimeId = aws.String(_bedrockagentcorecontrolAgentRuntimeId)
	}
	if len(_bedrockagentcorecontrolNetworkConfiguration) > 0 {
		if err := assignInputField(input, "NetworkConfiguration", _bedrockagentcorecontrolNetworkConfiguration); err != nil {
			log.Errorf("invalid --network-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolRoleArn) > 0 {
		input.RoleArn = aws.String(_bedrockagentcorecontrolRoleArn)
	}
	if len(_bedrockagentcorecontrolAuthorizerConfiguration) > 0 {
		if err := assignInputField(input, "AuthorizerConfiguration", _bedrockagentcorecontrolAuthorizerConfiguration); err != nil {
			log.Errorf("invalid --authorizer-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolClientToken) > 0 {
		input.ClientToken = aws.String(_bedrockagentcorecontrolClientToken)
	}
	if len(_bedrockagentcorecontrolDescription) > 0 {
		input.Description = aws.String(_bedrockagentcorecontrolDescription)
	}
	if len(_bedrockagentcorecontrolEnvironmentVariables) > 0 {
		if err := assignInputField(input, "EnvironmentVariables", _bedrockagentcorecontrolEnvironmentVariables); err != nil {
			log.Errorf("invalid --environment-variables: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolLifecycleConfiguration) > 0 {
		if err := assignInputField(input, "LifecycleConfiguration", _bedrockagentcorecontrolLifecycleConfiguration); err != nil {
			log.Errorf("invalid --lifecycle-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolMetadataConfiguration) > 0 {
		if err := assignInputField(input, "MetadataConfiguration", _bedrockagentcorecontrolMetadataConfiguration); err != nil {
			log.Errorf("invalid --metadata-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolProtocolConfiguration) > 0 {
		if err := assignInputField(input, "ProtocolConfiguration", _bedrockagentcorecontrolProtocolConfiguration); err != nil {
			log.Errorf("invalid --protocol-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolRequestHeaderConfiguration) > 0 {
		if err := assignInputField(input, "RequestHeaderConfiguration", _bedrockagentcorecontrolRequestHeaderConfiguration); err != nil {
			log.Errorf("invalid --request-header-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateAgentRuntime(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing Amazon Bedrock AgentCore Runtime endpoint.
func bedrockagentcorecontrol_UpdateAgentRuntimeEndpoint(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.UpdateAgentRuntimeEndpointInput{
		// AgentRuntimeId: *string, // Required
		// EndpointName: *string, // Required
	}

	if len(_bedrockagentcorecontrolAgentRuntimeId) > 0 {
		input.AgentRuntimeId = aws.String(_bedrockagentcorecontrolAgentRuntimeId)
	}
	if len(_bedrockagentcorecontrolEndpointName) > 0 {
		input.EndpointName = aws.String(_bedrockagentcorecontrolEndpointName)
	}
	if len(_bedrockagentcorecontrolAgentRuntimeVersion) > 0 {
		input.AgentRuntimeVersion = aws.String(_bedrockagentcorecontrolAgentRuntimeVersion)
	}
	if len(_bedrockagentcorecontrolClientToken) > 0 {
		input.ClientToken = aws.String(_bedrockagentcorecontrolClientToken)
	}
	if len(_bedrockagentcorecontrolDescription) > 0 {
		input.Description = aws.String(_bedrockagentcorecontrolDescription)
	}

	if resp, err := client.UpdateAgentRuntimeEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing API key credential provider.
func bedrockagentcorecontrol_UpdateApiKeyCredentialProvider(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.UpdateApiKeyCredentialProviderInput{
		// ApiKey: *string, // Required
		// Name: *string, // Required
	}

	if len(_bedrockagentcorecontrolApiKey) > 0 {
		input.ApiKey = aws.String(_bedrockagentcorecontrolApiKey)
	}
	if len(_bedrockagentcorecontrolName) > 0 {
		input.Name = aws.String(_bedrockagentcorecontrolName)
	}

	if resp, err := client.UpdateApiKeyCredentialProvider(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a custom evaluator's configuration, description, or evaluation level.
// Built-in evaluators cannot be updated. The evaluator must not be locked for
// modification.
func bedrockagentcorecontrol_UpdateEvaluator(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.UpdateEvaluatorInput{
		// EvaluatorId: *string, // Required
	}

	if len(_bedrockagentcorecontrolEvaluatorId) > 0 {
		input.EvaluatorId = aws.String(_bedrockagentcorecontrolEvaluatorId)
	}
	if len(_bedrockagentcorecontrolClientToken) > 0 {
		input.ClientToken = aws.String(_bedrockagentcorecontrolClientToken)
	}
	if len(_bedrockagentcorecontrolDescription) > 0 {
		input.Description = aws.String(_bedrockagentcorecontrolDescription)
	}
	if len(_bedrockagentcorecontrolEvaluatorConfig) > 0 {
		if err := assignInputField(input, "EvaluatorConfig", _bedrockagentcorecontrolEvaluatorConfig); err != nil {
			log.Errorf("invalid --evaluator-config: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolLevel) > 0 {
		if err := assignInputField(input, "Level", _bedrockagentcorecontrolLevel); err != nil {
			log.Errorf("invalid --level: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateEvaluator(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing gateway.
func bedrockagentcorecontrol_UpdateGateway(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.UpdateGatewayInput{
		// AuthorizerType: types.AuthorizerType, // Required
		// GatewayIdentifier: *string, // Required
		// Name: *string, // Required
		// ProtocolType: types.GatewayProtocolType, // Required
		// RoleArn: *string, // Required
	}

	if len(_bedrockagentcorecontrolAuthorizerType) > 0 {
		if err := assignInputField(input, "AuthorizerType", _bedrockagentcorecontrolAuthorizerType); err != nil {
			log.Errorf("invalid --authorizer-type: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolGatewayIdentifier) > 0 {
		input.GatewayIdentifier = aws.String(_bedrockagentcorecontrolGatewayIdentifier)
	}
	if len(_bedrockagentcorecontrolName) > 0 {
		input.Name = aws.String(_bedrockagentcorecontrolName)
	}
	if len(_bedrockagentcorecontrolProtocolType) > 0 {
		if err := assignInputField(input, "ProtocolType", _bedrockagentcorecontrolProtocolType); err != nil {
			log.Errorf("invalid --protocol-type: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolRoleArn) > 0 {
		input.RoleArn = aws.String(_bedrockagentcorecontrolRoleArn)
	}
	if len(_bedrockagentcorecontrolAuthorizerConfiguration) > 0 {
		if err := assignInputField(input, "AuthorizerConfiguration", _bedrockagentcorecontrolAuthorizerConfiguration); err != nil {
			log.Errorf("invalid --authorizer-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolDescription) > 0 {
		input.Description = aws.String(_bedrockagentcorecontrolDescription)
	}
	if len(_bedrockagentcorecontrolExceptionLevel) > 0 {
		if err := assignInputField(input, "ExceptionLevel", _bedrockagentcorecontrolExceptionLevel); err != nil {
			log.Errorf("invalid --exception-level: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolInterceptorConfigurations) > 0 {
		if err := assignInputField(input, "InterceptorConfigurations", _bedrockagentcorecontrolInterceptorConfigurations); err != nil {
			log.Errorf("invalid --interceptor-configurations: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolKmsKeyArn) > 0 {
		input.KmsKeyArn = aws.String(_bedrockagentcorecontrolKmsKeyArn)
	}
	if len(_bedrockagentcorecontrolPolicyEngineConfiguration) > 0 {
		if err := assignInputField(input, "PolicyEngineConfiguration", _bedrockagentcorecontrolPolicyEngineConfiguration); err != nil {
			log.Errorf("invalid --policy-engine-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolProtocolConfiguration) > 0 {
		if err := assignInputField(input, "ProtocolConfiguration", _bedrockagentcorecontrolProtocolConfiguration); err != nil {
			log.Errorf("invalid --protocol-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateGateway(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing gateway target.
func bedrockagentcorecontrol_UpdateGatewayTarget(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.UpdateGatewayTargetInput{
		// GatewayIdentifier: *string, // Required
		// Name: *string, // Required
		// TargetConfiguration: types.TargetConfiguration, // Required
		// TargetId: *string, // Required
	}

	if len(_bedrockagentcorecontrolGatewayIdentifier) > 0 {
		input.GatewayIdentifier = aws.String(_bedrockagentcorecontrolGatewayIdentifier)
	}
	if len(_bedrockagentcorecontrolName) > 0 {
		input.Name = aws.String(_bedrockagentcorecontrolName)
	}
	if len(_bedrockagentcorecontrolTargetConfiguration) > 0 {
		if err := assignInputField(input, "TargetConfiguration", _bedrockagentcorecontrolTargetConfiguration); err != nil {
			log.Errorf("invalid --target-configuration: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolTargetId) > 0 {
		input.TargetId = aws.String(_bedrockagentcorecontrolTargetId)
	}
	if len(_bedrockagentcorecontrolCredentialProviderConfigurations) > 0 {
		if err := assignInputField(input, "CredentialProviderConfigurations", _bedrockagentcorecontrolCredentialProviderConfigurations); err != nil {
			log.Errorf("invalid --credential-provider-configurations: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolDescription) > 0 {
		input.Description = aws.String(_bedrockagentcorecontrolDescription)
	}
	if len(_bedrockagentcorecontrolMetadataConfiguration) > 0 {
		if err := assignInputField(input, "MetadataConfiguration", _bedrockagentcorecontrolMetadataConfiguration); err != nil {
			log.Errorf("invalid --metadata-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateGatewayTarget(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update an Amazon Bedrock AgentCore Memory resource memory.
func bedrockagentcorecontrol_UpdateMemory(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.UpdateMemoryInput{
		// MemoryId: *string, // Required
	}

	if len(_bedrockagentcorecontrolMemoryId) > 0 {
		input.MemoryId = aws.String(_bedrockagentcorecontrolMemoryId)
	}
	if len(_bedrockagentcorecontrolClientToken) > 0 {
		input.ClientToken = aws.String(_bedrockagentcorecontrolClientToken)
	}
	if len(_bedrockagentcorecontrolDescription) > 0 {
		input.Description = aws.String(_bedrockagentcorecontrolDescription)
	}
	if len(_bedrockagentcorecontrolEventExpiryDuration) > 0 {
		if err := assignInputField(input, "EventExpiryDuration", _bedrockagentcorecontrolEventExpiryDuration); err != nil {
			log.Errorf("invalid --event-expiry-duration: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolMemoryExecutionRoleArn) > 0 {
		input.MemoryExecutionRoleArn = aws.String(_bedrockagentcorecontrolMemoryExecutionRoleArn)
	}
	if len(_bedrockagentcorecontrolMemoryStrategies) > 0 {
		if err := assignInputField(input, "MemoryStrategies", _bedrockagentcorecontrolMemoryStrategies); err != nil {
			log.Errorf("invalid --memory-strategies: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateMemory(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing OAuth2 credential provider.
func bedrockagentcorecontrol_UpdateOauth2CredentialProvider(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.UpdateOauth2CredentialProviderInput{
		// CredentialProviderVendor: types.CredentialProviderVendorType, // Required
		// Name: *string, // Required
		// Oauth2ProviderConfigInput: types.Oauth2ProviderConfigInput, // Required
	}

	if len(_bedrockagentcorecontrolCredentialProviderVendor) > 0 {
		if err := assignInputField(input, "CredentialProviderVendor", _bedrockagentcorecontrolCredentialProviderVendor); err != nil {
			log.Errorf("invalid --credential-provider-vendor: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolName) > 0 {
		input.Name = aws.String(_bedrockagentcorecontrolName)
	}
	if len(_bedrockagentcorecontrolOauth2ProviderConfigInput) > 0 {
		if err := assignInputField(input, "Oauth2ProviderConfigInput", _bedrockagentcorecontrolOauth2ProviderConfigInput); err != nil {
			log.Errorf("invalid --oauth2-provider-config-input: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateOauth2CredentialProvider(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an online evaluation configuration's settings, including rules, data
// sources, evaluators, and execution status. Changes take effect immediately for
// ongoing evaluations.
func bedrockagentcorecontrol_UpdateOnlineEvaluationConfig(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.UpdateOnlineEvaluationConfigInput{
		// OnlineEvaluationConfigId: *string, // Required
	}

	if len(_bedrockagentcorecontrolOnlineEvaluationConfigId) > 0 {
		input.OnlineEvaluationConfigId = aws.String(_bedrockagentcorecontrolOnlineEvaluationConfigId)
	}
	if len(_bedrockagentcorecontrolClientToken) > 0 {
		input.ClientToken = aws.String(_bedrockagentcorecontrolClientToken)
	}
	if len(_bedrockagentcorecontrolDataSourceConfig) > 0 {
		if err := assignInputField(input, "DataSourceConfig", _bedrockagentcorecontrolDataSourceConfig); err != nil {
			log.Errorf("invalid --data-source-config: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolDescription) > 0 {
		input.Description = aws.String(_bedrockagentcorecontrolDescription)
	}
	if len(_bedrockagentcorecontrolEvaluationExecutionRoleArn) > 0 {
		input.EvaluationExecutionRoleArn = aws.String(_bedrockagentcorecontrolEvaluationExecutionRoleArn)
	}
	if len(_bedrockagentcorecontrolEvaluators) > 0 {
		if err := assignInputField(input, "Evaluators", _bedrockagentcorecontrolEvaluators); err != nil {
			log.Errorf("invalid --evaluators: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolExecutionStatus) > 0 {
		if err := assignInputField(input, "ExecutionStatus", _bedrockagentcorecontrolExecutionStatus); err != nil {
			log.Errorf("invalid --execution-status: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolRule) > 0 {
		if err := assignInputField(input, "Rule", _bedrockagentcorecontrolRule); err != nil {
			log.Errorf("invalid --rule: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateOnlineEvaluationConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing policy within the AgentCore Policy system. This operation
// allows modification of the policy description and definition while maintaining
// the policy's identity. The updated policy is validated against the Cedar schema
// before being applied. This is an asynchronous operation. Use the GetPolicy
// operation to poll the status field to track completion.
func bedrockagentcorecontrol_UpdatePolicy(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.UpdatePolicyInput{
		// PolicyEngineId: *string, // Required
		// PolicyId: *string, // Required
	}

	if len(_bedrockagentcorecontrolPolicyEngineId) > 0 {
		input.PolicyEngineId = aws.String(_bedrockagentcorecontrolPolicyEngineId)
	}
	if len(_bedrockagentcorecontrolPolicyId) > 0 {
		input.PolicyId = aws.String(_bedrockagentcorecontrolPolicyId)
	}
	if len(_bedrockagentcorecontrolDefinition) > 0 {
		if err := assignInputField(input, "Definition", _bedrockagentcorecontrolDefinition); err != nil {
			log.Errorf("invalid --definition: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolDescription) > 0 {
		if err := assignInputField(input, "Description", _bedrockagentcorecontrolDescription); err != nil {
			log.Errorf("invalid --description: %s", err.Error())
			return
		}
	}
	if len(_bedrockagentcorecontrolValidationMode) > 0 {
		if err := assignInputField(input, "ValidationMode", _bedrockagentcorecontrolValidationMode); err != nil {
			log.Errorf("invalid --validation-mode: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdatePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing policy engine within the AgentCore Policy system. This
// operation allows modification of the policy engine description while maintaining
// its identity. This is an asynchronous operation. Use the GetPolicyEngine
// operation to poll the status field to track completion.
func bedrockagentcorecontrol_UpdatePolicyEngine(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.UpdatePolicyEngineInput{
		// PolicyEngineId: *string, // Required
	}

	if len(_bedrockagentcorecontrolPolicyEngineId) > 0 {
		input.PolicyEngineId = aws.String(_bedrockagentcorecontrolPolicyEngineId)
	}
	if len(_bedrockagentcorecontrolDescription) > 0 {
		if err := assignInputField(input, "Description", _bedrockagentcorecontrolDescription); err != nil {
			log.Errorf("invalid --description: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdatePolicyEngine(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing workload identity.
func bedrockagentcorecontrol_UpdateWorkloadIdentity(cfg aws.Config, client *bedrockagentcorecontrol.Client) {
	input := &bedrockagentcorecontrol.UpdateWorkloadIdentityInput{
		// Name: *string, // Required
	}

	if len(_bedrockagentcorecontrolName) > 0 {
		input.Name = aws.String(_bedrockagentcorecontrolName)
	}
	if len(_bedrockagentcorecontrolAllowedResourceOauth2ReturnUrls) > 0 {
		input.AllowedResourceOauth2ReturnUrls = append([]string(nil), _bedrockagentcorecontrolAllowedResourceOauth2ReturnUrls...)
	}

	if resp, err := client.UpdateWorkloadIdentity(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_bedrockagentcorecontrolCmd)
	_bedrockagentcorecontrolCmd.Flags().SortFlags = false

	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolAgentRuntimeArtifact, "agent-runtime-artifact", "", "", "Agent Runtime Artifact")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolAgentRuntimeId, "agent-runtime-id", "", "", "Agent Runtime ID")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolAgentRuntimeName, "agent-runtime-name", "", "", "Agent Runtime Name")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolAgentRuntimeVersion, "agent-runtime-version", "", "", "Agent Runtime Version")
	_bedrockagentcorecontrolCmd.Flags().StringSliceVarP(&_bedrockagentcorecontrolAllowedResourceOauth2ReturnUrls, "allowed-resource-oauth2-return-urls", "", nil, "Allowed Resource Oauth2 Return Urls")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolApiKey, "api-key", "", "", "API Key")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolAuthorizerConfiguration, "authorizer-configuration", "", "", "Authorizer Configuration")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolAuthorizerType, "authorizer-type", "", "", "Authorizer Type")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolBrowserId, "browser-id", "", "", "Browser ID")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolBrowserSigning, "browser-signing", "", "", "Browser Signing")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolClientToken, "client-token", "", "", "Client Token")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolCodeInterpreterId, "code-interpreter-id", "", "", "Code Interpreter ID")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolContent, "content", "", "", "Content")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolCredentialProviderConfigurations, "credential-provider-configurations", "", "", "Credential Provider Configurations")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolCredentialProviderVendor, "credential-provider-vendor", "", "", "Credential Provider Vendor")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolDataSourceConfig, "data-source-config", "", "", "Data Source Config")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolDefinition, "definition", "", "", "Definition")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolDescription, "description", "", "", "Description")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolEnableOnCreate, "enable-on-create", "", "", "Enable On Create")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolEncryptionKeyArn, "encryption-key-arn", "", "", "Encryption Key ARN")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolEndpointName, "endpoint-name", "", "", "Endpoint Name")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolEnvironmentVariables, "environment-variables", "", "", "Environment Variables")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolEvaluationExecutionRoleArn, "evaluation-execution-role-arn", "", "", "Evaluation Execution Role ARN")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolEvaluatorConfig, "evaluator-config", "", "", "Evaluator Config")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolEvaluatorId, "evaluator-id", "", "", "Evaluator ID")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolEvaluatorName, "evaluator-name", "", "", "Evaluator Name")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolEvaluators, "evaluators", "", "", "Evaluators")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolEventExpiryDuration, "event-expiry-duration", "", "", "Event Expiry Duration")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolExceptionLevel, "exception-level", "", "", "Exception Level")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolExecutionRoleArn, "execution-role-arn", "", "", "Execution Role ARN")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolExecutionStatus, "execution-status", "", "", "Execution Status")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolGatewayIdentifier, "gateway-identifier", "", "", "Gateway Identifier")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolInterceptorConfigurations, "interceptor-configurations", "", "", "Interceptor Configurations")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolKmsConfiguration, "kms-configuration", "", "", "KMS Configuration")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolKmsKeyArn, "kms-key-arn", "", "", "KMS Key ARN")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolLevel, "level", "", "", "Level")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolLifecycleConfiguration, "lifecycle-configuration", "", "", "Lifecycle Configuration")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolMaxResults, "max-results", "", "", "Max Results")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolMemoryExecutionRoleArn, "memory-execution-role-arn", "", "", "Memory Execution Role ARN")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolMemoryId, "memory-id", "", "", "Memory ID")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolMemoryStrategies, "memory-strategies", "", "", "Memory Strategies")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolMetadataConfiguration, "metadata-configuration", "", "", "Metadata Configuration")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolName, "name", "", "", "Name")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolNetworkConfiguration, "network-configuration", "", "", "Network Configuration")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolNextToken, "next-token", "", "", "Next Token")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolOauth2ProviderConfigInput, "oauth2-provider-config-input", "", "", "Oauth2 Provider Config Input")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolOnlineEvaluationConfigId, "online-evaluation-config-id", "", "", "Online Evaluation Config ID")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolOnlineEvaluationConfigName, "online-evaluation-config-name", "", "", "Online Evaluation Config Name")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolPolicy, "policy", "", "", "Policy")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolPolicyEngineConfiguration, "policy-engine-configuration", "", "", "Policy Engine Configuration")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolPolicyEngineId, "policy-engine-id", "", "", "Policy Engine ID")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolPolicyGenerationId, "policy-generation-id", "", "", "Policy Generation ID")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolPolicyId, "policy-id", "", "", "Policy ID")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolProfileId, "profile-id", "", "", "Profile ID")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolProtocolConfiguration, "protocol-configuration", "", "", "Protocol Configuration")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolProtocolType, "protocol-type", "", "", "Protocol Type")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolRecording, "recording", "", "", "Recording")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolRequestHeaderConfiguration, "request-header-configuration", "", "", "Request Header Configuration")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolResource, "resource", "", "", "Resource")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolResourceArn, "resource-arn", "", "", "Resource ARN")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolRoleArn, "role-arn", "", "", "Role ARN")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolRule, "rule", "", "", "Rule")
	_bedrockagentcorecontrolCmd.Flags().StringSliceVarP(&_bedrockagentcorecontrolTagKeys, "tag-keys", "", nil, "Tag Keys")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolTags, "tags", "", "", "Tags")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolTargetConfiguration, "target-configuration", "", "", "Target Configuration")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolTargetId, "target-id", "", "", "Target ID")
	_bedrockagentcorecontrolCmd.Flags().StringSliceVarP(&_bedrockagentcorecontrolTargetIdList, "target-id-list", "", nil, "Target ID List")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolTargetResourceScope, "target-resource-scope", "", "", "Target Resource Scope")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolTokenVaultId, "token-vault-id", "", "", "Token Vault ID")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolType, "type", "", "", "Type")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolValidationMode, "validation-mode", "", "", "Validation Mode")
	_bedrockagentcorecontrolCmd.Flags().StringVarP(&_bedrockagentcorecontrolView, "view", "", "", "View")

	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolCreateAgentRuntime, "create-agent-runtime", "", false, "Create Agent Runtime")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolCreateAgentRuntimeEndpoint, "create-agent-runtime-endpoint", "", false, "Create Agent Runtime Endpoint")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolCreateApiKeyCredentialProvider, "create-api-key-credential-provider", "", false, "Create API Key Credential Provider")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolCreateBrowser, "create-browser", "", false, "Create Browser")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolCreateBrowserProfile, "create-browser-profile", "", false, "Create Browser Profile")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolCreateCodeInterpreter, "create-code-interpreter", "", false, "Create Code Interpreter")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolCreateEvaluator, "create-evaluator", "", false, "Create Evaluator")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolCreateGateway, "create-gateway", "", false, "Create Gateway")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolCreateGatewayTarget, "create-gateway-target", "", false, "Create Gateway Target")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolCreateMemory, "create-memory", "", false, "Create Memory")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolCreateOauth2CredentialProvider, "create-oauth2-credential-provider", "", false, "Create Oauth2 Credential Provider")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolCreateOnlineEvaluationConfig, "create-online-evaluation-config", "", false, "Create Online Evaluation Config")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolCreatePolicy, "create-policy", "", false, "Create Policy")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolCreatePolicyEngine, "create-policy-engine", "", false, "Create Policy Engine")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolCreateWorkloadIdentity, "create-workload-identity", "", false, "Create Workload Identity")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolDeleteAgentRuntime, "delete-agent-runtime", "", false, "Delete Agent Runtime")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolDeleteAgentRuntimeEndpoint, "delete-agent-runtime-endpoint", "", false, "Delete Agent Runtime Endpoint")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolDeleteApiKeyCredentialProvider, "delete-api-key-credential-provider", "", false, "Delete API Key Credential Provider")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolDeleteBrowser, "delete-browser", "", false, "Delete Browser")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolDeleteBrowserProfile, "delete-browser-profile", "", false, "Delete Browser Profile")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolDeleteCodeInterpreter, "delete-code-interpreter", "", false, "Delete Code Interpreter")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolDeleteEvaluator, "delete-evaluator", "", false, "Delete Evaluator")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolDeleteGateway, "delete-gateway", "", false, "Delete Gateway")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolDeleteGatewayTarget, "delete-gateway-target", "", false, "Delete Gateway Target")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolDeleteMemory, "delete-memory", "", false, "Delete Memory")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolDeleteOauth2CredentialProvider, "delete-oauth2-credential-provider", "", false, "Delete Oauth2 Credential Provider")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolDeleteOnlineEvaluationConfig, "delete-online-evaluation-config", "", false, "Delete Online Evaluation Config")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolDeletePolicy, "delete-policy", "", false, "Delete Policy")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolDeletePolicyEngine, "delete-policy-engine", "", false, "Delete Policy Engine")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolDeleteResourcePolicy, "delete-resource-policy", "", false, "Delete Resource Policy")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolDeleteWorkloadIdentity, "delete-workload-identity", "", false, "Delete Workload Identity")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolGetAgentRuntime, "get-agent-runtime", "", false, "Get Agent Runtime")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolGetAgentRuntimeEndpoint, "get-agent-runtime-endpoint", "", false, "Get Agent Runtime Endpoint")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolGetApiKeyCredentialProvider, "get-api-key-credential-provider", "", false, "Get API Key Credential Provider")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolGetBrowser, "get-browser", "", false, "Get Browser")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolGetBrowserProfile, "get-browser-profile", "", false, "Get Browser Profile")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolGetCodeInterpreter, "get-code-interpreter", "", false, "Get Code Interpreter")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolGetEvaluator, "get-evaluator", "", false, "Get Evaluator")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolGetGateway, "get-gateway", "", false, "Get Gateway")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolGetGatewayTarget, "get-gateway-target", "", false, "Get Gateway Target")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolGetMemory, "get-memory", "", false, "Get Memory")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolGetOauth2CredentialProvider, "get-oauth2-credential-provider", "", false, "Get Oauth2 Credential Provider")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolGetOnlineEvaluationConfig, "get-online-evaluation-config", "", false, "Get Online Evaluation Config")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolGetPolicy, "get-policy", "", false, "Get Policy")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolGetPolicyEngine, "get-policy-engine", "", false, "Get Policy Engine")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolGetPolicyGeneration, "get-policy-generation", "", false, "Get Policy Generation")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolGetResourcePolicy, "get-resource-policy", "", false, "Get Resource Policy")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolGetTokenVault, "get-token-vault", "", false, "Get Token Vault")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolGetWorkloadIdentity, "get-workload-identity", "", false, "Get Workload Identity")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolListAgentRuntimeEndpoints, "list-agent-runtime-endpoints", "", false, "List Agent Runtime Endpoints")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolListAgentRuntimeVersions, "list-agent-runtime-versions", "", false, "List Agent Runtime Versions")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolListAgentRuntimes, "list-agent-runtimes", "", false, "List Agent Runtimes")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolListApiKeyCredentialProviders, "list-api-key-credential-providers", "", false, "List API Key Credential Providers")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolListBrowserProfiles, "list-browser-profiles", "", false, "List Browser Profiles")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolListBrowsers, "list-browsers", "", false, "List Browsers")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolListCodeInterpreters, "list-code-interpreters", "", false, "List Code Interpreters")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolListEvaluators, "list-evaluators", "", false, "List Evaluators")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolListGatewayTargets, "list-gateway-targets", "", false, "List Gateway Targets")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolListGateways, "list-gateways", "", false, "List Gateways")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolListMemories, "list-memories", "", false, "List Memories")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolListOauth2CredentialProviders, "list-oauth2-credential-providers", "", false, "List Oauth2 Credential Providers")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolListOnlineEvaluationConfigs, "list-online-evaluation-configs", "", false, "List Online Evaluation Configs")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolListPolicies, "list-policies", "", false, "List Policies")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolListPolicyEngines, "list-policy-engines", "", false, "List Policy Engines")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolListPolicyGenerationAssets, "list-policy-generation-assets", "", false, "List Policy Generation Assets")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolListPolicyGenerations, "list-policy-generations", "", false, "List Policy Generations")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolListWorkloadIdentities, "list-workload-identities", "", false, "List Workload Identities")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolPutResourcePolicy, "put-resource-policy", "", false, "Put Resource Policy")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolSetTokenVaultCMK, "set-token-vault-cmk", "", false, "Set Token Vault Cmk")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolStartPolicyGeneration, "start-policy-generation", "", false, "Start Policy Generation")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolSynchronizeGatewayTargets, "synchronize-gateway-targets", "", false, "Synchronize Gateway Targets")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolTagResource, "tag-resource", "", false, "Tag Resource")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolUntagResource, "untag-resource", "", false, "Untag Resource")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolUpdateAgentRuntime, "update-agent-runtime", "", false, "Update Agent Runtime")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolUpdateAgentRuntimeEndpoint, "update-agent-runtime-endpoint", "", false, "Update Agent Runtime Endpoint")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolUpdateApiKeyCredentialProvider, "update-api-key-credential-provider", "", false, "Update API Key Credential Provider")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolUpdateEvaluator, "update-evaluator", "", false, "Update Evaluator")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolUpdateGateway, "update-gateway", "", false, "Update Gateway")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolUpdateGatewayTarget, "update-gateway-target", "", false, "Update Gateway Target")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolUpdateMemory, "update-memory", "", false, "Update Memory")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolUpdateOauth2CredentialProvider, "update-oauth2-credential-provider", "", false, "Update Oauth2 Credential Provider")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolUpdateOnlineEvaluationConfig, "update-online-evaluation-config", "", false, "Update Online Evaluation Config")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolUpdatePolicy, "update-policy", "", false, "Update Policy")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolUpdatePolicyEngine, "update-policy-engine", "", false, "Update Policy Engine")
	_bedrockagentcorecontrolCmd.Flags().BoolVarP(&_bedrockagentcorecontrolUpdateWorkloadIdentity, "update-workload-identity", "", false, "Update Workload Identity")

}
