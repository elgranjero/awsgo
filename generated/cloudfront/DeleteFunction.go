package cloudfront

// DeleteFunction is generated as a reference stub.
// Executable command wiring lives under cmd/cloudfront.go.
//
// Deletes a CloudFront function.
//
// You cannot delete a function if it's associated with a cache behavior. First,
// update your distributions to remove the function association from all cache
// behaviors, then delete the function.
//
// To delete a function, you must provide the function's name and version ( ETag
// value). To get these values, you can use ListFunctions and DescribeFunction .
