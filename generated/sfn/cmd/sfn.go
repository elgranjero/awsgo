package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sfn"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// sfnCmd represents the sfn command
var _sfnCmd = &cobra.Command{
	Use:   "sfn",
	Short: "AWS sfn CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := sfn.NewFromConfig(cfg)
		if _sfnCreateActivity {
			sfn_CreateActivity(cfg, client)
			return
		}
		if _sfnCreateStateMachine {
			sfn_CreateStateMachine(cfg, client)
			return
		}
		if _sfnCreateStateMachineAlias {
			sfn_CreateStateMachineAlias(cfg, client)
			return
		}
		if _sfnDeleteActivity {
			sfn_DeleteActivity(cfg, client)
			return
		}
		if _sfnDeleteStateMachine {
			sfn_DeleteStateMachine(cfg, client)
			return
		}
		if _sfnDeleteStateMachineAlias {
			sfn_DeleteStateMachineAlias(cfg, client)
			return
		}
		if _sfnDeleteStateMachineVersion {
			sfn_DeleteStateMachineVersion(cfg, client)
			return
		}
		if _sfnDescribeActivity {
			sfn_DescribeActivity(cfg, client)
			return
		}
		if _sfnDescribeExecution {
			sfn_DescribeExecution(cfg, client)
			return
		}
		if _sfnDescribeMapRun {
			sfn_DescribeMapRun(cfg, client)
			return
		}
		if _sfnDescribeStateMachine {
			sfn_DescribeStateMachine(cfg, client)
			return
		}
		if _sfnDescribeStateMachineAlias {
			sfn_DescribeStateMachineAlias(cfg, client)
			return
		}
		if _sfnDescribeStateMachineForExecution {
			sfn_DescribeStateMachineForExecution(cfg, client)
			return
		}
		if _sfnGetActivityTask {
			sfn_GetActivityTask(cfg, client)
			return
		}
		if _sfnGetExecutionHistory {
			sfn_GetExecutionHistory(cfg, client)
			return
		}
		if _sfnListActivities {
			sfn_ListActivities(cfg, client)
			return
		}
		if _sfnListExecutions {
			sfn_ListExecutions(cfg, client)
			return
		}
		if _sfnListMapRuns {
			sfn_ListMapRuns(cfg, client)
			return
		}
		if _sfnListStateMachineAliases {
			sfn_ListStateMachineAliases(cfg, client)
			return
		}
		if _sfnListStateMachineVersions {
			sfn_ListStateMachineVersions(cfg, client)
			return
		}
		if _sfnListStateMachines {
			sfn_ListStateMachines(cfg, client)
			return
		}
		if _sfnListTagsForResource {
			sfn_ListTagsForResource(cfg, client)
			return
		}
		if _sfnPublishStateMachineVersion {
			sfn_PublishStateMachineVersion(cfg, client)
			return
		}
		if _sfnRedriveExecution {
			sfn_RedriveExecution(cfg, client)
			return
		}
		if _sfnSendTaskFailure {
			sfn_SendTaskFailure(cfg, client)
			return
		}
		if _sfnSendTaskHeartbeat {
			sfn_SendTaskHeartbeat(cfg, client)
			return
		}
		if _sfnSendTaskSuccess {
			sfn_SendTaskSuccess(cfg, client)
			return
		}
		if _sfnStartExecution {
			sfn_StartExecution(cfg, client)
			return
		}
		if _sfnStartSyncExecution {
			sfn_StartSyncExecution(cfg, client)
			return
		}
		if _sfnStopExecution {
			sfn_StopExecution(cfg, client)
			return
		}
		if _sfnTagResource {
			sfn_TagResource(cfg, client)
			return
		}
		if _sfnTestState {
			sfn_TestState(cfg, client)
			return
		}
		if _sfnUntagResource {
			sfn_UntagResource(cfg, client)
			return
		}
		if _sfnUpdateMapRun {
			sfn_UpdateMapRun(cfg, client)
			return
		}
		if _sfnUpdateStateMachine {
			sfn_UpdateStateMachine(cfg, client)
			return
		}
		if _sfnUpdateStateMachineAlias {
			sfn_UpdateStateMachineAlias(cfg, client)
			return
		}
		if _sfnValidateStateMachineDefinition {
			sfn_ValidateStateMachineDefinition(cfg, client)
			return
		}

	},
}

var (
	_sfnCreateActivity                   bool
	_sfnCreateStateMachine               bool
	_sfnCreateStateMachineAlias          bool
	_sfnDeleteActivity                   bool
	_sfnDeleteStateMachine               bool
	_sfnDeleteStateMachineAlias          bool
	_sfnDeleteStateMachineVersion        bool
	_sfnDescribeActivity                 bool
	_sfnDescribeExecution                bool
	_sfnDescribeMapRun                   bool
	_sfnDescribeStateMachine             bool
	_sfnDescribeStateMachineAlias        bool
	_sfnDescribeStateMachineForExecution bool
	_sfnGetActivityTask                  bool
	_sfnGetExecutionHistory              bool
	_sfnListActivities                   bool
	_sfnListExecutions                   bool
	_sfnListMapRuns                      bool
	_sfnListStateMachineAliases          bool
	_sfnListStateMachineVersions         bool
	_sfnListStateMachines                bool
	_sfnListTagsForResource              bool
	_sfnPublishStateMachineVersion       bool
	_sfnRedriveExecution                 bool
	_sfnSendTaskFailure                  bool
	_sfnSendTaskHeartbeat                bool
	_sfnSendTaskSuccess                  bool
	_sfnStartExecution                   bool
	_sfnStartSyncExecution               bool
	_sfnStopExecution                    bool
	_sfnTagResource                      bool
	_sfnTestState                        bool
	_sfnUntagResource                    bool
	_sfnUpdateMapRun                     bool
	_sfnUpdateStateMachine               bool
	_sfnUpdateStateMachineAlias          bool
	_sfnValidateStateMachineDefinition   bool

	_sfnActivityArn                string
	_sfnCause                      string
	_sfnClientToken                string
	_sfnContext                    string
	_sfnDefinition                 string
	_sfnDescription                string
	_sfnEncryptionConfiguration    string
	_sfnError                      string
	_sfnExecutionArn               string
	_sfnIncludeExecutionData       string
	_sfnIncludedData               string
	_sfnInput                      string
	_sfnInspectionLevel            string
	_sfnLoggingConfiguration       string
	_sfnMapRunArn                  string
	_sfnMaxConcurrency             string
	_sfnMaxResults                 string
	_sfnMock                       string
	_sfnName                       string
	_sfnNextToken                  string
	_sfnPublish                    string
	_sfnRedriveFilter              string
	_sfnResourceArn                string
	_sfnRevealSecrets              string
	_sfnReverseOrder               string
	_sfnRevisionId                 string
	_sfnRoleArn                    string
	_sfnRoutingConfiguration       string
	_sfnSeverity                   string
	_sfnStateConfiguration         string
	_sfnStateMachineAliasArn       string
	_sfnStateMachineArn            string
	_sfnStateMachineVersionArn     string
	_sfnStateName                  string
	_sfnStatusFilter               string
	_sfnTagKeys                    []string
	_sfnTags                       string
	_sfnTaskToken                  string
	_sfnToleratedFailureCount      string
	_sfnToleratedFailurePercentage string
	_sfnTraceHeader                string
	_sfnTracingConfiguration       string
	_sfnType                       string
	_sfnVariables                  string
	_sfnVersionDescription         string
	_sfnWorkerName                 string
)

// Creates an activity. An activity is a task that you write in any programming
// language and host on any machine that has access to Step Functions. Activities
// must poll Step Functions using the GetActivityTask API action and respond using
// SendTask* API actions. This function lets Step Functions know the existence of
// your activity and returns an identifier for use in a state machine and when
// polling from the activity.
//
// This operation is eventually consistent. The results are best effort and may
// not reflect very recent updates and changes.
//
// CreateActivity is an idempotent API. Subsequent requests won’t create a
// duplicate resource if it was already created. CreateActivity 's idempotency
// check is based on the activity name . If a following request has different tags
// values, Step Functions will ignore these differences and treat it as an
// idempotent request of the previous. In this case, tags will not be updated,
// even if they are different.
func sfn_CreateActivity(cfg aws.Config, client *sfn.Client) {
	input := &sfn.CreateActivityInput{
		// Name: *string, // Required
	}

	if len(_sfnName) > 0 {
		input.Name = aws.String(_sfnName)
	}
	if len(_sfnEncryptionConfiguration) > 0 {
		if err := assignInputField(input, "EncryptionConfiguration", _sfnEncryptionConfiguration); err != nil {
			log.Errorf("invalid --encryption-configuration: %s", err.Error())
			return
		}
	}
	if len(_sfnTags) > 0 {
		if err := assignInputField(input, "Tags", _sfnTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateActivity(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a state machine. A state machine consists of a collection of states
// that can do work ( Task states), determine to which states to transition next (
// Choice states), stop an execution with an error ( Fail states), and so on.
// State machines are specified using a JSON-based, structured language. For more
// information, see [Amazon States Language]in the Step Functions User Guide.
//
// If you set the publish parameter of this API action to true , it publishes
// version 1 as the first revision of the state machine.
//
// For additional control over security, you can encrypt your data using a
// customer-managed key for Step Functions state machines. You can configure a
// symmetric KMS key and data key reuse period when creating or updating a State
// Machine. The execution history and state machine definition will be encrypted
// with the key applied to the State Machine.
//
// This operation is eventually consistent. The results are best effort and may
// not reflect very recent updates and changes.
//
// CreateStateMachine is an idempotent API. Subsequent requests won’t create a
// duplicate resource if it was already created. CreateStateMachine 's idempotency
// check is based on the state machine name , definition , type ,
// LoggingConfiguration , TracingConfiguration , and EncryptionConfiguration The
// check is also based on the publish and versionDescription parameters. If a
// following request has a different roleArn or tags , Step Functions will ignore
// these differences and treat it as an idempotent request of the previous. In this
// case, roleArn and tags will not be updated, even if they are different.
//
// [Amazon States Language]: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-amazon-states-language.html
func sfn_CreateStateMachine(cfg aws.Config, client *sfn.Client) {
	input := &sfn.CreateStateMachineInput{
		// Definition: *string, // Required
		// Name: *string, // Required
		// RoleArn: *string, // Required
	}

	if len(_sfnDefinition) > 0 {
		input.Definition = aws.String(_sfnDefinition)
	}
	if len(_sfnName) > 0 {
		input.Name = aws.String(_sfnName)
	}
	if len(_sfnRoleArn) > 0 {
		input.RoleArn = aws.String(_sfnRoleArn)
	}
	if len(_sfnEncryptionConfiguration) > 0 {
		if err := assignInputField(input, "EncryptionConfiguration", _sfnEncryptionConfiguration); err != nil {
			log.Errorf("invalid --encryption-configuration: %s", err.Error())
			return
		}
	}
	if len(_sfnLoggingConfiguration) > 0 {
		if err := assignInputField(input, "LoggingConfiguration", _sfnLoggingConfiguration); err != nil {
			log.Errorf("invalid --logging-configuration: %s", err.Error())
			return
		}
	}
	if len(_sfnPublish) > 0 {
		if err := assignInputField(input, "Publish", _sfnPublish); err != nil {
			log.Errorf("invalid --publish: %s", err.Error())
			return
		}
	}
	if len(_sfnTags) > 0 {
		if err := assignInputField(input, "Tags", _sfnTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_sfnTracingConfiguration) > 0 {
		if err := assignInputField(input, "TracingConfiguration", _sfnTracingConfiguration); err != nil {
			log.Errorf("invalid --tracing-configuration: %s", err.Error())
			return
		}
	}
	if len(_sfnType) > 0 {
		if err := assignInputField(input, "Type", _sfnType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_sfnVersionDescription) > 0 {
		input.VersionDescription = aws.String(_sfnVersionDescription)
	}

	if resp, err := client.CreateStateMachine(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an [alias] for a state machine that points to one or two [versions] of the same state
// machine. You can set your application to call StartExecutionwith an alias and update the
// version the alias uses without changing the client's code.
//
// You can also map an alias to split StartExecution requests between two versions of a state
// machine. To do this, add a second RoutingConfig object in the
// routingConfiguration parameter. You must also specify the percentage of
// execution run requests each version should receive in both RoutingConfig
// objects. Step Functions randomly chooses which version runs a given execution
// based on the percentage you specify.
//
// To create an alias that points to a single version, specify a single
// RoutingConfig object with a weight set to 100.
//
// You can create up to 100 aliases for each state machine. You must delete unused
// aliases using the DeleteStateMachineAliasAPI action.
//
// CreateStateMachineAlias is an idempotent API. Step Functions bases the
// idempotency check on the stateMachineArn , description , name , and
// routingConfiguration parameters. Requests that contain the same values for these
// parameters return a successful idempotent response without creating a duplicate
// resource.
//
// Related operations:
//
// # DescribeStateMachineAlias
//
// # ListStateMachineAliases
//
// # UpdateStateMachineAlias
//
// # DeleteStateMachineAlias
//
// [versions]: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-state-machine-version.html
// [alias]: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-state-machine-alias.html
func sfn_CreateStateMachineAlias(cfg aws.Config, client *sfn.Client) {
	input := &sfn.CreateStateMachineAliasInput{
		// Name: *string, // Required
		// RoutingConfiguration: []types.RoutingConfigurationListItem, // Required
	}

	if len(_sfnName) > 0 {
		input.Name = aws.String(_sfnName)
	}
	if len(_sfnRoutingConfiguration) > 0 {
		if err := assignInputField(input, "RoutingConfiguration", _sfnRoutingConfiguration); err != nil {
			log.Errorf("invalid --routing-configuration: %s", err.Error())
			return
		}
	}
	if len(_sfnDescription) > 0 {
		input.Description = aws.String(_sfnDescription)
	}

	if resp, err := client.CreateStateMachineAlias(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an activity.
func sfn_DeleteActivity(cfg aws.Config, client *sfn.Client) {
	input := &sfn.DeleteActivityInput{
		// ActivityArn: *string, // Required
	}

	if len(_sfnActivityArn) > 0 {
		input.ActivityArn = aws.String(_sfnActivityArn)
	}

	if resp, err := client.DeleteActivity(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a state machine. This is an asynchronous operation. It sets the state
// machine's status to DELETING and begins the deletion process. A state machine
// is deleted only when all its executions are completed. On the next state
// transition, the state machine's executions are terminated.
//
// A qualified state machine ARN can either refer to a Distributed Map state
// defined within a state machine, a version ARN, or an alias ARN.
//
// The following are some examples of qualified and unqualified state machine ARNs:
//
// - The following qualified state machine ARN refers to a Distributed Map state
// with a label mapStateLabel in a state machine named myStateMachine .
//
// arn:partition:states:region:account-id:stateMachine:myStateMachine/mapStateLabel
//
// # If you provide a qualified state machine ARN that refers to a Distributed Map
//
// state, the request fails with ValidationException .
//
// - The following unqualified state machine ARN refers to a state machine named
// myStateMachine .
//
// arn:partition:states:region:account-id:stateMachine:myStateMachine
//
// This API action also deletes all [versions] and [aliases] associated with a state machine.
//
// For EXPRESS state machines, the deletion happens eventually (usually in less
// than a minute). Running executions may emit logs after DeleteStateMachine API
// is called.
//
// [aliases]: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-state-machine-alias.html
// [versions]: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-state-machine-version.html
func sfn_DeleteStateMachine(cfg aws.Config, client *sfn.Client) {
	input := &sfn.DeleteStateMachineInput{
		// StateMachineArn: *string, // Required
	}

	if len(_sfnStateMachineArn) > 0 {
		input.StateMachineArn = aws.String(_sfnStateMachineArn)
	}

	if resp, err := client.DeleteStateMachine(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a state machine [alias].
// After you delete a state machine alias, you can't use it to start executions.
// When you delete a state machine alias, Step Functions doesn't delete the state
// machine versions that alias references.
//
// Related operations:
//
// # CreateStateMachineAlias
//
// # DescribeStateMachineAlias
//
// # ListStateMachineAliases
//
// # UpdateStateMachineAlias
//
// [alias]: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-state-machine-alias.html
func sfn_DeleteStateMachineAlias(cfg aws.Config, client *sfn.Client) {
	input := &sfn.DeleteStateMachineAliasInput{
		// StateMachineAliasArn: *string, // Required
	}

	if len(_sfnStateMachineAliasArn) > 0 {
		input.StateMachineAliasArn = aws.String(_sfnStateMachineAliasArn)
	}

	if resp, err := client.DeleteStateMachineAlias(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a state machine [version]. After you delete a version, you can't call StartExecution using
// that version's ARN or use the version with a state machine [alias].
//
// Deleting a state machine version won't terminate its in-progress executions.
//
// You can't delete a state machine version currently referenced by one or more
// aliases. Before you delete a version, you must either delete the aliases or
// update them to point to another state machine version.
//
// Related operations:
//
// # PublishStateMachineVersion
//
// # ListStateMachineVersions
//
// [alias]: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-state-machine-alias.html
// [version]: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-state-machine-version.html
func sfn_DeleteStateMachineVersion(cfg aws.Config, client *sfn.Client) {
	input := &sfn.DeleteStateMachineVersionInput{
		// StateMachineVersionArn: *string, // Required
	}

	if len(_sfnStateMachineVersionArn) > 0 {
		input.StateMachineVersionArn = aws.String(_sfnStateMachineVersionArn)
	}

	if resp, err := client.DeleteStateMachineVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes an activity.
// This operation is eventually consistent. The results are best effort and may
// not reflect very recent updates and changes.
func sfn_DescribeActivity(cfg aws.Config, client *sfn.Client) {
	input := &sfn.DescribeActivityInput{
		// ActivityArn: *string, // Required
	}

	if len(_sfnActivityArn) > 0 {
		input.ActivityArn = aws.String(_sfnActivityArn)
	}

	if resp, err := client.DescribeActivity(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides information about a state machine execution, such as the state machine
// associated with the execution, the execution input and output, and relevant
// execution metadata. If you've [redriven]an execution, you can use this API action to
// return information about the redrives of that execution. In addition, you can
// use this API action to return the Map Run Amazon Resource Name (ARN) if the
// execution was dispatched by a Map Run.
//
// If you specify a version or alias ARN when you call the StartExecution API action,
// DescribeExecution returns that ARN.
//
// This operation is eventually consistent. The results are best effort and may
// not reflect very recent updates and changes.
//
// Executions of an EXPRESS state machine aren't supported by DescribeExecution
// unless a Map Run dispatched them.
//
// [redriven]: https://docs.aws.amazon.com/step-functions/latest/dg/redrive-executions.html
func sfn_DescribeExecution(cfg aws.Config, client *sfn.Client) {
	input := &sfn.DescribeExecutionInput{
		// ExecutionArn: *string, // Required
	}

	if len(_sfnExecutionArn) > 0 {
		input.ExecutionArn = aws.String(_sfnExecutionArn)
	}
	if len(_sfnIncludedData) > 0 {
		if err := assignInputField(input, "IncludedData", _sfnIncludedData); err != nil {
			log.Errorf("invalid --included-data: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides information about a Map Run's configuration, progress, and results. If
// you've [redriven]a Map Run, this API action also returns information about the redrives
// of that Map Run. For more information, see [Examining Map Run]in the Step Functions Developer
// Guide.
//
// [redriven]: https://docs.aws.amazon.com/step-functions/latest/dg/redrive-map-run.html
// [Examining Map Run]: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-examine-map-run.html
func sfn_DescribeMapRun(cfg aws.Config, client *sfn.Client) {
	input := &sfn.DescribeMapRunInput{
		// MapRunArn: *string, // Required
	}

	if len(_sfnMapRunArn) > 0 {
		input.MapRunArn = aws.String(_sfnMapRunArn)
	}

	if resp, err := client.DescribeMapRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides information about a state machine's definition, its IAM role Amazon
// Resource Name (ARN), and configuration.
//
// A qualified state machine ARN can either refer to a Distributed Map state
// defined within a state machine, a version ARN, or an alias ARN.
//
// The following are some examples of qualified and unqualified state machine ARNs:
//
// - The following qualified state machine ARN refers to a Distributed Map state
// with a label mapStateLabel in a state machine named myStateMachine .
//
// arn:partition:states:region:account-id:stateMachine:myStateMachine/mapStateLabel
//
// # If you provide a qualified state machine ARN that refers to a Distributed Map
//
// state, the request fails with ValidationException .
//
// - The following qualified state machine ARN refers to an alias named PROD .
//
// arn::states:::stateMachine:
//
// # If you provide a qualified state machine ARN that refers to a version ARN or an
//
// alias ARN, the request starts execution for that version or alias.
//
// - The following unqualified state machine ARN refers to a state machine named
// myStateMachine .
//
// arn::states:::stateMachine:
//
// This API action returns the details for a state machine version if the
// stateMachineArn you specify is a state machine version ARN.
//
// This operation is eventually consistent. The results are best effort and may
// not reflect very recent updates and changes.
func sfn_DescribeStateMachine(cfg aws.Config, client *sfn.Client) {
	input := &sfn.DescribeStateMachineInput{
		// StateMachineArn: *string, // Required
	}

	if len(_sfnStateMachineArn) > 0 {
		input.StateMachineArn = aws.String(_sfnStateMachineArn)
	}
	if len(_sfnIncludedData) > 0 {
		if err := assignInputField(input, "IncludedData", _sfnIncludedData); err != nil {
			log.Errorf("invalid --included-data: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeStateMachine(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns details about a state machine [alias].
// Related operations:
//
// # CreateStateMachineAlias
//
// # ListStateMachineAliases
//
// # UpdateStateMachineAlias
//
// # DeleteStateMachineAlias
//
// [alias]: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-state-machine-alias.html
func sfn_DescribeStateMachineAlias(cfg aws.Config, client *sfn.Client) {
	input := &sfn.DescribeStateMachineAliasInput{
		// StateMachineAliasArn: *string, // Required
	}

	if len(_sfnStateMachineAliasArn) > 0 {
		input.StateMachineAliasArn = aws.String(_sfnStateMachineAliasArn)
	}

	if resp, err := client.DescribeStateMachineAlias(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides information about a state machine's definition, its execution role
// ARN, and configuration. If a Map Run dispatched the execution, this action
// returns the Map Run Amazon Resource Name (ARN) in the response. The state
// machine returned is the state machine associated with the Map Run.
//
// This operation is eventually consistent. The results are best effort and may
// not reflect very recent updates and changes.
//
// This API action is not supported by EXPRESS state machines.
func sfn_DescribeStateMachineForExecution(cfg aws.Config, client *sfn.Client) {
	input := &sfn.DescribeStateMachineForExecutionInput{
		// ExecutionArn: *string, // Required
	}

	if len(_sfnExecutionArn) > 0 {
		input.ExecutionArn = aws.String(_sfnExecutionArn)
	}
	if len(_sfnIncludedData) > 0 {
		if err := assignInputField(input, "IncludedData", _sfnIncludedData); err != nil {
			log.Errorf("invalid --included-data: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeStateMachineForExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Used by workers to retrieve a task (with the specified activity ARN) which has
// been scheduled for execution by a running state machine. This initiates a long
// poll, where the service holds the HTTP connection open and responds as soon as a
// task becomes available (i.e. an execution of a task of this type is needed.) The
// maximum time the service holds on to the request before responding is 60
// seconds. If no task is available within 60 seconds, the poll returns a taskToken
// with a null string.
//
// This API action isn't logged in CloudTrail.
//
// Workers should set their client side socket timeout to at least 65 seconds (5
// seconds higher than the maximum time the service may hold the poll request).
//
// Polling with GetActivityTask can cause latency in some implementations. See [Avoid Latency When Polling for Activity Tasks] in
// the Step Functions Developer Guide.
//
// [Avoid Latency When Polling for Activity Tasks]: https://docs.aws.amazon.com/step-functions/latest/dg/bp-activity-pollers.html
func sfn_GetActivityTask(cfg aws.Config, client *sfn.Client) {
	input := &sfn.GetActivityTaskInput{
		// ActivityArn: *string, // Required
	}

	if len(_sfnActivityArn) > 0 {
		input.ActivityArn = aws.String(_sfnActivityArn)
	}
	if len(_sfnWorkerName) > 0 {
		input.WorkerName = aws.String(_sfnWorkerName)
	}

	if resp, err := client.GetActivityTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the history of the specified execution as a list of events. By default,
// the results are returned in ascending order of the timeStamp of the events. Use
// the reverseOrder parameter to get the latest events first.
//
// If nextToken is returned, there are more results available. The value of
// nextToken is a unique pagination token for each page. Make the call again using
// the returned token to retrieve the next page. Keep all other arguments
// unchanged. Each pagination token expires after 24 hours. Using an expired
// pagination token will return an HTTP 400 InvalidToken error.
//
// This API action is not supported by EXPRESS state machines.
func sfn_GetExecutionHistory(cfg aws.Config, client *sfn.Client) {
	input := &sfn.GetExecutionHistoryInput{
		// ExecutionArn: *string, // Required
	}

	if len(_sfnExecutionArn) > 0 {
		input.ExecutionArn = aws.String(_sfnExecutionArn)
	}
	if len(_sfnIncludeExecutionData) > 0 {
		if err := assignInputField(input, "IncludeExecutionData", _sfnIncludeExecutionData); err != nil {
			log.Errorf("invalid --include-execution-data: %s", err.Error())
			return
		}
	}
	if len(_sfnMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sfnMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sfnNextToken) > 0 {
		input.NextToken = aws.String(_sfnNextToken)
	}
	if len(_sfnReverseOrder) > 0 {
		if err := assignInputField(input, "ReverseOrder", _sfnReverseOrder); err != nil {
			log.Errorf("invalid --reverse-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.GetExecutionHistory(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sfn.GetExecutionHistoryOutput
	p := sfn.NewGetExecutionHistoryPaginator(client, input)
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

// Lists the existing activities.
// If nextToken is returned, there are more results available. The value of
// nextToken is a unique pagination token for each page. Make the call again using
// the returned token to retrieve the next page. Keep all other arguments
// unchanged. Each pagination token expires after 24 hours. Using an expired
// pagination token will return an HTTP 400 InvalidToken error.
//
// This operation is eventually consistent. The results are best effort and may
// not reflect very recent updates and changes.
func sfn_ListActivities(cfg aws.Config, client *sfn.Client) {
	input := &sfn.ListActivitiesInput{}

	if len(_sfnMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sfnMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sfnNextToken) > 0 {
		input.NextToken = aws.String(_sfnNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListActivities(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sfn.ListActivitiesOutput
	p := sfn.NewListActivitiesPaginator(client, input)
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

// Lists all executions of a state machine or a Map Run. You can list all
// executions related to a state machine by specifying a state machine Amazon
// Resource Name (ARN), or those related to a Map Run by specifying a Map Run ARN.
// Using this API action, you can also list all [redriven]executions.
//
// You can also provide a state machine [alias] ARN or [version] ARN to list the executions
// associated with a specific alias or version.
//
// Results are sorted by time, with the most recent execution first.
//
// If nextToken is returned, there are more results available. The value of
// nextToken is a unique pagination token for each page. Make the call again using
// the returned token to retrieve the next page. Keep all other arguments
// unchanged. Each pagination token expires after 24 hours. Using an expired
// pagination token will return an HTTP 400 InvalidToken error.
//
// This operation is eventually consistent. The results are best effort and may
// not reflect very recent updates and changes.
//
// This API action is not supported by EXPRESS state machines.
//
// [redriven]: https://docs.aws.amazon.com/step-functions/latest/dg/redrive-executions.html
// [alias]: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-state-machine-alias.html
// [version]: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-state-machine-version.html
func sfn_ListExecutions(cfg aws.Config, client *sfn.Client) {
	input := &sfn.ListExecutionsInput{}

	if len(_sfnMapRunArn) > 0 {
		input.MapRunArn = aws.String(_sfnMapRunArn)
	}
	if len(_sfnMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sfnMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sfnNextToken) > 0 {
		input.NextToken = aws.String(_sfnNextToken)
	}
	if len(_sfnRedriveFilter) > 0 {
		if err := assignInputField(input, "RedriveFilter", _sfnRedriveFilter); err != nil {
			log.Errorf("invalid --redrive-filter: %s", err.Error())
			return
		}
	}
	if len(_sfnStateMachineArn) > 0 {
		input.StateMachineArn = aws.String(_sfnStateMachineArn)
	}
	if len(_sfnStatusFilter) > 0 {
		if err := assignInputField(input, "StatusFilter", _sfnStatusFilter); err != nil {
			log.Errorf("invalid --status-filter: %s", err.Error())
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

	var results []*sfn.ListExecutionsOutput
	p := sfn.NewListExecutionsPaginator(client, input)
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

// Lists all Map Runs that were started by a given state machine execution. Use
// this API action to obtain Map Run ARNs, and then call DescribeMapRun to obtain
// more information, if needed.
func sfn_ListMapRuns(cfg aws.Config, client *sfn.Client) {
	input := &sfn.ListMapRunsInput{
		// ExecutionArn: *string, // Required
	}

	if len(_sfnExecutionArn) > 0 {
		input.ExecutionArn = aws.String(_sfnExecutionArn)
	}
	if len(_sfnMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sfnMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sfnNextToken) > 0 {
		input.NextToken = aws.String(_sfnNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListMapRuns(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sfn.ListMapRunsOutput
	p := sfn.NewListMapRunsPaginator(client, input)
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

// Lists [aliases] for a specified state machine ARN. Results are sorted by time, with the
// most recently created aliases listed first.
//
// To list aliases that reference a state machine [version], you can specify the version
// ARN in the stateMachineArn parameter.
//
// If nextToken is returned, there are more results available. The value of
// nextToken is a unique pagination token for each page. Make the call again using
// the returned token to retrieve the next page. Keep all other arguments
// unchanged. Each pagination token expires after 24 hours. Using an expired
// pagination token will return an HTTP 400 InvalidToken error.
//
// Related operations:
//
// # CreateStateMachineAlias
//
// # DescribeStateMachineAlias
//
// # UpdateStateMachineAlias
//
// # DeleteStateMachineAlias
//
// [version]: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-state-machine-version.html
// [aliases]: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-state-machine-alias.html
func sfn_ListStateMachineAliases(cfg aws.Config, client *sfn.Client) {
	input := &sfn.ListStateMachineAliasesInput{
		// StateMachineArn: *string, // Required
	}

	if len(_sfnStateMachineArn) > 0 {
		input.StateMachineArn = aws.String(_sfnStateMachineArn)
	}
	if len(_sfnMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sfnMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sfnNextToken) > 0 {
		input.NextToken = aws.String(_sfnNextToken)
	}

	if resp, err := client.ListStateMachineAliases(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists [versions] for the specified state machine Amazon Resource Name (ARN).
// The results are sorted in descending order of the version creation time.
//
// If nextToken is returned, there are more results available. The value of
// nextToken is a unique pagination token for each page. Make the call again using
// the returned token to retrieve the next page. Keep all other arguments
// unchanged. Each pagination token expires after 24 hours. Using an expired
// pagination token will return an HTTP 400 InvalidToken error.
//
// Related operations:
//
// # PublishStateMachineVersion
//
// # DeleteStateMachineVersion
//
// [versions]: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-state-machine-version.html
func sfn_ListStateMachineVersions(cfg aws.Config, client *sfn.Client) {
	input := &sfn.ListStateMachineVersionsInput{
		// StateMachineArn: *string, // Required
	}

	if len(_sfnStateMachineArn) > 0 {
		input.StateMachineArn = aws.String(_sfnStateMachineArn)
	}
	if len(_sfnMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sfnMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sfnNextToken) > 0 {
		input.NextToken = aws.String(_sfnNextToken)
	}

	if resp, err := client.ListStateMachineVersions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the existing state machines.
// If nextToken is returned, there are more results available. The value of
// nextToken is a unique pagination token for each page. Make the call again using
// the returned token to retrieve the next page. Keep all other arguments
// unchanged. Each pagination token expires after 24 hours. Using an expired
// pagination token will return an HTTP 400 InvalidToken error.
//
// This operation is eventually consistent. The results are best effort and may
// not reflect very recent updates and changes.
func sfn_ListStateMachines(cfg aws.Config, client *sfn.Client) {
	input := &sfn.ListStateMachinesInput{}

	if len(_sfnMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sfnMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sfnNextToken) > 0 {
		input.NextToken = aws.String(_sfnNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListStateMachines(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sfn.ListStateMachinesOutput
	p := sfn.NewListStateMachinesPaginator(client, input)
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

// List tags for a given resource.
// Tags may only contain Unicode letters, digits, white space, or these symbols: _
// . : / = + - (at) .
func sfn_ListTagsForResource(cfg aws.Config, client *sfn.Client) {
	input := &sfn.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_sfnResourceArn) > 0 {
		input.ResourceArn = aws.String(_sfnResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a [version] from the current revision of a state machine. Use versions to create
// immutable snapshots of your state machine. You can start executions from
// versions either directly or with an alias. To create an alias, use CreateStateMachineAlias.
//
// You can publish up to 1000 versions for each state machine. You must manually
// delete unused versions using the DeleteStateMachineVersionAPI action.
//
// PublishStateMachineVersion is an idempotent API. It doesn't create a duplicate
// state machine version if it already exists for the current revision. Step
// Functions bases PublishStateMachineVersion 's idempotency check on the
// stateMachineArn , name , and revisionId parameters. Requests with the same
// parameters return a successful idempotent response. If you don't specify a
// revisionId , Step Functions checks for a previously published version of the
// state machine's current revision.
//
// Related operations:
//
// # DeleteStateMachineVersion
//
// # ListStateMachineVersions
//
// [version]: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-state-machine-version.html
func sfn_PublishStateMachineVersion(cfg aws.Config, client *sfn.Client) {
	input := &sfn.PublishStateMachineVersionInput{
		// StateMachineArn: *string, // Required
	}

	if len(_sfnStateMachineArn) > 0 {
		input.StateMachineArn = aws.String(_sfnStateMachineArn)
	}
	if len(_sfnDescription) > 0 {
		input.Description = aws.String(_sfnDescription)
	}
	if len(_sfnRevisionId) > 0 {
		input.RevisionId = aws.String(_sfnRevisionId)
	}

	if resp, err := client.PublishStateMachineVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Restarts unsuccessful executions of Standard workflows that didn't complete
// successfully in the last 14 days. These include failed, aborted, or timed out
// executions. When you [redrive]an execution, it continues the failed execution from the
// unsuccessful step and uses the same input. Step Functions preserves the results
// and execution history of the successful steps, and doesn't rerun these steps
// when you redrive an execution. Redriven executions use the same state machine
// definition and execution ARN as the original execution attempt.
//
// For workflows that include an [Inline Map] or [Parallel] state, RedriveExecution API action
// reschedules and redrives only the iterations and branches that failed or
// aborted.
//
// To redrive a workflow that includes a Distributed Map state whose Map Run
// failed, you must redrive the [parent workflow]. The parent workflow redrives all the
// unsuccessful states, including a failed Map Run. If a Map Run was not started in
// the original execution attempt, the redriven parent workflow starts the Map Run.
//
// This API action is not supported by EXPRESS state machines.
//
// However, you can restart the unsuccessful executions of Express child workflows
// in a Distributed Map by redriving its Map Run. When you redrive a Map Run, the
// Express child workflows are rerun using the StartExecutionAPI action. For more information,
// see [Redriving Map Runs].
//
// You can redrive executions if your original execution meets the following
// conditions:
//
// - The execution status isn't SUCCEEDED .
//
// - Your workflow execution has not exceeded the redrivable period of 14 days.
// Redrivable period refers to the time during which you can redrive a given
// execution. This period starts from the day a state machine completes its
// execution.
//
// - The workflow execution has not exceeded the maximum open time of one year.
// For more information about state machine quotas, see [Quotas related to state machine executions].
//
// - The execution event history count is less than 24,999. Redriven executions
// append their event history to the existing event history. Make sure your
// workflow execution contains less than 24,999 events to accommodate the
// ExecutionRedriven history event and at least one other history event.
//
// [redrive]: https://docs.aws.amazon.com/step-functions/latest/dg/redrive-executions.html
// [parent workflow]: https://docs.aws.amazon.com/step-functions/latest/dg/use-dist-map-orchestrate-large-scale-parallel-workloads.html#dist-map-orchestrate-parallel-workloads-key-terms
// [Quotas related to state machine executions]: https://docs.aws.amazon.com/step-functions/latest/dg/limits-overview.html#service-limits-state-machine-executions
// [Parallel]: https://docs.aws.amazon.com/step-functions/latest/dg/amazon-states-language-parallel-state.html
// [Inline Map]: https://docs.aws.amazon.com/step-functions/latest/dg/amazon-states-language-map-state.html
// [Redriving Map Runs]: https://docs.aws.amazon.com/step-functions/latest/dg/redrive-map-run.html
func sfn_RedriveExecution(cfg aws.Config, client *sfn.Client) {
	input := &sfn.RedriveExecutionInput{
		// ExecutionArn: *string, // Required
	}

	if len(_sfnExecutionArn) > 0 {
		input.ExecutionArn = aws.String(_sfnExecutionArn)
	}
	if len(_sfnClientToken) > 0 {
		input.ClientToken = aws.String(_sfnClientToken)
	}

	if resp, err := client.RedriveExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Used by activity workers, Task states using the [callback] pattern, and optionally Task
// states using the [job run]pattern to report that the task identified by the taskToken
// failed.
//
// For an execution with encryption enabled, Step Functions will encrypt the error
// and cause fields using the KMS key for the execution role.
//
// A caller can mark a task as fail without using any KMS permissions in the
// execution role if the caller provides a null value for both error and cause
// fields because no data needs to be encrypted.
//
// [job run]: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-sync
// [callback]: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token
func sfn_SendTaskFailure(cfg aws.Config, client *sfn.Client) {
	input := &sfn.SendTaskFailureInput{
		// TaskToken: *string, // Required
	}

	if len(_sfnTaskToken) > 0 {
		input.TaskToken = aws.String(_sfnTaskToken)
	}
	if len(_sfnCause) > 0 {
		input.Cause = aws.String(_sfnCause)
	}
	if len(_sfnError) > 0 {
		input.Error = aws.String(_sfnError)
	}

	if resp, err := client.SendTaskFailure(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Used by activity workers and Task states using the [callback] pattern, and optionally
// Task states using the [job run]pattern to report to Step Functions that the task
// represented by the specified taskToken is still making progress. This action
// resets the Heartbeat clock. The Heartbeat threshold is specified in the state
// machine's Amazon States Language definition ( HeartbeatSeconds ). This action
// does not in itself create an event in the execution history. However, if the
// task times out, the execution history contains an ActivityTimedOut entry for
// activities, or a TaskTimedOut entry for tasks using the [job run] or [callback] pattern.
//
// The Timeout of a task, defined in the state machine's Amazon States Language
// definition, is its maximum allowed duration, regardless of the number of SendTaskHeartbeat
// requests received. Use HeartbeatSeconds to configure the timeout interval for
// heartbeats.
//
// [callback]: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token
// [job run]: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-sync
func sfn_SendTaskHeartbeat(cfg aws.Config, client *sfn.Client) {
	input := &sfn.SendTaskHeartbeatInput{
		// TaskToken: *string, // Required
	}

	if len(_sfnTaskToken) > 0 {
		input.TaskToken = aws.String(_sfnTaskToken)
	}

	if resp, err := client.SendTaskHeartbeat(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Used by activity workers, Task states using the [callback] pattern, and optionally Task
// states using the [job run]pattern to report that the task identified by the taskToken
// completed successfully.
//
// [job run]: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-sync
// [callback]: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token
func sfn_SendTaskSuccess(cfg aws.Config, client *sfn.Client) {
	input := &sfn.SendTaskSuccessInput{
		// Output: *string, // Required
		// TaskToken: *string, // Required
	}

	if len(_sfnTaskToken) > 0 {
		input.TaskToken = aws.String(_sfnTaskToken)
	}

	if resp, err := client.SendTaskSuccess(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a state machine execution.
// A qualified state machine ARN can either refer to a Distributed Map state
// defined within a state machine, a version ARN, or an alias ARN.
//
// The following are some examples of qualified and unqualified state machine ARNs:
//
// - The following qualified state machine ARN refers to a Distributed Map state
// with a label mapStateLabel in a state machine named myStateMachine .
//
// arn:partition:states:region:account-id:stateMachine:myStateMachine/mapStateLabel
//
// # If you provide a qualified state machine ARN that refers to a Distributed Map
//
// state, the request fails with ValidationException .
//
// - The following qualified state machine ARN refers to an alias named PROD .
//
// arn::states:::stateMachine:
//
// # If you provide a qualified state machine ARN that refers to a version ARN or an
//
// alias ARN, the request starts execution for that version or alias.
//
// - The following unqualified state machine ARN refers to a state machine named
// myStateMachine .
//
// arn::states:::stateMachine:
//
// If you start an execution with an unqualified state machine ARN, Step Functions
// uses the latest revision of the state machine for the execution.
//
// To start executions of a state machine [version], call StartExecution and provide the
// version ARN or the ARN of an [alias]that points to the version.
//
// StartExecution is idempotent for STANDARD workflows. For a STANDARD workflow,
// if you call StartExecution with the same name and input as a running execution,
// the call succeeds and return the same response as the original request. If the
// execution is closed or if the input is different, it returns a 400
// ExecutionAlreadyExists error. You can reuse names after 90 days.
//
// StartExecution isn't idempotent for EXPRESS workflows.
//
// [alias]: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-state-machine-alias.html
// [version]: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-state-machine-version.html
func sfn_StartExecution(cfg aws.Config, client *sfn.Client) {
	input := &sfn.StartExecutionInput{
		// StateMachineArn: *string, // Required
	}

	if len(_sfnStateMachineArn) > 0 {
		input.StateMachineArn = aws.String(_sfnStateMachineArn)
	}
	if len(_sfnInput) > 0 {
		input.Input = aws.String(_sfnInput)
	}
	if len(_sfnName) > 0 {
		input.Name = aws.String(_sfnName)
	}
	if len(_sfnTraceHeader) > 0 {
		input.TraceHeader = aws.String(_sfnTraceHeader)
	}

	if resp, err := client.StartExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a Synchronous Express state machine execution. StartSyncExecution is not
// available for STANDARD workflows.
//
// StartSyncExecution will return a 200 OK response, even if your execution fails,
// because the status code in the API response doesn't reflect function errors.
// Error codes are reserved for errors that prevent your execution from running,
// such as permissions errors, limit errors, or issues with your state machine code
// and configuration.
//
// This API action isn't logged in CloudTrail.
func sfn_StartSyncExecution(cfg aws.Config, client *sfn.Client) {
	input := &sfn.StartSyncExecutionInput{
		// StateMachineArn: *string, // Required
	}

	if len(_sfnStateMachineArn) > 0 {
		input.StateMachineArn = aws.String(_sfnStateMachineArn)
	}
	if len(_sfnIncludedData) > 0 {
		if err := assignInputField(input, "IncludedData", _sfnIncludedData); err != nil {
			log.Errorf("invalid --included-data: %s", err.Error())
			return
		}
	}
	if len(_sfnInput) > 0 {
		input.Input = aws.String(_sfnInput)
	}
	if len(_sfnName) > 0 {
		input.Name = aws.String(_sfnName)
	}
	if len(_sfnTraceHeader) > 0 {
		input.TraceHeader = aws.String(_sfnTraceHeader)
	}

	if resp, err := client.StartSyncExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops an execution.
// This API action is not supported by EXPRESS state machines.
//
// For an execution with encryption enabled, Step Functions will encrypt the error
// and cause fields using the KMS key for the execution role.
//
// A caller can stop an execution without using any KMS permissions in the
// execution role if the caller provides a null value for both error and cause
// fields because no data needs to be encrypted.
func sfn_StopExecution(cfg aws.Config, client *sfn.Client) {
	input := &sfn.StopExecutionInput{
		// ExecutionArn: *string, // Required
	}

	if len(_sfnExecutionArn) > 0 {
		input.ExecutionArn = aws.String(_sfnExecutionArn)
	}
	if len(_sfnCause) > 0 {
		input.Cause = aws.String(_sfnCause)
	}
	if len(_sfnError) > 0 {
		input.Error = aws.String(_sfnError)
	}

	if resp, err := client.StopExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Add a tag to a Step Functions resource.
// An array of key-value pairs. For more information, see [Using Cost Allocation Tags] in the Amazon Web
// Services Billing and Cost Management User Guide, and [Controlling Access Using IAM Tags].
//
// Tags may only contain Unicode letters, digits, white space, or these symbols: _
// . : / = + - (at) .
//
// [Controlling Access Using IAM Tags]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_iam-tags.html
// [Using Cost Allocation Tags]: https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html
func sfn_TagResource(cfg aws.Config, client *sfn.Client) {
	input := &sfn.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_sfnResourceArn) > 0 {
		input.ResourceArn = aws.String(_sfnResourceArn)
	}
	if len(_sfnTags) > 0 {
		if err := assignInputField(input, "Tags", _sfnTags); err != nil {
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

// Accepts the definition of a single state and executes it. You can test a state
// without creating a state machine or updating an existing state machine. Using
// this API, you can test the following:
//
// - A state's [input and output processing]data flow
//
// - An [Amazon Web Services service integration]request and response
//
// - An [HTTP Task]request and response
//
// You can call this API on only one state at a time. The states that you can test
// include the following:
//
// [All Task types]
// - except [Activity]
//
// [Pass]
//
// [Wait]
//
// [Choice]
//
// [Succeed]
//
// [Fail]
//
// The TestState API assumes an IAM role which must contain the required IAM
// permissions for the resources your state is accessing. For information about the
// permissions a state might need, see [IAM permissions to test a state].
//
// The TestState API can run for up to five minutes. If the execution of a state
// exceeds this duration, it fails with the States.Timeout error.
//
// TestState only supports the following when a mock is specified: [Activity tasks], .sync or
// .waitForTaskToken[service integration patterns] , [Parallel], or [Map] states.
//
// [Amazon Web Services service integration]: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-services.html
// [All Task types]: https://docs.aws.amazon.com/step-functions/latest/dg/amazon-states-language-task-state.html#task-types
// [Choice]: https://docs.aws.amazon.com/step-functions/latest/dg/amazon-states-language-choice-state.html
// [Activity tasks]: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-activities.html
// [HTTP Task]: https://docs.aws.amazon.com/step-functions/latest/dg/call-https-apis.html
// [input and output processing]: https://docs.aws.amazon.com/step-functions/latest/dg/test-state-isolation.html#test-state-input-output-dataflow
// [Activity]: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-activities.html
// [Parallel]: https://docs.aws.amazon.com/step-functions/latest/dg/amazon-states-language-parallel-state.html
// [Succeed]: https://docs.aws.amazon.com/step-functions/latest/dg/amazon-states-language-succeed-state.html
// [service integration patterns]: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html
// [Pass]: https://docs.aws.amazon.com/step-functions/latest/dg/amazon-states-language-pass-state.html
// [IAM permissions to test a state]: https://docs.aws.amazon.com/step-functions/latest/dg/test-state-isolation.html#test-state-permissions
// [Wait]: https://docs.aws.amazon.com/step-functions/latest/dg/amazon-states-language-wait-state.html
// [Map]: https://docs.aws.amazon.com/step-functions/latest/dg/amazon-states-language-map-state.html
// [Fail]: https://docs.aws.amazon.com/step-functions/latest/dg/amazon-states-language-fail-state.html
func sfn_TestState(cfg aws.Config, client *sfn.Client) {
	input := &sfn.TestStateInput{
		// Definition: *string, // Required
	}

	if len(_sfnDefinition) > 0 {
		input.Definition = aws.String(_sfnDefinition)
	}
	if len(_sfnContext) > 0 {
		input.Context = aws.String(_sfnContext)
	}
	if len(_sfnInput) > 0 {
		input.Input = aws.String(_sfnInput)
	}
	if len(_sfnInspectionLevel) > 0 {
		if err := assignInputField(input, "InspectionLevel", _sfnInspectionLevel); err != nil {
			log.Errorf("invalid --inspection-level: %s", err.Error())
			return
		}
	}
	if len(_sfnMock) > 0 {
		if err := assignInputField(input, "Mock", _sfnMock); err != nil {
			log.Errorf("invalid --mock: %s", err.Error())
			return
		}
	}
	if len(_sfnRevealSecrets) > 0 {
		if err := assignInputField(input, "RevealSecrets", _sfnRevealSecrets); err != nil {
			log.Errorf("invalid --reveal-secrets: %s", err.Error())
			return
		}
	}
	if len(_sfnRoleArn) > 0 {
		input.RoleArn = aws.String(_sfnRoleArn)
	}
	if len(_sfnStateConfiguration) > 0 {
		if err := assignInputField(input, "StateConfiguration", _sfnStateConfiguration); err != nil {
			log.Errorf("invalid --state-configuration: %s", err.Error())
			return
		}
	}
	if len(_sfnStateName) > 0 {
		input.StateName = aws.String(_sfnStateName)
	}
	if len(_sfnVariables) > 0 {
		input.Variables = aws.String(_sfnVariables)
	}

	if resp, err := client.TestState(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Remove a tag from a Step Functions resource
func sfn_UntagResource(cfg aws.Config, client *sfn.Client) {
	input := &sfn.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_sfnResourceArn) > 0 {
		input.ResourceArn = aws.String(_sfnResourceArn)
	}
	if len(_sfnTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _sfnTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an in-progress Map Run's configuration to include changes to the
// settings that control maximum concurrency and Map Run failure.
func sfn_UpdateMapRun(cfg aws.Config, client *sfn.Client) {
	input := &sfn.UpdateMapRunInput{
		// MapRunArn: *string, // Required
	}

	if len(_sfnMapRunArn) > 0 {
		input.MapRunArn = aws.String(_sfnMapRunArn)
	}
	if len(_sfnMaxConcurrency) > 0 {
		if err := assignInputField(input, "MaxConcurrency", _sfnMaxConcurrency); err != nil {
			log.Errorf("invalid --max-concurrency: %s", err.Error())
			return
		}
	}
	if len(_sfnToleratedFailureCount) > 0 {
		if err := assignInputField(input, "ToleratedFailureCount", _sfnToleratedFailureCount); err != nil {
			log.Errorf("invalid --tolerated-failure-count: %s", err.Error())
			return
		}
	}
	if len(_sfnToleratedFailurePercentage) > 0 {
		if err := assignInputField(input, "ToleratedFailurePercentage", _sfnToleratedFailurePercentage); err != nil {
			log.Errorf("invalid --tolerated-failure-percentage: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateMapRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing state machine by modifying its definition , roleArn ,
// loggingConfiguration , or EncryptionConfiguration . Running executions will
// continue to use the previous definition and roleArn . You must include at least
// one of definition or roleArn or you will receive a MissingRequiredParameter
// error.
//
// A qualified state machine ARN refers to a Distributed Map state defined within
// a state machine. For example, the qualified state machine ARN
// arn:partition:states:region:account-id:stateMachine:stateMachineName/mapStateLabel
// refers to a Distributed Map state with a label mapStateLabel in the state
// machine named stateMachineName .
//
// A qualified state machine ARN can either refer to a Distributed Map state
// defined within a state machine, a version ARN, or an alias ARN.
//
// The following are some examples of qualified and unqualified state machine ARNs:
//
// - The following qualified state machine ARN refers to a Distributed Map state
// with a label mapStateLabel in a state machine named myStateMachine .
//
// arn:partition:states:region:account-id:stateMachine:myStateMachine/mapStateLabel
//
// # If you provide a qualified state machine ARN that refers to a Distributed Map
//
// state, the request fails with ValidationException .
//
// - The following qualified state machine ARN refers to an alias named PROD .
//
// arn::states:::stateMachine:
//
// # If you provide a qualified state machine ARN that refers to a version ARN or an
//
// alias ARN, the request starts execution for that version or alias.
//
// - The following unqualified state machine ARN refers to a state machine named
// myStateMachine .
//
// arn::states:::stateMachine:
//
// After you update your state machine, you can set the publish parameter to true
// in the same action to publish a new [version]. This way, you can opt-in to strict
// versioning of your state machine.
//
// Step Functions assigns monotonically increasing integers for state machine
// versions, starting at version number 1.
//
// All StartExecution calls within a few seconds use the updated definition and
// roleArn . Executions started immediately after you call UpdateStateMachine may
// use the previous state machine definition and roleArn .
//
// [version]: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-state-machine-version.html
func sfn_UpdateStateMachine(cfg aws.Config, client *sfn.Client) {
	input := &sfn.UpdateStateMachineInput{
		// StateMachineArn: *string, // Required
	}

	if len(_sfnStateMachineArn) > 0 {
		input.StateMachineArn = aws.String(_sfnStateMachineArn)
	}
	if len(_sfnDefinition) > 0 {
		input.Definition = aws.String(_sfnDefinition)
	}
	if len(_sfnEncryptionConfiguration) > 0 {
		if err := assignInputField(input, "EncryptionConfiguration", _sfnEncryptionConfiguration); err != nil {
			log.Errorf("invalid --encryption-configuration: %s", err.Error())
			return
		}
	}
	if len(_sfnLoggingConfiguration) > 0 {
		if err := assignInputField(input, "LoggingConfiguration", _sfnLoggingConfiguration); err != nil {
			log.Errorf("invalid --logging-configuration: %s", err.Error())
			return
		}
	}
	if len(_sfnPublish) > 0 {
		if err := assignInputField(input, "Publish", _sfnPublish); err != nil {
			log.Errorf("invalid --publish: %s", err.Error())
			return
		}
	}
	if len(_sfnRoleArn) > 0 {
		input.RoleArn = aws.String(_sfnRoleArn)
	}
	if len(_sfnTracingConfiguration) > 0 {
		if err := assignInputField(input, "TracingConfiguration", _sfnTracingConfiguration); err != nil {
			log.Errorf("invalid --tracing-configuration: %s", err.Error())
			return
		}
	}
	if len(_sfnVersionDescription) > 0 {
		input.VersionDescription = aws.String(_sfnVersionDescription)
	}

	if resp, err := client.UpdateStateMachine(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration of an existing state machine [alias] by modifying its
// description or routingConfiguration .
//
// You must specify at least one of the description or routingConfiguration
// parameters to update a state machine alias.
//
// UpdateStateMachineAlias is an idempotent API. Step Functions bases the
// idempotency check on the stateMachineAliasArn , description , and
// routingConfiguration parameters. Requests with the same parameters return an
// idempotent response.
//
// This operation is eventually consistent. All StartExecution requests made within a few
// seconds use the latest alias configuration. Executions started immediately after
// calling UpdateStateMachineAlias may use the previous routing configuration.
//
// Related operations:
//
// # CreateStateMachineAlias
//
// # DescribeStateMachineAlias
//
// # ListStateMachineAliases
//
// # DeleteStateMachineAlias
//
// [alias]: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-state-machine-alias.html
func sfn_UpdateStateMachineAlias(cfg aws.Config, client *sfn.Client) {
	input := &sfn.UpdateStateMachineAliasInput{
		// StateMachineAliasArn: *string, // Required
	}

	if len(_sfnStateMachineAliasArn) > 0 {
		input.StateMachineAliasArn = aws.String(_sfnStateMachineAliasArn)
	}
	if len(_sfnDescription) > 0 {
		input.Description = aws.String(_sfnDescription)
	}
	if len(_sfnRoutingConfiguration) > 0 {
		if err := assignInputField(input, "RoutingConfiguration", _sfnRoutingConfiguration); err != nil {
			log.Errorf("invalid --routing-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateStateMachineAlias(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Validates the syntax of a state machine definition specified in [Amazon States Language] (ASL), a
// JSON-based, structured language.
//
// You can validate that a state machine definition is correct without creating a
// state machine resource.
//
// Suggested uses for ValidateStateMachineDefinition :
//
// - Integrate automated checks into your code review or Continuous Integration
// (CI) process to check state machine definitions before starting deployments.
//
// - Run validation from a Git pre-commit hook to verify the definition before
// committing to your source repository.
//
// Validation will look for problems in your state machine definition and return a
// result and a list of diagnostic elements.
//
// The result value will be OK when your workflow definition can be successfully
// created or updated. Note the result can be OK even when diagnostic warnings are
// present in the response. The result value will be FAIL when the workflow
// definition contains errors that would prevent you from creating or updating your
// state machine.
//
// The list of [ValidateStateMachineDefinitionDiagnostic] data elements can contain zero or more WARNING and/or ERROR
// elements.
//
// The ValidateStateMachineDefinition API might add new diagnostics in the future,
// adjust diagnostic codes, or change the message wording. Your automated processes
// should only rely on the value of the result field value (OK, FAIL). Do not rely
// on the exact order, count, or wording of diagnostic messages.
//
// [ValidateStateMachineDefinitionDiagnostic]: https://docs.aws.amazon.com/step-functions/latest/apireference/API_ValidateStateMachineDefinitionDiagnostic.html
// [Amazon States Language]: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-amazon-states-language.html
func sfn_ValidateStateMachineDefinition(cfg aws.Config, client *sfn.Client) {
	input := &sfn.ValidateStateMachineDefinitionInput{
		// Definition: *string, // Required
	}

	if len(_sfnDefinition) > 0 {
		input.Definition = aws.String(_sfnDefinition)
	}
	if len(_sfnMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sfnMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sfnSeverity) > 0 {
		if err := assignInputField(input, "Severity", _sfnSeverity); err != nil {
			log.Errorf("invalid --severity: %s", err.Error())
			return
		}
	}
	if len(_sfnType) > 0 {
		if err := assignInputField(input, "Type", _sfnType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if resp, err := client.ValidateStateMachineDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_sfnCmd)
	_sfnCmd.Flags().SortFlags = false

	_sfnCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_sfnCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_sfnCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_sfnCmd.Flags().StringVarP(&_sfnActivityArn, "activity-arn", "", "", "Activity ARN")
	_sfnCmd.Flags().StringVarP(&_sfnCause, "cause", "", "", "Cause")
	_sfnCmd.Flags().StringVarP(&_sfnClientToken, "client-token", "", "", "Client Token")
	_sfnCmd.Flags().StringVarP(&_sfnContext, "context", "", "", "Context")
	_sfnCmd.Flags().StringVarP(&_sfnDefinition, "definition", "", "", "Definition")
	_sfnCmd.Flags().StringVarP(&_sfnDescription, "description", "", "", "Description")
	_sfnCmd.Flags().StringVarP(&_sfnEncryptionConfiguration, "encryption-configuration", "", "", "Encryption Configuration")
	_sfnCmd.Flags().StringVarP(&_sfnError, "error", "", "", "Error")
	_sfnCmd.Flags().StringVarP(&_sfnExecutionArn, "execution-arn", "", "", "Execution ARN")
	_sfnCmd.Flags().StringVarP(&_sfnIncludeExecutionData, "include-execution-data", "", "", "Include Execution Data")
	_sfnCmd.Flags().StringVarP(&_sfnIncludedData, "included-data", "", "", "Included Data")
	_sfnCmd.Flags().StringVarP(&_sfnInput, "input", "", "", "Input")
	_sfnCmd.Flags().StringVarP(&_sfnInspectionLevel, "inspection-level", "", "", "Inspection Level")
	_sfnCmd.Flags().StringVarP(&_sfnLoggingConfiguration, "logging-configuration", "", "", "Logging Configuration")
	_sfnCmd.Flags().StringVarP(&_sfnMapRunArn, "map-run-arn", "", "", "Map Run ARN")
	_sfnCmd.Flags().StringVarP(&_sfnMaxConcurrency, "max-concurrency", "", "", "Max Concurrency")
	_sfnCmd.Flags().StringVarP(&_sfnMaxResults, "max-results", "", "", "Max Results")
	_sfnCmd.Flags().StringVarP(&_sfnMock, "mock", "", "", "Mock")
	_sfnCmd.Flags().StringVarP(&_sfnName, "name", "", "", "Name")
	_sfnCmd.Flags().StringVarP(&_sfnNextToken, "next-token", "", "", "Next Token")
	_sfnCmd.Flags().StringVarP(&_sfnPublish, "publish", "", "", "Publish")
	_sfnCmd.Flags().StringVarP(&_sfnRedriveFilter, "redrive-filter", "", "", "Redrive Filter")
	_sfnCmd.Flags().StringVarP(&_sfnResourceArn, "resource-arn", "", "", "Resource ARN")
	_sfnCmd.Flags().StringVarP(&_sfnRevealSecrets, "reveal-secrets", "", "", "Reveal Secrets")
	_sfnCmd.Flags().StringVarP(&_sfnReverseOrder, "reverse-order", "", "", "Reverse Order")
	_sfnCmd.Flags().StringVarP(&_sfnRevisionId, "revision-id", "", "", "Revision ID")
	_sfnCmd.Flags().StringVarP(&_sfnRoleArn, "role-arn", "", "", "Role ARN")
	_sfnCmd.Flags().StringVarP(&_sfnRoutingConfiguration, "routing-configuration", "", "", "Routing Configuration")
	_sfnCmd.Flags().StringVarP(&_sfnSeverity, "severity", "", "", "Severity")
	_sfnCmd.Flags().StringVarP(&_sfnStateConfiguration, "state-configuration", "", "", "State Configuration")
	_sfnCmd.Flags().StringVarP(&_sfnStateMachineAliasArn, "state-machine-alias-arn", "", "", "State Machine Alias ARN")
	_sfnCmd.Flags().StringVarP(&_sfnStateMachineArn, "state-machine-arn", "", "", "State Machine ARN")
	_sfnCmd.Flags().StringVarP(&_sfnStateMachineVersionArn, "state-machine-version-arn", "", "", "State Machine Version ARN")
	_sfnCmd.Flags().StringVarP(&_sfnStateName, "state-name", "", "", "State Name")
	_sfnCmd.Flags().StringVarP(&_sfnStatusFilter, "status-filter", "", "", "Status Filter")
	_sfnCmd.Flags().StringSliceVarP(&_sfnTagKeys, "tag-keys", "", nil, "Tag Keys")
	_sfnCmd.Flags().StringVarP(&_sfnTags, "tags", "", "", "Tags")
	_sfnCmd.Flags().StringVarP(&_sfnTaskToken, "task-token", "", "", "Task Token")
	_sfnCmd.Flags().StringVarP(&_sfnToleratedFailureCount, "tolerated-failure-count", "", "", "Tolerated Failure Count")
	_sfnCmd.Flags().StringVarP(&_sfnToleratedFailurePercentage, "tolerated-failure-percentage", "", "", "Tolerated Failure Percentage")
	_sfnCmd.Flags().StringVarP(&_sfnTraceHeader, "trace-header", "", "", "Trace Header")
	_sfnCmd.Flags().StringVarP(&_sfnTracingConfiguration, "tracing-configuration", "", "", "Tracing Configuration")
	_sfnCmd.Flags().StringVarP(&_sfnType, "type", "", "", "Type")
	_sfnCmd.Flags().StringVarP(&_sfnVariables, "variables", "", "", "Variables")
	_sfnCmd.Flags().StringVarP(&_sfnVersionDescription, "version-description", "", "", "Version Description")
	_sfnCmd.Flags().StringVarP(&_sfnWorkerName, "worker-name", "", "", "Worker Name")

	_sfnCmd.Flags().BoolVarP(&_sfnCreateActivity, "create-activity", "", false, "Create Activity")
	_sfnCmd.Flags().BoolVarP(&_sfnCreateStateMachine, "create-state-machine", "", false, "Create State Machine")
	_sfnCmd.Flags().BoolVarP(&_sfnCreateStateMachineAlias, "create-state-machine-alias", "", false, "Create State Machine Alias")
	_sfnCmd.Flags().BoolVarP(&_sfnDeleteActivity, "delete-activity", "", false, "Delete Activity")
	_sfnCmd.Flags().BoolVarP(&_sfnDeleteStateMachine, "delete-state-machine", "", false, "Delete State Machine")
	_sfnCmd.Flags().BoolVarP(&_sfnDeleteStateMachineAlias, "delete-state-machine-alias", "", false, "Delete State Machine Alias")
	_sfnCmd.Flags().BoolVarP(&_sfnDeleteStateMachineVersion, "delete-state-machine-version", "", false, "Delete State Machine Version")
	_sfnCmd.Flags().BoolVarP(&_sfnDescribeActivity, "describe-activity", "", false, "Describe Activity")
	_sfnCmd.Flags().BoolVarP(&_sfnDescribeExecution, "describe-execution", "", false, "Describe Execution")
	_sfnCmd.Flags().BoolVarP(&_sfnDescribeMapRun, "describe-map-run", "", false, "Describe Map Run")
	_sfnCmd.Flags().BoolVarP(&_sfnDescribeStateMachine, "describe-state-machine", "", false, "Describe State Machine")
	_sfnCmd.Flags().BoolVarP(&_sfnDescribeStateMachineAlias, "describe-state-machine-alias", "", false, "Describe State Machine Alias")
	_sfnCmd.Flags().BoolVarP(&_sfnDescribeStateMachineForExecution, "describe-state-machine-for-execution", "", false, "Describe State Machine For Execution")
	_sfnCmd.Flags().BoolVarP(&_sfnGetActivityTask, "get-activity-task", "", false, "Get Activity Task")
	_sfnCmd.Flags().BoolVarP(&_sfnGetExecutionHistory, "get-execution-history", "", false, "Get Execution History")
	_sfnCmd.Flags().BoolVarP(&_sfnListActivities, "list-activities", "", false, "List Activities")
	_sfnCmd.Flags().BoolVarP(&_sfnListExecutions, "list-executions", "", false, "List Executions")
	_sfnCmd.Flags().BoolVarP(&_sfnListMapRuns, "list-map-runs", "", false, "List Map Runs")
	_sfnCmd.Flags().BoolVarP(&_sfnListStateMachineAliases, "list-state-machine-aliases", "", false, "List State Machine Aliases")
	_sfnCmd.Flags().BoolVarP(&_sfnListStateMachineVersions, "list-state-machine-versions", "", false, "List State Machine Versions")
	_sfnCmd.Flags().BoolVarP(&_sfnListStateMachines, "list-state-machines", "", false, "List State Machines")
	_sfnCmd.Flags().BoolVarP(&_sfnListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_sfnCmd.Flags().BoolVarP(&_sfnPublishStateMachineVersion, "publish-state-machine-version", "", false, "Publish State Machine Version")
	_sfnCmd.Flags().BoolVarP(&_sfnRedriveExecution, "redrive-execution", "", false, "Redrive Execution")
	_sfnCmd.Flags().BoolVarP(&_sfnSendTaskFailure, "send-task-failure", "", false, "Send Task Failure")
	_sfnCmd.Flags().BoolVarP(&_sfnSendTaskHeartbeat, "send-task-heartbeat", "", false, "Send Task Heartbeat")
	_sfnCmd.Flags().BoolVarP(&_sfnSendTaskSuccess, "send-task-success", "", false, "Send Task Success")
	_sfnCmd.Flags().BoolVarP(&_sfnStartExecution, "start-execution", "", false, "Start Execution")
	_sfnCmd.Flags().BoolVarP(&_sfnStartSyncExecution, "start-sync-execution", "", false, "Start Sync Execution")
	_sfnCmd.Flags().BoolVarP(&_sfnStopExecution, "stop-execution", "", false, "Stop Execution")
	_sfnCmd.Flags().BoolVarP(&_sfnTagResource, "tag-resource", "", false, "Tag Resource")
	_sfnCmd.Flags().BoolVarP(&_sfnTestState, "test-state", "", false, "Test State")
	_sfnCmd.Flags().BoolVarP(&_sfnUntagResource, "untag-resource", "", false, "Untag Resource")
	_sfnCmd.Flags().BoolVarP(&_sfnUpdateMapRun, "update-map-run", "", false, "Update Map Run")
	_sfnCmd.Flags().BoolVarP(&_sfnUpdateStateMachine, "update-state-machine", "", false, "Update State Machine")
	_sfnCmd.Flags().BoolVarP(&_sfnUpdateStateMachineAlias, "update-state-machine-alias", "", false, "Update State Machine Alias")
	_sfnCmd.Flags().BoolVarP(&_sfnValidateStateMachineDefinition, "validate-state-machine-definition", "", false, "Validate State Machine Definition")

}
