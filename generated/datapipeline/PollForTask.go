package datapipeline

// PollForTask is generated as a reference stub.
// Executable command wiring lives under cmd/datapipeline.go.
//
// Task runners call PollForTask to receive a task to perform from AWS Data
// Pipeline. The task runner specifies which tasks it can perform by setting a
// value for the workerGroup parameter. The task returned can come from any of the
// pipelines that match the workerGroup value passed in by the task runner and
// that was launched using the IAM user credentials specified by the task runner.
//
// If tasks are ready in the work queue, PollForTask returns a response
// immediately. If no tasks are available in the queue, PollForTask uses
// long-polling and holds on to a poll connection for up to a 90 seconds, during
// which time the first newly scheduled task is handed to the task runner. To
// accomodate this, set the socket timeout in your task runner to 90 seconds. The
// task runner should not call PollForTask again on the same workerGroup until it
// receives a response, and this can take up to 90 seconds.
//
// POST / HTTP/1.1 Content-Type: application/x-amz-json-1.1 X-Amz-Target:
// DataPipeline.PollForTask Content-Length: 59 Host:
// datapipeline.us-east-1.amazonaws.com X-Amz-Date: Mon, 12 Nov 2012 17:49:52 GMT
// Authorization: AuthParams
//
// {"workerGroup": "MyworkerGroup", "hostname": "example.com"}
//
// x-amzn-RequestId: 41c713d2-0775-11e2-af6f-6bc7a6be60d9 Content-Type:
// application/x-amz-json-1.1 Content-Length: 39 Date: Mon, 12 Nov 2012 17:50:53
// GMT
//
// {"taskObject": {"attemptId": "(at)SayHello_2012-12-12T00:00:00_Attempt=1",
// "objects": {"(at)SayHello_2012-12-12T00:00:00_Attempt=1": {"fields": [ {"key":
// "(at)componentParent", "refValue": "SayHello"}, {"key": "(at)scheduledStartTime",
// "stringValue": "2012-12-12T00:00:00"}, {"key": "parent", "refValue":
// "SayHello"}, {"key": "(at)sphere", "stringValue": "ATTEMPT"}, {"key":
// "workerGroup", "stringValue": "workerGroup"}, {"key": "(at)instanceParent",
// "refValue": "(at)SayHello_2012-12-12T00:00:00"}, {"key": "type", "stringValue":
// "ShellCommandActivity"}, {"key": "(at)status", "stringValue":
// "WAITING_FOR_RUNNER"}, {"key": "(at)version", "stringValue": "1"}, {"key":
// "schedule", "refValue": "Schedule"}, {"key": "(at)actualStartTime", "stringValue":
// "2012-12-13T01:40:50"}, {"key": "command", "stringValue": "echo hello"}, {"key":
// "(at)scheduledEndTime", "stringValue": "2012-12-12T01:00:00"}, {"key":
// "(at)activeInstances", "refValue": "(at)SayHello_2012-12-12T00:00:00"}, {"key":
// "(at)pipelineId", "stringValue": "df-0937003356ZJEXAMPLE"} ], "id":
// "(at)SayHello_2012-12-12T00:00:00_Attempt=1", "name":
// "(at)SayHello_2012-12-12T00:00:00_Attempt=1"} }, "pipelineId":
// "df-0937003356ZJEXAMPLE", "taskId":
// "2xaM4wRs5zOsIH+g9U3oVHfAgAlbSqU6XduncB0HhZ3xMnmvfePZPn4dIbYXHyWyRK+cU15MqDHwdrvftx/4wv+sNS4w34vJfv7QA9aOoOazW28l1GYSb2ZRR0N0paiQp+d1MhSKo10hOTWOsVK5S5Lnx9Qm6omFgXHyIvZRIvTlrQMpr1xuUrflyGOfbFOGpOLpvPE172MYdqpZKnbSS4TcuqgQKSWV2833fEubI57DPOP7ghWa2TcYeSIv4pdLYG53fTuwfbnbdc98g2LNUQzSVhSnt7BoqyNwht2aQ6b/UHg9A80+KVpuXuqmz3m1MXwHFgxjdmuesXNOrrlGpeLCcRWD+aGo0RN1NqhQRzNAig8V4GlaPTQzMsRCljKqvrIyAoP3Tt2XEGsHkkQo12rEX8Z90957XX2qKRwhruwYzqGkSLWjINoLdAxUJdpRXRc5DJTrBd3D5mdzn7kY1l7NEh4kFHJDt3Cx4Z3Mk8MYCACyCk/CEyy9DwuPi66cLz0NBcgbCM5LKjTBOwo1m+am+pvM1kSposE9FPP1+RFGb8k6jQBTJx3TRz1yKilnGXQTZ5xvdOFpJrklIT0OXP1MG3+auM9FlJA+1dX90QoNJE5z7axmK//MOGXUdkqFe2kiDkorqjxwDvc0Js9pVKfKvAmW8YqUbmI9l0ERpWCXXnLVHNmPWz3jaPY+OBAmuJWDmxB/Z8p94aEDg4BVXQ7LvsKQ3DLYhaB7yJ390CJT+i0mm+EBqY60V6YikPSWDFrYQ/NPi2b1DgE19mX8zHqw8qprIl4yh1Ckx2Iige4En/N5ktOoIxnASxAw/TzcE2skxdw5KlHDF+UTj71m16CR/dIaKlXijlfNlNzUBo/bNSadCQn3G5NoO501wPKI:XO50TgDNyo8EXAMPLE/g==:1"}
// }
