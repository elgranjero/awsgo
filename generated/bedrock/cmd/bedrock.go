package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/bedrock"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// bedrockCmd represents the bedrock command
var _bedrockCmd = &cobra.Command{
	Use:   "bedrock",
	Short: "AWS bedrock CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := bedrock.NewFromConfig(cfg)
		if _bedrockBatchDeleteEvaluationJob {
			bedrock_BatchDeleteEvaluationJob(cfg, client)
			return
		}
		if _bedrockCancelAutomatedReasoningPolicyBuildWorkflow {
			bedrock_CancelAutomatedReasoningPolicyBuildWorkflow(cfg, client)
			return
		}
		if _bedrockCreateAutomatedReasoningPolicy {
			bedrock_CreateAutomatedReasoningPolicy(cfg, client)
			return
		}
		if _bedrockCreateAutomatedReasoningPolicyTestCase {
			bedrock_CreateAutomatedReasoningPolicyTestCase(cfg, client)
			return
		}
		if _bedrockCreateAutomatedReasoningPolicyVersion {
			bedrock_CreateAutomatedReasoningPolicyVersion(cfg, client)
			return
		}
		if _bedrockCreateCustomModel {
			bedrock_CreateCustomModel(cfg, client)
			return
		}
		if _bedrockCreateCustomModelDeployment {
			bedrock_CreateCustomModelDeployment(cfg, client)
			return
		}
		if _bedrockCreateEvaluationJob {
			bedrock_CreateEvaluationJob(cfg, client)
			return
		}
		if _bedrockCreateFoundationModelAgreement {
			bedrock_CreateFoundationModelAgreement(cfg, client)
			return
		}
		if _bedrockCreateGuardrail {
			bedrock_CreateGuardrail(cfg, client)
			return
		}
		if _bedrockCreateGuardrailVersion {
			bedrock_CreateGuardrailVersion(cfg, client)
			return
		}
		if _bedrockCreateInferenceProfile {
			bedrock_CreateInferenceProfile(cfg, client)
			return
		}
		if _bedrockCreateMarketplaceModelEndpoint {
			bedrock_CreateMarketplaceModelEndpoint(cfg, client)
			return
		}
		if _bedrockCreateModelCopyJob {
			bedrock_CreateModelCopyJob(cfg, client)
			return
		}
		if _bedrockCreateModelCustomizationJob {
			bedrock_CreateModelCustomizationJob(cfg, client)
			return
		}
		if _bedrockCreateModelImportJob {
			bedrock_CreateModelImportJob(cfg, client)
			return
		}
		if _bedrockCreateModelInvocationJob {
			bedrock_CreateModelInvocationJob(cfg, client)
			return
		}
		if _bedrockCreatePromptRouter {
			bedrock_CreatePromptRouter(cfg, client)
			return
		}
		if _bedrockCreateProvisionedModelThroughput {
			bedrock_CreateProvisionedModelThroughput(cfg, client)
			return
		}
		if _bedrockDeleteAutomatedReasoningPolicy {
			bedrock_DeleteAutomatedReasoningPolicy(cfg, client)
			return
		}
		if _bedrockDeleteAutomatedReasoningPolicyBuildWorkflow {
			bedrock_DeleteAutomatedReasoningPolicyBuildWorkflow(cfg, client)
			return
		}
		if _bedrockDeleteAutomatedReasoningPolicyTestCase {
			bedrock_DeleteAutomatedReasoningPolicyTestCase(cfg, client)
			return
		}
		if _bedrockDeleteCustomModel {
			bedrock_DeleteCustomModel(cfg, client)
			return
		}
		if _bedrockDeleteCustomModelDeployment {
			bedrock_DeleteCustomModelDeployment(cfg, client)
			return
		}
		if _bedrockDeleteEnforcedGuardrailConfiguration {
			bedrock_DeleteEnforcedGuardrailConfiguration(cfg, client)
			return
		}
		if _bedrockDeleteFoundationModelAgreement {
			bedrock_DeleteFoundationModelAgreement(cfg, client)
			return
		}
		if _bedrockDeleteGuardrail {
			bedrock_DeleteGuardrail(cfg, client)
			return
		}
		if _bedrockDeleteImportedModel {
			bedrock_DeleteImportedModel(cfg, client)
			return
		}
		if _bedrockDeleteInferenceProfile {
			bedrock_DeleteInferenceProfile(cfg, client)
			return
		}
		if _bedrockDeleteMarketplaceModelEndpoint {
			bedrock_DeleteMarketplaceModelEndpoint(cfg, client)
			return
		}
		if _bedrockDeleteModelInvocationLoggingConfiguration {
			bedrock_DeleteModelInvocationLoggingConfiguration(cfg, client)
			return
		}
		if _bedrockDeletePromptRouter {
			bedrock_DeletePromptRouter(cfg, client)
			return
		}
		if _bedrockDeleteProvisionedModelThroughput {
			bedrock_DeleteProvisionedModelThroughput(cfg, client)
			return
		}
		if _bedrockDeregisterMarketplaceModelEndpoint {
			bedrock_DeregisterMarketplaceModelEndpoint(cfg, client)
			return
		}
		if _bedrockExportAutomatedReasoningPolicyVersion {
			bedrock_ExportAutomatedReasoningPolicyVersion(cfg, client)
			return
		}
		if _bedrockGetAutomatedReasoningPolicy {
			bedrock_GetAutomatedReasoningPolicy(cfg, client)
			return
		}
		if _bedrockGetAutomatedReasoningPolicyAnnotations {
			bedrock_GetAutomatedReasoningPolicyAnnotations(cfg, client)
			return
		}
		if _bedrockGetAutomatedReasoningPolicyBuildWorkflow {
			bedrock_GetAutomatedReasoningPolicyBuildWorkflow(cfg, client)
			return
		}
		if _bedrockGetAutomatedReasoningPolicyBuildWorkflowResultAssets {
			bedrock_GetAutomatedReasoningPolicyBuildWorkflowResultAssets(cfg, client)
			return
		}
		if _bedrockGetAutomatedReasoningPolicyNextScenario {
			bedrock_GetAutomatedReasoningPolicyNextScenario(cfg, client)
			return
		}
		if _bedrockGetAutomatedReasoningPolicyTestCase {
			bedrock_GetAutomatedReasoningPolicyTestCase(cfg, client)
			return
		}
		if _bedrockGetAutomatedReasoningPolicyTestResult {
			bedrock_GetAutomatedReasoningPolicyTestResult(cfg, client)
			return
		}
		if _bedrockGetCustomModel {
			bedrock_GetCustomModel(cfg, client)
			return
		}
		if _bedrockGetCustomModelDeployment {
			bedrock_GetCustomModelDeployment(cfg, client)
			return
		}
		if _bedrockGetEvaluationJob {
			bedrock_GetEvaluationJob(cfg, client)
			return
		}
		if _bedrockGetFoundationModel {
			bedrock_GetFoundationModel(cfg, client)
			return
		}
		if _bedrockGetFoundationModelAvailability {
			bedrock_GetFoundationModelAvailability(cfg, client)
			return
		}
		if _bedrockGetGuardrail {
			bedrock_GetGuardrail(cfg, client)
			return
		}
		if _bedrockGetImportedModel {
			bedrock_GetImportedModel(cfg, client)
			return
		}
		if _bedrockGetInferenceProfile {
			bedrock_GetInferenceProfile(cfg, client)
			return
		}
		if _bedrockGetMarketplaceModelEndpoint {
			bedrock_GetMarketplaceModelEndpoint(cfg, client)
			return
		}
		if _bedrockGetModelCopyJob {
			bedrock_GetModelCopyJob(cfg, client)
			return
		}
		if _bedrockGetModelCustomizationJob {
			bedrock_GetModelCustomizationJob(cfg, client)
			return
		}
		if _bedrockGetModelImportJob {
			bedrock_GetModelImportJob(cfg, client)
			return
		}
		if _bedrockGetModelInvocationJob {
			bedrock_GetModelInvocationJob(cfg, client)
			return
		}
		if _bedrockGetModelInvocationLoggingConfiguration {
			bedrock_GetModelInvocationLoggingConfiguration(cfg, client)
			return
		}
		if _bedrockGetPromptRouter {
			bedrock_GetPromptRouter(cfg, client)
			return
		}
		if _bedrockGetProvisionedModelThroughput {
			bedrock_GetProvisionedModelThroughput(cfg, client)
			return
		}
		if _bedrockGetUseCaseForModelAccess {
			bedrock_GetUseCaseForModelAccess(cfg, client)
			return
		}
		if _bedrockListAutomatedReasoningPolicies {
			bedrock_ListAutomatedReasoningPolicies(cfg, client)
			return
		}
		if _bedrockListAutomatedReasoningPolicyBuildWorkflows {
			bedrock_ListAutomatedReasoningPolicyBuildWorkflows(cfg, client)
			return
		}
		if _bedrockListAutomatedReasoningPolicyTestCases {
			bedrock_ListAutomatedReasoningPolicyTestCases(cfg, client)
			return
		}
		if _bedrockListAutomatedReasoningPolicyTestResults {
			bedrock_ListAutomatedReasoningPolicyTestResults(cfg, client)
			return
		}
		if _bedrockListCustomModelDeployments {
			bedrock_ListCustomModelDeployments(cfg, client)
			return
		}
		if _bedrockListCustomModels {
			bedrock_ListCustomModels(cfg, client)
			return
		}
		if _bedrockListEnforcedGuardrailsConfiguration {
			bedrock_ListEnforcedGuardrailsConfiguration(cfg, client)
			return
		}
		if _bedrockListEvaluationJobs {
			bedrock_ListEvaluationJobs(cfg, client)
			return
		}
		if _bedrockListFoundationModelAgreementOffers {
			bedrock_ListFoundationModelAgreementOffers(cfg, client)
			return
		}
		if _bedrockListFoundationModels {
			bedrock_ListFoundationModels(cfg, client)
			return
		}
		if _bedrockListGuardrails {
			bedrock_ListGuardrails(cfg, client)
			return
		}
		if _bedrockListImportedModels {
			bedrock_ListImportedModels(cfg, client)
			return
		}
		if _bedrockListInferenceProfiles {
			bedrock_ListInferenceProfiles(cfg, client)
			return
		}
		if _bedrockListMarketplaceModelEndpoints {
			bedrock_ListMarketplaceModelEndpoints(cfg, client)
			return
		}
		if _bedrockListModelCopyJobs {
			bedrock_ListModelCopyJobs(cfg, client)
			return
		}
		if _bedrockListModelCustomizationJobs {
			bedrock_ListModelCustomizationJobs(cfg, client)
			return
		}
		if _bedrockListModelImportJobs {
			bedrock_ListModelImportJobs(cfg, client)
			return
		}
		if _bedrockListModelInvocationJobs {
			bedrock_ListModelInvocationJobs(cfg, client)
			return
		}
		if _bedrockListPromptRouters {
			bedrock_ListPromptRouters(cfg, client)
			return
		}
		if _bedrockListProvisionedModelThroughputs {
			bedrock_ListProvisionedModelThroughputs(cfg, client)
			return
		}
		if _bedrockListTagsForResource {
			bedrock_ListTagsForResource(cfg, client)
			return
		}
		if _bedrockPutEnforcedGuardrailConfiguration {
			bedrock_PutEnforcedGuardrailConfiguration(cfg, client)
			return
		}
		if _bedrockPutModelInvocationLoggingConfiguration {
			bedrock_PutModelInvocationLoggingConfiguration(cfg, client)
			return
		}
		if _bedrockPutUseCaseForModelAccess {
			bedrock_PutUseCaseForModelAccess(cfg, client)
			return
		}
		if _bedrockRegisterMarketplaceModelEndpoint {
			bedrock_RegisterMarketplaceModelEndpoint(cfg, client)
			return
		}
		if _bedrockStartAutomatedReasoningPolicyBuildWorkflow {
			bedrock_StartAutomatedReasoningPolicyBuildWorkflow(cfg, client)
			return
		}
		if _bedrockStartAutomatedReasoningPolicyTestWorkflow {
			bedrock_StartAutomatedReasoningPolicyTestWorkflow(cfg, client)
			return
		}
		if _bedrockStopEvaluationJob {
			bedrock_StopEvaluationJob(cfg, client)
			return
		}
		if _bedrockStopModelCustomizationJob {
			bedrock_StopModelCustomizationJob(cfg, client)
			return
		}
		if _bedrockStopModelInvocationJob {
			bedrock_StopModelInvocationJob(cfg, client)
			return
		}
		if _bedrockTagResource {
			bedrock_TagResource(cfg, client)
			return
		}
		if _bedrockUntagResource {
			bedrock_UntagResource(cfg, client)
			return
		}
		if _bedrockUpdateAutomatedReasoningPolicy {
			bedrock_UpdateAutomatedReasoningPolicy(cfg, client)
			return
		}
		if _bedrockUpdateAutomatedReasoningPolicyAnnotations {
			bedrock_UpdateAutomatedReasoningPolicyAnnotations(cfg, client)
			return
		}
		if _bedrockUpdateAutomatedReasoningPolicyTestCase {
			bedrock_UpdateAutomatedReasoningPolicyTestCase(cfg, client)
			return
		}
		if _bedrockUpdateCustomModelDeployment {
			bedrock_UpdateCustomModelDeployment(cfg, client)
			return
		}
		if _bedrockUpdateGuardrail {
			bedrock_UpdateGuardrail(cfg, client)
			return
		}
		if _bedrockUpdateMarketplaceModelEndpoint {
			bedrock_UpdateMarketplaceModelEndpoint(cfg, client)
			return
		}
		if _bedrockUpdateProvisionedModelThroughput {
			bedrock_UpdateProvisionedModelThroughput(cfg, client)
			return
		}

	},
}

var (
	_bedrockBatchDeleteEvaluationJob                             bool
	_bedrockCancelAutomatedReasoningPolicyBuildWorkflow          bool
	_bedrockCreateAutomatedReasoningPolicy                       bool
	_bedrockCreateAutomatedReasoningPolicyTestCase               bool
	_bedrockCreateAutomatedReasoningPolicyVersion                bool
	_bedrockCreateCustomModel                                    bool
	_bedrockCreateCustomModelDeployment                          bool
	_bedrockCreateEvaluationJob                                  bool
	_bedrockCreateFoundationModelAgreement                       bool
	_bedrockCreateGuardrail                                      bool
	_bedrockCreateGuardrailVersion                               bool
	_bedrockCreateInferenceProfile                               bool
	_bedrockCreateMarketplaceModelEndpoint                       bool
	_bedrockCreateModelCopyJob                                   bool
	_bedrockCreateModelCustomizationJob                          bool
	_bedrockCreateModelImportJob                                 bool
	_bedrockCreateModelInvocationJob                             bool
	_bedrockCreatePromptRouter                                   bool
	_bedrockCreateProvisionedModelThroughput                     bool
	_bedrockDeleteAutomatedReasoningPolicy                       bool
	_bedrockDeleteAutomatedReasoningPolicyBuildWorkflow          bool
	_bedrockDeleteAutomatedReasoningPolicyTestCase               bool
	_bedrockDeleteCustomModel                                    bool
	_bedrockDeleteCustomModelDeployment                          bool
	_bedrockDeleteEnforcedGuardrailConfiguration                 bool
	_bedrockDeleteFoundationModelAgreement                       bool
	_bedrockDeleteGuardrail                                      bool
	_bedrockDeleteImportedModel                                  bool
	_bedrockDeleteInferenceProfile                               bool
	_bedrockDeleteMarketplaceModelEndpoint                       bool
	_bedrockDeleteModelInvocationLoggingConfiguration            bool
	_bedrockDeletePromptRouter                                   bool
	_bedrockDeleteProvisionedModelThroughput                     bool
	_bedrockDeregisterMarketplaceModelEndpoint                   bool
	_bedrockExportAutomatedReasoningPolicyVersion                bool
	_bedrockGetAutomatedReasoningPolicy                          bool
	_bedrockGetAutomatedReasoningPolicyAnnotations               bool
	_bedrockGetAutomatedReasoningPolicyBuildWorkflow             bool
	_bedrockGetAutomatedReasoningPolicyBuildWorkflowResultAssets bool
	_bedrockGetAutomatedReasoningPolicyNextScenario              bool
	_bedrockGetAutomatedReasoningPolicyTestCase                  bool
	_bedrockGetAutomatedReasoningPolicyTestResult                bool
	_bedrockGetCustomModel                                       bool
	_bedrockGetCustomModelDeployment                             bool
	_bedrockGetEvaluationJob                                     bool
	_bedrockGetFoundationModel                                   bool
	_bedrockGetFoundationModelAvailability                       bool
	_bedrockGetGuardrail                                         bool
	_bedrockGetImportedModel                                     bool
	_bedrockGetInferenceProfile                                  bool
	_bedrockGetMarketplaceModelEndpoint                          bool
	_bedrockGetModelCopyJob                                      bool
	_bedrockGetModelCustomizationJob                             bool
	_bedrockGetModelImportJob                                    bool
	_bedrockGetModelInvocationJob                                bool
	_bedrockGetModelInvocationLoggingConfiguration               bool
	_bedrockGetPromptRouter                                      bool
	_bedrockGetProvisionedModelThroughput                        bool
	_bedrockGetUseCaseForModelAccess                             bool
	_bedrockListAutomatedReasoningPolicies                       bool
	_bedrockListAutomatedReasoningPolicyBuildWorkflows           bool
	_bedrockListAutomatedReasoningPolicyTestCases                bool
	_bedrockListAutomatedReasoningPolicyTestResults              bool
	_bedrockListCustomModelDeployments                           bool
	_bedrockListCustomModels                                     bool
	_bedrockListEnforcedGuardrailsConfiguration                  bool
	_bedrockListEvaluationJobs                                   bool
	_bedrockListFoundationModelAgreementOffers                   bool
	_bedrockListFoundationModels                                 bool
	_bedrockListGuardrails                                       bool
	_bedrockListImportedModels                                   bool
	_bedrockListInferenceProfiles                                bool
	_bedrockListMarketplaceModelEndpoints                        bool
	_bedrockListModelCopyJobs                                    bool
	_bedrockListModelCustomizationJobs                           bool
	_bedrockListModelImportJobs                                  bool
	_bedrockListModelInvocationJobs                              bool
	_bedrockListPromptRouters                                    bool
	_bedrockListProvisionedModelThroughputs                      bool
	_bedrockListTagsForResource                                  bool
	_bedrockPutEnforcedGuardrailConfiguration                    bool
	_bedrockPutModelInvocationLoggingConfiguration               bool
	_bedrockPutUseCaseForModelAccess                             bool
	_bedrockRegisterMarketplaceModelEndpoint                     bool
	_bedrockStartAutomatedReasoningPolicyBuildWorkflow           bool
	_bedrockStartAutomatedReasoningPolicyTestWorkflow            bool
	_bedrockStopEvaluationJob                                    bool
	_bedrockStopModelCustomizationJob                            bool
	_bedrockStopModelInvocationJob                               bool
	_bedrockTagResource                                          bool
	_bedrockUntagResource                                        bool
	_bedrockUpdateAutomatedReasoningPolicy                       bool
	_bedrockUpdateAutomatedReasoningPolicyAnnotations            bool
	_bedrockUpdateAutomatedReasoningPolicyTestCase               bool
	_bedrockUpdateCustomModelDeployment                          bool
	_bedrockUpdateGuardrail                                      bool
	_bedrockUpdateMarketplaceModelEndpoint                       bool
	_bedrockUpdateProvisionedModelThroughput                     bool

	_bedrockAcceptEula                       string
	_bedrockAnnotations                      string
	_bedrockApplicationType                  string
	_bedrockApplicationTypeEquals            string
	_bedrockAssetId                          string
	_bedrockAssetType                        string
	_bedrockAutomatedReasoningPolicyConfig   string
	_bedrockBaseModelArnEquals               string
	_bedrockBaseModelIdentifier              string
	_bedrockBlockedInputMessaging            string
	_bedrockBlockedOutputsMessaging          string
	_bedrockBuildWorkflowId                  string
	_bedrockBuildWorkflowType                string
	_bedrockByCustomizationType              string
	_bedrockByInferenceType                  string
	_bedrockByOutputModality                 string
	_bedrockByProvider                       string
	_bedrockClientRequestToken               string
	_bedrockCommitmentDuration               string
	_bedrockConfidenceThreshold              string
	_bedrockConfigId                         string
	_bedrockContentPolicyConfig              string
	_bedrockContextualGroundingPolicyConfig  string
	_bedrockCreatedAfter                     string
	_bedrockCreatedBefore                    string
	_bedrockCreationTimeAfter                string
	_bedrockCreationTimeBefore               string
	_bedrockCrossRegionConfig                string
	_bedrockCustomModelDeploymentIdentifier  string
	_bedrockCustomModelKmsKeyId              string
	_bedrockCustomModelName                  string
	_bedrockCustomModelTags                  string
	_bedrockCustomerEncryptionKeyId          string
	_bedrockCustomizationConfig              string
	_bedrockCustomizationType                string
	_bedrockDescription                      string
	_bedrockDesiredModelId                   string
	_bedrockDesiredProvisionedModelName      string
	_bedrockEndpointArn                      string
	_bedrockEndpointConfig                   string
	_bedrockEndpointIdentifier               string
	_bedrockEndpointName                     string
	_bedrockEvaluationConfig                 string
	_bedrockExpectedAggregatedFindingsResult string
	_bedrockFallbackModel                    string
	_bedrockForce                            string
	_bedrockFormData                         string
	_bedrockFoundationModelArnEquals         string
	_bedrockGuardContent                     string
	_bedrockGuardrailIdentifier              string
	_bedrockGuardrailInferenceConfig         string
	_bedrockGuardrailVersion                 string
	_bedrockHyperParameters                  string
	_bedrockImportedModelKmsKeyId            string
	_bedrockImportedModelName                string
	_bedrockImportedModelTags                string
	_bedrockInferenceConfig                  string
	_bedrockInferenceProfileIdentifier       string
	_bedrockInferenceProfileName             string
	_bedrockInputDataConfig                  string
	_bedrockIsOwned                          string
	_bedrockJobArn                           string
	_bedrockJobDescription                   string
	_bedrockJobIdentifier                    string
	_bedrockJobIdentifiers                   []string
	_bedrockJobName                          string
	_bedrockJobTags                          string
	_bedrockKmsKeyId                         string
	_bedrockLastUpdatedAnnotationSetHash     string
	_bedrockLastUpdatedAt                    string
	_bedrockLastUpdatedDefinitionHash        string
	_bedrockLoggingConfig                    string
	_bedrockMaxResults                       string
	_bedrockModelArn                         string
	_bedrockModelArnEquals                   string
	_bedrockModelDataSource                  string
	_bedrockModelDeploymentName              string
	_bedrockModelId                          string
	_bedrockModelIdentifier                  string
	_bedrockModelInvocationType              string
	_bedrockModelKmsKeyArn                   string
	_bedrockModelKmsKeyId                    string
	_bedrockModelName                        string
	_bedrockModelSource                      string
	_bedrockModelSourceConfig                string
	_bedrockModelSourceEquals                string
	_bedrockModelSourceIdentifier            string
	_bedrockModelStatus                      string
	_bedrockModelTags                        string
	_bedrockModelUnits                       string
	_bedrockModels                           string
	_bedrockName                             string
	_bedrockNameContains                     string
	_bedrockNextToken                        string
	_bedrockOfferToken                       string
	_bedrockOfferType                        string
	_bedrockOutputDataConfig                 string
	_bedrockPolicyArn                        string
	_bedrockPolicyDefinition                 string
	_bedrockPromptRouterArn                  string
	_bedrockPromptRouterName                 string
	_bedrockProvisionedModelId               string
	_bedrockProvisionedModelName             string
	_bedrockQueryContent                     string
	_bedrockResourceARN                      string
	_bedrockRoleArn                          string
	_bedrockRoutingCriteria                  string
	_bedrockSensitiveInformationPolicyConfig string
	_bedrockSortBy                           string
	_bedrockSortOrder                        string
	_bedrockSourceAccountEquals              string
	_bedrockSourceContent                    string
	_bedrockSourceModelArn                   string
	_bedrockSourceModelArnEquals             string
	_bedrockStatusEquals                     string
	_bedrockSubmitTimeAfter                  string
	_bedrockSubmitTimeBefore                 string
	_bedrockTagKeys                          []string
	_bedrockTags                             string
	_bedrockTargetModelName                  string
	_bedrockTargetModelNameContains          string
	_bedrockTargetModelTags                  string
	_bedrockTestCaseId                       string
	_bedrockTestCaseIds                      []string
	_bedrockTimeoutDurationInHours           string
	_bedrockTopicPolicyConfig                string
	_bedrockTrainingDataConfig               string
	_bedrockType                             string
	_bedrockTypeEquals                       string
	_bedrockValidationDataConfig             string
	_bedrockVpcConfig                        string
	_bedrockWordPolicyConfig                 string
)

// Deletes a batch of evaluation jobs. An evaluation job can only be deleted if it
// has following status FAILED , COMPLETED , and STOPPED . You can request up to 25
// model evaluation jobs be deleted in a single request.
func bedrock_BatchDeleteEvaluationJob(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.BatchDeleteEvaluationJobInput{
		// JobIdentifiers: []string, // Required
	}

	if len(_bedrockJobIdentifiers) > 0 {
		input.JobIdentifiers = append([]string(nil), _bedrockJobIdentifiers...)
	}

	if resp, err := client.BatchDeleteEvaluationJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels a running Automated Reasoning policy build workflow. This stops the
// policy generation process and prevents further processing of the source
// documents.
func bedrock_CancelAutomatedReasoningPolicyBuildWorkflow(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.CancelAutomatedReasoningPolicyBuildWorkflowInput{
		// BuildWorkflowId: *string, // Required
		// PolicyArn: *string, // Required
	}

	if len(_bedrockBuildWorkflowId) > 0 {
		input.BuildWorkflowId = aws.String(_bedrockBuildWorkflowId)
	}
	if len(_bedrockPolicyArn) > 0 {
		input.PolicyArn = aws.String(_bedrockPolicyArn)
	}

	if resp, err := client.CancelAutomatedReasoningPolicyBuildWorkflow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Automated Reasoning policy for Amazon Bedrock Guardrails. Automated
// Reasoning policies use mathematical techniques to detect hallucinations, suggest
// corrections, and highlight unstated assumptions in the responses of your GenAI
// application.
//
// To create a policy, you upload a source document that describes the rules that
// you're encoding. Automated Reasoning extracts important concepts from the source
// document that will become variables in the policy and infers policy rules.
func bedrock_CreateAutomatedReasoningPolicy(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.CreateAutomatedReasoningPolicyInput{
		// Name: *string, // Required
	}

	if len(_bedrockName) > 0 {
		input.Name = aws.String(_bedrockName)
	}
	if len(_bedrockClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_bedrockClientRequestToken)
	}
	if len(_bedrockDescription) > 0 {
		input.Description = aws.String(_bedrockDescription)
	}
	if len(_bedrockKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_bedrockKmsKeyId)
	}
	if len(_bedrockPolicyDefinition) > 0 {
		if err := assignInputField(input, "PolicyDefinition", _bedrockPolicyDefinition); err != nil {
			log.Errorf("invalid --policy-definition: %s", err.Error())
			return
		}
	}
	if len(_bedrockTags) > 0 {
		if err := assignInputField(input, "Tags", _bedrockTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAutomatedReasoningPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a test for an Automated Reasoning policy. Tests validate that your
// policy works as expected by providing sample inputs and expected outcomes. Use
// tests to verify policy behavior before deploying to production.
func bedrock_CreateAutomatedReasoningPolicyTestCase(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.CreateAutomatedReasoningPolicyTestCaseInput{
		// ExpectedAggregatedFindingsResult: types.AutomatedReasoningCheckResult, // Required
		// GuardContent: *string, // Required
		// PolicyArn: *string, // Required
	}

	if len(_bedrockExpectedAggregatedFindingsResult) > 0 {
		if err := assignInputField(input, "ExpectedAggregatedFindingsResult", _bedrockExpectedAggregatedFindingsResult); err != nil {
			log.Errorf("invalid --expected-aggregated-findings-result: %s", err.Error())
			return
		}
	}
	if len(_bedrockGuardContent) > 0 {
		input.GuardContent = aws.String(_bedrockGuardContent)
	}
	if len(_bedrockPolicyArn) > 0 {
		input.PolicyArn = aws.String(_bedrockPolicyArn)
	}
	if len(_bedrockClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_bedrockClientRequestToken)
	}
	if len(_bedrockConfidenceThreshold) > 0 {
		if err := assignInputField(input, "ConfidenceThreshold", _bedrockConfidenceThreshold); err != nil {
			log.Errorf("invalid --confidence-threshold: %s", err.Error())
			return
		}
	}
	if len(_bedrockQueryContent) > 0 {
		input.QueryContent = aws.String(_bedrockQueryContent)
	}

	if resp, err := client.CreateAutomatedReasoningPolicyTestCase(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new version of an existing Automated Reasoning policy. This allows
// you to iterate on your policy rules while maintaining previous versions for
// rollback or comparison purposes.
func bedrock_CreateAutomatedReasoningPolicyVersion(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.CreateAutomatedReasoningPolicyVersionInput{
		// LastUpdatedDefinitionHash: *string, // Required
		// PolicyArn: *string, // Required
	}

	if len(_bedrockLastUpdatedDefinitionHash) > 0 {
		input.LastUpdatedDefinitionHash = aws.String(_bedrockLastUpdatedDefinitionHash)
	}
	if len(_bedrockPolicyArn) > 0 {
		input.PolicyArn = aws.String(_bedrockPolicyArn)
	}
	if len(_bedrockClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_bedrockClientRequestToken)
	}
	if len(_bedrockTags) > 0 {
		if err := assignInputField(input, "Tags", _bedrockTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAutomatedReasoningPolicyVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new custom model in Amazon Bedrock. After the model is active, you
// can use it for inference.
//
// To use the model for inference, you must purchase Provisioned Throughput for
// it. You can't use On-demand inference with these custom models. For more
// information about Provisioned Throughput, see [Provisioned Throughput].
//
// The model appears in ListCustomModels with a customizationType of imported . To
// track the status of the new model, you use the GetCustomModel API operation.
// The model can be in the following states:
//
// - Creating - Initial state during validation and registration
//
// - Active - Model is ready for use in inference
//
// - Failed - Creation process encountered an error
//
// # Related APIs
//
// [GetCustomModel]
//
// [ListCustomModels]
//
// [DeleteCustomModel]
//
// [Provisioned Throughput]: https://docs.aws.amazon.com/bedrock/latest/userguide/prov-throughput.html
// [ListCustomModels]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_ListCustomModels.html
// [DeleteCustomModel]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_DeleteCustomModel.html
// [GetCustomModel]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_GetCustomModel.html
func bedrock_CreateCustomModel(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.CreateCustomModelInput{
		// ModelName: *string, // Required
		// ModelSourceConfig: types.ModelDataSource, // Required
	}

	if len(_bedrockModelName) > 0 {
		input.ModelName = aws.String(_bedrockModelName)
	}
	if len(_bedrockModelSourceConfig) > 0 {
		if err := assignInputField(input, "ModelSourceConfig", _bedrockModelSourceConfig); err != nil {
			log.Errorf("invalid --model-source-config: %s", err.Error())
			return
		}
	}
	if len(_bedrockClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_bedrockClientRequestToken)
	}
	if len(_bedrockModelKmsKeyArn) > 0 {
		input.ModelKmsKeyArn = aws.String(_bedrockModelKmsKeyArn)
	}
	if len(_bedrockModelTags) > 0 {
		if err := assignInputField(input, "ModelTags", _bedrockModelTags); err != nil {
			log.Errorf("invalid --model-tags: %s", err.Error())
			return
		}
	}
	if len(_bedrockRoleArn) > 0 {
		input.RoleArn = aws.String(_bedrockRoleArn)
	}

	if resp, err := client.CreateCustomModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deploys a custom model for on-demand inference in Amazon Bedrock. After you
// deploy your custom model, you use the deployment's Amazon Resource Name (ARN) as
// the modelId parameter when you submit prompts and generate responses with model
// inference.
//
// For more information about setting up on-demand inference for custom models,
// see [Set up inference for a custom model].
//
// The following actions are related to the CreateCustomModelDeployment operation:
//
// [GetCustomModelDeployment]
//
// [ListCustomModelDeployments]
//
// [DeleteCustomModelDeployment]
//
// [ListCustomModelDeployments]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_ListCustomModelDeployments.html
// [GetCustomModelDeployment]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_GetCustomModelDeployment.html
// [Set up inference for a custom model]: https://docs.aws.amazon.com/bedrock/latest/userguide/model-customization-use.html
// [DeleteCustomModelDeployment]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_DeleteCustomModelDeployment.html
func bedrock_CreateCustomModelDeployment(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.CreateCustomModelDeploymentInput{
		// ModelArn: *string, // Required
		// ModelDeploymentName: *string, // Required
	}

	if len(_bedrockModelArn) > 0 {
		input.ModelArn = aws.String(_bedrockModelArn)
	}
	if len(_bedrockModelDeploymentName) > 0 {
		input.ModelDeploymentName = aws.String(_bedrockModelDeploymentName)
	}
	if len(_bedrockClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_bedrockClientRequestToken)
	}
	if len(_bedrockDescription) > 0 {
		input.Description = aws.String(_bedrockDescription)
	}
	if len(_bedrockTags) > 0 {
		if err := assignInputField(input, "Tags", _bedrockTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCustomModelDeployment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an evaluation job.
func bedrock_CreateEvaluationJob(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.CreateEvaluationJobInput{
		// EvaluationConfig: types.EvaluationConfig, // Required
		// InferenceConfig: types.EvaluationInferenceConfig, // Required
		// JobName: *string, // Required
		// OutputDataConfig: *types.EvaluationOutputDataConfig, // Required
		// RoleArn: *string, // Required
	}

	if len(_bedrockEvaluationConfig) > 0 {
		if err := assignInputField(input, "EvaluationConfig", _bedrockEvaluationConfig); err != nil {
			log.Errorf("invalid --evaluation-config: %s", err.Error())
			return
		}
	}
	if len(_bedrockInferenceConfig) > 0 {
		if err := assignInputField(input, "InferenceConfig", _bedrockInferenceConfig); err != nil {
			log.Errorf("invalid --inference-config: %s", err.Error())
			return
		}
	}
	if len(_bedrockJobName) > 0 {
		input.JobName = aws.String(_bedrockJobName)
	}
	if len(_bedrockOutputDataConfig) > 0 {
		if err := assignInputField(input, "OutputDataConfig", _bedrockOutputDataConfig); err != nil {
			log.Errorf("invalid --output-data-config: %s", err.Error())
			return
		}
	}
	if len(_bedrockRoleArn) > 0 {
		input.RoleArn = aws.String(_bedrockRoleArn)
	}
	if len(_bedrockApplicationType) > 0 {
		if err := assignInputField(input, "ApplicationType", _bedrockApplicationType); err != nil {
			log.Errorf("invalid --application-type: %s", err.Error())
			return
		}
	}
	if len(_bedrockClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_bedrockClientRequestToken)
	}
	if len(_bedrockCustomerEncryptionKeyId) > 0 {
		input.CustomerEncryptionKeyId = aws.String(_bedrockCustomerEncryptionKeyId)
	}
	if len(_bedrockJobDescription) > 0 {
		input.JobDescription = aws.String(_bedrockJobDescription)
	}
	if len(_bedrockJobTags) > 0 {
		if err := assignInputField(input, "JobTags", _bedrockJobTags); err != nil {
			log.Errorf("invalid --job-tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateEvaluationJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Request a model access agreement for the specified model.
func bedrock_CreateFoundationModelAgreement(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.CreateFoundationModelAgreementInput{
		// ModelId: *string, // Required
		// OfferToken: *string, // Required
	}

	if len(_bedrockModelId) > 0 {
		input.ModelId = aws.String(_bedrockModelId)
	}
	if len(_bedrockOfferToken) > 0 {
		input.OfferToken = aws.String(_bedrockOfferToken)
	}

	if resp, err := client.CreateFoundationModelAgreement(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a guardrail to block topics and to implement safeguards for your
// generative AI applications.
//
// You can configure the following policies in a guardrail to avoid undesirable
// and harmful content, filter out denied topics and words, and remove sensitive
// information for privacy protection.
//
// - Content filters - Adjust filter strengths to block input prompts or model
// responses containing harmful content.
//
// - Denied topics - Define a set of topics that are undesirable in the context
// of your application. These topics will be blocked if detected in user queries or
// model responses.
//
// - Word filters - Configure filters to block undesirable words, phrases, and
// profanity. Such words can include offensive terms, competitor names etc.
//
// - Sensitive information filters - Block or mask sensitive information such as
// personally identifiable information (PII) or custom regex in user inputs and
// model responses.
//
// In addition to the above policies, you can also configure the messages to be
// returned to the user if a user input or model response is in violation of the
// policies defined in the guardrail.
//
// For more information, see [Amazon Bedrock Guardrails] in the Amazon Bedrock User Guide.
//
// [Amazon Bedrock Guardrails]: https://docs.aws.amazon.com/bedrock/latest/userguide/guardrails.html
func bedrock_CreateGuardrail(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.CreateGuardrailInput{
		// BlockedInputMessaging: *string, // Required
		// BlockedOutputsMessaging: *string, // Required
		// Name: *string, // Required
	}

	if len(_bedrockBlockedInputMessaging) > 0 {
		input.BlockedInputMessaging = aws.String(_bedrockBlockedInputMessaging)
	}
	if len(_bedrockBlockedOutputsMessaging) > 0 {
		input.BlockedOutputsMessaging = aws.String(_bedrockBlockedOutputsMessaging)
	}
	if len(_bedrockName) > 0 {
		input.Name = aws.String(_bedrockName)
	}
	if len(_bedrockAutomatedReasoningPolicyConfig) > 0 {
		if err := assignInputField(input, "AutomatedReasoningPolicyConfig", _bedrockAutomatedReasoningPolicyConfig); err != nil {
			log.Errorf("invalid --automated-reasoning-policy-config: %s", err.Error())
			return
		}
	}
	if len(_bedrockClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_bedrockClientRequestToken)
	}
	if len(_bedrockContentPolicyConfig) > 0 {
		if err := assignInputField(input, "ContentPolicyConfig", _bedrockContentPolicyConfig); err != nil {
			log.Errorf("invalid --content-policy-config: %s", err.Error())
			return
		}
	}
	if len(_bedrockContextualGroundingPolicyConfig) > 0 {
		if err := assignInputField(input, "ContextualGroundingPolicyConfig", _bedrockContextualGroundingPolicyConfig); err != nil {
			log.Errorf("invalid --contextual-grounding-policy-config: %s", err.Error())
			return
		}
	}
	if len(_bedrockCrossRegionConfig) > 0 {
		if err := assignInputField(input, "CrossRegionConfig", _bedrockCrossRegionConfig); err != nil {
			log.Errorf("invalid --cross-region-config: %s", err.Error())
			return
		}
	}
	if len(_bedrockDescription) > 0 {
		input.Description = aws.String(_bedrockDescription)
	}
	if len(_bedrockKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_bedrockKmsKeyId)
	}
	if len(_bedrockSensitiveInformationPolicyConfig) > 0 {
		if err := assignInputField(input, "SensitiveInformationPolicyConfig", _bedrockSensitiveInformationPolicyConfig); err != nil {
			log.Errorf("invalid --sensitive-information-policy-config: %s", err.Error())
			return
		}
	}
	if len(_bedrockTags) > 0 {
		if err := assignInputField(input, "Tags", _bedrockTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_bedrockTopicPolicyConfig) > 0 {
		if err := assignInputField(input, "TopicPolicyConfig", _bedrockTopicPolicyConfig); err != nil {
			log.Errorf("invalid --topic-policy-config: %s", err.Error())
			return
		}
	}
	if len(_bedrockWordPolicyConfig) > 0 {
		if err := assignInputField(input, "WordPolicyConfig", _bedrockWordPolicyConfig); err != nil {
			log.Errorf("invalid --word-policy-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateGuardrail(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a version of the guardrail. Use this API to create a snapshot of the
// guardrail when you are satisfied with a configuration, or to compare the
// configuration with another version.
func bedrock_CreateGuardrailVersion(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.CreateGuardrailVersionInput{
		// GuardrailIdentifier: *string, // Required
	}

	if len(_bedrockGuardrailIdentifier) > 0 {
		input.GuardrailIdentifier = aws.String(_bedrockGuardrailIdentifier)
	}
	if len(_bedrockClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_bedrockClientRequestToken)
	}
	if len(_bedrockDescription) > 0 {
		input.Description = aws.String(_bedrockDescription)
	}

	if resp, err := client.CreateGuardrailVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an application inference profile to track metrics and costs when
// invoking a model. To create an application inference profile for a foundation
// model in one region, specify the ARN of the model in that region. To create an
// application inference profile for a foundation model across multiple regions,
// specify the ARN of the system-defined inference profile that contains the
// regions that you want to route requests to. For more information, see [Increase throughput and resilience with cross-region inference in Amazon Bedrock]. in the
// Amazon Bedrock User Guide.
//
// [Increase throughput and resilience with cross-region inference in Amazon Bedrock]: https://docs.aws.amazon.com/bedrock/latest/userguide/cross-region-inference.html
func bedrock_CreateInferenceProfile(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.CreateInferenceProfileInput{
		// InferenceProfileName: *string, // Required
		// ModelSource: types.InferenceProfileModelSource, // Required
	}

	if len(_bedrockInferenceProfileName) > 0 {
		input.InferenceProfileName = aws.String(_bedrockInferenceProfileName)
	}
	if len(_bedrockModelSource) > 0 {
		if err := assignInputField(input, "ModelSource", _bedrockModelSource); err != nil {
			log.Errorf("invalid --model-source: %s", err.Error())
			return
		}
	}
	if len(_bedrockClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_bedrockClientRequestToken)
	}
	if len(_bedrockDescription) > 0 {
		input.Description = aws.String(_bedrockDescription)
	}
	if len(_bedrockTags) > 0 {
		if err := assignInputField(input, "Tags", _bedrockTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateInferenceProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an endpoint for a model from Amazon Bedrock Marketplace. The endpoint
// is hosted by Amazon SageMaker.
func bedrock_CreateMarketplaceModelEndpoint(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.CreateMarketplaceModelEndpointInput{
		// EndpointConfig: types.EndpointConfig, // Required
		// EndpointName: *string, // Required
		// ModelSourceIdentifier: *string, // Required
	}

	if len(_bedrockEndpointConfig) > 0 {
		if err := assignInputField(input, "EndpointConfig", _bedrockEndpointConfig); err != nil {
			log.Errorf("invalid --endpoint-config: %s", err.Error())
			return
		}
	}
	if len(_bedrockEndpointName) > 0 {
		input.EndpointName = aws.String(_bedrockEndpointName)
	}
	if len(_bedrockModelSourceIdentifier) > 0 {
		input.ModelSourceIdentifier = aws.String(_bedrockModelSourceIdentifier)
	}
	if len(_bedrockAcceptEula) > 0 {
		if err := assignInputField(input, "AcceptEula", _bedrockAcceptEula); err != nil {
			log.Errorf("invalid --accept-eula: %s", err.Error())
			return
		}
	}
	if len(_bedrockClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_bedrockClientRequestToken)
	}
	if len(_bedrockTags) > 0 {
		if err := assignInputField(input, "Tags", _bedrockTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateMarketplaceModelEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Copies a model to another region so that it can be used there. For more
// information, see [Copy models to be used in other regions]in the [Amazon Bedrock User Guide].
//
// [Copy models to be used in other regions]: https://docs.aws.amazon.com/bedrock/latest/userguide/copy-model.html
// [Amazon Bedrock User Guide]: https://docs.aws.amazon.com/bedrock/latest/userguide/what-is-service.html
func bedrock_CreateModelCopyJob(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.CreateModelCopyJobInput{
		// SourceModelArn: *string, // Required
		// TargetModelName: *string, // Required
	}

	if len(_bedrockSourceModelArn) > 0 {
		input.SourceModelArn = aws.String(_bedrockSourceModelArn)
	}
	if len(_bedrockTargetModelName) > 0 {
		input.TargetModelName = aws.String(_bedrockTargetModelName)
	}
	if len(_bedrockClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_bedrockClientRequestToken)
	}
	if len(_bedrockModelKmsKeyId) > 0 {
		input.ModelKmsKeyId = aws.String(_bedrockModelKmsKeyId)
	}
	if len(_bedrockTargetModelTags) > 0 {
		if err := assignInputField(input, "TargetModelTags", _bedrockTargetModelTags); err != nil {
			log.Errorf("invalid --target-model-tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateModelCopyJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a fine-tuning job to customize a base model.
// You specify the base foundation model and the location of the training data.
// After the model-customization job completes successfully, your custom model
// resource will be ready to use. Amazon Bedrock returns validation loss metrics
// and output generations after the job completes.
//
// For information on the format of training and validation data, see [Prepare the datasets].
//
// Model-customization jobs are asynchronous and the completion time depends on
// the base model and the training/validation data size. To monitor a job, use the
// GetModelCustomizationJob operation to retrieve the job status.
//
// For more information, see [Custom models] in the [Amazon Bedrock User Guide].
//
// [Custom models]: https://docs.aws.amazon.com/bedrock/latest/userguide/custom-models.html
// [Amazon Bedrock User Guide]: https://docs.aws.amazon.com/bedrock/latest/userguide/what-is-service.html
// [Prepare the datasets]: https://docs.aws.amazon.com/bedrock/latest/userguide/model-customization-prepare.html
func bedrock_CreateModelCustomizationJob(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.CreateModelCustomizationJobInput{
		// BaseModelIdentifier: *string, // Required
		// CustomModelName: *string, // Required
		// JobName: *string, // Required
		// OutputDataConfig: *types.OutputDataConfig, // Required
		// RoleArn: *string, // Required
		// TrainingDataConfig: *types.TrainingDataConfig, // Required
	}

	if len(_bedrockBaseModelIdentifier) > 0 {
		input.BaseModelIdentifier = aws.String(_bedrockBaseModelIdentifier)
	}
	if len(_bedrockCustomModelName) > 0 {
		input.CustomModelName = aws.String(_bedrockCustomModelName)
	}
	if len(_bedrockJobName) > 0 {
		input.JobName = aws.String(_bedrockJobName)
	}
	if len(_bedrockOutputDataConfig) > 0 {
		if err := assignInputField(input, "OutputDataConfig", _bedrockOutputDataConfig); err != nil {
			log.Errorf("invalid --output-data-config: %s", err.Error())
			return
		}
	}
	if len(_bedrockRoleArn) > 0 {
		input.RoleArn = aws.String(_bedrockRoleArn)
	}
	if len(_bedrockTrainingDataConfig) > 0 {
		if err := assignInputField(input, "TrainingDataConfig", _bedrockTrainingDataConfig); err != nil {
			log.Errorf("invalid --training-data-config: %s", err.Error())
			return
		}
	}
	if len(_bedrockClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_bedrockClientRequestToken)
	}
	if len(_bedrockCustomModelKmsKeyId) > 0 {
		input.CustomModelKmsKeyId = aws.String(_bedrockCustomModelKmsKeyId)
	}
	if len(_bedrockCustomModelTags) > 0 {
		if err := assignInputField(input, "CustomModelTags", _bedrockCustomModelTags); err != nil {
			log.Errorf("invalid --custom-model-tags: %s", err.Error())
			return
		}
	}
	if len(_bedrockCustomizationConfig) > 0 {
		if err := assignInputField(input, "CustomizationConfig", _bedrockCustomizationConfig); err != nil {
			log.Errorf("invalid --customization-config: %s", err.Error())
			return
		}
	}
	if len(_bedrockCustomizationType) > 0 {
		if err := assignInputField(input, "CustomizationType", _bedrockCustomizationType); err != nil {
			log.Errorf("invalid --customization-type: %s", err.Error())
			return
		}
	}
	if len(_bedrockHyperParameters) > 0 {
		if err := assignInputField(input, "HyperParameters", _bedrockHyperParameters); err != nil {
			log.Errorf("invalid --hyper-parameters: %s", err.Error())
			return
		}
	}
	if len(_bedrockJobTags) > 0 {
		if err := assignInputField(input, "JobTags", _bedrockJobTags); err != nil {
			log.Errorf("invalid --job-tags: %s", err.Error())
			return
		}
	}
	if len(_bedrockValidationDataConfig) > 0 {
		if err := assignInputField(input, "ValidationDataConfig", _bedrockValidationDataConfig); err != nil {
			log.Errorf("invalid --validation-data-config: %s", err.Error())
			return
		}
	}
	if len(_bedrockVpcConfig) > 0 {
		if err := assignInputField(input, "VpcConfig", _bedrockVpcConfig); err != nil {
			log.Errorf("invalid --vpc-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateModelCustomizationJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a model import job to import model that you have customized in other
// environments, such as Amazon SageMaker. For more information, see [Import a customized model]
//
// [Import a customized model]: https://docs.aws.amazon.com/bedrock/latest/userguide/model-customization-import-model.html
func bedrock_CreateModelImportJob(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.CreateModelImportJobInput{
		// ImportedModelName: *string, // Required
		// JobName: *string, // Required
		// ModelDataSource: types.ModelDataSource, // Required
		// RoleArn: *string, // Required
	}

	if len(_bedrockImportedModelName) > 0 {
		input.ImportedModelName = aws.String(_bedrockImportedModelName)
	}
	if len(_bedrockJobName) > 0 {
		input.JobName = aws.String(_bedrockJobName)
	}
	if len(_bedrockModelDataSource) > 0 {
		if err := assignInputField(input, "ModelDataSource", _bedrockModelDataSource); err != nil {
			log.Errorf("invalid --model-data-source: %s", err.Error())
			return
		}
	}
	if len(_bedrockRoleArn) > 0 {
		input.RoleArn = aws.String(_bedrockRoleArn)
	}
	if len(_bedrockClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_bedrockClientRequestToken)
	}
	if len(_bedrockImportedModelKmsKeyId) > 0 {
		input.ImportedModelKmsKeyId = aws.String(_bedrockImportedModelKmsKeyId)
	}
	if len(_bedrockImportedModelTags) > 0 {
		if err := assignInputField(input, "ImportedModelTags", _bedrockImportedModelTags); err != nil {
			log.Errorf("invalid --imported-model-tags: %s", err.Error())
			return
		}
	}
	if len(_bedrockJobTags) > 0 {
		if err := assignInputField(input, "JobTags", _bedrockJobTags); err != nil {
			log.Errorf("invalid --job-tags: %s", err.Error())
			return
		}
	}
	if len(_bedrockVpcConfig) > 0 {
		if err := assignInputField(input, "VpcConfig", _bedrockVpcConfig); err != nil {
			log.Errorf("invalid --vpc-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateModelImportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a batch inference job to invoke a model on multiple prompts. Format
// your data according to [Format your inference data]and upload it to an Amazon S3 bucket. For more
// information, see [Process multiple prompts with batch inference].
//
// The response returns a jobArn that you can use to stop or get details about the
// job.
//
// [Process multiple prompts with batch inference]: https://docs.aws.amazon.com/bedrock/latest/userguide/batch-inference.html
// [Format your inference data]: https://docs.aws.amazon.com/bedrock/latest/userguide/batch-inference-data
func bedrock_CreateModelInvocationJob(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.CreateModelInvocationJobInput{
		// InputDataConfig: types.ModelInvocationJobInputDataConfig, // Required
		// JobName: *string, // Required
		// ModelId: *string, // Required
		// OutputDataConfig: types.ModelInvocationJobOutputDataConfig, // Required
		// RoleArn: *string, // Required
	}

	if len(_bedrockInputDataConfig) > 0 {
		if err := assignInputField(input, "InputDataConfig", _bedrockInputDataConfig); err != nil {
			log.Errorf("invalid --input-data-config: %s", err.Error())
			return
		}
	}
	if len(_bedrockJobName) > 0 {
		input.JobName = aws.String(_bedrockJobName)
	}
	if len(_bedrockModelId) > 0 {
		input.ModelId = aws.String(_bedrockModelId)
	}
	if len(_bedrockOutputDataConfig) > 0 {
		if err := assignInputField(input, "OutputDataConfig", _bedrockOutputDataConfig); err != nil {
			log.Errorf("invalid --output-data-config: %s", err.Error())
			return
		}
	}
	if len(_bedrockRoleArn) > 0 {
		input.RoleArn = aws.String(_bedrockRoleArn)
	}
	if len(_bedrockClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_bedrockClientRequestToken)
	}
	if len(_bedrockModelInvocationType) > 0 {
		if err := assignInputField(input, "ModelInvocationType", _bedrockModelInvocationType); err != nil {
			log.Errorf("invalid --model-invocation-type: %s", err.Error())
			return
		}
	}
	if len(_bedrockTags) > 0 {
		if err := assignInputField(input, "Tags", _bedrockTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_bedrockTimeoutDurationInHours) > 0 {
		if err := assignInputField(input, "TimeoutDurationInHours", _bedrockTimeoutDurationInHours); err != nil {
			log.Errorf("invalid --timeout-duration-in-hours: %s", err.Error())
			return
		}
	}
	if len(_bedrockVpcConfig) > 0 {
		if err := assignInputField(input, "VpcConfig", _bedrockVpcConfig); err != nil {
			log.Errorf("invalid --vpc-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateModelInvocationJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a prompt router that manages the routing of requests between multiple
// foundation models based on the routing criteria.
func bedrock_CreatePromptRouter(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.CreatePromptRouterInput{
		// FallbackModel: *types.PromptRouterTargetModel, // Required
		// Models: []types.PromptRouterTargetModel, // Required
		// PromptRouterName: *string, // Required
		// RoutingCriteria: *types.RoutingCriteria, // Required
	}

	if len(_bedrockFallbackModel) > 0 {
		if err := assignInputField(input, "FallbackModel", _bedrockFallbackModel); err != nil {
			log.Errorf("invalid --fallback-model: %s", err.Error())
			return
		}
	}
	if len(_bedrockModels) > 0 {
		if err := assignInputField(input, "Models", _bedrockModels); err != nil {
			log.Errorf("invalid --models: %s", err.Error())
			return
		}
	}
	if len(_bedrockPromptRouterName) > 0 {
		input.PromptRouterName = aws.String(_bedrockPromptRouterName)
	}
	if len(_bedrockRoutingCriteria) > 0 {
		if err := assignInputField(input, "RoutingCriteria", _bedrockRoutingCriteria); err != nil {
			log.Errorf("invalid --routing-criteria: %s", err.Error())
			return
		}
	}
	if len(_bedrockClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_bedrockClientRequestToken)
	}
	if len(_bedrockDescription) > 0 {
		input.Description = aws.String(_bedrockDescription)
	}
	if len(_bedrockTags) > 0 {
		if err := assignInputField(input, "Tags", _bedrockTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreatePromptRouter(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates dedicated throughput for a base or custom model with the model units
// and for the duration that you specify. For pricing details, see [Amazon Bedrock Pricing]. For more
// information, see [Provisioned Throughput]in the [Amazon Bedrock User Guide].
//
// [Amazon Bedrock Pricing]: http://aws.amazon.com/bedrock/pricing/
// [Provisioned Throughput]: https://docs.aws.amazon.com/bedrock/latest/userguide/prov-throughput.html
// [Amazon Bedrock User Guide]: https://docs.aws.amazon.com/bedrock/latest/userguide/what-is-service.html
func bedrock_CreateProvisionedModelThroughput(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.CreateProvisionedModelThroughputInput{
		// ModelId: *string, // Required
		// ModelUnits: *int32, // Required
		// ProvisionedModelName: *string, // Required
	}

	if len(_bedrockModelId) > 0 {
		input.ModelId = aws.String(_bedrockModelId)
	}
	if len(_bedrockModelUnits) > 0 {
		if err := assignInputField(input, "ModelUnits", _bedrockModelUnits); err != nil {
			log.Errorf("invalid --model-units: %s", err.Error())
			return
		}
	}
	if len(_bedrockProvisionedModelName) > 0 {
		input.ProvisionedModelName = aws.String(_bedrockProvisionedModelName)
	}
	if len(_bedrockClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_bedrockClientRequestToken)
	}
	if len(_bedrockCommitmentDuration) > 0 {
		if err := assignInputField(input, "CommitmentDuration", _bedrockCommitmentDuration); err != nil {
			log.Errorf("invalid --commitment-duration: %s", err.Error())
			return
		}
	}
	if len(_bedrockTags) > 0 {
		if err := assignInputField(input, "Tags", _bedrockTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateProvisionedModelThroughput(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Automated Reasoning policy or policy version. This operation is
// idempotent. If you delete a policy more than once, each call succeeds. Deleting
// a policy removes it permanently and cannot be undone.
func bedrock_DeleteAutomatedReasoningPolicy(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.DeleteAutomatedReasoningPolicyInput{
		// PolicyArn: *string, // Required
	}

	if len(_bedrockPolicyArn) > 0 {
		input.PolicyArn = aws.String(_bedrockPolicyArn)
	}
	if len(_bedrockForce) > 0 {
		if err := assignInputField(input, "Force", _bedrockForce); err != nil {
			log.Errorf("invalid --force: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteAutomatedReasoningPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Automated Reasoning policy build workflow and its associated
// artifacts. This permanently removes the workflow history and any generated
// assets.
func bedrock_DeleteAutomatedReasoningPolicyBuildWorkflow(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.DeleteAutomatedReasoningPolicyBuildWorkflowInput{
		// BuildWorkflowId: *string, // Required
		// LastUpdatedAt: *time.Time, // Required
		// PolicyArn: *string, // Required
	}

	if len(_bedrockBuildWorkflowId) > 0 {
		input.BuildWorkflowId = aws.String(_bedrockBuildWorkflowId)
	}
	if len(_bedrockLastUpdatedAt) > 0 {
		if err := assignInputField(input, "LastUpdatedAt", _bedrockLastUpdatedAt); err != nil {
			log.Errorf("invalid --last-updated-at: %s", err.Error())
			return
		}
	}
	if len(_bedrockPolicyArn) > 0 {
		input.PolicyArn = aws.String(_bedrockPolicyArn)
	}

	if resp, err := client.DeleteAutomatedReasoningPolicyBuildWorkflow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Automated Reasoning policy test. This operation is idempotent; if
// you delete a test more than once, each call succeeds.
func bedrock_DeleteAutomatedReasoningPolicyTestCase(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.DeleteAutomatedReasoningPolicyTestCaseInput{
		// LastUpdatedAt: *time.Time, // Required
		// PolicyArn: *string, // Required
		// TestCaseId: *string, // Required
	}

	if len(_bedrockLastUpdatedAt) > 0 {
		if err := assignInputField(input, "LastUpdatedAt", _bedrockLastUpdatedAt); err != nil {
			log.Errorf("invalid --last-updated-at: %s", err.Error())
			return
		}
	}
	if len(_bedrockPolicyArn) > 0 {
		input.PolicyArn = aws.String(_bedrockPolicyArn)
	}
	if len(_bedrockTestCaseId) > 0 {
		input.TestCaseId = aws.String(_bedrockTestCaseId)
	}

	if resp, err := client.DeleteAutomatedReasoningPolicyTestCase(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a custom model that you created earlier. For more information, see [Custom models] in
// the [Amazon Bedrock User Guide].
//
// [Amazon Bedrock User Guide]: https://docs.aws.amazon.com/bedrock/latest/userguide/what-is-service.html
// [Custom models]: https://docs.aws.amazon.com/bedrock/latest/userguide/custom-models.html
func bedrock_DeleteCustomModel(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.DeleteCustomModelInput{
		// ModelIdentifier: *string, // Required
	}

	if len(_bedrockModelIdentifier) > 0 {
		input.ModelIdentifier = aws.String(_bedrockModelIdentifier)
	}

	if resp, err := client.DeleteCustomModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a custom model deployment. This operation stops the deployment and
// removes it from your account. After deletion, the deployment ARN can no longer
// be used for inference requests.
//
// The following actions are related to the DeleteCustomModelDeployment operation:
//
// [CreateCustomModelDeployment]
//
// [GetCustomModelDeployment]
//
// [ListCustomModelDeployments]
//
// [ListCustomModelDeployments]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_ListCustomModelDeployments.html
// [GetCustomModelDeployment]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_GetCustomModelDeployment.html
// [CreateCustomModelDeployment]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_CreateCustomModelDeployment.html
func bedrock_DeleteCustomModelDeployment(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.DeleteCustomModelDeploymentInput{
		// CustomModelDeploymentIdentifier: *string, // Required
	}

	if len(_bedrockCustomModelDeploymentIdentifier) > 0 {
		input.CustomModelDeploymentIdentifier = aws.String(_bedrockCustomModelDeploymentIdentifier)
	}

	if resp, err := client.DeleteCustomModelDeployment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the account-level enforced guardrail configuration.
func bedrock_DeleteEnforcedGuardrailConfiguration(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.DeleteEnforcedGuardrailConfigurationInput{
		// ConfigId: *string, // Required
	}

	if len(_bedrockConfigId) > 0 {
		input.ConfigId = aws.String(_bedrockConfigId)
	}

	if resp, err := client.DeleteEnforcedGuardrailConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete the model access agreement for the specified model.
func bedrock_DeleteFoundationModelAgreement(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.DeleteFoundationModelAgreementInput{
		// ModelId: *string, // Required
	}

	if len(_bedrockModelId) > 0 {
		input.ModelId = aws.String(_bedrockModelId)
	}

	if resp, err := client.DeleteFoundationModelAgreement(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a guardrail.
// - To delete a guardrail, only specify the ARN of the guardrail in the
// guardrailIdentifier field. If you delete a guardrail, all of its versions will
// be deleted.
//
// - To delete a version of a guardrail, specify the ARN of the guardrail in the
// guardrailIdentifier field and the version in the guardrailVersion field.
func bedrock_DeleteGuardrail(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.DeleteGuardrailInput{
		// GuardrailIdentifier: *string, // Required
	}

	if len(_bedrockGuardrailIdentifier) > 0 {
		input.GuardrailIdentifier = aws.String(_bedrockGuardrailIdentifier)
	}
	if len(_bedrockGuardrailVersion) > 0 {
		input.GuardrailVersion = aws.String(_bedrockGuardrailVersion)
	}

	if resp, err := client.DeleteGuardrail(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a custom model that you imported earlier. For more information, see [Import a customized model] in
// the [Amazon Bedrock User Guide].
//
// [Amazon Bedrock User Guide]: https://docs.aws.amazon.com/bedrock/latest/userguide/what-is-service.html
// [Import a customized model]: https://docs.aws.amazon.com/bedrock/latest/userguide/model-customization-import-model.html
func bedrock_DeleteImportedModel(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.DeleteImportedModelInput{
		// ModelIdentifier: *string, // Required
	}

	if len(_bedrockModelIdentifier) > 0 {
		input.ModelIdentifier = aws.String(_bedrockModelIdentifier)
	}

	if resp, err := client.DeleteImportedModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an application inference profile. For more information, see [Increase throughput and resilience with cross-region inference in Amazon Bedrock]. in the
// Amazon Bedrock User Guide.
//
// [Increase throughput and resilience with cross-region inference in Amazon Bedrock]: https://docs.aws.amazon.com/bedrock/latest/userguide/cross-region-inference.html
func bedrock_DeleteInferenceProfile(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.DeleteInferenceProfileInput{
		// InferenceProfileIdentifier: *string, // Required
	}

	if len(_bedrockInferenceProfileIdentifier) > 0 {
		input.InferenceProfileIdentifier = aws.String(_bedrockInferenceProfileIdentifier)
	}

	if resp, err := client.DeleteInferenceProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an endpoint for a model from Amazon Bedrock Marketplace.
func bedrock_DeleteMarketplaceModelEndpoint(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.DeleteMarketplaceModelEndpointInput{
		// EndpointArn: *string, // Required
	}

	if len(_bedrockEndpointArn) > 0 {
		input.EndpointArn = aws.String(_bedrockEndpointArn)
	}

	if resp, err := client.DeleteMarketplaceModelEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete the invocation logging.
func bedrock_DeleteModelInvocationLoggingConfiguration(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.DeleteModelInvocationLoggingConfigurationInput{}

	if resp, err := client.DeleteModelInvocationLoggingConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a specified prompt router. This action cannot be undone.
func bedrock_DeletePromptRouter(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.DeletePromptRouterInput{
		// PromptRouterArn: *string, // Required
	}

	if len(_bedrockPromptRouterArn) > 0 {
		input.PromptRouterArn = aws.String(_bedrockPromptRouterArn)
	}

	if resp, err := client.DeletePromptRouter(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a Provisioned Throughput. You can't delete a Provisioned Throughput
// before the commitment term is over. For more information, see [Provisioned Throughput]in the [Amazon Bedrock User Guide].
//
// [Provisioned Throughput]: https://docs.aws.amazon.com/bedrock/latest/userguide/prov-throughput.html
// [Amazon Bedrock User Guide]: https://docs.aws.amazon.com/bedrock/latest/userguide/what-is-service.html
func bedrock_DeleteProvisionedModelThroughput(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.DeleteProvisionedModelThroughputInput{
		// ProvisionedModelId: *string, // Required
	}

	if len(_bedrockProvisionedModelId) > 0 {
		input.ProvisionedModelId = aws.String(_bedrockProvisionedModelId)
	}

	if resp, err := client.DeleteProvisionedModelThroughput(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deregisters an endpoint for a model from Amazon Bedrock Marketplace. This
// operation removes the endpoint's association with Amazon Bedrock but does not
// delete the underlying Amazon SageMaker endpoint.
func bedrock_DeregisterMarketplaceModelEndpoint(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.DeregisterMarketplaceModelEndpointInput{
		// EndpointArn: *string, // Required
	}

	if len(_bedrockEndpointArn) > 0 {
		input.EndpointArn = aws.String(_bedrockEndpointArn)
	}

	if resp, err := client.DeregisterMarketplaceModelEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Exports the policy definition for an Automated Reasoning policy version.
// Returns the complete policy definition including rules, variables, and custom
// variable types in a structured format.
func bedrock_ExportAutomatedReasoningPolicyVersion(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.ExportAutomatedReasoningPolicyVersionInput{
		// PolicyArn: *string, // Required
	}

	if len(_bedrockPolicyArn) > 0 {
		input.PolicyArn = aws.String(_bedrockPolicyArn)
	}

	if resp, err := client.ExportAutomatedReasoningPolicyVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves details about an Automated Reasoning policy or policy version.
// Returns information including the policy definition, metadata, and timestamps.
func bedrock_GetAutomatedReasoningPolicy(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.GetAutomatedReasoningPolicyInput{
		// PolicyArn: *string, // Required
	}

	if len(_bedrockPolicyArn) > 0 {
		input.PolicyArn = aws.String(_bedrockPolicyArn)
	}

	if resp, err := client.GetAutomatedReasoningPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the current annotations for an Automated Reasoning policy build
// workflow. Annotations contain corrections to the rules, variables and types to
// be applied to the policy.
func bedrock_GetAutomatedReasoningPolicyAnnotations(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.GetAutomatedReasoningPolicyAnnotationsInput{
		// BuildWorkflowId: *string, // Required
		// PolicyArn: *string, // Required
	}

	if len(_bedrockBuildWorkflowId) > 0 {
		input.BuildWorkflowId = aws.String(_bedrockBuildWorkflowId)
	}
	if len(_bedrockPolicyArn) > 0 {
		input.PolicyArn = aws.String(_bedrockPolicyArn)
	}

	if resp, err := client.GetAutomatedReasoningPolicyAnnotations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves detailed information about an Automated Reasoning policy build
// workflow, including its status, configuration, and metadata.
func bedrock_GetAutomatedReasoningPolicyBuildWorkflow(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.GetAutomatedReasoningPolicyBuildWorkflowInput{
		// BuildWorkflowId: *string, // Required
		// PolicyArn: *string, // Required
	}

	if len(_bedrockBuildWorkflowId) > 0 {
		input.BuildWorkflowId = aws.String(_bedrockBuildWorkflowId)
	}
	if len(_bedrockPolicyArn) > 0 {
		input.PolicyArn = aws.String(_bedrockPolicyArn)
	}

	if resp, err := client.GetAutomatedReasoningPolicyBuildWorkflow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the resulting assets from a completed Automated Reasoning policy
// build workflow, including build logs, quality reports, and generated policy
// artifacts.
func bedrock_GetAutomatedReasoningPolicyBuildWorkflowResultAssets(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.GetAutomatedReasoningPolicyBuildWorkflowResultAssetsInput{
		// AssetType: types.AutomatedReasoningPolicyBuildResultAssetType, // Required
		// BuildWorkflowId: *string, // Required
		// PolicyArn: *string, // Required
	}

	if len(_bedrockAssetType) > 0 {
		if err := assignInputField(input, "AssetType", _bedrockAssetType); err != nil {
			log.Errorf("invalid --asset-type: %s", err.Error())
			return
		}
	}
	if len(_bedrockBuildWorkflowId) > 0 {
		input.BuildWorkflowId = aws.String(_bedrockBuildWorkflowId)
	}
	if len(_bedrockPolicyArn) > 0 {
		input.PolicyArn = aws.String(_bedrockPolicyArn)
	}
	if len(_bedrockAssetId) > 0 {
		input.AssetId = aws.String(_bedrockAssetId)
	}

	if resp, err := client.GetAutomatedReasoningPolicyBuildWorkflowResultAssets(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the next test scenario for validating an Automated Reasoning policy.
// This is used during the interactive policy refinement process to test policy
// behavior.
func bedrock_GetAutomatedReasoningPolicyNextScenario(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.GetAutomatedReasoningPolicyNextScenarioInput{
		// BuildWorkflowId: *string, // Required
		// PolicyArn: *string, // Required
	}

	if len(_bedrockBuildWorkflowId) > 0 {
		input.BuildWorkflowId = aws.String(_bedrockBuildWorkflowId)
	}
	if len(_bedrockPolicyArn) > 0 {
		input.PolicyArn = aws.String(_bedrockPolicyArn)
	}

	if resp, err := client.GetAutomatedReasoningPolicyNextScenario(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves details about a specific Automated Reasoning policy test.
func bedrock_GetAutomatedReasoningPolicyTestCase(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.GetAutomatedReasoningPolicyTestCaseInput{
		// PolicyArn: *string, // Required
		// TestCaseId: *string, // Required
	}

	if len(_bedrockPolicyArn) > 0 {
		input.PolicyArn = aws.String(_bedrockPolicyArn)
	}
	if len(_bedrockTestCaseId) > 0 {
		input.TestCaseId = aws.String(_bedrockTestCaseId)
	}

	if resp, err := client.GetAutomatedReasoningPolicyTestCase(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the test result for a specific Automated Reasoning policy test.
// Returns detailed validation findings and execution status.
func bedrock_GetAutomatedReasoningPolicyTestResult(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.GetAutomatedReasoningPolicyTestResultInput{
		// BuildWorkflowId: *string, // Required
		// PolicyArn: *string, // Required
		// TestCaseId: *string, // Required
	}

	if len(_bedrockBuildWorkflowId) > 0 {
		input.BuildWorkflowId = aws.String(_bedrockBuildWorkflowId)
	}
	if len(_bedrockPolicyArn) > 0 {
		input.PolicyArn = aws.String(_bedrockPolicyArn)
	}
	if len(_bedrockTestCaseId) > 0 {
		input.TestCaseId = aws.String(_bedrockTestCaseId)
	}

	if resp, err := client.GetAutomatedReasoningPolicyTestResult(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get the properties associated with a Amazon Bedrock custom model that you have
// created. For more information, see [Custom models]in the [Amazon Bedrock User Guide].
//
// [Custom models]: https://docs.aws.amazon.com/bedrock/latest/userguide/custom-models.html
// [Amazon Bedrock User Guide]: https://docs.aws.amazon.com/bedrock/latest/userguide/what-is-service.html
func bedrock_GetCustomModel(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.GetCustomModelInput{
		// ModelIdentifier: *string, // Required
	}

	if len(_bedrockModelIdentifier) > 0 {
		input.ModelIdentifier = aws.String(_bedrockModelIdentifier)
	}

	if resp, err := client.GetCustomModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a custom model deployment, including its status,
// configuration, and metadata. Use this operation to monitor the deployment status
// and retrieve details needed for inference requests.
//
// The following actions are related to the GetCustomModelDeployment operation:
//
// [CreateCustomModelDeployment]
//
// [ListCustomModelDeployments]
//
// [DeleteCustomModelDeployment]
//
// [ListCustomModelDeployments]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_ListCustomModelDeployments.html
// [CreateCustomModelDeployment]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_CreateCustomModelDeployment.html
// [DeleteCustomModelDeployment]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_DeleteCustomModelDeployment.html
func bedrock_GetCustomModelDeployment(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.GetCustomModelDeploymentInput{
		// CustomModelDeploymentIdentifier: *string, // Required
	}

	if len(_bedrockCustomModelDeploymentIdentifier) > 0 {
		input.CustomModelDeploymentIdentifier = aws.String(_bedrockCustomModelDeploymentIdentifier)
	}

	if resp, err := client.GetCustomModelDeployment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about an evaluation job, such as the status of the job.
func bedrock_GetEvaluationJob(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.GetEvaluationJobInput{
		// JobIdentifier: *string, // Required
	}

	if len(_bedrockJobIdentifier) > 0 {
		input.JobIdentifier = aws.String(_bedrockJobIdentifier)
	}

	if resp, err := client.GetEvaluationJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get details about a Amazon Bedrock foundation model.
func bedrock_GetFoundationModel(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.GetFoundationModelInput{
		// ModelIdentifier: *string, // Required
	}

	if len(_bedrockModelIdentifier) > 0 {
		input.ModelIdentifier = aws.String(_bedrockModelIdentifier)
	}

	if resp, err := client.GetFoundationModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get information about the Foundation model availability.
func bedrock_GetFoundationModelAvailability(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.GetFoundationModelAvailabilityInput{
		// ModelId: *string, // Required
	}

	if len(_bedrockModelId) > 0 {
		input.ModelId = aws.String(_bedrockModelId)
	}

	if resp, err := client.GetFoundationModelAvailability(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets details about a guardrail. If you don't specify a version, the response
// returns details for the DRAFT version.
func bedrock_GetGuardrail(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.GetGuardrailInput{
		// GuardrailIdentifier: *string, // Required
	}

	if len(_bedrockGuardrailIdentifier) > 0 {
		input.GuardrailIdentifier = aws.String(_bedrockGuardrailIdentifier)
	}
	if len(_bedrockGuardrailVersion) > 0 {
		input.GuardrailVersion = aws.String(_bedrockGuardrailVersion)
	}

	if resp, err := client.GetGuardrail(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets properties associated with a customized model you imported.
func bedrock_GetImportedModel(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.GetImportedModelInput{
		// ModelIdentifier: *string, // Required
	}

	if len(_bedrockModelIdentifier) > 0 {
		input.ModelIdentifier = aws.String(_bedrockModelIdentifier)
	}

	if resp, err := client.GetImportedModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about an inference profile. For more information, see [Increase throughput and resilience with cross-region inference in Amazon Bedrock]. in the
// Amazon Bedrock User Guide.
//
// [Increase throughput and resilience with cross-region inference in Amazon Bedrock]: https://docs.aws.amazon.com/bedrock/latest/userguide/cross-region-inference.html
func bedrock_GetInferenceProfile(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.GetInferenceProfileInput{
		// InferenceProfileIdentifier: *string, // Required
	}

	if len(_bedrockInferenceProfileIdentifier) > 0 {
		input.InferenceProfileIdentifier = aws.String(_bedrockInferenceProfileIdentifier)
	}

	if resp, err := client.GetInferenceProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves details about a specific endpoint for a model from Amazon Bedrock
// Marketplace.
func bedrock_GetMarketplaceModelEndpoint(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.GetMarketplaceModelEndpointInput{
		// EndpointArn: *string, // Required
	}

	if len(_bedrockEndpointArn) > 0 {
		input.EndpointArn = aws.String(_bedrockEndpointArn)
	}

	if resp, err := client.GetMarketplaceModelEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a model copy job. For more information, see [Copy models to be used in other regions] in the [Amazon Bedrock User Guide]
// .
//
// [Copy models to be used in other regions]: https://docs.aws.amazon.com/bedrock/latest/userguide/copy-model.html
// [Amazon Bedrock User Guide]: https://docs.aws.amazon.com/bedrock/latest/userguide/what-is-service.html
func bedrock_GetModelCopyJob(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.GetModelCopyJobInput{
		// JobArn: *string, // Required
	}

	if len(_bedrockJobArn) > 0 {
		input.JobArn = aws.String(_bedrockJobArn)
	}

	if resp, err := client.GetModelCopyJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the properties associated with a model-customization job, including
// the status of the job. For more information, see [Custom models]in the [Amazon Bedrock User Guide].
//
// [Custom models]: https://docs.aws.amazon.com/bedrock/latest/userguide/custom-models.html
// [Amazon Bedrock User Guide]: https://docs.aws.amazon.com/bedrock/latest/userguide/what-is-service.html
func bedrock_GetModelCustomizationJob(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.GetModelCustomizationJobInput{
		// JobIdentifier: *string, // Required
	}

	if len(_bedrockJobIdentifier) > 0 {
		input.JobIdentifier = aws.String(_bedrockJobIdentifier)
	}

	if resp, err := client.GetModelCustomizationJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the properties associated with import model job, including the status
// of the job. For more information, see [Import a customized model]in the [Amazon Bedrock User Guide].
//
// [Import a customized model]: https://docs.aws.amazon.com/bedrock/latest/userguide/model-customization-import-model.html
// [Amazon Bedrock User Guide]: https://docs.aws.amazon.com/bedrock/latest/userguide/what-is-service.html
func bedrock_GetModelImportJob(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.GetModelImportJobInput{
		// JobIdentifier: *string, // Required
	}

	if len(_bedrockJobIdentifier) > 0 {
		input.JobIdentifier = aws.String(_bedrockJobIdentifier)
	}

	if resp, err := client.GetModelImportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets details about a batch inference job. For more information, see [Monitor batch inference jobs]
//
// [Monitor batch inference jobs]: https://docs.aws.amazon.com/bedrock/latest/userguide/batch-inference-monitor
func bedrock_GetModelInvocationJob(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.GetModelInvocationJobInput{
		// JobIdentifier: *string, // Required
	}

	if len(_bedrockJobIdentifier) > 0 {
		input.JobIdentifier = aws.String(_bedrockJobIdentifier)
	}

	if resp, err := client.GetModelInvocationJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get the current configuration values for model invocation logging.
func bedrock_GetModelInvocationLoggingConfiguration(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.GetModelInvocationLoggingConfigurationInput{}

	if resp, err := client.GetModelInvocationLoggingConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves details about a prompt router.
func bedrock_GetPromptRouter(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.GetPromptRouterInput{
		// PromptRouterArn: *string, // Required
	}

	if len(_bedrockPromptRouterArn) > 0 {
		input.PromptRouterArn = aws.String(_bedrockPromptRouterArn)
	}

	if resp, err := client.GetPromptRouter(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns details for a Provisioned Throughput. For more information, see [Provisioned Throughput] in the [Amazon Bedrock User Guide]
// .
//
// [Provisioned Throughput]: https://docs.aws.amazon.com/bedrock/latest/userguide/prov-throughput.html
// [Amazon Bedrock User Guide]: https://docs.aws.amazon.com/bedrock/latest/userguide/what-is-service.html
func bedrock_GetProvisionedModelThroughput(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.GetProvisionedModelThroughputInput{
		// ProvisionedModelId: *string, // Required
	}

	if len(_bedrockProvisionedModelId) > 0 {
		input.ProvisionedModelId = aws.String(_bedrockProvisionedModelId)
	}

	if resp, err := client.GetProvisionedModelThroughput(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get usecase for model access.
func bedrock_GetUseCaseForModelAccess(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.GetUseCaseForModelAccessInput{}

	if resp, err := client.GetUseCaseForModelAccess(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all Automated Reasoning policies in your account, with optional filtering
// by policy ARN. This helps you manage and discover existing policies.
func bedrock_ListAutomatedReasoningPolicies(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.ListAutomatedReasoningPoliciesInput{}

	if len(_bedrockMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bedrockMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bedrockNextToken) > 0 {
		input.NextToken = aws.String(_bedrockNextToken)
	}
	if len(_bedrockPolicyArn) > 0 {
		input.PolicyArn = aws.String(_bedrockPolicyArn)
	}

	if disablePaginator() {
		if resp, err := client.ListAutomatedReasoningPolicies(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrock.ListAutomatedReasoningPoliciesOutput
	p := bedrock.NewListAutomatedReasoningPoliciesPaginator(client, input)
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

// Lists all build workflows for an Automated Reasoning policy, showing the
// history of policy creation and modification attempts.
func bedrock_ListAutomatedReasoningPolicyBuildWorkflows(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.ListAutomatedReasoningPolicyBuildWorkflowsInput{
		// PolicyArn: *string, // Required
	}

	if len(_bedrockPolicyArn) > 0 {
		input.PolicyArn = aws.String(_bedrockPolicyArn)
	}
	if len(_bedrockMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bedrockMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bedrockNextToken) > 0 {
		input.NextToken = aws.String(_bedrockNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAutomatedReasoningPolicyBuildWorkflows(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrock.ListAutomatedReasoningPolicyBuildWorkflowsOutput
	p := bedrock.NewListAutomatedReasoningPolicyBuildWorkflowsPaginator(client, input)
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

// Lists tests for an Automated Reasoning policy. We recommend using pagination to
// ensure that the operation returns quickly and successfully.
func bedrock_ListAutomatedReasoningPolicyTestCases(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.ListAutomatedReasoningPolicyTestCasesInput{
		// PolicyArn: *string, // Required
	}

	if len(_bedrockPolicyArn) > 0 {
		input.PolicyArn = aws.String(_bedrockPolicyArn)
	}
	if len(_bedrockMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bedrockMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bedrockNextToken) > 0 {
		input.NextToken = aws.String(_bedrockNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAutomatedReasoningPolicyTestCases(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrock.ListAutomatedReasoningPolicyTestCasesOutput
	p := bedrock.NewListAutomatedReasoningPolicyTestCasesPaginator(client, input)
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

// Lists test results for an Automated Reasoning policy, showing how the policy
// performed against various test scenarios and validation checks.
func bedrock_ListAutomatedReasoningPolicyTestResults(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.ListAutomatedReasoningPolicyTestResultsInput{
		// BuildWorkflowId: *string, // Required
		// PolicyArn: *string, // Required
	}

	if len(_bedrockBuildWorkflowId) > 0 {
		input.BuildWorkflowId = aws.String(_bedrockBuildWorkflowId)
	}
	if len(_bedrockPolicyArn) > 0 {
		input.PolicyArn = aws.String(_bedrockPolicyArn)
	}
	if len(_bedrockMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bedrockMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bedrockNextToken) > 0 {
		input.NextToken = aws.String(_bedrockNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAutomatedReasoningPolicyTestResults(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrock.ListAutomatedReasoningPolicyTestResultsOutput
	p := bedrock.NewListAutomatedReasoningPolicyTestResultsPaginator(client, input)
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

// Lists custom model deployments in your account. You can filter the results by
// creation time, name, status, and associated model. Use this operation to manage
// and monitor your custom model deployments.
//
// We recommend using pagination to ensure that the operation returns quickly and
// successfully.
//
// The following actions are related to the ListCustomModelDeployments operation:
//
// [CreateCustomModelDeployment]
//
// [GetCustomModelDeployment]
//
// [DeleteCustomModelDeployment]
//
// [GetCustomModelDeployment]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_GetCustomModelDeployment.html
// [CreateCustomModelDeployment]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_CreateCustomModelDeployment.html
// [DeleteCustomModelDeployment]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_DeleteCustomModelDeployment.html
func bedrock_ListCustomModelDeployments(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.ListCustomModelDeploymentsInput{}

	if len(_bedrockCreatedAfter) > 0 {
		if err := assignInputField(input, "CreatedAfter", _bedrockCreatedAfter); err != nil {
			log.Errorf("invalid --created-after: %s", err.Error())
			return
		}
	}
	if len(_bedrockCreatedBefore) > 0 {
		if err := assignInputField(input, "CreatedBefore", _bedrockCreatedBefore); err != nil {
			log.Errorf("invalid --created-before: %s", err.Error())
			return
		}
	}
	if len(_bedrockMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bedrockMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bedrockModelArnEquals) > 0 {
		input.ModelArnEquals = aws.String(_bedrockModelArnEquals)
	}
	if len(_bedrockNameContains) > 0 {
		input.NameContains = aws.String(_bedrockNameContains)
	}
	if len(_bedrockNextToken) > 0 {
		input.NextToken = aws.String(_bedrockNextToken)
	}
	if len(_bedrockSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _bedrockSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_bedrockSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _bedrockSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}
	if len(_bedrockStatusEquals) > 0 {
		if err := assignInputField(input, "StatusEquals", _bedrockStatusEquals); err != nil {
			log.Errorf("invalid --status-equals: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListCustomModelDeployments(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrock.ListCustomModelDeploymentsOutput
	p := bedrock.NewListCustomModelDeploymentsPaginator(client, input)
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

// Returns a list of the custom models that you have created with the
// CreateModelCustomizationJob operation.
//
// For more information, see [Custom models] in the [Amazon Bedrock User Guide].
//
// [Custom models]: https://docs.aws.amazon.com/bedrock/latest/userguide/custom-models.html
// [Amazon Bedrock User Guide]: https://docs.aws.amazon.com/bedrock/latest/userguide/what-is-service.html
func bedrock_ListCustomModels(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.ListCustomModelsInput{}

	if len(_bedrockBaseModelArnEquals) > 0 {
		input.BaseModelArnEquals = aws.String(_bedrockBaseModelArnEquals)
	}
	if len(_bedrockCreationTimeAfter) > 0 {
		if err := assignInputField(input, "CreationTimeAfter", _bedrockCreationTimeAfter); err != nil {
			log.Errorf("invalid --creation-time-after: %s", err.Error())
			return
		}
	}
	if len(_bedrockCreationTimeBefore) > 0 {
		if err := assignInputField(input, "CreationTimeBefore", _bedrockCreationTimeBefore); err != nil {
			log.Errorf("invalid --creation-time-before: %s", err.Error())
			return
		}
	}
	if len(_bedrockFoundationModelArnEquals) > 0 {
		input.FoundationModelArnEquals = aws.String(_bedrockFoundationModelArnEquals)
	}
	if len(_bedrockIsOwned) > 0 {
		if err := assignInputField(input, "IsOwned", _bedrockIsOwned); err != nil {
			log.Errorf("invalid --is-owned: %s", err.Error())
			return
		}
	}
	if len(_bedrockMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bedrockMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bedrockModelStatus) > 0 {
		if err := assignInputField(input, "ModelStatus", _bedrockModelStatus); err != nil {
			log.Errorf("invalid --model-status: %s", err.Error())
			return
		}
	}
	if len(_bedrockNameContains) > 0 {
		input.NameContains = aws.String(_bedrockNameContains)
	}
	if len(_bedrockNextToken) > 0 {
		input.NextToken = aws.String(_bedrockNextToken)
	}
	if len(_bedrockSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _bedrockSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_bedrockSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _bedrockSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListCustomModels(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrock.ListCustomModelsOutput
	p := bedrock.NewListCustomModelsPaginator(client, input)
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

// Lists the account-level enforced guardrail configurations.
func bedrock_ListEnforcedGuardrailsConfiguration(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.ListEnforcedGuardrailsConfigurationInput{}

	if len(_bedrockNextToken) > 0 {
		input.NextToken = aws.String(_bedrockNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListEnforcedGuardrailsConfiguration(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrock.ListEnforcedGuardrailsConfigurationOutput
	p := bedrock.NewListEnforcedGuardrailsConfigurationPaginator(client, input)
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

// Lists all existing evaluation jobs.
func bedrock_ListEvaluationJobs(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.ListEvaluationJobsInput{}

	if len(_bedrockApplicationTypeEquals) > 0 {
		if err := assignInputField(input, "ApplicationTypeEquals", _bedrockApplicationTypeEquals); err != nil {
			log.Errorf("invalid --application-type-equals: %s", err.Error())
			return
		}
	}
	if len(_bedrockCreationTimeAfter) > 0 {
		if err := assignInputField(input, "CreationTimeAfter", _bedrockCreationTimeAfter); err != nil {
			log.Errorf("invalid --creation-time-after: %s", err.Error())
			return
		}
	}
	if len(_bedrockCreationTimeBefore) > 0 {
		if err := assignInputField(input, "CreationTimeBefore", _bedrockCreationTimeBefore); err != nil {
			log.Errorf("invalid --creation-time-before: %s", err.Error())
			return
		}
	}
	if len(_bedrockMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bedrockMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bedrockNameContains) > 0 {
		input.NameContains = aws.String(_bedrockNameContains)
	}
	if len(_bedrockNextToken) > 0 {
		input.NextToken = aws.String(_bedrockNextToken)
	}
	if len(_bedrockSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _bedrockSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_bedrockSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _bedrockSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}
	if len(_bedrockStatusEquals) > 0 {
		if err := assignInputField(input, "StatusEquals", _bedrockStatusEquals); err != nil {
			log.Errorf("invalid --status-equals: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListEvaluationJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrock.ListEvaluationJobsOutput
	p := bedrock.NewListEvaluationJobsPaginator(client, input)
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

// Get the offers associated with the specified model.
func bedrock_ListFoundationModelAgreementOffers(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.ListFoundationModelAgreementOffersInput{
		// ModelId: *string, // Required
	}

	if len(_bedrockModelId) > 0 {
		input.ModelId = aws.String(_bedrockModelId)
	}
	if len(_bedrockOfferType) > 0 {
		if err := assignInputField(input, "OfferType", _bedrockOfferType); err != nil {
			log.Errorf("invalid --offer-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListFoundationModelAgreementOffers(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists Amazon Bedrock foundation models that you can use. You can filter the
// results with the request parameters. For more information, see [Foundation models]in the [Amazon Bedrock User Guide].
//
// [Foundation models]: https://docs.aws.amazon.com/bedrock/latest/userguide/foundation-models.html
// [Amazon Bedrock User Guide]: https://docs.aws.amazon.com/bedrock/latest/userguide/what-is-service.html
func bedrock_ListFoundationModels(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.ListFoundationModelsInput{}

	if len(_bedrockByCustomizationType) > 0 {
		if err := assignInputField(input, "ByCustomizationType", _bedrockByCustomizationType); err != nil {
			log.Errorf("invalid --by-customization-type: %s", err.Error())
			return
		}
	}
	if len(_bedrockByInferenceType) > 0 {
		if err := assignInputField(input, "ByInferenceType", _bedrockByInferenceType); err != nil {
			log.Errorf("invalid --by-inference-type: %s", err.Error())
			return
		}
	}
	if len(_bedrockByOutputModality) > 0 {
		if err := assignInputField(input, "ByOutputModality", _bedrockByOutputModality); err != nil {
			log.Errorf("invalid --by-output-modality: %s", err.Error())
			return
		}
	}
	if len(_bedrockByProvider) > 0 {
		input.ByProvider = aws.String(_bedrockByProvider)
	}

	if resp, err := client.ListFoundationModels(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists details about all the guardrails in an account. To list the DRAFT version
// of all your guardrails, don't specify the guardrailIdentifier field. To list
// all versions of a guardrail, specify the ARN of the guardrail in the
// guardrailIdentifier field.
//
// You can set the maximum number of results to return in a response in the
// maxResults field. If there are more results than the number you set, the
// response returns a nextToken that you can send in another ListGuardrails
// request to see the next batch of results.
func bedrock_ListGuardrails(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.ListGuardrailsInput{}

	if len(_bedrockGuardrailIdentifier) > 0 {
		input.GuardrailIdentifier = aws.String(_bedrockGuardrailIdentifier)
	}
	if len(_bedrockMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bedrockMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bedrockNextToken) > 0 {
		input.NextToken = aws.String(_bedrockNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListGuardrails(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrock.ListGuardrailsOutput
	p := bedrock.NewListGuardrailsPaginator(client, input)
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

// Returns a list of models you've imported. You can filter the results to return
// based on one or more criteria. For more information, see [Import a customized model]in the [Amazon Bedrock User Guide].
//
// [Import a customized model]: https://docs.aws.amazon.com/bedrock/latest/userguide/model-customization-import-model.html
// [Amazon Bedrock User Guide]: https://docs.aws.amazon.com/bedrock/latest/userguide/what-is-service.html
func bedrock_ListImportedModels(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.ListImportedModelsInput{}

	if len(_bedrockCreationTimeAfter) > 0 {
		if err := assignInputField(input, "CreationTimeAfter", _bedrockCreationTimeAfter); err != nil {
			log.Errorf("invalid --creation-time-after: %s", err.Error())
			return
		}
	}
	if len(_bedrockCreationTimeBefore) > 0 {
		if err := assignInputField(input, "CreationTimeBefore", _bedrockCreationTimeBefore); err != nil {
			log.Errorf("invalid --creation-time-before: %s", err.Error())
			return
		}
	}
	if len(_bedrockMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bedrockMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bedrockNameContains) > 0 {
		input.NameContains = aws.String(_bedrockNameContains)
	}
	if len(_bedrockNextToken) > 0 {
		input.NextToken = aws.String(_bedrockNextToken)
	}
	if len(_bedrockSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _bedrockSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_bedrockSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _bedrockSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListImportedModels(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrock.ListImportedModelsOutput
	p := bedrock.NewListImportedModelsPaginator(client, input)
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

// Returns a list of inference profiles that you can use. For more information,
// see [Increase throughput and resilience with cross-region inference in Amazon Bedrock]. in the Amazon Bedrock User Guide.
//
// [Increase throughput and resilience with cross-region inference in Amazon Bedrock]: https://docs.aws.amazon.com/bedrock/latest/userguide/cross-region-inference.html
func bedrock_ListInferenceProfiles(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.ListInferenceProfilesInput{}

	if len(_bedrockMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bedrockMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bedrockNextToken) > 0 {
		input.NextToken = aws.String(_bedrockNextToken)
	}
	if len(_bedrockTypeEquals) > 0 {
		if err := assignInputField(input, "TypeEquals", _bedrockTypeEquals); err != nil {
			log.Errorf("invalid --type-equals: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListInferenceProfiles(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrock.ListInferenceProfilesOutput
	p := bedrock.NewListInferenceProfilesPaginator(client, input)
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

// Lists the endpoints for models from Amazon Bedrock Marketplace in your Amazon
// Web Services account.
func bedrock_ListMarketplaceModelEndpoints(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.ListMarketplaceModelEndpointsInput{}

	if len(_bedrockMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bedrockMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bedrockModelSourceEquals) > 0 {
		input.ModelSourceEquals = aws.String(_bedrockModelSourceEquals)
	}
	if len(_bedrockNextToken) > 0 {
		input.NextToken = aws.String(_bedrockNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListMarketplaceModelEndpoints(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrock.ListMarketplaceModelEndpointsOutput
	p := bedrock.NewListMarketplaceModelEndpointsPaginator(client, input)
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

// Returns a list of model copy jobs that you have submitted. You can filter the
// jobs to return based on one or more criteria. For more information, see [Copy models to be used in other regions]in the [Amazon Bedrock User Guide].
//
// [Copy models to be used in other regions]: https://docs.aws.amazon.com/bedrock/latest/userguide/copy-model.html
// [Amazon Bedrock User Guide]: https://docs.aws.amazon.com/bedrock/latest/userguide/what-is-service.html
func bedrock_ListModelCopyJobs(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.ListModelCopyJobsInput{}

	if len(_bedrockCreationTimeAfter) > 0 {
		if err := assignInputField(input, "CreationTimeAfter", _bedrockCreationTimeAfter); err != nil {
			log.Errorf("invalid --creation-time-after: %s", err.Error())
			return
		}
	}
	if len(_bedrockCreationTimeBefore) > 0 {
		if err := assignInputField(input, "CreationTimeBefore", _bedrockCreationTimeBefore); err != nil {
			log.Errorf("invalid --creation-time-before: %s", err.Error())
			return
		}
	}
	if len(_bedrockMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bedrockMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bedrockNextToken) > 0 {
		input.NextToken = aws.String(_bedrockNextToken)
	}
	if len(_bedrockSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _bedrockSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_bedrockSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _bedrockSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}
	if len(_bedrockSourceAccountEquals) > 0 {
		input.SourceAccountEquals = aws.String(_bedrockSourceAccountEquals)
	}
	if len(_bedrockSourceModelArnEquals) > 0 {
		input.SourceModelArnEquals = aws.String(_bedrockSourceModelArnEquals)
	}
	if len(_bedrockStatusEquals) > 0 {
		if err := assignInputField(input, "StatusEquals", _bedrockStatusEquals); err != nil {
			log.Errorf("invalid --status-equals: %s", err.Error())
			return
		}
	}
	if len(_bedrockTargetModelNameContains) > 0 {
		input.TargetModelNameContains = aws.String(_bedrockTargetModelNameContains)
	}

	if disablePaginator() {
		if resp, err := client.ListModelCopyJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrock.ListModelCopyJobsOutput
	p := bedrock.NewListModelCopyJobsPaginator(client, input)
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

// Returns a list of model customization jobs that you have submitted. You can
// filter the jobs to return based on one or more criteria.
//
// For more information, see [Custom models] in the [Amazon Bedrock User Guide].
//
// [Custom models]: https://docs.aws.amazon.com/bedrock/latest/userguide/custom-models.html
// [Amazon Bedrock User Guide]: https://docs.aws.amazon.com/bedrock/latest/userguide/what-is-service.html
func bedrock_ListModelCustomizationJobs(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.ListModelCustomizationJobsInput{}

	if len(_bedrockCreationTimeAfter) > 0 {
		if err := assignInputField(input, "CreationTimeAfter", _bedrockCreationTimeAfter); err != nil {
			log.Errorf("invalid --creation-time-after: %s", err.Error())
			return
		}
	}
	if len(_bedrockCreationTimeBefore) > 0 {
		if err := assignInputField(input, "CreationTimeBefore", _bedrockCreationTimeBefore); err != nil {
			log.Errorf("invalid --creation-time-before: %s", err.Error())
			return
		}
	}
	if len(_bedrockMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bedrockMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bedrockNameContains) > 0 {
		input.NameContains = aws.String(_bedrockNameContains)
	}
	if len(_bedrockNextToken) > 0 {
		input.NextToken = aws.String(_bedrockNextToken)
	}
	if len(_bedrockSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _bedrockSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_bedrockSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _bedrockSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}
	if len(_bedrockStatusEquals) > 0 {
		if err := assignInputField(input, "StatusEquals", _bedrockStatusEquals); err != nil {
			log.Errorf("invalid --status-equals: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListModelCustomizationJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrock.ListModelCustomizationJobsOutput
	p := bedrock.NewListModelCustomizationJobsPaginator(client, input)
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

// Returns a list of import jobs you've submitted. You can filter the results to
// return based on one or more criteria. For more information, see [Import a customized model]in the [Amazon Bedrock User Guide].
//
// [Import a customized model]: https://docs.aws.amazon.com/bedrock/latest/userguide/model-customization-import-model.html
// [Amazon Bedrock User Guide]: https://docs.aws.amazon.com/bedrock/latest/userguide/what-is-service.html
func bedrock_ListModelImportJobs(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.ListModelImportJobsInput{}

	if len(_bedrockCreationTimeAfter) > 0 {
		if err := assignInputField(input, "CreationTimeAfter", _bedrockCreationTimeAfter); err != nil {
			log.Errorf("invalid --creation-time-after: %s", err.Error())
			return
		}
	}
	if len(_bedrockCreationTimeBefore) > 0 {
		if err := assignInputField(input, "CreationTimeBefore", _bedrockCreationTimeBefore); err != nil {
			log.Errorf("invalid --creation-time-before: %s", err.Error())
			return
		}
	}
	if len(_bedrockMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bedrockMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bedrockNameContains) > 0 {
		input.NameContains = aws.String(_bedrockNameContains)
	}
	if len(_bedrockNextToken) > 0 {
		input.NextToken = aws.String(_bedrockNextToken)
	}
	if len(_bedrockSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _bedrockSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_bedrockSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _bedrockSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}
	if len(_bedrockStatusEquals) > 0 {
		if err := assignInputField(input, "StatusEquals", _bedrockStatusEquals); err != nil {
			log.Errorf("invalid --status-equals: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListModelImportJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrock.ListModelImportJobsOutput
	p := bedrock.NewListModelImportJobsPaginator(client, input)
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

// Lists all batch inference jobs in the account. For more information, see [View details about a batch inference job].
//
// [View details about a batch inference job]: https://docs.aws.amazon.com/bedrock/latest/userguide/batch-inference-view.html
func bedrock_ListModelInvocationJobs(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.ListModelInvocationJobsInput{}

	if len(_bedrockMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bedrockMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bedrockNameContains) > 0 {
		input.NameContains = aws.String(_bedrockNameContains)
	}
	if len(_bedrockNextToken) > 0 {
		input.NextToken = aws.String(_bedrockNextToken)
	}
	if len(_bedrockSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _bedrockSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_bedrockSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _bedrockSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}
	if len(_bedrockStatusEquals) > 0 {
		if err := assignInputField(input, "StatusEquals", _bedrockStatusEquals); err != nil {
			log.Errorf("invalid --status-equals: %s", err.Error())
			return
		}
	}
	if len(_bedrockSubmitTimeAfter) > 0 {
		if err := assignInputField(input, "SubmitTimeAfter", _bedrockSubmitTimeAfter); err != nil {
			log.Errorf("invalid --submit-time-after: %s", err.Error())
			return
		}
	}
	if len(_bedrockSubmitTimeBefore) > 0 {
		if err := assignInputField(input, "SubmitTimeBefore", _bedrockSubmitTimeBefore); err != nil {
			log.Errorf("invalid --submit-time-before: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListModelInvocationJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrock.ListModelInvocationJobsOutput
	p := bedrock.NewListModelInvocationJobsPaginator(client, input)
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

// Retrieves a list of prompt routers.
func bedrock_ListPromptRouters(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.ListPromptRoutersInput{}

	if len(_bedrockMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bedrockMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bedrockNextToken) > 0 {
		input.NextToken = aws.String(_bedrockNextToken)
	}
	if len(_bedrockType) > 0 {
		if err := assignInputField(input, "Type", _bedrockType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListPromptRouters(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrock.ListPromptRoutersOutput
	p := bedrock.NewListPromptRoutersPaginator(client, input)
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

// Lists the Provisioned Throughputs in the account. For more information, see [Provisioned Throughput] in
// the [Amazon Bedrock User Guide].
//
// [Amazon Bedrock User Guide]: https://docs.aws.amazon.com/bedrock/latest/userguide/what-is-service.html
// [Provisioned Throughput]: https://docs.aws.amazon.com/bedrock/latest/userguide/prov-throughput.html
func bedrock_ListProvisionedModelThroughputs(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.ListProvisionedModelThroughputsInput{}

	if len(_bedrockCreationTimeAfter) > 0 {
		if err := assignInputField(input, "CreationTimeAfter", _bedrockCreationTimeAfter); err != nil {
			log.Errorf("invalid --creation-time-after: %s", err.Error())
			return
		}
	}
	if len(_bedrockCreationTimeBefore) > 0 {
		if err := assignInputField(input, "CreationTimeBefore", _bedrockCreationTimeBefore); err != nil {
			log.Errorf("invalid --creation-time-before: %s", err.Error())
			return
		}
	}
	if len(_bedrockMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bedrockMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bedrockModelArnEquals) > 0 {
		input.ModelArnEquals = aws.String(_bedrockModelArnEquals)
	}
	if len(_bedrockNameContains) > 0 {
		input.NameContains = aws.String(_bedrockNameContains)
	}
	if len(_bedrockNextToken) > 0 {
		input.NextToken = aws.String(_bedrockNextToken)
	}
	if len(_bedrockSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _bedrockSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_bedrockSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _bedrockSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}
	if len(_bedrockStatusEquals) > 0 {
		if err := assignInputField(input, "StatusEquals", _bedrockStatusEquals); err != nil {
			log.Errorf("invalid --status-equals: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListProvisionedModelThroughputs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrock.ListProvisionedModelThroughputsOutput
	p := bedrock.NewListProvisionedModelThroughputsPaginator(client, input)
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

// List the tags associated with the specified resource.
// For more information, see [Tagging resources] in the [Amazon Bedrock User Guide].
//
// [Amazon Bedrock User Guide]: https://docs.aws.amazon.com/bedrock/latest/userguide/what-is-service.html
// [Tagging resources]: https://docs.aws.amazon.com/bedrock/latest/userguide/what-is-service.html
func bedrock_ListTagsForResource(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.ListTagsForResourceInput{
		// ResourceARN: *string, // Required
	}

	if len(_bedrockResourceARN) > 0 {
		input.ResourceARN = aws.String(_bedrockResourceARN)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets the account-level enforced guardrail configuration.
func bedrock_PutEnforcedGuardrailConfiguration(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.PutEnforcedGuardrailConfigurationInput{
		// GuardrailInferenceConfig: *types.AccountEnforcedGuardrailInferenceInputConfiguration, // Required
	}

	if len(_bedrockGuardrailInferenceConfig) > 0 {
		if err := assignInputField(input, "GuardrailInferenceConfig", _bedrockGuardrailInferenceConfig); err != nil {
			log.Errorf("invalid --guardrail-inference-config: %s", err.Error())
			return
		}
	}
	if len(_bedrockConfigId) > 0 {
		input.ConfigId = aws.String(_bedrockConfigId)
	}

	if resp, err := client.PutEnforcedGuardrailConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Set the configuration values for model invocation logging.
func bedrock_PutModelInvocationLoggingConfiguration(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.PutModelInvocationLoggingConfigurationInput{
		// LoggingConfig: *types.LoggingConfig, // Required
	}

	if len(_bedrockLoggingConfig) > 0 {
		if err := assignInputField(input, "LoggingConfig", _bedrockLoggingConfig); err != nil {
			log.Errorf("invalid --logging-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutModelInvocationLoggingConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Put usecase for model access.
func bedrock_PutUseCaseForModelAccess(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.PutUseCaseForModelAccessInput{
		// FormData: []byte, // Required
	}

	if len(_bedrockFormData) > 0 {
		if err := assignInputField(input, "FormData", _bedrockFormData); err != nil {
			log.Errorf("invalid --form-data: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutUseCaseForModelAccess(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Registers an existing Amazon SageMaker endpoint with Amazon Bedrock
// Marketplace, allowing it to be used with Amazon Bedrock APIs.
func bedrock_RegisterMarketplaceModelEndpoint(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.RegisterMarketplaceModelEndpointInput{
		// EndpointIdentifier: *string, // Required
		// ModelSourceIdentifier: *string, // Required
	}

	if len(_bedrockEndpointIdentifier) > 0 {
		input.EndpointIdentifier = aws.String(_bedrockEndpointIdentifier)
	}
	if len(_bedrockModelSourceIdentifier) > 0 {
		input.ModelSourceIdentifier = aws.String(_bedrockModelSourceIdentifier)
	}

	if resp, err := client.RegisterMarketplaceModelEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a new build workflow for an Automated Reasoning policy. This initiates
// the process of analyzing source documents and generating policy rules,
// variables, and types.
func bedrock_StartAutomatedReasoningPolicyBuildWorkflow(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.StartAutomatedReasoningPolicyBuildWorkflowInput{
		// BuildWorkflowType: types.AutomatedReasoningPolicyBuildWorkflowType, // Required
		// PolicyArn: *string, // Required
		// SourceContent: *types.AutomatedReasoningPolicyBuildWorkflowSource, // Required
	}

	if len(_bedrockBuildWorkflowType) > 0 {
		if err := assignInputField(input, "BuildWorkflowType", _bedrockBuildWorkflowType); err != nil {
			log.Errorf("invalid --build-workflow-type: %s", err.Error())
			return
		}
	}
	if len(_bedrockPolicyArn) > 0 {
		input.PolicyArn = aws.String(_bedrockPolicyArn)
	}
	if len(_bedrockSourceContent) > 0 {
		if err := assignInputField(input, "SourceContent", _bedrockSourceContent); err != nil {
			log.Errorf("invalid --source-content: %s", err.Error())
			return
		}
	}
	if len(_bedrockClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_bedrockClientRequestToken)
	}

	if resp, err := client.StartAutomatedReasoningPolicyBuildWorkflow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Initiates a test workflow to validate Automated Reasoning policy tests. The
// workflow executes the specified tests against the policy and generates
// validation results.
func bedrock_StartAutomatedReasoningPolicyTestWorkflow(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.StartAutomatedReasoningPolicyTestWorkflowInput{
		// BuildWorkflowId: *string, // Required
		// PolicyArn: *string, // Required
	}

	if len(_bedrockBuildWorkflowId) > 0 {
		input.BuildWorkflowId = aws.String(_bedrockBuildWorkflowId)
	}
	if len(_bedrockPolicyArn) > 0 {
		input.PolicyArn = aws.String(_bedrockPolicyArn)
	}
	if len(_bedrockClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_bedrockClientRequestToken)
	}
	if len(_bedrockTestCaseIds) > 0 {
		input.TestCaseIds = append([]string(nil), _bedrockTestCaseIds...)
	}

	if resp, err := client.StartAutomatedReasoningPolicyTestWorkflow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops an evaluation job that is current being created or running.
func bedrock_StopEvaluationJob(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.StopEvaluationJobInput{
		// JobIdentifier: *string, // Required
	}

	if len(_bedrockJobIdentifier) > 0 {
		input.JobIdentifier = aws.String(_bedrockJobIdentifier)
	}

	if resp, err := client.StopEvaluationJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops an active model customization job. For more information, see [Custom models] in the [Amazon Bedrock User Guide].
//
// [Custom models]: https://docs.aws.amazon.com/bedrock/latest/userguide/custom-models.html
// [Amazon Bedrock User Guide]: https://docs.aws.amazon.com/bedrock/latest/userguide/what-is-service.html
func bedrock_StopModelCustomizationJob(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.StopModelCustomizationJobInput{
		// JobIdentifier: *string, // Required
	}

	if len(_bedrockJobIdentifier) > 0 {
		input.JobIdentifier = aws.String(_bedrockJobIdentifier)
	}

	if resp, err := client.StopModelCustomizationJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a batch inference job. You're only charged for tokens that were already
// processed. For more information, see [Stop a batch inference job].
//
// [Stop a batch inference job]: https://docs.aws.amazon.com/bedrock/latest/userguide/batch-inference-stop.html
func bedrock_StopModelInvocationJob(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.StopModelInvocationJobInput{
		// JobIdentifier: *string, // Required
	}

	if len(_bedrockJobIdentifier) > 0 {
		input.JobIdentifier = aws.String(_bedrockJobIdentifier)
	}

	if resp, err := client.StopModelInvocationJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associate tags with a resource. For more information, see [Tagging resources] in the [Amazon Bedrock User Guide].
//
// [Amazon Bedrock User Guide]: https://docs.aws.amazon.com/bedrock/latest/userguide/what-is-service.html
// [Tagging resources]: https://docs.aws.amazon.com/bedrock/latest/userguide/what-is-service.html
func bedrock_TagResource(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.TagResourceInput{
		// ResourceARN: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_bedrockResourceARN) > 0 {
		input.ResourceARN = aws.String(_bedrockResourceARN)
	}
	if len(_bedrockTags) > 0 {
		if err := assignInputField(input, "Tags", _bedrockTags); err != nil {
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

// Remove one or more tags from a resource. For more information, see [Tagging resources] in the [Amazon Bedrock User Guide].
//
// [Amazon Bedrock User Guide]: https://docs.aws.amazon.com/bedrock/latest/userguide/what-is-service.html
// [Tagging resources]: https://docs.aws.amazon.com/bedrock/latest/userguide/what-is-service.html
func bedrock_UntagResource(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.UntagResourceInput{
		// ResourceARN: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_bedrockResourceARN) > 0 {
		input.ResourceARN = aws.String(_bedrockResourceARN)
	}
	if len(_bedrockTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _bedrockTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing Automated Reasoning policy with new rules, variables, or
// configuration. This creates a new version of the policy while preserving the
// previous version.
func bedrock_UpdateAutomatedReasoningPolicy(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.UpdateAutomatedReasoningPolicyInput{
		// PolicyArn: *string, // Required
		// PolicyDefinition: *types.AutomatedReasoningPolicyDefinition, // Required
	}

	if len(_bedrockPolicyArn) > 0 {
		input.PolicyArn = aws.String(_bedrockPolicyArn)
	}
	if len(_bedrockPolicyDefinition) > 0 {
		if err := assignInputField(input, "PolicyDefinition", _bedrockPolicyDefinition); err != nil {
			log.Errorf("invalid --policy-definition: %s", err.Error())
			return
		}
	}
	if len(_bedrockDescription) > 0 {
		input.Description = aws.String(_bedrockDescription)
	}
	if len(_bedrockName) > 0 {
		input.Name = aws.String(_bedrockName)
	}

	if resp, err := client.UpdateAutomatedReasoningPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the annotations for an Automated Reasoning policy build workflow. This
// allows you to modify extracted rules, variables, and types before finalizing the
// policy.
func bedrock_UpdateAutomatedReasoningPolicyAnnotations(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.UpdateAutomatedReasoningPolicyAnnotationsInput{
		// Annotations: []types.AutomatedReasoningPolicyAnnotation, // Required
		// BuildWorkflowId: *string, // Required
		// LastUpdatedAnnotationSetHash: *string, // Required
		// PolicyArn: *string, // Required
	}

	if len(_bedrockAnnotations) > 0 {
		if err := assignInputField(input, "Annotations", _bedrockAnnotations); err != nil {
			log.Errorf("invalid --annotations: %s", err.Error())
			return
		}
	}
	if len(_bedrockBuildWorkflowId) > 0 {
		input.BuildWorkflowId = aws.String(_bedrockBuildWorkflowId)
	}
	if len(_bedrockLastUpdatedAnnotationSetHash) > 0 {
		input.LastUpdatedAnnotationSetHash = aws.String(_bedrockLastUpdatedAnnotationSetHash)
	}
	if len(_bedrockPolicyArn) > 0 {
		input.PolicyArn = aws.String(_bedrockPolicyArn)
	}

	if resp, err := client.UpdateAutomatedReasoningPolicyAnnotations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing Automated Reasoning policy test. You can modify the
// content, query, expected result, and confidence threshold.
func bedrock_UpdateAutomatedReasoningPolicyTestCase(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.UpdateAutomatedReasoningPolicyTestCaseInput{
		// ExpectedAggregatedFindingsResult: types.AutomatedReasoningCheckResult, // Required
		// GuardContent: *string, // Required
		// LastUpdatedAt: *time.Time, // Required
		// PolicyArn: *string, // Required
		// TestCaseId: *string, // Required
	}

	if len(_bedrockExpectedAggregatedFindingsResult) > 0 {
		if err := assignInputField(input, "ExpectedAggregatedFindingsResult", _bedrockExpectedAggregatedFindingsResult); err != nil {
			log.Errorf("invalid --expected-aggregated-findings-result: %s", err.Error())
			return
		}
	}
	if len(_bedrockGuardContent) > 0 {
		input.GuardContent = aws.String(_bedrockGuardContent)
	}
	if len(_bedrockLastUpdatedAt) > 0 {
		if err := assignInputField(input, "LastUpdatedAt", _bedrockLastUpdatedAt); err != nil {
			log.Errorf("invalid --last-updated-at: %s", err.Error())
			return
		}
	}
	if len(_bedrockPolicyArn) > 0 {
		input.PolicyArn = aws.String(_bedrockPolicyArn)
	}
	if len(_bedrockTestCaseId) > 0 {
		input.TestCaseId = aws.String(_bedrockTestCaseId)
	}
	if len(_bedrockClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_bedrockClientRequestToken)
	}
	if len(_bedrockConfidenceThreshold) > 0 {
		if err := assignInputField(input, "ConfidenceThreshold", _bedrockConfidenceThreshold); err != nil {
			log.Errorf("invalid --confidence-threshold: %s", err.Error())
			return
		}
	}
	if len(_bedrockQueryContent) > 0 {
		input.QueryContent = aws.String(_bedrockQueryContent)
	}

	if resp, err := client.UpdateAutomatedReasoningPolicyTestCase(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a custom model deployment with a new custom model. This allows you to
// deploy updated models without creating new deployment endpoints.
func bedrock_UpdateCustomModelDeployment(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.UpdateCustomModelDeploymentInput{
		// CustomModelDeploymentIdentifier: *string, // Required
		// ModelArn: *string, // Required
	}

	if len(_bedrockCustomModelDeploymentIdentifier) > 0 {
		input.CustomModelDeploymentIdentifier = aws.String(_bedrockCustomModelDeploymentIdentifier)
	}
	if len(_bedrockModelArn) > 0 {
		input.ModelArn = aws.String(_bedrockModelArn)
	}

	if resp, err := client.UpdateCustomModelDeployment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a guardrail with the values you specify.
// - Specify a name and optional description .
//
// - Specify messages for when the guardrail successfully blocks a prompt or a
// model response in the blockedInputMessaging and blockedOutputsMessaging fields.
//
// - Specify topics for the guardrail to deny in the topicPolicyConfig object.
// Each [GuardrailTopicConfig]object in the topicsConfig list pertains to one topic.
//
// - Give a name and description so that the guardrail can properly identify the
// topic.
//
// - Specify DENY in the type field.
//
// - (Optional) Provide up to five prompts that you would categorize as
// belonging to the topic in the examples list.
//
// - Specify filter strengths for the harmful categories defined in Amazon
// Bedrock in the contentPolicyConfig object. Each [GuardrailContentFilterConfig]object in the filtersConfig
// list pertains to a harmful category. For more information, see [Content filters]. For more
// information about the fields in a content filter, see [GuardrailContentFilterConfig].
//
// - Specify the category in the type field.
//
// - Specify the strength of the filter for prompts in the inputStrength field
// and for model responses in the strength field of the [GuardrailContentFilterConfig].
//
// - (Optional) For security, include the ARN of a KMS key in the kmsKeyId field.
//
// [GuardrailContentFilterConfig]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_GuardrailContentFilterConfig.html
// [Content filters]: https://docs.aws.amazon.com/bedrock/latest/userguide/guardrails-content-filters
// [GuardrailTopicConfig]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_GuardrailTopicConfig.html
func bedrock_UpdateGuardrail(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.UpdateGuardrailInput{
		// BlockedInputMessaging: *string, // Required
		// BlockedOutputsMessaging: *string, // Required
		// GuardrailIdentifier: *string, // Required
		// Name: *string, // Required
	}

	if len(_bedrockBlockedInputMessaging) > 0 {
		input.BlockedInputMessaging = aws.String(_bedrockBlockedInputMessaging)
	}
	if len(_bedrockBlockedOutputsMessaging) > 0 {
		input.BlockedOutputsMessaging = aws.String(_bedrockBlockedOutputsMessaging)
	}
	if len(_bedrockGuardrailIdentifier) > 0 {
		input.GuardrailIdentifier = aws.String(_bedrockGuardrailIdentifier)
	}
	if len(_bedrockName) > 0 {
		input.Name = aws.String(_bedrockName)
	}
	if len(_bedrockAutomatedReasoningPolicyConfig) > 0 {
		if err := assignInputField(input, "AutomatedReasoningPolicyConfig", _bedrockAutomatedReasoningPolicyConfig); err != nil {
			log.Errorf("invalid --automated-reasoning-policy-config: %s", err.Error())
			return
		}
	}
	if len(_bedrockContentPolicyConfig) > 0 {
		if err := assignInputField(input, "ContentPolicyConfig", _bedrockContentPolicyConfig); err != nil {
			log.Errorf("invalid --content-policy-config: %s", err.Error())
			return
		}
	}
	if len(_bedrockContextualGroundingPolicyConfig) > 0 {
		if err := assignInputField(input, "ContextualGroundingPolicyConfig", _bedrockContextualGroundingPolicyConfig); err != nil {
			log.Errorf("invalid --contextual-grounding-policy-config: %s", err.Error())
			return
		}
	}
	if len(_bedrockCrossRegionConfig) > 0 {
		if err := assignInputField(input, "CrossRegionConfig", _bedrockCrossRegionConfig); err != nil {
			log.Errorf("invalid --cross-region-config: %s", err.Error())
			return
		}
	}
	if len(_bedrockDescription) > 0 {
		input.Description = aws.String(_bedrockDescription)
	}
	if len(_bedrockKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_bedrockKmsKeyId)
	}
	if len(_bedrockSensitiveInformationPolicyConfig) > 0 {
		if err := assignInputField(input, "SensitiveInformationPolicyConfig", _bedrockSensitiveInformationPolicyConfig); err != nil {
			log.Errorf("invalid --sensitive-information-policy-config: %s", err.Error())
			return
		}
	}
	if len(_bedrockTopicPolicyConfig) > 0 {
		if err := assignInputField(input, "TopicPolicyConfig", _bedrockTopicPolicyConfig); err != nil {
			log.Errorf("invalid --topic-policy-config: %s", err.Error())
			return
		}
	}
	if len(_bedrockWordPolicyConfig) > 0 {
		if err := assignInputField(input, "WordPolicyConfig", _bedrockWordPolicyConfig); err != nil {
			log.Errorf("invalid --word-policy-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateGuardrail(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration of an existing endpoint for a model from Amazon
// Bedrock Marketplace.
func bedrock_UpdateMarketplaceModelEndpoint(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.UpdateMarketplaceModelEndpointInput{
		// EndpointArn: *string, // Required
		// EndpointConfig: types.EndpointConfig, // Required
	}

	if len(_bedrockEndpointArn) > 0 {
		input.EndpointArn = aws.String(_bedrockEndpointArn)
	}
	if len(_bedrockEndpointConfig) > 0 {
		if err := assignInputField(input, "EndpointConfig", _bedrockEndpointConfig); err != nil {
			log.Errorf("invalid --endpoint-config: %s", err.Error())
			return
		}
	}
	if len(_bedrockClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_bedrockClientRequestToken)
	}

	if resp, err := client.UpdateMarketplaceModelEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the name or associated model for a Provisioned Throughput. For more
// information, see [Provisioned Throughput]in the [Amazon Bedrock User Guide].
//
// [Provisioned Throughput]: https://docs.aws.amazon.com/bedrock/latest/userguide/prov-throughput.html
// [Amazon Bedrock User Guide]: https://docs.aws.amazon.com/bedrock/latest/userguide/what-is-service.html
func bedrock_UpdateProvisionedModelThroughput(cfg aws.Config, client *bedrock.Client) {
	input := &bedrock.UpdateProvisionedModelThroughputInput{
		// ProvisionedModelId: *string, // Required
	}

	if len(_bedrockProvisionedModelId) > 0 {
		input.ProvisionedModelId = aws.String(_bedrockProvisionedModelId)
	}
	if len(_bedrockDesiredModelId) > 0 {
		input.DesiredModelId = aws.String(_bedrockDesiredModelId)
	}
	if len(_bedrockDesiredProvisionedModelName) > 0 {
		input.DesiredProvisionedModelName = aws.String(_bedrockDesiredProvisionedModelName)
	}

	if resp, err := client.UpdateProvisionedModelThroughput(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_bedrockCmd)
	_bedrockCmd.Flags().SortFlags = false

	_bedrockCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_bedrockCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_bedrockCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_bedrockCmd.Flags().StringVarP(&_bedrockAcceptEula, "accept-eula", "", "", "Accept Eula")
	_bedrockCmd.Flags().StringVarP(&_bedrockAnnotations, "annotations", "", "", "Annotations")
	_bedrockCmd.Flags().StringVarP(&_bedrockApplicationType, "application-type", "", "", "Application Type")
	_bedrockCmd.Flags().StringVarP(&_bedrockApplicationTypeEquals, "application-type-equals", "", "", "Application Type Equals")
	_bedrockCmd.Flags().StringVarP(&_bedrockAssetId, "asset-id", "", "", "Asset ID")
	_bedrockCmd.Flags().StringVarP(&_bedrockAssetType, "asset-type", "", "", "Asset Type")
	_bedrockCmd.Flags().StringVarP(&_bedrockAutomatedReasoningPolicyConfig, "automated-reasoning-policy-config", "", "", "Automated Reasoning Policy Config")
	_bedrockCmd.Flags().StringVarP(&_bedrockBaseModelArnEquals, "base-model-arn-equals", "", "", "Base Model ARN Equals")
	_bedrockCmd.Flags().StringVarP(&_bedrockBaseModelIdentifier, "base-model-identifier", "", "", "Base Model Identifier")
	_bedrockCmd.Flags().StringVarP(&_bedrockBlockedInputMessaging, "blocked-input-messaging", "", "", "Blocked Input Messaging")
	_bedrockCmd.Flags().StringVarP(&_bedrockBlockedOutputsMessaging, "blocked-outputs-messaging", "", "", "Blocked Outputs Messaging")
	_bedrockCmd.Flags().StringVarP(&_bedrockBuildWorkflowId, "build-workflow-id", "", "", "Build Workflow ID")
	_bedrockCmd.Flags().StringVarP(&_bedrockBuildWorkflowType, "build-workflow-type", "", "", "Build Workflow Type")
	_bedrockCmd.Flags().StringVarP(&_bedrockByCustomizationType, "by-customization-type", "", "", "By Customization Type")
	_bedrockCmd.Flags().StringVarP(&_bedrockByInferenceType, "by-inference-type", "", "", "By Inference Type")
	_bedrockCmd.Flags().StringVarP(&_bedrockByOutputModality, "by-output-modality", "", "", "By Output Modality")
	_bedrockCmd.Flags().StringVarP(&_bedrockByProvider, "by-provider", "", "", "By Provider")
	_bedrockCmd.Flags().StringVarP(&_bedrockClientRequestToken, "client-request-token", "", "", "Client Request Token")
	_bedrockCmd.Flags().StringVarP(&_bedrockCommitmentDuration, "commitment-duration", "", "", "Commitment Duration")
	_bedrockCmd.Flags().StringVarP(&_bedrockConfidenceThreshold, "confidence-threshold", "", "", "Confidence Threshold")
	_bedrockCmd.Flags().StringVarP(&_bedrockConfigId, "config-id", "", "", "Config ID")
	_bedrockCmd.Flags().StringVarP(&_bedrockContentPolicyConfig, "content-policy-config", "", "", "Content Policy Config")
	_bedrockCmd.Flags().StringVarP(&_bedrockContextualGroundingPolicyConfig, "contextual-grounding-policy-config", "", "", "Contextual Grounding Policy Config")
	_bedrockCmd.Flags().StringVarP(&_bedrockCreatedAfter, "created-after", "", "", "Created After")
	_bedrockCmd.Flags().StringVarP(&_bedrockCreatedBefore, "created-before", "", "", "Created Before")
	_bedrockCmd.Flags().StringVarP(&_bedrockCreationTimeAfter, "creation-time-after", "", "", "Creation Time After")
	_bedrockCmd.Flags().StringVarP(&_bedrockCreationTimeBefore, "creation-time-before", "", "", "Creation Time Before")
	_bedrockCmd.Flags().StringVarP(&_bedrockCrossRegionConfig, "cross-region-config", "", "", "Cross Region Config")
	_bedrockCmd.Flags().StringVarP(&_bedrockCustomModelDeploymentIdentifier, "custom-model-deployment-identifier", "", "", "Custom Model Deployment Identifier")
	_bedrockCmd.Flags().StringVarP(&_bedrockCustomModelKmsKeyId, "custom-model-kms-key-id", "", "", "Custom Model KMS Key ID")
	_bedrockCmd.Flags().StringVarP(&_bedrockCustomModelName, "custom-model-name", "", "", "Custom Model Name")
	_bedrockCmd.Flags().StringVarP(&_bedrockCustomModelTags, "custom-model-tags", "", "", "Custom Model Tags")
	_bedrockCmd.Flags().StringVarP(&_bedrockCustomerEncryptionKeyId, "customer-encryption-key-id", "", "", "Customer Encryption Key ID")
	_bedrockCmd.Flags().StringVarP(&_bedrockCustomizationConfig, "customization-config", "", "", "Customization Config")
	_bedrockCmd.Flags().StringVarP(&_bedrockCustomizationType, "customization-type", "", "", "Customization Type")
	_bedrockCmd.Flags().StringVarP(&_bedrockDescription, "description", "", "", "Description")
	_bedrockCmd.Flags().StringVarP(&_bedrockDesiredModelId, "desired-model-id", "", "", "Desired Model ID")
	_bedrockCmd.Flags().StringVarP(&_bedrockDesiredProvisionedModelName, "desired-provisioned-model-name", "", "", "Desired Provisioned Model Name")
	_bedrockCmd.Flags().StringVarP(&_bedrockEndpointArn, "endpoint-arn", "", "", "Endpoint ARN")
	_bedrockCmd.Flags().StringVarP(&_bedrockEndpointConfig, "endpoint-config", "", "", "Endpoint Config")
	_bedrockCmd.Flags().StringVarP(&_bedrockEndpointIdentifier, "endpoint-identifier", "", "", "Endpoint Identifier")
	_bedrockCmd.Flags().StringVarP(&_bedrockEndpointName, "endpoint-name", "", "", "Endpoint Name")
	_bedrockCmd.Flags().StringVarP(&_bedrockEvaluationConfig, "evaluation-config", "", "", "Evaluation Config")
	_bedrockCmd.Flags().StringVarP(&_bedrockExpectedAggregatedFindingsResult, "expected-aggregated-findings-result", "", "", "Expected Aggregated Findings Result")
	_bedrockCmd.Flags().StringVarP(&_bedrockFallbackModel, "fallback-model", "", "", "Fallback Model")
	_bedrockCmd.Flags().StringVarP(&_bedrockForce, "force", "", "", "Force")
	_bedrockCmd.Flags().StringVarP(&_bedrockFormData, "form-data", "", "", "Form Data")
	_bedrockCmd.Flags().StringVarP(&_bedrockFoundationModelArnEquals, "foundation-model-arn-equals", "", "", "Foundation Model ARN Equals")
	_bedrockCmd.Flags().StringVarP(&_bedrockGuardContent, "guard-content", "", "", "Guard Content")
	_bedrockCmd.Flags().StringVarP(&_bedrockGuardrailIdentifier, "guardrail-identifier", "", "", "Guardrail Identifier")
	_bedrockCmd.Flags().StringVarP(&_bedrockGuardrailInferenceConfig, "guardrail-inference-config", "", "", "Guardrail Inference Config")
	_bedrockCmd.Flags().StringVarP(&_bedrockGuardrailVersion, "guardrail-version", "", "", "Guardrail Version")
	_bedrockCmd.Flags().StringVarP(&_bedrockHyperParameters, "hyper-parameters", "", "", "Hyper Parameters")
	_bedrockCmd.Flags().StringVarP(&_bedrockImportedModelKmsKeyId, "imported-model-kms-key-id", "", "", "Imported Model KMS Key ID")
	_bedrockCmd.Flags().StringVarP(&_bedrockImportedModelName, "imported-model-name", "", "", "Imported Model Name")
	_bedrockCmd.Flags().StringVarP(&_bedrockImportedModelTags, "imported-model-tags", "", "", "Imported Model Tags")
	_bedrockCmd.Flags().StringVarP(&_bedrockInferenceConfig, "inference-config", "", "", "Inference Config")
	_bedrockCmd.Flags().StringVarP(&_bedrockInferenceProfileIdentifier, "inference-profile-identifier", "", "", "Inference Profile Identifier")
	_bedrockCmd.Flags().StringVarP(&_bedrockInferenceProfileName, "inference-profile-name", "", "", "Inference Profile Name")
	_bedrockCmd.Flags().StringVarP(&_bedrockInputDataConfig, "input-data-config", "", "", "Input Data Config")
	_bedrockCmd.Flags().StringVarP(&_bedrockIsOwned, "is-owned", "", "", "Is Owned")
	_bedrockCmd.Flags().StringVarP(&_bedrockJobArn, "job-arn", "", "", "Job ARN")
	_bedrockCmd.Flags().StringVarP(&_bedrockJobDescription, "job-description", "", "", "Job Description")
	_bedrockCmd.Flags().StringVarP(&_bedrockJobIdentifier, "job-identifier", "", "", "Job Identifier")
	_bedrockCmd.Flags().StringSliceVarP(&_bedrockJobIdentifiers, "job-identifiers", "", nil, "Job Identifiers")
	_bedrockCmd.Flags().StringVarP(&_bedrockJobName, "job-name", "", "", "Job Name")
	_bedrockCmd.Flags().StringVarP(&_bedrockJobTags, "job-tags", "", "", "Job Tags")
	_bedrockCmd.Flags().StringVarP(&_bedrockKmsKeyId, "kms-key-id", "", "", "KMS Key ID")
	_bedrockCmd.Flags().StringVarP(&_bedrockLastUpdatedAnnotationSetHash, "last-updated-annotation-set-hash", "", "", "Last Updated Annotation Set Hash")
	_bedrockCmd.Flags().StringVarP(&_bedrockLastUpdatedAt, "last-updated-at", "", "", "Last Updated At")
	_bedrockCmd.Flags().StringVarP(&_bedrockLastUpdatedDefinitionHash, "last-updated-definition-hash", "", "", "Last Updated Definition Hash")
	_bedrockCmd.Flags().StringVarP(&_bedrockLoggingConfig, "logging-config", "", "", "Logging Config")
	_bedrockCmd.Flags().StringVarP(&_bedrockMaxResults, "max-results", "", "", "Max Results")
	_bedrockCmd.Flags().StringVarP(&_bedrockModelArn, "model-arn", "", "", "Model ARN")
	_bedrockCmd.Flags().StringVarP(&_bedrockModelArnEquals, "model-arn-equals", "", "", "Model ARN Equals")
	_bedrockCmd.Flags().StringVarP(&_bedrockModelDataSource, "model-data-source", "", "", "Model Data Source")
	_bedrockCmd.Flags().StringVarP(&_bedrockModelDeploymentName, "model-deployment-name", "", "", "Model Deployment Name")
	_bedrockCmd.Flags().StringVarP(&_bedrockModelId, "model-id", "", "", "Model ID")
	_bedrockCmd.Flags().StringVarP(&_bedrockModelIdentifier, "model-identifier", "", "", "Model Identifier")
	_bedrockCmd.Flags().StringVarP(&_bedrockModelInvocationType, "model-invocation-type", "", "", "Model Invocation Type")
	_bedrockCmd.Flags().StringVarP(&_bedrockModelKmsKeyArn, "model-kms-key-arn", "", "", "Model KMS Key ARN")
	_bedrockCmd.Flags().StringVarP(&_bedrockModelKmsKeyId, "model-kms-key-id", "", "", "Model KMS Key ID")
	_bedrockCmd.Flags().StringVarP(&_bedrockModelName, "model-name", "", "", "Model Name")
	_bedrockCmd.Flags().StringVarP(&_bedrockModelSource, "model-source", "", "", "Model Source")
	_bedrockCmd.Flags().StringVarP(&_bedrockModelSourceConfig, "model-source-config", "", "", "Model Source Config")
	_bedrockCmd.Flags().StringVarP(&_bedrockModelSourceEquals, "model-source-equals", "", "", "Model Source Equals")
	_bedrockCmd.Flags().StringVarP(&_bedrockModelSourceIdentifier, "model-source-identifier", "", "", "Model Source Identifier")
	_bedrockCmd.Flags().StringVarP(&_bedrockModelStatus, "model-status", "", "", "Model Status")
	_bedrockCmd.Flags().StringVarP(&_bedrockModelTags, "model-tags", "", "", "Model Tags")
	_bedrockCmd.Flags().StringVarP(&_bedrockModelUnits, "model-units", "", "", "Model Units")
	_bedrockCmd.Flags().StringVarP(&_bedrockModels, "models", "", "", "Models")
	_bedrockCmd.Flags().StringVarP(&_bedrockName, "name", "", "", "Name")
	_bedrockCmd.Flags().StringVarP(&_bedrockNameContains, "name-contains", "", "", "Name Contains")
	_bedrockCmd.Flags().StringVarP(&_bedrockNextToken, "next-token", "", "", "Next Token")
	_bedrockCmd.Flags().StringVarP(&_bedrockOfferToken, "offer-token", "", "", "Offer Token")
	_bedrockCmd.Flags().StringVarP(&_bedrockOfferType, "offer-type", "", "", "Offer Type")
	_bedrockCmd.Flags().StringVarP(&_bedrockOutputDataConfig, "output-data-config", "", "", "Output Data Config")
	_bedrockCmd.Flags().StringVarP(&_bedrockPolicyArn, "policy-arn", "", "", "Policy ARN")
	_bedrockCmd.Flags().StringVarP(&_bedrockPolicyDefinition, "policy-definition", "", "", "Policy Definition")
	_bedrockCmd.Flags().StringVarP(&_bedrockPromptRouterArn, "prompt-router-arn", "", "", "Prompt Router ARN")
	_bedrockCmd.Flags().StringVarP(&_bedrockPromptRouterName, "prompt-router-name", "", "", "Prompt Router Name")
	_bedrockCmd.Flags().StringVarP(&_bedrockProvisionedModelId, "provisioned-model-id", "", "", "Provisioned Model ID")
	_bedrockCmd.Flags().StringVarP(&_bedrockProvisionedModelName, "provisioned-model-name", "", "", "Provisioned Model Name")
	_bedrockCmd.Flags().StringVarP(&_bedrockQueryContent, "query-content", "", "", "Query Content")
	_bedrockCmd.Flags().StringVarP(&_bedrockResourceARN, "resource-arn", "", "", "Resource ARN")
	_bedrockCmd.Flags().StringVarP(&_bedrockRoleArn, "role-arn", "", "", "Role ARN")
	_bedrockCmd.Flags().StringVarP(&_bedrockRoutingCriteria, "routing-criteria", "", "", "Routing Criteria")
	_bedrockCmd.Flags().StringVarP(&_bedrockSensitiveInformationPolicyConfig, "sensitive-information-policy-config", "", "", "Sensitive Information Policy Config")
	_bedrockCmd.Flags().StringVarP(&_bedrockSortBy, "sort-by", "", "", "Sort By")
	_bedrockCmd.Flags().StringVarP(&_bedrockSortOrder, "sort-order", "", "", "Sort Order")
	_bedrockCmd.Flags().StringVarP(&_bedrockSourceAccountEquals, "source-account-equals", "", "", "Source Account Equals")
	_bedrockCmd.Flags().StringVarP(&_bedrockSourceContent, "source-content", "", "", "Source Content")
	_bedrockCmd.Flags().StringVarP(&_bedrockSourceModelArn, "source-model-arn", "", "", "Source Model ARN")
	_bedrockCmd.Flags().StringVarP(&_bedrockSourceModelArnEquals, "source-model-arn-equals", "", "", "Source Model ARN Equals")
	_bedrockCmd.Flags().StringVarP(&_bedrockStatusEquals, "status-equals", "", "", "Status Equals")
	_bedrockCmd.Flags().StringVarP(&_bedrockSubmitTimeAfter, "submit-time-after", "", "", "Submit Time After")
	_bedrockCmd.Flags().StringVarP(&_bedrockSubmitTimeBefore, "submit-time-before", "", "", "Submit Time Before")
	_bedrockCmd.Flags().StringSliceVarP(&_bedrockTagKeys, "tag-keys", "", nil, "Tag Keys")
	_bedrockCmd.Flags().StringVarP(&_bedrockTags, "tags", "", "", "Tags")
	_bedrockCmd.Flags().StringVarP(&_bedrockTargetModelName, "target-model-name", "", "", "Target Model Name")
	_bedrockCmd.Flags().StringVarP(&_bedrockTargetModelNameContains, "target-model-name-contains", "", "", "Target Model Name Contains")
	_bedrockCmd.Flags().StringVarP(&_bedrockTargetModelTags, "target-model-tags", "", "", "Target Model Tags")
	_bedrockCmd.Flags().StringVarP(&_bedrockTestCaseId, "test-case-id", "", "", "Test Case ID")
	_bedrockCmd.Flags().StringSliceVarP(&_bedrockTestCaseIds, "test-case-ids", "", nil, "Test Case Ids")
	_bedrockCmd.Flags().StringVarP(&_bedrockTimeoutDurationInHours, "timeout-duration-in-hours", "", "", "Timeout Duration In Hours")
	_bedrockCmd.Flags().StringVarP(&_bedrockTopicPolicyConfig, "topic-policy-config", "", "", "Topic Policy Config")
	_bedrockCmd.Flags().StringVarP(&_bedrockTrainingDataConfig, "training-data-config", "", "", "Training Data Config")
	_bedrockCmd.Flags().StringVarP(&_bedrockType, "type", "", "", "Type")
	_bedrockCmd.Flags().StringVarP(&_bedrockTypeEquals, "type-equals", "", "", "Type Equals")
	_bedrockCmd.Flags().StringVarP(&_bedrockValidationDataConfig, "validation-data-config", "", "", "Validation Data Config")
	_bedrockCmd.Flags().StringVarP(&_bedrockVpcConfig, "vpc-config", "", "", "VPC Config")
	_bedrockCmd.Flags().StringVarP(&_bedrockWordPolicyConfig, "word-policy-config", "", "", "Word Policy Config")

	_bedrockCmd.Flags().BoolVarP(&_bedrockBatchDeleteEvaluationJob, "batch-delete-evaluation-job", "", false, "Batch Delete Evaluation Job")
	_bedrockCmd.Flags().BoolVarP(&_bedrockCancelAutomatedReasoningPolicyBuildWorkflow, "cancel-automated-reasoning-policy-build-workflow", "", false, "Cancel Automated Reasoning Policy Build Workflow")
	_bedrockCmd.Flags().BoolVarP(&_bedrockCreateAutomatedReasoningPolicy, "create-automated-reasoning-policy", "", false, "Create Automated Reasoning Policy")
	_bedrockCmd.Flags().BoolVarP(&_bedrockCreateAutomatedReasoningPolicyTestCase, "create-automated-reasoning-policy-test-case", "", false, "Create Automated Reasoning Policy Test Case")
	_bedrockCmd.Flags().BoolVarP(&_bedrockCreateAutomatedReasoningPolicyVersion, "create-automated-reasoning-policy-version", "", false, "Create Automated Reasoning Policy Version")
	_bedrockCmd.Flags().BoolVarP(&_bedrockCreateCustomModel, "create-custom-model", "", false, "Create Custom Model")
	_bedrockCmd.Flags().BoolVarP(&_bedrockCreateCustomModelDeployment, "create-custom-model-deployment", "", false, "Create Custom Model Deployment")
	_bedrockCmd.Flags().BoolVarP(&_bedrockCreateEvaluationJob, "create-evaluation-job", "", false, "Create Evaluation Job")
	_bedrockCmd.Flags().BoolVarP(&_bedrockCreateFoundationModelAgreement, "create-foundation-model-agreement", "", false, "Create Foundation Model Agreement")
	_bedrockCmd.Flags().BoolVarP(&_bedrockCreateGuardrail, "create-guardrail", "", false, "Create Guardrail")
	_bedrockCmd.Flags().BoolVarP(&_bedrockCreateGuardrailVersion, "create-guardrail-version", "", false, "Create Guardrail Version")
	_bedrockCmd.Flags().BoolVarP(&_bedrockCreateInferenceProfile, "create-inference-profile", "", false, "Create Inference Profile")
	_bedrockCmd.Flags().BoolVarP(&_bedrockCreateMarketplaceModelEndpoint, "create-marketplace-model-endpoint", "", false, "Create Marketplace Model Endpoint")
	_bedrockCmd.Flags().BoolVarP(&_bedrockCreateModelCopyJob, "create-model-copy-job", "", false, "Create Model Copy Job")
	_bedrockCmd.Flags().BoolVarP(&_bedrockCreateModelCustomizationJob, "create-model-customization-job", "", false, "Create Model Customization Job")
	_bedrockCmd.Flags().BoolVarP(&_bedrockCreateModelImportJob, "create-model-import-job", "", false, "Create Model Import Job")
	_bedrockCmd.Flags().BoolVarP(&_bedrockCreateModelInvocationJob, "create-model-invocation-job", "", false, "Create Model Invocation Job")
	_bedrockCmd.Flags().BoolVarP(&_bedrockCreatePromptRouter, "create-prompt-router", "", false, "Create Prompt Router")
	_bedrockCmd.Flags().BoolVarP(&_bedrockCreateProvisionedModelThroughput, "create-provisioned-model-throughput", "", false, "Create Provisioned Model Throughput")
	_bedrockCmd.Flags().BoolVarP(&_bedrockDeleteAutomatedReasoningPolicy, "delete-automated-reasoning-policy", "", false, "Delete Automated Reasoning Policy")
	_bedrockCmd.Flags().BoolVarP(&_bedrockDeleteAutomatedReasoningPolicyBuildWorkflow, "delete-automated-reasoning-policy-build-workflow", "", false, "Delete Automated Reasoning Policy Build Workflow")
	_bedrockCmd.Flags().BoolVarP(&_bedrockDeleteAutomatedReasoningPolicyTestCase, "delete-automated-reasoning-policy-test-case", "", false, "Delete Automated Reasoning Policy Test Case")
	_bedrockCmd.Flags().BoolVarP(&_bedrockDeleteCustomModel, "delete-custom-model", "", false, "Delete Custom Model")
	_bedrockCmd.Flags().BoolVarP(&_bedrockDeleteCustomModelDeployment, "delete-custom-model-deployment", "", false, "Delete Custom Model Deployment")
	_bedrockCmd.Flags().BoolVarP(&_bedrockDeleteEnforcedGuardrailConfiguration, "delete-enforced-guardrail-configuration", "", false, "Delete Enforced Guardrail Configuration")
	_bedrockCmd.Flags().BoolVarP(&_bedrockDeleteFoundationModelAgreement, "delete-foundation-model-agreement", "", false, "Delete Foundation Model Agreement")
	_bedrockCmd.Flags().BoolVarP(&_bedrockDeleteGuardrail, "delete-guardrail", "", false, "Delete Guardrail")
	_bedrockCmd.Flags().BoolVarP(&_bedrockDeleteImportedModel, "delete-imported-model", "", false, "Delete Imported Model")
	_bedrockCmd.Flags().BoolVarP(&_bedrockDeleteInferenceProfile, "delete-inference-profile", "", false, "Delete Inference Profile")
	_bedrockCmd.Flags().BoolVarP(&_bedrockDeleteMarketplaceModelEndpoint, "delete-marketplace-model-endpoint", "", false, "Delete Marketplace Model Endpoint")
	_bedrockCmd.Flags().BoolVarP(&_bedrockDeleteModelInvocationLoggingConfiguration, "delete-model-invocation-logging-configuration", "", false, "Delete Model Invocation Logging Configuration")
	_bedrockCmd.Flags().BoolVarP(&_bedrockDeletePromptRouter, "delete-prompt-router", "", false, "Delete Prompt Router")
	_bedrockCmd.Flags().BoolVarP(&_bedrockDeleteProvisionedModelThroughput, "delete-provisioned-model-throughput", "", false, "Delete Provisioned Model Throughput")
	_bedrockCmd.Flags().BoolVarP(&_bedrockDeregisterMarketplaceModelEndpoint, "deregister-marketplace-model-endpoint", "", false, "Deregister Marketplace Model Endpoint")
	_bedrockCmd.Flags().BoolVarP(&_bedrockExportAutomatedReasoningPolicyVersion, "export-automated-reasoning-policy-version", "", false, "Export Automated Reasoning Policy Version")
	_bedrockCmd.Flags().BoolVarP(&_bedrockGetAutomatedReasoningPolicy, "get-automated-reasoning-policy", "", false, "Get Automated Reasoning Policy")
	_bedrockCmd.Flags().BoolVarP(&_bedrockGetAutomatedReasoningPolicyAnnotations, "get-automated-reasoning-policy-annotations", "", false, "Get Automated Reasoning Policy Annotations")
	_bedrockCmd.Flags().BoolVarP(&_bedrockGetAutomatedReasoningPolicyBuildWorkflow, "get-automated-reasoning-policy-build-workflow", "", false, "Get Automated Reasoning Policy Build Workflow")
	_bedrockCmd.Flags().BoolVarP(&_bedrockGetAutomatedReasoningPolicyBuildWorkflowResultAssets, "get-automated-reasoning-policy-build-workflow-result-assets", "", false, "Get Automated Reasoning Policy Build Workflow Result Assets")
	_bedrockCmd.Flags().BoolVarP(&_bedrockGetAutomatedReasoningPolicyNextScenario, "get-automated-reasoning-policy-next-scenario", "", false, "Get Automated Reasoning Policy Next Scenario")
	_bedrockCmd.Flags().BoolVarP(&_bedrockGetAutomatedReasoningPolicyTestCase, "get-automated-reasoning-policy-test-case", "", false, "Get Automated Reasoning Policy Test Case")
	_bedrockCmd.Flags().BoolVarP(&_bedrockGetAutomatedReasoningPolicyTestResult, "get-automated-reasoning-policy-test-result", "", false, "Get Automated Reasoning Policy Test Result")
	_bedrockCmd.Flags().BoolVarP(&_bedrockGetCustomModel, "get-custom-model", "", false, "Get Custom Model")
	_bedrockCmd.Flags().BoolVarP(&_bedrockGetCustomModelDeployment, "get-custom-model-deployment", "", false, "Get Custom Model Deployment")
	_bedrockCmd.Flags().BoolVarP(&_bedrockGetEvaluationJob, "get-evaluation-job", "", false, "Get Evaluation Job")
	_bedrockCmd.Flags().BoolVarP(&_bedrockGetFoundationModel, "get-foundation-model", "", false, "Get Foundation Model")
	_bedrockCmd.Flags().BoolVarP(&_bedrockGetFoundationModelAvailability, "get-foundation-model-availability", "", false, "Get Foundation Model Availability")
	_bedrockCmd.Flags().BoolVarP(&_bedrockGetGuardrail, "get-guardrail", "", false, "Get Guardrail")
	_bedrockCmd.Flags().BoolVarP(&_bedrockGetImportedModel, "get-imported-model", "", false, "Get Imported Model")
	_bedrockCmd.Flags().BoolVarP(&_bedrockGetInferenceProfile, "get-inference-profile", "", false, "Get Inference Profile")
	_bedrockCmd.Flags().BoolVarP(&_bedrockGetMarketplaceModelEndpoint, "get-marketplace-model-endpoint", "", false, "Get Marketplace Model Endpoint")
	_bedrockCmd.Flags().BoolVarP(&_bedrockGetModelCopyJob, "get-model-copy-job", "", false, "Get Model Copy Job")
	_bedrockCmd.Flags().BoolVarP(&_bedrockGetModelCustomizationJob, "get-model-customization-job", "", false, "Get Model Customization Job")
	_bedrockCmd.Flags().BoolVarP(&_bedrockGetModelImportJob, "get-model-import-job", "", false, "Get Model Import Job")
	_bedrockCmd.Flags().BoolVarP(&_bedrockGetModelInvocationJob, "get-model-invocation-job", "", false, "Get Model Invocation Job")
	_bedrockCmd.Flags().BoolVarP(&_bedrockGetModelInvocationLoggingConfiguration, "get-model-invocation-logging-configuration", "", false, "Get Model Invocation Logging Configuration")
	_bedrockCmd.Flags().BoolVarP(&_bedrockGetPromptRouter, "get-prompt-router", "", false, "Get Prompt Router")
	_bedrockCmd.Flags().BoolVarP(&_bedrockGetProvisionedModelThroughput, "get-provisioned-model-throughput", "", false, "Get Provisioned Model Throughput")
	_bedrockCmd.Flags().BoolVarP(&_bedrockGetUseCaseForModelAccess, "get-use-case-for-model-access", "", false, "Get Use Case For Model Access")
	_bedrockCmd.Flags().BoolVarP(&_bedrockListAutomatedReasoningPolicies, "list-automated-reasoning-policies", "", false, "List Automated Reasoning Policies")
	_bedrockCmd.Flags().BoolVarP(&_bedrockListAutomatedReasoningPolicyBuildWorkflows, "list-automated-reasoning-policy-build-workflows", "", false, "List Automated Reasoning Policy Build Workflows")
	_bedrockCmd.Flags().BoolVarP(&_bedrockListAutomatedReasoningPolicyTestCases, "list-automated-reasoning-policy-test-cases", "", false, "List Automated Reasoning Policy Test Cases")
	_bedrockCmd.Flags().BoolVarP(&_bedrockListAutomatedReasoningPolicyTestResults, "list-automated-reasoning-policy-test-results", "", false, "List Automated Reasoning Policy Test Results")
	_bedrockCmd.Flags().BoolVarP(&_bedrockListCustomModelDeployments, "list-custom-model-deployments", "", false, "List Custom Model Deployments")
	_bedrockCmd.Flags().BoolVarP(&_bedrockListCustomModels, "list-custom-models", "", false, "List Custom Models")
	_bedrockCmd.Flags().BoolVarP(&_bedrockListEnforcedGuardrailsConfiguration, "list-enforced-guardrails-configuration", "", false, "List Enforced Guardrails Configuration")
	_bedrockCmd.Flags().BoolVarP(&_bedrockListEvaluationJobs, "list-evaluation-jobs", "", false, "List Evaluation Jobs")
	_bedrockCmd.Flags().BoolVarP(&_bedrockListFoundationModelAgreementOffers, "list-foundation-model-agreement-offers", "", false, "List Foundation Model Agreement Offers")
	_bedrockCmd.Flags().BoolVarP(&_bedrockListFoundationModels, "list-foundation-models", "", false, "List Foundation Models")
	_bedrockCmd.Flags().BoolVarP(&_bedrockListGuardrails, "list-guardrails", "", false, "List Guardrails")
	_bedrockCmd.Flags().BoolVarP(&_bedrockListImportedModels, "list-imported-models", "", false, "List Imported Models")
	_bedrockCmd.Flags().BoolVarP(&_bedrockListInferenceProfiles, "list-inference-profiles", "", false, "List Inference Profiles")
	_bedrockCmd.Flags().BoolVarP(&_bedrockListMarketplaceModelEndpoints, "list-marketplace-model-endpoints", "", false, "List Marketplace Model Endpoints")
	_bedrockCmd.Flags().BoolVarP(&_bedrockListModelCopyJobs, "list-model-copy-jobs", "", false, "List Model Copy Jobs")
	_bedrockCmd.Flags().BoolVarP(&_bedrockListModelCustomizationJobs, "list-model-customization-jobs", "", false, "List Model Customization Jobs")
	_bedrockCmd.Flags().BoolVarP(&_bedrockListModelImportJobs, "list-model-import-jobs", "", false, "List Model Import Jobs")
	_bedrockCmd.Flags().BoolVarP(&_bedrockListModelInvocationJobs, "list-model-invocation-jobs", "", false, "List Model Invocation Jobs")
	_bedrockCmd.Flags().BoolVarP(&_bedrockListPromptRouters, "list-prompt-routers", "", false, "List Prompt Routers")
	_bedrockCmd.Flags().BoolVarP(&_bedrockListProvisionedModelThroughputs, "list-provisioned-model-throughputs", "", false, "List Provisioned Model Throughputs")
	_bedrockCmd.Flags().BoolVarP(&_bedrockListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_bedrockCmd.Flags().BoolVarP(&_bedrockPutEnforcedGuardrailConfiguration, "put-enforced-guardrail-configuration", "", false, "Put Enforced Guardrail Configuration")
	_bedrockCmd.Flags().BoolVarP(&_bedrockPutModelInvocationLoggingConfiguration, "put-model-invocation-logging-configuration", "", false, "Put Model Invocation Logging Configuration")
	_bedrockCmd.Flags().BoolVarP(&_bedrockPutUseCaseForModelAccess, "put-use-case-for-model-access", "", false, "Put Use Case For Model Access")
	_bedrockCmd.Flags().BoolVarP(&_bedrockRegisterMarketplaceModelEndpoint, "register-marketplace-model-endpoint", "", false, "Register Marketplace Model Endpoint")
	_bedrockCmd.Flags().BoolVarP(&_bedrockStartAutomatedReasoningPolicyBuildWorkflow, "start-automated-reasoning-policy-build-workflow", "", false, "Start Automated Reasoning Policy Build Workflow")
	_bedrockCmd.Flags().BoolVarP(&_bedrockStartAutomatedReasoningPolicyTestWorkflow, "start-automated-reasoning-policy-test-workflow", "", false, "Start Automated Reasoning Policy Test Workflow")
	_bedrockCmd.Flags().BoolVarP(&_bedrockStopEvaluationJob, "stop-evaluation-job", "", false, "Stop Evaluation Job")
	_bedrockCmd.Flags().BoolVarP(&_bedrockStopModelCustomizationJob, "stop-model-customization-job", "", false, "Stop Model Customization Job")
	_bedrockCmd.Flags().BoolVarP(&_bedrockStopModelInvocationJob, "stop-model-invocation-job", "", false, "Stop Model Invocation Job")
	_bedrockCmd.Flags().BoolVarP(&_bedrockTagResource, "tag-resource", "", false, "Tag Resource")
	_bedrockCmd.Flags().BoolVarP(&_bedrockUntagResource, "untag-resource", "", false, "Untag Resource")
	_bedrockCmd.Flags().BoolVarP(&_bedrockUpdateAutomatedReasoningPolicy, "update-automated-reasoning-policy", "", false, "Update Automated Reasoning Policy")
	_bedrockCmd.Flags().BoolVarP(&_bedrockUpdateAutomatedReasoningPolicyAnnotations, "update-automated-reasoning-policy-annotations", "", false, "Update Automated Reasoning Policy Annotations")
	_bedrockCmd.Flags().BoolVarP(&_bedrockUpdateAutomatedReasoningPolicyTestCase, "update-automated-reasoning-policy-test-case", "", false, "Update Automated Reasoning Policy Test Case")
	_bedrockCmd.Flags().BoolVarP(&_bedrockUpdateCustomModelDeployment, "update-custom-model-deployment", "", false, "Update Custom Model Deployment")
	_bedrockCmd.Flags().BoolVarP(&_bedrockUpdateGuardrail, "update-guardrail", "", false, "Update Guardrail")
	_bedrockCmd.Flags().BoolVarP(&_bedrockUpdateMarketplaceModelEndpoint, "update-marketplace-model-endpoint", "", false, "Update Marketplace Model Endpoint")
	_bedrockCmd.Flags().BoolVarP(&_bedrockUpdateProvisionedModelThroughput, "update-provisioned-model-throughput", "", false, "Update Provisioned Model Throughput")

}
