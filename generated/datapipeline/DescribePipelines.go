package datapipeline

// DescribePipelines is generated as a reference stub.
// Executable command wiring lives under cmd/datapipeline.go.
//
// Retrieves metadata about one or more pipelines. The information retrieved
// includes the name of the pipeline, the pipeline identifier, its current state,
// and the user account that owns the pipeline. Using account credentials, you can
// retrieve metadata about pipelines that you or your IAM users have created. If
// you are using an IAM user account, you can retrieve metadata about only those
// pipelines for which you have read permissions.
//
// To retrieve the full pipeline definition instead of metadata about the
// pipeline, call GetPipelineDefinition.
//
// POST / HTTP/1.1 Content-Type: application/x-amz-json-1.1 X-Amz-Target:
// DataPipeline.DescribePipelines Content-Length: 70 Host:
// datapipeline.us-east-1.amazonaws.com X-Amz-Date: Mon, 12 Nov 2012 17:49:52 GMT
// Authorization: AuthParams
//
// {"pipelineIds": ["df-08785951KAKJEXAMPLE"] }
//
// x-amzn-RequestId: 02870eb7-0736-11e2-af6f-6bc7a6be60d9 Content-Type:
// application/x-amz-json-1.1 Content-Length: 767 Date: Mon, 12 Nov 2012 17:50:53
// GMT
//
// {"pipelineDescriptionList": [ {"description": "This is my first pipeline",
// "fields": [ {"key": "(at)pipelineState", "stringValue": "SCHEDULED"}, {"key":
// "description", "stringValue": "This is my first pipeline"}, {"key": "name",
// "stringValue": "myPipeline"}, {"key": "(at)creationTime", "stringValue":
// "2012-12-13T01:24:06"}, {"key": "(at)id", "stringValue": "df-0937003356ZJEXAMPLE"},
// {"key": "(at)sphere", "stringValue": "PIPELINE"}, {"key": "(at)version",
// "stringValue": "1"}, {"key": "(at)userId", "stringValue": "924374875933"}, {"key":
// "(at)accountId", "stringValue": "924374875933"}, {"key": "uniqueId", "stringValue":
// "1234567890"} ], "name": "myPipeline", "pipelineId": "df-0937003356ZJEXAMPLE"} ]
// }
