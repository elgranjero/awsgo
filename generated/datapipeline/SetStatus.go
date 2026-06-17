package datapipeline

// SetStatus is generated as a reference stub.
// Executable command wiring lives under cmd/datapipeline.go.
//
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
