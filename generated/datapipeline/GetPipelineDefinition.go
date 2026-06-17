package datapipeline

// GetPipelineDefinition is generated as a reference stub.
// Executable command wiring lives under cmd/datapipeline.go.
//
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
