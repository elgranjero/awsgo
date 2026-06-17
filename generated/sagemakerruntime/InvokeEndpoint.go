package sagemakerruntime

// InvokeEndpoint is generated as a reference stub.
// Executable command wiring lives under cmd/sagemakerruntime.go.
//
// After you deploy a model into production using Amazon SageMaker AI hosting
// services, your client applications use this API to get inferences from the model
// hosted at the specified endpoint.
//
// For an overview of Amazon SageMaker AI, see [How It Works].
//
// Amazon SageMaker AI strips all POST headers except those supported by the API.
// Amazon SageMaker AI might add additional headers. You should not rely on the
// behavior of headers outside those enumerated in the request syntax.
//
// Calls to InvokeEndpoint are authenticated by using Amazon Web Services
// Signature Version 4. For information, see [Authenticating Requests (Amazon Web Services Signature Version 4)]in the Amazon S3 API Reference.
//
// A customer's model containers must respond to requests within 60 seconds. The
// model itself can have a maximum processing time of 60 seconds before responding
// to invocations. If your model is going to take 50-60 seconds of processing time,
// the SDK socket timeout should be set to be 70 seconds.
//
// Endpoints are scoped to an individual account, and are not public. The URL does
// not contain the account ID, but Amazon SageMaker AI determines the account ID
// from the authentication token that is supplied by the caller.
//
// [How It Works]: https://docs.aws.amazon.com/sagemaker/latest/dg/how-it-works.html
// [Authenticating Requests (Amazon Web Services Signature Version 4)]: https://docs.aws.amazon.com/AmazonS3/latest/API/sig-v4-authenticating-requests.html
