package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/datapipeline"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// datapipelineCmd represents the datapipeline command
var _datapipelineCmd = &cobra.Command{
	Use:   "datapipeline",
	Short: "AWS datapipeline CLI",
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
		client := datapipeline.NewFromConfig(cfg)
		if _datapipelineActivatePipeline {
			datapipeline_ActivatePipeline(cfg, client)
			return
		}
		if _datapipelineAddTags {
			datapipeline_AddTags(cfg, client)
			return
		}
		if _datapipelineCreatePipeline {
			datapipeline_CreatePipeline(cfg, client)
			return
		}
		if _datapipelineDeactivatePipeline {
			datapipeline_DeactivatePipeline(cfg, client)
			return
		}
		if _datapipelineDeletePipeline {
			datapipeline_DeletePipeline(cfg, client)
			return
		}
		if _datapipelineDescribeObjects {
			datapipeline_DescribeObjects(cfg, client)
			return
		}
		if _datapipelineDescribePipelines {
			datapipeline_DescribePipelines(cfg, client)
			return
		}
		if _datapipelineEvaluateExpression {
			datapipeline_EvaluateExpression(cfg, client)
			return
		}
		if _datapipelineGetPipelineDefinition {
			datapipeline_GetPipelineDefinition(cfg, client)
			return
		}
		if _datapipelineListPipelines {
			datapipeline_ListPipelines(cfg, client)
			return
		}
		if _datapipelinePollForTask {
			datapipeline_PollForTask(cfg, client)
			return
		}
		if _datapipelinePutPipelineDefinition {
			datapipeline_PutPipelineDefinition(cfg, client)
			return
		}
		if _datapipelineQueryObjects {
			datapipeline_QueryObjects(cfg, client)
			return
		}
		if _datapipelineRemoveTags {
			datapipeline_RemoveTags(cfg, client)
			return
		}
		if _datapipelineReportTaskProgress {
			datapipeline_ReportTaskProgress(cfg, client)
			return
		}
		if _datapipelineReportTaskRunnerHeartbeat {
			datapipeline_ReportTaskRunnerHeartbeat(cfg, client)
			return
		}
		if _datapipelineSetStatus {
			datapipeline_SetStatus(cfg, client)
			return
		}
		if _datapipelineSetTaskStatus {
			datapipeline_SetTaskStatus(cfg, client)
			return
		}
		if _datapipelineValidatePipelineDefinition {
			datapipeline_ValidatePipelineDefinition(cfg, client)
			return
		}

	},
}

var (
	_datapipelineActivatePipeline           bool
	_datapipelineAddTags                    bool
	_datapipelineCreatePipeline             bool
	_datapipelineDeactivatePipeline         bool
	_datapipelineDeletePipeline             bool
	_datapipelineDescribeObjects            bool
	_datapipelineDescribePipelines          bool
	_datapipelineEvaluateExpression         bool
	_datapipelineGetPipelineDefinition      bool
	_datapipelineListPipelines              bool
	_datapipelinePollForTask                bool
	_datapipelinePutPipelineDefinition      bool
	_datapipelineQueryObjects               bool
	_datapipelineRemoveTags                 bool
	_datapipelineReportTaskProgress         bool
	_datapipelineReportTaskRunnerHeartbeat  bool
	_datapipelineSetStatus                  bool
	_datapipelineSetTaskStatus              bool
	_datapipelineValidatePipelineDefinition bool

	_datapipelineCancelActive        string
	_datapipelineDescription         string
	_datapipelineErrorId             string
	_datapipelineErrorMessage        string
	_datapipelineErrorStackTrace     string
	_datapipelineEvaluateExpressions string
	_datapipelineExpression          string
	_datapipelineFields              string
	_datapipelineHostname            string
	_datapipelineInstanceIdentity    string
	_datapipelineLimit               string
	_datapipelineMarker              string
	_datapipelineName                string
	_datapipelineObjectId            string
	_datapipelineObjectIds           []string
	_datapipelineParameterObjects    string
	_datapipelineParameterValues     string
	_datapipelinePipelineId          string
	_datapipelinePipelineIds         []string
	_datapipelinePipelineObjects     string
	_datapipelineQuery               string
	_datapipelineSphere              string
	_datapipelineStartTimestamp      string
	_datapipelineStatus              string
	_datapipelineTagKeys             []string
	_datapipelineTags                string
	_datapipelineTaskId              string
	_datapipelineTaskStatus          string
	_datapipelineTaskrunnerId        string
	_datapipelineUniqueId            string
	_datapipelineVersion             string
	_datapipelineWorkerGroup         string
)

// Validates the specified pipeline and starts processing pipeline tasks. If the
// pipeline does not pass validation, activation fails.
//
// If you need to pause the pipeline to investigate an issue with a component,
// such as a data source or script, call DeactivatePipeline.
//
// To activate a finished pipeline, modify the end date for the pipeline and then
// activate it.
//
// POST / HTTP/1.1 Content-Type: application/x-amz-json-1.1 X-Amz-Target:
// DataPipeline.ActivatePipeline Content-Length: 39 Host:
// datapipeline.us-east-1.amazonaws.com X-Amz-Date: Mon, 12 Nov 2012 17:49:52 GMT
// Authorization: AuthParams
//
// {"pipelineId": "df-06372391ZG65EXAMPLE"}
//
// HTTP/1.1 200 x-amzn-RequestId: ee19d5bf-074e-11e2-af6f-6bc7a6be60d9
// Content-Type: application/x-amz-json-1.1 Content-Length: 2 Date: Mon, 12 Nov
// 2012 17:50:53 GMT
//
// {}
func datapipeline_ActivatePipeline(cfg aws.Config, client *datapipeline.Client) {
	input := &datapipeline.ActivatePipelineInput{
		// PipelineId: *string, // Required
	}

	if len(_datapipelinePipelineId) > 0 {
		input.PipelineId = aws.String(_datapipelinePipelineId)
	}
	if len(_datapipelineParameterValues) > 0 {
		if err := assignInputField(input, "ParameterValues", _datapipelineParameterValues); err != nil {
			log.Errorf("invalid --parameter-values: %s", err.Error())
			return
		}
	}
	if len(_datapipelineStartTimestamp) > 0 {
		if err := assignInputField(input, "StartTimestamp", _datapipelineStartTimestamp); err != nil {
			log.Errorf("invalid --start-timestamp: %s", err.Error())
			return
		}
	}

	if resp, err := client.ActivatePipeline(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds or modifies tags for the specified pipeline.
func datapipeline_AddTags(cfg aws.Config, client *datapipeline.Client) {
	input := &datapipeline.AddTagsInput{
		// PipelineId: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_datapipelinePipelineId) > 0 {
		input.PipelineId = aws.String(_datapipelinePipelineId)
	}
	if len(_datapipelineTags) > 0 {
		if err := assignInputField(input, "Tags", _datapipelineTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.AddTags(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new, empty pipeline. Use PutPipelineDefinition to populate the pipeline.
// POST / HTTP/1.1 Content-Type: application/x-amz-json-1.1 X-Amz-Target:
// DataPipeline.CreatePipeline Content-Length: 91 Host:
// datapipeline.us-east-1.amazonaws.com X-Amz-Date: Mon, 12 Nov 2012 17:49:52 GMT
// Authorization: AuthParams
//
// {"name": "myPipeline", "uniqueId": "123456789", "description": "This is my
// first pipeline"}
//
// HTTP/1.1 200 x-amzn-RequestId: b16911ce-0774-11e2-af6f-6bc7a6be60d9
// Content-Type: application/x-amz-json-1.1 Content-Length: 40 Date: Mon, 12 Nov
// 2012 17:50:53 GMT
//
// {"pipelineId": "df-06372391ZG65EXAMPLE"}
func datapipeline_CreatePipeline(cfg aws.Config, client *datapipeline.Client) {
	input := &datapipeline.CreatePipelineInput{
		// Name: *string, // Required
		// UniqueId: *string, // Required
	}

	if len(_datapipelineName) > 0 {
		input.Name = aws.String(_datapipelineName)
	}
	if len(_datapipelineUniqueId) > 0 {
		input.UniqueId = aws.String(_datapipelineUniqueId)
	}
	if len(_datapipelineDescription) > 0 {
		input.Description = aws.String(_datapipelineDescription)
	}
	if len(_datapipelineTags) > 0 {
		if err := assignInputField(input, "Tags", _datapipelineTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreatePipeline(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deactivates the specified running pipeline. The pipeline is set to the
// DEACTIVATING state until the deactivation process completes.
//
// To resume a deactivated pipeline, use ActivatePipeline. By default, the pipeline resumes from
// the last completed execution. Optionally, you can specify the date and time to
// resume the pipeline.
func datapipeline_DeactivatePipeline(cfg aws.Config, client *datapipeline.Client) {
	input := &datapipeline.DeactivatePipelineInput{
		// PipelineId: *string, // Required
	}

	if len(_datapipelinePipelineId) > 0 {
		input.PipelineId = aws.String(_datapipelinePipelineId)
	}
	if len(_datapipelineCancelActive) > 0 {
		if err := assignInputField(input, "CancelActive", _datapipelineCancelActive); err != nil {
			log.Errorf("invalid --cancel-active: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeactivatePipeline(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a pipeline, its pipeline definition, and its run history. AWS Data
// Pipeline attempts to cancel instances associated with the pipeline that are
// currently being processed by task runners.
//
// Deleting a pipeline cannot be undone. You cannot query or restore a deleted
// pipeline. To temporarily pause a pipeline instead of deleting it, call SetStatuswith the
// status set to PAUSE on individual components. Components that are paused by SetStatus
// can be resumed.
//
// POST / HTTP/1.1 Content-Type: application/x-amz-json-1.1 X-Amz-Target:
// DataPipeline.DeletePipeline Content-Length: 50 Host:
// datapipeline.us-east-1.amazonaws.com X-Amz-Date: Mon, 12 Nov 2012 17:49:52 GMT
// Authorization: AuthParams
//
// {"pipelineId": "df-06372391ZG65EXAMPLE"}
//
// x-amzn-RequestId: b7a88c81-0754-11e2-af6f-6bc7a6be60d9 Content-Type:
// application/x-amz-json-1.1 Content-Length: 0 Date: Mon, 12 Nov 2012 17:50:53 GMT
//
// Unexpected response: 200, OK, undefined
func datapipeline_DeletePipeline(cfg aws.Config, client *datapipeline.Client) {
	input := &datapipeline.DeletePipelineInput{
		// PipelineId: *string, // Required
	}

	if len(_datapipelinePipelineId) > 0 {
		input.PipelineId = aws.String(_datapipelinePipelineId)
	}

	if resp, err := client.DeletePipeline(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the object definitions for a set of objects associated with the pipeline.
// Object definitions are composed of a set of fields that define the properties of
// the object.
//
// POST / HTTP/1.1 Content-Type: application/x-amz-json-1.1 X-Amz-Target:
// DataPipeline.DescribeObjects Content-Length: 98 Host:
// datapipeline.us-east-1.amazonaws.com X-Amz-Date: Mon, 12 Nov 2012 17:49:52 GMT
// Authorization: AuthParams
//
// {"pipelineId": "df-06372391ZG65EXAMPLE", "objectIds": ["Schedule"],
// "evaluateExpressions": true}
//
// x-amzn-RequestId: 4c18ea5d-0777-11e2-8a14-21bb8a1f50ef Content-Type:
// application/x-amz-json-1.1 Content-Length: 1488 Date: Mon, 12 Nov 2012 17:50:53
// GMT
//
// {"hasMoreResults": false, "pipelineObjects": [ {"fields": [ {"key":
// "startDateTime", "stringValue": "2012-12-12T00:00:00"}, {"key": "parent",
// "refValue": "Default"}, {"key": "(at)sphere", "stringValue": "COMPONENT"}, {"key":
// "type", "stringValue": "Schedule"}, {"key": "period", "stringValue": "1 hour"},
// {"key": "endDateTime", "stringValue": "2012-12-21T18:00:00"}, {"key":
// "(at)version", "stringValue": "1"}, {"key": "(at)status", "stringValue": "PENDING"},
// {"key": "(at)pipelineId", "stringValue": "df-06372391ZG65EXAMPLE"} ], "id":
// "Schedule", "name": "Schedule"} ] }
func datapipeline_DescribeObjects(cfg aws.Config, client *datapipeline.Client) {
	input := &datapipeline.DescribeObjectsInput{
		// ObjectIds: []string, // Required
		// PipelineId: *string, // Required
	}

	if len(_datapipelineObjectIds) > 0 {
		input.ObjectIds = append([]string(nil), _datapipelineObjectIds...)
	}
	if len(_datapipelinePipelineId) > 0 {
		input.PipelineId = aws.String(_datapipelinePipelineId)
	}
	if len(_datapipelineEvaluateExpressions) > 0 {
		if err := assignInputField(input, "EvaluateExpressions", _datapipelineEvaluateExpressions); err != nil {
			log.Errorf("invalid --evaluate-expressions: %s", err.Error())
			return
		}
	}
	if len(_datapipelineMarker) > 0 {
		input.Marker = aws.String(_datapipelineMarker)
	}

	if disablePaginator() {
		if resp, err := client.DescribeObjects(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*datapipeline.DescribeObjectsOutput
	p := datapipeline.NewDescribeObjectsPaginator(client, input)
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

// Retrieves metadata about one or more pipelines. The information retrieved
// includes the name of the pipeline, the pipeline identifier, its current state,
// and the user account that owns the pipeline. Using account credentials, you can
// retrieve metadata about pipelines that you or your IAM users have created. If
// you are using an IAM user account, you can retrieve metadata about only those
// pipelines for which you have read permissions.
//
// To retrieve the full pipeline definition instead of metadata about the
// pipeline, call GetPipelineDefinition.
//
// POST / HTTP/1.1 Content-Type: application/x-amz-json-1.1 X-Amz-Target:
// DataPipeline.DescribePipelines Content-Length: 70 Host:
// datapipeline.us-east-1.amazonaws.com X-Amz-Date: Mon, 12 Nov 2012 17:49:52 GMT
// Authorization: AuthParams
//
// {"pipelineIds": ["df-08785951KAKJEXAMPLE"] }
//
// x-amzn-RequestId: 02870eb7-0736-11e2-af6f-6bc7a6be60d9 Content-Type:
// application/x-amz-json-1.1 Content-Length: 767 Date: Mon, 12 Nov 2012 17:50:53
// GMT
//
// {"pipelineDescriptionList": [ {"description": "This is my first pipeline",
// "fields": [ {"key": "(at)pipelineState", "stringValue": "SCHEDULED"}, {"key":
// "description", "stringValue": "This is my first pipeline"}, {"key": "name",
// "stringValue": "myPipeline"}, {"key": "(at)creationTime", "stringValue":
// "2012-12-13T01:24:06"}, {"key": "(at)id", "stringValue": "df-0937003356ZJEXAMPLE"},
// {"key": "(at)sphere", "stringValue": "PIPELINE"}, {"key": "(at)version",
// "stringValue": "1"}, {"key": "(at)userId", "stringValue": "924374875933"}, {"key":
// "(at)accountId", "stringValue": "924374875933"}, {"key": "uniqueId", "stringValue":
// "1234567890"} ], "name": "myPipeline", "pipelineId": "df-0937003356ZJEXAMPLE"} ]
// }
func datapipeline_DescribePipelines(cfg aws.Config, client *datapipeline.Client) {
	input := &datapipeline.DescribePipelinesInput{
		// PipelineIds: []string, // Required
	}

	if len(_datapipelinePipelineIds) > 0 {
		input.PipelineIds = append([]string(nil), _datapipelinePipelineIds...)
	}

	if resp, err := client.DescribePipelines(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Task runners call EvaluateExpression to evaluate a string in the context of the
// specified object. For example, a task runner can evaluate SQL queries stored in
// Amazon S3.
//
// POST / HTTP/1.1 Content-Type: application/x-amz-json-1.1 X-Amz-Target:
// DataPipeline.DescribePipelines Content-Length: 164 Host:
// datapipeline.us-east-1.amazonaws.com X-Amz-Date: Mon, 12 Nov 2012 17:49:52 GMT
// Authorization: AuthParams
//
// {"pipelineId": "df-08785951KAKJEXAMPLE", "objectId": "Schedule", "expression":
// "Transform started at #{startDateTime} and finished at #{endDateTime}"}
//
// x-amzn-RequestId: 02870eb7-0736-11e2-af6f-6bc7a6be60d9 Content-Type:
// application/x-amz-json-1.1 Content-Length: 103 Date: Mon, 12 Nov 2012 17:50:53
// GMT
//
// {"evaluatedExpression": "Transform started at 2012-12-12T00:00:00 and finished
// at 2012-12-21T18:00:00"}
func datapipeline_EvaluateExpression(cfg aws.Config, client *datapipeline.Client) {
	input := &datapipeline.EvaluateExpressionInput{
		// Expression: *string, // Required
		// ObjectId: *string, // Required
		// PipelineId: *string, // Required
	}

	if len(_datapipelineExpression) > 0 {
		input.Expression = aws.String(_datapipelineExpression)
	}
	if len(_datapipelineObjectId) > 0 {
		input.ObjectId = aws.String(_datapipelineObjectId)
	}
	if len(_datapipelinePipelineId) > 0 {
		input.PipelineId = aws.String(_datapipelinePipelineId)
	}

	if resp, err := client.EvaluateExpression(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the definition of the specified pipeline. You can call
// GetPipelineDefinition to retrieve the pipeline definition that you provided
// using PutPipelineDefinition.
//
// POST / HTTP/1.1 Content-Type: application/x-amz-json-1.1 X-Amz-Target:
// DataPipeline.GetPipelineDefinition Content-Length: 40 Host:
// datapipeline.us-east-1.amazonaws.com X-Amz-Date: Mon, 12 Nov 2012 17:49:52 GMT
// Authorization: AuthParams
//
// {"pipelineId": "df-06372391ZG65EXAMPLE"}
//
// x-amzn-RequestId: e28309e5-0776-11e2-8a14-21bb8a1f50ef Content-Type:
// application/x-amz-json-1.1 Content-Length: 890 Date: Mon, 12 Nov 2012 17:50:53
// GMT
//
// {"pipelineObjects": [ {"fields": [ {"key": "workerGroup", "stringValue":
// "workerGroup"} ], "id": "Default", "name": "Default"}, {"fields": [ {"key":
// "startDateTime", "stringValue": "2012-09-25T17:00:00"}, {"key": "type",
// "stringValue": "Schedule"}, {"key": "period", "stringValue": "1 hour"}, {"key":
// "endDateTime", "stringValue": "2012-09-25T18:00:00"} ], "id": "Schedule",
// "name": "Schedule"}, {"fields": [ {"key": "schedule", "refValue": "Schedule"},
// {"key": "command", "stringValue": "echo hello"}, {"key": "parent", "refValue":
// "Default"}, {"key": "type", "stringValue": "ShellCommandActivity"} ], "id":
// "SayHello", "name": "SayHello"} ] }
func datapipeline_GetPipelineDefinition(cfg aws.Config, client *datapipeline.Client) {
	input := &datapipeline.GetPipelineDefinitionInput{
		// PipelineId: *string, // Required
	}

	if len(_datapipelinePipelineId) > 0 {
		input.PipelineId = aws.String(_datapipelinePipelineId)
	}
	if len(_datapipelineVersion) > 0 {
		input.Version = aws.String(_datapipelineVersion)
	}

	if resp, err := client.GetPipelineDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the pipeline identifiers for all active pipelines that you have
// permission to access.
//
// POST / HTTP/1.1 Content-Type: application/x-amz-json-1.1 X-Amz-Target:
// DataPipeline.ListPipelines Content-Length: 14 Host:
// datapipeline.us-east-1.amazonaws.com X-Amz-Date: Mon, 12 Nov 2012 17:49:52 GMT
// Authorization: AuthParams
//
// {}
//
// Status: x-amzn-RequestId: b3104dc5-0734-11e2-af6f-6bc7a6be60d9 Content-Type:
// application/x-amz-json-1.1 Content-Length: 39 Date: Mon, 12 Nov 2012 17:50:53
// GMT
//
// {"PipelineIdList": [ {"id": "df-08785951KAKJEXAMPLE", "name": "MyPipeline"},
// {"id": "df-08662578ISYEXAMPLE", "name": "MySecondPipeline"} ] }
func datapipeline_ListPipelines(cfg aws.Config, client *datapipeline.Client) {
	input := &datapipeline.ListPipelinesInput{}

	if len(_datapipelineMarker) > 0 {
		input.Marker = aws.String(_datapipelineMarker)
	}

	if disablePaginator() {
		if resp, err := client.ListPipelines(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*datapipeline.ListPipelinesOutput
	p := datapipeline.NewListPipelinesPaginator(client, input)
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

// Task runners call PollForTask to receive a task to perform from AWS Data
// Pipeline. The task runner specifies which tasks it can perform by setting a
// value for the workerGroup parameter. The task returned can come from any of the
// pipelines that match the workerGroup value passed in by the task runner and
// that was launched using the IAM user credentials specified by the task runner.
//
// If tasks are ready in the work queue, PollForTask returns a response
// immediately. If no tasks are available in the queue, PollForTask uses
// long-polling and holds on to a poll connection for up to a 90 seconds, during
// which time the first newly scheduled task is handed to the task runner. To
// accomodate this, set the socket timeout in your task runner to 90 seconds. The
// task runner should not call PollForTask again on the same workerGroup until it
// receives a response, and this can take up to 90 seconds.
//
// POST / HTTP/1.1 Content-Type: application/x-amz-json-1.1 X-Amz-Target:
// DataPipeline.PollForTask Content-Length: 59 Host:
// datapipeline.us-east-1.amazonaws.com X-Amz-Date: Mon, 12 Nov 2012 17:49:52 GMT
// Authorization: AuthParams
//
// {"workerGroup": "MyworkerGroup", "hostname": "example.com"}
//
// x-amzn-RequestId: 41c713d2-0775-11e2-af6f-6bc7a6be60d9 Content-Type:
// application/x-amz-json-1.1 Content-Length: 39 Date: Mon, 12 Nov 2012 17:50:53
// GMT
//
// {"taskObject": {"attemptId": "(at)SayHello_2012-12-12T00:00:00_Attempt=1",
// "objects": {"(at)SayHello_2012-12-12T00:00:00_Attempt=1": {"fields": [ {"key":
// "(at)componentParent", "refValue": "SayHello"}, {"key": "(at)scheduledStartTime",
// "stringValue": "2012-12-12T00:00:00"}, {"key": "parent", "refValue":
// "SayHello"}, {"key": "(at)sphere", "stringValue": "ATTEMPT"}, {"key":
// "workerGroup", "stringValue": "workerGroup"}, {"key": "(at)instanceParent",
// "refValue": "(at)SayHello_2012-12-12T00:00:00"}, {"key": "type", "stringValue":
// "ShellCommandActivity"}, {"key": "(at)status", "stringValue":
// "WAITING_FOR_RUNNER"}, {"key": "(at)version", "stringValue": "1"}, {"key":
// "schedule", "refValue": "Schedule"}, {"key": "(at)actualStartTime", "stringValue":
// "2012-12-13T01:40:50"}, {"key": "command", "stringValue": "echo hello"}, {"key":
// "(at)scheduledEndTime", "stringValue": "2012-12-12T01:00:00"}, {"key":
// "(at)activeInstances", "refValue": "(at)SayHello_2012-12-12T00:00:00"}, {"key":
// "(at)pipelineId", "stringValue": "df-0937003356ZJEXAMPLE"} ], "id":
// "(at)SayHello_2012-12-12T00:00:00_Attempt=1", "name":
// "(at)SayHello_2012-12-12T00:00:00_Attempt=1"} }, "pipelineId":
// "df-0937003356ZJEXAMPLE", "taskId":
// "2xaM4wRs5zOsIH+g9U3oVHfAgAlbSqU6XduncB0HhZ3xMnmvfePZPn4dIbYXHyWyRK+cU15MqDHwdrvftx/4wv+sNS4w34vJfv7QA9aOoOazW28l1GYSb2ZRR0N0paiQp+d1MhSKo10hOTWOsVK5S5Lnx9Qm6omFgXHyIvZRIvTlrQMpr1xuUrflyGOfbFOGpOLpvPE172MYdqpZKnbSS4TcuqgQKSWV2833fEubI57DPOP7ghWa2TcYeSIv4pdLYG53fTuwfbnbdc98g2LNUQzSVhSnt7BoqyNwht2aQ6b/UHg9A80+KVpuXuqmz3m1MXwHFgxjdmuesXNOrrlGpeLCcRWD+aGo0RN1NqhQRzNAig8V4GlaPTQzMsRCljKqvrIyAoP3Tt2XEGsHkkQo12rEX8Z90957XX2qKRwhruwYzqGkSLWjINoLdAxUJdpRXRc5DJTrBd3D5mdzn7kY1l7NEh4kFHJDt3Cx4Z3Mk8MYCACyCk/CEyy9DwuPi66cLz0NBcgbCM5LKjTBOwo1m+am+pvM1kSposE9FPP1+RFGb8k6jQBTJx3TRz1yKilnGXQTZ5xvdOFpJrklIT0OXP1MG3+auM9FlJA+1dX90QoNJE5z7axmK//MOGXUdkqFe2kiDkorqjxwDvc0Js9pVKfKvAmW8YqUbmI9l0ERpWCXXnLVHNmPWz3jaPY+OBAmuJWDmxB/Z8p94aEDg4BVXQ7LvsKQ3DLYhaB7yJ390CJT+i0mm+EBqY60V6YikPSWDFrYQ/NPi2b1DgE19mX8zHqw8qprIl4yh1Ckx2Iige4En/N5ktOoIxnASxAw/TzcE2skxdw5KlHDF+UTj71m16CR/dIaKlXijlfNlNzUBo/bNSadCQn3G5NoO501wPKI:XO50TgDNyo8EXAMPLE/g==:1"}
// }
func datapipeline_PollForTask(cfg aws.Config, client *datapipeline.Client) {
	input := &datapipeline.PollForTaskInput{
		// WorkerGroup: *string, // Required
	}

	if len(_datapipelineWorkerGroup) > 0 {
		input.WorkerGroup = aws.String(_datapipelineWorkerGroup)
	}
	if len(_datapipelineHostname) > 0 {
		input.Hostname = aws.String(_datapipelineHostname)
	}
	if len(_datapipelineInstanceIdentity) > 0 {
		if err := assignInputField(input, "InstanceIdentity", _datapipelineInstanceIdentity); err != nil {
			log.Errorf("invalid --instance-identity: %s", err.Error())
			return
		}
	}

	if resp, err := client.PollForTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds tasks, schedules, and preconditions to the specified pipeline. You can use
// PutPipelineDefinition to populate a new pipeline.
//
// PutPipelineDefinition also validates the configuration as it adds it to the
// pipeline. Changes to the pipeline are saved unless one of the following three
// validation errors exists in the pipeline.
//
// - An object is missing a name or identifier field.
// - A string or reference field is empty.
// - The number of objects in the pipeline exceeds the maximum allowed objects.
// - The pipeline is in a FINISHED state.
//
// Pipeline object definitions are passed to the PutPipelineDefinition action and
// returned by the GetPipelineDefinitionaction.
//
// Example 1 This example sets an valid pipeline configuration and returns
// success. POST / HTTP/1.1 Content-Type: application/x-amz-json-1.1 X-Amz-Target:
// DataPipeline.PutPipelineDefinition Content-Length: 914 Host:
// datapipeline.us-east-1.amazonaws.com X-Amz-Date: Mon, 12 Nov 2012 17:49:52 GMT
// Authorization: AuthParams
//
// {"pipelineId": "df-0937003356ZJEXAMPLE", "pipelineObjects": [ {"id": "Default",
// "name": "Default", "fields": [ {"key": "workerGroup", "stringValue":
// "workerGroup"} ] }, {"id": "Schedule", "name": "Schedule", "fields": [ {"key":
// "startDateTime", "stringValue": "2012-12-12T00:00:00"}, {"key": "type",
// "stringValue": "Schedule"}, {"key": "period", "stringValue": "1 hour"}, {"key":
// "endDateTime", "stringValue": "2012-12-21T18:00:00"} ] }, {"id": "SayHello",
// "name": "SayHello", "fields": [ {"key": "type", "stringValue":
// "ShellCommandActivity"}, {"key": "command", "stringValue": "echo hello"},
// {"key": "parent", "refValue": "Default"}, {"key": "schedule", "refValue":
// "Schedule"} ] } ] }
//
// HTTP/1.1 200 x-amzn-RequestId: f74afc14-0754-11e2-af6f-6bc7a6be60d9
// Content-Type: application/x-amz-json-1.1 Content-Length: 18 Date: Mon, 12 Nov
// 2012 17:50:53 GMT
//
// {"errored": false}
//
// Example 2 This example sets an invalid pipeline configuration (the value for
// workerGroup is an empty string) and returns an error message.
//
// POST / HTTP/1.1 Content-Type: application/x-amz-json-1.1 X-Amz-Target:
// DataPipeline.PutPipelineDefinition Content-Length: 903 Host:
// datapipeline.us-east-1.amazonaws.com X-Amz-Date: Mon, 12 Nov 2012 17:49:52 GMT
// Authorization: AuthParams
//
// {"pipelineId": "df-06372391ZG65EXAMPLE", "pipelineObjects": [ {"id": "Default",
// "name": "Default", "fields": [ {"key": "workerGroup", "stringValue": ""} ] },
// {"id": "Schedule", "name": "Schedule", "fields": [ {"key": "startDateTime",
// "stringValue": "2012-09-25T17:00:00"}, {"key": "type", "stringValue":
// "Schedule"}, {"key": "period", "stringValue": "1 hour"}, {"key": "endDateTime",
// "stringValue": "2012-09-25T18:00:00"} ] }, {"id": "SayHello", "name":
// "SayHello", "fields": [ {"key": "type", "stringValue": "ShellCommandActivity"},
// {"key": "command", "stringValue": "echo hello"}, {"key": "parent", "refValue":
// "Default"}, {"key": "schedule", "refValue": "Schedule"}
//
// ] } ] }
//
// HTTP/1.1 200 x-amzn-RequestId: f74afc14-0754-11e2-af6f-6bc7a6be60d9
// Content-Type: application/x-amz-json-1.1 Content-Length: 18 Date: Mon, 12 Nov
// 2012 17:50:53 GMT
//
// {"__type": "com.amazon.setl.webservice#InvalidRequestException", "message":
// "Pipeline definition has errors: Could not save the pipeline definition due to
// FATAL errors: [com.amazon.setl.webservice.ValidationError(at)108d7ea9] Please call
// Validate to validate your pipeline"}
func datapipeline_PutPipelineDefinition(cfg aws.Config, client *datapipeline.Client) {
	input := &datapipeline.PutPipelineDefinitionInput{
		// PipelineId: *string, // Required
		// PipelineObjects: []types.PipelineObject, // Required
	}

	if len(_datapipelinePipelineId) > 0 {
		input.PipelineId = aws.String(_datapipelinePipelineId)
	}
	if len(_datapipelinePipelineObjects) > 0 {
		if err := assignInputField(input, "PipelineObjects", _datapipelinePipelineObjects); err != nil {
			log.Errorf("invalid --pipeline-objects: %s", err.Error())
			return
		}
	}
	if len(_datapipelineParameterObjects) > 0 {
		if err := assignInputField(input, "ParameterObjects", _datapipelineParameterObjects); err != nil {
			log.Errorf("invalid --parameter-objects: %s", err.Error())
			return
		}
	}
	if len(_datapipelineParameterValues) > 0 {
		if err := assignInputField(input, "ParameterValues", _datapipelineParameterValues); err != nil {
			log.Errorf("invalid --parameter-values: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutPipelineDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Queries the specified pipeline for the names of objects that match the
// specified set of conditions.
//
// POST / HTTP/1.1 Content-Type: application/x-amz-json-1.1 X-Amz-Target:
// DataPipeline.QueryObjects Content-Length: 123 Host:
// datapipeline.us-east-1.amazonaws.com X-Amz-Date: Mon, 12 Nov 2012 17:49:52 GMT
// Authorization: AuthParams
//
// {"pipelineId": "df-06372391ZG65EXAMPLE", "query": {"selectors": [ ] },
// "sphere": "INSTANCE", "marker": "", "limit": 10}
//
// x-amzn-RequestId: 14d704c1-0775-11e2-af6f-6bc7a6be60d9 Content-Type:
// application/x-amz-json-1.1 Content-Length: 72 Date: Mon, 12 Nov 2012 17:50:53
// GMT
//
// {"hasMoreResults": false, "ids": ["(at)SayHello_1_2012-09-25T17:00:00"] }
func datapipeline_QueryObjects(cfg aws.Config, client *datapipeline.Client) {
	input := &datapipeline.QueryObjectsInput{
		// PipelineId: *string, // Required
		// Sphere: *string, // Required
	}

	if len(_datapipelinePipelineId) > 0 {
		input.PipelineId = aws.String(_datapipelinePipelineId)
	}
	if len(_datapipelineSphere) > 0 {
		input.Sphere = aws.String(_datapipelineSphere)
	}
	if len(_datapipelineLimit) > 0 {
		if err := assignInputField(input, "Limit", _datapipelineLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_datapipelineMarker) > 0 {
		input.Marker = aws.String(_datapipelineMarker)
	}
	if len(_datapipelineQuery) > 0 {
		if err := assignInputField(input, "Query", _datapipelineQuery); err != nil {
			log.Errorf("invalid --query: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.QueryObjects(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*datapipeline.QueryObjectsOutput
	p := datapipeline.NewQueryObjectsPaginator(client, input)
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

// Removes existing tags from the specified pipeline.
func datapipeline_RemoveTags(cfg aws.Config, client *datapipeline.Client) {
	input := &datapipeline.RemoveTagsInput{
		// PipelineId: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_datapipelinePipelineId) > 0 {
		input.PipelineId = aws.String(_datapipelinePipelineId)
	}
	if len(_datapipelineTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _datapipelineTagKeys...)
	}

	if resp, err := client.RemoveTags(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Task runners call ReportTaskProgress when assigned a task to acknowledge that
// it has the task. If the web service does not receive this acknowledgement within
// 2 minutes, it assigns the task in a subsequent PollForTaskcall. After this initial
// acknowledgement, the task runner only needs to report progress every 15 minutes
// to maintain its ownership of the task. You can change this reporting time from
// 15 minutes by specifying a reportProgressTimeout field in your pipeline.
//
// If a task runner does not report its status after 5 minutes, AWS Data Pipeline
// assumes that the task runner is unable to process the task and reassigns the
// task in a subsequent response to PollForTask. Task runners should call ReportTaskProgress
// every 60 seconds.
//
// POST / HTTP/1.1 Content-Type: application/x-amz-json-1.1 X-Amz-Target:
// DataPipeline.ReportTaskProgress Content-Length: 832 Host:
// datapipeline.us-east-1.amazonaws.com X-Amz-Date: Mon, 12 Nov 2012 17:49:52 GMT
// Authorization: AuthParams
//
// {"taskId":
// "aaGgHT4LuH0T0Y0oLrJRjas5qH0d8cDPADxqq3tn+zCWGELkCdV2JprLreXm1oxeP5EFZHFLJ69kjSsLYE0iYHYBYVGBrB+E/pYq7ANEEeGJFnSBMRiXZVA+8UJ3OzcInvXeinqBmBaKwii7hnnKb/AXjXiNTXyxgydX1KAyg1AxkwBYG4cfPYMZbuEbQJFJvv5C/2+GVXz1w94nKYTeUeepwUOFOuRLS6JVtZoYwpF56E+Yfk1IcGpFOvCZ01B4Bkuu7x3J+MD/j6kJgZLAgbCJQtI3eiW3kdGmX0p0I2BdY1ZsX6b4UiSvM3OMj6NEHJCJL4E0ZfitnhCoe24Kvjo6C2hFbZq+ei/HPgSXBQMSagkr4vS9c0ChzxH2+LNYvec6bY4kymkaZI1dvOzmpa0FcnGf5AjSK4GpsViZ/ujz6zxFv81qBXzjF0/4M1775rjV1VUdyKaixiA/sJiACNezqZqETidp8d24BDPRhGsj6pBCrnelqGFrk/gXEXUsJ+xwMifRC8UVwiKekpAvHUywVk7Ku4jH/n3i2VoLRP6FXwpUbelu34iiZ9czpXyLtyPKwxa87dlrnRVURwkcVjOt2Mcrcaqe+cbWHvNRhyrPkkdfSF3ac8/wfgVbXvLEB2k9mKc67aD9rvdc1PKX09Tk8BKklsMTpZ3TRCd4NzQlJKigMe8Jat9+1tKj4Ole5ZzW6uyTu2s2iFjEV8KXu4MaiRJyNKCdKeGhhZWY37Qk4NBK4Ppgu+C6Y41dpfOh288SLDEVx0/UySlqOEdhba7c6BiPp5r3hKj3mk9lFy5OYp1aoGLeeFmjXveTnPdf2gkWqXXg7AUbJ7jEs1F0lKZQg4szep2gcKyAJXgvXLfJJHcha8Lfb/Ee7wYmyOcAaRpDBoFNSbtoVXar46teIrpho+ZDvynUXvU0grHWGOk=:wn3SgymHZM99bEXAMPLE",
// "fields": [ {"key": "percentComplete", "stringValue": "50"} ] }
//
// x-amzn-RequestId: 640bd023-0775-11e2-af6f-6bc7a6be60d9 Content-Type:
// application/x-amz-json-1.1 Content-Length: 18 Date: Mon, 12 Nov 2012 17:50:53
// GMT
//
// {"canceled": false}
func datapipeline_ReportTaskProgress(cfg aws.Config, client *datapipeline.Client) {
	input := &datapipeline.ReportTaskProgressInput{
		// TaskId: *string, // Required
	}

	if len(_datapipelineTaskId) > 0 {
		input.TaskId = aws.String(_datapipelineTaskId)
	}
	if len(_datapipelineFields) > 0 {
		if err := assignInputField(input, "Fields", _datapipelineFields); err != nil {
			log.Errorf("invalid --fields: %s", err.Error())
			return
		}
	}

	if resp, err := client.ReportTaskProgress(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Task runners call ReportTaskRunnerHeartbeat every 15 minutes to indicate that
// they are operational. If the AWS Data Pipeline Task Runner is launched on a
// resource managed by AWS Data Pipeline, the web service can use this call to
// detect when the task runner application has failed and restart a new instance.
//
// POST / HTTP/1.1 Content-Type: application/x-amz-json-1.1 X-Amz-Target:
// DataPipeline.ReportTaskRunnerHeartbeat Content-Length: 84 Host:
// datapipeline.us-east-1.amazonaws.com X-Amz-Date: Mon, 12 Nov 2012 17:49:52 GMT
// Authorization: AuthParams
//
// {"taskrunnerId": "1234567890", "workerGroup": "wg-12345", "hostname":
// "example.com"}
//
// Status: x-amzn-RequestId: b3104dc5-0734-11e2-af6f-6bc7a6be60d9 Content-Type:
// application/x-amz-json-1.1 Content-Length: 20 Date: Mon, 12 Nov 2012 17:50:53
// GMT
//
// {"terminate": false}
func datapipeline_ReportTaskRunnerHeartbeat(cfg aws.Config, client *datapipeline.Client) {
	input := &datapipeline.ReportTaskRunnerHeartbeatInput{
		// TaskrunnerId: *string, // Required
	}

	if len(_datapipelineTaskrunnerId) > 0 {
		input.TaskrunnerId = aws.String(_datapipelineTaskrunnerId)
	}
	if len(_datapipelineHostname) > 0 {
		input.Hostname = aws.String(_datapipelineHostname)
	}
	if len(_datapipelineWorkerGroup) > 0 {
		input.WorkerGroup = aws.String(_datapipelineWorkerGroup)
	}

	if resp, err := client.ReportTaskRunnerHeartbeat(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Requests that the status of the specified physical or logical pipeline objects
// be updated in the specified pipeline. This update might not occur immediately,
// but is eventually consistent. The status that can be set depends on the type of
// object (for example, DataNode or Activity). You cannot perform this operation on
// FINISHED pipelines and attempting to do so returns InvalidRequestException .
//
// POST / HTTP/1.1 Content-Type: application/x-amz-json-1.1 X-Amz-Target:
// DataPipeline.SetStatus Content-Length: 100 Host:
// datapipeline.us-east-1.amazonaws.com X-Amz-Date: Mon, 12 Nov 2012 17:49:52 GMT
// Authorization: AuthParams
//
// {"pipelineId": "df-0634701J7KEXAMPLE", "objectIds": ["o-08600941GHJWMBR9E2"],
// "status": "pause"}
//
// x-amzn-RequestId: e83b8ab7-076a-11e2-af6f-6bc7a6be60d9 Content-Type:
// application/x-amz-json-1.1 Content-Length: 0 Date: Mon, 12 Nov 2012 17:50:53 GMT
//
// Unexpected response: 200, OK, undefined
func datapipeline_SetStatus(cfg aws.Config, client *datapipeline.Client) {
	input := &datapipeline.SetStatusInput{
		// ObjectIds: []string, // Required
		// PipelineId: *string, // Required
		// Status: *string, // Required
	}

	if len(_datapipelineObjectIds) > 0 {
		input.ObjectIds = append([]string(nil), _datapipelineObjectIds...)
	}
	if len(_datapipelinePipelineId) > 0 {
		input.PipelineId = aws.String(_datapipelinePipelineId)
	}
	if len(_datapipelineStatus) > 0 {
		input.Status = aws.String(_datapipelineStatus)
	}

	if resp, err := client.SetStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Task runners call SetTaskStatus to notify AWS Data Pipeline that a task is
// completed and provide information about the final status. A task runner makes
// this call regardless of whether the task was sucessful. A task runner does not
// need to call SetTaskStatus for tasks that are canceled by the web service
// during a call to ReportTaskProgress.
//
// POST / HTTP/1.1 Content-Type: application/x-amz-json-1.1 X-Amz-Target:
// DataPipeline.SetTaskStatus Content-Length: 847 Host:
// datapipeline.us-east-1.amazonaws.com X-Amz-Date: Mon, 12 Nov 2012 17:49:52 GMT
// Authorization: AuthParams
//
// {"taskId":
// "aaGgHT4LuH0T0Y0oLrJRjas5qH0d8cDPADxqq3tn+zCWGELkCdV2JprLreXm1oxeP5EFZHFLJ69kjSsLYE0iYHYBYVGBrB+E/pYq7ANEEeGJFnSBMRiXZVA+8UJ3OzcInvXeinqBmBaKwii7hnnKb/AXjXiNTXyxgydX1KAyg1AxkwBYG4cfPYMZbuEbQJFJvv5C/2+GVXz1w94nKYTeUeepwUOFOuRLS6JVtZoYwpF56E+Yfk1IcGpFOvCZ01B4Bkuu7x3J+MD/j6kJgZLAgbCJQtI3eiW3kdGmX0p0I2BdY1ZsX6b4UiSvM3OMj6NEHJCJL4E0ZfitnhCoe24Kvjo6C2hFbZq+ei/HPgSXBQMSagkr4vS9c0ChzxH2+LNYvec6bY4kymkaZI1dvOzmpa0FcnGf5AjSK4GpsViZ/ujz6zxFv81qBXzjF0/4M1775rjV1VUdyKaixiA/sJiACNezqZqETidp8d24BDPRhGsj6pBCrnelqGFrk/gXEXUsJ+xwMifRC8UVwiKekpAvHUywVk7Ku4jH/n3i2VoLRP6FXwpUbelu34iiZ9czpXyLtyPKwxa87dlrnRVURwkcVjOt2Mcrcaqe+cbWHvNRhyrPkkdfSF3ac8/wfgVbXvLEB2k9mKc67aD9rvdc1PKX09Tk8BKklsMTpZ3TRCd4NzQlJKigMe8Jat9+1tKj4Ole5ZzW6uyTu2s2iFjEV8KXu4MaiRJyNKCdKeGhhZWY37Qk4NBK4Ppgu+C6Y41dpfOh288SLDEVx0/UySlqOEdhba7c6BiPp5r3hKj3mk9lFy5OYp1aoGLeeFmjXveTnPdf2gkWqXXg7AUbJ7jEs1F0lKZQg4szep2gcKyAJXgvXLfJJHcha8Lfb/Ee7wYmyOcAaRpDBoFNSbtoVXar46teIrpho+ZDvynUXvU0grHWGOk=:wn3SgymHZM99bEXAMPLE",
// "taskStatus": "FINISHED"}
//
// x-amzn-RequestId: 8c8deb53-0788-11e2-af9c-6bc7a6be6qr8 Content-Type:
// application/x-amz-json-1.1 Content-Length: 0 Date: Mon, 12 Nov 2012 17:50:53 GMT
//
// {}
func datapipeline_SetTaskStatus(cfg aws.Config, client *datapipeline.Client) {
	input := &datapipeline.SetTaskStatusInput{
		// TaskId: *string, // Required
		// TaskStatus: types.TaskStatus, // Required
	}

	if len(_datapipelineTaskId) > 0 {
		input.TaskId = aws.String(_datapipelineTaskId)
	}
	if len(_datapipelineTaskStatus) > 0 {
		if err := assignInputField(input, "TaskStatus", _datapipelineTaskStatus); err != nil {
			log.Errorf("invalid --task-status: %s", err.Error())
			return
		}
	}
	if len(_datapipelineErrorId) > 0 {
		input.ErrorId = aws.String(_datapipelineErrorId)
	}
	if len(_datapipelineErrorMessage) > 0 {
		input.ErrorMessage = aws.String(_datapipelineErrorMessage)
	}
	if len(_datapipelineErrorStackTrace) > 0 {
		input.ErrorStackTrace = aws.String(_datapipelineErrorStackTrace)
	}

	if resp, err := client.SetTaskStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Validates the specified pipeline definition to ensure that it is well formed
// and can be run without error.
//
// Example 1 This example sets an valid pipeline configuration and returns
// success.
//
// POST / HTTP/1.1 Content-Type: application/x-amz-json-1.1 X-Amz-Target:
// DataPipeline.ValidatePipelineDefinition Content-Length: 936 Host:
// datapipeline.us-east-1.amazonaws.com X-Amz-Date: Mon, 12 Nov 2012 17:49:52 GMT
// Authorization: AuthParams
//
// {"pipelineId": "df-06372391ZG65EXAMPLE", "pipelineObjects": [ {"id": "Default",
// "name": "Default", "fields": [ {"key": "workerGroup", "stringValue":
// "MyworkerGroup"} ] }, {"id": "Schedule", "name": "Schedule", "fields": [ {"key":
// "startDateTime", "stringValue": "2012-09-25T17:00:00"}, {"key": "type",
// "stringValue": "Schedule"}, {"key": "period", "stringValue": "1 hour"}, {"key":
// "endDateTime", "stringValue": "2012-09-25T18:00:00"} ] }, {"id": "SayHello",
// "name": "SayHello", "fields": [ {"key": "type", "stringValue":
// "ShellCommandActivity"}, {"key": "command", "stringValue": "echo hello"},
// {"key": "parent", "refValue": "Default"}, {"key": "schedule", "refValue":
// "Schedule"}
//
// ] } ] }
//
// x-amzn-RequestId: 92c9f347-0776-11e2-8a14-21bb8a1f50ef Content-Type:
// application/x-amz-json-1.1 Content-Length: 18 Date: Mon, 12 Nov 2012 17:50:53
// GMT
//
// {"errored": false}
//
// Example 2 This example sets an invalid pipeline configuration and returns the
// associated set of validation errors.
//
// POST / HTTP/1.1 Content-Type: application/x-amz-json-1.1 X-Amz-Target:
// DataPipeline.ValidatePipelineDefinition Content-Length: 903 Host:
// datapipeline.us-east-1.amazonaws.com X-Amz-Date: Mon, 12 Nov 2012 17:49:52 GMT
// Authorization: AuthParams
//
// {"pipelineId": "df-06372391ZG65EXAMPLE", "pipelineObjects": [ {"id": "Default",
// "name": "Default", "fields": [ {"key": "workerGroup", "stringValue":
// "MyworkerGroup"} ] }, {"id": "Schedule", "name": "Schedule", "fields": [ {"key":
// "startDateTime", "stringValue": "bad-time"}, {"key": "type", "stringValue":
// "Schedule"}, {"key": "period", "stringValue": "1 hour"}, {"key": "endDateTime",
// "stringValue": "2012-09-25T18:00:00"} ] }, {"id": "SayHello", "name":
// "SayHello", "fields": [ {"key": "type", "stringValue": "ShellCommandActivity"},
// {"key": "command", "stringValue": "echo hello"}, {"key": "parent", "refValue":
// "Default"}, {"key": "schedule", "refValue": "Schedule"}
//
// ] } ] }
//
// x-amzn-RequestId: 496a1f5a-0e6a-11e2-a61c-bd6312c92ddd Content-Type:
// application/x-amz-json-1.1 Content-Length: 278 Date: Mon, 12 Nov 2012 17:50:53
// GMT
//
// {"errored": true, "validationErrors": [ {"errors": ["INVALID_FIELD_VALUE:
// 'startDateTime' value must be a literal datetime value."], "id": "Schedule"} ] }
func datapipeline_ValidatePipelineDefinition(cfg aws.Config, client *datapipeline.Client) {
	input := &datapipeline.ValidatePipelineDefinitionInput{
		// PipelineId: *string, // Required
		// PipelineObjects: []types.PipelineObject, // Required
	}

	if len(_datapipelinePipelineId) > 0 {
		input.PipelineId = aws.String(_datapipelinePipelineId)
	}
	if len(_datapipelinePipelineObjects) > 0 {
		if err := assignInputField(input, "PipelineObjects", _datapipelinePipelineObjects); err != nil {
			log.Errorf("invalid --pipeline-objects: %s", err.Error())
			return
		}
	}
	if len(_datapipelineParameterObjects) > 0 {
		if err := assignInputField(input, "ParameterObjects", _datapipelineParameterObjects); err != nil {
			log.Errorf("invalid --parameter-objects: %s", err.Error())
			return
		}
	}
	if len(_datapipelineParameterValues) > 0 {
		if err := assignInputField(input, "ParameterValues", _datapipelineParameterValues); err != nil {
			log.Errorf("invalid --parameter-values: %s", err.Error())
			return
		}
	}

	if resp, err := client.ValidatePipelineDefinition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_datapipelineCmd)
	_datapipelineCmd.Flags().SortFlags = false

	_datapipelineCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_datapipelineCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_datapipelineCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_datapipelineCmd.Flags().StringVarP(&_datapipelineCancelActive, "cancel-active", "", "", "Cancel Active")
	_datapipelineCmd.Flags().StringVarP(&_datapipelineDescription, "description", "", "", "Description")
	_datapipelineCmd.Flags().StringVarP(&_datapipelineErrorId, "error-id", "", "", "Error ID")
	_datapipelineCmd.Flags().StringVarP(&_datapipelineErrorMessage, "error-message", "", "", "Error Message")
	_datapipelineCmd.Flags().StringVarP(&_datapipelineErrorStackTrace, "error-stack-trace", "", "", "Error Stack Trace")
	_datapipelineCmd.Flags().StringVarP(&_datapipelineEvaluateExpressions, "evaluate-expressions", "", "", "Evaluate Expressions")
	_datapipelineCmd.Flags().StringVarP(&_datapipelineExpression, "expression", "", "", "Expression")
	_datapipelineCmd.Flags().StringVarP(&_datapipelineFields, "fields", "", "", "Fields")
	_datapipelineCmd.Flags().StringVarP(&_datapipelineHostname, "hostname", "", "", "Hostname")
	_datapipelineCmd.Flags().StringVarP(&_datapipelineInstanceIdentity, "instance-identity", "", "", "Instance Identity")
	_datapipelineCmd.Flags().StringVarP(&_datapipelineLimit, "limit", "", "", "Limit")
	_datapipelineCmd.Flags().StringVarP(&_datapipelineMarker, "marker", "", "", "Marker")
	_datapipelineCmd.Flags().StringVarP(&_datapipelineName, "name", "", "", "Name")
	_datapipelineCmd.Flags().StringVarP(&_datapipelineObjectId, "object-id", "", "", "Object ID")
	_datapipelineCmd.Flags().StringSliceVarP(&_datapipelineObjectIds, "object-ids", "", nil, "Object Ids")
	_datapipelineCmd.Flags().StringVarP(&_datapipelineParameterObjects, "parameter-objects", "", "", "Parameter Objects")
	_datapipelineCmd.Flags().StringVarP(&_datapipelineParameterValues, "parameter-values", "", "", "Parameter Values")
	_datapipelineCmd.Flags().StringVarP(&_datapipelinePipelineId, "pipeline-id", "", "", "Pipeline ID")
	_datapipelineCmd.Flags().StringSliceVarP(&_datapipelinePipelineIds, "pipeline-ids", "", nil, "Pipeline Ids")
	_datapipelineCmd.Flags().StringVarP(&_datapipelinePipelineObjects, "pipeline-objects", "", "", "Pipeline Objects")
	_datapipelineCmd.Flags().StringVarP(&_datapipelineQuery, "query", "", "", "Query")
	_datapipelineCmd.Flags().StringVarP(&_datapipelineSphere, "sphere", "", "", "Sphere")
	_datapipelineCmd.Flags().StringVarP(&_datapipelineStartTimestamp, "start-timestamp", "", "", "Start Timestamp")
	_datapipelineCmd.Flags().StringVarP(&_datapipelineStatus, "status", "", "", "Status")
	_datapipelineCmd.Flags().StringSliceVarP(&_datapipelineTagKeys, "tag-keys", "", nil, "Tag Keys")
	_datapipelineCmd.Flags().StringVarP(&_datapipelineTags, "tags", "", "", "Tags")
	_datapipelineCmd.Flags().StringVarP(&_datapipelineTaskId, "task-id", "", "", "Task ID")
	_datapipelineCmd.Flags().StringVarP(&_datapipelineTaskStatus, "task-status", "", "", "Task Status")
	_datapipelineCmd.Flags().StringVarP(&_datapipelineTaskrunnerId, "taskrunner-id", "", "", "Taskrunner ID")
	_datapipelineCmd.Flags().StringVarP(&_datapipelineUniqueId, "unique-id", "", "", "Unique ID")
	_datapipelineCmd.Flags().StringVarP(&_datapipelineVersion, "version", "", "", "Version")
	_datapipelineCmd.Flags().StringVarP(&_datapipelineWorkerGroup, "worker-group", "", "", "Worker Group")

	_datapipelineCmd.Flags().BoolVarP(&_datapipelineActivatePipeline, "activate-pipeline", "", false, "Activate Pipeline")
	_datapipelineCmd.Flags().BoolVarP(&_datapipelineAddTags, "add-tags", "", false, "Add Tags")
	_datapipelineCmd.Flags().BoolVarP(&_datapipelineCreatePipeline, "create-pipeline", "", false, "Create Pipeline")
	_datapipelineCmd.Flags().BoolVarP(&_datapipelineDeactivatePipeline, "deactivate-pipeline", "", false, "Deactivate Pipeline")
	_datapipelineCmd.Flags().BoolVarP(&_datapipelineDeletePipeline, "delete-pipeline", "", false, "Delete Pipeline")
	_datapipelineCmd.Flags().BoolVarP(&_datapipelineDescribeObjects, "describe-objects", "", false, "Describe Objects")
	_datapipelineCmd.Flags().BoolVarP(&_datapipelineDescribePipelines, "describe-pipelines", "", false, "Describe Pipelines")
	_datapipelineCmd.Flags().BoolVarP(&_datapipelineEvaluateExpression, "evaluate-expression", "", false, "Evaluate Expression")
	_datapipelineCmd.Flags().BoolVarP(&_datapipelineGetPipelineDefinition, "get-pipeline-definition", "", false, "Get Pipeline Definition")
	_datapipelineCmd.Flags().BoolVarP(&_datapipelineListPipelines, "list-pipelines", "", false, "List Pipelines")
	_datapipelineCmd.Flags().BoolVarP(&_datapipelinePollForTask, "poll-for-task", "", false, "Poll For Task")
	_datapipelineCmd.Flags().BoolVarP(&_datapipelinePutPipelineDefinition, "put-pipeline-definition", "", false, "Put Pipeline Definition")
	_datapipelineCmd.Flags().BoolVarP(&_datapipelineQueryObjects, "query-objects", "", false, "Query Objects")
	_datapipelineCmd.Flags().BoolVarP(&_datapipelineRemoveTags, "remove-tags", "", false, "Remove Tags")
	_datapipelineCmd.Flags().BoolVarP(&_datapipelineReportTaskProgress, "report-task-progress", "", false, "Report Task Progress")
	_datapipelineCmd.Flags().BoolVarP(&_datapipelineReportTaskRunnerHeartbeat, "report-task-runner-heartbeat", "", false, "Report Task Runner Heartbeat")
	_datapipelineCmd.Flags().BoolVarP(&_datapipelineSetStatus, "set-status", "", false, "Set Status")
	_datapipelineCmd.Flags().BoolVarP(&_datapipelineSetTaskStatus, "set-task-status", "", false, "Set Task Status")
	_datapipelineCmd.Flags().BoolVarP(&_datapipelineValidatePipelineDefinition, "validate-pipeline-definition", "", false, "Validate Pipeline Definition")

}
