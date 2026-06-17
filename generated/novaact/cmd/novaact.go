package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/novaact"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// novaactCmd represents the novaact command
var _novaactCmd = &cobra.Command{
	Use:   "novaact",
	Short: "AWS novaact CLI",
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
		client := novaact.NewFromConfig(cfg)
		if _novaactCreateAct {
			novaact_CreateAct(cfg, client)
			return
		}
		if _novaactCreateSession {
			novaact_CreateSession(cfg, client)
			return
		}
		if _novaactCreateWorkflowDefinition {
			novaact_CreateWorkflowDefinition(cfg, client)
			return
		}
		if _novaactCreateWorkflowRun {
			novaact_CreateWorkflowRun(cfg, client)
			return
		}
		if _novaactDeleteWorkflowDefinition {
			novaact_DeleteWorkflowDefinition(cfg, client)
			return
		}
		if _novaactDeleteWorkflowRun {
			novaact_DeleteWorkflowRun(cfg, client)
			return
		}
		if _novaactGetWorkflowDefinition {
			novaact_GetWorkflowDefinition(cfg, client)
			return
		}
		if _novaactGetWorkflowRun {
			novaact_GetWorkflowRun(cfg, client)
			return
		}
		if _novaactInvokeActStep {
			novaact_InvokeActStep(cfg, client)
			return
		}
		if _novaactListActs {
			novaact_ListActs(cfg, client)
			return
		}
		if _novaactListModels {
			novaact_ListModels(cfg, client)
			return
		}
		if _novaactListSessions {
			novaact_ListSessions(cfg, client)
			return
		}
		if _novaactListWorkflowDefinitions {
			novaact_ListWorkflowDefinitions(cfg, client)
			return
		}
		if _novaactListWorkflowRuns {
			novaact_ListWorkflowRuns(cfg, client)
			return
		}
		if _novaactUpdateAct {
			novaact_UpdateAct(cfg, client)
			return
		}
		if _novaactUpdateWorkflowRun {
			novaact_UpdateWorkflowRun(cfg, client)
			return
		}

	},
}

var (
	_novaactCreateAct                bool
	_novaactCreateSession            bool
	_novaactCreateWorkflowDefinition bool
	_novaactCreateWorkflowRun        bool
	_novaactDeleteWorkflowDefinition bool
	_novaactDeleteWorkflowRun        bool
	_novaactGetWorkflowDefinition    bool
	_novaactGetWorkflowRun           bool
	_novaactInvokeActStep            bool
	_novaactListActs                 bool
	_novaactListModels               bool
	_novaactListSessions             bool
	_novaactListWorkflowDefinitions  bool
	_novaactListWorkflowRuns         bool
	_novaactUpdateAct                bool
	_novaactUpdateWorkflowRun        bool

	_novaactActId                      string
	_novaactCallResults                string
	_novaactClientCompatibilityVersion string
	_novaactClientInfo                 string
	_novaactClientToken                string
	_novaactDescription                string
	_novaactError                      string
	_novaactExportConfig               string
	_novaactLogGroupName               string
	_novaactMaxResults                 string
	_novaactModelId                    string
	_novaactName                       string
	_novaactNextToken                  string
	_novaactPreviousStepId             string
	_novaactSessionId                  string
	_novaactSortOrder                  string
	_novaactStatus                     string
	_novaactTask                       string
	_novaactToolSpecs                  string
	_novaactWorkflowDefinitionName     string
	_novaactWorkflowRunId              string
)

// Creates a new AI task (act) within a session that can interact with tools and
// perform specific actions.
func novaact_CreateAct(cfg aws.Config, client *novaact.Client) {
	input := &novaact.CreateActInput{
		// SessionId: *string, // Required
		// Task: *string, // Required
		// WorkflowDefinitionName: *string, // Required
		// WorkflowRunId: *string, // Required
	}

	if len(_novaactSessionId) > 0 {
		input.SessionId = aws.String(_novaactSessionId)
	}
	if len(_novaactTask) > 0 {
		input.Task = aws.String(_novaactTask)
	}
	if len(_novaactWorkflowDefinitionName) > 0 {
		input.WorkflowDefinitionName = aws.String(_novaactWorkflowDefinitionName)
	}
	if len(_novaactWorkflowRunId) > 0 {
		input.WorkflowRunId = aws.String(_novaactWorkflowRunId)
	}
	if len(_novaactClientToken) > 0 {
		input.ClientToken = aws.String(_novaactClientToken)
	}
	if len(_novaactToolSpecs) > 0 {
		if err := assignInputField(input, "ToolSpecs", _novaactToolSpecs); err != nil {
			log.Errorf("invalid --tool-specs: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAct(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new session context within a workflow run to manage conversation
// state and acts.
func novaact_CreateSession(cfg aws.Config, client *novaact.Client) {
	input := &novaact.CreateSessionInput{
		// WorkflowDefinitionName: *string, // Required
		// WorkflowRunId: *string, // Required
	}

	if len(_novaactWorkflowDefinitionName) > 0 {
		input.WorkflowDefinitionName = aws.String(_novaactWorkflowDefinitionName)
	}
	if len(_novaactWorkflowRunId) > 0 {
		input.WorkflowRunId = aws.String(_novaactWorkflowRunId)
	}
	if len(_novaactClientToken) > 0 {
		input.ClientToken = aws.String(_novaactClientToken)
	}

	if resp, err := client.CreateSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new workflow definition template that can be used to execute multiple
// workflow runs.
func novaact_CreateWorkflowDefinition(cfg aws.Config, client *novaact.Client) {
	input := &novaact.CreateWorkflowDefinitionInput{
		// Name: *string, // Required
	}

	if len(_novaactName) > 0 {
		input.Name = aws.String(_novaactName)
	}
	if len(_novaactClientToken) > 0 {
		input.ClientToken = aws.String(_novaactClientToken)
	}
	if len(_novaactDescription) > 0 {
		input.Description = aws.String(_novaactDescription)
	}
	if len(_novaactExportConfig) > 0 {
		if err := assignInputField(input, "ExportConfig", _novaactExportConfig); err != nil {
			log.Errorf("invalid --export-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateWorkflowDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new execution instance of a workflow definition with specified
// parameters.
func novaact_CreateWorkflowRun(cfg aws.Config, client *novaact.Client) {
	input := &novaact.CreateWorkflowRunInput{
		// ClientInfo: *types.ClientInfo, // Required
		// ModelId: *string, // Required
		// WorkflowDefinitionName: *string, // Required
	}

	if len(_novaactClientInfo) > 0 {
		if err := assignInputField(input, "ClientInfo", _novaactClientInfo); err != nil {
			log.Errorf("invalid --client-info: %s", err.Error())
			return
		}
	}
	if len(_novaactModelId) > 0 {
		input.ModelId = aws.String(_novaactModelId)
	}
	if len(_novaactWorkflowDefinitionName) > 0 {
		input.WorkflowDefinitionName = aws.String(_novaactWorkflowDefinitionName)
	}
	if len(_novaactClientToken) > 0 {
		input.ClientToken = aws.String(_novaactClientToken)
	}
	if len(_novaactLogGroupName) > 0 {
		input.LogGroupName = aws.String(_novaactLogGroupName)
	}

	if resp, err := client.CreateWorkflowRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a workflow definition and all associated resources. This operation
// cannot be undone.
func novaact_DeleteWorkflowDefinition(cfg aws.Config, client *novaact.Client) {
	input := &novaact.DeleteWorkflowDefinitionInput{
		// WorkflowDefinitionName: *string, // Required
	}

	if len(_novaactWorkflowDefinitionName) > 0 {
		input.WorkflowDefinitionName = aws.String(_novaactWorkflowDefinitionName)
	}

	if resp, err := client.DeleteWorkflowDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Terminates and cleans up a workflow run, stopping all associated acts and
// sessions.
func novaact_DeleteWorkflowRun(cfg aws.Config, client *novaact.Client) {
	input := &novaact.DeleteWorkflowRunInput{
		// WorkflowDefinitionName: *string, // Required
		// WorkflowRunId: *string, // Required
	}

	if len(_novaactWorkflowDefinitionName) > 0 {
		input.WorkflowDefinitionName = aws.String(_novaactWorkflowDefinitionName)
	}
	if len(_novaactWorkflowRunId) > 0 {
		input.WorkflowRunId = aws.String(_novaactWorkflowRunId)
	}

	if resp, err := client.DeleteWorkflowRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the details and configuration of a specific workflow definition.
func novaact_GetWorkflowDefinition(cfg aws.Config, client *novaact.Client) {
	input := &novaact.GetWorkflowDefinitionInput{
		// WorkflowDefinitionName: *string, // Required
	}

	if len(_novaactWorkflowDefinitionName) > 0 {
		input.WorkflowDefinitionName = aws.String(_novaactWorkflowDefinitionName)
	}

	if resp, err := client.GetWorkflowDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the current state, configuration, and execution details of a workflow
// run.
func novaact_GetWorkflowRun(cfg aws.Config, client *novaact.Client) {
	input := &novaact.GetWorkflowRunInput{
		// WorkflowDefinitionName: *string, // Required
		// WorkflowRunId: *string, // Required
	}

	if len(_novaactWorkflowDefinitionName) > 0 {
		input.WorkflowDefinitionName = aws.String(_novaactWorkflowDefinitionName)
	}
	if len(_novaactWorkflowRunId) > 0 {
		input.WorkflowRunId = aws.String(_novaactWorkflowRunId)
	}

	if resp, err := client.GetWorkflowRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Executes the next step of an act, processing tool call results and returning
// new tool calls if needed.
func novaact_InvokeActStep(cfg aws.Config, client *novaact.Client) {
	input := &novaact.InvokeActStepInput{
		// ActId: *string, // Required
		// CallResults: []types.CallResult, // Required
		// SessionId: *string, // Required
		// WorkflowDefinitionName: *string, // Required
		// WorkflowRunId: *string, // Required
	}

	if len(_novaactActId) > 0 {
		input.ActId = aws.String(_novaactActId)
	}
	if len(_novaactCallResults) > 0 {
		if err := assignInputField(input, "CallResults", _novaactCallResults); err != nil {
			log.Errorf("invalid --call-results: %s", err.Error())
			return
		}
	}
	if len(_novaactSessionId) > 0 {
		input.SessionId = aws.String(_novaactSessionId)
	}
	if len(_novaactWorkflowDefinitionName) > 0 {
		input.WorkflowDefinitionName = aws.String(_novaactWorkflowDefinitionName)
	}
	if len(_novaactWorkflowRunId) > 0 {
		input.WorkflowRunId = aws.String(_novaactWorkflowRunId)
	}
	if len(_novaactPreviousStepId) > 0 {
		input.PreviousStepId = aws.String(_novaactPreviousStepId)
	}

	if resp, err := client.InvokeActStep(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all acts within a specific session with their current status and
// execution details.
func novaact_ListActs(cfg aws.Config, client *novaact.Client) {
	input := &novaact.ListActsInput{
		// WorkflowDefinitionName: *string, // Required
	}

	if len(_novaactWorkflowDefinitionName) > 0 {
		input.WorkflowDefinitionName = aws.String(_novaactWorkflowDefinitionName)
	}
	if len(_novaactMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _novaactMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_novaactNextToken) > 0 {
		input.NextToken = aws.String(_novaactNextToken)
	}
	if len(_novaactSessionId) > 0 {
		input.SessionId = aws.String(_novaactSessionId)
	}
	if len(_novaactSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _novaactSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}
	if len(_novaactWorkflowRunId) > 0 {
		input.WorkflowRunId = aws.String(_novaactWorkflowRunId)
	}

	if disablePaginator() {
		if resp, err := client.ListActs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*novaact.ListActsOutput
	p := novaact.NewListActsPaginator(client, input)
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

// Lists all available AI models that can be used for workflow execution,
// including their status and compatibility information.
func novaact_ListModels(cfg aws.Config, client *novaact.Client) {
	input := &novaact.ListModelsInput{
		// ClientCompatibilityVersion: *int32, // Required
	}

	if len(_novaactClientCompatibilityVersion) > 0 {
		if err := assignInputField(input, "ClientCompatibilityVersion", _novaactClientCompatibilityVersion); err != nil {
			log.Errorf("invalid --client-compatibility-version: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListModels(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all sessions within a specific workflow run.
func novaact_ListSessions(cfg aws.Config, client *novaact.Client) {
	input := &novaact.ListSessionsInput{
		// WorkflowDefinitionName: *string, // Required
		// WorkflowRunId: *string, // Required
	}

	if len(_novaactWorkflowDefinitionName) > 0 {
		input.WorkflowDefinitionName = aws.String(_novaactWorkflowDefinitionName)
	}
	if len(_novaactWorkflowRunId) > 0 {
		input.WorkflowRunId = aws.String(_novaactWorkflowRunId)
	}
	if len(_novaactMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _novaactMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_novaactNextToken) > 0 {
		input.NextToken = aws.String(_novaactNextToken)
	}
	if len(_novaactSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _novaactSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListSessions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*novaact.ListSessionsOutput
	p := novaact.NewListSessionsPaginator(client, input)
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

// Lists all workflow definitions in your account with optional filtering and
// pagination.
func novaact_ListWorkflowDefinitions(cfg aws.Config, client *novaact.Client) {
	input := &novaact.ListWorkflowDefinitionsInput{}

	if len(_novaactMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _novaactMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_novaactNextToken) > 0 {
		input.NextToken = aws.String(_novaactNextToken)
	}
	if len(_novaactSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _novaactSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListWorkflowDefinitions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*novaact.ListWorkflowDefinitionsOutput
	p := novaact.NewListWorkflowDefinitionsPaginator(client, input)
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

// Lists all workflow runs for a specific workflow definition with optional
// filtering and pagination.
func novaact_ListWorkflowRuns(cfg aws.Config, client *novaact.Client) {
	input := &novaact.ListWorkflowRunsInput{
		// WorkflowDefinitionName: *string, // Required
	}

	if len(_novaactWorkflowDefinitionName) > 0 {
		input.WorkflowDefinitionName = aws.String(_novaactWorkflowDefinitionName)
	}
	if len(_novaactMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _novaactMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_novaactNextToken) > 0 {
		input.NextToken = aws.String(_novaactNextToken)
	}
	if len(_novaactSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _novaactSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
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

	var results []*novaact.ListWorkflowRunsOutput
	p := novaact.NewListWorkflowRunsPaginator(client, input)
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

// Updates an existing act's configuration, status, or error information.
func novaact_UpdateAct(cfg aws.Config, client *novaact.Client) {
	input := &novaact.UpdateActInput{
		// ActId: *string, // Required
		// SessionId: *string, // Required
		// Status: types.ActStatus, // Required
		// WorkflowDefinitionName: *string, // Required
		// WorkflowRunId: *string, // Required
	}

	if len(_novaactActId) > 0 {
		input.ActId = aws.String(_novaactActId)
	}
	if len(_novaactSessionId) > 0 {
		input.SessionId = aws.String(_novaactSessionId)
	}
	if len(_novaactStatus) > 0 {
		if err := assignInputField(input, "Status", _novaactStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_novaactWorkflowDefinitionName) > 0 {
		input.WorkflowDefinitionName = aws.String(_novaactWorkflowDefinitionName)
	}
	if len(_novaactWorkflowRunId) > 0 {
		input.WorkflowRunId = aws.String(_novaactWorkflowRunId)
	}
	if len(_novaactError) > 0 {
		if err := assignInputField(input, "Error", _novaactError); err != nil {
			log.Errorf("invalid --error: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateAct(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration or state of an active workflow run.
func novaact_UpdateWorkflowRun(cfg aws.Config, client *novaact.Client) {
	input := &novaact.UpdateWorkflowRunInput{
		// Status: types.WorkflowRunStatus, // Required
		// WorkflowDefinitionName: *string, // Required
		// WorkflowRunId: *string, // Required
	}

	if len(_novaactStatus) > 0 {
		if err := assignInputField(input, "Status", _novaactStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_novaactWorkflowDefinitionName) > 0 {
		input.WorkflowDefinitionName = aws.String(_novaactWorkflowDefinitionName)
	}
	if len(_novaactWorkflowRunId) > 0 {
		input.WorkflowRunId = aws.String(_novaactWorkflowRunId)
	}

	if resp, err := client.UpdateWorkflowRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_novaactCmd)
	_novaactCmd.Flags().SortFlags = false

	_novaactCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_novaactCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_novaactCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_novaactCmd.Flags().StringVarP(&_novaactActId, "act-id", "", "", "Act ID")
	_novaactCmd.Flags().StringVarP(&_novaactCallResults, "call-results", "", "", "Call Results")
	_novaactCmd.Flags().StringVarP(&_novaactClientCompatibilityVersion, "client-compatibility-version", "", "", "Client Compatibility Version")
	_novaactCmd.Flags().StringVarP(&_novaactClientInfo, "client-info", "", "", "Client Info")
	_novaactCmd.Flags().StringVarP(&_novaactClientToken, "client-token", "", "", "Client Token")
	_novaactCmd.Flags().StringVarP(&_novaactDescription, "description", "", "", "Description")
	_novaactCmd.Flags().StringVarP(&_novaactError, "error", "", "", "Error")
	_novaactCmd.Flags().StringVarP(&_novaactExportConfig, "export-config", "", "", "Export Config")
	_novaactCmd.Flags().StringVarP(&_novaactLogGroupName, "log-group-name", "", "", "Log Group Name")
	_novaactCmd.Flags().StringVarP(&_novaactMaxResults, "max-results", "", "", "Max Results")
	_novaactCmd.Flags().StringVarP(&_novaactModelId, "model-id", "", "", "Model ID")
	_novaactCmd.Flags().StringVarP(&_novaactName, "name", "", "", "Name")
	_novaactCmd.Flags().StringVarP(&_novaactNextToken, "next-token", "", "", "Next Token")
	_novaactCmd.Flags().StringVarP(&_novaactPreviousStepId, "previous-step-id", "", "", "Previous Step ID")
	_novaactCmd.Flags().StringVarP(&_novaactSessionId, "session-id", "", "", "Session ID")
	_novaactCmd.Flags().StringVarP(&_novaactSortOrder, "sort-order", "", "", "Sort Order")
	_novaactCmd.Flags().StringVarP(&_novaactStatus, "status", "", "", "Status")
	_novaactCmd.Flags().StringVarP(&_novaactTask, "task", "", "", "Task")
	_novaactCmd.Flags().StringVarP(&_novaactToolSpecs, "tool-specs", "", "", "Tool Specs")
	_novaactCmd.Flags().StringVarP(&_novaactWorkflowDefinitionName, "workflow-definition-name", "", "", "Workflow Definition Name")
	_novaactCmd.Flags().StringVarP(&_novaactWorkflowRunId, "workflow-run-id", "", "", "Workflow Run ID")

	_novaactCmd.Flags().BoolVarP(&_novaactCreateAct, "create-act", "", false, "Create Act")
	_novaactCmd.Flags().BoolVarP(&_novaactCreateSession, "create-session", "", false, "Create Session")
	_novaactCmd.Flags().BoolVarP(&_novaactCreateWorkflowDefinition, "create-workflow-definition", "", false, "Create Workflow Definition")
	_novaactCmd.Flags().BoolVarP(&_novaactCreateWorkflowRun, "create-workflow-run", "", false, "Create Workflow Run")
	_novaactCmd.Flags().BoolVarP(&_novaactDeleteWorkflowDefinition, "delete-workflow-definition", "", false, "Delete Workflow Definition")
	_novaactCmd.Flags().BoolVarP(&_novaactDeleteWorkflowRun, "delete-workflow-run", "", false, "Delete Workflow Run")
	_novaactCmd.Flags().BoolVarP(&_novaactGetWorkflowDefinition, "get-workflow-definition", "", false, "Get Workflow Definition")
	_novaactCmd.Flags().BoolVarP(&_novaactGetWorkflowRun, "get-workflow-run", "", false, "Get Workflow Run")
	_novaactCmd.Flags().BoolVarP(&_novaactInvokeActStep, "invoke-act-step", "", false, "Invoke Act Step")
	_novaactCmd.Flags().BoolVarP(&_novaactListActs, "list-acts", "", false, "List Acts")
	_novaactCmd.Flags().BoolVarP(&_novaactListModels, "list-models", "", false, "List Models")
	_novaactCmd.Flags().BoolVarP(&_novaactListSessions, "list-sessions", "", false, "List Sessions")
	_novaactCmd.Flags().BoolVarP(&_novaactListWorkflowDefinitions, "list-workflow-definitions", "", false, "List Workflow Definitions")
	_novaactCmd.Flags().BoolVarP(&_novaactListWorkflowRuns, "list-workflow-runs", "", false, "List Workflow Runs")
	_novaactCmd.Flags().BoolVarP(&_novaactUpdateAct, "update-act", "", false, "Update Act")
	_novaactCmd.Flags().BoolVarP(&_novaactUpdateWorkflowRun, "update-workflow-run", "", false, "Update Workflow Run")

}
