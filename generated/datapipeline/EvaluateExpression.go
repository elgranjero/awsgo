package datapipeline

// EvaluateExpression is generated as a reference stub.
// Executable command wiring lives under cmd/datapipeline.go.
//
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
