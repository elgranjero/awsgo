package wafregional

// UpdateByteMatchSet is generated as a reference stub.
// Executable command wiring lives under cmd/wafregional.go.
//
// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Inserts or deletes ByteMatchTuple objects (filters) in a ByteMatchSet. For each ByteMatchTuple object,
// you specify the following values:
//
// - Whether to insert or delete the object from the array. If you want to
// change a ByteMatchSetUpdate object, you delete the existing object and add a
// new one.
//
// - The part of a web request that you want AWS WAF to inspect, such as a query
// string or the value of the User-Agent header.
//
// - The bytes (typically a string that corresponds with ASCII characters) that
// you want AWS WAF to look for. For more information, including how you specify
// the values for the AWS WAF API and the AWS CLI or SDKs, see TargetString in
// the ByteMatchTupledata type.
//
// - Where to look, such as at the beginning or the end of a query string.
//
// - Whether to perform any conversions on the request, such as converting it to
// lowercase, before inspecting it for the specified string.
//
// For example, you can add a ByteMatchSetUpdate object that matches web requests
// in which User-Agent headers contain the string BadBot . You can then configure
// AWS WAF to block those requests.
//
// To create and configure a ByteMatchSet , perform the following steps:
//
// - Create a ByteMatchSet. For more information, see CreateByteMatchSet.
//
// - Use GetChangeTokento get the change token that you provide in the ChangeToken parameter of
// an UpdateByteMatchSet request.
//
// - Submit an UpdateByteMatchSet request to specify the part of the request that
// you want AWS WAF to inspect (for example, the header or the URI) and the value
// that you want AWS WAF to watch for.
//
// For more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the [AWS WAF Developer Guide].
//
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/
