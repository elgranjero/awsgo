package cloudfront

// TestFunction is generated as a reference stub.
// Executable command wiring lives under cmd/cloudfront.go.
//
// Tests a CloudFront function.
//
// To test a function, you provide an event object that represents an HTTP request
// or response that your CloudFront distribution could receive in production.
// CloudFront runs the function, passing it the event object that you provided, and
// returns the function's result (the modified event object) in the response. The
// response also contains function logs and error messages, if any exist. For more
// information about testing functions, see [Testing functions]in the Amazon CloudFront Developer
// Guide.
//
// To test a function, you provide the function's name and version ( ETag value)
// along with the event object. To get the function's name and version, you can use
// ListFunctions and DescribeFunction .
//
// [Testing functions]: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/managing-functions.html#test-function
