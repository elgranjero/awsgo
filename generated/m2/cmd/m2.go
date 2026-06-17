package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/m2"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// m2Cmd represents the m2 command
var _m2Cmd = &cobra.Command{
	Use:   "m2",
	Short: "AWS m2 CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := m2.NewFromConfig(cfg)
		if _m2CancelBatchJobExecution {
			m2_CancelBatchJobExecution(cfg, client)
			return
		}
		if _m2CreateApplication {
			m2_CreateApplication(cfg, client)
			return
		}
		if _m2CreateDataSetExportTask {
			m2_CreateDataSetExportTask(cfg, client)
			return
		}
		if _m2CreateDataSetImportTask {
			m2_CreateDataSetImportTask(cfg, client)
			return
		}
		if _m2CreateDeployment {
			m2_CreateDeployment(cfg, client)
			return
		}
		if _m2CreateEnvironment {
			m2_CreateEnvironment(cfg, client)
			return
		}
		if _m2DeleteApplication {
			m2_DeleteApplication(cfg, client)
			return
		}
		if _m2DeleteApplicationFromEnvironment {
			m2_DeleteApplicationFromEnvironment(cfg, client)
			return
		}
		if _m2DeleteEnvironment {
			m2_DeleteEnvironment(cfg, client)
			return
		}
		if _m2GetApplication {
			m2_GetApplication(cfg, client)
			return
		}
		if _m2GetApplicationVersion {
			m2_GetApplicationVersion(cfg, client)
			return
		}
		if _m2GetBatchJobExecution {
			m2_GetBatchJobExecution(cfg, client)
			return
		}
		if _m2GetDataSetDetails {
			m2_GetDataSetDetails(cfg, client)
			return
		}
		if _m2GetDataSetExportTask {
			m2_GetDataSetExportTask(cfg, client)
			return
		}
		if _m2GetDataSetImportTask {
			m2_GetDataSetImportTask(cfg, client)
			return
		}
		if _m2GetDeployment {
			m2_GetDeployment(cfg, client)
			return
		}
		if _m2GetEnvironment {
			m2_GetEnvironment(cfg, client)
			return
		}
		if _m2GetSignedBluinsightsUrl {
			m2_GetSignedBluinsightsUrl(cfg, client)
			return
		}
		if _m2ListApplicationVersions {
			m2_ListApplicationVersions(cfg, client)
			return
		}
		if _m2ListApplications {
			m2_ListApplications(cfg, client)
			return
		}
		if _m2ListBatchJobDefinitions {
			m2_ListBatchJobDefinitions(cfg, client)
			return
		}
		if _m2ListBatchJobExecutions {
			m2_ListBatchJobExecutions(cfg, client)
			return
		}
		if _m2ListBatchJobRestartPoints {
			m2_ListBatchJobRestartPoints(cfg, client)
			return
		}
		if _m2ListDataSetExportHistory {
			m2_ListDataSetExportHistory(cfg, client)
			return
		}
		if _m2ListDataSetImportHistory {
			m2_ListDataSetImportHistory(cfg, client)
			return
		}
		if _m2ListDataSets {
			m2_ListDataSets(cfg, client)
			return
		}
		if _m2ListDeployments {
			m2_ListDeployments(cfg, client)
			return
		}
		if _m2ListEngineVersions {
			m2_ListEngineVersions(cfg, client)
			return
		}
		if _m2ListEnvironments {
			m2_ListEnvironments(cfg, client)
			return
		}
		if _m2ListTagsForResource {
			m2_ListTagsForResource(cfg, client)
			return
		}
		if _m2StartApplication {
			m2_StartApplication(cfg, client)
			return
		}
		if _m2StartBatchJob {
			m2_StartBatchJob(cfg, client)
			return
		}
		if _m2StopApplication {
			m2_StopApplication(cfg, client)
			return
		}
		if _m2TagResource {
			m2_TagResource(cfg, client)
			return
		}
		if _m2UntagResource {
			m2_UntagResource(cfg, client)
			return
		}
		if _m2UpdateApplication {
			m2_UpdateApplication(cfg, client)
			return
		}
		if _m2UpdateEnvironment {
			m2_UpdateEnvironment(cfg, client)
			return
		}

	},
}

var (
	_m2CancelBatchJobExecution          bool
	_m2CreateApplication                bool
	_m2CreateDataSetExportTask          bool
	_m2CreateDataSetImportTask          bool
	_m2CreateDeployment                 bool
	_m2CreateEnvironment                bool
	_m2DeleteApplication                bool
	_m2DeleteApplicationFromEnvironment bool
	_m2DeleteEnvironment                bool
	_m2GetApplication                   bool
	_m2GetApplicationVersion            bool
	_m2GetBatchJobExecution             bool
	_m2GetDataSetDetails                bool
	_m2GetDataSetExportTask             bool
	_m2GetDataSetImportTask             bool
	_m2GetDeployment                    bool
	_m2GetEnvironment                   bool
	_m2GetSignedBluinsightsUrl          bool
	_m2ListApplicationVersions          bool
	_m2ListApplications                 bool
	_m2ListBatchJobDefinitions          bool
	_m2ListBatchJobExecutions           bool
	_m2ListBatchJobRestartPoints        bool
	_m2ListDataSetExportHistory         bool
	_m2ListDataSetImportHistory         bool
	_m2ListDataSets                     bool
	_m2ListDeployments                  bool
	_m2ListEngineVersions               bool
	_m2ListEnvironments                 bool
	_m2ListTagsForResource              bool
	_m2StartApplication                 bool
	_m2StartBatchJob                    bool
	_m2StopApplication                  bool
	_m2TagResource                      bool
	_m2UntagResource                    bool
	_m2UpdateApplication                bool
	_m2UpdateEnvironment                bool

	_m2ApplicationId                string
	_m2ApplicationVersion           string
	_m2ApplyDuringMaintenanceWindow string
	_m2AuthSecretsManagerArn        string
	_m2BatchJobIdentifier           string
	_m2ClientToken                  string
	_m2CurrentApplicationVersion    string
	_m2DataSetName                  string
	_m2Definition                   string
	_m2DeploymentId                 string
	_m2Description                  string
	_m2DesiredCapacity              string
	_m2EngineType                   string
	_m2EngineVersion                string
	_m2EnvironmentId                string
	_m2ExecutionId                  string
	_m2ExecutionIds                 []string
	_m2ExportConfig                 string
	_m2ForceStop                    string
	_m2ForceUpdate                  string
	_m2HighAvailabilityConfig       string
	_m2ImportConfig                 string
	_m2InstanceType                 string
	_m2JobName                      string
	_m2JobParams                    string
	_m2KmsKeyId                     string
	_m2MaxResults                   string
	_m2Name                         string
	_m2NameFilter                   string
	_m2Names                        []string
	_m2NetworkType                  string
	_m2NextToken                    string
	_m2PreferredMaintenanceWindow   string
	_m2Prefix                       string
	_m2PubliclyAccessible           string
	_m2ResourceArn                  string
	_m2RoleArn                      string
	_m2SecurityGroupIds             []string
	_m2StartedAfter                 string
	_m2StartedBefore                string
	_m2Status                       string
	_m2StorageConfigurations        string
	_m2SubnetIds                    []string
	_m2TagKeys                      []string
	_m2Tags                         string
	_m2TaskId                       string
)

// Cancels the running of a specific batch job execution.
func m2_CancelBatchJobExecution(cfg aws.Config, client *m2.Client) {
	input := &m2.CancelBatchJobExecutionInput{
		// ApplicationId: *string, // Required
		// ExecutionId: *string, // Required
	}

	if len(_m2ApplicationId) > 0 {
		input.ApplicationId = aws.String(_m2ApplicationId)
	}
	if len(_m2ExecutionId) > 0 {
		input.ExecutionId = aws.String(_m2ExecutionId)
	}
	if len(_m2AuthSecretsManagerArn) > 0 {
		input.AuthSecretsManagerArn = aws.String(_m2AuthSecretsManagerArn)
	}

	if resp, err := client.CancelBatchJobExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new application with given parameters. Requires an existing runtime
// environment and application definition file.
func m2_CreateApplication(cfg aws.Config, client *m2.Client) {
	input := &m2.CreateApplicationInput{
		// Definition: types.Definition, // Required
		// EngineType: types.EngineType, // Required
		// Name: *string, // Required
	}

	if len(_m2Definition) > 0 {
		if err := assignInputField(input, "Definition", _m2Definition); err != nil {
			log.Errorf("invalid --definition: %s", err.Error())
			return
		}
	}
	if len(_m2EngineType) > 0 {
		if err := assignInputField(input, "EngineType", _m2EngineType); err != nil {
			log.Errorf("invalid --engine-type: %s", err.Error())
			return
		}
	}
	if len(_m2Name) > 0 {
		input.Name = aws.String(_m2Name)
	}
	if len(_m2ClientToken) > 0 {
		input.ClientToken = aws.String(_m2ClientToken)
	}
	if len(_m2Description) > 0 {
		input.Description = aws.String(_m2Description)
	}
	if len(_m2KmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_m2KmsKeyId)
	}
	if len(_m2RoleArn) > 0 {
		input.RoleArn = aws.String(_m2RoleArn)
	}
	if len(_m2Tags) > 0 {
		if err := assignInputField(input, "Tags", _m2Tags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a data set export task for a specific application.
func m2_CreateDataSetExportTask(cfg aws.Config, client *m2.Client) {
	input := &m2.CreateDataSetExportTaskInput{
		// ApplicationId: *string, // Required
		// ExportConfig: types.DataSetExportConfig, // Required
	}

	if len(_m2ApplicationId) > 0 {
		input.ApplicationId = aws.String(_m2ApplicationId)
	}
	if len(_m2ExportConfig) > 0 {
		if err := assignInputField(input, "ExportConfig", _m2ExportConfig); err != nil {
			log.Errorf("invalid --export-config: %s", err.Error())
			return
		}
	}
	if len(_m2ClientToken) > 0 {
		input.ClientToken = aws.String(_m2ClientToken)
	}
	if len(_m2KmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_m2KmsKeyId)
	}

	if resp, err := client.CreateDataSetExportTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a data set import task for a specific application.
func m2_CreateDataSetImportTask(cfg aws.Config, client *m2.Client) {
	input := &m2.CreateDataSetImportTaskInput{
		// ApplicationId: *string, // Required
		// ImportConfig: types.DataSetImportConfig, // Required
	}

	if len(_m2ApplicationId) > 0 {
		input.ApplicationId = aws.String(_m2ApplicationId)
	}
	if len(_m2ImportConfig) > 0 {
		if err := assignInputField(input, "ImportConfig", _m2ImportConfig); err != nil {
			log.Errorf("invalid --import-config: %s", err.Error())
			return
		}
	}
	if len(_m2ClientToken) > 0 {
		input.ClientToken = aws.String(_m2ClientToken)
	}

	if resp, err := client.CreateDataSetImportTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates and starts a deployment to deploy an application into a runtime
// environment.
func m2_CreateDeployment(cfg aws.Config, client *m2.Client) {
	input := &m2.CreateDeploymentInput{
		// ApplicationId: *string, // Required
		// ApplicationVersion: *int32, // Required
		// EnvironmentId: *string, // Required
	}

	if len(_m2ApplicationId) > 0 {
		input.ApplicationId = aws.String(_m2ApplicationId)
	}
	if len(_m2ApplicationVersion) > 0 {
		if err := assignInputField(input, "ApplicationVersion", _m2ApplicationVersion); err != nil {
			log.Errorf("invalid --application-version: %s", err.Error())
			return
		}
	}
	if len(_m2EnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_m2EnvironmentId)
	}
	if len(_m2ClientToken) > 0 {
		input.ClientToken = aws.String(_m2ClientToken)
	}

	if resp, err := client.CreateDeployment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a runtime environment for a given runtime engine.
func m2_CreateEnvironment(cfg aws.Config, client *m2.Client) {
	input := &m2.CreateEnvironmentInput{
		// EngineType: types.EngineType, // Required
		// InstanceType: *string, // Required
		// Name: *string, // Required
	}

	if len(_m2EngineType) > 0 {
		if err := assignInputField(input, "EngineType", _m2EngineType); err != nil {
			log.Errorf("invalid --engine-type: %s", err.Error())
			return
		}
	}
	if len(_m2InstanceType) > 0 {
		input.InstanceType = aws.String(_m2InstanceType)
	}
	if len(_m2Name) > 0 {
		input.Name = aws.String(_m2Name)
	}
	if len(_m2ClientToken) > 0 {
		input.ClientToken = aws.String(_m2ClientToken)
	}
	if len(_m2Description) > 0 {
		input.Description = aws.String(_m2Description)
	}
	if len(_m2EngineVersion) > 0 {
		input.EngineVersion = aws.String(_m2EngineVersion)
	}
	if len(_m2HighAvailabilityConfig) > 0 {
		if err := assignInputField(input, "HighAvailabilityConfig", _m2HighAvailabilityConfig); err != nil {
			log.Errorf("invalid --high-availability-config: %s", err.Error())
			return
		}
	}
	if len(_m2KmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_m2KmsKeyId)
	}
	if len(_m2NetworkType) > 0 {
		if err := assignInputField(input, "NetworkType", _m2NetworkType); err != nil {
			log.Errorf("invalid --network-type: %s", err.Error())
			return
		}
	}
	if len(_m2PreferredMaintenanceWindow) > 0 {
		input.PreferredMaintenanceWindow = aws.String(_m2PreferredMaintenanceWindow)
	}
	if len(_m2PubliclyAccessible) > 0 {
		if err := assignInputField(input, "PubliclyAccessible", _m2PubliclyAccessible); err != nil {
			log.Errorf("invalid --publicly-accessible: %s", err.Error())
			return
		}
	}
	if len(_m2SecurityGroupIds) > 0 {
		input.SecurityGroupIds = append([]string(nil), _m2SecurityGroupIds...)
	}
	if len(_m2StorageConfigurations) > 0 {
		if err := assignInputField(input, "StorageConfigurations", _m2StorageConfigurations); err != nil {
			log.Errorf("invalid --storage-configurations: %s", err.Error())
			return
		}
	}
	if len(_m2SubnetIds) > 0 {
		input.SubnetIds = append([]string(nil), _m2SubnetIds...)
	}
	if len(_m2Tags) > 0 {
		if err := assignInputField(input, "Tags", _m2Tags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateEnvironment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a specific application. You cannot delete a running application.
func m2_DeleteApplication(cfg aws.Config, client *m2.Client) {
	input := &m2.DeleteApplicationInput{
		// ApplicationId: *string, // Required
	}

	if len(_m2ApplicationId) > 0 {
		input.ApplicationId = aws.String(_m2ApplicationId)
	}

	if resp, err := client.DeleteApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a specific application from the specific runtime environment where it
// was previously deployed. You cannot delete a runtime environment using
// DeleteEnvironment if any application has ever been deployed to it. This API
// removes the association of the application with the runtime environment so you
// can delete the environment smoothly.
func m2_DeleteApplicationFromEnvironment(cfg aws.Config, client *m2.Client) {
	input := &m2.DeleteApplicationFromEnvironmentInput{
		// ApplicationId: *string, // Required
		// EnvironmentId: *string, // Required
	}

	if len(_m2ApplicationId) > 0 {
		input.ApplicationId = aws.String(_m2ApplicationId)
	}
	if len(_m2EnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_m2EnvironmentId)
	}

	if resp, err := client.DeleteApplicationFromEnvironment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a specific runtime environment. The environment cannot contain deployed
// applications. If it does, you must delete those applications before you delete
// the environment.
func m2_DeleteEnvironment(cfg aws.Config, client *m2.Client) {
	input := &m2.DeleteEnvironmentInput{
		// EnvironmentId: *string, // Required
	}

	if len(_m2EnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_m2EnvironmentId)
	}

	if resp, err := client.DeleteEnvironment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the details of a specific application.
func m2_GetApplication(cfg aws.Config, client *m2.Client) {
	input := &m2.GetApplicationInput{
		// ApplicationId: *string, // Required
	}

	if len(_m2ApplicationId) > 0 {
		input.ApplicationId = aws.String(_m2ApplicationId)
	}

	if resp, err := client.GetApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns details about a specific version of a specific application.
func m2_GetApplicationVersion(cfg aws.Config, client *m2.Client) {
	input := &m2.GetApplicationVersionInput{
		// ApplicationId: *string, // Required
		// ApplicationVersion: *int32, // Required
	}

	if len(_m2ApplicationId) > 0 {
		input.ApplicationId = aws.String(_m2ApplicationId)
	}
	if len(_m2ApplicationVersion) > 0 {
		if err := assignInputField(input, "ApplicationVersion", _m2ApplicationVersion); err != nil {
			log.Errorf("invalid --application-version: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetApplicationVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the details of a specific batch job execution for a specific application.
func m2_GetBatchJobExecution(cfg aws.Config, client *m2.Client) {
	input := &m2.GetBatchJobExecutionInput{
		// ApplicationId: *string, // Required
		// ExecutionId: *string, // Required
	}

	if len(_m2ApplicationId) > 0 {
		input.ApplicationId = aws.String(_m2ApplicationId)
	}
	if len(_m2ExecutionId) > 0 {
		input.ExecutionId = aws.String(_m2ExecutionId)
	}

	if resp, err := client.GetBatchJobExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the details of a specific data set.
func m2_GetDataSetDetails(cfg aws.Config, client *m2.Client) {
	input := &m2.GetDataSetDetailsInput{
		// ApplicationId: *string, // Required
		// DataSetName: *string, // Required
	}

	if len(_m2ApplicationId) > 0 {
		input.ApplicationId = aws.String(_m2ApplicationId)
	}
	if len(_m2DataSetName) > 0 {
		input.DataSetName = aws.String(_m2DataSetName)
	}

	if resp, err := client.GetDataSetDetails(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the status of a data set import task initiated with the CreateDataSetExportTask operation.
func m2_GetDataSetExportTask(cfg aws.Config, client *m2.Client) {
	input := &m2.GetDataSetExportTaskInput{
		// ApplicationId: *string, // Required
		// TaskId: *string, // Required
	}

	if len(_m2ApplicationId) > 0 {
		input.ApplicationId = aws.String(_m2ApplicationId)
	}
	if len(_m2TaskId) > 0 {
		input.TaskId = aws.String(_m2TaskId)
	}

	if resp, err := client.GetDataSetExportTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the status of a data set import task initiated with the CreateDataSetImportTask operation.
func m2_GetDataSetImportTask(cfg aws.Config, client *m2.Client) {
	input := &m2.GetDataSetImportTaskInput{
		// ApplicationId: *string, // Required
		// TaskId: *string, // Required
	}

	if len(_m2ApplicationId) > 0 {
		input.ApplicationId = aws.String(_m2ApplicationId)
	}
	if len(_m2TaskId) > 0 {
		input.TaskId = aws.String(_m2TaskId)
	}

	if resp, err := client.GetDataSetImportTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets details of a specific deployment with a given deployment identifier.
func m2_GetDeployment(cfg aws.Config, client *m2.Client) {
	input := &m2.GetDeploymentInput{
		// ApplicationId: *string, // Required
		// DeploymentId: *string, // Required
	}

	if len(_m2ApplicationId) > 0 {
		input.ApplicationId = aws.String(_m2ApplicationId)
	}
	if len(_m2DeploymentId) > 0 {
		input.DeploymentId = aws.String(_m2DeploymentId)
	}

	if resp, err := client.GetDeployment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a specific runtime environment.
func m2_GetEnvironment(cfg aws.Config, client *m2.Client) {
	input := &m2.GetEnvironmentInput{
		// EnvironmentId: *string, // Required
	}

	if len(_m2EnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_m2EnvironmentId)
	}

	if resp, err := client.GetEnvironment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a single sign-on URL that can be used to connect to AWS Blu Insights.
func m2_GetSignedBluinsightsUrl(cfg aws.Config, client *m2.Client) {
	input := &m2.GetSignedBluinsightsUrlInput{}

	if resp, err := client.GetSignedBluinsightsUrl(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of the application versions for a specific application.
func m2_ListApplicationVersions(cfg aws.Config, client *m2.Client) {
	input := &m2.ListApplicationVersionsInput{
		// ApplicationId: *string, // Required
	}

	if len(_m2ApplicationId) > 0 {
		input.ApplicationId = aws.String(_m2ApplicationId)
	}
	if len(_m2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _m2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_m2NextToken) > 0 {
		input.NextToken = aws.String(_m2NextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListApplicationVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*m2.ListApplicationVersionsOutput
	p := m2.NewListApplicationVersionsPaginator(client, input)
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

// Lists the applications associated with a specific Amazon Web Services account.
// You can provide the unique identifier of a specific runtime environment in a
// query parameter to see all applications associated with that environment.
func m2_ListApplications(cfg aws.Config, client *m2.Client) {
	input := &m2.ListApplicationsInput{}

	if len(_m2EnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_m2EnvironmentId)
	}
	if len(_m2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _m2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_m2Names) > 0 {
		input.Names = append([]string(nil), _m2Names...)
	}
	if len(_m2NextToken) > 0 {
		input.NextToken = aws.String(_m2NextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListApplications(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*m2.ListApplicationsOutput
	p := m2.NewListApplicationsPaginator(client, input)
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

// Lists all the available batch job definitions based on the batch job resources
// uploaded during the application creation. You can use the batch job definitions
// in the list to start a batch job.
func m2_ListBatchJobDefinitions(cfg aws.Config, client *m2.Client) {
	input := &m2.ListBatchJobDefinitionsInput{
		// ApplicationId: *string, // Required
	}

	if len(_m2ApplicationId) > 0 {
		input.ApplicationId = aws.String(_m2ApplicationId)
	}
	if len(_m2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _m2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_m2NextToken) > 0 {
		input.NextToken = aws.String(_m2NextToken)
	}
	if len(_m2Prefix) > 0 {
		input.Prefix = aws.String(_m2Prefix)
	}

	if disablePaginator() {
		if resp, err := client.ListBatchJobDefinitions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*m2.ListBatchJobDefinitionsOutput
	p := m2.NewListBatchJobDefinitionsPaginator(client, input)
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

// Lists historical, current, and scheduled batch job executions for a specific
// application.
func m2_ListBatchJobExecutions(cfg aws.Config, client *m2.Client) {
	input := &m2.ListBatchJobExecutionsInput{
		// ApplicationId: *string, // Required
	}

	if len(_m2ApplicationId) > 0 {
		input.ApplicationId = aws.String(_m2ApplicationId)
	}
	if len(_m2ExecutionIds) > 0 {
		input.ExecutionIds = append([]string(nil), _m2ExecutionIds...)
	}
	if len(_m2JobName) > 0 {
		input.JobName = aws.String(_m2JobName)
	}
	if len(_m2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _m2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_m2NextToken) > 0 {
		input.NextToken = aws.String(_m2NextToken)
	}
	if len(_m2StartedAfter) > 0 {
		if err := assignInputField(input, "StartedAfter", _m2StartedAfter); err != nil {
			log.Errorf("invalid --started-after: %s", err.Error())
			return
		}
	}
	if len(_m2StartedBefore) > 0 {
		if err := assignInputField(input, "StartedBefore", _m2StartedBefore); err != nil {
			log.Errorf("invalid --started-before: %s", err.Error())
			return
		}
	}
	if len(_m2Status) > 0 {
		if err := assignInputField(input, "Status", _m2Status); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListBatchJobExecutions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*m2.ListBatchJobExecutionsOutput
	p := m2.NewListBatchJobExecutionsPaginator(client, input)
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

// Lists all the job steps for a JCL file to restart a batch job. This is only
// applicable for Micro Focus engine with versions 8.0.6 and above.
func m2_ListBatchJobRestartPoints(cfg aws.Config, client *m2.Client) {
	input := &m2.ListBatchJobRestartPointsInput{
		// ApplicationId: *string, // Required
		// ExecutionId: *string, // Required
	}

	if len(_m2ApplicationId) > 0 {
		input.ApplicationId = aws.String(_m2ApplicationId)
	}
	if len(_m2ExecutionId) > 0 {
		input.ExecutionId = aws.String(_m2ExecutionId)
	}
	if len(_m2AuthSecretsManagerArn) > 0 {
		input.AuthSecretsManagerArn = aws.String(_m2AuthSecretsManagerArn)
	}

	if resp, err := client.ListBatchJobRestartPoints(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the data set exports for the specified application.
func m2_ListDataSetExportHistory(cfg aws.Config, client *m2.Client) {
	input := &m2.ListDataSetExportHistoryInput{
		// ApplicationId: *string, // Required
	}

	if len(_m2ApplicationId) > 0 {
		input.ApplicationId = aws.String(_m2ApplicationId)
	}
	if len(_m2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _m2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_m2NextToken) > 0 {
		input.NextToken = aws.String(_m2NextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDataSetExportHistory(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*m2.ListDataSetExportHistoryOutput
	p := m2.NewListDataSetExportHistoryPaginator(client, input)
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

// Lists the data set imports for the specified application.
func m2_ListDataSetImportHistory(cfg aws.Config, client *m2.Client) {
	input := &m2.ListDataSetImportHistoryInput{
		// ApplicationId: *string, // Required
	}

	if len(_m2ApplicationId) > 0 {
		input.ApplicationId = aws.String(_m2ApplicationId)
	}
	if len(_m2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _m2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_m2NextToken) > 0 {
		input.NextToken = aws.String(_m2NextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDataSetImportHistory(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*m2.ListDataSetImportHistoryOutput
	p := m2.NewListDataSetImportHistoryPaginator(client, input)
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

// Lists the data sets imported for a specific application. In Amazon Web Services
// Mainframe Modernization, data sets are associated with applications deployed on
// runtime environments. This is known as importing data sets. Currently, Amazon
// Web Services Mainframe Modernization can import data sets into catalogs using [CreateDataSetImportTask].
//
// [CreateDataSetImportTask]: https://docs.aws.amazon.com/m2/latest/APIReference/API_CreateDataSetImportTask.html
func m2_ListDataSets(cfg aws.Config, client *m2.Client) {
	input := &m2.ListDataSetsInput{
		// ApplicationId: *string, // Required
	}

	if len(_m2ApplicationId) > 0 {
		input.ApplicationId = aws.String(_m2ApplicationId)
	}
	if len(_m2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _m2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_m2NameFilter) > 0 {
		input.NameFilter = aws.String(_m2NameFilter)
	}
	if len(_m2NextToken) > 0 {
		input.NextToken = aws.String(_m2NextToken)
	}
	if len(_m2Prefix) > 0 {
		input.Prefix = aws.String(_m2Prefix)
	}

	if disablePaginator() {
		if resp, err := client.ListDataSets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*m2.ListDataSetsOutput
	p := m2.NewListDataSetsPaginator(client, input)
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

// Returns a list of all deployments of a specific application. A deployment is a
// combination of a specific application and a specific version of that
// application. Each deployment is mapped to a particular application version.
func m2_ListDeployments(cfg aws.Config, client *m2.Client) {
	input := &m2.ListDeploymentsInput{
		// ApplicationId: *string, // Required
	}

	if len(_m2ApplicationId) > 0 {
		input.ApplicationId = aws.String(_m2ApplicationId)
	}
	if len(_m2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _m2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_m2NextToken) > 0 {
		input.NextToken = aws.String(_m2NextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDeployments(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*m2.ListDeploymentsOutput
	p := m2.NewListDeploymentsPaginator(client, input)
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

// Lists the available engine versions.
func m2_ListEngineVersions(cfg aws.Config, client *m2.Client) {
	input := &m2.ListEngineVersionsInput{}

	if len(_m2EngineType) > 0 {
		if err := assignInputField(input, "EngineType", _m2EngineType); err != nil {
			log.Errorf("invalid --engine-type: %s", err.Error())
			return
		}
	}
	if len(_m2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _m2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_m2NextToken) > 0 {
		input.NextToken = aws.String(_m2NextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListEngineVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*m2.ListEngineVersionsOutput
	p := m2.NewListEngineVersionsPaginator(client, input)
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

// Lists the runtime environments.
func m2_ListEnvironments(cfg aws.Config, client *m2.Client) {
	input := &m2.ListEnvironmentsInput{}

	if len(_m2EngineType) > 0 {
		if err := assignInputField(input, "EngineType", _m2EngineType); err != nil {
			log.Errorf("invalid --engine-type: %s", err.Error())
			return
		}
	}
	if len(_m2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _m2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_m2Names) > 0 {
		input.Names = append([]string(nil), _m2Names...)
	}
	if len(_m2NextToken) > 0 {
		input.NextToken = aws.String(_m2NextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListEnvironments(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*m2.ListEnvironmentsOutput
	p := m2.NewListEnvironmentsPaginator(client, input)
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

// Lists the tags for the specified resource.
func m2_ListTagsForResource(cfg aws.Config, client *m2.Client) {
	input := &m2.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_m2ResourceArn) > 0 {
		input.ResourceArn = aws.String(_m2ResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts an application that is currently stopped.
func m2_StartApplication(cfg aws.Config, client *m2.Client) {
	input := &m2.StartApplicationInput{
		// ApplicationId: *string, // Required
	}

	if len(_m2ApplicationId) > 0 {
		input.ApplicationId = aws.String(_m2ApplicationId)
	}

	if resp, err := client.StartApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a batch job and returns the unique identifier of this execution of the
// batch job. The associated application must be running in order to start the
// batch job.
func m2_StartBatchJob(cfg aws.Config, client *m2.Client) {
	input := &m2.StartBatchJobInput{
		// ApplicationId: *string, // Required
		// BatchJobIdentifier: types.BatchJobIdentifier, // Required
	}

	if len(_m2ApplicationId) > 0 {
		input.ApplicationId = aws.String(_m2ApplicationId)
	}
	if len(_m2BatchJobIdentifier) > 0 {
		if err := assignInputField(input, "BatchJobIdentifier", _m2BatchJobIdentifier); err != nil {
			log.Errorf("invalid --batch-job-identifier: %s", err.Error())
			return
		}
	}
	if len(_m2AuthSecretsManagerArn) > 0 {
		input.AuthSecretsManagerArn = aws.String(_m2AuthSecretsManagerArn)
	}
	if len(_m2JobParams) > 0 {
		if err := assignInputField(input, "JobParams", _m2JobParams); err != nil {
			log.Errorf("invalid --job-params: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartBatchJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a running application.
func m2_StopApplication(cfg aws.Config, client *m2.Client) {
	input := &m2.StopApplicationInput{
		// ApplicationId: *string, // Required
	}

	if len(_m2ApplicationId) > 0 {
		input.ApplicationId = aws.String(_m2ApplicationId)
	}
	if len(_m2ForceStop) > 0 {
		if err := assignInputField(input, "ForceStop", _m2ForceStop); err != nil {
			log.Errorf("invalid --force-stop: %s", err.Error())
			return
		}
	}

	if resp, err := client.StopApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds one or more tags to the specified resource.
func m2_TagResource(cfg aws.Config, client *m2.Client) {
	input := &m2.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_m2ResourceArn) > 0 {
		input.ResourceArn = aws.String(_m2ResourceArn)
	}
	if len(_m2Tags) > 0 {
		if err := assignInputField(input, "Tags", _m2Tags); err != nil {
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

// Removes one or more tags from the specified resource.
func m2_UntagResource(cfg aws.Config, client *m2.Client) {
	input := &m2.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_m2ResourceArn) > 0 {
		input.ResourceArn = aws.String(_m2ResourceArn)
	}
	if len(_m2TagKeys) > 0 {
		input.TagKeys = append([]string(nil), _m2TagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an application and creates a new version.
func m2_UpdateApplication(cfg aws.Config, client *m2.Client) {
	input := &m2.UpdateApplicationInput{
		// ApplicationId: *string, // Required
		// CurrentApplicationVersion: *int32, // Required
	}

	if len(_m2ApplicationId) > 0 {
		input.ApplicationId = aws.String(_m2ApplicationId)
	}
	if len(_m2CurrentApplicationVersion) > 0 {
		if err := assignInputField(input, "CurrentApplicationVersion", _m2CurrentApplicationVersion); err != nil {
			log.Errorf("invalid --current-application-version: %s", err.Error())
			return
		}
	}
	if len(_m2Definition) > 0 {
		if err := assignInputField(input, "Definition", _m2Definition); err != nil {
			log.Errorf("invalid --definition: %s", err.Error())
			return
		}
	}
	if len(_m2Description) > 0 {
		input.Description = aws.String(_m2Description)
	}

	if resp, err := client.UpdateApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration details for a specific runtime environment.
func m2_UpdateEnvironment(cfg aws.Config, client *m2.Client) {
	input := &m2.UpdateEnvironmentInput{
		// EnvironmentId: *string, // Required
	}

	if len(_m2EnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_m2EnvironmentId)
	}
	if len(_m2ApplyDuringMaintenanceWindow) > 0 {
		if err := assignInputField(input, "ApplyDuringMaintenanceWindow", _m2ApplyDuringMaintenanceWindow); err != nil {
			log.Errorf("invalid --apply-during-maintenance-window: %s", err.Error())
			return
		}
	}
	if len(_m2DesiredCapacity) > 0 {
		if err := assignInputField(input, "DesiredCapacity", _m2DesiredCapacity); err != nil {
			log.Errorf("invalid --desired-capacity: %s", err.Error())
			return
		}
	}
	if len(_m2EngineVersion) > 0 {
		input.EngineVersion = aws.String(_m2EngineVersion)
	}
	if len(_m2ForceUpdate) > 0 {
		if err := assignInputField(input, "ForceUpdate", _m2ForceUpdate); err != nil {
			log.Errorf("invalid --force-update: %s", err.Error())
			return
		}
	}
	if len(_m2InstanceType) > 0 {
		input.InstanceType = aws.String(_m2InstanceType)
	}
	if len(_m2PreferredMaintenanceWindow) > 0 {
		input.PreferredMaintenanceWindow = aws.String(_m2PreferredMaintenanceWindow)
	}

	if resp, err := client.UpdateEnvironment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_m2Cmd)
	_m2Cmd.Flags().SortFlags = false

	_m2Cmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_m2Cmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_m2Cmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_m2Cmd.Flags().StringVarP(&_m2ApplicationId, "application-id", "", "", "Application ID")
	_m2Cmd.Flags().StringVarP(&_m2ApplicationVersion, "application-version", "", "", "Application Version")
	_m2Cmd.Flags().StringVarP(&_m2ApplyDuringMaintenanceWindow, "apply-during-maintenance-window", "", "", "Apply During Maintenance Window")
	_m2Cmd.Flags().StringVarP(&_m2AuthSecretsManagerArn, "auth-secrets-manager-arn", "", "", "Auth Secrets Manager ARN")
	_m2Cmd.Flags().StringVarP(&_m2BatchJobIdentifier, "batch-job-identifier", "", "", "Batch Job Identifier")
	_m2Cmd.Flags().StringVarP(&_m2ClientToken, "client-token", "", "", "Client Token")
	_m2Cmd.Flags().StringVarP(&_m2CurrentApplicationVersion, "current-application-version", "", "", "Current Application Version")
	_m2Cmd.Flags().StringVarP(&_m2DataSetName, "data-set-name", "", "", "Data Set Name")
	_m2Cmd.Flags().StringVarP(&_m2Definition, "definition", "", "", "Definition")
	_m2Cmd.Flags().StringVarP(&_m2DeploymentId, "deployment-id", "", "", "Deployment ID")
	_m2Cmd.Flags().StringVarP(&_m2Description, "description", "", "", "Description")
	_m2Cmd.Flags().StringVarP(&_m2DesiredCapacity, "desired-capacity", "", "", "Desired Capacity")
	_m2Cmd.Flags().StringVarP(&_m2EngineType, "engine-type", "", "", "Engine Type")
	_m2Cmd.Flags().StringVarP(&_m2EngineVersion, "engine-version", "", "", "Engine Version")
	_m2Cmd.Flags().StringVarP(&_m2EnvironmentId, "environment-id", "", "", "Environment ID")
	_m2Cmd.Flags().StringVarP(&_m2ExecutionId, "execution-id", "", "", "Execution ID")
	_m2Cmd.Flags().StringSliceVarP(&_m2ExecutionIds, "execution-ids", "", nil, "Execution Ids")
	_m2Cmd.Flags().StringVarP(&_m2ExportConfig, "export-config", "", "", "Export Config")
	_m2Cmd.Flags().StringVarP(&_m2ForceStop, "force-stop", "", "", "Force Stop")
	_m2Cmd.Flags().StringVarP(&_m2ForceUpdate, "force-update", "", "", "Force Update")
	_m2Cmd.Flags().StringVarP(&_m2HighAvailabilityConfig, "high-availability-config", "", "", "High Availability Config")
	_m2Cmd.Flags().StringVarP(&_m2ImportConfig, "import-config", "", "", "Import Config")
	_m2Cmd.Flags().StringVarP(&_m2InstanceType, "instance-type", "", "", "Instance Type")
	_m2Cmd.Flags().StringVarP(&_m2JobName, "job-name", "", "", "Job Name")
	_m2Cmd.Flags().StringVarP(&_m2JobParams, "job-params", "", "", "Job Params")
	_m2Cmd.Flags().StringVarP(&_m2KmsKeyId, "kms-key-id", "", "", "KMS Key ID")
	_m2Cmd.Flags().StringVarP(&_m2MaxResults, "max-results", "", "", "Max Results")
	_m2Cmd.Flags().StringVarP(&_m2Name, "name", "", "", "Name")
	_m2Cmd.Flags().StringVarP(&_m2NameFilter, "name-filter", "", "", "Name Filter")
	_m2Cmd.Flags().StringSliceVarP(&_m2Names, "names", "", nil, "Names")
	_m2Cmd.Flags().StringVarP(&_m2NetworkType, "network-type", "", "", "Network Type")
	_m2Cmd.Flags().StringVarP(&_m2NextToken, "next-token", "", "", "Next Token")
	_m2Cmd.Flags().StringVarP(&_m2PreferredMaintenanceWindow, "preferred-maintenance-window", "", "", "Preferred Maintenance Window")
	_m2Cmd.Flags().StringVarP(&_m2Prefix, "prefix", "", "", "Prefix")
	_m2Cmd.Flags().StringVarP(&_m2PubliclyAccessible, "publicly-accessible", "", "", "Publicly Accessible")
	_m2Cmd.Flags().StringVarP(&_m2ResourceArn, "resource-arn", "", "", "Resource ARN")
	_m2Cmd.Flags().StringVarP(&_m2RoleArn, "role-arn", "", "", "Role ARN")
	_m2Cmd.Flags().StringSliceVarP(&_m2SecurityGroupIds, "security-group-ids", "", nil, "Security Group Ids")
	_m2Cmd.Flags().StringVarP(&_m2StartedAfter, "started-after", "", "", "Started After")
	_m2Cmd.Flags().StringVarP(&_m2StartedBefore, "started-before", "", "", "Started Before")
	_m2Cmd.Flags().StringVarP(&_m2Status, "status", "", "", "Status")
	_m2Cmd.Flags().StringVarP(&_m2StorageConfigurations, "storage-configurations", "", "", "Storage Configurations")
	_m2Cmd.Flags().StringSliceVarP(&_m2SubnetIds, "subnet-ids", "", nil, "Subnet Ids")
	_m2Cmd.Flags().StringSliceVarP(&_m2TagKeys, "tag-keys", "", nil, "Tag Keys")
	_m2Cmd.Flags().StringVarP(&_m2Tags, "tags", "", "", "Tags")
	_m2Cmd.Flags().StringVarP(&_m2TaskId, "task-id", "", "", "Task ID")

	_m2Cmd.Flags().BoolVarP(&_m2CancelBatchJobExecution, "cancel-batch-job-execution", "", false, "Cancel Batch Job Execution")
	_m2Cmd.Flags().BoolVarP(&_m2CreateApplication, "create-application", "", false, "Create Application")
	_m2Cmd.Flags().BoolVarP(&_m2CreateDataSetExportTask, "create-data-set-export-task", "", false, "Create Data Set Export Task")
	_m2Cmd.Flags().BoolVarP(&_m2CreateDataSetImportTask, "create-data-set-import-task", "", false, "Create Data Set Import Task")
	_m2Cmd.Flags().BoolVarP(&_m2CreateDeployment, "create-deployment", "", false, "Create Deployment")
	_m2Cmd.Flags().BoolVarP(&_m2CreateEnvironment, "create-environment", "", false, "Create Environment")
	_m2Cmd.Flags().BoolVarP(&_m2DeleteApplication, "delete-application", "", false, "Delete Application")
	_m2Cmd.Flags().BoolVarP(&_m2DeleteApplicationFromEnvironment, "delete-application-from-environment", "", false, "Delete Application From Environment")
	_m2Cmd.Flags().BoolVarP(&_m2DeleteEnvironment, "delete-environment", "", false, "Delete Environment")
	_m2Cmd.Flags().BoolVarP(&_m2GetApplication, "get-application", "", false, "Get Application")
	_m2Cmd.Flags().BoolVarP(&_m2GetApplicationVersion, "get-application-version", "", false, "Get Application Version")
	_m2Cmd.Flags().BoolVarP(&_m2GetBatchJobExecution, "get-batch-job-execution", "", false, "Get Batch Job Execution")
	_m2Cmd.Flags().BoolVarP(&_m2GetDataSetDetails, "get-data-set-details", "", false, "Get Data Set Details")
	_m2Cmd.Flags().BoolVarP(&_m2GetDataSetExportTask, "get-data-set-export-task", "", false, "Get Data Set Export Task")
	_m2Cmd.Flags().BoolVarP(&_m2GetDataSetImportTask, "get-data-set-import-task", "", false, "Get Data Set Import Task")
	_m2Cmd.Flags().BoolVarP(&_m2GetDeployment, "get-deployment", "", false, "Get Deployment")
	_m2Cmd.Flags().BoolVarP(&_m2GetEnvironment, "get-environment", "", false, "Get Environment")
	_m2Cmd.Flags().BoolVarP(&_m2GetSignedBluinsightsUrl, "get-signed-bluinsights-url", "", false, "Get Signed Bluinsights URL")
	_m2Cmd.Flags().BoolVarP(&_m2ListApplicationVersions, "list-application-versions", "", false, "List Application Versions")
	_m2Cmd.Flags().BoolVarP(&_m2ListApplications, "list-applications", "", false, "List Applications")
	_m2Cmd.Flags().BoolVarP(&_m2ListBatchJobDefinitions, "list-batch-job-definitions", "", false, "List Batch Job Definitions")
	_m2Cmd.Flags().BoolVarP(&_m2ListBatchJobExecutions, "list-batch-job-executions", "", false, "List Batch Job Executions")
	_m2Cmd.Flags().BoolVarP(&_m2ListBatchJobRestartPoints, "list-batch-job-restart-points", "", false, "List Batch Job Restart Points")
	_m2Cmd.Flags().BoolVarP(&_m2ListDataSetExportHistory, "list-data-set-export-history", "", false, "List Data Set Export History")
	_m2Cmd.Flags().BoolVarP(&_m2ListDataSetImportHistory, "list-data-set-import-history", "", false, "List Data Set Import History")
	_m2Cmd.Flags().BoolVarP(&_m2ListDataSets, "list-data-sets", "", false, "List Data Sets")
	_m2Cmd.Flags().BoolVarP(&_m2ListDeployments, "list-deployments", "", false, "List Deployments")
	_m2Cmd.Flags().BoolVarP(&_m2ListEngineVersions, "list-engine-versions", "", false, "List Engine Versions")
	_m2Cmd.Flags().BoolVarP(&_m2ListEnvironments, "list-environments", "", false, "List Environments")
	_m2Cmd.Flags().BoolVarP(&_m2ListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_m2Cmd.Flags().BoolVarP(&_m2StartApplication, "start-application", "", false, "Start Application")
	_m2Cmd.Flags().BoolVarP(&_m2StartBatchJob, "start-batch-job", "", false, "Start Batch Job")
	_m2Cmd.Flags().BoolVarP(&_m2StopApplication, "stop-application", "", false, "Stop Application")
	_m2Cmd.Flags().BoolVarP(&_m2TagResource, "tag-resource", "", false, "Tag Resource")
	_m2Cmd.Flags().BoolVarP(&_m2UntagResource, "untag-resource", "", false, "Untag Resource")
	_m2Cmd.Flags().BoolVarP(&_m2UpdateApplication, "update-application", "", false, "Update Application")
	_m2Cmd.Flags().BoolVarP(&_m2UpdateEnvironment, "update-environment", "", false, "Update Environment")

}
