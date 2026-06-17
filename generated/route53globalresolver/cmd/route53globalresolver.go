package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53globalresolver"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// route53globalresolverCmd represents the route53globalresolver command
var _route53globalresolverCmd = &cobra.Command{
	Use:   "route53globalresolver",
	Short: "AWS route53globalresolver CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := route53globalresolver.NewFromConfig(cfg)
		if _route53globalresolverAssociateHostedZone {
			route53globalresolver_AssociateHostedZone(cfg, client)
			return
		}
		if _route53globalresolverBatchCreateFirewallRule {
			route53globalresolver_BatchCreateFirewallRule(cfg, client)
			return
		}
		if _route53globalresolverBatchDeleteFirewallRule {
			route53globalresolver_BatchDeleteFirewallRule(cfg, client)
			return
		}
		if _route53globalresolverBatchUpdateFirewallRule {
			route53globalresolver_BatchUpdateFirewallRule(cfg, client)
			return
		}
		if _route53globalresolverCreateAccessSource {
			route53globalresolver_CreateAccessSource(cfg, client)
			return
		}
		if _route53globalresolverCreateAccessToken {
			route53globalresolver_CreateAccessToken(cfg, client)
			return
		}
		if _route53globalresolverCreateDNSView {
			route53globalresolver_CreateDNSView(cfg, client)
			return
		}
		if _route53globalresolverCreateFirewallDomainList {
			route53globalresolver_CreateFirewallDomainList(cfg, client)
			return
		}
		if _route53globalresolverCreateFirewallRule {
			route53globalresolver_CreateFirewallRule(cfg, client)
			return
		}
		if _route53globalresolverCreateGlobalResolver {
			route53globalresolver_CreateGlobalResolver(cfg, client)
			return
		}
		if _route53globalresolverDeleteAccessSource {
			route53globalresolver_DeleteAccessSource(cfg, client)
			return
		}
		if _route53globalresolverDeleteAccessToken {
			route53globalresolver_DeleteAccessToken(cfg, client)
			return
		}
		if _route53globalresolverDeleteDNSView {
			route53globalresolver_DeleteDNSView(cfg, client)
			return
		}
		if _route53globalresolverDeleteFirewallDomainList {
			route53globalresolver_DeleteFirewallDomainList(cfg, client)
			return
		}
		if _route53globalresolverDeleteFirewallRule {
			route53globalresolver_DeleteFirewallRule(cfg, client)
			return
		}
		if _route53globalresolverDeleteGlobalResolver {
			route53globalresolver_DeleteGlobalResolver(cfg, client)
			return
		}
		if _route53globalresolverDisableDNSView {
			route53globalresolver_DisableDNSView(cfg, client)
			return
		}
		if _route53globalresolverDisassociateHostedZone {
			route53globalresolver_DisassociateHostedZone(cfg, client)
			return
		}
		if _route53globalresolverEnableDNSView {
			route53globalresolver_EnableDNSView(cfg, client)
			return
		}
		if _route53globalresolverGetAccessSource {
			route53globalresolver_GetAccessSource(cfg, client)
			return
		}
		if _route53globalresolverGetAccessToken {
			route53globalresolver_GetAccessToken(cfg, client)
			return
		}
		if _route53globalresolverGetDNSView {
			route53globalresolver_GetDNSView(cfg, client)
			return
		}
		if _route53globalresolverGetFirewallDomainList {
			route53globalresolver_GetFirewallDomainList(cfg, client)
			return
		}
		if _route53globalresolverGetFirewallRule {
			route53globalresolver_GetFirewallRule(cfg, client)
			return
		}
		if _route53globalresolverGetGlobalResolver {
			route53globalresolver_GetGlobalResolver(cfg, client)
			return
		}
		if _route53globalresolverGetHostedZoneAssociation {
			route53globalresolver_GetHostedZoneAssociation(cfg, client)
			return
		}
		if _route53globalresolverGetManagedFirewallDomainList {
			route53globalresolver_GetManagedFirewallDomainList(cfg, client)
			return
		}
		if _route53globalresolverImportFirewallDomains {
			route53globalresolver_ImportFirewallDomains(cfg, client)
			return
		}
		if _route53globalresolverListAccessSources {
			route53globalresolver_ListAccessSources(cfg, client)
			return
		}
		if _route53globalresolverListAccessTokens {
			route53globalresolver_ListAccessTokens(cfg, client)
			return
		}
		if _route53globalresolverListDNSViews {
			route53globalresolver_ListDNSViews(cfg, client)
			return
		}
		if _route53globalresolverListFirewallDomainLists {
			route53globalresolver_ListFirewallDomainLists(cfg, client)
			return
		}
		if _route53globalresolverListFirewallDomains {
			route53globalresolver_ListFirewallDomains(cfg, client)
			return
		}
		if _route53globalresolverListFirewallRules {
			route53globalresolver_ListFirewallRules(cfg, client)
			return
		}
		if _route53globalresolverListGlobalResolvers {
			route53globalresolver_ListGlobalResolvers(cfg, client)
			return
		}
		if _route53globalresolverListHostedZoneAssociations {
			route53globalresolver_ListHostedZoneAssociations(cfg, client)
			return
		}
		if _route53globalresolverListManagedFirewallDomainLists {
			route53globalresolver_ListManagedFirewallDomainLists(cfg, client)
			return
		}
		if _route53globalresolverListTagsForResource {
			route53globalresolver_ListTagsForResource(cfg, client)
			return
		}
		if _route53globalresolverTagResource {
			route53globalresolver_TagResource(cfg, client)
			return
		}
		if _route53globalresolverUntagResource {
			route53globalresolver_UntagResource(cfg, client)
			return
		}
		if _route53globalresolverUpdateAccessSource {
			route53globalresolver_UpdateAccessSource(cfg, client)
			return
		}
		if _route53globalresolverUpdateAccessToken {
			route53globalresolver_UpdateAccessToken(cfg, client)
			return
		}
		if _route53globalresolverUpdateDNSView {
			route53globalresolver_UpdateDNSView(cfg, client)
			return
		}
		if _route53globalresolverUpdateFirewallDomains {
			route53globalresolver_UpdateFirewallDomains(cfg, client)
			return
		}
		if _route53globalresolverUpdateFirewallRule {
			route53globalresolver_UpdateFirewallRule(cfg, client)
			return
		}
		if _route53globalresolverUpdateGlobalResolver {
			route53globalresolver_UpdateGlobalResolver(cfg, client)
			return
		}
		if _route53globalresolverUpdateHostedZoneAssociation {
			route53globalresolver_UpdateHostedZoneAssociation(cfg, client)
			return
		}

	},
}

var (
	_route53globalresolverAssociateHostedZone            bool
	_route53globalresolverBatchCreateFirewallRule        bool
	_route53globalresolverBatchDeleteFirewallRule        bool
	_route53globalresolverBatchUpdateFirewallRule        bool
	_route53globalresolverCreateAccessSource             bool
	_route53globalresolverCreateAccessToken              bool
	_route53globalresolverCreateDNSView                  bool
	_route53globalresolverCreateFirewallDomainList       bool
	_route53globalresolverCreateFirewallRule             bool
	_route53globalresolverCreateGlobalResolver           bool
	_route53globalresolverDeleteAccessSource             bool
	_route53globalresolverDeleteAccessToken              bool
	_route53globalresolverDeleteDNSView                  bool
	_route53globalresolverDeleteFirewallDomainList       bool
	_route53globalresolverDeleteFirewallRule             bool
	_route53globalresolverDeleteGlobalResolver           bool
	_route53globalresolverDisableDNSView                 bool
	_route53globalresolverDisassociateHostedZone         bool
	_route53globalresolverEnableDNSView                  bool
	_route53globalresolverGetAccessSource                bool
	_route53globalresolverGetAccessToken                 bool
	_route53globalresolverGetDNSView                     bool
	_route53globalresolverGetFirewallDomainList          bool
	_route53globalresolverGetFirewallRule                bool
	_route53globalresolverGetGlobalResolver              bool
	_route53globalresolverGetHostedZoneAssociation       bool
	_route53globalresolverGetManagedFirewallDomainList   bool
	_route53globalresolverImportFirewallDomains          bool
	_route53globalresolverListAccessSources              bool
	_route53globalresolverListAccessTokens               bool
	_route53globalresolverListDNSViews                   bool
	_route53globalresolverListFirewallDomainLists        bool
	_route53globalresolverListFirewallDomains            bool
	_route53globalresolverListFirewallRules              bool
	_route53globalresolverListGlobalResolvers            bool
	_route53globalresolverListHostedZoneAssociations     bool
	_route53globalresolverListManagedFirewallDomainLists bool
	_route53globalresolverListTagsForResource            bool
	_route53globalresolverTagResource                    bool
	_route53globalresolverUntagResource                  bool
	_route53globalresolverUpdateAccessSource             bool
	_route53globalresolverUpdateAccessToken              bool
	_route53globalresolverUpdateDNSView                  bool
	_route53globalresolverUpdateFirewallDomains          bool
	_route53globalresolverUpdateFirewallRule             bool
	_route53globalresolverUpdateGlobalResolver           bool
	_route53globalresolverUpdateHostedZoneAssociation    bool

	_route53globalresolverAccessSourceId                string
	_route53globalresolverAccessTokenId                 string
	_route53globalresolverAction                        string
	_route53globalresolverBlockOverrideDnsType          string
	_route53globalresolverBlockOverrideDomain           string
	_route53globalresolverBlockOverrideTtl              string
	_route53globalresolverBlockResponse                 string
	_route53globalresolverCidr                          string
	_route53globalresolverClientToken                   string
	_route53globalresolverConfidenceThreshold           string
	_route53globalresolverDescription                   string
	_route53globalresolverDnsAdvancedProtection         string
	_route53globalresolverDnsViewId                     string
	_route53globalresolverDnssecValidation              string
	_route53globalresolverDomainFileUrl                 string
	_route53globalresolverDomains                       []string
	_route53globalresolverEdnsClientSubnet              string
	_route53globalresolverExpiresAt                     string
	_route53globalresolverFilters                       string
	_route53globalresolverFirewallDomainListId          string
	_route53globalresolverFirewallRuleId                string
	_route53globalresolverFirewallRules                 string
	_route53globalresolverFirewallRulesFailOpen         string
	_route53globalresolverGlobalResolverId              string
	_route53globalresolverHostedZoneAssociationId       string
	_route53globalresolverHostedZoneId                  string
	_route53globalresolverIpAddressType                 string
	_route53globalresolverManagedFirewallDomainListId   string
	_route53globalresolverManagedFirewallDomainListType string
	_route53globalresolverMaxResults                    string
	_route53globalresolverName                          string
	_route53globalresolverNextToken                     string
	_route53globalresolverObservabilityRegion           string
	_route53globalresolverOperation                     string
	_route53globalresolverPriority                      string
	_route53globalresolverProtocol                      string
	_route53globalresolverQType                         string
	_route53globalresolverRegions                       []string
	_route53globalresolverResourceArn                   string
	_route53globalresolverTagKeys                       []string
	_route53globalresolverTags                          string
)

// Associates a Route 53 private hosted zone with a Route 53 Global Resolver
// resource. This allows the resolver to resolve DNS queries for the private hosted
// zone from anywhere globally.
func route53globalresolver_AssociateHostedZone(cfg aws.Config, client *route53globalresolver.Client) {
	input := &route53globalresolver.AssociateHostedZoneInput{
		// HostedZoneId: *string, // Required
		// Name: *string, // Required
		// ResourceArn: *string, // Required
	}

	if len(_route53globalresolverHostedZoneId) > 0 {
		input.HostedZoneId = aws.String(_route53globalresolverHostedZoneId)
	}
	if len(_route53globalresolverName) > 0 {
		input.Name = aws.String(_route53globalresolverName)
	}
	if len(_route53globalresolverResourceArn) > 0 {
		input.ResourceArn = aws.String(_route53globalresolverResourceArn)
	}

	if resp, err := client.AssociateHostedZone(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates multiple DNS firewall rules in a single operation. This is more
// efficient than creating rules individually when you need to set up multiple
// rules at once.
func route53globalresolver_BatchCreateFirewallRule(cfg aws.Config, client *route53globalresolver.Client) {
	input := &route53globalresolver.BatchCreateFirewallRuleInput{
		// FirewallRules: []types.BatchCreateFirewallRuleInputItem, // Required
	}

	if len(_route53globalresolverFirewallRules) > 0 {
		if err := assignInputField(input, "FirewallRules", _route53globalresolverFirewallRules); err != nil {
			log.Errorf("invalid --firewall-rules: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchCreateFirewallRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes multiple DNS firewall rules in a single operation. This is more
// efficient than deleting rules individually.
func route53globalresolver_BatchDeleteFirewallRule(cfg aws.Config, client *route53globalresolver.Client) {
	input := &route53globalresolver.BatchDeleteFirewallRuleInput{
		// FirewallRules: []types.BatchDeleteFirewallRuleInputItem, // Required
	}

	if len(_route53globalresolverFirewallRules) > 0 {
		if err := assignInputField(input, "FirewallRules", _route53globalresolverFirewallRules); err != nil {
			log.Errorf("invalid --firewall-rules: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchDeleteFirewallRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates multiple DNS firewall rules in a single operation. This is more
// efficient than updating rules individually.
func route53globalresolver_BatchUpdateFirewallRule(cfg aws.Config, client *route53globalresolver.Client) {
	input := &route53globalresolver.BatchUpdateFirewallRuleInput{
		// FirewallRules: []types.BatchUpdateFirewallRuleInputItem, // Required
	}

	if len(_route53globalresolverFirewallRules) > 0 {
		if err := assignInputField(input, "FirewallRules", _route53globalresolverFirewallRules); err != nil {
			log.Errorf("invalid --firewall-rules: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchUpdateFirewallRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an access source for a DNS view. Access sources define IP addresses or
// CIDR ranges that are allowed to send DNS queries to the Route 53 Global
// Resolver, along with the permitted DNS protocols.
func route53globalresolver_CreateAccessSource(cfg aws.Config, client *route53globalresolver.Client) {
	input := &route53globalresolver.CreateAccessSourceInput{
		// Cidr: *string, // Required
		// DnsViewId: *string, // Required
		// Protocol: types.DnsProtocol, // Required
	}

	if len(_route53globalresolverCidr) > 0 {
		input.Cidr = aws.String(_route53globalresolverCidr)
	}
	if len(_route53globalresolverDnsViewId) > 0 {
		input.DnsViewId = aws.String(_route53globalresolverDnsViewId)
	}
	if len(_route53globalresolverProtocol) > 0 {
		if err := assignInputField(input, "Protocol", _route53globalresolverProtocol); err != nil {
			log.Errorf("invalid --protocol: %s", err.Error())
			return
		}
	}
	if len(_route53globalresolverClientToken) > 0 {
		input.ClientToken = aws.String(_route53globalresolverClientToken)
	}
	if len(_route53globalresolverIpAddressType) > 0 {
		if err := assignInputField(input, "IpAddressType", _route53globalresolverIpAddressType); err != nil {
			log.Errorf("invalid --ip-address-type: %s", err.Error())
			return
		}
	}
	if len(_route53globalresolverName) > 0 {
		input.Name = aws.String(_route53globalresolverName)
	}
	if len(_route53globalresolverTags) > 0 {
		if err := assignInputField(input, "Tags", _route53globalresolverTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAccessSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an access token for a DNS view. Access tokens provide token-based
// authentication for DNS-over-HTTPS (DoH) and DNS-over-TLS (DoT) connections to
// the Route 53 Global Resolver.
func route53globalresolver_CreateAccessToken(cfg aws.Config, client *route53globalresolver.Client) {
	input := &route53globalresolver.CreateAccessTokenInput{
		// DnsViewId: *string, // Required
	}

	if len(_route53globalresolverDnsViewId) > 0 {
		input.DnsViewId = aws.String(_route53globalresolverDnsViewId)
	}
	if len(_route53globalresolverClientToken) > 0 {
		input.ClientToken = aws.String(_route53globalresolverClientToken)
	}
	if len(_route53globalresolverExpiresAt) > 0 {
		if err := assignInputField(input, "ExpiresAt", _route53globalresolverExpiresAt); err != nil {
			log.Errorf("invalid --expires-at: %s", err.Error())
			return
		}
	}
	if len(_route53globalresolverName) > 0 {
		input.Name = aws.String(_route53globalresolverName)
	}
	if len(_route53globalresolverTags) > 0 {
		if err := assignInputField(input, "Tags", _route53globalresolverTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAccessToken(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a DNS view within a Route 53 Global Resolver. A DNS view models end
// users, user groups, networks, and devices, and serves as a parent resource that
// holds configurations controlling access, authorization, DNS firewall rules, and
// forwarding rules.
func route53globalresolver_CreateDNSView(cfg aws.Config, client *route53globalresolver.Client) {
	input := &route53globalresolver.CreateDNSViewInput{
		// GlobalResolverId: *string, // Required
		// Name: *string, // Required
	}

	if len(_route53globalresolverGlobalResolverId) > 0 {
		input.GlobalResolverId = aws.String(_route53globalresolverGlobalResolverId)
	}
	if len(_route53globalresolverName) > 0 {
		input.Name = aws.String(_route53globalresolverName)
	}
	if len(_route53globalresolverClientToken) > 0 {
		input.ClientToken = aws.String(_route53globalresolverClientToken)
	}
	if len(_route53globalresolverDescription) > 0 {
		input.Description = aws.String(_route53globalresolverDescription)
	}
	if len(_route53globalresolverDnssecValidation) > 0 {
		if err := assignInputField(input, "DnssecValidation", _route53globalresolverDnssecValidation); err != nil {
			log.Errorf("invalid --dnssec-validation: %s", err.Error())
			return
		}
	}
	if len(_route53globalresolverEdnsClientSubnet) > 0 {
		if err := assignInputField(input, "EdnsClientSubnet", _route53globalresolverEdnsClientSubnet); err != nil {
			log.Errorf("invalid --edns-client-subnet: %s", err.Error())
			return
		}
	}
	if len(_route53globalresolverFirewallRulesFailOpen) > 0 {
		if err := assignInputField(input, "FirewallRulesFailOpen", _route53globalresolverFirewallRulesFailOpen); err != nil {
			log.Errorf("invalid --firewall-rules-fail-open: %s", err.Error())
			return
		}
	}
	if len(_route53globalresolverTags) > 0 {
		if err := assignInputField(input, "Tags", _route53globalresolverTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDNSView(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a firewall domain list. Domain lists are reusable sets of domain
// specifications that you use in DNS firewall rules to allow, block, or alert on
// DNS queries to specific domains.
func route53globalresolver_CreateFirewallDomainList(cfg aws.Config, client *route53globalresolver.Client) {
	input := &route53globalresolver.CreateFirewallDomainListInput{
		// GlobalResolverId: *string, // Required
		// Name: *string, // Required
	}

	if len(_route53globalresolverGlobalResolverId) > 0 {
		input.GlobalResolverId = aws.String(_route53globalresolverGlobalResolverId)
	}
	if len(_route53globalresolverName) > 0 {
		input.Name = aws.String(_route53globalresolverName)
	}
	if len(_route53globalresolverClientToken) > 0 {
		input.ClientToken = aws.String(_route53globalresolverClientToken)
	}
	if len(_route53globalresolverDescription) > 0 {
		input.Description = aws.String(_route53globalresolverDescription)
	}
	if len(_route53globalresolverTags) > 0 {
		if err := assignInputField(input, "Tags", _route53globalresolverTags); err != nil {
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

// Creates a DNS firewall rule. Firewall rules define actions (ALLOW, BLOCK, or
// ALERT) to take on DNS queries that match specified domain lists, managed domain
// lists, or advanced threat protections.
func route53globalresolver_CreateFirewallRule(cfg aws.Config, client *route53globalresolver.Client) {
	input := &route53globalresolver.CreateFirewallRuleInput{
		// Action: types.FirewallRuleAction, // Required
		// DnsViewId: *string, // Required
		// Name: *string, // Required
	}

	if len(_route53globalresolverAction) > 0 {
		if err := assignInputField(input, "Action", _route53globalresolverAction); err != nil {
			log.Errorf("invalid --action: %s", err.Error())
			return
		}
	}
	if len(_route53globalresolverDnsViewId) > 0 {
		input.DnsViewId = aws.String(_route53globalresolverDnsViewId)
	}
	if len(_route53globalresolverName) > 0 {
		input.Name = aws.String(_route53globalresolverName)
	}
	if len(_route53globalresolverBlockOverrideDnsType) > 0 {
		if err := assignInputField(input, "BlockOverrideDnsType", _route53globalresolverBlockOverrideDnsType); err != nil {
			log.Errorf("invalid --block-override-dns-type: %s", err.Error())
			return
		}
	}
	if len(_route53globalresolverBlockOverrideDomain) > 0 {
		input.BlockOverrideDomain = aws.String(_route53globalresolverBlockOverrideDomain)
	}
	if len(_route53globalresolverBlockOverrideTtl) > 0 {
		if err := assignInputField(input, "BlockOverrideTtl", _route53globalresolverBlockOverrideTtl); err != nil {
			log.Errorf("invalid --block-override-ttl: %s", err.Error())
			return
		}
	}
	if len(_route53globalresolverBlockResponse) > 0 {
		if err := assignInputField(input, "BlockResponse", _route53globalresolverBlockResponse); err != nil {
			log.Errorf("invalid --block-response: %s", err.Error())
			return
		}
	}
	if len(_route53globalresolverClientToken) > 0 {
		input.ClientToken = aws.String(_route53globalresolverClientToken)
	}
	if len(_route53globalresolverConfidenceThreshold) > 0 {
		if err := assignInputField(input, "ConfidenceThreshold", _route53globalresolverConfidenceThreshold); err != nil {
			log.Errorf("invalid --confidence-threshold: %s", err.Error())
			return
		}
	}
	if len(_route53globalresolverDescription) > 0 {
		input.Description = aws.String(_route53globalresolverDescription)
	}
	if len(_route53globalresolverDnsAdvancedProtection) > 0 {
		if err := assignInputField(input, "DnsAdvancedProtection", _route53globalresolverDnsAdvancedProtection); err != nil {
			log.Errorf("invalid --dns-advanced-protection: %s", err.Error())
			return
		}
	}
	if len(_route53globalresolverFirewallDomainListId) > 0 {
		input.FirewallDomainListId = aws.String(_route53globalresolverFirewallDomainListId)
	}
	if len(_route53globalresolverPriority) > 0 {
		if err := assignInputField(input, "Priority", _route53globalresolverPriority); err != nil {
			log.Errorf("invalid --priority: %s", err.Error())
			return
		}
	}
	if len(_route53globalresolverQType) > 0 {
		input.QType = aws.String(_route53globalresolverQType)
	}

	if resp, err := client.CreateFirewallRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new Route 53 Global Resolver instance. A Route 53 Global Resolver is
// a global, internet-accessible DNS resolver that provides secure DNS resolution
// for both public and private domains through global anycast IP addresses.
func route53globalresolver_CreateGlobalResolver(cfg aws.Config, client *route53globalresolver.Client) {
	input := &route53globalresolver.CreateGlobalResolverInput{
		// Name: *string, // Required
		// Regions: []string, // Required
	}

	if len(_route53globalresolverName) > 0 {
		input.Name = aws.String(_route53globalresolverName)
	}
	if len(_route53globalresolverRegions) > 0 {
		input.Regions = append([]string(nil), _route53globalresolverRegions...)
	}
	if len(_route53globalresolverClientToken) > 0 {
		input.ClientToken = aws.String(_route53globalresolverClientToken)
	}
	if len(_route53globalresolverDescription) > 0 {
		input.Description = aws.String(_route53globalresolverDescription)
	}
	if len(_route53globalresolverObservabilityRegion) > 0 {
		input.ObservabilityRegion = aws.String(_route53globalresolverObservabilityRegion)
	}
	if len(_route53globalresolverTags) > 0 {
		if err := assignInputField(input, "Tags", _route53globalresolverTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateGlobalResolver(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an access source. This operation cannot be undone.
func route53globalresolver_DeleteAccessSource(cfg aws.Config, client *route53globalresolver.Client) {
	input := &route53globalresolver.DeleteAccessSourceInput{
		// AccessSourceId: *string, // Required
	}

	if len(_route53globalresolverAccessSourceId) > 0 {
		input.AccessSourceId = aws.String(_route53globalresolverAccessSourceId)
	}

	if resp, err := client.DeleteAccessSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an access token. This operation cannot be undone.
func route53globalresolver_DeleteAccessToken(cfg aws.Config, client *route53globalresolver.Client) {
	input := &route53globalresolver.DeleteAccessTokenInput{
		// AccessTokenId: *string, // Required
	}

	if len(_route53globalresolverAccessTokenId) > 0 {
		input.AccessTokenId = aws.String(_route53globalresolverAccessTokenId)
	}

	if resp, err := client.DeleteAccessToken(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a DNS view. This operation cannot be undone.
func route53globalresolver_DeleteDNSView(cfg aws.Config, client *route53globalresolver.Client) {
	input := &route53globalresolver.DeleteDNSViewInput{
		// DnsViewId: *string, // Required
	}

	if len(_route53globalresolverDnsViewId) > 0 {
		input.DnsViewId = aws.String(_route53globalresolverDnsViewId)
	}

	if resp, err := client.DeleteDNSView(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a firewall domain list. This operation cannot be undone.
func route53globalresolver_DeleteFirewallDomainList(cfg aws.Config, client *route53globalresolver.Client) {
	input := &route53globalresolver.DeleteFirewallDomainListInput{
		// FirewallDomainListId: *string, // Required
	}

	if len(_route53globalresolverFirewallDomainListId) > 0 {
		input.FirewallDomainListId = aws.String(_route53globalresolverFirewallDomainListId)
	}

	if resp, err := client.DeleteFirewallDomainList(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a DNS firewall rule. This operation cannot be undone.
func route53globalresolver_DeleteFirewallRule(cfg aws.Config, client *route53globalresolver.Client) {
	input := &route53globalresolver.DeleteFirewallRuleInput{
		// FirewallRuleId: *string, // Required
	}

	if len(_route53globalresolverFirewallRuleId) > 0 {
		input.FirewallRuleId = aws.String(_route53globalresolverFirewallRuleId)
	}

	if resp, err := client.DeleteFirewallRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a Route 53 Global Resolver instance. This operation cannot be undone.
// All associated DNS views, access sources, tokens, and firewall rules are also
// deleted.
func route53globalresolver_DeleteGlobalResolver(cfg aws.Config, client *route53globalresolver.Client) {
	input := &route53globalresolver.DeleteGlobalResolverInput{
		// GlobalResolverId: *string, // Required
	}

	if len(_route53globalresolverGlobalResolverId) > 0 {
		input.GlobalResolverId = aws.String(_route53globalresolverGlobalResolverId)
	}

	if resp, err := client.DeleteGlobalResolver(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disables a DNS view, preventing it from serving DNS queries.
func route53globalresolver_DisableDNSView(cfg aws.Config, client *route53globalresolver.Client) {
	input := &route53globalresolver.DisableDNSViewInput{
		// DnsViewId: *string, // Required
	}

	if len(_route53globalresolverDnsViewId) > 0 {
		input.DnsViewId = aws.String(_route53globalresolverDnsViewId)
	}

	if resp, err := client.DisableDNSView(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates a Route 53 private hosted zone from a Route 53 Global Resolver
// resource.
func route53globalresolver_DisassociateHostedZone(cfg aws.Config, client *route53globalresolver.Client) {
	input := &route53globalresolver.DisassociateHostedZoneInput{
		// HostedZoneId: *string, // Required
		// ResourceArn: *string, // Required
	}

	if len(_route53globalresolverHostedZoneId) > 0 {
		input.HostedZoneId = aws.String(_route53globalresolverHostedZoneId)
	}
	if len(_route53globalresolverResourceArn) > 0 {
		input.ResourceArn = aws.String(_route53globalresolverResourceArn)
	}

	if resp, err := client.DisassociateHostedZone(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables a disabled DNS view, allowing it to serve DNS queries again.
func route53globalresolver_EnableDNSView(cfg aws.Config, client *route53globalresolver.Client) {
	input := &route53globalresolver.EnableDNSViewInput{
		// DnsViewId: *string, // Required
	}

	if len(_route53globalresolverDnsViewId) > 0 {
		input.DnsViewId = aws.String(_route53globalresolverDnsViewId)
	}

	if resp, err := client.EnableDNSView(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about an access source.
func route53globalresolver_GetAccessSource(cfg aws.Config, client *route53globalresolver.Client) {
	input := &route53globalresolver.GetAccessSourceInput{
		// AccessSourceId: *string, // Required
	}

	if len(_route53globalresolverAccessSourceId) > 0 {
		input.AccessSourceId = aws.String(_route53globalresolverAccessSourceId)
	}

	if resp, err := client.GetAccessSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about an access token.
func route53globalresolver_GetAccessToken(cfg aws.Config, client *route53globalresolver.Client) {
	input := &route53globalresolver.GetAccessTokenInput{
		// AccessTokenId: *string, // Required
	}

	if len(_route53globalresolverAccessTokenId) > 0 {
		input.AccessTokenId = aws.String(_route53globalresolverAccessTokenId)
	}

	if resp, err := client.GetAccessToken(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a DNS view.
func route53globalresolver_GetDNSView(cfg aws.Config, client *route53globalresolver.Client) {
	input := &route53globalresolver.GetDNSViewInput{
		// DnsViewId: *string, // Required
	}

	if len(_route53globalresolverDnsViewId) > 0 {
		input.DnsViewId = aws.String(_route53globalresolverDnsViewId)
	}

	if resp, err := client.GetDNSView(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a firewall domain list.
func route53globalresolver_GetFirewallDomainList(cfg aws.Config, client *route53globalresolver.Client) {
	input := &route53globalresolver.GetFirewallDomainListInput{
		// FirewallDomainListId: *string, // Required
	}

	if len(_route53globalresolverFirewallDomainListId) > 0 {
		input.FirewallDomainListId = aws.String(_route53globalresolverFirewallDomainListId)
	}

	if resp, err := client.GetFirewallDomainList(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a DNS firewall rule.
func route53globalresolver_GetFirewallRule(cfg aws.Config, client *route53globalresolver.Client) {
	input := &route53globalresolver.GetFirewallRuleInput{
		// FirewallRuleId: *string, // Required
	}

	if len(_route53globalresolverFirewallRuleId) > 0 {
		input.FirewallRuleId = aws.String(_route53globalresolverFirewallRuleId)
	}

	if resp, err := client.GetFirewallRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a Route 53 Global Resolver instance.
func route53globalresolver_GetGlobalResolver(cfg aws.Config, client *route53globalresolver.Client) {
	input := &route53globalresolver.GetGlobalResolverInput{
		// GlobalResolverId: *string, // Required
	}

	if len(_route53globalresolverGlobalResolverId) > 0 {
		input.GlobalResolverId = aws.String(_route53globalresolverGlobalResolverId)
	}

	if resp, err := client.GetGlobalResolver(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a hosted zone association.
func route53globalresolver_GetHostedZoneAssociation(cfg aws.Config, client *route53globalresolver.Client) {
	input := &route53globalresolver.GetHostedZoneAssociationInput{
		// HostedZoneAssociationId: *string, // Required
	}

	if len(_route53globalresolverHostedZoneAssociationId) > 0 {
		input.HostedZoneAssociationId = aws.String(_route53globalresolverHostedZoneAssociationId)
	}

	if resp, err := client.GetHostedZoneAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about an AWS-managed firewall domain list. Managed domain
// lists contain domains associated with malicious activity, content categories, or
// specific threats.
func route53globalresolver_GetManagedFirewallDomainList(cfg aws.Config, client *route53globalresolver.Client) {
	input := &route53globalresolver.GetManagedFirewallDomainListInput{
		// ManagedFirewallDomainListId: *string, // Required
	}

	if len(_route53globalresolverManagedFirewallDomainListId) > 0 {
		input.ManagedFirewallDomainListId = aws.String(_route53globalresolverManagedFirewallDomainListId)
	}

	if resp, err := client.GetManagedFirewallDomainList(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Imports a list of domains from an Amazon S3 file into a firewall domain list.
// The file should contain one domain per line.
func route53globalresolver_ImportFirewallDomains(cfg aws.Config, client *route53globalresolver.Client) {
	input := &route53globalresolver.ImportFirewallDomainsInput{
		// DomainFileUrl: *string, // Required
		// FirewallDomainListId: *string, // Required
		// Operation: *string, // Required
	}

	if len(_route53globalresolverDomainFileUrl) > 0 {
		input.DomainFileUrl = aws.String(_route53globalresolverDomainFileUrl)
	}
	if len(_route53globalresolverFirewallDomainListId) > 0 {
		input.FirewallDomainListId = aws.String(_route53globalresolverFirewallDomainListId)
	}
	if len(_route53globalresolverOperation) > 0 {
		input.Operation = aws.String(_route53globalresolverOperation)
	}

	if resp, err := client.ImportFirewallDomains(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all access sources with pagination support.
func route53globalresolver_ListAccessSources(cfg aws.Config, client *route53globalresolver.Client) {
	input := &route53globalresolver.ListAccessSourcesInput{}

	if len(_route53globalresolverFilters) > 0 {
		if err := assignInputField(input, "Filters", _route53globalresolverFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_route53globalresolverMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _route53globalresolverMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_route53globalresolverNextToken) > 0 {
		input.NextToken = aws.String(_route53globalresolverNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAccessSources(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*route53globalresolver.ListAccessSourcesOutput
	p := route53globalresolver.NewListAccessSourcesPaginator(client, input)
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

// Lists all access tokens for a DNS view with pagination support.
func route53globalresolver_ListAccessTokens(cfg aws.Config, client *route53globalresolver.Client) {
	input := &route53globalresolver.ListAccessTokensInput{
		// DnsViewId: *string, // Required
	}

	if len(_route53globalresolverDnsViewId) > 0 {
		input.DnsViewId = aws.String(_route53globalresolverDnsViewId)
	}
	if len(_route53globalresolverFilters) > 0 {
		if err := assignInputField(input, "Filters", _route53globalresolverFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_route53globalresolverMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _route53globalresolverMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_route53globalresolverNextToken) > 0 {
		input.NextToken = aws.String(_route53globalresolverNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAccessTokens(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*route53globalresolver.ListAccessTokensOutput
	p := route53globalresolver.NewListAccessTokensPaginator(client, input)
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

// Lists all DNS views for a Route 53 Global Resolver with pagination support.
func route53globalresolver_ListDNSViews(cfg aws.Config, client *route53globalresolver.Client) {
	input := &route53globalresolver.ListDNSViewsInput{
		// GlobalResolverId: *string, // Required
	}

	if len(_route53globalresolverGlobalResolverId) > 0 {
		input.GlobalResolverId = aws.String(_route53globalresolverGlobalResolverId)
	}
	if len(_route53globalresolverMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _route53globalresolverMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_route53globalresolverNextToken) > 0 {
		input.NextToken = aws.String(_route53globalresolverNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDNSViews(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*route53globalresolver.ListDNSViewsOutput
	p := route53globalresolver.NewListDNSViewsPaginator(client, input)
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

// Lists all firewall domain lists for a Route 53 Global Resolver with pagination
// support.
func route53globalresolver_ListFirewallDomainLists(cfg aws.Config, client *route53globalresolver.Client) {
	input := &route53globalresolver.ListFirewallDomainListsInput{}

	if len(_route53globalresolverGlobalResolverId) > 0 {
		input.GlobalResolverId = aws.String(_route53globalresolverGlobalResolverId)
	}
	if len(_route53globalresolverMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _route53globalresolverMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_route53globalresolverNextToken) > 0 {
		input.NextToken = aws.String(_route53globalresolverNextToken)
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

	var results []*route53globalresolver.ListFirewallDomainListsOutput
	p := route53globalresolver.NewListFirewallDomainListsPaginator(client, input)
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

// Lists all the domains in DNS Firewall domain list you have created.
func route53globalresolver_ListFirewallDomains(cfg aws.Config, client *route53globalresolver.Client) {
	input := &route53globalresolver.ListFirewallDomainsInput{
		// FirewallDomainListId: *string, // Required
	}

	if len(_route53globalresolverFirewallDomainListId) > 0 {
		input.FirewallDomainListId = aws.String(_route53globalresolverFirewallDomainListId)
	}
	if len(_route53globalresolverMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _route53globalresolverMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_route53globalresolverNextToken) > 0 {
		input.NextToken = aws.String(_route53globalresolverNextToken)
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

	var results []*route53globalresolver.ListFirewallDomainsOutput
	p := route53globalresolver.NewListFirewallDomainsPaginator(client, input)
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

// Lists all DNS firewall rules for a DNS view with pagination support.
func route53globalresolver_ListFirewallRules(cfg aws.Config, client *route53globalresolver.Client) {
	input := &route53globalresolver.ListFirewallRulesInput{
		// DnsViewId: *string, // Required
	}

	if len(_route53globalresolverDnsViewId) > 0 {
		input.DnsViewId = aws.String(_route53globalresolverDnsViewId)
	}
	if len(_route53globalresolverFilters) > 0 {
		if err := assignInputField(input, "Filters", _route53globalresolverFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_route53globalresolverMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _route53globalresolverMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_route53globalresolverNextToken) > 0 {
		input.NextToken = aws.String(_route53globalresolverNextToken)
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

	var results []*route53globalresolver.ListFirewallRulesOutput
	p := route53globalresolver.NewListFirewallRulesPaginator(client, input)
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

// Lists all Route 53 Global Resolver instances in your account with pagination
// support.
func route53globalresolver_ListGlobalResolvers(cfg aws.Config, client *route53globalresolver.Client) {
	input := &route53globalresolver.ListGlobalResolversInput{}

	if len(_route53globalresolverMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _route53globalresolverMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_route53globalresolverNextToken) > 0 {
		input.NextToken = aws.String(_route53globalresolverNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListGlobalResolvers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*route53globalresolver.ListGlobalResolversOutput
	p := route53globalresolver.NewListGlobalResolversPaginator(client, input)
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

// Lists all hosted zone associations for a Route 53 Global Resolver resource with
// pagination support.
func route53globalresolver_ListHostedZoneAssociations(cfg aws.Config, client *route53globalresolver.Client) {
	input := &route53globalresolver.ListHostedZoneAssociationsInput{
		// ResourceArn: *string, // Required
	}

	if len(_route53globalresolverResourceArn) > 0 {
		input.ResourceArn = aws.String(_route53globalresolverResourceArn)
	}
	if len(_route53globalresolverMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _route53globalresolverMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_route53globalresolverNextToken) > 0 {
		input.NextToken = aws.String(_route53globalresolverNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListHostedZoneAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*route53globalresolver.ListHostedZoneAssociationsOutput
	p := route53globalresolver.NewListHostedZoneAssociationsPaginator(client, input)
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

// Returns a paginated list of the AWS Managed DNS Lists and the categories for
// DNS Firewall. The categories are either THREAT or CONTENT .
func route53globalresolver_ListManagedFirewallDomainLists(cfg aws.Config, client *route53globalresolver.Client) {
	input := &route53globalresolver.ListManagedFirewallDomainListsInput{
		// ManagedFirewallDomainListType: *string, // Required
	}

	if len(_route53globalresolverManagedFirewallDomainListType) > 0 {
		input.ManagedFirewallDomainListType = aws.String(_route53globalresolverManagedFirewallDomainListType)
	}
	if len(_route53globalresolverMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _route53globalresolverMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_route53globalresolverNextToken) > 0 {
		input.NextToken = aws.String(_route53globalresolverNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListManagedFirewallDomainLists(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*route53globalresolver.ListManagedFirewallDomainListsOutput
	p := route53globalresolver.NewListManagedFirewallDomainListsPaginator(client, input)
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

// Lists the tags associated with a Route 53 Global Resolver resource.
func route53globalresolver_ListTagsForResource(cfg aws.Config, client *route53globalresolver.Client) {
	input := &route53globalresolver.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_route53globalresolverResourceArn) > 0 {
		input.ResourceArn = aws.String(_route53globalresolverResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds or updates tags for a Route 53 Global Resolver resource. Tags are
// key-value pairs that help you organize and identify your resources.
func route53globalresolver_TagResource(cfg aws.Config, client *route53globalresolver.Client) {
	input := &route53globalresolver.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_route53globalresolverResourceArn) > 0 {
		input.ResourceArn = aws.String(_route53globalresolverResourceArn)
	}
	if len(_route53globalresolverTags) > 0 {
		if err := assignInputField(input, "Tags", _route53globalresolverTags); err != nil {
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

// Removes tags from a Route 53 Global Resolver resource.
func route53globalresolver_UntagResource(cfg aws.Config, client *route53globalresolver.Client) {
	input := &route53globalresolver.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_route53globalresolverResourceArn) > 0 {
		input.ResourceArn = aws.String(_route53globalresolverResourceArn)
	}
	if len(_route53globalresolverTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _route53globalresolverTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration of an access source.
func route53globalresolver_UpdateAccessSource(cfg aws.Config, client *route53globalresolver.Client) {
	input := &route53globalresolver.UpdateAccessSourceInput{
		// AccessSourceId: *string, // Required
	}

	if len(_route53globalresolverAccessSourceId) > 0 {
		input.AccessSourceId = aws.String(_route53globalresolverAccessSourceId)
	}
	if len(_route53globalresolverCidr) > 0 {
		input.Cidr = aws.String(_route53globalresolverCidr)
	}
	if len(_route53globalresolverIpAddressType) > 0 {
		if err := assignInputField(input, "IpAddressType", _route53globalresolverIpAddressType); err != nil {
			log.Errorf("invalid --ip-address-type: %s", err.Error())
			return
		}
	}
	if len(_route53globalresolverName) > 0 {
		input.Name = aws.String(_route53globalresolverName)
	}
	if len(_route53globalresolverProtocol) > 0 {
		if err := assignInputField(input, "Protocol", _route53globalresolverProtocol); err != nil {
			log.Errorf("invalid --protocol: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateAccessSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration of an access token.
func route53globalresolver_UpdateAccessToken(cfg aws.Config, client *route53globalresolver.Client) {
	input := &route53globalresolver.UpdateAccessTokenInput{
		// AccessTokenId: *string, // Required
		// Name: *string, // Required
	}

	if len(_route53globalresolverAccessTokenId) > 0 {
		input.AccessTokenId = aws.String(_route53globalresolverAccessTokenId)
	}
	if len(_route53globalresolverName) > 0 {
		input.Name = aws.String(_route53globalresolverName)
	}

	if resp, err := client.UpdateAccessToken(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration of a DNS view.
func route53globalresolver_UpdateDNSView(cfg aws.Config, client *route53globalresolver.Client) {
	input := &route53globalresolver.UpdateDNSViewInput{
		// DnsViewId: *string, // Required
	}

	if len(_route53globalresolverDnsViewId) > 0 {
		input.DnsViewId = aws.String(_route53globalresolverDnsViewId)
	}
	if len(_route53globalresolverDescription) > 0 {
		input.Description = aws.String(_route53globalresolverDescription)
	}
	if len(_route53globalresolverDnssecValidation) > 0 {
		if err := assignInputField(input, "DnssecValidation", _route53globalresolverDnssecValidation); err != nil {
			log.Errorf("invalid --dnssec-validation: %s", err.Error())
			return
		}
	}
	if len(_route53globalresolverEdnsClientSubnet) > 0 {
		if err := assignInputField(input, "EdnsClientSubnet", _route53globalresolverEdnsClientSubnet); err != nil {
			log.Errorf("invalid --edns-client-subnet: %s", err.Error())
			return
		}
	}
	if len(_route53globalresolverFirewallRulesFailOpen) > 0 {
		if err := assignInputField(input, "FirewallRulesFailOpen", _route53globalresolverFirewallRulesFailOpen); err != nil {
			log.Errorf("invalid --firewall-rules-fail-open: %s", err.Error())
			return
		}
	}
	if len(_route53globalresolverName) > 0 {
		input.Name = aws.String(_route53globalresolverName)
	}

	if resp, err := client.UpdateDNSView(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a DNS Firewall domain list from an array of specified domains.
func route53globalresolver_UpdateFirewallDomains(cfg aws.Config, client *route53globalresolver.Client) {
	input := &route53globalresolver.UpdateFirewallDomainsInput{
		// Domains: []string, // Required
		// FirewallDomainListId: *string, // Required
		// Operation: *string, // Required
	}

	if len(_route53globalresolverDomains) > 0 {
		input.Domains = append([]string(nil), _route53globalresolverDomains...)
	}
	if len(_route53globalresolverFirewallDomainListId) > 0 {
		input.FirewallDomainListId = aws.String(_route53globalresolverFirewallDomainListId)
	}
	if len(_route53globalresolverOperation) > 0 {
		input.Operation = aws.String(_route53globalresolverOperation)
	}

	if resp, err := client.UpdateFirewallDomains(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration of a DNS firewall rule.
func route53globalresolver_UpdateFirewallRule(cfg aws.Config, client *route53globalresolver.Client) {
	input := &route53globalresolver.UpdateFirewallRuleInput{
		// ClientToken: *string, // Required
		// FirewallRuleId: *string, // Required
	}

	if len(_route53globalresolverClientToken) > 0 {
		input.ClientToken = aws.String(_route53globalresolverClientToken)
	}
	if len(_route53globalresolverFirewallRuleId) > 0 {
		input.FirewallRuleId = aws.String(_route53globalresolverFirewallRuleId)
	}
	if len(_route53globalresolverAction) > 0 {
		if err := assignInputField(input, "Action", _route53globalresolverAction); err != nil {
			log.Errorf("invalid --action: %s", err.Error())
			return
		}
	}
	if len(_route53globalresolverBlockOverrideDnsType) > 0 {
		if err := assignInputField(input, "BlockOverrideDnsType", _route53globalresolverBlockOverrideDnsType); err != nil {
			log.Errorf("invalid --block-override-dns-type: %s", err.Error())
			return
		}
	}
	if len(_route53globalresolverBlockOverrideDomain) > 0 {
		input.BlockOverrideDomain = aws.String(_route53globalresolverBlockOverrideDomain)
	}
	if len(_route53globalresolverBlockOverrideTtl) > 0 {
		if err := assignInputField(input, "BlockOverrideTtl", _route53globalresolverBlockOverrideTtl); err != nil {
			log.Errorf("invalid --block-override-ttl: %s", err.Error())
			return
		}
	}
	if len(_route53globalresolverBlockResponse) > 0 {
		if err := assignInputField(input, "BlockResponse", _route53globalresolverBlockResponse); err != nil {
			log.Errorf("invalid --block-response: %s", err.Error())
			return
		}
	}
	if len(_route53globalresolverConfidenceThreshold) > 0 {
		if err := assignInputField(input, "ConfidenceThreshold", _route53globalresolverConfidenceThreshold); err != nil {
			log.Errorf("invalid --confidence-threshold: %s", err.Error())
			return
		}
	}
	if len(_route53globalresolverDescription) > 0 {
		input.Description = aws.String(_route53globalresolverDescription)
	}
	if len(_route53globalresolverDnsAdvancedProtection) > 0 {
		if err := assignInputField(input, "DnsAdvancedProtection", _route53globalresolverDnsAdvancedProtection); err != nil {
			log.Errorf("invalid --dns-advanced-protection: %s", err.Error())
			return
		}
	}
	if len(_route53globalresolverName) > 0 {
		input.Name = aws.String(_route53globalresolverName)
	}
	if len(_route53globalresolverPriority) > 0 {
		if err := assignInputField(input, "Priority", _route53globalresolverPriority); err != nil {
			log.Errorf("invalid --priority: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateFirewallRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration of a Route 53 Global Resolver instance. You can
// modify the name, description, and observability region.
func route53globalresolver_UpdateGlobalResolver(cfg aws.Config, client *route53globalresolver.Client) {
	input := &route53globalresolver.UpdateGlobalResolverInput{
		// GlobalResolverId: *string, // Required
	}

	if len(_route53globalresolverGlobalResolverId) > 0 {
		input.GlobalResolverId = aws.String(_route53globalresolverGlobalResolverId)
	}
	if len(_route53globalresolverDescription) > 0 {
		input.Description = aws.String(_route53globalresolverDescription)
	}
	if len(_route53globalresolverName) > 0 {
		input.Name = aws.String(_route53globalresolverName)
	}
	if len(_route53globalresolverObservabilityRegion) > 0 {
		input.ObservabilityRegion = aws.String(_route53globalresolverObservabilityRegion)
	}

	if resp, err := client.UpdateGlobalResolver(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration of a hosted zone association.
func route53globalresolver_UpdateHostedZoneAssociation(cfg aws.Config, client *route53globalresolver.Client) {
	input := &route53globalresolver.UpdateHostedZoneAssociationInput{
		// HostedZoneAssociationId: *string, // Required
	}

	if len(_route53globalresolverHostedZoneAssociationId) > 0 {
		input.HostedZoneAssociationId = aws.String(_route53globalresolverHostedZoneAssociationId)
	}
	if len(_route53globalresolverName) > 0 {
		input.Name = aws.String(_route53globalresolverName)
	}

	if resp, err := client.UpdateHostedZoneAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_route53globalresolverCmd)
	_route53globalresolverCmd.Flags().SortFlags = false

	_route53globalresolverCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_route53globalresolverCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_route53globalresolverCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_route53globalresolverCmd.Flags().StringVarP(&_route53globalresolverAccessSourceId, "access-source-id", "", "", "Access Source ID")
	_route53globalresolverCmd.Flags().StringVarP(&_route53globalresolverAccessTokenId, "access-token-id", "", "", "Access Token ID")
	_route53globalresolverCmd.Flags().StringVarP(&_route53globalresolverAction, "action", "", "", "Action")
	_route53globalresolverCmd.Flags().StringVarP(&_route53globalresolverBlockOverrideDnsType, "block-override-dns-type", "", "", "Block Override DNS Type")
	_route53globalresolverCmd.Flags().StringVarP(&_route53globalresolverBlockOverrideDomain, "block-override-domain", "", "", "Block Override Domain")
	_route53globalresolverCmd.Flags().StringVarP(&_route53globalresolverBlockOverrideTtl, "block-override-ttl", "", "", "Block Override TTL")
	_route53globalresolverCmd.Flags().StringVarP(&_route53globalresolverBlockResponse, "block-response", "", "", "Block Response")
	_route53globalresolverCmd.Flags().StringVarP(&_route53globalresolverCidr, "cidr", "", "", "CIDR")
	_route53globalresolverCmd.Flags().StringVarP(&_route53globalresolverClientToken, "client-token", "", "", "Client Token")
	_route53globalresolverCmd.Flags().StringVarP(&_route53globalresolverConfidenceThreshold, "confidence-threshold", "", "", "Confidence Threshold")
	_route53globalresolverCmd.Flags().StringVarP(&_route53globalresolverDescription, "description", "", "", "Description")
	_route53globalresolverCmd.Flags().StringVarP(&_route53globalresolverDnsAdvancedProtection, "dns-advanced-protection", "", "", "DNS Advanced Protection")
	_route53globalresolverCmd.Flags().StringVarP(&_route53globalresolverDnsViewId, "dns-view-id", "", "", "DNS View ID")
	_route53globalresolverCmd.Flags().StringVarP(&_route53globalresolverDnssecValidation, "dnssec-validation", "", "", "Dnssec Validation")
	_route53globalresolverCmd.Flags().StringVarP(&_route53globalresolverDomainFileUrl, "domain-file-url", "", "", "Domain File URL")
	_route53globalresolverCmd.Flags().StringSliceVarP(&_route53globalresolverDomains, "domains", "", nil, "Domains")
	_route53globalresolverCmd.Flags().StringVarP(&_route53globalresolverEdnsClientSubnet, "edns-client-subnet", "", "", "Edns Client Subnet")
	_route53globalresolverCmd.Flags().StringVarP(&_route53globalresolverExpiresAt, "expires-at", "", "", "Expires At")
	_route53globalresolverCmd.Flags().StringVarP(&_route53globalresolverFilters, "filters", "", "", "Filters")
	_route53globalresolverCmd.Flags().StringVarP(&_route53globalresolverFirewallDomainListId, "firewall-domain-list-id", "", "", "Firewall Domain List ID")
	_route53globalresolverCmd.Flags().StringVarP(&_route53globalresolverFirewallRuleId, "firewall-rule-id", "", "", "Firewall Rule ID")
	_route53globalresolverCmd.Flags().StringVarP(&_route53globalresolverFirewallRules, "firewall-rules", "", "", "Firewall Rules")
	_route53globalresolverCmd.Flags().StringVarP(&_route53globalresolverFirewallRulesFailOpen, "firewall-rules-fail-open", "", "", "Firewall Rules Fail Open")
	_route53globalresolverCmd.Flags().StringVarP(&_route53globalresolverGlobalResolverId, "global-resolver-id", "", "", "Global Resolver ID")
	_route53globalresolverCmd.Flags().StringVarP(&_route53globalresolverHostedZoneAssociationId, "hosted-zone-association-id", "", "", "Hosted Zone Association ID")
	_route53globalresolverCmd.Flags().StringVarP(&_route53globalresolverHostedZoneId, "hosted-zone-id", "", "", "Hosted Zone ID")
	_route53globalresolverCmd.Flags().StringVarP(&_route53globalresolverIpAddressType, "ip-address-type", "", "", "IP Address Type")
	_route53globalresolverCmd.Flags().StringVarP(&_route53globalresolverManagedFirewallDomainListId, "managed-firewall-domain-list-id", "", "", "Managed Firewall Domain List ID")
	_route53globalresolverCmd.Flags().StringVarP(&_route53globalresolverManagedFirewallDomainListType, "managed-firewall-domain-list-type", "", "", "Managed Firewall Domain List Type")
	_route53globalresolverCmd.Flags().StringVarP(&_route53globalresolverMaxResults, "max-results", "", "", "Max Results")
	_route53globalresolverCmd.Flags().StringVarP(&_route53globalresolverName, "name", "", "", "Name")
	_route53globalresolverCmd.Flags().StringVarP(&_route53globalresolverNextToken, "next-token", "", "", "Next Token")
	_route53globalresolverCmd.Flags().StringVarP(&_route53globalresolverObservabilityRegion, "observability-region", "", "", "Observability Region")
	_route53globalresolverCmd.Flags().StringVarP(&_route53globalresolverOperation, "operation", "", "", "Operation")
	_route53globalresolverCmd.Flags().StringVarP(&_route53globalresolverPriority, "priority", "", "", "Priority")
	_route53globalresolverCmd.Flags().StringVarP(&_route53globalresolverProtocol, "protocol", "", "", "Protocol")
	_route53globalresolverCmd.Flags().StringVarP(&_route53globalresolverQType, "qtype", "", "", "Qtype")
	_route53globalresolverCmd.Flags().StringSliceVarP(&_route53globalresolverRegions, "regions", "", nil, "Regions")
	_route53globalresolverCmd.Flags().StringVarP(&_route53globalresolverResourceArn, "resource-arn", "", "", "Resource ARN")
	_route53globalresolverCmd.Flags().StringSliceVarP(&_route53globalresolverTagKeys, "tag-keys", "", nil, "Tag Keys")
	_route53globalresolverCmd.Flags().StringVarP(&_route53globalresolverTags, "tags", "", "", "Tags")

	_route53globalresolverCmd.Flags().BoolVarP(&_route53globalresolverAssociateHostedZone, "associate-hosted-zone", "", false, "Associate Hosted Zone")
	_route53globalresolverCmd.Flags().BoolVarP(&_route53globalresolverBatchCreateFirewallRule, "batch-create-firewall-rule", "", false, "Batch Create Firewall Rule")
	_route53globalresolverCmd.Flags().BoolVarP(&_route53globalresolverBatchDeleteFirewallRule, "batch-delete-firewall-rule", "", false, "Batch Delete Firewall Rule")
	_route53globalresolverCmd.Flags().BoolVarP(&_route53globalresolverBatchUpdateFirewallRule, "batch-update-firewall-rule", "", false, "Batch Update Firewall Rule")
	_route53globalresolverCmd.Flags().BoolVarP(&_route53globalresolverCreateAccessSource, "create-access-source", "", false, "Create Access Source")
	_route53globalresolverCmd.Flags().BoolVarP(&_route53globalresolverCreateAccessToken, "create-access-token", "", false, "Create Access Token")
	_route53globalresolverCmd.Flags().BoolVarP(&_route53globalresolverCreateDNSView, "create-dns-view", "", false, "Create DNS View")
	_route53globalresolverCmd.Flags().BoolVarP(&_route53globalresolverCreateFirewallDomainList, "create-firewall-domain-list", "", false, "Create Firewall Domain List")
	_route53globalresolverCmd.Flags().BoolVarP(&_route53globalresolverCreateFirewallRule, "create-firewall-rule", "", false, "Create Firewall Rule")
	_route53globalresolverCmd.Flags().BoolVarP(&_route53globalresolverCreateGlobalResolver, "create-global-resolver", "", false, "Create Global Resolver")
	_route53globalresolverCmd.Flags().BoolVarP(&_route53globalresolverDeleteAccessSource, "delete-access-source", "", false, "Delete Access Source")
	_route53globalresolverCmd.Flags().BoolVarP(&_route53globalresolverDeleteAccessToken, "delete-access-token", "", false, "Delete Access Token")
	_route53globalresolverCmd.Flags().BoolVarP(&_route53globalresolverDeleteDNSView, "delete-dns-view", "", false, "Delete DNS View")
	_route53globalresolverCmd.Flags().BoolVarP(&_route53globalresolverDeleteFirewallDomainList, "delete-firewall-domain-list", "", false, "Delete Firewall Domain List")
	_route53globalresolverCmd.Flags().BoolVarP(&_route53globalresolverDeleteFirewallRule, "delete-firewall-rule", "", false, "Delete Firewall Rule")
	_route53globalresolverCmd.Flags().BoolVarP(&_route53globalresolverDeleteGlobalResolver, "delete-global-resolver", "", false, "Delete Global Resolver")
	_route53globalresolverCmd.Flags().BoolVarP(&_route53globalresolverDisableDNSView, "disable-dns-view", "", false, "Disable DNS View")
	_route53globalresolverCmd.Flags().BoolVarP(&_route53globalresolverDisassociateHostedZone, "disassociate-hosted-zone", "", false, "Disassociate Hosted Zone")
	_route53globalresolverCmd.Flags().BoolVarP(&_route53globalresolverEnableDNSView, "enable-dns-view", "", false, "Enable DNS View")
	_route53globalresolverCmd.Flags().BoolVarP(&_route53globalresolverGetAccessSource, "get-access-source", "", false, "Get Access Source")
	_route53globalresolverCmd.Flags().BoolVarP(&_route53globalresolverGetAccessToken, "get-access-token", "", false, "Get Access Token")
	_route53globalresolverCmd.Flags().BoolVarP(&_route53globalresolverGetDNSView, "get-dns-view", "", false, "Get DNS View")
	_route53globalresolverCmd.Flags().BoolVarP(&_route53globalresolverGetFirewallDomainList, "get-firewall-domain-list", "", false, "Get Firewall Domain List")
	_route53globalresolverCmd.Flags().BoolVarP(&_route53globalresolverGetFirewallRule, "get-firewall-rule", "", false, "Get Firewall Rule")
	_route53globalresolverCmd.Flags().BoolVarP(&_route53globalresolverGetGlobalResolver, "get-global-resolver", "", false, "Get Global Resolver")
	_route53globalresolverCmd.Flags().BoolVarP(&_route53globalresolverGetHostedZoneAssociation, "get-hosted-zone-association", "", false, "Get Hosted Zone Association")
	_route53globalresolverCmd.Flags().BoolVarP(&_route53globalresolverGetManagedFirewallDomainList, "get-managed-firewall-domain-list", "", false, "Get Managed Firewall Domain List")
	_route53globalresolverCmd.Flags().BoolVarP(&_route53globalresolverImportFirewallDomains, "import-firewall-domains", "", false, "Import Firewall Domains")
	_route53globalresolverCmd.Flags().BoolVarP(&_route53globalresolverListAccessSources, "list-access-sources", "", false, "List Access Sources")
	_route53globalresolverCmd.Flags().BoolVarP(&_route53globalresolverListAccessTokens, "list-access-tokens", "", false, "List Access Tokens")
	_route53globalresolverCmd.Flags().BoolVarP(&_route53globalresolverListDNSViews, "list-dns-views", "", false, "List DNS Views")
	_route53globalresolverCmd.Flags().BoolVarP(&_route53globalresolverListFirewallDomainLists, "list-firewall-domain-lists", "", false, "List Firewall Domain Lists")
	_route53globalresolverCmd.Flags().BoolVarP(&_route53globalresolverListFirewallDomains, "list-firewall-domains", "", false, "List Firewall Domains")
	_route53globalresolverCmd.Flags().BoolVarP(&_route53globalresolverListFirewallRules, "list-firewall-rules", "", false, "List Firewall Rules")
	_route53globalresolverCmd.Flags().BoolVarP(&_route53globalresolverListGlobalResolvers, "list-global-resolvers", "", false, "List Global Resolvers")
	_route53globalresolverCmd.Flags().BoolVarP(&_route53globalresolverListHostedZoneAssociations, "list-hosted-zone-associations", "", false, "List Hosted Zone Associations")
	_route53globalresolverCmd.Flags().BoolVarP(&_route53globalresolverListManagedFirewallDomainLists, "list-managed-firewall-domain-lists", "", false, "List Managed Firewall Domain Lists")
	_route53globalresolverCmd.Flags().BoolVarP(&_route53globalresolverListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_route53globalresolverCmd.Flags().BoolVarP(&_route53globalresolverTagResource, "tag-resource", "", false, "Tag Resource")
	_route53globalresolverCmd.Flags().BoolVarP(&_route53globalresolverUntagResource, "untag-resource", "", false, "Untag Resource")
	_route53globalresolverCmd.Flags().BoolVarP(&_route53globalresolverUpdateAccessSource, "update-access-source", "", false, "Update Access Source")
	_route53globalresolverCmd.Flags().BoolVarP(&_route53globalresolverUpdateAccessToken, "update-access-token", "", false, "Update Access Token")
	_route53globalresolverCmd.Flags().BoolVarP(&_route53globalresolverUpdateDNSView, "update-dns-view", "", false, "Update DNS View")
	_route53globalresolverCmd.Flags().BoolVarP(&_route53globalresolverUpdateFirewallDomains, "update-firewall-domains", "", false, "Update Firewall Domains")
	_route53globalresolverCmd.Flags().BoolVarP(&_route53globalresolverUpdateFirewallRule, "update-firewall-rule", "", false, "Update Firewall Rule")
	_route53globalresolverCmd.Flags().BoolVarP(&_route53globalresolverUpdateGlobalResolver, "update-global-resolver", "", false, "Update Global Resolver")
	_route53globalresolverCmd.Flags().BoolVarP(&_route53globalresolverUpdateHostedZoneAssociation, "update-hosted-zone-association", "", false, "Update Hosted Zone Association")

}
