package datapipeline

// ListPipelines is generated as a reference stub.
// Executable command wiring lives under cmd/datapipeline.go.
//
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
