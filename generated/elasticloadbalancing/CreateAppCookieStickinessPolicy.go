package elasticloadbalancing

// CreateAppCookieStickinessPolicy is generated as a reference stub.
// Executable command wiring lives under cmd/elasticloadbalancing.go.
//
// Generates a stickiness policy with sticky session lifetimes that follow that of
// an application-generated cookie. This policy can be associated only with
// HTTP/HTTPS listeners.
//
// This policy is similar to the policy created by CreateLBCookieStickinessPolicy, except that the lifetime of
// the special Elastic Load Balancing cookie, AWSELB , follows the lifetime of the
// application-generated cookie specified in the policy configuration. The load
// balancer only inserts a new stickiness cookie when the application response
// includes a new application cookie.
//
// If the application cookie is explicitly removed or expires, the session stops
// being sticky until a new application cookie is issued.
//
// For more information, see [Application-Controlled Session Stickiness] in the Classic Load Balancers Guide.
//
// [Application-Controlled Session Stickiness]: https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/elb-sticky-sessions.html#enable-sticky-sessions-application
