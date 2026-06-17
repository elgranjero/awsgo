package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/swf"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// swfCmd represents the swf command
var _swfCmd = &cobra.Command{
	Use:   "swf",
	Short: "AWS swf CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := swf.NewFromConfig(cfg)
		if _swfCountClosedWorkflowExecutions {
			swf_CountClosedWorkflowExecutions(cfg, client)
			return
		}
		if _swfCountOpenWorkflowExecutions {
			swf_CountOpenWorkflowExecutions(cfg, client)
			return
		}
		if _swfCountPendingActivityTasks {
			swf_CountPendingActivityTasks(cfg, client)
			return
		}
		if _swfCountPendingDecisionTasks {
			swf_CountPendingDecisionTasks(cfg, client)
			return
		}
		if _swfDeleteActivityType {
			swf_DeleteActivityType(cfg, client)
			return
		}
		if _swfDeleteWorkflowType {
			swf_DeleteWorkflowType(cfg, client)
			return
		}
		if _swfDeprecateActivityType {
			swf_DeprecateActivityType(cfg, client)
			return
		}
		if _swfDeprecateDomain {
			swf_DeprecateDomain(cfg, client)
			return
		}
		if _swfDeprecateWorkflowType {
			swf_DeprecateWorkflowType(cfg, client)
			return
		}
		if _swfDescribeActivityType {
			swf_DescribeActivityType(cfg, client)
			return
		}
		if _swfDescribeDomain {
			swf_DescribeDomain(cfg, client)
			return
		}
		if _swfDescribeWorkflowExecution {
			swf_DescribeWorkflowExecution(cfg, client)
			return
		}
		if _swfDescribeWorkflowType {
			swf_DescribeWorkflowType(cfg, client)
			return
		}
		if _swfGetWorkflowExecutionHistory {
			swf_GetWorkflowExecutionHistory(cfg, client)
			return
		}
		if _swfListActivityTypes {
			swf_ListActivityTypes(cfg, client)
			return
		}
		if _swfListClosedWorkflowExecutions {
			swf_ListClosedWorkflowExecutions(cfg, client)
			return
		}
		if _swfListDomains {
			swf_ListDomains(cfg, client)
			return
		}
		if _swfListOpenWorkflowExecutions {
			swf_ListOpenWorkflowExecutions(cfg, client)
			return
		}
		if _swfListTagsForResource {
			swf_ListTagsForResource(cfg, client)
			return
		}
		if _swfListWorkflowTypes {
			swf_ListWorkflowTypes(cfg, client)
			return
		}
		if _swfPollForActivityTask {
			swf_PollForActivityTask(cfg, client)
			return
		}
		if _swfPollForDecisionTask {
			swf_PollForDecisionTask(cfg, client)
			return
		}
		if _swfRecordActivityTaskHeartbeat {
			swf_RecordActivityTaskHeartbeat(cfg, client)
			return
		}
		if _swfRegisterActivityType {
			swf_RegisterActivityType(cfg, client)
			return
		}
		if _swfRegisterDomain {
			swf_RegisterDomain(cfg, client)
			return
		}
		if _swfRegisterWorkflowType {
			swf_RegisterWorkflowType(cfg, client)
			return
		}
		if _swfRequestCancelWorkflowExecution {
			swf_RequestCancelWorkflowExecution(cfg, client)
			return
		}
		if _swfRespondActivityTaskCanceled {
			swf_RespondActivityTaskCanceled(cfg, client)
			return
		}
		if _swfRespondActivityTaskCompleted {
			swf_RespondActivityTaskCompleted(cfg, client)
			return
		}
		if _swfRespondActivityTaskFailed {
			swf_RespondActivityTaskFailed(cfg, client)
			return
		}
		if _swfRespondDecisionTaskCompleted {
			swf_RespondDecisionTaskCompleted(cfg, client)
			return
		}
		if _swfSignalWorkflowExecution {
			swf_SignalWorkflowExecution(cfg, client)
			return
		}
		if _swfStartWorkflowExecution {
			swf_StartWorkflowExecution(cfg, client)
			return
		}
		if _swfTagResource {
			swf_TagResource(cfg, client)
			return
		}
		if _swfTerminateWorkflowExecution {
			swf_TerminateWorkflowExecution(cfg, client)
			return
		}
		if _swfUndeprecateActivityType {
			swf_UndeprecateActivityType(cfg, client)
			return
		}
		if _swfUndeprecateDomain {
			swf_UndeprecateDomain(cfg, client)
			return
		}
		if _swfUndeprecateWorkflowType {
			swf_UndeprecateWorkflowType(cfg, client)
			return
		}
		if _swfUntagResource {
			swf_UntagResource(cfg, client)
			return
		}

	},
}

var (
	_swfCountClosedWorkflowExecutions  bool
	_swfCountOpenWorkflowExecutions    bool
	_swfCountPendingActivityTasks      bool
	_swfCountPendingDecisionTasks      bool
	_swfDeleteActivityType             bool
	_swfDeleteWorkflowType             bool
	_swfDeprecateActivityType          bool
	_swfDeprecateDomain                bool
	_swfDeprecateWorkflowType          bool
	_swfDescribeActivityType           bool
	_swfDescribeDomain                 bool
	_swfDescribeWorkflowExecution      bool
	_swfDescribeWorkflowType           bool
	_swfGetWorkflowExecutionHistory    bool
	_swfListActivityTypes              bool
	_swfListClosedWorkflowExecutions   bool
	_swfListDomains                    bool
	_swfListOpenWorkflowExecutions     bool
	_swfListTagsForResource            bool
	_swfListWorkflowTypes              bool
	_swfPollForActivityTask            bool
	_swfPollForDecisionTask            bool
	_swfRecordActivityTaskHeartbeat    bool
	_swfRegisterActivityType           bool
	_swfRegisterDomain                 bool
	_swfRegisterWorkflowType           bool
	_swfRequestCancelWorkflowExecution bool
	_swfRespondActivityTaskCanceled    bool
	_swfRespondActivityTaskCompleted   bool
	_swfRespondActivityTaskFailed      bool
	_swfRespondDecisionTaskCompleted   bool
	_swfSignalWorkflowExecution        bool
	_swfStartWorkflowExecution         bool
	_swfTagResource                    bool
	_swfTerminateWorkflowExecution     bool
	_swfUndeprecateActivityType        bool
	_swfUndeprecateDomain              bool
	_swfUndeprecateWorkflowType        bool
	_swfUntagResource                  bool

	_swfActivityType                           string
	_swfChildPolicy                            string
	_swfCloseStatusFilter                      string
	_swfCloseTimeFilter                        string
	_swfDecisions                              string
	_swfDefaultChildPolicy                     string
	_swfDefaultExecutionStartToCloseTimeout    string
	_swfDefaultLambdaRole                      string
	_swfDefaultTaskHeartbeatTimeout            string
	_swfDefaultTaskList                        string
	_swfDefaultTaskPriority                    string
	_swfDefaultTaskScheduleToCloseTimeout      string
	_swfDefaultTaskScheduleToStartTimeout      string
	_swfDefaultTaskStartToCloseTimeout         string
	_swfDescription                            string
	_swfDetails                                string
	_swfDomain                                 string
	_swfExecution                              string
	_swfExecutionContext                       string
	_swfExecutionFilter                        string
	_swfExecutionStartToCloseTimeout           string
	_swfIdentity                               string
	_swfInput                                  string
	_swfLambdaRole                             string
	_swfMaximumPageSize                        string
	_swfName                                   string
	_swfNextPageToken                          string
	_swfReason                                 string
	_swfRegistrationStatus                     string
	_swfResourceArn                            string
	_swfResult                                 string
	_swfReverseOrder                           string
	_swfRunId                                  string
	_swfSignalName                             string
	_swfStartAtPreviousStartedEvent            string
	_swfStartTimeFilter                        string
	_swfTagFilter                              string
	_swfTagKeys                                []string
	_swfTagList                                []string
	_swfTags                                   string
	_swfTaskList                               string
	_swfTaskListScheduleToStartTimeout         string
	_swfTaskPriority                           string
	_swfTaskStartToCloseTimeout                string
	_swfTaskToken                              string
	_swfTypeFilter                             string
	_swfVersion                                string
	_swfWorkflowExecutionRetentionPeriodInDays string
	_swfWorkflowId                             string
	_swfWorkflowType                           string
)

// Returns the number of closed workflow executions within the given domain that
// meet the specified filtering criteria.
//
// This operation is eventually consistent. The results are best effort and may
// not exactly reflect recent updates and changes.
//
// # Access Control
//
// You can use IAM policies to control this action's access to Amazon SWF
// resources as follows:
//
// - Use a Resource element with the domain name to limit the action to only
// specified domains.
//
// - Use an Action element to allow or deny permission to call this action.
//
// - Constrain the following parameters by using a Condition element with the
// appropriate keys.
//
// - tagFilter.tag : String constraint. The key is swf:tagFilter.tag .
//
// - typeFilter.name : String constraint. The key is swf:typeFilter.name .
//
// - typeFilter.version : String constraint. The key is swf:typeFilter.version .
//
// If the caller doesn't have sufficient permissions to invoke the action, or the
// parameter values fall outside the specified constraints, the action fails. The
// associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED .
// For details and example IAM policies, see [Using IAM to Manage Access to Amazon SWF Workflows]in the Amazon SWF Developer Guide.
//
// [Using IAM to Manage Access to Amazon SWF Workflows]: https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html
func swf_CountClosedWorkflowExecutions(cfg aws.Config, client *swf.Client) {
	input := &swf.CountClosedWorkflowExecutionsInput{
		// Domain: *string, // Required
	}

	if len(_swfDomain) > 0 {
		input.Domain = aws.String(_swfDomain)
	}
	if len(_swfCloseStatusFilter) > 0 {
		if err := assignInputField(input, "CloseStatusFilter", _swfCloseStatusFilter); err != nil {
			log.Errorf("invalid --close-status-filter: %s", err.Error())
			return
		}
	}
	if len(_swfCloseTimeFilter) > 0 {
		if err := assignInputField(input, "CloseTimeFilter", _swfCloseTimeFilter); err != nil {
			log.Errorf("invalid --close-time-filter: %s", err.Error())
			return
		}
	}
	if len(_swfExecutionFilter) > 0 {
		if err := assignInputField(input, "ExecutionFilter", _swfExecutionFilter); err != nil {
			log.Errorf("invalid --execution-filter: %s", err.Error())
			return
		}
	}
	if len(_swfStartTimeFilter) > 0 {
		if err := assignInputField(input, "StartTimeFilter", _swfStartTimeFilter); err != nil {
			log.Errorf("invalid --start-time-filter: %s", err.Error())
			return
		}
	}
	if len(_swfTagFilter) > 0 {
		if err := assignInputField(input, "TagFilter", _swfTagFilter); err != nil {
			log.Errorf("invalid --tag-filter: %s", err.Error())
			return
		}
	}
	if len(_swfTypeFilter) > 0 {
		if err := assignInputField(input, "TypeFilter", _swfTypeFilter); err != nil {
			log.Errorf("invalid --type-filter: %s", err.Error())
			return
		}
	}

	if resp, err := client.CountClosedWorkflowExecutions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the number of open workflow executions within the given domain that
// meet the specified filtering criteria.
//
// This operation is eventually consistent. The results are best effort and may
// not exactly reflect recent updates and changes.
//
// # Access Control
//
// You can use IAM policies to control this action's access to Amazon SWF
// resources as follows:
//
// - Use a Resource element with the domain name to limit the action to only
// specified domains.
//
// - Use an Action element to allow or deny permission to call this action.
//
// - Constrain the following parameters by using a Condition element with the
// appropriate keys.
//
// - tagFilter.tag : String constraint. The key is swf:tagFilter.tag .
//
// - typeFilter.name : String constraint. The key is swf:typeFilter.name .
//
// - typeFilter.version : String constraint. The key is swf:typeFilter.version .
//
// If the caller doesn't have sufficient permissions to invoke the action, or the
// parameter values fall outside the specified constraints, the action fails. The
// associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED .
// For details and example IAM policies, see [Using IAM to Manage Access to Amazon SWF Workflows]in the Amazon SWF Developer Guide.
//
// [Using IAM to Manage Access to Amazon SWF Workflows]: https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html
func swf_CountOpenWorkflowExecutions(cfg aws.Config, client *swf.Client) {
	input := &swf.CountOpenWorkflowExecutionsInput{
		// Domain: *string, // Required
		// StartTimeFilter: *types.ExecutionTimeFilter, // Required
	}

	if len(_swfDomain) > 0 {
		input.Domain = aws.String(_swfDomain)
	}
	if len(_swfStartTimeFilter) > 0 {
		if err := assignInputField(input, "StartTimeFilter", _swfStartTimeFilter); err != nil {
			log.Errorf("invalid --start-time-filter: %s", err.Error())
			return
		}
	}
	if len(_swfExecutionFilter) > 0 {
		if err := assignInputField(input, "ExecutionFilter", _swfExecutionFilter); err != nil {
			log.Errorf("invalid --execution-filter: %s", err.Error())
			return
		}
	}
	if len(_swfTagFilter) > 0 {
		if err := assignInputField(input, "TagFilter", _swfTagFilter); err != nil {
			log.Errorf("invalid --tag-filter: %s", err.Error())
			return
		}
	}
	if len(_swfTypeFilter) > 0 {
		if err := assignInputField(input, "TypeFilter", _swfTypeFilter); err != nil {
			log.Errorf("invalid --type-filter: %s", err.Error())
			return
		}
	}

	if resp, err := client.CountOpenWorkflowExecutions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the estimated number of activity tasks in the specified task list. The
// count returned is an approximation and isn't guaranteed to be exact. If you
// specify a task list that no activity task was ever scheduled in then 0 is
// returned.
//
// # Access Control
//
// You can use IAM policies to control this action's access to Amazon SWF
// resources as follows:
//
// - Use a Resource element with the domain name to limit the action to only
// specified domains.
//
// - Use an Action element to allow or deny permission to call this action.
//
// - Constrain the taskList.name parameter by using a Condition element with the
// swf:taskList.name key to allow the action to access only certain task lists.
//
// If the caller doesn't have sufficient permissions to invoke the action, or the
// parameter values fall outside the specified constraints, the action fails. The
// associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED .
// For details and example IAM policies, see [Using IAM to Manage Access to Amazon SWF Workflows]in the Amazon SWF Developer Guide.
//
// [Using IAM to Manage Access to Amazon SWF Workflows]: https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html
func swf_CountPendingActivityTasks(cfg aws.Config, client *swf.Client) {
	input := &swf.CountPendingActivityTasksInput{
		// Domain: *string, // Required
		// TaskList: *types.TaskList, // Required
	}

	if len(_swfDomain) > 0 {
		input.Domain = aws.String(_swfDomain)
	}
	if len(_swfTaskList) > 0 {
		if err := assignInputField(input, "TaskList", _swfTaskList); err != nil {
			log.Errorf("invalid --task-list: %s", err.Error())
			return
		}
	}

	if resp, err := client.CountPendingActivityTasks(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the estimated number of decision tasks in the specified task list. The
// count returned is an approximation and isn't guaranteed to be exact. If you
// specify a task list that no decision task was ever scheduled in then 0 is
// returned.
//
// # Access Control
//
// You can use IAM policies to control this action's access to Amazon SWF
// resources as follows:
//
// - Use a Resource element with the domain name to limit the action to only
// specified domains.
//
// - Use an Action element to allow or deny permission to call this action.
//
// - Constrain the taskList.name parameter by using a Condition element with the
// swf:taskList.name key to allow the action to access only certain task lists.
//
// If the caller doesn't have sufficient permissions to invoke the action, or the
// parameter values fall outside the specified constraints, the action fails. The
// associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED .
// For details and example IAM policies, see [Using IAM to Manage Access to Amazon SWF Workflows]in the Amazon SWF Developer Guide.
//
// [Using IAM to Manage Access to Amazon SWF Workflows]: https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html
func swf_CountPendingDecisionTasks(cfg aws.Config, client *swf.Client) {
	input := &swf.CountPendingDecisionTasksInput{
		// Domain: *string, // Required
		// TaskList: *types.TaskList, // Required
	}

	if len(_swfDomain) > 0 {
		input.Domain = aws.String(_swfDomain)
	}
	if len(_swfTaskList) > 0 {
		if err := assignInputField(input, "TaskList", _swfTaskList); err != nil {
			log.Errorf("invalid --task-list: %s", err.Error())
			return
		}
	}

	if resp, err := client.CountPendingDecisionTasks(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified activity type.
// Note: Prior to deletion, activity types must first be deprecated.
//
// After an activity type has been deleted, you cannot schedule new activities of
// that type. Activities that started before the type was deleted will continue to
// run.
//
// # Access Control
//
// You can use IAM policies to control this action's access to Amazon SWF
// resources as follows:
//
// - Use a Resource element with the domain name to limit the action to only
// specified domains.
//
// - Use an Action element to allow or deny permission to call this action.
//
// - Constrain the following parameters by using a Condition element with the
// appropriate keys.
//
// - activityType.name : String constraint. The key is swf:activityType.name .
//
// - activityType.version : String constraint. The key is
// swf:activityType.version .
//
// If the caller doesn't have sufficient permissions to invoke the action, or the
// parameter values fall outside the specified constraints, the action fails. The
// associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED .
// For details and example IAM policies, see [Using IAM to Manage Access to Amazon SWF Workflows]in the Amazon SWF Developer Guide.
//
// [Using IAM to Manage Access to Amazon SWF Workflows]: https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html
func swf_DeleteActivityType(cfg aws.Config, client *swf.Client) {
	input := &swf.DeleteActivityTypeInput{
		// ActivityType: *types.ActivityType, // Required
		// Domain: *string, // Required
	}

	if len(_swfActivityType) > 0 {
		if err := assignInputField(input, "ActivityType", _swfActivityType); err != nil {
			log.Errorf("invalid --activity-type: %s", err.Error())
			return
		}
	}
	if len(_swfDomain) > 0 {
		input.Domain = aws.String(_swfDomain)
	}

	if resp, err := client.DeleteActivityType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified workflow type.
// Note: Prior to deletion, workflow types must first be deprecated.
//
// After a workflow type has been deleted, you cannot create new executions of
// that type. Executions that started before the type was deleted will continue to
// run.
//
// # Access Control
//
// You can use IAM policies to control this action's access to Amazon SWF
// resources as follows:
//
// - Use a Resource element with the domain name to limit the action to only
// specified domains.
//
// - Use an Action element to allow or deny permission to call this action.
//
// - Constrain the following parameters by using a Condition element with the
// appropriate keys.
//
// - workflowType.name : String constraint. The key is swf:workflowType.name .
//
// - workflowType.version : String constraint. The key is
// swf:workflowType.version .
//
// If the caller doesn't have sufficient permissions to invoke the action, or the
// parameter values fall outside the specified constraints, the action fails. The
// associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED .
// For details and example IAM policies, see [Using IAM to Manage Access to Amazon SWF Workflows]in the Amazon SWF Developer Guide.
//
// [Using IAM to Manage Access to Amazon SWF Workflows]: https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html
func swf_DeleteWorkflowType(cfg aws.Config, client *swf.Client) {
	input := &swf.DeleteWorkflowTypeInput{
		// Domain: *string, // Required
		// WorkflowType: *types.WorkflowType, // Required
	}

	if len(_swfDomain) > 0 {
		input.Domain = aws.String(_swfDomain)
	}
	if len(_swfWorkflowType) > 0 {
		if err := assignInputField(input, "WorkflowType", _swfWorkflowType); err != nil {
			log.Errorf("invalid --workflow-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteWorkflowType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deprecates the specified activity type. After an activity type has been
// deprecated, you cannot create new tasks of that activity type. Tasks of this
// type that were scheduled before the type was deprecated continue to run.
//
// # Access Control
//
// You can use IAM policies to control this action's access to Amazon SWF
// resources as follows:
//
// - Use a Resource element with the domain name to limit the action to only
// specified domains.
//
// - Use an Action element to allow or deny permission to call this action.
//
// - Constrain the following parameters by using a Condition element with the
// appropriate keys.
//
// - activityType.name : String constraint. The key is swf:activityType.name .
//
// - activityType.version : String constraint. The key is
// swf:activityType.version .
//
// If the caller doesn't have sufficient permissions to invoke the action, or the
// parameter values fall outside the specified constraints, the action fails. The
// associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED .
// For details and example IAM policies, see [Using IAM to Manage Access to Amazon SWF Workflows]in the Amazon SWF Developer Guide.
//
// [Using IAM to Manage Access to Amazon SWF Workflows]: https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html
func swf_DeprecateActivityType(cfg aws.Config, client *swf.Client) {
	input := &swf.DeprecateActivityTypeInput{
		// ActivityType: *types.ActivityType, // Required
		// Domain: *string, // Required
	}

	if len(_swfActivityType) > 0 {
		if err := assignInputField(input, "ActivityType", _swfActivityType); err != nil {
			log.Errorf("invalid --activity-type: %s", err.Error())
			return
		}
	}
	if len(_swfDomain) > 0 {
		input.Domain = aws.String(_swfDomain)
	}

	if resp, err := client.DeprecateActivityType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deprecates the specified domain. After a domain has been deprecated it cannot
// be used to create new workflow executions or register new types. However, you
// can still use visibility actions on this domain. Deprecating a domain also
// deprecates all activity and workflow types registered in the domain. Executions
// that were started before the domain was deprecated continues to run.
//
// This operation is eventually consistent. The results are best effort and may
// not exactly reflect recent updates and changes.
//
// # Access Control
//
// You can use IAM policies to control this action's access to Amazon SWF
// resources as follows:
//
// - Use a Resource element with the domain name to limit the action to only
// specified domains.
//
// - Use an Action element to allow or deny permission to call this action.
//
// - You cannot use an IAM policy to constrain this action's parameters.
//
// If the caller doesn't have sufficient permissions to invoke the action, or the
// parameter values fall outside the specified constraints, the action fails. The
// associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED .
// For details and example IAM policies, see [Using IAM to Manage Access to Amazon SWF Workflows]in the Amazon SWF Developer Guide.
//
// [Using IAM to Manage Access to Amazon SWF Workflows]: https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html
func swf_DeprecateDomain(cfg aws.Config, client *swf.Client) {
	input := &swf.DeprecateDomainInput{
		// Name: *string, // Required
	}

	if len(_swfName) > 0 {
		input.Name = aws.String(_swfName)
	}

	if resp, err := client.DeprecateDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deprecates the specified workflow type. After a workflow type has been
// deprecated, you cannot create new executions of that type. Executions that were
// started before the type was deprecated continues to run. A deprecated workflow
// type may still be used when calling visibility actions.
//
// This operation is eventually consistent. The results are best effort and may
// not exactly reflect recent updates and changes.
//
// # Access Control
//
// You can use IAM policies to control this action's access to Amazon SWF
// resources as follows:
//
// - Use a Resource element with the domain name to limit the action to only
// specified domains.
//
// - Use an Action element to allow or deny permission to call this action.
//
// - Constrain the following parameters by using a Condition element with the
// appropriate keys.
//
// - workflowType.name : String constraint. The key is swf:workflowType.name .
//
// - workflowType.version : String constraint. The key is
// swf:workflowType.version .
//
// If the caller doesn't have sufficient permissions to invoke the action, or the
// parameter values fall outside the specified constraints, the action fails. The
// associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED .
// For details and example IAM policies, see [Using IAM to Manage Access to Amazon SWF Workflows]in the Amazon SWF Developer Guide.
//
// [Using IAM to Manage Access to Amazon SWF Workflows]: https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html
func swf_DeprecateWorkflowType(cfg aws.Config, client *swf.Client) {
	input := &swf.DeprecateWorkflowTypeInput{
		// Domain: *string, // Required
		// WorkflowType: *types.WorkflowType, // Required
	}

	if len(_swfDomain) > 0 {
		input.Domain = aws.String(_swfDomain)
	}
	if len(_swfWorkflowType) > 0 {
		if err := assignInputField(input, "WorkflowType", _swfWorkflowType); err != nil {
			log.Errorf("invalid --workflow-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeprecateWorkflowType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the specified activity type. This includes
// configuration settings provided when the type was registered and other general
// information about the type.
//
// # Access Control
//
// You can use IAM policies to control this action's access to Amazon SWF
// resources as follows:
//
// - Use a Resource element with the domain name to limit the action to only
// specified domains.
//
// - Use an Action element to allow or deny permission to call this action.
//
// - Constrain the following parameters by using a Condition element with the
// appropriate keys.
//
// - activityType.name : String constraint. The key is swf:activityType.name .
//
// - activityType.version : String constraint. The key is
// swf:activityType.version .
//
// If the caller doesn't have sufficient permissions to invoke the action, or the
// parameter values fall outside the specified constraints, the action fails. The
// associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED .
// For details and example IAM policies, see [Using IAM to Manage Access to Amazon SWF Workflows]in the Amazon SWF Developer Guide.
//
// [Using IAM to Manage Access to Amazon SWF Workflows]: https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html
func swf_DescribeActivityType(cfg aws.Config, client *swf.Client) {
	input := &swf.DescribeActivityTypeInput{
		// ActivityType: *types.ActivityType, // Required
		// Domain: *string, // Required
	}

	if len(_swfActivityType) > 0 {
		if err := assignInputField(input, "ActivityType", _swfActivityType); err != nil {
			log.Errorf("invalid --activity-type: %s", err.Error())
			return
		}
	}
	if len(_swfDomain) > 0 {
		input.Domain = aws.String(_swfDomain)
	}

	if resp, err := client.DescribeActivityType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the specified domain, including description and
// status.
//
// # Access Control
//
// You can use IAM policies to control this action's access to Amazon SWF
// resources as follows:
//
// - Use a Resource element with the domain name to limit the action to only
// specified domains.
//
// - Use an Action element to allow or deny permission to call this action.
//
// - You cannot use an IAM policy to constrain this action's parameters.
//
// If the caller doesn't have sufficient permissions to invoke the action, or the
// parameter values fall outside the specified constraints, the action fails. The
// associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED .
// For details and example IAM policies, see [Using IAM to Manage Access to Amazon SWF Workflows]in the Amazon SWF Developer Guide.
//
// [Using IAM to Manage Access to Amazon SWF Workflows]: https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html
func swf_DescribeDomain(cfg aws.Config, client *swf.Client) {
	input := &swf.DescribeDomainInput{
		// Name: *string, // Required
	}

	if len(_swfName) > 0 {
		input.Name = aws.String(_swfName)
	}

	if resp, err := client.DescribeDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the specified workflow execution including its type
// and some statistics.
//
// This operation is eventually consistent. The results are best effort and may
// not exactly reflect recent updates and changes.
//
// # Access Control
//
// You can use IAM policies to control this action's access to Amazon SWF
// resources as follows:
//
// - Use a Resource element with the domain name to limit the action to only
// specified domains.
//
// - Use an Action element to allow or deny permission to call this action.
//
// - You cannot use an IAM policy to constrain this action's parameters.
//
// If the caller doesn't have sufficient permissions to invoke the action, or the
// parameter values fall outside the specified constraints, the action fails. The
// associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED .
// For details and example IAM policies, see [Using IAM to Manage Access to Amazon SWF Workflows]in the Amazon SWF Developer Guide.
//
// [Using IAM to Manage Access to Amazon SWF Workflows]: https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html
func swf_DescribeWorkflowExecution(cfg aws.Config, client *swf.Client) {
	input := &swf.DescribeWorkflowExecutionInput{
		// Domain: *string, // Required
		// Execution: *types.WorkflowExecution, // Required
	}

	if len(_swfDomain) > 0 {
		input.Domain = aws.String(_swfDomain)
	}
	if len(_swfExecution) > 0 {
		if err := assignInputField(input, "Execution", _swfExecution); err != nil {
			log.Errorf("invalid --execution: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeWorkflowExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the specified workflow type. This includes
// configuration settings specified when the type was registered and other
// information such as creation date, current status, etc.
//
// # Access Control
//
// You can use IAM policies to control this action's access to Amazon SWF
// resources as follows:
//
// - Use a Resource element with the domain name to limit the action to only
// specified domains.
//
// - Use an Action element to allow or deny permission to call this action.
//
// - Constrain the following parameters by using a Condition element with the
// appropriate keys.
//
// - workflowType.name : String constraint. The key is swf:workflowType.name .
//
// - workflowType.version : String constraint. The key is
// swf:workflowType.version .
//
// If the caller doesn't have sufficient permissions to invoke the action, or the
// parameter values fall outside the specified constraints, the action fails. The
// associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED .
// For details and example IAM policies, see [Using IAM to Manage Access to Amazon SWF Workflows]in the Amazon SWF Developer Guide.
//
// [Using IAM to Manage Access to Amazon SWF Workflows]: https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html
func swf_DescribeWorkflowType(cfg aws.Config, client *swf.Client) {
	input := &swf.DescribeWorkflowTypeInput{
		// Domain: *string, // Required
		// WorkflowType: *types.WorkflowType, // Required
	}

	if len(_swfDomain) > 0 {
		input.Domain = aws.String(_swfDomain)
	}
	if len(_swfWorkflowType) > 0 {
		if err := assignInputField(input, "WorkflowType", _swfWorkflowType); err != nil {
			log.Errorf("invalid --workflow-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeWorkflowType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the history of the specified workflow execution. The results may be
// split into multiple pages. To retrieve subsequent pages, make the call again
// using the nextPageToken returned by the initial call.
//
// This operation is eventually consistent. The results are best effort and may
// not exactly reflect recent updates and changes.
//
// # Access Control
//
// You can use IAM policies to control this action's access to Amazon SWF
// resources as follows:
//
// - Use a Resource element with the domain name to limit the action to only
// specified domains.
//
// - Use an Action element to allow or deny permission to call this action.
//
// - You cannot use an IAM policy to constrain this action's parameters.
//
// If the caller doesn't have sufficient permissions to invoke the action, or the
// parameter values fall outside the specified constraints, the action fails. The
// associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED .
// For details and example IAM policies, see [Using IAM to Manage Access to Amazon SWF Workflows]in the Amazon SWF Developer Guide.
//
// [Using IAM to Manage Access to Amazon SWF Workflows]: https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html
func swf_GetWorkflowExecutionHistory(cfg aws.Config, client *swf.Client) {
	input := &swf.GetWorkflowExecutionHistoryInput{
		// Domain: *string, // Required
		// Execution: *types.WorkflowExecution, // Required
	}

	if len(_swfDomain) > 0 {
		input.Domain = aws.String(_swfDomain)
	}
	if len(_swfExecution) > 0 {
		if err := assignInputField(input, "Execution", _swfExecution); err != nil {
			log.Errorf("invalid --execution: %s", err.Error())
			return
		}
	}
	if len(_swfMaximumPageSize) > 0 {
		if err := assignInputField(input, "MaximumPageSize", _swfMaximumPageSize); err != nil {
			log.Errorf("invalid --maximum-page-size: %s", err.Error())
			return
		}
	}
	if len(_swfNextPageToken) > 0 {
		input.NextPageToken = aws.String(_swfNextPageToken)
	}
	if len(_swfReverseOrder) > 0 {
		if err := assignInputField(input, "ReverseOrder", _swfReverseOrder); err != nil {
			log.Errorf("invalid --reverse-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.GetWorkflowExecutionHistory(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*swf.GetWorkflowExecutionHistoryOutput
	p := swf.NewGetWorkflowExecutionHistoryPaginator(client, input)
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

// Returns information about all activities registered in the specified domain
// that match the specified name and registration status. The result includes
// information like creation date, current status of the activity, etc. The results
// may be split into multiple pages. To retrieve subsequent pages, make the call
// again using the nextPageToken returned by the initial call.
//
// # Access Control
//
// You can use IAM policies to control this action's access to Amazon SWF
// resources as follows:
//
// - Use a Resource element with the domain name to limit the action to only
// specified domains.
//
// - Use an Action element to allow or deny permission to call this action.
//
// - You cannot use an IAM policy to constrain this action's parameters.
//
// If the caller doesn't have sufficient permissions to invoke the action, or the
// parameter values fall outside the specified constraints, the action fails. The
// associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED .
// For details and example IAM policies, see [Using IAM to Manage Access to Amazon SWF Workflows]in the Amazon SWF Developer Guide.
//
// [Using IAM to Manage Access to Amazon SWF Workflows]: https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html
func swf_ListActivityTypes(cfg aws.Config, client *swf.Client) {
	input := &swf.ListActivityTypesInput{
		// Domain: *string, // Required
		// RegistrationStatus: types.RegistrationStatus, // Required
	}

	if len(_swfDomain) > 0 {
		input.Domain = aws.String(_swfDomain)
	}
	if len(_swfRegistrationStatus) > 0 {
		if err := assignInputField(input, "RegistrationStatus", _swfRegistrationStatus); err != nil {
			log.Errorf("invalid --registration-status: %s", err.Error())
			return
		}
	}
	if len(_swfMaximumPageSize) > 0 {
		if err := assignInputField(input, "MaximumPageSize", _swfMaximumPageSize); err != nil {
			log.Errorf("invalid --maximum-page-size: %s", err.Error())
			return
		}
	}
	if len(_swfName) > 0 {
		input.Name = aws.String(_swfName)
	}
	if len(_swfNextPageToken) > 0 {
		input.NextPageToken = aws.String(_swfNextPageToken)
	}
	if len(_swfReverseOrder) > 0 {
		if err := assignInputField(input, "ReverseOrder", _swfReverseOrder); err != nil {
			log.Errorf("invalid --reverse-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListActivityTypes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*swf.ListActivityTypesOutput
	p := swf.NewListActivityTypesPaginator(client, input)
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

// Returns a list of closed workflow executions in the specified domain that meet
// the filtering criteria. The results may be split into multiple pages. To
// retrieve subsequent pages, make the call again using the nextPageToken returned
// by the initial call.
//
// This operation is eventually consistent. The results are best effort and may
// not exactly reflect recent updates and changes.
//
// # Access Control
//
// You can use IAM policies to control this action's access to Amazon SWF
// resources as follows:
//
// - Use a Resource element with the domain name to limit the action to only
// specified domains.
//
// - Use an Action element to allow or deny permission to call this action.
//
// - Constrain the following parameters by using a Condition element with the
// appropriate keys.
//
// - tagFilter.tag : String constraint. The key is swf:tagFilter.tag .
//
// - typeFilter.name : String constraint. The key is swf:typeFilter.name .
//
// - typeFilter.version : String constraint. The key is swf:typeFilter.version .
//
// If the caller doesn't have sufficient permissions to invoke the action, or the
// parameter values fall outside the specified constraints, the action fails. The
// associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED .
// For details and example IAM policies, see [Using IAM to Manage Access to Amazon SWF Workflows]in the Amazon SWF Developer Guide.
//
// [Using IAM to Manage Access to Amazon SWF Workflows]: https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html
func swf_ListClosedWorkflowExecutions(cfg aws.Config, client *swf.Client) {
	input := &swf.ListClosedWorkflowExecutionsInput{
		// Domain: *string, // Required
	}

	if len(_swfDomain) > 0 {
		input.Domain = aws.String(_swfDomain)
	}
	if len(_swfCloseStatusFilter) > 0 {
		if err := assignInputField(input, "CloseStatusFilter", _swfCloseStatusFilter); err != nil {
			log.Errorf("invalid --close-status-filter: %s", err.Error())
			return
		}
	}
	if len(_swfCloseTimeFilter) > 0 {
		if err := assignInputField(input, "CloseTimeFilter", _swfCloseTimeFilter); err != nil {
			log.Errorf("invalid --close-time-filter: %s", err.Error())
			return
		}
	}
	if len(_swfExecutionFilter) > 0 {
		if err := assignInputField(input, "ExecutionFilter", _swfExecutionFilter); err != nil {
			log.Errorf("invalid --execution-filter: %s", err.Error())
			return
		}
	}
	if len(_swfMaximumPageSize) > 0 {
		if err := assignInputField(input, "MaximumPageSize", _swfMaximumPageSize); err != nil {
			log.Errorf("invalid --maximum-page-size: %s", err.Error())
			return
		}
	}
	if len(_swfNextPageToken) > 0 {
		input.NextPageToken = aws.String(_swfNextPageToken)
	}
	if len(_swfReverseOrder) > 0 {
		if err := assignInputField(input, "ReverseOrder", _swfReverseOrder); err != nil {
			log.Errorf("invalid --reverse-order: %s", err.Error())
			return
		}
	}
	if len(_swfStartTimeFilter) > 0 {
		if err := assignInputField(input, "StartTimeFilter", _swfStartTimeFilter); err != nil {
			log.Errorf("invalid --start-time-filter: %s", err.Error())
			return
		}
	}
	if len(_swfTagFilter) > 0 {
		if err := assignInputField(input, "TagFilter", _swfTagFilter); err != nil {
			log.Errorf("invalid --tag-filter: %s", err.Error())
			return
		}
	}
	if len(_swfTypeFilter) > 0 {
		if err := assignInputField(input, "TypeFilter", _swfTypeFilter); err != nil {
			log.Errorf("invalid --type-filter: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListClosedWorkflowExecutions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*swf.ListClosedWorkflowExecutionsOutput
	p := swf.NewListClosedWorkflowExecutionsPaginator(client, input)
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

// Returns the list of domains registered in the account. The results may be split
// into multiple pages. To retrieve subsequent pages, make the call again using the
// nextPageToken returned by the initial call.
//
// This operation is eventually consistent. The results are best effort and may
// not exactly reflect recent updates and changes.
//
// # Access Control
//
// You can use IAM policies to control this action's access to Amazon SWF
// resources as follows:
//
// - Use a Resource element with the domain name to limit the action to only
// specified domains. The element must be set to arn:aws:swf::AccountID:domain/*
// , where AccountID is the account ID, with no dashes.
//
// - Use an Action element to allow or deny permission to call this action.
//
// - You cannot use an IAM policy to constrain this action's parameters.
//
// If the caller doesn't have sufficient permissions to invoke the action, or the
// parameter values fall outside the specified constraints, the action fails. The
// associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED .
// For details and example IAM policies, see [Using IAM to Manage Access to Amazon SWF Workflows]in the Amazon SWF Developer Guide.
//
// [Using IAM to Manage Access to Amazon SWF Workflows]: https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html
func swf_ListDomains(cfg aws.Config, client *swf.Client) {
	input := &swf.ListDomainsInput{
		// RegistrationStatus: types.RegistrationStatus, // Required
	}

	if len(_swfRegistrationStatus) > 0 {
		if err := assignInputField(input, "RegistrationStatus", _swfRegistrationStatus); err != nil {
			log.Errorf("invalid --registration-status: %s", err.Error())
			return
		}
	}
	if len(_swfMaximumPageSize) > 0 {
		if err := assignInputField(input, "MaximumPageSize", _swfMaximumPageSize); err != nil {
			log.Errorf("invalid --maximum-page-size: %s", err.Error())
			return
		}
	}
	if len(_swfNextPageToken) > 0 {
		input.NextPageToken = aws.String(_swfNextPageToken)
	}
	if len(_swfReverseOrder) > 0 {
		if err := assignInputField(input, "ReverseOrder", _swfReverseOrder); err != nil {
			log.Errorf("invalid --reverse-order: %s", err.Error())
			return
		}
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

	var results []*swf.ListDomainsOutput
	p := swf.NewListDomainsPaginator(client, input)
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

// Returns a list of open workflow executions in the specified domain that meet
// the filtering criteria. The results may be split into multiple pages. To
// retrieve subsequent pages, make the call again using the nextPageToken returned
// by the initial call.
//
// This operation is eventually consistent. The results are best effort and may
// not exactly reflect recent updates and changes.
//
// # Access Control
//
// You can use IAM policies to control this action's access to Amazon SWF
// resources as follows:
//
// - Use a Resource element with the domain name to limit the action to only
// specified domains.
//
// - Use an Action element to allow or deny permission to call this action.
//
// - Constrain the following parameters by using a Condition element with the
// appropriate keys.
//
// - tagFilter.tag : String constraint. The key is swf:tagFilter.tag .
//
// - typeFilter.name : String constraint. The key is swf:typeFilter.name .
//
// - typeFilter.version : String constraint. The key is swf:typeFilter.version .
//
// If the caller doesn't have sufficient permissions to invoke the action, or the
// parameter values fall outside the specified constraints, the action fails. The
// associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED .
// For details and example IAM policies, see [Using IAM to Manage Access to Amazon SWF Workflows]in the Amazon SWF Developer Guide.
//
// [Using IAM to Manage Access to Amazon SWF Workflows]: https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html
func swf_ListOpenWorkflowExecutions(cfg aws.Config, client *swf.Client) {
	input := &swf.ListOpenWorkflowExecutionsInput{
		// Domain: *string, // Required
		// StartTimeFilter: *types.ExecutionTimeFilter, // Required
	}

	if len(_swfDomain) > 0 {
		input.Domain = aws.String(_swfDomain)
	}
	if len(_swfStartTimeFilter) > 0 {
		if err := assignInputField(input, "StartTimeFilter", _swfStartTimeFilter); err != nil {
			log.Errorf("invalid --start-time-filter: %s", err.Error())
			return
		}
	}
	if len(_swfExecutionFilter) > 0 {
		if err := assignInputField(input, "ExecutionFilter", _swfExecutionFilter); err != nil {
			log.Errorf("invalid --execution-filter: %s", err.Error())
			return
		}
	}
	if len(_swfMaximumPageSize) > 0 {
		if err := assignInputField(input, "MaximumPageSize", _swfMaximumPageSize); err != nil {
			log.Errorf("invalid --maximum-page-size: %s", err.Error())
			return
		}
	}
	if len(_swfNextPageToken) > 0 {
		input.NextPageToken = aws.String(_swfNextPageToken)
	}
	if len(_swfReverseOrder) > 0 {
		if err := assignInputField(input, "ReverseOrder", _swfReverseOrder); err != nil {
			log.Errorf("invalid --reverse-order: %s", err.Error())
			return
		}
	}
	if len(_swfTagFilter) > 0 {
		if err := assignInputField(input, "TagFilter", _swfTagFilter); err != nil {
			log.Errorf("invalid --tag-filter: %s", err.Error())
			return
		}
	}
	if len(_swfTypeFilter) > 0 {
		if err := assignInputField(input, "TypeFilter", _swfTypeFilter); err != nil {
			log.Errorf("invalid --type-filter: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListOpenWorkflowExecutions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*swf.ListOpenWorkflowExecutionsOutput
	p := swf.NewListOpenWorkflowExecutionsPaginator(client, input)
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

// List tags for a given domain.
func swf_ListTagsForResource(cfg aws.Config, client *swf.Client) {
	input := &swf.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_swfResourceArn) > 0 {
		input.ResourceArn = aws.String(_swfResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about workflow types in the specified domain. The results
// may be split into multiple pages that can be retrieved by making the call
// repeatedly.
//
// # Access Control
//
// You can use IAM policies to control this action's access to Amazon SWF
// resources as follows:
//
// - Use a Resource element with the domain name to limit the action to only
// specified domains.
//
// - Use an Action element to allow or deny permission to call this action.
//
// - You cannot use an IAM policy to constrain this action's parameters.
//
// If the caller doesn't have sufficient permissions to invoke the action, or the
// parameter values fall outside the specified constraints, the action fails. The
// associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED .
// For details and example IAM policies, see [Using IAM to Manage Access to Amazon SWF Workflows]in the Amazon SWF Developer Guide.
//
// [Using IAM to Manage Access to Amazon SWF Workflows]: https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html
func swf_ListWorkflowTypes(cfg aws.Config, client *swf.Client) {
	input := &swf.ListWorkflowTypesInput{
		// Domain: *string, // Required
		// RegistrationStatus: types.RegistrationStatus, // Required
	}

	if len(_swfDomain) > 0 {
		input.Domain = aws.String(_swfDomain)
	}
	if len(_swfRegistrationStatus) > 0 {
		if err := assignInputField(input, "RegistrationStatus", _swfRegistrationStatus); err != nil {
			log.Errorf("invalid --registration-status: %s", err.Error())
			return
		}
	}
	if len(_swfMaximumPageSize) > 0 {
		if err := assignInputField(input, "MaximumPageSize", _swfMaximumPageSize); err != nil {
			log.Errorf("invalid --maximum-page-size: %s", err.Error())
			return
		}
	}
	if len(_swfName) > 0 {
		input.Name = aws.String(_swfName)
	}
	if len(_swfNextPageToken) > 0 {
		input.NextPageToken = aws.String(_swfNextPageToken)
	}
	if len(_swfReverseOrder) > 0 {
		if err := assignInputField(input, "ReverseOrder", _swfReverseOrder); err != nil {
			log.Errorf("invalid --reverse-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListWorkflowTypes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*swf.ListWorkflowTypesOutput
	p := swf.NewListWorkflowTypesPaginator(client, input)
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

// Used by workers to get an ActivityTask from the specified activity taskList . This initiates
// a long poll, where the service holds the HTTP connection open and responds as
// soon as a task becomes available. The maximum time the service holds on to the
// request before responding is 60 seconds. If no task is available within 60
// seconds, the poll returns an empty result. An empty result, in this context,
// means that an ActivityTask is returned, but that the value of taskToken is an
// empty string. If a task is returned, the worker should use its type to identify
// and process it correctly.
//
// Workers should set their client side socket timeout to at least 70 seconds (10
// seconds higher than the maximum time service may hold the poll request).
//
// # Access Control
//
// You can use IAM policies to control this action's access to Amazon SWF
// resources as follows:
//
// - Use a Resource element with the domain name to limit the action to only
// specified domains.
//
// - Use an Action element to allow or deny permission to call this action.
//
// - Constrain the taskList.name parameter by using a Condition element with the
// swf:taskList.name key to allow the action to access only certain task lists.
//
// If the caller doesn't have sufficient permissions to invoke the action, or the
// parameter values fall outside the specified constraints, the action fails. The
// associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED .
// For details and example IAM policies, see [Using IAM to Manage Access to Amazon SWF Workflows]in the Amazon SWF Developer Guide.
//
// [Using IAM to Manage Access to Amazon SWF Workflows]: https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html
func swf_PollForActivityTask(cfg aws.Config, client *swf.Client) {
	input := &swf.PollForActivityTaskInput{
		// Domain: *string, // Required
		// TaskList: *types.TaskList, // Required
	}

	if len(_swfDomain) > 0 {
		input.Domain = aws.String(_swfDomain)
	}
	if len(_swfTaskList) > 0 {
		if err := assignInputField(input, "TaskList", _swfTaskList); err != nil {
			log.Errorf("invalid --task-list: %s", err.Error())
			return
		}
	}
	if len(_swfIdentity) > 0 {
		input.Identity = aws.String(_swfIdentity)
	}

	if resp, err := client.PollForActivityTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Used by deciders to get a DecisionTask from the specified decision taskList . A decision
// task may be returned for any open workflow execution that is using the specified
// task list. The task includes a paginated view of the history of the workflow
// execution. The decider should use the workflow type and the history to determine
// how to properly handle the task.
//
// This action initiates a long poll, where the service holds the HTTP connection
// open and responds as soon a task becomes available. If no decision task is
// available in the specified task list before the timeout of 60 seconds expires,
// an empty result is returned. An empty result, in this context, means that a
// DecisionTask is returned, but that the value of taskToken is an empty string.
//
// Deciders should set their client side socket timeout to at least 70 seconds (10
// seconds higher than the timeout).
//
// Because the number of workflow history events for a single workflow execution
// might be very large, the result returned might be split up across a number of
// pages. To retrieve subsequent pages, make additional calls to
// PollForDecisionTask using the nextPageToken returned by the initial call. Note
// that you do not call GetWorkflowExecutionHistory with this nextPageToken .
// Instead, call PollForDecisionTask again.
//
// # Access Control
//
// You can use IAM policies to control this action's access to Amazon SWF
// resources as follows:
//
// - Use a Resource element with the domain name to limit the action to only
// specified domains.
//
// - Use an Action element to allow or deny permission to call this action.
//
// - Constrain the taskList.name parameter by using a Condition element with the
// swf:taskList.name key to allow the action to access only certain task lists.
//
// If the caller doesn't have sufficient permissions to invoke the action, or the
// parameter values fall outside the specified constraints, the action fails. The
// associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED .
// For details and example IAM policies, see [Using IAM to Manage Access to Amazon SWF Workflows]in the Amazon SWF Developer Guide.
//
// [Using IAM to Manage Access to Amazon SWF Workflows]: https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html
func swf_PollForDecisionTask(cfg aws.Config, client *swf.Client) {
	input := &swf.PollForDecisionTaskInput{
		// Domain: *string, // Required
		// TaskList: *types.TaskList, // Required
	}

	if len(_swfDomain) > 0 {
		input.Domain = aws.String(_swfDomain)
	}
	if len(_swfTaskList) > 0 {
		if err := assignInputField(input, "TaskList", _swfTaskList); err != nil {
			log.Errorf("invalid --task-list: %s", err.Error())
			return
		}
	}
	if len(_swfIdentity) > 0 {
		input.Identity = aws.String(_swfIdentity)
	}
	if len(_swfMaximumPageSize) > 0 {
		if err := assignInputField(input, "MaximumPageSize", _swfMaximumPageSize); err != nil {
			log.Errorf("invalid --maximum-page-size: %s", err.Error())
			return
		}
	}
	if len(_swfNextPageToken) > 0 {
		input.NextPageToken = aws.String(_swfNextPageToken)
	}
	if len(_swfReverseOrder) > 0 {
		if err := assignInputField(input, "ReverseOrder", _swfReverseOrder); err != nil {
			log.Errorf("invalid --reverse-order: %s", err.Error())
			return
		}
	}
	if len(_swfStartAtPreviousStartedEvent) > 0 {
		if err := assignInputField(input, "StartAtPreviousStartedEvent", _swfStartAtPreviousStartedEvent); err != nil {
			log.Errorf("invalid --start-at-previous-started-event: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.PollForDecisionTask(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*swf.PollForDecisionTaskOutput
	p := swf.NewPollForDecisionTaskPaginator(client, input)
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

// Used by activity workers to report to the service that the ActivityTask represented by the
// specified taskToken is still making progress. The worker can also specify
// details of the progress, for example percent complete, using the details
// parameter. This action can also be used by the worker as a mechanism to check if
// cancellation is being requested for the activity task. If a cancellation is
// being attempted for the specified task, then the boolean cancelRequested flag
// returned by the service is set to true .
//
// This action resets the taskHeartbeatTimeout clock. The taskHeartbeatTimeout is
// specified in RegisterActivityType.
//
// This action doesn't in itself create an event in the workflow execution
// history. However, if the task times out, the workflow execution history contains
// a ActivityTaskTimedOut event that contains the information from the last
// heartbeat generated by the activity worker.
//
// The taskStartToCloseTimeout of an activity type is the maximum duration of an
// activity task, regardless of the number of RecordActivityTaskHeartbeatrequests received. The
// taskStartToCloseTimeout is also specified in RegisterActivityType.
//
// This operation is only useful for long-lived activities to report liveliness of
// the task and to determine if a cancellation is being attempted.
//
// If the cancelRequested flag returns true , a cancellation is being attempted. If
// the worker can cancel the activity, it should respond with RespondActivityTaskCanceled. Otherwise, it
// should ignore the cancellation request.
//
// # Access Control
//
// You can use IAM policies to control this action's access to Amazon SWF
// resources as follows:
//
// - Use a Resource element with the domain name to limit the action to only
// specified domains.
//
// - Use an Action element to allow or deny permission to call this action.
//
// - You cannot use an IAM policy to constrain this action's parameters.
//
// If the caller doesn't have sufficient permissions to invoke the action, or the
// parameter values fall outside the specified constraints, the action fails. The
// associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED .
// For details and example IAM policies, see [Using IAM to Manage Access to Amazon SWF Workflows]in the Amazon SWF Developer Guide.
//
// [Using IAM to Manage Access to Amazon SWF Workflows]: https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html
func swf_RecordActivityTaskHeartbeat(cfg aws.Config, client *swf.Client) {
	input := &swf.RecordActivityTaskHeartbeatInput{
		// TaskToken: *string, // Required
	}

	if len(_swfTaskToken) > 0 {
		input.TaskToken = aws.String(_swfTaskToken)
	}
	if len(_swfDetails) > 0 {
		input.Details = aws.String(_swfDetails)
	}

	if resp, err := client.RecordActivityTaskHeartbeat(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Registers a new activity type along with its configuration settings in the
// specified domain.
//
// A TypeAlreadyExists fault is returned if the type already exists in the domain.
// You cannot change any configuration settings of the type after its registration,
// and it must be registered as a new version.
//
// # Access Control
//
// You can use IAM policies to control this action's access to Amazon SWF
// resources as follows:
//
// - Use a Resource element with the domain name to limit the action to only
// specified domains.
//
// - Use an Action element to allow or deny permission to call this action.
//
// - Constrain the following parameters by using a Condition element with the
// appropriate keys.
//
// - defaultTaskList.name : String constraint. The key is
// swf:defaultTaskList.name .
//
// - name : String constraint. The key is swf:name .
//
// - version : String constraint. The key is swf:version .
//
// If the caller doesn't have sufficient permissions to invoke the action, or the
// parameter values fall outside the specified constraints, the action fails. The
// associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED .
// For details and example IAM policies, see [Using IAM to Manage Access to Amazon SWF Workflows]in the Amazon SWF Developer Guide.
//
// [Using IAM to Manage Access to Amazon SWF Workflows]: https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html
func swf_RegisterActivityType(cfg aws.Config, client *swf.Client) {
	input := &swf.RegisterActivityTypeInput{
		// Domain: *string, // Required
		// Name: *string, // Required
		// Version: *string, // Required
	}

	if len(_swfDomain) > 0 {
		input.Domain = aws.String(_swfDomain)
	}
	if len(_swfName) > 0 {
		input.Name = aws.String(_swfName)
	}
	if len(_swfVersion) > 0 {
		input.Version = aws.String(_swfVersion)
	}
	if len(_swfDefaultTaskHeartbeatTimeout) > 0 {
		input.DefaultTaskHeartbeatTimeout = aws.String(_swfDefaultTaskHeartbeatTimeout)
	}
	if len(_swfDefaultTaskList) > 0 {
		if err := assignInputField(input, "DefaultTaskList", _swfDefaultTaskList); err != nil {
			log.Errorf("invalid --default-task-list: %s", err.Error())
			return
		}
	}
	if len(_swfDefaultTaskPriority) > 0 {
		input.DefaultTaskPriority = aws.String(_swfDefaultTaskPriority)
	}
	if len(_swfDefaultTaskScheduleToCloseTimeout) > 0 {
		input.DefaultTaskScheduleToCloseTimeout = aws.String(_swfDefaultTaskScheduleToCloseTimeout)
	}
	if len(_swfDefaultTaskScheduleToStartTimeout) > 0 {
		input.DefaultTaskScheduleToStartTimeout = aws.String(_swfDefaultTaskScheduleToStartTimeout)
	}
	if len(_swfDefaultTaskStartToCloseTimeout) > 0 {
		input.DefaultTaskStartToCloseTimeout = aws.String(_swfDefaultTaskStartToCloseTimeout)
	}
	if len(_swfDescription) > 0 {
		input.Description = aws.String(_swfDescription)
	}

	if resp, err := client.RegisterActivityType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Registers a new domain.
// # Access Control
//
// You can use IAM policies to control this action's access to Amazon SWF
// resources as follows:
//
// - You cannot use an IAM policy to control domain access for this action. The
// name of the domain being registered is available as the resource of this action.
//
// - Use an Action element to allow or deny permission to call this action.
//
// - You cannot use an IAM policy to constrain this action's parameters.
//
// If the caller doesn't have sufficient permissions to invoke the action, or the
// parameter values fall outside the specified constraints, the action fails. The
// associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED .
// For details and example IAM policies, see [Using IAM to Manage Access to Amazon SWF Workflows]in the Amazon SWF Developer Guide.
//
// [Using IAM to Manage Access to Amazon SWF Workflows]: https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html
func swf_RegisterDomain(cfg aws.Config, client *swf.Client) {
	input := &swf.RegisterDomainInput{
		// Name: *string, // Required
		// WorkflowExecutionRetentionPeriodInDays: *string, // Required
	}

	if len(_swfName) > 0 {
		input.Name = aws.String(_swfName)
	}
	if len(_swfWorkflowExecutionRetentionPeriodInDays) > 0 {
		input.WorkflowExecutionRetentionPeriodInDays = aws.String(_swfWorkflowExecutionRetentionPeriodInDays)
	}
	if len(_swfDescription) > 0 {
		input.Description = aws.String(_swfDescription)
	}
	if len(_swfTags) > 0 {
		if err := assignInputField(input, "Tags", _swfTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.RegisterDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Registers a new workflow type and its configuration settings in the specified
// domain.
//
// The retention period for the workflow history is set by the RegisterDomain action.
//
// If the type already exists, then a TypeAlreadyExists fault is returned. You
// cannot change the configuration settings of a workflow type once it is
// registered and it must be registered as a new version.
//
// # Access Control
//
// You can use IAM policies to control this action's access to Amazon SWF
// resources as follows:
//
// - Use a Resource element with the domain name to limit the action to only
// specified domains.
//
// - Use an Action element to allow or deny permission to call this action.
//
// - Constrain the following parameters by using a Condition element with the
// appropriate keys.
//
// - defaultTaskList.name : String constraint. The key is
// swf:defaultTaskList.name .
//
// - name : String constraint. The key is swf:name .
//
// - version : String constraint. The key is swf:version .
//
// If the caller doesn't have sufficient permissions to invoke the action, or the
// parameter values fall outside the specified constraints, the action fails. The
// associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED .
// For details and example IAM policies, see [Using IAM to Manage Access to Amazon SWF Workflows]in the Amazon SWF Developer Guide.
//
// [Using IAM to Manage Access to Amazon SWF Workflows]: https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html
func swf_RegisterWorkflowType(cfg aws.Config, client *swf.Client) {
	input := &swf.RegisterWorkflowTypeInput{
		// Domain: *string, // Required
		// Name: *string, // Required
		// Version: *string, // Required
	}

	if len(_swfDomain) > 0 {
		input.Domain = aws.String(_swfDomain)
	}
	if len(_swfName) > 0 {
		input.Name = aws.String(_swfName)
	}
	if len(_swfVersion) > 0 {
		input.Version = aws.String(_swfVersion)
	}
	if len(_swfDefaultChildPolicy) > 0 {
		if err := assignInputField(input, "DefaultChildPolicy", _swfDefaultChildPolicy); err != nil {
			log.Errorf("invalid --default-child-policy: %s", err.Error())
			return
		}
	}
	if len(_swfDefaultExecutionStartToCloseTimeout) > 0 {
		input.DefaultExecutionStartToCloseTimeout = aws.String(_swfDefaultExecutionStartToCloseTimeout)
	}
	if len(_swfDefaultLambdaRole) > 0 {
		input.DefaultLambdaRole = aws.String(_swfDefaultLambdaRole)
	}
	if len(_swfDefaultTaskList) > 0 {
		if err := assignInputField(input, "DefaultTaskList", _swfDefaultTaskList); err != nil {
			log.Errorf("invalid --default-task-list: %s", err.Error())
			return
		}
	}
	if len(_swfDefaultTaskPriority) > 0 {
		input.DefaultTaskPriority = aws.String(_swfDefaultTaskPriority)
	}
	if len(_swfDefaultTaskStartToCloseTimeout) > 0 {
		input.DefaultTaskStartToCloseTimeout = aws.String(_swfDefaultTaskStartToCloseTimeout)
	}
	if len(_swfDescription) > 0 {
		input.Description = aws.String(_swfDescription)
	}

	if resp, err := client.RegisterWorkflowType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Records a WorkflowExecutionCancelRequested event in the currently running
// workflow execution identified by the given domain, workflowId, and runId. This
// logically requests the cancellation of the workflow execution as a whole. It is
// up to the decider to take appropriate actions when it receives an execution
// history with this event.
//
// If the runId isn't specified, the WorkflowExecutionCancelRequested event is
// recorded in the history of the current open workflow execution with the
// specified workflowId in the domain.
//
// Because this action allows the workflow to properly clean up and gracefully
// close, it should be used instead of TerminateWorkflowExecutionwhen possible.
//
// # Access Control
//
// You can use IAM policies to control this action's access to Amazon SWF
// resources as follows:
//
// - Use a Resource element with the domain name to limit the action to only
// specified domains.
//
// - Use an Action element to allow or deny permission to call this action.
//
// - You cannot use an IAM policy to constrain this action's parameters.
//
// If the caller doesn't have sufficient permissions to invoke the action, or the
// parameter values fall outside the specified constraints, the action fails. The
// associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED .
// For details and example IAM policies, see [Using IAM to Manage Access to Amazon SWF Workflows]in the Amazon SWF Developer Guide.
//
// [Using IAM to Manage Access to Amazon SWF Workflows]: https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html
func swf_RequestCancelWorkflowExecution(cfg aws.Config, client *swf.Client) {
	input := &swf.RequestCancelWorkflowExecutionInput{
		// Domain: *string, // Required
		// WorkflowId: *string, // Required
	}

	if len(_swfDomain) > 0 {
		input.Domain = aws.String(_swfDomain)
	}
	if len(_swfWorkflowId) > 0 {
		input.WorkflowId = aws.String(_swfWorkflowId)
	}
	if len(_swfRunId) > 0 {
		input.RunId = aws.String(_swfRunId)
	}

	if resp, err := client.RequestCancelWorkflowExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Used by workers to tell the service that the ActivityTask identified by the taskToken was
// successfully canceled. Additional details can be provided using the details
// argument.
//
// These details (if provided) appear in the ActivityTaskCanceled event added to
// the workflow history.
//
// Only use this operation if the canceled flag of a RecordActivityTaskHeartbeat request returns true and if
// the activity can be safely undone or abandoned.
//
// A task is considered open from the time that it is scheduled until it is
// closed. Therefore a task is reported as open while a worker is processing it. A
// task is closed after it has been specified in a call to RespondActivityTaskCompleted,
// RespondActivityTaskCanceled, RespondActivityTaskFailed, or the task has [timed out].
//
// # Access Control
//
// You can use IAM policies to control this action's access to Amazon SWF
// resources as follows:
//
// - Use a Resource element with the domain name to limit the action to only
// specified domains.
//
// - Use an Action element to allow or deny permission to call this action.
//
// - You cannot use an IAM policy to constrain this action's parameters.
//
// If the caller doesn't have sufficient permissions to invoke the action, or the
// parameter values fall outside the specified constraints, the action fails. The
// associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED .
// For details and example IAM policies, see [Using IAM to Manage Access to Amazon SWF Workflows]in the Amazon SWF Developer Guide.
//
// [timed out]: https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dg-basic.html#swf-dev-timeout-types
// [Using IAM to Manage Access to Amazon SWF Workflows]: https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html
func swf_RespondActivityTaskCanceled(cfg aws.Config, client *swf.Client) {
	input := &swf.RespondActivityTaskCanceledInput{
		// TaskToken: *string, // Required
	}

	if len(_swfTaskToken) > 0 {
		input.TaskToken = aws.String(_swfTaskToken)
	}
	if len(_swfDetails) > 0 {
		input.Details = aws.String(_swfDetails)
	}

	if resp, err := client.RespondActivityTaskCanceled(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Used by workers to tell the service that the ActivityTask identified by the taskToken
// completed successfully with a result (if provided). The result appears in the
// ActivityTaskCompleted event in the workflow history.
//
// If the requested task doesn't complete successfully, use RespondActivityTaskFailed instead. If the
// worker finds that the task is canceled through the canceled flag returned by RecordActivityTaskHeartbeat,
// it should cancel the task, clean up and then call RespondActivityTaskCanceled.
//
// A task is considered open from the time that it is scheduled until it is
// closed. Therefore a task is reported as open while a worker is processing it. A
// task is closed after it has been specified in a call to
// RespondActivityTaskCompleted, RespondActivityTaskCanceled, RespondActivityTaskFailed, or the task has [timed out].
//
// # Access Control
//
// You can use IAM policies to control this action's access to Amazon SWF
// resources as follows:
//
// - Use a Resource element with the domain name to limit the action to only
// specified domains.
//
// - Use an Action element to allow or deny permission to call this action.
//
// - You cannot use an IAM policy to constrain this action's parameters.
//
// If the caller doesn't have sufficient permissions to invoke the action, or the
// parameter values fall outside the specified constraints, the action fails. The
// associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED .
// For details and example IAM policies, see [Using IAM to Manage Access to Amazon SWF Workflows]in the Amazon SWF Developer Guide.
//
// [timed out]: https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dg-basic.html#swf-dev-timeout-types
// [Using IAM to Manage Access to Amazon SWF Workflows]: https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html
func swf_RespondActivityTaskCompleted(cfg aws.Config, client *swf.Client) {
	input := &swf.RespondActivityTaskCompletedInput{
		// TaskToken: *string, // Required
	}

	if len(_swfTaskToken) > 0 {
		input.TaskToken = aws.String(_swfTaskToken)
	}
	if len(_swfResult) > 0 {
		input.Result = aws.String(_swfResult)
	}

	if resp, err := client.RespondActivityTaskCompleted(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Used by workers to tell the service that the ActivityTask identified by the taskToken has
// failed with reason (if specified). The reason and details appear in the
// ActivityTaskFailed event added to the workflow history.
//
// A task is considered open from the time that it is scheduled until it is
// closed. Therefore a task is reported as open while a worker is processing it. A
// task is closed after it has been specified in a call to RespondActivityTaskCompleted, RespondActivityTaskCanceled,
// RespondActivityTaskFailed, or the task has [timed out].
//
// # Access Control
//
// You can use IAM policies to control this action's access to Amazon SWF
// resources as follows:
//
// - Use a Resource element with the domain name to limit the action to only
// specified domains.
//
// - Use an Action element to allow or deny permission to call this action.
//
// - You cannot use an IAM policy to constrain this action's parameters.
//
// If the caller doesn't have sufficient permissions to invoke the action, or the
// parameter values fall outside the specified constraints, the action fails. The
// associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED .
// For details and example IAM policies, see [Using IAM to Manage Access to Amazon SWF Workflows]in the Amazon SWF Developer Guide.
//
// [timed out]: https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dg-basic.html#swf-dev-timeout-types
// [Using IAM to Manage Access to Amazon SWF Workflows]: https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html
func swf_RespondActivityTaskFailed(cfg aws.Config, client *swf.Client) {
	input := &swf.RespondActivityTaskFailedInput{
		// TaskToken: *string, // Required
	}

	if len(_swfTaskToken) > 0 {
		input.TaskToken = aws.String(_swfTaskToken)
	}
	if len(_swfDetails) > 0 {
		input.Details = aws.String(_swfDetails)
	}
	if len(_swfReason) > 0 {
		input.Reason = aws.String(_swfReason)
	}

	if resp, err := client.RespondActivityTaskFailed(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Used by deciders to tell the service that the DecisionTask identified by the taskToken has
// successfully completed. The decisions argument specifies the list of decisions
// made while processing the task.
//
// A DecisionTaskCompleted event is added to the workflow history. The
// executionContext specified is attached to the event in the workflow execution
// history.
//
// # Access Control
//
// If an IAM policy grants permission to use RespondDecisionTaskCompleted , it can
// express permissions for the list of decisions in the decisions parameter. Each
// of the decisions has one or more parameters, much like a regular API call. To
// allow for policies to be as readable as possible, you can express permissions on
// decisions as if they were actual API calls, including applying conditions to
// some parameters. For more information, see [Using IAM to Manage Access to Amazon SWF Workflows]in the Amazon SWF Developer Guide.
//
// [Using IAM to Manage Access to Amazon SWF Workflows]: https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html
func swf_RespondDecisionTaskCompleted(cfg aws.Config, client *swf.Client) {
	input := &swf.RespondDecisionTaskCompletedInput{
		// TaskToken: *string, // Required
	}

	if len(_swfTaskToken) > 0 {
		input.TaskToken = aws.String(_swfTaskToken)
	}
	if len(_swfDecisions) > 0 {
		if err := assignInputField(input, "Decisions", _swfDecisions); err != nil {
			log.Errorf("invalid --decisions: %s", err.Error())
			return
		}
	}
	if len(_swfExecutionContext) > 0 {
		input.ExecutionContext = aws.String(_swfExecutionContext)
	}
	if len(_swfTaskList) > 0 {
		if err := assignInputField(input, "TaskList", _swfTaskList); err != nil {
			log.Errorf("invalid --task-list: %s", err.Error())
			return
		}
	}
	if len(_swfTaskListScheduleToStartTimeout) > 0 {
		input.TaskListScheduleToStartTimeout = aws.String(_swfTaskListScheduleToStartTimeout)
	}

	if resp, err := client.RespondDecisionTaskCompleted(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Records a WorkflowExecutionSignaled event in the workflow execution history and
// creates a decision task for the workflow execution identified by the given
// domain, workflowId and runId. The event is recorded with the specified user
// defined signalName and input (if provided).
//
// If a runId isn't specified, then the WorkflowExecutionSignaled event is
// recorded in the history of the current open workflow with the matching
// workflowId in the domain.
//
// If the specified workflow execution isn't open, this method fails with
// UnknownResource .
//
// # Access Control
//
// You can use IAM policies to control this action's access to Amazon SWF
// resources as follows:
//
// - Use a Resource element with the domain name to limit the action to only
// specified domains.
//
// - Use an Action element to allow or deny permission to call this action.
//
// - You cannot use an IAM policy to constrain this action's parameters.
//
// If the caller doesn't have sufficient permissions to invoke the action, or the
// parameter values fall outside the specified constraints, the action fails. The
// associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED .
// For details and example IAM policies, see [Using IAM to Manage Access to Amazon SWF Workflows]in the Amazon SWF Developer Guide.
//
// [Using IAM to Manage Access to Amazon SWF Workflows]: https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html
func swf_SignalWorkflowExecution(cfg aws.Config, client *swf.Client) {
	input := &swf.SignalWorkflowExecutionInput{
		// Domain: *string, // Required
		// SignalName: *string, // Required
		// WorkflowId: *string, // Required
	}

	if len(_swfDomain) > 0 {
		input.Domain = aws.String(_swfDomain)
	}
	if len(_swfSignalName) > 0 {
		input.SignalName = aws.String(_swfSignalName)
	}
	if len(_swfWorkflowId) > 0 {
		input.WorkflowId = aws.String(_swfWorkflowId)
	}
	if len(_swfInput) > 0 {
		input.Input = aws.String(_swfInput)
	}
	if len(_swfRunId) > 0 {
		input.RunId = aws.String(_swfRunId)
	}

	if resp, err := client.SignalWorkflowExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts an execution of the workflow type in the specified domain using the
// provided workflowId and input data.
//
// This action returns the newly started workflow execution.
//
// # Access Control
//
// You can use IAM policies to control this action's access to Amazon SWF
// resources as follows:
//
// - Use a Resource element with the domain name to limit the action to only
// specified domains.
//
// - Use an Action element to allow or deny permission to call this action.
//
// - Constrain the following parameters by using a Condition element with the
// appropriate keys.
//
// - tagList.member.0 : The key is swf:tagList.member.0 .
//
// - tagList.member.1 : The key is swf:tagList.member.1 .
//
// - tagList.member.2 : The key is swf:tagList.member.2 .
//
// - tagList.member.3 : The key is swf:tagList.member.3 .
//
// - tagList.member.4 : The key is swf:tagList.member.4 .
//
// - taskList : String constraint. The key is swf:taskList.name .
//
// - workflowType.name : String constraint. The key is swf:workflowType.name .
//
// - workflowType.version : String constraint. The key is
// swf:workflowType.version .
//
// If the caller doesn't have sufficient permissions to invoke the action, or the
// parameter values fall outside the specified constraints, the action fails. The
// associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED .
// For details and example IAM policies, see [Using IAM to Manage Access to Amazon SWF Workflows]in the Amazon SWF Developer Guide.
//
// [Using IAM to Manage Access to Amazon SWF Workflows]: https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html
func swf_StartWorkflowExecution(cfg aws.Config, client *swf.Client) {
	input := &swf.StartWorkflowExecutionInput{
		// Domain: *string, // Required
		// WorkflowId: *string, // Required
		// WorkflowType: *types.WorkflowType, // Required
	}

	if len(_swfDomain) > 0 {
		input.Domain = aws.String(_swfDomain)
	}
	if len(_swfWorkflowId) > 0 {
		input.WorkflowId = aws.String(_swfWorkflowId)
	}
	if len(_swfWorkflowType) > 0 {
		if err := assignInputField(input, "WorkflowType", _swfWorkflowType); err != nil {
			log.Errorf("invalid --workflow-type: %s", err.Error())
			return
		}
	}
	if len(_swfChildPolicy) > 0 {
		if err := assignInputField(input, "ChildPolicy", _swfChildPolicy); err != nil {
			log.Errorf("invalid --child-policy: %s", err.Error())
			return
		}
	}
	if len(_swfExecutionStartToCloseTimeout) > 0 {
		input.ExecutionStartToCloseTimeout = aws.String(_swfExecutionStartToCloseTimeout)
	}
	if len(_swfInput) > 0 {
		input.Input = aws.String(_swfInput)
	}
	if len(_swfLambdaRole) > 0 {
		input.LambdaRole = aws.String(_swfLambdaRole)
	}
	if len(_swfTagList) > 0 {
		input.TagList = append([]string(nil), _swfTagList...)
	}
	if len(_swfTaskList) > 0 {
		if err := assignInputField(input, "TaskList", _swfTaskList); err != nil {
			log.Errorf("invalid --task-list: %s", err.Error())
			return
		}
	}
	if len(_swfTaskPriority) > 0 {
		input.TaskPriority = aws.String(_swfTaskPriority)
	}
	if len(_swfTaskStartToCloseTimeout) > 0 {
		input.TaskStartToCloseTimeout = aws.String(_swfTaskStartToCloseTimeout)
	}

	if resp, err := client.StartWorkflowExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Add a tag to a Amazon SWF domain.
// Amazon SWF supports a maximum of 50 tags per resource.
func swf_TagResource(cfg aws.Config, client *swf.Client) {
	input := &swf.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: []types.ResourceTag, // Required
	}

	if len(_swfResourceArn) > 0 {
		input.ResourceArn = aws.String(_swfResourceArn)
	}
	if len(_swfTags) > 0 {
		if err := assignInputField(input, "Tags", _swfTags); err != nil {
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

// Records a WorkflowExecutionTerminated event and forces closure of the workflow
// execution identified by the given domain, runId, and workflowId. The child
// policy, registered with the workflow type or specified when starting this
// execution, is applied to any open child workflow executions of this workflow
// execution.
//
// If the identified workflow execution was in progress, it is terminated
// immediately.
//
// If a runId isn't specified, then the WorkflowExecutionTerminated event is
// recorded in the history of the current open workflow with the matching
// workflowId in the domain.
//
// You should consider using RequestCancelWorkflowExecution action instead because it allows the workflow to
// gracefully close while TerminateWorkflowExecutiondoesn't.
//
// # Access Control
//
// You can use IAM policies to control this action's access to Amazon SWF
// resources as follows:
//
// - Use a Resource element with the domain name to limit the action to only
// specified domains.
//
// - Use an Action element to allow or deny permission to call this action.
//
// - You cannot use an IAM policy to constrain this action's parameters.
//
// If the caller doesn't have sufficient permissions to invoke the action, or the
// parameter values fall outside the specified constraints, the action fails. The
// associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED .
// For details and example IAM policies, see [Using IAM to Manage Access to Amazon SWF Workflows]in the Amazon SWF Developer Guide.
//
// [Using IAM to Manage Access to Amazon SWF Workflows]: https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html
func swf_TerminateWorkflowExecution(cfg aws.Config, client *swf.Client) {
	input := &swf.TerminateWorkflowExecutionInput{
		// Domain: *string, // Required
		// WorkflowId: *string, // Required
	}

	if len(_swfDomain) > 0 {
		input.Domain = aws.String(_swfDomain)
	}
	if len(_swfWorkflowId) > 0 {
		input.WorkflowId = aws.String(_swfWorkflowId)
	}
	if len(_swfChildPolicy) > 0 {
		if err := assignInputField(input, "ChildPolicy", _swfChildPolicy); err != nil {
			log.Errorf("invalid --child-policy: %s", err.Error())
			return
		}
	}
	if len(_swfDetails) > 0 {
		input.Details = aws.String(_swfDetails)
	}
	if len(_swfReason) > 0 {
		input.Reason = aws.String(_swfReason)
	}
	if len(_swfRunId) > 0 {
		input.RunId = aws.String(_swfRunId)
	}

	if resp, err := client.TerminateWorkflowExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Undeprecates a previously deprecated activity type. After an activity type has
// been undeprecated, you can create new tasks of that activity type.
//
// This operation is eventually consistent. The results are best effort and may
// not exactly reflect recent updates and changes.
//
// # Access Control
//
// You can use IAM policies to control this action's access to Amazon SWF
// resources as follows:
//
// - Use a Resource element with the domain name to limit the action to only
// specified domains.
//
// - Use an Action element to allow or deny permission to call this action.
//
// - Constrain the following parameters by using a Condition element with the
// appropriate keys.
//
// - activityType.name : String constraint. The key is swf:activityType.name .
//
// - activityType.version : String constraint. The key is
// swf:activityType.version .
//
// If the caller doesn't have sufficient permissions to invoke the action, or the
// parameter values fall outside the specified constraints, the action fails. The
// associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED .
// For details and example IAM policies, see [Using IAM to Manage Access to Amazon SWF Workflows]in the Amazon SWF Developer Guide.
//
// [Using IAM to Manage Access to Amazon SWF Workflows]: https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html
func swf_UndeprecateActivityType(cfg aws.Config, client *swf.Client) {
	input := &swf.UndeprecateActivityTypeInput{
		// ActivityType: *types.ActivityType, // Required
		// Domain: *string, // Required
	}

	if len(_swfActivityType) > 0 {
		if err := assignInputField(input, "ActivityType", _swfActivityType); err != nil {
			log.Errorf("invalid --activity-type: %s", err.Error())
			return
		}
	}
	if len(_swfDomain) > 0 {
		input.Domain = aws.String(_swfDomain)
	}

	if resp, err := client.UndeprecateActivityType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Undeprecates a previously deprecated domain. After a domain has been
// undeprecated it can be used to create new workflow executions or register new
// types.
//
// This operation is eventually consistent. The results are best effort and may
// not exactly reflect recent updates and changes.
//
// # Access Control
//
// You can use IAM policies to control this action's access to Amazon SWF
// resources as follows:
//
// - Use a Resource element with the domain name to limit the action to only
// specified domains.
//
// - Use an Action element to allow or deny permission to call this action.
//
// - You cannot use an IAM policy to constrain this action's parameters.
//
// If the caller doesn't have sufficient permissions to invoke the action, or the
// parameter values fall outside the specified constraints, the action fails. The
// associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED .
// For details and example IAM policies, see [Using IAM to Manage Access to Amazon SWF Workflows]in the Amazon SWF Developer Guide.
//
// [Using IAM to Manage Access to Amazon SWF Workflows]: https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html
func swf_UndeprecateDomain(cfg aws.Config, client *swf.Client) {
	input := &swf.UndeprecateDomainInput{
		// Name: *string, // Required
	}

	if len(_swfName) > 0 {
		input.Name = aws.String(_swfName)
	}

	if resp, err := client.UndeprecateDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Undeprecates a previously deprecated workflow type. After a workflow type has
// been undeprecated, you can create new executions of that type.
//
// This operation is eventually consistent. The results are best effort and may
// not exactly reflect recent updates and changes.
//
// # Access Control
//
// You can use IAM policies to control this action's access to Amazon SWF
// resources as follows:
//
// - Use a Resource element with the domain name to limit the action to only
// specified domains.
//
// - Use an Action element to allow or deny permission to call this action.
//
// - Constrain the following parameters by using a Condition element with the
// appropriate keys.
//
// - workflowType.name : String constraint. The key is swf:workflowType.name .
//
// - workflowType.version : String constraint. The key is
// swf:workflowType.version .
//
// If the caller doesn't have sufficient permissions to invoke the action, or the
// parameter values fall outside the specified constraints, the action fails. The
// associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED .
// For details and example IAM policies, see [Using IAM to Manage Access to Amazon SWF Workflows]in the Amazon SWF Developer Guide.
//
// [Using IAM to Manage Access to Amazon SWF Workflows]: https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html
func swf_UndeprecateWorkflowType(cfg aws.Config, client *swf.Client) {
	input := &swf.UndeprecateWorkflowTypeInput{
		// Domain: *string, // Required
		// WorkflowType: *types.WorkflowType, // Required
	}

	if len(_swfDomain) > 0 {
		input.Domain = aws.String(_swfDomain)
	}
	if len(_swfWorkflowType) > 0 {
		if err := assignInputField(input, "WorkflowType", _swfWorkflowType); err != nil {
			log.Errorf("invalid --workflow-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.UndeprecateWorkflowType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Remove a tag from a Amazon SWF domain.
func swf_UntagResource(cfg aws.Config, client *swf.Client) {
	input := &swf.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_swfResourceArn) > 0 {
		input.ResourceArn = aws.String(_swfResourceArn)
	}
	if len(_swfTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _swfTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_swfCmd)
	_swfCmd.Flags().SortFlags = false

	_swfCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_swfCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_swfCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_swfCmd.Flags().StringVarP(&_swfActivityType, "activity-type", "", "", "Activity Type")
	_swfCmd.Flags().StringVarP(&_swfChildPolicy, "child-policy", "", "", "Child Policy")
	_swfCmd.Flags().StringVarP(&_swfCloseStatusFilter, "close-status-filter", "", "", "Close Status Filter")
	_swfCmd.Flags().StringVarP(&_swfCloseTimeFilter, "close-time-filter", "", "", "Close Time Filter")
	_swfCmd.Flags().StringVarP(&_swfDecisions, "decisions", "", "", "Decisions")
	_swfCmd.Flags().StringVarP(&_swfDefaultChildPolicy, "default-child-policy", "", "", "Default Child Policy")
	_swfCmd.Flags().StringVarP(&_swfDefaultExecutionStartToCloseTimeout, "default-execution-start-to-close-timeout", "", "", "Default Execution Start To Close Timeout")
	_swfCmd.Flags().StringVarP(&_swfDefaultLambdaRole, "default-lambda-role", "", "", "Default Lambda Role")
	_swfCmd.Flags().StringVarP(&_swfDefaultTaskHeartbeatTimeout, "default-task-heartbeat-timeout", "", "", "Default Task Heartbeat Timeout")
	_swfCmd.Flags().StringVarP(&_swfDefaultTaskList, "default-task-list", "", "", "Default Task List")
	_swfCmd.Flags().StringVarP(&_swfDefaultTaskPriority, "default-task-priority", "", "", "Default Task Priority")
	_swfCmd.Flags().StringVarP(&_swfDefaultTaskScheduleToCloseTimeout, "default-task-schedule-to-close-timeout", "", "", "Default Task Schedule To Close Timeout")
	_swfCmd.Flags().StringVarP(&_swfDefaultTaskScheduleToStartTimeout, "default-task-schedule-to-start-timeout", "", "", "Default Task Schedule To Start Timeout")
	_swfCmd.Flags().StringVarP(&_swfDefaultTaskStartToCloseTimeout, "default-task-start-to-close-timeout", "", "", "Default Task Start To Close Timeout")
	_swfCmd.Flags().StringVarP(&_swfDescription, "description", "", "", "Description")
	_swfCmd.Flags().StringVarP(&_swfDetails, "details", "", "", "Details")
	_swfCmd.Flags().StringVarP(&_swfDomain, "domain", "", "", "Domain")
	_swfCmd.Flags().StringVarP(&_swfExecution, "execution", "", "", "Execution")
	_swfCmd.Flags().StringVarP(&_swfExecutionContext, "execution-context", "", "", "Execution Context")
	_swfCmd.Flags().StringVarP(&_swfExecutionFilter, "execution-filter", "", "", "Execution Filter")
	_swfCmd.Flags().StringVarP(&_swfExecutionStartToCloseTimeout, "execution-start-to-close-timeout", "", "", "Execution Start To Close Timeout")
	_swfCmd.Flags().StringVarP(&_swfIdentity, "identity", "", "", "Identity")
	_swfCmd.Flags().StringVarP(&_swfInput, "input", "", "", "Input")
	_swfCmd.Flags().StringVarP(&_swfLambdaRole, "lambda-role", "", "", "Lambda Role")
	_swfCmd.Flags().StringVarP(&_swfMaximumPageSize, "maximum-page-size", "", "", "Maximum Page Size")
	_swfCmd.Flags().StringVarP(&_swfName, "name", "", "", "Name")
	_swfCmd.Flags().StringVarP(&_swfNextPageToken, "next-page-token", "", "", "Next Page Token")
	_swfCmd.Flags().StringVarP(&_swfReason, "reason", "", "", "Reason")
	_swfCmd.Flags().StringVarP(&_swfRegistrationStatus, "registration-status", "", "", "Registration Status")
	_swfCmd.Flags().StringVarP(&_swfResourceArn, "resource-arn", "", "", "Resource ARN")
	_swfCmd.Flags().StringVarP(&_swfResult, "result", "", "", "Result")
	_swfCmd.Flags().StringVarP(&_swfReverseOrder, "reverse-order", "", "", "Reverse Order")
	_swfCmd.Flags().StringVarP(&_swfRunId, "run-id", "", "", "Run ID")
	_swfCmd.Flags().StringVarP(&_swfSignalName, "signal-name", "", "", "Signal Name")
	_swfCmd.Flags().StringVarP(&_swfStartAtPreviousStartedEvent, "start-at-previous-started-event", "", "", "Start At Previous Started Event")
	_swfCmd.Flags().StringVarP(&_swfStartTimeFilter, "start-time-filter", "", "", "Start Time Filter")
	_swfCmd.Flags().StringVarP(&_swfTagFilter, "tag-filter", "", "", "Tag Filter")
	_swfCmd.Flags().StringSliceVarP(&_swfTagKeys, "tag-keys", "", nil, "Tag Keys")
	_swfCmd.Flags().StringSliceVarP(&_swfTagList, "tag-list", "", nil, "Tag List")
	_swfCmd.Flags().StringVarP(&_swfTags, "tags", "", "", "Tags")
	_swfCmd.Flags().StringVarP(&_swfTaskList, "task-list", "", "", "Task List")
	_swfCmd.Flags().StringVarP(&_swfTaskListScheduleToStartTimeout, "task-list-schedule-to-start-timeout", "", "", "Task List Schedule To Start Timeout")
	_swfCmd.Flags().StringVarP(&_swfTaskPriority, "task-priority", "", "", "Task Priority")
	_swfCmd.Flags().StringVarP(&_swfTaskStartToCloseTimeout, "task-start-to-close-timeout", "", "", "Task Start To Close Timeout")
	_swfCmd.Flags().StringVarP(&_swfTaskToken, "task-token", "", "", "Task Token")
	_swfCmd.Flags().StringVarP(&_swfTypeFilter, "type-filter", "", "", "Type Filter")
	_swfCmd.Flags().StringVarP(&_swfVersion, "version", "", "", "Version")
	_swfCmd.Flags().StringVarP(&_swfWorkflowExecutionRetentionPeriodInDays, "workflow-execution-retention-period-in-days", "", "", "Workflow Execution Retention Period In Days")
	_swfCmd.Flags().StringVarP(&_swfWorkflowId, "workflow-id", "", "", "Workflow ID")
	_swfCmd.Flags().StringVarP(&_swfWorkflowType, "workflow-type", "", "", "Workflow Type")

	_swfCmd.Flags().BoolVarP(&_swfCountClosedWorkflowExecutions, "count-closed-workflow-executions", "", false, "Count Closed Workflow Executions")
	_swfCmd.Flags().BoolVarP(&_swfCountOpenWorkflowExecutions, "count-open-workflow-executions", "", false, "Count Open Workflow Executions")
	_swfCmd.Flags().BoolVarP(&_swfCountPendingActivityTasks, "count-pending-activity-tasks", "", false, "Count Pending Activity Tasks")
	_swfCmd.Flags().BoolVarP(&_swfCountPendingDecisionTasks, "count-pending-decision-tasks", "", false, "Count Pending Decision Tasks")
	_swfCmd.Flags().BoolVarP(&_swfDeleteActivityType, "delete-activity-type", "", false, "Delete Activity Type")
	_swfCmd.Flags().BoolVarP(&_swfDeleteWorkflowType, "delete-workflow-type", "", false, "Delete Workflow Type")
	_swfCmd.Flags().BoolVarP(&_swfDeprecateActivityType, "deprecate-activity-type", "", false, "Deprecate Activity Type")
	_swfCmd.Flags().BoolVarP(&_swfDeprecateDomain, "deprecate-domain", "", false, "Deprecate Domain")
	_swfCmd.Flags().BoolVarP(&_swfDeprecateWorkflowType, "deprecate-workflow-type", "", false, "Deprecate Workflow Type")
	_swfCmd.Flags().BoolVarP(&_swfDescribeActivityType, "describe-activity-type", "", false, "Describe Activity Type")
	_swfCmd.Flags().BoolVarP(&_swfDescribeDomain, "describe-domain", "", false, "Describe Domain")
	_swfCmd.Flags().BoolVarP(&_swfDescribeWorkflowExecution, "describe-workflow-execution", "", false, "Describe Workflow Execution")
	_swfCmd.Flags().BoolVarP(&_swfDescribeWorkflowType, "describe-workflow-type", "", false, "Describe Workflow Type")
	_swfCmd.Flags().BoolVarP(&_swfGetWorkflowExecutionHistory, "get-workflow-execution-history", "", false, "Get Workflow Execution History")
	_swfCmd.Flags().BoolVarP(&_swfListActivityTypes, "list-activity-types", "", false, "List Activity Types")
	_swfCmd.Flags().BoolVarP(&_swfListClosedWorkflowExecutions, "list-closed-workflow-executions", "", false, "List Closed Workflow Executions")
	_swfCmd.Flags().BoolVarP(&_swfListDomains, "list-domains", "", false, "List Domains")
	_swfCmd.Flags().BoolVarP(&_swfListOpenWorkflowExecutions, "list-open-workflow-executions", "", false, "List Open Workflow Executions")
	_swfCmd.Flags().BoolVarP(&_swfListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_swfCmd.Flags().BoolVarP(&_swfListWorkflowTypes, "list-workflow-types", "", false, "List Workflow Types")
	_swfCmd.Flags().BoolVarP(&_swfPollForActivityTask, "poll-for-activity-task", "", false, "Poll For Activity Task")
	_swfCmd.Flags().BoolVarP(&_swfPollForDecisionTask, "poll-for-decision-task", "", false, "Poll For Decision Task")
	_swfCmd.Flags().BoolVarP(&_swfRecordActivityTaskHeartbeat, "record-activity-task-heartbeat", "", false, "Record Activity Task Heartbeat")
	_swfCmd.Flags().BoolVarP(&_swfRegisterActivityType, "register-activity-type", "", false, "Register Activity Type")
	_swfCmd.Flags().BoolVarP(&_swfRegisterDomain, "register-domain", "", false, "Register Domain")
	_swfCmd.Flags().BoolVarP(&_swfRegisterWorkflowType, "register-workflow-type", "", false, "Register Workflow Type")
	_swfCmd.Flags().BoolVarP(&_swfRequestCancelWorkflowExecution, "request-cancel-workflow-execution", "", false, "Request Cancel Workflow Execution")
	_swfCmd.Flags().BoolVarP(&_swfRespondActivityTaskCanceled, "respond-activity-task-canceled", "", false, "Respond Activity Task Canceled")
	_swfCmd.Flags().BoolVarP(&_swfRespondActivityTaskCompleted, "respond-activity-task-completed", "", false, "Respond Activity Task Completed")
	_swfCmd.Flags().BoolVarP(&_swfRespondActivityTaskFailed, "respond-activity-task-failed", "", false, "Respond Activity Task Failed")
	_swfCmd.Flags().BoolVarP(&_swfRespondDecisionTaskCompleted, "respond-decision-task-completed", "", false, "Respond Decision Task Completed")
	_swfCmd.Flags().BoolVarP(&_swfSignalWorkflowExecution, "signal-workflow-execution", "", false, "Signal Workflow Execution")
	_swfCmd.Flags().BoolVarP(&_swfStartWorkflowExecution, "start-workflow-execution", "", false, "Start Workflow Execution")
	_swfCmd.Flags().BoolVarP(&_swfTagResource, "tag-resource", "", false, "Tag Resource")
	_swfCmd.Flags().BoolVarP(&_swfTerminateWorkflowExecution, "terminate-workflow-execution", "", false, "Terminate Workflow Execution")
	_swfCmd.Flags().BoolVarP(&_swfUndeprecateActivityType, "undeprecate-activity-type", "", false, "Undeprecate Activity Type")
	_swfCmd.Flags().BoolVarP(&_swfUndeprecateDomain, "undeprecate-domain", "", false, "Undeprecate Domain")
	_swfCmd.Flags().BoolVarP(&_swfUndeprecateWorkflowType, "undeprecate-workflow-type", "", false, "Undeprecate Workflow Type")
	_swfCmd.Flags().BoolVarP(&_swfUntagResource, "untag-resource", "", false, "Untag Resource")

}
