package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iotjobsdataplane"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// iotjobsdataplaneCmd represents the iotjobsdataplane command
var _iotjobsdataplaneCmd = &cobra.Command{
	Use:   "iotjobsdataplane",
	Short: "AWS iotjobsdataplane CLI",
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
		client := iotjobsdataplane.NewFromConfig(cfg)
		if _iotjobsdataplaneDescribeJobExecution {
			iotjobsdataplane_DescribeJobExecution(cfg, client)
			return
		}
		if _iotjobsdataplaneGetPendingJobExecutions {
			iotjobsdataplane_GetPendingJobExecutions(cfg, client)
			return
		}
		if _iotjobsdataplaneStartCommandExecution {
			iotjobsdataplane_StartCommandExecution(cfg, client)
			return
		}
		if _iotjobsdataplaneStartNextPendingJobExecution {
			iotjobsdataplane_StartNextPendingJobExecution(cfg, client)
			return
		}
		if _iotjobsdataplaneUpdateJobExecution {
			iotjobsdataplane_UpdateJobExecution(cfg, client)
			return
		}

	},
}

var (
	_iotjobsdataplaneDescribeJobExecution         bool
	_iotjobsdataplaneGetPendingJobExecutions      bool
	_iotjobsdataplaneStartCommandExecution        bool
	_iotjobsdataplaneStartNextPendingJobExecution bool
	_iotjobsdataplaneUpdateJobExecution           bool

	_iotjobsdataplaneClientToken              string
	_iotjobsdataplaneCommandArn               string
	_iotjobsdataplaneExecutionNumber          string
	_iotjobsdataplaneExecutionTimeoutSeconds  string
	_iotjobsdataplaneExpectedVersion          string
	_iotjobsdataplaneIncludeJobDocument       string
	_iotjobsdataplaneIncludeJobExecutionState string
	_iotjobsdataplaneJobId                    string
	_iotjobsdataplaneParameters               string
	_iotjobsdataplaneStatus                   string
	_iotjobsdataplaneStatusDetails            string
	_iotjobsdataplaneStepTimeoutInMinutes     string
	_iotjobsdataplaneTargetArn                string
	_iotjobsdataplaneThingName                string
)

// Gets details of a job execution.
// Requires permission to access the [DescribeJobExecution] action.
//
// [DescribeJobExecution]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iotjobsdataplane_DescribeJobExecution(cfg aws.Config, client *iotjobsdataplane.Client) {
	input := &iotjobsdataplane.DescribeJobExecutionInput{
		// JobId: *string, // Required
		// ThingName: *string, // Required
	}

	if len(_iotjobsdataplaneJobId) > 0 {
		input.JobId = aws.String(_iotjobsdataplaneJobId)
	}
	if len(_iotjobsdataplaneThingName) > 0 {
		input.ThingName = aws.String(_iotjobsdataplaneThingName)
	}
	if len(_iotjobsdataplaneExecutionNumber) > 0 {
		if err := assignInputField(input, "ExecutionNumber", _iotjobsdataplaneExecutionNumber); err != nil {
			log.Errorf("invalid --execution-number: %s", err.Error())
			return
		}
	}
	if len(_iotjobsdataplaneIncludeJobDocument) > 0 {
		if err := assignInputField(input, "IncludeJobDocument", _iotjobsdataplaneIncludeJobDocument); err != nil {
			log.Errorf("invalid --include-job-document: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeJobExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the list of all jobs for a thing that are not in a terminal status.
// Requires permission to access the [GetPendingJobExecutions] action.
//
// [GetPendingJobExecutions]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iotjobsdataplane_GetPendingJobExecutions(cfg aws.Config, client *iotjobsdataplane.Client) {
	input := &iotjobsdataplane.GetPendingJobExecutionsInput{
		// ThingName: *string, // Required
	}

	if len(_iotjobsdataplaneThingName) > 0 {
		input.ThingName = aws.String(_iotjobsdataplaneThingName)
	}

	if resp, err := client.GetPendingJobExecutions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Using the command created with the CreateCommand API, start a command execution
// on a specific device.
func iotjobsdataplane_StartCommandExecution(cfg aws.Config, client *iotjobsdataplane.Client) {
	input := &iotjobsdataplane.StartCommandExecutionInput{
		// CommandArn: *string, // Required
		// TargetArn: *string, // Required
	}

	if len(_iotjobsdataplaneCommandArn) > 0 {
		input.CommandArn = aws.String(_iotjobsdataplaneCommandArn)
	}
	if len(_iotjobsdataplaneTargetArn) > 0 {
		input.TargetArn = aws.String(_iotjobsdataplaneTargetArn)
	}
	if len(_iotjobsdataplaneClientToken) > 0 {
		input.ClientToken = aws.String(_iotjobsdataplaneClientToken)
	}
	if len(_iotjobsdataplaneExecutionTimeoutSeconds) > 0 {
		if err := assignInputField(input, "ExecutionTimeoutSeconds", _iotjobsdataplaneExecutionTimeoutSeconds); err != nil {
			log.Errorf("invalid --execution-timeout-seconds: %s", err.Error())
			return
		}
	}
	if len(_iotjobsdataplaneParameters) > 0 {
		if err := assignInputField(input, "Parameters", _iotjobsdataplaneParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
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

// Gets and starts the next pending (status IN_PROGRESS or QUEUED) job execution
// for a thing.
//
// Requires permission to access the [StartNextPendingJobExecution] action.
//
// [StartNextPendingJobExecution]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iotjobsdataplane_StartNextPendingJobExecution(cfg aws.Config, client *iotjobsdataplane.Client) {
	input := &iotjobsdataplane.StartNextPendingJobExecutionInput{
		// ThingName: *string, // Required
	}

	if len(_iotjobsdataplaneThingName) > 0 {
		input.ThingName = aws.String(_iotjobsdataplaneThingName)
	}
	if len(_iotjobsdataplaneStatusDetails) > 0 {
		if err := assignInputField(input, "StatusDetails", _iotjobsdataplaneStatusDetails); err != nil {
			log.Errorf("invalid --status-details: %s", err.Error())
			return
		}
	}
	if len(_iotjobsdataplaneStepTimeoutInMinutes) > 0 {
		if err := assignInputField(input, "StepTimeoutInMinutes", _iotjobsdataplaneStepTimeoutInMinutes); err != nil {
			log.Errorf("invalid --step-timeout-in-minutes: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartNextPendingJobExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the status of a job execution.
// Requires permission to access the [UpdateJobExecution] action.
//
// [UpdateJobExecution]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiotjobsdataplane.html
func iotjobsdataplane_UpdateJobExecution(cfg aws.Config, client *iotjobsdataplane.Client) {
	input := &iotjobsdataplane.UpdateJobExecutionInput{
		// JobId: *string, // Required
		// Status: types.JobExecutionStatus, // Required
		// ThingName: *string, // Required
	}

	if len(_iotjobsdataplaneJobId) > 0 {
		input.JobId = aws.String(_iotjobsdataplaneJobId)
	}
	if len(_iotjobsdataplaneStatus) > 0 {
		if err := assignInputField(input, "Status", _iotjobsdataplaneStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_iotjobsdataplaneThingName) > 0 {
		input.ThingName = aws.String(_iotjobsdataplaneThingName)
	}
	if len(_iotjobsdataplaneExecutionNumber) > 0 {
		if err := assignInputField(input, "ExecutionNumber", _iotjobsdataplaneExecutionNumber); err != nil {
			log.Errorf("invalid --execution-number: %s", err.Error())
			return
		}
	}
	if len(_iotjobsdataplaneExpectedVersion) > 0 {
		if err := assignInputField(input, "ExpectedVersion", _iotjobsdataplaneExpectedVersion); err != nil {
			log.Errorf("invalid --expected-version: %s", err.Error())
			return
		}
	}
	if len(_iotjobsdataplaneIncludeJobDocument) > 0 {
		if err := assignInputField(input, "IncludeJobDocument", _iotjobsdataplaneIncludeJobDocument); err != nil {
			log.Errorf("invalid --include-job-document: %s", err.Error())
			return
		}
	}
	if len(_iotjobsdataplaneIncludeJobExecutionState) > 0 {
		if err := assignInputField(input, "IncludeJobExecutionState", _iotjobsdataplaneIncludeJobExecutionState); err != nil {
			log.Errorf("invalid --include-job-execution-state: %s", err.Error())
			return
		}
	}
	if len(_iotjobsdataplaneStatusDetails) > 0 {
		if err := assignInputField(input, "StatusDetails", _iotjobsdataplaneStatusDetails); err != nil {
			log.Errorf("invalid --status-details: %s", err.Error())
			return
		}
	}
	if len(_iotjobsdataplaneStepTimeoutInMinutes) > 0 {
		if err := assignInputField(input, "StepTimeoutInMinutes", _iotjobsdataplaneStepTimeoutInMinutes); err != nil {
			log.Errorf("invalid --step-timeout-in-minutes: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateJobExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_iotjobsdataplaneCmd)
	_iotjobsdataplaneCmd.Flags().SortFlags = false

	_iotjobsdataplaneCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_iotjobsdataplaneCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_iotjobsdataplaneCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_iotjobsdataplaneCmd.Flags().StringVarP(&_iotjobsdataplaneClientToken, "client-token", "", "", "Client Token")
	_iotjobsdataplaneCmd.Flags().StringVarP(&_iotjobsdataplaneCommandArn, "command-arn", "", "", "Command ARN")
	_iotjobsdataplaneCmd.Flags().StringVarP(&_iotjobsdataplaneExecutionNumber, "execution-number", "", "", "Execution Number")
	_iotjobsdataplaneCmd.Flags().StringVarP(&_iotjobsdataplaneExecutionTimeoutSeconds, "execution-timeout-seconds", "", "", "Execution Timeout Seconds")
	_iotjobsdataplaneCmd.Flags().StringVarP(&_iotjobsdataplaneExpectedVersion, "expected-version", "", "", "Expected Version")
	_iotjobsdataplaneCmd.Flags().StringVarP(&_iotjobsdataplaneIncludeJobDocument, "include-job-document", "", "", "Include Job Document")
	_iotjobsdataplaneCmd.Flags().StringVarP(&_iotjobsdataplaneIncludeJobExecutionState, "include-job-execution-state", "", "", "Include Job Execution State")
	_iotjobsdataplaneCmd.Flags().StringVarP(&_iotjobsdataplaneJobId, "job-id", "", "", "Job ID")
	_iotjobsdataplaneCmd.Flags().StringVarP(&_iotjobsdataplaneParameters, "parameters", "", "", "Parameters")
	_iotjobsdataplaneCmd.Flags().StringVarP(&_iotjobsdataplaneStatus, "status", "", "", "Status")
	_iotjobsdataplaneCmd.Flags().StringVarP(&_iotjobsdataplaneStatusDetails, "status-details", "", "", "Status Details")
	_iotjobsdataplaneCmd.Flags().StringVarP(&_iotjobsdataplaneStepTimeoutInMinutes, "step-timeout-in-minutes", "", "", "Step Timeout In Minutes")
	_iotjobsdataplaneCmd.Flags().StringVarP(&_iotjobsdataplaneTargetArn, "target-arn", "", "", "Target ARN")
	_iotjobsdataplaneCmd.Flags().StringVarP(&_iotjobsdataplaneThingName, "thing-name", "", "", "Thing Name")

	_iotjobsdataplaneCmd.Flags().BoolVarP(&_iotjobsdataplaneDescribeJobExecution, "describe-job-execution", "", false, "Describe Job Execution")
	_iotjobsdataplaneCmd.Flags().BoolVarP(&_iotjobsdataplaneGetPendingJobExecutions, "get-pending-job-executions", "", false, "Get Pending Job Executions")
	_iotjobsdataplaneCmd.Flags().BoolVarP(&_iotjobsdataplaneStartCommandExecution, "start-command-execution", "", false, "Start Command Execution")
	_iotjobsdataplaneCmd.Flags().BoolVarP(&_iotjobsdataplaneStartNextPendingJobExecution, "start-next-pending-job-execution", "", false, "Start Next Pending Job Execution")
	_iotjobsdataplaneCmd.Flags().BoolVarP(&_iotjobsdataplaneUpdateJobExecution, "update-job-execution", "", false, "Update Job Execution")

}
