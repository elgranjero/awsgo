package cloudfront

// CreateKeyGroup is generated as a reference stub.
// Executable command wiring lives under cmd/cloudfront.go.
//
// Creates a key group that you can use with [CloudFront signed URLs and signed cookies].
//
// To create a key group, you must specify at least one public key for the key
// group. After you create a key group, you can reference it from one or more cache
// behaviors. When you reference a key group in a cache behavior, CloudFront
// requires signed URLs or signed cookies for all requests that match the cache
// behavior. The URLs or cookies must be signed with a private key whose
// corresponding public key is in the key group. The signed URL or cookie contains
// information about which public key CloudFront should use to verify the
// signature. For more information, see [Serving private content]in the Amazon CloudFront Developer Guide.
//
// [Serving private content]: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PrivateContent.html
// [CloudFront signed URLs and signed cookies]: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PrivateContent.html
