package kinesisvideo

// TagStream is generated as a reference stub.
// Executable command wiring lives under cmd/kinesisvideo.go.
//
// Adds one or more tags to a stream. A tag is a key-value pair (the value is
// optional) that you can define and assign to Amazon Web Services resources. If
// you specify a tag that already exists, the tag value is replaced with the value
// that you specify in the request. For more information, see [Using Cost Allocation Tags]in the Billing and
// Cost Management and Cost Management User Guide.
//
// You must provide either the StreamName or the StreamARN .
//
// This operation requires permission for the KinesisVideo:TagStream action.
//
// A Kinesis video stream can support up to 50 tags.
//
// [Using Cost Allocation Tags]: https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html
