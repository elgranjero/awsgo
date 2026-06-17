package datapipeline

// QueryObjects is generated as a reference stub.
// Executable command wiring lives under cmd/datapipeline.go.
//
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
