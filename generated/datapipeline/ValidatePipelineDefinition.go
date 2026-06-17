package datapipeline

// ValidatePipelineDefinition is generated as a reference stub.
// Executable command wiring lives under cmd/datapipeline.go.
//
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
