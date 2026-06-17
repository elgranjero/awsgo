package waf

// PutLoggingConfiguration is generated as a reference stub.
// Executable command wiring lives under cmd/waf.go.
//
// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Associates a LoggingConfiguration with a specified web ACL.
//
// You can access information about all traffic that AWS WAF inspects using the
// following steps:
//
// - Create an Amazon Kinesis Data Firehose.
//
// Create the data firehose with a PUT source and in the region that you are
//
// operating. However, if you are capturing logs for Amazon CloudFront, always
// create the firehose in US East (N. Virginia).
//
// Do not create the data firehose using a Kinesis stream as your source.
//
// - Associate that firehose to your web ACL using a PutLoggingConfiguration
// request.
//
// When you successfully enable logging using a PutLoggingConfiguration request,
// AWS WAF will create a service linked role with the necessary permissions to
// write logs to the Amazon Kinesis Data Firehose. For more information, see [Logging Web ACL Traffic Information]in
// the AWS WAF Developer Guide.
//
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
// [Logging Web ACL Traffic Information]: https://docs.aws.amazon.com/waf/latest/developerguide/logging.html
