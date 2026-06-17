package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/arcregionswitch"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// arcregionswitchCmd represents the arcregionswitch command
var _arcregionswitchCmd = &cobra.Command{
	Use:   "arcregionswitch",
	Short: "AWS arcregionswitch CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := arcregionswitch.NewFromConfig(cfg)
		if _arcregionswitchApprovePlanExecutionStep {
			arcregionswitch_ApprovePlanExecutionStep(cfg, client)
			return
		}
		if _arcregionswitchCancelPlanExecution {
			arcregionswitch_CancelPlanExecution(cfg, client)
			return
		}
		if _arcregionswitchCreatePlan {
			arcregionswitch_CreatePlan(cfg, client)
			return
		}
		if _arcregionswitchDeletePlan {
			arcregionswitch_DeletePlan(cfg, client)
			return
		}
		if _arcregionswitchGetPlan {
			arcregionswitch_GetPlan(cfg, client)
			return
		}
		if _arcregionswitchGetPlanEvaluationStatus {
			arcregionswitch_GetPlanEvaluationStatus(cfg, client)
			return
		}
		if _arcregionswitchGetPlanExecution {
			arcregionswitch_GetPlanExecution(cfg, client)
			return
		}
		if _arcregionswitchGetPlanInRegion {
			arcregionswitch_GetPlanInRegion(cfg, client)
			return
		}
		if _arcregionswitchListPlanExecutionEvents {
			arcregionswitch_ListPlanExecutionEvents(cfg, client)
			return
		}
		if _arcregionswitchListPlanExecutions {
			arcregionswitch_ListPlanExecutions(cfg, client)
			return
		}
		if _arcregionswitchListPlans {
			arcregionswitch_ListPlans(cfg, client)
			return
		}
		if _arcregionswitchListPlansInRegion {
			arcregionswitch_ListPlansInRegion(cfg, client)
			return
		}
		if _arcregionswitchListRoute53HealthChecks {
			arcregionswitch_ListRoute53HealthChecks(cfg, client)
			return
		}
		if _arcregionswitchListRoute53HealthChecksInRegion {
			arcregionswitch_ListRoute53HealthChecksInRegion(cfg, client)
			return
		}
		if _arcregionswitchListTagsForResource {
			arcregionswitch_ListTagsForResource(cfg, client)
			return
		}
		if _arcregionswitchStartPlanExecution {
			arcregionswitch_StartPlanExecution(cfg, client)
			return
		}
		if _arcregionswitchTagResource {
			arcregionswitch_TagResource(cfg, client)
			return
		}
		if _arcregionswitchUntagResource {
			arcregionswitch_UntagResource(cfg, client)
			return
		}
		if _arcregionswitchUpdatePlan {
			arcregionswitch_UpdatePlan(cfg, client)
			return
		}
		if _arcregionswitchUpdatePlanExecution {
			arcregionswitch_UpdatePlanExecution(cfg, client)
			return
		}
		if _arcregionswitchUpdatePlanExecutionStep {
			arcregionswitch_UpdatePlanExecutionStep(cfg, client)
			return
		}

	},
}

var (
	_arcregionswitchApprovePlanExecutionStep        bool
	_arcregionswitchCancelPlanExecution             bool
	_arcregionswitchCreatePlan                      bool
	_arcregionswitchDeletePlan                      bool
	_arcregionswitchGetPlan                         bool
	_arcregionswitchGetPlanEvaluationStatus         bool
	_arcregionswitchGetPlanExecution                bool
	_arcregionswitchGetPlanInRegion                 bool
	_arcregionswitchListPlanExecutionEvents         bool
	_arcregionswitchListPlanExecutions              bool
	_arcregionswitchListPlans                       bool
	_arcregionswitchListPlansInRegion               bool
	_arcregionswitchListRoute53HealthChecks         bool
	_arcregionswitchListRoute53HealthChecksInRegion bool
	_arcregionswitchListTagsForResource             bool
	_arcregionswitchStartPlanExecution              bool
	_arcregionswitchTagResource                     bool
	_arcregionswitchUntagResource                   bool
	_arcregionswitchUpdatePlan                      bool
	_arcregionswitchUpdatePlanExecution             bool
	_arcregionswitchUpdatePlanExecutionStep         bool

	_arcregionswitchAction                       string
	_arcregionswitchActionToTake                 string
	_arcregionswitchApproval                     string
	_arcregionswitchArn                          string
	_arcregionswitchAssociatedAlarms             string
	_arcregionswitchComment                      string
	_arcregionswitchDescription                  string
	_arcregionswitchExecutionId                  string
	_arcregionswitchExecutionRole                string
	_arcregionswitchHostedZoneId                 string
	_arcregionswitchLatestVersion                string
	_arcregionswitchMaxResults                   string
	_arcregionswitchMode                         string
	_arcregionswitchName                         string
	_arcregionswitchNextToken                    string
	_arcregionswitchPlanArn                      string
	_arcregionswitchPrimaryRegion                string
	_arcregionswitchRecordName                   string
	_arcregionswitchRecoveryApproach             string
	_arcregionswitchRecoveryExecutionId          string
	_arcregionswitchRecoveryTimeObjectiveMinutes string
	_arcregionswitchRegions                      []string
	_arcregionswitchReportConfiguration          string
	_arcregionswitchResourceTagKeys              []string
	_arcregionswitchState                        string
	_arcregionswitchStepName                     string
	_arcregionswitchTags                         string
	_arcregionswitchTargetRegion                 string
	_arcregionswitchTriggers                     string
	_arcregionswitchWorkflows                    string
)

// Approves a step in a plan execution that requires manual approval. When you
// create a plan, you can include approval steps that require manual intervention
// before the execution can proceed. This operation allows you to provide that
// approval.
//
// You must specify the plan ARN, execution ID, step name, and approval status.
// You can also provide an optional comment explaining the approval decision.
func arcregionswitch_ApprovePlanExecutionStep(cfg aws.Config, client *arcregionswitch.Client) {
	input := &arcregionswitch.ApprovePlanExecutionStepInput{
		// Approval: types.Approval, // Required
		// ExecutionId: *string, // Required
		// PlanArn: *string, // Required
		// StepName: *string, // Required
	}

	if len(_arcregionswitchApproval) > 0 {
		if err := assignInputField(input, "Approval", _arcregionswitchApproval); err != nil {
			log.Errorf("invalid --approval: %s", err.Error())
			return
		}
	}
	if len(_arcregionswitchExecutionId) > 0 {
		input.ExecutionId = aws.String(_arcregionswitchExecutionId)
	}
	if len(_arcregionswitchPlanArn) > 0 {
		input.PlanArn = aws.String(_arcregionswitchPlanArn)
	}
	if len(_arcregionswitchStepName) > 0 {
		input.StepName = aws.String(_arcregionswitchStepName)
	}
	if len(_arcregionswitchComment) > 0 {
		input.Comment = aws.String(_arcregionswitchComment)
	}

	if resp, err := client.ApprovePlanExecutionStep(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels an in-progress plan execution. This operation stops the execution of
// the plan and prevents any further steps from being processed.
//
// You must specify the plan ARN and execution ID. You can also provide an
// optional comment explaining why the execution was canceled.
func arcregionswitch_CancelPlanExecution(cfg aws.Config, client *arcregionswitch.Client) {
	input := &arcregionswitch.CancelPlanExecutionInput{
		// ExecutionId: *string, // Required
		// PlanArn: *string, // Required
	}

	if len(_arcregionswitchExecutionId) > 0 {
		input.ExecutionId = aws.String(_arcregionswitchExecutionId)
	}
	if len(_arcregionswitchPlanArn) > 0 {
		input.PlanArn = aws.String(_arcregionswitchPlanArn)
	}
	if len(_arcregionswitchComment) > 0 {
		input.Comment = aws.String(_arcregionswitchComment)
	}

	if resp, err := client.CancelPlanExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new Region switch plan. A plan defines the steps required to shift
// traffic from one Amazon Web Services Region to another.
//
// You must specify a name for the plan, the primary Region, and at least one
// additional Region. You can also provide a description, execution role, recovery
// time objective, associated alarms, triggers, and workflows that define the steps
// to execute during a Region switch.
func arcregionswitch_CreatePlan(cfg aws.Config, client *arcregionswitch.Client) {
	input := &arcregionswitch.CreatePlanInput{
		// ExecutionRole: *string, // Required
		// Name: *string, // Required
		// RecoveryApproach: types.RecoveryApproach, // Required
		// Regions: []string, // Required
		// Workflows: []types.Workflow, // Required
	}

	if len(_arcregionswitchExecutionRole) > 0 {
		input.ExecutionRole = aws.String(_arcregionswitchExecutionRole)
	}
	if len(_arcregionswitchName) > 0 {
		input.Name = aws.String(_arcregionswitchName)
	}
	if len(_arcregionswitchRecoveryApproach) > 0 {
		if err := assignInputField(input, "RecoveryApproach", _arcregionswitchRecoveryApproach); err != nil {
			log.Errorf("invalid --recovery-approach: %s", err.Error())
			return
		}
	}
	if len(_arcregionswitchRegions) > 0 {
		input.Regions = append([]string(nil), _arcregionswitchRegions...)
	}
	if len(_arcregionswitchWorkflows) > 0 {
		if err := assignInputField(input, "Workflows", _arcregionswitchWorkflows); err != nil {
			log.Errorf("invalid --workflows: %s", err.Error())
			return
		}
	}
	if len(_arcregionswitchAssociatedAlarms) > 0 {
		if err := assignInputField(input, "AssociatedAlarms", _arcregionswitchAssociatedAlarms); err != nil {
			log.Errorf("invalid --associated-alarms: %s", err.Error())
			return
		}
	}
	if len(_arcregionswitchDescription) > 0 {
		input.Description = aws.String(_arcregionswitchDescription)
	}
	if len(_arcregionswitchPrimaryRegion) > 0 {
		input.PrimaryRegion = aws.String(_arcregionswitchPrimaryRegion)
	}
	if len(_arcregionswitchRecoveryTimeObjectiveMinutes) > 0 {
		if err := assignInputField(input, "RecoveryTimeObjectiveMinutes", _arcregionswitchRecoveryTimeObjectiveMinutes); err != nil {
			log.Errorf("invalid --recovery-time-objective-minutes: %s", err.Error())
			return
		}
	}
	if len(_arcregionswitchReportConfiguration) > 0 {
		if err := assignInputField(input, "ReportConfiguration", _arcregionswitchReportConfiguration); err != nil {
			log.Errorf("invalid --report-configuration: %s", err.Error())
			return
		}
	}
	if len(_arcregionswitchTags) > 0 {
		if err := assignInputField(input, "Tags", _arcregionswitchTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_arcregionswitchTriggers) > 0 {
		if err := assignInputField(input, "Triggers", _arcregionswitchTriggers); err != nil {
			log.Errorf("invalid --triggers: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreatePlan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a Region switch plan. You must specify the ARN of the plan to delete.
// You cannot delete a plan that has an active execution in progress.
func arcregionswitch_DeletePlan(cfg aws.Config, client *arcregionswitch.Client) {
	input := &arcregionswitch.DeletePlanInput{
		// Arn: *string, // Required
	}

	if len(_arcregionswitchArn) > 0 {
		input.Arn = aws.String(_arcregionswitchArn)
	}

	if resp, err := client.DeletePlan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves detailed information about a Region switch plan. You must specify the
// ARN of the plan.
func arcregionswitch_GetPlan(cfg aws.Config, client *arcregionswitch.Client) {
	input := &arcregionswitch.GetPlanInput{
		// Arn: *string, // Required
	}

	if len(_arcregionswitchArn) > 0 {
		input.Arn = aws.String(_arcregionswitchArn)
	}

	if resp, err := client.GetPlan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the evaluation status of a Region switch plan. The evaluation status
// provides information about the last time the plan was evaluated and any warnings
// or issues detected.
func arcregionswitch_GetPlanEvaluationStatus(cfg aws.Config, client *arcregionswitch.Client) {
	input := &arcregionswitch.GetPlanEvaluationStatusInput{
		// PlanArn: *string, // Required
	}

	if len(_arcregionswitchPlanArn) > 0 {
		input.PlanArn = aws.String(_arcregionswitchPlanArn)
	}
	if len(_arcregionswitchMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _arcregionswitchMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_arcregionswitchNextToken) > 0 {
		input.NextToken = aws.String(_arcregionswitchNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetPlanEvaluationStatus(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*arcregionswitch.GetPlanEvaluationStatusOutput
	p := arcregionswitch.NewGetPlanEvaluationStatusPaginator(client, input)
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

// Retrieves detailed information about a specific plan execution. You must
// specify the plan ARN and execution ID.
func arcregionswitch_GetPlanExecution(cfg aws.Config, client *arcregionswitch.Client) {
	input := &arcregionswitch.GetPlanExecutionInput{
		// ExecutionId: *string, // Required
		// PlanArn: *string, // Required
	}

	if len(_arcregionswitchExecutionId) > 0 {
		input.ExecutionId = aws.String(_arcregionswitchExecutionId)
	}
	if len(_arcregionswitchPlanArn) > 0 {
		input.PlanArn = aws.String(_arcregionswitchPlanArn)
	}
	if len(_arcregionswitchMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _arcregionswitchMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_arcregionswitchNextToken) > 0 {
		input.NextToken = aws.String(_arcregionswitchNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetPlanExecution(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*arcregionswitch.GetPlanExecutionOutput
	p := arcregionswitch.NewGetPlanExecutionPaginator(client, input)
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

// Retrieves information about a Region switch plan in a specific Amazon Web
// Services Region. This operation is useful for getting Region-specific
// information about a plan.
func arcregionswitch_GetPlanInRegion(cfg aws.Config, client *arcregionswitch.Client) {
	input := &arcregionswitch.GetPlanInRegionInput{
		// Arn: *string, // Required
	}

	if len(_arcregionswitchArn) > 0 {
		input.Arn = aws.String(_arcregionswitchArn)
	}

	if resp, err := client.GetPlanInRegion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the events that occurred during a plan execution. These events provide a
// detailed timeline of the execution process.
func arcregionswitch_ListPlanExecutionEvents(cfg aws.Config, client *arcregionswitch.Client) {
	input := &arcregionswitch.ListPlanExecutionEventsInput{
		// ExecutionId: *string, // Required
		// PlanArn: *string, // Required
	}

	if len(_arcregionswitchExecutionId) > 0 {
		input.ExecutionId = aws.String(_arcregionswitchExecutionId)
	}
	if len(_arcregionswitchPlanArn) > 0 {
		input.PlanArn = aws.String(_arcregionswitchPlanArn)
	}
	if len(_arcregionswitchMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _arcregionswitchMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_arcregionswitchName) > 0 {
		input.Name = aws.String(_arcregionswitchName)
	}
	if len(_arcregionswitchNextToken) > 0 {
		input.NextToken = aws.String(_arcregionswitchNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPlanExecutionEvents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*arcregionswitch.ListPlanExecutionEventsOutput
	p := arcregionswitch.NewListPlanExecutionEventsPaginator(client, input)
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

// Lists the executions of a Region switch plan. This operation returns
// information about both current and historical executions.
func arcregionswitch_ListPlanExecutions(cfg aws.Config, client *arcregionswitch.Client) {
	input := &arcregionswitch.ListPlanExecutionsInput{
		// PlanArn: *string, // Required
	}

	if len(_arcregionswitchPlanArn) > 0 {
		input.PlanArn = aws.String(_arcregionswitchPlanArn)
	}
	if len(_arcregionswitchMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _arcregionswitchMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_arcregionswitchNextToken) > 0 {
		input.NextToken = aws.String(_arcregionswitchNextToken)
	}
	if len(_arcregionswitchState) > 0 {
		if err := assignInputField(input, "State", _arcregionswitchState); err != nil {
			log.Errorf("invalid --state: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListPlanExecutions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*arcregionswitch.ListPlanExecutionsOutput
	p := arcregionswitch.NewListPlanExecutionsPaginator(client, input)
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

// Lists all Region switch plans in your Amazon Web Services account.
func arcregionswitch_ListPlans(cfg aws.Config, client *arcregionswitch.Client) {
	input := &arcregionswitch.ListPlansInput{}

	if len(_arcregionswitchMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _arcregionswitchMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_arcregionswitchNextToken) > 0 {
		input.NextToken = aws.String(_arcregionswitchNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPlans(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*arcregionswitch.ListPlansOutput
	p := arcregionswitch.NewListPlansPaginator(client, input)
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

// Lists all Region switch plans in your Amazon Web Services account that are
// available in the current Amazon Web Services Region.
func arcregionswitch_ListPlansInRegion(cfg aws.Config, client *arcregionswitch.Client) {
	input := &arcregionswitch.ListPlansInRegionInput{}

	if len(_arcregionswitchMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _arcregionswitchMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_arcregionswitchNextToken) > 0 {
		input.NextToken = aws.String(_arcregionswitchNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPlansInRegion(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*arcregionswitch.ListPlansInRegionOutput
	p := arcregionswitch.NewListPlansInRegionPaginator(client, input)
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

// List the Amazon Route 53 health checks.
func arcregionswitch_ListRoute53HealthChecks(cfg aws.Config, client *arcregionswitch.Client) {
	input := &arcregionswitch.ListRoute53HealthChecksInput{
		// Arn: *string, // Required
	}

	if len(_arcregionswitchArn) > 0 {
		input.Arn = aws.String(_arcregionswitchArn)
	}
	if len(_arcregionswitchHostedZoneId) > 0 {
		input.HostedZoneId = aws.String(_arcregionswitchHostedZoneId)
	}
	if len(_arcregionswitchMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _arcregionswitchMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_arcregionswitchNextToken) > 0 {
		input.NextToken = aws.String(_arcregionswitchNextToken)
	}
	if len(_arcregionswitchRecordName) > 0 {
		input.RecordName = aws.String(_arcregionswitchRecordName)
	}

	if disablePaginator() {
		if resp, err := client.ListRoute53HealthChecks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*arcregionswitch.ListRoute53HealthChecksOutput
	p := arcregionswitch.NewListRoute53HealthChecksPaginator(client, input)
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

// List the Amazon Route 53 health checks in a specific Amazon Web Services Region.
func arcregionswitch_ListRoute53HealthChecksInRegion(cfg aws.Config, client *arcregionswitch.Client) {
	input := &arcregionswitch.ListRoute53HealthChecksInRegionInput{
		// Arn: *string, // Required
	}

	if len(_arcregionswitchArn) > 0 {
		input.Arn = aws.String(_arcregionswitchArn)
	}
	if len(_arcregionswitchHostedZoneId) > 0 {
		input.HostedZoneId = aws.String(_arcregionswitchHostedZoneId)
	}
	if len(_arcregionswitchMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _arcregionswitchMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_arcregionswitchNextToken) > 0 {
		input.NextToken = aws.String(_arcregionswitchNextToken)
	}
	if len(_arcregionswitchRecordName) > 0 {
		input.RecordName = aws.String(_arcregionswitchRecordName)
	}

	if disablePaginator() {
		if resp, err := client.ListRoute53HealthChecksInRegion(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*arcregionswitch.ListRoute53HealthChecksInRegionOutput
	p := arcregionswitch.NewListRoute53HealthChecksInRegionPaginator(client, input)
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

// Lists the tags attached to a Region switch resource.
func arcregionswitch_ListTagsForResource(cfg aws.Config, client *arcregionswitch.Client) {
	input := &arcregionswitch.ListTagsForResourceInput{
		// Arn: *string, // Required
	}

	if len(_arcregionswitchArn) > 0 {
		input.Arn = aws.String(_arcregionswitchArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts the execution of a Region switch plan. You can execute a plan in either
// graceful or ungraceful mode.
//
// Specifing ungraceful mode either changes the behavior of the execution blocks
// in a workflow or skips specific execution blocks.
func arcregionswitch_StartPlanExecution(cfg aws.Config, client *arcregionswitch.Client) {
	input := &arcregionswitch.StartPlanExecutionInput{
		// Action: types.ExecutionAction, // Required
		// PlanArn: *string, // Required
		// TargetRegion: *string, // Required
	}

	if len(_arcregionswitchAction) > 0 {
		if err := assignInputField(input, "Action", _arcregionswitchAction); err != nil {
			log.Errorf("invalid --action: %s", err.Error())
			return
		}
	}
	if len(_arcregionswitchPlanArn) > 0 {
		input.PlanArn = aws.String(_arcregionswitchPlanArn)
	}
	if len(_arcregionswitchTargetRegion) > 0 {
		input.TargetRegion = aws.String(_arcregionswitchTargetRegion)
	}
	if len(_arcregionswitchComment) > 0 {
		input.Comment = aws.String(_arcregionswitchComment)
	}
	if len(_arcregionswitchLatestVersion) > 0 {
		input.LatestVersion = aws.String(_arcregionswitchLatestVersion)
	}
	if len(_arcregionswitchMode) > 0 {
		if err := assignInputField(input, "Mode", _arcregionswitchMode); err != nil {
			log.Errorf("invalid --mode: %s", err.Error())
			return
		}
	}
	if len(_arcregionswitchRecoveryExecutionId) > 0 {
		input.RecoveryExecutionId = aws.String(_arcregionswitchRecoveryExecutionId)
	}

	if resp, err := client.StartPlanExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds or updates tags for a Region switch resource. You can assign metadata to
// your resources in the form of tags, which are key-value pairs.
func arcregionswitch_TagResource(cfg aws.Config, client *arcregionswitch.Client) {
	input := &arcregionswitch.TagResourceInput{
		// Arn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_arcregionswitchArn) > 0 {
		input.Arn = aws.String(_arcregionswitchArn)
	}
	if len(_arcregionswitchTags) > 0 {
		if err := assignInputField(input, "Tags", _arcregionswitchTags); err != nil {
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

// Removes tags from a Region switch resource.
func arcregionswitch_UntagResource(cfg aws.Config, client *arcregionswitch.Client) {
	input := &arcregionswitch.UntagResourceInput{
		// Arn: *string, // Required
		// ResourceTagKeys: []string, // Required
	}

	if len(_arcregionswitchArn) > 0 {
		input.Arn = aws.String(_arcregionswitchArn)
	}
	if len(_arcregionswitchResourceTagKeys) > 0 {
		input.ResourceTagKeys = append([]string(nil), _arcregionswitchResourceTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing Region switch plan. You can modify the plan's description,
// workflows, execution role, recovery time objective, associated alarms, and
// triggers.
func arcregionswitch_UpdatePlan(cfg aws.Config, client *arcregionswitch.Client) {
	input := &arcregionswitch.UpdatePlanInput{
		// Arn: *string, // Required
		// ExecutionRole: *string, // Required
		// Workflows: []types.Workflow, // Required
	}

	if len(_arcregionswitchArn) > 0 {
		input.Arn = aws.String(_arcregionswitchArn)
	}
	if len(_arcregionswitchExecutionRole) > 0 {
		input.ExecutionRole = aws.String(_arcregionswitchExecutionRole)
	}
	if len(_arcregionswitchWorkflows) > 0 {
		if err := assignInputField(input, "Workflows", _arcregionswitchWorkflows); err != nil {
			log.Errorf("invalid --workflows: %s", err.Error())
			return
		}
	}
	if len(_arcregionswitchAssociatedAlarms) > 0 {
		if err := assignInputField(input, "AssociatedAlarms", _arcregionswitchAssociatedAlarms); err != nil {
			log.Errorf("invalid --associated-alarms: %s", err.Error())
			return
		}
	}
	if len(_arcregionswitchDescription) > 0 {
		input.Description = aws.String(_arcregionswitchDescription)
	}
	if len(_arcregionswitchRecoveryTimeObjectiveMinutes) > 0 {
		if err := assignInputField(input, "RecoveryTimeObjectiveMinutes", _arcregionswitchRecoveryTimeObjectiveMinutes); err != nil {
			log.Errorf("invalid --recovery-time-objective-minutes: %s", err.Error())
			return
		}
	}
	if len(_arcregionswitchReportConfiguration) > 0 {
		if err := assignInputField(input, "ReportConfiguration", _arcregionswitchReportConfiguration); err != nil {
			log.Errorf("invalid --report-configuration: %s", err.Error())
			return
		}
	}
	if len(_arcregionswitchTriggers) > 0 {
		if err := assignInputField(input, "Triggers", _arcregionswitchTriggers); err != nil {
			log.Errorf("invalid --triggers: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdatePlan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an in-progress plan execution. This operation allows you to modify
// certain aspects of the execution, such as adding a comment or changing the
// action.
func arcregionswitch_UpdatePlanExecution(cfg aws.Config, client *arcregionswitch.Client) {
	input := &arcregionswitch.UpdatePlanExecutionInput{
		// Action: types.UpdatePlanExecutionAction, // Required
		// ExecutionId: *string, // Required
		// PlanArn: *string, // Required
	}

	if len(_arcregionswitchAction) > 0 {
		if err := assignInputField(input, "Action", _arcregionswitchAction); err != nil {
			log.Errorf("invalid --action: %s", err.Error())
			return
		}
	}
	if len(_arcregionswitchExecutionId) > 0 {
		input.ExecutionId = aws.String(_arcregionswitchExecutionId)
	}
	if len(_arcregionswitchPlanArn) > 0 {
		input.PlanArn = aws.String(_arcregionswitchPlanArn)
	}
	if len(_arcregionswitchComment) > 0 {
		input.Comment = aws.String(_arcregionswitchComment)
	}

	if resp, err := client.UpdatePlanExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a specific step in an in-progress plan execution. This operation allows
// you to modify the step's comment or action.
func arcregionswitch_UpdatePlanExecutionStep(cfg aws.Config, client *arcregionswitch.Client) {
	input := &arcregionswitch.UpdatePlanExecutionStepInput{
		// ActionToTake: types.UpdatePlanExecutionStepAction, // Required
		// Comment: *string, // Required
		// ExecutionId: *string, // Required
		// PlanArn: *string, // Required
		// StepName: *string, // Required
	}

	if len(_arcregionswitchActionToTake) > 0 {
		if err := assignInputField(input, "ActionToTake", _arcregionswitchActionToTake); err != nil {
			log.Errorf("invalid --action-to-take: %s", err.Error())
			return
		}
	}
	if len(_arcregionswitchComment) > 0 {
		input.Comment = aws.String(_arcregionswitchComment)
	}
	if len(_arcregionswitchExecutionId) > 0 {
		input.ExecutionId = aws.String(_arcregionswitchExecutionId)
	}
	if len(_arcregionswitchPlanArn) > 0 {
		input.PlanArn = aws.String(_arcregionswitchPlanArn)
	}
	if len(_arcregionswitchStepName) > 0 {
		input.StepName = aws.String(_arcregionswitchStepName)
	}

	if resp, err := client.UpdatePlanExecutionStep(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_arcregionswitchCmd)
	_arcregionswitchCmd.Flags().SortFlags = false

	_arcregionswitchCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_arcregionswitchCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_arcregionswitchCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_arcregionswitchCmd.Flags().StringVarP(&_arcregionswitchAction, "action", "", "", "Action")
	_arcregionswitchCmd.Flags().StringVarP(&_arcregionswitchActionToTake, "action-to-take", "", "", "Action To Take")
	_arcregionswitchCmd.Flags().StringVarP(&_arcregionswitchApproval, "approval", "", "", "Approval")
	_arcregionswitchCmd.Flags().StringVarP(&_arcregionswitchArn, "arn", "", "", "ARN")
	_arcregionswitchCmd.Flags().StringVarP(&_arcregionswitchAssociatedAlarms, "associated-alarms", "", "", "Associated Alarms")
	_arcregionswitchCmd.Flags().StringVarP(&_arcregionswitchComment, "comment", "", "", "Comment")
	_arcregionswitchCmd.Flags().StringVarP(&_arcregionswitchDescription, "description", "", "", "Description")
	_arcregionswitchCmd.Flags().StringVarP(&_arcregionswitchExecutionId, "execution-id", "", "", "Execution ID")
	_arcregionswitchCmd.Flags().StringVarP(&_arcregionswitchExecutionRole, "execution-role", "", "", "Execution Role")
	_arcregionswitchCmd.Flags().StringVarP(&_arcregionswitchHostedZoneId, "hosted-zone-id", "", "", "Hosted Zone ID")
	_arcregionswitchCmd.Flags().StringVarP(&_arcregionswitchLatestVersion, "latest-version", "", "", "Latest Version")
	_arcregionswitchCmd.Flags().StringVarP(&_arcregionswitchMaxResults, "max-results", "", "", "Max Results")
	_arcregionswitchCmd.Flags().StringVarP(&_arcregionswitchMode, "mode", "", "", "Mode")
	_arcregionswitchCmd.Flags().StringVarP(&_arcregionswitchName, "name", "", "", "Name")
	_arcregionswitchCmd.Flags().StringVarP(&_arcregionswitchNextToken, "next-token", "", "", "Next Token")
	_arcregionswitchCmd.Flags().StringVarP(&_arcregionswitchPlanArn, "plan-arn", "", "", "Plan ARN")
	_arcregionswitchCmd.Flags().StringVarP(&_arcregionswitchPrimaryRegion, "primary-region", "", "", "Primary Region")
	_arcregionswitchCmd.Flags().StringVarP(&_arcregionswitchRecordName, "record-name", "", "", "Record Name")
	_arcregionswitchCmd.Flags().StringVarP(&_arcregionswitchRecoveryApproach, "recovery-approach", "", "", "Recovery Approach")
	_arcregionswitchCmd.Flags().StringVarP(&_arcregionswitchRecoveryExecutionId, "recovery-execution-id", "", "", "Recovery Execution ID")
	_arcregionswitchCmd.Flags().StringVarP(&_arcregionswitchRecoveryTimeObjectiveMinutes, "recovery-time-objective-minutes", "", "", "Recovery Time Objective Minutes")
	_arcregionswitchCmd.Flags().StringSliceVarP(&_arcregionswitchRegions, "regions", "", nil, "Regions")
	_arcregionswitchCmd.Flags().StringVarP(&_arcregionswitchReportConfiguration, "report-configuration", "", "", "Report Configuration")
	_arcregionswitchCmd.Flags().StringSliceVarP(&_arcregionswitchResourceTagKeys, "resource-tag-keys", "", nil, "Resource Tag Keys")
	_arcregionswitchCmd.Flags().StringVarP(&_arcregionswitchState, "state", "", "", "State")
	_arcregionswitchCmd.Flags().StringVarP(&_arcregionswitchStepName, "step-name", "", "", "Step Name")
	_arcregionswitchCmd.Flags().StringVarP(&_arcregionswitchTags, "tags", "", "", "Tags")
	_arcregionswitchCmd.Flags().StringVarP(&_arcregionswitchTargetRegion, "target-region", "", "", "Target Region")
	_arcregionswitchCmd.Flags().StringVarP(&_arcregionswitchTriggers, "triggers", "", "", "Triggers")
	_arcregionswitchCmd.Flags().StringVarP(&_arcregionswitchWorkflows, "workflows", "", "", "Workflows")

	_arcregionswitchCmd.Flags().BoolVarP(&_arcregionswitchApprovePlanExecutionStep, "approve-plan-execution-step", "", false, "Approve Plan Execution Step")
	_arcregionswitchCmd.Flags().BoolVarP(&_arcregionswitchCancelPlanExecution, "cancel-plan-execution", "", false, "Cancel Plan Execution")
	_arcregionswitchCmd.Flags().BoolVarP(&_arcregionswitchCreatePlan, "create-plan", "", false, "Create Plan")
	_arcregionswitchCmd.Flags().BoolVarP(&_arcregionswitchDeletePlan, "delete-plan", "", false, "Delete Plan")
	_arcregionswitchCmd.Flags().BoolVarP(&_arcregionswitchGetPlan, "get-plan", "", false, "Get Plan")
	_arcregionswitchCmd.Flags().BoolVarP(&_arcregionswitchGetPlanEvaluationStatus, "get-plan-evaluation-status", "", false, "Get Plan Evaluation Status")
	_arcregionswitchCmd.Flags().BoolVarP(&_arcregionswitchGetPlanExecution, "get-plan-execution", "", false, "Get Plan Execution")
	_arcregionswitchCmd.Flags().BoolVarP(&_arcregionswitchGetPlanInRegion, "get-plan-in-region", "", false, "Get Plan In Region")
	_arcregionswitchCmd.Flags().BoolVarP(&_arcregionswitchListPlanExecutionEvents, "list-plan-execution-events", "", false, "List Plan Execution Events")
	_arcregionswitchCmd.Flags().BoolVarP(&_arcregionswitchListPlanExecutions, "list-plan-executions", "", false, "List Plan Executions")
	_arcregionswitchCmd.Flags().BoolVarP(&_arcregionswitchListPlans, "list-plans", "", false, "List Plans")
	_arcregionswitchCmd.Flags().BoolVarP(&_arcregionswitchListPlansInRegion, "list-plans-in-region", "", false, "List Plans In Region")
	_arcregionswitchCmd.Flags().BoolVarP(&_arcregionswitchListRoute53HealthChecks, "list-route53-health-checks", "", false, "List Route53 Health Checks")
	_arcregionswitchCmd.Flags().BoolVarP(&_arcregionswitchListRoute53HealthChecksInRegion, "list-route53-health-checks-in-region", "", false, "List Route53 Health Checks In Region")
	_arcregionswitchCmd.Flags().BoolVarP(&_arcregionswitchListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_arcregionswitchCmd.Flags().BoolVarP(&_arcregionswitchStartPlanExecution, "start-plan-execution", "", false, "Start Plan Execution")
	_arcregionswitchCmd.Flags().BoolVarP(&_arcregionswitchTagResource, "tag-resource", "", false, "Tag Resource")
	_arcregionswitchCmd.Flags().BoolVarP(&_arcregionswitchUntagResource, "untag-resource", "", false, "Untag Resource")
	_arcregionswitchCmd.Flags().BoolVarP(&_arcregionswitchUpdatePlan, "update-plan", "", false, "Update Plan")
	_arcregionswitchCmd.Flags().BoolVarP(&_arcregionswitchUpdatePlanExecution, "update-plan-execution", "", false, "Update Plan Execution")
	_arcregionswitchCmd.Flags().BoolVarP(&_arcregionswitchUpdatePlanExecutionStep, "update-plan-execution-step", "", false, "Update Plan Execution Step")

}
