package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// route53resolverCmd represents the route53resolver command
var _route53resolverCmd = &cobra.Command{
	Use:   "route53resolver",
	Short: "AWS route53resolver CLI",
	Run: func(cmd *cobra.Command, args []string) {
		_awsOutput = resolveAWSOutput(_awsProfile, cmd.Flags().Changed("output"))
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := route53resolver.NewFromConfig(cfg)
		if _route53resolverAssociateFirewallRuleGroup {
			route53resolver_AssociateFirewallRuleGroup(cfg, client)
			return
		}
		if _route53resolverAssociateResolverEndpointIpAddress {
			route53resolver_AssociateResolverEndpointIpAddress(cfg, client)
			return
		}
		if _route53resolverAssociateResolverQueryLogConfig {
			route53resolver_AssociateResolverQueryLogConfig(cfg, client)
			return
		}
		if _route53resolverAssociateResolverRule {
			route53resolver_AssociateResolverRule(cfg, client)
			return
		}
		if _route53resolverCreateFirewallDomainList {
			route53resolver_CreateFirewallDomainList(cfg, client)
			return
		}
		if _route53resolverCreateFirewallRule {
			route53resolver_CreateFirewallRule(cfg, client)
			return
		}
		if _route53resolverCreateFirewallRuleGroup {
			route53resolver_CreateFirewallRuleGroup(cfg, client)
			return
		}
		if _route53resolverCreateOutpostResolver {
			route53resolver_CreateOutpostResolver(cfg, client)
			return
		}
		if _route53resolverCreateResolverEndpoint {
			route53resolver_CreateResolverEndpoint(cfg, client)
			return
		}
		if _route53resolverCreateResolverQueryLogConfig {
			route53resolver_CreateResolverQueryLogConfig(cfg, client)
			return
		}
		if _route53resolverCreateResolverRule {
			route53resolver_CreateResolverRule(cfg, client)
			return
		}
		if _route53resolverDeleteFirewallDomainList {
			route53resolver_DeleteFirewallDomainList(cfg, client)
			return
		}
		if _route53resolverDeleteFirewallRule {
			route53resolver_DeleteFirewallRule(cfg, client)
			return
		}
		if _route53resolverDeleteFirewallRuleGroup {
			route53resolver_DeleteFirewallRuleGroup(cfg, client)
			return
		}
		if _route53resolverDeleteOutpostResolver {
			route53resolver_DeleteOutpostResolver(cfg, client)
			return
		}
		if _route53resolverDeleteResolverEndpoint {
			route53resolver_DeleteResolverEndpoint(cfg, client)
			return
		}
		if _route53resolverDeleteResolverQueryLogConfig {
			route53resolver_DeleteResolverQueryLogConfig(cfg, client)
			return
		}
		if _route53resolverDeleteResolverRule {
			route53resolver_DeleteResolverRule(cfg, client)
			return
		}
		if _route53resolverDisassociateFirewallRuleGroup {
			route53resolver_DisassociateFirewallRuleGroup(cfg, client)
			return
		}
		if _route53resolverDisassociateResolverEndpointIpAddress {
			route53resolver_DisassociateResolverEndpointIpAddress(cfg, client)
			return
		}
		if _route53resolverDisassociateResolverQueryLogConfig {
			route53resolver_DisassociateResolverQueryLogConfig(cfg, client)
			return
		}
		if _route53resolverDisassociateResolverRule {
			route53resolver_DisassociateResolverRule(cfg, client)
			return
		}
		if _route53resolverGetFirewallConfig {
			route53resolver_GetFirewallConfig(cfg, client)
			return
		}
		if _route53resolverGetFirewallDomainList {
			route53resolver_GetFirewallDomainList(cfg, client)
			return
		}
		if _route53resolverGetFirewallRuleGroup {
			route53resolver_GetFirewallRuleGroup(cfg, client)
			return
		}
		if _route53resolverGetFirewallRuleGroupAssociation {
			route53resolver_GetFirewallRuleGroupAssociation(cfg, client)
			return
		}
		if _route53resolverGetFirewallRuleGroupPolicy {
			route53resolver_GetFirewallRuleGroupPolicy(cfg, client)
			return
		}
		if _route53resolverGetOutpostResolver {
			route53resolver_GetOutpostResolver(cfg, client)
			return
		}
		if _route53resolverGetResolverConfig {
			route53resolver_GetResolverConfig(cfg, client)
			return
		}
		if _route53resolverGetResolverDnssecConfig {
			route53resolver_GetResolverDnssecConfig(cfg, client)
			return
		}
		if _route53resolverGetResolverEndpoint {
			route53resolver_GetResolverEndpoint(cfg, client)
			return
		}
		if _route53resolverGetResolverQueryLogConfig {
			route53resolver_GetResolverQueryLogConfig(cfg, client)
			return
		}
		if _route53resolverGetResolverQueryLogConfigAssociation {
			route53resolver_GetResolverQueryLogConfigAssociation(cfg, client)
			return
		}
		if _route53resolverGetResolverQueryLogConfigPolicy {
			route53resolver_GetResolverQueryLogConfigPolicy(cfg, client)
			return
		}
		if _route53resolverGetResolverRule {
			route53resolver_GetResolverRule(cfg, client)
			return
		}
		if _route53resolverGetResolverRuleAssociation {
			route53resolver_GetResolverRuleAssociation(cfg, client)
			return
		}
		if _route53resolverGetResolverRulePolicy {
			route53resolver_GetResolverRulePolicy(cfg, client)
			return
		}
		if _route53resolverImportFirewallDomains {
			route53resolver_ImportFirewallDomains(cfg, client)
			return
		}
		if _route53resolverListFirewallConfigs {
			route53resolver_ListFirewallConfigs(cfg, client)
			return
		}
		if _route53resolverListFirewallDomainLists {
			route53resolver_ListFirewallDomainLists(cfg, client)
			return
		}
		if _route53resolverListFirewallDomains {
			route53resolver_ListFirewallDomains(cfg, client)
			return
		}
		if _route53resolverListFirewallRuleGroupAssociations {
			route53resolver_ListFirewallRuleGroupAssociations(cfg, client)
			return
		}
		if _route53resolverListFirewallRuleGroups {
			route53resolver_ListFirewallRuleGroups(cfg, client)
			return
		}
		if _route53resolverListFirewallRules {
			route53resolver_ListFirewallRules(cfg, client)
			return
		}
		if _route53resolverListOutpostResolvers {
			route53resolver_ListOutpostResolvers(cfg, client)
			return
		}
		if _route53resolverListResolverConfigs {
			route53resolver_ListResolverConfigs(cfg, client)
			return
		}
		if _route53resolverListResolverDnssecConfigs {
			route53resolver_ListResolverDnssecConfigs(cfg, client)
			return
		}
		if _route53resolverListResolverEndpointIpAddresses {
			route53resolver_ListResolverEndpointIpAddresses(cfg, client)
			return
		}
		if _route53resolverListResolverEndpoints {
			route53resolver_ListResolverEndpoints(cfg, client)
			return
		}
		if _route53resolverListResolverQueryLogConfigAssociations {
			route53resolver_ListResolverQueryLogConfigAssociations(cfg, client)
			return
		}
		if _route53resolverListResolverQueryLogConfigs {
			route53resolver_ListResolverQueryLogConfigs(cfg, client)
			return
		}
		if _route53resolverListResolverRuleAssociations {
			route53resolver_ListResolverRuleAssociations(cfg, client)
			return
		}
		if _route53resolverListResolverRules {
			route53resolver_ListResolverRules(cfg, client)
			return
		}
		if _route53resolverListTagsForResource {
			route53resolver_ListTagsForResource(cfg, client)
			return
		}
		if _route53resolverPutFirewallRuleGroupPolicy {
			route53resolver_PutFirewallRuleGroupPolicy(cfg, client)
			return
		}
		if _route53resolverPutResolverQueryLogConfigPolicy {
			route53resolver_PutResolverQueryLogConfigPolicy(cfg, client)
			return
		}
		if _route53resolverPutResolverRulePolicy {
			route53resolver_PutResolverRulePolicy(cfg, client)
			return
		}
		if _route53resolverTagResource {
			route53resolver_TagResource(cfg, client)
			return
		}
		if _route53resolverUntagResource {
			route53resolver_UntagResource(cfg, client)
			return
		}
		if _route53resolverUpdateFirewallConfig {
			route53resolver_UpdateFirewallConfig(cfg, client)
			return
		}
		if _route53resolverUpdateFirewallDomains {
			route53resolver_UpdateFirewallDomains(cfg, client)
			return
		}
		if _route53resolverUpdateFirewallRule {
			route53resolver_UpdateFirewallRule(cfg, client)
			return
		}
		if _route53resolverUpdateFirewallRuleGroupAssociation {
			route53resolver_UpdateFirewallRuleGroupAssociation(cfg, client)
			return
		}
		if _route53resolverUpdateOutpostResolver {
			route53resolver_UpdateOutpostResolver(cfg, client)
			return
		}
		if _route53resolverUpdateResolverConfig {
			route53resolver_UpdateResolverConfig(cfg, client)
			return
		}
		if _route53resolverUpdateResolverDnssecConfig {
			route53resolver_UpdateResolverDnssecConfig(cfg, client)
			return
		}
		if _route53resolverUpdateResolverEndpoint {
			route53resolver_UpdateResolverEndpoint(cfg, client)
			return
		}
		if _route53resolverUpdateResolverRule {
			route53resolver_UpdateResolverRule(cfg, client)
			return
		}

	},
}

var (
	_route53resolverAssociateFirewallRuleGroup             bool
	_route53resolverAssociateResolverEndpointIpAddress     bool
	_route53resolverAssociateResolverQueryLogConfig        bool
	_route53resolverAssociateResolverRule                  bool
	_route53resolverCreateFirewallDomainList               bool
	_route53resolverCreateFirewallRule                     bool
	_route53resolverCreateFirewallRuleGroup                bool
	_route53resolverCreateOutpostResolver                  bool
	_route53resolverCreateResolverEndpoint                 bool
	_route53resolverCreateResolverQueryLogConfig           bool
	_route53resolverCreateResolverRule                     bool
	_route53resolverDeleteFirewallDomainList               bool
	_route53resolverDeleteFirewallRule                     bool
	_route53resolverDeleteFirewallRuleGroup                bool
	_route53resolverDeleteOutpostResolver                  bool
	_route53resolverDeleteResolverEndpoint                 bool
	_route53resolverDeleteResolverQueryLogConfig           bool
	_route53resolverDeleteResolverRule                     bool
	_route53resolverDisassociateFirewallRuleGroup          bool
	_route53resolverDisassociateResolverEndpointIpAddress  bool
	_route53resolverDisassociateResolverQueryLogConfig     bool
	_route53resolverDisassociateResolverRule               bool
	_route53resolverGetFirewallConfig                      bool
	_route53resolverGetFirewallDomainList                  bool
	_route53resolverGetFirewallRuleGroup                   bool
	_route53resolverGetFirewallRuleGroupAssociation        bool
	_route53resolverGetFirewallRuleGroupPolicy             bool
	_route53resolverGetOutpostResolver                     bool
	_route53resolverGetResolverConfig                      bool
	_route53resolverGetResolverDnssecConfig                bool
	_route53resolverGetResolverEndpoint                    bool
	_route53resolverGetResolverQueryLogConfig              bool
	_route53resolverGetResolverQueryLogConfigAssociation   bool
	_route53resolverGetResolverQueryLogConfigPolicy        bool
	_route53resolverGetResolverRule                        bool
	_route53resolverGetResolverRuleAssociation             bool
	_route53resolverGetResolverRulePolicy                  bool
	_route53resolverImportFirewallDomains                  bool
	_route53resolverListFirewallConfigs                    bool
	_route53resolverListFirewallDomainLists                bool
	_route53resolverListFirewallDomains                    bool
	_route53resolverListFirewallRuleGroupAssociations      bool
	_route53resolverListFirewallRuleGroups                 bool
	_route53resolverListFirewallRules                      bool
	_route53resolverListOutpostResolvers                   bool
	_route53resolverListResolverConfigs                    bool
	_route53resolverListResolverDnssecConfigs              bool
	_route53resolverListResolverEndpointIpAddresses        bool
	_route53resolverListResolverEndpoints                  bool
	_route53resolverListResolverQueryLogConfigAssociations bool
	_route53resolverListResolverQueryLogConfigs            bool
	_route53resolverListResolverRuleAssociations           bool
	_route53resolverListResolverRules                      bool
	_route53resolverListTagsForResource                    bool
	_route53resolverPutFirewallRuleGroupPolicy             bool
	_route53resolverPutResolverQueryLogConfigPolicy        bool
	_route53resolverPutResolverRulePolicy                  bool
	_route53resolverTagResource                            bool
	_route53resolverUntagResource                          bool
	_route53resolverUpdateFirewallConfig                   bool
	_route53resolverUpdateFirewallDomains                  bool
	_route53resolverUpdateFirewallRule                     bool
	_route53resolverUpdateFirewallRuleGroupAssociation     bool
	_route53resolverUpdateOutpostResolver                  bool
	_route53resolverUpdateResolverConfig                   bool
	_route53resolverUpdateResolverDnssecConfig             bool
	_route53resolverUpdateResolverEndpoint                 bool
	_route53resolverUpdateResolverRule                     bool

	_route53resolverAction                              string
	_route53resolverArn                                 string
	_route53resolverAutodefinedReverseFlag              string
	_route53resolverBlockOverrideDnsType                string
	_route53resolverBlockOverrideDomain                 string
	_route53resolverBlockOverrideTtl                    string
	_route53resolverBlockResponse                       string
	_route53resolverConfidenceThreshold                 string
	_route53resolverConfig                              string
	_route53resolverCreatorRequestId                    string
	_route53resolverDelegationRecord                    string
	_route53resolverDestinationArn                      string
	_route53resolverDirection                           string
	_route53resolverDnsThreatProtection                 string
	_route53resolverDomainFileUrl                       string
	_route53resolverDomainName                          string
	_route53resolverDomains                             []string
	_route53resolverFilters                             string
	_route53resolverFirewallDomainListId                string
	_route53resolverFirewallDomainRedirectionAction     string
	_route53resolverFirewallFailOpen                    string
	_route53resolverFirewallRuleGroupAssociationId      string
	_route53resolverFirewallRuleGroupId                 string
	_route53resolverFirewallRuleGroupPolicy             string
	_route53resolverFirewallThreatProtectionId          string
	_route53resolverId                                  string
	_route53resolverInstanceCount                       string
	_route53resolverIpAddress                           string
	_route53resolverIpAddresses                         string
	_route53resolverMaxResults                          string
	_route53resolverMutationProtection                  string
	_route53resolverName                                string
	_route53resolverNextToken                           string
	_route53resolverOperation                           string
	_route53resolverOutpostArn                          string
	_route53resolverPreferredInstanceType               string
	_route53resolverPriority                            string
	_route53resolverProtocols                           string
	_route53resolverQtype                               string
	_route53resolverResolverEndpointId                  string
	_route53resolverResolverEndpointType                string
	_route53resolverResolverQueryLogConfigAssociationId string
	_route53resolverResolverQueryLogConfigId            string
	_route53resolverResolverQueryLogConfigPolicy        string
	_route53resolverResolverRuleAssociationId           string
	_route53resolverResolverRuleId                      string
	_route53resolverResolverRulePolicy                  string
	_route53resolverResourceArn                         string
	_route53resolverResourceId                          string
	_route53resolverRniEnhancedMetricsEnabled           string
	_route53resolverRuleType                            string
	_route53resolverSecurityGroupIds                    []string
	_route53resolverSortBy                              string
	_route53resolverSortOrder                           string
	_route53resolverStatus                              string
	_route53resolverTagKeys                             []string
	_route53resolverTags                                string
	_route53resolverTargetIps                           string
	_route53resolverTargetNameServerMetricsEnabled      string
	_route53resolverUpdateIpAddresses                   string
	_route53resolverValidation                          string
	_route53resolverVpcId                               string
	_route53resolverVPCId                               string
)

// Associates a FirewallRuleGroup with a VPC, to provide DNS filtering for the VPC.
func route53resolver_AssociateFirewallRuleGroup(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.AssociateFirewallRuleGroupInput{
		// CreatorRequestId: *string, // Required
		// FirewallRuleGroupId: *string, // Required
		// Name: *string, // Required
		// Priority: *int32, // Required
		// VpcId: *string, // Required
	}

	if len(_route53resolverCreatorRequestId) > 0 {
		input.CreatorRequestId = aws.String(_route53resolverCreatorRequestId)
	}
	if len(_route53resolverFirewallRuleGroupId) > 0 {
		input.FirewallRuleGroupId = aws.String(_route53resolverFirewallRuleGroupId)
	}
	if len(_route53resolverName) > 0 {
		input.Name = aws.String(_route53resolverName)
	}
	if len(_route53resolverPriority) > 0 {
		if err := assignInputField(input, "Priority", _route53resolverPriority); err != nil {
			log.Errorf("invalid --priority: %s", err.Error())
			return
		}
	}
	if len(_route53resolverVpcId) > 0 {
		input.VpcId = aws.String(_route53resolverVpcId)
	}
	if len(_route53resolverMutationProtection) > 0 {
		if err := assignInputField(input, "MutationProtection", _route53resolverMutationProtection); err != nil {
			log.Errorf("invalid --mutation-protection: %s", err.Error())
			return
		}
	}
	if len(_route53resolverTags) > 0 {
		if err := assignInputField(input, "Tags", _route53resolverTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.AssociateFirewallRuleGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds IP addresses to an inbound or an outbound Resolver endpoint. If you want
// to add more than one IP address, submit one AssociateResolverEndpointIpAddress
// request for each IP address.
//
// To remove an IP address from an endpoint, see [DisassociateResolverEndpointIpAddress].
//
// [DisassociateResolverEndpointIpAddress]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_DisassociateResolverEndpointIpAddress.html
func route53resolver_AssociateResolverEndpointIpAddress(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.AssociateResolverEndpointIpAddressInput{
		// IpAddress: *types.IpAddressUpdate, // Required
		// ResolverEndpointId: *string, // Required
	}

	if len(_route53resolverIpAddress) > 0 {
		if err := assignInputField(input, "IpAddress", _route53resolverIpAddress); err != nil {
			log.Errorf("invalid --ip-address: %s", err.Error())
			return
		}
	}
	if len(_route53resolverResolverEndpointId) > 0 {
		input.ResolverEndpointId = aws.String(_route53resolverResolverEndpointId)
	}

	if resp, err := client.AssociateResolverEndpointIpAddress(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates an Amazon VPC with a specified query logging configuration. Route 53
// Resolver logs DNS queries that originate in all of the Amazon VPCs that are
// associated with a specified query logging configuration. To associate more than
// one VPC with a configuration, submit one AssociateResolverQueryLogConfig
// request for each VPC.
//
// The VPCs that you associate with a query logging configuration must be in the
// same Region as the configuration.
//
// To remove a VPC from a query logging configuration, see [DisassociateResolverQueryLogConfig].
//
// [DisassociateResolverQueryLogConfig]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_DisassociateResolverQueryLogConfig.html
func route53resolver_AssociateResolverQueryLogConfig(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.AssociateResolverQueryLogConfigInput{
		// ResolverQueryLogConfigId: *string, // Required
		// ResourceId: *string, // Required
	}

	if len(_route53resolverResolverQueryLogConfigId) > 0 {
		input.ResolverQueryLogConfigId = aws.String(_route53resolverResolverQueryLogConfigId)
	}
	if len(_route53resolverResourceId) > 0 {
		input.ResourceId = aws.String(_route53resolverResourceId)
	}

	if resp, err := client.AssociateResolverQueryLogConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates a Resolver rule with a VPC. When you associate a rule with a VPC,
// Resolver forwards all DNS queries for the domain name that is specified in the
// rule and that originate in the VPC. The queries are forwarded to the IP
// addresses for the DNS resolvers that are specified in the rule. For more
// information about rules, see [CreateResolverRule].
//
// [CreateResolverRule]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_CreateResolverRule.html
func route53resolver_AssociateResolverRule(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.AssociateResolverRuleInput{
		// ResolverRuleId: *string, // Required
		// VPCId: *string, // Required
	}

	if len(_route53resolverResolverRuleId) > 0 {
		input.ResolverRuleId = aws.String(_route53resolverResolverRuleId)
	}
	if len(_route53resolverVPCId) > 0 {
		input.VPCId = aws.String(_route53resolverVPCId)
	}
	if len(_route53resolverName) > 0 {
		input.Name = aws.String(_route53resolverName)
	}

	if resp, err := client.AssociateResolverRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an empty firewall domain list for use in DNS Firewall rules. You can
// populate the domains for the new list with a file, using ImportFirewallDomains, or with domain
// strings, using UpdateFirewallDomains.
func route53resolver_CreateFirewallDomainList(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.CreateFirewallDomainListInput{
		// CreatorRequestId: *string, // Required
		// Name: *string, // Required
	}

	if len(_route53resolverCreatorRequestId) > 0 {
		input.CreatorRequestId = aws.String(_route53resolverCreatorRequestId)
	}
	if len(_route53resolverName) > 0 {
		input.Name = aws.String(_route53resolverName)
	}
	if len(_route53resolverTags) > 0 {
		if err := assignInputField(input, "Tags", _route53resolverTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateFirewallDomainList(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a single DNS Firewall rule in the specified rule group, using the
// specified domain list.
func route53resolver_CreateFirewallRule(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.CreateFirewallRuleInput{
		// Action: types.Action, // Required
		// CreatorRequestId: *string, // Required
		// FirewallRuleGroupId: *string, // Required
		// Name: *string, // Required
		// Priority: *int32, // Required
	}

	if len(_route53resolverAction) > 0 {
		if err := assignInputField(input, "Action", _route53resolverAction); err != nil {
			log.Errorf("invalid --action: %s", err.Error())
			return
		}
	}
	if len(_route53resolverCreatorRequestId) > 0 {
		input.CreatorRequestId = aws.String(_route53resolverCreatorRequestId)
	}
	if len(_route53resolverFirewallRuleGroupId) > 0 {
		input.FirewallRuleGroupId = aws.String(_route53resolverFirewallRuleGroupId)
	}
	if len(_route53resolverName) > 0 {
		input.Name = aws.String(_route53resolverName)
	}
	if len(_route53resolverPriority) > 0 {
		if err := assignInputField(input, "Priority", _route53resolverPriority); err != nil {
			log.Errorf("invalid --priority: %s", err.Error())
			return
		}
	}
	if len(_route53resolverBlockOverrideDnsType) > 0 {
		if err := assignInputField(input, "BlockOverrideDnsType", _route53resolverBlockOverrideDnsType); err != nil {
			log.Errorf("invalid --block-override-dns-type: %s", err.Error())
			return
		}
	}
	if len(_route53resolverBlockOverrideDomain) > 0 {
		input.BlockOverrideDomain = aws.String(_route53resolverBlockOverrideDomain)
	}
	if len(_route53resolverBlockOverrideTtl) > 0 {
		if err := assignInputField(input, "BlockOverrideTtl", _route53resolverBlockOverrideTtl); err != nil {
			log.Errorf("invalid --block-override-ttl: %s", err.Error())
			return
		}
	}
	if len(_route53resolverBlockResponse) > 0 {
		if err := assignInputField(input, "BlockResponse", _route53resolverBlockResponse); err != nil {
			log.Errorf("invalid --block-response: %s", err.Error())
			return
		}
	}
	if len(_route53resolverConfidenceThreshold) > 0 {
		if err := assignInputField(input, "ConfidenceThreshold", _route53resolverConfidenceThreshold); err != nil {
			log.Errorf("invalid --confidence-threshold: %s", err.Error())
			return
		}
	}
	if len(_route53resolverDnsThreatProtection) > 0 {
		if err := assignInputField(input, "DnsThreatProtection", _route53resolverDnsThreatProtection); err != nil {
			log.Errorf("invalid --dns-threat-protection: %s", err.Error())
			return
		}
	}
	if len(_route53resolverFirewallDomainListId) > 0 {
		input.FirewallDomainListId = aws.String(_route53resolverFirewallDomainListId)
	}
	if len(_route53resolverFirewallDomainRedirectionAction) > 0 {
		if err := assignInputField(input, "FirewallDomainRedirectionAction", _route53resolverFirewallDomainRedirectionAction); err != nil {
			log.Errorf("invalid --firewall-domain-redirection-action: %s", err.Error())
			return
		}
	}
	if len(_route53resolverQtype) > 0 {
		input.Qtype = aws.String(_route53resolverQtype)
	}

	if resp, err := client.CreateFirewallRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an empty DNS Firewall rule group for filtering DNS network traffic in a
// VPC. You can add rules to the new rule group by calling CreateFirewallRule.
func route53resolver_CreateFirewallRuleGroup(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.CreateFirewallRuleGroupInput{
		// CreatorRequestId: *string, // Required
		// Name: *string, // Required
	}

	if len(_route53resolverCreatorRequestId) > 0 {
		input.CreatorRequestId = aws.String(_route53resolverCreatorRequestId)
	}
	if len(_route53resolverName) > 0 {
		input.Name = aws.String(_route53resolverName)
	}
	if len(_route53resolverTags) > 0 {
		if err := assignInputField(input, "Tags", _route53resolverTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateFirewallRuleGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a Route 53 Resolver on an Outpost.
func route53resolver_CreateOutpostResolver(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.CreateOutpostResolverInput{
		// CreatorRequestId: *string, // Required
		// Name: *string, // Required
		// OutpostArn: *string, // Required
		// PreferredInstanceType: *string, // Required
	}

	if len(_route53resolverCreatorRequestId) > 0 {
		input.CreatorRequestId = aws.String(_route53resolverCreatorRequestId)
	}
	if len(_route53resolverName) > 0 {
		input.Name = aws.String(_route53resolverName)
	}
	if len(_route53resolverOutpostArn) > 0 {
		input.OutpostArn = aws.String(_route53resolverOutpostArn)
	}
	if len(_route53resolverPreferredInstanceType) > 0 {
		input.PreferredInstanceType = aws.String(_route53resolverPreferredInstanceType)
	}
	if len(_route53resolverInstanceCount) > 0 {
		if err := assignInputField(input, "InstanceCount", _route53resolverInstanceCount); err != nil {
			log.Errorf("invalid --instance-count: %s", err.Error())
			return
		}
	}
	if len(_route53resolverTags) > 0 {
		if err := assignInputField(input, "Tags", _route53resolverTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateOutpostResolver(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a Resolver endpoint. There are two types of Resolver endpoints, inbound
// and outbound:
//
// - An inbound Resolver endpoint forwards DNS queries to the DNS service for a
// VPC from your network.
//
// - An outbound Resolver endpoint forwards DNS queries from the DNS service for
// a VPC to your network.
func route53resolver_CreateResolverEndpoint(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.CreateResolverEndpointInput{
		// CreatorRequestId: *string, // Required
		// Direction: types.ResolverEndpointDirection, // Required
		// IpAddresses: []types.IpAddressRequest, // Required
		// SecurityGroupIds: []string, // Required
	}

	if len(_route53resolverCreatorRequestId) > 0 {
		input.CreatorRequestId = aws.String(_route53resolverCreatorRequestId)
	}
	if len(_route53resolverDirection) > 0 {
		if err := assignInputField(input, "Direction", _route53resolverDirection); err != nil {
			log.Errorf("invalid --direction: %s", err.Error())
			return
		}
	}
	if len(_route53resolverIpAddresses) > 0 {
		if err := assignInputField(input, "IpAddresses", _route53resolverIpAddresses); err != nil {
			log.Errorf("invalid --ip-addresses: %s", err.Error())
			return
		}
	}
	if len(_route53resolverSecurityGroupIds) > 0 {
		input.SecurityGroupIds = append([]string(nil), _route53resolverSecurityGroupIds...)
	}
	if len(_route53resolverName) > 0 {
		input.Name = aws.String(_route53resolverName)
	}
	if len(_route53resolverOutpostArn) > 0 {
		input.OutpostArn = aws.String(_route53resolverOutpostArn)
	}
	if len(_route53resolverPreferredInstanceType) > 0 {
		input.PreferredInstanceType = aws.String(_route53resolverPreferredInstanceType)
	}
	if len(_route53resolverProtocols) > 0 {
		if err := assignInputField(input, "Protocols", _route53resolverProtocols); err != nil {
			log.Errorf("invalid --protocols: %s", err.Error())
			return
		}
	}
	if len(_route53resolverResolverEndpointType) > 0 {
		if err := assignInputField(input, "ResolverEndpointType", _route53resolverResolverEndpointType); err != nil {
			log.Errorf("invalid --resolver-endpoint-type: %s", err.Error())
			return
		}
	}
	if len(_route53resolverRniEnhancedMetricsEnabled) > 0 {
		if err := assignInputField(input, "RniEnhancedMetricsEnabled", _route53resolverRniEnhancedMetricsEnabled); err != nil {
			log.Errorf("invalid --rni-enhanced-metrics-enabled: %s", err.Error())
			return
		}
	}
	if len(_route53resolverTags) > 0 {
		if err := assignInputField(input, "Tags", _route53resolverTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_route53resolverTargetNameServerMetricsEnabled) > 0 {
		if err := assignInputField(input, "TargetNameServerMetricsEnabled", _route53resolverTargetNameServerMetricsEnabled); err != nil {
			log.Errorf("invalid --target-name-server-metrics-enabled: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateResolverEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a Resolver query logging configuration, which defines where you want
// Resolver to save DNS query logs that originate in your VPCs. Resolver can log
// queries only for VPCs that are in the same Region as the query logging
// configuration.
//
// To specify which VPCs you want to log queries for, you use
// AssociateResolverQueryLogConfig . For more information, see [AssociateResolverQueryLogConfig].
//
// You can optionally use Resource Access Manager (RAM) to share a query logging
// configuration with other Amazon Web Services accounts. The other accounts can
// then associate VPCs with the configuration. The query logs that Resolver creates
// for a configuration include all DNS queries that originate in all VPCs that are
// associated with the configuration.
//
// [AssociateResolverQueryLogConfig]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_AssociateResolverQueryLogConfig.html
func route53resolver_CreateResolverQueryLogConfig(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.CreateResolverQueryLogConfigInput{
		// CreatorRequestId: *string, // Required
		// DestinationArn: *string, // Required
		// Name: *string, // Required
	}

	if len(_route53resolverCreatorRequestId) > 0 {
		input.CreatorRequestId = aws.String(_route53resolverCreatorRequestId)
	}
	if len(_route53resolverDestinationArn) > 0 {
		input.DestinationArn = aws.String(_route53resolverDestinationArn)
	}
	if len(_route53resolverName) > 0 {
		input.Name = aws.String(_route53resolverName)
	}
	if len(_route53resolverTags) > 0 {
		if err := assignInputField(input, "Tags", _route53resolverTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateResolverQueryLogConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// For DNS queries that originate in your VPCs, specifies which Resolver endpoint
// the queries pass through, one domain name that you want to forward to your
// network, and the IP addresses of the DNS resolvers in your network.
func route53resolver_CreateResolverRule(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.CreateResolverRuleInput{
		// CreatorRequestId: *string, // Required
		// RuleType: types.RuleTypeOption, // Required
	}

	if len(_route53resolverCreatorRequestId) > 0 {
		input.CreatorRequestId = aws.String(_route53resolverCreatorRequestId)
	}
	if len(_route53resolverRuleType) > 0 {
		if err := assignInputField(input, "RuleType", _route53resolverRuleType); err != nil {
			log.Errorf("invalid --rule-type: %s", err.Error())
			return
		}
	}
	if len(_route53resolverDelegationRecord) > 0 {
		input.DelegationRecord = aws.String(_route53resolverDelegationRecord)
	}
	if len(_route53resolverDomainName) > 0 {
		input.DomainName = aws.String(_route53resolverDomainName)
	}
	if len(_route53resolverName) > 0 {
		input.Name = aws.String(_route53resolverName)
	}
	if len(_route53resolverResolverEndpointId) > 0 {
		input.ResolverEndpointId = aws.String(_route53resolverResolverEndpointId)
	}
	if len(_route53resolverTags) > 0 {
		if err := assignInputField(input, "Tags", _route53resolverTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_route53resolverTargetIps) > 0 {
		if err := assignInputField(input, "TargetIps", _route53resolverTargetIps); err != nil {
			log.Errorf("invalid --target-ips: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateResolverRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified domain list.
func route53resolver_DeleteFirewallDomainList(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.DeleteFirewallDomainListInput{
		// FirewallDomainListId: *string, // Required
	}

	if len(_route53resolverFirewallDomainListId) > 0 {
		input.FirewallDomainListId = aws.String(_route53resolverFirewallDomainListId)
	}

	if resp, err := client.DeleteFirewallDomainList(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified firewall rule.
func route53resolver_DeleteFirewallRule(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.DeleteFirewallRuleInput{
		// FirewallRuleGroupId: *string, // Required
	}

	if len(_route53resolverFirewallRuleGroupId) > 0 {
		input.FirewallRuleGroupId = aws.String(_route53resolverFirewallRuleGroupId)
	}
	if len(_route53resolverFirewallDomainListId) > 0 {
		input.FirewallDomainListId = aws.String(_route53resolverFirewallDomainListId)
	}
	if len(_route53resolverFirewallThreatProtectionId) > 0 {
		input.FirewallThreatProtectionId = aws.String(_route53resolverFirewallThreatProtectionId)
	}
	if len(_route53resolverQtype) > 0 {
		input.Qtype = aws.String(_route53resolverQtype)
	}

	if resp, err := client.DeleteFirewallRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified firewall rule group.
func route53resolver_DeleteFirewallRuleGroup(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.DeleteFirewallRuleGroupInput{
		// FirewallRuleGroupId: *string, // Required
	}

	if len(_route53resolverFirewallRuleGroupId) > 0 {
		input.FirewallRuleGroupId = aws.String(_route53resolverFirewallRuleGroupId)
	}

	if resp, err := client.DeleteFirewallRuleGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a Resolver on the Outpost.
func route53resolver_DeleteOutpostResolver(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.DeleteOutpostResolverInput{
		// Id: *string, // Required
	}

	if len(_route53resolverId) > 0 {
		input.Id = aws.String(_route53resolverId)
	}

	if resp, err := client.DeleteOutpostResolver(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a Resolver endpoint. The effect of deleting a Resolver endpoint depends
// on whether it's an inbound or an outbound Resolver endpoint:
//
// - Inbound: DNS queries from your network are no longer routed to the DNS
// service for the specified VPC.
//
// - Outbound: DNS queries from a VPC are no longer routed to your network.
func route53resolver_DeleteResolverEndpoint(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.DeleteResolverEndpointInput{
		// ResolverEndpointId: *string, // Required
	}

	if len(_route53resolverResolverEndpointId) > 0 {
		input.ResolverEndpointId = aws.String(_route53resolverResolverEndpointId)
	}

	if resp, err := client.DeleteResolverEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a query logging configuration. When you delete a configuration,
// Resolver stops logging DNS queries for all of the Amazon VPCs that are
// associated with the configuration. This also applies if the query logging
// configuration is shared with other Amazon Web Services accounts, and the other
// accounts have associated VPCs with the shared configuration.
//
// Before you can delete a query logging configuration, you must first
// disassociate all VPCs from the configuration. See [DisassociateResolverQueryLogConfig].
//
// If you used Resource Access Manager (RAM) to share a query logging
// configuration with other accounts, you must stop sharing the configuration
// before you can delete a configuration. The accounts that you shared the
// configuration with can first disassociate VPCs that they associated with the
// configuration, but that's not necessary. If you stop sharing the configuration,
// those VPCs are automatically disassociated from the configuration.
//
// [DisassociateResolverQueryLogConfig]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_DisassociateResolverQueryLogConfig.html
func route53resolver_DeleteResolverQueryLogConfig(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.DeleteResolverQueryLogConfigInput{
		// ResolverQueryLogConfigId: *string, // Required
	}

	if len(_route53resolverResolverQueryLogConfigId) > 0 {
		input.ResolverQueryLogConfigId = aws.String(_route53resolverResolverQueryLogConfigId)
	}

	if resp, err := client.DeleteResolverQueryLogConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a Resolver rule. Before you can delete a Resolver rule, you must
// disassociate it from all the VPCs that you associated the Resolver rule with.
// For more information, see [DisassociateResolverRule].
//
// [DisassociateResolverRule]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_DisassociateResolverRule.html
func route53resolver_DeleteResolverRule(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.DeleteResolverRuleInput{
		// ResolverRuleId: *string, // Required
	}

	if len(_route53resolverResolverRuleId) > 0 {
		input.ResolverRuleId = aws.String(_route53resolverResolverRuleId)
	}

	if resp, err := client.DeleteResolverRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates a FirewallRuleGroup from a VPC, to remove DNS filtering from the VPC.
func route53resolver_DisassociateFirewallRuleGroup(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.DisassociateFirewallRuleGroupInput{
		// FirewallRuleGroupAssociationId: *string, // Required
	}

	if len(_route53resolverFirewallRuleGroupAssociationId) > 0 {
		input.FirewallRuleGroupAssociationId = aws.String(_route53resolverFirewallRuleGroupAssociationId)
	}

	if resp, err := client.DisassociateFirewallRuleGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes IP addresses from an inbound or an outbound Resolver endpoint. If you
// want to remove more than one IP address, submit one
// DisassociateResolverEndpointIpAddress request for each IP address.
//
// To add an IP address to an endpoint, see [AssociateResolverEndpointIpAddress].
//
// [AssociateResolverEndpointIpAddress]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_AssociateResolverEndpointIpAddress.html
func route53resolver_DisassociateResolverEndpointIpAddress(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.DisassociateResolverEndpointIpAddressInput{
		// IpAddress: *types.IpAddressUpdate, // Required
		// ResolverEndpointId: *string, // Required
	}

	if len(_route53resolverIpAddress) > 0 {
		if err := assignInputField(input, "IpAddress", _route53resolverIpAddress); err != nil {
			log.Errorf("invalid --ip-address: %s", err.Error())
			return
		}
	}
	if len(_route53resolverResolverEndpointId) > 0 {
		input.ResolverEndpointId = aws.String(_route53resolverResolverEndpointId)
	}

	if resp, err := client.DisassociateResolverEndpointIpAddress(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates a VPC from a query logging configuration.
// Before you can delete a query logging configuration, you must first
// disassociate all VPCs from the configuration. If you used Resource Access
// Manager (RAM) to share a query logging configuration with other accounts, VPCs
// can be disassociated from the configuration in the following ways:
//
// - The accounts that you shared the configuration with can disassociate VPCs
// from the configuration.
//
// - You can stop sharing the configuration.
func route53resolver_DisassociateResolverQueryLogConfig(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.DisassociateResolverQueryLogConfigInput{
		// ResolverQueryLogConfigId: *string, // Required
		// ResourceId: *string, // Required
	}

	if len(_route53resolverResolverQueryLogConfigId) > 0 {
		input.ResolverQueryLogConfigId = aws.String(_route53resolverResolverQueryLogConfigId)
	}
	if len(_route53resolverResourceId) > 0 {
		input.ResourceId = aws.String(_route53resolverResourceId)
	}

	if resp, err := client.DisassociateResolverQueryLogConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the association between a specified Resolver rule and a specified VPC.
// If you disassociate a Resolver rule from a VPC, Resolver stops forwarding DNS
// queries for the domain name that you specified in the Resolver rule.
func route53resolver_DisassociateResolverRule(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.DisassociateResolverRuleInput{
		// ResolverRuleId: *string, // Required
		// VPCId: *string, // Required
	}

	if len(_route53resolverResolverRuleId) > 0 {
		input.ResolverRuleId = aws.String(_route53resolverResolverRuleId)
	}
	if len(_route53resolverVPCId) > 0 {
		input.VPCId = aws.String(_route53resolverVPCId)
	}

	if resp, err := client.DisassociateResolverRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the configuration of the firewall behavior provided by DNS Firewall
// for a single VPC from Amazon Virtual Private Cloud (Amazon VPC).
func route53resolver_GetFirewallConfig(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.GetFirewallConfigInput{
		// ResourceId: *string, // Required
	}

	if len(_route53resolverResourceId) > 0 {
		input.ResourceId = aws.String(_route53resolverResourceId)
	}

	if resp, err := client.GetFirewallConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the specified firewall domain list.
func route53resolver_GetFirewallDomainList(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.GetFirewallDomainListInput{
		// FirewallDomainListId: *string, // Required
	}

	if len(_route53resolverFirewallDomainListId) > 0 {
		input.FirewallDomainListId = aws.String(_route53resolverFirewallDomainListId)
	}

	if resp, err := client.GetFirewallDomainList(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the specified firewall rule group.
func route53resolver_GetFirewallRuleGroup(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.GetFirewallRuleGroupInput{
		// FirewallRuleGroupId: *string, // Required
	}

	if len(_route53resolverFirewallRuleGroupId) > 0 {
		input.FirewallRuleGroupId = aws.String(_route53resolverFirewallRuleGroupId)
	}

	if resp, err := client.GetFirewallRuleGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a firewall rule group association, which enables DNS filtering for a
// VPC with one rule group. A VPC can have more than one firewall rule group
// association, and a rule group can be associated with more than one VPC.
func route53resolver_GetFirewallRuleGroupAssociation(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.GetFirewallRuleGroupAssociationInput{
		// FirewallRuleGroupAssociationId: *string, // Required
	}

	if len(_route53resolverFirewallRuleGroupAssociationId) > 0 {
		input.FirewallRuleGroupAssociationId = aws.String(_route53resolverFirewallRuleGroupAssociationId)
	}

	if resp, err := client.GetFirewallRuleGroupAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the Identity and Access Management (Amazon Web Services IAM) policy for
// sharing the specified rule group. You can use the policy to share the rule group
// using Resource Access Manager (RAM).
func route53resolver_GetFirewallRuleGroupPolicy(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.GetFirewallRuleGroupPolicyInput{
		// Arn: *string, // Required
	}

	if len(_route53resolverArn) > 0 {
		input.Arn = aws.String(_route53resolverArn)
	}

	if resp, err := client.GetFirewallRuleGroupPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a specified Resolver on the Outpost, such as its
// instance count and type, name, and the current status of the Resolver.
func route53resolver_GetOutpostResolver(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.GetOutpostResolverInput{
		// Id: *string, // Required
	}

	if len(_route53resolverId) > 0 {
		input.Id = aws.String(_route53resolverId)
	}

	if resp, err := client.GetOutpostResolver(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the behavior configuration of Route 53 Resolver behavior for a single
// VPC from Amazon Virtual Private Cloud.
func route53resolver_GetResolverConfig(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.GetResolverConfigInput{
		// ResourceId: *string, // Required
	}

	if len(_route53resolverResourceId) > 0 {
		input.ResourceId = aws.String(_route53resolverResourceId)
	}

	if resp, err := client.GetResolverConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets DNSSEC validation information for a specified resource.
func route53resolver_GetResolverDnssecConfig(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.GetResolverDnssecConfigInput{
		// ResourceId: *string, // Required
	}

	if len(_route53resolverResourceId) > 0 {
		input.ResourceId = aws.String(_route53resolverResourceId)
	}

	if resp, err := client.GetResolverDnssecConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a specified Resolver endpoint, such as whether it's an
// inbound or an outbound Resolver endpoint, and the current status of the
// endpoint.
func route53resolver_GetResolverEndpoint(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.GetResolverEndpointInput{
		// ResolverEndpointId: *string, // Required
	}

	if len(_route53resolverResolverEndpointId) > 0 {
		input.ResolverEndpointId = aws.String(_route53resolverResolverEndpointId)
	}

	if resp, err := client.GetResolverEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a specified Resolver query logging configuration, such
// as the number of VPCs that the configuration is logging queries for and the
// location that logs are sent to.
func route53resolver_GetResolverQueryLogConfig(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.GetResolverQueryLogConfigInput{
		// ResolverQueryLogConfigId: *string, // Required
	}

	if len(_route53resolverResolverQueryLogConfigId) > 0 {
		input.ResolverQueryLogConfigId = aws.String(_route53resolverResolverQueryLogConfigId)
	}

	if resp, err := client.GetResolverQueryLogConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a specified association between a Resolver query logging
// configuration and an Amazon VPC. When you associate a VPC with a query logging
// configuration, Resolver logs DNS queries that originate in that VPC.
func route53resolver_GetResolverQueryLogConfigAssociation(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.GetResolverQueryLogConfigAssociationInput{
		// ResolverQueryLogConfigAssociationId: *string, // Required
	}

	if len(_route53resolverResolverQueryLogConfigAssociationId) > 0 {
		input.ResolverQueryLogConfigAssociationId = aws.String(_route53resolverResolverQueryLogConfigAssociationId)
	}

	if resp, err := client.GetResolverQueryLogConfigAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a query logging policy. A query logging policy specifies
// the Resolver query logging operations and resources that you want to allow
// another Amazon Web Services account to be able to use.
func route53resolver_GetResolverQueryLogConfigPolicy(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.GetResolverQueryLogConfigPolicyInput{
		// Arn: *string, // Required
	}

	if len(_route53resolverArn) > 0 {
		input.Arn = aws.String(_route53resolverArn)
	}

	if resp, err := client.GetResolverQueryLogConfigPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a specified Resolver rule, such as the domain name that
// the rule forwards DNS queries for and the ID of the outbound Resolver endpoint
// that the rule is associated with.
func route53resolver_GetResolverRule(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.GetResolverRuleInput{
		// ResolverRuleId: *string, // Required
	}

	if len(_route53resolverResolverRuleId) > 0 {
		input.ResolverRuleId = aws.String(_route53resolverResolverRuleId)
	}

	if resp, err := client.GetResolverRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about an association between a specified Resolver rule and a
// VPC. You associate a Resolver rule and a VPC using [AssociateResolverRule].
//
// [AssociateResolverRule]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_AssociateResolverRule.html
func route53resolver_GetResolverRuleAssociation(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.GetResolverRuleAssociationInput{
		// ResolverRuleAssociationId: *string, // Required
	}

	if len(_route53resolverResolverRuleAssociationId) > 0 {
		input.ResolverRuleAssociationId = aws.String(_route53resolverResolverRuleAssociationId)
	}

	if resp, err := client.GetResolverRuleAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about the Resolver rule policy for a specified rule. A
// Resolver rule policy includes the rule that you want to share with another
// account, the account that you want to share the rule with, and the Resolver
// operations that you want to allow the account to use.
func route53resolver_GetResolverRulePolicy(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.GetResolverRulePolicyInput{
		// Arn: *string, // Required
	}

	if len(_route53resolverArn) > 0 {
		input.Arn = aws.String(_route53resolverArn)
	}

	if resp, err := client.GetResolverRulePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Imports domain names from a file into a domain list, for use in a DNS firewall
// rule group.
//
// Each domain specification in your domain list must satisfy the following
// requirements:
//
// - It can optionally start with * (asterisk).
//
// - With the exception of the optional starting asterisk, it must only contain
// the following characters: A-Z , a-z , 0-9 , - (hyphen).
//
// - It must be from 1-255 characters in length.
func route53resolver_ImportFirewallDomains(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.ImportFirewallDomainsInput{
		// DomainFileUrl: *string, // Required
		// FirewallDomainListId: *string, // Required
		// Operation: types.FirewallDomainImportOperation, // Required
	}

	if len(_route53resolverDomainFileUrl) > 0 {
		input.DomainFileUrl = aws.String(_route53resolverDomainFileUrl)
	}
	if len(_route53resolverFirewallDomainListId) > 0 {
		input.FirewallDomainListId = aws.String(_route53resolverFirewallDomainListId)
	}
	if len(_route53resolverOperation) > 0 {
		if err := assignInputField(input, "Operation", _route53resolverOperation); err != nil {
			log.Errorf("invalid --operation: %s", err.Error())
			return
		}
	}

	if resp, err := client.ImportFirewallDomains(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the firewall configurations that you have defined. DNS Firewall uses
// the configurations to manage firewall behavior for your VPCs.
//
// A single call might return only a partial list of the configurations. For
// information, see MaxResults .
func route53resolver_ListFirewallConfigs(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.ListFirewallConfigsInput{}

	if len(_route53resolverMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _route53resolverMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_route53resolverNextToken) > 0 {
		input.NextToken = aws.String(_route53resolverNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListFirewallConfigs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*route53resolver.ListFirewallConfigsOutput
	p := route53resolver.NewListFirewallConfigsPaginator(client, input)
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

// Retrieves the firewall domain lists that you have defined. For each firewall
// domain list, you can retrieve the domains that are defined for a list by calling
// ListFirewallDomains.
//
// A single call to this list operation might return only a partial list of the
// domain lists. For information, see MaxResults .
func route53resolver_ListFirewallDomainLists(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.ListFirewallDomainListsInput{}

	if len(_route53resolverMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _route53resolverMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_route53resolverNextToken) > 0 {
		input.NextToken = aws.String(_route53resolverNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListFirewallDomainLists(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*route53resolver.ListFirewallDomainListsOutput
	p := route53resolver.NewListFirewallDomainListsPaginator(client, input)
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

// Retrieves the domains that you have defined for the specified firewall domain
// list.
//
// A single call might return only a partial list of the domains. For information,
// see MaxResults .
func route53resolver_ListFirewallDomains(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.ListFirewallDomainsInput{
		// FirewallDomainListId: *string, // Required
	}

	if len(_route53resolverFirewallDomainListId) > 0 {
		input.FirewallDomainListId = aws.String(_route53resolverFirewallDomainListId)
	}
	if len(_route53resolverMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _route53resolverMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_route53resolverNextToken) > 0 {
		input.NextToken = aws.String(_route53resolverNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListFirewallDomains(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*route53resolver.ListFirewallDomainsOutput
	p := route53resolver.NewListFirewallDomainsPaginator(client, input)
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

// Retrieves the firewall rule group associations that you have defined. Each
// association enables DNS filtering for a VPC with one rule group.
//
// A single call might return only a partial list of the associations. For
// information, see MaxResults .
func route53resolver_ListFirewallRuleGroupAssociations(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.ListFirewallRuleGroupAssociationsInput{}

	if len(_route53resolverFirewallRuleGroupId) > 0 {
		input.FirewallRuleGroupId = aws.String(_route53resolverFirewallRuleGroupId)
	}
	if len(_route53resolverMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _route53resolverMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_route53resolverNextToken) > 0 {
		input.NextToken = aws.String(_route53resolverNextToken)
	}
	if len(_route53resolverPriority) > 0 {
		if err := assignInputField(input, "Priority", _route53resolverPriority); err != nil {
			log.Errorf("invalid --priority: %s", err.Error())
			return
		}
	}
	if len(_route53resolverStatus) > 0 {
		if err := assignInputField(input, "Status", _route53resolverStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_route53resolverVpcId) > 0 {
		input.VpcId = aws.String(_route53resolverVpcId)
	}

	if disablePaginator() {
		if resp, err := client.ListFirewallRuleGroupAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*route53resolver.ListFirewallRuleGroupAssociationsOutput
	p := route53resolver.NewListFirewallRuleGroupAssociationsPaginator(client, input)
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

// Retrieves the minimal high-level information for the rule groups that you have
// defined.
//
// A single call might return only a partial list of the rule groups. For
// information, see MaxResults .
func route53resolver_ListFirewallRuleGroups(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.ListFirewallRuleGroupsInput{}

	if len(_route53resolverMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _route53resolverMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_route53resolverNextToken) > 0 {
		input.NextToken = aws.String(_route53resolverNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListFirewallRuleGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*route53resolver.ListFirewallRuleGroupsOutput
	p := route53resolver.NewListFirewallRuleGroupsPaginator(client, input)
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

// Retrieves the firewall rules that you have defined for the specified firewall
// rule group. DNS Firewall uses the rules in a rule group to filter DNS network
// traffic for a VPC.
//
// A single call might return only a partial list of the rules. For information,
// see MaxResults .
func route53resolver_ListFirewallRules(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.ListFirewallRulesInput{
		// FirewallRuleGroupId: *string, // Required
	}

	if len(_route53resolverFirewallRuleGroupId) > 0 {
		input.FirewallRuleGroupId = aws.String(_route53resolverFirewallRuleGroupId)
	}
	if len(_route53resolverAction) > 0 {
		if err := assignInputField(input, "Action", _route53resolverAction); err != nil {
			log.Errorf("invalid --action: %s", err.Error())
			return
		}
	}
	if len(_route53resolverMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _route53resolverMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_route53resolverNextToken) > 0 {
		input.NextToken = aws.String(_route53resolverNextToken)
	}
	if len(_route53resolverPriority) > 0 {
		if err := assignInputField(input, "Priority", _route53resolverPriority); err != nil {
			log.Errorf("invalid --priority: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListFirewallRules(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*route53resolver.ListFirewallRulesOutput
	p := route53resolver.NewListFirewallRulesPaginator(client, input)
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

// Lists all the Resolvers on Outposts that were created using the current Amazon
// Web Services account.
func route53resolver_ListOutpostResolvers(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.ListOutpostResolversInput{}

	if len(_route53resolverMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _route53resolverMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_route53resolverNextToken) > 0 {
		input.NextToken = aws.String(_route53resolverNextToken)
	}
	if len(_route53resolverOutpostArn) > 0 {
		input.OutpostArn = aws.String(_route53resolverOutpostArn)
	}

	if disablePaginator() {
		if resp, err := client.ListOutpostResolvers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*route53resolver.ListOutpostResolversOutput
	p := route53resolver.NewListOutpostResolversPaginator(client, input)
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

// Retrieves the Resolver configurations that you have defined. Route 53 Resolver
// uses the configurations to manage DNS resolution behavior for your VPCs.
func route53resolver_ListResolverConfigs(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.ListResolverConfigsInput{}

	if len(_route53resolverMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _route53resolverMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_route53resolverNextToken) > 0 {
		input.NextToken = aws.String(_route53resolverNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListResolverConfigs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*route53resolver.ListResolverConfigsOutput
	p := route53resolver.NewListResolverConfigsPaginator(client, input)
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

// Lists the configurations for DNSSEC validation that are associated with the
// current Amazon Web Services account.
func route53resolver_ListResolverDnssecConfigs(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.ListResolverDnssecConfigsInput{}

	if len(_route53resolverFilters) > 0 {
		if err := assignInputField(input, "Filters", _route53resolverFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_route53resolverMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _route53resolverMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_route53resolverNextToken) > 0 {
		input.NextToken = aws.String(_route53resolverNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListResolverDnssecConfigs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*route53resolver.ListResolverDnssecConfigsOutput
	p := route53resolver.NewListResolverDnssecConfigsPaginator(client, input)
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

// Gets the IP addresses for a specified Resolver endpoint.
func route53resolver_ListResolverEndpointIpAddresses(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.ListResolverEndpointIpAddressesInput{
		// ResolverEndpointId: *string, // Required
	}

	if len(_route53resolverResolverEndpointId) > 0 {
		input.ResolverEndpointId = aws.String(_route53resolverResolverEndpointId)
	}
	if len(_route53resolverMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _route53resolverMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_route53resolverNextToken) > 0 {
		input.NextToken = aws.String(_route53resolverNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListResolverEndpointIpAddresses(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*route53resolver.ListResolverEndpointIpAddressesOutput
	p := route53resolver.NewListResolverEndpointIpAddressesPaginator(client, input)
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

// Lists all the Resolver endpoints that were created using the current Amazon Web
// Services account.
func route53resolver_ListResolverEndpoints(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.ListResolverEndpointsInput{}

	if len(_route53resolverFilters) > 0 {
		if err := assignInputField(input, "Filters", _route53resolverFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_route53resolverMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _route53resolverMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_route53resolverNextToken) > 0 {
		input.NextToken = aws.String(_route53resolverNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListResolverEndpoints(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*route53resolver.ListResolverEndpointsOutput
	p := route53resolver.NewListResolverEndpointsPaginator(client, input)
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

// Lists information about associations between Amazon VPCs and query logging
// configurations.
func route53resolver_ListResolverQueryLogConfigAssociations(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.ListResolverQueryLogConfigAssociationsInput{}

	if len(_route53resolverFilters) > 0 {
		if err := assignInputField(input, "Filters", _route53resolverFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_route53resolverMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _route53resolverMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_route53resolverNextToken) > 0 {
		input.NextToken = aws.String(_route53resolverNextToken)
	}
	if len(_route53resolverSortBy) > 0 {
		input.SortBy = aws.String(_route53resolverSortBy)
	}
	if len(_route53resolverSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _route53resolverSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListResolverQueryLogConfigAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*route53resolver.ListResolverQueryLogConfigAssociationsOutput
	p := route53resolver.NewListResolverQueryLogConfigAssociationsPaginator(client, input)
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

// Lists information about the specified query logging configurations. Each
// configuration defines where you want Resolver to save DNS query logs and
// specifies the VPCs that you want to log queries for.
func route53resolver_ListResolverQueryLogConfigs(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.ListResolverQueryLogConfigsInput{}

	if len(_route53resolverFilters) > 0 {
		if err := assignInputField(input, "Filters", _route53resolverFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_route53resolverMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _route53resolverMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_route53resolverNextToken) > 0 {
		input.NextToken = aws.String(_route53resolverNextToken)
	}
	if len(_route53resolverSortBy) > 0 {
		input.SortBy = aws.String(_route53resolverSortBy)
	}
	if len(_route53resolverSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _route53resolverSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListResolverQueryLogConfigs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*route53resolver.ListResolverQueryLogConfigsOutput
	p := route53resolver.NewListResolverQueryLogConfigsPaginator(client, input)
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

// Lists the associations that were created between Resolver rules and VPCs using
// the current Amazon Web Services account.
func route53resolver_ListResolverRuleAssociations(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.ListResolverRuleAssociationsInput{}

	if len(_route53resolverFilters) > 0 {
		if err := assignInputField(input, "Filters", _route53resolverFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_route53resolverMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _route53resolverMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_route53resolverNextToken) > 0 {
		input.NextToken = aws.String(_route53resolverNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListResolverRuleAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*route53resolver.ListResolverRuleAssociationsOutput
	p := route53resolver.NewListResolverRuleAssociationsPaginator(client, input)
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

// Lists the Resolver rules that were created using the current Amazon Web
// Services account.
func route53resolver_ListResolverRules(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.ListResolverRulesInput{}

	if len(_route53resolverFilters) > 0 {
		if err := assignInputField(input, "Filters", _route53resolverFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_route53resolverMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _route53resolverMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_route53resolverNextToken) > 0 {
		input.NextToken = aws.String(_route53resolverNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListResolverRules(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*route53resolver.ListResolverRulesOutput
	p := route53resolver.NewListResolverRulesPaginator(client, input)
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

// Lists the tags that you associated with the specified resource.
func route53resolver_ListTagsForResource(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_route53resolverResourceArn) > 0 {
		input.ResourceArn = aws.String(_route53resolverResourceArn)
	}
	if len(_route53resolverMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _route53resolverMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_route53resolverNextToken) > 0 {
		input.NextToken = aws.String(_route53resolverNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*route53resolver.ListTagsForResourceOutput
	p := route53resolver.NewListTagsForResourcePaginator(client, input)
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

// Attaches an Identity and Access Management (Amazon Web Services IAM) policy for
// sharing the rule group. You can use the policy to share the rule group using
// Resource Access Manager (RAM).
func route53resolver_PutFirewallRuleGroupPolicy(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.PutFirewallRuleGroupPolicyInput{
		// Arn: *string, // Required
		// FirewallRuleGroupPolicy: *string, // Required
	}

	if len(_route53resolverArn) > 0 {
		input.Arn = aws.String(_route53resolverArn)
	}
	if len(_route53resolverFirewallRuleGroupPolicy) > 0 {
		input.FirewallRuleGroupPolicy = aws.String(_route53resolverFirewallRuleGroupPolicy)
	}

	if resp, err := client.PutFirewallRuleGroupPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Specifies an Amazon Web Services account that you want to share a query logging
// configuration with, the query logging configuration that you want to share, and
// the operations that you want the account to be able to perform on the
// configuration.
func route53resolver_PutResolverQueryLogConfigPolicy(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.PutResolverQueryLogConfigPolicyInput{
		// Arn: *string, // Required
		// ResolverQueryLogConfigPolicy: *string, // Required
	}

	if len(_route53resolverArn) > 0 {
		input.Arn = aws.String(_route53resolverArn)
	}
	if len(_route53resolverResolverQueryLogConfigPolicy) > 0 {
		input.ResolverQueryLogConfigPolicy = aws.String(_route53resolverResolverQueryLogConfigPolicy)
	}

	if resp, err := client.PutResolverQueryLogConfigPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Specifies an Amazon Web Services rule that you want to share with another
// account, the account that you want to share the rule with, and the operations
// that you want the account to be able to perform on the rule.
func route53resolver_PutResolverRulePolicy(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.PutResolverRulePolicyInput{
		// Arn: *string, // Required
		// ResolverRulePolicy: *string, // Required
	}

	if len(_route53resolverArn) > 0 {
		input.Arn = aws.String(_route53resolverArn)
	}
	if len(_route53resolverResolverRulePolicy) > 0 {
		input.ResolverRulePolicy = aws.String(_route53resolverResolverRulePolicy)
	}

	if resp, err := client.PutResolverRulePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds one or more tags to a specified resource.
func route53resolver_TagResource(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_route53resolverResourceArn) > 0 {
		input.ResourceArn = aws.String(_route53resolverResourceArn)
	}
	if len(_route53resolverTags) > 0 {
		if err := assignInputField(input, "Tags", _route53resolverTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.TagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes one or more tags from a specified resource.
func route53resolver_UntagResource(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_route53resolverResourceArn) > 0 {
		input.ResourceArn = aws.String(_route53resolverResourceArn)
	}
	if len(_route53resolverTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _route53resolverTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration of the firewall behavior provided by DNS Firewall for
// a single VPC from Amazon Virtual Private Cloud (Amazon VPC).
func route53resolver_UpdateFirewallConfig(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.UpdateFirewallConfigInput{
		// FirewallFailOpen: types.FirewallFailOpenStatus, // Required
		// ResourceId: *string, // Required
	}

	if len(_route53resolverFirewallFailOpen) > 0 {
		if err := assignInputField(input, "FirewallFailOpen", _route53resolverFirewallFailOpen); err != nil {
			log.Errorf("invalid --firewall-fail-open: %s", err.Error())
			return
		}
	}
	if len(_route53resolverResourceId) > 0 {
		input.ResourceId = aws.String(_route53resolverResourceId)
	}

	if resp, err := client.UpdateFirewallConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the firewall domain list from an array of domain specifications.
func route53resolver_UpdateFirewallDomains(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.UpdateFirewallDomainsInput{
		// Domains: []string, // Required
		// FirewallDomainListId: *string, // Required
		// Operation: types.FirewallDomainUpdateOperation, // Required
	}

	if len(_route53resolverDomains) > 0 {
		input.Domains = append([]string(nil), _route53resolverDomains...)
	}
	if len(_route53resolverFirewallDomainListId) > 0 {
		input.FirewallDomainListId = aws.String(_route53resolverFirewallDomainListId)
	}
	if len(_route53resolverOperation) > 0 {
		if err := assignInputField(input, "Operation", _route53resolverOperation); err != nil {
			log.Errorf("invalid --operation: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateFirewallDomains(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified firewall rule.
func route53resolver_UpdateFirewallRule(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.UpdateFirewallRuleInput{
		// FirewallRuleGroupId: *string, // Required
	}

	if len(_route53resolverFirewallRuleGroupId) > 0 {
		input.FirewallRuleGroupId = aws.String(_route53resolverFirewallRuleGroupId)
	}
	if len(_route53resolverAction) > 0 {
		if err := assignInputField(input, "Action", _route53resolverAction); err != nil {
			log.Errorf("invalid --action: %s", err.Error())
			return
		}
	}
	if len(_route53resolverBlockOverrideDnsType) > 0 {
		if err := assignInputField(input, "BlockOverrideDnsType", _route53resolverBlockOverrideDnsType); err != nil {
			log.Errorf("invalid --block-override-dns-type: %s", err.Error())
			return
		}
	}
	if len(_route53resolverBlockOverrideDomain) > 0 {
		input.BlockOverrideDomain = aws.String(_route53resolverBlockOverrideDomain)
	}
	if len(_route53resolverBlockOverrideTtl) > 0 {
		if err := assignInputField(input, "BlockOverrideTtl", _route53resolverBlockOverrideTtl); err != nil {
			log.Errorf("invalid --block-override-ttl: %s", err.Error())
			return
		}
	}
	if len(_route53resolverBlockResponse) > 0 {
		if err := assignInputField(input, "BlockResponse", _route53resolverBlockResponse); err != nil {
			log.Errorf("invalid --block-response: %s", err.Error())
			return
		}
	}
	if len(_route53resolverConfidenceThreshold) > 0 {
		if err := assignInputField(input, "ConfidenceThreshold", _route53resolverConfidenceThreshold); err != nil {
			log.Errorf("invalid --confidence-threshold: %s", err.Error())
			return
		}
	}
	if len(_route53resolverDnsThreatProtection) > 0 {
		if err := assignInputField(input, "DnsThreatProtection", _route53resolverDnsThreatProtection); err != nil {
			log.Errorf("invalid --dns-threat-protection: %s", err.Error())
			return
		}
	}
	if len(_route53resolverFirewallDomainListId) > 0 {
		input.FirewallDomainListId = aws.String(_route53resolverFirewallDomainListId)
	}
	if len(_route53resolverFirewallDomainRedirectionAction) > 0 {
		if err := assignInputField(input, "FirewallDomainRedirectionAction", _route53resolverFirewallDomainRedirectionAction); err != nil {
			log.Errorf("invalid --firewall-domain-redirection-action: %s", err.Error())
			return
		}
	}
	if len(_route53resolverFirewallThreatProtectionId) > 0 {
		input.FirewallThreatProtectionId = aws.String(_route53resolverFirewallThreatProtectionId)
	}
	if len(_route53resolverName) > 0 {
		input.Name = aws.String(_route53resolverName)
	}
	if len(_route53resolverPriority) > 0 {
		if err := assignInputField(input, "Priority", _route53resolverPriority); err != nil {
			log.Errorf("invalid --priority: %s", err.Error())
			return
		}
	}
	if len(_route53resolverQtype) > 0 {
		input.Qtype = aws.String(_route53resolverQtype)
	}

	if resp, err := client.UpdateFirewallRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Changes the association of a FirewallRuleGroup with a VPC. The association enables DNS filtering
// for the VPC.
func route53resolver_UpdateFirewallRuleGroupAssociation(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.UpdateFirewallRuleGroupAssociationInput{
		// FirewallRuleGroupAssociationId: *string, // Required
	}

	if len(_route53resolverFirewallRuleGroupAssociationId) > 0 {
		input.FirewallRuleGroupAssociationId = aws.String(_route53resolverFirewallRuleGroupAssociationId)
	}
	if len(_route53resolverMutationProtection) > 0 {
		if err := assignInputField(input, "MutationProtection", _route53resolverMutationProtection); err != nil {
			log.Errorf("invalid --mutation-protection: %s", err.Error())
			return
		}
	}
	if len(_route53resolverName) > 0 {
		input.Name = aws.String(_route53resolverName)
	}
	if len(_route53resolverPriority) > 0 {
		if err := assignInputField(input, "Priority", _route53resolverPriority); err != nil {
			log.Errorf("invalid --priority: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateFirewallRuleGroupAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// You can use UpdateOutpostResolver to update the instance count, type, or name
// of a Resolver on an Outpost.
func route53resolver_UpdateOutpostResolver(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.UpdateOutpostResolverInput{
		// Id: *string, // Required
	}

	if len(_route53resolverId) > 0 {
		input.Id = aws.String(_route53resolverId)
	}
	if len(_route53resolverInstanceCount) > 0 {
		if err := assignInputField(input, "InstanceCount", _route53resolverInstanceCount); err != nil {
			log.Errorf("invalid --instance-count: %s", err.Error())
			return
		}
	}
	if len(_route53resolverName) > 0 {
		input.Name = aws.String(_route53resolverName)
	}
	if len(_route53resolverPreferredInstanceType) > 0 {
		input.PreferredInstanceType = aws.String(_route53resolverPreferredInstanceType)
	}

	if resp, err := client.UpdateOutpostResolver(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the behavior configuration of Route 53 Resolver behavior for a single
// VPC from Amazon Virtual Private Cloud.
func route53resolver_UpdateResolverConfig(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.UpdateResolverConfigInput{
		// AutodefinedReverseFlag: types.AutodefinedReverseFlag, // Required
		// ResourceId: *string, // Required
	}

	if len(_route53resolverAutodefinedReverseFlag) > 0 {
		if err := assignInputField(input, "AutodefinedReverseFlag", _route53resolverAutodefinedReverseFlag); err != nil {
			log.Errorf("invalid --autodefined-reverse-flag: %s", err.Error())
			return
		}
	}
	if len(_route53resolverResourceId) > 0 {
		input.ResourceId = aws.String(_route53resolverResourceId)
	}

	if resp, err := client.UpdateResolverConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing DNSSEC validation configuration. If there is no existing
// DNSSEC validation configuration, one is created.
func route53resolver_UpdateResolverDnssecConfig(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.UpdateResolverDnssecConfigInput{
		// ResourceId: *string, // Required
		// Validation: types.Validation, // Required
	}

	if len(_route53resolverResourceId) > 0 {
		input.ResourceId = aws.String(_route53resolverResourceId)
	}
	if len(_route53resolverValidation) > 0 {
		if err := assignInputField(input, "Validation", _route53resolverValidation); err != nil {
			log.Errorf("invalid --validation: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateResolverDnssecConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the name, or endpoint type for an inbound or an outbound Resolver
// endpoint. You can only update between IPV4 and DUALSTACK, IPV6 endpoint type
// can't be updated to other type.
func route53resolver_UpdateResolverEndpoint(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.UpdateResolverEndpointInput{
		// ResolverEndpointId: *string, // Required
	}

	if len(_route53resolverResolverEndpointId) > 0 {
		input.ResolverEndpointId = aws.String(_route53resolverResolverEndpointId)
	}
	if len(_route53resolverName) > 0 {
		input.Name = aws.String(_route53resolverName)
	}
	if len(_route53resolverProtocols) > 0 {
		if err := assignInputField(input, "Protocols", _route53resolverProtocols); err != nil {
			log.Errorf("invalid --protocols: %s", err.Error())
			return
		}
	}
	if len(_route53resolverResolverEndpointType) > 0 {
		if err := assignInputField(input, "ResolverEndpointType", _route53resolverResolverEndpointType); err != nil {
			log.Errorf("invalid --resolver-endpoint-type: %s", err.Error())
			return
		}
	}
	if len(_route53resolverRniEnhancedMetricsEnabled) > 0 {
		if err := assignInputField(input, "RniEnhancedMetricsEnabled", _route53resolverRniEnhancedMetricsEnabled); err != nil {
			log.Errorf("invalid --rni-enhanced-metrics-enabled: %s", err.Error())
			return
		}
	}
	if len(_route53resolverTargetNameServerMetricsEnabled) > 0 {
		if err := assignInputField(input, "TargetNameServerMetricsEnabled", _route53resolverTargetNameServerMetricsEnabled); err != nil {
			log.Errorf("invalid --target-name-server-metrics-enabled: %s", err.Error())
			return
		}
	}
	if len(_route53resolverUpdateIpAddresses) > 0 {
		if err := assignInputField(input, "UpdateIpAddresses", _route53resolverUpdateIpAddresses); err != nil {
			log.Errorf("invalid --update-ip-addresses: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateResolverEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates settings for a specified Resolver rule. ResolverRuleId is required, and
// all other parameters are optional. If you don't specify a parameter, it retains
// its current value.
func route53resolver_UpdateResolverRule(cfg aws.Config, client *route53resolver.Client) {
	input := &route53resolver.UpdateResolverRuleInput{
		// Config: *types.ResolverRuleConfig, // Required
		// ResolverRuleId: *string, // Required
	}

	if len(_route53resolverConfig) > 0 {
		if err := assignInputField(input, "Config", _route53resolverConfig); err != nil {
			log.Errorf("invalid --config: %s", err.Error())
			return
		}
	}
	if len(_route53resolverResolverRuleId) > 0 {
		input.ResolverRuleId = aws.String(_route53resolverResolverRuleId)
	}

	if resp, err := client.UpdateResolverRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_route53resolverCmd)
	_route53resolverCmd.Flags().SortFlags = false

	_route53resolverCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_route53resolverCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_route53resolverCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_route53resolverCmd.Flags().StringVarP(&_route53resolverAction, "action", "", "", "Action")
	_route53resolverCmd.Flags().StringVarP(&_route53resolverArn, "arn", "", "", "ARN")
	_route53resolverCmd.Flags().StringVarP(&_route53resolverAutodefinedReverseFlag, "autodefined-reverse-flag", "", "", "Autodefined Reverse Flag")
	_route53resolverCmd.Flags().StringVarP(&_route53resolverBlockOverrideDnsType, "block-override-dns-type", "", "", "Block Override DNS Type")
	_route53resolverCmd.Flags().StringVarP(&_route53resolverBlockOverrideDomain, "block-override-domain", "", "", "Block Override Domain")
	_route53resolverCmd.Flags().StringVarP(&_route53resolverBlockOverrideTtl, "block-override-ttl", "", "", "Block Override TTL")
	_route53resolverCmd.Flags().StringVarP(&_route53resolverBlockResponse, "block-response", "", "", "Block Response")
	_route53resolverCmd.Flags().StringVarP(&_route53resolverConfidenceThreshold, "confidence-threshold", "", "", "Confidence Threshold")
	_route53resolverCmd.Flags().StringVarP(&_route53resolverConfig, "config", "", "", "Config")
	_route53resolverCmd.Flags().StringVarP(&_route53resolverCreatorRequestId, "creator-request-id", "", "", "Creator Request ID")
	_route53resolverCmd.Flags().StringVarP(&_route53resolverDelegationRecord, "delegation-record", "", "", "Delegation Record")
	_route53resolverCmd.Flags().StringVarP(&_route53resolverDestinationArn, "destination-arn", "", "", "Destination ARN")
	_route53resolverCmd.Flags().StringVarP(&_route53resolverDirection, "direction", "", "", "Direction")
	_route53resolverCmd.Flags().StringVarP(&_route53resolverDnsThreatProtection, "dns-threat-protection", "", "", "DNS Threat Protection")
	_route53resolverCmd.Flags().StringVarP(&_route53resolverDomainFileUrl, "domain-file-url", "", "", "Domain File URL")
	_route53resolverCmd.Flags().StringVarP(&_route53resolverDomainName, "domain-name", "", "", "Domain Name")
	_route53resolverCmd.Flags().StringSliceVarP(&_route53resolverDomains, "domains", "", nil, "Domains")
	_route53resolverCmd.Flags().StringVarP(&_route53resolverFilters, "filters", "", "", "Filters")
	_route53resolverCmd.Flags().StringVarP(&_route53resolverFirewallDomainListId, "firewall-domain-list-id", "", "", "Firewall Domain List ID")
	_route53resolverCmd.Flags().StringVarP(&_route53resolverFirewallDomainRedirectionAction, "firewall-domain-redirection-action", "", "", "Firewall Domain Redirection Action")
	_route53resolverCmd.Flags().StringVarP(&_route53resolverFirewallFailOpen, "firewall-fail-open", "", "", "Firewall Fail Open")
	_route53resolverCmd.Flags().StringVarP(&_route53resolverFirewallRuleGroupAssociationId, "firewall-rule-group-association-id", "", "", "Firewall Rule Group Association ID")
	_route53resolverCmd.Flags().StringVarP(&_route53resolverFirewallRuleGroupId, "firewall-rule-group-id", "", "", "Firewall Rule Group ID")
	_route53resolverCmd.Flags().StringVarP(&_route53resolverFirewallRuleGroupPolicy, "firewall-rule-group-policy", "", "", "Firewall Rule Group Policy")
	_route53resolverCmd.Flags().StringVarP(&_route53resolverFirewallThreatProtectionId, "firewall-threat-protection-id", "", "", "Firewall Threat Protection ID")
	_route53resolverCmd.Flags().StringVarP(&_route53resolverId, "id", "", "", "ID")
	_route53resolverCmd.Flags().StringVarP(&_route53resolverInstanceCount, "instance-count", "", "", "Instance Count")
	_route53resolverCmd.Flags().StringVarP(&_route53resolverIpAddress, "ip-address", "", "", "IP Address")
	_route53resolverCmd.Flags().StringVarP(&_route53resolverIpAddresses, "ip-addresses", "", "", "IP Addresses")
	_route53resolverCmd.Flags().StringVarP(&_route53resolverMaxResults, "max-results", "", "", "Max Results")
	_route53resolverCmd.Flags().StringVarP(&_route53resolverMutationProtection, "mutation-protection", "", "", "Mutation Protection")
	_route53resolverCmd.Flags().StringVarP(&_route53resolverName, "name", "", "", "Name")
	_route53resolverCmd.Flags().StringVarP(&_route53resolverNextToken, "next-token", "", "", "Next Token")
	_route53resolverCmd.Flags().StringVarP(&_route53resolverOperation, "operation", "", "", "Operation")
	_route53resolverCmd.Flags().StringVarP(&_route53resolverOutpostArn, "outpost-arn", "", "", "Outpost ARN")
	_route53resolverCmd.Flags().StringVarP(&_route53resolverPreferredInstanceType, "preferred-instance-type", "", "", "Preferred Instance Type")
	_route53resolverCmd.Flags().StringVarP(&_route53resolverPriority, "priority", "", "", "Priority")
	_route53resolverCmd.Flags().StringVarP(&_route53resolverProtocols, "protocols", "", "", "Protocols")
	_route53resolverCmd.Flags().StringVarP(&_route53resolverQtype, "qtype", "", "", "Qtype")
	_route53resolverCmd.Flags().StringVarP(&_route53resolverResolverEndpointId, "resolver-endpoint-id", "", "", "Resolver Endpoint ID")
	_route53resolverCmd.Flags().StringVarP(&_route53resolverResolverEndpointType, "resolver-endpoint-type", "", "", "Resolver Endpoint Type")
	_route53resolverCmd.Flags().StringVarP(&_route53resolverResolverQueryLogConfigAssociationId, "resolver-query-log-config-association-id", "", "", "Resolver Query Log Config Association ID")
	_route53resolverCmd.Flags().StringVarP(&_route53resolverResolverQueryLogConfigId, "resolver-query-log-config-id", "", "", "Resolver Query Log Config ID")
	_route53resolverCmd.Flags().StringVarP(&_route53resolverResolverQueryLogConfigPolicy, "resolver-query-log-config-policy", "", "", "Resolver Query Log Config Policy")
	_route53resolverCmd.Flags().StringVarP(&_route53resolverResolverRuleAssociationId, "resolver-rule-association-id", "", "", "Resolver Rule Association ID")
	_route53resolverCmd.Flags().StringVarP(&_route53resolverResolverRuleId, "resolver-rule-id", "", "", "Resolver Rule ID")
	_route53resolverCmd.Flags().StringVarP(&_route53resolverResolverRulePolicy, "resolver-rule-policy", "", "", "Resolver Rule Policy")
	_route53resolverCmd.Flags().StringVarP(&_route53resolverResourceArn, "resource-arn", "", "", "Resource ARN")
	_route53resolverCmd.Flags().StringVarP(&_route53resolverResourceId, "resource-id", "", "", "Resource ID")
	_route53resolverCmd.Flags().StringVarP(&_route53resolverRniEnhancedMetricsEnabled, "rni-enhanced-metrics-enabled", "", "", "Rni Enhanced Metrics Enabled")
	_route53resolverCmd.Flags().StringVarP(&_route53resolverRuleType, "rule-type", "", "", "Rule Type")
	_route53resolverCmd.Flags().StringSliceVarP(&_route53resolverSecurityGroupIds, "security-group-ids", "", nil, "Security Group Ids")
	_route53resolverCmd.Flags().StringVarP(&_route53resolverSortBy, "sort-by", "", "", "Sort By")
	_route53resolverCmd.Flags().StringVarP(&_route53resolverSortOrder, "sort-order", "", "", "Sort Order")
	_route53resolverCmd.Flags().StringVarP(&_route53resolverStatus, "status", "", "", "Status")
	_route53resolverCmd.Flags().StringSliceVarP(&_route53resolverTagKeys, "tag-keys", "", nil, "Tag Keys")
	_route53resolverCmd.Flags().StringVarP(&_route53resolverTags, "tags", "", "", "Tags")
	_route53resolverCmd.Flags().StringVarP(&_route53resolverTargetIps, "target-ips", "", "", "Target Ips")
	_route53resolverCmd.Flags().StringVarP(&_route53resolverTargetNameServerMetricsEnabled, "target-name-server-metrics-enabled", "", "", "Target Name Server Metrics Enabled")
	_route53resolverCmd.Flags().StringVarP(&_route53resolverUpdateIpAddresses, "update-ip-addresses", "", "", "Update IP Addresses")
	_route53resolverCmd.Flags().StringVarP(&_route53resolverValidation, "validation", "", "", "Validation")
	_route53resolverCmd.Flags().StringVarP(&_route53resolverVpcId, "vpc-id", "", "", "VPC ID")
	_route53resolverCmd.Flags().StringVarP(&_route53resolverVPCId, "vpcid", "", "", "Vpcid")

	_route53resolverCmd.Flags().BoolVarP(&_route53resolverAssociateFirewallRuleGroup, "associate-firewall-rule-group", "", false, "Associate Firewall Rule Group")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverAssociateResolverEndpointIpAddress, "associate-resolver-endpoint-ip-address", "", false, "Associate Resolver Endpoint IP Address")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverAssociateResolverQueryLogConfig, "associate-resolver-query-log-config", "", false, "Associate Resolver Query Log Config")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverAssociateResolverRule, "associate-resolver-rule", "", false, "Associate Resolver Rule")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverCreateFirewallDomainList, "create-firewall-domain-list", "", false, "Create Firewall Domain List")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverCreateFirewallRule, "create-firewall-rule", "", false, "Create Firewall Rule")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverCreateFirewallRuleGroup, "create-firewall-rule-group", "", false, "Create Firewall Rule Group")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverCreateOutpostResolver, "create-outpost-resolver", "", false, "Create Outpost Resolver")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverCreateResolverEndpoint, "create-resolver-endpoint", "", false, "Create Resolver Endpoint")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverCreateResolverQueryLogConfig, "create-resolver-query-log-config", "", false, "Create Resolver Query Log Config")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverCreateResolverRule, "create-resolver-rule", "", false, "Create Resolver Rule")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverDeleteFirewallDomainList, "delete-firewall-domain-list", "", false, "Delete Firewall Domain List")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverDeleteFirewallRule, "delete-firewall-rule", "", false, "Delete Firewall Rule")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverDeleteFirewallRuleGroup, "delete-firewall-rule-group", "", false, "Delete Firewall Rule Group")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverDeleteOutpostResolver, "delete-outpost-resolver", "", false, "Delete Outpost Resolver")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverDeleteResolverEndpoint, "delete-resolver-endpoint", "", false, "Delete Resolver Endpoint")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverDeleteResolverQueryLogConfig, "delete-resolver-query-log-config", "", false, "Delete Resolver Query Log Config")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverDeleteResolverRule, "delete-resolver-rule", "", false, "Delete Resolver Rule")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverDisassociateFirewallRuleGroup, "disassociate-firewall-rule-group", "", false, "Disassociate Firewall Rule Group")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverDisassociateResolverEndpointIpAddress, "disassociate-resolver-endpoint-ip-address", "", false, "Disassociate Resolver Endpoint IP Address")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverDisassociateResolverQueryLogConfig, "disassociate-resolver-query-log-config", "", false, "Disassociate Resolver Query Log Config")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverDisassociateResolverRule, "disassociate-resolver-rule", "", false, "Disassociate Resolver Rule")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverGetFirewallConfig, "get-firewall-config", "", false, "Get Firewall Config")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverGetFirewallDomainList, "get-firewall-domain-list", "", false, "Get Firewall Domain List")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverGetFirewallRuleGroup, "get-firewall-rule-group", "", false, "Get Firewall Rule Group")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverGetFirewallRuleGroupAssociation, "get-firewall-rule-group-association", "", false, "Get Firewall Rule Group Association")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverGetFirewallRuleGroupPolicy, "get-firewall-rule-group-policy", "", false, "Get Firewall Rule Group Policy")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverGetOutpostResolver, "get-outpost-resolver", "", false, "Get Outpost Resolver")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverGetResolverConfig, "get-resolver-config", "", false, "Get Resolver Config")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverGetResolverDnssecConfig, "get-resolver-dnssec-config", "", false, "Get Resolver Dnssec Config")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverGetResolverEndpoint, "get-resolver-endpoint", "", false, "Get Resolver Endpoint")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverGetResolverQueryLogConfig, "get-resolver-query-log-config", "", false, "Get Resolver Query Log Config")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverGetResolverQueryLogConfigAssociation, "get-resolver-query-log-config-association", "", false, "Get Resolver Query Log Config Association")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverGetResolverQueryLogConfigPolicy, "get-resolver-query-log-config-policy", "", false, "Get Resolver Query Log Config Policy")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverGetResolverRule, "get-resolver-rule", "", false, "Get Resolver Rule")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverGetResolverRuleAssociation, "get-resolver-rule-association", "", false, "Get Resolver Rule Association")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverGetResolverRulePolicy, "get-resolver-rule-policy", "", false, "Get Resolver Rule Policy")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverImportFirewallDomains, "import-firewall-domains", "", false, "Import Firewall Domains")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverListFirewallConfigs, "list-firewall-configs", "", false, "List Firewall Configs")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverListFirewallDomainLists, "list-firewall-domain-lists", "", false, "List Firewall Domain Lists")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverListFirewallDomains, "list-firewall-domains", "", false, "List Firewall Domains")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverListFirewallRuleGroupAssociations, "list-firewall-rule-group-associations", "", false, "List Firewall Rule Group Associations")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverListFirewallRuleGroups, "list-firewall-rule-groups", "", false, "List Firewall Rule Groups")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverListFirewallRules, "list-firewall-rules", "", false, "List Firewall Rules")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverListOutpostResolvers, "list-outpost-resolvers", "", false, "List Outpost Resolvers")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverListResolverConfigs, "list-resolver-configs", "", false, "List Resolver Configs")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverListResolverDnssecConfigs, "list-resolver-dnssec-configs", "", false, "List Resolver Dnssec Configs")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverListResolverEndpointIpAddresses, "list-resolver-endpoint-ip-addresses", "", false, "List Resolver Endpoint IP Addresses")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverListResolverEndpoints, "list-resolver-endpoints", "", false, "List Resolver Endpoints")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverListResolverQueryLogConfigAssociations, "list-resolver-query-log-config-associations", "", false, "List Resolver Query Log Config Associations")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverListResolverQueryLogConfigs, "list-resolver-query-log-configs", "", false, "List Resolver Query Log Configs")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverListResolverRuleAssociations, "list-resolver-rule-associations", "", false, "List Resolver Rule Associations")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverListResolverRules, "list-resolver-rules", "", false, "List Resolver Rules")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverPutFirewallRuleGroupPolicy, "put-firewall-rule-group-policy", "", false, "Put Firewall Rule Group Policy")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverPutResolverQueryLogConfigPolicy, "put-resolver-query-log-config-policy", "", false, "Put Resolver Query Log Config Policy")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverPutResolverRulePolicy, "put-resolver-rule-policy", "", false, "Put Resolver Rule Policy")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverTagResource, "tag-resource", "", false, "Tag Resource")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverUntagResource, "untag-resource", "", false, "Untag Resource")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverUpdateFirewallConfig, "update-firewall-config", "", false, "Update Firewall Config")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverUpdateFirewallDomains, "update-firewall-domains", "", false, "Update Firewall Domains")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverUpdateFirewallRule, "update-firewall-rule", "", false, "Update Firewall Rule")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverUpdateFirewallRuleGroupAssociation, "update-firewall-rule-group-association", "", false, "Update Firewall Rule Group Association")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverUpdateOutpostResolver, "update-outpost-resolver", "", false, "Update Outpost Resolver")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverUpdateResolverConfig, "update-resolver-config", "", false, "Update Resolver Config")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverUpdateResolverDnssecConfig, "update-resolver-dnssec-config", "", false, "Update Resolver Dnssec Config")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverUpdateResolverEndpoint, "update-resolver-endpoint", "", false, "Update Resolver Endpoint")
	_route53resolverCmd.Flags().BoolVarP(&_route53resolverUpdateResolverRule, "update-resolver-rule", "", false, "Update Resolver Rule")

}
