package datapipeline

// DeletePipeline is generated as a reference stub.
// Executable command wiring lives under cmd/datapipeline.go.
//
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
