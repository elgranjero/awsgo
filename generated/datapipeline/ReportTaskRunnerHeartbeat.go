package datapipeline

// ReportTaskRunnerHeartbeat is generated as a reference stub.
// Executable command wiring lives under cmd/datapipeline.go.
//
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
