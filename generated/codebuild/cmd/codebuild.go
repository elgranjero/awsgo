package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codebuild"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// codebuildCmd represents the codebuild command
var _codebuildCmd = &cobra.Command{
	Use:   "codebuild",
	Short: "AWS codebuild CLI",
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
		client := codebuild.NewFromConfig(cfg)
		if _codebuildBatchDeleteBuilds {
			codebuild_BatchDeleteBuilds(cfg, client)
			return
		}
		if _codebuildBatchGetBuildBatches {
			codebuild_BatchGetBuildBatches(cfg, client)
			return
		}
		if _codebuildBatchGetBuilds {
			codebuild_BatchGetBuilds(cfg, client)
			return
		}
		if _codebuildBatchGetCommandExecutions {
			codebuild_BatchGetCommandExecutions(cfg, client)
			return
		}
		if _codebuildBatchGetFleets {
			codebuild_BatchGetFleets(cfg, client)
			return
		}
		if _codebuildBatchGetProjects {
			codebuild_BatchGetProjects(cfg, client)
			return
		}
		if _codebuildBatchGetReportGroups {
			codebuild_BatchGetReportGroups(cfg, client)
			return
		}
		if _codebuildBatchGetReports {
			codebuild_BatchGetReports(cfg, client)
			return
		}
		if _codebuildBatchGetSandboxes {
			codebuild_BatchGetSandboxes(cfg, client)
			return
		}
		if _codebuildCreateFleet {
			codebuild_CreateFleet(cfg, client)
			return
		}
		if _codebuildCreateProject {
			codebuild_CreateProject(cfg, client)
			return
		}
		if _codebuildCreateReportGroup {
			codebuild_CreateReportGroup(cfg, client)
			return
		}
		if _codebuildCreateWebhook {
			codebuild_CreateWebhook(cfg, client)
			return
		}
		if _codebuildDeleteBuildBatch {
			codebuild_DeleteBuildBatch(cfg, client)
			return
		}
		if _codebuildDeleteFleet {
			codebuild_DeleteFleet(cfg, client)
			return
		}
		if _codebuildDeleteProject {
			codebuild_DeleteProject(cfg, client)
			return
		}
		if _codebuildDeleteReport {
			codebuild_DeleteReport(cfg, client)
			return
		}
		if _codebuildDeleteReportGroup {
			codebuild_DeleteReportGroup(cfg, client)
			return
		}
		if _codebuildDeleteResourcePolicy {
			codebuild_DeleteResourcePolicy(cfg, client)
			return
		}
		if _codebuildDeleteSourceCredentials {
			codebuild_DeleteSourceCredentials(cfg, client)
			return
		}
		if _codebuildDeleteWebhook {
			codebuild_DeleteWebhook(cfg, client)
			return
		}
		if _codebuildDescribeCodeCoverages {
			codebuild_DescribeCodeCoverages(cfg, client)
			return
		}
		if _codebuildDescribeTestCases {
			codebuild_DescribeTestCases(cfg, client)
			return
		}
		if _codebuildGetReportGroupTrend {
			codebuild_GetReportGroupTrend(cfg, client)
			return
		}
		if _codebuildGetResourcePolicy {
			codebuild_GetResourcePolicy(cfg, client)
			return
		}
		if _codebuildImportSourceCredentials {
			codebuild_ImportSourceCredentials(cfg, client)
			return
		}
		if _codebuildInvalidateProjectCache {
			codebuild_InvalidateProjectCache(cfg, client)
			return
		}
		if _codebuildListBuildBatches {
			codebuild_ListBuildBatches(cfg, client)
			return
		}
		if _codebuildListBuildBatchesForProject {
			codebuild_ListBuildBatchesForProject(cfg, client)
			return
		}
		if _codebuildListBuilds {
			codebuild_ListBuilds(cfg, client)
			return
		}
		if _codebuildListBuildsForProject {
			codebuild_ListBuildsForProject(cfg, client)
			return
		}
		if _codebuildListCommandExecutionsForSandbox {
			codebuild_ListCommandExecutionsForSandbox(cfg, client)
			return
		}
		if _codebuildListCuratedEnvironmentImages {
			codebuild_ListCuratedEnvironmentImages(cfg, client)
			return
		}
		if _codebuildListFleets {
			codebuild_ListFleets(cfg, client)
			return
		}
		if _codebuildListProjects {
			codebuild_ListProjects(cfg, client)
			return
		}
		if _codebuildListReportGroups {
			codebuild_ListReportGroups(cfg, client)
			return
		}
		if _codebuildListReports {
			codebuild_ListReports(cfg, client)
			return
		}
		if _codebuildListReportsForReportGroup {
			codebuild_ListReportsForReportGroup(cfg, client)
			return
		}
		if _codebuildListSandboxes {
			codebuild_ListSandboxes(cfg, client)
			return
		}
		if _codebuildListSandboxesForProject {
			codebuild_ListSandboxesForProject(cfg, client)
			return
		}
		if _codebuildListSharedProjects {
			codebuild_ListSharedProjects(cfg, client)
			return
		}
		if _codebuildListSharedReportGroups {
			codebuild_ListSharedReportGroups(cfg, client)
			return
		}
		if _codebuildListSourceCredentials {
			codebuild_ListSourceCredentials(cfg, client)
			return
		}
		if _codebuildPutResourcePolicy {
			codebuild_PutResourcePolicy(cfg, client)
			return
		}
		if _codebuildRetryBuild {
			codebuild_RetryBuild(cfg, client)
			return
		}
		if _codebuildRetryBuildBatch {
			codebuild_RetryBuildBatch(cfg, client)
			return
		}
		if _codebuildStartBuild {
			codebuild_StartBuild(cfg, client)
			return
		}
		if _codebuildStartBuildBatch {
			codebuild_StartBuildBatch(cfg, client)
			return
		}
		if _codebuildStartCommandExecution {
			codebuild_StartCommandExecution(cfg, client)
			return
		}
		if _codebuildStartSandbox {
			codebuild_StartSandbox(cfg, client)
			return
		}
		if _codebuildStartSandboxConnection {
			codebuild_StartSandboxConnection(cfg, client)
			return
		}
		if _codebuildStopBuild {
			codebuild_StopBuild(cfg, client)
			return
		}
		if _codebuildStopBuildBatch {
			codebuild_StopBuildBatch(cfg, client)
			return
		}
		if _codebuildStopSandbox {
			codebuild_StopSandbox(cfg, client)
			return
		}
		if _codebuildUpdateFleet {
			codebuild_UpdateFleet(cfg, client)
			return
		}
		if _codebuildUpdateProject {
			codebuild_UpdateProject(cfg, client)
			return
		}
		if _codebuildUpdateProjectVisibility {
			codebuild_UpdateProjectVisibility(cfg, client)
			return
		}
		if _codebuildUpdateReportGroup {
			codebuild_UpdateReportGroup(cfg, client)
			return
		}
		if _codebuildUpdateWebhook {
			codebuild_UpdateWebhook(cfg, client)
			return
		}

	},
}

var (
	_codebuildBatchDeleteBuilds               bool
	_codebuildBatchGetBuildBatches            bool
	_codebuildBatchGetBuilds                  bool
	_codebuildBatchGetCommandExecutions       bool
	_codebuildBatchGetFleets                  bool
	_codebuildBatchGetProjects                bool
	_codebuildBatchGetReportGroups            bool
	_codebuildBatchGetReports                 bool
	_codebuildBatchGetSandboxes               bool
	_codebuildCreateFleet                     bool
	_codebuildCreateProject                   bool
	_codebuildCreateReportGroup               bool
	_codebuildCreateWebhook                   bool
	_codebuildDeleteBuildBatch                bool
	_codebuildDeleteFleet                     bool
	_codebuildDeleteProject                   bool
	_codebuildDeleteReport                    bool
	_codebuildDeleteReportGroup               bool
	_codebuildDeleteResourcePolicy            bool
	_codebuildDeleteSourceCredentials         bool
	_codebuildDeleteWebhook                   bool
	_codebuildDescribeCodeCoverages           bool
	_codebuildDescribeTestCases               bool
	_codebuildGetReportGroupTrend             bool
	_codebuildGetResourcePolicy               bool
	_codebuildImportSourceCredentials         bool
	_codebuildInvalidateProjectCache          bool
	_codebuildListBuildBatches                bool
	_codebuildListBuildBatchesForProject      bool
	_codebuildListBuilds                      bool
	_codebuildListBuildsForProject            bool
	_codebuildListCommandExecutionsForSandbox bool
	_codebuildListCuratedEnvironmentImages    bool
	_codebuildListFleets                      bool
	_codebuildListProjects                    bool
	_codebuildListReportGroups                bool
	_codebuildListReports                     bool
	_codebuildListReportsForReportGroup       bool
	_codebuildListSandboxes                   bool
	_codebuildListSandboxesForProject         bool
	_codebuildListSharedProjects              bool
	_codebuildListSharedReportGroups          bool
	_codebuildListSourceCredentials           bool
	_codebuildPutResourcePolicy               bool
	_codebuildRetryBuild                      bool
	_codebuildRetryBuildBatch                 bool
	_codebuildStartBuild                      bool
	_codebuildStartBuildBatch                 bool
	_codebuildStartCommandExecution           bool
	_codebuildStartSandbox                    bool
	_codebuildStartSandboxConnection          bool
	_codebuildStopBuild                       bool
	_codebuildStopBuildBatch                  bool
	_codebuildStopSandbox                     bool
	_codebuildUpdateFleet                     bool
	_codebuildUpdateProject                   bool
	_codebuildUpdateProjectVisibility         bool
	_codebuildUpdateReportGroup               bool
	_codebuildUpdateWebhook                   bool

	_codebuildArn                              string
	_codebuildArtifacts                        string
	_codebuildArtifactsOverride                string
	_codebuildAuthType                         string
	_codebuildAutoRetryLimit                   string
	_codebuildAutoRetryLimitOverride           string
	_codebuildBadgeEnabled                     string
	_codebuildBaseCapacity                     string
	_codebuildBranchFilter                     string
	_codebuildBuildBatchConfig                 string
	_codebuildBuildBatchConfigOverride         string
	_codebuildBuildStatusConfigOverride        string
	_codebuildBuildTimeoutInMinutesOverride    string
	_codebuildBuildType                        string
	_codebuildBuildspecOverride                string
	_codebuildCache                            string
	_codebuildCacheOverride                    string
	_codebuildCertificateOverride              string
	_codebuildCommand                          string
	_codebuildCommandExecutionIds              []string
	_codebuildComputeConfiguration             string
	_codebuildComputeType                      string
	_codebuildComputeTypeOverride              string
	_codebuildConcurrentBuildLimit             string
	_codebuildDebugSessionEnabled              string
	_codebuildDeleteReports                    string
	_codebuildDescription                      string
	_codebuildEncryptionKey                    string
	_codebuildEncryptionKeyOverride            string
	_codebuildEnvironment                      string
	_codebuildEnvironmentType                  string
	_codebuildEnvironmentTypeOverride          string
	_codebuildEnvironmentVariablesOverride     string
	_codebuildExportConfig                     string
	_codebuildFileSystemLocations              string
	_codebuildFilter                           string
	_codebuildFilterGroups                     string
	_codebuildFleetOverride                    string
	_codebuildFleetServiceRole                 string
	_codebuildGitCloneDepthOverride            string
	_codebuildGitSubmodulesConfigOverride      string
	_codebuildId                               string
	_codebuildIdempotencyToken                 string
	_codebuildIds                              []string
	_codebuildImageId                          string
	_codebuildImageOverride                    string
	_codebuildImagePullCredentialsTypeOverride string
	_codebuildInsecureSslOverride              string
	_codebuildLogsConfig                       string
	_codebuildLogsConfigOverride               string
	_codebuildManualCreation                   string
	_codebuildMaxLineCoveragePercentage        string
	_codebuildMaxResults                       string
	_codebuildMinLineCoveragePercentage        string
	_codebuildName                             string
	_codebuildNames                            []string
	_codebuildNextToken                        string
	_codebuildNumOfReports                     string
	_codebuildOverflowBehavior                 string
	_codebuildPolicy                           string
	_codebuildPrivilegedModeOverride           string
	_codebuildProjectArn                       string
	_codebuildProjectName                      string
	_codebuildProjectVisibility                string
	_codebuildProxyConfiguration               string
	_codebuildPullRequestBuildPolicy           string
	_codebuildQueuedTimeoutInMinutes           string
	_codebuildQueuedTimeoutInMinutesOverride   string
	_codebuildRegistryCredentialOverride       string
	_codebuildReportArn                        string
	_codebuildReportArns                       []string
	_codebuildReportBuildBatchStatusOverride   string
	_codebuildReportBuildStatusOverride        string
	_codebuildReportGroupArn                   string
	_codebuildReportGroupArns                  []string
	_codebuildResourceAccessRole               string
	_codebuildResourceArn                      string
	_codebuildRetryType                        string
	_codebuildRotateSecret                     string
	_codebuildSandboxId                        string
	_codebuildScalingConfiguration             string
	_codebuildScopeConfiguration               string
	_codebuildSecondaryArtifacts               string
	_codebuildSecondaryArtifactsOverride       string
	_codebuildSecondarySourceVersions          string
	_codebuildSecondarySources                 string
	_codebuildSecondarySourcesOverride         string
	_codebuildSecondarySourcesVersionOverride  string
	_codebuildServerType                       string
	_codebuildServiceRole                      string
	_codebuildServiceRoleOverride              string
	_codebuildShouldOverwrite                  string
	_codebuildSortBy                           string
	_codebuildSortOrder                        string
	_codebuildSource                           string
	_codebuildSourceAuthOverride               string
	_codebuildSourceLocationOverride           string
	_codebuildSourceTypeOverride               string
	_codebuildSourceVersion                    string
	_codebuildTags                             string
	_codebuildTimeoutInMinutes                 string
	_codebuildTimeoutInMinutesOverride         string
	_codebuildToken                            string
	_codebuildTrendField                       string
	_codebuildType                             string
	_codebuildUsername                         string
	_codebuildVpcConfig                        string
)

// Deletes one or more builds.
func codebuild_BatchDeleteBuilds(cfg aws.Config, client *codebuild.Client) {
	input := &codebuild.BatchDeleteBuildsInput{
		// Ids: []string, // Required
	}

	if len(_codebuildIds) > 0 {
		input.Ids = append([]string(nil), _codebuildIds...)
	}

	if resp, err := client.BatchDeleteBuilds(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about one or more batch builds.
func codebuild_BatchGetBuildBatches(cfg aws.Config, client *codebuild.Client) {
	input := &codebuild.BatchGetBuildBatchesInput{
		// Ids: []string, // Required
	}

	if len(_codebuildIds) > 0 {
		input.Ids = append([]string(nil), _codebuildIds...)
	}

	if resp, err := client.BatchGetBuildBatches(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about one or more builds.
func codebuild_BatchGetBuilds(cfg aws.Config, client *codebuild.Client) {
	input := &codebuild.BatchGetBuildsInput{
		// Ids: []string, // Required
	}

	if len(_codebuildIds) > 0 {
		input.Ids = append([]string(nil), _codebuildIds...)
	}

	if resp, err := client.BatchGetBuilds(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about the command executions.
func codebuild_BatchGetCommandExecutions(cfg aws.Config, client *codebuild.Client) {
	input := &codebuild.BatchGetCommandExecutionsInput{
		// CommandExecutionIds: []string, // Required
		// SandboxId: *string, // Required
	}

	if len(_codebuildCommandExecutionIds) > 0 {
		input.CommandExecutionIds = append([]string(nil), _codebuildCommandExecutionIds...)
	}
	if len(_codebuildSandboxId) > 0 {
		input.SandboxId = aws.String(_codebuildSandboxId)
	}

	if resp, err := client.BatchGetCommandExecutions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about one or more compute fleets.
func codebuild_BatchGetFleets(cfg aws.Config, client *codebuild.Client) {
	input := &codebuild.BatchGetFleetsInput{
		// Names: []string, // Required
	}

	if len(_codebuildNames) > 0 {
		input.Names = append([]string(nil), _codebuildNames...)
	}

	if resp, err := client.BatchGetFleets(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about one or more build projects.
func codebuild_BatchGetProjects(cfg aws.Config, client *codebuild.Client) {
	input := &codebuild.BatchGetProjectsInput{
		// Names: []string, // Required
	}

	if len(_codebuildNames) > 0 {
		input.Names = append([]string(nil), _codebuildNames...)
	}

	if resp, err := client.BatchGetProjects(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns an array of report groups.
func codebuild_BatchGetReportGroups(cfg aws.Config, client *codebuild.Client) {
	input := &codebuild.BatchGetReportGroupsInput{
		// ReportGroupArns: []string, // Required
	}

	if len(_codebuildReportGroupArns) > 0 {
		input.ReportGroupArns = append([]string(nil), _codebuildReportGroupArns...)
	}

	if resp, err := client.BatchGetReportGroups(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns an array of reports.
func codebuild_BatchGetReports(cfg aws.Config, client *codebuild.Client) {
	input := &codebuild.BatchGetReportsInput{
		// ReportArns: []string, // Required
	}

	if len(_codebuildReportArns) > 0 {
		input.ReportArns = append([]string(nil), _codebuildReportArns...)
	}

	if resp, err := client.BatchGetReports(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about the sandbox status.
func codebuild_BatchGetSandboxes(cfg aws.Config, client *codebuild.Client) {
	input := &codebuild.BatchGetSandboxesInput{
		// Ids: []string, // Required
	}

	if len(_codebuildIds) > 0 {
		input.Ids = append([]string(nil), _codebuildIds...)
	}

	if resp, err := client.BatchGetSandboxes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a compute fleet.
func codebuild_CreateFleet(cfg aws.Config, client *codebuild.Client) {
	input := &codebuild.CreateFleetInput{
		// BaseCapacity: *int32, // Required
		// ComputeType: types.ComputeType, // Required
		// EnvironmentType: types.EnvironmentType, // Required
		// Name: *string, // Required
	}

	if len(_codebuildBaseCapacity) > 0 {
		if err := assignInputField(input, "BaseCapacity", _codebuildBaseCapacity); err != nil {
			log.Errorf("invalid --base-capacity: %s", err.Error())
			return
		}
	}
	if len(_codebuildComputeType) > 0 {
		if err := assignInputField(input, "ComputeType", _codebuildComputeType); err != nil {
			log.Errorf("invalid --compute-type: %s", err.Error())
			return
		}
	}
	if len(_codebuildEnvironmentType) > 0 {
		if err := assignInputField(input, "EnvironmentType", _codebuildEnvironmentType); err != nil {
			log.Errorf("invalid --environment-type: %s", err.Error())
			return
		}
	}
	if len(_codebuildName) > 0 {
		input.Name = aws.String(_codebuildName)
	}
	if len(_codebuildComputeConfiguration) > 0 {
		if err := assignInputField(input, "ComputeConfiguration", _codebuildComputeConfiguration); err != nil {
			log.Errorf("invalid --compute-configuration: %s", err.Error())
			return
		}
	}
	if len(_codebuildFleetServiceRole) > 0 {
		input.FleetServiceRole = aws.String(_codebuildFleetServiceRole)
	}
	if len(_codebuildImageId) > 0 {
		input.ImageId = aws.String(_codebuildImageId)
	}
	if len(_codebuildOverflowBehavior) > 0 {
		if err := assignInputField(input, "OverflowBehavior", _codebuildOverflowBehavior); err != nil {
			log.Errorf("invalid --overflow-behavior: %s", err.Error())
			return
		}
	}
	if len(_codebuildProxyConfiguration) > 0 {
		if err := assignInputField(input, "ProxyConfiguration", _codebuildProxyConfiguration); err != nil {
			log.Errorf("invalid --proxy-configuration: %s", err.Error())
			return
		}
	}
	if len(_codebuildScalingConfiguration) > 0 {
		if err := assignInputField(input, "ScalingConfiguration", _codebuildScalingConfiguration); err != nil {
			log.Errorf("invalid --scaling-configuration: %s", err.Error())
			return
		}
	}
	if len(_codebuildTags) > 0 {
		if err := assignInputField(input, "Tags", _codebuildTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_codebuildVpcConfig) > 0 {
		if err := assignInputField(input, "VpcConfig", _codebuildVpcConfig); err != nil {
			log.Errorf("invalid --vpc-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateFleet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a build project.
func codebuild_CreateProject(cfg aws.Config, client *codebuild.Client) {
	input := &codebuild.CreateProjectInput{
		// Artifacts: *types.ProjectArtifacts, // Required
		// Environment: *types.ProjectEnvironment, // Required
		// Name: *string, // Required
		// ServiceRole: *string, // Required
		// Source: *types.ProjectSource, // Required
	}

	if len(_codebuildArtifacts) > 0 {
		if err := assignInputField(input, "Artifacts", _codebuildArtifacts); err != nil {
			log.Errorf("invalid --artifacts: %s", err.Error())
			return
		}
	}
	if len(_codebuildEnvironment) > 0 {
		if err := assignInputField(input, "Environment", _codebuildEnvironment); err != nil {
			log.Errorf("invalid --environment: %s", err.Error())
			return
		}
	}
	if len(_codebuildName) > 0 {
		input.Name = aws.String(_codebuildName)
	}
	if len(_codebuildServiceRole) > 0 {
		input.ServiceRole = aws.String(_codebuildServiceRole)
	}
	if len(_codebuildSource) > 0 {
		if err := assignInputField(input, "Source", _codebuildSource); err != nil {
			log.Errorf("invalid --source: %s", err.Error())
			return
		}
	}
	if len(_codebuildAutoRetryLimit) > 0 {
		if err := assignInputField(input, "AutoRetryLimit", _codebuildAutoRetryLimit); err != nil {
			log.Errorf("invalid --auto-retry-limit: %s", err.Error())
			return
		}
	}
	if len(_codebuildBadgeEnabled) > 0 {
		if err := assignInputField(input, "BadgeEnabled", _codebuildBadgeEnabled); err != nil {
			log.Errorf("invalid --badge-enabled: %s", err.Error())
			return
		}
	}
	if len(_codebuildBuildBatchConfig) > 0 {
		if err := assignInputField(input, "BuildBatchConfig", _codebuildBuildBatchConfig); err != nil {
			log.Errorf("invalid --build-batch-config: %s", err.Error())
			return
		}
	}
	if len(_codebuildCache) > 0 {
		if err := assignInputField(input, "Cache", _codebuildCache); err != nil {
			log.Errorf("invalid --cache: %s", err.Error())
			return
		}
	}
	if len(_codebuildConcurrentBuildLimit) > 0 {
		if err := assignInputField(input, "ConcurrentBuildLimit", _codebuildConcurrentBuildLimit); err != nil {
			log.Errorf("invalid --concurrent-build-limit: %s", err.Error())
			return
		}
	}
	if len(_codebuildDescription) > 0 {
		input.Description = aws.String(_codebuildDescription)
	}
	if len(_codebuildEncryptionKey) > 0 {
		input.EncryptionKey = aws.String(_codebuildEncryptionKey)
	}
	if len(_codebuildFileSystemLocations) > 0 {
		if err := assignInputField(input, "FileSystemLocations", _codebuildFileSystemLocations); err != nil {
			log.Errorf("invalid --file-system-locations: %s", err.Error())
			return
		}
	}
	if len(_codebuildLogsConfig) > 0 {
		if err := assignInputField(input, "LogsConfig", _codebuildLogsConfig); err != nil {
			log.Errorf("invalid --logs-config: %s", err.Error())
			return
		}
	}
	if len(_codebuildQueuedTimeoutInMinutes) > 0 {
		if err := assignInputField(input, "QueuedTimeoutInMinutes", _codebuildQueuedTimeoutInMinutes); err != nil {
			log.Errorf("invalid --queued-timeout-in-minutes: %s", err.Error())
			return
		}
	}
	if len(_codebuildSecondaryArtifacts) > 0 {
		if err := assignInputField(input, "SecondaryArtifacts", _codebuildSecondaryArtifacts); err != nil {
			log.Errorf("invalid --secondary-artifacts: %s", err.Error())
			return
		}
	}
	if len(_codebuildSecondarySourceVersions) > 0 {
		if err := assignInputField(input, "SecondarySourceVersions", _codebuildSecondarySourceVersions); err != nil {
			log.Errorf("invalid --secondary-source-versions: %s", err.Error())
			return
		}
	}
	if len(_codebuildSecondarySources) > 0 {
		if err := assignInputField(input, "SecondarySources", _codebuildSecondarySources); err != nil {
			log.Errorf("invalid --secondary-sources: %s", err.Error())
			return
		}
	}
	if len(_codebuildSourceVersion) > 0 {
		input.SourceVersion = aws.String(_codebuildSourceVersion)
	}
	if len(_codebuildTags) > 0 {
		if err := assignInputField(input, "Tags", _codebuildTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_codebuildTimeoutInMinutes) > 0 {
		if err := assignInputField(input, "TimeoutInMinutes", _codebuildTimeoutInMinutes); err != nil {
			log.Errorf("invalid --timeout-in-minutes: %s", err.Error())
			return
		}
	}
	if len(_codebuildVpcConfig) > 0 {
		if err := assignInputField(input, "VpcConfig", _codebuildVpcConfig); err != nil {
			log.Errorf("invalid --vpc-config: %s", err.Error())
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

// Creates a report group. A report group contains a collection of reports.
func codebuild_CreateReportGroup(cfg aws.Config, client *codebuild.Client) {
	input := &codebuild.CreateReportGroupInput{
		// ExportConfig: *types.ReportExportConfig, // Required
		// Name: *string, // Required
		// Type: types.ReportType, // Required
	}

	if len(_codebuildExportConfig) > 0 {
		if err := assignInputField(input, "ExportConfig", _codebuildExportConfig); err != nil {
			log.Errorf("invalid --export-config: %s", err.Error())
			return
		}
	}
	if len(_codebuildName) > 0 {
		input.Name = aws.String(_codebuildName)
	}
	if len(_codebuildType) > 0 {
		if err := assignInputField(input, "Type", _codebuildType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_codebuildTags) > 0 {
		if err := assignInputField(input, "Tags", _codebuildTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateReportGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// For an existing CodeBuild build project that has its source code stored in a
// GitHub or Bitbucket repository, enables CodeBuild to start rebuilding the source
// code every time a code change is pushed to the repository.
//
// If you enable webhooks for an CodeBuild project, and the project is used as a
// build step in CodePipeline, then two identical builds are created for each
// commit. One build is triggered through webhooks, and one through CodePipeline.
// Because billing is on a per-build basis, you are billed for both builds.
// Therefore, if you are using CodePipeline, we recommend that you disable webhooks
// in CodeBuild. In the CodeBuild console, clear the Webhook box. For more
// information, see step 5 in [Change a Build Project's Settings].
//
// [Change a Build Project's Settings]: https://docs.aws.amazon.com/codebuild/latest/userguide/change-project.html#change-project-console
func codebuild_CreateWebhook(cfg aws.Config, client *codebuild.Client) {
	input := &codebuild.CreateWebhookInput{
		// ProjectName: *string, // Required
	}

	if len(_codebuildProjectName) > 0 {
		input.ProjectName = aws.String(_codebuildProjectName)
	}
	if len(_codebuildBranchFilter) > 0 {
		input.BranchFilter = aws.String(_codebuildBranchFilter)
	}
	if len(_codebuildBuildType) > 0 {
		if err := assignInputField(input, "BuildType", _codebuildBuildType); err != nil {
			log.Errorf("invalid --build-type: %s", err.Error())
			return
		}
	}
	if len(_codebuildFilterGroups) > 0 {
		if err := assignInputField(input, "FilterGroups", _codebuildFilterGroups); err != nil {
			log.Errorf("invalid --filter-groups: %s", err.Error())
			return
		}
	}
	if len(_codebuildManualCreation) > 0 {
		if err := assignInputField(input, "ManualCreation", _codebuildManualCreation); err != nil {
			log.Errorf("invalid --manual-creation: %s", err.Error())
			return
		}
	}
	if len(_codebuildPullRequestBuildPolicy) > 0 {
		if err := assignInputField(input, "PullRequestBuildPolicy", _codebuildPullRequestBuildPolicy); err != nil {
			log.Errorf("invalid --pull-request-build-policy: %s", err.Error())
			return
		}
	}
	if len(_codebuildScopeConfiguration) > 0 {
		if err := assignInputField(input, "ScopeConfiguration", _codebuildScopeConfiguration); err != nil {
			log.Errorf("invalid --scope-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateWebhook(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a batch build.
func codebuild_DeleteBuildBatch(cfg aws.Config, client *codebuild.Client) {
	input := &codebuild.DeleteBuildBatchInput{
		// Id: *string, // Required
	}

	if len(_codebuildId) > 0 {
		input.Id = aws.String(_codebuildId)
	}

	if resp, err := client.DeleteBuildBatch(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a compute fleet. When you delete a compute fleet, its builds are not
// deleted.
func codebuild_DeleteFleet(cfg aws.Config, client *codebuild.Client) {
	input := &codebuild.DeleteFleetInput{
		// Arn: *string, // Required
	}

	if len(_codebuildArn) > 0 {
		input.Arn = aws.String(_codebuildArn)
	}

	if resp, err := client.DeleteFleet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a build project. When you delete a project, its builds are not
// deleted.
func codebuild_DeleteProject(cfg aws.Config, client *codebuild.Client) {
	input := &codebuild.DeleteProjectInput{
		// Name: *string, // Required
	}

	if len(_codebuildName) > 0 {
		input.Name = aws.String(_codebuildName)
	}

	if resp, err := client.DeleteProject(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a report.
func codebuild_DeleteReport(cfg aws.Config, client *codebuild.Client) {
	input := &codebuild.DeleteReportInput{
		// Arn: *string, // Required
	}

	if len(_codebuildArn) > 0 {
		input.Arn = aws.String(_codebuildArn)
	}

	if resp, err := client.DeleteReport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a report group. Before you delete a report group, you must delete its
// reports.
func codebuild_DeleteReportGroup(cfg aws.Config, client *codebuild.Client) {
	input := &codebuild.DeleteReportGroupInput{
		// Arn: *string, // Required
	}

	if len(_codebuildArn) > 0 {
		input.Arn = aws.String(_codebuildArn)
	}
	if len(_codebuildDeleteReports) > 0 {
		if err := assignInputField(input, "DeleteReports", _codebuildDeleteReports); err != nil {
			log.Errorf("invalid --delete-reports: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteReportGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a resource policy that is identified by its resource ARN.
func codebuild_DeleteResourcePolicy(cfg aws.Config, client *codebuild.Client) {
	input := &codebuild.DeleteResourcePolicyInput{
		// ResourceArn: *string, // Required
	}

	if len(_codebuildResourceArn) > 0 {
		input.ResourceArn = aws.String(_codebuildResourceArn)
	}

	if resp, err := client.DeleteResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a set of GitHub, GitHub Enterprise, or Bitbucket source credentials.
func codebuild_DeleteSourceCredentials(cfg aws.Config, client *codebuild.Client) {
	input := &codebuild.DeleteSourceCredentialsInput{
		// Arn: *string, // Required
	}

	if len(_codebuildArn) > 0 {
		input.Arn = aws.String(_codebuildArn)
	}

	if resp, err := client.DeleteSourceCredentials(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// For an existing CodeBuild build project that has its source code stored in a
// GitHub or Bitbucket repository, stops CodeBuild from rebuilding the source code
// every time a code change is pushed to the repository.
func codebuild_DeleteWebhook(cfg aws.Config, client *codebuild.Client) {
	input := &codebuild.DeleteWebhookInput{
		// ProjectName: *string, // Required
	}

	if len(_codebuildProjectName) > 0 {
		input.ProjectName = aws.String(_codebuildProjectName)
	}

	if resp, err := client.DeleteWebhook(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves one or more code coverage reports.
func codebuild_DescribeCodeCoverages(cfg aws.Config, client *codebuild.Client) {
	input := &codebuild.DescribeCodeCoveragesInput{
		// ReportArn: *string, // Required
	}

	if len(_codebuildReportArn) > 0 {
		input.ReportArn = aws.String(_codebuildReportArn)
	}
	if len(_codebuildMaxLineCoveragePercentage) > 0 {
		if err := assignInputField(input, "MaxLineCoveragePercentage", _codebuildMaxLineCoveragePercentage); err != nil {
			log.Errorf("invalid --max-line-coverage-percentage: %s", err.Error())
			return
		}
	}
	if len(_codebuildMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codebuildMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codebuildMinLineCoveragePercentage) > 0 {
		if err := assignInputField(input, "MinLineCoveragePercentage", _codebuildMinLineCoveragePercentage); err != nil {
			log.Errorf("invalid --min-line-coverage-percentage: %s", err.Error())
			return
		}
	}
	if len(_codebuildNextToken) > 0 {
		input.NextToken = aws.String(_codebuildNextToken)
	}
	if len(_codebuildSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _codebuildSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_codebuildSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _codebuildSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeCodeCoverages(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codebuild.DescribeCodeCoveragesOutput
	p := codebuild.NewDescribeCodeCoveragesPaginator(client, input)
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

// Returns a list of details about test cases for a report.
func codebuild_DescribeTestCases(cfg aws.Config, client *codebuild.Client) {
	input := &codebuild.DescribeTestCasesInput{
		// ReportArn: *string, // Required
	}

	if len(_codebuildReportArn) > 0 {
		input.ReportArn = aws.String(_codebuildReportArn)
	}
	if len(_codebuildFilter) > 0 {
		if err := assignInputField(input, "Filter", _codebuildFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_codebuildMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codebuildMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codebuildNextToken) > 0 {
		input.NextToken = aws.String(_codebuildNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeTestCases(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codebuild.DescribeTestCasesOutput
	p := codebuild.NewDescribeTestCasesPaginator(client, input)
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

// Analyzes and accumulates test report values for the specified test reports.
func codebuild_GetReportGroupTrend(cfg aws.Config, client *codebuild.Client) {
	input := &codebuild.GetReportGroupTrendInput{
		// ReportGroupArn: *string, // Required
		// TrendField: types.ReportGroupTrendFieldType, // Required
	}

	if len(_codebuildReportGroupArn) > 0 {
		input.ReportGroupArn = aws.String(_codebuildReportGroupArn)
	}
	if len(_codebuildTrendField) > 0 {
		if err := assignInputField(input, "TrendField", _codebuildTrendField); err != nil {
			log.Errorf("invalid --trend-field: %s", err.Error())
			return
		}
	}
	if len(_codebuildNumOfReports) > 0 {
		if err := assignInputField(input, "NumOfReports", _codebuildNumOfReports); err != nil {
			log.Errorf("invalid --num-of-reports: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetReportGroupTrend(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a resource policy that is identified by its resource ARN.
func codebuild_GetResourcePolicy(cfg aws.Config, client *codebuild.Client) {
	input := &codebuild.GetResourcePolicyInput{
		// ResourceArn: *string, // Required
	}

	if len(_codebuildResourceArn) > 0 {
		input.ResourceArn = aws.String(_codebuildResourceArn)
	}

	if resp, err := client.GetResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Imports the source repository credentials for an CodeBuild project that has
// its source code stored in a GitHub, GitHub Enterprise, GitLab, GitLab Self
// Managed, or Bitbucket repository.
func codebuild_ImportSourceCredentials(cfg aws.Config, client *codebuild.Client) {
	input := &codebuild.ImportSourceCredentialsInput{
		// AuthType: types.AuthType, // Required
		// ServerType: types.ServerType, // Required
		// Token: *string, // Required
	}

	if len(_codebuildAuthType) > 0 {
		if err := assignInputField(input, "AuthType", _codebuildAuthType); err != nil {
			log.Errorf("invalid --auth-type: %s", err.Error())
			return
		}
	}
	if len(_codebuildServerType) > 0 {
		if err := assignInputField(input, "ServerType", _codebuildServerType); err != nil {
			log.Errorf("invalid --server-type: %s", err.Error())
			return
		}
	}
	if len(_codebuildToken) > 0 {
		input.Token = aws.String(_codebuildToken)
	}
	if len(_codebuildShouldOverwrite) > 0 {
		if err := assignInputField(input, "ShouldOverwrite", _codebuildShouldOverwrite); err != nil {
			log.Errorf("invalid --should-overwrite: %s", err.Error())
			return
		}
	}
	if len(_codebuildUsername) > 0 {
		input.Username = aws.String(_codebuildUsername)
	}

	if resp, err := client.ImportSourceCredentials(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Resets the cache for a project.
func codebuild_InvalidateProjectCache(cfg aws.Config, client *codebuild.Client) {
	input := &codebuild.InvalidateProjectCacheInput{
		// ProjectName: *string, // Required
	}

	if len(_codebuildProjectName) > 0 {
		input.ProjectName = aws.String(_codebuildProjectName)
	}

	if resp, err := client.InvalidateProjectCache(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the identifiers of your build batches in the current region.
func codebuild_ListBuildBatches(cfg aws.Config, client *codebuild.Client) {
	input := &codebuild.ListBuildBatchesInput{}

	if len(_codebuildFilter) > 0 {
		if err := assignInputField(input, "Filter", _codebuildFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_codebuildMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codebuildMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codebuildNextToken) > 0 {
		input.NextToken = aws.String(_codebuildNextToken)
	}
	if len(_codebuildSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _codebuildSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListBuildBatches(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codebuild.ListBuildBatchesOutput
	p := codebuild.NewListBuildBatchesPaginator(client, input)
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

// Retrieves the identifiers of the build batches for a specific project.
func codebuild_ListBuildBatchesForProject(cfg aws.Config, client *codebuild.Client) {
	input := &codebuild.ListBuildBatchesForProjectInput{}

	if len(_codebuildFilter) > 0 {
		if err := assignInputField(input, "Filter", _codebuildFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_codebuildMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codebuildMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codebuildNextToken) > 0 {
		input.NextToken = aws.String(_codebuildNextToken)
	}
	if len(_codebuildProjectName) > 0 {
		input.ProjectName = aws.String(_codebuildProjectName)
	}
	if len(_codebuildSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _codebuildSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListBuildBatchesForProject(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codebuild.ListBuildBatchesForProjectOutput
	p := codebuild.NewListBuildBatchesForProjectPaginator(client, input)
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

// Gets a list of build IDs, with each build ID representing a single build.
func codebuild_ListBuilds(cfg aws.Config, client *codebuild.Client) {
	input := &codebuild.ListBuildsInput{}

	if len(_codebuildNextToken) > 0 {
		input.NextToken = aws.String(_codebuildNextToken)
	}
	if len(_codebuildSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _codebuildSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListBuilds(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codebuild.ListBuildsOutput
	p := codebuild.NewListBuildsPaginator(client, input)
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

// Gets a list of build identifiers for the specified build project, with each
// build identifier representing a single build.
func codebuild_ListBuildsForProject(cfg aws.Config, client *codebuild.Client) {
	input := &codebuild.ListBuildsForProjectInput{
		// ProjectName: *string, // Required
	}

	if len(_codebuildProjectName) > 0 {
		input.ProjectName = aws.String(_codebuildProjectName)
	}
	if len(_codebuildNextToken) > 0 {
		input.NextToken = aws.String(_codebuildNextToken)
	}
	if len(_codebuildSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _codebuildSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListBuildsForProject(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codebuild.ListBuildsForProjectOutput
	p := codebuild.NewListBuildsForProjectPaginator(client, input)
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

// Gets a list of command executions for a sandbox.
func codebuild_ListCommandExecutionsForSandbox(cfg aws.Config, client *codebuild.Client) {
	input := &codebuild.ListCommandExecutionsForSandboxInput{
		// SandboxId: *string, // Required
	}

	if len(_codebuildSandboxId) > 0 {
		input.SandboxId = aws.String(_codebuildSandboxId)
	}
	if len(_codebuildMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codebuildMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codebuildNextToken) > 0 {
		input.NextToken = aws.String(_codebuildNextToken)
	}
	if len(_codebuildSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _codebuildSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListCommandExecutionsForSandbox(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codebuild.ListCommandExecutionsForSandboxOutput
	p := codebuild.NewListCommandExecutionsForSandboxPaginator(client, input)
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

// Gets information about Docker images that are managed by CodeBuild.
func codebuild_ListCuratedEnvironmentImages(cfg aws.Config, client *codebuild.Client) {
	input := &codebuild.ListCuratedEnvironmentImagesInput{}

	if resp, err := client.ListCuratedEnvironmentImages(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a list of compute fleet names with each compute fleet name representing a
// single compute fleet.
func codebuild_ListFleets(cfg aws.Config, client *codebuild.Client) {
	input := &codebuild.ListFleetsInput{}

	if len(_codebuildMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codebuildMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codebuildNextToken) > 0 {
		input.NextToken = aws.String(_codebuildNextToken)
	}
	if len(_codebuildSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _codebuildSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_codebuildSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _codebuildSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListFleets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codebuild.ListFleetsOutput
	p := codebuild.NewListFleetsPaginator(client, input)
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

// Gets a list of build project names, with each build project name representing a
// single build project.
func codebuild_ListProjects(cfg aws.Config, client *codebuild.Client) {
	input := &codebuild.ListProjectsInput{}

	if len(_codebuildNextToken) > 0 {
		input.NextToken = aws.String(_codebuildNextToken)
	}
	if len(_codebuildSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _codebuildSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_codebuildSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _codebuildSortOrder); err != nil {
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

	var results []*codebuild.ListProjectsOutput
	p := codebuild.NewListProjectsPaginator(client, input)
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

// Gets a list ARNs for the report groups in the current Amazon Web Services
// account.
func codebuild_ListReportGroups(cfg aws.Config, client *codebuild.Client) {
	input := &codebuild.ListReportGroupsInput{}

	if len(_codebuildMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codebuildMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codebuildNextToken) > 0 {
		input.NextToken = aws.String(_codebuildNextToken)
	}
	if len(_codebuildSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _codebuildSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_codebuildSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _codebuildSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListReportGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codebuild.ListReportGroupsOutput
	p := codebuild.NewListReportGroupsPaginator(client, input)
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

// Returns a list of ARNs for the reports in the current Amazon Web Services
// account.
func codebuild_ListReports(cfg aws.Config, client *codebuild.Client) {
	input := &codebuild.ListReportsInput{}

	if len(_codebuildFilter) > 0 {
		if err := assignInputField(input, "Filter", _codebuildFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_codebuildMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codebuildMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codebuildNextToken) > 0 {
		input.NextToken = aws.String(_codebuildNextToken)
	}
	if len(_codebuildSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _codebuildSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListReports(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codebuild.ListReportsOutput
	p := codebuild.NewListReportsPaginator(client, input)
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

// Returns a list of ARNs for the reports that belong to a ReportGroup .
func codebuild_ListReportsForReportGroup(cfg aws.Config, client *codebuild.Client) {
	input := &codebuild.ListReportsForReportGroupInput{
		// ReportGroupArn: *string, // Required
	}

	if len(_codebuildReportGroupArn) > 0 {
		input.ReportGroupArn = aws.String(_codebuildReportGroupArn)
	}
	if len(_codebuildFilter) > 0 {
		if err := assignInputField(input, "Filter", _codebuildFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_codebuildMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codebuildMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codebuildNextToken) > 0 {
		input.NextToken = aws.String(_codebuildNextToken)
	}
	if len(_codebuildSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _codebuildSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListReportsForReportGroup(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codebuild.ListReportsForReportGroupOutput
	p := codebuild.NewListReportsForReportGroupPaginator(client, input)
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

// Gets a list of sandboxes.
func codebuild_ListSandboxes(cfg aws.Config, client *codebuild.Client) {
	input := &codebuild.ListSandboxesInput{}

	if len(_codebuildMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codebuildMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codebuildNextToken) > 0 {
		input.NextToken = aws.String(_codebuildNextToken)
	}
	if len(_codebuildSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _codebuildSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListSandboxes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codebuild.ListSandboxesOutput
	p := codebuild.NewListSandboxesPaginator(client, input)
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

// Gets a list of sandboxes for a given project.
func codebuild_ListSandboxesForProject(cfg aws.Config, client *codebuild.Client) {
	input := &codebuild.ListSandboxesForProjectInput{
		// ProjectName: *string, // Required
	}

	if len(_codebuildProjectName) > 0 {
		input.ProjectName = aws.String(_codebuildProjectName)
	}
	if len(_codebuildMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codebuildMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codebuildNextToken) > 0 {
		input.NextToken = aws.String(_codebuildNextToken)
	}
	if len(_codebuildSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _codebuildSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListSandboxesForProject(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codebuild.ListSandboxesForProjectOutput
	p := codebuild.NewListSandboxesForProjectPaginator(client, input)
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

// Gets a list of projects that are shared with other Amazon Web Services
// accounts or users.
func codebuild_ListSharedProjects(cfg aws.Config, client *codebuild.Client) {
	input := &codebuild.ListSharedProjectsInput{}

	if len(_codebuildMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codebuildMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codebuildNextToken) > 0 {
		input.NextToken = aws.String(_codebuildNextToken)
	}
	if len(_codebuildSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _codebuildSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_codebuildSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _codebuildSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListSharedProjects(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codebuild.ListSharedProjectsOutput
	p := codebuild.NewListSharedProjectsPaginator(client, input)
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

// Gets a list of report groups that are shared with other Amazon Web Services
// accounts or users.
func codebuild_ListSharedReportGroups(cfg aws.Config, client *codebuild.Client) {
	input := &codebuild.ListSharedReportGroupsInput{}

	if len(_codebuildMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codebuildMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codebuildNextToken) > 0 {
		input.NextToken = aws.String(_codebuildNextToken)
	}
	if len(_codebuildSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _codebuildSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_codebuildSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _codebuildSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListSharedReportGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codebuild.ListSharedReportGroupsOutput
	p := codebuild.NewListSharedReportGroupsPaginator(client, input)
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

// Returns a list of SourceCredentialsInfo objects.
func codebuild_ListSourceCredentials(cfg aws.Config, client *codebuild.Client) {
	input := &codebuild.ListSourceCredentialsInput{}

	if resp, err := client.ListSourceCredentials(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stores a resource policy for the ARN of a Project or ReportGroup object.
func codebuild_PutResourcePolicy(cfg aws.Config, client *codebuild.Client) {
	input := &codebuild.PutResourcePolicyInput{
		// Policy: *string, // Required
		// ResourceArn: *string, // Required
	}

	if len(_codebuildPolicy) > 0 {
		input.Policy = aws.String(_codebuildPolicy)
	}
	if len(_codebuildResourceArn) > 0 {
		input.ResourceArn = aws.String(_codebuildResourceArn)
	}

	if resp, err := client.PutResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Restarts a build.
func codebuild_RetryBuild(cfg aws.Config, client *codebuild.Client) {
	input := &codebuild.RetryBuildInput{}

	if len(_codebuildId) > 0 {
		input.Id = aws.String(_codebuildId)
	}
	if len(_codebuildIdempotencyToken) > 0 {
		input.IdempotencyToken = aws.String(_codebuildIdempotencyToken)
	}

	if resp, err := client.RetryBuild(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Restarts a failed batch build. Only batch builds that have failed can be
// retried.
func codebuild_RetryBuildBatch(cfg aws.Config, client *codebuild.Client) {
	input := &codebuild.RetryBuildBatchInput{}

	if len(_codebuildId) > 0 {
		input.Id = aws.String(_codebuildId)
	}
	if len(_codebuildIdempotencyToken) > 0 {
		input.IdempotencyToken = aws.String(_codebuildIdempotencyToken)
	}
	if len(_codebuildRetryType) > 0 {
		if err := assignInputField(input, "RetryType", _codebuildRetryType); err != nil {
			log.Errorf("invalid --retry-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.RetryBuildBatch(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts running a build with the settings defined in the project. These setting
// include: how to run a build, where to get the source code, which build
// environment to use, which build commands to run, and where to store the build
// output.
//
// You can also start a build run by overriding some of the build settings in the
// project. The overrides only apply for that specific start build request. The
// settings in the project are unaltered.
func codebuild_StartBuild(cfg aws.Config, client *codebuild.Client) {
	input := &codebuild.StartBuildInput{
		// ProjectName: *string, // Required
	}

	if len(_codebuildProjectName) > 0 {
		input.ProjectName = aws.String(_codebuildProjectName)
	}
	if len(_codebuildArtifactsOverride) > 0 {
		if err := assignInputField(input, "ArtifactsOverride", _codebuildArtifactsOverride); err != nil {
			log.Errorf("invalid --artifacts-override: %s", err.Error())
			return
		}
	}
	if len(_codebuildAutoRetryLimitOverride) > 0 {
		if err := assignInputField(input, "AutoRetryLimitOverride", _codebuildAutoRetryLimitOverride); err != nil {
			log.Errorf("invalid --auto-retry-limit-override: %s", err.Error())
			return
		}
	}
	if len(_codebuildBuildStatusConfigOverride) > 0 {
		if err := assignInputField(input, "BuildStatusConfigOverride", _codebuildBuildStatusConfigOverride); err != nil {
			log.Errorf("invalid --build-status-config-override: %s", err.Error())
			return
		}
	}
	if len(_codebuildBuildspecOverride) > 0 {
		input.BuildspecOverride = aws.String(_codebuildBuildspecOverride)
	}
	if len(_codebuildCacheOverride) > 0 {
		if err := assignInputField(input, "CacheOverride", _codebuildCacheOverride); err != nil {
			log.Errorf("invalid --cache-override: %s", err.Error())
			return
		}
	}
	if len(_codebuildCertificateOverride) > 0 {
		input.CertificateOverride = aws.String(_codebuildCertificateOverride)
	}
	if len(_codebuildComputeTypeOverride) > 0 {
		if err := assignInputField(input, "ComputeTypeOverride", _codebuildComputeTypeOverride); err != nil {
			log.Errorf("invalid --compute-type-override: %s", err.Error())
			return
		}
	}
	if len(_codebuildDebugSessionEnabled) > 0 {
		if err := assignInputField(input, "DebugSessionEnabled", _codebuildDebugSessionEnabled); err != nil {
			log.Errorf("invalid --debug-session-enabled: %s", err.Error())
			return
		}
	}
	if len(_codebuildEncryptionKeyOverride) > 0 {
		input.EncryptionKeyOverride = aws.String(_codebuildEncryptionKeyOverride)
	}
	if len(_codebuildEnvironmentTypeOverride) > 0 {
		if err := assignInputField(input, "EnvironmentTypeOverride", _codebuildEnvironmentTypeOverride); err != nil {
			log.Errorf("invalid --environment-type-override: %s", err.Error())
			return
		}
	}
	if len(_codebuildEnvironmentVariablesOverride) > 0 {
		if err := assignInputField(input, "EnvironmentVariablesOverride", _codebuildEnvironmentVariablesOverride); err != nil {
			log.Errorf("invalid --environment-variables-override: %s", err.Error())
			return
		}
	}
	if len(_codebuildFleetOverride) > 0 {
		if err := assignInputField(input, "FleetOverride", _codebuildFleetOverride); err != nil {
			log.Errorf("invalid --fleet-override: %s", err.Error())
			return
		}
	}
	if len(_codebuildGitCloneDepthOverride) > 0 {
		if err := assignInputField(input, "GitCloneDepthOverride", _codebuildGitCloneDepthOverride); err != nil {
			log.Errorf("invalid --git-clone-depth-override: %s", err.Error())
			return
		}
	}
	if len(_codebuildGitSubmodulesConfigOverride) > 0 {
		if err := assignInputField(input, "GitSubmodulesConfigOverride", _codebuildGitSubmodulesConfigOverride); err != nil {
			log.Errorf("invalid --git-submodules-config-override: %s", err.Error())
			return
		}
	}
	if len(_codebuildIdempotencyToken) > 0 {
		input.IdempotencyToken = aws.String(_codebuildIdempotencyToken)
	}
	if len(_codebuildImageOverride) > 0 {
		input.ImageOverride = aws.String(_codebuildImageOverride)
	}
	if len(_codebuildImagePullCredentialsTypeOverride) > 0 {
		if err := assignInputField(input, "ImagePullCredentialsTypeOverride", _codebuildImagePullCredentialsTypeOverride); err != nil {
			log.Errorf("invalid --image-pull-credentials-type-override: %s", err.Error())
			return
		}
	}
	if len(_codebuildInsecureSslOverride) > 0 {
		if err := assignInputField(input, "InsecureSslOverride", _codebuildInsecureSslOverride); err != nil {
			log.Errorf("invalid --insecure-ssl-override: %s", err.Error())
			return
		}
	}
	if len(_codebuildLogsConfigOverride) > 0 {
		if err := assignInputField(input, "LogsConfigOverride", _codebuildLogsConfigOverride); err != nil {
			log.Errorf("invalid --logs-config-override: %s", err.Error())
			return
		}
	}
	if len(_codebuildPrivilegedModeOverride) > 0 {
		if err := assignInputField(input, "PrivilegedModeOverride", _codebuildPrivilegedModeOverride); err != nil {
			log.Errorf("invalid --privileged-mode-override: %s", err.Error())
			return
		}
	}
	if len(_codebuildQueuedTimeoutInMinutesOverride) > 0 {
		if err := assignInputField(input, "QueuedTimeoutInMinutesOverride", _codebuildQueuedTimeoutInMinutesOverride); err != nil {
			log.Errorf("invalid --queued-timeout-in-minutes-override: %s", err.Error())
			return
		}
	}
	if len(_codebuildRegistryCredentialOverride) > 0 {
		if err := assignInputField(input, "RegistryCredentialOverride", _codebuildRegistryCredentialOverride); err != nil {
			log.Errorf("invalid --registry-credential-override: %s", err.Error())
			return
		}
	}
	if len(_codebuildReportBuildStatusOverride) > 0 {
		if err := assignInputField(input, "ReportBuildStatusOverride", _codebuildReportBuildStatusOverride); err != nil {
			log.Errorf("invalid --report-build-status-override: %s", err.Error())
			return
		}
	}
	if len(_codebuildSecondaryArtifactsOverride) > 0 {
		if err := assignInputField(input, "SecondaryArtifactsOverride", _codebuildSecondaryArtifactsOverride); err != nil {
			log.Errorf("invalid --secondary-artifacts-override: %s", err.Error())
			return
		}
	}
	if len(_codebuildSecondarySourcesOverride) > 0 {
		if err := assignInputField(input, "SecondarySourcesOverride", _codebuildSecondarySourcesOverride); err != nil {
			log.Errorf("invalid --secondary-sources-override: %s", err.Error())
			return
		}
	}
	if len(_codebuildSecondarySourcesVersionOverride) > 0 {
		if err := assignInputField(input, "SecondarySourcesVersionOverride", _codebuildSecondarySourcesVersionOverride); err != nil {
			log.Errorf("invalid --secondary-sources-version-override: %s", err.Error())
			return
		}
	}
	if len(_codebuildServiceRoleOverride) > 0 {
		input.ServiceRoleOverride = aws.String(_codebuildServiceRoleOverride)
	}
	if len(_codebuildSourceAuthOverride) > 0 {
		if err := assignInputField(input, "SourceAuthOverride", _codebuildSourceAuthOverride); err != nil {
			log.Errorf("invalid --source-auth-override: %s", err.Error())
			return
		}
	}
	if len(_codebuildSourceLocationOverride) > 0 {
		input.SourceLocationOverride = aws.String(_codebuildSourceLocationOverride)
	}
	if len(_codebuildSourceTypeOverride) > 0 {
		if err := assignInputField(input, "SourceTypeOverride", _codebuildSourceTypeOverride); err != nil {
			log.Errorf("invalid --source-type-override: %s", err.Error())
			return
		}
	}
	if len(_codebuildSourceVersion) > 0 {
		input.SourceVersion = aws.String(_codebuildSourceVersion)
	}
	if len(_codebuildTimeoutInMinutesOverride) > 0 {
		if err := assignInputField(input, "TimeoutInMinutesOverride", _codebuildTimeoutInMinutesOverride); err != nil {
			log.Errorf("invalid --timeout-in-minutes-override: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartBuild(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a batch build for a project.
func codebuild_StartBuildBatch(cfg aws.Config, client *codebuild.Client) {
	input := &codebuild.StartBuildBatchInput{
		// ProjectName: *string, // Required
	}

	if len(_codebuildProjectName) > 0 {
		input.ProjectName = aws.String(_codebuildProjectName)
	}
	if len(_codebuildArtifactsOverride) > 0 {
		if err := assignInputField(input, "ArtifactsOverride", _codebuildArtifactsOverride); err != nil {
			log.Errorf("invalid --artifacts-override: %s", err.Error())
			return
		}
	}
	if len(_codebuildBuildBatchConfigOverride) > 0 {
		if err := assignInputField(input, "BuildBatchConfigOverride", _codebuildBuildBatchConfigOverride); err != nil {
			log.Errorf("invalid --build-batch-config-override: %s", err.Error())
			return
		}
	}
	if len(_codebuildBuildTimeoutInMinutesOverride) > 0 {
		if err := assignInputField(input, "BuildTimeoutInMinutesOverride", _codebuildBuildTimeoutInMinutesOverride); err != nil {
			log.Errorf("invalid --build-timeout-in-minutes-override: %s", err.Error())
			return
		}
	}
	if len(_codebuildBuildspecOverride) > 0 {
		input.BuildspecOverride = aws.String(_codebuildBuildspecOverride)
	}
	if len(_codebuildCacheOverride) > 0 {
		if err := assignInputField(input, "CacheOverride", _codebuildCacheOverride); err != nil {
			log.Errorf("invalid --cache-override: %s", err.Error())
			return
		}
	}
	if len(_codebuildCertificateOverride) > 0 {
		input.CertificateOverride = aws.String(_codebuildCertificateOverride)
	}
	if len(_codebuildComputeTypeOverride) > 0 {
		if err := assignInputField(input, "ComputeTypeOverride", _codebuildComputeTypeOverride); err != nil {
			log.Errorf("invalid --compute-type-override: %s", err.Error())
			return
		}
	}
	if len(_codebuildDebugSessionEnabled) > 0 {
		if err := assignInputField(input, "DebugSessionEnabled", _codebuildDebugSessionEnabled); err != nil {
			log.Errorf("invalid --debug-session-enabled: %s", err.Error())
			return
		}
	}
	if len(_codebuildEncryptionKeyOverride) > 0 {
		input.EncryptionKeyOverride = aws.String(_codebuildEncryptionKeyOverride)
	}
	if len(_codebuildEnvironmentTypeOverride) > 0 {
		if err := assignInputField(input, "EnvironmentTypeOverride", _codebuildEnvironmentTypeOverride); err != nil {
			log.Errorf("invalid --environment-type-override: %s", err.Error())
			return
		}
	}
	if len(_codebuildEnvironmentVariablesOverride) > 0 {
		if err := assignInputField(input, "EnvironmentVariablesOverride", _codebuildEnvironmentVariablesOverride); err != nil {
			log.Errorf("invalid --environment-variables-override: %s", err.Error())
			return
		}
	}
	if len(_codebuildGitCloneDepthOverride) > 0 {
		if err := assignInputField(input, "GitCloneDepthOverride", _codebuildGitCloneDepthOverride); err != nil {
			log.Errorf("invalid --git-clone-depth-override: %s", err.Error())
			return
		}
	}
	if len(_codebuildGitSubmodulesConfigOverride) > 0 {
		if err := assignInputField(input, "GitSubmodulesConfigOverride", _codebuildGitSubmodulesConfigOverride); err != nil {
			log.Errorf("invalid --git-submodules-config-override: %s", err.Error())
			return
		}
	}
	if len(_codebuildIdempotencyToken) > 0 {
		input.IdempotencyToken = aws.String(_codebuildIdempotencyToken)
	}
	if len(_codebuildImageOverride) > 0 {
		input.ImageOverride = aws.String(_codebuildImageOverride)
	}
	if len(_codebuildImagePullCredentialsTypeOverride) > 0 {
		if err := assignInputField(input, "ImagePullCredentialsTypeOverride", _codebuildImagePullCredentialsTypeOverride); err != nil {
			log.Errorf("invalid --image-pull-credentials-type-override: %s", err.Error())
			return
		}
	}
	if len(_codebuildInsecureSslOverride) > 0 {
		if err := assignInputField(input, "InsecureSslOverride", _codebuildInsecureSslOverride); err != nil {
			log.Errorf("invalid --insecure-ssl-override: %s", err.Error())
			return
		}
	}
	if len(_codebuildLogsConfigOverride) > 0 {
		if err := assignInputField(input, "LogsConfigOverride", _codebuildLogsConfigOverride); err != nil {
			log.Errorf("invalid --logs-config-override: %s", err.Error())
			return
		}
	}
	if len(_codebuildPrivilegedModeOverride) > 0 {
		if err := assignInputField(input, "PrivilegedModeOverride", _codebuildPrivilegedModeOverride); err != nil {
			log.Errorf("invalid --privileged-mode-override: %s", err.Error())
			return
		}
	}
	if len(_codebuildQueuedTimeoutInMinutesOverride) > 0 {
		if err := assignInputField(input, "QueuedTimeoutInMinutesOverride", _codebuildQueuedTimeoutInMinutesOverride); err != nil {
			log.Errorf("invalid --queued-timeout-in-minutes-override: %s", err.Error())
			return
		}
	}
	if len(_codebuildRegistryCredentialOverride) > 0 {
		if err := assignInputField(input, "RegistryCredentialOverride", _codebuildRegistryCredentialOverride); err != nil {
			log.Errorf("invalid --registry-credential-override: %s", err.Error())
			return
		}
	}
	if len(_codebuildReportBuildBatchStatusOverride) > 0 {
		if err := assignInputField(input, "ReportBuildBatchStatusOverride", _codebuildReportBuildBatchStatusOverride); err != nil {
			log.Errorf("invalid --report-build-batch-status-override: %s", err.Error())
			return
		}
	}
	if len(_codebuildSecondaryArtifactsOverride) > 0 {
		if err := assignInputField(input, "SecondaryArtifactsOverride", _codebuildSecondaryArtifactsOverride); err != nil {
			log.Errorf("invalid --secondary-artifacts-override: %s", err.Error())
			return
		}
	}
	if len(_codebuildSecondarySourcesOverride) > 0 {
		if err := assignInputField(input, "SecondarySourcesOverride", _codebuildSecondarySourcesOverride); err != nil {
			log.Errorf("invalid --secondary-sources-override: %s", err.Error())
			return
		}
	}
	if len(_codebuildSecondarySourcesVersionOverride) > 0 {
		if err := assignInputField(input, "SecondarySourcesVersionOverride", _codebuildSecondarySourcesVersionOverride); err != nil {
			log.Errorf("invalid --secondary-sources-version-override: %s", err.Error())
			return
		}
	}
	if len(_codebuildServiceRoleOverride) > 0 {
		input.ServiceRoleOverride = aws.String(_codebuildServiceRoleOverride)
	}
	if len(_codebuildSourceAuthOverride) > 0 {
		if err := assignInputField(input, "SourceAuthOverride", _codebuildSourceAuthOverride); err != nil {
			log.Errorf("invalid --source-auth-override: %s", err.Error())
			return
		}
	}
	if len(_codebuildSourceLocationOverride) > 0 {
		input.SourceLocationOverride = aws.String(_codebuildSourceLocationOverride)
	}
	if len(_codebuildSourceTypeOverride) > 0 {
		if err := assignInputField(input, "SourceTypeOverride", _codebuildSourceTypeOverride); err != nil {
			log.Errorf("invalid --source-type-override: %s", err.Error())
			return
		}
	}
	if len(_codebuildSourceVersion) > 0 {
		input.SourceVersion = aws.String(_codebuildSourceVersion)
	}

	if resp, err := client.StartBuildBatch(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a command execution.
func codebuild_StartCommandExecution(cfg aws.Config, client *codebuild.Client) {
	input := &codebuild.StartCommandExecutionInput{
		// Command: *string, // Required
		// SandboxId: *string, // Required
	}

	if len(_codebuildCommand) > 0 {
		input.Command = aws.String(_codebuildCommand)
	}
	if len(_codebuildSandboxId) > 0 {
		input.SandboxId = aws.String(_codebuildSandboxId)
	}
	if len(_codebuildType) > 0 {
		if err := assignInputField(input, "Type", _codebuildType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartCommandExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a sandbox.
func codebuild_StartSandbox(cfg aws.Config, client *codebuild.Client) {
	input := &codebuild.StartSandboxInput{}

	if len(_codebuildIdempotencyToken) > 0 {
		input.IdempotencyToken = aws.String(_codebuildIdempotencyToken)
	}
	if len(_codebuildProjectName) > 0 {
		input.ProjectName = aws.String(_codebuildProjectName)
	}

	if resp, err := client.StartSandbox(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a sandbox connection.
func codebuild_StartSandboxConnection(cfg aws.Config, client *codebuild.Client) {
	input := &codebuild.StartSandboxConnectionInput{
		// SandboxId: *string, // Required
	}

	if len(_codebuildSandboxId) > 0 {
		input.SandboxId = aws.String(_codebuildSandboxId)
	}

	if resp, err := client.StartSandboxConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Attempts to stop running a build.
func codebuild_StopBuild(cfg aws.Config, client *codebuild.Client) {
	input := &codebuild.StopBuildInput{
		// Id: *string, // Required
	}

	if len(_codebuildId) > 0 {
		input.Id = aws.String(_codebuildId)
	}

	if resp, err := client.StopBuild(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a running batch build.
func codebuild_StopBuildBatch(cfg aws.Config, client *codebuild.Client) {
	input := &codebuild.StopBuildBatchInput{
		// Id: *string, // Required
	}

	if len(_codebuildId) > 0 {
		input.Id = aws.String(_codebuildId)
	}

	if resp, err := client.StopBuildBatch(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a sandbox.
func codebuild_StopSandbox(cfg aws.Config, client *codebuild.Client) {
	input := &codebuild.StopSandboxInput{
		// Id: *string, // Required
	}

	if len(_codebuildId) > 0 {
		input.Id = aws.String(_codebuildId)
	}

	if resp, err := client.StopSandbox(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a compute fleet.
func codebuild_UpdateFleet(cfg aws.Config, client *codebuild.Client) {
	input := &codebuild.UpdateFleetInput{
		// Arn: *string, // Required
	}

	if len(_codebuildArn) > 0 {
		input.Arn = aws.String(_codebuildArn)
	}
	if len(_codebuildBaseCapacity) > 0 {
		if err := assignInputField(input, "BaseCapacity", _codebuildBaseCapacity); err != nil {
			log.Errorf("invalid --base-capacity: %s", err.Error())
			return
		}
	}
	if len(_codebuildComputeConfiguration) > 0 {
		if err := assignInputField(input, "ComputeConfiguration", _codebuildComputeConfiguration); err != nil {
			log.Errorf("invalid --compute-configuration: %s", err.Error())
			return
		}
	}
	if len(_codebuildComputeType) > 0 {
		if err := assignInputField(input, "ComputeType", _codebuildComputeType); err != nil {
			log.Errorf("invalid --compute-type: %s", err.Error())
			return
		}
	}
	if len(_codebuildEnvironmentType) > 0 {
		if err := assignInputField(input, "EnvironmentType", _codebuildEnvironmentType); err != nil {
			log.Errorf("invalid --environment-type: %s", err.Error())
			return
		}
	}
	if len(_codebuildFleetServiceRole) > 0 {
		input.FleetServiceRole = aws.String(_codebuildFleetServiceRole)
	}
	if len(_codebuildImageId) > 0 {
		input.ImageId = aws.String(_codebuildImageId)
	}
	if len(_codebuildOverflowBehavior) > 0 {
		if err := assignInputField(input, "OverflowBehavior", _codebuildOverflowBehavior); err != nil {
			log.Errorf("invalid --overflow-behavior: %s", err.Error())
			return
		}
	}
	if len(_codebuildProxyConfiguration) > 0 {
		if err := assignInputField(input, "ProxyConfiguration", _codebuildProxyConfiguration); err != nil {
			log.Errorf("invalid --proxy-configuration: %s", err.Error())
			return
		}
	}
	if len(_codebuildScalingConfiguration) > 0 {
		if err := assignInputField(input, "ScalingConfiguration", _codebuildScalingConfiguration); err != nil {
			log.Errorf("invalid --scaling-configuration: %s", err.Error())
			return
		}
	}
	if len(_codebuildTags) > 0 {
		if err := assignInputField(input, "Tags", _codebuildTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_codebuildVpcConfig) > 0 {
		if err := assignInputField(input, "VpcConfig", _codebuildVpcConfig); err != nil {
			log.Errorf("invalid --vpc-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateFleet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Changes the settings of a build project.
func codebuild_UpdateProject(cfg aws.Config, client *codebuild.Client) {
	input := &codebuild.UpdateProjectInput{
		// Name: *string, // Required
	}

	if len(_codebuildName) > 0 {
		input.Name = aws.String(_codebuildName)
	}
	if len(_codebuildArtifacts) > 0 {
		if err := assignInputField(input, "Artifacts", _codebuildArtifacts); err != nil {
			log.Errorf("invalid --artifacts: %s", err.Error())
			return
		}
	}
	if len(_codebuildAutoRetryLimit) > 0 {
		if err := assignInputField(input, "AutoRetryLimit", _codebuildAutoRetryLimit); err != nil {
			log.Errorf("invalid --auto-retry-limit: %s", err.Error())
			return
		}
	}
	if len(_codebuildBadgeEnabled) > 0 {
		if err := assignInputField(input, "BadgeEnabled", _codebuildBadgeEnabled); err != nil {
			log.Errorf("invalid --badge-enabled: %s", err.Error())
			return
		}
	}
	if len(_codebuildBuildBatchConfig) > 0 {
		if err := assignInputField(input, "BuildBatchConfig", _codebuildBuildBatchConfig); err != nil {
			log.Errorf("invalid --build-batch-config: %s", err.Error())
			return
		}
	}
	if len(_codebuildCache) > 0 {
		if err := assignInputField(input, "Cache", _codebuildCache); err != nil {
			log.Errorf("invalid --cache: %s", err.Error())
			return
		}
	}
	if len(_codebuildConcurrentBuildLimit) > 0 {
		if err := assignInputField(input, "ConcurrentBuildLimit", _codebuildConcurrentBuildLimit); err != nil {
			log.Errorf("invalid --concurrent-build-limit: %s", err.Error())
			return
		}
	}
	if len(_codebuildDescription) > 0 {
		input.Description = aws.String(_codebuildDescription)
	}
	if len(_codebuildEncryptionKey) > 0 {
		input.EncryptionKey = aws.String(_codebuildEncryptionKey)
	}
	if len(_codebuildEnvironment) > 0 {
		if err := assignInputField(input, "Environment", _codebuildEnvironment); err != nil {
			log.Errorf("invalid --environment: %s", err.Error())
			return
		}
	}
	if len(_codebuildFileSystemLocations) > 0 {
		if err := assignInputField(input, "FileSystemLocations", _codebuildFileSystemLocations); err != nil {
			log.Errorf("invalid --file-system-locations: %s", err.Error())
			return
		}
	}
	if len(_codebuildLogsConfig) > 0 {
		if err := assignInputField(input, "LogsConfig", _codebuildLogsConfig); err != nil {
			log.Errorf("invalid --logs-config: %s", err.Error())
			return
		}
	}
	if len(_codebuildQueuedTimeoutInMinutes) > 0 {
		if err := assignInputField(input, "QueuedTimeoutInMinutes", _codebuildQueuedTimeoutInMinutes); err != nil {
			log.Errorf("invalid --queued-timeout-in-minutes: %s", err.Error())
			return
		}
	}
	if len(_codebuildSecondaryArtifacts) > 0 {
		if err := assignInputField(input, "SecondaryArtifacts", _codebuildSecondaryArtifacts); err != nil {
			log.Errorf("invalid --secondary-artifacts: %s", err.Error())
			return
		}
	}
	if len(_codebuildSecondarySourceVersions) > 0 {
		if err := assignInputField(input, "SecondarySourceVersions", _codebuildSecondarySourceVersions); err != nil {
			log.Errorf("invalid --secondary-source-versions: %s", err.Error())
			return
		}
	}
	if len(_codebuildSecondarySources) > 0 {
		if err := assignInputField(input, "SecondarySources", _codebuildSecondarySources); err != nil {
			log.Errorf("invalid --secondary-sources: %s", err.Error())
			return
		}
	}
	if len(_codebuildServiceRole) > 0 {
		input.ServiceRole = aws.String(_codebuildServiceRole)
	}
	if len(_codebuildSource) > 0 {
		if err := assignInputField(input, "Source", _codebuildSource); err != nil {
			log.Errorf("invalid --source: %s", err.Error())
			return
		}
	}
	if len(_codebuildSourceVersion) > 0 {
		input.SourceVersion = aws.String(_codebuildSourceVersion)
	}
	if len(_codebuildTags) > 0 {
		if err := assignInputField(input, "Tags", _codebuildTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_codebuildTimeoutInMinutes) > 0 {
		if err := assignInputField(input, "TimeoutInMinutes", _codebuildTimeoutInMinutes); err != nil {
			log.Errorf("invalid --timeout-in-minutes: %s", err.Error())
			return
		}
	}
	if len(_codebuildVpcConfig) > 0 {
		if err := assignInputField(input, "VpcConfig", _codebuildVpcConfig); err != nil {
			log.Errorf("invalid --vpc-config: %s", err.Error())
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

// Changes the public visibility for a project. The project's build results, logs,
// and artifacts are available to the general public. For more information, see [Public build projects]in
// the CodeBuild User Guide.
//
// The following should be kept in mind when making your projects public:
//
// - All of a project's build results, logs, and artifacts, including builds
// that were run when the project was private, are available to the general public.
//
// - All build logs and artifacts are available to the public. Environment
// variables, source code, and other sensitive information may have been output to
// the build logs and artifacts. You must be careful about what information is
// output to the build logs. Some best practice are:
//
// - Do not store sensitive values in environment variables. We recommend that
// you use an Amazon EC2 Systems Manager Parameter Store or Secrets Manager to
// store sensitive values.
//
// - Follow [Best practices for using webhooks]in the CodeBuild User Guide to limit which entities can trigger a
// build, and do not store the buildspec in the project itself, to ensure that your
// webhooks are as secure as possible.
//
// - A malicious user can use public builds to distribute malicious artifacts.
// We recommend that you review all pull requests to verify that the pull request
// is a legitimate change. We also recommend that you validate any artifacts with
// their checksums to make sure that the correct artifacts are being downloaded.
//
// [Public build projects]: https://docs.aws.amazon.com/codebuild/latest/userguide/public-builds.html
// [Best practices for using webhooks]: https://docs.aws.amazon.com/codebuild/latest/userguide/webhooks.html#webhook-best-practices
func codebuild_UpdateProjectVisibility(cfg aws.Config, client *codebuild.Client) {
	input := &codebuild.UpdateProjectVisibilityInput{
		// ProjectArn: *string, // Required
		// ProjectVisibility: types.ProjectVisibilityType, // Required
	}

	if len(_codebuildProjectArn) > 0 {
		input.ProjectArn = aws.String(_codebuildProjectArn)
	}
	if len(_codebuildProjectVisibility) > 0 {
		if err := assignInputField(input, "ProjectVisibility", _codebuildProjectVisibility); err != nil {
			log.Errorf("invalid --project-visibility: %s", err.Error())
			return
		}
	}
	if len(_codebuildResourceAccessRole) > 0 {
		input.ResourceAccessRole = aws.String(_codebuildResourceAccessRole)
	}

	if resp, err := client.UpdateProjectVisibility(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a report group.
func codebuild_UpdateReportGroup(cfg aws.Config, client *codebuild.Client) {
	input := &codebuild.UpdateReportGroupInput{
		// Arn: *string, // Required
	}

	if len(_codebuildArn) > 0 {
		input.Arn = aws.String(_codebuildArn)
	}
	if len(_codebuildExportConfig) > 0 {
		if err := assignInputField(input, "ExportConfig", _codebuildExportConfig); err != nil {
			log.Errorf("invalid --export-config: %s", err.Error())
			return
		}
	}
	if len(_codebuildTags) > 0 {
		if err := assignInputField(input, "Tags", _codebuildTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateReportGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the webhook associated with an CodeBuild build project.
// If you use Bitbucket for your repository, rotateSecret is ignored.
func codebuild_UpdateWebhook(cfg aws.Config, client *codebuild.Client) {
	input := &codebuild.UpdateWebhookInput{
		// ProjectName: *string, // Required
	}

	if len(_codebuildProjectName) > 0 {
		input.ProjectName = aws.String(_codebuildProjectName)
	}
	if len(_codebuildBranchFilter) > 0 {
		input.BranchFilter = aws.String(_codebuildBranchFilter)
	}
	if len(_codebuildBuildType) > 0 {
		if err := assignInputField(input, "BuildType", _codebuildBuildType); err != nil {
			log.Errorf("invalid --build-type: %s", err.Error())
			return
		}
	}
	if len(_codebuildFilterGroups) > 0 {
		if err := assignInputField(input, "FilterGroups", _codebuildFilterGroups); err != nil {
			log.Errorf("invalid --filter-groups: %s", err.Error())
			return
		}
	}
	if len(_codebuildPullRequestBuildPolicy) > 0 {
		if err := assignInputField(input, "PullRequestBuildPolicy", _codebuildPullRequestBuildPolicy); err != nil {
			log.Errorf("invalid --pull-request-build-policy: %s", err.Error())
			return
		}
	}
	if len(_codebuildRotateSecret) > 0 {
		if err := assignInputField(input, "RotateSecret", _codebuildRotateSecret); err != nil {
			log.Errorf("invalid --rotate-secret: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateWebhook(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_codebuildCmd)
	_codebuildCmd.Flags().SortFlags = false

	_codebuildCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_codebuildCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_codebuildCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_codebuildCmd.Flags().StringVarP(&_codebuildArn, "arn", "", "", "ARN")
	_codebuildCmd.Flags().StringVarP(&_codebuildArtifacts, "artifacts", "", "", "Artifacts")
	_codebuildCmd.Flags().StringVarP(&_codebuildArtifactsOverride, "artifacts-override", "", "", "Artifacts Override")
	_codebuildCmd.Flags().StringVarP(&_codebuildAuthType, "auth-type", "", "", "Auth Type")
	_codebuildCmd.Flags().StringVarP(&_codebuildAutoRetryLimit, "auto-retry-limit", "", "", "Auto Retry Limit")
	_codebuildCmd.Flags().StringVarP(&_codebuildAutoRetryLimitOverride, "auto-retry-limit-override", "", "", "Auto Retry Limit Override")
	_codebuildCmd.Flags().StringVarP(&_codebuildBadgeEnabled, "badge-enabled", "", "", "Badge Enabled")
	_codebuildCmd.Flags().StringVarP(&_codebuildBaseCapacity, "base-capacity", "", "", "Base Capacity")
	_codebuildCmd.Flags().StringVarP(&_codebuildBranchFilter, "branch-filter", "", "", "Branch Filter")
	_codebuildCmd.Flags().StringVarP(&_codebuildBuildBatchConfig, "build-batch-config", "", "", "Build Batch Config")
	_codebuildCmd.Flags().StringVarP(&_codebuildBuildBatchConfigOverride, "build-batch-config-override", "", "", "Build Batch Config Override")
	_codebuildCmd.Flags().StringVarP(&_codebuildBuildStatusConfigOverride, "build-status-config-override", "", "", "Build Status Config Override")
	_codebuildCmd.Flags().StringVarP(&_codebuildBuildTimeoutInMinutesOverride, "build-timeout-in-minutes-override", "", "", "Build Timeout In Minutes Override")
	_codebuildCmd.Flags().StringVarP(&_codebuildBuildType, "build-type", "", "", "Build Type")
	_codebuildCmd.Flags().StringVarP(&_codebuildBuildspecOverride, "buildspec-override", "", "", "Buildspec Override")
	_codebuildCmd.Flags().StringVarP(&_codebuildCache, "cache", "", "", "Cache")
	_codebuildCmd.Flags().StringVarP(&_codebuildCacheOverride, "cache-override", "", "", "Cache Override")
	_codebuildCmd.Flags().StringVarP(&_codebuildCertificateOverride, "certificate-override", "", "", "Certificate Override")
	_codebuildCmd.Flags().StringVarP(&_codebuildCommand, "command", "", "", "Command")
	_codebuildCmd.Flags().StringSliceVarP(&_codebuildCommandExecutionIds, "command-execution-ids", "", nil, "Command Execution Ids")
	_codebuildCmd.Flags().StringVarP(&_codebuildComputeConfiguration, "compute-configuration", "", "", "Compute Configuration")
	_codebuildCmd.Flags().StringVarP(&_codebuildComputeType, "compute-type", "", "", "Compute Type")
	_codebuildCmd.Flags().StringVarP(&_codebuildComputeTypeOverride, "compute-type-override", "", "", "Compute Type Override")
	_codebuildCmd.Flags().StringVarP(&_codebuildConcurrentBuildLimit, "concurrent-build-limit", "", "", "Concurrent Build Limit")
	_codebuildCmd.Flags().StringVarP(&_codebuildDebugSessionEnabled, "debug-session-enabled", "", "", "Debug Session Enabled")
	_codebuildCmd.Flags().StringVarP(&_codebuildDeleteReports, "delete-reports", "", "", "Delete Reports")
	_codebuildCmd.Flags().StringVarP(&_codebuildDescription, "description", "", "", "Description")
	_codebuildCmd.Flags().StringVarP(&_codebuildEncryptionKey, "encryption-key", "", "", "Encryption Key")
	_codebuildCmd.Flags().StringVarP(&_codebuildEncryptionKeyOverride, "encryption-key-override", "", "", "Encryption Key Override")
	_codebuildCmd.Flags().StringVarP(&_codebuildEnvironment, "environment", "", "", "Environment")
	_codebuildCmd.Flags().StringVarP(&_codebuildEnvironmentType, "environment-type", "", "", "Environment Type")
	_codebuildCmd.Flags().StringVarP(&_codebuildEnvironmentTypeOverride, "environment-type-override", "", "", "Environment Type Override")
	_codebuildCmd.Flags().StringVarP(&_codebuildEnvironmentVariablesOverride, "environment-variables-override", "", "", "Environment Variables Override")
	_codebuildCmd.Flags().StringVarP(&_codebuildExportConfig, "export-config", "", "", "Export Config")
	_codebuildCmd.Flags().StringVarP(&_codebuildFileSystemLocations, "file-system-locations", "", "", "File System Locations")
	_codebuildCmd.Flags().StringVarP(&_codebuildFilter, "filter", "", "", "Filter")
	_codebuildCmd.Flags().StringVarP(&_codebuildFilterGroups, "filter-groups", "", "", "Filter Groups")
	_codebuildCmd.Flags().StringVarP(&_codebuildFleetOverride, "fleet-override", "", "", "Fleet Override")
	_codebuildCmd.Flags().StringVarP(&_codebuildFleetServiceRole, "fleet-service-role", "", "", "Fleet Service Role")
	_codebuildCmd.Flags().StringVarP(&_codebuildGitCloneDepthOverride, "git-clone-depth-override", "", "", "Git Clone Depth Override")
	_codebuildCmd.Flags().StringVarP(&_codebuildGitSubmodulesConfigOverride, "git-submodules-config-override", "", "", "Git Submodules Config Override")
	_codebuildCmd.Flags().StringVarP(&_codebuildId, "id", "", "", "ID")
	_codebuildCmd.Flags().StringVarP(&_codebuildIdempotencyToken, "idempotency-token", "", "", "Idempotency Token")
	_codebuildCmd.Flags().StringSliceVarP(&_codebuildIds, "ids", "", nil, "Ids")
	_codebuildCmd.Flags().StringVarP(&_codebuildImageId, "image-id", "", "", "Image ID")
	_codebuildCmd.Flags().StringVarP(&_codebuildImageOverride, "image-override", "", "", "Image Override")
	_codebuildCmd.Flags().StringVarP(&_codebuildImagePullCredentialsTypeOverride, "image-pull-credentials-type-override", "", "", "Image Pull Credentials Type Override")
	_codebuildCmd.Flags().StringVarP(&_codebuildInsecureSslOverride, "insecure-ssl-override", "", "", "Insecure SSL Override")
	_codebuildCmd.Flags().StringVarP(&_codebuildLogsConfig, "logs-config", "", "", "Logs Config")
	_codebuildCmd.Flags().StringVarP(&_codebuildLogsConfigOverride, "logs-config-override", "", "", "Logs Config Override")
	_codebuildCmd.Flags().StringVarP(&_codebuildManualCreation, "manual-creation", "", "", "Manual Creation")
	_codebuildCmd.Flags().StringVarP(&_codebuildMaxLineCoveragePercentage, "max-line-coverage-percentage", "", "", "Max Line Coverage Percentage")
	_codebuildCmd.Flags().StringVarP(&_codebuildMaxResults, "max-results", "", "", "Max Results")
	_codebuildCmd.Flags().StringVarP(&_codebuildMinLineCoveragePercentage, "min-line-coverage-percentage", "", "", "Min Line Coverage Percentage")
	_codebuildCmd.Flags().StringVarP(&_codebuildName, "name", "", "", "Name")
	_codebuildCmd.Flags().StringSliceVarP(&_codebuildNames, "names", "", nil, "Names")
	_codebuildCmd.Flags().StringVarP(&_codebuildNextToken, "next-token", "", "", "Next Token")
	_codebuildCmd.Flags().StringVarP(&_codebuildNumOfReports, "num-of-reports", "", "", "Num Of Reports")
	_codebuildCmd.Flags().StringVarP(&_codebuildOverflowBehavior, "overflow-behavior", "", "", "Overflow Behavior")
	_codebuildCmd.Flags().StringVarP(&_codebuildPolicy, "policy", "", "", "Policy")
	_codebuildCmd.Flags().StringVarP(&_codebuildPrivilegedModeOverride, "privileged-mode-override", "", "", "Privileged Mode Override")
	_codebuildCmd.Flags().StringVarP(&_codebuildProjectArn, "project-arn", "", "", "Project ARN")
	_codebuildCmd.Flags().StringVarP(&_codebuildProjectName, "project-name", "", "", "Project Name")
	_codebuildCmd.Flags().StringVarP(&_codebuildProjectVisibility, "project-visibility", "", "", "Project Visibility")
	_codebuildCmd.Flags().StringVarP(&_codebuildProxyConfiguration, "proxy-configuration", "", "", "Proxy Configuration")
	_codebuildCmd.Flags().StringVarP(&_codebuildPullRequestBuildPolicy, "pull-request-build-policy", "", "", "Pull Request Build Policy")
	_codebuildCmd.Flags().StringVarP(&_codebuildQueuedTimeoutInMinutes, "queued-timeout-in-minutes", "", "", "Queued Timeout In Minutes")
	_codebuildCmd.Flags().StringVarP(&_codebuildQueuedTimeoutInMinutesOverride, "queued-timeout-in-minutes-override", "", "", "Queued Timeout In Minutes Override")
	_codebuildCmd.Flags().StringVarP(&_codebuildRegistryCredentialOverride, "registry-credential-override", "", "", "Registry Credential Override")
	_codebuildCmd.Flags().StringVarP(&_codebuildReportArn, "report-arn", "", "", "Report ARN")
	_codebuildCmd.Flags().StringSliceVarP(&_codebuildReportArns, "report-arns", "", nil, "Report Arns")
	_codebuildCmd.Flags().StringVarP(&_codebuildReportBuildBatchStatusOverride, "report-build-batch-status-override", "", "", "Report Build Batch Status Override")
	_codebuildCmd.Flags().StringVarP(&_codebuildReportBuildStatusOverride, "report-build-status-override", "", "", "Report Build Status Override")
	_codebuildCmd.Flags().StringVarP(&_codebuildReportGroupArn, "report-group-arn", "", "", "Report Group ARN")
	_codebuildCmd.Flags().StringSliceVarP(&_codebuildReportGroupArns, "report-group-arns", "", nil, "Report Group Arns")
	_codebuildCmd.Flags().StringVarP(&_codebuildResourceAccessRole, "resource-access-role", "", "", "Resource Access Role")
	_codebuildCmd.Flags().StringVarP(&_codebuildResourceArn, "resource-arn", "", "", "Resource ARN")
	_codebuildCmd.Flags().StringVarP(&_codebuildRetryType, "retry-type", "", "", "Retry Type")
	_codebuildCmd.Flags().StringVarP(&_codebuildRotateSecret, "rotate-secret", "", "", "Rotate Secret")
	_codebuildCmd.Flags().StringVarP(&_codebuildSandboxId, "sandbox-id", "", "", "Sandbox ID")
	_codebuildCmd.Flags().StringVarP(&_codebuildScalingConfiguration, "scaling-configuration", "", "", "Scaling Configuration")
	_codebuildCmd.Flags().StringVarP(&_codebuildScopeConfiguration, "scope-configuration", "", "", "Scope Configuration")
	_codebuildCmd.Flags().StringVarP(&_codebuildSecondaryArtifacts, "secondary-artifacts", "", "", "Secondary Artifacts")
	_codebuildCmd.Flags().StringVarP(&_codebuildSecondaryArtifactsOverride, "secondary-artifacts-override", "", "", "Secondary Artifacts Override")
	_codebuildCmd.Flags().StringVarP(&_codebuildSecondarySourceVersions, "secondary-source-versions", "", "", "Secondary Source Versions")
	_codebuildCmd.Flags().StringVarP(&_codebuildSecondarySources, "secondary-sources", "", "", "Secondary Sources")
	_codebuildCmd.Flags().StringVarP(&_codebuildSecondarySourcesOverride, "secondary-sources-override", "", "", "Secondary Sources Override")
	_codebuildCmd.Flags().StringVarP(&_codebuildSecondarySourcesVersionOverride, "secondary-sources-version-override", "", "", "Secondary Sources Version Override")
	_codebuildCmd.Flags().StringVarP(&_codebuildServerType, "server-type", "", "", "Server Type")
	_codebuildCmd.Flags().StringVarP(&_codebuildServiceRole, "service-role", "", "", "Service Role")
	_codebuildCmd.Flags().StringVarP(&_codebuildServiceRoleOverride, "service-role-override", "", "", "Service Role Override")
	_codebuildCmd.Flags().StringVarP(&_codebuildShouldOverwrite, "should-overwrite", "", "", "Should Overwrite")
	_codebuildCmd.Flags().StringVarP(&_codebuildSortBy, "sort-by", "", "", "Sort By")
	_codebuildCmd.Flags().StringVarP(&_codebuildSortOrder, "sort-order", "", "", "Sort Order")
	_codebuildCmd.Flags().StringVarP(&_codebuildSource, "source", "", "", "Source")
	_codebuildCmd.Flags().StringVarP(&_codebuildSourceAuthOverride, "source-auth-override", "", "", "Source Auth Override")
	_codebuildCmd.Flags().StringVarP(&_codebuildSourceLocationOverride, "source-location-override", "", "", "Source Location Override")
	_codebuildCmd.Flags().StringVarP(&_codebuildSourceTypeOverride, "source-type-override", "", "", "Source Type Override")
	_codebuildCmd.Flags().StringVarP(&_codebuildSourceVersion, "source-version", "", "", "Source Version")
	_codebuildCmd.Flags().StringVarP(&_codebuildTags, "tags", "", "", "Tags")
	_codebuildCmd.Flags().StringVarP(&_codebuildTimeoutInMinutes, "timeout-in-minutes", "", "", "Timeout In Minutes")
	_codebuildCmd.Flags().StringVarP(&_codebuildTimeoutInMinutesOverride, "timeout-in-minutes-override", "", "", "Timeout In Minutes Override")
	_codebuildCmd.Flags().StringVarP(&_codebuildToken, "token", "", "", "Token")
	_codebuildCmd.Flags().StringVarP(&_codebuildTrendField, "trend-field", "", "", "Trend Field")
	_codebuildCmd.Flags().StringVarP(&_codebuildType, "type", "", "", "Type")
	_codebuildCmd.Flags().StringVarP(&_codebuildUsername, "username", "", "", "Username")
	_codebuildCmd.Flags().StringVarP(&_codebuildVpcConfig, "vpc-config", "", "", "VPC Config")

	_codebuildCmd.Flags().BoolVarP(&_codebuildBatchDeleteBuilds, "batch-delete-builds", "", false, "Batch Delete Builds")
	_codebuildCmd.Flags().BoolVarP(&_codebuildBatchGetBuildBatches, "batch-get-build-batches", "", false, "Batch Get Build Batches")
	_codebuildCmd.Flags().BoolVarP(&_codebuildBatchGetBuilds, "batch-get-builds", "", false, "Batch Get Builds")
	_codebuildCmd.Flags().BoolVarP(&_codebuildBatchGetCommandExecutions, "batch-get-command-executions", "", false, "Batch Get Command Executions")
	_codebuildCmd.Flags().BoolVarP(&_codebuildBatchGetFleets, "batch-get-fleets", "", false, "Batch Get Fleets")
	_codebuildCmd.Flags().BoolVarP(&_codebuildBatchGetProjects, "batch-get-projects", "", false, "Batch Get Projects")
	_codebuildCmd.Flags().BoolVarP(&_codebuildBatchGetReportGroups, "batch-get-report-groups", "", false, "Batch Get Report Groups")
	_codebuildCmd.Flags().BoolVarP(&_codebuildBatchGetReports, "batch-get-reports", "", false, "Batch Get Reports")
	_codebuildCmd.Flags().BoolVarP(&_codebuildBatchGetSandboxes, "batch-get-sandboxes", "", false, "Batch Get Sandboxes")
	_codebuildCmd.Flags().BoolVarP(&_codebuildCreateFleet, "create-fleet", "", false, "Create Fleet")
	_codebuildCmd.Flags().BoolVarP(&_codebuildCreateProject, "create-project", "", false, "Create Project")
	_codebuildCmd.Flags().BoolVarP(&_codebuildCreateReportGroup, "create-report-group", "", false, "Create Report Group")
	_codebuildCmd.Flags().BoolVarP(&_codebuildCreateWebhook, "create-webhook", "", false, "Create Webhook")
	_codebuildCmd.Flags().BoolVarP(&_codebuildDeleteBuildBatch, "delete-build-batch", "", false, "Delete Build Batch")
	_codebuildCmd.Flags().BoolVarP(&_codebuildDeleteFleet, "delete-fleet", "", false, "Delete Fleet")
	_codebuildCmd.Flags().BoolVarP(&_codebuildDeleteProject, "delete-project", "", false, "Delete Project")
	_codebuildCmd.Flags().BoolVarP(&_codebuildDeleteReport, "delete-report", "", false, "Delete Report")
	_codebuildCmd.Flags().BoolVarP(&_codebuildDeleteReportGroup, "delete-report-group", "", false, "Delete Report Group")
	_codebuildCmd.Flags().BoolVarP(&_codebuildDeleteResourcePolicy, "delete-resource-policy", "", false, "Delete Resource Policy")
	_codebuildCmd.Flags().BoolVarP(&_codebuildDeleteSourceCredentials, "delete-source-credentials", "", false, "Delete Source Credentials")
	_codebuildCmd.Flags().BoolVarP(&_codebuildDeleteWebhook, "delete-webhook", "", false, "Delete Webhook")
	_codebuildCmd.Flags().BoolVarP(&_codebuildDescribeCodeCoverages, "describe-code-coverages", "", false, "Describe Code Coverages")
	_codebuildCmd.Flags().BoolVarP(&_codebuildDescribeTestCases, "describe-test-cases", "", false, "Describe Test Cases")
	_codebuildCmd.Flags().BoolVarP(&_codebuildGetReportGroupTrend, "get-report-group-trend", "", false, "Get Report Group Trend")
	_codebuildCmd.Flags().BoolVarP(&_codebuildGetResourcePolicy, "get-resource-policy", "", false, "Get Resource Policy")
	_codebuildCmd.Flags().BoolVarP(&_codebuildImportSourceCredentials, "import-source-credentials", "", false, "Import Source Credentials")
	_codebuildCmd.Flags().BoolVarP(&_codebuildInvalidateProjectCache, "invalidate-project-cache", "", false, "Invalidate Project Cache")
	_codebuildCmd.Flags().BoolVarP(&_codebuildListBuildBatches, "list-build-batches", "", false, "List Build Batches")
	_codebuildCmd.Flags().BoolVarP(&_codebuildListBuildBatchesForProject, "list-build-batches-for-project", "", false, "List Build Batches For Project")
	_codebuildCmd.Flags().BoolVarP(&_codebuildListBuilds, "list-builds", "", false, "List Builds")
	_codebuildCmd.Flags().BoolVarP(&_codebuildListBuildsForProject, "list-builds-for-project", "", false, "List Builds For Project")
	_codebuildCmd.Flags().BoolVarP(&_codebuildListCommandExecutionsForSandbox, "list-command-executions-for-sandbox", "", false, "List Command Executions For Sandbox")
	_codebuildCmd.Flags().BoolVarP(&_codebuildListCuratedEnvironmentImages, "list-curated-environment-images", "", false, "List Curated Environment Images")
	_codebuildCmd.Flags().BoolVarP(&_codebuildListFleets, "list-fleets", "", false, "List Fleets")
	_codebuildCmd.Flags().BoolVarP(&_codebuildListProjects, "list-projects", "", false, "List Projects")
	_codebuildCmd.Flags().BoolVarP(&_codebuildListReportGroups, "list-report-groups", "", false, "List Report Groups")
	_codebuildCmd.Flags().BoolVarP(&_codebuildListReports, "list-reports", "", false, "List Reports")
	_codebuildCmd.Flags().BoolVarP(&_codebuildListReportsForReportGroup, "list-reports-for-report-group", "", false, "List Reports For Report Group")
	_codebuildCmd.Flags().BoolVarP(&_codebuildListSandboxes, "list-sandboxes", "", false, "List Sandboxes")
	_codebuildCmd.Flags().BoolVarP(&_codebuildListSandboxesForProject, "list-sandboxes-for-project", "", false, "List Sandboxes For Project")
	_codebuildCmd.Flags().BoolVarP(&_codebuildListSharedProjects, "list-shared-projects", "", false, "List Shared Projects")
	_codebuildCmd.Flags().BoolVarP(&_codebuildListSharedReportGroups, "list-shared-report-groups", "", false, "List Shared Report Groups")
	_codebuildCmd.Flags().BoolVarP(&_codebuildListSourceCredentials, "list-source-credentials", "", false, "List Source Credentials")
	_codebuildCmd.Flags().BoolVarP(&_codebuildPutResourcePolicy, "put-resource-policy", "", false, "Put Resource Policy")
	_codebuildCmd.Flags().BoolVarP(&_codebuildRetryBuild, "retry-build", "", false, "Retry Build")
	_codebuildCmd.Flags().BoolVarP(&_codebuildRetryBuildBatch, "retry-build-batch", "", false, "Retry Build Batch")
	_codebuildCmd.Flags().BoolVarP(&_codebuildStartBuild, "start-build", "", false, "Start Build")
	_codebuildCmd.Flags().BoolVarP(&_codebuildStartBuildBatch, "start-build-batch", "", false, "Start Build Batch")
	_codebuildCmd.Flags().BoolVarP(&_codebuildStartCommandExecution, "start-command-execution", "", false, "Start Command Execution")
	_codebuildCmd.Flags().BoolVarP(&_codebuildStartSandbox, "start-sandbox", "", false, "Start Sandbox")
	_codebuildCmd.Flags().BoolVarP(&_codebuildStartSandboxConnection, "start-sandbox-connection", "", false, "Start Sandbox Connection")
	_codebuildCmd.Flags().BoolVarP(&_codebuildStopBuild, "stop-build", "", false, "Stop Build")
	_codebuildCmd.Flags().BoolVarP(&_codebuildStopBuildBatch, "stop-build-batch", "", false, "Stop Build Batch")
	_codebuildCmd.Flags().BoolVarP(&_codebuildStopSandbox, "stop-sandbox", "", false, "Stop Sandbox")
	_codebuildCmd.Flags().BoolVarP(&_codebuildUpdateFleet, "update-fleet", "", false, "Update Fleet")
	_codebuildCmd.Flags().BoolVarP(&_codebuildUpdateProject, "update-project", "", false, "Update Project")
	_codebuildCmd.Flags().BoolVarP(&_codebuildUpdateProjectVisibility, "update-project-visibility", "", false, "Update Project Visibility")
	_codebuildCmd.Flags().BoolVarP(&_codebuildUpdateReportGroup, "update-report-group", "", false, "Update Report Group")
	_codebuildCmd.Flags().BoolVarP(&_codebuildUpdateWebhook, "update-webhook", "", false, "Update Webhook")

}
