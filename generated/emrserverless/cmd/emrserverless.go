package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/emrserverless"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// emrserverlessCmd represents the emrserverless command
var _emrserverlessCmd = &cobra.Command{
	Use:   "emrserverless",
	Short: "AWS emrserverless CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := emrserverless.NewFromConfig(cfg)
		if _emrserverlessCancelJobRun {
			emrserverless_CancelJobRun(cfg, client)
			return
		}
		if _emrserverlessCreateApplication {
			emrserverless_CreateApplication(cfg, client)
			return
		}
		if _emrserverlessDeleteApplication {
			emrserverless_DeleteApplication(cfg, client)
			return
		}
		if _emrserverlessGetApplication {
			emrserverless_GetApplication(cfg, client)
			return
		}
		if _emrserverlessGetDashboardForJobRun {
			emrserverless_GetDashboardForJobRun(cfg, client)
			return
		}
		if _emrserverlessGetJobRun {
			emrserverless_GetJobRun(cfg, client)
			return
		}
		if _emrserverlessListApplications {
			emrserverless_ListApplications(cfg, client)
			return
		}
		if _emrserverlessListJobRunAttempts {
			emrserverless_ListJobRunAttempts(cfg, client)
			return
		}
		if _emrserverlessListJobRuns {
			emrserverless_ListJobRuns(cfg, client)
			return
		}
		if _emrserverlessListTagsForResource {
			emrserverless_ListTagsForResource(cfg, client)
			return
		}
		if _emrserverlessStartApplication {
			emrserverless_StartApplication(cfg, client)
			return
		}
		if _emrserverlessStartJobRun {
			emrserverless_StartJobRun(cfg, client)
			return
		}
		if _emrserverlessStopApplication {
			emrserverless_StopApplication(cfg, client)
			return
		}
		if _emrserverlessTagResource {
			emrserverless_TagResource(cfg, client)
			return
		}
		if _emrserverlessUntagResource {
			emrserverless_UntagResource(cfg, client)
			return
		}
		if _emrserverlessUpdateApplication {
			emrserverless_UpdateApplication(cfg, client)
			return
		}

	},
}

var (
	_emrserverlessCancelJobRun          bool
	_emrserverlessCreateApplication     bool
	_emrserverlessDeleteApplication     bool
	_emrserverlessGetApplication        bool
	_emrserverlessGetDashboardForJobRun bool
	_emrserverlessGetJobRun             bool
	_emrserverlessListApplications      bool
	_emrserverlessListJobRunAttempts    bool
	_emrserverlessListJobRuns           bool
	_emrserverlessListTagsForResource   bool
	_emrserverlessStartApplication      bool
	_emrserverlessStartJobRun           bool
	_emrserverlessStopApplication       bool
	_emrserverlessTagResource           bool
	_emrserverlessUntagResource         bool
	_emrserverlessUpdateApplication     bool

	_emrserverlessAccessSystemProfileLogs             string
	_emrserverlessApplicationId                       string
	_emrserverlessArchitecture                        string
	_emrserverlessAttempt                             string
	_emrserverlessAutoStartConfiguration              string
	_emrserverlessAutoStopConfiguration               string
	_emrserverlessClientToken                         string
	_emrserverlessConfigurationOverrides              string
	_emrserverlessCreatedAtAfter                      string
	_emrserverlessCreatedAtBefore                     string
	_emrserverlessDiskEncryptionConfiguration         string
	_emrserverlessExecutionIamPolicy                  string
	_emrserverlessExecutionRoleArn                    string
	_emrserverlessExecutionTimeoutMinutes             string
	_emrserverlessIdentityCenterConfiguration         string
	_emrserverlessImageConfiguration                  string
	_emrserverlessInitialCapacity                     string
	_emrserverlessInteractiveConfiguration            string
	_emrserverlessJobDriver                           string
	_emrserverlessJobLevelCostAllocationConfiguration string
	_emrserverlessJobRunId                            string
	_emrserverlessMaxResults                          string
	_emrserverlessMaximumCapacity                     string
	_emrserverlessMode                                string
	_emrserverlessMonitoringConfiguration             string
	_emrserverlessName                                string
	_emrserverlessNetworkConfiguration                string
	_emrserverlessNextToken                           string
	_emrserverlessReleaseLabel                        string
	_emrserverlessResourceArn                         string
	_emrserverlessRetryPolicy                         string
	_emrserverlessRuntimeConfiguration                string
	_emrserverlessSchedulerConfiguration              string
	_emrserverlessShutdownGracePeriodInSeconds        string
	_emrserverlessStates                              string
	_emrserverlessTagKeys                             []string
	_emrserverlessTags                                string
	_emrserverlessType                                string
	_emrserverlessWorkerTypeSpecifications            string
)

// Cancels a job run.
func emrserverless_CancelJobRun(cfg aws.Config, client *emrserverless.Client) {
	input := &emrserverless.CancelJobRunInput{
		// ApplicationId: *string, // Required
		// JobRunId: *string, // Required
	}

	if len(_emrserverlessApplicationId) > 0 {
		input.ApplicationId = aws.String(_emrserverlessApplicationId)
	}
	if len(_emrserverlessJobRunId) > 0 {
		input.JobRunId = aws.String(_emrserverlessJobRunId)
	}
	if len(_emrserverlessShutdownGracePeriodInSeconds) > 0 {
		if err := assignInputField(input, "ShutdownGracePeriodInSeconds", _emrserverlessShutdownGracePeriodInSeconds); err != nil {
			log.Errorf("invalid --shutdown-grace-period-in-seconds: %s", err.Error())
			return
		}
	}

	if resp, err := client.CancelJobRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an application.
func emrserverless_CreateApplication(cfg aws.Config, client *emrserverless.Client) {
	input := &emrserverless.CreateApplicationInput{
		// ClientToken: *string, // Required
		// ReleaseLabel: *string, // Required
		// Type: *string, // Required
	}

	if len(_emrserverlessClientToken) > 0 {
		input.ClientToken = aws.String(_emrserverlessClientToken)
	}
	if len(_emrserverlessReleaseLabel) > 0 {
		input.ReleaseLabel = aws.String(_emrserverlessReleaseLabel)
	}
	if len(_emrserverlessType) > 0 {
		input.Type = aws.String(_emrserverlessType)
	}
	if len(_emrserverlessArchitecture) > 0 {
		if err := assignInputField(input, "Architecture", _emrserverlessArchitecture); err != nil {
			log.Errorf("invalid --architecture: %s", err.Error())
			return
		}
	}
	if len(_emrserverlessAutoStartConfiguration) > 0 {
		if err := assignInputField(input, "AutoStartConfiguration", _emrserverlessAutoStartConfiguration); err != nil {
			log.Errorf("invalid --auto-start-configuration: %s", err.Error())
			return
		}
	}
	if len(_emrserverlessAutoStopConfiguration) > 0 {
		if err := assignInputField(input, "AutoStopConfiguration", _emrserverlessAutoStopConfiguration); err != nil {
			log.Errorf("invalid --auto-stop-configuration: %s", err.Error())
			return
		}
	}
	if len(_emrserverlessDiskEncryptionConfiguration) > 0 {
		if err := assignInputField(input, "DiskEncryptionConfiguration", _emrserverlessDiskEncryptionConfiguration); err != nil {
			log.Errorf("invalid --disk-encryption-configuration: %s", err.Error())
			return
		}
	}
	if len(_emrserverlessIdentityCenterConfiguration) > 0 {
		if err := assignInputField(input, "IdentityCenterConfiguration", _emrserverlessIdentityCenterConfiguration); err != nil {
			log.Errorf("invalid --identity-center-configuration: %s", err.Error())
			return
		}
	}
	if len(_emrserverlessImageConfiguration) > 0 {
		if err := assignInputField(input, "ImageConfiguration", _emrserverlessImageConfiguration); err != nil {
			log.Errorf("invalid --image-configuration: %s", err.Error())
			return
		}
	}
	if len(_emrserverlessInitialCapacity) > 0 {
		if err := assignInputField(input, "InitialCapacity", _emrserverlessInitialCapacity); err != nil {
			log.Errorf("invalid --initial-capacity: %s", err.Error())
			return
		}
	}
	if len(_emrserverlessInteractiveConfiguration) > 0 {
		if err := assignInputField(input, "InteractiveConfiguration", _emrserverlessInteractiveConfiguration); err != nil {
			log.Errorf("invalid --interactive-configuration: %s", err.Error())
			return
		}
	}
	if len(_emrserverlessJobLevelCostAllocationConfiguration) > 0 {
		if err := assignInputField(input, "JobLevelCostAllocationConfiguration", _emrserverlessJobLevelCostAllocationConfiguration); err != nil {
			log.Errorf("invalid --job-level-cost-allocation-configuration: %s", err.Error())
			return
		}
	}
	if len(_emrserverlessMaximumCapacity) > 0 {
		if err := assignInputField(input, "MaximumCapacity", _emrserverlessMaximumCapacity); err != nil {
			log.Errorf("invalid --maximum-capacity: %s", err.Error())
			return
		}
	}
	if len(_emrserverlessMonitoringConfiguration) > 0 {
		if err := assignInputField(input, "MonitoringConfiguration", _emrserverlessMonitoringConfiguration); err != nil {
			log.Errorf("invalid --monitoring-configuration: %s", err.Error())
			return
		}
	}
	if len(_emrserverlessName) > 0 {
		input.Name = aws.String(_emrserverlessName)
	}
	if len(_emrserverlessNetworkConfiguration) > 0 {
		if err := assignInputField(input, "NetworkConfiguration", _emrserverlessNetworkConfiguration); err != nil {
			log.Errorf("invalid --network-configuration: %s", err.Error())
			return
		}
	}
	if len(_emrserverlessRuntimeConfiguration) > 0 {
		if err := assignInputField(input, "RuntimeConfiguration", _emrserverlessRuntimeConfiguration); err != nil {
			log.Errorf("invalid --runtime-configuration: %s", err.Error())
			return
		}
	}
	if len(_emrserverlessSchedulerConfiguration) > 0 {
		if err := assignInputField(input, "SchedulerConfiguration", _emrserverlessSchedulerConfiguration); err != nil {
			log.Errorf("invalid --scheduler-configuration: %s", err.Error())
			return
		}
	}
	if len(_emrserverlessTags) > 0 {
		if err := assignInputField(input, "Tags", _emrserverlessTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_emrserverlessWorkerTypeSpecifications) > 0 {
		if err := assignInputField(input, "WorkerTypeSpecifications", _emrserverlessWorkerTypeSpecifications); err != nil {
			log.Errorf("invalid --worker-type-specifications: %s", err.Error())
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

// Deletes an application. An application has to be in a stopped or created state
// in order to be deleted.
func emrserverless_DeleteApplication(cfg aws.Config, client *emrserverless.Client) {
	input := &emrserverless.DeleteApplicationInput{
		// ApplicationId: *string, // Required
	}

	if len(_emrserverlessApplicationId) > 0 {
		input.ApplicationId = aws.String(_emrserverlessApplicationId)
	}

	if resp, err := client.DeleteApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Displays detailed information about a specified application.
func emrserverless_GetApplication(cfg aws.Config, client *emrserverless.Client) {
	input := &emrserverless.GetApplicationInput{
		// ApplicationId: *string, // Required
	}

	if len(_emrserverlessApplicationId) > 0 {
		input.ApplicationId = aws.String(_emrserverlessApplicationId)
	}

	if resp, err := client.GetApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates and returns a URL that you can use to access the application UIs for a
// job run.
//
// For jobs in a running state, the application UI is a live user interface such
// as the Spark or Tez web UI. For completed jobs, the application UI is a
// persistent application user interface such as the Spark History Server or
// persistent Tez UI.
//
// The URL is valid for one hour after you generate it. To access the application
// UI after that hour elapses, you must invoke the API again to generate a new URL.
func emrserverless_GetDashboardForJobRun(cfg aws.Config, client *emrserverless.Client) {
	input := &emrserverless.GetDashboardForJobRunInput{
		// ApplicationId: *string, // Required
		// JobRunId: *string, // Required
	}

	if len(_emrserverlessApplicationId) > 0 {
		input.ApplicationId = aws.String(_emrserverlessApplicationId)
	}
	if len(_emrserverlessJobRunId) > 0 {
		input.JobRunId = aws.String(_emrserverlessJobRunId)
	}
	if len(_emrserverlessAccessSystemProfileLogs) > 0 {
		if err := assignInputField(input, "AccessSystemProfileLogs", _emrserverlessAccessSystemProfileLogs); err != nil {
			log.Errorf("invalid --access-system-profile-logs: %s", err.Error())
			return
		}
	}
	if len(_emrserverlessAttempt) > 0 {
		if err := assignInputField(input, "Attempt", _emrserverlessAttempt); err != nil {
			log.Errorf("invalid --attempt: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetDashboardForJobRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Displays detailed information about a job run.
func emrserverless_GetJobRun(cfg aws.Config, client *emrserverless.Client) {
	input := &emrserverless.GetJobRunInput{
		// ApplicationId: *string, // Required
		// JobRunId: *string, // Required
	}

	if len(_emrserverlessApplicationId) > 0 {
		input.ApplicationId = aws.String(_emrserverlessApplicationId)
	}
	if len(_emrserverlessJobRunId) > 0 {
		input.JobRunId = aws.String(_emrserverlessJobRunId)
	}
	if len(_emrserverlessAttempt) > 0 {
		if err := assignInputField(input, "Attempt", _emrserverlessAttempt); err != nil {
			log.Errorf("invalid --attempt: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetJobRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists applications based on a set of parameters.
func emrserverless_ListApplications(cfg aws.Config, client *emrserverless.Client) {
	input := &emrserverless.ListApplicationsInput{}

	if len(_emrserverlessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _emrserverlessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_emrserverlessNextToken) > 0 {
		input.NextToken = aws.String(_emrserverlessNextToken)
	}
	if len(_emrserverlessStates) > 0 {
		if err := assignInputField(input, "States", _emrserverlessStates); err != nil {
			log.Errorf("invalid --states: %s", err.Error())
			return
		}
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

	var results []*emrserverless.ListApplicationsOutput
	p := emrserverless.NewListApplicationsPaginator(client, input)
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

// Lists all attempt of a job run.
func emrserverless_ListJobRunAttempts(cfg aws.Config, client *emrserverless.Client) {
	input := &emrserverless.ListJobRunAttemptsInput{
		// ApplicationId: *string, // Required
		// JobRunId: *string, // Required
	}

	if len(_emrserverlessApplicationId) > 0 {
		input.ApplicationId = aws.String(_emrserverlessApplicationId)
	}
	if len(_emrserverlessJobRunId) > 0 {
		input.JobRunId = aws.String(_emrserverlessJobRunId)
	}
	if len(_emrserverlessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _emrserverlessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_emrserverlessNextToken) > 0 {
		input.NextToken = aws.String(_emrserverlessNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListJobRunAttempts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*emrserverless.ListJobRunAttemptsOutput
	p := emrserverless.NewListJobRunAttemptsPaginator(client, input)
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

// Lists job runs based on a set of parameters.
func emrserverless_ListJobRuns(cfg aws.Config, client *emrserverless.Client) {
	input := &emrserverless.ListJobRunsInput{
		// ApplicationId: *string, // Required
	}

	if len(_emrserverlessApplicationId) > 0 {
		input.ApplicationId = aws.String(_emrserverlessApplicationId)
	}
	if len(_emrserverlessCreatedAtAfter) > 0 {
		if err := assignInputField(input, "CreatedAtAfter", _emrserverlessCreatedAtAfter); err != nil {
			log.Errorf("invalid --created-at-after: %s", err.Error())
			return
		}
	}
	if len(_emrserverlessCreatedAtBefore) > 0 {
		if err := assignInputField(input, "CreatedAtBefore", _emrserverlessCreatedAtBefore); err != nil {
			log.Errorf("invalid --created-at-before: %s", err.Error())
			return
		}
	}
	if len(_emrserverlessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _emrserverlessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_emrserverlessMode) > 0 {
		if err := assignInputField(input, "Mode", _emrserverlessMode); err != nil {
			log.Errorf("invalid --mode: %s", err.Error())
			return
		}
	}
	if len(_emrserverlessNextToken) > 0 {
		input.NextToken = aws.String(_emrserverlessNextToken)
	}
	if len(_emrserverlessStates) > 0 {
		if err := assignInputField(input, "States", _emrserverlessStates); err != nil {
			log.Errorf("invalid --states: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListJobRuns(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*emrserverless.ListJobRunsOutput
	p := emrserverless.NewListJobRunsPaginator(client, input)
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

// Lists the tags assigned to the resources.
func emrserverless_ListTagsForResource(cfg aws.Config, client *emrserverless.Client) {
	input := &emrserverless.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_emrserverlessResourceArn) > 0 {
		input.ResourceArn = aws.String(_emrserverlessResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a specified application and initializes initial capacity if configured.
func emrserverless_StartApplication(cfg aws.Config, client *emrserverless.Client) {
	input := &emrserverless.StartApplicationInput{
		// ApplicationId: *string, // Required
	}

	if len(_emrserverlessApplicationId) > 0 {
		input.ApplicationId = aws.String(_emrserverlessApplicationId)
	}

	if resp, err := client.StartApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a job run.
func emrserverless_StartJobRun(cfg aws.Config, client *emrserverless.Client) {
	input := &emrserverless.StartJobRunInput{
		// ApplicationId: *string, // Required
		// ClientToken: *string, // Required
		// ExecutionRoleArn: *string, // Required
	}

	if len(_emrserverlessApplicationId) > 0 {
		input.ApplicationId = aws.String(_emrserverlessApplicationId)
	}
	if len(_emrserverlessClientToken) > 0 {
		input.ClientToken = aws.String(_emrserverlessClientToken)
	}
	if len(_emrserverlessExecutionRoleArn) > 0 {
		input.ExecutionRoleArn = aws.String(_emrserverlessExecutionRoleArn)
	}
	if len(_emrserverlessConfigurationOverrides) > 0 {
		if err := assignInputField(input, "ConfigurationOverrides", _emrserverlessConfigurationOverrides); err != nil {
			log.Errorf("invalid --configuration-overrides: %s", err.Error())
			return
		}
	}
	if len(_emrserverlessExecutionIamPolicy) > 0 {
		if err := assignInputField(input, "ExecutionIamPolicy", _emrserverlessExecutionIamPolicy); err != nil {
			log.Errorf("invalid --execution-iam-policy: %s", err.Error())
			return
		}
	}
	if len(_emrserverlessExecutionTimeoutMinutes) > 0 {
		if err := assignInputField(input, "ExecutionTimeoutMinutes", _emrserverlessExecutionTimeoutMinutes); err != nil {
			log.Errorf("invalid --execution-timeout-minutes: %s", err.Error())
			return
		}
	}
	if len(_emrserverlessJobDriver) > 0 {
		if err := assignInputField(input, "JobDriver", _emrserverlessJobDriver); err != nil {
			log.Errorf("invalid --job-driver: %s", err.Error())
			return
		}
	}
	if len(_emrserverlessMode) > 0 {
		if err := assignInputField(input, "Mode", _emrserverlessMode); err != nil {
			log.Errorf("invalid --mode: %s", err.Error())
			return
		}
	}
	if len(_emrserverlessName) > 0 {
		input.Name = aws.String(_emrserverlessName)
	}
	if len(_emrserverlessRetryPolicy) > 0 {
		if err := assignInputField(input, "RetryPolicy", _emrserverlessRetryPolicy); err != nil {
			log.Errorf("invalid --retry-policy: %s", err.Error())
			return
		}
	}
	if len(_emrserverlessTags) > 0 {
		if err := assignInputField(input, "Tags", _emrserverlessTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartJobRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a specified application and releases initial capacity if configured. All
// scheduled and running jobs must be completed or cancelled before stopping an
// application.
func emrserverless_StopApplication(cfg aws.Config, client *emrserverless.Client) {
	input := &emrserverless.StopApplicationInput{
		// ApplicationId: *string, // Required
	}

	if len(_emrserverlessApplicationId) > 0 {
		input.ApplicationId = aws.String(_emrserverlessApplicationId)
	}

	if resp, err := client.StopApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Assigns tags to resources. A tag is a label that you assign to an Amazon Web
// Services resource. Each tag consists of a key and an optional value, both of
// which you define. Tags enable you to categorize your Amazon Web Services
// resources by attributes such as purpose, owner, or environment. When you have
// many resources of the same type, you can quickly identify a specific resource
// based on the tags you've assigned to it.
func emrserverless_TagResource(cfg aws.Config, client *emrserverless.Client) {
	input := &emrserverless.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_emrserverlessResourceArn) > 0 {
		input.ResourceArn = aws.String(_emrserverlessResourceArn)
	}
	if len(_emrserverlessTags) > 0 {
		if err := assignInputField(input, "Tags", _emrserverlessTags); err != nil {
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

// Removes tags from resources.
func emrserverless_UntagResource(cfg aws.Config, client *emrserverless.Client) {
	input := &emrserverless.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_emrserverlessResourceArn) > 0 {
		input.ResourceArn = aws.String(_emrserverlessResourceArn)
	}
	if len(_emrserverlessTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _emrserverlessTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a specified application. An application has to be in a stopped or
// created state in order to be updated.
func emrserverless_UpdateApplication(cfg aws.Config, client *emrserverless.Client) {
	input := &emrserverless.UpdateApplicationInput{
		// ApplicationId: *string, // Required
		// ClientToken: *string, // Required
	}

	if len(_emrserverlessApplicationId) > 0 {
		input.ApplicationId = aws.String(_emrserverlessApplicationId)
	}
	if len(_emrserverlessClientToken) > 0 {
		input.ClientToken = aws.String(_emrserverlessClientToken)
	}
	if len(_emrserverlessArchitecture) > 0 {
		if err := assignInputField(input, "Architecture", _emrserverlessArchitecture); err != nil {
			log.Errorf("invalid --architecture: %s", err.Error())
			return
		}
	}
	if len(_emrserverlessAutoStartConfiguration) > 0 {
		if err := assignInputField(input, "AutoStartConfiguration", _emrserverlessAutoStartConfiguration); err != nil {
			log.Errorf("invalid --auto-start-configuration: %s", err.Error())
			return
		}
	}
	if len(_emrserverlessAutoStopConfiguration) > 0 {
		if err := assignInputField(input, "AutoStopConfiguration", _emrserverlessAutoStopConfiguration); err != nil {
			log.Errorf("invalid --auto-stop-configuration: %s", err.Error())
			return
		}
	}
	if len(_emrserverlessDiskEncryptionConfiguration) > 0 {
		if err := assignInputField(input, "DiskEncryptionConfiguration", _emrserverlessDiskEncryptionConfiguration); err != nil {
			log.Errorf("invalid --disk-encryption-configuration: %s", err.Error())
			return
		}
	}
	if len(_emrserverlessIdentityCenterConfiguration) > 0 {
		if err := assignInputField(input, "IdentityCenterConfiguration", _emrserverlessIdentityCenterConfiguration); err != nil {
			log.Errorf("invalid --identity-center-configuration: %s", err.Error())
			return
		}
	}
	if len(_emrserverlessImageConfiguration) > 0 {
		if err := assignInputField(input, "ImageConfiguration", _emrserverlessImageConfiguration); err != nil {
			log.Errorf("invalid --image-configuration: %s", err.Error())
			return
		}
	}
	if len(_emrserverlessInitialCapacity) > 0 {
		if err := assignInputField(input, "InitialCapacity", _emrserverlessInitialCapacity); err != nil {
			log.Errorf("invalid --initial-capacity: %s", err.Error())
			return
		}
	}
	if len(_emrserverlessInteractiveConfiguration) > 0 {
		if err := assignInputField(input, "InteractiveConfiguration", _emrserverlessInteractiveConfiguration); err != nil {
			log.Errorf("invalid --interactive-configuration: %s", err.Error())
			return
		}
	}
	if len(_emrserverlessJobLevelCostAllocationConfiguration) > 0 {
		if err := assignInputField(input, "JobLevelCostAllocationConfiguration", _emrserverlessJobLevelCostAllocationConfiguration); err != nil {
			log.Errorf("invalid --job-level-cost-allocation-configuration: %s", err.Error())
			return
		}
	}
	if len(_emrserverlessMaximumCapacity) > 0 {
		if err := assignInputField(input, "MaximumCapacity", _emrserverlessMaximumCapacity); err != nil {
			log.Errorf("invalid --maximum-capacity: %s", err.Error())
			return
		}
	}
	if len(_emrserverlessMonitoringConfiguration) > 0 {
		if err := assignInputField(input, "MonitoringConfiguration", _emrserverlessMonitoringConfiguration); err != nil {
			log.Errorf("invalid --monitoring-configuration: %s", err.Error())
			return
		}
	}
	if len(_emrserverlessNetworkConfiguration) > 0 {
		if err := assignInputField(input, "NetworkConfiguration", _emrserverlessNetworkConfiguration); err != nil {
			log.Errorf("invalid --network-configuration: %s", err.Error())
			return
		}
	}
	if len(_emrserverlessReleaseLabel) > 0 {
		input.ReleaseLabel = aws.String(_emrserverlessReleaseLabel)
	}
	if len(_emrserverlessRuntimeConfiguration) > 0 {
		if err := assignInputField(input, "RuntimeConfiguration", _emrserverlessRuntimeConfiguration); err != nil {
			log.Errorf("invalid --runtime-configuration: %s", err.Error())
			return
		}
	}
	if len(_emrserverlessSchedulerConfiguration) > 0 {
		if err := assignInputField(input, "SchedulerConfiguration", _emrserverlessSchedulerConfiguration); err != nil {
			log.Errorf("invalid --scheduler-configuration: %s", err.Error())
			return
		}
	}
	if len(_emrserverlessWorkerTypeSpecifications) > 0 {
		if err := assignInputField(input, "WorkerTypeSpecifications", _emrserverlessWorkerTypeSpecifications); err != nil {
			log.Errorf("invalid --worker-type-specifications: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_emrserverlessCmd)
	_emrserverlessCmd.Flags().SortFlags = false

	_emrserverlessCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_emrserverlessCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_emrserverlessCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_emrserverlessCmd.Flags().StringVarP(&_emrserverlessAccessSystemProfileLogs, "access-system-profile-logs", "", "", "Access System Profile Logs")
	_emrserverlessCmd.Flags().StringVarP(&_emrserverlessApplicationId, "application-id", "", "", "Application ID")
	_emrserverlessCmd.Flags().StringVarP(&_emrserverlessArchitecture, "architecture", "", "", "Architecture")
	_emrserverlessCmd.Flags().StringVarP(&_emrserverlessAttempt, "attempt", "", "", "Attempt")
	_emrserverlessCmd.Flags().StringVarP(&_emrserverlessAutoStartConfiguration, "auto-start-configuration", "", "", "Auto Start Configuration")
	_emrserverlessCmd.Flags().StringVarP(&_emrserverlessAutoStopConfiguration, "auto-stop-configuration", "", "", "Auto Stop Configuration")
	_emrserverlessCmd.Flags().StringVarP(&_emrserverlessClientToken, "client-token", "", "", "Client Token")
	_emrserverlessCmd.Flags().StringVarP(&_emrserverlessConfigurationOverrides, "configuration-overrides", "", "", "Configuration Overrides")
	_emrserverlessCmd.Flags().StringVarP(&_emrserverlessCreatedAtAfter, "created-at-after", "", "", "Created At After")
	_emrserverlessCmd.Flags().StringVarP(&_emrserverlessCreatedAtBefore, "created-at-before", "", "", "Created At Before")
	_emrserverlessCmd.Flags().StringVarP(&_emrserverlessDiskEncryptionConfiguration, "disk-encryption-configuration", "", "", "Disk Encryption Configuration")
	_emrserverlessCmd.Flags().StringVarP(&_emrserverlessExecutionIamPolicy, "execution-iam-policy", "", "", "Execution IAM Policy")
	_emrserverlessCmd.Flags().StringVarP(&_emrserverlessExecutionRoleArn, "execution-role-arn", "", "", "Execution Role ARN")
	_emrserverlessCmd.Flags().StringVarP(&_emrserverlessExecutionTimeoutMinutes, "execution-timeout-minutes", "", "", "Execution Timeout Minutes")
	_emrserverlessCmd.Flags().StringVarP(&_emrserverlessIdentityCenterConfiguration, "identity-center-configuration", "", "", "Identity Center Configuration")
	_emrserverlessCmd.Flags().StringVarP(&_emrserverlessImageConfiguration, "image-configuration", "", "", "Image Configuration")
	_emrserverlessCmd.Flags().StringVarP(&_emrserverlessInitialCapacity, "initial-capacity", "", "", "Initial Capacity")
	_emrserverlessCmd.Flags().StringVarP(&_emrserverlessInteractiveConfiguration, "interactive-configuration", "", "", "Interactive Configuration")
	_emrserverlessCmd.Flags().StringVarP(&_emrserverlessJobDriver, "job-driver", "", "", "Job Driver")
	_emrserverlessCmd.Flags().StringVarP(&_emrserverlessJobLevelCostAllocationConfiguration, "job-level-cost-allocation-configuration", "", "", "Job Level Cost Allocation Configuration")
	_emrserverlessCmd.Flags().StringVarP(&_emrserverlessJobRunId, "job-run-id", "", "", "Job Run ID")
	_emrserverlessCmd.Flags().StringVarP(&_emrserverlessMaxResults, "max-results", "", "", "Max Results")
	_emrserverlessCmd.Flags().StringVarP(&_emrserverlessMaximumCapacity, "maximum-capacity", "", "", "Maximum Capacity")
	_emrserverlessCmd.Flags().StringVarP(&_emrserverlessMode, "mode", "", "", "Mode")
	_emrserverlessCmd.Flags().StringVarP(&_emrserverlessMonitoringConfiguration, "monitoring-configuration", "", "", "Monitoring Configuration")
	_emrserverlessCmd.Flags().StringVarP(&_emrserverlessName, "name", "", "", "Name")
	_emrserverlessCmd.Flags().StringVarP(&_emrserverlessNetworkConfiguration, "network-configuration", "", "", "Network Configuration")
	_emrserverlessCmd.Flags().StringVarP(&_emrserverlessNextToken, "next-token", "", "", "Next Token")
	_emrserverlessCmd.Flags().StringVarP(&_emrserverlessReleaseLabel, "release-label", "", "", "Release Label")
	_emrserverlessCmd.Flags().StringVarP(&_emrserverlessResourceArn, "resource-arn", "", "", "Resource ARN")
	_emrserverlessCmd.Flags().StringVarP(&_emrserverlessRetryPolicy, "retry-policy", "", "", "Retry Policy")
	_emrserverlessCmd.Flags().StringVarP(&_emrserverlessRuntimeConfiguration, "runtime-configuration", "", "", "Runtime Configuration")
	_emrserverlessCmd.Flags().StringVarP(&_emrserverlessSchedulerConfiguration, "scheduler-configuration", "", "", "Scheduler Configuration")
	_emrserverlessCmd.Flags().StringVarP(&_emrserverlessShutdownGracePeriodInSeconds, "shutdown-grace-period-in-seconds", "", "", "Shutdown Grace Period In Seconds")
	_emrserverlessCmd.Flags().StringVarP(&_emrserverlessStates, "states", "", "", "States")
	_emrserverlessCmd.Flags().StringSliceVarP(&_emrserverlessTagKeys, "tag-keys", "", nil, "Tag Keys")
	_emrserverlessCmd.Flags().StringVarP(&_emrserverlessTags, "tags", "", "", "Tags")
	_emrserverlessCmd.Flags().StringVarP(&_emrserverlessType, "type", "", "", "Type")
	_emrserverlessCmd.Flags().StringVarP(&_emrserverlessWorkerTypeSpecifications, "worker-type-specifications", "", "", "Worker Type Specifications")

	_emrserverlessCmd.Flags().BoolVarP(&_emrserverlessCancelJobRun, "cancel-job-run", "", false, "Cancel Job Run")
	_emrserverlessCmd.Flags().BoolVarP(&_emrserverlessCreateApplication, "create-application", "", false, "Create Application")
	_emrserverlessCmd.Flags().BoolVarP(&_emrserverlessDeleteApplication, "delete-application", "", false, "Delete Application")
	_emrserverlessCmd.Flags().BoolVarP(&_emrserverlessGetApplication, "get-application", "", false, "Get Application")
	_emrserverlessCmd.Flags().BoolVarP(&_emrserverlessGetDashboardForJobRun, "get-dashboard-for-job-run", "", false, "Get Dashboard For Job Run")
	_emrserverlessCmd.Flags().BoolVarP(&_emrserverlessGetJobRun, "get-job-run", "", false, "Get Job Run")
	_emrserverlessCmd.Flags().BoolVarP(&_emrserverlessListApplications, "list-applications", "", false, "List Applications")
	_emrserverlessCmd.Flags().BoolVarP(&_emrserverlessListJobRunAttempts, "list-job-run-attempts", "", false, "List Job Run Attempts")
	_emrserverlessCmd.Flags().BoolVarP(&_emrserverlessListJobRuns, "list-job-runs", "", false, "List Job Runs")
	_emrserverlessCmd.Flags().BoolVarP(&_emrserverlessListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_emrserverlessCmd.Flags().BoolVarP(&_emrserverlessStartApplication, "start-application", "", false, "Start Application")
	_emrserverlessCmd.Flags().BoolVarP(&_emrserverlessStartJobRun, "start-job-run", "", false, "Start Job Run")
	_emrserverlessCmd.Flags().BoolVarP(&_emrserverlessStopApplication, "stop-application", "", false, "Stop Application")
	_emrserverlessCmd.Flags().BoolVarP(&_emrserverlessTagResource, "tag-resource", "", false, "Tag Resource")
	_emrserverlessCmd.Flags().BoolVarP(&_emrserverlessUntagResource, "untag-resource", "", false, "Untag Resource")
	_emrserverlessCmd.Flags().BoolVarP(&_emrserverlessUpdateApplication, "update-application", "", false, "Update Application")

}
