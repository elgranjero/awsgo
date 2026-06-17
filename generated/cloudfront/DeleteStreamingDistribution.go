package cloudfront

// DeleteStreamingDistribution is generated as a reference stub.
// Executable command wiring lives under cmd/cloudfront.go.
//
// Delete a streaming distribution. To delete an RTMP distribution using the
// CloudFront API, perform the following steps.
//
// To delete an RTMP distribution using the CloudFront API:
//
// - Disable the RTMP distribution.
//
// - Submit a GET Streaming Distribution Config request to get the current
// configuration and the Etag header for the distribution.
//
// - Update the XML document that was returned in the response to your GET
// Streaming Distribution Config request to change the value of Enabled to false .
//
// - Submit a PUT Streaming Distribution Config request to update the
// configuration for your distribution. In the request body, include the XML
// document that you updated in Step 3. Then set the value of the HTTP If-Match
// header to the value of the ETag header that CloudFront returned when you
// submitted the GET Streaming Distribution Config request in Step 2.
//
// - Review the response to the PUT Streaming Distribution Config request to
// confirm that the distribution was successfully disabled.
//
// - Submit a GET Streaming Distribution Config request to confirm that your
// changes have propagated. When propagation is complete, the value of Status is
// Deployed .
//
// - Submit a DELETE Streaming Distribution request. Set the value of the HTTP
// If-Match header to the value of the ETag header that CloudFront returned when
// you submitted the GET Streaming Distribution Config request in Step 2.
//
// - Review the response to your DELETE Streaming Distribution request to confirm
// that the distribution was successfully deleted.
//
// For information about deleting a distribution using the CloudFront console, see [Deleting a Distribution]
// in the Amazon CloudFront Developer Guide.
//
// [Deleting a Distribution]: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/HowToDeleteDistribution.html
