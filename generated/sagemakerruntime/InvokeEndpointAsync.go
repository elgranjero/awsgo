package sagemakerruntime

// InvokeEndpointAsync is generated as a reference stub.
// Executable command wiring lives under cmd/sagemakerruntime.go.
//
// After you deploy a model into production using Amazon SageMaker AI hosting
// services, your client applications use this API to get inferences from the model
// hosted at the specified endpoint in an asynchronous manner.
//
// Inference requests sent to this API are enqueued for asynchronous processing.
// The processing of the inference request may or may not complete before you
// receive a response from this API. The response from this API will not contain
// the result of the inference request but contain information about where you can
// locate it.
//
// Amazon SageMaker AI strips all POST headers except those supported by the API.
// Amazon SageMaker AI might add additional headers. You should not rely on the
// behavior of headers outside those enumerated in the request syntax.
//
// Calls to InvokeEndpointAsync are authenticated by using Amazon Web Services
// Signature Version 4. For information, see [Authenticating Requests (Amazon Web Services Signature Version 4)]in the Amazon S3 API Reference.
//
// [Authenticating Requests (Amazon Web Services Signature Version 4)]: https://docs.aws.amazon.com/AmazonS3/latest/API/sig-v4-authenticating-requests.html
