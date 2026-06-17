package sagemakerruntime

// InvokeEndpointWithResponseStream is generated as a reference stub.
// Executable command wiring lives under cmd/sagemakerruntime.go.
//
// Invokes a model at the specified endpoint to return the inference response as a
// stream. The inference stream provides the response payload incrementally as a
// series of parts. Before you can get an inference stream, you must have access to
// a model that's deployed using Amazon SageMaker AI hosting services, and the
// container for that model must support inference streaming.
//
// For more information that can help you use this API, see the following sections
// in the Amazon SageMaker AI Developer Guide:
//
// - For information about how to add streaming support to a model, see [How Containers Serve Requests].
//
// - For information about how to process the streaming response, see [Invoke real-time endpoints].
//
// Before you can use this operation, your IAM permissions must allow the
// sagemaker:InvokeEndpoint action. For more information about Amazon SageMaker AI
// actions for IAM policies, see [Actions, resources, and condition keys for Amazon SageMaker AI]in the IAM Service Authorization Reference.
//
// Amazon SageMaker AI strips all POST headers except those supported by the API.
// Amazon SageMaker AI might add additional headers. You should not rely on the
// behavior of headers outside those enumerated in the request syntax.
//
// Calls to InvokeEndpointWithResponseStream are authenticated by using Amazon Web
// Services Signature Version 4. For information, see [Authenticating Requests (Amazon Web Services Signature Version 4)]in the Amazon S3 API
// Reference.
//
// [How Containers Serve Requests]: https://docs.aws.amazon.com/sagemaker/latest/dg/your-algorithms-inference-code.html#your-algorithms-inference-code-how-containe-serves-requests
// [Invoke real-time endpoints]: https://docs.aws.amazon.com/sagemaker/latest/dg/realtime-endpoints-test-endpoints.html
// [Actions, resources, and condition keys for Amazon SageMaker AI]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_amazonsagemaker.html
// [Authenticating Requests (Amazon Web Services Signature Version 4)]: https://docs.aws.amazon.com/AmazonS3/latest/API/sig-v4-authenticating-requests.html
