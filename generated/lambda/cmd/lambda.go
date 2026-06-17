package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// lambdaCmd represents the lambda command
var _lambdaCmd = &cobra.Command{
	Use:   "lambda",
	Short: "AWS lambda CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := lambda.NewFromConfig(cfg)
		if _lambdaAddLayerVersionPermission {
			lambda_AddLayerVersionPermission(cfg, client)
			return
		}
		if _lambdaAddPermission {
			lambda_AddPermission(cfg, client)
			return
		}
		if _lambdaCheckpointDurableExecution {
			lambda_CheckpointDurableExecution(cfg, client)
			return
		}
		if _lambdaCreateAlias {
			lambda_CreateAlias(cfg, client)
			return
		}
		if _lambdaCreateCapacityProvider {
			lambda_CreateCapacityProvider(cfg, client)
			return
		}
		if _lambdaCreateCodeSigningConfig {
			lambda_CreateCodeSigningConfig(cfg, client)
			return
		}
		if _lambdaCreateEventSourceMapping {
			lambda_CreateEventSourceMapping(cfg, client)
			return
		}
		if _lambdaCreateFunction {
			lambda_CreateFunction(cfg, client)
			return
		}
		if _lambdaCreateFunctionUrlConfig {
			lambda_CreateFunctionUrlConfig(cfg, client)
			return
		}
		if _lambdaDeleteAlias {
			lambda_DeleteAlias(cfg, client)
			return
		}
		if _lambdaDeleteCapacityProvider {
			lambda_DeleteCapacityProvider(cfg, client)
			return
		}
		if _lambdaDeleteCodeSigningConfig {
			lambda_DeleteCodeSigningConfig(cfg, client)
			return
		}
		if _lambdaDeleteEventSourceMapping {
			lambda_DeleteEventSourceMapping(cfg, client)
			return
		}
		if _lambdaDeleteFunction {
			lambda_DeleteFunction(cfg, client)
			return
		}
		if _lambdaDeleteFunctionCodeSigningConfig {
			lambda_DeleteFunctionCodeSigningConfig(cfg, client)
			return
		}
		if _lambdaDeleteFunctionConcurrency {
			lambda_DeleteFunctionConcurrency(cfg, client)
			return
		}
		if _lambdaDeleteFunctionEventInvokeConfig {
			lambda_DeleteFunctionEventInvokeConfig(cfg, client)
			return
		}
		if _lambdaDeleteFunctionUrlConfig {
			lambda_DeleteFunctionUrlConfig(cfg, client)
			return
		}
		if _lambdaDeleteLayerVersion {
			lambda_DeleteLayerVersion(cfg, client)
			return
		}
		if _lambdaDeleteProvisionedConcurrencyConfig {
			lambda_DeleteProvisionedConcurrencyConfig(cfg, client)
			return
		}
		if _lambdaGetAccountSettings {
			lambda_GetAccountSettings(cfg, client)
			return
		}
		if _lambdaGetAlias {
			lambda_GetAlias(cfg, client)
			return
		}
		if _lambdaGetCapacityProvider {
			lambda_GetCapacityProvider(cfg, client)
			return
		}
		if _lambdaGetCodeSigningConfig {
			lambda_GetCodeSigningConfig(cfg, client)
			return
		}
		if _lambdaGetDurableExecution {
			lambda_GetDurableExecution(cfg, client)
			return
		}
		if _lambdaGetDurableExecutionHistory {
			lambda_GetDurableExecutionHistory(cfg, client)
			return
		}
		if _lambdaGetDurableExecutionState {
			lambda_GetDurableExecutionState(cfg, client)
			return
		}
		if _lambdaGetEventSourceMapping {
			lambda_GetEventSourceMapping(cfg, client)
			return
		}
		if _lambdaGetFunction {
			lambda_GetFunction(cfg, client)
			return
		}
		if _lambdaGetFunctionCodeSigningConfig {
			lambda_GetFunctionCodeSigningConfig(cfg, client)
			return
		}
		if _lambdaGetFunctionConcurrency {
			lambda_GetFunctionConcurrency(cfg, client)
			return
		}
		if _lambdaGetFunctionConfiguration {
			lambda_GetFunctionConfiguration(cfg, client)
			return
		}
		if _lambdaGetFunctionEventInvokeConfig {
			lambda_GetFunctionEventInvokeConfig(cfg, client)
			return
		}
		if _lambdaGetFunctionRecursionConfig {
			lambda_GetFunctionRecursionConfig(cfg, client)
			return
		}
		if _lambdaGetFunctionScalingConfig {
			lambda_GetFunctionScalingConfig(cfg, client)
			return
		}
		if _lambdaGetFunctionUrlConfig {
			lambda_GetFunctionUrlConfig(cfg, client)
			return
		}
		if _lambdaGetLayerVersion {
			lambda_GetLayerVersion(cfg, client)
			return
		}
		if _lambdaGetLayerVersionByArn {
			lambda_GetLayerVersionByArn(cfg, client)
			return
		}
		if _lambdaGetLayerVersionPolicy {
			lambda_GetLayerVersionPolicy(cfg, client)
			return
		}
		if _lambdaGetPolicy {
			lambda_GetPolicy(cfg, client)
			return
		}
		if _lambdaGetProvisionedConcurrencyConfig {
			lambda_GetProvisionedConcurrencyConfig(cfg, client)
			return
		}
		if _lambdaGetRuntimeManagementConfig {
			lambda_GetRuntimeManagementConfig(cfg, client)
			return
		}
		if _lambdaInvoke {
			lambda_Invoke(cfg, client)
			return
		}
		if _lambdaInvokeAsync {
			lambda_InvokeAsync(cfg, client)
			return
		}
		if _lambdaInvokeWithResponseStream {
			lambda_InvokeWithResponseStream(cfg, client)
			return
		}
		if _lambdaListAliases {
			lambda_ListAliases(cfg, client)
			return
		}
		if _lambdaListCapacityProviders {
			lambda_ListCapacityProviders(cfg, client)
			return
		}
		if _lambdaListCodeSigningConfigs {
			lambda_ListCodeSigningConfigs(cfg, client)
			return
		}
		if _lambdaListDurableExecutionsByFunction {
			lambda_ListDurableExecutionsByFunction(cfg, client)
			return
		}
		if _lambdaListEventSourceMappings {
			lambda_ListEventSourceMappings(cfg, client)
			return
		}
		if _lambdaListFunctionEventInvokeConfigs {
			lambda_ListFunctionEventInvokeConfigs(cfg, client)
			return
		}
		if _lambdaListFunctionUrlConfigs {
			lambda_ListFunctionUrlConfigs(cfg, client)
			return
		}
		if _lambdaListFunctionVersionsByCapacityProvider {
			lambda_ListFunctionVersionsByCapacityProvider(cfg, client)
			return
		}
		if _lambdaListFunctions {
			lambda_ListFunctions(cfg, client)
			return
		}
		if _lambdaListFunctionsByCodeSigningConfig {
			lambda_ListFunctionsByCodeSigningConfig(cfg, client)
			return
		}
		if _lambdaListLayerVersions {
			lambda_ListLayerVersions(cfg, client)
			return
		}
		if _lambdaListLayers {
			lambda_ListLayers(cfg, client)
			return
		}
		if _lambdaListProvisionedConcurrencyConfigs {
			lambda_ListProvisionedConcurrencyConfigs(cfg, client)
			return
		}
		if _lambdaListTags {
			lambda_ListTags(cfg, client)
			return
		}
		if _lambdaListVersionsByFunction {
			lambda_ListVersionsByFunction(cfg, client)
			return
		}
		if _lambdaPublishLayerVersion {
			lambda_PublishLayerVersion(cfg, client)
			return
		}
		if _lambdaPublishVersion {
			lambda_PublishVersion(cfg, client)
			return
		}
		if _lambdaPutFunctionCodeSigningConfig {
			lambda_PutFunctionCodeSigningConfig(cfg, client)
			return
		}
		if _lambdaPutFunctionConcurrency {
			lambda_PutFunctionConcurrency(cfg, client)
			return
		}
		if _lambdaPutFunctionEventInvokeConfig {
			lambda_PutFunctionEventInvokeConfig(cfg, client)
			return
		}
		if _lambdaPutFunctionRecursionConfig {
			lambda_PutFunctionRecursionConfig(cfg, client)
			return
		}
		if _lambdaPutFunctionScalingConfig {
			lambda_PutFunctionScalingConfig(cfg, client)
			return
		}
		if _lambdaPutProvisionedConcurrencyConfig {
			lambda_PutProvisionedConcurrencyConfig(cfg, client)
			return
		}
		if _lambdaPutRuntimeManagementConfig {
			lambda_PutRuntimeManagementConfig(cfg, client)
			return
		}
		if _lambdaRemoveLayerVersionPermission {
			lambda_RemoveLayerVersionPermission(cfg, client)
			return
		}
		if _lambdaRemovePermission {
			lambda_RemovePermission(cfg, client)
			return
		}
		if _lambdaSendDurableExecutionCallbackFailure {
			lambda_SendDurableExecutionCallbackFailure(cfg, client)
			return
		}
		if _lambdaSendDurableExecutionCallbackHeartbeat {
			lambda_SendDurableExecutionCallbackHeartbeat(cfg, client)
			return
		}
		if _lambdaSendDurableExecutionCallbackSuccess {
			lambda_SendDurableExecutionCallbackSuccess(cfg, client)
			return
		}
		if _lambdaStopDurableExecution {
			lambda_StopDurableExecution(cfg, client)
			return
		}
		if _lambdaTagResource {
			lambda_TagResource(cfg, client)
			return
		}
		if _lambdaUntagResource {
			lambda_UntagResource(cfg, client)
			return
		}
		if _lambdaUpdateAlias {
			lambda_UpdateAlias(cfg, client)
			return
		}
		if _lambdaUpdateCapacityProvider {
			lambda_UpdateCapacityProvider(cfg, client)
			return
		}
		if _lambdaUpdateCodeSigningConfig {
			lambda_UpdateCodeSigningConfig(cfg, client)
			return
		}
		if _lambdaUpdateEventSourceMapping {
			lambda_UpdateEventSourceMapping(cfg, client)
			return
		}
		if _lambdaUpdateFunctionCode {
			lambda_UpdateFunctionCode(cfg, client)
			return
		}
		if _lambdaUpdateFunctionConfiguration {
			lambda_UpdateFunctionConfiguration(cfg, client)
			return
		}
		if _lambdaUpdateFunctionEventInvokeConfig {
			lambda_UpdateFunctionEventInvokeConfig(cfg, client)
			return
		}
		if _lambdaUpdateFunctionUrlConfig {
			lambda_UpdateFunctionUrlConfig(cfg, client)
			return
		}

	},
}

var (
	_lambdaAddLayerVersionPermission              bool
	_lambdaAddPermission                          bool
	_lambdaCheckpointDurableExecution             bool
	_lambdaCreateAlias                            bool
	_lambdaCreateCapacityProvider                 bool
	_lambdaCreateCodeSigningConfig                bool
	_lambdaCreateEventSourceMapping               bool
	_lambdaCreateFunction                         bool
	_lambdaCreateFunctionUrlConfig                bool
	_lambdaDeleteAlias                            bool
	_lambdaDeleteCapacityProvider                 bool
	_lambdaDeleteCodeSigningConfig                bool
	_lambdaDeleteEventSourceMapping               bool
	_lambdaDeleteFunction                         bool
	_lambdaDeleteFunctionCodeSigningConfig        bool
	_lambdaDeleteFunctionConcurrency              bool
	_lambdaDeleteFunctionEventInvokeConfig        bool
	_lambdaDeleteFunctionUrlConfig                bool
	_lambdaDeleteLayerVersion                     bool
	_lambdaDeleteProvisionedConcurrencyConfig     bool
	_lambdaGetAccountSettings                     bool
	_lambdaGetAlias                               bool
	_lambdaGetCapacityProvider                    bool
	_lambdaGetCodeSigningConfig                   bool
	_lambdaGetDurableExecution                    bool
	_lambdaGetDurableExecutionHistory             bool
	_lambdaGetDurableExecutionState               bool
	_lambdaGetEventSourceMapping                  bool
	_lambdaGetFunction                            bool
	_lambdaGetFunctionCodeSigningConfig           bool
	_lambdaGetFunctionConcurrency                 bool
	_lambdaGetFunctionConfiguration               bool
	_lambdaGetFunctionEventInvokeConfig           bool
	_lambdaGetFunctionRecursionConfig             bool
	_lambdaGetFunctionScalingConfig               bool
	_lambdaGetFunctionUrlConfig                   bool
	_lambdaGetLayerVersion                        bool
	_lambdaGetLayerVersionByArn                   bool
	_lambdaGetLayerVersionPolicy                  bool
	_lambdaGetPolicy                              bool
	_lambdaGetProvisionedConcurrencyConfig        bool
	_lambdaGetRuntimeManagementConfig             bool
	_lambdaInvoke                                 bool
	_lambdaInvokeAsync                            bool
	_lambdaInvokeWithResponseStream               bool
	_lambdaListAliases                            bool
	_lambdaListCapacityProviders                  bool
	_lambdaListCodeSigningConfigs                 bool
	_lambdaListDurableExecutionsByFunction        bool
	_lambdaListEventSourceMappings                bool
	_lambdaListFunctionEventInvokeConfigs         bool
	_lambdaListFunctionUrlConfigs                 bool
	_lambdaListFunctionVersionsByCapacityProvider bool
	_lambdaListFunctions                          bool
	_lambdaListFunctionsByCodeSigningConfig       bool
	_lambdaListLayerVersions                      bool
	_lambdaListLayers                             bool
	_lambdaListProvisionedConcurrencyConfigs      bool
	_lambdaListTags                               bool
	_lambdaListVersionsByFunction                 bool
	_lambdaPublishLayerVersion                    bool
	_lambdaPublishVersion                         bool
	_lambdaPutFunctionCodeSigningConfig           bool
	_lambdaPutFunctionConcurrency                 bool
	_lambdaPutFunctionEventInvokeConfig           bool
	_lambdaPutFunctionRecursionConfig             bool
	_lambdaPutFunctionScalingConfig               bool
	_lambdaPutProvisionedConcurrencyConfig        bool
	_lambdaPutRuntimeManagementConfig             bool
	_lambdaRemoveLayerVersionPermission           bool
	_lambdaRemovePermission                       bool
	_lambdaSendDurableExecutionCallbackFailure    bool
	_lambdaSendDurableExecutionCallbackHeartbeat  bool
	_lambdaSendDurableExecutionCallbackSuccess    bool
	_lambdaStopDurableExecution                   bool
	_lambdaTagResource                            bool
	_lambdaUntagResource                          bool
	_lambdaUpdateAlias                            bool
	_lambdaUpdateCapacityProvider                 bool
	_lambdaUpdateCodeSigningConfig                bool
	_lambdaUpdateEventSourceMapping               bool
	_lambdaUpdateFunctionCode                     bool
	_lambdaUpdateFunctionConfiguration            bool
	_lambdaUpdateFunctionEventInvokeConfig        bool
	_lambdaUpdateFunctionUrlConfig                bool

	_lambdaAction                              string
	_lambdaAllowedPublishers                   string
	_lambdaAmazonManagedKafkaEventSourceConfig string
	_lambdaArchitectures                       string
	_lambdaArn                                 string
	_lambdaAuthType                            string
	_lambdaBatchSize                           string
	_lambdaBisectBatchOnFunctionError          string
	_lambdaCallbackId                          string
	_lambdaCapacityProviderConfig              string
	_lambdaCapacityProviderName                string
	_lambdaCapacityProviderScalingConfig       string
	_lambdaCheckpointToken                     string
	_lambdaClientContext                       string
	_lambdaClientToken                         string
	_lambdaCode                                string
	_lambdaCodeSha256                          string
	_lambdaCodeSigningConfigArn                string
	_lambdaCodeSigningPolicies                 string
	_lambdaCompatibleArchitecture              string
	_lambdaCompatibleArchitectures             string
	_lambdaCompatibleRuntime                   string
	_lambdaCompatibleRuntimes                  string
	_lambdaContent                             string
	_lambdaCors                                string
	_lambdaDeadLetterConfig                    string
	_lambdaDescription                         string
	_lambdaDestinationConfig                   string
	_lambdaDocumentDBEventSourceConfig         string
	_lambdaDryRun                              string
	_lambdaDurableConfig                       string
	_lambdaDurableExecutionArn                 string
	_lambdaDurableExecutionName                string
	_lambdaEnabled                             string
	_lambdaEnvironment                         string
	_lambdaEphemeralStorage                    string
	_lambdaError                               string
	_lambdaEventSourceArn                      string
	_lambdaEventSourceToken                    string
	_lambdaFileSystemConfigs                   string
	_lambdaFilterCriteria                      string
	_lambdaFunctionName                        string
	_lambdaFunctionResponseTypes               string
	_lambdaFunctionScalingConfig               string
	_lambdaFunctionUrlAuthType                 string
	_lambdaFunctionVersion                     string
	_lambdaHandler                             string
	_lambdaImageConfig                         string
	_lambdaImageUri                            string
	_lambdaIncludeExecutionData                string
	_lambdaInstanceRequirements                string
	_lambdaInvocationType                      string
	_lambdaInvokeArgs                          string
	_lambdaInvokeMode                          string
	_lambdaInvokedViaFunctionUrl               string
	_lambdaKMSKeyArn                           string
	_lambdaLayerName                           string
	_lambdaLayers                              []string
	_lambdaLicenseInfo                         string
	_lambdaLogType                             string
	_lambdaLoggingConfig                       string
	_lambdaMarker                              string
	_lambdaMasterRegion                        string
	_lambdaMaxItems                            string
	_lambdaMaximumBatchingWindowInSeconds      string
	_lambdaMaximumEventAgeInSeconds            string
	_lambdaMaximumRecordAgeInSeconds           string
	_lambdaMaximumRetryAttempts                string
	_lambdaMemorySize                          string
	_lambdaMetricsConfig                       string
	_lambdaName                                string
	_lambdaOrganizationId                      string
	_lambdaPackageType                         string
	_lambdaParallelizationFactor               string
	_lambdaPayload                             string
	_lambdaPermissionsConfig                   string
	_lambdaPrincipal                           string
	_lambdaPrincipalOrgID                      string
	_lambdaProvisionedConcurrentExecutions     string
	_lambdaProvisionedPollerConfig             string
	_lambdaPublish                             string
	_lambdaPublishTo                           string
	_lambdaQualifier                           string
	_lambdaQueues                              []string
	_lambdaRecursiveLoop                       string
	_lambdaReservedConcurrentExecutions        string
	_lambdaResource                            string
	_lambdaResult                              string
	_lambdaReverseOrder                        string
	_lambdaRevisionId                          string
	_lambdaRole                                string
	_lambdaRoutingConfig                       string
	_lambdaRuntime                             string
	_lambdaRuntimeVersionArn                   string
	_lambdaS3Bucket                            string
	_lambdaS3Key                               string
	_lambdaS3ObjectVersion                     string
	_lambdaScalingConfig                       string
	_lambdaSelfManagedEventSource              string
	_lambdaSelfManagedKafkaEventSourceConfig   string
	_lambdaSnapStart                           string
	_lambdaSourceAccessConfigurations          string
	_lambdaSourceAccount                       string
	_lambdaSourceArn                           string
	_lambdaSourceKMSKeyArn                     string
	_lambdaStartedAfter                        string
	_lambdaStartedBefore                       string
	_lambdaStartingPosition                    string
	_lambdaStartingPositionTimestamp           string
	_lambdaState                               string
	_lambdaStatementId                         string
	_lambdaStatuses                            string
	_lambdaTagKeys                             []string
	_lambdaTags                                string
	_lambdaTenancyConfig                       string
	_lambdaTenantId                            string
	_lambdaTimeout                             string
	_lambdaTopics                              []string
	_lambdaTracingConfig                       string
	_lambdaTumblingWindowInSeconds             string
	_lambdaUpdateRuntimeOn                     string
	_lambdaUpdates                             string
	_lambdaUUID                                string
	_lambdaVersionNumber                       string
	_lambdaVpcConfig                           string
	_lambdaZipFile                             string
)

// Adds permissions to the resource-based policy of a version of an [Lambda layer]. Use this
// action to grant layer usage permission to other accounts. You can grant
// permission to a single account, all accounts in an organization, or all Amazon
// Web Services accounts.
//
// To revoke permission, call RemoveLayerVersionPermission with the statement ID that you specified when you
// added it.
//
// [Lambda layer]: https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html
func lambda_AddLayerVersionPermission(cfg aws.Config, client *lambda.Client) {
	input := &lambda.AddLayerVersionPermissionInput{
		// Action: *string, // Required
		// LayerName: *string, // Required
		// Principal: *string, // Required
		// StatementId: *string, // Required
		// VersionNumber: *int64, // Required
	}

	if len(_lambdaAction) > 0 {
		input.Action = aws.String(_lambdaAction)
	}
	if len(_lambdaLayerName) > 0 {
		input.LayerName = aws.String(_lambdaLayerName)
	}
	if len(_lambdaPrincipal) > 0 {
		input.Principal = aws.String(_lambdaPrincipal)
	}
	if len(_lambdaStatementId) > 0 {
		input.StatementId = aws.String(_lambdaStatementId)
	}
	if len(_lambdaVersionNumber) > 0 {
		if err := assignInputField(input, "VersionNumber", _lambdaVersionNumber); err != nil {
			log.Errorf("invalid --version-number: %s", err.Error())
			return
		}
	}
	if len(_lambdaOrganizationId) > 0 {
		input.OrganizationId = aws.String(_lambdaOrganizationId)
	}
	if len(_lambdaRevisionId) > 0 {
		input.RevisionId = aws.String(_lambdaRevisionId)
	}

	if resp, err := client.AddLayerVersionPermission(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Grants a [principal] permission to use a function. You can apply the policy at the
// function level, or specify a qualifier to restrict access to a single version or
// alias. If you use a qualifier, the invoker must use the full Amazon Resource
// Name (ARN) of that version or alias to invoke the function. Note: Lambda does
// not support adding policies to version $LATEST.
//
// To grant permission to another account, specify the account ID as the Principal
// . To grant permission to an organization defined in Organizations, specify the
// organization ID as the PrincipalOrgID . For Amazon Web Services services, the
// principal is a domain-style identifier that the service defines, such as
// s3.amazonaws.com or sns.amazonaws.com . For Amazon Web Services services, you
// can also specify the ARN of the associated resource as the SourceArn . If you
// grant permission to a service principal without specifying the source, other
// accounts could potentially configure resources in their account to invoke your
// Lambda function.
//
// This operation adds a statement to a resource-based permissions policy for the
// function. For more information about function policies, see [Using resource-based policies for Lambda].
//
// [Using resource-based policies for Lambda]: https://docs.aws.amazon.com/lambda/latest/dg/access-control-resource-based.html
// [principal]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_principal.html#Principal_specifying
func lambda_AddPermission(cfg aws.Config, client *lambda.Client) {
	input := &lambda.AddPermissionInput{
		// Action: *string, // Required
		// FunctionName: *string, // Required
		// Principal: *string, // Required
		// StatementId: *string, // Required
	}

	if len(_lambdaAction) > 0 {
		input.Action = aws.String(_lambdaAction)
	}
	if len(_lambdaFunctionName) > 0 {
		input.FunctionName = aws.String(_lambdaFunctionName)
	}
	if len(_lambdaPrincipal) > 0 {
		input.Principal = aws.String(_lambdaPrincipal)
	}
	if len(_lambdaStatementId) > 0 {
		input.StatementId = aws.String(_lambdaStatementId)
	}
	if len(_lambdaEventSourceToken) > 0 {
		input.EventSourceToken = aws.String(_lambdaEventSourceToken)
	}
	if len(_lambdaFunctionUrlAuthType) > 0 {
		if err := assignInputField(input, "FunctionUrlAuthType", _lambdaFunctionUrlAuthType); err != nil {
			log.Errorf("invalid --function-url-auth-type: %s", err.Error())
			return
		}
	}
	if len(_lambdaInvokedViaFunctionUrl) > 0 {
		if err := assignInputField(input, "InvokedViaFunctionUrl", _lambdaInvokedViaFunctionUrl); err != nil {
			log.Errorf("invalid --invoked-via-function-url: %s", err.Error())
			return
		}
	}
	if len(_lambdaPrincipalOrgID) > 0 {
		input.PrincipalOrgID = aws.String(_lambdaPrincipalOrgID)
	}
	if len(_lambdaQualifier) > 0 {
		input.Qualifier = aws.String(_lambdaQualifier)
	}
	if len(_lambdaRevisionId) > 0 {
		input.RevisionId = aws.String(_lambdaRevisionId)
	}
	if len(_lambdaSourceAccount) > 0 {
		input.SourceAccount = aws.String(_lambdaSourceAccount)
	}
	if len(_lambdaSourceArn) > 0 {
		input.SourceArn = aws.String(_lambdaSourceArn)
	}

	if resp, err := client.AddPermission(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Saves the progress of a [durable function] execution during runtime. This API is used by the
// Lambda durable functions SDK to checkpoint completed steps and schedule
// asynchronous operations. You typically don't need to call this API directly as
// the SDK handles checkpointing automatically.
//
// Each checkpoint operation consumes the current checkpoint token and returns a
// new one for the next checkpoint. This ensures that checkpoints are applied in
// the correct order and prevents duplicate or out-of-order state updates.
//
// [durable function]: https://docs.aws.amazon.com/lambda/latest/dg/durable-functions.html
func lambda_CheckpointDurableExecution(cfg aws.Config, client *lambda.Client) {
	input := &lambda.CheckpointDurableExecutionInput{
		// CheckpointToken: *string, // Required
		// DurableExecutionArn: *string, // Required
	}

	if len(_lambdaCheckpointToken) > 0 {
		input.CheckpointToken = aws.String(_lambdaCheckpointToken)
	}
	if len(_lambdaDurableExecutionArn) > 0 {
		input.DurableExecutionArn = aws.String(_lambdaDurableExecutionArn)
	}
	if len(_lambdaClientToken) > 0 {
		input.ClientToken = aws.String(_lambdaClientToken)
	}
	if len(_lambdaUpdates) > 0 {
		if err := assignInputField(input, "Updates", _lambdaUpdates); err != nil {
			log.Errorf("invalid --updates: %s", err.Error())
			return
		}
	}

	if resp, err := client.CheckpointDurableExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an [alias] for a Lambda function version. Use aliases to provide clients with
// a function identifier that you can update to invoke a different version.
//
// You can also map an alias to split invocation requests between two versions.
// Use the RoutingConfig parameter to specify a second version and the percentage
// of invocation requests that it receives.
//
// [alias]: https://docs.aws.amazon.com/lambda/latest/dg/configuration-aliases.html
func lambda_CreateAlias(cfg aws.Config, client *lambda.Client) {
	input := &lambda.CreateAliasInput{
		// FunctionName: *string, // Required
		// FunctionVersion: *string, // Required
		// Name: *string, // Required
	}

	if len(_lambdaFunctionName) > 0 {
		input.FunctionName = aws.String(_lambdaFunctionName)
	}
	if len(_lambdaFunctionVersion) > 0 {
		input.FunctionVersion = aws.String(_lambdaFunctionVersion)
	}
	if len(_lambdaName) > 0 {
		input.Name = aws.String(_lambdaName)
	}
	if len(_lambdaDescription) > 0 {
		input.Description = aws.String(_lambdaDescription)
	}
	if len(_lambdaRoutingConfig) > 0 {
		if err := assignInputField(input, "RoutingConfig", _lambdaRoutingConfig); err != nil {
			log.Errorf("invalid --routing-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAlias(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a capacity provider that manages compute resources for Lambda functions
func lambda_CreateCapacityProvider(cfg aws.Config, client *lambda.Client) {
	input := &lambda.CreateCapacityProviderInput{
		// CapacityProviderName: *string, // Required
		// PermissionsConfig: *types.CapacityProviderPermissionsConfig, // Required
		// VpcConfig: *types.CapacityProviderVpcConfig, // Required
	}

	if len(_lambdaCapacityProviderName) > 0 {
		input.CapacityProviderName = aws.String(_lambdaCapacityProviderName)
	}
	if len(_lambdaPermissionsConfig) > 0 {
		if err := assignInputField(input, "PermissionsConfig", _lambdaPermissionsConfig); err != nil {
			log.Errorf("invalid --permissions-config: %s", err.Error())
			return
		}
	}
	if len(_lambdaVpcConfig) > 0 {
		if err := assignInputField(input, "VpcConfig", _lambdaVpcConfig); err != nil {
			log.Errorf("invalid --vpc-config: %s", err.Error())
			return
		}
	}
	if len(_lambdaCapacityProviderScalingConfig) > 0 {
		if err := assignInputField(input, "CapacityProviderScalingConfig", _lambdaCapacityProviderScalingConfig); err != nil {
			log.Errorf("invalid --capacity-provider-scaling-config: %s", err.Error())
			return
		}
	}
	if len(_lambdaInstanceRequirements) > 0 {
		if err := assignInputField(input, "InstanceRequirements", _lambdaInstanceRequirements); err != nil {
			log.Errorf("invalid --instance-requirements: %s", err.Error())
			return
		}
	}
	if len(_lambdaKMSKeyArn) > 0 {
		input.KmsKeyArn = aws.String(_lambdaKMSKeyArn)
	}
	if len(_lambdaTags) > 0 {
		if err := assignInputField(input, "Tags", _lambdaTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCapacityProvider(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a code signing configuration. A [code signing configuration] defines a list of allowed signing
// profiles and defines the code-signing validation policy (action to be taken if
// deployment validation checks fail).
//
// [code signing configuration]: https://docs.aws.amazon.com/lambda/latest/dg/configuration-codesigning.html
func lambda_CreateCodeSigningConfig(cfg aws.Config, client *lambda.Client) {
	input := &lambda.CreateCodeSigningConfigInput{
		// AllowedPublishers: *types.AllowedPublishers, // Required
	}

	if len(_lambdaAllowedPublishers) > 0 {
		if err := assignInputField(input, "AllowedPublishers", _lambdaAllowedPublishers); err != nil {
			log.Errorf("invalid --allowed-publishers: %s", err.Error())
			return
		}
	}
	if len(_lambdaCodeSigningPolicies) > 0 {
		if err := assignInputField(input, "CodeSigningPolicies", _lambdaCodeSigningPolicies); err != nil {
			log.Errorf("invalid --code-signing-policies: %s", err.Error())
			return
		}
	}
	if len(_lambdaDescription) > 0 {
		input.Description = aws.String(_lambdaDescription)
	}
	if len(_lambdaTags) > 0 {
		if err := assignInputField(input, "Tags", _lambdaTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCodeSigningConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a mapping between an event source and an Lambda function. Lambda reads
// items from the event source and invokes the function.
//
// For details about how to configure different event sources, see the following
// topics.
//
// [Amazon DynamoDB Streams]
//
// [Amazon Kinesis]
//
// [Amazon SQS]
//
// [Amazon MQ and RabbitMQ]
//
// [Amazon MSK]
//
// [Apache Kafka]
//
// [Amazon DocumentDB]
//
// The following error handling options are available for stream sources
// (DynamoDB, Kinesis, Amazon MSK, and self-managed Apache Kafka):
//
// - BisectBatchOnFunctionError – If the function returns an error, split the
// batch in two and retry.
//
// - MaximumRecordAgeInSeconds – Discard records older than the specified age.
// The default value is infinite (-1). When set to infinite (-1), failed records
// are retried until the record expires
//
// - MaximumRetryAttempts – Discard records after the specified number of
// retries. The default value is infinite (-1). When set to infinite (-1), failed
// records are retried until the record expires.
//
// - OnFailure – Send discarded records to an Amazon SQS queue, Amazon SNS topic,
// Kafka topic, or Amazon S3 bucket. For more information, see [Adding a destination].
//
// The following option is available only for DynamoDB and Kinesis event sources:
//
// - ParallelizationFactor – Process multiple batches from each shard
// concurrently.
//
// For information about which configuration parameters apply to each event
// source, see the following topics.
//
// [Amazon DynamoDB Streams]
//
// [Amazon Kinesis]
//
// [Amazon SQS]
//
// [Amazon MQ and RabbitMQ]
//
// [Amazon MSK]
//
// [Apache Kafka]
//
// [Amazon DocumentDB]
//
// [Amazon DynamoDB Streams]: https://docs.aws.amazon.com/lambda/latest/dg/with-ddb.html#services-ddb-params
// [Amazon SQS]: https://docs.aws.amazon.com/lambda/latest/dg/with-sqs.html#services-sqs-params
// [Amazon MSK]: https://docs.aws.amazon.com/lambda/latest/dg/with-msk.html#services-msk-parms
// [Amazon Kinesis]: https://docs.aws.amazon.com/lambda/latest/dg/with-kinesis.html#services-kinesis-params
// [Amazon MQ and RabbitMQ]: https://docs.aws.amazon.com/lambda/latest/dg/with-mq.html#services-mq-params
// [Apache Kafka]: https://docs.aws.amazon.com/lambda/latest/dg/with-kafka.html#services-kafka-parms
// [Amazon DocumentDB]: https://docs.aws.amazon.com/lambda/latest/dg/with-documentdb.html#docdb-configuration
// [Adding a destination]: https://docs.aws.amazon.com/lambda/latest/dg/invocation-async-retain-records.html#invocation-async-destinations
func lambda_CreateEventSourceMapping(cfg aws.Config, client *lambda.Client) {
	input := &lambda.CreateEventSourceMappingInput{
		// FunctionName: *string, // Required
	}

	if len(_lambdaFunctionName) > 0 {
		input.FunctionName = aws.String(_lambdaFunctionName)
	}
	if len(_lambdaAmazonManagedKafkaEventSourceConfig) > 0 {
		if err := assignInputField(input, "AmazonManagedKafkaEventSourceConfig", _lambdaAmazonManagedKafkaEventSourceConfig); err != nil {
			log.Errorf("invalid --amazon-managed-kafka-event-source-config: %s", err.Error())
			return
		}
	}
	if len(_lambdaBatchSize) > 0 {
		if err := assignInputField(input, "BatchSize", _lambdaBatchSize); err != nil {
			log.Errorf("invalid --batch-size: %s", err.Error())
			return
		}
	}
	if len(_lambdaBisectBatchOnFunctionError) > 0 {
		if err := assignInputField(input, "BisectBatchOnFunctionError", _lambdaBisectBatchOnFunctionError); err != nil {
			log.Errorf("invalid --bisect-batch-on-function-error: %s", err.Error())
			return
		}
	}
	if len(_lambdaDestinationConfig) > 0 {
		if err := assignInputField(input, "DestinationConfig", _lambdaDestinationConfig); err != nil {
			log.Errorf("invalid --destination-config: %s", err.Error())
			return
		}
	}
	if len(_lambdaDocumentDBEventSourceConfig) > 0 {
		if err := assignInputField(input, "DocumentDBEventSourceConfig", _lambdaDocumentDBEventSourceConfig); err != nil {
			log.Errorf("invalid --document-db-event-source-config: %s", err.Error())
			return
		}
	}
	if len(_lambdaEnabled) > 0 {
		if err := assignInputField(input, "Enabled", _lambdaEnabled); err != nil {
			log.Errorf("invalid --enabled: %s", err.Error())
			return
		}
	}
	if len(_lambdaEventSourceArn) > 0 {
		input.EventSourceArn = aws.String(_lambdaEventSourceArn)
	}
	if len(_lambdaFilterCriteria) > 0 {
		if err := assignInputField(input, "FilterCriteria", _lambdaFilterCriteria); err != nil {
			log.Errorf("invalid --filter-criteria: %s", err.Error())
			return
		}
	}
	if len(_lambdaFunctionResponseTypes) > 0 {
		if err := assignInputField(input, "FunctionResponseTypes", _lambdaFunctionResponseTypes); err != nil {
			log.Errorf("invalid --function-response-types: %s", err.Error())
			return
		}
	}
	if len(_lambdaKMSKeyArn) > 0 {
		input.KMSKeyArn = aws.String(_lambdaKMSKeyArn)
	}
	if len(_lambdaLoggingConfig) > 0 {
		if err := assignInputField(input, "LoggingConfig", _lambdaLoggingConfig); err != nil {
			log.Errorf("invalid --logging-config: %s", err.Error())
			return
		}
	}
	if len(_lambdaMaximumBatchingWindowInSeconds) > 0 {
		if err := assignInputField(input, "MaximumBatchingWindowInSeconds", _lambdaMaximumBatchingWindowInSeconds); err != nil {
			log.Errorf("invalid --maximum-batching-window-in-seconds: %s", err.Error())
			return
		}
	}
	if len(_lambdaMaximumRecordAgeInSeconds) > 0 {
		if err := assignInputField(input, "MaximumRecordAgeInSeconds", _lambdaMaximumRecordAgeInSeconds); err != nil {
			log.Errorf("invalid --maximum-record-age-in-seconds: %s", err.Error())
			return
		}
	}
	if len(_lambdaMaximumRetryAttempts) > 0 {
		if err := assignInputField(input, "MaximumRetryAttempts", _lambdaMaximumRetryAttempts); err != nil {
			log.Errorf("invalid --maximum-retry-attempts: %s", err.Error())
			return
		}
	}
	if len(_lambdaMetricsConfig) > 0 {
		if err := assignInputField(input, "MetricsConfig", _lambdaMetricsConfig); err != nil {
			log.Errorf("invalid --metrics-config: %s", err.Error())
			return
		}
	}
	if len(_lambdaParallelizationFactor) > 0 {
		if err := assignInputField(input, "ParallelizationFactor", _lambdaParallelizationFactor); err != nil {
			log.Errorf("invalid --parallelization-factor: %s", err.Error())
			return
		}
	}
	if len(_lambdaProvisionedPollerConfig) > 0 {
		if err := assignInputField(input, "ProvisionedPollerConfig", _lambdaProvisionedPollerConfig); err != nil {
			log.Errorf("invalid --provisioned-poller-config: %s", err.Error())
			return
		}
	}
	if len(_lambdaQueues) > 0 {
		input.Queues = append([]string(nil), _lambdaQueues...)
	}
	if len(_lambdaScalingConfig) > 0 {
		if err := assignInputField(input, "ScalingConfig", _lambdaScalingConfig); err != nil {
			log.Errorf("invalid --scaling-config: %s", err.Error())
			return
		}
	}
	if len(_lambdaSelfManagedEventSource) > 0 {
		if err := assignInputField(input, "SelfManagedEventSource", _lambdaSelfManagedEventSource); err != nil {
			log.Errorf("invalid --self-managed-event-source: %s", err.Error())
			return
		}
	}
	if len(_lambdaSelfManagedKafkaEventSourceConfig) > 0 {
		if err := assignInputField(input, "SelfManagedKafkaEventSourceConfig", _lambdaSelfManagedKafkaEventSourceConfig); err != nil {
			log.Errorf("invalid --self-managed-kafka-event-source-config: %s", err.Error())
			return
		}
	}
	if len(_lambdaSourceAccessConfigurations) > 0 {
		if err := assignInputField(input, "SourceAccessConfigurations", _lambdaSourceAccessConfigurations); err != nil {
			log.Errorf("invalid --source-access-configurations: %s", err.Error())
			return
		}
	}
	if len(_lambdaStartingPosition) > 0 {
		if err := assignInputField(input, "StartingPosition", _lambdaStartingPosition); err != nil {
			log.Errorf("invalid --starting-position: %s", err.Error())
			return
		}
	}
	if len(_lambdaStartingPositionTimestamp) > 0 {
		if err := assignInputField(input, "StartingPositionTimestamp", _lambdaStartingPositionTimestamp); err != nil {
			log.Errorf("invalid --starting-position-timestamp: %s", err.Error())
			return
		}
	}
	if len(_lambdaTags) > 0 {
		if err := assignInputField(input, "Tags", _lambdaTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_lambdaTopics) > 0 {
		input.Topics = append([]string(nil), _lambdaTopics...)
	}
	if len(_lambdaTumblingWindowInSeconds) > 0 {
		if err := assignInputField(input, "TumblingWindowInSeconds", _lambdaTumblingWindowInSeconds); err != nil {
			log.Errorf("invalid --tumbling-window-in-seconds: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateEventSourceMapping(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a Lambda function. To create a function, you need a [deployment package] and an [execution role]. The
// deployment package is a .zip file archive or container image that contains your
// function code. The execution role grants the function permission to use Amazon
// Web Services services, such as Amazon CloudWatch Logs for log streaming and
// X-Ray for request tracing.
//
// If the deployment package is a [container image], then you set the package type to Image . For a
// container image, the code property must include the URI of a container image in
// the Amazon ECR registry. You do not need to specify the handler and runtime
// properties.
//
// If the deployment package is a [.zip file archive], then you set the package type to Zip . For a
// .zip file archive, the code property specifies the location of the .zip file.
// You must also specify the handler and runtime properties. The code in the
// deployment package must be compatible with the target instruction set
// architecture of the function ( x86-64 or arm64 ). If you do not specify the
// architecture, then the default value is x86-64 .
//
// When you create a function, Lambda provisions an instance of the function and
// its supporting resources. If your function connects to a VPC, this process can
// take a minute or so. During this time, you can't invoke or modify the function.
// The State , StateReason , and StateReasonCode fields in the response from GetFunctionConfiguration
// indicate when the function is ready to invoke. For more information, see [Lambda function states].
//
// A function has an unpublished version, and can have published versions and
// aliases. The unpublished version changes when you update your function's code
// and configuration. A published version is a snapshot of your function code and
// configuration that can't be changed. An alias is a named resource that maps to a
// version, and can be changed to map to a different version. Use the Publish
// parameter to create version 1 of your function from its initial configuration.
//
// The other parameters let you configure version-specific and function-level
// settings. You can modify version-specific settings later with UpdateFunctionConfiguration. Function-level
// settings apply to both the unpublished and published versions of the function,
// and include tags (TagResource ) and per-function concurrency limits (PutFunctionConcurrency ).
//
// You can use code signing if your deployment package is a .zip file archive. To
// enable code signing for this function, specify the ARN of a code-signing
// configuration. When a user attempts to deploy a code package with UpdateFunctionCode, Lambda
// checks that the code package has a valid signature from a trusted publisher. The
// code-signing configuration includes set of signing profiles, which define the
// trusted publishers for this function.
//
// If another Amazon Web Services account or an Amazon Web Services service
// invokes your function, use AddPermissionto grant permission by creating a resource-based
// Identity and Access Management (IAM) policy. You can grant permissions at the
// function level, on a version, or on an alias.
//
// To invoke your function directly, use Invoke. To invoke your function in response to
// events in other Amazon Web Services services, create an event source mapping (CreateEventSourceMapping
// ), or configure a function trigger in the other service. For more information,
// see [Invoking Lambda functions].
//
// [Invoking Lambda functions]: https://docs.aws.amazon.com/lambda/latest/dg/lambda-invocation.html
// [Lambda function states]: https://docs.aws.amazon.com/lambda/latest/dg/functions-states.html
// [.zip file archive]: https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-package.html#gettingstarted-package-zip
// [container image]: https://docs.aws.amazon.com/lambda/latest/dg/lambda-images.html
// [execution role]: https://docs.aws.amazon.com/lambda/latest/dg/intro-permission-model.html#lambda-intro-execution-role
// [deployment package]: https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-package.html
func lambda_CreateFunction(cfg aws.Config, client *lambda.Client) {
	input := &lambda.CreateFunctionInput{
		// Code: *types.FunctionCode, // Required
		// FunctionName: *string, // Required
		// Role: *string, // Required
	}

	if len(_lambdaCode) > 0 {
		if err := assignInputField(input, "Code", _lambdaCode); err != nil {
			log.Errorf("invalid --code: %s", err.Error())
			return
		}
	}
	if len(_lambdaFunctionName) > 0 {
		input.FunctionName = aws.String(_lambdaFunctionName)
	}
	if len(_lambdaRole) > 0 {
		input.Role = aws.String(_lambdaRole)
	}
	if len(_lambdaArchitectures) > 0 {
		if err := assignInputField(input, "Architectures", _lambdaArchitectures); err != nil {
			log.Errorf("invalid --architectures: %s", err.Error())
			return
		}
	}
	if len(_lambdaCapacityProviderConfig) > 0 {
		if err := assignInputField(input, "CapacityProviderConfig", _lambdaCapacityProviderConfig); err != nil {
			log.Errorf("invalid --capacity-provider-config: %s", err.Error())
			return
		}
	}
	if len(_lambdaCodeSigningConfigArn) > 0 {
		input.CodeSigningConfigArn = aws.String(_lambdaCodeSigningConfigArn)
	}
	if len(_lambdaDeadLetterConfig) > 0 {
		if err := assignInputField(input, "DeadLetterConfig", _lambdaDeadLetterConfig); err != nil {
			log.Errorf("invalid --dead-letter-config: %s", err.Error())
			return
		}
	}
	if len(_lambdaDescription) > 0 {
		input.Description = aws.String(_lambdaDescription)
	}
	if len(_lambdaDurableConfig) > 0 {
		if err := assignInputField(input, "DurableConfig", _lambdaDurableConfig); err != nil {
			log.Errorf("invalid --durable-config: %s", err.Error())
			return
		}
	}
	if len(_lambdaEnvironment) > 0 {
		if err := assignInputField(input, "Environment", _lambdaEnvironment); err != nil {
			log.Errorf("invalid --environment: %s", err.Error())
			return
		}
	}
	if len(_lambdaEphemeralStorage) > 0 {
		if err := assignInputField(input, "EphemeralStorage", _lambdaEphemeralStorage); err != nil {
			log.Errorf("invalid --ephemeral-storage: %s", err.Error())
			return
		}
	}
	if len(_lambdaFileSystemConfigs) > 0 {
		if err := assignInputField(input, "FileSystemConfigs", _lambdaFileSystemConfigs); err != nil {
			log.Errorf("invalid --file-system-configs: %s", err.Error())
			return
		}
	}
	if len(_lambdaHandler) > 0 {
		input.Handler = aws.String(_lambdaHandler)
	}
	if len(_lambdaImageConfig) > 0 {
		if err := assignInputField(input, "ImageConfig", _lambdaImageConfig); err != nil {
			log.Errorf("invalid --image-config: %s", err.Error())
			return
		}
	}
	if len(_lambdaKMSKeyArn) > 0 {
		input.KMSKeyArn = aws.String(_lambdaKMSKeyArn)
	}
	if len(_lambdaLayers) > 0 {
		input.Layers = append([]string(nil), _lambdaLayers...)
	}
	if len(_lambdaLoggingConfig) > 0 {
		if err := assignInputField(input, "LoggingConfig", _lambdaLoggingConfig); err != nil {
			log.Errorf("invalid --logging-config: %s", err.Error())
			return
		}
	}
	if len(_lambdaMemorySize) > 0 {
		if err := assignInputField(input, "MemorySize", _lambdaMemorySize); err != nil {
			log.Errorf("invalid --memory-size: %s", err.Error())
			return
		}
	}
	if len(_lambdaPackageType) > 0 {
		if err := assignInputField(input, "PackageType", _lambdaPackageType); err != nil {
			log.Errorf("invalid --package-type: %s", err.Error())
			return
		}
	}
	if len(_lambdaPublish) > 0 {
		if err := assignInputField(input, "Publish", _lambdaPublish); err != nil {
			log.Errorf("invalid --publish: %s", err.Error())
			return
		}
	}
	if len(_lambdaPublishTo) > 0 {
		if err := assignInputField(input, "PublishTo", _lambdaPublishTo); err != nil {
			log.Errorf("invalid --publish-to: %s", err.Error())
			return
		}
	}
	if len(_lambdaRuntime) > 0 {
		if err := assignInputField(input, "Runtime", _lambdaRuntime); err != nil {
			log.Errorf("invalid --runtime: %s", err.Error())
			return
		}
	}
	if len(_lambdaSnapStart) > 0 {
		if err := assignInputField(input, "SnapStart", _lambdaSnapStart); err != nil {
			log.Errorf("invalid --snap-start: %s", err.Error())
			return
		}
	}
	if len(_lambdaTags) > 0 {
		if err := assignInputField(input, "Tags", _lambdaTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_lambdaTenancyConfig) > 0 {
		if err := assignInputField(input, "TenancyConfig", _lambdaTenancyConfig); err != nil {
			log.Errorf("invalid --tenancy-config: %s", err.Error())
			return
		}
	}
	if len(_lambdaTimeout) > 0 {
		if err := assignInputField(input, "Timeout", _lambdaTimeout); err != nil {
			log.Errorf("invalid --timeout: %s", err.Error())
			return
		}
	}
	if len(_lambdaTracingConfig) > 0 {
		if err := assignInputField(input, "TracingConfig", _lambdaTracingConfig); err != nil {
			log.Errorf("invalid --tracing-config: %s", err.Error())
			return
		}
	}
	if len(_lambdaVpcConfig) > 0 {
		if err := assignInputField(input, "VpcConfig", _lambdaVpcConfig); err != nil {
			log.Errorf("invalid --vpc-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateFunction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a Lambda function URL with the specified configuration parameters. A
// function URL is a dedicated HTTP(S) endpoint that you can use to invoke your
// function.
func lambda_CreateFunctionUrlConfig(cfg aws.Config, client *lambda.Client) {
	input := &lambda.CreateFunctionUrlConfigInput{
		// AuthType: types.FunctionUrlAuthType, // Required
		// FunctionName: *string, // Required
	}

	if len(_lambdaAuthType) > 0 {
		if err := assignInputField(input, "AuthType", _lambdaAuthType); err != nil {
			log.Errorf("invalid --auth-type: %s", err.Error())
			return
		}
	}
	if len(_lambdaFunctionName) > 0 {
		input.FunctionName = aws.String(_lambdaFunctionName)
	}
	if len(_lambdaCors) > 0 {
		if err := assignInputField(input, "Cors", _lambdaCors); err != nil {
			log.Errorf("invalid --cors: %s", err.Error())
			return
		}
	}
	if len(_lambdaInvokeMode) > 0 {
		if err := assignInputField(input, "InvokeMode", _lambdaInvokeMode); err != nil {
			log.Errorf("invalid --invoke-mode: %s", err.Error())
			return
		}
	}
	if len(_lambdaQualifier) > 0 {
		input.Qualifier = aws.String(_lambdaQualifier)
	}

	if resp, err := client.CreateFunctionUrlConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a Lambda function [alias].
//
// [alias]: https://docs.aws.amazon.com/lambda/latest/dg/configuration-aliases.html
func lambda_DeleteAlias(cfg aws.Config, client *lambda.Client) {
	input := &lambda.DeleteAliasInput{
		// FunctionName: *string, // Required
		// Name: *string, // Required
	}

	if len(_lambdaFunctionName) > 0 {
		input.FunctionName = aws.String(_lambdaFunctionName)
	}
	if len(_lambdaName) > 0 {
		input.Name = aws.String(_lambdaName)
	}

	if resp, err := client.DeleteAlias(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a capacity provider. You cannot delete a capacity provider that is
// currently being used by Lambda functions.
func lambda_DeleteCapacityProvider(cfg aws.Config, client *lambda.Client) {
	input := &lambda.DeleteCapacityProviderInput{
		// CapacityProviderName: *string, // Required
	}

	if len(_lambdaCapacityProviderName) > 0 {
		input.CapacityProviderName = aws.String(_lambdaCapacityProviderName)
	}

	if resp, err := client.DeleteCapacityProvider(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the code signing configuration. You can delete the code signing
// configuration only if no function is using it.
func lambda_DeleteCodeSigningConfig(cfg aws.Config, client *lambda.Client) {
	input := &lambda.DeleteCodeSigningConfigInput{
		// CodeSigningConfigArn: *string, // Required
	}

	if len(_lambdaCodeSigningConfigArn) > 0 {
		input.CodeSigningConfigArn = aws.String(_lambdaCodeSigningConfigArn)
	}

	if resp, err := client.DeleteCodeSigningConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an [event source mapping]. You can get the identifier of a mapping from the output of ListEventSourceMappings.
// When you delete an event source mapping, it enters a Deleting state and might
// not be completely deleted for several seconds.
//
// [event source mapping]: https://docs.aws.amazon.com/lambda/latest/dg/intro-invocation-modes.html
func lambda_DeleteEventSourceMapping(cfg aws.Config, client *lambda.Client) {
	input := &lambda.DeleteEventSourceMappingInput{
		// UUID: *string, // Required
	}

	if len(_lambdaUUID) > 0 {
		input.UUID = aws.String(_lambdaUUID)
	}

	if resp, err := client.DeleteEventSourceMapping(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a Lambda function. To delete a specific function version, use the
// Qualifier parameter. Otherwise, all versions and aliases are deleted. This
// doesn't require the user to have explicit permissions for DeleteAlias.
//
// A deleted Lambda function cannot be recovered. Ensure that you specify the
// correct function name and version before deleting.
//
// To delete Lambda event source mappings that invoke a function, use DeleteEventSourceMapping. For Amazon
// Web Services services and resources that invoke your function directly, delete
// the trigger in the service where you originally configured it.
func lambda_DeleteFunction(cfg aws.Config, client *lambda.Client) {
	input := &lambda.DeleteFunctionInput{
		// FunctionName: *string, // Required
	}

	if len(_lambdaFunctionName) > 0 {
		input.FunctionName = aws.String(_lambdaFunctionName)
	}
	if len(_lambdaQualifier) > 0 {
		input.Qualifier = aws.String(_lambdaQualifier)
	}

	if resp, err := client.DeleteFunction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the code signing configuration from the function.
func lambda_DeleteFunctionCodeSigningConfig(cfg aws.Config, client *lambda.Client) {
	input := &lambda.DeleteFunctionCodeSigningConfigInput{
		// FunctionName: *string, // Required
	}

	if len(_lambdaFunctionName) > 0 {
		input.FunctionName = aws.String(_lambdaFunctionName)
	}

	if resp, err := client.DeleteFunctionCodeSigningConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a concurrent execution limit from a function.
func lambda_DeleteFunctionConcurrency(cfg aws.Config, client *lambda.Client) {
	input := &lambda.DeleteFunctionConcurrencyInput{
		// FunctionName: *string, // Required
	}

	if len(_lambdaFunctionName) > 0 {
		input.FunctionName = aws.String(_lambdaFunctionName)
	}

	if resp, err := client.DeleteFunctionConcurrency(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the configuration for asynchronous invocation for a function, version,
// or alias.
//
// To configure options for asynchronous invocation, use PutFunctionEventInvokeConfig.
func lambda_DeleteFunctionEventInvokeConfig(cfg aws.Config, client *lambda.Client) {
	input := &lambda.DeleteFunctionEventInvokeConfigInput{
		// FunctionName: *string, // Required
	}

	if len(_lambdaFunctionName) > 0 {
		input.FunctionName = aws.String(_lambdaFunctionName)
	}
	if len(_lambdaQualifier) > 0 {
		input.Qualifier = aws.String(_lambdaQualifier)
	}

	if resp, err := client.DeleteFunctionEventInvokeConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a Lambda function URL. When you delete a function URL, you can't
// recover it. Creating a new function URL results in a different URL address.
func lambda_DeleteFunctionUrlConfig(cfg aws.Config, client *lambda.Client) {
	input := &lambda.DeleteFunctionUrlConfigInput{
		// FunctionName: *string, // Required
	}

	if len(_lambdaFunctionName) > 0 {
		input.FunctionName = aws.String(_lambdaFunctionName)
	}
	if len(_lambdaQualifier) > 0 {
		input.Qualifier = aws.String(_lambdaQualifier)
	}

	if resp, err := client.DeleteFunctionUrlConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a version of an [Lambda layer]. Deleted versions can no longer be viewed or added to
// functions. To avoid breaking functions, a copy of the version remains in Lambda
// until no functions refer to it.
//
// [Lambda layer]: https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html
func lambda_DeleteLayerVersion(cfg aws.Config, client *lambda.Client) {
	input := &lambda.DeleteLayerVersionInput{
		// LayerName: *string, // Required
		// VersionNumber: *int64, // Required
	}

	if len(_lambdaLayerName) > 0 {
		input.LayerName = aws.String(_lambdaLayerName)
	}
	if len(_lambdaVersionNumber) > 0 {
		if err := assignInputField(input, "VersionNumber", _lambdaVersionNumber); err != nil {
			log.Errorf("invalid --version-number: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteLayerVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the provisioned concurrency configuration for a function.
func lambda_DeleteProvisionedConcurrencyConfig(cfg aws.Config, client *lambda.Client) {
	input := &lambda.DeleteProvisionedConcurrencyConfigInput{
		// FunctionName: *string, // Required
		// Qualifier: *string, // Required
	}

	if len(_lambdaFunctionName) > 0 {
		input.FunctionName = aws.String(_lambdaFunctionName)
	}
	if len(_lambdaQualifier) > 0 {
		input.Qualifier = aws.String(_lambdaQualifier)
	}

	if resp, err := client.DeleteProvisionedConcurrencyConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves details about your account's [limits] and usage in an Amazon Web Services
// Region.
//
// [limits]: https://docs.aws.amazon.com/lambda/latest/dg/limits.html
func lambda_GetAccountSettings(cfg aws.Config, client *lambda.Client) {
	input := &lambda.GetAccountSettingsInput{}

	if resp, err := client.GetAccountSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns details about a Lambda function [alias].
//
// [alias]: https://docs.aws.amazon.com/lambda/latest/dg/configuration-aliases.html
func lambda_GetAlias(cfg aws.Config, client *lambda.Client) {
	input := &lambda.GetAliasInput{
		// FunctionName: *string, // Required
		// Name: *string, // Required
	}

	if len(_lambdaFunctionName) > 0 {
		input.FunctionName = aws.String(_lambdaFunctionName)
	}
	if len(_lambdaName) > 0 {
		input.Name = aws.String(_lambdaName)
	}

	if resp, err := client.GetAlias(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a specific capacity provider, including its
// configuration, state, and associated resources.
func lambda_GetCapacityProvider(cfg aws.Config, client *lambda.Client) {
	input := &lambda.GetCapacityProviderInput{
		// CapacityProviderName: *string, // Required
	}

	if len(_lambdaCapacityProviderName) > 0 {
		input.CapacityProviderName = aws.String(_lambdaCapacityProviderName)
	}

	if resp, err := client.GetCapacityProvider(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the specified code signing configuration.
func lambda_GetCodeSigningConfig(cfg aws.Config, client *lambda.Client) {
	input := &lambda.GetCodeSigningConfigInput{
		// CodeSigningConfigArn: *string, // Required
	}

	if len(_lambdaCodeSigningConfigArn) > 0 {
		input.CodeSigningConfigArn = aws.String(_lambdaCodeSigningConfigArn)
	}

	if resp, err := client.GetCodeSigningConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves detailed information about a specific [durable execution], including its current status,
// input payload, result or error information, and execution metadata such as start
// time and usage statistics.
//
// [durable execution]: https://docs.aws.amazon.com/lambda/latest/dg/durable-functions.html
func lambda_GetDurableExecution(cfg aws.Config, client *lambda.Client) {
	input := &lambda.GetDurableExecutionInput{
		// DurableExecutionArn: *string, // Required
	}

	if len(_lambdaDurableExecutionArn) > 0 {
		input.DurableExecutionArn = aws.String(_lambdaDurableExecutionArn)
	}

	if resp, err := client.GetDurableExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the execution history for a [durable execution], showing all the steps, callbacks, and
// events that occurred during the execution. This provides a detailed audit trail
// of the execution's progress over time.
//
// The history is available while the execution is running and for a retention
// period after it completes (1-90 days, default 30 days). You can control whether
// to include execution data such as step results and callback payloads.
//
// [durable execution]: https://docs.aws.amazon.com/lambda/latest/dg/durable-functions.html
func lambda_GetDurableExecutionHistory(cfg aws.Config, client *lambda.Client) {
	input := &lambda.GetDurableExecutionHistoryInput{
		// DurableExecutionArn: *string, // Required
	}

	if len(_lambdaDurableExecutionArn) > 0 {
		input.DurableExecutionArn = aws.String(_lambdaDurableExecutionArn)
	}
	if len(_lambdaIncludeExecutionData) > 0 {
		if err := assignInputField(input, "IncludeExecutionData", _lambdaIncludeExecutionData); err != nil {
			log.Errorf("invalid --include-execution-data: %s", err.Error())
			return
		}
	}
	if len(_lambdaMarker) > 0 {
		input.Marker = aws.String(_lambdaMarker)
	}
	if len(_lambdaMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _lambdaMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}
	if len(_lambdaReverseOrder) > 0 {
		if err := assignInputField(input, "ReverseOrder", _lambdaReverseOrder); err != nil {
			log.Errorf("invalid --reverse-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.GetDurableExecutionHistory(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lambda.GetDurableExecutionHistoryOutput
	p := lambda.NewGetDurableExecutionHistoryPaginator(client, input)
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

// Retrieves the current execution state required for the replay process during [durable function]
// execution. This API is used by the Lambda durable functions SDK to get state
// information needed for replay. You typically don't need to call this API
// directly as the SDK handles state management automatically.
//
// The response contains operations ordered by start sequence number in ascending
// order. Completed operations with children don't include child operation details
// since they don't need to be replayed.
//
// [durable function]: https://docs.aws.amazon.com/lambda/latest/dg/durable-functions.html
func lambda_GetDurableExecutionState(cfg aws.Config, client *lambda.Client) {
	input := &lambda.GetDurableExecutionStateInput{
		// CheckpointToken: *string, // Required
		// DurableExecutionArn: *string, // Required
	}

	if len(_lambdaCheckpointToken) > 0 {
		input.CheckpointToken = aws.String(_lambdaCheckpointToken)
	}
	if len(_lambdaDurableExecutionArn) > 0 {
		input.DurableExecutionArn = aws.String(_lambdaDurableExecutionArn)
	}
	if len(_lambdaMarker) > 0 {
		input.Marker = aws.String(_lambdaMarker)
	}
	if len(_lambdaMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _lambdaMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.GetDurableExecutionState(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lambda.GetDurableExecutionStateOutput
	p := lambda.NewGetDurableExecutionStatePaginator(client, input)
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

// Returns details about an event source mapping. You can get the identifier of a
// mapping from the output of ListEventSourceMappings.
func lambda_GetEventSourceMapping(cfg aws.Config, client *lambda.Client) {
	input := &lambda.GetEventSourceMappingInput{
		// UUID: *string, // Required
	}

	if len(_lambdaUUID) > 0 {
		input.UUID = aws.String(_lambdaUUID)
	}

	if resp, err := client.GetEventSourceMapping(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the function or function version, with a link to
// download the deployment package that's valid for 10 minutes. If you specify a
// function version, only details that are specific to that version are returned.
func lambda_GetFunction(cfg aws.Config, client *lambda.Client) {
	input := &lambda.GetFunctionInput{
		// FunctionName: *string, // Required
	}

	if len(_lambdaFunctionName) > 0 {
		input.FunctionName = aws.String(_lambdaFunctionName)
	}
	if len(_lambdaQualifier) > 0 {
		input.Qualifier = aws.String(_lambdaQualifier)
	}

	if resp, err := client.GetFunction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the code signing configuration for the specified function.
func lambda_GetFunctionCodeSigningConfig(cfg aws.Config, client *lambda.Client) {
	input := &lambda.GetFunctionCodeSigningConfigInput{
		// FunctionName: *string, // Required
	}

	if len(_lambdaFunctionName) > 0 {
		input.FunctionName = aws.String(_lambdaFunctionName)
	}

	if resp, err := client.GetFunctionCodeSigningConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns details about the reserved concurrency configuration for a function. To
// set a concurrency limit for a function, use PutFunctionConcurrency.
func lambda_GetFunctionConcurrency(cfg aws.Config, client *lambda.Client) {
	input := &lambda.GetFunctionConcurrencyInput{
		// FunctionName: *string, // Required
	}

	if len(_lambdaFunctionName) > 0 {
		input.FunctionName = aws.String(_lambdaFunctionName)
	}

	if resp, err := client.GetFunctionConcurrency(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the version-specific settings of a Lambda function or version. The
// output includes only options that can vary between versions of a function. To
// modify these settings, use UpdateFunctionConfiguration.
//
// To get all of a function's details, including function-level settings, use GetFunction.
func lambda_GetFunctionConfiguration(cfg aws.Config, client *lambda.Client) {
	input := &lambda.GetFunctionConfigurationInput{
		// FunctionName: *string, // Required
	}

	if len(_lambdaFunctionName) > 0 {
		input.FunctionName = aws.String(_lambdaFunctionName)
	}
	if len(_lambdaQualifier) > 0 {
		input.Qualifier = aws.String(_lambdaQualifier)
	}

	if resp, err := client.GetFunctionConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the configuration for asynchronous invocation for a function,
// version, or alias.
//
// To configure options for asynchronous invocation, use PutFunctionEventInvokeConfig.
func lambda_GetFunctionEventInvokeConfig(cfg aws.Config, client *lambda.Client) {
	input := &lambda.GetFunctionEventInvokeConfigInput{
		// FunctionName: *string, // Required
	}

	if len(_lambdaFunctionName) > 0 {
		input.FunctionName = aws.String(_lambdaFunctionName)
	}
	if len(_lambdaQualifier) > 0 {
		input.Qualifier = aws.String(_lambdaQualifier)
	}

	if resp, err := client.GetFunctionEventInvokeConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns your function's [recursive loop detection] configuration.
//
// [recursive loop detection]: https://docs.aws.amazon.com/lambda/latest/dg/invocation-recursion.html
func lambda_GetFunctionRecursionConfig(cfg aws.Config, client *lambda.Client) {
	input := &lambda.GetFunctionRecursionConfigInput{
		// FunctionName: *string, // Required
	}

	if len(_lambdaFunctionName) > 0 {
		input.FunctionName = aws.String(_lambdaFunctionName)
	}

	if resp, err := client.GetFunctionRecursionConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the scaling configuration for a Lambda Managed Instances function.
func lambda_GetFunctionScalingConfig(cfg aws.Config, client *lambda.Client) {
	input := &lambda.GetFunctionScalingConfigInput{
		// FunctionName: *string, // Required
		// Qualifier: *string, // Required
	}

	if len(_lambdaFunctionName) > 0 {
		input.FunctionName = aws.String(_lambdaFunctionName)
	}
	if len(_lambdaQualifier) > 0 {
		input.Qualifier = aws.String(_lambdaQualifier)
	}

	if resp, err := client.GetFunctionScalingConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns details about a Lambda function URL.
func lambda_GetFunctionUrlConfig(cfg aws.Config, client *lambda.Client) {
	input := &lambda.GetFunctionUrlConfigInput{
		// FunctionName: *string, // Required
	}

	if len(_lambdaFunctionName) > 0 {
		input.FunctionName = aws.String(_lambdaFunctionName)
	}
	if len(_lambdaQualifier) > 0 {
		input.Qualifier = aws.String(_lambdaQualifier)
	}

	if resp, err := client.GetFunctionUrlConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a version of an [Lambda layer], with a link to download the layer
// archive that's valid for 10 minutes.
//
// [Lambda layer]: https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html
func lambda_GetLayerVersion(cfg aws.Config, client *lambda.Client) {
	input := &lambda.GetLayerVersionInput{
		// LayerName: *string, // Required
		// VersionNumber: *int64, // Required
	}

	if len(_lambdaLayerName) > 0 {
		input.LayerName = aws.String(_lambdaLayerName)
	}
	if len(_lambdaVersionNumber) > 0 {
		if err := assignInputField(input, "VersionNumber", _lambdaVersionNumber); err != nil {
			log.Errorf("invalid --version-number: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetLayerVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a version of an [Lambda layer], with a link to download the layer
// archive that's valid for 10 minutes.
//
// [Lambda layer]: https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html
func lambda_GetLayerVersionByArn(cfg aws.Config, client *lambda.Client) {
	input := &lambda.GetLayerVersionByArnInput{
		// Arn: *string, // Required
	}

	if len(_lambdaArn) > 0 {
		input.Arn = aws.String(_lambdaArn)
	}

	if resp, err := client.GetLayerVersionByArn(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the permission policy for a version of an [Lambda layer]. For more information, see AddLayerVersionPermission.
//
// [Lambda layer]: https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html
func lambda_GetLayerVersionPolicy(cfg aws.Config, client *lambda.Client) {
	input := &lambda.GetLayerVersionPolicyInput{
		// LayerName: *string, // Required
		// VersionNumber: *int64, // Required
	}

	if len(_lambdaLayerName) > 0 {
		input.LayerName = aws.String(_lambdaLayerName)
	}
	if len(_lambdaVersionNumber) > 0 {
		if err := assignInputField(input, "VersionNumber", _lambdaVersionNumber); err != nil {
			log.Errorf("invalid --version-number: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetLayerVersionPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the [resource-based IAM policy] for a function, version, or alias.
//
// [resource-based IAM policy]: https://docs.aws.amazon.com/lambda/latest/dg/access-control-resource-based.html
func lambda_GetPolicy(cfg aws.Config, client *lambda.Client) {
	input := &lambda.GetPolicyInput{
		// FunctionName: *string, // Required
	}

	if len(_lambdaFunctionName) > 0 {
		input.FunctionName = aws.String(_lambdaFunctionName)
	}
	if len(_lambdaQualifier) > 0 {
		input.Qualifier = aws.String(_lambdaQualifier)
	}

	if resp, err := client.GetPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the provisioned concurrency configuration for a function's alias or
// version.
func lambda_GetProvisionedConcurrencyConfig(cfg aws.Config, client *lambda.Client) {
	input := &lambda.GetProvisionedConcurrencyConfigInput{
		// FunctionName: *string, // Required
		// Qualifier: *string, // Required
	}

	if len(_lambdaFunctionName) > 0 {
		input.FunctionName = aws.String(_lambdaFunctionName)
	}
	if len(_lambdaQualifier) > 0 {
		input.Qualifier = aws.String(_lambdaQualifier)
	}

	if resp, err := client.GetProvisionedConcurrencyConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the runtime management configuration for a function's version. If the
// runtime update mode is Manual, this includes the ARN of the runtime version and
// the runtime update mode. If the runtime update mode is Auto or Function update,
// this includes the runtime update mode and null is returned for the ARN. For
// more information, see [Runtime updates].
//
// [Runtime updates]: https://docs.aws.amazon.com/lambda/latest/dg/runtimes-update.html
func lambda_GetRuntimeManagementConfig(cfg aws.Config, client *lambda.Client) {
	input := &lambda.GetRuntimeManagementConfigInput{
		// FunctionName: *string, // Required
	}

	if len(_lambdaFunctionName) > 0 {
		input.FunctionName = aws.String(_lambdaFunctionName)
	}
	if len(_lambdaQualifier) > 0 {
		input.Qualifier = aws.String(_lambdaQualifier)
	}

	if resp, err := client.GetRuntimeManagementConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Invokes a Lambda function. You can invoke a function synchronously (and wait
// for the response), or asynchronously. By default, Lambda invokes your function
// synchronously (i.e. the InvocationType is RequestResponse ). To invoke a
// function asynchronously, set InvocationType to Event . Lambda passes the
// ClientContext object to your function for synchronous invocations only.
//
// For synchronous invocations, the maximum payload size is 6 MB. For asynchronous
// invocations, the maximum payload size is 1 MB.
//
// For [synchronous invocation], details about the function response, including errors, are included in
// the response body and headers. For either invocation type, you can find more
// information in the [execution log]and [trace].
//
// When an error occurs, your function may be invoked multiple times. Retry
// behavior varies by error type, client, event source, and invocation type. For
// example, if you invoke a function asynchronously and it returns an error, Lambda
// executes the function up to two more times. For more information, see [Error handling and automatic retries in Lambda].
//
// For [asynchronous invocation], Lambda adds events to a queue before sending them to your function. If
// your function does not have enough capacity to keep up with the queue, events
// may be lost. Occasionally, your function may receive the same event multiple
// times, even if no error occurs. To retain events that were not processed,
// configure your function with a [dead-letter queue].
//
// The status code in the API response doesn't reflect function errors. Error
// codes are reserved for errors that prevent your function from executing, such as
// permissions errors, [quota]errors, or issues with your function's code and
// configuration. For example, Lambda returns TooManyRequestsException if running
// the function would cause you to exceed a concurrency limit at either the account
// level ( ConcurrentInvocationLimitExceeded ) or function level (
// ReservedFunctionConcurrentInvocationLimitExceeded ).
//
// For functions with a long timeout, your client might disconnect during
// synchronous invocation while it waits for a response. Configure your HTTP
// client, SDK, firewall, proxy, or operating system to allow for long connections
// with timeout or keep-alive settings.
//
// This operation requires permission for the [lambda:InvokeFunction] action. For details on how to set
// up permissions for cross-account invocations, see [Granting function access to other accounts].
//
// [execution log]: https://docs.aws.amazon.com/lambda/latest/dg/monitoring-functions.html
// [asynchronous invocation]: https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html
// [trace]: https://docs.aws.amazon.com/lambda/latest/dg/lambda-x-ray.html
// [dead-letter queue]: https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html#invocation-dlq
// [Error handling and automatic retries in Lambda]: https://docs.aws.amazon.com/lambda/latest/dg/invocation-retries.html
// [lambda:InvokeFunction]: https://docs.aws.amazon.com/IAM/latest/UserGuide/list_awslambda.html
// [quota]: https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-limits.html
// [synchronous invocation]: https://docs.aws.amazon.com/lambda/latest/dg/invocation-sync.html
// [Granting function access to other accounts]: https://docs.aws.amazon.com/lambda/latest/dg/access-control-resource-based.html#permissions-resource-xaccountinvoke
func lambda_Invoke(cfg aws.Config, client *lambda.Client) {
	input := &lambda.InvokeInput{
		// FunctionName: *string, // Required
	}

	if len(_lambdaFunctionName) > 0 {
		input.FunctionName = aws.String(_lambdaFunctionName)
	}
	if len(_lambdaClientContext) > 0 {
		input.ClientContext = aws.String(_lambdaClientContext)
	}
	if len(_lambdaDurableExecutionName) > 0 {
		input.DurableExecutionName = aws.String(_lambdaDurableExecutionName)
	}
	if len(_lambdaInvocationType) > 0 {
		if err := assignInputField(input, "InvocationType", _lambdaInvocationType); err != nil {
			log.Errorf("invalid --invocation-type: %s", err.Error())
			return
		}
	}
	if len(_lambdaLogType) > 0 {
		if err := assignInputField(input, "LogType", _lambdaLogType); err != nil {
			log.Errorf("invalid --log-type: %s", err.Error())
			return
		}
	}
	if len(_lambdaPayload) > 0 {
		if err := assignInputField(input, "Payload", _lambdaPayload); err != nil {
			log.Errorf("invalid --payload: %s", err.Error())
			return
		}
	}
	if len(_lambdaQualifier) > 0 {
		input.Qualifier = aws.String(_lambdaQualifier)
	}
	if len(_lambdaTenantId) > 0 {
		input.TenantId = aws.String(_lambdaTenantId)
	}

	if resp, err := client.Invoke(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// For asynchronous function invocation, use Invoke.
// Invokes a function asynchronously.
//
// The payload limit is 256KB. For larger payloads, for up to 1MB, use Invoke.
//
// If you do use the InvokeAsync action, note that it doesn't support the use of
// X-Ray active tracing. Trace ID is not propagated to the function, even if X-Ray
// active tracing is turned on.
//
// Deprecated: This operation has been deprecated.
func lambda_InvokeAsync(cfg aws.Config, client *lambda.Client) {
	input := &lambda.InvokeAsyncInput{
		// FunctionName: *string, // Required
		// InvokeArgs: io.Reader, // Required
	}

	if len(_lambdaFunctionName) > 0 {
		input.FunctionName = aws.String(_lambdaFunctionName)
	}
	if len(_lambdaInvokeArgs) > 0 {
		if err := assignInputField(input, "InvokeArgs", _lambdaInvokeArgs); err != nil {
			log.Errorf("invalid --invoke-args: %s", err.Error())
			return
		}
	}

	if resp, err := client.InvokeAsync(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Configure your Lambda functions to stream response payloads back to clients.
// For more information, see [Configuring a Lambda function to stream responses].
//
// This operation requires permission for the [lambda:InvokeFunction] action. For details on how to set
// up permissions for cross-account invocations, see [Granting function access to other accounts].
//
// [Configuring a Lambda function to stream responses]: https://docs.aws.amazon.com/lambda/latest/dg/configuration-response-streaming.html
// [lambda:InvokeFunction]: https://docs.aws.amazon.com/IAM/latest/UserGuide/list_awslambda.html
// [Granting function access to other accounts]: https://docs.aws.amazon.com/lambda/latest/dg/access-control-resource-based.html#permissions-resource-xaccountinvoke
func lambda_InvokeWithResponseStream(cfg aws.Config, client *lambda.Client) {
	input := &lambda.InvokeWithResponseStreamInput{
		// FunctionName: *string, // Required
	}

	if len(_lambdaFunctionName) > 0 {
		input.FunctionName = aws.String(_lambdaFunctionName)
	}
	if len(_lambdaClientContext) > 0 {
		input.ClientContext = aws.String(_lambdaClientContext)
	}
	if len(_lambdaInvocationType) > 0 {
		if err := assignInputField(input, "InvocationType", _lambdaInvocationType); err != nil {
			log.Errorf("invalid --invocation-type: %s", err.Error())
			return
		}
	}
	if len(_lambdaLogType) > 0 {
		if err := assignInputField(input, "LogType", _lambdaLogType); err != nil {
			log.Errorf("invalid --log-type: %s", err.Error())
			return
		}
	}
	if len(_lambdaPayload) > 0 {
		if err := assignInputField(input, "Payload", _lambdaPayload); err != nil {
			log.Errorf("invalid --payload: %s", err.Error())
			return
		}
	}
	if len(_lambdaQualifier) > 0 {
		input.Qualifier = aws.String(_lambdaQualifier)
	}
	if len(_lambdaTenantId) > 0 {
		input.TenantId = aws.String(_lambdaTenantId)
	}

	if resp, err := client.InvokeWithResponseStream(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of [aliases] for a Lambda function.
//
// [aliases]: https://docs.aws.amazon.com/lambda/latest/dg/configuration-aliases.html
func lambda_ListAliases(cfg aws.Config, client *lambda.Client) {
	input := &lambda.ListAliasesInput{
		// FunctionName: *string, // Required
	}

	if len(_lambdaFunctionName) > 0 {
		input.FunctionName = aws.String(_lambdaFunctionName)
	}
	if len(_lambdaFunctionVersion) > 0 {
		input.FunctionVersion = aws.String(_lambdaFunctionVersion)
	}
	if len(_lambdaMarker) > 0 {
		input.Marker = aws.String(_lambdaMarker)
	}
	if len(_lambdaMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _lambdaMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListAliases(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lambda.ListAliasesOutput
	p := lambda.NewListAliasesPaginator(client, input)
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

// Returns a list of capacity providers in your account.
func lambda_ListCapacityProviders(cfg aws.Config, client *lambda.Client) {
	input := &lambda.ListCapacityProvidersInput{}

	if len(_lambdaMarker) > 0 {
		input.Marker = aws.String(_lambdaMarker)
	}
	if len(_lambdaMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _lambdaMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}
	if len(_lambdaState) > 0 {
		if err := assignInputField(input, "State", _lambdaState); err != nil {
			log.Errorf("invalid --state: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListCapacityProviders(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lambda.ListCapacityProvidersOutput
	p := lambda.NewListCapacityProvidersPaginator(client, input)
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

// Returns a list of [code signing configurations]. A request returns up to 10,000 configurations per call. You
// can use the MaxItems parameter to return fewer configurations per call.
//
// [code signing configurations]: https://docs.aws.amazon.com/lambda/latest/dg/configuring-codesigning.html
func lambda_ListCodeSigningConfigs(cfg aws.Config, client *lambda.Client) {
	input := &lambda.ListCodeSigningConfigsInput{}

	if len(_lambdaMarker) > 0 {
		input.Marker = aws.String(_lambdaMarker)
	}
	if len(_lambdaMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _lambdaMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListCodeSigningConfigs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lambda.ListCodeSigningConfigsOutput
	p := lambda.NewListCodeSigningConfigsPaginator(client, input)
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

// Returns a list of [durable executions] for a specified Lambda function. You can filter the results
// by execution name, status, and start time range. This API supports pagination
// for large result sets.
//
// [durable executions]: https://docs.aws.amazon.com/lambda/latest/dg/durable-functions.html
func lambda_ListDurableExecutionsByFunction(cfg aws.Config, client *lambda.Client) {
	input := &lambda.ListDurableExecutionsByFunctionInput{
		// FunctionName: *string, // Required
	}

	if len(_lambdaFunctionName) > 0 {
		input.FunctionName = aws.String(_lambdaFunctionName)
	}
	if len(_lambdaDurableExecutionName) > 0 {
		input.DurableExecutionName = aws.String(_lambdaDurableExecutionName)
	}
	if len(_lambdaMarker) > 0 {
		input.Marker = aws.String(_lambdaMarker)
	}
	if len(_lambdaMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _lambdaMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}
	if len(_lambdaQualifier) > 0 {
		input.Qualifier = aws.String(_lambdaQualifier)
	}
	if len(_lambdaReverseOrder) > 0 {
		if err := assignInputField(input, "ReverseOrder", _lambdaReverseOrder); err != nil {
			log.Errorf("invalid --reverse-order: %s", err.Error())
			return
		}
	}
	if len(_lambdaStartedAfter) > 0 {
		if err := assignInputField(input, "StartedAfter", _lambdaStartedAfter); err != nil {
			log.Errorf("invalid --started-after: %s", err.Error())
			return
		}
	}
	if len(_lambdaStartedBefore) > 0 {
		if err := assignInputField(input, "StartedBefore", _lambdaStartedBefore); err != nil {
			log.Errorf("invalid --started-before: %s", err.Error())
			return
		}
	}
	if len(_lambdaStatuses) > 0 {
		if err := assignInputField(input, "Statuses", _lambdaStatuses); err != nil {
			log.Errorf("invalid --statuses: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListDurableExecutionsByFunction(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lambda.ListDurableExecutionsByFunctionOutput
	p := lambda.NewListDurableExecutionsByFunctionPaginator(client, input)
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

// Lists event source mappings. Specify an EventSourceArn to show only event
// source mappings for a single event source.
func lambda_ListEventSourceMappings(cfg aws.Config, client *lambda.Client) {
	input := &lambda.ListEventSourceMappingsInput{}

	if len(_lambdaEventSourceArn) > 0 {
		input.EventSourceArn = aws.String(_lambdaEventSourceArn)
	}
	if len(_lambdaFunctionName) > 0 {
		input.FunctionName = aws.String(_lambdaFunctionName)
	}
	if len(_lambdaMarker) > 0 {
		input.Marker = aws.String(_lambdaMarker)
	}
	if len(_lambdaMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _lambdaMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListEventSourceMappings(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lambda.ListEventSourceMappingsOutput
	p := lambda.NewListEventSourceMappingsPaginator(client, input)
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

// Retrieves a list of configurations for asynchronous invocation for a function.
// To configure options for asynchronous invocation, use PutFunctionEventInvokeConfig.
func lambda_ListFunctionEventInvokeConfigs(cfg aws.Config, client *lambda.Client) {
	input := &lambda.ListFunctionEventInvokeConfigsInput{
		// FunctionName: *string, // Required
	}

	if len(_lambdaFunctionName) > 0 {
		input.FunctionName = aws.String(_lambdaFunctionName)
	}
	if len(_lambdaMarker) > 0 {
		input.Marker = aws.String(_lambdaMarker)
	}
	if len(_lambdaMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _lambdaMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListFunctionEventInvokeConfigs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lambda.ListFunctionEventInvokeConfigsOutput
	p := lambda.NewListFunctionEventInvokeConfigsPaginator(client, input)
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

// Returns a list of Lambda function URLs for the specified function.
func lambda_ListFunctionUrlConfigs(cfg aws.Config, client *lambda.Client) {
	input := &lambda.ListFunctionUrlConfigsInput{
		// FunctionName: *string, // Required
	}

	if len(_lambdaFunctionName) > 0 {
		input.FunctionName = aws.String(_lambdaFunctionName)
	}
	if len(_lambdaMarker) > 0 {
		input.Marker = aws.String(_lambdaMarker)
	}
	if len(_lambdaMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _lambdaMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListFunctionUrlConfigs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lambda.ListFunctionUrlConfigsOutput
	p := lambda.NewListFunctionUrlConfigsPaginator(client, input)
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

// Returns a list of function versions that are configured to use a specific
// capacity provider.
func lambda_ListFunctionVersionsByCapacityProvider(cfg aws.Config, client *lambda.Client) {
	input := &lambda.ListFunctionVersionsByCapacityProviderInput{
		// CapacityProviderName: *string, // Required
	}

	if len(_lambdaCapacityProviderName) > 0 {
		input.CapacityProviderName = aws.String(_lambdaCapacityProviderName)
	}
	if len(_lambdaMarker) > 0 {
		input.Marker = aws.String(_lambdaMarker)
	}
	if len(_lambdaMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _lambdaMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListFunctionVersionsByCapacityProvider(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lambda.ListFunctionVersionsByCapacityProviderOutput
	p := lambda.NewListFunctionVersionsByCapacityProviderPaginator(client, input)
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

// Returns a list of Lambda functions, with the version-specific configuration of
// each. Lambda returns up to 50 functions per call.
//
// Set FunctionVersion to ALL to include all published versions of each function
// in addition to the unpublished version.
//
// The ListFunctions operation returns a subset of the FunctionConfiguration fields. To get the
// additional fields (State, StateReasonCode, StateReason, LastUpdateStatus,
// LastUpdateStatusReason, LastUpdateStatusReasonCode, RuntimeVersionConfig) for a
// function or version, use GetFunction.
func lambda_ListFunctions(cfg aws.Config, client *lambda.Client) {
	input := &lambda.ListFunctionsInput{}

	if len(_lambdaFunctionVersion) > 0 {
		if err := assignInputField(input, "FunctionVersion", _lambdaFunctionVersion); err != nil {
			log.Errorf("invalid --function-version: %s", err.Error())
			return
		}
	}
	if len(_lambdaMarker) > 0 {
		input.Marker = aws.String(_lambdaMarker)
	}
	if len(_lambdaMasterRegion) > 0 {
		input.MasterRegion = aws.String(_lambdaMasterRegion)
	}
	if len(_lambdaMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _lambdaMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListFunctions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lambda.ListFunctionsOutput
	p := lambda.NewListFunctionsPaginator(client, input)
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

// List the functions that use the specified code signing configuration. You can
// use this method prior to deleting a code signing configuration, to verify that
// no functions are using it.
func lambda_ListFunctionsByCodeSigningConfig(cfg aws.Config, client *lambda.Client) {
	input := &lambda.ListFunctionsByCodeSigningConfigInput{
		// CodeSigningConfigArn: *string, // Required
	}

	if len(_lambdaCodeSigningConfigArn) > 0 {
		input.CodeSigningConfigArn = aws.String(_lambdaCodeSigningConfigArn)
	}
	if len(_lambdaMarker) > 0 {
		input.Marker = aws.String(_lambdaMarker)
	}
	if len(_lambdaMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _lambdaMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListFunctionsByCodeSigningConfig(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lambda.ListFunctionsByCodeSigningConfigOutput
	p := lambda.NewListFunctionsByCodeSigningConfigPaginator(client, input)
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

// Lists the versions of an [Lambda layer]. Versions that have been deleted aren't listed.
// Specify a [runtime identifier]to list only versions that indicate that they're compatible with that
// runtime. Specify a compatible architecture to include only layer versions that
// are compatible with that architecture.
//
// [runtime identifier]: https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html
// [Lambda layer]: https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html
func lambda_ListLayerVersions(cfg aws.Config, client *lambda.Client) {
	input := &lambda.ListLayerVersionsInput{
		// LayerName: *string, // Required
	}

	if len(_lambdaLayerName) > 0 {
		input.LayerName = aws.String(_lambdaLayerName)
	}
	if len(_lambdaCompatibleArchitecture) > 0 {
		if err := assignInputField(input, "CompatibleArchitecture", _lambdaCompatibleArchitecture); err != nil {
			log.Errorf("invalid --compatible-architecture: %s", err.Error())
			return
		}
	}
	if len(_lambdaCompatibleRuntime) > 0 {
		if err := assignInputField(input, "CompatibleRuntime", _lambdaCompatibleRuntime); err != nil {
			log.Errorf("invalid --compatible-runtime: %s", err.Error())
			return
		}
	}
	if len(_lambdaMarker) > 0 {
		input.Marker = aws.String(_lambdaMarker)
	}
	if len(_lambdaMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _lambdaMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListLayerVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lambda.ListLayerVersionsOutput
	p := lambda.NewListLayerVersionsPaginator(client, input)
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

// Lists [Lambda layers] and shows information about the latest version of each. Specify a [runtime identifier] to
// list only layers that indicate that they're compatible with that runtime.
// Specify a compatible architecture to include only layers that are compatible
// with that [instruction set architecture].
//
// [instruction set architecture]: https://docs.aws.amazon.com/lambda/latest/dg/foundation-arch.html
// [runtime identifier]: https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html
// [Lambda layers]: https://docs.aws.amazon.com/lambda/latest/dg/invocation-layers.html
func lambda_ListLayers(cfg aws.Config, client *lambda.Client) {
	input := &lambda.ListLayersInput{}

	if len(_lambdaCompatibleArchitecture) > 0 {
		if err := assignInputField(input, "CompatibleArchitecture", _lambdaCompatibleArchitecture); err != nil {
			log.Errorf("invalid --compatible-architecture: %s", err.Error())
			return
		}
	}
	if len(_lambdaCompatibleRuntime) > 0 {
		if err := assignInputField(input, "CompatibleRuntime", _lambdaCompatibleRuntime); err != nil {
			log.Errorf("invalid --compatible-runtime: %s", err.Error())
			return
		}
	}
	if len(_lambdaMarker) > 0 {
		input.Marker = aws.String(_lambdaMarker)
	}
	if len(_lambdaMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _lambdaMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListLayers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lambda.ListLayersOutput
	p := lambda.NewListLayersPaginator(client, input)
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

// Retrieves a list of provisioned concurrency configurations for a function.
func lambda_ListProvisionedConcurrencyConfigs(cfg aws.Config, client *lambda.Client) {
	input := &lambda.ListProvisionedConcurrencyConfigsInput{
		// FunctionName: *string, // Required
	}

	if len(_lambdaFunctionName) > 0 {
		input.FunctionName = aws.String(_lambdaFunctionName)
	}
	if len(_lambdaMarker) > 0 {
		input.Marker = aws.String(_lambdaMarker)
	}
	if len(_lambdaMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _lambdaMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListProvisionedConcurrencyConfigs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lambda.ListProvisionedConcurrencyConfigsOutput
	p := lambda.NewListProvisionedConcurrencyConfigsPaginator(client, input)
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

// Returns a function, event source mapping, or code signing configuration's [tags]. You
// can also view function tags with GetFunction.
//
// [tags]: https://docs.aws.amazon.com/lambda/latest/dg/tagging.html
func lambda_ListTags(cfg aws.Config, client *lambda.Client) {
	input := &lambda.ListTagsInput{
		// Resource: *string, // Required
	}

	if len(_lambdaResource) > 0 {
		input.Resource = aws.String(_lambdaResource)
	}

	if resp, err := client.ListTags(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of [versions], with the version-specific configuration of each. Lambda
// returns up to 50 versions per call.
//
// [versions]: https://docs.aws.amazon.com/lambda/latest/dg/versioning-aliases.html
func lambda_ListVersionsByFunction(cfg aws.Config, client *lambda.Client) {
	input := &lambda.ListVersionsByFunctionInput{
		// FunctionName: *string, // Required
	}

	if len(_lambdaFunctionName) > 0 {
		input.FunctionName = aws.String(_lambdaFunctionName)
	}
	if len(_lambdaMarker) > 0 {
		input.Marker = aws.String(_lambdaMarker)
	}
	if len(_lambdaMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _lambdaMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListVersionsByFunction(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lambda.ListVersionsByFunctionOutput
	p := lambda.NewListVersionsByFunctionPaginator(client, input)
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

// Creates an [Lambda layer] from a ZIP archive. Each time you call PublishLayerVersion with the
// same layer name, a new version is created.
//
// Add layers to your function with CreateFunction or UpdateFunctionConfiguration.
//
// [Lambda layer]: https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html
func lambda_PublishLayerVersion(cfg aws.Config, client *lambda.Client) {
	input := &lambda.PublishLayerVersionInput{
		// Content: *types.LayerVersionContentInput, // Required
		// LayerName: *string, // Required
	}

	if len(_lambdaContent) > 0 {
		if err := assignInputField(input, "Content", _lambdaContent); err != nil {
			log.Errorf("invalid --content: %s", err.Error())
			return
		}
	}
	if len(_lambdaLayerName) > 0 {
		input.LayerName = aws.String(_lambdaLayerName)
	}
	if len(_lambdaCompatibleArchitectures) > 0 {
		if err := assignInputField(input, "CompatibleArchitectures", _lambdaCompatibleArchitectures); err != nil {
			log.Errorf("invalid --compatible-architectures: %s", err.Error())
			return
		}
	}
	if len(_lambdaCompatibleRuntimes) > 0 {
		if err := assignInputField(input, "CompatibleRuntimes", _lambdaCompatibleRuntimes); err != nil {
			log.Errorf("invalid --compatible-runtimes: %s", err.Error())
			return
		}
	}
	if len(_lambdaDescription) > 0 {
		input.Description = aws.String(_lambdaDescription)
	}
	if len(_lambdaLicenseInfo) > 0 {
		input.LicenseInfo = aws.String(_lambdaLicenseInfo)
	}

	if resp, err := client.PublishLayerVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a [version] from the current code and configuration of a function. Use versions
// to create a snapshot of your function code and configuration that doesn't
// change.
//
// Lambda doesn't publish a version if the function's configuration and code
// haven't changed since the last version. Use UpdateFunctionCodeor UpdateFunctionConfiguration to update the function before
// publishing a version.
//
// Clients can invoke versions directly or with an alias. To create an alias, use CreateAlias.
//
// [version]: https://docs.aws.amazon.com/lambda/latest/dg/versioning-aliases.html
func lambda_PublishVersion(cfg aws.Config, client *lambda.Client) {
	input := &lambda.PublishVersionInput{
		// FunctionName: *string, // Required
	}

	if len(_lambdaFunctionName) > 0 {
		input.FunctionName = aws.String(_lambdaFunctionName)
	}
	if len(_lambdaCodeSha256) > 0 {
		input.CodeSha256 = aws.String(_lambdaCodeSha256)
	}
	if len(_lambdaDescription) > 0 {
		input.Description = aws.String(_lambdaDescription)
	}
	if len(_lambdaPublishTo) > 0 {
		if err := assignInputField(input, "PublishTo", _lambdaPublishTo); err != nil {
			log.Errorf("invalid --publish-to: %s", err.Error())
			return
		}
	}
	if len(_lambdaRevisionId) > 0 {
		input.RevisionId = aws.String(_lambdaRevisionId)
	}

	if resp, err := client.PublishVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update the code signing configuration for the function. Changes to the code
// signing configuration take effect the next time a user tries to deploy a code
// package to the function.
func lambda_PutFunctionCodeSigningConfig(cfg aws.Config, client *lambda.Client) {
	input := &lambda.PutFunctionCodeSigningConfigInput{
		// CodeSigningConfigArn: *string, // Required
		// FunctionName: *string, // Required
	}

	if len(_lambdaCodeSigningConfigArn) > 0 {
		input.CodeSigningConfigArn = aws.String(_lambdaCodeSigningConfigArn)
	}
	if len(_lambdaFunctionName) > 0 {
		input.FunctionName = aws.String(_lambdaFunctionName)
	}

	if resp, err := client.PutFunctionCodeSigningConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets the maximum number of simultaneous executions for a function, and reserves
// capacity for that concurrency level.
//
// Concurrency settings apply to the function as a whole, including all published
// versions and the unpublished version. Reserving concurrency both ensures that
// your function has capacity to process the specified number of events
// simultaneously, and prevents it from scaling beyond that level. Use GetFunctionto see the
// current setting for a function.
//
// Use GetAccountSettings to see your Regional concurrency limit. You can reserve concurrency for as
// many functions as you like, as long as you leave at least 100 simultaneous
// executions unreserved for functions that aren't configured with a per-function
// limit. For more information, see [Lambda function scaling].
//
// [Lambda function scaling]: https://docs.aws.amazon.com/lambda/latest/dg/invocation-scaling.html
func lambda_PutFunctionConcurrency(cfg aws.Config, client *lambda.Client) {
	input := &lambda.PutFunctionConcurrencyInput{
		// FunctionName: *string, // Required
		// ReservedConcurrentExecutions: *int32, // Required
	}

	if len(_lambdaFunctionName) > 0 {
		input.FunctionName = aws.String(_lambdaFunctionName)
	}
	if len(_lambdaReservedConcurrentExecutions) > 0 {
		if err := assignInputField(input, "ReservedConcurrentExecutions", _lambdaReservedConcurrentExecutions); err != nil {
			log.Errorf("invalid --reserved-concurrent-executions: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutFunctionConcurrency(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Configures options for [asynchronous invocation] on a function, version, or alias. If a configuration
// already exists for a function, version, or alias, this operation overwrites it.
// If you exclude any settings, they are removed. To set one option without
// affecting existing settings for other options, use UpdateFunctionEventInvokeConfig.
//
// By default, Lambda retries an asynchronous invocation twice if the function
// returns an error. It retains events in a queue for up to six hours. When an
// event fails all processing attempts or stays in the asynchronous invocation
// queue for too long, Lambda discards it. To retain discarded events, configure a
// dead-letter queue with UpdateFunctionConfiguration.
//
// To send an invocation record to a queue, topic, S3 bucket, function, or event
// bus, specify a [destination]. You can configure separate destinations for successful
// invocations (on-success) and events that fail all processing attempts
// (on-failure). You can configure destinations in addition to or instead of a
// dead-letter queue.
//
// S3 buckets are supported only for on-failure destinations. To retain records of
// successful invocations, use another destination type.
//
// [destination]: https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html#invocation-async-destinations
// [asynchronous invocation]: https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html
func lambda_PutFunctionEventInvokeConfig(cfg aws.Config, client *lambda.Client) {
	input := &lambda.PutFunctionEventInvokeConfigInput{
		// FunctionName: *string, // Required
	}

	if len(_lambdaFunctionName) > 0 {
		input.FunctionName = aws.String(_lambdaFunctionName)
	}
	if len(_lambdaDestinationConfig) > 0 {
		if err := assignInputField(input, "DestinationConfig", _lambdaDestinationConfig); err != nil {
			log.Errorf("invalid --destination-config: %s", err.Error())
			return
		}
	}
	if len(_lambdaMaximumEventAgeInSeconds) > 0 {
		if err := assignInputField(input, "MaximumEventAgeInSeconds", _lambdaMaximumEventAgeInSeconds); err != nil {
			log.Errorf("invalid --maximum-event-age-in-seconds: %s", err.Error())
			return
		}
	}
	if len(_lambdaMaximumRetryAttempts) > 0 {
		if err := assignInputField(input, "MaximumRetryAttempts", _lambdaMaximumRetryAttempts); err != nil {
			log.Errorf("invalid --maximum-retry-attempts: %s", err.Error())
			return
		}
	}
	if len(_lambdaQualifier) > 0 {
		input.Qualifier = aws.String(_lambdaQualifier)
	}

	if resp, err := client.PutFunctionEventInvokeConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets your function's [recursive loop detection] configuration.
// When you configure a Lambda function to output to the same service or resource
// that invokes the function, it's possible to create an infinite recursive loop.
// For example, a Lambda function might write a message to an Amazon Simple Queue
// Service (Amazon SQS) queue, which then invokes the same function. This
// invocation causes the function to write another message to the queue, which in
// turn invokes the function again.
//
// Lambda can detect certain types of recursive loops shortly after they occur.
// When Lambda detects a recursive loop and your function's recursive loop
// detection configuration is set to Terminate , it stops your function being
// invoked and notifies you.
//
// [recursive loop detection]: https://docs.aws.amazon.com/lambda/latest/dg/invocation-recursion.html
func lambda_PutFunctionRecursionConfig(cfg aws.Config, client *lambda.Client) {
	input := &lambda.PutFunctionRecursionConfigInput{
		// FunctionName: *string, // Required
		// RecursiveLoop: types.RecursiveLoop, // Required
	}

	if len(_lambdaFunctionName) > 0 {
		input.FunctionName = aws.String(_lambdaFunctionName)
	}
	if len(_lambdaRecursiveLoop) > 0 {
		if err := assignInputField(input, "RecursiveLoop", _lambdaRecursiveLoop); err != nil {
			log.Errorf("invalid --recursive-loop: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutFunctionRecursionConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets the scaling configuration for a Lambda Managed Instances function. The
// scaling configuration defines the minimum and maximum number of execution
// environments that can be provisioned for the function, allowing you to control
// scaling behavior and resource allocation.
func lambda_PutFunctionScalingConfig(cfg aws.Config, client *lambda.Client) {
	input := &lambda.PutFunctionScalingConfigInput{
		// FunctionName: *string, // Required
		// Qualifier: *string, // Required
	}

	if len(_lambdaFunctionName) > 0 {
		input.FunctionName = aws.String(_lambdaFunctionName)
	}
	if len(_lambdaQualifier) > 0 {
		input.Qualifier = aws.String(_lambdaQualifier)
	}
	if len(_lambdaFunctionScalingConfig) > 0 {
		if err := assignInputField(input, "FunctionScalingConfig", _lambdaFunctionScalingConfig); err != nil {
			log.Errorf("invalid --function-scaling-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutFunctionScalingConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds a provisioned concurrency configuration to a function's alias or version.
func lambda_PutProvisionedConcurrencyConfig(cfg aws.Config, client *lambda.Client) {
	input := &lambda.PutProvisionedConcurrencyConfigInput{
		// FunctionName: *string, // Required
		// ProvisionedConcurrentExecutions: *int32, // Required
		// Qualifier: *string, // Required
	}

	if len(_lambdaFunctionName) > 0 {
		input.FunctionName = aws.String(_lambdaFunctionName)
	}
	if len(_lambdaProvisionedConcurrentExecutions) > 0 {
		if err := assignInputField(input, "ProvisionedConcurrentExecutions", _lambdaProvisionedConcurrentExecutions); err != nil {
			log.Errorf("invalid --provisioned-concurrent-executions: %s", err.Error())
			return
		}
	}
	if len(_lambdaQualifier) > 0 {
		input.Qualifier = aws.String(_lambdaQualifier)
	}

	if resp, err := client.PutProvisionedConcurrencyConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets the runtime management configuration for a function's version. For more
// information, see [Runtime updates].
//
// [Runtime updates]: https://docs.aws.amazon.com/lambda/latest/dg/runtimes-update.html
func lambda_PutRuntimeManagementConfig(cfg aws.Config, client *lambda.Client) {
	input := &lambda.PutRuntimeManagementConfigInput{
		// FunctionName: *string, // Required
		// UpdateRuntimeOn: types.UpdateRuntimeOn, // Required
	}

	if len(_lambdaFunctionName) > 0 {
		input.FunctionName = aws.String(_lambdaFunctionName)
	}
	if len(_lambdaUpdateRuntimeOn) > 0 {
		if err := assignInputField(input, "UpdateRuntimeOn", _lambdaUpdateRuntimeOn); err != nil {
			log.Errorf("invalid --update-runtime-on: %s", err.Error())
			return
		}
	}
	if len(_lambdaQualifier) > 0 {
		input.Qualifier = aws.String(_lambdaQualifier)
	}
	if len(_lambdaRuntimeVersionArn) > 0 {
		input.RuntimeVersionArn = aws.String(_lambdaRuntimeVersionArn)
	}

	if resp, err := client.PutRuntimeManagementConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a statement from the permissions policy for a version of an [Lambda layer]. For more
// information, see AddLayerVersionPermission.
//
// [Lambda layer]: https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html
func lambda_RemoveLayerVersionPermission(cfg aws.Config, client *lambda.Client) {
	input := &lambda.RemoveLayerVersionPermissionInput{
		// LayerName: *string, // Required
		// StatementId: *string, // Required
		// VersionNumber: *int64, // Required
	}

	if len(_lambdaLayerName) > 0 {
		input.LayerName = aws.String(_lambdaLayerName)
	}
	if len(_lambdaStatementId) > 0 {
		input.StatementId = aws.String(_lambdaStatementId)
	}
	if len(_lambdaVersionNumber) > 0 {
		if err := assignInputField(input, "VersionNumber", _lambdaVersionNumber); err != nil {
			log.Errorf("invalid --version-number: %s", err.Error())
			return
		}
	}
	if len(_lambdaRevisionId) > 0 {
		input.RevisionId = aws.String(_lambdaRevisionId)
	}

	if resp, err := client.RemoveLayerVersionPermission(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Revokes function-use permission from an Amazon Web Services service or another
// Amazon Web Services account. You can get the ID of the statement from the output
// of GetPolicy.
func lambda_RemovePermission(cfg aws.Config, client *lambda.Client) {
	input := &lambda.RemovePermissionInput{
		// FunctionName: *string, // Required
		// StatementId: *string, // Required
	}

	if len(_lambdaFunctionName) > 0 {
		input.FunctionName = aws.String(_lambdaFunctionName)
	}
	if len(_lambdaStatementId) > 0 {
		input.StatementId = aws.String(_lambdaStatementId)
	}
	if len(_lambdaQualifier) > 0 {
		input.Qualifier = aws.String(_lambdaQualifier)
	}
	if len(_lambdaRevisionId) > 0 {
		input.RevisionId = aws.String(_lambdaRevisionId)
	}

	if resp, err := client.RemovePermission(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sends a failure response for a callback operation in a durable execution. Use
// this API when an external system cannot complete a callback operation
// successfully.
func lambda_SendDurableExecutionCallbackFailure(cfg aws.Config, client *lambda.Client) {
	input := &lambda.SendDurableExecutionCallbackFailureInput{
		// CallbackId: *string, // Required
	}

	if len(_lambdaCallbackId) > 0 {
		input.CallbackId = aws.String(_lambdaCallbackId)
	}
	if len(_lambdaError) > 0 {
		if err := assignInputField(input, "Error", _lambdaError); err != nil {
			log.Errorf("invalid --error: %s", err.Error())
			return
		}
	}

	if resp, err := client.SendDurableExecutionCallbackFailure(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sends a heartbeat signal for a long-running callback operation to prevent
// timeout. Use this API to extend the callback timeout period while the external
// operation is still in progress.
func lambda_SendDurableExecutionCallbackHeartbeat(cfg aws.Config, client *lambda.Client) {
	input := &lambda.SendDurableExecutionCallbackHeartbeatInput{
		// CallbackId: *string, // Required
	}

	if len(_lambdaCallbackId) > 0 {
		input.CallbackId = aws.String(_lambdaCallbackId)
	}

	if resp, err := client.SendDurableExecutionCallbackHeartbeat(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sends a successful completion response for a callback operation in a durable
// execution. Use this API when an external system has successfully completed a
// callback operation.
func lambda_SendDurableExecutionCallbackSuccess(cfg aws.Config, client *lambda.Client) {
	input := &lambda.SendDurableExecutionCallbackSuccessInput{
		// CallbackId: *string, // Required
	}

	if len(_lambdaCallbackId) > 0 {
		input.CallbackId = aws.String(_lambdaCallbackId)
	}
	if len(_lambdaResult) > 0 {
		if err := assignInputField(input, "Result", _lambdaResult); err != nil {
			log.Errorf("invalid --result: %s", err.Error())
			return
		}
	}

	if resp, err := client.SendDurableExecutionCallbackSuccess(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a running [durable execution]. The execution transitions to STOPPED status and cannot be
// resumed. Any in-progress operations are terminated.
//
// [durable execution]: https://docs.aws.amazon.com/lambda/latest/dg/durable-functions.html
func lambda_StopDurableExecution(cfg aws.Config, client *lambda.Client) {
	input := &lambda.StopDurableExecutionInput{
		// DurableExecutionArn: *string, // Required
	}

	if len(_lambdaDurableExecutionArn) > 0 {
		input.DurableExecutionArn = aws.String(_lambdaDurableExecutionArn)
	}
	if len(_lambdaError) > 0 {
		if err := assignInputField(input, "Error", _lambdaError); err != nil {
			log.Errorf("invalid --error: %s", err.Error())
			return
		}
	}

	if resp, err := client.StopDurableExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds [tags] to a function, event source mapping, or code signing configuration.
//
// [tags]: https://docs.aws.amazon.com/lambda/latest/dg/tagging.html
func lambda_TagResource(cfg aws.Config, client *lambda.Client) {
	input := &lambda.TagResourceInput{
		// Resource: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_lambdaResource) > 0 {
		input.Resource = aws.String(_lambdaResource)
	}
	if len(_lambdaTags) > 0 {
		if err := assignInputField(input, "Tags", _lambdaTags); err != nil {
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

// Removes [tags] from a function, event source mapping, or code signing configuration.
//
// [tags]: https://docs.aws.amazon.com/lambda/latest/dg/tagging.html
func lambda_UntagResource(cfg aws.Config, client *lambda.Client) {
	input := &lambda.UntagResourceInput{
		// Resource: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_lambdaResource) > 0 {
		input.Resource = aws.String(_lambdaResource)
	}
	if len(_lambdaTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _lambdaTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration of a Lambda function [alias].
//
// [alias]: https://docs.aws.amazon.com/lambda/latest/dg/configuration-aliases.html
func lambda_UpdateAlias(cfg aws.Config, client *lambda.Client) {
	input := &lambda.UpdateAliasInput{
		// FunctionName: *string, // Required
		// Name: *string, // Required
	}

	if len(_lambdaFunctionName) > 0 {
		input.FunctionName = aws.String(_lambdaFunctionName)
	}
	if len(_lambdaName) > 0 {
		input.Name = aws.String(_lambdaName)
	}
	if len(_lambdaDescription) > 0 {
		input.Description = aws.String(_lambdaDescription)
	}
	if len(_lambdaFunctionVersion) > 0 {
		input.FunctionVersion = aws.String(_lambdaFunctionVersion)
	}
	if len(_lambdaRevisionId) > 0 {
		input.RevisionId = aws.String(_lambdaRevisionId)
	}
	if len(_lambdaRoutingConfig) > 0 {
		if err := assignInputField(input, "RoutingConfig", _lambdaRoutingConfig); err != nil {
			log.Errorf("invalid --routing-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateAlias(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration of an existing capacity provider.
func lambda_UpdateCapacityProvider(cfg aws.Config, client *lambda.Client) {
	input := &lambda.UpdateCapacityProviderInput{
		// CapacityProviderName: *string, // Required
	}

	if len(_lambdaCapacityProviderName) > 0 {
		input.CapacityProviderName = aws.String(_lambdaCapacityProviderName)
	}
	if len(_lambdaCapacityProviderScalingConfig) > 0 {
		if err := assignInputField(input, "CapacityProviderScalingConfig", _lambdaCapacityProviderScalingConfig); err != nil {
			log.Errorf("invalid --capacity-provider-scaling-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateCapacityProvider(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update the code signing configuration. Changes to the code signing
// configuration take effect the next time a user tries to deploy a code package to
// the function.
func lambda_UpdateCodeSigningConfig(cfg aws.Config, client *lambda.Client) {
	input := &lambda.UpdateCodeSigningConfigInput{
		// CodeSigningConfigArn: *string, // Required
	}

	if len(_lambdaCodeSigningConfigArn) > 0 {
		input.CodeSigningConfigArn = aws.String(_lambdaCodeSigningConfigArn)
	}
	if len(_lambdaAllowedPublishers) > 0 {
		if err := assignInputField(input, "AllowedPublishers", _lambdaAllowedPublishers); err != nil {
			log.Errorf("invalid --allowed-publishers: %s", err.Error())
			return
		}
	}
	if len(_lambdaCodeSigningPolicies) > 0 {
		if err := assignInputField(input, "CodeSigningPolicies", _lambdaCodeSigningPolicies); err != nil {
			log.Errorf("invalid --code-signing-policies: %s", err.Error())
			return
		}
	}
	if len(_lambdaDescription) > 0 {
		input.Description = aws.String(_lambdaDescription)
	}

	if resp, err := client.UpdateCodeSigningConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an event source mapping. You can change the function that Lambda
// invokes, or pause invocation and resume later from the same location.
//
// For details about how to configure different event sources, see the following
// topics.
//
// [Amazon DynamoDB Streams]
//
// [Amazon Kinesis]
//
// [Amazon SQS]
//
// [Amazon MQ and RabbitMQ]
//
// [Amazon MSK]
//
// [Apache Kafka]
//
// [Amazon DocumentDB]
//
// The following error handling options are available for stream sources
// (DynamoDB, Kinesis, Amazon MSK, and self-managed Apache Kafka):
//
// - BisectBatchOnFunctionError – If the function returns an error, split the
// batch in two and retry.
//
// - MaximumRecordAgeInSeconds – Discard records older than the specified age.
// The default value is infinite (-1). When set to infinite (-1), failed records
// are retried until the record expires
//
// - MaximumRetryAttempts – Discard records after the specified number of
// retries. The default value is infinite (-1). When set to infinite (-1), failed
// records are retried until the record expires.
//
// - OnFailure – Send discarded records to an Amazon SQS queue, Amazon SNS topic,
// Kafka topic, or Amazon S3 bucket. For more information, see [Adding a destination].
//
// The following option is available only for DynamoDB and Kinesis event sources:
//
// - ParallelizationFactor – Process multiple batches from each shard
// concurrently.
//
// For information about which configuration parameters apply to each event
// source, see the following topics.
//
// [Amazon DynamoDB Streams]
//
// [Amazon Kinesis]
//
// [Amazon SQS]
//
// [Amazon MQ and RabbitMQ]
//
// [Amazon MSK]
//
// [Apache Kafka]
//
// [Amazon DocumentDB]
//
// [Amazon DynamoDB Streams]: https://docs.aws.amazon.com/lambda/latest/dg/with-ddb.html#services-ddb-params
// [Amazon SQS]: https://docs.aws.amazon.com/lambda/latest/dg/with-sqs.html#services-sqs-params
// [Amazon MSK]: https://docs.aws.amazon.com/lambda/latest/dg/with-msk.html#services-msk-parms
// [Amazon Kinesis]: https://docs.aws.amazon.com/lambda/latest/dg/with-kinesis.html#services-kinesis-params
// [Amazon MQ and RabbitMQ]: https://docs.aws.amazon.com/lambda/latest/dg/with-mq.html#services-mq-params
// [Apache Kafka]: https://docs.aws.amazon.com/lambda/latest/dg/with-kafka.html#services-kafka-parms
// [Amazon DocumentDB]: https://docs.aws.amazon.com/lambda/latest/dg/with-documentdb.html#docdb-configuration
// [Adding a destination]: https://docs.aws.amazon.com/lambda/latest/dg/invocation-async-retain-records.html#invocation-async-destinations
func lambda_UpdateEventSourceMapping(cfg aws.Config, client *lambda.Client) {
	input := &lambda.UpdateEventSourceMappingInput{
		// UUID: *string, // Required
	}

	if len(_lambdaUUID) > 0 {
		input.UUID = aws.String(_lambdaUUID)
	}
	if len(_lambdaAmazonManagedKafkaEventSourceConfig) > 0 {
		if err := assignInputField(input, "AmazonManagedKafkaEventSourceConfig", _lambdaAmazonManagedKafkaEventSourceConfig); err != nil {
			log.Errorf("invalid --amazon-managed-kafka-event-source-config: %s", err.Error())
			return
		}
	}
	if len(_lambdaBatchSize) > 0 {
		if err := assignInputField(input, "BatchSize", _lambdaBatchSize); err != nil {
			log.Errorf("invalid --batch-size: %s", err.Error())
			return
		}
	}
	if len(_lambdaBisectBatchOnFunctionError) > 0 {
		if err := assignInputField(input, "BisectBatchOnFunctionError", _lambdaBisectBatchOnFunctionError); err != nil {
			log.Errorf("invalid --bisect-batch-on-function-error: %s", err.Error())
			return
		}
	}
	if len(_lambdaDestinationConfig) > 0 {
		if err := assignInputField(input, "DestinationConfig", _lambdaDestinationConfig); err != nil {
			log.Errorf("invalid --destination-config: %s", err.Error())
			return
		}
	}
	if len(_lambdaDocumentDBEventSourceConfig) > 0 {
		if err := assignInputField(input, "DocumentDBEventSourceConfig", _lambdaDocumentDBEventSourceConfig); err != nil {
			log.Errorf("invalid --document-db-event-source-config: %s", err.Error())
			return
		}
	}
	if len(_lambdaEnabled) > 0 {
		if err := assignInputField(input, "Enabled", _lambdaEnabled); err != nil {
			log.Errorf("invalid --enabled: %s", err.Error())
			return
		}
	}
	if len(_lambdaFilterCriteria) > 0 {
		if err := assignInputField(input, "FilterCriteria", _lambdaFilterCriteria); err != nil {
			log.Errorf("invalid --filter-criteria: %s", err.Error())
			return
		}
	}
	if len(_lambdaFunctionName) > 0 {
		input.FunctionName = aws.String(_lambdaFunctionName)
	}
	if len(_lambdaFunctionResponseTypes) > 0 {
		if err := assignInputField(input, "FunctionResponseTypes", _lambdaFunctionResponseTypes); err != nil {
			log.Errorf("invalid --function-response-types: %s", err.Error())
			return
		}
	}
	if len(_lambdaKMSKeyArn) > 0 {
		input.KMSKeyArn = aws.String(_lambdaKMSKeyArn)
	}
	if len(_lambdaLoggingConfig) > 0 {
		if err := assignInputField(input, "LoggingConfig", _lambdaLoggingConfig); err != nil {
			log.Errorf("invalid --logging-config: %s", err.Error())
			return
		}
	}
	if len(_lambdaMaximumBatchingWindowInSeconds) > 0 {
		if err := assignInputField(input, "MaximumBatchingWindowInSeconds", _lambdaMaximumBatchingWindowInSeconds); err != nil {
			log.Errorf("invalid --maximum-batching-window-in-seconds: %s", err.Error())
			return
		}
	}
	if len(_lambdaMaximumRecordAgeInSeconds) > 0 {
		if err := assignInputField(input, "MaximumRecordAgeInSeconds", _lambdaMaximumRecordAgeInSeconds); err != nil {
			log.Errorf("invalid --maximum-record-age-in-seconds: %s", err.Error())
			return
		}
	}
	if len(_lambdaMaximumRetryAttempts) > 0 {
		if err := assignInputField(input, "MaximumRetryAttempts", _lambdaMaximumRetryAttempts); err != nil {
			log.Errorf("invalid --maximum-retry-attempts: %s", err.Error())
			return
		}
	}
	if len(_lambdaMetricsConfig) > 0 {
		if err := assignInputField(input, "MetricsConfig", _lambdaMetricsConfig); err != nil {
			log.Errorf("invalid --metrics-config: %s", err.Error())
			return
		}
	}
	if len(_lambdaParallelizationFactor) > 0 {
		if err := assignInputField(input, "ParallelizationFactor", _lambdaParallelizationFactor); err != nil {
			log.Errorf("invalid --parallelization-factor: %s", err.Error())
			return
		}
	}
	if len(_lambdaProvisionedPollerConfig) > 0 {
		if err := assignInputField(input, "ProvisionedPollerConfig", _lambdaProvisionedPollerConfig); err != nil {
			log.Errorf("invalid --provisioned-poller-config: %s", err.Error())
			return
		}
	}
	if len(_lambdaScalingConfig) > 0 {
		if err := assignInputField(input, "ScalingConfig", _lambdaScalingConfig); err != nil {
			log.Errorf("invalid --scaling-config: %s", err.Error())
			return
		}
	}
	if len(_lambdaSelfManagedKafkaEventSourceConfig) > 0 {
		if err := assignInputField(input, "SelfManagedKafkaEventSourceConfig", _lambdaSelfManagedKafkaEventSourceConfig); err != nil {
			log.Errorf("invalid --self-managed-kafka-event-source-config: %s", err.Error())
			return
		}
	}
	if len(_lambdaSourceAccessConfigurations) > 0 {
		if err := assignInputField(input, "SourceAccessConfigurations", _lambdaSourceAccessConfigurations); err != nil {
			log.Errorf("invalid --source-access-configurations: %s", err.Error())
			return
		}
	}
	if len(_lambdaTumblingWindowInSeconds) > 0 {
		if err := assignInputField(input, "TumblingWindowInSeconds", _lambdaTumblingWindowInSeconds); err != nil {
			log.Errorf("invalid --tumbling-window-in-seconds: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateEventSourceMapping(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a Lambda function's code. If code signing is enabled for the function,
// the code package must be signed by a trusted publisher. For more information,
// see [Configuring code signing for Lambda].
//
// If the function's package type is Image , then you must specify the code package
// in ImageUri as the URI of a [container image] in the Amazon ECR registry.
//
// If the function's package type is Zip , then you must specify the deployment
// package as a [.zip file archive]. Enter the Amazon S3 bucket and key of the code .zip file
// location. You can also provide the function code inline using the ZipFile field.
//
// The code in the deployment package must be compatible with the target
// instruction set architecture of the function ( x86-64 or arm64 ).
//
// The function's code is locked when you publish a version. You can't modify the
// code of a published version, only the unpublished version.
//
// For a function defined as a container image, Lambda resolves the image tag to
// an image digest. In Amazon ECR, if you update the image tag to a new image,
// Lambda does not automatically update the function.
//
// [.zip file archive]: https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-package.html#gettingstarted-package-zip
// [Configuring code signing for Lambda]: https://docs.aws.amazon.com/lambda/latest/dg/configuration-codesigning.html
// [container image]: https://docs.aws.amazon.com/lambda/latest/dg/lambda-images.html
func lambda_UpdateFunctionCode(cfg aws.Config, client *lambda.Client) {
	input := &lambda.UpdateFunctionCodeInput{
		// FunctionName: *string, // Required
	}

	if len(_lambdaFunctionName) > 0 {
		input.FunctionName = aws.String(_lambdaFunctionName)
	}
	if len(_lambdaArchitectures) > 0 {
		if err := assignInputField(input, "Architectures", _lambdaArchitectures); err != nil {
			log.Errorf("invalid --architectures: %s", err.Error())
			return
		}
	}
	if len(_lambdaDryRun) > 0 {
		if err := assignInputField(input, "DryRun", _lambdaDryRun); err != nil {
			log.Errorf("invalid --dry-run: %s", err.Error())
			return
		}
	}
	if len(_lambdaImageUri) > 0 {
		input.ImageUri = aws.String(_lambdaImageUri)
	}
	if len(_lambdaPublish) > 0 {
		if err := assignInputField(input, "Publish", _lambdaPublish); err != nil {
			log.Errorf("invalid --publish: %s", err.Error())
			return
		}
	}
	if len(_lambdaPublishTo) > 0 {
		if err := assignInputField(input, "PublishTo", _lambdaPublishTo); err != nil {
			log.Errorf("invalid --publish-to: %s", err.Error())
			return
		}
	}
	if len(_lambdaRevisionId) > 0 {
		input.RevisionId = aws.String(_lambdaRevisionId)
	}
	if len(_lambdaS3Bucket) > 0 {
		input.S3Bucket = aws.String(_lambdaS3Bucket)
	}
	if len(_lambdaS3Key) > 0 {
		input.S3Key = aws.String(_lambdaS3Key)
	}
	if len(_lambdaS3ObjectVersion) > 0 {
		input.S3ObjectVersion = aws.String(_lambdaS3ObjectVersion)
	}
	if len(_lambdaSourceKMSKeyArn) > 0 {
		input.SourceKMSKeyArn = aws.String(_lambdaSourceKMSKeyArn)
	}
	if len(_lambdaZipFile) > 0 {
		if err := assignInputField(input, "ZipFile", _lambdaZipFile); err != nil {
			log.Errorf("invalid --zip-file: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateFunctionCode(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modify the version-specific settings of a Lambda function.
// When you update a function, Lambda provisions an instance of the function and
// its supporting resources. If your function connects to a VPC, this process can
// take a minute. During this time, you can't modify the function, but you can
// still invoke it. The LastUpdateStatus , LastUpdateStatusReason , and
// LastUpdateStatusReasonCode fields in the response from GetFunctionConfiguration indicate when the
// update is complete and the function is processing events with the new
// configuration. For more information, see [Lambda function states].
//
// These settings can vary between versions of a function and are locked when you
// publish a version. You can't modify the configuration of a published version,
// only the unpublished version.
//
// To configure function concurrency, use PutFunctionConcurrency. To grant invoke permissions to an
// Amazon Web Services account or Amazon Web Services service, use AddPermission.
//
// [Lambda function states]: https://docs.aws.amazon.com/lambda/latest/dg/functions-states.html
func lambda_UpdateFunctionConfiguration(cfg aws.Config, client *lambda.Client) {
	input := &lambda.UpdateFunctionConfigurationInput{
		// FunctionName: *string, // Required
	}

	if len(_lambdaFunctionName) > 0 {
		input.FunctionName = aws.String(_lambdaFunctionName)
	}
	if len(_lambdaCapacityProviderConfig) > 0 {
		if err := assignInputField(input, "CapacityProviderConfig", _lambdaCapacityProviderConfig); err != nil {
			log.Errorf("invalid --capacity-provider-config: %s", err.Error())
			return
		}
	}
	if len(_lambdaDeadLetterConfig) > 0 {
		if err := assignInputField(input, "DeadLetterConfig", _lambdaDeadLetterConfig); err != nil {
			log.Errorf("invalid --dead-letter-config: %s", err.Error())
			return
		}
	}
	if len(_lambdaDescription) > 0 {
		input.Description = aws.String(_lambdaDescription)
	}
	if len(_lambdaDurableConfig) > 0 {
		if err := assignInputField(input, "DurableConfig", _lambdaDurableConfig); err != nil {
			log.Errorf("invalid --durable-config: %s", err.Error())
			return
		}
	}
	if len(_lambdaEnvironment) > 0 {
		if err := assignInputField(input, "Environment", _lambdaEnvironment); err != nil {
			log.Errorf("invalid --environment: %s", err.Error())
			return
		}
	}
	if len(_lambdaEphemeralStorage) > 0 {
		if err := assignInputField(input, "EphemeralStorage", _lambdaEphemeralStorage); err != nil {
			log.Errorf("invalid --ephemeral-storage: %s", err.Error())
			return
		}
	}
	if len(_lambdaFileSystemConfigs) > 0 {
		if err := assignInputField(input, "FileSystemConfigs", _lambdaFileSystemConfigs); err != nil {
			log.Errorf("invalid --file-system-configs: %s", err.Error())
			return
		}
	}
	if len(_lambdaHandler) > 0 {
		input.Handler = aws.String(_lambdaHandler)
	}
	if len(_lambdaImageConfig) > 0 {
		if err := assignInputField(input, "ImageConfig", _lambdaImageConfig); err != nil {
			log.Errorf("invalid --image-config: %s", err.Error())
			return
		}
	}
	if len(_lambdaKMSKeyArn) > 0 {
		input.KMSKeyArn = aws.String(_lambdaKMSKeyArn)
	}
	if len(_lambdaLayers) > 0 {
		input.Layers = append([]string(nil), _lambdaLayers...)
	}
	if len(_lambdaLoggingConfig) > 0 {
		if err := assignInputField(input, "LoggingConfig", _lambdaLoggingConfig); err != nil {
			log.Errorf("invalid --logging-config: %s", err.Error())
			return
		}
	}
	if len(_lambdaMemorySize) > 0 {
		if err := assignInputField(input, "MemorySize", _lambdaMemorySize); err != nil {
			log.Errorf("invalid --memory-size: %s", err.Error())
			return
		}
	}
	if len(_lambdaRevisionId) > 0 {
		input.RevisionId = aws.String(_lambdaRevisionId)
	}
	if len(_lambdaRole) > 0 {
		input.Role = aws.String(_lambdaRole)
	}
	if len(_lambdaRuntime) > 0 {
		if err := assignInputField(input, "Runtime", _lambdaRuntime); err != nil {
			log.Errorf("invalid --runtime: %s", err.Error())
			return
		}
	}
	if len(_lambdaSnapStart) > 0 {
		if err := assignInputField(input, "SnapStart", _lambdaSnapStart); err != nil {
			log.Errorf("invalid --snap-start: %s", err.Error())
			return
		}
	}
	if len(_lambdaTimeout) > 0 {
		if err := assignInputField(input, "Timeout", _lambdaTimeout); err != nil {
			log.Errorf("invalid --timeout: %s", err.Error())
			return
		}
	}
	if len(_lambdaTracingConfig) > 0 {
		if err := assignInputField(input, "TracingConfig", _lambdaTracingConfig); err != nil {
			log.Errorf("invalid --tracing-config: %s", err.Error())
			return
		}
	}
	if len(_lambdaVpcConfig) > 0 {
		if err := assignInputField(input, "VpcConfig", _lambdaVpcConfig); err != nil {
			log.Errorf("invalid --vpc-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateFunctionConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration for asynchronous invocation for a function, version,
// or alias.
//
// To configure options for asynchronous invocation, use PutFunctionEventInvokeConfig.
func lambda_UpdateFunctionEventInvokeConfig(cfg aws.Config, client *lambda.Client) {
	input := &lambda.UpdateFunctionEventInvokeConfigInput{
		// FunctionName: *string, // Required
	}

	if len(_lambdaFunctionName) > 0 {
		input.FunctionName = aws.String(_lambdaFunctionName)
	}
	if len(_lambdaDestinationConfig) > 0 {
		if err := assignInputField(input, "DestinationConfig", _lambdaDestinationConfig); err != nil {
			log.Errorf("invalid --destination-config: %s", err.Error())
			return
		}
	}
	if len(_lambdaMaximumEventAgeInSeconds) > 0 {
		if err := assignInputField(input, "MaximumEventAgeInSeconds", _lambdaMaximumEventAgeInSeconds); err != nil {
			log.Errorf("invalid --maximum-event-age-in-seconds: %s", err.Error())
			return
		}
	}
	if len(_lambdaMaximumRetryAttempts) > 0 {
		if err := assignInputField(input, "MaximumRetryAttempts", _lambdaMaximumRetryAttempts); err != nil {
			log.Errorf("invalid --maximum-retry-attempts: %s", err.Error())
			return
		}
	}
	if len(_lambdaQualifier) > 0 {
		input.Qualifier = aws.String(_lambdaQualifier)
	}

	if resp, err := client.UpdateFunctionEventInvokeConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration for a Lambda function URL.
func lambda_UpdateFunctionUrlConfig(cfg aws.Config, client *lambda.Client) {
	input := &lambda.UpdateFunctionUrlConfigInput{
		// FunctionName: *string, // Required
	}

	if len(_lambdaFunctionName) > 0 {
		input.FunctionName = aws.String(_lambdaFunctionName)
	}
	if len(_lambdaAuthType) > 0 {
		if err := assignInputField(input, "AuthType", _lambdaAuthType); err != nil {
			log.Errorf("invalid --auth-type: %s", err.Error())
			return
		}
	}
	if len(_lambdaCors) > 0 {
		if err := assignInputField(input, "Cors", _lambdaCors); err != nil {
			log.Errorf("invalid --cors: %s", err.Error())
			return
		}
	}
	if len(_lambdaInvokeMode) > 0 {
		if err := assignInputField(input, "InvokeMode", _lambdaInvokeMode); err != nil {
			log.Errorf("invalid --invoke-mode: %s", err.Error())
			return
		}
	}
	if len(_lambdaQualifier) > 0 {
		input.Qualifier = aws.String(_lambdaQualifier)
	}

	if resp, err := client.UpdateFunctionUrlConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_lambdaCmd)
	_lambdaCmd.Flags().SortFlags = false

	_lambdaCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_lambdaCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_lambdaCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_lambdaCmd.Flags().StringVarP(&_lambdaAction, "action", "", "", "Action")
	_lambdaCmd.Flags().StringVarP(&_lambdaAllowedPublishers, "allowed-publishers", "", "", "Allowed Publishers")
	_lambdaCmd.Flags().StringVarP(&_lambdaAmazonManagedKafkaEventSourceConfig, "amazon-managed-kafka-event-source-config", "", "", "Amazon Managed Kafka Event Source Config")
	_lambdaCmd.Flags().StringVarP(&_lambdaArchitectures, "architectures", "", "", "Architectures")
	_lambdaCmd.Flags().StringVarP(&_lambdaArn, "arn", "", "", "ARN")
	_lambdaCmd.Flags().StringVarP(&_lambdaAuthType, "auth-type", "", "", "Auth Type")
	_lambdaCmd.Flags().StringVarP(&_lambdaBatchSize, "batch-size", "", "", "Batch Size")
	_lambdaCmd.Flags().StringVarP(&_lambdaBisectBatchOnFunctionError, "bisect-batch-on-function-error", "", "", "Bisect Batch On Function Error")
	_lambdaCmd.Flags().StringVarP(&_lambdaCallbackId, "callback-id", "", "", "Callback ID")
	_lambdaCmd.Flags().StringVarP(&_lambdaCapacityProviderConfig, "capacity-provider-config", "", "", "Capacity Provider Config")
	_lambdaCmd.Flags().StringVarP(&_lambdaCapacityProviderName, "capacity-provider-name", "", "", "Capacity Provider Name")
	_lambdaCmd.Flags().StringVarP(&_lambdaCapacityProviderScalingConfig, "capacity-provider-scaling-config", "", "", "Capacity Provider Scaling Config")
	_lambdaCmd.Flags().StringVarP(&_lambdaCheckpointToken, "checkpoint-token", "", "", "Checkpoint Token")
	_lambdaCmd.Flags().StringVarP(&_lambdaClientContext, "client-context", "", "", "Client Context")
	_lambdaCmd.Flags().StringVarP(&_lambdaClientToken, "client-token", "", "", "Client Token")
	_lambdaCmd.Flags().StringVarP(&_lambdaCode, "code", "", "", "Code")
	_lambdaCmd.Flags().StringVarP(&_lambdaCodeSha256, "code-sha256", "", "", "Code SHA256")
	_lambdaCmd.Flags().StringVarP(&_lambdaCodeSigningConfigArn, "code-signing-config-arn", "", "", "Code Signing Config ARN")
	_lambdaCmd.Flags().StringVarP(&_lambdaCodeSigningPolicies, "code-signing-policies", "", "", "Code Signing Policies")
	_lambdaCmd.Flags().StringVarP(&_lambdaCompatibleArchitecture, "compatible-architecture", "", "", "Compatible Architecture")
	_lambdaCmd.Flags().StringVarP(&_lambdaCompatibleArchitectures, "compatible-architectures", "", "", "Compatible Architectures")
	_lambdaCmd.Flags().StringVarP(&_lambdaCompatibleRuntime, "compatible-runtime", "", "", "Compatible Runtime")
	_lambdaCmd.Flags().StringVarP(&_lambdaCompatibleRuntimes, "compatible-runtimes", "", "", "Compatible Runtimes")
	_lambdaCmd.Flags().StringVarP(&_lambdaContent, "content", "", "", "Content")
	_lambdaCmd.Flags().StringVarP(&_lambdaCors, "cors", "", "", "Cors")
	_lambdaCmd.Flags().StringVarP(&_lambdaDeadLetterConfig, "dead-letter-config", "", "", "Dead Letter Config")
	_lambdaCmd.Flags().StringVarP(&_lambdaDescription, "description", "", "", "Description")
	_lambdaCmd.Flags().StringVarP(&_lambdaDestinationConfig, "destination-config", "", "", "Destination Config")
	_lambdaCmd.Flags().StringVarP(&_lambdaDocumentDBEventSourceConfig, "document-db-event-source-config", "", "", "Document DB Event Source Config")
	_lambdaCmd.Flags().StringVarP(&_lambdaDryRun, "dry-run", "", "", "Dry Run")
	_lambdaCmd.Flags().StringVarP(&_lambdaDurableConfig, "durable-config", "", "", "Durable Config")
	_lambdaCmd.Flags().StringVarP(&_lambdaDurableExecutionArn, "durable-execution-arn", "", "", "Durable Execution ARN")
	_lambdaCmd.Flags().StringVarP(&_lambdaDurableExecutionName, "durable-execution-name", "", "", "Durable Execution Name")
	_lambdaCmd.Flags().StringVarP(&_lambdaEnabled, "enabled", "", "", "Enabled")
	_lambdaCmd.Flags().StringVarP(&_lambdaEnvironment, "environment", "", "", "Environment")
	_lambdaCmd.Flags().StringVarP(&_lambdaEphemeralStorage, "ephemeral-storage", "", "", "Ephemeral Storage")
	_lambdaCmd.Flags().StringVarP(&_lambdaError, "error", "", "", "Error")
	_lambdaCmd.Flags().StringVarP(&_lambdaEventSourceArn, "event-source-arn", "", "", "Event Source ARN")
	_lambdaCmd.Flags().StringVarP(&_lambdaEventSourceToken, "event-source-token", "", "", "Event Source Token")
	_lambdaCmd.Flags().StringVarP(&_lambdaFileSystemConfigs, "file-system-configs", "", "", "File System Configs")
	_lambdaCmd.Flags().StringVarP(&_lambdaFilterCriteria, "filter-criteria", "", "", "Filter Criteria")
	_lambdaCmd.Flags().StringVarP(&_lambdaFunctionName, "function-name", "", "", "Function Name")
	_lambdaCmd.Flags().StringVarP(&_lambdaFunctionResponseTypes, "function-response-types", "", "", "Function Response Types")
	_lambdaCmd.Flags().StringVarP(&_lambdaFunctionScalingConfig, "function-scaling-config", "", "", "Function Scaling Config")
	_lambdaCmd.Flags().StringVarP(&_lambdaFunctionUrlAuthType, "function-url-auth-type", "", "", "Function URL Auth Type")
	_lambdaCmd.Flags().StringVarP(&_lambdaFunctionVersion, "function-version", "", "", "Function Version")
	_lambdaCmd.Flags().StringVarP(&_lambdaHandler, "handler", "", "", "Handler")
	_lambdaCmd.Flags().StringVarP(&_lambdaImageConfig, "image-config", "", "", "Image Config")
	_lambdaCmd.Flags().StringVarP(&_lambdaImageUri, "image-uri", "", "", "Image URI")
	_lambdaCmd.Flags().StringVarP(&_lambdaIncludeExecutionData, "include-execution-data", "", "", "Include Execution Data")
	_lambdaCmd.Flags().StringVarP(&_lambdaInstanceRequirements, "instance-requirements", "", "", "Instance Requirements")
	_lambdaCmd.Flags().StringVarP(&_lambdaInvocationType, "invocation-type", "", "", "Invocation Type")
	_lambdaCmd.Flags().StringVarP(&_lambdaInvokeArgs, "invoke-args", "", "", "Invoke Args")
	_lambdaCmd.Flags().StringVarP(&_lambdaInvokeMode, "invoke-mode", "", "", "Invoke Mode")
	_lambdaCmd.Flags().StringVarP(&_lambdaInvokedViaFunctionUrl, "invoked-via-function-url", "", "", "Invoked Via Function URL")
	_lambdaCmd.Flags().StringVarP(&_lambdaKMSKeyArn, "kms-key-arn", "", "", "KMS Key ARN")
	_lambdaCmd.Flags().StringVarP(&_lambdaLayerName, "layer-name", "", "", "Layer Name")
	_lambdaCmd.Flags().StringSliceVarP(&_lambdaLayers, "layers", "", nil, "Layers")
	_lambdaCmd.Flags().StringVarP(&_lambdaLicenseInfo, "license-info", "", "", "License Info")
	_lambdaCmd.Flags().StringVarP(&_lambdaLogType, "log-type", "", "", "Log Type")
	_lambdaCmd.Flags().StringVarP(&_lambdaLoggingConfig, "logging-config", "", "", "Logging Config")
	_lambdaCmd.Flags().StringVarP(&_lambdaMarker, "marker", "", "", "Marker")
	_lambdaCmd.Flags().StringVarP(&_lambdaMasterRegion, "master-region", "", "", "Master Region")
	_lambdaCmd.Flags().StringVarP(&_lambdaMaxItems, "max-items", "", "", "Max Items")
	_lambdaCmd.Flags().StringVarP(&_lambdaMaximumBatchingWindowInSeconds, "maximum-batching-window-in-seconds", "", "", "Maximum Batching Window In Seconds")
	_lambdaCmd.Flags().StringVarP(&_lambdaMaximumEventAgeInSeconds, "maximum-event-age-in-seconds", "", "", "Maximum Event Age In Seconds")
	_lambdaCmd.Flags().StringVarP(&_lambdaMaximumRecordAgeInSeconds, "maximum-record-age-in-seconds", "", "", "Maximum Record Age In Seconds")
	_lambdaCmd.Flags().StringVarP(&_lambdaMaximumRetryAttempts, "maximum-retry-attempts", "", "", "Maximum Retry Attempts")
	_lambdaCmd.Flags().StringVarP(&_lambdaMemorySize, "memory-size", "", "", "Memory Size")
	_lambdaCmd.Flags().StringVarP(&_lambdaMetricsConfig, "metrics-config", "", "", "Metrics Config")
	_lambdaCmd.Flags().StringVarP(&_lambdaName, "name", "", "", "Name")
	_lambdaCmd.Flags().StringVarP(&_lambdaOrganizationId, "organization-id", "", "", "Organization ID")
	_lambdaCmd.Flags().StringVarP(&_lambdaPackageType, "package-type", "", "", "Package Type")
	_lambdaCmd.Flags().StringVarP(&_lambdaParallelizationFactor, "parallelization-factor", "", "", "Parallelization Factor")
	_lambdaCmd.Flags().StringVarP(&_lambdaPayload, "payload", "", "", "Payload")
	_lambdaCmd.Flags().StringVarP(&_lambdaPermissionsConfig, "permissions-config", "", "", "Permissions Config")
	_lambdaCmd.Flags().StringVarP(&_lambdaPrincipal, "principal", "", "", "Principal")
	_lambdaCmd.Flags().StringVarP(&_lambdaPrincipalOrgID, "principal-org-id", "", "", "Principal Org ID")
	_lambdaCmd.Flags().StringVarP(&_lambdaProvisionedConcurrentExecutions, "provisioned-concurrent-executions", "", "", "Provisioned Concurrent Executions")
	_lambdaCmd.Flags().StringVarP(&_lambdaProvisionedPollerConfig, "provisioned-poller-config", "", "", "Provisioned Poller Config")
	_lambdaCmd.Flags().StringVarP(&_lambdaPublish, "publish", "", "", "Publish")
	_lambdaCmd.Flags().StringVarP(&_lambdaPublishTo, "publish-to", "", "", "Publish To")
	_lambdaCmd.Flags().StringVarP(&_lambdaQualifier, "qualifier", "", "", "Qualifier")
	_lambdaCmd.Flags().StringSliceVarP(&_lambdaQueues, "queues", "", nil, "Queues")
	_lambdaCmd.Flags().StringVarP(&_lambdaRecursiveLoop, "recursive-loop", "", "", "Recursive Loop")
	_lambdaCmd.Flags().StringVarP(&_lambdaReservedConcurrentExecutions, "reserved-concurrent-executions", "", "", "Reserved Concurrent Executions")
	_lambdaCmd.Flags().StringVarP(&_lambdaResource, "resource", "", "", "Resource")
	_lambdaCmd.Flags().StringVarP(&_lambdaResult, "result", "", "", "Result")
	_lambdaCmd.Flags().StringVarP(&_lambdaReverseOrder, "reverse-order", "", "", "Reverse Order")
	_lambdaCmd.Flags().StringVarP(&_lambdaRevisionId, "revision-id", "", "", "Revision ID")
	_lambdaCmd.Flags().StringVarP(&_lambdaRole, "role", "", "", "Role")
	_lambdaCmd.Flags().StringVarP(&_lambdaRoutingConfig, "routing-config", "", "", "Routing Config")
	_lambdaCmd.Flags().StringVarP(&_lambdaRuntime, "runtime", "", "", "Runtime")
	_lambdaCmd.Flags().StringVarP(&_lambdaRuntimeVersionArn, "runtime-version-arn", "", "", "Runtime Version ARN")
	_lambdaCmd.Flags().StringVarP(&_lambdaS3Bucket, "s3-bucket", "", "", "S3 Bucket")
	_lambdaCmd.Flags().StringVarP(&_lambdaS3Key, "s3-key", "", "", "S3 Key")
	_lambdaCmd.Flags().StringVarP(&_lambdaS3ObjectVersion, "s3-object-version", "", "", "S3 Object Version")
	_lambdaCmd.Flags().StringVarP(&_lambdaScalingConfig, "scaling-config", "", "", "Scaling Config")
	_lambdaCmd.Flags().StringVarP(&_lambdaSelfManagedEventSource, "self-managed-event-source", "", "", "Self Managed Event Source")
	_lambdaCmd.Flags().StringVarP(&_lambdaSelfManagedKafkaEventSourceConfig, "self-managed-kafka-event-source-config", "", "", "Self Managed Kafka Event Source Config")
	_lambdaCmd.Flags().StringVarP(&_lambdaSnapStart, "snap-start", "", "", "Snap Start")
	_lambdaCmd.Flags().StringVarP(&_lambdaSourceAccessConfigurations, "source-access-configurations", "", "", "Source Access Configurations")
	_lambdaCmd.Flags().StringVarP(&_lambdaSourceAccount, "source-account", "", "", "Source Account")
	_lambdaCmd.Flags().StringVarP(&_lambdaSourceArn, "source-arn", "", "", "Source ARN")
	_lambdaCmd.Flags().StringVarP(&_lambdaSourceKMSKeyArn, "source-kms-key-arn", "", "", "Source KMS Key ARN")
	_lambdaCmd.Flags().StringVarP(&_lambdaStartedAfter, "started-after", "", "", "Started After")
	_lambdaCmd.Flags().StringVarP(&_lambdaStartedBefore, "started-before", "", "", "Started Before")
	_lambdaCmd.Flags().StringVarP(&_lambdaStartingPosition, "starting-position", "", "", "Starting Position")
	_lambdaCmd.Flags().StringVarP(&_lambdaStartingPositionTimestamp, "starting-position-timestamp", "", "", "Starting Position Timestamp")
	_lambdaCmd.Flags().StringVarP(&_lambdaState, "state", "", "", "State")
	_lambdaCmd.Flags().StringVarP(&_lambdaStatementId, "statement-id", "", "", "Statement ID")
	_lambdaCmd.Flags().StringVarP(&_lambdaStatuses, "statuses", "", "", "Statuses")
	_lambdaCmd.Flags().StringSliceVarP(&_lambdaTagKeys, "tag-keys", "", nil, "Tag Keys")
	_lambdaCmd.Flags().StringVarP(&_lambdaTags, "tags", "", "", "Tags")
	_lambdaCmd.Flags().StringVarP(&_lambdaTenancyConfig, "tenancy-config", "", "", "Tenancy Config")
	_lambdaCmd.Flags().StringVarP(&_lambdaTenantId, "tenant-id", "", "", "Tenant ID")
	_lambdaCmd.Flags().StringVarP(&_lambdaTimeout, "timeout", "", "", "Timeout")
	_lambdaCmd.Flags().StringSliceVarP(&_lambdaTopics, "topics", "", nil, "Topics")
	_lambdaCmd.Flags().StringVarP(&_lambdaTracingConfig, "tracing-config", "", "", "Tracing Config")
	_lambdaCmd.Flags().StringVarP(&_lambdaTumblingWindowInSeconds, "tumbling-window-in-seconds", "", "", "Tumbling Window In Seconds")
	_lambdaCmd.Flags().StringVarP(&_lambdaUpdateRuntimeOn, "update-runtime-on", "", "", "Update Runtime On")
	_lambdaCmd.Flags().StringVarP(&_lambdaUpdates, "updates", "", "", "Updates")
	_lambdaCmd.Flags().StringVarP(&_lambdaUUID, "uuid", "", "", "UUID")
	_lambdaCmd.Flags().StringVarP(&_lambdaVersionNumber, "version-number", "", "", "Version Number")
	_lambdaCmd.Flags().StringVarP(&_lambdaVpcConfig, "vpc-config", "", "", "VPC Config")
	_lambdaCmd.Flags().StringVarP(&_lambdaZipFile, "zip-file", "", "", "Zip File")

	_lambdaCmd.Flags().BoolVarP(&_lambdaAddLayerVersionPermission, "add-layer-version-permission", "", false, "Add Layer Version Permission")
	_lambdaCmd.Flags().BoolVarP(&_lambdaAddPermission, "add-permission", "", false, "Add Permission")
	_lambdaCmd.Flags().BoolVarP(&_lambdaCheckpointDurableExecution, "checkpoint-durable-execution", "", false, "Checkpoint Durable Execution")
	_lambdaCmd.Flags().BoolVarP(&_lambdaCreateAlias, "create-alias", "", false, "Create Alias")
	_lambdaCmd.Flags().BoolVarP(&_lambdaCreateCapacityProvider, "create-capacity-provider", "", false, "Create Capacity Provider")
	_lambdaCmd.Flags().BoolVarP(&_lambdaCreateCodeSigningConfig, "create-code-signing-config", "", false, "Create Code Signing Config")
	_lambdaCmd.Flags().BoolVarP(&_lambdaCreateEventSourceMapping, "create-event-source-mapping", "", false, "Create Event Source Mapping")
	_lambdaCmd.Flags().BoolVarP(&_lambdaCreateFunction, "create-function", "", false, "Create Function")
	_lambdaCmd.Flags().BoolVarP(&_lambdaCreateFunctionUrlConfig, "create-function-url-config", "", false, "Create Function URL Config")
	_lambdaCmd.Flags().BoolVarP(&_lambdaDeleteAlias, "delete-alias", "", false, "Delete Alias")
	_lambdaCmd.Flags().BoolVarP(&_lambdaDeleteCapacityProvider, "delete-capacity-provider", "", false, "Delete Capacity Provider")
	_lambdaCmd.Flags().BoolVarP(&_lambdaDeleteCodeSigningConfig, "delete-code-signing-config", "", false, "Delete Code Signing Config")
	_lambdaCmd.Flags().BoolVarP(&_lambdaDeleteEventSourceMapping, "delete-event-source-mapping", "", false, "Delete Event Source Mapping")
	_lambdaCmd.Flags().BoolVarP(&_lambdaDeleteFunction, "delete-function", "", false, "Delete Function")
	_lambdaCmd.Flags().BoolVarP(&_lambdaDeleteFunctionCodeSigningConfig, "delete-function-code-signing-config", "", false, "Delete Function Code Signing Config")
	_lambdaCmd.Flags().BoolVarP(&_lambdaDeleteFunctionConcurrency, "delete-function-concurrency", "", false, "Delete Function Concurrency")
	_lambdaCmd.Flags().BoolVarP(&_lambdaDeleteFunctionEventInvokeConfig, "delete-function-event-invoke-config", "", false, "Delete Function Event Invoke Config")
	_lambdaCmd.Flags().BoolVarP(&_lambdaDeleteFunctionUrlConfig, "delete-function-url-config", "", false, "Delete Function URL Config")
	_lambdaCmd.Flags().BoolVarP(&_lambdaDeleteLayerVersion, "delete-layer-version", "", false, "Delete Layer Version")
	_lambdaCmd.Flags().BoolVarP(&_lambdaDeleteProvisionedConcurrencyConfig, "delete-provisioned-concurrency-config", "", false, "Delete Provisioned Concurrency Config")
	_lambdaCmd.Flags().BoolVarP(&_lambdaGetAccountSettings, "get-account-settings", "", false, "Get Account Settings")
	_lambdaCmd.Flags().BoolVarP(&_lambdaGetAlias, "get-alias", "", false, "Get Alias")
	_lambdaCmd.Flags().BoolVarP(&_lambdaGetCapacityProvider, "get-capacity-provider", "", false, "Get Capacity Provider")
	_lambdaCmd.Flags().BoolVarP(&_lambdaGetCodeSigningConfig, "get-code-signing-config", "", false, "Get Code Signing Config")
	_lambdaCmd.Flags().BoolVarP(&_lambdaGetDurableExecution, "get-durable-execution", "", false, "Get Durable Execution")
	_lambdaCmd.Flags().BoolVarP(&_lambdaGetDurableExecutionHistory, "get-durable-execution-history", "", false, "Get Durable Execution History")
	_lambdaCmd.Flags().BoolVarP(&_lambdaGetDurableExecutionState, "get-durable-execution-state", "", false, "Get Durable Execution State")
	_lambdaCmd.Flags().BoolVarP(&_lambdaGetEventSourceMapping, "get-event-source-mapping", "", false, "Get Event Source Mapping")
	_lambdaCmd.Flags().BoolVarP(&_lambdaGetFunction, "get-function", "", false, "Get Function")
	_lambdaCmd.Flags().BoolVarP(&_lambdaGetFunctionCodeSigningConfig, "get-function-code-signing-config", "", false, "Get Function Code Signing Config")
	_lambdaCmd.Flags().BoolVarP(&_lambdaGetFunctionConcurrency, "get-function-concurrency", "", false, "Get Function Concurrency")
	_lambdaCmd.Flags().BoolVarP(&_lambdaGetFunctionConfiguration, "get-function-configuration", "", false, "Get Function Configuration")
	_lambdaCmd.Flags().BoolVarP(&_lambdaGetFunctionEventInvokeConfig, "get-function-event-invoke-config", "", false, "Get Function Event Invoke Config")
	_lambdaCmd.Flags().BoolVarP(&_lambdaGetFunctionRecursionConfig, "get-function-recursion-config", "", false, "Get Function Recursion Config")
	_lambdaCmd.Flags().BoolVarP(&_lambdaGetFunctionScalingConfig, "get-function-scaling-config", "", false, "Get Function Scaling Config")
	_lambdaCmd.Flags().BoolVarP(&_lambdaGetFunctionUrlConfig, "get-function-url-config", "", false, "Get Function URL Config")
	_lambdaCmd.Flags().BoolVarP(&_lambdaGetLayerVersion, "get-layer-version", "", false, "Get Layer Version")
	_lambdaCmd.Flags().BoolVarP(&_lambdaGetLayerVersionByArn, "get-layer-version-by-arn", "", false, "Get Layer Version By ARN")
	_lambdaCmd.Flags().BoolVarP(&_lambdaGetLayerVersionPolicy, "get-layer-version-policy", "", false, "Get Layer Version Policy")
	_lambdaCmd.Flags().BoolVarP(&_lambdaGetPolicy, "get-policy", "", false, "Get Policy")
	_lambdaCmd.Flags().BoolVarP(&_lambdaGetProvisionedConcurrencyConfig, "get-provisioned-concurrency-config", "", false, "Get Provisioned Concurrency Config")
	_lambdaCmd.Flags().BoolVarP(&_lambdaGetRuntimeManagementConfig, "get-runtime-management-config", "", false, "Get Runtime Management Config")
	_lambdaCmd.Flags().BoolVarP(&_lambdaInvoke, "invoke", "", false, "Invoke")
	_lambdaCmd.Flags().BoolVarP(&_lambdaInvokeAsync, "invoke-async", "", false, "Invoke Async")
	_lambdaCmd.Flags().BoolVarP(&_lambdaInvokeWithResponseStream, "invoke-with-response-stream", "", false, "Invoke With Response Stream")
	_lambdaCmd.Flags().BoolVarP(&_lambdaListAliases, "list-aliases", "", false, "List Aliases")
	_lambdaCmd.Flags().BoolVarP(&_lambdaListCapacityProviders, "list-capacity-providers", "", false, "List Capacity Providers")
	_lambdaCmd.Flags().BoolVarP(&_lambdaListCodeSigningConfigs, "list-code-signing-configs", "", false, "List Code Signing Configs")
	_lambdaCmd.Flags().BoolVarP(&_lambdaListDurableExecutionsByFunction, "list-durable-executions-by-function", "", false, "List Durable Executions By Function")
	_lambdaCmd.Flags().BoolVarP(&_lambdaListEventSourceMappings, "list-event-source-mappings", "", false, "List Event Source Mappings")
	_lambdaCmd.Flags().BoolVarP(&_lambdaListFunctionEventInvokeConfigs, "list-function-event-invoke-configs", "", false, "List Function Event Invoke Configs")
	_lambdaCmd.Flags().BoolVarP(&_lambdaListFunctionUrlConfigs, "list-function-url-configs", "", false, "List Function URL Configs")
	_lambdaCmd.Flags().BoolVarP(&_lambdaListFunctionVersionsByCapacityProvider, "list-function-versions-by-capacity-provider", "", false, "List Function Versions By Capacity Provider")
	_lambdaCmd.Flags().BoolVarP(&_lambdaListFunctions, "list-functions", "", false, "List Functions")
	_lambdaCmd.Flags().BoolVarP(&_lambdaListFunctionsByCodeSigningConfig, "list-functions-by-code-signing-config", "", false, "List Functions By Code Signing Config")
	_lambdaCmd.Flags().BoolVarP(&_lambdaListLayerVersions, "list-layer-versions", "", false, "List Layer Versions")
	_lambdaCmd.Flags().BoolVarP(&_lambdaListLayers, "list-layers", "", false, "List Layers")
	_lambdaCmd.Flags().BoolVarP(&_lambdaListProvisionedConcurrencyConfigs, "list-provisioned-concurrency-configs", "", false, "List Provisioned Concurrency Configs")
	_lambdaCmd.Flags().BoolVarP(&_lambdaListTags, "list-tags", "", false, "List Tags")
	_lambdaCmd.Flags().BoolVarP(&_lambdaListVersionsByFunction, "list-versions-by-function", "", false, "List Versions By Function")
	_lambdaCmd.Flags().BoolVarP(&_lambdaPublishLayerVersion, "publish-layer-version", "", false, "Publish Layer Version")
	_lambdaCmd.Flags().BoolVarP(&_lambdaPublishVersion, "publish-version", "", false, "Publish Version")
	_lambdaCmd.Flags().BoolVarP(&_lambdaPutFunctionCodeSigningConfig, "put-function-code-signing-config", "", false, "Put Function Code Signing Config")
	_lambdaCmd.Flags().BoolVarP(&_lambdaPutFunctionConcurrency, "put-function-concurrency", "", false, "Put Function Concurrency")
	_lambdaCmd.Flags().BoolVarP(&_lambdaPutFunctionEventInvokeConfig, "put-function-event-invoke-config", "", false, "Put Function Event Invoke Config")
	_lambdaCmd.Flags().BoolVarP(&_lambdaPutFunctionRecursionConfig, "put-function-recursion-config", "", false, "Put Function Recursion Config")
	_lambdaCmd.Flags().BoolVarP(&_lambdaPutFunctionScalingConfig, "put-function-scaling-config", "", false, "Put Function Scaling Config")
	_lambdaCmd.Flags().BoolVarP(&_lambdaPutProvisionedConcurrencyConfig, "put-provisioned-concurrency-config", "", false, "Put Provisioned Concurrency Config")
	_lambdaCmd.Flags().BoolVarP(&_lambdaPutRuntimeManagementConfig, "put-runtime-management-config", "", false, "Put Runtime Management Config")
	_lambdaCmd.Flags().BoolVarP(&_lambdaRemoveLayerVersionPermission, "remove-layer-version-permission", "", false, "Remove Layer Version Permission")
	_lambdaCmd.Flags().BoolVarP(&_lambdaRemovePermission, "remove-permission", "", false, "Remove Permission")
	_lambdaCmd.Flags().BoolVarP(&_lambdaSendDurableExecutionCallbackFailure, "send-durable-execution-callback-failure", "", false, "Send Durable Execution Callback Failure")
	_lambdaCmd.Flags().BoolVarP(&_lambdaSendDurableExecutionCallbackHeartbeat, "send-durable-execution-callback-heartbeat", "", false, "Send Durable Execution Callback Heartbeat")
	_lambdaCmd.Flags().BoolVarP(&_lambdaSendDurableExecutionCallbackSuccess, "send-durable-execution-callback-success", "", false, "Send Durable Execution Callback Success")
	_lambdaCmd.Flags().BoolVarP(&_lambdaStopDurableExecution, "stop-durable-execution", "", false, "Stop Durable Execution")
	_lambdaCmd.Flags().BoolVarP(&_lambdaTagResource, "tag-resource", "", false, "Tag Resource")
	_lambdaCmd.Flags().BoolVarP(&_lambdaUntagResource, "untag-resource", "", false, "Untag Resource")
	_lambdaCmd.Flags().BoolVarP(&_lambdaUpdateAlias, "update-alias", "", false, "Update Alias")
	_lambdaCmd.Flags().BoolVarP(&_lambdaUpdateCapacityProvider, "update-capacity-provider", "", false, "Update Capacity Provider")
	_lambdaCmd.Flags().BoolVarP(&_lambdaUpdateCodeSigningConfig, "update-code-signing-config", "", false, "Update Code Signing Config")
	_lambdaCmd.Flags().BoolVarP(&_lambdaUpdateEventSourceMapping, "update-event-source-mapping", "", false, "Update Event Source Mapping")
	_lambdaCmd.Flags().BoolVarP(&_lambdaUpdateFunctionCode, "update-function-code", "", false, "Update Function Code")
	_lambdaCmd.Flags().BoolVarP(&_lambdaUpdateFunctionConfiguration, "update-function-configuration", "", false, "Update Function Configuration")
	_lambdaCmd.Flags().BoolVarP(&_lambdaUpdateFunctionEventInvokeConfig, "update-function-event-invoke-config", "", false, "Update Function Event Invoke Config")
	_lambdaCmd.Flags().BoolVarP(&_lambdaUpdateFunctionUrlConfig, "update-function-url-config", "", false, "Update Function URL Config")

}
