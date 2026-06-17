package wafregional

// CreateByteMatchSet is generated as a reference stub.
// Executable command wiring lives under cmd/wafregional.go.
//
// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Creates a ByteMatchSet . You then use UpdateByteMatchSet to identify the part of a web request
// that you want AWS WAF to inspect, such as the values of the User-Agent header
// or the query string. For example, you can create a ByteMatchSet that matches
// any requests with User-Agent headers that contain the string BadBot . You can
// then configure AWS WAF to reject those requests.
//
// To create and configure a ByteMatchSet , perform the following steps:
//
// - Use GetChangeTokento get the change token that you provide in the ChangeToken parameter of
// a CreateByteMatchSet request.
//
// - Submit a CreateByteMatchSet request.
//
// - Use GetChangeToken to get the change token that you provide in the
// ChangeToken parameter of an UpdateByteMatchSet request.
//
// - Submit an UpdateByteMatchSetrequest to specify the part of the request that you want AWS WAF
// to inspect (for example, the header or the URI) and the value that you want AWS
// WAF to watch for.
//
// For more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the [AWS WAF Developer Guide].
//
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/
