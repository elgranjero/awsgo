package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iotsitewise"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// iotsitewiseCmd represents the iotsitewise command
var _iotsitewiseCmd = &cobra.Command{
	Use:   "iotsitewise",
	Short: "AWS iotsitewise CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := iotsitewise.NewFromConfig(cfg)
		if _iotsitewiseAssociateAssets {
			iotsitewise_AssociateAssets(cfg, client)
			return
		}
		if _iotsitewiseAssociateTimeSeriesToAssetProperty {
			iotsitewise_AssociateTimeSeriesToAssetProperty(cfg, client)
			return
		}
		if _iotsitewiseBatchAssociateProjectAssets {
			iotsitewise_BatchAssociateProjectAssets(cfg, client)
			return
		}
		if _iotsitewiseBatchDisassociateProjectAssets {
			iotsitewise_BatchDisassociateProjectAssets(cfg, client)
			return
		}
		if _iotsitewiseBatchGetAssetPropertyAggregates {
			iotsitewise_BatchGetAssetPropertyAggregates(cfg, client)
			return
		}
		if _iotsitewiseBatchGetAssetPropertyValue {
			iotsitewise_BatchGetAssetPropertyValue(cfg, client)
			return
		}
		if _iotsitewiseBatchGetAssetPropertyValueHistory {
			iotsitewise_BatchGetAssetPropertyValueHistory(cfg, client)
			return
		}
		if _iotsitewiseBatchPutAssetPropertyValue {
			iotsitewise_BatchPutAssetPropertyValue(cfg, client)
			return
		}
		if _iotsitewiseCreateAccessPolicy {
			iotsitewise_CreateAccessPolicy(cfg, client)
			return
		}
		if _iotsitewiseCreateAsset {
			iotsitewise_CreateAsset(cfg, client)
			return
		}
		if _iotsitewiseCreateAssetModel {
			iotsitewise_CreateAssetModel(cfg, client)
			return
		}
		if _iotsitewiseCreateAssetModelCompositeModel {
			iotsitewise_CreateAssetModelCompositeModel(cfg, client)
			return
		}
		if _iotsitewiseCreateBulkImportJob {
			iotsitewise_CreateBulkImportJob(cfg, client)
			return
		}
		if _iotsitewiseCreateComputationModel {
			iotsitewise_CreateComputationModel(cfg, client)
			return
		}
		if _iotsitewiseCreateDashboard {
			iotsitewise_CreateDashboard(cfg, client)
			return
		}
		if _iotsitewiseCreateDataset {
			iotsitewise_CreateDataset(cfg, client)
			return
		}
		if _iotsitewiseCreateGateway {
			iotsitewise_CreateGateway(cfg, client)
			return
		}
		if _iotsitewiseCreatePortal {
			iotsitewise_CreatePortal(cfg, client)
			return
		}
		if _iotsitewiseCreateProject {
			iotsitewise_CreateProject(cfg, client)
			return
		}
		if _iotsitewiseDeleteAccessPolicy {
			iotsitewise_DeleteAccessPolicy(cfg, client)
			return
		}
		if _iotsitewiseDeleteAsset {
			iotsitewise_DeleteAsset(cfg, client)
			return
		}
		if _iotsitewiseDeleteAssetModel {
			iotsitewise_DeleteAssetModel(cfg, client)
			return
		}
		if _iotsitewiseDeleteAssetModelCompositeModel {
			iotsitewise_DeleteAssetModelCompositeModel(cfg, client)
			return
		}
		if _iotsitewiseDeleteAssetModelInterfaceRelationship {
			iotsitewise_DeleteAssetModelInterfaceRelationship(cfg, client)
			return
		}
		if _iotsitewiseDeleteComputationModel {
			iotsitewise_DeleteComputationModel(cfg, client)
			return
		}
		if _iotsitewiseDeleteDashboard {
			iotsitewise_DeleteDashboard(cfg, client)
			return
		}
		if _iotsitewiseDeleteDataset {
			iotsitewise_DeleteDataset(cfg, client)
			return
		}
		if _iotsitewiseDeleteGateway {
			iotsitewise_DeleteGateway(cfg, client)
			return
		}
		if _iotsitewiseDeletePortal {
			iotsitewise_DeletePortal(cfg, client)
			return
		}
		if _iotsitewiseDeleteProject {
			iotsitewise_DeleteProject(cfg, client)
			return
		}
		if _iotsitewiseDeleteTimeSeries {
			iotsitewise_DeleteTimeSeries(cfg, client)
			return
		}
		if _iotsitewiseDescribeAccessPolicy {
			iotsitewise_DescribeAccessPolicy(cfg, client)
			return
		}
		if _iotsitewiseDescribeAction {
			iotsitewise_DescribeAction(cfg, client)
			return
		}
		if _iotsitewiseDescribeAsset {
			iotsitewise_DescribeAsset(cfg, client)
			return
		}
		if _iotsitewiseDescribeAssetCompositeModel {
			iotsitewise_DescribeAssetCompositeModel(cfg, client)
			return
		}
		if _iotsitewiseDescribeAssetModel {
			iotsitewise_DescribeAssetModel(cfg, client)
			return
		}
		if _iotsitewiseDescribeAssetModelCompositeModel {
			iotsitewise_DescribeAssetModelCompositeModel(cfg, client)
			return
		}
		if _iotsitewiseDescribeAssetModelInterfaceRelationship {
			iotsitewise_DescribeAssetModelInterfaceRelationship(cfg, client)
			return
		}
		if _iotsitewiseDescribeAssetProperty {
			iotsitewise_DescribeAssetProperty(cfg, client)
			return
		}
		if _iotsitewiseDescribeBulkImportJob {
			iotsitewise_DescribeBulkImportJob(cfg, client)
			return
		}
		if _iotsitewiseDescribeComputationModel {
			iotsitewise_DescribeComputationModel(cfg, client)
			return
		}
		if _iotsitewiseDescribeComputationModelExecutionSummary {
			iotsitewise_DescribeComputationModelExecutionSummary(cfg, client)
			return
		}
		if _iotsitewiseDescribeDashboard {
			iotsitewise_DescribeDashboard(cfg, client)
			return
		}
		if _iotsitewiseDescribeDataset {
			iotsitewise_DescribeDataset(cfg, client)
			return
		}
		if _iotsitewiseDescribeDefaultEncryptionConfiguration {
			iotsitewise_DescribeDefaultEncryptionConfiguration(cfg, client)
			return
		}
		if _iotsitewiseDescribeExecution {
			iotsitewise_DescribeExecution(cfg, client)
			return
		}
		if _iotsitewiseDescribeGateway {
			iotsitewise_DescribeGateway(cfg, client)
			return
		}
		if _iotsitewiseDescribeGatewayCapabilityConfiguration {
			iotsitewise_DescribeGatewayCapabilityConfiguration(cfg, client)
			return
		}
		if _iotsitewiseDescribeLoggingOptions {
			iotsitewise_DescribeLoggingOptions(cfg, client)
			return
		}
		if _iotsitewiseDescribePortal {
			iotsitewise_DescribePortal(cfg, client)
			return
		}
		if _iotsitewiseDescribeProject {
			iotsitewise_DescribeProject(cfg, client)
			return
		}
		if _iotsitewiseDescribeStorageConfiguration {
			iotsitewise_DescribeStorageConfiguration(cfg, client)
			return
		}
		if _iotsitewiseDescribeTimeSeries {
			iotsitewise_DescribeTimeSeries(cfg, client)
			return
		}
		if _iotsitewiseDisassociateAssets {
			iotsitewise_DisassociateAssets(cfg, client)
			return
		}
		if _iotsitewiseDisassociateTimeSeriesFromAssetProperty {
			iotsitewise_DisassociateTimeSeriesFromAssetProperty(cfg, client)
			return
		}
		if _iotsitewiseExecuteAction {
			iotsitewise_ExecuteAction(cfg, client)
			return
		}
		if _iotsitewiseExecuteQuery {
			iotsitewise_ExecuteQuery(cfg, client)
			return
		}
		if _iotsitewiseGetAssetPropertyAggregates {
			iotsitewise_GetAssetPropertyAggregates(cfg, client)
			return
		}
		if _iotsitewiseGetAssetPropertyValue {
			iotsitewise_GetAssetPropertyValue(cfg, client)
			return
		}
		if _iotsitewiseGetAssetPropertyValueHistory {
			iotsitewise_GetAssetPropertyValueHistory(cfg, client)
			return
		}
		if _iotsitewiseGetInterpolatedAssetPropertyValues {
			iotsitewise_GetInterpolatedAssetPropertyValues(cfg, client)
			return
		}
		if _iotsitewiseInvokeAssistant {
			iotsitewise_InvokeAssistant(cfg, client)
			return
		}
		if _iotsitewiseListAccessPolicies {
			iotsitewise_ListAccessPolicies(cfg, client)
			return
		}
		if _iotsitewiseListActions {
			iotsitewise_ListActions(cfg, client)
			return
		}
		if _iotsitewiseListAssetModelCompositeModels {
			iotsitewise_ListAssetModelCompositeModels(cfg, client)
			return
		}
		if _iotsitewiseListAssetModelProperties {
			iotsitewise_ListAssetModelProperties(cfg, client)
			return
		}
		if _iotsitewiseListAssetModels {
			iotsitewise_ListAssetModels(cfg, client)
			return
		}
		if _iotsitewiseListAssetProperties {
			iotsitewise_ListAssetProperties(cfg, client)
			return
		}
		if _iotsitewiseListAssetRelationships {
			iotsitewise_ListAssetRelationships(cfg, client)
			return
		}
		if _iotsitewiseListAssets {
			iotsitewise_ListAssets(cfg, client)
			return
		}
		if _iotsitewiseListAssociatedAssets {
			iotsitewise_ListAssociatedAssets(cfg, client)
			return
		}
		if _iotsitewiseListBulkImportJobs {
			iotsitewise_ListBulkImportJobs(cfg, client)
			return
		}
		if _iotsitewiseListCompositionRelationships {
			iotsitewise_ListCompositionRelationships(cfg, client)
			return
		}
		if _iotsitewiseListComputationModelDataBindingUsages {
			iotsitewise_ListComputationModelDataBindingUsages(cfg, client)
			return
		}
		if _iotsitewiseListComputationModelResolveToResources {
			iotsitewise_ListComputationModelResolveToResources(cfg, client)
			return
		}
		if _iotsitewiseListComputationModels {
			iotsitewise_ListComputationModels(cfg, client)
			return
		}
		if _iotsitewiseListDashboards {
			iotsitewise_ListDashboards(cfg, client)
			return
		}
		if _iotsitewiseListDatasets {
			iotsitewise_ListDatasets(cfg, client)
			return
		}
		if _iotsitewiseListExecutions {
			iotsitewise_ListExecutions(cfg, client)
			return
		}
		if _iotsitewiseListGateways {
			iotsitewise_ListGateways(cfg, client)
			return
		}
		if _iotsitewiseListInterfaceRelationships {
			iotsitewise_ListInterfaceRelationships(cfg, client)
			return
		}
		if _iotsitewiseListPortals {
			iotsitewise_ListPortals(cfg, client)
			return
		}
		if _iotsitewiseListProjectAssets {
			iotsitewise_ListProjectAssets(cfg, client)
			return
		}
		if _iotsitewiseListProjects {
			iotsitewise_ListProjects(cfg, client)
			return
		}
		if _iotsitewiseListTagsForResource {
			iotsitewise_ListTagsForResource(cfg, client)
			return
		}
		if _iotsitewiseListTimeSeries {
			iotsitewise_ListTimeSeries(cfg, client)
			return
		}
		if _iotsitewisePutAssetModelInterfaceRelationship {
			iotsitewise_PutAssetModelInterfaceRelationship(cfg, client)
			return
		}
		if _iotsitewisePutDefaultEncryptionConfiguration {
			iotsitewise_PutDefaultEncryptionConfiguration(cfg, client)
			return
		}
		if _iotsitewisePutLoggingOptions {
			iotsitewise_PutLoggingOptions(cfg, client)
			return
		}
		if _iotsitewisePutStorageConfiguration {
			iotsitewise_PutStorageConfiguration(cfg, client)
			return
		}
		if _iotsitewiseTagResource {
			iotsitewise_TagResource(cfg, client)
			return
		}
		if _iotsitewiseUntagResource {
			iotsitewise_UntagResource(cfg, client)
			return
		}
		if _iotsitewiseUpdateAccessPolicy {
			iotsitewise_UpdateAccessPolicy(cfg, client)
			return
		}
		if _iotsitewiseUpdateAsset {
			iotsitewise_UpdateAsset(cfg, client)
			return
		}
		if _iotsitewiseUpdateAssetModel {
			iotsitewise_UpdateAssetModel(cfg, client)
			return
		}
		if _iotsitewiseUpdateAssetModelCompositeModel {
			iotsitewise_UpdateAssetModelCompositeModel(cfg, client)
			return
		}
		if _iotsitewiseUpdateAssetProperty {
			iotsitewise_UpdateAssetProperty(cfg, client)
			return
		}
		if _iotsitewiseUpdateComputationModel {
			iotsitewise_UpdateComputationModel(cfg, client)
			return
		}
		if _iotsitewiseUpdateDashboard {
			iotsitewise_UpdateDashboard(cfg, client)
			return
		}
		if _iotsitewiseUpdateDataset {
			iotsitewise_UpdateDataset(cfg, client)
			return
		}
		if _iotsitewiseUpdateGateway {
			iotsitewise_UpdateGateway(cfg, client)
			return
		}
		if _iotsitewiseUpdateGatewayCapabilityConfiguration {
			iotsitewise_UpdateGatewayCapabilityConfiguration(cfg, client)
			return
		}
		if _iotsitewiseUpdatePortal {
			iotsitewise_UpdatePortal(cfg, client)
			return
		}
		if _iotsitewiseUpdateProject {
			iotsitewise_UpdateProject(cfg, client)
			return
		}

	},
}

var (
	_iotsitewiseAssociateAssets                          bool
	_iotsitewiseAssociateTimeSeriesToAssetProperty       bool
	_iotsitewiseBatchAssociateProjectAssets              bool
	_iotsitewiseBatchDisassociateProjectAssets           bool
	_iotsitewiseBatchGetAssetPropertyAggregates          bool
	_iotsitewiseBatchGetAssetPropertyValue               bool
	_iotsitewiseBatchGetAssetPropertyValueHistory        bool
	_iotsitewiseBatchPutAssetPropertyValue               bool
	_iotsitewiseCreateAccessPolicy                       bool
	_iotsitewiseCreateAsset                              bool
	_iotsitewiseCreateAssetModel                         bool
	_iotsitewiseCreateAssetModelCompositeModel           bool
	_iotsitewiseCreateBulkImportJob                      bool
	_iotsitewiseCreateComputationModel                   bool
	_iotsitewiseCreateDashboard                          bool
	_iotsitewiseCreateDataset                            bool
	_iotsitewiseCreateGateway                            bool
	_iotsitewiseCreatePortal                             bool
	_iotsitewiseCreateProject                            bool
	_iotsitewiseDeleteAccessPolicy                       bool
	_iotsitewiseDeleteAsset                              bool
	_iotsitewiseDeleteAssetModel                         bool
	_iotsitewiseDeleteAssetModelCompositeModel           bool
	_iotsitewiseDeleteAssetModelInterfaceRelationship    bool
	_iotsitewiseDeleteComputationModel                   bool
	_iotsitewiseDeleteDashboard                          bool
	_iotsitewiseDeleteDataset                            bool
	_iotsitewiseDeleteGateway                            bool
	_iotsitewiseDeletePortal                             bool
	_iotsitewiseDeleteProject                            bool
	_iotsitewiseDeleteTimeSeries                         bool
	_iotsitewiseDescribeAccessPolicy                     bool
	_iotsitewiseDescribeAction                           bool
	_iotsitewiseDescribeAsset                            bool
	_iotsitewiseDescribeAssetCompositeModel              bool
	_iotsitewiseDescribeAssetModel                       bool
	_iotsitewiseDescribeAssetModelCompositeModel         bool
	_iotsitewiseDescribeAssetModelInterfaceRelationship  bool
	_iotsitewiseDescribeAssetProperty                    bool
	_iotsitewiseDescribeBulkImportJob                    bool
	_iotsitewiseDescribeComputationModel                 bool
	_iotsitewiseDescribeComputationModelExecutionSummary bool
	_iotsitewiseDescribeDashboard                        bool
	_iotsitewiseDescribeDataset                          bool
	_iotsitewiseDescribeDefaultEncryptionConfiguration   bool
	_iotsitewiseDescribeExecution                        bool
	_iotsitewiseDescribeGateway                          bool
	_iotsitewiseDescribeGatewayCapabilityConfiguration   bool
	_iotsitewiseDescribeLoggingOptions                   bool
	_iotsitewiseDescribePortal                           bool
	_iotsitewiseDescribeProject                          bool
	_iotsitewiseDescribeStorageConfiguration             bool
	_iotsitewiseDescribeTimeSeries                       bool
	_iotsitewiseDisassociateAssets                       bool
	_iotsitewiseDisassociateTimeSeriesFromAssetProperty  bool
	_iotsitewiseExecuteAction                            bool
	_iotsitewiseExecuteQuery                             bool
	_iotsitewiseGetAssetPropertyAggregates               bool
	_iotsitewiseGetAssetPropertyValue                    bool
	_iotsitewiseGetAssetPropertyValueHistory             bool
	_iotsitewiseGetInterpolatedAssetPropertyValues       bool
	_iotsitewiseInvokeAssistant                          bool
	_iotsitewiseListAccessPolicies                       bool
	_iotsitewiseListActions                              bool
	_iotsitewiseListAssetModelCompositeModels            bool
	_iotsitewiseListAssetModelProperties                 bool
	_iotsitewiseListAssetModels                          bool
	_iotsitewiseListAssetProperties                      bool
	_iotsitewiseListAssetRelationships                   bool
	_iotsitewiseListAssets                               bool
	_iotsitewiseListAssociatedAssets                     bool
	_iotsitewiseListBulkImportJobs                       bool
	_iotsitewiseListCompositionRelationships             bool
	_iotsitewiseListComputationModelDataBindingUsages    bool
	_iotsitewiseListComputationModelResolveToResources   bool
	_iotsitewiseListComputationModels                    bool
	_iotsitewiseListDashboards                           bool
	_iotsitewiseListDatasets                             bool
	_iotsitewiseListExecutions                           bool
	_iotsitewiseListGateways                             bool
	_iotsitewiseListInterfaceRelationships               bool
	_iotsitewiseListPortals                              bool
	_iotsitewiseListProjectAssets                        bool
	_iotsitewiseListProjects                             bool
	_iotsitewiseListTagsForResource                      bool
	_iotsitewiseListTimeSeries                           bool
	_iotsitewisePutAssetModelInterfaceRelationship       bool
	_iotsitewisePutDefaultEncryptionConfiguration        bool
	_iotsitewisePutLoggingOptions                        bool
	_iotsitewisePutStorageConfiguration                  bool
	_iotsitewiseTagResource                              bool
	_iotsitewiseUntagResource                            bool
	_iotsitewiseUpdateAccessPolicy                       bool
	_iotsitewiseUpdateAsset                              bool
	_iotsitewiseUpdateAssetModel                         bool
	_iotsitewiseUpdateAssetModelCompositeModel           bool
	_iotsitewiseUpdateAssetProperty                      bool
	_iotsitewiseUpdateComputationModel                   bool
	_iotsitewiseUpdateDashboard                          bool
	_iotsitewiseUpdateDataset                            bool
	_iotsitewiseUpdateGateway                            bool
	_iotsitewiseUpdateGatewayCapabilityConfiguration     bool
	_iotsitewiseUpdatePortal                             bool
	_iotsitewiseUpdateProject                            bool

	_iotsitewiseAccessPolicyId                      string
	_iotsitewiseAccessPolicyIdentity                string
	_iotsitewiseAccessPolicyPermission              string
	_iotsitewiseAccessPolicyResource                string
	_iotsitewiseActionDefinitionId                  string
	_iotsitewiseActionId                            string
	_iotsitewiseActionPayload                       string
	_iotsitewiseActionType                          string
	_iotsitewiseAdaptiveIngestion                   string
	_iotsitewiseAggregateTypes                      string
	_iotsitewiseAlarms                              string
	_iotsitewiseAlias                               string
	_iotsitewiseAliasPrefix                         string
	_iotsitewiseAssetCompositeModelId               string
	_iotsitewiseAssetDescription                    string
	_iotsitewiseAssetExternalId                     string
	_iotsitewiseAssetId                             string
	_iotsitewiseAssetIds                            []string
	_iotsitewiseAssetModelCompositeModelDescription string
	_iotsitewiseAssetModelCompositeModelExternalId  string
	_iotsitewiseAssetModelCompositeModelId          string
	_iotsitewiseAssetModelCompositeModelName        string
	_iotsitewiseAssetModelCompositeModelProperties  string
	_iotsitewiseAssetModelCompositeModelType        string
	_iotsitewiseAssetModelCompositeModels           string
	_iotsitewiseAssetModelDescription               string
	_iotsitewiseAssetModelExternalId                string
	_iotsitewiseAssetModelHierarchies               string
	_iotsitewiseAssetModelId                        string
	_iotsitewiseAssetModelName                      string
	_iotsitewiseAssetModelProperties                string
	_iotsitewiseAssetModelType                      string
	_iotsitewiseAssetModelTypes                     string
	_iotsitewiseAssetModelVersion                   string
	_iotsitewiseAssetName                           string
	_iotsitewiseCapabilityConfiguration             string
	_iotsitewiseCapabilityNamespace                 string
	_iotsitewiseChildAssetId                        string
	_iotsitewiseClientToken                         string
	_iotsitewiseComposedAssetModelId                string
	_iotsitewiseComputationModelConfiguration       string
	_iotsitewiseComputationModelDataBinding         string
	_iotsitewiseComputationModelDescription         string
	_iotsitewiseComputationModelId                  string
	_iotsitewiseComputationModelName                string
	_iotsitewiseComputationModelType                string
	_iotsitewiseComputationModelVersion             string
	_iotsitewiseConversationId                      string
	_iotsitewiseDashboardDefinition                 string
	_iotsitewiseDashboardDescription                string
	_iotsitewiseDashboardId                         string
	_iotsitewiseDashboardName                       string
	_iotsitewiseDataBindingValueFilter              string
	_iotsitewiseDatasetDescription                  string
	_iotsitewiseDatasetId                           string
	_iotsitewiseDatasetName                         string
	_iotsitewiseDatasetSource                       string
	_iotsitewiseDeleteFilesAfterImport              string
	_iotsitewiseDisallowIngestNullNaN               string
	_iotsitewiseDisassociatedDataStorage            string
	_iotsitewiseEnablePartialEntryProcessing        string
	_iotsitewiseEnableTrace                         string
	_iotsitewiseEncryptionType                      string
	_iotsitewiseEndDate                             string
	_iotsitewiseEndTimeInSeconds                    string
	_iotsitewiseEndTimeOffsetInNanos                string
	_iotsitewiseEntries                             string
	_iotsitewiseErrorReportLocation                 string
	_iotsitewiseExcludeProperties                   string
	_iotsitewiseExecutionId                         string
	_iotsitewiseFiles                               string
	_iotsitewiseFilter                              string
	_iotsitewiseGatewayId                           string
	_iotsitewiseGatewayName                         string
	_iotsitewiseGatewayPlatform                     string
	_iotsitewiseGatewayVersion                      string
	_iotsitewiseHierarchyId                         string
	_iotsitewiseIamArn                              string
	_iotsitewiseIdentityId                          string
	_iotsitewiseIdentityType                        string
	_iotsitewiseIfMatch                             string
	_iotsitewiseIfNoneMatch                         string
	_iotsitewiseInterfaceAssetModelId               string
	_iotsitewiseIntervalInSeconds                   string
	_iotsitewiseIntervalWindowInSeconds             string
	_iotsitewiseJobConfiguration                    string
	_iotsitewiseJobId                               string
	_iotsitewiseJobName                             string
	_iotsitewiseJobRoleArn                          string
	_iotsitewiseKmsKeyId                            string
	_iotsitewiseLoggingOptions                      string
	_iotsitewiseMatchForVersionType                 string
	_iotsitewiseMaxResults                          string
	_iotsitewiseMessage                             string
	_iotsitewiseMultiLayerStorage                   string
	_iotsitewiseNextToken                           string
	_iotsitewiseNotificationSenderEmail             string
	_iotsitewiseParentAssetModelCompositeModelId    string
	_iotsitewisePortalAuthMode                      string
	_iotsitewisePortalContactEmail                  string
	_iotsitewisePortalDescription                   string
	_iotsitewisePortalId                            string
	_iotsitewisePortalLogoImage                     string
	_iotsitewisePortalLogoImageFile                 string
	_iotsitewisePortalName                          string
	_iotsitewisePortalType                          string
	_iotsitewisePortalTypeConfiguration             string
	_iotsitewiseProjectDescription                  string
	_iotsitewiseProjectId                           string
	_iotsitewiseProjectName                         string
	_iotsitewisePropertyAlias                       string
	_iotsitewisePropertyId                          string
	_iotsitewisePropertyMappingConfiguration        string
	_iotsitewisePropertyNotificationState           string
	_iotsitewisePropertyUnit                        string
	_iotsitewiseQualities                           string
	_iotsitewiseQuality                             string
	_iotsitewiseQueryStatement                      string
	_iotsitewiseResolution                          string
	_iotsitewiseResolveTo                           string
	_iotsitewiseResolveToResourceId                 string
	_iotsitewiseResolveToResourceType               string
	_iotsitewiseResourceArn                         string
	_iotsitewiseResourceId                          string
	_iotsitewiseResourceType                        string
	_iotsitewiseRetentionPeriod                     string
	_iotsitewiseRoleArn                             string
	_iotsitewiseSourceType                          string
	_iotsitewiseStartDate                           string
	_iotsitewiseStartTimeInSeconds                  string
	_iotsitewiseStartTimeOffsetInNanos              string
	_iotsitewiseStorageType                         string
	_iotsitewiseTagKeys                             []string
	_iotsitewiseTags                                string
	_iotsitewiseTargetResource                      string
	_iotsitewiseTargetResourceId                    string
	_iotsitewiseTargetResourceType                  string
	_iotsitewiseTimeOrdering                        string
	_iotsitewiseTimeSeriesType                      string
	_iotsitewiseTraversalDirection                  string
	_iotsitewiseTraversalType                       string
	_iotsitewiseType                                string
	_iotsitewiseWarmTier                            string
	_iotsitewiseWarmTierRetentionPeriod             string
)

// Associates a child asset with the given parent asset through a hierarchy
// defined in the parent asset's model. For more information, see [Associating assets]in the IoT
// SiteWise User Guide.
//
// [Associating assets]: https://docs.aws.amazon.com/iot-sitewise/latest/userguide/add-associated-assets.html
func iotsitewise_AssociateAssets(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.AssociateAssetsInput{
		// AssetId: *string, // Required
		// ChildAssetId: *string, // Required
		// HierarchyId: *string, // Required
	}

	if len(_iotsitewiseAssetId) > 0 {
		input.AssetId = aws.String(_iotsitewiseAssetId)
	}
	if len(_iotsitewiseChildAssetId) > 0 {
		input.ChildAssetId = aws.String(_iotsitewiseChildAssetId)
	}
	if len(_iotsitewiseHierarchyId) > 0 {
		input.HierarchyId = aws.String(_iotsitewiseHierarchyId)
	}
	if len(_iotsitewiseClientToken) > 0 {
		input.ClientToken = aws.String(_iotsitewiseClientToken)
	}

	if resp, err := client.AssociateAssets(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates a time series (data stream) with an asset property.
func iotsitewise_AssociateTimeSeriesToAssetProperty(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.AssociateTimeSeriesToAssetPropertyInput{
		// Alias: *string, // Required
		// AssetId: *string, // Required
		// PropertyId: *string, // Required
	}

	if len(_iotsitewiseAlias) > 0 {
		input.Alias = aws.String(_iotsitewiseAlias)
	}
	if len(_iotsitewiseAssetId) > 0 {
		input.AssetId = aws.String(_iotsitewiseAssetId)
	}
	if len(_iotsitewisePropertyId) > 0 {
		input.PropertyId = aws.String(_iotsitewisePropertyId)
	}
	if len(_iotsitewiseClientToken) > 0 {
		input.ClientToken = aws.String(_iotsitewiseClientToken)
	}

	if resp, err := client.AssociateTimeSeriesToAssetProperty(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates a group (batch) of assets with an IoT SiteWise Monitor project.
func iotsitewise_BatchAssociateProjectAssets(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.BatchAssociateProjectAssetsInput{
		// AssetIds: []string, // Required
		// ProjectId: *string, // Required
	}

	if len(_iotsitewiseAssetIds) > 0 {
		input.AssetIds = append([]string(nil), _iotsitewiseAssetIds...)
	}
	if len(_iotsitewiseProjectId) > 0 {
		input.ProjectId = aws.String(_iotsitewiseProjectId)
	}
	if len(_iotsitewiseClientToken) > 0 {
		input.ClientToken = aws.String(_iotsitewiseClientToken)
	}

	if resp, err := client.BatchAssociateProjectAssets(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates a group (batch) of assets from an IoT SiteWise Monitor project.
func iotsitewise_BatchDisassociateProjectAssets(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.BatchDisassociateProjectAssetsInput{
		// AssetIds: []string, // Required
		// ProjectId: *string, // Required
	}

	if len(_iotsitewiseAssetIds) > 0 {
		input.AssetIds = append([]string(nil), _iotsitewiseAssetIds...)
	}
	if len(_iotsitewiseProjectId) > 0 {
		input.ProjectId = aws.String(_iotsitewiseProjectId)
	}
	if len(_iotsitewiseClientToken) > 0 {
		input.ClientToken = aws.String(_iotsitewiseClientToken)
	}

	if resp, err := client.BatchDisassociateProjectAssets(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets aggregated values (for example, average, minimum, and maximum) for one or
// more asset properties. For more information, see [Querying aggregates]in the IoT SiteWise User Guide.
//
// [Querying aggregates]: https://docs.aws.amazon.com/iot-sitewise/latest/userguide/query-industrial-data.html#aggregates
func iotsitewise_BatchGetAssetPropertyAggregates(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.BatchGetAssetPropertyAggregatesInput{
		// Entries: []types.BatchGetAssetPropertyAggregatesEntry, // Required
	}

	if len(_iotsitewiseEntries) > 0 {
		if err := assignInputField(input, "Entries", _iotsitewiseEntries); err != nil {
			log.Errorf("invalid --entries: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotsitewiseMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseNextToken) > 0 {
		input.NextToken = aws.String(_iotsitewiseNextToken)
	}

	if disablePaginator() {
		if resp, err := client.BatchGetAssetPropertyAggregates(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotsitewise.BatchGetAssetPropertyAggregatesOutput
	p := iotsitewise.NewBatchGetAssetPropertyAggregatesPaginator(client, input)
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

// Gets the current value for one or more asset properties. For more information,
// see [Querying current values]in the IoT SiteWise User Guide.
//
// [Querying current values]: https://docs.aws.amazon.com/iot-sitewise/latest/userguide/query-industrial-data.html#current-values
func iotsitewise_BatchGetAssetPropertyValue(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.BatchGetAssetPropertyValueInput{
		// Entries: []types.BatchGetAssetPropertyValueEntry, // Required
	}

	if len(_iotsitewiseEntries) > 0 {
		if err := assignInputField(input, "Entries", _iotsitewiseEntries); err != nil {
			log.Errorf("invalid --entries: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseNextToken) > 0 {
		input.NextToken = aws.String(_iotsitewiseNextToken)
	}

	if disablePaginator() {
		if resp, err := client.BatchGetAssetPropertyValue(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotsitewise.BatchGetAssetPropertyValueOutput
	p := iotsitewise.NewBatchGetAssetPropertyValuePaginator(client, input)
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

// Gets the historical values for one or more asset properties. For more
// information, see [Querying historical values]in the IoT SiteWise User Guide.
//
// [Querying historical values]: https://docs.aws.amazon.com/iot-sitewise/latest/userguide/query-industrial-data.html#historical-values
func iotsitewise_BatchGetAssetPropertyValueHistory(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.BatchGetAssetPropertyValueHistoryInput{
		// Entries: []types.BatchGetAssetPropertyValueHistoryEntry, // Required
	}

	if len(_iotsitewiseEntries) > 0 {
		if err := assignInputField(input, "Entries", _iotsitewiseEntries); err != nil {
			log.Errorf("invalid --entries: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotsitewiseMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseNextToken) > 0 {
		input.NextToken = aws.String(_iotsitewiseNextToken)
	}

	if disablePaginator() {
		if resp, err := client.BatchGetAssetPropertyValueHistory(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotsitewise.BatchGetAssetPropertyValueHistoryOutput
	p := iotsitewise.NewBatchGetAssetPropertyValueHistoryPaginator(client, input)
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

// Sends a list of asset property values to IoT SiteWise. Each value is a
// timestamp-quality-value (TQV) data point. For more information, see [Ingesting data using the API]in the IoT
// SiteWise User Guide.
//
// To identify an asset property, you must specify one of the following:
//
// - The assetId and propertyId of an asset property.
//
// - A propertyAlias , which is a data stream alias (for example,
// /company/windfarm/3/turbine/7/temperature ). To define an asset property's
// alias, see [UpdateAssetProperty].
//
// With respect to Unix epoch time, IoT SiteWise accepts only TQVs that have a
// timestamp of no more than 7 days in the past and no more than 10 minutes in the
// future. IoT SiteWise rejects timestamps outside of the inclusive range of [-7
// days, +10 minutes] and returns a TimestampOutOfRangeException error.
//
// For each asset property, IoT SiteWise overwrites TQVs with duplicate timestamps
// unless the newer TQV has a different quality. For example, if you store a TQV
// {T1, GOOD, V1} , then storing {T1, GOOD, V2} replaces the existing TQV.
//
// IoT SiteWise authorizes access to each BatchPutAssetPropertyValue entry
// individually. For more information, see [BatchPutAssetPropertyValue authorization]in the IoT SiteWise User Guide.
//
// [Ingesting data using the API]: https://docs.aws.amazon.com/iot-sitewise/latest/userguide/ingest-api.html
// [BatchPutAssetPropertyValue authorization]: https://docs.aws.amazon.com/iot-sitewise/latest/userguide/security_iam_service-with-iam.html#security_iam_service-with-iam-id-based-policies-batchputassetpropertyvalue-action
// [UpdateAssetProperty]: https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_UpdateAssetProperty.html
func iotsitewise_BatchPutAssetPropertyValue(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.BatchPutAssetPropertyValueInput{
		// Entries: []types.PutAssetPropertyValueEntry, // Required
	}

	if len(_iotsitewiseEntries) > 0 {
		if err := assignInputField(input, "Entries", _iotsitewiseEntries); err != nil {
			log.Errorf("invalid --entries: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseEnablePartialEntryProcessing) > 0 {
		if err := assignInputField(input, "EnablePartialEntryProcessing", _iotsitewiseEnablePartialEntryProcessing); err != nil {
			log.Errorf("invalid --enable-partial-entry-processing: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchPutAssetPropertyValue(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an access policy that grants the specified identity (IAM Identity
// Center user, IAM Identity Center group, or IAM user) access to the specified IoT
// SiteWise Monitor portal or project resource.
//
// Support for access policies that use an SSO Group as the identity is not
// supported at this time.
func iotsitewise_CreateAccessPolicy(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.CreateAccessPolicyInput{
		// AccessPolicyIdentity: *types.Identity, // Required
		// AccessPolicyPermission: types.Permission, // Required
		// AccessPolicyResource: *types.Resource, // Required
	}

	if len(_iotsitewiseAccessPolicyIdentity) > 0 {
		if err := assignInputField(input, "AccessPolicyIdentity", _iotsitewiseAccessPolicyIdentity); err != nil {
			log.Errorf("invalid --access-policy-identity: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseAccessPolicyPermission) > 0 {
		if err := assignInputField(input, "AccessPolicyPermission", _iotsitewiseAccessPolicyPermission); err != nil {
			log.Errorf("invalid --access-policy-permission: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseAccessPolicyResource) > 0 {
		if err := assignInputField(input, "AccessPolicyResource", _iotsitewiseAccessPolicyResource); err != nil {
			log.Errorf("invalid --access-policy-resource: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseClientToken) > 0 {
		input.ClientToken = aws.String(_iotsitewiseClientToken)
	}
	if len(_iotsitewiseTags) > 0 {
		if err := assignInputField(input, "Tags", _iotsitewiseTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAccessPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an asset from an existing asset model. For more information, see [Creating assets] in
// the IoT SiteWise User Guide.
//
// [Creating assets]: https://docs.aws.amazon.com/iot-sitewise/latest/userguide/create-assets.html
func iotsitewise_CreateAsset(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.CreateAssetInput{
		// AssetModelId: *string, // Required
		// AssetName: *string, // Required
	}

	if len(_iotsitewiseAssetModelId) > 0 {
		input.AssetModelId = aws.String(_iotsitewiseAssetModelId)
	}
	if len(_iotsitewiseAssetName) > 0 {
		input.AssetName = aws.String(_iotsitewiseAssetName)
	}
	if len(_iotsitewiseAssetDescription) > 0 {
		input.AssetDescription = aws.String(_iotsitewiseAssetDescription)
	}
	if len(_iotsitewiseAssetExternalId) > 0 {
		input.AssetExternalId = aws.String(_iotsitewiseAssetExternalId)
	}
	if len(_iotsitewiseAssetId) > 0 {
		input.AssetId = aws.String(_iotsitewiseAssetId)
	}
	if len(_iotsitewiseClientToken) > 0 {
		input.ClientToken = aws.String(_iotsitewiseClientToken)
	}
	if len(_iotsitewiseTags) > 0 {
		if err := assignInputField(input, "Tags", _iotsitewiseTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAsset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an asset model from specified property and hierarchy definitions. You
// create assets from asset models. With asset models, you can easily create assets
// of the same type that have standardized definitions. Each asset created from a
// model inherits the asset model's property and hierarchy definitions. For more
// information, see [Defining asset models]in the IoT SiteWise User Guide.
//
// You can create three types of asset models, ASSET_MODEL , COMPONENT_MODEL , or
// an INTERFACE .
//
// - ASSET_MODEL – (default) An asset model that you can use to create assets.
// Can't be included as a component in another asset model.
//
// - COMPONENT_MODEL – A reusable component that you can include in the
// composite models of other asset models. You can't create assets directly from
// this type of asset model.
//
// - INTERFACE – An interface is a type of model that defines a standard
// structure that can be applied to different asset models.
//
// [Defining asset models]: https://docs.aws.amazon.com/iot-sitewise/latest/userguide/define-models.html
func iotsitewise_CreateAssetModel(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.CreateAssetModelInput{
		// AssetModelName: *string, // Required
	}

	if len(_iotsitewiseAssetModelName) > 0 {
		input.AssetModelName = aws.String(_iotsitewiseAssetModelName)
	}
	if len(_iotsitewiseAssetModelCompositeModels) > 0 {
		if err := assignInputField(input, "AssetModelCompositeModels", _iotsitewiseAssetModelCompositeModels); err != nil {
			log.Errorf("invalid --asset-model-composite-models: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseAssetModelDescription) > 0 {
		input.AssetModelDescription = aws.String(_iotsitewiseAssetModelDescription)
	}
	if len(_iotsitewiseAssetModelExternalId) > 0 {
		input.AssetModelExternalId = aws.String(_iotsitewiseAssetModelExternalId)
	}
	if len(_iotsitewiseAssetModelHierarchies) > 0 {
		if err := assignInputField(input, "AssetModelHierarchies", _iotsitewiseAssetModelHierarchies); err != nil {
			log.Errorf("invalid --asset-model-hierarchies: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseAssetModelId) > 0 {
		input.AssetModelId = aws.String(_iotsitewiseAssetModelId)
	}
	if len(_iotsitewiseAssetModelProperties) > 0 {
		if err := assignInputField(input, "AssetModelProperties", _iotsitewiseAssetModelProperties); err != nil {
			log.Errorf("invalid --asset-model-properties: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseAssetModelType) > 0 {
		if err := assignInputField(input, "AssetModelType", _iotsitewiseAssetModelType); err != nil {
			log.Errorf("invalid --asset-model-type: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseClientToken) > 0 {
		input.ClientToken = aws.String(_iotsitewiseClientToken)
	}
	if len(_iotsitewiseTags) > 0 {
		if err := assignInputField(input, "Tags", _iotsitewiseTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAssetModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a custom composite model from specified property and hierarchy
// definitions. There are two types of custom composite models, inline and
// component-model-based .
//
// Use component-model-based custom composite models to define standard, reusable
// components. A component-model-based custom composite model consists of a name, a
// description, and the ID of the component model it references. A
// component-model-based custom composite model has no properties of its own; its
// referenced component model provides its associated properties to any created
// assets. For more information, see [Custom composite models (Components)]in the IoT SiteWise User Guide.
//
// Use inline custom composite models to organize the properties of an asset
// model. The properties of inline custom composite models are local to the asset
// model where they are included and can't be used to create multiple assets.
//
// To create a component-model-based model, specify the composedAssetModelId of an
// existing asset model with assetModelType of COMPONENT_MODEL .
//
// To create an inline model, specify the assetModelCompositeModelProperties and
// don't include an composedAssetModelId .
//
// [Custom composite models (Components)]: https://docs.aws.amazon.com/iot-sitewise/latest/userguide/custom-composite-models.html
func iotsitewise_CreateAssetModelCompositeModel(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.CreateAssetModelCompositeModelInput{
		// AssetModelCompositeModelName: *string, // Required
		// AssetModelCompositeModelType: *string, // Required
		// AssetModelId: *string, // Required
	}

	if len(_iotsitewiseAssetModelCompositeModelName) > 0 {
		input.AssetModelCompositeModelName = aws.String(_iotsitewiseAssetModelCompositeModelName)
	}
	if len(_iotsitewiseAssetModelCompositeModelType) > 0 {
		input.AssetModelCompositeModelType = aws.String(_iotsitewiseAssetModelCompositeModelType)
	}
	if len(_iotsitewiseAssetModelId) > 0 {
		input.AssetModelId = aws.String(_iotsitewiseAssetModelId)
	}
	if len(_iotsitewiseAssetModelCompositeModelDescription) > 0 {
		input.AssetModelCompositeModelDescription = aws.String(_iotsitewiseAssetModelCompositeModelDescription)
	}
	if len(_iotsitewiseAssetModelCompositeModelExternalId) > 0 {
		input.AssetModelCompositeModelExternalId = aws.String(_iotsitewiseAssetModelCompositeModelExternalId)
	}
	if len(_iotsitewiseAssetModelCompositeModelId) > 0 {
		input.AssetModelCompositeModelId = aws.String(_iotsitewiseAssetModelCompositeModelId)
	}
	if len(_iotsitewiseAssetModelCompositeModelProperties) > 0 {
		if err := assignInputField(input, "AssetModelCompositeModelProperties", _iotsitewiseAssetModelCompositeModelProperties); err != nil {
			log.Errorf("invalid --asset-model-composite-model-properties: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseClientToken) > 0 {
		input.ClientToken = aws.String(_iotsitewiseClientToken)
	}
	if len(_iotsitewiseComposedAssetModelId) > 0 {
		input.ComposedAssetModelId = aws.String(_iotsitewiseComposedAssetModelId)
	}
	if len(_iotsitewiseIfMatch) > 0 {
		input.IfMatch = aws.String(_iotsitewiseIfMatch)
	}
	if len(_iotsitewiseIfNoneMatch) > 0 {
		input.IfNoneMatch = aws.String(_iotsitewiseIfNoneMatch)
	}
	if len(_iotsitewiseMatchForVersionType) > 0 {
		if err := assignInputField(input, "MatchForVersionType", _iotsitewiseMatchForVersionType); err != nil {
			log.Errorf("invalid --match-for-version-type: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseParentAssetModelCompositeModelId) > 0 {
		input.ParentAssetModelCompositeModelId = aws.String(_iotsitewiseParentAssetModelCompositeModelId)
	}

	if resp, err := client.CreateAssetModelCompositeModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Defines a job to ingest data to IoT SiteWise from Amazon S3. For more
// information, see [Create a bulk import job (CLI)]in the Amazon Simple Storage Service User Guide.
//
// Before you create a bulk import job, you must enable IoT SiteWise warm tier or
// IoT SiteWise cold tier. For more information about how to configure storage
// settings, see [PutStorageConfiguration].
//
// Bulk import is designed to store historical data to IoT SiteWise.
//
// - Newly ingested data in the hot tier triggers notifications and computations.
//
// - After data moves from the hot tier to the warm or cold tier based on
// retention settings, it does not trigger computations or notifications.
//
// - Data older than 7 days does not trigger computations or notifications.
//
// [Create a bulk import job (CLI)]: https://docs.aws.amazon.com/iot-sitewise/latest/userguide/CreateBulkImportJob.html
// [PutStorageConfiguration]: https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_PutStorageConfiguration.html
func iotsitewise_CreateBulkImportJob(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.CreateBulkImportJobInput{
		// ErrorReportLocation: *types.ErrorReportLocation, // Required
		// Files: []types.File, // Required
		// JobConfiguration: *types.JobConfiguration, // Required
		// JobName: *string, // Required
		// JobRoleArn: *string, // Required
	}

	if len(_iotsitewiseErrorReportLocation) > 0 {
		if err := assignInputField(input, "ErrorReportLocation", _iotsitewiseErrorReportLocation); err != nil {
			log.Errorf("invalid --error-report-location: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseFiles) > 0 {
		if err := assignInputField(input, "Files", _iotsitewiseFiles); err != nil {
			log.Errorf("invalid --files: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseJobConfiguration) > 0 {
		if err := assignInputField(input, "JobConfiguration", _iotsitewiseJobConfiguration); err != nil {
			log.Errorf("invalid --job-configuration: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseJobName) > 0 {
		input.JobName = aws.String(_iotsitewiseJobName)
	}
	if len(_iotsitewiseJobRoleArn) > 0 {
		input.JobRoleArn = aws.String(_iotsitewiseJobRoleArn)
	}
	if len(_iotsitewiseAdaptiveIngestion) > 0 {
		if err := assignInputField(input, "AdaptiveIngestion", _iotsitewiseAdaptiveIngestion); err != nil {
			log.Errorf("invalid --adaptive-ingestion: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseDeleteFilesAfterImport) > 0 {
		if err := assignInputField(input, "DeleteFilesAfterImport", _iotsitewiseDeleteFilesAfterImport); err != nil {
			log.Errorf("invalid --delete-files-after-import: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateBulkImportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a computation model with a configuration and data binding.
func iotsitewise_CreateComputationModel(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.CreateComputationModelInput{
		// ComputationModelConfiguration: *types.ComputationModelConfiguration, // Required
		// ComputationModelDataBinding: map[string]types.ComputationModelDataBindingValue, // Required
		// ComputationModelName: *string, // Required
	}

	if len(_iotsitewiseComputationModelConfiguration) > 0 {
		if err := assignInputField(input, "ComputationModelConfiguration", _iotsitewiseComputationModelConfiguration); err != nil {
			log.Errorf("invalid --computation-model-configuration: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseComputationModelDataBinding) > 0 {
		if err := assignInputField(input, "ComputationModelDataBinding", _iotsitewiseComputationModelDataBinding); err != nil {
			log.Errorf("invalid --computation-model-data-binding: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseComputationModelName) > 0 {
		input.ComputationModelName = aws.String(_iotsitewiseComputationModelName)
	}
	if len(_iotsitewiseClientToken) > 0 {
		input.ClientToken = aws.String(_iotsitewiseClientToken)
	}
	if len(_iotsitewiseComputationModelDescription) > 0 {
		input.ComputationModelDescription = aws.String(_iotsitewiseComputationModelDescription)
	}
	if len(_iotsitewiseTags) > 0 {
		if err := assignInputField(input, "Tags", _iotsitewiseTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateComputationModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a dashboard in an IoT SiteWise Monitor project.
func iotsitewise_CreateDashboard(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.CreateDashboardInput{
		// DashboardDefinition: *string, // Required
		// DashboardName: *string, // Required
		// ProjectId: *string, // Required
	}

	if len(_iotsitewiseDashboardDefinition) > 0 {
		input.DashboardDefinition = aws.String(_iotsitewiseDashboardDefinition)
	}
	if len(_iotsitewiseDashboardName) > 0 {
		input.DashboardName = aws.String(_iotsitewiseDashboardName)
	}
	if len(_iotsitewiseProjectId) > 0 {
		input.ProjectId = aws.String(_iotsitewiseProjectId)
	}
	if len(_iotsitewiseClientToken) > 0 {
		input.ClientToken = aws.String(_iotsitewiseClientToken)
	}
	if len(_iotsitewiseDashboardDescription) > 0 {
		input.DashboardDescription = aws.String(_iotsitewiseDashboardDescription)
	}
	if len(_iotsitewiseTags) > 0 {
		if err := assignInputField(input, "Tags", _iotsitewiseTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDashboard(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a dataset to connect an external datasource.
func iotsitewise_CreateDataset(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.CreateDatasetInput{
		// DatasetName: *string, // Required
		// DatasetSource: *types.DatasetSource, // Required
	}

	if len(_iotsitewiseDatasetName) > 0 {
		input.DatasetName = aws.String(_iotsitewiseDatasetName)
	}
	if len(_iotsitewiseDatasetSource) > 0 {
		if err := assignInputField(input, "DatasetSource", _iotsitewiseDatasetSource); err != nil {
			log.Errorf("invalid --dataset-source: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseClientToken) > 0 {
		input.ClientToken = aws.String(_iotsitewiseClientToken)
	}
	if len(_iotsitewiseDatasetDescription) > 0 {
		input.DatasetDescription = aws.String(_iotsitewiseDatasetDescription)
	}
	if len(_iotsitewiseDatasetId) > 0 {
		input.DatasetId = aws.String(_iotsitewiseDatasetId)
	}
	if len(_iotsitewiseTags) > 0 {
		if err := assignInputField(input, "Tags", _iotsitewiseTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDataset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a gateway, which is a virtual or edge device that delivers industrial
// data streams from local servers to IoT SiteWise. For more information, see [Ingesting data using a gateway]in
// the IoT SiteWise User Guide.
//
// [Ingesting data using a gateway]: https://docs.aws.amazon.com/iot-sitewise/latest/userguide/gateway-connector.html
func iotsitewise_CreateGateway(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.CreateGatewayInput{
		// GatewayName: *string, // Required
		// GatewayPlatform: *types.GatewayPlatform, // Required
	}

	if len(_iotsitewiseGatewayName) > 0 {
		input.GatewayName = aws.String(_iotsitewiseGatewayName)
	}
	if len(_iotsitewiseGatewayPlatform) > 0 {
		if err := assignInputField(input, "GatewayPlatform", _iotsitewiseGatewayPlatform); err != nil {
			log.Errorf("invalid --gateway-platform: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseGatewayVersion) > 0 {
		input.GatewayVersion = aws.String(_iotsitewiseGatewayVersion)
	}
	if len(_iotsitewiseTags) > 0 {
		if err := assignInputField(input, "Tags", _iotsitewiseTags); err != nil {
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

// Creates a portal, which can contain projects and dashboards. IoT SiteWise
// Monitor uses IAM Identity Center or IAM to authenticate portal users and manage
// user permissions.
//
// Before you can sign in to a new portal, you must add at least one identity to
// that portal. For more information, see [Adding or removing portal administrators]in the IoT SiteWise User Guide.
//
// [Adding or removing portal administrators]: https://docs.aws.amazon.com/iot-sitewise/latest/userguide/administer-portals.html#portal-change-admins
func iotsitewise_CreatePortal(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.CreatePortalInput{
		// PortalContactEmail: *string, // Required
		// PortalName: *string, // Required
		// RoleArn: *string, // Required
	}

	if len(_iotsitewisePortalContactEmail) > 0 {
		input.PortalContactEmail = aws.String(_iotsitewisePortalContactEmail)
	}
	if len(_iotsitewisePortalName) > 0 {
		input.PortalName = aws.String(_iotsitewisePortalName)
	}
	if len(_iotsitewiseRoleArn) > 0 {
		input.RoleArn = aws.String(_iotsitewiseRoleArn)
	}
	if len(_iotsitewiseAlarms) > 0 {
		if err := assignInputField(input, "Alarms", _iotsitewiseAlarms); err != nil {
			log.Errorf("invalid --alarms: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseClientToken) > 0 {
		input.ClientToken = aws.String(_iotsitewiseClientToken)
	}
	if len(_iotsitewiseNotificationSenderEmail) > 0 {
		input.NotificationSenderEmail = aws.String(_iotsitewiseNotificationSenderEmail)
	}
	if len(_iotsitewisePortalAuthMode) > 0 {
		if err := assignInputField(input, "PortalAuthMode", _iotsitewisePortalAuthMode); err != nil {
			log.Errorf("invalid --portal-auth-mode: %s", err.Error())
			return
		}
	}
	if len(_iotsitewisePortalDescription) > 0 {
		input.PortalDescription = aws.String(_iotsitewisePortalDescription)
	}
	if len(_iotsitewisePortalLogoImageFile) > 0 {
		if err := assignInputField(input, "PortalLogoImageFile", _iotsitewisePortalLogoImageFile); err != nil {
			log.Errorf("invalid --portal-logo-image-file: %s", err.Error())
			return
		}
	}
	if len(_iotsitewisePortalType) > 0 {
		if err := assignInputField(input, "PortalType", _iotsitewisePortalType); err != nil {
			log.Errorf("invalid --portal-type: %s", err.Error())
			return
		}
	}
	if len(_iotsitewisePortalTypeConfiguration) > 0 {
		if err := assignInputField(input, "PortalTypeConfiguration", _iotsitewisePortalTypeConfiguration); err != nil {
			log.Errorf("invalid --portal-type-configuration: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseTags) > 0 {
		if err := assignInputField(input, "Tags", _iotsitewiseTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreatePortal(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a project in the specified portal.
// Make sure that the project name and description don't contain confidential
// information.
func iotsitewise_CreateProject(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.CreateProjectInput{
		// PortalId: *string, // Required
		// ProjectName: *string, // Required
	}

	if len(_iotsitewisePortalId) > 0 {
		input.PortalId = aws.String(_iotsitewisePortalId)
	}
	if len(_iotsitewiseProjectName) > 0 {
		input.ProjectName = aws.String(_iotsitewiseProjectName)
	}
	if len(_iotsitewiseClientToken) > 0 {
		input.ClientToken = aws.String(_iotsitewiseClientToken)
	}
	if len(_iotsitewiseProjectDescription) > 0 {
		input.ProjectDescription = aws.String(_iotsitewiseProjectDescription)
	}
	if len(_iotsitewiseTags) > 0 {
		if err := assignInputField(input, "Tags", _iotsitewiseTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateProject(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an access policy that grants the specified identity access to the
// specified IoT SiteWise Monitor resource. You can use this operation to revoke
// access to an IoT SiteWise Monitor resource.
func iotsitewise_DeleteAccessPolicy(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.DeleteAccessPolicyInput{
		// AccessPolicyId: *string, // Required
	}

	if len(_iotsitewiseAccessPolicyId) > 0 {
		input.AccessPolicyId = aws.String(_iotsitewiseAccessPolicyId)
	}
	if len(_iotsitewiseClientToken) > 0 {
		input.ClientToken = aws.String(_iotsitewiseClientToken)
	}

	if resp, err := client.DeleteAccessPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an asset. This action can't be undone. For more information, see [Deleting assets and models] in
// the IoT SiteWise User Guide.
//
// You can't delete an asset that's associated to another asset. For more
// information, see [DisassociateAssets].
//
// [DisassociateAssets]: https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_DisassociateAssets.html
// [Deleting assets and models]: https://docs.aws.amazon.com/iot-sitewise/latest/userguide/delete-assets-and-models.html
func iotsitewise_DeleteAsset(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.DeleteAssetInput{
		// AssetId: *string, // Required
	}

	if len(_iotsitewiseAssetId) > 0 {
		input.AssetId = aws.String(_iotsitewiseAssetId)
	}
	if len(_iotsitewiseClientToken) > 0 {
		input.ClientToken = aws.String(_iotsitewiseClientToken)
	}

	if resp, err := client.DeleteAsset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an asset model. This action can't be undone. You must delete all assets
// created from an asset model before you can delete the model. Also, you can't
// delete an asset model if a parent asset model exists that contains a property
// formula expression that depends on the asset model that you want to delete. For
// more information, see [Deleting assets and models]in the IoT SiteWise User Guide.
//
// [Deleting assets and models]: https://docs.aws.amazon.com/iot-sitewise/latest/userguide/delete-assets-and-models.html
func iotsitewise_DeleteAssetModel(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.DeleteAssetModelInput{
		// AssetModelId: *string, // Required
	}

	if len(_iotsitewiseAssetModelId) > 0 {
		input.AssetModelId = aws.String(_iotsitewiseAssetModelId)
	}
	if len(_iotsitewiseClientToken) > 0 {
		input.ClientToken = aws.String(_iotsitewiseClientToken)
	}
	if len(_iotsitewiseIfMatch) > 0 {
		input.IfMatch = aws.String(_iotsitewiseIfMatch)
	}
	if len(_iotsitewiseIfNoneMatch) > 0 {
		input.IfNoneMatch = aws.String(_iotsitewiseIfNoneMatch)
	}
	if len(_iotsitewiseMatchForVersionType) > 0 {
		if err := assignInputField(input, "MatchForVersionType", _iotsitewiseMatchForVersionType); err != nil {
			log.Errorf("invalid --match-for-version-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteAssetModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a composite model. This action can't be undone. You must delete all
// assets created from a composite model before you can delete the model. Also, you
// can't delete a composite model if a parent asset model exists that contains a
// property formula expression that depends on the asset model that you want to
// delete. For more information, see [Deleting assets and models]in the IoT SiteWise User Guide.
//
// [Deleting assets and models]: https://docs.aws.amazon.com/iot-sitewise/latest/userguide/delete-assets-and-models.html
func iotsitewise_DeleteAssetModelCompositeModel(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.DeleteAssetModelCompositeModelInput{
		// AssetModelCompositeModelId: *string, // Required
		// AssetModelId: *string, // Required
	}

	if len(_iotsitewiseAssetModelCompositeModelId) > 0 {
		input.AssetModelCompositeModelId = aws.String(_iotsitewiseAssetModelCompositeModelId)
	}
	if len(_iotsitewiseAssetModelId) > 0 {
		input.AssetModelId = aws.String(_iotsitewiseAssetModelId)
	}
	if len(_iotsitewiseClientToken) > 0 {
		input.ClientToken = aws.String(_iotsitewiseClientToken)
	}
	if len(_iotsitewiseIfMatch) > 0 {
		input.IfMatch = aws.String(_iotsitewiseIfMatch)
	}
	if len(_iotsitewiseIfNoneMatch) > 0 {
		input.IfNoneMatch = aws.String(_iotsitewiseIfNoneMatch)
	}
	if len(_iotsitewiseMatchForVersionType) > 0 {
		if err := assignInputField(input, "MatchForVersionType", _iotsitewiseMatchForVersionType); err != nil {
			log.Errorf("invalid --match-for-version-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteAssetModelCompositeModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an interface relationship between an asset model and an interface asset
// model.
func iotsitewise_DeleteAssetModelInterfaceRelationship(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.DeleteAssetModelInterfaceRelationshipInput{
		// AssetModelId: *string, // Required
		// InterfaceAssetModelId: *string, // Required
	}

	if len(_iotsitewiseAssetModelId) > 0 {
		input.AssetModelId = aws.String(_iotsitewiseAssetModelId)
	}
	if len(_iotsitewiseInterfaceAssetModelId) > 0 {
		input.InterfaceAssetModelId = aws.String(_iotsitewiseInterfaceAssetModelId)
	}
	if len(_iotsitewiseClientToken) > 0 {
		input.ClientToken = aws.String(_iotsitewiseClientToken)
	}

	if resp, err := client.DeleteAssetModelInterfaceRelationship(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a computation model. This action can't be undone.
func iotsitewise_DeleteComputationModel(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.DeleteComputationModelInput{
		// ComputationModelId: *string, // Required
	}

	if len(_iotsitewiseComputationModelId) > 0 {
		input.ComputationModelId = aws.String(_iotsitewiseComputationModelId)
	}
	if len(_iotsitewiseClientToken) > 0 {
		input.ClientToken = aws.String(_iotsitewiseClientToken)
	}

	if resp, err := client.DeleteComputationModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a dashboard from IoT SiteWise Monitor.
func iotsitewise_DeleteDashboard(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.DeleteDashboardInput{
		// DashboardId: *string, // Required
	}

	if len(_iotsitewiseDashboardId) > 0 {
		input.DashboardId = aws.String(_iotsitewiseDashboardId)
	}
	if len(_iotsitewiseClientToken) > 0 {
		input.ClientToken = aws.String(_iotsitewiseClientToken)
	}

	if resp, err := client.DeleteDashboard(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a dataset. This cannot be undone.
func iotsitewise_DeleteDataset(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.DeleteDatasetInput{
		// DatasetId: *string, // Required
	}

	if len(_iotsitewiseDatasetId) > 0 {
		input.DatasetId = aws.String(_iotsitewiseDatasetId)
	}
	if len(_iotsitewiseClientToken) > 0 {
		input.ClientToken = aws.String(_iotsitewiseClientToken)
	}

	if resp, err := client.DeleteDataset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a gateway from IoT SiteWise. When you delete a gateway, some of the
// gateway's files remain in your gateway's file system.
func iotsitewise_DeleteGateway(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.DeleteGatewayInput{
		// GatewayId: *string, // Required
	}

	if len(_iotsitewiseGatewayId) > 0 {
		input.GatewayId = aws.String(_iotsitewiseGatewayId)
	}

	if resp, err := client.DeleteGateway(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a portal from IoT SiteWise Monitor.
func iotsitewise_DeletePortal(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.DeletePortalInput{
		// PortalId: *string, // Required
	}

	if len(_iotsitewisePortalId) > 0 {
		input.PortalId = aws.String(_iotsitewisePortalId)
	}
	if len(_iotsitewiseClientToken) > 0 {
		input.ClientToken = aws.String(_iotsitewiseClientToken)
	}

	if resp, err := client.DeletePortal(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a project from IoT SiteWise Monitor.
func iotsitewise_DeleteProject(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.DeleteProjectInput{
		// ProjectId: *string, // Required
	}

	if len(_iotsitewiseProjectId) > 0 {
		input.ProjectId = aws.String(_iotsitewiseProjectId)
	}
	if len(_iotsitewiseClientToken) > 0 {
		input.ClientToken = aws.String(_iotsitewiseClientToken)
	}

	if resp, err := client.DeleteProject(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a time series (data stream). If you delete a time series that's
// associated with an asset property, the asset property still exists, but the time
// series will no longer be associated with this asset property.
//
// To identify a time series, do one of the following:
//
// - If the time series isn't associated with an asset property, specify the
// alias of the time series.
//
// - If the time series is associated with an asset property, specify one of the
// following:
//
// - The alias of the time series.
//
// - The assetId and propertyId that identifies the asset property.
func iotsitewise_DeleteTimeSeries(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.DeleteTimeSeriesInput{}

	if len(_iotsitewiseAlias) > 0 {
		input.Alias = aws.String(_iotsitewiseAlias)
	}
	if len(_iotsitewiseAssetId) > 0 {
		input.AssetId = aws.String(_iotsitewiseAssetId)
	}
	if len(_iotsitewiseClientToken) > 0 {
		input.ClientToken = aws.String(_iotsitewiseClientToken)
	}
	if len(_iotsitewisePropertyId) > 0 {
		input.PropertyId = aws.String(_iotsitewisePropertyId)
	}

	if resp, err := client.DeleteTimeSeries(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes an access policy, which specifies an identity's access to an IoT
// SiteWise Monitor portal or project.
func iotsitewise_DescribeAccessPolicy(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.DescribeAccessPolicyInput{
		// AccessPolicyId: *string, // Required
	}

	if len(_iotsitewiseAccessPolicyId) > 0 {
		input.AccessPolicyId = aws.String(_iotsitewiseAccessPolicyId)
	}

	if resp, err := client.DescribeAccessPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about an action.
func iotsitewise_DescribeAction(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.DescribeActionInput{
		// ActionId: *string, // Required
	}

	if len(_iotsitewiseActionId) > 0 {
		input.ActionId = aws.String(_iotsitewiseActionId)
	}

	if resp, err := client.DescribeAction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about an asset.
func iotsitewise_DescribeAsset(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.DescribeAssetInput{
		// AssetId: *string, // Required
	}

	if len(_iotsitewiseAssetId) > 0 {
		input.AssetId = aws.String(_iotsitewiseAssetId)
	}
	if len(_iotsitewiseExcludeProperties) > 0 {
		if err := assignInputField(input, "ExcludeProperties", _iotsitewiseExcludeProperties); err != nil {
			log.Errorf("invalid --exclude-properties: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeAsset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about an asset composite model (also known as an asset
// component). An AssetCompositeModel is an instance of an AssetModelCompositeModel
// . If you want to see information about the model this is based on, call [DescribeAssetModelCompositeModel].
//
// [DescribeAssetModelCompositeModel]: https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_DescribeAssetModelCompositeModel.html
func iotsitewise_DescribeAssetCompositeModel(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.DescribeAssetCompositeModelInput{
		// AssetCompositeModelId: *string, // Required
		// AssetId: *string, // Required
	}

	if len(_iotsitewiseAssetCompositeModelId) > 0 {
		input.AssetCompositeModelId = aws.String(_iotsitewiseAssetCompositeModelId)
	}
	if len(_iotsitewiseAssetId) > 0 {
		input.AssetId = aws.String(_iotsitewiseAssetId)
	}

	if resp, err := client.DescribeAssetCompositeModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about an asset model. This includes details about the
// asset model's properties, hierarchies, composite models, and any interface
// relationships if the asset model implements interfaces.
func iotsitewise_DescribeAssetModel(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.DescribeAssetModelInput{
		// AssetModelId: *string, // Required
	}

	if len(_iotsitewiseAssetModelId) > 0 {
		input.AssetModelId = aws.String(_iotsitewiseAssetModelId)
	}
	if len(_iotsitewiseAssetModelVersion) > 0 {
		input.AssetModelVersion = aws.String(_iotsitewiseAssetModelVersion)
	}
	if len(_iotsitewiseExcludeProperties) > 0 {
		if err := assignInputField(input, "ExcludeProperties", _iotsitewiseExcludeProperties); err != nil {
			log.Errorf("invalid --exclude-properties: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeAssetModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about an asset model composite model (also known as an
// asset model component). For more information, see [Custom composite models (Components)]in the IoT SiteWise User
// Guide.
//
// [Custom composite models (Components)]: https://docs.aws.amazon.com/iot-sitewise/latest/userguide/custom-composite-models.html
func iotsitewise_DescribeAssetModelCompositeModel(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.DescribeAssetModelCompositeModelInput{
		// AssetModelCompositeModelId: *string, // Required
		// AssetModelId: *string, // Required
	}

	if len(_iotsitewiseAssetModelCompositeModelId) > 0 {
		input.AssetModelCompositeModelId = aws.String(_iotsitewiseAssetModelCompositeModelId)
	}
	if len(_iotsitewiseAssetModelId) > 0 {
		input.AssetModelId = aws.String(_iotsitewiseAssetModelId)
	}
	if len(_iotsitewiseAssetModelVersion) > 0 {
		input.AssetModelVersion = aws.String(_iotsitewiseAssetModelVersion)
	}

	if resp, err := client.DescribeAssetModelCompositeModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about an interface relationship between an asset model
// and an interface asset model.
func iotsitewise_DescribeAssetModelInterfaceRelationship(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.DescribeAssetModelInterfaceRelationshipInput{
		// AssetModelId: *string, // Required
		// InterfaceAssetModelId: *string, // Required
	}

	if len(_iotsitewiseAssetModelId) > 0 {
		input.AssetModelId = aws.String(_iotsitewiseAssetModelId)
	}
	if len(_iotsitewiseInterfaceAssetModelId) > 0 {
		input.InterfaceAssetModelId = aws.String(_iotsitewiseInterfaceAssetModelId)
	}

	if resp, err := client.DescribeAssetModelInterfaceRelationship(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about an asset property.
// When you call this operation for an attribute property, this response includes
// the default attribute value that you define in the asset model. If you update
// the default value in the model, this operation's response includes the new
// default value.
//
// This operation doesn't return the value of the asset property. To get the value
// of an asset property, use [GetAssetPropertyValue].
//
// [GetAssetPropertyValue]: https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_GetAssetPropertyValue.html
func iotsitewise_DescribeAssetProperty(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.DescribeAssetPropertyInput{
		// AssetId: *string, // Required
		// PropertyId: *string, // Required
	}

	if len(_iotsitewiseAssetId) > 0 {
		input.AssetId = aws.String(_iotsitewiseAssetId)
	}
	if len(_iotsitewisePropertyId) > 0 {
		input.PropertyId = aws.String(_iotsitewisePropertyId)
	}

	if resp, err := client.DescribeAssetProperty(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a bulk import job request. For more information,
// see [Describe a bulk import job (CLI)]in the Amazon Simple Storage Service User Guide.
//
// [Describe a bulk import job (CLI)]: https://docs.aws.amazon.com/iot-sitewise/latest/userguide/DescribeBulkImportJob.html
func iotsitewise_DescribeBulkImportJob(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.DescribeBulkImportJobInput{
		// JobId: *string, // Required
	}

	if len(_iotsitewiseJobId) > 0 {
		input.JobId = aws.String(_iotsitewiseJobId)
	}

	if resp, err := client.DescribeBulkImportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a computation model.
func iotsitewise_DescribeComputationModel(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.DescribeComputationModelInput{
		// ComputationModelId: *string, // Required
	}

	if len(_iotsitewiseComputationModelId) > 0 {
		input.ComputationModelId = aws.String(_iotsitewiseComputationModelId)
	}
	if len(_iotsitewiseComputationModelVersion) > 0 {
		input.ComputationModelVersion = aws.String(_iotsitewiseComputationModelVersion)
	}

	if resp, err := client.DescribeComputationModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the execution summary of a computation model.
func iotsitewise_DescribeComputationModelExecutionSummary(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.DescribeComputationModelExecutionSummaryInput{
		// ComputationModelId: *string, // Required
	}

	if len(_iotsitewiseComputationModelId) > 0 {
		input.ComputationModelId = aws.String(_iotsitewiseComputationModelId)
	}
	if len(_iotsitewiseResolveToResourceId) > 0 {
		input.ResolveToResourceId = aws.String(_iotsitewiseResolveToResourceId)
	}
	if len(_iotsitewiseResolveToResourceType) > 0 {
		if err := assignInputField(input, "ResolveToResourceType", _iotsitewiseResolveToResourceType); err != nil {
			log.Errorf("invalid --resolve-to-resource-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeComputationModelExecutionSummary(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a dashboard.
func iotsitewise_DescribeDashboard(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.DescribeDashboardInput{
		// DashboardId: *string, // Required
	}

	if len(_iotsitewiseDashboardId) > 0 {
		input.DashboardId = aws.String(_iotsitewiseDashboardId)
	}

	if resp, err := client.DescribeDashboard(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a dataset.
func iotsitewise_DescribeDataset(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.DescribeDatasetInput{
		// DatasetId: *string, // Required
	}

	if len(_iotsitewiseDatasetId) > 0 {
		input.DatasetId = aws.String(_iotsitewiseDatasetId)
	}

	if resp, err := client.DescribeDataset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the default encryption configuration for the Amazon
// Web Services account in the default or specified Region. For more information,
// see [Key management]in the IoT SiteWise User Guide.
//
// [Key management]: https://docs.aws.amazon.com/iot-sitewise/latest/userguide/key-management.html
func iotsitewise_DescribeDefaultEncryptionConfiguration(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.DescribeDefaultEncryptionConfigurationInput{}

	if resp, err := client.DescribeDefaultEncryptionConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the execution.
func iotsitewise_DescribeExecution(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.DescribeExecutionInput{
		// ExecutionId: *string, // Required
	}

	if len(_iotsitewiseExecutionId) > 0 {
		input.ExecutionId = aws.String(_iotsitewiseExecutionId)
	}

	if resp, err := client.DescribeExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a gateway.
func iotsitewise_DescribeGateway(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.DescribeGatewayInput{
		// GatewayId: *string, // Required
	}

	if len(_iotsitewiseGatewayId) > 0 {
		input.GatewayId = aws.String(_iotsitewiseGatewayId)
	}

	if resp, err := client.DescribeGateway(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Each gateway capability defines data sources for a gateway. This is the
// namespace of the gateway capability.
//
// . The namespace follows the format service:capability:version , where:
//
// - service - The service providing the capability, or iotsitewise .
//
// - capability - The specific capability type. Options include: opcuacollector
// for the OPC UA data source collector, or publisher for data publisher
// capability.
//
// - version - The version number of the capability. Option include 2 for Classic
// streams, V2 gateways, and 3 for MQTT-enabled, V3 gateways.
//
// After updating a capability configuration, the sync status becomes OUT_OF_SYNC
// until the gateway processes the configuration.Use
// DescribeGatewayCapabilityConfiguration to check the sync status and verify the
// configuration was applied.
//
// A gateway can have multiple capability configurations with different namespaces.
func iotsitewise_DescribeGatewayCapabilityConfiguration(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.DescribeGatewayCapabilityConfigurationInput{
		// CapabilityNamespace: *string, // Required
		// GatewayId: *string, // Required
	}

	if len(_iotsitewiseCapabilityNamespace) > 0 {
		input.CapabilityNamespace = aws.String(_iotsitewiseCapabilityNamespace)
	}
	if len(_iotsitewiseGatewayId) > 0 {
		input.GatewayId = aws.String(_iotsitewiseGatewayId)
	}

	if resp, err := client.DescribeGatewayCapabilityConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the current IoT SiteWise logging options.
func iotsitewise_DescribeLoggingOptions(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.DescribeLoggingOptionsInput{}

	if resp, err := client.DescribeLoggingOptions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a portal.
func iotsitewise_DescribePortal(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.DescribePortalInput{
		// PortalId: *string, // Required
	}

	if len(_iotsitewisePortalId) > 0 {
		input.PortalId = aws.String(_iotsitewisePortalId)
	}

	if resp, err := client.DescribePortal(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a project.
func iotsitewise_DescribeProject(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.DescribeProjectInput{
		// ProjectId: *string, // Required
	}

	if len(_iotsitewiseProjectId) > 0 {
		input.ProjectId = aws.String(_iotsitewiseProjectId)
	}

	if resp, err := client.DescribeProject(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the storage configuration for IoT SiteWise.
func iotsitewise_DescribeStorageConfiguration(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.DescribeStorageConfigurationInput{}

	if resp, err := client.DescribeStorageConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a time series (data stream).
// To identify a time series, do one of the following:
//
// - If the time series isn't associated with an asset property, specify the
// alias of the time series.
//
// - If the time series is associated with an asset property, specify one of the
// following:
//
// - The alias of the time series.
//
// - The assetId and propertyId that identifies the asset property.
func iotsitewise_DescribeTimeSeries(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.DescribeTimeSeriesInput{}

	if len(_iotsitewiseAlias) > 0 {
		input.Alias = aws.String(_iotsitewiseAlias)
	}
	if len(_iotsitewiseAssetId) > 0 {
		input.AssetId = aws.String(_iotsitewiseAssetId)
	}
	if len(_iotsitewisePropertyId) > 0 {
		input.PropertyId = aws.String(_iotsitewisePropertyId)
	}

	if resp, err := client.DescribeTimeSeries(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates a child asset from the given parent asset through a hierarchy
// defined in the parent asset's model.
func iotsitewise_DisassociateAssets(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.DisassociateAssetsInput{
		// AssetId: *string, // Required
		// ChildAssetId: *string, // Required
		// HierarchyId: *string, // Required
	}

	if len(_iotsitewiseAssetId) > 0 {
		input.AssetId = aws.String(_iotsitewiseAssetId)
	}
	if len(_iotsitewiseChildAssetId) > 0 {
		input.ChildAssetId = aws.String(_iotsitewiseChildAssetId)
	}
	if len(_iotsitewiseHierarchyId) > 0 {
		input.HierarchyId = aws.String(_iotsitewiseHierarchyId)
	}
	if len(_iotsitewiseClientToken) > 0 {
		input.ClientToken = aws.String(_iotsitewiseClientToken)
	}

	if resp, err := client.DisassociateAssets(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates a time series (data stream) from an asset property.
func iotsitewise_DisassociateTimeSeriesFromAssetProperty(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.DisassociateTimeSeriesFromAssetPropertyInput{
		// Alias: *string, // Required
		// AssetId: *string, // Required
		// PropertyId: *string, // Required
	}

	if len(_iotsitewiseAlias) > 0 {
		input.Alias = aws.String(_iotsitewiseAlias)
	}
	if len(_iotsitewiseAssetId) > 0 {
		input.AssetId = aws.String(_iotsitewiseAssetId)
	}
	if len(_iotsitewisePropertyId) > 0 {
		input.PropertyId = aws.String(_iotsitewisePropertyId)
	}
	if len(_iotsitewiseClientToken) > 0 {
		input.ClientToken = aws.String(_iotsitewiseClientToken)
	}

	if resp, err := client.DisassociateTimeSeriesFromAssetProperty(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Executes an action on a target resource.
func iotsitewise_ExecuteAction(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.ExecuteActionInput{
		// ActionDefinitionId: *string, // Required
		// ActionPayload: *types.ActionPayload, // Required
		// TargetResource: *types.TargetResource, // Required
	}

	if len(_iotsitewiseActionDefinitionId) > 0 {
		input.ActionDefinitionId = aws.String(_iotsitewiseActionDefinitionId)
	}
	if len(_iotsitewiseActionPayload) > 0 {
		if err := assignInputField(input, "ActionPayload", _iotsitewiseActionPayload); err != nil {
			log.Errorf("invalid --action-payload: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseTargetResource) > 0 {
		if err := assignInputField(input, "TargetResource", _iotsitewiseTargetResource); err != nil {
			log.Errorf("invalid --target-resource: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseClientToken) > 0 {
		input.ClientToken = aws.String(_iotsitewiseClientToken)
	}
	if len(_iotsitewiseResolveTo) > 0 {
		if err := assignInputField(input, "ResolveTo", _iotsitewiseResolveTo); err != nil {
			log.Errorf("invalid --resolve-to: %s", err.Error())
			return
		}
	}

	if resp, err := client.ExecuteAction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Run SQL queries to retrieve metadata and time-series data from asset models,
// assets, measurements, metrics, transforms, and aggregates.
func iotsitewise_ExecuteQuery(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.ExecuteQueryInput{
		// QueryStatement: *string, // Required
	}

	if len(_iotsitewiseQueryStatement) > 0 {
		input.QueryStatement = aws.String(_iotsitewiseQueryStatement)
	}
	if len(_iotsitewiseClientToken) > 0 {
		input.ClientToken = aws.String(_iotsitewiseClientToken)
	}
	if len(_iotsitewiseMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotsitewiseMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseNextToken) > 0 {
		input.NextToken = aws.String(_iotsitewiseNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ExecuteQuery(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotsitewise.ExecuteQueryOutput
	p := iotsitewise.NewExecuteQueryPaginator(client, input)
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

// Gets aggregated values for an asset property. For more information, see [Querying aggregates] in the
// IoT SiteWise User Guide.
//
// To identify an asset property, you must specify one of the following:
//
// - The assetId and propertyId of an asset property.
//
// - A propertyAlias , which is a data stream alias (for example,
// /company/windfarm/3/turbine/7/temperature ). To define an asset property's
// alias, see [UpdateAssetProperty].
//
// [UpdateAssetProperty]: https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_UpdateAssetProperty.html
// [Querying aggregates]: https://docs.aws.amazon.com/iot-sitewise/latest/userguide/query-industrial-data.html#aggregates
func iotsitewise_GetAssetPropertyAggregates(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.GetAssetPropertyAggregatesInput{
		// AggregateTypes: []types.AggregateType, // Required
		// EndDate: *time.Time, // Required
		// Resolution: *string, // Required
		// StartDate: *time.Time, // Required
	}

	if len(_iotsitewiseAggregateTypes) > 0 {
		if err := assignInputField(input, "AggregateTypes", _iotsitewiseAggregateTypes); err != nil {
			log.Errorf("invalid --aggregate-types: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseEndDate) > 0 {
		if err := assignInputField(input, "EndDate", _iotsitewiseEndDate); err != nil {
			log.Errorf("invalid --end-date: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseResolution) > 0 {
		input.Resolution = aws.String(_iotsitewiseResolution)
	}
	if len(_iotsitewiseStartDate) > 0 {
		if err := assignInputField(input, "StartDate", _iotsitewiseStartDate); err != nil {
			log.Errorf("invalid --start-date: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseAssetId) > 0 {
		input.AssetId = aws.String(_iotsitewiseAssetId)
	}
	if len(_iotsitewiseMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotsitewiseMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseNextToken) > 0 {
		input.NextToken = aws.String(_iotsitewiseNextToken)
	}
	if len(_iotsitewisePropertyAlias) > 0 {
		input.PropertyAlias = aws.String(_iotsitewisePropertyAlias)
	}
	if len(_iotsitewisePropertyId) > 0 {
		input.PropertyId = aws.String(_iotsitewisePropertyId)
	}
	if len(_iotsitewiseQualities) > 0 {
		if err := assignInputField(input, "Qualities", _iotsitewiseQualities); err != nil {
			log.Errorf("invalid --qualities: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseTimeOrdering) > 0 {
		if err := assignInputField(input, "TimeOrdering", _iotsitewiseTimeOrdering); err != nil {
			log.Errorf("invalid --time-ordering: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.GetAssetPropertyAggregates(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotsitewise.GetAssetPropertyAggregatesOutput
	p := iotsitewise.NewGetAssetPropertyAggregatesPaginator(client, input)
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

// Gets an asset property's current value. For more information, see [Querying current values] in the IoT
// SiteWise User Guide.
//
// To identify an asset property, you must specify one of the following:
//
// - The assetId and propertyId of an asset property.
//
// - A propertyAlias , which is a data stream alias (for example,
// /company/windfarm/3/turbine/7/temperature ). To define an asset property's
// alias, see [UpdateAssetProperty].
//
// [UpdateAssetProperty]: https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_UpdateAssetProperty.html
// [Querying current values]: https://docs.aws.amazon.com/iot-sitewise/latest/userguide/query-industrial-data.html#current-values
func iotsitewise_GetAssetPropertyValue(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.GetAssetPropertyValueInput{}

	if len(_iotsitewiseAssetId) > 0 {
		input.AssetId = aws.String(_iotsitewiseAssetId)
	}
	if len(_iotsitewisePropertyAlias) > 0 {
		input.PropertyAlias = aws.String(_iotsitewisePropertyAlias)
	}
	if len(_iotsitewisePropertyId) > 0 {
		input.PropertyId = aws.String(_iotsitewisePropertyId)
	}

	if resp, err := client.GetAssetPropertyValue(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the history of an asset property's values. For more information, see [Querying historical values] in
// the IoT SiteWise User Guide.
//
// To identify an asset property, you must specify one of the following:
//
// - The assetId and propertyId of an asset property.
//
// - A propertyAlias , which is a data stream alias (for example,
// /company/windfarm/3/turbine/7/temperature ). To define an asset property's
// alias, see [UpdateAssetProperty].
//
// [UpdateAssetProperty]: https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_UpdateAssetProperty.html
// [Querying historical values]: https://docs.aws.amazon.com/iot-sitewise/latest/userguide/query-industrial-data.html#historical-values
func iotsitewise_GetAssetPropertyValueHistory(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.GetAssetPropertyValueHistoryInput{}

	if len(_iotsitewiseAssetId) > 0 {
		input.AssetId = aws.String(_iotsitewiseAssetId)
	}
	if len(_iotsitewiseEndDate) > 0 {
		if err := assignInputField(input, "EndDate", _iotsitewiseEndDate); err != nil {
			log.Errorf("invalid --end-date: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotsitewiseMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseNextToken) > 0 {
		input.NextToken = aws.String(_iotsitewiseNextToken)
	}
	if len(_iotsitewisePropertyAlias) > 0 {
		input.PropertyAlias = aws.String(_iotsitewisePropertyAlias)
	}
	if len(_iotsitewisePropertyId) > 0 {
		input.PropertyId = aws.String(_iotsitewisePropertyId)
	}
	if len(_iotsitewiseQualities) > 0 {
		if err := assignInputField(input, "Qualities", _iotsitewiseQualities); err != nil {
			log.Errorf("invalid --qualities: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseStartDate) > 0 {
		if err := assignInputField(input, "StartDate", _iotsitewiseStartDate); err != nil {
			log.Errorf("invalid --start-date: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseTimeOrdering) > 0 {
		if err := assignInputField(input, "TimeOrdering", _iotsitewiseTimeOrdering); err != nil {
			log.Errorf("invalid --time-ordering: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.GetAssetPropertyValueHistory(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotsitewise.GetAssetPropertyValueHistoryOutput
	p := iotsitewise.NewGetAssetPropertyValueHistoryPaginator(client, input)
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

// Get interpolated values for an asset property for a specified time interval,
// during a period of time. If your time series is missing data points during the
// specified time interval, you can use interpolation to estimate the missing data.
//
// For example, you can use this operation to return the interpolated temperature
// values for a wind turbine every 24 hours over a duration of 7 days.
//
// To identify an asset property, you must specify one of the following:
//
// - The assetId and propertyId of an asset property.
//
// - A propertyAlias , which is a data stream alias (for example,
// /company/windfarm/3/turbine/7/temperature ). To define an asset property's
// alias, see [UpdateAssetProperty].
//
// [UpdateAssetProperty]: https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_UpdateAssetProperty.html
func iotsitewise_GetInterpolatedAssetPropertyValues(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.GetInterpolatedAssetPropertyValuesInput{
		// EndTimeInSeconds: *int64, // Required
		// IntervalInSeconds: *int64, // Required
		// Quality: types.Quality, // Required
		// StartTimeInSeconds: *int64, // Required
		// Type: *string, // Required
	}

	if len(_iotsitewiseEndTimeInSeconds) > 0 {
		if err := assignInputField(input, "EndTimeInSeconds", _iotsitewiseEndTimeInSeconds); err != nil {
			log.Errorf("invalid --end-time-in-seconds: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseIntervalInSeconds) > 0 {
		if err := assignInputField(input, "IntervalInSeconds", _iotsitewiseIntervalInSeconds); err != nil {
			log.Errorf("invalid --interval-in-seconds: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseQuality) > 0 {
		if err := assignInputField(input, "Quality", _iotsitewiseQuality); err != nil {
			log.Errorf("invalid --quality: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseStartTimeInSeconds) > 0 {
		if err := assignInputField(input, "StartTimeInSeconds", _iotsitewiseStartTimeInSeconds); err != nil {
			log.Errorf("invalid --start-time-in-seconds: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseType) > 0 {
		input.Type = aws.String(_iotsitewiseType)
	}
	if len(_iotsitewiseAssetId) > 0 {
		input.AssetId = aws.String(_iotsitewiseAssetId)
	}
	if len(_iotsitewiseEndTimeOffsetInNanos) > 0 {
		if err := assignInputField(input, "EndTimeOffsetInNanos", _iotsitewiseEndTimeOffsetInNanos); err != nil {
			log.Errorf("invalid --end-time-offset-in-nanos: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseIntervalWindowInSeconds) > 0 {
		if err := assignInputField(input, "IntervalWindowInSeconds", _iotsitewiseIntervalWindowInSeconds); err != nil {
			log.Errorf("invalid --interval-window-in-seconds: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotsitewiseMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseNextToken) > 0 {
		input.NextToken = aws.String(_iotsitewiseNextToken)
	}
	if len(_iotsitewisePropertyAlias) > 0 {
		input.PropertyAlias = aws.String(_iotsitewisePropertyAlias)
	}
	if len(_iotsitewisePropertyId) > 0 {
		input.PropertyId = aws.String(_iotsitewisePropertyId)
	}
	if len(_iotsitewiseStartTimeOffsetInNanos) > 0 {
		if err := assignInputField(input, "StartTimeOffsetInNanos", _iotsitewiseStartTimeOffsetInNanos); err != nil {
			log.Errorf("invalid --start-time-offset-in-nanos: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.GetInterpolatedAssetPropertyValues(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotsitewise.GetInterpolatedAssetPropertyValuesOutput
	p := iotsitewise.NewGetInterpolatedAssetPropertyValuesPaginator(client, input)
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

// Invokes SiteWise Assistant to start or continue a conversation.
func iotsitewise_InvokeAssistant(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.InvokeAssistantInput{
		// Message: *string, // Required
	}

	if len(_iotsitewiseMessage) > 0 {
		input.Message = aws.String(_iotsitewiseMessage)
	}
	if len(_iotsitewiseConversationId) > 0 {
		input.ConversationId = aws.String(_iotsitewiseConversationId)
	}
	if len(_iotsitewiseEnableTrace) > 0 {
		if err := assignInputField(input, "EnableTrace", _iotsitewiseEnableTrace); err != nil {
			log.Errorf("invalid --enable-trace: %s", err.Error())
			return
		}
	}

	if resp, err := client.InvokeAssistant(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a paginated list of access policies for an identity (an IAM Identity
// Center user, an IAM Identity Center group, or an IAM user) or an IoT SiteWise
// Monitor resource (a portal or project).
func iotsitewise_ListAccessPolicies(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.ListAccessPoliciesInput{}

	if len(_iotsitewiseIamArn) > 0 {
		input.IamArn = aws.String(_iotsitewiseIamArn)
	}
	if len(_iotsitewiseIdentityId) > 0 {
		input.IdentityId = aws.String(_iotsitewiseIdentityId)
	}
	if len(_iotsitewiseIdentityType) > 0 {
		if err := assignInputField(input, "IdentityType", _iotsitewiseIdentityType); err != nil {
			log.Errorf("invalid --identity-type: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotsitewiseMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseNextToken) > 0 {
		input.NextToken = aws.String(_iotsitewiseNextToken)
	}
	if len(_iotsitewiseResourceId) > 0 {
		input.ResourceId = aws.String(_iotsitewiseResourceId)
	}
	if len(_iotsitewiseResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _iotsitewiseResourceType); err != nil {
			log.Errorf("invalid --resource-type: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListAccessPolicies(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotsitewise.ListAccessPoliciesOutput
	p := iotsitewise.NewListAccessPoliciesPaginator(client, input)
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

// Retrieves a paginated list of actions for a specific target resource.
func iotsitewise_ListActions(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.ListActionsInput{
		// TargetResourceId: *string, // Required
		// TargetResourceType: types.TargetResourceType, // Required
	}

	if len(_iotsitewiseTargetResourceId) > 0 {
		input.TargetResourceId = aws.String(_iotsitewiseTargetResourceId)
	}
	if len(_iotsitewiseTargetResourceType) > 0 {
		if err := assignInputField(input, "TargetResourceType", _iotsitewiseTargetResourceType); err != nil {
			log.Errorf("invalid --target-resource-type: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotsitewiseMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseNextToken) > 0 {
		input.NextToken = aws.String(_iotsitewiseNextToken)
	}
	if len(_iotsitewiseResolveToResourceId) > 0 {
		input.ResolveToResourceId = aws.String(_iotsitewiseResolveToResourceId)
	}
	if len(_iotsitewiseResolveToResourceType) > 0 {
		if err := assignInputField(input, "ResolveToResourceType", _iotsitewiseResolveToResourceType); err != nil {
			log.Errorf("invalid --resolve-to-resource-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListActions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a paginated list of composite models associated with the asset model
func iotsitewise_ListAssetModelCompositeModels(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.ListAssetModelCompositeModelsInput{
		// AssetModelId: *string, // Required
	}

	if len(_iotsitewiseAssetModelId) > 0 {
		input.AssetModelId = aws.String(_iotsitewiseAssetModelId)
	}
	if len(_iotsitewiseAssetModelVersion) > 0 {
		input.AssetModelVersion = aws.String(_iotsitewiseAssetModelVersion)
	}
	if len(_iotsitewiseMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotsitewiseMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseNextToken) > 0 {
		input.NextToken = aws.String(_iotsitewiseNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAssetModelCompositeModels(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotsitewise.ListAssetModelCompositeModelsOutput
	p := iotsitewise.NewListAssetModelCompositeModelsPaginator(client, input)
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

// Retrieves a paginated list of properties associated with an asset model. If you
// update properties associated with the model before you finish listing all the
// properties, you need to start all over again.
func iotsitewise_ListAssetModelProperties(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.ListAssetModelPropertiesInput{
		// AssetModelId: *string, // Required
	}

	if len(_iotsitewiseAssetModelId) > 0 {
		input.AssetModelId = aws.String(_iotsitewiseAssetModelId)
	}
	if len(_iotsitewiseAssetModelVersion) > 0 {
		input.AssetModelVersion = aws.String(_iotsitewiseAssetModelVersion)
	}
	if len(_iotsitewiseFilter) > 0 {
		if err := assignInputField(input, "Filter", _iotsitewiseFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotsitewiseMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseNextToken) > 0 {
		input.NextToken = aws.String(_iotsitewiseNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAssetModelProperties(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotsitewise.ListAssetModelPropertiesOutput
	p := iotsitewise.NewListAssetModelPropertiesPaginator(client, input)
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

// Retrieves a paginated list of summaries of all asset models.
func iotsitewise_ListAssetModels(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.ListAssetModelsInput{}

	if len(_iotsitewiseAssetModelTypes) > 0 {
		if err := assignInputField(input, "AssetModelTypes", _iotsitewiseAssetModelTypes); err != nil {
			log.Errorf("invalid --asset-model-types: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseAssetModelVersion) > 0 {
		input.AssetModelVersion = aws.String(_iotsitewiseAssetModelVersion)
	}
	if len(_iotsitewiseMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotsitewiseMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseNextToken) > 0 {
		input.NextToken = aws.String(_iotsitewiseNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAssetModels(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotsitewise.ListAssetModelsOutput
	p := iotsitewise.NewListAssetModelsPaginator(client, input)
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

// Retrieves a paginated list of properties associated with an asset. If you
// update properties associated with the model before you finish listing all the
// properties, you need to start all over again.
func iotsitewise_ListAssetProperties(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.ListAssetPropertiesInput{
		// AssetId: *string, // Required
	}

	if len(_iotsitewiseAssetId) > 0 {
		input.AssetId = aws.String(_iotsitewiseAssetId)
	}
	if len(_iotsitewiseFilter) > 0 {
		if err := assignInputField(input, "Filter", _iotsitewiseFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotsitewiseMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseNextToken) > 0 {
		input.NextToken = aws.String(_iotsitewiseNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAssetProperties(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotsitewise.ListAssetPropertiesOutput
	p := iotsitewise.NewListAssetPropertiesPaginator(client, input)
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

// Retrieves a paginated list of asset relationships for an asset. You can use
// this operation to identify an asset's root asset and all associated assets
// between that asset and its root.
func iotsitewise_ListAssetRelationships(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.ListAssetRelationshipsInput{
		// AssetId: *string, // Required
		// TraversalType: types.TraversalType, // Required
	}

	if len(_iotsitewiseAssetId) > 0 {
		input.AssetId = aws.String(_iotsitewiseAssetId)
	}
	if len(_iotsitewiseTraversalType) > 0 {
		if err := assignInputField(input, "TraversalType", _iotsitewiseTraversalType); err != nil {
			log.Errorf("invalid --traversal-type: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotsitewiseMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseNextToken) > 0 {
		input.NextToken = aws.String(_iotsitewiseNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAssetRelationships(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotsitewise.ListAssetRelationshipsOutput
	p := iotsitewise.NewListAssetRelationshipsPaginator(client, input)
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

// Retrieves a paginated list of asset summaries.
// You can use this operation to do the following:
//
// - List assets based on a specific asset model.
//
// - List top-level assets.
//
// You can't use this operation to list all assets. To retrieve summaries for all
// of your assets, use [ListAssetModels]to get all of your asset model IDs. Then, use ListAssets to
// get all assets for each asset model.
//
// [ListAssetModels]: https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_ListAssetModels.html
func iotsitewise_ListAssets(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.ListAssetsInput{}

	if len(_iotsitewiseAssetModelId) > 0 {
		input.AssetModelId = aws.String(_iotsitewiseAssetModelId)
	}
	if len(_iotsitewiseFilter) > 0 {
		if err := assignInputField(input, "Filter", _iotsitewiseFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotsitewiseMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseNextToken) > 0 {
		input.NextToken = aws.String(_iotsitewiseNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAssets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotsitewise.ListAssetsOutput
	p := iotsitewise.NewListAssetsPaginator(client, input)
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

// Retrieves a paginated list of associated assets.
// You can use this operation to do the following:
//
// - CHILD - List all child assets associated to the asset.
//
// - PARENT - List the asset's parent asset.
func iotsitewise_ListAssociatedAssets(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.ListAssociatedAssetsInput{
		// AssetId: *string, // Required
	}

	if len(_iotsitewiseAssetId) > 0 {
		input.AssetId = aws.String(_iotsitewiseAssetId)
	}
	if len(_iotsitewiseHierarchyId) > 0 {
		input.HierarchyId = aws.String(_iotsitewiseHierarchyId)
	}
	if len(_iotsitewiseMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotsitewiseMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseNextToken) > 0 {
		input.NextToken = aws.String(_iotsitewiseNextToken)
	}
	if len(_iotsitewiseTraversalDirection) > 0 {
		if err := assignInputField(input, "TraversalDirection", _iotsitewiseTraversalDirection); err != nil {
			log.Errorf("invalid --traversal-direction: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListAssociatedAssets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotsitewise.ListAssociatedAssetsOutput
	p := iotsitewise.NewListAssociatedAssetsPaginator(client, input)
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

// Retrieves a paginated list of bulk import job requests. For more information,
// see [List bulk import jobs (CLI)]in the IoT SiteWise User Guide.
//
// [List bulk import jobs (CLI)]: https://docs.aws.amazon.com/iot-sitewise/latest/userguide/ListBulkImportJobs.html
func iotsitewise_ListBulkImportJobs(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.ListBulkImportJobsInput{}

	if len(_iotsitewiseFilter) > 0 {
		if err := assignInputField(input, "Filter", _iotsitewiseFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotsitewiseMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseNextToken) > 0 {
		input.NextToken = aws.String(_iotsitewiseNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListBulkImportJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotsitewise.ListBulkImportJobsOutput
	p := iotsitewise.NewListBulkImportJobsPaginator(client, input)
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

// Retrieves a paginated list of composition relationships for an asset model of
// type COMPONENT_MODEL .
func iotsitewise_ListCompositionRelationships(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.ListCompositionRelationshipsInput{
		// AssetModelId: *string, // Required
	}

	if len(_iotsitewiseAssetModelId) > 0 {
		input.AssetModelId = aws.String(_iotsitewiseAssetModelId)
	}
	if len(_iotsitewiseMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotsitewiseMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseNextToken) > 0 {
		input.NextToken = aws.String(_iotsitewiseNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCompositionRelationships(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotsitewise.ListCompositionRelationshipsOutput
	p := iotsitewise.NewListCompositionRelationshipsPaginator(client, input)
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

// Lists all data binding usages for computation models. This allows to identify
// where specific data bindings are being utilized across the computation models.
// This track dependencies between data sources and computation models.
func iotsitewise_ListComputationModelDataBindingUsages(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.ListComputationModelDataBindingUsagesInput{
		// DataBindingValueFilter: *types.DataBindingValueFilter, // Required
	}

	if len(_iotsitewiseDataBindingValueFilter) > 0 {
		if err := assignInputField(input, "DataBindingValueFilter", _iotsitewiseDataBindingValueFilter); err != nil {
			log.Errorf("invalid --data-binding-value-filter: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotsitewiseMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseNextToken) > 0 {
		input.NextToken = aws.String(_iotsitewiseNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListComputationModelDataBindingUsages(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotsitewise.ListComputationModelDataBindingUsagesOutput
	p := iotsitewise.NewListComputationModelDataBindingUsagesPaginator(client, input)
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

// Lists all distinct resources that are resolved from the executed actions of the
// computation model.
func iotsitewise_ListComputationModelResolveToResources(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.ListComputationModelResolveToResourcesInput{
		// ComputationModelId: *string, // Required
	}

	if len(_iotsitewiseComputationModelId) > 0 {
		input.ComputationModelId = aws.String(_iotsitewiseComputationModelId)
	}
	if len(_iotsitewiseMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotsitewiseMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseNextToken) > 0 {
		input.NextToken = aws.String(_iotsitewiseNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListComputationModelResolveToResources(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotsitewise.ListComputationModelResolveToResourcesOutput
	p := iotsitewise.NewListComputationModelResolveToResourcesPaginator(client, input)
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

// Retrieves a paginated list of summaries of all computation models.
func iotsitewise_ListComputationModels(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.ListComputationModelsInput{}

	if len(_iotsitewiseComputationModelType) > 0 {
		if err := assignInputField(input, "ComputationModelType", _iotsitewiseComputationModelType); err != nil {
			log.Errorf("invalid --computation-model-type: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotsitewiseMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseNextToken) > 0 {
		input.NextToken = aws.String(_iotsitewiseNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListComputationModels(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotsitewise.ListComputationModelsOutput
	p := iotsitewise.NewListComputationModelsPaginator(client, input)
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

// Retrieves a paginated list of dashboards for an IoT SiteWise Monitor project.
func iotsitewise_ListDashboards(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.ListDashboardsInput{
		// ProjectId: *string, // Required
	}

	if len(_iotsitewiseProjectId) > 0 {
		input.ProjectId = aws.String(_iotsitewiseProjectId)
	}
	if len(_iotsitewiseMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotsitewiseMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseNextToken) > 0 {
		input.NextToken = aws.String(_iotsitewiseNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDashboards(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotsitewise.ListDashboardsOutput
	p := iotsitewise.NewListDashboardsPaginator(client, input)
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

// Retrieves a paginated list of datasets for a specific target resource.
func iotsitewise_ListDatasets(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.ListDatasetsInput{
		// SourceType: types.DatasetSourceType, // Required
	}

	if len(_iotsitewiseSourceType) > 0 {
		if err := assignInputField(input, "SourceType", _iotsitewiseSourceType); err != nil {
			log.Errorf("invalid --source-type: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotsitewiseMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseNextToken) > 0 {
		input.NextToken = aws.String(_iotsitewiseNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDatasets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotsitewise.ListDatasetsOutput
	p := iotsitewise.NewListDatasetsPaginator(client, input)
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

// Retrieves a paginated list of summaries of all executions.
func iotsitewise_ListExecutions(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.ListExecutionsInput{
		// TargetResourceId: *string, // Required
		// TargetResourceType: types.TargetResourceType, // Required
	}

	if len(_iotsitewiseTargetResourceId) > 0 {
		input.TargetResourceId = aws.String(_iotsitewiseTargetResourceId)
	}
	if len(_iotsitewiseTargetResourceType) > 0 {
		if err := assignInputField(input, "TargetResourceType", _iotsitewiseTargetResourceType); err != nil {
			log.Errorf("invalid --target-resource-type: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseActionType) > 0 {
		input.ActionType = aws.String(_iotsitewiseActionType)
	}
	if len(_iotsitewiseMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotsitewiseMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseNextToken) > 0 {
		input.NextToken = aws.String(_iotsitewiseNextToken)
	}
	if len(_iotsitewiseResolveToResourceId) > 0 {
		input.ResolveToResourceId = aws.String(_iotsitewiseResolveToResourceId)
	}
	if len(_iotsitewiseResolveToResourceType) > 0 {
		if err := assignInputField(input, "ResolveToResourceType", _iotsitewiseResolveToResourceType); err != nil {
			log.Errorf("invalid --resolve-to-resource-type: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListExecutions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotsitewise.ListExecutionsOutput
	p := iotsitewise.NewListExecutionsPaginator(client, input)
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

// Retrieves a paginated list of gateways.
func iotsitewise_ListGateways(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.ListGatewaysInput{}

	if len(_iotsitewiseMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotsitewiseMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseNextToken) > 0 {
		input.NextToken = aws.String(_iotsitewiseNextToken)
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

	var results []*iotsitewise.ListGatewaysOutput
	p := iotsitewise.NewListGatewaysPaginator(client, input)
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

// Retrieves a paginated list of asset models that have a specific interface asset
// model applied to them.
func iotsitewise_ListInterfaceRelationships(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.ListInterfaceRelationshipsInput{
		// InterfaceAssetModelId: *string, // Required
	}

	if len(_iotsitewiseInterfaceAssetModelId) > 0 {
		input.InterfaceAssetModelId = aws.String(_iotsitewiseInterfaceAssetModelId)
	}
	if len(_iotsitewiseMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotsitewiseMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseNextToken) > 0 {
		input.NextToken = aws.String(_iotsitewiseNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListInterfaceRelationships(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotsitewise.ListInterfaceRelationshipsOutput
	p := iotsitewise.NewListInterfaceRelationshipsPaginator(client, input)
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

// Retrieves a paginated list of IoT SiteWise Monitor portals.
func iotsitewise_ListPortals(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.ListPortalsInput{}

	if len(_iotsitewiseMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotsitewiseMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseNextToken) > 0 {
		input.NextToken = aws.String(_iotsitewiseNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPortals(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotsitewise.ListPortalsOutput
	p := iotsitewise.NewListPortalsPaginator(client, input)
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

// Retrieves a paginated list of assets associated with an IoT SiteWise Monitor
// project.
func iotsitewise_ListProjectAssets(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.ListProjectAssetsInput{
		// ProjectId: *string, // Required
	}

	if len(_iotsitewiseProjectId) > 0 {
		input.ProjectId = aws.String(_iotsitewiseProjectId)
	}
	if len(_iotsitewiseMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotsitewiseMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseNextToken) > 0 {
		input.NextToken = aws.String(_iotsitewiseNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListProjectAssets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotsitewise.ListProjectAssetsOutput
	p := iotsitewise.NewListProjectAssetsPaginator(client, input)
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

// Retrieves a paginated list of projects for an IoT SiteWise Monitor portal.
func iotsitewise_ListProjects(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.ListProjectsInput{
		// PortalId: *string, // Required
	}

	if len(_iotsitewisePortalId) > 0 {
		input.PortalId = aws.String(_iotsitewisePortalId)
	}
	if len(_iotsitewiseMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotsitewiseMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseNextToken) > 0 {
		input.NextToken = aws.String(_iotsitewiseNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListProjects(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotsitewise.ListProjectsOutput
	p := iotsitewise.NewListProjectsPaginator(client, input)
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

// Retrieves the list of tags for an IoT SiteWise resource.
func iotsitewise_ListTagsForResource(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_iotsitewiseResourceArn) > 0 {
		input.ResourceArn = aws.String(_iotsitewiseResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a paginated list of time series (data streams).
func iotsitewise_ListTimeSeries(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.ListTimeSeriesInput{}

	if len(_iotsitewiseAliasPrefix) > 0 {
		input.AliasPrefix = aws.String(_iotsitewiseAliasPrefix)
	}
	if len(_iotsitewiseAssetId) > 0 {
		input.AssetId = aws.String(_iotsitewiseAssetId)
	}
	if len(_iotsitewiseMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotsitewiseMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseNextToken) > 0 {
		input.NextToken = aws.String(_iotsitewiseNextToken)
	}
	if len(_iotsitewiseTimeSeriesType) > 0 {
		if err := assignInputField(input, "TimeSeriesType", _iotsitewiseTimeSeriesType); err != nil {
			log.Errorf("invalid --time-series-type: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListTimeSeries(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotsitewise.ListTimeSeriesOutput
	p := iotsitewise.NewListTimeSeriesPaginator(client, input)
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

// Creates or updates an interface relationship between an asset model and an
// interface asset model. This operation applies an interface to an asset model.
func iotsitewise_PutAssetModelInterfaceRelationship(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.PutAssetModelInterfaceRelationshipInput{
		// AssetModelId: *string, // Required
		// InterfaceAssetModelId: *string, // Required
		// PropertyMappingConfiguration: *types.PropertyMappingConfiguration, // Required
	}

	if len(_iotsitewiseAssetModelId) > 0 {
		input.AssetModelId = aws.String(_iotsitewiseAssetModelId)
	}
	if len(_iotsitewiseInterfaceAssetModelId) > 0 {
		input.InterfaceAssetModelId = aws.String(_iotsitewiseInterfaceAssetModelId)
	}
	if len(_iotsitewisePropertyMappingConfiguration) > 0 {
		if err := assignInputField(input, "PropertyMappingConfiguration", _iotsitewisePropertyMappingConfiguration); err != nil {
			log.Errorf("invalid --property-mapping-configuration: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseClientToken) > 0 {
		input.ClientToken = aws.String(_iotsitewiseClientToken)
	}

	if resp, err := client.PutAssetModelInterfaceRelationship(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets the default encryption configuration for the Amazon Web Services account.
// For more information, see [Key management]in the IoT SiteWise User Guide.
//
// [Key management]: https://docs.aws.amazon.com/iot-sitewise/latest/userguide/key-management.html
func iotsitewise_PutDefaultEncryptionConfiguration(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.PutDefaultEncryptionConfigurationInput{
		// EncryptionType: types.EncryptionType, // Required
	}

	if len(_iotsitewiseEncryptionType) > 0 {
		if err := assignInputField(input, "EncryptionType", _iotsitewiseEncryptionType); err != nil {
			log.Errorf("invalid --encryption-type: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_iotsitewiseKmsKeyId)
	}

	if resp, err := client.PutDefaultEncryptionConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets logging options for IoT SiteWise.
func iotsitewise_PutLoggingOptions(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.PutLoggingOptionsInput{
		// LoggingOptions: *types.LoggingOptions, // Required
	}

	if len(_iotsitewiseLoggingOptions) > 0 {
		if err := assignInputField(input, "LoggingOptions", _iotsitewiseLoggingOptions); err != nil {
			log.Errorf("invalid --logging-options: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutLoggingOptions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Configures storage settings for IoT SiteWise.
func iotsitewise_PutStorageConfiguration(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.PutStorageConfigurationInput{
		// StorageType: types.StorageType, // Required
	}

	if len(_iotsitewiseStorageType) > 0 {
		if err := assignInputField(input, "StorageType", _iotsitewiseStorageType); err != nil {
			log.Errorf("invalid --storage-type: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseDisallowIngestNullNaN) > 0 {
		if err := assignInputField(input, "DisallowIngestNullNaN", _iotsitewiseDisallowIngestNullNaN); err != nil {
			log.Errorf("invalid --disallow-ingest-null-na-n: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseDisassociatedDataStorage) > 0 {
		if err := assignInputField(input, "DisassociatedDataStorage", _iotsitewiseDisassociatedDataStorage); err != nil {
			log.Errorf("invalid --disassociated-data-storage: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseMultiLayerStorage) > 0 {
		if err := assignInputField(input, "MultiLayerStorage", _iotsitewiseMultiLayerStorage); err != nil {
			log.Errorf("invalid --multi-layer-storage: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseRetentionPeriod) > 0 {
		if err := assignInputField(input, "RetentionPeriod", _iotsitewiseRetentionPeriod); err != nil {
			log.Errorf("invalid --retention-period: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseWarmTier) > 0 {
		if err := assignInputField(input, "WarmTier", _iotsitewiseWarmTier); err != nil {
			log.Errorf("invalid --warm-tier: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseWarmTierRetentionPeriod) > 0 {
		if err := assignInputField(input, "WarmTierRetentionPeriod", _iotsitewiseWarmTierRetentionPeriod); err != nil {
			log.Errorf("invalid --warm-tier-retention-period: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutStorageConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds tags to an IoT SiteWise resource. If a tag already exists for the
// resource, this operation updates the tag's value.
func iotsitewise_TagResource(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_iotsitewiseResourceArn) > 0 {
		input.ResourceArn = aws.String(_iotsitewiseResourceArn)
	}
	if len(_iotsitewiseTags) > 0 {
		if err := assignInputField(input, "Tags", _iotsitewiseTags); err != nil {
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

// Removes a tag from an IoT SiteWise resource.
func iotsitewise_UntagResource(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_iotsitewiseResourceArn) > 0 {
		input.ResourceArn = aws.String(_iotsitewiseResourceArn)
	}
	if len(_iotsitewiseTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _iotsitewiseTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing access policy that specifies an identity's access to an IoT
// SiteWise Monitor portal or project resource.
func iotsitewise_UpdateAccessPolicy(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.UpdateAccessPolicyInput{
		// AccessPolicyId: *string, // Required
		// AccessPolicyIdentity: *types.Identity, // Required
		// AccessPolicyPermission: types.Permission, // Required
		// AccessPolicyResource: *types.Resource, // Required
	}

	if len(_iotsitewiseAccessPolicyId) > 0 {
		input.AccessPolicyId = aws.String(_iotsitewiseAccessPolicyId)
	}
	if len(_iotsitewiseAccessPolicyIdentity) > 0 {
		if err := assignInputField(input, "AccessPolicyIdentity", _iotsitewiseAccessPolicyIdentity); err != nil {
			log.Errorf("invalid --access-policy-identity: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseAccessPolicyPermission) > 0 {
		if err := assignInputField(input, "AccessPolicyPermission", _iotsitewiseAccessPolicyPermission); err != nil {
			log.Errorf("invalid --access-policy-permission: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseAccessPolicyResource) > 0 {
		if err := assignInputField(input, "AccessPolicyResource", _iotsitewiseAccessPolicyResource); err != nil {
			log.Errorf("invalid --access-policy-resource: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseClientToken) > 0 {
		input.ClientToken = aws.String(_iotsitewiseClientToken)
	}

	if resp, err := client.UpdateAccessPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an asset's name. For more information, see [Updating assets and models] in the IoT SiteWise User
// Guide.
//
// [Updating assets and models]: https://docs.aws.amazon.com/iot-sitewise/latest/userguide/update-assets-and-models.html
func iotsitewise_UpdateAsset(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.UpdateAssetInput{
		// AssetId: *string, // Required
		// AssetName: *string, // Required
	}

	if len(_iotsitewiseAssetId) > 0 {
		input.AssetId = aws.String(_iotsitewiseAssetId)
	}
	if len(_iotsitewiseAssetName) > 0 {
		input.AssetName = aws.String(_iotsitewiseAssetName)
	}
	if len(_iotsitewiseAssetDescription) > 0 {
		input.AssetDescription = aws.String(_iotsitewiseAssetDescription)
	}
	if len(_iotsitewiseAssetExternalId) > 0 {
		input.AssetExternalId = aws.String(_iotsitewiseAssetExternalId)
	}
	if len(_iotsitewiseClientToken) > 0 {
		input.ClientToken = aws.String(_iotsitewiseClientToken)
	}

	if resp, err := client.UpdateAsset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an asset model and all of the assets that were created from the model.
// Each asset created from the model inherits the updated asset model's property
// and hierarchy definitions. For more information, see [Updating assets and models]in the IoT SiteWise User
// Guide.
//
// If you remove a property from an asset model, IoT SiteWise deletes all previous
// data for that property. You can’t change the type or data type of an existing
// property.
//
// To replace an existing asset model property with a new one with the same name ,
// do the following:
//
// - Submit an UpdateAssetModel request with the entire existing property removed.
//
// - Submit a second UpdateAssetModel request that includes the new property. The
// new asset property will have the same name as the previous one and IoT
// SiteWise will generate a new unique id .
//
// [Updating assets and models]: https://docs.aws.amazon.com/iot-sitewise/latest/userguide/update-assets-and-models.html
func iotsitewise_UpdateAssetModel(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.UpdateAssetModelInput{
		// AssetModelId: *string, // Required
		// AssetModelName: *string, // Required
	}

	if len(_iotsitewiseAssetModelId) > 0 {
		input.AssetModelId = aws.String(_iotsitewiseAssetModelId)
	}
	if len(_iotsitewiseAssetModelName) > 0 {
		input.AssetModelName = aws.String(_iotsitewiseAssetModelName)
	}
	if len(_iotsitewiseAssetModelCompositeModels) > 0 {
		if err := assignInputField(input, "AssetModelCompositeModels", _iotsitewiseAssetModelCompositeModels); err != nil {
			log.Errorf("invalid --asset-model-composite-models: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseAssetModelDescription) > 0 {
		input.AssetModelDescription = aws.String(_iotsitewiseAssetModelDescription)
	}
	if len(_iotsitewiseAssetModelExternalId) > 0 {
		input.AssetModelExternalId = aws.String(_iotsitewiseAssetModelExternalId)
	}
	if len(_iotsitewiseAssetModelHierarchies) > 0 {
		if err := assignInputField(input, "AssetModelHierarchies", _iotsitewiseAssetModelHierarchies); err != nil {
			log.Errorf("invalid --asset-model-hierarchies: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseAssetModelProperties) > 0 {
		if err := assignInputField(input, "AssetModelProperties", _iotsitewiseAssetModelProperties); err != nil {
			log.Errorf("invalid --asset-model-properties: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseClientToken) > 0 {
		input.ClientToken = aws.String(_iotsitewiseClientToken)
	}
	if len(_iotsitewiseIfMatch) > 0 {
		input.IfMatch = aws.String(_iotsitewiseIfMatch)
	}
	if len(_iotsitewiseIfNoneMatch) > 0 {
		input.IfNoneMatch = aws.String(_iotsitewiseIfNoneMatch)
	}
	if len(_iotsitewiseMatchForVersionType) > 0 {
		if err := assignInputField(input, "MatchForVersionType", _iotsitewiseMatchForVersionType); err != nil {
			log.Errorf("invalid --match-for-version-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateAssetModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a composite model and all of the assets that were created from the
// model. Each asset created from the model inherits the updated asset model's
// property and hierarchy definitions. For more information, see [Updating assets and models]in the IoT
// SiteWise User Guide.
//
// If you remove a property from a composite asset model, IoT SiteWise deletes all
// previous data for that property. You can’t change the type or data type of an
// existing property.
//
// To replace an existing composite asset model property with a new one with the
// same name , do the following:
//
// - Submit an UpdateAssetModelCompositeModel request with the entire existing
// property removed.
//
// - Submit a second UpdateAssetModelCompositeModel request that includes the new
// property. The new asset property will have the same name as the previous one
// and IoT SiteWise will generate a new unique id .
//
// [Updating assets and models]: https://docs.aws.amazon.com/iot-sitewise/latest/userguide/update-assets-and-models.html
func iotsitewise_UpdateAssetModelCompositeModel(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.UpdateAssetModelCompositeModelInput{
		// AssetModelCompositeModelId: *string, // Required
		// AssetModelCompositeModelName: *string, // Required
		// AssetModelId: *string, // Required
	}

	if len(_iotsitewiseAssetModelCompositeModelId) > 0 {
		input.AssetModelCompositeModelId = aws.String(_iotsitewiseAssetModelCompositeModelId)
	}
	if len(_iotsitewiseAssetModelCompositeModelName) > 0 {
		input.AssetModelCompositeModelName = aws.String(_iotsitewiseAssetModelCompositeModelName)
	}
	if len(_iotsitewiseAssetModelId) > 0 {
		input.AssetModelId = aws.String(_iotsitewiseAssetModelId)
	}
	if len(_iotsitewiseAssetModelCompositeModelDescription) > 0 {
		input.AssetModelCompositeModelDescription = aws.String(_iotsitewiseAssetModelCompositeModelDescription)
	}
	if len(_iotsitewiseAssetModelCompositeModelExternalId) > 0 {
		input.AssetModelCompositeModelExternalId = aws.String(_iotsitewiseAssetModelCompositeModelExternalId)
	}
	if len(_iotsitewiseAssetModelCompositeModelProperties) > 0 {
		if err := assignInputField(input, "AssetModelCompositeModelProperties", _iotsitewiseAssetModelCompositeModelProperties); err != nil {
			log.Errorf("invalid --asset-model-composite-model-properties: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseClientToken) > 0 {
		input.ClientToken = aws.String(_iotsitewiseClientToken)
	}
	if len(_iotsitewiseIfMatch) > 0 {
		input.IfMatch = aws.String(_iotsitewiseIfMatch)
	}
	if len(_iotsitewiseIfNoneMatch) > 0 {
		input.IfNoneMatch = aws.String(_iotsitewiseIfNoneMatch)
	}
	if len(_iotsitewiseMatchForVersionType) > 0 {
		if err := assignInputField(input, "MatchForVersionType", _iotsitewiseMatchForVersionType); err != nil {
			log.Errorf("invalid --match-for-version-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateAssetModelCompositeModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an asset property's alias and notification state.
// This operation overwrites the property's existing alias and notification state.
// To keep your existing property's alias or notification state, you must include
// the existing values in the UpdateAssetProperty request. For more information,
// see [DescribeAssetProperty].
//
// [DescribeAssetProperty]: https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_DescribeAssetProperty.html
func iotsitewise_UpdateAssetProperty(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.UpdateAssetPropertyInput{
		// AssetId: *string, // Required
		// PropertyId: *string, // Required
	}

	if len(_iotsitewiseAssetId) > 0 {
		input.AssetId = aws.String(_iotsitewiseAssetId)
	}
	if len(_iotsitewisePropertyId) > 0 {
		input.PropertyId = aws.String(_iotsitewisePropertyId)
	}
	if len(_iotsitewiseClientToken) > 0 {
		input.ClientToken = aws.String(_iotsitewiseClientToken)
	}
	if len(_iotsitewisePropertyAlias) > 0 {
		input.PropertyAlias = aws.String(_iotsitewisePropertyAlias)
	}
	if len(_iotsitewisePropertyNotificationState) > 0 {
		if err := assignInputField(input, "PropertyNotificationState", _iotsitewisePropertyNotificationState); err != nil {
			log.Errorf("invalid --property-notification-state: %s", err.Error())
			return
		}
	}
	if len(_iotsitewisePropertyUnit) > 0 {
		input.PropertyUnit = aws.String(_iotsitewisePropertyUnit)
	}

	if resp, err := client.UpdateAssetProperty(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the computation model.
func iotsitewise_UpdateComputationModel(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.UpdateComputationModelInput{
		// ComputationModelConfiguration: *types.ComputationModelConfiguration, // Required
		// ComputationModelDataBinding: map[string]types.ComputationModelDataBindingValue, // Required
		// ComputationModelId: *string, // Required
		// ComputationModelName: *string, // Required
	}

	if len(_iotsitewiseComputationModelConfiguration) > 0 {
		if err := assignInputField(input, "ComputationModelConfiguration", _iotsitewiseComputationModelConfiguration); err != nil {
			log.Errorf("invalid --computation-model-configuration: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseComputationModelDataBinding) > 0 {
		if err := assignInputField(input, "ComputationModelDataBinding", _iotsitewiseComputationModelDataBinding); err != nil {
			log.Errorf("invalid --computation-model-data-binding: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseComputationModelId) > 0 {
		input.ComputationModelId = aws.String(_iotsitewiseComputationModelId)
	}
	if len(_iotsitewiseComputationModelName) > 0 {
		input.ComputationModelName = aws.String(_iotsitewiseComputationModelName)
	}
	if len(_iotsitewiseClientToken) > 0 {
		input.ClientToken = aws.String(_iotsitewiseClientToken)
	}
	if len(_iotsitewiseComputationModelDescription) > 0 {
		input.ComputationModelDescription = aws.String(_iotsitewiseComputationModelDescription)
	}

	if resp, err := client.UpdateComputationModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an IoT SiteWise Monitor dashboard.
func iotsitewise_UpdateDashboard(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.UpdateDashboardInput{
		// DashboardDefinition: *string, // Required
		// DashboardId: *string, // Required
		// DashboardName: *string, // Required
	}

	if len(_iotsitewiseDashboardDefinition) > 0 {
		input.DashboardDefinition = aws.String(_iotsitewiseDashboardDefinition)
	}
	if len(_iotsitewiseDashboardId) > 0 {
		input.DashboardId = aws.String(_iotsitewiseDashboardId)
	}
	if len(_iotsitewiseDashboardName) > 0 {
		input.DashboardName = aws.String(_iotsitewiseDashboardName)
	}
	if len(_iotsitewiseClientToken) > 0 {
		input.ClientToken = aws.String(_iotsitewiseClientToken)
	}
	if len(_iotsitewiseDashboardDescription) > 0 {
		input.DashboardDescription = aws.String(_iotsitewiseDashboardDescription)
	}

	if resp, err := client.UpdateDashboard(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a dataset.
func iotsitewise_UpdateDataset(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.UpdateDatasetInput{
		// DatasetId: *string, // Required
		// DatasetName: *string, // Required
		// DatasetSource: *types.DatasetSource, // Required
	}

	if len(_iotsitewiseDatasetId) > 0 {
		input.DatasetId = aws.String(_iotsitewiseDatasetId)
	}
	if len(_iotsitewiseDatasetName) > 0 {
		input.DatasetName = aws.String(_iotsitewiseDatasetName)
	}
	if len(_iotsitewiseDatasetSource) > 0 {
		if err := assignInputField(input, "DatasetSource", _iotsitewiseDatasetSource); err != nil {
			log.Errorf("invalid --dataset-source: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseClientToken) > 0 {
		input.ClientToken = aws.String(_iotsitewiseClientToken)
	}
	if len(_iotsitewiseDatasetDescription) > 0 {
		input.DatasetDescription = aws.String(_iotsitewiseDatasetDescription)
	}

	if resp, err := client.UpdateDataset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a gateway's name.
func iotsitewise_UpdateGateway(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.UpdateGatewayInput{
		// GatewayId: *string, // Required
		// GatewayName: *string, // Required
	}

	if len(_iotsitewiseGatewayId) > 0 {
		input.GatewayId = aws.String(_iotsitewiseGatewayId)
	}
	if len(_iotsitewiseGatewayName) > 0 {
		input.GatewayName = aws.String(_iotsitewiseGatewayName)
	}

	if resp, err := client.UpdateGateway(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a gateway capability configuration or defines a new capability
// configuration. Each gateway capability defines data sources for a gateway.
//
// Important workflow notes:
//
// Each gateway capability defines data sources for a gateway. This is the
// namespace of the gateway capability.
//
// . The namespace follows the format service:capability:version , where:
//
// - service - The service providing the capability, or iotsitewise .
//
// - capability - The specific capability type. Options include: opcuacollector
// for the OPC UA data source collector, or publisher for data publisher
// capability.
//
// - version - The version number of the capability. Option include 2 for Classic
// streams, V2 gateways, and 3 for MQTT-enabled, V3 gateways.
//
// After updating a capability configuration, the sync status becomes OUT_OF_SYNC
// until the gateway processes the configuration.Use
// DescribeGatewayCapabilityConfiguration to check the sync status and verify the
// configuration was applied.
//
// A gateway can have multiple capability configurations with different namespaces.
func iotsitewise_UpdateGatewayCapabilityConfiguration(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.UpdateGatewayCapabilityConfigurationInput{
		// CapabilityConfiguration: *string, // Required
		// CapabilityNamespace: *string, // Required
		// GatewayId: *string, // Required
	}

	if len(_iotsitewiseCapabilityConfiguration) > 0 {
		input.CapabilityConfiguration = aws.String(_iotsitewiseCapabilityConfiguration)
	}
	if len(_iotsitewiseCapabilityNamespace) > 0 {
		input.CapabilityNamespace = aws.String(_iotsitewiseCapabilityNamespace)
	}
	if len(_iotsitewiseGatewayId) > 0 {
		input.GatewayId = aws.String(_iotsitewiseGatewayId)
	}

	if resp, err := client.UpdateGatewayCapabilityConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an IoT SiteWise Monitor portal.
func iotsitewise_UpdatePortal(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.UpdatePortalInput{
		// PortalContactEmail: *string, // Required
		// PortalId: *string, // Required
		// PortalName: *string, // Required
		// RoleArn: *string, // Required
	}

	if len(_iotsitewisePortalContactEmail) > 0 {
		input.PortalContactEmail = aws.String(_iotsitewisePortalContactEmail)
	}
	if len(_iotsitewisePortalId) > 0 {
		input.PortalId = aws.String(_iotsitewisePortalId)
	}
	if len(_iotsitewisePortalName) > 0 {
		input.PortalName = aws.String(_iotsitewisePortalName)
	}
	if len(_iotsitewiseRoleArn) > 0 {
		input.RoleArn = aws.String(_iotsitewiseRoleArn)
	}
	if len(_iotsitewiseAlarms) > 0 {
		if err := assignInputField(input, "Alarms", _iotsitewiseAlarms); err != nil {
			log.Errorf("invalid --alarms: %s", err.Error())
			return
		}
	}
	if len(_iotsitewiseClientToken) > 0 {
		input.ClientToken = aws.String(_iotsitewiseClientToken)
	}
	if len(_iotsitewiseNotificationSenderEmail) > 0 {
		input.NotificationSenderEmail = aws.String(_iotsitewiseNotificationSenderEmail)
	}
	if len(_iotsitewisePortalDescription) > 0 {
		input.PortalDescription = aws.String(_iotsitewisePortalDescription)
	}
	if len(_iotsitewisePortalLogoImage) > 0 {
		if err := assignInputField(input, "PortalLogoImage", _iotsitewisePortalLogoImage); err != nil {
			log.Errorf("invalid --portal-logo-image: %s", err.Error())
			return
		}
	}
	if len(_iotsitewisePortalType) > 0 {
		if err := assignInputField(input, "PortalType", _iotsitewisePortalType); err != nil {
			log.Errorf("invalid --portal-type: %s", err.Error())
			return
		}
	}
	if len(_iotsitewisePortalTypeConfiguration) > 0 {
		if err := assignInputField(input, "PortalTypeConfiguration", _iotsitewisePortalTypeConfiguration); err != nil {
			log.Errorf("invalid --portal-type-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdatePortal(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an IoT SiteWise Monitor project.
func iotsitewise_UpdateProject(cfg aws.Config, client *iotsitewise.Client) {
	input := &iotsitewise.UpdateProjectInput{
		// ProjectId: *string, // Required
		// ProjectName: *string, // Required
	}

	if len(_iotsitewiseProjectId) > 0 {
		input.ProjectId = aws.String(_iotsitewiseProjectId)
	}
	if len(_iotsitewiseProjectName) > 0 {
		input.ProjectName = aws.String(_iotsitewiseProjectName)
	}
	if len(_iotsitewiseClientToken) > 0 {
		input.ClientToken = aws.String(_iotsitewiseClientToken)
	}
	if len(_iotsitewiseProjectDescription) > 0 {
		input.ProjectDescription = aws.String(_iotsitewiseProjectDescription)
	}

	if resp, err := client.UpdateProject(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_iotsitewiseCmd)
	_iotsitewiseCmd.Flags().SortFlags = false

	_iotsitewiseCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_iotsitewiseCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_iotsitewiseCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseAccessPolicyId, "access-policy-id", "", "", "Access Policy ID")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseAccessPolicyIdentity, "access-policy-identity", "", "", "Access Policy Identity")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseAccessPolicyPermission, "access-policy-permission", "", "", "Access Policy Permission")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseAccessPolicyResource, "access-policy-resource", "", "", "Access Policy Resource")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseActionDefinitionId, "action-definition-id", "", "", "Action Definition ID")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseActionId, "action-id", "", "", "Action ID")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseActionPayload, "action-payload", "", "", "Action Payload")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseActionType, "action-type", "", "", "Action Type")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseAdaptiveIngestion, "adaptive-ingestion", "", "", "Adaptive Ingestion")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseAggregateTypes, "aggregate-types", "", "", "Aggregate Types")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseAlarms, "alarms", "", "", "Alarms")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseAlias, "alias", "", "", "Alias")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseAliasPrefix, "alias-prefix", "", "", "Alias Prefix")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseAssetCompositeModelId, "asset-composite-model-id", "", "", "Asset Composite Model ID")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseAssetDescription, "asset-description", "", "", "Asset Description")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseAssetExternalId, "asset-external-id", "", "", "Asset External ID")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseAssetId, "asset-id", "", "", "Asset ID")
	_iotsitewiseCmd.Flags().StringSliceVarP(&_iotsitewiseAssetIds, "asset-ids", "", nil, "Asset Ids")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseAssetModelCompositeModelDescription, "asset-model-composite-model-description", "", "", "Asset Model Composite Model Description")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseAssetModelCompositeModelExternalId, "asset-model-composite-model-external-id", "", "", "Asset Model Composite Model External ID")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseAssetModelCompositeModelId, "asset-model-composite-model-id", "", "", "Asset Model Composite Model ID")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseAssetModelCompositeModelName, "asset-model-composite-model-name", "", "", "Asset Model Composite Model Name")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseAssetModelCompositeModelProperties, "asset-model-composite-model-properties", "", "", "Asset Model Composite Model Properties")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseAssetModelCompositeModelType, "asset-model-composite-model-type", "", "", "Asset Model Composite Model Type")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseAssetModelCompositeModels, "asset-model-composite-models", "", "", "Asset Model Composite Models")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseAssetModelDescription, "asset-model-description", "", "", "Asset Model Description")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseAssetModelExternalId, "asset-model-external-id", "", "", "Asset Model External ID")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseAssetModelHierarchies, "asset-model-hierarchies", "", "", "Asset Model Hierarchies")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseAssetModelId, "asset-model-id", "", "", "Asset Model ID")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseAssetModelName, "asset-model-name", "", "", "Asset Model Name")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseAssetModelProperties, "asset-model-properties", "", "", "Asset Model Properties")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseAssetModelType, "asset-model-type", "", "", "Asset Model Type")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseAssetModelTypes, "asset-model-types", "", "", "Asset Model Types")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseAssetModelVersion, "asset-model-version", "", "", "Asset Model Version")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseAssetName, "asset-name", "", "", "Asset Name")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseCapabilityConfiguration, "capability-configuration", "", "", "Capability Configuration")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseCapabilityNamespace, "capability-namespace", "", "", "Capability Namespace")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseChildAssetId, "child-asset-id", "", "", "Child Asset ID")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseClientToken, "client-token", "", "", "Client Token")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseComposedAssetModelId, "composed-asset-model-id", "", "", "Composed Asset Model ID")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseComputationModelConfiguration, "computation-model-configuration", "", "", "Computation Model Configuration")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseComputationModelDataBinding, "computation-model-data-binding", "", "", "Computation Model Data Binding")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseComputationModelDescription, "computation-model-description", "", "", "Computation Model Description")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseComputationModelId, "computation-model-id", "", "", "Computation Model ID")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseComputationModelName, "computation-model-name", "", "", "Computation Model Name")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseComputationModelType, "computation-model-type", "", "", "Computation Model Type")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseComputationModelVersion, "computation-model-version", "", "", "Computation Model Version")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseConversationId, "conversation-id", "", "", "Conversation ID")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseDashboardDefinition, "dashboard-definition", "", "", "Dashboard Definition")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseDashboardDescription, "dashboard-description", "", "", "Dashboard Description")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseDashboardId, "dashboard-id", "", "", "Dashboard ID")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseDashboardName, "dashboard-name", "", "", "Dashboard Name")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseDataBindingValueFilter, "data-binding-value-filter", "", "", "Data Binding Value Filter")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseDatasetDescription, "dataset-description", "", "", "Dataset Description")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseDatasetId, "dataset-id", "", "", "Dataset ID")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseDatasetName, "dataset-name", "", "", "Dataset Name")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseDatasetSource, "dataset-source", "", "", "Dataset Source")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseDeleteFilesAfterImport, "delete-files-after-import", "", "", "Delete Files After Import")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseDisallowIngestNullNaN, "disallow-ingest-null-na-n", "", "", "Disallow Ingest Null Na N")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseDisassociatedDataStorage, "disassociated-data-storage", "", "", "Disassociated Data Storage")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseEnablePartialEntryProcessing, "enable-partial-entry-processing", "", "", "Enable Partial Entry Processing")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseEnableTrace, "enable-trace", "", "", "Enable Trace")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseEncryptionType, "encryption-type", "", "", "Encryption Type")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseEndDate, "end-date", "", "", "End Date")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseEndTimeInSeconds, "end-time-in-seconds", "", "", "End Time In Seconds")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseEndTimeOffsetInNanos, "end-time-offset-in-nanos", "", "", "End Time Offset In Nanos")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseEntries, "entries", "", "", "Entries")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseErrorReportLocation, "error-report-location", "", "", "Error Report Location")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseExcludeProperties, "exclude-properties", "", "", "Exclude Properties")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseExecutionId, "execution-id", "", "", "Execution ID")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseFiles, "files", "", "", "Files")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseFilter, "filter", "", "", "Filter")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseGatewayId, "gateway-id", "", "", "Gateway ID")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseGatewayName, "gateway-name", "", "", "Gateway Name")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseGatewayPlatform, "gateway-platform", "", "", "Gateway Platform")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseGatewayVersion, "gateway-version", "", "", "Gateway Version")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseHierarchyId, "hierarchy-id", "", "", "Hierarchy ID")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseIamArn, "iam-arn", "", "", "IAM ARN")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseIdentityId, "identity-id", "", "", "Identity ID")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseIdentityType, "identity-type", "", "", "Identity Type")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseIfMatch, "if-match", "", "", "If Match")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseIfNoneMatch, "if-none-match", "", "", "If None Match")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseInterfaceAssetModelId, "interface-asset-model-id", "", "", "Interface Asset Model ID")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseIntervalInSeconds, "interval-in-seconds", "", "", "Interval In Seconds")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseIntervalWindowInSeconds, "interval-window-in-seconds", "", "", "Interval Window In Seconds")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseJobConfiguration, "job-configuration", "", "", "Job Configuration")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseJobId, "job-id", "", "", "Job ID")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseJobName, "job-name", "", "", "Job Name")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseJobRoleArn, "job-role-arn", "", "", "Job Role ARN")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseKmsKeyId, "kms-key-id", "", "", "KMS Key ID")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseLoggingOptions, "logging-options", "", "", "Logging Options")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseMatchForVersionType, "match-for-version-type", "", "", "Match For Version Type")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseMaxResults, "max-results", "", "", "Max Results")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseMessage, "message", "", "", "Message")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseMultiLayerStorage, "multi-layer-storage", "", "", "Multi Layer Storage")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseNextToken, "next-token", "", "", "Next Token")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseNotificationSenderEmail, "notification-sender-email", "", "", "Notification Sender Email")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseParentAssetModelCompositeModelId, "parent-asset-model-composite-model-id", "", "", "Parent Asset Model Composite Model ID")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewisePortalAuthMode, "portal-auth-mode", "", "", "Portal Auth Mode")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewisePortalContactEmail, "portal-contact-email", "", "", "Portal Contact Email")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewisePortalDescription, "portal-description", "", "", "Portal Description")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewisePortalId, "portal-id", "", "", "Portal ID")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewisePortalLogoImage, "portal-logo-image", "", "", "Portal Logo Image")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewisePortalLogoImageFile, "portal-logo-image-file", "", "", "Portal Logo Image File")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewisePortalName, "portal-name", "", "", "Portal Name")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewisePortalType, "portal-type", "", "", "Portal Type")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewisePortalTypeConfiguration, "portal-type-configuration", "", "", "Portal Type Configuration")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseProjectDescription, "project-description", "", "", "Project Description")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseProjectId, "project-id", "", "", "Project ID")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseProjectName, "project-name", "", "", "Project Name")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewisePropertyAlias, "property-alias", "", "", "Property Alias")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewisePropertyId, "property-id", "", "", "Property ID")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewisePropertyMappingConfiguration, "property-mapping-configuration", "", "", "Property Mapping Configuration")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewisePropertyNotificationState, "property-notification-state", "", "", "Property Notification State")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewisePropertyUnit, "property-unit", "", "", "Property Unit")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseQualities, "qualities", "", "", "Qualities")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseQuality, "quality", "", "", "Quality")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseQueryStatement, "query-statement", "", "", "Query Statement")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseResolution, "resolution", "", "", "Resolution")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseResolveTo, "resolve-to", "", "", "Resolve To")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseResolveToResourceId, "resolve-to-resource-id", "", "", "Resolve To Resource ID")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseResolveToResourceType, "resolve-to-resource-type", "", "", "Resolve To Resource Type")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseResourceArn, "resource-arn", "", "", "Resource ARN")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseResourceId, "resource-id", "", "", "Resource ID")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseResourceType, "resource-type", "", "", "Resource Type")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseRetentionPeriod, "retention-period", "", "", "Retention Period")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseRoleArn, "role-arn", "", "", "Role ARN")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseSourceType, "source-type", "", "", "Source Type")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseStartDate, "start-date", "", "", "Start Date")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseStartTimeInSeconds, "start-time-in-seconds", "", "", "Start Time In Seconds")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseStartTimeOffsetInNanos, "start-time-offset-in-nanos", "", "", "Start Time Offset In Nanos")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseStorageType, "storage-type", "", "", "Storage Type")
	_iotsitewiseCmd.Flags().StringSliceVarP(&_iotsitewiseTagKeys, "tag-keys", "", nil, "Tag Keys")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseTags, "tags", "", "", "Tags")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseTargetResource, "target-resource", "", "", "Target Resource")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseTargetResourceId, "target-resource-id", "", "", "Target Resource ID")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseTargetResourceType, "target-resource-type", "", "", "Target Resource Type")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseTimeOrdering, "time-ordering", "", "", "Time Ordering")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseTimeSeriesType, "time-series-type", "", "", "Time Series Type")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseTraversalDirection, "traversal-direction", "", "", "Traversal Direction")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseTraversalType, "traversal-type", "", "", "Traversal Type")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseType, "type", "", "", "Type")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseWarmTier, "warm-tier", "", "", "Warm Tier")
	_iotsitewiseCmd.Flags().StringVarP(&_iotsitewiseWarmTierRetentionPeriod, "warm-tier-retention-period", "", "", "Warm Tier Retention Period")

	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseAssociateAssets, "associate-assets", "", false, "Associate Assets")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseAssociateTimeSeriesToAssetProperty, "associate-time-series-to-asset-property", "", false, "Associate Time Series To Asset Property")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseBatchAssociateProjectAssets, "batch-associate-project-assets", "", false, "Batch Associate Project Assets")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseBatchDisassociateProjectAssets, "batch-disassociate-project-assets", "", false, "Batch Disassociate Project Assets")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseBatchGetAssetPropertyAggregates, "batch-get-asset-property-aggregates", "", false, "Batch Get Asset Property Aggregates")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseBatchGetAssetPropertyValue, "batch-get-asset-property-value", "", false, "Batch Get Asset Property Value")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseBatchGetAssetPropertyValueHistory, "batch-get-asset-property-value-history", "", false, "Batch Get Asset Property Value History")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseBatchPutAssetPropertyValue, "batch-put-asset-property-value", "", false, "Batch Put Asset Property Value")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseCreateAccessPolicy, "create-access-policy", "", false, "Create Access Policy")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseCreateAsset, "create-asset", "", false, "Create Asset")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseCreateAssetModel, "create-asset-model", "", false, "Create Asset Model")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseCreateAssetModelCompositeModel, "create-asset-model-composite-model", "", false, "Create Asset Model Composite Model")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseCreateBulkImportJob, "create-bulk-import-job", "", false, "Create Bulk Import Job")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseCreateComputationModel, "create-computation-model", "", false, "Create Computation Model")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseCreateDashboard, "create-dashboard", "", false, "Create Dashboard")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseCreateDataset, "create-dataset", "", false, "Create Dataset")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseCreateGateway, "create-gateway", "", false, "Create Gateway")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseCreatePortal, "create-portal", "", false, "Create Portal")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseCreateProject, "create-project", "", false, "Create Project")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseDeleteAccessPolicy, "delete-access-policy", "", false, "Delete Access Policy")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseDeleteAsset, "delete-asset", "", false, "Delete Asset")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseDeleteAssetModel, "delete-asset-model", "", false, "Delete Asset Model")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseDeleteAssetModelCompositeModel, "delete-asset-model-composite-model", "", false, "Delete Asset Model Composite Model")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseDeleteAssetModelInterfaceRelationship, "delete-asset-model-interface-relationship", "", false, "Delete Asset Model Interface Relationship")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseDeleteComputationModel, "delete-computation-model", "", false, "Delete Computation Model")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseDeleteDashboard, "delete-dashboard", "", false, "Delete Dashboard")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseDeleteDataset, "delete-dataset", "", false, "Delete Dataset")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseDeleteGateway, "delete-gateway", "", false, "Delete Gateway")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseDeletePortal, "delete-portal", "", false, "Delete Portal")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseDeleteProject, "delete-project", "", false, "Delete Project")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseDeleteTimeSeries, "delete-time-series", "", false, "Delete Time Series")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseDescribeAccessPolicy, "describe-access-policy", "", false, "Describe Access Policy")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseDescribeAction, "describe-action", "", false, "Describe Action")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseDescribeAsset, "describe-asset", "", false, "Describe Asset")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseDescribeAssetCompositeModel, "describe-asset-composite-model", "", false, "Describe Asset Composite Model")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseDescribeAssetModel, "describe-asset-model", "", false, "Describe Asset Model")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseDescribeAssetModelCompositeModel, "describe-asset-model-composite-model", "", false, "Describe Asset Model Composite Model")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseDescribeAssetModelInterfaceRelationship, "describe-asset-model-interface-relationship", "", false, "Describe Asset Model Interface Relationship")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseDescribeAssetProperty, "describe-asset-property", "", false, "Describe Asset Property")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseDescribeBulkImportJob, "describe-bulk-import-job", "", false, "Describe Bulk Import Job")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseDescribeComputationModel, "describe-computation-model", "", false, "Describe Computation Model")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseDescribeComputationModelExecutionSummary, "describe-computation-model-execution-summary", "", false, "Describe Computation Model Execution Summary")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseDescribeDashboard, "describe-dashboard", "", false, "Describe Dashboard")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseDescribeDataset, "describe-dataset", "", false, "Describe Dataset")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseDescribeDefaultEncryptionConfiguration, "describe-default-encryption-configuration", "", false, "Describe Default Encryption Configuration")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseDescribeExecution, "describe-execution", "", false, "Describe Execution")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseDescribeGateway, "describe-gateway", "", false, "Describe Gateway")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseDescribeGatewayCapabilityConfiguration, "describe-gateway-capability-configuration", "", false, "Describe Gateway Capability Configuration")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseDescribeLoggingOptions, "describe-logging-options", "", false, "Describe Logging Options")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseDescribePortal, "describe-portal", "", false, "Describe Portal")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseDescribeProject, "describe-project", "", false, "Describe Project")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseDescribeStorageConfiguration, "describe-storage-configuration", "", false, "Describe Storage Configuration")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseDescribeTimeSeries, "describe-time-series", "", false, "Describe Time Series")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseDisassociateAssets, "disassociate-assets", "", false, "Disassociate Assets")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseDisassociateTimeSeriesFromAssetProperty, "disassociate-time-series-from-asset-property", "", false, "Disassociate Time Series From Asset Property")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseExecuteAction, "execute-action", "", false, "Execute Action")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseExecuteQuery, "execute-query", "", false, "Execute Query")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseGetAssetPropertyAggregates, "get-asset-property-aggregates", "", false, "Get Asset Property Aggregates")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseGetAssetPropertyValue, "get-asset-property-value", "", false, "Get Asset Property Value")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseGetAssetPropertyValueHistory, "get-asset-property-value-history", "", false, "Get Asset Property Value History")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseGetInterpolatedAssetPropertyValues, "get-interpolated-asset-property-values", "", false, "Get Interpolated Asset Property Values")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseInvokeAssistant, "invoke-assistant", "", false, "Invoke Assistant")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseListAccessPolicies, "list-access-policies", "", false, "List Access Policies")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseListActions, "list-actions", "", false, "List Actions")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseListAssetModelCompositeModels, "list-asset-model-composite-models", "", false, "List Asset Model Composite Models")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseListAssetModelProperties, "list-asset-model-properties", "", false, "List Asset Model Properties")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseListAssetModels, "list-asset-models", "", false, "List Asset Models")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseListAssetProperties, "list-asset-properties", "", false, "List Asset Properties")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseListAssetRelationships, "list-asset-relationships", "", false, "List Asset Relationships")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseListAssets, "list-assets", "", false, "List Assets")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseListAssociatedAssets, "list-associated-assets", "", false, "List Associated Assets")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseListBulkImportJobs, "list-bulk-import-jobs", "", false, "List Bulk Import Jobs")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseListCompositionRelationships, "list-composition-relationships", "", false, "List Composition Relationships")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseListComputationModelDataBindingUsages, "list-computation-model-data-binding-usages", "", false, "List Computation Model Data Binding Usages")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseListComputationModelResolveToResources, "list-computation-model-resolve-to-resources", "", false, "List Computation Model Resolve To Resources")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseListComputationModels, "list-computation-models", "", false, "List Computation Models")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseListDashboards, "list-dashboards", "", false, "List Dashboards")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseListDatasets, "list-datasets", "", false, "List Datasets")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseListExecutions, "list-executions", "", false, "List Executions")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseListGateways, "list-gateways", "", false, "List Gateways")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseListInterfaceRelationships, "list-interface-relationships", "", false, "List Interface Relationships")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseListPortals, "list-portals", "", false, "List Portals")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseListProjectAssets, "list-project-assets", "", false, "List Project Assets")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseListProjects, "list-projects", "", false, "List Projects")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseListTimeSeries, "list-time-series", "", false, "List Time Series")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewisePutAssetModelInterfaceRelationship, "put-asset-model-interface-relationship", "", false, "Put Asset Model Interface Relationship")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewisePutDefaultEncryptionConfiguration, "put-default-encryption-configuration", "", false, "Put Default Encryption Configuration")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewisePutLoggingOptions, "put-logging-options", "", false, "Put Logging Options")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewisePutStorageConfiguration, "put-storage-configuration", "", false, "Put Storage Configuration")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseTagResource, "tag-resource", "", false, "Tag Resource")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseUntagResource, "untag-resource", "", false, "Untag Resource")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseUpdateAccessPolicy, "update-access-policy", "", false, "Update Access Policy")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseUpdateAsset, "update-asset", "", false, "Update Asset")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseUpdateAssetModel, "update-asset-model", "", false, "Update Asset Model")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseUpdateAssetModelCompositeModel, "update-asset-model-composite-model", "", false, "Update Asset Model Composite Model")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseUpdateAssetProperty, "update-asset-property", "", false, "Update Asset Property")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseUpdateComputationModel, "update-computation-model", "", false, "Update Computation Model")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseUpdateDashboard, "update-dashboard", "", false, "Update Dashboard")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseUpdateDataset, "update-dataset", "", false, "Update Dataset")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseUpdateGateway, "update-gateway", "", false, "Update Gateway")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseUpdateGatewayCapabilityConfiguration, "update-gateway-capability-configuration", "", false, "Update Gateway Capability Configuration")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseUpdatePortal, "update-portal", "", false, "Update Portal")
	_iotsitewiseCmd.Flags().BoolVarP(&_iotsitewiseUpdateProject, "update-project", "", false, "Update Project")

}
