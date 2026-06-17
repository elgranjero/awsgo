package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// cloudformationCmd represents the cloudformation command
var _cloudformationCmd = &cobra.Command{
	Use:   "cloudformation",
	Short: "AWS cloudformation CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := cloudformation.NewFromConfig(cfg)
		if _cloudformationActivateOrganizationsAccess {
			cloudformation_ActivateOrganizationsAccess(cfg, client)
			return
		}
		if _cloudformationActivateType {
			cloudformation_ActivateType(cfg, client)
			return
		}
		if _cloudformationBatchDescribeTypeConfigurations {
			cloudformation_BatchDescribeTypeConfigurations(cfg, client)
			return
		}
		if _cloudformationCancelUpdateStack {
			cloudformation_CancelUpdateStack(cfg, client)
			return
		}
		if _cloudformationContinueUpdateRollback {
			cloudformation_ContinueUpdateRollback(cfg, client)
			return
		}
		if _cloudformationCreateChangeSet {
			cloudformation_CreateChangeSet(cfg, client)
			return
		}
		if _cloudformationCreateGeneratedTemplate {
			cloudformation_CreateGeneratedTemplate(cfg, client)
			return
		}
		if _cloudformationCreateStack {
			cloudformation_CreateStack(cfg, client)
			return
		}
		if _cloudformationCreateStackInstances {
			cloudformation_CreateStackInstances(cfg, client)
			return
		}
		if _cloudformationCreateStackRefactor {
			cloudformation_CreateStackRefactor(cfg, client)
			return
		}
		if _cloudformationCreateStackSet {
			cloudformation_CreateStackSet(cfg, client)
			return
		}
		if _cloudformationDeactivateOrganizationsAccess {
			cloudformation_DeactivateOrganizationsAccess(cfg, client)
			return
		}
		if _cloudformationDeactivateType {
			cloudformation_DeactivateType(cfg, client)
			return
		}
		if _cloudformationDeleteChangeSet {
			cloudformation_DeleteChangeSet(cfg, client)
			return
		}
		if _cloudformationDeleteGeneratedTemplate {
			cloudformation_DeleteGeneratedTemplate(cfg, client)
			return
		}
		if _cloudformationDeleteStack {
			cloudformation_DeleteStack(cfg, client)
			return
		}
		if _cloudformationDeleteStackInstances {
			cloudformation_DeleteStackInstances(cfg, client)
			return
		}
		if _cloudformationDeleteStackSet {
			cloudformation_DeleteStackSet(cfg, client)
			return
		}
		if _cloudformationDeregisterType {
			cloudformation_DeregisterType(cfg, client)
			return
		}
		if _cloudformationDescribeAccountLimits {
			cloudformation_DescribeAccountLimits(cfg, client)
			return
		}
		if _cloudformationDescribeChangeSet {
			cloudformation_DescribeChangeSet(cfg, client)
			return
		}
		if _cloudformationDescribeChangeSetHooks {
			cloudformation_DescribeChangeSetHooks(cfg, client)
			return
		}
		if _cloudformationDescribeEvents {
			cloudformation_DescribeEvents(cfg, client)
			return
		}
		if _cloudformationDescribeGeneratedTemplate {
			cloudformation_DescribeGeneratedTemplate(cfg, client)
			return
		}
		if _cloudformationDescribeOrganizationsAccess {
			cloudformation_DescribeOrganizationsAccess(cfg, client)
			return
		}
		if _cloudformationDescribePublisher {
			cloudformation_DescribePublisher(cfg, client)
			return
		}
		if _cloudformationDescribeResourceScan {
			cloudformation_DescribeResourceScan(cfg, client)
			return
		}
		if _cloudformationDescribeStackDriftDetectionStatus {
			cloudformation_DescribeStackDriftDetectionStatus(cfg, client)
			return
		}
		if _cloudformationDescribeStackEvents {
			cloudformation_DescribeStackEvents(cfg, client)
			return
		}
		if _cloudformationDescribeStackInstance {
			cloudformation_DescribeStackInstance(cfg, client)
			return
		}
		if _cloudformationDescribeStackRefactor {
			cloudformation_DescribeStackRefactor(cfg, client)
			return
		}
		if _cloudformationDescribeStackResource {
			cloudformation_DescribeStackResource(cfg, client)
			return
		}
		if _cloudformationDescribeStackResourceDrifts {
			cloudformation_DescribeStackResourceDrifts(cfg, client)
			return
		}
		if _cloudformationDescribeStackResources {
			cloudformation_DescribeStackResources(cfg, client)
			return
		}
		if _cloudformationDescribeStackSet {
			cloudformation_DescribeStackSet(cfg, client)
			return
		}
		if _cloudformationDescribeStackSetOperation {
			cloudformation_DescribeStackSetOperation(cfg, client)
			return
		}
		if _cloudformationDescribeStacks {
			cloudformation_DescribeStacks(cfg, client)
			return
		}
		if _cloudformationDescribeType {
			cloudformation_DescribeType(cfg, client)
			return
		}
		if _cloudformationDescribeTypeRegistration {
			cloudformation_DescribeTypeRegistration(cfg, client)
			return
		}
		if _cloudformationDetectStackDrift {
			cloudformation_DetectStackDrift(cfg, client)
			return
		}
		if _cloudformationDetectStackResourceDrift {
			cloudformation_DetectStackResourceDrift(cfg, client)
			return
		}
		if _cloudformationDetectStackSetDrift {
			cloudformation_DetectStackSetDrift(cfg, client)
			return
		}
		if _cloudformationEstimateTemplateCost {
			cloudformation_EstimateTemplateCost(cfg, client)
			return
		}
		if _cloudformationExecuteChangeSet {
			cloudformation_ExecuteChangeSet(cfg, client)
			return
		}
		if _cloudformationExecuteStackRefactor {
			cloudformation_ExecuteStackRefactor(cfg, client)
			return
		}
		if _cloudformationGetGeneratedTemplate {
			cloudformation_GetGeneratedTemplate(cfg, client)
			return
		}
		if _cloudformationGetHookResult {
			cloudformation_GetHookResult(cfg, client)
			return
		}
		if _cloudformationGetStackPolicy {
			cloudformation_GetStackPolicy(cfg, client)
			return
		}
		if _cloudformationGetTemplate {
			cloudformation_GetTemplate(cfg, client)
			return
		}
		if _cloudformationGetTemplateSummary {
			cloudformation_GetTemplateSummary(cfg, client)
			return
		}
		if _cloudformationImportStacksToStackSet {
			cloudformation_ImportStacksToStackSet(cfg, client)
			return
		}
		if _cloudformationListChangeSets {
			cloudformation_ListChangeSets(cfg, client)
			return
		}
		if _cloudformationListExports {
			cloudformation_ListExports(cfg, client)
			return
		}
		if _cloudformationListGeneratedTemplates {
			cloudformation_ListGeneratedTemplates(cfg, client)
			return
		}
		if _cloudformationListHookResults {
			cloudformation_ListHookResults(cfg, client)
			return
		}
		if _cloudformationListImports {
			cloudformation_ListImports(cfg, client)
			return
		}
		if _cloudformationListResourceScanRelatedResources {
			cloudformation_ListResourceScanRelatedResources(cfg, client)
			return
		}
		if _cloudformationListResourceScanResources {
			cloudformation_ListResourceScanResources(cfg, client)
			return
		}
		if _cloudformationListResourceScans {
			cloudformation_ListResourceScans(cfg, client)
			return
		}
		if _cloudformationListStackInstanceResourceDrifts {
			cloudformation_ListStackInstanceResourceDrifts(cfg, client)
			return
		}
		if _cloudformationListStackInstances {
			cloudformation_ListStackInstances(cfg, client)
			return
		}
		if _cloudformationListStackRefactorActions {
			cloudformation_ListStackRefactorActions(cfg, client)
			return
		}
		if _cloudformationListStackRefactors {
			cloudformation_ListStackRefactors(cfg, client)
			return
		}
		if _cloudformationListStackResources {
			cloudformation_ListStackResources(cfg, client)
			return
		}
		if _cloudformationListStackSetAutoDeploymentTargets {
			cloudformation_ListStackSetAutoDeploymentTargets(cfg, client)
			return
		}
		if _cloudformationListStackSetOperationResults {
			cloudformation_ListStackSetOperationResults(cfg, client)
			return
		}
		if _cloudformationListStackSetOperations {
			cloudformation_ListStackSetOperations(cfg, client)
			return
		}
		if _cloudformationListStackSets {
			cloudformation_ListStackSets(cfg, client)
			return
		}
		if _cloudformationListStacks {
			cloudformation_ListStacks(cfg, client)
			return
		}
		if _cloudformationListTypeRegistrations {
			cloudformation_ListTypeRegistrations(cfg, client)
			return
		}
		if _cloudformationListTypeVersions {
			cloudformation_ListTypeVersions(cfg, client)
			return
		}
		if _cloudformationListTypes {
			cloudformation_ListTypes(cfg, client)
			return
		}
		if _cloudformationPublishType {
			cloudformation_PublishType(cfg, client)
			return
		}
		if _cloudformationRecordHandlerProgress {
			cloudformation_RecordHandlerProgress(cfg, client)
			return
		}
		if _cloudformationRegisterPublisher {
			cloudformation_RegisterPublisher(cfg, client)
			return
		}
		if _cloudformationRegisterType {
			cloudformation_RegisterType(cfg, client)
			return
		}
		if _cloudformationRollbackStack {
			cloudformation_RollbackStack(cfg, client)
			return
		}
		if _cloudformationSetStackPolicy {
			cloudformation_SetStackPolicy(cfg, client)
			return
		}
		if _cloudformationSetTypeConfiguration {
			cloudformation_SetTypeConfiguration(cfg, client)
			return
		}
		if _cloudformationSetTypeDefaultVersion {
			cloudformation_SetTypeDefaultVersion(cfg, client)
			return
		}
		if _cloudformationSignalResource {
			cloudformation_SignalResource(cfg, client)
			return
		}
		if _cloudformationStartResourceScan {
			cloudformation_StartResourceScan(cfg, client)
			return
		}
		if _cloudformationStopStackSetOperation {
			cloudformation_StopStackSetOperation(cfg, client)
			return
		}
		if _cloudformationTestType {
			cloudformation_TestType(cfg, client)
			return
		}
		if _cloudformationUpdateGeneratedTemplate {
			cloudformation_UpdateGeneratedTemplate(cfg, client)
			return
		}
		if _cloudformationUpdateStack {
			cloudformation_UpdateStack(cfg, client)
			return
		}
		if _cloudformationUpdateStackInstances {
			cloudformation_UpdateStackInstances(cfg, client)
			return
		}
		if _cloudformationUpdateStackSet {
			cloudformation_UpdateStackSet(cfg, client)
			return
		}
		if _cloudformationUpdateTerminationProtection {
			cloudformation_UpdateTerminationProtection(cfg, client)
			return
		}
		if _cloudformationValidateTemplate {
			cloudformation_ValidateTemplate(cfg, client)
			return
		}

	},
}

var (
	_cloudformationActivateOrganizationsAccess       bool
	_cloudformationActivateType                      bool
	_cloudformationBatchDescribeTypeConfigurations   bool
	_cloudformationCancelUpdateStack                 bool
	_cloudformationContinueUpdateRollback            bool
	_cloudformationCreateChangeSet                   bool
	_cloudformationCreateGeneratedTemplate           bool
	_cloudformationCreateStack                       bool
	_cloudformationCreateStackInstances              bool
	_cloudformationCreateStackRefactor               bool
	_cloudformationCreateStackSet                    bool
	_cloudformationDeactivateOrganizationsAccess     bool
	_cloudformationDeactivateType                    bool
	_cloudformationDeleteChangeSet                   bool
	_cloudformationDeleteGeneratedTemplate           bool
	_cloudformationDeleteStack                       bool
	_cloudformationDeleteStackInstances              bool
	_cloudformationDeleteStackSet                    bool
	_cloudformationDeregisterType                    bool
	_cloudformationDescribeAccountLimits             bool
	_cloudformationDescribeChangeSet                 bool
	_cloudformationDescribeChangeSetHooks            bool
	_cloudformationDescribeEvents                    bool
	_cloudformationDescribeGeneratedTemplate         bool
	_cloudformationDescribeOrganizationsAccess       bool
	_cloudformationDescribePublisher                 bool
	_cloudformationDescribeResourceScan              bool
	_cloudformationDescribeStackDriftDetectionStatus bool
	_cloudformationDescribeStackEvents               bool
	_cloudformationDescribeStackInstance             bool
	_cloudformationDescribeStackRefactor             bool
	_cloudformationDescribeStackResource             bool
	_cloudformationDescribeStackResourceDrifts       bool
	_cloudformationDescribeStackResources            bool
	_cloudformationDescribeStackSet                  bool
	_cloudformationDescribeStackSetOperation         bool
	_cloudformationDescribeStacks                    bool
	_cloudformationDescribeType                      bool
	_cloudformationDescribeTypeRegistration          bool
	_cloudformationDetectStackDrift                  bool
	_cloudformationDetectStackResourceDrift          bool
	_cloudformationDetectStackSetDrift               bool
	_cloudformationEstimateTemplateCost              bool
	_cloudformationExecuteChangeSet                  bool
	_cloudformationExecuteStackRefactor              bool
	_cloudformationGetGeneratedTemplate              bool
	_cloudformationGetHookResult                     bool
	_cloudformationGetStackPolicy                    bool
	_cloudformationGetTemplate                       bool
	_cloudformationGetTemplateSummary                bool
	_cloudformationImportStacksToStackSet            bool
	_cloudformationListChangeSets                    bool
	_cloudformationListExports                       bool
	_cloudformationListGeneratedTemplates            bool
	_cloudformationListHookResults                   bool
	_cloudformationListImports                       bool
	_cloudformationListResourceScanRelatedResources  bool
	_cloudformationListResourceScanResources         bool
	_cloudformationListResourceScans                 bool
	_cloudformationListStackInstanceResourceDrifts   bool
	_cloudformationListStackInstances                bool
	_cloudformationListStackRefactorActions          bool
	_cloudformationListStackRefactors                bool
	_cloudformationListStackResources                bool
	_cloudformationListStackSetAutoDeploymentTargets bool
	_cloudformationListStackSetOperationResults      bool
	_cloudformationListStackSetOperations            bool
	_cloudformationListStackSets                     bool
	_cloudformationListStacks                        bool
	_cloudformationListTypeRegistrations             bool
	_cloudformationListTypeVersions                  bool
	_cloudformationListTypes                         bool
	_cloudformationPublishType                       bool
	_cloudformationRecordHandlerProgress             bool
	_cloudformationRegisterPublisher                 bool
	_cloudformationRegisterType                      bool
	_cloudformationRollbackStack                     bool
	_cloudformationSetStackPolicy                    bool
	_cloudformationSetTypeConfiguration              bool
	_cloudformationSetTypeDefaultVersion             bool
	_cloudformationSignalResource                    bool
	_cloudformationStartResourceScan                 bool
	_cloudformationStopStackSetOperation             bool
	_cloudformationTestType                          bool
	_cloudformationUpdateGeneratedTemplate           bool
	_cloudformationUpdateStack                       bool
	_cloudformationUpdateStackInstances              bool
	_cloudformationUpdateStackSet                    bool
	_cloudformationUpdateTerminationProtection       bool
	_cloudformationValidateTemplate                  bool

	_cloudformationAcceptTermsAndConditions           string
	_cloudformationAccounts                           []string
	_cloudformationAddResources                       string
	_cloudformationAdministrationRoleARN              string
	_cloudformationArn                                string
	_cloudformationAutoDeployment                     string
	_cloudformationAutoUpdate                         string
	_cloudformationBearerToken                        string
	_cloudformationCallAs                             string
	_cloudformationCapabilities                       string
	_cloudformationChangeSetName                      string
	_cloudformationChangeSetType                      string
	_cloudformationClientRequestToken                 string
	_cloudformationClientToken                        string
	_cloudformationConfiguration                      string
	_cloudformationConfigurationAlias                 string
	_cloudformationConnectionArn                      string
	_cloudformationCurrentOperationStatus             string
	_cloudformationDeletionMode                       string
	_cloudformationDeploymentMode                     string
	_cloudformationDeploymentTargets                  string
	_cloudformationDeprecatedStatus                   string
	_cloudformationDescription                        string
	_cloudformationDisableRollback                    string
	_cloudformationEnableStackCreation                string
	_cloudformationEnableTerminationProtection        string
	_cloudformationErrorCode                          string
	_cloudformationExecutionRoleArn                   string
	_cloudformationExecutionRoleName                  string
	_cloudformationExecutionStatusFilter              string
	_cloudformationExportName                         string
	_cloudformationFilters                            string
	_cloudformationFormat                             string
	_cloudformationGeneratedTemplateName              string
	_cloudformationHookResultId                       string
	_cloudformationImportExistingResources            string
	_cloudformationIncludeNestedStacks                string
	_cloudformationIncludePropertyValues              string
	_cloudformationLogDeliveryBucket                  string
	_cloudformationLoggingConfig                      string
	_cloudformationLogicalResourceId                  string
	_cloudformationLogicalResourceIds                 []string
	_cloudformationMajorVersion                       string
	_cloudformationManagedExecution                   string
	_cloudformationMaxResults                         string
	_cloudformationNewGeneratedTemplateName           string
	_cloudformationNextToken                          string
	_cloudformationNotificationARNs                   []string
	_cloudformationOnFailure                          string
	_cloudformationOnStackFailure                     string
	_cloudformationOperationId                        string
	_cloudformationOperationPreferences               string
	_cloudformationOperationStatus                    string
	_cloudformationOrganizationalUnitIds              []string
	_cloudformationParameterOverrides                 string
	_cloudformationParameters                         string
	_cloudformationPermissionModel                    string
	_cloudformationPhysicalResourceId                 string
	_cloudformationProvisioningType                   string
	_cloudformationPublicTypeArn                      string
	_cloudformationPublicVersionNumber                string
	_cloudformationPublisherId                        string
	_cloudformationRefreshAllResources                string
	_cloudformationRegions                            []string
	_cloudformationRegistrationStatusFilter           string
	_cloudformationRegistrationToken                  string
	_cloudformationRemoveResources                    []string
	_cloudformationResourceIdentifier                 string
	_cloudformationResourceMappings                   string
	_cloudformationResourceModel                      string
	_cloudformationResourceScanId                     string
	_cloudformationResourceTypePrefix                 string
	_cloudformationResourceTypes                      []string
	_cloudformationResources                          string
	_cloudformationResourcesToImport                  string
	_cloudformationResourcesToSkip                    []string
	_cloudformationRetainExceptOnCreate               string
	_cloudformationRetainResources                    []string
	_cloudformationRetainStacks                       string
	_cloudformationRoleARN                            string
	_cloudformationRollbackConfiguration              string
	_cloudformationScanFilters                        string
	_cloudformationScanTypeFilter                     string
	_cloudformationSchemaHandlerPackage               string
	_cloudformationStackDefinitions                   string
	_cloudformationStackDriftDetectionId              string
	_cloudformationStackId                            string
	_cloudformationStackIds                           []string
	_cloudformationStackIdsUrl                        string
	_cloudformationStackInstanceAccount               string
	_cloudformationStackInstanceRegion                string
	_cloudformationStackInstanceResourceDriftStatuses string
	_cloudformationStackName                          string
	_cloudformationStackPolicyBody                    string
	_cloudformationStackPolicyDuringUpdateBody        string
	_cloudformationStackPolicyDuringUpdateURL         string
	_cloudformationStackPolicyURL                     string
	_cloudformationStackRefactorId                    string
	_cloudformationStackResourceDriftStatusFilters    string
	_cloudformationStackSetName                       string
	_cloudformationStackStatusFilter                  string
	_cloudformationStatus                             string
	_cloudformationStatusMessage                      string
	_cloudformationTagKey                             string
	_cloudformationTagValue                           string
	_cloudformationTags                               string
	_cloudformationTargetId                           string
	_cloudformationTargetType                         string
	_cloudformationTemplateBody                       string
	_cloudformationTemplateConfiguration              string
	_cloudformationTemplateStage                      string
	_cloudformationTemplateSummaryConfig              string
	_cloudformationTemplateURL                        string
	_cloudformationTimeoutInMinutes                   string
	_cloudformationType                               string
	_cloudformationTypeArn                            string
	_cloudformationTypeConfigurationIdentifiers       string
	_cloudformationTypeName                           string
	_cloudformationTypeNameAlias                      string
	_cloudformationUniqueId                           string
	_cloudformationUsePreviousTemplate                string
	_cloudformationVersionBump                        string
	_cloudformationVersionId                          string
	_cloudformationVisibility                         string
)

// Activate trusted access with Organizations. With trusted access between
// StackSets and Organizations activated, the management account has permissions to
// create and manage StackSets for your organization.
func cloudformation_ActivateOrganizationsAccess(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.ActivateOrganizationsAccessInput{}

	if resp, err := client.ActivateOrganizationsAccess(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Activates a public third-party extension, such as a resource or module, to make
// it available for use in stack templates in your current account and Region. It
// can also create CloudFormation Hooks, which allow you to evaluate resource
// configurations before CloudFormation provisions them. Hooks integrate with both
// CloudFormation and Cloud Control API operations.
//
// After you activate an extension, you can use [SetTypeConfiguration] to set specific properties for
// the extension.
//
// To see which extensions have been activated, use [ListTypes]. To see configuration details
// for an extension, use [DescribeType].
//
// For more information, see [Activate a third-party public extension in your account] in the CloudFormation User Guide. For information
// about creating Hooks, see the [CloudFormation Hooks User Guide].
//
// [DescribeType]: https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_DescribeType.html
// [SetTypeConfiguration]: https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_SetTypeConfiguration.html
// [ListTypes]: https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_ListTypes.html
// [Activate a third-party public extension in your account]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/registry-public-activate-extension.html
// [CloudFormation Hooks User Guide]: https://docs.aws.amazon.com/cloudformation-cli/latest/hooks-userguide/what-is-cloudformation-hooks.html
func cloudformation_ActivateType(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.ActivateTypeInput{}

	if len(_cloudformationAutoUpdate) > 0 {
		if err := assignInputField(input, "AutoUpdate", _cloudformationAutoUpdate); err != nil {
			log.Errorf("invalid --auto-update: %s", err.Error())
			return
		}
	}
	if len(_cloudformationExecutionRoleArn) > 0 {
		input.ExecutionRoleArn = aws.String(_cloudformationExecutionRoleArn)
	}
	if len(_cloudformationLoggingConfig) > 0 {
		if err := assignInputField(input, "LoggingConfig", _cloudformationLoggingConfig); err != nil {
			log.Errorf("invalid --logging-config: %s", err.Error())
			return
		}
	}
	if len(_cloudformationMajorVersion) > 0 {
		if err := assignInputField(input, "MajorVersion", _cloudformationMajorVersion); err != nil {
			log.Errorf("invalid --major-version: %s", err.Error())
			return
		}
	}
	if len(_cloudformationPublicTypeArn) > 0 {
		input.PublicTypeArn = aws.String(_cloudformationPublicTypeArn)
	}
	if len(_cloudformationPublisherId) > 0 {
		input.PublisherId = aws.String(_cloudformationPublisherId)
	}
	if len(_cloudformationType) > 0 {
		if err := assignInputField(input, "Type", _cloudformationType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_cloudformationTypeName) > 0 {
		input.TypeName = aws.String(_cloudformationTypeName)
	}
	if len(_cloudformationTypeNameAlias) > 0 {
		input.TypeNameAlias = aws.String(_cloudformationTypeNameAlias)
	}
	if len(_cloudformationVersionBump) > 0 {
		if err := assignInputField(input, "VersionBump", _cloudformationVersionBump); err != nil {
			log.Errorf("invalid --version-bump: %s", err.Error())
			return
		}
	}

	if resp, err := client.ActivateType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns configuration data for the specified CloudFormation extensions, from
// the CloudFormation registry in your current account and Region.
//
// For more information, see [Edit configuration data for extensions in your account] in the CloudFormation User Guide.
//
// [Edit configuration data for extensions in your account]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/registry-set-configuration.html
func cloudformation_BatchDescribeTypeConfigurations(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.BatchDescribeTypeConfigurationsInput{
		// TypeConfigurationIdentifiers: []types.TypeConfigurationIdentifier, // Required
	}

	if len(_cloudformationTypeConfigurationIdentifiers) > 0 {
		if err := assignInputField(input, "TypeConfigurationIdentifiers", _cloudformationTypeConfigurationIdentifiers); err != nil {
			log.Errorf("invalid --type-configuration-identifiers: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchDescribeTypeConfigurations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels an update on the specified stack. If the call completes successfully,
// the stack rolls back the update and reverts to the previous stack configuration.
//
// You can cancel only stacks that are in the UPDATE_IN_PROGRESS state.
func cloudformation_CancelUpdateStack(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.CancelUpdateStackInput{
		// StackName: *string, // Required
	}

	if len(_cloudformationStackName) > 0 {
		input.StackName = aws.String(_cloudformationStackName)
	}
	if len(_cloudformationClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_cloudformationClientRequestToken)
	}

	if resp, err := client.CancelUpdateStack(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Continues rolling back a stack from UPDATE_ROLLBACK_FAILED to
// UPDATE_ROLLBACK_COMPLETE state. Depending on the cause of the failure, you can
// manually fix the error and continue the rollback. By continuing the rollback,
// you can return your stack to a working state (the UPDATE_ROLLBACK_COMPLETE
// state) and then try to update the stack again.
//
// A stack enters the UPDATE_ROLLBACK_FAILED state when CloudFormation can't roll
// back all changes after a failed stack update. For example, this might occur when
// a stack attempts to roll back to an old database that was deleted outside of
// CloudFormation. Because CloudFormation doesn't know the instance was deleted, it
// assumes the instance still exists and attempts to roll back to it, causing the
// update rollback to fail.
//
// For more information, see [Continue rolling back an update] in the CloudFormation User Guide. For information
// for troubleshooting a failed update rollback, see [Update rollback failed].
//
// [Continue rolling back an update]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-continueupdaterollback.html
// [Update rollback failed]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/troubleshooting.html#troubleshooting-errors-update-rollback-failed
func cloudformation_ContinueUpdateRollback(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.ContinueUpdateRollbackInput{
		// StackName: *string, // Required
	}

	if len(_cloudformationStackName) > 0 {
		input.StackName = aws.String(_cloudformationStackName)
	}
	if len(_cloudformationClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_cloudformationClientRequestToken)
	}
	if len(_cloudformationResourcesToSkip) > 0 {
		input.ResourcesToSkip = append([]string(nil), _cloudformationResourcesToSkip...)
	}
	if len(_cloudformationRoleARN) > 0 {
		input.RoleARN = aws.String(_cloudformationRoleARN)
	}

	if resp, err := client.ContinueUpdateRollback(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a list of changes that will be applied to a stack so that you can
// review the changes before executing them. You can create a change set for a
// stack that doesn't exist or an existing stack. If you create a change set for a
// stack that doesn't exist, the change set shows all of the resources that
// CloudFormation will create. If you create a change set for an existing stack,
// CloudFormation compares the stack's information with the information that you
// submit in the change set and lists the differences. Use change sets to
// understand which resources CloudFormation will create or change, and how it will
// change resources in an existing stack, before you create or update a stack.
//
// To create a change set for a stack that doesn't exist, for the ChangeSetType
// parameter, specify CREATE . To create a change set for an existing stack,
// specify UPDATE for the ChangeSetType parameter. To create a change set for an
// import operation, specify IMPORT for the ChangeSetType parameter. After the
// CreateChangeSet call successfully completes, CloudFormation starts creating the
// change set. To check the status of the change set or to review it, use the DescribeChangeSet
// action.
//
// When you are satisfied with the changes the change set will make, execute the
// change set by using the ExecuteChangeSetaction. CloudFormation doesn't make changes until you
// execute the change set.
//
// To create a change set for the entire stack hierarchy, set IncludeNestedStacks
// to True .
func cloudformation_CreateChangeSet(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.CreateChangeSetInput{
		// ChangeSetName: *string, // Required
		// StackName: *string, // Required
	}

	if len(_cloudformationChangeSetName) > 0 {
		input.ChangeSetName = aws.String(_cloudformationChangeSetName)
	}
	if len(_cloudformationStackName) > 0 {
		input.StackName = aws.String(_cloudformationStackName)
	}
	if len(_cloudformationCapabilities) > 0 {
		if err := assignInputField(input, "Capabilities", _cloudformationCapabilities); err != nil {
			log.Errorf("invalid --capabilities: %s", err.Error())
			return
		}
	}
	if len(_cloudformationChangeSetType) > 0 {
		if err := assignInputField(input, "ChangeSetType", _cloudformationChangeSetType); err != nil {
			log.Errorf("invalid --change-set-type: %s", err.Error())
			return
		}
	}
	if len(_cloudformationClientToken) > 0 {
		input.ClientToken = aws.String(_cloudformationClientToken)
	}
	if len(_cloudformationDeploymentMode) > 0 {
		if err := assignInputField(input, "DeploymentMode", _cloudformationDeploymentMode); err != nil {
			log.Errorf("invalid --deployment-mode: %s", err.Error())
			return
		}
	}
	if len(_cloudformationDescription) > 0 {
		input.Description = aws.String(_cloudformationDescription)
	}
	if len(_cloudformationImportExistingResources) > 0 {
		if err := assignInputField(input, "ImportExistingResources", _cloudformationImportExistingResources); err != nil {
			log.Errorf("invalid --import-existing-resources: %s", err.Error())
			return
		}
	}
	if len(_cloudformationIncludeNestedStacks) > 0 {
		if err := assignInputField(input, "IncludeNestedStacks", _cloudformationIncludeNestedStacks); err != nil {
			log.Errorf("invalid --include-nested-stacks: %s", err.Error())
			return
		}
	}
	if len(_cloudformationNotificationARNs) > 0 {
		input.NotificationARNs = append([]string(nil), _cloudformationNotificationARNs...)
	}
	if len(_cloudformationOnStackFailure) > 0 {
		if err := assignInputField(input, "OnStackFailure", _cloudformationOnStackFailure); err != nil {
			log.Errorf("invalid --on-stack-failure: %s", err.Error())
			return
		}
	}
	if len(_cloudformationParameters) > 0 {
		if err := assignInputField(input, "Parameters", _cloudformationParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}
	if len(_cloudformationResourceTypes) > 0 {
		input.ResourceTypes = append([]string(nil), _cloudformationResourceTypes...)
	}
	if len(_cloudformationResourcesToImport) > 0 {
		if err := assignInputField(input, "ResourcesToImport", _cloudformationResourcesToImport); err != nil {
			log.Errorf("invalid --resources-to-import: %s", err.Error())
			return
		}
	}
	if len(_cloudformationRoleARN) > 0 {
		input.RoleARN = aws.String(_cloudformationRoleARN)
	}
	if len(_cloudformationRollbackConfiguration) > 0 {
		if err := assignInputField(input, "RollbackConfiguration", _cloudformationRollbackConfiguration); err != nil {
			log.Errorf("invalid --rollback-configuration: %s", err.Error())
			return
		}
	}
	if len(_cloudformationTags) > 0 {
		if err := assignInputField(input, "Tags", _cloudformationTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_cloudformationTemplateBody) > 0 {
		input.TemplateBody = aws.String(_cloudformationTemplateBody)
	}
	if len(_cloudformationTemplateURL) > 0 {
		input.TemplateURL = aws.String(_cloudformationTemplateURL)
	}
	if len(_cloudformationUsePreviousTemplate) > 0 {
		if err := assignInputField(input, "UsePreviousTemplate", _cloudformationUsePreviousTemplate); err != nil {
			log.Errorf("invalid --use-previous-template: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateChangeSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a template from existing resources that are not already managed with
// CloudFormation. You can check the status of the template generation using the
// DescribeGeneratedTemplate API action.
func cloudformation_CreateGeneratedTemplate(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.CreateGeneratedTemplateInput{
		// GeneratedTemplateName: *string, // Required
	}

	if len(_cloudformationGeneratedTemplateName) > 0 {
		input.GeneratedTemplateName = aws.String(_cloudformationGeneratedTemplateName)
	}
	if len(_cloudformationResources) > 0 {
		if err := assignInputField(input, "Resources", _cloudformationResources); err != nil {
			log.Errorf("invalid --resources: %s", err.Error())
			return
		}
	}
	if len(_cloudformationStackName) > 0 {
		input.StackName = aws.String(_cloudformationStackName)
	}
	if len(_cloudformationTemplateConfiguration) > 0 {
		if err := assignInputField(input, "TemplateConfiguration", _cloudformationTemplateConfiguration); err != nil {
			log.Errorf("invalid --template-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateGeneratedTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a stack as specified in the template. After the call completes
// successfully, the stack creation starts. You can check the status of the stack
// through the DescribeStacksoperation.
//
// For more information about creating a stack and monitoring stack progress, see [Managing Amazon Web Services resources as a single unit with CloudFormation stacks]
// in the CloudFormation User Guide.
//
// [Managing Amazon Web Services resources as a single unit with CloudFormation stacks]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacks.html
func cloudformation_CreateStack(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.CreateStackInput{
		// StackName: *string, // Required
	}

	if len(_cloudformationStackName) > 0 {
		input.StackName = aws.String(_cloudformationStackName)
	}
	if len(_cloudformationCapabilities) > 0 {
		if err := assignInputField(input, "Capabilities", _cloudformationCapabilities); err != nil {
			log.Errorf("invalid --capabilities: %s", err.Error())
			return
		}
	}
	if len(_cloudformationClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_cloudformationClientRequestToken)
	}
	if len(_cloudformationDisableRollback) > 0 {
		if err := assignInputField(input, "DisableRollback", _cloudformationDisableRollback); err != nil {
			log.Errorf("invalid --disable-rollback: %s", err.Error())
			return
		}
	}
	if len(_cloudformationEnableTerminationProtection) > 0 {
		if err := assignInputField(input, "EnableTerminationProtection", _cloudformationEnableTerminationProtection); err != nil {
			log.Errorf("invalid --enable-termination-protection: %s", err.Error())
			return
		}
	}
	if len(_cloudformationNotificationARNs) > 0 {
		input.NotificationARNs = append([]string(nil), _cloudformationNotificationARNs...)
	}
	if len(_cloudformationOnFailure) > 0 {
		if err := assignInputField(input, "OnFailure", _cloudformationOnFailure); err != nil {
			log.Errorf("invalid --on-failure: %s", err.Error())
			return
		}
	}
	if len(_cloudformationParameters) > 0 {
		if err := assignInputField(input, "Parameters", _cloudformationParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}
	if len(_cloudformationResourceTypes) > 0 {
		input.ResourceTypes = append([]string(nil), _cloudformationResourceTypes...)
	}
	if len(_cloudformationRetainExceptOnCreate) > 0 {
		if err := assignInputField(input, "RetainExceptOnCreate", _cloudformationRetainExceptOnCreate); err != nil {
			log.Errorf("invalid --retain-except-on-create: %s", err.Error())
			return
		}
	}
	if len(_cloudformationRoleARN) > 0 {
		input.RoleARN = aws.String(_cloudformationRoleARN)
	}
	if len(_cloudformationRollbackConfiguration) > 0 {
		if err := assignInputField(input, "RollbackConfiguration", _cloudformationRollbackConfiguration); err != nil {
			log.Errorf("invalid --rollback-configuration: %s", err.Error())
			return
		}
	}
	if len(_cloudformationStackPolicyBody) > 0 {
		input.StackPolicyBody = aws.String(_cloudformationStackPolicyBody)
	}
	if len(_cloudformationStackPolicyURL) > 0 {
		input.StackPolicyURL = aws.String(_cloudformationStackPolicyURL)
	}
	if len(_cloudformationTags) > 0 {
		if err := assignInputField(input, "Tags", _cloudformationTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_cloudformationTemplateBody) > 0 {
		input.TemplateBody = aws.String(_cloudformationTemplateBody)
	}
	if len(_cloudformationTemplateURL) > 0 {
		input.TemplateURL = aws.String(_cloudformationTemplateURL)
	}
	if len(_cloudformationTimeoutInMinutes) > 0 {
		if err := assignInputField(input, "TimeoutInMinutes", _cloudformationTimeoutInMinutes); err != nil {
			log.Errorf("invalid --timeout-in-minutes: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateStack(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates stack instances for the specified accounts, within the specified Amazon
// Web Services Regions. A stack instance refers to a stack in a specific account
// and Region. You must specify at least one value for either Accounts or
// DeploymentTargets , and you must specify at least one value for Regions .
//
// The maximum number of organizational unit (OUs) supported by a
// CreateStackInstances operation is 50.
//
// If you need more than 50, consider the following options:
//
// - Batch processing: If you don't want to expose your OU hierarchy, split up
// the operations into multiple calls with less than 50 OUs each.
//
// - Parent OU strategy: If you don't mind exposing the OU hierarchy, target a
// parent OU that contains all desired child OUs.
func cloudformation_CreateStackInstances(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.CreateStackInstancesInput{
		// Regions: []string, // Required
		// StackSetName: *string, // Required
	}

	if len(_cloudformationRegions) > 0 {
		input.Regions = append([]string(nil), _cloudformationRegions...)
	}
	if len(_cloudformationStackSetName) > 0 {
		input.StackSetName = aws.String(_cloudformationStackSetName)
	}
	if len(_cloudformationAccounts) > 0 {
		input.Accounts = append([]string(nil), _cloudformationAccounts...)
	}
	if len(_cloudformationCallAs) > 0 {
		if err := assignInputField(input, "CallAs", _cloudformationCallAs); err != nil {
			log.Errorf("invalid --call-as: %s", err.Error())
			return
		}
	}
	if len(_cloudformationDeploymentTargets) > 0 {
		if err := assignInputField(input, "DeploymentTargets", _cloudformationDeploymentTargets); err != nil {
			log.Errorf("invalid --deployment-targets: %s", err.Error())
			return
		}
	}
	if len(_cloudformationOperationId) > 0 {
		input.OperationId = aws.String(_cloudformationOperationId)
	}
	if len(_cloudformationOperationPreferences) > 0 {
		if err := assignInputField(input, "OperationPreferences", _cloudformationOperationPreferences); err != nil {
			log.Errorf("invalid --operation-preferences: %s", err.Error())
			return
		}
	}
	if len(_cloudformationParameterOverrides) > 0 {
		if err := assignInputField(input, "ParameterOverrides", _cloudformationParameterOverrides); err != nil {
			log.Errorf("invalid --parameter-overrides: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateStackInstances(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a refactor across multiple stacks, with the list of stacks and
// resources that are affected.
func cloudformation_CreateStackRefactor(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.CreateStackRefactorInput{
		// StackDefinitions: []types.StackDefinition, // Required
	}

	if len(_cloudformationStackDefinitions) > 0 {
		if err := assignInputField(input, "StackDefinitions", _cloudformationStackDefinitions); err != nil {
			log.Errorf("invalid --stack-definitions: %s", err.Error())
			return
		}
	}
	if len(_cloudformationDescription) > 0 {
		input.Description = aws.String(_cloudformationDescription)
	}
	if len(_cloudformationEnableStackCreation) > 0 {
		if err := assignInputField(input, "EnableStackCreation", _cloudformationEnableStackCreation); err != nil {
			log.Errorf("invalid --enable-stack-creation: %s", err.Error())
			return
		}
	}
	if len(_cloudformationResourceMappings) > 0 {
		if err := assignInputField(input, "ResourceMappings", _cloudformationResourceMappings); err != nil {
			log.Errorf("invalid --resource-mappings: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateStackRefactor(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a StackSet.
func cloudformation_CreateStackSet(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.CreateStackSetInput{
		// StackSetName: *string, // Required
	}

	if len(_cloudformationStackSetName) > 0 {
		input.StackSetName = aws.String(_cloudformationStackSetName)
	}
	if len(_cloudformationAdministrationRoleARN) > 0 {
		input.AdministrationRoleARN = aws.String(_cloudformationAdministrationRoleARN)
	}
	if len(_cloudformationAutoDeployment) > 0 {
		if err := assignInputField(input, "AutoDeployment", _cloudformationAutoDeployment); err != nil {
			log.Errorf("invalid --auto-deployment: %s", err.Error())
			return
		}
	}
	if len(_cloudformationCallAs) > 0 {
		if err := assignInputField(input, "CallAs", _cloudformationCallAs); err != nil {
			log.Errorf("invalid --call-as: %s", err.Error())
			return
		}
	}
	if len(_cloudformationCapabilities) > 0 {
		if err := assignInputField(input, "Capabilities", _cloudformationCapabilities); err != nil {
			log.Errorf("invalid --capabilities: %s", err.Error())
			return
		}
	}
	if len(_cloudformationClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_cloudformationClientRequestToken)
	}
	if len(_cloudformationDescription) > 0 {
		input.Description = aws.String(_cloudformationDescription)
	}
	if len(_cloudformationExecutionRoleName) > 0 {
		input.ExecutionRoleName = aws.String(_cloudformationExecutionRoleName)
	}
	if len(_cloudformationManagedExecution) > 0 {
		if err := assignInputField(input, "ManagedExecution", _cloudformationManagedExecution); err != nil {
			log.Errorf("invalid --managed-execution: %s", err.Error())
			return
		}
	}
	if len(_cloudformationParameters) > 0 {
		if err := assignInputField(input, "Parameters", _cloudformationParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}
	if len(_cloudformationPermissionModel) > 0 {
		if err := assignInputField(input, "PermissionModel", _cloudformationPermissionModel); err != nil {
			log.Errorf("invalid --permission-model: %s", err.Error())
			return
		}
	}
	if len(_cloudformationStackId) > 0 {
		input.StackId = aws.String(_cloudformationStackId)
	}
	if len(_cloudformationTags) > 0 {
		if err := assignInputField(input, "Tags", _cloudformationTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_cloudformationTemplateBody) > 0 {
		input.TemplateBody = aws.String(_cloudformationTemplateBody)
	}
	if len(_cloudformationTemplateURL) > 0 {
		input.TemplateURL = aws.String(_cloudformationTemplateURL)
	}

	if resp, err := client.CreateStackSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deactivates trusted access with Organizations. If trusted access is
// deactivated, the management account does not have permissions to create and
// manage service-managed StackSets for your organization.
func cloudformation_DeactivateOrganizationsAccess(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.DeactivateOrganizationsAccessInput{}

	if resp, err := client.DeactivateOrganizationsAccess(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deactivates a public third-party extension, such as a resource or module, or a
// CloudFormation Hook when you no longer use it.
//
// Deactivating an extension deletes the configuration details that are associated
// with it. To temporarily disable a CloudFormation Hook instead, you can use [SetTypeConfiguration].
//
// Once deactivated, an extension can't be used in any CloudFormation operation.
// This includes stack update operations where the stack template includes the
// extension, even if no updates are being made to the extension. In addition,
// deactivated extensions aren't automatically updated if a new version of the
// extension is released.
//
// To see which extensions are currently activated, use [ListTypes].
//
// [SetTypeConfiguration]: https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_SetTypeConfiguration.html
// [ListTypes]: https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_ListTypes.html
func cloudformation_DeactivateType(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.DeactivateTypeInput{}

	if len(_cloudformationArn) > 0 {
		input.Arn = aws.String(_cloudformationArn)
	}
	if len(_cloudformationType) > 0 {
		if err := assignInputField(input, "Type", _cloudformationType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_cloudformationTypeName) > 0 {
		input.TypeName = aws.String(_cloudformationTypeName)
	}

	if resp, err := client.DeactivateType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified change set. Deleting change sets ensures that no one
// executes the wrong change set.
//
// If the call successfully completes, CloudFormation successfully deleted the
// change set.
//
// If IncludeNestedStacks specifies True during the creation of the nested change
// set, then DeleteChangeSet will delete all change sets that belong to the stacks
// hierarchy and will also delete all change sets for nested stacks with the status
// of REVIEW_IN_PROGRESS .
func cloudformation_DeleteChangeSet(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.DeleteChangeSetInput{
		// ChangeSetName: *string, // Required
	}

	if len(_cloudformationChangeSetName) > 0 {
		input.ChangeSetName = aws.String(_cloudformationChangeSetName)
	}
	if len(_cloudformationStackName) > 0 {
		input.StackName = aws.String(_cloudformationStackName)
	}

	if resp, err := client.DeleteChangeSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deleted a generated template.
func cloudformation_DeleteGeneratedTemplate(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.DeleteGeneratedTemplateInput{
		// GeneratedTemplateName: *string, // Required
	}

	if len(_cloudformationGeneratedTemplateName) > 0 {
		input.GeneratedTemplateName = aws.String(_cloudformationGeneratedTemplateName)
	}

	if resp, err := client.DeleteGeneratedTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a specified stack. Once the call completes successfully, stack deletion
// starts. Deleted stacks don't show up in the DescribeStacksoperation if the deletion has been
// completed successfully.
//
// For more information about deleting a stack, see [Delete a stack from the CloudFormation console] in the CloudFormation User
// Guide.
//
// [Delete a stack from the CloudFormation console]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/cfn-console-delete-stack.html
func cloudformation_DeleteStack(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.DeleteStackInput{
		// StackName: *string, // Required
	}

	if len(_cloudformationStackName) > 0 {
		input.StackName = aws.String(_cloudformationStackName)
	}
	if len(_cloudformationClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_cloudformationClientRequestToken)
	}
	if len(_cloudformationDeletionMode) > 0 {
		if err := assignInputField(input, "DeletionMode", _cloudformationDeletionMode); err != nil {
			log.Errorf("invalid --deletion-mode: %s", err.Error())
			return
		}
	}
	if len(_cloudformationRetainResources) > 0 {
		input.RetainResources = append([]string(nil), _cloudformationRetainResources...)
	}
	if len(_cloudformationRoleARN) > 0 {
		input.RoleARN = aws.String(_cloudformationRoleARN)
	}

	if resp, err := client.DeleteStack(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes stack instances for the specified accounts, in the specified Amazon Web
// Services Regions.
//
// The maximum number of organizational unit (OUs) supported by a
// DeleteStackInstances operation is 50.
//
// If you need more than 50, consider the following options:
//
// - Batch processing: If you don't want to expose your OU hierarchy, split up
// the operations into multiple calls with less than 50 OUs each.
//
// - Parent OU strategy: If you don't mind exposing the OU hierarchy, target a
// parent OU that contains all desired child OUs.
func cloudformation_DeleteStackInstances(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.DeleteStackInstancesInput{
		// Regions: []string, // Required
		// RetainStacks: *bool, // Required
		// StackSetName: *string, // Required
	}

	if len(_cloudformationRegions) > 0 {
		input.Regions = append([]string(nil), _cloudformationRegions...)
	}
	if len(_cloudformationRetainStacks) > 0 {
		if err := assignInputField(input, "RetainStacks", _cloudformationRetainStacks); err != nil {
			log.Errorf("invalid --retain-stacks: %s", err.Error())
			return
		}
	}
	if len(_cloudformationStackSetName) > 0 {
		input.StackSetName = aws.String(_cloudformationStackSetName)
	}
	if len(_cloudformationAccounts) > 0 {
		input.Accounts = append([]string(nil), _cloudformationAccounts...)
	}
	if len(_cloudformationCallAs) > 0 {
		if err := assignInputField(input, "CallAs", _cloudformationCallAs); err != nil {
			log.Errorf("invalid --call-as: %s", err.Error())
			return
		}
	}
	if len(_cloudformationDeploymentTargets) > 0 {
		if err := assignInputField(input, "DeploymentTargets", _cloudformationDeploymentTargets); err != nil {
			log.Errorf("invalid --deployment-targets: %s", err.Error())
			return
		}
	}
	if len(_cloudformationOperationId) > 0 {
		input.OperationId = aws.String(_cloudformationOperationId)
	}
	if len(_cloudformationOperationPreferences) > 0 {
		if err := assignInputField(input, "OperationPreferences", _cloudformationOperationPreferences); err != nil {
			log.Errorf("invalid --operation-preferences: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteStackInstances(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a StackSet. Before you can delete a StackSet, all its member stack
// instances must be deleted. For more information about how to complete this, see DeleteStackInstances
// .
func cloudformation_DeleteStackSet(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.DeleteStackSetInput{
		// StackSetName: *string, // Required
	}

	if len(_cloudformationStackSetName) > 0 {
		input.StackSetName = aws.String(_cloudformationStackSetName)
	}
	if len(_cloudformationCallAs) > 0 {
		if err := assignInputField(input, "CallAs", _cloudformationCallAs); err != nil {
			log.Errorf("invalid --call-as: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteStackSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Marks an extension or extension version as DEPRECATED in the CloudFormation
// registry, removing it from active use. Deprecated extensions or extension
// versions cannot be used in CloudFormation operations.
//
// To deregister an entire extension, you must individually deregister all active
// versions of that extension. If an extension has only a single active version,
// deregistering that version results in the extension itself being deregistered
// and marked as deprecated in the registry.
//
// You can't deregister the default version of an extension if there are other
// active version of that extension. If you do deregister the default version of an
// extension, the extension type itself is deregistered as well and marked as
// deprecated.
//
// To view the deprecation status of an extension or extension version, use [DescribeType].
//
// For more information, see [Remove third-party private extensions from your account] in the CloudFormation User Guide.
//
// [DescribeType]: https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_DescribeType.html
// [Remove third-party private extensions from your account]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/registry-private-deregister-extension.html
func cloudformation_DeregisterType(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.DeregisterTypeInput{}

	if len(_cloudformationArn) > 0 {
		input.Arn = aws.String(_cloudformationArn)
	}
	if len(_cloudformationType) > 0 {
		if err := assignInputField(input, "Type", _cloudformationType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_cloudformationTypeName) > 0 {
		input.TypeName = aws.String(_cloudformationTypeName)
	}
	if len(_cloudformationVersionId) > 0 {
		input.VersionId = aws.String(_cloudformationVersionId)
	}

	if resp, err := client.DeregisterType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves your account's CloudFormation limits, such as the maximum number of
// stacks that you can create in your account. For more information about account
// limits, see [Understand CloudFormation quotas]in the CloudFormation User Guide.
//
// [Understand CloudFormation quotas]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/cloudformation-limits.html
func cloudformation_DescribeAccountLimits(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.DescribeAccountLimitsInput{}

	if len(_cloudformationNextToken) > 0 {
		input.NextToken = aws.String(_cloudformationNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeAccountLimits(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudformation.DescribeAccountLimitsOutput
	p := cloudformation.NewDescribeAccountLimitsPaginator(client, input)
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

// Returns the inputs for the change set and a list of changes that CloudFormation
// will make if you execute the change set. For more information, see [Update CloudFormation stacks using change sets]in the
// CloudFormation User Guide.
//
// [Update CloudFormation stacks using change sets]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-changesets.html
func cloudformation_DescribeChangeSet(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.DescribeChangeSetInput{
		// ChangeSetName: *string, // Required
	}

	if len(_cloudformationChangeSetName) > 0 {
		input.ChangeSetName = aws.String(_cloudformationChangeSetName)
	}
	if len(_cloudformationIncludePropertyValues) > 0 {
		if err := assignInputField(input, "IncludePropertyValues", _cloudformationIncludePropertyValues); err != nil {
			log.Errorf("invalid --include-property-values: %s", err.Error())
			return
		}
	}
	if len(_cloudformationNextToken) > 0 {
		input.NextToken = aws.String(_cloudformationNextToken)
	}
	if len(_cloudformationStackName) > 0 {
		input.StackName = aws.String(_cloudformationStackName)
	}

	if disablePaginator() {
		if resp, err := client.DescribeChangeSet(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudformation.DescribeChangeSetOutput
	p := cloudformation.NewDescribeChangeSetPaginator(client, input)
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

// Returns Hook-related information for the change set and a list of changes that
// CloudFormation makes when you run the change set.
func cloudformation_DescribeChangeSetHooks(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.DescribeChangeSetHooksInput{
		// ChangeSetName: *string, // Required
	}

	if len(_cloudformationChangeSetName) > 0 {
		input.ChangeSetName = aws.String(_cloudformationChangeSetName)
	}
	if len(_cloudformationLogicalResourceId) > 0 {
		input.LogicalResourceId = aws.String(_cloudformationLogicalResourceId)
	}
	if len(_cloudformationNextToken) > 0 {
		input.NextToken = aws.String(_cloudformationNextToken)
	}
	if len(_cloudformationStackName) > 0 {
		input.StackName = aws.String(_cloudformationStackName)
	}

	if resp, err := client.DescribeChangeSetHooks(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns CloudFormation events based on flexible query criteria. Groups events
// by operation ID, enabling you to focus on individual stack operations during
// deployment.
//
// An operation is any action performed on a stack, including stack lifecycle
// actions (Create, Update, Delete, Rollback), change set creation, nested stack
// creation, and automatic rollbacks triggered by failures. Each operation has a
// unique identifier (Operation ID) and represents a discrete change attempt on the
// stack.
//
// Returns different types of events including:
//
// - Progress events - Status updates during stack operation execution.
//
// - Validation errors - Failures from CloudFormation Early Validations.
//
// - Provisioning errors - Resource creation and update failures.
//
// - Hook invocation errors - Failures from CloudFormation Hook during stack
// operations.
//
// One of ChangeSetName , OperationId or StackName must be specified as input.
func cloudformation_DescribeEvents(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.DescribeEventsInput{}

	if len(_cloudformationChangeSetName) > 0 {
		input.ChangeSetName = aws.String(_cloudformationChangeSetName)
	}
	if len(_cloudformationFilters) > 0 {
		if err := assignInputField(input, "Filters", _cloudformationFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_cloudformationNextToken) > 0 {
		input.NextToken = aws.String(_cloudformationNextToken)
	}
	if len(_cloudformationOperationId) > 0 {
		input.OperationId = aws.String(_cloudformationOperationId)
	}
	if len(_cloudformationStackName) > 0 {
		input.StackName = aws.String(_cloudformationStackName)
	}

	if disablePaginator() {
		if resp, err := client.DescribeEvents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudformation.DescribeEventsOutput
	p := cloudformation.NewDescribeEventsPaginator(client, input)
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

// Describes a generated template. The output includes details about the progress
// of the creation of a generated template started by a CreateGeneratedTemplate
// API action or the update of a generated template started with an
// UpdateGeneratedTemplate API action.
func cloudformation_DescribeGeneratedTemplate(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.DescribeGeneratedTemplateInput{
		// GeneratedTemplateName: *string, // Required
	}

	if len(_cloudformationGeneratedTemplateName) > 0 {
		input.GeneratedTemplateName = aws.String(_cloudformationGeneratedTemplateName)
	}

	if resp, err := client.DescribeGeneratedTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the account's OrganizationAccess status. This API
// can be called either by the management account or the delegated administrator by
// using the CallAs parameter. This API can also be called without the CallAs
// parameter by the management account.
func cloudformation_DescribeOrganizationsAccess(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.DescribeOrganizationsAccessInput{}

	if len(_cloudformationCallAs) > 0 {
		if err := assignInputField(input, "CallAs", _cloudformationCallAs); err != nil {
			log.Errorf("invalid --call-as: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeOrganizationsAccess(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a CloudFormation extension publisher.
// If you don't supply a PublisherId , and you have registered as an extension
// publisher, DescribePublisher returns information about your own publisher
// account.
//
// For more information about registering as a publisher, see:
//
// [RegisterPublisher]
//
// [Publishing extensions to make them available for public use]
// - in the CloudFormation Command Line Interface (CLI) User Guide
//
// [Publishing extensions to make them available for public use]: https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/publish-extension.html
// [RegisterPublisher]: https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_RegisterPublisher.html
func cloudformation_DescribePublisher(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.DescribePublisherInput{}

	if len(_cloudformationPublisherId) > 0 {
		input.PublisherId = aws.String(_cloudformationPublisherId)
	}

	if resp, err := client.DescribePublisher(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes details of a resource scan.
func cloudformation_DescribeResourceScan(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.DescribeResourceScanInput{
		// ResourceScanId: *string, // Required
	}

	if len(_cloudformationResourceScanId) > 0 {
		input.ResourceScanId = aws.String(_cloudformationResourceScanId)
	}

	if resp, err := client.DescribeResourceScan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a stack drift detection operation. A stack drift
// detection operation detects whether a stack's actual configuration differs, or
// has drifted, from its expected configuration, as defined in the stack template
// and any values specified as template parameters. A stack is considered to have
// drifted if one or more of its resources have drifted. For more information about
// stack and resource drift, see [Detect unmanaged configuration changes to stacks and resources with drift detection].
//
// Use DetectStackDrift to initiate a stack drift detection operation. DetectStackDrift returns a
// StackDriftDetectionId you can use to monitor the progress of the operation using
// DescribeStackDriftDetectionStatus . Once the drift detection operation has
// completed, use DescribeStackResourceDriftsto return drift information about the stack and its resources.
//
// [Detect unmanaged configuration changes to stacks and resources with drift detection]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-stack-drift.html
func cloudformation_DescribeStackDriftDetectionStatus(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.DescribeStackDriftDetectionStatusInput{
		// StackDriftDetectionId: *string, // Required
	}

	if len(_cloudformationStackDriftDetectionId) > 0 {
		input.StackDriftDetectionId = aws.String(_cloudformationStackDriftDetectionId)
	}

	if resp, err := client.DescribeStackDriftDetectionStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns all stack related events for a specified stack in reverse chronological
// order. For more information about a stack's event history, see [Understand CloudFormation stack creation events]in the
// CloudFormation User Guide.
//
// You can list events for stacks that have failed to create or have been deleted
// by specifying the unique stack identifier (stack ID).
//
// [Understand CloudFormation stack creation events]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stack-resource-configuration-complete.html
func cloudformation_DescribeStackEvents(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.DescribeStackEventsInput{
		// StackName: *string, // Required
	}

	if len(_cloudformationStackName) > 0 {
		input.StackName = aws.String(_cloudformationStackName)
	}
	if len(_cloudformationNextToken) > 0 {
		input.NextToken = aws.String(_cloudformationNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeStackEvents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudformation.DescribeStackEventsOutput
	p := cloudformation.NewDescribeStackEventsPaginator(client, input)
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

// Returns the stack instance that's associated with the specified StackSet,
// Amazon Web Services account, and Amazon Web Services Region.
//
// For a list of stack instances that are associated with a specific StackSet, use ListStackInstances
// .
func cloudformation_DescribeStackInstance(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.DescribeStackInstanceInput{
		// StackInstanceAccount: *string, // Required
		// StackInstanceRegion: *string, // Required
		// StackSetName: *string, // Required
	}

	if len(_cloudformationStackInstanceAccount) > 0 {
		input.StackInstanceAccount = aws.String(_cloudformationStackInstanceAccount)
	}
	if len(_cloudformationStackInstanceRegion) > 0 {
		input.StackInstanceRegion = aws.String(_cloudformationStackInstanceRegion)
	}
	if len(_cloudformationStackSetName) > 0 {
		input.StackSetName = aws.String(_cloudformationStackSetName)
	}
	if len(_cloudformationCallAs) > 0 {
		if err := assignInputField(input, "CallAs", _cloudformationCallAs); err != nil {
			log.Errorf("invalid --call-as: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeStackInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the stack refactor status.
func cloudformation_DescribeStackRefactor(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.DescribeStackRefactorInput{
		// StackRefactorId: *string, // Required
	}

	if len(_cloudformationStackRefactorId) > 0 {
		input.StackRefactorId = aws.String(_cloudformationStackRefactorId)
	}

	if resp, err := client.DescribeStackRefactor(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a description of the specified resource in the specified stack.
// For deleted stacks, DescribeStackResource returns resource information for up
// to 90 days after the stack has been deleted.
func cloudformation_DescribeStackResource(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.DescribeStackResourceInput{
		// LogicalResourceId: *string, // Required
		// StackName: *string, // Required
	}

	if len(_cloudformationLogicalResourceId) > 0 {
		input.LogicalResourceId = aws.String(_cloudformationLogicalResourceId)
	}
	if len(_cloudformationStackName) > 0 {
		input.StackName = aws.String(_cloudformationStackName)
	}

	if resp, err := client.DescribeStackResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns drift information for the resources that have been checked for drift in
// the specified stack. This includes actual and expected configuration values for
// resources where CloudFormation detects configuration drift.
//
// For a given stack, there will be one StackResourceDrift for each stack resource
// that has been checked for drift. Resources that haven't yet been checked for
// drift aren't included. Resources that don't currently support drift detection
// aren't checked, and so not included. For a list of resources that support drift
// detection, see [Resource type support for imports and drift detection].
//
// Use DetectStackResourceDrift to detect drift on individual resources, or DetectStackDrift to detect drift on all
// supported resources for a given stack.
//
// [Resource type support for imports and drift detection]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/resource-import-supported-resources.html
func cloudformation_DescribeStackResourceDrifts(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.DescribeStackResourceDriftsInput{
		// StackName: *string, // Required
	}

	if len(_cloudformationStackName) > 0 {
		input.StackName = aws.String(_cloudformationStackName)
	}
	if len(_cloudformationMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cloudformationMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cloudformationNextToken) > 0 {
		input.NextToken = aws.String(_cloudformationNextToken)
	}
	if len(_cloudformationStackResourceDriftStatusFilters) > 0 {
		if err := assignInputField(input, "StackResourceDriftStatusFilters", _cloudformationStackResourceDriftStatusFilters); err != nil {
			log.Errorf("invalid --stack-resource-drift-status-filters: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeStackResourceDrifts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudformation.DescribeStackResourceDriftsOutput
	p := cloudformation.NewDescribeStackResourceDriftsPaginator(client, input)
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

// Returns Amazon Web Services resource descriptions for running and deleted
// stacks. If StackName is specified, all the associated resources that are part
// of the stack are returned. If PhysicalResourceId is specified, the associated
// resources of the stack that the resource belongs to are returned.
//
// Only the first 100 resources will be returned. If your stack has more resources
// than this, you should use ListStackResources instead.
//
// For deleted stacks, DescribeStackResources returns resource information for up
// to 90 days after the stack has been deleted.
//
// You must specify either StackName or PhysicalResourceId , but not both. In
// addition, you can specify LogicalResourceId to filter the returned result. For
// more information about resources, the LogicalResourceId and PhysicalResourceId ,
// see the [CloudFormation User Guide].
//
// A ValidationError is returned if you specify both StackName and
// PhysicalResourceId in the same request.
//
// [CloudFormation User Guide]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/
func cloudformation_DescribeStackResources(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.DescribeStackResourcesInput{}

	if len(_cloudformationLogicalResourceId) > 0 {
		input.LogicalResourceId = aws.String(_cloudformationLogicalResourceId)
	}
	if len(_cloudformationPhysicalResourceId) > 0 {
		input.PhysicalResourceId = aws.String(_cloudformationPhysicalResourceId)
	}
	if len(_cloudformationStackName) > 0 {
		input.StackName = aws.String(_cloudformationStackName)
	}

	if resp, err := client.DescribeStackResources(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the description of the specified StackSet.
// This API provides strongly consistent reads meaning it will always return the
// most up-to-date data.
func cloudformation_DescribeStackSet(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.DescribeStackSetInput{
		// StackSetName: *string, // Required
	}

	if len(_cloudformationStackSetName) > 0 {
		input.StackSetName = aws.String(_cloudformationStackSetName)
	}
	if len(_cloudformationCallAs) > 0 {
		if err := assignInputField(input, "CallAs", _cloudformationCallAs); err != nil {
			log.Errorf("invalid --call-as: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeStackSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the description of the specified StackSet operation.
// This API provides strongly consistent reads meaning it will always return the
// most up-to-date data.
func cloudformation_DescribeStackSetOperation(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.DescribeStackSetOperationInput{
		// OperationId: *string, // Required
		// StackSetName: *string, // Required
	}

	if len(_cloudformationOperationId) > 0 {
		input.OperationId = aws.String(_cloudformationOperationId)
	}
	if len(_cloudformationStackSetName) > 0 {
		input.StackSetName = aws.String(_cloudformationStackSetName)
	}
	if len(_cloudformationCallAs) > 0 {
		if err := assignInputField(input, "CallAs", _cloudformationCallAs); err != nil {
			log.Errorf("invalid --call-as: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeStackSetOperation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the description for the specified stack; if no stack name was
// specified, then it returns the description for all the stacks created. For more
// information about a stack's event history, see [Understand CloudFormation stack creation events]in the CloudFormation User Guide.
//
// If the stack doesn't exist, a ValidationError is returned.
//
// [Understand CloudFormation stack creation events]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stack-resource-configuration-complete.html
func cloudformation_DescribeStacks(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.DescribeStacksInput{}

	if len(_cloudformationNextToken) > 0 {
		input.NextToken = aws.String(_cloudformationNextToken)
	}
	if len(_cloudformationStackName) > 0 {
		input.StackName = aws.String(_cloudformationStackName)
	}

	if disablePaginator() {
		if resp, err := client.DescribeStacks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudformation.DescribeStacksOutput
	p := cloudformation.NewDescribeStacksPaginator(client, input)
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

// Returns detailed information about an extension from the CloudFormation
// registry in your current account and Region.
//
// If you specify a VersionId , DescribeType returns information about that
// specific extension version. Otherwise, it returns information about the default
// extension version.
//
// For more information, see [Edit configuration data for extensions in your account] in the CloudFormation User Guide.
//
// [Edit configuration data for extensions in your account]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/registry-set-configuration.html
func cloudformation_DescribeType(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.DescribeTypeInput{}

	if len(_cloudformationArn) > 0 {
		input.Arn = aws.String(_cloudformationArn)
	}
	if len(_cloudformationPublicVersionNumber) > 0 {
		input.PublicVersionNumber = aws.String(_cloudformationPublicVersionNumber)
	}
	if len(_cloudformationPublisherId) > 0 {
		input.PublisherId = aws.String(_cloudformationPublisherId)
	}
	if len(_cloudformationType) > 0 {
		if err := assignInputField(input, "Type", _cloudformationType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_cloudformationTypeName) > 0 {
		input.TypeName = aws.String(_cloudformationTypeName)
	}
	if len(_cloudformationVersionId) > 0 {
		input.VersionId = aws.String(_cloudformationVersionId)
	}

	if resp, err := client.DescribeType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about an extension's registration, including its current
// status and type and version identifiers.
//
// When you initiate a registration request using RegisterType, you can then use DescribeTypeRegistration to monitor
// the progress of that registration request.
//
// Once the registration request has completed, use DescribeType to return detailed
// information about an extension.
func cloudformation_DescribeTypeRegistration(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.DescribeTypeRegistrationInput{
		// RegistrationToken: *string, // Required
	}

	if len(_cloudformationRegistrationToken) > 0 {
		input.RegistrationToken = aws.String(_cloudformationRegistrationToken)
	}

	if resp, err := client.DescribeTypeRegistration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Detects whether a stack's actual configuration differs, or has drifted, from
// its expected configuration, as defined in the stack template and any values
// specified as template parameters. For each resource in the stack that supports
// drift detection, CloudFormation compares the actual configuration of the
// resource with its expected template configuration. Only resource properties
// explicitly defined in the stack template are checked for drift. A stack is
// considered to have drifted if one or more of its resources differ from their
// expected template configurations. For more information, see [Detect unmanaged configuration changes to stacks and resources with drift detection].
//
// Use DetectStackDrift to detect drift on all supported resources for a given
// stack, or DetectStackResourceDriftto detect drift on individual resources.
//
// For a list of stack resources that currently support drift detection, see [Resource type support for imports and drift detection].
//
// DetectStackDrift can take up to several minutes, depending on the number of
// resources contained within the stack. Use DescribeStackDriftDetectionStatusto monitor the progress of a detect
// stack drift operation. Once the drift detection operation has completed, use DescribeStackResourceDriftsto
// return drift information about the stack and its resources.
//
// When detecting drift on a stack, CloudFormation doesn't detect drift on any
// nested stacks belonging to that stack. Perform DetectStackDrift directly on the
// nested stack itself.
//
// [Resource type support for imports and drift detection]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/resource-import-supported-resources.html
// [Detect unmanaged configuration changes to stacks and resources with drift detection]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-stack-drift.html
func cloudformation_DetectStackDrift(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.DetectStackDriftInput{
		// StackName: *string, // Required
	}

	if len(_cloudformationStackName) > 0 {
		input.StackName = aws.String(_cloudformationStackName)
	}
	if len(_cloudformationLogicalResourceIds) > 0 {
		input.LogicalResourceIds = append([]string(nil), _cloudformationLogicalResourceIds...)
	}

	if resp, err := client.DetectStackDrift(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about whether a resource's actual configuration differs, or
// has drifted, from its expected configuration, as defined in the stack template
// and any values specified as template parameters. This information includes
// actual and expected property values for resources in which CloudFormation
// detects drift. Only resource properties explicitly defined in the stack template
// are checked for drift. For more information about stack and resource drift, see [Detect unmanaged configuration changes to stacks and resources with drift detection]
// .
//
// Use DetectStackResourceDrift to detect drift on individual resources, or DetectStackDrift to
// detect drift on all resources in a given stack that support drift detection.
//
// Resources that don't currently support drift detection can't be checked. For a
// list of resources that support drift detection, see [Resource type support for imports and drift detection].
//
// [Resource type support for imports and drift detection]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/resource-import-supported-resources.html
// [Detect unmanaged configuration changes to stacks and resources with drift detection]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-stack-drift.html
func cloudformation_DetectStackResourceDrift(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.DetectStackResourceDriftInput{
		// LogicalResourceId: *string, // Required
		// StackName: *string, // Required
	}

	if len(_cloudformationLogicalResourceId) > 0 {
		input.LogicalResourceId = aws.String(_cloudformationLogicalResourceId)
	}
	if len(_cloudformationStackName) > 0 {
		input.StackName = aws.String(_cloudformationStackName)
	}

	if resp, err := client.DetectStackResourceDrift(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Detect drift on a StackSet. When CloudFormation performs drift detection on a
// StackSet, it performs drift detection on the stack associated with each stack
// instance in the StackSet. For more information, see [Performing drift detection on CloudFormation StackSets].
//
// DetectStackSetDrift returns the OperationId of the StackSet drift detection
// operation. Use this operation id with DescribeStackSetOperationto monitor the progress of the drift
// detection operation. The drift detection operation may take some time, depending
// on the number of stack instances included in the StackSet, in addition to the
// number of resources included in each stack.
//
// Once the operation has completed, use the following actions to return drift
// information:
//
// - Use DescribeStackSetto return detailed information about the stack set, including detailed
// information about the last completed drift operation performed on the StackSet.
// (Information about drift operations that are in progress isn't included.)
//
// - Use ListStackInstancesto return a list of stack instances belonging to the StackSet,
// including the drift status and last drift time checked of each instance.
//
// - Use DescribeStackInstanceto return detailed information about a specific stack instance,
// including its drift status and last drift time checked.
//
// You can only run a single drift detection operation on a given StackSet at one
// time.
//
// To stop a drift detection StackSet operation, use StopStackSetOperation.
//
// [Performing drift detection on CloudFormation StackSets]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-drift.html
func cloudformation_DetectStackSetDrift(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.DetectStackSetDriftInput{
		// StackSetName: *string, // Required
	}

	if len(_cloudformationStackSetName) > 0 {
		input.StackSetName = aws.String(_cloudformationStackSetName)
	}
	if len(_cloudformationCallAs) > 0 {
		if err := assignInputField(input, "CallAs", _cloudformationCallAs); err != nil {
			log.Errorf("invalid --call-as: %s", err.Error())
			return
		}
	}
	if len(_cloudformationOperationId) > 0 {
		input.OperationId = aws.String(_cloudformationOperationId)
	}
	if len(_cloudformationOperationPreferences) > 0 {
		if err := assignInputField(input, "OperationPreferences", _cloudformationOperationPreferences); err != nil {
			log.Errorf("invalid --operation-preferences: %s", err.Error())
			return
		}
	}

	if resp, err := client.DetectStackSetDrift(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the estimated monthly cost of a template. The return value is an Amazon
// Web Services Simple Monthly Calculator URL with a query string that describes
// the resources required to run the template.
func cloudformation_EstimateTemplateCost(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.EstimateTemplateCostInput{}

	if len(_cloudformationParameters) > 0 {
		if err := assignInputField(input, "Parameters", _cloudformationParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}
	if len(_cloudformationTemplateBody) > 0 {
		input.TemplateBody = aws.String(_cloudformationTemplateBody)
	}
	if len(_cloudformationTemplateURL) > 0 {
		input.TemplateURL = aws.String(_cloudformationTemplateURL)
	}

	if resp, err := client.EstimateTemplateCost(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a stack using the input information that was provided when the
// specified change set was created. After the call successfully completes,
// CloudFormation starts updating the stack. Use the DescribeStacksaction to view the status of
// the update.
//
// When you execute a change set, CloudFormation deletes all other change sets
// associated with the stack because they aren't valid for the updated stack.
//
// If a stack policy is associated with the stack, CloudFormation enforces the
// policy during the update. You can't specify a temporary stack policy that
// overrides the current policy.
//
// To create a change set for the entire stack hierarchy, IncludeNestedStacks must
// have been set to True .
func cloudformation_ExecuteChangeSet(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.ExecuteChangeSetInput{
		// ChangeSetName: *string, // Required
	}

	if len(_cloudformationChangeSetName) > 0 {
		input.ChangeSetName = aws.String(_cloudformationChangeSetName)
	}
	if len(_cloudformationClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_cloudformationClientRequestToken)
	}
	if len(_cloudformationDisableRollback) > 0 {
		if err := assignInputField(input, "DisableRollback", _cloudformationDisableRollback); err != nil {
			log.Errorf("invalid --disable-rollback: %s", err.Error())
			return
		}
	}
	if len(_cloudformationRetainExceptOnCreate) > 0 {
		if err := assignInputField(input, "RetainExceptOnCreate", _cloudformationRetainExceptOnCreate); err != nil {
			log.Errorf("invalid --retain-except-on-create: %s", err.Error())
			return
		}
	}
	if len(_cloudformationStackName) > 0 {
		input.StackName = aws.String(_cloudformationStackName)
	}

	if resp, err := client.ExecuteChangeSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Executes the stack refactor operation.
func cloudformation_ExecuteStackRefactor(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.ExecuteStackRefactorInput{
		// StackRefactorId: *string, // Required
	}

	if len(_cloudformationStackRefactorId) > 0 {
		input.StackRefactorId = aws.String(_cloudformationStackRefactorId)
	}

	if resp, err := client.ExecuteStackRefactor(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a generated template. If the template is in an InProgress or Pending
// status then the template returned will be the template when the template was
// last in a Complete status. If the template has not yet been in a Complete
// status then an empty template will be returned.
func cloudformation_GetGeneratedTemplate(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.GetGeneratedTemplateInput{
		// GeneratedTemplateName: *string, // Required
	}

	if len(_cloudformationGeneratedTemplateName) > 0 {
		input.GeneratedTemplateName = aws.String(_cloudformationGeneratedTemplateName)
	}
	if len(_cloudformationFormat) > 0 {
		if err := assignInputField(input, "Format", _cloudformationFormat); err != nil {
			log.Errorf("invalid --format: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetGeneratedTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves detailed information and remediation guidance for a Hook invocation
// result.
//
// If the Hook uses a KMS key to encrypt annotations, callers of the GetHookResult
// operation must have kms:Decrypt permissions. For more information, see [KMS key policy and permissions for encrypting CloudFormation Hooks results at rest] in the
// CloudFormation Hooks User Guide.
//
// [KMS key policy and permissions for encrypting CloudFormation Hooks results at rest]: https://docs.aws.amazon.com/cloudformation-cli/latest/hooks-userguide/hooks-kms-key-policy.html
func cloudformation_GetHookResult(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.GetHookResultInput{}

	if len(_cloudformationHookResultId) > 0 {
		input.HookResultId = aws.String(_cloudformationHookResultId)
	}

	if resp, err := client.GetHookResult(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the stack policy for a specified stack. If a stack doesn't have a
// policy, a null value is returned.
func cloudformation_GetStackPolicy(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.GetStackPolicyInput{
		// StackName: *string, // Required
	}

	if len(_cloudformationStackName) > 0 {
		input.StackName = aws.String(_cloudformationStackName)
	}

	if resp, err := client.GetStackPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the template body for a specified stack. You can get the template for
// running or deleted stacks.
//
// For deleted stacks, GetTemplate returns the template for up to 90 days after
// the stack has been deleted.
//
// If the template doesn't exist, a ValidationError is returned.
func cloudformation_GetTemplate(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.GetTemplateInput{}

	if len(_cloudformationChangeSetName) > 0 {
		input.ChangeSetName = aws.String(_cloudformationChangeSetName)
	}
	if len(_cloudformationStackName) > 0 {
		input.StackName = aws.String(_cloudformationStackName)
	}
	if len(_cloudformationTemplateStage) > 0 {
		if err := assignInputField(input, "TemplateStage", _cloudformationTemplateStage); err != nil {
			log.Errorf("invalid --template-stage: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a new or existing template. The GetTemplateSummary
// action is useful for viewing parameter information, such as default parameter
// values and parameter types, before you create or update a stack or StackSet.
//
// You can use the GetTemplateSummary action when you submit a template, or you
// can get template information for a StackSet, or a running or deleted stack.
//
// For deleted stacks, GetTemplateSummary returns the template information for up
// to 90 days after the stack has been deleted. If the template doesn't exist, a
// ValidationError is returned.
func cloudformation_GetTemplateSummary(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.GetTemplateSummaryInput{}

	if len(_cloudformationCallAs) > 0 {
		if err := assignInputField(input, "CallAs", _cloudformationCallAs); err != nil {
			log.Errorf("invalid --call-as: %s", err.Error())
			return
		}
	}
	if len(_cloudformationStackName) > 0 {
		input.StackName = aws.String(_cloudformationStackName)
	}
	if len(_cloudformationStackSetName) > 0 {
		input.StackSetName = aws.String(_cloudformationStackSetName)
	}
	if len(_cloudformationTemplateBody) > 0 {
		input.TemplateBody = aws.String(_cloudformationTemplateBody)
	}
	if len(_cloudformationTemplateSummaryConfig) > 0 {
		if err := assignInputField(input, "TemplateSummaryConfig", _cloudformationTemplateSummaryConfig); err != nil {
			log.Errorf("invalid --template-summary-config: %s", err.Error())
			return
		}
	}
	if len(_cloudformationTemplateURL) > 0 {
		input.TemplateURL = aws.String(_cloudformationTemplateURL)
	}

	if resp, err := client.GetTemplateSummary(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Import existing stacks into a new StackSets. Use the stack import operation to
// import up to 10 stacks into a new StackSet in the same account as the source
// stack or in a different administrator account and Region, by specifying the
// stack ID of the stack you intend to import.
func cloudformation_ImportStacksToStackSet(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.ImportStacksToStackSetInput{
		// StackSetName: *string, // Required
	}

	if len(_cloudformationStackSetName) > 0 {
		input.StackSetName = aws.String(_cloudformationStackSetName)
	}
	if len(_cloudformationCallAs) > 0 {
		if err := assignInputField(input, "CallAs", _cloudformationCallAs); err != nil {
			log.Errorf("invalid --call-as: %s", err.Error())
			return
		}
	}
	if len(_cloudformationOperationId) > 0 {
		input.OperationId = aws.String(_cloudformationOperationId)
	}
	if len(_cloudformationOperationPreferences) > 0 {
		if err := assignInputField(input, "OperationPreferences", _cloudformationOperationPreferences); err != nil {
			log.Errorf("invalid --operation-preferences: %s", err.Error())
			return
		}
	}
	if len(_cloudformationOrganizationalUnitIds) > 0 {
		input.OrganizationalUnitIds = append([]string(nil), _cloudformationOrganizationalUnitIds...)
	}
	if len(_cloudformationStackIds) > 0 {
		input.StackIds = append([]string(nil), _cloudformationStackIds...)
	}
	if len(_cloudformationStackIdsUrl) > 0 {
		input.StackIdsUrl = aws.String(_cloudformationStackIdsUrl)
	}

	if resp, err := client.ImportStacksToStackSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the ID and status of each active change set for a stack. For example,
// CloudFormation lists change sets that are in the CREATE_IN_PROGRESS or
// CREATE_PENDING state.
func cloudformation_ListChangeSets(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.ListChangeSetsInput{
		// StackName: *string, // Required
	}

	if len(_cloudformationStackName) > 0 {
		input.StackName = aws.String(_cloudformationStackName)
	}
	if len(_cloudformationNextToken) > 0 {
		input.NextToken = aws.String(_cloudformationNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListChangeSets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudformation.ListChangeSetsOutput
	p := cloudformation.NewListChangeSetsPaginator(client, input)
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

// Lists all exported output values in the account and Region in which you call
// this action. Use this action to see the exported output values that you can
// import into other stacks. To import values, use the [Fn::ImportValue]function.
//
// For more information, see [Get exported outputs from a deployed CloudFormation stack].
//
// [Get exported outputs from a deployed CloudFormation stack]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-stack-exports.html
// [Fn::ImportValue]: https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/intrinsic-function-reference-importvalue.html
func cloudformation_ListExports(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.ListExportsInput{}

	if len(_cloudformationNextToken) > 0 {
		input.NextToken = aws.String(_cloudformationNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListExports(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudformation.ListExportsOutput
	p := cloudformation.NewListExportsPaginator(client, input)
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

// Lists your generated templates in this Region.
func cloudformation_ListGeneratedTemplates(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.ListGeneratedTemplatesInput{}

	if len(_cloudformationMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cloudformationMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cloudformationNextToken) > 0 {
		input.NextToken = aws.String(_cloudformationNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListGeneratedTemplates(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudformation.ListGeneratedTemplatesOutput
	p := cloudformation.NewListGeneratedTemplatesPaginator(client, input)
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

// Returns summaries of invoked Hooks. For more information, see [View invocation summaries for CloudFormation Hooks] in the
// CloudFormation Hooks User Guide.
//
// This operation supports the following parameter combinations:
//
// - No parameters: Returns all Hook invocation summaries.
//
// - TypeArn only: Returns summaries for a specific Hook.
//
// - TypeArn and Status : Returns summaries for a specific Hook filtered by
// status.
//
// - TargetId and TargetType : Returns summaries for a specific Hook invocation
// target.
//
// [View invocation summaries for CloudFormation Hooks]: https://docs.aws.amazon.com/cloudformation-cli/latest/hooks-userguide/hooks-view-invocations.html
func cloudformation_ListHookResults(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.ListHookResultsInput{}

	if len(_cloudformationNextToken) > 0 {
		input.NextToken = aws.String(_cloudformationNextToken)
	}
	if len(_cloudformationStatus) > 0 {
		if err := assignInputField(input, "Status", _cloudformationStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_cloudformationTargetId) > 0 {
		input.TargetId = aws.String(_cloudformationTargetId)
	}
	if len(_cloudformationTargetType) > 0 {
		if err := assignInputField(input, "TargetType", _cloudformationTargetType); err != nil {
			log.Errorf("invalid --target-type: %s", err.Error())
			return
		}
	}
	if len(_cloudformationTypeArn) > 0 {
		input.TypeArn = aws.String(_cloudformationTypeArn)
	}

	if resp, err := client.ListHookResults(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all stacks that are importing an exported output value. To modify or
// remove an exported output value, first use this action to see which stacks are
// using it. To see the exported output values in your account, see ListExports.
//
// For more information about importing an exported output value, see the [Fn::ImportValue]
// function.
//
// [Fn::ImportValue]: https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/intrinsic-function-reference-importvalue.html
func cloudformation_ListImports(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.ListImportsInput{
		// ExportName: *string, // Required
	}

	if len(_cloudformationExportName) > 0 {
		input.ExportName = aws.String(_cloudformationExportName)
	}
	if len(_cloudformationNextToken) > 0 {
		input.NextToken = aws.String(_cloudformationNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListImports(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudformation.ListImportsOutput
	p := cloudformation.NewListImportsPaginator(client, input)
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

// Lists the related resources for a list of resources from a resource scan. The
// response indicates whether each returned resource is already managed by
// CloudFormation.
func cloudformation_ListResourceScanRelatedResources(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.ListResourceScanRelatedResourcesInput{
		// ResourceScanId: *string, // Required
		// Resources: []types.ScannedResourceIdentifier, // Required
	}

	if len(_cloudformationResourceScanId) > 0 {
		input.ResourceScanId = aws.String(_cloudformationResourceScanId)
	}
	if len(_cloudformationResources) > 0 {
		if err := assignInputField(input, "Resources", _cloudformationResources); err != nil {
			log.Errorf("invalid --resources: %s", err.Error())
			return
		}
	}
	if len(_cloudformationMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cloudformationMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cloudformationNextToken) > 0 {
		input.NextToken = aws.String(_cloudformationNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListResourceScanRelatedResources(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudformation.ListResourceScanRelatedResourcesOutput
	p := cloudformation.NewListResourceScanRelatedResourcesPaginator(client, input)
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

// Lists the resources from a resource scan. The results can be filtered by
// resource identifier, resource type prefix, tag key, and tag value. Only
// resources that match all specified filters are returned. The response indicates
// whether each returned resource is already managed by CloudFormation.
func cloudformation_ListResourceScanResources(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.ListResourceScanResourcesInput{
		// ResourceScanId: *string, // Required
	}

	if len(_cloudformationResourceScanId) > 0 {
		input.ResourceScanId = aws.String(_cloudformationResourceScanId)
	}
	if len(_cloudformationMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cloudformationMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cloudformationNextToken) > 0 {
		input.NextToken = aws.String(_cloudformationNextToken)
	}
	if len(_cloudformationResourceIdentifier) > 0 {
		input.ResourceIdentifier = aws.String(_cloudformationResourceIdentifier)
	}
	if len(_cloudformationResourceTypePrefix) > 0 {
		input.ResourceTypePrefix = aws.String(_cloudformationResourceTypePrefix)
	}
	if len(_cloudformationTagKey) > 0 {
		input.TagKey = aws.String(_cloudformationTagKey)
	}
	if len(_cloudformationTagValue) > 0 {
		input.TagValue = aws.String(_cloudformationTagValue)
	}

	if disablePaginator() {
		if resp, err := client.ListResourceScanResources(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudformation.ListResourceScanResourcesOutput
	p := cloudformation.NewListResourceScanResourcesPaginator(client, input)
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

// List the resource scans from newest to oldest. By default it will return up to
// 10 resource scans.
func cloudformation_ListResourceScans(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.ListResourceScansInput{}

	if len(_cloudformationMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cloudformationMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cloudformationNextToken) > 0 {
		input.NextToken = aws.String(_cloudformationNextToken)
	}
	if len(_cloudformationScanTypeFilter) > 0 {
		if err := assignInputField(input, "ScanTypeFilter", _cloudformationScanTypeFilter); err != nil {
			log.Errorf("invalid --scan-type-filter: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListResourceScans(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudformation.ListResourceScansOutput
	p := cloudformation.NewListResourceScansPaginator(client, input)
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

// Returns drift information for resources in a stack instance.
// ListStackInstanceResourceDrifts returns drift information for the most recent
// drift detection operation. If an operation is in progress, it may only return
// partial results.
func cloudformation_ListStackInstanceResourceDrifts(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.ListStackInstanceResourceDriftsInput{
		// OperationId: *string, // Required
		// StackInstanceAccount: *string, // Required
		// StackInstanceRegion: *string, // Required
		// StackSetName: *string, // Required
	}

	if len(_cloudformationOperationId) > 0 {
		input.OperationId = aws.String(_cloudformationOperationId)
	}
	if len(_cloudformationStackInstanceAccount) > 0 {
		input.StackInstanceAccount = aws.String(_cloudformationStackInstanceAccount)
	}
	if len(_cloudformationStackInstanceRegion) > 0 {
		input.StackInstanceRegion = aws.String(_cloudformationStackInstanceRegion)
	}
	if len(_cloudformationStackSetName) > 0 {
		input.StackSetName = aws.String(_cloudformationStackSetName)
	}
	if len(_cloudformationCallAs) > 0 {
		if err := assignInputField(input, "CallAs", _cloudformationCallAs); err != nil {
			log.Errorf("invalid --call-as: %s", err.Error())
			return
		}
	}
	if len(_cloudformationMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cloudformationMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cloudformationNextToken) > 0 {
		input.NextToken = aws.String(_cloudformationNextToken)
	}
	if len(_cloudformationStackInstanceResourceDriftStatuses) > 0 {
		if err := assignInputField(input, "StackInstanceResourceDriftStatuses", _cloudformationStackInstanceResourceDriftStatuses); err != nil {
			log.Errorf("invalid --stack-instance-resource-drift-statuses: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListStackInstanceResourceDrifts(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns summary information about stack instances that are associated with the
// specified StackSet. You can filter for stack instances that are associated with
// a specific Amazon Web Services account name or Region, or that have a specific
// status.
func cloudformation_ListStackInstances(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.ListStackInstancesInput{
		// StackSetName: *string, // Required
	}

	if len(_cloudformationStackSetName) > 0 {
		input.StackSetName = aws.String(_cloudformationStackSetName)
	}
	if len(_cloudformationCallAs) > 0 {
		if err := assignInputField(input, "CallAs", _cloudformationCallAs); err != nil {
			log.Errorf("invalid --call-as: %s", err.Error())
			return
		}
	}
	if len(_cloudformationFilters) > 0 {
		if err := assignInputField(input, "Filters", _cloudformationFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_cloudformationMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cloudformationMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cloudformationNextToken) > 0 {
		input.NextToken = aws.String(_cloudformationNextToken)
	}
	if len(_cloudformationStackInstanceAccount) > 0 {
		input.StackInstanceAccount = aws.String(_cloudformationStackInstanceAccount)
	}
	if len(_cloudformationStackInstanceRegion) > 0 {
		input.StackInstanceRegion = aws.String(_cloudformationStackInstanceRegion)
	}

	if disablePaginator() {
		if resp, err := client.ListStackInstances(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudformation.ListStackInstancesOutput
	p := cloudformation.NewListStackInstancesPaginator(client, input)
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

// Lists the stack refactor actions that will be taken after calling the ExecuteStackRefactor action.
func cloudformation_ListStackRefactorActions(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.ListStackRefactorActionsInput{
		// StackRefactorId: *string, // Required
	}

	if len(_cloudformationStackRefactorId) > 0 {
		input.StackRefactorId = aws.String(_cloudformationStackRefactorId)
	}
	if len(_cloudformationMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cloudformationMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cloudformationNextToken) > 0 {
		input.NextToken = aws.String(_cloudformationNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListStackRefactorActions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudformation.ListStackRefactorActionsOutput
	p := cloudformation.NewListStackRefactorActionsPaginator(client, input)
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

// Lists all account stack refactor operations and their statuses.
func cloudformation_ListStackRefactors(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.ListStackRefactorsInput{}

	if len(_cloudformationExecutionStatusFilter) > 0 {
		if err := assignInputField(input, "ExecutionStatusFilter", _cloudformationExecutionStatusFilter); err != nil {
			log.Errorf("invalid --execution-status-filter: %s", err.Error())
			return
		}
	}
	if len(_cloudformationMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cloudformationMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cloudformationNextToken) > 0 {
		input.NextToken = aws.String(_cloudformationNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListStackRefactors(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudformation.ListStackRefactorsOutput
	p := cloudformation.NewListStackRefactorsPaginator(client, input)
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

// Returns descriptions of all resources of the specified stack.
// For deleted stacks, ListStackResources returns resource information for up to
// 90 days after the stack has been deleted.
func cloudformation_ListStackResources(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.ListStackResourcesInput{
		// StackName: *string, // Required
	}

	if len(_cloudformationStackName) > 0 {
		input.StackName = aws.String(_cloudformationStackName)
	}
	if len(_cloudformationNextToken) > 0 {
		input.NextToken = aws.String(_cloudformationNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListStackResources(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudformation.ListStackResourcesOutput
	p := cloudformation.NewListStackResourcesPaginator(client, input)
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

// Returns summary information about deployment targets for a StackSet.
func cloudformation_ListStackSetAutoDeploymentTargets(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.ListStackSetAutoDeploymentTargetsInput{
		// StackSetName: *string, // Required
	}

	if len(_cloudformationStackSetName) > 0 {
		input.StackSetName = aws.String(_cloudformationStackSetName)
	}
	if len(_cloudformationCallAs) > 0 {
		if err := assignInputField(input, "CallAs", _cloudformationCallAs); err != nil {
			log.Errorf("invalid --call-as: %s", err.Error())
			return
		}
	}
	if len(_cloudformationMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cloudformationMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cloudformationNextToken) > 0 {
		input.NextToken = aws.String(_cloudformationNextToken)
	}

	if resp, err := client.ListStackSetAutoDeploymentTargets(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns summary information about the results of a StackSet operation.
// This API provides eventually consistent reads meaning it may take some time but
// will eventually return the most up-to-date data.
func cloudformation_ListStackSetOperationResults(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.ListStackSetOperationResultsInput{
		// OperationId: *string, // Required
		// StackSetName: *string, // Required
	}

	if len(_cloudformationOperationId) > 0 {
		input.OperationId = aws.String(_cloudformationOperationId)
	}
	if len(_cloudformationStackSetName) > 0 {
		input.StackSetName = aws.String(_cloudformationStackSetName)
	}
	if len(_cloudformationCallAs) > 0 {
		if err := assignInputField(input, "CallAs", _cloudformationCallAs); err != nil {
			log.Errorf("invalid --call-as: %s", err.Error())
			return
		}
	}
	if len(_cloudformationFilters) > 0 {
		if err := assignInputField(input, "Filters", _cloudformationFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_cloudformationMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cloudformationMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cloudformationNextToken) > 0 {
		input.NextToken = aws.String(_cloudformationNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListStackSetOperationResults(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudformation.ListStackSetOperationResultsOutput
	p := cloudformation.NewListStackSetOperationResultsPaginator(client, input)
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

// Returns summary information about operations performed on a StackSet.
// This API provides eventually consistent reads meaning it may take some time but
// will eventually return the most up-to-date data.
func cloudformation_ListStackSetOperations(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.ListStackSetOperationsInput{
		// StackSetName: *string, // Required
	}

	if len(_cloudformationStackSetName) > 0 {
		input.StackSetName = aws.String(_cloudformationStackSetName)
	}
	if len(_cloudformationCallAs) > 0 {
		if err := assignInputField(input, "CallAs", _cloudformationCallAs); err != nil {
			log.Errorf("invalid --call-as: %s", err.Error())
			return
		}
	}
	if len(_cloudformationMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cloudformationMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cloudformationNextToken) > 0 {
		input.NextToken = aws.String(_cloudformationNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListStackSetOperations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudformation.ListStackSetOperationsOutput
	p := cloudformation.NewListStackSetOperationsPaginator(client, input)
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

// Returns summary information about StackSets that are associated with the user.
// This API provides strongly consistent reads meaning it will always return the
// most up-to-date data.
//
// - [Self-managed permissions] If you set the CallAs parameter to SELF while
// signed in to your Amazon Web Services account, ListStackSets returns all
// self-managed StackSets in your Amazon Web Services account.
//
// - [Service-managed permissions] If you set the CallAs parameter to SELF while
// signed in to the organization's management account, ListStackSets returns all
// StackSets in the management account.
//
// - [Service-managed permissions] If you set the CallAs parameter to
// DELEGATED_ADMIN while signed in to your member account, ListStackSets returns
// all StackSets with service-managed permissions in the management account.
func cloudformation_ListStackSets(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.ListStackSetsInput{}

	if len(_cloudformationCallAs) > 0 {
		if err := assignInputField(input, "CallAs", _cloudformationCallAs); err != nil {
			log.Errorf("invalid --call-as: %s", err.Error())
			return
		}
	}
	if len(_cloudformationMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cloudformationMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cloudformationNextToken) > 0 {
		input.NextToken = aws.String(_cloudformationNextToken)
	}
	if len(_cloudformationStatus) > 0 {
		if err := assignInputField(input, "Status", _cloudformationStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListStackSets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudformation.ListStackSetsOutput
	p := cloudformation.NewListStackSetsPaginator(client, input)
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

// Returns the summary information for stacks whose status matches the specified
// StackStatusFilter . Summary information for stacks that have been deleted is
// kept for 90 days after the stack is deleted. If no StackStatusFilter is
// specified, summary information for all stacks is returned (including existing
// stacks and stacks that have been deleted).
func cloudformation_ListStacks(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.ListStacksInput{}

	if len(_cloudformationNextToken) > 0 {
		input.NextToken = aws.String(_cloudformationNextToken)
	}
	if len(_cloudformationStackStatusFilter) > 0 {
		if err := assignInputField(input, "StackStatusFilter", _cloudformationStackStatusFilter); err != nil {
			log.Errorf("invalid --stack-status-filter: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListStacks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudformation.ListStacksOutput
	p := cloudformation.NewListStacksPaginator(client, input)
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

// Returns a list of registration tokens for the specified extension(s).
func cloudformation_ListTypeRegistrations(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.ListTypeRegistrationsInput{}

	if len(_cloudformationMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cloudformationMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cloudformationNextToken) > 0 {
		input.NextToken = aws.String(_cloudformationNextToken)
	}
	if len(_cloudformationRegistrationStatusFilter) > 0 {
		if err := assignInputField(input, "RegistrationStatusFilter", _cloudformationRegistrationStatusFilter); err != nil {
			log.Errorf("invalid --registration-status-filter: %s", err.Error())
			return
		}
	}
	if len(_cloudformationType) > 0 {
		if err := assignInputField(input, "Type", _cloudformationType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_cloudformationTypeArn) > 0 {
		input.TypeArn = aws.String(_cloudformationTypeArn)
	}
	if len(_cloudformationTypeName) > 0 {
		input.TypeName = aws.String(_cloudformationTypeName)
	}

	if disablePaginator() {
		if resp, err := client.ListTypeRegistrations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudformation.ListTypeRegistrationsOutput
	p := cloudformation.NewListTypeRegistrationsPaginator(client, input)
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

// Returns summary information about the versions of an extension.
func cloudformation_ListTypeVersions(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.ListTypeVersionsInput{}

	if len(_cloudformationArn) > 0 {
		input.Arn = aws.String(_cloudformationArn)
	}
	if len(_cloudformationDeprecatedStatus) > 0 {
		if err := assignInputField(input, "DeprecatedStatus", _cloudformationDeprecatedStatus); err != nil {
			log.Errorf("invalid --deprecated-status: %s", err.Error())
			return
		}
	}
	if len(_cloudformationMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cloudformationMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cloudformationNextToken) > 0 {
		input.NextToken = aws.String(_cloudformationNextToken)
	}
	if len(_cloudformationPublisherId) > 0 {
		input.PublisherId = aws.String(_cloudformationPublisherId)
	}
	if len(_cloudformationType) > 0 {
		if err := assignInputField(input, "Type", _cloudformationType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_cloudformationTypeName) > 0 {
		input.TypeName = aws.String(_cloudformationTypeName)
	}

	if disablePaginator() {
		if resp, err := client.ListTypeVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudformation.ListTypeVersionsOutput
	p := cloudformation.NewListTypeVersionsPaginator(client, input)
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

// Returns summary information about all extensions, including your private
// resource types, modules, and Hooks as well as all public extensions from Amazon
// Web Services and third-party publishers.
func cloudformation_ListTypes(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.ListTypesInput{}

	if len(_cloudformationDeprecatedStatus) > 0 {
		if err := assignInputField(input, "DeprecatedStatus", _cloudformationDeprecatedStatus); err != nil {
			log.Errorf("invalid --deprecated-status: %s", err.Error())
			return
		}
	}
	if len(_cloudformationFilters) > 0 {
		if err := assignInputField(input, "Filters", _cloudformationFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_cloudformationMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cloudformationMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cloudformationNextToken) > 0 {
		input.NextToken = aws.String(_cloudformationNextToken)
	}
	if len(_cloudformationProvisioningType) > 0 {
		if err := assignInputField(input, "ProvisioningType", _cloudformationProvisioningType); err != nil {
			log.Errorf("invalid --provisioning-type: %s", err.Error())
			return
		}
	}
	if len(_cloudformationType) > 0 {
		if err := assignInputField(input, "Type", _cloudformationType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_cloudformationVisibility) > 0 {
		if err := assignInputField(input, "Visibility", _cloudformationVisibility); err != nil {
			log.Errorf("invalid --visibility: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListTypes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudformation.ListTypesOutput
	p := cloudformation.NewListTypesPaginator(client, input)
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

// Publishes the specified extension to the CloudFormation registry as a public
// extension in this Region. Public extensions are available for use by all
// CloudFormation users. For more information about publishing extensions, see [Publishing extensions to make them available for public use]in
// the CloudFormation Command Line Interface (CLI) User Guide.
//
// To publish an extension, you must be registered as a publisher with
// CloudFormation. For more information, see [RegisterPublisher].
//
// [Publishing extensions to make them available for public use]: https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/publish-extension.html
// [RegisterPublisher]: https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_RegisterPublisher.html
func cloudformation_PublishType(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.PublishTypeInput{}

	if len(_cloudformationArn) > 0 {
		input.Arn = aws.String(_cloudformationArn)
	}
	if len(_cloudformationPublicVersionNumber) > 0 {
		input.PublicVersionNumber = aws.String(_cloudformationPublicVersionNumber)
	}
	if len(_cloudformationType) > 0 {
		if err := assignInputField(input, "Type", _cloudformationType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_cloudformationTypeName) > 0 {
		input.TypeName = aws.String(_cloudformationTypeName)
	}

	if resp, err := client.PublishType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Reports progress of a resource handler to CloudFormation.
// Reserved for use by the [CloudFormation CLI]. Don't use this API in your code.
//
// [CloudFormation CLI]: https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/what-is-cloudformation-cli.html
func cloudformation_RecordHandlerProgress(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.RecordHandlerProgressInput{
		// BearerToken: *string, // Required
		// OperationStatus: types.OperationStatus, // Required
	}

	if len(_cloudformationBearerToken) > 0 {
		input.BearerToken = aws.String(_cloudformationBearerToken)
	}
	if len(_cloudformationOperationStatus) > 0 {
		if err := assignInputField(input, "OperationStatus", _cloudformationOperationStatus); err != nil {
			log.Errorf("invalid --operation-status: %s", err.Error())
			return
		}
	}
	if len(_cloudformationClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_cloudformationClientRequestToken)
	}
	if len(_cloudformationCurrentOperationStatus) > 0 {
		if err := assignInputField(input, "CurrentOperationStatus", _cloudformationCurrentOperationStatus); err != nil {
			log.Errorf("invalid --current-operation-status: %s", err.Error())
			return
		}
	}
	if len(_cloudformationErrorCode) > 0 {
		if err := assignInputField(input, "ErrorCode", _cloudformationErrorCode); err != nil {
			log.Errorf("invalid --error-code: %s", err.Error())
			return
		}
	}
	if len(_cloudformationResourceModel) > 0 {
		input.ResourceModel = aws.String(_cloudformationResourceModel)
	}
	if len(_cloudformationStatusMessage) > 0 {
		input.StatusMessage = aws.String(_cloudformationStatusMessage)
	}

	if resp, err := client.RecordHandlerProgress(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Registers your account as a publisher of public extensions in the
// CloudFormation registry. Public extensions are available for use by all
// CloudFormation users. This publisher ID applies to your account in all Amazon
// Web Services Regions.
//
// For information about requirements for registering as a public extension
// publisher, see [Prerequisite: Registering your account to publish CloudFormation extensions]in the CloudFormation Command Line Interface (CLI) User Guide.
//
// [Prerequisite: Registering your account to publish CloudFormation extensions]: https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/publish-extension.html#publish-extension-prereqs
func cloudformation_RegisterPublisher(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.RegisterPublisherInput{}

	if len(_cloudformationAcceptTermsAndConditions) > 0 {
		if err := assignInputField(input, "AcceptTermsAndConditions", _cloudformationAcceptTermsAndConditions); err != nil {
			log.Errorf("invalid --accept-terms-and-conditions: %s", err.Error())
			return
		}
	}
	if len(_cloudformationConnectionArn) > 0 {
		input.ConnectionArn = aws.String(_cloudformationConnectionArn)
	}

	if resp, err := client.RegisterPublisher(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Registers an extension with the CloudFormation service. Registering an
// extension makes it available for use in CloudFormation templates in your Amazon
// Web Services account, and includes:
//
// - Validating the extension schema.
//
// - Determining which handlers, if any, have been specified for the extension.
//
// - Making the extension available for use in your account.
//
// For more information about how to develop extensions and ready them for
// registration, see [Creating resource types using the CloudFormation CLI]in the CloudFormation Command Line Interface (CLI) User Guide.
//
// You can have a maximum of 50 resource extension versions registered at a time.
// This maximum is per account and per Region. Use [DeregisterType]to deregister specific
// extension versions if necessary.
//
// Once you have initiated a registration request using RegisterType, you can use DescribeTypeRegistration to monitor
// the progress of the registration request.
//
// Once you have registered a private extension in your account and Region, use [SetTypeConfiguration]
// to specify configuration properties for the extension. For more information, see
// [Edit configuration data for extensions in your account]in the CloudFormation User Guide.
//
// [SetTypeConfiguration]: https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_SetTypeConfiguration.html
// [Edit configuration data for extensions in your account]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/registry-set-configuration.html
// [Creating resource types using the CloudFormation CLI]: https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-types.html
// [DeregisterType]: https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_DeregisterType.html
func cloudformation_RegisterType(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.RegisterTypeInput{
		// SchemaHandlerPackage: *string, // Required
		// TypeName: *string, // Required
	}

	if len(_cloudformationSchemaHandlerPackage) > 0 {
		input.SchemaHandlerPackage = aws.String(_cloudformationSchemaHandlerPackage)
	}
	if len(_cloudformationTypeName) > 0 {
		input.TypeName = aws.String(_cloudformationTypeName)
	}
	if len(_cloudformationClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_cloudformationClientRequestToken)
	}
	if len(_cloudformationExecutionRoleArn) > 0 {
		input.ExecutionRoleArn = aws.String(_cloudformationExecutionRoleArn)
	}
	if len(_cloudformationLoggingConfig) > 0 {
		if err := assignInputField(input, "LoggingConfig", _cloudformationLoggingConfig); err != nil {
			log.Errorf("invalid --logging-config: %s", err.Error())
			return
		}
	}
	if len(_cloudformationType) > 0 {
		if err := assignInputField(input, "Type", _cloudformationType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if resp, err := client.RegisterType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// When specifying RollbackStack , you preserve the state of previously provisioned
// resources when an operation fails. You can check the status of the stack through
// the DescribeStacksoperation.
//
// Rolls back the specified stack to the last known stable state from CREATE_FAILED
// or UPDATE_FAILED stack statuses.
//
// This operation will delete a stack if it doesn't contain a last known stable
// state. A last known stable state includes any status in a *_COMPLETE . This
// includes the following stack statuses.
//
// - CREATE_COMPLETE
//
// - UPDATE_COMPLETE
//
// - UPDATE_ROLLBACK_COMPLETE
//
// - IMPORT_COMPLETE
//
// - IMPORT_ROLLBACK_COMPLETE
func cloudformation_RollbackStack(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.RollbackStackInput{
		// StackName: *string, // Required
	}

	if len(_cloudformationStackName) > 0 {
		input.StackName = aws.String(_cloudformationStackName)
	}
	if len(_cloudformationClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_cloudformationClientRequestToken)
	}
	if len(_cloudformationRetainExceptOnCreate) > 0 {
		if err := assignInputField(input, "RetainExceptOnCreate", _cloudformationRetainExceptOnCreate); err != nil {
			log.Errorf("invalid --retain-except-on-create: %s", err.Error())
			return
		}
	}
	if len(_cloudformationRoleARN) > 0 {
		input.RoleARN = aws.String(_cloudformationRoleARN)
	}

	if resp, err := client.RollbackStack(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets a stack policy for a specified stack.
func cloudformation_SetStackPolicy(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.SetStackPolicyInput{
		// StackName: *string, // Required
	}

	if len(_cloudformationStackName) > 0 {
		input.StackName = aws.String(_cloudformationStackName)
	}
	if len(_cloudformationStackPolicyBody) > 0 {
		input.StackPolicyBody = aws.String(_cloudformationStackPolicyBody)
	}
	if len(_cloudformationStackPolicyURL) > 0 {
		input.StackPolicyURL = aws.String(_cloudformationStackPolicyURL)
	}

	if resp, err := client.SetStackPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Specifies the configuration data for a CloudFormation extension, such as a
// resource or Hook, in the given account and Region.
//
// For more information, see [Edit configuration data for extensions in your account] in the CloudFormation User Guide.
//
// To view the current configuration data for an extension, refer to the
// ConfigurationSchema element of [DescribeType].
//
// It's strongly recommended that you use dynamic references to restrict sensitive
// configuration definitions, such as third-party credentials. For more
// information, see [Specify values stored in other services using dynamic references]in the CloudFormation User Guide.
//
// For more information about setting the configuration data for resource types,
// see [Defining the account-level configuration of an extension]in the CloudFormation Command Line Interface (CLI) User Guide. For more
// information about setting the configuration data for Hooks, see the [CloudFormation Hooks User Guide].
//
// [DescribeType]: https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_DescribeType.html
// [Edit configuration data for extensions in your account]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/registry-set-configuration.html
// [Specify values stored in other services using dynamic references]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/dynamic-references.html
// [Defining the account-level configuration of an extension]: https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-type-model.html#resource-type-howto-configuration
// [CloudFormation Hooks User Guide]: https://docs.aws.amazon.com/cloudformation-cli/latest/hooks-userguide/what-is-cloudformation-hooks.html
func cloudformation_SetTypeConfiguration(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.SetTypeConfigurationInput{
		// Configuration: *string, // Required
	}

	if len(_cloudformationConfiguration) > 0 {
		input.Configuration = aws.String(_cloudformationConfiguration)
	}
	if len(_cloudformationConfigurationAlias) > 0 {
		input.ConfigurationAlias = aws.String(_cloudformationConfigurationAlias)
	}
	if len(_cloudformationType) > 0 {
		if err := assignInputField(input, "Type", _cloudformationType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_cloudformationTypeArn) > 0 {
		input.TypeArn = aws.String(_cloudformationTypeArn)
	}
	if len(_cloudformationTypeName) > 0 {
		input.TypeName = aws.String(_cloudformationTypeName)
	}

	if resp, err := client.SetTypeConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Specify the default version of an extension. The default version of an
// extension will be used in CloudFormation operations.
func cloudformation_SetTypeDefaultVersion(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.SetTypeDefaultVersionInput{}

	if len(_cloudformationArn) > 0 {
		input.Arn = aws.String(_cloudformationArn)
	}
	if len(_cloudformationType) > 0 {
		if err := assignInputField(input, "Type", _cloudformationType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_cloudformationTypeName) > 0 {
		input.TypeName = aws.String(_cloudformationTypeName)
	}
	if len(_cloudformationVersionId) > 0 {
		input.VersionId = aws.String(_cloudformationVersionId)
	}

	if resp, err := client.SetTypeDefaultVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sends a signal to the specified resource with a success or failure status. You
// can use the SignalResource operation in conjunction with a creation policy or
// update policy. CloudFormation doesn't proceed with a stack creation or update
// until resources receive the required number of signals or the timeout period is
// exceeded. The SignalResource operation is useful in cases where you want to
// send signals from anywhere other than an Amazon EC2 instance.
func cloudformation_SignalResource(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.SignalResourceInput{
		// LogicalResourceId: *string, // Required
		// StackName: *string, // Required
		// Status: types.ResourceSignalStatus, // Required
		// UniqueId: *string, // Required
	}

	if len(_cloudformationLogicalResourceId) > 0 {
		input.LogicalResourceId = aws.String(_cloudformationLogicalResourceId)
	}
	if len(_cloudformationStackName) > 0 {
		input.StackName = aws.String(_cloudformationStackName)
	}
	if len(_cloudformationStatus) > 0 {
		if err := assignInputField(input, "Status", _cloudformationStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_cloudformationUniqueId) > 0 {
		input.UniqueId = aws.String(_cloudformationUniqueId)
	}

	if resp, err := client.SignalResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a scan of the resources in this account in this Region. You can the
// status of a scan using the ListResourceScans API action.
func cloudformation_StartResourceScan(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.StartResourceScanInput{}

	if len(_cloudformationClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_cloudformationClientRequestToken)
	}
	if len(_cloudformationScanFilters) > 0 {
		if err := assignInputField(input, "ScanFilters", _cloudformationScanFilters); err != nil {
			log.Errorf("invalid --scan-filters: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartResourceScan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops an in-progress operation on a StackSet and its associated stack
// instances. StackSets will cancel all the unstarted stack instance deployments
// and wait for those are in-progress to complete.
func cloudformation_StopStackSetOperation(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.StopStackSetOperationInput{
		// OperationId: *string, // Required
		// StackSetName: *string, // Required
	}

	if len(_cloudformationOperationId) > 0 {
		input.OperationId = aws.String(_cloudformationOperationId)
	}
	if len(_cloudformationStackSetName) > 0 {
		input.StackSetName = aws.String(_cloudformationStackSetName)
	}
	if len(_cloudformationCallAs) > 0 {
		if err := assignInputField(input, "CallAs", _cloudformationCallAs); err != nil {
			log.Errorf("invalid --call-as: %s", err.Error())
			return
		}
	}

	if resp, err := client.StopStackSetOperation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Tests a registered extension to make sure it meets all necessary requirements
// for being published in the CloudFormation registry.
//
// - For resource types, this includes passing all contracts tests defined for
// the type.
//
// - For modules, this includes determining if the module's model meets all
// necessary requirements.
//
// For more information, see [Testing your public extension before publishing] in the CloudFormation Command Line Interface (CLI)
// User Guide.
//
// If you don't specify a version, CloudFormation uses the default version of the
// extension in your account and Region for testing.
//
// To perform testing, CloudFormation assumes the execution role specified when
// the type was registered. For more information, see [RegisterType].
//
// Once you've initiated testing on an extension using TestType , you can pass the
// returned TypeVersionArn into [DescribeType] to monitor the current test status and test
// status description for the extension.
//
// An extension must have a test status of PASSED before it can be published. For
// more information, see [Publishing extensions to make them available for public use]in the CloudFormation Command Line Interface (CLI) User
// Guide.
//
// [DescribeType]: https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_DescribeType.html
// [Testing your public extension before publishing]: https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/publish-extension.html#publish-extension-testing
// [RegisterType]: https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_RegisterType.html
// [Publishing extensions to make them available for public use]: https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-type-publish.html
func cloudformation_TestType(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.TestTypeInput{}

	if len(_cloudformationArn) > 0 {
		input.Arn = aws.String(_cloudformationArn)
	}
	if len(_cloudformationLogDeliveryBucket) > 0 {
		input.LogDeliveryBucket = aws.String(_cloudformationLogDeliveryBucket)
	}
	if len(_cloudformationType) > 0 {
		if err := assignInputField(input, "Type", _cloudformationType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_cloudformationTypeName) > 0 {
		input.TypeName = aws.String(_cloudformationTypeName)
	}
	if len(_cloudformationVersionId) > 0 {
		input.VersionId = aws.String(_cloudformationVersionId)
	}

	if resp, err := client.TestType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a generated template. This can be used to change the name, add and
// remove resources, refresh resources, and change the DeletionPolicy and
// UpdateReplacePolicy settings. You can check the status of the update to the
// generated template using the DescribeGeneratedTemplate API action.
func cloudformation_UpdateGeneratedTemplate(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.UpdateGeneratedTemplateInput{
		// GeneratedTemplateName: *string, // Required
	}

	if len(_cloudformationGeneratedTemplateName) > 0 {
		input.GeneratedTemplateName = aws.String(_cloudformationGeneratedTemplateName)
	}
	if len(_cloudformationAddResources) > 0 {
		if err := assignInputField(input, "AddResources", _cloudformationAddResources); err != nil {
			log.Errorf("invalid --add-resources: %s", err.Error())
			return
		}
	}
	if len(_cloudformationNewGeneratedTemplateName) > 0 {
		input.NewGeneratedTemplateName = aws.String(_cloudformationNewGeneratedTemplateName)
	}
	if len(_cloudformationRefreshAllResources) > 0 {
		if err := assignInputField(input, "RefreshAllResources", _cloudformationRefreshAllResources); err != nil {
			log.Errorf("invalid --refresh-all-resources: %s", err.Error())
			return
		}
	}
	if len(_cloudformationRemoveResources) > 0 {
		input.RemoveResources = append([]string(nil), _cloudformationRemoveResources...)
	}
	if len(_cloudformationTemplateConfiguration) > 0 {
		if err := assignInputField(input, "TemplateConfiguration", _cloudformationTemplateConfiguration); err != nil {
			log.Errorf("invalid --template-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateGeneratedTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a stack as specified in the template. After the call completes
// successfully, the stack update starts. You can check the status of the stack
// through the DescribeStacksaction.
//
// To get a copy of the template for an existing stack, you can use the GetTemplate action.
//
// For more information about updating a stack and monitoring the progress of the
// update, see [Managing Amazon Web Services resources as a single unit with CloudFormation stacks]in the CloudFormation User Guide.
//
// [Managing Amazon Web Services resources as a single unit with CloudFormation stacks]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacks.html
func cloudformation_UpdateStack(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.UpdateStackInput{
		// StackName: *string, // Required
	}

	if len(_cloudformationStackName) > 0 {
		input.StackName = aws.String(_cloudformationStackName)
	}
	if len(_cloudformationCapabilities) > 0 {
		if err := assignInputField(input, "Capabilities", _cloudformationCapabilities); err != nil {
			log.Errorf("invalid --capabilities: %s", err.Error())
			return
		}
	}
	if len(_cloudformationClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_cloudformationClientRequestToken)
	}
	if len(_cloudformationDisableRollback) > 0 {
		if err := assignInputField(input, "DisableRollback", _cloudformationDisableRollback); err != nil {
			log.Errorf("invalid --disable-rollback: %s", err.Error())
			return
		}
	}
	if len(_cloudformationNotificationARNs) > 0 {
		input.NotificationARNs = append([]string(nil), _cloudformationNotificationARNs...)
	}
	if len(_cloudformationParameters) > 0 {
		if err := assignInputField(input, "Parameters", _cloudformationParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}
	if len(_cloudformationResourceTypes) > 0 {
		input.ResourceTypes = append([]string(nil), _cloudformationResourceTypes...)
	}
	if len(_cloudformationRetainExceptOnCreate) > 0 {
		if err := assignInputField(input, "RetainExceptOnCreate", _cloudformationRetainExceptOnCreate); err != nil {
			log.Errorf("invalid --retain-except-on-create: %s", err.Error())
			return
		}
	}
	if len(_cloudformationRoleARN) > 0 {
		input.RoleARN = aws.String(_cloudformationRoleARN)
	}
	if len(_cloudformationRollbackConfiguration) > 0 {
		if err := assignInputField(input, "RollbackConfiguration", _cloudformationRollbackConfiguration); err != nil {
			log.Errorf("invalid --rollback-configuration: %s", err.Error())
			return
		}
	}
	if len(_cloudformationStackPolicyBody) > 0 {
		input.StackPolicyBody = aws.String(_cloudformationStackPolicyBody)
	}
	if len(_cloudformationStackPolicyDuringUpdateBody) > 0 {
		input.StackPolicyDuringUpdateBody = aws.String(_cloudformationStackPolicyDuringUpdateBody)
	}
	if len(_cloudformationStackPolicyDuringUpdateURL) > 0 {
		input.StackPolicyDuringUpdateURL = aws.String(_cloudformationStackPolicyDuringUpdateURL)
	}
	if len(_cloudformationStackPolicyURL) > 0 {
		input.StackPolicyURL = aws.String(_cloudformationStackPolicyURL)
	}
	if len(_cloudformationTags) > 0 {
		if err := assignInputField(input, "Tags", _cloudformationTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_cloudformationTemplateBody) > 0 {
		input.TemplateBody = aws.String(_cloudformationTemplateBody)
	}
	if len(_cloudformationTemplateURL) > 0 {
		input.TemplateURL = aws.String(_cloudformationTemplateURL)
	}
	if len(_cloudformationUsePreviousTemplate) > 0 {
		if err := assignInputField(input, "UsePreviousTemplate", _cloudformationUsePreviousTemplate); err != nil {
			log.Errorf("invalid --use-previous-template: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateStack(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the parameter values for stack instances for the specified accounts,
// within the specified Amazon Web Services Regions. A stack instance refers to a
// stack in a specific account and Region.
//
// You can only update stack instances in Amazon Web Services Regions and accounts
// where they already exist; to create additional stack instances, use [CreateStackInstances].
//
// During StackSet updates, any parameters overridden for a stack instance aren't
// updated, but retain their overridden value.
//
// You can only update the parameter values that are specified in the StackSet. To
// add or delete a parameter itself, use [UpdateStackSet]to update the StackSet template. If you
// add a parameter to a template, before you can override the parameter value
// specified in the StackSet you must first use [UpdateStackSet]to update all stack instances with
// the updated template and parameter value specified in the StackSet. Once a stack
// instance has been updated with the new parameter, you can then override the
// parameter value using UpdateStackInstances .
//
// The maximum number of organizational unit (OUs) supported by a
// UpdateStackInstances operation is 50.
//
// If you need more than 50, consider the following options:
//
// - Batch processing: If you don't want to expose your OU hierarchy, split up
// the operations into multiple calls with less than 50 OUs each.
//
// - Parent OU strategy: If you don't mind exposing the OU hierarchy, target a
// parent OU that contains all desired child OUs.
//
// [CreateStackInstances]: https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_CreateStackInstances.html
// [UpdateStackSet]: https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_UpdateStackSet.html
func cloudformation_UpdateStackInstances(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.UpdateStackInstancesInput{
		// Regions: []string, // Required
		// StackSetName: *string, // Required
	}

	if len(_cloudformationRegions) > 0 {
		input.Regions = append([]string(nil), _cloudformationRegions...)
	}
	if len(_cloudformationStackSetName) > 0 {
		input.StackSetName = aws.String(_cloudformationStackSetName)
	}
	if len(_cloudformationAccounts) > 0 {
		input.Accounts = append([]string(nil), _cloudformationAccounts...)
	}
	if len(_cloudformationCallAs) > 0 {
		if err := assignInputField(input, "CallAs", _cloudformationCallAs); err != nil {
			log.Errorf("invalid --call-as: %s", err.Error())
			return
		}
	}
	if len(_cloudformationDeploymentTargets) > 0 {
		if err := assignInputField(input, "DeploymentTargets", _cloudformationDeploymentTargets); err != nil {
			log.Errorf("invalid --deployment-targets: %s", err.Error())
			return
		}
	}
	if len(_cloudformationOperationId) > 0 {
		input.OperationId = aws.String(_cloudformationOperationId)
	}
	if len(_cloudformationOperationPreferences) > 0 {
		if err := assignInputField(input, "OperationPreferences", _cloudformationOperationPreferences); err != nil {
			log.Errorf("invalid --operation-preferences: %s", err.Error())
			return
		}
	}
	if len(_cloudformationParameterOverrides) > 0 {
		if err := assignInputField(input, "ParameterOverrides", _cloudformationParameterOverrides); err != nil {
			log.Errorf("invalid --parameter-overrides: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateStackInstances(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the StackSet and associated stack instances in the specified accounts
// and Amazon Web Services Regions.
//
// Even if the StackSet operation created by updating the StackSet fails
// (completely or partially, below or above a specified failure tolerance), the
// StackSet is updated with your changes. Subsequent CreateStackInstancescalls on the specified
// StackSet use the updated StackSet.
//
// The maximum number of organizational unit (OUs) supported by a UpdateStackSet
// operation is 50.
//
// If you need more than 50, consider the following options:
//
// - Batch processing: If you don't want to expose your OU hierarchy, split up
// the operations into multiple calls with less than 50 OUs each.
//
// - Parent OU strategy: If you don't mind exposing the OU hierarchy, target a
// parent OU that contains all desired child OUs.
func cloudformation_UpdateStackSet(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.UpdateStackSetInput{
		// StackSetName: *string, // Required
	}

	if len(_cloudformationStackSetName) > 0 {
		input.StackSetName = aws.String(_cloudformationStackSetName)
	}
	if len(_cloudformationAccounts) > 0 {
		input.Accounts = append([]string(nil), _cloudformationAccounts...)
	}
	if len(_cloudformationAdministrationRoleARN) > 0 {
		input.AdministrationRoleARN = aws.String(_cloudformationAdministrationRoleARN)
	}
	if len(_cloudformationAutoDeployment) > 0 {
		if err := assignInputField(input, "AutoDeployment", _cloudformationAutoDeployment); err != nil {
			log.Errorf("invalid --auto-deployment: %s", err.Error())
			return
		}
	}
	if len(_cloudformationCallAs) > 0 {
		if err := assignInputField(input, "CallAs", _cloudformationCallAs); err != nil {
			log.Errorf("invalid --call-as: %s", err.Error())
			return
		}
	}
	if len(_cloudformationCapabilities) > 0 {
		if err := assignInputField(input, "Capabilities", _cloudformationCapabilities); err != nil {
			log.Errorf("invalid --capabilities: %s", err.Error())
			return
		}
	}
	if len(_cloudformationDeploymentTargets) > 0 {
		if err := assignInputField(input, "DeploymentTargets", _cloudformationDeploymentTargets); err != nil {
			log.Errorf("invalid --deployment-targets: %s", err.Error())
			return
		}
	}
	if len(_cloudformationDescription) > 0 {
		input.Description = aws.String(_cloudformationDescription)
	}
	if len(_cloudformationExecutionRoleName) > 0 {
		input.ExecutionRoleName = aws.String(_cloudformationExecutionRoleName)
	}
	if len(_cloudformationManagedExecution) > 0 {
		if err := assignInputField(input, "ManagedExecution", _cloudformationManagedExecution); err != nil {
			log.Errorf("invalid --managed-execution: %s", err.Error())
			return
		}
	}
	if len(_cloudformationOperationId) > 0 {
		input.OperationId = aws.String(_cloudformationOperationId)
	}
	if len(_cloudformationOperationPreferences) > 0 {
		if err := assignInputField(input, "OperationPreferences", _cloudformationOperationPreferences); err != nil {
			log.Errorf("invalid --operation-preferences: %s", err.Error())
			return
		}
	}
	if len(_cloudformationParameters) > 0 {
		if err := assignInputField(input, "Parameters", _cloudformationParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}
	if len(_cloudformationPermissionModel) > 0 {
		if err := assignInputField(input, "PermissionModel", _cloudformationPermissionModel); err != nil {
			log.Errorf("invalid --permission-model: %s", err.Error())
			return
		}
	}
	if len(_cloudformationRegions) > 0 {
		input.Regions = append([]string(nil), _cloudformationRegions...)
	}
	if len(_cloudformationTags) > 0 {
		if err := assignInputField(input, "Tags", _cloudformationTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_cloudformationTemplateBody) > 0 {
		input.TemplateBody = aws.String(_cloudformationTemplateBody)
	}
	if len(_cloudformationTemplateURL) > 0 {
		input.TemplateURL = aws.String(_cloudformationTemplateURL)
	}
	if len(_cloudformationUsePreviousTemplate) > 0 {
		if err := assignInputField(input, "UsePreviousTemplate", _cloudformationUsePreviousTemplate); err != nil {
			log.Errorf("invalid --use-previous-template: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateStackSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates termination protection for the specified stack. If a user attempts to
// delete a stack with termination protection enabled, the operation fails and the
// stack remains unchanged. For more information, see [Protect a CloudFormation stack from being deleted]in the CloudFormation User
// Guide.
//
// For [nested stacks], termination protection is set on the root stack and can't be changed
// directly on the nested stack.
//
// [nested stacks]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-nested-stacks.html
// [Protect a CloudFormation stack from being deleted]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-protect-stacks.html
func cloudformation_UpdateTerminationProtection(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.UpdateTerminationProtectionInput{
		// EnableTerminationProtection: *bool, // Required
		// StackName: *string, // Required
	}

	if len(_cloudformationEnableTerminationProtection) > 0 {
		if err := assignInputField(input, "EnableTerminationProtection", _cloudformationEnableTerminationProtection); err != nil {
			log.Errorf("invalid --enable-termination-protection: %s", err.Error())
			return
		}
	}
	if len(_cloudformationStackName) > 0 {
		input.StackName = aws.String(_cloudformationStackName)
	}

	if resp, err := client.UpdateTerminationProtection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Validates a specified template. CloudFormation first checks if the template is
// valid JSON. If it isn't, CloudFormation checks if the template is valid YAML. If
// both these checks fail, CloudFormation returns a template validation error.
func cloudformation_ValidateTemplate(cfg aws.Config, client *cloudformation.Client) {
	input := &cloudformation.ValidateTemplateInput{}

	if len(_cloudformationTemplateBody) > 0 {
		input.TemplateBody = aws.String(_cloudformationTemplateBody)
	}
	if len(_cloudformationTemplateURL) > 0 {
		input.TemplateURL = aws.String(_cloudformationTemplateURL)
	}

	if resp, err := client.ValidateTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_cloudformationCmd)
	_cloudformationCmd.Flags().SortFlags = false

	_cloudformationCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_cloudformationCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_cloudformationCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_cloudformationCmd.Flags().StringVarP(&_cloudformationAcceptTermsAndConditions, "accept-terms-and-conditions", "", "", "Accept Terms And Conditions")
	_cloudformationCmd.Flags().StringSliceVarP(&_cloudformationAccounts, "accounts", "", nil, "Accounts")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationAddResources, "add-resources", "", "", "Add Resources")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationAdministrationRoleARN, "administration-role-arn", "", "", "Administration Role ARN")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationArn, "arn", "", "", "ARN")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationAutoDeployment, "auto-deployment", "", "", "Auto Deployment")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationAutoUpdate, "auto-update", "", "", "Auto Update")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationBearerToken, "bearer-token", "", "", "Bearer Token")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationCallAs, "call-as", "", "", "Call As")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationCapabilities, "capabilities", "", "", "Capabilities")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationChangeSetName, "change-set-name", "", "", "Change Set Name")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationChangeSetType, "change-set-type", "", "", "Change Set Type")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationClientRequestToken, "client-request-token", "", "", "Client Request Token")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationClientToken, "client-token", "", "", "Client Token")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationConfiguration, "configuration", "", "", "Configuration")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationConfigurationAlias, "configuration-alias", "", "", "Configuration Alias")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationConnectionArn, "connection-arn", "", "", "Connection ARN")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationCurrentOperationStatus, "current-operation-status", "", "", "Current Operation Status")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationDeletionMode, "deletion-mode", "", "", "Deletion Mode")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationDeploymentMode, "deployment-mode", "", "", "Deployment Mode")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationDeploymentTargets, "deployment-targets", "", "", "Deployment Targets")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationDeprecatedStatus, "deprecated-status", "", "", "Deprecated Status")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationDescription, "description", "", "", "Description")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationDisableRollback, "disable-rollback", "", "", "Disable Rollback")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationEnableStackCreation, "enable-stack-creation", "", "", "Enable Stack Creation")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationEnableTerminationProtection, "enable-termination-protection", "", "", "Enable Termination Protection")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationErrorCode, "error-code", "", "", "Error Code")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationExecutionRoleArn, "execution-role-arn", "", "", "Execution Role ARN")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationExecutionRoleName, "execution-role-name", "", "", "Execution Role Name")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationExecutionStatusFilter, "execution-status-filter", "", "", "Execution Status Filter")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationExportName, "export-name", "", "", "Export Name")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationFilters, "filters", "", "", "Filters")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationFormat, "format", "", "", "Format")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationGeneratedTemplateName, "generated-template-name", "", "", "Generated Template Name")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationHookResultId, "hook-result-id", "", "", "Hook Result ID")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationImportExistingResources, "import-existing-resources", "", "", "Import Existing Resources")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationIncludeNestedStacks, "include-nested-stacks", "", "", "Include Nested Stacks")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationIncludePropertyValues, "include-property-values", "", "", "Include Property Values")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationLogDeliveryBucket, "log-delivery-bucket", "", "", "Log Delivery Bucket")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationLoggingConfig, "logging-config", "", "", "Logging Config")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationLogicalResourceId, "logical-resource-id", "", "", "Logical Resource ID")
	_cloudformationCmd.Flags().StringSliceVarP(&_cloudformationLogicalResourceIds, "logical-resource-ids", "", nil, "Logical Resource Ids")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationMajorVersion, "major-version", "", "", "Major Version")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationManagedExecution, "managed-execution", "", "", "Managed Execution")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationMaxResults, "max-results", "", "", "Max Results")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationNewGeneratedTemplateName, "new-generated-template-name", "", "", "New Generated Template Name")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationNextToken, "next-token", "", "", "Next Token")
	_cloudformationCmd.Flags().StringSliceVarP(&_cloudformationNotificationARNs, "notification-arns", "", nil, "Notification Arns")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationOnFailure, "on-failure", "", "", "On Failure")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationOnStackFailure, "on-stack-failure", "", "", "On Stack Failure")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationOperationId, "operation-id", "", "", "Operation ID")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationOperationPreferences, "operation-preferences", "", "", "Operation Preferences")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationOperationStatus, "operation-status", "", "", "Operation Status")
	_cloudformationCmd.Flags().StringSliceVarP(&_cloudformationOrganizationalUnitIds, "organizational-unit-ids", "", nil, "Organizational Unit Ids")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationParameterOverrides, "parameter-overrides", "", "", "Parameter Overrides")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationParameters, "parameters", "", "", "Parameters")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationPermissionModel, "permission-model", "", "", "Permission Model")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationPhysicalResourceId, "physical-resource-id", "", "", "Physical Resource ID")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationProvisioningType, "provisioning-type", "", "", "Provisioning Type")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationPublicTypeArn, "public-type-arn", "", "", "Public Type ARN")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationPublicVersionNumber, "public-version-number", "", "", "Public Version Number")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationPublisherId, "publisher-id", "", "", "Publisher ID")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationRefreshAllResources, "refresh-all-resources", "", "", "Refresh All Resources")
	_cloudformationCmd.Flags().StringSliceVarP(&_cloudformationRegions, "regions", "", nil, "Regions")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationRegistrationStatusFilter, "registration-status-filter", "", "", "Registration Status Filter")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationRegistrationToken, "registration-token", "", "", "Registration Token")
	_cloudformationCmd.Flags().StringSliceVarP(&_cloudformationRemoveResources, "remove-resources", "", nil, "Remove Resources")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationResourceIdentifier, "resource-identifier", "", "", "Resource Identifier")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationResourceMappings, "resource-mappings", "", "", "Resource Mappings")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationResourceModel, "resource-model", "", "", "Resource Model")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationResourceScanId, "resource-scan-id", "", "", "Resource Scan ID")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationResourceTypePrefix, "resource-type-prefix", "", "", "Resource Type Prefix")
	_cloudformationCmd.Flags().StringSliceVarP(&_cloudformationResourceTypes, "resource-types", "", nil, "Resource Types")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationResources, "resources", "", "", "Resources")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationResourcesToImport, "resources-to-import", "", "", "Resources To Import")
	_cloudformationCmd.Flags().StringSliceVarP(&_cloudformationResourcesToSkip, "resources-to-skip", "", nil, "Resources To Skip")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationRetainExceptOnCreate, "retain-except-on-create", "", "", "Retain Except On Create")
	_cloudformationCmd.Flags().StringSliceVarP(&_cloudformationRetainResources, "retain-resources", "", nil, "Retain Resources")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationRetainStacks, "retain-stacks", "", "", "Retain Stacks")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationRoleARN, "role-arn", "", "", "Role ARN")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationRollbackConfiguration, "rollback-configuration", "", "", "Rollback Configuration")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationScanFilters, "scan-filters", "", "", "Scan Filters")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationScanTypeFilter, "scan-type-filter", "", "", "Scan Type Filter")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationSchemaHandlerPackage, "schema-handler-package", "", "", "Schema Handler Package")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationStackDefinitions, "stack-definitions", "", "", "Stack Definitions")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationStackDriftDetectionId, "stack-drift-detection-id", "", "", "Stack Drift Detection ID")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationStackId, "stack-id", "", "", "Stack ID")
	_cloudformationCmd.Flags().StringSliceVarP(&_cloudformationStackIds, "stack-ids", "", nil, "Stack Ids")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationStackIdsUrl, "stack-ids-url", "", "", "Stack Ids URL")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationStackInstanceAccount, "stack-instance-account", "", "", "Stack Instance Account")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationStackInstanceRegion, "stack-instance-region", "", "", "Stack Instance Region")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationStackInstanceResourceDriftStatuses, "stack-instance-resource-drift-statuses", "", "", "Stack Instance Resource Drift Statuses")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationStackName, "stack-name", "", "", "Stack Name")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationStackPolicyBody, "stack-policy-body", "", "", "Stack Policy Body")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationStackPolicyDuringUpdateBody, "stack-policy-during-update-body", "", "", "Stack Policy During Update Body")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationStackPolicyDuringUpdateURL, "stack-policy-during-update-url", "", "", "Stack Policy During Update URL")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationStackPolicyURL, "stack-policy-url", "", "", "Stack Policy URL")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationStackRefactorId, "stack-refactor-id", "", "", "Stack Refactor ID")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationStackResourceDriftStatusFilters, "stack-resource-drift-status-filters", "", "", "Stack Resource Drift Status Filters")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationStackSetName, "stack-set-name", "", "", "Stack Set Name")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationStackStatusFilter, "stack-status-filter", "", "", "Stack Status Filter")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationStatus, "status", "", "", "Status")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationStatusMessage, "status-message", "", "", "Status Message")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationTagKey, "tag-key", "", "", "Tag Key")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationTagValue, "tag-value", "", "", "Tag Value")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationTags, "tags", "", "", "Tags")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationTargetId, "target-id", "", "", "Target ID")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationTargetType, "target-type", "", "", "Target Type")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationTemplateBody, "template-body", "", "", "Template Body")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationTemplateConfiguration, "template-configuration", "", "", "Template Configuration")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationTemplateStage, "template-stage", "", "", "Template Stage")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationTemplateSummaryConfig, "template-summary-config", "", "", "Template Summary Config")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationTemplateURL, "template-url", "", "", "Template URL")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationTimeoutInMinutes, "timeout-in-minutes", "", "", "Timeout In Minutes")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationType, "type", "", "", "Type")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationTypeArn, "type-arn", "", "", "Type ARN")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationTypeConfigurationIdentifiers, "type-configuration-identifiers", "", "", "Type Configuration Identifiers")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationTypeName, "type-name", "", "", "Type Name")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationTypeNameAlias, "type-name-alias", "", "", "Type Name Alias")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationUniqueId, "unique-id", "", "", "Unique ID")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationUsePreviousTemplate, "use-previous-template", "", "", "Use Previous Template")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationVersionBump, "version-bump", "", "", "Version Bump")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationVersionId, "version-id", "", "", "Version ID")
	_cloudformationCmd.Flags().StringVarP(&_cloudformationVisibility, "visibility", "", "", "Visibility")

	_cloudformationCmd.Flags().BoolVarP(&_cloudformationActivateOrganizationsAccess, "activate-organizations-access", "", false, "Activate Organizations Access")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationActivateType, "activate-type", "", false, "Activate Type")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationBatchDescribeTypeConfigurations, "batch-describe-type-configurations", "", false, "Batch Describe Type Configurations")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationCancelUpdateStack, "cancel-update-stack", "", false, "Cancel Update Stack")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationContinueUpdateRollback, "continue-update-rollback", "", false, "Continue Update Rollback")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationCreateChangeSet, "create-change-set", "", false, "Create Change Set")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationCreateGeneratedTemplate, "create-generated-template", "", false, "Create Generated Template")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationCreateStack, "create-stack", "", false, "Create Stack")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationCreateStackInstances, "create-stack-instances", "", false, "Create Stack Instances")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationCreateStackRefactor, "create-stack-refactor", "", false, "Create Stack Refactor")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationCreateStackSet, "create-stack-set", "", false, "Create Stack Set")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationDeactivateOrganizationsAccess, "deactivate-organizations-access", "", false, "Deactivate Organizations Access")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationDeactivateType, "deactivate-type", "", false, "Deactivate Type")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationDeleteChangeSet, "delete-change-set", "", false, "Delete Change Set")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationDeleteGeneratedTemplate, "delete-generated-template", "", false, "Delete Generated Template")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationDeleteStack, "delete-stack", "", false, "Delete Stack")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationDeleteStackInstances, "delete-stack-instances", "", false, "Delete Stack Instances")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationDeleteStackSet, "delete-stack-set", "", false, "Delete Stack Set")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationDeregisterType, "deregister-type", "", false, "Deregister Type")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationDescribeAccountLimits, "describe-account-limits", "", false, "Describe Account Limits")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationDescribeChangeSet, "describe-change-set", "", false, "Describe Change Set")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationDescribeChangeSetHooks, "describe-change-set-hooks", "", false, "Describe Change Set Hooks")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationDescribeEvents, "describe-events", "", false, "Describe Events")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationDescribeGeneratedTemplate, "describe-generated-template", "", false, "Describe Generated Template")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationDescribeOrganizationsAccess, "describe-organizations-access", "", false, "Describe Organizations Access")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationDescribePublisher, "describe-publisher", "", false, "Describe Publisher")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationDescribeResourceScan, "describe-resource-scan", "", false, "Describe Resource Scan")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationDescribeStackDriftDetectionStatus, "describe-stack-drift-detection-status", "", false, "Describe Stack Drift Detection Status")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationDescribeStackEvents, "describe-stack-events", "", false, "Describe Stack Events")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationDescribeStackInstance, "describe-stack-instance", "", false, "Describe Stack Instance")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationDescribeStackRefactor, "describe-stack-refactor", "", false, "Describe Stack Refactor")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationDescribeStackResource, "describe-stack-resource", "", false, "Describe Stack Resource")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationDescribeStackResourceDrifts, "describe-stack-resource-drifts", "", false, "Describe Stack Resource Drifts")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationDescribeStackResources, "describe-stack-resources", "", false, "Describe Stack Resources")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationDescribeStackSet, "describe-stack-set", "", false, "Describe Stack Set")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationDescribeStackSetOperation, "describe-stack-set-operation", "", false, "Describe Stack Set Operation")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationDescribeStacks, "describe-stacks", "", false, "Describe Stacks")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationDescribeType, "describe-type", "", false, "Describe Type")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationDescribeTypeRegistration, "describe-type-registration", "", false, "Describe Type Registration")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationDetectStackDrift, "detect-stack-drift", "", false, "Detect Stack Drift")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationDetectStackResourceDrift, "detect-stack-resource-drift", "", false, "Detect Stack Resource Drift")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationDetectStackSetDrift, "detect-stack-set-drift", "", false, "Detect Stack Set Drift")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationEstimateTemplateCost, "estimate-template-cost", "", false, "Estimate Template Cost")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationExecuteChangeSet, "execute-change-set", "", false, "Execute Change Set")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationExecuteStackRefactor, "execute-stack-refactor", "", false, "Execute Stack Refactor")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationGetGeneratedTemplate, "get-generated-template", "", false, "Get Generated Template")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationGetHookResult, "get-hook-result", "", false, "Get Hook Result")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationGetStackPolicy, "get-stack-policy", "", false, "Get Stack Policy")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationGetTemplate, "get-template", "", false, "Get Template")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationGetTemplateSummary, "get-template-summary", "", false, "Get Template Summary")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationImportStacksToStackSet, "import-stacks-to-stack-set", "", false, "Import Stacks To Stack Set")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationListChangeSets, "list-change-sets", "", false, "List Change Sets")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationListExports, "list-exports", "", false, "List Exports")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationListGeneratedTemplates, "list-generated-templates", "", false, "List Generated Templates")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationListHookResults, "list-hook-results", "", false, "List Hook Results")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationListImports, "list-imports", "", false, "List Imports")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationListResourceScanRelatedResources, "list-resource-scan-related-resources", "", false, "List Resource Scan Related Resources")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationListResourceScanResources, "list-resource-scan-resources", "", false, "List Resource Scan Resources")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationListResourceScans, "list-resource-scans", "", false, "List Resource Scans")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationListStackInstanceResourceDrifts, "list-stack-instance-resource-drifts", "", false, "List Stack Instance Resource Drifts")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationListStackInstances, "list-stack-instances", "", false, "List Stack Instances")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationListStackRefactorActions, "list-stack-refactor-actions", "", false, "List Stack Refactor Actions")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationListStackRefactors, "list-stack-refactors", "", false, "List Stack Refactors")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationListStackResources, "list-stack-resources", "", false, "List Stack Resources")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationListStackSetAutoDeploymentTargets, "list-stack-set-auto-deployment-targets", "", false, "List Stack Set Auto Deployment Targets")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationListStackSetOperationResults, "list-stack-set-operation-results", "", false, "List Stack Set Operation Results")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationListStackSetOperations, "list-stack-set-operations", "", false, "List Stack Set Operations")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationListStackSets, "list-stack-sets", "", false, "List Stack Sets")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationListStacks, "list-stacks", "", false, "List Stacks")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationListTypeRegistrations, "list-type-registrations", "", false, "List Type Registrations")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationListTypeVersions, "list-type-versions", "", false, "List Type Versions")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationListTypes, "list-types", "", false, "List Types")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationPublishType, "publish-type", "", false, "Publish Type")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationRecordHandlerProgress, "record-handler-progress", "", false, "Record Handler Progress")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationRegisterPublisher, "register-publisher", "", false, "Register Publisher")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationRegisterType, "register-type", "", false, "Register Type")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationRollbackStack, "rollback-stack", "", false, "Rollback Stack")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationSetStackPolicy, "set-stack-policy", "", false, "Set Stack Policy")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationSetTypeConfiguration, "set-type-configuration", "", false, "Set Type Configuration")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationSetTypeDefaultVersion, "set-type-default-version", "", false, "Set Type Default Version")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationSignalResource, "signal-resource", "", false, "Signal Resource")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationStartResourceScan, "start-resource-scan", "", false, "Start Resource Scan")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationStopStackSetOperation, "stop-stack-set-operation", "", false, "Stop Stack Set Operation")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationTestType, "test-type", "", false, "Test Type")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationUpdateGeneratedTemplate, "update-generated-template", "", false, "Update Generated Template")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationUpdateStack, "update-stack", "", false, "Update Stack")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationUpdateStackInstances, "update-stack-instances", "", false, "Update Stack Instances")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationUpdateStackSet, "update-stack-set", "", false, "Update Stack Set")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationUpdateTerminationProtection, "update-termination-protection", "", false, "Update Termination Protection")
	_cloudformationCmd.Flags().BoolVarP(&_cloudformationValidateTemplate, "validate-template", "", false, "Validate Template")

}
