package wafregional

// GetChangeToken is generated as a reference stub.
// Executable command wiring lives under cmd/wafregional.go.
//
// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// When you want to create, update, or delete AWS WAF objects, get a change token
// and include the change token in the create, update, or delete request. Change
// tokens ensure that your application doesn't submit conflicting requests to AWS
// WAF.
//
// Each create, update, or delete request must use a unique change token. If your
// application submits a GetChangeToken request and then submits a second
// GetChangeToken request before submitting a create, update, or delete request,
// the second GetChangeToken request returns the same value as the first
// GetChangeToken request.
//
// When you use a change token in a create, update, or delete request, the status
// of the change token changes to PENDING , which indicates that AWS WAF is
// propagating the change to all AWS WAF servers. Use GetChangeTokenStatus to
// determine the status of your change token.
//
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
