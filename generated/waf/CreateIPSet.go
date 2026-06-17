package waf

// CreateIPSet is generated as a reference stub.
// Executable command wiring lives under cmd/waf.go.
//
// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Creates an IPSet, which you use to specify which web requests that you want to allow
// or block based on the IP addresses that the requests originate from. For
// example, if you're receiving a lot of requests from one or more individual IP
// addresses or one or more ranges of IP addresses and you want to block the
// requests, you can create an IPSet that contains those IP addresses and then
// configure AWS WAF to block the requests.
//
// To create and configure an IPSet , perform the following steps:
//
// - Use GetChangeTokento get the change token that you provide in the ChangeToken parameter of
// a CreateIPSet request.
//
// - Submit a CreateIPSet request.
//
// - Use GetChangeToken to get the change token that you provide in the
// ChangeToken parameter of an UpdateIPSetrequest.
//
// - Submit an UpdateIPSet request to specify the IP addresses that you want AWS
// WAF to watch for.
//
// For more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the [AWS WAF Developer Guide].
//
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/
