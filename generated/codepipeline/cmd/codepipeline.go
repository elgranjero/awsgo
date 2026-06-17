package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// codepipelineCmd represents the codepipeline command
var _codepipelineCmd = &cobra.Command{
	Use:   "codepipeline",
	Short: "AWS codepipeline CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := codepipeline.NewFromConfig(cfg)
		if _codepipelineAcknowledgeJob {
			codepipeline_AcknowledgeJob(cfg, client)
			return
		}
		if _codepipelineAcknowledgeThirdPartyJob {
			codepipeline_AcknowledgeThirdPartyJob(cfg, client)
			return
		}
		if _codepipelineCreateCustomActionType {
			codepipeline_CreateCustomActionType(cfg, client)
			return
		}
		if _codepipelineCreatePipeline {
			codepipeline_CreatePipeline(cfg, client)
			return
		}
		if _codepipelineDeleteCustomActionType {
			codepipeline_DeleteCustomActionType(cfg, client)
			return
		}
		if _codepipelineDeletePipeline {
			codepipeline_DeletePipeline(cfg, client)
			return
		}
		if _codepipelineDeleteWebhook {
			codepipeline_DeleteWebhook(cfg, client)
			return
		}
		if _codepipelineDeregisterWebhookWithThirdParty {
			codepipeline_DeregisterWebhookWithThirdParty(cfg, client)
			return
		}
		if _codepipelineDisableStageTransition {
			codepipeline_DisableStageTransition(cfg, client)
			return
		}
		if _codepipelineEnableStageTransition {
			codepipeline_EnableStageTransition(cfg, client)
			return
		}
		if _codepipelineGetActionType {
			codepipeline_GetActionType(cfg, client)
			return
		}
		if _codepipelineGetJobDetails {
			codepipeline_GetJobDetails(cfg, client)
			return
		}
		if _codepipelineGetPipeline {
			codepipeline_GetPipeline(cfg, client)
			return
		}
		if _codepipelineGetPipelineExecution {
			codepipeline_GetPipelineExecution(cfg, client)
			return
		}
		if _codepipelineGetPipelineState {
			codepipeline_GetPipelineState(cfg, client)
			return
		}
		if _codepipelineGetThirdPartyJobDetails {
			codepipeline_GetThirdPartyJobDetails(cfg, client)
			return
		}
		if _codepipelineListActionExecutions {
			codepipeline_ListActionExecutions(cfg, client)
			return
		}
		if _codepipelineListActionTypes {
			codepipeline_ListActionTypes(cfg, client)
			return
		}
		if _codepipelineListDeployActionExecutionTargets {
			codepipeline_ListDeployActionExecutionTargets(cfg, client)
			return
		}
		if _codepipelineListPipelineExecutions {
			codepipeline_ListPipelineExecutions(cfg, client)
			return
		}
		if _codepipelineListPipelines {
			codepipeline_ListPipelines(cfg, client)
			return
		}
		if _codepipelineListRuleExecutions {
			codepipeline_ListRuleExecutions(cfg, client)
			return
		}
		if _codepipelineListRuleTypes {
			codepipeline_ListRuleTypes(cfg, client)
			return
		}
		if _codepipelineListTagsForResource {
			codepipeline_ListTagsForResource(cfg, client)
			return
		}
		if _codepipelineListWebhooks {
			codepipeline_ListWebhooks(cfg, client)
			return
		}
		if _codepipelineOverrideStageCondition {
			codepipeline_OverrideStageCondition(cfg, client)
			return
		}
		if _codepipelinePollForJobs {
			codepipeline_PollForJobs(cfg, client)
			return
		}
		if _codepipelinePollForThirdPartyJobs {
			codepipeline_PollForThirdPartyJobs(cfg, client)
			return
		}
		if _codepipelinePutActionRevision {
			codepipeline_PutActionRevision(cfg, client)
			return
		}
		if _codepipelinePutApprovalResult {
			codepipeline_PutApprovalResult(cfg, client)
			return
		}
		if _codepipelinePutJobFailureResult {
			codepipeline_PutJobFailureResult(cfg, client)
			return
		}
		if _codepipelinePutJobSuccessResult {
			codepipeline_PutJobSuccessResult(cfg, client)
			return
		}
		if _codepipelinePutThirdPartyJobFailureResult {
			codepipeline_PutThirdPartyJobFailureResult(cfg, client)
			return
		}
		if _codepipelinePutThirdPartyJobSuccessResult {
			codepipeline_PutThirdPartyJobSuccessResult(cfg, client)
			return
		}
		if _codepipelinePutWebhook {
			codepipeline_PutWebhook(cfg, client)
			return
		}
		if _codepipelineRegisterWebhookWithThirdParty {
			codepipeline_RegisterWebhookWithThirdParty(cfg, client)
			return
		}
		if _codepipelineRetryStageExecution {
			codepipeline_RetryStageExecution(cfg, client)
			return
		}
		if _codepipelineRollbackStage {
			codepipeline_RollbackStage(cfg, client)
			return
		}
		if _codepipelineStartPipelineExecution {
			codepipeline_StartPipelineExecution(cfg, client)
			return
		}
		if _codepipelineStopPipelineExecution {
			codepipeline_StopPipelineExecution(cfg, client)
			return
		}
		if _codepipelineTagResource {
			codepipeline_TagResource(cfg, client)
			return
		}
		if _codepipelineUntagResource {
			codepipeline_UntagResource(cfg, client)
			return
		}
		if _codepipelineUpdateActionType {
			codepipeline_UpdateActionType(cfg, client)
			return
		}
		if _codepipelineUpdatePipeline {
			codepipeline_UpdatePipeline(cfg, client)
			return
		}

	},
}

var (
	_codepipelineAcknowledgeJob                   bool
	_codepipelineAcknowledgeThirdPartyJob         bool
	_codepipelineCreateCustomActionType           bool
	_codepipelineCreatePipeline                   bool
	_codepipelineDeleteCustomActionType           bool
	_codepipelineDeletePipeline                   bool
	_codepipelineDeleteWebhook                    bool
	_codepipelineDeregisterWebhookWithThirdParty  bool
	_codepipelineDisableStageTransition           bool
	_codepipelineEnableStageTransition            bool
	_codepipelineGetActionType                    bool
	_codepipelineGetJobDetails                    bool
	_codepipelineGetPipeline                      bool
	_codepipelineGetPipelineExecution             bool
	_codepipelineGetPipelineState                 bool
	_codepipelineGetThirdPartyJobDetails          bool
	_codepipelineListActionExecutions             bool
	_codepipelineListActionTypes                  bool
	_codepipelineListDeployActionExecutionTargets bool
	_codepipelineListPipelineExecutions           bool
	_codepipelineListPipelines                    bool
	_codepipelineListRuleExecutions               bool
	_codepipelineListRuleTypes                    bool
	_codepipelineListTagsForResource              bool
	_codepipelineListWebhooks                     bool
	_codepipelineOverrideStageCondition           bool
	_codepipelinePollForJobs                      bool
	_codepipelinePollForThirdPartyJobs            bool
	_codepipelinePutActionRevision                bool
	_codepipelinePutApprovalResult                bool
	_codepipelinePutJobFailureResult              bool
	_codepipelinePutJobSuccessResult              bool
	_codepipelinePutThirdPartyJobFailureResult    bool
	_codepipelinePutThirdPartyJobSuccessResult    bool
	_codepipelinePutWebhook                       bool
	_codepipelineRegisterWebhookWithThirdParty    bool
	_codepipelineRetryStageExecution              bool
	_codepipelineRollbackStage                    bool
	_codepipelineStartPipelineExecution           bool
	_codepipelineStopPipelineExecution            bool
	_codepipelineTagResource                      bool
	_codepipelineUntagResource                    bool
	_codepipelineUpdateActionType                 bool
	_codepipelineUpdatePipeline                   bool

	_codepipelineAbandon                   string
	_codepipelineActionExecutionId         string
	_codepipelineActionName                string
	_codepipelineActionOwnerFilter         string
	_codepipelineActionRevision            string
	_codepipelineActionType                string
	_codepipelineActionTypeId              string
	_codepipelineCategory                  string
	_codepipelineClientRequestToken        string
	_codepipelineClientToken               string
	_codepipelineConditionType             string
	_codepipelineConfigurationProperties   string
	_codepipelineContinuationToken         string
	_codepipelineCurrentRevision           string
	_codepipelineExecutionDetails          string
	_codepipelineFailureDetails            string
	_codepipelineFilter                    string
	_codepipelineFilters                   string
	_codepipelineInputArtifactDetails      string
	_codepipelineJobId                     string
	_codepipelineMaxBatchSize              string
	_codepipelineMaxResults                string
	_codepipelineName                      string
	_codepipelineNextToken                 string
	_codepipelineNonce                     string
	_codepipelineOutputArtifactDetails     string
	_codepipelineOutputVariables           string
	_codepipelineOwner                     string
	_codepipelinePipeline                  string
	_codepipelinePipelineExecutionId       string
	_codepipelinePipelineName              string
	_codepipelineProvider                  string
	_codepipelineQueryParam                string
	_codepipelineReason                    string
	_codepipelineRegionFilter              string
	_codepipelineResourceArn               string
	_codepipelineResult                    string
	_codepipelineRetryMode                 string
	_codepipelineRuleOwnerFilter           string
	_codepipelineSettings                  string
	_codepipelineSourceRevisions           string
	_codepipelineStageName                 string
	_codepipelineTagKeys                   []string
	_codepipelineTags                      string
	_codepipelineTargetPipelineExecutionId string
	_codepipelineToken                     string
	_codepipelineTransitionType            string
	_codepipelineVariables                 string
	_codepipelineVersion                   string
	_codepipelineWebhook                   string
	_codepipelineWebhookName               string
)

// Returns information about a specified job and whether that job has been
// received by the job worker. Used for custom actions only.
func codepipeline_AcknowledgeJob(cfg aws.Config, client *codepipeline.Client) {
	input := &codepipeline.AcknowledgeJobInput{
		// JobId: *string, // Required
		// Nonce: *string, // Required
	}

	if len(_codepipelineJobId) > 0 {
		input.JobId = aws.String(_codepipelineJobId)
	}
	if len(_codepipelineNonce) > 0 {
		input.Nonce = aws.String(_codepipelineNonce)
	}

	if resp, err := client.AcknowledgeJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Confirms a job worker has received the specified job. Used for partner actions
// only.
func codepipeline_AcknowledgeThirdPartyJob(cfg aws.Config, client *codepipeline.Client) {
	input := &codepipeline.AcknowledgeThirdPartyJobInput{
		// ClientToken: *string, // Required
		// JobId: *string, // Required
		// Nonce: *string, // Required
	}

	if len(_codepipelineClientToken) > 0 {
		input.ClientToken = aws.String(_codepipelineClientToken)
	}
	if len(_codepipelineJobId) > 0 {
		input.JobId = aws.String(_codepipelineJobId)
	}
	if len(_codepipelineNonce) > 0 {
		input.Nonce = aws.String(_codepipelineNonce)
	}

	if resp, err := client.AcknowledgeThirdPartyJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new custom action that can be used in all pipelines associated with
// the Amazon Web Services account. Only used for custom actions.
func codepipeline_CreateCustomActionType(cfg aws.Config, client *codepipeline.Client) {
	input := &codepipeline.CreateCustomActionTypeInput{
		// Category: types.ActionCategory, // Required
		// InputArtifactDetails: *types.ArtifactDetails, // Required
		// OutputArtifactDetails: *types.ArtifactDetails, // Required
		// Provider: *string, // Required
		// Version: *string, // Required
	}

	if len(_codepipelineCategory) > 0 {
		if err := assignInputField(input, "Category", _codepipelineCategory); err != nil {
			log.Errorf("invalid --category: %s", err.Error())
			return
		}
	}
	if len(_codepipelineInputArtifactDetails) > 0 {
		if err := assignInputField(input, "InputArtifactDetails", _codepipelineInputArtifactDetails); err != nil {
			log.Errorf("invalid --input-artifact-details: %s", err.Error())
			return
		}
	}
	if len(_codepipelineOutputArtifactDetails) > 0 {
		if err := assignInputField(input, "OutputArtifactDetails", _codepipelineOutputArtifactDetails); err != nil {
			log.Errorf("invalid --output-artifact-details: %s", err.Error())
			return
		}
	}
	if len(_codepipelineProvider) > 0 {
		input.Provider = aws.String(_codepipelineProvider)
	}
	if len(_codepipelineVersion) > 0 {
		input.Version = aws.String(_codepipelineVersion)
	}
	if len(_codepipelineConfigurationProperties) > 0 {
		if err := assignInputField(input, "ConfigurationProperties", _codepipelineConfigurationProperties); err != nil {
			log.Errorf("invalid --configuration-properties: %s", err.Error())
			return
		}
	}
	if len(_codepipelineSettings) > 0 {
		if err := assignInputField(input, "Settings", _codepipelineSettings); err != nil {
			log.Errorf("invalid --settings: %s", err.Error())
			return
		}
	}
	if len(_codepipelineTags) > 0 {
		if err := assignInputField(input, "Tags", _codepipelineTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCustomActionType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a pipeline.
// In the pipeline structure, you must include either artifactStore or
// artifactStores in your pipeline, but you cannot use both. If you create a
// cross-region action in your pipeline, you must use artifactStores .
func codepipeline_CreatePipeline(cfg aws.Config, client *codepipeline.Client) {
	input := &codepipeline.CreatePipelineInput{
		// Pipeline: *types.PipelineDeclaration, // Required
	}

	if len(_codepipelinePipeline) > 0 {
		if err := assignInputField(input, "Pipeline", _codepipelinePipeline); err != nil {
			log.Errorf("invalid --pipeline: %s", err.Error())
			return
		}
	}
	if len(_codepipelineTags) > 0 {
		if err := assignInputField(input, "Tags", _codepipelineTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreatePipeline(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Marks a custom action as deleted. PollForJobs for the custom action fails after
// the action is marked for deletion. Used for custom actions only.
//
// To re-create a custom action after it has been deleted you must use a string in
// the version field that has never been used before. This string can be an
// incremented version number, for example. To restore a deleted custom action, use
// a JSON file that is identical to the deleted action, including the original
// string in the version field.
func codepipeline_DeleteCustomActionType(cfg aws.Config, client *codepipeline.Client) {
	input := &codepipeline.DeleteCustomActionTypeInput{
		// Category: types.ActionCategory, // Required
		// Provider: *string, // Required
		// Version: *string, // Required
	}

	if len(_codepipelineCategory) > 0 {
		if err := assignInputField(input, "Category", _codepipelineCategory); err != nil {
			log.Errorf("invalid --category: %s", err.Error())
			return
		}
	}
	if len(_codepipelineProvider) > 0 {
		input.Provider = aws.String(_codepipelineProvider)
	}
	if len(_codepipelineVersion) > 0 {
		input.Version = aws.String(_codepipelineVersion)
	}

	if resp, err := client.DeleteCustomActionType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified pipeline.
func codepipeline_DeletePipeline(cfg aws.Config, client *codepipeline.Client) {
	input := &codepipeline.DeletePipelineInput{
		// Name: *string, // Required
	}

	if len(_codepipelineName) > 0 {
		input.Name = aws.String(_codepipelineName)
	}

	if resp, err := client.DeletePipeline(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a previously created webhook by name. Deleting the webhook stops
// CodePipeline from starting a pipeline every time an external event occurs. The
// API returns successfully when trying to delete a webhook that is already
// deleted. If a deleted webhook is re-created by calling PutWebhook with the same
// name, it will have a different URL.
func codepipeline_DeleteWebhook(cfg aws.Config, client *codepipeline.Client) {
	input := &codepipeline.DeleteWebhookInput{
		// Name: *string, // Required
	}

	if len(_codepipelineName) > 0 {
		input.Name = aws.String(_codepipelineName)
	}

	if resp, err := client.DeleteWebhook(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the connection between the webhook that was created by CodePipeline and
// the external tool with events to be detected. Currently supported only for
// webhooks that target an action type of GitHub.
func codepipeline_DeregisterWebhookWithThirdParty(cfg aws.Config, client *codepipeline.Client) {
	input := &codepipeline.DeregisterWebhookWithThirdPartyInput{}

	if len(_codepipelineWebhookName) > 0 {
		input.WebhookName = aws.String(_codepipelineWebhookName)
	}

	if resp, err := client.DeregisterWebhookWithThirdParty(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Prevents artifacts in a pipeline from transitioning to the next stage in the
// pipeline.
func codepipeline_DisableStageTransition(cfg aws.Config, client *codepipeline.Client) {
	input := &codepipeline.DisableStageTransitionInput{
		// PipelineName: *string, // Required
		// Reason: *string, // Required
		// StageName: *string, // Required
		// TransitionType: types.StageTransitionType, // Required
	}

	if len(_codepipelinePipelineName) > 0 {
		input.PipelineName = aws.String(_codepipelinePipelineName)
	}
	if len(_codepipelineReason) > 0 {
		input.Reason = aws.String(_codepipelineReason)
	}
	if len(_codepipelineStageName) > 0 {
		input.StageName = aws.String(_codepipelineStageName)
	}
	if len(_codepipelineTransitionType) > 0 {
		if err := assignInputField(input, "TransitionType", _codepipelineTransitionType); err != nil {
			log.Errorf("invalid --transition-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.DisableStageTransition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables artifacts in a pipeline to transition to a stage in a pipeline.
func codepipeline_EnableStageTransition(cfg aws.Config, client *codepipeline.Client) {
	input := &codepipeline.EnableStageTransitionInput{
		// PipelineName: *string, // Required
		// StageName: *string, // Required
		// TransitionType: types.StageTransitionType, // Required
	}

	if len(_codepipelinePipelineName) > 0 {
		input.PipelineName = aws.String(_codepipelinePipelineName)
	}
	if len(_codepipelineStageName) > 0 {
		input.StageName = aws.String(_codepipelineStageName)
	}
	if len(_codepipelineTransitionType) > 0 {
		if err := assignInputField(input, "TransitionType", _codepipelineTransitionType); err != nil {
			log.Errorf("invalid --transition-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.EnableStageTransition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about an action type created for an external provider,
// where the action is to be used by customers of the external provider. The action
// can be created with any supported integration model.
func codepipeline_GetActionType(cfg aws.Config, client *codepipeline.Client) {
	input := &codepipeline.GetActionTypeInput{
		// Category: types.ActionCategory, // Required
		// Owner: *string, // Required
		// Provider: *string, // Required
		// Version: *string, // Required
	}

	if len(_codepipelineCategory) > 0 {
		if err := assignInputField(input, "Category", _codepipelineCategory); err != nil {
			log.Errorf("invalid --category: %s", err.Error())
			return
		}
	}
	if len(_codepipelineOwner) > 0 {
		input.Owner = aws.String(_codepipelineOwner)
	}
	if len(_codepipelineProvider) > 0 {
		input.Provider = aws.String(_codepipelineProvider)
	}
	if len(_codepipelineVersion) > 0 {
		input.Version = aws.String(_codepipelineVersion)
	}

	if resp, err := client.GetActionType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a job. Used for custom actions only.
// When this API is called, CodePipeline returns temporary credentials for the S3
// bucket used to store artifacts for the pipeline, if the action requires access
// to that S3 bucket for input or output artifacts. This API also returns any
// secret values defined for the action.
func codepipeline_GetJobDetails(cfg aws.Config, client *codepipeline.Client) {
	input := &codepipeline.GetJobDetailsInput{
		// JobId: *string, // Required
	}

	if len(_codepipelineJobId) > 0 {
		input.JobId = aws.String(_codepipelineJobId)
	}

	if resp, err := client.GetJobDetails(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the metadata, structure, stages, and actions of a pipeline. Can be used
// to return the entire structure of a pipeline in JSON format, which can then be
// modified and used to update the pipeline structure with UpdatePipeline.
func codepipeline_GetPipeline(cfg aws.Config, client *codepipeline.Client) {
	input := &codepipeline.GetPipelineInput{
		// Name: *string, // Required
	}

	if len(_codepipelineName) > 0 {
		input.Name = aws.String(_codepipelineName)
	}
	if len(_codepipelineVersion) > 0 {
		if err := assignInputField(input, "Version", _codepipelineVersion); err != nil {
			log.Errorf("invalid --version: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetPipeline(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about an execution of a pipeline, including details about
// artifacts, the pipeline execution ID, and the name, version, and status of the
// pipeline.
func codepipeline_GetPipelineExecution(cfg aws.Config, client *codepipeline.Client) {
	input := &codepipeline.GetPipelineExecutionInput{
		// PipelineExecutionId: *string, // Required
		// PipelineName: *string, // Required
	}

	if len(_codepipelinePipelineExecutionId) > 0 {
		input.PipelineExecutionId = aws.String(_codepipelinePipelineExecutionId)
	}
	if len(_codepipelinePipelineName) > 0 {
		input.PipelineName = aws.String(_codepipelinePipelineName)
	}

	if resp, err := client.GetPipelineExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the state of a pipeline, including the stages and
// actions.
//
// Values returned in the revisionId and revisionUrl fields indicate the source
// revision information, such as the commit ID, for the current state.
func codepipeline_GetPipelineState(cfg aws.Config, client *codepipeline.Client) {
	input := &codepipeline.GetPipelineStateInput{
		// Name: *string, // Required
	}

	if len(_codepipelineName) > 0 {
		input.Name = aws.String(_codepipelineName)
	}

	if resp, err := client.GetPipelineState(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Requests the details of a job for a third party action. Used for partner
// actions only.
//
// When this API is called, CodePipeline returns temporary credentials for the S3
// bucket used to store artifacts for the pipeline, if the action requires access
// to that S3 bucket for input or output artifacts. This API also returns any
// secret values defined for the action.
func codepipeline_GetThirdPartyJobDetails(cfg aws.Config, client *codepipeline.Client) {
	input := &codepipeline.GetThirdPartyJobDetailsInput{
		// ClientToken: *string, // Required
		// JobId: *string, // Required
	}

	if len(_codepipelineClientToken) > 0 {
		input.ClientToken = aws.String(_codepipelineClientToken)
	}
	if len(_codepipelineJobId) > 0 {
		input.JobId = aws.String(_codepipelineJobId)
	}

	if resp, err := client.GetThirdPartyJobDetails(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the action executions that have occurred in a pipeline.
func codepipeline_ListActionExecutions(cfg aws.Config, client *codepipeline.Client) {
	input := &codepipeline.ListActionExecutionsInput{
		// PipelineName: *string, // Required
	}

	if len(_codepipelinePipelineName) > 0 {
		input.PipelineName = aws.String(_codepipelinePipelineName)
	}
	if len(_codepipelineFilter) > 0 {
		if err := assignInputField(input, "Filter", _codepipelineFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_codepipelineMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codepipelineMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codepipelineNextToken) > 0 {
		input.NextToken = aws.String(_codepipelineNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListActionExecutions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codepipeline.ListActionExecutionsOutput
	p := codepipeline.NewListActionExecutionsPaginator(client, input)
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

// Gets a summary of all CodePipeline action types associated with your account.
func codepipeline_ListActionTypes(cfg aws.Config, client *codepipeline.Client) {
	input := &codepipeline.ListActionTypesInput{}

	if len(_codepipelineActionOwnerFilter) > 0 {
		if err := assignInputField(input, "ActionOwnerFilter", _codepipelineActionOwnerFilter); err != nil {
			log.Errorf("invalid --action-owner-filter: %s", err.Error())
			return
		}
	}
	if len(_codepipelineNextToken) > 0 {
		input.NextToken = aws.String(_codepipelineNextToken)
	}
	if len(_codepipelineRegionFilter) > 0 {
		input.RegionFilter = aws.String(_codepipelineRegionFilter)
	}

	if disablePaginator() {
		if resp, err := client.ListActionTypes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codepipeline.ListActionTypesOutput
	p := codepipeline.NewListActionTypesPaginator(client, input)
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

// Lists the targets for the deploy action.
func codepipeline_ListDeployActionExecutionTargets(cfg aws.Config, client *codepipeline.Client) {
	input := &codepipeline.ListDeployActionExecutionTargetsInput{
		// ActionExecutionId: *string, // Required
	}

	if len(_codepipelineActionExecutionId) > 0 {
		input.ActionExecutionId = aws.String(_codepipelineActionExecutionId)
	}
	if len(_codepipelineFilters) > 0 {
		if err := assignInputField(input, "Filters", _codepipelineFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_codepipelineMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codepipelineMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codepipelineNextToken) > 0 {
		input.NextToken = aws.String(_codepipelineNextToken)
	}
	if len(_codepipelinePipelineName) > 0 {
		input.PipelineName = aws.String(_codepipelinePipelineName)
	}

	if disablePaginator() {
		if resp, err := client.ListDeployActionExecutionTargets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codepipeline.ListDeployActionExecutionTargetsOutput
	p := codepipeline.NewListDeployActionExecutionTargetsPaginator(client, input)
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

// Gets a summary of the most recent executions for a pipeline.
// When applying the filter for pipeline executions that have succeeded in the
// stage, the operation returns all executions in the current pipeline version
// beginning on February 1, 2024.
func codepipeline_ListPipelineExecutions(cfg aws.Config, client *codepipeline.Client) {
	input := &codepipeline.ListPipelineExecutionsInput{
		// PipelineName: *string, // Required
	}

	if len(_codepipelinePipelineName) > 0 {
		input.PipelineName = aws.String(_codepipelinePipelineName)
	}
	if len(_codepipelineFilter) > 0 {
		if err := assignInputField(input, "Filter", _codepipelineFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_codepipelineMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codepipelineMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codepipelineNextToken) > 0 {
		input.NextToken = aws.String(_codepipelineNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPipelineExecutions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codepipeline.ListPipelineExecutionsOutput
	p := codepipeline.NewListPipelineExecutionsPaginator(client, input)
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

// Gets a summary of all of the pipelines associated with your account.
func codepipeline_ListPipelines(cfg aws.Config, client *codepipeline.Client) {
	input := &codepipeline.ListPipelinesInput{}

	if len(_codepipelineMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codepipelineMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codepipelineNextToken) > 0 {
		input.NextToken = aws.String(_codepipelineNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPipelines(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codepipeline.ListPipelinesOutput
	p := codepipeline.NewListPipelinesPaginator(client, input)
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

// Lists the rule executions that have occurred in a pipeline configured for
// conditions with rules.
func codepipeline_ListRuleExecutions(cfg aws.Config, client *codepipeline.Client) {
	input := &codepipeline.ListRuleExecutionsInput{
		// PipelineName: *string, // Required
	}

	if len(_codepipelinePipelineName) > 0 {
		input.PipelineName = aws.String(_codepipelinePipelineName)
	}
	if len(_codepipelineFilter) > 0 {
		if err := assignInputField(input, "Filter", _codepipelineFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_codepipelineMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codepipelineMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codepipelineNextToken) > 0 {
		input.NextToken = aws.String(_codepipelineNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRuleExecutions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codepipeline.ListRuleExecutionsOutput
	p := codepipeline.NewListRuleExecutionsPaginator(client, input)
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

// Lists the rules for the condition. For more information about conditions, see [Stage conditions]
// and [How do stage conditions work?].For more information about rules, see the [CodePipeline rule reference].
//
// [How do stage conditions work?]: https://docs.aws.amazon.com/codepipeline/latest/userguide/concepts-how-it-works-conditions.html
// [CodePipeline rule reference]: https://docs.aws.amazon.com/codepipeline/latest/userguide/rule-reference.html
// [Stage conditions]: https://docs.aws.amazon.com/codepipeline/latest/userguide/stage-conditions.html
func codepipeline_ListRuleTypes(cfg aws.Config, client *codepipeline.Client) {
	input := &codepipeline.ListRuleTypesInput{}

	if len(_codepipelineRegionFilter) > 0 {
		input.RegionFilter = aws.String(_codepipelineRegionFilter)
	}
	if len(_codepipelineRuleOwnerFilter) > 0 {
		if err := assignInputField(input, "RuleOwnerFilter", _codepipelineRuleOwnerFilter); err != nil {
			log.Errorf("invalid --rule-owner-filter: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListRuleTypes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the set of key-value pairs (metadata) that are used to manage the resource.
func codepipeline_ListTagsForResource(cfg aws.Config, client *codepipeline.Client) {
	input := &codepipeline.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_codepipelineResourceArn) > 0 {
		input.ResourceArn = aws.String(_codepipelineResourceArn)
	}
	if len(_codepipelineMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codepipelineMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codepipelineNextToken) > 0 {
		input.NextToken = aws.String(_codepipelineNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codepipeline.ListTagsForResourceOutput
	p := codepipeline.NewListTagsForResourcePaginator(client, input)
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

// Gets a listing of all the webhooks in this Amazon Web Services Region for this
// account. The output lists all webhooks and includes the webhook URL and ARN and
// the configuration for each webhook.
//
// If a secret token was provided, it will be redacted in the response.
func codepipeline_ListWebhooks(cfg aws.Config, client *codepipeline.Client) {
	input := &codepipeline.ListWebhooksInput{}

	if len(_codepipelineMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codepipelineMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codepipelineNextToken) > 0 {
		input.NextToken = aws.String(_codepipelineNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListWebhooks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codepipeline.ListWebhooksOutput
	p := codepipeline.NewListWebhooksPaginator(client, input)
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

// Used to override a stage condition. For more information about conditions, see [Stage conditions]
// and [How do stage conditions work?].
//
// [How do stage conditions work?]: https://docs.aws.amazon.com/codepipeline/latest/userguide/concepts-how-it-works-conditions.html
// [Stage conditions]: https://docs.aws.amazon.com/codepipeline/latest/userguide/stage-conditions.html
func codepipeline_OverrideStageCondition(cfg aws.Config, client *codepipeline.Client) {
	input := &codepipeline.OverrideStageConditionInput{
		// ConditionType: types.ConditionType, // Required
		// PipelineExecutionId: *string, // Required
		// PipelineName: *string, // Required
		// StageName: *string, // Required
	}

	if len(_codepipelineConditionType) > 0 {
		if err := assignInputField(input, "ConditionType", _codepipelineConditionType); err != nil {
			log.Errorf("invalid --condition-type: %s", err.Error())
			return
		}
	}
	if len(_codepipelinePipelineExecutionId) > 0 {
		input.PipelineExecutionId = aws.String(_codepipelinePipelineExecutionId)
	}
	if len(_codepipelinePipelineName) > 0 {
		input.PipelineName = aws.String(_codepipelinePipelineName)
	}
	if len(_codepipelineStageName) > 0 {
		input.StageName = aws.String(_codepipelineStageName)
	}

	if resp, err := client.OverrideStageCondition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about any jobs for CodePipeline to act on. PollForJobs is
// valid only for action types with "Custom" in the owner field. If the action type
// contains AWS or ThirdParty in the owner field, the PollForJobs action returns
// an error.
//
// When this API is called, CodePipeline returns temporary credentials for the S3
// bucket used to store artifacts for the pipeline, if the action requires access
// to that S3 bucket for input or output artifacts. This API also returns any
// secret values defined for the action.
func codepipeline_PollForJobs(cfg aws.Config, client *codepipeline.Client) {
	input := &codepipeline.PollForJobsInput{
		// ActionTypeId: *types.ActionTypeId, // Required
	}

	if len(_codepipelineActionTypeId) > 0 {
		if err := assignInputField(input, "ActionTypeId", _codepipelineActionTypeId); err != nil {
			log.Errorf("invalid --action-type-id: %s", err.Error())
			return
		}
	}
	if len(_codepipelineMaxBatchSize) > 0 {
		if err := assignInputField(input, "MaxBatchSize", _codepipelineMaxBatchSize); err != nil {
			log.Errorf("invalid --max-batch-size: %s", err.Error())
			return
		}
	}
	if len(_codepipelineQueryParam) > 0 {
		if err := assignInputField(input, "QueryParam", _codepipelineQueryParam); err != nil {
			log.Errorf("invalid --query-param: %s", err.Error())
			return
		}
	}

	if resp, err := client.PollForJobs(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Determines whether there are any third party jobs for a job worker to act on.
// Used for partner actions only.
//
// When this API is called, CodePipeline returns temporary credentials for the S3
// bucket used to store artifacts for the pipeline, if the action requires access
// to that S3 bucket for input or output artifacts.
func codepipeline_PollForThirdPartyJobs(cfg aws.Config, client *codepipeline.Client) {
	input := &codepipeline.PollForThirdPartyJobsInput{
		// ActionTypeId: *types.ActionTypeId, // Required
	}

	if len(_codepipelineActionTypeId) > 0 {
		if err := assignInputField(input, "ActionTypeId", _codepipelineActionTypeId); err != nil {
			log.Errorf("invalid --action-type-id: %s", err.Error())
			return
		}
	}
	if len(_codepipelineMaxBatchSize) > 0 {
		if err := assignInputField(input, "MaxBatchSize", _codepipelineMaxBatchSize); err != nil {
			log.Errorf("invalid --max-batch-size: %s", err.Error())
			return
		}
	}

	if resp, err := client.PollForThirdPartyJobs(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides information to CodePipeline about new revisions to a source.
func codepipeline_PutActionRevision(cfg aws.Config, client *codepipeline.Client) {
	input := &codepipeline.PutActionRevisionInput{
		// ActionName: *string, // Required
		// ActionRevision: *types.ActionRevision, // Required
		// PipelineName: *string, // Required
		// StageName: *string, // Required
	}

	if len(_codepipelineActionName) > 0 {
		input.ActionName = aws.String(_codepipelineActionName)
	}
	if len(_codepipelineActionRevision) > 0 {
		if err := assignInputField(input, "ActionRevision", _codepipelineActionRevision); err != nil {
			log.Errorf("invalid --action-revision: %s", err.Error())
			return
		}
	}
	if len(_codepipelinePipelineName) > 0 {
		input.PipelineName = aws.String(_codepipelinePipelineName)
	}
	if len(_codepipelineStageName) > 0 {
		input.StageName = aws.String(_codepipelineStageName)
	}

	if resp, err := client.PutActionRevision(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides the response to a manual approval request to CodePipeline. Valid
// responses include Approved and Rejected.
func codepipeline_PutApprovalResult(cfg aws.Config, client *codepipeline.Client) {
	input := &codepipeline.PutApprovalResultInput{
		// ActionName: *string, // Required
		// PipelineName: *string, // Required
		// Result: *types.ApprovalResult, // Required
		// StageName: *string, // Required
		// Token: *string, // Required
	}

	if len(_codepipelineActionName) > 0 {
		input.ActionName = aws.String(_codepipelineActionName)
	}
	if len(_codepipelinePipelineName) > 0 {
		input.PipelineName = aws.String(_codepipelinePipelineName)
	}
	if len(_codepipelineResult) > 0 {
		if err := assignInputField(input, "Result", _codepipelineResult); err != nil {
			log.Errorf("invalid --result: %s", err.Error())
			return
		}
	}
	if len(_codepipelineStageName) > 0 {
		input.StageName = aws.String(_codepipelineStageName)
	}
	if len(_codepipelineToken) > 0 {
		input.Token = aws.String(_codepipelineToken)
	}

	if resp, err := client.PutApprovalResult(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Represents the failure of a job as returned to the pipeline by a job worker.
// Used for custom actions only.
func codepipeline_PutJobFailureResult(cfg aws.Config, client *codepipeline.Client) {
	input := &codepipeline.PutJobFailureResultInput{
		// FailureDetails: *types.FailureDetails, // Required
		// JobId: *string, // Required
	}

	if len(_codepipelineFailureDetails) > 0 {
		if err := assignInputField(input, "FailureDetails", _codepipelineFailureDetails); err != nil {
			log.Errorf("invalid --failure-details: %s", err.Error())
			return
		}
	}
	if len(_codepipelineJobId) > 0 {
		input.JobId = aws.String(_codepipelineJobId)
	}

	if resp, err := client.PutJobFailureResult(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Represents the success of a job as returned to the pipeline by a job worker.
// Used for custom actions only.
func codepipeline_PutJobSuccessResult(cfg aws.Config, client *codepipeline.Client) {
	input := &codepipeline.PutJobSuccessResultInput{
		// JobId: *string, // Required
	}

	if len(_codepipelineJobId) > 0 {
		input.JobId = aws.String(_codepipelineJobId)
	}
	if len(_codepipelineContinuationToken) > 0 {
		input.ContinuationToken = aws.String(_codepipelineContinuationToken)
	}
	if len(_codepipelineCurrentRevision) > 0 {
		if err := assignInputField(input, "CurrentRevision", _codepipelineCurrentRevision); err != nil {
			log.Errorf("invalid --current-revision: %s", err.Error())
			return
		}
	}
	if len(_codepipelineExecutionDetails) > 0 {
		if err := assignInputField(input, "ExecutionDetails", _codepipelineExecutionDetails); err != nil {
			log.Errorf("invalid --execution-details: %s", err.Error())
			return
		}
	}
	if len(_codepipelineOutputVariables) > 0 {
		if err := assignInputField(input, "OutputVariables", _codepipelineOutputVariables); err != nil {
			log.Errorf("invalid --output-variables: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutJobSuccessResult(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Represents the failure of a third party job as returned to the pipeline by a
// job worker. Used for partner actions only.
func codepipeline_PutThirdPartyJobFailureResult(cfg aws.Config, client *codepipeline.Client) {
	input := &codepipeline.PutThirdPartyJobFailureResultInput{
		// ClientToken: *string, // Required
		// FailureDetails: *types.FailureDetails, // Required
		// JobId: *string, // Required
	}

	if len(_codepipelineClientToken) > 0 {
		input.ClientToken = aws.String(_codepipelineClientToken)
	}
	if len(_codepipelineFailureDetails) > 0 {
		if err := assignInputField(input, "FailureDetails", _codepipelineFailureDetails); err != nil {
			log.Errorf("invalid --failure-details: %s", err.Error())
			return
		}
	}
	if len(_codepipelineJobId) > 0 {
		input.JobId = aws.String(_codepipelineJobId)
	}

	if resp, err := client.PutThirdPartyJobFailureResult(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Represents the success of a third party job as returned to the pipeline by a
// job worker. Used for partner actions only.
func codepipeline_PutThirdPartyJobSuccessResult(cfg aws.Config, client *codepipeline.Client) {
	input := &codepipeline.PutThirdPartyJobSuccessResultInput{
		// ClientToken: *string, // Required
		// JobId: *string, // Required
	}

	if len(_codepipelineClientToken) > 0 {
		input.ClientToken = aws.String(_codepipelineClientToken)
	}
	if len(_codepipelineJobId) > 0 {
		input.JobId = aws.String(_codepipelineJobId)
	}
	if len(_codepipelineContinuationToken) > 0 {
		input.ContinuationToken = aws.String(_codepipelineContinuationToken)
	}
	if len(_codepipelineCurrentRevision) > 0 {
		if err := assignInputField(input, "CurrentRevision", _codepipelineCurrentRevision); err != nil {
			log.Errorf("invalid --current-revision: %s", err.Error())
			return
		}
	}
	if len(_codepipelineExecutionDetails) > 0 {
		if err := assignInputField(input, "ExecutionDetails", _codepipelineExecutionDetails); err != nil {
			log.Errorf("invalid --execution-details: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutThirdPartyJobSuccessResult(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Defines a webhook and returns a unique webhook URL generated by CodePipeline.
// This URL can be supplied to third party source hosting providers to call every
// time there's a code change. When CodePipeline receives a POST request on this
// URL, the pipeline defined in the webhook is started as long as the POST request
// satisfied the authentication and filtering requirements supplied when defining
// the webhook. RegisterWebhookWithThirdParty and DeregisterWebhookWithThirdParty
// APIs can be used to automatically configure supported third parties to call the
// generated webhook URL.
//
// When creating CodePipeline webhooks, do not use your own credentials or reuse
// the same secret token across multiple webhooks. For optimal security, generate a
// unique secret token for each webhook you create. The secret token is an
// arbitrary string that you provide, which GitHub uses to compute and sign the
// webhook payloads sent to CodePipeline, for protecting the integrity and
// authenticity of the webhook payloads. Using your own credentials or reusing the
// same token across multiple webhooks can lead to security vulnerabilities.
//
// If a secret token was provided, it will be redacted in the response.
func codepipeline_PutWebhook(cfg aws.Config, client *codepipeline.Client) {
	input := &codepipeline.PutWebhookInput{
		// Webhook: *types.WebhookDefinition, // Required
	}

	if len(_codepipelineWebhook) > 0 {
		if err := assignInputField(input, "Webhook", _codepipelineWebhook); err != nil {
			log.Errorf("invalid --webhook: %s", err.Error())
			return
		}
	}
	if len(_codepipelineTags) > 0 {
		if err := assignInputField(input, "Tags", _codepipelineTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutWebhook(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Configures a connection between the webhook that was created and the external
// tool with events to be detected.
func codepipeline_RegisterWebhookWithThirdParty(cfg aws.Config, client *codepipeline.Client) {
	input := &codepipeline.RegisterWebhookWithThirdPartyInput{}

	if len(_codepipelineWebhookName) > 0 {
		input.WebhookName = aws.String(_codepipelineWebhookName)
	}

	if resp, err := client.RegisterWebhookWithThirdParty(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// You can retry a stage that has failed without having to run a pipeline again
// from the beginning. You do this by either retrying the failed actions in a stage
// or by retrying all actions in the stage starting from the first action in the
// stage. When you retry the failed actions in a stage, all actions that are still
// in progress continue working, and failed actions are triggered again. When you
// retry a failed stage from the first action in the stage, the stage cannot have
// any actions in progress. Before a stage can be retried, it must either have all
// actions failed or some actions failed and some succeeded.
func codepipeline_RetryStageExecution(cfg aws.Config, client *codepipeline.Client) {
	input := &codepipeline.RetryStageExecutionInput{
		// PipelineExecutionId: *string, // Required
		// PipelineName: *string, // Required
		// RetryMode: types.StageRetryMode, // Required
		// StageName: *string, // Required
	}

	if len(_codepipelinePipelineExecutionId) > 0 {
		input.PipelineExecutionId = aws.String(_codepipelinePipelineExecutionId)
	}
	if len(_codepipelinePipelineName) > 0 {
		input.PipelineName = aws.String(_codepipelinePipelineName)
	}
	if len(_codepipelineRetryMode) > 0 {
		if err := assignInputField(input, "RetryMode", _codepipelineRetryMode); err != nil {
			log.Errorf("invalid --retry-mode: %s", err.Error())
			return
		}
	}
	if len(_codepipelineStageName) > 0 {
		input.StageName = aws.String(_codepipelineStageName)
	}

	if resp, err := client.RetryStageExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Rolls back a stage execution.
func codepipeline_RollbackStage(cfg aws.Config, client *codepipeline.Client) {
	input := &codepipeline.RollbackStageInput{
		// PipelineName: *string, // Required
		// StageName: *string, // Required
		// TargetPipelineExecutionId: *string, // Required
	}

	if len(_codepipelinePipelineName) > 0 {
		input.PipelineName = aws.String(_codepipelinePipelineName)
	}
	if len(_codepipelineStageName) > 0 {
		input.StageName = aws.String(_codepipelineStageName)
	}
	if len(_codepipelineTargetPipelineExecutionId) > 0 {
		input.TargetPipelineExecutionId = aws.String(_codepipelineTargetPipelineExecutionId)
	}

	if resp, err := client.RollbackStage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts the specified pipeline. Specifically, it begins processing the latest
// commit to the source location specified as part of the pipeline.
func codepipeline_StartPipelineExecution(cfg aws.Config, client *codepipeline.Client) {
	input := &codepipeline.StartPipelineExecutionInput{
		// Name: *string, // Required
	}

	if len(_codepipelineName) > 0 {
		input.Name = aws.String(_codepipelineName)
	}
	if len(_codepipelineClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_codepipelineClientRequestToken)
	}
	if len(_codepipelineSourceRevisions) > 0 {
		if err := assignInputField(input, "SourceRevisions", _codepipelineSourceRevisions); err != nil {
			log.Errorf("invalid --source-revisions: %s", err.Error())
			return
		}
	}
	if len(_codepipelineVariables) > 0 {
		if err := assignInputField(input, "Variables", _codepipelineVariables); err != nil {
			log.Errorf("invalid --variables: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartPipelineExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops the specified pipeline execution. You choose to either stop the pipeline
// execution by completing in-progress actions without starting subsequent actions,
// or by abandoning in-progress actions. While completing or abandoning in-progress
// actions, the pipeline execution is in a Stopping state. After all in-progress
// actions are completed or abandoned, the pipeline execution is in a Stopped
// state.
func codepipeline_StopPipelineExecution(cfg aws.Config, client *codepipeline.Client) {
	input := &codepipeline.StopPipelineExecutionInput{
		// PipelineExecutionId: *string, // Required
		// PipelineName: *string, // Required
	}

	if len(_codepipelinePipelineExecutionId) > 0 {
		input.PipelineExecutionId = aws.String(_codepipelinePipelineExecutionId)
	}
	if len(_codepipelinePipelineName) > 0 {
		input.PipelineName = aws.String(_codepipelinePipelineName)
	}
	if len(_codepipelineAbandon) > 0 {
		if err := assignInputField(input, "Abandon", _codepipelineAbandon); err != nil {
			log.Errorf("invalid --abandon: %s", err.Error())
			return
		}
	}
	if len(_codepipelineReason) > 0 {
		input.Reason = aws.String(_codepipelineReason)
	}

	if resp, err := client.StopPipelineExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds to or modifies the tags of the given resource. Tags are metadata that can
// be used to manage a resource.
func codepipeline_TagResource(cfg aws.Config, client *codepipeline.Client) {
	input := &codepipeline.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_codepipelineResourceArn) > 0 {
		input.ResourceArn = aws.String(_codepipelineResourceArn)
	}
	if len(_codepipelineTags) > 0 {
		if err := assignInputField(input, "Tags", _codepipelineTags); err != nil {
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

// Removes tags from an Amazon Web Services resource.
func codepipeline_UntagResource(cfg aws.Config, client *codepipeline.Client) {
	input := &codepipeline.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_codepipelineResourceArn) > 0 {
		input.ResourceArn = aws.String(_codepipelineResourceArn)
	}
	if len(_codepipelineTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _codepipelineTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an action type that was created with any supported integration model,
// where the action type is to be used by customers of the action type provider.
// Use a JSON file with the action definition and UpdateActionType to provide the
// full structure.
func codepipeline_UpdateActionType(cfg aws.Config, client *codepipeline.Client) {
	input := &codepipeline.UpdateActionTypeInput{
		// ActionType: *types.ActionTypeDeclaration, // Required
	}

	if len(_codepipelineActionType) > 0 {
		if err := assignInputField(input, "ActionType", _codepipelineActionType); err != nil {
			log.Errorf("invalid --action-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateActionType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a specified pipeline with edits or changes to its structure. Use a JSON
// file with the pipeline structure and UpdatePipeline to provide the full
// structure of the pipeline. Updating the pipeline increases the version number of
// the pipeline by 1.
func codepipeline_UpdatePipeline(cfg aws.Config, client *codepipeline.Client) {
	input := &codepipeline.UpdatePipelineInput{
		// Pipeline: *types.PipelineDeclaration, // Required
	}

	if len(_codepipelinePipeline) > 0 {
		if err := assignInputField(input, "Pipeline", _codepipelinePipeline); err != nil {
			log.Errorf("invalid --pipeline: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdatePipeline(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_codepipelineCmd)
	_codepipelineCmd.Flags().SortFlags = false

	_codepipelineCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_codepipelineCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_codepipelineCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_codepipelineCmd.Flags().StringVarP(&_codepipelineAbandon, "abandon", "", "", "Abandon")
	_codepipelineCmd.Flags().StringVarP(&_codepipelineActionExecutionId, "action-execution-id", "", "", "Action Execution ID")
	_codepipelineCmd.Flags().StringVarP(&_codepipelineActionName, "action-name", "", "", "Action Name")
	_codepipelineCmd.Flags().StringVarP(&_codepipelineActionOwnerFilter, "action-owner-filter", "", "", "Action Owner Filter")
	_codepipelineCmd.Flags().StringVarP(&_codepipelineActionRevision, "action-revision", "", "", "Action Revision")
	_codepipelineCmd.Flags().StringVarP(&_codepipelineActionType, "action-type", "", "", "Action Type")
	_codepipelineCmd.Flags().StringVarP(&_codepipelineActionTypeId, "action-type-id", "", "", "Action Type ID")
	_codepipelineCmd.Flags().StringVarP(&_codepipelineCategory, "category", "", "", "Category")
	_codepipelineCmd.Flags().StringVarP(&_codepipelineClientRequestToken, "client-request-token", "", "", "Client Request Token")
	_codepipelineCmd.Flags().StringVarP(&_codepipelineClientToken, "client-token", "", "", "Client Token")
	_codepipelineCmd.Flags().StringVarP(&_codepipelineConditionType, "condition-type", "", "", "Condition Type")
	_codepipelineCmd.Flags().StringVarP(&_codepipelineConfigurationProperties, "configuration-properties", "", "", "Configuration Properties")
	_codepipelineCmd.Flags().StringVarP(&_codepipelineContinuationToken, "continuation-token", "", "", "Continuation Token")
	_codepipelineCmd.Flags().StringVarP(&_codepipelineCurrentRevision, "current-revision", "", "", "Current Revision")
	_codepipelineCmd.Flags().StringVarP(&_codepipelineExecutionDetails, "execution-details", "", "", "Execution Details")
	_codepipelineCmd.Flags().StringVarP(&_codepipelineFailureDetails, "failure-details", "", "", "Failure Details")
	_codepipelineCmd.Flags().StringVarP(&_codepipelineFilter, "filter", "", "", "Filter")
	_codepipelineCmd.Flags().StringVarP(&_codepipelineFilters, "filters", "", "", "Filters")
	_codepipelineCmd.Flags().StringVarP(&_codepipelineInputArtifactDetails, "input-artifact-details", "", "", "Input Artifact Details")
	_codepipelineCmd.Flags().StringVarP(&_codepipelineJobId, "job-id", "", "", "Job ID")
	_codepipelineCmd.Flags().StringVarP(&_codepipelineMaxBatchSize, "max-batch-size", "", "", "Max Batch Size")
	_codepipelineCmd.Flags().StringVarP(&_codepipelineMaxResults, "max-results", "", "", "Max Results")
	_codepipelineCmd.Flags().StringVarP(&_codepipelineName, "name", "", "", "Name")
	_codepipelineCmd.Flags().StringVarP(&_codepipelineNextToken, "next-token", "", "", "Next Token")
	_codepipelineCmd.Flags().StringVarP(&_codepipelineNonce, "nonce", "", "", "Nonce")
	_codepipelineCmd.Flags().StringVarP(&_codepipelineOutputArtifactDetails, "output-artifact-details", "", "", "Output Artifact Details")
	_codepipelineCmd.Flags().StringVarP(&_codepipelineOutputVariables, "output-variables", "", "", "Output Variables")
	_codepipelineCmd.Flags().StringVarP(&_codepipelineOwner, "owner", "", "", "Owner")
	_codepipelineCmd.Flags().StringVarP(&_codepipelinePipeline, "pipeline", "", "", "Pipeline")
	_codepipelineCmd.Flags().StringVarP(&_codepipelinePipelineExecutionId, "pipeline-execution-id", "", "", "Pipeline Execution ID")
	_codepipelineCmd.Flags().StringVarP(&_codepipelinePipelineName, "pipeline-name", "", "", "Pipeline Name")
	_codepipelineCmd.Flags().StringVarP(&_codepipelineProvider, "provider", "", "", "Provider")
	_codepipelineCmd.Flags().StringVarP(&_codepipelineQueryParam, "query-param", "", "", "Query Param")
	_codepipelineCmd.Flags().StringVarP(&_codepipelineReason, "reason", "", "", "Reason")
	_codepipelineCmd.Flags().StringVarP(&_codepipelineRegionFilter, "region-filter", "", "", "Region Filter")
	_codepipelineCmd.Flags().StringVarP(&_codepipelineResourceArn, "resource-arn", "", "", "Resource ARN")
	_codepipelineCmd.Flags().StringVarP(&_codepipelineResult, "result", "", "", "Result")
	_codepipelineCmd.Flags().StringVarP(&_codepipelineRetryMode, "retry-mode", "", "", "Retry Mode")
	_codepipelineCmd.Flags().StringVarP(&_codepipelineRuleOwnerFilter, "rule-owner-filter", "", "", "Rule Owner Filter")
	_codepipelineCmd.Flags().StringVarP(&_codepipelineSettings, "settings", "", "", "Settings")
	_codepipelineCmd.Flags().StringVarP(&_codepipelineSourceRevisions, "source-revisions", "", "", "Source Revisions")
	_codepipelineCmd.Flags().StringVarP(&_codepipelineStageName, "stage-name", "", "", "Stage Name")
	_codepipelineCmd.Flags().StringSliceVarP(&_codepipelineTagKeys, "tag-keys", "", nil, "Tag Keys")
	_codepipelineCmd.Flags().StringVarP(&_codepipelineTags, "tags", "", "", "Tags")
	_codepipelineCmd.Flags().StringVarP(&_codepipelineTargetPipelineExecutionId, "target-pipeline-execution-id", "", "", "Target Pipeline Execution ID")
	_codepipelineCmd.Flags().StringVarP(&_codepipelineToken, "token", "", "", "Token")
	_codepipelineCmd.Flags().StringVarP(&_codepipelineTransitionType, "transition-type", "", "", "Transition Type")
	_codepipelineCmd.Flags().StringVarP(&_codepipelineVariables, "variables", "", "", "Variables")
	_codepipelineCmd.Flags().StringVarP(&_codepipelineVersion, "version", "", "", "Version")
	_codepipelineCmd.Flags().StringVarP(&_codepipelineWebhook, "webhook", "", "", "Webhook")
	_codepipelineCmd.Flags().StringVarP(&_codepipelineWebhookName, "webhook-name", "", "", "Webhook Name")

	_codepipelineCmd.Flags().BoolVarP(&_codepipelineAcknowledgeJob, "acknowledge-job", "", false, "Acknowledge Job")
	_codepipelineCmd.Flags().BoolVarP(&_codepipelineAcknowledgeThirdPartyJob, "acknowledge-third-party-job", "", false, "Acknowledge Third Party Job")
	_codepipelineCmd.Flags().BoolVarP(&_codepipelineCreateCustomActionType, "create-custom-action-type", "", false, "Create Custom Action Type")
	_codepipelineCmd.Flags().BoolVarP(&_codepipelineCreatePipeline, "create-pipeline", "", false, "Create Pipeline")
	_codepipelineCmd.Flags().BoolVarP(&_codepipelineDeleteCustomActionType, "delete-custom-action-type", "", false, "Delete Custom Action Type")
	_codepipelineCmd.Flags().BoolVarP(&_codepipelineDeletePipeline, "delete-pipeline", "", false, "Delete Pipeline")
	_codepipelineCmd.Flags().BoolVarP(&_codepipelineDeleteWebhook, "delete-webhook", "", false, "Delete Webhook")
	_codepipelineCmd.Flags().BoolVarP(&_codepipelineDeregisterWebhookWithThirdParty, "deregister-webhook-with-third-party", "", false, "Deregister Webhook With Third Party")
	_codepipelineCmd.Flags().BoolVarP(&_codepipelineDisableStageTransition, "disable-stage-transition", "", false, "Disable Stage Transition")
	_codepipelineCmd.Flags().BoolVarP(&_codepipelineEnableStageTransition, "enable-stage-transition", "", false, "Enable Stage Transition")
	_codepipelineCmd.Flags().BoolVarP(&_codepipelineGetActionType, "get-action-type", "", false, "Get Action Type")
	_codepipelineCmd.Flags().BoolVarP(&_codepipelineGetJobDetails, "get-job-details", "", false, "Get Job Details")
	_codepipelineCmd.Flags().BoolVarP(&_codepipelineGetPipeline, "get-pipeline", "", false, "Get Pipeline")
	_codepipelineCmd.Flags().BoolVarP(&_codepipelineGetPipelineExecution, "get-pipeline-execution", "", false, "Get Pipeline Execution")
	_codepipelineCmd.Flags().BoolVarP(&_codepipelineGetPipelineState, "get-pipeline-state", "", false, "Get Pipeline State")
	_codepipelineCmd.Flags().BoolVarP(&_codepipelineGetThirdPartyJobDetails, "get-third-party-job-details", "", false, "Get Third Party Job Details")
	_codepipelineCmd.Flags().BoolVarP(&_codepipelineListActionExecutions, "list-action-executions", "", false, "List Action Executions")
	_codepipelineCmd.Flags().BoolVarP(&_codepipelineListActionTypes, "list-action-types", "", false, "List Action Types")
	_codepipelineCmd.Flags().BoolVarP(&_codepipelineListDeployActionExecutionTargets, "list-deploy-action-execution-targets", "", false, "List Deploy Action Execution Targets")
	_codepipelineCmd.Flags().BoolVarP(&_codepipelineListPipelineExecutions, "list-pipeline-executions", "", false, "List Pipeline Executions")
	_codepipelineCmd.Flags().BoolVarP(&_codepipelineListPipelines, "list-pipelines", "", false, "List Pipelines")
	_codepipelineCmd.Flags().BoolVarP(&_codepipelineListRuleExecutions, "list-rule-executions", "", false, "List Rule Executions")
	_codepipelineCmd.Flags().BoolVarP(&_codepipelineListRuleTypes, "list-rule-types", "", false, "List Rule Types")
	_codepipelineCmd.Flags().BoolVarP(&_codepipelineListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_codepipelineCmd.Flags().BoolVarP(&_codepipelineListWebhooks, "list-webhooks", "", false, "List Webhooks")
	_codepipelineCmd.Flags().BoolVarP(&_codepipelineOverrideStageCondition, "override-stage-condition", "", false, "Override Stage Condition")
	_codepipelineCmd.Flags().BoolVarP(&_codepipelinePollForJobs, "poll-for-jobs", "", false, "Poll For Jobs")
	_codepipelineCmd.Flags().BoolVarP(&_codepipelinePollForThirdPartyJobs, "poll-for-third-party-jobs", "", false, "Poll For Third Party Jobs")
	_codepipelineCmd.Flags().BoolVarP(&_codepipelinePutActionRevision, "put-action-revision", "", false, "Put Action Revision")
	_codepipelineCmd.Flags().BoolVarP(&_codepipelinePutApprovalResult, "put-approval-result", "", false, "Put Approval Result")
	_codepipelineCmd.Flags().BoolVarP(&_codepipelinePutJobFailureResult, "put-job-failure-result", "", false, "Put Job Failure Result")
	_codepipelineCmd.Flags().BoolVarP(&_codepipelinePutJobSuccessResult, "put-job-success-result", "", false, "Put Job Success Result")
	_codepipelineCmd.Flags().BoolVarP(&_codepipelinePutThirdPartyJobFailureResult, "put-third-party-job-failure-result", "", false, "Put Third Party Job Failure Result")
	_codepipelineCmd.Flags().BoolVarP(&_codepipelinePutThirdPartyJobSuccessResult, "put-third-party-job-success-result", "", false, "Put Third Party Job Success Result")
	_codepipelineCmd.Flags().BoolVarP(&_codepipelinePutWebhook, "put-webhook", "", false, "Put Webhook")
	_codepipelineCmd.Flags().BoolVarP(&_codepipelineRegisterWebhookWithThirdParty, "register-webhook-with-third-party", "", false, "Register Webhook With Third Party")
	_codepipelineCmd.Flags().BoolVarP(&_codepipelineRetryStageExecution, "retry-stage-execution", "", false, "Retry Stage Execution")
	_codepipelineCmd.Flags().BoolVarP(&_codepipelineRollbackStage, "rollback-stage", "", false, "Rollback Stage")
	_codepipelineCmd.Flags().BoolVarP(&_codepipelineStartPipelineExecution, "start-pipeline-execution", "", false, "Start Pipeline Execution")
	_codepipelineCmd.Flags().BoolVarP(&_codepipelineStopPipelineExecution, "stop-pipeline-execution", "", false, "Stop Pipeline Execution")
	_codepipelineCmd.Flags().BoolVarP(&_codepipelineTagResource, "tag-resource", "", false, "Tag Resource")
	_codepipelineCmd.Flags().BoolVarP(&_codepipelineUntagResource, "untag-resource", "", false, "Untag Resource")
	_codepipelineCmd.Flags().BoolVarP(&_codepipelineUpdateActionType, "update-action-type", "", false, "Update Action Type")
	_codepipelineCmd.Flags().BoolVarP(&_codepipelineUpdatePipeline, "update-pipeline", "", false, "Update Pipeline")

}
