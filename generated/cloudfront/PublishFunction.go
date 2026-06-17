package cloudfront

// PublishFunction is generated as a reference stub.
// Executable command wiring lives under cmd/cloudfront.go.
//
// Publishes a CloudFront function by copying the function code from the
// DEVELOPMENT stage to LIVE . This automatically updates all cache behaviors that
// are using this function to use the newly published copy in the LIVE stage.
//
// When a function is published to the LIVE stage, you can attach the function to
// a distribution's cache behavior, using the function's Amazon Resource Name
// (ARN).
//
// To publish a function, you must provide the function's name and version ( ETag
// value). To get these values, you can use ListFunctions and DescribeFunction .
