package datapipeline

// DescribeObjects is generated as a reference stub.
// Executable command wiring lives under cmd/datapipeline.go.
//
// Gets the object definitions for a set of objects associated with the pipeline.
// Object definitions are composed of a set of fields that define the properties of
// the object.
//
// POST / HTTP/1.1 Content-Type: application/x-amz-json-1.1 X-Amz-Target:
// DataPipeline.DescribeObjects Content-Length: 98 Host:
// datapipeline.us-east-1.amazonaws.com X-Amz-Date: Mon, 12 Nov 2012 17:49:52 GMT
// Authorization: AuthParams
//
// {"pipelineId": "df-06372391ZG65EXAMPLE", "objectIds": ["Schedule"],
// "evaluateExpressions": true}
//
// x-amzn-RequestId: 4c18ea5d-0777-11e2-8a14-21bb8a1f50ef Content-Type:
// application/x-amz-json-1.1 Content-Length: 1488 Date: Mon, 12 Nov 2012 17:50:53
// GMT
//
// {"hasMoreResults": false, "pipelineObjects": [ {"fields": [ {"key":
// "startDateTime", "stringValue": "2012-12-12T00:00:00"}, {"key": "parent",
// "refValue": "Default"}, {"key": "(at)sphere", "stringValue": "COMPONENT"}, {"key":
// "type", "stringValue": "Schedule"}, {"key": "period", "stringValue": "1 hour"},
// {"key": "endDateTime", "stringValue": "2012-12-21T18:00:00"}, {"key":
// "(at)version", "stringValue": "1"}, {"key": "(at)status", "stringValue": "PENDING"},
// {"key": "(at)pipelineId", "stringValue": "df-06372391ZG65EXAMPLE"} ], "id":
// "Schedule", "name": "Schedule"} ] }
