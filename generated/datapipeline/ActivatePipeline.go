package datapipeline

// ActivatePipeline is generated as a reference stub.
// Executable command wiring lives under cmd/datapipeline.go.
//
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
