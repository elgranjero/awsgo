package elasticloadbalancing

// SetLoadBalancerPoliciesForBackendServer is generated as a reference stub.
// Executable command wiring lives under cmd/elasticloadbalancing.go.
//
// Replaces the set of policies associated with the specified port on which the
// EC2 instance is listening with a new set of policies. At this time, only the
// back-end server authentication policy type can be applied to the instance ports;
// this policy type is composed of multiple public key policies.
//
// Each time you use SetLoadBalancerPoliciesForBackendServer to enable the
// policies, use the PolicyNames parameter to list the policies that you want to
// enable.
//
// You can use DescribeLoadBalancers or DescribeLoadBalancerPolicies to verify that the policy is associated with the EC2 instance.
//
// For more information about enabling back-end instance authentication, see [Configure Back-end Instance Authentication] in
// the Classic Load Balancers Guide. For more information about Proxy Protocol, see
// [Configure Proxy Protocol Support]in the Classic Load Balancers Guide.
//
// [Configure Back-end Instance Authentication]: https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/elb-create-https-ssl-load-balancer.html#configure_backendauth_clt
// [Configure Proxy Protocol Support]: https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/enable-proxy-protocol.html
