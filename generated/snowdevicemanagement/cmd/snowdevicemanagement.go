package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/snowdevicemanagement"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// snowdevicemanagementCmd represents the snowdevicemanagement command
var _snowdevicemanagementCmd = &cobra.Command{
	Use:   "snowdevicemanagement",
	Short: "AWS snowdevicemanagement CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := snowdevicemanagement.NewFromConfig(cfg)
		if _snowdevicemanagementCancelTask {
			snowdevicemanagement_CancelTask(cfg, client)
			return
		}
		if _snowdevicemanagementCreateTask {
			snowdevicemanagement_CreateTask(cfg, client)
			return
		}
		if _snowdevicemanagementDescribeDevice {
			snowdevicemanagement_DescribeDevice(cfg, client)
			return
		}
		if _snowdevicemanagementDescribeDeviceEc2Instances {
			snowdevicemanagement_DescribeDeviceEc2Instances(cfg, client)
			return
		}
		if _snowdevicemanagementDescribeExecution {
			snowdevicemanagement_DescribeExecution(cfg, client)
			return
		}
		if _snowdevicemanagementDescribeTask {
			snowdevicemanagement_DescribeTask(cfg, client)
			return
		}
		if _snowdevicemanagementListDeviceResources {
			snowdevicemanagement_ListDeviceResources(cfg, client)
			return
		}
		if _snowdevicemanagementListDevices {
			snowdevicemanagement_ListDevices(cfg, client)
			return
		}
		if _snowdevicemanagementListExecutions {
			snowdevicemanagement_ListExecutions(cfg, client)
			return
		}
		if _snowdevicemanagementListTagsForResource {
			snowdevicemanagement_ListTagsForResource(cfg, client)
			return
		}
		if _snowdevicemanagementListTasks {
			snowdevicemanagement_ListTasks(cfg, client)
			return
		}
		if _snowdevicemanagementTagResource {
			snowdevicemanagement_TagResource(cfg, client)
			return
		}
		if _snowdevicemanagementUntagResource {
			snowdevicemanagement_UntagResource(cfg, client)
			return
		}

	},
}

var (
	_snowdevicemanagementCancelTask                 bool
	_snowdevicemanagementCreateTask                 bool
	_snowdevicemanagementDescribeDevice             bool
	_snowdevicemanagementDescribeDeviceEc2Instances bool
	_snowdevicemanagementDescribeExecution          bool
	_snowdevicemanagementDescribeTask               bool
	_snowdevicemanagementListDeviceResources        bool
	_snowdevicemanagementListDevices                bool
	_snowdevicemanagementListExecutions             bool
	_snowdevicemanagementListTagsForResource        bool
	_snowdevicemanagementListTasks                  bool
	_snowdevicemanagementTagResource                bool
	_snowdevicemanagementUntagResource              bool

	_snowdevicemanagementClientToken     string
	_snowdevicemanagementCommand         string
	_snowdevicemanagementDescription     string
	_snowdevicemanagementInstanceIds     []string
	_snowdevicemanagementJobId           string
	_snowdevicemanagementManagedDeviceId string
	_snowdevicemanagementMaxResults      string
	_snowdevicemanagementNextToken       string
	_snowdevicemanagementResourceArn     string
	_snowdevicemanagementState           string
	_snowdevicemanagementTagKeys         []string
	_snowdevicemanagementTags            string
	_snowdevicemanagementTargets         []string
	_snowdevicemanagementTaskId          string
	_snowdevicemanagementType            string
)

// Sends a cancel request for a specified task. You can cancel a task only if it's
// still in a QUEUED state. Tasks that are already running can't be cancelled.
//
// A task might still run if it's processed from the queue before the CancelTask
// operation changes the task's state.
func snowdevicemanagement_CancelTask(cfg aws.Config, client *snowdevicemanagement.Client) {
	input := &snowdevicemanagement.CancelTaskInput{
		// TaskId: *string, // Required
	}

	if len(_snowdevicemanagementTaskId) > 0 {
		input.TaskId = aws.String(_snowdevicemanagementTaskId)
	}

	if resp, err := client.CancelTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Instructs one or more devices to start a task, such as unlocking or rebooting.
func snowdevicemanagement_CreateTask(cfg aws.Config, client *snowdevicemanagement.Client) {
	input := &snowdevicemanagement.CreateTaskInput{
		// Command: types.Command, // Required
		// Targets: []string, // Required
	}

	if len(_snowdevicemanagementCommand) > 0 {
		if err := assignInputField(input, "Command", _snowdevicemanagementCommand); err != nil {
			log.Errorf("invalid --command: %s", err.Error())
			return
		}
	}
	if len(_snowdevicemanagementTargets) > 0 {
		input.Targets = append([]string(nil), _snowdevicemanagementTargets...)
	}
	if len(_snowdevicemanagementClientToken) > 0 {
		input.ClientToken = aws.String(_snowdevicemanagementClientToken)
	}
	if len(_snowdevicemanagementDescription) > 0 {
		input.Description = aws.String(_snowdevicemanagementDescription)
	}
	if len(_snowdevicemanagementTags) > 0 {
		if err := assignInputField(input, "Tags", _snowdevicemanagementTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Checks device-specific information, such as the device type, software version,
// IP addresses, and lock status.
func snowdevicemanagement_DescribeDevice(cfg aws.Config, client *snowdevicemanagement.Client) {
	input := &snowdevicemanagement.DescribeDeviceInput{
		// ManagedDeviceId: *string, // Required
	}

	if len(_snowdevicemanagementManagedDeviceId) > 0 {
		input.ManagedDeviceId = aws.String(_snowdevicemanagementManagedDeviceId)
	}

	if resp, err := client.DescribeDevice(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Checks the current state of the Amazon EC2 instances. The output is similar to
// describeDevice , but the results are sourced from the device cache in the Amazon
// Web Services Cloud and include a subset of the available fields.
func snowdevicemanagement_DescribeDeviceEc2Instances(cfg aws.Config, client *snowdevicemanagement.Client) {
	input := &snowdevicemanagement.DescribeDeviceEc2InstancesInput{
		// InstanceIds: []string, // Required
		// ManagedDeviceId: *string, // Required
	}

	if len(_snowdevicemanagementInstanceIds) > 0 {
		input.InstanceIds = append([]string(nil), _snowdevicemanagementInstanceIds...)
	}
	if len(_snowdevicemanagementManagedDeviceId) > 0 {
		input.ManagedDeviceId = aws.String(_snowdevicemanagementManagedDeviceId)
	}

	if resp, err := client.DescribeDeviceEc2Instances(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Checks the status of a remote task running on one or more target devices.
func snowdevicemanagement_DescribeExecution(cfg aws.Config, client *snowdevicemanagement.Client) {
	input := &snowdevicemanagement.DescribeExecutionInput{
		// ManagedDeviceId: *string, // Required
		// TaskId: *string, // Required
	}

	if len(_snowdevicemanagementManagedDeviceId) > 0 {
		input.ManagedDeviceId = aws.String(_snowdevicemanagementManagedDeviceId)
	}
	if len(_snowdevicemanagementTaskId) > 0 {
		input.TaskId = aws.String(_snowdevicemanagementTaskId)
	}

	if resp, err := client.DescribeExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Checks the metadata for a given task on a device.
func snowdevicemanagement_DescribeTask(cfg aws.Config, client *snowdevicemanagement.Client) {
	input := &snowdevicemanagement.DescribeTaskInput{
		// TaskId: *string, // Required
	}

	if len(_snowdevicemanagementTaskId) > 0 {
		input.TaskId = aws.String(_snowdevicemanagementTaskId)
	}

	if resp, err := client.DescribeTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of the Amazon Web Services resources available for a device.
// Currently, Amazon EC2 instances are the only supported resource type.
func snowdevicemanagement_ListDeviceResources(cfg aws.Config, client *snowdevicemanagement.Client) {
	input := &snowdevicemanagement.ListDeviceResourcesInput{
		// ManagedDeviceId: *string, // Required
	}

	if len(_snowdevicemanagementManagedDeviceId) > 0 {
		input.ManagedDeviceId = aws.String(_snowdevicemanagementManagedDeviceId)
	}
	if len(_snowdevicemanagementMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _snowdevicemanagementMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_snowdevicemanagementNextToken) > 0 {
		input.NextToken = aws.String(_snowdevicemanagementNextToken)
	}
	if len(_snowdevicemanagementType) > 0 {
		input.Type = aws.String(_snowdevicemanagementType)
	}

	if disablePaginator() {
		if resp, err := client.ListDeviceResources(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*snowdevicemanagement.ListDeviceResourcesOutput
	p := snowdevicemanagement.NewListDeviceResourcesPaginator(client, input)
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

// Returns a list of all devices on your Amazon Web Services account that have
// Amazon Web Services Snow Device Management enabled in the Amazon Web Services
// Region where the command is run.
func snowdevicemanagement_ListDevices(cfg aws.Config, client *snowdevicemanagement.Client) {
	input := &snowdevicemanagement.ListDevicesInput{}

	if len(_snowdevicemanagementJobId) > 0 {
		input.JobId = aws.String(_snowdevicemanagementJobId)
	}
	if len(_snowdevicemanagementMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _snowdevicemanagementMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_snowdevicemanagementNextToken) > 0 {
		input.NextToken = aws.String(_snowdevicemanagementNextToken)
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

	var results []*snowdevicemanagement.ListDevicesOutput
	p := snowdevicemanagement.NewListDevicesPaginator(client, input)
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

// Returns the status of tasks for one or more target devices.
func snowdevicemanagement_ListExecutions(cfg aws.Config, client *snowdevicemanagement.Client) {
	input := &snowdevicemanagement.ListExecutionsInput{
		// TaskId: *string, // Required
	}

	if len(_snowdevicemanagementTaskId) > 0 {
		input.TaskId = aws.String(_snowdevicemanagementTaskId)
	}
	if len(_snowdevicemanagementMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _snowdevicemanagementMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_snowdevicemanagementNextToken) > 0 {
		input.NextToken = aws.String(_snowdevicemanagementNextToken)
	}
	if len(_snowdevicemanagementState) > 0 {
		if err := assignInputField(input, "State", _snowdevicemanagementState); err != nil {
			log.Errorf("invalid --state: %s", err.Error())
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

	var results []*snowdevicemanagement.ListExecutionsOutput
	p := snowdevicemanagement.NewListExecutionsPaginator(client, input)
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

// Returns a list of tags for a managed device or task.
func snowdevicemanagement_ListTagsForResource(cfg aws.Config, client *snowdevicemanagement.Client) {
	input := &snowdevicemanagement.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_snowdevicemanagementResourceArn) > 0 {
		input.ResourceArn = aws.String(_snowdevicemanagementResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of tasks that can be filtered by state.
func snowdevicemanagement_ListTasks(cfg aws.Config, client *snowdevicemanagement.Client) {
	input := &snowdevicemanagement.ListTasksInput{}

	if len(_snowdevicemanagementMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _snowdevicemanagementMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_snowdevicemanagementNextToken) > 0 {
		input.NextToken = aws.String(_snowdevicemanagementNextToken)
	}
	if len(_snowdevicemanagementState) > 0 {
		if err := assignInputField(input, "State", _snowdevicemanagementState); err != nil {
			log.Errorf("invalid --state: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListTasks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*snowdevicemanagement.ListTasksOutput
	p := snowdevicemanagement.NewListTasksPaginator(client, input)
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

// Adds or replaces tags on a device or task.
func snowdevicemanagement_TagResource(cfg aws.Config, client *snowdevicemanagement.Client) {
	input := &snowdevicemanagement.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_snowdevicemanagementResourceArn) > 0 {
		input.ResourceArn = aws.String(_snowdevicemanagementResourceArn)
	}
	if len(_snowdevicemanagementTags) > 0 {
		if err := assignInputField(input, "Tags", _snowdevicemanagementTags); err != nil {
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

// Removes a tag from a device or task.
func snowdevicemanagement_UntagResource(cfg aws.Config, client *snowdevicemanagement.Client) {
	input := &snowdevicemanagement.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_snowdevicemanagementResourceArn) > 0 {
		input.ResourceArn = aws.String(_snowdevicemanagementResourceArn)
	}
	if len(_snowdevicemanagementTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _snowdevicemanagementTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_snowdevicemanagementCmd)
	_snowdevicemanagementCmd.Flags().SortFlags = false

	_snowdevicemanagementCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_snowdevicemanagementCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_snowdevicemanagementCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_snowdevicemanagementCmd.Flags().StringVarP(&_snowdevicemanagementClientToken, "client-token", "", "", "Client Token")
	_snowdevicemanagementCmd.Flags().StringVarP(&_snowdevicemanagementCommand, "command", "", "", "Command")
	_snowdevicemanagementCmd.Flags().StringVarP(&_snowdevicemanagementDescription, "description", "", "", "Description")
	_snowdevicemanagementCmd.Flags().StringSliceVarP(&_snowdevicemanagementInstanceIds, "instance-ids", "", nil, "Instance Ids")
	_snowdevicemanagementCmd.Flags().StringVarP(&_snowdevicemanagementJobId, "job-id", "", "", "Job ID")
	_snowdevicemanagementCmd.Flags().StringVarP(&_snowdevicemanagementManagedDeviceId, "managed-device-id", "", "", "Managed Device ID")
	_snowdevicemanagementCmd.Flags().StringVarP(&_snowdevicemanagementMaxResults, "max-results", "", "", "Max Results")
	_snowdevicemanagementCmd.Flags().StringVarP(&_snowdevicemanagementNextToken, "next-token", "", "", "Next Token")
	_snowdevicemanagementCmd.Flags().StringVarP(&_snowdevicemanagementResourceArn, "resource-arn", "", "", "Resource ARN")
	_snowdevicemanagementCmd.Flags().StringVarP(&_snowdevicemanagementState, "state", "", "", "State")
	_snowdevicemanagementCmd.Flags().StringSliceVarP(&_snowdevicemanagementTagKeys, "tag-keys", "", nil, "Tag Keys")
	_snowdevicemanagementCmd.Flags().StringVarP(&_snowdevicemanagementTags, "tags", "", "", "Tags")
	_snowdevicemanagementCmd.Flags().StringSliceVarP(&_snowdevicemanagementTargets, "targets", "", nil, "Targets")
	_snowdevicemanagementCmd.Flags().StringVarP(&_snowdevicemanagementTaskId, "task-id", "", "", "Task ID")
	_snowdevicemanagementCmd.Flags().StringVarP(&_snowdevicemanagementType, "type", "", "", "Type")

	_snowdevicemanagementCmd.Flags().BoolVarP(&_snowdevicemanagementCancelTask, "cancel-task", "", false, "Cancel Task")
	_snowdevicemanagementCmd.Flags().BoolVarP(&_snowdevicemanagementCreateTask, "create-task", "", false, "Create Task")
	_snowdevicemanagementCmd.Flags().BoolVarP(&_snowdevicemanagementDescribeDevice, "describe-device", "", false, "Describe Device")
	_snowdevicemanagementCmd.Flags().BoolVarP(&_snowdevicemanagementDescribeDeviceEc2Instances, "describe-device-ec2-instances", "", false, "Describe Device EC2 Instances")
	_snowdevicemanagementCmd.Flags().BoolVarP(&_snowdevicemanagementDescribeExecution, "describe-execution", "", false, "Describe Execution")
	_snowdevicemanagementCmd.Flags().BoolVarP(&_snowdevicemanagementDescribeTask, "describe-task", "", false, "Describe Task")
	_snowdevicemanagementCmd.Flags().BoolVarP(&_snowdevicemanagementListDeviceResources, "list-device-resources", "", false, "List Device Resources")
	_snowdevicemanagementCmd.Flags().BoolVarP(&_snowdevicemanagementListDevices, "list-devices", "", false, "List Devices")
	_snowdevicemanagementCmd.Flags().BoolVarP(&_snowdevicemanagementListExecutions, "list-executions", "", false, "List Executions")
	_snowdevicemanagementCmd.Flags().BoolVarP(&_snowdevicemanagementListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_snowdevicemanagementCmd.Flags().BoolVarP(&_snowdevicemanagementListTasks, "list-tasks", "", false, "List Tasks")
	_snowdevicemanagementCmd.Flags().BoolVarP(&_snowdevicemanagementTagResource, "tag-resource", "", false, "Tag Resource")
	_snowdevicemanagementCmd.Flags().BoolVarP(&_snowdevicemanagementUntagResource, "untag-resource", "", false, "Untag Resource")

}
