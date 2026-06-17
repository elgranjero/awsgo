package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// sagemakerCmd represents the sagemaker command
var _sagemakerCmd = &cobra.Command{
	Use:   "sagemaker",
	Short: "AWS sagemaker CLI",
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
		client := sagemaker.NewFromConfig(cfg)
		if _sagemakerAddAssociation {
			sagemaker_AddAssociation(cfg, client)
			return
		}
		if _sagemakerAddTags {
			sagemaker_AddTags(cfg, client)
			return
		}
		if _sagemakerAssociateTrialComponent {
			sagemaker_AssociateTrialComponent(cfg, client)
			return
		}
		if _sagemakerAttachClusterNodeVolume {
			sagemaker_AttachClusterNodeVolume(cfg, client)
			return
		}
		if _sagemakerBatchAddClusterNodes {
			sagemaker_BatchAddClusterNodes(cfg, client)
			return
		}
		if _sagemakerBatchDeleteClusterNodes {
			sagemaker_BatchDeleteClusterNodes(cfg, client)
			return
		}
		if _sagemakerBatchDescribeModelPackage {
			sagemaker_BatchDescribeModelPackage(cfg, client)
			return
		}
		if _sagemakerBatchRebootClusterNodes {
			sagemaker_BatchRebootClusterNodes(cfg, client)
			return
		}
		if _sagemakerBatchReplaceClusterNodes {
			sagemaker_BatchReplaceClusterNodes(cfg, client)
			return
		}
		if _sagemakerCreateAction {
			sagemaker_CreateAction(cfg, client)
			return
		}
		if _sagemakerCreateAlgorithm {
			sagemaker_CreateAlgorithm(cfg, client)
			return
		}
		if _sagemakerCreateApp {
			sagemaker_CreateApp(cfg, client)
			return
		}
		if _sagemakerCreateAppImageConfig {
			sagemaker_CreateAppImageConfig(cfg, client)
			return
		}
		if _sagemakerCreateArtifact {
			sagemaker_CreateArtifact(cfg, client)
			return
		}
		if _sagemakerCreateAutoMLJob {
			sagemaker_CreateAutoMLJob(cfg, client)
			return
		}
		if _sagemakerCreateAutoMLJobV2 {
			sagemaker_CreateAutoMLJobV2(cfg, client)
			return
		}
		if _sagemakerCreateCluster {
			sagemaker_CreateCluster(cfg, client)
			return
		}
		if _sagemakerCreateClusterSchedulerConfig {
			sagemaker_CreateClusterSchedulerConfig(cfg, client)
			return
		}
		if _sagemakerCreateCodeRepository {
			sagemaker_CreateCodeRepository(cfg, client)
			return
		}
		if _sagemakerCreateCompilationJob {
			sagemaker_CreateCompilationJob(cfg, client)
			return
		}
		if _sagemakerCreateComputeQuota {
			sagemaker_CreateComputeQuota(cfg, client)
			return
		}
		if _sagemakerCreateContext {
			sagemaker_CreateContext(cfg, client)
			return
		}
		if _sagemakerCreateDataQualityJobDefinition {
			sagemaker_CreateDataQualityJobDefinition(cfg, client)
			return
		}
		if _sagemakerCreateDeviceFleet {
			sagemaker_CreateDeviceFleet(cfg, client)
			return
		}
		if _sagemakerCreateDomain {
			sagemaker_CreateDomain(cfg, client)
			return
		}
		if _sagemakerCreateEdgeDeploymentPlan {
			sagemaker_CreateEdgeDeploymentPlan(cfg, client)
			return
		}
		if _sagemakerCreateEdgeDeploymentStage {
			sagemaker_CreateEdgeDeploymentStage(cfg, client)
			return
		}
		if _sagemakerCreateEdgePackagingJob {
			sagemaker_CreateEdgePackagingJob(cfg, client)
			return
		}
		if _sagemakerCreateEndpoint {
			sagemaker_CreateEndpoint(cfg, client)
			return
		}
		if _sagemakerCreateEndpointConfig {
			sagemaker_CreateEndpointConfig(cfg, client)
			return
		}
		if _sagemakerCreateExperiment {
			sagemaker_CreateExperiment(cfg, client)
			return
		}
		if _sagemakerCreateFeatureGroup {
			sagemaker_CreateFeatureGroup(cfg, client)
			return
		}
		if _sagemakerCreateFlowDefinition {
			sagemaker_CreateFlowDefinition(cfg, client)
			return
		}
		if _sagemakerCreateHub {
			sagemaker_CreateHub(cfg, client)
			return
		}
		if _sagemakerCreateHubContentPresignedUrls {
			sagemaker_CreateHubContentPresignedUrls(cfg, client)
			return
		}
		if _sagemakerCreateHubContentReference {
			sagemaker_CreateHubContentReference(cfg, client)
			return
		}
		if _sagemakerCreateHumanTaskUi {
			sagemaker_CreateHumanTaskUi(cfg, client)
			return
		}
		if _sagemakerCreateHyperParameterTuningJob {
			sagemaker_CreateHyperParameterTuningJob(cfg, client)
			return
		}
		if _sagemakerCreateImage {
			sagemaker_CreateImage(cfg, client)
			return
		}
		if _sagemakerCreateImageVersion {
			sagemaker_CreateImageVersion(cfg, client)
			return
		}
		if _sagemakerCreateInferenceComponent {
			sagemaker_CreateInferenceComponent(cfg, client)
			return
		}
		if _sagemakerCreateInferenceExperiment {
			sagemaker_CreateInferenceExperiment(cfg, client)
			return
		}
		if _sagemakerCreateInferenceRecommendationsJob {
			sagemaker_CreateInferenceRecommendationsJob(cfg, client)
			return
		}
		if _sagemakerCreateLabelingJob {
			sagemaker_CreateLabelingJob(cfg, client)
			return
		}
		if _sagemakerCreateMlflowApp {
			sagemaker_CreateMlflowApp(cfg, client)
			return
		}
		if _sagemakerCreateMlflowTrackingServer {
			sagemaker_CreateMlflowTrackingServer(cfg, client)
			return
		}
		if _sagemakerCreateModel {
			sagemaker_CreateModel(cfg, client)
			return
		}
		if _sagemakerCreateModelBiasJobDefinition {
			sagemaker_CreateModelBiasJobDefinition(cfg, client)
			return
		}
		if _sagemakerCreateModelCard {
			sagemaker_CreateModelCard(cfg, client)
			return
		}
		if _sagemakerCreateModelCardExportJob {
			sagemaker_CreateModelCardExportJob(cfg, client)
			return
		}
		if _sagemakerCreateModelExplainabilityJobDefinition {
			sagemaker_CreateModelExplainabilityJobDefinition(cfg, client)
			return
		}
		if _sagemakerCreateModelPackage {
			sagemaker_CreateModelPackage(cfg, client)
			return
		}
		if _sagemakerCreateModelPackageGroup {
			sagemaker_CreateModelPackageGroup(cfg, client)
			return
		}
		if _sagemakerCreateModelQualityJobDefinition {
			sagemaker_CreateModelQualityJobDefinition(cfg, client)
			return
		}
		if _sagemakerCreateMonitoringSchedule {
			sagemaker_CreateMonitoringSchedule(cfg, client)
			return
		}
		if _sagemakerCreateNotebookInstance {
			sagemaker_CreateNotebookInstance(cfg, client)
			return
		}
		if _sagemakerCreateNotebookInstanceLifecycleConfig {
			sagemaker_CreateNotebookInstanceLifecycleConfig(cfg, client)
			return
		}
		if _sagemakerCreateOptimizationJob {
			sagemaker_CreateOptimizationJob(cfg, client)
			return
		}
		if _sagemakerCreatePartnerApp {
			sagemaker_CreatePartnerApp(cfg, client)
			return
		}
		if _sagemakerCreatePartnerAppPresignedUrl {
			sagemaker_CreatePartnerAppPresignedUrl(cfg, client)
			return
		}
		if _sagemakerCreatePipeline {
			sagemaker_CreatePipeline(cfg, client)
			return
		}
		if _sagemakerCreatePresignedDomainUrl {
			sagemaker_CreatePresignedDomainUrl(cfg, client)
			return
		}
		if _sagemakerCreatePresignedMlflowAppUrl {
			sagemaker_CreatePresignedMlflowAppUrl(cfg, client)
			return
		}
		if _sagemakerCreatePresignedMlflowTrackingServerUrl {
			sagemaker_CreatePresignedMlflowTrackingServerUrl(cfg, client)
			return
		}
		if _sagemakerCreatePresignedNotebookInstanceUrl {
			sagemaker_CreatePresignedNotebookInstanceUrl(cfg, client)
			return
		}
		if _sagemakerCreateProcessingJob {
			sagemaker_CreateProcessingJob(cfg, client)
			return
		}
		if _sagemakerCreateProject {
			sagemaker_CreateProject(cfg, client)
			return
		}
		if _sagemakerCreateSpace {
			sagemaker_CreateSpace(cfg, client)
			return
		}
		if _sagemakerCreateStudioLifecycleConfig {
			sagemaker_CreateStudioLifecycleConfig(cfg, client)
			return
		}
		if _sagemakerCreateTrainingJob {
			sagemaker_CreateTrainingJob(cfg, client)
			return
		}
		if _sagemakerCreateTrainingPlan {
			sagemaker_CreateTrainingPlan(cfg, client)
			return
		}
		if _sagemakerCreateTransformJob {
			sagemaker_CreateTransformJob(cfg, client)
			return
		}
		if _sagemakerCreateTrial {
			sagemaker_CreateTrial(cfg, client)
			return
		}
		if _sagemakerCreateTrialComponent {
			sagemaker_CreateTrialComponent(cfg, client)
			return
		}
		if _sagemakerCreateUserProfile {
			sagemaker_CreateUserProfile(cfg, client)
			return
		}
		if _sagemakerCreateWorkforce {
			sagemaker_CreateWorkforce(cfg, client)
			return
		}
		if _sagemakerCreateWorkteam {
			sagemaker_CreateWorkteam(cfg, client)
			return
		}
		if _sagemakerDeleteAction {
			sagemaker_DeleteAction(cfg, client)
			return
		}
		if _sagemakerDeleteAlgorithm {
			sagemaker_DeleteAlgorithm(cfg, client)
			return
		}
		if _sagemakerDeleteApp {
			sagemaker_DeleteApp(cfg, client)
			return
		}
		if _sagemakerDeleteAppImageConfig {
			sagemaker_DeleteAppImageConfig(cfg, client)
			return
		}
		if _sagemakerDeleteArtifact {
			sagemaker_DeleteArtifact(cfg, client)
			return
		}
		if _sagemakerDeleteAssociation {
			sagemaker_DeleteAssociation(cfg, client)
			return
		}
		if _sagemakerDeleteCluster {
			sagemaker_DeleteCluster(cfg, client)
			return
		}
		if _sagemakerDeleteClusterSchedulerConfig {
			sagemaker_DeleteClusterSchedulerConfig(cfg, client)
			return
		}
		if _sagemakerDeleteCodeRepository {
			sagemaker_DeleteCodeRepository(cfg, client)
			return
		}
		if _sagemakerDeleteCompilationJob {
			sagemaker_DeleteCompilationJob(cfg, client)
			return
		}
		if _sagemakerDeleteComputeQuota {
			sagemaker_DeleteComputeQuota(cfg, client)
			return
		}
		if _sagemakerDeleteContext {
			sagemaker_DeleteContext(cfg, client)
			return
		}
		if _sagemakerDeleteDataQualityJobDefinition {
			sagemaker_DeleteDataQualityJobDefinition(cfg, client)
			return
		}
		if _sagemakerDeleteDeviceFleet {
			sagemaker_DeleteDeviceFleet(cfg, client)
			return
		}
		if _sagemakerDeleteDomain {
			sagemaker_DeleteDomain(cfg, client)
			return
		}
		if _sagemakerDeleteEdgeDeploymentPlan {
			sagemaker_DeleteEdgeDeploymentPlan(cfg, client)
			return
		}
		if _sagemakerDeleteEdgeDeploymentStage {
			sagemaker_DeleteEdgeDeploymentStage(cfg, client)
			return
		}
		if _sagemakerDeleteEndpoint {
			sagemaker_DeleteEndpoint(cfg, client)
			return
		}
		if _sagemakerDeleteEndpointConfig {
			sagemaker_DeleteEndpointConfig(cfg, client)
			return
		}
		if _sagemakerDeleteExperiment {
			sagemaker_DeleteExperiment(cfg, client)
			return
		}
		if _sagemakerDeleteFeatureGroup {
			sagemaker_DeleteFeatureGroup(cfg, client)
			return
		}
		if _sagemakerDeleteFlowDefinition {
			sagemaker_DeleteFlowDefinition(cfg, client)
			return
		}
		if _sagemakerDeleteHub {
			sagemaker_DeleteHub(cfg, client)
			return
		}
		if _sagemakerDeleteHubContent {
			sagemaker_DeleteHubContent(cfg, client)
			return
		}
		if _sagemakerDeleteHubContentReference {
			sagemaker_DeleteHubContentReference(cfg, client)
			return
		}
		if _sagemakerDeleteHumanTaskUi {
			sagemaker_DeleteHumanTaskUi(cfg, client)
			return
		}
		if _sagemakerDeleteHyperParameterTuningJob {
			sagemaker_DeleteHyperParameterTuningJob(cfg, client)
			return
		}
		if _sagemakerDeleteImage {
			sagemaker_DeleteImage(cfg, client)
			return
		}
		if _sagemakerDeleteImageVersion {
			sagemaker_DeleteImageVersion(cfg, client)
			return
		}
		if _sagemakerDeleteInferenceComponent {
			sagemaker_DeleteInferenceComponent(cfg, client)
			return
		}
		if _sagemakerDeleteInferenceExperiment {
			sagemaker_DeleteInferenceExperiment(cfg, client)
			return
		}
		if _sagemakerDeleteMlflowApp {
			sagemaker_DeleteMlflowApp(cfg, client)
			return
		}
		if _sagemakerDeleteMlflowTrackingServer {
			sagemaker_DeleteMlflowTrackingServer(cfg, client)
			return
		}
		if _sagemakerDeleteModel {
			sagemaker_DeleteModel(cfg, client)
			return
		}
		if _sagemakerDeleteModelBiasJobDefinition {
			sagemaker_DeleteModelBiasJobDefinition(cfg, client)
			return
		}
		if _sagemakerDeleteModelCard {
			sagemaker_DeleteModelCard(cfg, client)
			return
		}
		if _sagemakerDeleteModelExplainabilityJobDefinition {
			sagemaker_DeleteModelExplainabilityJobDefinition(cfg, client)
			return
		}
		if _sagemakerDeleteModelPackage {
			sagemaker_DeleteModelPackage(cfg, client)
			return
		}
		if _sagemakerDeleteModelPackageGroup {
			sagemaker_DeleteModelPackageGroup(cfg, client)
			return
		}
		if _sagemakerDeleteModelPackageGroupPolicy {
			sagemaker_DeleteModelPackageGroupPolicy(cfg, client)
			return
		}
		if _sagemakerDeleteModelQualityJobDefinition {
			sagemaker_DeleteModelQualityJobDefinition(cfg, client)
			return
		}
		if _sagemakerDeleteMonitoringSchedule {
			sagemaker_DeleteMonitoringSchedule(cfg, client)
			return
		}
		if _sagemakerDeleteNotebookInstance {
			sagemaker_DeleteNotebookInstance(cfg, client)
			return
		}
		if _sagemakerDeleteNotebookInstanceLifecycleConfig {
			sagemaker_DeleteNotebookInstanceLifecycleConfig(cfg, client)
			return
		}
		if _sagemakerDeleteOptimizationJob {
			sagemaker_DeleteOptimizationJob(cfg, client)
			return
		}
		if _sagemakerDeletePartnerApp {
			sagemaker_DeletePartnerApp(cfg, client)
			return
		}
		if _sagemakerDeletePipeline {
			sagemaker_DeletePipeline(cfg, client)
			return
		}
		if _sagemakerDeleteProcessingJob {
			sagemaker_DeleteProcessingJob(cfg, client)
			return
		}
		if _sagemakerDeleteProject {
			sagemaker_DeleteProject(cfg, client)
			return
		}
		if _sagemakerDeleteSpace {
			sagemaker_DeleteSpace(cfg, client)
			return
		}
		if _sagemakerDeleteStudioLifecycleConfig {
			sagemaker_DeleteStudioLifecycleConfig(cfg, client)
			return
		}
		if _sagemakerDeleteTags {
			sagemaker_DeleteTags(cfg, client)
			return
		}
		if _sagemakerDeleteTrainingJob {
			sagemaker_DeleteTrainingJob(cfg, client)
			return
		}
		if _sagemakerDeleteTrial {
			sagemaker_DeleteTrial(cfg, client)
			return
		}
		if _sagemakerDeleteTrialComponent {
			sagemaker_DeleteTrialComponent(cfg, client)
			return
		}
		if _sagemakerDeleteUserProfile {
			sagemaker_DeleteUserProfile(cfg, client)
			return
		}
		if _sagemakerDeleteWorkforce {
			sagemaker_DeleteWorkforce(cfg, client)
			return
		}
		if _sagemakerDeleteWorkteam {
			sagemaker_DeleteWorkteam(cfg, client)
			return
		}
		if _sagemakerDeregisterDevices {
			sagemaker_DeregisterDevices(cfg, client)
			return
		}
		if _sagemakerDescribeAction {
			sagemaker_DescribeAction(cfg, client)
			return
		}
		if _sagemakerDescribeAlgorithm {
			sagemaker_DescribeAlgorithm(cfg, client)
			return
		}
		if _sagemakerDescribeApp {
			sagemaker_DescribeApp(cfg, client)
			return
		}
		if _sagemakerDescribeAppImageConfig {
			sagemaker_DescribeAppImageConfig(cfg, client)
			return
		}
		if _sagemakerDescribeArtifact {
			sagemaker_DescribeArtifact(cfg, client)
			return
		}
		if _sagemakerDescribeAutoMLJob {
			sagemaker_DescribeAutoMLJob(cfg, client)
			return
		}
		if _sagemakerDescribeAutoMLJobV2 {
			sagemaker_DescribeAutoMLJobV2(cfg, client)
			return
		}
		if _sagemakerDescribeCluster {
			sagemaker_DescribeCluster(cfg, client)
			return
		}
		if _sagemakerDescribeClusterEvent {
			sagemaker_DescribeClusterEvent(cfg, client)
			return
		}
		if _sagemakerDescribeClusterNode {
			sagemaker_DescribeClusterNode(cfg, client)
			return
		}
		if _sagemakerDescribeClusterSchedulerConfig {
			sagemaker_DescribeClusterSchedulerConfig(cfg, client)
			return
		}
		if _sagemakerDescribeCodeRepository {
			sagemaker_DescribeCodeRepository(cfg, client)
			return
		}
		if _sagemakerDescribeCompilationJob {
			sagemaker_DescribeCompilationJob(cfg, client)
			return
		}
		if _sagemakerDescribeComputeQuota {
			sagemaker_DescribeComputeQuota(cfg, client)
			return
		}
		if _sagemakerDescribeContext {
			sagemaker_DescribeContext(cfg, client)
			return
		}
		if _sagemakerDescribeDataQualityJobDefinition {
			sagemaker_DescribeDataQualityJobDefinition(cfg, client)
			return
		}
		if _sagemakerDescribeDevice {
			sagemaker_DescribeDevice(cfg, client)
			return
		}
		if _sagemakerDescribeDeviceFleet {
			sagemaker_DescribeDeviceFleet(cfg, client)
			return
		}
		if _sagemakerDescribeDomain {
			sagemaker_DescribeDomain(cfg, client)
			return
		}
		if _sagemakerDescribeEdgeDeploymentPlan {
			sagemaker_DescribeEdgeDeploymentPlan(cfg, client)
			return
		}
		if _sagemakerDescribeEdgePackagingJob {
			sagemaker_DescribeEdgePackagingJob(cfg, client)
			return
		}
		if _sagemakerDescribeEndpoint {
			sagemaker_DescribeEndpoint(cfg, client)
			return
		}
		if _sagemakerDescribeEndpointConfig {
			sagemaker_DescribeEndpointConfig(cfg, client)
			return
		}
		if _sagemakerDescribeExperiment {
			sagemaker_DescribeExperiment(cfg, client)
			return
		}
		if _sagemakerDescribeFeatureGroup {
			sagemaker_DescribeFeatureGroup(cfg, client)
			return
		}
		if _sagemakerDescribeFeatureMetadata {
			sagemaker_DescribeFeatureMetadata(cfg, client)
			return
		}
		if _sagemakerDescribeFlowDefinition {
			sagemaker_DescribeFlowDefinition(cfg, client)
			return
		}
		if _sagemakerDescribeHub {
			sagemaker_DescribeHub(cfg, client)
			return
		}
		if _sagemakerDescribeHubContent {
			sagemaker_DescribeHubContent(cfg, client)
			return
		}
		if _sagemakerDescribeHumanTaskUi {
			sagemaker_DescribeHumanTaskUi(cfg, client)
			return
		}
		if _sagemakerDescribeHyperParameterTuningJob {
			sagemaker_DescribeHyperParameterTuningJob(cfg, client)
			return
		}
		if _sagemakerDescribeImage {
			sagemaker_DescribeImage(cfg, client)
			return
		}
		if _sagemakerDescribeImageVersion {
			sagemaker_DescribeImageVersion(cfg, client)
			return
		}
		if _sagemakerDescribeInferenceComponent {
			sagemaker_DescribeInferenceComponent(cfg, client)
			return
		}
		if _sagemakerDescribeInferenceExperiment {
			sagemaker_DescribeInferenceExperiment(cfg, client)
			return
		}
		if _sagemakerDescribeInferenceRecommendationsJob {
			sagemaker_DescribeInferenceRecommendationsJob(cfg, client)
			return
		}
		if _sagemakerDescribeLabelingJob {
			sagemaker_DescribeLabelingJob(cfg, client)
			return
		}
		if _sagemakerDescribeLineageGroup {
			sagemaker_DescribeLineageGroup(cfg, client)
			return
		}
		if _sagemakerDescribeMlflowApp {
			sagemaker_DescribeMlflowApp(cfg, client)
			return
		}
		if _sagemakerDescribeMlflowTrackingServer {
			sagemaker_DescribeMlflowTrackingServer(cfg, client)
			return
		}
		if _sagemakerDescribeModel {
			sagemaker_DescribeModel(cfg, client)
			return
		}
		if _sagemakerDescribeModelBiasJobDefinition {
			sagemaker_DescribeModelBiasJobDefinition(cfg, client)
			return
		}
		if _sagemakerDescribeModelCard {
			sagemaker_DescribeModelCard(cfg, client)
			return
		}
		if _sagemakerDescribeModelCardExportJob {
			sagemaker_DescribeModelCardExportJob(cfg, client)
			return
		}
		if _sagemakerDescribeModelExplainabilityJobDefinition {
			sagemaker_DescribeModelExplainabilityJobDefinition(cfg, client)
			return
		}
		if _sagemakerDescribeModelPackage {
			sagemaker_DescribeModelPackage(cfg, client)
			return
		}
		if _sagemakerDescribeModelPackageGroup {
			sagemaker_DescribeModelPackageGroup(cfg, client)
			return
		}
		if _sagemakerDescribeModelQualityJobDefinition {
			sagemaker_DescribeModelQualityJobDefinition(cfg, client)
			return
		}
		if _sagemakerDescribeMonitoringSchedule {
			sagemaker_DescribeMonitoringSchedule(cfg, client)
			return
		}
		if _sagemakerDescribeNotebookInstance {
			sagemaker_DescribeNotebookInstance(cfg, client)
			return
		}
		if _sagemakerDescribeNotebookInstanceLifecycleConfig {
			sagemaker_DescribeNotebookInstanceLifecycleConfig(cfg, client)
			return
		}
		if _sagemakerDescribeOptimizationJob {
			sagemaker_DescribeOptimizationJob(cfg, client)
			return
		}
		if _sagemakerDescribePartnerApp {
			sagemaker_DescribePartnerApp(cfg, client)
			return
		}
		if _sagemakerDescribePipeline {
			sagemaker_DescribePipeline(cfg, client)
			return
		}
		if _sagemakerDescribePipelineDefinitionForExecution {
			sagemaker_DescribePipelineDefinitionForExecution(cfg, client)
			return
		}
		if _sagemakerDescribePipelineExecution {
			sagemaker_DescribePipelineExecution(cfg, client)
			return
		}
		if _sagemakerDescribeProcessingJob {
			sagemaker_DescribeProcessingJob(cfg, client)
			return
		}
		if _sagemakerDescribeProject {
			sagemaker_DescribeProject(cfg, client)
			return
		}
		if _sagemakerDescribeReservedCapacity {
			sagemaker_DescribeReservedCapacity(cfg, client)
			return
		}
		if _sagemakerDescribeSpace {
			sagemaker_DescribeSpace(cfg, client)
			return
		}
		if _sagemakerDescribeStudioLifecycleConfig {
			sagemaker_DescribeStudioLifecycleConfig(cfg, client)
			return
		}
		if _sagemakerDescribeSubscribedWorkteam {
			sagemaker_DescribeSubscribedWorkteam(cfg, client)
			return
		}
		if _sagemakerDescribeTrainingJob {
			sagemaker_DescribeTrainingJob(cfg, client)
			return
		}
		if _sagemakerDescribeTrainingPlan {
			sagemaker_DescribeTrainingPlan(cfg, client)
			return
		}
		if _sagemakerDescribeTransformJob {
			sagemaker_DescribeTransformJob(cfg, client)
			return
		}
		if _sagemakerDescribeTrial {
			sagemaker_DescribeTrial(cfg, client)
			return
		}
		if _sagemakerDescribeTrialComponent {
			sagemaker_DescribeTrialComponent(cfg, client)
			return
		}
		if _sagemakerDescribeUserProfile {
			sagemaker_DescribeUserProfile(cfg, client)
			return
		}
		if _sagemakerDescribeWorkforce {
			sagemaker_DescribeWorkforce(cfg, client)
			return
		}
		if _sagemakerDescribeWorkteam {
			sagemaker_DescribeWorkteam(cfg, client)
			return
		}
		if _sagemakerDetachClusterNodeVolume {
			sagemaker_DetachClusterNodeVolume(cfg, client)
			return
		}
		if _sagemakerDisableSagemakerServicecatalogPortfolio {
			sagemaker_DisableSagemakerServicecatalogPortfolio(cfg, client)
			return
		}
		if _sagemakerDisassociateTrialComponent {
			sagemaker_DisassociateTrialComponent(cfg, client)
			return
		}
		if _sagemakerEnableSagemakerServicecatalogPortfolio {
			sagemaker_EnableSagemakerServicecatalogPortfolio(cfg, client)
			return
		}
		if _sagemakerGetDeviceFleetReport {
			sagemaker_GetDeviceFleetReport(cfg, client)
			return
		}
		if _sagemakerGetLineageGroupPolicy {
			sagemaker_GetLineageGroupPolicy(cfg, client)
			return
		}
		if _sagemakerGetModelPackageGroupPolicy {
			sagemaker_GetModelPackageGroupPolicy(cfg, client)
			return
		}
		if _sagemakerGetSagemakerServicecatalogPortfolioStatus {
			sagemaker_GetSagemakerServicecatalogPortfolioStatus(cfg, client)
			return
		}
		if _sagemakerGetScalingConfigurationRecommendation {
			sagemaker_GetScalingConfigurationRecommendation(cfg, client)
			return
		}
		if _sagemakerGetSearchSuggestions {
			sagemaker_GetSearchSuggestions(cfg, client)
			return
		}
		if _sagemakerImportHubContent {
			sagemaker_ImportHubContent(cfg, client)
			return
		}
		if _sagemakerListActions {
			sagemaker_ListActions(cfg, client)
			return
		}
		if _sagemakerListAlgorithms {
			sagemaker_ListAlgorithms(cfg, client)
			return
		}
		if _sagemakerListAliases {
			sagemaker_ListAliases(cfg, client)
			return
		}
		if _sagemakerListAppImageConfigs {
			sagemaker_ListAppImageConfigs(cfg, client)
			return
		}
		if _sagemakerListApps {
			sagemaker_ListApps(cfg, client)
			return
		}
		if _sagemakerListArtifacts {
			sagemaker_ListArtifacts(cfg, client)
			return
		}
		if _sagemakerListAssociations {
			sagemaker_ListAssociations(cfg, client)
			return
		}
		if _sagemakerListAutoMLJobs {
			sagemaker_ListAutoMLJobs(cfg, client)
			return
		}
		if _sagemakerListCandidatesForAutoMLJob {
			sagemaker_ListCandidatesForAutoMLJob(cfg, client)
			return
		}
		if _sagemakerListClusterEvents {
			sagemaker_ListClusterEvents(cfg, client)
			return
		}
		if _sagemakerListClusterNodes {
			sagemaker_ListClusterNodes(cfg, client)
			return
		}
		if _sagemakerListClusterSchedulerConfigs {
			sagemaker_ListClusterSchedulerConfigs(cfg, client)
			return
		}
		if _sagemakerListClusters {
			sagemaker_ListClusters(cfg, client)
			return
		}
		if _sagemakerListCodeRepositories {
			sagemaker_ListCodeRepositories(cfg, client)
			return
		}
		if _sagemakerListCompilationJobs {
			sagemaker_ListCompilationJobs(cfg, client)
			return
		}
		if _sagemakerListComputeQuotas {
			sagemaker_ListComputeQuotas(cfg, client)
			return
		}
		if _sagemakerListContexts {
			sagemaker_ListContexts(cfg, client)
			return
		}
		if _sagemakerListDataQualityJobDefinitions {
			sagemaker_ListDataQualityJobDefinitions(cfg, client)
			return
		}
		if _sagemakerListDeviceFleets {
			sagemaker_ListDeviceFleets(cfg, client)
			return
		}
		if _sagemakerListDevices {
			sagemaker_ListDevices(cfg, client)
			return
		}
		if _sagemakerListDomains {
			sagemaker_ListDomains(cfg, client)
			return
		}
		if _sagemakerListEdgeDeploymentPlans {
			sagemaker_ListEdgeDeploymentPlans(cfg, client)
			return
		}
		if _sagemakerListEdgePackagingJobs {
			sagemaker_ListEdgePackagingJobs(cfg, client)
			return
		}
		if _sagemakerListEndpointConfigs {
			sagemaker_ListEndpointConfigs(cfg, client)
			return
		}
		if _sagemakerListEndpoints {
			sagemaker_ListEndpoints(cfg, client)
			return
		}
		if _sagemakerListExperiments {
			sagemaker_ListExperiments(cfg, client)
			return
		}
		if _sagemakerListFeatureGroups {
			sagemaker_ListFeatureGroups(cfg, client)
			return
		}
		if _sagemakerListFlowDefinitions {
			sagemaker_ListFlowDefinitions(cfg, client)
			return
		}
		if _sagemakerListHubContentVersions {
			sagemaker_ListHubContentVersions(cfg, client)
			return
		}
		if _sagemakerListHubContents {
			sagemaker_ListHubContents(cfg, client)
			return
		}
		if _sagemakerListHubs {
			sagemaker_ListHubs(cfg, client)
			return
		}
		if _sagemakerListHumanTaskUis {
			sagemaker_ListHumanTaskUis(cfg, client)
			return
		}
		if _sagemakerListHyperParameterTuningJobs {
			sagemaker_ListHyperParameterTuningJobs(cfg, client)
			return
		}
		if _sagemakerListImageVersions {
			sagemaker_ListImageVersions(cfg, client)
			return
		}
		if _sagemakerListImages {
			sagemaker_ListImages(cfg, client)
			return
		}
		if _sagemakerListInferenceComponents {
			sagemaker_ListInferenceComponents(cfg, client)
			return
		}
		if _sagemakerListInferenceExperiments {
			sagemaker_ListInferenceExperiments(cfg, client)
			return
		}
		if _sagemakerListInferenceRecommendationsJobSteps {
			sagemaker_ListInferenceRecommendationsJobSteps(cfg, client)
			return
		}
		if _sagemakerListInferenceRecommendationsJobs {
			sagemaker_ListInferenceRecommendationsJobs(cfg, client)
			return
		}
		if _sagemakerListLabelingJobs {
			sagemaker_ListLabelingJobs(cfg, client)
			return
		}
		if _sagemakerListLabelingJobsForWorkteam {
			sagemaker_ListLabelingJobsForWorkteam(cfg, client)
			return
		}
		if _sagemakerListLineageGroups {
			sagemaker_ListLineageGroups(cfg, client)
			return
		}
		if _sagemakerListMlflowApps {
			sagemaker_ListMlflowApps(cfg, client)
			return
		}
		if _sagemakerListMlflowTrackingServers {
			sagemaker_ListMlflowTrackingServers(cfg, client)
			return
		}
		if _sagemakerListModelBiasJobDefinitions {
			sagemaker_ListModelBiasJobDefinitions(cfg, client)
			return
		}
		if _sagemakerListModelCardExportJobs {
			sagemaker_ListModelCardExportJobs(cfg, client)
			return
		}
		if _sagemakerListModelCardVersions {
			sagemaker_ListModelCardVersions(cfg, client)
			return
		}
		if _sagemakerListModelCards {
			sagemaker_ListModelCards(cfg, client)
			return
		}
		if _sagemakerListModelExplainabilityJobDefinitions {
			sagemaker_ListModelExplainabilityJobDefinitions(cfg, client)
			return
		}
		if _sagemakerListModelMetadata {
			sagemaker_ListModelMetadata(cfg, client)
			return
		}
		if _sagemakerListModelPackageGroups {
			sagemaker_ListModelPackageGroups(cfg, client)
			return
		}
		if _sagemakerListModelPackages {
			sagemaker_ListModelPackages(cfg, client)
			return
		}
		if _sagemakerListModelQualityJobDefinitions {
			sagemaker_ListModelQualityJobDefinitions(cfg, client)
			return
		}
		if _sagemakerListModels {
			sagemaker_ListModels(cfg, client)
			return
		}
		if _sagemakerListMonitoringAlertHistory {
			sagemaker_ListMonitoringAlertHistory(cfg, client)
			return
		}
		if _sagemakerListMonitoringAlerts {
			sagemaker_ListMonitoringAlerts(cfg, client)
			return
		}
		if _sagemakerListMonitoringExecutions {
			sagemaker_ListMonitoringExecutions(cfg, client)
			return
		}
		if _sagemakerListMonitoringSchedules {
			sagemaker_ListMonitoringSchedules(cfg, client)
			return
		}
		if _sagemakerListNotebookInstanceLifecycleConfigs {
			sagemaker_ListNotebookInstanceLifecycleConfigs(cfg, client)
			return
		}
		if _sagemakerListNotebookInstances {
			sagemaker_ListNotebookInstances(cfg, client)
			return
		}
		if _sagemakerListOptimizationJobs {
			sagemaker_ListOptimizationJobs(cfg, client)
			return
		}
		if _sagemakerListPartnerApps {
			sagemaker_ListPartnerApps(cfg, client)
			return
		}
		if _sagemakerListPipelineExecutionSteps {
			sagemaker_ListPipelineExecutionSteps(cfg, client)
			return
		}
		if _sagemakerListPipelineExecutions {
			sagemaker_ListPipelineExecutions(cfg, client)
			return
		}
		if _sagemakerListPipelineParametersForExecution {
			sagemaker_ListPipelineParametersForExecution(cfg, client)
			return
		}
		if _sagemakerListPipelineVersions {
			sagemaker_ListPipelineVersions(cfg, client)
			return
		}
		if _sagemakerListPipelines {
			sagemaker_ListPipelines(cfg, client)
			return
		}
		if _sagemakerListProcessingJobs {
			sagemaker_ListProcessingJobs(cfg, client)
			return
		}
		if _sagemakerListProjects {
			sagemaker_ListProjects(cfg, client)
			return
		}
		if _sagemakerListResourceCatalogs {
			sagemaker_ListResourceCatalogs(cfg, client)
			return
		}
		if _sagemakerListSpaces {
			sagemaker_ListSpaces(cfg, client)
			return
		}
		if _sagemakerListStageDevices {
			sagemaker_ListStageDevices(cfg, client)
			return
		}
		if _sagemakerListStudioLifecycleConfigs {
			sagemaker_ListStudioLifecycleConfigs(cfg, client)
			return
		}
		if _sagemakerListSubscribedWorkteams {
			sagemaker_ListSubscribedWorkteams(cfg, client)
			return
		}
		if _sagemakerListTags {
			sagemaker_ListTags(cfg, client)
			return
		}
		if _sagemakerListTrainingJobs {
			sagemaker_ListTrainingJobs(cfg, client)
			return
		}
		if _sagemakerListTrainingJobsForHyperParameterTuningJob {
			sagemaker_ListTrainingJobsForHyperParameterTuningJob(cfg, client)
			return
		}
		if _sagemakerListTrainingPlans {
			sagemaker_ListTrainingPlans(cfg, client)
			return
		}
		if _sagemakerListTransformJobs {
			sagemaker_ListTransformJobs(cfg, client)
			return
		}
		if _sagemakerListTrialComponents {
			sagemaker_ListTrialComponents(cfg, client)
			return
		}
		if _sagemakerListTrials {
			sagemaker_ListTrials(cfg, client)
			return
		}
		if _sagemakerListUltraServersByReservedCapacity {
			sagemaker_ListUltraServersByReservedCapacity(cfg, client)
			return
		}
		if _sagemakerListUserProfiles {
			sagemaker_ListUserProfiles(cfg, client)
			return
		}
		if _sagemakerListWorkforces {
			sagemaker_ListWorkforces(cfg, client)
			return
		}
		if _sagemakerListWorkteams {
			sagemaker_ListWorkteams(cfg, client)
			return
		}
		if _sagemakerPutModelPackageGroupPolicy {
			sagemaker_PutModelPackageGroupPolicy(cfg, client)
			return
		}
		if _sagemakerQueryLineage {
			sagemaker_QueryLineage(cfg, client)
			return
		}
		if _sagemakerRegisterDevices {
			sagemaker_RegisterDevices(cfg, client)
			return
		}
		if _sagemakerRenderUiTemplate {
			sagemaker_RenderUiTemplate(cfg, client)
			return
		}
		if _sagemakerRetryPipelineExecution {
			sagemaker_RetryPipelineExecution(cfg, client)
			return
		}
		if _sagemakerSearch {
			sagemaker_Search(cfg, client)
			return
		}
		if _sagemakerSearchTrainingPlanOfferings {
			sagemaker_SearchTrainingPlanOfferings(cfg, client)
			return
		}
		if _sagemakerSendPipelineExecutionStepFailure {
			sagemaker_SendPipelineExecutionStepFailure(cfg, client)
			return
		}
		if _sagemakerSendPipelineExecutionStepSuccess {
			sagemaker_SendPipelineExecutionStepSuccess(cfg, client)
			return
		}
		if _sagemakerStartEdgeDeploymentStage {
			sagemaker_StartEdgeDeploymentStage(cfg, client)
			return
		}
		if _sagemakerStartInferenceExperiment {
			sagemaker_StartInferenceExperiment(cfg, client)
			return
		}
		if _sagemakerStartMlflowTrackingServer {
			sagemaker_StartMlflowTrackingServer(cfg, client)
			return
		}
		if _sagemakerStartMonitoringSchedule {
			sagemaker_StartMonitoringSchedule(cfg, client)
			return
		}
		if _sagemakerStartNotebookInstance {
			sagemaker_StartNotebookInstance(cfg, client)
			return
		}
		if _sagemakerStartPipelineExecution {
			sagemaker_StartPipelineExecution(cfg, client)
			return
		}
		if _sagemakerStartSession {
			sagemaker_StartSession(cfg, client)
			return
		}
		if _sagemakerStopAutoMLJob {
			sagemaker_StopAutoMLJob(cfg, client)
			return
		}
		if _sagemakerStopCompilationJob {
			sagemaker_StopCompilationJob(cfg, client)
			return
		}
		if _sagemakerStopEdgeDeploymentStage {
			sagemaker_StopEdgeDeploymentStage(cfg, client)
			return
		}
		if _sagemakerStopEdgePackagingJob {
			sagemaker_StopEdgePackagingJob(cfg, client)
			return
		}
		if _sagemakerStopHyperParameterTuningJob {
			sagemaker_StopHyperParameterTuningJob(cfg, client)
			return
		}
		if _sagemakerStopInferenceExperiment {
			sagemaker_StopInferenceExperiment(cfg, client)
			return
		}
		if _sagemakerStopInferenceRecommendationsJob {
			sagemaker_StopInferenceRecommendationsJob(cfg, client)
			return
		}
		if _sagemakerStopLabelingJob {
			sagemaker_StopLabelingJob(cfg, client)
			return
		}
		if _sagemakerStopMlflowTrackingServer {
			sagemaker_StopMlflowTrackingServer(cfg, client)
			return
		}
		if _sagemakerStopMonitoringSchedule {
			sagemaker_StopMonitoringSchedule(cfg, client)
			return
		}
		if _sagemakerStopNotebookInstance {
			sagemaker_StopNotebookInstance(cfg, client)
			return
		}
		if _sagemakerStopOptimizationJob {
			sagemaker_StopOptimizationJob(cfg, client)
			return
		}
		if _sagemakerStopPipelineExecution {
			sagemaker_StopPipelineExecution(cfg, client)
			return
		}
		if _sagemakerStopProcessingJob {
			sagemaker_StopProcessingJob(cfg, client)
			return
		}
		if _sagemakerStopTrainingJob {
			sagemaker_StopTrainingJob(cfg, client)
			return
		}
		if _sagemakerStopTransformJob {
			sagemaker_StopTransformJob(cfg, client)
			return
		}
		if _sagemakerUpdateAction {
			sagemaker_UpdateAction(cfg, client)
			return
		}
		if _sagemakerUpdateAppImageConfig {
			sagemaker_UpdateAppImageConfig(cfg, client)
			return
		}
		if _sagemakerUpdateArtifact {
			sagemaker_UpdateArtifact(cfg, client)
			return
		}
		if _sagemakerUpdateCluster {
			sagemaker_UpdateCluster(cfg, client)
			return
		}
		if _sagemakerUpdateClusterSchedulerConfig {
			sagemaker_UpdateClusterSchedulerConfig(cfg, client)
			return
		}
		if _sagemakerUpdateClusterSoftware {
			sagemaker_UpdateClusterSoftware(cfg, client)
			return
		}
		if _sagemakerUpdateCodeRepository {
			sagemaker_UpdateCodeRepository(cfg, client)
			return
		}
		if _sagemakerUpdateComputeQuota {
			sagemaker_UpdateComputeQuota(cfg, client)
			return
		}
		if _sagemakerUpdateContext {
			sagemaker_UpdateContext(cfg, client)
			return
		}
		if _sagemakerUpdateDeviceFleet {
			sagemaker_UpdateDeviceFleet(cfg, client)
			return
		}
		if _sagemakerUpdateDevices {
			sagemaker_UpdateDevices(cfg, client)
			return
		}
		if _sagemakerUpdateDomain {
			sagemaker_UpdateDomain(cfg, client)
			return
		}
		if _sagemakerUpdateEndpoint {
			sagemaker_UpdateEndpoint(cfg, client)
			return
		}
		if _sagemakerUpdateEndpointWeightsAndCapacities {
			sagemaker_UpdateEndpointWeightsAndCapacities(cfg, client)
			return
		}
		if _sagemakerUpdateExperiment {
			sagemaker_UpdateExperiment(cfg, client)
			return
		}
		if _sagemakerUpdateFeatureGroup {
			sagemaker_UpdateFeatureGroup(cfg, client)
			return
		}
		if _sagemakerUpdateFeatureMetadata {
			sagemaker_UpdateFeatureMetadata(cfg, client)
			return
		}
		if _sagemakerUpdateHub {
			sagemaker_UpdateHub(cfg, client)
			return
		}
		if _sagemakerUpdateHubContent {
			sagemaker_UpdateHubContent(cfg, client)
			return
		}
		if _sagemakerUpdateHubContentReference {
			sagemaker_UpdateHubContentReference(cfg, client)
			return
		}
		if _sagemakerUpdateImage {
			sagemaker_UpdateImage(cfg, client)
			return
		}
		if _sagemakerUpdateImageVersion {
			sagemaker_UpdateImageVersion(cfg, client)
			return
		}
		if _sagemakerUpdateInferenceComponent {
			sagemaker_UpdateInferenceComponent(cfg, client)
			return
		}
		if _sagemakerUpdateInferenceComponentRuntimeConfig {
			sagemaker_UpdateInferenceComponentRuntimeConfig(cfg, client)
			return
		}
		if _sagemakerUpdateInferenceExperiment {
			sagemaker_UpdateInferenceExperiment(cfg, client)
			return
		}
		if _sagemakerUpdateMlflowApp {
			sagemaker_UpdateMlflowApp(cfg, client)
			return
		}
		if _sagemakerUpdateMlflowTrackingServer {
			sagemaker_UpdateMlflowTrackingServer(cfg, client)
			return
		}
		if _sagemakerUpdateModelCard {
			sagemaker_UpdateModelCard(cfg, client)
			return
		}
		if _sagemakerUpdateModelPackage {
			sagemaker_UpdateModelPackage(cfg, client)
			return
		}
		if _sagemakerUpdateMonitoringAlert {
			sagemaker_UpdateMonitoringAlert(cfg, client)
			return
		}
		if _sagemakerUpdateMonitoringSchedule {
			sagemaker_UpdateMonitoringSchedule(cfg, client)
			return
		}
		if _sagemakerUpdateNotebookInstance {
			sagemaker_UpdateNotebookInstance(cfg, client)
			return
		}
		if _sagemakerUpdateNotebookInstanceLifecycleConfig {
			sagemaker_UpdateNotebookInstanceLifecycleConfig(cfg, client)
			return
		}
		if _sagemakerUpdatePartnerApp {
			sagemaker_UpdatePartnerApp(cfg, client)
			return
		}
		if _sagemakerUpdatePipeline {
			sagemaker_UpdatePipeline(cfg, client)
			return
		}
		if _sagemakerUpdatePipelineExecution {
			sagemaker_UpdatePipelineExecution(cfg, client)
			return
		}
		if _sagemakerUpdatePipelineVersion {
			sagemaker_UpdatePipelineVersion(cfg, client)
			return
		}
		if _sagemakerUpdateProject {
			sagemaker_UpdateProject(cfg, client)
			return
		}
		if _sagemakerUpdateSpace {
			sagemaker_UpdateSpace(cfg, client)
			return
		}
		if _sagemakerUpdateTrainingJob {
			sagemaker_UpdateTrainingJob(cfg, client)
			return
		}
		if _sagemakerUpdateTrial {
			sagemaker_UpdateTrial(cfg, client)
			return
		}
		if _sagemakerUpdateTrialComponent {
			sagemaker_UpdateTrialComponent(cfg, client)
			return
		}
		if _sagemakerUpdateUserProfile {
			sagemaker_UpdateUserProfile(cfg, client)
			return
		}
		if _sagemakerUpdateWorkforce {
			sagemaker_UpdateWorkforce(cfg, client)
			return
		}
		if _sagemakerUpdateWorkteam {
			sagemaker_UpdateWorkteam(cfg, client)
			return
		}

	},
}

var (
	_sagemakerAddAssociation                             bool
	_sagemakerAddTags                                    bool
	_sagemakerAssociateTrialComponent                    bool
	_sagemakerAttachClusterNodeVolume                    bool
	_sagemakerBatchAddClusterNodes                       bool
	_sagemakerBatchDeleteClusterNodes                    bool
	_sagemakerBatchDescribeModelPackage                  bool
	_sagemakerBatchRebootClusterNodes                    bool
	_sagemakerBatchReplaceClusterNodes                   bool
	_sagemakerCreateAction                               bool
	_sagemakerCreateAlgorithm                            bool
	_sagemakerCreateApp                                  bool
	_sagemakerCreateAppImageConfig                       bool
	_sagemakerCreateArtifact                             bool
	_sagemakerCreateAutoMLJob                            bool
	_sagemakerCreateAutoMLJobV2                          bool
	_sagemakerCreateCluster                              bool
	_sagemakerCreateClusterSchedulerConfig               bool
	_sagemakerCreateCodeRepository                       bool
	_sagemakerCreateCompilationJob                       bool
	_sagemakerCreateComputeQuota                         bool
	_sagemakerCreateContext                              bool
	_sagemakerCreateDataQualityJobDefinition             bool
	_sagemakerCreateDeviceFleet                          bool
	_sagemakerCreateDomain                               bool
	_sagemakerCreateEdgeDeploymentPlan                   bool
	_sagemakerCreateEdgeDeploymentStage                  bool
	_sagemakerCreateEdgePackagingJob                     bool
	_sagemakerCreateEndpoint                             bool
	_sagemakerCreateEndpointConfig                       bool
	_sagemakerCreateExperiment                           bool
	_sagemakerCreateFeatureGroup                         bool
	_sagemakerCreateFlowDefinition                       bool
	_sagemakerCreateHub                                  bool
	_sagemakerCreateHubContentPresignedUrls              bool
	_sagemakerCreateHubContentReference                  bool
	_sagemakerCreateHumanTaskUi                          bool
	_sagemakerCreateHyperParameterTuningJob              bool
	_sagemakerCreateImage                                bool
	_sagemakerCreateImageVersion                         bool
	_sagemakerCreateInferenceComponent                   bool
	_sagemakerCreateInferenceExperiment                  bool
	_sagemakerCreateInferenceRecommendationsJob          bool
	_sagemakerCreateLabelingJob                          bool
	_sagemakerCreateMlflowApp                            bool
	_sagemakerCreateMlflowTrackingServer                 bool
	_sagemakerCreateModel                                bool
	_sagemakerCreateModelBiasJobDefinition               bool
	_sagemakerCreateModelCard                            bool
	_sagemakerCreateModelCardExportJob                   bool
	_sagemakerCreateModelExplainabilityJobDefinition     bool
	_sagemakerCreateModelPackage                         bool
	_sagemakerCreateModelPackageGroup                    bool
	_sagemakerCreateModelQualityJobDefinition            bool
	_sagemakerCreateMonitoringSchedule                   bool
	_sagemakerCreateNotebookInstance                     bool
	_sagemakerCreateNotebookInstanceLifecycleConfig      bool
	_sagemakerCreateOptimizationJob                      bool
	_sagemakerCreatePartnerApp                           bool
	_sagemakerCreatePartnerAppPresignedUrl               bool
	_sagemakerCreatePipeline                             bool
	_sagemakerCreatePresignedDomainUrl                   bool
	_sagemakerCreatePresignedMlflowAppUrl                bool
	_sagemakerCreatePresignedMlflowTrackingServerUrl     bool
	_sagemakerCreatePresignedNotebookInstanceUrl         bool
	_sagemakerCreateProcessingJob                        bool
	_sagemakerCreateProject                              bool
	_sagemakerCreateSpace                                bool
	_sagemakerCreateStudioLifecycleConfig                bool
	_sagemakerCreateTrainingJob                          bool
	_sagemakerCreateTrainingPlan                         bool
	_sagemakerCreateTransformJob                         bool
	_sagemakerCreateTrial                                bool
	_sagemakerCreateTrialComponent                       bool
	_sagemakerCreateUserProfile                          bool
	_sagemakerCreateWorkforce                            bool
	_sagemakerCreateWorkteam                             bool
	_sagemakerDeleteAction                               bool
	_sagemakerDeleteAlgorithm                            bool
	_sagemakerDeleteApp                                  bool
	_sagemakerDeleteAppImageConfig                       bool
	_sagemakerDeleteArtifact                             bool
	_sagemakerDeleteAssociation                          bool
	_sagemakerDeleteCluster                              bool
	_sagemakerDeleteClusterSchedulerConfig               bool
	_sagemakerDeleteCodeRepository                       bool
	_sagemakerDeleteCompilationJob                       bool
	_sagemakerDeleteComputeQuota                         bool
	_sagemakerDeleteContext                              bool
	_sagemakerDeleteDataQualityJobDefinition             bool
	_sagemakerDeleteDeviceFleet                          bool
	_sagemakerDeleteDomain                               bool
	_sagemakerDeleteEdgeDeploymentPlan                   bool
	_sagemakerDeleteEdgeDeploymentStage                  bool
	_sagemakerDeleteEndpoint                             bool
	_sagemakerDeleteEndpointConfig                       bool
	_sagemakerDeleteExperiment                           bool
	_sagemakerDeleteFeatureGroup                         bool
	_sagemakerDeleteFlowDefinition                       bool
	_sagemakerDeleteHub                                  bool
	_sagemakerDeleteHubContent                           bool
	_sagemakerDeleteHubContentReference                  bool
	_sagemakerDeleteHumanTaskUi                          bool
	_sagemakerDeleteHyperParameterTuningJob              bool
	_sagemakerDeleteImage                                bool
	_sagemakerDeleteImageVersion                         bool
	_sagemakerDeleteInferenceComponent                   bool
	_sagemakerDeleteInferenceExperiment                  bool
	_sagemakerDeleteMlflowApp                            bool
	_sagemakerDeleteMlflowTrackingServer                 bool
	_sagemakerDeleteModel                                bool
	_sagemakerDeleteModelBiasJobDefinition               bool
	_sagemakerDeleteModelCard                            bool
	_sagemakerDeleteModelExplainabilityJobDefinition     bool
	_sagemakerDeleteModelPackage                         bool
	_sagemakerDeleteModelPackageGroup                    bool
	_sagemakerDeleteModelPackageGroupPolicy              bool
	_sagemakerDeleteModelQualityJobDefinition            bool
	_sagemakerDeleteMonitoringSchedule                   bool
	_sagemakerDeleteNotebookInstance                     bool
	_sagemakerDeleteNotebookInstanceLifecycleConfig      bool
	_sagemakerDeleteOptimizationJob                      bool
	_sagemakerDeletePartnerApp                           bool
	_sagemakerDeletePipeline                             bool
	_sagemakerDeleteProcessingJob                        bool
	_sagemakerDeleteProject                              bool
	_sagemakerDeleteSpace                                bool
	_sagemakerDeleteStudioLifecycleConfig                bool
	_sagemakerDeleteTags                                 bool
	_sagemakerDeleteTrainingJob                          bool
	_sagemakerDeleteTrial                                bool
	_sagemakerDeleteTrialComponent                       bool
	_sagemakerDeleteUserProfile                          bool
	_sagemakerDeleteWorkforce                            bool
	_sagemakerDeleteWorkteam                             bool
	_sagemakerDeregisterDevices                          bool
	_sagemakerDescribeAction                             bool
	_sagemakerDescribeAlgorithm                          bool
	_sagemakerDescribeApp                                bool
	_sagemakerDescribeAppImageConfig                     bool
	_sagemakerDescribeArtifact                           bool
	_sagemakerDescribeAutoMLJob                          bool
	_sagemakerDescribeAutoMLJobV2                        bool
	_sagemakerDescribeCluster                            bool
	_sagemakerDescribeClusterEvent                       bool
	_sagemakerDescribeClusterNode                        bool
	_sagemakerDescribeClusterSchedulerConfig             bool
	_sagemakerDescribeCodeRepository                     bool
	_sagemakerDescribeCompilationJob                     bool
	_sagemakerDescribeComputeQuota                       bool
	_sagemakerDescribeContext                            bool
	_sagemakerDescribeDataQualityJobDefinition           bool
	_sagemakerDescribeDevice                             bool
	_sagemakerDescribeDeviceFleet                        bool
	_sagemakerDescribeDomain                             bool
	_sagemakerDescribeEdgeDeploymentPlan                 bool
	_sagemakerDescribeEdgePackagingJob                   bool
	_sagemakerDescribeEndpoint                           bool
	_sagemakerDescribeEndpointConfig                     bool
	_sagemakerDescribeExperiment                         bool
	_sagemakerDescribeFeatureGroup                       bool
	_sagemakerDescribeFeatureMetadata                    bool
	_sagemakerDescribeFlowDefinition                     bool
	_sagemakerDescribeHub                                bool
	_sagemakerDescribeHubContent                         bool
	_sagemakerDescribeHumanTaskUi                        bool
	_sagemakerDescribeHyperParameterTuningJob            bool
	_sagemakerDescribeImage                              bool
	_sagemakerDescribeImageVersion                       bool
	_sagemakerDescribeInferenceComponent                 bool
	_sagemakerDescribeInferenceExperiment                bool
	_sagemakerDescribeInferenceRecommendationsJob        bool
	_sagemakerDescribeLabelingJob                        bool
	_sagemakerDescribeLineageGroup                       bool
	_sagemakerDescribeMlflowApp                          bool
	_sagemakerDescribeMlflowTrackingServer               bool
	_sagemakerDescribeModel                              bool
	_sagemakerDescribeModelBiasJobDefinition             bool
	_sagemakerDescribeModelCard                          bool
	_sagemakerDescribeModelCardExportJob                 bool
	_sagemakerDescribeModelExplainabilityJobDefinition   bool
	_sagemakerDescribeModelPackage                       bool
	_sagemakerDescribeModelPackageGroup                  bool
	_sagemakerDescribeModelQualityJobDefinition          bool
	_sagemakerDescribeMonitoringSchedule                 bool
	_sagemakerDescribeNotebookInstance                   bool
	_sagemakerDescribeNotebookInstanceLifecycleConfig    bool
	_sagemakerDescribeOptimizationJob                    bool
	_sagemakerDescribePartnerApp                         bool
	_sagemakerDescribePipeline                           bool
	_sagemakerDescribePipelineDefinitionForExecution     bool
	_sagemakerDescribePipelineExecution                  bool
	_sagemakerDescribeProcessingJob                      bool
	_sagemakerDescribeProject                            bool
	_sagemakerDescribeReservedCapacity                   bool
	_sagemakerDescribeSpace                              bool
	_sagemakerDescribeStudioLifecycleConfig              bool
	_sagemakerDescribeSubscribedWorkteam                 bool
	_sagemakerDescribeTrainingJob                        bool
	_sagemakerDescribeTrainingPlan                       bool
	_sagemakerDescribeTransformJob                       bool
	_sagemakerDescribeTrial                              bool
	_sagemakerDescribeTrialComponent                     bool
	_sagemakerDescribeUserProfile                        bool
	_sagemakerDescribeWorkforce                          bool
	_sagemakerDescribeWorkteam                           bool
	_sagemakerDetachClusterNodeVolume                    bool
	_sagemakerDisableSagemakerServicecatalogPortfolio    bool
	_sagemakerDisassociateTrialComponent                 bool
	_sagemakerEnableSagemakerServicecatalogPortfolio     bool
	_sagemakerGetDeviceFleetReport                       bool
	_sagemakerGetLineageGroupPolicy                      bool
	_sagemakerGetModelPackageGroupPolicy                 bool
	_sagemakerGetSagemakerServicecatalogPortfolioStatus  bool
	_sagemakerGetScalingConfigurationRecommendation      bool
	_sagemakerGetSearchSuggestions                       bool
	_sagemakerImportHubContent                           bool
	_sagemakerListActions                                bool
	_sagemakerListAlgorithms                             bool
	_sagemakerListAliases                                bool
	_sagemakerListAppImageConfigs                        bool
	_sagemakerListApps                                   bool
	_sagemakerListArtifacts                              bool
	_sagemakerListAssociations                           bool
	_sagemakerListAutoMLJobs                             bool
	_sagemakerListCandidatesForAutoMLJob                 bool
	_sagemakerListClusterEvents                          bool
	_sagemakerListClusterNodes                           bool
	_sagemakerListClusterSchedulerConfigs                bool
	_sagemakerListClusters                               bool
	_sagemakerListCodeRepositories                       bool
	_sagemakerListCompilationJobs                        bool
	_sagemakerListComputeQuotas                          bool
	_sagemakerListContexts                               bool
	_sagemakerListDataQualityJobDefinitions              bool
	_sagemakerListDeviceFleets                           bool
	_sagemakerListDevices                                bool
	_sagemakerListDomains                                bool
	_sagemakerListEdgeDeploymentPlans                    bool
	_sagemakerListEdgePackagingJobs                      bool
	_sagemakerListEndpointConfigs                        bool
	_sagemakerListEndpoints                              bool
	_sagemakerListExperiments                            bool
	_sagemakerListFeatureGroups                          bool
	_sagemakerListFlowDefinitions                        bool
	_sagemakerListHubContentVersions                     bool
	_sagemakerListHubContents                            bool
	_sagemakerListHubs                                   bool
	_sagemakerListHumanTaskUis                           bool
	_sagemakerListHyperParameterTuningJobs               bool
	_sagemakerListImageVersions                          bool
	_sagemakerListImages                                 bool
	_sagemakerListInferenceComponents                    bool
	_sagemakerListInferenceExperiments                   bool
	_sagemakerListInferenceRecommendationsJobSteps       bool
	_sagemakerListInferenceRecommendationsJobs           bool
	_sagemakerListLabelingJobs                           bool
	_sagemakerListLabelingJobsForWorkteam                bool
	_sagemakerListLineageGroups                          bool
	_sagemakerListMlflowApps                             bool
	_sagemakerListMlflowTrackingServers                  bool
	_sagemakerListModelBiasJobDefinitions                bool
	_sagemakerListModelCardExportJobs                    bool
	_sagemakerListModelCardVersions                      bool
	_sagemakerListModelCards                             bool
	_sagemakerListModelExplainabilityJobDefinitions      bool
	_sagemakerListModelMetadata                          bool
	_sagemakerListModelPackageGroups                     bool
	_sagemakerListModelPackages                          bool
	_sagemakerListModelQualityJobDefinitions             bool
	_sagemakerListModels                                 bool
	_sagemakerListMonitoringAlertHistory                 bool
	_sagemakerListMonitoringAlerts                       bool
	_sagemakerListMonitoringExecutions                   bool
	_sagemakerListMonitoringSchedules                    bool
	_sagemakerListNotebookInstanceLifecycleConfigs       bool
	_sagemakerListNotebookInstances                      bool
	_sagemakerListOptimizationJobs                       bool
	_sagemakerListPartnerApps                            bool
	_sagemakerListPipelineExecutionSteps                 bool
	_sagemakerListPipelineExecutions                     bool
	_sagemakerListPipelineParametersForExecution         bool
	_sagemakerListPipelineVersions                       bool
	_sagemakerListPipelines                              bool
	_sagemakerListProcessingJobs                         bool
	_sagemakerListProjects                               bool
	_sagemakerListResourceCatalogs                       bool
	_sagemakerListSpaces                                 bool
	_sagemakerListStageDevices                           bool
	_sagemakerListStudioLifecycleConfigs                 bool
	_sagemakerListSubscribedWorkteams                    bool
	_sagemakerListTags                                   bool
	_sagemakerListTrainingJobs                           bool
	_sagemakerListTrainingJobsForHyperParameterTuningJob bool
	_sagemakerListTrainingPlans                          bool
	_sagemakerListTransformJobs                          bool
	_sagemakerListTrialComponents                        bool
	_sagemakerListTrials                                 bool
	_sagemakerListUltraServersByReservedCapacity         bool
	_sagemakerListUserProfiles                           bool
	_sagemakerListWorkforces                             bool
	_sagemakerListWorkteams                              bool
	_sagemakerPutModelPackageGroupPolicy                 bool
	_sagemakerQueryLineage                               bool
	_sagemakerRegisterDevices                            bool
	_sagemakerRenderUiTemplate                           bool
	_sagemakerRetryPipelineExecution                     bool
	_sagemakerSearch                                     bool
	_sagemakerSearchTrainingPlanOfferings                bool
	_sagemakerSendPipelineExecutionStepFailure           bool
	_sagemakerSendPipelineExecutionStepSuccess           bool
	_sagemakerStartEdgeDeploymentStage                   bool
	_sagemakerStartInferenceExperiment                   bool
	_sagemakerStartMlflowTrackingServer                  bool
	_sagemakerStartMonitoringSchedule                    bool
	_sagemakerStartNotebookInstance                      bool
	_sagemakerStartPipelineExecution                     bool
	_sagemakerStartSession                               bool
	_sagemakerStopAutoMLJob                              bool
	_sagemakerStopCompilationJob                         bool
	_sagemakerStopEdgeDeploymentStage                    bool
	_sagemakerStopEdgePackagingJob                       bool
	_sagemakerStopHyperParameterTuningJob                bool
	_sagemakerStopInferenceExperiment                    bool
	_sagemakerStopInferenceRecommendationsJob            bool
	_sagemakerStopLabelingJob                            bool
	_sagemakerStopMlflowTrackingServer                   bool
	_sagemakerStopMonitoringSchedule                     bool
	_sagemakerStopNotebookInstance                       bool
	_sagemakerStopOptimizationJob                        bool
	_sagemakerStopPipelineExecution                      bool
	_sagemakerStopProcessingJob                          bool
	_sagemakerStopTrainingJob                            bool
	_sagemakerStopTransformJob                           bool
	_sagemakerUpdateAction                               bool
	_sagemakerUpdateAppImageConfig                       bool
	_sagemakerUpdateArtifact                             bool
	_sagemakerUpdateCluster                              bool
	_sagemakerUpdateClusterSchedulerConfig               bool
	_sagemakerUpdateClusterSoftware                      bool
	_sagemakerUpdateCodeRepository                       bool
	_sagemakerUpdateComputeQuota                         bool
	_sagemakerUpdateContext                              bool
	_sagemakerUpdateDeviceFleet                          bool
	_sagemakerUpdateDevices                              bool
	_sagemakerUpdateDomain                               bool
	_sagemakerUpdateEndpoint                             bool
	_sagemakerUpdateEndpointWeightsAndCapacities         bool
	_sagemakerUpdateExperiment                           bool
	_sagemakerUpdateFeatureGroup                         bool
	_sagemakerUpdateFeatureMetadata                      bool
	_sagemakerUpdateHub                                  bool
	_sagemakerUpdateHubContent                           bool
	_sagemakerUpdateHubContentReference                  bool
	_sagemakerUpdateImage                                bool
	_sagemakerUpdateImageVersion                         bool
	_sagemakerUpdateInferenceComponent                   bool
	_sagemakerUpdateInferenceComponentRuntimeConfig      bool
	_sagemakerUpdateInferenceExperiment                  bool
	_sagemakerUpdateMlflowApp                            bool
	_sagemakerUpdateMlflowTrackingServer                 bool
	_sagemakerUpdateModelCard                            bool
	_sagemakerUpdateModelPackage                         bool
	_sagemakerUpdateMonitoringAlert                      bool
	_sagemakerUpdateMonitoringSchedule                   bool
	_sagemakerUpdateNotebookInstance                     bool
	_sagemakerUpdateNotebookInstanceLifecycleConfig      bool
	_sagemakerUpdatePartnerApp                           bool
	_sagemakerUpdatePipeline                             bool
	_sagemakerUpdatePipelineExecution                    bool
	_sagemakerUpdatePipelineVersion                      bool
	_sagemakerUpdateProject                              bool
	_sagemakerUpdateSpace                                bool
	_sagemakerUpdateTrainingJob                          bool
	_sagemakerUpdateTrial                                bool
	_sagemakerUpdateTrialComponent                       bool
	_sagemakerUpdateUserProfile                          bool
	_sagemakerUpdateWorkforce                            bool
	_sagemakerUpdateWorkteam                             bool

	_sagemakerAcceleratorTypes                            string
	_sagemakerAccessConfig                                string
	_sagemakerAccountDefaultStatus                        string
	_sagemakerActionName                                  string
	_sagemakerActionType                                  string
	_sagemakerActivationState                             string
	_sagemakerAdditionalCodeRepositories                  []string
	_sagemakerAdditionalCodeRepositoryEquals              string
	_sagemakerAdditionalInferenceSpecifications           string
	_sagemakerAdditionalInferenceSpecificationsToAdd      string
	_sagemakerAlgorithmDescription                        string
	_sagemakerAlgorithmName                               string
	_sagemakerAlgorithmSpecification                      string
	_sagemakerAlias                                       string
	_sagemakerAliases                                     []string
	_sagemakerAliasesToAdd                                []string
	_sagemakerAliasesToDelete                             []string
	_sagemakerAppImageConfigName                          string
	_sagemakerAppName                                     string
	_sagemakerAppNetworkAccessType                        string
	_sagemakerAppSecurityGroupManagement                  string
	_sagemakerAppSpecification                            string
	_sagemakerAppType                                     string
	_sagemakerAppTypeEquals                               string
	_sagemakerAppVersion                                  string
	_sagemakerApplicationConfig                           string
	_sagemakerApprovalDescription                         string
	_sagemakerArn                                         string
	_sagemakerArtifactArn                                 string
	_sagemakerArtifactName                                string
	_sagemakerArtifactStoreUri                            string
	_sagemakerArtifactType                                string
	_sagemakerAssociationType                             string
	_sagemakerAsyncInferenceConfig                        string
	_sagemakerAuthMode                                    string
	_sagemakerAuthType                                    string
	_sagemakerAutoMLComputeConfig                         string
	_sagemakerAutoMLJobConfig                             string
	_sagemakerAutoMLJobInputDataConfig                    string
	_sagemakerAutoMLJobName                               string
	_sagemakerAutoMLJobObjective                          string
	_sagemakerAutoMLProblemTypeConfig                     string
	_sagemakerAutoScaling                                 string
	_sagemakerAutomaticModelRegistration                  string
	_sagemakerAutotune                                    string
	_sagemakerBaseImage                                   string
	_sagemakerBatchStrategy                               string
	_sagemakerCallbackToken                               string
	_sagemakerCandidateNameEquals                         string
	_sagemakerCertifyForMarketplace                       string
	_sagemakerCheckpointConfig                            string
	_sagemakerClientRequestToken                          string
	_sagemakerClientToken                                 string
	_sagemakerClusterArn                                  string
	_sagemakerClusterName                                 string
	_sagemakerClusterRole                                 string
	_sagemakerClusterSchedulerConfigId                    string
	_sagemakerClusterSchedulerConfigVersion               string
	_sagemakerCodeEditorAppImageConfig                    string
	_sagemakerCodeRepositoryName                          string
	_sagemakerCognitoConfig                               string
	_sagemakerCompilationJobName                          string
	_sagemakerComputeQuotaConfig                          string
	_sagemakerComputeQuotaId                              string
	_sagemakerComputeQuotaTarget                          string
	_sagemakerComputeQuotaVersion                         string
	_sagemakerContainers                                  string
	_sagemakerContent                                     string
	_sagemakerContextName                                 string
	_sagemakerContextType                                 string
	_sagemakerCreatedAfter                                string
	_sagemakerCreatedBefore                               string
	_sagemakerCreationTimeAfter                           string
	_sagemakerCreationTimeBefore                          string
	_sagemakerCrossAccountFilterOption                    string
	_sagemakerCustomerMetadataProperties                  string
	_sagemakerCustomerMetadataPropertiesToRemove          []string
	_sagemakerDataCaptureConfig                           string
	_sagemakerDataProcessing                              string
	_sagemakerDataQualityAppSpecification                 string
	_sagemakerDataQualityBaselineConfig                   string
	_sagemakerDataQualityJobInput                         string
	_sagemakerDataQualityJobOutputConfig                  string
	_sagemakerDataSplitConfig                             string
	_sagemakerDataStorageConfig                           string
	_sagemakerDatapointsToAlert                           string
	_sagemakerDebugHookConfig                             string
	_sagemakerDebugRuleConfigurations                     string
	_sagemakerDefaultCodeRepository                       string
	_sagemakerDefaultCodeRepositoryContains               string
	_sagemakerDefaultDomainIdList                         []string
	_sagemakerDefaultForDomainId                          string
	_sagemakerDefaultSpaceSettings                        string
	_sagemakerDefaultUserSettings                         string
	_sagemakerDeleteProperties                            []string
	_sagemakerDeploymentConfig                            string
	_sagemakerDeploymentInstanceType                      string
	_sagemakerDescription                                 string
	_sagemakerDesiredModelVariants                        string
	_sagemakerDesiredRuntimeConfig                        string
	_sagemakerDesiredState                                string
	_sagemakerDesiredWeightsAndCapacities                 string
	_sagemakerDestinationArn                              string
	_sagemakerDestinationType                             string
	_sagemakerDeviceFleetName                             string
	_sagemakerDeviceFleetNameContains                     string
	_sagemakerDeviceName                                  string
	_sagemakerDeviceNames                                 []string
	_sagemakerDevices                                     string
	_sagemakerDirectInternetAccess                        string
	_sagemakerDirection                                   string
	_sagemakerDisassociateAcceleratorTypes                string
	_sagemakerDisassociateAdditionalCodeRepositories      string
	_sagemakerDisassociateDefaultCodeRepository           string
	_sagemakerDisassociateLifecycleConfig                 string
	_sagemakerDisplayName                                 string
	_sagemakerDocumentSchemaVersion                       string
	_sagemakerDomain                                      string
	_sagemakerDomainId                                    string
	_sagemakerDomainIdEquals                              string
	_sagemakerDomainName                                  string
	_sagemakerDomainSettings                              string
	_sagemakerDomainSettingsForUpdate                     string
	_sagemakerDriftCheckBaselines                         string
	_sagemakerDurationHours                               string
	_sagemakerEdgeDeploymentPlanName                      string
	_sagemakerEdgePackagingJobName                        string
	_sagemakerEnableAutoMinorVersionUpgrade               string
	_sagemakerEnableIamSessionBasedIdentity               string
	_sagemakerEnableInterContainerTrafficEncryption       string
	_sagemakerEnableIotRoleAlias                          string
	_sagemakerEnableManagedSpotTraining                   string
	_sagemakerEnableNetworkIsolation                      string
	_sagemakerEndTime                                     string
	_sagemakerEndTimeBefore                               string
	_sagemakerEndpointConfigName                          string
	_sagemakerEndpointName                                string
	_sagemakerEndpointNameEquals                          string
	_sagemakerEnvironment                                 string
	_sagemakerEvaluationPeriod                            string
	_sagemakerEventId                                     string
	_sagemakerEventTimeAfter                              string
	_sagemakerEventTimeBefore                             string
	_sagemakerEventTimeFeatureName                        string
	_sagemakerExcludeDevicesDeployedInOtherStage          string
	_sagemakerExcludeRetainedVariantProperties            string
	_sagemakerExecutionRoleArn                            string
	_sagemakerExperimentConfig                            string
	_sagemakerExperimentName                              string
	_sagemakerExpiresInSeconds                            string
	_sagemakerExplainerConfig                             string
	_sagemakerFailureReason                               string
	_sagemakerFeatureAdditions                            string
	_sagemakerFeatureDefinitions                          string
	_sagemakerFeatureGroupName                            string
	_sagemakerFeatureGroupStatusEquals                    string
	_sagemakerFeatureName                                 string
	_sagemakerFilters                                     string
	_sagemakerFlowDefinitionName                          string
	_sagemakerGenerateCandidateDefinitionsOnly            string
	_sagemakerGitConfig                                   string
	_sagemakerHomeEfsFileSystemKmsKeyId                   string
	_sagemakerHorovod                                     string
	_sagemakerHubContentDescription                       string
	_sagemakerHubContentDisplayName                       string
	_sagemakerHubContentDocument                          string
	_sagemakerHubContentMarkdown                          string
	_sagemakerHubContentName                              string
	_sagemakerHubContentSearchKeywords                    []string
	_sagemakerHubContentType                              string
	_sagemakerHubContentVersion                           string
	_sagemakerHubDescription                              string
	_sagemakerHubDisplayName                              string
	_sagemakerHubName                                     string
	_sagemakerHubSearchKeywords                           []string
	_sagemakerHumanLoopActivationConfig                   string
	_sagemakerHumanLoopConfig                             string
	_sagemakerHumanLoopRequestSource                      string
	_sagemakerHumanTaskConfig                             string
	_sagemakerHumanTaskUiArn                              string
	_sagemakerHumanTaskUiName                             string
	_sagemakerHyperParameterTuningJobConfig               string
	_sagemakerHyperParameterTuningJobName                 string
	_sagemakerHyperParameters                             string
	_sagemakerImageId                                     string
	_sagemakerImageName                                   string
	_sagemakerIncludeAvailableUpgrade                     string
	_sagemakerIncludeEdges                                string
	_sagemakerIncludeNodeLogicalIds                       string
	_sagemakerInferenceComponentName                      string
	_sagemakerInferenceExecutionConfig                    string
	_sagemakerInferenceRecommendationsJobName             string
	_sagemakerInferenceSpecification                      string
	_sagemakerInfraCheckConfig                            string
	_sagemakerInputArtifacts                              string
	_sagemakerInputArtifactsToRemove                      []string
	_sagemakerInputConfig                                 string
	_sagemakerInputDataConfig                             string
	_sagemakerInstanceCount                               string
	_sagemakerInstanceGroupName                           string
	_sagemakerInstanceGroupNameContains                   string
	_sagemakerInstanceGroups                              string
	_sagemakerInstanceGroupsToDelete                      []string
	_sagemakerInstanceMetadataServiceConfiguration        string
	_sagemakerInstanceType                                string
	_sagemakerIpAddressType                               string
	_sagemakerJobDefinitionName                           string
	_sagemakerJobDescription                              string
	_sagemakerJobName                                     string
	_sagemakerJobReferenceCodeContains                    string
	_sagemakerJobResources                                string
	_sagemakerJobType                                     string
	_sagemakerJupyterLabAppImageConfig                    string
	_sagemakerKernelGatewayImageConfig                    string
	_sagemakerKmsKey                                      string
	_sagemakerKmsKeyId                                    string
	_sagemakerLabelAttributeName                          string
	_sagemakerLabelCategoryConfigS3Uri                    string
	_sagemakerLabelingJobAlgorithmsConfig                 string
	_sagemakerLabelingJobName                             string
	_sagemakerLandingUri                                  string
	_sagemakerLastModifiedTimeAfter                       string
	_sagemakerLastModifiedTimeBefore                      string
	_sagemakerLatestHeartbeatAfter                        string
	_sagemakerLifecycleConfigName                         string
	_sagemakerLineageGroupName                            string
	_sagemakerMaintenanceConfig                           string
	_sagemakerMaxConcurrentTransforms                     string
	_sagemakerMaxDepth                                    string
	_sagemakerMaxInstanceCount                            string
	_sagemakerMaxPayloadInMB                              string
	_sagemakerMaxResults                                  string
	_sagemakerMaxSchemaVersion                            string
	_sagemakerMemberDefinitions                           string
	_sagemakerMetadataProperties                          string
	_sagemakerMetricsConfig                               string
	_sagemakerMinVersion                                  string
	_sagemakerMLFramework                                 string
	_sagemakerMlflowConfig                                string
	_sagemakerMlflowExperimentName                        string
	_sagemakerMlflowVersion                               string
	_sagemakerModelApprovalStatus                         string
	_sagemakerModelBiasAppSpecification                   string
	_sagemakerModelBiasBaselineConfig                     string
	_sagemakerModelBiasJobInput                           string
	_sagemakerModelBiasJobOutputConfig                    string
	_sagemakerModelCard                                   string
	_sagemakerModelCardExportJobArn                       string
	_sagemakerModelCardExportJobName                      string
	_sagemakerModelCardExportJobNameContains              string
	_sagemakerModelCardName                               string
	_sagemakerModelCardStatus                             string
	_sagemakerModelCardVersion                            string
	_sagemakerModelClientConfig                           string
	_sagemakerModelConfigs                                string
	_sagemakerModelDeployConfig                           string
	_sagemakerModelExplainabilityAppSpecification         string
	_sagemakerModelExplainabilityBaselineConfig           string
	_sagemakerModelExplainabilityJobInput                 string
	_sagemakerModelExplainabilityJobOutputConfig          string
	_sagemakerModelLifeCycle                              string
	_sagemakerModelMetrics                                string
	_sagemakerModelName                                   string
	_sagemakerModelNameContains                           string
	_sagemakerModelNameEquals                             string
	_sagemakerModelPackageArn                             string
	_sagemakerModelPackageArnList                         []string
	_sagemakerModelPackageConfig                          string
	_sagemakerModelPackageDescription                     string
	_sagemakerModelPackageGroupDescription                string
	_sagemakerModelPackageGroupName                       string
	_sagemakerModelPackageName                            string
	_sagemakerModelPackageRegistrationType                string
	_sagemakerModelPackageType                            string
	_sagemakerModelPackageVersionArn                      string
	_sagemakerModelPackageVersionArnEquals                string
	_sagemakerModelQualityAppSpecification                string
	_sagemakerModelQualityBaselineConfig                  string
	_sagemakerModelQualityJobInput                        string
	_sagemakerModelQualityJobOutputConfig                 string
	_sagemakerModelRegistrationMode                       string
	_sagemakerModelSource                                 string
	_sagemakerModelVariantActions                         string
	_sagemakerModelVariants                               string
	_sagemakerModelVersion                                string
	_sagemakerModifiedTimeAfter                           string
	_sagemakerModifiedTimeBefore                          string
	_sagemakerMonitoringAlertName                         string
	_sagemakerMonitoringJobDefinitionName                 string
	_sagemakerMonitoringScheduleConfig                    string
	_sagemakerMonitoringScheduleName                      string
	_sagemakerMonitoringTypeEquals                        string
	_sagemakerName                                        string
	_sagemakerNameContains                                string
	_sagemakerNetworkConfig                               string
	_sagemakerNextToken                                   string
	_sagemakerNodeId                                      string
	_sagemakerNodeIds                                     []string
	_sagemakerNodeLogicalId                               string
	_sagemakerNodeLogicalIds                              []string
	_sagemakerNodeProvisioningMode                        string
	_sagemakerNodeRecovery                                string
	_sagemakerNodesToAdd                                  string
	_sagemakerNotebookInstanceLifecycleConfigName         string
	_sagemakerNotebookInstanceLifecycleConfigNameContains string
	_sagemakerNotebookInstanceName                        string
	_sagemakerNotificationConfiguration                   string
	_sagemakerOfflineStoreConfig                          string
	_sagemakerOfflineStoreStatusEquals                    string
	_sagemakerOidcConfig                                  string
	_sagemakerOnCreate                                    string
	_sagemakerOnStart                                     string
	_sagemakerOnlineStoreConfig                           string
	_sagemakerOptimizationConfigs                         string
	_sagemakerOptimizationContains                        string
	_sagemakerOptimizationEnvironment                     string
	_sagemakerOptimizationJobName                         string
	_sagemakerOrchestrator                                string
	_sagemakerOutputArtifacts                             string
	_sagemakerOutputArtifactsToRemove                     []string
	_sagemakerOutputConfig                                string
	_sagemakerOutputDataConfig                            string
	_sagemakerOutputParameters                            string
	_sagemakerOwnershipSettings                           string
	_sagemakerParallelismConfiguration                    string
	_sagemakerParameterAdditions                          string
	_sagemakerParameterRemovals                           []string
	_sagemakerParameters                                  string
	_sagemakerParametersToRemove                          []string
	_sagemakerPipelineArn                                 string
	_sagemakerPipelineDefinition                          string
	_sagemakerPipelineDefinitionS3Location                string
	_sagemakerPipelineDescription                         string
	_sagemakerPipelineDisplayName                         string
	_sagemakerPipelineExecutionArn                        string
	_sagemakerPipelineExecutionDescription                string
	_sagemakerPipelineExecutionDisplayName                string
	_sagemakerPipelineName                                string
	_sagemakerPipelineNamePrefix                          string
	_sagemakerPipelineParameters                          string
	_sagemakerPipelineVersionDescription                  string
	_sagemakerPipelineVersionDisplayName                  string
	_sagemakerPipelineVersionId                           string
	_sagemakerPlatformIdentifier                          string
	_sagemakerPrimaryContainer                            string
	_sagemakerProblemType                                 string
	_sagemakerProcessingInputs                            string
	_sagemakerProcessingJobName                           string
	_sagemakerProcessingOutputConfig                      string
	_sagemakerProcessingResources                         string
	_sagemakerProcessor                                   string
	_sagemakerProductionVariants                          string
	_sagemakerProfilerConfig                              string
	_sagemakerProfilerRuleConfigurations                  string
	_sagemakerProgrammingLang                             string
	_sagemakerProjectDescription                          string
	_sagemakerProjectName                                 string
	_sagemakerProperties                                  string
	_sagemakerPropertiesToRemove                          []string
	_sagemakerReason                                      string
	_sagemakerRecommendationId                            string
	_sagemakerRecordIdentifierFeatureName                 string
	_sagemakerRecoveryMode                                string
	_sagemakerReleaseNotes                                string
	_sagemakerRemoteDebugConfig                           string
	_sagemakerReservedCapacityArn                         string
	_sagemakerResource                                    string
	_sagemakerResourceArn                                 string
	_sagemakerResourceConfig                              string
	_sagemakerResourceIdentifier                          string
	_sagemakerResourceKey                                 string
	_sagemakerResourcePolicy                              string
	_sagemakerResourceSpec                                string
	_sagemakerResourceType                                string
	_sagemakerRestrictedInstanceGroups                    string
	_sagemakerRetainAllVariantProperties                  string
	_sagemakerRetainDeploymentConfig                      string
	_sagemakerRetentionPolicy                             string
	_sagemakerRetryStrategy                               string
	_sagemakerRoleArn                                     string
	_sagemakerRootAccess                                  string
	_sagemakerRuntimeConfig                               string
	_sagemakerS3StorageConfig                             string
	_sagemakerSageMakerPublicHubContentArn                string
	_sagemakerSamplePayloadUrl                            string
	_sagemakerScalingPolicyObjective                      string
	_sagemakerSchedule                                    string
	_sagemakerScheduledTimeAfter                          string
	_sagemakerScheduledTimeBefore                         string
	_sagemakerSchedulerConfig                             string
	_sagemakerSearchExpression                            string
	_sagemakerSecurityConfig                              string
	_sagemakerSecurityGroupIds                            []string
	_sagemakerSelectiveExecutionConfig                    string
	_sagemakerServerlessJobConfig                         string
	_sagemakerServiceCatalogProvisioningDetails           string
	_sagemakerServiceCatalogProvisioningUpdateDetails     string
	_sagemakerSessionChainingConfig                       string
	_sagemakerSessionExpirationDurationInSeconds          string
	_sagemakerShadowModeConfig                            string
	_sagemakerShadowProductionVariants                    string
	_sagemakerSingleSignOnUserIdentifier                  string
	_sagemakerSingleSignOnUserValue                       string
	_sagemakerSkipModelValidation                         string
	_sagemakerSortBy                                      string
	_sagemakerSortOrder                                   string
	_sagemakerSource                                      string
	_sagemakerSourceAlgorithmSpecification                string
	_sagemakerSourceArn                                   string
	_sagemakerSourceIpConfig                              string
	_sagemakerSourceType                                  string
	_sagemakerSourceUri                                   string
	_sagemakerSpaceDisplayName                            string
	_sagemakerSpaceName                                   string
	_sagemakerSpaceNameContains                           string
	_sagemakerSpaceNameEquals                             string
	_sagemakerSpaceSettings                               string
	_sagemakerSpaceSharingSettings                        string
	_sagemakerSpareInstanceCountPerUltraServer            string
	_sagemakerSpecification                               string
	_sagemakerStageName                                   string
	_sagemakerStages                                      string
	_sagemakerStartArns                                   []string
	_sagemakerStartTime                                   string
	_sagemakerStartTimeAfter                              string
	_sagemakerStartTimeBefore                             string
	_sagemakerStatus                                      string
	_sagemakerStatusEquals                                string
	_sagemakerStepType                                    string
	_sagemakerStoppingCondition                           string
	_sagemakerStoppingConditions                          string
	_sagemakerStudioLifecycleConfigAppType                string
	_sagemakerStudioLifecycleConfigContent                string
	_sagemakerStudioLifecycleConfigName                   string
	_sagemakerSubnetId                                    string
	_sagemakerSubnetIds                                   []string
	_sagemakerSuggestionQuery                             string
	_sagemakerSupportStatus                               string
	_sagemakerTagKeys                                     []string
	_sagemakerTagPropagation                              string
	_sagemakerTags                                        string
	_sagemakerTargetCpuUtilizationPerCore                 string
	_sagemakerTargetResources                             string
	_sagemakerTargetVersion                               string
	_sagemakerTask                                        string
	_sagemakerTemplateProviders                           string
	_sagemakerTemplateProvidersToUpdate                   string
	_sagemakerTensorBoardOutputConfig                     string
	_sagemakerThroughputConfig                            string
	_sagemakerTier                                        string
	_sagemakerTieredStorageConfig                         string
	_sagemakerTrackingServerName                          string
	_sagemakerTrackingServerSize                          string
	_sagemakerTrackingServerStatus                        string
	_sagemakerTrainingJobDefinition                       string
	_sagemakerTrainingJobDefinitions                      string
	_sagemakerTrainingJobName                             string
	_sagemakerTrainingPlanArn                             string
	_sagemakerTrainingPlanArnEquals                       string
	_sagemakerTrainingPlanName                            string
	_sagemakerTrainingPlanOfferingId                      string
	_sagemakerTrainingSpecification                       string
	_sagemakerTransformInput                              string
	_sagemakerTransformJobName                            string
	_sagemakerTransformOutput                             string
	_sagemakerTransformResources                          string
	_sagemakerTrialComponentName                          string
	_sagemakerTrialName                                   string
	_sagemakerType                                        string
	_sagemakerUiTemplate                                  string
	_sagemakerUltraServerCount                            string
	_sagemakerUltraServerType                             string
	_sagemakerUserProfileName                             string
	_sagemakerUserProfileNameContains                     string
	_sagemakerUserProfileNameEquals                       string
	_sagemakerUserSettings                                string
	_sagemakerValidationSpecification                     string
	_sagemakerVariantName                                 string
	_sagemakerVariantNameEquals                           string
	_sagemakerVendorGuidance                              string
	_sagemakerVersion                                     string
	_sagemakerVisibilityConditions                        string
	_sagemakerVolumeId                                    string
	_sagemakerVolumeSizeInGB                              string
	_sagemakerVpcConfig                                   string
	_sagemakerVpcId                                       string
	_sagemakerWarmPoolStatusEquals                        string
	_sagemakerWarmStartConfig                             string
	_sagemakerWeeklyMaintenanceWindowStart                string
	_sagemakerWorkerAccessConfiguration                   string
	_sagemakerWorkforceName                               string
	_sagemakerWorkforceVpcConfig                          string
	_sagemakerWorkteamArn                                 string
	_sagemakerWorkteamName                                string
)

// Creates an association between the source and the destination. A source can be
// associated with multiple destinations, and a destination can be associated with
// multiple sources. An association is a lineage tracking entity. For more
// information, see [Amazon SageMaker ML Lineage Tracking].
//
// [Amazon SageMaker ML Lineage Tracking]: https://docs.aws.amazon.com/sagemaker/latest/dg/lineage-tracking.html
func sagemaker_AddAssociation(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.AddAssociationInput{
		// DestinationArn: *string, // Required
		// SourceArn: *string, // Required
	}

	if len(_sagemakerDestinationArn) > 0 {
		input.DestinationArn = aws.String(_sagemakerDestinationArn)
	}
	if len(_sagemakerSourceArn) > 0 {
		input.SourceArn = aws.String(_sagemakerSourceArn)
	}
	if len(_sagemakerAssociationType) > 0 {
		if err := assignInputField(input, "AssociationType", _sagemakerAssociationType); err != nil {
			log.Errorf("invalid --association-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.AddAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds or overwrites one or more tags for the specified SageMaker resource. You
// can add tags to notebook instances, training jobs, hyperparameter tuning jobs,
// batch transform jobs, models, labeling jobs, work teams, endpoint
// configurations, and endpoints.
//
// Each tag consists of a key and an optional value. Tag keys must be unique per
// resource. For more information about tags, see For more information, see [Amazon Web Services Tagging Strategies].
//
// Tags that you add to a hyperparameter tuning job by calling this API are also
// added to any training jobs that the hyperparameter tuning job launches after you
// call this API, but not to training jobs that the hyperparameter tuning job
// launched before you called this API. To make sure that the tags associated with
// a hyperparameter tuning job are also added to all training jobs that the
// hyperparameter tuning job launches, add the tags when you first create the
// tuning job by specifying them in the Tags parameter of [CreateHyperParameterTuningJob]
//
// Tags that you add to a SageMaker Domain or User Profile by calling this API are
// also added to any Apps that the Domain or User Profile launches after you call
// this API, but not to Apps that the Domain or User Profile launched before you
// called this API. To make sure that the tags associated with a Domain or User
// Profile are also added to all Apps that the Domain or User Profile launches, add
// the tags when you first create the Domain or User Profile by specifying them in
// the Tags parameter of [CreateDomain] or [CreateUserProfile].
//
// [CreateHyperParameterTuningJob]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateHyperParameterTuningJob.html
// [Amazon Web Services Tagging Strategies]: https://aws.amazon.com/answers/account-management/aws-tagging-strategies/
// [CreateUserProfile]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateUserProfile.html
// [CreateDomain]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateDomain.html
func sagemaker_AddTags(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.AddTagsInput{
		// ResourceArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_sagemakerResourceArn) > 0 {
		input.ResourceArn = aws.String(_sagemakerResourceArn)
	}
	if len(_sagemakerTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.AddTags(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates a trial component with a trial. A trial component can be associated
// with multiple trials. To disassociate a trial component from a trial, call the [DisassociateTrialComponent]
// API.
//
// [DisassociateTrialComponent]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DisassociateTrialComponent.html
func sagemaker_AssociateTrialComponent(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.AssociateTrialComponentInput{
		// TrialComponentName: *string, // Required
		// TrialName: *string, // Required
	}

	if len(_sagemakerTrialComponentName) > 0 {
		input.TrialComponentName = aws.String(_sagemakerTrialComponentName)
	}
	if len(_sagemakerTrialName) > 0 {
		input.TrialName = aws.String(_sagemakerTrialName)
	}

	if resp, err := client.AssociateTrialComponent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Attaches your Amazon Elastic Block Store (Amazon EBS) volume to a node in your
// EKS orchestrated HyperPod cluster.
//
// This API works with the Amazon Elastic Block Store (Amazon EBS) Container
// Storage Interface (CSI) driver to manage the lifecycle of persistent storage in
// your HyperPod EKS clusters.
func sagemaker_AttachClusterNodeVolume(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.AttachClusterNodeVolumeInput{
		// ClusterArn: *string, // Required
		// NodeId: *string, // Required
		// VolumeId: *string, // Required
	}

	if len(_sagemakerClusterArn) > 0 {
		input.ClusterArn = aws.String(_sagemakerClusterArn)
	}
	if len(_sagemakerNodeId) > 0 {
		input.NodeId = aws.String(_sagemakerNodeId)
	}
	if len(_sagemakerVolumeId) > 0 {
		input.VolumeId = aws.String(_sagemakerVolumeId)
	}

	if resp, err := client.AttachClusterNodeVolume(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds nodes to a HyperPod cluster by incrementing the target count for one or
// more instance groups. This operation returns a unique NodeLogicalId for each
// node being added, which can be used to track the provisioning status of the
// node. This API provides a safer alternative to UpdateCluster for scaling
// operations by avoiding unintended configuration changes.
//
// This API is only supported for clusters using Continuous as the
// NodeProvisioningMode .
func sagemaker_BatchAddClusterNodes(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.BatchAddClusterNodesInput{
		// ClusterName: *string, // Required
		// NodesToAdd: []types.AddClusterNodeSpecification, // Required
	}

	if len(_sagemakerClusterName) > 0 {
		input.ClusterName = aws.String(_sagemakerClusterName)
	}
	if len(_sagemakerNodesToAdd) > 0 {
		if err := assignInputField(input, "NodesToAdd", _sagemakerNodesToAdd); err != nil {
			log.Errorf("invalid --nodes-to-add: %s", err.Error())
			return
		}
	}
	if len(_sagemakerClientToken) > 0 {
		input.ClientToken = aws.String(_sagemakerClientToken)
	}

	if resp, err := client.BatchAddClusterNodes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes specific nodes within a SageMaker HyperPod cluster.
// BatchDeleteClusterNodes accepts a cluster name and a list of node IDs.
//
// - To safeguard your work, back up your data to Amazon S3 or an FSx for Lustre
// file system before invoking the API on a worker node group. This will help
// prevent any potential data loss from the instance root volume. For more
// information about backup, see [Use the backup script provided by SageMaker HyperPod].
//
// - If you want to invoke this API on an existing cluster, you'll first need to
// patch the cluster by running the [UpdateClusterSoftware API]. For more information about patching a
// cluster, see [Update the SageMaker HyperPod platform software of a cluster].
//
// [UpdateClusterSoftware API]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_UpdateClusterSoftware.html
// [Use the backup script provided by SageMaker HyperPod]: https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-hyperpod-operate-cli-command.html#sagemaker-hyperpod-operate-cli-command-update-cluster-software-backup
// [Update the SageMaker HyperPod platform software of a cluster]: https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-hyperpod-operate-cli-command.html#sagemaker-hyperpod-operate-cli-command-update-cluster-software
func sagemaker_BatchDeleteClusterNodes(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.BatchDeleteClusterNodesInput{
		// ClusterName: *string, // Required
	}

	if len(_sagemakerClusterName) > 0 {
		input.ClusterName = aws.String(_sagemakerClusterName)
	}
	if len(_sagemakerNodeIds) > 0 {
		input.NodeIds = append([]string(nil), _sagemakerNodeIds...)
	}
	if len(_sagemakerNodeLogicalIds) > 0 {
		input.NodeLogicalIds = append([]string(nil), _sagemakerNodeLogicalIds...)
	}

	if resp, err := client.BatchDeleteClusterNodes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This action batch describes a list of versioned model packages
func sagemaker_BatchDescribeModelPackage(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.BatchDescribeModelPackageInput{
		// ModelPackageArnList: []string, // Required
	}

	if len(_sagemakerModelPackageArnList) > 0 {
		input.ModelPackageArnList = append([]string(nil), _sagemakerModelPackageArnList...)
	}

	if resp, err := client.BatchDescribeModelPackage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Reboots specific nodes within a SageMaker HyperPod cluster using a soft
// recovery mechanism. BatchRebootClusterNodes performs a graceful reboot of the
// specified nodes by calling the Amazon Elastic Compute Cloud RebootInstances
// API, which attempts to cleanly shut down the operating system before restarting
// the instance.
//
// This operation is useful for recovering from transient issues or applying
// certain configuration changes that require a restart.
//
// - Rebooting a node may cause temporary service interruption for workloads
// running on that node. Ensure your workloads can handle node restarts or use
// appropriate scheduling to minimize impact.
//
// - You can reboot up to 25 nodes in a single request.
//
// - For SageMaker HyperPod clusters using the Slurm workload manager, ensure
// rebooting nodes will not disrupt critical cluster operations.
func sagemaker_BatchRebootClusterNodes(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.BatchRebootClusterNodesInput{
		// ClusterName: *string, // Required
	}

	if len(_sagemakerClusterName) > 0 {
		input.ClusterName = aws.String(_sagemakerClusterName)
	}
	if len(_sagemakerNodeIds) > 0 {
		input.NodeIds = append([]string(nil), _sagemakerNodeIds...)
	}
	if len(_sagemakerNodeLogicalIds) > 0 {
		input.NodeLogicalIds = append([]string(nil), _sagemakerNodeLogicalIds...)
	}

	if resp, err := client.BatchRebootClusterNodes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Replaces specific nodes within a SageMaker HyperPod cluster with new hardware.
// BatchReplaceClusterNodes terminates the specified instances and provisions new
// replacement instances with the same configuration but fresh hardware. The Amazon
// Machine Image (AMI) and instance configuration remain the same.
//
// This operation is useful for recovering from hardware failures or persistent
// issues that cannot be resolved through a reboot.
//
// - Data Loss Warning: Replacing nodes destroys all instance volumes, including
// both root and secondary volumes. All data stored on these volumes will be
// permanently lost and cannot be recovered.
//
// - To safeguard your work, back up your data to Amazon S3 or an FSx for Lustre
// file system before invoking the API on a worker node group. This will help
// prevent any potential data loss from the instance root volume. For more
// information about backup, see [Use the backup script provided by SageMaker HyperPod].
//
// - If you want to invoke this API on an existing cluster, you'll first need to
// patch the cluster by running the [UpdateClusterSoftware API]. For more information about patching a
// cluster, see [Update the SageMaker HyperPod platform software of a cluster].
//
// - You can replace up to 25 nodes in a single request.
//
// [UpdateClusterSoftware API]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_UpdateClusterSoftware.html
// [Use the backup script provided by SageMaker HyperPod]: https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-hyperpod-operate-cli-command.html#sagemaker-hyperpod-operate-cli-command-update-cluster-software-backup
// [Update the SageMaker HyperPod platform software of a cluster]: https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-hyperpod-operate-cli-command.html#sagemaker-hyperpod-operate-cli-command-update-cluster-software
func sagemaker_BatchReplaceClusterNodes(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.BatchReplaceClusterNodesInput{
		// ClusterName: *string, // Required
	}

	if len(_sagemakerClusterName) > 0 {
		input.ClusterName = aws.String(_sagemakerClusterName)
	}
	if len(_sagemakerNodeIds) > 0 {
		input.NodeIds = append([]string(nil), _sagemakerNodeIds...)
	}
	if len(_sagemakerNodeLogicalIds) > 0 {
		input.NodeLogicalIds = append([]string(nil), _sagemakerNodeLogicalIds...)
	}

	if resp, err := client.BatchReplaceClusterNodes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an action. An action is a lineage tracking entity that represents an
// action or activity. For example, a model deployment or an HPO job. Generally, an
// action involves at least one input or output artifact. For more information, see
// [Amazon SageMaker ML Lineage Tracking].
//
// [Amazon SageMaker ML Lineage Tracking]: https://docs.aws.amazon.com/sagemaker/latest/dg/lineage-tracking.html
func sagemaker_CreateAction(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreateActionInput{
		// ActionName: *string, // Required
		// ActionType: *string, // Required
		// Source: *types.ActionSource, // Required
	}

	if len(_sagemakerActionName) > 0 {
		input.ActionName = aws.String(_sagemakerActionName)
	}
	if len(_sagemakerActionType) > 0 {
		input.ActionType = aws.String(_sagemakerActionType)
	}
	if len(_sagemakerSource) > 0 {
		if err := assignInputField(input, "Source", _sagemakerSource); err != nil {
			log.Errorf("invalid --source: %s", err.Error())
			return
		}
	}
	if len(_sagemakerDescription) > 0 {
		input.Description = aws.String(_sagemakerDescription)
	}
	if len(_sagemakerMetadataProperties) > 0 {
		if err := assignInputField(input, "MetadataProperties", _sagemakerMetadataProperties); err != nil {
			log.Errorf("invalid --metadata-properties: %s", err.Error())
			return
		}
	}
	if len(_sagemakerProperties) > 0 {
		if err := assignInputField(input, "Properties", _sagemakerProperties); err != nil {
			log.Errorf("invalid --properties: %s", err.Error())
			return
		}
	}
	if len(_sagemakerStatus) > 0 {
		if err := assignInputField(input, "Status", _sagemakerStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_sagemakerTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a machine learning algorithm that you can use in SageMaker and list in
// the Amazon Web Services Marketplace.
func sagemaker_CreateAlgorithm(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreateAlgorithmInput{
		// AlgorithmName: *string, // Required
		// TrainingSpecification: *types.TrainingSpecification, // Required
	}

	if len(_sagemakerAlgorithmName) > 0 {
		input.AlgorithmName = aws.String(_sagemakerAlgorithmName)
	}
	if len(_sagemakerTrainingSpecification) > 0 {
		if err := assignInputField(input, "TrainingSpecification", _sagemakerTrainingSpecification); err != nil {
			log.Errorf("invalid --training-specification: %s", err.Error())
			return
		}
	}
	if len(_sagemakerAlgorithmDescription) > 0 {
		input.AlgorithmDescription = aws.String(_sagemakerAlgorithmDescription)
	}
	if len(_sagemakerCertifyForMarketplace) > 0 {
		if err := assignInputField(input, "CertifyForMarketplace", _sagemakerCertifyForMarketplace); err != nil {
			log.Errorf("invalid --certify-for-marketplace: %s", err.Error())
			return
		}
	}
	if len(_sagemakerInferenceSpecification) > 0 {
		if err := assignInputField(input, "InferenceSpecification", _sagemakerInferenceSpecification); err != nil {
			log.Errorf("invalid --inference-specification: %s", err.Error())
			return
		}
	}
	if len(_sagemakerTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_sagemakerValidationSpecification) > 0 {
		if err := assignInputField(input, "ValidationSpecification", _sagemakerValidationSpecification); err != nil {
			log.Errorf("invalid --validation-specification: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAlgorithm(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a running app for the specified UserProfile. This operation is
// automatically invoked by Amazon SageMaker AI upon access to the associated
// Domain, and when new kernel configurations are selected by the user. A user may
// have multiple Apps active simultaneously.
func sagemaker_CreateApp(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreateAppInput{
		// AppName: *string, // Required
		// AppType: types.AppType, // Required
		// DomainId: *string, // Required
	}

	if len(_sagemakerAppName) > 0 {
		input.AppName = aws.String(_sagemakerAppName)
	}
	if len(_sagemakerAppType) > 0 {
		if err := assignInputField(input, "AppType", _sagemakerAppType); err != nil {
			log.Errorf("invalid --app-type: %s", err.Error())
			return
		}
	}
	if len(_sagemakerDomainId) > 0 {
		input.DomainId = aws.String(_sagemakerDomainId)
	}
	if len(_sagemakerRecoveryMode) > 0 {
		if err := assignInputField(input, "RecoveryMode", _sagemakerRecoveryMode); err != nil {
			log.Errorf("invalid --recovery-mode: %s", err.Error())
			return
		}
	}
	if len(_sagemakerResourceSpec) > 0 {
		if err := assignInputField(input, "ResourceSpec", _sagemakerResourceSpec); err != nil {
			log.Errorf("invalid --resource-spec: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSpaceName) > 0 {
		input.SpaceName = aws.String(_sagemakerSpaceName)
	}
	if len(_sagemakerTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_sagemakerUserProfileName) > 0 {
		input.UserProfileName = aws.String(_sagemakerUserProfileName)
	}

	if resp, err := client.CreateApp(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a configuration for running a SageMaker AI image as a KernelGateway
// app. The configuration specifies the Amazon Elastic File System storage volume
// on the image, and a list of the kernels in the image.
func sagemaker_CreateAppImageConfig(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreateAppImageConfigInput{
		// AppImageConfigName: *string, // Required
	}

	if len(_sagemakerAppImageConfigName) > 0 {
		input.AppImageConfigName = aws.String(_sagemakerAppImageConfigName)
	}
	if len(_sagemakerCodeEditorAppImageConfig) > 0 {
		if err := assignInputField(input, "CodeEditorAppImageConfig", _sagemakerCodeEditorAppImageConfig); err != nil {
			log.Errorf("invalid --code-editor-app-image-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerJupyterLabAppImageConfig) > 0 {
		if err := assignInputField(input, "JupyterLabAppImageConfig", _sagemakerJupyterLabAppImageConfig); err != nil {
			log.Errorf("invalid --jupyter-lab-app-image-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerKernelGatewayImageConfig) > 0 {
		if err := assignInputField(input, "KernelGatewayImageConfig", _sagemakerKernelGatewayImageConfig); err != nil {
			log.Errorf("invalid --kernel-gateway-image-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAppImageConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an artifact. An artifact is a lineage tracking entity that represents a
// URI addressable object or data. Some examples are the S3 URI of a dataset and
// the ECR registry path of an image. For more information, see [Amazon SageMaker ML Lineage Tracking].
//
// [Amazon SageMaker ML Lineage Tracking]: https://docs.aws.amazon.com/sagemaker/latest/dg/lineage-tracking.html
func sagemaker_CreateArtifact(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreateArtifactInput{
		// ArtifactType: *string, // Required
		// Source: *types.ArtifactSource, // Required
	}

	if len(_sagemakerArtifactType) > 0 {
		input.ArtifactType = aws.String(_sagemakerArtifactType)
	}
	if len(_sagemakerSource) > 0 {
		if err := assignInputField(input, "Source", _sagemakerSource); err != nil {
			log.Errorf("invalid --source: %s", err.Error())
			return
		}
	}
	if len(_sagemakerArtifactName) > 0 {
		input.ArtifactName = aws.String(_sagemakerArtifactName)
	}
	if len(_sagemakerMetadataProperties) > 0 {
		if err := assignInputField(input, "MetadataProperties", _sagemakerMetadataProperties); err != nil {
			log.Errorf("invalid --metadata-properties: %s", err.Error())
			return
		}
	}
	if len(_sagemakerProperties) > 0 {
		if err := assignInputField(input, "Properties", _sagemakerProperties); err != nil {
			log.Errorf("invalid --properties: %s", err.Error())
			return
		}
	}
	if len(_sagemakerTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateArtifact(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Autopilot job also referred to as Autopilot experiment or AutoML job.
// An AutoML job in SageMaker AI is a fully automated process that allows you to
// build machine learning models with minimal effort and machine learning
// expertise. When initiating an AutoML job, you provide your data and optionally
// specify parameters tailored to your use case. SageMaker AI then automates the
// entire model development lifecycle, including data preprocessing, model
// training, tuning, and evaluation. AutoML jobs are designed to simplify and
// accelerate the model building process by automating various tasks and exploring
// different combinations of machine learning algorithms, data preprocessing
// techniques, and hyperparameter values. The output of an AutoML job comprises one
// or more trained models ready for deployment and inference. Additionally,
// SageMaker AI AutoML jobs generate a candidate model leaderboard, allowing you to
// select the best-performing model for deployment.
//
// For more information about AutoML jobs, see [https://docs.aws.amazon.com/sagemaker/latest/dg/autopilot-automate-model-development.html] in the SageMaker AI developer
// guide.
//
// We recommend using the new versions [CreateAutoMLJobV2] and [DescribeAutoMLJobV2], which offer backward compatibility.
//
// CreateAutoMLJobV2 can manage tabular problem types identical to those of its
// previous version CreateAutoMLJob , as well as time-series forecasting,
// non-tabular problem types such as image or text classification, and text
// generation (LLMs fine-tuning).
//
// Find guidelines about how to migrate a CreateAutoMLJob to CreateAutoMLJobV2 in [Migrate a CreateAutoMLJob to CreateAutoMLJobV2].
//
// You can find the best-performing model after you run an AutoML job by calling [DescribeAutoMLJobV2]
// (recommended) or [DescribeAutoMLJob].
//
// [DescribeAutoMLJob]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DescribeAutoMLJob.html
// [DescribeAutoMLJobV2]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DescribeAutoMLJobV2.html
// [https://docs.aws.amazon.com/sagemaker/latest/dg/autopilot-automate-model-development.html]: https://docs.aws.amazon.com/sagemaker/latest/dg/autopilot-automate-model-development.html
// [CreateAutoMLJobV2]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateAutoMLJobV2.html
// [Migrate a CreateAutoMLJob to CreateAutoMLJobV2]: https://docs.aws.amazon.com/sagemaker/latest/dg/autopilot-automate-model-development-create-experiment.html#autopilot-create-experiment-api-migrate-v1-v2
func sagemaker_CreateAutoMLJob(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreateAutoMLJobInput{
		// AutoMLJobName: *string, // Required
		// InputDataConfig: []types.AutoMLChannel, // Required
		// OutputDataConfig: *types.AutoMLOutputDataConfig, // Required
		// RoleArn: *string, // Required
	}

	if len(_sagemakerAutoMLJobName) > 0 {
		input.AutoMLJobName = aws.String(_sagemakerAutoMLJobName)
	}
	if len(_sagemakerInputDataConfig) > 0 {
		if err := assignInputField(input, "InputDataConfig", _sagemakerInputDataConfig); err != nil {
			log.Errorf("invalid --input-data-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerOutputDataConfig) > 0 {
		if err := assignInputField(input, "OutputDataConfig", _sagemakerOutputDataConfig); err != nil {
			log.Errorf("invalid --output-data-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerRoleArn) > 0 {
		input.RoleArn = aws.String(_sagemakerRoleArn)
	}
	if len(_sagemakerAutoMLJobConfig) > 0 {
		if err := assignInputField(input, "AutoMLJobConfig", _sagemakerAutoMLJobConfig); err != nil {
			log.Errorf("invalid --auto-ml-job-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerAutoMLJobObjective) > 0 {
		if err := assignInputField(input, "AutoMLJobObjective", _sagemakerAutoMLJobObjective); err != nil {
			log.Errorf("invalid --auto-ml-job-objective: %s", err.Error())
			return
		}
	}
	if len(_sagemakerGenerateCandidateDefinitionsOnly) > 0 {
		if err := assignInputField(input, "GenerateCandidateDefinitionsOnly", _sagemakerGenerateCandidateDefinitionsOnly); err != nil {
			log.Errorf("invalid --generate-candidate-definitions-only: %s", err.Error())
			return
		}
	}
	if len(_sagemakerModelDeployConfig) > 0 {
		if err := assignInputField(input, "ModelDeployConfig", _sagemakerModelDeployConfig); err != nil {
			log.Errorf("invalid --model-deploy-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerProblemType) > 0 {
		if err := assignInputField(input, "ProblemType", _sagemakerProblemType); err != nil {
			log.Errorf("invalid --problem-type: %s", err.Error())
			return
		}
	}
	if len(_sagemakerTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAutoMLJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Autopilot job also referred to as Autopilot experiment or AutoML job
// V2.
//
// An AutoML job in SageMaker AI is a fully automated process that allows you to
// build machine learning models with minimal effort and machine learning
// expertise. When initiating an AutoML job, you provide your data and optionally
// specify parameters tailored to your use case. SageMaker AI then automates the
// entire model development lifecycle, including data preprocessing, model
// training, tuning, and evaluation. AutoML jobs are designed to simplify and
// accelerate the model building process by automating various tasks and exploring
// different combinations of machine learning algorithms, data preprocessing
// techniques, and hyperparameter values. The output of an AutoML job comprises one
// or more trained models ready for deployment and inference. Additionally,
// SageMaker AI AutoML jobs generate a candidate model leaderboard, allowing you to
// select the best-performing model for deployment.
//
// For more information about AutoML jobs, see [https://docs.aws.amazon.com/sagemaker/latest/dg/autopilot-automate-model-development.html] in the SageMaker AI developer
// guide.
//
// AutoML jobs V2 support various problem types such as regression, binary, and
// multiclass classification with tabular data, text and image classification,
// time-series forecasting, and fine-tuning of large language models (LLMs) for
// text generation.
//
// [CreateAutoMLJobV2]and [DescribeAutoMLJobV2] are new versions of [CreateAutoMLJob] and [DescribeAutoMLJob] which offer backward compatibility.
//
// CreateAutoMLJobV2 can manage tabular problem types identical to those of its
// previous version CreateAutoMLJob , as well as time-series forecasting,
// non-tabular problem types such as image or text classification, and text
// generation (LLMs fine-tuning).
//
// Find guidelines about how to migrate a CreateAutoMLJob to CreateAutoMLJobV2 in [Migrate a CreateAutoMLJob to CreateAutoMLJobV2].
//
// For the list of available problem types supported by CreateAutoMLJobV2 , see [AutoMLProblemTypeConfig].
//
// You can find the best-performing model after you run an AutoML job V2 by
// calling [DescribeAutoMLJobV2].
//
// [CreateAutoMLJob]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateAutoMLJob.html
// [DescribeAutoMLJob]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DescribeAutoMLJob.html
// [DescribeAutoMLJobV2]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DescribeAutoMLJobV2.html
// [https://docs.aws.amazon.com/sagemaker/latest/dg/autopilot-automate-model-development.html]: https://docs.aws.amazon.com/sagemaker/latest/dg/autopilot-automate-model-development.html
// [CreateAutoMLJobV2]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateAutoMLJobV2.html
// [Migrate a CreateAutoMLJob to CreateAutoMLJobV2]: https://docs.aws.amazon.com/sagemaker/latest/dg/autopilot-automate-model-development-create-experiment.html#autopilot-create-experiment-api-migrate-v1-v2
// [AutoMLProblemTypeConfig]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_AutoMLProblemTypeConfig.html
func sagemaker_CreateAutoMLJobV2(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreateAutoMLJobV2Input{
		// AutoMLJobInputDataConfig: []types.AutoMLJobChannel, // Required
		// AutoMLJobName: *string, // Required
		// AutoMLProblemTypeConfig: types.AutoMLProblemTypeConfig, // Required
		// OutputDataConfig: *types.AutoMLOutputDataConfig, // Required
		// RoleArn: *string, // Required
	}

	if len(_sagemakerAutoMLJobInputDataConfig) > 0 {
		if err := assignInputField(input, "AutoMLJobInputDataConfig", _sagemakerAutoMLJobInputDataConfig); err != nil {
			log.Errorf("invalid --auto-ml-job-input-data-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerAutoMLJobName) > 0 {
		input.AutoMLJobName = aws.String(_sagemakerAutoMLJobName)
	}
	if len(_sagemakerAutoMLProblemTypeConfig) > 0 {
		if err := assignInputField(input, "AutoMLProblemTypeConfig", _sagemakerAutoMLProblemTypeConfig); err != nil {
			log.Errorf("invalid --auto-ml-problem-type-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerOutputDataConfig) > 0 {
		if err := assignInputField(input, "OutputDataConfig", _sagemakerOutputDataConfig); err != nil {
			log.Errorf("invalid --output-data-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerRoleArn) > 0 {
		input.RoleArn = aws.String(_sagemakerRoleArn)
	}
	if len(_sagemakerAutoMLComputeConfig) > 0 {
		if err := assignInputField(input, "AutoMLComputeConfig", _sagemakerAutoMLComputeConfig); err != nil {
			log.Errorf("invalid --auto-ml-compute-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerAutoMLJobObjective) > 0 {
		if err := assignInputField(input, "AutoMLJobObjective", _sagemakerAutoMLJobObjective); err != nil {
			log.Errorf("invalid --auto-ml-job-objective: %s", err.Error())
			return
		}
	}
	if len(_sagemakerDataSplitConfig) > 0 {
		if err := assignInputField(input, "DataSplitConfig", _sagemakerDataSplitConfig); err != nil {
			log.Errorf("invalid --data-split-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerModelDeployConfig) > 0 {
		if err := assignInputField(input, "ModelDeployConfig", _sagemakerModelDeployConfig); err != nil {
			log.Errorf("invalid --model-deploy-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSecurityConfig) > 0 {
		if err := assignInputField(input, "SecurityConfig", _sagemakerSecurityConfig); err != nil {
			log.Errorf("invalid --security-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAutoMLJobV2(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon SageMaker HyperPod cluster. SageMaker HyperPod is a
// capability of SageMaker for creating and managing persistent clusters for
// developing large machine learning models, such as large language models (LLMs)
// and diffusion models. To learn more, see [Amazon SageMaker HyperPod]in the Amazon SageMaker Developer
// Guide.
//
// [Amazon SageMaker HyperPod]: https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-hyperpod.html
func sagemaker_CreateCluster(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreateClusterInput{
		// ClusterName: *string, // Required
	}

	if len(_sagemakerClusterName) > 0 {
		input.ClusterName = aws.String(_sagemakerClusterName)
	}
	if len(_sagemakerAutoScaling) > 0 {
		if err := assignInputField(input, "AutoScaling", _sagemakerAutoScaling); err != nil {
			log.Errorf("invalid --auto-scaling: %s", err.Error())
			return
		}
	}
	if len(_sagemakerClusterRole) > 0 {
		input.ClusterRole = aws.String(_sagemakerClusterRole)
	}
	if len(_sagemakerInstanceGroups) > 0 {
		if err := assignInputField(input, "InstanceGroups", _sagemakerInstanceGroups); err != nil {
			log.Errorf("invalid --instance-groups: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNodeProvisioningMode) > 0 {
		if err := assignInputField(input, "NodeProvisioningMode", _sagemakerNodeProvisioningMode); err != nil {
			log.Errorf("invalid --node-provisioning-mode: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNodeRecovery) > 0 {
		if err := assignInputField(input, "NodeRecovery", _sagemakerNodeRecovery); err != nil {
			log.Errorf("invalid --node-recovery: %s", err.Error())
			return
		}
	}
	if len(_sagemakerOrchestrator) > 0 {
		if err := assignInputField(input, "Orchestrator", _sagemakerOrchestrator); err != nil {
			log.Errorf("invalid --orchestrator: %s", err.Error())
			return
		}
	}
	if len(_sagemakerRestrictedInstanceGroups) > 0 {
		if err := assignInputField(input, "RestrictedInstanceGroups", _sagemakerRestrictedInstanceGroups); err != nil {
			log.Errorf("invalid --restricted-instance-groups: %s", err.Error())
			return
		}
	}
	if len(_sagemakerTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_sagemakerTieredStorageConfig) > 0 {
		if err := assignInputField(input, "TieredStorageConfig", _sagemakerTieredStorageConfig); err != nil {
			log.Errorf("invalid --tiered-storage-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerVpcConfig) > 0 {
		if err := assignInputField(input, "VpcConfig", _sagemakerVpcConfig); err != nil {
			log.Errorf("invalid --vpc-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create cluster policy configuration. This policy is used for task
// prioritization and fair-share allocation of idle compute. This helps prioritize
// critical workloads and distributes idle compute across entities.
func sagemaker_CreateClusterSchedulerConfig(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreateClusterSchedulerConfigInput{
		// ClusterArn: *string, // Required
		// Name: *string, // Required
		// SchedulerConfig: *types.SchedulerConfig, // Required
	}

	if len(_sagemakerClusterArn) > 0 {
		input.ClusterArn = aws.String(_sagemakerClusterArn)
	}
	if len(_sagemakerName) > 0 {
		input.Name = aws.String(_sagemakerName)
	}
	if len(_sagemakerSchedulerConfig) > 0 {
		if err := assignInputField(input, "SchedulerConfig", _sagemakerSchedulerConfig); err != nil {
			log.Errorf("invalid --scheduler-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerDescription) > 0 {
		input.Description = aws.String(_sagemakerDescription)
	}
	if len(_sagemakerTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateClusterSchedulerConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a Git repository as a resource in your SageMaker AI account. You can
// associate the repository with notebook instances so that you can use Git source
// control for the notebooks you create. The Git repository is a resource in your
// SageMaker AI account, so it can be associated with more than one notebook
// instance, and it persists independently from the lifecycle of any notebook
// instances it is associated with.
//
// The repository can be hosted either in [Amazon Web Services CodeCommit] or in any other Git repository.
//
// [Amazon Web Services CodeCommit]: https://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html
func sagemaker_CreateCodeRepository(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreateCodeRepositoryInput{
		// CodeRepositoryName: *string, // Required
		// GitConfig: *types.GitConfig, // Required
	}

	if len(_sagemakerCodeRepositoryName) > 0 {
		input.CodeRepositoryName = aws.String(_sagemakerCodeRepositoryName)
	}
	if len(_sagemakerGitConfig) > 0 {
		if err := assignInputField(input, "GitConfig", _sagemakerGitConfig); err != nil {
			log.Errorf("invalid --git-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCodeRepository(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a model compilation job. After the model has been compiled, Amazon
// SageMaker AI saves the resulting model artifacts to an Amazon Simple Storage
// Service (Amazon S3) bucket that you specify.
//
// If you choose to host your model using Amazon SageMaker AI hosting services,
// you can use the resulting model artifacts as part of the model. You can also use
// the artifacts with Amazon Web Services IoT Greengrass. In that case, deploy them
// as an ML resource.
//
// In the request body, you provide the following:
//
// - A name for the compilation job
//
// - Information about the input model artifacts
//
// - The output location for the compiled model and the device (target) that the
// model runs on
//
// - The Amazon Resource Name (ARN) of the IAM role that Amazon SageMaker AI
// assumes to perform the model compilation job.
//
// You can also provide a Tag to track the model compilation job's resource use
// and costs. The response body contains the CompilationJobArn for the compiled
// job.
//
// To stop a model compilation job, use [StopCompilationJob]. To get information about a particular
// model compilation job, use [DescribeCompilationJob]. To get information about multiple model
// compilation jobs, use [ListCompilationJobs].
//
// [StopCompilationJob]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_StopCompilationJob.html
// [DescribeCompilationJob]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DescribeCompilationJob.html
// [ListCompilationJobs]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_ListCompilationJobs.html
func sagemaker_CreateCompilationJob(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreateCompilationJobInput{
		// CompilationJobName: *string, // Required
		// OutputConfig: *types.OutputConfig, // Required
		// RoleArn: *string, // Required
		// StoppingCondition: *types.StoppingCondition, // Required
	}

	if len(_sagemakerCompilationJobName) > 0 {
		input.CompilationJobName = aws.String(_sagemakerCompilationJobName)
	}
	if len(_sagemakerOutputConfig) > 0 {
		if err := assignInputField(input, "OutputConfig", _sagemakerOutputConfig); err != nil {
			log.Errorf("invalid --output-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerRoleArn) > 0 {
		input.RoleArn = aws.String(_sagemakerRoleArn)
	}
	if len(_sagemakerStoppingCondition) > 0 {
		if err := assignInputField(input, "StoppingCondition", _sagemakerStoppingCondition); err != nil {
			log.Errorf("invalid --stopping-condition: %s", err.Error())
			return
		}
	}
	if len(_sagemakerInputConfig) > 0 {
		if err := assignInputField(input, "InputConfig", _sagemakerInputConfig); err != nil {
			log.Errorf("invalid --input-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerModelPackageVersionArn) > 0 {
		input.ModelPackageVersionArn = aws.String(_sagemakerModelPackageVersionArn)
	}
	if len(_sagemakerTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_sagemakerVpcConfig) > 0 {
		if err := assignInputField(input, "VpcConfig", _sagemakerVpcConfig); err != nil {
			log.Errorf("invalid --vpc-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCompilationJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create compute allocation definition. This defines how compute is allocated,
// shared, and borrowed for specified entities. Specifically, how to lend and
// borrow idle compute and assign a fair-share weight to the specified entities.
func sagemaker_CreateComputeQuota(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreateComputeQuotaInput{
		// ClusterArn: *string, // Required
		// ComputeQuotaConfig: *types.ComputeQuotaConfig, // Required
		// ComputeQuotaTarget: *types.ComputeQuotaTarget, // Required
		// Name: *string, // Required
	}

	if len(_sagemakerClusterArn) > 0 {
		input.ClusterArn = aws.String(_sagemakerClusterArn)
	}
	if len(_sagemakerComputeQuotaConfig) > 0 {
		if err := assignInputField(input, "ComputeQuotaConfig", _sagemakerComputeQuotaConfig); err != nil {
			log.Errorf("invalid --compute-quota-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerComputeQuotaTarget) > 0 {
		if err := assignInputField(input, "ComputeQuotaTarget", _sagemakerComputeQuotaTarget); err != nil {
			log.Errorf("invalid --compute-quota-target: %s", err.Error())
			return
		}
	}
	if len(_sagemakerName) > 0 {
		input.Name = aws.String(_sagemakerName)
	}
	if len(_sagemakerActivationState) > 0 {
		if err := assignInputField(input, "ActivationState", _sagemakerActivationState); err != nil {
			log.Errorf("invalid --activation-state: %s", err.Error())
			return
		}
	}
	if len(_sagemakerDescription) > 0 {
		input.Description = aws.String(_sagemakerDescription)
	}
	if len(_sagemakerTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateComputeQuota(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a context. A context is a lineage tracking entity that represents a
// logical grouping of other tracking or experiment entities. Some examples are an
// endpoint and a model package. For more information, see [Amazon SageMaker ML Lineage Tracking].
//
// [Amazon SageMaker ML Lineage Tracking]: https://docs.aws.amazon.com/sagemaker/latest/dg/lineage-tracking.html
func sagemaker_CreateContext(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreateContextInput{
		// ContextName: *string, // Required
		// ContextType: *string, // Required
		// Source: *types.ContextSource, // Required
	}

	if len(_sagemakerContextName) > 0 {
		input.ContextName = aws.String(_sagemakerContextName)
	}
	if len(_sagemakerContextType) > 0 {
		input.ContextType = aws.String(_sagemakerContextType)
	}
	if len(_sagemakerSource) > 0 {
		if err := assignInputField(input, "Source", _sagemakerSource); err != nil {
			log.Errorf("invalid --source: %s", err.Error())
			return
		}
	}
	if len(_sagemakerDescription) > 0 {
		input.Description = aws.String(_sagemakerDescription)
	}
	if len(_sagemakerProperties) > 0 {
		if err := assignInputField(input, "Properties", _sagemakerProperties); err != nil {
			log.Errorf("invalid --properties: %s", err.Error())
			return
		}
	}
	if len(_sagemakerTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateContext(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a definition for a job that monitors data quality and drift. For
// information about model monitor, see [Amazon SageMaker AI Model Monitor].
//
// [Amazon SageMaker AI Model Monitor]: https://docs.aws.amazon.com/sagemaker/latest/dg/model-monitor.html
func sagemaker_CreateDataQualityJobDefinition(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreateDataQualityJobDefinitionInput{
		// DataQualityAppSpecification: *types.DataQualityAppSpecification, // Required
		// DataQualityJobInput: *types.DataQualityJobInput, // Required
		// DataQualityJobOutputConfig: *types.MonitoringOutputConfig, // Required
		// JobDefinitionName: *string, // Required
		// JobResources: *types.MonitoringResources, // Required
		// RoleArn: *string, // Required
	}

	if len(_sagemakerDataQualityAppSpecification) > 0 {
		if err := assignInputField(input, "DataQualityAppSpecification", _sagemakerDataQualityAppSpecification); err != nil {
			log.Errorf("invalid --data-quality-app-specification: %s", err.Error())
			return
		}
	}
	if len(_sagemakerDataQualityJobInput) > 0 {
		if err := assignInputField(input, "DataQualityJobInput", _sagemakerDataQualityJobInput); err != nil {
			log.Errorf("invalid --data-quality-job-input: %s", err.Error())
			return
		}
	}
	if len(_sagemakerDataQualityJobOutputConfig) > 0 {
		if err := assignInputField(input, "DataQualityJobOutputConfig", _sagemakerDataQualityJobOutputConfig); err != nil {
			log.Errorf("invalid --data-quality-job-output-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerJobDefinitionName) > 0 {
		input.JobDefinitionName = aws.String(_sagemakerJobDefinitionName)
	}
	if len(_sagemakerJobResources) > 0 {
		if err := assignInputField(input, "JobResources", _sagemakerJobResources); err != nil {
			log.Errorf("invalid --job-resources: %s", err.Error())
			return
		}
	}
	if len(_sagemakerRoleArn) > 0 {
		input.RoleArn = aws.String(_sagemakerRoleArn)
	}
	if len(_sagemakerDataQualityBaselineConfig) > 0 {
		if err := assignInputField(input, "DataQualityBaselineConfig", _sagemakerDataQualityBaselineConfig); err != nil {
			log.Errorf("invalid --data-quality-baseline-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNetworkConfig) > 0 {
		if err := assignInputField(input, "NetworkConfig", _sagemakerNetworkConfig); err != nil {
			log.Errorf("invalid --network-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerStoppingCondition) > 0 {
		if err := assignInputField(input, "StoppingCondition", _sagemakerStoppingCondition); err != nil {
			log.Errorf("invalid --stopping-condition: %s", err.Error())
			return
		}
	}
	if len(_sagemakerTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDataQualityJobDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a device fleet.
func sagemaker_CreateDeviceFleet(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreateDeviceFleetInput{
		// DeviceFleetName: *string, // Required
		// OutputConfig: *types.EdgeOutputConfig, // Required
	}

	if len(_sagemakerDeviceFleetName) > 0 {
		input.DeviceFleetName = aws.String(_sagemakerDeviceFleetName)
	}
	if len(_sagemakerOutputConfig) > 0 {
		if err := assignInputField(input, "OutputConfig", _sagemakerOutputConfig); err != nil {
			log.Errorf("invalid --output-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerDescription) > 0 {
		input.Description = aws.String(_sagemakerDescription)
	}
	if len(_sagemakerEnableIotRoleAlias) > 0 {
		if err := assignInputField(input, "EnableIotRoleAlias", _sagemakerEnableIotRoleAlias); err != nil {
			log.Errorf("invalid --enable-iot-role-alias: %s", err.Error())
			return
		}
	}
	if len(_sagemakerRoleArn) > 0 {
		input.RoleArn = aws.String(_sagemakerRoleArn)
	}
	if len(_sagemakerTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDeviceFleet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a Domain . A domain consists of an associated Amazon Elastic File System
// volume, a list of authorized users, and a variety of security, application,
// policy, and Amazon Virtual Private Cloud (VPC) configurations. Users within a
// domain can share notebook files and other artifacts with each other.
//
// # EFS storage
//
// When a domain is created, an EFS volume is created for use by all of the users
// within the domain. Each user receives a private home directory within the EFS
// volume for notebooks, Git repositories, and data files.
//
// SageMaker AI uses the Amazon Web Services Key Management Service (Amazon Web
// Services KMS) to encrypt the EFS volume attached to the domain with an Amazon
// Web Services managed key by default. For more control, you can specify a
// customer managed key. For more information, see [Protect Data at Rest Using Encryption].
//
// # VPC configuration
//
// All traffic between the domain and the Amazon EFS volume is through the
// specified VPC and subnets. For other traffic, you can specify the
// AppNetworkAccessType parameter. AppNetworkAccessType corresponds to the network
// access type that you choose when you onboard to the domain. The following
// options are available:
//
// - PublicInternetOnly - Non-EFS traffic goes through a VPC managed by Amazon
// SageMaker AI, which allows internet access. This is the default value.
//
// - VpcOnly - All traffic is through the specified VPC and subnets. Internet
// access is disabled by default. To allow internet access, you must specify a NAT
// gateway.
//
// When internet access is disabled, you won't be able to run a Amazon SageMaker
//
// AI Studio notebook or to train or host models unless your VPC has an interface
// endpoint to the SageMaker AI API and runtime or a NAT gateway and your security
// groups allow outbound connections.
//
// NFS traffic over TCP on port 2049 needs to be allowed in both inbound and
// outbound rules in order to launch a Amazon SageMaker AI Studio app successfully.
//
// For more information, see [Connect Amazon SageMaker AI Studio Notebooks to Resources in a VPC].
//
// [Connect Amazon SageMaker AI Studio Notebooks to Resources in a VPC]: https://docs.aws.amazon.com/sagemaker/latest/dg/studio-notebooks-and-internet-access.html
// [Protect Data at Rest Using Encryption]: https://docs.aws.amazon.com/sagemaker/latest/dg/encryption-at-rest.html
func sagemaker_CreateDomain(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreateDomainInput{
		// AuthMode: types.AuthMode, // Required
		// DefaultUserSettings: *types.UserSettings, // Required
		// DomainName: *string, // Required
	}

	if len(_sagemakerAuthMode) > 0 {
		if err := assignInputField(input, "AuthMode", _sagemakerAuthMode); err != nil {
			log.Errorf("invalid --auth-mode: %s", err.Error())
			return
		}
	}
	if len(_sagemakerDefaultUserSettings) > 0 {
		if err := assignInputField(input, "DefaultUserSettings", _sagemakerDefaultUserSettings); err != nil {
			log.Errorf("invalid --default-user-settings: %s", err.Error())
			return
		}
	}
	if len(_sagemakerDomainName) > 0 {
		input.DomainName = aws.String(_sagemakerDomainName)
	}
	if len(_sagemakerAppNetworkAccessType) > 0 {
		if err := assignInputField(input, "AppNetworkAccessType", _sagemakerAppNetworkAccessType); err != nil {
			log.Errorf("invalid --app-network-access-type: %s", err.Error())
			return
		}
	}
	if len(_sagemakerAppSecurityGroupManagement) > 0 {
		if err := assignInputField(input, "AppSecurityGroupManagement", _sagemakerAppSecurityGroupManagement); err != nil {
			log.Errorf("invalid --app-security-group-management: %s", err.Error())
			return
		}
	}
	if len(_sagemakerDefaultSpaceSettings) > 0 {
		if err := assignInputField(input, "DefaultSpaceSettings", _sagemakerDefaultSpaceSettings); err != nil {
			log.Errorf("invalid --default-space-settings: %s", err.Error())
			return
		}
	}
	if len(_sagemakerDomainSettings) > 0 {
		if err := assignInputField(input, "DomainSettings", _sagemakerDomainSettings); err != nil {
			log.Errorf("invalid --domain-settings: %s", err.Error())
			return
		}
	}
	if len(_sagemakerHomeEfsFileSystemKmsKeyId) > 0 {
		input.HomeEfsFileSystemKmsKeyId = aws.String(_sagemakerHomeEfsFileSystemKmsKeyId)
	}
	if len(_sagemakerKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_sagemakerKmsKeyId)
	}
	if len(_sagemakerSubnetIds) > 0 {
		input.SubnetIds = append([]string(nil), _sagemakerSubnetIds...)
	}
	if len(_sagemakerTagPropagation) > 0 {
		if err := assignInputField(input, "TagPropagation", _sagemakerTagPropagation); err != nil {
			log.Errorf("invalid --tag-propagation: %s", err.Error())
			return
		}
	}
	if len(_sagemakerTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_sagemakerVpcId) > 0 {
		input.VpcId = aws.String(_sagemakerVpcId)
	}

	if resp, err := client.CreateDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an edge deployment plan, consisting of multiple stages. Each stage may
// have a different deployment configuration and devices.
func sagemaker_CreateEdgeDeploymentPlan(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreateEdgeDeploymentPlanInput{
		// DeviceFleetName: *string, // Required
		// EdgeDeploymentPlanName: *string, // Required
		// ModelConfigs: []types.EdgeDeploymentModelConfig, // Required
	}

	if len(_sagemakerDeviceFleetName) > 0 {
		input.DeviceFleetName = aws.String(_sagemakerDeviceFleetName)
	}
	if len(_sagemakerEdgeDeploymentPlanName) > 0 {
		input.EdgeDeploymentPlanName = aws.String(_sagemakerEdgeDeploymentPlanName)
	}
	if len(_sagemakerModelConfigs) > 0 {
		if err := assignInputField(input, "ModelConfigs", _sagemakerModelConfigs); err != nil {
			log.Errorf("invalid --model-configs: %s", err.Error())
			return
		}
	}
	if len(_sagemakerStages) > 0 {
		if err := assignInputField(input, "Stages", _sagemakerStages); err != nil {
			log.Errorf("invalid --stages: %s", err.Error())
			return
		}
	}
	if len(_sagemakerTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateEdgeDeploymentPlan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new stage in an existing edge deployment plan.
func sagemaker_CreateEdgeDeploymentStage(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreateEdgeDeploymentStageInput{
		// EdgeDeploymentPlanName: *string, // Required
		// Stages: []types.DeploymentStage, // Required
	}

	if len(_sagemakerEdgeDeploymentPlanName) > 0 {
		input.EdgeDeploymentPlanName = aws.String(_sagemakerEdgeDeploymentPlanName)
	}
	if len(_sagemakerStages) > 0 {
		if err := assignInputField(input, "Stages", _sagemakerStages); err != nil {
			log.Errorf("invalid --stages: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateEdgeDeploymentStage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a SageMaker Edge Manager model packaging job. Edge Manager will use the
// model artifacts from the Amazon Simple Storage Service bucket that you specify.
// After the model has been packaged, Amazon SageMaker saves the resulting
// artifacts to an S3 bucket that you specify.
func sagemaker_CreateEdgePackagingJob(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreateEdgePackagingJobInput{
		// CompilationJobName: *string, // Required
		// EdgePackagingJobName: *string, // Required
		// ModelName: *string, // Required
		// ModelVersion: *string, // Required
		// OutputConfig: *types.EdgeOutputConfig, // Required
		// RoleArn: *string, // Required
	}

	if len(_sagemakerCompilationJobName) > 0 {
		input.CompilationJobName = aws.String(_sagemakerCompilationJobName)
	}
	if len(_sagemakerEdgePackagingJobName) > 0 {
		input.EdgePackagingJobName = aws.String(_sagemakerEdgePackagingJobName)
	}
	if len(_sagemakerModelName) > 0 {
		input.ModelName = aws.String(_sagemakerModelName)
	}
	if len(_sagemakerModelVersion) > 0 {
		input.ModelVersion = aws.String(_sagemakerModelVersion)
	}
	if len(_sagemakerOutputConfig) > 0 {
		if err := assignInputField(input, "OutputConfig", _sagemakerOutputConfig); err != nil {
			log.Errorf("invalid --output-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerRoleArn) > 0 {
		input.RoleArn = aws.String(_sagemakerRoleArn)
	}
	if len(_sagemakerResourceKey) > 0 {
		input.ResourceKey = aws.String(_sagemakerResourceKey)
	}
	if len(_sagemakerTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateEdgePackagingJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an endpoint using the endpoint configuration specified in the request.
// SageMaker uses the endpoint to provision resources and deploy models. You create
// the endpoint configuration with the [CreateEndpointConfig]API.
//
// Use this API to deploy models using SageMaker hosting services.
//
// You must not delete an EndpointConfig that is in use by an endpoint that is
// live or while the UpdateEndpoint or CreateEndpoint operations are being
// performed on the endpoint. To update an endpoint, you must create a new
// EndpointConfig .
//
// The endpoint name must be unique within an Amazon Web Services Region in your
// Amazon Web Services account.
//
// When it receives the request, SageMaker creates the endpoint, launches the
// resources (ML compute instances), and deploys the model(s) on them.
//
// When you call [CreateEndpoint], a load call is made to DynamoDB to verify that your endpoint
// configuration exists. When you read data from a DynamoDB table supporting [Eventually Consistent Reads]
// Eventually Consistent Reads , the response might not reflect the results of a
// recently completed write operation. The response might include some stale data.
// If the dependent entities are not yet in DynamoDB, this causes a validation
// error. If you repeat your read request after a short time, the response should
// return the latest data. So retry logic is recommended to handle these possible
// issues. We also recommend that customers call [DescribeEndpointConfig]before calling [CreateEndpoint] to minimize the
// potential impact of a DynamoDB eventually consistent read.
//
// When SageMaker receives the request, it sets the endpoint status to Creating .
// After it creates the endpoint, it sets the status to InService . SageMaker can
// then process incoming requests for inferences. To check the status of an
// endpoint, use the [DescribeEndpoint]API.
//
// If any of the models hosted at this endpoint get model data from an Amazon S3
// location, SageMaker uses Amazon Web Services Security Token Service to download
// model artifacts from the S3 path you provided. Amazon Web Services STS is
// activated in your Amazon Web Services account by default. If you previously
// deactivated Amazon Web Services STS for a region, you need to reactivate Amazon
// Web Services STS for that region. For more information, see [Activating and Deactivating Amazon Web Services STS in an Amazon Web Services Region]in the Amazon Web
// Services Identity and Access Management User Guide.
//
// To add the IAM role policies for using this API operation, go to the [IAM console], and
// choose Roles in the left navigation pane. Search the IAM role that you want to
// grant access to use the [CreateEndpoint]and [CreateEndpointConfig] API operations, add the following policies to the
// role.
//
// - Option 1: For a full SageMaker access, search and attach the
// AmazonSageMakerFullAccess policy.
//
// - Option 2: For granting a limited access to an IAM role, paste the following
// Action elements manually into the JSON file of the IAM role:
//
// "Action": ["sagemaker:CreateEndpoint", "sagemaker:CreateEndpointConfig"]
//
// "Resource": [
//
// "arn:aws:sagemaker:region:account-id:endpoint/endpointName"
//
// "arn:aws:sagemaker:region:account-id:endpoint-config/endpointConfigName"
//
// ]
//
// For more information, see [SageMaker API Permissions: Actions, Permissions, and Resources Reference].
//
// [Eventually Consistent Reads]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadConsistency.html
// [IAM console]: https://console.aws.amazon.com/iam/
// [DescribeEndpoint]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DescribeEndpoint.html
// [CreateEndpoint]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateEndpoint.html
// [CreateEndpointConfig]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateEndpointConfig.html
// [DescribeEndpointConfig]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DescribeEndpointConfig.html
// [Activating and Deactivating Amazon Web Services STS in an Amazon Web Services Region]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_enable-regions.html
// [SageMaker API Permissions: Actions, Permissions, and Resources Reference]: https://docs.aws.amazon.com/sagemaker/latest/dg/api-permissions-reference.html
func sagemaker_CreateEndpoint(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreateEndpointInput{
		// EndpointConfigName: *string, // Required
		// EndpointName: *string, // Required
	}

	if len(_sagemakerEndpointConfigName) > 0 {
		input.EndpointConfigName = aws.String(_sagemakerEndpointConfigName)
	}
	if len(_sagemakerEndpointName) > 0 {
		input.EndpointName = aws.String(_sagemakerEndpointName)
	}
	if len(_sagemakerDeploymentConfig) > 0 {
		if err := assignInputField(input, "DeploymentConfig", _sagemakerDeploymentConfig); err != nil {
			log.Errorf("invalid --deployment-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an endpoint configuration that SageMaker hosting services uses to
// deploy models. In the configuration, you identify one or more models, created
// using the CreateModel API, to deploy and the resources that you want SageMaker
// to provision. Then you call the [CreateEndpoint]API.
//
// Use this API if you want to use SageMaker hosting services to deploy models
// into production.
//
// In the request, you define a ProductionVariant , for each model that you want to
// deploy. Each ProductionVariant parameter also describes the resources that you
// want SageMaker to provision. This includes the number and type of ML compute
// instances to deploy.
//
// If you are hosting multiple models, you also assign a VariantWeight to specify
// how much traffic you want to allocate to each model. For example, suppose that
// you want to host two models, A and B, and you assign traffic weight 2 for model
// A and 1 for model B. SageMaker distributes two-thirds of the traffic to Model A,
// and one-third to model B.
//
// When you call [CreateEndpoint], a load call is made to DynamoDB to verify that your endpoint
// configuration exists. When you read data from a DynamoDB table supporting [Eventually Consistent Reads]
// Eventually Consistent Reads , the response might not reflect the results of a
// recently completed write operation. The response might include some stale data.
// If the dependent entities are not yet in DynamoDB, this causes a validation
// error. If you repeat your read request after a short time, the response should
// return the latest data. So retry logic is recommended to handle these possible
// issues. We also recommend that customers call [DescribeEndpointConfig]before calling [CreateEndpoint] to minimize the
// potential impact of a DynamoDB eventually consistent read.
//
// [Eventually Consistent Reads]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadConsistency.html
// [CreateEndpoint]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateEndpoint.html
// [DescribeEndpointConfig]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DescribeEndpointConfig.html
func sagemaker_CreateEndpointConfig(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreateEndpointConfigInput{
		// EndpointConfigName: *string, // Required
		// ProductionVariants: []types.ProductionVariant, // Required
	}

	if len(_sagemakerEndpointConfigName) > 0 {
		input.EndpointConfigName = aws.String(_sagemakerEndpointConfigName)
	}
	if len(_sagemakerProductionVariants) > 0 {
		if err := assignInputField(input, "ProductionVariants", _sagemakerProductionVariants); err != nil {
			log.Errorf("invalid --production-variants: %s", err.Error())
			return
		}
	}
	if len(_sagemakerAsyncInferenceConfig) > 0 {
		if err := assignInputField(input, "AsyncInferenceConfig", _sagemakerAsyncInferenceConfig); err != nil {
			log.Errorf("invalid --async-inference-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerDataCaptureConfig) > 0 {
		if err := assignInputField(input, "DataCaptureConfig", _sagemakerDataCaptureConfig); err != nil {
			log.Errorf("invalid --data-capture-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerEnableNetworkIsolation) > 0 {
		if err := assignInputField(input, "EnableNetworkIsolation", _sagemakerEnableNetworkIsolation); err != nil {
			log.Errorf("invalid --enable-network-isolation: %s", err.Error())
			return
		}
	}
	if len(_sagemakerExecutionRoleArn) > 0 {
		input.ExecutionRoleArn = aws.String(_sagemakerExecutionRoleArn)
	}
	if len(_sagemakerExplainerConfig) > 0 {
		if err := assignInputField(input, "ExplainerConfig", _sagemakerExplainerConfig); err != nil {
			log.Errorf("invalid --explainer-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_sagemakerKmsKeyId)
	}
	if len(_sagemakerMetricsConfig) > 0 {
		if err := assignInputField(input, "MetricsConfig", _sagemakerMetricsConfig); err != nil {
			log.Errorf("invalid --metrics-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerShadowProductionVariants) > 0 {
		if err := assignInputField(input, "ShadowProductionVariants", _sagemakerShadowProductionVariants); err != nil {
			log.Errorf("invalid --shadow-production-variants: %s", err.Error())
			return
		}
	}
	if len(_sagemakerTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_sagemakerVpcConfig) > 0 {
		if err := assignInputField(input, "VpcConfig", _sagemakerVpcConfig); err != nil {
			log.Errorf("invalid --vpc-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateEndpointConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a SageMaker experiment. An experiment is a collection of trials that
// are observed, compared and evaluated as a group. A trial is a set of steps,
// called trial components, that produce a machine learning model.
//
// In the Studio UI, trials are referred to as run groups and trial components are
// referred to as runs.
//
// The goal of an experiment is to determine the components that produce the best
// model. Multiple trials are performed, each one isolating and measuring the
// impact of a change to one or more inputs, while keeping the remaining inputs
// constant.
//
// When you use SageMaker Studio or the SageMaker Python SDK, all experiments,
// trials, and trial components are automatically tracked, logged, and indexed.
// When you use the Amazon Web Services SDK for Python (Boto), you must use the
// logging APIs provided by the SDK.
//
// You can add tags to experiments, trials, trial components and then use the [Search] API
// to search for the tags.
//
// To add a description to an experiment, specify the optional Description
// parameter. To add a description later, or to change the description, call the [UpdateExperiment]
// API.
//
// To get a list of all your experiments, call the [ListExperiments] API. To view an experiment's
// properties, call the [DescribeExperiment]API. To get a list of all the trials associated with an
// experiment, call the [ListTrials]API. To create a trial call the [CreateTrial] API.
//
// [DescribeExperiment]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DescribeExperiment.html
// [Search]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_Search.html
// [CreateTrial]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateTrial.html
// [ListExperiments]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_ListExperiments.html
// [UpdateExperiment]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_UpdateExperiment.html
// [ListTrials]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_ListTrials.html
func sagemaker_CreateExperiment(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreateExperimentInput{
		// ExperimentName: *string, // Required
	}

	if len(_sagemakerExperimentName) > 0 {
		input.ExperimentName = aws.String(_sagemakerExperimentName)
	}
	if len(_sagemakerDescription) > 0 {
		input.Description = aws.String(_sagemakerDescription)
	}
	if len(_sagemakerDisplayName) > 0 {
		input.DisplayName = aws.String(_sagemakerDisplayName)
	}
	if len(_sagemakerTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateExperiment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a new FeatureGroup . A FeatureGroup is a group of Features defined in
// the FeatureStore to describe a Record .
//
// The FeatureGroup defines the schema and features contained in the FeatureGroup .
// A FeatureGroup definition is composed of a list of Features , a
// RecordIdentifierFeatureName , an EventTimeFeatureName and configurations for
// its OnlineStore and OfflineStore . Check [Amazon Web Services service quotas] to see the FeatureGroup s quota for
// your Amazon Web Services account.
//
// Note that it can take approximately 10-15 minutes to provision an OnlineStore
// FeatureGroup with the InMemory StorageType .
//
// You must include at least one of OnlineStoreConfig and OfflineStoreConfig to
// create a FeatureGroup .
//
// [Amazon Web Services service quotas]: https://docs.aws.amazon.com/general/latest/gr/aws_service_limits.html
func sagemaker_CreateFeatureGroup(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreateFeatureGroupInput{
		// EventTimeFeatureName: *string, // Required
		// FeatureDefinitions: []types.FeatureDefinition, // Required
		// FeatureGroupName: *string, // Required
		// RecordIdentifierFeatureName: *string, // Required
	}

	if len(_sagemakerEventTimeFeatureName) > 0 {
		input.EventTimeFeatureName = aws.String(_sagemakerEventTimeFeatureName)
	}
	if len(_sagemakerFeatureDefinitions) > 0 {
		if err := assignInputField(input, "FeatureDefinitions", _sagemakerFeatureDefinitions); err != nil {
			log.Errorf("invalid --feature-definitions: %s", err.Error())
			return
		}
	}
	if len(_sagemakerFeatureGroupName) > 0 {
		input.FeatureGroupName = aws.String(_sagemakerFeatureGroupName)
	}
	if len(_sagemakerRecordIdentifierFeatureName) > 0 {
		input.RecordIdentifierFeatureName = aws.String(_sagemakerRecordIdentifierFeatureName)
	}
	if len(_sagemakerDescription) > 0 {
		input.Description = aws.String(_sagemakerDescription)
	}
	if len(_sagemakerOfflineStoreConfig) > 0 {
		if err := assignInputField(input, "OfflineStoreConfig", _sagemakerOfflineStoreConfig); err != nil {
			log.Errorf("invalid --offline-store-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerOnlineStoreConfig) > 0 {
		if err := assignInputField(input, "OnlineStoreConfig", _sagemakerOnlineStoreConfig); err != nil {
			log.Errorf("invalid --online-store-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerRoleArn) > 0 {
		input.RoleArn = aws.String(_sagemakerRoleArn)
	}
	if len(_sagemakerTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_sagemakerThroughputConfig) > 0 {
		if err := assignInputField(input, "ThroughputConfig", _sagemakerThroughputConfig); err != nil {
			log.Errorf("invalid --throughput-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateFeatureGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a flow definition.
func sagemaker_CreateFlowDefinition(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreateFlowDefinitionInput{
		// FlowDefinitionName: *string, // Required
		// OutputConfig: *types.FlowDefinitionOutputConfig, // Required
		// RoleArn: *string, // Required
	}

	if len(_sagemakerFlowDefinitionName) > 0 {
		input.FlowDefinitionName = aws.String(_sagemakerFlowDefinitionName)
	}
	if len(_sagemakerOutputConfig) > 0 {
		if err := assignInputField(input, "OutputConfig", _sagemakerOutputConfig); err != nil {
			log.Errorf("invalid --output-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerRoleArn) > 0 {
		input.RoleArn = aws.String(_sagemakerRoleArn)
	}
	if len(_sagemakerHumanLoopActivationConfig) > 0 {
		if err := assignInputField(input, "HumanLoopActivationConfig", _sagemakerHumanLoopActivationConfig); err != nil {
			log.Errorf("invalid --human-loop-activation-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerHumanLoopConfig) > 0 {
		if err := assignInputField(input, "HumanLoopConfig", _sagemakerHumanLoopConfig); err != nil {
			log.Errorf("invalid --human-loop-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerHumanLoopRequestSource) > 0 {
		if err := assignInputField(input, "HumanLoopRequestSource", _sagemakerHumanLoopRequestSource); err != nil {
			log.Errorf("invalid --human-loop-request-source: %s", err.Error())
			return
		}
	}
	if len(_sagemakerTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateFlowDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a hub.
func sagemaker_CreateHub(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreateHubInput{
		// HubDescription: *string, // Required
		// HubName: *string, // Required
	}

	if len(_sagemakerHubDescription) > 0 {
		input.HubDescription = aws.String(_sagemakerHubDescription)
	}
	if len(_sagemakerHubName) > 0 {
		input.HubName = aws.String(_sagemakerHubName)
	}
	if len(_sagemakerHubDisplayName) > 0 {
		input.HubDisplayName = aws.String(_sagemakerHubDisplayName)
	}
	if len(_sagemakerHubSearchKeywords) > 0 {
		input.HubSearchKeywords = append([]string(nil), _sagemakerHubSearchKeywords...)
	}
	if len(_sagemakerS3StorageConfig) > 0 {
		if err := assignInputField(input, "S3StorageConfig", _sagemakerS3StorageConfig); err != nil {
			log.Errorf("invalid --s3-storage-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateHub(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates presigned URLs for accessing hub content artifacts. This operation
// generates time-limited, secure URLs that allow direct download of model
// artifacts and associated files from Amazon SageMaker hub content, including
// gated models that require end-user license agreement acceptance.
func sagemaker_CreateHubContentPresignedUrls(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreateHubContentPresignedUrlsInput{
		// HubContentName: *string, // Required
		// HubContentType: types.HubContentType, // Required
		// HubName: *string, // Required
	}

	if len(_sagemakerHubContentName) > 0 {
		input.HubContentName = aws.String(_sagemakerHubContentName)
	}
	if len(_sagemakerHubContentType) > 0 {
		if err := assignInputField(input, "HubContentType", _sagemakerHubContentType); err != nil {
			log.Errorf("invalid --hub-content-type: %s", err.Error())
			return
		}
	}
	if len(_sagemakerHubName) > 0 {
		input.HubName = aws.String(_sagemakerHubName)
	}
	if len(_sagemakerAccessConfig) > 0 {
		if err := assignInputField(input, "AccessConfig", _sagemakerAccessConfig); err != nil {
			log.Errorf("invalid --access-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerHubContentVersion) > 0 {
		input.HubContentVersion = aws.String(_sagemakerHubContentVersion)
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.CreateHubContentPresignedUrls(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.CreateHubContentPresignedUrlsOutput
	p := sagemaker.NewCreateHubContentPresignedUrlsPaginator(client, input)
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

// Create a hub content reference in order to add a model in the JumpStart public
// hub to a private hub.
func sagemaker_CreateHubContentReference(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreateHubContentReferenceInput{
		// HubName: *string, // Required
		// SageMakerPublicHubContentArn: *string, // Required
	}

	if len(_sagemakerHubName) > 0 {
		input.HubName = aws.String(_sagemakerHubName)
	}
	if len(_sagemakerSageMakerPublicHubContentArn) > 0 {
		input.SageMakerPublicHubContentArn = aws.String(_sagemakerSageMakerPublicHubContentArn)
	}
	if len(_sagemakerHubContentName) > 0 {
		input.HubContentName = aws.String(_sagemakerHubContentName)
	}
	if len(_sagemakerMinVersion) > 0 {
		input.MinVersion = aws.String(_sagemakerMinVersion)
	}
	if len(_sagemakerTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateHubContentReference(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Defines the settings you will use for the human review workflow user interface.
// Reviewers will see a three-panel interface with an instruction area, the item to
// review, and an input area.
func sagemaker_CreateHumanTaskUi(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreateHumanTaskUiInput{
		// HumanTaskUiName: *string, // Required
		// UiTemplate: *types.UiTemplate, // Required
	}

	if len(_sagemakerHumanTaskUiName) > 0 {
		input.HumanTaskUiName = aws.String(_sagemakerHumanTaskUiName)
	}
	if len(_sagemakerUiTemplate) > 0 {
		if err := assignInputField(input, "UiTemplate", _sagemakerUiTemplate); err != nil {
			log.Errorf("invalid --ui-template: %s", err.Error())
			return
		}
	}
	if len(_sagemakerTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateHumanTaskUi(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a hyperparameter tuning job. A hyperparameter tuning job finds the best
// version of a model by running many training jobs on your dataset using the
// algorithm you choose and values for hyperparameters within ranges that you
// specify. It then chooses the hyperparameter values that result in a model that
// performs the best, as measured by an objective metric that you choose.
//
// A hyperparameter tuning job automatically creates Amazon SageMaker experiments,
// trials, and trial components for each training job that it runs. You can view
// these entities in Amazon SageMaker Studio. For more information, see [View Experiments, Trials, and Trial Components].
//
// Do not include any security-sensitive information including account access IDs,
// secrets, or tokens in any hyperparameter fields. As part of the shared
// responsibility model, you are responsible for any potential exposure,
// unauthorized access, or compromise of your sensitive data if caused by any
// security-sensitive information included in the request hyperparameter variable
// or plain text fields..
//
// [View Experiments, Trials, and Trial Components]: https://docs.aws.amazon.com/sagemaker/latest/dg/experiments-view-compare.html#experiments-view
func sagemaker_CreateHyperParameterTuningJob(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreateHyperParameterTuningJobInput{
		// HyperParameterTuningJobConfig: *types.HyperParameterTuningJobConfig, // Required
		// HyperParameterTuningJobName: *string, // Required
	}

	if len(_sagemakerHyperParameterTuningJobConfig) > 0 {
		if err := assignInputField(input, "HyperParameterTuningJobConfig", _sagemakerHyperParameterTuningJobConfig); err != nil {
			log.Errorf("invalid --hyper-parameter-tuning-job-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerHyperParameterTuningJobName) > 0 {
		input.HyperParameterTuningJobName = aws.String(_sagemakerHyperParameterTuningJobName)
	}
	if len(_sagemakerAutotune) > 0 {
		if err := assignInputField(input, "Autotune", _sagemakerAutotune); err != nil {
			log.Errorf("invalid --autotune: %s", err.Error())
			return
		}
	}
	if len(_sagemakerTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_sagemakerTrainingJobDefinition) > 0 {
		if err := assignInputField(input, "TrainingJobDefinition", _sagemakerTrainingJobDefinition); err != nil {
			log.Errorf("invalid --training-job-definition: %s", err.Error())
			return
		}
	}
	if len(_sagemakerTrainingJobDefinitions) > 0 {
		if err := assignInputField(input, "TrainingJobDefinitions", _sagemakerTrainingJobDefinitions); err != nil {
			log.Errorf("invalid --training-job-definitions: %s", err.Error())
			return
		}
	}
	if len(_sagemakerWarmStartConfig) > 0 {
		if err := assignInputField(input, "WarmStartConfig", _sagemakerWarmStartConfig); err != nil {
			log.Errorf("invalid --warm-start-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateHyperParameterTuningJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a custom SageMaker AI image. A SageMaker AI image is a set of image
// versions. Each image version represents a container image stored in Amazon ECR.
// For more information, see [Bring your own SageMaker AI image].
//
// [Bring your own SageMaker AI image]: https://docs.aws.amazon.com/sagemaker/latest/dg/studio-byoi.html
func sagemaker_CreateImage(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreateImageInput{
		// ImageName: *string, // Required
		// RoleArn: *string, // Required
	}

	if len(_sagemakerImageName) > 0 {
		input.ImageName = aws.String(_sagemakerImageName)
	}
	if len(_sagemakerRoleArn) > 0 {
		input.RoleArn = aws.String(_sagemakerRoleArn)
	}
	if len(_sagemakerDescription) > 0 {
		input.Description = aws.String(_sagemakerDescription)
	}
	if len(_sagemakerDisplayName) > 0 {
		input.DisplayName = aws.String(_sagemakerDisplayName)
	}
	if len(_sagemakerTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateImage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a version of the SageMaker AI image specified by ImageName . The version
// represents the Amazon ECR container image specified by BaseImage .
func sagemaker_CreateImageVersion(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreateImageVersionInput{
		// BaseImage: *string, // Required
		// ClientToken: *string, // Required
		// ImageName: *string, // Required
	}

	if len(_sagemakerBaseImage) > 0 {
		input.BaseImage = aws.String(_sagemakerBaseImage)
	}
	if len(_sagemakerClientToken) > 0 {
		input.ClientToken = aws.String(_sagemakerClientToken)
	}
	if len(_sagemakerImageName) > 0 {
		input.ImageName = aws.String(_sagemakerImageName)
	}
	if len(_sagemakerAliases) > 0 {
		input.Aliases = append([]string(nil), _sagemakerAliases...)
	}
	if len(_sagemakerHorovod) > 0 {
		if err := assignInputField(input, "Horovod", _sagemakerHorovod); err != nil {
			log.Errorf("invalid --horovod: %s", err.Error())
			return
		}
	}
	if len(_sagemakerJobType) > 0 {
		if err := assignInputField(input, "JobType", _sagemakerJobType); err != nil {
			log.Errorf("invalid --job-type: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMLFramework) > 0 {
		input.MLFramework = aws.String(_sagemakerMLFramework)
	}
	if len(_sagemakerProcessor) > 0 {
		if err := assignInputField(input, "Processor", _sagemakerProcessor); err != nil {
			log.Errorf("invalid --processor: %s", err.Error())
			return
		}
	}
	if len(_sagemakerProgrammingLang) > 0 {
		input.ProgrammingLang = aws.String(_sagemakerProgrammingLang)
	}
	if len(_sagemakerReleaseNotes) > 0 {
		input.ReleaseNotes = aws.String(_sagemakerReleaseNotes)
	}
	if len(_sagemakerVendorGuidance) > 0 {
		if err := assignInputField(input, "VendorGuidance", _sagemakerVendorGuidance); err != nil {
			log.Errorf("invalid --vendor-guidance: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateImageVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an inference component, which is a SageMaker AI hosting object that you
// can use to deploy a model to an endpoint. In the inference component settings,
// you specify the model, the endpoint, and how the model utilizes the resources
// that the endpoint hosts. You can optimize resource utilization by tailoring how
// the required CPU cores, accelerators, and memory are allocated. You can deploy
// multiple inference components to an endpoint, where each inference component
// contains one model and the resource utilization needs for that individual model.
// After you deploy an inference component, you can directly invoke the associated
// model when you use the InvokeEndpoint API action.
func sagemaker_CreateInferenceComponent(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreateInferenceComponentInput{
		// EndpointName: *string, // Required
		// InferenceComponentName: *string, // Required
		// Specification: *types.InferenceComponentSpecification, // Required
	}

	if len(_sagemakerEndpointName) > 0 {
		input.EndpointName = aws.String(_sagemakerEndpointName)
	}
	if len(_sagemakerInferenceComponentName) > 0 {
		input.InferenceComponentName = aws.String(_sagemakerInferenceComponentName)
	}
	if len(_sagemakerSpecification) > 0 {
		if err := assignInputField(input, "Specification", _sagemakerSpecification); err != nil {
			log.Errorf("invalid --specification: %s", err.Error())
			return
		}
	}
	if len(_sagemakerRuntimeConfig) > 0 {
		if err := assignInputField(input, "RuntimeConfig", _sagemakerRuntimeConfig); err != nil {
			log.Errorf("invalid --runtime-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_sagemakerVariantName) > 0 {
		input.VariantName = aws.String(_sagemakerVariantName)
	}

	if resp, err := client.CreateInferenceComponent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an inference experiment using the configurations specified in the
// request.
//
// Use this API to setup and schedule an experiment to compare model variants on a
// Amazon SageMaker inference endpoint. For more information about inference
// experiments, see [Shadow tests].
//
// Amazon SageMaker begins your experiment at the scheduled time and routes
// traffic to your endpoint's model variants based on your specified configuration.
//
// While the experiment is in progress or after it has concluded, you can view
// metrics that compare your model variants. For more information, see [View, monitor, and edit shadow tests].
//
// [Shadow tests]: https://docs.aws.amazon.com/sagemaker/latest/dg/shadow-tests.html
// [View, monitor, and edit shadow tests]: https://docs.aws.amazon.com/sagemaker/latest/dg/shadow-tests-view-monitor-edit.html
func sagemaker_CreateInferenceExperiment(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreateInferenceExperimentInput{
		// EndpointName: *string, // Required
		// ModelVariants: []types.ModelVariantConfig, // Required
		// Name: *string, // Required
		// RoleArn: *string, // Required
		// ShadowModeConfig: *types.ShadowModeConfig, // Required
		// Type: types.InferenceExperimentType, // Required
	}

	if len(_sagemakerEndpointName) > 0 {
		input.EndpointName = aws.String(_sagemakerEndpointName)
	}
	if len(_sagemakerModelVariants) > 0 {
		if err := assignInputField(input, "ModelVariants", _sagemakerModelVariants); err != nil {
			log.Errorf("invalid --model-variants: %s", err.Error())
			return
		}
	}
	if len(_sagemakerName) > 0 {
		input.Name = aws.String(_sagemakerName)
	}
	if len(_sagemakerRoleArn) > 0 {
		input.RoleArn = aws.String(_sagemakerRoleArn)
	}
	if len(_sagemakerShadowModeConfig) > 0 {
		if err := assignInputField(input, "ShadowModeConfig", _sagemakerShadowModeConfig); err != nil {
			log.Errorf("invalid --shadow-mode-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerType) > 0 {
		if err := assignInputField(input, "Type", _sagemakerType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_sagemakerDataStorageConfig) > 0 {
		if err := assignInputField(input, "DataStorageConfig", _sagemakerDataStorageConfig); err != nil {
			log.Errorf("invalid --data-storage-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerDescription) > 0 {
		input.Description = aws.String(_sagemakerDescription)
	}
	if len(_sagemakerKmsKey) > 0 {
		input.KmsKey = aws.String(_sagemakerKmsKey)
	}
	if len(_sagemakerSchedule) > 0 {
		if err := assignInputField(input, "Schedule", _sagemakerSchedule); err != nil {
			log.Errorf("invalid --schedule: %s", err.Error())
			return
		}
	}
	if len(_sagemakerTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateInferenceExperiment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a recommendation job. You can create either an instance recommendation
// or load test job.
func sagemaker_CreateInferenceRecommendationsJob(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreateInferenceRecommendationsJobInput{
		// InputConfig: *types.RecommendationJobInputConfig, // Required
		// JobName: *string, // Required
		// JobType: types.RecommendationJobType, // Required
		// RoleArn: *string, // Required
	}

	if len(_sagemakerInputConfig) > 0 {
		if err := assignInputField(input, "InputConfig", _sagemakerInputConfig); err != nil {
			log.Errorf("invalid --input-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerJobName) > 0 {
		input.JobName = aws.String(_sagemakerJobName)
	}
	if len(_sagemakerJobType) > 0 {
		if err := assignInputField(input, "JobType", _sagemakerJobType); err != nil {
			log.Errorf("invalid --job-type: %s", err.Error())
			return
		}
	}
	if len(_sagemakerRoleArn) > 0 {
		input.RoleArn = aws.String(_sagemakerRoleArn)
	}
	if len(_sagemakerJobDescription) > 0 {
		input.JobDescription = aws.String(_sagemakerJobDescription)
	}
	if len(_sagemakerOutputConfig) > 0 {
		if err := assignInputField(input, "OutputConfig", _sagemakerOutputConfig); err != nil {
			log.Errorf("invalid --output-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerStoppingConditions) > 0 {
		if err := assignInputField(input, "StoppingConditions", _sagemakerStoppingConditions); err != nil {
			log.Errorf("invalid --stopping-conditions: %s", err.Error())
			return
		}
	}
	if len(_sagemakerTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateInferenceRecommendationsJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a job that uses workers to label the data objects in your input
// dataset. You can use the labeled data to train machine learning models.
//
// You can select your workforce from one of three providers:
//
// - A private workforce that you create. It can include employees, contractors,
// and outside experts. Use a private workforce when want the data to stay within
// your organization or when a specific set of skills is required.
//
// - One or more vendors that you select from the Amazon Web Services
// Marketplace. Vendors provide expertise in specific areas.
//
// - The Amazon Mechanical Turk workforce. This is the largest workforce, but it
// should only be used for public data or data that has been stripped of any
// personally identifiable information.
//
// You can also use automated data labeling to reduce the number of data objects
// that need to be labeled by a human. Automated data labeling uses active learning
// to determine if a data object can be labeled by machine or if it needs to be
// sent to a human worker. For more information, see [Using Automated Data Labeling].
//
// The data objects to be labeled are contained in an Amazon S3 bucket. You create
// a manifest file that describes the location of each object. For more
// information, see [Using Input and Output Data].
//
// The output can be used as the manifest file for another labeling job or as
// training data for your machine learning models.
//
// You can use this operation to create a static labeling job or a streaming
// labeling job. A static labeling job stops if all data objects in the input
// manifest file identified in ManifestS3Uri have been labeled. A streaming
// labeling job runs perpetually until it is manually stopped, or remains idle for
// 10 days. You can send new data objects to an active ( InProgress ) streaming
// labeling job in real time. To learn how to create a static labeling job, see [Create a Labeling Job (API)]in
// the Amazon SageMaker Developer Guide. To learn how to create a streaming
// labeling job, see [Create a Streaming Labeling Job].
//
// [Using Automated Data Labeling]: https://docs.aws.amazon.com/sagemaker/latest/dg/sms-automated-labeling.html
// [Create a Streaming Labeling Job]: https://docs.aws.amazon.com/sagemaker/latest/dg/sms-streaming-create-job.html
// [Create a Labeling Job (API)]: https://docs.aws.amazon.com/sagemaker/latest/dg/sms-create-labeling-job-api.html
// [Using Input and Output Data]: https://docs.aws.amazon.com/sagemaker/latest/dg/sms-data.html
func sagemaker_CreateLabelingJob(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreateLabelingJobInput{
		// HumanTaskConfig: *types.HumanTaskConfig, // Required
		// InputConfig: *types.LabelingJobInputConfig, // Required
		// LabelAttributeName: *string, // Required
		// LabelingJobName: *string, // Required
		// OutputConfig: *types.LabelingJobOutputConfig, // Required
		// RoleArn: *string, // Required
	}

	if len(_sagemakerHumanTaskConfig) > 0 {
		if err := assignInputField(input, "HumanTaskConfig", _sagemakerHumanTaskConfig); err != nil {
			log.Errorf("invalid --human-task-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerInputConfig) > 0 {
		if err := assignInputField(input, "InputConfig", _sagemakerInputConfig); err != nil {
			log.Errorf("invalid --input-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerLabelAttributeName) > 0 {
		input.LabelAttributeName = aws.String(_sagemakerLabelAttributeName)
	}
	if len(_sagemakerLabelingJobName) > 0 {
		input.LabelingJobName = aws.String(_sagemakerLabelingJobName)
	}
	if len(_sagemakerOutputConfig) > 0 {
		if err := assignInputField(input, "OutputConfig", _sagemakerOutputConfig); err != nil {
			log.Errorf("invalid --output-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerRoleArn) > 0 {
		input.RoleArn = aws.String(_sagemakerRoleArn)
	}
	if len(_sagemakerLabelCategoryConfigS3Uri) > 0 {
		input.LabelCategoryConfigS3Uri = aws.String(_sagemakerLabelCategoryConfigS3Uri)
	}
	if len(_sagemakerLabelingJobAlgorithmsConfig) > 0 {
		if err := assignInputField(input, "LabelingJobAlgorithmsConfig", _sagemakerLabelingJobAlgorithmsConfig); err != nil {
			log.Errorf("invalid --labeling-job-algorithms-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerStoppingConditions) > 0 {
		if err := assignInputField(input, "StoppingConditions", _sagemakerStoppingConditions); err != nil {
			log.Errorf("invalid --stopping-conditions: %s", err.Error())
			return
		}
	}
	if len(_sagemakerTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateLabelingJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an MLflow Tracking Server using a general purpose Amazon S3 bucket as
// the artifact store.
func sagemaker_CreateMlflowApp(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreateMlflowAppInput{
		// ArtifactStoreUri: *string, // Required
		// Name: *string, // Required
		// RoleArn: *string, // Required
	}

	if len(_sagemakerArtifactStoreUri) > 0 {
		input.ArtifactStoreUri = aws.String(_sagemakerArtifactStoreUri)
	}
	if len(_sagemakerName) > 0 {
		input.Name = aws.String(_sagemakerName)
	}
	if len(_sagemakerRoleArn) > 0 {
		input.RoleArn = aws.String(_sagemakerRoleArn)
	}
	if len(_sagemakerAccountDefaultStatus) > 0 {
		if err := assignInputField(input, "AccountDefaultStatus", _sagemakerAccountDefaultStatus); err != nil {
			log.Errorf("invalid --account-default-status: %s", err.Error())
			return
		}
	}
	if len(_sagemakerDefaultDomainIdList) > 0 {
		input.DefaultDomainIdList = append([]string(nil), _sagemakerDefaultDomainIdList...)
	}
	if len(_sagemakerModelRegistrationMode) > 0 {
		if err := assignInputField(input, "ModelRegistrationMode", _sagemakerModelRegistrationMode); err != nil {
			log.Errorf("invalid --model-registration-mode: %s", err.Error())
			return
		}
	}
	if len(_sagemakerTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_sagemakerWeeklyMaintenanceWindowStart) > 0 {
		input.WeeklyMaintenanceWindowStart = aws.String(_sagemakerWeeklyMaintenanceWindowStart)
	}

	if resp, err := client.CreateMlflowApp(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an MLflow Tracking Server using a general purpose Amazon S3 bucket as
// the artifact store. For more information, see [Create an MLflow Tracking Server].
//
// [Create an MLflow Tracking Server]: https://docs.aws.amazon.com/sagemaker/latest/dg/mlflow-create-tracking-server.html
func sagemaker_CreateMlflowTrackingServer(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreateMlflowTrackingServerInput{
		// ArtifactStoreUri: *string, // Required
		// RoleArn: *string, // Required
		// TrackingServerName: *string, // Required
	}

	if len(_sagemakerArtifactStoreUri) > 0 {
		input.ArtifactStoreUri = aws.String(_sagemakerArtifactStoreUri)
	}
	if len(_sagemakerRoleArn) > 0 {
		input.RoleArn = aws.String(_sagemakerRoleArn)
	}
	if len(_sagemakerTrackingServerName) > 0 {
		input.TrackingServerName = aws.String(_sagemakerTrackingServerName)
	}
	if len(_sagemakerAutomaticModelRegistration) > 0 {
		if err := assignInputField(input, "AutomaticModelRegistration", _sagemakerAutomaticModelRegistration); err != nil {
			log.Errorf("invalid --automatic-model-registration: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMlflowVersion) > 0 {
		input.MlflowVersion = aws.String(_sagemakerMlflowVersion)
	}
	if len(_sagemakerTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_sagemakerTrackingServerSize) > 0 {
		if err := assignInputField(input, "TrackingServerSize", _sagemakerTrackingServerSize); err != nil {
			log.Errorf("invalid --tracking-server-size: %s", err.Error())
			return
		}
	}
	if len(_sagemakerWeeklyMaintenanceWindowStart) > 0 {
		input.WeeklyMaintenanceWindowStart = aws.String(_sagemakerWeeklyMaintenanceWindowStart)
	}

	if resp, err := client.CreateMlflowTrackingServer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a model in SageMaker. In the request, you name the model and describe a
// primary container. For the primary container, you specify the Docker image that
// contains inference code, artifacts (from prior training), and a custom
// environment map that the inference code uses when you deploy the model for
// predictions.
//
// Use this API to create a model if you want to use SageMaker hosting services or
// run a batch transform job.
//
// To host your model, you create an endpoint configuration with the
// CreateEndpointConfig API, and then create an endpoint with the CreateEndpoint
// API. SageMaker then deploys all of the containers that you defined for the model
// in the hosting environment.
//
// To run a batch transform using your model, you start a job with the
// CreateTransformJob API. SageMaker uses your model and your dataset to get
// inferences which are then saved to a specified S3 location.
//
// In the request, you also provide an IAM role that SageMaker can assume to
// access model artifacts and docker image for deployment on ML compute hosting
// instances or for batch transform jobs. In addition, you also use the IAM role to
// manage permissions the inference code needs. For example, if the inference code
// access any other Amazon Web Services resources, you grant necessary permissions
// via this role.
func sagemaker_CreateModel(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreateModelInput{
		// ModelName: *string, // Required
	}

	if len(_sagemakerModelName) > 0 {
		input.ModelName = aws.String(_sagemakerModelName)
	}
	if len(_sagemakerContainers) > 0 {
		if err := assignInputField(input, "Containers", _sagemakerContainers); err != nil {
			log.Errorf("invalid --containers: %s", err.Error())
			return
		}
	}
	if len(_sagemakerEnableNetworkIsolation) > 0 {
		if err := assignInputField(input, "EnableNetworkIsolation", _sagemakerEnableNetworkIsolation); err != nil {
			log.Errorf("invalid --enable-network-isolation: %s", err.Error())
			return
		}
	}
	if len(_sagemakerExecutionRoleArn) > 0 {
		input.ExecutionRoleArn = aws.String(_sagemakerExecutionRoleArn)
	}
	if len(_sagemakerInferenceExecutionConfig) > 0 {
		if err := assignInputField(input, "InferenceExecutionConfig", _sagemakerInferenceExecutionConfig); err != nil {
			log.Errorf("invalid --inference-execution-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerPrimaryContainer) > 0 {
		if err := assignInputField(input, "PrimaryContainer", _sagemakerPrimaryContainer); err != nil {
			log.Errorf("invalid --primary-container: %s", err.Error())
			return
		}
	}
	if len(_sagemakerTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_sagemakerVpcConfig) > 0 {
		if err := assignInputField(input, "VpcConfig", _sagemakerVpcConfig); err != nil {
			log.Errorf("invalid --vpc-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates the definition for a model bias job.
func sagemaker_CreateModelBiasJobDefinition(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreateModelBiasJobDefinitionInput{
		// JobDefinitionName: *string, // Required
		// JobResources: *types.MonitoringResources, // Required
		// ModelBiasAppSpecification: *types.ModelBiasAppSpecification, // Required
		// ModelBiasJobInput: *types.ModelBiasJobInput, // Required
		// ModelBiasJobOutputConfig: *types.MonitoringOutputConfig, // Required
		// RoleArn: *string, // Required
	}

	if len(_sagemakerJobDefinitionName) > 0 {
		input.JobDefinitionName = aws.String(_sagemakerJobDefinitionName)
	}
	if len(_sagemakerJobResources) > 0 {
		if err := assignInputField(input, "JobResources", _sagemakerJobResources); err != nil {
			log.Errorf("invalid --job-resources: %s", err.Error())
			return
		}
	}
	if len(_sagemakerModelBiasAppSpecification) > 0 {
		if err := assignInputField(input, "ModelBiasAppSpecification", _sagemakerModelBiasAppSpecification); err != nil {
			log.Errorf("invalid --model-bias-app-specification: %s", err.Error())
			return
		}
	}
	if len(_sagemakerModelBiasJobInput) > 0 {
		if err := assignInputField(input, "ModelBiasJobInput", _sagemakerModelBiasJobInput); err != nil {
			log.Errorf("invalid --model-bias-job-input: %s", err.Error())
			return
		}
	}
	if len(_sagemakerModelBiasJobOutputConfig) > 0 {
		if err := assignInputField(input, "ModelBiasJobOutputConfig", _sagemakerModelBiasJobOutputConfig); err != nil {
			log.Errorf("invalid --model-bias-job-output-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerRoleArn) > 0 {
		input.RoleArn = aws.String(_sagemakerRoleArn)
	}
	if len(_sagemakerModelBiasBaselineConfig) > 0 {
		if err := assignInputField(input, "ModelBiasBaselineConfig", _sagemakerModelBiasBaselineConfig); err != nil {
			log.Errorf("invalid --model-bias-baseline-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNetworkConfig) > 0 {
		if err := assignInputField(input, "NetworkConfig", _sagemakerNetworkConfig); err != nil {
			log.Errorf("invalid --network-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerStoppingCondition) > 0 {
		if err := assignInputField(input, "StoppingCondition", _sagemakerStoppingCondition); err != nil {
			log.Errorf("invalid --stopping-condition: %s", err.Error())
			return
		}
	}
	if len(_sagemakerTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateModelBiasJobDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon SageMaker Model Card.
// For information about how to use model cards, see [Amazon SageMaker Model Card].
//
// [Amazon SageMaker Model Card]: https://docs.aws.amazon.com/sagemaker/latest/dg/model-cards.html
func sagemaker_CreateModelCard(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreateModelCardInput{
		// Content: *string, // Required
		// ModelCardName: *string, // Required
		// ModelCardStatus: types.ModelCardStatus, // Required
	}

	if len(_sagemakerContent) > 0 {
		input.Content = aws.String(_sagemakerContent)
	}
	if len(_sagemakerModelCardName) > 0 {
		input.ModelCardName = aws.String(_sagemakerModelCardName)
	}
	if len(_sagemakerModelCardStatus) > 0 {
		if err := assignInputField(input, "ModelCardStatus", _sagemakerModelCardStatus); err != nil {
			log.Errorf("invalid --model-card-status: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSecurityConfig) > 0 {
		if err := assignInputField(input, "SecurityConfig", _sagemakerSecurityConfig); err != nil {
			log.Errorf("invalid --security-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateModelCard(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon SageMaker Model Card export job.
func sagemaker_CreateModelCardExportJob(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreateModelCardExportJobInput{
		// ModelCardExportJobName: *string, // Required
		// ModelCardName: *string, // Required
		// OutputConfig: *types.ModelCardExportOutputConfig, // Required
	}

	if len(_sagemakerModelCardExportJobName) > 0 {
		input.ModelCardExportJobName = aws.String(_sagemakerModelCardExportJobName)
	}
	if len(_sagemakerModelCardName) > 0 {
		input.ModelCardName = aws.String(_sagemakerModelCardName)
	}
	if len(_sagemakerOutputConfig) > 0 {
		if err := assignInputField(input, "OutputConfig", _sagemakerOutputConfig); err != nil {
			log.Errorf("invalid --output-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerModelCardVersion) > 0 {
		if err := assignInputField(input, "ModelCardVersion", _sagemakerModelCardVersion); err != nil {
			log.Errorf("invalid --model-card-version: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateModelCardExportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates the definition for a model explainability job.
func sagemaker_CreateModelExplainabilityJobDefinition(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreateModelExplainabilityJobDefinitionInput{
		// JobDefinitionName: *string, // Required
		// JobResources: *types.MonitoringResources, // Required
		// ModelExplainabilityAppSpecification: *types.ModelExplainabilityAppSpecification, // Required
		// ModelExplainabilityJobInput: *types.ModelExplainabilityJobInput, // Required
		// ModelExplainabilityJobOutputConfig: *types.MonitoringOutputConfig, // Required
		// RoleArn: *string, // Required
	}

	if len(_sagemakerJobDefinitionName) > 0 {
		input.JobDefinitionName = aws.String(_sagemakerJobDefinitionName)
	}
	if len(_sagemakerJobResources) > 0 {
		if err := assignInputField(input, "JobResources", _sagemakerJobResources); err != nil {
			log.Errorf("invalid --job-resources: %s", err.Error())
			return
		}
	}
	if len(_sagemakerModelExplainabilityAppSpecification) > 0 {
		if err := assignInputField(input, "ModelExplainabilityAppSpecification", _sagemakerModelExplainabilityAppSpecification); err != nil {
			log.Errorf("invalid --model-explainability-app-specification: %s", err.Error())
			return
		}
	}
	if len(_sagemakerModelExplainabilityJobInput) > 0 {
		if err := assignInputField(input, "ModelExplainabilityJobInput", _sagemakerModelExplainabilityJobInput); err != nil {
			log.Errorf("invalid --model-explainability-job-input: %s", err.Error())
			return
		}
	}
	if len(_sagemakerModelExplainabilityJobOutputConfig) > 0 {
		if err := assignInputField(input, "ModelExplainabilityJobOutputConfig", _sagemakerModelExplainabilityJobOutputConfig); err != nil {
			log.Errorf("invalid --model-explainability-job-output-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerRoleArn) > 0 {
		input.RoleArn = aws.String(_sagemakerRoleArn)
	}
	if len(_sagemakerModelExplainabilityBaselineConfig) > 0 {
		if err := assignInputField(input, "ModelExplainabilityBaselineConfig", _sagemakerModelExplainabilityBaselineConfig); err != nil {
			log.Errorf("invalid --model-explainability-baseline-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNetworkConfig) > 0 {
		if err := assignInputField(input, "NetworkConfig", _sagemakerNetworkConfig); err != nil {
			log.Errorf("invalid --network-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerStoppingCondition) > 0 {
		if err := assignInputField(input, "StoppingCondition", _sagemakerStoppingCondition); err != nil {
			log.Errorf("invalid --stopping-condition: %s", err.Error())
			return
		}
	}
	if len(_sagemakerTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateModelExplainabilityJobDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a model package that you can use to create SageMaker models or list on
// Amazon Web Services Marketplace, or a versioned model that is part of a model
// group. Buyers can subscribe to model packages listed on Amazon Web Services
// Marketplace to create models in SageMaker.
//
// To create a model package by specifying a Docker container that contains your
// inference code and the Amazon S3 location of your model artifacts, provide
// values for InferenceSpecification . To create a model from an algorithm resource
// that you created or subscribed to in Amazon Web Services Marketplace, provide a
// value for SourceAlgorithmSpecification .
//
// There are two types of model packages:
//
// - Versioned - a model that is part of a model group in the model registry.
//
// - Unversioned - a model package that is not part of a model group.
func sagemaker_CreateModelPackage(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreateModelPackageInput{}

	if len(_sagemakerAdditionalInferenceSpecifications) > 0 {
		if err := assignInputField(input, "AdditionalInferenceSpecifications", _sagemakerAdditionalInferenceSpecifications); err != nil {
			log.Errorf("invalid --additional-inference-specifications: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCertifyForMarketplace) > 0 {
		if err := assignInputField(input, "CertifyForMarketplace", _sagemakerCertifyForMarketplace); err != nil {
			log.Errorf("invalid --certify-for-marketplace: %s", err.Error())
			return
		}
	}
	if len(_sagemakerClientToken) > 0 {
		input.ClientToken = aws.String(_sagemakerClientToken)
	}
	if len(_sagemakerCustomerMetadataProperties) > 0 {
		if err := assignInputField(input, "CustomerMetadataProperties", _sagemakerCustomerMetadataProperties); err != nil {
			log.Errorf("invalid --customer-metadata-properties: %s", err.Error())
			return
		}
	}
	if len(_sagemakerDomain) > 0 {
		input.Domain = aws.String(_sagemakerDomain)
	}
	if len(_sagemakerDriftCheckBaselines) > 0 {
		if err := assignInputField(input, "DriftCheckBaselines", _sagemakerDriftCheckBaselines); err != nil {
			log.Errorf("invalid --drift-check-baselines: %s", err.Error())
			return
		}
	}
	if len(_sagemakerInferenceSpecification) > 0 {
		if err := assignInputField(input, "InferenceSpecification", _sagemakerInferenceSpecification); err != nil {
			log.Errorf("invalid --inference-specification: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMetadataProperties) > 0 {
		if err := assignInputField(input, "MetadataProperties", _sagemakerMetadataProperties); err != nil {
			log.Errorf("invalid --metadata-properties: %s", err.Error())
			return
		}
	}
	if len(_sagemakerModelApprovalStatus) > 0 {
		if err := assignInputField(input, "ModelApprovalStatus", _sagemakerModelApprovalStatus); err != nil {
			log.Errorf("invalid --model-approval-status: %s", err.Error())
			return
		}
	}
	if len(_sagemakerModelCard) > 0 {
		if err := assignInputField(input, "ModelCard", _sagemakerModelCard); err != nil {
			log.Errorf("invalid --model-card: %s", err.Error())
			return
		}
	}
	if len(_sagemakerModelLifeCycle) > 0 {
		if err := assignInputField(input, "ModelLifeCycle", _sagemakerModelLifeCycle); err != nil {
			log.Errorf("invalid --model-life-cycle: %s", err.Error())
			return
		}
	}
	if len(_sagemakerModelMetrics) > 0 {
		if err := assignInputField(input, "ModelMetrics", _sagemakerModelMetrics); err != nil {
			log.Errorf("invalid --model-metrics: %s", err.Error())
			return
		}
	}
	if len(_sagemakerModelPackageDescription) > 0 {
		input.ModelPackageDescription = aws.String(_sagemakerModelPackageDescription)
	}
	if len(_sagemakerModelPackageGroupName) > 0 {
		input.ModelPackageGroupName = aws.String(_sagemakerModelPackageGroupName)
	}
	if len(_sagemakerModelPackageName) > 0 {
		input.ModelPackageName = aws.String(_sagemakerModelPackageName)
	}
	if len(_sagemakerModelPackageRegistrationType) > 0 {
		if err := assignInputField(input, "ModelPackageRegistrationType", _sagemakerModelPackageRegistrationType); err != nil {
			log.Errorf("invalid --model-package-registration-type: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSamplePayloadUrl) > 0 {
		input.SamplePayloadUrl = aws.String(_sagemakerSamplePayloadUrl)
	}
	if len(_sagemakerSecurityConfig) > 0 {
		if err := assignInputField(input, "SecurityConfig", _sagemakerSecurityConfig); err != nil {
			log.Errorf("invalid --security-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSkipModelValidation) > 0 {
		if err := assignInputField(input, "SkipModelValidation", _sagemakerSkipModelValidation); err != nil {
			log.Errorf("invalid --skip-model-validation: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSourceAlgorithmSpecification) > 0 {
		if err := assignInputField(input, "SourceAlgorithmSpecification", _sagemakerSourceAlgorithmSpecification); err != nil {
			log.Errorf("invalid --source-algorithm-specification: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSourceUri) > 0 {
		input.SourceUri = aws.String(_sagemakerSourceUri)
	}
	if len(_sagemakerTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_sagemakerTask) > 0 {
		input.Task = aws.String(_sagemakerTask)
	}
	if len(_sagemakerValidationSpecification) > 0 {
		if err := assignInputField(input, "ValidationSpecification", _sagemakerValidationSpecification); err != nil {
			log.Errorf("invalid --validation-specification: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateModelPackage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a model group. A model group contains a group of model versions.
func sagemaker_CreateModelPackageGroup(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreateModelPackageGroupInput{
		// ModelPackageGroupName: *string, // Required
	}

	if len(_sagemakerModelPackageGroupName) > 0 {
		input.ModelPackageGroupName = aws.String(_sagemakerModelPackageGroupName)
	}
	if len(_sagemakerModelPackageGroupDescription) > 0 {
		input.ModelPackageGroupDescription = aws.String(_sagemakerModelPackageGroupDescription)
	}
	if len(_sagemakerTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateModelPackageGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a definition for a job that monitors model quality and drift. For
// information about model monitor, see [Amazon SageMaker AI Model Monitor].
//
// [Amazon SageMaker AI Model Monitor]: https://docs.aws.amazon.com/sagemaker/latest/dg/model-monitor.html
func sagemaker_CreateModelQualityJobDefinition(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreateModelQualityJobDefinitionInput{
		// JobDefinitionName: *string, // Required
		// JobResources: *types.MonitoringResources, // Required
		// ModelQualityAppSpecification: *types.ModelQualityAppSpecification, // Required
		// ModelQualityJobInput: *types.ModelQualityJobInput, // Required
		// ModelQualityJobOutputConfig: *types.MonitoringOutputConfig, // Required
		// RoleArn: *string, // Required
	}

	if len(_sagemakerJobDefinitionName) > 0 {
		input.JobDefinitionName = aws.String(_sagemakerJobDefinitionName)
	}
	if len(_sagemakerJobResources) > 0 {
		if err := assignInputField(input, "JobResources", _sagemakerJobResources); err != nil {
			log.Errorf("invalid --job-resources: %s", err.Error())
			return
		}
	}
	if len(_sagemakerModelQualityAppSpecification) > 0 {
		if err := assignInputField(input, "ModelQualityAppSpecification", _sagemakerModelQualityAppSpecification); err != nil {
			log.Errorf("invalid --model-quality-app-specification: %s", err.Error())
			return
		}
	}
	if len(_sagemakerModelQualityJobInput) > 0 {
		if err := assignInputField(input, "ModelQualityJobInput", _sagemakerModelQualityJobInput); err != nil {
			log.Errorf("invalid --model-quality-job-input: %s", err.Error())
			return
		}
	}
	if len(_sagemakerModelQualityJobOutputConfig) > 0 {
		if err := assignInputField(input, "ModelQualityJobOutputConfig", _sagemakerModelQualityJobOutputConfig); err != nil {
			log.Errorf("invalid --model-quality-job-output-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerRoleArn) > 0 {
		input.RoleArn = aws.String(_sagemakerRoleArn)
	}
	if len(_sagemakerModelQualityBaselineConfig) > 0 {
		if err := assignInputField(input, "ModelQualityBaselineConfig", _sagemakerModelQualityBaselineConfig); err != nil {
			log.Errorf("invalid --model-quality-baseline-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNetworkConfig) > 0 {
		if err := assignInputField(input, "NetworkConfig", _sagemakerNetworkConfig); err != nil {
			log.Errorf("invalid --network-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerStoppingCondition) > 0 {
		if err := assignInputField(input, "StoppingCondition", _sagemakerStoppingCondition); err != nil {
			log.Errorf("invalid --stopping-condition: %s", err.Error())
			return
		}
	}
	if len(_sagemakerTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateModelQualityJobDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a schedule that regularly starts Amazon SageMaker AI Processing Jobs to
// monitor the data captured for an Amazon SageMaker AI Endpoint.
func sagemaker_CreateMonitoringSchedule(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreateMonitoringScheduleInput{
		// MonitoringScheduleConfig: *types.MonitoringScheduleConfig, // Required
		// MonitoringScheduleName: *string, // Required
	}

	if len(_sagemakerMonitoringScheduleConfig) > 0 {
		if err := assignInputField(input, "MonitoringScheduleConfig", _sagemakerMonitoringScheduleConfig); err != nil {
			log.Errorf("invalid --monitoring-schedule-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMonitoringScheduleName) > 0 {
		input.MonitoringScheduleName = aws.String(_sagemakerMonitoringScheduleName)
	}
	if len(_sagemakerTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateMonitoringSchedule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an SageMaker AI notebook instance. A notebook instance is a machine
// learning (ML) compute instance running on a Jupyter notebook.
//
// In a CreateNotebookInstance request, specify the type of ML compute instance
// that you want to run. SageMaker AI launches the instance, installs common
// libraries that you can use to explore datasets for model training, and attaches
// an ML storage volume to the notebook instance.
//
// SageMaker AI also provides a set of example notebooks. Each notebook
// demonstrates how to use SageMaker AI with a specific algorithm or with a machine
// learning framework.
//
// After receiving the request, SageMaker AI does the following:
//
// - Creates a network interface in the SageMaker AI VPC.
//
// - (Option) If you specified SubnetId , SageMaker AI creates a network
// interface in your own VPC, which is inferred from the subnet ID that you provide
// in the input. When creating this network interface, SageMaker AI attaches the
// security group that you specified in the request to the network interface that
// it creates in your VPC.
//
// - Launches an EC2 instance of the type specified in the request in the
// SageMaker AI VPC. If you specified SubnetId of your VPC, SageMaker AI
// specifies both network interfaces when launching this instance. This enables
// inbound traffic from your own VPC to the notebook instance, assuming that the
// security groups allow it.
//
// After creating the notebook instance, SageMaker AI returns its Amazon Resource
// Name (ARN). You can't change the name of a notebook instance after you create
// it.
//
// After SageMaker AI creates the notebook instance, you can connect to the
// Jupyter server and work in Jupyter notebooks. For example, you can write code to
// explore a dataset that you can use for model training, train a model, host
// models by creating SageMaker AI endpoints, and validate hosted models.
//
// For more information, see [How It Works].
//
// [How It Works]: https://docs.aws.amazon.com/sagemaker/latest/dg/how-it-works.html
func sagemaker_CreateNotebookInstance(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreateNotebookInstanceInput{
		// InstanceType: types.InstanceType, // Required
		// NotebookInstanceName: *string, // Required
		// RoleArn: *string, // Required
	}

	if len(_sagemakerInstanceType) > 0 {
		if err := assignInputField(input, "InstanceType", _sagemakerInstanceType); err != nil {
			log.Errorf("invalid --instance-type: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNotebookInstanceName) > 0 {
		input.NotebookInstanceName = aws.String(_sagemakerNotebookInstanceName)
	}
	if len(_sagemakerRoleArn) > 0 {
		input.RoleArn = aws.String(_sagemakerRoleArn)
	}
	if len(_sagemakerAcceleratorTypes) > 0 {
		if err := assignInputField(input, "AcceleratorTypes", _sagemakerAcceleratorTypes); err != nil {
			log.Errorf("invalid --accelerator-types: %s", err.Error())
			return
		}
	}
	if len(_sagemakerAdditionalCodeRepositories) > 0 {
		input.AdditionalCodeRepositories = append([]string(nil), _sagemakerAdditionalCodeRepositories...)
	}
	if len(_sagemakerDefaultCodeRepository) > 0 {
		input.DefaultCodeRepository = aws.String(_sagemakerDefaultCodeRepository)
	}
	if len(_sagemakerDirectInternetAccess) > 0 {
		if err := assignInputField(input, "DirectInternetAccess", _sagemakerDirectInternetAccess); err != nil {
			log.Errorf("invalid --direct-internet-access: %s", err.Error())
			return
		}
	}
	if len(_sagemakerInstanceMetadataServiceConfiguration) > 0 {
		if err := assignInputField(input, "InstanceMetadataServiceConfiguration", _sagemakerInstanceMetadataServiceConfiguration); err != nil {
			log.Errorf("invalid --instance-metadata-service-configuration: %s", err.Error())
			return
		}
	}
	if len(_sagemakerIpAddressType) > 0 {
		if err := assignInputField(input, "IpAddressType", _sagemakerIpAddressType); err != nil {
			log.Errorf("invalid --ip-address-type: %s", err.Error())
			return
		}
	}
	if len(_sagemakerKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_sagemakerKmsKeyId)
	}
	if len(_sagemakerLifecycleConfigName) > 0 {
		input.LifecycleConfigName = aws.String(_sagemakerLifecycleConfigName)
	}
	if len(_sagemakerPlatformIdentifier) > 0 {
		input.PlatformIdentifier = aws.String(_sagemakerPlatformIdentifier)
	}
	if len(_sagemakerRootAccess) > 0 {
		if err := assignInputField(input, "RootAccess", _sagemakerRootAccess); err != nil {
			log.Errorf("invalid --root-access: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSecurityGroupIds) > 0 {
		input.SecurityGroupIds = append([]string(nil), _sagemakerSecurityGroupIds...)
	}
	if len(_sagemakerSubnetId) > 0 {
		input.SubnetId = aws.String(_sagemakerSubnetId)
	}
	if len(_sagemakerTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_sagemakerVolumeSizeInGB) > 0 {
		if err := assignInputField(input, "VolumeSizeInGB", _sagemakerVolumeSizeInGB); err != nil {
			log.Errorf("invalid --volume-size-in-gb: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateNotebookInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a lifecycle configuration that you can associate with a notebook
// instance. A lifecycle configuration is a collection of shell scripts that run
// when you create or start a notebook instance.
//
// Each lifecycle configuration script has a limit of 16384 characters.
//
// The value of the $PATH environment variable that is available to both scripts
// is /sbin:bin:/usr/sbin:/usr/bin .
//
// View Amazon CloudWatch Logs for notebook instance lifecycle configurations in
// log group /aws/sagemaker/NotebookInstances in log stream
// [notebook-instance-name]/[LifecycleConfigHook] .
//
// Lifecycle configuration scripts cannot run for longer than 5 minutes. If a
// script runs for longer than 5 minutes, it fails and the notebook instance is not
// created or started.
//
// For information about notebook instance lifestyle configurations, see [Step 2.1: (Optional) Customize a Notebook Instance].
//
// Lifecycle configuration scripts execute with root access and the notebook
// instance's IAM execution role privileges. Grant this permission only to trusted
// principals. See [Customize a Notebook Instance Using a Lifecycle Configuration Script]for security best practices.
//
// [Customize a Notebook Instance Using a Lifecycle Configuration Script]: https://docs.aws.amazon.com/sagemaker/latest/dg/notebook-lifecycle-config.html
// [Step 2.1: (Optional) Customize a Notebook Instance]: https://docs.aws.amazon.com/sagemaker/latest/dg/notebook-lifecycle-config.html
func sagemaker_CreateNotebookInstanceLifecycleConfig(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreateNotebookInstanceLifecycleConfigInput{
		// NotebookInstanceLifecycleConfigName: *string, // Required
	}

	if len(_sagemakerNotebookInstanceLifecycleConfigName) > 0 {
		input.NotebookInstanceLifecycleConfigName = aws.String(_sagemakerNotebookInstanceLifecycleConfigName)
	}
	if len(_sagemakerOnCreate) > 0 {
		if err := assignInputField(input, "OnCreate", _sagemakerOnCreate); err != nil {
			log.Errorf("invalid --on-create: %s", err.Error())
			return
		}
	}
	if len(_sagemakerOnStart) > 0 {
		if err := assignInputField(input, "OnStart", _sagemakerOnStart); err != nil {
			log.Errorf("invalid --on-start: %s", err.Error())
			return
		}
	}
	if len(_sagemakerTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateNotebookInstanceLifecycleConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a job that optimizes a model for inference performance. To create the
// job, you provide the location of a source model, and you provide the settings
// for the optimization techniques that you want the job to apply. When the job
// completes successfully, SageMaker uploads the new optimized model to the output
// destination that you specify.
//
// For more information about how to use this action, and about the supported
// optimization techniques, see [Optimize model inference with Amazon SageMaker].
//
// [Optimize model inference with Amazon SageMaker]: https://docs.aws.amazon.com/sagemaker/latest/dg/model-optimize.html
func sagemaker_CreateOptimizationJob(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreateOptimizationJobInput{
		// DeploymentInstanceType: types.OptimizationJobDeploymentInstanceType, // Required
		// ModelSource: *types.OptimizationJobModelSource, // Required
		// OptimizationConfigs: []types.OptimizationConfig, // Required
		// OptimizationJobName: *string, // Required
		// OutputConfig: *types.OptimizationJobOutputConfig, // Required
		// RoleArn: *string, // Required
		// StoppingCondition: *types.StoppingCondition, // Required
	}

	if len(_sagemakerDeploymentInstanceType) > 0 {
		if err := assignInputField(input, "DeploymentInstanceType", _sagemakerDeploymentInstanceType); err != nil {
			log.Errorf("invalid --deployment-instance-type: %s", err.Error())
			return
		}
	}
	if len(_sagemakerModelSource) > 0 {
		if err := assignInputField(input, "ModelSource", _sagemakerModelSource); err != nil {
			log.Errorf("invalid --model-source: %s", err.Error())
			return
		}
	}
	if len(_sagemakerOptimizationConfigs) > 0 {
		if err := assignInputField(input, "OptimizationConfigs", _sagemakerOptimizationConfigs); err != nil {
			log.Errorf("invalid --optimization-configs: %s", err.Error())
			return
		}
	}
	if len(_sagemakerOptimizationJobName) > 0 {
		input.OptimizationJobName = aws.String(_sagemakerOptimizationJobName)
	}
	if len(_sagemakerOutputConfig) > 0 {
		if err := assignInputField(input, "OutputConfig", _sagemakerOutputConfig); err != nil {
			log.Errorf("invalid --output-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerRoleArn) > 0 {
		input.RoleArn = aws.String(_sagemakerRoleArn)
	}
	if len(_sagemakerStoppingCondition) > 0 {
		if err := assignInputField(input, "StoppingCondition", _sagemakerStoppingCondition); err != nil {
			log.Errorf("invalid --stopping-condition: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMaxInstanceCount) > 0 {
		if err := assignInputField(input, "MaxInstanceCount", _sagemakerMaxInstanceCount); err != nil {
			log.Errorf("invalid --max-instance-count: %s", err.Error())
			return
		}
	}
	if len(_sagemakerOptimizationEnvironment) > 0 {
		if err := assignInputField(input, "OptimizationEnvironment", _sagemakerOptimizationEnvironment); err != nil {
			log.Errorf("invalid --optimization-environment: %s", err.Error())
			return
		}
	}
	if len(_sagemakerTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_sagemakerVpcConfig) > 0 {
		if err := assignInputField(input, "VpcConfig", _sagemakerVpcConfig); err != nil {
			log.Errorf("invalid --vpc-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateOptimizationJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon SageMaker Partner AI App.
func sagemaker_CreatePartnerApp(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreatePartnerAppInput{
		// AuthType: types.PartnerAppAuthType, // Required
		// ExecutionRoleArn: *string, // Required
		// Name: *string, // Required
		// Tier: *string, // Required
		// Type: types.PartnerAppType, // Required
	}

	if len(_sagemakerAuthType) > 0 {
		if err := assignInputField(input, "AuthType", _sagemakerAuthType); err != nil {
			log.Errorf("invalid --auth-type: %s", err.Error())
			return
		}
	}
	if len(_sagemakerExecutionRoleArn) > 0 {
		input.ExecutionRoleArn = aws.String(_sagemakerExecutionRoleArn)
	}
	if len(_sagemakerName) > 0 {
		input.Name = aws.String(_sagemakerName)
	}
	if len(_sagemakerTier) > 0 {
		input.Tier = aws.String(_sagemakerTier)
	}
	if len(_sagemakerType) > 0 {
		if err := assignInputField(input, "Type", _sagemakerType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_sagemakerApplicationConfig) > 0 {
		if err := assignInputField(input, "ApplicationConfig", _sagemakerApplicationConfig); err != nil {
			log.Errorf("invalid --application-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerClientToken) > 0 {
		input.ClientToken = aws.String(_sagemakerClientToken)
	}
	if len(_sagemakerEnableAutoMinorVersionUpgrade) > 0 {
		if err := assignInputField(input, "EnableAutoMinorVersionUpgrade", _sagemakerEnableAutoMinorVersionUpgrade); err != nil {
			log.Errorf("invalid --enable-auto-minor-version-upgrade: %s", err.Error())
			return
		}
	}
	if len(_sagemakerEnableIamSessionBasedIdentity) > 0 {
		if err := assignInputField(input, "EnableIamSessionBasedIdentity", _sagemakerEnableIamSessionBasedIdentity); err != nil {
			log.Errorf("invalid --enable-iam-session-based-identity: %s", err.Error())
			return
		}
	}
	if len(_sagemakerKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_sagemakerKmsKeyId)
	}
	if len(_sagemakerMaintenanceConfig) > 0 {
		if err := assignInputField(input, "MaintenanceConfig", _sagemakerMaintenanceConfig); err != nil {
			log.Errorf("invalid --maintenance-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreatePartnerApp(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a presigned URL to access an Amazon SageMaker Partner AI App.
func sagemaker_CreatePartnerAppPresignedUrl(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreatePartnerAppPresignedUrlInput{
		// Arn: *string, // Required
	}

	if len(_sagemakerArn) > 0 {
		input.Arn = aws.String(_sagemakerArn)
	}
	if len(_sagemakerExpiresInSeconds) > 0 {
		if err := assignInputField(input, "ExpiresInSeconds", _sagemakerExpiresInSeconds); err != nil {
			log.Errorf("invalid --expires-in-seconds: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSessionExpirationDurationInSeconds) > 0 {
		if err := assignInputField(input, "SessionExpirationDurationInSeconds", _sagemakerSessionExpirationDurationInSeconds); err != nil {
			log.Errorf("invalid --session-expiration-duration-in-seconds: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreatePartnerAppPresignedUrl(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a pipeline using a JSON pipeline definition.
func sagemaker_CreatePipeline(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreatePipelineInput{
		// ClientRequestToken: *string, // Required
		// PipelineName: *string, // Required
		// RoleArn: *string, // Required
	}

	if len(_sagemakerClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_sagemakerClientRequestToken)
	}
	if len(_sagemakerPipelineName) > 0 {
		input.PipelineName = aws.String(_sagemakerPipelineName)
	}
	if len(_sagemakerRoleArn) > 0 {
		input.RoleArn = aws.String(_sagemakerRoleArn)
	}
	if len(_sagemakerParallelismConfiguration) > 0 {
		if err := assignInputField(input, "ParallelismConfiguration", _sagemakerParallelismConfiguration); err != nil {
			log.Errorf("invalid --parallelism-configuration: %s", err.Error())
			return
		}
	}
	if len(_sagemakerPipelineDefinition) > 0 {
		input.PipelineDefinition = aws.String(_sagemakerPipelineDefinition)
	}
	if len(_sagemakerPipelineDefinitionS3Location) > 0 {
		if err := assignInputField(input, "PipelineDefinitionS3Location", _sagemakerPipelineDefinitionS3Location); err != nil {
			log.Errorf("invalid --pipeline-definition-s3-location: %s", err.Error())
			return
		}
	}
	if len(_sagemakerPipelineDescription) > 0 {
		input.PipelineDescription = aws.String(_sagemakerPipelineDescription)
	}
	if len(_sagemakerPipelineDisplayName) > 0 {
		input.PipelineDisplayName = aws.String(_sagemakerPipelineDisplayName)
	}
	if len(_sagemakerTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakerTags); err != nil {
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

// Creates a URL for a specified UserProfile in a Domain. When accessed in a web
// browser, the user will be automatically signed in to the domain, and granted
// access to all of the Apps and files associated with the Domain's Amazon Elastic
// File System volume. This operation can only be called when the authentication
// mode equals IAM.
//
// The IAM role or user passed to this API defines the permissions to access the
// app. Once the presigned URL is created, no additional permission is required to
// access this URL. IAM authorization policies for this API are also enforced for
// every HTTP request and WebSocket frame that attempts to connect to the app.
//
// You can restrict access to this API and to the URL that it returns to a list of
// IP addresses, Amazon VPCs or Amazon VPC Endpoints that you specify. For more
// information, see [Connect to Amazon SageMaker AI Studio Through an Interface VPC Endpoint].
//
// - The URL that you get from a call to CreatePresignedDomainUrl has a default
// timeout of 5 minutes. You can configure this value using ExpiresInSeconds . If
// you try to use the URL after the timeout limit expires, you are directed to the
// Amazon Web Services console sign-in page.
//
// - The JupyterLab session default expiration time is 12 hours. You can
// configure this value using SessionExpirationDurationInSeconds.
//
// [Connect to Amazon SageMaker AI Studio Through an Interface VPC Endpoint]: https://docs.aws.amazon.com/sagemaker/latest/dg/studio-interface-endpoint.html
func sagemaker_CreatePresignedDomainUrl(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreatePresignedDomainUrlInput{
		// DomainId: *string, // Required
		// UserProfileName: *string, // Required
	}

	if len(_sagemakerDomainId) > 0 {
		input.DomainId = aws.String(_sagemakerDomainId)
	}
	if len(_sagemakerUserProfileName) > 0 {
		input.UserProfileName = aws.String(_sagemakerUserProfileName)
	}
	if len(_sagemakerExpiresInSeconds) > 0 {
		if err := assignInputField(input, "ExpiresInSeconds", _sagemakerExpiresInSeconds); err != nil {
			log.Errorf("invalid --expires-in-seconds: %s", err.Error())
			return
		}
	}
	if len(_sagemakerLandingUri) > 0 {
		input.LandingUri = aws.String(_sagemakerLandingUri)
	}
	if len(_sagemakerSessionExpirationDurationInSeconds) > 0 {
		if err := assignInputField(input, "SessionExpirationDurationInSeconds", _sagemakerSessionExpirationDurationInSeconds); err != nil {
			log.Errorf("invalid --session-expiration-duration-in-seconds: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSpaceName) > 0 {
		input.SpaceName = aws.String(_sagemakerSpaceName)
	}

	if resp, err := client.CreatePresignedDomainUrl(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a presigned URL that you can use to connect to the MLflow UI attached
// to your MLflow App. For more information, see [Launch the MLflow UI using a presigned URL].
//
// [Launch the MLflow UI using a presigned URL]: https://docs.aws.amazon.com/sagemaker/latest/dg/mlflow-launch-ui.html
func sagemaker_CreatePresignedMlflowAppUrl(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreatePresignedMlflowAppUrlInput{
		// Arn: *string, // Required
	}

	if len(_sagemakerArn) > 0 {
		input.Arn = aws.String(_sagemakerArn)
	}
	if len(_sagemakerExpiresInSeconds) > 0 {
		if err := assignInputField(input, "ExpiresInSeconds", _sagemakerExpiresInSeconds); err != nil {
			log.Errorf("invalid --expires-in-seconds: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSessionExpirationDurationInSeconds) > 0 {
		if err := assignInputField(input, "SessionExpirationDurationInSeconds", _sagemakerSessionExpirationDurationInSeconds); err != nil {
			log.Errorf("invalid --session-expiration-duration-in-seconds: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreatePresignedMlflowAppUrl(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a presigned URL that you can use to connect to the MLflow UI attached
// to your tracking server. For more information, see [Launch the MLflow UI using a presigned URL].
//
// [Launch the MLflow UI using a presigned URL]: https://docs.aws.amazon.com/sagemaker/latest/dg/mlflow-launch-ui.html
func sagemaker_CreatePresignedMlflowTrackingServerUrl(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreatePresignedMlflowTrackingServerUrlInput{
		// TrackingServerName: *string, // Required
	}

	if len(_sagemakerTrackingServerName) > 0 {
		input.TrackingServerName = aws.String(_sagemakerTrackingServerName)
	}
	if len(_sagemakerExpiresInSeconds) > 0 {
		if err := assignInputField(input, "ExpiresInSeconds", _sagemakerExpiresInSeconds); err != nil {
			log.Errorf("invalid --expires-in-seconds: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSessionExpirationDurationInSeconds) > 0 {
		if err := assignInputField(input, "SessionExpirationDurationInSeconds", _sagemakerSessionExpirationDurationInSeconds); err != nil {
			log.Errorf("invalid --session-expiration-duration-in-seconds: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreatePresignedMlflowTrackingServerUrl(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a URL that you can use to connect to the Jupyter server from a notebook
// instance. In the SageMaker AI console, when you choose Open next to a notebook
// instance, SageMaker AI opens a new tab showing the Jupyter server home page from
// the notebook instance. The console uses this API to get the URL and show the
// page.
//
// The IAM role or user used to call this API defines the permissions to access
// the notebook instance. Once the presigned URL is created, no additional
// permission is required to access this URL. IAM authorization policies for this
// API are also enforced for every HTTP request and WebSocket frame that attempts
// to connect to the notebook instance.
//
// You can restrict access to this API and to the URL that it returns to a list of
// IP addresses that you specify. Use the NotIpAddress condition operator and the
// aws:SourceIP condition context key to specify the list of IP addresses that you
// want to have access to the notebook instance. For more information, see [Limit Access to a Notebook Instance by IP Address].
//
// The URL that you get from a call to [CreatePresignedNotebookInstanceUrl] is valid only for 5 minutes. If you try to
// use the URL after the 5-minute limit expires, you are directed to the Amazon Web
// Services console sign-in page.
//
// [CreatePresignedNotebookInstanceUrl]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreatePresignedNotebookInstanceUrl.html
// [Limit Access to a Notebook Instance by IP Address]: https://docs.aws.amazon.com/sagemaker/latest/dg/security_iam_id-based-policy-examples.html#nbi-ip-filter
func sagemaker_CreatePresignedNotebookInstanceUrl(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreatePresignedNotebookInstanceUrlInput{
		// NotebookInstanceName: *string, // Required
	}

	if len(_sagemakerNotebookInstanceName) > 0 {
		input.NotebookInstanceName = aws.String(_sagemakerNotebookInstanceName)
	}
	if len(_sagemakerSessionExpirationDurationInSeconds) > 0 {
		if err := assignInputField(input, "SessionExpirationDurationInSeconds", _sagemakerSessionExpirationDurationInSeconds); err != nil {
			log.Errorf("invalid --session-expiration-duration-in-seconds: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreatePresignedNotebookInstanceUrl(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a processing job.
func sagemaker_CreateProcessingJob(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreateProcessingJobInput{
		// AppSpecification: *types.AppSpecification, // Required
		// ProcessingJobName: *string, // Required
		// ProcessingResources: *types.ProcessingResources, // Required
		// RoleArn: *string, // Required
	}

	if len(_sagemakerAppSpecification) > 0 {
		if err := assignInputField(input, "AppSpecification", _sagemakerAppSpecification); err != nil {
			log.Errorf("invalid --app-specification: %s", err.Error())
			return
		}
	}
	if len(_sagemakerProcessingJobName) > 0 {
		input.ProcessingJobName = aws.String(_sagemakerProcessingJobName)
	}
	if len(_sagemakerProcessingResources) > 0 {
		if err := assignInputField(input, "ProcessingResources", _sagemakerProcessingResources); err != nil {
			log.Errorf("invalid --processing-resources: %s", err.Error())
			return
		}
	}
	if len(_sagemakerRoleArn) > 0 {
		input.RoleArn = aws.String(_sagemakerRoleArn)
	}
	if len(_sagemakerEnvironment) > 0 {
		if err := assignInputField(input, "Environment", _sagemakerEnvironment); err != nil {
			log.Errorf("invalid --environment: %s", err.Error())
			return
		}
	}
	if len(_sagemakerExperimentConfig) > 0 {
		if err := assignInputField(input, "ExperimentConfig", _sagemakerExperimentConfig); err != nil {
			log.Errorf("invalid --experiment-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNetworkConfig) > 0 {
		if err := assignInputField(input, "NetworkConfig", _sagemakerNetworkConfig); err != nil {
			log.Errorf("invalid --network-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerProcessingInputs) > 0 {
		if err := assignInputField(input, "ProcessingInputs", _sagemakerProcessingInputs); err != nil {
			log.Errorf("invalid --processing-inputs: %s", err.Error())
			return
		}
	}
	if len(_sagemakerProcessingOutputConfig) > 0 {
		if err := assignInputField(input, "ProcessingOutputConfig", _sagemakerProcessingOutputConfig); err != nil {
			log.Errorf("invalid --processing-output-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerStoppingCondition) > 0 {
		if err := assignInputField(input, "StoppingCondition", _sagemakerStoppingCondition); err != nil {
			log.Errorf("invalid --stopping-condition: %s", err.Error())
			return
		}
	}
	if len(_sagemakerTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateProcessingJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a machine learning (ML) project that can contain one or more templates
// that set up an ML pipeline from training to deploying an approved model.
func sagemaker_CreateProject(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreateProjectInput{
		// ProjectName: *string, // Required
	}

	if len(_sagemakerProjectName) > 0 {
		input.ProjectName = aws.String(_sagemakerProjectName)
	}
	if len(_sagemakerProjectDescription) > 0 {
		input.ProjectDescription = aws.String(_sagemakerProjectDescription)
	}
	if len(_sagemakerServiceCatalogProvisioningDetails) > 0 {
		if err := assignInputField(input, "ServiceCatalogProvisioningDetails", _sagemakerServiceCatalogProvisioningDetails); err != nil {
			log.Errorf("invalid --service-catalog-provisioning-details: %s", err.Error())
			return
		}
	}
	if len(_sagemakerTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_sagemakerTemplateProviders) > 0 {
		if err := assignInputField(input, "TemplateProviders", _sagemakerTemplateProviders); err != nil {
			log.Errorf("invalid --template-providers: %s", err.Error())
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

// Creates a private space or a space used for real time collaboration in a domain.
func sagemaker_CreateSpace(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreateSpaceInput{
		// DomainId: *string, // Required
		// SpaceName: *string, // Required
	}

	if len(_sagemakerDomainId) > 0 {
		input.DomainId = aws.String(_sagemakerDomainId)
	}
	if len(_sagemakerSpaceName) > 0 {
		input.SpaceName = aws.String(_sagemakerSpaceName)
	}
	if len(_sagemakerOwnershipSettings) > 0 {
		if err := assignInputField(input, "OwnershipSettings", _sagemakerOwnershipSettings); err != nil {
			log.Errorf("invalid --ownership-settings: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSpaceDisplayName) > 0 {
		input.SpaceDisplayName = aws.String(_sagemakerSpaceDisplayName)
	}
	if len(_sagemakerSpaceSettings) > 0 {
		if err := assignInputField(input, "SpaceSettings", _sagemakerSpaceSettings); err != nil {
			log.Errorf("invalid --space-settings: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSpaceSharingSettings) > 0 {
		if err := assignInputField(input, "SpaceSharingSettings", _sagemakerSpaceSharingSettings); err != nil {
			log.Errorf("invalid --space-sharing-settings: %s", err.Error())
			return
		}
	}
	if len(_sagemakerTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSpace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new Amazon SageMaker AI Studio Lifecycle Configuration.
func sagemaker_CreateStudioLifecycleConfig(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreateStudioLifecycleConfigInput{
		// StudioLifecycleConfigAppType: types.StudioLifecycleConfigAppType, // Required
		// StudioLifecycleConfigContent: *string, // Required
		// StudioLifecycleConfigName: *string, // Required
	}

	if len(_sagemakerStudioLifecycleConfigAppType) > 0 {
		if err := assignInputField(input, "StudioLifecycleConfigAppType", _sagemakerStudioLifecycleConfigAppType); err != nil {
			log.Errorf("invalid --studio-lifecycle-config-app-type: %s", err.Error())
			return
		}
	}
	if len(_sagemakerStudioLifecycleConfigContent) > 0 {
		input.StudioLifecycleConfigContent = aws.String(_sagemakerStudioLifecycleConfigContent)
	}
	if len(_sagemakerStudioLifecycleConfigName) > 0 {
		input.StudioLifecycleConfigName = aws.String(_sagemakerStudioLifecycleConfigName)
	}
	if len(_sagemakerTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateStudioLifecycleConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a model training job. After training completes, SageMaker saves the
// resulting model artifacts to an Amazon S3 location that you specify.
//
// If you choose to host your model using SageMaker hosting services, you can use
// the resulting model artifacts as part of the model. You can also use the
// artifacts in a machine learning service other than SageMaker, provided that you
// know how to use them for inference.
//
// In the request body, you provide the following:
//
// - AlgorithmSpecification - Identifies the training algorithm to use.
//
// - HyperParameters - Specify these algorithm-specific parameters to enable the
// estimation of model parameters during training. Hyperparameters can be tuned to
// optimize this learning process. For a list of hyperparameters for each training
// algorithm provided by SageMaker, see [Algorithms].
//
// Do not include any security-sensitive information including account access IDs,
//
// secrets, or tokens in any hyperparameter fields. As part of the shared
// responsibility model, you are responsible for any potential exposure,
// unauthorized access, or compromise of your sensitive data if caused by
// security-sensitive information included in the request hyperparameter variable
// or plain text fields.
//
// - InputDataConfig - Describes the input required by the training job and the
// Amazon S3, EFS, or FSx location where it is stored.
//
// - OutputDataConfig - Identifies the Amazon S3 bucket where you want SageMaker
// to save the results of model training.
//
// - ResourceConfig - Identifies the resources, ML compute instances, and ML
// storage volumes to deploy for model training. In distributed training, you
// specify more than one instance.
//
// - EnableManagedSpotTraining - Optimize the cost of training machine learning
// models by up to 80% by using Amazon EC2 Spot instances. For more information,
// see [Managed Spot Training].
//
// - RoleArn - The Amazon Resource Name (ARN) that SageMaker assumes to perform
// tasks on your behalf during model training. You must grant this role the
// necessary permissions so that SageMaker can successfully complete model
// training.
//
// - StoppingCondition - To help cap training costs, use MaxRuntimeInSeconds to
// set a time limit for training. Use MaxWaitTimeInSeconds to specify how long a
// managed spot training job has to complete.
//
// - Environment - The environment variables to set in the Docker container.
//
// Do not include any security-sensitive information including account access IDs,
//
// secrets, or tokens in any environment fields. As part of the shared
// responsibility model, you are responsible for any potential exposure,
// unauthorized access, or compromise of your sensitive data if caused by
// security-sensitive information included in the request environment variable or
// plain text fields.
//
// - RetryStrategy - The number of times to retry the job when the job fails due
// to an InternalServerError .
//
// For more information about SageMaker, see [How It Works].
//
// [Algorithms]: https://docs.aws.amazon.com/sagemaker/latest/dg/algos.html
// [How It Works]: https://docs.aws.amazon.com/sagemaker/latest/dg/how-it-works.html
// [Managed Spot Training]: https://docs.aws.amazon.com/sagemaker/latest/dg/model-managed-spot-training.html
func sagemaker_CreateTrainingJob(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreateTrainingJobInput{
		// OutputDataConfig: *types.OutputDataConfig, // Required
		// RoleArn: *string, // Required
		// TrainingJobName: *string, // Required
	}

	if len(_sagemakerOutputDataConfig) > 0 {
		if err := assignInputField(input, "OutputDataConfig", _sagemakerOutputDataConfig); err != nil {
			log.Errorf("invalid --output-data-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerRoleArn) > 0 {
		input.RoleArn = aws.String(_sagemakerRoleArn)
	}
	if len(_sagemakerTrainingJobName) > 0 {
		input.TrainingJobName = aws.String(_sagemakerTrainingJobName)
	}
	if len(_sagemakerAlgorithmSpecification) > 0 {
		if err := assignInputField(input, "AlgorithmSpecification", _sagemakerAlgorithmSpecification); err != nil {
			log.Errorf("invalid --algorithm-specification: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCheckpointConfig) > 0 {
		if err := assignInputField(input, "CheckpointConfig", _sagemakerCheckpointConfig); err != nil {
			log.Errorf("invalid --checkpoint-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerDebugHookConfig) > 0 {
		if err := assignInputField(input, "DebugHookConfig", _sagemakerDebugHookConfig); err != nil {
			log.Errorf("invalid --debug-hook-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerDebugRuleConfigurations) > 0 {
		if err := assignInputField(input, "DebugRuleConfigurations", _sagemakerDebugRuleConfigurations); err != nil {
			log.Errorf("invalid --debug-rule-configurations: %s", err.Error())
			return
		}
	}
	if len(_sagemakerEnableInterContainerTrafficEncryption) > 0 {
		if err := assignInputField(input, "EnableInterContainerTrafficEncryption", _sagemakerEnableInterContainerTrafficEncryption); err != nil {
			log.Errorf("invalid --enable-inter-container-traffic-encryption: %s", err.Error())
			return
		}
	}
	if len(_sagemakerEnableManagedSpotTraining) > 0 {
		if err := assignInputField(input, "EnableManagedSpotTraining", _sagemakerEnableManagedSpotTraining); err != nil {
			log.Errorf("invalid --enable-managed-spot-training: %s", err.Error())
			return
		}
	}
	if len(_sagemakerEnableNetworkIsolation) > 0 {
		if err := assignInputField(input, "EnableNetworkIsolation", _sagemakerEnableNetworkIsolation); err != nil {
			log.Errorf("invalid --enable-network-isolation: %s", err.Error())
			return
		}
	}
	if len(_sagemakerEnvironment) > 0 {
		if err := assignInputField(input, "Environment", _sagemakerEnvironment); err != nil {
			log.Errorf("invalid --environment: %s", err.Error())
			return
		}
	}
	if len(_sagemakerExperimentConfig) > 0 {
		if err := assignInputField(input, "ExperimentConfig", _sagemakerExperimentConfig); err != nil {
			log.Errorf("invalid --experiment-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerHyperParameters) > 0 {
		if err := assignInputField(input, "HyperParameters", _sagemakerHyperParameters); err != nil {
			log.Errorf("invalid --hyper-parameters: %s", err.Error())
			return
		}
	}
	if len(_sagemakerInfraCheckConfig) > 0 {
		if err := assignInputField(input, "InfraCheckConfig", _sagemakerInfraCheckConfig); err != nil {
			log.Errorf("invalid --infra-check-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerInputDataConfig) > 0 {
		if err := assignInputField(input, "InputDataConfig", _sagemakerInputDataConfig); err != nil {
			log.Errorf("invalid --input-data-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMlflowConfig) > 0 {
		if err := assignInputField(input, "MlflowConfig", _sagemakerMlflowConfig); err != nil {
			log.Errorf("invalid --mlflow-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerModelPackageConfig) > 0 {
		if err := assignInputField(input, "ModelPackageConfig", _sagemakerModelPackageConfig); err != nil {
			log.Errorf("invalid --model-package-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerProfilerConfig) > 0 {
		if err := assignInputField(input, "ProfilerConfig", _sagemakerProfilerConfig); err != nil {
			log.Errorf("invalid --profiler-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerProfilerRuleConfigurations) > 0 {
		if err := assignInputField(input, "ProfilerRuleConfigurations", _sagemakerProfilerRuleConfigurations); err != nil {
			log.Errorf("invalid --profiler-rule-configurations: %s", err.Error())
			return
		}
	}
	if len(_sagemakerRemoteDebugConfig) > 0 {
		if err := assignInputField(input, "RemoteDebugConfig", _sagemakerRemoteDebugConfig); err != nil {
			log.Errorf("invalid --remote-debug-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerResourceConfig) > 0 {
		if err := assignInputField(input, "ResourceConfig", _sagemakerResourceConfig); err != nil {
			log.Errorf("invalid --resource-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerRetryStrategy) > 0 {
		if err := assignInputField(input, "RetryStrategy", _sagemakerRetryStrategy); err != nil {
			log.Errorf("invalid --retry-strategy: %s", err.Error())
			return
		}
	}
	if len(_sagemakerServerlessJobConfig) > 0 {
		if err := assignInputField(input, "ServerlessJobConfig", _sagemakerServerlessJobConfig); err != nil {
			log.Errorf("invalid --serverless-job-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSessionChainingConfig) > 0 {
		if err := assignInputField(input, "SessionChainingConfig", _sagemakerSessionChainingConfig); err != nil {
			log.Errorf("invalid --session-chaining-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerStoppingCondition) > 0 {
		if err := assignInputField(input, "StoppingCondition", _sagemakerStoppingCondition); err != nil {
			log.Errorf("invalid --stopping-condition: %s", err.Error())
			return
		}
	}
	if len(_sagemakerTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_sagemakerTensorBoardOutputConfig) > 0 {
		if err := assignInputField(input, "TensorBoardOutputConfig", _sagemakerTensorBoardOutputConfig); err != nil {
			log.Errorf("invalid --tensor-board-output-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerVpcConfig) > 0 {
		if err := assignInputField(input, "VpcConfig", _sagemakerVpcConfig); err != nil {
			log.Errorf("invalid --vpc-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateTrainingJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new training plan in SageMaker to reserve compute capacity.
// Amazon SageMaker Training Plan is a capability within SageMaker that allows
// customers to reserve and manage GPU capacity for large-scale AI model training.
// It provides a way to secure predictable access to computational resources within
// specific timelines and budgets, without the need to manage underlying
// infrastructure.
//
// # How it works
//
// Plans can be created for specific resources such as SageMaker Training Jobs or
// SageMaker HyperPod clusters, automatically provisioning resources, setting up
// infrastructure, executing workloads, and handling infrastructure failures.
//
// # Plan creation workflow
//
// - Users search for available plan offerings based on their requirements
// (e.g., instance type, count, start time, duration) using the [SearchTrainingPlanOfferings]API operation.
//
// - They create a plan that best matches their needs using the ID of the plan
// offering they want to use.
//
// - After successful upfront payment, the plan's status becomes Scheduled .
//
// - The plan can be used to:
//
// - Queue training jobs.
//
// - Allocate to an instance group of a SageMaker HyperPod cluster.
//
// - When the plan start date arrives, it becomes Active . Based on available
// reserved capacity:
//
// - Training jobs are launched.
//
// - Instance groups are provisioned.
//
// # Plan composition
//
// A plan can consist of one or more Reserved Capacities, each defined by a
// specific instance type, quantity, Availability Zone, duration, and start and end
// times. For more information about Reserved Capacity, see [ReservedCapacitySummary].
//
// [SearchTrainingPlanOfferings]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_SearchTrainingPlanOfferings.html
// [ReservedCapacitySummary]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_ReservedCapacitySummary.html
func sagemaker_CreateTrainingPlan(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreateTrainingPlanInput{
		// TrainingPlanName: *string, // Required
		// TrainingPlanOfferingId: *string, // Required
	}

	if len(_sagemakerTrainingPlanName) > 0 {
		input.TrainingPlanName = aws.String(_sagemakerTrainingPlanName)
	}
	if len(_sagemakerTrainingPlanOfferingId) > 0 {
		input.TrainingPlanOfferingId = aws.String(_sagemakerTrainingPlanOfferingId)
	}
	if len(_sagemakerSpareInstanceCountPerUltraServer) > 0 {
		if err := assignInputField(input, "SpareInstanceCountPerUltraServer", _sagemakerSpareInstanceCountPerUltraServer); err != nil {
			log.Errorf("invalid --spare-instance-count-per-ultra-server: %s", err.Error())
			return
		}
	}
	if len(_sagemakerTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateTrainingPlan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a transform job. A transform job uses a trained model to get inferences
// on a dataset and saves these results to an Amazon S3 location that you specify.
//
// To perform batch transformations, you create a transform job and use the data
// that you have readily available.
//
// In the request body, you provide the following:
//
// - TransformJobName - Identifies the transform job. The name must be unique
// within an Amazon Web Services Region in an Amazon Web Services account.
//
// - ModelName - Identifies the model to use. ModelName must be the name of an
// existing Amazon SageMaker model in the same Amazon Web Services Region and
// Amazon Web Services account. For information on creating a model, see [CreateModel].
//
// - TransformInput - Describes the dataset to be transformed and the Amazon S3
// location where it is stored.
//
// - TransformOutput - Identifies the Amazon S3 location where you want Amazon
// SageMaker to save the results from the transform job.
//
// - TransformResources - Identifies the ML compute instances and AMI image
// versions for the transform job.
//
// For more information about how batch transformation works, see [Batch Transform].
//
// [CreateModel]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateModel.html
// [Batch Transform]: https://docs.aws.amazon.com/sagemaker/latest/dg/batch-transform.html
func sagemaker_CreateTransformJob(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreateTransformJobInput{
		// ModelName: *string, // Required
		// TransformInput: *types.TransformInput, // Required
		// TransformJobName: *string, // Required
		// TransformOutput: *types.TransformOutput, // Required
		// TransformResources: *types.TransformResources, // Required
	}

	if len(_sagemakerModelName) > 0 {
		input.ModelName = aws.String(_sagemakerModelName)
	}
	if len(_sagemakerTransformInput) > 0 {
		if err := assignInputField(input, "TransformInput", _sagemakerTransformInput); err != nil {
			log.Errorf("invalid --transform-input: %s", err.Error())
			return
		}
	}
	if len(_sagemakerTransformJobName) > 0 {
		input.TransformJobName = aws.String(_sagemakerTransformJobName)
	}
	if len(_sagemakerTransformOutput) > 0 {
		if err := assignInputField(input, "TransformOutput", _sagemakerTransformOutput); err != nil {
			log.Errorf("invalid --transform-output: %s", err.Error())
			return
		}
	}
	if len(_sagemakerTransformResources) > 0 {
		if err := assignInputField(input, "TransformResources", _sagemakerTransformResources); err != nil {
			log.Errorf("invalid --transform-resources: %s", err.Error())
			return
		}
	}
	if len(_sagemakerBatchStrategy) > 0 {
		if err := assignInputField(input, "BatchStrategy", _sagemakerBatchStrategy); err != nil {
			log.Errorf("invalid --batch-strategy: %s", err.Error())
			return
		}
	}
	if len(_sagemakerDataCaptureConfig) > 0 {
		if err := assignInputField(input, "DataCaptureConfig", _sagemakerDataCaptureConfig); err != nil {
			log.Errorf("invalid --data-capture-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerDataProcessing) > 0 {
		if err := assignInputField(input, "DataProcessing", _sagemakerDataProcessing); err != nil {
			log.Errorf("invalid --data-processing: %s", err.Error())
			return
		}
	}
	if len(_sagemakerEnvironment) > 0 {
		if err := assignInputField(input, "Environment", _sagemakerEnvironment); err != nil {
			log.Errorf("invalid --environment: %s", err.Error())
			return
		}
	}
	if len(_sagemakerExperimentConfig) > 0 {
		if err := assignInputField(input, "ExperimentConfig", _sagemakerExperimentConfig); err != nil {
			log.Errorf("invalid --experiment-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMaxConcurrentTransforms) > 0 {
		if err := assignInputField(input, "MaxConcurrentTransforms", _sagemakerMaxConcurrentTransforms); err != nil {
			log.Errorf("invalid --max-concurrent-transforms: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMaxPayloadInMB) > 0 {
		if err := assignInputField(input, "MaxPayloadInMB", _sagemakerMaxPayloadInMB); err != nil {
			log.Errorf("invalid --max-payload-in-mb: %s", err.Error())
			return
		}
	}
	if len(_sagemakerModelClientConfig) > 0 {
		if err := assignInputField(input, "ModelClientConfig", _sagemakerModelClientConfig); err != nil {
			log.Errorf("invalid --model-client-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateTransformJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an SageMaker trial. A trial is a set of steps called trial components
// that produce a machine learning model. A trial is part of a single SageMaker
// experiment.
//
// When you use SageMaker Studio or the SageMaker Python SDK, all experiments,
// trials, and trial components are automatically tracked, logged, and indexed.
// When you use the Amazon Web Services SDK for Python (Boto), you must use the
// logging APIs provided by the SDK.
//
// You can add tags to a trial and then use the [Search] API to search for the tags.
//
// To get a list of all your trials, call the [ListTrials] API. To view a trial's properties,
// call the [DescribeTrial]API. To create a trial component, call the [CreateTrialComponent] API.
//
// [DescribeTrial]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DescribeTrial.html
// [Search]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_Search.html
// [ListTrials]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_ListTrials.html
// [CreateTrialComponent]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateTrialComponent.html
func sagemaker_CreateTrial(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreateTrialInput{
		// ExperimentName: *string, // Required
		// TrialName: *string, // Required
	}

	if len(_sagemakerExperimentName) > 0 {
		input.ExperimentName = aws.String(_sagemakerExperimentName)
	}
	if len(_sagemakerTrialName) > 0 {
		input.TrialName = aws.String(_sagemakerTrialName)
	}
	if len(_sagemakerDisplayName) > 0 {
		input.DisplayName = aws.String(_sagemakerDisplayName)
	}
	if len(_sagemakerMetadataProperties) > 0 {
		if err := assignInputField(input, "MetadataProperties", _sagemakerMetadataProperties); err != nil {
			log.Errorf("invalid --metadata-properties: %s", err.Error())
			return
		}
	}
	if len(_sagemakerTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateTrial(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a trial component, which is a stage of a machine learning trial. A
// trial is composed of one or more trial components. A trial component can be used
// in multiple trials.
//
// Trial components include pre-processing jobs, training jobs, and batch
// transform jobs.
//
// When you use SageMaker Studio or the SageMaker Python SDK, all experiments,
// trials, and trial components are automatically tracked, logged, and indexed.
// When you use the Amazon Web Services SDK for Python (Boto), you must use the
// logging APIs provided by the SDK.
//
// You can add tags to a trial component and then use the [Search] API to search for the
// tags.
//
// [Search]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_Search.html
func sagemaker_CreateTrialComponent(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreateTrialComponentInput{
		// TrialComponentName: *string, // Required
	}

	if len(_sagemakerTrialComponentName) > 0 {
		input.TrialComponentName = aws.String(_sagemakerTrialComponentName)
	}
	if len(_sagemakerDisplayName) > 0 {
		input.DisplayName = aws.String(_sagemakerDisplayName)
	}
	if len(_sagemakerEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _sagemakerEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_sagemakerInputArtifacts) > 0 {
		if err := assignInputField(input, "InputArtifacts", _sagemakerInputArtifacts); err != nil {
			log.Errorf("invalid --input-artifacts: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMetadataProperties) > 0 {
		if err := assignInputField(input, "MetadataProperties", _sagemakerMetadataProperties); err != nil {
			log.Errorf("invalid --metadata-properties: %s", err.Error())
			return
		}
	}
	if len(_sagemakerOutputArtifacts) > 0 {
		if err := assignInputField(input, "OutputArtifacts", _sagemakerOutputArtifacts); err != nil {
			log.Errorf("invalid --output-artifacts: %s", err.Error())
			return
		}
	}
	if len(_sagemakerParameters) > 0 {
		if err := assignInputField(input, "Parameters", _sagemakerParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}
	if len(_sagemakerStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _sagemakerStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_sagemakerStatus) > 0 {
		if err := assignInputField(input, "Status", _sagemakerStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_sagemakerTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateTrialComponent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a user profile. A user profile represents a single user within a
// domain, and is the main way to reference a "person" for the purposes of sharing,
// reporting, and other user-oriented features. This entity is created when a user
// onboards to a domain. If an administrator invites a person by email or imports
// them from IAM Identity Center, a user profile is automatically created. A user
// profile is the primary holder of settings for an individual user and has a
// reference to the user's private Amazon Elastic File System home directory.
func sagemaker_CreateUserProfile(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreateUserProfileInput{
		// DomainId: *string, // Required
		// UserProfileName: *string, // Required
	}

	if len(_sagemakerDomainId) > 0 {
		input.DomainId = aws.String(_sagemakerDomainId)
	}
	if len(_sagemakerUserProfileName) > 0 {
		input.UserProfileName = aws.String(_sagemakerUserProfileName)
	}
	if len(_sagemakerSingleSignOnUserIdentifier) > 0 {
		input.SingleSignOnUserIdentifier = aws.String(_sagemakerSingleSignOnUserIdentifier)
	}
	if len(_sagemakerSingleSignOnUserValue) > 0 {
		input.SingleSignOnUserValue = aws.String(_sagemakerSingleSignOnUserValue)
	}
	if len(_sagemakerTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_sagemakerUserSettings) > 0 {
		if err := assignInputField(input, "UserSettings", _sagemakerUserSettings); err != nil {
			log.Errorf("invalid --user-settings: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateUserProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Use this operation to create a workforce. This operation will return an error
// if a workforce already exists in the Amazon Web Services Region that you
// specify. You can only create one workforce in each Amazon Web Services Region
// per Amazon Web Services account.
//
// If you want to create a new workforce in an Amazon Web Services Region where a
// workforce already exists, use the [DeleteWorkforce]API operation to delete the existing
// workforce and then use CreateWorkforce to create a new workforce.
//
// To create a private workforce using Amazon Cognito, you must specify a Cognito
// user pool in CognitoConfig . You can also create an Amazon Cognito workforce
// using the Amazon SageMaker console. For more information, see [Create a Private Workforce (Amazon Cognito)].
//
// To create a private workforce using your own OIDC Identity Provider (IdP),
// specify your IdP configuration in OidcConfig . Your OIDC IdP must support groups
// because groups are used by Ground Truth and Amazon A2I to create work teams. For
// more information, see [Create a Private Workforce (OIDC IdP)].
//
// [Create a Private Workforce (Amazon Cognito)]: https://docs.aws.amazon.com/sagemaker/latest/dg/sms-workforce-create-private.html
// [DeleteWorkforce]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DeleteWorkforce.html
// [Create a Private Workforce (OIDC IdP)]: https://docs.aws.amazon.com/sagemaker/latest/dg/sms-workforce-create-private-oidc.html
func sagemaker_CreateWorkforce(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreateWorkforceInput{
		// WorkforceName: *string, // Required
	}

	if len(_sagemakerWorkforceName) > 0 {
		input.WorkforceName = aws.String(_sagemakerWorkforceName)
	}
	if len(_sagemakerCognitoConfig) > 0 {
		if err := assignInputField(input, "CognitoConfig", _sagemakerCognitoConfig); err != nil {
			log.Errorf("invalid --cognito-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerIpAddressType) > 0 {
		if err := assignInputField(input, "IpAddressType", _sagemakerIpAddressType); err != nil {
			log.Errorf("invalid --ip-address-type: %s", err.Error())
			return
		}
	}
	if len(_sagemakerOidcConfig) > 0 {
		if err := assignInputField(input, "OidcConfig", _sagemakerOidcConfig); err != nil {
			log.Errorf("invalid --oidc-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSourceIpConfig) > 0 {
		if err := assignInputField(input, "SourceIpConfig", _sagemakerSourceIpConfig); err != nil {
			log.Errorf("invalid --source-ip-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_sagemakerWorkforceVpcConfig) > 0 {
		if err := assignInputField(input, "WorkforceVpcConfig", _sagemakerWorkforceVpcConfig); err != nil {
			log.Errorf("invalid --workforce-vpc-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateWorkforce(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new work team for labeling your data. A work team is defined by one
// or more Amazon Cognito user pools. You must first create the user pools before
// you can create a work team.
//
// You cannot create more than 25 work teams in an account and region.
func sagemaker_CreateWorkteam(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.CreateWorkteamInput{
		// Description: *string, // Required
		// MemberDefinitions: []types.MemberDefinition, // Required
		// WorkteamName: *string, // Required
	}

	if len(_sagemakerDescription) > 0 {
		input.Description = aws.String(_sagemakerDescription)
	}
	if len(_sagemakerMemberDefinitions) > 0 {
		if err := assignInputField(input, "MemberDefinitions", _sagemakerMemberDefinitions); err != nil {
			log.Errorf("invalid --member-definitions: %s", err.Error())
			return
		}
	}
	if len(_sagemakerWorkteamName) > 0 {
		input.WorkteamName = aws.String(_sagemakerWorkteamName)
	}
	if len(_sagemakerNotificationConfiguration) > 0 {
		if err := assignInputField(input, "NotificationConfiguration", _sagemakerNotificationConfiguration); err != nil {
			log.Errorf("invalid --notification-configuration: %s", err.Error())
			return
		}
	}
	if len(_sagemakerTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_sagemakerWorkerAccessConfiguration) > 0 {
		if err := assignInputField(input, "WorkerAccessConfiguration", _sagemakerWorkerAccessConfiguration); err != nil {
			log.Errorf("invalid --worker-access-configuration: %s", err.Error())
			return
		}
	}
	if len(_sagemakerWorkforceName) > 0 {
		input.WorkforceName = aws.String(_sagemakerWorkforceName)
	}

	if resp, err := client.CreateWorkteam(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an action.
func sagemaker_DeleteAction(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DeleteActionInput{
		// ActionName: *string, // Required
	}

	if len(_sagemakerActionName) > 0 {
		input.ActionName = aws.String(_sagemakerActionName)
	}

	if resp, err := client.DeleteAction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the specified algorithm from your account.
func sagemaker_DeleteAlgorithm(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DeleteAlgorithmInput{
		// AlgorithmName: *string, // Required
	}

	if len(_sagemakerAlgorithmName) > 0 {
		input.AlgorithmName = aws.String(_sagemakerAlgorithmName)
	}

	if resp, err := client.DeleteAlgorithm(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Used to stop and delete an app.
func sagemaker_DeleteApp(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DeleteAppInput{
		// AppName: *string, // Required
		// AppType: types.AppType, // Required
		// DomainId: *string, // Required
	}

	if len(_sagemakerAppName) > 0 {
		input.AppName = aws.String(_sagemakerAppName)
	}
	if len(_sagemakerAppType) > 0 {
		if err := assignInputField(input, "AppType", _sagemakerAppType); err != nil {
			log.Errorf("invalid --app-type: %s", err.Error())
			return
		}
	}
	if len(_sagemakerDomainId) > 0 {
		input.DomainId = aws.String(_sagemakerDomainId)
	}
	if len(_sagemakerSpaceName) > 0 {
		input.SpaceName = aws.String(_sagemakerSpaceName)
	}
	if len(_sagemakerUserProfileName) > 0 {
		input.UserProfileName = aws.String(_sagemakerUserProfileName)
	}

	if resp, err := client.DeleteApp(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an AppImageConfig.
func sagemaker_DeleteAppImageConfig(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DeleteAppImageConfigInput{
		// AppImageConfigName: *string, // Required
	}

	if len(_sagemakerAppImageConfigName) > 0 {
		input.AppImageConfigName = aws.String(_sagemakerAppImageConfigName)
	}

	if resp, err := client.DeleteAppImageConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an artifact. Either ArtifactArn or Source must be specified.
func sagemaker_DeleteArtifact(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DeleteArtifactInput{}

	if len(_sagemakerArtifactArn) > 0 {
		input.ArtifactArn = aws.String(_sagemakerArtifactArn)
	}
	if len(_sagemakerSource) > 0 {
		if err := assignInputField(input, "Source", _sagemakerSource); err != nil {
			log.Errorf("invalid --source: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteArtifact(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an association.
func sagemaker_DeleteAssociation(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DeleteAssociationInput{
		// DestinationArn: *string, // Required
		// SourceArn: *string, // Required
	}

	if len(_sagemakerDestinationArn) > 0 {
		input.DestinationArn = aws.String(_sagemakerDestinationArn)
	}
	if len(_sagemakerSourceArn) > 0 {
		input.SourceArn = aws.String(_sagemakerSourceArn)
	}

	if resp, err := client.DeleteAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete a SageMaker HyperPod cluster.
func sagemaker_DeleteCluster(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DeleteClusterInput{
		// ClusterName: *string, // Required
	}

	if len(_sagemakerClusterName) > 0 {
		input.ClusterName = aws.String(_sagemakerClusterName)
	}

	if resp, err := client.DeleteCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the cluster policy of the cluster.
func sagemaker_DeleteClusterSchedulerConfig(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DeleteClusterSchedulerConfigInput{
		// ClusterSchedulerConfigId: *string, // Required
	}

	if len(_sagemakerClusterSchedulerConfigId) > 0 {
		input.ClusterSchedulerConfigId = aws.String(_sagemakerClusterSchedulerConfigId)
	}

	if resp, err := client.DeleteClusterSchedulerConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified Git repository from your account.
func sagemaker_DeleteCodeRepository(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DeleteCodeRepositoryInput{
		// CodeRepositoryName: *string, // Required
	}

	if len(_sagemakerCodeRepositoryName) > 0 {
		input.CodeRepositoryName = aws.String(_sagemakerCodeRepositoryName)
	}

	if resp, err := client.DeleteCodeRepository(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified compilation job. This action deletes only the compilation
// job resource in Amazon SageMaker AI. It doesn't delete other resources that are
// related to that job, such as the model artifacts that the job creates, the
// compilation logs in CloudWatch, the compiled model, or the IAM role.
//
// You can delete a compilation job only if its current status is COMPLETED ,
// FAILED , or STOPPED . If the job status is STARTING or INPROGRESS , stop the
// job, and then delete it after its status becomes STOPPED .
func sagemaker_DeleteCompilationJob(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DeleteCompilationJobInput{
		// CompilationJobName: *string, // Required
	}

	if len(_sagemakerCompilationJobName) > 0 {
		input.CompilationJobName = aws.String(_sagemakerCompilationJobName)
	}

	if resp, err := client.DeleteCompilationJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the compute allocation from the cluster.
func sagemaker_DeleteComputeQuota(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DeleteComputeQuotaInput{
		// ComputeQuotaId: *string, // Required
	}

	if len(_sagemakerComputeQuotaId) > 0 {
		input.ComputeQuotaId = aws.String(_sagemakerComputeQuotaId)
	}

	if resp, err := client.DeleteComputeQuota(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an context.
func sagemaker_DeleteContext(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DeleteContextInput{
		// ContextName: *string, // Required
	}

	if len(_sagemakerContextName) > 0 {
		input.ContextName = aws.String(_sagemakerContextName)
	}

	if resp, err := client.DeleteContext(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a data quality monitoring job definition.
func sagemaker_DeleteDataQualityJobDefinition(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DeleteDataQualityJobDefinitionInput{
		// JobDefinitionName: *string, // Required
	}

	if len(_sagemakerJobDefinitionName) > 0 {
		input.JobDefinitionName = aws.String(_sagemakerJobDefinitionName)
	}

	if resp, err := client.DeleteDataQualityJobDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a fleet.
func sagemaker_DeleteDeviceFleet(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DeleteDeviceFleetInput{
		// DeviceFleetName: *string, // Required
	}

	if len(_sagemakerDeviceFleetName) > 0 {
		input.DeviceFleetName = aws.String(_sagemakerDeviceFleetName)
	}

	if resp, err := client.DeleteDeviceFleet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Used to delete a domain. If you onboarded with IAM mode, you will need to
// delete your domain to onboard again using IAM Identity Center. Use with caution.
// All of the members of the domain will lose access to their EFS volume, including
// data, notebooks, and other artifacts.
func sagemaker_DeleteDomain(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DeleteDomainInput{
		// DomainId: *string, // Required
	}

	if len(_sagemakerDomainId) > 0 {
		input.DomainId = aws.String(_sagemakerDomainId)
	}
	if len(_sagemakerRetentionPolicy) > 0 {
		if err := assignInputField(input, "RetentionPolicy", _sagemakerRetentionPolicy); err != nil {
			log.Errorf("invalid --retention-policy: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an edge deployment plan if (and only if) all the stages in the plan are
// inactive or there are no stages in the plan.
func sagemaker_DeleteEdgeDeploymentPlan(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DeleteEdgeDeploymentPlanInput{
		// EdgeDeploymentPlanName: *string, // Required
	}

	if len(_sagemakerEdgeDeploymentPlanName) > 0 {
		input.EdgeDeploymentPlanName = aws.String(_sagemakerEdgeDeploymentPlanName)
	}

	if resp, err := client.DeleteEdgeDeploymentPlan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete a stage in an edge deployment plan if (and only if) the stage is
// inactive.
func sagemaker_DeleteEdgeDeploymentStage(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DeleteEdgeDeploymentStageInput{
		// EdgeDeploymentPlanName: *string, // Required
		// StageName: *string, // Required
	}

	if len(_sagemakerEdgeDeploymentPlanName) > 0 {
		input.EdgeDeploymentPlanName = aws.String(_sagemakerEdgeDeploymentPlanName)
	}
	if len(_sagemakerStageName) > 0 {
		input.StageName = aws.String(_sagemakerStageName)
	}

	if resp, err := client.DeleteEdgeDeploymentStage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an endpoint. SageMaker frees up all of the resources that were deployed
// when the endpoint was created.
//
// SageMaker retires any custom KMS key grants associated with the endpoint,
// meaning you don't need to use the [RevokeGrant]API call.
//
// When you delete your endpoint, SageMaker asynchronously deletes associated
// endpoint resources such as KMS key grants. You might still see these resources
// in your account for a few minutes after deleting your endpoint. Do not delete or
// revoke the permissions for your [ExecutionRoleArn], otherwise SageMaker cannot delete these
// resources.
//
// [ExecutionRoleArn]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateModel.html#sagemaker-CreateModel-request-ExecutionRoleArn
// [RevokeGrant]: http://docs.aws.amazon.com/kms/latest/APIReference/API_RevokeGrant.html
func sagemaker_DeleteEndpoint(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DeleteEndpointInput{
		// EndpointName: *string, // Required
	}

	if len(_sagemakerEndpointName) > 0 {
		input.EndpointName = aws.String(_sagemakerEndpointName)
	}

	if resp, err := client.DeleteEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an endpoint configuration. The DeleteEndpointConfig API deletes only
// the specified configuration. It does not delete endpoints created using the
// configuration.
//
// You must not delete an EndpointConfig in use by an endpoint that is live or
// while the UpdateEndpoint or CreateEndpoint operations are being performed on
// the endpoint. If you delete the EndpointConfig of an endpoint that is active or
// being created or updated you may lose visibility into the instance type the
// endpoint is using. The endpoint must be deleted in order to stop incurring
// charges.
func sagemaker_DeleteEndpointConfig(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DeleteEndpointConfigInput{
		// EndpointConfigName: *string, // Required
	}

	if len(_sagemakerEndpointConfigName) > 0 {
		input.EndpointConfigName = aws.String(_sagemakerEndpointConfigName)
	}

	if resp, err := client.DeleteEndpointConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an SageMaker experiment. All trials associated with the experiment must
// be deleted first. Use the [ListTrials]API to get a list of the trials associated with the
// experiment.
//
// [ListTrials]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_ListTrials.html
func sagemaker_DeleteExperiment(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DeleteExperimentInput{
		// ExperimentName: *string, // Required
	}

	if len(_sagemakerExperimentName) > 0 {
		input.ExperimentName = aws.String(_sagemakerExperimentName)
	}

	if resp, err := client.DeleteExperiment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete the FeatureGroup and any data that was written to the OnlineStore of the
// FeatureGroup . Data cannot be accessed from the OnlineStore immediately after
// DeleteFeatureGroup is called.
//
// Data written into the OfflineStore will not be deleted. The Amazon Web Services
// Glue database and tables that are automatically created for your OfflineStore
// are not deleted.
//
// Note that it can take approximately 10-15 minutes to delete an OnlineStore
// FeatureGroup with the InMemory StorageType .
func sagemaker_DeleteFeatureGroup(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DeleteFeatureGroupInput{
		// FeatureGroupName: *string, // Required
	}

	if len(_sagemakerFeatureGroupName) > 0 {
		input.FeatureGroupName = aws.String(_sagemakerFeatureGroupName)
	}

	if resp, err := client.DeleteFeatureGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified flow definition.
func sagemaker_DeleteFlowDefinition(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DeleteFlowDefinitionInput{
		// FlowDefinitionName: *string, // Required
	}

	if len(_sagemakerFlowDefinitionName) > 0 {
		input.FlowDefinitionName = aws.String(_sagemakerFlowDefinitionName)
	}

	if resp, err := client.DeleteFlowDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete a hub.
func sagemaker_DeleteHub(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DeleteHubInput{
		// HubName: *string, // Required
	}

	if len(_sagemakerHubName) > 0 {
		input.HubName = aws.String(_sagemakerHubName)
	}

	if resp, err := client.DeleteHub(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete the contents of a hub.
func sagemaker_DeleteHubContent(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DeleteHubContentInput{
		// HubContentName: *string, // Required
		// HubContentType: types.HubContentType, // Required
		// HubContentVersion: *string, // Required
		// HubName: *string, // Required
	}

	if len(_sagemakerHubContentName) > 0 {
		input.HubContentName = aws.String(_sagemakerHubContentName)
	}
	if len(_sagemakerHubContentType) > 0 {
		if err := assignInputField(input, "HubContentType", _sagemakerHubContentType); err != nil {
			log.Errorf("invalid --hub-content-type: %s", err.Error())
			return
		}
	}
	if len(_sagemakerHubContentVersion) > 0 {
		input.HubContentVersion = aws.String(_sagemakerHubContentVersion)
	}
	if len(_sagemakerHubName) > 0 {
		input.HubName = aws.String(_sagemakerHubName)
	}

	if resp, err := client.DeleteHubContent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete a hub content reference in order to remove a model from a private hub.
func sagemaker_DeleteHubContentReference(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DeleteHubContentReferenceInput{
		// HubContentName: *string, // Required
		// HubContentType: types.HubContentType, // Required
		// HubName: *string, // Required
	}

	if len(_sagemakerHubContentName) > 0 {
		input.HubContentName = aws.String(_sagemakerHubContentName)
	}
	if len(_sagemakerHubContentType) > 0 {
		if err := assignInputField(input, "HubContentType", _sagemakerHubContentType); err != nil {
			log.Errorf("invalid --hub-content-type: %s", err.Error())
			return
		}
	}
	if len(_sagemakerHubName) > 0 {
		input.HubName = aws.String(_sagemakerHubName)
	}

	if resp, err := client.DeleteHubContentReference(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Use this operation to delete a human task user interface (worker task template).
// To see a list of human task user interfaces (work task templates) in your
// account, use [ListHumanTaskUis]. When you delete a worker task template, it no longer appears
// when you call ListHumanTaskUis .
//
// [ListHumanTaskUis]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_ListHumanTaskUis.html
func sagemaker_DeleteHumanTaskUi(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DeleteHumanTaskUiInput{
		// HumanTaskUiName: *string, // Required
	}

	if len(_sagemakerHumanTaskUiName) > 0 {
		input.HumanTaskUiName = aws.String(_sagemakerHumanTaskUiName)
	}

	if resp, err := client.DeleteHumanTaskUi(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a hyperparameter tuning job. The DeleteHyperParameterTuningJob API
// deletes only the tuning job entry that was created in SageMaker when you called
// the CreateHyperParameterTuningJob API. It does not delete training jobs,
// artifacts, or the IAM role that you specified when creating the model.
func sagemaker_DeleteHyperParameterTuningJob(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DeleteHyperParameterTuningJobInput{
		// HyperParameterTuningJobName: *string, // Required
	}

	if len(_sagemakerHyperParameterTuningJobName) > 0 {
		input.HyperParameterTuningJobName = aws.String(_sagemakerHyperParameterTuningJobName)
	}

	if resp, err := client.DeleteHyperParameterTuningJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a SageMaker AI image and all versions of the image. The container
// images aren't deleted.
func sagemaker_DeleteImage(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DeleteImageInput{
		// ImageName: *string, // Required
	}

	if len(_sagemakerImageName) > 0 {
		input.ImageName = aws.String(_sagemakerImageName)
	}

	if resp, err := client.DeleteImage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a version of a SageMaker AI image. The container image the version
// represents isn't deleted.
func sagemaker_DeleteImageVersion(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DeleteImageVersionInput{
		// ImageName: *string, // Required
	}

	if len(_sagemakerImageName) > 0 {
		input.ImageName = aws.String(_sagemakerImageName)
	}
	if len(_sagemakerAlias) > 0 {
		input.Alias = aws.String(_sagemakerAlias)
	}
	if len(_sagemakerVersion) > 0 {
		if err := assignInputField(input, "Version", _sagemakerVersion); err != nil {
			log.Errorf("invalid --version: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteImageVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an inference component.
func sagemaker_DeleteInferenceComponent(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DeleteInferenceComponentInput{
		// InferenceComponentName: *string, // Required
	}

	if len(_sagemakerInferenceComponentName) > 0 {
		input.InferenceComponentName = aws.String(_sagemakerInferenceComponentName)
	}

	if resp, err := client.DeleteInferenceComponent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an inference experiment.
// This operation does not delete your endpoint, variants, or any underlying
// resources. This operation only deletes the metadata of your experiment.
func sagemaker_DeleteInferenceExperiment(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DeleteInferenceExperimentInput{
		// Name: *string, // Required
	}

	if len(_sagemakerName) > 0 {
		input.Name = aws.String(_sagemakerName)
	}

	if resp, err := client.DeleteInferenceExperiment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an MLflow App.
func sagemaker_DeleteMlflowApp(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DeleteMlflowAppInput{
		// Arn: *string, // Required
	}

	if len(_sagemakerArn) > 0 {
		input.Arn = aws.String(_sagemakerArn)
	}

	if resp, err := client.DeleteMlflowApp(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an MLflow Tracking Server. For more information, see [Clean up MLflow resources].
//
// [Clean up MLflow resources]: https://docs.aws.amazon.com/sagemaker/latest/dg/mlflow-cleanup.html.html
func sagemaker_DeleteMlflowTrackingServer(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DeleteMlflowTrackingServerInput{
		// TrackingServerName: *string, // Required
	}

	if len(_sagemakerTrackingServerName) > 0 {
		input.TrackingServerName = aws.String(_sagemakerTrackingServerName)
	}

	if resp, err := client.DeleteMlflowTrackingServer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a model. The DeleteModel API deletes only the model entry that was
// created in SageMaker when you called the CreateModel API. It does not delete
// model artifacts, inference code, or the IAM role that you specified when
// creating the model.
func sagemaker_DeleteModel(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DeleteModelInput{
		// ModelName: *string, // Required
	}

	if len(_sagemakerModelName) > 0 {
		input.ModelName = aws.String(_sagemakerModelName)
	}

	if resp, err := client.DeleteModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Amazon SageMaker AI model bias job definition.
func sagemaker_DeleteModelBiasJobDefinition(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DeleteModelBiasJobDefinitionInput{
		// JobDefinitionName: *string, // Required
	}

	if len(_sagemakerJobDefinitionName) > 0 {
		input.JobDefinitionName = aws.String(_sagemakerJobDefinitionName)
	}

	if resp, err := client.DeleteModelBiasJobDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Amazon SageMaker Model Card.
func sagemaker_DeleteModelCard(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DeleteModelCardInput{
		// ModelCardName: *string, // Required
	}

	if len(_sagemakerModelCardName) > 0 {
		input.ModelCardName = aws.String(_sagemakerModelCardName)
	}

	if resp, err := client.DeleteModelCard(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Amazon SageMaker AI model explainability job definition.
func sagemaker_DeleteModelExplainabilityJobDefinition(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DeleteModelExplainabilityJobDefinitionInput{
		// JobDefinitionName: *string, // Required
	}

	if len(_sagemakerJobDefinitionName) > 0 {
		input.JobDefinitionName = aws.String(_sagemakerJobDefinitionName)
	}

	if resp, err := client.DeleteModelExplainabilityJobDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a model package.
// A model package is used to create SageMaker models or list on Amazon Web
// Services Marketplace. Buyers can subscribe to model packages listed on Amazon
// Web Services Marketplace to create models in SageMaker.
func sagemaker_DeleteModelPackage(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DeleteModelPackageInput{
		// ModelPackageName: *string, // Required
	}

	if len(_sagemakerModelPackageName) > 0 {
		input.ModelPackageName = aws.String(_sagemakerModelPackageName)
	}

	if resp, err := client.DeleteModelPackage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified model group.
func sagemaker_DeleteModelPackageGroup(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DeleteModelPackageGroupInput{
		// ModelPackageGroupName: *string, // Required
	}

	if len(_sagemakerModelPackageGroupName) > 0 {
		input.ModelPackageGroupName = aws.String(_sagemakerModelPackageGroupName)
	}

	if resp, err := client.DeleteModelPackageGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a model group resource policy.
func sagemaker_DeleteModelPackageGroupPolicy(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DeleteModelPackageGroupPolicyInput{
		// ModelPackageGroupName: *string, // Required
	}

	if len(_sagemakerModelPackageGroupName) > 0 {
		input.ModelPackageGroupName = aws.String(_sagemakerModelPackageGroupName)
	}

	if resp, err := client.DeleteModelPackageGroupPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the secified model quality monitoring job definition.
func sagemaker_DeleteModelQualityJobDefinition(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DeleteModelQualityJobDefinitionInput{
		// JobDefinitionName: *string, // Required
	}

	if len(_sagemakerJobDefinitionName) > 0 {
		input.JobDefinitionName = aws.String(_sagemakerJobDefinitionName)
	}

	if resp, err := client.DeleteModelQualityJobDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a monitoring schedule. Also stops the schedule had not already been
// stopped. This does not delete the job execution history of the monitoring
// schedule.
func sagemaker_DeleteMonitoringSchedule(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DeleteMonitoringScheduleInput{
		// MonitoringScheduleName: *string, // Required
	}

	if len(_sagemakerMonitoringScheduleName) > 0 {
		input.MonitoringScheduleName = aws.String(_sagemakerMonitoringScheduleName)
	}

	if resp, err := client.DeleteMonitoringSchedule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an SageMaker AI notebook instance. Before you can delete a notebook
// instance, you must call the StopNotebookInstance API.
//
// When you delete a notebook instance, you lose all of your data. SageMaker AI
// removes the ML compute instance, and deletes the ML storage volume and the
// network interface associated with the notebook instance.
func sagemaker_DeleteNotebookInstance(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DeleteNotebookInstanceInput{
		// NotebookInstanceName: *string, // Required
	}

	if len(_sagemakerNotebookInstanceName) > 0 {
		input.NotebookInstanceName = aws.String(_sagemakerNotebookInstanceName)
	}

	if resp, err := client.DeleteNotebookInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a notebook instance lifecycle configuration.
func sagemaker_DeleteNotebookInstanceLifecycleConfig(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DeleteNotebookInstanceLifecycleConfigInput{
		// NotebookInstanceLifecycleConfigName: *string, // Required
	}

	if len(_sagemakerNotebookInstanceLifecycleConfigName) > 0 {
		input.NotebookInstanceLifecycleConfigName = aws.String(_sagemakerNotebookInstanceLifecycleConfigName)
	}

	if resp, err := client.DeleteNotebookInstanceLifecycleConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an optimization job.
func sagemaker_DeleteOptimizationJob(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DeleteOptimizationJobInput{
		// OptimizationJobName: *string, // Required
	}

	if len(_sagemakerOptimizationJobName) > 0 {
		input.OptimizationJobName = aws.String(_sagemakerOptimizationJobName)
	}

	if resp, err := client.DeleteOptimizationJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a SageMaker Partner AI App.
func sagemaker_DeletePartnerApp(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DeletePartnerAppInput{
		// Arn: *string, // Required
	}

	if len(_sagemakerArn) > 0 {
		input.Arn = aws.String(_sagemakerArn)
	}
	if len(_sagemakerClientToken) > 0 {
		input.ClientToken = aws.String(_sagemakerClientToken)
	}

	if resp, err := client.DeletePartnerApp(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a pipeline if there are no running instances of the pipeline. To delete
// a pipeline, you must stop all running instances of the pipeline using the
// StopPipelineExecution API. When you delete a pipeline, all instances of the
// pipeline are deleted.
func sagemaker_DeletePipeline(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DeletePipelineInput{
		// ClientRequestToken: *string, // Required
		// PipelineName: *string, // Required
	}

	if len(_sagemakerClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_sagemakerClientRequestToken)
	}
	if len(_sagemakerPipelineName) > 0 {
		input.PipelineName = aws.String(_sagemakerPipelineName)
	}

	if resp, err := client.DeletePipeline(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a processing job. After Amazon SageMaker deletes a processing job, all
// of the metadata for the processing job is lost. You can delete only processing
// jobs that are in a terminal state ( Stopped , Failed , or Completed ). You
// cannot delete a job that is in the InProgress or Stopping state. After deleting
// the job, you can reuse its name to create another processing job.
func sagemaker_DeleteProcessingJob(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DeleteProcessingJobInput{
		// ProcessingJobName: *string, // Required
	}

	if len(_sagemakerProcessingJobName) > 0 {
		input.ProcessingJobName = aws.String(_sagemakerProcessingJobName)
	}

	if resp, err := client.DeleteProcessingJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete the specified project.
func sagemaker_DeleteProject(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DeleteProjectInput{
		// ProjectName: *string, // Required
	}

	if len(_sagemakerProjectName) > 0 {
		input.ProjectName = aws.String(_sagemakerProjectName)
	}

	if resp, err := client.DeleteProject(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Used to delete a space.
func sagemaker_DeleteSpace(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DeleteSpaceInput{
		// DomainId: *string, // Required
		// SpaceName: *string, // Required
	}

	if len(_sagemakerDomainId) > 0 {
		input.DomainId = aws.String(_sagemakerDomainId)
	}
	if len(_sagemakerSpaceName) > 0 {
		input.SpaceName = aws.String(_sagemakerSpaceName)
	}

	if resp, err := client.DeleteSpace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the Amazon SageMaker AI Studio Lifecycle Configuration. In order to
// delete the Lifecycle Configuration, there must be no running apps using the
// Lifecycle Configuration. You must also remove the Lifecycle Configuration from
// UserSettings in all Domains and UserProfiles.
func sagemaker_DeleteStudioLifecycleConfig(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DeleteStudioLifecycleConfigInput{
		// StudioLifecycleConfigName: *string, // Required
	}

	if len(_sagemakerStudioLifecycleConfigName) > 0 {
		input.StudioLifecycleConfigName = aws.String(_sagemakerStudioLifecycleConfigName)
	}

	if resp, err := client.DeleteStudioLifecycleConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified tags from an SageMaker resource.
// To list a resource's tags, use the ListTags API.
//
// When you call this API to delete tags from a hyperparameter tuning job, the
// deleted tags are not removed from training jobs that the hyperparameter tuning
// job launched before you called this API.
//
// When you call this API to delete tags from a SageMaker Domain or User Profile,
// the deleted tags are not removed from Apps that the SageMaker Domain or User
// Profile launched before you called this API.
func sagemaker_DeleteTags(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DeleteTagsInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_sagemakerResourceArn) > 0 {
		input.ResourceArn = aws.String(_sagemakerResourceArn)
	}
	if len(_sagemakerTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _sagemakerTagKeys...)
	}

	if resp, err := client.DeleteTags(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a training job. After SageMaker deletes a training job, all of the
// metadata for the training job is lost. You can delete only training jobs that
// are in a terminal state ( Stopped , Failed , or Completed ) and don't retain an
// Available[managed warm pool] . You cannot delete a job that is in the InProgress or Stopping
// state. After deleting the job, you can reuse its name to create another training
// job.
//
// [managed warm pool]: https://docs.aws.amazon.com/sagemaker/latest/dg/train-warm-pools.html
func sagemaker_DeleteTrainingJob(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DeleteTrainingJobInput{
		// TrainingJobName: *string, // Required
	}

	if len(_sagemakerTrainingJobName) > 0 {
		input.TrainingJobName = aws.String(_sagemakerTrainingJobName)
	}

	if resp, err := client.DeleteTrainingJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified trial. All trial components that make up the trial must
// be deleted first. Use the [DescribeTrialComponent]API to get the list of trial components.
//
// [DescribeTrialComponent]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DescribeTrialComponent.html
func sagemaker_DeleteTrial(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DeleteTrialInput{
		// TrialName: *string, // Required
	}

	if len(_sagemakerTrialName) > 0 {
		input.TrialName = aws.String(_sagemakerTrialName)
	}

	if resp, err := client.DeleteTrial(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified trial component. A trial component must be disassociated
// from all trials before the trial component can be deleted. To disassociate a
// trial component from a trial, call the [DisassociateTrialComponent]API.
//
// [DisassociateTrialComponent]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DisassociateTrialComponent.html
func sagemaker_DeleteTrialComponent(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DeleteTrialComponentInput{
		// TrialComponentName: *string, // Required
	}

	if len(_sagemakerTrialComponentName) > 0 {
		input.TrialComponentName = aws.String(_sagemakerTrialComponentName)
	}

	if resp, err := client.DeleteTrialComponent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a user profile. When a user profile is deleted, the user loses access
// to their EFS volume, including data, notebooks, and other artifacts.
func sagemaker_DeleteUserProfile(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DeleteUserProfileInput{
		// DomainId: *string, // Required
		// UserProfileName: *string, // Required
	}

	if len(_sagemakerDomainId) > 0 {
		input.DomainId = aws.String(_sagemakerDomainId)
	}
	if len(_sagemakerUserProfileName) > 0 {
		input.UserProfileName = aws.String(_sagemakerUserProfileName)
	}

	if resp, err := client.DeleteUserProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Use this operation to delete a workforce.
// If you want to create a new workforce in an Amazon Web Services Region where a
// workforce already exists, use this operation to delete the existing workforce
// and then use [CreateWorkforce]to create a new workforce.
//
// If a private workforce contains one or more work teams, you must use the [DeleteWorkteam]
// operation to delete all work teams before you delete the workforce. If you try
// to delete a workforce that contains one or more work teams, you will receive a
// ResourceInUse error.
//
// [CreateWorkforce]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateWorkforce.html
// [DeleteWorkteam]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DeleteWorkteam.html
func sagemaker_DeleteWorkforce(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DeleteWorkforceInput{
		// WorkforceName: *string, // Required
	}

	if len(_sagemakerWorkforceName) > 0 {
		input.WorkforceName = aws.String(_sagemakerWorkforceName)
	}

	if resp, err := client.DeleteWorkforce(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing work team. This operation can't be undone.
func sagemaker_DeleteWorkteam(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DeleteWorkteamInput{
		// WorkteamName: *string, // Required
	}

	if len(_sagemakerWorkteamName) > 0 {
		input.WorkteamName = aws.String(_sagemakerWorkteamName)
	}

	if resp, err := client.DeleteWorkteam(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deregisters the specified devices. After you deregister a device, you will need
// to re-register the devices.
func sagemaker_DeregisterDevices(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DeregisterDevicesInput{
		// DeviceFleetName: *string, // Required
		// DeviceNames: []string, // Required
	}

	if len(_sagemakerDeviceFleetName) > 0 {
		input.DeviceFleetName = aws.String(_sagemakerDeviceFleetName)
	}
	if len(_sagemakerDeviceNames) > 0 {
		input.DeviceNames = append([]string(nil), _sagemakerDeviceNames...)
	}

	if resp, err := client.DeregisterDevices(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes an action.
func sagemaker_DescribeAction(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeActionInput{
		// ActionName: *string, // Required
	}

	if len(_sagemakerActionName) > 0 {
		input.ActionName = aws.String(_sagemakerActionName)
	}

	if resp, err := client.DescribeAction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a description of the specified algorithm that is in your account.
func sagemaker_DescribeAlgorithm(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeAlgorithmInput{
		// AlgorithmName: *string, // Required
	}

	if len(_sagemakerAlgorithmName) > 0 {
		input.AlgorithmName = aws.String(_sagemakerAlgorithmName)
	}

	if resp, err := client.DescribeAlgorithm(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the app.
func sagemaker_DescribeApp(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeAppInput{
		// AppName: *string, // Required
		// AppType: types.AppType, // Required
		// DomainId: *string, // Required
	}

	if len(_sagemakerAppName) > 0 {
		input.AppName = aws.String(_sagemakerAppName)
	}
	if len(_sagemakerAppType) > 0 {
		if err := assignInputField(input, "AppType", _sagemakerAppType); err != nil {
			log.Errorf("invalid --app-type: %s", err.Error())
			return
		}
	}
	if len(_sagemakerDomainId) > 0 {
		input.DomainId = aws.String(_sagemakerDomainId)
	}
	if len(_sagemakerSpaceName) > 0 {
		input.SpaceName = aws.String(_sagemakerSpaceName)
	}
	if len(_sagemakerUserProfileName) > 0 {
		input.UserProfileName = aws.String(_sagemakerUserProfileName)
	}

	if resp, err := client.DescribeApp(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes an AppImageConfig.
func sagemaker_DescribeAppImageConfig(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeAppImageConfigInput{
		// AppImageConfigName: *string, // Required
	}

	if len(_sagemakerAppImageConfigName) > 0 {
		input.AppImageConfigName = aws.String(_sagemakerAppImageConfigName)
	}

	if resp, err := client.DescribeAppImageConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes an artifact.
func sagemaker_DescribeArtifact(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeArtifactInput{
		// ArtifactArn: *string, // Required
	}

	if len(_sagemakerArtifactArn) > 0 {
		input.ArtifactArn = aws.String(_sagemakerArtifactArn)
	}

	if resp, err := client.DescribeArtifact(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about an AutoML job created by calling [CreateAutoMLJob].
// AutoML jobs created by calling [CreateAutoMLJobV2] cannot be described by DescribeAutoMLJob .
//
// [CreateAutoMLJobV2]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateAutoMLJobV2.html
// [CreateAutoMLJob]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateAutoMLJob.html
func sagemaker_DescribeAutoMLJob(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeAutoMLJobInput{
		// AutoMLJobName: *string, // Required
	}

	if len(_sagemakerAutoMLJobName) > 0 {
		input.AutoMLJobName = aws.String(_sagemakerAutoMLJobName)
	}

	if resp, err := client.DescribeAutoMLJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about an AutoML job created by calling [CreateAutoMLJobV2] or [CreateAutoMLJob].
//
// [CreateAutoMLJob]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateAutoMLJob.html
// [CreateAutoMLJobV2]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateAutoMLJobV2.html
func sagemaker_DescribeAutoMLJobV2(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeAutoMLJobV2Input{
		// AutoMLJobName: *string, // Required
	}

	if len(_sagemakerAutoMLJobName) > 0 {
		input.AutoMLJobName = aws.String(_sagemakerAutoMLJobName)
	}

	if resp, err := client.DescribeAutoMLJobV2(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information of a SageMaker HyperPod cluster.
func sagemaker_DescribeCluster(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeClusterInput{
		// ClusterName: *string, // Required
	}

	if len(_sagemakerClusterName) > 0 {
		input.ClusterName = aws.String(_sagemakerClusterName)
	}

	if resp, err := client.DescribeCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves detailed information about a specific event for a given HyperPod
// cluster. This functionality is only supported when the NodeProvisioningMode is
// set to Continuous .
func sagemaker_DescribeClusterEvent(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeClusterEventInput{
		// ClusterName: *string, // Required
		// EventId: *string, // Required
	}

	if len(_sagemakerClusterName) > 0 {
		input.ClusterName = aws.String(_sagemakerClusterName)
	}
	if len(_sagemakerEventId) > 0 {
		input.EventId = aws.String(_sagemakerEventId)
	}

	if resp, err := client.DescribeClusterEvent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information of a node (also called a instance interchangeably) of a
// SageMaker HyperPod cluster.
func sagemaker_DescribeClusterNode(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeClusterNodeInput{
		// ClusterName: *string, // Required
	}

	if len(_sagemakerClusterName) > 0 {
		input.ClusterName = aws.String(_sagemakerClusterName)
	}
	if len(_sagemakerNodeId) > 0 {
		input.NodeId = aws.String(_sagemakerNodeId)
	}
	if len(_sagemakerNodeLogicalId) > 0 {
		input.NodeLogicalId = aws.String(_sagemakerNodeLogicalId)
	}

	if resp, err := client.DescribeClusterNode(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Description of the cluster policy. This policy is used for task prioritization
// and fair-share allocation. This helps prioritize critical workloads and
// distributes idle compute across entities.
func sagemaker_DescribeClusterSchedulerConfig(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeClusterSchedulerConfigInput{
		// ClusterSchedulerConfigId: *string, // Required
	}

	if len(_sagemakerClusterSchedulerConfigId) > 0 {
		input.ClusterSchedulerConfigId = aws.String(_sagemakerClusterSchedulerConfigId)
	}
	if len(_sagemakerClusterSchedulerConfigVersion) > 0 {
		if err := assignInputField(input, "ClusterSchedulerConfigVersion", _sagemakerClusterSchedulerConfigVersion); err != nil {
			log.Errorf("invalid --cluster-scheduler-config-version: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeClusterSchedulerConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets details about the specified Git repository.
func sagemaker_DescribeCodeRepository(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeCodeRepositoryInput{
		// CodeRepositoryName: *string, // Required
	}

	if len(_sagemakerCodeRepositoryName) > 0 {
		input.CodeRepositoryName = aws.String(_sagemakerCodeRepositoryName)
	}

	if resp, err := client.DescribeCodeRepository(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a model compilation job.
// To create a model compilation job, use [CreateCompilationJob]. To get information about multiple
// model compilation jobs, use [ListCompilationJobs].
//
// [CreateCompilationJob]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateCompilationJob.html
// [ListCompilationJobs]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_ListCompilationJobs.html
func sagemaker_DescribeCompilationJob(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeCompilationJobInput{
		// CompilationJobName: *string, // Required
	}

	if len(_sagemakerCompilationJobName) > 0 {
		input.CompilationJobName = aws.String(_sagemakerCompilationJobName)
	}

	if resp, err := client.DescribeCompilationJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Description of the compute allocation definition.
func sagemaker_DescribeComputeQuota(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeComputeQuotaInput{
		// ComputeQuotaId: *string, // Required
	}

	if len(_sagemakerComputeQuotaId) > 0 {
		input.ComputeQuotaId = aws.String(_sagemakerComputeQuotaId)
	}
	if len(_sagemakerComputeQuotaVersion) > 0 {
		if err := assignInputField(input, "ComputeQuotaVersion", _sagemakerComputeQuotaVersion); err != nil {
			log.Errorf("invalid --compute-quota-version: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeComputeQuota(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a context.
func sagemaker_DescribeContext(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeContextInput{
		// ContextName: *string, // Required
	}

	if len(_sagemakerContextName) > 0 {
		input.ContextName = aws.String(_sagemakerContextName)
	}

	if resp, err := client.DescribeContext(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the details of a data quality monitoring job definition.
func sagemaker_DescribeDataQualityJobDefinition(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeDataQualityJobDefinitionInput{
		// JobDefinitionName: *string, // Required
	}

	if len(_sagemakerJobDefinitionName) > 0 {
		input.JobDefinitionName = aws.String(_sagemakerJobDefinitionName)
	}

	if resp, err := client.DescribeDataQualityJobDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the device.
func sagemaker_DescribeDevice(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeDeviceInput{
		// DeviceFleetName: *string, // Required
		// DeviceName: *string, // Required
	}

	if len(_sagemakerDeviceFleetName) > 0 {
		input.DeviceFleetName = aws.String(_sagemakerDeviceFleetName)
	}
	if len(_sagemakerDeviceName) > 0 {
		input.DeviceName = aws.String(_sagemakerDeviceName)
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}

	if resp, err := client.DescribeDevice(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// A description of the fleet the device belongs to.
func sagemaker_DescribeDeviceFleet(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeDeviceFleetInput{
		// DeviceFleetName: *string, // Required
	}

	if len(_sagemakerDeviceFleetName) > 0 {
		input.DeviceFleetName = aws.String(_sagemakerDeviceFleetName)
	}

	if resp, err := client.DescribeDeviceFleet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The description of the domain.
func sagemaker_DescribeDomain(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeDomainInput{
		// DomainId: *string, // Required
	}

	if len(_sagemakerDomainId) > 0 {
		input.DomainId = aws.String(_sagemakerDomainId)
	}

	if resp, err := client.DescribeDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes an edge deployment plan with deployment status per stage.
func sagemaker_DescribeEdgeDeploymentPlan(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeEdgeDeploymentPlanInput{
		// EdgeDeploymentPlanName: *string, // Required
	}

	if len(_sagemakerEdgeDeploymentPlanName) > 0 {
		input.EdgeDeploymentPlanName = aws.String(_sagemakerEdgeDeploymentPlanName)
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}

	if resp, err := client.DescribeEdgeDeploymentPlan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// A description of edge packaging jobs.
func sagemaker_DescribeEdgePackagingJob(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeEdgePackagingJobInput{
		// EdgePackagingJobName: *string, // Required
	}

	if len(_sagemakerEdgePackagingJobName) > 0 {
		input.EdgePackagingJobName = aws.String(_sagemakerEdgePackagingJobName)
	}

	if resp, err := client.DescribeEdgePackagingJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the description of an endpoint.
func sagemaker_DescribeEndpoint(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeEndpointInput{
		// EndpointName: *string, // Required
	}

	if len(_sagemakerEndpointName) > 0 {
		input.EndpointName = aws.String(_sagemakerEndpointName)
	}

	if resp, err := client.DescribeEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the description of an endpoint configuration created using the
// CreateEndpointConfig API.
func sagemaker_DescribeEndpointConfig(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeEndpointConfigInput{
		// EndpointConfigName: *string, // Required
	}

	if len(_sagemakerEndpointConfigName) > 0 {
		input.EndpointConfigName = aws.String(_sagemakerEndpointConfigName)
	}

	if resp, err := client.DescribeEndpointConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides a list of an experiment's properties.
func sagemaker_DescribeExperiment(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeExperimentInput{
		// ExperimentName: *string, // Required
	}

	if len(_sagemakerExperimentName) > 0 {
		input.ExperimentName = aws.String(_sagemakerExperimentName)
	}

	if resp, err := client.DescribeExperiment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Use this operation to describe a FeatureGroup . The response includes
// information on the creation time, FeatureGroup name, the unique identifier for
// each FeatureGroup , and more.
func sagemaker_DescribeFeatureGroup(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeFeatureGroupInput{
		// FeatureGroupName: *string, // Required
	}

	if len(_sagemakerFeatureGroupName) > 0 {
		input.FeatureGroupName = aws.String(_sagemakerFeatureGroupName)
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}

	if resp, err := client.DescribeFeatureGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Shows the metadata for a feature within a feature group.
func sagemaker_DescribeFeatureMetadata(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeFeatureMetadataInput{
		// FeatureGroupName: *string, // Required
		// FeatureName: *string, // Required
	}

	if len(_sagemakerFeatureGroupName) > 0 {
		input.FeatureGroupName = aws.String(_sagemakerFeatureGroupName)
	}
	if len(_sagemakerFeatureName) > 0 {
		input.FeatureName = aws.String(_sagemakerFeatureName)
	}

	if resp, err := client.DescribeFeatureMetadata(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the specified flow definition.
func sagemaker_DescribeFlowDefinition(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeFlowDefinitionInput{
		// FlowDefinitionName: *string, // Required
	}

	if len(_sagemakerFlowDefinitionName) > 0 {
		input.FlowDefinitionName = aws.String(_sagemakerFlowDefinitionName)
	}

	if resp, err := client.DescribeFlowDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a hub.
func sagemaker_DescribeHub(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeHubInput{
		// HubName: *string, // Required
	}

	if len(_sagemakerHubName) > 0 {
		input.HubName = aws.String(_sagemakerHubName)
	}

	if resp, err := client.DescribeHub(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describe the content of a hub.
func sagemaker_DescribeHubContent(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeHubContentInput{
		// HubContentName: *string, // Required
		// HubContentType: types.HubContentType, // Required
		// HubName: *string, // Required
	}

	if len(_sagemakerHubContentName) > 0 {
		input.HubContentName = aws.String(_sagemakerHubContentName)
	}
	if len(_sagemakerHubContentType) > 0 {
		if err := assignInputField(input, "HubContentType", _sagemakerHubContentType); err != nil {
			log.Errorf("invalid --hub-content-type: %s", err.Error())
			return
		}
	}
	if len(_sagemakerHubName) > 0 {
		input.HubName = aws.String(_sagemakerHubName)
	}
	if len(_sagemakerHubContentVersion) > 0 {
		input.HubContentVersion = aws.String(_sagemakerHubContentVersion)
	}

	if resp, err := client.DescribeHubContent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the requested human task user interface (worker task
// template).
func sagemaker_DescribeHumanTaskUi(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeHumanTaskUiInput{
		// HumanTaskUiName: *string, // Required
	}

	if len(_sagemakerHumanTaskUiName) > 0 {
		input.HumanTaskUiName = aws.String(_sagemakerHumanTaskUiName)
	}

	if resp, err := client.DescribeHumanTaskUi(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a description of a hyperparameter tuning job, depending on the fields
// selected. These fields can include the name, Amazon Resource Name (ARN), job
// status of your tuning job and more.
func sagemaker_DescribeHyperParameterTuningJob(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeHyperParameterTuningJobInput{
		// HyperParameterTuningJobName: *string, // Required
	}

	if len(_sagemakerHyperParameterTuningJobName) > 0 {
		input.HyperParameterTuningJobName = aws.String(_sagemakerHyperParameterTuningJobName)
	}

	if resp, err := client.DescribeHyperParameterTuningJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a SageMaker AI image.
func sagemaker_DescribeImage(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeImageInput{
		// ImageName: *string, // Required
	}

	if len(_sagemakerImageName) > 0 {
		input.ImageName = aws.String(_sagemakerImageName)
	}

	if resp, err := client.DescribeImage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a version of a SageMaker AI image.
func sagemaker_DescribeImageVersion(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeImageVersionInput{
		// ImageName: *string, // Required
	}

	if len(_sagemakerImageName) > 0 {
		input.ImageName = aws.String(_sagemakerImageName)
	}
	if len(_sagemakerAlias) > 0 {
		input.Alias = aws.String(_sagemakerAlias)
	}
	if len(_sagemakerVersion) > 0 {
		if err := assignInputField(input, "Version", _sagemakerVersion); err != nil {
			log.Errorf("invalid --version: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeImageVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about an inference component.
func sagemaker_DescribeInferenceComponent(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeInferenceComponentInput{
		// InferenceComponentName: *string, // Required
	}

	if len(_sagemakerInferenceComponentName) > 0 {
		input.InferenceComponentName = aws.String(_sagemakerInferenceComponentName)
	}

	if resp, err := client.DescribeInferenceComponent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns details about an inference experiment.
func sagemaker_DescribeInferenceExperiment(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeInferenceExperimentInput{
		// Name: *string, // Required
	}

	if len(_sagemakerName) > 0 {
		input.Name = aws.String(_sagemakerName)
	}

	if resp, err := client.DescribeInferenceExperiment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides the results of the Inference Recommender job. One or more
// recommendation jobs are returned.
func sagemaker_DescribeInferenceRecommendationsJob(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeInferenceRecommendationsJobInput{
		// JobName: *string, // Required
	}

	if len(_sagemakerJobName) > 0 {
		input.JobName = aws.String(_sagemakerJobName)
	}

	if resp, err := client.DescribeInferenceRecommendationsJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a labeling job.
func sagemaker_DescribeLabelingJob(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeLabelingJobInput{
		// LabelingJobName: *string, // Required
	}

	if len(_sagemakerLabelingJobName) > 0 {
		input.LabelingJobName = aws.String(_sagemakerLabelingJobName)
	}

	if resp, err := client.DescribeLabelingJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides a list of properties for the requested lineage group. For more
// information, see [Cross-Account Lineage Tracking]in the Amazon SageMaker Developer Guide.
//
// [Cross-Account Lineage Tracking]: https://docs.aws.amazon.com/sagemaker/latest/dg/xaccount-lineage-tracking.html
func sagemaker_DescribeLineageGroup(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeLineageGroupInput{
		// LineageGroupName: *string, // Required
	}

	if len(_sagemakerLineageGroupName) > 0 {
		input.LineageGroupName = aws.String(_sagemakerLineageGroupName)
	}

	if resp, err := client.DescribeLineageGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about an MLflow App.
func sagemaker_DescribeMlflowApp(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeMlflowAppInput{
		// Arn: *string, // Required
	}

	if len(_sagemakerArn) > 0 {
		input.Arn = aws.String(_sagemakerArn)
	}

	if resp, err := client.DescribeMlflowApp(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about an MLflow Tracking Server.
func sagemaker_DescribeMlflowTrackingServer(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeMlflowTrackingServerInput{
		// TrackingServerName: *string, // Required
	}

	if len(_sagemakerTrackingServerName) > 0 {
		input.TrackingServerName = aws.String(_sagemakerTrackingServerName)
	}

	if resp, err := client.DescribeMlflowTrackingServer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a model that you created using the CreateModel API.
func sagemaker_DescribeModel(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeModelInput{
		// ModelName: *string, // Required
	}

	if len(_sagemakerModelName) > 0 {
		input.ModelName = aws.String(_sagemakerModelName)
	}

	if resp, err := client.DescribeModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a description of a model bias job definition.
func sagemaker_DescribeModelBiasJobDefinition(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeModelBiasJobDefinitionInput{
		// JobDefinitionName: *string, // Required
	}

	if len(_sagemakerJobDefinitionName) > 0 {
		input.JobDefinitionName = aws.String(_sagemakerJobDefinitionName)
	}

	if resp, err := client.DescribeModelBiasJobDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the content, creation time, and security configuration of an Amazon
// SageMaker Model Card.
func sagemaker_DescribeModelCard(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeModelCardInput{
		// ModelCardName: *string, // Required
	}

	if len(_sagemakerModelCardName) > 0 {
		input.ModelCardName = aws.String(_sagemakerModelCardName)
	}
	if len(_sagemakerModelCardVersion) > 0 {
		if err := assignInputField(input, "ModelCardVersion", _sagemakerModelCardVersion); err != nil {
			log.Errorf("invalid --model-card-version: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeModelCard(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes an Amazon SageMaker Model Card export job.
func sagemaker_DescribeModelCardExportJob(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeModelCardExportJobInput{
		// ModelCardExportJobArn: *string, // Required
	}

	if len(_sagemakerModelCardExportJobArn) > 0 {
		input.ModelCardExportJobArn = aws.String(_sagemakerModelCardExportJobArn)
	}

	if resp, err := client.DescribeModelCardExportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a description of a model explainability job definition.
func sagemaker_DescribeModelExplainabilityJobDefinition(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeModelExplainabilityJobDefinitionInput{
		// JobDefinitionName: *string, // Required
	}

	if len(_sagemakerJobDefinitionName) > 0 {
		input.JobDefinitionName = aws.String(_sagemakerJobDefinitionName)
	}

	if resp, err := client.DescribeModelExplainabilityJobDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a description of the specified model package, which is used to create
// SageMaker models or list them on Amazon Web Services Marketplace.
//
// If you provided a KMS Key ID when you created your model package, you will see
// the [KMS Decrypt]API call in your CloudTrail logs when you use this API.
//
// To create models in SageMaker, buyers can subscribe to model packages listed on
// Amazon Web Services Marketplace.
//
// [KMS Decrypt]: https://docs.aws.amazon.com/kms/latest/APIReference/API_Decrypt.html
func sagemaker_DescribeModelPackage(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeModelPackageInput{
		// ModelPackageName: *string, // Required
	}

	if len(_sagemakerModelPackageName) > 0 {
		input.ModelPackageName = aws.String(_sagemakerModelPackageName)
	}

	if resp, err := client.DescribeModelPackage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a description for the specified model group.
func sagemaker_DescribeModelPackageGroup(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeModelPackageGroupInput{
		// ModelPackageGroupName: *string, // Required
	}

	if len(_sagemakerModelPackageGroupName) > 0 {
		input.ModelPackageGroupName = aws.String(_sagemakerModelPackageGroupName)
	}

	if resp, err := client.DescribeModelPackageGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a description of a model quality job definition.
func sagemaker_DescribeModelQualityJobDefinition(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeModelQualityJobDefinitionInput{
		// JobDefinitionName: *string, // Required
	}

	if len(_sagemakerJobDefinitionName) > 0 {
		input.JobDefinitionName = aws.String(_sagemakerJobDefinitionName)
	}

	if resp, err := client.DescribeModelQualityJobDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the schedule for a monitoring job.
func sagemaker_DescribeMonitoringSchedule(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeMonitoringScheduleInput{
		// MonitoringScheduleName: *string, // Required
	}

	if len(_sagemakerMonitoringScheduleName) > 0 {
		input.MonitoringScheduleName = aws.String(_sagemakerMonitoringScheduleName)
	}

	if resp, err := client.DescribeMonitoringSchedule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a notebook instance.
func sagemaker_DescribeNotebookInstance(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeNotebookInstanceInput{
		// NotebookInstanceName: *string, // Required
	}

	if len(_sagemakerNotebookInstanceName) > 0 {
		input.NotebookInstanceName = aws.String(_sagemakerNotebookInstanceName)
	}

	if resp, err := client.DescribeNotebookInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a description of a notebook instance lifecycle configuration.
// For information about notebook instance lifestyle configurations, see [Step 2.1: (Optional) Customize a Notebook Instance].
//
// [Step 2.1: (Optional) Customize a Notebook Instance]: https://docs.aws.amazon.com/sagemaker/latest/dg/notebook-lifecycle-config.html
func sagemaker_DescribeNotebookInstanceLifecycleConfig(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeNotebookInstanceLifecycleConfigInput{
		// NotebookInstanceLifecycleConfigName: *string, // Required
	}

	if len(_sagemakerNotebookInstanceLifecycleConfigName) > 0 {
		input.NotebookInstanceLifecycleConfigName = aws.String(_sagemakerNotebookInstanceLifecycleConfigName)
	}

	if resp, err := client.DescribeNotebookInstanceLifecycleConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides the properties of the specified optimization job.
func sagemaker_DescribeOptimizationJob(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeOptimizationJobInput{
		// OptimizationJobName: *string, // Required
	}

	if len(_sagemakerOptimizationJobName) > 0 {
		input.OptimizationJobName = aws.String(_sagemakerOptimizationJobName)
	}

	if resp, err := client.DescribeOptimizationJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a SageMaker Partner AI App.
func sagemaker_DescribePartnerApp(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribePartnerAppInput{
		// Arn: *string, // Required
	}

	if len(_sagemakerArn) > 0 {
		input.Arn = aws.String(_sagemakerArn)
	}
	if len(_sagemakerIncludeAvailableUpgrade) > 0 {
		if err := assignInputField(input, "IncludeAvailableUpgrade", _sagemakerIncludeAvailableUpgrade); err != nil {
			log.Errorf("invalid --include-available-upgrade: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribePartnerApp(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the details of a pipeline.
func sagemaker_DescribePipeline(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribePipelineInput{
		// PipelineName: *string, // Required
	}

	if len(_sagemakerPipelineName) > 0 {
		input.PipelineName = aws.String(_sagemakerPipelineName)
	}
	if len(_sagemakerPipelineVersionId) > 0 {
		if err := assignInputField(input, "PipelineVersionId", _sagemakerPipelineVersionId); err != nil {
			log.Errorf("invalid --pipeline-version-id: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribePipeline(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the details of an execution's pipeline definition.
func sagemaker_DescribePipelineDefinitionForExecution(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribePipelineDefinitionForExecutionInput{
		// PipelineExecutionArn: *string, // Required
	}

	if len(_sagemakerPipelineExecutionArn) > 0 {
		input.PipelineExecutionArn = aws.String(_sagemakerPipelineExecutionArn)
	}

	if resp, err := client.DescribePipelineDefinitionForExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the details of a pipeline execution.
func sagemaker_DescribePipelineExecution(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribePipelineExecutionInput{
		// PipelineExecutionArn: *string, // Required
	}

	if len(_sagemakerPipelineExecutionArn) > 0 {
		input.PipelineExecutionArn = aws.String(_sagemakerPipelineExecutionArn)
	}

	if resp, err := client.DescribePipelineExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a description of a processing job.
func sagemaker_DescribeProcessingJob(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeProcessingJobInput{
		// ProcessingJobName: *string, // Required
	}

	if len(_sagemakerProcessingJobName) > 0 {
		input.ProcessingJobName = aws.String(_sagemakerProcessingJobName)
	}

	if resp, err := client.DescribeProcessingJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the details of a project.
func sagemaker_DescribeProject(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeProjectInput{
		// ProjectName: *string, // Required
	}

	if len(_sagemakerProjectName) > 0 {
		input.ProjectName = aws.String(_sagemakerProjectName)
	}

	if resp, err := client.DescribeProject(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves details about a reserved capacity.
func sagemaker_DescribeReservedCapacity(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeReservedCapacityInput{
		// ReservedCapacityArn: *string, // Required
	}

	if len(_sagemakerReservedCapacityArn) > 0 {
		input.ReservedCapacityArn = aws.String(_sagemakerReservedCapacityArn)
	}

	if resp, err := client.DescribeReservedCapacity(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the space.
func sagemaker_DescribeSpace(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeSpaceInput{
		// DomainId: *string, // Required
		// SpaceName: *string, // Required
	}

	if len(_sagemakerDomainId) > 0 {
		input.DomainId = aws.String(_sagemakerDomainId)
	}
	if len(_sagemakerSpaceName) > 0 {
		input.SpaceName = aws.String(_sagemakerSpaceName)
	}

	if resp, err := client.DescribeSpace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the Amazon SageMaker AI Studio Lifecycle Configuration.
func sagemaker_DescribeStudioLifecycleConfig(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeStudioLifecycleConfigInput{
		// StudioLifecycleConfigName: *string, // Required
	}

	if len(_sagemakerStudioLifecycleConfigName) > 0 {
		input.StudioLifecycleConfigName = aws.String(_sagemakerStudioLifecycleConfigName)
	}

	if resp, err := client.DescribeStudioLifecycleConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a work team provided by a vendor. It returns details
// about the subscription with a vendor in the Amazon Web Services Marketplace.
func sagemaker_DescribeSubscribedWorkteam(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeSubscribedWorkteamInput{
		// WorkteamArn: *string, // Required
	}

	if len(_sagemakerWorkteamArn) > 0 {
		input.WorkteamArn = aws.String(_sagemakerWorkteamArn)
	}

	if resp, err := client.DescribeSubscribedWorkteam(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a training job.
// Some of the attributes below only appear if the training job successfully
// starts. If the training job fails, TrainingJobStatus is Failed and, depending
// on the FailureReason , attributes like TrainingStartTime , TrainingTimeInSeconds
// , TrainingEndTime , and BillableTimeInSeconds may not be present in the
// response.
func sagemaker_DescribeTrainingJob(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeTrainingJobInput{
		// TrainingJobName: *string, // Required
	}

	if len(_sagemakerTrainingJobName) > 0 {
		input.TrainingJobName = aws.String(_sagemakerTrainingJobName)
	}

	if resp, err := client.DescribeTrainingJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves detailed information about a specific training plan.
func sagemaker_DescribeTrainingPlan(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeTrainingPlanInput{
		// TrainingPlanName: *string, // Required
	}

	if len(_sagemakerTrainingPlanName) > 0 {
		input.TrainingPlanName = aws.String(_sagemakerTrainingPlanName)
	}

	if resp, err := client.DescribeTrainingPlan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a transform job.
func sagemaker_DescribeTransformJob(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeTransformJobInput{
		// TransformJobName: *string, // Required
	}

	if len(_sagemakerTransformJobName) > 0 {
		input.TransformJobName = aws.String(_sagemakerTransformJobName)
	}

	if resp, err := client.DescribeTransformJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides a list of a trial's properties.
func sagemaker_DescribeTrial(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeTrialInput{
		// TrialName: *string, // Required
	}

	if len(_sagemakerTrialName) > 0 {
		input.TrialName = aws.String(_sagemakerTrialName)
	}

	if resp, err := client.DescribeTrial(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides a list of a trials component's properties.
func sagemaker_DescribeTrialComponent(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeTrialComponentInput{
		// TrialComponentName: *string, // Required
	}

	if len(_sagemakerTrialComponentName) > 0 {
		input.TrialComponentName = aws.String(_sagemakerTrialComponentName)
	}

	if resp, err := client.DescribeTrialComponent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a user profile. For more information, see CreateUserProfile .
func sagemaker_DescribeUserProfile(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeUserProfileInput{
		// DomainId: *string, // Required
		// UserProfileName: *string, // Required
	}

	if len(_sagemakerDomainId) > 0 {
		input.DomainId = aws.String(_sagemakerDomainId)
	}
	if len(_sagemakerUserProfileName) > 0 {
		input.UserProfileName = aws.String(_sagemakerUserProfileName)
	}

	if resp, err := client.DescribeUserProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists private workforce information, including workforce name, Amazon Resource
// Name (ARN), and, if applicable, allowed IP address ranges ([CIDRs] ). Allowable IP
// address ranges are the IP addresses that workers can use to access tasks.
//
// This operation applies only to private workforces.
//
// [CIDRs]: https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Subnets.html
func sagemaker_DescribeWorkforce(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeWorkforceInput{
		// WorkforceName: *string, // Required
	}

	if len(_sagemakerWorkforceName) > 0 {
		input.WorkforceName = aws.String(_sagemakerWorkforceName)
	}

	if resp, err := client.DescribeWorkforce(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a specific work team. You can see information such as
// the creation date, the last updated date, membership information, and the work
// team's Amazon Resource Name (ARN).
func sagemaker_DescribeWorkteam(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DescribeWorkteamInput{
		// WorkteamName: *string, // Required
	}

	if len(_sagemakerWorkteamName) > 0 {
		input.WorkteamName = aws.String(_sagemakerWorkteamName)
	}

	if resp, err := client.DescribeWorkteam(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Detaches your Amazon Elastic Block Store (Amazon EBS) volume from a node in
// your EKS orchestrated SageMaker HyperPod cluster.
//
// This API works with the Amazon Elastic Block Store (Amazon EBS) Container
// Storage Interface (CSI) driver to manage the lifecycle of persistent storage in
// your HyperPod EKS clusters.
func sagemaker_DetachClusterNodeVolume(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DetachClusterNodeVolumeInput{
		// ClusterArn: *string, // Required
		// NodeId: *string, // Required
		// VolumeId: *string, // Required
	}

	if len(_sagemakerClusterArn) > 0 {
		input.ClusterArn = aws.String(_sagemakerClusterArn)
	}
	if len(_sagemakerNodeId) > 0 {
		input.NodeId = aws.String(_sagemakerNodeId)
	}
	if len(_sagemakerVolumeId) > 0 {
		input.VolumeId = aws.String(_sagemakerVolumeId)
	}

	if resp, err := client.DetachClusterNodeVolume(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disables using Service Catalog in SageMaker. Service Catalog is used to create
// SageMaker projects.
func sagemaker_DisableSagemakerServicecatalogPortfolio(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DisableSagemakerServicecatalogPortfolioInput{}

	if resp, err := client.DisableSagemakerServicecatalogPortfolio(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates a trial component from a trial. This doesn't effect other trials
// the component is associated with. Before you can delete a component, you must
// disassociate the component from all trials it is associated with. To associate a
// trial component with a trial, call the [AssociateTrialComponent]API.
//
// To get a list of the trials a component is associated with, use the [Search] API.
// Specify ExperimentTrialComponent for the Resource parameter. The list appears
// in the response under Results.TrialComponent.Parents .
//
// [Search]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_Search.html
// [AssociateTrialComponent]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_AssociateTrialComponent.html
func sagemaker_DisassociateTrialComponent(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.DisassociateTrialComponentInput{
		// TrialComponentName: *string, // Required
		// TrialName: *string, // Required
	}

	if len(_sagemakerTrialComponentName) > 0 {
		input.TrialComponentName = aws.String(_sagemakerTrialComponentName)
	}
	if len(_sagemakerTrialName) > 0 {
		input.TrialName = aws.String(_sagemakerTrialName)
	}

	if resp, err := client.DisassociateTrialComponent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables using Service Catalog in SageMaker. Service Catalog is used to create
// SageMaker projects.
func sagemaker_EnableSagemakerServicecatalogPortfolio(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.EnableSagemakerServicecatalogPortfolioInput{}

	if resp, err := client.EnableSagemakerServicecatalogPortfolio(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a fleet.
func sagemaker_GetDeviceFleetReport(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.GetDeviceFleetReportInput{
		// DeviceFleetName: *string, // Required
	}

	if len(_sagemakerDeviceFleetName) > 0 {
		input.DeviceFleetName = aws.String(_sagemakerDeviceFleetName)
	}

	if resp, err := client.GetDeviceFleetReport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The resource policy for the lineage group.
func sagemaker_GetLineageGroupPolicy(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.GetLineageGroupPolicyInput{
		// LineageGroupName: *string, // Required
	}

	if len(_sagemakerLineageGroupName) > 0 {
		input.LineageGroupName = aws.String(_sagemakerLineageGroupName)
	}

	if resp, err := client.GetLineageGroupPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a resource policy that manages access for a model group. For information
// about resource policies, see [Identity-based policies and resource-based policies]in the Amazon Web Services Identity and Access
// Management User Guide..
//
// [Identity-based policies and resource-based policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_identity-vs-resource.html
func sagemaker_GetModelPackageGroupPolicy(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.GetModelPackageGroupPolicyInput{
		// ModelPackageGroupName: *string, // Required
	}

	if len(_sagemakerModelPackageGroupName) > 0 {
		input.ModelPackageGroupName = aws.String(_sagemakerModelPackageGroupName)
	}

	if resp, err := client.GetModelPackageGroupPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the status of Service Catalog in SageMaker. Service Catalog is used to
// create SageMaker projects.
func sagemaker_GetSagemakerServicecatalogPortfolioStatus(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.GetSagemakerServicecatalogPortfolioStatusInput{}

	if resp, err := client.GetSagemakerServicecatalogPortfolioStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts an Amazon SageMaker Inference Recommender autoscaling recommendation
// job. Returns recommendations for autoscaling policies that you can apply to your
// SageMaker endpoint.
func sagemaker_GetScalingConfigurationRecommendation(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.GetScalingConfigurationRecommendationInput{
		// InferenceRecommendationsJobName: *string, // Required
	}

	if len(_sagemakerInferenceRecommendationsJobName) > 0 {
		input.InferenceRecommendationsJobName = aws.String(_sagemakerInferenceRecommendationsJobName)
	}
	if len(_sagemakerEndpointName) > 0 {
		input.EndpointName = aws.String(_sagemakerEndpointName)
	}
	if len(_sagemakerRecommendationId) > 0 {
		input.RecommendationId = aws.String(_sagemakerRecommendationId)
	}
	if len(_sagemakerScalingPolicyObjective) > 0 {
		if err := assignInputField(input, "ScalingPolicyObjective", _sagemakerScalingPolicyObjective); err != nil {
			log.Errorf("invalid --scaling-policy-objective: %s", err.Error())
			return
		}
	}
	if len(_sagemakerTargetCpuUtilizationPerCore) > 0 {
		if err := assignInputField(input, "TargetCpuUtilizationPerCore", _sagemakerTargetCpuUtilizationPerCore); err != nil {
			log.Errorf("invalid --target-cpu-utilization-per-core: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetScalingConfigurationRecommendation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// An auto-complete API for the search functionality in the SageMaker console. It
// returns suggestions of possible matches for the property name to use in Search
// queries. Provides suggestions for HyperParameters , Tags , and Metrics .
func sagemaker_GetSearchSuggestions(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.GetSearchSuggestionsInput{
		// Resource: types.ResourceType, // Required
	}

	if len(_sagemakerResource) > 0 {
		if err := assignInputField(input, "Resource", _sagemakerResource); err != nil {
			log.Errorf("invalid --resource: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSuggestionQuery) > 0 {
		if err := assignInputField(input, "SuggestionQuery", _sagemakerSuggestionQuery); err != nil {
			log.Errorf("invalid --suggestion-query: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetSearchSuggestions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Import hub content.
func sagemaker_ImportHubContent(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ImportHubContentInput{
		// DocumentSchemaVersion: *string, // Required
		// HubContentDocument: *string, // Required
		// HubContentName: *string, // Required
		// HubContentType: types.HubContentType, // Required
		// HubName: *string, // Required
	}

	if len(_sagemakerDocumentSchemaVersion) > 0 {
		input.DocumentSchemaVersion = aws.String(_sagemakerDocumentSchemaVersion)
	}
	if len(_sagemakerHubContentDocument) > 0 {
		input.HubContentDocument = aws.String(_sagemakerHubContentDocument)
	}
	if len(_sagemakerHubContentName) > 0 {
		input.HubContentName = aws.String(_sagemakerHubContentName)
	}
	if len(_sagemakerHubContentType) > 0 {
		if err := assignInputField(input, "HubContentType", _sagemakerHubContentType); err != nil {
			log.Errorf("invalid --hub-content-type: %s", err.Error())
			return
		}
	}
	if len(_sagemakerHubName) > 0 {
		input.HubName = aws.String(_sagemakerHubName)
	}
	if len(_sagemakerHubContentDescription) > 0 {
		input.HubContentDescription = aws.String(_sagemakerHubContentDescription)
	}
	if len(_sagemakerHubContentDisplayName) > 0 {
		input.HubContentDisplayName = aws.String(_sagemakerHubContentDisplayName)
	}
	if len(_sagemakerHubContentMarkdown) > 0 {
		input.HubContentMarkdown = aws.String(_sagemakerHubContentMarkdown)
	}
	if len(_sagemakerHubContentSearchKeywords) > 0 {
		input.HubContentSearchKeywords = append([]string(nil), _sagemakerHubContentSearchKeywords...)
	}
	if len(_sagemakerHubContentVersion) > 0 {
		input.HubContentVersion = aws.String(_sagemakerHubContentVersion)
	}
	if len(_sagemakerSupportStatus) > 0 {
		if err := assignInputField(input, "SupportStatus", _sagemakerSupportStatus); err != nil {
			log.Errorf("invalid --support-status: %s", err.Error())
			return
		}
	}
	if len(_sagemakerTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.ImportHubContent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the actions in your account and their properties.
func sagemaker_ListActions(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListActionsInput{}

	if len(_sagemakerActionType) > 0 {
		input.ActionType = aws.String(_sagemakerActionType)
	}
	if len(_sagemakerCreatedAfter) > 0 {
		if err := assignInputField(input, "CreatedAfter", _sagemakerCreatedAfter); err != nil {
			log.Errorf("invalid --created-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreatedBefore) > 0 {
		if err := assignInputField(input, "CreatedBefore", _sagemakerCreatedBefore); err != nil {
			log.Errorf("invalid --created-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSourceUri) > 0 {
		input.SourceUri = aws.String(_sagemakerSourceUri)
	}

	if disablePaginator() {
		if resp, err := client.ListActions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListActionsOutput
	p := sagemaker.NewListActionsPaginator(client, input)
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

// Lists the machine learning algorithms that have been created.
func sagemaker_ListAlgorithms(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListAlgorithmsInput{}

	if len(_sagemakerCreationTimeAfter) > 0 {
		if err := assignInputField(input, "CreationTimeAfter", _sagemakerCreationTimeAfter); err != nil {
			log.Errorf("invalid --creation-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreationTimeBefore) > 0 {
		if err := assignInputField(input, "CreationTimeBefore", _sagemakerCreationTimeBefore); err != nil {
			log.Errorf("invalid --creation-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNameContains) > 0 {
		input.NameContains = aws.String(_sagemakerNameContains)
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListAlgorithms(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListAlgorithmsOutput
	p := sagemaker.NewListAlgorithmsPaginator(client, input)
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

// Lists the aliases of a specified image or image version.
func sagemaker_ListAliases(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListAliasesInput{
		// ImageName: *string, // Required
	}

	if len(_sagemakerImageName) > 0 {
		input.ImageName = aws.String(_sagemakerImageName)
	}
	if len(_sagemakerAlias) > 0 {
		input.Alias = aws.String(_sagemakerAlias)
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerVersion) > 0 {
		if err := assignInputField(input, "Version", _sagemakerVersion); err != nil {
			log.Errorf("invalid --version: %s", err.Error())
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

	var results []*sagemaker.ListAliasesOutput
	p := sagemaker.NewListAliasesPaginator(client, input)
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

// Lists the AppImageConfigs in your account and their properties. The list can be
// filtered by creation time or modified time, and whether the AppImageConfig name
// contains a specified string.
func sagemaker_ListAppImageConfigs(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListAppImageConfigsInput{}

	if len(_sagemakerCreationTimeAfter) > 0 {
		if err := assignInputField(input, "CreationTimeAfter", _sagemakerCreationTimeAfter); err != nil {
			log.Errorf("invalid --creation-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreationTimeBefore) > 0 {
		if err := assignInputField(input, "CreationTimeBefore", _sagemakerCreationTimeBefore); err != nil {
			log.Errorf("invalid --creation-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerModifiedTimeAfter) > 0 {
		if err := assignInputField(input, "ModifiedTimeAfter", _sagemakerModifiedTimeAfter); err != nil {
			log.Errorf("invalid --modified-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerModifiedTimeBefore) > 0 {
		if err := assignInputField(input, "ModifiedTimeBefore", _sagemakerModifiedTimeBefore); err != nil {
			log.Errorf("invalid --modified-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNameContains) > 0 {
		input.NameContains = aws.String(_sagemakerNameContains)
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListAppImageConfigs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListAppImageConfigsOutput
	p := sagemaker.NewListAppImageConfigsPaginator(client, input)
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

// Lists apps.
func sagemaker_ListApps(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListAppsInput{}

	if len(_sagemakerDomainIdEquals) > 0 {
		input.DomainIdEquals = aws.String(_sagemakerDomainIdEquals)
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSpaceNameEquals) > 0 {
		input.SpaceNameEquals = aws.String(_sagemakerSpaceNameEquals)
	}
	if len(_sagemakerUserProfileNameEquals) > 0 {
		input.UserProfileNameEquals = aws.String(_sagemakerUserProfileNameEquals)
	}

	if disablePaginator() {
		if resp, err := client.ListApps(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListAppsOutput
	p := sagemaker.NewListAppsPaginator(client, input)
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

// Lists the artifacts in your account and their properties.
func sagemaker_ListArtifacts(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListArtifactsInput{}

	if len(_sagemakerArtifactType) > 0 {
		input.ArtifactType = aws.String(_sagemakerArtifactType)
	}
	if len(_sagemakerCreatedAfter) > 0 {
		if err := assignInputField(input, "CreatedAfter", _sagemakerCreatedAfter); err != nil {
			log.Errorf("invalid --created-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreatedBefore) > 0 {
		if err := assignInputField(input, "CreatedBefore", _sagemakerCreatedBefore); err != nil {
			log.Errorf("invalid --created-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSourceUri) > 0 {
		input.SourceUri = aws.String(_sagemakerSourceUri)
	}

	if disablePaginator() {
		if resp, err := client.ListArtifacts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListArtifactsOutput
	p := sagemaker.NewListArtifactsPaginator(client, input)
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

// Lists the associations in your account and their properties.
func sagemaker_ListAssociations(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListAssociationsInput{}

	if len(_sagemakerAssociationType) > 0 {
		if err := assignInputField(input, "AssociationType", _sagemakerAssociationType); err != nil {
			log.Errorf("invalid --association-type: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreatedAfter) > 0 {
		if err := assignInputField(input, "CreatedAfter", _sagemakerCreatedAfter); err != nil {
			log.Errorf("invalid --created-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreatedBefore) > 0 {
		if err := assignInputField(input, "CreatedBefore", _sagemakerCreatedBefore); err != nil {
			log.Errorf("invalid --created-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerDestinationArn) > 0 {
		input.DestinationArn = aws.String(_sagemakerDestinationArn)
	}
	if len(_sagemakerDestinationType) > 0 {
		input.DestinationType = aws.String(_sagemakerDestinationType)
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSourceArn) > 0 {
		input.SourceArn = aws.String(_sagemakerSourceArn)
	}
	if len(_sagemakerSourceType) > 0 {
		input.SourceType = aws.String(_sagemakerSourceType)
	}

	if disablePaginator() {
		if resp, err := client.ListAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListAssociationsOutput
	p := sagemaker.NewListAssociationsPaginator(client, input)
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

// Request a list of jobs.
func sagemaker_ListAutoMLJobs(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListAutoMLJobsInput{}

	if len(_sagemakerCreationTimeAfter) > 0 {
		if err := assignInputField(input, "CreationTimeAfter", _sagemakerCreationTimeAfter); err != nil {
			log.Errorf("invalid --creation-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreationTimeBefore) > 0 {
		if err := assignInputField(input, "CreationTimeBefore", _sagemakerCreationTimeBefore); err != nil {
			log.Errorf("invalid --creation-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerLastModifiedTimeAfter) > 0 {
		if err := assignInputField(input, "LastModifiedTimeAfter", _sagemakerLastModifiedTimeAfter); err != nil {
			log.Errorf("invalid --last-modified-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerLastModifiedTimeBefore) > 0 {
		if err := assignInputField(input, "LastModifiedTimeBefore", _sagemakerLastModifiedTimeBefore); err != nil {
			log.Errorf("invalid --last-modified-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNameContains) > 0 {
		input.NameContains = aws.String(_sagemakerNameContains)
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}
	if len(_sagemakerStatusEquals) > 0 {
		if err := assignInputField(input, "StatusEquals", _sagemakerStatusEquals); err != nil {
			log.Errorf("invalid --status-equals: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListAutoMLJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListAutoMLJobsOutput
	p := sagemaker.NewListAutoMLJobsPaginator(client, input)
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

// List the candidates created for the job.
func sagemaker_ListCandidatesForAutoMLJob(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListCandidatesForAutoMLJobInput{
		// AutoMLJobName: *string, // Required
	}

	if len(_sagemakerAutoMLJobName) > 0 {
		input.AutoMLJobName = aws.String(_sagemakerAutoMLJobName)
	}
	if len(_sagemakerCandidateNameEquals) > 0 {
		input.CandidateNameEquals = aws.String(_sagemakerCandidateNameEquals)
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}
	if len(_sagemakerStatusEquals) > 0 {
		if err := assignInputField(input, "StatusEquals", _sagemakerStatusEquals); err != nil {
			log.Errorf("invalid --status-equals: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListCandidatesForAutoMLJob(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListCandidatesForAutoMLJobOutput
	p := sagemaker.NewListCandidatesForAutoMLJobPaginator(client, input)
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

// Retrieves a list of event summaries for a specified HyperPod cluster. The
// operation supports filtering, sorting, and pagination of results. This
// functionality is only supported when the NodeProvisioningMode is set to
// Continuous .
func sagemaker_ListClusterEvents(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListClusterEventsInput{
		// ClusterName: *string, // Required
	}

	if len(_sagemakerClusterName) > 0 {
		input.ClusterName = aws.String(_sagemakerClusterName)
	}
	if len(_sagemakerEventTimeAfter) > 0 {
		if err := assignInputField(input, "EventTimeAfter", _sagemakerEventTimeAfter); err != nil {
			log.Errorf("invalid --event-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerEventTimeBefore) > 0 {
		if err := assignInputField(input, "EventTimeBefore", _sagemakerEventTimeBefore); err != nil {
			log.Errorf("invalid --event-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerInstanceGroupName) > 0 {
		input.InstanceGroupName = aws.String(_sagemakerInstanceGroupName)
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerNodeId) > 0 {
		input.NodeId = aws.String(_sagemakerNodeId)
	}
	if len(_sagemakerResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _sagemakerResourceType); err != nil {
			log.Errorf("invalid --resource-type: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListClusterEvents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListClusterEventsOutput
	p := sagemaker.NewListClusterEventsPaginator(client, input)
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

// Retrieves the list of instances (also called nodes interchangeably) in a
// SageMaker HyperPod cluster.
func sagemaker_ListClusterNodes(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListClusterNodesInput{
		// ClusterName: *string, // Required
	}

	if len(_sagemakerClusterName) > 0 {
		input.ClusterName = aws.String(_sagemakerClusterName)
	}
	if len(_sagemakerCreationTimeAfter) > 0 {
		if err := assignInputField(input, "CreationTimeAfter", _sagemakerCreationTimeAfter); err != nil {
			log.Errorf("invalid --creation-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreationTimeBefore) > 0 {
		if err := assignInputField(input, "CreationTimeBefore", _sagemakerCreationTimeBefore); err != nil {
			log.Errorf("invalid --creation-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerIncludeNodeLogicalIds) > 0 {
		if err := assignInputField(input, "IncludeNodeLogicalIds", _sagemakerIncludeNodeLogicalIds); err != nil {
			log.Errorf("invalid --include-node-logical-ids: %s", err.Error())
			return
		}
	}
	if len(_sagemakerInstanceGroupNameContains) > 0 {
		input.InstanceGroupNameContains = aws.String(_sagemakerInstanceGroupNameContains)
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListClusterNodes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListClusterNodesOutput
	p := sagemaker.NewListClusterNodesPaginator(client, input)
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

// List the cluster policy configurations.
func sagemaker_ListClusterSchedulerConfigs(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListClusterSchedulerConfigsInput{}

	if len(_sagemakerClusterArn) > 0 {
		input.ClusterArn = aws.String(_sagemakerClusterArn)
	}
	if len(_sagemakerCreatedAfter) > 0 {
		if err := assignInputField(input, "CreatedAfter", _sagemakerCreatedAfter); err != nil {
			log.Errorf("invalid --created-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreatedBefore) > 0 {
		if err := assignInputField(input, "CreatedBefore", _sagemakerCreatedBefore); err != nil {
			log.Errorf("invalid --created-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNameContains) > 0 {
		input.NameContains = aws.String(_sagemakerNameContains)
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}
	if len(_sagemakerStatus) > 0 {
		if err := assignInputField(input, "Status", _sagemakerStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListClusterSchedulerConfigs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListClusterSchedulerConfigsOutput
	p := sagemaker.NewListClusterSchedulerConfigsPaginator(client, input)
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

// Retrieves the list of SageMaker HyperPod clusters.
func sagemaker_ListClusters(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListClustersInput{}

	if len(_sagemakerCreationTimeAfter) > 0 {
		if err := assignInputField(input, "CreationTimeAfter", _sagemakerCreationTimeAfter); err != nil {
			log.Errorf("invalid --creation-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreationTimeBefore) > 0 {
		if err := assignInputField(input, "CreationTimeBefore", _sagemakerCreationTimeBefore); err != nil {
			log.Errorf("invalid --creation-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNameContains) > 0 {
		input.NameContains = aws.String(_sagemakerNameContains)
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}
	if len(_sagemakerTrainingPlanArn) > 0 {
		input.TrainingPlanArn = aws.String(_sagemakerTrainingPlanArn)
	}

	if disablePaginator() {
		if resp, err := client.ListClusters(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListClustersOutput
	p := sagemaker.NewListClustersPaginator(client, input)
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

// Gets a list of the Git repositories in your account.
func sagemaker_ListCodeRepositories(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListCodeRepositoriesInput{}

	if len(_sagemakerCreationTimeAfter) > 0 {
		if err := assignInputField(input, "CreationTimeAfter", _sagemakerCreationTimeAfter); err != nil {
			log.Errorf("invalid --creation-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreationTimeBefore) > 0 {
		if err := assignInputField(input, "CreationTimeBefore", _sagemakerCreationTimeBefore); err != nil {
			log.Errorf("invalid --creation-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerLastModifiedTimeAfter) > 0 {
		if err := assignInputField(input, "LastModifiedTimeAfter", _sagemakerLastModifiedTimeAfter); err != nil {
			log.Errorf("invalid --last-modified-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerLastModifiedTimeBefore) > 0 {
		if err := assignInputField(input, "LastModifiedTimeBefore", _sagemakerLastModifiedTimeBefore); err != nil {
			log.Errorf("invalid --last-modified-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNameContains) > 0 {
		input.NameContains = aws.String(_sagemakerNameContains)
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListCodeRepositories(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListCodeRepositoriesOutput
	p := sagemaker.NewListCodeRepositoriesPaginator(client, input)
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

// Lists model compilation jobs that satisfy various filters.
// To create a model compilation job, use [CreateCompilationJob]. To get information about a particular
// model compilation job you have created, use [DescribeCompilationJob].
//
// [DescribeCompilationJob]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DescribeCompilationJob.html
// [CreateCompilationJob]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateCompilationJob.html
func sagemaker_ListCompilationJobs(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListCompilationJobsInput{}

	if len(_sagemakerCreationTimeAfter) > 0 {
		if err := assignInputField(input, "CreationTimeAfter", _sagemakerCreationTimeAfter); err != nil {
			log.Errorf("invalid --creation-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreationTimeBefore) > 0 {
		if err := assignInputField(input, "CreationTimeBefore", _sagemakerCreationTimeBefore); err != nil {
			log.Errorf("invalid --creation-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerLastModifiedTimeAfter) > 0 {
		if err := assignInputField(input, "LastModifiedTimeAfter", _sagemakerLastModifiedTimeAfter); err != nil {
			log.Errorf("invalid --last-modified-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerLastModifiedTimeBefore) > 0 {
		if err := assignInputField(input, "LastModifiedTimeBefore", _sagemakerLastModifiedTimeBefore); err != nil {
			log.Errorf("invalid --last-modified-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNameContains) > 0 {
		input.NameContains = aws.String(_sagemakerNameContains)
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}
	if len(_sagemakerStatusEquals) > 0 {
		if err := assignInputField(input, "StatusEquals", _sagemakerStatusEquals); err != nil {
			log.Errorf("invalid --status-equals: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListCompilationJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListCompilationJobsOutput
	p := sagemaker.NewListCompilationJobsPaginator(client, input)
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

// List the resource allocation definitions.
func sagemaker_ListComputeQuotas(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListComputeQuotasInput{}

	if len(_sagemakerClusterArn) > 0 {
		input.ClusterArn = aws.String(_sagemakerClusterArn)
	}
	if len(_sagemakerCreatedAfter) > 0 {
		if err := assignInputField(input, "CreatedAfter", _sagemakerCreatedAfter); err != nil {
			log.Errorf("invalid --created-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreatedBefore) > 0 {
		if err := assignInputField(input, "CreatedBefore", _sagemakerCreatedBefore); err != nil {
			log.Errorf("invalid --created-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNameContains) > 0 {
		input.NameContains = aws.String(_sagemakerNameContains)
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}
	if len(_sagemakerStatus) > 0 {
		if err := assignInputField(input, "Status", _sagemakerStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListComputeQuotas(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListComputeQuotasOutput
	p := sagemaker.NewListComputeQuotasPaginator(client, input)
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

// Lists the contexts in your account and their properties.
func sagemaker_ListContexts(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListContextsInput{}

	if len(_sagemakerContextType) > 0 {
		input.ContextType = aws.String(_sagemakerContextType)
	}
	if len(_sagemakerCreatedAfter) > 0 {
		if err := assignInputField(input, "CreatedAfter", _sagemakerCreatedAfter); err != nil {
			log.Errorf("invalid --created-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreatedBefore) > 0 {
		if err := assignInputField(input, "CreatedBefore", _sagemakerCreatedBefore); err != nil {
			log.Errorf("invalid --created-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSourceUri) > 0 {
		input.SourceUri = aws.String(_sagemakerSourceUri)
	}

	if disablePaginator() {
		if resp, err := client.ListContexts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListContextsOutput
	p := sagemaker.NewListContextsPaginator(client, input)
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

// Lists the data quality job definitions in your account.
func sagemaker_ListDataQualityJobDefinitions(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListDataQualityJobDefinitionsInput{}

	if len(_sagemakerCreationTimeAfter) > 0 {
		if err := assignInputField(input, "CreationTimeAfter", _sagemakerCreationTimeAfter); err != nil {
			log.Errorf("invalid --creation-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreationTimeBefore) > 0 {
		if err := assignInputField(input, "CreationTimeBefore", _sagemakerCreationTimeBefore); err != nil {
			log.Errorf("invalid --creation-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerEndpointName) > 0 {
		input.EndpointName = aws.String(_sagemakerEndpointName)
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNameContains) > 0 {
		input.NameContains = aws.String(_sagemakerNameContains)
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListDataQualityJobDefinitions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListDataQualityJobDefinitionsOutput
	p := sagemaker.NewListDataQualityJobDefinitionsPaginator(client, input)
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

// Returns a list of devices in the fleet.
func sagemaker_ListDeviceFleets(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListDeviceFleetsInput{}

	if len(_sagemakerCreationTimeAfter) > 0 {
		if err := assignInputField(input, "CreationTimeAfter", _sagemakerCreationTimeAfter); err != nil {
			log.Errorf("invalid --creation-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreationTimeBefore) > 0 {
		if err := assignInputField(input, "CreationTimeBefore", _sagemakerCreationTimeBefore); err != nil {
			log.Errorf("invalid --creation-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerLastModifiedTimeAfter) > 0 {
		if err := assignInputField(input, "LastModifiedTimeAfter", _sagemakerLastModifiedTimeAfter); err != nil {
			log.Errorf("invalid --last-modified-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerLastModifiedTimeBefore) > 0 {
		if err := assignInputField(input, "LastModifiedTimeBefore", _sagemakerLastModifiedTimeBefore); err != nil {
			log.Errorf("invalid --last-modified-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNameContains) > 0 {
		input.NameContains = aws.String(_sagemakerNameContains)
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListDeviceFleets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListDeviceFleetsOutput
	p := sagemaker.NewListDeviceFleetsPaginator(client, input)
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

// A list of devices.
func sagemaker_ListDevices(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListDevicesInput{}

	if len(_sagemakerDeviceFleetName) > 0 {
		input.DeviceFleetName = aws.String(_sagemakerDeviceFleetName)
	}
	if len(_sagemakerLatestHeartbeatAfter) > 0 {
		if err := assignInputField(input, "LatestHeartbeatAfter", _sagemakerLatestHeartbeatAfter); err != nil {
			log.Errorf("invalid --latest-heartbeat-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerModelName) > 0 {
		input.ModelName = aws.String(_sagemakerModelName)
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDevices(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListDevicesOutput
	p := sagemaker.NewListDevicesPaginator(client, input)
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

// Lists the domains.
func sagemaker_ListDomains(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListDomainsInput{}

	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDomains(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListDomainsOutput
	p := sagemaker.NewListDomainsPaginator(client, input)
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

// Lists all edge deployment plans.
func sagemaker_ListEdgeDeploymentPlans(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListEdgeDeploymentPlansInput{}

	if len(_sagemakerCreationTimeAfter) > 0 {
		if err := assignInputField(input, "CreationTimeAfter", _sagemakerCreationTimeAfter); err != nil {
			log.Errorf("invalid --creation-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreationTimeBefore) > 0 {
		if err := assignInputField(input, "CreationTimeBefore", _sagemakerCreationTimeBefore); err != nil {
			log.Errorf("invalid --creation-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerDeviceFleetNameContains) > 0 {
		input.DeviceFleetNameContains = aws.String(_sagemakerDeviceFleetNameContains)
	}
	if len(_sagemakerLastModifiedTimeAfter) > 0 {
		if err := assignInputField(input, "LastModifiedTimeAfter", _sagemakerLastModifiedTimeAfter); err != nil {
			log.Errorf("invalid --last-modified-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerLastModifiedTimeBefore) > 0 {
		if err := assignInputField(input, "LastModifiedTimeBefore", _sagemakerLastModifiedTimeBefore); err != nil {
			log.Errorf("invalid --last-modified-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNameContains) > 0 {
		input.NameContains = aws.String(_sagemakerNameContains)
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListEdgeDeploymentPlans(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListEdgeDeploymentPlansOutput
	p := sagemaker.NewListEdgeDeploymentPlansPaginator(client, input)
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

// Returns a list of edge packaging jobs.
func sagemaker_ListEdgePackagingJobs(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListEdgePackagingJobsInput{}

	if len(_sagemakerCreationTimeAfter) > 0 {
		if err := assignInputField(input, "CreationTimeAfter", _sagemakerCreationTimeAfter); err != nil {
			log.Errorf("invalid --creation-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreationTimeBefore) > 0 {
		if err := assignInputField(input, "CreationTimeBefore", _sagemakerCreationTimeBefore); err != nil {
			log.Errorf("invalid --creation-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerLastModifiedTimeAfter) > 0 {
		if err := assignInputField(input, "LastModifiedTimeAfter", _sagemakerLastModifiedTimeAfter); err != nil {
			log.Errorf("invalid --last-modified-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerLastModifiedTimeBefore) > 0 {
		if err := assignInputField(input, "LastModifiedTimeBefore", _sagemakerLastModifiedTimeBefore); err != nil {
			log.Errorf("invalid --last-modified-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerModelNameContains) > 0 {
		input.ModelNameContains = aws.String(_sagemakerModelNameContains)
	}
	if len(_sagemakerNameContains) > 0 {
		input.NameContains = aws.String(_sagemakerNameContains)
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}
	if len(_sagemakerStatusEquals) > 0 {
		if err := assignInputField(input, "StatusEquals", _sagemakerStatusEquals); err != nil {
			log.Errorf("invalid --status-equals: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListEdgePackagingJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListEdgePackagingJobsOutput
	p := sagemaker.NewListEdgePackagingJobsPaginator(client, input)
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

// Lists endpoint configurations.
func sagemaker_ListEndpointConfigs(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListEndpointConfigsInput{}

	if len(_sagemakerCreationTimeAfter) > 0 {
		if err := assignInputField(input, "CreationTimeAfter", _sagemakerCreationTimeAfter); err != nil {
			log.Errorf("invalid --creation-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreationTimeBefore) > 0 {
		if err := assignInputField(input, "CreationTimeBefore", _sagemakerCreationTimeBefore); err != nil {
			log.Errorf("invalid --creation-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNameContains) > 0 {
		input.NameContains = aws.String(_sagemakerNameContains)
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListEndpointConfigs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListEndpointConfigsOutput
	p := sagemaker.NewListEndpointConfigsPaginator(client, input)
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

// Lists endpoints.
func sagemaker_ListEndpoints(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListEndpointsInput{}

	if len(_sagemakerCreationTimeAfter) > 0 {
		if err := assignInputField(input, "CreationTimeAfter", _sagemakerCreationTimeAfter); err != nil {
			log.Errorf("invalid --creation-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreationTimeBefore) > 0 {
		if err := assignInputField(input, "CreationTimeBefore", _sagemakerCreationTimeBefore); err != nil {
			log.Errorf("invalid --creation-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerLastModifiedTimeAfter) > 0 {
		if err := assignInputField(input, "LastModifiedTimeAfter", _sagemakerLastModifiedTimeAfter); err != nil {
			log.Errorf("invalid --last-modified-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerLastModifiedTimeBefore) > 0 {
		if err := assignInputField(input, "LastModifiedTimeBefore", _sagemakerLastModifiedTimeBefore); err != nil {
			log.Errorf("invalid --last-modified-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNameContains) > 0 {
		input.NameContains = aws.String(_sagemakerNameContains)
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}
	if len(_sagemakerStatusEquals) > 0 {
		if err := assignInputField(input, "StatusEquals", _sagemakerStatusEquals); err != nil {
			log.Errorf("invalid --status-equals: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListEndpoints(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListEndpointsOutput
	p := sagemaker.NewListEndpointsPaginator(client, input)
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

// Lists all the experiments in your account. The list can be filtered to show
// only experiments that were created in a specific time range. The list can be
// sorted by experiment name or creation time.
func sagemaker_ListExperiments(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListExperimentsInput{}

	if len(_sagemakerCreatedAfter) > 0 {
		if err := assignInputField(input, "CreatedAfter", _sagemakerCreatedAfter); err != nil {
			log.Errorf("invalid --created-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreatedBefore) > 0 {
		if err := assignInputField(input, "CreatedBefore", _sagemakerCreatedBefore); err != nil {
			log.Errorf("invalid --created-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListExperiments(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListExperimentsOutput
	p := sagemaker.NewListExperimentsPaginator(client, input)
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

// List FeatureGroup s based on given filter and order.
func sagemaker_ListFeatureGroups(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListFeatureGroupsInput{}

	if len(_sagemakerCreationTimeAfter) > 0 {
		if err := assignInputField(input, "CreationTimeAfter", _sagemakerCreationTimeAfter); err != nil {
			log.Errorf("invalid --creation-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreationTimeBefore) > 0 {
		if err := assignInputField(input, "CreationTimeBefore", _sagemakerCreationTimeBefore); err != nil {
			log.Errorf("invalid --creation-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerFeatureGroupStatusEquals) > 0 {
		if err := assignInputField(input, "FeatureGroupStatusEquals", _sagemakerFeatureGroupStatusEquals); err != nil {
			log.Errorf("invalid --feature-group-status-equals: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNameContains) > 0 {
		input.NameContains = aws.String(_sagemakerNameContains)
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerOfflineStoreStatusEquals) > 0 {
		if err := assignInputField(input, "OfflineStoreStatusEquals", _sagemakerOfflineStoreStatusEquals); err != nil {
			log.Errorf("invalid --offline-store-status-equals: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListFeatureGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListFeatureGroupsOutput
	p := sagemaker.NewListFeatureGroupsPaginator(client, input)
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

// Returns information about the flow definitions in your account.
func sagemaker_ListFlowDefinitions(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListFlowDefinitionsInput{}

	if len(_sagemakerCreationTimeAfter) > 0 {
		if err := assignInputField(input, "CreationTimeAfter", _sagemakerCreationTimeAfter); err != nil {
			log.Errorf("invalid --creation-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreationTimeBefore) > 0 {
		if err := assignInputField(input, "CreationTimeBefore", _sagemakerCreationTimeBefore); err != nil {
			log.Errorf("invalid --creation-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListFlowDefinitions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListFlowDefinitionsOutput
	p := sagemaker.NewListFlowDefinitionsPaginator(client, input)
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

// List hub content versions.
func sagemaker_ListHubContentVersions(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListHubContentVersionsInput{
		// HubContentName: *string, // Required
		// HubContentType: types.HubContentType, // Required
		// HubName: *string, // Required
	}

	if len(_sagemakerHubContentName) > 0 {
		input.HubContentName = aws.String(_sagemakerHubContentName)
	}
	if len(_sagemakerHubContentType) > 0 {
		if err := assignInputField(input, "HubContentType", _sagemakerHubContentType); err != nil {
			log.Errorf("invalid --hub-content-type: %s", err.Error())
			return
		}
	}
	if len(_sagemakerHubName) > 0 {
		input.HubName = aws.String(_sagemakerHubName)
	}
	if len(_sagemakerCreationTimeAfter) > 0 {
		if err := assignInputField(input, "CreationTimeAfter", _sagemakerCreationTimeAfter); err != nil {
			log.Errorf("invalid --creation-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreationTimeBefore) > 0 {
		if err := assignInputField(input, "CreationTimeBefore", _sagemakerCreationTimeBefore); err != nil {
			log.Errorf("invalid --creation-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMaxSchemaVersion) > 0 {
		input.MaxSchemaVersion = aws.String(_sagemakerMaxSchemaVersion)
	}
	if len(_sagemakerMinVersion) > 0 {
		input.MinVersion = aws.String(_sagemakerMinVersion)
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListHubContentVersions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// List the contents of a hub.
func sagemaker_ListHubContents(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListHubContentsInput{
		// HubContentType: types.HubContentType, // Required
		// HubName: *string, // Required
	}

	if len(_sagemakerHubContentType) > 0 {
		if err := assignInputField(input, "HubContentType", _sagemakerHubContentType); err != nil {
			log.Errorf("invalid --hub-content-type: %s", err.Error())
			return
		}
	}
	if len(_sagemakerHubName) > 0 {
		input.HubName = aws.String(_sagemakerHubName)
	}
	if len(_sagemakerCreationTimeAfter) > 0 {
		if err := assignInputField(input, "CreationTimeAfter", _sagemakerCreationTimeAfter); err != nil {
			log.Errorf("invalid --creation-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreationTimeBefore) > 0 {
		if err := assignInputField(input, "CreationTimeBefore", _sagemakerCreationTimeBefore); err != nil {
			log.Errorf("invalid --creation-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMaxSchemaVersion) > 0 {
		input.MaxSchemaVersion = aws.String(_sagemakerMaxSchemaVersion)
	}
	if len(_sagemakerNameContains) > 0 {
		input.NameContains = aws.String(_sagemakerNameContains)
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListHubContents(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// List all existing hubs.
func sagemaker_ListHubs(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListHubsInput{}

	if len(_sagemakerCreationTimeAfter) > 0 {
		if err := assignInputField(input, "CreationTimeAfter", _sagemakerCreationTimeAfter); err != nil {
			log.Errorf("invalid --creation-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreationTimeBefore) > 0 {
		if err := assignInputField(input, "CreationTimeBefore", _sagemakerCreationTimeBefore); err != nil {
			log.Errorf("invalid --creation-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerLastModifiedTimeAfter) > 0 {
		if err := assignInputField(input, "LastModifiedTimeAfter", _sagemakerLastModifiedTimeAfter); err != nil {
			log.Errorf("invalid --last-modified-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerLastModifiedTimeBefore) > 0 {
		if err := assignInputField(input, "LastModifiedTimeBefore", _sagemakerLastModifiedTimeBefore); err != nil {
			log.Errorf("invalid --last-modified-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNameContains) > 0 {
		input.NameContains = aws.String(_sagemakerNameContains)
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListHubs(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the human task user interfaces in your account.
func sagemaker_ListHumanTaskUis(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListHumanTaskUisInput{}

	if len(_sagemakerCreationTimeAfter) > 0 {
		if err := assignInputField(input, "CreationTimeAfter", _sagemakerCreationTimeAfter); err != nil {
			log.Errorf("invalid --creation-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreationTimeBefore) > 0 {
		if err := assignInputField(input, "CreationTimeBefore", _sagemakerCreationTimeBefore); err != nil {
			log.Errorf("invalid --creation-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListHumanTaskUis(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListHumanTaskUisOutput
	p := sagemaker.NewListHumanTaskUisPaginator(client, input)
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

// Gets a list of [HyperParameterTuningJobSummary] objects that describe the hyperparameter tuning jobs launched
// in your account.
//
// [HyperParameterTuningJobSummary]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_HyperParameterTuningJobSummary.html
func sagemaker_ListHyperParameterTuningJobs(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListHyperParameterTuningJobsInput{}

	if len(_sagemakerCreationTimeAfter) > 0 {
		if err := assignInputField(input, "CreationTimeAfter", _sagemakerCreationTimeAfter); err != nil {
			log.Errorf("invalid --creation-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreationTimeBefore) > 0 {
		if err := assignInputField(input, "CreationTimeBefore", _sagemakerCreationTimeBefore); err != nil {
			log.Errorf("invalid --creation-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerLastModifiedTimeAfter) > 0 {
		if err := assignInputField(input, "LastModifiedTimeAfter", _sagemakerLastModifiedTimeAfter); err != nil {
			log.Errorf("invalid --last-modified-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerLastModifiedTimeBefore) > 0 {
		if err := assignInputField(input, "LastModifiedTimeBefore", _sagemakerLastModifiedTimeBefore); err != nil {
			log.Errorf("invalid --last-modified-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNameContains) > 0 {
		input.NameContains = aws.String(_sagemakerNameContains)
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}
	if len(_sagemakerStatusEquals) > 0 {
		if err := assignInputField(input, "StatusEquals", _sagemakerStatusEquals); err != nil {
			log.Errorf("invalid --status-equals: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListHyperParameterTuningJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListHyperParameterTuningJobsOutput
	p := sagemaker.NewListHyperParameterTuningJobsPaginator(client, input)
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

// Lists the versions of a specified image and their properties. The list can be
// filtered by creation time or modified time.
func sagemaker_ListImageVersions(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListImageVersionsInput{
		// ImageName: *string, // Required
	}

	if len(_sagemakerImageName) > 0 {
		input.ImageName = aws.String(_sagemakerImageName)
	}
	if len(_sagemakerCreationTimeAfter) > 0 {
		if err := assignInputField(input, "CreationTimeAfter", _sagemakerCreationTimeAfter); err != nil {
			log.Errorf("invalid --creation-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreationTimeBefore) > 0 {
		if err := assignInputField(input, "CreationTimeBefore", _sagemakerCreationTimeBefore); err != nil {
			log.Errorf("invalid --creation-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerLastModifiedTimeAfter) > 0 {
		if err := assignInputField(input, "LastModifiedTimeAfter", _sagemakerLastModifiedTimeAfter); err != nil {
			log.Errorf("invalid --last-modified-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerLastModifiedTimeBefore) > 0 {
		if err := assignInputField(input, "LastModifiedTimeBefore", _sagemakerLastModifiedTimeBefore); err != nil {
			log.Errorf("invalid --last-modified-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListImageVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListImageVersionsOutput
	p := sagemaker.NewListImageVersionsPaginator(client, input)
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

// Lists the images in your account and their properties. The list can be filtered
// by creation time or modified time, and whether the image name contains a
// specified string.
func sagemaker_ListImages(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListImagesInput{}

	if len(_sagemakerCreationTimeAfter) > 0 {
		if err := assignInputField(input, "CreationTimeAfter", _sagemakerCreationTimeAfter); err != nil {
			log.Errorf("invalid --creation-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreationTimeBefore) > 0 {
		if err := assignInputField(input, "CreationTimeBefore", _sagemakerCreationTimeBefore); err != nil {
			log.Errorf("invalid --creation-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerLastModifiedTimeAfter) > 0 {
		if err := assignInputField(input, "LastModifiedTimeAfter", _sagemakerLastModifiedTimeAfter); err != nil {
			log.Errorf("invalid --last-modified-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerLastModifiedTimeBefore) > 0 {
		if err := assignInputField(input, "LastModifiedTimeBefore", _sagemakerLastModifiedTimeBefore); err != nil {
			log.Errorf("invalid --last-modified-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNameContains) > 0 {
		input.NameContains = aws.String(_sagemakerNameContains)
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListImages(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListImagesOutput
	p := sagemaker.NewListImagesPaginator(client, input)
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

// Lists the inference components in your account and their properties.
func sagemaker_ListInferenceComponents(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListInferenceComponentsInput{}

	if len(_sagemakerCreationTimeAfter) > 0 {
		if err := assignInputField(input, "CreationTimeAfter", _sagemakerCreationTimeAfter); err != nil {
			log.Errorf("invalid --creation-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreationTimeBefore) > 0 {
		if err := assignInputField(input, "CreationTimeBefore", _sagemakerCreationTimeBefore); err != nil {
			log.Errorf("invalid --creation-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerEndpointNameEquals) > 0 {
		input.EndpointNameEquals = aws.String(_sagemakerEndpointNameEquals)
	}
	if len(_sagemakerLastModifiedTimeAfter) > 0 {
		if err := assignInputField(input, "LastModifiedTimeAfter", _sagemakerLastModifiedTimeAfter); err != nil {
			log.Errorf("invalid --last-modified-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerLastModifiedTimeBefore) > 0 {
		if err := assignInputField(input, "LastModifiedTimeBefore", _sagemakerLastModifiedTimeBefore); err != nil {
			log.Errorf("invalid --last-modified-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNameContains) > 0 {
		input.NameContains = aws.String(_sagemakerNameContains)
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}
	if len(_sagemakerStatusEquals) > 0 {
		if err := assignInputField(input, "StatusEquals", _sagemakerStatusEquals); err != nil {
			log.Errorf("invalid --status-equals: %s", err.Error())
			return
		}
	}
	if len(_sagemakerVariantNameEquals) > 0 {
		input.VariantNameEquals = aws.String(_sagemakerVariantNameEquals)
	}

	if disablePaginator() {
		if resp, err := client.ListInferenceComponents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListInferenceComponentsOutput
	p := sagemaker.NewListInferenceComponentsPaginator(client, input)
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

// Returns the list of all inference experiments.
func sagemaker_ListInferenceExperiments(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListInferenceExperimentsInput{}

	if len(_sagemakerCreationTimeAfter) > 0 {
		if err := assignInputField(input, "CreationTimeAfter", _sagemakerCreationTimeAfter); err != nil {
			log.Errorf("invalid --creation-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreationTimeBefore) > 0 {
		if err := assignInputField(input, "CreationTimeBefore", _sagemakerCreationTimeBefore); err != nil {
			log.Errorf("invalid --creation-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerLastModifiedTimeAfter) > 0 {
		if err := assignInputField(input, "LastModifiedTimeAfter", _sagemakerLastModifiedTimeAfter); err != nil {
			log.Errorf("invalid --last-modified-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerLastModifiedTimeBefore) > 0 {
		if err := assignInputField(input, "LastModifiedTimeBefore", _sagemakerLastModifiedTimeBefore); err != nil {
			log.Errorf("invalid --last-modified-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNameContains) > 0 {
		input.NameContains = aws.String(_sagemakerNameContains)
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}
	if len(_sagemakerStatusEquals) > 0 {
		if err := assignInputField(input, "StatusEquals", _sagemakerStatusEquals); err != nil {
			log.Errorf("invalid --status-equals: %s", err.Error())
			return
		}
	}
	if len(_sagemakerType) > 0 {
		if err := assignInputField(input, "Type", _sagemakerType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListInferenceExperiments(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListInferenceExperimentsOutput
	p := sagemaker.NewListInferenceExperimentsPaginator(client, input)
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

// Returns a list of the subtasks for an Inference Recommender job.
// The supported subtasks are benchmarks, which evaluate the performance of your
// model on different instance types.
func sagemaker_ListInferenceRecommendationsJobSteps(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListInferenceRecommendationsJobStepsInput{
		// JobName: *string, // Required
	}

	if len(_sagemakerJobName) > 0 {
		input.JobName = aws.String(_sagemakerJobName)
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerStatus) > 0 {
		if err := assignInputField(input, "Status", _sagemakerStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_sagemakerStepType) > 0 {
		if err := assignInputField(input, "StepType", _sagemakerStepType); err != nil {
			log.Errorf("invalid --step-type: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListInferenceRecommendationsJobSteps(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListInferenceRecommendationsJobStepsOutput
	p := sagemaker.NewListInferenceRecommendationsJobStepsPaginator(client, input)
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

// Lists recommendation jobs that satisfy various filters.
func sagemaker_ListInferenceRecommendationsJobs(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListInferenceRecommendationsJobsInput{}

	if len(_sagemakerCreationTimeAfter) > 0 {
		if err := assignInputField(input, "CreationTimeAfter", _sagemakerCreationTimeAfter); err != nil {
			log.Errorf("invalid --creation-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreationTimeBefore) > 0 {
		if err := assignInputField(input, "CreationTimeBefore", _sagemakerCreationTimeBefore); err != nil {
			log.Errorf("invalid --creation-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerLastModifiedTimeAfter) > 0 {
		if err := assignInputField(input, "LastModifiedTimeAfter", _sagemakerLastModifiedTimeAfter); err != nil {
			log.Errorf("invalid --last-modified-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerLastModifiedTimeBefore) > 0 {
		if err := assignInputField(input, "LastModifiedTimeBefore", _sagemakerLastModifiedTimeBefore); err != nil {
			log.Errorf("invalid --last-modified-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerModelNameEquals) > 0 {
		input.ModelNameEquals = aws.String(_sagemakerModelNameEquals)
	}
	if len(_sagemakerModelPackageVersionArnEquals) > 0 {
		input.ModelPackageVersionArnEquals = aws.String(_sagemakerModelPackageVersionArnEquals)
	}
	if len(_sagemakerNameContains) > 0 {
		input.NameContains = aws.String(_sagemakerNameContains)
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}
	if len(_sagemakerStatusEquals) > 0 {
		if err := assignInputField(input, "StatusEquals", _sagemakerStatusEquals); err != nil {
			log.Errorf("invalid --status-equals: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListInferenceRecommendationsJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListInferenceRecommendationsJobsOutput
	p := sagemaker.NewListInferenceRecommendationsJobsPaginator(client, input)
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

// Gets a list of labeling jobs.
func sagemaker_ListLabelingJobs(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListLabelingJobsInput{}

	if len(_sagemakerCreationTimeAfter) > 0 {
		if err := assignInputField(input, "CreationTimeAfter", _sagemakerCreationTimeAfter); err != nil {
			log.Errorf("invalid --creation-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreationTimeBefore) > 0 {
		if err := assignInputField(input, "CreationTimeBefore", _sagemakerCreationTimeBefore); err != nil {
			log.Errorf("invalid --creation-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerLastModifiedTimeAfter) > 0 {
		if err := assignInputField(input, "LastModifiedTimeAfter", _sagemakerLastModifiedTimeAfter); err != nil {
			log.Errorf("invalid --last-modified-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerLastModifiedTimeBefore) > 0 {
		if err := assignInputField(input, "LastModifiedTimeBefore", _sagemakerLastModifiedTimeBefore); err != nil {
			log.Errorf("invalid --last-modified-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNameContains) > 0 {
		input.NameContains = aws.String(_sagemakerNameContains)
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}
	if len(_sagemakerStatusEquals) > 0 {
		if err := assignInputField(input, "StatusEquals", _sagemakerStatusEquals); err != nil {
			log.Errorf("invalid --status-equals: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListLabelingJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListLabelingJobsOutput
	p := sagemaker.NewListLabelingJobsPaginator(client, input)
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

// Gets a list of labeling jobs assigned to a specified work team.
func sagemaker_ListLabelingJobsForWorkteam(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListLabelingJobsForWorkteamInput{
		// WorkteamArn: *string, // Required
	}

	if len(_sagemakerWorkteamArn) > 0 {
		input.WorkteamArn = aws.String(_sagemakerWorkteamArn)
	}
	if len(_sagemakerCreationTimeAfter) > 0 {
		if err := assignInputField(input, "CreationTimeAfter", _sagemakerCreationTimeAfter); err != nil {
			log.Errorf("invalid --creation-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreationTimeBefore) > 0 {
		if err := assignInputField(input, "CreationTimeBefore", _sagemakerCreationTimeBefore); err != nil {
			log.Errorf("invalid --creation-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerJobReferenceCodeContains) > 0 {
		input.JobReferenceCodeContains = aws.String(_sagemakerJobReferenceCodeContains)
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListLabelingJobsForWorkteam(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListLabelingJobsForWorkteamOutput
	p := sagemaker.NewListLabelingJobsForWorkteamPaginator(client, input)
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

// A list of lineage groups shared with your Amazon Web Services account. For more
// information, see [Cross-Account Lineage Tracking]in the Amazon SageMaker Developer Guide.
//
// [Cross-Account Lineage Tracking]: https://docs.aws.amazon.com/sagemaker/latest/dg/xaccount-lineage-tracking.html
func sagemaker_ListLineageGroups(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListLineageGroupsInput{}

	if len(_sagemakerCreatedAfter) > 0 {
		if err := assignInputField(input, "CreatedAfter", _sagemakerCreatedAfter); err != nil {
			log.Errorf("invalid --created-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreatedBefore) > 0 {
		if err := assignInputField(input, "CreatedBefore", _sagemakerCreatedBefore); err != nil {
			log.Errorf("invalid --created-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListLineageGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListLineageGroupsOutput
	p := sagemaker.NewListLineageGroupsPaginator(client, input)
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

// Lists all MLflow Apps
func sagemaker_ListMlflowApps(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListMlflowAppsInput{}

	if len(_sagemakerAccountDefaultStatus) > 0 {
		if err := assignInputField(input, "AccountDefaultStatus", _sagemakerAccountDefaultStatus); err != nil {
			log.Errorf("invalid --account-default-status: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreatedAfter) > 0 {
		if err := assignInputField(input, "CreatedAfter", _sagemakerCreatedAfter); err != nil {
			log.Errorf("invalid --created-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreatedBefore) > 0 {
		if err := assignInputField(input, "CreatedBefore", _sagemakerCreatedBefore); err != nil {
			log.Errorf("invalid --created-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerDefaultForDomainId) > 0 {
		input.DefaultForDomainId = aws.String(_sagemakerDefaultForDomainId)
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMlflowVersion) > 0 {
		input.MlflowVersion = aws.String(_sagemakerMlflowVersion)
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}
	if len(_sagemakerStatus) > 0 {
		if err := assignInputField(input, "Status", _sagemakerStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListMlflowApps(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListMlflowAppsOutput
	p := sagemaker.NewListMlflowAppsPaginator(client, input)
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

// Lists all MLflow Tracking Servers.
func sagemaker_ListMlflowTrackingServers(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListMlflowTrackingServersInput{}

	if len(_sagemakerCreatedAfter) > 0 {
		if err := assignInputField(input, "CreatedAfter", _sagemakerCreatedAfter); err != nil {
			log.Errorf("invalid --created-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreatedBefore) > 0 {
		if err := assignInputField(input, "CreatedBefore", _sagemakerCreatedBefore); err != nil {
			log.Errorf("invalid --created-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMlflowVersion) > 0 {
		input.MlflowVersion = aws.String(_sagemakerMlflowVersion)
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}
	if len(_sagemakerTrackingServerStatus) > 0 {
		if err := assignInputField(input, "TrackingServerStatus", _sagemakerTrackingServerStatus); err != nil {
			log.Errorf("invalid --tracking-server-status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListMlflowTrackingServers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListMlflowTrackingServersOutput
	p := sagemaker.NewListMlflowTrackingServersPaginator(client, input)
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

// Lists model bias jobs definitions that satisfy various filters.
func sagemaker_ListModelBiasJobDefinitions(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListModelBiasJobDefinitionsInput{}

	if len(_sagemakerCreationTimeAfter) > 0 {
		if err := assignInputField(input, "CreationTimeAfter", _sagemakerCreationTimeAfter); err != nil {
			log.Errorf("invalid --creation-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreationTimeBefore) > 0 {
		if err := assignInputField(input, "CreationTimeBefore", _sagemakerCreationTimeBefore); err != nil {
			log.Errorf("invalid --creation-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerEndpointName) > 0 {
		input.EndpointName = aws.String(_sagemakerEndpointName)
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNameContains) > 0 {
		input.NameContains = aws.String(_sagemakerNameContains)
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListModelBiasJobDefinitions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListModelBiasJobDefinitionsOutput
	p := sagemaker.NewListModelBiasJobDefinitionsPaginator(client, input)
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

// List the export jobs for the Amazon SageMaker Model Card.
func sagemaker_ListModelCardExportJobs(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListModelCardExportJobsInput{
		// ModelCardName: *string, // Required
	}

	if len(_sagemakerModelCardName) > 0 {
		input.ModelCardName = aws.String(_sagemakerModelCardName)
	}
	if len(_sagemakerCreationTimeAfter) > 0 {
		if err := assignInputField(input, "CreationTimeAfter", _sagemakerCreationTimeAfter); err != nil {
			log.Errorf("invalid --creation-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreationTimeBefore) > 0 {
		if err := assignInputField(input, "CreationTimeBefore", _sagemakerCreationTimeBefore); err != nil {
			log.Errorf("invalid --creation-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerModelCardExportJobNameContains) > 0 {
		input.ModelCardExportJobNameContains = aws.String(_sagemakerModelCardExportJobNameContains)
	}
	if len(_sagemakerModelCardVersion) > 0 {
		if err := assignInputField(input, "ModelCardVersion", _sagemakerModelCardVersion); err != nil {
			log.Errorf("invalid --model-card-version: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}
	if len(_sagemakerStatusEquals) > 0 {
		if err := assignInputField(input, "StatusEquals", _sagemakerStatusEquals); err != nil {
			log.Errorf("invalid --status-equals: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListModelCardExportJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListModelCardExportJobsOutput
	p := sagemaker.NewListModelCardExportJobsPaginator(client, input)
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

// List existing versions of an Amazon SageMaker Model Card.
func sagemaker_ListModelCardVersions(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListModelCardVersionsInput{
		// ModelCardName: *string, // Required
	}

	if len(_sagemakerModelCardName) > 0 {
		input.ModelCardName = aws.String(_sagemakerModelCardName)
	}
	if len(_sagemakerCreationTimeAfter) > 0 {
		if err := assignInputField(input, "CreationTimeAfter", _sagemakerCreationTimeAfter); err != nil {
			log.Errorf("invalid --creation-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreationTimeBefore) > 0 {
		if err := assignInputField(input, "CreationTimeBefore", _sagemakerCreationTimeBefore); err != nil {
			log.Errorf("invalid --creation-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerModelCardStatus) > 0 {
		if err := assignInputField(input, "ModelCardStatus", _sagemakerModelCardStatus); err != nil {
			log.Errorf("invalid --model-card-status: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListModelCardVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListModelCardVersionsOutput
	p := sagemaker.NewListModelCardVersionsPaginator(client, input)
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

// List existing model cards.
func sagemaker_ListModelCards(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListModelCardsInput{}

	if len(_sagemakerCreationTimeAfter) > 0 {
		if err := assignInputField(input, "CreationTimeAfter", _sagemakerCreationTimeAfter); err != nil {
			log.Errorf("invalid --creation-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreationTimeBefore) > 0 {
		if err := assignInputField(input, "CreationTimeBefore", _sagemakerCreationTimeBefore); err != nil {
			log.Errorf("invalid --creation-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerModelCardStatus) > 0 {
		if err := assignInputField(input, "ModelCardStatus", _sagemakerModelCardStatus); err != nil {
			log.Errorf("invalid --model-card-status: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNameContains) > 0 {
		input.NameContains = aws.String(_sagemakerNameContains)
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListModelCards(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListModelCardsOutput
	p := sagemaker.NewListModelCardsPaginator(client, input)
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

// Lists model explainability job definitions that satisfy various filters.
func sagemaker_ListModelExplainabilityJobDefinitions(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListModelExplainabilityJobDefinitionsInput{}

	if len(_sagemakerCreationTimeAfter) > 0 {
		if err := assignInputField(input, "CreationTimeAfter", _sagemakerCreationTimeAfter); err != nil {
			log.Errorf("invalid --creation-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreationTimeBefore) > 0 {
		if err := assignInputField(input, "CreationTimeBefore", _sagemakerCreationTimeBefore); err != nil {
			log.Errorf("invalid --creation-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerEndpointName) > 0 {
		input.EndpointName = aws.String(_sagemakerEndpointName)
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNameContains) > 0 {
		input.NameContains = aws.String(_sagemakerNameContains)
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListModelExplainabilityJobDefinitions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListModelExplainabilityJobDefinitionsOutput
	p := sagemaker.NewListModelExplainabilityJobDefinitionsPaginator(client, input)
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

// Lists the domain, framework, task, and model name of standard machine learning
// models found in common model zoos.
func sagemaker_ListModelMetadata(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListModelMetadataInput{}

	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSearchExpression) > 0 {
		if err := assignInputField(input, "SearchExpression", _sagemakerSearchExpression); err != nil {
			log.Errorf("invalid --search-expression: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListModelMetadata(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListModelMetadataOutput
	p := sagemaker.NewListModelMetadataPaginator(client, input)
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

// Gets a list of the model groups in your Amazon Web Services account.
func sagemaker_ListModelPackageGroups(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListModelPackageGroupsInput{}

	if len(_sagemakerCreationTimeAfter) > 0 {
		if err := assignInputField(input, "CreationTimeAfter", _sagemakerCreationTimeAfter); err != nil {
			log.Errorf("invalid --creation-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreationTimeBefore) > 0 {
		if err := assignInputField(input, "CreationTimeBefore", _sagemakerCreationTimeBefore); err != nil {
			log.Errorf("invalid --creation-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCrossAccountFilterOption) > 0 {
		if err := assignInputField(input, "CrossAccountFilterOption", _sagemakerCrossAccountFilterOption); err != nil {
			log.Errorf("invalid --cross-account-filter-option: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNameContains) > 0 {
		input.NameContains = aws.String(_sagemakerNameContains)
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListModelPackageGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListModelPackageGroupsOutput
	p := sagemaker.NewListModelPackageGroupsPaginator(client, input)
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

// Lists the model packages that have been created.
func sagemaker_ListModelPackages(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListModelPackagesInput{}

	if len(_sagemakerCreationTimeAfter) > 0 {
		if err := assignInputField(input, "CreationTimeAfter", _sagemakerCreationTimeAfter); err != nil {
			log.Errorf("invalid --creation-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreationTimeBefore) > 0 {
		if err := assignInputField(input, "CreationTimeBefore", _sagemakerCreationTimeBefore); err != nil {
			log.Errorf("invalid --creation-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerModelApprovalStatus) > 0 {
		if err := assignInputField(input, "ModelApprovalStatus", _sagemakerModelApprovalStatus); err != nil {
			log.Errorf("invalid --model-approval-status: %s", err.Error())
			return
		}
	}
	if len(_sagemakerModelPackageGroupName) > 0 {
		input.ModelPackageGroupName = aws.String(_sagemakerModelPackageGroupName)
	}
	if len(_sagemakerModelPackageType) > 0 {
		if err := assignInputField(input, "ModelPackageType", _sagemakerModelPackageType); err != nil {
			log.Errorf("invalid --model-package-type: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNameContains) > 0 {
		input.NameContains = aws.String(_sagemakerNameContains)
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListModelPackages(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListModelPackagesOutput
	p := sagemaker.NewListModelPackagesPaginator(client, input)
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

// Gets a list of model quality monitoring job definitions in your account.
func sagemaker_ListModelQualityJobDefinitions(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListModelQualityJobDefinitionsInput{}

	if len(_sagemakerCreationTimeAfter) > 0 {
		if err := assignInputField(input, "CreationTimeAfter", _sagemakerCreationTimeAfter); err != nil {
			log.Errorf("invalid --creation-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreationTimeBefore) > 0 {
		if err := assignInputField(input, "CreationTimeBefore", _sagemakerCreationTimeBefore); err != nil {
			log.Errorf("invalid --creation-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerEndpointName) > 0 {
		input.EndpointName = aws.String(_sagemakerEndpointName)
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNameContains) > 0 {
		input.NameContains = aws.String(_sagemakerNameContains)
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListModelQualityJobDefinitions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListModelQualityJobDefinitionsOutput
	p := sagemaker.NewListModelQualityJobDefinitionsPaginator(client, input)
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

// Lists models created with the CreateModel API.
func sagemaker_ListModels(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListModelsInput{}

	if len(_sagemakerCreationTimeAfter) > 0 {
		if err := assignInputField(input, "CreationTimeAfter", _sagemakerCreationTimeAfter); err != nil {
			log.Errorf("invalid --creation-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreationTimeBefore) > 0 {
		if err := assignInputField(input, "CreationTimeBefore", _sagemakerCreationTimeBefore); err != nil {
			log.Errorf("invalid --creation-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNameContains) > 0 {
		input.NameContains = aws.String(_sagemakerNameContains)
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListModels(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListModelsOutput
	p := sagemaker.NewListModelsPaginator(client, input)
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

// Gets a list of past alerts in a model monitoring schedule.
func sagemaker_ListMonitoringAlertHistory(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListMonitoringAlertHistoryInput{}

	if len(_sagemakerCreationTimeAfter) > 0 {
		if err := assignInputField(input, "CreationTimeAfter", _sagemakerCreationTimeAfter); err != nil {
			log.Errorf("invalid --creation-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreationTimeBefore) > 0 {
		if err := assignInputField(input, "CreationTimeBefore", _sagemakerCreationTimeBefore); err != nil {
			log.Errorf("invalid --creation-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMonitoringAlertName) > 0 {
		input.MonitoringAlertName = aws.String(_sagemakerMonitoringAlertName)
	}
	if len(_sagemakerMonitoringScheduleName) > 0 {
		input.MonitoringScheduleName = aws.String(_sagemakerMonitoringScheduleName)
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}
	if len(_sagemakerStatusEquals) > 0 {
		if err := assignInputField(input, "StatusEquals", _sagemakerStatusEquals); err != nil {
			log.Errorf("invalid --status-equals: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListMonitoringAlertHistory(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListMonitoringAlertHistoryOutput
	p := sagemaker.NewListMonitoringAlertHistoryPaginator(client, input)
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

// Gets the alerts for a single monitoring schedule.
func sagemaker_ListMonitoringAlerts(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListMonitoringAlertsInput{
		// MonitoringScheduleName: *string, // Required
	}

	if len(_sagemakerMonitoringScheduleName) > 0 {
		input.MonitoringScheduleName = aws.String(_sagemakerMonitoringScheduleName)
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListMonitoringAlerts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListMonitoringAlertsOutput
	p := sagemaker.NewListMonitoringAlertsPaginator(client, input)
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

// Returns list of all monitoring job executions.
func sagemaker_ListMonitoringExecutions(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListMonitoringExecutionsInput{}

	if len(_sagemakerCreationTimeAfter) > 0 {
		if err := assignInputField(input, "CreationTimeAfter", _sagemakerCreationTimeAfter); err != nil {
			log.Errorf("invalid --creation-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreationTimeBefore) > 0 {
		if err := assignInputField(input, "CreationTimeBefore", _sagemakerCreationTimeBefore); err != nil {
			log.Errorf("invalid --creation-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerEndpointName) > 0 {
		input.EndpointName = aws.String(_sagemakerEndpointName)
	}
	if len(_sagemakerLastModifiedTimeAfter) > 0 {
		if err := assignInputField(input, "LastModifiedTimeAfter", _sagemakerLastModifiedTimeAfter); err != nil {
			log.Errorf("invalid --last-modified-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerLastModifiedTimeBefore) > 0 {
		if err := assignInputField(input, "LastModifiedTimeBefore", _sagemakerLastModifiedTimeBefore); err != nil {
			log.Errorf("invalid --last-modified-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMonitoringJobDefinitionName) > 0 {
		input.MonitoringJobDefinitionName = aws.String(_sagemakerMonitoringJobDefinitionName)
	}
	if len(_sagemakerMonitoringScheduleName) > 0 {
		input.MonitoringScheduleName = aws.String(_sagemakerMonitoringScheduleName)
	}
	if len(_sagemakerMonitoringTypeEquals) > 0 {
		if err := assignInputField(input, "MonitoringTypeEquals", _sagemakerMonitoringTypeEquals); err != nil {
			log.Errorf("invalid --monitoring-type-equals: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerScheduledTimeAfter) > 0 {
		if err := assignInputField(input, "ScheduledTimeAfter", _sagemakerScheduledTimeAfter); err != nil {
			log.Errorf("invalid --scheduled-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerScheduledTimeBefore) > 0 {
		if err := assignInputField(input, "ScheduledTimeBefore", _sagemakerScheduledTimeBefore); err != nil {
			log.Errorf("invalid --scheduled-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}
	if len(_sagemakerStatusEquals) > 0 {
		if err := assignInputField(input, "StatusEquals", _sagemakerStatusEquals); err != nil {
			log.Errorf("invalid --status-equals: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListMonitoringExecutions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListMonitoringExecutionsOutput
	p := sagemaker.NewListMonitoringExecutionsPaginator(client, input)
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

// Returns list of all monitoring schedules.
func sagemaker_ListMonitoringSchedules(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListMonitoringSchedulesInput{}

	if len(_sagemakerCreationTimeAfter) > 0 {
		if err := assignInputField(input, "CreationTimeAfter", _sagemakerCreationTimeAfter); err != nil {
			log.Errorf("invalid --creation-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreationTimeBefore) > 0 {
		if err := assignInputField(input, "CreationTimeBefore", _sagemakerCreationTimeBefore); err != nil {
			log.Errorf("invalid --creation-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerEndpointName) > 0 {
		input.EndpointName = aws.String(_sagemakerEndpointName)
	}
	if len(_sagemakerLastModifiedTimeAfter) > 0 {
		if err := assignInputField(input, "LastModifiedTimeAfter", _sagemakerLastModifiedTimeAfter); err != nil {
			log.Errorf("invalid --last-modified-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerLastModifiedTimeBefore) > 0 {
		if err := assignInputField(input, "LastModifiedTimeBefore", _sagemakerLastModifiedTimeBefore); err != nil {
			log.Errorf("invalid --last-modified-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMonitoringJobDefinitionName) > 0 {
		input.MonitoringJobDefinitionName = aws.String(_sagemakerMonitoringJobDefinitionName)
	}
	if len(_sagemakerMonitoringTypeEquals) > 0 {
		if err := assignInputField(input, "MonitoringTypeEquals", _sagemakerMonitoringTypeEquals); err != nil {
			log.Errorf("invalid --monitoring-type-equals: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNameContains) > 0 {
		input.NameContains = aws.String(_sagemakerNameContains)
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}
	if len(_sagemakerStatusEquals) > 0 {
		if err := assignInputField(input, "StatusEquals", _sagemakerStatusEquals); err != nil {
			log.Errorf("invalid --status-equals: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListMonitoringSchedules(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListMonitoringSchedulesOutput
	p := sagemaker.NewListMonitoringSchedulesPaginator(client, input)
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

// Lists notebook instance lifestyle configurations created with the [CreateNotebookInstanceLifecycleConfig] API.
//
// [CreateNotebookInstanceLifecycleConfig]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateNotebookInstanceLifecycleConfig.html
func sagemaker_ListNotebookInstanceLifecycleConfigs(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListNotebookInstanceLifecycleConfigsInput{}

	if len(_sagemakerCreationTimeAfter) > 0 {
		if err := assignInputField(input, "CreationTimeAfter", _sagemakerCreationTimeAfter); err != nil {
			log.Errorf("invalid --creation-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreationTimeBefore) > 0 {
		if err := assignInputField(input, "CreationTimeBefore", _sagemakerCreationTimeBefore); err != nil {
			log.Errorf("invalid --creation-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerLastModifiedTimeAfter) > 0 {
		if err := assignInputField(input, "LastModifiedTimeAfter", _sagemakerLastModifiedTimeAfter); err != nil {
			log.Errorf("invalid --last-modified-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerLastModifiedTimeBefore) > 0 {
		if err := assignInputField(input, "LastModifiedTimeBefore", _sagemakerLastModifiedTimeBefore); err != nil {
			log.Errorf("invalid --last-modified-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNameContains) > 0 {
		input.NameContains = aws.String(_sagemakerNameContains)
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListNotebookInstanceLifecycleConfigs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListNotebookInstanceLifecycleConfigsOutput
	p := sagemaker.NewListNotebookInstanceLifecycleConfigsPaginator(client, input)
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

// Returns a list of the SageMaker AI notebook instances in the requester's
// account in an Amazon Web Services Region.
func sagemaker_ListNotebookInstances(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListNotebookInstancesInput{}

	if len(_sagemakerAdditionalCodeRepositoryEquals) > 0 {
		input.AdditionalCodeRepositoryEquals = aws.String(_sagemakerAdditionalCodeRepositoryEquals)
	}
	if len(_sagemakerCreationTimeAfter) > 0 {
		if err := assignInputField(input, "CreationTimeAfter", _sagemakerCreationTimeAfter); err != nil {
			log.Errorf("invalid --creation-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreationTimeBefore) > 0 {
		if err := assignInputField(input, "CreationTimeBefore", _sagemakerCreationTimeBefore); err != nil {
			log.Errorf("invalid --creation-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerDefaultCodeRepositoryContains) > 0 {
		input.DefaultCodeRepositoryContains = aws.String(_sagemakerDefaultCodeRepositoryContains)
	}
	if len(_sagemakerLastModifiedTimeAfter) > 0 {
		if err := assignInputField(input, "LastModifiedTimeAfter", _sagemakerLastModifiedTimeAfter); err != nil {
			log.Errorf("invalid --last-modified-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerLastModifiedTimeBefore) > 0 {
		if err := assignInputField(input, "LastModifiedTimeBefore", _sagemakerLastModifiedTimeBefore); err != nil {
			log.Errorf("invalid --last-modified-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNameContains) > 0 {
		input.NameContains = aws.String(_sagemakerNameContains)
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerNotebookInstanceLifecycleConfigNameContains) > 0 {
		input.NotebookInstanceLifecycleConfigNameContains = aws.String(_sagemakerNotebookInstanceLifecycleConfigNameContains)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}
	if len(_sagemakerStatusEquals) > 0 {
		if err := assignInputField(input, "StatusEquals", _sagemakerStatusEquals); err != nil {
			log.Errorf("invalid --status-equals: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListNotebookInstances(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListNotebookInstancesOutput
	p := sagemaker.NewListNotebookInstancesPaginator(client, input)
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

// Lists the optimization jobs in your account and their properties.
func sagemaker_ListOptimizationJobs(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListOptimizationJobsInput{}

	if len(_sagemakerCreationTimeAfter) > 0 {
		if err := assignInputField(input, "CreationTimeAfter", _sagemakerCreationTimeAfter); err != nil {
			log.Errorf("invalid --creation-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreationTimeBefore) > 0 {
		if err := assignInputField(input, "CreationTimeBefore", _sagemakerCreationTimeBefore); err != nil {
			log.Errorf("invalid --creation-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerLastModifiedTimeAfter) > 0 {
		if err := assignInputField(input, "LastModifiedTimeAfter", _sagemakerLastModifiedTimeAfter); err != nil {
			log.Errorf("invalid --last-modified-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerLastModifiedTimeBefore) > 0 {
		if err := assignInputField(input, "LastModifiedTimeBefore", _sagemakerLastModifiedTimeBefore); err != nil {
			log.Errorf("invalid --last-modified-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNameContains) > 0 {
		input.NameContains = aws.String(_sagemakerNameContains)
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerOptimizationContains) > 0 {
		input.OptimizationContains = aws.String(_sagemakerOptimizationContains)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}
	if len(_sagemakerStatusEquals) > 0 {
		if err := assignInputField(input, "StatusEquals", _sagemakerStatusEquals); err != nil {
			log.Errorf("invalid --status-equals: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListOptimizationJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListOptimizationJobsOutput
	p := sagemaker.NewListOptimizationJobsPaginator(client, input)
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

// Lists all of the SageMaker Partner AI Apps in an account.
func sagemaker_ListPartnerApps(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListPartnerAppsInput{}

	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPartnerApps(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListPartnerAppsOutput
	p := sagemaker.NewListPartnerAppsPaginator(client, input)
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

// Gets a list of PipeLineExecutionStep objects.
func sagemaker_ListPipelineExecutionSteps(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListPipelineExecutionStepsInput{}

	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerPipelineExecutionArn) > 0 {
		input.PipelineExecutionArn = aws.String(_sagemakerPipelineExecutionArn)
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListPipelineExecutionSteps(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListPipelineExecutionStepsOutput
	p := sagemaker.NewListPipelineExecutionStepsPaginator(client, input)
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

// Gets a list of the pipeline executions.
func sagemaker_ListPipelineExecutions(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListPipelineExecutionsInput{
		// PipelineName: *string, // Required
	}

	if len(_sagemakerPipelineName) > 0 {
		input.PipelineName = aws.String(_sagemakerPipelineName)
	}
	if len(_sagemakerCreatedAfter) > 0 {
		if err := assignInputField(input, "CreatedAfter", _sagemakerCreatedAfter); err != nil {
			log.Errorf("invalid --created-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreatedBefore) > 0 {
		if err := assignInputField(input, "CreatedBefore", _sagemakerCreatedBefore); err != nil {
			log.Errorf("invalid --created-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
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

	var results []*sagemaker.ListPipelineExecutionsOutput
	p := sagemaker.NewListPipelineExecutionsPaginator(client, input)
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

// Gets a list of parameters for a pipeline execution.
func sagemaker_ListPipelineParametersForExecution(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListPipelineParametersForExecutionInput{
		// PipelineExecutionArn: *string, // Required
	}

	if len(_sagemakerPipelineExecutionArn) > 0 {
		input.PipelineExecutionArn = aws.String(_sagemakerPipelineExecutionArn)
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPipelineParametersForExecution(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListPipelineParametersForExecutionOutput
	p := sagemaker.NewListPipelineParametersForExecutionPaginator(client, input)
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

// Gets a list of all versions of the pipeline.
func sagemaker_ListPipelineVersions(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListPipelineVersionsInput{
		// PipelineName: *string, // Required
	}

	if len(_sagemakerPipelineName) > 0 {
		input.PipelineName = aws.String(_sagemakerPipelineName)
	}
	if len(_sagemakerCreatedAfter) > 0 {
		if err := assignInputField(input, "CreatedAfter", _sagemakerCreatedAfter); err != nil {
			log.Errorf("invalid --created-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreatedBefore) > 0 {
		if err := assignInputField(input, "CreatedBefore", _sagemakerCreatedBefore); err != nil {
			log.Errorf("invalid --created-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListPipelineVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListPipelineVersionsOutput
	p := sagemaker.NewListPipelineVersionsPaginator(client, input)
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

// Gets a list of pipelines.
func sagemaker_ListPipelines(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListPipelinesInput{}

	if len(_sagemakerCreatedAfter) > 0 {
		if err := assignInputField(input, "CreatedAfter", _sagemakerCreatedAfter); err != nil {
			log.Errorf("invalid --created-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreatedBefore) > 0 {
		if err := assignInputField(input, "CreatedBefore", _sagemakerCreatedBefore); err != nil {
			log.Errorf("invalid --created-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerPipelineNamePrefix) > 0 {
		input.PipelineNamePrefix = aws.String(_sagemakerPipelineNamePrefix)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
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

	var results []*sagemaker.ListPipelinesOutput
	p := sagemaker.NewListPipelinesPaginator(client, input)
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

// Lists processing jobs that satisfy various filters.
func sagemaker_ListProcessingJobs(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListProcessingJobsInput{}

	if len(_sagemakerCreationTimeAfter) > 0 {
		if err := assignInputField(input, "CreationTimeAfter", _sagemakerCreationTimeAfter); err != nil {
			log.Errorf("invalid --creation-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreationTimeBefore) > 0 {
		if err := assignInputField(input, "CreationTimeBefore", _sagemakerCreationTimeBefore); err != nil {
			log.Errorf("invalid --creation-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerLastModifiedTimeAfter) > 0 {
		if err := assignInputField(input, "LastModifiedTimeAfter", _sagemakerLastModifiedTimeAfter); err != nil {
			log.Errorf("invalid --last-modified-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerLastModifiedTimeBefore) > 0 {
		if err := assignInputField(input, "LastModifiedTimeBefore", _sagemakerLastModifiedTimeBefore); err != nil {
			log.Errorf("invalid --last-modified-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNameContains) > 0 {
		input.NameContains = aws.String(_sagemakerNameContains)
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}
	if len(_sagemakerStatusEquals) > 0 {
		if err := assignInputField(input, "StatusEquals", _sagemakerStatusEquals); err != nil {
			log.Errorf("invalid --status-equals: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListProcessingJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListProcessingJobsOutput
	p := sagemaker.NewListProcessingJobsPaginator(client, input)
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

// Gets a list of the projects in an Amazon Web Services account.
func sagemaker_ListProjects(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListProjectsInput{}

	if len(_sagemakerCreationTimeAfter) > 0 {
		if err := assignInputField(input, "CreationTimeAfter", _sagemakerCreationTimeAfter); err != nil {
			log.Errorf("invalid --creation-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreationTimeBefore) > 0 {
		if err := assignInputField(input, "CreationTimeBefore", _sagemakerCreationTimeBefore); err != nil {
			log.Errorf("invalid --creation-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNameContains) > 0 {
		input.NameContains = aws.String(_sagemakerNameContains)
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
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

	var results []*sagemaker.ListProjectsOutput
	p := sagemaker.NewListProjectsPaginator(client, input)
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

// Lists Amazon SageMaker Catalogs based on given filters and orders. The maximum
// number of ResourceCatalog s viewable is 1000.
func sagemaker_ListResourceCatalogs(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListResourceCatalogsInput{}

	if len(_sagemakerCreationTimeAfter) > 0 {
		if err := assignInputField(input, "CreationTimeAfter", _sagemakerCreationTimeAfter); err != nil {
			log.Errorf("invalid --creation-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreationTimeBefore) > 0 {
		if err := assignInputField(input, "CreationTimeBefore", _sagemakerCreationTimeBefore); err != nil {
			log.Errorf("invalid --creation-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNameContains) > 0 {
		input.NameContains = aws.String(_sagemakerNameContains)
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListResourceCatalogs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListResourceCatalogsOutput
	p := sagemaker.NewListResourceCatalogsPaginator(client, input)
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

// Lists spaces.
func sagemaker_ListSpaces(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListSpacesInput{}

	if len(_sagemakerDomainIdEquals) > 0 {
		input.DomainIdEquals = aws.String(_sagemakerDomainIdEquals)
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSpaceNameContains) > 0 {
		input.SpaceNameContains = aws.String(_sagemakerSpaceNameContains)
	}

	if disablePaginator() {
		if resp, err := client.ListSpaces(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListSpacesOutput
	p := sagemaker.NewListSpacesPaginator(client, input)
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

// Lists devices allocated to the stage, containing detailed device information
// and deployment status.
func sagemaker_ListStageDevices(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListStageDevicesInput{
		// EdgeDeploymentPlanName: *string, // Required
		// StageName: *string, // Required
	}

	if len(_sagemakerEdgeDeploymentPlanName) > 0 {
		input.EdgeDeploymentPlanName = aws.String(_sagemakerEdgeDeploymentPlanName)
	}
	if len(_sagemakerStageName) > 0 {
		input.StageName = aws.String(_sagemakerStageName)
	}
	if len(_sagemakerExcludeDevicesDeployedInOtherStage) > 0 {
		if err := assignInputField(input, "ExcludeDevicesDeployedInOtherStage", _sagemakerExcludeDevicesDeployedInOtherStage); err != nil {
			log.Errorf("invalid --exclude-devices-deployed-in-other-stage: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListStageDevices(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListStageDevicesOutput
	p := sagemaker.NewListStageDevicesPaginator(client, input)
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

// Lists the Amazon SageMaker AI Studio Lifecycle Configurations in your Amazon
// Web Services Account.
func sagemaker_ListStudioLifecycleConfigs(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListStudioLifecycleConfigsInput{}

	if len(_sagemakerAppTypeEquals) > 0 {
		if err := assignInputField(input, "AppTypeEquals", _sagemakerAppTypeEquals); err != nil {
			log.Errorf("invalid --app-type-equals: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreationTimeAfter) > 0 {
		if err := assignInputField(input, "CreationTimeAfter", _sagemakerCreationTimeAfter); err != nil {
			log.Errorf("invalid --creation-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreationTimeBefore) > 0 {
		if err := assignInputField(input, "CreationTimeBefore", _sagemakerCreationTimeBefore); err != nil {
			log.Errorf("invalid --creation-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerModifiedTimeAfter) > 0 {
		if err := assignInputField(input, "ModifiedTimeAfter", _sagemakerModifiedTimeAfter); err != nil {
			log.Errorf("invalid --modified-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerModifiedTimeBefore) > 0 {
		if err := assignInputField(input, "ModifiedTimeBefore", _sagemakerModifiedTimeBefore); err != nil {
			log.Errorf("invalid --modified-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNameContains) > 0 {
		input.NameContains = aws.String(_sagemakerNameContains)
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListStudioLifecycleConfigs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListStudioLifecycleConfigsOutput
	p := sagemaker.NewListStudioLifecycleConfigsPaginator(client, input)
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

// Gets a list of the work teams that you are subscribed to in the Amazon Web
// Services Marketplace. The list may be empty if no work team satisfies the filter
// specified in the NameContains parameter.
func sagemaker_ListSubscribedWorkteams(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListSubscribedWorkteamsInput{}

	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNameContains) > 0 {
		input.NameContains = aws.String(_sagemakerNameContains)
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSubscribedWorkteams(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListSubscribedWorkteamsOutput
	p := sagemaker.NewListSubscribedWorkteamsPaginator(client, input)
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

// Returns the tags for the specified SageMaker resource.
func sagemaker_ListTags(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListTagsInput{
		// ResourceArn: *string, // Required
	}

	if len(_sagemakerResourceArn) > 0 {
		input.ResourceArn = aws.String(_sagemakerResourceArn)
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTags(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListTagsOutput
	p := sagemaker.NewListTagsPaginator(client, input)
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

// Lists training jobs.
// When StatusEquals and MaxResults are set at the same time, the MaxResults
// number of training jobs are first retrieved ignoring the StatusEquals parameter
// and then they are filtered by the StatusEquals parameter, which is returned as
// a response.
//
// For example, if ListTrainingJobs is invoked with the following parameters:
//
// { ... MaxResults: 100, StatusEquals: InProgress ... }
//
// First, 100 trainings jobs with any status, including those other than InProgress
// , are selected (sorted according to the creation time, from the most current to
// the oldest). Next, those with a status of InProgress are returned.
//
// You can quickly test the API using the following Amazon Web Services CLI code.
//
// aws sagemaker list-training-jobs --max-results 100 --status-equals InProgress
func sagemaker_ListTrainingJobs(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListTrainingJobsInput{}

	if len(_sagemakerCreationTimeAfter) > 0 {
		if err := assignInputField(input, "CreationTimeAfter", _sagemakerCreationTimeAfter); err != nil {
			log.Errorf("invalid --creation-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreationTimeBefore) > 0 {
		if err := assignInputField(input, "CreationTimeBefore", _sagemakerCreationTimeBefore); err != nil {
			log.Errorf("invalid --creation-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerLastModifiedTimeAfter) > 0 {
		if err := assignInputField(input, "LastModifiedTimeAfter", _sagemakerLastModifiedTimeAfter); err != nil {
			log.Errorf("invalid --last-modified-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerLastModifiedTimeBefore) > 0 {
		if err := assignInputField(input, "LastModifiedTimeBefore", _sagemakerLastModifiedTimeBefore); err != nil {
			log.Errorf("invalid --last-modified-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNameContains) > 0 {
		input.NameContains = aws.String(_sagemakerNameContains)
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}
	if len(_sagemakerStatusEquals) > 0 {
		if err := assignInputField(input, "StatusEquals", _sagemakerStatusEquals); err != nil {
			log.Errorf("invalid --status-equals: %s", err.Error())
			return
		}
	}
	if len(_sagemakerTrainingPlanArnEquals) > 0 {
		input.TrainingPlanArnEquals = aws.String(_sagemakerTrainingPlanArnEquals)
	}
	if len(_sagemakerWarmPoolStatusEquals) > 0 {
		if err := assignInputField(input, "WarmPoolStatusEquals", _sagemakerWarmPoolStatusEquals); err != nil {
			log.Errorf("invalid --warm-pool-status-equals: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListTrainingJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListTrainingJobsOutput
	p := sagemaker.NewListTrainingJobsPaginator(client, input)
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

// Gets a list of [TrainingJobSummary] objects that describe the training jobs that a hyperparameter
// tuning job launched.
//
// [TrainingJobSummary]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_TrainingJobSummary.html
func sagemaker_ListTrainingJobsForHyperParameterTuningJob(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListTrainingJobsForHyperParameterTuningJobInput{
		// HyperParameterTuningJobName: *string, // Required
	}

	if len(_sagemakerHyperParameterTuningJobName) > 0 {
		input.HyperParameterTuningJobName = aws.String(_sagemakerHyperParameterTuningJobName)
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}
	if len(_sagemakerStatusEquals) > 0 {
		if err := assignInputField(input, "StatusEquals", _sagemakerStatusEquals); err != nil {
			log.Errorf("invalid --status-equals: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListTrainingJobsForHyperParameterTuningJob(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListTrainingJobsForHyperParameterTuningJobOutput
	p := sagemaker.NewListTrainingJobsForHyperParameterTuningJobPaginator(client, input)
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

// Retrieves a list of training plans for the current account.
func sagemaker_ListTrainingPlans(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListTrainingPlansInput{}

	if len(_sagemakerFilters) > 0 {
		if err := assignInputField(input, "Filters", _sagemakerFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}
	if len(_sagemakerStartTimeAfter) > 0 {
		if err := assignInputField(input, "StartTimeAfter", _sagemakerStartTimeAfter); err != nil {
			log.Errorf("invalid --start-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerStartTimeBefore) > 0 {
		if err := assignInputField(input, "StartTimeBefore", _sagemakerStartTimeBefore); err != nil {
			log.Errorf("invalid --start-time-before: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListTrainingPlans(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListTrainingPlansOutput
	p := sagemaker.NewListTrainingPlansPaginator(client, input)
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

// Lists transform jobs.
func sagemaker_ListTransformJobs(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListTransformJobsInput{}

	if len(_sagemakerCreationTimeAfter) > 0 {
		if err := assignInputField(input, "CreationTimeAfter", _sagemakerCreationTimeAfter); err != nil {
			log.Errorf("invalid --creation-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreationTimeBefore) > 0 {
		if err := assignInputField(input, "CreationTimeBefore", _sagemakerCreationTimeBefore); err != nil {
			log.Errorf("invalid --creation-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerLastModifiedTimeAfter) > 0 {
		if err := assignInputField(input, "LastModifiedTimeAfter", _sagemakerLastModifiedTimeAfter); err != nil {
			log.Errorf("invalid --last-modified-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerLastModifiedTimeBefore) > 0 {
		if err := assignInputField(input, "LastModifiedTimeBefore", _sagemakerLastModifiedTimeBefore); err != nil {
			log.Errorf("invalid --last-modified-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNameContains) > 0 {
		input.NameContains = aws.String(_sagemakerNameContains)
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}
	if len(_sagemakerStatusEquals) > 0 {
		if err := assignInputField(input, "StatusEquals", _sagemakerStatusEquals); err != nil {
			log.Errorf("invalid --status-equals: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListTransformJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListTransformJobsOutput
	p := sagemaker.NewListTransformJobsPaginator(client, input)
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

// Lists the trial components in your account. You can sort the list by trial
// component name or creation time. You can filter the list to show only components
// that were created in a specific time range. You can also filter on one of the
// following:
//
// - ExperimentName
//
// - SourceArn
//
// - TrialName
func sagemaker_ListTrialComponents(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListTrialComponentsInput{}

	if len(_sagemakerCreatedAfter) > 0 {
		if err := assignInputField(input, "CreatedAfter", _sagemakerCreatedAfter); err != nil {
			log.Errorf("invalid --created-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreatedBefore) > 0 {
		if err := assignInputField(input, "CreatedBefore", _sagemakerCreatedBefore); err != nil {
			log.Errorf("invalid --created-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerExperimentName) > 0 {
		input.ExperimentName = aws.String(_sagemakerExperimentName)
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSourceArn) > 0 {
		input.SourceArn = aws.String(_sagemakerSourceArn)
	}
	if len(_sagemakerTrialName) > 0 {
		input.TrialName = aws.String(_sagemakerTrialName)
	}

	if disablePaginator() {
		if resp, err := client.ListTrialComponents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListTrialComponentsOutput
	p := sagemaker.NewListTrialComponentsPaginator(client, input)
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

// Lists the trials in your account. Specify an experiment name to limit the list
// to the trials that are part of that experiment. Specify a trial component name
// to limit the list to the trials that associated with that trial component. The
// list can be filtered to show only trials that were created in a specific time
// range. The list can be sorted by trial name or creation time.
func sagemaker_ListTrials(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListTrialsInput{}

	if len(_sagemakerCreatedAfter) > 0 {
		if err := assignInputField(input, "CreatedAfter", _sagemakerCreatedAfter); err != nil {
			log.Errorf("invalid --created-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCreatedBefore) > 0 {
		if err := assignInputField(input, "CreatedBefore", _sagemakerCreatedBefore); err != nil {
			log.Errorf("invalid --created-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerExperimentName) > 0 {
		input.ExperimentName = aws.String(_sagemakerExperimentName)
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}
	if len(_sagemakerTrialComponentName) > 0 {
		input.TrialComponentName = aws.String(_sagemakerTrialComponentName)
	}

	if disablePaginator() {
		if resp, err := client.ListTrials(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListTrialsOutput
	p := sagemaker.NewListTrialsPaginator(client, input)
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

// Lists all UltraServers that are part of a specified reserved capacity.
func sagemaker_ListUltraServersByReservedCapacity(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListUltraServersByReservedCapacityInput{
		// ReservedCapacityArn: *string, // Required
	}

	if len(_sagemakerReservedCapacityArn) > 0 {
		input.ReservedCapacityArn = aws.String(_sagemakerReservedCapacityArn)
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListUltraServersByReservedCapacity(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListUltraServersByReservedCapacityOutput
	p := sagemaker.NewListUltraServersByReservedCapacityPaginator(client, input)
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

// Lists user profiles.
func sagemaker_ListUserProfiles(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListUserProfilesInput{}

	if len(_sagemakerDomainIdEquals) > 0 {
		input.DomainIdEquals = aws.String(_sagemakerDomainIdEquals)
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}
	if len(_sagemakerUserProfileNameContains) > 0 {
		input.UserProfileNameContains = aws.String(_sagemakerUserProfileNameContains)
	}

	if disablePaginator() {
		if resp, err := client.ListUserProfiles(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListUserProfilesOutput
	p := sagemaker.NewListUserProfilesPaginator(client, input)
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

// Use this operation to list all private and vendor workforces in an Amazon Web
// Services Region. Note that you can only have one private workforce per Amazon
// Web Services Region.
func sagemaker_ListWorkforces(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListWorkforcesInput{}

	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNameContains) > 0 {
		input.NameContains = aws.String(_sagemakerNameContains)
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListWorkforces(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListWorkforcesOutput
	p := sagemaker.NewListWorkforcesPaginator(client, input)
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

// Gets a list of private work teams that you have defined in a region. The list
// may be empty if no work team satisfies the filter specified in the NameContains
// parameter.
func sagemaker_ListWorkteams(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.ListWorkteamsInput{}

	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNameContains) > 0 {
		input.NameContains = aws.String(_sagemakerNameContains)
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _sagemakerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListWorkteams(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.ListWorkteamsOutput
	p := sagemaker.NewListWorkteamsPaginator(client, input)
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

// Adds a resouce policy to control access to a model group. For information about
// resoure policies, see [Identity-based policies and resource-based policies]in the Amazon Web Services Identity and Access Management
// User Guide..
//
// [Identity-based policies and resource-based policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_identity-vs-resource.html
func sagemaker_PutModelPackageGroupPolicy(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.PutModelPackageGroupPolicyInput{
		// ModelPackageGroupName: *string, // Required
		// ResourcePolicy: *string, // Required
	}

	if len(_sagemakerModelPackageGroupName) > 0 {
		input.ModelPackageGroupName = aws.String(_sagemakerModelPackageGroupName)
	}
	if len(_sagemakerResourcePolicy) > 0 {
		input.ResourcePolicy = aws.String(_sagemakerResourcePolicy)
	}

	if resp, err := client.PutModelPackageGroupPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Use this action to inspect your lineage and discover relationships between
// entities. For more information, see [Querying Lineage Entities]in the Amazon SageMaker Developer Guide.
//
// [Querying Lineage Entities]: https://docs.aws.amazon.com/sagemaker/latest/dg/querying-lineage-entities.html
func sagemaker_QueryLineage(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.QueryLineageInput{}

	if len(_sagemakerDirection) > 0 {
		if err := assignInputField(input, "Direction", _sagemakerDirection); err != nil {
			log.Errorf("invalid --direction: %s", err.Error())
			return
		}
	}
	if len(_sagemakerFilters) > 0 {
		if err := assignInputField(input, "Filters", _sagemakerFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_sagemakerIncludeEdges) > 0 {
		if err := assignInputField(input, "IncludeEdges", _sagemakerIncludeEdges); err != nil {
			log.Errorf("invalid --include-edges: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMaxDepth) > 0 {
		if err := assignInputField(input, "MaxDepth", _sagemakerMaxDepth); err != nil {
			log.Errorf("invalid --max-depth: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerStartArns) > 0 {
		input.StartArns = append([]string(nil), _sagemakerStartArns...)
	}

	if disablePaginator() {
		if resp, err := client.QueryLineage(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.QueryLineageOutput
	p := sagemaker.NewQueryLineagePaginator(client, input)
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

// Register devices.
func sagemaker_RegisterDevices(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.RegisterDevicesInput{
		// DeviceFleetName: *string, // Required
		// Devices: []types.Device, // Required
	}

	if len(_sagemakerDeviceFleetName) > 0 {
		input.DeviceFleetName = aws.String(_sagemakerDeviceFleetName)
	}
	if len(_sagemakerDevices) > 0 {
		if err := assignInputField(input, "Devices", _sagemakerDevices); err != nil {
			log.Errorf("invalid --devices: %s", err.Error())
			return
		}
	}
	if len(_sagemakerTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.RegisterDevices(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Renders the UI template so that you can preview the worker's experience.
func sagemaker_RenderUiTemplate(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.RenderUiTemplateInput{
		// RoleArn: *string, // Required
		// Task: *types.RenderableTask, // Required
	}

	if len(_sagemakerRoleArn) > 0 {
		input.RoleArn = aws.String(_sagemakerRoleArn)
	}
	if len(_sagemakerTask) > 0 {
		if err := assignInputField(input, "Task", _sagemakerTask); err != nil {
			log.Errorf("invalid --task: %s", err.Error())
			return
		}
	}
	if len(_sagemakerHumanTaskUiArn) > 0 {
		input.HumanTaskUiArn = aws.String(_sagemakerHumanTaskUiArn)
	}
	if len(_sagemakerUiTemplate) > 0 {
		if err := assignInputField(input, "UiTemplate", _sagemakerUiTemplate); err != nil {
			log.Errorf("invalid --ui-template: %s", err.Error())
			return
		}
	}

	if resp, err := client.RenderUiTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retry the execution of the pipeline.
func sagemaker_RetryPipelineExecution(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.RetryPipelineExecutionInput{
		// ClientRequestToken: *string, // Required
		// PipelineExecutionArn: *string, // Required
	}

	if len(_sagemakerClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_sagemakerClientRequestToken)
	}
	if len(_sagemakerPipelineExecutionArn) > 0 {
		input.PipelineExecutionArn = aws.String(_sagemakerPipelineExecutionArn)
	}
	if len(_sagemakerParallelismConfiguration) > 0 {
		if err := assignInputField(input, "ParallelismConfiguration", _sagemakerParallelismConfiguration); err != nil {
			log.Errorf("invalid --parallelism-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.RetryPipelineExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Finds SageMaker resources that match a search query. Matching resources are
// returned as a list of SearchRecord objects in the response. You can sort the
// search results by any resource property in a ascending or descending order.
//
// You can query against the following value types: numeric, text, Boolean, and
// timestamp.
//
// The Search API may provide access to otherwise restricted data. See [Amazon SageMaker API Permissions: Actions, Permissions, and Resources Reference] for more
// information.
//
// [Amazon SageMaker API Permissions: Actions, Permissions, and Resources Reference]: https://docs.aws.amazon.com/sagemaker/latest/dg/api-permissions-reference.html
func sagemaker_Search(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.SearchInput{
		// Resource: types.ResourceType, // Required
	}

	if len(_sagemakerResource) > 0 {
		if err := assignInputField(input, "Resource", _sagemakerResource); err != nil {
			log.Errorf("invalid --resource: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCrossAccountFilterOption) > 0 {
		if err := assignInputField(input, "CrossAccountFilterOption", _sagemakerCrossAccountFilterOption); err != nil {
			log.Errorf("invalid --cross-account-filter-option: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sagemakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNextToken) > 0 {
		input.NextToken = aws.String(_sagemakerNextToken)
	}
	if len(_sagemakerSearchExpression) > 0 {
		if err := assignInputField(input, "SearchExpression", _sagemakerSearchExpression); err != nil {
			log.Errorf("invalid --search-expression: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSortBy) > 0 {
		input.SortBy = aws.String(_sagemakerSortBy)
	}
	if len(_sagemakerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _sagemakerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}
	if len(_sagemakerVisibilityConditions) > 0 {
		if err := assignInputField(input, "VisibilityConditions", _sagemakerVisibilityConditions); err != nil {
			log.Errorf("invalid --visibility-conditions: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.Search(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sagemaker.SearchOutput
	p := sagemaker.NewSearchPaginator(client, input)
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

// Searches for available training plan offerings based on specified criteria.
// - Users search for available plan offerings based on their requirements
// (e.g., instance type, count, start time, duration).
//
// - And then, they create a plan that best matches their needs using the ID of
// the plan offering they want to use.
//
// For more information about how to reserve GPU capacity for your SageMaker
// training jobs or SageMaker HyperPod clusters using Amazon SageMaker Training
// Plan , see [CreateTrainingPlan].
//
// [CreateTrainingPlan]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateTrainingPlan.html
func sagemaker_SearchTrainingPlanOfferings(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.SearchTrainingPlanOfferingsInput{}

	if len(_sagemakerDurationHours) > 0 {
		if err := assignInputField(input, "DurationHours", _sagemakerDurationHours); err != nil {
			log.Errorf("invalid --duration-hours: %s", err.Error())
			return
		}
	}
	if len(_sagemakerEndTimeBefore) > 0 {
		if err := assignInputField(input, "EndTimeBefore", _sagemakerEndTimeBefore); err != nil {
			log.Errorf("invalid --end-time-before: %s", err.Error())
			return
		}
	}
	if len(_sagemakerInstanceCount) > 0 {
		if err := assignInputField(input, "InstanceCount", _sagemakerInstanceCount); err != nil {
			log.Errorf("invalid --instance-count: %s", err.Error())
			return
		}
	}
	if len(_sagemakerInstanceType) > 0 {
		if err := assignInputField(input, "InstanceType", _sagemakerInstanceType); err != nil {
			log.Errorf("invalid --instance-type: %s", err.Error())
			return
		}
	}
	if len(_sagemakerStartTimeAfter) > 0 {
		if err := assignInputField(input, "StartTimeAfter", _sagemakerStartTimeAfter); err != nil {
			log.Errorf("invalid --start-time-after: %s", err.Error())
			return
		}
	}
	if len(_sagemakerTargetResources) > 0 {
		if err := assignInputField(input, "TargetResources", _sagemakerTargetResources); err != nil {
			log.Errorf("invalid --target-resources: %s", err.Error())
			return
		}
	}
	if len(_sagemakerUltraServerCount) > 0 {
		if err := assignInputField(input, "UltraServerCount", _sagemakerUltraServerCount); err != nil {
			log.Errorf("invalid --ultra-server-count: %s", err.Error())
			return
		}
	}
	if len(_sagemakerUltraServerType) > 0 {
		input.UltraServerType = aws.String(_sagemakerUltraServerType)
	}

	if resp, err := client.SearchTrainingPlanOfferings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Notifies the pipeline that the execution of a callback step failed, along with
// a message describing why. When a callback step is run, the pipeline generates a
// callback token and includes the token in a message sent to Amazon Simple Queue
// Service (Amazon SQS).
func sagemaker_SendPipelineExecutionStepFailure(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.SendPipelineExecutionStepFailureInput{
		// CallbackToken: *string, // Required
	}

	if len(_sagemakerCallbackToken) > 0 {
		input.CallbackToken = aws.String(_sagemakerCallbackToken)
	}
	if len(_sagemakerClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_sagemakerClientRequestToken)
	}
	if len(_sagemakerFailureReason) > 0 {
		input.FailureReason = aws.String(_sagemakerFailureReason)
	}

	if resp, err := client.SendPipelineExecutionStepFailure(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Notifies the pipeline that the execution of a callback step succeeded and
// provides a list of the step's output parameters. When a callback step is run,
// the pipeline generates a callback token and includes the token in a message sent
// to Amazon Simple Queue Service (Amazon SQS).
func sagemaker_SendPipelineExecutionStepSuccess(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.SendPipelineExecutionStepSuccessInput{
		// CallbackToken: *string, // Required
	}

	if len(_sagemakerCallbackToken) > 0 {
		input.CallbackToken = aws.String(_sagemakerCallbackToken)
	}
	if len(_sagemakerClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_sagemakerClientRequestToken)
	}
	if len(_sagemakerOutputParameters) > 0 {
		if err := assignInputField(input, "OutputParameters", _sagemakerOutputParameters); err != nil {
			log.Errorf("invalid --output-parameters: %s", err.Error())
			return
		}
	}

	if resp, err := client.SendPipelineExecutionStepSuccess(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a stage in an edge deployment plan.
func sagemaker_StartEdgeDeploymentStage(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.StartEdgeDeploymentStageInput{
		// EdgeDeploymentPlanName: *string, // Required
		// StageName: *string, // Required
	}

	if len(_sagemakerEdgeDeploymentPlanName) > 0 {
		input.EdgeDeploymentPlanName = aws.String(_sagemakerEdgeDeploymentPlanName)
	}
	if len(_sagemakerStageName) > 0 {
		input.StageName = aws.String(_sagemakerStageName)
	}

	if resp, err := client.StartEdgeDeploymentStage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts an inference experiment.
func sagemaker_StartInferenceExperiment(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.StartInferenceExperimentInput{
		// Name: *string, // Required
	}

	if len(_sagemakerName) > 0 {
		input.Name = aws.String(_sagemakerName)
	}

	if resp, err := client.StartInferenceExperiment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Programmatically start an MLflow Tracking Server.
func sagemaker_StartMlflowTrackingServer(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.StartMlflowTrackingServerInput{
		// TrackingServerName: *string, // Required
	}

	if len(_sagemakerTrackingServerName) > 0 {
		input.TrackingServerName = aws.String(_sagemakerTrackingServerName)
	}

	if resp, err := client.StartMlflowTrackingServer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a previously stopped monitoring schedule.
// By default, when you successfully create a new schedule, the status of a
// monitoring schedule is scheduled .
func sagemaker_StartMonitoringSchedule(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.StartMonitoringScheduleInput{
		// MonitoringScheduleName: *string, // Required
	}

	if len(_sagemakerMonitoringScheduleName) > 0 {
		input.MonitoringScheduleName = aws.String(_sagemakerMonitoringScheduleName)
	}

	if resp, err := client.StartMonitoringSchedule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Launches an ML compute instance with the latest version of the libraries and
// attaches your ML storage volume. After configuring the notebook instance,
// SageMaker AI sets the notebook instance status to InService . A notebook
// instance's status must be InService before you can connect to your Jupyter
// notebook.
func sagemaker_StartNotebookInstance(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.StartNotebookInstanceInput{
		// NotebookInstanceName: *string, // Required
	}

	if len(_sagemakerNotebookInstanceName) > 0 {
		input.NotebookInstanceName = aws.String(_sagemakerNotebookInstanceName)
	}

	if resp, err := client.StartNotebookInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a pipeline execution.
func sagemaker_StartPipelineExecution(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.StartPipelineExecutionInput{
		// ClientRequestToken: *string, // Required
		// PipelineName: *string, // Required
	}

	if len(_sagemakerClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_sagemakerClientRequestToken)
	}
	if len(_sagemakerPipelineName) > 0 {
		input.PipelineName = aws.String(_sagemakerPipelineName)
	}
	if len(_sagemakerMlflowExperimentName) > 0 {
		input.MlflowExperimentName = aws.String(_sagemakerMlflowExperimentName)
	}
	if len(_sagemakerParallelismConfiguration) > 0 {
		if err := assignInputField(input, "ParallelismConfiguration", _sagemakerParallelismConfiguration); err != nil {
			log.Errorf("invalid --parallelism-configuration: %s", err.Error())
			return
		}
	}
	if len(_sagemakerPipelineExecutionDescription) > 0 {
		input.PipelineExecutionDescription = aws.String(_sagemakerPipelineExecutionDescription)
	}
	if len(_sagemakerPipelineExecutionDisplayName) > 0 {
		input.PipelineExecutionDisplayName = aws.String(_sagemakerPipelineExecutionDisplayName)
	}
	if len(_sagemakerPipelineParameters) > 0 {
		if err := assignInputField(input, "PipelineParameters", _sagemakerPipelineParameters); err != nil {
			log.Errorf("invalid --pipeline-parameters: %s", err.Error())
			return
		}
	}
	if len(_sagemakerPipelineVersionId) > 0 {
		if err := assignInputField(input, "PipelineVersionId", _sagemakerPipelineVersionId); err != nil {
			log.Errorf("invalid --pipeline-version-id: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSelectiveExecutionConfig) > 0 {
		if err := assignInputField(input, "SelectiveExecutionConfig", _sagemakerSelectiveExecutionConfig); err != nil {
			log.Errorf("invalid --selective-execution-config: %s", err.Error())
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

// Initiates a remote connection session between a local integrated development
// environments (IDEs) and a remote SageMaker space.
func sagemaker_StartSession(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.StartSessionInput{
		// ResourceIdentifier: *string, // Required
	}

	if len(_sagemakerResourceIdentifier) > 0 {
		input.ResourceIdentifier = aws.String(_sagemakerResourceIdentifier)
	}

	if resp, err := client.StartSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// A method for forcing a running job to shut down.
func sagemaker_StopAutoMLJob(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.StopAutoMLJobInput{
		// AutoMLJobName: *string, // Required
	}

	if len(_sagemakerAutoMLJobName) > 0 {
		input.AutoMLJobName = aws.String(_sagemakerAutoMLJobName)
	}

	if resp, err := client.StopAutoMLJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a model compilation job.
// To stop a job, Amazon SageMaker AI sends the algorithm the SIGTERM signal. This
// gracefully shuts the job down. If the job hasn't stopped, it sends the SIGKILL
// signal.
//
// When it receives a StopCompilationJob request, Amazon SageMaker AI changes the
// CompilationJobStatus of the job to Stopping . After Amazon SageMaker stops the
// job, it sets the CompilationJobStatus to Stopped .
func sagemaker_StopCompilationJob(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.StopCompilationJobInput{
		// CompilationJobName: *string, // Required
	}

	if len(_sagemakerCompilationJobName) > 0 {
		input.CompilationJobName = aws.String(_sagemakerCompilationJobName)
	}

	if resp, err := client.StopCompilationJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a stage in an edge deployment plan.
func sagemaker_StopEdgeDeploymentStage(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.StopEdgeDeploymentStageInput{
		// EdgeDeploymentPlanName: *string, // Required
		// StageName: *string, // Required
	}

	if len(_sagemakerEdgeDeploymentPlanName) > 0 {
		input.EdgeDeploymentPlanName = aws.String(_sagemakerEdgeDeploymentPlanName)
	}
	if len(_sagemakerStageName) > 0 {
		input.StageName = aws.String(_sagemakerStageName)
	}

	if resp, err := client.StopEdgeDeploymentStage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Request to stop an edge packaging job.
func sagemaker_StopEdgePackagingJob(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.StopEdgePackagingJobInput{
		// EdgePackagingJobName: *string, // Required
	}

	if len(_sagemakerEdgePackagingJobName) > 0 {
		input.EdgePackagingJobName = aws.String(_sagemakerEdgePackagingJobName)
	}

	if resp, err := client.StopEdgePackagingJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a running hyperparameter tuning job and all running training jobs that
// the tuning job launched.
//
// All model artifacts output from the training jobs are stored in Amazon Simple
// Storage Service (Amazon S3). All data that the training jobs write to Amazon
// CloudWatch Logs are still available in CloudWatch. After the tuning job moves to
// the Stopped state, it releases all reserved resources for the tuning job.
func sagemaker_StopHyperParameterTuningJob(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.StopHyperParameterTuningJobInput{
		// HyperParameterTuningJobName: *string, // Required
	}

	if len(_sagemakerHyperParameterTuningJobName) > 0 {
		input.HyperParameterTuningJobName = aws.String(_sagemakerHyperParameterTuningJobName)
	}

	if resp, err := client.StopHyperParameterTuningJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops an inference experiment.
func sagemaker_StopInferenceExperiment(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.StopInferenceExperimentInput{
		// ModelVariantActions: map[string]types.ModelVariantAction, // Required
		// Name: *string, // Required
	}

	if len(_sagemakerModelVariantActions) > 0 {
		if err := assignInputField(input, "ModelVariantActions", _sagemakerModelVariantActions); err != nil {
			log.Errorf("invalid --model-variant-actions: %s", err.Error())
			return
		}
	}
	if len(_sagemakerName) > 0 {
		input.Name = aws.String(_sagemakerName)
	}
	if len(_sagemakerDesiredModelVariants) > 0 {
		if err := assignInputField(input, "DesiredModelVariants", _sagemakerDesiredModelVariants); err != nil {
			log.Errorf("invalid --desired-model-variants: %s", err.Error())
			return
		}
	}
	if len(_sagemakerDesiredState) > 0 {
		if err := assignInputField(input, "DesiredState", _sagemakerDesiredState); err != nil {
			log.Errorf("invalid --desired-state: %s", err.Error())
			return
		}
	}
	if len(_sagemakerReason) > 0 {
		input.Reason = aws.String(_sagemakerReason)
	}

	if resp, err := client.StopInferenceExperiment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops an Inference Recommender job.
func sagemaker_StopInferenceRecommendationsJob(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.StopInferenceRecommendationsJobInput{
		// JobName: *string, // Required
	}

	if len(_sagemakerJobName) > 0 {
		input.JobName = aws.String(_sagemakerJobName)
	}

	if resp, err := client.StopInferenceRecommendationsJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a running labeling job. A job that is stopped cannot be restarted. Any
// results obtained before the job is stopped are placed in the Amazon S3 output
// bucket.
func sagemaker_StopLabelingJob(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.StopLabelingJobInput{
		// LabelingJobName: *string, // Required
	}

	if len(_sagemakerLabelingJobName) > 0 {
		input.LabelingJobName = aws.String(_sagemakerLabelingJobName)
	}

	if resp, err := client.StopLabelingJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Programmatically stop an MLflow Tracking Server.
func sagemaker_StopMlflowTrackingServer(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.StopMlflowTrackingServerInput{
		// TrackingServerName: *string, // Required
	}

	if len(_sagemakerTrackingServerName) > 0 {
		input.TrackingServerName = aws.String(_sagemakerTrackingServerName)
	}

	if resp, err := client.StopMlflowTrackingServer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a previously started monitoring schedule.
func sagemaker_StopMonitoringSchedule(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.StopMonitoringScheduleInput{
		// MonitoringScheduleName: *string, // Required
	}

	if len(_sagemakerMonitoringScheduleName) > 0 {
		input.MonitoringScheduleName = aws.String(_sagemakerMonitoringScheduleName)
	}

	if resp, err := client.StopMonitoringSchedule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Terminates the ML compute instance. Before terminating the instance, SageMaker
// AI disconnects the ML storage volume from it. SageMaker AI preserves the ML
// storage volume. SageMaker AI stops charging you for the ML compute instance when
// you call StopNotebookInstance .
//
// To access data on the ML storage volume for a notebook instance that has been
// terminated, call the StartNotebookInstance API. StartNotebookInstance launches
// another ML compute instance, configures it, and attaches the preserved ML
// storage volume so you can continue your work.
func sagemaker_StopNotebookInstance(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.StopNotebookInstanceInput{
		// NotebookInstanceName: *string, // Required
	}

	if len(_sagemakerNotebookInstanceName) > 0 {
		input.NotebookInstanceName = aws.String(_sagemakerNotebookInstanceName)
	}

	if resp, err := client.StopNotebookInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Ends a running inference optimization job.
func sagemaker_StopOptimizationJob(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.StopOptimizationJobInput{
		// OptimizationJobName: *string, // Required
	}

	if len(_sagemakerOptimizationJobName) > 0 {
		input.OptimizationJobName = aws.String(_sagemakerOptimizationJobName)
	}

	if resp, err := client.StopOptimizationJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a pipeline execution.
// # Callback Step
//
// A pipeline execution won't stop while a callback step is running. When you call
// StopPipelineExecution on a pipeline execution with a running callback step,
// SageMaker Pipelines sends an additional Amazon SQS message to the specified SQS
// queue. The body of the SQS message contains a "Status" field which is set to
// "Stopping".
//
// You should add logic to your Amazon SQS message consumer to take any needed
// action (for example, resource cleanup) upon receipt of the message followed by a
// call to SendPipelineExecutionStepSuccess or SendPipelineExecutionStepFailure .
//
// Only when SageMaker Pipelines receives one of these calls will it stop the
// pipeline execution.
//
// # Lambda Step
//
// A pipeline execution can't be stopped while a lambda step is running because
// the Lambda function invoked by the lambda step can't be stopped. If you attempt
// to stop the execution while the Lambda function is running, the pipeline waits
// for the Lambda function to finish or until the timeout is hit, whichever occurs
// first, and then stops. If the Lambda function finishes, the pipeline execution
// status is Stopped . If the timeout is hit the pipeline execution status is
// Failed .
func sagemaker_StopPipelineExecution(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.StopPipelineExecutionInput{
		// ClientRequestToken: *string, // Required
		// PipelineExecutionArn: *string, // Required
	}

	if len(_sagemakerClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_sagemakerClientRequestToken)
	}
	if len(_sagemakerPipelineExecutionArn) > 0 {
		input.PipelineExecutionArn = aws.String(_sagemakerPipelineExecutionArn)
	}

	if resp, err := client.StopPipelineExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a processing job.
func sagemaker_StopProcessingJob(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.StopProcessingJobInput{
		// ProcessingJobName: *string, // Required
	}

	if len(_sagemakerProcessingJobName) > 0 {
		input.ProcessingJobName = aws.String(_sagemakerProcessingJobName)
	}

	if resp, err := client.StopProcessingJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a training job. To stop a job, SageMaker sends the algorithm the SIGTERM
// signal, which delays job termination for 120 seconds. Algorithms might use this
// 120-second window to save the model artifacts, so the results of the training is
// not lost.
//
// When it receives a StopTrainingJob request, SageMaker changes the status of the
// job to Stopping . After SageMaker stops the job, it sets the status to Stopped .
func sagemaker_StopTrainingJob(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.StopTrainingJobInput{
		// TrainingJobName: *string, // Required
	}

	if len(_sagemakerTrainingJobName) > 0 {
		input.TrainingJobName = aws.String(_sagemakerTrainingJobName)
	}

	if resp, err := client.StopTrainingJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a batch transform job.
// When Amazon SageMaker receives a StopTransformJob request, the status of the
// job changes to Stopping . After Amazon SageMaker stops the job, the status is
// set to Stopped . When you stop a batch transform job before it is completed,
// Amazon SageMaker doesn't store the job's output in Amazon S3.
func sagemaker_StopTransformJob(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.StopTransformJobInput{
		// TransformJobName: *string, // Required
	}

	if len(_sagemakerTransformJobName) > 0 {
		input.TransformJobName = aws.String(_sagemakerTransformJobName)
	}

	if resp, err := client.StopTransformJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an action.
func sagemaker_UpdateAction(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.UpdateActionInput{
		// ActionName: *string, // Required
	}

	if len(_sagemakerActionName) > 0 {
		input.ActionName = aws.String(_sagemakerActionName)
	}
	if len(_sagemakerDescription) > 0 {
		input.Description = aws.String(_sagemakerDescription)
	}
	if len(_sagemakerProperties) > 0 {
		if err := assignInputField(input, "Properties", _sagemakerProperties); err != nil {
			log.Errorf("invalid --properties: %s", err.Error())
			return
		}
	}
	if len(_sagemakerPropertiesToRemove) > 0 {
		input.PropertiesToRemove = append([]string(nil), _sagemakerPropertiesToRemove...)
	}
	if len(_sagemakerStatus) > 0 {
		if err := assignInputField(input, "Status", _sagemakerStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateAction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the properties of an AppImageConfig.
func sagemaker_UpdateAppImageConfig(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.UpdateAppImageConfigInput{
		// AppImageConfigName: *string, // Required
	}

	if len(_sagemakerAppImageConfigName) > 0 {
		input.AppImageConfigName = aws.String(_sagemakerAppImageConfigName)
	}
	if len(_sagemakerCodeEditorAppImageConfig) > 0 {
		if err := assignInputField(input, "CodeEditorAppImageConfig", _sagemakerCodeEditorAppImageConfig); err != nil {
			log.Errorf("invalid --code-editor-app-image-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerJupyterLabAppImageConfig) > 0 {
		if err := assignInputField(input, "JupyterLabAppImageConfig", _sagemakerJupyterLabAppImageConfig); err != nil {
			log.Errorf("invalid --jupyter-lab-app-image-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerKernelGatewayImageConfig) > 0 {
		if err := assignInputField(input, "KernelGatewayImageConfig", _sagemakerKernelGatewayImageConfig); err != nil {
			log.Errorf("invalid --kernel-gateway-image-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateAppImageConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an artifact.
func sagemaker_UpdateArtifact(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.UpdateArtifactInput{
		// ArtifactArn: *string, // Required
	}

	if len(_sagemakerArtifactArn) > 0 {
		input.ArtifactArn = aws.String(_sagemakerArtifactArn)
	}
	if len(_sagemakerArtifactName) > 0 {
		input.ArtifactName = aws.String(_sagemakerArtifactName)
	}
	if len(_sagemakerProperties) > 0 {
		if err := assignInputField(input, "Properties", _sagemakerProperties); err != nil {
			log.Errorf("invalid --properties: %s", err.Error())
			return
		}
	}
	if len(_sagemakerPropertiesToRemove) > 0 {
		input.PropertiesToRemove = append([]string(nil), _sagemakerPropertiesToRemove...)
	}

	if resp, err := client.UpdateArtifact(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a SageMaker HyperPod cluster.
func sagemaker_UpdateCluster(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.UpdateClusterInput{
		// ClusterName: *string, // Required
	}

	if len(_sagemakerClusterName) > 0 {
		input.ClusterName = aws.String(_sagemakerClusterName)
	}
	if len(_sagemakerAutoScaling) > 0 {
		if err := assignInputField(input, "AutoScaling", _sagemakerAutoScaling); err != nil {
			log.Errorf("invalid --auto-scaling: %s", err.Error())
			return
		}
	}
	if len(_sagemakerClusterRole) > 0 {
		input.ClusterRole = aws.String(_sagemakerClusterRole)
	}
	if len(_sagemakerInstanceGroups) > 0 {
		if err := assignInputField(input, "InstanceGroups", _sagemakerInstanceGroups); err != nil {
			log.Errorf("invalid --instance-groups: %s", err.Error())
			return
		}
	}
	if len(_sagemakerInstanceGroupsToDelete) > 0 {
		input.InstanceGroupsToDelete = append([]string(nil), _sagemakerInstanceGroupsToDelete...)
	}
	if len(_sagemakerNodeProvisioningMode) > 0 {
		if err := assignInputField(input, "NodeProvisioningMode", _sagemakerNodeProvisioningMode); err != nil {
			log.Errorf("invalid --node-provisioning-mode: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNodeRecovery) > 0 {
		if err := assignInputField(input, "NodeRecovery", _sagemakerNodeRecovery); err != nil {
			log.Errorf("invalid --node-recovery: %s", err.Error())
			return
		}
	}
	if len(_sagemakerOrchestrator) > 0 {
		if err := assignInputField(input, "Orchestrator", _sagemakerOrchestrator); err != nil {
			log.Errorf("invalid --orchestrator: %s", err.Error())
			return
		}
	}
	if len(_sagemakerRestrictedInstanceGroups) > 0 {
		if err := assignInputField(input, "RestrictedInstanceGroups", _sagemakerRestrictedInstanceGroups); err != nil {
			log.Errorf("invalid --restricted-instance-groups: %s", err.Error())
			return
		}
	}
	if len(_sagemakerTieredStorageConfig) > 0 {
		if err := assignInputField(input, "TieredStorageConfig", _sagemakerTieredStorageConfig); err != nil {
			log.Errorf("invalid --tiered-storage-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update the cluster policy configuration.
func sagemaker_UpdateClusterSchedulerConfig(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.UpdateClusterSchedulerConfigInput{
		// ClusterSchedulerConfigId: *string, // Required
		// TargetVersion: *int32, // Required
	}

	if len(_sagemakerClusterSchedulerConfigId) > 0 {
		input.ClusterSchedulerConfigId = aws.String(_sagemakerClusterSchedulerConfigId)
	}
	if len(_sagemakerTargetVersion) > 0 {
		if err := assignInputField(input, "TargetVersion", _sagemakerTargetVersion); err != nil {
			log.Errorf("invalid --target-version: %s", err.Error())
			return
		}
	}
	if len(_sagemakerDescription) > 0 {
		input.Description = aws.String(_sagemakerDescription)
	}
	if len(_sagemakerSchedulerConfig) > 0 {
		if err := assignInputField(input, "SchedulerConfig", _sagemakerSchedulerConfig); err != nil {
			log.Errorf("invalid --scheduler-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateClusterSchedulerConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the platform software of a SageMaker HyperPod cluster for security
// patching. To learn how to use this API, see [Update the SageMaker HyperPod platform software of a cluster].
//
// The UpgradeClusterSoftware API call may impact your SageMaker HyperPod cluster
// uptime and availability. Plan accordingly to mitigate potential disruptions to
// your workloads.
//
// [Update the SageMaker HyperPod platform software of a cluster]: https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-hyperpod-operate.html#sagemaker-hyperpod-operate-cli-command-update-cluster-software
func sagemaker_UpdateClusterSoftware(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.UpdateClusterSoftwareInput{
		// ClusterName: *string, // Required
	}

	if len(_sagemakerClusterName) > 0 {
		input.ClusterName = aws.String(_sagemakerClusterName)
	}
	if len(_sagemakerDeploymentConfig) > 0 {
		if err := assignInputField(input, "DeploymentConfig", _sagemakerDeploymentConfig); err != nil {
			log.Errorf("invalid --deployment-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerImageId) > 0 {
		input.ImageId = aws.String(_sagemakerImageId)
	}
	if len(_sagemakerInstanceGroups) > 0 {
		if err := assignInputField(input, "InstanceGroups", _sagemakerInstanceGroups); err != nil {
			log.Errorf("invalid --instance-groups: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateClusterSoftware(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified Git repository with the specified values.
func sagemaker_UpdateCodeRepository(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.UpdateCodeRepositoryInput{
		// CodeRepositoryName: *string, // Required
	}

	if len(_sagemakerCodeRepositoryName) > 0 {
		input.CodeRepositoryName = aws.String(_sagemakerCodeRepositoryName)
	}
	if len(_sagemakerGitConfig) > 0 {
		if err := assignInputField(input, "GitConfig", _sagemakerGitConfig); err != nil {
			log.Errorf("invalid --git-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateCodeRepository(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update the compute allocation definition.
func sagemaker_UpdateComputeQuota(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.UpdateComputeQuotaInput{
		// ComputeQuotaId: *string, // Required
		// TargetVersion: *int32, // Required
	}

	if len(_sagemakerComputeQuotaId) > 0 {
		input.ComputeQuotaId = aws.String(_sagemakerComputeQuotaId)
	}
	if len(_sagemakerTargetVersion) > 0 {
		if err := assignInputField(input, "TargetVersion", _sagemakerTargetVersion); err != nil {
			log.Errorf("invalid --target-version: %s", err.Error())
			return
		}
	}
	if len(_sagemakerActivationState) > 0 {
		if err := assignInputField(input, "ActivationState", _sagemakerActivationState); err != nil {
			log.Errorf("invalid --activation-state: %s", err.Error())
			return
		}
	}
	if len(_sagemakerComputeQuotaConfig) > 0 {
		if err := assignInputField(input, "ComputeQuotaConfig", _sagemakerComputeQuotaConfig); err != nil {
			log.Errorf("invalid --compute-quota-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerComputeQuotaTarget) > 0 {
		if err := assignInputField(input, "ComputeQuotaTarget", _sagemakerComputeQuotaTarget); err != nil {
			log.Errorf("invalid --compute-quota-target: %s", err.Error())
			return
		}
	}
	if len(_sagemakerDescription) > 0 {
		input.Description = aws.String(_sagemakerDescription)
	}

	if resp, err := client.UpdateComputeQuota(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a context.
func sagemaker_UpdateContext(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.UpdateContextInput{
		// ContextName: *string, // Required
	}

	if len(_sagemakerContextName) > 0 {
		input.ContextName = aws.String(_sagemakerContextName)
	}
	if len(_sagemakerDescription) > 0 {
		input.Description = aws.String(_sagemakerDescription)
	}
	if len(_sagemakerProperties) > 0 {
		if err := assignInputField(input, "Properties", _sagemakerProperties); err != nil {
			log.Errorf("invalid --properties: %s", err.Error())
			return
		}
	}
	if len(_sagemakerPropertiesToRemove) > 0 {
		input.PropertiesToRemove = append([]string(nil), _sagemakerPropertiesToRemove...)
	}

	if resp, err := client.UpdateContext(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a fleet of devices.
func sagemaker_UpdateDeviceFleet(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.UpdateDeviceFleetInput{
		// DeviceFleetName: *string, // Required
		// OutputConfig: *types.EdgeOutputConfig, // Required
	}

	if len(_sagemakerDeviceFleetName) > 0 {
		input.DeviceFleetName = aws.String(_sagemakerDeviceFleetName)
	}
	if len(_sagemakerOutputConfig) > 0 {
		if err := assignInputField(input, "OutputConfig", _sagemakerOutputConfig); err != nil {
			log.Errorf("invalid --output-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerDescription) > 0 {
		input.Description = aws.String(_sagemakerDescription)
	}
	if len(_sagemakerEnableIotRoleAlias) > 0 {
		if err := assignInputField(input, "EnableIotRoleAlias", _sagemakerEnableIotRoleAlias); err != nil {
			log.Errorf("invalid --enable-iot-role-alias: %s", err.Error())
			return
		}
	}
	if len(_sagemakerRoleArn) > 0 {
		input.RoleArn = aws.String(_sagemakerRoleArn)
	}

	if resp, err := client.UpdateDeviceFleet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates one or more devices in a fleet.
func sagemaker_UpdateDevices(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.UpdateDevicesInput{
		// DeviceFleetName: *string, // Required
		// Devices: []types.Device, // Required
	}

	if len(_sagemakerDeviceFleetName) > 0 {
		input.DeviceFleetName = aws.String(_sagemakerDeviceFleetName)
	}
	if len(_sagemakerDevices) > 0 {
		if err := assignInputField(input, "Devices", _sagemakerDevices); err != nil {
			log.Errorf("invalid --devices: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateDevices(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the default settings for new user profiles in the domain.
func sagemaker_UpdateDomain(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.UpdateDomainInput{
		// DomainId: *string, // Required
	}

	if len(_sagemakerDomainId) > 0 {
		input.DomainId = aws.String(_sagemakerDomainId)
	}
	if len(_sagemakerAppNetworkAccessType) > 0 {
		if err := assignInputField(input, "AppNetworkAccessType", _sagemakerAppNetworkAccessType); err != nil {
			log.Errorf("invalid --app-network-access-type: %s", err.Error())
			return
		}
	}
	if len(_sagemakerAppSecurityGroupManagement) > 0 {
		if err := assignInputField(input, "AppSecurityGroupManagement", _sagemakerAppSecurityGroupManagement); err != nil {
			log.Errorf("invalid --app-security-group-management: %s", err.Error())
			return
		}
	}
	if len(_sagemakerDefaultSpaceSettings) > 0 {
		if err := assignInputField(input, "DefaultSpaceSettings", _sagemakerDefaultSpaceSettings); err != nil {
			log.Errorf("invalid --default-space-settings: %s", err.Error())
			return
		}
	}
	if len(_sagemakerDefaultUserSettings) > 0 {
		if err := assignInputField(input, "DefaultUserSettings", _sagemakerDefaultUserSettings); err != nil {
			log.Errorf("invalid --default-user-settings: %s", err.Error())
			return
		}
	}
	if len(_sagemakerDomainSettingsForUpdate) > 0 {
		if err := assignInputField(input, "DomainSettingsForUpdate", _sagemakerDomainSettingsForUpdate); err != nil {
			log.Errorf("invalid --domain-settings-for-update: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSubnetIds) > 0 {
		input.SubnetIds = append([]string(nil), _sagemakerSubnetIds...)
	}
	if len(_sagemakerTagPropagation) > 0 {
		if err := assignInputField(input, "TagPropagation", _sagemakerTagPropagation); err != nil {
			log.Errorf("invalid --tag-propagation: %s", err.Error())
			return
		}
	}
	if len(_sagemakerVpcId) > 0 {
		input.VpcId = aws.String(_sagemakerVpcId)
	}

	if resp, err := client.UpdateDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deploys the EndpointConfig specified in the request to a new fleet of
// instances. SageMaker shifts endpoint traffic to the new instances with the
// updated endpoint configuration and then deletes the old instances using the
// previous EndpointConfig (there is no availability loss). For more information
// about how to control the update and traffic shifting process, see [Update models in production].
//
// When SageMaker receives the request, it sets the endpoint status to Updating .
// After updating the endpoint, it sets the status to InService . To check the
// status of an endpoint, use the [DescribeEndpoint]API.
//
// You must not delete an EndpointConfig in use by an endpoint that is live or
// while the UpdateEndpoint or CreateEndpoint operations are being performed on
// the endpoint. To update an endpoint, you must create a new EndpointConfig .
//
// If you delete the EndpointConfig of an endpoint that is active or being created
// or updated you may lose visibility into the instance type the endpoint is using.
// The endpoint must be deleted in order to stop incurring charges.
//
// [DescribeEndpoint]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DescribeEndpoint.html
// [Update models in production]: https://docs.aws.amazon.com/sagemaker/latest/dg/deployment-guardrails.html
func sagemaker_UpdateEndpoint(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.UpdateEndpointInput{
		// EndpointConfigName: *string, // Required
		// EndpointName: *string, // Required
	}

	if len(_sagemakerEndpointConfigName) > 0 {
		input.EndpointConfigName = aws.String(_sagemakerEndpointConfigName)
	}
	if len(_sagemakerEndpointName) > 0 {
		input.EndpointName = aws.String(_sagemakerEndpointName)
	}
	if len(_sagemakerDeploymentConfig) > 0 {
		if err := assignInputField(input, "DeploymentConfig", _sagemakerDeploymentConfig); err != nil {
			log.Errorf("invalid --deployment-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerExcludeRetainedVariantProperties) > 0 {
		if err := assignInputField(input, "ExcludeRetainedVariantProperties", _sagemakerExcludeRetainedVariantProperties); err != nil {
			log.Errorf("invalid --exclude-retained-variant-properties: %s", err.Error())
			return
		}
	}
	if len(_sagemakerRetainAllVariantProperties) > 0 {
		if err := assignInputField(input, "RetainAllVariantProperties", _sagemakerRetainAllVariantProperties); err != nil {
			log.Errorf("invalid --retain-all-variant-properties: %s", err.Error())
			return
		}
	}
	if len(_sagemakerRetainDeploymentConfig) > 0 {
		if err := assignInputField(input, "RetainDeploymentConfig", _sagemakerRetainDeploymentConfig); err != nil {
			log.Errorf("invalid --retain-deployment-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates variant weight of one or more variants associated with an existing
// endpoint, or capacity of one variant associated with an existing endpoint. When
// it receives the request, SageMaker sets the endpoint status to Updating . After
// updating the endpoint, it sets the status to InService . To check the status of
// an endpoint, use the [DescribeEndpoint]API.
//
// [DescribeEndpoint]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DescribeEndpoint.html
func sagemaker_UpdateEndpointWeightsAndCapacities(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.UpdateEndpointWeightsAndCapacitiesInput{
		// DesiredWeightsAndCapacities: []types.DesiredWeightAndCapacity, // Required
		// EndpointName: *string, // Required
	}

	if len(_sagemakerDesiredWeightsAndCapacities) > 0 {
		if err := assignInputField(input, "DesiredWeightsAndCapacities", _sagemakerDesiredWeightsAndCapacities); err != nil {
			log.Errorf("invalid --desired-weights-and-capacities: %s", err.Error())
			return
		}
	}
	if len(_sagemakerEndpointName) > 0 {
		input.EndpointName = aws.String(_sagemakerEndpointName)
	}

	if resp, err := client.UpdateEndpointWeightsAndCapacities(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds, updates, or removes the description of an experiment. Updates the display
// name of an experiment.
func sagemaker_UpdateExperiment(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.UpdateExperimentInput{
		// ExperimentName: *string, // Required
	}

	if len(_sagemakerExperimentName) > 0 {
		input.ExperimentName = aws.String(_sagemakerExperimentName)
	}
	if len(_sagemakerDescription) > 0 {
		input.Description = aws.String(_sagemakerDescription)
	}
	if len(_sagemakerDisplayName) > 0 {
		input.DisplayName = aws.String(_sagemakerDisplayName)
	}

	if resp, err := client.UpdateExperiment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the feature group by either adding features or updating the online
// store configuration. Use one of the following request parameters at a time while
// using the UpdateFeatureGroup API.
//
// You can add features for your feature group using the FeatureAdditions request
// parameter. Features cannot be removed from a feature group.
//
// You can update the online store configuration by using the OnlineStoreConfig
// request parameter. If a TtlDuration is specified, the default TtlDuration
// applies for all records added to the feature group after the feature group is
// updated. If a record level TtlDuration exists from using the PutRecord API, the
// record level TtlDuration applies to that record instead of the default
// TtlDuration . To remove the default TtlDuration from an existing feature group,
// use the UpdateFeatureGroup API and set the TtlDuration Unit and Value to null .
func sagemaker_UpdateFeatureGroup(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.UpdateFeatureGroupInput{
		// FeatureGroupName: *string, // Required
	}

	if len(_sagemakerFeatureGroupName) > 0 {
		input.FeatureGroupName = aws.String(_sagemakerFeatureGroupName)
	}
	if len(_sagemakerFeatureAdditions) > 0 {
		if err := assignInputField(input, "FeatureAdditions", _sagemakerFeatureAdditions); err != nil {
			log.Errorf("invalid --feature-additions: %s", err.Error())
			return
		}
	}
	if len(_sagemakerOnlineStoreConfig) > 0 {
		if err := assignInputField(input, "OnlineStoreConfig", _sagemakerOnlineStoreConfig); err != nil {
			log.Errorf("invalid --online-store-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerThroughputConfig) > 0 {
		if err := assignInputField(input, "ThroughputConfig", _sagemakerThroughputConfig); err != nil {
			log.Errorf("invalid --throughput-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateFeatureGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the description and parameters of the feature group.
func sagemaker_UpdateFeatureMetadata(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.UpdateFeatureMetadataInput{
		// FeatureGroupName: *string, // Required
		// FeatureName: *string, // Required
	}

	if len(_sagemakerFeatureGroupName) > 0 {
		input.FeatureGroupName = aws.String(_sagemakerFeatureGroupName)
	}
	if len(_sagemakerFeatureName) > 0 {
		input.FeatureName = aws.String(_sagemakerFeatureName)
	}
	if len(_sagemakerDescription) > 0 {
		input.Description = aws.String(_sagemakerDescription)
	}
	if len(_sagemakerParameterAdditions) > 0 {
		if err := assignInputField(input, "ParameterAdditions", _sagemakerParameterAdditions); err != nil {
			log.Errorf("invalid --parameter-additions: %s", err.Error())
			return
		}
	}
	if len(_sagemakerParameterRemovals) > 0 {
		input.ParameterRemovals = append([]string(nil), _sagemakerParameterRemovals...)
	}

	if resp, err := client.UpdateFeatureMetadata(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update a hub.
func sagemaker_UpdateHub(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.UpdateHubInput{
		// HubName: *string, // Required
	}

	if len(_sagemakerHubName) > 0 {
		input.HubName = aws.String(_sagemakerHubName)
	}
	if len(_sagemakerHubDescription) > 0 {
		input.HubDescription = aws.String(_sagemakerHubDescription)
	}
	if len(_sagemakerHubDisplayName) > 0 {
		input.HubDisplayName = aws.String(_sagemakerHubDisplayName)
	}
	if len(_sagemakerHubSearchKeywords) > 0 {
		input.HubSearchKeywords = append([]string(nil), _sagemakerHubSearchKeywords...)
	}

	if resp, err := client.UpdateHub(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates SageMaker hub content (either a Model or Notebook resource).
// You can update the metadata that describes the resource. In addition to the
// required request fields, specify at least one of the following fields to update:
//
// - HubContentDescription
//
// - HubContentDisplayName
//
// - HubContentMarkdown
//
// - HubContentSearchKeywords
//
// - SupportStatus
//
// For more information about hubs, see [Private curated hubs for foundation model access control in JumpStart].
//
// If you want to update a ModelReference resource in your hub, use the
// UpdateHubContentResource API instead.
//
// [Private curated hubs for foundation model access control in JumpStart]: https://docs.aws.amazon.com/sagemaker/latest/dg/jumpstart-curated-hubs.html
func sagemaker_UpdateHubContent(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.UpdateHubContentInput{
		// HubContentName: *string, // Required
		// HubContentType: types.HubContentType, // Required
		// HubContentVersion: *string, // Required
		// HubName: *string, // Required
	}

	if len(_sagemakerHubContentName) > 0 {
		input.HubContentName = aws.String(_sagemakerHubContentName)
	}
	if len(_sagemakerHubContentType) > 0 {
		if err := assignInputField(input, "HubContentType", _sagemakerHubContentType); err != nil {
			log.Errorf("invalid --hub-content-type: %s", err.Error())
			return
		}
	}
	if len(_sagemakerHubContentVersion) > 0 {
		input.HubContentVersion = aws.String(_sagemakerHubContentVersion)
	}
	if len(_sagemakerHubName) > 0 {
		input.HubName = aws.String(_sagemakerHubName)
	}
	if len(_sagemakerHubContentDescription) > 0 {
		input.HubContentDescription = aws.String(_sagemakerHubContentDescription)
	}
	if len(_sagemakerHubContentDisplayName) > 0 {
		input.HubContentDisplayName = aws.String(_sagemakerHubContentDisplayName)
	}
	if len(_sagemakerHubContentMarkdown) > 0 {
		input.HubContentMarkdown = aws.String(_sagemakerHubContentMarkdown)
	}
	if len(_sagemakerHubContentSearchKeywords) > 0 {
		input.HubContentSearchKeywords = append([]string(nil), _sagemakerHubContentSearchKeywords...)
	}
	if len(_sagemakerSupportStatus) > 0 {
		if err := assignInputField(input, "SupportStatus", _sagemakerSupportStatus); err != nil {
			log.Errorf("invalid --support-status: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateHubContent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the contents of a SageMaker hub for a ModelReference resource. A
// ModelReference allows you to access public SageMaker JumpStart models from
// within your private hub.
//
// When using this API, you can update the MinVersion field for additional
// flexibility in the model version. You shouldn't update any additional fields
// when using this API, because the metadata in your private hub should match the
// public JumpStart model's metadata.
//
// If you want to update a Model or Notebook resource in your hub, use the
// UpdateHubContent API instead.
//
// For more information about adding model references to your hub, see [Add models to a private hub].
//
// [Add models to a private hub]: https://docs.aws.amazon.com/sagemaker/latest/dg/jumpstart-curated-hubs-admin-guide-add-models.html
func sagemaker_UpdateHubContentReference(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.UpdateHubContentReferenceInput{
		// HubContentName: *string, // Required
		// HubContentType: types.HubContentType, // Required
		// HubName: *string, // Required
	}

	if len(_sagemakerHubContentName) > 0 {
		input.HubContentName = aws.String(_sagemakerHubContentName)
	}
	if len(_sagemakerHubContentType) > 0 {
		if err := assignInputField(input, "HubContentType", _sagemakerHubContentType); err != nil {
			log.Errorf("invalid --hub-content-type: %s", err.Error())
			return
		}
	}
	if len(_sagemakerHubName) > 0 {
		input.HubName = aws.String(_sagemakerHubName)
	}
	if len(_sagemakerMinVersion) > 0 {
		input.MinVersion = aws.String(_sagemakerMinVersion)
	}

	if resp, err := client.UpdateHubContentReference(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the properties of a SageMaker AI image. To change the image's tags, use
// the [AddTags]and [DeleteTags] APIs.
//
// [AddTags]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_AddTags.html
// [DeleteTags]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DeleteTags.html
func sagemaker_UpdateImage(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.UpdateImageInput{
		// ImageName: *string, // Required
	}

	if len(_sagemakerImageName) > 0 {
		input.ImageName = aws.String(_sagemakerImageName)
	}
	if len(_sagemakerDeleteProperties) > 0 {
		input.DeleteProperties = append([]string(nil), _sagemakerDeleteProperties...)
	}
	if len(_sagemakerDescription) > 0 {
		input.Description = aws.String(_sagemakerDescription)
	}
	if len(_sagemakerDisplayName) > 0 {
		input.DisplayName = aws.String(_sagemakerDisplayName)
	}
	if len(_sagemakerRoleArn) > 0 {
		input.RoleArn = aws.String(_sagemakerRoleArn)
	}

	if resp, err := client.UpdateImage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the properties of a SageMaker AI image version.
func sagemaker_UpdateImageVersion(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.UpdateImageVersionInput{
		// ImageName: *string, // Required
	}

	if len(_sagemakerImageName) > 0 {
		input.ImageName = aws.String(_sagemakerImageName)
	}
	if len(_sagemakerAlias) > 0 {
		input.Alias = aws.String(_sagemakerAlias)
	}
	if len(_sagemakerAliasesToAdd) > 0 {
		input.AliasesToAdd = append([]string(nil), _sagemakerAliasesToAdd...)
	}
	if len(_sagemakerAliasesToDelete) > 0 {
		input.AliasesToDelete = append([]string(nil), _sagemakerAliasesToDelete...)
	}
	if len(_sagemakerHorovod) > 0 {
		if err := assignInputField(input, "Horovod", _sagemakerHorovod); err != nil {
			log.Errorf("invalid --horovod: %s", err.Error())
			return
		}
	}
	if len(_sagemakerJobType) > 0 {
		if err := assignInputField(input, "JobType", _sagemakerJobType); err != nil {
			log.Errorf("invalid --job-type: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMLFramework) > 0 {
		input.MLFramework = aws.String(_sagemakerMLFramework)
	}
	if len(_sagemakerProcessor) > 0 {
		if err := assignInputField(input, "Processor", _sagemakerProcessor); err != nil {
			log.Errorf("invalid --processor: %s", err.Error())
			return
		}
	}
	if len(_sagemakerProgrammingLang) > 0 {
		input.ProgrammingLang = aws.String(_sagemakerProgrammingLang)
	}
	if len(_sagemakerReleaseNotes) > 0 {
		input.ReleaseNotes = aws.String(_sagemakerReleaseNotes)
	}
	if len(_sagemakerVendorGuidance) > 0 {
		if err := assignInputField(input, "VendorGuidance", _sagemakerVendorGuidance); err != nil {
			log.Errorf("invalid --vendor-guidance: %s", err.Error())
			return
		}
	}
	if len(_sagemakerVersion) > 0 {
		if err := assignInputField(input, "Version", _sagemakerVersion); err != nil {
			log.Errorf("invalid --version: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateImageVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an inference component.
func sagemaker_UpdateInferenceComponent(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.UpdateInferenceComponentInput{
		// InferenceComponentName: *string, // Required
	}

	if len(_sagemakerInferenceComponentName) > 0 {
		input.InferenceComponentName = aws.String(_sagemakerInferenceComponentName)
	}
	if len(_sagemakerDeploymentConfig) > 0 {
		if err := assignInputField(input, "DeploymentConfig", _sagemakerDeploymentConfig); err != nil {
			log.Errorf("invalid --deployment-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerRuntimeConfig) > 0 {
		if err := assignInputField(input, "RuntimeConfig", _sagemakerRuntimeConfig); err != nil {
			log.Errorf("invalid --runtime-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSpecification) > 0 {
		if err := assignInputField(input, "Specification", _sagemakerSpecification); err != nil {
			log.Errorf("invalid --specification: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateInferenceComponent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Runtime settings for a model that is deployed with an inference component.
func sagemaker_UpdateInferenceComponentRuntimeConfig(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.UpdateInferenceComponentRuntimeConfigInput{
		// DesiredRuntimeConfig: *types.InferenceComponentRuntimeConfig, // Required
		// InferenceComponentName: *string, // Required
	}

	if len(_sagemakerDesiredRuntimeConfig) > 0 {
		if err := assignInputField(input, "DesiredRuntimeConfig", _sagemakerDesiredRuntimeConfig); err != nil {
			log.Errorf("invalid --desired-runtime-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerInferenceComponentName) > 0 {
		input.InferenceComponentName = aws.String(_sagemakerInferenceComponentName)
	}

	if resp, err := client.UpdateInferenceComponentRuntimeConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an inference experiment that you created. The status of the inference
// experiment has to be either Created , Running . For more information on the
// status of an inference experiment, see [DescribeInferenceExperiment].
//
// [DescribeInferenceExperiment]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DescribeInferenceExperiment.html
func sagemaker_UpdateInferenceExperiment(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.UpdateInferenceExperimentInput{
		// Name: *string, // Required
	}

	if len(_sagemakerName) > 0 {
		input.Name = aws.String(_sagemakerName)
	}
	if len(_sagemakerDataStorageConfig) > 0 {
		if err := assignInputField(input, "DataStorageConfig", _sagemakerDataStorageConfig); err != nil {
			log.Errorf("invalid --data-storage-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerDescription) > 0 {
		input.Description = aws.String(_sagemakerDescription)
	}
	if len(_sagemakerModelVariants) > 0 {
		if err := assignInputField(input, "ModelVariants", _sagemakerModelVariants); err != nil {
			log.Errorf("invalid --model-variants: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSchedule) > 0 {
		if err := assignInputField(input, "Schedule", _sagemakerSchedule); err != nil {
			log.Errorf("invalid --schedule: %s", err.Error())
			return
		}
	}
	if len(_sagemakerShadowModeConfig) > 0 {
		if err := assignInputField(input, "ShadowModeConfig", _sagemakerShadowModeConfig); err != nil {
			log.Errorf("invalid --shadow-mode-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateInferenceExperiment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an MLflow App.
func sagemaker_UpdateMlflowApp(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.UpdateMlflowAppInput{
		// Arn: *string, // Required
	}

	if len(_sagemakerArn) > 0 {
		input.Arn = aws.String(_sagemakerArn)
	}
	if len(_sagemakerAccountDefaultStatus) > 0 {
		if err := assignInputField(input, "AccountDefaultStatus", _sagemakerAccountDefaultStatus); err != nil {
			log.Errorf("invalid --account-default-status: %s", err.Error())
			return
		}
	}
	if len(_sagemakerArtifactStoreUri) > 0 {
		input.ArtifactStoreUri = aws.String(_sagemakerArtifactStoreUri)
	}
	if len(_sagemakerDefaultDomainIdList) > 0 {
		input.DefaultDomainIdList = append([]string(nil), _sagemakerDefaultDomainIdList...)
	}
	if len(_sagemakerModelRegistrationMode) > 0 {
		if err := assignInputField(input, "ModelRegistrationMode", _sagemakerModelRegistrationMode); err != nil {
			log.Errorf("invalid --model-registration-mode: %s", err.Error())
			return
		}
	}
	if len(_sagemakerName) > 0 {
		input.Name = aws.String(_sagemakerName)
	}
	if len(_sagemakerWeeklyMaintenanceWindowStart) > 0 {
		input.WeeklyMaintenanceWindowStart = aws.String(_sagemakerWeeklyMaintenanceWindowStart)
	}

	if resp, err := client.UpdateMlflowApp(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates properties of an existing MLflow Tracking Server.
func sagemaker_UpdateMlflowTrackingServer(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.UpdateMlflowTrackingServerInput{
		// TrackingServerName: *string, // Required
	}

	if len(_sagemakerTrackingServerName) > 0 {
		input.TrackingServerName = aws.String(_sagemakerTrackingServerName)
	}
	if len(_sagemakerArtifactStoreUri) > 0 {
		input.ArtifactStoreUri = aws.String(_sagemakerArtifactStoreUri)
	}
	if len(_sagemakerAutomaticModelRegistration) > 0 {
		if err := assignInputField(input, "AutomaticModelRegistration", _sagemakerAutomaticModelRegistration); err != nil {
			log.Errorf("invalid --automatic-model-registration: %s", err.Error())
			return
		}
	}
	if len(_sagemakerTrackingServerSize) > 0 {
		if err := assignInputField(input, "TrackingServerSize", _sagemakerTrackingServerSize); err != nil {
			log.Errorf("invalid --tracking-server-size: %s", err.Error())
			return
		}
	}
	if len(_sagemakerWeeklyMaintenanceWindowStart) > 0 {
		input.WeeklyMaintenanceWindowStart = aws.String(_sagemakerWeeklyMaintenanceWindowStart)
	}

	if resp, err := client.UpdateMlflowTrackingServer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update an Amazon SageMaker Model Card.
// You cannot update both model card content and model card status in a single
// call.
func sagemaker_UpdateModelCard(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.UpdateModelCardInput{
		// ModelCardName: *string, // Required
	}

	if len(_sagemakerModelCardName) > 0 {
		input.ModelCardName = aws.String(_sagemakerModelCardName)
	}
	if len(_sagemakerContent) > 0 {
		input.Content = aws.String(_sagemakerContent)
	}
	if len(_sagemakerModelCardStatus) > 0 {
		if err := assignInputField(input, "ModelCardStatus", _sagemakerModelCardStatus); err != nil {
			log.Errorf("invalid --model-card-status: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateModelCard(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a versioned model.
func sagemaker_UpdateModelPackage(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.UpdateModelPackageInput{
		// ModelPackageArn: *string, // Required
	}

	if len(_sagemakerModelPackageArn) > 0 {
		input.ModelPackageArn = aws.String(_sagemakerModelPackageArn)
	}
	if len(_sagemakerAdditionalInferenceSpecificationsToAdd) > 0 {
		if err := assignInputField(input, "AdditionalInferenceSpecificationsToAdd", _sagemakerAdditionalInferenceSpecificationsToAdd); err != nil {
			log.Errorf("invalid --additional-inference-specifications-to-add: %s", err.Error())
			return
		}
	}
	if len(_sagemakerApprovalDescription) > 0 {
		input.ApprovalDescription = aws.String(_sagemakerApprovalDescription)
	}
	if len(_sagemakerClientToken) > 0 {
		input.ClientToken = aws.String(_sagemakerClientToken)
	}
	if len(_sagemakerCustomerMetadataProperties) > 0 {
		if err := assignInputField(input, "CustomerMetadataProperties", _sagemakerCustomerMetadataProperties); err != nil {
			log.Errorf("invalid --customer-metadata-properties: %s", err.Error())
			return
		}
	}
	if len(_sagemakerCustomerMetadataPropertiesToRemove) > 0 {
		input.CustomerMetadataPropertiesToRemove = append([]string(nil), _sagemakerCustomerMetadataPropertiesToRemove...)
	}
	if len(_sagemakerInferenceSpecification) > 0 {
		if err := assignInputField(input, "InferenceSpecification", _sagemakerInferenceSpecification); err != nil {
			log.Errorf("invalid --inference-specification: %s", err.Error())
			return
		}
	}
	if len(_sagemakerModelApprovalStatus) > 0 {
		if err := assignInputField(input, "ModelApprovalStatus", _sagemakerModelApprovalStatus); err != nil {
			log.Errorf("invalid --model-approval-status: %s", err.Error())
			return
		}
	}
	if len(_sagemakerModelCard) > 0 {
		if err := assignInputField(input, "ModelCard", _sagemakerModelCard); err != nil {
			log.Errorf("invalid --model-card: %s", err.Error())
			return
		}
	}
	if len(_sagemakerModelLifeCycle) > 0 {
		if err := assignInputField(input, "ModelLifeCycle", _sagemakerModelLifeCycle); err != nil {
			log.Errorf("invalid --model-life-cycle: %s", err.Error())
			return
		}
	}
	if len(_sagemakerModelPackageRegistrationType) > 0 {
		if err := assignInputField(input, "ModelPackageRegistrationType", _sagemakerModelPackageRegistrationType); err != nil {
			log.Errorf("invalid --model-package-registration-type: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSourceUri) > 0 {
		input.SourceUri = aws.String(_sagemakerSourceUri)
	}

	if resp, err := client.UpdateModelPackage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update the parameters of a model monitor alert.
func sagemaker_UpdateMonitoringAlert(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.UpdateMonitoringAlertInput{
		// DatapointsToAlert: *int32, // Required
		// EvaluationPeriod: *int32, // Required
		// MonitoringAlertName: *string, // Required
		// MonitoringScheduleName: *string, // Required
	}

	if len(_sagemakerDatapointsToAlert) > 0 {
		if err := assignInputField(input, "DatapointsToAlert", _sagemakerDatapointsToAlert); err != nil {
			log.Errorf("invalid --datapoints-to-alert: %s", err.Error())
			return
		}
	}
	if len(_sagemakerEvaluationPeriod) > 0 {
		if err := assignInputField(input, "EvaluationPeriod", _sagemakerEvaluationPeriod); err != nil {
			log.Errorf("invalid --evaluation-period: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMonitoringAlertName) > 0 {
		input.MonitoringAlertName = aws.String(_sagemakerMonitoringAlertName)
	}
	if len(_sagemakerMonitoringScheduleName) > 0 {
		input.MonitoringScheduleName = aws.String(_sagemakerMonitoringScheduleName)
	}

	if resp, err := client.UpdateMonitoringAlert(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a previously created schedule.
func sagemaker_UpdateMonitoringSchedule(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.UpdateMonitoringScheduleInput{
		// MonitoringScheduleConfig: *types.MonitoringScheduleConfig, // Required
		// MonitoringScheduleName: *string, // Required
	}

	if len(_sagemakerMonitoringScheduleConfig) > 0 {
		if err := assignInputField(input, "MonitoringScheduleConfig", _sagemakerMonitoringScheduleConfig); err != nil {
			log.Errorf("invalid --monitoring-schedule-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMonitoringScheduleName) > 0 {
		input.MonitoringScheduleName = aws.String(_sagemakerMonitoringScheduleName)
	}

	if resp, err := client.UpdateMonitoringSchedule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a notebook instance. NotebookInstance updates include upgrading or
// downgrading the ML compute instance used for your notebook instance to
// accommodate changes in your workload requirements.
//
// This API can attach lifecycle configurations to notebook instances. Lifecycle
// configuration scripts execute with root access and the notebook instance's IAM
// execution role privileges. Principals with this permission and access to
// lifecycle configurations can execute code with the execution role's credentials.
// See [Customize a Notebook Instance Using a Lifecycle Configuration Script]for security best practices.
//
// [Customize a Notebook Instance Using a Lifecycle Configuration Script]: https://docs.aws.amazon.com/sagemaker/latest/dg/notebook-lifecycle-config.html
func sagemaker_UpdateNotebookInstance(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.UpdateNotebookInstanceInput{
		// NotebookInstanceName: *string, // Required
	}

	if len(_sagemakerNotebookInstanceName) > 0 {
		input.NotebookInstanceName = aws.String(_sagemakerNotebookInstanceName)
	}
	if len(_sagemakerAcceleratorTypes) > 0 {
		if err := assignInputField(input, "AcceleratorTypes", _sagemakerAcceleratorTypes); err != nil {
			log.Errorf("invalid --accelerator-types: %s", err.Error())
			return
		}
	}
	if len(_sagemakerAdditionalCodeRepositories) > 0 {
		input.AdditionalCodeRepositories = append([]string(nil), _sagemakerAdditionalCodeRepositories...)
	}
	if len(_sagemakerDefaultCodeRepository) > 0 {
		input.DefaultCodeRepository = aws.String(_sagemakerDefaultCodeRepository)
	}
	if len(_sagemakerDisassociateAcceleratorTypes) > 0 {
		if err := assignInputField(input, "DisassociateAcceleratorTypes", _sagemakerDisassociateAcceleratorTypes); err != nil {
			log.Errorf("invalid --disassociate-accelerator-types: %s", err.Error())
			return
		}
	}
	if len(_sagemakerDisassociateAdditionalCodeRepositories) > 0 {
		if err := assignInputField(input, "DisassociateAdditionalCodeRepositories", _sagemakerDisassociateAdditionalCodeRepositories); err != nil {
			log.Errorf("invalid --disassociate-additional-code-repositories: %s", err.Error())
			return
		}
	}
	if len(_sagemakerDisassociateDefaultCodeRepository) > 0 {
		if err := assignInputField(input, "DisassociateDefaultCodeRepository", _sagemakerDisassociateDefaultCodeRepository); err != nil {
			log.Errorf("invalid --disassociate-default-code-repository: %s", err.Error())
			return
		}
	}
	if len(_sagemakerDisassociateLifecycleConfig) > 0 {
		if err := assignInputField(input, "DisassociateLifecycleConfig", _sagemakerDisassociateLifecycleConfig); err != nil {
			log.Errorf("invalid --disassociate-lifecycle-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerInstanceMetadataServiceConfiguration) > 0 {
		if err := assignInputField(input, "InstanceMetadataServiceConfiguration", _sagemakerInstanceMetadataServiceConfiguration); err != nil {
			log.Errorf("invalid --instance-metadata-service-configuration: %s", err.Error())
			return
		}
	}
	if len(_sagemakerInstanceType) > 0 {
		if err := assignInputField(input, "InstanceType", _sagemakerInstanceType); err != nil {
			log.Errorf("invalid --instance-type: %s", err.Error())
			return
		}
	}
	if len(_sagemakerIpAddressType) > 0 {
		if err := assignInputField(input, "IpAddressType", _sagemakerIpAddressType); err != nil {
			log.Errorf("invalid --ip-address-type: %s", err.Error())
			return
		}
	}
	if len(_sagemakerLifecycleConfigName) > 0 {
		input.LifecycleConfigName = aws.String(_sagemakerLifecycleConfigName)
	}
	if len(_sagemakerPlatformIdentifier) > 0 {
		input.PlatformIdentifier = aws.String(_sagemakerPlatformIdentifier)
	}
	if len(_sagemakerRoleArn) > 0 {
		input.RoleArn = aws.String(_sagemakerRoleArn)
	}
	if len(_sagemakerRootAccess) > 0 {
		if err := assignInputField(input, "RootAccess", _sagemakerRootAccess); err != nil {
			log.Errorf("invalid --root-access: %s", err.Error())
			return
		}
	}
	if len(_sagemakerVolumeSizeInGB) > 0 {
		if err := assignInputField(input, "VolumeSizeInGB", _sagemakerVolumeSizeInGB); err != nil {
			log.Errorf("invalid --volume-size-in-gb: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateNotebookInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a notebook instance lifecycle configuration created with the [CreateNotebookInstanceLifecycleConfig] API.
// Updates to lifecycle configurations affect all notebook instances using that
// configuration upon their next start. Lifecycle configuration scripts execute
// with root access and the notebook instance's IAM execution role privileges.
// Grant this permission only to trusted principals. See [Customize a Notebook Instance Using a Lifecycle Configuration Script]for security best
// practices.
//
// [Customize a Notebook Instance Using a Lifecycle Configuration Script]: https://docs.aws.amazon.com/sagemaker/latest/dg/notebook-lifecycle-config.html
// [CreateNotebookInstanceLifecycleConfig]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateNotebookInstanceLifecycleConfig.html
func sagemaker_UpdateNotebookInstanceLifecycleConfig(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.UpdateNotebookInstanceLifecycleConfigInput{
		// NotebookInstanceLifecycleConfigName: *string, // Required
	}

	if len(_sagemakerNotebookInstanceLifecycleConfigName) > 0 {
		input.NotebookInstanceLifecycleConfigName = aws.String(_sagemakerNotebookInstanceLifecycleConfigName)
	}
	if len(_sagemakerOnCreate) > 0 {
		if err := assignInputField(input, "OnCreate", _sagemakerOnCreate); err != nil {
			log.Errorf("invalid --on-create: %s", err.Error())
			return
		}
	}
	if len(_sagemakerOnStart) > 0 {
		if err := assignInputField(input, "OnStart", _sagemakerOnStart); err != nil {
			log.Errorf("invalid --on-start: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateNotebookInstanceLifecycleConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates all of the SageMaker Partner AI Apps in an account.
func sagemaker_UpdatePartnerApp(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.UpdatePartnerAppInput{
		// Arn: *string, // Required
	}

	if len(_sagemakerArn) > 0 {
		input.Arn = aws.String(_sagemakerArn)
	}
	if len(_sagemakerAppVersion) > 0 {
		input.AppVersion = aws.String(_sagemakerAppVersion)
	}
	if len(_sagemakerApplicationConfig) > 0 {
		if err := assignInputField(input, "ApplicationConfig", _sagemakerApplicationConfig); err != nil {
			log.Errorf("invalid --application-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerClientToken) > 0 {
		input.ClientToken = aws.String(_sagemakerClientToken)
	}
	if len(_sagemakerEnableAutoMinorVersionUpgrade) > 0 {
		if err := assignInputField(input, "EnableAutoMinorVersionUpgrade", _sagemakerEnableAutoMinorVersionUpgrade); err != nil {
			log.Errorf("invalid --enable-auto-minor-version-upgrade: %s", err.Error())
			return
		}
	}
	if len(_sagemakerEnableIamSessionBasedIdentity) > 0 {
		if err := assignInputField(input, "EnableIamSessionBasedIdentity", _sagemakerEnableIamSessionBasedIdentity); err != nil {
			log.Errorf("invalid --enable-iam-session-based-identity: %s", err.Error())
			return
		}
	}
	if len(_sagemakerMaintenanceConfig) > 0 {
		if err := assignInputField(input, "MaintenanceConfig", _sagemakerMaintenanceConfig); err != nil {
			log.Errorf("invalid --maintenance-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_sagemakerTier) > 0 {
		input.Tier = aws.String(_sagemakerTier)
	}

	if resp, err := client.UpdatePartnerApp(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a pipeline.
func sagemaker_UpdatePipeline(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.UpdatePipelineInput{
		// PipelineName: *string, // Required
	}

	if len(_sagemakerPipelineName) > 0 {
		input.PipelineName = aws.String(_sagemakerPipelineName)
	}
	if len(_sagemakerParallelismConfiguration) > 0 {
		if err := assignInputField(input, "ParallelismConfiguration", _sagemakerParallelismConfiguration); err != nil {
			log.Errorf("invalid --parallelism-configuration: %s", err.Error())
			return
		}
	}
	if len(_sagemakerPipelineDefinition) > 0 {
		input.PipelineDefinition = aws.String(_sagemakerPipelineDefinition)
	}
	if len(_sagemakerPipelineDefinitionS3Location) > 0 {
		if err := assignInputField(input, "PipelineDefinitionS3Location", _sagemakerPipelineDefinitionS3Location); err != nil {
			log.Errorf("invalid --pipeline-definition-s3-location: %s", err.Error())
			return
		}
	}
	if len(_sagemakerPipelineDescription) > 0 {
		input.PipelineDescription = aws.String(_sagemakerPipelineDescription)
	}
	if len(_sagemakerPipelineDisplayName) > 0 {
		input.PipelineDisplayName = aws.String(_sagemakerPipelineDisplayName)
	}
	if len(_sagemakerRoleArn) > 0 {
		input.RoleArn = aws.String(_sagemakerRoleArn)
	}

	if resp, err := client.UpdatePipeline(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a pipeline execution.
func sagemaker_UpdatePipelineExecution(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.UpdatePipelineExecutionInput{
		// PipelineExecutionArn: *string, // Required
	}

	if len(_sagemakerPipelineExecutionArn) > 0 {
		input.PipelineExecutionArn = aws.String(_sagemakerPipelineExecutionArn)
	}
	if len(_sagemakerParallelismConfiguration) > 0 {
		if err := assignInputField(input, "ParallelismConfiguration", _sagemakerParallelismConfiguration); err != nil {
			log.Errorf("invalid --parallelism-configuration: %s", err.Error())
			return
		}
	}
	if len(_sagemakerPipelineExecutionDescription) > 0 {
		input.PipelineExecutionDescription = aws.String(_sagemakerPipelineExecutionDescription)
	}
	if len(_sagemakerPipelineExecutionDisplayName) > 0 {
		input.PipelineExecutionDisplayName = aws.String(_sagemakerPipelineExecutionDisplayName)
	}

	if resp, err := client.UpdatePipelineExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a pipeline version.
func sagemaker_UpdatePipelineVersion(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.UpdatePipelineVersionInput{
		// PipelineArn: *string, // Required
		// PipelineVersionId: *int64, // Required
	}

	if len(_sagemakerPipelineArn) > 0 {
		input.PipelineArn = aws.String(_sagemakerPipelineArn)
	}
	if len(_sagemakerPipelineVersionId) > 0 {
		if err := assignInputField(input, "PipelineVersionId", _sagemakerPipelineVersionId); err != nil {
			log.Errorf("invalid --pipeline-version-id: %s", err.Error())
			return
		}
	}
	if len(_sagemakerPipelineVersionDescription) > 0 {
		input.PipelineVersionDescription = aws.String(_sagemakerPipelineVersionDescription)
	}
	if len(_sagemakerPipelineVersionDisplayName) > 0 {
		input.PipelineVersionDisplayName = aws.String(_sagemakerPipelineVersionDisplayName)
	}

	if resp, err := client.UpdatePipelineVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a machine learning (ML) project that is created from a template that
// sets up an ML pipeline from training to deploying an approved model.
//
// You must not update a project that is in use. If you update the
// ServiceCatalogProvisioningUpdateDetails of a project that is active or being
// created, or updated, you may lose resources already created by the project.
func sagemaker_UpdateProject(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.UpdateProjectInput{
		// ProjectName: *string, // Required
	}

	if len(_sagemakerProjectName) > 0 {
		input.ProjectName = aws.String(_sagemakerProjectName)
	}
	if len(_sagemakerProjectDescription) > 0 {
		input.ProjectDescription = aws.String(_sagemakerProjectDescription)
	}
	if len(_sagemakerServiceCatalogProvisioningUpdateDetails) > 0 {
		if err := assignInputField(input, "ServiceCatalogProvisioningUpdateDetails", _sagemakerServiceCatalogProvisioningUpdateDetails); err != nil {
			log.Errorf("invalid --service-catalog-provisioning-update-details: %s", err.Error())
			return
		}
	}
	if len(_sagemakerTags) > 0 {
		if err := assignInputField(input, "Tags", _sagemakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_sagemakerTemplateProvidersToUpdate) > 0 {
		if err := assignInputField(input, "TemplateProvidersToUpdate", _sagemakerTemplateProvidersToUpdate); err != nil {
			log.Errorf("invalid --template-providers-to-update: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateProject(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the settings of a space.
// You can't edit the app type of a space in the SpaceSettings .
func sagemaker_UpdateSpace(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.UpdateSpaceInput{
		// DomainId: *string, // Required
		// SpaceName: *string, // Required
	}

	if len(_sagemakerDomainId) > 0 {
		input.DomainId = aws.String(_sagemakerDomainId)
	}
	if len(_sagemakerSpaceName) > 0 {
		input.SpaceName = aws.String(_sagemakerSpaceName)
	}
	if len(_sagemakerSpaceDisplayName) > 0 {
		input.SpaceDisplayName = aws.String(_sagemakerSpaceDisplayName)
	}
	if len(_sagemakerSpaceSettings) > 0 {
		if err := assignInputField(input, "SpaceSettings", _sagemakerSpaceSettings); err != nil {
			log.Errorf("invalid --space-settings: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateSpace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update a model training job to request a new Debugger profiling configuration
// or to change warm pool retention length.
func sagemaker_UpdateTrainingJob(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.UpdateTrainingJobInput{
		// TrainingJobName: *string, // Required
	}

	if len(_sagemakerTrainingJobName) > 0 {
		input.TrainingJobName = aws.String(_sagemakerTrainingJobName)
	}
	if len(_sagemakerProfilerConfig) > 0 {
		if err := assignInputField(input, "ProfilerConfig", _sagemakerProfilerConfig); err != nil {
			log.Errorf("invalid --profiler-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerProfilerRuleConfigurations) > 0 {
		if err := assignInputField(input, "ProfilerRuleConfigurations", _sagemakerProfilerRuleConfigurations); err != nil {
			log.Errorf("invalid --profiler-rule-configurations: %s", err.Error())
			return
		}
	}
	if len(_sagemakerRemoteDebugConfig) > 0 {
		if err := assignInputField(input, "RemoteDebugConfig", _sagemakerRemoteDebugConfig); err != nil {
			log.Errorf("invalid --remote-debug-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerResourceConfig) > 0 {
		if err := assignInputField(input, "ResourceConfig", _sagemakerResourceConfig); err != nil {
			log.Errorf("invalid --resource-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateTrainingJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the display name of a trial.
func sagemaker_UpdateTrial(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.UpdateTrialInput{
		// TrialName: *string, // Required
	}

	if len(_sagemakerTrialName) > 0 {
		input.TrialName = aws.String(_sagemakerTrialName)
	}
	if len(_sagemakerDisplayName) > 0 {
		input.DisplayName = aws.String(_sagemakerDisplayName)
	}

	if resp, err := client.UpdateTrial(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates one or more properties of a trial component.
func sagemaker_UpdateTrialComponent(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.UpdateTrialComponentInput{
		// TrialComponentName: *string, // Required
	}

	if len(_sagemakerTrialComponentName) > 0 {
		input.TrialComponentName = aws.String(_sagemakerTrialComponentName)
	}
	if len(_sagemakerDisplayName) > 0 {
		input.DisplayName = aws.String(_sagemakerDisplayName)
	}
	if len(_sagemakerEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _sagemakerEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_sagemakerInputArtifacts) > 0 {
		if err := assignInputField(input, "InputArtifacts", _sagemakerInputArtifacts); err != nil {
			log.Errorf("invalid --input-artifacts: %s", err.Error())
			return
		}
	}
	if len(_sagemakerInputArtifactsToRemove) > 0 {
		input.InputArtifactsToRemove = append([]string(nil), _sagemakerInputArtifactsToRemove...)
	}
	if len(_sagemakerOutputArtifacts) > 0 {
		if err := assignInputField(input, "OutputArtifacts", _sagemakerOutputArtifacts); err != nil {
			log.Errorf("invalid --output-artifacts: %s", err.Error())
			return
		}
	}
	if len(_sagemakerOutputArtifactsToRemove) > 0 {
		input.OutputArtifactsToRemove = append([]string(nil), _sagemakerOutputArtifactsToRemove...)
	}
	if len(_sagemakerParameters) > 0 {
		if err := assignInputField(input, "Parameters", _sagemakerParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}
	if len(_sagemakerParametersToRemove) > 0 {
		input.ParametersToRemove = append([]string(nil), _sagemakerParametersToRemove...)
	}
	if len(_sagemakerStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _sagemakerStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_sagemakerStatus) > 0 {
		if err := assignInputField(input, "Status", _sagemakerStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateTrialComponent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a user profile.
func sagemaker_UpdateUserProfile(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.UpdateUserProfileInput{
		// DomainId: *string, // Required
		// UserProfileName: *string, // Required
	}

	if len(_sagemakerDomainId) > 0 {
		input.DomainId = aws.String(_sagemakerDomainId)
	}
	if len(_sagemakerUserProfileName) > 0 {
		input.UserProfileName = aws.String(_sagemakerUserProfileName)
	}
	if len(_sagemakerUserSettings) > 0 {
		if err := assignInputField(input, "UserSettings", _sagemakerUserSettings); err != nil {
			log.Errorf("invalid --user-settings: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateUserProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Use this operation to update your workforce. You can use this operation to
// require that workers use specific IP addresses to work on tasks and to update
// your OpenID Connect (OIDC) Identity Provider (IdP) workforce configuration.
//
// The worker portal is now supported in VPC and public internet.
//
// Use SourceIpConfig to restrict worker access to tasks to a specific range of IP
// addresses. You specify allowed IP addresses by creating a list of up to ten [CIDRs].
// By default, a workforce isn't restricted to specific IP addresses. If you
// specify a range of IP addresses, workers who attempt to access tasks using any
// IP address outside the specified range are denied and get a Not Found error
// message on the worker portal.
//
// To restrict public internet access for all workers, configure the SourceIpConfig
// CIDR value. For example, when using SourceIpConfig with an IpAddressType of IPv4
// , you can restrict access to the IPv4 CIDR block "10.0.0.0/16". When using an
// IpAddressType of dualstack , you can specify both the IPv4 and IPv6 CIDR blocks,
// such as "10.0.0.0/16" for IPv4 only, "2001:db8:1234:1a00::/56" for IPv6 only, or
// "10.0.0.0/16" and "2001:db8:1234:1a00::/56" for dual stack.
//
// Amazon SageMaker does not support Source Ip restriction for worker portals in
// VPC.
//
// Use OidcConfig to update the configuration of a workforce created using your
// own OIDC IdP.
//
// You can only update your OIDC IdP configuration when there are no work teams
// associated with your workforce. You can delete work teams using the [DeleteWorkteam]operation.
//
// After restricting access to a range of IP addresses or updating your OIDC IdP
// configuration with this operation, you can view details about your update
// workforce using the [DescribeWorkforce]operation.
//
// This operation only applies to private workforces.
//
// [DescribeWorkforce]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DescribeWorkforce.html
// [DeleteWorkteam]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DeleteWorkteam.html
// [CIDRs]: https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Subnets.html
func sagemaker_UpdateWorkforce(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.UpdateWorkforceInput{
		// WorkforceName: *string, // Required
	}

	if len(_sagemakerWorkforceName) > 0 {
		input.WorkforceName = aws.String(_sagemakerWorkforceName)
	}
	if len(_sagemakerIpAddressType) > 0 {
		if err := assignInputField(input, "IpAddressType", _sagemakerIpAddressType); err != nil {
			log.Errorf("invalid --ip-address-type: %s", err.Error())
			return
		}
	}
	if len(_sagemakerOidcConfig) > 0 {
		if err := assignInputField(input, "OidcConfig", _sagemakerOidcConfig); err != nil {
			log.Errorf("invalid --oidc-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerSourceIpConfig) > 0 {
		if err := assignInputField(input, "SourceIpConfig", _sagemakerSourceIpConfig); err != nil {
			log.Errorf("invalid --source-ip-config: %s", err.Error())
			return
		}
	}
	if len(_sagemakerWorkforceVpcConfig) > 0 {
		if err := assignInputField(input, "WorkforceVpcConfig", _sagemakerWorkforceVpcConfig); err != nil {
			log.Errorf("invalid --workforce-vpc-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateWorkforce(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing work team with new member definitions or description.
func sagemaker_UpdateWorkteam(cfg aws.Config, client *sagemaker.Client) {
	input := &sagemaker.UpdateWorkteamInput{
		// WorkteamName: *string, // Required
	}

	if len(_sagemakerWorkteamName) > 0 {
		input.WorkteamName = aws.String(_sagemakerWorkteamName)
	}
	if len(_sagemakerDescription) > 0 {
		input.Description = aws.String(_sagemakerDescription)
	}
	if len(_sagemakerMemberDefinitions) > 0 {
		if err := assignInputField(input, "MemberDefinitions", _sagemakerMemberDefinitions); err != nil {
			log.Errorf("invalid --member-definitions: %s", err.Error())
			return
		}
	}
	if len(_sagemakerNotificationConfiguration) > 0 {
		if err := assignInputField(input, "NotificationConfiguration", _sagemakerNotificationConfiguration); err != nil {
			log.Errorf("invalid --notification-configuration: %s", err.Error())
			return
		}
	}
	if len(_sagemakerWorkerAccessConfiguration) > 0 {
		if err := assignInputField(input, "WorkerAccessConfiguration", _sagemakerWorkerAccessConfiguration); err != nil {
			log.Errorf("invalid --worker-access-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateWorkteam(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_sagemakerCmd)
	_sagemakerCmd.Flags().SortFlags = false

	_sagemakerCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_sagemakerCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_sagemakerCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_sagemakerCmd.Flags().StringVarP(&_sagemakerAcceleratorTypes, "accelerator-types", "", "", "Accelerator Types")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerAccessConfig, "access-config", "", "", "Access Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerAccountDefaultStatus, "account-default-status", "", "", "Account Default Status")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerActionName, "action-name", "", "", "Action Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerActionType, "action-type", "", "", "Action Type")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerActivationState, "activation-state", "", "", "Activation State")
	_sagemakerCmd.Flags().StringSliceVarP(&_sagemakerAdditionalCodeRepositories, "additional-code-repositories", "", nil, "Additional Code Repositories")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerAdditionalCodeRepositoryEquals, "additional-code-repository-equals", "", "", "Additional Code Repository Equals")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerAdditionalInferenceSpecifications, "additional-inference-specifications", "", "", "Additional Inference Specifications")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerAdditionalInferenceSpecificationsToAdd, "additional-inference-specifications-to-add", "", "", "Additional Inference Specifications To Add")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerAlgorithmDescription, "algorithm-description", "", "", "Algorithm Description")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerAlgorithmName, "algorithm-name", "", "", "Algorithm Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerAlgorithmSpecification, "algorithm-specification", "", "", "Algorithm Specification")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerAlias, "alias", "", "", "Alias")
	_sagemakerCmd.Flags().StringSliceVarP(&_sagemakerAliases, "aliases", "", nil, "Aliases")
	_sagemakerCmd.Flags().StringSliceVarP(&_sagemakerAliasesToAdd, "aliases-to-add", "", nil, "Aliases To Add")
	_sagemakerCmd.Flags().StringSliceVarP(&_sagemakerAliasesToDelete, "aliases-to-delete", "", nil, "Aliases To Delete")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerAppImageConfigName, "app-image-config-name", "", "", "App Image Config Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerAppName, "app-name", "", "", "App Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerAppNetworkAccessType, "app-network-access-type", "", "", "App Network Access Type")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerAppSecurityGroupManagement, "app-security-group-management", "", "", "App Security Group Management")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerAppSpecification, "app-specification", "", "", "App Specification")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerAppType, "app-type", "", "", "App Type")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerAppTypeEquals, "app-type-equals", "", "", "App Type Equals")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerAppVersion, "app-version", "", "", "App Version")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerApplicationConfig, "application-config", "", "", "Application Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerApprovalDescription, "approval-description", "", "", "Approval Description")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerArn, "arn", "", "", "ARN")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerArtifactArn, "artifact-arn", "", "", "Artifact ARN")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerArtifactName, "artifact-name", "", "", "Artifact Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerArtifactStoreUri, "artifact-store-uri", "", "", "Artifact Store URI")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerArtifactType, "artifact-type", "", "", "Artifact Type")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerAssociationType, "association-type", "", "", "Association Type")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerAsyncInferenceConfig, "async-inference-config", "", "", "Async Inference Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerAuthMode, "auth-mode", "", "", "Auth Mode")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerAuthType, "auth-type", "", "", "Auth Type")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerAutoMLComputeConfig, "auto-ml-compute-config", "", "", "Auto Ml Compute Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerAutoMLJobConfig, "auto-ml-job-config", "", "", "Auto Ml Job Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerAutoMLJobInputDataConfig, "auto-ml-job-input-data-config", "", "", "Auto Ml Job Input Data Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerAutoMLJobName, "auto-ml-job-name", "", "", "Auto Ml Job Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerAutoMLJobObjective, "auto-ml-job-objective", "", "", "Auto Ml Job Objective")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerAutoMLProblemTypeConfig, "auto-ml-problem-type-config", "", "", "Auto Ml Problem Type Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerAutoScaling, "auto-scaling", "", "", "Auto Scaling")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerAutomaticModelRegistration, "automatic-model-registration", "", "", "Automatic Model Registration")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerAutotune, "autotune", "", "", "Autotune")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerBaseImage, "base-image", "", "", "Base Image")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerBatchStrategy, "batch-strategy", "", "", "Batch Strategy")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerCallbackToken, "callback-token", "", "", "Callback Token")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerCandidateNameEquals, "candidate-name-equals", "", "", "Candidate Name Equals")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerCertifyForMarketplace, "certify-for-marketplace", "", "", "Certify For Marketplace")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerCheckpointConfig, "checkpoint-config", "", "", "Checkpoint Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerClientRequestToken, "client-request-token", "", "", "Client Request Token")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerClientToken, "client-token", "", "", "Client Token")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerClusterArn, "cluster-arn", "", "", "Cluster ARN")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerClusterName, "cluster-name", "", "", "Cluster Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerClusterRole, "cluster-role", "", "", "Cluster Role")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerClusterSchedulerConfigId, "cluster-scheduler-config-id", "", "", "Cluster Scheduler Config ID")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerClusterSchedulerConfigVersion, "cluster-scheduler-config-version", "", "", "Cluster Scheduler Config Version")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerCodeEditorAppImageConfig, "code-editor-app-image-config", "", "", "Code Editor App Image Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerCodeRepositoryName, "code-repository-name", "", "", "Code Repository Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerCognitoConfig, "cognito-config", "", "", "Cognito Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerCompilationJobName, "compilation-job-name", "", "", "Compilation Job Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerComputeQuotaConfig, "compute-quota-config", "", "", "Compute Quota Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerComputeQuotaId, "compute-quota-id", "", "", "Compute Quota ID")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerComputeQuotaTarget, "compute-quota-target", "", "", "Compute Quota Target")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerComputeQuotaVersion, "compute-quota-version", "", "", "Compute Quota Version")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerContainers, "containers", "", "", "Containers")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerContent, "content", "", "", "Content")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerContextName, "context-name", "", "", "Context Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerContextType, "context-type", "", "", "Context Type")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerCreatedAfter, "created-after", "", "", "Created After")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerCreatedBefore, "created-before", "", "", "Created Before")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerCreationTimeAfter, "creation-time-after", "", "", "Creation Time After")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerCreationTimeBefore, "creation-time-before", "", "", "Creation Time Before")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerCrossAccountFilterOption, "cross-account-filter-option", "", "", "Cross Account Filter Option")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerCustomerMetadataProperties, "customer-metadata-properties", "", "", "Customer Metadata Properties")
	_sagemakerCmd.Flags().StringSliceVarP(&_sagemakerCustomerMetadataPropertiesToRemove, "customer-metadata-properties-to-remove", "", nil, "Customer Metadata Properties To Remove")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerDataCaptureConfig, "data-capture-config", "", "", "Data Capture Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerDataProcessing, "data-processing", "", "", "Data Processing")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerDataQualityAppSpecification, "data-quality-app-specification", "", "", "Data Quality App Specification")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerDataQualityBaselineConfig, "data-quality-baseline-config", "", "", "Data Quality Baseline Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerDataQualityJobInput, "data-quality-job-input", "", "", "Data Quality Job Input")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerDataQualityJobOutputConfig, "data-quality-job-output-config", "", "", "Data Quality Job Output Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerDataSplitConfig, "data-split-config", "", "", "Data Split Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerDataStorageConfig, "data-storage-config", "", "", "Data Storage Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerDatapointsToAlert, "datapoints-to-alert", "", "", "Datapoints To Alert")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerDebugHookConfig, "debug-hook-config", "", "", "Debug Hook Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerDebugRuleConfigurations, "debug-rule-configurations", "", "", "Debug Rule Configurations")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerDefaultCodeRepository, "default-code-repository", "", "", "Default Code Repository")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerDefaultCodeRepositoryContains, "default-code-repository-contains", "", "", "Default Code Repository Contains")
	_sagemakerCmd.Flags().StringSliceVarP(&_sagemakerDefaultDomainIdList, "default-domain-id-list", "", nil, "Default Domain ID List")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerDefaultForDomainId, "default-for-domain-id", "", "", "Default For Domain ID")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerDefaultSpaceSettings, "default-space-settings", "", "", "Default Space Settings")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerDefaultUserSettings, "default-user-settings", "", "", "Default User Settings")
	_sagemakerCmd.Flags().StringSliceVarP(&_sagemakerDeleteProperties, "delete-properties", "", nil, "Delete Properties")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerDeploymentConfig, "deployment-config", "", "", "Deployment Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerDeploymentInstanceType, "deployment-instance-type", "", "", "Deployment Instance Type")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerDescription, "description", "", "", "Description")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerDesiredModelVariants, "desired-model-variants", "", "", "Desired Model Variants")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerDesiredRuntimeConfig, "desired-runtime-config", "", "", "Desired Runtime Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerDesiredState, "desired-state", "", "", "Desired State")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerDesiredWeightsAndCapacities, "desired-weights-and-capacities", "", "", "Desired Weights And Capacities")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerDestinationArn, "destination-arn", "", "", "Destination ARN")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerDestinationType, "destination-type", "", "", "Destination Type")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerDeviceFleetName, "device-fleet-name", "", "", "Device Fleet Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerDeviceFleetNameContains, "device-fleet-name-contains", "", "", "Device Fleet Name Contains")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerDeviceName, "device-name", "", "", "Device Name")
	_sagemakerCmd.Flags().StringSliceVarP(&_sagemakerDeviceNames, "device-names", "", nil, "Device Names")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerDevices, "devices", "", "", "Devices")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerDirectInternetAccess, "direct-internet-access", "", "", "Direct Internet Access")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerDirection, "direction", "", "", "Direction")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerDisassociateAcceleratorTypes, "disassociate-accelerator-types", "", "", "Disassociate Accelerator Types")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerDisassociateAdditionalCodeRepositories, "disassociate-additional-code-repositories", "", "", "Disassociate Additional Code Repositories")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerDisassociateDefaultCodeRepository, "disassociate-default-code-repository", "", "", "Disassociate Default Code Repository")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerDisassociateLifecycleConfig, "disassociate-lifecycle-config", "", "", "Disassociate Lifecycle Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerDisplayName, "display-name", "", "", "Display Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerDocumentSchemaVersion, "document-schema-version", "", "", "Document Schema Version")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerDomain, "domain", "", "", "Domain")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerDomainId, "domain-id", "", "", "Domain ID")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerDomainIdEquals, "domain-id-equals", "", "", "Domain ID Equals")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerDomainName, "domain-name", "", "", "Domain Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerDomainSettings, "domain-settings", "", "", "Domain Settings")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerDomainSettingsForUpdate, "domain-settings-for-update", "", "", "Domain Settings For Update")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerDriftCheckBaselines, "drift-check-baselines", "", "", "Drift Check Baselines")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerDurationHours, "duration-hours", "", "", "Duration Hours")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerEdgeDeploymentPlanName, "edge-deployment-plan-name", "", "", "Edge Deployment Plan Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerEdgePackagingJobName, "edge-packaging-job-name", "", "", "Edge Packaging Job Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerEnableAutoMinorVersionUpgrade, "enable-auto-minor-version-upgrade", "", "", "Enable Auto Minor Version Upgrade")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerEnableIamSessionBasedIdentity, "enable-iam-session-based-identity", "", "", "Enable IAM Session Based Identity")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerEnableInterContainerTrafficEncryption, "enable-inter-container-traffic-encryption", "", "", "Enable Inter Container Traffic Encryption")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerEnableIotRoleAlias, "enable-iot-role-alias", "", "", "Enable Iot Role Alias")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerEnableManagedSpotTraining, "enable-managed-spot-training", "", "", "Enable Managed Spot Training")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerEnableNetworkIsolation, "enable-network-isolation", "", "", "Enable Network Isolation")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerEndTime, "end-time", "", "", "End Time")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerEndTimeBefore, "end-time-before", "", "", "End Time Before")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerEndpointConfigName, "endpoint-config-name", "", "", "Endpoint Config Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerEndpointName, "endpoint-name", "", "", "Endpoint Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerEndpointNameEquals, "endpoint-name-equals", "", "", "Endpoint Name Equals")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerEnvironment, "environment", "", "", "Environment")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerEvaluationPeriod, "evaluation-period", "", "", "Evaluation Period")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerEventId, "event-id", "", "", "Event ID")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerEventTimeAfter, "event-time-after", "", "", "Event Time After")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerEventTimeBefore, "event-time-before", "", "", "Event Time Before")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerEventTimeFeatureName, "event-time-feature-name", "", "", "Event Time Feature Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerExcludeDevicesDeployedInOtherStage, "exclude-devices-deployed-in-other-stage", "", "", "Exclude Devices Deployed In Other Stage")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerExcludeRetainedVariantProperties, "exclude-retained-variant-properties", "", "", "Exclude Retained Variant Properties")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerExecutionRoleArn, "execution-role-arn", "", "", "Execution Role ARN")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerExperimentConfig, "experiment-config", "", "", "Experiment Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerExperimentName, "experiment-name", "", "", "Experiment Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerExpiresInSeconds, "expires-in-seconds", "", "", "Expires In Seconds")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerExplainerConfig, "explainer-config", "", "", "Explainer Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerFailureReason, "failure-reason", "", "", "Failure Reason")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerFeatureAdditions, "feature-additions", "", "", "Feature Additions")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerFeatureDefinitions, "feature-definitions", "", "", "Feature Definitions")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerFeatureGroupName, "feature-group-name", "", "", "Feature Group Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerFeatureGroupStatusEquals, "feature-group-status-equals", "", "", "Feature Group Status Equals")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerFeatureName, "feature-name", "", "", "Feature Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerFilters, "filters", "", "", "Filters")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerFlowDefinitionName, "flow-definition-name", "", "", "Flow Definition Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerGenerateCandidateDefinitionsOnly, "generate-candidate-definitions-only", "", "", "Generate Candidate Definitions Only")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerGitConfig, "git-config", "", "", "Git Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerHomeEfsFileSystemKmsKeyId, "home-efs-file-system-kms-key-id", "", "", "Home EFS File System KMS Key ID")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerHorovod, "horovod", "", "", "Horovod")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerHubContentDescription, "hub-content-description", "", "", "Hub Content Description")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerHubContentDisplayName, "hub-content-display-name", "", "", "Hub Content Display Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerHubContentDocument, "hub-content-document", "", "", "Hub Content Document")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerHubContentMarkdown, "hub-content-markdown", "", "", "Hub Content Markdown")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerHubContentName, "hub-content-name", "", "", "Hub Content Name")
	_sagemakerCmd.Flags().StringSliceVarP(&_sagemakerHubContentSearchKeywords, "hub-content-search-keywords", "", nil, "Hub Content Search Keywords")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerHubContentType, "hub-content-type", "", "", "Hub Content Type")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerHubContentVersion, "hub-content-version", "", "", "Hub Content Version")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerHubDescription, "hub-description", "", "", "Hub Description")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerHubDisplayName, "hub-display-name", "", "", "Hub Display Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerHubName, "hub-name", "", "", "Hub Name")
	_sagemakerCmd.Flags().StringSliceVarP(&_sagemakerHubSearchKeywords, "hub-search-keywords", "", nil, "Hub Search Keywords")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerHumanLoopActivationConfig, "human-loop-activation-config", "", "", "Human Loop Activation Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerHumanLoopConfig, "human-loop-config", "", "", "Human Loop Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerHumanLoopRequestSource, "human-loop-request-source", "", "", "Human Loop Request Source")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerHumanTaskConfig, "human-task-config", "", "", "Human Task Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerHumanTaskUiArn, "human-task-ui-arn", "", "", "Human Task Ui ARN")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerHumanTaskUiName, "human-task-ui-name", "", "", "Human Task Ui Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerHyperParameterTuningJobConfig, "hyper-parameter-tuning-job-config", "", "", "Hyper Parameter Tuning Job Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerHyperParameterTuningJobName, "hyper-parameter-tuning-job-name", "", "", "Hyper Parameter Tuning Job Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerHyperParameters, "hyper-parameters", "", "", "Hyper Parameters")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerImageId, "image-id", "", "", "Image ID")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerImageName, "image-name", "", "", "Image Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerIncludeAvailableUpgrade, "include-available-upgrade", "", "", "Include Available Upgrade")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerIncludeEdges, "include-edges", "", "", "Include Edges")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerIncludeNodeLogicalIds, "include-node-logical-ids", "", "", "Include Node Logical Ids")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerInferenceComponentName, "inference-component-name", "", "", "Inference Component Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerInferenceExecutionConfig, "inference-execution-config", "", "", "Inference Execution Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerInferenceRecommendationsJobName, "inference-recommendations-job-name", "", "", "Inference Recommendations Job Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerInferenceSpecification, "inference-specification", "", "", "Inference Specification")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerInfraCheckConfig, "infra-check-config", "", "", "Infra Check Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerInputArtifacts, "input-artifacts", "", "", "Input Artifacts")
	_sagemakerCmd.Flags().StringSliceVarP(&_sagemakerInputArtifactsToRemove, "input-artifacts-to-remove", "", nil, "Input Artifacts To Remove")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerInputConfig, "input-config", "", "", "Input Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerInputDataConfig, "input-data-config", "", "", "Input Data Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerInstanceCount, "instance-count", "", "", "Instance Count")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerInstanceGroupName, "instance-group-name", "", "", "Instance Group Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerInstanceGroupNameContains, "instance-group-name-contains", "", "", "Instance Group Name Contains")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerInstanceGroups, "instance-groups", "", "", "Instance Groups")
	_sagemakerCmd.Flags().StringSliceVarP(&_sagemakerInstanceGroupsToDelete, "instance-groups-to-delete", "", nil, "Instance Groups To Delete")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerInstanceMetadataServiceConfiguration, "instance-metadata-service-configuration", "", "", "Instance Metadata Service Configuration")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerInstanceType, "instance-type", "", "", "Instance Type")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerIpAddressType, "ip-address-type", "", "", "IP Address Type")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerJobDefinitionName, "job-definition-name", "", "", "Job Definition Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerJobDescription, "job-description", "", "", "Job Description")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerJobName, "job-name", "", "", "Job Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerJobReferenceCodeContains, "job-reference-code-contains", "", "", "Job Reference Code Contains")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerJobResources, "job-resources", "", "", "Job Resources")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerJobType, "job-type", "", "", "Job Type")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerJupyterLabAppImageConfig, "jupyter-lab-app-image-config", "", "", "Jupyter Lab App Image Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerKernelGatewayImageConfig, "kernel-gateway-image-config", "", "", "Kernel Gateway Image Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerKmsKey, "kms-key", "", "", "KMS Key")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerKmsKeyId, "kms-key-id", "", "", "KMS Key ID")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerLabelAttributeName, "label-attribute-name", "", "", "Label Attribute Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerLabelCategoryConfigS3Uri, "label-category-config-s3-uri", "", "", "Label Category Config S3 URI")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerLabelingJobAlgorithmsConfig, "labeling-job-algorithms-config", "", "", "Labeling Job Algorithms Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerLabelingJobName, "labeling-job-name", "", "", "Labeling Job Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerLandingUri, "landing-uri", "", "", "Landing URI")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerLastModifiedTimeAfter, "last-modified-time-after", "", "", "Last Modified Time After")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerLastModifiedTimeBefore, "last-modified-time-before", "", "", "Last Modified Time Before")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerLatestHeartbeatAfter, "latest-heartbeat-after", "", "", "Latest Heartbeat After")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerLifecycleConfigName, "lifecycle-config-name", "", "", "Lifecycle Config Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerLineageGroupName, "lineage-group-name", "", "", "Lineage Group Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerMaintenanceConfig, "maintenance-config", "", "", "Maintenance Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerMaxConcurrentTransforms, "max-concurrent-transforms", "", "", "Max Concurrent Transforms")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerMaxDepth, "max-depth", "", "", "Max Depth")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerMaxInstanceCount, "max-instance-count", "", "", "Max Instance Count")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerMaxPayloadInMB, "max-payload-in-mb", "", "", "Max Payload In Mb")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerMaxResults, "max-results", "", "", "Max Results")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerMaxSchemaVersion, "max-schema-version", "", "", "Max Schema Version")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerMemberDefinitions, "member-definitions", "", "", "Member Definitions")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerMetadataProperties, "metadata-properties", "", "", "Metadata Properties")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerMetricsConfig, "metrics-config", "", "", "Metrics Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerMinVersion, "min-version", "", "", "Min Version")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerMLFramework, "ml-framework", "", "", "Ml Framework")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerMlflowConfig, "mlflow-config", "", "", "Mlflow Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerMlflowExperimentName, "mlflow-experiment-name", "", "", "Mlflow Experiment Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerMlflowVersion, "mlflow-version", "", "", "Mlflow Version")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerModelApprovalStatus, "model-approval-status", "", "", "Model Approval Status")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerModelBiasAppSpecification, "model-bias-app-specification", "", "", "Model Bias App Specification")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerModelBiasBaselineConfig, "model-bias-baseline-config", "", "", "Model Bias Baseline Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerModelBiasJobInput, "model-bias-job-input", "", "", "Model Bias Job Input")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerModelBiasJobOutputConfig, "model-bias-job-output-config", "", "", "Model Bias Job Output Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerModelCard, "model-card", "", "", "Model Card")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerModelCardExportJobArn, "model-card-export-job-arn", "", "", "Model Card Export Job ARN")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerModelCardExportJobName, "model-card-export-job-name", "", "", "Model Card Export Job Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerModelCardExportJobNameContains, "model-card-export-job-name-contains", "", "", "Model Card Export Job Name Contains")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerModelCardName, "model-card-name", "", "", "Model Card Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerModelCardStatus, "model-card-status", "", "", "Model Card Status")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerModelCardVersion, "model-card-version", "", "", "Model Card Version")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerModelClientConfig, "model-client-config", "", "", "Model Client Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerModelConfigs, "model-configs", "", "", "Model Configs")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerModelDeployConfig, "model-deploy-config", "", "", "Model Deploy Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerModelExplainabilityAppSpecification, "model-explainability-app-specification", "", "", "Model Explainability App Specification")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerModelExplainabilityBaselineConfig, "model-explainability-baseline-config", "", "", "Model Explainability Baseline Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerModelExplainabilityJobInput, "model-explainability-job-input", "", "", "Model Explainability Job Input")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerModelExplainabilityJobOutputConfig, "model-explainability-job-output-config", "", "", "Model Explainability Job Output Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerModelLifeCycle, "model-life-cycle", "", "", "Model Life Cycle")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerModelMetrics, "model-metrics", "", "", "Model Metrics")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerModelName, "model-name", "", "", "Model Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerModelNameContains, "model-name-contains", "", "", "Model Name Contains")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerModelNameEquals, "model-name-equals", "", "", "Model Name Equals")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerModelPackageArn, "model-package-arn", "", "", "Model Package ARN")
	_sagemakerCmd.Flags().StringSliceVarP(&_sagemakerModelPackageArnList, "model-package-arn-list", "", nil, "Model Package ARN List")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerModelPackageConfig, "model-package-config", "", "", "Model Package Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerModelPackageDescription, "model-package-description", "", "", "Model Package Description")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerModelPackageGroupDescription, "model-package-group-description", "", "", "Model Package Group Description")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerModelPackageGroupName, "model-package-group-name", "", "", "Model Package Group Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerModelPackageName, "model-package-name", "", "", "Model Package Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerModelPackageRegistrationType, "model-package-registration-type", "", "", "Model Package Registration Type")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerModelPackageType, "model-package-type", "", "", "Model Package Type")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerModelPackageVersionArn, "model-package-version-arn", "", "", "Model Package Version ARN")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerModelPackageVersionArnEquals, "model-package-version-arn-equals", "", "", "Model Package Version ARN Equals")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerModelQualityAppSpecification, "model-quality-app-specification", "", "", "Model Quality App Specification")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerModelQualityBaselineConfig, "model-quality-baseline-config", "", "", "Model Quality Baseline Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerModelQualityJobInput, "model-quality-job-input", "", "", "Model Quality Job Input")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerModelQualityJobOutputConfig, "model-quality-job-output-config", "", "", "Model Quality Job Output Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerModelRegistrationMode, "model-registration-mode", "", "", "Model Registration Mode")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerModelSource, "model-source", "", "", "Model Source")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerModelVariantActions, "model-variant-actions", "", "", "Model Variant Actions")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerModelVariants, "model-variants", "", "", "Model Variants")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerModelVersion, "model-version", "", "", "Model Version")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerModifiedTimeAfter, "modified-time-after", "", "", "Modified Time After")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerModifiedTimeBefore, "modified-time-before", "", "", "Modified Time Before")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerMonitoringAlertName, "monitoring-alert-name", "", "", "Monitoring Alert Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerMonitoringJobDefinitionName, "monitoring-job-definition-name", "", "", "Monitoring Job Definition Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerMonitoringScheduleConfig, "monitoring-schedule-config", "", "", "Monitoring Schedule Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerMonitoringScheduleName, "monitoring-schedule-name", "", "", "Monitoring Schedule Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerMonitoringTypeEquals, "monitoring-type-equals", "", "", "Monitoring Type Equals")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerName, "name", "", "", "Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerNameContains, "name-contains", "", "", "Name Contains")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerNetworkConfig, "network-config", "", "", "Network Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerNextToken, "next-token", "", "", "Next Token")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerNodeId, "node-id", "", "", "Node ID")
	_sagemakerCmd.Flags().StringSliceVarP(&_sagemakerNodeIds, "node-ids", "", nil, "Node Ids")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerNodeLogicalId, "node-logical-id", "", "", "Node Logical ID")
	_sagemakerCmd.Flags().StringSliceVarP(&_sagemakerNodeLogicalIds, "node-logical-ids", "", nil, "Node Logical Ids")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerNodeProvisioningMode, "node-provisioning-mode", "", "", "Node Provisioning Mode")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerNodeRecovery, "node-recovery", "", "", "Node Recovery")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerNodesToAdd, "nodes-to-add", "", "", "Nodes To Add")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerNotebookInstanceLifecycleConfigName, "notebook-instance-lifecycle-config-name", "", "", "Notebook Instance Lifecycle Config Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerNotebookInstanceLifecycleConfigNameContains, "notebook-instance-lifecycle-config-name-contains", "", "", "Notebook Instance Lifecycle Config Name Contains")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerNotebookInstanceName, "notebook-instance-name", "", "", "Notebook Instance Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerNotificationConfiguration, "notification-configuration", "", "", "Notification Configuration")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerOfflineStoreConfig, "offline-store-config", "", "", "Offline Store Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerOfflineStoreStatusEquals, "offline-store-status-equals", "", "", "Offline Store Status Equals")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerOidcConfig, "oidc-config", "", "", "OIDC Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerOnCreate, "on-create", "", "", "On Create")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerOnStart, "on-start", "", "", "On Start")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerOnlineStoreConfig, "online-store-config", "", "", "Online Store Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerOptimizationConfigs, "optimization-configs", "", "", "Optimization Configs")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerOptimizationContains, "optimization-contains", "", "", "Optimization Contains")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerOptimizationEnvironment, "optimization-environment", "", "", "Optimization Environment")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerOptimizationJobName, "optimization-job-name", "", "", "Optimization Job Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerOrchestrator, "orchestrator", "", "", "Orchestrator")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerOutputArtifacts, "output-artifacts", "", "", "Output Artifacts")
	_sagemakerCmd.Flags().StringSliceVarP(&_sagemakerOutputArtifactsToRemove, "output-artifacts-to-remove", "", nil, "Output Artifacts To Remove")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerOutputConfig, "output-config", "", "", "Output Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerOutputDataConfig, "output-data-config", "", "", "Output Data Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerOutputParameters, "output-parameters", "", "", "Output Parameters")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerOwnershipSettings, "ownership-settings", "", "", "Ownership Settings")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerParallelismConfiguration, "parallelism-configuration", "", "", "Parallelism Configuration")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerParameterAdditions, "parameter-additions", "", "", "Parameter Additions")
	_sagemakerCmd.Flags().StringSliceVarP(&_sagemakerParameterRemovals, "parameter-removals", "", nil, "Parameter Removals")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerParameters, "parameters", "", "", "Parameters")
	_sagemakerCmd.Flags().StringSliceVarP(&_sagemakerParametersToRemove, "parameters-to-remove", "", nil, "Parameters To Remove")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerPipelineArn, "pipeline-arn", "", "", "Pipeline ARN")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerPipelineDefinition, "pipeline-definition", "", "", "Pipeline Definition")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerPipelineDefinitionS3Location, "pipeline-definition-s3-location", "", "", "Pipeline Definition S3 Location")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerPipelineDescription, "pipeline-description", "", "", "Pipeline Description")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerPipelineDisplayName, "pipeline-display-name", "", "", "Pipeline Display Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerPipelineExecutionArn, "pipeline-execution-arn", "", "", "Pipeline Execution ARN")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerPipelineExecutionDescription, "pipeline-execution-description", "", "", "Pipeline Execution Description")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerPipelineExecutionDisplayName, "pipeline-execution-display-name", "", "", "Pipeline Execution Display Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerPipelineName, "pipeline-name", "", "", "Pipeline Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerPipelineNamePrefix, "pipeline-name-prefix", "", "", "Pipeline Name Prefix")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerPipelineParameters, "pipeline-parameters", "", "", "Pipeline Parameters")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerPipelineVersionDescription, "pipeline-version-description", "", "", "Pipeline Version Description")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerPipelineVersionDisplayName, "pipeline-version-display-name", "", "", "Pipeline Version Display Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerPipelineVersionId, "pipeline-version-id", "", "", "Pipeline Version ID")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerPlatformIdentifier, "platform-identifier", "", "", "Platform Identifier")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerPrimaryContainer, "primary-container", "", "", "Primary Container")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerProblemType, "problem-type", "", "", "Problem Type")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerProcessingInputs, "processing-inputs", "", "", "Processing Inputs")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerProcessingJobName, "processing-job-name", "", "", "Processing Job Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerProcessingOutputConfig, "processing-output-config", "", "", "Processing Output Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerProcessingResources, "processing-resources", "", "", "Processing Resources")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerProcessor, "processor", "", "", "Processor")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerProductionVariants, "production-variants", "", "", "Production Variants")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerProfilerConfig, "profiler-config", "", "", "Profiler Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerProfilerRuleConfigurations, "profiler-rule-configurations", "", "", "Profiler Rule Configurations")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerProgrammingLang, "programming-lang", "", "", "Programming Lang")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerProjectDescription, "project-description", "", "", "Project Description")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerProjectName, "project-name", "", "", "Project Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerProperties, "properties", "", "", "Properties")
	_sagemakerCmd.Flags().StringSliceVarP(&_sagemakerPropertiesToRemove, "properties-to-remove", "", nil, "Properties To Remove")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerReason, "reason", "", "", "Reason")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerRecommendationId, "recommendation-id", "", "", "Recommendation ID")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerRecordIdentifierFeatureName, "record-identifier-feature-name", "", "", "Record Identifier Feature Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerRecoveryMode, "recovery-mode", "", "", "Recovery Mode")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerReleaseNotes, "release-notes", "", "", "Release Notes")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerRemoteDebugConfig, "remote-debug-config", "", "", "Remote Debug Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerReservedCapacityArn, "reserved-capacity-arn", "", "", "Reserved Capacity ARN")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerResource, "resource", "", "", "Resource")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerResourceArn, "resource-arn", "", "", "Resource ARN")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerResourceConfig, "resource-config", "", "", "Resource Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerResourceIdentifier, "resource-identifier", "", "", "Resource Identifier")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerResourceKey, "resource-key", "", "", "Resource Key")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerResourcePolicy, "resource-policy", "", "", "Resource Policy")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerResourceSpec, "resource-spec", "", "", "Resource Spec")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerResourceType, "resource-type", "", "", "Resource Type")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerRestrictedInstanceGroups, "restricted-instance-groups", "", "", "Restricted Instance Groups")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerRetainAllVariantProperties, "retain-all-variant-properties", "", "", "Retain All Variant Properties")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerRetainDeploymentConfig, "retain-deployment-config", "", "", "Retain Deployment Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerRetentionPolicy, "retention-policy", "", "", "Retention Policy")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerRetryStrategy, "retry-strategy", "", "", "Retry Strategy")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerRoleArn, "role-arn", "", "", "Role ARN")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerRootAccess, "root-access", "", "", "Root Access")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerRuntimeConfig, "runtime-config", "", "", "Runtime Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerS3StorageConfig, "s3-storage-config", "", "", "S3 Storage Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerSageMakerPublicHubContentArn, "sage-maker-public-hub-content-arn", "", "", "Sage Maker Public Hub Content ARN")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerSamplePayloadUrl, "sample-payload-url", "", "", "Sample Payload URL")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerScalingPolicyObjective, "scaling-policy-objective", "", "", "Scaling Policy Objective")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerSchedule, "schedule", "", "", "Schedule")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerScheduledTimeAfter, "scheduled-time-after", "", "", "Scheduled Time After")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerScheduledTimeBefore, "scheduled-time-before", "", "", "Scheduled Time Before")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerSchedulerConfig, "scheduler-config", "", "", "Scheduler Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerSearchExpression, "search-expression", "", "", "Search Expression")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerSecurityConfig, "security-config", "", "", "Security Config")
	_sagemakerCmd.Flags().StringSliceVarP(&_sagemakerSecurityGroupIds, "security-group-ids", "", nil, "Security Group Ids")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerSelectiveExecutionConfig, "selective-execution-config", "", "", "Selective Execution Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerServerlessJobConfig, "serverless-job-config", "", "", "Serverless Job Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerServiceCatalogProvisioningDetails, "service-catalog-provisioning-details", "", "", "Service Catalog Provisioning Details")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerServiceCatalogProvisioningUpdateDetails, "service-catalog-provisioning-update-details", "", "", "Service Catalog Provisioning Update Details")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerSessionChainingConfig, "session-chaining-config", "", "", "Session Chaining Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerSessionExpirationDurationInSeconds, "session-expiration-duration-in-seconds", "", "", "Session Expiration Duration In Seconds")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerShadowModeConfig, "shadow-mode-config", "", "", "Shadow Mode Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerShadowProductionVariants, "shadow-production-variants", "", "", "Shadow Production Variants")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerSingleSignOnUserIdentifier, "single-sign-on-user-identifier", "", "", "Single Sign On User Identifier")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerSingleSignOnUserValue, "single-sign-on-user-value", "", "", "Single Sign On User Value")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerSkipModelValidation, "skip-model-validation", "", "", "Skip Model Validation")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerSortBy, "sort-by", "", "", "Sort By")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerSortOrder, "sort-order", "", "", "Sort Order")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerSource, "source", "", "", "Source")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerSourceAlgorithmSpecification, "source-algorithm-specification", "", "", "Source Algorithm Specification")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerSourceArn, "source-arn", "", "", "Source ARN")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerSourceIpConfig, "source-ip-config", "", "", "Source IP Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerSourceType, "source-type", "", "", "Source Type")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerSourceUri, "source-uri", "", "", "Source URI")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerSpaceDisplayName, "space-display-name", "", "", "Space Display Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerSpaceName, "space-name", "", "", "Space Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerSpaceNameContains, "space-name-contains", "", "", "Space Name Contains")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerSpaceNameEquals, "space-name-equals", "", "", "Space Name Equals")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerSpaceSettings, "space-settings", "", "", "Space Settings")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerSpaceSharingSettings, "space-sharing-settings", "", "", "Space Sharing Settings")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerSpareInstanceCountPerUltraServer, "spare-instance-count-per-ultra-server", "", "", "Spare Instance Count Per Ultra Server")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerSpecification, "specification", "", "", "Specification")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerStageName, "stage-name", "", "", "Stage Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerStages, "stages", "", "", "Stages")
	_sagemakerCmd.Flags().StringSliceVarP(&_sagemakerStartArns, "start-arns", "", nil, "Start Arns")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerStartTime, "start-time", "", "", "Start Time")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerStartTimeAfter, "start-time-after", "", "", "Start Time After")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerStartTimeBefore, "start-time-before", "", "", "Start Time Before")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerStatus, "status", "", "", "Status")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerStatusEquals, "status-equals", "", "", "Status Equals")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerStepType, "step-type", "", "", "Step Type")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerStoppingCondition, "stopping-condition", "", "", "Stopping Condition")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerStoppingConditions, "stopping-conditions", "", "", "Stopping Conditions")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerStudioLifecycleConfigAppType, "studio-lifecycle-config-app-type", "", "", "Studio Lifecycle Config App Type")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerStudioLifecycleConfigContent, "studio-lifecycle-config-content", "", "", "Studio Lifecycle Config Content")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerStudioLifecycleConfigName, "studio-lifecycle-config-name", "", "", "Studio Lifecycle Config Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerSubnetId, "subnet-id", "", "", "Subnet ID")
	_sagemakerCmd.Flags().StringSliceVarP(&_sagemakerSubnetIds, "subnet-ids", "", nil, "Subnet Ids")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerSuggestionQuery, "suggestion-query", "", "", "Suggestion Query")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerSupportStatus, "support-status", "", "", "Support Status")
	_sagemakerCmd.Flags().StringSliceVarP(&_sagemakerTagKeys, "tag-keys", "", nil, "Tag Keys")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerTagPropagation, "tag-propagation", "", "", "Tag Propagation")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerTags, "tags", "", "", "Tags")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerTargetCpuUtilizationPerCore, "target-cpu-utilization-per-core", "", "", "Target CPU Utilization Per Core")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerTargetResources, "target-resources", "", "", "Target Resources")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerTargetVersion, "target-version", "", "", "Target Version")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerTask, "task", "", "", "Task")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerTemplateProviders, "template-providers", "", "", "Template Providers")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerTemplateProvidersToUpdate, "template-providers-to-update", "", "", "Template Providers To Update")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerTensorBoardOutputConfig, "tensor-board-output-config", "", "", "Tensor Board Output Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerThroughputConfig, "throughput-config", "", "", "Throughput Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerTier, "tier", "", "", "Tier")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerTieredStorageConfig, "tiered-storage-config", "", "", "Tiered Storage Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerTrackingServerName, "tracking-server-name", "", "", "Tracking Server Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerTrackingServerSize, "tracking-server-size", "", "", "Tracking Server Size")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerTrackingServerStatus, "tracking-server-status", "", "", "Tracking Server Status")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerTrainingJobDefinition, "training-job-definition", "", "", "Training Job Definition")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerTrainingJobDefinitions, "training-job-definitions", "", "", "Training Job Definitions")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerTrainingJobName, "training-job-name", "", "", "Training Job Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerTrainingPlanArn, "training-plan-arn", "", "", "Training Plan ARN")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerTrainingPlanArnEquals, "training-plan-arn-equals", "", "", "Training Plan ARN Equals")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerTrainingPlanName, "training-plan-name", "", "", "Training Plan Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerTrainingPlanOfferingId, "training-plan-offering-id", "", "", "Training Plan Offering ID")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerTrainingSpecification, "training-specification", "", "", "Training Specification")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerTransformInput, "transform-input", "", "", "Transform Input")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerTransformJobName, "transform-job-name", "", "", "Transform Job Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerTransformOutput, "transform-output", "", "", "Transform Output")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerTransformResources, "transform-resources", "", "", "Transform Resources")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerTrialComponentName, "trial-component-name", "", "", "Trial Component Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerTrialName, "trial-name", "", "", "Trial Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerType, "type", "", "", "Type")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerUiTemplate, "ui-template", "", "", "Ui Template")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerUltraServerCount, "ultra-server-count", "", "", "Ultra Server Count")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerUltraServerType, "ultra-server-type", "", "", "Ultra Server Type")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerUserProfileName, "user-profile-name", "", "", "User Profile Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerUserProfileNameContains, "user-profile-name-contains", "", "", "User Profile Name Contains")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerUserProfileNameEquals, "user-profile-name-equals", "", "", "User Profile Name Equals")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerUserSettings, "user-settings", "", "", "User Settings")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerValidationSpecification, "validation-specification", "", "", "Validation Specification")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerVariantName, "variant-name", "", "", "Variant Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerVariantNameEquals, "variant-name-equals", "", "", "Variant Name Equals")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerVendorGuidance, "vendor-guidance", "", "", "Vendor Guidance")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerVersion, "version", "", "", "Version")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerVisibilityConditions, "visibility-conditions", "", "", "Visibility Conditions")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerVolumeId, "volume-id", "", "", "Volume ID")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerVolumeSizeInGB, "volume-size-in-gb", "", "", "Volume Size In Gb")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerVpcConfig, "vpc-config", "", "", "VPC Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerVpcId, "vpc-id", "", "", "VPC ID")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerWarmPoolStatusEquals, "warm-pool-status-equals", "", "", "Warm Pool Status Equals")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerWarmStartConfig, "warm-start-config", "", "", "Warm Start Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerWeeklyMaintenanceWindowStart, "weekly-maintenance-window-start", "", "", "Weekly Maintenance Window Start")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerWorkerAccessConfiguration, "worker-access-configuration", "", "", "Worker Access Configuration")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerWorkforceName, "workforce-name", "", "", "Workforce Name")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerWorkforceVpcConfig, "workforce-vpc-config", "", "", "Workforce VPC Config")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerWorkteamArn, "workteam-arn", "", "", "Workteam ARN")
	_sagemakerCmd.Flags().StringVarP(&_sagemakerWorkteamName, "workteam-name", "", "", "Workteam Name")

	_sagemakerCmd.Flags().BoolVarP(&_sagemakerAddAssociation, "add-association", "", false, "Add Association")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerAddTags, "add-tags", "", false, "Add Tags")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerAssociateTrialComponent, "associate-trial-component", "", false, "Associate Trial Component")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerAttachClusterNodeVolume, "attach-cluster-node-volume", "", false, "Attach Cluster Node Volume")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerBatchAddClusterNodes, "batch-add-cluster-nodes", "", false, "Batch Add Cluster Nodes")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerBatchDeleteClusterNodes, "batch-delete-cluster-nodes", "", false, "Batch Delete Cluster Nodes")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerBatchDescribeModelPackage, "batch-describe-model-package", "", false, "Batch Describe Model Package")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerBatchRebootClusterNodes, "batch-reboot-cluster-nodes", "", false, "Batch Reboot Cluster Nodes")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerBatchReplaceClusterNodes, "batch-replace-cluster-nodes", "", false, "Batch Replace Cluster Nodes")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreateAction, "create-action", "", false, "Create Action")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreateAlgorithm, "create-algorithm", "", false, "Create Algorithm")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreateApp, "create-app", "", false, "Create App")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreateAppImageConfig, "create-app-image-config", "", false, "Create App Image Config")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreateArtifact, "create-artifact", "", false, "Create Artifact")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreateAutoMLJob, "create-auto-ml-job", "", false, "Create Auto Ml Job")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreateAutoMLJobV2, "create-auto-ml-job-v2", "", false, "Create Auto Ml Job V2")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreateCluster, "create-cluster", "", false, "Create Cluster")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreateClusterSchedulerConfig, "create-cluster-scheduler-config", "", false, "Create Cluster Scheduler Config")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreateCodeRepository, "create-code-repository", "", false, "Create Code Repository")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreateCompilationJob, "create-compilation-job", "", false, "Create Compilation Job")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreateComputeQuota, "create-compute-quota", "", false, "Create Compute Quota")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreateContext, "create-context", "", false, "Create Context")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreateDataQualityJobDefinition, "create-data-quality-job-definition", "", false, "Create Data Quality Job Definition")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreateDeviceFleet, "create-device-fleet", "", false, "Create Device Fleet")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreateDomain, "create-domain", "", false, "Create Domain")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreateEdgeDeploymentPlan, "create-edge-deployment-plan", "", false, "Create Edge Deployment Plan")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreateEdgeDeploymentStage, "create-edge-deployment-stage", "", false, "Create Edge Deployment Stage")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreateEdgePackagingJob, "create-edge-packaging-job", "", false, "Create Edge Packaging Job")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreateEndpoint, "create-endpoint", "", false, "Create Endpoint")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreateEndpointConfig, "create-endpoint-config", "", false, "Create Endpoint Config")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreateExperiment, "create-experiment", "", false, "Create Experiment")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreateFeatureGroup, "create-feature-group", "", false, "Create Feature Group")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreateFlowDefinition, "create-flow-definition", "", false, "Create Flow Definition")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreateHub, "create-hub", "", false, "Create Hub")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreateHubContentPresignedUrls, "create-hub-content-presigned-urls", "", false, "Create Hub Content Presigned Urls")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreateHubContentReference, "create-hub-content-reference", "", false, "Create Hub Content Reference")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreateHumanTaskUi, "create-human-task-ui", "", false, "Create Human Task Ui")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreateHyperParameterTuningJob, "create-hyper-parameter-tuning-job", "", false, "Create Hyper Parameter Tuning Job")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreateImage, "create-image", "", false, "Create Image")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreateImageVersion, "create-image-version", "", false, "Create Image Version")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreateInferenceComponent, "create-inference-component", "", false, "Create Inference Component")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreateInferenceExperiment, "create-inference-experiment", "", false, "Create Inference Experiment")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreateInferenceRecommendationsJob, "create-inference-recommendations-job", "", false, "Create Inference Recommendations Job")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreateLabelingJob, "create-labeling-job", "", false, "Create Labeling Job")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreateMlflowApp, "create-mlflow-app", "", false, "Create Mlflow App")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreateMlflowTrackingServer, "create-mlflow-tracking-server", "", false, "Create Mlflow Tracking Server")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreateModel, "create-model", "", false, "Create Model")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreateModelBiasJobDefinition, "create-model-bias-job-definition", "", false, "Create Model Bias Job Definition")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreateModelCard, "create-model-card", "", false, "Create Model Card")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreateModelCardExportJob, "create-model-card-export-job", "", false, "Create Model Card Export Job")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreateModelExplainabilityJobDefinition, "create-model-explainability-job-definition", "", false, "Create Model Explainability Job Definition")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreateModelPackage, "create-model-package", "", false, "Create Model Package")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreateModelPackageGroup, "create-model-package-group", "", false, "Create Model Package Group")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreateModelQualityJobDefinition, "create-model-quality-job-definition", "", false, "Create Model Quality Job Definition")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreateMonitoringSchedule, "create-monitoring-schedule", "", false, "Create Monitoring Schedule")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreateNotebookInstance, "create-notebook-instance", "", false, "Create Notebook Instance")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreateNotebookInstanceLifecycleConfig, "create-notebook-instance-lifecycle-config", "", false, "Create Notebook Instance Lifecycle Config")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreateOptimizationJob, "create-optimization-job", "", false, "Create Optimization Job")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreatePartnerApp, "create-partner-app", "", false, "Create Partner App")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreatePartnerAppPresignedUrl, "create-partner-app-presigned-url", "", false, "Create Partner App Presigned URL")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreatePipeline, "create-pipeline", "", false, "Create Pipeline")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreatePresignedDomainUrl, "create-presigned-domain-url", "", false, "Create Presigned Domain URL")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreatePresignedMlflowAppUrl, "create-presigned-mlflow-app-url", "", false, "Create Presigned Mlflow App URL")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreatePresignedMlflowTrackingServerUrl, "create-presigned-mlflow-tracking-server-url", "", false, "Create Presigned Mlflow Tracking Server URL")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreatePresignedNotebookInstanceUrl, "create-presigned-notebook-instance-url", "", false, "Create Presigned Notebook Instance URL")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreateProcessingJob, "create-processing-job", "", false, "Create Processing Job")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreateProject, "create-project", "", false, "Create Project")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreateSpace, "create-space", "", false, "Create Space")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreateStudioLifecycleConfig, "create-studio-lifecycle-config", "", false, "Create Studio Lifecycle Config")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreateTrainingJob, "create-training-job", "", false, "Create Training Job")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreateTrainingPlan, "create-training-plan", "", false, "Create Training Plan")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreateTransformJob, "create-transform-job", "", false, "Create Transform Job")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreateTrial, "create-trial", "", false, "Create Trial")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreateTrialComponent, "create-trial-component", "", false, "Create Trial Component")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreateUserProfile, "create-user-profile", "", false, "Create User Profile")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreateWorkforce, "create-workforce", "", false, "Create Workforce")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerCreateWorkteam, "create-workteam", "", false, "Create Workteam")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDeleteAction, "delete-action", "", false, "Delete Action")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDeleteAlgorithm, "delete-algorithm", "", false, "Delete Algorithm")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDeleteApp, "delete-app", "", false, "Delete App")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDeleteAppImageConfig, "delete-app-image-config", "", false, "Delete App Image Config")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDeleteArtifact, "delete-artifact", "", false, "Delete Artifact")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDeleteAssociation, "delete-association", "", false, "Delete Association")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDeleteCluster, "delete-cluster", "", false, "Delete Cluster")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDeleteClusterSchedulerConfig, "delete-cluster-scheduler-config", "", false, "Delete Cluster Scheduler Config")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDeleteCodeRepository, "delete-code-repository", "", false, "Delete Code Repository")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDeleteCompilationJob, "delete-compilation-job", "", false, "Delete Compilation Job")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDeleteComputeQuota, "delete-compute-quota", "", false, "Delete Compute Quota")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDeleteContext, "delete-context", "", false, "Delete Context")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDeleteDataQualityJobDefinition, "delete-data-quality-job-definition", "", false, "Delete Data Quality Job Definition")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDeleteDeviceFleet, "delete-device-fleet", "", false, "Delete Device Fleet")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDeleteDomain, "delete-domain", "", false, "Delete Domain")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDeleteEdgeDeploymentPlan, "delete-edge-deployment-plan", "", false, "Delete Edge Deployment Plan")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDeleteEdgeDeploymentStage, "delete-edge-deployment-stage", "", false, "Delete Edge Deployment Stage")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDeleteEndpoint, "delete-endpoint", "", false, "Delete Endpoint")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDeleteEndpointConfig, "delete-endpoint-config", "", false, "Delete Endpoint Config")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDeleteExperiment, "delete-experiment", "", false, "Delete Experiment")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDeleteFeatureGroup, "delete-feature-group", "", false, "Delete Feature Group")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDeleteFlowDefinition, "delete-flow-definition", "", false, "Delete Flow Definition")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDeleteHub, "delete-hub", "", false, "Delete Hub")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDeleteHubContent, "delete-hub-content", "", false, "Delete Hub Content")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDeleteHubContentReference, "delete-hub-content-reference", "", false, "Delete Hub Content Reference")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDeleteHumanTaskUi, "delete-human-task-ui", "", false, "Delete Human Task Ui")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDeleteHyperParameterTuningJob, "delete-hyper-parameter-tuning-job", "", false, "Delete Hyper Parameter Tuning Job")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDeleteImage, "delete-image", "", false, "Delete Image")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDeleteImageVersion, "delete-image-version", "", false, "Delete Image Version")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDeleteInferenceComponent, "delete-inference-component", "", false, "Delete Inference Component")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDeleteInferenceExperiment, "delete-inference-experiment", "", false, "Delete Inference Experiment")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDeleteMlflowApp, "delete-mlflow-app", "", false, "Delete Mlflow App")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDeleteMlflowTrackingServer, "delete-mlflow-tracking-server", "", false, "Delete Mlflow Tracking Server")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDeleteModel, "delete-model", "", false, "Delete Model")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDeleteModelBiasJobDefinition, "delete-model-bias-job-definition", "", false, "Delete Model Bias Job Definition")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDeleteModelCard, "delete-model-card", "", false, "Delete Model Card")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDeleteModelExplainabilityJobDefinition, "delete-model-explainability-job-definition", "", false, "Delete Model Explainability Job Definition")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDeleteModelPackage, "delete-model-package", "", false, "Delete Model Package")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDeleteModelPackageGroup, "delete-model-package-group", "", false, "Delete Model Package Group")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDeleteModelPackageGroupPolicy, "delete-model-package-group-policy", "", false, "Delete Model Package Group Policy")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDeleteModelQualityJobDefinition, "delete-model-quality-job-definition", "", false, "Delete Model Quality Job Definition")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDeleteMonitoringSchedule, "delete-monitoring-schedule", "", false, "Delete Monitoring Schedule")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDeleteNotebookInstance, "delete-notebook-instance", "", false, "Delete Notebook Instance")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDeleteNotebookInstanceLifecycleConfig, "delete-notebook-instance-lifecycle-config", "", false, "Delete Notebook Instance Lifecycle Config")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDeleteOptimizationJob, "delete-optimization-job", "", false, "Delete Optimization Job")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDeletePartnerApp, "delete-partner-app", "", false, "Delete Partner App")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDeletePipeline, "delete-pipeline", "", false, "Delete Pipeline")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDeleteProcessingJob, "delete-processing-job", "", false, "Delete Processing Job")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDeleteProject, "delete-project", "", false, "Delete Project")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDeleteSpace, "delete-space", "", false, "Delete Space")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDeleteStudioLifecycleConfig, "delete-studio-lifecycle-config", "", false, "Delete Studio Lifecycle Config")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDeleteTags, "delete-tags", "", false, "Delete Tags")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDeleteTrainingJob, "delete-training-job", "", false, "Delete Training Job")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDeleteTrial, "delete-trial", "", false, "Delete Trial")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDeleteTrialComponent, "delete-trial-component", "", false, "Delete Trial Component")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDeleteUserProfile, "delete-user-profile", "", false, "Delete User Profile")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDeleteWorkforce, "delete-workforce", "", false, "Delete Workforce")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDeleteWorkteam, "delete-workteam", "", false, "Delete Workteam")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDeregisterDevices, "deregister-devices", "", false, "Deregister Devices")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeAction, "describe-action", "", false, "Describe Action")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeAlgorithm, "describe-algorithm", "", false, "Describe Algorithm")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeApp, "describe-app", "", false, "Describe App")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeAppImageConfig, "describe-app-image-config", "", false, "Describe App Image Config")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeArtifact, "describe-artifact", "", false, "Describe Artifact")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeAutoMLJob, "describe-auto-ml-job", "", false, "Describe Auto Ml Job")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeAutoMLJobV2, "describe-auto-ml-job-v2", "", false, "Describe Auto Ml Job V2")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeCluster, "describe-cluster", "", false, "Describe Cluster")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeClusterEvent, "describe-cluster-event", "", false, "Describe Cluster Event")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeClusterNode, "describe-cluster-node", "", false, "Describe Cluster Node")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeClusterSchedulerConfig, "describe-cluster-scheduler-config", "", false, "Describe Cluster Scheduler Config")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeCodeRepository, "describe-code-repository", "", false, "Describe Code Repository")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeCompilationJob, "describe-compilation-job", "", false, "Describe Compilation Job")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeComputeQuota, "describe-compute-quota", "", false, "Describe Compute Quota")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeContext, "describe-context", "", false, "Describe Context")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeDataQualityJobDefinition, "describe-data-quality-job-definition", "", false, "Describe Data Quality Job Definition")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeDevice, "describe-device", "", false, "Describe Device")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeDeviceFleet, "describe-device-fleet", "", false, "Describe Device Fleet")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeDomain, "describe-domain", "", false, "Describe Domain")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeEdgeDeploymentPlan, "describe-edge-deployment-plan", "", false, "Describe Edge Deployment Plan")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeEdgePackagingJob, "describe-edge-packaging-job", "", false, "Describe Edge Packaging Job")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeEndpoint, "describe-endpoint", "", false, "Describe Endpoint")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeEndpointConfig, "describe-endpoint-config", "", false, "Describe Endpoint Config")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeExperiment, "describe-experiment", "", false, "Describe Experiment")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeFeatureGroup, "describe-feature-group", "", false, "Describe Feature Group")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeFeatureMetadata, "describe-feature-metadata", "", false, "Describe Feature Metadata")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeFlowDefinition, "describe-flow-definition", "", false, "Describe Flow Definition")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeHub, "describe-hub", "", false, "Describe Hub")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeHubContent, "describe-hub-content", "", false, "Describe Hub Content")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeHumanTaskUi, "describe-human-task-ui", "", false, "Describe Human Task Ui")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeHyperParameterTuningJob, "describe-hyper-parameter-tuning-job", "", false, "Describe Hyper Parameter Tuning Job")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeImage, "describe-image", "", false, "Describe Image")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeImageVersion, "describe-image-version", "", false, "Describe Image Version")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeInferenceComponent, "describe-inference-component", "", false, "Describe Inference Component")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeInferenceExperiment, "describe-inference-experiment", "", false, "Describe Inference Experiment")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeInferenceRecommendationsJob, "describe-inference-recommendations-job", "", false, "Describe Inference Recommendations Job")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeLabelingJob, "describe-labeling-job", "", false, "Describe Labeling Job")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeLineageGroup, "describe-lineage-group", "", false, "Describe Lineage Group")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeMlflowApp, "describe-mlflow-app", "", false, "Describe Mlflow App")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeMlflowTrackingServer, "describe-mlflow-tracking-server", "", false, "Describe Mlflow Tracking Server")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeModel, "describe-model", "", false, "Describe Model")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeModelBiasJobDefinition, "describe-model-bias-job-definition", "", false, "Describe Model Bias Job Definition")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeModelCard, "describe-model-card", "", false, "Describe Model Card")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeModelCardExportJob, "describe-model-card-export-job", "", false, "Describe Model Card Export Job")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeModelExplainabilityJobDefinition, "describe-model-explainability-job-definition", "", false, "Describe Model Explainability Job Definition")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeModelPackage, "describe-model-package", "", false, "Describe Model Package")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeModelPackageGroup, "describe-model-package-group", "", false, "Describe Model Package Group")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeModelQualityJobDefinition, "describe-model-quality-job-definition", "", false, "Describe Model Quality Job Definition")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeMonitoringSchedule, "describe-monitoring-schedule", "", false, "Describe Monitoring Schedule")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeNotebookInstance, "describe-notebook-instance", "", false, "Describe Notebook Instance")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeNotebookInstanceLifecycleConfig, "describe-notebook-instance-lifecycle-config", "", false, "Describe Notebook Instance Lifecycle Config")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeOptimizationJob, "describe-optimization-job", "", false, "Describe Optimization Job")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribePartnerApp, "describe-partner-app", "", false, "Describe Partner App")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribePipeline, "describe-pipeline", "", false, "Describe Pipeline")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribePipelineDefinitionForExecution, "describe-pipeline-definition-for-execution", "", false, "Describe Pipeline Definition For Execution")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribePipelineExecution, "describe-pipeline-execution", "", false, "Describe Pipeline Execution")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeProcessingJob, "describe-processing-job", "", false, "Describe Processing Job")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeProject, "describe-project", "", false, "Describe Project")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeReservedCapacity, "describe-reserved-capacity", "", false, "Describe Reserved Capacity")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeSpace, "describe-space", "", false, "Describe Space")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeStudioLifecycleConfig, "describe-studio-lifecycle-config", "", false, "Describe Studio Lifecycle Config")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeSubscribedWorkteam, "describe-subscribed-workteam", "", false, "Describe Subscribed Workteam")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeTrainingJob, "describe-training-job", "", false, "Describe Training Job")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeTrainingPlan, "describe-training-plan", "", false, "Describe Training Plan")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeTransformJob, "describe-transform-job", "", false, "Describe Transform Job")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeTrial, "describe-trial", "", false, "Describe Trial")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeTrialComponent, "describe-trial-component", "", false, "Describe Trial Component")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeUserProfile, "describe-user-profile", "", false, "Describe User Profile")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeWorkforce, "describe-workforce", "", false, "Describe Workforce")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDescribeWorkteam, "describe-workteam", "", false, "Describe Workteam")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDetachClusterNodeVolume, "detach-cluster-node-volume", "", false, "Detach Cluster Node Volume")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDisableSagemakerServicecatalogPortfolio, "disable-sagemaker-servicecatalog-portfolio", "", false, "Disable Sagemaker Servicecatalog Portfolio")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerDisassociateTrialComponent, "disassociate-trial-component", "", false, "Disassociate Trial Component")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerEnableSagemakerServicecatalogPortfolio, "enable-sagemaker-servicecatalog-portfolio", "", false, "Enable Sagemaker Servicecatalog Portfolio")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerGetDeviceFleetReport, "get-device-fleet-report", "", false, "Get Device Fleet Report")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerGetLineageGroupPolicy, "get-lineage-group-policy", "", false, "Get Lineage Group Policy")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerGetModelPackageGroupPolicy, "get-model-package-group-policy", "", false, "Get Model Package Group Policy")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerGetSagemakerServicecatalogPortfolioStatus, "get-sagemaker-servicecatalog-portfolio-status", "", false, "Get Sagemaker Servicecatalog Portfolio Status")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerGetScalingConfigurationRecommendation, "get-scaling-configuration-recommendation", "", false, "Get Scaling Configuration Recommendation")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerGetSearchSuggestions, "get-search-suggestions", "", false, "Get Search Suggestions")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerImportHubContent, "import-hub-content", "", false, "Import Hub Content")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListActions, "list-actions", "", false, "List Actions")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListAlgorithms, "list-algorithms", "", false, "List Algorithms")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListAliases, "list-aliases", "", false, "List Aliases")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListAppImageConfigs, "list-app-image-configs", "", false, "List App Image Configs")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListApps, "list-apps", "", false, "List Apps")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListArtifacts, "list-artifacts", "", false, "List Artifacts")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListAssociations, "list-associations", "", false, "List Associations")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListAutoMLJobs, "list-auto-ml-jobs", "", false, "List Auto Ml Jobs")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListCandidatesForAutoMLJob, "list-candidates-for-auto-ml-job", "", false, "List Candidates For Auto Ml Job")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListClusterEvents, "list-cluster-events", "", false, "List Cluster Events")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListClusterNodes, "list-cluster-nodes", "", false, "List Cluster Nodes")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListClusterSchedulerConfigs, "list-cluster-scheduler-configs", "", false, "List Cluster Scheduler Configs")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListClusters, "list-clusters", "", false, "List Clusters")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListCodeRepositories, "list-code-repositories", "", false, "List Code Repositories")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListCompilationJobs, "list-compilation-jobs", "", false, "List Compilation Jobs")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListComputeQuotas, "list-compute-quotas", "", false, "List Compute Quotas")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListContexts, "list-contexts", "", false, "List Contexts")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListDataQualityJobDefinitions, "list-data-quality-job-definitions", "", false, "List Data Quality Job Definitions")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListDeviceFleets, "list-device-fleets", "", false, "List Device Fleets")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListDevices, "list-devices", "", false, "List Devices")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListDomains, "list-domains", "", false, "List Domains")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListEdgeDeploymentPlans, "list-edge-deployment-plans", "", false, "List Edge Deployment Plans")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListEdgePackagingJobs, "list-edge-packaging-jobs", "", false, "List Edge Packaging Jobs")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListEndpointConfigs, "list-endpoint-configs", "", false, "List Endpoint Configs")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListEndpoints, "list-endpoints", "", false, "List Endpoints")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListExperiments, "list-experiments", "", false, "List Experiments")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListFeatureGroups, "list-feature-groups", "", false, "List Feature Groups")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListFlowDefinitions, "list-flow-definitions", "", false, "List Flow Definitions")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListHubContentVersions, "list-hub-content-versions", "", false, "List Hub Content Versions")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListHubContents, "list-hub-contents", "", false, "List Hub Contents")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListHubs, "list-hubs", "", false, "List Hubs")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListHumanTaskUis, "list-human-task-uis", "", false, "List Human Task Uis")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListHyperParameterTuningJobs, "list-hyper-parameter-tuning-jobs", "", false, "List Hyper Parameter Tuning Jobs")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListImageVersions, "list-image-versions", "", false, "List Image Versions")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListImages, "list-images", "", false, "List Images")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListInferenceComponents, "list-inference-components", "", false, "List Inference Components")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListInferenceExperiments, "list-inference-experiments", "", false, "List Inference Experiments")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListInferenceRecommendationsJobSteps, "list-inference-recommendations-job-steps", "", false, "List Inference Recommendations Job Steps")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListInferenceRecommendationsJobs, "list-inference-recommendations-jobs", "", false, "List Inference Recommendations Jobs")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListLabelingJobs, "list-labeling-jobs", "", false, "List Labeling Jobs")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListLabelingJobsForWorkteam, "list-labeling-jobs-for-workteam", "", false, "List Labeling Jobs For Workteam")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListLineageGroups, "list-lineage-groups", "", false, "List Lineage Groups")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListMlflowApps, "list-mlflow-apps", "", false, "List Mlflow Apps")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListMlflowTrackingServers, "list-mlflow-tracking-servers", "", false, "List Mlflow Tracking Servers")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListModelBiasJobDefinitions, "list-model-bias-job-definitions", "", false, "List Model Bias Job Definitions")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListModelCardExportJobs, "list-model-card-export-jobs", "", false, "List Model Card Export Jobs")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListModelCardVersions, "list-model-card-versions", "", false, "List Model Card Versions")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListModelCards, "list-model-cards", "", false, "List Model Cards")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListModelExplainabilityJobDefinitions, "list-model-explainability-job-definitions", "", false, "List Model Explainability Job Definitions")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListModelMetadata, "list-model-metadata", "", false, "List Model Metadata")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListModelPackageGroups, "list-model-package-groups", "", false, "List Model Package Groups")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListModelPackages, "list-model-packages", "", false, "List Model Packages")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListModelQualityJobDefinitions, "list-model-quality-job-definitions", "", false, "List Model Quality Job Definitions")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListModels, "list-models", "", false, "List Models")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListMonitoringAlertHistory, "list-monitoring-alert-history", "", false, "List Monitoring Alert History")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListMonitoringAlerts, "list-monitoring-alerts", "", false, "List Monitoring Alerts")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListMonitoringExecutions, "list-monitoring-executions", "", false, "List Monitoring Executions")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListMonitoringSchedules, "list-monitoring-schedules", "", false, "List Monitoring Schedules")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListNotebookInstanceLifecycleConfigs, "list-notebook-instance-lifecycle-configs", "", false, "List Notebook Instance Lifecycle Configs")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListNotebookInstances, "list-notebook-instances", "", false, "List Notebook Instances")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListOptimizationJobs, "list-optimization-jobs", "", false, "List Optimization Jobs")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListPartnerApps, "list-partner-apps", "", false, "List Partner Apps")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListPipelineExecutionSteps, "list-pipeline-execution-steps", "", false, "List Pipeline Execution Steps")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListPipelineExecutions, "list-pipeline-executions", "", false, "List Pipeline Executions")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListPipelineParametersForExecution, "list-pipeline-parameters-for-execution", "", false, "List Pipeline Parameters For Execution")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListPipelineVersions, "list-pipeline-versions", "", false, "List Pipeline Versions")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListPipelines, "list-pipelines", "", false, "List Pipelines")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListProcessingJobs, "list-processing-jobs", "", false, "List Processing Jobs")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListProjects, "list-projects", "", false, "List Projects")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListResourceCatalogs, "list-resource-catalogs", "", false, "List Resource Catalogs")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListSpaces, "list-spaces", "", false, "List Spaces")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListStageDevices, "list-stage-devices", "", false, "List Stage Devices")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListStudioLifecycleConfigs, "list-studio-lifecycle-configs", "", false, "List Studio Lifecycle Configs")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListSubscribedWorkteams, "list-subscribed-workteams", "", false, "List Subscribed Workteams")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListTags, "list-tags", "", false, "List Tags")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListTrainingJobs, "list-training-jobs", "", false, "List Training Jobs")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListTrainingJobsForHyperParameterTuningJob, "list-training-jobs-for-hyper-parameter-tuning-job", "", false, "List Training Jobs For Hyper Parameter Tuning Job")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListTrainingPlans, "list-training-plans", "", false, "List Training Plans")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListTransformJobs, "list-transform-jobs", "", false, "List Transform Jobs")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListTrialComponents, "list-trial-components", "", false, "List Trial Components")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListTrials, "list-trials", "", false, "List Trials")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListUltraServersByReservedCapacity, "list-ultra-servers-by-reserved-capacity", "", false, "List Ultra Servers By Reserved Capacity")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListUserProfiles, "list-user-profiles", "", false, "List User Profiles")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListWorkforces, "list-workforces", "", false, "List Workforces")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerListWorkteams, "list-workteams", "", false, "List Workteams")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerPutModelPackageGroupPolicy, "put-model-package-group-policy", "", false, "Put Model Package Group Policy")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerQueryLineage, "query-lineage", "", false, "Query Lineage")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerRegisterDevices, "register-devices", "", false, "Register Devices")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerRenderUiTemplate, "render-ui-template", "", false, "Render Ui Template")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerRetryPipelineExecution, "retry-pipeline-execution", "", false, "Retry Pipeline Execution")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerSearch, "search", "", false, "Search")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerSearchTrainingPlanOfferings, "search-training-plan-offerings", "", false, "Search Training Plan Offerings")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerSendPipelineExecutionStepFailure, "send-pipeline-execution-step-failure", "", false, "Send Pipeline Execution Step Failure")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerSendPipelineExecutionStepSuccess, "send-pipeline-execution-step-success", "", false, "Send Pipeline Execution Step Success")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerStartEdgeDeploymentStage, "start-edge-deployment-stage", "", false, "Start Edge Deployment Stage")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerStartInferenceExperiment, "start-inference-experiment", "", false, "Start Inference Experiment")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerStartMlflowTrackingServer, "start-mlflow-tracking-server", "", false, "Start Mlflow Tracking Server")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerStartMonitoringSchedule, "start-monitoring-schedule", "", false, "Start Monitoring Schedule")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerStartNotebookInstance, "start-notebook-instance", "", false, "Start Notebook Instance")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerStartPipelineExecution, "start-pipeline-execution", "", false, "Start Pipeline Execution")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerStartSession, "start-session", "", false, "Start Session")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerStopAutoMLJob, "stop-auto-ml-job", "", false, "Stop Auto Ml Job")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerStopCompilationJob, "stop-compilation-job", "", false, "Stop Compilation Job")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerStopEdgeDeploymentStage, "stop-edge-deployment-stage", "", false, "Stop Edge Deployment Stage")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerStopEdgePackagingJob, "stop-edge-packaging-job", "", false, "Stop Edge Packaging Job")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerStopHyperParameterTuningJob, "stop-hyper-parameter-tuning-job", "", false, "Stop Hyper Parameter Tuning Job")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerStopInferenceExperiment, "stop-inference-experiment", "", false, "Stop Inference Experiment")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerStopInferenceRecommendationsJob, "stop-inference-recommendations-job", "", false, "Stop Inference Recommendations Job")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerStopLabelingJob, "stop-labeling-job", "", false, "Stop Labeling Job")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerStopMlflowTrackingServer, "stop-mlflow-tracking-server", "", false, "Stop Mlflow Tracking Server")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerStopMonitoringSchedule, "stop-monitoring-schedule", "", false, "Stop Monitoring Schedule")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerStopNotebookInstance, "stop-notebook-instance", "", false, "Stop Notebook Instance")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerStopOptimizationJob, "stop-optimization-job", "", false, "Stop Optimization Job")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerStopPipelineExecution, "stop-pipeline-execution", "", false, "Stop Pipeline Execution")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerStopProcessingJob, "stop-processing-job", "", false, "Stop Processing Job")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerStopTrainingJob, "stop-training-job", "", false, "Stop Training Job")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerStopTransformJob, "stop-transform-job", "", false, "Stop Transform Job")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerUpdateAction, "update-action", "", false, "Update Action")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerUpdateAppImageConfig, "update-app-image-config", "", false, "Update App Image Config")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerUpdateArtifact, "update-artifact", "", false, "Update Artifact")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerUpdateCluster, "update-cluster", "", false, "Update Cluster")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerUpdateClusterSchedulerConfig, "update-cluster-scheduler-config", "", false, "Update Cluster Scheduler Config")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerUpdateClusterSoftware, "update-cluster-software", "", false, "Update Cluster Software")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerUpdateCodeRepository, "update-code-repository", "", false, "Update Code Repository")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerUpdateComputeQuota, "update-compute-quota", "", false, "Update Compute Quota")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerUpdateContext, "update-context", "", false, "Update Context")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerUpdateDeviceFleet, "update-device-fleet", "", false, "Update Device Fleet")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerUpdateDevices, "update-devices", "", false, "Update Devices")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerUpdateDomain, "update-domain", "", false, "Update Domain")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerUpdateEndpoint, "update-endpoint", "", false, "Update Endpoint")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerUpdateEndpointWeightsAndCapacities, "update-endpoint-weights-and-capacities", "", false, "Update Endpoint Weights And Capacities")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerUpdateExperiment, "update-experiment", "", false, "Update Experiment")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerUpdateFeatureGroup, "update-feature-group", "", false, "Update Feature Group")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerUpdateFeatureMetadata, "update-feature-metadata", "", false, "Update Feature Metadata")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerUpdateHub, "update-hub", "", false, "Update Hub")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerUpdateHubContent, "update-hub-content", "", false, "Update Hub Content")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerUpdateHubContentReference, "update-hub-content-reference", "", false, "Update Hub Content Reference")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerUpdateImage, "update-image", "", false, "Update Image")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerUpdateImageVersion, "update-image-version", "", false, "Update Image Version")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerUpdateInferenceComponent, "update-inference-component", "", false, "Update Inference Component")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerUpdateInferenceComponentRuntimeConfig, "update-inference-component-runtime-config", "", false, "Update Inference Component Runtime Config")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerUpdateInferenceExperiment, "update-inference-experiment", "", false, "Update Inference Experiment")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerUpdateMlflowApp, "update-mlflow-app", "", false, "Update Mlflow App")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerUpdateMlflowTrackingServer, "update-mlflow-tracking-server", "", false, "Update Mlflow Tracking Server")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerUpdateModelCard, "update-model-card", "", false, "Update Model Card")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerUpdateModelPackage, "update-model-package", "", false, "Update Model Package")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerUpdateMonitoringAlert, "update-monitoring-alert", "", false, "Update Monitoring Alert")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerUpdateMonitoringSchedule, "update-monitoring-schedule", "", false, "Update Monitoring Schedule")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerUpdateNotebookInstance, "update-notebook-instance", "", false, "Update Notebook Instance")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerUpdateNotebookInstanceLifecycleConfig, "update-notebook-instance-lifecycle-config", "", false, "Update Notebook Instance Lifecycle Config")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerUpdatePartnerApp, "update-partner-app", "", false, "Update Partner App")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerUpdatePipeline, "update-pipeline", "", false, "Update Pipeline")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerUpdatePipelineExecution, "update-pipeline-execution", "", false, "Update Pipeline Execution")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerUpdatePipelineVersion, "update-pipeline-version", "", false, "Update Pipeline Version")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerUpdateProject, "update-project", "", false, "Update Project")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerUpdateSpace, "update-space", "", false, "Update Space")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerUpdateTrainingJob, "update-training-job", "", false, "Update Training Job")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerUpdateTrial, "update-trial", "", false, "Update Trial")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerUpdateTrialComponent, "update-trial-component", "", false, "Update Trial Component")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerUpdateUserProfile, "update-user-profile", "", false, "Update User Profile")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerUpdateWorkforce, "update-workforce", "", false, "Update Workforce")
	_sagemakerCmd.Flags().BoolVarP(&_sagemakerUpdateWorkteam, "update-workteam", "", false, "Update Workteam")

}
