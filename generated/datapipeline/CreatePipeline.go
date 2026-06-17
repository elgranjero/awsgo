package datapipeline

// CreatePipeline is generated as a reference stub.
// Executable command wiring lives under cmd/datapipeline.go.
//
// Creates a new, empty pipeline. Use PutPipelineDefinition to populate the pipeline.
//
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
