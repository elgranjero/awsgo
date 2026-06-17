package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mwaaserverless"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// mwaaserverlessCmd represents the mwaaserverless command
var _mwaaserverlessCmd = &cobra.Command{
	Use:   "mwaaserverless",
	Short: "AWS mwaaserverless CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := mwaaserverless.NewFromConfig(cfg)
		if _mwaaserverlessCreateWorkflow {
			mwaaserverless_CreateWorkflow(cfg, client)
			return
		}
		if _mwaaserverlessDeleteWorkflow {
			mwaaserverless_DeleteWorkflow(cfg, client)
			return
		}
		if _mwaaserverlessGetTaskInstance {
			mwaaserverless_GetTaskInstance(cfg, client)
			return
		}
		if _mwaaserverlessGetWorkflow {
			mwaaserverless_GetWorkflow(cfg, client)
			return
		}
		if _mwaaserverlessGetWorkflowRun {
			mwaaserverless_GetWorkflowRun(cfg, client)
			return
		}
		if _mwaaserverlessListTagsForResource {
			mwaaserverless_ListTagsForResource(cfg, client)
			return
		}
		if _mwaaserverlessListTaskInstances {
			mwaaserverless_ListTaskInstances(cfg, client)
			return
		}
		if _mwaaserverlessListWorkflowRuns {
			mwaaserverless_ListWorkflowRuns(cfg, client)
			return
		}
		if _mwaaserverlessListWorkflowVersions {
			mwaaserverless_ListWorkflowVersions(cfg, client)
			return
		}
		if _mwaaserverlessListWorkflows {
			mwaaserverless_ListWorkflows(cfg, client)
			return
		}
		if _mwaaserverlessStartWorkflowRun {
			mwaaserverless_StartWorkflowRun(cfg, client)
			return
		}
		if _mwaaserverlessStopWorkflowRun {
			mwaaserverless_StopWorkflowRun(cfg, client)
			return
		}
		if _mwaaserverlessTagResource {
			mwaaserverless_TagResource(cfg, client)
			return
		}
		if _mwaaserverlessUntagResource {
			mwaaserverless_UntagResource(cfg, client)
			return
		}
		if _mwaaserverlessUpdateWorkflow {
			mwaaserverless_UpdateWorkflow(cfg, client)
			return
		}

	},
}

var (
	_mwaaserverlessCreateWorkflow       bool
	_mwaaserverlessDeleteWorkflow       bool
	_mwaaserverlessGetTaskInstance      bool
	_mwaaserverlessGetWorkflow          bool
	_mwaaserverlessGetWorkflowRun       bool
	_mwaaserverlessListTagsForResource  bool
	_mwaaserverlessListTaskInstances    bool
	_mwaaserverlessListWorkflowRuns     bool
	_mwaaserverlessListWorkflowVersions bool
	_mwaaserverlessListWorkflows        bool
	_mwaaserverlessStartWorkflowRun     bool
	_mwaaserverlessStopWorkflowRun      bool
	_mwaaserverlessTagResource          bool
	_mwaaserverlessUntagResource        bool
	_mwaaserverlessUpdateWorkflow       bool

	_mwaaserverlessClientToken             string
	_mwaaserverlessDefinitionS3Location    string
	_mwaaserverlessDescription             string
	_mwaaserverlessEncryptionConfiguration string
	_mwaaserverlessEngineVersion           string
	_mwaaserverlessLoggingConfiguration    string
	_mwaaserverlessMaxResults              string
	_mwaaserverlessName                    string
	_mwaaserverlessNetworkConfiguration    string
	_mwaaserverlessNextToken               string
	_mwaaserverlessOverrideParameters      string
	_mwaaserverlessResourceArn             string
	_mwaaserverlessRoleArn                 string
	_mwaaserverlessRunId                   string
	_mwaaserverlessTagKeys                 []string
	_mwaaserverlessTags                    string
	_mwaaserverlessTaskInstanceId          string
	_mwaaserverlessTriggerMode             string
	_mwaaserverlessWorkflowArn             string
	_mwaaserverlessWorkflowVersion         string
)

// Creates a new workflow in Amazon Managed Workflows for Apache Airflow
// Serverless. This operation initializes a workflow with the specified
// configuration including the workflow definition, execution role, and optional
// settings for encryption, logging, and networking. You must provide the workflow
// definition as a YAML file stored in Amazon S3 that defines the DAG structure
// using supported Amazon Web Services operators. Amazon Managed Workflows for
// Apache Airflow Serverless automatically creates the first version of the
// workflow and sets up the necessary execution environment with multi-tenant
// isolation and security controls.
func mwaaserverless_CreateWorkflow(cfg aws.Config, client *mwaaserverless.Client) {
	input := &mwaaserverless.CreateWorkflowInput{
		// DefinitionS3Location: *types.DefinitionS3Location, // Required
		// Name: *string, // Required
		// RoleArn: *string, // Required
	}

	if len(_mwaaserverlessDefinitionS3Location) > 0 {
		if err := assignInputField(input, "DefinitionS3Location", _mwaaserverlessDefinitionS3Location); err != nil {
			log.Errorf("invalid --definition-s3-location: %s", err.Error())
			return
		}
	}
	if len(_mwaaserverlessName) > 0 {
		input.Name = aws.String(_mwaaserverlessName)
	}
	if len(_mwaaserverlessRoleArn) > 0 {
		input.RoleArn = aws.String(_mwaaserverlessRoleArn)
	}
	if len(_mwaaserverlessClientToken) > 0 {
		input.ClientToken = aws.String(_mwaaserverlessClientToken)
	}
	if len(_mwaaserverlessDescription) > 0 {
		input.Description = aws.String(_mwaaserverlessDescription)
	}
	if len(_mwaaserverlessEncryptionConfiguration) > 0 {
		if err := assignInputField(input, "EncryptionConfiguration", _mwaaserverlessEncryptionConfiguration); err != nil {
			log.Errorf("invalid --encryption-configuration: %s", err.Error())
			return
		}
	}
	if len(_mwaaserverlessEngineVersion) > 0 {
		if err := assignInputField(input, "EngineVersion", _mwaaserverlessEngineVersion); err != nil {
			log.Errorf("invalid --engine-version: %s", err.Error())
			return
		}
	}
	if len(_mwaaserverlessLoggingConfiguration) > 0 {
		if err := assignInputField(input, "LoggingConfiguration", _mwaaserverlessLoggingConfiguration); err != nil {
			log.Errorf("invalid --logging-configuration: %s", err.Error())
			return
		}
	}
	if len(_mwaaserverlessNetworkConfiguration) > 0 {
		if err := assignInputField(input, "NetworkConfiguration", _mwaaserverlessNetworkConfiguration); err != nil {
			log.Errorf("invalid --network-configuration: %s", err.Error())
			return
		}
	}
	if len(_mwaaserverlessTags) > 0 {
		if err := assignInputField(input, "Tags", _mwaaserverlessTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_mwaaserverlessTriggerMode) > 0 {
		input.TriggerMode = aws.String(_mwaaserverlessTriggerMode)
	}

	if resp, err := client.CreateWorkflow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a workflow and all its versions. This operation permanently removes the
// workflow and cannot be undone. Amazon Managed Workflows for Apache Airflow
// Serverless ensures that all associated resources are properly cleaned up,
// including stopping any running executions, removing scheduled triggers, and
// cleaning up execution history. The deletion process respects the multi-tenant
// isolation boundaries and ensures that no residual data or configurations remain
// that could affect other customers or workflows.
func mwaaserverless_DeleteWorkflow(cfg aws.Config, client *mwaaserverless.Client) {
	input := &mwaaserverless.DeleteWorkflowInput{
		// WorkflowArn: *string, // Required
	}

	if len(_mwaaserverlessWorkflowArn) > 0 {
		input.WorkflowArn = aws.String(_mwaaserverlessWorkflowArn)
	}
	if len(_mwaaserverlessWorkflowVersion) > 0 {
		input.WorkflowVersion = aws.String(_mwaaserverlessWorkflowVersion)
	}

	if resp, err := client.DeleteWorkflow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves detailed information about a specific task instance within a workflow
// run. Task instances represent individual tasks that are executed as part of a
// workflow in the Amazon Managed Workflows for Apache Airflow Serverless
// environment. Each task instance runs in an isolated ECS container with dedicated
// resources and security boundaries. The service tracks task execution state,
// retry attempts, and provides detailed timing and error information for
// troubleshooting and monitoring purposes.
func mwaaserverless_GetTaskInstance(cfg aws.Config, client *mwaaserverless.Client) {
	input := &mwaaserverless.GetTaskInstanceInput{
		// RunId: *string, // Required
		// TaskInstanceId: *string, // Required
		// WorkflowArn: *string, // Required
	}

	if len(_mwaaserverlessRunId) > 0 {
		input.RunId = aws.String(_mwaaserverlessRunId)
	}
	if len(_mwaaserverlessTaskInstanceId) > 0 {
		input.TaskInstanceId = aws.String(_mwaaserverlessTaskInstanceId)
	}
	if len(_mwaaserverlessWorkflowArn) > 0 {
		input.WorkflowArn = aws.String(_mwaaserverlessWorkflowArn)
	}

	if resp, err := client.GetTaskInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves detailed information about a workflow, including its configuration,
// status, and metadata.
func mwaaserverless_GetWorkflow(cfg aws.Config, client *mwaaserverless.Client) {
	input := &mwaaserverless.GetWorkflowInput{
		// WorkflowArn: *string, // Required
	}

	if len(_mwaaserverlessWorkflowArn) > 0 {
		input.WorkflowArn = aws.String(_mwaaserverlessWorkflowArn)
	}
	if len(_mwaaserverlessWorkflowVersion) > 0 {
		input.WorkflowVersion = aws.String(_mwaaserverlessWorkflowVersion)
	}

	if resp, err := client.GetWorkflow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves detailed information about a specific workflow run, including its
// status, execution details, and task instances.
func mwaaserverless_GetWorkflowRun(cfg aws.Config, client *mwaaserverless.Client) {
	input := &mwaaserverless.GetWorkflowRunInput{
		// RunId: *string, // Required
		// WorkflowArn: *string, // Required
	}

	if len(_mwaaserverlessRunId) > 0 {
		input.RunId = aws.String(_mwaaserverlessRunId)
	}
	if len(_mwaaserverlessWorkflowArn) > 0 {
		input.WorkflowArn = aws.String(_mwaaserverlessWorkflowArn)
	}

	if resp, err := client.GetWorkflowRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all tags that are associated with a specified Amazon Managed Workflows
// for Apache Airflow Serverless resource.
func mwaaserverless_ListTagsForResource(cfg aws.Config, client *mwaaserverless.Client) {
	input := &mwaaserverless.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_mwaaserverlessResourceArn) > 0 {
		input.ResourceArn = aws.String(_mwaaserverlessResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all task instances for a specific workflow run, with optional pagination
// support.
func mwaaserverless_ListTaskInstances(cfg aws.Config, client *mwaaserverless.Client) {
	input := &mwaaserverless.ListTaskInstancesInput{
		// RunId: *string, // Required
		// WorkflowArn: *string, // Required
	}

	if len(_mwaaserverlessRunId) > 0 {
		input.RunId = aws.String(_mwaaserverlessRunId)
	}
	if len(_mwaaserverlessWorkflowArn) > 0 {
		input.WorkflowArn = aws.String(_mwaaserverlessWorkflowArn)
	}
	if len(_mwaaserverlessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mwaaserverlessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mwaaserverlessNextToken) > 0 {
		input.NextToken = aws.String(_mwaaserverlessNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTaskInstances(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mwaaserverless.ListTaskInstancesOutput
	p := mwaaserverless.NewListTaskInstancesPaginator(client, input)
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

// Lists all runs for a specified workflow, with optional pagination and filtering
// support.
func mwaaserverless_ListWorkflowRuns(cfg aws.Config, client *mwaaserverless.Client) {
	input := &mwaaserverless.ListWorkflowRunsInput{
		// WorkflowArn: *string, // Required
	}

	if len(_mwaaserverlessWorkflowArn) > 0 {
		input.WorkflowArn = aws.String(_mwaaserverlessWorkflowArn)
	}
	if len(_mwaaserverlessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mwaaserverlessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mwaaserverlessNextToken) > 0 {
		input.NextToken = aws.String(_mwaaserverlessNextToken)
	}
	if len(_mwaaserverlessWorkflowVersion) > 0 {
		input.WorkflowVersion = aws.String(_mwaaserverlessWorkflowVersion)
	}

	if disablePaginator() {
		if resp, err := client.ListWorkflowRuns(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mwaaserverless.ListWorkflowRunsOutput
	p := mwaaserverless.NewListWorkflowRunsPaginator(client, input)
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

// Lists all versions of a specified workflow, with optional pagination support.
func mwaaserverless_ListWorkflowVersions(cfg aws.Config, client *mwaaserverless.Client) {
	input := &mwaaserverless.ListWorkflowVersionsInput{
		// WorkflowArn: *string, // Required
	}

	if len(_mwaaserverlessWorkflowArn) > 0 {
		input.WorkflowArn = aws.String(_mwaaserverlessWorkflowArn)
	}
	if len(_mwaaserverlessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mwaaserverlessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mwaaserverlessNextToken) > 0 {
		input.NextToken = aws.String(_mwaaserverlessNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListWorkflowVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mwaaserverless.ListWorkflowVersionsOutput
	p := mwaaserverless.NewListWorkflowVersionsPaginator(client, input)
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

// Lists all workflows in your account, with optional pagination support. This
// operation returns summary information for workflows, showing only the most
// recently created version of each workflow. Amazon Managed Workflows for Apache
// Airflow Serverless maintains workflow metadata in a highly available,
// distributed storage system that enables efficient querying and filtering. The
// service implements proper access controls to ensure you can only view workflows
// that you have permissions to access, supporting both individual and team-based
// workflow management scenarios.
func mwaaserverless_ListWorkflows(cfg aws.Config, client *mwaaserverless.Client) {
	input := &mwaaserverless.ListWorkflowsInput{}

	if len(_mwaaserverlessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mwaaserverlessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mwaaserverlessNextToken) > 0 {
		input.NextToken = aws.String(_mwaaserverlessNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListWorkflows(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mwaaserverless.ListWorkflowsOutput
	p := mwaaserverless.NewListWorkflowsPaginator(client, input)
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

// Starts a new execution of a workflow. This operation creates a workflow run
// that executes the tasks that are defined in the workflow. Amazon Managed
// Workflows for Apache Airflow Serverless schedules the workflow execution across
// its managed Airflow environment, automatically scaling ECS worker tasks based on
// the workload. The service handles task isolation, dependency resolution, and
// provides comprehensive monitoring and logging throughout the execution
// lifecycle.
func mwaaserverless_StartWorkflowRun(cfg aws.Config, client *mwaaserverless.Client) {
	input := &mwaaserverless.StartWorkflowRunInput{
		// WorkflowArn: *string, // Required
	}

	if len(_mwaaserverlessWorkflowArn) > 0 {
		input.WorkflowArn = aws.String(_mwaaserverlessWorkflowArn)
	}
	if len(_mwaaserverlessClientToken) > 0 {
		input.ClientToken = aws.String(_mwaaserverlessClientToken)
	}
	if len(_mwaaserverlessOverrideParameters) > 0 {
		if err := assignInputField(input, "OverrideParameters", _mwaaserverlessOverrideParameters); err != nil {
			log.Errorf("invalid --override-parameters: %s", err.Error())
			return
		}
	}
	if len(_mwaaserverlessWorkflowVersion) > 0 {
		input.WorkflowVersion = aws.String(_mwaaserverlessWorkflowVersion)
	}

	if resp, err := client.StartWorkflowRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a running workflow execution. This operation terminates all running tasks
// and prevents new tasks from starting. Amazon Managed Workflows for Apache
// Airflow Serverless gracefully shuts down the workflow execution by stopping task
// scheduling and terminating active ECS worker containers. The operation
// transitions the workflow run to a STOPPING state and then to STOPPED once all
// cleanup is complete. In-flight tasks may complete or be terminated depending on
// their current execution state.
func mwaaserverless_StopWorkflowRun(cfg aws.Config, client *mwaaserverless.Client) {
	input := &mwaaserverless.StopWorkflowRunInput{
		// RunId: *string, // Required
		// WorkflowArn: *string, // Required
	}

	if len(_mwaaserverlessRunId) > 0 {
		input.RunId = aws.String(_mwaaserverlessRunId)
	}
	if len(_mwaaserverlessWorkflowArn) > 0 {
		input.WorkflowArn = aws.String(_mwaaserverlessWorkflowArn)
	}

	if resp, err := client.StopWorkflowRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds tags to an Amazon Managed Workflows for Apache Airflow Serverless
// resource. Tags are key-value pairs that help you organize and categorize your
// resources.
func mwaaserverless_TagResource(cfg aws.Config, client *mwaaserverless.Client) {
	input := &mwaaserverless.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_mwaaserverlessResourceArn) > 0 {
		input.ResourceArn = aws.String(_mwaaserverlessResourceArn)
	}
	if len(_mwaaserverlessTags) > 0 {
		if err := assignInputField(input, "Tags", _mwaaserverlessTags); err != nil {
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

// Removes tags from an Amazon Managed Workflows for Apache Airflow Serverless
// resource. This operation removes the specified tags from the resource.
func mwaaserverless_UntagResource(cfg aws.Config, client *mwaaserverless.Client) {
	input := &mwaaserverless.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_mwaaserverlessResourceArn) > 0 {
		input.ResourceArn = aws.String(_mwaaserverlessResourceArn)
	}
	if len(_mwaaserverlessTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _mwaaserverlessTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing workflow with new configuration settings. This operation
// allows you to modify the workflow definition, role, and other settings. When you
// update a workflow, Amazon Managed Workflows for Apache Airflow Serverless
// automatically creates a new version with the updated configuration and disables
// scheduling on all previous versions to ensure only one version is actively
// scheduled at a time. The update operation maintains workflow history while
// providing a clean transition to the new configuration.
func mwaaserverless_UpdateWorkflow(cfg aws.Config, client *mwaaserverless.Client) {
	input := &mwaaserverless.UpdateWorkflowInput{
		// DefinitionS3Location: *types.DefinitionS3Location, // Required
		// RoleArn: *string, // Required
		// WorkflowArn: *string, // Required
	}

	if len(_mwaaserverlessDefinitionS3Location) > 0 {
		if err := assignInputField(input, "DefinitionS3Location", _mwaaserverlessDefinitionS3Location); err != nil {
			log.Errorf("invalid --definition-s3-location: %s", err.Error())
			return
		}
	}
	if len(_mwaaserverlessRoleArn) > 0 {
		input.RoleArn = aws.String(_mwaaserverlessRoleArn)
	}
	if len(_mwaaserverlessWorkflowArn) > 0 {
		input.WorkflowArn = aws.String(_mwaaserverlessWorkflowArn)
	}
	if len(_mwaaserverlessDescription) > 0 {
		input.Description = aws.String(_mwaaserverlessDescription)
	}
	if len(_mwaaserverlessEngineVersion) > 0 {
		if err := assignInputField(input, "EngineVersion", _mwaaserverlessEngineVersion); err != nil {
			log.Errorf("invalid --engine-version: %s", err.Error())
			return
		}
	}
	if len(_mwaaserverlessLoggingConfiguration) > 0 {
		if err := assignInputField(input, "LoggingConfiguration", _mwaaserverlessLoggingConfiguration); err != nil {
			log.Errorf("invalid --logging-configuration: %s", err.Error())
			return
		}
	}
	if len(_mwaaserverlessNetworkConfiguration) > 0 {
		if err := assignInputField(input, "NetworkConfiguration", _mwaaserverlessNetworkConfiguration); err != nil {
			log.Errorf("invalid --network-configuration: %s", err.Error())
			return
		}
	}
	if len(_mwaaserverlessTriggerMode) > 0 {
		input.TriggerMode = aws.String(_mwaaserverlessTriggerMode)
	}

	if resp, err := client.UpdateWorkflow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_mwaaserverlessCmd)
	_mwaaserverlessCmd.Flags().SortFlags = false

	_mwaaserverlessCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_mwaaserverlessCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_mwaaserverlessCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_mwaaserverlessCmd.Flags().StringVarP(&_mwaaserverlessClientToken, "client-token", "", "", "Client Token")
	_mwaaserverlessCmd.Flags().StringVarP(&_mwaaserverlessDefinitionS3Location, "definition-s3-location", "", "", "Definition S3 Location")
	_mwaaserverlessCmd.Flags().StringVarP(&_mwaaserverlessDescription, "description", "", "", "Description")
	_mwaaserverlessCmd.Flags().StringVarP(&_mwaaserverlessEncryptionConfiguration, "encryption-configuration", "", "", "Encryption Configuration")
	_mwaaserverlessCmd.Flags().StringVarP(&_mwaaserverlessEngineVersion, "engine-version", "", "", "Engine Version")
	_mwaaserverlessCmd.Flags().StringVarP(&_mwaaserverlessLoggingConfiguration, "logging-configuration", "", "", "Logging Configuration")
	_mwaaserverlessCmd.Flags().StringVarP(&_mwaaserverlessMaxResults, "max-results", "", "", "Max Results")
	_mwaaserverlessCmd.Flags().StringVarP(&_mwaaserverlessName, "name", "", "", "Name")
	_mwaaserverlessCmd.Flags().StringVarP(&_mwaaserverlessNetworkConfiguration, "network-configuration", "", "", "Network Configuration")
	_mwaaserverlessCmd.Flags().StringVarP(&_mwaaserverlessNextToken, "next-token", "", "", "Next Token")
	_mwaaserverlessCmd.Flags().StringVarP(&_mwaaserverlessOverrideParameters, "override-parameters", "", "", "Override Parameters")
	_mwaaserverlessCmd.Flags().StringVarP(&_mwaaserverlessResourceArn, "resource-arn", "", "", "Resource ARN")
	_mwaaserverlessCmd.Flags().StringVarP(&_mwaaserverlessRoleArn, "role-arn", "", "", "Role ARN")
	_mwaaserverlessCmd.Flags().StringVarP(&_mwaaserverlessRunId, "run-id", "", "", "Run ID")
	_mwaaserverlessCmd.Flags().StringSliceVarP(&_mwaaserverlessTagKeys, "tag-keys", "", nil, "Tag Keys")
	_mwaaserverlessCmd.Flags().StringVarP(&_mwaaserverlessTags, "tags", "", "", "Tags")
	_mwaaserverlessCmd.Flags().StringVarP(&_mwaaserverlessTaskInstanceId, "task-instance-id", "", "", "Task Instance ID")
	_mwaaserverlessCmd.Flags().StringVarP(&_mwaaserverlessTriggerMode, "trigger-mode", "", "", "Trigger Mode")
	_mwaaserverlessCmd.Flags().StringVarP(&_mwaaserverlessWorkflowArn, "workflow-arn", "", "", "Workflow ARN")
	_mwaaserverlessCmd.Flags().StringVarP(&_mwaaserverlessWorkflowVersion, "workflow-version", "", "", "Workflow Version")

	_mwaaserverlessCmd.Flags().BoolVarP(&_mwaaserverlessCreateWorkflow, "create-workflow", "", false, "Create Workflow")
	_mwaaserverlessCmd.Flags().BoolVarP(&_mwaaserverlessDeleteWorkflow, "delete-workflow", "", false, "Delete Workflow")
	_mwaaserverlessCmd.Flags().BoolVarP(&_mwaaserverlessGetTaskInstance, "get-task-instance", "", false, "Get Task Instance")
	_mwaaserverlessCmd.Flags().BoolVarP(&_mwaaserverlessGetWorkflow, "get-workflow", "", false, "Get Workflow")
	_mwaaserverlessCmd.Flags().BoolVarP(&_mwaaserverlessGetWorkflowRun, "get-workflow-run", "", false, "Get Workflow Run")
	_mwaaserverlessCmd.Flags().BoolVarP(&_mwaaserverlessListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_mwaaserverlessCmd.Flags().BoolVarP(&_mwaaserverlessListTaskInstances, "list-task-instances", "", false, "List Task Instances")
	_mwaaserverlessCmd.Flags().BoolVarP(&_mwaaserverlessListWorkflowRuns, "list-workflow-runs", "", false, "List Workflow Runs")
	_mwaaserverlessCmd.Flags().BoolVarP(&_mwaaserverlessListWorkflowVersions, "list-workflow-versions", "", false, "List Workflow Versions")
	_mwaaserverlessCmd.Flags().BoolVarP(&_mwaaserverlessListWorkflows, "list-workflows", "", false, "List Workflows")
	_mwaaserverlessCmd.Flags().BoolVarP(&_mwaaserverlessStartWorkflowRun, "start-workflow-run", "", false, "Start Workflow Run")
	_mwaaserverlessCmd.Flags().BoolVarP(&_mwaaserverlessStopWorkflowRun, "stop-workflow-run", "", false, "Stop Workflow Run")
	_mwaaserverlessCmd.Flags().BoolVarP(&_mwaaserverlessTagResource, "tag-resource", "", false, "Tag Resource")
	_mwaaserverlessCmd.Flags().BoolVarP(&_mwaaserverlessUntagResource, "untag-resource", "", false, "Untag Resource")
	_mwaaserverlessCmd.Flags().BoolVarP(&_mwaaserverlessUpdateWorkflow, "update-workflow", "", false, "Update Workflow")

}
