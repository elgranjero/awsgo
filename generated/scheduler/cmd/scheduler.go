package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/scheduler"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// schedulerCmd represents the scheduler command
var _schedulerCmd = &cobra.Command{
	Use:   "scheduler",
	Short: "AWS scheduler CLI",
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
		client := scheduler.NewFromConfig(cfg)
		if _schedulerCreateSchedule {
			scheduler_CreateSchedule(cfg, client)
			return
		}
		if _schedulerCreateScheduleGroup {
			scheduler_CreateScheduleGroup(cfg, client)
			return
		}
		if _schedulerDeleteSchedule {
			scheduler_DeleteSchedule(cfg, client)
			return
		}
		if _schedulerDeleteScheduleGroup {
			scheduler_DeleteScheduleGroup(cfg, client)
			return
		}
		if _schedulerGetSchedule {
			scheduler_GetSchedule(cfg, client)
			return
		}
		if _schedulerGetScheduleGroup {
			scheduler_GetScheduleGroup(cfg, client)
			return
		}
		if _schedulerListScheduleGroups {
			scheduler_ListScheduleGroups(cfg, client)
			return
		}
		if _schedulerListSchedules {
			scheduler_ListSchedules(cfg, client)
			return
		}
		if _schedulerListTagsForResource {
			scheduler_ListTagsForResource(cfg, client)
			return
		}
		if _schedulerTagResource {
			scheduler_TagResource(cfg, client)
			return
		}
		if _schedulerUntagResource {
			scheduler_UntagResource(cfg, client)
			return
		}
		if _schedulerUpdateSchedule {
			scheduler_UpdateSchedule(cfg, client)
			return
		}

	},
}

var (
	_schedulerCreateSchedule      bool
	_schedulerCreateScheduleGroup bool
	_schedulerDeleteSchedule      bool
	_schedulerDeleteScheduleGroup bool
	_schedulerGetSchedule         bool
	_schedulerGetScheduleGroup    bool
	_schedulerListScheduleGroups  bool
	_schedulerListSchedules       bool
	_schedulerListTagsForResource bool
	_schedulerTagResource         bool
	_schedulerUntagResource       bool
	_schedulerUpdateSchedule      bool

	_schedulerActionAfterCompletion      string
	_schedulerClientToken                string
	_schedulerDescription                string
	_schedulerEndDate                    string
	_schedulerFlexibleTimeWindow         string
	_schedulerGroupName                  string
	_schedulerKmsKeyArn                  string
	_schedulerMaxResults                 string
	_schedulerName                       string
	_schedulerNamePrefix                 string
	_schedulerNextToken                  string
	_schedulerResourceArn                string
	_schedulerScheduleExpression         string
	_schedulerScheduleExpressionTimezone string
	_schedulerStartDate                  string
	_schedulerState                      string
	_schedulerTagKeys                    []string
	_schedulerTags                       string
	_schedulerTarget                     string
)

// Creates the specified schedule.
func scheduler_CreateSchedule(cfg aws.Config, client *scheduler.Client) {
	input := &scheduler.CreateScheduleInput{
		// FlexibleTimeWindow: *types.FlexibleTimeWindow, // Required
		// Name: *string, // Required
		// ScheduleExpression: *string, // Required
		// Target: *types.Target, // Required
	}

	if len(_schedulerFlexibleTimeWindow) > 0 {
		if err := assignInputField(input, "FlexibleTimeWindow", _schedulerFlexibleTimeWindow); err != nil {
			log.Errorf("invalid --flexible-time-window: %s", err.Error())
			return
		}
	}
	if len(_schedulerName) > 0 {
		input.Name = aws.String(_schedulerName)
	}
	if len(_schedulerScheduleExpression) > 0 {
		input.ScheduleExpression = aws.String(_schedulerScheduleExpression)
	}
	if len(_schedulerTarget) > 0 {
		if err := assignInputField(input, "Target", _schedulerTarget); err != nil {
			log.Errorf("invalid --target: %s", err.Error())
			return
		}
	}
	if len(_schedulerActionAfterCompletion) > 0 {
		if err := assignInputField(input, "ActionAfterCompletion", _schedulerActionAfterCompletion); err != nil {
			log.Errorf("invalid --action-after-completion: %s", err.Error())
			return
		}
	}
	if len(_schedulerClientToken) > 0 {
		input.ClientToken = aws.String(_schedulerClientToken)
	}
	if len(_schedulerDescription) > 0 {
		input.Description = aws.String(_schedulerDescription)
	}
	if len(_schedulerEndDate) > 0 {
		if err := assignInputField(input, "EndDate", _schedulerEndDate); err != nil {
			log.Errorf("invalid --end-date: %s", err.Error())
			return
		}
	}
	if len(_schedulerGroupName) > 0 {
		input.GroupName = aws.String(_schedulerGroupName)
	}
	if len(_schedulerKmsKeyArn) > 0 {
		input.KmsKeyArn = aws.String(_schedulerKmsKeyArn)
	}
	if len(_schedulerScheduleExpressionTimezone) > 0 {
		input.ScheduleExpressionTimezone = aws.String(_schedulerScheduleExpressionTimezone)
	}
	if len(_schedulerStartDate) > 0 {
		if err := assignInputField(input, "StartDate", _schedulerStartDate); err != nil {
			log.Errorf("invalid --start-date: %s", err.Error())
			return
		}
	}
	if len(_schedulerState) > 0 {
		if err := assignInputField(input, "State", _schedulerState); err != nil {
			log.Errorf("invalid --state: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSchedule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates the specified schedule group.
func scheduler_CreateScheduleGroup(cfg aws.Config, client *scheduler.Client) {
	input := &scheduler.CreateScheduleGroupInput{
		// Name: *string, // Required
	}

	if len(_schedulerName) > 0 {
		input.Name = aws.String(_schedulerName)
	}
	if len(_schedulerClientToken) > 0 {
		input.ClientToken = aws.String(_schedulerClientToken)
	}
	if len(_schedulerTags) > 0 {
		if err := assignInputField(input, "Tags", _schedulerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateScheduleGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified schedule.
func scheduler_DeleteSchedule(cfg aws.Config, client *scheduler.Client) {
	input := &scheduler.DeleteScheduleInput{
		// Name: *string, // Required
	}

	if len(_schedulerName) > 0 {
		input.Name = aws.String(_schedulerName)
	}
	if len(_schedulerClientToken) > 0 {
		input.ClientToken = aws.String(_schedulerClientToken)
	}
	if len(_schedulerGroupName) > 0 {
		input.GroupName = aws.String(_schedulerGroupName)
	}

	if resp, err := client.DeleteSchedule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified schedule group. Deleting a schedule group results in
// EventBridge Scheduler deleting all schedules associated with the group. When you
// delete a group, it remains in a DELETING state until all of its associated
// schedules are deleted. Schedules associated with the group that are set to run
// while the schedule group is in the process of being deleted might continue to
// invoke their targets until the schedule group and its associated schedules are
// deleted.
//
// This operation is eventually consistent.
func scheduler_DeleteScheduleGroup(cfg aws.Config, client *scheduler.Client) {
	input := &scheduler.DeleteScheduleGroupInput{
		// Name: *string, // Required
	}

	if len(_schedulerName) > 0 {
		input.Name = aws.String(_schedulerName)
	}
	if len(_schedulerClientToken) > 0 {
		input.ClientToken = aws.String(_schedulerClientToken)
	}

	if resp, err := client.DeleteScheduleGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the specified schedule.
func scheduler_GetSchedule(cfg aws.Config, client *scheduler.Client) {
	input := &scheduler.GetScheduleInput{
		// Name: *string, // Required
	}

	if len(_schedulerName) > 0 {
		input.Name = aws.String(_schedulerName)
	}
	if len(_schedulerGroupName) > 0 {
		input.GroupName = aws.String(_schedulerGroupName)
	}

	if resp, err := client.GetSchedule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the specified schedule group.
func scheduler_GetScheduleGroup(cfg aws.Config, client *scheduler.Client) {
	input := &scheduler.GetScheduleGroupInput{
		// Name: *string, // Required
	}

	if len(_schedulerName) > 0 {
		input.Name = aws.String(_schedulerName)
	}

	if resp, err := client.GetScheduleGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a paginated list of your schedule groups.
func scheduler_ListScheduleGroups(cfg aws.Config, client *scheduler.Client) {
	input := &scheduler.ListScheduleGroupsInput{}

	if len(_schedulerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _schedulerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_schedulerNamePrefix) > 0 {
		input.NamePrefix = aws.String(_schedulerNamePrefix)
	}
	if len(_schedulerNextToken) > 0 {
		input.NextToken = aws.String(_schedulerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListScheduleGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*scheduler.ListScheduleGroupsOutput
	p := scheduler.NewListScheduleGroupsPaginator(client, input)
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

// Returns a paginated list of your EventBridge Scheduler schedules.
func scheduler_ListSchedules(cfg aws.Config, client *scheduler.Client) {
	input := &scheduler.ListSchedulesInput{}

	if len(_schedulerGroupName) > 0 {
		input.GroupName = aws.String(_schedulerGroupName)
	}
	if len(_schedulerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _schedulerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_schedulerNamePrefix) > 0 {
		input.NamePrefix = aws.String(_schedulerNamePrefix)
	}
	if len(_schedulerNextToken) > 0 {
		input.NextToken = aws.String(_schedulerNextToken)
	}
	if len(_schedulerState) > 0 {
		if err := assignInputField(input, "State", _schedulerState); err != nil {
			log.Errorf("invalid --state: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListSchedules(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*scheduler.ListSchedulesOutput
	p := scheduler.NewListSchedulesPaginator(client, input)
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

// Lists the tags associated with the Scheduler resource.
func scheduler_ListTagsForResource(cfg aws.Config, client *scheduler.Client) {
	input := &scheduler.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_schedulerResourceArn) > 0 {
		input.ResourceArn = aws.String(_schedulerResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Assigns one or more tags (key-value pairs) to the specified EventBridge
// Scheduler resource. You can only assign tags to schedule groups.
func scheduler_TagResource(cfg aws.Config, client *scheduler.Client) {
	input := &scheduler.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_schedulerResourceArn) > 0 {
		input.ResourceArn = aws.String(_schedulerResourceArn)
	}
	if len(_schedulerTags) > 0 {
		if err := assignInputField(input, "Tags", _schedulerTags); err != nil {
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

// Removes one or more tags from the specified EventBridge Scheduler schedule
// group.
func scheduler_UntagResource(cfg aws.Config, client *scheduler.Client) {
	input := &scheduler.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_schedulerResourceArn) > 0 {
		input.ResourceArn = aws.String(_schedulerResourceArn)
	}
	if len(_schedulerTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _schedulerTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified schedule. When you call UpdateSchedule , EventBridge
// Scheduler uses all values, including empty values, specified in the request and
// overrides the existing schedule. This is by design. This means that if you do
// not set an optional field in your request, that field will be set to its
// system-default value after the update.
//
// Before calling this operation, we recommend that you call the GetSchedule API
// operation and make a note of all optional parameters for your UpdateSchedule
// call.
func scheduler_UpdateSchedule(cfg aws.Config, client *scheduler.Client) {
	input := &scheduler.UpdateScheduleInput{
		// FlexibleTimeWindow: *types.FlexibleTimeWindow, // Required
		// Name: *string, // Required
		// ScheduleExpression: *string, // Required
		// Target: *types.Target, // Required
	}

	if len(_schedulerFlexibleTimeWindow) > 0 {
		if err := assignInputField(input, "FlexibleTimeWindow", _schedulerFlexibleTimeWindow); err != nil {
			log.Errorf("invalid --flexible-time-window: %s", err.Error())
			return
		}
	}
	if len(_schedulerName) > 0 {
		input.Name = aws.String(_schedulerName)
	}
	if len(_schedulerScheduleExpression) > 0 {
		input.ScheduleExpression = aws.String(_schedulerScheduleExpression)
	}
	if len(_schedulerTarget) > 0 {
		if err := assignInputField(input, "Target", _schedulerTarget); err != nil {
			log.Errorf("invalid --target: %s", err.Error())
			return
		}
	}
	if len(_schedulerActionAfterCompletion) > 0 {
		if err := assignInputField(input, "ActionAfterCompletion", _schedulerActionAfterCompletion); err != nil {
			log.Errorf("invalid --action-after-completion: %s", err.Error())
			return
		}
	}
	if len(_schedulerClientToken) > 0 {
		input.ClientToken = aws.String(_schedulerClientToken)
	}
	if len(_schedulerDescription) > 0 {
		input.Description = aws.String(_schedulerDescription)
	}
	if len(_schedulerEndDate) > 0 {
		if err := assignInputField(input, "EndDate", _schedulerEndDate); err != nil {
			log.Errorf("invalid --end-date: %s", err.Error())
			return
		}
	}
	if len(_schedulerGroupName) > 0 {
		input.GroupName = aws.String(_schedulerGroupName)
	}
	if len(_schedulerKmsKeyArn) > 0 {
		input.KmsKeyArn = aws.String(_schedulerKmsKeyArn)
	}
	if len(_schedulerScheduleExpressionTimezone) > 0 {
		input.ScheduleExpressionTimezone = aws.String(_schedulerScheduleExpressionTimezone)
	}
	if len(_schedulerStartDate) > 0 {
		if err := assignInputField(input, "StartDate", _schedulerStartDate); err != nil {
			log.Errorf("invalid --start-date: %s", err.Error())
			return
		}
	}
	if len(_schedulerState) > 0 {
		if err := assignInputField(input, "State", _schedulerState); err != nil {
			log.Errorf("invalid --state: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateSchedule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_schedulerCmd)
	_schedulerCmd.Flags().SortFlags = false

	_schedulerCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_schedulerCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_schedulerCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_schedulerCmd.Flags().StringVarP(&_schedulerActionAfterCompletion, "action-after-completion", "", "", "Action After Completion")
	_schedulerCmd.Flags().StringVarP(&_schedulerClientToken, "client-token", "", "", "Client Token")
	_schedulerCmd.Flags().StringVarP(&_schedulerDescription, "description", "", "", "Description")
	_schedulerCmd.Flags().StringVarP(&_schedulerEndDate, "end-date", "", "", "End Date")
	_schedulerCmd.Flags().StringVarP(&_schedulerFlexibleTimeWindow, "flexible-time-window", "", "", "Flexible Time Window")
	_schedulerCmd.Flags().StringVarP(&_schedulerGroupName, "group-name", "", "", "Group Name")
	_schedulerCmd.Flags().StringVarP(&_schedulerKmsKeyArn, "kms-key-arn", "", "", "KMS Key ARN")
	_schedulerCmd.Flags().StringVarP(&_schedulerMaxResults, "max-results", "", "", "Max Results")
	_schedulerCmd.Flags().StringVarP(&_schedulerName, "name", "", "", "Name")
	_schedulerCmd.Flags().StringVarP(&_schedulerNamePrefix, "name-prefix", "", "", "Name Prefix")
	_schedulerCmd.Flags().StringVarP(&_schedulerNextToken, "next-token", "", "", "Next Token")
	_schedulerCmd.Flags().StringVarP(&_schedulerResourceArn, "resource-arn", "", "", "Resource ARN")
	_schedulerCmd.Flags().StringVarP(&_schedulerScheduleExpression, "schedule-expression", "", "", "Schedule Expression")
	_schedulerCmd.Flags().StringVarP(&_schedulerScheduleExpressionTimezone, "schedule-expression-timezone", "", "", "Schedule Expression Timezone")
	_schedulerCmd.Flags().StringVarP(&_schedulerStartDate, "start-date", "", "", "Start Date")
	_schedulerCmd.Flags().StringVarP(&_schedulerState, "state", "", "", "State")
	_schedulerCmd.Flags().StringSliceVarP(&_schedulerTagKeys, "tag-keys", "", nil, "Tag Keys")
	_schedulerCmd.Flags().StringVarP(&_schedulerTags, "tags", "", "", "Tags")
	_schedulerCmd.Flags().StringVarP(&_schedulerTarget, "target", "", "", "Target")

	_schedulerCmd.Flags().BoolVarP(&_schedulerCreateSchedule, "create-schedule", "", false, "Create Schedule")
	_schedulerCmd.Flags().BoolVarP(&_schedulerCreateScheduleGroup, "create-schedule-group", "", false, "Create Schedule Group")
	_schedulerCmd.Flags().BoolVarP(&_schedulerDeleteSchedule, "delete-schedule", "", false, "Delete Schedule")
	_schedulerCmd.Flags().BoolVarP(&_schedulerDeleteScheduleGroup, "delete-schedule-group", "", false, "Delete Schedule Group")
	_schedulerCmd.Flags().BoolVarP(&_schedulerGetSchedule, "get-schedule", "", false, "Get Schedule")
	_schedulerCmd.Flags().BoolVarP(&_schedulerGetScheduleGroup, "get-schedule-group", "", false, "Get Schedule Group")
	_schedulerCmd.Flags().BoolVarP(&_schedulerListScheduleGroups, "list-schedule-groups", "", false, "List Schedule Groups")
	_schedulerCmd.Flags().BoolVarP(&_schedulerListSchedules, "list-schedules", "", false, "List Schedules")
	_schedulerCmd.Flags().BoolVarP(&_schedulerListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_schedulerCmd.Flags().BoolVarP(&_schedulerTagResource, "tag-resource", "", false, "Tag Resource")
	_schedulerCmd.Flags().BoolVarP(&_schedulerUntagResource, "untag-resource", "", false, "Untag Resource")
	_schedulerCmd.Flags().BoolVarP(&_schedulerUpdateSchedule, "update-schedule", "", false, "Update Schedule")

}
