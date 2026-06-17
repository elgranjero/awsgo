package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// elasticloadbalancingv2Cmd represents the elasticloadbalancingv2 command
var _elasticloadbalancingv2Cmd = &cobra.Command{
	Use:   "elasticloadbalancingv2",
	Short: "AWS elasticloadbalancingv2 CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := elasticloadbalancingv2.NewFromConfig(cfg)
		if _elasticloadbalancingv2AddListenerCertificates {
			elasticloadbalancingv2_AddListenerCertificates(cfg, client)
			return
		}
		if _elasticloadbalancingv2AddTags {
			elasticloadbalancingv2_AddTags(cfg, client)
			return
		}
		if _elasticloadbalancingv2AddTrustStoreRevocations {
			elasticloadbalancingv2_AddTrustStoreRevocations(cfg, client)
			return
		}
		if _elasticloadbalancingv2CreateListener {
			elasticloadbalancingv2_CreateListener(cfg, client)
			return
		}
		if _elasticloadbalancingv2CreateLoadBalancer {
			elasticloadbalancingv2_CreateLoadBalancer(cfg, client)
			return
		}
		if _elasticloadbalancingv2CreateRule {
			elasticloadbalancingv2_CreateRule(cfg, client)
			return
		}
		if _elasticloadbalancingv2CreateTargetGroup {
			elasticloadbalancingv2_CreateTargetGroup(cfg, client)
			return
		}
		if _elasticloadbalancingv2CreateTrustStore {
			elasticloadbalancingv2_CreateTrustStore(cfg, client)
			return
		}
		if _elasticloadbalancingv2DeleteListener {
			elasticloadbalancingv2_DeleteListener(cfg, client)
			return
		}
		if _elasticloadbalancingv2DeleteLoadBalancer {
			elasticloadbalancingv2_DeleteLoadBalancer(cfg, client)
			return
		}
		if _elasticloadbalancingv2DeleteRule {
			elasticloadbalancingv2_DeleteRule(cfg, client)
			return
		}
		if _elasticloadbalancingv2DeleteSharedTrustStoreAssociation {
			elasticloadbalancingv2_DeleteSharedTrustStoreAssociation(cfg, client)
			return
		}
		if _elasticloadbalancingv2DeleteTargetGroup {
			elasticloadbalancingv2_DeleteTargetGroup(cfg, client)
			return
		}
		if _elasticloadbalancingv2DeleteTrustStore {
			elasticloadbalancingv2_DeleteTrustStore(cfg, client)
			return
		}
		if _elasticloadbalancingv2DeregisterTargets {
			elasticloadbalancingv2_DeregisterTargets(cfg, client)
			return
		}
		if _elasticloadbalancingv2DescribeAccountLimits {
			elasticloadbalancingv2_DescribeAccountLimits(cfg, client)
			return
		}
		if _elasticloadbalancingv2DescribeCapacityReservation {
			elasticloadbalancingv2_DescribeCapacityReservation(cfg, client)
			return
		}
		if _elasticloadbalancingv2DescribeListenerAttributes {
			elasticloadbalancingv2_DescribeListenerAttributes(cfg, client)
			return
		}
		if _elasticloadbalancingv2DescribeListenerCertificates {
			elasticloadbalancingv2_DescribeListenerCertificates(cfg, client)
			return
		}
		if _elasticloadbalancingv2DescribeListeners {
			elasticloadbalancingv2_DescribeListeners(cfg, client)
			return
		}
		if _elasticloadbalancingv2DescribeLoadBalancerAttributes {
			elasticloadbalancingv2_DescribeLoadBalancerAttributes(cfg, client)
			return
		}
		if _elasticloadbalancingv2DescribeLoadBalancers {
			elasticloadbalancingv2_DescribeLoadBalancers(cfg, client)
			return
		}
		if _elasticloadbalancingv2DescribeRules {
			elasticloadbalancingv2_DescribeRules(cfg, client)
			return
		}
		if _elasticloadbalancingv2DescribeSSLPolicies {
			elasticloadbalancingv2_DescribeSSLPolicies(cfg, client)
			return
		}
		if _elasticloadbalancingv2DescribeTags {
			elasticloadbalancingv2_DescribeTags(cfg, client)
			return
		}
		if _elasticloadbalancingv2DescribeTargetGroupAttributes {
			elasticloadbalancingv2_DescribeTargetGroupAttributes(cfg, client)
			return
		}
		if _elasticloadbalancingv2DescribeTargetGroups {
			elasticloadbalancingv2_DescribeTargetGroups(cfg, client)
			return
		}
		if _elasticloadbalancingv2DescribeTargetHealth {
			elasticloadbalancingv2_DescribeTargetHealth(cfg, client)
			return
		}
		if _elasticloadbalancingv2DescribeTrustStoreAssociations {
			elasticloadbalancingv2_DescribeTrustStoreAssociations(cfg, client)
			return
		}
		if _elasticloadbalancingv2DescribeTrustStoreRevocations {
			elasticloadbalancingv2_DescribeTrustStoreRevocations(cfg, client)
			return
		}
		if _elasticloadbalancingv2DescribeTrustStores {
			elasticloadbalancingv2_DescribeTrustStores(cfg, client)
			return
		}
		if _elasticloadbalancingv2GetResourcePolicy {
			elasticloadbalancingv2_GetResourcePolicy(cfg, client)
			return
		}
		if _elasticloadbalancingv2GetTrustStoreCaCertificatesBundle {
			elasticloadbalancingv2_GetTrustStoreCaCertificatesBundle(cfg, client)
			return
		}
		if _elasticloadbalancingv2GetTrustStoreRevocationContent {
			elasticloadbalancingv2_GetTrustStoreRevocationContent(cfg, client)
			return
		}
		if _elasticloadbalancingv2ModifyCapacityReservation {
			elasticloadbalancingv2_ModifyCapacityReservation(cfg, client)
			return
		}
		if _elasticloadbalancingv2ModifyIpPools {
			elasticloadbalancingv2_ModifyIpPools(cfg, client)
			return
		}
		if _elasticloadbalancingv2ModifyListener {
			elasticloadbalancingv2_ModifyListener(cfg, client)
			return
		}
		if _elasticloadbalancingv2ModifyListenerAttributes {
			elasticloadbalancingv2_ModifyListenerAttributes(cfg, client)
			return
		}
		if _elasticloadbalancingv2ModifyLoadBalancerAttributes {
			elasticloadbalancingv2_ModifyLoadBalancerAttributes(cfg, client)
			return
		}
		if _elasticloadbalancingv2ModifyRule {
			elasticloadbalancingv2_ModifyRule(cfg, client)
			return
		}
		if _elasticloadbalancingv2ModifyTargetGroup {
			elasticloadbalancingv2_ModifyTargetGroup(cfg, client)
			return
		}
		if _elasticloadbalancingv2ModifyTargetGroupAttributes {
			elasticloadbalancingv2_ModifyTargetGroupAttributes(cfg, client)
			return
		}
		if _elasticloadbalancingv2ModifyTrustStore {
			elasticloadbalancingv2_ModifyTrustStore(cfg, client)
			return
		}
		if _elasticloadbalancingv2RegisterTargets {
			elasticloadbalancingv2_RegisterTargets(cfg, client)
			return
		}
		if _elasticloadbalancingv2RemoveListenerCertificates {
			elasticloadbalancingv2_RemoveListenerCertificates(cfg, client)
			return
		}
		if _elasticloadbalancingv2RemoveTags {
			elasticloadbalancingv2_RemoveTags(cfg, client)
			return
		}
		if _elasticloadbalancingv2RemoveTrustStoreRevocations {
			elasticloadbalancingv2_RemoveTrustStoreRevocations(cfg, client)
			return
		}
		if _elasticloadbalancingv2SetIpAddressType {
			elasticloadbalancingv2_SetIpAddressType(cfg, client)
			return
		}
		if _elasticloadbalancingv2SetRulePriorities {
			elasticloadbalancingv2_SetRulePriorities(cfg, client)
			return
		}
		if _elasticloadbalancingv2SetSecurityGroups {
			elasticloadbalancingv2_SetSecurityGroups(cfg, client)
			return
		}
		if _elasticloadbalancingv2SetSubnets {
			elasticloadbalancingv2_SetSubnets(cfg, client)
			return
		}

	},
}

var (
	_elasticloadbalancingv2AddListenerCertificates           bool
	_elasticloadbalancingv2AddTags                           bool
	_elasticloadbalancingv2AddTrustStoreRevocations          bool
	_elasticloadbalancingv2CreateListener                    bool
	_elasticloadbalancingv2CreateLoadBalancer                bool
	_elasticloadbalancingv2CreateRule                        bool
	_elasticloadbalancingv2CreateTargetGroup                 bool
	_elasticloadbalancingv2CreateTrustStore                  bool
	_elasticloadbalancingv2DeleteListener                    bool
	_elasticloadbalancingv2DeleteLoadBalancer                bool
	_elasticloadbalancingv2DeleteRule                        bool
	_elasticloadbalancingv2DeleteSharedTrustStoreAssociation bool
	_elasticloadbalancingv2DeleteTargetGroup                 bool
	_elasticloadbalancingv2DeleteTrustStore                  bool
	_elasticloadbalancingv2DeregisterTargets                 bool
	_elasticloadbalancingv2DescribeAccountLimits             bool
	_elasticloadbalancingv2DescribeCapacityReservation       bool
	_elasticloadbalancingv2DescribeListenerAttributes        bool
	_elasticloadbalancingv2DescribeListenerCertificates      bool
	_elasticloadbalancingv2DescribeListeners                 bool
	_elasticloadbalancingv2DescribeLoadBalancerAttributes    bool
	_elasticloadbalancingv2DescribeLoadBalancers             bool
	_elasticloadbalancingv2DescribeRules                     bool
	_elasticloadbalancingv2DescribeSSLPolicies               bool
	_elasticloadbalancingv2DescribeTags                      bool
	_elasticloadbalancingv2DescribeTargetGroupAttributes     bool
	_elasticloadbalancingv2DescribeTargetGroups              bool
	_elasticloadbalancingv2DescribeTargetHealth              bool
	_elasticloadbalancingv2DescribeTrustStoreAssociations    bool
	_elasticloadbalancingv2DescribeTrustStoreRevocations     bool
	_elasticloadbalancingv2DescribeTrustStores               bool
	_elasticloadbalancingv2GetResourcePolicy                 bool
	_elasticloadbalancingv2GetTrustStoreCaCertificatesBundle bool
	_elasticloadbalancingv2GetTrustStoreRevocationContent    bool
	_elasticloadbalancingv2ModifyCapacityReservation         bool
	_elasticloadbalancingv2ModifyIpPools                     bool
	_elasticloadbalancingv2ModifyListener                    bool
	_elasticloadbalancingv2ModifyListenerAttributes          bool
	_elasticloadbalancingv2ModifyLoadBalancerAttributes      bool
	_elasticloadbalancingv2ModifyRule                        bool
	_elasticloadbalancingv2ModifyTargetGroup                 bool
	_elasticloadbalancingv2ModifyTargetGroupAttributes       bool
	_elasticloadbalancingv2ModifyTrustStore                  bool
	_elasticloadbalancingv2RegisterTargets                   bool
	_elasticloadbalancingv2RemoveListenerCertificates        bool
	_elasticloadbalancingv2RemoveTags                        bool
	_elasticloadbalancingv2RemoveTrustStoreRevocations       bool
	_elasticloadbalancingv2SetIpAddressType                  bool
	_elasticloadbalancingv2SetRulePriorities                 bool
	_elasticloadbalancingv2SetSecurityGroups                 bool
	_elasticloadbalancingv2SetSubnets                        bool

	_elasticloadbalancingv2Actions                                              string
	_elasticloadbalancingv2AlpnPolicy                                           []string
	_elasticloadbalancingv2Attributes                                           string
	_elasticloadbalancingv2CaCertificatesBundleS3Bucket                         string
	_elasticloadbalancingv2CaCertificatesBundleS3Key                            string
	_elasticloadbalancingv2CaCertificatesBundleS3ObjectVersion                  string
	_elasticloadbalancingv2Certificates                                         string
	_elasticloadbalancingv2Conditions                                           string
	_elasticloadbalancingv2CustomerOwnedIpv4Pool                                string
	_elasticloadbalancingv2DefaultActions                                       string
	_elasticloadbalancingv2EnablePrefixForIpv6SourceNat                         string
	_elasticloadbalancingv2EnforceSecurityGroupInboundRulesOnPrivateLinkTraffic string
	_elasticloadbalancingv2HealthCheckEnabled                                   string
	_elasticloadbalancingv2HealthCheckIntervalSeconds                           string
	_elasticloadbalancingv2HealthCheckPath                                      string
	_elasticloadbalancingv2HealthCheckPort                                      string
	_elasticloadbalancingv2HealthCheckProtocol                                  string
	_elasticloadbalancingv2HealthCheckTimeoutSeconds                            string
	_elasticloadbalancingv2HealthyThresholdCount                                string
	_elasticloadbalancingv2Include                                              string
	_elasticloadbalancingv2IpAddressType                                        string
	_elasticloadbalancingv2IpamPools                                            string
	_elasticloadbalancingv2ListenerArn                                          string
	_elasticloadbalancingv2ListenerArns                                         []string
	_elasticloadbalancingv2LoadBalancerArn                                      string
	_elasticloadbalancingv2LoadBalancerArns                                     []string
	_elasticloadbalancingv2LoadBalancerType                                     string
	_elasticloadbalancingv2Marker                                               string
	_elasticloadbalancingv2Matcher                                              string
	_elasticloadbalancingv2MinimumLoadBalancerCapacity                          string
	_elasticloadbalancingv2MutualAuthentication                                 string
	_elasticloadbalancingv2Name                                                 string
	_elasticloadbalancingv2Names                                                []string
	_elasticloadbalancingv2PageSize                                             string
	_elasticloadbalancingv2Port                                                 string
	_elasticloadbalancingv2Priority                                             string
	_elasticloadbalancingv2Protocol                                             string
	_elasticloadbalancingv2ProtocolVersion                                      string
	_elasticloadbalancingv2RemoveIpamPools                                      string
	_elasticloadbalancingv2ResetCapacityReservation                             string
	_elasticloadbalancingv2ResetTransforms                                      string
	_elasticloadbalancingv2ResourceArn                                          string
	_elasticloadbalancingv2ResourceArns                                         []string
	_elasticloadbalancingv2RevocationContents                                   string
	_elasticloadbalancingv2RevocationId                                         string
	_elasticloadbalancingv2RevocationIds                                        string
	_elasticloadbalancingv2RuleArn                                              string
	_elasticloadbalancingv2RuleArns                                             []string
	_elasticloadbalancingv2RulePriorities                                       string
	_elasticloadbalancingv2Scheme                                               string
	_elasticloadbalancingv2SecurityGroups                                       []string
	_elasticloadbalancingv2SslPolicy                                            string
	_elasticloadbalancingv2SubnetMappings                                       string
	_elasticloadbalancingv2Subnets                                              []string
	_elasticloadbalancingv2TagKeys                                              []string
	_elasticloadbalancingv2Tags                                                 string
	_elasticloadbalancingv2TargetControlPort                                    string
	_elasticloadbalancingv2TargetGroupArn                                       string
	_elasticloadbalancingv2TargetGroupArns                                      []string
	_elasticloadbalancingv2TargetType                                           string
	_elasticloadbalancingv2Targets                                              string
	_elasticloadbalancingv2Transforms                                           string
	_elasticloadbalancingv2TrustStoreArn                                        string
	_elasticloadbalancingv2TrustStoreArns                                       []string
	_elasticloadbalancingv2Type                                                 string
	_elasticloadbalancingv2UnhealthyThresholdCount                              string
	_elasticloadbalancingv2VpcId                                                string
)

// Adds the specified SSL server certificate to the certificate list for the
// specified HTTPS or TLS listener.
//
// If the certificate in already in the certificate list, the call is successful
// but the certificate is not added again.
//
// For more information, see [SSL certificates] in the Application Load Balancers Guide or [Server certificates] in the
// Network Load Balancers Guide.
//
// [Server certificates]: https://docs.aws.amazon.com/elasticloadbalancing/latest/network/tls-listener-certificates.html
// [SSL certificates]: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/https-listener-certificates.html
func elasticloadbalancingv2_AddListenerCertificates(cfg aws.Config, client *elasticloadbalancingv2.Client) {
	input := &elasticloadbalancingv2.AddListenerCertificatesInput{
		// Certificates: []types.Certificate, // Required
		// ListenerArn: *string, // Required
	}

	if len(_elasticloadbalancingv2Certificates) > 0 {
		if err := assignInputField(input, "Certificates", _elasticloadbalancingv2Certificates); err != nil {
			log.Errorf("invalid --certificates: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingv2ListenerArn) > 0 {
		input.ListenerArn = aws.String(_elasticloadbalancingv2ListenerArn)
	}

	if resp, err := client.AddListenerCertificates(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds the specified tags to the specified Elastic Load Balancing resource. You
// can tag your Application Load Balancers, Network Load Balancers, Gateway Load
// Balancers, target groups, trust stores, listeners, and rules.
//
// Each tag consists of a key and an optional value. If a resource already has a
// tag with the same key, AddTags updates its value.
func elasticloadbalancingv2_AddTags(cfg aws.Config, client *elasticloadbalancingv2.Client) {
	input := &elasticloadbalancingv2.AddTagsInput{
		// ResourceArns: []string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_elasticloadbalancingv2ResourceArns) > 0 {
		input.ResourceArns = append([]string(nil), _elasticloadbalancingv2ResourceArns...)
	}
	if len(_elasticloadbalancingv2Tags) > 0 {
		if err := assignInputField(input, "Tags", _elasticloadbalancingv2Tags); err != nil {
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

// Adds the specified revocation file to the specified trust store.
func elasticloadbalancingv2_AddTrustStoreRevocations(cfg aws.Config, client *elasticloadbalancingv2.Client) {
	input := &elasticloadbalancingv2.AddTrustStoreRevocationsInput{
		// TrustStoreArn: *string, // Required
	}

	if len(_elasticloadbalancingv2TrustStoreArn) > 0 {
		input.TrustStoreArn = aws.String(_elasticloadbalancingv2TrustStoreArn)
	}
	if len(_elasticloadbalancingv2RevocationContents) > 0 {
		if err := assignInputField(input, "RevocationContents", _elasticloadbalancingv2RevocationContents); err != nil {
			log.Errorf("invalid --revocation-contents: %s", err.Error())
			return
		}
	}

	if resp, err := client.AddTrustStoreRevocations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a listener for the specified Application Load Balancer, Network Load
// Balancer, or Gateway Load Balancer.
//
// For more information, see the following:
//
// [Listeners for your Application Load Balancers]
//
// [Listeners for your Network Load Balancers]
//
// [Listeners for your Gateway Load Balancers]
//
// This operation is idempotent, which means that it completes at most one time.
// If you attempt to create multiple listeners with the same settings, each call
// succeeds.
//
// [Listeners for your Gateway Load Balancers]: https://docs.aws.amazon.com/elasticloadbalancing/latest/gateway/gateway-listeners.html
// [Listeners for your Application Load Balancers]: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-listeners.html
// [Listeners for your Network Load Balancers]: https://docs.aws.amazon.com/elasticloadbalancing/latest/network/load-balancer-listeners.html
func elasticloadbalancingv2_CreateListener(cfg aws.Config, client *elasticloadbalancingv2.Client) {
	input := &elasticloadbalancingv2.CreateListenerInput{
		// DefaultActions: []types.Action, // Required
		// LoadBalancerArn: *string, // Required
	}

	if len(_elasticloadbalancingv2DefaultActions) > 0 {
		if err := assignInputField(input, "DefaultActions", _elasticloadbalancingv2DefaultActions); err != nil {
			log.Errorf("invalid --default-actions: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingv2LoadBalancerArn) > 0 {
		input.LoadBalancerArn = aws.String(_elasticloadbalancingv2LoadBalancerArn)
	}
	if len(_elasticloadbalancingv2AlpnPolicy) > 0 {
		input.AlpnPolicy = append([]string(nil), _elasticloadbalancingv2AlpnPolicy...)
	}
	if len(_elasticloadbalancingv2Certificates) > 0 {
		if err := assignInputField(input, "Certificates", _elasticloadbalancingv2Certificates); err != nil {
			log.Errorf("invalid --certificates: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingv2MutualAuthentication) > 0 {
		if err := assignInputField(input, "MutualAuthentication", _elasticloadbalancingv2MutualAuthentication); err != nil {
			log.Errorf("invalid --mutual-authentication: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingv2Port) > 0 {
		if err := assignInputField(input, "Port", _elasticloadbalancingv2Port); err != nil {
			log.Errorf("invalid --port: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingv2Protocol) > 0 {
		if err := assignInputField(input, "Protocol", _elasticloadbalancingv2Protocol); err != nil {
			log.Errorf("invalid --protocol: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingv2SslPolicy) > 0 {
		input.SslPolicy = aws.String(_elasticloadbalancingv2SslPolicy)
	}
	if len(_elasticloadbalancingv2Tags) > 0 {
		if err := assignInputField(input, "Tags", _elasticloadbalancingv2Tags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateListener(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Application Load Balancer, Network Load Balancer, or Gateway Load
// Balancer.
//
// For more information, see the following:
//
// [Application Load Balancers]
//
// [Network Load Balancers]
//
// [Gateway Load Balancers]
//
// This operation is idempotent, which means that it completes at most one time.
// If you attempt to create multiple load balancers with the same settings, each
// call succeeds.
//
// [Gateway Load Balancers]: https://docs.aws.amazon.com/elasticloadbalancing/latest/gateway/gateway-load-balancers.html
// [Network Load Balancers]: https://docs.aws.amazon.com/elasticloadbalancing/latest/network/network-load-balancers.html
// [Application Load Balancers]: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/application-load-balancers.html
func elasticloadbalancingv2_CreateLoadBalancer(cfg aws.Config, client *elasticloadbalancingv2.Client) {
	input := &elasticloadbalancingv2.CreateLoadBalancerInput{
		// Name: *string, // Required
	}

	if len(_elasticloadbalancingv2Name) > 0 {
		input.Name = aws.String(_elasticloadbalancingv2Name)
	}
	if len(_elasticloadbalancingv2CustomerOwnedIpv4Pool) > 0 {
		input.CustomerOwnedIpv4Pool = aws.String(_elasticloadbalancingv2CustomerOwnedIpv4Pool)
	}
	if len(_elasticloadbalancingv2EnablePrefixForIpv6SourceNat) > 0 {
		if err := assignInputField(input, "EnablePrefixForIpv6SourceNat", _elasticloadbalancingv2EnablePrefixForIpv6SourceNat); err != nil {
			log.Errorf("invalid --enable-prefix-for-ipv6-source-nat: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingv2IpAddressType) > 0 {
		if err := assignInputField(input, "IpAddressType", _elasticloadbalancingv2IpAddressType); err != nil {
			log.Errorf("invalid --ip-address-type: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingv2IpamPools) > 0 {
		if err := assignInputField(input, "IpamPools", _elasticloadbalancingv2IpamPools); err != nil {
			log.Errorf("invalid --ipam-pools: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingv2Scheme) > 0 {
		if err := assignInputField(input, "Scheme", _elasticloadbalancingv2Scheme); err != nil {
			log.Errorf("invalid --scheme: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingv2SecurityGroups) > 0 {
		input.SecurityGroups = append([]string(nil), _elasticloadbalancingv2SecurityGroups...)
	}
	if len(_elasticloadbalancingv2SubnetMappings) > 0 {
		if err := assignInputField(input, "SubnetMappings", _elasticloadbalancingv2SubnetMappings); err != nil {
			log.Errorf("invalid --subnet-mappings: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingv2Subnets) > 0 {
		input.Subnets = append([]string(nil), _elasticloadbalancingv2Subnets...)
	}
	if len(_elasticloadbalancingv2Tags) > 0 {
		if err := assignInputField(input, "Tags", _elasticloadbalancingv2Tags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingv2Type) > 0 {
		if err := assignInputField(input, "Type", _elasticloadbalancingv2Type); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
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

// Creates a rule for the specified listener. The listener must be associated with
// an Application Load Balancer.
//
// Each rule consists of a priority, one or more actions, one or more conditions,
// and up to two optional transforms. Rules are evaluated in priority order, from
// the lowest value to the highest value. When the conditions for a rule are met,
// its actions are performed. If the conditions for no rules are met, the actions
// for the default rule are performed. For more information, see [Listener rules]in the
// Application Load Balancers Guide.
//
// [Listener rules]: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-listeners.html#listener-rules
func elasticloadbalancingv2_CreateRule(cfg aws.Config, client *elasticloadbalancingv2.Client) {
	input := &elasticloadbalancingv2.CreateRuleInput{
		// Actions: []types.Action, // Required
		// Conditions: []types.RuleCondition, // Required
		// ListenerArn: *string, // Required
		// Priority: *int32, // Required
	}

	if len(_elasticloadbalancingv2Actions) > 0 {
		if err := assignInputField(input, "Actions", _elasticloadbalancingv2Actions); err != nil {
			log.Errorf("invalid --actions: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingv2Conditions) > 0 {
		if err := assignInputField(input, "Conditions", _elasticloadbalancingv2Conditions); err != nil {
			log.Errorf("invalid --conditions: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingv2ListenerArn) > 0 {
		input.ListenerArn = aws.String(_elasticloadbalancingv2ListenerArn)
	}
	if len(_elasticloadbalancingv2Priority) > 0 {
		if err := assignInputField(input, "Priority", _elasticloadbalancingv2Priority); err != nil {
			log.Errorf("invalid --priority: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingv2Tags) > 0 {
		if err := assignInputField(input, "Tags", _elasticloadbalancingv2Tags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingv2Transforms) > 0 {
		if err := assignInputField(input, "Transforms", _elasticloadbalancingv2Transforms); err != nil {
			log.Errorf("invalid --transforms: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a target group.
// For more information, see the following:
//
// [Target groups for your Application Load Balancers]
//
// [Target groups for your Network Load Balancers]
//
// [Target groups for your Gateway Load Balancers]
//
// This operation is idempotent, which means that it completes at most one time.
// If you attempt to create multiple target groups with the same settings, each
// call succeeds.
//
// [Target groups for your Gateway Load Balancers]: https://docs.aws.amazon.com/elasticloadbalancing/latest/gateway/target-groups.html
// [Target groups for your Application Load Balancers]: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-target-groups.html
// [Target groups for your Network Load Balancers]: https://docs.aws.amazon.com/elasticloadbalancing/latest/network/load-balancer-target-groups.html
func elasticloadbalancingv2_CreateTargetGroup(cfg aws.Config, client *elasticloadbalancingv2.Client) {
	input := &elasticloadbalancingv2.CreateTargetGroupInput{
		// Name: *string, // Required
	}

	if len(_elasticloadbalancingv2Name) > 0 {
		input.Name = aws.String(_elasticloadbalancingv2Name)
	}
	if len(_elasticloadbalancingv2HealthCheckEnabled) > 0 {
		if err := assignInputField(input, "HealthCheckEnabled", _elasticloadbalancingv2HealthCheckEnabled); err != nil {
			log.Errorf("invalid --health-check-enabled: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingv2HealthCheckIntervalSeconds) > 0 {
		if err := assignInputField(input, "HealthCheckIntervalSeconds", _elasticloadbalancingv2HealthCheckIntervalSeconds); err != nil {
			log.Errorf("invalid --health-check-interval-seconds: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingv2HealthCheckPath) > 0 {
		input.HealthCheckPath = aws.String(_elasticloadbalancingv2HealthCheckPath)
	}
	if len(_elasticloadbalancingv2HealthCheckPort) > 0 {
		input.HealthCheckPort = aws.String(_elasticloadbalancingv2HealthCheckPort)
	}
	if len(_elasticloadbalancingv2HealthCheckProtocol) > 0 {
		if err := assignInputField(input, "HealthCheckProtocol", _elasticloadbalancingv2HealthCheckProtocol); err != nil {
			log.Errorf("invalid --health-check-protocol: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingv2HealthCheckTimeoutSeconds) > 0 {
		if err := assignInputField(input, "HealthCheckTimeoutSeconds", _elasticloadbalancingv2HealthCheckTimeoutSeconds); err != nil {
			log.Errorf("invalid --health-check-timeout-seconds: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingv2HealthyThresholdCount) > 0 {
		if err := assignInputField(input, "HealthyThresholdCount", _elasticloadbalancingv2HealthyThresholdCount); err != nil {
			log.Errorf("invalid --healthy-threshold-count: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingv2IpAddressType) > 0 {
		if err := assignInputField(input, "IpAddressType", _elasticloadbalancingv2IpAddressType); err != nil {
			log.Errorf("invalid --ip-address-type: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingv2Matcher) > 0 {
		if err := assignInputField(input, "Matcher", _elasticloadbalancingv2Matcher); err != nil {
			log.Errorf("invalid --matcher: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingv2Port) > 0 {
		if err := assignInputField(input, "Port", _elasticloadbalancingv2Port); err != nil {
			log.Errorf("invalid --port: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingv2Protocol) > 0 {
		if err := assignInputField(input, "Protocol", _elasticloadbalancingv2Protocol); err != nil {
			log.Errorf("invalid --protocol: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingv2ProtocolVersion) > 0 {
		input.ProtocolVersion = aws.String(_elasticloadbalancingv2ProtocolVersion)
	}
	if len(_elasticloadbalancingv2Tags) > 0 {
		if err := assignInputField(input, "Tags", _elasticloadbalancingv2Tags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingv2TargetControlPort) > 0 {
		if err := assignInputField(input, "TargetControlPort", _elasticloadbalancingv2TargetControlPort); err != nil {
			log.Errorf("invalid --target-control-port: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingv2TargetType) > 0 {
		if err := assignInputField(input, "TargetType", _elasticloadbalancingv2TargetType); err != nil {
			log.Errorf("invalid --target-type: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingv2UnhealthyThresholdCount) > 0 {
		if err := assignInputField(input, "UnhealthyThresholdCount", _elasticloadbalancingv2UnhealthyThresholdCount); err != nil {
			log.Errorf("invalid --unhealthy-threshold-count: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingv2VpcId) > 0 {
		input.VpcId = aws.String(_elasticloadbalancingv2VpcId)
	}

	if resp, err := client.CreateTargetGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a trust store.
// For more information, see [Mutual TLS for Application Load Balancers].
//
// [Mutual TLS for Application Load Balancers]: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/mutual-authentication.html
func elasticloadbalancingv2_CreateTrustStore(cfg aws.Config, client *elasticloadbalancingv2.Client) {
	input := &elasticloadbalancingv2.CreateTrustStoreInput{
		// CaCertificatesBundleS3Bucket: *string, // Required
		// CaCertificatesBundleS3Key: *string, // Required
		// Name: *string, // Required
	}

	if len(_elasticloadbalancingv2CaCertificatesBundleS3Bucket) > 0 {
		input.CaCertificatesBundleS3Bucket = aws.String(_elasticloadbalancingv2CaCertificatesBundleS3Bucket)
	}
	if len(_elasticloadbalancingv2CaCertificatesBundleS3Key) > 0 {
		input.CaCertificatesBundleS3Key = aws.String(_elasticloadbalancingv2CaCertificatesBundleS3Key)
	}
	if len(_elasticloadbalancingv2Name) > 0 {
		input.Name = aws.String(_elasticloadbalancingv2Name)
	}
	if len(_elasticloadbalancingv2CaCertificatesBundleS3ObjectVersion) > 0 {
		input.CaCertificatesBundleS3ObjectVersion = aws.String(_elasticloadbalancingv2CaCertificatesBundleS3ObjectVersion)
	}
	if len(_elasticloadbalancingv2Tags) > 0 {
		if err := assignInputField(input, "Tags", _elasticloadbalancingv2Tags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateTrustStore(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified listener.
// Alternatively, your listener is deleted when you delete the load balancer to
// which it is attached.
func elasticloadbalancingv2_DeleteListener(cfg aws.Config, client *elasticloadbalancingv2.Client) {
	input := &elasticloadbalancingv2.DeleteListenerInput{
		// ListenerArn: *string, // Required
	}

	if len(_elasticloadbalancingv2ListenerArn) > 0 {
		input.ListenerArn = aws.String(_elasticloadbalancingv2ListenerArn)
	}

	if resp, err := client.DeleteListener(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified Application Load Balancer, Network Load Balancer, or
// Gateway Load Balancer. Deleting a load balancer also deletes its listeners.
//
// You can't delete a load balancer if deletion protection is enabled. If the load
// balancer does not exist or has already been deleted, the call succeeds.
//
// Deleting a load balancer does not affect its registered targets. For example,
// your EC2 instances continue to run and are still registered to their target
// groups. If you no longer need these EC2 instances, you can stop or terminate
// them.
func elasticloadbalancingv2_DeleteLoadBalancer(cfg aws.Config, client *elasticloadbalancingv2.Client) {
	input := &elasticloadbalancingv2.DeleteLoadBalancerInput{
		// LoadBalancerArn: *string, // Required
	}

	if len(_elasticloadbalancingv2LoadBalancerArn) > 0 {
		input.LoadBalancerArn = aws.String(_elasticloadbalancingv2LoadBalancerArn)
	}

	if resp, err := client.DeleteLoadBalancer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified rule.
// You can't delete the default rule.
func elasticloadbalancingv2_DeleteRule(cfg aws.Config, client *elasticloadbalancingv2.Client) {
	input := &elasticloadbalancingv2.DeleteRuleInput{
		// RuleArn: *string, // Required
	}

	if len(_elasticloadbalancingv2RuleArn) > 0 {
		input.RuleArn = aws.String(_elasticloadbalancingv2RuleArn)
	}

	if resp, err := client.DeleteRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a shared trust store association.
func elasticloadbalancingv2_DeleteSharedTrustStoreAssociation(cfg aws.Config, client *elasticloadbalancingv2.Client) {
	input := &elasticloadbalancingv2.DeleteSharedTrustStoreAssociationInput{
		// ResourceArn: *string, // Required
		// TrustStoreArn: *string, // Required
	}

	if len(_elasticloadbalancingv2ResourceArn) > 0 {
		input.ResourceArn = aws.String(_elasticloadbalancingv2ResourceArn)
	}
	if len(_elasticloadbalancingv2TrustStoreArn) > 0 {
		input.TrustStoreArn = aws.String(_elasticloadbalancingv2TrustStoreArn)
	}

	if resp, err := client.DeleteSharedTrustStoreAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified target group.
// You can delete a target group if it is not referenced by any actions. Deleting
// a target group also deletes any associated health checks. Deleting a target
// group does not affect its registered targets. For example, any EC2 instances
// continue to run until you stop or terminate them.
func elasticloadbalancingv2_DeleteTargetGroup(cfg aws.Config, client *elasticloadbalancingv2.Client) {
	input := &elasticloadbalancingv2.DeleteTargetGroupInput{
		// TargetGroupArn: *string, // Required
	}

	if len(_elasticloadbalancingv2TargetGroupArn) > 0 {
		input.TargetGroupArn = aws.String(_elasticloadbalancingv2TargetGroupArn)
	}

	if resp, err := client.DeleteTargetGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a trust store.
func elasticloadbalancingv2_DeleteTrustStore(cfg aws.Config, client *elasticloadbalancingv2.Client) {
	input := &elasticloadbalancingv2.DeleteTrustStoreInput{
		// TrustStoreArn: *string, // Required
	}

	if len(_elasticloadbalancingv2TrustStoreArn) > 0 {
		input.TrustStoreArn = aws.String(_elasticloadbalancingv2TrustStoreArn)
	}

	if resp, err := client.DeleteTrustStore(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deregisters the specified targets from the specified target group. After the
// targets are deregistered, they no longer receive traffic from the load balancer.
//
// The load balancer stops sending requests to targets that are deregistering, but
// uses connection draining to ensure that in-flight traffic completes on the
// existing connections. This deregistration delay is configured by default but can
// be updated for each target group.
//
// For more information, see the following:
//
// [Deregistration delay]
// - in the Application Load Balancers User Guide
//
// [Deregistration delay]
// - in the Network Load Balancers User Guide
//
// [Deregistration delay]
// - in the Gateway Load Balancers User Guide
//
// Note: If the specified target does not exist, the action returns successfully.
//
// [Deregistration delay]: https://docs.aws.amazon.com/elasticloadbalancing/latest/gateway/edit-target-group-attributes.html#deregistration-delay
func elasticloadbalancingv2_DeregisterTargets(cfg aws.Config, client *elasticloadbalancingv2.Client) {
	input := &elasticloadbalancingv2.DeregisterTargetsInput{
		// TargetGroupArn: *string, // Required
		// Targets: []types.TargetDescription, // Required
	}

	if len(_elasticloadbalancingv2TargetGroupArn) > 0 {
		input.TargetGroupArn = aws.String(_elasticloadbalancingv2TargetGroupArn)
	}
	if len(_elasticloadbalancingv2Targets) > 0 {
		if err := assignInputField(input, "Targets", _elasticloadbalancingv2Targets); err != nil {
			log.Errorf("invalid --targets: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeregisterTargets(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the current Elastic Load Balancing resource limits for your Amazon
// Web Services account.
//
// For more information, see the following:
//
// [Quotas for your Application Load Balancers]
//
// [Quotas for your Network Load Balancers]
//
// [Quotas for your Gateway Load Balancers]
//
// [Quotas for your Gateway Load Balancers]: https://docs.aws.amazon.com/elasticloadbalancing/latest/gateway/quotas-limits.html
// [Quotas for your Application Load Balancers]: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-limits.html
// [Quotas for your Network Load Balancers]: https://docs.aws.amazon.com/elasticloadbalancing/latest/network/load-balancer-limits.html
func elasticloadbalancingv2_DescribeAccountLimits(cfg aws.Config, client *elasticloadbalancingv2.Client) {
	input := &elasticloadbalancingv2.DescribeAccountLimitsInput{}

	if len(_elasticloadbalancingv2Marker) > 0 {
		input.Marker = aws.String(_elasticloadbalancingv2Marker)
	}
	if len(_elasticloadbalancingv2PageSize) > 0 {
		if err := assignInputField(input, "PageSize", _elasticloadbalancingv2PageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeAccountLimits(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*elasticloadbalancingv2.DescribeAccountLimitsOutput
	p := elasticloadbalancingv2.NewDescribeAccountLimitsPaginator(client, input)
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

// Describes the capacity reservation status for the specified load balancer.
func elasticloadbalancingv2_DescribeCapacityReservation(cfg aws.Config, client *elasticloadbalancingv2.Client) {
	input := &elasticloadbalancingv2.DescribeCapacityReservationInput{
		// LoadBalancerArn: *string, // Required
	}

	if len(_elasticloadbalancingv2LoadBalancerArn) > 0 {
		input.LoadBalancerArn = aws.String(_elasticloadbalancingv2LoadBalancerArn)
	}

	if resp, err := client.DescribeCapacityReservation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the attributes for the specified listener.
func elasticloadbalancingv2_DescribeListenerAttributes(cfg aws.Config, client *elasticloadbalancingv2.Client) {
	input := &elasticloadbalancingv2.DescribeListenerAttributesInput{
		// ListenerArn: *string, // Required
	}

	if len(_elasticloadbalancingv2ListenerArn) > 0 {
		input.ListenerArn = aws.String(_elasticloadbalancingv2ListenerArn)
	}

	if resp, err := client.DescribeListenerAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the default certificate and the certificate list for the specified
// HTTPS or TLS listener.
//
// If the default certificate is also in the certificate list, it appears twice in
// the results (once with IsDefault set to true and once with IsDefault set to
// false).
//
// For more information, see [SSL certificates] in the Application Load Balancers Guide or [Server certificates] in the
// Network Load Balancers Guide.
//
// [Server certificates]: https://docs.aws.amazon.com/elasticloadbalancing/latest/network/tls-listener-certificates.html
// [SSL certificates]: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/https-listener-certificates.html
func elasticloadbalancingv2_DescribeListenerCertificates(cfg aws.Config, client *elasticloadbalancingv2.Client) {
	input := &elasticloadbalancingv2.DescribeListenerCertificatesInput{
		// ListenerArn: *string, // Required
	}

	if len(_elasticloadbalancingv2ListenerArn) > 0 {
		input.ListenerArn = aws.String(_elasticloadbalancingv2ListenerArn)
	}
	if len(_elasticloadbalancingv2Marker) > 0 {
		input.Marker = aws.String(_elasticloadbalancingv2Marker)
	}
	if len(_elasticloadbalancingv2PageSize) > 0 {
		if err := assignInputField(input, "PageSize", _elasticloadbalancingv2PageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeListenerCertificates(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*elasticloadbalancingv2.DescribeListenerCertificatesOutput
	p := elasticloadbalancingv2.NewDescribeListenerCertificatesPaginator(client, input)
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

// Describes the specified listeners or the listeners for the specified
// Application Load Balancer, Network Load Balancer, or Gateway Load Balancer. You
// must specify either a load balancer or one or more listeners.
func elasticloadbalancingv2_DescribeListeners(cfg aws.Config, client *elasticloadbalancingv2.Client) {
	input := &elasticloadbalancingv2.DescribeListenersInput{}

	if len(_elasticloadbalancingv2ListenerArns) > 0 {
		input.ListenerArns = append([]string(nil), _elasticloadbalancingv2ListenerArns...)
	}
	if len(_elasticloadbalancingv2LoadBalancerArn) > 0 {
		input.LoadBalancerArn = aws.String(_elasticloadbalancingv2LoadBalancerArn)
	}
	if len(_elasticloadbalancingv2Marker) > 0 {
		input.Marker = aws.String(_elasticloadbalancingv2Marker)
	}
	if len(_elasticloadbalancingv2PageSize) > 0 {
		if err := assignInputField(input, "PageSize", _elasticloadbalancingv2PageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeListeners(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*elasticloadbalancingv2.DescribeListenersOutput
	p := elasticloadbalancingv2.NewDescribeListenersPaginator(client, input)
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

// Describes the attributes for the specified Application Load Balancer, Network
// Load Balancer, or Gateway Load Balancer.
//
// For more information, see the following:
//
// [Load balancer attributes]
// - in the Application Load Balancers Guide
//
// [Load balancer attributes]
// - in the Network Load Balancers Guide
//
// [Load balancer attributes]
// - in the Gateway Load Balancers Guide
//
// [Load balancer attributes]: https://docs.aws.amazon.com/elasticloadbalancing/latest/gateway/gateway-load-balancers.html#load-balancer-attributes
func elasticloadbalancingv2_DescribeLoadBalancerAttributes(cfg aws.Config, client *elasticloadbalancingv2.Client) {
	input := &elasticloadbalancingv2.DescribeLoadBalancerAttributesInput{
		// LoadBalancerArn: *string, // Required
	}

	if len(_elasticloadbalancingv2LoadBalancerArn) > 0 {
		input.LoadBalancerArn = aws.String(_elasticloadbalancingv2LoadBalancerArn)
	}

	if resp, err := client.DescribeLoadBalancerAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the specified load balancers or all of your load balancers.
func elasticloadbalancingv2_DescribeLoadBalancers(cfg aws.Config, client *elasticloadbalancingv2.Client) {
	input := &elasticloadbalancingv2.DescribeLoadBalancersInput{}

	if len(_elasticloadbalancingv2LoadBalancerArns) > 0 {
		input.LoadBalancerArns = append([]string(nil), _elasticloadbalancingv2LoadBalancerArns...)
	}
	if len(_elasticloadbalancingv2Marker) > 0 {
		input.Marker = aws.String(_elasticloadbalancingv2Marker)
	}
	if len(_elasticloadbalancingv2Names) > 0 {
		input.Names = append([]string(nil), _elasticloadbalancingv2Names...)
	}
	if len(_elasticloadbalancingv2PageSize) > 0 {
		if err := assignInputField(input, "PageSize", _elasticloadbalancingv2PageSize); err != nil {
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

	var results []*elasticloadbalancingv2.DescribeLoadBalancersOutput
	p := elasticloadbalancingv2.NewDescribeLoadBalancersPaginator(client, input)
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

// Describes the specified rules or the rules for the specified listener. You must
// specify either a listener or rules.
func elasticloadbalancingv2_DescribeRules(cfg aws.Config, client *elasticloadbalancingv2.Client) {
	input := &elasticloadbalancingv2.DescribeRulesInput{}

	if len(_elasticloadbalancingv2ListenerArn) > 0 {
		input.ListenerArn = aws.String(_elasticloadbalancingv2ListenerArn)
	}
	if len(_elasticloadbalancingv2Marker) > 0 {
		input.Marker = aws.String(_elasticloadbalancingv2Marker)
	}
	if len(_elasticloadbalancingv2PageSize) > 0 {
		if err := assignInputField(input, "PageSize", _elasticloadbalancingv2PageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingv2RuleArns) > 0 {
		input.RuleArns = append([]string(nil), _elasticloadbalancingv2RuleArns...)
	}

	if disablePaginator() {
		if resp, err := client.DescribeRules(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*elasticloadbalancingv2.DescribeRulesOutput
	p := elasticloadbalancingv2.NewDescribeRulesPaginator(client, input)
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

// Describes the specified policies or all policies used for SSL negotiation.
// For more information, see [Security policies] in the Application Load Balancers Guide and [Security policies] in the
// Network Load Balancers Guide.
//
// [Security policies]: https://docs.aws.amazon.com/elasticloadbalancing/latest/network/describe-ssl-policies.html
func elasticloadbalancingv2_DescribeSSLPolicies(cfg aws.Config, client *elasticloadbalancingv2.Client) {
	input := &elasticloadbalancingv2.DescribeSSLPoliciesInput{}

	if len(_elasticloadbalancingv2LoadBalancerType) > 0 {
		if err := assignInputField(input, "LoadBalancerType", _elasticloadbalancingv2LoadBalancerType); err != nil {
			log.Errorf("invalid --load-balancer-type: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingv2Marker) > 0 {
		input.Marker = aws.String(_elasticloadbalancingv2Marker)
	}
	if len(_elasticloadbalancingv2Names) > 0 {
		input.Names = append([]string(nil), _elasticloadbalancingv2Names...)
	}
	if len(_elasticloadbalancingv2PageSize) > 0 {
		if err := assignInputField(input, "PageSize", _elasticloadbalancingv2PageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeSSLPolicies(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the tags for the specified Elastic Load Balancing resources. You can
// describe the tags for one or more Application Load Balancers, Network Load
// Balancers, Gateway Load Balancers, target groups, listeners, or rules.
func elasticloadbalancingv2_DescribeTags(cfg aws.Config, client *elasticloadbalancingv2.Client) {
	input := &elasticloadbalancingv2.DescribeTagsInput{
		// ResourceArns: []string, // Required
	}

	if len(_elasticloadbalancingv2ResourceArns) > 0 {
		input.ResourceArns = append([]string(nil), _elasticloadbalancingv2ResourceArns...)
	}

	if resp, err := client.DescribeTags(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the attributes for the specified target group.
// For more information, see the following:
//
// [Target group attributes]
// - in the Application Load Balancers Guide
//
// [Target group attributes]
// - in the Network Load Balancers Guide
//
// [Target group attributes]
// - in the Gateway Load Balancers Guide
//
// [Target group attributes]: https://docs.aws.amazon.com/elasticloadbalancing/latest/gateway/target-groups.html#target-group-attributes
func elasticloadbalancingv2_DescribeTargetGroupAttributes(cfg aws.Config, client *elasticloadbalancingv2.Client) {
	input := &elasticloadbalancingv2.DescribeTargetGroupAttributesInput{
		// TargetGroupArn: *string, // Required
	}

	if len(_elasticloadbalancingv2TargetGroupArn) > 0 {
		input.TargetGroupArn = aws.String(_elasticloadbalancingv2TargetGroupArn)
	}

	if resp, err := client.DescribeTargetGroupAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the specified target groups or all of your target groups. By default,
// all target groups are described. Alternatively, you can specify one of the
// following to filter the results: the ARN of the load balancer, the names of one
// or more target groups, or the ARNs of one or more target groups.
func elasticloadbalancingv2_DescribeTargetGroups(cfg aws.Config, client *elasticloadbalancingv2.Client) {
	input := &elasticloadbalancingv2.DescribeTargetGroupsInput{}

	if len(_elasticloadbalancingv2LoadBalancerArn) > 0 {
		input.LoadBalancerArn = aws.String(_elasticloadbalancingv2LoadBalancerArn)
	}
	if len(_elasticloadbalancingv2Marker) > 0 {
		input.Marker = aws.String(_elasticloadbalancingv2Marker)
	}
	if len(_elasticloadbalancingv2Names) > 0 {
		input.Names = append([]string(nil), _elasticloadbalancingv2Names...)
	}
	if len(_elasticloadbalancingv2PageSize) > 0 {
		if err := assignInputField(input, "PageSize", _elasticloadbalancingv2PageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingv2TargetGroupArns) > 0 {
		input.TargetGroupArns = append([]string(nil), _elasticloadbalancingv2TargetGroupArns...)
	}

	if disablePaginator() {
		if resp, err := client.DescribeTargetGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*elasticloadbalancingv2.DescribeTargetGroupsOutput
	p := elasticloadbalancingv2.NewDescribeTargetGroupsPaginator(client, input)
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

// Describes the health of the specified targets or all of your targets.
func elasticloadbalancingv2_DescribeTargetHealth(cfg aws.Config, client *elasticloadbalancingv2.Client) {
	input := &elasticloadbalancingv2.DescribeTargetHealthInput{
		// TargetGroupArn: *string, // Required
	}

	if len(_elasticloadbalancingv2TargetGroupArn) > 0 {
		input.TargetGroupArn = aws.String(_elasticloadbalancingv2TargetGroupArn)
	}
	if len(_elasticloadbalancingv2Include) > 0 {
		if err := assignInputField(input, "Include", _elasticloadbalancingv2Include); err != nil {
			log.Errorf("invalid --include: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingv2Targets) > 0 {
		if err := assignInputField(input, "Targets", _elasticloadbalancingv2Targets); err != nil {
			log.Errorf("invalid --targets: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeTargetHealth(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes all resources associated with the specified trust store.
func elasticloadbalancingv2_DescribeTrustStoreAssociations(cfg aws.Config, client *elasticloadbalancingv2.Client) {
	input := &elasticloadbalancingv2.DescribeTrustStoreAssociationsInput{
		// TrustStoreArn: *string, // Required
	}

	if len(_elasticloadbalancingv2TrustStoreArn) > 0 {
		input.TrustStoreArn = aws.String(_elasticloadbalancingv2TrustStoreArn)
	}
	if len(_elasticloadbalancingv2Marker) > 0 {
		input.Marker = aws.String(_elasticloadbalancingv2Marker)
	}
	if len(_elasticloadbalancingv2PageSize) > 0 {
		if err := assignInputField(input, "PageSize", _elasticloadbalancingv2PageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeTrustStoreAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*elasticloadbalancingv2.DescribeTrustStoreAssociationsOutput
	p := elasticloadbalancingv2.NewDescribeTrustStoreAssociationsPaginator(client, input)
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

// Describes the revocation files in use by the specified trust store or
// revocation files.
func elasticloadbalancingv2_DescribeTrustStoreRevocations(cfg aws.Config, client *elasticloadbalancingv2.Client) {
	input := &elasticloadbalancingv2.DescribeTrustStoreRevocationsInput{
		// TrustStoreArn: *string, // Required
	}

	if len(_elasticloadbalancingv2TrustStoreArn) > 0 {
		input.TrustStoreArn = aws.String(_elasticloadbalancingv2TrustStoreArn)
	}
	if len(_elasticloadbalancingv2Marker) > 0 {
		input.Marker = aws.String(_elasticloadbalancingv2Marker)
	}
	if len(_elasticloadbalancingv2PageSize) > 0 {
		if err := assignInputField(input, "PageSize", _elasticloadbalancingv2PageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingv2RevocationIds) > 0 {
		if err := assignInputField(input, "RevocationIds", _elasticloadbalancingv2RevocationIds); err != nil {
			log.Errorf("invalid --revocation-ids: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeTrustStoreRevocations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*elasticloadbalancingv2.DescribeTrustStoreRevocationsOutput
	p := elasticloadbalancingv2.NewDescribeTrustStoreRevocationsPaginator(client, input)
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

// Describes all trust stores for the specified account.
func elasticloadbalancingv2_DescribeTrustStores(cfg aws.Config, client *elasticloadbalancingv2.Client) {
	input := &elasticloadbalancingv2.DescribeTrustStoresInput{}

	if len(_elasticloadbalancingv2Marker) > 0 {
		input.Marker = aws.String(_elasticloadbalancingv2Marker)
	}
	if len(_elasticloadbalancingv2Names) > 0 {
		input.Names = append([]string(nil), _elasticloadbalancingv2Names...)
	}
	if len(_elasticloadbalancingv2PageSize) > 0 {
		if err := assignInputField(input, "PageSize", _elasticloadbalancingv2PageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingv2TrustStoreArns) > 0 {
		input.TrustStoreArns = append([]string(nil), _elasticloadbalancingv2TrustStoreArns...)
	}

	if disablePaginator() {
		if resp, err := client.DescribeTrustStores(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*elasticloadbalancingv2.DescribeTrustStoresOutput
	p := elasticloadbalancingv2.NewDescribeTrustStoresPaginator(client, input)
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

// Retrieves the resource policy for a specified resource.
func elasticloadbalancingv2_GetResourcePolicy(cfg aws.Config, client *elasticloadbalancingv2.Client) {
	input := &elasticloadbalancingv2.GetResourcePolicyInput{
		// ResourceArn: *string, // Required
	}

	if len(_elasticloadbalancingv2ResourceArn) > 0 {
		input.ResourceArn = aws.String(_elasticloadbalancingv2ResourceArn)
	}

	if resp, err := client.GetResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the ca certificate bundle.
// This action returns a pre-signed S3 URI which is active for ten minutes.
func elasticloadbalancingv2_GetTrustStoreCaCertificatesBundle(cfg aws.Config, client *elasticloadbalancingv2.Client) {
	input := &elasticloadbalancingv2.GetTrustStoreCaCertificatesBundleInput{
		// TrustStoreArn: *string, // Required
	}

	if len(_elasticloadbalancingv2TrustStoreArn) > 0 {
		input.TrustStoreArn = aws.String(_elasticloadbalancingv2TrustStoreArn)
	}

	if resp, err := client.GetTrustStoreCaCertificatesBundle(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the specified revocation file.
// This action returns a pre-signed S3 URI which is active for ten minutes.
func elasticloadbalancingv2_GetTrustStoreRevocationContent(cfg aws.Config, client *elasticloadbalancingv2.Client) {
	input := &elasticloadbalancingv2.GetTrustStoreRevocationContentInput{
		// RevocationId: *int64, // Required
		// TrustStoreArn: *string, // Required
	}

	if len(_elasticloadbalancingv2RevocationId) > 0 {
		if err := assignInputField(input, "RevocationId", _elasticloadbalancingv2RevocationId); err != nil {
			log.Errorf("invalid --revocation-id: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingv2TrustStoreArn) > 0 {
		input.TrustStoreArn = aws.String(_elasticloadbalancingv2TrustStoreArn)
	}

	if resp, err := client.GetTrustStoreRevocationContent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the capacity reservation of the specified load balancer.
// When modifying capacity reservation, you must include at least one
// MinimumLoadBalancerCapacity or ResetCapacityReservation .
func elasticloadbalancingv2_ModifyCapacityReservation(cfg aws.Config, client *elasticloadbalancingv2.Client) {
	input := &elasticloadbalancingv2.ModifyCapacityReservationInput{
		// LoadBalancerArn: *string, // Required
	}

	if len(_elasticloadbalancingv2LoadBalancerArn) > 0 {
		input.LoadBalancerArn = aws.String(_elasticloadbalancingv2LoadBalancerArn)
	}
	if len(_elasticloadbalancingv2MinimumLoadBalancerCapacity) > 0 {
		if err := assignInputField(input, "MinimumLoadBalancerCapacity", _elasticloadbalancingv2MinimumLoadBalancerCapacity); err != nil {
			log.Errorf("invalid --minimum-load-balancer-capacity: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingv2ResetCapacityReservation) > 0 {
		if err := assignInputField(input, "ResetCapacityReservation", _elasticloadbalancingv2ResetCapacityReservation); err != nil {
			log.Errorf("invalid --reset-capacity-reservation: %s", err.Error())
			return
		}
	}

	if resp, err := client.ModifyCapacityReservation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// [Application Load Balancers] Modify the IP pool associated to a load balancer.
func elasticloadbalancingv2_ModifyIpPools(cfg aws.Config, client *elasticloadbalancingv2.Client) {
	input := &elasticloadbalancingv2.ModifyIpPoolsInput{
		// LoadBalancerArn: *string, // Required
	}

	if len(_elasticloadbalancingv2LoadBalancerArn) > 0 {
		input.LoadBalancerArn = aws.String(_elasticloadbalancingv2LoadBalancerArn)
	}
	if len(_elasticloadbalancingv2IpamPools) > 0 {
		if err := assignInputField(input, "IpamPools", _elasticloadbalancingv2IpamPools); err != nil {
			log.Errorf("invalid --ipam-pools: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingv2RemoveIpamPools) > 0 {
		if err := assignInputField(input, "RemoveIpamPools", _elasticloadbalancingv2RemoveIpamPools); err != nil {
			log.Errorf("invalid --remove-ipam-pools: %s", err.Error())
			return
		}
	}

	if resp, err := client.ModifyIpPools(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Replaces the specified properties of the specified listener. Any properties
// that you do not specify remain unchanged.
//
// Changing the protocol from HTTPS to HTTP, or from TLS to TCP, removes the
// security policy and default certificate properties. If you change the protocol
// from HTTP to HTTPS, or from TCP to TLS, you must add the security policy and
// default certificate properties.
//
// To add an item to a list, remove an item from a list, or update an item in a
// list, you must provide the entire list. For example, to add an action, specify a
// list with the current actions plus the new action.
func elasticloadbalancingv2_ModifyListener(cfg aws.Config, client *elasticloadbalancingv2.Client) {
	input := &elasticloadbalancingv2.ModifyListenerInput{
		// ListenerArn: *string, // Required
	}

	if len(_elasticloadbalancingv2ListenerArn) > 0 {
		input.ListenerArn = aws.String(_elasticloadbalancingv2ListenerArn)
	}
	if len(_elasticloadbalancingv2AlpnPolicy) > 0 {
		input.AlpnPolicy = append([]string(nil), _elasticloadbalancingv2AlpnPolicy...)
	}
	if len(_elasticloadbalancingv2Certificates) > 0 {
		if err := assignInputField(input, "Certificates", _elasticloadbalancingv2Certificates); err != nil {
			log.Errorf("invalid --certificates: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingv2DefaultActions) > 0 {
		if err := assignInputField(input, "DefaultActions", _elasticloadbalancingv2DefaultActions); err != nil {
			log.Errorf("invalid --default-actions: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingv2MutualAuthentication) > 0 {
		if err := assignInputField(input, "MutualAuthentication", _elasticloadbalancingv2MutualAuthentication); err != nil {
			log.Errorf("invalid --mutual-authentication: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingv2Port) > 0 {
		if err := assignInputField(input, "Port", _elasticloadbalancingv2Port); err != nil {
			log.Errorf("invalid --port: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingv2Protocol) > 0 {
		if err := assignInputField(input, "Protocol", _elasticloadbalancingv2Protocol); err != nil {
			log.Errorf("invalid --protocol: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingv2SslPolicy) > 0 {
		input.SslPolicy = aws.String(_elasticloadbalancingv2SslPolicy)
	}

	if resp, err := client.ModifyListener(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the specified attributes of the specified listener.
func elasticloadbalancingv2_ModifyListenerAttributes(cfg aws.Config, client *elasticloadbalancingv2.Client) {
	input := &elasticloadbalancingv2.ModifyListenerAttributesInput{
		// Attributes: []types.ListenerAttribute, // Required
		// ListenerArn: *string, // Required
	}

	if len(_elasticloadbalancingv2Attributes) > 0 {
		if err := assignInputField(input, "Attributes", _elasticloadbalancingv2Attributes); err != nil {
			log.Errorf("invalid --attributes: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingv2ListenerArn) > 0 {
		input.ListenerArn = aws.String(_elasticloadbalancingv2ListenerArn)
	}

	if resp, err := client.ModifyListenerAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the specified attributes of the specified Application Load Balancer,
// Network Load Balancer, or Gateway Load Balancer.
//
// If any of the specified attributes can't be modified as requested, the call
// fails. Any existing attributes that you do not modify retain their current
// values.
func elasticloadbalancingv2_ModifyLoadBalancerAttributes(cfg aws.Config, client *elasticloadbalancingv2.Client) {
	input := &elasticloadbalancingv2.ModifyLoadBalancerAttributesInput{
		// Attributes: []types.LoadBalancerAttribute, // Required
		// LoadBalancerArn: *string, // Required
	}

	if len(_elasticloadbalancingv2Attributes) > 0 {
		if err := assignInputField(input, "Attributes", _elasticloadbalancingv2Attributes); err != nil {
			log.Errorf("invalid --attributes: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingv2LoadBalancerArn) > 0 {
		input.LoadBalancerArn = aws.String(_elasticloadbalancingv2LoadBalancerArn)
	}

	if resp, err := client.ModifyLoadBalancerAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Replaces the specified properties of the specified rule. Any properties that
// you do not specify are unchanged.
//
// To add an item to a list, remove an item from a list, or update an item in a
// list, you must provide the entire list. For example, to add an action, specify a
// list with the current actions plus the new action.
func elasticloadbalancingv2_ModifyRule(cfg aws.Config, client *elasticloadbalancingv2.Client) {
	input := &elasticloadbalancingv2.ModifyRuleInput{
		// RuleArn: *string, // Required
	}

	if len(_elasticloadbalancingv2RuleArn) > 0 {
		input.RuleArn = aws.String(_elasticloadbalancingv2RuleArn)
	}
	if len(_elasticloadbalancingv2Actions) > 0 {
		if err := assignInputField(input, "Actions", _elasticloadbalancingv2Actions); err != nil {
			log.Errorf("invalid --actions: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingv2Conditions) > 0 {
		if err := assignInputField(input, "Conditions", _elasticloadbalancingv2Conditions); err != nil {
			log.Errorf("invalid --conditions: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingv2ResetTransforms) > 0 {
		if err := assignInputField(input, "ResetTransforms", _elasticloadbalancingv2ResetTransforms); err != nil {
			log.Errorf("invalid --reset-transforms: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingv2Transforms) > 0 {
		if err := assignInputField(input, "Transforms", _elasticloadbalancingv2Transforms); err != nil {
			log.Errorf("invalid --transforms: %s", err.Error())
			return
		}
	}

	if resp, err := client.ModifyRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the health checks used when evaluating the health state of the targets
// in the specified target group.
func elasticloadbalancingv2_ModifyTargetGroup(cfg aws.Config, client *elasticloadbalancingv2.Client) {
	input := &elasticloadbalancingv2.ModifyTargetGroupInput{
		// TargetGroupArn: *string, // Required
	}

	if len(_elasticloadbalancingv2TargetGroupArn) > 0 {
		input.TargetGroupArn = aws.String(_elasticloadbalancingv2TargetGroupArn)
	}
	if len(_elasticloadbalancingv2HealthCheckEnabled) > 0 {
		if err := assignInputField(input, "HealthCheckEnabled", _elasticloadbalancingv2HealthCheckEnabled); err != nil {
			log.Errorf("invalid --health-check-enabled: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingv2HealthCheckIntervalSeconds) > 0 {
		if err := assignInputField(input, "HealthCheckIntervalSeconds", _elasticloadbalancingv2HealthCheckIntervalSeconds); err != nil {
			log.Errorf("invalid --health-check-interval-seconds: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingv2HealthCheckPath) > 0 {
		input.HealthCheckPath = aws.String(_elasticloadbalancingv2HealthCheckPath)
	}
	if len(_elasticloadbalancingv2HealthCheckPort) > 0 {
		input.HealthCheckPort = aws.String(_elasticloadbalancingv2HealthCheckPort)
	}
	if len(_elasticloadbalancingv2HealthCheckProtocol) > 0 {
		if err := assignInputField(input, "HealthCheckProtocol", _elasticloadbalancingv2HealthCheckProtocol); err != nil {
			log.Errorf("invalid --health-check-protocol: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingv2HealthCheckTimeoutSeconds) > 0 {
		if err := assignInputField(input, "HealthCheckTimeoutSeconds", _elasticloadbalancingv2HealthCheckTimeoutSeconds); err != nil {
			log.Errorf("invalid --health-check-timeout-seconds: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingv2HealthyThresholdCount) > 0 {
		if err := assignInputField(input, "HealthyThresholdCount", _elasticloadbalancingv2HealthyThresholdCount); err != nil {
			log.Errorf("invalid --healthy-threshold-count: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingv2Matcher) > 0 {
		if err := assignInputField(input, "Matcher", _elasticloadbalancingv2Matcher); err != nil {
			log.Errorf("invalid --matcher: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingv2UnhealthyThresholdCount) > 0 {
		if err := assignInputField(input, "UnhealthyThresholdCount", _elasticloadbalancingv2UnhealthyThresholdCount); err != nil {
			log.Errorf("invalid --unhealthy-threshold-count: %s", err.Error())
			return
		}
	}

	if resp, err := client.ModifyTargetGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the specified attributes of the specified target group.
func elasticloadbalancingv2_ModifyTargetGroupAttributes(cfg aws.Config, client *elasticloadbalancingv2.Client) {
	input := &elasticloadbalancingv2.ModifyTargetGroupAttributesInput{
		// Attributes: []types.TargetGroupAttribute, // Required
		// TargetGroupArn: *string, // Required
	}

	if len(_elasticloadbalancingv2Attributes) > 0 {
		if err := assignInputField(input, "Attributes", _elasticloadbalancingv2Attributes); err != nil {
			log.Errorf("invalid --attributes: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingv2TargetGroupArn) > 0 {
		input.TargetGroupArn = aws.String(_elasticloadbalancingv2TargetGroupArn)
	}

	if resp, err := client.ModifyTargetGroupAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update the ca certificate bundle for the specified trust store.
func elasticloadbalancingv2_ModifyTrustStore(cfg aws.Config, client *elasticloadbalancingv2.Client) {
	input := &elasticloadbalancingv2.ModifyTrustStoreInput{
		// CaCertificatesBundleS3Bucket: *string, // Required
		// CaCertificatesBundleS3Key: *string, // Required
		// TrustStoreArn: *string, // Required
	}

	if len(_elasticloadbalancingv2CaCertificatesBundleS3Bucket) > 0 {
		input.CaCertificatesBundleS3Bucket = aws.String(_elasticloadbalancingv2CaCertificatesBundleS3Bucket)
	}
	if len(_elasticloadbalancingv2CaCertificatesBundleS3Key) > 0 {
		input.CaCertificatesBundleS3Key = aws.String(_elasticloadbalancingv2CaCertificatesBundleS3Key)
	}
	if len(_elasticloadbalancingv2TrustStoreArn) > 0 {
		input.TrustStoreArn = aws.String(_elasticloadbalancingv2TrustStoreArn)
	}
	if len(_elasticloadbalancingv2CaCertificatesBundleS3ObjectVersion) > 0 {
		input.CaCertificatesBundleS3ObjectVersion = aws.String(_elasticloadbalancingv2CaCertificatesBundleS3ObjectVersion)
	}

	if resp, err := client.ModifyTrustStore(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Registers the specified targets with the specified target group.
// If the target is an EC2 instance, it must be in the running state when you
// register it.
//
// By default, the load balancer routes requests to registered targets using the
// protocol and port for the target group. Alternatively, you can override the port
// for a target when you register it. You can register each EC2 instance or IP
// address with the same target group multiple times using different ports.
//
// For more information, see the following:
//
// [Register targets for your Application Load Balancer]
//
// [Register targets for your Network Load Balancer]
//
// [Register targets for your Gateway Load Balancer]
//
// [Register targets for your Network Load Balancer]: https://docs.aws.amazon.com/elasticloadbalancing/latest/network/target-group-register-targets.html
// [Register targets for your Gateway Load Balancer]: https://docs.aws.amazon.com/elasticloadbalancing/latest/gateway/target-group-register-targets.html
// [Register targets for your Application Load Balancer]: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/target-group-register-targets.html
func elasticloadbalancingv2_RegisterTargets(cfg aws.Config, client *elasticloadbalancingv2.Client) {
	input := &elasticloadbalancingv2.RegisterTargetsInput{
		// TargetGroupArn: *string, // Required
		// Targets: []types.TargetDescription, // Required
	}

	if len(_elasticloadbalancingv2TargetGroupArn) > 0 {
		input.TargetGroupArn = aws.String(_elasticloadbalancingv2TargetGroupArn)
	}
	if len(_elasticloadbalancingv2Targets) > 0 {
		if err := assignInputField(input, "Targets", _elasticloadbalancingv2Targets); err != nil {
			log.Errorf("invalid --targets: %s", err.Error())
			return
		}
	}

	if resp, err := client.RegisterTargets(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the specified certificate from the certificate list for the specified
// HTTPS or TLS listener.
func elasticloadbalancingv2_RemoveListenerCertificates(cfg aws.Config, client *elasticloadbalancingv2.Client) {
	input := &elasticloadbalancingv2.RemoveListenerCertificatesInput{
		// Certificates: []types.Certificate, // Required
		// ListenerArn: *string, // Required
	}

	if len(_elasticloadbalancingv2Certificates) > 0 {
		if err := assignInputField(input, "Certificates", _elasticloadbalancingv2Certificates); err != nil {
			log.Errorf("invalid --certificates: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingv2ListenerArn) > 0 {
		input.ListenerArn = aws.String(_elasticloadbalancingv2ListenerArn)
	}

	if resp, err := client.RemoveListenerCertificates(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the specified tags from the specified Elastic Load Balancing resources.
// You can remove the tags for one or more Application Load Balancers, Network Load
// Balancers, Gateway Load Balancers, target groups, listeners, or rules.
func elasticloadbalancingv2_RemoveTags(cfg aws.Config, client *elasticloadbalancingv2.Client) {
	input := &elasticloadbalancingv2.RemoveTagsInput{
		// ResourceArns: []string, // Required
		// TagKeys: []string, // Required
	}

	if len(_elasticloadbalancingv2ResourceArns) > 0 {
		input.ResourceArns = append([]string(nil), _elasticloadbalancingv2ResourceArns...)
	}
	if len(_elasticloadbalancingv2TagKeys) > 0 {
		input.TagKeys = append([]string(nil), _elasticloadbalancingv2TagKeys...)
	}

	if resp, err := client.RemoveTags(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the specified revocation file from the specified trust store.
func elasticloadbalancingv2_RemoveTrustStoreRevocations(cfg aws.Config, client *elasticloadbalancingv2.Client) {
	input := &elasticloadbalancingv2.RemoveTrustStoreRevocationsInput{
		// RevocationIds: []int64, // Required
		// TrustStoreArn: *string, // Required
	}

	if len(_elasticloadbalancingv2RevocationIds) > 0 {
		if err := assignInputField(input, "RevocationIds", _elasticloadbalancingv2RevocationIds); err != nil {
			log.Errorf("invalid --revocation-ids: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingv2TrustStoreArn) > 0 {
		input.TrustStoreArn = aws.String(_elasticloadbalancingv2TrustStoreArn)
	}

	if resp, err := client.RemoveTrustStoreRevocations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets the type of IP addresses used by the subnets of the specified load
// balancer.
func elasticloadbalancingv2_SetIpAddressType(cfg aws.Config, client *elasticloadbalancingv2.Client) {
	input := &elasticloadbalancingv2.SetIpAddressTypeInput{
		// IpAddressType: types.IpAddressType, // Required
		// LoadBalancerArn: *string, // Required
	}

	if len(_elasticloadbalancingv2IpAddressType) > 0 {
		if err := assignInputField(input, "IpAddressType", _elasticloadbalancingv2IpAddressType); err != nil {
			log.Errorf("invalid --ip-address-type: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingv2LoadBalancerArn) > 0 {
		input.LoadBalancerArn = aws.String(_elasticloadbalancingv2LoadBalancerArn)
	}

	if resp, err := client.SetIpAddressType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets the priorities of the specified rules.
// You can reorder the rules as long as there are no priority conflicts in the new
// order. Any existing rules that you do not specify retain their current priority.
func elasticloadbalancingv2_SetRulePriorities(cfg aws.Config, client *elasticloadbalancingv2.Client) {
	input := &elasticloadbalancingv2.SetRulePrioritiesInput{
		// RulePriorities: []types.RulePriorityPair, // Required
	}

	if len(_elasticloadbalancingv2RulePriorities) > 0 {
		if err := assignInputField(input, "RulePriorities", _elasticloadbalancingv2RulePriorities); err != nil {
			log.Errorf("invalid --rule-priorities: %s", err.Error())
			return
		}
	}

	if resp, err := client.SetRulePriorities(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates the specified security groups with the specified Application Load
// Balancer or Network Load Balancer. The specified security groups override the
// previously associated security groups.
//
// You can't perform this operation on a Network Load Balancer unless you
// specified a security group for the load balancer when you created it.
//
// You can't associate a security group with a Gateway Load Balancer.
func elasticloadbalancingv2_SetSecurityGroups(cfg aws.Config, client *elasticloadbalancingv2.Client) {
	input := &elasticloadbalancingv2.SetSecurityGroupsInput{
		// LoadBalancerArn: *string, // Required
		// SecurityGroups: []string, // Required
	}

	if len(_elasticloadbalancingv2LoadBalancerArn) > 0 {
		input.LoadBalancerArn = aws.String(_elasticloadbalancingv2LoadBalancerArn)
	}
	if len(_elasticloadbalancingv2SecurityGroups) > 0 {
		input.SecurityGroups = append([]string(nil), _elasticloadbalancingv2SecurityGroups...)
	}
	if len(_elasticloadbalancingv2EnforceSecurityGroupInboundRulesOnPrivateLinkTraffic) > 0 {
		if err := assignInputField(input, "EnforceSecurityGroupInboundRulesOnPrivateLinkTraffic", _elasticloadbalancingv2EnforceSecurityGroupInboundRulesOnPrivateLinkTraffic); err != nil {
			log.Errorf("invalid --enforce-security-group-inbound-rules-on-private-link-traffic: %s", err.Error())
			return
		}
	}

	if resp, err := client.SetSecurityGroups(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables the Availability Zones for the specified public subnets for the
// specified Application Load Balancer, Network Load Balancer or Gateway Load
// Balancer. The specified subnets replace the previously enabled subnets.
func elasticloadbalancingv2_SetSubnets(cfg aws.Config, client *elasticloadbalancingv2.Client) {
	input := &elasticloadbalancingv2.SetSubnetsInput{
		// LoadBalancerArn: *string, // Required
	}

	if len(_elasticloadbalancingv2LoadBalancerArn) > 0 {
		input.LoadBalancerArn = aws.String(_elasticloadbalancingv2LoadBalancerArn)
	}
	if len(_elasticloadbalancingv2EnablePrefixForIpv6SourceNat) > 0 {
		if err := assignInputField(input, "EnablePrefixForIpv6SourceNat", _elasticloadbalancingv2EnablePrefixForIpv6SourceNat); err != nil {
			log.Errorf("invalid --enable-prefix-for-ipv6-source-nat: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingv2IpAddressType) > 0 {
		if err := assignInputField(input, "IpAddressType", _elasticloadbalancingv2IpAddressType); err != nil {
			log.Errorf("invalid --ip-address-type: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingv2SubnetMappings) > 0 {
		if err := assignInputField(input, "SubnetMappings", _elasticloadbalancingv2SubnetMappings); err != nil {
			log.Errorf("invalid --subnet-mappings: %s", err.Error())
			return
		}
	}
	if len(_elasticloadbalancingv2Subnets) > 0 {
		input.Subnets = append([]string(nil), _elasticloadbalancingv2Subnets...)
	}

	if resp, err := client.SetSubnets(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_elasticloadbalancingv2Cmd)
	_elasticloadbalancingv2Cmd.Flags().SortFlags = false

	_elasticloadbalancingv2Cmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_elasticloadbalancingv2Cmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_elasticloadbalancingv2Cmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_elasticloadbalancingv2Cmd.Flags().StringVarP(&_elasticloadbalancingv2Actions, "actions", "", "", "Actions")
	_elasticloadbalancingv2Cmd.Flags().StringSliceVarP(&_elasticloadbalancingv2AlpnPolicy, "alpn-policy", "", nil, "Alpn Policy")
	_elasticloadbalancingv2Cmd.Flags().StringVarP(&_elasticloadbalancingv2Attributes, "attributes", "", "", "Attributes")
	_elasticloadbalancingv2Cmd.Flags().StringVarP(&_elasticloadbalancingv2CaCertificatesBundleS3Bucket, "ca-certificates-bundle-s3-bucket", "", "", "Ca Certificates Bundle S3 Bucket")
	_elasticloadbalancingv2Cmd.Flags().StringVarP(&_elasticloadbalancingv2CaCertificatesBundleS3Key, "ca-certificates-bundle-s3-key", "", "", "Ca Certificates Bundle S3 Key")
	_elasticloadbalancingv2Cmd.Flags().StringVarP(&_elasticloadbalancingv2CaCertificatesBundleS3ObjectVersion, "ca-certificates-bundle-s3-object-version", "", "", "Ca Certificates Bundle S3 Object Version")
	_elasticloadbalancingv2Cmd.Flags().StringVarP(&_elasticloadbalancingv2Certificates, "certificates", "", "", "Certificates")
	_elasticloadbalancingv2Cmd.Flags().StringVarP(&_elasticloadbalancingv2Conditions, "conditions", "", "", "Conditions")
	_elasticloadbalancingv2Cmd.Flags().StringVarP(&_elasticloadbalancingv2CustomerOwnedIpv4Pool, "customer-owned-ipv4-pool", "", "", "Customer Owned IPV4 Pool")
	_elasticloadbalancingv2Cmd.Flags().StringVarP(&_elasticloadbalancingv2DefaultActions, "default-actions", "", "", "Default Actions")
	_elasticloadbalancingv2Cmd.Flags().StringVarP(&_elasticloadbalancingv2EnablePrefixForIpv6SourceNat, "enable-prefix-for-ipv6-source-nat", "", "", "Enable Prefix For IPV6 Source NAT")
	_elasticloadbalancingv2Cmd.Flags().StringVarP(&_elasticloadbalancingv2EnforceSecurityGroupInboundRulesOnPrivateLinkTraffic, "enforce-security-group-inbound-rules-on-private-link-traffic", "", "", "Enforce Security Group Inbound Rules On Private Link Traffic")
	_elasticloadbalancingv2Cmd.Flags().StringVarP(&_elasticloadbalancingv2HealthCheckEnabled, "health-check-enabled", "", "", "Health Check Enabled")
	_elasticloadbalancingv2Cmd.Flags().StringVarP(&_elasticloadbalancingv2HealthCheckIntervalSeconds, "health-check-interval-seconds", "", "", "Health Check Interval Seconds")
	_elasticloadbalancingv2Cmd.Flags().StringVarP(&_elasticloadbalancingv2HealthCheckPath, "health-check-path", "", "", "Health Check Path")
	_elasticloadbalancingv2Cmd.Flags().StringVarP(&_elasticloadbalancingv2HealthCheckPort, "health-check-port", "", "", "Health Check Port")
	_elasticloadbalancingv2Cmd.Flags().StringVarP(&_elasticloadbalancingv2HealthCheckProtocol, "health-check-protocol", "", "", "Health Check Protocol")
	_elasticloadbalancingv2Cmd.Flags().StringVarP(&_elasticloadbalancingv2HealthCheckTimeoutSeconds, "health-check-timeout-seconds", "", "", "Health Check Timeout Seconds")
	_elasticloadbalancingv2Cmd.Flags().StringVarP(&_elasticloadbalancingv2HealthyThresholdCount, "healthy-threshold-count", "", "", "Healthy Threshold Count")
	_elasticloadbalancingv2Cmd.Flags().StringVarP(&_elasticloadbalancingv2Include, "include", "", "", "Include")
	_elasticloadbalancingv2Cmd.Flags().StringVarP(&_elasticloadbalancingv2IpAddressType, "ip-address-type", "", "", "IP Address Type")
	_elasticloadbalancingv2Cmd.Flags().StringVarP(&_elasticloadbalancingv2IpamPools, "ipam-pools", "", "", "Ipam Pools")
	_elasticloadbalancingv2Cmd.Flags().StringVarP(&_elasticloadbalancingv2ListenerArn, "listener-arn", "", "", "Listener ARN")
	_elasticloadbalancingv2Cmd.Flags().StringSliceVarP(&_elasticloadbalancingv2ListenerArns, "listener-arns", "", nil, "Listener Arns")
	_elasticloadbalancingv2Cmd.Flags().StringVarP(&_elasticloadbalancingv2LoadBalancerArn, "load-balancer-arn", "", "", "Load Balancer ARN")
	_elasticloadbalancingv2Cmd.Flags().StringSliceVarP(&_elasticloadbalancingv2LoadBalancerArns, "load-balancer-arns", "", nil, "Load Balancer Arns")
	_elasticloadbalancingv2Cmd.Flags().StringVarP(&_elasticloadbalancingv2LoadBalancerType, "load-balancer-type", "", "", "Load Balancer Type")
	_elasticloadbalancingv2Cmd.Flags().StringVarP(&_elasticloadbalancingv2Marker, "marker", "", "", "Marker")
	_elasticloadbalancingv2Cmd.Flags().StringVarP(&_elasticloadbalancingv2Matcher, "matcher", "", "", "Matcher")
	_elasticloadbalancingv2Cmd.Flags().StringVarP(&_elasticloadbalancingv2MinimumLoadBalancerCapacity, "minimum-load-balancer-capacity", "", "", "Minimum Load Balancer Capacity")
	_elasticloadbalancingv2Cmd.Flags().StringVarP(&_elasticloadbalancingv2MutualAuthentication, "mutual-authentication", "", "", "Mutual Authentication")
	_elasticloadbalancingv2Cmd.Flags().StringVarP(&_elasticloadbalancingv2Name, "name", "", "", "Name")
	_elasticloadbalancingv2Cmd.Flags().StringSliceVarP(&_elasticloadbalancingv2Names, "names", "", nil, "Names")
	_elasticloadbalancingv2Cmd.Flags().StringVarP(&_elasticloadbalancingv2PageSize, "page-size", "", "", "Page Size")
	_elasticloadbalancingv2Cmd.Flags().StringVarP(&_elasticloadbalancingv2Port, "port", "", "", "Port")
	_elasticloadbalancingv2Cmd.Flags().StringVarP(&_elasticloadbalancingv2Priority, "priority", "", "", "Priority")
	_elasticloadbalancingv2Cmd.Flags().StringVarP(&_elasticloadbalancingv2Protocol, "protocol", "", "", "Protocol")
	_elasticloadbalancingv2Cmd.Flags().StringVarP(&_elasticloadbalancingv2ProtocolVersion, "protocol-version", "", "", "Protocol Version")
	_elasticloadbalancingv2Cmd.Flags().StringVarP(&_elasticloadbalancingv2RemoveIpamPools, "remove-ipam-pools", "", "", "Remove Ipam Pools")
	_elasticloadbalancingv2Cmd.Flags().StringVarP(&_elasticloadbalancingv2ResetCapacityReservation, "reset-capacity-reservation", "", "", "Reset Capacity Reservation")
	_elasticloadbalancingv2Cmd.Flags().StringVarP(&_elasticloadbalancingv2ResetTransforms, "reset-transforms", "", "", "Reset Transforms")
	_elasticloadbalancingv2Cmd.Flags().StringVarP(&_elasticloadbalancingv2ResourceArn, "resource-arn", "", "", "Resource ARN")
	_elasticloadbalancingv2Cmd.Flags().StringSliceVarP(&_elasticloadbalancingv2ResourceArns, "resource-arns", "", nil, "Resource Arns")
	_elasticloadbalancingv2Cmd.Flags().StringVarP(&_elasticloadbalancingv2RevocationContents, "revocation-contents", "", "", "Revocation Contents")
	_elasticloadbalancingv2Cmd.Flags().StringVarP(&_elasticloadbalancingv2RevocationId, "revocation-id", "", "", "Revocation ID")
	_elasticloadbalancingv2Cmd.Flags().StringVarP(&_elasticloadbalancingv2RevocationIds, "revocation-ids", "", "", "Revocation Ids")
	_elasticloadbalancingv2Cmd.Flags().StringVarP(&_elasticloadbalancingv2RuleArn, "rule-arn", "", "", "Rule ARN")
	_elasticloadbalancingv2Cmd.Flags().StringSliceVarP(&_elasticloadbalancingv2RuleArns, "rule-arns", "", nil, "Rule Arns")
	_elasticloadbalancingv2Cmd.Flags().StringVarP(&_elasticloadbalancingv2RulePriorities, "rule-priorities", "", "", "Rule Priorities")
	_elasticloadbalancingv2Cmd.Flags().StringVarP(&_elasticloadbalancingv2Scheme, "scheme", "", "", "Scheme")
	_elasticloadbalancingv2Cmd.Flags().StringSliceVarP(&_elasticloadbalancingv2SecurityGroups, "security-groups", "", nil, "Security Groups")
	_elasticloadbalancingv2Cmd.Flags().StringVarP(&_elasticloadbalancingv2SslPolicy, "ssl-policy", "", "", "SSL Policy")
	_elasticloadbalancingv2Cmd.Flags().StringVarP(&_elasticloadbalancingv2SubnetMappings, "subnet-mappings", "", "", "Subnet Mappings")
	_elasticloadbalancingv2Cmd.Flags().StringSliceVarP(&_elasticloadbalancingv2Subnets, "subnets", "", nil, "Subnets")
	_elasticloadbalancingv2Cmd.Flags().StringSliceVarP(&_elasticloadbalancingv2TagKeys, "tag-keys", "", nil, "Tag Keys")
	_elasticloadbalancingv2Cmd.Flags().StringVarP(&_elasticloadbalancingv2Tags, "tags", "", "", "Tags")
	_elasticloadbalancingv2Cmd.Flags().StringVarP(&_elasticloadbalancingv2TargetControlPort, "target-control-port", "", "", "Target Control Port")
	_elasticloadbalancingv2Cmd.Flags().StringVarP(&_elasticloadbalancingv2TargetGroupArn, "target-group-arn", "", "", "Target Group ARN")
	_elasticloadbalancingv2Cmd.Flags().StringSliceVarP(&_elasticloadbalancingv2TargetGroupArns, "target-group-arns", "", nil, "Target Group Arns")
	_elasticloadbalancingv2Cmd.Flags().StringVarP(&_elasticloadbalancingv2TargetType, "target-type", "", "", "Target Type")
	_elasticloadbalancingv2Cmd.Flags().StringVarP(&_elasticloadbalancingv2Targets, "targets", "", "", "Targets")
	_elasticloadbalancingv2Cmd.Flags().StringVarP(&_elasticloadbalancingv2Transforms, "transforms", "", "", "Transforms")
	_elasticloadbalancingv2Cmd.Flags().StringVarP(&_elasticloadbalancingv2TrustStoreArn, "trust-store-arn", "", "", "Trust Store ARN")
	_elasticloadbalancingv2Cmd.Flags().StringSliceVarP(&_elasticloadbalancingv2TrustStoreArns, "trust-store-arns", "", nil, "Trust Store Arns")
	_elasticloadbalancingv2Cmd.Flags().StringVarP(&_elasticloadbalancingv2Type, "type", "", "", "Type")
	_elasticloadbalancingv2Cmd.Flags().StringVarP(&_elasticloadbalancingv2UnhealthyThresholdCount, "unhealthy-threshold-count", "", "", "Unhealthy Threshold Count")
	_elasticloadbalancingv2Cmd.Flags().StringVarP(&_elasticloadbalancingv2VpcId, "vpc-id", "", "", "VPC ID")

	_elasticloadbalancingv2Cmd.Flags().BoolVarP(&_elasticloadbalancingv2AddListenerCertificates, "add-listener-certificates", "", false, "Add Listener Certificates")
	_elasticloadbalancingv2Cmd.Flags().BoolVarP(&_elasticloadbalancingv2AddTags, "add-tags", "", false, "Add Tags")
	_elasticloadbalancingv2Cmd.Flags().BoolVarP(&_elasticloadbalancingv2AddTrustStoreRevocations, "add-trust-store-revocations", "", false, "Add Trust Store Revocations")
	_elasticloadbalancingv2Cmd.Flags().BoolVarP(&_elasticloadbalancingv2CreateListener, "create-listener", "", false, "Create Listener")
	_elasticloadbalancingv2Cmd.Flags().BoolVarP(&_elasticloadbalancingv2CreateLoadBalancer, "create-load-balancer", "", false, "Create Load Balancer")
	_elasticloadbalancingv2Cmd.Flags().BoolVarP(&_elasticloadbalancingv2CreateRule, "create-rule", "", false, "Create Rule")
	_elasticloadbalancingv2Cmd.Flags().BoolVarP(&_elasticloadbalancingv2CreateTargetGroup, "create-target-group", "", false, "Create Target Group")
	_elasticloadbalancingv2Cmd.Flags().BoolVarP(&_elasticloadbalancingv2CreateTrustStore, "create-trust-store", "", false, "Create Trust Store")
	_elasticloadbalancingv2Cmd.Flags().BoolVarP(&_elasticloadbalancingv2DeleteListener, "delete-listener", "", false, "Delete Listener")
	_elasticloadbalancingv2Cmd.Flags().BoolVarP(&_elasticloadbalancingv2DeleteLoadBalancer, "delete-load-balancer", "", false, "Delete Load Balancer")
	_elasticloadbalancingv2Cmd.Flags().BoolVarP(&_elasticloadbalancingv2DeleteRule, "delete-rule", "", false, "Delete Rule")
	_elasticloadbalancingv2Cmd.Flags().BoolVarP(&_elasticloadbalancingv2DeleteSharedTrustStoreAssociation, "delete-shared-trust-store-association", "", false, "Delete Shared Trust Store Association")
	_elasticloadbalancingv2Cmd.Flags().BoolVarP(&_elasticloadbalancingv2DeleteTargetGroup, "delete-target-group", "", false, "Delete Target Group")
	_elasticloadbalancingv2Cmd.Flags().BoolVarP(&_elasticloadbalancingv2DeleteTrustStore, "delete-trust-store", "", false, "Delete Trust Store")
	_elasticloadbalancingv2Cmd.Flags().BoolVarP(&_elasticloadbalancingv2DeregisterTargets, "deregister-targets", "", false, "Deregister Targets")
	_elasticloadbalancingv2Cmd.Flags().BoolVarP(&_elasticloadbalancingv2DescribeAccountLimits, "describe-account-limits", "", false, "Describe Account Limits")
	_elasticloadbalancingv2Cmd.Flags().BoolVarP(&_elasticloadbalancingv2DescribeCapacityReservation, "describe-capacity-reservation", "", false, "Describe Capacity Reservation")
	_elasticloadbalancingv2Cmd.Flags().BoolVarP(&_elasticloadbalancingv2DescribeListenerAttributes, "describe-listener-attributes", "", false, "Describe Listener Attributes")
	_elasticloadbalancingv2Cmd.Flags().BoolVarP(&_elasticloadbalancingv2DescribeListenerCertificates, "describe-listener-certificates", "", false, "Describe Listener Certificates")
	_elasticloadbalancingv2Cmd.Flags().BoolVarP(&_elasticloadbalancingv2DescribeListeners, "describe-listeners", "", false, "Describe Listeners")
	_elasticloadbalancingv2Cmd.Flags().BoolVarP(&_elasticloadbalancingv2DescribeLoadBalancerAttributes, "describe-load-balancer-attributes", "", false, "Describe Load Balancer Attributes")
	_elasticloadbalancingv2Cmd.Flags().BoolVarP(&_elasticloadbalancingv2DescribeLoadBalancers, "describe-load-balancers", "", false, "Describe Load Balancers")
	_elasticloadbalancingv2Cmd.Flags().BoolVarP(&_elasticloadbalancingv2DescribeRules, "describe-rules", "", false, "Describe Rules")
	_elasticloadbalancingv2Cmd.Flags().BoolVarP(&_elasticloadbalancingv2DescribeSSLPolicies, "describe-ssl-policies", "", false, "Describe SSL Policies")
	_elasticloadbalancingv2Cmd.Flags().BoolVarP(&_elasticloadbalancingv2DescribeTags, "describe-tags", "", false, "Describe Tags")
	_elasticloadbalancingv2Cmd.Flags().BoolVarP(&_elasticloadbalancingv2DescribeTargetGroupAttributes, "describe-target-group-attributes", "", false, "Describe Target Group Attributes")
	_elasticloadbalancingv2Cmd.Flags().BoolVarP(&_elasticloadbalancingv2DescribeTargetGroups, "describe-target-groups", "", false, "Describe Target Groups")
	_elasticloadbalancingv2Cmd.Flags().BoolVarP(&_elasticloadbalancingv2DescribeTargetHealth, "describe-target-health", "", false, "Describe Target Health")
	_elasticloadbalancingv2Cmd.Flags().BoolVarP(&_elasticloadbalancingv2DescribeTrustStoreAssociations, "describe-trust-store-associations", "", false, "Describe Trust Store Associations")
	_elasticloadbalancingv2Cmd.Flags().BoolVarP(&_elasticloadbalancingv2DescribeTrustStoreRevocations, "describe-trust-store-revocations", "", false, "Describe Trust Store Revocations")
	_elasticloadbalancingv2Cmd.Flags().BoolVarP(&_elasticloadbalancingv2DescribeTrustStores, "describe-trust-stores", "", false, "Describe Trust Stores")
	_elasticloadbalancingv2Cmd.Flags().BoolVarP(&_elasticloadbalancingv2GetResourcePolicy, "get-resource-policy", "", false, "Get Resource Policy")
	_elasticloadbalancingv2Cmd.Flags().BoolVarP(&_elasticloadbalancingv2GetTrustStoreCaCertificatesBundle, "get-trust-store-ca-certificates-bundle", "", false, "Get Trust Store Ca Certificates Bundle")
	_elasticloadbalancingv2Cmd.Flags().BoolVarP(&_elasticloadbalancingv2GetTrustStoreRevocationContent, "get-trust-store-revocation-content", "", false, "Get Trust Store Revocation Content")
	_elasticloadbalancingv2Cmd.Flags().BoolVarP(&_elasticloadbalancingv2ModifyCapacityReservation, "modify-capacity-reservation", "", false, "Modify Capacity Reservation")
	_elasticloadbalancingv2Cmd.Flags().BoolVarP(&_elasticloadbalancingv2ModifyIpPools, "modify-ip-pools", "", false, "Modify IP Pools")
	_elasticloadbalancingv2Cmd.Flags().BoolVarP(&_elasticloadbalancingv2ModifyListener, "modify-listener", "", false, "Modify Listener")
	_elasticloadbalancingv2Cmd.Flags().BoolVarP(&_elasticloadbalancingv2ModifyListenerAttributes, "modify-listener-attributes", "", false, "Modify Listener Attributes")
	_elasticloadbalancingv2Cmd.Flags().BoolVarP(&_elasticloadbalancingv2ModifyLoadBalancerAttributes, "modify-load-balancer-attributes", "", false, "Modify Load Balancer Attributes")
	_elasticloadbalancingv2Cmd.Flags().BoolVarP(&_elasticloadbalancingv2ModifyRule, "modify-rule", "", false, "Modify Rule")
	_elasticloadbalancingv2Cmd.Flags().BoolVarP(&_elasticloadbalancingv2ModifyTargetGroup, "modify-target-group", "", false, "Modify Target Group")
	_elasticloadbalancingv2Cmd.Flags().BoolVarP(&_elasticloadbalancingv2ModifyTargetGroupAttributes, "modify-target-group-attributes", "", false, "Modify Target Group Attributes")
	_elasticloadbalancingv2Cmd.Flags().BoolVarP(&_elasticloadbalancingv2ModifyTrustStore, "modify-trust-store", "", false, "Modify Trust Store")
	_elasticloadbalancingv2Cmd.Flags().BoolVarP(&_elasticloadbalancingv2RegisterTargets, "register-targets", "", false, "Register Targets")
	_elasticloadbalancingv2Cmd.Flags().BoolVarP(&_elasticloadbalancingv2RemoveListenerCertificates, "remove-listener-certificates", "", false, "Remove Listener Certificates")
	_elasticloadbalancingv2Cmd.Flags().BoolVarP(&_elasticloadbalancingv2RemoveTags, "remove-tags", "", false, "Remove Tags")
	_elasticloadbalancingv2Cmd.Flags().BoolVarP(&_elasticloadbalancingv2RemoveTrustStoreRevocations, "remove-trust-store-revocations", "", false, "Remove Trust Store Revocations")
	_elasticloadbalancingv2Cmd.Flags().BoolVarP(&_elasticloadbalancingv2SetIpAddressType, "set-ip-address-type", "", false, "Set IP Address Type")
	_elasticloadbalancingv2Cmd.Flags().BoolVarP(&_elasticloadbalancingv2SetRulePriorities, "set-rule-priorities", "", false, "Set Rule Priorities")
	_elasticloadbalancingv2Cmd.Flags().BoolVarP(&_elasticloadbalancingv2SetSecurityGroups, "set-security-groups", "", false, "Set Security Groups")
	_elasticloadbalancingv2Cmd.Flags().BoolVarP(&_elasticloadbalancingv2SetSubnets, "set-subnets", "", false, "Set Subnets")

}
