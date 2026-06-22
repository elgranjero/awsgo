package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
)

var fields_add_tags = []leanruntime.Field{
	{Name: "LoadBalancerNames", Flag: "load-balancer-names", Type: "[]string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_apply_security_groups_to_load_balancer = []leanruntime.Field{
	{Name: "LoadBalancerName", Flag: "load-balancer-name", Type: "*string", Required: true},
	{Name: "SecurityGroups", Flag: "security-groups", Type: "[]string", Required: true},
}

var fields_attach_load_balancer_to_subnets = []leanruntime.Field{
	{Name: "LoadBalancerName", Flag: "load-balancer-name", Type: "*string", Required: true},
	{Name: "Subnets", Flag: "subnets", Type: "[]string", Required: true},
}

var fields_configure_health_check = []leanruntime.Field{
	{Name: "HealthCheck", Flag: "health-check", Type: "*types.HealthCheck", Required: true},
	{Name: "LoadBalancerName", Flag: "load-balancer-name", Type: "*string", Required: true},
}

var fields_create_app_cookie_stickiness_policy = []leanruntime.Field{
	{Name: "CookieName", Flag: "cookie-name", Type: "*string", Required: true},
	{Name: "LoadBalancerName", Flag: "load-balancer-name", Type: "*string", Required: true},
	{Name: "PolicyName", Flag: "policy-name", Type: "*string", Required: true},
}

var fields_create_lb_cookie_stickiness_policy = []leanruntime.Field{
	{Name: "CookieExpirationPeriod", Flag: "cookie-expiration-period", Type: "*int64", Required: false},
	{Name: "LoadBalancerName", Flag: "load-balancer-name", Type: "*string", Required: true},
	{Name: "PolicyName", Flag: "policy-name", Type: "*string", Required: true},
}

var fields_create_load_balancer = []leanruntime.Field{
	{Name: "AvailabilityZones", Flag: "availability-zones", Type: "[]string", Required: false},
	{Name: "Listeners", Flag: "listeners", Type: "[]types.Listener", Required: true},
	{Name: "LoadBalancerName", Flag: "load-balancer-name", Type: "*string", Required: true},
	{Name: "Scheme", Flag: "scheme", Type: "*string", Required: false},
	{Name: "SecurityGroups", Flag: "security-groups", Type: "[]string", Required: false},
	{Name: "Subnets", Flag: "subnets", Type: "[]string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_load_balancer_listeners = []leanruntime.Field{
	{Name: "Listeners", Flag: "listeners", Type: "[]types.Listener", Required: true},
	{Name: "LoadBalancerName", Flag: "load-balancer-name", Type: "*string", Required: true},
}

var fields_create_load_balancer_policy = []leanruntime.Field{
	{Name: "LoadBalancerName", Flag: "load-balancer-name", Type: "*string", Required: true},
	{Name: "PolicyAttributes", Flag: "policy-attributes", Type: "[]types.PolicyAttribute", Required: false},
	{Name: "PolicyName", Flag: "policy-name", Type: "*string", Required: true},
	{Name: "PolicyTypeName", Flag: "policy-type-name", Type: "*string", Required: true},
}

var fields_delete_load_balancer = []leanruntime.Field{
	{Name: "LoadBalancerName", Flag: "load-balancer-name", Type: "*string", Required: true},
}

var fields_delete_load_balancer_listeners = []leanruntime.Field{
	{Name: "LoadBalancerName", Flag: "load-balancer-name", Type: "*string", Required: true},
	{Name: "LoadBalancerPorts", Flag: "load-balancer-ports", Type: "[]int32", Required: true},
}

var fields_delete_load_balancer_policy = []leanruntime.Field{
	{Name: "LoadBalancerName", Flag: "load-balancer-name", Type: "*string", Required: true},
	{Name: "PolicyName", Flag: "policy-name", Type: "*string", Required: true},
}

var fields_deregister_instances_from_load_balancer = []leanruntime.Field{
	{Name: "Instances", Flag: "instances", Type: "[]types.Instance", Required: true},
	{Name: "LoadBalancerName", Flag: "load-balancer-name", Type: "*string", Required: true},
}

var fields_describe_account_limits = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
}

var fields_describe_instance_health = []leanruntime.Field{
	{Name: "Instances", Flag: "instances", Type: "[]types.Instance", Required: false},
	{Name: "LoadBalancerName", Flag: "load-balancer-name", Type: "*string", Required: true},
}

var fields_describe_load_balancer_attributes = []leanruntime.Field{
	{Name: "LoadBalancerName", Flag: "load-balancer-name", Type: "*string", Required: true},
}

var fields_describe_load_balancer_policies = []leanruntime.Field{
	{Name: "LoadBalancerName", Flag: "load-balancer-name", Type: "*string", Required: false},
	{Name: "PolicyNames", Flag: "policy-names", Type: "[]string", Required: false},
}

var fields_describe_load_balancer_policy_types = []leanruntime.Field{
	{Name: "PolicyTypeNames", Flag: "policy-type-names", Type: "[]string", Required: false},
}

var fields_describe_load_balancers = []leanruntime.Field{
	{Name: "LoadBalancerNames", Flag: "load-balancer-names", Type: "[]string", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
}

var fields_describe_tags = []leanruntime.Field{
	{Name: "LoadBalancerNames", Flag: "load-balancer-names", Type: "[]string", Required: true},
}

var fields_detach_load_balancer_from_subnets = []leanruntime.Field{
	{Name: "LoadBalancerName", Flag: "load-balancer-name", Type: "*string", Required: true},
	{Name: "Subnets", Flag: "subnets", Type: "[]string", Required: true},
}

var fields_disable_availability_zones_for_load_balancer = []leanruntime.Field{
	{Name: "AvailabilityZones", Flag: "availability-zones", Type: "[]string", Required: true},
	{Name: "LoadBalancerName", Flag: "load-balancer-name", Type: "*string", Required: true},
}

var fields_enable_availability_zones_for_load_balancer = []leanruntime.Field{
	{Name: "AvailabilityZones", Flag: "availability-zones", Type: "[]string", Required: true},
	{Name: "LoadBalancerName", Flag: "load-balancer-name", Type: "*string", Required: true},
}

var fields_modify_load_balancer_attributes = []leanruntime.Field{
	{Name: "LoadBalancerAttributes", Flag: "load-balancer-attributes", Type: "*types.LoadBalancerAttributes", Required: true},
	{Name: "LoadBalancerName", Flag: "load-balancer-name", Type: "*string", Required: true},
}

var fields_register_instances_with_load_balancer = []leanruntime.Field{
	{Name: "Instances", Flag: "instances", Type: "[]types.Instance", Required: true},
	{Name: "LoadBalancerName", Flag: "load-balancer-name", Type: "*string", Required: true},
}

var fields_remove_tags = []leanruntime.Field{
	{Name: "LoadBalancerNames", Flag: "load-balancer-names", Type: "[]string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.TagKeyOnly", Required: true},
}

var fields_set_load_balancer_listener_ssl_certificate = []leanruntime.Field{
	{Name: "LoadBalancerName", Flag: "load-balancer-name", Type: "*string", Required: true},
	{Name: "LoadBalancerPort", Flag: "load-balancer-port", Type: "int32", Required: true},
	{Name: "SSLCertificateId", Flag: "ssl-certificate-id", Type: "*string", Required: true},
}

var fields_set_load_balancer_policies_for_backend_server = []leanruntime.Field{
	{Name: "InstancePort", Flag: "instance-port", Type: "*int32", Required: true},
	{Name: "LoadBalancerName", Flag: "load-balancer-name", Type: "*string", Required: true},
	{Name: "PolicyNames", Flag: "policy-names", Type: "[]string", Required: true},
}

var fields_set_load_balancer_policies_of_listener = []leanruntime.Field{
	{Name: "LoadBalancerName", Flag: "load-balancer-name", Type: "*string", Required: true},
	{Name: "LoadBalancerPort", Flag: "load-balancer-port", Type: "int32", Required: true},
	{Name: "PolicyNames", Flag: "policy-names", Type: "[]string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"add-tags": {
			Name:   "add-tags",
			Fields: fields_add_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddTagsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_tags, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddTags(ctx, input)
			},
		},
		"apply-security-groups-to-load-balancer": {
			Name:   "apply-security-groups-to-load-balancer",
			Fields: fields_apply_security_groups_to_load_balancer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ApplySecurityGroupsToLoadBalancerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_apply_security_groups_to_load_balancer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ApplySecurityGroupsToLoadBalancer(ctx, input)
			},
		},
		"attach-load-balancer-to-subnets": {
			Name:   "attach-load-balancer-to-subnets",
			Fields: fields_attach_load_balancer_to_subnets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AttachLoadBalancerToSubnetsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_attach_load_balancer_to_subnets, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AttachLoadBalancerToSubnets(ctx, input)
			},
		},
		"configure-health-check": {
			Name:   "configure-health-check",
			Fields: fields_configure_health_check,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ConfigureHealthCheckInput{}
				if _, err := leanruntime.ApplyInput(input, fields_configure_health_check, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ConfigureHealthCheck(ctx, input)
			},
		},
		"create-app-cookie-stickiness-policy": {
			Name:   "create-app-cookie-stickiness-policy",
			Fields: fields_create_app_cookie_stickiness_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAppCookieStickinessPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_app_cookie_stickiness_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAppCookieStickinessPolicy(ctx, input)
			},
		},
		"create-lb-cookie-stickiness-policy": {
			Name:   "create-lb-cookie-stickiness-policy",
			Fields: fields_create_lb_cookie_stickiness_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLBCookieStickinessPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_lb_cookie_stickiness_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLBCookieStickinessPolicy(ctx, input)
			},
		},
		"create-load-balancer": {
			Name:   "create-load-balancer",
			Fields: fields_create_load_balancer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLoadBalancerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_load_balancer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLoadBalancer(ctx, input)
			},
		},
		"create-load-balancer-listeners": {
			Name:   "create-load-balancer-listeners",
			Fields: fields_create_load_balancer_listeners,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLoadBalancerListenersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_load_balancer_listeners, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLoadBalancerListeners(ctx, input)
			},
		},
		"create-load-balancer-policy": {
			Name:   "create-load-balancer-policy",
			Fields: fields_create_load_balancer_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLoadBalancerPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_load_balancer_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLoadBalancerPolicy(ctx, input)
			},
		},
		"delete-load-balancer": {
			Name:   "delete-load-balancer",
			Fields: fields_delete_load_balancer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteLoadBalancerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_load_balancer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteLoadBalancer(ctx, input)
			},
		},
		"delete-load-balancer-listeners": {
			Name:   "delete-load-balancer-listeners",
			Fields: fields_delete_load_balancer_listeners,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteLoadBalancerListenersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_load_balancer_listeners, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteLoadBalancerListeners(ctx, input)
			},
		},
		"delete-load-balancer-policy": {
			Name:   "delete-load-balancer-policy",
			Fields: fields_delete_load_balancer_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteLoadBalancerPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_load_balancer_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteLoadBalancerPolicy(ctx, input)
			},
		},
		"deregister-instances-from-load-balancer": {
			Name:   "deregister-instances-from-load-balancer",
			Fields: fields_deregister_instances_from_load_balancer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeregisterInstancesFromLoadBalancerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deregister_instances_from_load_balancer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeregisterInstancesFromLoadBalancer(ctx, input)
			},
		},
		"describe-account-limits": {
			Name:   "describe-account-limits",
			Fields: fields_describe_account_limits,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAccountLimitsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_account_limits, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAccountLimits(ctx, input)
			},
		},
		"describe-instance-health": {
			Name:   "describe-instance-health",
			Fields: fields_describe_instance_health,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeInstanceHealthInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_instance_health, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeInstanceHealth(ctx, input)
			},
		},
		"describe-load-balancer-attributes": {
			Name:   "describe-load-balancer-attributes",
			Fields: fields_describe_load_balancer_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeLoadBalancerAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_load_balancer_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeLoadBalancerAttributes(ctx, input)
			},
		},
		"describe-load-balancer-policies": {
			Name:   "describe-load-balancer-policies",
			Fields: fields_describe_load_balancer_policies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeLoadBalancerPoliciesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_load_balancer_policies, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeLoadBalancerPolicies(ctx, input)
			},
		},
		"describe-load-balancer-policy-types": {
			Name:   "describe-load-balancer-policy-types",
			Fields: fields_describe_load_balancer_policy_types,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeLoadBalancerPolicyTypesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_load_balancer_policy_types, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeLoadBalancerPolicyTypes(ctx, input)
			},
		},
		"describe-load-balancers": {
			Name:   "describe-load-balancers",
			Fields: fields_describe_load_balancers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeLoadBalancersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_load_balancers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeLoadBalancers(ctx, input)
				}
				var results []*svc.DescribeLoadBalancersOutput
				p := svc.NewDescribeLoadBalancersPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-tags": {
			Name:   "describe-tags",
			Fields: fields_describe_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTagsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_tags, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeTags(ctx, input)
			},
		},
		"detach-load-balancer-from-subnets": {
			Name:   "detach-load-balancer-from-subnets",
			Fields: fields_detach_load_balancer_from_subnets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DetachLoadBalancerFromSubnetsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_detach_load_balancer_from_subnets, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DetachLoadBalancerFromSubnets(ctx, input)
			},
		},
		"disable-availability-zones-for-load-balancer": {
			Name:   "disable-availability-zones-for-load-balancer",
			Fields: fields_disable_availability_zones_for_load_balancer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableAvailabilityZonesForLoadBalancerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_availability_zones_for_load_balancer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableAvailabilityZonesForLoadBalancer(ctx, input)
			},
		},
		"enable-availability-zones-for-load-balancer": {
			Name:   "enable-availability-zones-for-load-balancer",
			Fields: fields_enable_availability_zones_for_load_balancer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableAvailabilityZonesForLoadBalancerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_availability_zones_for_load_balancer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableAvailabilityZonesForLoadBalancer(ctx, input)
			},
		},
		"modify-load-balancer-attributes": {
			Name:   "modify-load-balancer-attributes",
			Fields: fields_modify_load_balancer_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyLoadBalancerAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_load_balancer_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyLoadBalancerAttributes(ctx, input)
			},
		},
		"register-instances-with-load-balancer": {
			Name:   "register-instances-with-load-balancer",
			Fields: fields_register_instances_with_load_balancer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterInstancesWithLoadBalancerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_instances_with_load_balancer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterInstancesWithLoadBalancer(ctx, input)
			},
		},
		"remove-tags": {
			Name:   "remove-tags",
			Fields: fields_remove_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveTagsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_tags, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveTags(ctx, input)
			},
		},
		"set-load-balancer-listener-ssl-certificate": {
			Name:   "set-load-balancer-listener-ssl-certificate",
			Fields: fields_set_load_balancer_listener_ssl_certificate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetLoadBalancerListenerSSLCertificateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_load_balancer_listener_ssl_certificate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetLoadBalancerListenerSSLCertificate(ctx, input)
			},
		},
		"set-load-balancer-policies-for-backend-server": {
			Name:   "set-load-balancer-policies-for-backend-server",
			Fields: fields_set_load_balancer_policies_for_backend_server,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetLoadBalancerPoliciesForBackendServerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_load_balancer_policies_for_backend_server, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetLoadBalancerPoliciesForBackendServer(ctx, input)
			},
		},
		"set-load-balancer-policies-of-listener": {
			Name:   "set-load-balancer-policies-of-listener",
			Fields: fields_set_load_balancer_policies_of_listener,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetLoadBalancerPoliciesOfListenerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_load_balancer_policies_of_listener, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetLoadBalancerPoliciesOfListener(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("elasticloadbalancing", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
