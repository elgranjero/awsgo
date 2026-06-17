package cloudfront

// CreateFunction is generated as a reference stub.
// Executable command wiring lives under cmd/cloudfront.go.
//
// Creates a CloudFront function.
//
// To create a function, you provide the function code and some configuration
// information about the function. The response contains an Amazon Resource Name
// (ARN) that uniquely identifies the function.
//
// When you create a function, it's in the DEVELOPMENT stage. In this stage, you
// can test the function with TestFunction , and update it with UpdateFunction .
//
// When you're ready to use your function with a CloudFront distribution, use
// PublishFunction to copy the function from the DEVELOPMENT stage to LIVE . When
// it's live, you can attach the function to a distribution's cache behavior, using
// the function's ARN.
