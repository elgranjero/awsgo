package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// elasticloadbalancingCmd represents the elasticloadbalancing command
var _elasticloadbalancingCmd = &cobra.Command{
	Use:   "elasticloadbalancing",
	Short: "AWS elasticloadbalancing CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := elasticloadbalancing.NewFromConfig(cfg)
		if _elasticloadbalancingAddTags {
			elasticloadbalancing_AddTags(cfg, client)
			return
		}
		if _elasticloadbalancingApplySecurityGroupsToLoadBalancer {
			elasticloadbalancing_ApplySecurityGroupsToLoadBalancer(cfg, client)
			return
		}
		if _elasticloadbalancingAttachLoadBalancerToSubnets {
			elasticloadbalancing_AttachLoadBalancerToSubnets(cfg, client)
			return
		}
		if _elasticloadbalancingConfigureHealthCheck {
			elasticloadbalancing_ConfigureHealthCheck(cfg, client)
			return
		}
		if _elasticloadbalancingCreateAppCookieStickinessPolicy {
			elasticloadbalancing_CreateAppCookieStickinessPolicy(cfg, client)
			return
		}
		if _elasticloadbalancingCreateLBCookieStickinessPolicy {
			elasticloadbalancing_CreateLBCookieStickinessPolicy(cfg, client)
			return
		}
		if _elasticloadbalancingCreateLoadBalancer {
			elasticloadbalancing_CreateLoadBalancer(cfg, client)
			return
		}
		if _elasticloadbalancingCreateLoadBalancerListeners {
			elasticloadbalancing_CreateLoadBalancerListeners(cfg, client)
			return
		}
		if _elasticloadbalancingCreateLoadBalancerPolicy {
			elasticloadbalancing_CreateLoadBalancerPolicy(cfg, client)
			return
		}
		if _elasticloadbalancingDeleteLoadBalancer {
			elasticloadbalancing_DeleteLoadBalancer(cfg, client)
			return
		}
		if _elasticloadbalancingDeleteLoadBalancerListeners {
			elasticloadbalancing_DeleteLoadBalancerListeners(cfg, client)
			return
		}
		if _elasticloadbalancingDeleteLoadBalancerPolicy {
			elasticloadbalancing_DeleteLoadBalancerPolicy(cfg, client)
			return
		}
		if _elasticloadbalancingDeregisterInstancesFromLoadBalancer {
			elasticloadbalancing_DeregisterInstancesFromLoadBalancer(cfg, client)
			return
		}
		if _elasticloadbalancingDescribeAccountLimits {
			elasticloadbalancing_DescribeAccountLimits(cfg, client)
			return
		}
		if _elasticloadbalancingDescribeInstanceHealth {
			elasticloadbalancing_DescribeInstanceHealth(cfg, client)
			return
		}
		if _elasticloadbalancingDescribeLoadBalancerAttributes {
			elasticloadbalancing_DescribeLoadBalancerAttributes(cfg, client)
			return
		}
		if _elasticloadbalancingDescribeLoadBalancerPolicies {
			elasticloadbalancing_DescribeLoadBalancerPolicies(cfg, client)
			return
		}
		if _elasticloadbalancingDescribeLoadBalancerPolicyTypes {
			elasticloadbalancing_DescribeLoadBalancerPolicyTypes(cfg, client)
			return
		}
		if _elasticloadbalancingDescribeLoadBalancers {
			elasticloadbalancing_DescribeLoadBalancers(cfg, client)
			return
		}
		if _elasticloadbalancingDescribeTags {
			elasticloadbalancing_DescribeTags(cfg, client)
			return
		}
		if _elasticloadbalancingDetachLoadBalancerFromSubnets {
			elasticloadbalancing_DetachLoadBalancerFromSubnets(cfg, client)
			return
		}
		if _elasticloadbalancingDisableAvailabilityZonesForLoadBalancer {
			elasticloadbalancing_DisableAvailabilityZonesForLoadBalancer(cfg, client)
			return
		}
		if _elasticloadbalancingEnableAvailabilityZonesForLoadBalancer {
			elasticloadbalancing_EnableAvailabilityZonesForLoadBalancer(cfg, client)
			return
		}
		if _elasticloadbalancingModifyLoadBalancerAttributes {
			elasticloadbalancing_ModifyLoadBalancerAttributes(cfg, client)
			return
		}
		if _elasticloadbalancingRegisterInstancesWithLoadBalancer {
			elasticloadbalancing_RegisterInstancesWithLoadBalancer(cfg, client)
			return
		}
		if _elasticloadbalancingRemoveTags {
			elasticloadbalancing_RemoveTags(cfg, client)
			return
		}
		if _elasticloadbalancingSetLoadBalancerListenerSSLCertificate {
			elasticloadbalancing_SetLoadBalancerListenerSSLCertificate(cfg, client)
			return
		}
		if _elasticloadbalancingSetLoadBalancerPoliciesForBackendServer {
			elasticloadbalancing_SetLoadBalancerPoliciesForBackendServer(cfg, client)
			return
		}
		if _elasticloadbalancingSetLoadBalancerPoliciesOfListener {
			elasticloadbalancing_SetLoadBalancerPoliciesOfListener(cfg, client)
			return
		}

	},
}

var (
	_elasticloadbalancingAddTags                                 bool
	_elasticloadbalancingApplySecurityGroupsToLoadBalancer       bool
	_elasticloadbalancingAttachLoadBalancerToSubnets             bool
	_elasticloadbalancingConfigureHealthCheck                    bool
	_elasticloadbalancingCreateAppCookieStickinessPolicy         bool
	_elasticloadbalancingCreateLBCookieStickinessPolicy          bool
	_elasticloadbalancingCreateLoadBalancer                      bool
	_elasticloadbalancingCreateLoadBalancerListeners             bool
	_elasticloadbalancingCreateLoadBalancerPolicy                bool
	_elasticloadbalancingDeleteLoadBalancer                      bool
	_elasticloadbalancingDeleteLoadBalancerListeners             bool
	_elasticloadbalancingDeleteLoadBalancerPolicy                bool
	_elasticloadbalancingDeregisterInstancesFromLoadBalancer     bool
	_elasticloadbalancingDescribeAccountLimits                   bool
	_elasticloadbalancingDescribeInstanceHealth                  bool
	_elasticloadbalancingDescribeLoadBalancerAttributes          bool
	_elasticloadbalancingDescribeLoadBalancerPolicies            bool
	_elasticloadbalancingDescribeLoadBalancerPolicyTypes         bool
	_elasticloadbalancingDescribeLoadBalancers                   bool
	_elasticloadbalancingDescribeTags                            bool
	_elasticloadbalancingDetachLoadBalancerFromSubnets           bool
	_elasticloadbalancingDisableAvailabilityZonesForLoadBalancer bool
	_elasticloadbalancingEnableAvailabilityZonesForLoadBalancer  bool
	_elasticloadbalancingModifyLoadBalancerAttributes            bool
	_elasticloadbalancingRegisterInstancesWithLoadBalancer       bool
	_elasticloadbalancingRemoveTags                              bool
	_elasticloadbalancingSetLoadBalancerListenerSSLCertificate   bool
	_elasticloadbalancingSetLoadBalancerPoliciesForBackendServer bool
	_elasticloadbalancingSetLoadBalancerPoliciesOfListener       bool

	_elasticloadbalancingAvailabilityZones      []string
	_elasticloadbalancingCookieExpirationPeriod string
	_elasticloadbalancingCookieName             string
	_elasticloadbalancingHealthCheck            string
	_elasticloadbalancingInstancePort           string
	_elasticloadbalancingInstances              string
	_elasticloadbalancingListeners              string
	_elasticloadbalancingLoadBalancerAttributes string
	_elasticloadbalancingLoadBalancerName       string
	_elasticloadbalancingLoadBalancerNames      []string
	_elasticloadbalancingLoadBalancerPort       string
	_elasticloadbalancingLoadBalancerPorts      string
	_elasticloadbalancingMarker                 string
	_elasticloadbalancingPageSize               string
	_elasticloadbalancingPolicyAttributes       string
	_elasticloadbalancingPolicyName             string
	_elasticloadbalancingPolicyNames            []string
	_elasticloadbalancingPolicyTypeName         string
	_elasticloadbalancingPolicyTypeNames        []string
	_elasticloadbalancingScheme                 string
	_elasticloadbalancingSecurityGroups         []string
	_elasticloadbalancingSSLCertificateId       string
	_elasticloadbalancingSubnets                []string
	_elasticloadbalancingTags                   string
)

// Adds the specified tags to the specified load balancer. Each load balancer can
// have a maximum of 10 tags.
//
// Each tag consists of a key and an optional value. If a tag with the same key is
// already associated with the load balancer, AddTags updates its value.
//
// For more information, see [Tag Your Classic Load Balancer] in the Classic Load Balancers Guide.
//
// [Tag Your Classic Load Balancer]: https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/add-remove-tags.html
func elasticloadbalancing_AddTags(cfg aws.Config, client *elasticloadbalancing.Client) {
	input := &elasticloadbalancing.AddTagsInput{
		// LoadBalancerNames: []string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_elasticloadbalancingLoadBalancerNames) > 0 {
		input.LoadBalancerNames = append([]string(nil), _elasticloadbalancingLoadBalancerNames...)
	}
	if len(_elasticloadbalancingTags) > 0 {
		if err := assignInputField(input, "Tags", _elasticloadbalancingTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.AddTags(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates one or more security groups with your load balancer in a virtual
// private cloud (VPC). The specified security groups override the previously
// associated security groups.
//
// For more information, see [Security Groups for Load Balancers in a VPC] in the Classic Load Balancers Guide.
//
// [Security Groups for Load Balancers in a VPC]: https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/elb-security-groups.html#elb-vpc-security-groups
func elasticloadbalancing_ApplySecurityGroupsToLoadBalancer(cfg aws.Config, client *elasticloadbalancing.Client) {
	input := &elasticloadbalancing.ApplySecurityGroupsToLoadBalancerInput{
		// LoadBalancerName: *string, // Required
		// SecurityGroups: []string, // Required
	}

	if len(_elasticloadbalancingLoadBalancerName) > 0 {
		input.LoadBalancerName = aws.String(_elasticloadbalancingLoadBalancerName)
	}
	if len(_elasticloadbalancingSecurityGroups) > 0 {
		input.SecurityGroups = append([]string(nil), _elasticloadbalancingSecurityGroups...)
	}

	if resp, err := client.ApplySecurityGroupsToLoadBalancer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds one or more subnets to the set of configured subnets for the specified
// load balancer.
//
// The load balancer evenly distributes requests across all registered subnets.
// For more information, see [Add or Remove Subnets for Your Load Balancer in a VPC]in the Classic Load Balancers Guide.
//
// [Add or Remove Subnets for Your Load Balancer in a VPC]: https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/elb-manage-subnets.html
func elasticloadbalancing_AttachLoadBalancerToSubnets(cfg aws.Config, client *elasticloadbalancing.Client) {
	input := &elasticloadbalancing.AttachLoadBalancerToSubnetsInput{
		// LoadBalancerName: *string, // Required
		// Subnets: []string, // Required
	}

	if len(_elasticloadbalancingLoadBalancerName) > 0 {
		input.LoadBalancerName = aws.String(_elasticloadbalancingLoadBalancerName)
	}
	if len(_elasticloadbalancingSubnets) > 0 {
		input.Subnets = append([]string(nil), _elasticloadbalancingSubnets...)
	}

	if resp, err := client.AttachLoadBalancerToSubnets(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Specifies the health check settings to use when evaluating the health state of
// your EC2 instances.
//
// For more information, see [Configure Health Checks for Your Load Balancer] in the Classic Load Balancers Guide.
//
// [Configure Health Checks for Your Load Balancer]: https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/elb-healthchecks.html
func elasticloadbalancing_ConfigureHealthCheck(cfg aws.Config, client *elasticloadbalancing.Client) {
	input := &elasticloadbalancing.ConfigureHealthCheckInput{
		// HealthCheck: *types.HealthCheck, // Required
		// LoadBalancerName: *string, // Required
	}

	if len(_elasticloadbalancingHealthCheck) > 0 {
		if err := assignInputField(input, "HealthCheck", _elasticloadbalancingHealthCheck); err != nil {
			log.Errorf("invalid --health-check: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingLoadBalancerName) > 0 {
		input.LoadBalancerName = aws.String(_elasticloadbalancingLoadBalancerName)
	}

	if resp, err := client.ConfigureHealthCheck(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

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
func elasticloadbalancing_CreateAppCookieStickinessPolicy(cfg aws.Config, client *elasticloadbalancing.Client) {
	input := &elasticloadbalancing.CreateAppCookieStickinessPolicyInput{
		// CookieName: *string, // Required
		// LoadBalancerName: *string, // Required
		// PolicyName: *string, // Required
	}

	if len(_elasticloadbalancingCookieName) > 0 {
		input.CookieName = aws.String(_elasticloadbalancingCookieName)
	}
	if len(_elasticloadbalancingLoadBalancerName) > 0 {
		input.LoadBalancerName = aws.String(_elasticloadbalancingLoadBalancerName)
	}
	if len(_elasticloadbalancingPolicyName) > 0 {
		input.PolicyName = aws.String(_elasticloadbalancingPolicyName)
	}

	if resp, err := client.CreateAppCookieStickinessPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Generates a stickiness policy with sticky session lifetimes controlled by the
// lifetime of the browser (user-agent) or a specified expiration period. This
// policy can be associated only with HTTP/HTTPS listeners.
//
// When a load balancer implements this policy, the load balancer uses a special
// cookie to track the instance for each request. When the load balancer receives a
// request, it first checks to see if this cookie is present in the request. If so,
// the load balancer sends the request to the application server specified in the
// cookie. If not, the load balancer sends the request to a server that is chosen
// based on the existing load-balancing algorithm.
//
// A cookie is inserted into the response for binding subsequent requests from the
// same user to that server. The validity of the cookie is based on the cookie
// expiration time, which is specified in the policy configuration.
//
// For more information, see [Duration-Based Session Stickiness] in the Classic Load Balancers Guide.
//
// [Duration-Based Session Stickiness]: https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/elb-sticky-sessions.html#enable-sticky-sessions-duration
func elasticloadbalancing_CreateLBCookieStickinessPolicy(cfg aws.Config, client *elasticloadbalancing.Client) {
	input := &elasticloadbalancing.CreateLBCookieStickinessPolicyInput{
		// LoadBalancerName: *string, // Required
		// PolicyName: *string, // Required
	}

	if len(_elasticloadbalancingLoadBalancerName) > 0 {
		input.LoadBalancerName = aws.String(_elasticloadbalancingLoadBalancerName)
	}
	if len(_elasticloadbalancingPolicyName) > 0 {
		input.PolicyName = aws.String(_elasticloadbalancingPolicyName)
	}
	if len(_elasticloadbalancingCookieExpirationPeriod) > 0 {
		if err := assignInputField(input, "CookieExpirationPeriod", _elasticloadbalancingCookieExpirationPeriod); err != nil {
			log.Errorf("invalid --cookie-expiration-period: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateLBCookieStickinessPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a Classic Load Balancer.
// You can add listeners, security groups, subnets, and tags when you create your
// load balancer, or you can add them later using CreateLoadBalancerListeners, ApplySecurityGroupsToLoadBalancer, AttachLoadBalancerToSubnets, and AddTags.
//
// To describe your current load balancers, see DescribeLoadBalancers. When you are finished with a
// load balancer, you can delete it using DeleteLoadBalancer.
//
// You can create up to 20 load balancers per region per account. You can request
// an increase for the number of load balancers for your account. For more
// information, see [Limits for Your Classic Load Balancer]in the Classic Load Balancers Guide.
//
// [Limits for Your Classic Load Balancer]: https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/elb-limits.html
func elasticloadbalancing_CreateLoadBalancer(cfg aws.Config, client *elasticloadbalancing.Client) {
	input := &elasticloadbalancing.CreateLoadBalancerInput{
		// Listeners: []types.Listener, // Required
		// LoadBalancerName: *string, // Required
	}

	if len(_elasticloadbalancingListeners) > 0 {
		if err := assignInputField(input, "Listeners", _elasticloadbalancingListeners); err != nil {
			log.Errorf("invalid --listeners: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingLoadBalancerName) > 0 {
		input.LoadBalancerName = aws.String(_elasticloadbalancingLoadBalancerName)
	}
	if len(_elasticloadbalancingAvailabilityZones) > 0 {
		input.AvailabilityZones = append([]string(nil), _elasticloadbalancingAvailabilityZones...)
	}
	if len(_elasticloadbalancingScheme) > 0 {
		input.Scheme = aws.String(_elasticloadbalancingScheme)
	}
	if len(_elasticloadbalancingSecurityGroups) > 0 {
		input.SecurityGroups = append([]string(nil), _elasticloadbalancingSecurityGroups...)
	}
	if len(_elasticloadbalancingSubnets) > 0 {
		input.Subnets = append([]string(nil), _elasticloadbalancingSubnets...)
	}
	if len(_elasticloadbalancingTags) > 0 {
		if err := assignInputField(input, "Tags", _elasticloadbalancingTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateLoadBalancer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates one or more listeners for the specified load balancer. If a listener
// with the specified port does not already exist, it is created; otherwise, the
// properties of the new listener must match the properties of the existing
// listener.
//
// For more information, see [Listeners for Your Classic Load Balancer] in the Classic Load Balancers Guide.
//
// [Listeners for Your Classic Load Balancer]: https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/elb-listener-config.html
func elasticloadbalancing_CreateLoadBalancerListeners(cfg aws.Config, client *elasticloadbalancing.Client) {
	input := &elasticloadbalancing.CreateLoadBalancerListenersInput{
		// Listeners: []types.Listener, // Required
		// LoadBalancerName: *string, // Required
	}

	if len(_elasticloadbalancingListeners) > 0 {
		if err := assignInputField(input, "Listeners", _elasticloadbalancingListeners); err != nil {
			log.Errorf("invalid --listeners: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingLoadBalancerName) > 0 {
		input.LoadBalancerName = aws.String(_elasticloadbalancingLoadBalancerName)
	}

	if resp, err := client.CreateLoadBalancerListeners(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a policy with the specified attributes for the specified load balancer.
// Policies are settings that are saved for your load balancer and that can be
// applied to the listener or the application server, depending on the policy type.
func elasticloadbalancing_CreateLoadBalancerPolicy(cfg aws.Config, client *elasticloadbalancing.Client) {
	input := &elasticloadbalancing.CreateLoadBalancerPolicyInput{
		// LoadBalancerName: *string, // Required
		// PolicyName: *string, // Required
		// PolicyTypeName: *string, // Required
	}

	if len(_elasticloadbalancingLoadBalancerName) > 0 {
		input.LoadBalancerName = aws.String(_elasticloadbalancingLoadBalancerName)
	}
	if len(_elasticloadbalancingPolicyName) > 0 {
		input.PolicyName = aws.String(_elasticloadbalancingPolicyName)
	}
	if len(_elasticloadbalancingPolicyTypeName) > 0 {
		input.PolicyTypeName = aws.String(_elasticloadbalancingPolicyTypeName)
	}
	if len(_elasticloadbalancingPolicyAttributes) > 0 {
		if err := assignInputField(input, "PolicyAttributes", _elasticloadbalancingPolicyAttributes); err != nil {
			log.Errorf("invalid --policy-attributes: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateLoadBalancerPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified load balancer.
// If you are attempting to recreate a load balancer, you must reconfigure all
// settings. The DNS name associated with a deleted load balancer are no longer
// usable. The name and associated DNS record of the deleted load balancer no
// longer exist and traffic sent to any of its IP addresses is no longer delivered
// to your instances.
//
// If the load balancer does not exist or has already been deleted, the call to
// DeleteLoadBalancer still succeeds.
func elasticloadbalancing_DeleteLoadBalancer(cfg aws.Config, client *elasticloadbalancing.Client) {
	input := &elasticloadbalancing.DeleteLoadBalancerInput{
		// LoadBalancerName: *string, // Required
	}

	if len(_elasticloadbalancingLoadBalancerName) > 0 {
		input.LoadBalancerName = aws.String(_elasticloadbalancingLoadBalancerName)
	}

	if resp, err := client.DeleteLoadBalancer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified listeners from the specified load balancer.
func elasticloadbalancing_DeleteLoadBalancerListeners(cfg aws.Config, client *elasticloadbalancing.Client) {
	input := &elasticloadbalancing.DeleteLoadBalancerListenersInput{
		// LoadBalancerName: *string, // Required
		// LoadBalancerPorts: []int32, // Required
	}

	if len(_elasticloadbalancingLoadBalancerName) > 0 {
		input.LoadBalancerName = aws.String(_elasticloadbalancingLoadBalancerName)
	}
	if len(_elasticloadbalancingLoadBalancerPorts) > 0 {
		if err := assignInputField(input, "LoadBalancerPorts", _elasticloadbalancingLoadBalancerPorts); err != nil {
			log.Errorf("invalid --load-balancer-ports: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteLoadBalancerListeners(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified policy from the specified load balancer. This policy must
// not be enabled for any listeners.
func elasticloadbalancing_DeleteLoadBalancerPolicy(cfg aws.Config, client *elasticloadbalancing.Client) {
	input := &elasticloadbalancing.DeleteLoadBalancerPolicyInput{
		// LoadBalancerName: *string, // Required
		// PolicyName: *string, // Required
	}

	if len(_elasticloadbalancingLoadBalancerName) > 0 {
		input.LoadBalancerName = aws.String(_elasticloadbalancingLoadBalancerName)
	}
	if len(_elasticloadbalancingPolicyName) > 0 {
		input.PolicyName = aws.String(_elasticloadbalancingPolicyName)
	}

	if resp, err := client.DeleteLoadBalancerPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deregisters the specified instances from the specified load balancer. After the
// instance is deregistered, it no longer receives traffic from the load balancer.
//
// You can use DescribeLoadBalancers to verify that the instance is deregistered from the load balancer.
//
// For more information, see [Register or De-Register EC2 Instances] in the Classic Load Balancers Guide.
//
// [Register or De-Register EC2 Instances]: https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/elb-deregister-register-instances.html
func elasticloadbalancing_DeregisterInstancesFromLoadBalancer(cfg aws.Config, client *elasticloadbalancing.Client) {
	input := &elasticloadbalancing.DeregisterInstancesFromLoadBalancerInput{
		// Instances: []types.Instance, // Required
		// LoadBalancerName: *string, // Required
	}

	if len(_elasticloadbalancingInstances) > 0 {
		if err := assignInputField(input, "Instances", _elasticloadbalancingInstances); err != nil {
			log.Errorf("invalid --instances: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingLoadBalancerName) > 0 {
		input.LoadBalancerName = aws.String(_elasticloadbalancingLoadBalancerName)
	}

	if resp, err := client.DeregisterInstancesFromLoadBalancer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the current Elastic Load Balancing resource limits for your AWS
// account.
//
// For more information, see [Limits for Your Classic Load Balancer] in the Classic Load Balancers Guide.
//
// [Limits for Your Classic Load Balancer]: https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/elb-limits.html
func elasticloadbalancing_DescribeAccountLimits(cfg aws.Config, client *elasticloadbalancing.Client) {
	input := &elasticloadbalancing.DescribeAccountLimitsInput{}

	if len(_elasticloadbalancingMarker) > 0 {
		input.Marker = aws.String(_elasticloadbalancingMarker)
	}
	if len(_elasticloadbalancingPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _elasticloadbalancingPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeAccountLimits(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the state of the specified instances with respect to the specified
// load balancer. If no instances are specified, the call describes the state of
// all instances that are currently registered with the load balancer. If instances
// are specified, their state is returned even if they are no longer registered
// with the load balancer. The state of terminated instances is not returned.
func elasticloadbalancing_DescribeInstanceHealth(cfg aws.Config, client *elasticloadbalancing.Client) {
	input := &elasticloadbalancing.DescribeInstanceHealthInput{
		// LoadBalancerName: *string, // Required
	}

	if len(_elasticloadbalancingLoadBalancerName) > 0 {
		input.LoadBalancerName = aws.String(_elasticloadbalancingLoadBalancerName)
	}
	if len(_elasticloadbalancingInstances) > 0 {
		if err := assignInputField(input, "Instances", _elasticloadbalancingInstances); err != nil {
			log.Errorf("invalid --instances: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeInstanceHealth(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the attributes for the specified load balancer.
func elasticloadbalancing_DescribeLoadBalancerAttributes(cfg aws.Config, client *elasticloadbalancing.Client) {
	input := &elasticloadbalancing.DescribeLoadBalancerAttributesInput{
		// LoadBalancerName: *string, // Required
	}

	if len(_elasticloadbalancingLoadBalancerName) > 0 {
		input.LoadBalancerName = aws.String(_elasticloadbalancingLoadBalancerName)
	}

	if resp, err := client.DescribeLoadBalancerAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the specified policies.
// If you specify a load balancer name, the action returns the descriptions of all
// policies created for the load balancer. If you specify a policy name associated
// with your load balancer, the action returns the description of that policy. If
// you don't specify a load balancer name, the action returns descriptions of the
// specified sample policies, or descriptions of all sample policies. The names of
// the sample policies have the ELBSample- prefix.
func elasticloadbalancing_DescribeLoadBalancerPolicies(cfg aws.Config, client *elasticloadbalancing.Client) {
	input := &elasticloadbalancing.DescribeLoadBalancerPoliciesInput{}

	if len(_elasticloadbalancingLoadBalancerName) > 0 {
		input.LoadBalancerName = aws.String(_elasticloadbalancingLoadBalancerName)
	}
	if len(_elasticloadbalancingPolicyNames) > 0 {
		input.PolicyNames = append([]string(nil), _elasticloadbalancingPolicyNames...)
	}

	if resp, err := client.DescribeLoadBalancerPolicies(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the specified load balancer policy types or all load balancer policy
// types.
//
// The description of each type indicates how it can be used. For example, some
// policies can be used only with layer 7 listeners, some policies can be used only
// with layer 4 listeners, and some policies can be used only with your EC2
// instances.
//
// You can use CreateLoadBalancerPolicy to create a policy configuration for any of these policy types.
// Then, depending on the policy type, use either SetLoadBalancerPoliciesOfListeneror SetLoadBalancerPoliciesForBackendServer to set the policy.
func elasticloadbalancing_DescribeLoadBalancerPolicyTypes(cfg aws.Config, client *elasticloadbalancing.Client) {
	input := &elasticloadbalancing.DescribeLoadBalancerPolicyTypesInput{}

	if len(_elasticloadbalancingPolicyTypeNames) > 0 {
		input.PolicyTypeNames = append([]string(nil), _elasticloadbalancingPolicyTypeNames...)
	}

	if resp, err := client.DescribeLoadBalancerPolicyTypes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the specified the load balancers. If no load balancers are specified,
// the call describes all of your load balancers.
func elasticloadbalancing_DescribeLoadBalancers(cfg aws.Config, client *elasticloadbalancing.Client) {
	input := &elasticloadbalancing.DescribeLoadBalancersInput{}

	if len(_elasticloadbalancingLoadBalancerNames) > 0 {
		input.LoadBalancerNames = append([]string(nil), _elasticloadbalancingLoadBalancerNames...)
	}
	if len(_elasticloadbalancingMarker) > 0 {
		input.Marker = aws.String(_elasticloadbalancingMarker)
	}
	if len(_elasticloadbalancingPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _elasticloadbalancingPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeLoadBalancers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*elasticloadbalancing.DescribeLoadBalancersOutput
	p := elasticloadbalancing.NewDescribeLoadBalancersPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Describes the tags associated with the specified load balancers.
func elasticloadbalancing_DescribeTags(cfg aws.Config, client *elasticloadbalancing.Client) {
	input := &elasticloadbalancing.DescribeTagsInput{
		// LoadBalancerNames: []string, // Required
	}

	if len(_elasticloadbalancingLoadBalancerNames) > 0 {
		input.LoadBalancerNames = append([]string(nil), _elasticloadbalancingLoadBalancerNames...)
	}

	if resp, err := client.DescribeTags(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the specified subnets from the set of configured subnets for the load
// balancer.
//
// After a subnet is removed, all EC2 instances registered with the load balancer
// in the removed subnet go into the OutOfService state. Then, the load balancer
// balances the traffic among the remaining routable subnets.
func elasticloadbalancing_DetachLoadBalancerFromSubnets(cfg aws.Config, client *elasticloadbalancing.Client) {
	input := &elasticloadbalancing.DetachLoadBalancerFromSubnetsInput{
		// LoadBalancerName: *string, // Required
		// Subnets: []string, // Required
	}

	if len(_elasticloadbalancingLoadBalancerName) > 0 {
		input.LoadBalancerName = aws.String(_elasticloadbalancingLoadBalancerName)
	}
	if len(_elasticloadbalancingSubnets) > 0 {
		input.Subnets = append([]string(nil), _elasticloadbalancingSubnets...)
	}

	if resp, err := client.DetachLoadBalancerFromSubnets(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the specified Availability Zones from the set of Availability Zones for
// the specified load balancer in EC2-Classic or a default VPC.
//
// For load balancers in a non-default VPC, use DetachLoadBalancerFromSubnets.
//
// There must be at least one Availability Zone registered with a load balancer at
// all times. After an Availability Zone is removed, all instances registered with
// the load balancer that are in the removed Availability Zone go into the
// OutOfService state. Then, the load balancer attempts to equally balance the
// traffic among its remaining Availability Zones.
//
// For more information, see [Add or Remove Availability Zones] in the Classic Load Balancers Guide.
//
// [Add or Remove Availability Zones]: https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/enable-disable-az.html
func elasticloadbalancing_DisableAvailabilityZonesForLoadBalancer(cfg aws.Config, client *elasticloadbalancing.Client) {
	input := &elasticloadbalancing.DisableAvailabilityZonesForLoadBalancerInput{
		// AvailabilityZones: []string, // Required
		// LoadBalancerName: *string, // Required
	}

	if len(_elasticloadbalancingAvailabilityZones) > 0 {
		input.AvailabilityZones = append([]string(nil), _elasticloadbalancingAvailabilityZones...)
	}
	if len(_elasticloadbalancingLoadBalancerName) > 0 {
		input.LoadBalancerName = aws.String(_elasticloadbalancingLoadBalancerName)
	}

	if resp, err := client.DisableAvailabilityZonesForLoadBalancer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds the specified Availability Zones to the set of Availability Zones for the
// specified load balancer in EC2-Classic or a default VPC.
//
// For load balancers in a non-default VPC, use AttachLoadBalancerToSubnets.
//
// The load balancer evenly distributes requests across all its registered
// Availability Zones that contain instances. For more information, see [Add or Remove Availability Zones]in the
// Classic Load Balancers Guide.
//
// [Add or Remove Availability Zones]: https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/enable-disable-az.html
func elasticloadbalancing_EnableAvailabilityZonesForLoadBalancer(cfg aws.Config, client *elasticloadbalancing.Client) {
	input := &elasticloadbalancing.EnableAvailabilityZonesForLoadBalancerInput{
		// AvailabilityZones: []string, // Required
		// LoadBalancerName: *string, // Required
	}

	if len(_elasticloadbalancingAvailabilityZones) > 0 {
		input.AvailabilityZones = append([]string(nil), _elasticloadbalancingAvailabilityZones...)
	}
	if len(_elasticloadbalancingLoadBalancerName) > 0 {
		input.LoadBalancerName = aws.String(_elasticloadbalancingLoadBalancerName)
	}

	if resp, err := client.EnableAvailabilityZonesForLoadBalancer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the attributes of the specified load balancer.
// You can modify the load balancer attributes, such as AccessLogs ,
// ConnectionDraining , and CrossZoneLoadBalancing by either enabling or disabling
// them. Or, you can modify the load balancer attribute ConnectionSettings by
// specifying an idle connection timeout value for your load balancer.
//
// For more information, see the following in the Classic Load Balancers Guide:
//
// [Cross-Zone Load Balancing]
//
// [Connection Draining]
//
// [Access Logs]
//
// [Idle Connection Timeout]
//
// [Cross-Zone Load Balancing]: https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/enable-disable-crosszone-lb.html
// [Idle Connection Timeout]: https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/config-idle-timeout.html
// [Access Logs]: https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/access-log-collection.html
// [Connection Draining]: https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/config-conn-drain.html
func elasticloadbalancing_ModifyLoadBalancerAttributes(cfg aws.Config, client *elasticloadbalancing.Client) {
	input := &elasticloadbalancing.ModifyLoadBalancerAttributesInput{
		// LoadBalancerAttributes: *types.LoadBalancerAttributes, // Required
		// LoadBalancerName: *string, // Required
	}

	if len(_elasticloadbalancingLoadBalancerAttributes) > 0 {
		if err := assignInputField(input, "LoadBalancerAttributes", _elasticloadbalancingLoadBalancerAttributes); err != nil {
			log.Errorf("invalid --load-balancer-attributes: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingLoadBalancerName) > 0 {
		input.LoadBalancerName = aws.String(_elasticloadbalancingLoadBalancerName)
	}

	if resp, err := client.ModifyLoadBalancerAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds the specified instances to the specified load balancer.
// The instance must be a running instance in the same network as the load
// balancer (EC2-Classic or the same VPC). If you have EC2-Classic instances and a
// load balancer in a VPC with ClassicLink enabled, you can link the EC2-Classic
// instances to that VPC and then register the linked EC2-Classic instances with
// the load balancer in the VPC.
//
// Note that RegisterInstanceWithLoadBalancer completes when the request has been
// registered. Instance registration takes a little time to complete. To check the
// state of the registered instances, use DescribeLoadBalancersor DescribeInstanceHealth.
//
// After the instance is registered, it starts receiving traffic and requests from
// the load balancer. Any instance that is not in one of the Availability Zones
// registered for the load balancer is moved to the OutOfService state. If an
// Availability Zone is added to the load balancer later, any instances registered
// with the load balancer move to the InService state.
//
// To deregister instances from a load balancer, use DeregisterInstancesFromLoadBalancer.
//
// For more information, see [Register or De-Register EC2 Instances] in the Classic Load Balancers Guide.
//
// [Register or De-Register EC2 Instances]: https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/elb-deregister-register-instances.html
func elasticloadbalancing_RegisterInstancesWithLoadBalancer(cfg aws.Config, client *elasticloadbalancing.Client) {
	input := &elasticloadbalancing.RegisterInstancesWithLoadBalancerInput{
		// Instances: []types.Instance, // Required
		// LoadBalancerName: *string, // Required
	}

	if len(_elasticloadbalancingInstances) > 0 {
		if err := assignInputField(input, "Instances", _elasticloadbalancingInstances); err != nil {
			log.Errorf("invalid --instances: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingLoadBalancerName) > 0 {
		input.LoadBalancerName = aws.String(_elasticloadbalancingLoadBalancerName)
	}

	if resp, err := client.RegisterInstancesWithLoadBalancer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes one or more tags from the specified load balancer.
func elasticloadbalancing_RemoveTags(cfg aws.Config, client *elasticloadbalancing.Client) {
	input := &elasticloadbalancing.RemoveTagsInput{
		// LoadBalancerNames: []string, // Required
		// Tags: []types.TagKeyOnly, // Required
	}

	if len(_elasticloadbalancingLoadBalancerNames) > 0 {
		input.LoadBalancerNames = append([]string(nil), _elasticloadbalancingLoadBalancerNames...)
	}
	if len(_elasticloadbalancingTags) > 0 {
		if err := assignInputField(input, "Tags", _elasticloadbalancingTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.RemoveTags(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets the certificate that terminates the specified listener's SSL connections.
// The specified certificate replaces any prior certificate that was used on the
// same load balancer and port.
//
// For more information about updating your SSL certificate, see [Replace the SSL Certificate for Your Load Balancer] in the Classic
// Load Balancers Guide.
//
// [Replace the SSL Certificate for Your Load Balancer]: https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/elb-update-ssl-cert.html
func elasticloadbalancing_SetLoadBalancerListenerSSLCertificate(cfg aws.Config, client *elasticloadbalancing.Client) {
	input := &elasticloadbalancing.SetLoadBalancerListenerSSLCertificateInput{
		// LoadBalancerName: *string, // Required
		// LoadBalancerPort: int32, // Required
		// SSLCertificateId: *string, // Required
	}

	if len(_elasticloadbalancingLoadBalancerName) > 0 {
		input.LoadBalancerName = aws.String(_elasticloadbalancingLoadBalancerName)
	}
	if len(_elasticloadbalancingLoadBalancerPort) > 0 {
		if err := assignInputField(input, "LoadBalancerPort", _elasticloadbalancingLoadBalancerPort); err != nil {
			log.Errorf("invalid --load-balancer-port: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingSSLCertificateId) > 0 {
		input.SSLCertificateId = aws.String(_elasticloadbalancingSSLCertificateId)
	}

	if resp, err := client.SetLoadBalancerListenerSSLCertificate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

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
func elasticloadbalancing_SetLoadBalancerPoliciesForBackendServer(cfg aws.Config, client *elasticloadbalancing.Client) {
	input := &elasticloadbalancing.SetLoadBalancerPoliciesForBackendServerInput{
		// InstancePort: *int32, // Required
		// LoadBalancerName: *string, // Required
		// PolicyNames: []string, // Required
	}

	if len(_elasticloadbalancingInstancePort) > 0 {
		if err := assignInputField(input, "InstancePort", _elasticloadbalancingInstancePort); err != nil {
			log.Errorf("invalid --instance-port: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingLoadBalancerName) > 0 {
		input.LoadBalancerName = aws.String(_elasticloadbalancingLoadBalancerName)
	}
	if len(_elasticloadbalancingPolicyNames) > 0 {
		input.PolicyNames = append([]string(nil), _elasticloadbalancingPolicyNames...)
	}

	if resp, err := client.SetLoadBalancerPoliciesForBackendServer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Replaces the current set of policies for the specified load balancer port with
// the specified set of policies.
//
// To enable back-end server authentication, use SetLoadBalancerPoliciesForBackendServer.
//
// For more information about setting policies, see [Update the SSL Negotiation Configuration], [Duration-Based Session Stickiness], and [Application-Controlled Session Stickiness] in the Classic Load
// Balancers Guide.
//
// [Update the SSL Negotiation Configuration]: https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/ssl-config-update.html
// [Duration-Based Session Stickiness]: https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/elb-sticky-sessions.html#enable-sticky-sessions-duration
// [Application-Controlled Session Stickiness]: https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/elb-sticky-sessions.html#enable-sticky-sessions-application
func elasticloadbalancing_SetLoadBalancerPoliciesOfListener(cfg aws.Config, client *elasticloadbalancing.Client) {
	input := &elasticloadbalancing.SetLoadBalancerPoliciesOfListenerInput{
		// LoadBalancerName: *string, // Required
		// LoadBalancerPort: int32, // Required
		// PolicyNames: []string, // Required
	}

	if len(_elasticloadbalancingLoadBalancerName) > 0 {
		input.LoadBalancerName = aws.String(_elasticloadbalancingLoadBalancerName)
	}
	if len(_elasticloadbalancingLoadBalancerPort) > 0 {
		if err := assignInputField(input, "LoadBalancerPort", _elasticloadbalancingLoadBalancerPort); err != nil {
			log.Errorf("invalid --load-balancer-port: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingPolicyNames) > 0 {
		input.PolicyNames = append([]string(nil), _elasticloadbalancingPolicyNames...)
	}

	if resp, err := client.SetLoadBalancerPoliciesOfListener(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_elasticloadbalancingCmd)
	_elasticloadbalancingCmd.Flags().SortFlags = false

	_elasticloadbalancingCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_elasticloadbalancingCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_elasticloadbalancingCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_elasticloadbalancingCmd.Flags().StringSliceVarP(&_elasticloadbalancingAvailabilityZones, "availability-zones", "", nil, "Availability Zones")
	_elasticloadbalancingCmd.Flags().StringVarP(&_elasticloadbalancingCookieExpirationPeriod, "cookie-expiration-period", "", "", "Cookie Expiration Period")
	_elasticloadbalancingCmd.Flags().StringVarP(&_elasticloadbalancingCookieName, "cookie-name", "", "", "Cookie Name")
	_elasticloadbalancingCmd.Flags().StringVarP(&_elasticloadbalancingHealthCheck, "health-check", "", "", "Health Check")
	_elasticloadbalancingCmd.Flags().StringVarP(&_elasticloadbalancingInstancePort, "instance-port", "", "", "Instance Port")
	_elasticloadbalancingCmd.Flags().StringVarP(&_elasticloadbalancingInstances, "instances", "", "", "Instances")
	_elasticloadbalancingCmd.Flags().StringVarP(&_elasticloadbalancingListeners, "listeners", "", "", "Listeners")
	_elasticloadbalancingCmd.Flags().StringVarP(&_elasticloadbalancingLoadBalancerAttributes, "load-balancer-attributes", "", "", "Load Balancer Attributes")
	_elasticloadbalancingCmd.Flags().StringVarP(&_elasticloadbalancingLoadBalancerName, "load-balancer-name", "", "", "Load Balancer Name")
	_elasticloadbalancingCmd.Flags().StringSliceVarP(&_elasticloadbalancingLoadBalancerNames, "load-balancer-names", "", nil, "Load Balancer Names")
	_elasticloadbalancingCmd.Flags().StringVarP(&_elasticloadbalancingLoadBalancerPort, "load-balancer-port", "", "", "Load Balancer Port")
	_elasticloadbalancingCmd.Flags().StringVarP(&_elasticloadbalancingLoadBalancerPorts, "load-balancer-ports", "", "", "Load Balancer Ports")
	_elasticloadbalancingCmd.Flags().StringVarP(&_elasticloadbalancingMarker, "marker", "", "", "Marker")
	_elasticloadbalancingCmd.Flags().StringVarP(&_elasticloadbalancingPageSize, "page-size", "", "", "Page Size")
	_elasticloadbalancingCmd.Flags().StringVarP(&_elasticloadbalancingPolicyAttributes, "policy-attributes", "", "", "Policy Attributes")
	_elasticloadbalancingCmd.Flags().StringVarP(&_elasticloadbalancingPolicyName, "policy-name", "", "", "Policy Name")
	_elasticloadbalancingCmd.Flags().StringSliceVarP(&_elasticloadbalancingPolicyNames, "policy-names", "", nil, "Policy Names")
	_elasticloadbalancingCmd.Flags().StringVarP(&_elasticloadbalancingPolicyTypeName, "policy-type-name", "", "", "Policy Type Name")
	_elasticloadbalancingCmd.Flags().StringSliceVarP(&_elasticloadbalancingPolicyTypeNames, "policy-type-names", "", nil, "Policy Type Names")
	_elasticloadbalancingCmd.Flags().StringVarP(&_elasticloadbalancingScheme, "scheme", "", "", "Scheme")
	_elasticloadbalancingCmd.Flags().StringSliceVarP(&_elasticloadbalancingSecurityGroups, "security-groups", "", nil, "Security Groups")
	_elasticloadbalancingCmd.Flags().StringVarP(&_elasticloadbalancingSSLCertificateId, "ssl-certificate-id", "", "", "SSL Certificate ID")
	_elasticloadbalancingCmd.Flags().StringSliceVarP(&_elasticloadbalancingSubnets, "subnets", "", nil, "Subnets")
	_elasticloadbalancingCmd.Flags().StringVarP(&_elasticloadbalancingTags, "tags", "", "", "Tags")

	_elasticloadbalancingCmd.Flags().BoolVarP(&_elasticloadbalancingAddTags, "add-tags", "", false, "Add Tags")
	_elasticloadbalancingCmd.Flags().BoolVarP(&_elasticloadbalancingApplySecurityGroupsToLoadBalancer, "apply-security-groups-to-load-balancer", "", false, "Apply Security Groups To Load Balancer")
	_elasticloadbalancingCmd.Flags().BoolVarP(&_elasticloadbalancingAttachLoadBalancerToSubnets, "attach-load-balancer-to-subnets", "", false, "Attach Load Balancer To Subnets")
	_elasticloadbalancingCmd.Flags().BoolVarP(&_elasticloadbalancingConfigureHealthCheck, "configure-health-check", "", false, "Configure Health Check")
	_elasticloadbalancingCmd.Flags().BoolVarP(&_elasticloadbalancingCreateAppCookieStickinessPolicy, "create-app-cookie-stickiness-policy", "", false, "Create App Cookie Stickiness Policy")
	_elasticloadbalancingCmd.Flags().BoolVarP(&_elasticloadbalancingCreateLBCookieStickinessPolicy, "create-lb-cookie-stickiness-policy", "", false, "Create Lb Cookie Stickiness Policy")
	_elasticloadbalancingCmd.Flags().BoolVarP(&_elasticloadbalancingCreateLoadBalancer, "create-load-balancer", "", false, "Create Load Balancer")
	_elasticloadbalancingCmd.Flags().BoolVarP(&_elasticloadbalancingCreateLoadBalancerListeners, "create-load-balancer-listeners", "", false, "Create Load Balancer Listeners")
	_elasticloadbalancingCmd.Flags().BoolVarP(&_elasticloadbalancingCreateLoadBalancerPolicy, "create-load-balancer-policy", "", false, "Create Load Balancer Policy")
	_elasticloadbalancingCmd.Flags().BoolVarP(&_elasticloadbalancingDeleteLoadBalancer, "delete-load-balancer", "", false, "Delete Load Balancer")
	_elasticloadbalancingCmd.Flags().BoolVarP(&_elasticloadbalancingDeleteLoadBalancerListeners, "delete-load-balancer-listeners", "", false, "Delete Load Balancer Listeners")
	_elasticloadbalancingCmd.Flags().BoolVarP(&_elasticloadbalancingDeleteLoadBalancerPolicy, "delete-load-balancer-policy", "", false, "Delete Load Balancer Policy")
	_elasticloadbalancingCmd.Flags().BoolVarP(&_elasticloadbalancingDeregisterInstancesFromLoadBalancer, "deregister-instances-from-load-balancer", "", false, "Deregister Instances From Load Balancer")
	_elasticloadbalancingCmd.Flags().BoolVarP(&_elasticloadbalancingDescribeAccountLimits, "describe-account-limits", "", false, "Describe Account Limits")
	_elasticloadbalancingCmd.Flags().BoolVarP(&_elasticloadbalancingDescribeInstanceHealth, "describe-instance-health", "", false, "Describe Instance Health")
	_elasticloadbalancingCmd.Flags().BoolVarP(&_elasticloadbalancingDescribeLoadBalancerAttributes, "describe-load-balancer-attributes", "", false, "Describe Load Balancer Attributes")
	_elasticloadbalancingCmd.Flags().BoolVarP(&_elasticloadbalancingDescribeLoadBalancerPolicies, "describe-load-balancer-policies", "", false, "Describe Load Balancer Policies")
	_elasticloadbalancingCmd.Flags().BoolVarP(&_elasticloadbalancingDescribeLoadBalancerPolicyTypes, "describe-load-balancer-policy-types", "", false, "Describe Load Balancer Policy Types")
	_elasticloadbalancingCmd.Flags().BoolVarP(&_elasticloadbalancingDescribeLoadBalancers, "describe-load-balancers", "", false, "Describe Load Balancers")
	_elasticloadbalancingCmd.Flags().BoolVarP(&_elasticloadbalancingDescribeTags, "describe-tags", "", false, "Describe Tags")
	_elasticloadbalancingCmd.Flags().BoolVarP(&_elasticloadbalancingDetachLoadBalancerFromSubnets, "detach-load-balancer-from-subnets", "", false, "Detach Load Balancer From Subnets")
	_elasticloadbalancingCmd.Flags().BoolVarP(&_elasticloadbalancingDisableAvailabilityZonesForLoadBalancer, "disable-availability-zones-for-load-balancer", "", false, "Disable Availability Zones For Load Balancer")
	_elasticloadbalancingCmd.Flags().BoolVarP(&_elasticloadbalancingEnableAvailabilityZonesForLoadBalancer, "enable-availability-zones-for-load-balancer", "", false, "Enable Availability Zones For Load Balancer")
	_elasticloadbalancingCmd.Flags().BoolVarP(&_elasticloadbalancingModifyLoadBalancerAttributes, "modify-load-balancer-attributes", "", false, "Modify Load Balancer Attributes")
	_elasticloadbalancingCmd.Flags().BoolVarP(&_elasticloadbalancingRegisterInstancesWithLoadBalancer, "register-instances-with-load-balancer", "", false, "Register Instances With Load Balancer")
	_elasticloadbalancingCmd.Flags().BoolVarP(&_elasticloadbalancingRemoveTags, "remove-tags", "", false, "Remove Tags")
	_elasticloadbalancingCmd.Flags().BoolVarP(&_elasticloadbalancingSetLoadBalancerListenerSSLCertificate, "set-load-balancer-listener-ssl-certificate", "", false, "Set Load Balancer Listener SSL Certificate")
	_elasticloadbalancingCmd.Flags().BoolVarP(&_elasticloadbalancingSetLoadBalancerPoliciesForBackendServer, "set-load-balancer-policies-for-backend-server", "", false, "Set Load Balancer Policies For Backend Server")
	_elasticloadbalancingCmd.Flags().BoolVarP(&_elasticloadbalancingSetLoadBalancerPoliciesOfListener, "set-load-balancer-policies-of-listener", "", false, "Set Load Balancer Policies Of Listener")

}
