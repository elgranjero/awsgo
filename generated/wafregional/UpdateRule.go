package wafregional

// UpdateRule is generated as a reference stub.
// Executable command wiring lives under cmd/wafregional.go.
//
// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Inserts or deletes Predicate objects in a Rule . Each Predicate object identifies a
// predicate, such as a ByteMatchSetor an IPSet, that specifies the web requests that you want to
// allow, block, or count. If you add more than one predicate to a Rule , a request
// must match all of the specifications to be allowed, blocked, or counted. For
// example, suppose that you add the following to a Rule :
//
// - A ByteMatchSet that matches the value BadBot in the User-Agent header
//
// - An IPSet that matches the IP address 192.0.2.44
//
// You then add the Rule to a WebACL and specify that you want to block requests
// that satisfy the Rule . For a request to be blocked, the User-Agent header in
// the request must contain the value BadBot and the request must originate from
// the IP address 192.0.2.44.
//
// To create and configure a Rule , perform the following steps:
//
// - Create and update the predicates that you want to include in the Rule .
//
// - Create the Rule . See CreateRule.
//
// - Use GetChangeToken to get the change token that you provide in the
// ChangeToken parameter of an UpdateRulerequest.
//
// - Submit an UpdateRule request to add predicates to the Rule .
//
// - Create and update a WebACL that contains the Rule . See CreateWebACL.
//
// If you want to replace one ByteMatchSet or IPSet with another, you delete the
// existing one and add the new one.
//
// For more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the [AWS WAF Developer Guide].
//
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/
